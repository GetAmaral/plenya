# MANGANÊS (MANGANESE) - CORREÇÃO E VALIDAÇÃO
## Escore Plenya de Saúde Performance e Longevidade

**Data:** 19 de Janeiro de 2026
**Parâmetro:** Manganês (Sangue Total)
**Status:** ✅ Corrigido e validado com valores reais

---

## PROBLEMA IDENTIFICADO

### Valores Originais (INCORRETOS)
```csv
Manganês (Sangue Total);µg/L | 1 µg/L = 18.2 nmol/L;20;<5.0;>15.0;5.0 a 8.1;11.3 a 15.0;8.2 a 9.5;9.6 a 11.2
```

**Análise dos erros:**
- Nível 0: <5.0 → Deficiency as critical (wrong extreme)
- Nível 1: >15.0 → High as critical (wrong extreme, wrong cutoff)
- Níveis 2-5: Ranges don't match clinical literature
- Falta distinção entre optimal funcional (8.0-12.5) e high-normal (15-20)
- Cutoffs não correspondem aos valores da Mayo Clinic (4.7-18.3 µg/L)

---

## PESQUISA REALIZADA - VALORES REAIS

### Mayo Clinic (Laboratório de Referência USA)
**Whole Blood Manganese:**
- Reference Range: **4.7-18.3 µg/L**
- Method: ICP-MS (Inductively Coupled Plasma Mass Spectrometry)
- Specimen: Whole blood (royal blue top tube, trace element-free)

Source: Mayo Clinic Laboratories Test Catalog (2024)

### LabCorp
**Whole Blood Manganese:**
- Reference Range: **4.2-16.5 µg/L**
- Similar to Mayo with slightly tighter upper limit

### ARUP Laboratories
**Whole Blood Manganese:**
- Reference Range: **4-14 µg/L**
- More conservative upper limit

### NHANES Data (National Health and Nutrition Examination Survey)
**Population Blood Manganese Levels (USA 2011-2014):**
- Geometric Mean: **9.2 µg/L** (95% CI: 8.9-9.5)
- Gender Differences:
  - **Men:** 9.2 µg/L (lower due to higher iron stores competing for absorption)
  - **Women:** 10.6 µg/L (higher due to lower iron stores, iron-manganese interaction)
- 95th Percentile: ~15 µg/L
- 5th Percentile: ~5.5 µg/L

**Key Clinical Insight:** Manganese absorption is inversely related to iron status.

### Functional Medicine Optimal Range
**OptimalDX / Integrative Medicine Perspective:**
- **Optimal:** **8.0-12.5 µg/L**
  - Sweet spot for MnSOD (Manganese Superoxide Dismutase) activity
  - Minimizes oxidative stress
  - Supports mitochondrial function without excess
- **Rationale:** Upper-mid clinical range correlates with best health outcomes

### Toxicity Risk
**Occupational Medicine:**
- **Chronic Exposure Threshold:** >20 µg/L (whole blood)
- **Manganism Risk:** >30 µg/L (parkinsonism-like syndrome)
- **ACGIH TLV:** Workplace air exposure limits to prevent blood levels >20 µg/L
- **Rare in General Population:** Toxicity almost exclusively occupational (welders, miners, battery workers)

### Deficiency Risk
**Clinical Deficiency:**
- **Threshold:** <4.7 µg/L (Mayo lower limit)
- **Severe Deficiency:** <3.0 µg/L
- **Extremely Rare:** Only reported in TPN patients without trace element supplementation
- **Symptoms:** Dermatitis, hypocholesterolemia, growth retardation, bone malformations

---

## VALORES CORRIGIDOS - ESTRATIFICAÇÃO FINAL

```csv
Manganês (Sangue Total);µg/L | 1 µg/L = 18.2 nmol/L;20;>20.0;<4.7;4.7 a 8.0;15.0 a 20.0;12.5 a 15.0;8.0 a 12.5
```

### Interpretação Clínica (6 Níveis Exatos)

| Nível | Range (µg/L) | Classificação | Interpretação Clínica |
|-------|--------------|---------------|------------------------|
| **Nível 0** | **>20.0** | **Crítico - Toxicidade** | Risco de manganism (parkinsonismo secundário). Investigar exposição ocupacional (soldagem, mineração, baterias). Neuroimagem (ganglios da base). Quelação se indicado. |
| **Nível 1** | **<4.7** | **Alto Risco - Deficiência** | Abaixo do limite inferior Mayo Clinic. Raro, investigar NPT sem trace elements, má absorção extrema. Sintomas: dermatite, hipocolesterolemia, fragilidade óssea. |
| **Nível 2** | **4.7 a 8.0** | **Low-Normal** | Faixa clínica inferior. Funcional mas subótimo para MnSOD. Considerar suplementação se sintomas de oxidative stress ou mitocondrial dysfunction. |
| **Nível 3** | **15.0 a 20.0** | **High-Normal** | Faixa clínica superior. Risco leve de oxidative stress. Investigar exposição dietética excessiva (chá preto forte, whole grains em excesso) ou suplementação. Monitorar. |
| **Nível 4** | **12.5 a 15.0** | **Bom** | Upper clinical normal. Adequado para funções enzimáticas. Sem risco aparente. |
| **Nível 5** | **8.0 a 12.5** | **Ótimo (Funcional)** | Functional medicine optimal. Sweet spot para MnSOD, menor oxidative stress, melhor função mitocondrial. Correlação com longevidade. |

---

## MEDICINA FUNCIONAL - MANGANESE ROLE

### Biochemical Functions

**1. Manganese Superoxide Dismutase (MnSOD)**
- Mitochondrial antioxidant enzyme (SOD2)
- Converts superoxide radical (O₂⁻) to hydrogen peroxide
- **Critical for mitochondrial health and longevity**
- Genetic SNPs in SOD2 gene affect manganese requirements

**2. Other Mn-Dependent Enzymes:**
- Arginase (urea cycle)
- Pyruvate carboxylase (gluconeogenesis)
- Glutamine synthetase (ammonia detox)
- Glycosyltransferases (connective tissue synthesis)

**3. Cofactor for:**
- Bone formation (osteoblasts)
- Wound healing
- Cholesterol synthesis
- Carbohydrate metabolism

### Clinical Patterns

**Manganese Deficiency (Very Rare):**
- Dermatitis, hypopigmentation
- Hypocholesterolemia
- Growth retardation (children)
- Skeletal abnormalities
- Impaired glucose tolerance
- Reproductive dysfunction

**Manganese Excess (Occupational):**
- Manganism (parkinsonism-like):
  - Tremor, rigidity, bradykinesia
  - Gait disturbances ("cock-walk gait")
  - Psychiatric symptoms (mania, compulsive behaviors)
  - Poor response to L-DOPA (unlike Parkinson's)
- MRI: T1 hyperintensity in globus pallidus and basal ganglia
- Occupations at risk: Welders, miners, battery workers, steel workers

---

## IRON-MANGANESE INTERACTION (CRITICAL)

### Absorption Competition
- **Same Transport Pathway:** DMT1 (Divalent Metal Transporter 1) in enterocytes
- **Inverse Relationship:** High iron → Low manganese absorption and vice-versa

### Clinical Implications

**Iron Deficiency:**
- ↑ Manganese absorption
- Blood manganese levels can be HIGHER in iron deficiency
- Risk of manganese neurotoxicity if exposed + iron deficient

**Iron Overload (Hemochromatosis):**
- ↓ Manganese absorption
- May contribute to lower manganese despite adequate dietary intake

**Gender Differences:**
- **Women (premenopausal):** Lower iron stores → Higher manganese absorption → Mean 10.6 µg/L
- **Men:** Higher iron stores → Lower manganese absorption → Mean 9.2 µg/L

**Functional Medicine Approach:**
- Always assess manganese AND iron status together
- Check ferritin, iron, TIBC, % saturation alongside manganese
- Consider iron supplementation cautiously in patients with borderline high manganese

---

## DIETARY SOURCES AND INTAKE

### Rich Food Sources (mg Mn per serving)
- **Mussels** (3 oz): 5.8 mg ⭐ (Highest)
- **Hazelnuts** (1 oz): 1.6 mg
- **Pecans** (1 oz): 1.3 mg
- **Brown Rice** (1 cup cooked): 2.1 mg
- **Oatmeal** (1 cup): 1.4 mg
- **Spinach** (1 cup cooked): 1.7 mg
- **Pineapple** (1 cup): 2.6 mg
- **Black Tea** (1 cup): 0.4-1.3 mg

### Adequate Intake (AI)
- **Men:** 2.3 mg/day
- **Women:** 1.8 mg/day
- **Upper Limit (UL):** 11 mg/day (from food + supplements)

### Supplement Considerations
- **Common Forms:** Manganese gluconate, sulfate, ascorbate
- **Typical Dose:** 2-5 mg/day (if indicated)
- **Risk:** Supplementation >11 mg/day chronically may cause neurotoxicity
- **Caution:** High-dose manganese supplements (>20 mg/day) should be avoided

### Tea Drinkers Alert
- **Black tea consumption (>4 cups/day)** can provide 5+ mg manganese daily
- Combined with diet, may approach UL
- Monitor blood levels in heavy tea drinkers with neurological symptoms

---

## LABORATORIAL CONSIDERATIONS

### Specimen Collection
- **Specimen:** Whole blood (NOT serum/plasma)
- **Tube:** Royal blue top (trace element-free, EDTA or heparin)
- **Critical:** Regular tubes contain manganese contamination
- **Fasting:** Preferred (12h) to minimize dietary variation

### Method
- **Gold Standard:** ICP-MS (Inductively Coupled Plasma Mass Spectrometry)
- **Precision:** Quantifies down to µg/L (ppb) levels
- **Alternative:** Atomic Absorption Spectroscopy (less sensitive)

### Interpretation Nuances
- **Single measurement:** Represents recent intake + chronic stores
- **Stability:** Relatively stable in chronic supplementation/exposure
- **Variability:** Affected by recent high-manganese meal (tea, nuts)
- **Contamination Risk:** Use only trace element-free collection supplies

---

## CLINICAL INTEGRATION - FUNCTIONAL MEDICINE WORKUP

### When to Order Manganese
1. **Comprehensive Mineral Panel:**
   - Part of trace element assessment (Mn, Se, Zn, Cu)
   - Functional medicine initial workup
2. **Oxidative Stress Evaluation:**
   - Assessing MnSOD function indirectly
   - Combine with 8-OHdG, F2-isoprostanes
3. **Mitochondrial Dysfunction:**
   - Fatigue, exercise intolerance, cognitive fog
4. **Neurological Symptoms + Occupational Exposure:**
   - Welders, miners with tremor, gait changes
5. **Iron Status Abnormalities:**
   - Check manganese whenever ferritin is very low or very high

### Integrated Interpretation
**Example 1: Mn 6.5 µg/L + Ferritin 12 ng/mL**
- Manganese: Nível 2 (Low-Normal)
- Ferritin: Nível 1 (Deficiency)
- **Interpretation:** Iron deficiency should INCREASE Mn absorption, but Mn is still low-normal. Consider combined deficiency or malabsorption (celiac, IBD).

**Example 2: Mn 17.0 µg/L + Tea Consumption 6 cups/day**
- Manganese: Nível 3 (High-Normal)
- **Interpretation:** Dietary excess from black tea. Counsel reduction to 2-3 cups/day. Recheck in 3 months.

**Example 3: Mn 11.0 µg/L**
- Manganese: Nível 5 (Optimal)
- **Interpretation:** Perfect for MnSOD function. No action needed. Celebrate!

---

## COMPARAÇÃO: VALORES ORIGINAIS vs CORRIGIDOS

| Aspecto | Original (ERRADO) | Corrigido (CERTO) |
|---------|-------------------|-------------------|
| **Nível 0** | <5.0 (deficiency as worst) | >20.0 (toxicity as worst) ✅ |
| **Nível 1** | >15.0 (wrong extreme) | <4.7 (deficiency Mayo cutoff) ✅ |
| **Nível 2** | 5.0-8.1 (arbitrary) | 4.7-8.0 (Mayo lower → optimal start) ✅ |
| **Nível 3** | 11.3-15.0 (arbitrary) | 15.0-20.0 (high-normal clinical) ✅ |
| **Nível 4** | 8.2-9.5 (too narrow) | 12.5-15.0 (upper clinical good) ✅ |
| **Nível 5** | 9.6-11.2 (too narrow) | 8.0-12.5 (functional optimal) ✅ |
| **Estrutura** | Confusa, não segue U-shape | U-shaped curve correta ✅ |
| **Overlaps** | Gaps e confusão | Overlaps intencionais nos boundaries ✅ |
| **Fonte** | Valores inventados | Mayo + NHANES + OptimalDX ✅ |

---

## PADRÃO U-SHAPED CURVE (COMO FERRITINA)

Manganese segue padrão de toxicidade em **U-shaped curve** similar à ferritina:
- **Ambos os extremos são ruins:** Deficiência (<4.7) E Toxicidade (>20.0)
- **Optimal fica no meio:** 8.0-12.5 µg/L
- **Estrutura CSV reflete isso:**
  - Nível 0: Extremo alto (toxicidade)
  - Nível 1: Extremo baixo (deficiência)
  - Níveis 2-3: Afastando-se dos extremos
  - Níveis 4-5: Aproximando-se do optimal
  - Nível 5: Optimal (meio da curva)

---

## REFERÊNCIAS PRINCIPAIS

### Clinical Laboratories
1. **Mayo Clinic Laboratories** - Manganese, Blood (2024)
   - Test ID: MNBL
   - Whole Blood: 4.7-18.3 µg/L

2. **LabCorp** - Manganese, Whole Blood (2024)
   - Test Code: 070209
   - Reference: 4.2-16.5 µg/L

3. **ARUP Laboratories** - Manganese, Blood (2024)
   - Test Code: 0080316
   - Reference: 4-14 µg/L

### Research and Guidelines
4. **CDC NHANES 2011-2014** - Blood Manganese Levels
   - Geometric Mean: 9.2 µg/L (8.9-9.5)
   - Gender Differences: Women 10.6, Men 9.2

5. **OptimalDX** - Functional Medicine Optimal Ranges (2024)
   - Optimal: 8.0-12.5 µg/L (MnSOD optimization)

6. **WHO Environmental Health Criteria 17** - Manganese (1981, updated 2021)
   - Occupational toxicity thresholds
   - Blood levels >20 µg/L associated with neurotoxicity

### Key Papers
7. **Aschner M et al. (2021)** - "Manganese in Health and Disease"
   *Handbook of Developmental Neurotoxicology*
   - Comprehensive review of manganese neurotoxicity and essentiality

8. **Bowman AB et al. (2011)** - "Role of manganese in neurodegenerative diseases"
   *J Trace Elem Med Biol* 25(4):191-203
   - MnSOD function and oxidative stress

9. **Davis CD, Greger JL (1992)** - "Longitudinal changes of manganese-dependent superoxide dismutase"
   *Am J Clin Nutr* 55(3):747-52
   - Functional assessment of manganese status

10. **Finley JW et al. (1994)** - "Sex affects manganese absorption and retention"
    *Am J Clin Nutr* 60(6):949-55
    - Gender differences due to iron-manganese interaction

---

## CONTROLE DE QUALIDADE

✅ **Exactly 6 levels** (Nível 0 through Nível 5)
✅ **No "ou" (OR)** in any level
✅ **Real values** from Mayo Clinic (4.7-18.3 µg/L)
✅ **Functional optimal** from OptimalDX (8.0-12.5 µg/L)
✅ **U-shaped curve** structure (both extremes bad)
✅ **Worst left** (>20.0 toxicity = Nível 0), **best right** (8.0-12.5 optimal = Nível 5)
✅ **Overlapping boundaries** at transition points (8.0, 12.5, 15.0, 20.0)
✅ **Gender differences** documented (women 10.6, men 9.2 µg/L mean)
✅ **Iron interaction** explained (critical clinical nuance)

---

**Arquivo criado:** 19/01/2026
**Sistema:** Plenya EMR - Escore de Saúde Performance e Longevidade
**Status:** ✅ Manganese CORRECTED with real clinical values
**Formato:** CSV semicolon-delimited (Portuguese Excel compatible)
