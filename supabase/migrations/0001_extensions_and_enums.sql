-- ════════════════════════════════════════════════════════════════════════════
-- 0001 · Extensions & enumerated types
-- ════════════════════════════════════════════════════════════════════════════
create extension if not exists "pgcrypto";   -- gen_random_uuid()

do $$ begin
  create type source_type as enum
    ('lab_pdf','apple_health','oura','whoop','garmin',
     'manual','imaging','doctor_note','photo','cgm','other');
exception when duplicate_object then null; end $$;

do $$ begin
  create type value_kind as enum
    ('quantity','ratio','code','text','boolean','range');
exception when duplicate_object then null; end $$;

do $$ begin
  create type ingest_status as enum
    ('uploaded','extracting','staged','normalized','error','superseded');
exception when duplicate_object then null; end $$;

-- direction of "good" for a metric, drives auto color-coding & trend meaning
do $$ begin
  create type metric_direction as enum
    ('higher_better','lower_better','target_band');
exception when duplicate_object then null; end $$;
