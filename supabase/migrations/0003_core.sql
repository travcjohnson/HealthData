-- ════════════════════════════════════════════════════════════════════════════
-- 0003 · Core data: profile, provenance, observations, clinical records.
--        Raw label/value/unit are IMMUTABLE provenance; value_num (canonical)
--        is derived and always re-computable from staging.
-- ════════════════════════════════════════════════════════════════════════════

-- one row per user; drives age/sex-specific reference ranges
create table if not exists profile (
  owner_id     uuid primary key references auth.users(id) on delete cascade,
  display_name text,
  birth_date   date,
  sex_at_birth text,                 -- 'male','female'
  height_cm    numeric,
  blood_type   text,
  baseline_pinned_at date,           -- "compare trends against this date"
  created_at   timestamptz not null default now(),
  updated_at   timestamptz not null default now()
);

-- ── Provenance layer ────────────────────────────────────────────────────────
-- one row per uploaded artifact (PDF, xml, csv, image)
create table if not exists source_documents (
  id             uuid primary key default gen_random_uuid(),
  owner_id       uuid not null references auth.users(id) on delete cascade,
  source_type    source_type not null,
  storage_bucket text not null default 'health-raw',
  storage_path   text not null,
  original_name  text,
  content_sha256 text not null,
  mime_type      text,
  byte_size      bigint,
  source_lab     text,               -- 'Quest','LabCorp','Mercy Hospital'
  document_date  date,
  status         ingest_status not null default 'uploaded',
  ingest_error   text,
  created_at     timestamptz not null default now(),
  unique (owner_id, content_sha256)  -- same file never ingested twice
);

-- one row per extraction run against a document (re-extractable)
create table if not exists extractions (
  id             uuid primary key default gen_random_uuid(),
  owner_id       uuid not null references auth.users(id) on delete cascade,
  document_id    uuid not null references source_documents(id) on delete cascade,
  extractor      text not null,      -- 'claude-sonnet-4-6','apple-xml-parser-v2'
  extractor_kind text not null,      -- 'llm' | 'deterministic'
  prompt_version text,
  raw_output     jsonb,
  page_count     int,
  created_at     timestamptz not null default now()
);

-- groups observations (a lab draw, an office visit, an imaging study)
create table if not exists encounters (
  id             uuid primary key default gen_random_uuid(),
  owner_id       uuid not null references auth.users(id) on delete cascade,
  encounter_date date not null,
  kind           text,              -- 'lab','office_visit','imaging','telehealth'
  provider       text,
  facility       text,
  document_id    uuid references source_documents(id) on delete set null,
  notes          text,
  created_at     timestamptz not null default now()
);

-- ── The universal observation event table (labs, vitals, wearables) ─────────
create table if not exists observations (
  id            uuid primary key default gen_random_uuid(),
  owner_id      uuid not null references auth.users(id) on delete cascade,
  metric_id     text not null references metrics(id),
  effective_at  timestamptz not null,
  value_kind    value_kind not null default 'quantity',

  -- normalized (canonical) value — what charts read
  value_num     numeric,
  value_unit    text,
  value_text    text,
  value_code    text,

  -- original as reported (immutable provenance)
  raw_value_num numeric,
  raw_value_text text,
  raw_unit      text,
  raw_label     text,

  -- reference range AS REPORTED for this specific result
  ref_low       numeric,
  ref_high      numeric,
  ref_text      text,
  abnormal_flag text,                -- 'H','L','HH','A', or null

  -- provenance + trust
  document_id   uuid references source_documents(id) on delete set null,
  extraction_id uuid references extractions(id) on delete set null,
  encounter_id  uuid references encounters(id) on delete set null,
  source_type   source_type not null,
  confidence    numeric check (confidence is null or confidence between 0 and 1),
  status        ingest_status not null default 'normalized',
  supersedes    uuid references observations(id) on delete set null,
  fasting       boolean,             -- context flag, gates some alerts
  context       text,                -- 'post_exercise','acute_illness',...

  meta          jsonb not null default '{}',
  created_at    timestamptz not null default now()
);

-- primary trend access path; also enforces datapoint-level dedup
create unique index if not exists uniq_obs_point
  on observations (owner_id, metric_id, effective_at, value_num)
  where status = 'normalized';
create index if not exists obs_owner_metric_time
  on observations (owner_id, metric_id, effective_at desc)
  where status = 'normalized';
create index if not exists obs_owner_time on observations (owner_id, effective_at desc);
create index if not exists obs_doc on observations (document_id);
create index if not exists obs_meta_gin on observations using gin (meta);

-- ── Staging: every extracted line item lands here BEFORE normalization ──────
create table if not exists staging_observations (
  id            uuid primary key default gen_random_uuid(),
  owner_id      uuid not null references auth.users(id) on delete cascade,
  extraction_id uuid not null references extractions(id) on delete cascade,
  document_id   uuid not null references source_documents(id) on delete cascade,
  raw_label     text,
  raw_value     text,
  raw_unit      text,
  ref_text      text,
  abnormal_flag text,
  effective_at  timestamptz,
  confidence    numeric,
  resolved_metric_id text references metrics(id),
  resolve_status text not null default 'pending', -- pending|matched|unmatched|ambiguous
  promoted_observation_id uuid references observations(id) on delete set null,
  created_at    timestamptz not null default now()
);
create index if not exists staging_resolve_idx on staging_observations (owner_id, resolve_status);

-- ── Non-time-series clinical records ────────────────────────────────────────
create table if not exists medications (
  id          uuid primary key default gen_random_uuid(),
  owner_id    uuid not null references auth.users(id) on delete cascade,
  name        text not null,
  rxnorm_code text,
  dose_value  numeric,
  dose_unit   text,
  route       text,
  frequency   text,
  status      text not null default 'active',  -- active|discontinued
  started_on  date,
  stopped_on  date,
  document_id uuid references source_documents(id) on delete set null,
  meta        jsonb not null default '{}',
  created_at  timestamptz not null default now()
);

create table if not exists conditions (
  id            uuid primary key default gen_random_uuid(),
  owner_id      uuid not null references auth.users(id) on delete cascade,
  name          text not null,
  icd10_code    text,
  snomed_code   text,
  status        text not null default 'active',
  onset_date    date,
  resolved_date date,
  document_id   uuid references source_documents(id) on delete set null,
  created_at    timestamptz not null default now()
);

-- imaging / narrative docs that resist structuring
create table if not exists clinical_documents (
  id           uuid primary key default gen_random_uuid(),
  owner_id     uuid not null references auth.users(id) on delete cascade,
  doc_type     text not null,        -- 'imaging','visit_note','pathology'
  title        text,
  effective_at timestamptz,
  body_text    text,
  structured   jsonb,                -- {modality, impression, findings[]}
  document_id  uuid references source_documents(id) on delete set null,
  encounter_id uuid references encounters(id) on delete set null,
  created_at   timestamptz not null default now()
);
