-- ────────────────────────────────────────────────────────────────────────────
-- LOCAL-ONLY shim. Recreates the minimal pieces of Supabase's `auth` schema so
-- the real migrations apply unchanged against a plain Postgres for verification.
-- On real Supabase, `auth.users`, `auth.uid()`, `auth.role()` already exist and
-- this file is NOT applied. (See scripts/db/migrate.ts --local)
-- ────────────────────────────────────────────────────────────────────────────
create schema if not exists auth;

create table if not exists auth.users (
  id    uuid primary key default gen_random_uuid(),
  email text unique
);

-- Mirrors Supabase: derives the current user id from the request JWT claims,
-- which our RLS tests set via `set local request.jwt.claims = '{"sub": "..."}'`.
create or replace function auth.uid() returns uuid
  language sql stable as $$
  select nullif(
    coalesce(
      current_setting('request.jwt.claim.sub', true),
      (nullif(current_setting('request.jwt.claims', true), '')::jsonb ->> 'sub')
    ), ''
  )::uuid
$$;

create or replace function auth.role() returns text
  language sql stable as $$
  select coalesce(
    nullif(current_setting('request.jwt.claim.role', true), ''),
    (nullif(current_setting('request.jwt.claims', true), '')::jsonb ->> 'role'),
    'anon'
  )
$$;

-- `authenticated` / `anon` / `service_role` exist on Supabase; create locally.
do $$
begin
  if not exists (select from pg_roles where rolname = 'authenticated') then
    create role authenticated nologin;
  end if;
  if not exists (select from pg_roles where rolname = 'anon') then
    create role anon nologin;
  end if;
  if not exists (select from pg_roles where rolname = 'service_role') then
    create role service_role nologin bypassrls;
  end if;
end $$;
