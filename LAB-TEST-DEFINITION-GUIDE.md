# Guia: Lab Test Definition - Estrutura de Exames Laboratoriais

## 1. Visão Geral

O modelo `LabTestDefinition` cria uma estrutura hierárquica e padronizada para definir todos os exames laboratoriais do sistema.

### Componentes Principais

```
LabTestDefinition (definição do exame/parâmetro)
  ├─ LabTestScoreMapping (mapeamento para ScoreItem)
  └─ LabTestReferenceRange (faixas de referência)
```

---

## 2. Campos Principais

### Identificação

- **`code`**: Código único interno (ex: `HEMOGRAMA_COMPLETO`, `HGB_M`)
- **`name`**: Nome completo (ex: "Hemograma Completo", "Hemoglobina - Homens")
- **`shortName`**: Nome curto/abreviação (ex: "Hemograma", "Hb")
- **`tussCode`**: Código TUSS para faturamento (ex: "40304485")
- **`loincCode`**: Código LOINC internacional (ex: "718-7")

### Hierarquia

- **`isRequestable`**: Pode ser solicitado individualmente?
  - `true`: Exame completo que pode ser pedido (Hemograma, Glicemia)
  - `false`: Parâmetro que só vem no resultado (Hemoglobina, Bilirrubina Indireta)

- **`parentTestId`**: ID do exame pai (para criar hierarquia)

### Categorização

- **`category`**: Categoria do exame
  - `hematology`: Hemogramas, coagulação
  - `biochemistry`: Glicose, eletrólitos, função renal/hepática
  - `hormones`: TSH, T4, testosterona, cortisol
  - `immunology`: Sorologias, autoimunes
  - `microbiology`: Culturas, antibiogramas
  - `urine`: EAS, urocultura
  - `imaging`: Exames de imagem
  - `functional`: Medicina funcional
  - `genetics`: Testes genéticos
  - `other`: Outros

### Resultado

- **`resultType`**: Tipo de resultado
  - `numeric`: Valores numéricos (ex: 14.5)
  - `text`: Texto livre (ex: "Negativo", "Positivo")
  - `boolean`: Sim/Não (ex: presença de sangue)
  - `categorical`: Categorias fixas (ex: "Escasso", "Moderado", "Abundante")

- **`unit`**: Unidade padrão (ex: "g/dL", "mg/dL", "%")
- **`unitConversion`**: Conversão entre unidades (ex: "1 g/dL = 10 g/L")

### Coleta

- **`collectionMethod`**: Como coletar (ex: "Sangue venoso com jejum")
- **`fastingHours`**: Horas de jejum necessárias
- **`specimenType`**: Material (ex: "Sangue total", "Soro", "Urina")

---

## 3. Exemplos Práticos

### 3.1 Hemograma Completo (Exame Composto Solicitável)

```json
{
  "code": "HEMOGRAMA_COMPLETO",
  "name": "Hemograma Completo",
  "shortName": "Hemograma",
  "tussCode": "40304485",
  "loincCode": "58410-2",
  "category": "hematology",
  "isRequestable": true,
  "parentTestId": null,
  "resultType": "numeric",
  "specimenType": "Sangue total com EDTA",
  "collectionMethod": "Punção venosa",
  "fastingHours": 0,
  "description": "Avaliação quantitativa e qualitativa das células sanguíneas",
  "displayOrder": 0,
  "isActive": true
}
```

### 3.2 Hemoglobina - Homens (Parâmetro NÃO Solicitável)

```json
{
  "code": "HGB_M",
  "name": "Hemoglobina - Homens",
  "shortName": "Hb",
  "loincCode": "718-7",
  "category": "hematology",
  "isRequestable": false,           // ← NÃO pode ser solicitado sozinho
  "parentTestId": "<uuid-hemograma-completo>",  // ← Sempre vem no Hemograma
  "unit": "g/dL",
  "unitConversion": "1 g/dL = 10 g/L",
  "resultType": "numeric",
  "displayOrder": 1,
  "isActive": true
}
```

### 3.3 Hemoglobina - Mulheres (Parâmetro com Faixas Diferentes)

```json
{
  "code": "HGB_F",
  "name": "Hemoglobina - Mulheres",
  "shortName": "Hb",
  "loincCode": "718-7",
  "category": "hematology",
  "isRequestable": false,
  "parentTestId": "<uuid-hemograma-completo>",
  "unit": "g/dL",
  "unitConversion": "1 g/dL = 10 g/L",
  "resultType": "numeric",
  "displayOrder": 2,
  "isActive": true
}
```

### 3.4 Glicemia de Jejum (Exame Individual Solicitável)

```json
{
  "code": "GLUCOSE_FASTING",
  "name": "Glicemia de Jejum",
  "shortName": "Gli",
  "tussCode": "40301044",
  "loincCode": "1558-6",
  "category": "biochemistry",
  "isRequestable": true,            // ← Pode ser solicitado individualmente
  "parentTestId": null,
  "unit": "mg/dL",
  "unitConversion": "1 mg/dL = 0.0555 mmol/L",
  "resultType": "numeric",
  "specimenType": "Soro ou plasma",
  "fastingHours": 8,
  "collectionMethod": "Sangue venoso após jejum de 8-12 horas",
  "displayOrder": 0,
  "isActive": true
}
```

### 3.5 Bilirrubina Total e Frações (Exame Composto)

```json
{
  "code": "BILIRUBIN_TOTAL_FRACTIONS",
  "name": "Bilirrubina Total e Frações",
  "shortName": "Bili TF",
  "tussCode": "40301052",
  "category": "biochemistry",
  "isRequestable": true,
  "parentTestId": null,
  "resultType": "numeric",
  "specimenType": "Soro",
  "fastingHours": 4,
  "displayOrder": 0,
  "isActive": true
}
```

### 3.6 Bilirrubina Direta (Subparâmetro)

```json
{
  "code": "BILIRUBIN_DIRECT",
  "name": "Bilirrubina Direta",
  "shortName": "BD",
  "category": "biochemistry",
  "isRequestable": false,           // ← Vem no resultado, não no pedido
  "parentTestId": "<uuid-bili-total-fracoes>",
  "unit": "mg/dL",
  "resultType": "numeric",
  "displayOrder": 1,
  "isActive": true
}
```

### 3.7 Bilirrubina Indireta (Subparâmetro Calculado)

```json
{
  "code": "BILIRUBIN_INDIRECT",
  "name": "Bilirrubina Indireta",
  "shortName": "BI",
  "category": "biochemistry",
  "isRequestable": false,           // ← NUNCA é solicitada
  "parentTestId": "<uuid-bili-total-fracoes>",
  "unit": "mg/dL",
  "resultType": "numeric",
  "description": "Calculada: Total - Direta",
  "displayOrder": 2,
  "isActive": true
}
```

---

## 4. Mapeamento para Score Items

### 4.1 LabTestScoreMapping

Relaciona um `LabTestDefinition` com um `ScoreItem`, podendo especificar gênero e faixa etária.

**Exemplo: Hemoglobina Homens → ScoreItem**

```json
{
  "labTestId": "<uuid-HGB_M>",
  "scoreItemId": "<uuid-score-item-hemoglobina-homens>",
  "gender": "male",
  "minAge": 18,
  "maxAge": null,
  "notes": "Referência para homens adultos",
  "isActive": true
}
```

**Exemplo: Hemoglobina Mulheres → ScoreItem**

```json
{
  "labTestId": "<uuid-HGB_F>",
  "scoreItemId": "<uuid-score-item-hemoglobina-mulheres>",
  "gender": "female",
  "minAge": 18,
  "maxAge": null,
  "isActive": true
}
```

**Exemplo: Glicemia (sem distinção de gênero)**

```json
{
  "labTestId": "<uuid-GLUCOSE_FASTING>",
  "scoreItemId": "<uuid-score-item-glicemia>",
  "gender": null,          // ← Ambos os gêneros
  "minAge": null,
  "maxAge": null,
  "isActive": true
}
```

---

## 5. Faixas de Referência

### 5.1 LabTestReferenceRange

Define valores de referência que podem variar por gênero e idade.

**Exemplo: Hemoglobina - Homens Adultos**

```json
{
  "labTestId": "<uuid-HGB_M>",
  "gender": "male",
  "minAge": 18,
  "maxAge": null,
  "lowerLimit": 13.5,
  "upperLimit": 17.5,
  "unit": "g/dL",
  "description": "Faixa de referência para homens adultos",
  "source": "Diretrizes SBP 2024",
  "isActive": true
}
```

**Exemplo: Hemoglobina - Mulheres Adultas**

```json
{
  "labTestId": "<uuid-HGB_F>",
  "gender": "female",
  "minAge": 18,
  "maxAge": null,
  "lowerLimit": 12.0,
  "upperLimit": 15.5,
  "unit": "g/dL",
  "description": "Faixa de referência para mulheres adultas",
  "source": "Diretrizes SBP 2024",
  "isActive": true
}
```

**Exemplo: TSH (múltiplas faixas por idade)**

```json
// Adultos jovens
{
  "labTestId": "<uuid-TSH>",
  "gender": null,
  "minAge": 18,
  "maxAge": 60,
  "lowerLimit": 0.4,
  "upperLimit": 4.0,
  "unit": "mU/L"
}

// Idosos
{
  "labTestId": "<uuid-TSH>",
  "gender": null,
  "minAge": 61,
  "maxAge": null,
  "lowerLimit": 0.5,
  "upperLimit": 6.0,
  "unit": "mU/L"
}
```

---

## 6. Hierarquia Completa - Exemplo Hemograma

```
HEMOGRAMA_COMPLETO (isRequestable: true, parentTestId: null)
├─ HGB_M (isRequestable: false, parentTestId: HEMOGRAMA_COMPLETO)
├─ HGB_F (isRequestable: false, parentTestId: HEMOGRAMA_COMPLETO)
├─ HCT_M (isRequestable: false, parentTestId: HEMOGRAMA_COMPLETO)
├─ HCT_F (isRequestable: false, parentTestId: HEMOGRAMA_COMPLETO)
├─ RBC_M (isRequestable: false, parentTestId: HEMOGRAMA_COMPLETO)
├─ RBC_F (isRequestable: false, parentTestId: HEMOGRAMA_COMPLETO)
├─ WBC (isRequestable: false, parentTestId: HEMOGRAMA_COMPLETO)
├─ NEUTROPHILS (isRequestable: false, parentTestId: HEMOGRAMA_COMPLETO)
├─ LYMPHOCYTES (isRequestable: false, parentTestId: HEMOGRAMA_COMPLETO)
├─ MONOCYTES (isRequestable: false, parentTestId: HEMOGRAMA_COMPLETO)
├─ EOSINOPHILS (isRequestable: false, parentTestId: HEMOGRAMA_COMPLETO)
├─ BASOPHILS (isRequestable: false, parentTestId: HEMOGRAMA_COMPLETO)
└─ PLATELETS (isRequestable: false, parentTestId: HEMOGRAMA_COMPLETO)
```

**Fluxo de Solicitação:**
1. Médico solicita → **HEMOGRAMA_COMPLETO** (único com isRequestable=true)
2. Laboratório retorna → Todos os 13 parâmetros filhos
3. Sistema armazena → LabResultValue para cada parâmetro

---

## 7. Casos de Uso

### 7.1 Solicitação de Exame

```typescript
// Frontend mostra apenas exames com isRequestable = true
GET /api/v1/lab-tests/requestable?category=hematology

Response:
[
  {
    "id": "uuid",
    "code": "HEMOGRAMA_COMPLETO",
    "name": "Hemograma Completo",
    "tussCode": "40304485",
    "category": "hematology"
  },
  // ... outros exames solicitáveis
]

// Médico seleciona "Hemograma Completo"
POST /api/v1/lab-results
{
  "patientId": "uuid",
  "testDefinitionId": "uuid-hemograma-completo",
  "status": "pending"
}
```

### 7.2 Entrada de Resultado

```typescript
// Sistema busca todos os subexames do Hemograma
GET /api/v1/lab-tests/{hemograma-id}/sub-tests

Response:
[
  {
    "id": "uuid-hgb-m",
    "code": "HGB_M",
    "name": "Hemoglobina - Homens",
    "unit": "g/dL",
    "referenceRange": { "lower": 13.5, "upper": 17.5 }
  },
  {
    "id": "uuid-wbc",
    "code": "WBC",
    "name": "Leucócitos Totais",
    "unit": "k/µL",
    "referenceRange": { "lower": 4.0, "upper": 11.0 }
  },
  // ... outros parâmetros
]

// Médico/técnico preenche valores estruturados
POST /api/v1/lab-results/{id}/values
{
  "values": [
    {
      "testDefinitionId": "uuid-hgb-m",
      "numericValue": 15.2,
      "unit": "g/dL"
    },
    {
      "testDefinitionId": "uuid-wbc",
      "numericValue": 7.8,
      "unit": "k/µL"
    }
    // ... outros valores
  ]
}
```

### 7.3 Cálculo Automático do Score

```typescript
// Sistema identifica mapeamentos para ScoreItems
LabResultValue (HGB_M = 15.2)
  → LabTestScoreMapping (gender=male)
  → ScoreItem "Hemoglobina - Homens"
  → ScoreLevels (avaliar 15.2 contra ranges)
  → PatientScoreItemValue (pontuação calculada)

// Se paciente for mulher, busca outro mapeamento
LabResultValue (HGB_F = 13.8)
  → LabTestScoreMapping (gender=female)
  → ScoreItem "Hemoglobina - Mulheres"
  → ScoreLevels diferentes
  → Pontuação diferente
```

---

## 8. Vantagens da Estrutura

### 8.1 Padronização
✅ Códigos TUSS para faturamento
✅ Códigos LOINC para interoperabilidade internacional
✅ Hierarquia clara entre exames compostos e parâmetros

### 8.2 Flexibilidade
✅ Suporte a múltiplas faixas de referência (por sexo, idade)
✅ Mapeamento condicional para ScoreItems
✅ Exames calculados (ex: Bilirrubina Indireta)

### 8.3 Usabilidade
✅ Filtrar exames solicitáveis vs apenas resultados
✅ Formulários dinâmicos baseados em definições
✅ Validação automática contra faixas de referência

### 8.4 Manutenção
✅ Centralizar mudanças de faixas de referência
✅ Ativar/desativar exames sem deletar
✅ Histórico de versões (via timestamps)

---

## 9. Próximos Passos

1. ✅ Criar models Go
2. ⏳ Criar migration
3. ⏳ Popular dados iniciais (hemograma completo)
4. ⏳ Criar endpoints API
5. ⏳ Criar formulários frontend
6. ⏳ Integrar com cálculo de escore

---

## 10. Script de População Inicial

Ver arquivo: `scripts/seed_lab_tests.sql` ou `scripts/seed_lab_tests.go`

Incluirá:
- Hemograma Completo com todos os 18+ parâmetros
- Bioquímica básica (glicose, ureia, creatinina, etc.)
- Perfil lipídico
- Função tireoidiana
- Eletrólitos
- Função hepática

Total estimado: ~200 definições de testes iniciais
