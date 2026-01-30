# PROLACTINA - MULHERES: Quick Summary

## ✅ ENRICHMENT COMPLETED SUCCESSFULLY

**Date:** 2026-01-29
**Score Item ID:** `7aca1ff2-49b7-431d-aa69-f02d2a686bbc`

---

## Results

### Content Added
- ✅ Clinical Relevance: **1,386 characters**
- ✅ Patient Explanation: **1,290 characters**
- ✅ Conduct Guidelines: **2,097 characters**
- ✅ Last Review: Updated to current date

### Scientific Articles
- ✅ **4 articles** from 2023-2024 added
- ✅ **4 article-score_item links** created
- ✅ Total links for this item: **13** (4 new + 9 pre-existing)

---

## Key Clinical Points

### Reference Values
- **Normal (non-pregnant women):** 2-29 ng/mL
- **Prolactinoma threshold:** >200 ng/mL
- **Macroprolactinemia:** 10-46% of hyperprolactinemia cases

### Prevalence
- **15-20%** of infertile women have hyperprolactinemia
- **10x higher** than general population

### Treatment
- **Cabergoline** is first-line therapy
- **80% success rate** in fertility restoration
- **1/3 of patients** achieve definitive cure

---

## Clinical Protocols (Stratified by Level)

### Normal (2-29 ng/mL)
→ Routine follow-up

### Mild (30-50 ng/mL)
→ Repeat test + investigate causes + macroprolactin test + thyroid function

### Moderate (51-100 ng/mL)
→ Full investigation + pituitary MRI + endocrinology referral

### Very High (>100 ng/mL)
→ Priority endocrinology referral + MRI + consider treatment start

---

## Files Generated

1. **SQL Script:** `/home/user/plenya/scripts/enrich_prolactina_mulheres.sql`
2. **Full Report:** `/home/user/plenya/PROLACTINA-MULHERES-ENRICHMENT-REPORT.md`
3. **Quick Summary:** `/home/user/plenya/PROLACTINA-MULHERES-QUICK-SUMMARY.md`

---

## Verification Query

```sql
SELECT id, name,
       LENGTH(clinical_relevance) as clinical,
       LENGTH(patient_explanation) as patient,
       LENGTH(conduct) as conduct,
       last_review
FROM score_items
WHERE id = '7aca1ff2-49b7-431d-aa69-f02d2a686bbc';
```

**Status:** ✅ All fields populated correctly
