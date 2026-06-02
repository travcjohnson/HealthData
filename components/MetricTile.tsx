import Link from "next/link";
import { Sparkline } from "./Sparkline";
import { FLAG_META, fmtValue, trendArrow } from "@/lib/health/status";
import type { MetricSummary } from "@/lib/types";

export function MetricTile({
  m,
  spark,
}: {
  m: MetricSummary;
  spark?: number[];
}) {
  const meta = FLAG_META[m.flag];
  const { arrow, good } = trendArrow(m.pct_change, m.direction);
  const arrowColor =
    good === null ? "text-slate-400" : good ? "text-emerald-600" : "text-red-600";

  return (
    <Link
      href={`/metrics/${m.metric_id}`}
      className={`card block p-4 ring-1 ${meta.ring} transition hover:shadow-md`}
    >
      <div className="flex items-start justify-between">
        <div className="min-w-0">
          <p className="truncate text-sm font-medium text-slate-600">{m.display_name}</p>
          <p className="mt-1 text-2xl font-semibold tabular-nums">
            {fmtValue(m.latest_value)}
            <span className="ml-1 text-sm font-normal text-slate-400">{m.canonical_unit}</span>
          </p>
        </div>
        <span className={`mt-1 h-2.5 w-2.5 flex-shrink-0 rounded-full ${meta.dot}`} title={meta.label} />
      </div>

      <div className="mt-3 flex items-center justify-between">
        <span className={`text-xs font-medium ${arrowColor}`}>
          {arrow} {m.pct_change != null ? `${Math.abs(m.pct_change)}%` : "—"}
        </span>
        {spark && spark.length > 1 && (
          <Sparkline values={spark} className={meta.text} />
        )}
      </div>

      {(m.eff_ref_low != null || m.eff_ref_high != null) && (
        <p className="mt-2 text-[11px] text-slate-400">
          ref {m.eff_ref_low ?? "–"}–{m.eff_ref_high ?? "–"} {m.canonical_unit}
        </p>
      )}
    </Link>
  );
}
