# Risk Stratification Tables: B Vitamins (Whole Blood HPLC)

**Document Version:** 1.0
**Last Updated:** January 2026
**Methodology:** Whole Blood HPLC (High-Performance Liquid Chromatography)
**Context:** Functional Medicine Optimal Ranges for EMR Integration

---

## Table of Contents

1. [Introduction](#introduction)
2. [Vitamin B1 (Thiamine) - Whole Blood](#vitamin-b1-thiamine---whole-blood)
3. [Vitamin B2 (Riboflavin) - Whole Blood](#vitamin-b2-riboflavin---whole-blood)
4. [Vitamin B6 (Pyridoxal-5-Phosphate) - Whole Blood](#vitamin-b6-pyridoxal-5-phosphate---whole-blood)
5. [Clinical Integration Notes](#clinical-integration-notes)
6. [References](#references)

---

## Introduction

This document provides evidence-based risk stratification tables for three critical B vitamins measured by **whole blood HPLC**, the gold standard for assessing intracellular vitamin status. Unlike serum/plasma measurements that reflect recent intake, whole blood values correlate better with tissue stores and functional status.

**Why Whole Blood HPLC?**
- Reflects intracellular status (>90% of vitamins are in RBCs)
- Better correlation with clinical outcomes
- Less influenced by recent dietary intake
- Gold standard for functional deficiency assessment

**Functional Medicine Approach:**
- Ranges are tighter than conventional reference intervals
- Focus on optimal function (mitochondrial, neurological, methylation)
- Earlier intervention for functional deficiencies
- Upper safe limits to prevent toxicity

---

## Vitamin B1 (Thiamine) - Whole Blood

### Overview

**Biomarker:** Thiamine Diphosphate (ThDP) - active form
**Method:** Whole Blood HPLC with fluorescence detection
**Units:** nmol/L (primary), µg/L (alternative)
**Molecular Weight:** 425.3 g/mol (ThDP)

### Conversion Factors

- **nmol/L to µg/L:** multiply by 0.425
- **µg/L to nmol/L:** multiply by 2.353
- **nmol/L to ng/mL:** multiply by 0.425

### Clinical Significance

Thiamine is essential for:
- **Carbohydrate metabolism:** Cofactor for pyruvate dehydrogenase, α-ketoglutarate dehydrogenase
- **Energy production:** Critical for mitochondrial ATP synthesis
- **Neurological function:** Myelin synthesis, neurotransmitter production
- **Cardiovascular health:** Cardiac muscle function, vascular tone

### Risk Stratification Table

| Level | Range (nmol/L) | Status | Clinical Implications | Interventions |
|-------|----------------|--------|----------------------|---------------|
| **0** | < 40 | Severe Deficiency | **CRITICAL:** High risk of Wernicke-Korsakoff syndrome, beriberi (wet/dry), cardiac failure, peripheral neuropathy, confusion, ataxia, ophthalmoplegia | Emergency IV thiamine 200-500mg TID × 3 days, then 100-250mg daily, hospitalization may be required |
| **1** | 40-69 | Moderate Deficiency | **HIGH RISK:** Subclinical beriberi, fatigue, poor glucose metabolism, cognitive impairment, peripheral neuropathy symptoms, tachycardia, dyspnea on exertion | High-dose oral thiamine 100-300mg daily (benfotiamine preferred), investigate causes (alcoholism, malabsorption, diuretics) |
| **2** | 70-84 | Mild Deficiency | **MODERATE RISK:** Functional insufficiency, impaired energy metabolism, early cognitive symptoms, reduced exercise tolerance, mild neuropathy | Thiamine 50-150mg daily, optimize diet (whole grains, legumes, pork), address absorption issues |
| **3** | 85-120 | Suboptimal | **LOW RISK:** Below optimal for mitochondrial function, may compromise glucose metabolism in high-demand states (athletes, pregnancy, diabetes) | Thiamine 25-100mg daily or dietary optimization, consider B-complex, reassess in 3 months |
| **4** | 121-180 | Optimal | **IDEAL RANGE:** Adequate for all metabolic functions, optimal mitochondrial health, cardiovascular protection, cognitive function | Maintenance: 10-50mg daily or comprehensive B-complex, continue healthy diet |
| **5** | 181-250 | High Optimal | **VERY GOOD:** Superior reserves, excellent for high-demand states (athletes, pregnancy, chronic illness recovery) | Continue current regimen, no intervention needed, retest annually |
| **6** | > 250 | Supraphysiological | **MONITOR:** Very high levels, typically from recent high-dose supplementation. No toxicity established for thiamine (water-soluble with excellent safety profile) | Review supplementation (may be unnecessary), reduce to maintenance dose, no safety concerns |

### CSV Format for Database Import

```csv
Vitamina B1 (Tiamina Difosfato) - Sangue Total;nmol/L | µg/L: ×0.425;20;0. Deficiência Grave (< 40): CRÍTICO - Risco de Wernicke-Korsakoff e beribéri;1. Deficiência Moderada (40-69): ALTO RISCO - Beribéri subclínico;2. Deficiência Leve (70-84): RISCO MODERADO - Insuficiência funcional;3. Subótimo (85-120): BAIXO RISCO - Abaixo do ótimo;4. Ótimo (121-180): IDEAL - Função metabólica adequada;5. Ótimo Alto (181-250): EXCELENTE - Reservas superiores;6. Suprafisiológico (> 250): MONITORAR - Níveis muito elevados
```

### Key Clinical Contexts

#### Beriberi (Thiamine Deficiency Disease)
- **Wet Beriberi:** Cardiovascular symptoms - high-output heart failure, tachycardia, edema, dyspnea, warm extremities
- **Dry Beriberi:** Neurological symptoms - peripheral neuropathy, muscle weakness, paresthesias, decreased reflexes
- **Onset:** Can occur within 3-4 weeks of severe dietary restriction
- **Reversibility:** Cardiac symptoms improve within 24-48 hours of treatment; neurological recovery may take weeks-months

#### Wernicke-Korsakoff Syndrome
- **Wernicke Encephalopathy:** Acute - confusion, ataxia, ophthalmoplegia (classic triad present in only 10-30%)
- **Korsakoff Syndrome:** Chronic - anterograde amnesia, confabulation, apathy
- **Risk Factors:** Alcoholism (most common), hyperemesis gravidarum, malnutrition, gastric bypass surgery
- **Prevention:** Thiamine 100-250mg before glucose administration in at-risk patients

#### Other High-Risk Groups
- **Diabetics:** Increased urinary thiamine loss (40-80% lower levels)
- **Heart Failure:** Both cause and consequence of deficiency
- **Chronic Diuretic Use:** Furosemide increases thiamine excretion
- **Post-Bariatric Surgery:** Malabsorption risk
- **Chronic Alcohol Use:** Impaired absorption, increased requirements

---

## Vitamin B2 (Riboflavin) - Whole Blood

### Overview

**Biomarker:** Riboflavin + FAD (Flavin Adenine Dinucleotide)
**Method:** Whole Blood HPLC with fluorescence detection
**Units:** nmol/L (primary), µg/L (alternative)
**Molecular Weight:** 376.4 g/mol (riboflavin), 785.6 g/mol (FAD)

### Conversion Factors

- **nmol/L to µg/L (riboflavin):** multiply by 0.376
- **µg/L to nmol/L:** multiply by 2.66
- **nmol/L to ng/mL:** multiply by 0.376

### Clinical Significance

Riboflavin is essential for:
- **Methylation:** FAD cofactor for MTHFR enzyme (folate cycle)
- **Antioxidant Defense:** Glutathione reductase cofactor
- **Energy Production:** Complex I and II in electron transport chain
- **MTHFR Support:** Critical for C677T variant carriers (30-40% reduced enzyme activity)
- **Migraine Prevention:** High-dose therapy (400mg/day) reduces frequency and severity

### Risk Stratification Table

| Level | Range (nmol/L) | Status | Clinical Implications | Interventions |
|-------|----------------|--------|----------------------|---------------|
| **0** | < 80 | Severe Deficiency | **CRITICAL:** Clinical riboflavin deficiency (ariboflavinosis), cheilosis, angular stomatitis, glossitis, seborrheic dermatitis, anemia, neuropathy, cataracts | High-dose riboflavin 50-100mg daily, investigate causes (malabsorption, vegan diet), assess for other B vitamin deficiencies |
| **1** | 80-120 | Moderate Deficiency | **HIGH RISK:** Impaired MTHFR function (especially C677T carriers), elevated homocysteine, reduced glutathione recycling, increased oxidative stress, low energy | Riboflavin 25-100mg daily, check MTHFR genotype, measure homocysteine, comprehensive B-complex |
| **2** | 121-160 | Mild Deficiency | **MODERATE RISK:** Suboptimal methylation, increased migraine risk, compromised antioxidant capacity, mild fatigue | Riboflavin 10-50mg daily, optimize diet (dairy, eggs, leafy greens), consider B-complex |
| **3** | 161-220 | Suboptimal | **LOW RISK:** Adequate for basic functions but below optimal for MTHFR support, migraine prevention, or high oxidative stress states | Riboflavin 10-25mg daily, dietary optimization, reassess in 3 months |
| **4** | 221-350 | Optimal | **IDEAL RANGE:** Full MTHFR activity, optimal methylation, adequate antioxidant support, excellent for most individuals | Maintenance: 5-10mg daily or B-complex, continue healthy diet |
| **5** | 351-500 | High Optimal | **VERY GOOD:** Superior methylation support, excellent for MTHFR C677T carriers, migraine prevention range (with 400mg/day dosing) | Continue current regimen if therapeutic (migraines), otherwise reduce to maintenance |
| **6** | > 500 | Supraphysiological | **MONITOR:** Very high levels from supplementation. Riboflavin has excellent safety profile (no established toxicity), but levels >500 nmol/L suggest recent high-dose intake | Review necessity of high-dose therapy, bright yellow urine is normal, reduce if not therapeutically indicated |

### CSV Format for Database Import

```csv
Vitamina B2 (Riboflavina) - Sangue Total;nmol/L | µg/L: ×0.376;20;0. Deficiência Grave (< 80): CRÍTICO - Ariboflavinose clínica;1. Deficiência Moderada (80-120): ALTO RISCO - MTHFR comprometida;2. Deficiência Leve (121-160): RISCO MODERADO - Metilação subótima;3. Subótimo (161-220): BAIXO RISCO - Abaixo do ótimo para MTHFR;4. Ótimo (221-350): IDEAL - Atividade plena de MTHFR;5. Ótimo Alto (351-500): EXCELENTE - Suporte superior de metilação;6. Suprafisiológico (> 500): MONITORAR - Níveis muito elevados
```

### Key Clinical Contexts

#### MTHFR Gene Variants
- **C677T Polymorphism:** 30-40% reduction in enzyme activity (homozygous TT: ~40-50% of population depending on ethnicity)
- **Riboflavin Response:** FAD cofactor stabilizes mutant MTHFR enzyme
- **Clinical Impact:** Elevated homocysteine, increased cardiovascular risk, pregnancy complications (neural tube defects), depression
- **Therapeutic Approach:** Riboflavin 5-50mg daily significantly improves MTHFR function in TT carriers
- **Evidence:** Riboflavin supplementation normalizes homocysteine in 40-60% of C677T carriers

#### Migraine Prevention
- **Mechanism:** Improves mitochondrial energy metabolism, reduces oxidative stress in neurons
- **Dosing:** 400mg daily for 3-4 months minimum
- **Efficacy:** 50-60% reduction in migraine frequency and severity in responders
- **Time to Effect:** 2-3 months of consistent supplementation
- **Evidence:** Multiple RCTs show efficacy comparable to prescription prophylactics
- **Safety:** Excellent - no significant adverse effects at 400mg daily

#### Other Clinical Applications
- **Homocysteine Management:** Essential component of methylation pathway
- **Oxidative Stress:** Glutathione reductase requires FAD
- **Anemia Support:** Required for iron metabolism and RBC production
- **Ocular Health:** Prevents cataracts, supports corneal health

---

## Vitamin B6 (Pyridoxal-5-Phosphate) - Whole Blood

### Overview

**Biomarker:** Pyridoxal-5'-Phosphate (PLP) - active form
**Method:** Whole Blood HPLC with fluorescence detection or LC-MS/MS
**Units:** nmol/L (primary), ng/mL or µg/L (alternative)
**Molecular Weight:** 247.1 g/mol (PLP)

### Conversion Factors

- **nmol/L to ng/mL:** multiply by 0.247
- **ng/mL to nmol/L:** multiply by 4.05
- **nmol/L to µg/L:** multiply by 0.247

### Clinical Significance

Pyridoxal-5-phosphate is essential for:
- **Homocysteine Metabolism:** Cofactor for cystathionine β-synthase (transsulfuration pathway)
- **Neurotransmitter Synthesis:** GABA, serotonin, dopamine, norepinephrine production
- **Amino Acid Metabolism:** Over 100 PLP-dependent enzymes
- **Immune Function:** Lymphocyte proliferation, antibody production
- **Hemoglobin Synthesis:** Required for heme production

### Risk Stratification Table

| Level | Range (nmol/L) | Status | Clinical Implications | Interventions |
|-------|----------------|--------|----------------------|---------------|
| **0** | < 20 | Severe Deficiency | **CRITICAL:** Clinical deficiency - peripheral neuropathy, seizures (especially infants), microcytic anemia, seborrheic dermatitis, glossitis, depression, confusion, weakened immunity | High-dose PLP or P5P 50-100mg daily, investigate causes (medications, malabsorption), comprehensive B-complex |
| **1** | 20-40 | Moderate Deficiency | **HIGH RISK:** Elevated homocysteine, increased cardiovascular risk, impaired neurotransmitter synthesis, depression/anxiety, cognitive impairment, carpal tunnel syndrome | PLP/P5P 25-75mg daily, measure homocysteine, check for medication interactions (PPIs, oral contraceptives) |
| **2** | 41-60 | Mild Deficiency | **MODERATE RISK:** Suboptimal methylation, mood dysregulation, increased PMS symptoms, mild neuropathy risk, impaired immune function | PLP/P5P 10-50mg daily, optimize diet (poultry, fish, chickpeas), B-complex support |
| **3** | 61-90 | Suboptimal | **LOW RISK:** Adequate for basic functions but below optimal for homocysteine metabolism, neurotransmitter production in high-demand states | PLP/P5P 10-25mg daily, dietary optimization, reassess in 3 months |
| **4** | 91-180 | Optimal | **IDEAL RANGE:** Full enzymatic activity, optimal homocysteine metabolism, adequate neurotransmitter synthesis, excellent immune function | Maintenance: 5-15mg daily or B-complex, continue healthy diet |
| **5** | 181-300 | High Optimal | **GOOD:** Therapeutic range for some conditions (PMS, carpal tunnel, morning sickness), no toxicity concerns at these levels | Continue if therapeutic indication, otherwise reduce to maintenance dose |
| **6** | 301-500 | Upper Safe Range | **CAUTION:** Approaching upper safety threshold. Monitor for neuropathy symptoms (tingling, numbness, burning in extremities). Risk increases with chronic intake >200mg/day | Reduce supplementation to <100mg daily, assess necessity, monitor neurological symptoms |
| **7** | > 500 | Potentially Toxic | **WARNING:** High risk of peripheral sensory neuropathy with prolonged elevation. Toxicity documented at >2000mg/day, but susceptible individuals may react to 100-300mg daily over months | STOP high-dose supplementation immediately, investigate source (supplements, medications, IV therapy), neurological evaluation if symptomatic |

### CSV Format for Database Import

```csv
Vitamina B6 (Piridoxal-5-Fosfato) - Sangue Total;nmol/L | ng/mL: ×0.247;20;0. Deficiência Grave (< 20): CRÍTICO - Neuropatia e convulsões;1. Deficiência Moderada (20-40): ALTO RISCO - Homocisteína elevada;2. Deficiência Leve (41-60): RISCO MODERADO - Metilação subótima;3. Subótimo (61-90): BAIXO RISCO - Abaixo do ótimo;4. Ótimo (91-180): IDEAL - Atividade enzimática plena;5. Ótimo Alto (181-300): BOM - Faixa terapêutica;6. Limite Superior Seguro (301-500): CAUTELA - Monitorar neuropatia;7. Potencialmente Tóxico (> 500): ALERTA - Alto risco de neuropatia
```

### Key Clinical Contexts

#### Homocysteine Metabolism
- **Pathway:** Transsulfuration - converts homocysteine to cysteine via cystathionine
- **Clinical Significance:** Elevated homocysteine increases cardiovascular disease risk, stroke, dementia
- **Synergy:** Works with B2 (MTHFR cofactor) and B12 (methionine synthase cofactor)
- **Optimal Homocysteine:** <8 µmol/L (functional medicine), <15 µmol/L (conventional)

#### Neurotransmitter Synthesis
- **GABA:** PLP-dependent decarboxylation of glutamate - calming neurotransmitter
- **Serotonin:** Tryptophan → 5-HTP → serotonin (PLP cofactor for aromatic amino acid decarboxylase)
- **Dopamine:** Tyrosine → L-DOPA → dopamine (PLP-dependent)
- **Clinical Impact:** Deficiency associated with depression, anxiety, insomnia, ADHD

#### Peripheral Neuropathy Risk (CRITICAL)
- **Mechanism:** Paradoxical toxicity - excessive pyridoxine (not PLP) may deplete active form or damage dorsal root ganglia
- **Dose-Response:** Risk increases with >200mg/day for >2 months
- **Symptoms:** Burning, tingling, numbness in hands/feet (stocking-glove distribution), ataxia, reduced proprioception
- **Susceptibility:** Individual variation - some patients develop symptoms at 50-100mg/day
- **Reversibility:** Partial to complete recovery over months after discontinuation, but may be permanent in severe cases
- **Prevention:** Use P5P (active form) at lower doses, avoid megadoses, monitor symptoms

#### Recent Regulatory Changes (2023-2025)
- **European Union (EFSA 2023):** Upper limit reduced to 12mg/day for adults
- **Australia (2023-2027):** >50mg requires pharmacy, >200mg requires prescription
- **US/Canada:** Maintains 100mg/day upper limit, but increasing safety concerns
- **Clinical Implication:** International consensus moving toward lower safe limits

#### Special Populations
- **Pregnancy:** Increased requirements (1.9mg/day RDA), high-dose (>50mg) for morning sickness should be short-term only
- **Oral Contraceptives:** Deplete B6 status (30-50% reduction)
- **PPIs/H2 Blockers:** Impair B6 absorption
- **Renal Disease:** Increased PLP catabolism, deficiency common
- **Alcoholism:** Impaired conversion to active PLP form

---

## Clinical Integration Notes

### Methylation Support: The B Vitamin Trinity

Optimal homocysteine metabolism requires synergistic action of three B vitamins:

1. **Vitamin B2 (Riboflavin/FAD):** Cofactor for MTHFR enzyme
   - Converts 5,10-methyleneTHF to 5-methylTHF (methyl donor)
   - Essential for remethylation pathway
   - Particularly critical for MTHFR C677T carriers

2. **Vitamin B6 (PLP):** Cofactor for transsulfuration pathway
   - Converts homocysteine → cystathionine → cysteine
   - Alternative pathway when remethylation saturated
   - Produces cysteine for glutathione synthesis

3. **Vitamin B12 (Methylcobalamin):** Cofactor for methionine synthase
   - Converts homocysteine → methionine
   - Uses 5-methylTHF as methyl donor
   - Also requires folate sufficiency

**Clinical Approach:**
- **Elevated Homocysteine:** Check all three B vitamins + folate
- **MTHFR C677T Carriers:** Emphasize B2 and methylfolate
- **Cardiovascular Risk:** Target homocysteine <8 µmol/L with comprehensive B support

### Mitochondrial Health: Energy Production

All three B vitamins are critical for mitochondrial function:

- **B1 (Thiamine):** Pyruvate dehydrogenase, α-ketoglutarate dehydrogenase (Krebs cycle)
- **B2 (Riboflavin/FAD):** Complex I and II (electron transport chain)
- **B6 (PLP):** Amino acid metabolism for ATP production

**Clinical Applications:**
- Chronic fatigue syndrome
- Fibromyalgia
- Post-viral syndromes
- Athletic performance optimization
- Aging and cellular senescence

### Neurological Health

Neurological manifestations vary by deficiency:

- **B1 Deficiency:** Peripheral neuropathy (dry beriberi), Wernicke-Korsakoff, confusion, ataxia
- **B2 Deficiency:** Migraine, cognitive impairment (via methylation defects)
- **B6 Deficiency:** Peripheral neuropathy, seizures (infants), depression, anxiety
- **B6 Excess:** Peripheral sensory neuropathy (paradoxical toxicity)

**Key Distinction:** Both B6 deficiency AND excess cause neuropathy - optimal range is critical.

### Whole Blood vs. Plasma/Serum

**Why Whole Blood is Superior:**

| Aspect | Whole Blood | Plasma/Serum |
|--------|-------------|--------------|
| **Reflects** | Intracellular status (tissue stores) | Recent intake |
| **Stability** | More stable | Variable with diet |
| **Clinical Correlation** | Better predicts functional deficiency | May miss subclinical deficiency |
| **Vitamin Distribution** | 90% in RBCs | 10% in plasma |
| **Recommended For** | Functional assessment | Acute deficiency screening |

**Clinical Implication:** Always specify testing method when interpreting results. Plasma reference ranges DO NOT apply to whole blood values.

### Supplementation Strategies

#### Forms and Bioavailability

| Vitamin | Standard Form | Active/Enhanced Forms | Clinical Preference |
|---------|---------------|----------------------|---------------------|
| **B1** | Thiamine HCl | Benfotiamine (fat-soluble), Sulbutiamine | Benfotiamine for neuropathy, diabetes |
| **B2** | Riboflavin | Riboflavin-5-phosphate (FMN) | Standard riboflavin adequate, high-dose for migraines |
| **B6** | Pyridoxine HCl | Pyridoxal-5-phosphate (P5P) | P5P preferred - bypasses conversion, lower neuropathy risk |

#### Dosing Principles

1. **Start Low, Go Slow:** Begin with modest doses, increase based on response
2. **Use Active Forms:** Especially for B6 (P5P) to minimize toxicity risk
3. **Monitor Clinically:** Retest whole blood levels after 3 months
4. **Comprehensive Approach:** B vitamins work synergistically - consider B-complex
5. **Address Root Cause:** Malabsorption, medications, dietary insufficiency

#### Safety Considerations

- **B1 (Thiamine):** No known toxicity - water-soluble with excellent safety profile
- **B2 (Riboflavin):** No known toxicity - bright yellow urine is normal
- **B6 (Pyridoxine/PLP):** CAUTION - toxicity risk >200mg/day chronically, use P5P at lower doses

---

## References

### Thiamine (Vitamin B1)

1. **Mayo Clinic Laboratories.** Thiamine (Vitamin B1), Whole Blood Test Catalog. Available at: https://www.mayocliniclabs.com/test-catalog/overview/42356 (Accessed January 2026)

2. **Labcorp.** Vitamin B1, Whole Blood Test 121186. Available at: https://www.labcorp.com/tests/121186/vitamin-b1-whole-blood (Accessed January 2026)

3. **National Center for Biotechnology Information (NCBI).** Vitamin B1 (Thiamine) - StatPearls. Available at: https://www.ncbi.nlm.nih.gov/books/NBK482360/ (Updated 2024)

4. **National Center for Biotechnology Information (NCBI).** Vitamin B1 (Thiamine) Deficiency - StatPearls. Available at: https://www.ncbi.nlm.nih.gov/books/NBK537204/ (Updated 2024)

5. **NCBI PubMed Central.** Thiamine deficiency disorders: diagnosis, prevalence, and a roadmap for global control programs. Available at: https://pmc.ncbi.nlm.nih.gov/articles/PMC6392124/ (2019)

6. **NCBI PubMed Central.** Hiding in Plain Sight: Modern Thiamine Deficiency. Available at: https://pmc.ncbi.nlm.nih.gov/articles/PMC8533683/ (2021)

7. **Medscape.** Vitamin B1 (Thiamine): Reference Range, Interpretation, Collection and Panels. Available at: https://emedicine.medscape.com/article/2088582-overview (Updated 2024)

8. **Medscape.** Beriberi (Thiamine Deficiency) Workup: Approach Considerations. Available at: https://emedicine.medscape.com/article/116930-workup (Updated 2024)

9. **NCBI PubMed.** Simple HPLC Method with Internal Standard for Evaluation of Vitamin B1 Status By Use of Whole Blood. Available at: https://pubmed.ncbi.nlm.nih.gov/33636838/ (2021)

10. **HealthMatters.io.** Vitamin B1 (Thiamine) - Lab Results Explained. Available at: https://healthmatters.io/understand-blood-test-results/vitamin-b1-thiamine (Accessed January 2026)

11. **UnitsLab.com.** Vitamin B1 (Thiamine) Unit Conversion Calculator. Available at: https://unitslab.com/node/81 (Accessed January 2026)

### Riboflavin (Vitamin B2)

12. **NCBI Bookshelf.** Riboflavin and Methylenetetrahydrofolate Reductase - Madame Curie Bioscience Database. Available at: https://www.ncbi.nlm.nih.gov/books/NBK6145/ (2013)

13. **Genetic Lifehacks.** Riboflavin (Vitamin B2), MTHFR, and Deficiency Symptoms. Available at: https://www.geneticlifehacks.com/riboflavin/ (Accessed January 2026)

14. **Genetic Lifehacks.** Riboflavin (Vitamin B2): MTHFR and FMO3 SNPs. Available at: https://www.geneticlifehacks.com/riboflavin-mthfr-and-other-genetic-polymorphisms/ (Accessed January 2026)

15. **ScienceDirect/PubMed.** Riboflavin supplementation alters global and gene-specific DNA methylation in adults with the MTHFR 677 TT genotype. Available at: https://pubmed.ncbi.nlm.nih.gov/32334045/ (2020)

16. **NCBI PubMed.** Riboflavin status, MTHFR genotype and blood pressure: current evidence and implications for personalised nutrition. Available at: https://pubmed.ncbi.nlm.nih.gov/27170501/ (2016)

17. **NCBI Bookshelf.** Vitamin B2 (Riboflavin) - StatPearls. Available at: https://www.ncbi.nlm.nih.gov/books/NBK525977/ (Updated 2024)

18. **Labcorp.** Vitamin B2, Whole Blood Test 123220. Available at: https://www.labcorp.com/tests/123220/vitamin-b2-whole-blood (Accessed January 2026)

19. **NCBI PubMed Central.** Experimental and Clinical Evidence of the Effectiveness of Riboflavin on Migraines. Available at: https://pmc.ncbi.nlm.nih.gov/articles/PMC8401857/ (2021)

20. **Springer Nature.** Nutraceuticals and Headache 2024: Riboflavin, Coenzyme Q10, Feverfew, Magnesium, Melatonin, and Butterbur. Available at: https://link.springer.com/article/10.1007/s11916-025-01358-3 (2025)

21. **Neurology Journal.** Effectiveness of high‐dose riboflavin in migraine prophylaxis: A randomized controlled trial. Available at: https://www.neurology.org/doi/10.1212/WNL.50.2.466 (1998)

22. **Frontiers in Nutrition.** Practical supplements for prevention and management of migraine attacks: a narrative review. Available at: https://www.frontiersin.org/journals/nutrition/articles/10.3389/fnut.2024.1433390/full (2024)

23. **NCBI PubMed Central.** Improvement of migraine symptoms with a proprietary supplement containing riboflavin, magnesium and Q10. Available at: https://pmc.ncbi.nlm.nih.gov/articles/PMC4393401/ (2015)

24. **NeuroTest NY.** The Benefits of Riboflavin for Migraines: A Comprehensive Guide (2024). Available at: https://neurotestny.com/benefits-of-riboflavin-for-migraines-a-comprehensive-guide-2024/ (2024)

25. **Linus Pauling Institute, Oregon State University.** Riboflavin Micronutrient Information Center. Available at: https://lpi.oregonstate.edu/mic/vitamins/riboflavin (Accessed January 2026)

### Pyridoxal-5-Phosphate (Vitamin B6)

26. **NCBI Bookshelf.** Vitamin B6 Toxicity - StatPearls. Available at: https://www.ncbi.nlm.nih.gov/books/NBK554500/ (Updated 2024)

27. **Mayo Clinic Laboratories.** Pyridoxal 5-Phosphate, Plasma Test Catalog. Available at: https://www.mayocliniclabs.com/test-catalog/Clinical+and+Interpretive/42359 (Accessed January 2026)

28. **Mayo Clinic Laboratories.** Vitamin B6 Profile (Pyridoxal 5-Phosphate and Pyridoxic Acid), Plasma. Available at: https://www.mayocliniclabs.com/test-catalog/overview/42360 (Accessed January 2026)

29. **Synnovis.** The prevalence of low and very high vitamin B6 (Pyridoxal 5'-phosphate) concentrations in hospital patients. Available at: https://www.synnovis.co.uk/articles-and-papers/the-prevalence-of-low-and-very-high-vitamin-b6-pyridoxal-5%E2%80%99-phosphate (Accessed January 2026)

30. **HealthMatters.io.** Pyridoxal 5-Phosphate (PLP) - Lab Results Explained. Available at: https://healthmatters.io/understand-blood-test-results/pyridoxal-5-phosphate-plp (Accessed January 2026)

31. **NCBI PubMed.** Semi-automated fluorometric determination of pyridoxal-5'-phosphate (vitamin B6) in whole blood by HPLC. Available at: https://pubmed.ncbi.nlm.nih.gov/7319721/ (1981)

32. **NCBI PubMed.** Optimisation and validation of a sensitive HPLC assay for routine measurement of pyridoxal 5-phosphate in human plasma and red cells. Available at: https://pubmed.ncbi.nlm.nih.gov/12860041/ (2003)

33. **Health Sciences Authority Singapore.** High-dose vitamin B6 and risk of peripheral neuropathy. Available at: https://www.hsa.gov.sg/announcements/safety-alert/high-dose-vitamin-b6-and-risk-of-peripheral-neuropathy (2024)

34. **Therapeutic Goods Administration (TGA) Australia.** Peripheral neuropathy with supplementary vitamin B6 (pyridoxine). Available at: https://www.tga.gov.au/news/safety-updates/peripheral-neuropathy-supplementary-vitamin-b6-pyridoxine (2023)

35. **NCBI PubMed Central.** Expert Consensus on Vitamin B6 Therapeutic Use for Patients: Guidance on Safe Dosage, Duration and Clinical Management. Available at: https://pmc.ncbi.nlm.nih.gov/articles/PMC12090844/ (2024)

36. **Medscape.** Vitamin B6 Pills Tied to Neuropathy: Stricter Rules Urged. Available at: https://www.medscape.com/viewarticle/vitamin-b6-pills-tied-neuropathy-stricter-rules-urged-2025a10010iq (2025)

37. **Wiley Online Library.** Impact of Regulatory Action on Dose Maximalization for Vitamin B6 Dietary Supplements on the Reporting Pattern for Neuropathy. Available at: https://onlinelibrary.wiley.com/doi/full/10.1002/pds.70108 (2025)

38. **NCBI PubMed Central.** Underestimated pyridoxine consumption and neurotoxicity: a novel manifestation with rheumatologic relevance - a case-based review. Available at: https://pmc.ncbi.nlm.nih.gov/articles/PMC12106547/ (2025)

39. **Brain Foundation.** What is Vitamin B6 Toxicity? Available at: https://brainfoundation.org.au/what-is-vitamin-b6-toxicity/ (Accessed January 2026)

40. **National Institutes of Health (NIH).** Vitamin B6 - Health Professional Fact Sheet. Available at: https://ods.od.nih.gov/factsheets/VitaminB6-HealthProfessional/ (Updated 2024)

41. **UnitsLab.com.** Vitamin B6 (Pyridoxal Phosphate) Unit Conversion Calculator. Available at: https://unitslab.com/node/214 (Accessed January 2026)

42. **OptimalDX.** Vitamin Biomarkers: Vitamin B6. Available at: https://www.optimaldx.com/research-blog/vitamin-biomarkers-vitamin-b6 (Accessed January 2026)

43. **NCBI PubMed Central.** Direct and Functional Biomarkers of Vitamin B6 Status. Available at: https://pmc.ncbi.nlm.nih.gov/articles/PMC5988249/ (2018)

### Homocysteine and Methylation

44. **Frontiers in Nutrition.** Homocysteine—a retrospective and prospective appraisal. Available at: https://www.frontiersin.org/journals/nutrition/articles/10.3389/fnut.2023.1179807/full (2023)

45. **Food for the Brain.** Methylation and Homocysteine: A Brief Overview. Available at: https://foodforthebrain.org/nutrition-for-healthcare-professionals/methylation-and-homocysteine/ (Accessed January 2026)

46. **Clinical Laboratory Diagnostics.** Chapter 13: Homocysteine, vitamin B12, folates, vitamin B6, choline, betaine. Available at: https://clinical-laboratory-diagnostics.com/k13.html (Accessed January 2026)

47. **NCBI PubMed.** Homocysteine metabolism. Available at: https://pubmed.ncbi.nlm.nih.gov/10448523/ (1999)

48. **Springer Nature EPMA Journal.** Homocysteine metabolism as the target for predictive medical approach, disease prevention, prognosis, and treatments tailored to the person. Available at: https://link.springer.com/article/10.1007/s13167-021-00263-0 (2021)

49. **Life Extension.** Homocysteine Reduction: Causes & Treatments Protocol. Available at: https://www.lifeextension.com/protocols/heart-circulatory/homocysteine-reduction (Accessed January 2026)

50. **Rupa Health.** Nutritional Support for Optimal Methylation: A Guide for Functional Medicine. Available at: https://www.rupahealth.com/post/nutritional-support-for-optimal-methylation-a-guide-for-functional-medicine (Accessed January 2026)

51. **NCBI PubMed Central.** Homocysteine, Vitamin B12 and Folate Level: Possible Risk Factors in the Progression of Chronic Heart and Kidney Disorders. Available at: https://pmc.ncbi.nlm.nih.gov/articles/PMC10494275/ (2023)

---

## Document Metadata

**Created By:** Clinical Research Analysis (Claude Code)
**Purpose:** EMR Integration for Plenya System
**Target Users:** Functional Medicine Practitioners, Clinical Laboratory Staff
**Testing Methodology:** Whole Blood HPLC (High-Performance Liquid Chromatography)
**Evidence Level:** Based on peer-reviewed literature (2018-2026), clinical laboratory standards, and regulatory guidance
**Revision Schedule:** Annual review or upon significant new evidence

**Database Import Format:** CSV tables provided for each vitamin compatible with PostgreSQL EMR system

**Clinical Use:** These ranges are intended for clinical decision support in the context of comprehensive patient assessment. Always consider individual patient context, symptoms, medications, and comorbidities. Consult with qualified healthcare providers for medical advice.

---

**END OF DOCUMENT**
