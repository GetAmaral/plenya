# Estratificação de Risco - Tomografia Cardíaca (CAC Score e CCTA)

## Índice

1. [CAC Score (Escore de Cálcio Coronariano)](#cac-score-escore-de-cálcio-coronariano)
2. [CCTA (Angiotomografia Coronariana) - CAD-RADS](#ccta-angiotomografia-coronariana---cad-rads)
3. [Placa Burden Score (PBS)](#placa-burden-score-pbs)
4. [High-Risk Plaque Features](#high-risk-plaque-features)
5. [Integração com Medicina Funcional](#integração-com-medicina-funcional)
6. [Algoritmos Clínicos de Decisão](#algoritmos-clínicos-de-decisão)
7. [Evidências Científicas](#evidências-científicas)
8. [Implementação no Sistema EMR](#implementação-no-sistema-emr)

---

## CAC Score (Escore de Cálcio Coronariano)

### Tabela de Estratificação de Risco - Agatston Score (QUANTITATIVO)

```csv
CAC Score (Agatston)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 | Nível 4 | Nível 5 |
| >1000 AU | 401-1000 AU | 101-400 AU | 11-100 AU | 1-10 AU | 0 AU |
Agatston Units (AU) | Absoluto (sem conversão)
```

**Interpretação:**
- **Nível 5 (0 AU)**: Ausência de placa, risco muito baixo (0.11% eventos/ano)
- **Nível 4 (1-10 AU)**: Placa mínima, risco baixo (~0.3% eventos/ano)
- **Nível 3 (11-100 AU)**: Placa leve, risco moderado (~0.7% eventos/ano)
- **Nível 2 (101-400 AU)**: Placa moderada, risco moderado-alto (~1.4% eventos/ano)
- **Nível 1 (401-1000 AU)**: Placa severa, alto risco (~2.1% eventos/ano)
- **Nível 0 (>1000 AU)**: Placa extensa, risco muito alto (3-4% eventos/ano)

### Tabela de Estratificação por Percentil (Idade/Sexo Específico)

```csv
CAC Percentil (Idade/Sexo - MESA Reference)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 | Nível 4 |
| ≥90th | 75-89th | 50-74th | 25-49th | <25th |
Percentil | Baseado em idade, sexo e raça (MESA study)
```

**Interpretação:**
- **Nível 4 (<25th percentil)**: Melhor que a média da população
- **Nível 3 (25-49th)**: Risco médio para idade/sexo
- **Nível 2 (50-74th)**: Acima da média
- **Nível 1 (75-89th)**: Alto para idade/sexo
- **Nível 0 (≥90th)**: Muito alto para idade/sexo

### Tabela de Progressão Anual do CAC

```csv
CAC Progressão Anual
| Nível 0 | Nível 1 | Nível 2 | Nível 3 |
| >30%/ano | 20-30%/ano | 15-20%/ano | <15%/ano |
Porcentagem | Mudança percentual anual do score
```

**Interpretação:**
- **Nível 3 (<15%/ano)**: Progressão normal/lenta
- **Nível 2 (15-20%/ano)**: Progressão típica esperada
- **Nível 1 (20-30%/ano)**: Progressão acelerada, requer intervenção intensiva
- **Nível 0 (>30%/ano)**: Progressão muito acelerada, alto risco

### CAC Density Score (Qualitativo)

```csv
CAC Densidade
| Nível 0 | Nível 1 | Nível 2 | Nível 3 |
| Baixa densidade (<130 HU) | Mista | Moderada (130-400 HU) | Alta densidade (>400 HU) |
Hounsfield Units (HU) | Densidade da placa calcificada
```

**Interpretação:**
- **Nível 3 (Alta densidade)**: Placa calcificada estável (menor risco agudo)
- **Nível 2 (Moderada)**: Risco intermediário
- **Nível 1 (Mista)**: Combinação de placa lipídica e calcificada
- **Nível 0 (Baixa densidade)**: Placa lipídica instável (maior risco agudo)

---

## CCTA (Angiotomografia Coronariana) - CAD-RADS

### Tabela de Estratificação - CAD-RADS 2.0 (SEMI-QUANTITATIVO)

```csv
CCTA - CAD-RADS (Stenosis Severity)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 | Nível 4 | Nível 5 | Nível 6 |
| CAD-RADS 5 (100%) | CAD-RADS 4B | CAD-RADS 4A | CAD-RADS 3 | CAD-RADS 2 | CAD-RADS 1 | CAD-RADS 0 |
Porcentagem | 0% = sem estenose, 100% = oclusão total
```

**Classificação Detalhada:**

- **Nível 6 (CAD-RADS 0)**: Sem placa ou estenose (0%)
  - Evento 5 anos: ~5% (95.2% livre de eventos)
  - Conduta: Sem necessidade de seguimento específico

- **Nível 5 (CAD-RADS 1)**: Estenose mínima (1-24%)
  - Evento 5 anos: ~7% (92.9% livre de eventos)
  - Conduta: Prevenção primária agressiva

- **Nível 4 (CAD-RADS 2)**: Estenose leve não-obstrutiva (25-49%)
  - Evento 5 anos: ~11% (88.7% livre de eventos)
  - Conduta: Prevenção secundária, controle fatores de risco

- **Nível 3 (CAD-RADS 3)**: Estenose moderada (50-69%)
  - Evento 5 anos: ~15% (84.5% livre de eventos)
  - Conduta: Considerar teste de isquemia (stress test, FFR-CT)

- **Nível 2 (CAD-RADS 4A)**: Estenose severa 1-2 vasos (70-99%)
  - Evento 5 anos: ~23% (76.7% livre de eventos)
  - Conduta: Provável angiografia invasiva + revascularização

- **Nível 1 (CAD-RADS 4B)**: Estenose severa 3 vasos ou TCE >50%
  - Evento 5 anos: ~23% (similar 4A)
  - Conduta: Angiografia invasiva urgente

- **Nível 0 (CAD-RADS 5)**: Oclusão total (100%)
  - Evento 5 anos: ~31% (69.3% livre de eventos)
  - Conduta: Angiografia invasiva + revascularização urgente

### Modificadores CAD-RADS 2.0

**CAD-RADS N (Non-diagnostic):**
- Exame tecnicamente inadequado
- Requer repetição ou método alternativo

**CAD-RADS S (Stent):**
- Avaliação de stent coronariano prévio

**CAD-RADS G (Graft):**
- Avaliação de enxerto de bypass prévio

---

## Plaque Burden Score (PBS)

### Tabela de Carga de Placa Total

```csv
Plaque Burden Score (CAD-RADS 2.0)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 | Nível 4 |
| P4 (Extenso) | P3 (Moderado-Severo) | P2 (Moderado) | P1 (Leve) | P0 (Ausente) |
Categórico | Baseado em SIS ou CAC ou avaliação visual
```

**Métodos de Quantificação:**

1. **CAC Score (se disponível):**
   - P0: CAC = 0
   - P1: CAC 1-99
   - P2: CAC 100-399
   - P3: CAC 400-999
   - P4: CAC ≥1000

2. **Segment Involvement Score (SIS) - Modelo 16 segmentos:**
   - P0: 0 segmentos
   - P1: 1-4 segmentos
   - P2: 5-8 segmentos
   - P3: 9-12 segmentos
   - P4: ≥13 segmentos

3. **Avaliação Visual Qualitativa:**
   - P0: Sem placa
   - P1: Placa focal em 1-2 segmentos
   - P2: Placa difusa em múltiplos segmentos
   - P3: Placa difusa + calcificação extensa
   - P4: Placa difusa + calcificação massiva

---

## High-Risk Plaque Features

### Características de Placa de Alto Risco (HRP)

**Critérios Principais (2024-2025):**

1. **Positive Remodeling (Remodelamento Positivo)**
   - Definição: Área vascular externa >10% maior que referência
   - Significado: Placa vulnerável em expansão
   - Risco: Alto para ruptura

2. **Low-Attenuation Plaque (LAP)**
   - Definição: <30 Hounsfield Units
   - Significado: Placa rica em lipídios
   - Risco: Altamente vulnerável

3. **Napkin-Ring Sign (Sinal do Anel)**
   - Definição: Core central hipodenso com anel hiperdenso periférico
   - Significado: Core necrótico lipídico com cápsula fibrótica fina
   - Risco: Muito alto para eventos agudos

4. **Spotty Calcification (Calcificação Pontual)**
   - Definição: Calcificação <3mm
   - Significado: Calcificação precoce em placa vulnerável
   - Risco: Marcador de instabilidade

### Tabela HRP Features Score

```csv
High-Risk Plaque Features (HRP)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 | Nível 4 |
| 4 features | 3 features | 2 features | 1 feature | 0 features |
Contagem | Número de características de alto risco presentes
```

**Interpretação:**
- **Nível 4 (0 features)**: Sem características de alto risco
- **Nível 3 (1 feature)**: Risco moderadamente aumentado
- **Nível 2 (2 features)**: Alto risco
- **Nível 1 (3 features)**: Risco muito alto
- **Nível 0 (4 features)**: Risco extremo

---

## Integração com Medicina Funcional

### Painel Avançado de Lipídios + CAC Score

**Abordagem Integrativa:**

```csv
ApoB (Apolipoprotein B)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 | Nível 4 | Nível 5 |
| >130 mg/dL | 100-130 mg/dL | 80-99 mg/dL | 65-79 mg/dL | 50-64 mg/dL | <50 mg/dL |
mg/dL | 1 mg/dL = 0.01 g/L
```

```csv
Lp(a) - Lipoproteína(a)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 | Nível 4 |
| >150 mg/dL | 100-150 mg/dL | 50-99 mg/dL | 30-49 mg/dL | <30 mg/dL |
mg/dL | 1 mg/dL = 2.5 nmol/L
```

```csv
LDL-P (LDL Particle Number)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 | Nível 4 |
| >2000 nmol/L | 1600-2000 nmol/L | 1300-1599 nmol/L | 1000-1299 nmol/L | <1000 nmol/L |
nmol/L | Partículas LDL por litro
```

### Marcadores Inflamatórios Complementares

```csv
hs-CRP (High-Sensitivity C-Reactive Protein)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 | Nível 4 |
| >10 mg/L | 3-10 mg/L | 1-3 mg/L | 0.5-1 mg/L | <0.5 mg/L |
mg/L | 1 mg/L = 1000 µg/dL
```

```csv
Lp-PLA2 (Lipoprotein-Associated Phospholipase A2)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 | Nível 4 |
| >235 ng/mL | 200-235 ng/mL | 175-199 ng/mL | 150-174 ng/mL | <150 ng/mL |
ng/mL | Marcador específico de inflamação vascular
```

### Modelo de Risco Combinado (CAC + ApoB + Lp(a))

**Estratificação de Risco Integrada:**

| CAC Score | ApoB | Lp(a) | Risco 10 anos | Estratégia |
|-----------|------|-------|---------------|------------|
| 0 | <80 | <50 | <5% | Lifestyle, sem estatina |
| 0 | >80 | <50 | 5-10% | Lifestyle + considerar estatina |
| 0 | Qualquer | >100 | 10-15% | Estatina + terapia anti-Lp(a) (futuro) |
| 1-99 | >80 | >50 | 15-20% | Estatina + ezetimiba |
| 100-400 | >80 | >50 | 20-30% | Estatina + ezetimiba + PCSK9i |
| >400 | Qualquer | Qualquer | >30% | Terapia intensiva máxima |

---

## Algoritmos Clínicos de Decisão

### Algoritmo 1: CAC Score para Decisão de Estatina (ACC/AHA 2019)

```
Risco Intermediário (PCE 7.5-20%)
         ↓
    Dúvida sobre estatina?
         ↓
    Solicitar CAC Score
         ↓
    ┌────────────┬──────────────┬──────────────┐
    ↓            ↓              ↓              ↓
CAC = 0      CAC 1-99      CAC 100-400    CAC >400
    ↓            ↓              ↓              ↓
Não tratar   Decisão       Estatina       Estatina
(reassess    compartilhada  moderada/      alta
5-10 anos)   (idade>55     alta           intensidade
             favorece)     intensidade    + ezetimiba
```

### Algoritmo 2: CAD-RADS para Conduta Clínica

```
CCTA Realizada
       ↓
  CAD-RADS Score
       ↓
    ┌──────┬──────┬──────┬──────┬──────┐
    ↓      ↓      ↓      ↓      ↓      ↓
   0-1     2      3      4A     4B     5
    ↓      ↓      ↓      ↓      ↓      ↓
Lifestyle  Prev.  Stress ICA +  ICA    ICA
only      2ária  Test   revasc urgente urgente
                        provável
```

### Algoritmo 3: Warranty Period (Período de Garantia CAC=0)

**Quando repetir CAC se score inicial = 0:**

| Idade | Sexo | Diabetes | Lp(a)>50 | Repetir CAC em: |
|-------|------|----------|----------|-----------------|
| <45   | M    | Não      | Não      | 10-15 anos      |
| <45   | M    | Não      | Sim      | 5-7 anos        |
| <45   | M    | Sim      | Não/Sim  | 3-5 anos        |
| 45-55 | M    | Não      | Não      | 7-10 anos       |
| 45-55 | M    | Não      | Sim      | 5 anos          |
| 45-55 | M    | Sim      | Qualquer | 3-5 anos        |
| >55   | M    | Não      | Não      | 5-7 anos        |
| >55   | M    | Sim      | Qualquer | 3 anos          |
| <55   | F    | Não      | Não      | 10-15 anos      |
| <55   | F    | Não      | Sim      | 7-10 anos       |
| >55   | F    | Sim      | Qualquer | 3-5 anos        |

**Fatores que encurtam warranty period:**
- Diabetes mellitus
- Lp(a) >50 mg/dL
- História familiar forte (parente 1º grau <55 anos M / <65 anos F)
- Tabagismo ativo
- HbA1c >6.5%
- ApoB >100 mg/dL

---

## Evidências Científicas

### CAC Score - Estudos Principais (2023-2026)

**1. Multi-Ethnic Study of Atherosclerosis (MESA)**
- **População:** 6,814 adultos multi-étnicos, idade 45-84 anos
- **Seguimento:** 15+ anos
- **Achados principais:**
  - CAC=0: Warranty period 5-15 anos (varia por demografia)
  - CAC percentil >75th: Risco equivalente a prevenção secundária
  - CAC prediz eventos melhor que Framingham/PCE
  - Progressão típica: 15-25%/ano

**2. CAC Consortium (Pooled Cohort - 2024)**
- **População:** >66,000 pacientes, 4 continentes
- **Achados principais:**
  - CAC >400: Risco 10 anos >20% (equivalente prevenção 2ária)
  - CAC 0 em diabetes: Ainda baixo risco (<1%/ano)
  - Cada 100 AU aumento = 10-15% aumento risco relativo

**3. CAC Staging Proposal (JACC Advances 2024)**
- **Proposta:** Sistema de estadiamento CAC similar a câncer
  - Stage 0: CAC=0
  - Stage I: CAC 1-99
  - Stage II: CAC 100-399
  - Stage III: CAC 400-999
  - Stage IV: CAC ≥1000
- **Objetivo:** Guiar intensidade de terapia preventiva

### CCTA - Estudos Principais

**1. CAD-RADS 2.0 Consensus (2022)**
- **Sociedades:** SCCT, ACC, ACR, NASCI
- **Inovações:**
  - Plaque burden score (P0-P4)
  - Integração FFR-CT
  - Maior ênfase em high-risk plaque features

**2. SCOT-HEART Trial**
- **População:** 4,146 pacientes com dor torácica estável
- **Achados:**
  - CCTA reduziu eventos cardiovasculares em 41%
  - CAD-RADS ≥3: Benefício de tratamento agressivo
  - High-risk plaque features: Preditor independente

**3. Quantitative Coronary Plaque Analysis (ACC 2025)**
- **Scientific Statement:** Padronização análise quantitativa
- **Destaque:**
  - Positive remodeling + LAP: Maior risco
  - Napkin-ring sign: HR 2.5 para MACE
  - Spotty calcification: Marcador precoce vulnerabilidade

### Estudos de Intervenção (Medicina Funcional)

**1. AVADEC Trial (2024) - Vitamin K2 + D3**
- **População:** Pacientes com CAC ≥400
- **Intervenção:** Vitamina K2 (720 µg/dia) + D3 (25 µg/dia) vs placebo, 2 anos
- **Resultados:**
  - Grupo intervenção: Progressão CAC reduzida (Δ288 vs Δ380, p=0.047)
  - Eventos cardiovasculares: 65% redução (1.9% vs 6.7%, p=0.048)
  - **Interpretação:** Vitamina K2+D3 desacelera progressão em CAC alto

**2. Ornish Lifestyle Medicine**
- **Intervenção:** Dieta plant-based, exercício, meditação, suporte social
- **Resultados:**
  - Reversão de estenose coronariana em 82% pacientes
  - Redução eventos cardiovasculares 2.5x
  - Mudança expressão gênica (downregulation genes inflamatórios)

**3. Low-Carb/Mediterranean Diet + Exercise (Meta-analysis 2024)**
- **Achados:**
  - Dieta low-carb: Redução triglicerídeos 25-30%
  - Mediterranean: Redução eventos 30%
  - Exercício aeróbico: Desacelera progressão CAC

---

## Implementação no Sistema EMR

### Estrutura de Dados - Go Models

**1. CAC Score Model:**

```go
// apps/api/internal/models/cac_score.go

type CACScore struct {
    ID                 uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    PatientID          uuid.UUID      `gorm:"type:uuid;not null;index" json:"patientId" validate:"required"`

    // Exam info
    ExamDate           time.Time      `gorm:"not null" json:"examDate" validate:"required"`
    Institution        string         `gorm:"type:varchar(200)" json:"institution"`

    // CAC Score (Agatston)
    AgatstonScore      int            `gorm:"not null" json:"agatstonScore" validate:"gte=0"`
    AgatstonRiskLevel  int            `gorm:"not null;check:agatston_risk_level BETWEEN 0 AND 5" json:"agatstonRiskLevel"`

    // Percentile (age/sex/race-specific)
    Percentile         *int           `gorm:"check:percentile BETWEEN 0 AND 100" json:"percentile,omitempty"`
    PercentileRiskLevel *int          `gorm:"check:percentile_risk_level BETWEEN 0 AND 4" json:"percentileRiskLevel,omitempty"`

    // CAC Density
    CACDensity         *string        `gorm:"type:varchar(50);check:cac_density IN ('low','mixed','moderate','high')" json:"cacDensity,omitempty"`
    CACDensityRiskLevel *int          `gorm:"check:cac_density_risk_level BETWEEN 0 AND 3" json:"cacDensityRiskLevel,omitempty"`

    // Vessel-specific scores
    LMScore            *int           `json:"lmScore,omitempty"` // Left Main
    LADScore           *int           `json:"ladScore,omitempty"` // Left Anterior Descending
    LCXScore           *int           `json:"lcxScore,omitempty"` // Left Circumflex
    RCAScore           *int           `json:"rcaScore,omitempty"` // Right Coronary Artery

    // Progression (if repeat scan)
    PreviousScanID     *uuid.UUID     `gorm:"type:uuid" json:"previousScanId,omitempty"`
    ProgressionRate    *float64       `json:"progressionRate,omitempty"` // Annual % change
    ProgressionRiskLevel *int         `gorm:"check:progression_risk_level BETWEEN 0 AND 3" json:"progressionRiskLevel,omitempty"`

    // Clinical recommendations
    StatinRecommended  bool           `gorm:"default:false" json:"statinRecommended"`
    WarrantyPeriodYears *int          `gorm:"check:warranty_period_years BETWEEN 1 AND 15" json:"warrantyPeriodYears,omitempty"`

    // Audit
    CreatedAt          time.Time      `gorm:"autoCreateTime" json:"createdAt"`
    UpdatedAt          time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`

    // Relations
    Patient            Patient        `gorm:"foreignKey:PatientID" json:"-"`
}

// Calculate risk level based on Agatston score
func (c *CACScore) CalculateRiskLevel() {
    switch {
    case c.AgatstonScore == 0:
        c.AgatstonRiskLevel = 5
    case c.AgatstonScore <= 10:
        c.AgatstonRiskLevel = 4
    case c.AgatstonScore <= 100:
        c.AgatstonRiskLevel = 3
    case c.AgatstonScore <= 400:
        c.AgatstonRiskLevel = 2
    case c.AgatstonScore <= 1000:
        c.AgatstonRiskLevel = 1
    default:
        c.AgatstonRiskLevel = 0
    }
}

// Calculate warranty period based on patient demographics
func (c *CACScore) CalculateWarrantyPeriod(patient *Patient, lpa *float64, diabetes bool) {
    if c.AgatstonScore > 0 {
        c.WarrantyPeriodYears = nil
        return
    }

    age := time.Now().Year() - patient.BirthDate.Year()
    hasHighLpa := lpa != nil && *lpa > 50

    // Base warranty period
    years := 10

    // Adjust for diabetes
    if diabetes {
        years = 3
        c.WarrantyPeriodYears = &years
        return
    }

    // Adjust for age and Lp(a)
    if age < 45 {
        if hasHighLpa {
            years = 5
        } else {
            years = 15
        }
    } else if age < 55 {
        if hasHighLpa {
            years = 5
        } else {
            years = 7
        }
    } else {
        if hasHighLpa {
            years = 3
        } else {
            years = 5
        }
    }

    c.WarrantyPeriodYears = &years
}
```

**2. CCTA Model:**

```go
// apps/api/internal/models/ccta.go

type CCTA struct {
    ID                    uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    PatientID             uuid.UUID      `gorm:"type:uuid;not null;index" json:"patientId" validate:"required"`

    // Exam info
    ExamDate              time.Time      `gorm:"not null" json:"examDate" validate:"required"`
    Institution           string         `gorm:"type:varchar(200)" json:"institution"`

    // CAD-RADS Classification
    CADRADSScore          string         `gorm:"type:varchar(10);not null;check:cad_rads_score IN ('0','1','2','3','4A','4B','5','N','S','G')" json:"cadRadsScore" validate:"required"`
    CADRADSRiskLevel      int            `gorm:"not null;check:cad_rads_risk_level BETWEEN 0 AND 6" json:"cadRadsRiskLevel"`

    // Plaque Burden Score (P0-P4)
    PlaqueBurdenScore     string         `gorm:"type:varchar(10);check:plaque_burden_score IN ('P0','P1','P2','P3','P4')" json:"plaqueBurdenScore,omitempty"`
    PlaqueBurdenRiskLevel *int           `gorm:"check:plaque_burden_risk_level BETWEEN 0 AND 4" json:"plaqueBurdenRiskLevel,omitempty"`

    // Segment Involvement Score
    SegmentInvolvementScore *int         `gorm:"check:segment_involvement_score BETWEEN 0 AND 16" json:"segmentInvolvementScore,omitempty"`

    // High-Risk Plaque Features
    PositiveRemodeling    bool           `gorm:"default:false" json:"positiveRemodeling"`
    LowAttenuationPlaque  bool           `gorm:"default:false" json:"lowAttenuationPlaque"`
    NapkinRingSign        bool           `gorm:"default:false" json:"napkinRingSign"`
    SpottyCalcification   bool           `gorm:"default:false" json:"spottyCalcification"`
    HRPFeatureCount       int            `gorm:"check:hrp_feature_count BETWEEN 0 AND 4" json:"hrpFeatureCount"`
    HRPRiskLevel          int            `gorm:"check:hrp_risk_level BETWEEN 0 AND 4" json:"hrpRiskLevel"`

    // Stenosis details (worst lesion)
    WorstStenosisPercent  *int           `gorm:"check:worst_stenosis_percent BETWEEN 0 AND 100" json:"worstStenosisPercent,omitempty"`
    WorstStenosisLocation string         `gorm:"type:varchar(50)" json:"worstStenosisLocation,omitempty"`

    // FFR-CT (if performed)
    FFRCT                 *float64       `gorm:"check:ffrct BETWEEN 0 AND 1" json:"ffrct,omitempty"`
    FFRCTPositive         *bool          `json:"ffrctPositive,omitempty"` // <0.80 = positive

    // Clinical recommendations
    InvasiveAngiography   bool           `gorm:"default:false" json:"invasiveAngiography"`
    StressTestRecommended bool           `gorm:"default:false" json:"stressTestRecommended"`
    MedicalTherapyIntensity string       `gorm:"type:varchar(50);check:medical_therapy_intensity IN ('lifestyle','moderate','intensive','maximal')" json:"medicalTherapyIntensity"`

    // Associated CAC Score
    CACScoreID            *uuid.UUID     `gorm:"type:uuid" json:"cacScoreId,omitempty"`

    // Audit
    CreatedAt             time.Time      `gorm:"autoCreateTime" json:"createdAt"`
    UpdatedAt             time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`

    // Relations
    Patient               Patient        `gorm:"foreignKey:PatientID" json:"-"`
    CACScore              *CACScore      `gorm:"foreignKey:CACScoreID" json:"cacScore,omitempty"`
}

// Calculate risk level based on CAD-RADS
func (c *CCTA) CalculateRiskLevel() {
    switch c.CADRADSScore {
    case "0":
        c.CADRADSRiskLevel = 6
    case "1":
        c.CADRADSRiskLevel = 5
    case "2":
        c.CADRADSRiskLevel = 4
    case "3":
        c.CADRADSRiskLevel = 3
    case "4A":
        c.CADRADSRiskLevel = 2
    case "4B":
        c.CADRADSRiskLevel = 1
    case "5":
        c.CADRADSRiskLevel = 0
    }
}

// Calculate HRP feature count and risk
func (c *CCTA) CalculateHRPRisk() {
    count := 0
    if c.PositiveRemodeling {
        count++
    }
    if c.LowAttenuationPlaque {
        count++
    }
    if c.NapkinRingSign {
        count++
    }
    if c.SpottyCalcification {
        count++
    }

    c.HRPFeatureCount = count
    c.HRPRiskLevel = 4 - count // 4=best, 0=worst
}

// Determine clinical recommendations
func (c *CCTA) DetermineRecommendations() {
    switch c.CADRADSScore {
    case "0", "1":
        c.InvasiveAngiography = false
        c.StressTestRecommended = false
        c.MedicalTherapyIntensity = "lifestyle"
    case "2":
        c.InvasiveAngiography = false
        c.StressTestRecommended = false
        c.MedicalTherapyIntensity = "moderate"
    case "3":
        c.InvasiveAngiography = false
        c.StressTestRecommended = true
        c.MedicalTherapyIntensity = "intensive"
    case "4A":
        c.InvasiveAngiography = true
        c.StressTestRecommended = false
        c.MedicalTherapyIntensity = "maximal"
    case "4B", "5":
        c.InvasiveAngiography = true
        c.StressTestRecommended = false
        c.MedicalTherapyIntensity = "maximal"
    }
}
```

### Exemplo de Endpoint API

```go
// POST /api/v1/patients/:id/cac-score
func (h *Handler) CreateCACScore(c *fiber.Ctx) error {
    var req struct {
        ExamDate           time.Time  `json:"examDate" validate:"required"`
        Institution        string     `json:"institution"`
        AgatstonScore      int        `json:"agatstonScore" validate:"required,gte=0"`
        Percentile         *int       `json:"percentile"`
        CACDensity         *string    `json:"cacDensity"`
        LMScore            *int       `json:"lmScore"`
        LADScore           *int       `json:"ladScore"`
        LCXScore           *int       `json:"lcxScore"`
        RCAScore           *int       `json:"rcaScore"`
        PreviousScanID     *uuid.UUID `json:"previousScanId"`
    }

    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
    }

    patientID := c.Params("id")

    // Get patient to calculate warranty period
    patient, err := h.patientRepo.FindByID(patientID)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Patient not found"})
    }

    // Get advanced lipids to check Lp(a)
    labs, _ := h.labRepo.GetLatestByPatient(patientID, "lpa")
    var lpaValue *float64
    if len(labs) > 0 {
        val := labs[0].Value
        lpaValue = &val
    }

    // Check diabetes from patient record
    hasDiabetes := patient.HasDiagnosis("diabetes")

    // Create CAC score
    cac := &models.CACScore{
        PatientID:      uuid.MustParse(patientID),
        ExamDate:       req.ExamDate,
        Institution:    req.Institution,
        AgatstonScore:  req.AgatstonScore,
        Percentile:     req.Percentile,
        CACDensity:     req.CACDensity,
        LMScore:        req.LMScore,
        LADScore:       req.LADScore,
        LCXScore:       req.LCXScore,
        RCAScore:       req.RCAScore,
        PreviousScanID: req.PreviousScanID,
    }

    // Calculate risk levels
    cac.CalculateRiskLevel()
    cac.CalculateWarrantyPeriod(patient, lpaValue, hasDiabetes)

    // Calculate progression if previous scan exists
    if req.PreviousScanID != nil {
        prevScan, err := h.cacRepo.FindByID(*req.PreviousScanID)
        if err == nil {
            years := cac.ExamDate.Sub(prevScan.ExamDate).Hours() / 24 / 365
            if years > 0 {
                progression := float64(cac.AgatstonScore-prevScan.AgatstonScore) / float64(prevScan.AgatstonScore) / years * 100
                cac.ProgressionRate = &progression

                // Determine progression risk level
                switch {
                case progression < 15:
                    level := 3
                    cac.ProgressionRiskLevel = &level
                case progression < 20:
                    level := 2
                    cac.ProgressionRiskLevel = &level
                case progression < 30:
                    level := 1
                    cac.ProgressionRiskLevel = &level
                default:
                    level := 0
                    cac.ProgressionRiskLevel = &level
                }
            }
        }
    }

    // Determine statin recommendation
    cac.StatinRecommended = cac.AgatstonScore >= 100

    // Save to database
    if err := h.cacRepo.Create(cac); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to create CAC score"})
    }

    // Create audit log
    h.auditLog.Log(c, "cac_score", "create", cac.ID.String())

    return c.Status(201).JSON(cac)
}
```

### Interface de Visualização (Frontend)

**Componente React para CAC Score Dashboard:**

```typescript
// apps/web/components/cac-score-card.tsx

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Progress } from '@/components/ui/progress';
import { TrendingUp, TrendingDown, AlertTriangle } from 'lucide-react';

interface CACScoreCardProps {
  agatstonScore: number;
  riskLevel: number;
  percentile?: number;
  progressionRate?: number;
  warrantyPeriod?: number;
  statinRecommended: boolean;
}

const riskLabels = [
  { level: 0, label: 'Risco Muito Alto', color: 'destructive', bg: 'bg-red-100' },
  { level: 1, label: 'Risco Alto', color: 'destructive', bg: 'bg-orange-100' },
  { level: 2, label: 'Risco Moderado-Alto', color: 'warning', bg: 'bg-yellow-100' },
  { level: 3, label: 'Risco Moderado', color: 'warning', bg: 'bg-yellow-50' },
  { level: 4, label: 'Risco Baixo', color: 'success', bg: 'bg-green-100' },
  { level: 5, label: 'Risco Muito Baixo', color: 'success', bg: 'bg-green-50' },
];

export function CACScoreCard({
  agatstonScore,
  riskLevel,
  percentile,
  progressionRate,
  warrantyPeriod,
  statinRecommended,
}: CACScoreCardProps) {
  const risk = riskLabels[riskLevel];

  return (
    <Card className={risk.bg}>
      <CardHeader>
        <CardTitle className="flex items-center justify-between">
          <span>CAC Score (Agatston)</span>
          <Badge variant={risk.color}>{risk.label}</Badge>
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        <div>
          <div className="text-4xl font-bold">{agatstonScore} AU</div>
          {percentile && (
            <p className="text-sm text-muted-foreground mt-1">
              Percentil {percentile} para idade/sexo
            </p>
          )}
        </div>

        <Progress value={(5 - riskLevel) * 20} className="h-2" />

        {progressionRate !== undefined && (
          <div className="flex items-center gap-2">
            {progressionRate > 20 ? (
              <TrendingUp className="h-4 w-4 text-red-500" />
            ) : (
              <TrendingDown className="h-4 w-4 text-green-500" />
            )}
            <span className="text-sm">
              Progressão: {progressionRate.toFixed(1)}%/ano
            </span>
          </div>
        )}

        {warrantyPeriod && agatstonScore === 0 && (
          <div className="bg-white/50 p-3 rounded-md">
            <p className="text-sm font-medium">Período de Garantia</p>
            <p className="text-xs text-muted-foreground">
              Repetir CAC em {warrantyPeriod} anos
            </p>
          </div>
        )}

        {statinRecommended && (
          <div className="flex items-center gap-2 bg-blue-50 p-3 rounded-md">
            <AlertTriangle className="h-4 w-4 text-blue-600" />
            <p className="text-sm font-medium text-blue-900">
              Estatina recomendada (CAC ≥100)
            </p>
          </div>
        )}
      </CardContent>
    </Card>
  );
}
```

---

## Referências Bibliográficas

### Guidelines e Consensos

1. **2019 ACC/AHA Guideline on the Primary Prevention of Cardiovascular Disease**
   - Recomendação CAC Score (Class IIa) para risco intermediário
   - CAC ≥100 ou ≥75th percentil = alto risco

2. **CAD-RADS 2.0 Consensus Document (2022)**
   - SCCT, ACC, ACR, NASCI
   - Padronização CCTA reporting

3. **2024 ACC Scientific Statement: Quantitative Coronary Plaque Analysis**
   - Análise quantitativa de placa
   - High-risk plaque features

### Estudos MESA (Multi-Ethnic Study of Atherosclerosis)

4. **Warranty Period of a Calcium Score of Zero (2020)**
   - JACC Cardiovascular Imaging
   - Warranty period: 3-7 anos (varia por demografia)

5. **Ten-year Trajectory of CAC and Risk (2024)**
   - Progressão CAC e eventos cardiovasculares
   - Trajetórias de risco longitudinais

### Intervenções e Medicina Funcional

6. **AVADEC Trial (2024) - Vitamin K2 and D3**
   - JACC Advances
   - Redução progressão CAC em scores ≥400

7. **Ornish Lifestyle Medicine**
   - Reversão de doença coronariana
   - Mudanças na expressão gênica

### Advanced Lipid Biomarkers

8. **Role of ApoB in Clinical Management (2025)**
   - National Lipid Association Expert Consensus
   - ApoB superior a LDL-C

9. **Lp(a) in Primary Prevention (2024)**
   - American College of Cardiology
   - Medição universal recomendada

---

## Glossário

**ACC:** American College of Cardiology
**AHA:** American Heart Association
**ApoB:** Apolipoprotein B (principal proteína das partículas LDL)
**ASCVD:** Atherosclerotic Cardiovascular Disease
**AU:** Agatston Units (unidade do CAC score)
**CAC:** Coronary Artery Calcium
**CAD-RADS:** Coronary Artery Disease - Reporting and Data System
**CCTA:** Coronary Computed Tomography Angiography
**FFR-CT:** Fractional Flow Reserve by CT
**HRP:** High-Risk Plaque
**HU:** Hounsfield Units (unidade de densidade tomográfica)
**ICA:** Invasive Coronary Angiography
**LAD:** Left Anterior Descending artery
**LAP:** Low-Attenuation Plaque
**LCX:** Left Circumflex artery
**LM:** Left Main coronary artery
**Lp(a):** Lipoprotein(a)
**MACE:** Major Adverse Cardiovascular Events
**MESA:** Multi-Ethnic Study of Atherosclerosis
**PCE:** Pooled Cohort Equations
**RCA:** Right Coronary Artery
**SIS:** Segment Involvement Score
**TCE:** Tronco de Coronária Esquerda (Left Main)

---

**Última atualização:** Janeiro 2026
**Versão:** 1.0
**Autor:** Sistema Plenya EMR - Medicina Funcional Integrativa
