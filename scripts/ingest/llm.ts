/**
 * LLM structured extraction for unstructured lab reports (PDF text or images).
 * Uses Anthropic tool-calling to force a strict JSON schema. The model only
 * transcribes what it sees (verbatim labels) — resolution to canonical metrics
 * happens deterministically in the normalizer, keeping the model auditable.
 *
 * Falls back to the deterministic text parser when ANTHROPIC_API_KEY is unset.
 */
import type { ExtractedItem } from "@/lib/ingest/normalize";
import { parseLabText } from "./parsers/labText";

export const PROMPT_VERSION = "lab-extract-v1";

const EXTRACT_TOOL = {
  name: "record_lab_results",
  description: "Record every lab result, vital, or measurement found in the document.",
  input_schema: {
    type: "object" as const,
    properties: {
      collected_date: { type: "string", description: "ISO date the specimen was collected/reported, if present" },
      source_lab: { type: "string", description: "Lab or facility name if present (Quest, LabCorp, ...)" },
      results: {
        type: "array",
        items: {
          type: "object",
          properties: {
            label: { type: "string", description: "Test name EXACTLY as printed" },
            value: { type: "string", description: "Numeric or qualitative result as printed" },
            unit: { type: "string" },
            reference_range: { type: "string", description: "Reference range as printed, e.g. '70-99' or '<5.7'" },
            flag: { type: "string", description: "Abnormal flag as printed: H, L, HH, A, ..." },
            confidence: { type: "number", description: "0-1 confidence this row was read correctly" },
          },
          required: ["label", "value"],
        },
      },
    },
    required: ["results"],
  },
};

export type LlmExtraction = {
  items: ExtractedItem[];
  extractor: string;
  extractorKind: "llm" | "deterministic";
  sourceLab: string | null;
  raw: unknown;
};

/** Extract from already-extracted PDF text (or any plain text). */
export async function extractFromText(text: string): Promise<LlmExtraction> {
  const apiKey = process.env.ANTHROPIC_API_KEY;
  if (!apiKey) {
    return {
      items: parseLabText(text),
      extractor: "labText-v1",
      extractorKind: "deterministic",
      sourceLab: null,
      raw: { fallback: "no ANTHROPIC_API_KEY; used deterministic parser" },
    };
  }

  const { default: Anthropic } = await import("@anthropic-ai/sdk");
  const client = new Anthropic({ apiKey });
  const model = process.env.INGEST_LLM_MODEL || "claude-sonnet-4-6";

  const resp = await client.messages.create({
    model,
    max_tokens: 4096,
    tools: [EXTRACT_TOOL],
    tool_choice: { type: "tool", name: "record_lab_results" },
    messages: [
      {
        role: "user",
        content:
          "Extract every lab result/vital/measurement from this report. " +
          "Transcribe labels and values EXACTLY as printed; do not normalize, " +
          "interpret, or invent values. Include the reference range and any " +
          "abnormal flag if shown.\n\n----- REPORT TEXT -----\n" +
          text.slice(0, 60_000),
      },
    ],
  });

  const toolUse = resp.content.find((c) => c.type === "tool_use") as
    | { input: { results?: any[]; collected_date?: string; source_lab?: string } }
    | undefined;
  const data = toolUse?.input ?? { results: [] };
  const collected = data.collected_date || null;
  const items: ExtractedItem[] = (data.results ?? []).map((r: any) => ({
    label: r.label,
    value: r.value,
    unit: r.unit ?? null,
    refText: r.reference_range ?? null,
    flag: r.flag ?? null,
    effectiveAt: collected,
    confidence: typeof r.confidence === "number" ? r.confidence : 0.85,
  }));

  return {
    items,
    extractor: model,
    extractorKind: "llm",
    sourceLab: data.source_lab ?? null,
    raw: data,
  };
}

/** Extract from a PDF file path (pdf-parse → text → LLM/deterministic). */
export async function extractFromPdf(buffer: Buffer): Promise<LlmExtraction> {
  // import the inner lib path to avoid pdf-parse's debug-mode auto-run on import
  const pdfParse = (await import("pdf-parse/lib/pdf-parse.js")).default as (
    b: Buffer,
  ) => Promise<{ text: string }>;
  const { text } = await pdfParse(buffer);
  return extractFromText(text);
}
