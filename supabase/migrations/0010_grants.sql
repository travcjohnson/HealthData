-- ════════════════════════════════════════════════════════════════════════════
-- 0010 · Role grants. Supabase grants these to anon/authenticated/service_role
--        by default; we declare them explicitly so the schema is self-contained
--        and behaves identically on a plain Postgres. RLS still filters rows.
-- ════════════════════════════════════════════════════════════════════════════
do $$
begin
  if exists (select 1 from pg_roles where rolname='authenticated') then
    grant usage on schema public to authenticated;
    grant select, insert, update, delete on all tables in schema public to authenticated;
    grant usage, select on all sequences in schema public to authenticated;
    grant execute on all functions in schema public to authenticated;
  end if;
  if exists (select 1 from pg_roles where rolname='anon') then
    grant usage on schema public to anon;
    grant execute on all functions in schema public to anon;
  end if;
  if exists (select 1 from pg_roles where rolname='service_role') then
    grant usage on schema public to service_role;
    grant all on all tables in schema public to service_role;
    grant all on all sequences in schema public to service_role;
    grant execute on all functions in schema public to service_role;
  end if;
end $$;
