/**
 * Generic wearable CSV parser (Oura / Whoop / Garmin daily exports).
 * Detects a date column + maps known header names to catalog aliases.
 * Unknown columns are skipped (surfaced as 'unmatched' only if you map them).
 */
import { parse } from "csv-parse/sync";
import type { ExtractedItem } from "@/lib/ingest/normalize";

// header (lowercased, spaces→_) → { alias, unit }
const COLUMN_MAP: Record<string, { label: string; unit?: string }> = {
  // Oura
  total_sleep_duration: { label: "SleepAnalysis", unit: "hr" }, // seconds handled below
  sleep_efficiency: { label: "Sleep Efficiency", unit: "%" },
  average_hrv: { label: "HRV", unit: "ms" },
  lowest_resting_heart_rate: { label: "RestingHeartRate", unit: "bpm" },
  resting_heart_rate: { label: "RestingHeartRate", unit: "bpm" },
  readiness_score: { label: "Readiness", unit: "score" },
  steps: { label: "StepCount", unit: "count" },
  active_calories: { label: "ActiveEnergyBurned", unit: "kcal" },
  // Whoop
  "resting heart rate (bpm)": { label: "RestingHeartRate", unit: "bpm" },
  "heart rate variability (ms)": { label: "HRV", unit: "ms" },
  "asleep duration (min)": { label: "SleepAnalysis", unit: "min" },
  "sleep performance %": { label: "Sleep Efficiency", unit: "%" },
  // Garmin / generic
  vo2max: { label: "VO2Max", unit: "mL/kg/min" },
  weight: { label: "BodyMass", unit: "kg" },
  respiratory_rate: { label: "RespiratoryRate", unit: "breaths/min" },
};

const DATE_KEYS = ["date", "day", "summary_date", "cycle start time", "start", "timestamp"];

export function parseWearableCsv(csv: string): ExtractedItem[] {
  const rows: Record<string, string>[] = parse(csv, {
    columns: (hdr: string[]) => hdr.map((h) => h.trim()),
    skip_empty_lines: true,
    relax_column_count: true,
  });
  const items: ExtractedItem[] = [];

  for (const row of rows) {
    const lower: Record<string, string> = {};
    for (const [k, v] of Object.entries(row)) lower[k.toLowerCase().trim()] = v;
    const dateKey = DATE_KEYS.find((d) => lower[d]);
    const day = dateKey ? String(lower[dateKey]).slice(0, 10) : null;
    if (!day) continue;

    for (const [col, def] of Object.entries(COLUMN_MAP)) {
      if (!(col in lower)) continue;
      const raw = lower[col];
      if (raw === "" || raw == null) continue;
      let value = parseFloat(String(raw).replace(/,/g, ""));
      if (!Number.isFinite(value)) continue;
      let unit = def.unit;
      // normalize sleep duration to hours
      if (def.label === "SleepAnalysis") {
        if (col.includes("min")) value = value / 60;
        else if (col === "total_sleep_duration" && value > 1000) value = value / 3600; // seconds
        unit = "hr";
      }
      items.push({
        label: def.label,
        value: Math.round(value * 100) / 100,
        unit,
        effectiveAt: `${day}T12:00:00Z`,
        confidence: 1,
      });
    }
  }
  return items;
}
