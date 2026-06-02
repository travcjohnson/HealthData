// pdf-parse ships no types for its inner lib path; we import it to avoid the
// package's debug-mode auto-run on import.
declare module "pdf-parse/lib/pdf-parse.js" {
  function pdfParse(buffer: Buffer): Promise<{ text: string; numpages: number }>;
  export default pdfParse;
}
