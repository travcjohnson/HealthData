import Link from "next/link";
import { AlertTriangle, TrendingDown, CheckCircle2, Clock, Trophy } from "lucide-react";
import { getSummaries, getSparkValues, buildTriage } from "@/lib/data/queries";
import { MetricTile } from "@/components/MetricTile";
import { FLAG_META, fmtValue, trendArrow } from "@/lib/health/status";
import type { MetricSummary } from "@/lib/types";

export const dynamic = "force-dynamic";

export default async function DashboardPage() {
  const [rows, sparks] = await Promise.all([getSummaries(), getSparkValues()]);
  const t = buildTriage(rows);
  const today = new Date().toLocaleDateString(undefined, {
    weekday: "long", month: "long", day: "numeric",
  });

  const empty = rows.length === 0;

  return (
    <div className="mx-auto max-w-6xl space-y-8">
      <header className="flex items-end justify-between">
        <div>
          <h1 className="text-2xl font-semibold">Command Center</h1>
          <p className="text-sm text-slate-500">{today}</p>
        </div>
        <Link href="/metrics" className="text-sm font-medium text-emerald-700 hover:underline">
          View all metrics →
        </Link>
      </header>

      {empty && (
        <div className="card p-8 text-center">
          <p className="text-slate-600">No health data yet.</p>
          <Link href="/data" className="mt-3 inline-block rounded-lg bg-emerald-600 px-4 py-2 text-sm font-medium text-white">
            Import your data
          </Link>
        </div>
      )}

      {!empty && (
        <>
          {/* Longevity scorecard — heroes get top billing */}
          {t.heroes.length > 0 && (
            <section>
              <SectionHead icon={Trophy} title="Longevity scorecard" tone="text-emerald-700"
                hint="Highest-leverage drivers of healthspan" />
              <div className="grid grid-cols-2 gap-3 sm:grid-cols-3 lg:grid-cols-5">
                {t.heroes.map((m) => (
                  <MetricTile key={m.metric_id} m={m} spark={sparks.get(m.metric_id)} />
                ))}
              </div>
            </section>
          )}

          <div className="grid gap-6 lg:grid-cols-2">
            {t.outOfRange.length > 0 && (
              <TriageList icon={AlertTriangle} tone="text-red-600"
                title="Needs attention" rows={t.outOfRange} />
            )}
            {t.trendingWrong.length > 0 && (
              <TriageList icon={TrendingDown} tone="text-amber-600"
                title="Trending the wrong way" rows={t.trendingWrong}
                hint="In range now, but moving toward risk" />
            )}
            {t.borderline.length > 0 && (
              <TriageList icon={AlertTriangle} tone="text-amber-500"
                title="Borderline" rows={t.borderline} />
            )}
            {t.wins.length > 0 && (
              <TriageList icon={CheckCircle2} tone="text-emerald-600"
                title="Wins — improving" rows={t.wins} />
            )}
            {t.overdue.length > 0 && (
              <TriageList icon={Clock} tone="text-slate-500"
                title="Due for a recheck" rows={t.overdue} overdue />
            )}
          </div>
        </>
      )}
    </div>
  );
}

function SectionHead({
  icon: Icon, title, tone, hint,
}: { icon: any; title: string; tone: string; hint?: string }) {
  return (
    <div className="mb-3 flex items-center gap-2">
      <Icon className={`h-5 w-5 ${tone}`} />
      <h2 className="font-semibold">{title}</h2>
      {hint && <span className="text-xs text-slate-400">· {hint}</span>}
    </div>
  );
}

function TriageList({
  icon, tone, title, rows, hint, overdue,
}: {
  icon: any; tone: string; title: string; rows: MetricSummary[]; hint?: string; overdue?: boolean;
}) {
  return (
    <section className="card p-4">
      <SectionHead icon={icon} title={title} tone={tone} hint={hint} />
      <ul className="divide-y divide-slate-100">
        {rows.map((m) => {
          const meta = FLAG_META[m.flag];
          const { arrow, good } = trendArrow(m.pct_change, m.direction);
          return (
            <li key={m.metric_id}>
              <Link href={`/metrics/${m.metric_id}`}
                className="flex items-center justify-between py-2 hover:opacity-80">
                <span className="flex items-center gap-2">
                  <span className={`h-2 w-2 rounded-full ${meta.dot}`} />
                  <span className="text-sm font-medium text-slate-700">{m.display_name}</span>
                </span>
                <span className="flex items-center gap-3 text-sm tabular-nums">
                  {overdue ? (
                    <span className="text-xs text-slate-500">{m.days_overdue}d overdue</span>
                  ) : (
                    <span className={good === false ? "text-red-600" : good ? "text-emerald-600" : "text-slate-400"}>
                      {arrow} {m.pct_change != null ? `${Math.abs(m.pct_change)}%` : ""}
                    </span>
                  )}
                  <span className="font-semibold">{fmtValue(m.latest_value, m.canonical_unit)}</span>
                </span>
              </Link>
            </li>
          );
        })}
      </ul>
    </section>
  );
}
