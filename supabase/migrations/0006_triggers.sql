-- ════════════════════════════════════════════════════════════════════════════
-- 0006 · Triggers: auto-create a profile row on signup; maintain updated_at.
-- ════════════════════════════════════════════════════════════════════════════

create or replace function fn_touch_updated_at() returns trigger
  language plpgsql as $$
begin new.updated_at = now(); return new; end $$;

drop trigger if exists profile_touch on profile;
create trigger profile_touch before update on profile
  for each row execute function fn_touch_updated_at();

-- When a new auth user is created, seed an empty profile row.
create or replace function fn_handle_new_user() returns trigger
  language plpgsql security definer set search_path = public as $$
begin
  insert into public.profile (owner_id, display_name)
  values (new.id, coalesce(new.email, 'me'))
  on conflict (owner_id) do nothing;
  return new;
end $$;

-- Guard: auth.users trigger only where we can create it (Supabase & local shim).
do $$
begin
  if exists (select 1 from information_schema.tables
             where table_schema='auth' and table_name='users') then
    drop trigger if exists on_auth_user_created on auth.users;
    create trigger on_auth_user_created after insert on auth.users
      for each row execute function fn_handle_new_user();
  end if;
end $$;
