-- ════════════════════════════════════════════════════════════════════════════
-- 0002 · Metric catalog — the single source of truth that collapses the many
--        lab/wearable synonyms into canonical, chartable, color-coded metrics.
-- ════════════════════════════════════════════════════════════════════════════

-- A canonical metric (e.g. 'hba1c'). Merges the data-architecture "metric"
-- with the clinical "metric definition" (direction, optimal band, flag bands).
create table if not exists metrics (
  id              text primary key,             -- slug: 'hba1c','ldl_c','vo2max'
  display_name    text not null,                -- 'Hemoglobin A1c'
  panel           text not null,                -- 'lipids','metabolic','cbc','vital','sleep'...
  canonical_unit  text not null,                -- UCUM-ish: '%','mg/dL','mmol/L','bpm'
  value_kind      value_kind not null default 'quantity',
  loinc_code      text,
  molar_mass      numeric,                      -- g/mol, for mass<->molar conversion

  direction       metric_direction not null default 'target_band',

  -- standard adult reference range (fallback when a report prints none)
  ref_low         numeric,
  ref_high        numeric,
  -- tighter "optimal"/longevity band (nullable)
  optimal_low     numeric,
  optimal_high    numeric,

  -- four flag thresholds (any may be null = unbounded on that side):
  --   value <= flag_low_red    OR value >= flag_high_red    -> red    (high risk)
  --   value <= flag_low_yellow OR value >= flag_high_yellow -> yellow (borderline)
  --   otherwise                                             -> green
  flag_low_red     numeric,
  flag_low_yellow  numeric,
  flag_high_yellow numeric,
  flag_high_red    numeric,

  sex_specific    boolean not null default false,
  fasting_required boolean not null default false,
  cadence_days    int,                          -- recommended retest cadence (null=continuous/none)
  hero            boolean not null default false,-- surface in the longevity scorecard
  unit_aliases    text[] not null default '{}',
  notes           text
);

-- every label any source has ever printed -> one metric
create table if not exists metric_aliases (
  id          bigserial primary key,
  metric_id   text not null references metrics(id) on delete cascade,
  alias       text not null,
  alias_norm  text generated always as
                 (lower(regexp_replace(alias, '[^a-zA-Z0-9]', '', 'g'))) stored,
  source_type source_type,
  loinc_code  text,
  unique (alias_norm, metric_id)
);
create index if not exists metric_aliases_norm_idx on metric_aliases (alias_norm);

-- deterministic, auditable unit conversions (canonicalization)
create table if not exists unit_conversions (
  id        bigserial primary key,
  metric_id text references metrics(id) on delete cascade, -- null = generic
  from_unit text not null,
  to_unit   text not null,
  factor    numeric not null,
  "offset"  numeric not null default 0,
  unique (metric_id, from_unit, to_unit)
);

-- sex/age-specific canonical bands, resolved against the user's profile
create table if not exists reference_ranges (
  id        bigserial primary key,
  metric_id text not null references metrics(id) on delete cascade,
  sex       text,                 -- 'male','female', null=any
  age_min   int,
  age_max   int,
  unit      text not null,
  low       numeric,
  high      numeric,
  source    text
);
create index if not exists reference_ranges_metric_idx on reference_ranges (metric_id);
