/**
 * Apple Health `export.xml` parser → daily-aggregated ExtractedItems.
 *
 * Apple records thousands of tiny samples; for trends we aggregate to one
 * value per day (SUM for cumulative metrics, MEAN for the rest). For very large
 * exports prefer a streaming SAX parser; this loads the document for simplicity.
 */
import { XMLParser } from "fast-xml-parser";
import type { ExtractedItem } from "@/lib/ingest/normalize";

// HK type → { alias label the catalog knows, how to aggregate per day }
const HK_MAP: Record<string, { label: string; agg: "sum" | "mean" | "last" }> = {
  HKQuantityTypeIdentifierBodyMass: { label: "BodyMass", agg: "last" },
  HKQuantityTypeIdentifierBodyMassIndex: { label: "Body Mass Index", agg: "last" },
  HKQuantityTypeIdentifierBodyFatPercentage: { label: "BodyFatPercentage", agg: "last" },
  HKQuantityTypeIdentifierHeartRateVariabilitySDNN: { label: "HeartRateVariabilitySDNN", agg: "mean" },
  HKQuantityTypeIdentifierRestingHeartRate: { label: "RestingHeartRate", agg: "mean" },
  HKQuantityTypeIdentifierStepCount: { label: "StepCount", agg: "sum" },
  HKQuantityTypeIdentifierActiveEnergyBurned: { label: "ActiveEnergyBurned", agg: "sum" },
  HKQuantityTypeIdentifierRespiratoryRate: { label: "RespiratoryRate", agg: "mean" },
  HKQuantityTypeIdentifierVO2Max: { label: "VO2Max", agg: "last" },
  HKQuantityTypeIdentifierBloodPressureSystolic: { label: "BloodPressureSystolic", agg: "mean" },
  HKQuantityTypeIdentifierBloodPressureDiastolic: { label: "BloodPressureDiastolic", agg: "mean" },
};

const PCT_BODYFAT_THRESHOLD = 1; // Apple stores body fat as fraction (0.18) → ×100

type Bucket = { label: string; unit: string; sum: number; n: number; last: number; agg: string };

export function parseAppleHealth(xml: string): ExtractedItem[] {
  const parser = new XMLParser({ ignoreAttributes: false, attributeNamePrefix: "@_" });
  const doc = parser.parse(xml);
  const records = asArray(doc?.HealthData?.Record);
  const sleepByDay = new Map<string, number>(); // hours asleep per day
  const buckets = new Map<string, Bucket>(); // key: `${label}|${day}`

  for (const r of records) {
    const type = r["@_type"];
    const day = (r["@_startDate"] || "").slice(0, 10);
    if (!day) continue;

    // Sleep is a category record; accumulate "asleep" duration in hours.
    if (type === "HKCategoryTypeIdentifierSleepAnalysis") {
      const val = String(r["@_value"] || "");
      if (/Asleep/i.test(val)) {
        const start = new Date(r["@_startDate"]);
        const end = new Date(r["@_endDate"]);
        const hrs = (end.getTime() - start.getTime()) / 3_600_000;
        if (hrs > 0 && hrs < 24) sleepByDay.set(day, (sleepByDay.get(day) ?? 0) + hrs);
      }
      continue;
    }

    const map = HK_MAP[type];
    if (!map) continue;
    let value = parseFloat(r["@_value"]);
    if (!Number.isFinite(value)) continue;
    let unit = r["@_unit"] || "";
    if (map.label === "BodyFatPercentage" && value < PCT_BODYFAT_THRESHOLD) {
      value *= 100;
      unit = "%";
    }
    const key = `${map.label}|${day}`;
    const b = buckets.get(key) ?? { label: map.label, unit, sum: 0, n: 0, last: value, agg: map.agg };
    b.sum += value;
    b.n += 1;
    b.last = value;
    b.unit = unit;
    buckets.set(key, b);
  }

  const items: ExtractedItem[] = [];
  for (const [key, b] of buckets) {
    const day = key.split("|")[1];
    const value =
      b.agg === "sum" ? b.sum : b.agg === "mean" ? b.sum / b.n : b.last;
    items.push({
      label: b.label,
      value: Math.round(value * 100) / 100,
      unit: b.unit,
      effectiveAt: `${day}T12:00:00Z`,
      confidence: 1,
    });
  }
  for (const [day, hrs] of sleepByDay) {
    items.push({
      label: "SleepAnalysis",
      value: Math.round(hrs * 100) / 100,
      unit: "hr",
      effectiveAt: `${day}T12:00:00Z`,
      confidence: 1,
    });
  }
  return items;
}

function asArray<T>(x: T | T[] | undefined): T[] {
  return x == null ? [] : Array.isArray(x) ? x : [x];
}
