# NT-proBNP (50-75 anos) - Enrichment Report

**Score Item ID:** 7998e69e-a0e0-488b-bcc3-9da32e59adfb
**Date:** 2026-01-29
**Status:** ✅ COMPLETE

---

## Executive Summary

Successfully enriched the NT-proBNP (50-75 anos) score item with comprehensive clinical content based on latest ESC 2023/ACC 2022 guidelines and 4 high-quality peer-reviewed articles (2023-2025).

### Key Focus Areas
- Age-specific NT-proBNP cutoffs for 50-75 year age group
- Heart failure diagnosis algorithms (ambulatory vs emergency)
- Rule-out vs rule-in thresholds
- Impact of atrial fibrillation and renal dysfunction
- ESC Heart Failure Association 2023 consensus recommendations

---

## Clinical Content Character Counts

| Field | Character Count | Target Range | Status |
|-------|----------------|--------------|--------|
| **clinical_relevance** | 1,911 | 1,500-2,000 | ✅ |
| **patient_explanation** | 1,385 | 1,000-1,500 | ✅ |
| **conduct** | 1,713 | 1,500-2,500 | ✅ |

All fields meet the required character count ranges.

---

## Linked Scientific Articles (n=4)

### 1. ESC HFA 2023 Consensus Statement
**PMID:** 37712339
**Title:** Practical algorithms for early diagnosis of heart failure and heart stress using NT-proBNP: A clinical consensus statement from the Heart Failure Association of the ESC
**Authors:** Antoni Bayes-Genis, Kieran F Docherty, Mark C Petrie, James L Januzzi, Christian Mueller, et al.
**Journal:** European Journal of Heart Failure
**Year:** 2023
**DOI:** 10.1002/ejhf.3036
**Type:** Review/Consensus

**Key Findings:**
- Age-adjusted rule-in thresholds: ≥250 pg/mL (50-75 years) for ambulatory setting
- Emergency department: ≥900 pg/mL for acute HF in 50-75 age group
- Introduced FIND-HF acronym (Fatigue, Increased water, Natriuretic peptide, Dyspnoea)
- No additional adjustments needed for AF, renal dysfunction, obesity in ED setting

---

### 2. Community Diagnostic Accuracy Study
**PMID:** 40717271
**Title:** Age-adjusted natriuretic peptide thresholds for a diagnosis of heart failure in the community: Diagnostic accuracy study
**Authors:** Clare J Taylor, Kathryn S Taylor, Nicholas R Jones, Jose M Ordóñez-Mena, Antoni Bayes-Genis, FD Richard Hobbs
**Journal:** ESC Heart Failure
**Year:** 2025
**DOI:** 10.1002/ehf2.15383
**Type:** Research Article

**Key Findings:**
- Analyzed 155,347 patients with NT-proBNP tests
- Age-adjusted thresholds increased specificity from 50.0% to 67.8% in 50-74 age group
- Sensitivity: 88.5% for age-adjusted threshold (≥250 pg/mL)
- Obesity-adjusted thresholds necessary to avoid missing HF in patients with elevated BMI

---

### 3. Optimal Thresholds in Elderly Patients
**PMID:** 38923835
**Title:** Setting the optimal threshold of NT-proBNP and BNP for the diagnosis of heart failure in patients over 75 years
**Authors:** Emmanuelle Berthelot, Minh Tam Bailly, Xenia Cerchez Lehova, et al.
**Journal:** ESC Heart Failure
**Year:** 2024
**DOI:** 10.1002/ehf2.14894
**Type:** Research Article

**Key Findings:**
- Optimal threshold for acute HF: 1748 ng/L in patients ≥75 years
- Among patients >85 years: threshold of 2235 pg/mL (84% PPV)
- Provides context for upper age boundary of 50-75 range

---

### 4. NT-proBNP in Atrial Fibrillation
**PMID:** 37320999
**Title:** NT-proBNP cut-off value for ruling out heart failure in atrial fibrillation patients - A prospective clinical study
**Authors:** Cecilie Budolfsen, Anders Sjørslev Schmidt, Kasper Glerup Lauridsen, et al.
**Journal:** American Journal of Emergency Medicine
**Year:** 2023
**DOI:** 10.1016/j.ajem.2023.05.041
**Type:** Clinical Trial

**Key Findings:**
- Median NT-proBNP in AF patients: 2577 ng/L (IQR: 1185-5438)
- 21% of AF patients had heart failure
- Exploratory analysis (≤75 years): cut-off of 1398 ng/L for rule-out
- AF present in ~30% of newly presenting HF patients

---

## Clinical Content Summary

### Clinical Relevance Highlights
1. **Age-specific cutoffs:** ≥250 pg/mL (ambulatory), ≥900 pg/mL (emergency) for 50-75 years
2. **Rule-out threshold:** <300 pg/mL (98% NPV across all ages)
3. **Grey zone:** 300-900 pg/mL requires additional testing
4. **Confounding factors:** AF, renal dysfunction (GFR <60), obesity
5. **Prognostic value:** >30% reduction = favorable response to therapy
6. **High-risk threshold:** Persistent values >1000 pg/mL

### Patient Explanation Highlights
- Simple language explanation of what NT-proBNP measures
- Clear interpretation thresholds (below 250, 250-900, above 900)
- Information about factors that affect results (AF, kidney problems, obesity)
- Guidance on comparing serial results
- Importance of clinical correlation

### Conduct Protocol Highlights
1. **Ambulatory setting:**
   - <250 pg/mL: HF unlikely
   - 250-500 pg/mL: grey zone, order echo
   - >500 pg/mL: refer to cardiology within 2 weeks

2. **Emergency setting:**
   - <300 pg/mL: acute HF unlikely (98% NPV)
   - 300-900 pg/mL: grey zone, bedside echo
   - >900 pg/mL: acute HF probable, initiate therapy

3. **Special adjustments:**
   - AF: increase threshold 30-50% or use ≥1400 pg/mL
   - Obesity (BMI >30): reduce thresholds by 50%
   - Renal dysfunction (GFR 30-60): interpret with caution

4. **Post-diagnosis management:**
   - Recheck every 3-6 months
   - Target: >30% reduction in 3 months
   - Persistent >1000 pg/mL: optimize GDMT therapy

---

## Database Verification

```sql
-- Score item verification
SELECT id, code, name,
       LENGTH(clinical_relevance) as clin_chars,
       LENGTH(patient_explanation) as pat_chars,
       LENGTH(conduct) as cond_chars
FROM score_items
WHERE id = '7998e69e-a0e0-488b-bcc3-9da32e59adfb';

-- Article links verification
SELECT COUNT(*) as total_articles,
       COUNT(CASE WHEN pm_id IS NOT NULL THEN 1 END) as with_pmid
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '7998e69e-a0e0-488b-bcc3-9da32e59adfb';
```

**Results:**
- ✅ Score item exists and is enriched
- ✅ 4 scientific articles linked (all with PMIDs)
- ✅ All character counts within target ranges
- ✅ Last review timestamp updated

---

## Scientific Evidence Quality

### Evidence Level
- **1 Consensus Statement** (ESC HFA 2023) - High authority
- **2 Diagnostic Accuracy Studies** (2024-2025) - Recent evidence
- **1 Clinical Trial** (2023) - Prospective study on AF patients

### Publication Timeline
- 2025: 1 article (most recent community study)
- 2024: 1 article (elderly thresholds)
- 2023: 2 articles (ESC consensus + AF study)

All articles published within the last 2-3 years, ensuring current evidence-based recommendations.

---

## Key Clinical Takeaways

1. **Age matters:** 50-75 years requires intermediate cutoffs (between <50 and >75)
2. **Context matters:** Different thresholds for ambulatory vs emergency settings
3. **Comorbidities matter:** AF and renal dysfunction significantly affect interpretation
4. **Grey zone exists:** 300-900 pg/mL requires additional diagnostic workup
5. **Serial monitoring:** Useful for assessing therapy response and prognosis

---

## Implementation Notes

### SQL Execution
```bash
# Enrichment executed via Docker PostgreSQL
cat scripts/enrich_nt_probnp_50_75.sql | docker compose exec -T db psql -U plenya_user -d plenya_db

# Content update for character counts
cat scripts/update_nt_probnp_50_75_content.sql | docker compose exec -T db psql -U plenya_user -d plenya_db
```

### Database Schema Compliance
- ✅ Table name: `articles` (not scientific_articles)
- ✅ Junction table: `article_score_items` (not score_item_articles)
- ✅ Column: `pm_id` (not pmid)
- ✅ Column: `publish_date` as date type (not year)
- ✅ No URL column used
- ✅ No created_at in junction table
- ✅ Valid article_type enum values

---

## Sources

- [ESC HFA Practical Algorithms 2023](https://pubmed.ncbi.nlm.nih.gov/37712339/)
- [Age-adjusted Thresholds Community Study 2025](https://pubmed.ncbi.nlm.nih.gov/40717271/)
- [Optimal Thresholds Elderly 2024](https://pubmed.ncbi.nlm.nih.gov/38923835/)
- [NT-proBNP in Atrial Fibrillation 2023](https://pubmed.ncbi.nlm.nih.gov/37320999/)
- [ESC HFA Guidelines 2023 - Wiley](https://onlinelibrary.wiley.com/doi/10.1002/ejhf.3036)
- [Community Diagnosis Study - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC12450762/)
- [Elderly Thresholds Study - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC11424345/)

---

## Completion Checklist

- [x] Search for 2-4 peer-reviewed articles (2020-2025)
- [x] Found 4 high-quality articles (2023-2025)
- [x] Create clinical_relevance (1500-2000 chars) - **1911 chars**
- [x] Create patient_explanation (1000-1500 chars) - **1385 chars**
- [x] Create conduct (1500-2500 chars) - **1713 chars**
- [x] Generate SQL with correct schema
- [x] Execute SQL via Docker
- [x] Verify character counts
- [x] Verify article links (4 articles with PMIDs)
- [x] Document sources with hyperlinks

**Status: COMPLETE ✅**

---

*Report generated: 2026-01-29*
*Score Item: NT-proBNP (50-75 anos)*
*Database: Plenya EMR PostgreSQL*
