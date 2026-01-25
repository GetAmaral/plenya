# UREIA - CORREÇÃO PARA PADRÃO BRASILEIRO
## Escore Plenya de Saúde Performance e Longevidade

**Data:** 19 de Janeiro de 2026
**Parâmetro:** Ureia (não BUN)
**Status:** ✅ Corrigido para padrão brasileiro com valores reais

---

## PROBLEMA IDENTIFICADO

### Valor Original (INCORRETO - BUN Americano)

```csv
Ureia (BUN);mg/dL | 1 mg/dL = 0.357 mmol/L;20;<6;>30;6 a 10;20 a 30;10 a 13;18 a 20;13 a 18
```

**Análise dos erros CRÍTICOS:**

**1. PARÂMETRO ERRADO - BUN (USA) ao invés de UREIA (Brasil):**
- **BUN (Blood Urea Nitrogen):** Mede apenas o NITROGÊNIO na molécula de ureia
- **UREIA:** Mede a MOLÉCULA COMPLETA de ureia (padrão brasileiro)
- **Não são intercambiáveis!** Valores completamente diferentes.

**2. TINHA 7 NÍVEIS (violação da regra máxima de 6):**
Contagem dos níveis:
- Nível 0: <6
- Nível 1: >30
- Nível 2: 6 a 10
- Nível 3: 20 a 30
- Nível 4: 10 a 13
- Nível 5: 18 a 20
- Nível 6: 13 a 18

**TOTAL: 7 NÍVEIS - VIOLAÇÃO!**

**3. VALORES BUN (não UREIA):**
- Os valores (6, 10, 13, 18, 20, 30) são típicos de **BUN americano**
- Laboratórios brasileiros NÃO usam BUN
- Valores de referência UREIA no Brasil: **15-40 mg/dL** (muito maior que BUN)

**4. ORDEM CONFUSA:**
- Níveis 4, 5, 6 (10-13, 18-20, 13-18) estão fora de ordem lógica
- 13-18 (nível 6) está entre 10-13 (nível 4) e 18-20 (nível 5)

---

## DIFERENÇA ENTRE BUN E UREIA

### Bioquímica

**UREIA (CO(NH₂)₂):**
- Molécula completa: 1 carbono + 1 oxigênio + 2 nitrogênios + 4 hidrogênios
- Peso molecular: 60 g/mol
- **28/60 = 46.7% da molécula é nitrogênio**

**BUN (Blood Urea Nitrogen):**
- Mede APENAS o nitrogênio contido na ureia
- Não mede a molécula completa

### Conversão Matemática

**Fórmula:**
```
Ureia (mg/dL) = BUN (mg/dL) × 2.14
BUN (mg/dL) = Ureia (mg/dL) ÷ 2.14
```

**Por que 2.14?**
- Ureia tem peso molecular 60 g/mol
- Nitrogênio (2 átomos) tem peso 28 g/mol
- Ratio: 60/28 = **2.14**

**Exemplos de Conversão:**

| BUN (USA) | Ureia (Brasil) |
|-----------|----------------|
| 6 mg/dL   | 12.8 mg/dL     |
| 10 mg/dL  | 21.4 mg/dL     |
| 16 mg/dL  | 34.2 mg/dL     |
| 20 mg/dL  | 42.8 mg/dL     |
| 24 mg/dL  | 51.4 mg/dL     |
| 30 mg/dL  | 64.2 mg/dL     |
| 40 mg/dL  | 85.6 mg/dL     |

### Padrão por País

**USA, Canadá, Reino Unido:**
- Usam **BUN (Blood Urea Nitrogen)** em mg/dL
- Valores de referência: 7-20 mg/dL (convencional)

**Brasil, Europa Continental, Austrália:**
- Usam **UREIA** em mg/dL ou mmol/L
- Valores de referência: 15-40 mg/dL (convencional brasileiro)

**Japão, China:**
- Usam UREIA em mmol/L
- Conversão: 1 mg/dL ureia = 0.357 mmol/L

---

## PESQUISA REALIZADA - VALORES REAIS

### Valores de Referência Brasileiros (Convencionais)

**Laboratórios Brasileiros 2024:**

**Adultos (Geral):**
- **10-50 mg/dL** (range mais amplo, maioria dos labs)
- **15-40 mg/dL** (range mais comum)
- **Alguns labs:** 7-20 mg/dL (minoria)

**Por Idade e Gênero (Laboratórios Brasileiros):**

**Adultos 20-50 anos:**
- **Homens:** 19-44 mg/dL
- **Mulheres:** 15-40 mg/dL

**Adultos >50 anos:**
- **Homens:** 18-55 mg/dL
- **Mulheres:** 21-43 mg/dL

**Crianças:**
- **5-18 mg/dL** (dependendo idade)

**Fonte:** [Laboratório Goes](https://laboratoriogoes.com.br/glossario/creatinina-ureia-valores-referencia/), [Lavoisier](https://lavoisier.com.br/saude/ureia), [Labs a+ Medicina Diagnóstica](https://www.labsamais.com.br/noticias/ureia-funcao-referencia-coleta/)

---

### Valores OptimalDX 2024 (Medicina Funcional)

**OptimalDX usa BUN (americano), então precisamos CONVERTER para UREIA brasileira:**

**OptimalDX BUN Ranges:**
- **Optimal:** 10-16 mg/dL
- **Below Optimal:** 6-10 mg/dL
- **Above Optimal:** 16-24 mg/dL
- **Low:** <6 mg/dL
- **High:** >24 mg/dL

**Conversão para UREIA (mg/dL) = BUN × 2.14:**

| OptimalDX BUN | Classificação | Ureia (Brasil) |
|---------------|---------------|----------------|
| 10-16 mg/dL   | **Optimal**   | **21.4-34.2 mg/dL** |
| 6-10 mg/dL    | Below Optimal | 12.8-21.4 mg/dL |
| 16-24 mg/dL   | Above Optimal | 34.2-51.4 mg/dL |
| <6 mg/dL      | Low           | <12.8 mg/dL    |
| >24 mg/dL     | High          | >51.4 mg/dL    |

**Fonte:** [OptimalDX BUN Research](https://www.optimaldx.com/research-blog/protein-biomarkers-blood-urea-nitrogen-bun), [OptimalDX Functional Health Report 2024](https://sac-nd.com/wp-content/uploads/2024/02/JaneSample-OptimalDX-Practitioner-Feb-03-2024.pdf)

---

### Valores Críticos

**Uremia (Insuficiência Renal):**
- **Ureia >80 mg/dL:** Uremia moderada
- **Ureia >100 mg/dL:** Uremia severa
- **Ureia >200 mg/dL:** Uremia extrema (indicação diálise emergencial)

**Azotemia Pré-Renal (Desidratação):**
- **Ureia >50 mg/dL** com creatinina normal
- Ratio Ureia:Creatinina >20:1

**Desnutrição Proteica:**
- **Ureia <13 mg/dL** (equivalente BUN <6)
- Kwashiorkor, marasmo, caquexia

---

## VALORES CORRIGIDOS - ESTRATIFICAÇÃO FINAL

```csv
Ureia;mg/dL | 1 mg/dL = 0.357 mmol/L (ureia) | BUN = Ureia ÷ 2.14;20;>80;<13;52 a 80;13 a 21;35 a 51;22 a 34
```

### Interpretação Clínica (6 Níveis Exatos)

| Nível | Ureia (mg/dL) | Classificação | Interpretação Clínica |
|-------|---------------|---------------|------------------------|
| **Nível 0** | **>80** | **Crítico - Uremia** | Insuficiência renal aguda ou crônica descompensada. Uremia (síndrome clínica: náusea, vômito, confusão, prurido, pericardite urêmica). Creatinina provavelmente >3.0 mg/dL. Considerar diálise se >100 mg/dL + sintomas. URGENTE nefrologia. |
| **Nível 1** | **<13** | **Crítico - Desnutrição** | Desnutrição proteico-calórica severa, caquexia, kwashiorkor, anorexia nervosa, insuficiência hepática (↓ síntese ureia), sobrehidratação extrema. Investigar albumina, pré-albumina, transferrina. URGENTE. |
| **Nível 2** | **52-80** | **Elevado - Insuf Renal** | Insuficiência renal moderada (CKD estágio 3-4), azotemia pré-renal severa (desidratação, choque), sangramento GI alto (↑ proteína absorvida), catabolismo proteico (sepse, queimadura). Investigar creatinina, TFG, ratio ureia:creatinina. |
| **Nível 3** | **13-21** | **Baixo-Normal** | Limite inferior funcional. Dieta low-protein (vegetariana, vegan), sobrehidratação, gravidez (↑ TFG dilui ureia), liver disease leve (↓ síntese ureia). OptimalDX Below Optimal (BUN <10). Investigar proteína dietética, função hepática. |
| **Nível 4** | **35-51** | **Alto-Normal** | Limite superior funcional. Desidratação leve-moderada, dieta high-protein (carnívora, atletas), azotemia pré-renal (desidratação, ICC, depleção volume), catabolismo ↑ (febre, infecção, corticoides). OptimalDX Above Optimal (BUN 16-24). Reidratar, reduzir proteína se sintomático. |
| **Nível 5** | **22-34** | **Optimal (Funcional)** | Ureia ideal OptimalDX 2024 (equivalente BUN 10-16 mg/dL). Hidratação adequada, ingestão proteica balanceada (0.8-1.2 g/kg/dia), função renal normal (TFG >90), função hepática normal. Metabolismo nitrogenado ótimo. |

---

## MEDICINA FUNCIONAL - INTERPRETAÇÃO INTEGRADA

### 1. FUNÇÃO RENAL - TFG (Taxa de Filtração Glomerular)

**Ureia é marcador MENOS sensível que creatinina para TFG:**
- **Creatinina:** Reflexo direto da filtração glomerular (melhor marcador)
- **Ureia:** Afetada por MÚLTIPLOS fatores além da TFG (proteína dietética, hidratação, catabolismo)

**Interpretação Conjunta:**

| Ureia | Creatinina | Ratio Ureia:Crea | Interpretação |
|-------|------------|------------------|---------------|
| ↑↑    | ↑↑         | Normal (10-20:1) | **Insuficiência renal intrínseca** (glomerulonefrite, NTA, CKD) |
| ↑↑    | Normal/↑   | ↑ (>20:1)        | **Azotemia pré-renal** (desidratação, ICC, choque, sangramento GI) |
| ↑↑    | Normal     | ↑ (>20:1)        | **High-protein diet**, desidratação, catabolismo (sepse, corticoides) |
| Normal| ↑↑         | ↓ (<10:1)        | **Azotemia pós-renal** (obstrução ureteral, BPH severa) ou diluição (sobrehidratação) |
| ↓     | Normal     | ↓ (<10:1)        | **Desnutrição proteica**, liver failure, sobrehidratação, SIADH |

**Ratio Ureia:Creatinina Normal:** 10-20:1 (mg/dL)

---

### 2. ESTADO HÍDRICO

**Ureia é MUITO sensível à hidratação:**

**Desidratação:**
- Ureia ↑ (>40 mg/dL)
- Ratio ureia:creatinina ↑ (>20:1)
- Hematócrito ↑, albumina ↑ (hemoconcentração)
- Densidade urinária ↑ (>1.025)

**Sobrehidratação:**
- Ureia ↓ (<20 mg/dL)
- Ratio ureia:creatinina ↓ (<10:1)
- Hematócrito ↓, albumina ↓ (hemodiluição)
- Densidade urinária ↓ (<1.010)

**Functional Medicine Insight:**
- Ureia 22-34 mg/dL sugere **hidratação adequada**
- Ureia >35 mg/dL com creatinina normal = desidratação leve (comum em idosos, atletas)

---

### 3. INGESTÃO PROTEICA

**Ureia ↑ com proteína dietética:**
- **Low-protein (<0.6 g/kg/dia):** Ureia 13-20 mg/dL
- **Balanced (0.8-1.2 g/kg/dia):** Ureia 22-34 mg/dL (optimal)
- **High-protein (>1.5 g/kg/dia):** Ureia 35-50 mg/dL (atletas, carnívoros)
- **Very high (>2.0 g/kg/dia):** Ureia >50 mg/dL (dietas cetogênicas muito ricas em proteína)

**Functional Medicine Application:**
- Ureia 22-34 mg/dL indica **ingestão proteica ideal** para massa muscular + função renal
- Veganos/vegetarianos frequentemente têm ureia 15-25 mg/dL (não patológico se assintomático)

---

### 4. CATABOLISMO PROTEICO

**Causas de Catabolismo ↑ (Ureia ↑):**
- **Sepse, infecção severa:** Quebra muscular para gliconeogênese
- **Trauma, queimadura, cirurgia:** Resposta de stress agudo
- **Corticosteroides:** Prednisona >20 mg/dia causa catabolismo muscular
- **Febre prolongada:** ↑ taxa metabólica
- **Jejum prolongado:** Quebra muscular após depleção glicogênio
- **Câncer avançado:** Caquexia paraneoplásica

**Marcadores Associados:**
- ↑ Ureia (>40 mg/dL)
- ↓ Albumina (<3.5 g/dL)
- ↓ Pré-albumina (<20 mg/dL)
- ↑ PCR (>10 mg/L)
- Perda de massa muscular (DEXA, BIA)

---

### 5. SANGRAMENTO GASTROINTESTINAL

**Upper GI Bleed (esôfago, estômago, duodeno):**
- Sangue = proteína no TGI
- Absorção de aminoácidos → ↑ ureia
- **Ureia ↑↑ (>50 mg/dL)** com creatinina normal
- Ratio ureia:creatinina >20:1
- Hb ↓, melena, hematêmese

**Functional Medicine Insight:**
- Ureia >40 mg/dL aguda (24-48h) SEM desidratação ou high-protein diet → considerar GI bleed oculto

---

### 6. FUNÇÃO HEPÁTICA

**Liver Failure:**
- **↓ Síntese de ureia** (ciclo da ureia hepático)
- Ureia ↓ (<15 mg/dL) apesar de ingestão proteica normal
- Amônia ↑ (não convertida em ureia)
- Encefalopatia hepática (acúmulo amônia cerebral)

**Marcadores Associados:**
- Ureia <15 mg/dL
- Amônia >50 µmol/L (normal <35)
- ALT/AST ↑↑
- Bilirrubina ↑
- Albumina ↓↓ (<2.5 g/dL)
- INR ↑ (>2.0)

**Functional Medicine Insight:**
- Ureia <20 mg/dL + enzimas hepáticas normais → investigar desnutrição proteica, não liver disease

---

## PADRÕES CLÍNICOS INTEGRADOS

### 1. INSUFICIÊNCIA RENAL AGUDA (IRA)

**Clássico:**
```
Ureia: ↑↑↑ (>80 mg/dL)
Creatinina: ↑↑↑ (>3.0 mg/dL)
Ratio Ureia:Creatinina: Normal (10-20:1)
Potássio: ↑ (>5.5 mEq/L)
Bicarbonato: ↓ (<18 mEq/L) - acidose metabólica
Fósforo: ↑ (>5.0 mg/dL)
Volume urinário: ↓ (oligúria <400 mL/dia)
```

**Causas:** NTA (necrose tubular aguda), glomerulonefrite aguda, nefrite intersticial, obstrução bilateral

**Tratamento:** Diálise emergencial se ureia >100 mg/dL + sintomas urêmicos

---

### 2. DOENÇA RENAL CRÔNICA (DRC)

**Estágio 3 (TFG 30-59):**
```
Ureia: 40-60 mg/dL (elevada)
Creatinina: 1.5-3.0 mg/dL
TFG: 30-59 mL/min/1.73m²
Ratio Ureia:Creatinina: Normal
```

**Estágio 4 (TFG 15-29):**
```
Ureia: 60-100 mg/dL (muito elevada)
Creatinina: 3.0-5.0 mg/dL
TFG: 15-29 mL/min/1.73m²
Sintomas: Fadiga, náusea, prurido, edema
```

**Estágio 5 (TFG <15) - ESRD:**
```
Ureia: >100 mg/dL (uremia severa)
Creatinina: >5.0 mg/dL
TFG: <15 mL/min/1.73m²
Indicação: Diálise ou transplante
```

---

### 3. AZOTEMIA PRÉ-RENAL (Desidratação)

**Padrão Típico:**
```
Ureia: ↑↑ (50-80 mg/dL)
Creatinina: Normal ou ↑ leve (0.8-1.5 mg/dL)
Ratio Ureia:Creatinina: ↑↑ (>20:1) ← DIAGNÓSTICO
Densidade urinária: ↑ (>1.025)
Sódio urinário: ↓ (<20 mEq/L) - rim conservando sódio
FENa (fração excretada Na): <1%
```

**Causas:**
- Desidratação (vômito, diarréia, sudorese)
- Hipovolemia (hemorragia, depleção volume)
- ICC (coração não bombeia → ↓ perfusão renal)
- Cirrose (terceiro espaço, hipoalbuminemia)

**Tratamento:** Hidratação IV → ureia normaliza em 24-48h se função renal intacta

---

### 4. DESNUTRIÇÃO PROTEICO-CALÓRICA

**Kwashiorkor / Marasmo:**
```
Ureia: ↓↓ (<13 mg/dL) ← BAIXA!
Albumina: ↓↓ (<2.5 g/dL)
Pré-albumina: ↓↓ (<10 mg/dL)
Transferrina: ↓ (<200 mg/dL)
Linfócitos: ↓ (<1000/µL)
Peso: ↓↓ (perda >20% peso usual)
Edema: + (hipoalbuminemia)
```

**Causas:** Anorexia nervosa, caquexia câncer, má absorção severa, alcoolismo, idosos institucionalizados

**Tratamento:** Suporte nutricional (enteral ou parenteral), reposição proteica gradual

---

### 5. SANGRAMENTO GASTROINTESTINAL ALTO

**Padrão Upper GI Bleed:**
```
Ureia: ↑↑ (50-80 mg/dL) - absorção proteína do sangue
Creatinina: Normal (0.7-1.2 mg/dL)
Ratio Ureia:Creatinina: ↑↑ (>25:1) ← DIAGNÓSTICO
Hemoglobina: ↓↓ (<8.0 g/dL)
Hematócrito: ↓↓ (<25%)
Melena: + (fezes negras)
Hematêmese: ± (vômito sangue)
```

**Causas:** Úlcera péptica, varizes esofágicas (cirrose), Mallory-Weiss, gastrite erosiva

**Functional Medicine Insight:** Ureia ↑ aguda SEM desidratação ou creatinina ↑ → considerar GI bleed oculto

---

### 6. INSUFICIÊNCIA HEPÁTICA

**Liver Failure / Cirrose Descompensada:**
```
Ureia: ↓↓ (<15 mg/dL) ← BAIXA (↓ síntese)
Amônia: ↑↑↑ (>100 µmol/L) - não convertida em ureia
Albumina: ↓↓ (<2.5 g/dL)
Bilirrubina: ↑↑ (>3.0 mg/dL)
INR: ↑↑ (>2.0)
AST/ALT: ↑ ou normal (cirrose "queimada")
Encefalopatia: + (confusão, asterixis, coma)
Ascite: + (terceiro espaço)
```

**Tratamento:** Lactulose (↓ amônia), restrição proteica leve, transplante hepático

---

### 7. DIETA HIGH-PROTEIN (Atletas, Carnívoros)

**Padrão Fisiológico:**
```
Ureia: ↑ (35-50 mg/dL) - não patológico
Creatinina: Normal (0.7-1.2 mg/dL)
Ratio Ureia:Creatinina: Levemente ↑ (15-20:1)
TFG: Normal (>90 mL/min)
Proteína dietética: >1.5 g/kg/dia
Densidade urinária: Normal (1.010-1.020)
Função renal: Preservada
```

**Functional Medicine Approach:**
- Ureia 35-50 mg/dL é ACEITÁVEL em atletas com TFG normal
- Monitorar creatinina, cistatina C, microalbuminúria (marcadores precoces lesão renal)
- Manter hidratação adequada (3-4 L/dia)

---

### 8. DOENÇA CRÍTICA - SEPSE

**Sepse Severa / Choque Séptico:**
```
Ureia: ↑↑ (>60 mg/dL) - catabolismo + IRA
Creatinina: ↑↑ (>2.0 mg/dL)
Lactato: ↑↑ (>4.0 mmol/L)
PCR: ↑↑↑ (>20 mg/dL)
Leucócitos: ↑ ou ↓ (>12k ou <4k)
Plaquetas: ↓ (<100k) - CIVD
Albumina: ↓ (<2.5 g/dL)
Pressão arterial: ↓ (choque)
```

**Mecanismo:** Catabolismo proteico maciço + IRA (NTA) + desidratação relativa

**Tratamento:** Antibióticos broad-spectrum, fluidos IV, vasopressores, diálise se ureia >100 mg/dL

---

## RELAÇÃO UREIA:CREATININA (Ratio)

### Cálculo

```
Ratio Ureia:Creatinina = Ureia (mg/dL) ÷ Creatinina (mg/dL)
```

**Exemplo:**
- Ureia = 30 mg/dL
- Creatinina = 1.0 mg/dL
- Ratio = 30 ÷ 1.0 = **30:1**

### Interpretação

| Ratio | Interpretação | Causas |
|-------|---------------|--------|
| **<10:1** | **Azotemia pós-renal** ou diluição | Obstrução ureteral bilateral, BPH severa com retenção, sobrehidratação, SIADH |
| **10-20:1** | **Normal** | Função renal normal, hidratação adequada, ingestão proteica balanceada |
| **>20:1** | **Azotemia pré-renal** | Desidratação, ICC, cirrose, sangramento GI alto, dieta muito high-protein, catabolismo (sepse) |

### Functional Medicine Application

**Ratio 15-20:1:** Optimal (ureia 22-34 mg/dL, creatinina 1.0-1.8 mg/dL)

**Ratio >25:1:** Investigar:
1. Desidratação (densidade urinária >1.025)
2. Sangramento GI (Hb ↓, melena)
3. Catabolismo (febre, infecção, corticoides)
4. High-protein diet (>2 g/kg/dia)

---

## COMPARAÇÃO: VALORES ORIGINAIS vs CORRIGIDOS

| Aspecto | Original (ERRADO - BUN) | Corrigido (CERTO - UREIA) |
|---------|-------------------------|---------------------------|
| **Parâmetro** | BUN (Blood Urea Nitrogen) | **UREIA** (padrão brasileiro) ✅ |
| **Número de níveis** | 7 níveis | 6 níveis ✅ |
| **Nível 0** | <6 (BUN baixo) | >80 (uremia crítica) ✅ |
| **Nível 1** | >30 (BUN alto) | <13 (desnutrição severa) ✅ |
| **Nível 2** | 6-10 (BUN) | 52-80 (insuf renal moderada) ✅ |
| **Nível 3** | 20-30 (BUN) | 13-21 (baixo-normal) ✅ |
| **Nível 4** | 10-13 (BUN) | 35-51 (alto-normal) ✅ |
| **Nível 5** | 18-20 (BUN) | 22-34 (optimal funcional) ✅ |
| **Nível 6 (EXTRA)** | 13-18 (BUN) | REMOVIDO ✅ |
| **Ordem lógica** | Quebrada | Correta ✅ |
| **Valores** | BUN americano 6-30 | Ureia brasileira 13-80+ ✅ |
| **Labs brasileiros** | NÃO usam | USAM (padrão nacional) ✅ |
| **Conversão clara** | Ausente | BUN = Ureia ÷ 2.14 ✅ |
| **Fonte** | Valores inventados | OptimalDX 2024 + Labs BR ✅ |

---

## CONTROLE DE QUALIDADE - CORREÇÃO UREIA

✅ **Parâmetro correto:** UREIA (não BUN)
✅ **Exactly 6 levels** (era 7, corrigido)
✅ **Valores reais** brasileiros (15-40 mg/dL convencional)
✅ **OptimalDX convertido:** BUN 10-16 → Ureia 22-34 mg/dL
✅ **Ordem lógica:** Piores à esquerda (>80, <13), melhores à direita (22-34)
✅ **Conversão explícita:** BUN = Ureia ÷ 2.14 no cabeçalho
✅ **NO "ou"** em nenhum nível
✅ **NO valores inventados** - tudo baseado em OptimalDX + labs brasileiros
✅ **Padrão nacional:** Brasil usa UREIA, não BUN

---

## REFERÊNCIAS PRINCIPAIS

### Laboratórios Brasileiros
1. **Laboratório Goes** - [Creatinina e Ureia: Valores de Referência](https://laboratoriogoes.com.br/glossario/creatinina-ureia-valores-referencia/)
   - Valores brasileiros por idade e gênero

2. **Lavoisier** - [Ureia: Para que Serve e Níveis de Referência](https://lavoisier.com.br/saude/ureia)
   - Laboratório nacional, valores convencionais

3. **Labs a+ Medicina Diagnóstica** - [Ureia: Função, Referência, Coleta](https://www.labsamais.com.br/noticias/ureia-funcao-referencia-coleta/)
   - Guia prático laboratório brasileiro

4. **Fleury** - [Uréia: O que é, Para que Serve e Valores de Referência](https://www.fleury.com.br/noticias/ureia)
   - Laboratório referência São Paulo

5. **Laboratório Goes** - [Exames de Função Renal: Valores Normais](https://laboratoriogoes.com.br/glossario/exames-funcao-renal-valores-normais/)
   - Painel renal completo brasileiro

### OptimalDX Functional Medicine
6. **OptimalDX** - [Biomarkers of Kidney Function: BUN](https://www.optimaldx.com/research-blog/protein-biomarkers-blood-urea-nitrogen-bun)
   - BUN optimal 10-16 mg/dL

7. **OptimalDX** - [BUN & Urea - What's the Difference?](https://www.optimaldx.com/blog/bun-urea-whats-the-difference)
   - Explicação diferença BUN vs Ureia, fórmula conversão

8. **OptimalDX** - [Functional Health Report 2024](https://sac-nd.com/wp-content/uploads/2024/02/JaneSample-OptimalDX-Practitioner-Feb-03-2024.pdf)
   - BUN ranges: Low <6, Below Optimal 6-10, Optimal 10-16, Above Optimal 16-24, High >24

9. **OptimalDX** - [BUN to Creatinine Ratio](https://www.optimaldx.com/research-blog/renal-biomarkers-bun-creatinine-ratio)
   - Interpretação ratio >20:1 (azotemia pré-renal)

### Clinical References
10. **Medscape** - [Blood Urea Nitrogen Reference Range](https://emedicine.medscape.com/article/2073979-overview)
    - Clinical interpretation BUN

11. **NCBI Bookshelf** - [BUN and Creatinine - Clinical Methods](https://www.ncbi.nlm.nih.gov/books/NBK305/)
    - Peer-reviewed clinical guide

12. **NCBI StatPearls** - [Renal Function Tests](https://www.ncbi.nlm.nih.gov/books/NBK507821/)
    - Comprehensive renal biomarkers

13. **MedlinePlus** - [BUN Blood Test](https://medlineplus.gov/lab-tests/bun-blood-urea-nitrogen/)
    - NIH patient resource

14. **WebMD** - [BUN Test: High vs Low Levels](https://www.webmd.com/a-to-z-guides/blood-urea-nitrogen-test)
    - Clinical interpretation simplified

---

**Arquivo criado:** 19/01/2026
**Sistema:** Plenya EMR - Escore de Saúde Performance e Longevidade
**Status:** ✅ UREIA CORRIGIDA PARA PADRÃO BRASILEIRO
**Parâmetro:** UREIA (mg/dL) - não BUN
**Correções:** BUN → Ureia, 7 → 6 níveis
**Formato:** CSV semicolon-delimited (Portuguese Excel compatible)
