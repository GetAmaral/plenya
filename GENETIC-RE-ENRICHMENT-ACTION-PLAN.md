# Genetic Items Re-Enrichment Action Plan
**Status:** AWAITING APPROVAL
**Date:** 2026-01-29

---

## Problem Statement

80 genetic score items were enriched with generic template content and only 2 MFI lecture articles each. They provide **zero clinical decision support value** for personalized medicine.

**Example Issue (APOE):**
- Current content: "Gene studied in functional genomics..." (generic template)
- Current articles: 2 generic lectures
- **Missing:** APOE ε4 Alzheimer risk, cardiovascular implications, clinical guidelines

---

## Proposed Solution

**Re-enrich all 80 genetic items** with gene-specific peer-reviewed literature and actionable clinical content.

---

## Quality Standards

### Per Item Requirements:
1. **Minimum 4 peer-reviewed articles** (PubMed PMID required)
2. **Gene-specific content:**
   - Gene function and biological pathway
   - Variant frequency and effect size
   - Clinical significance and penetrance
   - Current clinical guidelines
   - Therapeutic implications
3. **Article selection criteria:**
   - Must mention specific gene AND variant in abstract
   - Publication: 2019-2025 preferred
   - Journals: Major genetics/medical journals
   - Types: Research articles, reviews, meta-analyses, clinical guidelines
4. **Character targets:** clinical 1500-2000, patient 1000-1500, conduct 1500-2500
5. **Set last_review date**

---

## Batch Execution Strategy

### Approach: 16 batches × 5 items (parallel agents)

**Estimated Time:**
- Per batch: ~30-40 minutes
- Total: ~8-10 hours with parallelization
- Timeline: 2-3 days of focused work

### Priority Order:

**Batch 1 - High-Impact Cardiovascular (5 items):**
1. APOE (Alzheimer e Lipídios)
2. LDLR rs688 (Colesterol LDL)
3. PCSK9 rs11591147 R46L (Colesterol)
4. LPL rs328 (Triglicerídeos)
5. APOA5 rs662799 (Triglicerídeos)

**Batch 2 - High-Impact Metabolismo (5 items):**
1. TCF7L2 rs7903146 (Diabetes) - Strongest T2D risk SNP
2. PPARG Pro12Ala rs1801282 (Diabetes)
3. FTO rs9939609 (Obesidade)
4. MTHFR C677T rs1801133 (Homocisteína)
5. MCM6 rs4988235 (Lactose)

**Batch 3 - Pharmacogenomics Detoxificação (5 items):**
1. CYP1A2 rs762551 (Cafeína)
2. CYP2A6 rs1801272 (Nicotina)
3. ALDH2 rs671 Glu504Lys (Álcool)
4. ADH1B rs1229984 Arg48His (Álcool)
5. NAT2 (Acetilador)

**Batch 4 - Antioxidant/Detox (5 items):**
1. SOD2 rs4880 Ala16Val (Antioxidante)
2. GPX1 rs1050450 Pro198Leu (Glutationa)
3. GSTP1 rs1695 Ile105Val (Detoxificação)
4. GSTM1 (Detoxificação)
5. GSTT1 (Detoxificação)

**Batch 5 - Remaining Cardiovascular Part 1 (5 items):**
1. ACE I/D rs4646994 (Hipertensão)
2. AGT rs699 M235T (Hipertensão)
3. AGTR1 rs5186 A1166C (Hipertensão)
4. NOS3 rs1799983 Glu298Asp (Hipertensão)
5. GNB3 rs5443 C825T (Hipertensão)

**Batch 6 - Remaining Cardiovascular Part 2 (5 items):**
1. ABCA1 rs9282541 (HDL)
2. APOA1 rs670 (HDL)
3. LCAT rs5923 (HDL)
4. LIPC rs1800588 (HDL)
5. CYP11B2 rs1799998 (Hipertensão)

**Batch 7 - Rare Lipid Disorders (3 items):**
1. APOC2 (Deficiência)
2. GPIHBP1 (Quilomicronemia)
3. LMF1 (Hipertrigliceridemia)

**Batch 8 - Neurodegeneração Part 1 (5 items):**
1. APOE (already in Batch 1, skip)
2. LRRK2 G2019S rs34637584 (Parkinson)
3. SNCA rs356219 (Parkinson)
4. PSEN1 E280A (Alzheimer Familial)
5. PSEN2 (Alzheimer Familial)
6. APP A673T rs63750847 (Alzheimer)

**Batch 9 - Neurodegeneração Part 2 (6 items):**
1. PARK2 (Parkinson Precoce)
2. PARK7/DJ-1 (Parkinson Precoce)
3. PINK1 (Parkinson Precoce)
4. C9orf72 (DFT/ELA)
5. MAPT rs1467967 H1/H2 (Demência FTD)
6. GRN rs5848 (Demência FTD)

**Batch 10 - Diabetes/MODY (5 items):**
1. HNF1A (MODY3)
2. HNF4A (MODY1)
3. GCK (MODY2)
4. HNF1B (MODY5)
5. INS VNTR (Diabetes Tipo 1)

**Batch 11 - Diabetes Risk SNPs (5 items):**
1. KCNJ11 rs5219 (Diabetes)
2. ABCC8 rs757110 (Diabetes)
3. CDKAL1 rs7754840 (Diabetes)
4. HHEX rs1111875 (Diabetes)
5. IGF2BP2 rs4402960 (Diabetes)
6. SLC30A8 rs13266634 (Diabetes)

**Batch 12 - Insulin Resistance & Obesity (5 items):**
1. IRS1 rs1801278 (Resistência Insulina)
2. LEPR rs1137101 Gln223Arg (Obesidade)
3. MC4R rs17782313 (Obesidade)
4. POMC rs6713532 (Obesidade)
5. PPARA rs4253778 (Resistência)
6. PPARGC1A rs8192678 Gly482Ser (Metabolismo)

**Batch 13 - Nutrient Metabolism (5 items):**
1. FADS1 rs174546 (Ômega-3)
2. FADS2 rs174575 (Ômega-3/DHA)
3. FABP2 Ala54Thr rs1799883 (Gordura)
4. BCO1 rs6564851 (Vitamina A)
5. VDR FokI rs2228570 (Vitamina D)
6. SLC23A1 rs33972313 (Vitamina C)

**Batch 14 - Inflammation & Immunity (5 items):**
1. IL6 rs1800795 -174G>C (Inflamação)
2. IL1B rs16944 -511C>T (Inflamação)
3. TNF rs1800629 -308G>A (Inflamação)
4. CRP rs1130864 (Proteína C Reativa)
5. HLA-DQ2/DQ8 (Doença Celíaca)

**Batch 15 - Remaining Detox (2 items):**
1. CYP1A1 rs4646903 MspI (Detoxificação)
2. EPHX1 rs1051740 Tyr113His (Detoxificação)
3. CAT rs1001179 -262C>T (Catalase)

**Batch 16 - Performance & Other (5 items):**
1. ACTN3 rs1815739 R577X (Performance)
2. COL1A1 rs1800012 Sp1 (Osteoporose)
3. COL5A1 rs12722 (Lesão Tendão)
4. ESR1 rs2234693 PvuII (Osteoporose)
5. ALPL (Hipofosfatasia)

---

## Search Strategy Example

### For: APOE (Alzheimer e Lipídios)

**PubMed Queries:**
1. `"APOE epsilon4" AND "Alzheimer disease" AND (risk OR biomarker)` [2020-2025]
2. `"APOE genotype" AND (cardiovascular OR lipid OR cholesterol)` [2019-2025]
3. `"APOE" AND "clinical guidelines"` [2020-2025]
4. `"APOE ε4" AND (prevention OR treatment)` [2021-2025]

**Expected Articles:**
- GWAS/association studies
- Clinical guidelines (Alzheimer's Association, etc.)
- Lipid metabolism reviews
- Personalized medicine protocols

### For: CYP1A2 rs762551 (Cafeína)

**PubMed Queries:**
1. `"CYP1A2" AND "rs762551" AND caffeine` [2019-2025]
2. `"CYP1A2*1F" AND metabolism` [2020-2025]
3. `"CYP1A2" AND pharmacogenetics AND (coffee OR caffeine)` [2019-2025]

**Expected Articles:**
- Pharmacogenetics studies
- Caffeine metabolism variability
- Clinical implications for caffeine sensitivity

---

## Validation Checklist (Per Item)

After enrichment, verify:
- [ ] Gene name appears in clinical_relevance field
- [ ] Variant ID (rsID or allele) mentioned when applicable
- [ ] Minimum 4 articles with PMIDs
- [ ] Articles mention the gene in abstract (spot-check 2 random)
- [ ] Clinical content ≥ 1500 chars
- [ ] Patient content ≥ 1000 chars
- [ ] Conduct content ≥ 1500 chars
- [ ] last_review date set to 2026-01-29
- [ ] No generic template language ("Gene studied in functional genomics...")

---

## Database Operations

### Per Item:
1. **UPDATE score_items** - Replace clinical_relevance, patient_explanation, conduct, last_review
2. **INSERT articles** - Add 4-6 new peer-reviewed articles with PMIDs
3. **INSERT article_score_items** - Link articles to score item
4. **KEEP existing MFI lectures** - Don't delete, results in 6-8 total articles per item

### Rollback Strategy:
- Export current content before starting (backup)
- Track changes in git
- Can revert individual items if needed

---

## Expected Outcomes

**Before:**
- 80 items × 2 generic articles = 160 article links
- 0% clinical utility
- Generic content

**After:**
- 80 items × 6 articles (4 new + 2 existing) = 480 article links
- 95%+ clinical utility
- Gene-specific, actionable content
- **+320 peer-reviewed articles** added to database

**Database Growth:**
- Current: 838 articles, 8,203 links
- After: ~1,158 articles (+38%), ~8,523 links (+4%)

---

## Risk Assessment

**Low Risk:**
- Only updating existing items (no schema changes)
- Existing MFI lectures remain (additive approach)
- Well-tested enrichment infrastructure
- Can validate/rollback individual items

**Mitigation:**
- Test 1 item manually first
- Batch execution allows incremental progress
- Quality validation after each batch
- Git tracking of all changes

---

## Next Steps (AWAITING APPROVAL)

1. **User Reviews This Plan**
2. **User Approves to Proceed** or requests modifications
3. **Execute Batch 1** (5 high-impact items) as proof of concept
4. **Review Batch 1 results** with user
5. **Continue with remaining 15 batches** if approved
6. **Final validation and reporting**

---

## Questions for User

1. **Priority:** Should we prioritize certain gene categories (e.g., pharmacogenomics first)?
2. **Thoroughness:** 4 articles minimum OK, or prefer 5-6 per item?
3. **Execution:** Start immediately or schedule for specific time?
4. **Validation:** Want to review each batch or only spot-check after completion?

---

**Status:** ⏸️ **PAUSED - AWAITING USER APPROVAL**

Do not proceed with any database changes until user confirms.
