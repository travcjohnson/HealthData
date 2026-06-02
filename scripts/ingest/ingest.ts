/**
 * End-to-end ingestion for one file:
 *   detect type → parse → source_document → extraction → staging → normalize
 *   → observations (canonical, deduped). Every datapoint keeps full provenance.
 */
import { Client } from "pg";
import { createHash } from "node:crypto";
import { readFileSync } from "node:fs";
import { basename, extname } from "node:path";
import { loadCatalog } from "./catalog";
import { normalizeItem, type ExtractedItem } from "@/lib/ingest/normalize";
import { parseAppleHealth } from "./parsers/appleHealth";
import { parseWearableCsv } from "./parsers/wearableCsv";
import { extractFromPdf, extractFromText, PROMPT_VERSION } from "./llm";

export type IngestResult = {
  documentId: string | null;
  sourceType: string;
  duplicate: boolean;
  extracted: number;
  matched: number;
  unmatched: number;
  inserted: number;
};

type Parsed = {
  sourceType: string;
  items: ExtractedItem[];
  extractor: string;
  extractorKind: "llm" | "deterministic";
  sourceLab: string | null;
  raw: unknown;
};

async function parseFile(path: string, buf: Buffer): Promise<Parsed> {
  const ext = extname(path).toLowerCase();
  const name = basename(path).toLowerCase();
  if (ext === ".xml" || name.includes("export")) {
    return { sourceType: "apple_health", items: parseAppleHealth(buf.toString("utf8")),
             extractor: "apple-xml-parser-v1", extractorKind: "deterministic", sourceLab: null, raw: null };
  }
  if (ext === ".csv") {
    return { sourceType: "oura", items: parseWearableCsv(buf.toString("utf8")),
             extractor: "wearable-csv-v1", extractorKind: "deterministic", sourceLab: null, raw: null };
  }
  if (ext === ".pdf") {
    const e = await extractFromPdf(buf);
    return { sourceType: "lab_pdf", items: e.items, extractor: e.extractor,
             extractorKind: e.extractorKind, sourceLab: e.sourceLab, raw: e.raw };
  }
  // treat anything else as text (e.g. pasted lab text, .txt)
  const e = await extractFromText(buf.toString("utf8"));
  return { sourceType: "lab_pdf", items: e.items, extractor: e.extractor,
           extractorKind: e.extractorKind, sourceLab: e.sourceLab, raw: e.raw };
}

export async function ingestFile(
  db: Client,
  ownerId: string,
  path: string,
  opts: { storagePath?: string } = {},
): Promise<IngestResult> {
  const buf = readFileSync(path);
  const sha = createHash("sha256").update(buf).digest("hex");
  const parsed = await parseFile(path, buf);

  // 1) source_documents (dedup on content hash)
  const doc = await db.query(
    `insert into source_documents
       (owner_id, source_type, storage_path, original_name, content_sha256, byte_size, source_lab, status)
     values ($1,$2,$3,$4,$5,$6,$7,'staged')
     on conflict (owner_id, content_sha256) do nothing
     returning id`,
    [ownerId, parsed.sourceType, opts.storagePath ?? `local://${basename(path)}`,
     basename(path), sha, buf.byteLength, parsed.sourceLab],
  );
  if (doc.rowCount === 0) {
    return { documentId: null, sourceType: parsed.sourceType, duplicate: true,
             extracted: 0, matched: 0, unmatched: 0, inserted: 0 };
  }
  const documentId = doc.rows[0].id as string;

  // 2) extraction run
  const ext = await db.query(
    `insert into extractions (owner_id, document_id, extractor, extractor_kind, prompt_version, raw_output)
     values ($1,$2,$3,$4,$5,$6) returning id`,
    [ownerId, documentId, parsed.extractor, parsed.extractorKind, PROMPT_VERSION,
     parsed.raw ? JSON.stringify(parsed.raw) : null],
  );
  const extractionId = ext.rows[0].id as string;

  // 3) normalize + stage + promote
  const catalog = await loadCatalog(db);
  let matched = 0, unmatched = 0, inserted = 0;

  for (const item of parsed.items) {
    const n = normalizeItem(item, catalog);
    if (n.resolveStatus === "matched") matched++; else unmatched++;

    const stage = await db.query(
      `insert into staging_observations
         (owner_id, extraction_id, document_id, raw_label, raw_value, raw_unit, ref_text,
          abnormal_flag, effective_at, confidence, resolved_metric_id, resolve_status)
       values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) returning id`,
      [ownerId, extractionId, documentId, n.rawLabel, String(item.value ?? ""), n.rawUnit,
       n.refText, item.flag ?? null, n.effectiveAt, n.confidence, n.metricId,
       n.metricId ? "matched" : "unmatched"],
    );
    const stagingId = stage.rows[0].id as string;

    if (n.metricId && n.valueNum != null && n.effectiveAt) {
      const obs = await db.query(
        `insert into observations
           (owner_id, metric_id, effective_at, value_num, value_unit,
            raw_value_num, raw_unit, raw_label, ref_low, ref_high, ref_text, abnormal_flag,
            document_id, extraction_id, source_type, confidence, status)
         values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,'normalized')
         on conflict (owner_id, metric_id, effective_at, value_num) where status='normalized'
         do nothing
         returning id`,
        [ownerId, n.metricId, n.effectiveAt, n.valueNum, n.valueUnit,
         n.rawValueNum, n.rawUnit, n.rawLabel, n.refLow, n.refHigh, n.refText, item.flag ?? null,
         documentId, extractionId, parsed.sourceType, n.confidence],
      );
      if (obs.rowCount && obs.rowCount > 0) {
        inserted++;
        await db.query("update staging_observations set promoted_observation_id=$1 where id=$2",
          [obs.rows[0].id, stagingId]);
      }
    }
  }

  await db.query("update source_documents set status='normalized' where id=$1", [documentId]);
  return { documentId, sourceType: parsed.sourceType, duplicate: false,
           extracted: parsed.items.length, matched, unmatched, inserted };
}
