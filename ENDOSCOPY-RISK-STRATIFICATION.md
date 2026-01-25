# Endoscopic Procedures Risk Stratification Guide

**Version:** 1.0
**Last Updated:** January 2026
**Purpose:** Clinical decision support for endoscopic findings documentation and risk stratification

---

## Table of Contents

1. [Executive Summary](#executive-summary)
2. [Endoscopia Digestiva Alta (EDA)](#endoscopia-digestiva-alta-eda)
3. [Colonoscopy](#colonoscopy)
4. [CSV Stratification Feasibility](#csv-stratification-feasibility)
5. [Implementation Recommendations](#implementation-recommendations)
6. [References](#references)

---

## Executive Summary

### Key Findings

**CSV Stratification Potential:**
- **PARTIALLY SUITABLE** for both EDA and colonoscopy
- Most findings are **categorical/ordinal** rather than continuous
- Best approach: **Clinical reporting guidelines + surveillance interval algorithms**
- Some components can be converted to risk scores (e.g., OLGA/OLGIM, Baveno VII)

**Recommended Implementation Strategy:**

1. **Structured Categorical Documentation** (Primary)
   - Use standardized classifications (LA, OLGA/OLGIM, Prague, BBPS)
   - Implement dropdown/checkbox interfaces
   - Auto-calculate surveillance intervals

2. **Composite Risk Scores** (Secondary)
   - OLGA/OLGIM stages → Gastric cancer risk (0-5% annual risk)
   - Baveno VII criteria → Variceal bleeding risk
   - Adenoma burden → Colorectal cancer risk stratification

3. **Quality Metrics** (Mandatory)
   - Adenoma Detection Rate (ADR): >30% men, >20% women
   - Sessile Serrated Lesion Detection Rate (SSLDR)
   - Boston Bowel Preparation Scale (BBPS): ≥6 adequate, ≥7 optimal

---

## Endoscopia Digestiva Alta (EDA)

### Overview

**Procedure:** Esophagogastroduodenoscopy (EGD)
**Purpose:** Visualize upper GI tract (esophagus, stomach, duodenum)
**Clinical Uses:** Diagnose reflux, gastritis, ulcers, Barrett's, varices, cancer screening

### 1. Esophagitis: LA Classification

**System:** Los Angeles Classification (most widely validated)
**Type:** Semi-quantitative ordinal scale
**CSV Suitability:** MEDIUM - Can assign numeric values for trending

#### Classification Grades

| Grade | Definition | CSV Value | Clinical Significance |
|-------|------------|-----------|----------------------|
| **Grade A** | ≤5mm mucosal breaks, not continuous | 1 | Mild - Can attempt PPI tapering |
| **Grade B** | >5mm breaks, not continuous between mucosal folds | 2 | Moderate - Continue PPI therapy |
| **Grade C** | Continuous breaks <75% circumference | 3 | Severe - Repeat endoscopy in 8 weeks |
| **Grade D** | Continuous breaks ≥75% circumference | 4 | Very severe - Long-term PPI, consider fundoplication |

#### Management Guidelines (2024-2025)

- **Grade A/B:** Attempt to discontinue or decrease PPI after 8 weeks
- **Grade C/D:** Continue PPI indefinitely; virtually 100% recurrence if stopped
- **Grade C/D:** Repeat endoscopy after 8 weeks of treatment (ACG recommendation)

#### Database Schema

```json
{
  "finding": "esophagitis_la",
  "grade": "A|B|C|D",
  "numeric_value": 1-4,
  "requires_followup": boolean,
  "followup_interval_weeks": 8,
  "ppi_recommendation": "taper|continue|long_term"
}
```

**References:**
- [Endoscopy Campus - LA Classification](https://www.endoscopy-campus.com/en/classifications/reflux-esophagitis-los-angeles-classification/)
- [Clinical Gastroenterology and Hepatology - LA Grades Comparison](https://www.cghjournal.org/article/S1542-3565(24)00480-4/fulltext)

---

### 2. Barrett's Esophagus: Prague Classification

**System:** Prague C&M Criteria
**Type:** Quantitative (circumferential + maximal extent in cm)
**CSV Suitability:** HIGH - Continuous measurements

#### Prague C&M Criteria

| Parameter | Definition | Measurement |
|-----------|------------|-------------|
| **C (Circumferential)** | Length of circumferential Barrett's mucosa | cm from GE junction |
| **M (Maximal)** | Maximal extent of Barrett's tongues | cm from GE junction |
| **Landmarks** | Distance from incisors to GE junction | Document consistently |

#### Risk Stratification

| Segment Length | Classification | Annual Progression Risk | Surveillance Interval |
|----------------|----------------|-------------------------|----------------------|
| **C<3 cm** | Short-segment Barrett's | 0.06% to HGD/EAC | 5 years |
| **C≥3 cm** | Long-segment Barrett's | 0.31% to HGD/EAC | 3 years |

**Risk Factor:** Increasing segment length OR 1.25 per cm (95% CI 1.16-1.36)

#### Dysplasia Grading

| Grade | Management | Intervention |
|-------|------------|--------------|
| **Nondysplastic** | Surveillance per segment length | None |
| **Low-grade dysplasia (LGD)** | Confirm with 2nd pathologist | Endoscopic eradication therapy |
| **High-grade dysplasia (HGD)** | Confirm with 2nd pathologist | Endoscopic eradication therapy (mandatory) |

#### Database Schema

```json
{
  "finding": "barretts_esophagus",
  "prague_c_cm": 0-15,
  "prague_m_cm": 0-15,
  "segment_classification": "short|long",
  "dysplasia_grade": "none|LGD|HGD",
  "dysplasia_confirmed": boolean,
  "annual_progression_risk_percent": 0.06-0.31,
  "surveillance_interval_years": 3-5,
  "eradication_indicated": boolean
}
```

**References:**
- [ACG Guideline - Barrett's Esophagus Diagnosis and Management](https://pmc.ncbi.nlm.nih.gov/articles/PMC10259184/)
- [Prague Classification Validation Study](https://pmc.ncbi.nlm.nih.gov/articles/PMC4547779/)

---

### 3. Gastritis: Sydney System & OLGA/OLGIM Staging

**System:** Updated Sydney System (1996) + OLGA/OLGIM (risk stratification)
**Type:** Ordinal grading + Stage-based risk scoring
**CSV Suitability:** MEDIUM to HIGH for OLGA/OLGIM stages

#### Sydney System Components

| Parameter | Grading | Score |
|-----------|---------|-------|
| **Chronic inflammation** | None/Mild/Moderate/Severe | 0-3 |
| **Neutrophil activity** | None/Mild/Moderate/Severe | 0-3 |
| **Atrophy** | None/Mild/Moderate/Severe | 0-3 |
| **Intestinal metaplasia** | None/Mild/Moderate/Severe | 0-3 |
| **H. pylori density** | None/Mild/Moderate/Severe | 0-3 |

**Topography:** Antrum + Corpus biopsies (minimum 2 sites each)

#### OLGA/OLGIM Staging for Gastric Cancer Risk

**OLGA (Operative Link for Gastritis Assessment):** Based on atrophy
**OLGIM (Operative Link for Gastric Intestinal Metaplasia):** Based on intestinal metaplasia

| Stage | Definition | Gastric Cancer Risk | Surveillance |
|-------|------------|---------------------|--------------|
| **Stage 0** | No atrophy/IM | Very low | None |
| **Stage I** | Minimal atrophy/IM | Low | None or 10 years |
| **Stage II** | Moderate atrophy/IM | Intermediate | 3-5 years (high-incidence regions) |
| **Stage III** | Extensive atrophy/IM in one compartment | High | 3 years |
| **Stage IV** | Extensive atrophy/IM in both compartments | Very high | 1-3 years |

#### Quantitative Risk (2025 Meta-Analysis)

- **OLGA III-IV:** RR 32.31 (95% CI 9.14-114.21) vs Stage 0-II
- **OLGIM III-IV:** RR 12.38 (95% CI 5.75-26.65) vs Stage 0-II
- **Absolute Risk Increase:** OLGA III-IV = 4%, OLGIM III-IV = 5%

#### H. pylori Status

| Status | Detection Method | Clinical Action |
|--------|------------------|-----------------|
| **Positive** | Histology, RUT, UBT, stool Ag | Eradication therapy |
| **Negative** | No organisms seen | No treatment |
| **Eradicated** | Post-treatment testing | Follow-up OLGA/OLGIM in 3 years |

**Note:** H. pylori eradication can slow progression but may not reverse atrophy/IM.

#### Database Schema

```json
{
  "finding": "gastritis",
  "sydney_system": {
    "chronic_inflammation": 0-3,
    "neutrophil_activity": 0-3,
    "atrophy": 0-3,
    "intestinal_metaplasia": 0-3,
    "h_pylori_density": 0-3
  },
  "olga_stage": "0|I|II|III|IV",
  "olgim_stage": "0|I|II|III|IV",
  "gastric_cancer_risk": "very_low|low|intermediate|high|very_high",
  "annual_cancer_risk_percent": 0-5,
  "surveillance_interval_years": 1-10,
  "h_pylori_status": "positive|negative|eradicated",
  "eradication_indicated": boolean
}
```

**References:**
- [Sydney System and Gastric Cancer Risk](https://pmc.ncbi.nlm.nih.gov/articles/PMC6014253/)
- [OLGA/OLGIM Meta-Analysis 2025](https://journals.sagepub.com/doi/10.1177/17562848251325461)
- [OLGA Long-term Follow-up Study](https://pubmed.ncbi.nlm.nih.gov/30333540/)

---

### 4. Esophageal Varices: Baveno Classification

**System:** Baveno VII (2024) "Rule of Five"
**Type:** Risk stratification combining LSM + platelet count
**CSV Suitability:** HIGH - Quantitative scoring

#### Variceal Size Grading

| Grade | Definition | Bleeding Risk | Management |
|-------|------------|---------------|------------|
| **Small** | Minimal elevation, straight | Low | β-blocker if decompensated |
| **Medium** | Tortuous, <1/3 lumen | Moderate | β-blocker (carvedilol/nadolol) |
| **Large** | >1/3 lumen occlusion | High | β-blocker + consider EVL |

#### High-Risk Stigmata

- Red wale marks (longitudinal red streaks)
- Cherry-red spots
- Diffuse erythema
- **Any grade 2+ with stigmata = HIGH RISK**

#### Baveno VII Risk Stratification (2024)

**Non-invasive screening criteria to avoid endoscopy:**

| Liver Stiffness (LSM) | Platelet Count | Risk Category | Action |
|----------------------|----------------|---------------|--------|
| **2.5-9.9 kPa** | Any | Very low | No endoscopy needed |
| **10-14.9 kPa** | >150 × 10⁹/L | Low | No endoscopy; can skip screening |
| **15-19.9 kPa** | <110 × 10⁹/L | Moderate | Screen with endoscopy |
| **20-24.9 kPa** | <150 × 10⁹/L | High | Screen with endoscopy |
| **≥25 kPa** | Any | Very high | Mandatory screening |

**>90% NPV for ruling out high-risk varices** with LSM ≤15 kPa + PLT >150

#### Baveno VI Criteria (Still Valid)

**Can safely avoid screening endoscopy if:**
- LSM <20 kPa AND
- Platelet count >150 × 10⁹/L

#### Expanded Baveno VI with Spleen Stiffness

- **Spleen Stiffness Measurement (SSM) ≤40-46 kPa:** Can rule out VNT (varices needing treatment)

#### Database Schema

```json
{
  "finding": "esophageal_varices",
  "grade": "small|medium|large",
  "high_risk_stigmata": boolean,
  "red_wale_marks": boolean,
  "bleeding_risk": "low|moderate|high",
  "baveno_vii_category": "very_low|low|moderate|high|very_high",
  "liver_stiffness_kpa": 0-75,
  "platelet_count": 0-500,
  "spleen_stiffness_kpa": 0-100,
  "screening_indicated": boolean,
  "prophylaxis_indicated": boolean,
  "beta_blocker_recommended": boolean,
  "evl_recommended": boolean
}
```

**References:**
- [Baveno VII Validation Study 2024](https://pmc.ncbi.nlm.nih.gov/articles/PMC10177352/)
- [AASLD Practice Guidance 2024](https://pubmed.ncbi.nlm.nih.gov/37870298/)
- [Baveno VI Performance Meta-Analysis](https://www.cghjournal.org/article/S1542-3565(19)30494-X/fulltext)

---

### 5. Peptic Ulcer Disease

**CSV Suitability:** LOW - Mostly categorical

#### Forrest Classification (Bleeding Risk)

| Grade | Finding | Rebleeding Risk | Management |
|-------|---------|-----------------|------------|
| **Ia** | Spurting arterial bleeding | 90% | Urgent endoscopic hemostasis |
| **Ib** | Oozing bleeding | 50% | Endoscopic hemostasis |
| **IIa** | Visible vessel | 50% | Endoscopic hemostasis |
| **IIb** | Adherent clot | 25% | Consider hemostasis |
| **IIc** | Hematin-covered base | 10% | Medical management |
| **III** | Clean base | 5% | Medical management |

**Recommendation:** Document as categorical finding with automatic risk assignment.

---

## Colonoscopy

### Overview

**Procedure:** Colonoscopy with polypectomy
**Purpose:** Colorectal cancer screening and surveillance
**Quality Metrics:** ADR, SSLDR, BBPS

### 1. Polyp Characteristics

**CSV Suitability:** MEDIUM - Size categories, histology types

#### Polyp Size Classification

| Size Category | Measurement | Cancer Risk | Management |
|---------------|-------------|-------------|------------|
| **Diminutive** | <5 mm | Very low | "resect and discard" strategy |
| **Small** | 5-9 mm | Low | Polypectomy |
| **Large** | 10-19 mm | Moderate | Polypectomy, send to pathology |
| **Very large** | ≥20 mm | High | EMR/ESD, 3-6 month check |

#### Histology Types

| Type | Malignant Potential | Clinical Significance |
|------|---------------------|----------------------|
| **Hyperplastic** | None (unless ≥10mm proximal) | No surveillance |
| **Tubular adenoma** | Low | Standard surveillance |
| **Tubulovillous adenoma** | Moderate | High-risk if ≥10mm |
| **Villous adenoma** | High | High-risk if ≥25% villous |
| **Sessile serrated lesion (SSL)** | Moderate | High-risk if ≥10mm or dysplasia |
| **Traditional serrated adenoma (TSA)** | High | Always high-risk |

#### Advanced Adenoma Criteria

**Definition:** At least ONE of the following:
- Size ≥10 mm OR
- High-grade dysplasia OR
- Villous component ≥25%

---

### 2. Post-Polypectomy Surveillance: USMSTF 2020 Guidelines

**CSV Suitability:** HIGH - Algorithmic interval assignment

#### Surveillance Intervals

| Findings | Risk Category | Interval | Notes |
|----------|---------------|----------|-------|
| **No adenomas** | Average risk | 10 years | Return to screening |
| **1-2 tubular adenomas <10mm** | Low-risk | 7-10 years | Major change from 5-10 years |
| **3-4 tubular adenomas <10mm** | Intermediate | 3-5 years | Changed from 3 years |
| **5-10 adenomas** | High-risk | 3 years | |
| **>10 adenomas** | Very high-risk | 1 year | Consider polyposis syndrome |
| **≥1 adenoma ≥10mm** | High-risk | 3 years | |
| **≥1 adenoma with HGD** | High-risk | 3 years | |
| **≥1 adenoma with ≥25% villous** | High-risk | 3 years | USMSTF; not ESGE |

#### Sessile Serrated Lesions (SSLs)

| Findings | Interval |
|----------|----------|
| **1-2 SSLs <10mm, no dysplasia** | 5-10 years |
| **3-4 SSLs <10mm, no dysplasia** | 3-5 years |
| **5-10 SSLs <10mm, no dysplasia** | 3 years |
| **≥1 SSL ≥10mm** | 3 years |
| **≥1 SSL with dysplasia** | 3 years |

#### Piecemeal Resection (≥20mm)

- **3-6 months:** Early check for residual/recurrence
- **12 months:** First surveillance after clear early check
- **Then:** Standard interval based on total polyp burden

#### Database Schema

```json
{
  "procedure": "colonoscopy",
  "polyps": [
    {
      "number": 1-N,
      "location": "cecum|ascending|transverse|descending|sigmoid|rectum",
      "size_mm": 1-100,
      "size_category": "diminutive|small|large|very_large",
      "histology": "hyperplastic|tubular|tubulovillous|villous|SSL|TSA",
      "dysplasia_grade": "none|low_grade|high_grade",
      "villous_percent": 0-100,
      "is_advanced_adenoma": boolean,
      "resection_method": "biopsy|snare|EMR|ESD|piecemeal"
    }
  ],
  "total_adenomas": 0-N,
  "total_ssls": 0-N,
  "risk_category": "average|low|intermediate|high|very_high",
  "surveillance_interval_years": 1-10,
  "early_check_indicated": boolean,
  "early_check_months": 3-6
}
```

**References:**
- [USMSTF 2020 Guidelines PDF](https://www.asge.org/docs/default-source/guidelines/recommendations-for-follow-up-after-colonoscopy-and-polypectomy-a-consensus-update-by-the-us-multi-society-task-force-on-colorectal-cancer-2020-march-gie.pdf)
- [AGA Clinical Practice Update 2023](https://www.gastrojournal.org/article/S0016-5085(23)04771-6/fulltext)

---

### 3. ESGE 2020 Guidelines (European)

**Key Differences from USMSTF:**

| Finding | ESGE Recommendation | USMSTF Recommendation |
|---------|---------------------|----------------------|
| **1-4 adenomas <10mm, LGD** | No surveillance; return to screening (10 years if no program) | 7-10 years |
| **Villous histology alone** | Does NOT require surveillance | High-risk, 3 years |
| **5-10 adenomas** | 3 years | 3 years (same) |
| **≥1 adenoma ≥10mm** | 3 years | 3 years (same) |

**ESGE is more conservative** (less surveillance) for low-risk adenomas.

**References:**
- [ESGE Guideline 2020 Update](https://www.esge.com/assets/downloads/pdfs/guidelines/2020_a_1185_3109.pdf)

---

### 4. Quality Metrics: ADR & SSLDR

**CSV Suitability:** HIGH - Provider performance tracking

#### Adenoma Detection Rate (ADR)

**Definition:** % of screening colonoscopies with ≥1 adenoma detected

**Benchmarks (2024 ASGE/ACG):**
- **Men:** >30%
- **Women:** >20%
- **Overall:** >25%

**Clinical Impact:**
- Each 1% increase in ADR → 3% decrease in interval CRC risk
- ADR <20% associated with significantly higher interval cancer rates

#### Sessile Serrated Lesion Detection Rate (SSLDR)

**Definition:** % of screening colonoscopies with ≥1 SSL detected

**Benchmarks (2024):**
- **Emerging metric:** >7-10% (varies by guidelines)
- **Priority indicator:** Should be measured by all colonoscopists

#### Database Schema

```json
{
  "provider_quality_metrics": {
    "provider_id": "uuid",
    "time_period": "YYYY-MM to YYYY-MM",
    "total_screening_colonoscopies": 100-N,
    "colonoscopies_with_adenoma": 0-N,
    "adenoma_detection_rate": 0-100,
    "colonoscopies_with_ssl": 0-N,
    "ssl_detection_rate": 0-100,
    "meets_adr_benchmark": boolean,
    "meets_ssldr_benchmark": boolean
  }
}
```

**References:**
- [2024 ACG/ASGE Quality Indicators PDF](http://giquic.org/wp-content/uploads/2024/09/ASGE-ACG-quality-indicators-for-colonoscopy-2024.pdf)
- [Medscape: 19 Quality Indicators](https://www.medscape.com/viewarticle/acg-asge-task-force-identifies-19-indicators-achieving-2024a1000g6c)

---

### 5. Boston Bowel Preparation Scale (BBPS)

**CSV Suitability:** HIGH - Quantitative quality metric

#### BBPS Scoring

| Segment | Score | Definition | Clinical Implication |
|---------|-------|------------|---------------------|
| **Right colon** | 0 | Unprepared, mucosa not seen | Repeat procedure immediately |
| | 1 | Portion of mucosa seen | Inadequate; repeat in 3-6 months |
| | 2 | Minor residue, mucosa seen | Adequate |
| | 3 | Entire mucosa seen | Excellent |
| **Transverse** | 0-3 | Same as above | |
| **Left colon** | 0-3 | Same as above | |

**Total Score:** 0-9 (sum of three segments)

#### Adequacy Thresholds

| Total Score | Segment Scores | Adequacy | Action |
|-------------|----------------|----------|--------|
| **<5** | Any segment <2 | Inadequate | Repeat in <1 year |
| **5** | Debate; some say inadequate | Borderline | Clinical judgment |
| **6** | All segments ≥2, but not optimal | Adequate | Standard intervals |
| **7-9** | Most segments 2-3 | Optimal | Standard intervals |

**Evidence (2024 Study):**
- BBPS 6 has **more missed lesions** than BBPS 7-9
- Segments with BBPS <2 have **higher polyp detection at follow-up** (10% vs 5%)
- **Advanced polyps more common** at follow-up after inadequate prep (20% vs 4%)

#### Database Schema

```json
{
  "bowel_preparation": {
    "bbps_right": 0-3,
    "bbps_transverse": 0-3,
    "bbps_left": 0-3,
    "bbps_total": 0-9,
    "adequacy": "inadequate|borderline|adequate|optimal",
    "repeat_indicated": boolean,
    "repeat_interval_months": 3-12,
    "affects_surveillance": boolean
  }
}
```

**References:**
- [Endoscopy Campus - BBPS](https://www.endoscopy-campus.com/en/classifications/boston-bowel-preparation-scale/)
- [Nature Study: BBPS 6 vs 7-9 Missed Lesions](https://www.nature.com/articles/s41598-024-52244-8)
- [BBPS Validation Study](https://pmc.ncbi.nlm.nih.gov/articles/PMC2763922/)

---

## CSV Stratification Feasibility

### Overall Assessment

| Procedure | Component | CSV Suitability | Data Type | Recommendation |
|-----------|-----------|-----------------|-----------|----------------|
| **EDA** | LA Classification | MEDIUM | Ordinal (1-4) | Convert to numeric for trending |
| | Prague C&M | HIGH | Continuous (cm) | Direct numeric storage |
| | OLGA/OLGIM | HIGH | Stage + risk % | Stage-based risk score |
| | Baveno VII | HIGH | Quantitative formula | LSM + platelet algorithm |
| | Peptic ulcers | LOW | Categorical | Use Forrest classes |
| **Colonoscopy** | Polyp size | MEDIUM | Continuous + categories | Store mm + category |
| | Histology | LOW | Categorical | Dropdown selection |
| | USMSTF intervals | HIGH | Algorithmic | Auto-calculate from findings |
| | ADR/SSLDR | HIGH | Continuous % | Provider performance metric |
| | BBPS | HIGH | Numeric score (0-9) | Quality metric |

### Why NOT Pure CSV Risk Scores?

**Clinical Reasoning:**

1. **Categorical Nature:**
   - Most findings are **discrete grades** (LA A/B/C/D, polyp histology)
   - Not truly continuous variables like lab values

2. **Guideline-Driven:**
   - Management is based on **classification systems**, not continuous risk
   - Surveillance intervals are **categorical decisions** (3 vs 5 vs 10 years)

3. **Quality Standards:**
   - Endoscopy reports require **standardized terminology** (2024 ASGE/ACG mandate)
   - CSV-style scoring would deviate from clinical documentation standards

4. **Multi-Dimensional:**
   - Colonoscopy risk depends on **number + size + histology + dysplasia**
   - Cannot reduce to single continuous score without losing clinical utility

### Recommended Hybrid Approach

**1. Structured Data Entry (Primary):**
- Classification-based dropdowns (LA, Prague, OLGA, BBPS)
- Numeric inputs for measurements (polyp size, Prague C&M)
- Checkboxes for high-risk features

**2. Automatic Risk Calculation (Secondary):**
- OLGA/OLGIM → Annual gastric cancer risk %
- Baveno VII → Portal hypertension risk category
- Polyp burden → USMSTF surveillance interval
- BBPS → Adequacy flag + repeat recommendation

**3. Quality Dashboard (Tertiary):**
- Provider ADR/SSLDR tracking
- Guideline adherence monitoring
- Preparation quality trends

---

## Implementation Recommendations

### Database Schema Design

#### Endoscopy Procedures Table

```sql
CREATE TABLE endoscopy_procedures (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    provider_id UUID NOT NULL REFERENCES users(id),
    procedure_type VARCHAR(50) NOT NULL CHECK (procedure_type IN ('EGD', 'colonoscopy', 'flexible_sigmoidoscopy')),
    procedure_date TIMESTAMP NOT NULL,
    indication TEXT NOT NULL,
    sedation_type VARCHAR(50),

    -- Quality metrics
    procedure_duration_minutes INT,
    photo_documentation_complete BOOLEAN DEFAULT false,

    -- Follow-up
    surveillance_indicated BOOLEAN,
    surveillance_interval_years DECIMAL(3,1),
    next_surveillance_date DATE,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

#### EGD-Specific Findings Table

```sql
CREATE TABLE egd_findings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    procedure_id UUID NOT NULL REFERENCES endoscopy_procedures(id),

    -- Esophagitis
    esophagitis_present BOOLEAN,
    esophagitis_la_grade VARCHAR(1) CHECK (esophagitis_la_grade IN ('A', 'B', 'C', 'D')),
    esophagitis_numeric_value INT CHECK (esophagitis_numeric_value BETWEEN 1 AND 4),

    -- Barrett's
    barretts_present BOOLEAN,
    prague_c_cm DECIMAL(3,1),
    prague_m_cm DECIMAL(3,1),
    barretts_dysplasia VARCHAR(20) CHECK (barretts_dysplasia IN ('none', 'LGD', 'HGD')),
    barretts_segment_type VARCHAR(20) CHECK (barretts_segment_type IN ('short', 'long')),
    annual_progression_risk_percent DECIMAL(4,2),

    -- Gastritis
    gastritis_present BOOLEAN,
    sydney_chronic_inflammation INT CHECK (sydney_chronic_inflammation BETWEEN 0 AND 3),
    sydney_neutrophil_activity INT CHECK (sydney_neutrophil_activity BETWEEN 0 AND 3),
    sydney_atrophy INT CHECK (sydney_atrophy BETWEEN 0 AND 3),
    sydney_intestinal_metaplasia INT CHECK (sydney_intestinal_metaplasia BETWEEN 0 AND 3),
    sydney_h_pylori_density INT CHECK (sydney_h_pylori_density BETWEEN 0 AND 3),
    olga_stage VARCHAR(3) CHECK (olga_stage IN ('0', 'I', 'II', 'III', 'IV')),
    olgim_stage VARCHAR(3) CHECK (olgim_stage IN ('0', 'I', 'II', 'III', 'IV')),
    gastric_cancer_risk_category VARCHAR(20),
    annual_gastric_cancer_risk_percent DECIMAL(4,2),

    -- Varices
    varices_present BOOLEAN,
    varices_grade VARCHAR(20) CHECK (varices_grade IN ('small', 'medium', 'large')),
    varices_high_risk_stigmata BOOLEAN,
    baveno_vii_category VARCHAR(20),
    liver_stiffness_kpa DECIMAL(4,1),
    platelet_count INT,

    -- Ulcers
    ulcer_present BOOLEAN,
    ulcer_location TEXT,
    forrest_classification VARCHAR(3),

    created_at TIMESTAMP DEFAULT NOW()
);
```

#### Colonoscopy Polyps Table

```sql
CREATE TABLE colonoscopy_polyps (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    procedure_id UUID NOT NULL REFERENCES endoscopy_procedures(id),
    polyp_number INT NOT NULL,

    -- Location
    location VARCHAR(50) NOT NULL CHECK (location IN (
        'cecum', 'ascending', 'hepatic_flexure', 'transverse',
        'splenic_flexure', 'descending', 'sigmoid', 'rectum'
    )),

    -- Size
    size_mm INT NOT NULL,
    size_category VARCHAR(20) CHECK (size_category IN ('diminutive', 'small', 'large', 'very_large')),

    -- Histology
    histology VARCHAR(50) CHECK (histology IN (
        'hyperplastic', 'tubular_adenoma', 'tubulovillous_adenoma',
        'villous_adenoma', 'SSL', 'TSA', 'adenocarcinoma'
    )),
    dysplasia_grade VARCHAR(20) CHECK (dysplasia_grade IN ('none', 'low_grade', 'high_grade')),
    villous_percent INT CHECK (villous_percent BETWEEN 0 AND 100),

    -- Risk classification
    is_advanced_adenoma BOOLEAN,
    is_high_risk BOOLEAN,

    -- Resection
    resection_method VARCHAR(50),
    resection_complete BOOLEAN,

    created_at TIMESTAMP DEFAULT NOW()
);
```

#### Colonoscopy Quality Table

```sql
CREATE TABLE colonoscopy_quality (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    procedure_id UUID NOT NULL REFERENCES endoscopy_procedures(id),

    -- Bowel prep
    bbps_right INT NOT NULL CHECK (bbps_right BETWEEN 0 AND 3),
    bbps_transverse INT NOT NULL CHECK (bbps_transverse BETWEEN 0 AND 3),
    bbps_left INT NOT NULL CHECK (bbps_left BETWEEN 0 AND 3),
    bbps_total INT GENERATED ALWAYS AS (bbps_right + bbps_transverse + bbps_left) STORED,
    prep_adequacy VARCHAR(20) CHECK (prep_adequacy IN ('inadequate', 'borderline', 'adequate', 'optimal')),

    -- Procedure quality
    cecal_intubation_achieved BOOLEAN,
    cecal_photo_documented BOOLEAN,
    withdrawal_time_minutes INT,

    -- Detection
    adenomas_detected INT DEFAULT 0,
    ssls_detected INT DEFAULT 0,

    created_at TIMESTAMP DEFAULT NOW()
);
```

#### Risk Calculation Functions

```sql
-- Function: Calculate USMSTF surveillance interval
CREATE OR REPLACE FUNCTION calculate_usmstf_interval(
    p_procedure_id UUID
) RETURNS TABLE (
    risk_category VARCHAR(20),
    interval_years DECIMAL(3,1)
) AS $$
DECLARE
    v_total_adenomas INT;
    v_total_ssls INT;
    v_has_advanced BOOLEAN;
    v_max_size_mm INT;
    v_has_hgd BOOLEAN;
    v_has_high_villous BOOLEAN;
BEGIN
    -- Count polyps and check for high-risk features
    SELECT
        COUNT(*) FILTER (WHERE histology LIKE '%adenoma'),
        COUNT(*) FILTER (WHERE histology IN ('SSL', 'TSA')),
        BOOL_OR(is_advanced_adenoma),
        MAX(size_mm),
        BOOL_OR(dysplasia_grade = 'high_grade'),
        BOOL_OR(villous_percent >= 25)
    INTO
        v_total_adenomas, v_total_ssls, v_has_advanced,
        v_max_size_mm, v_has_hgd, v_has_high_villous
    FROM colonoscopy_polyps
    WHERE procedure_id = p_procedure_id;

    -- Apply USMSTF 2020 logic
    IF v_total_adenomas = 0 AND v_total_ssls = 0 THEN
        RETURN QUERY SELECT 'average_risk'::VARCHAR(20), 10::DECIMAL(3,1);
    ELSIF v_total_adenomas BETWEEN 1 AND 2 AND v_max_size_mm < 10 THEN
        RETURN QUERY SELECT 'low_risk'::VARCHAR(20), 7.5::DECIMAL(3,1); -- 7-10 years
    ELSIF v_total_adenomas BETWEEN 3 AND 4 AND NOT v_has_advanced THEN
        RETURN QUERY SELECT 'intermediate_risk'::VARCHAR(20), 4::DECIMAL(3,1); -- 3-5 years
    ELSIF v_total_adenomas BETWEEN 5 AND 10 THEN
        RETURN QUERY SELECT 'high_risk'::VARCHAR(20), 3::DECIMAL(3,1);
    ELSIF v_total_adenomas > 10 THEN
        RETURN QUERY SELECT 'very_high_risk'::VARCHAR(20), 1::DECIMAL(3,1);
    ELSIF v_has_advanced OR v_has_hgd OR v_has_high_villous THEN
        RETURN QUERY SELECT 'high_risk'::VARCHAR(20), 3::DECIMAL(3,1);
    ELSE
        RETURN QUERY SELECT 'average_risk'::VARCHAR(20), 10::DECIMAL(3,1);
    END IF;
END;
$$ LANGUAGE plpgsql;
```

---

### UI/UX Recommendations

#### 1. Procedure Documentation Interface

**Structured Form Sections:**

```typescript
// EDA Documentation Form
interface EDGForm {
  // Section 1: Esophagus
  esophagitis: {
    present: boolean;
    laGrade?: 'A' | 'B' | 'C' | 'D';
    autoFollowup?: { weeks: 8; indication: 'Grade C/D reassessment' }; // Auto-populated
  };

  barretts: {
    present: boolean;
    pragueC?: number; // cm
    pragueM?: number; // cm
    dysplasia?: 'none' | 'LGD' | 'HGD';
    autoSurveillance?: { years: number; reason: string }; // Auto-calculated
  };

  // Section 2: Stomach
  gastritis: {
    present: boolean;
    sydneyScores?: {
      chronicInflammation: 0 | 1 | 2 | 3;
      neutrophilActivity: 0 | 1 | 2 | 3;
      atrophy: 0 | 1 | 2 | 3;
      intestinalMetaplasia: 0 | 1 | 2 | 3;
      hPyloriDensity: 0 | 1 | 2 | 3;
    };
    olgaStage?: '0' | 'I' | 'II' | 'III' | 'IV';
    olgimStage?: '0' | 'I' | 'II' | 'III' | 'IV';
    autoRisk?: { cancerRiskPercent: number; surveillance: string }; // Auto-calculated
  };

  varices: {
    present: boolean;
    grade?: 'small' | 'medium' | 'large';
    highRiskStigmata?: boolean;
    liverStiffnessKpa?: number;
    plateletCount?: number;
    autoRiskCategory?: string; // Baveno VII auto-calculated
  };
}

// Colonoscopy Documentation Form
interface ColonoscopyForm {
  // Section 1: Bowel Preparation
  bowelPrep: {
    bbpsRight: 0 | 1 | 2 | 3;
    bbpsTransverse: 0 | 1 | 2 | 3;
    bbpsLeft: 0 | 1 | 2 | 3;
    autoTotal?: number; // Auto-sum
    autoAdequacy?: 'inadequate' | 'adequate' | 'optimal'; // Auto-determined
    autoRepeat?: { indicated: boolean; months: number }; // Auto-flagged
  };

  // Section 2: Polyps (Dynamic Array)
  polyps: Array<{
    number: number;
    location: 'cecum' | 'ascending' | 'transverse' | 'descending' | 'sigmoid' | 'rectum';
    sizeMm: number;
    autoSizeCategory?: 'diminutive' | 'small' | 'large' | 'very_large'; // Auto-categorized
    histology?: 'hyperplastic' | 'tubular_adenoma' | 'tubulovillous_adenoma' | 'villous_adenoma' | 'SSL' | 'TSA';
    dysplasiaGrade?: 'none' | 'low_grade' | 'high_grade';
    villousPercent?: number;
    autoAdvancedAdenoma?: boolean; // Auto-flagged
  }>;

  // Section 3: Auto-Generated Surveillance
  autoSurveillance: {
    riskCategory: string; // Auto-calculated from polyps
    intervalYears: number; // Auto-calculated (USMSTF)
    nextDate: Date; // Auto-calculated
    earlyCheckIndicated: boolean; // If piecemeal ≥20mm
  };
}
```

#### 2. Real-Time Validation & Guidance

**Smart Assistance Features:**

- **Auto-Categorization:** Size mm → category, polyp count → risk tier
- **Guideline Hints:** Show relevant USMSTF/ESGE criteria as tooltip
- **Risk Alerts:** Red flag for high-risk findings (HGD, OLGA IV, large varices)
- **Quality Checks:** Warn if ADR below benchmark, BBPS inadequate

#### 3. Surveillance Dashboard

**Provider View:**

```typescript
interface SurveillanceDashboard {
  upcomingProcedures: Array<{
    patientName: string;
    lastProcedureDate: Date;
    nextDueDate: Date;
    indication: string; // "3-year post-polypectomy", "Barrett's long-segment"
    daysUntilDue: number;
    overdue: boolean;
  }>;

  qualityMetrics: {
    adr: { value: number; benchmark: 30; meets: boolean };
    ssldr: { value: number; benchmark: 7; meets: boolean };
    bbpsAdequacy: { value: number; benchmark: 90; meets: boolean };
    withdrawalTime: { avgMinutes: number; benchmark: 6; meets: boolean };
  };

  riskStratification: {
    highRiskPatients: number; // OLGA III-IV, Barrett's HGD, large varices
    dueForSurveillance: number;
    overdueForSurveillance: number;
  };
}
```

---

### Functional Medicine Integration

#### Gut Health Beyond Endoscopy

**Microbiome Context (2024 Research):**

1. **IBD and Dysbiosis:**
   - Adherent-invasive E. coli linked to mucosal dysbiosis and endoscopic recurrence in Crohn's
   - Diagnostic models using 10 bacterial species for UC, 9 for CD (AUC >0.90)

2. **FMT and Endoscopic Remission:**
   - Patients with endoscopic remission post-FMT show enrichment of:
     - *Eubacterium hallii*
     - *Roseburia inulivorans*
   - Increased short-chain fatty acid biosynthesis

3. **Precision Microbiome Therapy:**
   - GUT-108 (11 bacterial strains) reduced inflammatory cytokines in colitis models
   - 2 new FMT-based products introduced in 2024

**Integration Recommendations:**

```json
{
  "endoscopy_finding": "LA Grade C esophagitis",
  "functional_medicine_context": {
    "microbiome_testing_indicated": true,
    "dietary_interventions": [
      "Low-FODMAP trial",
      "Anti-inflammatory diet",
      "Elimination of trigger foods"
    ],
    "supplementation": [
      "Probiotics (Lactobacillus, Bifidobacterium)",
      "Digestive enzymes",
      "L-glutamine for mucosal healing"
    ],
    "lifestyle_factors": [
      "Stress management (vagal tone)",
      "Sleep optimization",
      "Eating pace and meal timing"
    ]
  }
}
```

**References:**
- [Microbiome 2.0: 2024 Gut Microbiota Summit](https://pmc.ncbi.nlm.nih.gov/articles/PMC11404604/)
- [Precision Microbiome Therapy for IBD 2025](https://www.tandfonline.com/doi/full/10.1080/19490976.2025.2489067)
- [Noninvasive Microbiome-Based IBD Diagnosis 2024](https://www.nature.com/articles/s41591-024-03280-4)

---

## Key Takeaways

### 1. Use Structured Classifications, Not Pure CSV Scores

**Rationale:**
- Endoscopic findings are inherently **categorical/ordinal**
- Clinical guidelines are **classification-driven** (LA, OLGA, Prague)
- 2024 ASGE/ACG mandate **standardized terminology**

**Implementation:**
- Dropdown menus for grades/stages
- Numeric inputs for measurements (cm, mm)
- Auto-calculation of risk categories and surveillance intervals

---

### 2. Leverage Algorithmic Risk Stratification

**Where It Works Well:**
- **OLGA/OLGIM:** Stage → gastric cancer risk % → surveillance interval
- **Baveno VII:** LSM + platelets → variceal risk → screening decision
- **USMSTF:** Polyp count/size/histology → CRC risk → surveillance interval
- **BBPS:** Segment scores → adequacy → repeat decision

**Implementation:**
- Write SQL/TypeScript functions for automatic calculation
- Display calculated risk and recommendations in real-time
- Store both raw data and calculated outputs

---

### 3. Track Quality Metrics for Provider Performance

**Mandatory Metrics (2024 ASGE/ACG):**
- **ADR:** >30% men, >20% women
- **SSLDR:** >7-10%
- **BBPS adequacy rate:** >90% with score ≥6
- **Withdrawal time:** ≥6 minutes average

**Implementation:**
- Provider dashboard with benchmark comparison
- Quarterly quality reports
- Alerts for below-benchmark performance

---

### 4. Automate Surveillance Scheduling

**Benefits:**
- Reduce recall errors
- Improve guideline adherence
- Identify overdue patients

**Implementation:**
- Auto-calculate next surveillance date from procedure findings
- Calendar reminders 3 months before due date
- Automated patient outreach for overdue surveillance

---

### 5. Integrate Functional Medicine Context

**Holistic Approach:**
- Link endoscopic inflammation to microbiome dysbiosis
- Recommend functional testing (stool analysis, food sensitivity)
- Provide dietary and lifestyle interventions alongside medical therapy

**Implementation:**
- Conditional recommendations based on findings
- Patient education materials on gut health
- Track functional interventions and outcomes

---

## References

### Esophagogastroduodenoscopy (EDA)

1. [Endoscopy Campus - LA Classification](https://www.endoscopy-campus.com/en/classifications/reflux-esophagitis-los-angeles-classification/)
2. [Clinical Gastroenterology and Hepatology - LA Grades Comparison 2024](https://www.cghjournal.org/article/S1542-3565(24)00480-4/fulltext)
3. [ACG Guideline - Barrett's Esophagus 2023](https://pmc.ncbi.nlm.nih.gov/articles/PMC10259184/)
4. [Sydney System and H. pylori Gastritis](https://pmc.ncbi.nlm.nih.gov/articles/PMC6014253/)
5. [OLGA/OLGIM Meta-Analysis 2025](https://journals.sagepub.com/doi/10.1177/17562848251325461)
6. [OLGA Long-term Follow-up Study - 7436 Patients](https://pubmed.ncbi.nlm.nih.gov/30333540/)
7. [Baveno VII Validation Study 2024](https://pmc.ncbi.nlm.nih.gov/articles/PMC10177352/)
8. [AASLD Practice Guidance on Portal Hypertension 2024](https://pubmed.ncbi.nlm.nih.gov/37870298/)
9. [Updated EGD Quality Indicators 2024](https://www.gastroendonews.com/Endoscopy-Suite/Article/04-25/Updated-EGD-Quality-Indicators-Released/76762)

### Colonoscopy

10. [USMSTF 2020 Guidelines - Follow-Up After Colonoscopy](https://www.asge.org/docs/default-source/guidelines/recommendations-for-follow-up-after-colonoscopy-and-polypectomy-a-consensus-update-by-the-us-multi-society-task-force-on-colorectal-cancer-2020-march-gie.pdf)
11. [USMSTF 2020 Consensus Update - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC7389642/)
12. [AGA Clinical Practice Update 2023 - Risk Stratification for CRC Screening](https://www.gastrojournal.org/article/S0016-5085(23)04771-6/fulltext)
13. [ESGE Post-Polypectomy Surveillance Guideline 2020 Update](https://www.esge.com/assets/downloads/pdfs/guidelines/2020_a_1185_3109.pdf)
14. [ASGE/ACG Quality Indicators for Colonoscopy 2024](http://giquic.org/wp-content/uploads/2024/09/ASGE-ACG-quality-indicators-for-colonoscopy-2024.pdf)
15. [Medscape: 19 Quality Indicators for GI Endoscopy 2024](https://www.medscape.com/viewarticle/acg-asge-task-force-identifies-19-indicators-achieving-2024a1000g6c)

### Quality Metrics

16. [Boston Bowel Preparation Scale - Endoscopy Campus](https://www.endoscopy-campus.com/en/classifications/boston-bowel-preparation-scale/)
17. [Nature Study: BBPS 6 vs 7-9 Missed Lesions 2024](https://www.nature.com/articles/s41598-024-52244-8)
18. [BBPS Validation Study - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC2763922/)
19. [Quality Indicators Common to All GI Endoscopy 2024](https://www.giejournal.org/article/S0016-5107(24)03183-3/fulltext)

### Functional Medicine & Microbiome

20. [Microbiome 2.0: 2024 Gut Microbiota for Health World Summit](https://pmc.ncbi.nlm.nih.gov/articles/PMC11404604/)
21. [Precision Microbiome Therapy for IBD 2025](https://www.tandfonline.com/doi/full/10.1080/19490976.2025.2489067)
22. [Noninvasive Microbiome-Based IBD Diagnosis 2024](https://www.nature.com/articles/s41591-024-03280-4)
23. [Key Advances in Gut Microbiome During 2024](https://www.gutmicrobiotaforhealth.com/key-advances-in-the-gut-microbiome-during-2024/)
24. [International Consensus on Microbiome Testing 2024](https://www.thelancet.com/journals/langas/article/PIIS2468-1253(24)00311-X/abstract)

---

**Document Status:** Complete
**Next Review:** July 2026 (for new guideline updates)
**Contact:** EMR Development Team
