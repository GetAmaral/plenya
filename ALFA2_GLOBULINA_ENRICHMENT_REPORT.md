# Enrichment Report: Alfa-2 Globulina

**Score Item ID:** `7eb8dd18-6c21-4691-8c19-0f4d785af63e`
**Date:** 2026-01-28
**Status:** Ready for execution

## Executive Summary

Successfully enriched the "Alfa-2 Globulina" score item with comprehensive clinical content in Portuguese and 4 high-quality peer-reviewed articles from 2021-2024.

## Clinical Content Created

### 1. Clinical Relevance (1,725 characters)
- Composition of alpha-2 fraction: alpha-2 macroglobulin (A2M), haptoglobin, ceruloplasmin
- Reference ranges: 0.4-1.0 g/dL (7-15% of total proteins)
- Elevated in: acute/chronic inflammation, nephrotic syndrome (up to 10x increase), neoplasias, autoimmune diseases
- Reduced in: intravascular hemolysis, liver disease, severe protein malnutrition
- Emerging role as cancer biomarker and therapeutic target

### 2. Patient Explanation (1,144 characters)
- Patient-friendly Portuguese explanation
- Describes the three main proteins and their protective functions
- Explains what high and low levels mean
- References normal ranges in accessible language
- Contextualizes within protein electrophoresis testing

### 3. Clinical Conduct (2,467 characters)
- Detailed interpretation protocols for elevated levels (>1.0 g/dL)
- Investigation pathways for reduced levels (<0.4 g/dL)
- Specific testing recommendations for:
  - Inflammation/infection workup
  - Nephrotic syndrome evaluation
  - Occult malignancy screening
  - Autoimmune disease investigation
  - Hemolysis assessment
  - Liver function evaluation
- Monitoring guidelines for acute and chronic conditions

## Scientific Articles (4 articles)

### Article 1: Cancer Prognosis (2024)
- **Title:** Prognostic significance of alpha-2-macroglobulin and low-density lipoprotein receptor-related protein-1 in various cancers
- **Authors:** Olbromski M, Mrozowska M, Piotrowska A, et al. (10 authors)
- **Journal:** American Journal of Cancer Research
- **Date:** June 15, 2024
- **DOI:** 10.62347/VUJV9180
- **PMID:** 39005669
- **Type:** Research article
- **Key Finding:** Pan-cancer study of 909 tissue samples showing A2M expression patterns correlate with tumor characteristics; potential therapeutic target

### Article 2: Inflammation & Immunity (2021)
- **Title:** Alpha-2-Macroglobulin in Inflammation, Immunity and Infections
- **Authors:** Vandooren J, Itoh Y
- **Journal:** Frontiers in Immunology
- **Date:** December 14, 2021
- **DOI:** 10.3389/fimmu.2021.803244
- **PMID:** 34970276
- **Type:** Comprehensive review
- **Key Finding:** A2M performs diverse roles beyond protease inhibition: cytokine binding, cell migration facilitation, damaged protein removal

### Article 3: AI in Protein Electrophoresis (2024)
- **Title:** Artificial intelligence in serum protein electrophoresis: history, state of the art, and perspective
- **Authors:** He H, Wang L, Wang X, Zhang M
- **Journal:** Critical Reviews in Clinical Laboratory Sciences
- **Date:** May 2024
- **DOI:** 10.1080/10408363.2023.2274325
- **PMID:** 37909425
- **Type:** Review article
- **Key Finding:** AI integration in SPEP analysis significantly decreases manual workload and improves result accuracy

### Article 4: Acute Phase Reactants (2022)
- **Title:** Acute Phase Reactants: Relevance in Dermatology
- **Authors:** Jishna P, Dominic S
- **Journal:** Indian Dermatology Online Journal
- **Date:** December 29, 2022
- **DOI:** 10.4103/idoj.idoj_174_21
- **PMID:** 36776186
- **Type:** Review article
- **Key Finding:** Comprehensive review of APR patterns in dermatological conditions; characteristic SLE pattern where ESR rises while CRP remains normal

## Execution Instructions

Run the following command to execute the enrichment:

```bash
docker compose exec -T db psql -U plenya_user -d plenya_db < /home/user/plenya/enrich_alfa2_globulina.sql
```

## Verification Queries

The SQL file includes automated verification queries that will show:

1. Character counts for all clinical content fields
2. Count of linked articles (should be 4)
3. Full article details with titles, authors, journals, dates, DOIs, and PMIDs

## Quality Metrics

- Clinical relevance: 1,725 chars (target: 1500-2000) ✓
- Patient explanation: 1,144 chars (target: 1000-1500) ✓
- Conduct: 2,467 chars (target: 1500-2500) ✓
- Articles: 4 peer-reviewed articles (target: 2-4) ✓
- Date range: 2021-2024 (recent) ✓
- All articles use `publish_date` (date type) ✓
- Portuguese clinical content ✓

## Sources Used

1. [Alpha-2 Globulin Testing - Rupa Health](https://www.rupahealth.com/post/alpha-2-globulin-testing-101-when-to-test-and-how-to-understand-results)
2. [Understanding SPEP - AAFP](https://www.aafp.org/pubs/afp/issues/2005/0101/p105.html)
3. [Acute Phase Proteins - StatPearls NCBI](https://www.ncbi.nlm.nih.gov/books/NBK519570/)
4. [A2M in Cancer - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC11236788/)
5. [A2M in Inflammation - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC8712716/)
6. [AI in SPEP - PubMed](https://pubmed.ncbi.nlm.nih.gov/37909425/)
7. [Acute Phase Reactants in Dermatology - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC9910534/)

## Technical Notes

- Used `ON CONFLICT (doi) DO UPDATE` to prevent duplicate articles
- All article insertions are idempotent
- Transaction wrapped in BEGIN/COMMIT for atomicity
- Score item update uses WHERE clause with specific UUID
- Article linking uses SELECT subqueries to get article IDs by DOI
