# IGF-1 (Somatomedina C) - Enrichment Report

**Date:** 2026-01-29
**Score Item ID:** 039f1542-7596-4671-8ed0-049b3b41cfc4
**Status:** ✅ COMPLETE

---

## Summary

Successfully enriched the IGF-1 (Somatomedina C) score item with comprehensive clinical content in Portuguese (PT-BR) and linked 4 high-quality peer-reviewed articles.

---

## Content Statistics

| Field | Character Count | Target Range | Status |
|-------|----------------|--------------|--------|
| **clinical_relevance** | 1,656 chars | 1,500-2,000 | ✅ PASS |
| **patient_explanation** | 1,468 chars | 1,000-1,500 | ✅ PASS |
| **conduct** | 2,290 chars | 1,500-2,500 | ✅ PASS |
| **Total Articles** | 13 articles | 2-4 new | ✅ PASS (4 added) |
| **Recent Articles (2014+)** | 13 articles | Current | ✅ PASS |

---

## Scientific Articles Added

### 1. Bailes & Soloviev (2021) - Review
**Title:** Insulin-Like Growth Factor-1 (IGF-1) and Its Monitoring in Medical Diagnostic and in Sports
**Journal:** Biomolecules
**Date:** February 4, 2021
**DOI:** 10.3390/biom11020217
**Type:** Review
**Focus:** Comprehensive review of IGF-1 structure, function, and clinical monitoring challenges

### 2. Bidlingmaier et al. (2014) - Multicenter Study
**Title:** Reference Intervals for Insulin-like Growth Factor-1 (IGF-I) From Birth to Senescence
**Journal:** J Clin Endocrinol Metab
**Date:** May 2014
**DOI:** 10.1210/jc.2013-3059
**Type:** Research Article
**Focus:** Establishes age- and sex-specific reference ranges for IGF-1 (15,014 subjects, 0-94 years)

### 3. Ibba et al. (2020) - Diagnostic Reappraisal
**Title:** IGF1 for the diagnosis of growth hormone deficiency in children and adolescents: a reappraisal
**Journal:** Endocr Connect
**Date:** November 2020
**DOI:** 10.1530/EC-20-0347
**Type:** Research Article
**Focus:** Critical evaluation of IGF-1 diagnostic accuracy in pediatric GHD (sensitivity 67.61%, specificity 62.62%)

### 4. Katznelson et al. (2014) - Clinical Guideline
**Title:** Acromegaly: An Endocrine Society Clinical Practice Guideline
**Journal:** J Clin Endocrinol Metab
**Date:** November 2014
**DOI:** 10.1210/jc.2014-2700
**Type:** Review/Clinical Guideline
**Focus:** Endocrine Society guidelines for IGF-1 in acromegaly diagnosis and monitoring

---

## Clinical Content Overview

### Clinical Relevance (1,656 chars)
Comprehensive coverage of:
- **IGF-1 as GH surrogate marker:** Stable serum levels vs. pulsatile GH secretion
- **Biochemistry:** 98% bound to IGFBPs, especially IGFBP-3 in ternary complex
- **Clinical applications:**
  - GH deficiency diagnosis (pediatric and adult)
  - Acromegaly/gigantism evaluation
  - GH replacement therapy monitoring
  - Growth disorder investigation
- **Age-dependent reference ranges:** Birth → puberty peak (15y) → senescence decline
- **Diagnostic cutoffs:**
  - Low: <2.5th percentile (Z-score <-2) suggests GH deficiency/resistance
  - High: >1.3× ULN with clinical features confirms acromegaly
- **Confounding factors:** Nutrition, chronic disease, hypothyroidism, liver/renal disease

### Patient Explanation (1,468 chars)
Patient-friendly Portuguese explanation covering:
- **What is IGF-1:** Liver-produced protein stimulated by GH
- **Why measure IGF-1:** More stable than GH, reliable marker
- **Clinical importance:** Growth in childhood, metabolic functions in adults
- **Low IGF-1 implications:**
  - Children: short stature
  - Adults: fatigue, muscle loss, bone loss, increased body fat
- **High IGF-1 implications:** Acromegaly (enlarged hands/feet/facial features)
- **Age variability:** Low in infants, peak in adolescence, decline with aging
- **Interpretation context:** Must compare with age-specific references

### Clinical Conduct (2,290 chars)
Detailed clinical protocols including:

**LOW IGF-1 (< 2.5th percentile):**
- Request IGFBP-3
- GH stimulation testing if GHD suspected
- Rule out secondary causes: malnutrition, hypothyroidism, liver disease, renal disease, diabetes, glucocorticoid use
- Pediatric considerations: limited accuracy (67% sensitivity, 63% specificity), better in pubertal stage
- Adult evaluation: consider in symptomatic patients with hypothalamic-pituitary history

**HIGH IGF-1 (>1.3× ULN):**
- Clinical acromegaly features assessment
- Oral glucose tolerance test (OGTT) with GH suppression (confirm if GH doesn't suppress <1 µg/L)
- MRI sella turcica for pituitary adenoma
- Treatment monitoring: target normalized IGF-1 (lower third of normal range)

**GH REPLACEMENT MONITORING:**
- Target: IGF-1 in normal range (mid-to-upper third)
- Serial measurements for dose titration
- Avoid supraphysiological levels

**SPECIAL CONSIDERATIONS:**
- Transition age patients: retest after growth completion if isolated idiopathic GHD
- Multiple pituitary deficiencies (≥3): may not require retest
- Timing: preferably morning, fasting recommended
- Use laboratory-specific reference ranges

---

## Key Clinical Points from Literature

### IGF-1 as GH Axis Marker
- **Stability advantage:** Unlike pulsatile GH secretion with significant diurnal variation, IGF-1 levels show minimal fluctuations (Bailes & Soloviev, 2021)
- **Reference range complexity:** Must account for age, pubertal stage, BMI, pregnancy, sex, ethnicity (Bidlingmaier et al., 2014)
- **Binding proteins:** 98% of circulating IGF-1 bound to IGFBPs, affecting immunoassay measurements

### Diagnostic Limitations
- **Pediatric GHD:** IGF-1 alone has poor accuracy discriminating GHD from non-GHD (AUC 0.69, best cutoff -1.5 SDS) (Ibba et al., 2020)
- **Age effect:** Better diagnostic performance in pubertal (AUC 0.81) vs. prepubertal children (AUC 0.64)
- **Combined interpretation:** Must interpret with other clinical and biochemical parameters

### Acromegaly Guidelines
- **Diagnostic threshold:** IGF-1 >1.3× ULN in patient with typical features confirms diagnosis (Endocrine Society, 2014)
- **Treatment goal:** Normalize IGF-1 to age-specific reference range
- **Prognostic value:** Biochemical control (normalized IGF-1 and GH) is strongest predictor of outcomes and mortality

### Reference Ranges
- **Multicenter validation:** 15,014 subjects (0-94 years) across US, Canada, Europe (Bidlingmaier et al., 2014)
- **Life trajectory:** Decline after birth → increase to pubertal peak (age 15) → continuous decline
- **Population specificity:** Regional, ethnic, racial, nutritional factors affect applicability

---

## Verification Results

```sql
SELECT
  si.name,
  LENGTH(si.clinical_relevance) as clinical_chars,
  LENGTH(si.patient_explanation) as patient_chars,
  LENGTH(si.conduct) as conduct_chars,
  COUNT(asi.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '039f1542-7596-4671-8ed0-049b3b41cfc4'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct;
```

**Results:**
- Name: IGF-1 (Somatomedina C)
- Clinical Relevance: 1,656 characters ✅
- Patient Explanation: 1,468 characters ✅
- Conduct: 2,290 characters ✅
- Linked Articles: 13 total (4 newly added) ✅

---

## Database Changes

### Tables Modified
1. **articles** - 4 new articles inserted
2. **score_items** - 1 item updated (clinical content fields)
3. **article_score_items** - 4 new links created

### SQL Execution
- File: `/home/user/plenya/enrich_igf1_somatomedina.sql`
- Status: Successfully executed
- Transaction: COMMITTED

---

## Quality Assurance

### Content Quality
✅ All content in Portuguese (PT-BR)
✅ Clinical relevance covers pathophysiology, applications, and interpretation
✅ Patient explanation uses accessible language without medical jargon
✅ Conduct provides actionable clinical protocols
✅ Age-specific considerations included
✅ Diagnostic limitations acknowledged

### Scientific Rigor
✅ 4 peer-reviewed articles from high-impact journals
✅ Publication dates: 2014-2021 (current guidelines and research)
✅ Mix of article types: reviews, research articles, clinical guidelines
✅ Large multicenter study (15,014 subjects) for reference ranges
✅ Endocrine Society clinical practice guideline included

### Technical Compliance
✅ Correct table names (articles, article_score_items)
✅ Valid article_type values (review, research_article)
✅ publish_date as date type (not year)
✅ DOI uniqueness preserved
✅ Foreign key relationships maintained

---

## Sources

### Articles Added to Database
- [Insulin-Like Growth Factor-1 (IGF-1) and Its Monitoring in Medical Diagnostic and in Sports](https://www.mdpi.com/2218-273X/11/2/217)
- [Reference Intervals for Insulin-like Growth Factor-1 (IGF-I) From Birth to Senescence](https://academic.oup.com/jcem/article/99/5/1712/2537423)
- [IGF1 for the diagnosis of growth hormone deficiency in children and adolescents: a reappraisal](https://ec.bioscientifica.com/view/journals/ec/9/11/EC-20-0347.xml)
- [Acromegaly: An Endocrine Society Clinical Practice Guideline](https://academic.oup.com/jcem/article/99/11/3933/2836347)

### Additional References Consulted
- [Mayo Clinic Laboratories - IGF-1 Test Catalog](https://www.mayocliniclabs.com/test-catalog/overview/62750)
- [MedlinePlus - IGF-1 Test](https://medlineplus.gov/lab-tests/igf-1-insulin-like-growth-factor-1-test/)
- [Serum IGF-1 Age-Specific Reference Values for Serbia](https://pmc.ncbi.nlm.nih.gov/articles/PMC9206165/)
- [Consensus on acromegaly therapeutic outcomes - Nature Reviews Endocrinology](https://www.nature.com/articles/s41574-025-01148-2)

---

## Conclusion

The IGF-1 (Somatomedina C) score item has been successfully enriched with comprehensive, evidence-based clinical content supported by recent peer-reviewed literature. The content provides healthcare professionals with actionable guidance for IGF-1 interpretation across different clinical contexts (GH deficiency, acromegaly, treatment monitoring) while maintaining patient-centered language for patient education materials.

**Status:** ✅ COMPLETE AND VALIDATED
