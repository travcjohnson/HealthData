# Health Command Center

A personal health **dashboard & command center** that turns messy, mixed health
data — lab PDFs from different labs, Apple Health exports, wearable CSVs, doctor
notes — into **clear, color-coded trends** so you can actually understand your
health over time.

Built on **Next.js + Supabase (Postgres, Auth, RLS, Storage)**, deployable to
**Vercel**. Designed for one person, multi-user ready from day one.

> ⚕️ **Not a medical device.** Informational tracking only — reference ranges are
> general adult values, your lab's printed range wins, and nothing here is
> diagnosis or treatment. Always discuss results with a clinician.

---

## What it does

- **Command Center** — a triage-first home screen that surfaces, in priority
  order: what's **out of range now**, what's **trending the wrong way** (in range
  but moving toward risk — the thing a normal patient portal never tells you),
  **borderline** values, **wins** you're making, and labs that are **overdue**.
  A **longevity scorecard** gives ApoB, Lp(a), HbA1c, VO₂max, BP, and fasting
  glucose top billing.
- **Metrics explorer** — every biomarker grouped by panel (lipids, metabolic,
  kidney, liver, thyroid, CBC, vitamins, hormones, wearables…) with latest value,
  flag dot, trend arrow, and an inline sparkline.
- **Metric detail** — full trend chart with **optimal** (green) and **reference**
  (amber) bands, trend slope, and a **provenance table**: every datapoint traces
  back to the exact source file, how it was reported, and extraction confidence.
- **Mixed-data ingestion** — drop in lab PDFs, `export.xml`, or wearable CSVs;
  they're parsed, **normalized to canonical units**, de-duplicated, and matched
  to a 63-metric catalog. Anything it can't confidently map is **surfaced for
  review** rather than silently guessed.

## How it's built

| Layer | Choice |
|---|---|
| App | Next.js 15 (App Router, RSC), TypeScript, Tailwind |
| Data | Supabase Postgres — provenance-first schema, a universal `observations` event table, RLS on every row |
| Auth | Supabase Auth (email/password), middleware-gated routes |
| Charts | Recharts |
| Ingestion | Deterministic parsers (Apple Health XML, wearable CSV) + LLM structured extraction (Anthropic) for lab PDFs/photos, with a deterministic fallback |
| Deploy | Vercel + Supabase |

See **[docs/ARCHITECTURE.md](docs/ARCHITECTURE.md)** for the data model and
**[docs/DEPLOY.md](docs/DEPLOY.md)** to ship it.

---

## Quick start (local)

```bash
npm install
cp .env.example .env.local          # fill in Supabase + DATABASE_URL

# apply schema + 63-metric catalog to your database
npm run db:migrate                  # add `-- --local` against a plain Postgres

# optional: load a realistic demo dataset for an existing auth user
npm run db:seed -- --owner <your-supabase-user-id>

npm run dev                         # http://localhost:3000
```

Import your own data from the **Data & Sources** page, or via CLI:

```bash
npm run ingest -- --owner <user-id> ~/Desktop/labs.pdf apple_export.xml oura.csv
```

## Verify

```bash
npm run typecheck        # tsc
npm run test:parsers     # 28 assertions: parsers, units, derived metrics, trends
npm run build            # next build
```

## Repository layout

```
app/                 Next.js routes (dashboard, metrics, data, records, auth)
components/           UI: tiles, charts, nav, disclaimer
lib/
  supabase/          browser / server / service / middleware clients
  health/            derived metrics, trend math, status mapping
  ingest/            pure normalization (units, alias resolution)
  data/queries.ts    RLS-scoped data access + triage builder
scripts/
  db/                migrate & seed runners (+ local auth shim)
  ingest/            parsers (Apple Health, CSV, lab text/LLM), CLI, tests
supabase/migrations/ 10 ordered SQL migrations (schema, catalog, RLS, grants)
docs/                ARCHITECTURE.md, DEPLOY.md
```
