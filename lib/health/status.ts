/** Maps a flag to display label + Tailwind classes. UI-side, presentational. */
export type Flag = "red" | "yellow" | "green" | "unknown";

export const FLAG_META: Record<
  Flag,
  { label: string; dot: string; text: string; bg: string; ring: string }
> = {
  red: { label: "High risk", dot: "bg-status-high", text: "text-status-high", bg: "bg-red-50", ring: "ring-red-200" },
  yellow: { label: "Borderline", dot: "bg-status-borderline", text: "text-status-borderline", bg: "bg-amber-50", ring: "ring-amber-200" },
  green: { label: "In range", dot: "bg-status-normal", text: "text-status-normal", bg: "bg-emerald-50", ring: "ring-emerald-200" },
  unknown: { label: "No range", dot: "bg-status-unknown", text: "text-slate-400", bg: "bg-slate-50", ring: "ring-slate-200" },
};

export const PANEL_LABELS: Record<string, string> = {
  lipids: "Cardiometabolic · Lipids",
  metabolic: "Metabolic · Glycemic",
  inflammation: "Inflammation",
  body: "Body Composition",
  vital: "Blood Pressure & HR",
  kidney: "Kidney",
  liver: "Liver",
  thyroid: "Thyroid",
  cbc: "Complete Blood Count",
  vitamins: "Vitamins & Minerals",
  hormones: "Hormones",
  wearable: "Wearable",
  sleep: "Sleep",
  activity: "Activity",
};

export const PANEL_ORDER = [
  "lipids", "metabolic", "inflammation", "vital", "body",
  "kidney", "liver", "thyroid", "cbc", "vitamins", "hormones",
  "wearable", "sleep", "activity",
];

/** Format a numeric value compactly with its unit. */
export function fmtValue(v: number | null | undefined, unit?: string | null): string {
  if (v == null) return "—";
  const n = Math.abs(v) >= 1000 ? v.toLocaleString() : String(Math.round(v * 100) / 100);
  return unit ? `${n} ${unit}` : n;
}

/** Arrow + class for a percent change given metric direction. */
export function trendArrow(
  pctChange: number | null | undefined,
  direction: "higher_better" | "lower_better" | "target_band",
): { arrow: string; good: boolean | null } {
  if (pctChange == null || Math.abs(pctChange) < 1) return { arrow: "→", good: null };
  const rising = pctChange > 0;
  if (direction === "target_band") return { arrow: rising ? "↑" : "↓", good: null };
  const good = direction === "higher_better" ? rising : !rising;
  return { arrow: rising ? "↑" : "↓", good };
}
