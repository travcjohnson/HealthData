/**
 * Pure normalization helpers used by the ingestion pipeline.
 * No DB / IO here so they can be unit-tested in isolation.
 */

/** Canonical form of a label for alias matching (mirrors metric_aliases.alias_norm). */
export function normalizeLabel(label: string): string {
  return (label || "").toLowerCase().replace(/[^a-z0-9]/g, "");
}

/** Extract the first numeric value from a messy string ("5.4 %", "1,234", "<0.5"). */
export function parseNumeric(raw: string | number | null | undefined): number | null {
  if (raw == null) return null;
  if (typeof raw === "number") return Number.isFinite(raw) ? raw : null;
  const m = String(raw).replace(/,/g, "").match(/-?\d+(\.\d+)?/);
  return m ? parseFloat(m[0]) : null;
}

export type RefRange = { low: number | null; high: number | null; text: string | null };

/** Parse a reference-range string: "<5.7", ">40", "70-99", "0.4 - 4.0", "Negative". */
export function parseRefRange(raw: string | null | undefined): RefRange {
  if (!raw) return { low: null, high: null, text: null };
  const s = String(raw).trim();
  const text = s;
  const lt = s.match(/^[<≤]\s*(-?\d+(\.\d+)?)/);
  if (lt) return { low: null, high: parseFloat(lt[1]), text };
  const gt = s.match(/^[>≥]\s*(-?\d+(\.\d+)?)/);
  if (gt) return { low: parseFloat(gt[1]), high: null, text };
  const range = s.match(/(-?\d+(\.\d+)?)\s*[-–to]+\s*(-?\d+(\.\d+)?)/i);
  if (range) return { low: parseFloat(range[1]), high: parseFloat(range[3]), text };
  return { low: null, high: null, text };
}

export type Conversion = { from_unit: string; to_unit: string; factor: number; offset?: number };

/**
 * Convert `value` from `fromUnit` to `canonicalUnit`. Tries metric-specific
 * conversions first, then generic ones. Returns null (don't guess) if no path.
 */
export function convertUnit(
  value: number,
  fromUnit: string | null,
  canonicalUnit: string,
  conversions: Conversion[],
): { value: number; unit: string } | null {
  const from = (fromUnit || "").trim();
  // already canonical (or no unit reported — assume canonical for the metric)
  if (!from || eqUnit(from, canonicalUnit)) return { value, unit: canonicalUnit };
  const c = conversions.find(
    (x) => eqUnit(x.from_unit, from) && eqUnit(x.to_unit, canonicalUnit),
  );
  if (!c) return null;
  return { value: value * c.factor + (c.offset ?? 0), unit: canonicalUnit };
}

function eqUnit(a: string, b: string): boolean {
  return a.toLowerCase().replace(/\s+/g, "") === b.toLowerCase().replace(/\s+/g, "");
}

export type Catalog = {
  /** alias_norm -> metric_id */
  aliases: Map<string, string>;
  /** metric_id -> { canonical_unit } */
  metrics: Map<string, { canonical_unit: string }>;
  /** metric_id -> conversions (plus generic under key '*') */
  conversions: Map<string, Conversion[]>;
};

export type ExtractedItem = {
  label: string;
  value: string | number | null;
  unit?: string | null;
  refText?: string | null;
  flag?: string | null;
  effectiveAt?: string | Date | null;
  confidence?: number | null;
};

export type NormalizedItem = {
  metricId: string | null;
  resolveStatus: "matched" | "unmatched";
  valueNum: number | null;
  valueUnit: string | null;
  rawValueNum: number | null;
  rawUnit: string | null;
  rawLabel: string;
  refLow: number | null;
  refHigh: number | null;
  refText: string | null;
  confidence: number | null;
  effectiveAt: string | Date | null;
};

/** Resolve one extracted item against the catalog into a canonical observation. */
export function normalizeItem(item: ExtractedItem, cat: Catalog): NormalizedItem {
  const rawLabel = item.label ?? "";
  const metricId = cat.aliases.get(normalizeLabel(rawLabel)) ?? null;
  const rawValueNum = parseNumeric(item.value);
  const ref = parseRefRange(item.refText);
  const base: NormalizedItem = {
    metricId,
    resolveStatus: metricId ? "matched" : "unmatched",
    valueNum: null,
    valueUnit: null,
    rawValueNum,
    rawUnit: item.unit ?? null,
    rawLabel,
    refLow: ref.low,
    refHigh: ref.high,
    refText: ref.text,
    confidence: item.confidence ?? null,
    effectiveAt: item.effectiveAt ?? null,
  };
  if (!metricId || rawValueNum == null) return base;

  const meta = cat.metrics.get(metricId);
  if (!meta) return base;
  const convs = [
    ...(cat.conversions.get(metricId) ?? []),
    ...(cat.conversions.get("*") ?? []),
  ];
  const conv = convertUnit(rawValueNum, item.unit ?? null, meta.canonical_unit, convs);
  if (!conv) return base; // no conversion path → leave normalized value null (review)
  return { ...base, valueNum: round(conv.value, 4), valueUnit: conv.unit };
}

function round(n: number, d: number): number {
  const f = Math.pow(10, d);
  return Math.round(n * f) / f;
}
