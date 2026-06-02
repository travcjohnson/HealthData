/**
 * Trend analysis over a time-series of observations. Pure & unit-tested.
 */
export type Point = { t: number | Date | string; v: number };
export type Direction = "higher_better" | "lower_better" | "target_band";
export type TrendVerdict = "improving" | "worsening" | "stable";

function toMs(t: Point["t"]): number {
  return t instanceof Date ? t.getTime() : new Date(t).getTime();
}

/** Ordinary-least-squares slope in value-units per DAY. null if < 2 points. */
export function slopePerDay(points: Point[]): number | null {
  const pts = [...points]
    .map((p) => ({ x: toMs(p.t), y: p.v }))
    .filter((p) => Number.isFinite(p.x) && Number.isFinite(p.y))
    .sort((a, b) => a.x - b.x);
  if (pts.length < 2) return null;
  const n = pts.length;
  const meanX = pts.reduce((s, p) => s + p.x, 0) / n;
  const meanY = pts.reduce((s, p) => s + p.y, 0) / n;
  let num = 0,
    den = 0;
  for (const p of pts) {
    num += (p.x - meanX) * (p.y - meanY);
    den += (p.x - meanX) ** 2;
  }
  if (den === 0) return null;
  const perMs = num / den;
  return perMs * 86_400_000; // per day
}

/**
 * Classify the trend in clinical terms using the metric's `direction`.
 * `noise` is a magnitude floor (per reading) below which we call it stable.
 */
export function classifyTrend(
  points: Point[],
  direction: Direction,
  noise = 0,
): TrendVerdict {
  if (points.length < 3) return "stable";
  const sorted = [...points].sort((a, b) => toMs(a.t) - toMs(b.t));
  const first = sorted[0].v;
  const last = sorted[sorted.length - 1].v;
  const change = last - first;
  if (Math.abs(change) <= noise) return "stable";
  // target_band: judged elsewhere (depends on which side of the band); treat as stable here.
  if (direction === "target_band") return "stable";
  const rising = change > 0;
  if (direction === "higher_better") return rising ? "improving" : "worsening";
  return rising ? "worsening" : "improving"; // lower_better
}

/** Trailing rolling average; window = number of points. */
export function rollingAverage(points: Point[], window: number): Point[] {
  const sorted = [...points].sort((a, b) => toMs(a.t) - toMs(b.t));
  return sorted.map((p, i) => {
    const start = Math.max(0, i - window + 1);
    const slice = sorted.slice(start, i + 1);
    const avg = slice.reduce((s, q) => s + q.v, 0) / slice.length;
    return { t: p.t, v: Math.round(avg * 1000) / 1000 };
  });
}

/** Percent change of latest vs a baseline value. */
export function pctChange(latest: number, baseline: number): number | null {
  if (baseline === 0 || baseline == null) return null;
  return Math.round(((latest - baseline) / Math.abs(baseline)) * 1000) / 10;
}

/**
 * "Trending the wrong way while still in range" — the command-center
 * differentiator. Returns true when a green/yellow metric is moving toward red.
 */
export function trendingTowardRisk(
  points: Point[],
  direction: Direction,
  currentFlag: "green" | "yellow" | "red" | "unknown",
  noise = 0,
): boolean {
  if (currentFlag === "red" || currentFlag === "unknown") return false;
  const verdict = classifyTrend(points, direction, noise);
  return verdict === "worsening";
}
