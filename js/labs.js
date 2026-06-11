/* Ridgeline — biomarker reference catalog (generic, longevity-oriented).
 *
 * Deliberately contains NO personal data: this file ships on a public site.
 * The personal lab record (report timeline, recovered values, analysis
 * findings) lives in a private JSON file imported on the Data tab and is
 * stored only in this browser's localStorage.
 *
 * ⚕️ Informational tracking only — never diagnosis or treatment. Your lab's
 * printed reference range wins. Discuss anything here with a clinician.
 */
var LABS_CATALOG = [
  { slug: "apob", name: "ApoB", cat: "Cardiovascular", unit: "mg/dL", target: "< 80", why: "Single best causal marker of atherosclerotic risk. Aggressive target <60." },
  { slug: "lpa", name: "Lp(a)", cat: "Cardiovascular", unit: "nmol/L", target: "< 75", why: "Genetically determined, lifelong-stable. <75 nmol/L is low-risk." },
  { slug: "ldl_c", name: "LDL-C", cat: "Cardiovascular", unit: "mg/dL", target: "< 100", why: "Longevity-oriented target <70." },
  { slug: "triglycerides", name: "Triglycerides", cat: "Cardiovascular", unit: "mg/dL", target: "< 100", why: "Metabolic + cardiovascular marker." },
  { slug: "omega_index", name: "Omega-3 Index", cat: "Cardiovascular", unit: "%", target: "≥ 5.5", why: "Lower sudden-cardiac-death risk." },
  { slug: "hba1c", name: "HbA1c", cat: "Metabolic", unit: "%", target: "< 5.4", why: "90-day average glucose; prediabetes ≥5.7." },
  { slug: "glucose_fasting", name: "Fasting Glucose", cat: "Metabolic", unit: "mg/dL", target: "70–90", why: "" },
  { slug: "insulin_fasting", name: "Fasting Insulin", cat: "Metabolic", unit: "µIU/mL", target: "< 5", why: "Earliest insulin-resistance marker." },
  { slug: "lp_ir", name: "LP-IR Score", cat: "Metabolic", unit: "score", target: "< 27", why: "Flags insulin resistance before glucose moves." },
  { slug: "testosterone_total", name: "Testosterone, Total", cat: "Hormones", unit: "ng/dL", target: "500–900", why: "Mid-upper range favorable. Assays differ across labs." },
  { slug: "shbg", name: "SHBG", cat: "Hormones", unit: "nmol/L", target: "20–55", why: "High SHBG binds free testosterone." },
  { slug: "dhea_s", name: "DHEA-S", cat: "Hormones", unit: "µg/dL", target: "mid-range", why: "" },
  { slug: "tsh", name: "TSH", cat: "Thyroid", unit: "µIU/mL", target: "0.5–2.5", why: "High-normal worth monitoring." },
  { slug: "free_t3", name: "Free T3", cat: "Thyroid", unit: "pg/mL", target: "3.0–4.4", why: "" },
  { slug: "tpo_ab", name: "TPO Antibodies", cat: "Thyroid", unit: "IU/mL", target: "< 34", why: "" },
  { slug: "wbc", name: "White Blood Cells", cat: "Blood", unit: "×10³/µL", target: "4.0–7.0", why: "Persistent lows deserve a clinical conversation." },
  { slug: "neutrophils_abs", name: "Absolute Neutrophils", cat: "Blood", unit: "/µL", target: "> 2000", why: "ANC below 1,000 warrants attention." },
  { slug: "ferritin", name: "Ferritin", cat: "Iron", unit: "ng/mL", target: "50–150", why: "Iron stores + inflammation marker." },
  { slug: "iron_sat", name: "Iron Saturation", cat: "Iron", unit: "%", target: "25–45", why: "" },
  { slug: "hscrp", name: "hs-CRP", cat: "Inflammation", unit: "mg/L", target: "< 0.5", why: "Silent inflammation; ideally <0.3." },
  { slug: "homocysteine", name: "Homocysteine", cat: "Inflammation", unit: "µmol/L", target: "< 9", why: "Vascular + cognitive risk; responds to B-vitamin status." },
  { slug: "egfr", name: "eGFR", cat: "Kidney", unit: "mL/min", target: "> 90", why: "" },
  { slug: "cystatin_c", name: "Cystatin C", cat: "Kidney", unit: "mg/L", target: "low", why: "Muscle-independent kidney marker — better for athletes." },
  { slug: "alt", name: "ALT", cat: "Liver", unit: "IU/L", target: "< 25", why: "" },
  { slug: "vitamin_d", name: "Vitamin D (25-OH)", cat: "Vitamins", unit: "ng/mL", target: "40–60", why: "More is not better past the optimal band." },
  { slug: "vitamin_b12", name: "Vitamin B12", cat: "Vitamins", unit: "pg/mL", target: "400–800", why: "Very high values are usually supplement-driven; MMA confirms function." },
  { slug: "ebv_ea_igg", name: "EBV Early Antigen IgG", cat: "Immune", unit: "U/mL", target: "< 9", why: "Positive suggests reactivation; watch the trend." },
  { slug: "ana_screen", name: "ANA Screen", cat: "Immune", unit: "", target: "negative", why: "New positives are watch-items, especially with other immune signals." },
];
