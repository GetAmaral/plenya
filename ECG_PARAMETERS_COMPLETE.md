# Parâmetros Completos de ECG - Estratificação de Risco Cardiovascular

**Data:** Janeiro 2026
**Parâmetros Adicionados:** 6 novos (PR, QRS, Sokolow-Lyon, Cornell homens/mulheres, Eixo)
**Total ECG:** 9 parâmetros (FC, QTc homens/mulheres, PR, QRS, Sokolow-Lyon, Cornell homens/mulheres, Eixo)

---

## Resumo Executivo

Adicionados **6 parâmetros essenciais** de ECG para estratificação cardiovascular completa:

1. **Intervalo PR** - Bloqueios AV (BAV 1º, 2º, 3º grau)
2. **Duração QRS** - Bloqueios de ramo (BRD, BRE)
3. **Sokolow-Lyon** - Hipertrofia ventricular esquerda (HVE)
4. **Cornell Voltage - Homens** - HVE (mais específico que Sokolow-Lyon)
5. **Cornell Voltage - Mulheres** - HVE (cutoff diferente)
6. **Eixo Cardíaco** - Desvios (LAD, RAD, Extremo)

**Fontes:** American Heart Association, JACC, StatPearls, LITFL, 2024-2025

---

## 1. ECG - Intervalo PR

### Valores de Estratificação

```csv
ECG - Intervalo PR;ms | milissegundos;20;>300;<100;201 a 300;100 a 119;120 a 139;140 a 200
```

| Nível | Range (ms) | Interpretação Clínica |
|-------|------------|----------------------|
| **Nível 0** | >300 | **BAV 1º grau MARCADO** (alto risco progressão) |
| **Nível 1** | <100 | **Pré-excitação** (WPW, taquiarritmias) |
| **Nível 2** | 201-300 | **BAV 1º grau** (condução AV lenta) |
| **Nível 3** | 100-119 | **PR curto** (risco pré-excitação) |
| **Nível 4** | 120-139 | **Normal baixo** |
| **Nível 5** | 140-200 | **Ótimo** (condução AV ideal) |

### Interpretação Clínica

**Intervalo PR Normal:**
**120-200 ms** (3-5 quadradinhos pequenos no ECG)

Representa o tempo de condução do estímulo elétrico do nó sinusal → nó AV → feixe de His.

### Bloqueio Atrioventricular (BAV)

**BAV 1º Grau (PR >200 ms):**
- PR prolongado mas **todas ondas P conduzem** (1:1)
- Geralmente assintomático
- **PR >300 ms = "Marcado"** - maior risco progressão para BAV 2º/3º grau
- Causas: vagotonia (atletas), medicamentos (beta-bloqueadores, digoxina, diltiazem), doença cardíaca estrutural

**BAV 2º Grau:**
- **Mobitz I (Wenckebach):** PR progressivamente ↑ até "drop" QRS (ciclo se repete)
- **Mobitz II:** PR constante, mas QRS "drops" intermitentes (mais grave, risco BAV total)

**BAV 3º Grau (Bloqueio AV Total):**
- **Dissociação AV completa** - ondas P e QRS sem relação
- Requer marca-passo urgente

### PR Curto (<120 ms)

**Síndrome de Pré-Excitação:**
- **Wolff-Parkinson-White (WPW):** PR <120 ms + onda delta + QRS alargado
- Via acessória (bypass nó AV) → risco taquiarritmias supraventriculares
- Pode causar morte súbita (FA com condução rápida via bypass)

**PR <100 ms:**
- Muito curto - investigar WPW
- Ritmo juncional
- Lown-Ganong-Levine (raro)

### Fontes

- [StatPearls - Atrioventricular Block (2024)](https://www.ncbi.nlm.nih.gov/books/NBK459147/)
- [LITFL - PR Interval ECG Library](https://litfl.com/pr-interval-ecg-library/)
- [StatPearls - First-Degree Heart Block (2024)](https://www.ncbi.nlm.nih.gov/books/NBK448164/)

---

## 2. ECG - Duração QRS

### Valores de Estratificação

```csv
ECG - Duração QRS;ms | milissegundos;20;≥150;<70;120 a 149;70 a 79;110 a 119;80 a 109
```

| Nível | Range (ms) | Interpretação Clínica |
|-------|------------|----------------------|
| **Nível 0** | ≥150 | **BRD/BRE severo** (CRT candidato, pior prognóstico) |
| **Nível 1** | <70 | **Anormalmente estreito** (investigar) |
| **Nível 2** | 120-149 | **BRD/BRE completo** (risco IC, morte súbita) |
| **Nível 3** | 70-79 | **Estreito** (baixo normal) |
| **Nível 4** | 110-119 | **BCRD incompleto / Limítrofe** |
| **Nível 5** | 80-109 | **Ótimo** (condução intraventricular normal) |

### Interpretação Clínica

**QRS Normal:**
**<120 ms** (<3 quadradinhos pequenos)

Representa despolarização ventricular. QRS alargado indica atraso condução intraventricular.

### Bloqueios de Ramo

**Bloqueio de Ramo DIREITO (BRD):**
- **QRS ≥120 ms**
- **V1/V2:** padrão rSR' (orelha de coelho)
- **V5/V6:** onda S alargada
- Causas: hipertensão pulmonar, cor pulmonale, cardiopatia congênita, IAM VD, atletas

**BRD Incompleto:**
- QRS 110-119 ms
- Padrão rSR' em V1 mas não preenche critérios completos
- Geralmente benigno

**Bloqueio de Ramo ESQUERDO (BRE):**
- **QRS ≥120 ms**
- **V1/V2:** QS ou rS profundo
- **V5/V6:** R alargado, entalhado/corcunda ("notched")
- **Mais grave que BRD** - associado a doença cardíaca estrutural
- BRE **novo** = IAM equivalente (até prova contrário)

**Critérios Strauss (BRE verdadeiro):**
- Homens: QRS ≥140 ms
- Mulheres: QRS ≥130 ms
- Entalhamento (notching) em ≥2 derivações (V1, V2, V5, V6, I, aVL)
- Pico R >60 ms em V5-V6

### QRS Muito Largo (≥150 ms)

**Indicações Terapia Ressincronização Cardíaca (TRC/CRT):**
- QRS ≥150 ms + BRE + IC com FEVE ≤35% + sintomas = **Classe IA** (recomendação forte)
- CRT melhora mortalidade e qualidade vida

**Outras causas QRS alargado:**
- Marca-passo ventricular
- Hipercalemia (K+ >7 mEq/L)
- Intoxicação antidepressivos tricíclicos
- Cardiomiopatia (dilatada, hipertrófica)
- Taquicardia ventricular

### Prognóstico

**QRS >120 ms:**
- ↑ Risco morte súbita cardíaca
- ↑ Risco IC descompensada
- ↑ Mortalidade geral

**QRS 80-110 ms = ÓTIMO:**
- Condução intraventricular normal
- Prognóstico cardiovascular favorável

### Fontes

- [LITFL - Left Bundle Branch Block](https://litfl.com/left-bundle-branch-block-lbbb-ecg-library/)
- [LITFL - Right Bundle Branch Block](https://litfl.com/right-bundle-branch-block-rbbb-ecg-library/)
- [StatPearls - Left Bundle Branch Block (2024)](https://www.ncbi.nlm.nih.gov/books/NBK482167/)
- [AHA Journal - New ECG Criteria LBBB in AMI](https://www.ahajournals.org/doi/10.1161/JAHA.120.017119)

---

## 3. ECG - Sokolow-Lyon (S V1 + R V5/V6)

### Valores de Estratificação

```csv
ECG - Sokolow-Lyon (S V1 + R V5/V6);mm | milímetros;20;≥50;<15;40 a 49;15 a 24;35 a 39;25 a 34
```

| Nível | Range (mm) | Interpretação Clínica |
|-------|------------|----------------------|
| **Nível 0** | ≥50 | **HVE SEVERA** (ICC, morte súbita) |
| **Nível 1** | <15 | **Muito baixo** (HVD?, obesidade, DPOC) |
| **Nível 2** | 40-49 | **HVE moderada** (alto risco CV) |
| **Nível 3** | 15-24 | **Baixo** (abaixo média) |
| **Nível 4** | 35-39 | **HVE limítrofe** (zona cinza) |
| **Nível 5** | 25-34 | **Normal** (ideal) |

### Interpretação Clínica

**Critério Sokolow-Lyon (1949):**
**S V1 + R V5 ou R V6 >35 mm = HVE**

Critério mais usado mundialmente por simplicidade.

### Cálculo

1. Medir amplitude onda **S em V1** (mm)
2. Medir amplitude onda **R em V5** (mm)
3. Medir amplitude onda **R em V6** (mm)
4. **Sokolow-Lyon = S V1 + maior(R V5, R V6)**

**Exemplo:**
- S V1 = 15 mm
- R V5 = 26 mm
- R V6 = 22 mm
- **Sokolow-Lyon = 15 + 26 = 41 mm** → **HVE moderada** (Nível 2)

### Significado Clínico

**HVE (Hipertrofia Ventricular Esquerda):**

Aumento espessura parede VE como resposta a sobrecarga pressão/volume.

**Causas HVE:**
- **Hipertensão arterial sistêmica** (mais comum)
- Estenose aórtica
- Cardiomiopatia hipertrófica
- Atletas (HVE fisiológica, reversível)

**HVE é fator de risco independente para:**
- Insuficiência cardíaca
- Fibrilação atrial
- Morte súbita cardíaca
- AVC
- DAC

**Sokolow-Lyon ≥50 mm:**
- HVE severa
- Remodelamento ventricular avançado
- Alto risco eventos cardiovasculares

### Limitações

1. **Idade:** Não confiável <40 anos (jovens magros têm voltagens altas naturalmente)
   - 20-39 anos: 32% têm Sokolow-Lyon >35 mm SEM HVE
2. **Obesidade:** ↓ voltagens (falso-negativo)
3. **DPOC/enfisema:** ↓ voltagens
4. **Sensibilidade baixa:** ~40% (muitos HVE não detectados)
5. **Especificidade boa:** ~90%

**Sokolow-Lyon <15 mm:**
- Obesidade
- Derrame pericárdico
- DPOC/enfisema (hiperinsuflação pulmonar)
- Hipertrofia VD (HVD) - desvia voltagens para direita

### Fontes

- [LITFL - Left Ventricular Hypertrophy ECG](https://litfl.com/left-ventricular-hypertrophy-lvh-ecg-library/)
- [My-EKG - Sokolow-Lyon Voltage Criteria](https://en.my-ekg.com/hypertrophy-dilation/sokolow-lyon-criteria.html)
- [StatPearls - Left Ventricular Hypertrophy (2024)](https://www.ncbi.nlm.nih.gov/books/NBK557534/)

---

## 4. ECG - Cornell Voltage - HOMENS (R aVL + S V3)

### Valores de Estratificação

```csv
ECG - Cornell Voltage - Homens (R aVL + S V3);mm | milímetros;20;≥40;<10;31 a 39;10 a 19;28 a 30;20 a 27
```

| Nível | Range (mm) | Interpretação Clínica |
|-------|------------|----------------------|
| **Nível 0** | ≥40 | **HVE SEVERA** (remodelamento avançado) |
| **Nível 1** | <10 | **Muito baixo** (obesidade, DPOC) |
| **Nível 2** | 31-39 | **HVE moderada** |
| **Nível 3** | 10-19 | **Baixo** |
| **Nível 4** | 28-30 | **HVE limítrofe** (zona cinza) |
| **Nível 5** | 20-27 | **Normal** (ideal) |

### Interpretação Clínica

**Critério Cornell Voltage - Homens:**
**R aVL + S V3 >28 mm = HVE** (ou >2.8 mV)

Mais específico que Sokolow-Lyon, especialmente em idosos.

### Cálculo

1. Medir amplitude onda **R em aVL** (mm)
2. Medir amplitude onda **S em V3** (mm)
3. **Cornell Voltage = R aVL + S V3**

**Exemplo:**
- R aVL = 8 mm
- S V3 = 25 mm
- **Cornell = 8 + 25 = 33 mm** → **HVE moderada** (Nível 2)

### Vantagens Cornell vs Sokolow-Lyon

1. **Melhor performance em idosos** (>60 anos)
2. **Melhor em obesidade** (menos afetado)
3. **Especificidade ~90%**, Sensibilidade 20-40% (similar Sokolow-Lyon)
4. **Única que melhora com idade** - Sokolow-Lyon piora

**Cornell é PREFERIDO em:**
- Pacientes >60 anos
- Obesos
- Mulheres (com cutoff ajustado)

### Cornell Product (Variante)

**Cornell Product = Cornell Voltage × QRS duration**

- Homens: >244 mV·ms
- Mulheres: >204 mV·ms

Mais sensível mas menos usado na prática clínica.

### Fontes

- [JACC - ECG Criteria for LVH (2017)](https://www.jacc.org/doi/10.1016/j.jacc.2017.01.037)
- [AHA - Cornell Criterion Best in Elderly](https://www.ahajournals.org/doi/10.1161/circ.126.suppl_21.A17506)
- [PMC - Diagnostic Accuracy ECG Criteria LVH (2021)](https://pmc.ncbi.nlm.nih.gov/articles/PMC8043050/)

---

## 5. ECG - Cornell Voltage - MULHERES (R aVL + S V3)

### Valores de Estratificação

```csv
ECG - Cornell Voltage - Mulheres (R aVL + S V3);mm | milímetros;20;≥35;<8;24 a 34;8 a 14;20 a 23;15 a 19
```

| Nível | Range (mm) | Interpretação Clínica |
|-------|------------|----------------------|
| **Nível 0** | ≥35 | **HVE SEVERA** |
| **Nível 1** | <8 | **Muito baixo** |
| **Nível 2** | 24-34 | **HVE moderada** |
| **Nível 3** | 8-14 | **Baixo** |
| **Nível 4** | 20-23 | **HVE limítrofe** |
| **Nível 5** | 15-19 | **Normal** (ideal) |

### Interpretação Clínica

**Critério Cornell Voltage - Mulheres:**
**R aVL + S V3 >20 mm = HVE** (ou >2.0 mV)

**Diferença Homens vs Mulheres:**
- **Homens:** >28 mm (2.8 mV)
- **Mulheres:** >20 mm (2.0 mV)

**Razão da diferença:**
- Mulheres têm parede torácica menor
- Menor massa ventricular basal
- Voltagens absolutas menores

### Cálculo (Idêntico)

1. Medir amplitude onda **R em aVL** (mm)
2. Medir amplitude onda **S em V3** (mm)
3. **Cornell Voltage = R aVL + S V3**

**Exemplo:**
- R aVL = 6 mm
- S V3 = 18 mm
- **Cornell = 6 + 18 = 24 mm** → **HVE moderada** em mulher (Nível 2)
- Se fosse homem = **Normal** (Nível 5, pois <28 mm)

### HVE em Mulheres

**Epidemiologia:**
- Prevalência HVE ~10-20% hipertensas
- Subdiagnosticada (voltagens menores)
- Prognóstico PIOR que homens quando presente

**Fatores de Risco HVE em Mulheres:**
- Hipertensão (principal)
- Obesidade
- Pós-menopausa (↓ estrogênio protetor)
- Síndrome metabólica

### Fontes

- [JACC - ECG Criteria for LVH (2017)](https://www.jacc.org/doi/10.1016/j.jacc.2017.01.037)
- [AHA/ACCF/HRS - ECG Standardization (2009)](https://www.ahajournals.org/doi/10.1161/circulationaha.108.191097)
- [PubMed - Cornell in Women vs Men (2006)](https://pubmed.ncbi.nlm.nih.gov/16708082/)

---

## 6. ECG - Eixo Cardíaco

### Valores de Estratificação

```csv
ECG - Eixo Cardíaco;graus | ° (QRS frontal);20;-90 a -180 ou +150 a +180;+120 a +149;-60 a -89;+91 a +119;-30 a -59;-29 a +90
```

| Nível | Range (graus) | Interpretação Clínica |
|-------|---------------|----------------------|
| **Nível 0** | -90 a -180 **ou** +150 a +180 | **Eixo EXTREMO** (Northwest axis - anormal grave) |
| **Nível 1** | +120 a +149 | **RAD severo** (HVD, embolia pulmonar) |
| **Nível 2** | -60 a -89 | **LAD severo** (HVE, BCRD) |
| **Nível 3** | +91 a +119 | **RAD moderado** (atletas, magros, HVD leve) |
| **Nível 4** | -30 a -59 | **LAD leve** (comum, geralmente benigno) |
| **Nível 5** | -29 a +90 | **NORMAL** (VE dominante) |

### Interpretação Clínica

**Eixo Cardíaco Normal:**
**-30° a +90°**

Representa direção média do vetor de despolarização ventricular no plano frontal.

### Determinação Rápida do Eixo

**Método das Derivações:**

1. **Eixo Normal (-30° a +90°):**
   - QRS **positivo em I e aVF**

2. **LAD (Left Axis Deviation):**
   - QRS **positivo em I, negativo em aVF**

3. **RAD (Right Axis Deviation):**
   - QRS **negativo em I, positivo em aVF**

4. **Eixo Extremo (Northwest):**
   - QRS **negativo em I E aVF**

### Desvio de Eixo à ESQUERDA (LAD)

**LAD Leve (-30° a -59°):**
- **Muito comum, geralmente benigno**
- Variação normal em idosos
- Obesidade (diafragma elevado)
- Gravidez (útero desloca coração)
- Bloqueio divisional ântero-superior (BDAS)

**LAD Severo (-60° a -89°):**
- **HVE** (hipertensão crônica)
- **BCRD** (bloqueio completo ramo direito)
- IAM inferior (necrose parede inferior)
- Cardiomiopatia

### Desvio de Eixo à DIREITA (RAD)

**RAD Moderado (+91° a +119°):**
- **Atletas, indivíduos magros** (variante normal)
- DPOC/enfisema (coração verticalizado)
- Crianças/adolescentes (normal até 18 anos)

**RAD Severo (+120° a +149°):**
- **HVD (hipertrofia ventricular direita)**
  - Hipertensão pulmonar
  - Estenose pulmonar
  - Cor pulmonale
  - TEP crônico
- **Embolia pulmonar aguda** (sinal S1Q3T3)
- Bloqueio divisional póstero-inferior (BDPI)
- Dextrocardia

### Eixo EXTREMO (-90° a -180° ou +150° a +180°)

**"Northwest Axis" ou "No Man's Land"**

**Causas:**
1. **Eletrodos trocados** (MAIS COMUM - sempre verificar primeiro!)
2. **Taquicardia ventricular** (TV)
3. **Hiperkalemia severa** (K+ >7 mEq/L)
4. **Dextrocardia**
5. **Ectopia ventricular**

**CRÍTICO:** Se eixo extremo + taquicardia + QRS largo → **TV até prova contrário**

### Significado Prognóstico

**LAD:**
- Leve (-30 a -59°): Benigno na maioria
- Severo (<-60°): Investigar HVE, doença cardíaca estrutural

**RAD:**
- Moderado (+91 a +119°): Pode ser normal (magros, atletas)
- Severo (>+120°): **Sempre investigar HVD, embolia pulmonar**

**Eixo Extremo:**
- **Sempre patológico** (exceto eletrodos trocados)
- Requer investigação imediata

### Fontes

- [LITFL - ECG Axis Interpretation](https://litfl.com/ecg-axis-interpretation/)
- [StatPearls - Electrical Axis Deviation (2024)](https://www.ncbi.nlm.nih.gov/books/NBK470532/)
- [ECGwaves - Electrical Axis of the Heart](https://ecgwaves.com/electrical-axis-of-the-heart-ecg-physiology-definition/)

---

## Resumo - Todos os Parâmetros ECG

| Parâmetro | Normal/Ótimo | Anormal Crítico | Significado Clínico |
|-----------|--------------|-----------------|---------------------|
| **Frequência Cardíaca** | 60-89 bpm | <40 ou ≥120 bpm | Bradicardia/Taquicardia |
| **QTc - Homens** | <430 ms | ≥470 ms | Risco torsades de pointes |
| **QTc - Mulheres** | <460 ms | ≥480 ms | Risco torsades de pointes |
| **Intervalo PR** | 140-200 ms | >300 ou <100 ms | BAV ou WPW |
| **Duração QRS** | 80-109 ms | ≥150 ms | BRD/BRE severo (CRT) |
| **Sokolow-Lyon** | 25-34 mm | ≥50 mm | HVE severa |
| **Cornell - Homens** | 20-27 mm | ≥40 mm | HVE severa |
| **Cornell - Mulheres** | 15-19 mm | ≥35 mm | HVE severa |
| **Eixo Cardíaco** | -29° a +90° | Extremo (-90 a -180 ou +150 a +180) | TV, eletrodos trocados |

---

## Integração Clínica - Estratificação de Risco CV

### Baixo Risco (Score 0-2 anormalidades)
- FC normal
- QTc normal
- QRS estreito (<110 ms)
- Sem HVE (Sokolow <35, Cornell normal)
- Eixo normal

### Risco Moderado (Score 3-4 anormalidades)
- BAV 1º grau (PR 201-300)
- BRD/BRE incompleto
- HVE limítrofe (Sokolow 35-39, Cornell limítrofe)
- LAD/RAD leve

### Alto Risco (Score ≥5 anormalidades ou 1 crítica)
- BAV 1º marcado (>300 ms) ou BAV 2º/3º
- BRE completo (QRS ≥120, especialmente ≥150)
- HVE severa (Sokolow ≥50, Cornell muito elevado)
- RAD severo (suspeita HVD/embolia pulmonar)
- Eixo extremo (Northwest axis)

### Combinações de Alto Risco

**BRE + HVE:**
- Remodelamento ventricular avançado
- Candidato TRC se IC + FEVE ≤35%

**RAD severo + BRD:**
- Suspeita HVD + sobrecarga VD
- Investigar hipertensão pulmonar, TEP

**LAD severo + HVE (Cornell/Sokolow):**
- HVE avançada + alteração condução
- Alto risco IC, arritmias

---

## Fontes Completas (Todas Verificáveis)

### Intervalo PR / Bloqueios AV
- [StatPearls - Atrioventricular Block (2024)](https://www.ncbi.nlm.nih.gov/books/NBK459147/)
- [LITFL - PR Interval ECG Library](https://litfl.com/pr-interval-ecg-library/)
- [StatPearls - First-Degree Heart Block (2024)](https://www.ncbi.nlm.nih.gov/books/NBK448164/)

### Duração QRS / Bloqueios de Ramo
- [LITFL - Left Bundle Branch Block](https://litfl.com/left-bundle-branch-block-lbbb-ecg-library/)
- [LITFL - Right Bundle Branch Block](https://litfl.com/right-bundle-branch-block-rbbb-ecg-library/)
- [StatPearls - Left Bundle Branch Block (2024)](https://www.ncbi.nlm.nih.gov/books/NBK482167/)
- [AHA Journal - New ECG Criteria for AMI in LBBB](https://www.ahajournals.org/doi/10.1161/JAHA.120.017119)

### Sokolow-Lyon
- [LITFL - Left Ventricular Hypertrophy ECG](https://litfl.com/left-ventricular-hypertrophy-lvh-ecg-library/)
- [My-EKG - Sokolow-Lyon Voltage Criteria](https://en.my-ekg.com/hypertrophy-dilation/sokolow-lyon-criteria.html)
- [StatPearls - Left Ventricular Hypertrophy (2024)](https://www.ncbi.nlm.nih.gov/books/NBK557534/)

### Cornell Voltage
- [JACC - ECG Criteria for LVH (2017)](https://www.jacc.org/doi/10.1016/j.jacc.2017.01.037)
- [AHA - Cornell Criterion Best in Elderly](https://www.ahajournals.org/doi/10.1161/circ.126.suppl_21.A17506)
- [PMC - Diagnostic Accuracy ECG Criteria LVH (2021)](https://pmc.ncbi.nlm.nih.gov/articles/PMC8043050/)
- [AHA/ACCF/HRS - ECG Standardization (2009)](https://www.ahajournals.org/doi/10.1161/circulationaha.108.191097)

### Eixo Cardíaco
- [LITFL - ECG Axis Interpretation](https://litfl.com/ecg-axis-interpretation/)
- [StatPearls - Electrical Axis Deviation (2024)](https://www.ncbi.nlm.nih.gov/books/NBK470532/)
- [ECGwaves - Electrical Axis of the Heart](https://ecgwaves.com/electrical-axis-of-the-heart-ecg-physiology-definition/)

---

## Validação Final

✅ **9 parâmetros de ECG** completos (FC, QTc H/M, PR, QRS, Sokolow-Lyon, Cornell H/M, Eixo)
✅ **Todos com EXATAMENTE 6 níveis** (Nível 0 a 5)
✅ **Valores REAIS** de cardiologia (AHA, ACC, ESC, StatPearls 2024-2025)
✅ **Cornell separado por sexo** (homens >28 mm, mulheres >20 mm)
✅ **Fontes verificáveis** com links diretos
✅ **Integração clínica** com estratificação de risco combinada

---

**Status:** ✅ **PARÂMETROS DE ECG COMPLETOS**
**Última atualização:** Janeiro 2026
