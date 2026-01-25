# Imaging Exams for Risk Stratification - Technical Documentation

**Document Version:** 1.0
**Last Updated:** January 19, 2026
**Target System:** Plenya EMR - Laboratory Results Module

---

## Table of Contents

1. [Executive Summary](#executive-summary)
2. [Mamografia Digital Bilateral](#1-mamografia-digital-bilateral)
3. [USG Mamas (Breast Ultrasound)](#2-usg-mamas-breast-ultrasound)
4. [USG Transvaginal](#3-usg-transvaginal)
5. [CSV Data Structure Recommendations](#csv-data-structure-recommendations)
6. [Clinical Decision Support Algorithms](#clinical-decision-support-algorithms)
7. [Database Schema Design](#database-schema-design)
8. [References](#references)

---

## Executive Summary

### CSV Suitability Assessment

| Exam | Quantitative Components | Semi-Quantitative Components | Qualitative Components | CSV Suitability |
|------|------------------------|----------------------------|----------------------|----------------|
| **Mamografia Digital** | None | BI-RADS (0-6), Density (A-D) | Findings description | **PARTIALLY** - Ordinal categories |
| **USG Mamas** | None | BI-RADS (0-6) | Findings description | **PARTIALLY** - Ordinal categories |
| **USG Transvaginal** | Endometrial thickness (mm), Ovarian volume (cm³) | O-RADS (1-5), IOTA Simple Rules | Mass characteristics | **YES** - Mixed quantitative/ordinal |

### Key Findings

**BI-RADS as Ordinal Variable:**
- BI-RADS categories (0-6) represent **ordinal risk stratification** with specific cancer probability ranges
- Suitable for CSV storage as integer values with mapped risk percentages
- Enable automated clinical decision support and risk tracking over time

**Quantitative Measurements:**
- Endometrial thickness (mm) - Fully quantitative, threshold-based risk stratification
- Ovarian volume (cm³) - Quantitative measurement for PCOS diagnosis
- Suitable for statistical analysis, trending, and automated alerts

**Recommended Approach:**
- Store structured data in relational database (PostgreSQL)
- Generate CSV exports for analytics and risk scoring calculations
- Implement automated risk stratification algorithms based on evidence-based thresholds

---

## 1. Mamografia Digital Bilateral

### Overview

Digital mammography is the primary screening tool for breast cancer detection, using low-dose X-rays to create detailed images of breast tissue. The standardized BI-RADS (Breast Imaging Reporting and Data System) classification provides consistent risk stratification.

### 1.1 BI-RADS Classification System

**ACR BI-RADS 5th Edition (2013) - Current Standard**

| Category | Description | Cancer Risk | Management | CSV Value |
|----------|-------------|-------------|------------|-----------|
| **0** | Incomplete - Need Additional Imaging | N/A | Recall for additional views/ultrasound | 0 |
| **1** | Negative | 0.1% | Routine screening (annual/biennial) | 1 |
| **2** | Benign Finding | 0.1% | Routine screening (annual/biennial) | 2 |
| **3** | Probably Benign | <2% | Short-term follow-up (6 months) | 3 |
| **4A** | Low Suspicion | 2-10% | Biopsy recommended | 4 |
| **4B** | Moderate Suspicion | 10-50% | Biopsy recommended | 5 |
| **4C** | High Suspicion | 50-95% | Biopsy recommended | 6 |
| **5** | Highly Suggestive of Malignancy | ≥95% | Biopsy required + oncology referral | 7 |
| **6** | Known Biopsy-Proven Malignancy | 100% | Treatment planning | 8 |

**Clinical Decision Rules:**
- BI-RADS 0: Incomplete study, requires additional imaging within 30 days
- BI-RADS 1-2: Return to routine screening schedule
- BI-RADS 3: 6-month follow-up mammogram (bilateral)
- BI-RADS 4-5: Tissue diagnosis required (biopsy)
- BI-RADS 6: Active treatment monitoring

### 1.2 Breast Density Classification

**ACR BI-RADS Breast Density Categories**

| Category | Description | Fibroglandular Tissue | Cancer Risk Multiplier | Mammography Sensitivity | CSV Value |
|----------|-------------|---------------------|----------------------|------------------------|-----------|
| **A** | Almost entirely fatty | <25% | 1.0x (baseline) | High (>90%) | 1 |
| **B** | Scattered fibroglandular | 25-50% | 1.2x | High (85-90%) | 2 |
| **C** | Heterogeneously dense | 51-75% | 1.5-2.0x | Moderate (70-85%) | 3 |
| **D** | Extremely dense | >75% | 2.0-4.0x | Low (<70%) | 4 |

**Clinical Significance:**
- **Categories C and D ("Dense Breasts"):** 40% of screening-age women
- **Cancer Risk:** 4-6 fold increased risk with Category D vs Category A
- **Supplemental Screening:** Consider ultrasound or MRI for Category C/D with additional risk factors
- **FDA Mandate (2024):** All mammogram reports must include breast density notification

**Supplemental Screening Recommendations (ACR 2024):**
- Category C + elevated risk: Consider MRI or ultrasound
- Category D: Strong consideration for supplemental screening
- BRCA+ or >20% lifetime risk: Annual MRI regardless of density

### 1.3 Screening Guidelines Comparison

#### USPSTF Guidelines (Updated April 2024)

**Key Recommendations (Grade B):**
- **Age to start:** 40 years (changed from 50 years)
- **Screening interval:** Biennial (every 2 years)
- **Age to stop:** 74 years
- **Age 75+:** Insufficient evidence (individualized decision)

**Rationale for Change:**
- 19% more lives saved by starting at age 40
- Addresses racial disparities (Black women 40% higher mortality)
- Breast cancer incidence increasing 2%/year in women 40-49

**Insufficient Evidence:**
- Supplemental screening for dense breasts (ultrasound/MRI)
- Optimal screening interval for high-risk women
- Digital breast tomosynthesis vs standard mammography (cost-effectiveness)

#### ACR Guidelines (2023-2024)

**More Aggressive than USPSTF:**
- **Age to start:** 40 years (average risk)
- **Screening interval:** Annual (every year)
- **High-risk screening:** Begin at age 25-30 with MRI + mammography

#### Brazilian Society of Mastology (SBM/CBR/FEBRASGO 2023)

**Aligned with ACR:**
- **Age to start:** 40 years
- **Screening interval:** Annual mammography
- **Technology:** Preferably digital (Category A recommendation)
- **Age to stop:** 74 years (continue if life expectancy >7 years)

**Higher-Risk Groups Requiring Enhanced Screening:**
- Dense breasts (BI-RADS C/D)
- Personal history of atypical hyperplasia or LCIS
- Prior breast cancer treatment
- Chest irradiation before age 30

### 1.4 Digital Breast Tomosynthesis (DBT/3D Mammography)

**Technology Overview:**
- Multiple X-ray images from different angles
- Computer reconstruction into 3D image set
- Reduces tissue overlap compared to 2D mammography

**2023-2024 Evidence:**

| Metric | 2D Digital Mammography | 3D Tomosynthesis (DBT) | Improvement |
|--------|----------------------|----------------------|-------------|
| Cancer Detection Rate | Baseline | +15-20% | Higher |
| Recall Rate | 10-12% | 7-9% | 25-30% reduction |
| Interval Cancer Rate | Same | Same | No difference |
| Advanced Cancer Rate | Higher | Lower | Significant reduction |
| Radiologist Performance | 27.9% meet benchmarks | 34.2% meet benchmarks | +23% |

**10-Year Study Results (RSNA 2024):**
- Significantly reduces advanced cancer rates
- Sustained benefits across multiple screening rounds
- Currently available at 86% of US certified facilities (March 2023)

**Cost-Effectiveness Considerations:**
- Higher equipment cost (~$150,000-$400,000)
- Longer reading time for radiologists
- May not be reimbursed by all insurers
- Not yet universally adopted in Brazil

**Recommendation for Plenya EMR:**
- Track modality type (2D vs 3D) in database
- Flag 3D availability in facility settings
- Include in quality metrics reporting

### 1.5 Functional Medicine Perspective

**Root Cause Approach to Breast Health:**

**Hormonal Balance:**
- Estrogen dominance assessment (estradiol, progesterone ratio)
- Xenoestrogen exposure reduction (plastics, pesticides)
- Support estrogen metabolism (cruciferous vegetables, DIM, I3C)

**Inflammation Reduction:**
- Omega-3 fatty acids (EPA/DHA 2-3g daily)
- Anti-inflammatory diet (Mediterranean pattern)
- Gut microbiome optimization (estrobolome)

**Metabolic Health:**
- Insulin resistance correction (primary driver)
- Body composition optimization (reduce adiposity)
- Regular physical activity (150 min/week moderate intensity)

**Nutritional Support:**
- Vitamin D3 (target 40-60 ng/mL)
- Iodine sufficiency (breast tissue concentrates iodine)
- Selenium (200 mcg daily)
- Green tea (EGCG 400-600 mg daily)

**Stress Management:**
- Cortisol dysregulation linked to breast cancer risk
- Sleep optimization (7-9 hours nightly)
- Mindfulness practices

---

## 2. USG Mamas (Breast Ultrasound)

### Overview

Breast ultrasound uses sound waves to create images of breast tissue. It serves as a complementary modality to mammography, particularly valuable for dense breasts, palpable masses, and young women.

### 2.1 BI-RADS Classification for Ultrasound

**Same 0-6 scale as mammography** (see Section 1.1)

**Ultrasound-Specific Applications:**
- Characterization of masses (solid vs cystic)
- Evaluation of palpable abnormalities
- Dense breast supplemental screening
- Pregnancy/lactation (no radiation)
- Guidance for biopsy procedures

### 2.2 Clinical Indications

| Indication | Primary Use | Clinical Scenario |
|------------|------------|------------------|
| **Supplemental Screening** | Dense breasts (BI-RADS C/D) | Adjunct to mammography |
| **Diagnostic Workup** | Mammography callback | Characterize BI-RADS 0 findings |
| **Palpable Mass** | Clinical examination finding | First-line in women <30 years |
| **Pregnancy/Lactation** | Symptomatic patient | Avoid radiation exposure |
| **Biopsy Guidance** | Tissue sampling | Real-time needle visualization |

### 2.3 Automated Breast Ultrasound (ABUS)

**Technology:**
- Standardized automated scanning
- Operator-independent acquisition
- 3D volumetric imaging
- FDA/Europe approved for screening

**2024 Evidence:**

| Study | Finding | Clinical Impact |
|-------|---------|-----------------|
| **Cancer Detection** | +1.9-7.7 per 1000 women | Incremental benefit over mammography |
| **Sensitivity Increase** | +21.6-41.0% | Better detection in dense tissue |
| **5-Year Experience** | 2.4 additional cancers per 1000 exams | Academic center data |
| **Recall Rate** | 5.2% | Lower than handheld ultrasound |
| **False Positives** | Substantially increased | Main limitation |

**Clinical Application:**
- Approved for asymptomatic women with dense breasts
- Supplement to mammography (not replacement)
- All malignant lesions detected (high sensitivity)
- Workflow efficiency vs handheld ultrasound

**Limitations:**
- Higher non-cancer recall rates
- Requires specialized equipment
- Not yet widely available in Brazil
- Cost-effectiveness debated

### 2.4 Combined Mammography + Ultrasound Performance

**Synergistic Effect:**
- Combined sensitivity: ~95%
- Mammography: Excellent for calcifications
- Ultrasound: Superior for solid masses in dense tissue
- Complementary modalities, not competing

**Evidence-Based Protocol:**
1. Mammography (BI-RADS C/D with dense breasts)
2. If negative + elevated risk → Add ultrasound
3. If suspicious findings → Diagnostic ultrasound
4. If high risk (>20% lifetime) → Consider MRI

### 2.5 Functional Medicine Integration

**Prevention Strategy:**
- Regular thermography (controversial, not evidence-based screening)
- Lymphatic drainage support (dry brushing, rebounding)
- Reduce underwire bra use (theoretical lymphatic obstruction)
- Castor oil packs (traditional support, limited evidence)

---

## 3. USG Transvaginal

### Overview

Transvaginal ultrasound is the primary imaging modality for evaluating the uterus, ovaries, and adnexal structures. It provides both quantitative measurements and qualitative assessments critical for gynecologic risk stratification.

### 3.1 Endometrial Thickness Assessment

#### Quantitative Measurement (mm)

**Normal Values by Reproductive Status:**

| Status | Phase/Condition | Normal Thickness | Clinical Action Threshold | CSV Column |
|--------|----------------|-----------------|-------------------------|------------|
| **Premenopausal** | Early follicular (days 1-6) | 4-8 mm | N/A | `endometrial_thickness_mm` |
| | Late follicular (days 7-14) | 8-14 mm | N/A | |
| | Secretory (days 15-28) | 10-16 mm | N/A | |
| **Postmenopausal** | No HRT | <5 mm (optimal <4 mm) | >5 mm | `endometrial_thickness_mm` |
| | On HRT | <8 mm | >8 mm | |
| | Tamoxifen therapy | <8 mm | >8 mm | |

**Clinical Decision Thresholds:**

**Postmenopausal WITHOUT Bleeding:**
- ≤11 mm: No further workup required (0.002% cancer risk)
- >11 mm: Consider biopsy (6.7% cancer risk)

**Postmenopausal WITH Bleeding:**
- ≤4 mm: <0.07% cancer risk (negative predictive value >99%)
- ≤5 mm: <0.07% cancer risk (most commonly used threshold)
- >5 mm: 7.3% cancer risk → Endometrial biopsy required

**2024 Guidelines (ACOG, International Consensus):**
- Symptomatic postmenopausal: 4-5 mm cutoff
- Asymptomatic postmenopausal: 11 mm cutoff
- These thresholds remain standard in 2024-2026

#### Automated Risk Stratification Algorithm

```python
def stratify_endometrial_risk(thickness_mm, postmenopausal, has_bleeding, on_hrt, on_tamoxifen):
    """
    Returns: (risk_level, recommendation, cancer_probability)
    """
    if not postmenopausal:
        return ("normal_premenopausal", "Cycle-dependent variation", None)

    # Postmenopausal WITH bleeding
    if has_bleeding:
        if thickness_mm <= 4:
            return ("low", "Routine surveillance", "<0.07%")
        elif thickness_mm <= 5:
            return ("low-moderate", "Consider biopsy based on clinical context", "~0.5%")
        else:
            return ("moderate-high", "Endometrial biopsy required", "7.3%")

    # Postmenopausal WITHOUT bleeding
    else:
        if on_tamoxifen:
            threshold = 8
        elif on_hrt:
            threshold = 8
        else:
            threshold = 11

        if thickness_mm <= threshold:
            return ("low", "Routine surveillance", "0.002%")
        else:
            return ("elevated", "Consider biopsy or close surveillance", "6.7%")
```

### 3.2 Ovarian Assessment - O-RADS Classification

**O-RADS US (Ovarian-Adnexal Reporting and Data System Ultrasound)**

Developed by ACR in 2020, O-RADS provides standardized risk stratification for adnexal masses.

#### Classification System

| Category | Description | Malignancy Risk | Management | CSV Value |
|----------|-------------|----------------|------------|-----------|
| **O-RADS 0** | Incomplete Evaluation | N/A | Complete exam or use MRI | 0 |
| **O-RADS 1** | Physiologic/Normal | 0% | Routine surveillance | 1 |
| **O-RADS 2** | Almost Certainly Benign | <1% | Annual follow-up or discharge | 2 |
| **O-RADS 3** | Low Risk | 1-<10% | Annual follow-up | 3 |
| **O-RADS 4** | Intermediate Risk | 10-<50% | Surgical consultation, consider surgery | 4 |
| **O-RADS 5** | High Risk | ≥50% | Gynecologic oncology referral + surgery | 5 |

**Diagnostic Performance (Meta-Analysis):**
- **Sensitivity:** 97% (95% CI: 94-98%)
- **Specificity:** 77% (95% CI: 68-84%)
- **Positive Likelihood Ratio:** 4.2
- **Negative Likelihood Ratio:** 0.04
- **Diagnostic Odds Ratio:** 96

**Clinical Decision Algorithm:**

```
O-RADS 1: Normal ovary
├── Premenopausal: Follicles/corpus luteum normal
└── Postmenopausal: Small ovaries (<10 cm³) normal

O-RADS 2: Simple cysts
├── Premenopausal: <10 cm simple cyst
├── Postmenopausal: <3 cm simple cyst
└── Other benign lesions (teratoma, endometrioma)

O-RADS 3: Low risk features
├── Smooth walls, thin septa (<3mm)
├── No solid components or papillations
└── No ascites

O-RADS 4: Intermediate risk
├── Thick septa (≥3mm) or nodularity
├── Moderate vascularity
└── Some solid components

O-RADS 5: High risk
├── Large solid components
├── High vascularity
├── Ascites present
└── Peritoneal implants
```

### 3.3 IOTA Simple Rules (Alternative Classification)

**International Ovarian Tumor Analysis Group**

Alternative to O-RADS, uses specific ultrasound features:

**5 Benign (B) Features:**
1. Unilocular cyst
2. Solid components <7 mm
3. Acoustic shadowing
4. Smooth multilocular cyst <100 mm
5. No blood flow (Color Doppler)

**5 Malignant (M) Features:**
1. Irregular solid tumor
2. Ascites
3. ≥4 papillary structures
4. Irregular multilocular-solid cyst ≥100 mm
5. Very strong blood flow (Color Doppler)

**Classification Logic:**
- **Benign:** ≥1 B-feature AND 0 M-features
- **Malignant:** ≥1 M-feature AND 0 B-features
- **Indeterminate:** Mixed features or no features

**Performance (2024 Validation):**
- **Applicable:** 75-94% of masses
- **Sensitivity:** 84.8-93%
- **Specificity:** 95-98.9%
- **PPV:** 87.5%
- **NPV:** 98.6%

**Advantage:** Can be used by non-expert sonographers after training

### 3.4 Ovarian Volume and PCOS Diagnosis

**Quantitative Measurement:**

| Parameter | Normal | PCOS Threshold | CSV Column |
|-----------|--------|----------------|------------|
| **Ovarian Volume** | 3-10 cm³ | ≥10 cm³ | `ovarian_volume_cm3` |
| **Follicle Count** | Variable | ≥12 follicles 2-9mm (FNPO) | `follicle_count` |

**Rotterdam Criteria for PCOS (2 of 3 required):**
1. Oligo/anovulation
2. Hyperandrogenism (clinical or biochemical)
3. Polycystic ovarian morphology (PCOM):
   - ≥12 follicles 2-9 mm in diameter, OR
   - Ovarian volume ≥10 cm³

**2023 Updated Guidelines:**
- FNPO (Follicle Number Per Ovary) replaces older AFC
- Uses transducers with frequency ≥8 MHz
- Measurement in early follicular phase (days 2-5)

### 3.5 Other Structures Assessed

#### Uterine Findings

| Finding | Description | Clinical Significance | CSV Storage |
|---------|-------------|---------------------|-------------|
| **Myomas** | FIGO classification (Types 0-8) | Fertility impact, bleeding | `myoma_figo_type` |
| **Adenomyosis** | Diffuse vs focal | Dysmenorrhea, infertility | `adenomyosis_present` |
| **Congenital anomalies** | Septate, bicornuate, etc. | Pregnancy outcomes | `uterine_anomaly` |

#### Adnexal Findings

| Finding | Description | Clinical Action | CSV Storage |
|---------|-------------|----------------|-------------|
| **Hydrosalpinx** | Fluid-filled fallopian tube | IVF outcomes, infection risk | `hydrosalpinx_present` |
| **Endometrioma** | Ovarian endometriosis | Pain, infertility, surgery | `endometrioma_present` |
| **Paraovarian cyst** | Separate from ovary | Usually benign, observation | `paraovarian_cyst` |

### 3.6 Functional Medicine Approach to Gynecologic Health

#### PCOS Management

**Root Cause Factors:**
- Insulin resistance (primary driver in 70%)
- Chronic inflammation
- Adrenal dysfunction
- Gut dysbiosis

**Therapeutic Interventions:**

**1. Insulin Sensitization:**
- Low-glycemic Mediterranean diet
- Inositol supplementation (Myo-inositol 2-4g + D-chiro-inositol 50-100mg daily)
- Berberine 500mg TID or Metformin
- Regular exercise (resistance + HIIT)

**2. Anti-Inflammatory:**
- Omega-3 EPA/DHA 2-3g daily
- Curcumin 500-1000mg daily
- Eliminate processed foods, refined sugars

**3. Hormone Balance:**
- Vitex (chasteberry) 400mg daily for prolactin modulation
- Saw palmetto 320mg daily for androgen blocking
- Spearmint tea (2 cups daily) reduces hirsutism
- N-acetylcysteine (NAC) 600mg BID

**4. Metabolic Support:**
- Chromium picolinate 200-500 mcg daily
- Alpha-lipoic acid 600mg daily
- Vitamin D3 optimization (target 50-70 ng/mL)
- Magnesium glycinate 400-600mg daily

**5. Gut Health:**
- Probiotic diversity (>25 billion CFU)
- Prebiotics (inulin, resistant starch)
- Address SIBO if present
- Optimize digestive enzymes

#### Endometriosis Support

**Anti-Inflammatory Protocol:**
- Eliminate dairy, gluten, soy (potential triggers)
- Increase cruciferous vegetables (DIM 200mg)
- Resveratrol 500mg daily
- Pycnogenol 30-60mg daily

**Hormonal Modulation:**
- Support estrogen metabolism (Phase I & II detox)
- Calcium-D-glucarate 500mg TID
- Choline and betaine support methylation

**Pain Management:**
- Acupuncture (evidence-based for dysmenorrhea)
- Pelvic floor physical therapy
- Transcutaneous electrical nerve stimulation (TENS)

---

## 4. CSV Data Structure Recommendations

### 4.1 Mamografia Digital CSV Schema

```csv
patient_id,exam_date,exam_type,facility_id,radiologist_id,birads_category,birads_numeric,cancer_risk_percent,breast_density,breast_density_numeric,modality,findings_description,calcifications_present,mass_present,asymmetry_present,architectural_distortion,recommendation,follow_up_months,biopsy_recommended,created_at

12345,2024-01-15,screening_mammogram,FAC001,RAD001,4B,5,30.0,heterogeneously_dense,3,3D_tomosynthesis,"Irregular mass 1.2cm upper outer left breast with suspicious calcifications",true,true,false,false,biopsy_required,0,true,2024-01-15T14:30:00Z
```

**Key Fields:**
- `birads_numeric`: Integer 0-8 for statistical analysis (maps 4A→4, 4B→5, 4C→6)
- `cancer_risk_percent`: Numeric risk for trending
- `breast_density_numeric`: Integer 1-4 for correlation analysis
- `modality`: Track 2D vs 3D technology

### 4.2 USG Mamas CSV Schema

```csv
patient_id,exam_date,exam_type,indication,birads_category,birads_numeric,cancer_risk_percent,mass_characteristics,solid_component,vascular_pattern,elastography_score,findings_description,correlate_with_mammogram,recommendation,biopsy_recommended,created_at

12345,2024-01-16,diagnostic_ultrasound,mammogram_callback,4A,4,6.0,irregular_margins,true,moderate_vascularity,4,"Solid hypoechoic mass 1.1cm corresponding to mammographic finding",true,biopsy_recommended,true,2024-01-16T10:15:00Z
```

### 4.3 USG Transvaginal CSV Schema (Most Complex)

```csv
patient_id,exam_date,indication,postmenopausal,has_bleeding,on_hrt,on_tamoxifen,endometrial_thickness_mm,endometrial_risk_level,endometrial_recommendation,right_ovary_volume_cm3,left_ovary_volume_cm3,right_ovary_orads,left_ovary_orads,right_ovary_risk_percent,left_ovary_risk_percent,follicle_count_right,follicle_count_left,pcos_criteria_met,myoma_present,myoma_figo_type,adenomyosis_present,free_fluid,findings_description,recommendation,referral_needed,created_at

12345,2024-01-20,postmenopausal_bleeding,true,true,false,false,7.5,moderate-high,"Endometrial biopsy required",2.8,3.1,2,2,0.5,0.5,0,0,false,false,NULL,false,false,"Thickened endometrium 7.5mm in symptomatic postmenopausal patient. Normal ovaries.",endometrial_biopsy,gynecology,2024-01-20T09:00:00Z

23456,2024-01-21,pcos_evaluation,false,false,false,false,8.0,normal_premenopausal,"Routine surveillance",12.5,11.8,1,1,0,0,15,14,true,false,NULL,false,false,"Bilateral enlarged ovaries with increased follicle count consistent with PCOS. Endometrium appropriate for cycle day 5.",lifestyle_management,endocrinology,2024-01-21T11:30:00Z
```

**Key Features:**
- Separate risk stratification for endometrium and each ovary
- Boolean flags for clinical context (postmenopausal, bleeding, medications)
- Quantitative measurements (thickness, volume) for trending
- PCOS criteria assessment
- Automated recommendation field

---

## 5. Clinical Decision Support Algorithms

### 5.1 Mammography Risk Stratification

```python
def mammography_risk_stratification(birads, density, age, family_history, prior_biopsy):
    """
    Comprehensive risk assessment combining imaging and clinical factors
    """
    # Base risk from BI-RADS
    base_risk = {
        0: None,  # Incomplete
        1: 0.1,
        2: 0.1,
        3: 2.0,
        4: 6.0,   # Average of 4A (6%) for initial categorization
        5: 30.0,  # Average of 4B
        6: 72.5,  # Average of 4C
        7: 95.0,
        8: 100.0
    }

    if birads not in base_risk or base_risk[birads] is None:
        return {"error": "Incomplete exam or invalid BI-RADS"}

    risk = base_risk[birads]

    # Density multiplier for future cancer risk (not current exam)
    density_multiplier = {1: 1.0, 2: 1.2, 3: 1.6, 4: 2.5}

    # Family history adjustment
    if family_history:
        risk *= 1.5

    # Prior atypical biopsy
    if prior_biopsy == "atypical":
        risk *= 2.0

    # Recommendations
    if birads in [0]:
        recommendation = "Complete imaging workup"
    elif birads in [1, 2]:
        if age >= 40 and age <= 74:
            recommendation = "Annual screening (SBM/ACR) or Biennial (USPSTF)"
        else:
            recommendation = "Individualized screening plan"
    elif birads == 3:
        recommendation = "6-month bilateral follow-up mammogram"
    elif birads in [4, 5, 6]:
        recommendation = "Tissue diagnosis (biopsy) required"
    elif birads == 7:
        recommendation = "Immediate biopsy + oncology referral"
    elif birads == 8:
        recommendation = "Treatment planning and monitoring"

    # Supplemental screening for dense breasts
    supplemental = None
    if density in [3, 4] and birads in [1, 2]:
        supplemental = "Consider supplemental ultrasound or MRI based on risk factors"

    return {
        "current_exam_risk_percent": risk,
        "density_category": density,
        "recommendation": recommendation,
        "supplemental_screening": supplemental,
        "follow_up_months": 6 if birads == 3 else (12 if birads in [1,2] else 0)
    }
```

### 5.2 Transvaginal Ultrasound - Endometrial Algorithm

```python
def endometrial_risk_algorithm(thickness_mm, postmenopausal, has_bleeding, on_hrt, on_tamoxifen):
    """
    Evidence-based endometrial thickness risk stratification
    Returns dict with risk level, cancer probability, and recommendation
    """
    # Premenopausal - cycle-dependent, not risk-stratified by thickness alone
    if not postmenopausal:
        if thickness_mm > 16:
            return {
                "risk_level": "investigate",
                "cancer_probability": "N/A",
                "recommendation": "Rule out endometrial pathology - consider biopsy if abnormal bleeding",
                "requires_action": True
            }
        else:
            return {
                "risk_level": "normal",
                "cancer_probability": "N/A",
                "recommendation": "Normal cycle-dependent variation",
                "requires_action": False
            }

    # Postmenopausal WITH bleeding
    if has_bleeding:
        if thickness_mm <= 4:
            return {
                "risk_level": "low",
                "cancer_probability": "<0.07%",
                "recommendation": "Routine surveillance, biopsy may be deferred",
                "requires_action": False
            }
        elif thickness_mm <= 5:
            return {
                "risk_level": "low-moderate",
                "cancer_probability": "~0.5%",
                "recommendation": "Consider endometrial biopsy based on clinical context",
                "requires_action": True
            }
        else:
            return {
                "risk_level": "moderate-high",
                "cancer_probability": "7.3%",
                "recommendation": "Endometrial biopsy required",
                "requires_action": True
            }

    # Postmenopausal WITHOUT bleeding
    else:
        # Determine threshold based on medication use
        if on_tamoxifen or on_hrt:
            threshold = 8
        else:
            threshold = 11

        if thickness_mm <= threshold:
            return {
                "risk_level": "low",
                "cancer_probability": "0.002%",
                "recommendation": "Routine surveillance, no immediate action needed",
                "requires_action": False
            }
        else:
            return {
                "risk_level": "elevated",
                "cancer_probability": "6.7%",
                "recommendation": f"Endometrial thickness >{threshold}mm - consider biopsy or close surveillance",
                "requires_action": True
            }
```

### 5.3 Ovarian Mass Risk Stratification (O-RADS)

```python
def ovarian_risk_stratification(orads_category, postmenopausal, ca125_level=None, age=None):
    """
    O-RADS classification with management recommendations
    Optionally incorporates CA-125 tumor marker
    """
    risk_data = {
        0: {"risk": "N/A", "description": "Incomplete evaluation", "action": "Complete ultrasound or obtain MRI"},
        1: {"risk": "0%", "description": "Physiologic/normal", "action": "Routine surveillance"},
        2: {"risk": "<1%", "description": "Almost certainly benign", "action": "Annual follow-up or discharge"},
        3: {"risk": "1-10%", "description": "Low risk", "action": "Annual ultrasound follow-up"},
        4: {"risk": "10-50%", "description": "Intermediate risk", "action": "Surgical consultation - consider surgery"},
        5: {"risk": "≥50%", "description": "High risk", "action": "Gynecologic oncology referral + surgery"}
    }

    if orads_category not in risk_data:
        return {"error": "Invalid O-RADS category"}

    result = risk_data[orads_category].copy()

    # Risk of Malignancy Index (RMI) if CA-125 available
    if ca125_level and postmenopausal is not None:
        u_score = 3 if orads_category >= 4 else 1  # Simplified ultrasound score
        m_score = 3 if postmenopausal else 1
        rmi = u_score * m_score * ca125_level

        result["rmi_score"] = rmi
        result["rmi_interpretation"] = "High risk (>200)" if rmi > 200 else "Low-moderate risk"

    # Urgency flag
    result["urgent_referral"] = orads_category >= 4

    return result
```

### 5.4 Integrated Women's Health Risk Score

```python
def womens_health_composite_score(patient_data):
    """
    Composite risk scoring across multiple imaging modalities
    Useful for population health management and prevention programs
    """
    score = 0
    alerts = []

    # Breast cancer risk factors
    if patient_data.get("birads_numeric", 0) >= 4:
        score += 10
        alerts.append("Suspicious breast finding - biopsy needed")
    elif patient_data.get("breast_density") in [3, 4]:
        score += 2
        alerts.append("Dense breasts - consider supplemental screening")

    # Endometrial risk
    if patient_data.get("endometrial_requires_action"):
        score += 8
        alerts.append("Endometrial thickness warrants biopsy")

    # Ovarian risk
    max_orads = max(
        patient_data.get("right_ovary_orads", 1),
        patient_data.get("left_ovary_orads", 1)
    )
    if max_orads >= 4:
        score += 9
        alerts.append("Ovarian mass requires surgical evaluation")
    elif max_orads == 3:
        score += 3
        alerts.append("Low-risk ovarian finding - annual follow-up")

    # PCOS metabolic risk
    if patient_data.get("pcos_criteria_met"):
        score += 3
        alerts.append("PCOS diagnosis - metabolic screening recommended")

    # Risk stratification
    if score >= 10:
        category = "HIGH"
        recommendation = "Immediate specialist referral required"
    elif score >= 5:
        category = "MODERATE"
        recommendation = "Close surveillance and specialist consultation"
    elif score >= 2:
        category = "LOW-MODERATE"
        recommendation = "Enhanced screening protocols"
    else:
        category = "LOW"
        recommendation = "Routine age-appropriate screening"

    return {
        "composite_score": score,
        "risk_category": category,
        "recommendation": recommendation,
        "alerts": alerts
    }
```

---

## 6. Database Schema Design

### 6.1 Recommended PostgreSQL Schema

```sql
-- Extend existing lab_results table or create imaging_exams table

CREATE TYPE birads_category AS ENUM ('0', '1', '2', '3', '4A', '4B', '4C', '5', '6');
CREATE TYPE breast_density_category AS ENUM ('A', 'B', 'C', 'D');
CREATE TYPE orads_category AS ENUM ('0', '1', '2', '3', '4', '5');
CREATE TYPE mammography_modality AS ENUM ('2D_digital', '3D_tomosynthesis', 'film');

-- Mammography results table
CREATE TABLE mammography_results (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    exam_date DATE NOT NULL,
    exam_type VARCHAR(50) NOT NULL, -- screening, diagnostic, follow_up
    facility_id UUID REFERENCES facilities(id),
    radiologist_id UUID REFERENCES users(id),

    -- BI-RADS classification
    birads_category birads_category NOT NULL,
    birads_numeric INTEGER NOT NULL CHECK (birads_numeric BETWEEN 0 AND 8),
    cancer_risk_percent DECIMAL(5,2),

    -- Breast density
    breast_density breast_density_category NOT NULL,
    breast_density_numeric INTEGER NOT NULL CHECK (breast_density_numeric BETWEEN 1 AND 4),

    -- Technology
    modality mammography_modality NOT NULL DEFAULT '2D_digital',

    -- Findings (boolean flags for structured querying)
    calcifications_present BOOLEAN DEFAULT FALSE,
    mass_present BOOLEAN DEFAULT FALSE,
    asymmetry_present BOOLEAN DEFAULT FALSE,
    architectural_distortion BOOLEAN DEFAULT FALSE,

    -- Free text
    findings_description TEXT,

    -- Recommendations
    recommendation TEXT NOT NULL,
    follow_up_months INTEGER,
    biopsy_recommended BOOLEAN DEFAULT FALSE,
    supplemental_screening_recommended BOOLEAN DEFAULT FALSE,

    -- Audit
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by UUID REFERENCES users(id),

    -- Indexes for performance
    INDEX idx_patient_exam_date (patient_id, exam_date DESC),
    INDEX idx_birads_category (birads_category),
    INDEX idx_biopsy_recommended (biopsy_recommended) WHERE biopsy_recommended = TRUE
);

-- Breast ultrasound results table
CREATE TABLE breast_ultrasound_results (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    exam_date DATE NOT NULL,
    exam_type VARCHAR(50) NOT NULL, -- screening_supplemental, diagnostic, biopsy_guidance
    indication TEXT,

    -- BI-RADS classification
    birads_category birads_category NOT NULL,
    birads_numeric INTEGER NOT NULL CHECK (birads_numeric BETWEEN 0 AND 8),
    cancer_risk_percent DECIMAL(5,2),

    -- Findings characteristics
    mass_present BOOLEAN DEFAULT FALSE,
    mass_characteristics TEXT,
    solid_component BOOLEAN,
    vascular_pattern VARCHAR(50), -- none, mild, moderate, marked
    elastography_score INTEGER CHECK (elastography_score BETWEEN 1 AND 5),

    -- Correlation
    correlates_with_mammogram BOOLEAN,
    mammography_result_id UUID REFERENCES mammography_results(id),

    -- Free text
    findings_description TEXT,

    -- Recommendations
    recommendation TEXT NOT NULL,
    biopsy_recommended BOOLEAN DEFAULT FALSE,

    -- Audit
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by UUID REFERENCES users(id),

    INDEX idx_patient_exam_date (patient_id, exam_date DESC),
    INDEX idx_birads_category (birads_category)
);

-- Transvaginal ultrasound results table
CREATE TABLE transvaginal_ultrasound_results (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    exam_date DATE NOT NULL,
    indication TEXT,

    -- Patient context
    postmenopausal BOOLEAN NOT NULL,
    has_bleeding BOOLEAN DEFAULT FALSE,
    on_hrt BOOLEAN DEFAULT FALSE,
    on_tamoxifen BOOLEAN DEFAULT FALSE,
    menstrual_cycle_day INTEGER CHECK (menstrual_cycle_day BETWEEN 1 AND 35),

    -- Endometrial assessment
    endometrial_thickness_mm DECIMAL(4,1) NOT NULL,
    endometrial_risk_level VARCHAR(50),
    endometrial_recommendation TEXT,
    endometrial_biopsy_recommended BOOLEAN DEFAULT FALSE,

    -- Right ovary
    right_ovary_volume_cm3 DECIMAL(5,2),
    right_ovary_orads orads_category,
    right_ovary_risk_percent DECIMAL(5,2),
    right_follicle_count INTEGER,

    -- Left ovary
    left_ovary_volume_cm3 DECIMAL(5,2),
    left_ovary_orads orads_category,
    left_ovary_risk_percent DECIMAL(5,2),
    left_follicle_count INTEGER,

    -- PCOS assessment
    pcos_criteria_met BOOLEAN DEFAULT FALSE,

    -- Uterine findings
    myoma_present BOOLEAN DEFAULT FALSE,
    myoma_figo_type VARCHAR(10), -- Type 0-8
    adenomyosis_present BOOLEAN DEFAULT FALSE,
    uterine_anomaly VARCHAR(100),

    -- Adnexal findings
    hydrosalpinx_present BOOLEAN DEFAULT FALSE,
    endometrioma_present BOOLEAN DEFAULT FALSE,
    free_fluid_present BOOLEAN DEFAULT FALSE,

    -- Free text
    findings_description TEXT,

    -- Recommendations
    recommendation TEXT NOT NULL,
    referral_specialty VARCHAR(50), -- gynecology, gynecologic_oncology, reproductive_endocrinology

    -- Audit
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by UUID REFERENCES users(id),

    INDEX idx_patient_exam_date (patient_id, exam_date DESC),
    INDEX idx_endometrial_thickness (endometrial_thickness_mm),
    INDEX idx_biopsy_recommended (endometrial_biopsy_recommended) WHERE endometrial_biopsy_recommended = TRUE,
    INDEX idx_orads_high_risk (right_ovary_orads, left_ovary_orads) WHERE right_ovary_orads >= '4' OR left_ovary_orads >= '4',
    INDEX idx_pcos (pcos_criteria_met) WHERE pcos_criteria_met = TRUE
);

-- Composite risk tracking table
CREATE TABLE womens_health_risk_scores (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    calculation_date DATE NOT NULL,

    -- Component scores
    breast_risk_score INTEGER DEFAULT 0,
    endometrial_risk_score INTEGER DEFAULT 0,
    ovarian_risk_score INTEGER DEFAULT 0,
    metabolic_risk_score INTEGER DEFAULT 0,

    -- Composite
    total_risk_score INTEGER NOT NULL,
    risk_category VARCHAR(20) NOT NULL, -- LOW, LOW-MODERATE, MODERATE, HIGH

    -- Recommendations
    recommendation TEXT,
    alerts TEXT[],

    -- Linked exams
    mammography_result_id UUID REFERENCES mammography_results(id),
    breast_ultrasound_result_id UUID REFERENCES breast_ultrasound_results(id),
    transvaginal_ultrasound_result_id UUID REFERENCES transvaginal_ultrasound_results(id),

    -- Audit
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_patient_calculation_date (patient_id, calculation_date DESC),
    INDEX idx_risk_category (risk_category)
);

-- Audit log for all imaging access (LGPD requirement)
CREATE TABLE imaging_audit_log (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    patient_id UUID NOT NULL REFERENCES patients(id),
    action VARCHAR(50) NOT NULL, -- view, create, update, export
    resource_type VARCHAR(50) NOT NULL, -- mammography, ultrasound, etc
    resource_id UUID NOT NULL,
    ip_address INET,
    user_agent TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_patient_access (patient_id, created_at DESC),
    INDEX idx_user_actions (user_id, action, created_at DESC)
);
```

### 6.2 Go Model Implementation (Única Fonte de Verdade)

```go
// apps/api/internal/models/mammography_result.go

package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

// @Description Mammography examination result
type MammographyResult struct {
    ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    PatientID  uuid.UUID `gorm:"type:uuid;not null;index" json:"patient_id" validate:"required"`
    ExamDate   time.Time `gorm:"type:date;not null" json:"exam_date" validate:"required"`
    ExamType   string    `gorm:"type:varchar(50);not null" json:"exam_type" validate:"required,oneof=screening diagnostic follow_up"`

    // @enum 0,1,2,3,4A,4B,4C,5,6
    BiradsCategory string `gorm:"type:varchar(2);not null" json:"birads_category" validate:"required,oneof=0 1 2 3 4A 4B 4C 5 6"`

    // @minimum 0
    // @maximum 8
    BiradsNumeric int `gorm:"type:integer;not null;check:birads_numeric BETWEEN 0 AND 8" json:"birads_numeric" validate:"required,min=0,max=8"`

    // @minimum 0
    // @maximum 100
    CancerRiskPercent *float64 `gorm:"type:decimal(5,2)" json:"cancer_risk_percent,omitempty"`

    // @enum A,B,C,D
    BreastDensity string `gorm:"type:varchar(1);not null" json:"breast_density" validate:"required,oneof=A B C D"`

    // @minimum 1
    // @maximum 4
    BreastDensityNumeric int `gorm:"type:integer;not null;check:breast_density_numeric BETWEEN 1 AND 4" json:"breast_density_numeric" validate:"required,min=1,max=4"`

    // @enum 2D_digital,3D_tomosynthesis,film
    Modality string `gorm:"type:varchar(20);not null;default:'2D_digital'" json:"modality" validate:"oneof=2D_digital 3D_tomosynthesis film"`

    CalcificationsPresent     bool    `gorm:"default:false" json:"calcifications_present"`
    MassPresent               bool    `gorm:"default:false" json:"mass_present"`
    AsymmetryPresent          bool    `gorm:"default:false" json:"asymmetry_present"`
    ArchitecturalDistortion   bool    `gorm:"default:false" json:"architectural_distortion"`

    FindingsDescription       *string `gorm:"type:text" json:"findings_description,omitempty"`

    Recommendation            string  `gorm:"type:text;not null" json:"recommendation" validate:"required"`
    FollowUpMonths            *int    `gorm:"type:integer" json:"follow_up_months,omitempty"`
    BiopsyRecommended         bool    `gorm:"default:false" json:"biopsy_recommended"`
    SupplementalScreening     bool    `gorm:"default:false" json:"supplemental_screening_recommended"`

    CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
    CreatedBy *uuid.UUID     `gorm:"type:uuid" json:"created_by,omitempty"`
}

// TableName overrides the table name
func (MammographyResult) TableName() string {
    return "mammography_results"
}

// BeforeCreate hook for additional validation
func (m *MammographyResult) BeforeCreate(tx *gorm.DB) error {
    // Auto-calculate cancer risk if not provided
    if m.CancerRiskPercent == nil {
        risk := calculateBiradsRisk(m.BiradsNumeric)
        m.CancerRiskPercent = &risk
    }
    return nil
}

func calculateBiradsRisk(biradsNumeric int) float64 {
    riskMap := map[int]float64{
        0: 0.0, 1: 0.1, 2: 0.1, 3: 2.0,
        4: 6.0, 5: 30.0, 6: 72.5, 7: 95.0, 8: 100.0,
    }
    return riskMap[biradsNumeric]
}
```

```go
// apps/api/internal/models/transvaginal_ultrasound_result.go

package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

// @Description Transvaginal ultrasound examination result
type TransvaginalUltrasoundResult struct {
    ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    PatientID   uuid.UUID `gorm:"type:uuid;not null;index" json:"patient_id" validate:"required"`
    ExamDate    time.Time `gorm:"type:date;not null" json:"exam_date" validate:"required"`
    Indication  *string   `gorm:"type:text" json:"indication,omitempty"`

    // Patient context
    Postmenopausal bool `gorm:"not null" json:"postmenopausal"`
    HasBleeding    bool `gorm:"default:false" json:"has_bleeding"`
    OnHRT          bool `gorm:"default:false" json:"on_hrt"`
    OnTamoxifen    bool `gorm:"default:false" json:"on_tamoxifen"`

    // @minimum 1
    // @maximum 35
    MenstrualCycleDay *int `gorm:"type:integer;check:menstrual_cycle_day BETWEEN 1 AND 35" json:"menstrual_cycle_day,omitempty"`

    // Endometrial assessment
    // @minimum 0
    // @maximum 50
    EndometrialThicknessMm    float64 `gorm:"type:decimal(4,1);not null" json:"endometrial_thickness_mm" validate:"required,min=0,max=50"`
    EndometrialRiskLevel      *string `gorm:"type:varchar(50)" json:"endometrial_risk_level,omitempty"`
    EndometrialRecommendation *string `gorm:"type:text" json:"endometrial_recommendation,omitempty"`
    EndometrialBiopsy         bool    `gorm:"default:false" json:"endometrial_biopsy_recommended"`

    // Right ovary
    // @minimum 0
    // @maximum 100
    RightOvaryVolumeCm3   *float64 `gorm:"type:decimal(5,2)" json:"right_ovary_volume_cm3,omitempty"`

    // @enum 0,1,2,3,4,5
    RightOvaryOrads       *string  `gorm:"type:varchar(1)" json:"right_ovary_orads,omitempty" validate:"omitempty,oneof=0 1 2 3 4 5"`
    RightOvaryRiskPercent *float64 `gorm:"type:decimal(5,2)" json:"right_ovary_risk_percent,omitempty"`
    RightFollicleCount    *int     `gorm:"type:integer" json:"right_follicle_count,omitempty"`

    // Left ovary
    LeftOvaryVolumeCm3   *float64 `gorm:"type:decimal(5,2)" json:"left_ovary_volume_cm3,omitempty"`
    LeftOvaryOrads       *string  `gorm:"type:varchar(1)" json:"left_ovary_orads,omitempty" validate:"omitempty,oneof=0 1 2 3 4 5"`
    LeftOvaryRiskPercent *float64 `gorm:"type:decimal(5,2)" json:"left_ovary_risk_percent,omitempty"`
    LeftFollicleCount    *int     `gorm:"type:integer" json:"left_follicle_count,omitempty"`

    // PCOS assessment
    PcosCriteriaMet bool `gorm:"default:false" json:"pcos_criteria_met"`

    // Uterine findings
    MyomaPresent    bool    `gorm:"default:false" json:"myoma_present"`
    MyomaFigoType   *string `gorm:"type:varchar(10)" json:"myoma_figo_type,omitempty"`
    Adenomyosis     bool    `gorm:"default:false" json:"adenomyosis_present"`
    UterineAnomaly  *string `gorm:"type:varchar(100)" json:"uterine_anomaly,omitempty"`

    // Adnexal findings
    Hydrosalpinx    bool `gorm:"default:false" json:"hydrosalpinx_present"`
    Endometrioma    bool `gorm:"default:false" json:"endometrioma_present"`
    FreeFluid       bool `gorm:"default:false" json:"free_fluid_present"`

    FindingsDescription *string `gorm:"type:text" json:"findings_description,omitempty"`

    Recommendation     string  `gorm:"type:text;not null" json:"recommendation" validate:"required"`
    ReferralSpecialty  *string `gorm:"type:varchar(50)" json:"referral_specialty,omitempty"`

    CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
    CreatedBy *uuid.UUID `gorm:"type:uuid" json:"created_by,omitempty"`
}

func (TransvaginalUltrasoundResult) TableName() string {
    return "transvaginal_ultrasound_results"
}

// BeforeCreate hook for automated risk calculation
func (t *TransvaginalUltrasoundResult) BeforeCreate(tx *gorm.DB) error {
    // Calculate endometrial risk
    t.calculateEndometrialRisk()

    // Calculate ovarian risks
    t.calculateOvarianRisks()

    // Assess PCOS criteria
    t.assessPCOS()

    return nil
}

func (t *TransvaginalUltrasoundResult) calculateEndometrialRisk() {
    thickness := t.EndometrialThicknessMm

    if !t.Postmenopausal {
        if thickness > 16 {
            risk := "investigate"
            t.EndometrialRiskLevel = &risk
            rec := "Rule out endometrial pathology"
            t.EndometrialRecommendation = &rec
            t.EndometrialBiopsy = true
        }
        return
    }

    // Postmenopausal with bleeding
    if t.HasBleeding {
        if thickness <= 4 {
            risk := "low"
            t.EndometrialRiskLevel = &risk
            rec := "Routine surveillance"
            t.EndometrialRecommendation = &rec
        } else if thickness <= 5 {
            risk := "low-moderate"
            t.EndometrialRiskLevel = &risk
            rec := "Consider biopsy based on clinical context"
            t.EndometrialRecommendation = &rec
        } else {
            risk := "moderate-high"
            t.EndometrialRiskLevel = &risk
            rec := "Endometrial biopsy required"
            t.EndometrialRecommendation = &rec
            t.EndometrialBiopsy = true
        }
        return
    }

    // Postmenopausal without bleeding
    threshold := 11.0
    if t.OnHRT || t.OnTamoxifen {
        threshold = 8.0
    }

    if thickness <= threshold {
        risk := "low"
        t.EndometrialRiskLevel = &risk
        rec := "Routine surveillance"
        t.EndometrialRecommendation = &rec
    } else {
        risk := "elevated"
        t.EndometrialRiskLevel = &risk
        rec := "Consider biopsy or close surveillance"
        t.EndometrialRecommendation = &rec
    }
}

func (t *TransvaginalUltrasoundResult) calculateOvarianRisks() {
    if t.RightOvaryOrads != nil {
        risk := oradsToRiskPercent(*t.RightOvaryOrads)
        t.RightOvaryRiskPercent = &risk
    }

    if t.LeftOvaryOrads != nil {
        risk := oradsToRiskPercent(*t.LeftOvaryOrads)
        t.LeftOvaryRiskPercent = &risk
    }
}

func (t *TransvaginalUltrasoundResult) assessPCOS() {
    // Rotterdam criteria: need 2 of 3
    // 1. Ovarian volume >=10 cm³ OR >=12 follicles 2-9mm
    morphologyCriteria := false

    if (t.RightOvaryVolumeCm3 != nil && *t.RightOvaryVolumeCm3 >= 10) ||
       (t.LeftOvaryVolumeCm3 != nil && *t.LeftOvaryVolumeCm3 >= 10) ||
       (t.RightFollicleCount != nil && *t.RightFollicleCount >= 12) ||
       (t.LeftFollicleCount != nil && *t.LeftFollicleCount >= 12) {
        morphologyCriteria = true
    }

    // Note: Clinical criteria (oligo/anovulation, hyperandrogenism)
    // must be assessed separately - this only flags ultrasound findings
    if morphologyCriteria {
        t.PcosCriteriaMet = true
    }
}

func oradsToRiskPercent(orads string) float64 {
    riskMap := map[string]float64{
        "0": 0.0, "1": 0.0, "2": 0.5, "3": 5.0, "4": 30.0, "5": 75.0,
    }
    return riskMap[orads]
}
```

---

## 7. Implementation Recommendations for Plenya EMR

### 7.1 Phased Rollout

**Phase 1: Data Capture (Foundation)**
- Implement Go models for all three exam types
- Generate Atlas migrations
- Create OpenAPI spec → TypeScript types
- Build basic CRUD endpoints (POST, GET by ID, GET list)

**Phase 2: Automated Risk Stratification**
- Implement BeforeCreate hooks for risk calculation
- Add validation rules
- Create automated recommendation generation
- Build alert system for high-risk findings

**Phase 3: Clinical Decision Support**
- Integrate risk algorithms into frontend
- Display visual risk indicators (color-coded)
- Automated follow-up scheduling
- Referral workflow triggers

**Phase 4: Analytics & Reporting**
- CSV export functionality
- Population health dashboards
- Screening compliance tracking
- Quality metrics (recall rates, cancer detection rates)

**Phase 5: Integration**
- HL7 FHIR integration for external lab systems
- PDF report parsing (OCR + NLP)
- Automated data extraction from radiology reports

### 7.2 Frontend UI/UX Recommendations

**Risk Visualization:**
- Traffic light system (green/yellow/red) for BI-RADS and O-RADS
- Timeline view showing progression of findings over time
- Comparative charts (current vs previous exams)

**Clinical Workflows:**
- Smart forms with conditional fields (show HRT question only if postmenopausal)
- Auto-populated recommendations based on entered values
- One-click referral generation
- Follow-up reminder creation

**Patient Portal:**
- Simplified risk explanations (avoid medical jargon)
- Educational resources linked to findings
- Preventive care recommendations
- Shared decision-making tools

### 7.3 LGPD Compliance Considerations

**Data Sensitivity Classification:**
- Imaging findings: HIGH sensitivity (health data)
- BI-RADS/O-RADS categories: MEDIUM (semi-public medical codes)
- Risk percentages: HIGH (derived health data)

**Security Measures:**
- Encrypt findings_description fields (free text may contain PII)
- Audit all access to imaging results (immutable log)
- Role-based access control (patients see own data only)
- Retention policy: 20 years for imaging reports (Brazilian requirement)

**Patient Rights:**
- Data export in human-readable format (PDF summary)
- Right to rectification (allow corrections with audit trail)
- Right to erasure (soft delete with anonymization after retention period)

### 7.4 Quality Assurance & Validation

**Data Quality Checks:**
- Flag impossible values (e.g., endometrial thickness >50mm)
- Cross-validation (postmenopausal + menstrual cycle day = error)
- Mandatory fields enforcement
- Standardized vocabularies (SNOMED CT mapping)

**Clinical Validation:**
- Peer review for high-risk findings (BI-RADS 5-6, O-RADS 5)
- Double-read program for screening mammography
- Biopsy correlation tracking (validate positive predictive value)

### 7.5 Performance Optimization

**Database Indexes:**
```sql
-- Fast retrieval of recent exams
CREATE INDEX idx_patient_recent_mammo ON mammography_results(patient_id, exam_date DESC);

-- High-risk patient queries
CREATE INDEX idx_high_risk_breast ON mammography_results(birads_numeric)
    WHERE birads_numeric >= 4;

CREATE INDEX idx_high_risk_ovarian ON transvaginal_ultrasound_results(
    patient_id, right_ovary_orads, left_ovary_orads
) WHERE right_ovary_orads >= '4' OR left_ovary_orads >= '4';

-- PCOS cohort identification
CREATE INDEX idx_pcos_cohort ON transvaginal_ultrasound_results(patient_id)
    WHERE pcos_criteria_met = TRUE;
```

**Caching Strategy:**
- Cache patient risk scores (recalculate daily)
- Precompute population statistics for dashboards
- Materialize views for reporting queries

---

## 8. References

### Clinical Guidelines

1. **ACR BI-RADS Atlas, 5th Edition (2013)**
   - Current standard for mammography and ultrasound reporting
   - [Mammogram Results and BI-RADS Score Category](https://www.cancercenter.com/cancer-types/breast-cancer/diagnosis-and-detection/mammography/results-bi-rads)
   - [BI-RADS Categories & Modern Mammography](https://www.mtmi.net/blog/bi-rads-categories-explained-by-mammographers)

2. **USPSTF Breast Cancer Screening Guidelines (April 2024)**
   - Grade B recommendation: Biennial screening ages 40-74
   - [Final Recommendation Statement](https://www.uspreventiveservicestaskforce.org/uspstf/recommendation/breast-cancer-screening)
   - [USPSTF Screening for Breast Cancer](https://jamanetwork.com/pages/uspstf-screening-for-breast-cancer)

3. **Brazilian Society of Mastology Guidelines (2023)**
   - Annual screening starting age 40 (aligned with ACR)
   - [SBM/CBR/FEBRASGO Recommendations](https://pmc.ncbi.nlm.nih.gov/articles/PMC10491472/)
   - [Updated Screening Recommendations](https://journalrbgo.org/article/recommendations-for-the-screening-of-breast-cancer-of-the-brazilian-college-of-radiology-and-diagnostic-imaging-brazilian-society-of-mastology-and-brazilian-federation-of-gynecology-and-obstetrics-as/)

4. **O-RADS US Risk Stratification (ACR 2020)**
   - [ACR O-RADS Official Resource](https://www.acr.org/Clinical-Resources/Clinical-Tools-and-Reference/Reporting-and-Data-Systems/O-RADS)
   - [O-RADS Consensus Guideline - Radiology](https://pubs.rsna.org/doi/full/10.1148/radiol.2019191150)
   - [Radiopaedia Reference](https://radiopaedia.org/articles/ovarian-adnexal-reporting-and-data-system-ultrasound-o-rads-us)

5. **IOTA Simple Rules**
   - [Radiopaedia IOTA Rules](https://radiopaedia.org/articles/iota-ultrasound-rules-for-ovarian-masses)
   - [IOTA Simple Rules Review - RadioGraphics](https://pubs.rsna.org/doi/full/10.1148/rg.210068)
   - [2024 Validation Study](https://pmc.ncbi.nlm.nih.gov/articles/PMC10827489/)

6. **ACOG Endometrial Thickness Guidelines**
   - [Committee Opinion on Transvaginal Ultrasonography](https://www.acog.org/clinical/clinical-guidance/committee-opinion/articles/2018/05/the-role-of-transvaginal-ultrasonography-in-evaluating-the-endometrium-of-women-with-postmenopausal-bleeding)
   - [Endometrial Thickness - Radiopaedia](https://radiopaedia.org/articles/endometrial-thickness)
   - [2024 Re-evaluation Meta-Analysis](https://www.jacr.org/article/S1546-1440(24)00918-9/fulltext)

### Recent Evidence (2023-2026)

7. **Digital Breast Tomosynthesis Performance**
   - [10-Year Study - RSNA 2024](https://www.rsna.org/news/2024/september/tomosynthesis-improves-breast-cancer-detection)
   - [Radiology 2023 - Screening Performance Metrics](https://pubs.rsna.org/doi/10.1148/radiol.230505)
   - [DBT Improves Mammography Screening](https://pubs.rsna.org/doi/10.1148/radiol.230306)
   - [BCSC Successive Screening Rounds](https://pmc.ncbi.nlm.nih.gov/articles/PMC10315524/)

8. **Automated Breast Ultrasound (ABUS)**
   - [2024 Egyptian Radiology Study](https://ejrnm.springeropen.com/articles/10.1186/s43055-024-01258-3)
   - [5-Year Academic Experience](https://academic.oup.com/jbi/article/7/1/35/7775421)
   - [European Radiology 2023](https://link.springer.com/article/10.1007/s00330-023-10568-5)
   - [ABUS Screening Narrative Review](https://pmc.ncbi.nlm.nih.gov/articles/PMC8400952/)

9. **Breast Density Research**
   - [FDA 2024 Mandate on Density Reporting](https://www.cancer.org/cancer/types/breast-cancer/screening-tests-and-early-detection/mammograms/breast-density-and-your-mammogram-report.html)
   - [ACR 2024 Supplemental Screening Update](https://pubmed.ncbi.nlm.nih.gov/40409891/)
   - [Density Impact on Risk Assessment 2024](https://ejbc.kr/DOIx.php?id=10.4048/jbc.2024.0101)

### Functional Medicine Resources

10. **Hormone Balance & PCOS**
    - [Institute for Functional Medicine - Hormone APM](https://www.ifm.org/hormone)
    - [IFM PCOS Treatment & Care](https://www.ifm.org/news-insights/pcos-treatment-care/)
    - [Functional Medicine PCOS Protocol](https://www.rupahealth.com/post/a-functional-medicine-pcos-protocol-comprehensive-testing-therapeutic-diet-and-supplements)
    - [Root Functional Medicine - PCOS](https://rootfunctionalmedicine.com/pcos-functional-medicine)

### Technical Implementation

11. **Database & Schema Design**
    - PostgreSQL 17 Official Documentation
    - GORM v1.25 Documentation
    - Atlas Migration Tool Documentation

12. **Standards & Interoperability**
    - HL7 FHIR R4 Diagnostic Report Profile
    - SNOMED CT Imaging Terminology
    - LOINC Codes for Imaging Procedures

---

## Appendix A: Quick Reference Tables

### A.1 BI-RADS Quick Reference

| Category | Risk | Action | Timeframe |
|----------|------|--------|-----------|
| 0 | N/A | Additional imaging | 30 days |
| 1 | 0.1% | Routine screening | 1-2 years |
| 2 | 0.1% | Routine screening | 1-2 years |
| 3 | <2% | Short follow-up | 6 months |
| 4A | 2-10% | Biopsy | 1-2 weeks |
| 4B | 10-50% | Biopsy | 1 week |
| 4C | 50-95% | Biopsy | <1 week |
| 5 | ≥95% | Biopsy + oncology | Urgent |
| 6 | 100% | Treatment planning | Immediate |

### A.2 O-RADS Quick Reference

| Category | Risk | Management | Timeframe |
|----------|------|-----------|-----------|
| 0 | N/A | Complete evaluation | 30 days |
| 1 | 0% | Routine surveillance | N/A |
| 2 | <1% | Annual follow-up | 12 months |
| 3 | 1-10% | Annual ultrasound | 12 months |
| 4 | 10-50% | Surgery consultation | 4-8 weeks |
| 5 | ≥50% | Gyn-onc referral | <2 weeks |

### A.3 Endometrial Thickness Decision Matrix

| Status | Bleeding | Thickness | Action |
|--------|----------|-----------|--------|
| Premenopausal | Any | <16mm | Normal |
| Premenopausal | Yes | ≥16mm | Investigate |
| Postmenopausal | Yes | ≤4mm | Surveillance |
| Postmenopausal | Yes | >5mm | Biopsy |
| Postmenopausal | No | ≤11mm | Surveillance |
| Postmenopausal | No | >11mm | Consider biopsy |
| On HRT/Tamoxifen | No | ≤8mm | Surveillance |
| On HRT/Tamoxifen | No | >8mm | Consider biopsy |

---

## Document Control

**Prepared by:** Claude Sonnet 4.5 (Anthropic)
**Approved by:** [Clinical Director Name]
**Review Date:** Every 6 months or upon guideline updates
**Version History:**
- v1.0 (2026-01-19): Initial comprehensive documentation

**Revision Schedule:**
- Monitor USPSTF, ACR, SBM guideline updates
- Review emerging evidence on DBT, ABUS, AI-assisted interpretation
- Update risk stratification algorithms based on institutional data
- Validate automated clinical decision support accuracy

---

**END OF DOCUMENT**
