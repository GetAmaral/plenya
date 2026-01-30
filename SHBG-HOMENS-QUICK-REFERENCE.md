# SHBG - Homens: Quick Reference Guide

**Score Item ID:** `fe938255-6b7a-4fbd-ac8f-8f3ba0c2d291`
**Status:** ✅ ENRICHED | **Date:** 2026-01-29

---

## At a Glance

| Metric | Value | Status |
|--------|-------|--------|
| Clinical Relevance | 1,494 chars | ✅ |
| Patient Explanation | 1,367 chars | ✅ |
| Conduct | 2,000 chars | ✅ |
| Scientific Articles | 4 (PMIDs) | ✅ |
| Last Review | 2026-01-29 | ✅ |

---

## Clinical Quick Facts

### Reference Ranges
- **Adult Men:** 13-71 nmol/L
- **Metabolic Risk Threshold:** <17 nmol/L
- **Age Effect:** Increases ~1.3% per year after 50

### Testosterone Binding Distribution
- **SHBG-bound (inactive):** 45%
- **Albumin-bound (bioavailable):** 50%
- **Free (active):** 2-3%

### Key Clinical Associations

**Low SHBG (<17 nmol/L):**
- Insulin resistance (independent predictor, p=0.029)
- Metabolic syndrome
- Type 2 diabetes (RR 0.10 highest vs lowest quartile)
- Visceral obesity
- Hepatic steatosis

**High SHBG (>71 nmol/L):**
- Aging (physiological)
- Primary hypogonadism
- Hyperthyroidism
- Cirrhosis
- Anticonvulsants/estrogen use

---

## Clinical Protocol Summary

### When to Order SHBG
1. Suspected functional hypogonadism
2. Metabolic syndrome evaluation
3. Obesity with androgenic symptoms
4. Testosterone total discordant with symptoms
5. Men >50 years with fatigue/low libido

### Testing Requirements
- **Timing:** Morning (7-11h)
- **Fasting:** 8 hours
- **Concurrent Tests:** Total testosterone, LH, FSH, prolactin

### Free Testosterone Calculation
**When:** SHBG outside reference range
**Methods:**
- Gold standard: Equilibrium dialysis
- Practical: Vermeulen formula (total T + SHBG + albumin)
**Cutoff:** <65 pg/mL with symptoms

### Additional Workup (Low SHBG)
- HOMA-IR
- HbA1c
- Lipid panel
- Liver enzymes
- Abdominal ultrasound

### Monitoring During TRT
- **Frequency:** Every 6 months initially
- **Rationale:** Androgen normalization may elevate SHBG

---

## Age-Related Testosterone Decline

| Parameter | Annual Decline | Clinical Significance |
|-----------|----------------|----------------------|
| SHBG | **+1.3%** (increases) | Reduces bioavailable testosterone |
| Total Testosterone | -1.6% | Moderate decline |
| Free Testosterone | **-2.3%** | Most significant decline |

**Clinical Pearl:** Free testosterone declines faster than total due to SHBG increase, explaining hypogonadal symptoms despite "normal" total testosterone in older men.

---

## Metabolic Vicious Cycle

```
Visceral Obesity
    ↓
Hyperinsulinemia
    ↓
↓ SHBG Production (liver)
    ↓
↓ Bioavailable Testosterone
    ↓
↓ Muscle Mass, ↑ Fat Mass
    ↓
↑ Insulin Resistance
    ↓
[CYCLE REPEATS]
```

**Intervention Points:**
- Metformin (diabetic men with SHBG <17)
- Glycemic control intensification
- Weight loss
- Testosterone replacement (if indicated)

---

## Treatment Decision Algorithm

### Candidate for TRT Trial (Men >50 years)
All 3 criteria must be met:
1. ✓ Low-normal total testosterone + elevated SHBG
2. ✓ Clinical symptoms (fatigue, muscle loss, low libido)
3. ✓ No contraindications (PSA >4, Hct >54%, severe untreated apnea)

### Not a TRT Candidate
- Normal free testosterone (calculated)
- SHBG elevation explained by medication (estrogens, anticonvulsants)
- Active contraindications present

---

## Scientific Evidence

### 4 Peer-Reviewed Articles (2023-2025)

**1. Biomedicines 2025** (PMID: 40427034)
- Review of SHBG diagnostic potential
- Beyond hormone transport: metabolic syndrome, CVD risk

**2. Diabetes Therapy 2023** (PMID: 37462840)
- SHBG independent predictor of insulin resistance (p=0.029)
- Optimal cutoff: 17.2 nmol/L in newly diagnosed T2DM men

**3. Arch Endocrinol Metab 2024** (PMID: 40857627)
- Age-specific SHBG elevation in men >50 years
- Inverse correlation with insulin resistance

**4. J Clin Endocrinol Metab 2023** (PMID: 36995891)
- Clinical approach to low testosterone in men ≥50 years
- JCEM clinical guidance review

---

## Patient Communication Template

### "What is SHBG?"
"SHBG is a protein that carries testosterone in your blood. Think of it like a train car: when testosterone is inside the car (bound to SHBG), it can't get into your cells to work. Only the 'free' testosterone that's not bound can actually increase your muscle, energy, and libido."

### "Why is mine low?"
"Low SHBG (below 17) usually indicates metabolic problems like insulin resistance or increased diabetes risk. It's common in men with belly fat or metabolic syndrome."

### "Why is mine high?"
"High SHBG increases with age (normal after 50). However, it means less of your testosterone is available to use, even if your total testosterone looks normal. That's why we calculate your 'free testosterone' separately."

### "Why test it with testosterone?"
"Your total testosterone number can be misleading without SHBG. Two men with the same total testosterone can have very different amounts of active, usable testosterone depending on their SHBG level."

---

## Common Clinical Scenarios

### Scenario 1: "Normal" Testosterone, Symptomatic Patient
- **Finding:** Total T = 350 ng/dL (low-normal), SHBG = 85 nmol/L (high)
- **Interpretation:** Elevated SHBG reduces bioavailable testosterone
- **Action:** Calculate free testosterone, consider TRT trial if <65 pg/mL

### Scenario 2: Metabolic Syndrome Patient
- **Finding:** SHBG = 15 nmol/L, BMI = 34, prediabetic
- **Interpretation:** Low SHBG indicates insulin resistance
- **Action:** Intensify glycemic control, consider metformin, lifestyle modification

### Scenario 3: Aging Male, Mild Symptoms
- **Finding:** Age 62, SHBG = 68 nmol/L (age-appropriate elevation)
- **Interpretation:** Physiological SHBG increase with aging
- **Action:** Calculate free testosterone to assess true hypogonadism

---

## Key Takeaways

1. **SHBG is NOT optional:** Mandatory for accurate testosterone assessment
2. **Low SHBG = Metabolic red flag:** Screen for insulin resistance and diabetes
3. **High SHBG in older men:** Calculate free testosterone to avoid missing hypogonadism
4. **Free > Total:** Free testosterone is more clinically meaningful than total
5. **Breaking the cycle:** Address insulin resistance to improve SHBG and testosterone

---

## Files

- **Full Report:** `/home/user/plenya/SHBG-HOMENS-ENRICHMENT-REPORT.md`
- **SQL Script:** `/home/user/plenya/scripts/enrich_shbg_homens_final.sql`
- **Quick Reference:** `/home/user/plenya/SHBG-HOMENS-QUICK-REFERENCE.md` (this file)

---

## Database Query

```sql
-- View enriched item
SELECT
    name,
    LENGTH(clinical_relevance) as clin_rel,
    LENGTH(patient_explanation) as pat_exp,
    LENGTH(conduct) as conduct,
    last_review
FROM score_items
WHERE id = 'fe938255-6b7a-4fbd-ac8f-8f3ba0c2d291';

-- View linked articles
SELECT a.pm_id, a.title, a.journal, a.publish_date
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = 'fe938255-6b7a-4fbd-ac8f-8f3ba0c2d291'
  AND a.pm_id IS NOT NULL
ORDER BY a.publish_date DESC;
```

---

*Generated: 2026-01-29 | Plenya EMR System*
