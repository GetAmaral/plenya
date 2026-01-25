# CORREÇÃO: ESTRATIFICAÇÃO DE RISCO DO ESTRADIOL
## Escore Plenya - Análise de Erros e Correções

**Data:** 19 de Janeiro de 2026
**Motivo:** Valores incorretos identificados pelo usuário no CSV

---

## ERROS IDENTIFICADOS NO CSV ORIGINAL

### 1. Estradiol - Homens
**ANTES (INCORRETO):**
```csv
Estradiol - Homens;pg/mL | 1 pg/mL = 3.671 pmol/L;20;<10;10 a 20;20 a 30;21.80 a 30.11;30 a 40;40 a 60;>60
```

**Problemas identificados:**
- ❌ Nível 3 contém valor estranho: "21.80 a 30.11" (provavelmente erro de digitação ou corrupção de dados)
- ❌ Estratificação de risco invertida (níveis altos como ótimos, mas excesso também é prejudicial)
- ❌ Não reflete curva em U (tanto deficiência quanto excesso são ruins)

**DEPOIS (CORRETO):**
```csv
Estradiol - Homens;pg/mL | 1 pg/mL = 3.671 pmol/L;20;<10;10 a 19;20 a 40;41 a 50;51 a 60;>60
```

**Justificativa:**
- ✅ Nível 5 (Ótimo): 20-40 pg/mL - faixa funcional ótima para saúde óssea, cardiovascular, sexual, cognitiva
- ✅ Nível 4 (Bom): 41-50 pg/mL - limítrofe superior, monitorar aromatização
- ✅ Nível 3 (Borderline): 51-60 pg/mL - risco de aromatização excessiva
- ✅ Nível 2 (Elevado): 10-19 pg/mL - subótimo, risco de perda óssea
- ✅ Nível 1 (Alto risco): <10 pg/mL - deficiência severa
- ✅ Nível 0 (Crítico): >60 pg/mL - excesso (risco CV, ginecomastia, mortalidade)

**Referências:**
- Optimal range: 20-40 pg/mL (Finkelstein et al., NEJM 2013; OptimalDX 2024)
- Age-specific medians: 29.6-37.0 pg/mL ages 20-59 (NHANES 2023)
- Bone loss prevention: ≥19 pg/mL (Finkelstein et al.)
- Cardiovascular risk increase: >50 pg/mL (multiple studies)

---

### 2. Estradiol - Mulheres Fase Folicular Inicial
**ANTES (NECESSITA REVISÃO):**
```csv
Estradiol - Mulheres Fase Folicular Inicial;pg/mL | 1 pg/mL = 3.671 pmol/L;20;<15;15 a 30;30 a 50;50 a 100;100 a 150;;
```

**Problemas identificados:**
- ⚠️ Nível ótimo (5) deveria ser 30-60 pg/mL para folicular inicial (dias 1-7)
- ⚠️ Valores >100 pg/mL em fase folicular inicial podem indicar PCOS ou disfunção ovulatória
- ⚠️ Estratificação não reflete risco adequadamente

**DEPOIS (CORRETO):**
```csv
Estradiol - Mulheres Fase Folicular Inicial (Dias 1-7);pg/mL | 1 pg/mL = 3.671 pmol/L;20;<20;20 a 29;30 a 60;61 a 80;81 a 100;>100
```

**Justificativa:**
- ✅ Nível 5 (Ótimo): 30-60 pg/mL - desenvolvimento folicular saudável
- ✅ Nível 4 (Bom): 61-80 pg/mL - limítrofe alto para fase inicial
- ✅ Nível 3 (Aceitável): 81-100 pg/mL - elevado para fase inicial
- ✅ Nível 2 (Subótimo): 20-29 pg/mL - desenvolvimento folicular fraco
- ✅ Nível 1 (Deficiente): <20 pg/mL - insuficiência ovariana ou anovulação
- ✅ Nível 0 (Crítico): >100 pg/mL - investigar PCOS, cistos, disfunção

**Referências:**
- Optimal early follicular: 30-60 pg/mL (OptimalDX 2024)
- Conventional range: 25-75 pg/mL (Mayo Clinic 2024)
- LC-MS/MS: 31-771 pmol/L (Bae et al., JCEM 2019)

---

### 3. Estradiol - Mulheres Fase Ovulatória
**ANTES (NECESSITA REVISÃO):**
```csv
Estradiol - Mulheres Fase Ovulatória;pg/mL | 1 pg/mL = 3.671 pmol/L;20;<100;100 a 150;150 a 250;250 a 500;500 a 750;>750;
```

**Problemas identificados:**
- ⚠️ Nível ótimo deveria ser 200-400 pg/mL (pico ovulatório)
- ⚠️ Valores <100 pg/mL indicam ovulação fraca/ausente
- ⚠️ Valores >500 pg/mL podem indicar hiperestimulação ou PCOS
- ⚠️ Estratificação atual não reflete fisiologia normal

**DEPOIS (CORRETO):**
```csv
Estradiol - Mulheres Fase Ovulatória (Pico Meio do Ciclo);pg/mL | 1 pg/mL = 3.671 pmol/L;20;<100;100 a 199;200 a 400;401 a 500;501 a 600;>600
```

**Justificativa:**
- ✅ Nível 5 (Ótimo): 200-400 pg/mL - pico ovulatório saudável, ovulação confirmada
- ✅ Nível 4 (Bom): 401-500 pg/mL - pico alto mas aceitável
- ✅ Nível 3 (Aceitável): 501-600 pg/mL - muito alto, monitorar
- ✅ Nível 2 (Subótimo): 100-199 pg/mL - ovulação fraca
- ✅ Nível 1 (Deficiente): <100 pg/mL - anovulação provável
- ✅ Nível 0 (Crítico): >600 pg/mL - hiperestimulação, PCOS severo

**Referências:**
- Optimal ovulatory peak: 200-400 pg/mL (OptimalDX 2024)
- Conventional range: 100-400 pg/mL (LabCorp, Quest)
- LC-MS/MS: 900-1600 pmol/L optimal (Bae et al., JCEM 2019)

---

### 4. Estradiol - Mulheres Fase Lútea
**ANTES (NECESSITA REVISÃO):**
```csv
Estradiol - Mulheres Fase Lútea;pg/mL | 1 pg/mL = 3.671 pmol/L;20;<30;30 a 70;70 a 100;100 a 200;200 a 300;>300;
```

**Problemas identificados:**
- ⚠️ Nível ótimo deveria ser 80-150 pg/mL (fase lútea média, dias 19-21)
- ⚠️ Valores >200 pg/mL podem indicar dominância estrogênica (baixa progesterona)
- ⚠️ Estratificação não reflete balanço estrogênio-progesterona

**DEPOIS (CORRETO):**
```csv
Estradiol - Mulheres Fase Lútea (Dias 19-21);pg/mL | 1 pg/mL = 3.671 pmol/L;20;<40;40 a 79;80 a 150;151 a 200;201 a 250;>250
```

**Justificativa:**
- ✅ Nível 5 (Ótimo): 80-150 pg/mL - balanço hormonal saudável com progesterona adequada
- ✅ Nível 4 (Bom): 151-200 pg/mL - aceitável se progesterona proporcional (Pg/E2 >100)
- ✅ Nível 3 (Aceitável): 201-250 pg/mL - elevado, avaliar dominância estrogênica
- ✅ Nível 2 (Subótimo): 40-79 pg/mL - baixo para fase lútea
- ✅ Nível 1 (Deficiente): <40 pg/mL - deficiência lútea, ovulação fraca
- ✅ Nível 0 (Crítico): >250 pg/mL - dominância estrogênica provável (Pg/E2 <100)

**Referências:**
- Optimal mid-luteal: 80-150 pg/mL with progesterone 8-30 ng/mL (OptimalDX 2024)
- Conventional range: 75-300 pg/mL (wide range, not functional optimal)
- Pg/E2 ratio optimal: 100-500 (target ~300)

---

### 5. Estradiol - Mulheres Pós-Menopausa
**ANTES (NECESSITA REVISÃO SIGNIFICATIVA):**
```csv
Estradiol - Mulheres Pós-Menopausa;pg/mL | 1 pg/mL = 3.671 pmol/L;20;<10;10 a 30;30 a 60;60 a 100;100 a 150;>150;
```

**Problemas identificados:**
- ❌ **ERRO CRÍTICO:** Estratificação está INVERTIDA
- ❌ Em mulheres pós-menopausa **SEM TRH**, valores baixos (<25 pg/mL) são ESPERADOS e normais
- ❌ Em mulheres pós-menopausa **COM TRH**, valores 60-150 pg/mL são ÓTIMOS
- ❌ Tabela atual não diferencia SEM TRH vs COM TRH (contextos completamente opostos)

**SOLUÇÃO:** Criar **DUAS TABELAS SEPARADAS**

#### 5a. Pós-Menopausa SEM TRH (Valores Naturais)
**DEPOIS (CORRETO):**
```csv
Estradiol - Mulheres Pós-Menopausa (Sem TRH);pg/mL | 1 pg/mL = 3.671 pmol/L;20;>60;40 a 60;25 a 39;10 a 24;5 a 9;<5
```

**Justificativa:**
- ✅ Nível 5 (Ótimo): 5-9 pg/mL - produção adrenal residual saudável
- ✅ Nível 4 (Bom): 10-24 pg/mL - aceitável, sintomas mínimos
- ✅ Nível 3 (Aceitável): 25-39 pg/mL - elevado (fonte periférica aumentada?)
- ✅ Nível 2 (Subótimo): <5 pg/mL - sintomas vasomotores severos esperados
- ✅ Nível 1 (Anormal): 40-60 pg/mL - investigar fonte (tumor secretor?)
- ✅ Nível 0 (Crítico): >60 pg/mL - patológico, investigar neoplasia

**Referências:**
- Expected range without HRT: 5-25 pg/mL (Mayo Clinic 2024)
- Hot flash threshold: <5 pg/mL (ZRT Laboratory)
- Symptom improvement: ≥15 pg/mL (natural production)

#### 5b. Pós-Menopausa COM TRH (Valores Terapêuticos)
**DEPOIS (CORRETO - NOVA TABELA):**
```csv
Estradiol - Mulheres Pós-Menopausa (Com TRH);pg/mL | 1 pg/mL = 3.671 pmol/L;20;<30;30 a 59;60 a 150;151 a 200;201 a 250;>250
```

**Justificativa:**
- ✅ Nível 5 (Ótimo): 60-150 pg/mL - alívio sintomático + proteção óssea (meta terapêutica)
- ✅ Nível 4 (Bom): 151-200 pg/mL - aceitável, monitorar sintomas
- ✅ Nível 3 (Alto): 201-250 pg/mL - dose alta, ajustar se assintomática
- ✅ Nível 2 (Subótimo): 30-59 pg/mL - subterapêutico, sintomas podem persistir
- ✅ Nível 1 (Ineficaz): <30 pg/mL - dose inadequada, ajustar TRH
- ✅ Nível 0 (Crítico): >250 pg/mL - suprafisiológico, risco aumentado

**Referências:**
- Symptom relief threshold: 60 pg/mL achieves 50% control (Menopause Journal 2024)
- Real-world median on HRT: 96 pg/mL (transdermal)
- Bone protection: ≥60 pg/mL (NAMS 2022)
- 25% of women on highest dose remain <60 pg/mL (underdosed)

---

## RESUMO DAS CORREÇÕES

| Tabela | Status | Ação |
|--------|--------|------|
| **Homens** | ❌ ERRO CRÍTICO | Corrigir valor "21.80 a 30.11" e re-estratificar |
| **Fase Folicular** | ⚠️ AJUSTE MENOR | Re-estratificar para ótimo 30-60 pg/mL |
| **Fase Ovulatória** | ⚠️ AJUSTE MODERADO | Re-estratificar para ótimo 200-400 pg/mL |
| **Fase Lútea** | ⚠️ AJUSTE MODERADO | Re-estratificar para ótimo 80-150 pg/mL |
| **Pós-Menopausa** | ❌ ERRO CRÍTICO | **DIVIDIR EM DUAS TABELAS** (Sem TRH vs Com TRH) |

---

## TABELAS CORRIGIDAS - FORMATO CSV FINAL

### 1. Homens (CORRIGIDO)
```csv
Estradiol - Homens;pg/mL | 1 pg/mL = 3.671 pmol/L;20;<10;10 a 19;20 a 40;41 a 50;51 a 60;>60
```

### 2. Mulheres Fase Folicular Inicial (CORRIGIDO)
```csv
Estradiol - Mulheres Fase Folicular Inicial (Dias 1-7);pg/mL | 1 pg/mL = 3.671 pmol/L;20;<20;20 a 29;30 a 60;61 a 80;81 a 100;>100
```

### 3. Mulheres Fase Ovulatória (CORRIGIDO)
```csv
Estradiol - Mulheres Fase Ovulatória (Pico Meio do Ciclo);pg/mL | 1 pg/mL = 3.671 pmol/L;20;<100;100 a 199;200 a 400;401 a 500;501 a 600;>600
```

### 4. Mulheres Fase Lútea (CORRIGIDO)
```csv
Estradiol - Mulheres Fase Lútea (Dias 19-21);pg/mL | 1 pg/mL = 3.671 pmol/L;20;<40;40 a 79;80 a 150;151 a 200;201 a 250;>250
```

### 5a. Mulheres Pós-Menopausa SEM TRH (NOVA - CORRIGIDO)
```csv
Estradiol - Mulheres Pós-Menopausa (Sem TRH);pg/mL | 1 pg/mL = 3.671 pmol/L;20;>60;40 a 60;25 a 39;10 a 24;5 a 9;<5
```

### 5b. Mulheres Pós-Menopausa COM TRH (NOVA - TABELA ADICIONAL)
```csv
Estradiol - Mulheres Pós-Menopausa (Com TRH);pg/mL | 1 pg/mL = 3.671 pmol/L;20;<30;30 a 59;60 a 150;151 a 200;201 a 250;>250
```

---

## IMPACTO NO CSV PRINCIPAL

**Linhas a REMOVER (5):**
```
Estradiol - Homens;pg/mL | 1 pg/mL = 3.671 pmol/L;20;<10;10 a 20;20 a 30;21.80 a 30.11;30 a 40;40 a 60;>60
Estradiol - Mulheres Fase Folicular Inicial;pg/mL | 1 pg/mL = 3.671 pmol/L;20;<15;15 a 30;30 a 50;50 a 100;100 a 150;;
Estradiol - Mulheres Fase Ovulatória;pg/mL | 1 pg/mL = 3.671 pmol/L;20;<100;100 a 150;150 a 250;250 a 500;500 a 750;>750;
Estradiol - Mulheres Fase Lútea;pg/mL | 1 pg/mL = 3.671 pmol/L;20;<30;30 a 70;70 a 100;100 a 200;200 a 300;>300;
Estradiol - Mulheres Pós-Menopausa;pg/mL | 1 pg/mL = 3.671 pmol/L;20;<10;10 a 30;30 a 60;60 a 100;100 a 150;>150;
```

**Linhas a ADICIONAR (6):**
```
Estradiol - Homens;pg/mL | 1 pg/mL = 3.671 pmol/L;20;<10;10 a 19;20 a 40;41 a 50;51 a 60;>60
Estradiol - Mulheres Fase Folicular Inicial (Dias 1-7);pg/mL | 1 pg/mL = 3.671 pmol/L;20;<20;20 a 29;30 a 60;61 a 80;81 a 100;>100
Estradiol - Mulheres Fase Ovulatória (Pico Meio do Ciclo);pg/mL | 1 pg/mL = 3.671 pmol/L;20;<100;100 a 199;200 a 400;401 a 500;501 a 600;>600
Estradiol - Mulheres Fase Lútea (Dias 19-21);pg/mL | 1 pg/mL = 3.671 pmol/L;20;<40;40 a 79;80 a 150;151 a 200;201 a 250;>250
Estradiol - Mulheres Pós-Menopausa (Sem TRH);pg/mL | 1 pg/mL = 3.671 pmol/L;20;>60;40 a 60;25 a 39;10 a 24;5 a 9;<5
Estradiol - Mulheres Pós-Menopausa (Com TRH);pg/mL | 1 pg/mL = 3.671 pmol/L;20;<30;30 a 59;60 a 150;151 a 200;201 a 250;>250
```

**Mudança líquida:** +1 linha (agora 6 linhas em vez de 5)
**Total esperado no CSV:** 151 linhas (150 atual + 1 nova)

---

## REFERÊNCIAS PRINCIPAIS (2023-2026)

1. **Finkelstein JS, et al. (2013).** Gonadal Steroids and Body Composition, Strength, and Sexual Function in Men. *N Engl J Med*. 369:1011-1022.
   - Estudo definitivo sobre limiares de estradiol em homens

2. **NHANES Data (2023).** Age-Specific Serum Estradiol Concentrations in Healthy Men. *PMC6749840*.
   - Dados populacionais de referência por faixa etária

3. **Bae YJ, et al. (2019).** Reference Intervals of Nine Steroid Hormones via LC-MS/MS. *J Clin Endocrinol Metab*. 104(9):3089-3100.
   - Intervalos de referência LC-MS/MS (método gold standard)

4. **OptimalDX (2024).** Hormone Biomarkers: Estradiol in Men and Women.
   - Faixas funcionais ótimas (medicina funcional)

5. **Menopause Journal (2024).** The Range and Variation in Serum Estradiol During Menopausal Hormone Therapy. *Menopause*. 32(2):107-113.
   - Dados do mundo real de TRH, limiar terapêutico 60 pg/mL

6. **ZRT Laboratory (2024).** Progesterone to Estradiol (Pg/E2) Ratio Reference Ranges.
   - Importância do balanço Pg/E2

7. **Institute for Functional Medicine (2024).** Perimenopause & Menopause Protocol.
   - Abordagem funcional para reposição hormonal

---

**Documento compilado:** 19/01/2026
**Sistema:** Plenya EMR
**Propósito:** Correção de erros no CSV de estratificação de risco
