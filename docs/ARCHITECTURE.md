# Architecture

## Design principles

1. **Provenance is immutable; canonical values are derived.** Raw label, raw
   value, and raw unit (exactly as printed/exported) are never overwritten.
   `value_num` in canonical units is always recomputable from staging — so we can
   re-normalize or re-extract with a better model without losing ground truth.
2. **One universal event table.** Labs, vitals, and wearable metrics are
   structurally identical (metric + time + value + unit + range), so they all
   live in `observations`. `metrics.panel` gives logical grouping without
   physical fragmentation — and makes cross-domain correlation trivial.
3. **A metric catalog collapses the chaos.** Hundreds of lab/wearable synonyms
   resolve to ~63 canonical metrics with canonical units, flag bands, and
   reference ranges. This is the single source of truth for "what was measured."
4. **Multi-tenant from row one.** Every row carries `owner_id`; RLS is on
   everywhere. Single-user today, zero migration to add users later.

## Data model (key tables)

```
metrics              canonical metric defs: unit, direction, ref/optimal/flag bands, cadence
metric_aliases       every printed synonym  → metric_id  (normalized, indexed)
unit_conversions     deterministic, auditable unit canonicalization
reference_ranges     sex/age-specific bands, resolved against profile

source_documents     one row per uploaded file (sha256 dedup, storage path, status)
extractions          one row per extraction run (LLM or deterministic; raw_output)
staging_observations every extracted line item BEFORE normalization (resolve_status)
observations         canonical event stream — value_num + raw_* + ref_* + provenance + supersedes
encounters           groups observations (a lab draw, a visit)

profile              birth_date, sex, height → drives age/sex ranges & derived metrics
medications          conditions          clinical_documents (imaging/notes: typed header + JSONB/text)
```

### Flagging

Each metric carries four thresholds. A value is **red** if
`≤ flag_low_red` or `≥ flag_high_red`; **yellow** if `≤ flag_low_yellow` or
`≥ flag_high_yellow`; else **green**. Null bounds = unbounded on that side. This
one model cleanly expresses `lower_better`, `higher_better`, and `target_band`
metrics. `direction` decouples *trend meaning* (a falling LDL is improving; a
falling HDL is worsening) from the color.

Views (`v_observations`, `v_metric_summary`) are `security_invoker`, so the
querying user's RLS applies. `v_metric_summary` computes latest/prev/delta/%
change and days-overdue per metric in one query.

## Ingestion pipeline

```
Upload ─▶ source_documents (sha256 dedup, status)
   │
   ├── export.xml      → deterministic Apple Health parser (daily aggregation)
   ├── *.csv           → deterministic wearable mapper (Oura/Whoop/Garmin)
   └── *.pdf / *.txt   → LLM structured extraction (strict JSON tool schema)
                          ↳ deterministic regex parser if no ANTHROPIC_API_KEY
   │
   ▼ extractions (raw_output kept for re-processing)
   ▼ staging_observations (verbatim line items)
   ▼ normalize: alias → metric_id, unit → canonical, parse ref range, confidence
   ▼ observations (canonical, deduped on owner+metric+time+value)
```

The LLM only transcribes what it sees (verbatim labels); resolution to canonical
metrics happens **deterministically** in the normalizer, keeping the model dumb
and auditable. No conversion path → the value is held back for review, never
guessed. Conflicting values for the same metric+time use `supersedes` (old row
flagged `superseded`, charts show only `normalized`).

## Health logic (`lib/health`)

- `derived.ts` — BMI, HOMA-IR, **CKD-EPI 2021** (race-free) eGFR, Friedewald LDL
  (invalid ≥400 TG), non-HDL, TG/HDL, remnant cholesterol, waist:height.
- `trends.ts` — OLS slope/day, rolling average, % change, and a clinical
  `classifyTrend` that reads `direction` to call improving vs worsening, plus
  `trendingTowardRisk` (the command-center differentiator).
- `status.ts` — flag → color/label, panel ordering, value/trend formatting.

## Sources behind the catalog

Reference ranges and flag thresholds follow preventive-medicine guidelines:
ADA 2024 (glucose/HbA1c), ACC/AHA 2017 (blood pressure), NLA/ESC (ApoB, Lp(a)),
CKD-EPI 2021 (eGFR). They are general adult values for color-coding only — a
report's own printed range overrides on import, and none of it is diagnosis.
