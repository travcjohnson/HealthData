import { createClient } from "@/lib/supabase/server";
import type {
  MetricSummary, Observation, SourceDocument, Profile,
} from "@/lib/types";

/** All per-metric summaries for the logged-in user (RLS-scoped). */
export async function getSummaries(): Promise<MetricSummary[]> {
  const supabase = await createClient();
  const { data, error } = await supabase
    .from("v_metric_summary")
    .select("*")
    .order("panel");
  if (error) throw error;
  return (data ?? []) as MetricSummary[];
}

/** Full time-series for one metric (oldest → newest). */
export async function getMetricSeries(metricId: string): Promise<Observation[]> {
  const supabase = await createClient();
  const { data, error } = await supabase
    .from("v_observations")
    .select(
      "id, metric_id, display_name, panel, effective_at, value_num, value_unit, raw_label, raw_unit, source_type, confidence, eff_ref_low, eff_ref_high, optimal_low, optimal_high, flag, direction",
    )
    .eq("metric_id", metricId)
    .order("effective_at", { ascending: true });
  if (error) throw error;
  return (data ?? []) as Observation[];
}

/** metric_id → ordered value array, for inline sparklines (single query). */
export async function getSparkValues(): Promise<Map<string, number[]>> {
  const supabase = await createClient();
  const { data, error } = await supabase
    .from("observations")
    .select("metric_id, value_num, effective_at")
    .eq("status", "normalized")
    .not("value_num", "is", null)
    .order("effective_at", { ascending: true })
    .limit(20000);
  if (error) throw error;
  const map = new Map<string, number[]>();
  for (const r of (data ?? []) as { metric_id: string; value_num: number }[]) {
    const arr = map.get(r.metric_id) ?? [];
    arr.push(Number(r.value_num));
    map.set(r.metric_id, arr);
  }
  return map;
}

export async function getProfile(): Promise<Profile | null> {
  const supabase = await createClient();
  const { data } = await supabase.from("profile").select("*").maybeSingle();
  return (data as Profile) ?? null;
}

export async function getSources(): Promise<SourceDocument[]> {
  const supabase = await createClient();
  const { data, error } = await supabase
    .from("source_documents")
    .select("id, source_type, original_name, source_lab, document_date, status, created_at")
    .order("created_at", { ascending: false });
  if (error) throw error;
  return (data ?? []) as SourceDocument[];
}

export async function getMedications() {
  const supabase = await createClient();
  const { data } = await supabase
    .from("medications")
    .select("id, name, dose_value, dose_unit, frequency, status, started_on")
    .order("status")
    .order("name");
  return data ?? [];
}

export async function getConditions() {
  const supabase = await createClient();
  const { data } = await supabase
    .from("conditions")
    .select("id, name, status, onset_date, icd10_code")
    .order("status")
    .order("name");
  return data ?? [];
}

export type Triage = {
  outOfRange: MetricSummary[]; // red
  borderline: MetricSummary[]; // yellow
  trendingWrong: MetricSummary[]; // green/yellow but moving toward risk
  wins: MetricSummary[]; // improving meaningfully
  overdue: MetricSummary[];
  heroes: MetricSummary[];
};

/** Build the command-center triage stacks from the summaries. */
export function buildTriage(rows: MetricSummary[]): Triage {
  const worseningPct = (m: MetricSummary): boolean => {
    if (m.pct_change == null || Math.abs(m.pct_change) < 3) return false;
    if (m.direction === "target_band") return false;
    const rising = m.pct_change > 0;
    return m.direction === "higher_better" ? !rising : rising;
  };
  const improvingPct = (m: MetricSummary): boolean => {
    if (m.pct_change == null || Math.abs(m.pct_change) < 3) return false;
    if (m.direction === "target_band") return false;
    const rising = m.pct_change > 0;
    return m.direction === "higher_better" ? rising : !rising;
  };

  return {
    outOfRange: rows.filter((m) => m.flag === "red"),
    borderline: rows.filter((m) => m.flag === "yellow"),
    trendingWrong: rows.filter((m) => m.flag !== "red" && m.flag !== "unknown" && worseningPct(m)),
    wins: rows.filter((m) => improvingPct(m)),
    overdue: rows.filter((m) => (m.days_overdue ?? 0) > 0),
    heroes: rows.filter((m) => m.hero),
  };
}
