-- ════════════════════════════════════════════════════════════════════════════
-- 0009 · Supabase Storage: private bucket for raw health files, owner-scoped.
--        Guarded so it is a no-op on a plain local Postgres (no storage schema).
-- ════════════════════════════════════════════════════════════════════════════
do $$
begin
  if exists (select 1 from information_schema.tables
             where table_schema='storage' and table_name='buckets') then

    insert into storage.buckets (id, name, public)
    values ('health-raw','health-raw', false)
    on conflict (id) do nothing;

    -- a user may only touch objects whose first path segment is their uid
    drop policy if exists own_health_files on storage.objects;
    create policy own_health_files on storage.objects for all to authenticated
      using (bucket_id = 'health-raw'
             and (storage.foldername(name))[1] = (select auth.uid())::text)
      with check (bucket_id = 'health-raw'
             and (storage.foldername(name))[1] = (select auth.uid())::text);
  end if;
end $$;
