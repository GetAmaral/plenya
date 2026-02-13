# Backend - Go Models Patterns

## Template Obrigatório

**TODOS os models DEVEM seguir este template:**

```go
package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type MyModel struct {
    // 1. ID: SEMPRE UUID v7 (time-ordered, RFC 9562)
    // @example 019c1a1e-0579-7f3b-a1bd-4767008e844c
    ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

    // 2. Campos obrigatórios (NOT NULL)
    // @minLength 3
    // @maxLength 200
    Name string `gorm:"type:varchar(200);not null" json:"name" validate:"required,min=3,max=200"`

    // 3. Campos opcionais (nullable via pointer)
    Phone *string `gorm:"type:varchar(20)" json:"phone,omitempty" validate:"omitempty,min=10"`

    // 4. Enums com CHECK constraint
    // @enum active,inactive,archived
    Status string `gorm:"type:varchar(20);not null;check:status IN ('active','inactive','archived')" json:"status"`

    // 5. Foreign Keys com índice
    ParentID uuid.UUID `gorm:"type:uuid;not null;index" json:"parentId"`

    // 6. Relações (GORM Preload)
    Parent   *MyParent  `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE" json:"parent,omitempty"`
    Children []MyChild  `gorm:"foreignKey:ParentID" json:"children,omitempty"`

    // 7. Timestamps automáticos
    CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete
}

// 8. OBRIGATÓRIO: BeforeCreate hook para UUID v7
func (m *MyModel) BeforeCreate(tx *gorm.DB) error {
    if m.ID == uuid.Nil {
        m.ID = uuid.Must(uuid.NewV7())
    }
    return nil
}
```

## UUID v7 (RFC 9562) - OBRIGATÓRIO

### Por quê UUID v7?

- **Time-ordered:** Sortable cronologicamente
- **B-tree friendly:** 20-30% menos fragmentação de índice
- **Rastreabilidade:** Primeiros 48 bits = Unix timestamp ms
- **Performance:** Inserts sequenciais mais rápidos

### Estrutura

```
019c1a1e-0579-7f3b-a1bd-4767008e844c
│       │    │    │    │
│       │    │    │    └─ Random (48 bits)
│       │    │    └────── Version 7 + Variant
│       │    └─────────── Random (12 bits)
│       └──────────────── Timestamp milliseconds (48 bits)
└──────────────────────── Unix epoch ms
```

### Geração

**SEMPRE pela aplicação Go (BeforeCreate hook):**

```go
func (m *Model) BeforeCreate(tx *gorm.DB) error {
    if m.ID == uuid.Nil {
        m.ID = uuid.Must(uuid.NewV7())
    }
    return nil
}
```

**Fallback PostgreSQL (operações manuais SQL):**

```sql
-- Função helper (já criada no init.sql)
CREATE OR REPLACE FUNCTION uuid_generate_v7()
RETURNS UUID AS $$
DECLARE
  unix_ts_ms BIGINT;
  uuid_bytes BYTEA;
BEGIN
  unix_ts_ms := (EXTRACT(EPOCH FROM NOW()) * 1000)::BIGINT;
  uuid_bytes := (
    int8send(unix_ts_ms) ||
    gen_random_bytes(10)
  );
  RETURN uuid_bytes::UUID;
END;
$$ LANGUAGE plpgsql VOLATILE;

-- Uso
INSERT INTO patients (id, name) VALUES (uuid_generate_v7(), 'João');
```

## Tags GORM Essenciais

### Tipos e Constraints

```go
type Example struct {
    // VARCHAR com tamanho
    Name string `gorm:"type:varchar(200)"`

    // TEXT (ilimitado)
    Description string `gorm:"type:text"`

    // INTEGER
    Age int `gorm:"type:integer"`

    // DOUBLE PRECISION (float64)
    Score float64 `gorm:"type:double precision"`

    // BOOLEAN
    Active bool `gorm:"type:boolean"`

    // DATE
    BirthDate time.Time `gorm:"type:date"`

    // TIMESTAMP
    CreatedAt time.Time `gorm:"type:timestamp"`

    // UUID
    ID uuid.UUID `gorm:"type:uuid"`

    // NOT NULL
    Required string `gorm:"not null"`

    // UNIQUE
    Email string `gorm:"unique"`

    // DEFAULT
    Status string `gorm:"default:'active'"`

    // CHECK constraint
    Gender string `gorm:"check:gender IN ('male','female','other')"`

    // Múltiplos constraints
    Username string `gorm:"type:varchar(50);not null;unique;index"`
}
```

### Índices

```go
type Example struct {
    // Índice simples
    Email string `gorm:"index"`

    // Índice nomeado
    Username string `gorm:"index:idx_username"`

    // Índice composto (declarar em ambos os campos)
    FirstName string `gorm:"index:idx_full_name"`
    LastName  string `gorm:"index:idx_full_name"`

    // Índice único
    CPF string `gorm:"uniqueIndex"`

    // Índice parcial (via migration customizada)
    DeletedAt gorm.DeletedAt `gorm:"index"` // Apenas NOT NULL
}
```

### Foreign Keys e Relações

```go
type Child struct {
    // Foreign key (coluna)
    ParentID uuid.UUID `gorm:"type:uuid;not null;index" json:"parentId"`

    // Relação (carregada via Preload)
    Parent *Parent `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE" json:"parent,omitempty"`
}

type Parent struct {
    ID uuid.UUID `gorm:"type:uuid;primaryKey"`

    // Relação reversa (1:N)
    Children []Child `gorm:"foreignKey:ParentID" json:"children,omitempty"`
}
```

**Constraints ON DELETE:**

- `CASCADE` - Deleta filhos quando pai deletado
- `SET NULL` - Seta NULL nos filhos
- `RESTRICT` - Impede deletar pai se tiver filhos

### Timestamps Automáticos

```go
type Example struct {
    // Auto-preenche na criação
    CreatedAt time.Time `gorm:"autoCreateTime"`

    // Auto-atualiza em UPDATE
    UpdatedAt time.Time `gorm:"autoUpdateTime"`

    // Soft delete (seta timestamp em vez de deletar)
    DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

**Comportamento:**

- `Create()`: Preenche `CreatedAt` e `UpdatedAt`
- `Update()`: Atualiza apenas `UpdatedAt`
- `Delete()`: Seta `DeletedAt` (soft delete)
- Queries: GORM ignora automaticamente registros com `DeletedAt != NULL`

## Tags JSON (API Response)

```go
type Example struct {
    // Incluído no JSON com nome "id"
    ID uuid.UUID `json:"id"`

    // Incluído, omitido se vazio/nil
    Phone *string `json:"phone,omitempty"`

    // NUNCA incluído no JSON (segurança)
    Password string `json:"-"`

    // Renomeado no JSON
    FirstName string `json:"first_name"`

    // Snake_case automático (via Fiber config)
    BirthDate time.Time `json:"birthDate"` // serializa como "birthDate"
}
```

**Regras:**

- `json:"field"` → Sempre inclui
- `json:"field,omitempty"` → Omite se zero value (0, "", nil, false, empty slice/map)
- `json:"-"` → NUNCA serializa (senhas, tokens, dados sensíveis)

## Tags Validate (go-playground/validator)

```go
type CreatePatientDTO struct {
    // Obrigatório
    Name string `validate:"required"`

    // Tamanho string
    Username string `validate:"min=3,max=20"`

    // Email válido
    Email string `validate:"email"`

    // URL válida
    Website string `validate:"url"`

    // Enum (um dos valores)
    Status string `validate:"oneof=active inactive"`

    // Números
    Age int `validate:"gte=0,lte=150"` // Greater/Less Than or Equal

    // UUID válido
    ID string `validate:"uuid"`

    // CPF brasileiro (custom validator)
    CPF string `validate:"cpf"`

    // Telefone
    Phone string `validate:"e164"` // E.164 format: +5511999999999

    // Opcional com validação
    SecondaryEmail *string `validate:"omitempty,email"`

    // Data
    BirthDate string `validate:"required,datetime=2006-01-02"`

    // Combinações
    Password string `validate:"required,min=8,containsany=!@#$%"`
}
```

**Validação no handler:**

```go
func (h *Handler) Create(c *fiber.Ctx) error {
    var dto CreatePatientDTO
    if err := c.BodyParser(&dto); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
    }

    // Validar
    if err := h.validator.Struct(dto); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    // Processar...
}
```

## Annotations Swagger (Documentação)

### No Go Model

```go
type Patient struct {
    // ID único do paciente (UUID v7)
    // @example 019c1a1e-0579-7f3b-a1bd-4767008e844c
    ID uuid.UUID `json:"id"`

    // Nome completo do paciente
    // @minLength 3
    // @maxLength 200
    // @example João Silva Santos
    Name string `json:"name" validate:"required,min=3,max=200"`

    // Idade em anos (calculado automaticamente)
    // @minimum 0
    // @maximum 150
    // @example 45
    Age int `json:"age"`

    // Gênero do paciente
    // @enum male,female,other
    // @example male
    Gender string `json:"gender" validate:"oneof=male female other"`

    // Email de contato (opcional)
    // @example joao.silva@email.com
    Email *string `json:"email,omitempty" validate:"omitempty,email"`

    // CPF (criptografado no banco)
    // @pattern ^\d{3}\.\d{3}\.\d{3}-\d{2}$
    // @example 123.456.789-00
    CPF *string `json:"cpf,omitempty"`
}
```

### No Handler

```go
// CreatePatient godoc
// @Summary Create a new patient
// @Description Create a new patient record (admin/doctor only)
// @Tags Patients
// @Accept json
// @Produce json
// @Param body body services.CreatePatientDTO true "Patient Data"
// @Success 201 {object} models.Patient
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Failure 500 {object} map[string]interface{} "Internal error"
// @Security BearerAuth
// @Router /api/v1/patients [post]
func (h *PatientHandler) Create(c *fiber.Ctx) error {
    // Implementation...
}
```

**Gerar Swagger:**

```bash
cd apps/api
swag init -g cmd/server/main.go -o docs
```

## Padrões Específicos

### Nullable Fields (Ponteiros)

```go
type Patient struct {
    // Obrigatório (NOT NULL)
    Name string `gorm:"type:varchar(200);not null"`

    // Opcional (NULL permitido)
    Phone *string `gorm:"type:varchar(20)"`
    Email *string `gorm:"type:varchar(100)"`
}

// Uso
patient := &Patient{
    Name:  "João",
    Phone: stringPtr("+5511999999999"), // Helper function
    Email: nil, // Explicitamente NULL
}

// Helper
func stringPtr(s string) *string {
    return &s
}
```

### Enums com Validation

```go
type Patient struct {
    // @enum male,female,other
    Gender string `gorm:"type:varchar(10);not null;check:gender IN ('male','female','other')" json:"gender" validate:"oneof=male female other"`

    // @enum active,inactive,archived
    Status string `gorm:"type:varchar(20);default:'active';check:status IN ('active','inactive','archived')" json:"status"`
}
```

**Constantes (opcional):**

```go
const (
    GenderMale   = "male"
    GenderFemale = "female"
    GenderOther  = "other"
)

// Uso type-safe
patient.Gender = GenderMale
```

### Campos Calculados (Não Persistidos)

```go
type Patient struct {
    BirthDate time.Time `gorm:"type:date;not null" json:"birthDate"`

    // Calculado (não salvo no DB)
    Age int `gorm:"-" json:"age"`
}

// Calcular no hook AfterFind
func (p *Patient) AfterFind(tx *gorm.DB) error {
    p.CalculateAge()
    return nil
}

func (p *Patient) CalculateAge() {
    p.Age = int(time.Since(p.BirthDate).Hours() / 24 / 365)
}
```

### Campos Sensíveis (Criptografados)

```go
type Patient struct {
    // Criptografado antes de salvar
    // @pattern ^\d{3}\.\d{3}\.\d{3}-\d{2}$
    CPF *string `gorm:"type:text" json:"cpf,omitempty"`

    // Criptografado antes de salvar
    RG *string `gorm:"type:text" json:"rg,omitempty"`
}

// Hooks (ver hooks.md para implementação completa)
func (p *Patient) BeforeSave(tx *gorm.DB) error {
    if p.CPF != nil && *p.CPF != "" && !isEncrypted(*p.CPF) {
        encrypted, _ := crypto.EncryptWithDefaultKey(*p.CPF)
        p.CPF = &encrypted
    }
    return nil
}

func (p *Patient) AfterFind(tx *gorm.DB) error {
    if p.CPF != nil && *p.CPF != "" && isEncrypted(*p.CPF) {
        decrypted, _ := crypto.DecryptWithDefaultKey(*p.CPF)
        p.CPF = &decrypted
    }
    return nil
}
```

### JSON/JSONB Fields

```go
type Patient struct {
    // PostgreSQL JSONB
    Metadata map[string]interface{} `gorm:"type:jsonb" json:"metadata,omitempty"`

    // OU com tipo específico
    Preferences PatientPreferences `gorm:"type:jsonb;serializer:json" json:"preferences"`
}

type PatientPreferences struct {
    Language     string `json:"language"`
    Timezone     string `json:"timezone"`
    NotifyEmail  bool   `json:"notifyEmail"`
    NotifySMS    bool   `json:"notifySMS"`
}
```

**Query JSONB:**

```go
// Buscar por campo JSON
db.Where("metadata->>'key' = ?", "value").Find(&patients)

// Operador JSONB
db.Where("preferences @> ?", `{"language":"pt-BR"}`).Find(&patients)
```

### Self-Referencing (Hierarquia)

```go
type ScoreItem struct {
    ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
    Name string    `gorm:"type:varchar(300)"`

    // Self-reference (parent)
    ParentItemID *uuid.UUID `gorm:"type:uuid;index" json:"parentItemId,omitempty"`

    // Relações
    ParentItem *ScoreItem  `gorm:"foreignKey:ParentItemID;constraint:OnDelete:SET NULL" json:"parentItem,omitempty"`
    ChildItems []ScoreItem `gorm:"foreignKey:ParentItemID" json:"childItems,omitempty"`
}
```

### Many-to-Many

```go
type Article struct {
    ID    uuid.UUID   `gorm:"type:uuid;primaryKey"`
    Title string      `gorm:"type:varchar(500)"`

    // M:N
    ScoreItems []ScoreItem `gorm:"many2many:article_score_items;" json:"scoreItems,omitempty"`
}

type ScoreItem struct {
    ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
    Name string    `gorm:"type:varchar(300)"`

    // M:N (bidirecional)
    Articles []Article `gorm:"many2many:article_score_items;" json:"articles,omitempty"`
}
```

**GORM cria automaticamente:**

```sql
CREATE TABLE article_score_items (
    article_id UUID REFERENCES articles(id) ON DELETE CASCADE,
    score_item_id UUID REFERENCES score_items(id) ON DELETE CASCADE,
    PRIMARY KEY (article_id, score_item_id)
);
```

## Métodos de Negócio

**Colocar lógica de domínio no model:**

```go
type ScoreItem struct {
    Name          string  `gorm:"type:varchar(300)"`
    Gender        *string `gorm:"type:varchar(20)"`
    AgeRangeMin   *int    `gorm:"type:integer"`
    AgeRangeMax   *int    `gorm:"type:integer"`
    PostMenopause *bool   `gorm:"type:boolean"`
}

// Verifica se item se aplica a um paciente
func (si *ScoreItem) AppliesToPatient(patient *Patient) bool {
    if si.Gender != nil && *si.Gender != "not_applicable" {
        if *si.Gender != string(patient.Gender) {
            return false
        }
    }

    if si.AgeRangeMin != nil && patient.Age < *si.AgeRangeMin {
        return false
    }

    if si.AgeRangeMax != nil && patient.Age > *si.AgeRangeMax {
        return false
    }

    if patient.Gender == "female" && si.PostMenopause != nil {
        if patient.Menopause == nil || *si.PostMenopause != *patient.Menopause {
            return false
        }
    }

    return true
}
```

**Uso:**

```go
if scoreItem.AppliesToPatient(patient) {
    // Aplicar este score item ao paciente
}
```

## Checklist de Criação

Ao criar novo model:

- [ ] UUID v7 como Primary Key
- [ ] BeforeCreate hook para gerar UUID
- [ ] Timestamps (CreatedAt, UpdatedAt, DeletedAt)
- [ ] Tags GORM corretas (type, not null, index, etc.)
- [ ] Tags JSON (omitempty para opcionais, - para sensíveis)
- [ ] Tags Validate (DTOs)
- [ ] Annotations Swagger (examples, min/max, enum)
- [ ] Foreign keys com constraint (CASCADE, SET NULL)
- [ ] Check constraints para enums
- [ ] Métodos de negócio (se aplicável)
- [ ] Hooks especiais (criptografia, cálculos)
- [ ] Testar geração: `pnpm generate`

## Exemplos Completos

Ver arquivos:
- `apps/api/internal/models/patient.go`
- `apps/api/internal/models/score_item.go`
- `apps/api/internal/models/score_level.go`
- `apps/api/internal/models/article.go`
