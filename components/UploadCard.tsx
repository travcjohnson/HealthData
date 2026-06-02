"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { UploadCloud, Loader2 } from "lucide-react";

type Result = {
  sourceType?: string; duplicate?: boolean; extracted?: number;
  matched?: number; inserted?: number; error?: string;
};

export function UploadCard() {
  const router = useRouter();
  const [busy, setBusy] = useState(false);
  const [results, setResults] = useState<{ name: string; r: Result }[]>([]);

  async function handleFiles(files: FileList | null) {
    if (!files || files.length === 0) return;
    setBusy(true);
    const out: { name: string; r: Result }[] = [];
    for (const file of Array.from(files)) {
      const fd = new FormData();
      fd.append("file", file);
      try {
        const res = await fetch("/api/ingest", { method: "POST", body: fd });
        out.push({ name: file.name, r: await res.json() });
      } catch (e) {
        out.push({ name: file.name, r: { error: (e as Error).message } });
      }
      setResults([...out]);
    }
    setBusy(false);
    router.refresh();
  }

  return (
    <div className="card p-6">
      <label className="flex cursor-pointer flex-col items-center justify-center gap-2 rounded-lg border-2 border-dashed border-slate-200 py-10 text-center hover:border-emerald-400">
        {busy ? (
          <Loader2 className="h-8 w-8 animate-spin text-emerald-600" />
        ) : (
          <UploadCloud className="h-8 w-8 text-slate-400" />
        )}
        <span className="text-sm font-medium text-slate-600">
          {busy ? "Processing…" : "Drop or choose files"}
        </span>
        <span className="text-xs text-slate-400">
          Lab PDFs · Apple Health export.xml · Oura/Whoop/Garmin CSV · lab text
        </span>
        <input
          type="file"
          multiple
          className="hidden"
          disabled={busy}
          accept=".pdf,.xml,.csv,.txt"
          onChange={(e) => handleFiles(e.target.files)}
        />
      </label>

      {results.length > 0 && (
        <ul className="mt-4 space-y-1 text-sm">
          {results.map(({ name, r }, i) => (
            <li key={i} className="flex items-center justify-between">
              <span className="truncate text-slate-600">{name}</span>
              <span className="text-xs text-slate-400">
                {r.error
                  ? `error: ${r.error}`
                  : r.duplicate
                  ? "already imported"
                  : `${r.sourceType}: +${r.inserted ?? 0} of ${r.extracted ?? 0}`}
              </span>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
