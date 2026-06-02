import { getSources } from "@/lib/data/queries";
import { UploadCard } from "@/components/UploadCard";
import { FileText } from "lucide-react";

export const dynamic = "force-dynamic";

export default async function DataPage() {
  const sources = await getSources();
  return (
    <div className="mx-auto max-w-4xl space-y-6">
      <header>
        <h1 className="text-2xl font-semibold">Data & Sources</h1>
        <p className="text-sm text-slate-500">
          Import your mixed health data — it&apos;s parsed, normalized, and de-duplicated automatically.
        </p>
      </header>

      <UploadCard />

      <section>
        <h2 className="mb-3 text-sm font-semibold uppercase tracking-wide text-slate-500">
          Imported documents ({sources.length})
        </h2>
        {sources.length === 0 ? (
          <p className="text-sm text-slate-400">Nothing imported yet.</p>
        ) : (
          <ul className="card divide-y divide-slate-100">
            {sources.map((s) => (
              <li key={s.id} className="flex items-center justify-between px-4 py-3">
                <span className="flex items-center gap-3">
                  <FileText className="h-4 w-4 text-slate-400" />
                  <span>
                    <span className="text-sm font-medium text-slate-700">
                      {s.original_name ?? "(file)"}
                    </span>
                    <span className="ml-2 text-xs text-slate-400">
                      {s.source_lab ?? s.source_type}
                    </span>
                  </span>
                </span>
                <span className="flex items-center gap-3 text-xs text-slate-400">
                  <span className="chip bg-slate-100 text-slate-500">{s.status}</span>
                  {new Date(s.created_at).toLocaleDateString()}
                </span>
              </li>
            ))}
          </ul>
        )}
      </section>
    </div>
  );
}
