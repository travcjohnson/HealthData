/**
 * CLI: ingest one or more health data files into the database.
 *
 *   tsx scripts/ingest/cli.ts --owner <uuid> path/to/export.xml labs.pdf oura.csv
 *
 * Detects Apple Health XML, wearable CSV, and lab PDFs/text automatically.
 * Uses DATABASE_URL (service-role / direct connection — bypasses RLS).
 */
import { Client } from "pg";
import { ingestFile } from "./ingest";

async function main() {
  const args = process.argv.slice(2);
  const ownerIdx = args.indexOf("--owner");
  const ownerId = ownerIdx >= 0 ? args[ownerIdx + 1] : process.env.INGEST_OWNER_ID;
  const files = args.filter((a, i) => !a.startsWith("--") && i !== ownerIdx + 1);

  if (!ownerId) throw new Error("Provide --owner <uuid> (or set INGEST_OWNER_ID)");
  if (files.length === 0) throw new Error("Provide at least one file to ingest");

  const db = new Client({ connectionString: process.env.DATABASE_URL });
  await db.connect();
  try {
    for (const f of files) {
      const r = await ingestFile(db, ownerId, f);
      if (r.duplicate) {
        console.log(`• ${f}: already ingested (duplicate) — skipped`);
      } else {
        console.log(
          `• ${f}: ${r.sourceType} — extracted ${r.extracted}, matched ${r.matched}, ` +
            `unmatched ${r.unmatched}, inserted ${r.inserted}`,
        );
      }
    }
  } finally {
    await db.end();
  }
}

main().catch((e) => {
  console.error("✗ ingest failed:", e.message);
  process.exit(1);
});
