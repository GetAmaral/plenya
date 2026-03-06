# Plano: Migração do App Treinador para o Plenya

## Contexto

O `app_treinador` é um sistema Streamlit/Python + Supabase com funcionalidades de avaliação física, gerador de treinos, periodização e IA (Gemini + ChromaDB). O objetivo é migrar **tudo exceto nutrição** para o Plenya (Go + Next.js + PostgreSQL), tornando o sistema treinador dispensável. A migração deve reutilizar a infraestrutura existente do Plenya (RAG com pgvector, Claude API, Article system, Score system, Patient model).

---

## Mapeamento: App Treinador → Plenya

| Componente App Treinador | Mapeamento no Plenya |
|--------------------------|---------------------|
| `anamneses` (Supabase) | `Patient` (existente) + novo `PhysicalAssessment` |
| `treinos` (Supabase) | Novo `WorkoutPlan` + `WorkoutPlanExercise` |
| `resultados_avaliacao` | `PatientScoreItemResult` (sistema de escores existente) |
| `exercises.json` (5000+) | Novo `Exercise` (seed data importado) |
| ChromaDB (ciência esporte) | `Article` system + pgvector (já existe) |
| Google Gemini + Embeddings | Claude API + Voyage/OpenAI (já configurado) |
| Estratificação ACSM (risk) | `ScoreGroup: "Avaliação Física"` + ScoreItems |
| GIFs de exercícios (/media/) | Migrar para VPS Plenya (servir localmente) |
| Periodização (Bomba) | Novo `WorkoutPeriodization` + `WorkoutMesocycle` |
| Agente IA | Novo handler usando `AIService` (já existe) |

---

## Fase 1: Go Models (Fonte Única de Verdade)

**Arquivo alvo:** `apps/api/internal/models/`

### Novos models a criar:

#### `exercise.go`
```
Exercise {
  ID: UUID (v7)
  ExerciseDbId: string (unique, ex: "trmte8s")
  Name: string (inglês original)
  NamePt: string (português)
  BodyPart: string (ex: "chest")
  BodyPartPt: string (ex: "Peito")
  TargetMuscle: string (ex: "pectorals")
  TargetMusclePt: string (ex: "Peitorais")
  Equipment: string (ex: "barbell")
  EquipmentPt: string (ex: "Barra")
  SecondaryMuscles: []string (JSONB)
  Instructions: []string (JSONB, inglês)
  InstructionsPt: []string (JSONB, português)
  GifUrl: string (local: /media/exercises/{id}.gif, fallback: CDN)
  IsActive: bool
}
```

#### `workout_plan.go`
```
WorkoutPlan {
  ID: UUID (v7)
  PatientID: UUID (FK → Patient)
  CreatedByID: UUID (FK → User)
  Name: string (ex: "Treino A - Hipertrofia")
  Objective: enum {hypertrophy, strength, endurance, weight_loss, conditioning, rehabilitation}
  Intensity: enum {very_light, light, moderate, high, very_high}
  DurationMinutes: int
  WeeklyFrequency: int
  PublicCode: string (8 chars, ex: "MUYY5R3H")
  IsActive: bool
  Sessions: []WorkoutPlanSession
}
```

#### `workout_plan_session.go`
```
WorkoutPlanSession {
  ID: UUID (v7)
  PlanID: UUID (FK → WorkoutPlan)
  Name: string (ex: "Sessão A", "Segunda-feira")
  Order: int
  Exercises: []WorkoutSessionExercise
}
```

#### `workout_session_exercise.go`
```
WorkoutSessionExercise {
  ID: UUID (v7)
  SessionID: UUID (FK → WorkoutPlanSession)
  ExerciseID: UUID (FK → Exercise)
  Phase: enum {warmup, main, cooldown}
  Order: int
  Sets: int
  Reps: string (ex: "8-12" ou "10")
  Cadence: enum {normal, slow, paused, explosive, free}
  RestBetweenSetsSec: int
  RestBetweenExercisesSec: int
  Notes: *string
}
```

#### `physical_assessment.go`
Avaliação física ACSM — dados brutos coletados:
```
PhysicalAssessment {
  ID: UUID (v7)
  PatientID: UUID (FK → Patient)
  CreatedByID: UUID (FK → User)
  AssessmentDate: time.Time

  // Antropometria
  Weight: *float64 (kg)
  Height: *float64 (cm)
  WaistCircumference: *float64 (cm)

  // Composição Corporal (calculados no BeforeSave)
  BMI: *float64 (calculado: peso/altura²)
  BRI: *float64 (Body Roundness Index)
  BodyFatPercent: *float64 (Pollock)
  LeanMass: *float64

  // Cardiovascular
  SystolicBP: *int (mmHg)
  DiastolicBP: *int (mmHg)
  RestingHeartRate: *int (bpm)

  // Laboratorial — preenchido automaticamente do último LabResult do paciente.
  // Se não houver LabResult, aceita input manual.
  LDL: *float64
  HDL: *float64
  TotalCholesterol: *float64
  Triglycerides: *float64
  FastingGlucose: *float64
  HbA1c: *float64

  // Histórico (ACSM risk factors)
  FamilyHistory: *bool
  SmokingStatus: enum {never, former, current}
  PhysicalActivityLevel: enum {sedentary, insufficient, moderate, active}

  // Condições
  CardiovascularDisease: *bool
  DiabetesType: *string
  Symptoms: *string
  ClinicalAlert: *bool

  // ACSM Estratificação (calculada no BeforeSave)
  ACSMRiskLevel: enum {low, moderate, high}
  ACSMRiskFactorsCount: int
  ACSMRecommendation: *string

  // Fotos
  FrontPhotoURL: *string
  SidePhotoURL: *string

  // Texto clínico (gerado por IA via Claude)
  AIRecommendation: *string
}
```

#### `workout_periodization.go`
```
WorkoutPeriodization {
  ID: UUID (v7)
  PatientID: UUID (FK → Patient)
  CreatedByID: UUID (FK → User)
  Framework: enum {bomba, linear, undulating, block}
  StartDate: time.Time
  TotalWeeks: int
  Objective: string
  ScientificJustification: *string (gerado por RAG + Claude)
  Mesocycles: []WorkoutMesocycle
}
```

#### `workout_mesocycle.go`
```
WorkoutMesocycle {
  ID: UUID (v7)
  PeriodizationID: UUID (FK)
  Order: int (1-6)
  Phase: enum {accumulation, transformation, realization, hypertrophy, strength, endurance}
  DurationWeeks: int
  VolumePercent: int (0-100)
  IntensityPercent: int (0-100)
  PhysiologicalFocus: string
  StartDate: time.Time
  EndDate: time.Time
}
```

**Nota:** Todos os models precisam de `BeforeCreate` hook para UUID v7.

---

## Fase 2: Migração de Dados

### 2.1 Importar Exercises.json → PostgreSQL
- Criar Go script em `apps/api/cmd/import-exercises/main.go`
- Lê `/root/app_treinador/dados/exercises.json` da VPS treinador (via SCP ou arquivo local)
- Usa `exercicios_mapeados_pt.json` para traduções PT
- Insere via GORM em batch (upsert por `exercise_db_id`)
- ~5000 exercícios, operação única

### 2.2 Migrar GIFs de Exercícios
```bash
# Copiar ~5000 GIFs da VPS treinador para a VPS Plenya
rsync -avz root@72.62.108.11:/root/app_treinador/media/ /path/plenya/media/exercises/

# Alternativa: Go script cmd/import-exercise-media que baixa do ExerciseDB CDN
# para cada exercício sem arquivo local
```
- Campo `GifUrl` no model usa caminho local
- Fallback automático para CDN se arquivo não existir

### 2.3 Importar Livros Científicos → Article System
Os livros já estão em `/root/app_treinador/dados/` como `.txt` e `.pdf`:
- NSCA Essentials 4th Ed.
- Bioquímica do Exercício (Mougios)
- Exercise Physiology (Plowman)
- Biomecânica (Hall)
- Tudor Bomba (Conditioning Young Athletes)
- ACSM's Exercise Testing
- etc.

**Usar o comando existente:** `apps/api/cmd/import-books-from-dir` (já existe no Plenya)
```bash
scp -r root@72.62.108.11:/root/app_treinador/dados/*.txt ./tmp-books/
go run ./apps/api/cmd/import-books-from-dir --dir=./tmp-books/ --specialty=sports_science
```
Checar `file_hash` para não reimportar livros já existentes.

### 2.4 Gerar Embeddings dos Livros
```bash
./scripts/enrichment/1-regenerate-embeddings.sh
```
Gera embeddings via Voyage/OpenAI e indexa em pgvector (pipeline já configurado).

### 2.5 Criar ScoreGroup "Avaliação Física"
Via SQL direto (`docker compose exec db psql ...`):
```sql
-- ScoreGroup: Avaliação Física
-- ScoreSubgroups: Composição Corporal, Risco Cardiovascular, Capacidade Funcional
-- ScoreItems: IMC, BRI, Body Fat%, ACSM Risk Level, Zona FC, etc.
-- ScoreLevels: faixas com classificação (ex: IMC 18.5-24.9 = "Normal")
```

---

## Fase 3: Backend Go (Handlers + Services)

### Novos handlers (padrão RESTful existente):

**`apps/api/internal/handlers/`**

| Handler | Endpoints Principais |
|---------|---------------------|
| `exercise_handler.go` | `GET /exercises` (search + filters), `GET /exercises/:id` |
| `workout_plan_handler.go` | CRUD `/workout-plans`, `GET /workout-plans/public/:code` |
| `physical_assessment_handler.go` | CRUD `/physical-assessments` (scoped por patient) |
| `periodization_handler.go` | CRUD `/periodizations`, `POST /periodizations/generate` (IA) |
| `training_ai_handler.go` | `POST /training/ai/chat`, `POST /training/ai/recommendations` |

### Novos services:

**`apps/api/internal/services/`**

| Service | Responsabilidade |
|---------|-----------------|
| `exercise_service.go` | Busca com filtros, paginação |
| `physical_assessment_service.go` | Cálculos ACSM (IMC, BRI, estratificação de risco) |
| `workout_plan_service.go` | Criação de planos, geração de public code |
| `periodization_service.go` | Geração Bomba via RAG + Claude |
| `training_ai_service.go` | Agente IA usando `AIService` existente |

### Cálculos ACSM (em `physical_assessment_service.go`):
```go
// Replicar lógica de /root/app_treinador/paginas/avaliadoschk.py

func CalculateACSMRisk(a *PhysicalAssessment, patient *Patient) ACSMRisk {
  // 7 fatores de risco cardiovascular ACSM 2018:
  // 1. Idade (M>=45, F>=55)
  // 2. Histórico familiar (FamilyHistory=true)
  // 3. Tabagismo (SmokingStatus=current)
  // 4. Sedentarismo (PhysicalActivityLevel=sedentary|insufficient)
  // 5. Obesidade (BMI >= 30)
  // 6. Dislipidemia (LDL>=130 ou HDL<40)
  // 7. Pré-diabetes (FastingGlucose>=100 ou HbA1c>=5.7)
  // Fator protetor: HDL >= 60 (desconta 1 fator)
  // Low: 0-1 fatores | Moderate: 2+ fatores | High: sintomas/doença cardiovascular
}

func CalculateBMI(weightKg, heightCm float64) float64
func CalculateBRI(waistCm, heightCm float64) float64
func CalculateMaxHR(age int) int
func CalculateKarvonenZones(maxHR, restHR int) [5]HRZone
```

### Geração de periodização com RAG + Claude:
```go
// POST /api/v1/periodizations/generate
// 1. Busca chunks científicos via embedding similarity
//    (reutilizar ArticleVectorRepository existente)
//    query: "Tudor Bomba periodization {objetivo} mesocycle"
// 2. Monta prompt com: dados do paciente + objetivo + chunks científicos (top 5-10)
// 3. Claude gera JSON estruturado com 6 mesociclos (fase, volume%, intensidade%, foco)
// 4. Salva em WorkoutPeriodization + WorkoutMesocycle
```

---

## Fase 4: Frontend Next.js

**Estrutura de páginas:** `apps/web/app/(authenticated)/training/`

```
training/
├── page.tsx                      # Dashboard: planos e avaliações do paciente
├── exercises/
│   └── page.tsx                  # Biblioteca (busca + filtros + GIFs)
├── workout-plans/
│   ├── page.tsx                  # Lista planos do paciente
│   ├── new/
│   │   └── page.tsx              # Criar plano + gerador de sessões
│   └── [id]/
│       ├── page.tsx              # Ver plano
│       └── edit/
│           └── page.tsx          # Editar plano
├── physical-assessments/
│   ├── new/
│   │   └── page.tsx              # Nova avaliação física (form ACSM)
│   └── [id]/
│       └── page.tsx              # Ver resultado (ACSM tags + gráficos)
├── periodization/
│   ├── page.tsx                  # Ver macrociclo ativo + calendário
│   └── new/
│       └── page.tsx              # Gerar periodização com IA
└── ai-agent/
    └── page.tsx                  # Chat com agente IA de treino
```

### Componentes principais (`apps/web/components/training/`):

| Componente | Descrição |
|-----------|-----------|
| `ExerciseCard.tsx` | Card com GIF, nome PT, músculos-alvo |
| `ExercisePicker.tsx` | Seletor com busca (corpo/músculo/equipamento) |
| `WorkoutPlanBuilder.tsx` | Drag-and-drop de exercícios por fase (warmup/main/cooldown) |
| `PhysicalAssessmentForm.tsx` | Form ACSM com `useFormNavigation` obrigatório |
| `ACSMRiskBadge.tsx` | Badge BAIXO/MODERADO/ALTO com cor |
| `PhysicalAssessmentResult.tsx` | Cards: IMC, BRI, BF%, Zonas FC |
| `PeriodizationCalendar.tsx` | Calendário mensal por fase (cor por tipo) |
| `MesocycleCard.tsx` | Card: fase, volume%, intensidade%, foco fisiológico |

### API clients (`apps/web/lib/api/`):

| Arquivo | Função |
|---------|--------|
| `exercise-api.ts` | GET /exercises com filtros |
| `workout-plan-api.ts` | CRUD workout plans |
| `physical-assessment-api.ts` | CRUD physical assessments |
| `periodization-api.ts` | CRUD + generate |
| `training-ai-api.ts` | Chat + recommendations |

**Padrões obrigatórios (igual ao restante do sistema):**
- `useFormNavigation({ formRef })` em todos os formulários
- `useRequireSelectedPatient()` em todas as páginas de dados do paciente
- TanStack Query para fetch/mutation

### Navegação (sidebar):
Adicionar seção "Treinamento" no sidebar com:
- Avaliação Física
- Planos de Treino
- Periodização
- Biblioteca de Exercícios
- Agente IA

---

## Fase 5: IA e RAG

### 5.1 Substituição de LLM (Gemini → Claude)
- `physical_assessment_service.go`: `AIRecommendation` via `AIService` existente
- `periodization_service.go`: mesociclos via RAG + Claude (mesmo padrão de `ScoreEnrichmentPreparationService`)
- `training_ai_service.go`: agente de chat/recomendações

### 5.2 RAG para Periodização
Reutilizar `ArticleVectorRepository` + `EmbeddingService` (já existem):
```go
// query embedding → pgvector similarity search → top chunks
// prompt = dados_paciente + objetivo + chunks_científicos
// Claude retorna JSON: 6 mesociclos com justificativa científica
```

### 5.3 RAG para Recomendações de Treino
```go
// query = "NSCA resistance training {objetivo} {nivel}"
// → top chunks de Essentials 4th Ed, Bomba, Exercise Physiology
// → Claude gera sugestão com referência científica
```

### 5.4 Agente IA de Treino
`training/ai-agent/page.tsx`:
- Contexto automático: dados do paciente selecionado (avaliação física, plano atual, periodização)
- Endpoint: `POST /api/v1/training/ai/chat`
- Histórico de conversa em localStorage

---

## Ordem de Execução (Sprints)

### Sprint 1 — Avaliação Física ACSM (prioridade máxima)
```
1. Go model: physical_assessment.go → pnpm generate
2. ScoreGroup "Avaliação Física" + ScoreItems ACSM → INSERT via SQL
3. Go service: physical_assessment_service.go (cálculos ACSM)
4. Go handler: physical_assessment_handler.go
5. Frontend: physical-assessments/new (form + useFormNavigation)
6. Frontend: physical-assessments/[id] (resultado com ACSMRiskBadge)
```

### Sprint 2 — Catálogo de Exercícios + Gerador de Treinos
```
7. Go models: exercise.go, workout_plan.go, workout_plan_session.go, workout_session_exercise.go
   → pnpm generate
8. Go script: cmd/import-exercises (exercises.json + exercicios_mapeados_pt.json)
9. Migrar GIFs: rsync VPS treinador → VPS Plenya
10. Go handler: exercise_handler.go (busca com filtros)
11. Go handler: workout_plan_handler.go (CRUD + public link)
12. Frontend: exercises/ (biblioteca com filtros + GIF preview)
13. Frontend: workout-plans/new (WorkoutPlanBuilder por fase)
14. Frontend: workout-plans/[id] (visualização + link público)
```

### Sprint 3 — RAG + Periodização + Agente IA
```
15. Importar livros esportivos → cmd/import-books-from-dir (já existe)
    → regenerate-embeddings (já existe)
16. Go models: workout_periodization.go, workout_mesocycle.go → pnpm generate
17. Go service: periodization_service.go (RAG + Claude)
18. Go handler: periodization_handler.go
19. Go service/handler: training_ai_service.go + training_ai_handler.go
20. Frontend: periodization/ (calendário macrociclo)
21. Frontend: ai-agent/ (chat com contexto do paciente)
```

### Sprint 4 — Migração Supabase (ao final, quando tudo estiver pronto)
```
22. Go script: cmd/migrate-from-supabase
    (anamneses → patients + physical_assessments, treinos → workout_plans)
```

---

## Decisões de Arquitetura

| Decisão | Escolha |
|---------|---------|
| Dados laboratoriais (LDL, HDL, glicemia) | `PhysicalAssessment` busca automaticamente o último `LabResult` do paciente. Aceita input manual se não houver. |
| GIFs dos exercícios | Migrar ~5000 GIFs da VPS treinador para VPS Plenya (servir localmente) |
| Migração Supabase | Sprint 4 — apenas após toda a estrutura nova estar funcional |
| Prioridade | **Avaliação física (ACSM) primeiro**, depois treinos, depois periodização/IA |
| Anamnese clínica vs. física | `Anamnesis` (notas clínicas) ≠ `PhysicalAssessment` (ACSM). Separados. |
| Peso/altura | `Patient` tem os valores atuais. `PhysicalAssessment` armazena snapshot da data da avaliação. |
| LLM | Gemini → Claude API (já configurada no Plenya via `AIService`) |
| Vector DB | ChromaDB → pgvector (já configurado no Plenya) |
| Livros científicos | Importar via `import-books-from-dir` (existente). Checar `file_hash` para evitar duplicatas. |

---

## Verificação / Testes End-to-End

1. `GET /api/v1/exercises?bodyPart=chest&equipment=barbell` → lista com GIF URLs locais
2. `POST /api/v1/workout-plans` com exercícios → retorna `public_code`
3. `GET /api/v1/workout-plans/public/MUYY5R3H` (sem auth) → plano público
4. `POST /api/v1/physical-assessments` com dados clínicos → `acsm_risk_level` calculado automaticamente
5. `POST /api/v1/periodizations/generate` → 6 mesociclos com justificativa científica
6. Frontend: `training/` → criar avaliação → ver resultado ACSM → criar plano → adicionar exercícios → link público

---

## Arquivos Críticos de Referência

**Plenya (padrões a seguir):**
- `apps/api/internal/models/score_item.go` — padrão de model + hooks
- `apps/api/internal/services/score_enrichment_preparation_service.go` — padrão RAG
- `apps/api/internal/services/ai_service.go` — Claude API configurada
- `apps/api/internal/services/embedding_service.go` — embeddings Voyage/OpenAI
- `apps/api/cmd/import-books-from-dir/` — importar livros (reutilizar)
- `apps/web/components/anamnesis/AnamnesisForm.tsx` — padrão de formulário

**App Treinador (lógica a replicar):**
- VPS `/root/app_treinador/paginas/avaliadoschk.py` — cálculos ACSM
- VPS `/root/app_treinador/paginas/gerador_treinos.py` — lógica do gerador
- VPS `/root/app_treinador/paginas/modulo_periodizacao.py` — periodização + RAG
- VPS `/root/app_treinador/dados/exercises.json` — catálogo de exercícios
- VPS `/root/app_treinador/dados/exercicios_mapeados_pt.json` — traduções PT
