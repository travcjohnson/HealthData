import { getMedications, getConditions } from "@/lib/data/queries";
import { Pill, Stethoscope } from "lucide-react";

export const dynamic = "force-dynamic";

export default async function RecordsPage() {
  const [meds, conditions] = await Promise.all([getMedications(), getConditions()]);

  return (
    <div className="mx-auto max-w-4xl space-y-8">
      <h1 className="text-2xl font-semibold">Medications & Conditions</h1>

      <section>
        <h2 className="mb-3 flex items-center gap-2 text-sm font-semibold uppercase tracking-wide text-slate-500">
          <Pill className="h-4 w-4" /> Medications
        </h2>
        {meds.length === 0 ? (
          <p className="text-sm text-slate-400">No medications recorded.</p>
        ) : (
          <ul className="card divide-y divide-slate-100">
            {meds.map((m: any) => (
              <li key={m.id} className="flex items-center justify-between px-4 py-3 text-sm">
                <span className="font-medium text-slate-700">
                  {m.name}
                  {m.dose_value ? ` · ${m.dose_value}${m.dose_unit ?? ""}` : ""}
                </span>
                <span className="text-xs text-slate-400">
                  {m.frequency ?? ""} <span className="chip bg-slate-100 text-slate-500">{m.status}</span>
                </span>
              </li>
            ))}
          </ul>
        )}
      </section>

      <section>
        <h2 className="mb-3 flex items-center gap-2 text-sm font-semibold uppercase tracking-wide text-slate-500">
          <Stethoscope className="h-4 w-4" /> Conditions
        </h2>
        {conditions.length === 0 ? (
          <p className="text-sm text-slate-400">No conditions recorded.</p>
        ) : (
          <ul className="card divide-y divide-slate-100">
            {conditions.map((c: any) => (
              <li key={c.id} className="flex items-center justify-between px-4 py-3 text-sm">
                <span className="font-medium text-slate-700">{c.name}</span>
                <span className="text-xs text-slate-400">
                  {c.onset_date ?? ""} <span className="chip bg-slate-100 text-slate-500">{c.status}</span>
                </span>
              </li>
            ))}
          </ul>
        )}
      </section>

      <p className="text-xs text-slate-400">
        Medications & conditions are extracted from imported documents, or can be added manually
        (coming soon). Imaging and visit notes are stored as clinical documents.
      </p>
    </div>
  );
}
