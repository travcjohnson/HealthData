import { getSummaries, getSparkValues } from "@/lib/data/queries";
import { MetricTile } from "@/components/MetricTile";
import { PANEL_LABELS, PANEL_ORDER } from "@/lib/health/status";

export const dynamic = "force-dynamic";

export default async function MetricsPage() {
  const [rows, sparks] = await Promise.all([getSummaries(), getSparkValues()]);
  const byPanel = new Map<string, typeof rows>();
  for (const m of rows) {
    const arr = byPanel.get(m.panel) ?? [];
    arr.push(m);
    byPanel.set(m.panel, arr);
  }
  const panels = [...byPanel.keys()].sort(
    (a, b) => (PANEL_ORDER.indexOf(a) + 1 || 99) - (PANEL_ORDER.indexOf(b) + 1 || 99),
  );

  return (
    <div className="mx-auto max-w-6xl space-y-8">
      <h1 className="text-2xl font-semibold">Metrics</h1>
      {rows.length === 0 && <p className="text-slate-500">No metrics yet — import data first.</p>}
      {panels.map((panel) => (
        <section key={panel}>
          <h2 className="mb-3 text-sm font-semibold uppercase tracking-wide text-slate-500">
            {PANEL_LABELS[panel] ?? panel}
          </h2>
          <div className="grid grid-cols-2 gap-3 sm:grid-cols-3 lg:grid-cols-4">
            {byPanel
              .get(panel)!
              .sort((a, b) => a.display_name.localeCompare(b.display_name))
              .map((m) => (
                <MetricTile key={m.metric_id} m={m} spark={sparks.get(m.metric_id)} />
              ))}
          </div>
        </section>
      ))}
    </div>
  );
}
