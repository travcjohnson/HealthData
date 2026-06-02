/**
 * Lightweight assertion tests for parsers + pure health logic (no DB, no deps).
 *   tsx scripts/ingest/parsers.test.ts
 */
import {
  normalizeLabel, parseNumeric, parseRefRange, convertUnit, normalizeItem,
  type Catalog,
} from "@/lib/ingest/normalize";
import { bmi, homaIr, egfrCkdEpi2021, friedewaldLdl, nonHdl } from "@/lib/health/derived";
import { slopePerDay, classifyTrend, pctChange } from "@/lib/health/trends";
import { parseAppleHealth } from "./parsers/appleHealth";
import { parseWearableCsv } from "./parsers/wearableCsv";
import { parseLabText } from "./parsers/labText";

let passed = 0, failed = 0;
function ok(name: string, cond: boolean, extra?: unknown) {
  if (cond) { passed++; }
  else { failed++; console.error(`✗ ${name}`, extra ?? ""); }
}
function near(a: number | null, b: number, eps = 0.01) {
  return a != null && Math.abs(a - b) <= eps;
}

// ── normalize helpers ───────────────────────────────────────────────────────
ok("normalizeLabel strips", normalizeLabel("Hemoglobin A1c") === "hemoglobina1c");
ok("parseNumeric messy", parseNumeric("<0.5") === 0.5 && parseNumeric("1,234") === 1234);
const r1 = parseRefRange("70-99"); ok("ref range", r1.low === 70 && r1.high === 99);
const r2 = parseRefRange("<5.7"); ok("ref lt", r2.low === null && r2.high === 5.7);
const conv = convertUnit(5.5, "mmol/L", "mg/dL", [{ from_unit: "mmol/L", to_unit: "mg/dL", factor: 18.0156 }]);
ok("glucose mmol→mg", near(conv?.value ?? null, 99.09, 0.1));
ok("no conversion path → null", convertUnit(5, "foo", "mg/dL", []) === null);

const cat: Catalog = {
  aliases: new Map([["glucose", "glucose_fasting"], ["a1c", "hba1c"]]),
  metrics: new Map([
    ["glucose_fasting", { canonical_unit: "mg/dL" }],
    ["hba1c", { canonical_unit: "%" }],
  ]),
  conversions: new Map([["glucose_fasting", [{ from_unit: "mmol/L", to_unit: "mg/dL", factor: 18.0156 }]]]),
};
const ni = normalizeItem({ label: "Glucose", value: "5.5", unit: "mmol/L", effectiveAt: "2024-01-01" }, cat);
ok("normalizeItem converts", ni.metricId === "glucose_fasting" && near(ni.valueNum, 99.09, 0.1));
const niU = normalizeItem({ label: "Mystery Test", value: "1" }, cat);
ok("normalizeItem unmatched", niU.metricId === null && niU.resolveStatus === "unmatched");

// ── derived metrics ─────────────────────────────────────────────────────────
ok("bmi", near(bmi(80, 180), 24.7, 0.1));
ok("homaIr", near(homaIr(90, 6), 1.33, 0.01));
ok("friedewald invalid >=400", friedewaldLdl(250, 50, 400) === null);
ok("nonHdl", nonHdl(200, 50) === 150);
// CKD-EPI 2021: male, Scr 0.9, age 40 → ~109; female same → ~104
ok("egfr male", near(egfrCkdEpi2021(0.9, 40, "male"), 109, 2));
ok("egfr female < male", (egfrCkdEpi2021(0.9, 40, "female") ?? 0) < (egfrCkdEpi2021(0.9, 40, "male") ?? 0));

// ── trends ──────────────────────────────────────────────────────────────────
const series = [
  { t: "2024-01-01", v: 100 }, { t: "2024-02-01", v: 110 }, { t: "2024-03-01", v: 120 },
];
ok("slope positive", (slopePerDay(series) ?? 0) > 0);
ok("lower_better rising = worsening", classifyTrend(series, "lower_better") === "worsening");
ok("higher_better rising = improving", classifyTrend(series, "higher_better") === "improving");
ok("pctChange", pctChange(110, 100) === 10);

// ── Apple Health XML ────────────────────────────────────────────────────────
const appleXml = `<?xml version="1.0"?><HealthData>
<Record type="HKQuantityTypeIdentifierStepCount" unit="count" value="2000" startDate="2024-03-01 08:00:00 -0000" endDate="2024-03-01 09:00:00 -0000"/>
<Record type="HKQuantityTypeIdentifierStepCount" unit="count" value="3000" startDate="2024-03-01 10:00:00 -0000" endDate="2024-03-01 11:00:00 -0000"/>
<Record type="HKQuantityTypeIdentifierRestingHeartRate" unit="count/min" value="58" startDate="2024-03-01 07:00:00 -0000" endDate="2024-03-01 07:01:00 -0000"/>
<Record type="HKQuantityTypeIdentifierBodyFatPercentage" unit="%" value="0.18" startDate="2024-03-01 07:00:00 -0000" endDate="2024-03-01 07:01:00 -0000"/>
<Record type="HKCategoryTypeIdentifierSleepAnalysis" value="HKCategoryValueSleepAnalysisAsleepCore" startDate="2024-03-01 23:00:00 -0000" endDate="2024-03-02 05:00:00 -0000"/>
</HealthData>`;
const apple = parseAppleHealth(appleXml);
const steps = apple.find((i) => i.label === "StepCount");
ok("apple steps summed", steps?.value === 5000, steps);
ok("apple resting hr mean", apple.find((i) => i.label === "RestingHeartRate")?.value === 58);
ok("apple bodyfat ×100", apple.find((i) => i.label === "BodyFatPercentage")?.value === 18);
ok("apple sleep hours", near(Number(apple.find((i) => i.label === "SleepAnalysis")?.value), 6, 0.1));

// ── Wearable CSV (Oura-like) ────────────────────────────────────────────────
const csv = `date,average_hrv,lowest_resting_heart_rate,total_sleep_duration,sleep_efficiency,steps
2024-03-01,65,52,28800,91,8200
2024-03-02,48,55,25200,84,5400`;
const wear = parseWearableCsv(csv);
ok("csv hrv", wear.find((i) => i.label === "HRV" && i.effectiveAt === "2024-03-01T12:00:00Z")?.value === 65);
ok("csv sleep seconds→hours", near(Number(wear.find((i) => i.label === "SleepAnalysis")?.value), 8, 0.1));
ok("csv steps", wear.find((i) => i.label === "StepCount" && String(i.effectiveAt).startsWith("2024-03-02"))?.value === 5400);

// ── Lab text ────────────────────────────────────────────────────────────────
const labTxt = `Collected: 03/14/2024
Glucose 95 mg/dL 70-99
Hemoglobin A1c 5.4 % <5.7
LDL Cholesterol 88 mg/dL H 0-99
Patient Name: John Doe`;
const lab = parseLabText(labTxt);
ok("labText glucose", lab.some((i) => i.label === "Glucose" && i.value === "95"));
ok("labText a1c ref", lab.some((i) => i.label.includes("A1c") && i.refText === "<5.7"));
ok("labText skips patient line", !lab.some((i) => /patient name/i.test(i.label)));

console.log(`\n${passed} passed, ${failed} failed`);
process.exit(failed ? 1 : 0);
