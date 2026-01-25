# Comprehensive Analysis: Cardiac Diagnostic Exams for Risk Stratification

**Last Updated:** January 2026
**Purpose:** Evidence-based analysis of ECG and Echocardiography for EMR integration and risk stratification

---

## Table of Contents

1. [Executive Summary](#executive-summary)
2. [ECG 12-Lead Electrocardiogram](#ecg-12-lead-electrocardiogram)
3. [Transthoracic Echocardiography (TTE)](#transthoracic-echocardiography-tte)
4. [CSV and Database Integration](#csv-and-database-integration)
5. [Clinical Decision Support Algorithms](#clinical-decision-support-algorithms)
6. [Functional Medicine Integration](#functional-medicine-integration)
7. [Implementation Recommendations](#implementation-recommendations)
8. [References](#references)

---

## Executive Summary

### Key Findings

**ECG 12-Lead:**
- **CSV Suitability:** PARTIAL - Quantitative parameters (HR, QTc, PR, QRS duration) suitable; qualitative findings require categorical encoding
- **Critical Parameters:** Heart rate, QTc interval (sudden death risk), rhythm classification
- **Risk Stratification:** Strong evidence for QTc ≥450/460 ms (HR 1.72) and HR ≥90 bpm (HR 2.35) predicting cardiovascular death

**Echocardiography:**
- **CSV Suitability:** HIGH - Multiple quantitative parameters ideal for structured data
- **Critical Parameters:** Ejection fraction (EF%) is single most important; E/e' ratio and LAVI for diastolic function
- **Risk Stratification:** EF <40% defines HFrEF; E/e' >14 indicates elevated filling pressures; LAVI >34 mL/m² predicts adverse events

### Recommended Database Approach

1. **ECG:** Single CSV column for key quantitative parameters + separate table for detailed findings
2. **Echo:** Dedicated structured table with 15-20 key parameters OR single "EF%" column for screening
3. **Both:** Timestamp + physician overread flag + raw report storage (PDF/DICOM)

---

## ECG 12-Lead Electrocardiogram

### Overview

The 12-lead ECG records electrical activity of the heart from 12 different perspectives, providing comprehensive assessment of cardiac rhythm, conduction, ischemia, hypertrophy, and arrhythmias.

### 1. Quantitative Parameters (CSV-Suitable)

#### 1.1 Heart Rate (HR)

**Definition:** Ventricular rate in beats per minute (bpm)

**Normal Values:**
- Bradycardia: <60 bpm
- Normal: 60-100 bpm
- Tachycardia: >100 bpm

**Risk Stratification (2024-2025 Evidence):**
- **HR ≥90 bpm:** Hazard ratio 2.35 for cardiovascular death
- **Cancer therapy patients:** Higher baseline HR (74±13 vs 72±12 bpm) associated with cardiac dysfunction (HR 1.02 per bpm)
- **Continuous risk:** Linear increase in mortality risk above 75 bpm

**Clinical Significance:**
- Persistent tachycardia may indicate heart failure, hyperthyroidism, or autonomic dysfunction
- Severe bradycardia (<40 bpm) may require pacemaker evaluation

**CSV Format:** `heart_rate_bpm` (integer or decimal)

---

#### 1.2 QT Interval and QTc (Corrected)

**Definition:** Time from start of QRS complex to end of T wave, representing ventricular depolarization and repolarization

**Correction Formula:**
- **Bazett Formula:** QTc = QT / √RR interval
- Most commonly used, though has rate-dependency limitations
- Alternative formulas: Fridericia, Framingham (less rate-dependent)

**Normal Values:**
- **Men:** <450 ms (borderline 431-450 ms)
- **Women:** <460 ms (borderline 451-470 ms)

**Prolonged QTc Thresholds:**
- **Abnormal:** >450 ms (men), >470 ms (women)
- **High Risk:** >500 ms (significant sudden cardiac death risk)

**Risk Stratification (2024-2025 Evidence):**
- **QTc ≥450/460 ms:** Hazard ratio 1.72 for cardiovascular death
- Independent risk factor for sudden death due to cardiac arrest
- **Cancer therapy:** QTc was longer in patients who developed cardiac dysfunction (435±22 vs 428±21 ms; HR 1.01 per ms)

**Clinical Significance:**
- **Congenital Long QT Syndrome (LQTS):** 13 genetic subtypes, QTc typically >480 ms
- **Acquired QT Prolongation:** Medications (antiarrhythmics, antipsychotics, antibiotics), electrolyte abnormalities (hypokalemia, hypomagnesemia, hypocalcemia)
- **Torsades de Pointes Risk:** Polymorphic VT that can degenerate to VF

**Medication Screening:**
- Critical to screen all medications against QT-prolonging drug lists
- Examples: Azithromycin, methadone, antipsychotics (haloperidol, quetiapine), Class IA/III antiarrhythmics

**CSV Format:**
- `qt_interval_ms` (integer)
- `qtc_interval_ms` (integer, Bazett-corrected)
- `qtc_formula` (text: "Bazett", "Fridericia", "Framingham")

---

#### 1.3 PR Interval

**Definition:** Time from start of P wave to start of QRS complex, representing AV nodal conduction time

**Normal Values:** 120-200 ms

**Abnormal Values:**
- **Short PR (<120 ms):** Pre-excitation syndromes (WPW), enhanced AV nodal conduction
- **Prolonged PR (>200 ms):** First-degree AV block
- **Progressive prolongation:** Second-degree AV block Type I (Wenckebach)
- **Fixed PR with dropped beats:** Second-degree AV block Type II (high risk)

**Risk Stratification:**
- PR interval showed HR 1.01 per ms for cardiac dysfunction in cancer therapy patients
- Type II second-degree and complete (third-degree) AV block carry significant risk of progression to complete heart block

**CSV Format:** `pr_interval_ms` (integer)

---

#### 1.4 QRS Duration

**Definition:** Time from beginning to end of QRS complex, representing ventricular depolarization

**Normal Values:** <120 ms (80-100 ms typical)

**Abnormal Values:**
- **120-150 ms:** Incomplete bundle branch block (RBBB or LBBB)
- **≥120 ms:** Complete bundle branch block
- **≥150 ms:** Severe conduction delay (poor prognosis in heart failure)

**Bundle Branch Block Patterns:**

**Right Bundle Branch Block (RBBB):**
- QRS ≥120 ms
- RSR' pattern in V1-V2 ("M-shaped")
- Wide S wave in I, V5-V6
- **Clinical:** May be normal variant in young adults; associated with RV strain if acute

**Left Bundle Branch Block (LBBB):**
- QRS ≥120 ms
- Broad monophasic R wave in I, aVL, V5-V6
- Absent Q waves in I, V5-V6
- **Clinical:** Usually indicates underlying cardiac disease; new LBBB with chest pain = MI until proven otherwise (Sgarbossa criteria)

**Risk Stratification:**
- QRS ≥150 ms in heart failure: Increased mortality, candidate for CRT (Cardiac Resynchronization Therapy)
- New LBBB with acute MI: High-risk, requires urgent revascularization

**CSV Format:**
- `qrs_duration_ms` (integer)
- `bundle_branch_block` (text: "None", "RBBB", "LBBB", "IVCD")

---

### 2. Qualitative Findings (Categorical/Semi-Structured)

#### 2.1 Cardiac Rhythm

**Categories:**
- **Sinus Rhythm:** Normal P wave before each QRS, rate 60-100 bpm
- **Sinus Bradycardia:** Sinus rhythm <60 bpm
- **Sinus Tachycardia:** Sinus rhythm >100 bpm
- **Atrial Fibrillation (AFib):** Irregularly irregular rhythm, absent P waves, fibrillatory waves
- **Atrial Flutter:** Regular atrial rate 250-350 bpm, "sawtooth" pattern
- **Supraventricular Tachycardia (SVT):** Regular narrow-complex tachycardia >150 bpm
- **Ventricular Tachycardia (VT):** Wide-complex tachycardia ≥3 beats, rate >100 bpm
- **Ventricular Fibrillation (VF):** Chaotic ventricular activity, no organized QRS
- **Paced Rhythm:** Pacemaker spikes before QRS or P waves

**Risk Stratification:**
- **AFib:** 5x increased stroke risk, requires anticoagulation (CHA2DS2-VASc score)
- **VT/VF:** Life-threatening, requires immediate intervention and ICD evaluation

**CSV Format:** `rhythm` (categorical enum)

---

#### 2.2 Axis Deviation

**Normal Axis:** -30° to +90°

**Abnormal Axes:**
- **Left Axis Deviation (LAD):** -30° to -90° → Left anterior fascicular block, LVH, inferior MI
- **Right Axis Deviation (RAD):** +90° to +180° → RVH, lateral MI, PE, COPD
- **Extreme Axis Deviation:** -90° to ±180° → Ventricular tachycardia, hyperkalemia

**CSV Format:** `axis_degrees` (integer) + `axis_category` (text: "Normal", "LAD", "RAD", "Extreme")

---

#### 2.3 ST-T Wave Changes

**Ischemia Patterns:**

**ST Elevation (STEMI):**
- **Criteria:** ≥1 mm in 2+ contiguous leads (≥2 mm in V2-V3)
- **Locations:**
  - Anterior: V1-V4 (LAD territory)
  - Inferior: II, III, aVF (RCA or LCx)
  - Lateral: I, aVL, V5-V6 (LCx)
  - Posterior: Tall R waves V1-V2, ST depression V1-V3

**ST Depression (NSTEMI/Ischemia):**
- **Criteria:** ≥0.5 mm horizontal or downsloping in 2+ leads
- May indicate subendocardial ischemia or reciprocal changes

**T Wave Inversion:**
- **Deep T wave inversion:** Wellens syndrome (critical LAD stenosis), hypertrophic cardiomyopathy
- **Global T wave inversion:** Myocarditis, takotsubo cardiomyopathy

**Risk Stratification:**
- STEMI requires immediate catheterization (door-to-balloon <90 min)
- Wellens syndrome: 75% develop anterior MI within weeks without intervention

**CSV Format:**
- `st_elevation` (boolean + affected leads)
- `st_depression` (boolean + affected leads)
- `t_wave_inversion` (boolean + affected leads)
- `wellens_pattern` (boolean)

---

#### 2.4 Left Ventricular Hypertrophy (LVH)

**Sokolow-Lyon Criteria:**
- **Formula:** S wave in V1 + tallest R wave in V5 or V6
- **Positive:** >35 mm (3.5 mV)
- **Sensitivity:** 20%
- **Specificity:** >85%

**Cornell Voltage Criteria:**
- **Formula:** R wave in aVL + S wave in V3
- **Positive:** ≥28 mm (men), ≥20 mm (women)
- **Sensitivity:** 51%
- **Specificity:** 95%

**Voltage-Duration Product (More Accurate):**
- **(RaVL + SV3) × QRS duration**
- **Positive:** ≥2440 mm·ms (add 8 mm to voltage for women)
- Performs better than voltage alone, especially in borderline cases

**Additional LVH Indicators:**
- Left axis deviation
- Prolonged QRS duration
- ST-T changes in V5-V6 (strain pattern)
- Left atrial enlargement (P wave duration >120 ms in lead II)

**Clinical Significance:**
- Strong predictor of cardiovascular events and heart failure
- Regression with treatment (BP control) improves prognosis

**Note on Bundle Branch Block:**
- Voltage-based criteria perform poorly with LBBB
- QRS duration >120 ms itself suggests LVH when LBBB present
- Sokolow-Lyon ≥3.0 mV has best performance in LBBB patients

**CSV Format:**
- `lvh_sokolow_lyon_mm` (integer)
- `lvh_cornell_voltage_mm` (integer)
- `lvh_cornell_product_mm_ms` (integer)
- `lvh_present` (boolean)

---

### 3. Automated ECG Interpretation

**Current State (2024):**
- Most modern ECG machines provide automated computer interpretation
- Algorithms analyze waveforms and provide preliminary diagnosis

**Performance:**
- **Sensitivity:** 60-90% depending on condition
- **Specificity:** 70-95%
- **False Positives:** Common (artifact, baseline wander, muscle tremor)
- **Missed Findings:** Subtle ST changes, early ischemia

**Critical Requirement:**
- **Physician overread is MANDATORY**
- Automated interpretation should be considered preliminary
- Up to 20-30% of automated interpretations contain clinically significant errors

**Implementation in EMR:**
- Store automated interpretation separately from verified interpretation
- Flag for physician review
- Track physician overread completion
- Compare automated vs. physician interpretation for quality metrics

**CSV Format:**
- `automated_interpretation` (text, JSON blob)
- `physician_overread` (text)
- `physician_overread_date` (timestamp)
- `significant_discrepancy` (boolean)

---

### 4. Screening Guidelines (USPSTF 2018-2024)

#### 4.1 General CVD Risk Screening with ECG

**USPSTF Recommendation (2018):**
- **Grade I** (Insufficient Evidence) for screening asymptomatic low-risk adults
- Balance of benefits and harms cannot be determined

**Rationale:**
- Potential harms: False positives, unnecessary testing, anxiety
- ECG misinterpretation leading to inappropriate treatment
- Cost without proven benefit in asymptomatic population

#### 4.2 Atrial Fibrillation Screening

**USPSTF Recommendation (2022):**
- **Grade I** (Insufficient Evidence) for routine AFib screening in asymptomatic adults ≥50 years
- Applies to patients without AFib diagnosis, symptoms, or history of TIA/stroke

**Screening Methods:**
- Single ECG at clinic visit
- Intermittent monitoring with portable device (AliveCor, Apple Watch)
- Continuous monitoring (event recorders, patches)

**Potential Harms:**
- Misdiagnosis (artifact, premature beats)
- Inappropriate anticoagulation (bleeding risk)
- Overdiagnosis of paroxysmal AFib

#### 4.3 When ECG IS Indicated

Despite lack of screening recommendation, ECG is clinically indicated for:

**Symptoms:**
- Chest pain, palpitations, syncope, dyspnea
- Presyncope, lightheadedness
- Stroke/TIA evaluation

**High-Risk Populations:**
- Family history of sudden cardiac death
- Known or suspected heart disease
- Hypertension, diabetes, hyperlipidemia (moderate-high CVD risk)
- Pre-operative evaluation (intermediate-high risk surgery)

**Athletes:**
- AHA recommends history and physical (not routine ECG)
- European guidelines recommend ECG for competitive athletes
- Controversy continues

**Occupational:**
- Pilots, professional drivers, public safety personnel

---

### 5. ECG CSV Data Structure Recommendation

#### Option A: Minimal (Single Table)

```csv
patient_id, exam_date, heart_rate_bpm, qtc_ms, rhythm, ef_normal, physician_verified, report_url
```

**Pros:** Simple, sufficient for basic risk flagging
**Cons:** Loses detailed data for research and advanced algorithms

#### Option B: Comprehensive (Separate ECG Table)

```sql
CREATE TABLE ecg_results (
    id UUID PRIMARY KEY,
    patient_id UUID REFERENCES patients(id),
    exam_date TIMESTAMP NOT NULL,

    -- Quantitative parameters
    heart_rate_bpm INTEGER,
    pr_interval_ms INTEGER,
    qrs_duration_ms INTEGER,
    qt_interval_ms INTEGER,
    qtc_interval_ms INTEGER,
    qtc_formula VARCHAR(20), -- 'Bazett', 'Fridericia', etc.
    axis_degrees INTEGER,

    -- LVH parameters
    sokolow_lyon_mm INTEGER,
    cornell_voltage_mm INTEGER,
    cornell_product_mm_ms INTEGER,

    -- Categorical findings (enums)
    rhythm VARCHAR(50), -- 'sinus', 'afib', 'aflutter', 'vt', etc.
    axis_category VARCHAR(20), -- 'normal', 'lad', 'rad', 'extreme'
    bundle_branch_block VARCHAR(20), -- 'none', 'rbbb', 'lbbb', 'ivcd'
    av_block_degree VARCHAR(20), -- 'none', 'first', 'second_i', 'second_ii', 'third'

    -- Boolean flags
    st_elevation BOOLEAN,
    st_elevation_leads TEXT, -- JSON array: ["V2", "V3", "V4"]
    st_depression BOOLEAN,
    st_depression_leads TEXT,
    t_wave_inversion BOOLEAN,
    t_wave_inversion_leads TEXT,
    lvh_present BOOLEAN,
    wellens_pattern BOOLEAN,

    -- Interpretations
    automated_interpretation TEXT, -- JSON blob from ECG machine
    physician_interpretation TEXT,
    physician_overread_by UUID REFERENCES users(id),
    physician_overread_date TIMESTAMP,
    significant_discrepancy BOOLEAN,

    -- Raw data
    report_pdf_url TEXT,
    dicom_file_url TEXT,

    -- Metadata
    machine_model VARCHAR(100),
    technician_id UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Indexes for performance
CREATE INDEX idx_ecg_patient_date ON ecg_results(patient_id, exam_date DESC);
CREATE INDEX idx_ecg_qtc ON ecg_results(qtc_interval_ms) WHERE qtc_interval_ms >= 450;
CREATE INDEX idx_ecg_rhythm ON ecg_results(rhythm);
CREATE INDEX idx_ecg_unverified ON ecg_results(patient_id) WHERE physician_overread_date IS NULL;
```

**Pros:** Comprehensive, enables advanced queries, research-ready
**Cons:** More complex to implement and maintain

#### Recommended Approach

**Phase 1 (MVP):** Option A with key parameters only
**Phase 2:** Migrate to Option B as data needs grow
**Always include:** `heart_rate_bpm`, `qtc_ms`, `rhythm`, `physician_verified`, `report_url`

---

## Transthoracic Echocardiography (TTE)

### Overview

Transthoracic echocardiography is the most valuable initial diagnostic imaging tool for assessing cardiac structure and function. It provides real-time visualization of cardiac chambers, valves, and blood flow.

**Modalities:**
- **2D Echocardiography:** Real-time cross-sectional imaging
- **3D Echocardiography:** Volumetric assessment (more accurate for EF)
- **Doppler:** Blood flow velocity and direction
- **Tissue Doppler Imaging (TDI):** Myocardial tissue velocity
- **Speckle Tracking:** Strain and strain rate analysis

---

### 1. Left Ventricular Systolic Function

#### 1.1 Ejection Fraction (EF%)

**Definition:** Percentage of blood ejected from left ventricle with each contraction

**Formula:** EF% = [(LVEDV - LVESV) / LVEDV] × 100

**Measurement Methods:**

**1. Biplane Method of Disks (Modified Simpson's Rule) - GOLD STANDARD (2D)**
- Traces LV endocardium in apical 4-chamber and 2-chamber views
- Divides LV into multiple cylindrical disks
- Most accurate 2D method
- Normal EF: 63 ± 5%
- Normal range: 53-73% (in adults >20 years)

**2. 3D Echocardiography (Preferred When Available)**
- Most accurate and reproducible
- Less geometric assumptions
- Should be used when technically feasible

**3. Visual Estimation (Eyeball Method)**
- Experienced observers can estimate within ±5%
- Useful for rapid assessment
- Not recommended for quantitative reporting

**Classification (2024 Guidelines):**

| Category | EF Range | Clinical Significance |
|----------|----------|----------------------|
| **Normal** | ≥55% (52-72%) | Normal systolic function |
| **Mildly Reduced** | 45-54% | HFmrEF (Heart Failure with Mildly Reduced EF) |
| **Moderately Reduced** | 30-44% | HFrEF (Heart Failure with Reduced EF) |
| **Severely Reduced** | <30% | High mortality risk, ICD candidate |

**Heart Failure Phenotypes:**
- **HFrEF:** EF ≤40% - Beta-blockers, ACE-I/ARB/ARNI, MRAs, SGLT2i
- **HFmrEF:** EF 41-49% - Transitional category, similar treatment to HFrEF
- **HFpEF:** EF ≥50% - Diastolic dysfunction, treat comorbidities

**Risk Stratification:**
- **EF <35%:** Primary prevention ICD candidacy (if >40 days post-MI, >3 months on optimal medical therapy)
- **EF <20%:** Very high mortality risk, consider advanced heart failure therapies (LVAD, transplant)
- **Each 5% decrease in EF:** Incrementally increases mortality and hospitalization risk

**Prognostic Value:**
- Strongest single predictor of mortality in heart failure
- Serial measurements guide medication titration
- Required for device therapy eligibility (ICD, CRT)

**Limitations:**
- Load-dependent (affected by preload, afterload)
- May be normal in early heart failure with diastolic dysfunction
- Doesn't capture regional wall motion abnormalities
- Suboptimal in poor imaging windows (obesity, COPD)

**CSV Format:**
- `ef_percent` (decimal, 1 decimal place)
- `ef_method` (text: "biplane_simpson", "3d", "visual")
- `ef_category` (text: "normal", "mild_reduction", "moderate_reduction", "severe_reduction")

---

#### 1.2 Left Ventricular Volumes

**LV End-Diastolic Volume (LVEDV):**
- Volume at end of diastole (maximum filling)
- Normal: 62-150 mL (absolute), should be indexed to BSA
- Indexed: Men 34-74 mL/m², Women 29-61 mL/m²

**LV End-Systolic Volume (LVESV):**
- Volume at end of systole (after contraction)
- Normal: 16-60 mL (absolute)
- Indexed: Men 11-31 mL/m², Women 8-24 mL/m²

**Clinical Significance:**
- **Elevated LVEDV:** Dilated cardiomyopathy, volume overload (MR, AR)
- **Elevated LVESV:** Impaired systolic function
- **LV remodeling:** Progressive increase in volumes post-MI indicates adverse remodeling

**Body Surface Area (BSA) Indexing:**
- Essential for comparing across different body sizes
- Mosteller formula: BSA (m²) = √[(height(cm) × weight(kg)) / 3600]

**CSV Format:**
- `lvedv_ml` (integer)
- `lvesv_ml` (integer)
- `lvedv_indexed` (decimal, mL/m²)
- `lvesv_indexed` (decimal, mL/m²)

---

#### 1.3 Global Longitudinal Strain (GLS)

**Definition:** Percentage change in myocardial fiber length during systole using speckle-tracking echocardiography

**Normal Values:** -18% and lower (more negative), e.g., -20%, -22%

**Abnormal Values:** -16% or higher (less negative), e.g., -14%, -12%

**Measurement:**
- Speckle-tracking echocardiography analyzes frame-to-frame motion of "acoustic markers"
- Average of 17 myocardial segments
- More sensitive than EF for detecting subclinical dysfunction

**Clinical Applications (2024-2025 Evidence):**

**1. Subclinical Dysfunction:**
- Detects early cardiomyopathy before EF reduction
- Cancer therapy monitoring: ≥15% relative decrease in GLS = cardiotoxicity

**2. Risk Stratification:**
- Lower GLS associated with higher cardiovascular events (HR 1.12 per 1% decrease)
- Provides independent prognostic information beyond EF
- Predicts long-term mortality in general population

**3. Specific Conditions:**
- **Hypertrophic cardiomyopathy:** Standardized reference values improve prognosis prediction
- **Arrhythmogenic cardiomyopathy:** RV GLS -13±6% (with events) vs -23±6% (without)
- **Systemic sclerosis:** Predicts death and CV events

**4. Serial Monitoring:**
- More reproducible than EF for tracking changes over time
- Less load-dependent than EF
- Superior reliability for assessing LV function evolution

**Advantages Over EF:**
- Detects regional dysfunction
- Earlier marker of disease
- Superior inter-observer reproducibility
- Disease-specific diagnostic insights

**Limitations:**
- Requires adequate image quality
- Vendor-dependent (different software algorithms)
- More time-consuming than EF
- Not universally available

**CSV Format:**
- `gls_percent` (decimal, typically negative value)
- `gls_category` (text: "normal", "borderline", "abnormal")
- `gls_vendor` (text: GE, Philips, Siemens, etc.)

---

#### 1.4 LV Mass and Hypertrophy

**LV Mass Calculation:**
- Derived from LV dimensions and wall thickness
- Devereux formula (most common)

**LV Mass Index (Indexed to BSA):**
- **Normal:** Men <115 g/m², Women <95 g/m²
- **LVH:** Above these thresholds

**Geometric Patterns:**

| Pattern | RWT | LV Mass Index | Characteristics |
|---------|-----|--------------|----------------|
| **Normal** | <0.42 | Normal | Normal geometry |
| **Concentric Remodeling** | ≥0.42 | Normal | Increased wall thickness, normal mass |
| **Eccentric Hypertrophy** | <0.42 | Increased | Dilated LV, proportional wall thickness |
| **Concentric Hypertrophy** | ≥0.42 | Increased | Thick walls, small cavity |

**Relative Wall Thickness (RWT):**
- RWT = (2 × posterior wall thickness) / LV end-diastolic dimension
- Normal: <0.42

**Clinical Significance:**
- **Concentric LVH:** Hypertension, aortic stenosis; highest CVD risk
- **Eccentric LVH:** Volume overload (MR, AR); HFrEF pattern
- **Concentric Remodeling:** Early hypertensive changes

**Prognosis:**
- Concentric LVH carries highest mortality risk
- LVH regression with treatment improves outcomes

**CSV Format:**
- `lv_mass_g` (integer)
- `lv_mass_index_g_m2` (decimal)
- `rwt` (decimal, 2 decimal places)
- `lv_geometry` (text: "normal", "concentric_remodeling", "eccentric_lvh", "concentric_lvh")

---

### 2. Left Ventricular Diastolic Function

**Context:** Diastolic dysfunction is an early marker of heart failure, often present before systolic dysfunction. Assessment is integral for dyspnea evaluation and HFpEF diagnosis.

#### 2.1 Mitral Inflow Velocities (Pulsed-Wave Doppler)

**E Wave (Early Diastolic Filling):**
- Peak velocity of early LV filling
- Normal: 0.6-1.3 m/s (age-dependent, decreases with age)

**A Wave (Atrial Contraction):**
- Peak velocity of filling due to atrial kick
- Normal: 0.4-0.9 m/s

**E/A Ratio:**

| E/A Ratio | Age | Interpretation |
|-----------|-----|----------------|
| <0.8 | Any | Impaired relaxation (Grade I diastolic dysfunction) |
| 0.8-2.0 | <50 years | Normal |
| 0.8-2.0 | >50 years | May be pseudonormal (Grade II) or normal - requires additional parameters |
| >2.0 | Any | Restrictive pattern (Grade III) or normal in young athletes |

**Deceleration Time (DT):**
- Time for E wave velocity to decrease to baseline
- Normal: 160-240 ms
- Prolonged (>240 ms): Impaired relaxation
- Shortened (<160 ms): Restrictive physiology, high LA pressure

**Clinical Significance:**
- E/A ratio alone insufficient due to "pseudonormalization"
- Must integrate with other parameters (see diastolic function algorithm)

**CSV Format:**
- `mitral_e_velocity_ms` (decimal, m/s)
- `mitral_a_velocity_ms` (decimal, m/s)
- `e_a_ratio` (decimal, 2 decimals)
- `deceleration_time_ms` (integer)

---

#### 2.2 Tissue Doppler Imaging (TDI) - e' Velocities

**e' Wave:** Early diastolic mitral annular velocity (myocardial relaxation)

**Measurement Sites:**
- **Septal e':** Medial mitral annulus
- **Lateral e':** Lateral mitral annulus
- **Average e':** (Septal + Lateral) / 2

**Normal Values (2025 ASE Guidelines):**
- **Septal e':** >7 cm/s
- **Lateral e':** >10 cm/s
- **Average e':** >9 cm/s

**Abnormal Values:**
- **Septal e' <7 cm/s** or **Lateral e' <10 cm/s:** Impaired relaxation

**Clinical Significance:**
- Less load-dependent than E wave
- Reflects intrinsic myocardial relaxation
- Decreases with age and cardiac disease

**CSV Format:**
- `tdi_septal_e_prime_cms` (decimal, cm/s)
- `tdi_lateral_e_prime_cms` (decimal, cm/s)
- `tdi_average_e_prime_cms` (decimal, cm/s)

---

#### 2.3 E/e' Ratio - Filling Pressure Estimate

**Definition:** Ratio of early mitral inflow velocity (E) to early diastolic mitral annular velocity (e')

**Calculation:** E/e' = (Mitral E velocity in m/s × 100) / Average e' velocity in cm/s

**Interpretation (2024-2025 Guidelines):**

| E/e' Ratio | LA Pressure | Clinical Significance |
|------------|-------------|----------------------|
| <8 | Normal | Normal LV filling pressure |
| 8-14 | Intermediate | Indeterminate - requires additional parameters |
| >14 | Elevated | Elevated LV filling pressures |

**Updated 2025 ASE Guidelines:**
- **Average E/e' >14:** Marker of elevated LV filling pressures
- Near-linear association with heart failure incidence or death
- No clear threshold with respect to prognosis (continuous risk)

**Clinical Applications:**
- **E/e' >14:** Strong predictor of elevated PCWP (pulmonary capillary wedge pressure)
- Correlates with NT-proBNP elevation
- Guides diuretic therapy in heart failure

**Limitations:**
- Intermediate range (8-14) requires additional assessment
- May be less reliable in:
  - Mitral annular calcification
  - Mitral valve disease (stenosis/regurgitation)
  - Paced rhythm
  - Atrial fibrillation

**CSV Format:**
- `e_e_prime_ratio` (decimal, 2 decimals)
- `filling_pressure_category` (text: "normal", "intermediate", "elevated")

---

#### 2.4 Left Atrial Volume Index (LAVI)

**Definition:** Left atrial volume indexed to body surface area

**Measurement:** Biplane area-length method (apical 4-chamber and 2-chamber views)

**Normal Values:** <34 mL/m²

**Abnormal Values:**
- **Mildly increased:** 34-38 mL/m²
- **Moderately increased:** 39-42 mL/m²
- **Severely increased:** >42 mL/m²

**Clinical Significance (2024-2025 Evidence):**

**1. Chronic Pressure/Volume Overload:**
- Reflects cumulative exposure to elevated LV filling pressures
- "Barometer" of diastolic burden over time

**2. Heart Failure Risk:**
- Near-linear association with incidence of HF or death
- Independent predictor of HF hospitalization
- LAVI >34 mL/m² predicts poor prognosis in HF patients

**3. Atrial Fibrillation Risk:**
- LA enlargement is strongest predictor of AFib development
- Each 5 mL/m² increase = 40% higher AFib risk

**4. Stroke Risk:**
- Independent predictor even in sinus rhythm
- Contributes to thromboembolism risk

**5. Hypertension:**
- Independent predictor of hypertensive response to exercise
- Marker of target organ damage

**Limitations (2024-2025 Updates):**
- **Inconsistency with other indices:** Often discordant with other LA pressure markers
- **Weak correlation with LAP:** Not strong correlation with directly measured LA pressure
- **Fails to track changes:** LA volumes don't reliably follow acute changes in LA pressure
- **Better for chronic assessment:** Reflects long-term remodeling rather than acute hemodynamics

**Recommendation:**
- LAVI remains in diastolic assessment algorithms but with reduced weight
- Left atrial strain (LARS) emerging as superior marker (see below)

**CSV Format:**
- `la_volume_ml` (integer)
- `la_volume_index_ml_m2` (decimal, 1 decimal)
- `la_size_category` (text: "normal", "mild", "moderate", "severe")

---

#### 2.5 Left Atrial Reservoir Strain (LARS) - Emerging Gold Standard

**Definition:** Peak longitudinal strain of left atrium during reservoir phase (ventricular systole)

**Measurement:** Speckle-tracking strain analysis of LA in apical 4-chamber view

**Normal Values:** >24%

**Thresholds (2024-2025 Guidelines):**
- **LARS ≤18%:** Elevated LA pressure (high specificity)
- **LARS ≤24%:** Risk stratification cutpoint for adverse events

**Clinical Applications (2025 Evidence):**

**1. Superior to LAVI for Filling Pressures:**
- More accurate than LAVI for assessing elevated LVFP
- LARS emerged as useful predictor of elevated LVFP
- Better correlation with invasive measurements

**2. Risk Stratification in HFpEF:**
- **High Risk:** Diastolic dysfunction + elevated filling pressure OR indeterminate diastolic function + LARS ≤24%
- **Lower Risk:** Other patterns
- Patients at high risk showed significant event rate in follow-up

**3. Incremental Prognostic Value:**
- Improves risk stratification beyond traditional diastolic parameters
- Predicts cardiovascular events in preserved EF
- Association with combined events: HF, AFib, ischemic stroke

**4. Early Detection:**
- Detects LA dysfunction before structural enlargement
- More sensitive than LAVI for early diastolic disease

**Advantages:**
- Less affected by acute loading conditions than E/e'
- Detects LA dysfunction before volume increase
- Strong prognostic value
- Incorporates LA function, not just size

**Limitations:**
- Requires adequate image quality
- Not yet universally available
- Vendor-specific analysis software
- Learning curve for acquisition and analysis

**Implementation Status:**
- Included in 2025 ASE diastolic function guidelines
- Recommended when available and technically feasible
- Expected to become standard of care within 2-3 years

**CSV Format:**
- `lars_percent` (decimal, 1 decimal)
- `lars_category` (text: "normal", "borderline", "abnormal")

---

#### 2.6 Diastolic Function Grading (2024-2025 Algorithm)

**2025 ASE Guidelines Update:** Refined approach incorporating LA strain

**Grade 0 - Normal:**
- Average E/e' <8 (or 9-13 with normal LARS >24%)
- LAVI <34 mL/m²
- No evidence of elevated filling pressure

**Grade I - Impaired Relaxation:**
- E/A <0.8
- Average E/e' <8
- Normal filling pressures at rest

**Grade II - Pseudonormal (Elevated Filling Pressures):**
- E/A 0.8-2.0
- Average E/e' >14 OR
- LAVI >34 mL/m² OR
- TR velocity >2.8 m/s OR
- LARS ≤18%

**Grade III - Restrictive:**
- E/A >2.0
- DT <160 ms
- Average E/e' >14
- Severely elevated filling pressures

**Indeterminate:**
- E/A 0.8-2.0
- Average E/e' 9-14
- Mixed findings on other parameters
- Requires additional testing (exercise echo, invasive hemodynamics)

**Clinical Correlation:**
- Must integrate with symptoms, BNP/NT-proBNP, other imaging
- Atrial fibrillation complicates assessment (no A wave)

**CSV Format:**
- `diastolic_function_grade` (text: "normal", "grade_i", "grade_ii", "grade_iii", "indeterminate")
- `elevated_filling_pressure` (boolean)

---

#### 2.7 Tricuspid Regurgitation (TR) Velocity

**Definition:** Peak velocity of tricuspid regurgitation jet by continuous-wave Doppler

**Normal:** <2.8 m/s

**Abnormal:** ≥2.8 m/s

**Clinical Significance:**
- Marker of elevated right atrial pressure
- Included in diastolic function algorithm (2025 ASE)
- Contributes to assessment of elevated LV filling pressures

**Limitation:**
- Requires presence of TR jet
- Not measurable in all patients (~30% lack sufficient TR)

**CSV Format:**
- `tr_velocity_ms` (decimal, m/s)
- `tr_velocity_elevated` (boolean)

---

### 3. Valvular Heart Disease

#### 3.1 Aortic Stenosis (AS)

**Primary Parameters:**

| Severity | Peak Velocity (Vmax) | Mean Gradient | Aortic Valve Area (AVA) |
|----------|---------------------|---------------|------------------------|
| **Normal** | <2.0 m/s | <10 mmHg | >2.0 cm² |
| **Mild** | 2.0-2.9 m/s | <20 mmHg | 1.5-2.0 cm² |
| **Moderate** | 3.0-3.9 m/s | 20-39 mmHg | 1.0-1.5 cm² |
| **Severe** | ≥4.0 m/s | ≥40 mmHg | ≤1.0 cm² |

**Indexed AVA:**
- AVAi ≤0.6 cm²/m² = Severe AS (alternative to absolute AVA)

**Discordant Grading:**
- **Low-Flow, Low-Gradient AS:** AVA ≤1.0 cm² but mean gradient 20-40 mmHg
- Occurs in 20-40% of patients
- Defined by stroke volume index <35 mL/m²
- **Dobutamine stress echo** to distinguish:
  - True severe AS: Gradient increases, AVA remains ≤1.0 cm²
  - Pseudo-severe AS: AVA increases >1.0 cm², not truly severe
- **Calcium score** verification (CT):
  - Severe AS: ≥2,000 Agatston units (men), ≥1,250 AU (women)

**Inconsistencies in Criteria:**
- Gradient of 40 mmHg correlates with AVA ~0.8 cm² (not 1.0 cm²)
- AVA 1.0 cm² relates to mean gradient ~26 mmHg (not 40 mmHg)
- Must consider all parameters and clinical context

**Risk Stratification:**
- **Symptomatic severe AS:** Very high mortality without intervention (2-year survival ~50%)
- **Asymptomatic severe AS with Vmax >5.0 m/s:** High risk, consider early intervention
- **Very severe AS:** Vmax ≥5.0 m/s, mean gradient ≥60 mmHg → expedited surgery

**CSV Format:**
- `as_peak_velocity_ms` (decimal, m/s)
- `as_mean_gradient_mmhg` (integer)
- `as_ava_cm2` (decimal, 2 decimals)
- `as_severity` (text: "none", "mild", "moderate", "severe")
- `as_low_flow` (boolean, SVI <35)

---

#### 3.2 Mitral Regurgitation (MR)

**Severity Grading (Semi-Quantitative):**

| Severity | EROA (cm²) | Regurgitant Volume (mL) | LA Size | LV Size |
|----------|-----------|-------------------------|---------|---------|
| **Mild** | <0.2 | <30 | Normal | Normal |
| **Moderate** | 0.2-0.4 | 30-59 | Mild enlargement | Mild enlargement |
| **Severe** | ≥0.4 | ≥60 | Moderate-severe enlargement | Dilated |

**EROA:** Effective Regurgitant Orifice Area (by PISA method)

**Primary vs. Secondary MR:**
- **Primary (Degenerative):** Leaflet pathology (prolapse, flail, perforation)
- **Secondary (Functional):** Structurally normal leaflets, LV/annular dilatation

**Risk Stratification:**
- **Severe symptomatic MR:** High mortality, requires intervention
- **Severe asymptomatic primary MR:** Consider surgery if LVESD >40 mm or EF <60%
- **Severe secondary MR:** Medical optimization of HF first; surgery controversial

**CSV Format:**
- `mr_eroa_cm2` (decimal, 2 decimals)
- `mr_regurgitant_volume_ml` (integer)
- `mr_severity` (text: "none", "trace", "mild", "moderate", "severe")
- `mr_type` (text: "primary", "secondary")

---

#### 3.3 Aortic Regurgitation (AR)

**Severity Grading:**

| Severity | Jet Width/LVOT | EROA (cm²) | Regurgitant Volume (mL) | Vena Contracta (cm) |
|----------|---------------|-----------|-------------------------|-------------------|
| **Mild** | <25% | <0.10 | <30 | <0.3 |
| **Moderate** | 25-64% | 0.10-0.29 | 30-59 | 0.3-0.6 |
| **Severe** | ≥65% | ≥0.30 | ≥60 | >0.6 |

**Additional Findings in Severe AR:**
- LV dilation
- Holodiastolic flow reversal in descending aorta
- Pressure half-time <200 ms (rapid equilibration = severe AR)

**Risk Stratification:**
- **Symptomatic severe AR:** Requires surgery
- **Asymptomatic severe AR:** Surgery if LVESD >50 mm or EF <50%

**CSV Format:**
- `ar_eroa_cm2` (decimal, 2 decimals)
- `ar_severity` (text: "none", "trace", "mild", "moderate", "severe")

---

### 4. Right Ventricular Function and Pulmonary Hypertension

#### 4.1 Right Ventricular Systolic Pressure (RVSP)

**Definition:** Estimated RV systolic pressure derived from TR velocity using simplified Bernoulli equation

**Calculation:** RVSP = 4 × (TR velocity)² + RAP

**Right Atrial Pressure (RAP) Estimation:**
- Based on IVC diameter and collapsibility
- **RAP 3 mmHg:** IVC <2.1 cm, collapses >50%
- **RAP 8 mmHg:** IVC >2.1 cm, collapses >50% OR IVC <2.1 cm, collapses <50%
- **RAP 15 mmHg:** IVC >2.1 cm, collapses <50%

**Normal Values:**
- **Healthy adults:** Mean PAP 21 ± 4 mmHg (by right heart catheterization)
- **RVSP normal:** <30 mmHg
- **Abnormal:** ≥30 mmHg

**Pulmonary Hypertension Thresholds (2024 Guidelines):**

**Updated Definition:**
- **PH redefined:** Mean PAP >20 mmHg (lowered from previous 25 mmHg)
- **Systolic PAP 30 mmHg** typically implies mean PAP >20 mmHg

**Echocardiographic Categories:**

| Category | RVSP | Clinical Significance |
|----------|------|----------------------|
| **Normal** | <30 mmHg | Normal pulmonary pressure |
| **Borderline/Mild** | 30-39 mmHg | Increased risk begins at 27 mmHg |
| **Mild PH** | 40-49 mmHg | ACC/AHA: Further evaluation if dyspnea |
| **Moderate PH** | 50-59 mmHg | Significant PH |
| **Severe PH** | ≥60 mmHg | Severe PH |

**Risk Stratification (2024 Evidence):**
- **Increased risk begins at 27 mmHg** (well within "normal range")
- **Risk doubles at 35 mmHg** vs. reference group
- **RVSP >40 mmHg with dyspnea:** Requires comprehensive PH evaluation (right heart catheterization)

**Etiology of PH (WHO Classification):**
1. **Group 1:** Pulmonary arterial hypertension (PAH)
2. **Group 2:** Left heart disease (most common, ~70%)
3. **Group 3:** Lung disease/hypoxia
4. **Group 4:** Chronic thromboembolic PH (CTEPH)
5. **Group 5:** Unclear/multifactorial

**Limitations:**
- Requires adequate TR jet (absent in ~30%)
- Overestimates or underestimates in ~20-30%
- RAP estimation adds variability
- Gold standard: Right heart catheterization

**Clinical Action:**
- RVSP >40 mmHg → Comprehensive evaluation (ECG, PFTs, VQ scan, right heart cath)
- Assess RV function (TAPSE, FAC, S' wave)
- Search for underlying cause

**CSV Format:**
- `tr_velocity_ms` (decimal, m/s)
- `rap_estimate_mmhg` (integer: 3, 8, or 15)
- `rvsp_mmhg` (integer)
- `ph_severity` (text: "none", "borderline", "mild", "moderate", "severe")

---

#### 4.2 RV Systolic Function Parameters

**TAPSE (Tricuspid Annular Plane Systolic Excursion):**
- M-mode measurement of RV free wall longitudinal motion
- **Normal:** ≥17 mm
- **Abnormal:** <17 mm (RV systolic dysfunction)

**RV Fractional Area Change (FAC):**
- (RV end-diastolic area - RV end-systolic area) / RV end-diastolic area × 100
- **Normal:** ≥35%
- **Abnormal:** <35%

**RV S' Wave (Tissue Doppler):**
- Peak systolic velocity of tricuspid lateral annulus
- **Normal:** ≥9.5 cm/s
- **Abnormal:** <9.5 cm/s

**Clinical Significance:**
- RV dysfunction predicts poor prognosis in HF, PH, PE
- Serial measurements guide therapy in PH

**CSV Format:**
- `tapse_mm` (decimal, 1 decimal)
- `rv_fac_percent` (integer)
- `rv_s_prime_cms` (decimal, 1 decimal)
- `rv_dysfunction` (boolean)

---

### 5. Additional Echo Parameters

#### 5.1 Pericardial Effusion

**Size Categories:**
- **Trace/Trivial:** Echo-free space in systole only
- **Small:** <10 mm in diastole
- **Moderate:** 10-20 mm in diastole
- **Large:** >20 mm in diastole

**Tamponade Physiology:**
- RA collapse in systole
- RV collapse in diastole
- Respiratory variation in mitral/tricuspid inflow (>25%)
- IVC plethora (dilated, non-collapsing)

**CSV Format:**
- `pericardial_effusion_size` (text: "none", "trace", "small", "moderate", "large")
- `tamponade_physiology` (boolean)

---

#### 5.2 Wall Motion Abnormalities

**Regional Wall Motion:**
- 17-segment model (ASE/EACVI)
- Each segment scored:
  - 1 = Normal
  - 2 = Hypokinetic
  - 3 = Akinetic
  - 4 = Dyskinetic (paradoxical motion)

**Wall Motion Score Index (WMSI):**
- Sum of scores / Number of segments
- **Normal:** 1.0
- **Abnormal:** >1.0

**Clinical Significance:**
- Identifies coronary artery territory ischemia/infarction
- LAD → Anterior/septal/apex
- LCx → Lateral/posterior
- RCA → Inferior/inferoseptal

**CSV Format:**
- `wall_motion_abnormality` (boolean)
- `wmsi` (decimal, 2 decimals)
- `affected_territories` (text, JSON array)

---

## CSV and Database Integration

### Recommended Strategy: Hybrid Approach

#### 1. ECG Data Structure

**Minimal Screening Approach:**
```csv
patient_id, exam_date, heart_rate_bpm, qtc_ms, rhythm, critical_flag, report_url
```

**Comprehensive Approach (Separate Table):**
See detailed SQL schema in ECG section above.

**Critical Flags for Automated Alerts:**
- QTc ≥500 ms
- HR <40 or >120 bpm
- New AFib
- ST elevation
- High-grade AV block

---

#### 2. Echo Data Structure

**Option A: Single Parameter (Minimal)**

For **screening** purposes, EF alone may suffice:

```csv
patient_id, exam_date, ef_percent, ef_category, diastolic_dysfunction, report_url
```

**Pros:** Simple, captures most important parameter
**Cons:** Loses rich data for comprehensive risk stratification

**Recommendation:** Use for MVP or when TTE is infrequent

---

**Option B: Core Parameters (Recommended)**

Balanced approach with key risk stratification parameters:

```csv
patient_id, exam_date, ef_percent, lavi_ml_m2, e_e_prime_ratio, gls_percent, rvsp_mmhg, as_severity, mr_severity, report_url
```

**Included:**
- **EF:** Systolic function
- **LAVI:** LA remodeling, AFib risk
- **E/e' ratio:** Filling pressures
- **GLS:** Subclinical dysfunction (if available)
- **RVSP:** Pulmonary hypertension screening
- **AS/MR severity:** Major valvular disease

**Pros:** Comprehensive risk stratification, manageable complexity
**Cons:** More fields to maintain

**Recommendation:** Best for robust EMR with active cardiology practice

---

**Option C: Full Structured Table (Research-Grade)**

Complete echocardiographic dataset:

```sql
CREATE TABLE echo_results (
    id UUID PRIMARY KEY,
    patient_id UUID REFERENCES patients(id),
    exam_date TIMESTAMP NOT NULL,

    -- LV Systolic Function
    ef_percent DECIMAL(4,1),
    ef_method VARCHAR(30), -- 'biplane_simpson', '3d', 'visual'
    lvedv_ml INTEGER,
    lvesv_ml INTEGER,
    lvedv_indexed DECIMAL(5,1), -- mL/m²
    lvesv_indexed DECIMAL(5,1),
    gls_percent DECIMAL(4,1), -- Usually negative, e.g., -18.5

    -- LV Geometry
    lv_mass_g INTEGER,
    lv_mass_index_g_m2 DECIMAL(5,1),
    rwt DECIMAL(3,2),
    lv_geometry VARCHAR(30), -- 'normal', 'concentric_remodeling', etc.

    -- Diastolic Function
    mitral_e_velocity_ms DECIMAL(3,2),
    mitral_a_velocity_ms DECIMAL(3,2),
    e_a_ratio DECIMAL(3,2),
    deceleration_time_ms INTEGER,
    tdi_septal_e_prime_cms DECIMAL(3,1),
    tdi_lateral_e_prime_cms DECIMAL(3,1),
    tdi_average_e_prime_cms DECIMAL(3,1),
    e_e_prime_ratio DECIMAL(4,1),
    la_volume_ml INTEGER,
    la_volume_index_ml_m2 DECIMAL(4,1),
    lars_percent DECIMAL(4,1), -- Left atrial reservoir strain
    tr_velocity_ms DECIMAL(3,2),
    diastolic_function_grade VARCHAR(20), -- 'normal', 'grade_i', 'grade_ii', 'grade_iii'
    elevated_filling_pressure BOOLEAN,

    -- Valvular Disease
    as_peak_velocity_ms DECIMAL(3,2),
    as_mean_gradient_mmhg INTEGER,
    as_ava_cm2 DECIMAL(3,2),
    as_severity VARCHAR(20), -- 'none', 'mild', 'moderate', 'severe'
    mr_severity VARCHAR(20),
    mr_type VARCHAR(20), -- 'primary', 'secondary'
    ar_severity VARCHAR(20),
    tr_severity VARCHAR(20),

    -- RV Function & PH
    rvsp_mmhg INTEGER,
    ph_severity VARCHAR(20),
    tapse_mm DECIMAL(3,1),
    rv_fac_percent INTEGER,
    rv_s_prime_cms DECIMAL(3,1),
    rv_dysfunction BOOLEAN,

    -- Other
    pericardial_effusion_size VARCHAR(20),
    tamponade_physiology BOOLEAN,
    wall_motion_abnormality BOOLEAN,
    wmsi DECIMAL(3,2),

    -- Report
    report_pdf_url TEXT,
    dicom_study_url TEXT,

    -- Quality & Verification
    image_quality VARCHAR(20), -- 'excellent', 'good', 'adequate', 'poor'
    sonographer_id UUID REFERENCES users(id),
    interpreting_cardiologist_id UUID REFERENCES users(id),
    interpretation_date TIMESTAMP,

    -- Metadata
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Indexes
CREATE INDEX idx_echo_patient_date ON echo_results(patient_id, exam_date DESC);
CREATE INDEX idx_echo_ef ON echo_results(ef_percent) WHERE ef_percent < 55;
CREATE INDEX idx_echo_rvsp ON echo_results(rvsp_mmhg) WHERE rvsp_mmhg >= 40;
CREATE INDEX idx_echo_uninterpreted ON echo_results(patient_id)
    WHERE interpreting_cardiologist_id IS NULL;
```

**Pros:** Complete dataset, research-ready, comprehensive risk algorithms
**Cons:** Complex, requires disciplined data entry

**Recommendation:** Use for academic centers or large cardiology practices

---

### 3. Risk Stratification Queries

**Example: High-Risk Cardiac Patients**

```sql
-- Patients with EF <40% OR severe AS OR severe PH
SELECT
    p.id,
    p.name,
    e.exam_date,
    e.ef_percent,
    e.as_severity,
    e.rvsp_mmhg,
    e.diastolic_function_grade
FROM patients p
JOIN echo_results e ON p.id = e.patient_id
WHERE
    e.ef_percent < 40
    OR e.as_severity = 'severe'
    OR e.rvsp_mmhg >= 60
    OR (e.ef_percent >= 50 AND e.elevated_filling_pressure = TRUE) -- HFpEF
ORDER BY e.exam_date DESC;
```

**Example: QTc Prolongation + Diastolic Dysfunction (Medication Caution)**

```sql
-- Patients with prolonged QTc and HFpEF (high risk for arrhythmia with QT drugs)
SELECT
    p.id,
    p.name,
    ecg.qtc_interval_ms,
    echo.ef_percent,
    echo.e_e_prime_ratio,
    echo.diastolic_function_grade
FROM patients p
LEFT JOIN (
    SELECT DISTINCT ON (patient_id) *
    FROM ecg_results
    ORDER BY patient_id, exam_date DESC
) ecg ON p.id = ecg.patient_id
LEFT JOIN (
    SELECT DISTINCT ON (patient_id) *
    FROM echo_results
    ORDER BY patient_id, exam_date DESC
) echo ON p.id = echo.patient_id
WHERE
    ecg.qtc_interval_ms >= 450
    AND echo.ef_percent >= 50
    AND echo.elevated_filling_pressure = TRUE
ORDER BY ecg.qtc_interval_ms DESC;
```

---

## Clinical Decision Support Algorithms

### Algorithm 1: Heart Failure Risk Stratification

```
INPUT: Latest Echo + ECG + Clinical Data

STEP 1: Systolic Function Assessment
  IF EF < 40%:
    → HFrEF
    → Risk Score += 10
    → Recommend: Beta-blocker, ACE-I/ARB/ARNI, MRA, SGLT2i
    → If EF < 35% AND >40 days post-MI: Consider ICD
  ELSE IF EF 40-49%:
    → HFmrEF
    → Risk Score += 6
    → Recommend: Similar to HFrEF treatment
  ELSE IF EF ≥ 50%:
    → Proceed to diastolic assessment

STEP 2: Diastolic Function Assessment (if EF ≥50%)
  IF Average E/e' > 14:
    → Elevated filling pressure
    → Risk Score += 8
  ELSE IF LAVI > 34 mL/m²:
    → LA remodeling
    → Risk Score += 6
  ELSE IF LARS ≤ 24% (if available):
    → LA dysfunction
    → Risk Score += 5

  IF any of above TRUE:
    → HFpEF likely
    → Recommend: Diuretics, treat comorbidities (HTN, DM, AF), SGLT2i

STEP 3: Valvular Disease
  IF Severe AS (Vmax ≥4.0, AVA ≤1.0, gradient ≥40):
    → Risk Score += 12
    → Urgent cardiology referral for AVR evaluation
  IF Severe MR or AR:
    → Risk Score += 8
    → Cardiology referral

STEP 4: Pulmonary Hypertension
  IF RVSP ≥ 60 mmHg:
    → Severe PH
    → Risk Score += 10
    → Comprehensive PH workup
  ELSE IF RVSP 40-59 mmHg:
    → Moderate PH
    → Risk Score += 6

STEP 5: Arrhythmia Risk (ECG)
  IF QTc ≥ 500 ms:
    → High sudden death risk
    → Risk Score += 10
    → Avoid QT-prolonging drugs, monitor electrolytes
  ELSE IF QTc 450-499 ms:
    → Moderate risk
    → Risk Score += 5
  IF AFib:
    → Risk Score += 6
    → Calculate CHA2DS2-VASc, consider anticoagulation
  IF HR ≥ 90 bpm at rest:
    → Risk Score += 4

STEP 6: Advanced Markers (if available)
  IF GLS > -16% (less negative):
    → Subclinical dysfunction
    → Risk Score += 5
    → Closer monitoring

RISK SCORE INTERPRETATION:
  0-10: Low Risk → Routine follow-up
  11-20: Moderate Risk → 6-month cardiology follow-up
  21-35: High Risk → 3-month follow-up, optimization of therapy
  >35: Very High Risk → Urgent cardiology evaluation, consider advanced HF referral

OUTPUT: Risk Category + Specific Recommendations
```

---

### Algorithm 2: Pre-Operative Cardiac Risk Assessment

```
INPUT: Echo + ECG + Patient Demographics + Surgical Risk

STEP 1: Surgical Risk Classification
  Low Risk (<1% cardiac event rate): Endoscopy, cataract, superficial
  Intermediate Risk (1-5%): Carotid, peripheral vascular, major orthopedic
  High Risk (>5%): Vascular surgery (AAA), major thoracic/abdominal

STEP 2: Functional Capacity (METS)
  ≥4 METS (climb 1 flight of stairs): Proceed if low-intermediate surgery
  <4 METS OR Unknown: Proceed to Step 3

STEP 3: Echo Risk Factors
  IF EF < 35%:
    → High Risk
    → Consider pre-op hemodynamic optimization
    → May need ICU post-op
  IF Severe AS (symptomatic):
    → Very High Risk (up to 10% mortality)
    → Require AVR or TAVR before elective surgery
  IF Severe PH (RVSP >60):
    → High Risk
    → Pre-op RV function optimization

STEP 4: ECG Risk Factors
  IF AFib:
    → Ensure rate control (<100 bpm at rest)
    → Continue anticoagulation per surgical bleeding risk
  IF QTc ≥ 500 ms:
    → Arrhythmia risk with anesthesia
    → Cardiology clearance
  IF High-grade AV block:
    → Consider temporary pacemaker

STEP 5: Calculate RCRI (Revised Cardiac Risk Index)
  Points (1 each):
    - High-risk surgery
    - Ischemic heart disease
    - Heart failure (EF <40% or HF history)
    - Cerebrovascular disease
    - Diabetes on insulin
    - Creatinine >2 mg/dL

  RCRI 0-1: Risk 0.4-1% → Proceed
  RCRI 2: Risk 2.4% → Optimize, proceed
  RCRI ≥3: Risk 5.4%+ → Intensive optimization, consider surgery necessity

OUTPUT: Risk Category + Recommendations (proceed, optimize, delay)
```

---

### Algorithm 3: Atrial Fibrillation Risk Prediction

```
INPUT: Echo + Clinical Data

STEP 1: Structural Risk Factors
  IF LAVI > 34 mL/m²:
    → LA enlargement
    → Risk Score += 8
    → Each 5 mL/m² above 34: +40% AFib risk
  IF LARS < 24% (if available):
    → LA dysfunction
    → Risk Score += 6
  IF LVH present (mass index >115/95 g/m²):
    → Risk Score += 5

STEP 2: Diastolic Dysfunction
  IF E/e' > 14:
    → Elevated filling pressure
    → Risk Score += 6
  IF Diastolic dysfunction Grade II or III:
    → Risk Score += 5

STEP 3: Age & Comorbidities
  Age ≥65: Risk Score += 5
  Age ≥75: Risk Score += 8
  Hypertension: Risk Score += 4
  Heart failure: Risk Score += 6
  Diabetes: Risk Score += 3
  Obesity (BMI >30): Risk Score += 3

STEP 4: Valvular Disease
  Moderate-severe MR: Risk Score += 5
  Moderate-severe MS: Risk Score += 8

RISK SCORE INTERPRETATION:
  0-10: Low Risk (1-2% annual AFib incidence) → No specific monitoring
  11-20: Moderate Risk (3-5% annual) → Consider annual ECG or event monitor if symptoms
  21-35: High Risk (6-10% annual) → Annual ECG, low threshold for extended monitoring
  >35: Very High Risk (>10% annual) → Consider screening with mobile ECG, wearables

IF AFib Detected:
  → Calculate CHA2DS2-VASc for stroke risk
  → Score ≥2 (men) or ≥3 (women): Anticoagulation indicated

OUTPUT: AFib Risk Category + Monitoring Recommendations
```

---

### Algorithm 4: Heart Failure Therapy Optimization

```
INPUT: Echo + Clinical Status + Current Medications

STEP 1: Phenotype Classification
  IF EF < 40%:
    → HFrEF Pathway
  ELSE IF EF 40-49%:
    → HFmrEF Pathway (treat similar to HFrEF)
  ELSE IF EF ≥ 50% AND (E/e' >14 OR LAVI >34 OR elevated BNP):
    → HFpEF Pathway

--- HFrEF/HFmrEF Pathway ---

STEP 2: Guideline-Directed Medical Therapy (GDMT) - "Fantastic Four"

  1. Beta-Blocker (carvedilol, metoprolol succinate, bisoprolol)
     IF NOT on beta-blocker:
       → Start low, titrate up
       Target: HR 50-60 bpm

  2. ACE-I/ARB/ARNI (sacubitril-valsartan)
     IF NOT on RAAS inhibitor:
       → Start ACE-I or ARB
     IF on ACE-I/ARB AND EF still <40%:
       → Consider switch to ARNI (sacubitril-valsartan)
       → Superior to ACE-I for mortality reduction

  3. Mineralocorticoid Receptor Antagonist (spironolactone, eplerenone)
     IF EF ≤35% AND NYHA II-IV:
       → Add MRA
       → Monitor K+ and creatinine closely

  4. SGLT2 Inhibitor (dapagliflozin, empagliflozin)
     IF EF ≤40%:
       → Add SGLT2i (even without diabetes)
       → Proven mortality benefit

STEP 3: Device Therapy
  IF EF ≤35% AND NYHA II-III AND on optimal medical therapy >3 months:
    IF QRS ≥150 ms with LBBB:
      → CRT-D (cardiac resynchronization + defibrillator)
    ELSE IF QRS ≥120 ms:
      → Consider CRT
    ELSE:
      → ICD (implantable cardioverter-defibrillator)

  IF >40 days post-MI AND EF ≤30%:
    → ICD for primary prevention

STEP 4: Serial Echo Monitoring
  Repeat echo in 3-6 months after medication optimization
  IF EF improved ≥10% points OR now >40%:
    → Reverse remodeling
    → Continue current therapy
    → Less intensive monitoring
  IF no improvement or worsening:
    → Reassess adherence, titration
    → Consider advanced HF referral

--- HFpEF Pathway ---

STEP 2: Treat Underlying Conditions
  Hypertension: Target BP <130/80 mmHg
  Diabetes: Glycemic control, add SGLT2i (shown to reduce HF hospitalization)
  Obesity: Weight loss (bariatric surgery if BMI >40)
  AFib: Rate control, consider rhythm control, anticoagulation
  OSA: CPAP therapy

STEP 3: Diuretics for Volume Overload
  IF congestion present (elevated JVP, edema, DOE):
    → Loop diuretic (furosemide, torsemide)
    → Titrate to euvolemia
    → Monitor electrolytes

STEP 4: Evidence-Based HFpEF Therapy
  SGLT2 Inhibitor (dapagliflozin, empagliflozin):
    → Reduces HF hospitalization in HFpEF
    → Start in all HFpEF patients unless contraindicated

  IF AFib present:
    → Rate control critical (beta-blocker, diltiazem)

  IF Obesity + HFpEF:
    → GLP-1 agonist (emerging evidence for benefit)

STEP 5: Serial Echo Monitoring
  Repeat echo annually or with clinical change
  Monitor E/e', LAVI, LARS (if available)
  Track progression of diastolic dysfunction

OUTPUT: Specific Medication Recommendations + Device Evaluation + Follow-Up Plan
```

---

## Functional Medicine Integration

### Overview

Functional medicine emphasizes root-cause analysis, nutrient optimization, and evidence-based supplementation to complement conventional cardiac care. The following protocols integrate recent evidence for cardiac risk reduction.

---

### 1. Coenzyme Q10 (CoQ10)

#### Evidence Summary (2024-2025)

**Systematic Reviews & Meta-Analyses:**
- **14 RCTs (2,149 HF patients):** CoQ10 supplementation (100-200 mg/day for 3-12 months) significantly reduced mortality vs. placebo (RR = 0.69; 95% CI 0.50–0.95; p = 0.02)
- **Q-SYMBIO Trial:** Landmark RCT showing reduction in major adverse cardiovascular events in contemporary heart failure population

**2025 RCT Findings:**
- **Functional capacity:** 6-minute walk test significantly improved in CoQ10 group (349.3 ± 100.6 m vs. 267.0 ± 96.5 m in placebo, p = 0.008)
- Quality of life improvements in NYHA class

**Diastolic Dysfunction Specific:**
- **Isolated diastolic dysfunction:** Improvement in ≥1 diastolic parameter in hypertrophic cardiomyopathy patients
- **HFpEF pilot trial:** Mixed results - did NOT significantly affect echocardiographic indices or NT-proBNP in elderly HFpEF patients (suggests benefit may be condition-specific)

#### Mechanism of Action

- **Mitochondrial bioenergetics:** Essential for ATP production in myocardium
- **Antioxidant properties:** Reduces oxidative stress
- **Statin depletion:** Statins inhibit CoQ10 synthesis; supplementation restores levels
- **Endothelial function:** Improves NO bioavailability

#### Clinical Recommendations

**Indications:**
- Heart failure (HFrEF, HFmrEF)
- Statin therapy (all patients on statins should consider CoQ10)
- Hypertrophic cardiomyopathy
- Diastolic dysfunction (selected cases)
- Chronic coronary syndrome

**Dosing:**
- **Standard:** 100-200 mg daily (divided doses)
- **Heart failure:** 200-300 mg daily (Q-SYMBIO used 300 mg)
- **With meals:** Fat-soluble; absorption enhanced with dietary fat

**Formulation:**
- **Ubiquinone:** Standard form, requires conversion
- **Ubiquinol:** Reduced, active form; better absorption in elderly or those with HF
- Prefer ubiquinol in patients >60 years or with severe HF

**Monitoring:**
- **Baseline:** Serum CoQ10 level (optional)
- **3-6 months:** Reassess symptoms, functional capacity (6MWT), echo parameters
- **Safety:** Excellent safety profile; rare GI upset

**Contraindications:**
- None absolute
- Use caution with warfarin (may reduce warfarin effect; monitor INR)

**EMR Integration:**
- Flag all statin prescriptions for CoQ10 recommendation
- Track supplementation adherence
- Document functional improvements (NYHA class, 6MWT)

#### Evidence Quality: ★★★★☆ (Strong evidence for HFrEF, moderate for diastolic dysfunction)

---

### 2. Omega-3 Fatty Acids (EPA/DHA)

#### Evidence Summary (2024-2025)

**Complex and Dose-Dependent Effects:**

**Heart Failure:**
- **Systematic review:** Omega-3 mildly reduces risk of hospitalization and all-cause mortality in HF patients
- **AHA/ACC/HFSA 2022 Guideline:** Class IIb recommendation - omega-3 for HF (NYHA II-IV) to reduce CV hospitalization and death

**Arrhythmia - CONFLICTING EVIDENCE:**

**Protective (Dietary Amounts):**
- Higher dietary omega-3 consumption associated with decreased AFib risk
- Mechanism: Anti-inflammatory, membrane stabilization

**Harmful (High Pharmaceutical Doses):**
- **Dose-dependent AFib risk:** High-risk CVD patients on high doses (>1,500 mg/day) showed increased AF risk (pooled OR 1.48; 95% CI 1.21-1.81)
- **Meta-analysis (34 trials, 114,326 subjects, 2025):** EPA/DHA most likely to increase AFib in high-risk CVD patients on high doses
- **Mechanism hypothesis:** High doses may increase vagal tone, promoting AFib

**Resolution:**
- **Dietary omega-3 (from fish) = PROTECTIVE**
- **Pharmaceutical omega-3 >1.5g/day in high-risk patients = RISK**

**Recent 2025/2026 Research:**
- **Mortality benefit:** Association of dietary omega-3 intake with reduced all-cause and CVD-specific mortality in CVD patients
- **Biomarkers (Jan 2026):** Omega-3 biomarkers support lower atrial fibrillation risk (dietary, not supplemental)

#### Mechanism of Action

- **Anti-inflammatory:** Reduces CRP, IL-6
- **Membrane incorporation:** Alters myocyte electrical properties
- **Triglyceride lowering:** Especially EPA (icosapent ethyl)
- **Endothelial function:** Improves flow-mediated dilation

#### Clinical Recommendations

**Indications:**
- **Heart failure (NYHA II-IV):** Modest benefit for hospitalization/mortality reduction
- **Post-MI:** Omega-3 may reduce recurrent events (dose-dependent)
- **Hypertriglyceridemia:** High-dose EPA (icosapent ethyl 4g) for TG >150 mg/dL
- **General CVD prevention:** Dietary sources preferred

**Dosing - CRITICAL DISTINCTION:**

**1. Dietary Omega-3 (PREFERRED FOR MOST):**
- **Target:** 2-3 servings fatty fish/week (salmon, mackerel, sardines, herring)
- **Equivalent:** ~1,000-1,500 mg EPA+DHA per week from food
- **Mechanism:** Anti-inflammatory, AFib protective

**2. Supplemental Omega-3 (MODERATE DOSE):**
- **Standard:** 500-1,000 mg EPA+DHA daily (combined)
- **Use for:** General CV health, mild TG elevation, HF
- **Safety:** Appears safe for AFib risk at these doses

**3. High-Dose Pharmaceutical (USE WITH CAUTION):**
- **Icosapent ethyl (Vascepa):** 4g/day EPA (FDA-approved for TG >150 mg/dL + CVD)
- **WARNING:** May increase AFib risk in high-risk patients
- **Risk stratification before high-dose omega-3:**
  - History of AFib? → AVOID high-dose
  - LAVI >34 mL/m² or E/e' >14? → CAUTION (high AFib risk substrate)
  - Elderly with HTN + LVH? → CAUTION
  - Younger, normal echo? → Lower risk

**Monitoring:**
- **Baseline echo:** Assess LAVI, E/e' (AFib risk substrates)
- **ECG:** Baseline, then 3-6 months after starting high-dose (if used)
- **Symptoms:** Palpitations (may indicate AFib)
- **Triglycerides:** If indication was TG lowering

**Contraindications:**
- **Allergy to fish/shellfish**
- **Active AFib on high-dose omega-3:** Discontinue high-dose, switch to dietary or low-dose

**EMR Integration:**
- **Decision support alert:** If prescribing high-dose omega-3 (>1.5g) + patient has LAVI >34 or AFib history → WARNING
- Recommend dietary omega-3 over supplemental for most patients
- Track AFib incidence in patients on high-dose omega-3

#### Evidence Quality:
- **HF benefit:** ★★★☆☆ (Moderate, modest effect)
- **AFib risk (high-dose):** ★★★★☆ (Strong evidence for dose-dependent risk)
- **Dietary protective effect:** ★★★★☆ (Strong epidemiological evidence)

---

### 3. Magnesium

#### Evidence Summary (2024-2025)

**Diastolic Dysfunction:**
- **Animal model (2021):** Low-Mg mice exhibited impaired cardiac relaxation, reversible diastolic cardiomyopathy associated with mitochondrial dysfunction and oxidative modification of cMyBPC
- **Mechanism:** Magnesium deficiency leads to myocardial fibrosis, LVH, and dysfunction → incident HF in diabetes patients
- **Plausibility:** Magnesium supplementation may delay HF onset by reducing LV fibrosis and improving systolic/diastolic function

**Heart Failure Mortality:**
- **Propensity-matched cohort:** Low magnesium associated with increased cardiovascular mortality (HR 1.38; 95% CI 1.04–1.83; p=0.024)
- **Rates:** 20% CV mortality in normal-Mg vs. 24% in low-Mg patients

**Arrhythmias:**
- Oral magnesium supplementation reduced:
  - Mean arterial pressure
  - Systolic vascular resistance
  - Frequency of isolated ventricular premature complexes, couplets, non-sustained VT
- **Clinical use:** Torsades de pointes (IV magnesium), atrial fibrillation, ventricular arrhythmias

**Recent Research:**
- **2025 propensity score study (HFpEF in ICU):** Magnesium supplementation effects on mortality in critically ill HFpEF
- **Diabetes cohort:** Nonprescription magnesium supplement use reduced risk of HF in patients with diabetes (target trial emulation)

#### Mechanism of Action

- **Calcium antagonism:** Natural calcium channel blocker, promotes vasodilation
- **Arrhythmia suppression:** Stabilizes myocyte membranes, reduces ectopy
- **Mitochondrial function:** Essential cofactor in ATP synthesis
- **Anti-fibrotic:** Reduces myocardial collagen deposition
- **Blood pressure:** Reduces vascular resistance

#### Clinical Recommendations

**Indications:**
- **Diastolic dysfunction (especially HFpEF)**
- **Arrhythmias:** VPCs, non-sustained VT, AFib
- **Heart failure (low magnesium)**
- **Hypertension**
- **Diabetes with LVH or diastolic dysfunction**

**Screening:**
- **Serum magnesium:** Often normal despite tissue depletion (only 1% of Mg is extracellular)
- **Better markers (if available):** RBC magnesium, ionized magnesium
- **Clinical clues:**
  - PPI use (reduces absorption)
  - Diuretic use (increases loss)
  - Diabetes (renal wasting)
  - Alcohol use disorder

**Dosing:**
- **Supplementation:** 300-400 mg elemental magnesium daily
- **Forms:**
  - **Magnesium glycinate:** Best absorption, least GI upset
  - **Magnesium citrate:** Good absorption, mild laxative effect
  - **Magnesium oxide:** Poor absorption, not recommended
- **Dietary sources:** Leafy greens, nuts, seeds, whole grains, dark chocolate

**Monitoring:**
- **Baseline:** Serum magnesium, creatinine (clearance)
- **3-6 months:** Recheck magnesium, assess symptoms (palpitations, fatigue), consider repeat echo for diastolic parameters
- **Safety:** Monitor in renal impairment (Mg is renally excreted)

**Contraindications:**
- **Severe renal impairment (eGFR <30):** Risk of hypermagnesemia
- **AV block:** Magnesium can slow AV conduction
- **Myasthenia gravis:** May worsen weakness

**Drug Interactions:**
- **Bisphosphonates:** Separate dosing by 2 hours
- **Antibiotics (fluoroquinolones, tetracyclines):** Separate by 2-4 hours

**EMR Integration:**
- Flag patients on diuretics for magnesium supplementation
- Alert if low serum Mg + arrhythmia detected
- Track correlation between Mg supplementation and arrhythmia burden (if using continuous monitoring)

#### Evidence Quality:
- **Deficiency harms:** ★★★★☆ (Strong observational evidence)
- **Supplementation benefit:** ★★★☆☆ (Moderate; human RCTs for diastolic HF still needed)
- **Arrhythmia benefit:** ★★★★☆ (Strong for IV use in acute settings, moderate for oral prevention)

---

### 4. Additional Functional Medicine Considerations

#### L-Carnitine
- **Mechanism:** Facilitates fatty acid transport into mitochondria
- **Evidence:** Small studies show improvement in exercise capacity in HF
- **Dose:** 1-2g daily
- **Quality:** ★★☆☆☆ (Limited evidence)

#### D-Ribose
- **Mechanism:** Substrate for ATP regeneration
- **Evidence:** May improve diastolic function in HFpEF (very small studies)
- **Dose:** 5g TID
- **Quality:** ★☆☆☆☆ (Preliminary)

#### Hawthorn Extract
- **Mechanism:** Positive inotrope, vasodilator
- **Evidence:** European studies show modest benefit in NYHA II HF
- **Dose:** 900-1,800 mg daily (standardized extract)
- **Quality:** ★★★☆☆ (Moderate)

#### Vitamin D
- **Mechanism:** Modulates RAAS, reduces inflammation
- **Evidence:** Deficiency associated with worse HF outcomes; supplementation trials show modest functional improvement
- **Target:** 25-OH vitamin D >30 ng/mL (ideally 40-60)
- **Dose:** 2,000-5,000 IU daily (titrate to level)
- **Quality:** ★★★☆☆ (Moderate)

---

### Functional Medicine Protocol for Cardiac Patients

**Step 1: Assess Baseline**
- Comprehensive echo (EF, diastolic function, LAVI, GLS)
- ECG (QTc, rhythm, LVH)
- Labs: BNP/NT-proBNP, CRP, lipids, HbA1c, TSH, vitamin D, magnesium, CoQ10 (optional)

**Step 2: Risk Stratification**
- Use algorithms above to classify HF risk
- Identify specific deficits (systolic vs. diastolic, valvular, arrhythmia)

**Step 3: Conventional Therapy Optimization**
- Ensure GDMT for HFrEF (beta-blocker, RAAS inhibitor, MRA, SGLT2i)
- Treat HTN, DM, dyslipidemia aggressively

**Step 4: Functional Medicine Add-Ons**

**All Cardiac Patients:**
- **CoQ10 100-200 mg daily** (especially if on statins)
- **Magnesium glycinate 300-400 mg daily** (if not contraindicated)
- **Dietary omega-3:** 2-3 servings fatty fish/week
- **Vitamin D:** Supplement to achieve 40-60 ng/mL

**Heart Failure (HFrEF/HFmrEF):**
- **CoQ10 200-300 mg daily** (strong evidence)
- **Omega-3 1,000 mg daily** (moderate evidence)
- Consider L-carnitine, hawthorn (moderate evidence)

**Diastolic Dysfunction/HFpEF:**
- **Magnesium glycinate 400 mg daily** (emerging evidence)
- **CoQ10 100-200 mg daily** (mixed evidence; trial in selected patients)
- Aggressive BP control, weight loss, SGLT2i

**Arrhythmia (VPCs, AFib):**
- **Magnesium glycinate 400 mg daily** (strong evidence for VPC reduction)
- **Dietary omega-3 ONLY** (avoid high-dose supplements >1.5g)
- Correct electrolyte abnormalities (K+, Mg++)

**Step 5: Monitor & Adjust**
- **3 months:** Reassess symptoms, functional capacity, labs
- **6 months:** Repeat echo (if abnormal at baseline), ECG
- **Document improvements:** NYHA class, 6MWT, QOL scores, echo parameters

---

## Implementation Recommendations

### For Plenya EMR System

#### 1. Database Schema

**Phase 1 (MVP):**
- Implement **Option B** for ECG (comprehensive table with key parameters)
- Implement **Option B** for Echo (core parameters: EF, LAVI, E/e', GLS, RVSP, valve severity)
- Both should store `report_pdf_url` for full report access

**Phase 2 (Advanced):**
- Expand to full structured tables (Option C)
- Add time-series queries for trend analysis
- Integrate with risk stratification algorithms

---

#### 2. Clinical Decision Support (CDS)

**Automated Alerts:**

**Critical Alerts (Immediate Action):**
- ECG: QTc ≥500 ms, new STEMI, high-grade AV block, sustained VT
- Echo: EF <30%, severe symptomatic AS, tamponade physiology

**Warning Alerts (Requires Follow-Up):**
- ECG: QTc 450-499 ms, new AFib, HR >120 at rest
- Echo: EF 30-40%, RVSP ≥60, severe MR/AR, E/e' >14

**Medication Interaction Alerts:**
- QTc >450 ms + prescription of QT-prolonging drug → WARNING
- High-dose omega-3 (>1.5g) + LAVI >34 or AFib history → WARNING
- Magnesium supplement + eGFR <30 → WARNING

**Guideline-Based Prompts:**
- EF <35% + >3 months on optimal therapy → "Consider ICD evaluation"
- Severe AS + symptoms → "Urgent cardiology referral for AVR"
- HFrEF without SGLT2i → "Consider adding SGLT2 inhibitor (mortality benefit)"

---

#### 3. Risk Scores & Dashboards

**Heart Failure Risk Dashboard:**
- Display latest EF, E/e', LAVI, GLS, BNP
- Calculate composite risk score (Algorithm 1)
- Show trend over time (serial echos)
- Flag patients overdue for follow-up echo

**AFib Risk Predictor:**
- Input: LAVI, LARS, E/e', age, HTN, LVH
- Output: Estimated 5-year AFib risk
- Recommendation for monitoring intensity

**Pre-Op Risk Calculator:**
- Integrate RCRI with echo parameters
- Stratify surgical risk by procedure type
- Generate pre-op clearance note

---

#### 4. Physician Workflows

**ECG Workflow:**
1. Technician performs ECG → Automated interpretation stored
2. ECG flagged for physician review (required for all)
3. Physician reviews automated + raw ECG → Documents overread
4. Critical findings trigger alerts
5. ECG marked as "verified" → Timestamped

**Echo Workflow:**
1. Sonographer performs study → Preliminary measurements
2. Interpreting cardiologist reviews images → Enters final measurements
3. Auto-population of structured fields (EF, LAVI, etc.)
4. Free-text impression + recommendations
5. PDF report generated + stored
6. CDS algorithms run on structured data → Alerts/prompts

---

#### 5. Patient-Facing Features

**Patient Portal:**
- Layman-friendly echo results:
  - "Your heart's pumping strength (ejection fraction) is normal at 60%"
  - Visual: Green/yellow/red zones for key parameters
- ECG results:
  - "Your heart rhythm is normal sinus rhythm"
  - "Your heart rate is 75 beats per minute (normal range)"
- Trend graphs for serial measurements
- Educational content based on findings (e.g., "What is diastolic dysfunction?")

---

#### 6. LGPD/Security Considerations

**Access Control:**
- Cardiac data is PHI (Protected Health Information)
- Role-based access: Cardiologists, primary care, patient
- Audit log for all accesses (required by LGPD)

**Data Retention:**
- Raw DICOM/PDF: Minimum 20 years (legal requirement for medical records)
- Structured data: Permanent retention
- Audit logs: Minimum 5 years

**De-Identification for Research:**
- Remove patient_id, replace with study ID
- Retain age ranges (not DOB), sex, echo/ECG parameters
- LGPD compliant for aggregated analysis

---

## References

### ECG 12-Lead Electrocardiogram

1. [Electrocardiographic Measures and Prediction of Cardiovascular and Noncardiovascular Death in CKD - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC4731112/)
2. [QTc prolongation measured by standard 12-lead electrocardiography is an independent risk factor for sudden death due to cardiac arrest - PubMed](https://pubmed.ncbi.nlm.nih.gov/2040041/)
3. [Normalization of Electrocardiogram-Derived Cardiac Risk Indices: A Scoping Review of the Open-Access Literature | MDPI](https://www.mdpi.com/2076-3417/14/20/9457)
4. [Independent and Incremental Value of ECG Markers for Prediction of Cancer Therapy‐Related Cardiac Dysfunction - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC12184591/)
5. [The 12-lead electrocardiogram and risk of sudden death: current utility and future prospects - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC4812823/)
6. [The Diagnostic and Prognostic Value of the 12-Lead ECG in Arrhythmogenic Left Ventricular Cardiomyopathy - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC12141911/)

### ECG Screening Guidelines (USPSTF)

7. [Recommendation: Atrial Fibrillation: Screening | United States Preventive Services Taskforce](https://www.uspreventiveservicestaskforce.org/uspstf/recommendation/atrial-fibrillation-screening)
8. [Recommendation: Cardiovascular Disease Risk: Screening With Electrocardiography | United States Preventive Services Taskforce](https://www.uspreventiveservicestaskforce.org/uspstf/recommendation/cardiovascular-disease-risk-screening-with-electrocardiography)
9. [Screening for Atrial Fibrillation: US Preventive Services Task Force Recommendation Statement - PubMed](https://pubmed.ncbi.nlm.nih.gov/35076659/)

### ECG LVH Criteria

10. [Evaluation of the electrocardiographic criteria for left ventricular hypertrophy diagnosis](http://cardiolatina.com/wp-content/uploads/2020/01/Electrocardiographic-criteria-for-Left-Ventricular-Hypertrophy-diagnosis.pdf)
11. [Electrocardiogram Performance in the Diagnosis of Left Ventricular Hypertrophy in Hypertensive Patients With Left Bundle Branch Block - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC5245847/)
12. [A comparison of Cornell and Sokolow-Lyon electrocardiographic criteria for left ventricular hypertrophy in a military male population in Taiwan - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC5440264/)
13. [ECG in left ventricular hypertrophy (LVH): criteria and implications](https://ecgwaves.com/topic/ecg-left-ventricular-hypertrophy-lvh-clinical-characteristics/)

### Echocardiography - General Guidelines

14. [Recommendations for Cardiac Chamber Quantification by Echocardiography in Adults: An Update from the American Society of Echocardiography and the European Association of Cardiovascular Imaging - ASE](https://www.asecho.org/wp-content/uploads/2016/02/2015_ChamberQuantificationREV.pdf)
15. [Standardization of adult transthoracic echocardiography reporting in agreement with recent chamber quantification, diastolic function, and heart valve disease recommendations - PubMed](https://pubmed.ncbi.nlm.nih.gov/29045589/)

### Echocardiography - Ejection Fraction & Systolic Function

16. [Ejection fraction, B-lines, and global longitudinal strain evaluated with rest transthoracic echocardiography to assess prognosis in patients with chronic coronary syndromes](https://www.explorationpub.com/Journals/ec/Article/10127)
17. [2024 Guidelines of the Taiwan Society of Cardiology for the Diagnosis and Treatment of Heart Failure with Preserved Ejection Fraction - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC10961629/)
18. [Left Ventricular Ejection Fraction - StatPearls - NCBI Bookshelf](https://www.ncbi.nlm.nih.gov/books/NBK459131/)
19. [Left ventricular ejection fraction: clinical, pathophysiological, and technical limitations | Frontiers](https://www.frontiersin.org/journals/cardiovascular-medicine/articles/10.3389/fcvm.2024.1340708/full)

### Echocardiography - Diastolic Function (2024-2025 Guidelines)

20. [Recommendations for the Evaluation of Left Ventricular Diastolic Function by Echocardiography and for Heart Failure With Preserved Ejection Fraction Diagnosis: An Update From the American Society of Echocardiography - Journal of the American Society of Echocardiography](https://onlinejase.com/article/S0894-7317(25)00157-9/fulltext)
21. [Diastolic dysfunction: a comparison of 2025 ASE, 2024 BSE and 2022 EACVI guidelines | European Heart Journal - Cardiovascular Imaging](https://academic.oup.com/ehjcimaging/article/26/11/1725/8254556)
22. [Diastolic function and cardiovascular events in patients with preserved left ventricular ejection fraction. Improving risk stratification with left atrial strain | Frontiers](https://www.frontiersin.org/journals/cardiovascular-medicine/articles/10.3389/fcvm.2025.1565052/full)
23. [Increased left atrial volume index predicts a poor prognosis in patients with heart failure - PubMed](https://pubmed.ncbi.nlm.nih.gov/21362529/)

### Echocardiography - Global Longitudinal Strain

24. [Global Longitudinal Strain by Echocardiography Predicts Long-Term Risk of Cardiovascular Morbidity and Mortality in a Low-Risk General Population | Circulation: Cardiovascular Imaging](https://www.ahajournals.org/doi/10.1161/circimaging.116.005521)
25. [Clinical Applications of Strain - ASE Guidelines August 2025](https://www.asecho.org/wp-content/uploads/2025/08/Strain-Guideline-AIP-August-2025.pdf)
26. [Speckle-Tracking Strain Echocardiography for the Assessment of Left Ventricular Structure and Function: A Scientific Statement From the American Heart Association | Circulation](https://www.ahajournals.org/doi/10.1161/CIR.0000000000001354)
27. [Prognostic Value of Strain by Speckle Tracking Echocardiography in Patients with Arrhythmogenic Right Ventricular Cardiomyopathy | MDPI](https://www.mdpi.com/2308-3425/11/12/388)

### Echocardiography - Valvular Disease

28. [Recommendations on the Echocardiographic Assessment of Aortic Valve Stenosis - EACVI/ASE 2017](https://www.asecho.org/wp-content/uploads/2017/04/2017ValveStenosisGuideline.pdf)
29. [Echocardiographic Evaluation of Aortic Stenosis: A Comprehensive Review - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC10417789/)
30. [Moderate gradient severe aortic stenosis: diagnosis, prognosis and therapy - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC8503314/)
31. [Echocardiographic assessment of aortic stenosis: a practical guideline from the British Society of Echocardiography - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC8115410/)

### Echocardiography - Pulmonary Hypertension & RV Function

32. [Guidelines for the Echocardiographic Assessment of Pulmonary Hypertension - ASE 2025](https://www.asecho.org/wp-content/uploads/2025/03/PIIS0894731725000379.pdf)
33. [What the Heck is the Cut-Off Value for RVSP?! - Cardioserv](https://www.cardioserv.net/echo-rvsp-cutoff-values/)
34. [Right Ventricular Systolic Pressure Trajectory as a Predictor of Hospitalization and Mortality in Patients With Chronic Heart Failure - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC10516718/)
35. [Association of Mild Echocardiographic Pulmonary Hypertension With Mortality and Right Ventricular Function | JAMA Cardiology](https://jamanetwork.com/journals/jamacardiology/fullarticle/2751319)
36. [Elevated Pulmonary Pressure Noted on Echocardiogram: A Simplified Approach to Next Steps | Journal of the American Heart Association](https://www.ahajournals.org/doi/10.1161/JAHA.120.017684)

### Functional Medicine - Coenzyme Q10

37. [Coenzyme Q10 and Heart Failure | Circulation: Heart Failure](https://www.ahajournals.org/doi/10.1161/circheartfailure.115.002639)
38. [Clinical Evidence for Q10 Coenzyme Supplementation in Heart Failure: From Energetics to Functional Improvement | MDPI](https://www.mdpi.com/2077-0383/9/5/1266)
39. [Effect of Coenzyme Q10 Supplementation on Cardiac Function and Quality of Life in Patients with Heart Failure: A Randomized Controlled Trial | MDPI 2025](https://www.mdpi.com/2077-0383/14/11/3675)
40. [Coenzyme Q10 in the Treatment of Heart Failure with Preserved Ejection Fraction: A Prospective, Randomized, Double-Blind, Placebo-Controlled Trial | Drugs in R&D](https://link.springer.com/article/10.1007/s40268-021-00372-1)
41. [Isolated diastolic dysfunction of the myocardium and its response to CoQ10 treatment - PubMed](https://pubmed.ncbi.nlm.nih.gov/8241699/)
42. [Coenzyme Q10 (CoQ10) in isolated diastolic heart failure in hypertrophic cardiomyopathy (HCM) - PubMed](https://pubmed.ncbi.nlm.nih.gov/19096110/)

### Functional Medicine - Omega-3 Fatty Acids

43. [Effect of omega-3 fatty acids on cardiovascular outcomes: A systematic review and meta-analysis - eClinicalMedicine](https://www.thelancet.com/journals/eclinm/article/PIIS2589-5370(21)00277-7/fulltext)
44. [Association of dietary omega-3 fatty acids intake with all-cause and cardiovascular disease-specific mortality | Scientific Reports 2025](https://www.nature.com/articles/s41598-025-21193-1)
45. [Omega-3 Fatty Acids and Arrhythmias | Circulation 2024](https://www.ahajournals.org/doi/10.1161/CIRCULATIONAHA.123.065769)
46. [Associations Between Plasma Omega‐3 and Fish Oil Use With Risk of Atrial Fibrillation in the UK Biobank | Journal of the American Heart Association 2025](https://www.ahajournals.org/doi/full/10.1161/JAHA.125.043031)
47. [Effects of Omega-3 Fatty Acid Treatment on Risk for Atrial Fibrillation: An Updated Meta-Analysis of 34 Trials including 114,326 Individuals | medRxiv 2025](https://www.medrxiv.org/content/10.64898/2025.12.14.25342167v1.full)
48. [Omega-3 and Risk of atrial fibrillation: Vagally-mediated double-edged sword - ScienceDirect 2024](https://www.sciencedirect.com/science/article/abs/pii/S0033062024001683)
49. [Biomarkers show omega-3s support lower atrial fibrillation risk - NutraIngredients Jan 2026](https://www.nutraingredients.com/Article/2026/01/02/omega-3s-support-lower-atrial-fibrillation-risk/)

### Functional Medicine - Magnesium

50. [Magnesium Deficiency Causes a Reversible, Metabolic, Diastolic Cardiomyopathy | Journal of the American Heart Association 2021](https://www.ahajournals.org/doi/full/10.1161/JAHA.120.020205)
51. [Nonprescription Magnesium Supplement Use and Risk of Heart Failure in Patients With Diabetes | Journal of the American Heart Association 2024](https://www.ahajournals.org/doi/10.1161/JAHA.124.038870)
52. [Magnesium for the prevention and treatment of cardiovascular disease - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC6045762/)
53. [Magnesium, Oxidative Stress, Inflammation, and Cardiovascular Disease | MDPI](https://www.mdpi.com/2076-3921/9/10/907)
54. [Low serum magnesium and cardiovascular mortality in chronic heart failure: a propensity-matched study - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC2721016/)
55. [Propensity score matched cohort study on magnesium supplementation and mortality in critically ill patients with HFpEF | Scientific Reports 2025](https://www.nature.com/articles/s41598-025-85931-1)

---

## Appendices

### Appendix A: Abbreviations

- **AFib/AF:** Atrial Fibrillation
- **AR:** Aortic Regurgitation
- **AS:** Aortic Stenosis
- **ASE:** American Society of Echocardiography
- **AVA:** Aortic Valve Area
- **AV Block:** Atrioventricular Block
- **BSA:** Body Surface Area
- **CRT:** Cardiac Resynchronization Therapy
- **CVD:** Cardiovascular Disease
- **DT:** Deceleration Time
- **EACVI:** European Association of Cardiovascular Imaging
- **ECG:** Electrocardiogram
- **EF:** Ejection Fraction
- **EROA:** Effective Regurgitant Orifice Area
- **FAC:** Fractional Area Change
- **GDMT:** Guideline-Directed Medical Therapy
- **GLS:** Global Longitudinal Strain
- **HF:** Heart Failure
- **HFpEF:** Heart Failure with Preserved Ejection Fraction
- **HFrEF:** Heart Failure with Reduced Ejection Fraction
- **HFmrEF:** Heart Failure with Mildly Reduced Ejection Fraction
- **HR:** Hazard Ratio (in statistics) or Heart Rate (in clinical context)
- **ICD:** Implantable Cardioverter-Defibrillator
- **IVC:** Inferior Vena Cava
- **LAD:** Left Axis Deviation (ECG) or Left Anterior Descending artery
- **LARS:** Left Atrial Reservoir Strain
- **LAVI:** Left Atrial Volume Index
- **LBBB:** Left Bundle Branch Block
- **LV:** Left Ventricle/Ventricular
- **LVEDV:** Left Ventricular End-Diastolic Volume
- **LVEF:** Left Ventricular Ejection Fraction
- **LVH:** Left Ventricular Hypertrophy
- **LVESV:** Left Ventricular End-Systolic Volume
- **LVFP:** Left Ventricular Filling Pressure
- **MR:** Mitral Regurgitation
- **MRA:** Mineralocorticoid Receptor Antagonist
- **PAP:** Pulmonary Artery Pressure
- **PH:** Pulmonary Hypertension
- **QTc:** Corrected QT Interval
- **RAD:** Right Axis Deviation
- **RAP:** Right Atrial Pressure
- **RBBB:** Right Bundle Branch Block
- **RCRI:** Revised Cardiac Risk Index
- **RV:** Right Ventricle/Ventricular
- **RVSP:** Right Ventricular Systolic Pressure
- **RWT:** Relative Wall Thickness
- **STEMI:** ST-Elevation Myocardial Infarction
- **SVT:** Supraventricular Tachycardia
- **TAPSE:** Tricuspid Annular Plane Systolic Excursion
- **TDI:** Tissue Doppler Imaging
- **TR:** Tricuspid Regurgitation
- **TTE:** Transthoracic Echocardiography
- **USPSTF:** United States Preventive Services Task Force
- **VF:** Ventricular Fibrillation
- **VT:** Ventricular Tachycardia
- **WMSI:** Wall Motion Score Index
- **WPW:** Wolff-Parkinson-White Syndrome

### Appendix B: Normal Reference Values Quick Reference

| Parameter | Normal Value | Units |
|-----------|--------------|-------|
| **ECG** |
| Heart Rate | 60-100 | bpm |
| PR Interval | 120-200 | ms |
| QRS Duration | <120 | ms |
| QTc (men) | <450 | ms |
| QTc (women) | <460 | ms |
| **Echo - LV Systolic** |
| EF (biplane Simpson) | 53-73% | % |
| LVEDV indexed (men) | 34-74 | mL/m² |
| LVEDV indexed (women) | 29-61 | mL/m² |
| GLS | ≤-18% (more negative) | % |
| LV Mass Index (men) | <115 | g/m² |
| LV Mass Index (women) | <95 | g/m² |
| **Echo - Diastolic** |
| E/A Ratio | 0.8-2.0 | - |
| Septal e' | >7 | cm/s |
| Lateral e' | >10 | cm/s |
| E/e' Ratio | <8 | - |
| LAVI | <34 | mL/m² |
| LARS | >24% | % |
| DT | 160-240 | ms |
| **Echo - RV/PH** |
| RVSP | <30 | mmHg |
| TAPSE | ≥17 | mm |
| RV FAC | ≥35% | % |
| RV S' | ≥9.5 | cm/s |

---

**Document Version:** 1.0
**Author:** Claude Sonnet 4.5 (Anthropic)
**For:** Plenya EMR System
**Last Updated:** January 19, 2026
