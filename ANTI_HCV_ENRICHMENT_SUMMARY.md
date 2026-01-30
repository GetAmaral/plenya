# Enrichment Summary: Hepatite C - Anti-HCV

**Score Item ID:** 73411e66-c180-46ad-8f5b-7d67b26ef877
**Generated:** 2026-01-29
**Status:** Ready for execution

---

## Scientific Articles Added (4 articles)

### 1. Global HCV Elimination Progress Report (2025)
- **Title:** A 2024 global report on national policies, programmes, and progress towards hepatitis C elimination: findings from 33 hepatitis elimination profiles
- **Authors:** Hiebert-Suwondo L, Manning J, Tohme RA, et al.
- **Journal:** Lancet Gastroenterol Hepatol
- **Published:** May 20, 2025
- **DOI:** 10.1016/S2468-1253(25)00068-8
- **Type:** Review
- **Key Focus:** Analysis of 33 countries representing 73% of global HCV cases; only 30% met WHO 2025 diagnosis targets

### 2. HCV Epidemiology and Global Elimination (2025)
- **Title:** Hepatitis C Virus: Epidemiological Challenges and Global Strategies for Elimination
- **Authors:** Toma D, Anghel L, Patraș D, Ciubară A
- **Journal:** Viruses
- **Published:** July 31, 2025
- **DOI:** 10.3390/v17081069
- **Type:** Review
- **Key Focus:** 1 million annual new infections, 242,000 deaths in 2022; DAAs achieve >95% cure rates; prison prevalence 17.7%

### 3. SASLT HCV Treatment Guidelines 2024
- **Title:** SASLT guidelines: Update in treatment of hepatitis C virus infection, 2024
- **Authors:** Alghamdi AS, Alghamdi H, Alserehi HA, et al.
- **Journal:** Saudi J Gastroenterol
- **Published:** January 3, 2024
- **DOI:** 10.4103/sjg.sjg_333_23
- **Type:** Clinical Guideline
- **Key Focus:** Pangenotypic DAAs as first-line treatment; all active HCV patients should receive treatment without delay

### 4. DAA Treatment and Immune Restoration (2025)
- **Title:** Cure of chronic hepatitis C virus infection after DAA treatment only partially restores the functional capacity of exhausted T cell subsets: a systematic review
- **Authors:** Frontiers in Immunology authors
- **Journal:** Front Immunol
- **Published:** January 15, 2025
- **DOI:** 10.3389/fimmu.2025.1546915
- **Type:** Systematic Review
- **Key Focus:** DAAs achieve >95% cure rates but T cell restoration is partial; importance of early treatment

---

## Clinical Content Summary

### Clinical Relevance (1,958 characters)
Comprehensive overview covering:
- Anti-HCV test mechanism and clinical utility
- Global epidemiology (58M chronic infections, 242K annual deaths)
- DAA revolution (>95% cure rates in 8-12 weeks)
- WHO 2030 elimination targets (90% diagnosed, 80% treated)
- High-risk populations (PWID, incarcerated, healthcare workers)
- Treatment simplification with pangenotypic regimens
- Post-SVR surveillance requirements for cirrhosis patients

### Patient Explanation (1,338 characters)
Patient-friendly Portuguese explanation covering:
- What anti-HCV test detects (antibodies vs active infection)
- Why confirmation with HCV RNA is needed
- Transmission routes (blood contact, PWID, procedures)
- Modern cure rates with DAAs (>95% cure in 8-12 weeks)
- Natural history if untreated (cirrhosis, HCC, liver failure)
- Importance of early detection and treatment

### Conduct (2,466 characters)
Detailed clinical protocols including:
- Anti-HCV interpretation (negative, positive, window period)
- Mandatory HCV RNA confirmation for positive results
- Comprehensive workup (liver function, coinfections, fibrosis staging)
- Treatment initiation criteria and priority populations
- Pangenotypic DAA regimens (sofosbuvir/velpatasvir, glecaprevir/pibrentasvir)
- Drug-drug interaction assessment
- SVR12 definition and post-cure surveillance
- Management of resolved infections and reinfection risk

---

## Execution Instructions

### Step 1: Execute SQL enrichment
```bash
docker compose exec -T db psql -U plenya_user -d plenya_db < /home/user/plenya/enrich_anti_hcv.sql
```

### Step 2: Verify results
The SQL file includes an automatic verification query that will display:
- Score item ID and name
- Character counts for each clinical field
- Last review timestamp
- Number of linked articles (should be 4)

---

## Expected Verification Output

```
                  id                  |         name         | clinical_relevance_chars | patient_explanation_chars | conduct_chars | last_review | linked_articles_count
--------------------------------------+---------------------+--------------------------+---------------------------+---------------+-------------+----------------------
 73411e66-c180-46ad-8f5b-7d67b26ef877 | Hepatite C - Anti-HCV|         1958            |           1338           |     2466      | 2026-01-29  |          4
```

---

## Quality Assurance Checklist

- [x] 4 peer-reviewed articles from 2024-2025
- [x] Mix of article types (review, systematic review, clinical guideline)
- [x] All articles have DOI and URLs
- [x] Used `publish_date` (date) not `year`
- [x] Clinical relevance: 1,500-2,000 characters
- [x] Patient explanation: 1,000-1,500 characters (PT-BR)
- [x] Conduct: 1,500-2,500 characters
- [x] All content evidence-based from cited articles
- [x] DAA cure rates >95% prominently featured
- [x] HCV RNA confirmation protocols included
- [x] SVR12 definition and monitoring included
- [x] Post-cure surveillance for cirrhosis patients included

---

## Sources Referenced

- [A 2024 global report on national policies and hepatitis C elimination](https://pmc.ncbi.nlm.nih.gov/articles/PMC12308979/)
- [Hepatitis C Virus: Epidemiological Challenges and Global Strategies](https://pmc.ncbi.nlm.nih.gov/articles/PMC12390683/)
- [SASLT guidelines: Update in HCV treatment 2024](https://pmc.ncbi.nlm.nih.gov/articles/PMC10856511/)
- [Cure of chronic HCV after DAA treatment - systematic review](https://www.frontiersin.org/journals/immunology/articles/10.3389/fimmu.2025.1546915/full)
- [WHO Hepatitis C Fact Sheet](https://www.who.int/news-room/fact-sheets/detail/hepatitis-c)
- [CDC Clinical Screening and Diagnosis for Hepatitis C](https://www.cdc.gov/hepatitis-c/hcp/diagnosis-testing/)
- [AASLD/IDSA HCV Guidance 2023](https://www.hcvguidelines.org/)

---

**File Location:** `/home/user/plenya/enrich_anti_hcv.sql`
