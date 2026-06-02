/**
 * Derived / composite health metrics. Pure functions, unit-tested.
 * Formulas per clinical spec (CKD-EPI 2021 race-free, Friedewald, HOMA-IR).
 */

/** Body Mass Index (kg/m²). */
export function bmi(weightKg: number, heightCm: number): number | null {
  if (!weightKg || !heightCm) return null;
  const m = heightCm / 100;
  return round(weightKg / (m * m), 1);
}

/** HOMA-IR insulin-resistance index. glucose mg/dL, insulin µIU/mL. */
export function homaIr(glucoseMgDl: number, insulinUIuMl: number): number | null {
  if (!glucoseMgDl || !insulinUIuMl) return null;
  return round((glucoseMgDl * insulinUIuMl) / 405, 2);
}

/** Friedewald LDL (mg/dL). Invalid when TG ≥ 400 → returns null. */
export function friedewaldLdl(tc: number, hdl: number, tg: number): number | null {
  if (tg >= 400) return null;
  if (tc == null || hdl == null || tg == null) return null;
  return round(tc - hdl - tg / 5, 0);
}

/** Non-HDL cholesterol (mg/dL). */
export function nonHdl(tc: number, hdl: number): number | null {
  if (tc == null || hdl == null) return null;
  return round(tc - hdl, 0);
}

/** Triglyceride : HDL ratio (insulin-resistance surrogate). */
export function tgHdlRatio(tg: number, hdl: number): number | null {
  if (!hdl) return null;
  return round(tg / hdl, 2);
}

/** Remnant cholesterol = TC − HDL − LDL (mg/dL). */
export function remnantCholesterol(tc: number, hdl: number, ldl: number): number | null {
  if (tc == null || hdl == null || ldl == null) return null;
  return round(tc - hdl - ldl, 0);
}

/** Waist : height ratio (same units). */
export function waistHeightRatio(waistCm: number, heightCm: number): number | null {
  if (!heightCm) return null;
  return round(waistCm / heightCm, 2);
}

/**
 * eGFR via the 2021 race-free CKD-EPI creatinine equation.
 * @param scr serum creatinine mg/dL, @param age years, @param sex 'male'|'female'
 */
export function egfrCkdEpi2021(
  scr: number,
  age: number,
  sex: "male" | "female",
): number | null {
  if (!scr || !age || !sex) return null;
  const female = sex === "female";
  const kappa = female ? 0.7 : 0.9;
  const alpha = female ? -0.241 : -0.302;
  const ratio = scr / kappa;
  const egfr =
    142 *
    Math.pow(Math.min(ratio, 1), alpha) *
    Math.pow(Math.max(ratio, 1), -1.2) *
    Math.pow(0.9938, age) *
    (female ? 1.012 : 1);
  return round(egfr, 0);
}

function round(n: number, d: number): number {
  const f = Math.pow(10, d);
  return Math.round(n * f) / f;
}
