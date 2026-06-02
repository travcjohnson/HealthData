/**
 * Seed realistic demo health data for one owner.
 *
 *   tsx scripts/db/seed.ts --owner <uuid> [--create-user]
 *
 * LOCAL: pass --create-user to also create the auth.users row (shim).
 * SUPABASE: sign up first, then run with --owner <your auth user id> (NO
 *           --create-user — the row already exists; do not write auth.users).
 *
 * The generated story: someone actively improving cardiometabolic health —
 * ApoB/LDL/A1c/CRP trending down (wins), Lp(a) genetically high (red),
 * Vitamin D recovering from deficiency, resting HR creeping up (trending wrong).
 */
import { Client } from "pg";

const OWNER = argVal("--owner") ?? "00000000-0000-0000-0000-000000000001";
const CREATE_USER = process.argv.includes("--create-user");

// deterministic PRNG so seeds are reproducible
let _s = 12345;
function rnd() { _s = (_s * 1103515245 + 12345) & 0x7fffffff; return _s / 0x7fffffff; }
function noise(amp: number) { return (rnd() - 0.5) * 2 * amp; }

type Row = { metric: string; at: Date; value: number; src: string };

/** Quarterly lab series interpolating start→end over `quarters`, with noise. */
function labSeries(metric: string, start: number, end: number, quarters: number, amp: number): Row[] {
  const rows: Row[] = [];
  const now = new Date();
  for (let i = 0; i < quarters; i++) {
    const frac = quarters === 1 ? 1 : i / (quarters - 1);
    const at = new Date(now.getTime() - (quarters - 1 - i) * 91 * 86400000);
    const value = round(start + (end - start) * frac + noise(amp), 2);
    rows.push({ metric, at, value, src: "lab_pdf" });
  }
  return rows;
}

/** Daily wearable series start→end over `days`, with noise. */
function dailySeries(metric: string, start: number, end: number, days: number, amp: number, src = "oura"): Row[] {
  const rows: Row[] = [];
  const now = new Date();
  for (let i = 0; i < days; i++) {
    const frac = i / (days - 1);
    const at = new Date(now.getTime() - (days - 1 - i) * 86400000);
    const value = round(start + (end - start) * frac + noise(amp), 1);
    rows.push({ metric, at, value, src });
  }
  return rows;
}

function build(): Row[] {
  const rows: Row[] = [];
  // ── Labs (12 quarterly draws ≈ 3 years) ──────────────────────────────────
  rows.push(...labSeries("apob", 105, 82, 8, 3));            // improving, still borderline
  rows.push(...labSeries("ldl_c", 140, 95, 8, 5));
  rows.push(...labSeries("hdl_c", 48, 55, 8, 2));
  rows.push(...labSeries("triglycerides", 165, 95, 8, 10));
  rows.push(...labSeries("total_cholesterol", 212, 176, 8, 6));
  rows.push(...labSeries("hba1c", 5.9, 5.4, 8, 0.05));       // win
  rows.push(...labSeries("glucose_fasting", 103, 91, 8, 3));
  rows.push(...labSeries("hs_crp", 2.6, 0.8, 8, 0.2));       // win
  rows.push(...labSeries("vitamin_d", 23, 39, 6, 2));        // recovering from deficient
  rows.push(...labSeries("tsh", 2.1, 1.9, 6, 0.2));
  rows.push(...labSeries("alt", 31, 22, 6, 2));
  rows.push(...labSeries("egfr", 96, 94, 6, 2));
  rows.push(...labSeries("creatinine", 0.98, 1.0, 6, 0.05));
  rows.push(...labSeries("ferritin", 85, 95, 6, 8));
  rows.push(...labSeries("testosterone_total", 470, 615, 6, 25));
  rows.push(...labSeries("psa", 0.9, 1.0, 6, 0.1));
  rows.push(...labSeries("lpa", 132, 134, 1, 0));            // genetically high (red), one reading
  // ── Wearables / vitals ───────────────────────────────────────────────────
  rows.push(...dailySeries("resting_hr_wearable", 56, 63, 120, 2));   // creeping UP (trending wrong)
  rows.push(...dailySeries("hrv_rmssd", 68, 58, 120, 6));             // drifting down
  rows.push(...dailySeries("sleep_duration", 7.3, 7.0, 120, 0.6));
  rows.push(...dailySeries("sleep_efficiency", 88, 87, 120, 3));
  rows.push(...dailySeries("steps", 8600, 8800, 120, 1800));
  rows.push(...dailySeries("weight", 84, 80, 120, 0.4, "apple_health"));
  rows.push(...dailySeries("vo2max", 44, 47, 120, 0.5, "apple_health").filter((_, i) => i % 7 === 0)); // weekly
  // BP — a dozen manual readings
  rows.push(...dailySeries("bp_systolic", 122, 118, 12, 4, "manual"));
  rows.push(...dailySeries("bp_diastolic", 80, 76, 12, 3, "manual"));
  return rows;
}

function round(n: number, d: number) { const f = 10 ** d; return Math.round(n * f) / f; }
function argVal(flag: string) { const i = process.argv.indexOf(flag); return i >= 0 ? process.argv[i + 1] : undefined; }

async function main() {
  const db = new Client({ connectionString: process.env.DATABASE_URL });
  await db.connect();
  try {
    if (CREATE_USER) {
      await db.query(
        "insert into auth.users (id, email) values ($1,$2) on conflict (id) do nothing",
        [OWNER, "demo@local"],
      );
    }
    await db.query(
      `insert into profile (owner_id, display_name, birth_date, sex_at_birth, height_cm)
       values ($1,'Demo User','1985-05-20','male',180)
       on conflict (owner_id) do update set
         display_name=excluded.display_name, birth_date=excluded.birth_date,
         sex_at_birth=excluded.sex_at_birth, height_cm=excluded.height_cm`,
      [OWNER],
    );
    // idempotent: clear prior seeded observations for this owner
    await db.query("delete from observations where owner_id=$1", [OWNER]);

    const rows = build();
    const unitByMetric = new Map<string, string>(
      (await db.query("select id, canonical_unit from metrics")).rows.map(
        (r: any) => [r.id, r.canonical_unit],
      ),
    );
    let n = 0;
    for (const r of rows) {
      const unit = unitByMetric.get(r.metric) ?? null;
      const res = await db.query(
        `insert into observations
           (owner_id, metric_id, effective_at, value_num, value_unit, raw_value_num, raw_unit,
            raw_label, source_type, confidence, status)
         values ($1,$2,$3,$4,$5,$4,$5,$2,$6,1.0,'normalized')
         on conflict (owner_id, metric_id, effective_at, value_num) where status='normalized'
         do nothing`,
        [OWNER, r.metric, r.at.toISOString(), r.value, unit, r.src],
      );
      n += res.rowCount ?? 0;
    }
    const summary = await db.query(
      "select count(*)::int obs, count(distinct metric_id)::int metrics from observations where owner_id=$1",
      [OWNER],
    );
    console.log(`✓ seeded ${n} observations for ${OWNER}`);
    console.log(`  ${summary.rows[0].obs} total observations across ${summary.rows[0].metrics} metrics`);
  } finally {
    await db.end();
  }
}

main().catch((e) => { console.error("✗ seed failed:", e.message); process.exit(1); });
