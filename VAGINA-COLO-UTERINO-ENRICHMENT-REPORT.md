# Enrichment Report: Vagina e Colo Uterino

## Executive Summary

Successfully enriched the score item "Vagina e Colo Uterino" (ID: `1ff1e7f8-5b15-4f0c-b76b-806c1d6ff1fd`) with comprehensive clinical content in Portuguese and 4 peer-reviewed scientific articles.

**Completion Date:** 2026-01-29

---

## Clinical Content Character Counts

| Field | Character Count | Target Range | Status |
|-------|----------------|--------------|--------|
| **clinical_relevance** | 1,528 | 1,500-2,000 | ✅ Within range |
| **patient_explanation** | 1,370 | 1,000-1,500 | ✅ Within range |
| **conduct** | 2,229 | 1,500-2,500 | ✅ Within range |
| **last_review** | 2026-01-29 | Current date | ✅ Updated |

**All fields meet the specified requirements.**

---

## Linked Scientific Articles

### Article 1: Self-Collected Vaginal Specimens (2025) ⭐ MOST RECENT
- **Title:** Self-Collected Vaginal Specimens for HPV Testing: Recommendations From the Enduring Consensus Cervical Cancer Screening and Management Guidelines Committee
- **Authors:** Wentzensen N, Perkins RB, Guido RS, Castle PE, Massad LS, Einstein MH, Smith RA, Waxman AG, Khan MJ, Chelmow D, Fontham ETH, Saraiya M, Conageski C, Feldman S, Tedeschi C, Stier EA, Mayeaux EJ, Saslow D, Schiffman M
- **Journal:** J Low Genit Tract Dis
- **PubMed ID:** 39982254
- **Publication Date:** 2025-02-21
- **Article Type:** clinical_trial
- **Focus:** FDA-approved self-collection for HPV testing, expanding access to screening

### Article 2: Primary HPV Testing Technologies (2023)
- **Title:** Primary Human Papillomavirus Testing and Other New Technologies for Cervical Cancer Screening
- **Authors:** Einstein MH, Zhou N, Gabor L, Sahasrabuddhe VV
- **Journal:** Obstet Gynecol
- **PubMed ID:** 37708516
- **Publication Date:** 2023-10-01
- **Article Type:** review
- **Focus:** Transition to HPV-based screening, p16/Ki-67 dual stain, HPV methylation testing

### Article 3: Cervical Cancer Screening Recommendations (2023)
- **Title:** Cervical Cancer Screening Recommendations: Now and for the Future
- **Authors:** Rayner M, Welp A, Stoler MH, Cantrell LA
- **Journal:** Healthcare (Basel)
- **PubMed ID:** 37628471
- **Publication Date:** 2023-08-15
- **Article Type:** review
- **Focus:** Global guidelines comparison (WHO, USPSTF, ACOG, ACS, ASCO)

### Article 4: Colposcopy Systematic Approach (2017)
- **Title:** A systematic approach to colposcopic examination
- **Authors:** International Agency for Research on Cancer (IARC)
- **Journal:** IARC Technical Publications No. 45 - Colposcopy and Treatment of Cervical Precancer
- **PubMed ID:** 29897450
- **Publication Date:** 2017-12-01
- **Article Type:** lecture
- **Focus:** Technical guidelines for colposcopic examination technique and training

---

## Clinical Content Summary

### Clinical Relevance (1,528 chars)
Covers:
- Fundamental role in women's preventive health
- Visual inspection with speculum, Pap smear, HPV testing
- 2024-2025 updated guidelines (USPSTF, ACS, ASCCP)
- Screening intervals by age group
- HPV genotypes 16/18 responsibility for 70% cervical cancer
- Colposcopy indications (ASCUS, LGSIL, HGSIL)
- Proper examination technique (acetic acid, Lugol's solution)
- FDA-approved self-collection advances (2024)

### Patient Explanation (1,370 chars)
Covers:
- Simple explanation of Pap smear and preventive exam
- HPV virus role in cervical cancer
- Screening schedule recommendations
- Self-collection option for patients with barriers
- Colposcopy procedure when needed
- Importance of HPV vaccination + screening
- Cancer prevention through early detection

### Conduct (2,229 chars)
Detailed protocols including:
1. Age 21-29: Cytology every 3 years
2. Age 30-65: Primary HPV every 5 years OR co-test every 5 years OR cytology every 3 years
3. Self-collection option with 3-year retest if negative
4. ASCUS management with HPV genotyping
5. LGSIL/HGSIL colposcopy referral criteria
6. Colposcopy technique (speculum selection, acetic acid, transformation zone types)
7. Endocervical curettage indications
8. Molecular tests (p16/Ki-67, HPV methylation)
9. Post-treatment follow-up (25 years or until age 65)
10. Screening cessation criteria (≥65 years with adequate negative screening)
11. Vaccinated women follow same guidelines
12. Documentation requirements

---

## Key Clinical Guidelines Covered

### USPSTF 2024-2025 Updates
- Primary HPV testing every 5 years (ages 30-65)
- Self-collected HPV tests now included
- Cytology every 3 years (ages 21-29)
- Co-testing remains acceptable alternative

### ASCCP 2024 Guidelines
- Risk-based management approach
- CINtec PLUS Cytology triage test
- HPV 16/18 direct colposcopy referral
- Extended genotyping protocols

### American Cancer Society 2025
- Start screening at age 25 (updated from 21)
- Primary HPV testing preferred
- Self-collection guidance
- Screening exit criteria

### FDA Approvals 2024
- Roche cobas HPV test with FLOQSwab or Evalyn Brush
- BD Onclarity HPV test with FLOQSwab
- Abbott Alinity m with Evalyn Brush or Qvintip swab

---

## Technical Implementation

### Database Tables Used
- **score_items:** Updated clinical fields and last_review
- **articles:** Inserted 4 new peer-reviewed articles
- **article_score_items:** Created 4 junction table links

### SQL Script Location
`/home/user/plenya/scripts/enrich_vagina_colo_uterino.sql`

### Execution Method
```bash
cat /home/user/plenya/scripts/enrich_vagina_colo_uterino.sql | \
  docker compose exec -T db psql -U plenya_user -d plenya_db
```

### Database Schema Compliance
- ✅ Table name: `articles` (not scientific_articles)
- ✅ Junction table: `article_score_items` (not score_item_articles)
- ✅ Column: `pm_id` (not pmid)
- ✅ Column: `publish_date` as date type (not year)
- ✅ Column: `abstract` (not summary)
- ✅ Column: `last_review` (not last_reviewed)
- ✅ Article types: review, clinical_trial, lecture (valid enum values)

---

## Quality Assurance

### Content Validation
- [x] All character counts within specified ranges
- [x] Portuguese language clinical content
- [x] Evidence-based guidelines (2020-2025)
- [x] Proper medical terminology
- [x] Patient-friendly explanations
- [x] Detailed clinical protocols

### Article Validation
- [x] 4 peer-reviewed articles linked
- [x] PubMed IDs verified
- [x] Publication dates 2017-2025
- [x] Mix of article types (review, clinical_trial, lecture)
- [x] Relevant to vaginal/cervical examination and screening
- [x] Portuguese abstracts included

### Database Validation
- [x] No duplicate articles (WHERE NOT EXISTS clause)
- [x] No duplicate junction links (WHERE NOT EXISTS clause)
- [x] Proper UUID format for score_item_id
- [x] Transaction wrapped in BEGIN/COMMIT
- [x] Verification queries included

---

## Clinical Focus Areas

### Primary Topics
1. **Cervical Cancer Screening**
   - Pap smear technique
   - HPV testing (primary and co-testing)
   - Self-collection methods
   - Screening intervals and age guidelines

2. **Colposcopy**
   - Indications and referral criteria
   - Systematic examination approach
   - Transformation zone assessment
   - Biopsy techniques

3. **HPV Management**
   - Genotyping (especially 16/18)
   - Risk stratification
   - Molecular markers (p16/Ki-67, methylation)
   - Vaccination integration

4. **Clinical Guidelines**
   - USPSTF recommendations
   - ASCCP risk-based management
   - ACS screening updates
   - WHO global guidelines

### Patient-Centered Care
- Barrier reduction through self-collection
- Trauma-informed approaches
- Shared decision-making
- Health equity considerations

---

## Sources for User Response

All web search results used to inform this enrichment:

### Guidelines and Official Sources
- [New Cervical Cancer Screening Guideline Aims to Reduce Deaths | American Cancer Society](https://www.cancer.org/research/acs-research-news/new-cervical-cancer-screening-guidelines-include-self-collection-for-hpv.html)
- [Recommendation: Cervical Cancer: Screening | USPSTF](https://www.uspreventiveservicestaskforce.org/uspstf/recommendation/cervical-cancer-screening)
- [Updated Cervical Cancer Screening Guidelines | ACOG](https://www.acog.org/clinical/clinical-guidance/practice-advisory/articles/2021/04/updated-cervical-cancer-screening-guidelines)
- [What Are the ASCCP Guidelines?](https://www.asccp.org/guidelines/)

### Research Articles
- [Primary Human Papillomavirus Testing - PubMed](https://pubmed.ncbi.nlm.nih.gov/37708516/)
- [Cervical Cancer Screening Recommendations: Now and for the Future - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC10454304/)
- [Self-Collected Vaginal Specimens for HPV Testing - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC11939108/)
- [A systematic approach to colposcopic examination - NCBI Bookshelf](https://www.ncbi.nlm.nih.gov/books/NBK568388/)

### Technical Resources
- [Colposcopy - StatPearls - NCBI Bookshelf](https://www.ncbi.nlm.nih.gov/books/NBK564514/)
- [It Is Time to Switch to Primary HPV Screening | AAFP](https://www.aafp.org/pubs/afp/issues/2024/0100/editorial-hpv-screening-cervical-cancer.html)

---

## Conclusion

The "Vagina e Colo Uterino" score item has been successfully enriched with comprehensive, evidence-based clinical content reflecting the most current 2024-2025 guidelines for cervical cancer screening, including breakthrough advances in self-collection technology. The enrichment includes 4 high-quality scientific articles spanning 2017-2025, covering both foundational techniques and cutting-edge screening innovations.

**Status:** ✅ COMPLETE
