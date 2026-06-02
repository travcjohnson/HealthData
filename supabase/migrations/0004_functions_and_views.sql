-- ════════════════════════════════════════════════════════════════════════════
-- 0004 · Flagging functions + enriched / summary views.
--        Views use security_invoker so the querying user's RLS applies (PG15+).
-- ════════════════════════════════════════════════════════════════════════════

-- Color-code a value against a metric's flag bands -> 'red'|'yellow'|'green'|'unknown'
create or replace function fn_metric_flag(p_metric_id text, p_value numeric)
returns text language sql stable as $$
  select case
    when p_value is null then 'unknown'
    when m.id is null then 'unknown'
    when (m.flag_low_red    is not null and p_value <= m.flag_low_red)
      or (m.flag_high_red   is not null and p_value >= m.flag_high_red)   then 'red'
    when (m.flag_low_yellow is not null and p_value <= m.flag_low_yellow)
      or (m.flag_high_yellow is not null and p_value >= m.flag_high_yellow) then 'yellow'
    when m.flag_low_red is null and m.flag_low_yellow is null
     and m.flag_high_red is null and m.flag_high_yellow is null          then 'unknown'
    else 'green'
  end
  from (select * from metrics where id = p_metric_id) m
$$;

-- Within the tighter "optimal" band?
create or replace function fn_metric_optimal(p_metric_id text, p_value numeric)
returns boolean language sql stable as $$
  select case
    when p_value is null then false
    else coalesce(p_value >= m.optimal_low, true)
     and coalesce(p_value <= m.optimal_high, true)
     and (m.optimal_low is not null or m.optimal_high is not null)
  end
  from (select * from metrics where id = p_metric_id) m
$$;

-- Enriched, owner-scoped observation stream (catalog joined, ranges resolved).
create or replace view v_observations
with (security_invoker = true) as
select
  o.*,
  m.display_name,
  m.panel,
  m.canonical_unit,
  m.direction,
  m.hero,
  coalesce(o.ref_low,  m.ref_low)  as eff_ref_low,
  coalesce(o.ref_high, m.ref_high) as eff_ref_high,
  m.optimal_low,
  m.optimal_high,
  fn_metric_flag(o.metric_id, o.value_num)    as flag,
  fn_metric_optimal(o.metric_id, o.value_num) as is_optimal
from observations o
join metrics m on m.id = o.metric_id
where o.status = 'normalized';

-- One row per metric: latest value + previous value + delta + flag.
-- The app computes trend slope client-side from the full series.
create or replace view v_metric_summary
with (security_invoker = true) as
with ranked as (
  select
    o.owner_id, o.metric_id, o.value_num, o.value_unit, o.effective_at,
    row_number() over (partition by o.owner_id, o.metric_id
                       order by o.effective_at desc) as rn,
    count(*)    over (partition by o.owner_id, o.metric_id) as n_readings
  from observations o
  where o.status = 'normalized' and o.value_num is not null
)
select
  cur.owner_id,
  cur.metric_id,
  m.display_name,
  m.panel,
  m.canonical_unit,
  m.direction,
  m.hero,
  m.cadence_days,
  cur.value_num                       as latest_value,
  cur.effective_at                    as latest_at,
  prev.value_num                      as prev_value,
  prev.effective_at                   as prev_at,
  (cur.value_num - prev.value_num)    as delta,
  case when prev.value_num is not null and prev.value_num <> 0
       then round((cur.value_num - prev.value_num) / abs(prev.value_num) * 100, 1)
       end                            as pct_change,
  cur.n_readings,
  coalesce(o.ref_low,  m.ref_low)     as eff_ref_low,
  coalesce(o.ref_high, m.ref_high)    as eff_ref_high,
  m.optimal_low,
  m.optimal_high,
  fn_metric_flag(cur.metric_id, cur.value_num)    as flag,
  fn_metric_optimal(cur.metric_id, cur.value_num) as is_optimal,
  -- days overdue vs recommended cadence (null if no cadence / not overdue)
  case when m.cadence_days is not null
       then greatest(0, (current_date - cur.effective_at::date) - m.cadence_days)
       end                            as days_overdue
from ranked cur
join metrics m on m.id = cur.metric_id
left join ranked prev
  on prev.owner_id = cur.owner_id and prev.metric_id = cur.metric_id and prev.rn = 2
-- pull the latest observation's reported range, if any
left join lateral (
  select ref_low, ref_high from observations
  where owner_id = cur.owner_id and metric_id = cur.metric_id and status='normalized'
  order by effective_at desc limit 1
) o on true
where cur.rn = 1;
