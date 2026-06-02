/** Load the metric catalog from Postgres into in-memory maps for the normalizer. */
import type { Client } from "pg";
import type { Catalog, Conversion } from "@/lib/ingest/normalize";

export async function loadCatalog(db: Client): Promise<Catalog> {
  const aliases = new Map<string, string>();
  const metrics = new Map<string, { canonical_unit: string }>();
  const conversions = new Map<string, Conversion[]>();

  const a = await db.query<{ alias_norm: string; metric_id: string }>(
    "select alias_norm, metric_id from metric_aliases",
  );
  for (const r of a.rows) aliases.set(r.alias_norm, r.metric_id);

  const m = await db.query<{ id: string; canonical_unit: string }>(
    "select id, canonical_unit from metrics",
  );
  for (const r of m.rows) metrics.set(r.id, { canonical_unit: r.canonical_unit });

  const c = await db.query<{
    metric_id: string | null;
    from_unit: string;
    to_unit: string;
    factor: string;
    offset: string;
  }>('select metric_id, from_unit, to_unit, factor, "offset" from unit_conversions');
  for (const r of c.rows) {
    const key = r.metric_id ?? "*";
    const list = conversions.get(key) ?? [];
    list.push({
      from_unit: r.from_unit,
      to_unit: r.to_unit,
      factor: Number(r.factor),
      offset: Number(r.offset),
    });
    conversions.set(key, list);
  }
  return { aliases, metrics, conversions };
}
