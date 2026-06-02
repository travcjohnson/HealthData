/**
 * Deterministic lab-report line parser. Works on text extracted from a lab PDF
 * (via pdf-parse) when no LLM key is configured, or as a cross-check. Heuristic:
 * looks for "<label> <value> <unit?> <flag?> <reference?>" patterns per line.
 */
import type { ExtractedItem } from "@/lib/ingest/normalize";

const LINE_RE = new RegExp(
  [
    "^\\s*",
    "(?<label>[A-Za-z][A-Za-z0-9 ,()/+.\\-]*?)", // label (lazy: stop at first value)
    "[\\s.:]+",
    "(?<value>-?\\d+(?:\\.\\d+)?)", // value
    "\\s*",
    "(?<unit>%|x?10\\^?\\d+/?[a-zA-Zµ]+|[a-zA-Zµ°]+(?:/[a-zA-Z0-9.]+){0,2})?", // unit
    "\\s*",
    "(?<flag>HH|LL|H|L|A)?", // abnormal flag
    "\\s*",
    "(?<ref>(?:[<>≤≥]\\s*-?\\d+(?:\\.\\d+)?)|(?:-?\\d+(?:\\.\\d+)?\\s*[-–]\\s*-?\\d+(?:\\.\\d+)?))?", // ref range
    "\\s*$",
  ].join(""),
);

// dates like "Collected: 03/14/2024" or "Date of Service 2024-03-14"
const DATE_RE = /(\d{4}-\d{2}-\d{2})|(\d{1,2}\/\d{1,2}\/\d{2,4})/;

function findDocumentDate(text: string): string | null {
  const labeled = text.match(
    /(?:collected|drawn|reported|date of service|service date|specimen)[^\d]{0,20}(\d{4}-\d{2}-\d{2}|\d{1,2}\/\d{1,2}\/\d{2,4})/i,
  );
  const m = labeled?.[1] ?? text.match(DATE_RE)?.[0];
  if (!m) return null;
  const d = new Date(m.includes("/") ? m : `${m}T12:00:00Z`);
  return isNaN(d.getTime()) ? null : d.toISOString();
}

export function parseLabText(text: string): ExtractedItem[] {
  const docDate = findDocumentDate(text);
  const items: ExtractedItem[] = [];
  for (const line of text.split(/\r?\n/)) {
    const trimmed = line.trim();
    if (trimmed.length < 4 || trimmed.length > 80) continue;
    const m = LINE_RE.exec(trimmed);
    if (!m?.groups) continue;
    const label = m.groups.label.trim();
    // reject obvious non-metric lines (addresses, page numbers, phone, ids)
    if (/^(page|fax|phone|tel|patient|account|dob|mrn|ordered|provider)/i.test(label)) continue;
    if (label.length < 2) continue;
    const value = m.groups.value;
    if (value == null) continue;
    items.push({
      label,
      value,
      unit: m.groups.unit?.trim() || null,
      flag: m.groups.flag || null,
      refText: m.groups.ref?.trim() || null,
      effectiveAt: docDate,
      confidence: 0.6, // deterministic text parse is lower-trust than LLM/structured
    });
  }
  return items;
}
