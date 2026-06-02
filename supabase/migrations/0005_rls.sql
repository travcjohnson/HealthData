-- ════════════════════════════════════════════════════════════════════════════
-- 0005 · Row Level Security. Every owner-scoped table: owner-only CRUD.
--        Catalog tables: read-only to authenticated. Writes happen via the
--        service_role key in the server-side ingestion pipeline (bypasses RLS).
-- ════════════════════════════════════════════════════════════════════════════

-- Owner-scoped tables get the identical 4-policy pattern, scoped by owner_id.
do $$
declare t text;
begin
  foreach t in array array[
    'profile','source_documents','extractions','encounters','observations',
    'staging_observations','medications','conditions','clinical_documents'
  ] loop
    execute format('alter table %I enable row level security', t);

    execute format('drop policy if exists owner_select on %I', t);
    execute format($p$create policy owner_select on %I for select to authenticated
                      using ((select auth.uid()) = owner_id)$p$, t);

    execute format('drop policy if exists owner_insert on %I', t);
    execute format($p$create policy owner_insert on %I for insert to authenticated
                      with check ((select auth.uid()) = owner_id)$p$, t);

    execute format('drop policy if exists owner_update on %I', t);
    execute format($p$create policy owner_update on %I for update to authenticated
                      using ((select auth.uid()) = owner_id)
                      with check ((select auth.uid()) = owner_id)$p$, t);

    execute format('drop policy if exists owner_delete on %I', t);
    execute format($p$create policy owner_delete on %I for delete to authenticated
                      using ((select auth.uid()) = owner_id)$p$, t);
  end loop;
end $$;

-- Global, read-only catalog tables.
do $$
declare t text;
begin
  foreach t in array array['metrics','metric_aliases','unit_conversions','reference_ranges']
  loop
    execute format('alter table %I enable row level security', t);
    execute format('drop policy if exists read_catalog on %I', t);
    execute format($p$create policy read_catalog on %I for select to authenticated
                      using (true)$p$, t);
  end loop;
end $$;
