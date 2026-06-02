import Link from "next/link";
import { notFound } from "next/navigation";
import { ArrowLeft } from "lucide-react";
import { getMetricSeries } from "@/lib/data/queries";
import { TrendChart, type ChartPoint } from "@/components/TrendChart";
import { FLAG_META, fmtValue } from "@/lib/health/status";
import { classifyTrend, slopePerDay } from "@/lib/health/trends";

export const dynamic = "force-dynamic";

export default async function MetricDetail({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;
  const series = await getMetricSeries(id);
  if (series.length === 0) notFound();

  const latest = series[series.length - 1];
  const first = series[0];
  const meta = FLAG_META[latest.flag];
  const unit = latest.value_unit ?? "";
  const points: ChartPoint[] = series
    .filter((o) => o.value_num != null)
    .map((o) => ({
      t: new Date(o.effective_at).getTime(),
      v: Number(o.value_num),
      label: o.effective_at,
    }));

  const trendPts = points.map((p) => ({ t: p.t, v: p.v }));
  const verdict = classifyTrend(trendPts, latest.direction);
  const slope = slopePerDay(trendPts);

  return (
    <div className="mx-auto max-w-4xl space-y-6">
      <Link href="/metrics" className="inline-flex items-center gap-1 text-sm text-slate-500 hover:text-slate-700">
        <ArrowLeft className="h-4 w-4" /> Metrics
      </Link>

      <header className="flex flex-wrap items-end justify-between gap-4">
        <div>
          <h1 className="text-2xl font-semibold">{latest.display_name}</h1>
          <p className="mt-1 flex items-center gap-2 text-sm">
            <span className={`chip ${meta.bg} ${meta.text}`}>
              <span className={`h-1.5 w-1.5 rounded-full ${meta.dot}`} /> {meta.label}
            </span>
            <span className="text-slate-400">
              {series.length} readings since {new Date(first.effective_at).toLocaleDateString()}
            </span>
          </p>
        </div>
        <div className="text-right">
          <p className="text-3xl font-semibold tabular-nums">{fmtValue(latest.value_num, unit)}</p>
          <p className="text-xs text-slate-400 capitalize">trend: {verdict}</p>
        </div>
      </header>

      <div className="card p-4">
        <TrendChart
          points={points}
          unit={unit}
          refLow={latest.eff_ref_low}
          refHigh={latest.eff_ref_high}
          optimalLow={latest.optimal_low}
          optimalHigh={latest.optimal_high}
        />
        <div className="mt-2 flex flex-wrap gap-4 px-2 text-xs text-slate-400">
          {latest.optimal_low != null && (
            <span><span className="mr-1 inline-block h-2 w-3 rounded bg-emerald-200" />optimal band</span>
          )}
          {(latest.eff_ref_low != null || latest.eff_ref_high != null) && (
            <span><span className="mr-1 inline-block h-2 w-3 rounded bg-amber-300" />reference limits</span>
          )}
          {slope != null && (
            <span>slope ≈ {slope > 0 ? "+" : ""}{(slope * 30).toFixed(2)} {unit}/mo</span>
          )}
        </div>
      </div>

      {/* Provenance table — every datapoint traces to its source */}
      <div className="card overflow-hidden">
        <table className="w-full text-sm">
          <thead className="bg-slate-50 text-left text-xs uppercase tracking-wide text-slate-400">
            <tr>
              <th className="px-4 py-2 font-medium">Date</th>
              <th className="px-4 py-2 font-medium">Value</th>
              <th className="px-4 py-2 font-medium">Reported as</th>
              <th className="px-4 py-2 font-medium">Source</th>
              <th className="px-4 py-2 font-medium">Conf.</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-100">
            {[...series].reverse().map((o) => (
              <tr key={o.id}>
                <td className="px-4 py-2 tabular-nums">{new Date(o.effective_at).toLocaleDateString()}</td>
                <td className="px-4 py-2 font-medium tabular-nums">{fmtValue(o.value_num, o.value_unit)}</td>
                <td className="px-4 py-2 text-slate-400">
                  {o.raw_label}{o.raw_unit ? ` (${o.raw_unit})` : ""}
                </td>
                <td className="px-4 py-2"><span className="chip bg-slate-100 text-slate-500">{o.source_type}</span></td>
                <td className="px-4 py-2 text-slate-400">{o.confidence != null ? `${Math.round(o.confidence * 100)}%` : "—"}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
