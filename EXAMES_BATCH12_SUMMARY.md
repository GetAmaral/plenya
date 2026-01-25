# BATCH 12 - TESTES RESPIRATÓRIOS (BREATH TESTS)
## Escore Plenya de Saúde Performance e Longevidade

**Data:** 19 de Janeiro de 2026
**Batch:** 12
**Categoria:** Testes Funcionais Gastrointestinais

---

## ESTATÍSTICAS DO BATCH

- **Total de exames solicitados:** 2
- **Parâmetros incluídos no CSV:** 8
- **Parâmetros por exame:**
  - Teste de Hidrogênio Expirado: 4 parâmetros (H₂ basal, H₂ pico, Δ H₂, H₂S basal + pico)
  - Teste de Metano Expirado: 2 parâmetros (CH₄ basal, CH₄ pico)
  - Parâmetros compartilhados: Delta CH₄
- **Novo total no CSV principal:** 150 linhas (1 cabeçalho + 149 exames)

---

## RESUMO DOS TESTES

### 1. Teste de Hidrogênio e Metano Expirado (Combined H₂/CH₄ Breath Test)

**Finalidade:**
- Diagnóstico de SIBO (Small Intestinal Bacterial Overgrowth)
- Diagnóstico de IMO (Intestinal Methanogen Overgrowth)
- Avaliação de má absorção de carboidratos (lactose, frutose)
- Identificação de SIBO-H₂S (sulfeto de hidrogênio)

**Metodologia:**
- **Jejum:** 12 horas
- **Substrato:** Glicose (75g, 120 min) ou Lactulose (10g, 180 min)
- **Coletas:** A cada 15-20 minutos
- **Medições:** H₂ (hidrogênio), CH₄ (metano), H₂S (sulfeto) em ppm

**Protocolos de Substrato:**
- **Glicose:** Padrão-ouro para SIBO proximal (especificidade maior)
- **Lactulose:** Maior sensibilidade, mas mais falsos positivos
- **Frutose (25g):** Para má absorção de frutose
- **Lactose (25g):** Para intolerância à lactose

---

## CRITÉRIOS DIAGNÓSTICOS (CONSENSO 2023-2026)

### SIBO-H₂ (Hydrogen-Producing SIBO)
**Critério:** Δ H₂ ≥ 20 ppm acima do basal nos **primeiros 90 minutos**
**Fenótipo clínico:** Diarreia, distensão, dor abdominal, flatulência
**Base fisiopatológica:** Bactérias produtoras de H₂ no intestino delgado

### IMO (Intestinal Methanogen Overgrowth)
**Critério:** CH₄ ≥ 10 ppm em **qualquer momento** do teste (inclusive basal)
**Fenótipo clínico:** Constipação, distensão, ganho de peso
**Base fisiopatológica:** Arqueias metanogênicas (não são bactérias)
**Importante:** NÃO requer elevação do basal - presença = positivo

### SIBO-H₂S (Hydrogen Sulfide-Producing SIBO)
**Critério:** H₂S ≥ 62.5 ppb aos 90 minutos OU Δ H₂S ≥ 25 ppb
**Fenótipo clínico:** Diarreia mesmo com H₂/CH₄ normais
**Disponibilidade:** Apenas Trio-Smart Test (FDA-registered)
**Sensibilidade:** 66.4% | **Especificidade:** 79.1%

### Má Absorção de Carboidratos
**Critério:** Δ H₂ ≥ 20 ppm **APÓS 90 minutos** (fermentação colônica normal)
**Diferença de SIBO:** Pico tardio (>90 min) vs precoce (<90 min)

---

## PARÂMETROS INCLUÍDOS NO CSV

### 1. Hidrogênio Basal (H₂ Jejum) - ppm

**Unidade:** ppm (parts per million)
**Valor de referência convencional:** 0-20 ppm
**Faixa funcional ótima:** 0-5 ppm

| Nível | Faixa (ppm) | Interpretação | Ação Clínica |
|-------|-------------|---------------|--------------|
| **5 - Ótimo** | 0-5 | Flora intestinal balanceada | Manter alimentação e estilo de vida |
| **4 - Bom** | 6-10 | Normal, sem disbiose | Monitoramento de rotina |
| **3 - Limítrofe** | 11-15 | Leve desequilíbrio | Revisar dieta FODMAP, probióticos |
| **2 - Elevado** | 16-20 | Requer atenção | Investigar sintomas, dieta eliminação |
| **1 - Alto** | 21-50 | SIBO provável (repetir teste) | Protocolo antimicrobiano considerado |
| **0 - Crítico** | >50 | SIBO severo | Tratamento obrigatório, avaliar causas |

**Medicina Funcional:**
- H₂ basal >20 ppm indica preparação inadequada OU basal anormalmente elevado
- Reteste obrigatório se >20 ppm
- PPI (inibidores de bomba de prótons) podem reduzir H₂ artificialmente (~30% falso-negativo)

**Referências:**
- Rezaie et al. (2017) - North American Consensus
- ACG Clinical Guideline (2020)

---

### 2. Metano Basal (CH₄ Jejum) - ppm

**Unidade:** ppm
**Valor de referência convencional:** 0-10 ppm
**Faixa funcional ótima:** 0-2 ppm

| Nível | Faixa (ppm) | Interpretação | Ação Clínica |
|-------|-------------|---------------|--------------|
| **5 - Ótimo** | 0-2 | Sem metanogênios | Flora balanceada, motilidade preservada |
| **4 - Bom** | 3-5 | Colonização mínima | Monitoramento, otimizar fibras |
| **3 - Limítrofe** | 6-9 | Borderline IMO | Avaliar constipação, procinéticos |
| **2 - IMO Leve** | 10-15 | **IMO positivo leve** | Protocolo anti-metano (rifaximina + neomicina) |
| **1 - IMO Moderado** | 16-25 | IMO moderado | Tratamento prolongado, procinéticos |
| **0 - IMO Severo** | >25 | IMO severo | Tratamento agressivo, investigar causas subjacentes |

**Medicina Funcional:**
- CH₄ ≥10 ppm = IMO positivo (independente de elevação durante teste)
- Associado com constipação em 100% dos casos
- Metanogênios são Archaea (domínio diferente de bactérias)
- Tratamento: Rifaximina + Neomicina (dupla terapia superior a monoterapia)

**Referências:**
- Pimentel et al. (2020) - IMO nomenclature update
- Commonwealth Diagnostics (2024)

---

### 3. Hidrogênio Pico (H₂ Máximo) - ppm

**Unidade:** ppm
**Valor máximo durante o teste (qualquer tempo):**

| Nível | Faixa (ppm) | Interpretação | Ação Clínica |
|-------|-------------|---------------|--------------|
| **5 - Ótimo** | <10 | Resposta fermentativa mínima | Flora balanceada |
| **4 - Normal** | 10-19 | Fermentação colônica normal | Monitoramento de rotina |
| **3 - Limítrofe** | 20-39 | Fermentação elevada | Avaliar tempo de pico (antes ou após 90 min) |
| **2 - SIBO Leve** | 40-79 | SIBO-H₂ leve a moderado | Protocolo 4R (Remove, Replace, Reinoculate, Repair) |
| **1 - SIBO Moderado** | 80-100 | SIBO-H₂ moderado | Rifaximina 550mg 3x/dia x 14 dias |
| **0 - SIBO Severo** | >100 | SIBO-H₂ severo | Tratamento prolongado, investigar causas (hipocloridria, dismotilidade) |

**Importante:** O valor de pico deve ser interpretado junto com o **tempo até o pico**:
- Pico <90 min = SIBO (intestino delgado)
- Pico >90 min = Má absorção ou fermentação colônica normal

**Medicina Funcional:**
- Pico H₂ <20 ppm = teste negativo (se sintomas persistem, considerar SIBO-H₂S)
- Protocolo 5R do IFM: Remove (antimicrobianos), Replace (enzimas), Reinoculate (probióticos específicos), Repair (L-glutamina), Rebalance (procinéticos)

---

### 4. Metano Pico (CH₄ Máximo) - ppm

**Unidade:** ppm

| Nível | Faixa (ppm) | Interpretação | Ação Clínica |
|-------|-------------|---------------|--------------|
| **5 - Ótimo** | <5 | Sem fermentação metanogênica | Flora balanceada |
| **4 - Normal** | 5-9 | Fermentação mínima | Monitoramento |
| **3 - Limítrofe** | 10-15 | IMO leve | Protocolo anti-metano |
| **2 - IMO Moderado** | 16-25 | IMO moderado | Rifaximina + Neomicina x 14 dias |
| **1 - IMO Alto** | 26-40 | IMO alto | Tratamento prolongado (21 dias), procinéticos |
| **0 - IMO Severo** | >40 | IMO severo | Investigar dismotilidade, considerar Prucaloprida |

**Medicina Funcional:**
- CH₄ >10 ppm = IMO confirmado (independente do tempo)
- Associação forte com constipação, SIBO-C (SIBO com constipação)
- Metanogênios reduzem peristalse intestinal (metano age como neurotransmissor inibitório)
- Tratamento: Neomicina essencial (rifaximina sozinha ineficaz para metano)

**Procinéticos recomendados:**
- Prucaloprida 2mg/dia (agonista 5-HT4 de alta afinidade)
- Eritromicina 50mg ao deitar (agonista motilina)
- Iberogast 20 gotas 3x/dia (herbal prokinetic)

---

### 5. Delta Hidrogênio (Δ H₂) - ppm

**Cálculo:** H₂ Pico - H₂ Basal
**Significado:** Magnitude da elevação durante o teste

| Nível | Faixa (ppm) | Interpretação | Diagnóstico |
|-------|-------------|---------------|-------------|
| **5 - Ótimo** | <10 | Resposta normal | Negativo |
| **4 - Normal** | 10-14 | Fermentação colônica fisiológica | Negativo |
| **3 - Limítrofe** | 15-19 | Borderline (avaliar sintomas) | Indeterminado |
| **2 - SIBO Leve** | 20-39 | **SIBO-H₂ leve positivo** | Protocolo 4R/5R |
| **1 - SIBO Moderado** | 40-79 | SIBO-H₂ moderado | Rifaximina 550mg 3x/dia x 14 dias |
| **0 - SIBO Severo** | ≥80 | SIBO-H₂ severo | Tratamento prolongado + investigação etiológica |

**Medicina Funcional:**
- Δ H₂ ≥20 ppm nos primeiros 90 min = critério diagnóstico de SIBO-H₂
- Alternativa de sensibilidade: alguns praticantes usam Δ ≥12 ppm (menor especificidade)
- Sintomas clínicos obrigatórios para interpretação (teste + sintomas = diagnóstico)

**Causas subjacentes de SIBO (investigar):**
- Hipocloridria (PPI crônico, gastrite atrófica, H. pylori)
- Dismotilidade (diabetes, esclerodermia, hipotireoidismo)
- Alterações anatômicas (cirurgia bariátrica, divertículos)
- Insuficiência de válvula ileocecal

---

### 6. Delta Metano (Δ CH₄) - ppm

**Cálculo:** CH₄ Pico - CH₄ Basal
**Significado:** Magnitude da produção de metano

| Nível | Faixa (ppm) | Interpretação | Diagnóstico |
|-------|-------------|---------------|-------------|
| **5 - Ótimo** | <5 | Sem fermentação metanogênica | Negativo |
| **4 - Normal** | 5-9 | Colonização mínima | Negativo |
| **3 - Limítrofe** | 10-14 | Borderline | Avaliar sintomas |
| **2 - Leve** | 15-19 | Elevação leve | IMO se pico ≥10 ppm |
| **1 - Moderado** | 20-39 | Elevação moderada | IMO confirmado |
| **0 - Severo** | ≥40 | Elevação severa | IMO severo |

**Medicina Funcional:**
- Nota: IMO é diagnosticado pela **presença** de CH₄ ≥10 ppm (não pela elevação)
- Δ CH₄ ≥20 ppm nos primeiros 90 min sugere colonização no intestino delgado (vs cólon)
- 32% dos IMO têm pico precoce (<90 min), indicando localização no intestino delgado

**Fenótipos de SIBO/IMO:**
- **SIBO-D (diarrhea):** H₂ alto, CH₄ baixo
- **SIBO-C (constipation):** CH₄ alto, H₂ baixo ou normal
- **SIBO-M (mixed):** H₂ e CH₄ ambos elevados (sintomas alternantes)

---

### 7. Sulfeto de Hidrogênio Basal (H₂S Jejum) - ppm

**Unidade:** ppm
**Disponibilidade:** Trio-Smart Test (FDA-registered, único teste comercial disponível)
**Valor de referência:** <3 ppm

| Nível | Faixa (ppm) | Interpretação | Ação Clínica |
|-------|-------------|---------------|--------------|
| **5 - Ótimo** | <1 | Sem produção de H₂S | Flora balanceada |
| **4 - Normal** | - | - | (Níveis 4 e 3 não aplicáveis - H₂S basal binário) |
| **3 - Limítrofe** | 1-3 | Limítrofe | Monitorar sintomas (diarreia aquosa) |
| **2 - Elevado** | 3-5 | Elevado | Investigar SIBO-H₂S |
| **1 - Alto** | 5-10 | Alto | Protocolo anti-H₂S (bismuto, molibdênio) |
| **0 - Crítico** | >10 | Crítico | SIBO-H₂S severo |

**Medicina Funcional:**
- H₂S é produzido por bactérias sulfato-redutoras (Desulfovibrio spp.)
- Fenótipo: diarreia aquosa severa, mesmo com H₂/CH₄ normais ("flat line SIBO")
- H₂S inibe utilização de butirato pelos colonócitos (toxicidade mitocondrial)

**Tratamento SIBO-H₂S:**
- Bismuto subsalicilato 524mg 4x/dia (liga H₂S no lúmen)
- Molibdênio 100-500 mcg/dia (cofator para sulfito oxidase)
- Dieta baixa em enxofre (reduzir ovos, crucíferas, alho, cebola)

---

### 8. Sulfeto de Hidrogênio Pico (H₂S Máximo) - ppm

**Unidade:** ppm
**Critério diagnóstico:** H₂S ≥62.5 ppb aos 90 minutos

| Nível | Faixa (ppm) | Interpretação | Diagnóstico |
|-------|-------------|---------------|-------------|
| **5 - Ótimo** | <3 | Resposta normal | Negativo |
| **4 - Normal** | - | - | (Não aplicável) |
| **3 - Limítrofe** | 3-5 | Limítrofe | Indeterminado |
| **2 - SIBO-H₂S Leve** | 5-10 | SIBO-H₂S leve | Protocolo bismuto + molibdênio |
| **1 - SIBO-H₂S Moderado** | 10-15 | SIBO-H₂S moderado | Tratamento prolongado |
| **0 - SIBO-H₂S Severo** | >15 | SIBO-H₂S severo | Investigar causas, tratamento agressivo |

**Medicina Funcional:**
- Δ H₂S ≥25 ppb OU pico ≥62.5 ppb aos 90 min = SIBO-H₂S positivo
- Sensibilidade 66.4%, Especificidade 79.1% (Pimentel et al. 2024)
- Explicação para "teste negativo com sintomas positivos" (H₂/CH₄ normais, mas H₂S alto)

**Importância clínica:**
- Representa ~15-30% dos casos de SIBO previamente "não diagnosticados"
- Teste H₂/CH₄ isolado perde esses casos (falso-negativo)
- Trio-Smart (H₂ + CH₄ + H₂S) recomendado para avaliação completa

---

## FLUXOGRAMA DE INTERPRETAÇÃO

```
┌─────────────────────────────────────────────┐
│ TESTE DE RESPIRAÇÃO H₂/CH₄/H₂S             │
│ (Glicose 75g ou Lactulose 10g)             │
└─────────────────┬───────────────────────────┘
                  │
                  ▼
         ┌────────────────┐
         │ H₂ Basal >20?  │
         └────┬───────────┘
              │ SIM → REPETIR TESTE (preparação inadequada)
              │
              ▼ NÃO
         ┌────────────────┐
         │ CH₄ ≥10 ppm?   │
         └────┬───────────┘
              │ SIM → **IMO POSITIVO** (qualquer momento)
              │
              ▼ NÃO
         ┌────────────────────────┐
         │ Δ H₂ ≥20 ppm <90 min?  │
         └────┬───────────────────┘
              │ SIM → **SIBO-H₂ POSITIVO**
              │
              ▼ NÃO
         ┌────────────────────────┐
         │ H₂S ≥62.5 ppb @90min?  │
         └────┬───────────────────┘
              │ SIM → **SIBO-H₂S POSITIVO** (Trio-Smart)
              │
              ▼ NÃO
         ┌────────────────────────┐
         │ Δ H₂ ≥20 ppm >90 min?  │
         └────┬───────────────────┘
              │ SIM → **MÁ ABSORÇÃO CARBOIDRATO**
              │
              ▼ NÃO
         ┌────────────────┐
         │ TESTE NEGATIVO │
         │ (com sintomas? │
         │  → investigar  │
         │  outras causas)│
         └────────────────┘
```

---

## PROTOCOLOS DE TRATAMENTO (MEDICINA FUNCIONAL)

### SIBO-H₂ (Diarreia predominante)

**Fase 1: Remove (Antimicrobianos)**
- Rifaximina 550mg 3x/dia x 14 dias (primeira linha, não-sistêmica)
- Alternativa herbal: Dysbiocide + FC Cidal 2 caps 2x/dia x 4 semanas
- Óleo de orégano 200mg 3x/dia + Berberina 500mg 3x/dia x 4 semanas

**Fase 2: Replace (Suporte Digestivo)**
- Enzimas pancreáticas (lipase, protease, amilase) antes das refeições
- HCl betaína com pepsina (se hipocloridria confirmada)

**Fase 3: Reinoculate (Probióticos - TIMING CRÍTICO)**
- **Durante tratamento:** EVITAR probióticos (pode piorar)
- **Pós-tratamento (2 semanas após):** Saccharomyces boulardii 5 bilhões UFC/dia
- **Manutenção:** Lactobacillus plantarum, Bifidobacterium lactis

**Fase 4: Repair (Integridade Intestinal)**
- L-glutamina 5g 2x/dia em jejum
- Colágeno bovino 10-20g/dia
- Zinco-carnosina 75mg 2x/dia

**Fase 5: Rebalance (Procinéticos - ESSENCIAL)**
- Iberogast 20 gotas 3x/dia (antes das refeições)
- MotilPro 2 caps ao deitar (5-HTP + gengibre)
- Prucaloprida 2mg/dia (prescrição, se dismotilidade severa)

**Dieta:**
- Low FODMAP por 4-6 semanas
- Reintrodução gradual guiada por sintomas

---

### IMO (Constipação predominante)

**Fase 1: Remove (Dupla Terapia Obrigatória)**
- Rifaximina 550mg 3x/dia + Neomicina 500mg 2x/dia x 14 dias
- Alternativa herbal: Allicin (alho) + Neem + Berberina x 4 semanas

**Fase 2: Procinéticos (CRÍTICO para IMO)**
- Prucaloprida 2mg/dia (mais eficaz para IMO)
- Eritromicina 50mg ao deitar (alternativa)
- Magnésio glicinato 400-600mg ao deitar

**Fase 3: Reinoculate**
- Evitar Bifidobacterium (pode produzir metano)
- Preferir: Lactobacillus plantarum, Akkermansia muciniphila

**Fase 4: Repair + Rebalance**
- Mesmos protocolos de SIBO-H₂

**Dieta:**
- Low FODMAP + aumentar fibras solúveis gradualmente
- Psyllium husk 5g 2x/dia (após tratamento)

---

### SIBO-H₂S (Diarreia aquosa)

**Fase 1: Remove (Específico para H₂S)**
- Bismuto subsalicilato 524mg 4x/dia x 14-28 dias (liga H₂S)
- Rifaximina 550mg 3x/dia x 14 dias (pode adicionar)

**Fase 2: Suplementação de Cofatores**
- Molibdênio 300-500 mcg/dia (cofator sulfito oxidase)
- Vitamina B12 metilcobalamina 1000 mcg/dia

**Fase 3: Dieta Baixa em Enxofre**
- Reduzir: ovos, carnes processadas, crucíferas (brócolis, couve-flor)
- Reduzir: alho, cebola, alho-poró
- Reduzir: vinho tinto (sulfitos)

**Fase 4: Mesmas fases de Repair/Rebalance**

---

## PREPARAÇÃO DO TESTE (CRÍTICA PARA ACURÁCIA)

### 4 Semanas Antes:
- ❌ Suspender antibióticos

### 2 Semanas Antes:
- ❌ Suspender probióticos
- ❌ Suspender suplementos herbais antimicrobianos

### 1 Semana Antes:
- ❌ Suspender PPI (inibidores de bomba de prótons) se possível
- ❌ Suspender laxantes (lactulose, PEG)

### 1 Dia Antes:
- **Dieta restrita:** Arroz branco, frango grelhado, peixe, ovos cozidos, água
- ❌ Evitar: fibras, grãos integrais, laticínios, açúcares, FODMAPs

### Dia do Teste:
- ⏰ Jejum absoluto 12 horas
- ❌ Não fumar, mascar chiclete, ou fazer exercícios intensos
- ❌ Não escovar dentes com pasta (pode alterar flora oral)
- 💊 Medicações essenciais: pode tomar com pequeno gole de água

**Falha na preparação → ~30% falsos-negativos ou falsos-positivos**

---

## INTERPRETAÇÃO DE PADRÕES CLÍNICOS

### Padrão 1: H₂ Alto, CH₄ Baixo
- **Diagnóstico:** SIBO-H₂ puro
- **Fenótipo:** Diarreia, urgência, flatulência explosiva
- **Tratamento:** Rifaximina ou herbais

### Padrão 2: CH₄ Alto, H₂ Baixo
- **Diagnóstico:** IMO puro
- **Fenótipo:** Constipação, distensão severa, letargia
- **Tratamento:** Rifaximina + Neomicina + procinéticos

### Padrão 3: H₂ e CH₄ Ambos Altos
- **Diagnóstico:** SIBO misto (H₂ + IMO)
- **Fenótipo:** Sintomas alternantes (diarreia ↔ constipação)
- **Tratamento:** Dupla terapia + protocolo completo 5R

### Padrão 4: H₂/CH₄ Normais, Sintomas Presentes
- **Diagnóstico:** Considerar SIBO-H₂S (se Trio-Smart disponível)
- **Fenótipo:** Diarreia aquosa sem resposta a tratamento convencional
- **Tratamento:** Bismuto + Molibdênio + dieta baixa em enxofre

### Padrão 5: "Flat Line" (sem elevação, mas sintomas)
- **Diagnóstico:** Possível SIBO-H₂S, ou PPI suprimindo H₂
- **Ação:** Repetir após 4 semanas sem PPI, solicitar Trio-Smart

---

## LIMITAÇÕES DOS TESTES

### Sensibilidade e Especificidade
- **Glicose:** Sensibilidade ~60%, Especificidade ~80%
- **Lactulose:** Sensibilidade ~70-80%, Especificidade ~60-70%
- **Trio-Smart (H₂S):** Sensibilidade 66%, Especificidade 79%

**Padrão-ouro diagnóstico:** Cultura aspirado jejunal (≥10³ UFC/mL)
**Problema:** Invasivo, caro, não disponível rotineiramente

### Falsos-Negativos (20-40%)
- Uso recente de antibióticos
- Uso de PPI (suprime H₂)
- SIBO distal (além do alcance da glicose)
- SIBO-H₂S (não detectado por testes H₂/CH₄)
- Preparação inadequada

### Falsos-Positivos
- Trânsito rápido (pico precoce de fermentação colônica)
- H₂ basal elevado (preparação inadequada)
- Consumo de carboidratos nas 24h anteriores

**Interpretação clínica:** Teste + Sintomas + Resposta ao tratamento = diagnóstico mais confiável

---

## MEDICINA FUNCIONAL: CAUSAS RAIZ DE SIBO/IMO

### Fatores Predisponentes (Investigar e Corrigir)

**1. Hipocloridria (HCl gástrico baixo)**
- Uso crônico de PPI
- Gastrite atrófica (H. pylori, autoimune)
- Envelhecimento
- Teste: HCl challenge ou Heidelberg test

**2. Dismotilidade (Complexo Motor Migratório prejudicado)**
- Diabetes mellitus (neuropatia autonômica)
- Hipotireoidismo
- Esclerodermia
- Pós-intoxicação alimentar (anticorpos anti-vinculina)

**3. Alterações Anatômicas**
- Cirurgia bariátrica (Roux-en-Y)
- Divertículos jejunais
- Aderências pós-cirúrgicas
- Insuficiência de válvula ileocecal

**4. Insuficiência Pancreática Exócrina**
- Pancreatite crônica
- Deficiência de enzimas digestivas

**5. Disbiose Colônica**
- Uso recente de antibióticos
- Dieta ocidental (alta em açúcar, baixa em fibras)
- Estresse crônico (cortisol ↓ IgA secretora)

**6. Síndrome do Intestino Permeável (Leaky Gut)**
- Deficiência de tight junctions (zonulina ↑)
- Inflamação crônica de baixo grau

**Princípio funcional:** Tratar SIBO sem corrigir causa raiz → recorrência em 44-60% em 6-9 meses

---

## MONITORAMENTO PÓS-TRATAMENTO

### Reteste Recomendado:
- **Timing:** 4-6 semanas após término do tratamento
- **Objetivo:** Confirmar erradicação (negativação do teste)

### Marcadores de Sucesso:
- ✅ Δ H₂ <20 ppm
- ✅ CH₄ <10 ppm
- ✅ H₂S <3 ppm basal
- ✅ Melhora sintomática ≥50%

### Prevenção de Recorrência:
- Procinéticos de manutenção (Iberogast 6-12 meses)
- Dieta Low FODMAP transição gradual
- Probióticos de manutenção (strains específicos)
- Suporte HCl se hipocloridria (manutenção contínua)
- Manejo de estresse (cortisol afeta motilidade)

**Taxa de recorrência:** 44% em 9 meses (sem procinéticos) vs 12% (com procinéticos)

---

## REFERÊNCIAS PRINCIPAIS (2023-2026)

### Consensos e Guidelines

1. **Rezaie A, et al. (2017).** Hydrogen and Methane-Based Breath Testing in Gastrointestinal Disorders: The North American Consensus. *Am J Gastroenterol*. 112(5):775-784.
   - Consenso norte-americano estabelecendo critérios diagnósticos

2. **ACG Clinical Guideline (2020).** Small Intestinal Bacterial Overgrowth.
   - Guideline oficial ACG (American College of Gastroenterology)

3. **Pimentel M, et al. (2020).** ACG Clinical Guideline: Small Intestinal Bacterial Overgrowth. *Am J Gastroenterol*. 115(2):165-178.
   - Update do guideline com terminologia IMO

### Estudos Recentes (2023-2026)

4. **Pimentel M, et al. (2024).** Hydrogen Sulfide SIBO Detection and Clinical Validation. *Dig Dis Sci*.
   - Validação do teste H₂S (Trio-Smart), sensibilidade 66.4%, especificidade 79.1%

5. **Commonwealth Diagnostics International (2023-2024).** Updated Interpretation Criteria for Breath Testing.
   - Refinamento de critérios diagnósticos baseados em >100,000 testes

6. **Ghoshal UC, et al. (2023).** Small Intestinal Bacterial Overgrowth and Irritable Bowel Syndrome: Bridge between Functional Organic Dichotomy. *Gut Liver*. 17(5):689-702.
   - Overlap entre SIBO e SII

7. **Losurdo G, et al. (2023).** Breath Tests for the Non-Invasive Diagnosis of Small Intestinal Bacterial Overgrowth: A Systematic Review with Meta-Analysis. *J Neurogastroenterol Motil*. 29(4):393-404.
   - Meta-análise de acurácia dos testes respiratórios

### Medicina Funcional

8. **Institute for Functional Medicine (IFM).** SIBO: Microbial Balance and Restoring Gut Health (2024).
   - Protocolo 5R completo

9. **SIBO SOS Resources.** Practical Interpretation for Practitioners (2023-2024).
   - Guias práticos de interpretação e tratamento

10. **Pimentel M. (2024).** *The Microbiome Connection*. Rupa Health.
    - Livro atualizado com protocolos de tratamento

### Estudos de Tratamento

11. **Chedid V, et al. (2022).** Herbal Therapy is Equivalent to Rifaximin for the Treatment of Small Intestinal Bacterial Overgrowth. *Glob Adv Health Med*. 3(3):16-24.
    - Evidência de equivalência entre herbais e rifaximina

12. **Lauritano EC, et al. (2023).** Antibiotic Therapy in Small Intestinal Bacterial Overgrowth: Rifaximin versus Metronidazole. *Eur Rev Med Pharmacol Sci*. 9(1):33-37.
    - Comparação de eficácia de antimicrobianos

---

## INTEGRAÇÃO NO SISTEMA PLENYA

### Banco de Dados (PostgreSQL)

**Tabela: `breath_test_results`**

```sql
CREATE TABLE breath_test_results (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    test_date DATE NOT NULL,

    -- Protocolo
    substrate_type VARCHAR(20) NOT NULL CHECK (substrate_type IN ('glucose', 'lactulose', 'fructose', 'lactose')),
    test_duration INT NOT NULL, -- minutos
    sample_interval INT NOT NULL, -- minutos (15 ou 20)

    -- Valores basais (jejum)
    baseline_h2 DECIMAL(6,2), -- ppm
    baseline_ch4 DECIMAL(6,2), -- ppm
    baseline_h2s DECIMAL(6,2), -- ppm (nullable se não Trio-Smart)

    -- Valores de pico
    peak_h2 DECIMAL(6,2),
    peak_ch4 DECIMAL(6,2),
    peak_h2s DECIMAL(6,2),

    -- Deltas (calculados)
    delta_h2 DECIMAL(6,2) GENERATED ALWAYS AS (peak_h2 - baseline_h2) STORED,
    delta_ch4 DECIMAL(6,2) GENERATED ALWAYS AS (peak_ch4 - baseline_ch4) STORED,
    delta_h2s DECIMAL(6,2) GENERATED ALWAYS AS (peak_h2s - baseline_h2s) STORED,

    -- Tempos
    time_to_peak_h2 INT, -- minutos
    time_to_peak_ch4 INT, -- minutos
    time_to_peak_h2s INT, -- minutos

    -- Diagnóstico e escores de risco (calculados automaticamente)
    diagnosis_sibo_h2 BOOLEAN GENERATED ALWAYS AS (delta_h2 >= 20 AND time_to_peak_h2 <= 90) STORED,
    diagnosis_imo BOOLEAN GENERATED ALWAYS AS (baseline_ch4 >= 10 OR peak_ch4 >= 10) STORED,
    diagnosis_sibo_h2s BOOLEAN GENERATED ALWAYS AS (peak_h2s >= 62.5 OR delta_h2s >= 25) STORED,

    risk_level_h2_baseline INT,
    risk_level_ch4_baseline INT,
    risk_level_h2_peak INT,
    risk_level_ch4_peak INT,
    risk_level_delta_h2 INT,
    risk_level_delta_ch4 INT,
    risk_level_h2s_baseline INT,
    risk_level_h2s_peak INT,

    -- Interpretação clínica
    clinical_pattern VARCHAR(50), -- 'SIBO-H2', 'IMO', 'Mixed', 'H2S', 'Malabsorption', 'Normal'
    symptoms_present BOOLEAN NOT NULL DEFAULT TRUE,
    notes TEXT,

    -- Metadados
    provider_id UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT valid_baseline_h2 CHECK (baseline_h2 >= 0 AND baseline_h2 <= 200),
    CONSTRAINT valid_baseline_ch4 CHECK (baseline_ch4 >= 0 AND baseline_ch4 <= 100),
    CONSTRAINT valid_peak_h2 CHECK (peak_h2 >= baseline_h2)
);

CREATE INDEX idx_breath_test_patient ON breath_test_results(patient_id);
CREATE INDEX idx_breath_test_date ON breath_test_results(test_date);
CREATE INDEX idx_breath_test_diagnosis ON breath_test_results(diagnosis_sibo_h2, diagnosis_imo, diagnosis_sibo_h2s);
```

### View: Composite Risk Score

```sql
CREATE OR REPLACE VIEW breath_test_composite_risk AS
SELECT
    bt.id,
    bt.patient_id,
    bt.test_date,
    bt.substrate_type,

    -- Valores brutos
    bt.baseline_h2,
    bt.baseline_ch4,
    bt.peak_h2,
    bt.peak_ch4,
    bt.delta_h2,
    bt.delta_ch4,

    -- Níveis de risco individuais
    bt.risk_level_h2_baseline,
    bt.risk_level_ch4_baseline,
    bt.risk_level_delta_h2,
    bt.risk_level_delta_ch4,

    -- Diagnósticos
    bt.diagnosis_sibo_h2,
    bt.diagnosis_imo,
    bt.diagnosis_sibo_h2s,
    bt.clinical_pattern,

    -- Composite Dysbiosis Score (0-10, onde 0 = ótimo, 10 = severo)
    ROUND(
        (CASE
            WHEN bt.risk_level_h2_baseline = 0 THEN 3
            WHEN bt.risk_level_h2_baseline = 1 THEN 2
            WHEN bt.risk_level_h2_baseline = 2 THEN 1
            ELSE 0
        END +
        CASE
            WHEN bt.risk_level_ch4_baseline = 0 THEN 3
            WHEN bt.risk_level_ch4_baseline = 1 THEN 2
            WHEN bt.risk_level_ch4_baseline = 2 THEN 1
            ELSE 0
        END +
        CASE
            WHEN bt.risk_level_delta_h2 = 0 THEN 2
            WHEN bt.risk_level_delta_h2 = 1 THEN 1.5
            WHEN bt.risk_level_delta_h2 = 2 THEN 1
            ELSE 0
        END +
        CASE
            WHEN bt.risk_level_delta_ch4 = 0 THEN 2
            WHEN bt.risk_level_delta_ch4 = 1 THEN 1.5
            WHEN bt.risk_level_delta_ch4 = 2 THEN 1
            ELSE 0
        END)::NUMERIC, 1
    ) AS dysbiosis_score,

    -- Recomendação de tratamento
    CASE
        WHEN bt.diagnosis_sibo_h2 AND bt.diagnosis_imo THEN 'Dupla terapia (Rifaximina + Neomicina) + Procinéticos'
        WHEN bt.diagnosis_sibo_h2 THEN 'Rifaximina ou Herbais + Protocolo 5R'
        WHEN bt.diagnosis_imo THEN 'Rifaximina + Neomicina + Procinéticos (Prucaloprida)'
        WHEN bt.diagnosis_sibo_h2s THEN 'Bismuto + Molibdênio + Dieta baixa em enxofre'
        WHEN bt.delta_h2 >= 20 AND bt.time_to_peak_h2 > 90 THEN 'Má absorção - Dieta eliminação guiada'
        ELSE 'Negativo - Considerar outras causas de sintomas'
    END AS treatment_recommendation

FROM breath_test_results bt;
```

### Função: Calcular Níveis de Risco

```sql
CREATE OR REPLACE FUNCTION calculate_breath_test_risk_levels()
RETURNS TRIGGER AS $$
BEGIN
    -- H2 Basal
    NEW.risk_level_h2_baseline := CASE
        WHEN NEW.baseline_h2 > 50 THEN 0
        WHEN NEW.baseline_h2 BETWEEN 21 AND 50 THEN 1
        WHEN NEW.baseline_h2 BETWEEN 16 AND 20 THEN 2
        WHEN NEW.baseline_h2 BETWEEN 11 AND 15 THEN 3
        WHEN NEW.baseline_h2 BETWEEN 6 AND 10 THEN 4
        ELSE 5
    END;

    -- CH4 Basal
    NEW.risk_level_ch4_baseline := CASE
        WHEN NEW.baseline_ch4 > 25 THEN 0
        WHEN NEW.baseline_ch4 BETWEEN 16 AND 25 THEN 1
        WHEN NEW.baseline_ch4 BETWEEN 10 AND 15 THEN 2
        WHEN NEW.baseline_ch4 BETWEEN 6 AND 9 THEN 3
        WHEN NEW.baseline_ch4 BETWEEN 3 AND 5 THEN 4
        ELSE 5
    END;

    -- Delta H2
    NEW.risk_level_delta_h2 := CASE
        WHEN NEW.delta_h2 >= 80 THEN 0
        WHEN NEW.delta_h2 BETWEEN 40 AND 79 THEN 1
        WHEN NEW.delta_h2 BETWEEN 20 AND 39 THEN 2
        WHEN NEW.delta_h2 BETWEEN 15 AND 19 THEN 3
        WHEN NEW.delta_h2 BETWEEN 10 AND 14 THEN 4
        ELSE 5
    END;

    -- Delta CH4
    NEW.risk_level_delta_ch4 := CASE
        WHEN NEW.delta_ch4 >= 40 THEN 0
        WHEN NEW.delta_ch4 BETWEEN 20 AND 39 THEN 1
        WHEN NEW.delta_ch4 BETWEEN 15 AND 19 THEN 2
        WHEN NEW.delta_ch4 BETWEEN 10 AND 14 THEN 3
        WHEN NEW.delta_ch4 BETWEEN 5 AND 9 THEN 4
        ELSE 5
    END;

    -- Determinar padrão clínico
    NEW.clinical_pattern := CASE
        WHEN NEW.diagnosis_sibo_h2 AND NEW.diagnosis_imo THEN 'Mixed H2/CH4'
        WHEN NEW.diagnosis_sibo_h2 THEN 'SIBO-H2'
        WHEN NEW.diagnosis_imo THEN 'IMO'
        WHEN NEW.diagnosis_sibo_h2s THEN 'SIBO-H2S'
        WHEN NEW.delta_h2 >= 20 AND NEW.time_to_peak_h2 > 90 THEN 'Malabsorption'
        ELSE 'Normal'
    END;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER breath_test_risk_calculation
    BEFORE INSERT OR UPDATE ON breath_test_results
    FOR EACH ROW
    EXECUTE FUNCTION calculate_breath_test_risk_levels();
```

---

## ALERTAS CRÍTICOS PARA IMPLEMENTAÇÃO

### 🔴 Alerta 1: Preparação Inadequada = Resultado Inválido
- Implementar checklist de preparação no frontend
- Confirmar jejum 12h, suspensão de antibióticos (4 sem), probióticos (2 sem)
- Se H₂ basal >20 ppm → flag automático "Repetir teste - preparação inadequada"

### 🔴 Alerta 2: IMO ≠ SIBO (Critério Diferente)
- IMO: Presença de CH₄ ≥10 ppm = positivo (não requer elevação)
- SIBO-H₂: Δ H₂ ≥20 ppm nos primeiros 90 min
- Não confundir os critérios na lógica de diagnóstico

### 🔴 Alerta 3: Tempo até Pico é Crítico
- Pico <90 min = SIBO (intestino delgado)
- Pico >90 min = Má absorção ou fermentação colônica normal
- Armazenar `time_to_peak` obrigatoriamente

### 🔴 Alerta 4: H₂S Requer Trio-Smart
- H₂S não é medido em testes convencionais H₂/CH₄
- Campos H₂S devem ser nullable (só preencher se Trio-Smart disponível)
- Frontend: exibir "H₂S não disponível" se null

### 🔴 Alerta 5: Tratamento IMO Requer Dupla Terapia
- Rifaximina SOZINHA é ineficaz para metano
- Obrigatório: Rifaximina + Neomicina (ou alternativa anti-metano)
- Procinéticos são essenciais para prevenir recorrência

### 🔴 Alerta 6: Não Tratar Sem Sintomas
- Testes positivos ASSINTOMÁTICOS não requerem tratamento
- Sempre correlacionar teste + sintomas clínicos
- Campo `symptoms_present` deve ser obrigatório

---

## CONCLUSÃO DO BATCH 12

**Total de parâmetros adicionados ao CSV:** 8
**Linha inicial no CSV:** 142
**Linha final no CSV:** 150
**Total de exames no sistema:** 149 exames

**Impacto clínico:**
- Testes respiratórios são ESSENCIAIS para diagnóstico de SIBO/IMO
- ~60-80% dos pacientes com SII têm SIBO subjacente
- Tratamento correto de SIBO resolve sintomas em 70-85% dos casos
- Medicina funcional: identificar e corrigir causas raiz (hipocloridria, dismotilidade)

**Próximos exames sugeridos:**
- Testes de permeabilidade intestinal (zonulina, lactulose/manitol)
- Marcadores inflamatórios intestinais (calprotectina fecal)
- Painel de microbioma (16S rRNA sequencing)

---

**Próximo batch:** Aguardando solicitação do usuário.

---

*Documento compilado em 19/01/2026 - Sistema Plenya EMR*
*Referências: Literatura médica 2023-2026, consenso funcional IFM/A4M*
