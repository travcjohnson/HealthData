import type { Flag } from "@/lib/health/status";

export type Direction = "higher_better" | "lower_better" | "target_band";

export type MetricSummary = {
  metric_id: string;
  display_name: string;
  panel: string;
  canonical_unit: string;
  direction: Direction;
  hero: boolean;
  cadence_days: number | null;
  latest_value: number | null;
  latest_at: string | null;
  prev_value: number | null;
  prev_at: string | null;
  delta: number | null;
  pct_change: number | null;
  n_readings: number;
  eff_ref_low: number | null;
  eff_ref_high: number | null;
  optimal_low: number | null;
  optimal_high: number | null;
  flag: Flag;
  is_optimal: boolean;
  days_overdue: number | null;
};

export type Observation = {
  id: string;
  metric_id: string;
  display_name: string;
  panel: string;
  effective_at: string;
  value_num: number | null;
  value_unit: string | null;
  raw_label: string | null;
  raw_unit: string | null;
  source_type: string;
  confidence: number | null;
  eff_ref_low: number | null;
  eff_ref_high: number | null;
  optimal_low: number | null;
  optimal_high: number | null;
  flag: Flag;
  direction: Direction;
};

export type SourceDocument = {
  id: string;
  source_type: string;
  original_name: string | null;
  source_lab: string | null;
  document_date: string | null;
  status: string;
  created_at: string;
};

export type Profile = {
  owner_id: string;
  display_name: string | null;
  birth_date: string | null;
  sex_at_birth: string | null;
  height_cm: number | null;
};
