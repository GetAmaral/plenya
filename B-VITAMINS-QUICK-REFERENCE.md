# B Vitamins Risk Stratification - Quick Reference Guide

**For:** Whole Blood HPLC Measurements
**Version:** 1.0 (January 2026)

---

## CSV Import Formats (Ready for Database)

### Vitamin B1 (Thiamine Diphosphate) - Whole Blood

```csv
Vitamina B1 (Tiamina Difosfato) - Sangue Total;nmol/L | µg/L: ×0.425;20;0. Deficiência Grave (< 40): CRÍTICO - Risco de Wernicke-Korsakoff e beribéri;1. Deficiência Moderada (40-69): ALTO RISCO - Beribéri subclínico;2. Deficiência Leve (70-84): RISCO MODERADO - Insuficiência funcional;3. Subótimo (85-120): BAIXO RISCO - Abaixo do ótimo;4. Ótimo (121-180): IDEAL - Função metabólica adequada;5. Ótimo Alto (181-250): EXCELENTE - Reservas superiores;6. Suprafisiológico (> 250): MONITORAR - Níveis muito elevados
```

**Optimal Range:** 121-180 nmol/L
**No known toxicity** - water-soluble with excellent safety profile

---

### Vitamin B2 (Riboflavin) - Whole Blood

```csv
Vitamina B2 (Riboflavina) - Sangue Total;nmol/L | µg/L: ×0.376;20;0. Deficiência Grave (< 80): CRÍTICO - Ariboflavinose clínica;1. Deficiência Moderada (80-120): ALTO RISCO - MTHFR comprometida;2. Deficiência Leve (121-160): RISCO MODERADO - Metilação subótima;3. Subótimo (161-220): BAIXO RISCO - Abaixo do ótimo para MTHFR;4. Ótimo (221-350): IDEAL - Atividade plena de MTHFR;5. Ótimo Alto (351-500): EXCELENTE - Suporte superior de metilação;6. Suprafisiológico (> 500): MONITORAR - Níveis muito elevados
```

**Optimal Range:** 221-350 nmol/L
**Therapeutic Dose for Migraines:** 400mg/day (achieves 351-500 nmol/L range)
**No known toxicity** - bright yellow urine is normal

---

### Vitamin B6 (Pyridoxal-5-Phosphate) - Whole Blood

```csv
Vitamina B6 (Piridoxal-5-Fosfato) - Sangue Total;nmol/L | ng/mL: ×0.247;20;0. Deficiência Grave (< 20): CRÍTICO - Neuropatia e convulsões;1. Deficiência Moderada (20-40): ALTO RISCO - Homocisteína elevada;2. Deficiência Leve (41-60): RISCO MODERADO - Metilação subótima;3. Subótimo (61-90): BAIXO RISCO - Abaixo do ótimo;4. Ótimo (91-180): IDEAL - Atividade enzimática plena;5. Ótimo Alto (181-300): BOM - Faixa terapêutica;6. Limite Superior Seguro (301-500): CAUTELA - Monitorar neuropatia;7. Potencialmente Tóxico (> 500): ALERTA - Alto risco de neuropatia
```

**Optimal Range:** 91-180 nmol/L
**CAUTION:** Risk of peripheral neuropathy at >200mg/day chronically
**Upper Safe Limit:** 301-500 nmol/L (monitor symptoms)
**Potentially Toxic:** >500 nmol/L (stop supplementation)

---

## Quick Reference Tables

### Vitamin B1 (Thiamine) - Clinical Quick Guide

| Range (nmol/L) | Status | Red Flags | Action |
|----------------|--------|-----------|--------|
| **< 40** | CRITICAL | Wernicke-Korsakoff, beriberi | Emergency IV thiamine 200-500mg TID |
| **40-69** | HIGH RISK | Subclinical beriberi | High-dose oral 100-300mg daily |
| **70-84** | MODERATE | Functional insufficiency | 50-150mg daily + investigate cause |
| **85-120** | SUBOPTIMAL | Below optimal for mitochondria | 25-100mg daily or B-complex |
| **121-180** | OPTIMAL ✓ | - | Maintenance 10-50mg daily |
| **181-250** | HIGH OPTIMAL | - | Continue current regimen |
| **> 250** | MONITOR | Very high from supplementation | Review necessity, no toxicity |

**Key Risks:** Alcoholism, diuretics, diabetes, post-bariatric surgery, pregnancy
**Clinical Emergency:** Always give thiamine BEFORE glucose in at-risk patients

---

### Vitamin B2 (Riboflavin) - Clinical Quick Guide

| Range (nmol/L) | Status | Red Flags | Action |
|----------------|--------|-----------|--------|
| **< 80** | CRITICAL | Ariboflavinosis, cheilosis | 50-100mg daily + B-complex |
| **80-120** | HIGH RISK | MTHFR impaired, ↑homocysteine | 25-100mg daily, check MTHFR genotype |
| **121-160** | MODERATE | Suboptimal methylation | 10-50mg daily + B-complex |
| **161-220** | SUBOPTIMAL | Below optimal for MTHFR | 10-25mg daily |
| **221-350** | OPTIMAL ✓ | - | Maintenance 5-10mg daily |
| **351-500** | HIGH OPTIMAL | Migraine prevention range | Continue if therapeutic indication |
| **> 500** | MONITOR | Recent high-dose intake | Review necessity, no toxicity |

**Key Applications:** MTHFR C677T support, migraine prevention (400mg/day), methylation optimization
**Bright Yellow Urine:** Normal with supplementation (excess excreted safely)

---

### Vitamin B6 (PLP) - Clinical Quick Guide

| Range (nmol/L) | Status | Red Flags | Action |
|----------------|--------|-----------|--------|
| **< 20** | CRITICAL | Seizures, neuropathy, anemia | High-dose P5P 50-100mg daily |
| **20-40** | HIGH RISK | ↑Homocysteine, depression | P5P 25-75mg daily, check medications |
| **41-60** | MODERATE | Mood dysregulation, PMS | P5P 10-50mg daily + B-complex |
| **61-90** | SUBOPTIMAL | Below optimal for neurotransmitters | P5P 10-25mg daily |
| **91-180** | OPTIMAL ✓ | - | Maintenance 5-15mg daily |
| **181-300** | HIGH OPTIMAL | Therapeutic for PMS, carpal tunnel | Continue if indicated |
| **301-500** | CAUTION ⚠️ | Approaching upper safety limit | Reduce to <100mg, monitor symptoms |
| **> 500** | TOXIC 🚨 | High neuropathy risk | STOP supplements, neurological exam |

**CRITICAL WARNING:** Both deficiency AND excess cause peripheral neuropathy
**Neuropathy Risk:** >200mg/day for >2 months, some susceptible at 50-100mg/day
**Safe Supplementation:** Use P5P (active form) at lower doses (<50mg)

---

## Methylation Trinity: B2, B6, B12 + Folate

**For Optimal Homocysteine Metabolism (<8 µmol/L):**

1. **B2 (Riboflavin/FAD):** Cofactor for MTHFR
   - Critical for C677T carriers (30-40% of population)
   - Target: 221-350 nmol/L

2. **B6 (PLP):** Transsulfuration pathway
   - Converts homocysteine → cysteine
   - Target: 91-180 nmol/L

3. **B12 (Methylcobalamin):** Methionine synthase cofactor
   - Remethylation pathway
   - Target: >400 pmol/L (not covered in this document)

4. **Folate (5-MTHF):** Methyl donor
   - Used by methionine synthase
   - Target: >900 nmol/L (not covered in this document)

**Clinical Approach:**
- Test all four markers together
- Elevated homocysteine = check all B vitamins
- MTHFR C677T = emphasize B2 + methylfolate
- Cardiovascular risk = comprehensive B support

---

## Conversion Factors (Quick Reference)

### Thiamine (B1)
- **nmol/L × 0.425 = µg/L**
- **µg/L × 2.353 = nmol/L**
- Molecular weight: 425.3 g/mol (ThDP)

### Riboflavin (B2)
- **nmol/L × 0.376 = µg/L**
- **µg/L × 2.66 = nmol/L**
- Molecular weight: 376.4 g/mol

### Pyridoxal-5-Phosphate (B6)
- **nmol/L × 0.247 = ng/mL**
- **ng/mL × 4.05 = nmol/L**
- Molecular weight: 247.1 g/mol

---

## Supplementation Quick Guide

### Recommended Forms

| Vitamin | Standard | Enhanced/Active | Preferred |
|---------|----------|-----------------|-----------|
| **B1** | Thiamine HCl | Benfotiamine (fat-soluble) | Benfotiamine for diabetes/neuropathy |
| **B2** | Riboflavin | Riboflavin-5-phosphate | Standard adequate (400mg for migraines) |
| **B6** | Pyridoxine HCl | **P5P (active)** | **Always use P5P** to minimize neuropathy risk |

### Dosing Principles

✓ **Start low, increase gradually**
✓ **Use active forms** (especially P5P for B6)
✓ **Retest after 3 months**
✓ **Consider B-complex** for synergy
✓ **Address root cause** (diet, absorption, medications)

### Safety Limits

- **B1:** No established toxicity (water-soluble)
- **B2:** No established toxicity (bright yellow urine normal)
- **B6:** ⚠️ **CAUTION >200mg/day** - neuropathy risk increases

---

## High-Risk Populations

### Thiamine (B1) Deficiency Risk
- Chronic alcohol use (most common)
- Diabetes (increased urinary loss)
- Heart failure patients
- Chronic diuretic use (furosemide)
- Post-bariatric surgery
- Hyperemesis gravidarum
- Chronic diarrhea/malabsorption

### Riboflavin (B2) Deficiency Risk
- MTHFR C677T homozygous (TT genotype)
- Vegan/vegetarian diets
- Chronic migraines
- Malabsorption disorders
- Elderly populations
- High oxidative stress conditions

### Pyridoxal-5-Phosphate (B6) Issues
**Deficiency Risk:**
- Oral contraceptive users (30-50% reduction)
- PPI/H2 blocker use
- Renal disease
- Alcoholism
- Pregnancy (increased requirements)

**Excess/Toxicity Risk:**
- High-dose supplement users (>200mg/day)
- Multivitamin stacking (check total B6 intake)
- IV therapy recipients
- Individuals sensitive to pyridoxine

---

## Clinical Red Flags

### Emergency Situations

**Thiamine (B1):**
- 🚨 Confusion + ataxia + eye movement abnormalities = **Wernicke Encephalopathy**
- 🚨 High-output heart failure in alcoholic = **Wet Beriberi**
- 🚨 Give thiamine BEFORE glucose in at-risk patients (prevents Wernicke)

**B6 Toxicity:**
- 🚨 New-onset tingling/burning in hands/feet = **Stop B6 immediately**
- 🚨 Whole blood PLP >500 nmol/L = **Neuropathy risk**
- 🚨 Check total daily B6 intake (all sources)

### Laboratory Alerts

| Finding | Action Required |
|---------|-----------------|
| Thiamine <40 nmol/L | Consider IV thiamine if symptomatic |
| All three B vitamins low | Screen for malabsorption, B12, folate |
| Elevated homocysteine + low B2/B6 | Comprehensive methylation support |
| B6 >300 nmol/L | Reduce supplementation, monitor symptoms |
| B6 >500 nmol/L | Stop supplements, neurological evaluation |

---

## Testing Recommendations

### When to Test Whole Blood B Vitamins

**Mandatory:**
- Chronic fatigue/fibromyalgia workup
- Elevated homocysteine (any level >8 µmol/L)
- Peripheral neuropathy of unknown cause
- Chronic migraine
- Pre-conception/pregnancy planning
- MTHFR genetic variants
- Post-bariatric surgery follow-up
- Chronic alcohol use (any level)

**Consider:**
- Depression/anxiety refractory to treatment
- Cognitive impairment in elderly
- Cardiovascular disease risk assessment
- Diabetes with complications
- Malabsorption disorders
- Long-term medication use (PPIs, diuretics, OCPs)

### Retest Intervals

- **Baseline Deficiency:** Retest after 3 months of supplementation
- **Optimization:** Retest after 3-6 months
- **Maintenance:** Annual testing
- **High-dose B6:** Retest every 3-6 months (safety monitoring)

---

## Database Integration Notes

### Field Structure for EMR

```sql
-- Suggested PostgreSQL table structure

CREATE TABLE lab_b_vitamins_whole_blood (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    test_date DATE NOT NULL,

    -- Thiamine (B1)
    thiamine_value DECIMAL(6,2),  -- nmol/L
    thiamine_unit VARCHAR(20) DEFAULT 'nmol/L',
    thiamine_status VARCHAR(50),  -- e.g., "Optimal", "Deficiency Moderate"
    thiamine_risk_level INTEGER,  -- 0-6

    -- Riboflavin (B2)
    riboflavin_value DECIMAL(6,2),  -- nmol/L
    riboflavin_unit VARCHAR(20) DEFAULT 'nmol/L',
    riboflavin_status VARCHAR(50),
    riboflavin_risk_level INTEGER,  -- 0-6

    -- Pyridoxal-5-Phosphate (B6)
    plp_value DECIMAL(6,2),  -- nmol/L
    plp_unit VARCHAR(20) DEFAULT 'nmol/L',
    plp_status VARCHAR(50),
    plp_risk_level INTEGER,  -- 0-7 (B6 has extra toxicity level)

    -- Clinical context
    laboratory VARCHAR(200),
    method VARCHAR(100) DEFAULT 'HPLC',
    notes TEXT,

    -- Metadata
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_by UUID REFERENCES users(id)
);

-- Risk stratification function example
CREATE OR REPLACE FUNCTION calculate_thiamine_risk(value DECIMAL)
RETURNS INTEGER AS $$
BEGIN
    RETURN CASE
        WHEN value < 40 THEN 0   -- Critical
        WHEN value < 70 THEN 1   -- High Risk
        WHEN value < 85 THEN 2   -- Moderate
        WHEN value < 121 THEN 3  -- Suboptimal
        WHEN value < 181 THEN 4  -- Optimal
        WHEN value < 251 THEN 5  -- High Optimal
        ELSE 6                   -- Supraphysiological
    END;
END;
$$ LANGUAGE plpgsql IMMUTABLE;
```

---

## Full Documentation

For comprehensive clinical details, mechanisms of action, extensive references, and detailed risk stratification tables, see:

**📄 B-VITAMINS-RISK-STRATIFICATION.md** (Full document with 51 references)

---

**Created:** January 2026
**For:** Plenya EMR System
**Purpose:** Clinical decision support for whole blood HPLC B vitamin interpretation
**Evidence Base:** Peer-reviewed literature 2018-2026 + clinical laboratory standards

---

**END OF QUICK REFERENCE**
