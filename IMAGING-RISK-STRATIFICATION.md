# Estratificação de Risco - Exames de Imagem

## Visão Geral

Este documento analisa três exames de imagem essenciais para estratificação de risco em EMR:

1. **USG Abdome Total** - Avaliação hepática, vesícula, rins, baço
2. **USG Próstata (via Abdominal)** - Volume prostático, BPH, densidade PSA
3. **TC Tórax** - Nódulos pulmonares, enfisema, rastreamento de câncer

**Questão Central:** Quais componentes são **quantitativos/semi-quantitativos** suficientes para integração no CSV de estratificação de risco vs. quais devem ser documentados como **diretrizes qualitativas de laudos**?

---

## 1. USG Abdome Total

### Terminologia Atualizada (2023-2026)

**MASLD** (Metabolic dysfunction-associated steatotic liver disease) substituiu **NAFLD** (nonalcoholic fatty liver disease) como terminologia preferencial, estabelecida por consenso em 2023 e formalmente adotada nas diretrizes clínicas EASL-EASD-EASO de 2024.

### 1.1. Esteatose Hepática (Hepatic Steatosis)

**Status:** ✅ **ADEQUADO PARA CSV** - Grau Semi-quantitativo

#### Classificação por Ultrassonografia B-Mode

| Grau | Achado Ultrassonográfico | % Hepatócitos |
|------|-------------------------|---------------|
| **Grau 0** | Ecogenicidade hepática normal | 0% (sem esteatose) |
| **Grau I** | Aumento difuso da ecogenicidade hepática, mas ecogenicidade periportal e diafragmática ainda visíveis | 5-33% |
| **Grau II** | Aumento difuso da ecogenicidade hepática obscurecendo ecogenicidade periportal, mas diafragmática ainda visível | 34-66% |
| **Grau III** | Aumento difuso da ecogenicidade hepática obscurecendo ecogenicidade periportal e diafragmática | >66% |

#### Estratificação de Risco para CSV

```
Esteatose Hepática (Grau USG)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 |
| Grau III (>66%) | Grau II (34-66%) | Grau I (5-33%) | Grau 0 (Normal) |
Grau USG | Correlação com histologia r=0.82 (2024)
```

#### Métodos Quantitativos Avançados

**CAP (Controlled Attenuation Parameter)** - Disponível em elastografia transitória:

```
Esteatose Hepática (CAP)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 |
| >290 dB/m | 260-290 dB/m | 238-260 dB/m | <238 dB/m |
dB/m | CAP ≥275 dB/m detecta esteatose grau ≥1 (AASLD 2024)
```

**Correspondência CAP com Grau:**
- 238-260 dB/m = S1 (Grau I) = 11-33%
- 260-290 dB/m = S2 (Grau II) = 34-66%
- >290 dB/m = S3 (Grau III) = ≥67%

### 1.2. Estratificação de Risco MASLD (Abordagem 2024-2025)

**Duas Etapas de Avaliação:**

#### Etapa 1: Avaliação Primária - FIB-4 Score

**FIB-4** (Fibrosis-4) é recomendado como teste sorológico inicial com alto valor preditivo negativo para excluir fibrose avançada (AASLD 2024).

```
FIB-4 Score (Estratificação Primária)
| Nível 0 | Nível 1 | Nível 2 |
| >2.67 (idade 36-65) ou >2.0 (idade >65) | 1.3-2.67 (36-65) ou 2.0-2.67 (>65) | <1.3 (idade 36-65) ou <2.0 (idade >65) |
Índice | FIB-4 = (Idade × AST) / (Plaquetas × √ALT)
```

**Interpretação:**
- **Nível 0 (Alto Risco):** Fibrose avançada provável → Encaminhar hepatologia
- **Nível 1 (Risco Intermediário):** Avaliação secundária necessária (VCTE ou ELF)
- **Nível 2 (Baixo Risco):** Seguimento em atenção primária

**Evidência:** Estudo 2024 mostrou que estratificação com FIB-4 caracterizou 77% como baixo risco, 17% indeterminado, e 6% alto risco.

#### Etapa 2: Avaliação Secundária (se FIB-4 elevado)

**VCTE-LSM** (Vibration-Controlled Transient Elastography - Liver Stiffness Measurement):
- LSM <8 kPa: Fibrose improvável
- LSM 8-12 kPa: Fibrose significativa possível
- LSM >12 kPa: Fibrose avançada/cirrose provável

### 1.3. Outros Achados de USG Abdome Total

**Status:** ❌ **NÃO ADEQUADO PARA CSV** - Reportar qualitativamente

Achados que devem ser documentados como presença/ausência ou descrição textual:

- **Colelitíase (Cálculos Biliares):** Presente/Ausente + características (tamanho, número, mobilidade)
- **Nefrolitíase (Cálculos Renais):** Presente/Ausente + localização, tamanho
- **Lesões Hepáticas Focais:** Presente/Ausente + características, necessita correlação
- **Esplenomegalia:** Presente/Ausente + medida do diâmetro (normal <13 cm)
- **Dilatação de Vias Biliares:** Presente/Ausente + calibre
- **Cistos Renais:** Classificação de Bosniak (I-IV)

### 1.4. Perspectiva de Medicina Funcional - MASLD

#### Reversibilidade e Progressão

**Evidência 2024-2025:**
- MASLD afeta ~25-30% da população adulta mundial
- Projeção 2040: Prevalência aumentará para >55%
- **Intervenções de estilo de vida** (dieta + atividade física) permanecem pedra angular do tratamento

#### Medicações FDA-Aprovadas (2024-2025)

**Resmetirom (Março 2024):**
- Primeiro medicamento FDA-aprovado para MASH não-cirrótico com fibrose moderada/avançada
- Agonista do receptor de hormônio tireoidiano-β

**Semaglutida (Agosto 2025):**
- Agonista GLP-1 aprovado para tratamento de MASH com fibrose moderada/avançada
- Demonstrou reduções significativas no conteúdo de gordura hepática

#### Abordagem Integrativa

1. **Perda de Peso:** 7-10% de redução pode reverter esteatose e melhorar fibrose
2. **Nutrição:** Dieta mediterrânea, redução de frutose, aumento de ômega-3
3. **Exercício:** 150-200 min/semana de atividade moderada
4. **Sensibilizadores de Insulina:** Pioglitazona mostrou melhora histológica
5. **Inibidores SGLT2:** Empagliflozina, dapagliflozina - melhoram sensibilidade insulínica
6. **Suplementação:** Vitamina E (800 IU/dia) em não-diabéticos com MASH

---

## 2. USG Próstata (via Abdominal)

### 2.1. Volume Prostático

**Status:** ✅ **ADEQUADO PARA CSV** - Medida Quantitativa

#### Cálculo do Volume

**Fórmula do Elipsoide:**
```
Volume (cc) = Comprimento × Largura × Altura × 0.52
```

#### Classificação por Volume

```
Volume Prostático (USG Transabdominal)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 |
| >80 cc | 50-80 cc | 30-50 cc | <30 cc |
cc (mL) | Volume normal: <30 cc
```

**Correlação com BPH (Benign Prostatic Hyperplasia):**
- <30 cc: Normal
- 30-50 cc: Aumento leve
- 50-80 cc: Aumento moderado
- >80 cc: Aumento severo

**Correlação com Sintomas (IPSS - International Prostate Symptom Score):**
- Volumes maiores geralmente correlacionam com scores IPSS mais altos
- Necessidade de avaliação clínica complementar (não apenas volume)

### 2.2. Densidade PSA (PSA Density)

**Status:** ✅ **ADEQUADO PARA CSV** - Parâmetro Calculado

#### Cálculo

```
PSAD = PSA Total (ng/mL) / Volume Prostático (cc)
```

#### Estratificação de Risco para Câncer de Próstata

```
Densidade PSA (PSAD)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 | Nível 4 |
| >0.20 ng/mL/cc | 0.15-0.20 ng/mL/cc | 0.10-0.15 ng/mL/cc | <0.10 ng/mL/cc | N/A |
ng/mL/cc | PSAD ≥0.15: considerar biópsia (guidelines 2024-2025)
```

**Interpretação Clínica (2024-2025):**
- **<0.10:** Baixo risco
- **0.10-0.15:** Risco intermediário
- **0.15-0.20:** Alto risco
- **>0.20:** Risco muito alto

**Evidência Recente (2025):**
- **4K Density:** Novo biomarcador ajustado por volume que supera PSAD e 4Kscore em detecção de câncer clinicamente significativo
- **Cutoff PSAD 0.30:** Melhor opção para pacientes com BPH e PSA elevado mas RM negativa (especialmente em próstatas grandes)

### 2.3. Acurácia da USG Transabdominal (2024-2025)

#### Comparação TAUS vs. TRUS vs. RM

**Estudo Prospectivo (2025):**
- Diferença média TAUS vs. RM: **2.5 mL** (SD: 16.4)
- Diferença média TAUS vs. TRUS: **11.5 mL** (SD: 20.4)
- Diferença média TRUS vs. RM: **-9.0 mL** (SD: 21.1)

**Conclusão:** TAUS é equivalente à RM para estimativa de volume prostático.

#### Vantagens TAUS

- Não invasivo, mais confortável para paciente
- Acessível, pode ser realizado na avaliação inicial
- Reduz superestimação de risco em 6% comparado a TRUS (7% vs. 13%)

#### Limitações TAUS

- Qualidade de imagem menor que TRUS/RM
- Dependente da experiência do operador
- Variabilidade interobservador necessita exploração adicional
- **Próstatas >30 cc:** Para próstatas maiores onde tamanho pode alterar manejo cirúrgico, TRUS ou RM recomendados

### 2.4. Outros Achados USG Próstata

**Status:** ❌ **NÃO ADEQUADO PARA CSV** - Reportar qualitativamente

- **Ecotextura:** Homogênea/Heterogênea
- **Calcificações:** Presente/Ausente + distribuição
- **Lesões Focais:** Presente/Ausente + descrição (necessita TRUS ou RM multiparamétrica)
- **Resíduo Pós-Miccional:** Medida em mL (quando aplicável)

### 2.5. Limitações para Detecção de Câncer

**IMPORTANTE:** USG transabdominal **NÃO** é adequada para detecção ou caracterização de câncer de próstata:

- Baixa sensibilidade para lesões pequenas
- Incapaz de diferenciar BPH de câncer
- **RM Multiparamétrica (mpMRI):** Padrão-ouro para avaliação de lesões suspeitas
- **TRUS com Biópsia:** Necessário para confirmação histológica

**Orientação Clínica:**
- PSA elevado + PSAD alto → mpMRI seguida de biópsia guiada por RM se PI-RADS ≥3
- Diretrizes EAU/AUA 2025 endossam PSAD como adjuvante ao PSA em decisões de biópsia

---

## 3. TC Tórax (Chest CT)

### 3.1. Nódulos Pulmonares

**Status:** ⚠️ **PARCIALMENTE ADEQUADO PARA CSV** - Categorização por Tamanho

#### Diretrizes Fleischner Society (2017 - Ainda Vigentes em 2024-2025)

**Aplicabilidade:**
- Nódulos descobertos **FORA** de programas de rastreamento de câncer de pulmão
- **NÃO** aplicável para: rastreamento, pacientes <35 anos, história de câncer primário, imunossupressão

**Objetivo:** Limitar avaliações adicionais quando probabilidade de câncer <1%; prosseguir quando ≥1%

#### Nódulos Sólidos - Estratificação por Tamanho e Risco

**Nódulo Solitário:**

```
Nódulo Pulmonar Sólido Solitário (Tamanho)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 |
| >8 mm | 6-8 mm | <6 mm (baixo risco) | Ausente |
mm | Seguimento conforme Fleischner 2017
```

**Seguimento Recomendado (Baixo Risco):**
- **<6 mm:** Sem seguimento necessário (consenso clínico: risco <1%)
- **6-8 mm:** CT em 6-12 meses, depois considerar 18-24 meses
- **>8 mm:** CT em 3 meses, PET/CT, ou biópsia

**Múltiplos Nódulos:**
- Seguimento baseado no nódulo de maior tamanho e características
- Maior número de nódulos pode indicar risco aumentado ou doença benigna difusa (necessita correlação clínica)

#### Nódulos em Vidro Fosco (Ground-Glass Nodules)

**Status:** ❌ **NÃO ADEQUADO PARA CSV** - Requer avaliação qualitativa especializada

- Maior risco de malignidade que nódulos sólidos de mesmo tamanho
- Crescimento mais lento
- Seguimento prolongado frequentemente necessário (3-5 anos)
- Requer avaliação por radiologista e pneumologista

#### Rastreamento de Câncer de Pulmão (USPSTF Grade B)

**Elegibilidade (2024):**
- Idade 50-80 anos
- História de tabagismo ≥20 pack-years
- Fumante atual ou que parou nos últimos 15 anos

**TC de Baixa Dose (LDCT):**
- Reduz mortalidade por câncer de pulmão em ~20%
- Rastreamento anual recomendado

**Diretrizes Japonesas (6ª Edição - 2024):**
- Conformam-se ao Fleischner Society
- Utilizam diâmetro médio de 6 mm como limiar

### 3.2. Enfisema

**Status:** ⚠️ **PARCIALMENTE ADEQUADO PARA CSV** - Grau Semi-quantitativo

#### Goddard Score - Scoring Visual

**Metodologia:**
- Pulmão dividido em 6 áreas (superior, médio, inferior bilateralmente)
- Cada área pontuada 0-4 baseado em extensão de enfisema
- Pontuação total: 0-24 (soma das 6 áreas)

**Pontuação por Área:**
- **0 pontos:** Sem alterações enfisematosas
- **1 ponto:** <25% de enfisema
- **2 pontos:** 25-50% de enfisema
- **3 pontos:** 50-75% de enfisema
- **4 pontos:** ≥75% de enfisema

```
Enfisema (Goddard Score)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 | Nível 4 |
| Score 13-24 | Score 7-12 | Score 1-6 | Score 0 | N/A |
Score 0-24 | Goddard Score - avaliação visual semi-quantitativa
```

**Correlação Clínica:**
- **0 pontos:** Sem enfisema
- **1-6 pontos:** Enfisema leve
- **7-12 pontos:** Enfisema moderado
- **13-24 pontos:** Enfisema severo

**Vantagens do Goddard Score (2024-2025):**
- Método visual simples, amplamente usado na prática clínica
- Não requer software especializado
- Valioso para avaliações de emergência, estratificação pré-operatória, monitoramento longitudinal
- Confiável mesmo em TC com dose equivalente a raio-X (PCD-CT 2025)

#### Métodos Quantitativos Avançados

**LAA% (Low Attenuation Area):**
- Percentual de áreas com atenuação <-950 HU
- Requer software especializado
- Mais objetivo que Goddard Score

```
Enfisema (LAA%-950)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 |
| >25% | 10-25% | 1-10% | <1% |
% | Percentual de área <-950 HU
```

**Machine Learning (2024-2025):**
- Modelos integram parênquima pulmonar, vias aéreas, vasos
- Swin Transformer + características de enfisema: AUC 84.3% (COPDGene)
- ePRM (elastic Parametric Response Mapping): Classifica volumes por grau de doença de pequenas vias aéreas, normal, e tecido enfisematoso

#### Correlação com Severidade COPD (GOLD)

**Classificação GOLD (Global Initiative for Chronic Obstructive Lung Disease):**
- Baseada em VEF1 (Volume Expiratório Forçado em 1 segundo)
- TC quantitativa correlaciona bem com GOLD, mas não substitui espirometria

**Aplicação Clínica:**
- Estratificação pré-operatória (risco de complicações, vazamento aéreo)
- Monitoramento de progressão
- Guiar estratégias de manejo personalizadas

### 3.3. Outros Achados TC Tórax

**Status:** ❌ **NÃO ADEQUADO PARA CSV** - Reportar qualitativamente

Achados que requerem descrição detalhada em laudo:

- **Opacidades em Vidro Fosco (GGO):** Podem indicar infecção, inflamação, neoplasia, hemorragia
- **Fibrose Pulmonar:** Classificação UIP (Usual Interstitial Pneumonia) pattern, NSIP, etc.
- **Bronquiectasias:** Presente/Ausente + distribuição, severidade
- **Derrame Pleural:** Presente/Ausente + volume estimado
- **Linfadenopatia Mediastinal/Hilar:** Presente/Ausente + tamanho, características
- **Consolidações:** Presente/Ausente + distribuição, padrão
- **Atelectasia:** Presente/Ausente + tipo, extensão
- **Achados Cardíacos:** Cardiomegalia, calcificação coronária (CAC score se protocolo específico)

---

## 4. Resumo Executivo - Integração CSV

### 4.1. Componentes ADEQUADOS para CSV

| Exame | Componente | Tipo | Colunas CSV |
|-------|-----------|------|-------------|
| **USG Abdome** | Esteatose Hepática (Grau) | Semi-quantitativo | `hepatic_steatosis_grade` (0, I, II, III) |
| **USG Abdome** | CAP (se disponível) | Quantitativo | `cap_score_dbm` (dB/m) |
| **USG Abdome** | FIB-4 Score | Quantitativo (calculado) | `fib4_score` (índice) |
| **USG Próstata** | Volume Prostático | Quantitativo | `prostate_volume_cc` (cc) |
| **USG Próstata** | Densidade PSA | Quantitativo (calculado) | `psa_density` (ng/mL/cc) |
| **TC Tórax** | Nódulo Sólido Maior | Semi-quantitativo | `largest_solid_nodule_mm` (mm) |
| **TC Tórax** | Goddard Score | Semi-quantitativo | `emphysema_goddard_score` (0-24) |
| **TC Tórax** | LAA%-950 (se disponível) | Quantitativo | `emphysema_laa_percent` (%) |

### 4.2. Componentes NÃO ADEQUADOS para CSV (Laudo Qualitativo)

**USG Abdome Total:**
- Colelitíase, nefrolitíase, lesões hepáticas focais, esplenomegalia, cistos renais

**USG Próstata:**
- Ecotextura, calcificações, lesões focais, resíduo pós-miccional

**TC Tórax:**
- Nódulos em vidro fosco, fibrose pulmonar, bronquiectasias, derrame pleural, linfadenopatia, consolidações

**Recomendação:** Esses achados devem ser documentados em campos de texto livre (`report_text`, `additional_findings`) no banco de dados EMR.

### 4.3. Decisão de Suporte Clínico (CDS)

#### Algoritmo MASLD (USG Abdome)

```
SE esteatose_hepatica_grade IN (I, II, III)
  E fib4_score IS NULL
  ENTÃO: ALERTAR "Calcular FIB-4 Score para estratificação de fibrose"

SE fib4_score >= 1.3 (idade 36-65) OU >= 2.0 (idade >65)
  ENTÃO: RECOMENDAR "Elastografia transitória (VCTE) ou ELF para avaliação secundária"

SE fib4_score > 2.67 (idade 36-65) OU > 2.67 (idade >65)
  ENTÃO: ALERTAR "Risco alto de fibrose avançada - Encaminhar hepatologia"

SE esteatose_hepatica_grade = III
  ENTÃO: RECOMENDAR "Intervenção estilo de vida: perda peso 7-10%, dieta mediterrânea, exercício 150-200 min/sem"
```

#### Algoritmo Densidade PSA (USG Próstata)

```
SE psa_density >= 0.15
  ENTÃO: RECOMENDAR "Considerar RM multiparamétrica próstata (mpMRI)"

SE psa_density >= 0.20
  ENTÃO: ALERTAR "Risco muito alto câncer próstata - Avaliar mpMRI + biópsia"

SE prostate_volume_cc > 80 E psa_density >= 0.30
  ENTÃO: RECOMENDAR "BPH severo com PSAD elevado - Considerar biópsia mesmo com RM negativa"
```

#### Algoritmo Nódulo Pulmonar (TC Tórax)

```
SE largest_solid_nodule_mm > 8
  ENTÃO: ALERTAR "Nódulo >8mm - CT follow-up 3 meses, considerar PET/CT ou biópsia"

SE largest_solid_nodule_mm >= 6 E largest_solid_nodule_mm <= 8
  ENTÃO: RECOMENDAR "CT follow-up 6-12 meses, depois considerar 18-24 meses (Fleischner 2017)"

SE largest_solid_nodule_mm < 6
  ENTÃO: INFO "Nódulo <6mm baixo risco - Seguimento não necessário (consenso clínico)"
```

#### Algoritmo Enfisema (TC Tórax)

```
SE emphysema_goddard_score >= 13
  ENTÃO: ALERTAR "Enfisema severo - Avaliar espirometria, encaminhar pneumologia"

SE emphysema_goddard_score >= 7 E emphysema_goddard_score < 13
  ENTÃO: RECOMENDAR "Enfisema moderado - Espirometria, cessação tabagismo, broncodilatadores"

SE emphysema_goddard_score >= 1 E emphysema_goddard_score < 7
  ENTÃO: INFO "Enfisema leve - Cessação tabagismo, seguimento clínico"
```

---

## 5. Implementação no Backend Go

### 5.1. Models (Go)

```go
// apps/api/internal/models/imaging_exam.go

type ImagingExam struct {
    ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    PatientID   uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId" validate:"required"`
    ExamType    string    `gorm:"type:varchar(50);not null;check:exam_type IN ('usg_abdomen','usg_prostate','ct_chest')" json:"examType"`
    ExamDate    time.Time `gorm:"type:date;not null" json:"examDate" validate:"required"`

    // USG Abdome - Esteatose Hepática
    HepaticSteatosisGrade *string  `gorm:"type:varchar(10);check:hepatic_steatosis_grade IN ('0','I','II','III')" json:"hepaticSteatosisGrade,omitempty"`
    CAPScore              *float64 `gorm:"type:decimal(5,1)" json:"capScore,omitempty"` // dB/m

    // USG Próstata
    ProstateVolume *float64 `gorm:"type:decimal(6,2)" json:"prostateVolume,omitempty"` // cc
    PSADensity     *float64 `gorm:"type:decimal(5,3)" json:"psaDensity,omitempty"` // ng/mL/cc (calculado)

    // TC Tórax - Nódulos
    LargestNoduleSize *float64 `gorm:"type:decimal(5,2)" json:"largestNoduleSize,omitempty"` // mm

    // TC Tórax - Enfisema
    EmphysemaGoddardScore *int     `gorm:"type:smallint;check:emphysema_goddard_score BETWEEN 0 AND 24" json:"emphysemaGoddardScore,omitempty"`
    EmphysemaLAAPercent   *float64 `gorm:"type:decimal(5,2)" json:"emphysemaLaaPercent,omitempty"` // %

    // Laudo Textual Qualitativo
    ReportText         string  `gorm:"type:text" json:"reportText"`
    AdditionalFindings *string `gorm:"type:text" json:"additionalFindings,omitempty"`

    // Radiologista
    RadiologistID *uuid.UUID `gorm:"type:uuid" json:"radiologistId,omitempty"`

    CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Calcular Densidade PSA
func (ie *ImagingExam) CalculatePSADensity(psaTotal float64) {
    if ie.ProstateVolume != nil && *ie.ProstateVolume > 0 {
        density := psaTotal / *ie.ProstateVolume
        ie.PSADensity = &density
    }
}

// Estratificação de Risco - Esteatose
func (ie *ImagingExam) GetSteatosisRiskLevel() int {
    if ie.HepaticSteatosisGrade == nil {
        return -1 // N/A
    }
    switch *ie.HepaticSteatosisGrade {
    case "III":
        return 0 // Pior
    case "II":
        return 1
    case "I":
        return 2
    case "0":
        return 3 // Melhor
    default:
        return -1
    }
}

// Estratificação de Risco - Volume Prostático
func (ie *ImagingExam) GetProstateVolumeRiskLevel() int {
    if ie.ProstateVolume == nil {
        return -1
    }
    vol := *ie.ProstateVolume
    if vol > 80 {
        return 0 // Severo
    } else if vol >= 50 {
        return 1 // Moderado
    } else if vol >= 30 {
        return 2 // Leve
    } else {
        return 3 // Normal
    }
}

// Estratificação de Risco - Densidade PSA
func (ie *ImagingExam) GetPSADensityRiskLevel() int {
    if ie.PSADensity == nil {
        return -1
    }
    psad := *ie.PSADensity
    if psad > 0.20 {
        return 0 // Muito alto
    } else if psad >= 0.15 {
        return 1 // Alto
    } else if psad >= 0.10 {
        return 2 // Intermediário
    } else {
        return 3 // Baixo
    }
}

// Estratificação de Risco - Nódulo Pulmonar
func (ie *ImagingExam) GetNoduleRiskLevel() int {
    if ie.LargestNoduleSize == nil {
        return 3 // Sem nódulo = melhor
    }
    size := *ie.LargestNoduleSize
    if size > 8 {
        return 0 // Alto risco
    } else if size >= 6 {
        return 1 // Risco intermediário
    } else {
        return 2 // Baixo risco
    }
}

// Estratificação de Risco - Enfisema
func (ie *ImagingExam) GetEmphysemaRiskLevel() int {
    if ie.EmphysemaGoddardScore == nil {
        return 3 // Sem enfisema = melhor
    }
    score := *ie.EmphysemaGoddardScore
    if score >= 13 {
        return 0 // Severo
    } else if score >= 7 {
        return 1 // Moderado
    } else if score >= 1 {
        return 2 // Leve
    } else {
        return 3 // Normal
    }
}
```

### 5.2. Clinical Decision Support (CDS)

```go
// apps/api/internal/services/cds_imaging.go

type ImagingCDSAlert struct {
    Severity    string `json:"severity"` // "info", "warning", "critical"
    Message     string `json:"message"`
    Recommendation string `json:"recommendation,omitempty"`
}

func GenerateImagingCDSAlerts(exam *models.ImagingExam, patient *models.Patient, labResults *models.LabResults) []ImagingCDSAlert {
    alerts := []ImagingCDSAlert{}

    // MASLD/NAFLD Alerts
    if exam.ExamType == "usg_abdomen" && exam.HepaticSteatosisGrade != nil {
        grade := *exam.HepaticSteatosisGrade
        if grade == "I" || grade == "II" || grade == "III" {
            // Calcular FIB-4 se dados disponíveis
            if labResults != nil && labResults.AST != nil && labResults.ALT != nil && labResults.Platelets != nil {
                age := calculateAge(patient.BirthDate)
                fib4 := calculateFIB4(age, *labResults.AST, *labResults.ALT, *labResults.Platelets)

                // Estratificar por idade
                if age >= 36 && age <= 65 {
                    if fib4 > 2.67 {
                        alerts = append(alerts, ImagingCDSAlert{
                            Severity: "critical",
                            Message: fmt.Sprintf("FIB-4 Score: %.2f (>2.67) - Alto risco de fibrose avançada", fib4),
                            Recommendation: "Encaminhar para hepatologia com urgência",
                        })
                    } else if fib4 >= 1.3 {
                        alerts = append(alerts, ImagingCDSAlert{
                            Severity: "warning",
                            Message: fmt.Sprintf("FIB-4 Score: %.2f (1.3-2.67) - Risco intermediário", fib4),
                            Recommendation: "Solicitar elastografia transitória (VCTE) ou ELF para avaliação secundária",
                        })
                    }
                } else if age > 65 {
                    if fib4 > 2.67 {
                        alerts = append(alerts, ImagingCDSAlert{
                            Severity: "critical",
                            Message: fmt.Sprintf("FIB-4 Score: %.2f (>2.67) - Alto risco", fib4),
                            Recommendation: "Encaminhar para hepatologia",
                        })
                    } else if fib4 >= 2.0 {
                        alerts = append(alerts, ImagingCDSAlert{
                            Severity: "warning",
                            Message: fmt.Sprintf("FIB-4 Score: %.2f (≥2.0) - Avaliar fibrose", fib4),
                            Recommendation: "VCTE ou ELF para avaliação secundária",
                        })
                    }
                }
            } else {
                alerts = append(alerts, ImagingCDSAlert{
                    Severity: "info",
                    Message: "Esteatose hepática detectada",
                    Recommendation: "Solicitar AST, ALT, contagem de plaquetas para cálculo de FIB-4 Score",
                })
            }

            // Intervenção de estilo de vida
            if grade == "II" || grade == "III" {
                alerts = append(alerts, ImagingCDSAlert{
                    Severity: "warning",
                    Message: fmt.Sprintf("Esteatose hepática Grau %s - Intervenção necessária", grade),
                    Recommendation: "Perda de peso 7-10%, dieta mediterrânea, exercício 150-200 min/semana, considerar GLP-1 agonistas se diabético",
                })
            }
        }
    }

    // PSA Density Alerts
    if exam.ExamType == "usg_prostate" && exam.PSADensity != nil {
        psad := *exam.PSADensity
        if psad >= 0.20 {
            alerts = append(alerts, ImagingCDSAlert{
                Severity: "critical",
                Message: fmt.Sprintf("Densidade PSA: %.3f ng/mL/cc - Risco muito alto", psad),
                Recommendation: "Urgente: RM multiparamétrica próstata + avaliação urologia para biópsia",
            })
        } else if psad >= 0.15 {
            alerts = append(alerts, ImagingCDSAlert{
                Severity: "warning",
                Message: fmt.Sprintf("Densidade PSA: %.3f ng/mL/cc - Alto risco", psad),
                Recommendation: "Solicitar RM multiparamétrica próstata para avaliação de lesões",
            })
        } else if psad >= 0.10 {
            alerts = append(alerts, ImagingCDSAlert{
                Severity: "info",
                Message: fmt.Sprintf("Densidade PSA: %.3f ng/mL/cc - Risco intermediário", psad),
                Recommendation: "Seguimento clínico rigoroso, considerar mpMRI se PSA em elevação",
            })
        }
    }

    // Prostate Volume + PSAD Special Case
    if exam.ExamType == "usg_prostate" && exam.ProstateVolume != nil && exam.PSADensity != nil {
        if *exam.ProstateVolume > 80 && *exam.PSADensity >= 0.30 {
            alerts = append(alerts, ImagingCDSAlert{
                Severity: "warning",
                Message: "BPH severo com PSAD ≥0.30",
                Recommendation: "Considerar biópsia mesmo com RM negativa (evidência 2025)",
            })
        }
    }

    // Pulmonary Nodule Alerts
    if exam.ExamType == "ct_chest" && exam.LargestNoduleSize != nil {
        size := *exam.LargestNoduleSize
        if size > 8 {
            alerts = append(alerts, ImagingCDSAlert{
                Severity: "critical",
                Message: fmt.Sprintf("Nódulo pulmonar sólido: %.1f mm - Alto risco", size),
                Recommendation: "CT follow-up 3 meses, considerar PET/CT ou biópsia. Encaminhar pneumologia/oncologia",
            })
        } else if size >= 6 {
            alerts = append(alerts, ImagingCDSAlert{
                Severity: "warning",
                Message: fmt.Sprintf("Nódulo pulmonar sólido: %.1f mm", size),
                Recommendation: "CT follow-up 6-12 meses, depois 18-24 meses (Fleischner 2017)",
            })
        } else if size > 0 {
            alerts = append(alerts, ImagingCDSAlert{
                Severity: "info",
                Message: fmt.Sprintf("Nódulo pulmonar sólido: %.1f mm - Baixo risco", size),
                Recommendation: "Sem seguimento necessário (<6mm consenso clínico). Reavaliar se sintomas novos",
            })
        }
    }

    // Emphysema Alerts
    if exam.ExamType == "ct_chest" && exam.EmphysemaGoddardScore != nil {
        score := *exam.EmphysemaGoddardScore
        if score >= 13 {
            alerts = append(alerts, ImagingCDSAlert{
                Severity: "critical",
                Message: fmt.Sprintf("Enfisema severo (Goddard: %d/24)", score),
                Recommendation: "Solicitar espirometria, encaminhar pneumologia. Avaliação GOLD COPD. Cessação tabagismo crítica",
            })
        } else if score >= 7 {
            alerts = append(alerts, ImagingCDSAlert{
                Severity: "warning",
                Message: fmt.Sprintf("Enfisema moderado (Goddard: %d/24)", score),
                Recommendation: "Espirometria, cessação tabagismo, considerar broncodilatadores, vacinação influenza/pneumococo",
            })
        } else if score >= 1 {
            alerts = append(alerts, ImagingCDSAlert{
                Severity: "info",
                Message: fmt.Sprintf("Enfisema leve (Goddard: %d/24)", score),
                Recommendation: "Cessação tabagismo imediata. Seguimento clínico anual com espirometria",
            })
        }
    }

    return alerts
}

func calculateFIB4(age int, ast, alt, platelets float64) float64 {
    // FIB-4 = (Age × AST) / (Platelets × √ALT)
    // Plaquetas em 10^9/L
    return (float64(age) * ast) / (platelets * math.Sqrt(alt))
}
```

---

## 6. Fontes e Evidências

### USG Abdome Total / MASLD

- [Current Update on MASLD Nomenclature and Management - RadioGraphics (RSNA)](https://pubs.rsna.org/doi/10.1148/rg.240221)
- [Hepatic Steatosis: Sonographic Assessment vs Histology - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC11423484/)
- [Global Consensus Recommendations for MASLD - Gastroenterology](https://www.gastrojournal.org/article/S0016-5085(25)00632-8/fulltext)
- [MASLD StatPearls - NCBI Bookshelf](https://www.ncbi.nlm.nih.gov/books/NBK541033/)
- [Noninvasive Assessment of MASLD - AASLD](https://www.aasld.org/liver-fellow-network/core-series/clinical-pearls/spare-me-jab-noninvasive-assessment-patients-masld)
- [MAFLD Guideline 2024 - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC11557364/)
- [Diffuse Hepatic Steatosis Grading - Radiopaedia](https://radiopaedia.org/articles/diffuse-hepatic-steatosis-grading?lang=us)
- [Imaging Biomarkers in MASLD - Current Hepatology Reports](https://link.springer.com/article/10.1007/s11901-025-00711-9)
- [MASLD Treatment Mechanisms and Advances - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC12627968/)
- [Semaglutide Therapy for MASH - AASLD 2025 - PubMed](https://pubmed.ncbi.nlm.nih.gov/41201884/)
- [KASL Clinical Practice Guidelines MASLD 2025 - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC11925433/)
- [MASLD in People With Diabetes - Diabetes Care (ADA)](https://diabetesjournals.org/care/article/48/7/1057/160536/)

### FIB-4 Score

- [Why FIB-4 Used in Clinical Practice - AASLD](https://www.aasld.org/liver-fellow-network/core-series/why-series/why-are-non-invasive-risk-scores-such-fib-4-used)
- [Age-dependent FIB-4 in MASLD - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC11637747/)
- [Validation FIB-4 for Liver Cirrhosis - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC12269257/)
- [Identifying Patients with MASLD Advanced Fibrosis - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC11861828/)
- [FIB-4 Calculator - MDCalc](https://www.mdcalc.com/calc/2200/fibrosis-4-fib-4-index-liver-fibrosis)
- [FIB-4 Despite Normal Aminotransferases - ACP Gastroenterology](https://gastroenterology.acponline.org/archives/2024/08/23/2.htm)

### USG Próstata / PSA Density

- [4K Density for Risk Stratification - Urology Today](https://www.urotoday.com/index.php?option=com_content&view=article&id=164263)
- [4K Density Prostate Cancer - PubMed](https://pubmed.ncbi.nlm.nih.gov/41107473/)
- [Prostate Volume TAUS vs MRI vs TRUS 2025 - The Prostate (Wiley)](https://onlinelibrary.wiley.com/doi/10.1002/pros.70036)
- [Prostate Volume Calculator - RadAtHand](https://radathand.com/radiology-calculators/body-imaging/prostate-volume-calculator-psa-density-calculator/)
- [Optimal PSAD Threshold BPH - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC11874838/)
- [Optimal PSAD Threshold BPH - BMC Urology](https://bmcurol.biomedcentral.com/articles/10.1186/s12894-025-01719-5)
- [EAU Guidelines Prostate Cancer 2025 - PDF](https://d56bochluxqnz.cloudfront.net/documents/full-guideline/EAU-EANM-ESTRO-ESUR-ISUP-SIOG-Guidelines-on-Prostate-Cancer-2025_updated.pdf)
- [EAU Guidelines Prostate Cancer - Uroweb](https://uroweb.org/guidelines/prostate-cancer)
- [Risk Stratification MRI PSA Density - British Journal of Radiology](https://academic.oup.com/bjr/article/97/1153/113/7470409)

### TC Tórax / Nódulos Pulmonares

- [Japanese Guidelines Pulmonary Nodules 6th Edition 2024 - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC11868311/)
- [Fleischner Society Recommendations - Radiopaedia](https://radiopaedia.org/articles/fleischner-society-pulmonary-nodule-recommendations-1?lang=us)
- [Pulmonary Nodules Q&A - AAFP](https://www.aafp.org/pubs/afp/issues/2023/0300/pulmonary-nodules.html)
- [Follow-Up Lung Nodules - STS](https://www.sts.org/follow-lung-nodules)
- [Incidental Pulmonary Nodules Management - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC12094709/)
- [Fleischner Guidelines and Cancer Probability - JACR](https://www.jacr.org/article/S1546-1440(22)00564-6/abstract)
- [Fleischner Guidelines Calculator - MDCalc](https://www.mdcalc.com/calc/10062/fleischner-society-guidelines-incidental-pulmonary-nodules)
- [Japanese Guidelines LDCT Screening 2024 - Japanese Journal of Radiology](https://link.springer.com/article/10.1007/s11604-024-01695-0)

### TC Tórax / Enfisema COPD

- [COPD Severity Deep Learning - European Radiology](https://link.springer.com/article/10.1007/s00330-023-10540-3)
- [COPD Diagnosis Machine Learning - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC12360390/)
- [CT Emphysema Single Lobe - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC11250351/)
- [CT Visual Classification Emphysema Mortality - Radiology](https://pubs.rsna.org/doi/full/10.1148/radiol.2018172294)
- [CT Reconstruction Parameters Impact Goddard Score 2026 - International Journal Biomedical Imaging](https://onlinelibrary.wiley.com/doi/10.1155/ijbi/7436511)
- [COPD-TransNet Swin Transformer - Journal of Imaging Informatics in Medicine](https://link.springer.com/article/10.1007/s10278-025-01785-z)
- [Quantitative CT in COPD - PMC](https://pmc.ncbi.nlm.nih.gov/articles/PMC4161463/)
- [Diagnosis and COPD Severity Machine Learning - Dove Press](https://www.dovepress.com/diagnosis-and-severity-assessment-of-copd-based-on-machine-learning-of-peer-reviewed-fulltext-article-COPD)
- [Goddard Score Air Leak Prediction 2024 - PubMed](https://pubmed.ncbi.nlm.nih.gov/37858880/)
- [Goddard Score Air Leak - Annals of Thoracic Surgery](https://www.annalsthoracicsurgery.org/article/S0003-4975(23)01068-8/fulltext)
- [Emphysema X-ray Dose PCD-CT 2025 - PubMed](https://pubmed.ncbi.nlm.nih.gov/39729642/)

---

## 7. Conclusões e Próximos Passos

### 7.1. Resumo de Integração

**Componentes Quantitativos CSV-Ready:**
1. **USG Abdome:** Grau de esteatose (0/I/II/III), CAP score, FIB-4 calculado
2. **USG Próstata:** Volume prostático (cc), Densidade PSA (ng/mL/cc)
3. **TC Tórax:** Tamanho maior nódulo (mm), Goddard Score (0-24), LAA% (-950 HU)

**Total:** 8 campos quantitativos/semi-quantitativos para estratificação CSV

### 7.2. Decisão de Suporte Clínico Implementável

- **Alertas Automáticos:** 4 algoritmos prontos (MASLD, PSA Density, Nódulos, Enfisema)
- **Níveis de Severidade:** Info / Warning / Critical
- **Recomendações Acionáveis:** Baseadas em guidelines 2024-2025

### 7.3. Próximas Ações

1. **Backend:**
   - Criar migration para tabela `imaging_exams`
   - Implementar handlers CRUD
   - Implementar serviço CDS (alertas automáticos)
   - Adicionar cálculo automático de FIB-4 quando labs disponíveis
   - Adicionar cálculo automático de PSAD quando PSA disponível

2. **Frontend:**
   - Forms para cadastro de exames de imagem
   - Dashboard com visualização de alertas CDS
   - Cards de "Risk Stratification" mostrando níveis coloridos
   - Integração com timeline de paciente

3. **Documentação:**
   - Criar guia para radiologistas sobre campos CSV
   - Criar guia de interpretação de alertas CDS
   - Documentar limitações e quando usar laudo qualitativo

### 7.4. Limitações e Avisos

**IMPORTANTE:** Exames de imagem têm limitações inerentes:

- **USG Abdome:** Operador-dependente, limitado para fibrose (necessita VCTE/RM/biópsia)
- **USG Próstata Transabdominal:** Não detecta câncer (necessita mpMRI + biópsia)
- **TC Tórax:** Nódulos em vidro fosco não estratificados aqui (requerem avaliação especializada)

**Nunca substituir julgamento clínico:** Algoritmos CDS são ferramentas de suporte, não substitutos de decisão médica.

---

**Documento elaborado:** Janeiro 2026
**Última atualização evidências:** Janeiro 2026
**Próxima revisão:** Janeiro 2027 ou quando novas diretrizes publicadas
