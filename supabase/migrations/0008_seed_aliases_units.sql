-- ════════════════════════════════════════════════════════════════════════════
-- 0008 · Synonyms, unit conversions, sex/age reference ranges (re-runnable).
-- ════════════════════════════════════════════════════════════════════════════

-- ── Aliases: every label a lab/wearable might print -> canonical metric ──────
delete from metric_aliases;
insert into metric_aliases (metric_id, alias) values
('apob','ApoB'),('apob','Apolipoprotein B'),('apob','Apo B'),
('ldl_c','LDL'),('ldl_c','LDL-C'),('ldl_c','LDL Cholesterol'),('ldl_c','LDL Chol Calc'),('ldl_c','LDL-Cholesterol'),('ldl_c','Cholesterol, LDL'),
('hdl_c','HDL'),('hdl_c','HDL-C'),('hdl_c','HDL Cholesterol'),('hdl_c','Cholesterol, HDL'),
('triglycerides','Triglycerides'),('triglycerides','TG'),('triglycerides','Trig'),('triglycerides','Triglyceride'),
('lpa','Lp(a)'),('lpa','Lipoprotein (a)'),('lpa','Lipoprotein(a)'),('lpa','LPA'),('lpa','Lp(a) Mass'),
('total_cholesterol','Total Cholesterol'),('total_cholesterol','Cholesterol, Total'),('total_cholesterol','Cholesterol Total'),('total_cholesterol','TC'),
('non_hdl_c','Non-HDL Cholesterol'),('non_hdl_c','Non HDL Cholesterol'),('non_hdl_c','Non-HDL-C'),
('glucose_fasting','Glucose'),('glucose_fasting','Glucose, Fasting'),('glucose_fasting','Fasting Glucose'),('glucose_fasting','Glucose Fasting'),('glucose_fasting','FBG'),('glucose_fasting','Glucose, Serum'),
('hba1c','HbA1c'),('hba1c','A1c'),('hba1c','Hemoglobin A1c'),('hba1c','Hgb A1C'),('hba1c','GHb'),('hba1c','Glycohemoglobin'),('hba1c','Hemoglobin A1C (HbA1c)'),
('insulin_fasting','Insulin'),('insulin_fasting','Fasting Insulin'),('insulin_fasting','Insulin, Fasting'),
('homa_ir','HOMA-IR'),('homa_ir','HOMA IR'),
('hs_crp','hs-CRP'),('hs_crp','hsCRP'),('hs_crp','C-Reactive Protein, Cardiac'),('hs_crp','CRP, High Sensitivity'),('hs_crp','High Sensitivity CRP'),
('homocysteine','Homocysteine'),('homocysteine','Homocyst(e)ine'),
('uric_acid','Uric Acid'),('uric_acid','Urate'),
('weight','Weight'),('weight','Body Weight'),('weight','BodyMass'),
('bmi','BMI'),('bmi','Body Mass Index'),
('waist_cm','Waist'),('waist_cm','Waist Circumference'),('waist_cm','WaistCircumference'),
('body_fat_pct','Body Fat'),('body_fat_pct','Body Fat Percentage'),('body_fat_pct','BodyFatPercentage'),
('bp_systolic','Systolic'),('bp_systolic','Systolic Blood Pressure'),('bp_systolic','BloodPressureSystolic'),('bp_systolic','SBP'),
('bp_diastolic','Diastolic'),('bp_diastolic','Diastolic Blood Pressure'),('bp_diastolic','BloodPressureDiastolic'),('bp_diastolic','DBP'),
('resting_hr','Resting HR'),('resting_hr','Pulse'),('resting_hr','Heart Rate'),
('egfr','eGFR'),('egfr','GFR'),('egfr','eGFR (CKD-EPI)'),('egfr','Estimated GFR'),('egfr','GFR Estimated'),
('creatinine','Creatinine'),('creatinine','Creatinine, Serum'),('creatinine','Creat'),
('bun','BUN'),('bun','Urea Nitrogen'),('bun','Blood Urea Nitrogen'),
('uacr','UACR'),('uacr','Albumin/Creatinine Ratio'),('uacr','Microalbumin/Creatinine'),
('alt','ALT'),('alt','ALT (SGPT)'),('alt','SGPT'),('alt','Alanine Aminotransferase'),
('ast','AST'),('ast','AST (SGOT)'),('ast','SGOT'),('ast','Aspartate Aminotransferase'),
('ggt','GGT'),('ggt','Gamma GT'),('ggt','Gamma-Glutamyl Transferase'),
('alp','ALP'),('alp','Alkaline Phosphatase'),('alp','Alk Phos'),
('albumin','Albumin'),('albumin','Albumin, Serum'),
('bilirubin_total','Bilirubin'),('bilirubin_total','Total Bilirubin'),('bilirubin_total','Bilirubin, Total'),
('tsh','TSH'),('tsh','Thyroid Stimulating Hormone'),('tsh','Thyrotropin'),
('free_t4','Free T4'),('free_t4','FT4'),('free_t4','T4, Free'),('free_t4','Thyroxine, Free'),
('free_t3','Free T3'),('free_t3','FT3'),('free_t3','T3, Free'),('free_t3','Triiodothyronine, Free'),
('hemoglobin','Hemoglobin'),('hemoglobin','Hgb'),('hemoglobin','Hb'),
('hematocrit','Hematocrit'),('hematocrit','Hct'),
('wbc','WBC'),('wbc','White Blood Cell Count'),('wbc','White Blood Cells'),('wbc','Leukocytes'),
('platelets','Platelets'),('platelets','Platelet Count'),('platelets','PLT'),
('rdw','RDW'),('rdw','Red Cell Distribution Width'),
('mcv','MCV'),('mcv','Mean Corpuscular Volume'),
('vitamin_d','Vitamin D'),('vitamin_d','Vitamin D, 25-Hydroxy'),('vitamin_d','25-OH Vitamin D'),('vitamin_d','25-Hydroxyvitamin D'),('vitamin_d','Vit D 25 Hydroxy'),
('vitamin_b12','Vitamin B12'),('vitamin_b12','B12'),('vitamin_b12','Cobalamin'),
('folate','Folate'),('folate','Folic Acid'),('folate','Folate, Serum'),
('ferritin','Ferritin'),('ferritin','Ferritin, Serum'),
('magnesium','Magnesium'),('magnesium','Mg'),('magnesium','Magnesium, Serum'),
('omega3_index','Omega-3 Index'),('omega3_index','Omega 3 Index'),
('testosterone_total','Testosterone'),('testosterone_total','Total Testosterone'),('testosterone_total','Testosterone, Total'),
('free_testosterone','Free Testosterone'),('free_testosterone','Testosterone, Free'),
('psa','PSA'),('psa','Prostate Specific Antigen'),('psa','PSA, Total'),
('cortisol_am','Cortisol'),('cortisol_am','Cortisol, AM'),('cortisol_am','Cortisol A.M.'),
('igf1','IGF-1'),('igf1','IGF 1'),('igf1','Insulin-like Growth Factor 1'),('igf1','Somatomedin C'),
('shbg','SHBG'),('shbg','Sex Hormone Binding Globulin'),
('hrv_rmssd','HRV'),('hrv_rmssd','Heart Rate Variability'),('hrv_rmssd','HeartRateVariabilitySDNN'),('hrv_rmssd','RMSSD'),
('resting_hr_wearable','RestingHeartRate'),('resting_hr_wearable','Resting Heart Rate (wearable)'),
('vo2max','VO2max'),('vo2max','VO2 Max'),('vo2max','VO2Max'),('vo2max','Cardio Fitness'),
('sleep_duration','Sleep Duration'),('sleep_duration','SleepAnalysis'),('sleep_duration','Total Sleep'),('sleep_duration','Time Asleep'),
('sleep_efficiency','Sleep Efficiency'),('sleep_efficiency','SleepEfficiency'),
('steps','Steps'),('steps','StepCount'),('steps','Step Count'),('steps','Daily Steps'),
('active_energy','Active Energy'),('active_energy','ActiveEnergyBurned'),('active_energy','Active Calories'),
('respiratory_rate','Respiratory Rate'),('respiratory_rate','RespiratoryRate'),('respiratory_rate','Breathing Rate')
on conflict (alias_norm, metric_id) do nothing;

-- ── Unit conversions (canonicalize to each metric's canonical_unit) ──────────
delete from unit_conversions;
insert into unit_conversions (metric_id, from_unit, to_unit, factor) values
-- glucose mmol/L -> mg/dL
('glucose_fasting','mmol/L','mg/dL',18.0156),
-- cholesterols mmol/L -> mg/dL (x38.67)
('total_cholesterol','mmol/L','mg/dL',38.67),
('ldl_c','mmol/L','mg/dL',38.67),
('hdl_c','mmol/L','mg/dL',38.67),
('non_hdl_c','mmol/L','mg/dL',38.67),
-- triglycerides mmol/L -> mg/dL (x88.57)
('triglycerides','mmol/L','mg/dL',88.57),
-- creatinine umol/L -> mg/dL (x0.0113)
('creatinine','umol/L','mg/dL',0.0113),
-- vitamin D nmol/L -> ng/mL (/2.496 == x0.4006)
('vitamin_d','nmol/L','ng/mL',0.4006),
-- VO2max unit spellings (Apple/Garmin variants) -> canonical mL/kg/min
('vo2max','mL/min·kg','mL/kg/min',1),
('vo2max','mL/min/kg','mL/kg/min',1),
('vo2max','ml/kg/min','mL/kg/min',1),
-- insulin uU/mL spelling
('insulin_fasting','uU/mL','uIU/mL',1),
-- generic mass conversions (metric_id null = applies to any metric)
(null,'g','kg',0.001),
(null,'lb','kg',0.453592),
(null,'lbs','kg',0.453592),
(null,'cm','cm',1),
(null,'in','cm',2.54),
(null,'%','%',1)
on conflict (metric_id, from_unit, to_unit) do nothing;

-- ── Sex/age-specific reference ranges (resolved against profile) ────────────
delete from reference_ranges;
insert into reference_ranges (metric_id, sex, age_min, age_max, unit, low, high, source) values
('hdl_c','female',null,null,'mg/dL',50,null,'NCEP'),
('hdl_c','male',null,null,'mg/dL',40,null,'NCEP'),
('waist_cm','female',null,null,'cm',null,88,'IDF'),
('waist_cm','male',null,null,'cm',null,102,'IDF'),
('uric_acid','female',null,null,'mg/dL',2.4,6.0,'population'),
('uric_acid','male',null,null,'mg/dL',3.4,7.0,'population'),
('alt','female',null,null,'U/L',7,45,'population'),
('alt','male',null,null,'U/L',7,55,'population'),
('ggt','female',null,null,'U/L',5,38,'population'),
('ggt','male',null,null,'U/L',8,55,'population'),
('creatinine','female',null,null,'mg/dL',0.59,1.04,'population'),
('creatinine','male',null,null,'mg/dL',0.74,1.35,'population'),
('hemoglobin','female',null,null,'g/dL',12.0,15.5,'population'),
('hemoglobin','male',null,null,'g/dL',13.5,17.5,'population'),
('hematocrit','female',null,null,'%',36,44,'population'),
('hematocrit','male',null,null,'%',41,50,'population'),
('ferritin','female',null,null,'ng/mL',15,200,'population'),
('ferritin','male',null,null,'ng/mL',30,400,'population'),
('body_fat_pct','female',null,null,'%',21,33,'ACE'),
('body_fat_pct','male',null,null,'%',8,20,'ACE');
