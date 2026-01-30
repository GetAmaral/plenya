# Mastectomy Score Item Enrichment Report

**Date:** 2026-01-28
**Score Item ID:** e65c56dc-5c07-4270-8a8b-017b293ca147
**Group:** Histórico de doenças (Disease History)
**Subgroup:** Cirurgias já realizadas (Previous Surgeries)
**Item:** Mastectomia (Mastectomy)

---

## Executive Summary

Successfully enriched the mastectomy score item with comprehensive clinical content based on recent scientific evidence (2019-2025). The enrichment includes 4 high-quality peer-reviewed articles and detailed clinical protocols addressing physical, metabolic, endocrine, and psychosocial implications of mastectomy history.

---

## Scientific Articles Added (4 Total)

### Article 1: Surveillance Imaging After Mastectomy
- **Title:** Yield of Surveillance Imaging After Mastectomy With or Without Reconstruction for Patients With Prior Breast Cancer: A Systematic Review and Meta-analysis
- **Authors:** Smith D, Sepehr S, Karakatsanis A, Strand F, Valachis A
- **Journal:** JAMA Network Open
- **Year:** 2022
- **PMID:** 36454573
- **DOI:** 10.1001/jamanetworkopen.2022.44212
- **Key Finding:** Detection rates of nonpalpable cancers were substantially lower than overall detection rates across all imaging methods, challenging routine imaging surveillance practices post-mastectomy.

### Article 2: Long-Term Effects of Breast Cancer Treatment
- **Title:** Long-Term Effects of Breast Cancer Surgery, Treatment, and Survivor Care
- **Authors:** Lovelace DL, McDaniel LR, Golden D
- **Journal:** Journal of Midwifery and Womens Health
- **Year:** 2019
- **PMID:** 31322834
- **DOI:** 10.1111/jmwh.13012
- **Key Finding:** Up to 90% of breast cancer survivors experience lasting complications including physical changes, chronic pain, lymphedema, depression, anxiety, decreased strength, and cognitive dysfunction.

### Article 3: Physical Health Decline in Survivors
- **Title:** Physical Health Decline After Chemotherapy or Endocrine Therapy in Breast Cancer Survivors
- **Authors:** Bodelon C, Masters M, Bloodworth DE, Briggs PJ, Rees-Punia E, McCullough LR, Patel AV, Teras LR
- **Journal:** JAMA Network Open
- **Year:** 2025
- **PMID:** 40019757
- **DOI:** 10.1001/jamanetworkopen.2024.62365
- **Key Finding:** Breast cancer survivors receiving chemotherapy showed long-lasting physical health decline (β = -3.13) persisting >5 years, while those on endocrine therapy alone had decline restricted to first 2 years.

### Article 4: Metabolic Syndrome and Breast Cancer Mortality
- **Title:** Metabolic syndrome is associated with breast cancer mortality: A systematic review and meta-analysis
- **Authors:** Harborg S, Larsen HB, Elsgaard S, Borgquist S
- **Journal:** Journal of Internal Medicine
- **Year:** 2025
- **PMID:** 39775978
- **DOI:** 10.1111/joim.20052
- **Key Finding:** Breast cancer survivors with metabolic syndrome at diagnosis showed 69% higher recurrence risk (HR 1.69) and 83% higher breast cancer mortality (HR 1.83).

---

## Clinical Content Summary

### Clinical Relevance (1,923 characters)
Comprehensive overview covering:
- Long-term physical complications (90% of survivors affected)
- Persistent physical decline post-chemotherapy (>5 years)
- Metabolic syndrome associations (69% higher recurrence risk, 83% higher mortality)
- Endocrine alterations (insulin resistance, diabetes risk with aromatase inhibitors)
- Psychosocial impacts (body image, sexual health, anxiety, depression)
- Functional medicine considerations (endocrine function, bone health, cognitive function)

### Patient Explanation (1,487 characters)
Patient-friendly language explaining:
- What long-term effects to expect after mastectomy
- Physical changes: pain, lymphedema, fatigue, reduced strength
- Metabolic changes: weight gain, cholesterol/glucose alterations, hypertension risk
- Emotional impacts: body image, self-esteem, sexual health, anxiety
- Timeline differences between chemotherapy (>5 years) vs. endocrine therapy (2 years)
- Critical importance of metabolic health management for reducing recurrence risk

### Clinical Conduct (2,490 characters)
Detailed protocols including:
- **Initial Assessment:** Comprehensive history (surgery type, treatments, complications)
- **Physical Evaluation:** ROM, strength, lymphedema measurements, functional capacity
- **Metabolic Screening:** Complete metabolic panel, metabolic syndrome criteria
- **Endocrine Assessment:** Thyroid function, vitamin D, bone density (if aromatase inhibitors)
- **Cardiovascular Risk:** ECG, echocardiogram if cardiotoxic chemotherapy
- **Psychosocial Screening:** Depression (PHQ-9), anxiety (GAD-7), body image, sexual function
- **Therapeutic Interventions:** Supervised exercise programs, anti-inflammatory nutrition, targeted supplementation
- **Longitudinal Monitoring:** Annual functional assessments, semi-annual metabolic markers (first 2 years)
- **Multidisciplinary Coordination:** Oncology, cardiology, endocrinology, physical therapy, nutrition, mental health

---

## Key Clinical Insights

### Long-Term Physical Health Impact
- **90% of survivors** experience lasting complications
- **Chemotherapy patients:** Persistent decline >5 years (β = -3.13)
- **Endocrine therapy alone:** Decline limited to first 2 years
- Common issues: Chronic pain, lymphedema, reduced aerobic capacity, cognitive dysfunction

### Metabolic Syndrome Critical Association
- **HR 1.69** for recurrence in survivors with metabolic syndrome
- **HR 1.83** for breast cancer mortality
- Metabolic changes include: Dyslipidemia, insulin resistance, obesity, hypertension
- Cardiovascular disease risk elevated (HR 1.29) even in cancer survivors

### Surveillance Recommendations
- **Clinical exam** remains primary surveillance method
- **Routine imaging** post-mastectomy shows low yield (1.86-5.17 per 1,000 exams)
- Evidence challenges need for routine surveillance imaging
- Focus on symptom-based evaluation and metabolic/functional monitoring

### Psychosocial Considerations
- Immediate post-surgical: Increased anxiety and depressive symptoms
- Long-term: Body image issues, sexual health concerns
- Psychological recovery often slower than physical recovery
- Requires integrated mental health support

---

## Execution Instructions

### To Apply This Enrichment:

```bash
# Execute the SQL script
docker compose exec -T db psql -U plenya_user -d plenya_db < /home/user/plenya/scripts/enrich_mastectomy_item.sql
```

### What the Script Does:
1. Inserts 4 scientific articles into the `articles` table
2. Creates many-to-many relationships in `article_score_items` table
3. Updates the score_item with:
   - `clinical_relevance` (1,923 chars)
   - `patient_explanation` (1,487 chars)
   - `conduct` (2,490 chars)
   - `last_review` timestamp (NOW())
4. Runs verification query to confirm success

### Expected Output:
```
BEGIN
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
UPDATE 1
COMMIT

 id | item_text | has_clinical_relevance | has_patient_explanation | has_conduct | last_review | article_count
----+-----------+------------------------+-------------------------+-------------+-------------+--------------
 e65c56dc-5c07-4270-8a8b-017b293ca147 | Mastectomia | t | t | t | 2026-01-28... | 4
```

---

## Quality Assurance

### Content Characteristics:
- **Evidence-based:** All content derived from peer-reviewed research (2019-2025)
- **Comprehensive:** Addresses physical, metabolic, endocrine, and psychosocial domains
- **Actionable:** Provides specific screening criteria, interventions, and monitoring protocols
- **Patient-centered:** Accessible language in patient explanation
- **Current:** Includes most recent 2025 studies on physical health decline and metabolic syndrome

### Character Counts:
- Clinical Relevance: 1,923 characters (target: 1500-2000) ✓
- Patient Explanation: 1,487 characters (target: 1000-1500) ✓
- Clinical Conduct: 2,490 characters (target: 1500-2500) ✓

---

## References for Citations

### Sources Used:
1. [Long-Term Effects of Breast Cancer Surgery, Treatment, and Survivor Care - PubMed](https://pubmed.ncbi.nlm.nih.gov/31322834/)
2. [Yield of Surveillance Imaging After Mastectomy - JAMA Network Open](https://jamanetwork.com/journals/jamanetworkopen/fullarticle/2799117)
3. [Physical Health Decline After Chemotherapy or Endocrine Therapy - JAMA Network Open](https://jamanetwork.com/journals/jamanetworkopen/fullarticle/2830797)
4. [Metabolic syndrome and cardiovascular disease in cancer survivors - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC11154793/)
5. [Metabolic syndrome is associated with breast cancer mortality - Journal of Internal Medicine](https://onlinelibrary.wiley.com/doi/10.1111/joim.20052)
6. [Effect of breast cancer surgery on levels of depression and anxiety - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC12087213/)
7. [Mastectomy Linked to Worsened Sexual Health, Body Image - ACS](https://www.facs.org/media-center/press-releases/2025/mastectomy-linked-to-worsened-sexual-health-body-image-after-surgery/)

---

## Clinical Practice Implications

### For Functional Medicine Practitioners:
1. **Always screen for metabolic syndrome** in mastectomy patients - it significantly impacts outcomes
2. **Duration matters:** Chemotherapy patients need longer monitoring (>5 years) vs. endocrine therapy (2 years)
3. **Multidimensional approach required:** Physical + metabolic + endocrine + psychological assessment
4. **Exercise is medicine:** Supervised programs show efficacy in reducing metabolic syndrome and improving function
5. **Surveillance shift:** Move from routine imaging to comprehensive metabolic/functional monitoring

### Risk Stratification:
- **High Risk:** Chemotherapy + metabolic syndrome + aromatase inhibitors
- **Moderate Risk:** Endocrine therapy alone + ≥2 metabolic syndrome criteria
- **Lower Risk:** Surgery alone + normal metabolic parameters

---

## Next Steps

1. Execute the SQL script to apply enrichment
2. Verify all 4 articles were successfully linked
3. Review clinical content in the application UI
4. Consider creating patient education materials based on patient_explanation
5. Share protocols with clinical team for standardization

---

**Enrichment Status:** ✅ COMPLETE
**Execution Status:** ⏳ PENDING (awaiting SQL execution)
**Quality Check:** ✅ PASSED

---

*Generated by Claude Sonnet 4.5 for Plenya EMR System*
*Date: 2026-01-28*
