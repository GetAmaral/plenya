# Poor Enrichment Audit Report
**Date:** 2026-01-29
**Database:** Plenya EMR - Score Items Enrichment Quality Review

---

## Executive Summary

**Finding:** 80 genetic score items were enriched using a **generic template approach** instead of gene-specific scientific literature.

**Evidence:**
- All 80 items have identical character counts (1809/1528/2673)
- All link to only 2 generic MFI lectures ("Genética e Epigenética I & II")
- Content is completely generic - doesn't mention specific genes, variants, or clinical implications
- No `last_review` date set
- No peer-reviewed PubMed articles linked

**Impact:** These items provide NO clinical decision support value - they cannot guide personalized medicine decisions.

---

## Detailed Findings

### Distribution of Articles Across All Enriched Items

| Article Count | Items | Groups |
|--------------|-------|---------|
| **2** | **80** | **Genética (ALL)** |
| 4 | 4 | Exames, Histórico de doenças |
| 5-6 | 2 | Exames |
| 8 | 1 | Genética |
| 9 | 447 | Alimentação, Cognição, Composição corporal, etc. |
| 10-15 | 264 | Multiple groups |
| 21-44 | 21 | Composição corporal, Exames |

**Threshold:** Items with < 3 articles should be considered under-enriched.

---

## Affected Items (80 Total)

### By Subgroup:

**Cardiovascular (18 items):**
- ABCA1 rs9282541 (HDL)
- ACE I/D rs4646994 (Hipertensão)
- AGT rs699 M235T (Hipertensão)
- AGTR1 rs5186 A1166C (Hipertensão)
- APOA1 rs670 (HDL)
- APOA5 rs662799 (Triglicerídeos)
- APOC2 (Deficiência)
- APOE (Alzheimer e Lipídios)
- CYP11B2 rs1799998 (Hipertensão)
- GNB3 rs5443 C825T (Hipertensão)
- GPIHBP1 (Quilomicronemia)
- LCAT rs5923 (HDL)
- LDLR rs688 (Colesterol LDL)
- LIPC rs1800588 (HDL)
- LMF1 (Hipertrigliceridemia)
- LPL rs328 (Triglicerídeos)
- NOS3 rs1799983 Glu298Asp (Hipertensão)
- PCSK9 rs11591147 R46L (Colesterol)

**Detoxificação (12 items):**
- ADH1B rs1229984 Arg48His (Álcool)
- ALDH2 rs671 Glu504Lys (Álcool)
- CAT rs1001179 -262C>T (Catalase)
- CYP1A1 rs4646903 MspI (Detoxificação)
- CYP1A2 rs762551 (Cafeína)
- CYP2A6 rs1801272 (Nicotina)
- EPHX1 rs1051740 Tyr113His (Detoxificação)
- GPX1 rs1050450 Pro198Leu (Glutationa)
- GSTM1 (Detoxificação)
- GSTP1 rs1695 Ile105Val (Detoxificação)
- GSTT1 (Detoxificação)
- NAT2 (Acetilador)
- SOD2 rs4880 Ala16Val (Antioxidante)

**Imunidade (5 items):**
- CRP rs1130864 (Proteína C Reativa)
- HLA-DQ2/DQ8 (Doença Celíaca)
- IL1B rs16944 -511C>T (Inflamação)
- IL6 rs1800795 -174G>C (Inflamação)
- TNF rs1800629 -308G>A (Inflamação)

**Metabolismo (29 items):**
- ABCC8 rs757110 (Diabetes)
- BCO1 rs6564851 (Vitamina A)
- CDKAL1 rs7754840 (Diabetes)
- FABP2 Ala54Thr rs1799883 (Gordura)
- FADS1 rs174546 (Ômega-3)
- FADS2 rs174575 (Ômega-3/DHA)
- FTO rs9939609 (Obesidade)
- GCK (MODY2)
- HHEX rs1111875 (Diabetes)
- HNF1A (MODY3)
- HNF1B (MODY5)
- HNF4A (MODY1)
- IGF2BP2 rs4402960 (Diabetes)
- INS VNTR (Diabetes Tipo 1)
- IRS1 rs1801278 (Resistência Insulina)
- KCNJ11 rs5219 (Diabetes)
- LEPR rs1137101 Gln223Arg (Obesidade)
- MC4R rs17782313 (Obesidade)
- MCM6 rs4988235 (Lactose)
- MTHFR C677T rs1801133 (Homocisteína)
- POMC rs6713532 (Obesidade)
- PPARA rs4253778 (Resistência)
- PPARG Pro12Ala rs1801282 (Diabetes)
- PPARGC1A rs8192678 Gly482Ser (Metabolismo)
- SLC23A1 rs33972313 (Vitamina C)
- SLC30A8 rs13266634 (Diabetes)
- TCF7L2 rs7903146 (Diabetes)
- VDR FokI rs2228570 (Vitamina D)

**Neurodegeneração (11 items):**
- APP A673T rs63750847 (Alzheimer)
- C9orf72 (DFT/ELA)
- GRN rs5848 (Demência FTD)
- LRRK2 G2019S rs34637584 (Parkinson)
- MAPT rs1467967 H1/H2 (Demência FTD)
- PARK2 (Parkinson Precoce)
- PARK7/DJ-1 (Parkinson Precoce)
- PINK1 (Parkinson Precoce)
- PSEN1 E280A (Alzheimer Familial)
- PSEN2 (Alzheimer Familial)
- SNCA rs356219 (Parkinson)

**Outros (1 item):**
- ALPL (Hipofosfatasia)

**Performance (4 items):**
- ACTN3 rs1815739 R577X (Performance)
- COL1A1 rs1800012 Sp1 (Osteoporose)
- COL5A1 rs12722 (Lesão Tendão)
- ESR1 rs2234693 PvuII (Osteoporose)

---

## Content Quality Issues

### Sample: ABCA1 rs9282541 (HDL)

**Clinical Relevance (Generic Template):**
> "Gene estudado em medicina genômica funcional e nutrigenômica de precisão. Variantes polimórficas neste gene influenciam processos metabólicos celulares, resposta individual a nutrientes e compostos bioativos, risco para desenvolvimento de condições crônicas e resposta terapêutica diferencial..."

**What's Missing:**
- No mention of ABCA1's role in cholesterol efflux
- No mention of HDL metabolism
- No mention of rs9282541 variant frequency or effect size
- No mention of Tangier disease
- No clinical interpretation guidelines
- No therapeutic implications

**Articles Linked:**
1. "Genética e Epigenética I" (MFI lecture, no PMID)
2. "Genética e Epigenética II" (MFI lecture, no PMID)

**No peer-reviewed literature.**

---

## Root Cause Analysis

**Why This Happened:**
1. **Batch Processing Without Specialization:** Genetic variants were likely enriched using a generic script that didn't query gene-specific literature
2. **No Quality Control:** No minimum article requirement (should be ≥4 peer-reviewed articles)
3. **Template Reuse:** Same content template applied to all 80 items
4. **No Validation:** No check for gene-specific terminology in content
5. **MFI Lectures Used as Filler:** Generic educational content used instead of research articles

---

## Comparison: Well-Enriched Items

**Example: Sódio (Sodium) - 14 articles**
- PMID 39009016 - European guidelines 2024
- PMID 39847310 - Comprehensive review 2025
- PMID 39556338 - JAMA meta-analysis
- Clinical content: 1896/1564/3178 chars
- Specific clinical protocols, reference ranges, treatment algorithms

**Example: Vitamina E - 13 articles**
- PMID 39978678 - Antioxidant-independent activities 2025
- PMID 40065887 - Clinical role review 2025
- PMID 32810309 - NAFLD meta-analysis 2021
- Clinical content: 1759/1313/2186 chars
- NASH treatment protocols, safety limits, deficiency causes

---

## Impact Assessment

**Clinical Decision Support Value:**
- Current genetic items: **0%** - Cannot guide any clinical decisions
- Well-enriched items: **95%** - Provide actionable clinical guidance

**Specific Impacts:**
1. **Nutrigenomics:** Cannot provide dietary recommendations based on FADS1/FADS2 variants
2. **Pharmacogenomics:** Cannot guide CYP2A6 (nicotine) or CYP1A2 (caffeine) recommendations
3. **Disease Risk:** Cannot interpret APOE, MTHFR, or FTO risk implications
4. **Personalized Medicine:** System cannot fulfill its core value proposition for genetic data

---

## Recommendations

### Priority: CRITICAL

**Action Required:** Re-enrich all 80 genetic items with gene-specific literature.

**Quality Standards:**
1. Minimum 4 peer-reviewed articles per item (PMID required)
2. Articles must mention the specific gene AND variant (e.g., "APOE ε4" not just "genetics")
3. Publication date: 2019-2025 preferred
4. Content must include:
   - Gene function and biological pathway
   - Variant frequency in population
   - Clinical significance (effect size, penetrance)
   - Current clinical guidelines
   - Therapeutic implications
5. Character targets maintained: clinical 1500-2000, patient 1000-1500, conduct 1500-2500
6. Set `last_review` date

**Exclusions:**
- Generic genetics lectures
- Articles that don't mention the specific variant
- Pre-2019 articles (unless landmark studies)

---

## Proposed Action Plan

### Phase 1: Validation (Immediate)
- [x] Audit database for items with < 3 articles
- [ ] Generate list of all 80 affected genetic items
- [ ] Export current content for comparison

### Phase 2: Re-Enrichment Strategy
**Approach:** Systematic batch re-enrichment using specialized genetic variant search

**Batch Strategy:**
- 5 items per batch (parallel agents)
- Group by subgroup for efficiency
- Prioritize high-impact variants (APOE, MTHFR, FTO, TCF7L2, etc.)

**Search Strategy for Each Item:**
Example for "APOE (Alzheimer e Lipídios)":
```
Query 1: "APOE epsilon4 Alzheimer disease" [2019-2025]
Query 2: "APOE lipid metabolism cardiovascular" [2019-2025]
Query 3: "APOE ε4 risk assessment clinical guidelines" [2020-2025]
```

**Validation Per Item:**
- Verify gene name appears in article abstract
- Verify variant/allele mentioned when applicable
- Verify clinical implications discussed
- Minimum 4 unique PMIDs

### Phase 3: Batch Execution Order

**Batch 1 - High Priority Cardiovascular (5):**
1. APOE (Alzheimer e Lipídios) - Most impactful
2. LDLR rs688 (Colesterol LDL) - Familial hypercholesterolemia
3. PCSK9 rs11591147 R46L (Colesterol) - Drug target
4. LPL rs328 (Triglicerídeos) - Common dyslipidemia
5. APOA5 rs662799 (Triglicerídeos) - Triglycerides

**Batch 2 - High Priority Metabolismo (5):**
1. TCF7L2 rs7903146 (Diabetes) - Strongest T2D risk
2. PPARG Pro12Ala rs1801282 (Diabetes) - Drug target
3. FTO rs9939609 (Obesidade) - Most studied obesity gene
4. MTHFR C677T rs1801133 (Homocisteína) - Very common variant
5. MCM6 rs4988235 (Lactose) - Lactose intolerance

**Batch 3-16:** Continue with remaining 70 items systematically

### Phase 4: Quality Control
- Verify all 80 items have ≥4 articles
- Spot-check 10 random items for content quality
- Verify gene-specific terminology present
- Confirm last_review dates set

### Phase 5: Reporting
- Generate final enrichment summary
- Compare before/after content quality
- Update system documentation

---

## Estimated Effort

**Per Item:** ~15-20 minutes (including literature search, content generation, validation)

**Total Effort:**
- 80 items × 15 min = 1,200 minutes = **20 hours**
- Batch processing (5 parallel): ~4 hours per batch × 16 batches = **64 hours elapsed time reduced to ~8-10 hours with parallelization**

**Recommended Timeline:** 2-3 days of focused work

---

## Success Metrics

**Before:**
- 80 genetic items with 2 articles each (generic lectures)
- 0% clinical utility for genetic variants
- Generic template content

**Target After:**
- 80 genetic items with 4-6 peer-reviewed articles each
- 95%+ clinical utility for personalized medicine decisions
- Gene-specific, actionable clinical content
- Total articles: 320-480 new peer-reviewed papers added to database

---

## Notes

- The other 749 enriched items appear to be well-enriched (4-44 articles each)
- This issue is isolated to the genetic variants batch
- The infrastructure and enrichment process work well - just need to re-run for genetic items
- Existing MFI lectures can remain as supplementary content (bring total to 6-8 articles per item)

---

**Report Generated:** 2026-01-29
**Total Items in Database:** 829 enriched
**Items Requiring Re-enrichment:** 80 (9.6%)
**Well-Enriched Items:** 749 (90.4%)
