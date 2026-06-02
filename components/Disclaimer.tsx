import { ShieldAlert } from "lucide-react";

export function Disclaimer() {
  return (
    <footer className="border-t border-slate-200 bg-white px-6 py-3">
      <p className="flex items-start gap-2 text-[11px] leading-relaxed text-slate-400">
        <ShieldAlert className="mt-0.5 h-3.5 w-3.5 flex-shrink-0" />
        <span>
          This dashboard is an informational tool for personal tracking — <strong>not a medical
          device</strong> and not a source of diagnosis or treatment. Reference ranges are general
          adult values and may not apply to you; your lab&apos;s own printed range takes precedence.
          Do not start, stop, or change any treatment based on this dashboard. Discuss all results
          with a qualified clinician. For emergencies, call your local emergency number.
        </span>
      </p>
    </footer>
  );
}
