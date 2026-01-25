# Proposta: Correlação de Dados Clínicos com Sistema de Escore Plenya

## 1. Análise da Estrutura Atual

### 1.1 Sistema de Escore (já implementado)
```
ScoreGroup → ScoreSubgroup → ScoreItem → ScoreLevel
```

**ScoreItem** contém:
- `name`: Nome do parâmetro (ex: "Hemoglobina - Homens")
- `unit`: Unidade de medida (ex: "g/dL")
- `unitConversion`: Conversão (ex: "1 g/dL = 10 g/L")
- `points`: Pontuação máxima do item

**ScoreLevel** contém:
- `level`: 0-5 (onde 5 = ótimo, 0 = crítico)
- `lowerLimit` / `upperLimit`: Faixas de valores
- `operator`: between, >, <, >=, <=, =
- `name`: Nome do nível (usado quando não há valores numéricos)

### 1.2 Dados Clínicos (já implementados)

**Patient:**
- Dados básicos: altura, peso (permite calcular IMC)
- Campos em texto livre, não estruturados

**Anamnesis:**
- **PROBLEMA:** Totalmente texto livre
- Campos: chiefComplaint, pastMedicalHistory, familyHistory, socialHistory, etc.
- **Impossível** correlacionar automaticamente com escore

**LabResult:**
- **PROBLEMA:** Campo `result` é texto livre
- Não há estrutura para mapear valores para ScoreItems
- Exemplo atual: `result: "Hemoglobina: 14.5 g/dL"`

### 1.3 Tipos de Dados no Escore Plenya

Analisando o CSV do escore, temos **3 categorias**:

#### A) Dados Laboratoriais Quantitativos (70% do escore)
```
Item: "Hemoglobina - Homens"
Unidade: "g/dL"
Níveis: N5=14.0-15.0, N4=15.0-16.5, N3=13.0-14.0, N2=10.0-13.0, N1=<10.0, N0=>16.5
Pontos: 20
```

#### B) Dados de Anamnese Qualitativos (25% do escore)
```
Item: "Aleitamento materno até qual idade"
Níveis: N5="2+ anos", N4="1 ano", N3="6m", N2="3m", N1="2m", N0="Nenhum"
Pontos: 5
```

#### C) Dados Calculados (5% do escore)
```
Item: "IMC"
Calculado a partir de: peso / (altura²)
Níveis: ranges numéricos
```

---

## 2. Problemas a Resolver

### 2.1 Falta de Correlação
- ❌ Não há ligação entre `LabResult` e `ScoreItem`
- ❌ Não há ligação entre `Anamnesis` e `ScoreItem`
- ❌ Não há modelo para armazenar "valor clínico estruturado de um paciente para um item do escore"

### 2.2 Dados Não Estruturados
- ❌ `LabResult.result` é texto livre
- ❌ `Anamnesis` é texto livre
- ❌ Impossível fazer queries ou cálculos automáticos

### 2.3 Versionamento do Escore
- ❌ Se o escore mudar (novos níveis, faixas ajustadas), como manter histórico?
- ❌ Como recalcular scores antigos com nova versão?

### 2.4 Identificação de Itens
- ❌ ScoreItem não tem código único (ex: "HGB_M", "GLUCOSE_FASTING")
- ❌ Dependência de string matching (frágil)

---

## 3. Solução Proposta

### 3.1 Arquitetura de 3 Camadas

```
┌─────────────────────────────────────────────────────────────┐
│ CAMADA 1: Score Definition (já existe)                      │
│ ScoreGroup → ScoreSubgroup → ScoreItem → ScoreLevel         │
│ + ADICIONAR: ScoreItem.code (identificador único)           │
│ + ADICIONAR: ScoreItem.dataType (lab, anamnesis, calculated)│
│ + ADICIONAR: ScoreVersion (versionamento)                   │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│ CAMADA 2: Clinical Data (estruturado)                       │
│ LabResultValue → mapeia para ScoreItem via code             │
│ AnamnesisValue → mapeia para ScoreItem via code             │
│ CalculatedValue → calculado automaticamente                 │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│ CAMADA 3: Patient Score (calculado automaticamente)         │
│ PatientScoreSnapshot → escore completo do paciente em T     │
│ PatientScoreItemValue → valor individual por item           │
└─────────────────────────────────────────────────────────────┘
```

---

## 4. Novos Models Propostos

### 4.1 Adicionar em ScoreItem

```go
type ScoreItem struct {
    // ... campos existentes ...

    // Código único para identificação programática
    // @example HGB_M, GLUCOSE_FASTING, IMC, BREASTFEEDING_DURATION
    Code *string `gorm:"type:varchar(100);unique;index" json:"code,omitempty"`

    // Tipo de dado
    // @enum lab, anamnesis, calculated, anthropometric
    DataType *string `gorm:"type:varchar(20)" json:"dataType,omitempty"`

    // Para dados calculados, fórmula em JSON
    // @example {"formula": "weight / (height * height)", "inputs": ["weight", "height"]}
    CalculationFormula *string `gorm:"type:text" json:"calculationFormula,omitempty"`
}
```

### 4.2 ScoreVersion (versionamento do escore)

```go
// ScoreVersion representa uma versão específica do sistema de escore
type ScoreVersion struct {
    ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
    Version     string         `gorm:"type:varchar(20);not null;unique" json:"version"` // Ex: "1.0", "1.1", "2.0"
    Description string         `gorm:"type:text" json:"description"`
    IsActive    bool           `gorm:"type:boolean;not null;default:false" json:"isActive"`
    EffectiveAt time.Time      `gorm:"type:timestamp;not null" json:"effectiveAt"`
    CreatedAt   time.Time      `json:"createdAt"`
    UpdatedAt   time.Time      `json:"updatedAt"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
```

**Adicionar em ScoreGroup:**
```go
VersionID *uuid.UUID `gorm:"type:uuid;index" json:"versionId,omitempty"`
Version   *ScoreVersion `gorm:"foreignKey:VersionID" json:"version,omitempty"`
```

### 4.3 LabResultValue (valores estruturados de exames)

```go
// LabResultValue armazena valores estruturados de exames laboratoriais
// Substitui/complementa o campo texto livre de LabResult
type LabResultValue struct {
    ID           uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
    LabResultID  uuid.UUID      `gorm:"type:uuid;not null;index" json:"labResultId"`
    ScoreItemID  *uuid.UUID     `gorm:"type:uuid;index" json:"scoreItemId,omitempty"` // Mapeia para item do escore

    // Identificação do parâmetro (se não mapear para ScoreItem)
    ParameterName string        `gorm:"type:varchar(200);not null" json:"parameterName"`
    ParameterCode *string       `gorm:"type:varchar(100);index" json:"parameterCode,omitempty"` // Ex: HGB, WBC

    // Valor (numérico ou texto, dependendo do tipo)
    NumericValue  *float64      `gorm:"type:double precision" json:"numericValue,omitempty"`
    TextValue     *string       `gorm:"type:text" json:"textValue,omitempty"`

    // Metadados
    Unit              *string   `gorm:"type:varchar(50)" json:"unit,omitempty"`
    ReferenceRange    *string   `gorm:"type:varchar(200)" json:"referenceRange,omitempty"`
    IsAbnormal        bool      `gorm:"type:boolean;default:false" json:"isAbnormal"`

    CreatedAt    time.Time      `json:"createdAt"`
    UpdatedAt    time.Time      `json:"updatedAt"`
    DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

    // Relacionamentos
    LabResult LabResult  `gorm:"foreignKey:LabResultID" json:"labResult,omitempty"`
    ScoreItem *ScoreItem `gorm:"foreignKey:ScoreItemID" json:"scoreItem,omitempty"`
}
```

### 4.4 AnamnesisValue (valores estruturados de anamnese)

```go
// AnamnesisValue armazena respostas estruturadas de anamnese
type AnamnesisValue struct {
    ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
    PatientID   uuid.UUID      `gorm:"type:uuid;not null;index" json:"patientId"`
    ScoreItemID uuid.UUID      `gorm:"type:uuid;not null;index" json:"scoreItemId"`
    DoctorID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"doctorId"`

    // Valor (numérico para escalas, texto para categórico, booleano para sim/não)
    NumericValue  *float64  `gorm:"type:double precision" json:"numericValue,omitempty"`
    TextValue     *string   `gorm:"type:varchar(500)" json:"textValue,omitempty"`
    BooleanValue  *bool     `gorm:"type:boolean" json:"booleanValue,omitempty"`

    // Nível do escore selecionado (se aplicável)
    ScoreLevelID  *uuid.UUID `gorm:"type:uuid" json:"scoreLevelId,omitempty"`

    // Data da avaliação
    AssessmentDate time.Time `gorm:"type:timestamp;not null;index" json:"assessmentDate"`

    Notes       *string        `gorm:"type:text" json:"notes,omitempty"`
    CreatedAt   time.Time      `json:"createdAt"`
    UpdatedAt   time.Time      `json:"updatedAt"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

    // Relacionamentos
    Patient   Patient    `gorm:"foreignKey:PatientID" json:"patient,omitempty"`
    ScoreItem ScoreItem  `gorm:"foreignKey:ScoreItemID" json:"scoreItem,omitempty"`
    ScoreLevel *ScoreLevel `gorm:"foreignKey:ScoreLevelID" json:"scoreLevel,omitempty"`
    Doctor    User       `gorm:"foreignKey:DoctorID" json:"doctor,omitempty"`
}
```

### 4.5 PatientScoreSnapshot (escore calculado)

```go
// PatientScoreSnapshot representa o escore completo de um paciente em um momento específico
type PatientScoreSnapshot struct {
    ID              uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
    PatientID       uuid.UUID      `gorm:"type:uuid;not null;index" json:"patientId"`
    ScoreVersionID  uuid.UUID      `gorm:"type:uuid;not null;index" json:"scoreVersionId"`

    // Score total
    TotalScore      float64        `gorm:"type:double precision;not null" json:"totalScore"`
    MaxPossibleScore float64       `gorm:"type:double precision;not null" json:"maxPossibleScore"`
    ScorePercentage float64        `gorm:"type:double precision;not null" json:"scorePercentage"` // 0-100

    // Classificação
    // @enum critical, low, medium, good, excellent
    Classification  string         `gorm:"type:varchar(20);not null" json:"classification"`

    // Data base para o cálculo
    CalculatedAt    time.Time      `gorm:"type:timestamp;not null;index" json:"calculatedAt"`

    // Metadados
    ItemsEvaluated  int            `gorm:"type:integer;not null" json:"itemsEvaluated"`
    ItemsMissing    int            `gorm:"type:integer;not null" json:"itemsMissing"`
    CompletionRate  float64        `gorm:"type:double precision;not null" json:"completionRate"` // 0-100

    Notes           *string        `gorm:"type:text" json:"notes,omitempty"`
    CreatedAt       time.Time      `json:"createdAt"`
    UpdatedAt       time.Time      `json:"updatedAt"`
    DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

    // Relacionamentos
    Patient      Patient       `gorm:"foreignKey:PatientID" json:"patient,omitempty"`
    ScoreVersion ScoreVersion  `gorm:"foreignKey:ScoreVersionID" json:"scoreVersion,omitempty"`
    ItemValues   []PatientScoreItemValue `gorm:"foreignKey:SnapshotID" json:"itemValues,omitempty"`
}
```

### 4.6 PatientScoreItemValue (valor individual por item)

```go
// PatientScoreItemValue armazena o valor e pontuação de cada item do escore para um paciente
type PatientScoreItemValue struct {
    ID              uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
    SnapshotID      uuid.UUID      `gorm:"type:uuid;not null;index" json:"snapshotId"`
    ScoreItemID     uuid.UUID      `gorm:"type:uuid;not null;index" json:"scoreItemId"`
    ScoreLevelID    *uuid.UUID     `gorm:"type:uuid" json:"scoreLevelId,omitempty"` // Nível alcançado

    // Valor obtido (numérico, texto ou booleano dependendo do tipo)
    NumericValue    *float64       `gorm:"type:double precision" json:"numericValue,omitempty"`
    TextValue       *string        `gorm:"type:varchar(500)" json:"textValue,omitempty"`
    BooleanValue    *bool          `gorm:"type:boolean" json:"booleanValue,omitempty"`

    // Pontuação
    PointsObtained  float64        `gorm:"type:double precision;not null" json:"pointsObtained"`
    MaxPoints       float64        `gorm:"type:double precision;not null" json:"maxPoints"`

    // Origem do dado
    // @enum lab_result, anamnesis, calculated, anthropometric, manual
    DataSource      string         `gorm:"type:varchar(20);not null" json:"dataSource"`
    SourceID        *uuid.UUID     `gorm:"type:uuid" json:"sourceId,omitempty"` // ID do LabResultValue ou AnamnesisValue

    CreatedAt       time.Time      `json:"createdAt"`
    UpdatedAt       time.Time      `json:"updatedAt"`

    // Relacionamentos
    Snapshot   PatientScoreSnapshot `gorm:"foreignKey:SnapshotID" json:"snapshot,omitempty"`
    ScoreItem  ScoreItem            `gorm:"foreignKey:ScoreItemID" json:"scoreItem,omitempty"`
    ScoreLevel *ScoreLevel          `gorm:"foreignKey:ScoreLevelID" json:"scoreLevel,omitempty"`
}
```

---

## 5. Fluxo de Dados

### 5.1 Registro de Exame Laboratorial

```
1. Médico solicita exame → LabResult (status: pending)
2. Resultado chega do laboratório → LabResult (status: completed)
3. Sistema parseia resultado em valores estruturados:

   LabResult {
       id: uuid-1
       testName: "Hemograma Completo"
       result: "Hemoglobina: 14.5 g/dL, Leucócitos: 7.2 k/µL..."
   }

   ↓ AUTO-PARSING ou ENTRADA MANUAL ESTRUTURADA ↓

   LabResultValue {
       labResultId: uuid-1
       parameterCode: "HGB_M"
       scoreItemId: (lookup by code "HGB_M")
       numericValue: 14.5
       unit: "g/dL"
   }

   LabResultValue {
       labResultId: uuid-1
       parameterCode: "WBC"
       scoreItemId: (lookup by code "WBC")
       numericValue: 7.2
       unit: "k/µL"
   }
```

### 5.2 Registro de Anamnese Estruturada

```
1. Médico preenche questionário estruturado na interface
2. Para cada pergunta do escore:

   AnamnesisValue {
       patientId: patient-uuid
       scoreItemId: (item "Aleitamento materno")
       textValue: "1 ano"
       scoreLevelId: (level N4 correspondente)
       assessmentDate: 2025-01-25
   }
```

### 5.3 Cálculo Automático do Escore

```
TRIGGER: Novo exame completo OU requisição manual

1. Buscar última versão ativa do escore (ScoreVersion)
2. Para cada ScoreItem na versão:

   A) Se dataType = "lab":
      → Buscar último LabResultValue com scoreItemId correspondente
      → Avaliar valor contra ScoreLevels (operator + limits)
      → Atribuir pontuação

   B) Se dataType = "anamnesis":
      → Buscar último AnamnesisValue com scoreItemId correspondente
      → Pegar scoreLevelId direto
      → Atribuir pontuação

   C) Se dataType = "calculated":
      → Executar cálculo (ex: IMC = peso / altura²)
      → Avaliar contra ScoreLevels
      → Atribuir pontuação

   D) Se dataType = "anthropometric":
      → Pegar valor de Patient (height, weight)
      → Avaliar contra ScoreLevels
      → Atribuir pontuação

3. Criar PatientScoreSnapshot com total
4. Criar PatientScoreItemValue para cada item avaliado
```

---

## 6. Interface Proposta

### 6.1 Entrada de Exames (Web/Mobile)

**Opção A: Upload de PDF + OCR (Fase futura)**
- Upload do PDF do laboratório
- OCR extrai valores automaticamente
- Sistema mapeia para ScoreItems via códigos

**Opção B: Formulário Estruturado (Fase atual - RECOMENDADO)**
```tsx
// Ao criar LabResult, apresentar formulário para cada parâmetro conhecido

<Form>
  <h3>Hemograma Completo</h3>

  <Input
    label="Hemoglobina"
    unit="g/dL"
    scoreItemCode="HGB_M"
    type="number"
    referenceRange="14.0 - 18.0"
  />

  <Input
    label="Leucócitos"
    unit="k/µL"
    scoreItemCode="WBC"
    type="number"
    referenceRange="4.0 - 11.0"
  />

  {/* ... outros parâmetros */}
</Form>

// Ao salvar, cria automaticamente LabResultValues mapeados
```

### 6.2 Formulário de Anamnese Estruturada

```tsx
// Renderizar dinamicamente com base nos ScoreItems de type="anamnesis"

<Form>
  <h3>Histórico - Infância</h3>

  <Select
    label="Aleitamento materno até qual idade?"
    scoreItemId={item.id}
    options={item.levels.map(level => ({
      value: level.id,
      label: level.name, // "Nenhum", "2m", "3m", "6m", "1 ano", "2+ anos"
      points: calculatePoints(level.level, item.points)
    }))}
  />

  <Select
    label="Qualidade da alimentação na idade pré-escolar"
    scoreItemId={item.id}
    options={item.levels.map(...)}
  />
</Form>
```

### 6.3 Dashboard de Escore

```tsx
// Exibir último PatientScoreSnapshot

<ScoreDashboard>
  <ScoreTotal>
    Score: {snapshot.totalScore} / {snapshot.maxPossibleScore}
    ({snapshot.scorePercentage}%)

    Classificação: {snapshot.classification}
    Data: {snapshot.calculatedAt}
  </ScoreTotal>

  <CompletionRate>
    {snapshot.itemsEvaluated} / {snapshot.itemsEvaluated + snapshot.itemsMissing} itens
    ({snapshot.completionRate}%)
  </CompletionRate>

  <ItemBreakdown>
    {snapshot.itemValues.map(itemValue => (
      <ScoreItemCard>
        <Name>{itemValue.scoreItem.name}</Name>
        <Value>{itemValue.numericValue} {itemValue.scoreItem.unit}</Value>
        <Level color={getLevelColor(itemValue.scoreLevel.level)}>
          Nível {itemValue.scoreLevel.level}
        </Level>
        <Points>{itemValue.pointsObtained} / {itemValue.maxPoints} pts</Points>
      </ScoreItemCard>
    ))}
  </ItemBreakdown>
</ScoreDashboard>
```

---

## 7. Migração e Implementação

### 7.1 Fase 1: Adicionar Campos em ScoreItem (1 dia)
- Adicionar `code`, `dataType`, `calculationFormula`
- Rodar migration
- Popular códigos para itens existentes via script

### 7.2 Fase 2: Criar Novos Models (2 dias)
- ScoreVersion
- LabResultValue
- AnamnesisValue
- PatientScoreSnapshot
- PatientScoreItemValue
- Migrations
- Repositories
- Services

### 7.3 Fase 3: API Endpoints (3 dias)
- POST /api/v1/lab-results/:id/values (criar valores estruturados)
- POST /api/v1/patients/:id/anamnesis-values (criar respostas estruturadas)
- POST /api/v1/patients/:id/score/calculate (calcular escore)
- GET /api/v1/patients/:id/score/latest (último snapshot)
- GET /api/v1/patients/:id/score/history (histórico de snapshots)

### 7.4 Fase 4: Frontend (5 dias)
- Formulário estruturado de exames
- Formulário estruturado de anamnese
- Dashboard de escore calculado
- Gráficos de evolução temporal

### 7.5 Fase 5: Cálculo Automático (3 dias)
- Engine de cálculo de escore
- Triggers automáticos (novo exame → recalcular)
- Classificação automática
- Notificações (score crítico, melhora significativa)

---

## 8. Benefícios da Solução

### 8.1 Automação
✅ Escore calculado automaticamente ao inserir exames
✅ Reduz tempo médico de análise manual
✅ Elimina erros de cálculo

### 8.2 Rastreabilidade
✅ Cada valor do escore tem origem rastreável (qual exame, qual consulta)
✅ Histórico completo de evolução do paciente
✅ Auditoria de mudanças

### 8.3 Versionamento
✅ Permite ajustar faixas de referência sem perder histórico
✅ Recalcular scores antigos com nova metodologia
✅ Comparar versões do escore

### 8.4 Insights
✅ Queries SQL complexas possíveis (ex: "pacientes com hemoglobina <10")
✅ Relatórios populacionais
✅ Machine learning futuro (predição de scores)

### 8.5 Experiência do Usuário
✅ Formulários inteligentes e validados
✅ Feedback imediato sobre score
✅ Visualizações ricas (gráficos, trends, alerts)

---

## 9. Exemplo de Uso Real

### Cenário: Paciente João, 45 anos

**1. Cadastro inicial:**
```
Patient: altura=175cm, peso=82kg
→ Sistema calcula IMC = 26.8 (sobrepeso)
→ PatientScoreItemValue criado para item "IMC"
```

**2. Primeira consulta - Anamnese:**
```
Médico preenche formulário estruturado:
- Aleitamento materno: "1 ano" → AnamnesisValue
- Alimentação infantil: "Alimentação adequada" → AnamnesisValue
- Exercício atual: "3x/semana moderado" → AnamnesisValue

→ Sistema calcula pontuação parcial
```

**3. Exames solicitados:**
```
LabResult criado (status: pending)
Após resultado, médico preenche:
- Hemoglobina: 15.2 g/dL → LabResultValue
- Glicemia jejum: 105 mg/dL → LabResultValue
- Colesterol total: 220 mg/dL → LabResultValue
- HDL: 45 mg/dL → LabResultValue

→ Sistema mapeia cada valor para ScoreItem via code
→ Avalia contra ScoreLevels
→ Calcula pontuação
```

**4. Cálculo Final do Escore:**
```
POST /api/v1/patients/{joao-id}/score/calculate

→ PatientScoreSnapshot criado:
  totalScore: 2.340
  maxPossibleScore: 4.000
  scorePercentage: 58.5%
  classification: "medium"
  itemsEvaluated: 87
  itemsMissing: 213
  completionRate: 29%

→ 87 PatientScoreItemValue criados (um por item avaliado)
```

**5. Dashboard mostra:**
- Score total: 58.5% (Médio)
- Áreas críticas: Glicemia (limítrofe), HDL (baixo)
- Áreas boas: Hemoglobina (ótimo), História alimentar (bom)
- Recomendação: Completar exames faltantes para score completo

---

## 10. Recomendações Finais

### Prioridade ALTA (Fazer agora):
1. ✅ Adicionar `code` em ScoreItem
2. ✅ Criar LabResultValue model
3. ✅ Criar PatientScoreSnapshot + PatientScoreItemValue
4. ✅ API básica de cálculo de escore

### Prioridade MÉDIA (Próximas sprints):
5. ⏳ Criar AnamnesisValue model
6. ⏳ Formulários estruturados frontend
7. ⏳ Dashboard de escore

### Prioridade BAIXA (Futuro):
8. ⏸️ ScoreVersion (versionamento)
9. ⏸️ OCR de PDFs de laboratório
10. ⏸️ Machine learning para predições

---

## 11. Perguntas para Validação

Antes de implementar, precisamos decidir:

**Q1:** Preferência de entrada de exames?
- [ ] Formulário estruturado (mais trabalhoso, dados perfeitos)
- [ ] Texto livre + parsing posterior (mais rápido, menos preciso)
- [ ] Híbrido (texto + campos principais estruturados)

**Q2:** Quando recalcular escore?
- [ ] Manual (médico clica "Calcular Escore")
- [ ] Automático (qualquer novo dado → recalcula)
- [ ] Agendado (1x/dia recalcula todos os pacientes)

**Q3:** Dados de anamnese históricos (pré-natal, infância):
- [ ] Formulário único gigante (questionário inicial completo)
- [ ] Formulários modulares (preencher ao longo do tempo)
- [ ] Texto livre + campos-chave estruturados

**Q4:** Prioridade:
- [ ] Começar com exames laboratoriais apenas (70% do escore)
- [ ] Começar com anamnese apenas (contexto qualitativo)
- [ ] Implementar ambos em paralelo

---

**Proposta preparada por:** Claude Sonnet 4.5
**Data:** 2025-01-25
**Status:** Aguardando validação e aprovação
