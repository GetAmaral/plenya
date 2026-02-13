# 03 - Arquitetura: Única Fonte de Verdade

## Princípio Fundamental

**CRÍTICO:** Go models são a **ÚNICA fonte de verdade**. Tudo é gerado automaticamente a partir deles.

```
apps/api/internal/models/*.go  ← ÚNICA FONTE DE VERDADE
         │
         ├─→ Atlas CLI → SQL migrations (*.sql)
         ├─→ Swag → OpenAPI spec (swagger.json)
         │    │
         │    ├─→ openapi-typescript → TypeScript types
         │    └─→ zod-openapi → Zod schemas
         │
         └─→ pnpm generate (orquestra tudo)
```

## Workflow de Geração

### 1. Editar Go Model (Fonte Única)

```go
// apps/api/internal/models/patient.go
package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Patient struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name      string    `gorm:"type:varchar(200);not null" json:"name" validate:"required,min=3,max=200"`
    BirthDate time.Time `gorm:"type:date;not null" json:"birthDate" validate:"required"`
    Email     *string   `gorm:"type:varchar(100);unique" json:"email,omitempty" validate:"omitempty,email"`

    CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (p *Patient) BeforeCreate(tx *gorm.DB) error {
    if p.ID == uuid.Nil {
        p.ID = uuid.Must(uuid.NewV7())
    }
    return nil
}
```

### 2. Gerar Tudo Automaticamente

```bash
pnpm generate
```

Esse comando executa:

```bash
# Passo 1: Gerar migration SQL
cd apps/api && atlas migrate diff --env dev

# Passo 2: Gerar OpenAPI
cd apps/api && swag init -g cmd/server/main.go -o docs

# Passo 3: Gerar TypeScript types + Zod schemas
cd packages/types && pnpm run generate
```

### 3. Resultado: Arquivos Gerados

#### Migration SQL (gerado)

```sql
-- apps/api/database/migrations/20260212_add_email_to_patients.sql
ALTER TABLE patients ADD COLUMN email VARCHAR(100) UNIQUE;
CREATE UNIQUE INDEX IF NOT EXISTS idx_patients_email ON patients(email);
```

#### OpenAPI Spec (gerado)

```json
// apps/api/docs/swagger.json
{
  "components": {
    "schemas": {
      "Patient": {
        "type": "object",
        "properties": {
          "id": { "type": "string", "format": "uuid" },
          "name": { "type": "string", "minLength": 3, "maxLength": 200 },
          "birthDate": { "type": "string", "format": "date" },
          "email": { "type": "string", "format": "email" }
        },
        "required": ["id", "name", "birthDate"]
      }
    }
  }
}
```

#### TypeScript Types (gerado)

```typescript
// packages/types/src/generated/api.ts
export interface Patient {
  id: string
  name: string
  birthDate: string
  email?: string
  createdAt: string
  updatedAt: string
}
```

#### Zod Schemas (gerado)

```typescript
// packages/types/src/generated/schemas.ts
import { z } from 'zod'

export const PatientSchema = z.object({
  id: z.string().uuid(),
  name: z.string().min(3).max(200),
  birthDate: z.string(),
  email: z.string().email().optional(),
  createdAt: z.string(),
  updatedAt: z.string(),
})
```

## Regras de Ouro

### ✅ SEMPRE Editar

- `apps/api/internal/models/*.go` - Go models (fonte única)
- `apps/api/internal/handlers/*.go` - Swagger annotations
- `apps/api/internal/services/*.go` - Business logic
- `apps/api/internal/repository/*.go` - Database queries

### ❌ NUNCA Editar

- `apps/api/database/migrations/*.sql` - Gerado pelo Atlas
- `apps/api/docs/swagger.json` - Gerado pelo Swag
- `packages/types/src/generated/*.ts` - Gerado do OpenAPI

### Exceção: Migrations Customizadas

Se precisar de migration manual (dados, funções SQL):

```bash
# Criar migration vazia
atlas migrate new add_custom_function --env dev

# Editar manualmente
vim apps/api/database/migrations/<timestamp>_add_custom_function.sql
```

```sql
-- Migration customizada
CREATE OR REPLACE FUNCTION calculate_age(birth_date DATE)
RETURNS INTEGER AS $$
BEGIN
  RETURN DATE_PART('year', AGE(birth_date));
END;
$$ LANGUAGE plpgsql IMMUTABLE;
```

## Anatomia de um Go Model

### Tags GORM (Database)

```go
type Example struct {
    // Primary Key
    ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

    // Not Null
    Name string `gorm:"type:varchar(200);not null" json:"name"`

    // Nullable (pointer)
    Phone *string `gorm:"type:varchar(20)" json:"phone,omitempty"`

    // Unique
    Email string `gorm:"type:varchar(100);unique" json:"email"`

    // Default value
    Status string `gorm:"type:varchar(20);default:'active'" json:"status"`

    // Index
    UserID uuid.UUID `gorm:"type:uuid;index" json:"userId"`

    // Check constraint
    Age int `gorm:"type:integer;check:age >= 0 AND age <= 150" json:"age"`

    // Enum with check
    Gender string `gorm:"type:varchar(10);check:gender IN ('male','female','other')" json:"gender"`
}
```

### Tags JSON (API Response)

```go
type Example struct {
    // Incluído no JSON
    Name string `json:"name"`

    // Omitido se vazio
    Phone *string `json:"phone,omitempty"`

    // NUNCA incluído no JSON (segurança)
    Password string `json:"-"`
}
```

### Tags Validate (Validation)

```go
type Example struct {
    // Obrigatório
    Name string `validate:"required"`

    // Tamanho mínimo/máximo
    Username string `validate:"required,min=3,max=20"`

    // Email válido
    Email string `validate:"required,email"`

    // Enum
    Status string `validate:"oneof=active inactive"`

    // Número
    Age int `validate:"gte=0,lte=150"`

    // UUID
    UserID string `validate:"uuid"`

    // Opcional com validação
    Phone *string `validate:"omitempty,min=10"`
}
```

### Annotations Swagger (API Docs)

```go
type Patient struct {
    // ID do paciente
    // @example 019c1a1e-0579-7f3b-a1bd-4767008e844c
    ID uuid.UUID `json:"id"`

    // Nome completo do paciente
    // @minLength 3
    // @maxLength 200
    // @example João Silva Santos
    Name string `json:"name" validate:"required,min=3,max=200"`

    // Data de nascimento
    // @example 1985-03-15
    BirthDate time.Time `json:"birthDate" validate:"required"`

    // Gênero do paciente
    // @enum male,female,other
    // @example male
    Gender string `json:"gender" validate:"oneof=male female other"`
}
```

## Relações (GORM)

### One-to-Many

```go
// Parent
type Patient struct {
    ID       uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
    Name     string      `gorm:"type:varchar(200)" json:"name"`

    // Relação 1:N
    Anamneses []Anamnesis `gorm:"foreignKey:PatientID" json:"anamneses,omitempty"`
}

// Child
type Anamnesis struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`

    // Relação N:1
    Patient *Patient `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
}
```

**SQL gerado:**

```sql
CREATE TABLE patients (
  id UUID PRIMARY KEY,
  name VARCHAR(200)
);

CREATE TABLE anamneses (
  id UUID PRIMARY KEY,
  patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
  CONSTRAINT fk_anamneses_patient FOREIGN KEY (patient_id) REFERENCES patients(id)
);

CREATE INDEX idx_anamneses_patient_id ON anamneses(patient_id);
```

### Many-to-Many

```go
// ScoreItem
type ScoreItem struct {
    ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name string    `gorm:"type:varchar(300)" json:"name"`

    // Relação M:N
    Articles []Article `gorm:"many2many:article_score_items;" json:"articles,omitempty"`
}

// Article
type Article struct {
    ID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Title string    `gorm:"type:varchar(500)" json:"title"`

    // Relação M:N (bidirecional)
    ScoreItems []ScoreItem `gorm:"many2many:article_score_items;" json:"scoreItems,omitempty"`
}
```

**SQL gerado:**

```sql
CREATE TABLE score_items (
  id UUID PRIMARY KEY,
  name VARCHAR(300)
);

CREATE TABLE articles (
  id UUID PRIMARY KEY,
  title VARCHAR(500)
);

-- Tabela de junção
CREATE TABLE article_score_items (
  article_id UUID REFERENCES articles(id) ON DELETE CASCADE,
  score_item_id UUID REFERENCES score_items(id) ON DELETE CASCADE,
  PRIMARY KEY (article_id, score_item_id)
);
```

### Self-Referencing

```go
type ScoreItem struct {
    ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name string    `gorm:"type:varchar(300)" json:"name"`

    // Self-referencing (hierarquia)
    ParentItemID *uuid.UUID  `gorm:"type:uuid" json:"parentItemId,omitempty"`
    ParentItem   *ScoreItem  `gorm:"foreignKey:ParentItemID;constraint:OnDelete:SET NULL" json:"parentItem,omitempty"`
    ChildItems   []ScoreItem `gorm:"foreignKey:ParentItemID" json:"childItems,omitempty"`
}
```

**SQL gerado:**

```sql
CREATE TABLE score_items (
  id UUID PRIMARY KEY,
  name VARCHAR(300),
  parent_item_id UUID REFERENCES score_items(id) ON DELETE SET NULL
);
```

## Camadas da Arquitetura

### Layer 1: Models (Domain)

```go
// apps/api/internal/models/patient.go
package models

type Patient struct {
    ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name string    `gorm:"type:varchar(200);not null" json:"name"`

    // Métodos de negócio
    func (p *Patient) CalculateAge() int {
        return int(time.Since(p.BirthDate).Hours() / 24 / 365)
    }
}
```

### Layer 2: Repository (Data Access)

```go
// apps/api/internal/repository/patient_repository.go
package repository

type PatientRepository struct {
    db *gorm.DB
}

func (r *PatientRepository) GetByID(id uuid.UUID) (*models.Patient, error) {
    var patient models.Patient
    err := r.db.First(&patient, "id = ?", id).Error
    return &patient, err
}

func (r *PatientRepository) Create(patient *models.Patient) error {
    return r.db.Create(patient).Error
}
```

### Layer 3: Service (Business Logic)

```go
// apps/api/internal/services/patient_service.go
package services

type CreatePatientDTO struct {
    Name      string `json:"name" validate:"required,min=3,max=200"`
    BirthDate string `json:"birthDate" validate:"required"`
}

type PatientService struct {
    repo *repository.PatientRepository
}

func (s *PatientService) Create(dto CreatePatientDTO) (*models.Patient, error) {
    birthDate, _ := time.Parse("2006-01-02", dto.BirthDate)

    patient := &models.Patient{
        Name:      dto.Name,
        BirthDate: birthDate,
    }

    if err := s.repo.Create(patient); err != nil {
        return nil, err
    }

    return patient, nil
}
```

### Layer 4: Handler (HTTP/API)

```go
// apps/api/internal/handlers/patient_handler.go
package handlers

// CreatePatient godoc
// @Summary Create a new patient
// @Tags Patients
// @Accept json
// @Produce json
// @Param body body services.CreatePatientDTO true "Patient Data"
// @Success 201 {object} models.Patient
// @Router /api/v1/patients [post]
func (h *PatientHandler) Create(c *fiber.Ctx) error {
    var dto services.CreatePatientDTO
    if err := c.BodyParser(&dto); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
    }

    patient, err := h.service.Create(dto)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(201).JSON(patient)
}
```

### Layer 5: Routes (Registration)

```go
// apps/api/cmd/server/main.go
func setupRoutes(app *fiber.App, handlers *Handlers) {
    api := app.Group("/api/v1")

    // Patients
    api.Post("/patients", middleware.Auth(cfg), handlers.Patient.Create)
    api.Get("/patients/:id", middleware.Auth(cfg), handlers.Patient.GetByID)
}
```

## Fluxo de Dados Completo

### Request → Response

```
1. HTTP POST /api/v1/patients
   ↓
2. Fiber Router → Handler.Create()
   ↓
3. Handler valida DTO com validator
   ↓
4. Handler chama Service.Create(dto)
   ↓
5. Service cria Model
   ↓
6. Service chama Repository.Create(model)
   ↓
7. Repository salva no DB via GORM
   ↓  (BeforeCreate hook executa aqui - gera UUID v7)
8. Repository retorna Model
   ↓
9. Service retorna Model ao Handler
   ↓
10. Handler serializa Model → JSON
    ↓
11. Fiber retorna HTTP 201 + JSON
```

### Frontend Consumption

```typescript
// apps/web/lib/api/patients.ts
import { apiClient } from './client'
import { Patient, CreatePatientDTO } from '@plenya/types'

export function useCreatePatient() {
  return useMutation({
    mutationFn: (data: CreatePatientDTO) =>
      apiClient.post<Patient>('/api/v1/patients', data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['patients'] })
    },
  })
}
```

```tsx
// apps/web/app/(authenticated)/patients/new/page.tsx
function NewPatientPage() {
  const { mutate, isPending } = useCreatePatient()

  const onSubmit = (data: CreatePatientDTO) => {
    mutate(data, {
      onSuccess: () => router.push('/patients'),
    })
  }

  return <PatientForm onSubmit={onSubmit} isPending={isPending} />
}
```

## Versionamento de API

### Atual: v1

```
/api/v1/patients
/api/v1/anamneses
/api/v1/score-groups
```

### Futuro: v2 (Backward Compatible)

```
/api/v2/patients  # Nova versão
/api/v1/patients  # Ainda funciona (deprecated)
```

**Handler:**

```go
// v1
app.Get("/api/v1/patients/:id", handlers.V1.Patient.GetByID)

// v2
app.Get("/api/v2/patients/:id", handlers.V2.Patient.GetByID)
```

## Decisões de Design

### Por que Go models como fonte única?

1. **Type safety** - Go é statically typed
2. **Performance** - Go compila para binário nativo
3. **GORM tags** - Definem schema SQL automaticamente
4. **Swagger annotations** - Geram OpenAPI diretamente
5. **Single source** - 1 lugar para mudar, tudo propaga

### Por que Atlas para migrations?

1. **Schema-as-code** - Migrations geradas do Go
2. **Declarative** - Define estado desejado, não passos
3. **Rollback safety** - Detecta mudanças perigosas
4. **Versioning** - Git-friendly migrations

### Por que OpenAPI como intermediário?

1. **Standard** - OpenAPI 3.0 é padrão industry
2. **Tooling** - Geradores de código existem para tudo
3. **Documentation** - Swagger UI grátis
4. **Validation** - TypeScript + Zod gerados automaticamente

## Troubleshooting

### Migration conflita com state atual

```bash
# Ver diff
atlas migrate diff --env dev

# Se houver conflito, resetar (dev only!)
docker compose down -v
docker compose up -d
atlas migrate apply --env dev
```

### OpenAPI não reflete mudanças

```bash
# Regerar Swagger
cd apps/api
swag init -g cmd/server/main.go -o docs

# Verificar
cat docs/swagger.json | jq '.components.schemas.Patient'
```

### TypeScript types desatualizados

```bash
# Regerar types
cd packages/types
pnpm run generate

# Verificar
cat src/generated/api.ts | grep "interface Patient"
```

### Tudo dessincronizado

```bash
# Nuclear option: regenerar tudo
pnpm generate

# Revisar o que mudou
git diff
```
