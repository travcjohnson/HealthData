/**
 * Apply SQL migrations in order against DATABASE_URL.
 *
 *   tsx scripts/db/migrate.ts            # apply supabase/migrations/*.sql
 *   tsx scripts/db/migrate.ts --local    # also apply the local auth shim first
 *
 * Re-runnable: migrations are written to be idempotent (IF NOT EXISTS / ON
 * CONFLICT / guarded DO blocks), so this doubles as a "sync schema" command.
 */
import { Client } from "pg";
import { readFileSync, readdirSync } from "node:fs";
import { join } from "node:path";

const MIGRATIONS_DIR = join(process.cwd(), "supabase", "migrations");
const SHIM = join(process.cwd(), "scripts", "db", "_local_shim.sql");

async function main() {
  const local = process.argv.includes("--local");
  const url = process.env.DATABASE_URL;
  if (!url) throw new Error("DATABASE_URL is not set");

  const client = new Client({ connectionString: url });
  await client.connect();
  console.log(`→ connected: ${url.replace(/:[^:@/]*@/, ":***@")}`);

  try {
    if (local) {
      console.log("→ applying local auth shim");
      await client.query(readFileSync(SHIM, "utf8"));
    }
    const files = readdirSync(MIGRATIONS_DIR)
      .filter((f) => f.endsWith(".sql"))
      .sort();
    for (const f of files) {
      process.stdout.write(`→ ${f} ... `);
      await client.query(readFileSync(join(MIGRATIONS_DIR, f), "utf8"));
      console.log("ok");
    }
    const { rows } = await client.query("select count(*)::int n from metrics");
    console.log(`✓ migrations applied. metric catalog: ${rows[0].n} metrics`);
  } finally {
    await client.end();
  }
}

main().catch((e) => {
  console.error("✗ migration failed:\n", e.message);
  process.exit(1);
});
