# Deploy

End-to-end: connect your Supabase project, load the schema, deploy to Vercel.

## 1. Supabase

You already have a project (`NEXT_PUBLIC_SUPABASE_URL` points at it). Grab these
from **Project Settings → API** and **→ Database**:

- `NEXT_PUBLIC_SUPABASE_URL`
- `NEXT_PUBLIC_SUPABASE_PUBLISHABLE_KEY` (anon / publishable — safe in browser)
- `SUPABASE_SECRET_KEY` (service-role / secret — **server only**)
- `DATABASE_URL` (Connection string → URI; use the **session/pooler** URI for
  serverless, or the direct `db.<ref>.supabase.co:5432` URI for migrations)

### Apply the schema + metric catalog

Either with the Supabase CLI:

```bash
supabase link --project-ref <your-ref>
supabase db push        # applies supabase/migrations/*.sql in order
```

…or directly (no Docker needed):

```bash
DATABASE_URL='postgresql://postgres:<pw>@db.<ref>.supabase.co:5432/postgres' \
  npm run db:migrate
```

This creates the schema, RLS policies, role grants, the **63-metric catalog**,
synonyms, unit conversions, and the private `health-raw` storage bucket. The
migrations are idempotent — safe to re-run to pick up catalog updates.

> On Supabase you do **not** apply `scripts/db/_local_shim.sql` — that exists
> only so the same migrations run against a plain local Postgres. `auth.users`,
> `auth.uid()`, and the `authenticated`/`anon`/`service_role` roles already
> exist on Supabase.

### Auth

Email/password is enabled by default. In **Authentication → URL Configuration**
set the Site URL to your Vercel domain and add
`https://<domain>/auth/callback` as a redirect URL (and `http://localhost:3000/auth/callback`
for local dev). Sign up once — a `profile` row is auto-created by trigger.

### (Optional) seed demo data

```bash
DATABASE_URL=... npm run db:seed -- --owner <your-auth-user-id>
```

(omit `--create-user`; your user already exists). This loads ~3 years of
realistic labs + 120 days of wearable data so the dashboard has a story to tell.

## 2. Vercel

1. Import the repo in Vercel (framework auto-detected as Next.js).
2. Add the four env vars above **plus**:
   - `ANTHROPIC_API_KEY` and `INGEST_LLM_MODEL` (optional — enables reading lab
     PDFs/photos; without it the deterministic text parser is used).
   - `NEXT_PUBLIC_APP_URL=https://<your-domain>`
3. Deploy.

The `/api/ingest` route runs on the Node.js runtime (PDF parsing, direct
Postgres) — `maxDuration` is set to 60s for large lab PDFs / Apple exports.

## 3. Load your real data

- **In-app:** Data & Sources → drop lab PDFs, `export.xml`, wearable CSVs.
- **Bulk/CLI:** `npm run ingest -- --owner <user-id> <files...>`

Large Apple Health `export.xml` files (often >100 MB) are best ingested via the
CLI. Unmatched line items land in `staging_observations` with
`resolve_status='unmatched'` — add an alias to `metric_aliases` and re-run to
teach the system permanently.

## Environment / network notes

This project was scaffolded and verified in a sandbox where outbound access to
`*.supabase.co` was blocked, so migrations/seed must be run from a network that
can reach your project (your laptop or Vercel's build). Everything else —
schema, RLS, parsers, build — was verified locally against Postgres 16.
