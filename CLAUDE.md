# Plenya - Sistema de Prontuário Médico Eletrônico (EMR)

## PARTE 1: VISÃO GERAL E FUNDAÇÃO

### 1. Visão Geral

Sistema EMR completo com web app, mobile apps e backend Go. Foco em segurança LGPD e baixo custo.

**Escala:** Centenas de pacientes
**Plataformas:** Web (Next.js) + Mobile (iOS/Android via Expo)
**Infraestrutura:** Self-hosted Hetzner VPS + Coolify

---

### 2. Stack Técnica

**Frontend**
- Web: Next.js 16.1 + React 19.2
- Mobile: React Native 0.77 + Expo SDK 56
- UI: Tailwind CSS + shadcn/ui
- Charts: Recharts + Tremor Raw
- Forms: React Hook Form + Zod (gerado automaticamente)
- State: TanStack Query + Zustand

**Backend**
- Language: Go 1.25
- Framework: Fiber v2.52.10
- ORM: GORM v1.25
- Migrations: Atlas (gerado dos Go models)
- Auth: JWT (golang-jwt/jwt v5)
- Validation: go-playground/validator v10

**Database**
- SGBD: PostgreSQL 17
- Extensions: pgcrypto (criptografia), uuid-ossp

**Infraestrutura**
- Monorepo: Turborepo 2.7.5 + pnpm 10.28.1
- Node.js: 24 LTS (Krypton)
- TypeScript: 5.9.3
- Containers: Docker 27
- VPS: Hetzner CPX31 (4 vCPU, 8GB, ~R$50/mês)
- PaaS: Coolify 4.0
- CDN/DNS: Cloudflare Free

---

### 3. Arquitetura: Única Fonte de Verdade

**CRÍTICO:** Go models são a única fonte de verdade. Tudo é gerado automaticamente.

```
apps/api/internal/models/*.go  ← ÚNICA FONTE
         │
         ├─→ Atlas → SQL migrations
         ├─→ Swag → OpenAPI spec
         │    ├─→ TypeScript types (gerado)
         │    └─→ Zod schemas (gerado)
         │
         └─→ pnpm generate
```

**Nunca editar:**
- `apps/api/database/migrations/*.sql` (gerado)
- `packages/types/src/generated/*.ts` (gerado)

**Sempre editar:**
- `apps/api/internal/models/*.go` (fonte única)

**Workflow de Mudanças:**

```bash
# 1. Editar APENAS o Go model
vim apps/api/internal/models/patient.go

# 2. Gerar tudo automaticamente
pnpm generate

# 3. Revisar o que foi gerado
git diff
```

---

### 4. Estrutura do Monorepo

```
plenya/
├── apps/
│   ├── web/              # Next.js 16.1
│   ├── mobile/           # Expo SDK 56
│   └── api/              # Go backend
│       ├── cmd/server/
│       ├── internal/
│       │   ├── handlers/      # HTTP handlers (Swagger annotations)
│       │   ├── models/        ← ÚNICA FONTE DE VERDADE
│       │   ├── repository/    # Database queries (Preload, joins)
│       │   ├── services/      # Business logic (DTOs, validação)
│       │   └── middleware/    # Auth, RBAC, WithSelectedPatient
│       ├── database/
│       │   └── migrations/    ← GERADO (não editar)
│       └── docs/
│           └── swagger.json   ← GERADO (não editar)
├── packages/
│   ├── types/
│   │   └── src/
│   │       └── generated/     ← GERADO (não editar)
│   ├── ui/
│   └── api-client/
└── ARQUITETURA.md             # Documentação técnica
```

---

## PARTE 2: BACKEND PATTERNS

### 5. Go Models: Anatomia Completa

**Template obrigatório para todos os models:**

```go
package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type MyModel struct {
    // ID: SEMPRE UUID v7 (time-ordered, RFC 9562)
    // @example 019c1a1e-0579-7f3b-a1bd-4767008e844c
    ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

    // Campos obrigatórios
    // @minLength 3
    // @maxLength 200
    Name string `gorm:"type:varchar(200);not null" json:"name" validate:"required,min=3,max=200"`

    // Campos opcionais (nullable)
    Phone *string `gorm:"type:varchar(20)" json:"phone,omitempty" validate:"omitempty,min=10"`

    // Enums com CHECK constraint
    // @enum active,inactive,archived
    Status string `gorm:"type:varchar(20);not null;check:status IN ('active','inactive','archived')" json:"status"`

    // Foreign Keys com índice
    ParentID uuid.UUID `gorm:"type:uuid;not null;index" json:"parentId"`

    // Relações (GORM Preload)
    Parent *MyParent `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE" json:"parent,omitempty"`
    Children []MyChild `gorm:"foreignKey:ParentID" json:"children,omitempty"`

    // Timestamps automáticos
    CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete
}

// OBRIGATÓRIO: BeforeCreate hook para UUID v7
func (m *MyModel) BeforeCreate(tx *gorm.DB) error {
    if m.ID == uuid.Nil {
        m.ID = uuid.Must(uuid.NewV7())
    }
    return nil
}
```

**UUID v7 (RFC 9562) - OBRIGATÓRIO:**

- **Por quê?** Time-ordered (sortable), melhor performance em B-tree indexes (~20-30% menos fragmentação), rastreabilidade cronológica
- **Aplicação gera:** SEMPRE via `BeforeCreate` hook (não usar default PostgreSQL)
- **Estrutura:** `019c1a1e-0579-7f3b-a1bd-4767008e844c` (primeiros 48 bits = Unix timestamp ms)
- **PostgreSQL fallback:** Função `uuid_generate_v7()` para operações manuais

**Tags GORM essenciais:**

| Tag | Exemplo | Resultado SQL |
|-----|---------|---------------|
| `type:varchar(200)` | - | `VARCHAR(200)` |
| `not null` | - | `NOT NULL` |
| `unique` | - | `UNIQUE` constraint |
| `index` | - | `CREATE INDEX` |
| `check:field IN (...)` | `check:status IN ('a','b')` | `CHECK` constraint |
| `default:'value'` | `default:'active'` | `DEFAULT 'active'` |
| `autoCreateTime` | CreatedAt | Timestamp automático |
| `foreignKey:ParentID` | Relação | JOIN via ParentID |
| `constraint:OnDelete:CASCADE` | Relação | `ON DELETE CASCADE` |

**Tags JSON:**

- `json:"field"` → Campo exposto no JSON
- `json:"field,omitempty"` → Omite se vazio
- `json:"-"` → NUNCA expõe (segurança para CPF, senhas)

**Annotations Swagger (comentários no código):**

- `// @example valor` → Exemplo na documentação
- `// @minLength 3` → Validação mínima
- `// @maxLength 200` → Validação máxima
- `// @minimum 0` → Número mínimo
- `// @maximum 100` → Número máximo
- `// @enum a,b,c` → Valores permitidos
- `// @pattern ^\d{3}\.` → Regex de validação

**Tags de Validação (go-playground/validator):**

- `validate:"required"` → Campo obrigatório
- `validate:"min=3,max=200"` → Tamanho string
- `validate:"gte=0,lte=100"` → Números
- `validate:"email"` → Email válido
- `validate:"oneof=a b c"` → Enum
- `validate:"uuid"` → UUID válido

---

### 6. GORM Hooks

**Hooks disponíveis (ordem de execução):**

1. **BeforeCreate** → Antes de INSERT (gerar UUID, defaults)
2. **BeforeSave** → Antes de INSERT/UPDATE (validação, criptografia)
3. **BeforeUpdate** → Antes de UPDATE (campos automáticos)
4. **AfterFind** → Após SELECT (descriptografia, cálculos)

**Exemplo 1: BeforeCreate - UUID v7 (OBRIGATÓRIO em todos os models)**

```go
func (m *MyModel) BeforeCreate(tx *gorm.DB) error {
    if m.ID == uuid.Nil {
        m.ID = uuid.Must(uuid.NewV7())
    }
    return nil
}
```

**Exemplo 2: BeforeUpdate - LastReview automático (ScoreItem, ScoreLevel)**

Atualiza `LastReview` quando campos clínicos mudam:

```go
func (si *ScoreItem) BeforeUpdate(tx *gorm.DB) error {
    // Verifica se campos clínicos mudaram
    if tx.Statement.Changed("ClinicalRelevance") ||
       tx.Statement.Changed("PatientExplanation") ||
       tx.Statement.Changed("Conduct") {
        now := time.Now()
        si.LastReview = &now
    }
    return nil
}
```

**Exemplo 3: BeforeSave + AfterFind - Criptografia de dados sensíveis (Patient)**

Criptografa CPF/RG antes de salvar, descriptografa após buscar:

```go
func (p *Patient) BeforeSave(tx *gorm.DB) error {
    // Criptografar CPF se não criptografado
    if p.CPF != nil && *p.CPF != "" && !isEncrypted(*p.CPF) {
        encrypted, err := crypto.EncryptWithDefaultKey(*p.CPF)
        if err != nil {
            return err
        }
        p.CPF = &encrypted
    }

    // Criptografar RG se não criptografado
    if p.RG != nil && *p.RG != "" && !isEncrypted(*p.RG) {
        encrypted, err := crypto.EncryptWithDefaultKey(*p.RG)
        if err != nil {
            return err
        }
        p.RG = &encrypted
    }

    // Recalcula idade antes de salvar
    p.CalculateAge()

    return nil
}

func (p *Patient) AfterFind(tx *gorm.DB) error {
    // Descriptografar CPF
    if p.CPF != nil && *p.CPF != "" && isEncrypted(*p.CPF) {
        decrypted, err := crypto.DecryptWithDefaultKey(*p.CPF)
        if err != nil {
            return nil // Manter criptografado se falhar
        }
        p.CPF = &decrypted
    }

    // Descriptografar RG
    if p.RG != nil && *p.RG != "" && isEncrypted(*p.RG) {
        decrypted, err := crypto.DecryptWithDefaultKey(*p.RG)
        if err != nil {
            return nil
        }
        p.RG = &decrypted
    }

    return nil
}
```

**Exemplo 4: Métodos de negócio - Lógica de domínio no model**

**ScoreItem.AppliesToPatient** - Filtros demográficos:

```go
func (si *ScoreItem) AppliesToPatient(patient *Patient) bool {
    // Filtro de gênero
    if si.Gender != nil && *si.Gender != "not_applicable" {
        if *si.Gender != string(patient.Gender) {
            return false
        }
    }

    // Filtro de idade mínima
    if si.AgeRangeMin != nil && patient.Age < *si.AgeRangeMin {
        return false
    }

    // Filtro de idade máxima
    if si.AgeRangeMax != nil && patient.Age > *si.AgeRangeMax {
        return false
    }

    // Filtro de pós-menopausa (apenas para mulheres)
    if patient.Gender == "female" && si.PostMenopause != nil {
        if patient.Menopause == nil {
            return false
        }
        if *si.PostMenopause != *patient.Menopause {
            return false
        }
    }

    return true
}
```

**ScoreLevel.EvaluatesTrue** - Avaliação de operadores:

```go
func (sl *ScoreLevel) EvaluatesTrue(value float64) bool {
    parseFloat := func(s *string) float64 {
        if s == nil {
            return 0
        }
        var result float64
        _, _ = fmt.Sscanf(*s, "%f", &result)
        return result
    }

    switch sl.Operator {
    case "=":
        return value == parseFloat(sl.LowerLimit)
    case ">":
        return value > parseFloat(sl.LowerLimit)
    case ">=":
        return value >= parseFloat(sl.LowerLimit)
    case "<":
        return value < parseFloat(sl.LowerLimit)
    case "<=":
        return value <= parseFloat(sl.LowerLimit)
    case "between":
        lower := parseFloat(sl.LowerLimit)
        upper := parseFloat(sl.UpperLimit)
        return value >= lower && value <= upper
    default:
        return false
    }
}
```

**Regras:**
- ✅ Hooks para geração automática (UUID, timestamps)
- ✅ Hooks para criptografia de dados sensíveis (CPF, RG)
- ✅ Hooks para campos calculados (Age, LastReview)
- ✅ Métodos de negócio para lógica de domínio complexa
- ❌ NÃO colocar lógica de apresentação (formatação de strings)
- ❌ NÃO fazer chamadas HTTP ou operações lentas

---

### 7. Service Layer

**Responsabilidades:**
1. Receber DTOs (Data Transfer Objects) do handler
2. Validar regras de negócio complexas
3. Converter DTOs → Models
4. Chamar repository para persistência
5. Retornar models ao handler

**DTOs (separação API contract vs model interno):**

```go
// CreateScoreItemDTO - O que a API recebe
type CreateScoreItemDTO struct {
    Name           string     `json:"name" validate:"required,min=2,max=300"`
    Unit           *string    `json:"unit,omitempty" validate:"omitempty,max=50"`
    UnitConversion *string    `json:"unitConversion,omitempty"`
    Gender         *string    `json:"gender,omitempty" validate:"omitempty,oneof=not_applicable male female"`
    AgeRangeMin    *int       `json:"ageRangeMin,omitempty" validate:"omitempty,gte=0,lte=150"`
    AgeRangeMax    *int       `json:"ageRangeMax,omitempty" validate:"omitempty,gte=0,lte=150"`
    PostMenopause  *bool      `json:"postMenopause,omitempty"`
    Points         *float64   `json:"points,omitempty" validate:"omitempty,gte=0,lte=100"`
    SubgroupID     uuid.UUID  `json:"subgroupId" validate:"required"`
    ParentItemID   *uuid.UUID `json:"parentItemId,omitempty" validate:"omitempty,uuid"`
    Order          *int       `json:"order,omitempty"`
}

// UpdateScoreItemDTO - Todos os campos opcionais
type UpdateScoreItemDTO struct {
    Name               *string     `json:"name,omitempty" validate:"omitempty,min=2,max=300"`
    Unit               *string     `json:"unit,omitempty" validate:"omitempty,max=50"`
    Gender             *string     `json:"gender,omitempty" validate:"omitempty,oneof=not_applicable male female"`
    AgeRangeMin        *int        `json:"ageRangeMin,omitempty" validate:"omitempty,gte=0,lte=150"`
    AgeRangeMax        *int        `json:"ageRangeMax,omitempty" validate:"omitempty,gte=0,lte=150"`
    ClinicalRelevance  *string     `json:"clinicalRelevance,omitempty"`
    PatientExplanation *string     `json:"patientExplanation,omitempty"`
    Conduct            *string     `json:"conduct,omitempty"`
    Points             *float64    `json:"points,omitempty" validate:"omitempty,gte=0,lte=100"`
    Order              *int        `json:"order,omitempty" validate:"omitempty,gte=0,lte=9999"`
    SubgroupID         *uuid.UUID  `json:"subgroupId,omitempty"`
    ParentItemID       *uuid.UUID  `json:"parentItemId,omitempty"`
}
```

**Exemplo completo de CreateItem:**

```go
func (s *ScoreService) CreateItem(dto CreateScoreItemDTO) (*models.ScoreItem, error) {
    // 1. Validar foreign keys
    _, err := s.repo.GetScoreSubgroupByID(dto.SubgroupID)
    if err != nil {
        return nil, errors.New("parent subgroup not found")
    }

    if dto.ParentItemID != nil {
        _, err := s.repo.GetScoreItemByID(*dto.ParentItemID)
        if err != nil {
            return nil, errors.New("parent item not found")
        }
    }

    // 2. Auto-incrementar order se não fornecido
    order := 0
    if dto.Order != nil {
        order = *dto.Order
    } else {
        maxOrder, err := s.repo.GetMaxOrderForItem(dto.SubgroupID)
        if err != nil {
            return nil, err
        }
        order = maxOrder + 1
    }

    // 3. Converter DTO → Model
    item := &models.ScoreItem{
        Name:           dto.Name,
        Unit:           dto.Unit,
        UnitConversion: dto.UnitConversion,
        Gender:         dto.Gender,
        AgeRangeMin:    dto.AgeRangeMin,
        AgeRangeMax:    dto.AgeRangeMax,
        PostMenopause:  dto.PostMenopause,
        Points:         dto.Points,
        SubgroupID:     dto.SubgroupID,
        ParentItemID:   dto.ParentItemID,
        Order:          order,
    }

    // 4. Salvar (hooks automáticos executam aqui)
    if err := s.repo.CreateScoreItem(item); err != nil {
        return nil, err
    }

    // 5. Recarregar com relações (Preload)
    item, err = s.repo.GetScoreItemByID(item.ID)
    if err != nil {
        return nil, err
    }

    return item, nil
}
```

**Exemplo de UpdateItem (partial update):**

```go
func (s *ScoreService) UpdateItem(id uuid.UUID, dto UpdateScoreItemDTO) (*models.ScoreItem, error) {
    // 1. Buscar item existente
    item, err := s.repo.GetScoreItemByID(id)
    if err != nil {
        return nil, err
    }

    // 2. Atualizar apenas campos fornecidos (partial update)
    if dto.Name != nil {
        item.Name = *dto.Name
    }
    if dto.Points != nil {
        item.Points = dto.Points
    }
    if dto.Order != nil {
        item.Order = *dto.Order
    }
    if dto.ClinicalRelevance != nil {
        item.ClinicalRelevance = dto.ClinicalRelevance
    }
    // ... outros campos

    // 3. ParentItemID especial (pode ser nil para "unindent")
    if dto.Order != nil {
        // Se order mudou, parentItemId também foi enviado (drag-drop)
        item.ParentItemID = dto.ParentItemID
    }

    // 4. Salvar (BeforeUpdate hook executa aqui)
    if err := s.repo.UpdateScoreItem(item); err != nil {
        return nil, err
    }

    // 5. Recarregar com relações
    item, err = s.repo.GetScoreItemByID(item.ID)
    if err != nil {
        return nil, err
    }

    return item, nil
}
```

**Regras:**
- ✅ Validar foreign keys antes de criar
- ✅ Auto-incrementar `order` se não fornecido
- ✅ Sempre recarregar após salvar (para obter relações Preload)
- ✅ Update: apenas campos fornecidos (partial update)
- ❌ NÃO fazer queries diretas ao DB (usar repository)
- ❌ NÃO retornar erros de DB diretamente (encapsular)

---

### 8. Repository Pattern

**Responsabilidades:**
- Queries GORM (Preload, joins, filters)
- CRUD básico (Create, Read, Update, Delete)
- Queries complexas (GetMaxOrder, GetTree)

**Exemplo 1: GetScoreItemByID - Preload de relações:**

```go
func (r *ScoreRepository) GetScoreItemByID(id uuid.UUID) (*models.ScoreItem, error) {
    var item models.ScoreItem
    err := r.db.
        Preload("Subgroup").
        Preload("Subgroup.Group").
        Preload("ParentItem").
        Preload("ChildItems").
        Preload("Levels", func(db *gorm.DB) *gorm.DB {
            return db.Order("level ASC")
        }).
        First(&item, "id = ?", id).Error

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("score item not found")
        }
        return nil, err
    }
    return &item, nil
}
```

**Exemplo 2: GetAllScoreGroupTrees - Nested Preload:**

```go
func (r *ScoreRepository) GetAllScoreGroupTrees() ([]models.ScoreGroup, error) {
    var groups []models.ScoreGroup
    err := r.db.
        Preload("Subgroups", func(db *gorm.DB) *gorm.DB {
            return db.Order("\"order\" ASC, name ASC")
        }).
        Preload("Subgroups.Items", func(db *gorm.DB) *gorm.DB {
            return db.Order("\"order\" ASC, name ASC")
        }).
        Preload("Subgroups.Items.Levels", func(db *gorm.DB) *gorm.DB {
            return db.Order("level ASC")
        }).
        Preload("Subgroups.Items.ParentItem").
        Preload("Subgroups.Items.ChildItems", func(db *gorm.DB) *gorm.DB {
            return db.Order("\"order\" ASC, name ASC")
        }).
        Preload("Subgroups.Items.ChildItems.Levels", func(db *gorm.DB) *gorm.DB {
            return db.Order("level ASC")
        }).
        Order("\"order\" ASC, name ASC").
        Find(&groups).Error

    return groups, err
}
```

**Exemplo 3: UpdateScoreItem - Atualizar nil fields:**

```go
func (r *ScoreRepository) UpdateScoreItem(item *models.ScoreItem) error {
    // GORM Updates() ignora nil pointer fields mesmo com Select("*")
    // Use map para explicitamente setar todos os campos incluindo nil
    updates := map[string]interface{}{
        "name":                item.Name,
        "lab_test_code":       item.LabTestCode,
        "unit":                item.Unit,
        "gender":              item.Gender,
        "age_range_min":       item.AgeRangeMin,
        "age_range_max":       item.AgeRangeMax,
        "clinical_relevance":  item.ClinicalRelevance,
        "patient_explanation": item.PatientExplanation,
        "conduct":             item.Conduct,
        "points":              item.Points,
        "order":               item.Order,
        "subgroup_id":         item.SubgroupID,
        "parent_item_id":      item.ParentItemID, // Pode ser nil (unindent)
    }

    return r.db.Model(&models.ScoreItem{}).Where("id = ?", item.ID).Updates(updates).Error
}
```

**Exemplo 4: GetMaxOrderForItem - Queries auxiliares:**

```go
func (r *ScoreRepository) GetMaxOrderForItem(subgroupID uuid.UUID) (int, error) {
    var maxOrder int
    err := r.db.Model(&models.ScoreItem{}).
        Where("subgroup_id = ?", subgroupID).
        Select("COALESCE(MAX(\"order\"), -1)").
        Scan(&maxOrder).Error
    return maxOrder, err
}
```

**Regras:**
- ✅ Usar `Preload` para carregar relações (evita N+1)
- ✅ Ordenar com `func(db *gorm.DB) *gorm.DB` inline
- ✅ PostgreSQL: escapar `"order"` (palavra reservada)
- ✅ Update nil fields: usar `map[string]interface{}`
- ✅ Queries auxiliares: `Select("COALESCE(MAX(...))")`
- ❌ NÃO fazer lógica de negócio (mover para service)
- ❌ NÃO retornar erros customizados (deixar para service)

---

### 9. Handlers

**Responsabilidades:**
- Receber request HTTP (Fiber context)
- Validar DTO com go-playground/validator
- Chamar service
- Retornar response JSON

**Exemplo completo de CreateScoreItem:**

```go
// CreateScoreItem godoc
// @Summary Create a new score item
// @Description Create a new score item (admin only)
// @Tags Score Items
// @Accept json
// @Produce json
// @Param body body services.CreateScoreItemDTO true "Score Item Data"
// @Success 201 {object} models.ScoreItem
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-items [post]
func (h *ScoreHandler) CreateScoreItem(c *fiber.Ctx) error {
    // 1. Parse request body
    var dto services.CreateScoreItemDTO
    if err := c.BodyParser(&dto); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    // 2. Validar DTO
    if err := h.validator.Struct(dto); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // 3. Chamar service
    item, err := h.service.CreateItem(dto)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // 4. Retornar 201 Created
    return c.Status(http.StatusCreated).JSON(item)
}
```

**Exemplo de UpdateScoreItem:**

```go
// UpdateScoreItem godoc
// @Summary Update a score item
// @Description Update an existing score item (admin only)
// @Tags Score Items
// @Accept json
// @Produce json
// @Param id path string true "Score Item ID (UUID)"
// @Param body body services.UpdateScoreItemDTO true "Score Item Update Data"
// @Success 200 {object} models.ScoreItem
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-items/{id} [put]
func (h *ScoreHandler) UpdateScoreItem(c *fiber.Ctx) error {
    // 1. Parse path param
    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid item ID",
        })
    }

    // 2. Parse request body
    var dto services.UpdateScoreItemDTO
    if err := c.BodyParser(&dto); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    // 3. Validar DTO
    if err := h.validator.Struct(dto); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // 4. Chamar service
    item, err := h.service.UpdateItem(id, dto)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // 5. Retornar 200 OK
    return c.JSON(item)
}
```

**Regras:**
- ✅ Swagger annotations completas (@Summary, @Description, @Tags, @Param, @Success, @Failure, @Security)
- ✅ Validar DTO com `h.validator.Struct(dto)`
- ✅ Status codes corretos (201 Created, 204 No Content, 400 Bad Request, 404 Not Found, 500 Internal Server Error)
- ✅ Retornar erros em formato `fiber.Map{"error": "mensagem"}`
- ❌ NÃO fazer lógica de negócio (mover para service)
- ❌ NÃO fazer queries DB diretas (usar service)

---

### 10. Middleware

**Auth Middleware - JWT validation:**

```go
func Auth(cfg *config.Config) fiber.Handler {
    return func(c *fiber.Ctx) error {
        // 1. Extrair token do header Authorization
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "missing authorization header",
            })
        }

        // 2. Formato esperado: "Bearer <token>"
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "invalid authorization format",
            })
        }

        tokenString := parts[1]

        // 3. Validar token JWT
        token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
            return []byte(cfg.JWT.Secret), nil
        })

        if err != nil || !token.Valid {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "invalid token",
            })
        }

        claims, ok := token.Claims.(*AuthClaims)
        if !ok {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "invalid token claims",
            })
        }

        // 4. Armazenar informações no contexto
        userID, _ := uuid.Parse(claims.UserID)
        c.Locals("userID", userID)
        c.Locals("userEmail", claims.Email)
        c.Locals("userRoles", claims.Roles)

        return c.Next()
    }
}
```

**RBAC Middleware - Role-Based Access Control:**

```go
func RequireRole(allowedRoles ...models.Role) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userRoles := c.Locals("userRoles").([]string)

        // Verificar se o usuário tem pelo menos uma das roles permitidas
        for _, userRole := range userRoles {
            for _, allowedRole := range allowedRoles {
                if userRole == string(allowedRole) {
                    return c.Next()
                }
            }
        }

        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "error": "insufficient permissions",
        })
    }
}

// Helpers específicos
func RequireAdmin() fiber.Handler {
    return RequireRole(models.RoleAdmin)
}

func RequireMedicalStaff() fiber.Handler {
    return RequireRole(models.RoleAdmin, models.RoleDoctor, models.RoleNurse)
}
```

**Uso em rotas:**

```go
// Rotas protegidas por Auth + RBAC
router.Post("/api/v1/score-items",
    middleware.Auth(cfg),
    middleware.RequireAdmin(),
    handler.CreateScoreItem)

router.Get("/api/v1/patients",
    middleware.Auth(cfg),
    middleware.RequireMedicalStaff(),
    handler.GetPatients)
```

---

## PARTE 3: SISTEMA DE ESCORES ⭐

### 11. Sistema de Escores: Visão Geral

**O Sistema de Escores é o coração clínico do Plenya.** Permite estratificação de risco baseada em exames laboratoriais, sinais vitais e parâmetros clínicos.

**Hierarquia (4 níveis):**

```
ScoreGroup (ex: "Exames", "Cardiovascular")
  └─ ScoreSubgroup (ex: "Hemograma", "Função Cardíaca")
      └─ ScoreItem (ex: "Hemoglobina - Homens", "FEVE")
          └─ ScoreLevel (ex: "Nível 5: 55-70% (Ótimo)")
```

**Casos de Uso:**

1. **Avaliação Cardiovascular:**
   - Grupo: Cardiovascular
   - Subgrupo: Função Cardíaca
   - Item: FEVE (Fração de Ejeção do Ventrículo Esquerdo)
   - Levels:
     - Nível 6: FEVE ≥ 70% (Ótimo)
     - Nível 5: FEVE 55-69% (Normal)
     - Nível 3: FEVE 40-54% (Disfunção leve)
     - Nível 1: FEVE < 40% (Disfunção grave)

2. **Hemograma com filtros demográficos:**
   - Item: "Hemoglobina - Homens" (gender=male, age 18-65)
   - Item: "Hemoglobina - Mulheres pré-menopausa" (gender=female, postMenopause=false)
   - Item: "Hemoglobina - Mulheres pós-menopausa" (gender=female, postMenopause=true)

**Relações:**

- **1:N:** Group → Subgroups → Items → Levels
- **M:N:** ScoreItem ↔ Articles (artigos científicos de referência)
- **Self-referencing:** ScoreItem → ParentItem (hierarquia de itens)

**Campos Clínicos (enriquecimento):**

Todos os Items e Levels têm 3 campos de texto enriquecido:
- `ClinicalRelevance` - Explicação técnica para profissionais
- `PatientExplanation` - Explicação em linguagem simples para pacientes
- `Conduct` - Conduta clínica recomendada
- `LastReview` - Auto-updated quando campos clínicos mudam (BeforeUpdate hook)

---

### 12. ScoreGroup

**Estrutura:**

```go
type ScoreGroup struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name      string    `gorm:"type:varchar(200);not null" json:"name"`
    Order     int       `gorm:"type:integer;not null;default:0;index" json:"order"`

    // Relationships
    Subgroups []ScoreSubgroup `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE" json:"subgroups,omitempty"`

    CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
```

**Exemplo de dados:**

```json
{
  "id": "019c1a1e-0579-7f3b-a1bd-4767008e844c",
  "name": "Cardiovascular",
  "order": 1,
  "subgroups": [...]
}
```

**Endpoints:**

- `GET /api/v1/score-groups` - Listar todos
- `GET /api/v1/score-groups/:id` - Detalhes
- `GET /api/v1/score-groups/:id/tree` - Com nested data (subgroups, items, levels)
- `GET /api/v1/score-groups/tree` - Todos com nested data (para mindmap frontend)
- `POST /api/v1/score-groups` - Criar (admin)
- `PUT /api/v1/score-groups/:id` - Atualizar (admin)
- `DELETE /api/v1/score-groups/:id` - Soft delete (admin)

---

### 13. ScoreSubgroup

**Estrutura:**

```go
type ScoreSubgroup struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name      string    `gorm:"type:varchar(200);not null" json:"name"`
    Order     int       `gorm:"type:integer;not null;default:0;index" json:"order"`
    MaxSelect int       `gorm:"type:integer;not null;default:0" json:"maxSelect"` // Para formulários

    // Foreign Keys
    GroupID uuid.UUID `gorm:"type:uuid;not null;index" json:"groupId"`

    // Relationships
    Group *ScoreGroup `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE" json:"group,omitempty"`
    Items []ScoreItem `gorm:"foreignKey:SubgroupID;constraint:OnDelete:CASCADE" json:"items,omitempty"`

    CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
```

**Campo MaxSelect:**
- Indica quantos items podem ser selecionados simultaneamente em formulários
- 0 = sem limite
- 1 = radio button (apenas um)
- N > 1 = checkbox com limite

**Endpoints:**

- `GET /api/v1/score-groups/:groupId/subgroups` - Por grupo
- `GET /api/v1/score-subgroups/:id` - Detalhes
- `POST /api/v1/score-subgroups` - Criar (admin)
- `PUT /api/v1/score-subgroups/:id` - Atualizar (admin)
- `DELETE /api/v1/score-subgroups/:id` - Soft delete (admin)

---

### 14. ScoreItem (Model Crítico)

**Estrutura completa:**

```go
type ScoreItem struct {
    ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

    // Ligação com exames laboratoriais
    // Liga ao lab_test_definitions.code para associar resultados
    LabTestCode *string `gorm:"type:varchar(100);index" json:"labTestCode,omitempty"`

    Name string `gorm:"type:varchar(300);not null" json:"name"`
    Unit *string `gorm:"type:varchar(50)" json:"unit,omitempty"`
    UnitConversion *string `gorm:"type:text" json:"unitConversion,omitempty"` // Ex: "1 g/dL = 10 g/L"

    // Filtros demográficos (aplicabilidade)
    Gender        *string `gorm:"type:varchar(20);default:'not_applicable'" json:"gender,omitempty"` // not_applicable, male, female
    AgeRangeMin   *int    `gorm:"type:integer;check:age_range_min >= 0 AND age_range_min <= 150" json:"ageRangeMin,omitempty"`
    AgeRangeMax   *int    `gorm:"type:integer;check:age_range_max >= 0 AND age_range_max <= 150" json:"ageRangeMax,omitempty"`
    PostMenopause *bool   `gorm:"type:boolean" json:"postMenopause,omitempty"` // Apenas para gender=female

    // Enriquecimento clínico (campos de texto enriquecido)
    ClinicalRelevance  *string    `gorm:"type:text" json:"clinicalRelevance,omitempty"`  // Para profissionais
    PatientExplanation *string    `gorm:"type:text" json:"patientExplanation,omitempty"` // Para pacientes
    Conduct            *string    `gorm:"type:text" json:"conduct,omitempty"`            // Conduta clínica
    LastReview         *time.Time `gorm:"type:timestamp" json:"lastReview,omitempty"`    // Auto-updated

    Points *float64 `gorm:"type:double precision" json:"points,omitempty"`
    Order  int      `gorm:"type:integer;not null;default:0;index" json:"order"`

    // Foreign Keys
    SubgroupID   uuid.UUID  `gorm:"type:uuid;not null;index" json:"subgroupId"`
    ParentItemID *uuid.UUID `gorm:"type:uuid;index" json:"parentItemId,omitempty"` // Self-referencing

    // Relationships
    Subgroup   *ScoreSubgroup `gorm:"foreignKey:SubgroupID;constraint:OnDelete:CASCADE" json:"subgroup,omitempty"`
    ParentItem *ScoreItem     `gorm:"foreignKey:ParentItemID;constraint:OnDelete:SET NULL" json:"parentItem,omitempty"`
    ChildItems []ScoreItem    `gorm:"foreignKey:ParentItemID;constraint:OnDelete:SET NULL" json:"childItems,omitempty"`
    Levels     []ScoreLevel   `gorm:"foreignKey:ItemID;constraint:OnDelete:CASCADE" json:"levels,omitempty"`

    // Many-to-many: artigos científicos de referência
    Articles []Article `gorm:"many2many:article_score_items;" json:"articles,omitempty"`

    CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
```

**Hooks:**

```go
func (si *ScoreItem) BeforeCreate(tx *gorm.DB) error {
    if si.ID == uuid.Nil {
        si.ID = uuid.Must(uuid.NewV7())
    }
    return nil
}

func (si *ScoreItem) BeforeUpdate(tx *gorm.DB) error {
    // Auto-update LastReview quando campos clínicos mudam
    if tx.Statement.Changed("ClinicalRelevance") ||
       tx.Statement.Changed("PatientExplanation") ||
       tx.Statement.Changed("Conduct") {
        now := time.Now()
        si.LastReview = &now
    }
    return nil
}
```

**Método de negócio AppliesToPatient:**

```go
func (si *ScoreItem) AppliesToPatient(patient *Patient) bool {
    // Filtro de gênero
    if si.Gender != nil && *si.Gender != "not_applicable" {
        if *si.Gender != string(patient.Gender) {
            return false
        }
    }

    // Filtro de idade
    if si.AgeRangeMin != nil && patient.Age < *si.AgeRangeMin {
        return false
    }
    if si.AgeRangeMax != nil && patient.Age > *si.AgeRangeMax {
        return false
    }

    // Filtro de menopausa (apenas mulheres)
    if patient.Gender == "female" && si.PostMenopause != nil {
        if patient.Menopause == nil || *si.PostMenopause != *patient.Menopause {
            return false
        }
    }

    return true
}
```

**Exemplo de dados:**

```json
{
  "id": "019c1a20-1234-7f3b-a1bd-4767008e844c",
  "labTestCode": "PLN585CE3E3",
  "name": "Hemoglobina - Homens",
  "unit": "g/dL",
  "unitConversion": "1 g/dL = 10 g/L",
  "gender": "male",
  "ageRangeMin": 18,
  "ageRangeMax": 65,
  "postMenopause": null,
  "clinicalRelevance": "Valores baixos indicam anemia, associada a fadiga e risco cardiovascular aumentado",
  "patientExplanation": "Hemoglobina transporta oxigênio. Valores baixos podem causar cansaço",
  "conduct": "Investigar causa (deficiência de ferro, B12, folato). Suplementação conforme indicado. Encaminhar ao hematologista se Hb < 10 g/dL",
  "lastReview": "2026-01-25T10:30:00Z",
  "points": 20,
  "order": 1,
  "subgroupId": "019c1a1f-5678-7f3b-a1bd-4767008e844c",
  "levels": [...]
}
```

**Endpoints:**

- `GET /api/v1/score-items` - Listar todos (ordenado por Group > Subgroup > Order)
- `GET /api/v1/score-subgroups/:subgroupId/items` - Por subgrupo
- `GET /api/v1/score-items/:id` - Detalhes com levels
- `POST /api/v1/score-items` - Criar (admin)
- `PUT /api/v1/score-items/:id` - Atualizar (admin)
- `DELETE /api/v1/score-items/:id` - Soft delete (admin)

---

### 15. ScoreLevel

**Estrutura:**

```go
type ScoreLevel struct {
    ID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Level int       `gorm:"type:integer;not null;check:level >= 0 AND level <= 6" json:"level"` // 0-6
    Name  string    `gorm:"type:varchar(500);not null" json:"name"` // Ex: "55 a 70 (Ótimo)"

    // Limites numéricos (operador define como usar)
    LowerLimit *string `gorm:"type:varchar(50)" json:"lowerLimit,omitempty"` // Ex: "55"
    UpperLimit *string `gorm:"type:varchar(50)" json:"upperLimit,omitempty"` // Ex: "70"
    Operator   string  `gorm:"type:varchar(10);not null;check:operator IN ('=','>','>=','<','<=','between')" json:"operator"`

    // Enriquecimento clínico
    ClinicalRelevance  *string    `gorm:"type:text" json:"clinicalRelevance,omitempty"`
    PatientExplanation *string    `gorm:"type:text" json:"patientExplanation,omitempty"`
    Conduct            *string    `gorm:"type:text" json:"conduct,omitempty"`
    LastReview         *time.Time `gorm:"type:timestamp" json:"lastReview,omitempty"`

    // Foreign Keys
    ItemID uuid.UUID `gorm:"type:uuid;not null;index" json:"itemId"`

    // Relationships
    Item *ScoreItem `gorm:"foreignKey:ItemID;constraint:OnDelete:CASCADE" json:"item,omitempty"`

    CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
```

**Operadores e uso de limites:**

| Operator | LowerLimit | UpperLimit | Exemplo | Significado |
|----------|------------|------------|---------|-------------|
| `=` | Obrigatório | - | `lowerLimit="50"` | Valor exatamente 50 |
| `>` | Obrigatório | - | `lowerLimit="70"` | Valor maior que 70 |
| `>=` | Obrigatório | - | `lowerLimit="70"` | Valor ≥ 70 |
| `<` | Obrigatório | - | `lowerLimit="40"` | Valor menor que 40 |
| `<=` | Obrigatório | - | `lowerLimit="40"` | Valor ≤ 40 |
| `between` | Obrigatório | Obrigatório | `lowerLimit="55", upperLimit="70"` | Valor entre 55 e 70 (inclusive) |

**Método de negócio EvaluatesTrue:**

```go
func (sl *ScoreLevel) EvaluatesTrue(value float64) bool {
    parseFloat := func(s *string) float64 {
        if s == nil {
            return 0
        }
        var result float64
        _, _ = fmt.Sscanf(*s, "%f", &result)
        return result
    }

    switch sl.Operator {
    case "=":
        return value == parseFloat(sl.LowerLimit)
    case ">":
        return value > parseFloat(sl.LowerLimit)
    case ">=":
        return value >= parseFloat(sl.LowerLimit)
    case "<":
        return value < parseFloat(sl.LowerLimit)
    case "<=":
        return value <= parseFloat(sl.LowerLimit)
    case "between":
        lower := parseFloat(sl.LowerLimit)
        upper := parseFloat(sl.UpperLimit)
        return value >= lower && value <= upper
    default:
        return false
    }
}
```

**Exemplo de dados (FEVE - Fração de Ejeção):**

```json
[
  {
    "level": 6,
    "name": "≥ 70% (Ótimo)",
    "lowerLimit": "70",
    "upperLimit": null,
    "operator": ">=",
    "clinicalRelevance": "FEVE ≥ 70% indica função cardíaca excelente",
    "patientExplanation": "Seu coração está bombeando sangue muito eficientemente",
    "conduct": "Manter acompanhamento regular"
  },
  {
    "level": 5,
    "name": "55 a 70% (Normal)",
    "lowerLimit": "55",
    "upperLimit": "69.9",
    "operator": "between",
    "clinicalRelevance": "FEVE normal indica função sistólica preservada",
    "patientExplanation": "Resultado normal e saudável",
    "conduct": "Acompanhamento anual com ecocardiograma"
  },
  {
    "level": 3,
    "name": "40 a 54% (Disfunção leve)",
    "lowerLimit": "40",
    "upperLimit": "54.9",
    "operator": "between",
    "clinicalRelevance": "Disfunção sistólica leve, risco aumentado de IC",
    "patientExplanation": "Função cardíaca levemente reduzida",
    "conduct": "Otimizar tratamento, eco em 6 meses"
  },
  {
    "level": 1,
    "name": "< 40% (Disfunção grave)",
    "lowerLimit": "40",
    "upperLimit": null,
    "operator": "<",
    "clinicalRelevance": "Disfunção sistólica grave, alto risco de eventos",
    "patientExplanation": "Função cardíaca significativamente reduzida",
    "conduct": "Encaminhar ao cardiologista urgente, otimizar IC"
  }
]
```

**Hooks:**

```go
func (sl *ScoreLevel) BeforeCreate(tx *gorm.DB) error {
    if sl.ID == uuid.Nil {
        sl.ID = uuid.Must(uuid.NewV7())
    }
    return nil
}

func (sl *ScoreLevel) BeforeUpdate(tx *gorm.DB) error {
    // Auto-update LastReview
    if tx.Statement.Changed("ClinicalRelevance") ||
       tx.Statement.Changed("PatientExplanation") ||
       tx.Statement.Changed("Conduct") {
        now := time.Now()
        sl.LastReview = &now
    }
    return nil
}
```

**Endpoints:**

- `GET /api/v1/score-items/:itemId/levels` - Por item
- `GET /api/v1/score-levels/:id` - Detalhes
- `POST /api/v1/score-levels` - Criar (admin)
- `PUT /api/v1/score-levels/:id` - Atualizar (admin)
- `DELETE /api/v1/score-levels/:id` - Soft delete (admin)

---

### 16. Workflows com Escores

**Workflow 1: Adicionar campo em ScoreItem**

Exemplo: Adicionar campo `ReferenceSource` para fonte de referência clínica.

```bash
# Passo 1: Editar Go model
vim apps/api/internal/models/score_item.go

# Adicionar campo:
# ReferenceSource *string `gorm:"type:text" json:"referenceSource,omitempty"`

# Passo 2: Gerar tudo automaticamente
pnpm generate

# Resultado:
# ✅ Migration SQL gerada: ALTER TABLE score_items ADD COLUMN reference_source TEXT;
# ✅ OpenAPI atualizado com novo campo
# ✅ TypeScript types gerados com ReferenceSource?: string
# ✅ Zod schemas atualizados

# Passo 3: Atualizar DTOs
vim apps/api/internal/services/score_service.go

# Adicionar em CreateScoreItemDTO:
# ReferenceSource *string `json:"referenceSource,omitempty"`

# Adicionar em UpdateScoreItemDTO:
# ReferenceSource *string `json:"referenceSource,omitempty"`

# Passo 4: Atualizar conversão DTO → Model
# Em CreateItem e UpdateItem, adicionar:
# item.ReferenceSource = dto.ReferenceSource

# Passo 5: Frontend já funciona automaticamente!
# Types TypeScript já incluem o campo, formulários podem usar
```

**Workflow 2: Duplicar ScoreItem com níveis**

Caso de uso: Criar "Hemoglobina - Mulheres pós-menopausa" baseado em "Hemoglobina - Mulheres".

Backend (adicionar método no service):

```go
func (s *ScoreService) DuplicateItem(itemID uuid.UUID, newName string, newGender *string, newPostMenopause *bool) (*models.ScoreItem, error) {
    // 1. Buscar item original com níveis
    original, err := s.repo.GetScoreItemByID(itemID)
    if err != nil {
        return nil, err
    }

    // 2. Criar novo item (cópia)
    newItem := &models.ScoreItem{
        Name:           newName,
        Gender:         newGender,
        PostMenopause:  newPostMenopause,
        // Copiar outros campos
        LabTestCode:    original.LabTestCode,
        Unit:           original.Unit,
        UnitConversion: original.UnitConversion,
        AgeRangeMin:    original.AgeRangeMin,
        AgeRangeMax:    original.AgeRangeMax,
        SubgroupID:     original.SubgroupID,
    }

    if err := s.repo.CreateScoreItem(newItem); err != nil {
        return nil, err
    }

    // 3. Duplicar níveis
    for _, level := range original.Levels {
        newLevel := &models.ScoreLevel{
            Level:      level.Level,
            Name:       level.Name,
            LowerLimit: level.LowerLimit,
            UpperLimit: level.UpperLimit,
            Operator:   level.Operator,
            ItemID:     newItem.ID,
        }
        if err := s.repo.CreateScoreLevel(newLevel); err != nil {
            return nil, err
        }
    }

    // 4. Recarregar com relações
    return s.repo.GetScoreItemByID(newItem.ID)
}
```

Handler:

```go
// POST /api/v1/score-items/:id/duplicate
func (h *ScoreHandler) DuplicateScoreItem(c *fiber.Ctx) error {
    id, _ := uuid.Parse(c.Params("id"))

    var dto struct {
        NewName         string  `json:"newName" validate:"required"`
        NewGender       *string `json:"newGender,omitempty"`
        NewPostMenopause *bool  `json:"newPostMenopause,omitempty"`
    }

    if err := c.BodyParser(&dto); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
    }

    item, err := h.service.DuplicateItem(id, dto.NewName, dto.NewGender, dto.NewPostMenopause)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(201).JSON(item)
}
```

---

## PARTE 4: FRONTEND PATTERNS

### 17. UX Frontend: Navegação de Formulário (OBRIGATÓRIO)

**CRÍTICO: TODOS os formulários DEVEM usar `useFormNavigation`**

```tsx
import { useFormNavigation } from '@/lib/use-form-navigation'

function MyForm() {
  const formRef = useRef<HTMLFormElement>(null)

  // OBRIGATÓRIO: aplicar em todos os forms
  useFormNavigation({ formRef, disabled: !isDialogOpen }) // disabled se for Dialog

  return (
    <form ref={formRef} onSubmit={handleSubmit}>
      <Input name="field1" />
      <Input name="field2" />
      <button type="submit">Salvar</button>
    </form>
  )
}
```

**Comportamento esperado:**
- Auto-focus: Foca automaticamente no primeiro campo quando o formulário carrega
- Enter em input → Foca próximo campo
- Enter no último campo → Submete formulário
- Tab continua funcionando normalmente
- Textareas permitem quebra de linha (Enter não navega)

**Regras:**
1. Sempre criar `formRef` com `useRef<HTMLFormElement>(null)`
2. Sempre passar `ref={formRef}` no `<form>`
3. Sempre chamar `useFormNavigation({ formRef })`
4. Se o form está em Dialog/Sheet: `disabled: !isOpen`
5. Se o form delega para outro componente, aplicar o hook NO componente do form
6. Para desabilitar auto-focus (raro): `useFormNavigation({ formRef, autoFocus: false })`

---

### 18. Contexto de Paciente (OBRIGATÓRIO)

**CRÍTICO: TODAS as páginas relacionadas a dados de paciente DEVEM usar `useRequireSelectedPatient`**

**Conceito Fundamental:**

O sistema EMR opera no contexto de UM paciente por vez. O médico/profissional seleciona um paciente e então todas as telas mostram apenas dados desse paciente.

**Entidades que pertencem a um paciente:**
- Anamnesis, Lab Results, Prescriptions, Appointments, Health Scores

**Implementação Obrigatória:**

```tsx
import { useRequireSelectedPatient } from '@/lib/use-require-selected-patient'
import { SelectedPatientHeader } from '@/components/patients/SelectedPatientHeader'

export default function MyPatientDataPage() {
  const { selectedPatient, isLoading } = useRequireSelectedPatient()

  // Se não tiver selectedPatient, o hook redireciona automaticamente
  // para /patients/select e salva a URL atual para voltar depois

  return (
    <div>
      <SelectedPatientHeader /> {/* Mostra qual paciente está selecionado */}

      {/* Seus dados aqui - já filtrados pelo selectedPatient */}
    </div>
  )
}
```

**Backend automático:**

O backend usa middleware `WithSelectedPatient()` que:
- Lê o `selectedPatient` do token JWT
- Preenche automaticamente `patientId` nas criações
- Filtra automaticamente queries pelo `patientId`

**Formulários NÃO devem ter select de paciente:**

❌ ERRADO:
```tsx
<Select value={patientId} onChange={setPatientId}>
  <SelectItem value="123">João Silva</SelectItem>
</Select>
```

✅ CORRETO:
```tsx
// Paciente já está selecionado no contexto global
// Formulário só precisa dos campos específicos
<Input name="chiefComplaint" placeholder="Queixa principal" />
```

**Fluxo de Uso:**

1. Usuário acessa `/lab-requests`
2. Sistema verifica: Tem `selectedPatient`?
   - ❌ Não → Redireciona para `/patients/select`
   - ✅ Sim → Mostra página com dados filtrados
3. Usuário seleciona paciente em `/patients/select`
4. Sistema salva `selectedPatient` no Zustand + localStorage
5. Sistema redireciona de volta para `/lab-requests`
6. Página carrega dados apenas do paciente selecionado

**Exceções (páginas que NÃO precisam de selectedPatient):**
- `/patients` - Lista de todos os pacientes
- `/patients/select` - Seleção de paciente
- `/articles` - Artigos científicos
- `/lab-test-definitions` - Definições de exames (template)
- `/score-groups` - Sistema de escores (não é de paciente)

---

### 19. TanStack Query Patterns

**Query keys hierárquicos:**

```typescript
export const scoreKeys = {
  all: ['scores'] as const,
  groups: () => [...scoreKeys.all, 'groups'] as const,
  group: (id: string) => [...scoreKeys.groups(), id] as const,
  groupTree: (id: string) => [...scoreKeys.group(id), 'tree'] as const,
  allGroupTrees: () => [...scoreKeys.groups(), 'trees'] as const,
  items: () => [...scoreKeys.all, 'items'] as const,
  item: (id: string) => [...scoreKeys.items(), id] as const,
  itemsBySubgroup: (subgroupId: string) => [...scoreKeys.items(), 'subgroup', subgroupId] as const,
  levels: () => [...scoreKeys.all, 'levels'] as const,
  level: (id: string) => [...scoreKeys.levels(), id] as const,
}
```

**Padrão de Queries:**

```typescript
export function useScoreGroups() {
  return useQuery({
    queryKey: scoreKeys.groups(),
    queryFn: () => apiClient.get<ScoreGroup[]>('/api/v1/score-groups'),
    staleTime: 5 * 60 * 1000, // 5 minutes (cache)
  })
}

export function useScoreGroup(id: string) {
  return useQuery({
    queryKey: scoreKeys.group(id),
    queryFn: () => apiClient.get<ScoreGroup>(`/api/v1/score-groups/${id}`),
    enabled: !!id, // Só executa se id estiver presente
  })
}
```

**Mutations com invalidação:**

```typescript
export function useCreateScoreItem() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (data: CreateScoreItemDTO) =>
      apiClient.post<ScoreItem>('/api/v1/score-items', data),
    onSuccess: (newItem) => {
      // Invalidar queries relacionadas (força refetch)
      queryClient.invalidateQueries({ queryKey: scoreKeys.items() })
      queryClient.invalidateQueries({ queryKey: scoreKeys.itemsBySubgroup(newItem.subgroupId) })
      queryClient.invalidateQueries({ queryKey: scoreKeys.allGroupTrees() })
    },
  })
}

export function useUpdateScoreItem() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateScoreItemDTO }) =>
      apiClient.put<ScoreItem>(`/api/v1/score-items/${id}`, data),
    onSuccess: (updatedItem, { id }) => {
      queryClient.invalidateQueries({ queryKey: scoreKeys.item(id) })
      queryClient.invalidateQueries({ queryKey: scoreKeys.items() })
      queryClient.invalidateQueries({ queryKey: scoreKeys.allGroupTrees() })
    },
  })
}
```

**Optimistic Updates (opcional, para melhor UX):**

```typescript
export function useDeleteScoreLevel() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) => apiClient.delete(`/api/v1/score-levels/${id}`),
    onMutate: async (id) => {
      // Cancel outgoing refetches
      await queryClient.cancelQueries({ queryKey: scoreKeys.allGroupTrees() })

      // Snapshot previous value
      const previousData = queryClient.getQueryData<ScoreGroup[]>(scoreKeys.allGroupTrees())

      // Optimistically update (remove level da UI imediatamente)
      queryClient.setQueryData<ScoreGroup[]>(
        scoreKeys.allGroupTrees(),
        (oldData) => {
          if (!oldData) return oldData

          return oldData.map(group => ({
            ...group,
            subgroups: group.subgroups?.map(subgroup => ({
              ...subgroup,
              items: subgroup.items?.map(item => ({
                ...item,
                levels: item.levels?.filter(level => level.id !== id)
              }))
            }))
          }))
        }
      )

      return { previousData }
    },
    onError: (_err, _id, context) => {
      // Rollback on error
      if (context?.previousData) {
        queryClient.setQueryData(scoreKeys.allGroupTrees(), context.previousData)
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: scoreKeys.levels() })
    },
  })
}
```

**Regras:**
- ✅ Query keys hierárquicos (`all` → `groups` → `group(id)`)
- ✅ Invalidar queries relacionadas após mutations
- ✅ `enabled: !!id` para queries condicionais
- ✅ `staleTime` para caching (5 min para listas)
- ✅ Optimistic updates para melhor UX (delete, toggle)
- ❌ NÃO fazer invalidações excessivas (performance)
- ❌ NÃO usar `refetchOnWindowFocus` em tudo (excesso de requests)

---

### 20. API Client

**Auto-refresh token:**

O `apiClient` automaticamente renova o token JWT quando expira (401).

```typescript
class APIClient {
  private isRefreshing = false
  private refreshPromise: Promise<boolean> | null = null

  private async tryRefreshToken(): Promise<boolean> {
    // Se já está fazendo refresh, aguarda a promise existente
    if (this.isRefreshing && this.refreshPromise) {
      return this.refreshPromise
    }

    this.isRefreshing = true
    this.refreshPromise = this._doRefresh()

    try {
      const result = await this.refreshPromise
      return result
    } finally {
      this.isRefreshing = false
      this.refreshPromise = null
    }
  }

  private async _doRefresh(): Promise<boolean> {
    const { refreshToken } = useAuthStore.getState()

    if (!refreshToken) {
      return false
    }

    try {
      const response = await fetch(`${this.baseURL}/api/v1/auth/refresh`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ refreshToken }),
      })

      if (response.ok) {
        const data = await response.json()
        useAuthStore.getState().setAuth(data.user, data.accessToken, data.refreshToken)
        return true
      }
    } catch (error) {
      console.error("Error refreshing token:", error)
    }

    return false
  }

  private async request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
    // 1. Primeira tentativa
    let response = await this.fetchWithAuth(endpoint, options)

    // 2. Se 401, tenta renovar token
    if (response.status === 401 && endpoint !== "/api/v1/auth/refresh") {
      const refreshed = await this.tryRefreshToken()

      if (refreshed) {
        // Retry com novo token
        response = await this.fetchWithAuth(endpoint, options)
      } else {
        // Refresh falhou - fazer logout
        useAuthStore.getState().clearAuth()
        toast.error("Sua sessão expirou")
        setTimeout(() => window.location.href = "/login", 1500)
        throw new Error("Session expired")
      }
    }

    if (!response.ok) {
      const error = await response.json().catch(() => ({ error: "Unknown error" }))
      throw new Error(error.error || error.message || response.statusText)
    }

    // 3. Handle 204 No Content
    if (response.status === 204) {
      return {} as T
    }

    return response.json()
  }
}
```

**Regras:**
- ✅ Auto-refresh token em 401 (exceto endpoint /refresh)
- ✅ Retry request após refresh bem-sucedido
- ✅ Logout automático se refresh falhar
- ✅ Deduplicar refreshes simultâneos (promise única)
- ✅ Suporte FormData (upload de arquivos)
- ❌ NÃO fazer refresh em loop infinito
- ❌ NÃO expor erros de rede diretamente (encapsular)

---

## PARTE 5: DOMÍNIO E SEGURANÇA

### 21. Estrutura de Dados

**Entidades Principais:**

```
Users (base para todos)
  ├─ id, email, password_hash, roles (multi-role RBAC)
  │
  ├─ Patients (extends Users)
  │    └─ cpf_encrypted, name, birth_date, gender, menopause
  │
  ├─ Doctors (extends Users)
  │    └─ crm, specialty
  │
  └─ Nurses (extends Users)

Patients (1:N)
  ├─ Anamnesis (histórico médico)
  ├─ LabResults (exames laboratoriais)
  ├─ LabRequests (pedidos de exames)
  └─ Prescriptions (prescrições)

ScoreGroups (Sistema de Escores)
  └─ ScoreSubgroups (1:N)
      └─ ScoreItems (1:N)
          ├─ ScoreLevels (1:N)
          └─ Articles (M:N) - artigos científicos

LabTestDefinitions (templates de exames)
  └─ LabResults (1:N) - resultados de exames

Articles (artigos científicos)
  └─ ScoreItems (M:N) - relação inversa

AuditLogs (rastreabilidade)
  └─ user_id, action, resource, resource_id, ip, created_at
```

**Regras de Negócio:**

- Health scores: low (0-30), medium (31-60), high (61-100)
- Recalculados quando novos exames adicionados
- Relatórios PDF com timestamp + hash SHA-256
- Soft delete em todas entidades principais
- UUID v7 (time-ordered) em todos os models

---

### 22. Segurança EMR/LGPD (CRÍTICO)

**Autenticação:**
- JWT access token: 15min
- JWT refresh token: 7 dias
- Senhas: Argon2id ou Bcrypt (cost 12+)
- 2FA/TOTP obrigatório para profissionais
- Biometria em mobile apps

**Autorização:**
- RBAC multi-role: admin, manager, doctor, psychologist, nutritionist, nurse, secretary, patient
- Patients só acessam próprios dados
- Doctors acessam pacientes atribuídos
- Admins têm acesso total (com audit log)

**Dados Sensíveis:**
- CPF, RG: criptografados com pgcrypto (BeforeSave/AfterFind hooks)
- `json:"-"` em campos sensíveis (nunca expor)
- Mascaramento em logs
- Backups criptografados

**Audit Logging:**
- TODOS os acessos a prontuários geram log
- Campos: user_id, action, resource, resource_id, ip_address, created_at
- Retenção: mínimo 5 anos
- Logs imutáveis (append-only)

**Rate Limiting:**
- 100 requests/min por IP
- Alertas de brute force

**Implementação via Hooks (Patient model):**

```go
func (p *Patient) BeforeSave(tx *gorm.DB) error {
    // Criptografar CPF se não criptografado
    if p.CPF != nil && *p.CPF != "" && !isEncrypted(*p.CPF) {
        encrypted, _ := crypto.EncryptWithDefaultKey(*p.CPF)
        p.CPF = &encrypted
    }
    return nil
}

func (p *Patient) AfterFind(tx *gorm.DB) error {
    // Descriptografar CPF após buscar
    if p.CPF != nil && *p.CPF != "" && isEncrypted(*p.CPF) {
        decrypted, _ := crypto.DecryptWithDefaultKey(*p.CPF)
        p.CPF = &decrypted
    }
    return nil
}
```

---

## PARTE 6: WORKFLOWS E DESENVOLVIMENTO

### 23. Comandos Essenciais

**Desenvolvimento (via Docker):**

```bash
# Iniciar todos os serviços (PostgreSQL, API Go, Web Next.js)
docker compose up -d

# Reconstruir imagens (necessário quando mudar Dockerfile ou dependências)
docker compose up -d --build

# Ver logs
docker compose logs -f          # Todos os serviços
docker compose logs -f web      # Apenas web
docker compose logs -f api      # Apenas API
docker compose logs -f db       # Apenas database

# Parar serviços
docker compose down

# Instalar dependência no web
docker compose exec web pnpm add <pacote> --filter web
docker compose restart web

# Rodar comandos no container
docker compose exec web pnpm dev --filter web
docker compose exec api go mod tidy
```

**Hot Reload:** Código sincronizado via volumes Docker. Mudanças em `.ts`, `.tsx`, `.go` são detectadas automaticamente.

**Migrations:**

```bash
# Gerar migration do Go model
pnpm generate:migrations

# Aplicar migrations
atlas migrate apply --env dev

# Status
atlas migrate status --env dev
```

---

### 24. Workflows Práticos

**Workflow 1: Criar novo Model do zero**

Exemplo: Criar `ClinicalProtocol` (protocolo clínico).

```bash
# Passo 1: Criar Go model
vim apps/api/internal/models/clinical_protocol.go
```

```go
package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type ClinicalProtocol struct {
    ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name        string    `gorm:"type:varchar(200);not null" json:"name" validate:"required,min=3,max=200"`
    Description *string   `gorm:"type:text" json:"description,omitempty"`
    Active      bool      `gorm:"type:boolean;not null;default:true" json:"active"`

    CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (ClinicalProtocol) TableName() string {
    return "clinical_protocols"
}

func (cp *ClinicalProtocol) BeforeCreate(tx *gorm.DB) error {
    if cp.ID == uuid.Nil {
        cp.ID = uuid.Must(uuid.NewV7())
    }
    return nil
}
```

```bash
# Passo 2: Gerar migrations + OpenAPI + TypeScript
pnpm generate

# Passo 3: Criar repository
vim apps/api/internal/repository/clinical_protocol_repository.go
```

```go
package repository

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
    "github.com/plenya/api/internal/models"
)

type ClinicalProtocolRepository struct {
    db *gorm.DB
}

func NewClinicalProtocolRepository(db *gorm.DB) *ClinicalProtocolRepository {
    return &ClinicalProtocolRepository{db: db}
}

func (r *ClinicalProtocolRepository) Create(protocol *models.ClinicalProtocol) error {
    return r.db.Create(protocol).Error
}

func (r *ClinicalProtocolRepository) GetByID(id uuid.UUID) (*models.ClinicalProtocol, error) {
    var protocol models.ClinicalProtocol
    err := r.db.First(&protocol, "id = ?", id).Error
    return &protocol, err
}

func (r *ClinicalProtocolRepository) GetAll() ([]models.ClinicalProtocol, error) {
    var protocols []models.ClinicalProtocol
    err := r.db.Order("name ASC").Find(&protocols).Error
    return protocols, err
}

func (r *ClinicalProtocolRepository) Update(protocol *models.ClinicalProtocol) error {
    return r.db.Save(protocol).Error
}

func (r *ClinicalProtocolRepository) Delete(id uuid.UUID) error {
    return r.db.Delete(&models.ClinicalProtocol{}, "id = ?", id).Error
}
```

```bash
# Passo 4: Criar service com DTOs
vim apps/api/internal/services/clinical_protocol_service.go
```

```go
package services

import (
    "github.com/google/uuid"
    "github.com/plenya/api/internal/models"
    "github.com/plenya/api/internal/repository"
)

type CreateClinicalProtocolDTO struct {
    Name        string  `json:"name" validate:"required,min=3,max=200"`
    Description *string `json:"description,omitempty"`
    Active      *bool   `json:"active,omitempty"`
}

type UpdateClinicalProtocolDTO struct {
    Name        *string `json:"name,omitempty" validate:"omitempty,min=3,max=200"`
    Description *string `json:"description,omitempty"`
    Active      *bool   `json:"active,omitempty"`
}

type ClinicalProtocolService struct {
    repo *repository.ClinicalProtocolRepository
}

func NewClinicalProtocolService(repo *repository.ClinicalProtocolRepository) *ClinicalProtocolService {
    return &ClinicalProtocolService{repo: repo}
}

func (s *ClinicalProtocolService) Create(dto CreateClinicalProtocolDTO) (*models.ClinicalProtocol, error) {
    active := true
    if dto.Active != nil {
        active = *dto.Active
    }

    protocol := &models.ClinicalProtocol{
        Name:        dto.Name,
        Description: dto.Description,
        Active:      active,
    }

    if err := s.repo.Create(protocol); err != nil {
        return nil, err
    }

    return protocol, nil
}

func (s *ClinicalProtocolService) GetByID(id uuid.UUID) (*models.ClinicalProtocol, error) {
    return s.repo.GetByID(id)
}

func (s *ClinicalProtocolService) GetAll() ([]models.ClinicalProtocol, error) {
    return s.repo.GetAll()
}

func (s *ClinicalProtocolService) Update(id uuid.UUID, dto UpdateClinicalProtocolDTO) (*models.ClinicalProtocol, error) {
    protocol, err := s.repo.GetByID(id)
    if err != nil {
        return nil, err
    }

    if dto.Name != nil {
        protocol.Name = *dto.Name
    }
    if dto.Description != nil {
        protocol.Description = dto.Description
    }
    if dto.Active != nil {
        protocol.Active = *dto.Active
    }

    if err := s.repo.Update(protocol); err != nil {
        return nil, err
    }

    return protocol, nil
}

func (s *ClinicalProtocolService) Delete(id uuid.UUID) error {
    return s.repo.Delete(id)
}
```

```bash
# Passo 5: Criar handler
vim apps/api/internal/handlers/clinical_protocol_handler.go
```

```go
package handlers

import (
    "net/http"
    "github.com/gofiber/fiber/v2"
    "github.com/go-playground/validator/v10"
    "github.com/google/uuid"
    "github.com/plenya/api/internal/services"
)

type ClinicalProtocolHandler struct {
    service   *services.ClinicalProtocolService
    validator *validator.Validate
}

func NewClinicalProtocolHandler(service *services.ClinicalProtocolService, validator *validator.Validate) *ClinicalProtocolHandler {
    return &ClinicalProtocolHandler{service: service, validator: validator}
}

// GetAll godoc
// @Summary List all clinical protocols
// @Tags Clinical Protocols
// @Produce json
// @Success 200 {array} models.ClinicalProtocol
// @Security BearerAuth
// @Router /api/v1/clinical-protocols [get]
func (h *ClinicalProtocolHandler) GetAll(c *fiber.Ctx) error {
    protocols, err := h.service.GetAll()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch protocols"})
    }
    return c.JSON(protocols)
}

// Create godoc
// @Summary Create a new clinical protocol
// @Tags Clinical Protocols
// @Accept json
// @Produce json
// @Param body body services.CreateClinicalProtocolDTO true "Protocol Data"
// @Success 201 {object} models.ClinicalProtocol
// @Security BearerAuth
// @Router /api/v1/clinical-protocols [post]
func (h *ClinicalProtocolHandler) Create(c *fiber.Ctx) error {
    var dto services.CreateClinicalProtocolDTO
    if err := c.BodyParser(&dto); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
    }
    if err := h.validator.Struct(dto); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }
    protocol, err := h.service.Create(dto)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(201).JSON(protocol)
}

// ... outros métodos (GetByID, Update, Delete)
```

```bash
# Passo 6: Registrar rotas
vim apps/api/cmd/server/main.go

# Adicionar no setupRoutes:
# protocolRepo := repository.NewClinicalProtocolRepository(db)
# protocolService := services.NewClinicalProtocolService(protocolRepo)
# protocolHandler := handlers.NewClinicalProtocolHandler(protocolService, validate)
#
# router.Get("/api/v1/clinical-protocols", middleware.Auth(cfg), protocolHandler.GetAll)
# router.Post("/api/v1/clinical-protocols", middleware.Auth(cfg), middleware.RequireAdmin(), protocolHandler.Create)

# Passo 7: Gerar OpenAPI
pnpm generate:swagger

# Passo 8: Frontend - criar queries
vim apps/web/lib/api/clinical-protocol-api.ts
```

```typescript
export interface ClinicalProtocol {
  id: string
  name: string
  description?: string
  active: boolean
  createdAt: string
  updatedAt: string
}

export const clinicalProtocolKeys = {
  all: ['clinical-protocols'] as const,
  list: () => [...clinicalProtocolKeys.all, 'list'] as const,
  detail: (id: string) => [...clinicalProtocolKeys.all, id] as const,
}

export function useClinicalProtocols() {
  return useQuery({
    queryKey: clinicalProtocolKeys.list(),
    queryFn: () => apiClient.get<ClinicalProtocol[]>('/api/v1/clinical-protocols'),
  })
}

export function useCreateClinicalProtocol() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (data: CreateClinicalProtocolDTO) =>
      apiClient.post<ClinicalProtocol>('/api/v1/clinical-protocols', data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: clinicalProtocolKeys.list() })
    },
  })
}
```

```bash
# Passo 9: Criar página
vim apps/web/app/(authenticated)/clinical-protocols/page.tsx
```

**Checklist completo:**
- [x] Criar Go model com BeforeCreate hook
- [x] Rodar `pnpm generate`
- [x] Criar repository
- [x] Criar service com DTOs
- [x] Criar handler com Swagger annotations
- [x] Registrar rotas
- [x] Gerar OpenAPI
- [x] Criar queries frontend
- [x] Criar página

---

**Workflow 2: Adicionar endpoint em model existente**

Exemplo: `GET /api/v1/score-items/search?query=hemoglobina`

```bash
# Passo 1: Adicionar método no repository
vim apps/api/internal/repository/score_repository.go
```

```go
func (r *ScoreRepository) SearchItems(query string) ([]models.ScoreItem, error) {
    var items []models.ScoreItem
    err := r.db.
        Where("name ILIKE ?", "%"+query+"%").
        Preload("Subgroup").
        Preload("Levels", func(db *gorm.DB) *gorm.DB {
            return db.Order("level ASC")
        }).
        Order("name ASC").
        Limit(20).
        Find(&items).Error
    return items, err
}
```

```bash
# Passo 2: Adicionar método no service
vim apps/api/internal/services/score_service.go
```

```go
func (s *ScoreService) SearchItems(query string) ([]models.ScoreItem, error) {
    if len(query) < 2 {
        return nil, errors.New("query must be at least 2 characters")
    }
    return s.repo.SearchItems(query)
}
```

```bash
# Passo 3: Adicionar handler
vim apps/api/internal/handlers/score_handler.go
```

```go
// SearchScoreItems godoc
// @Summary Search score items by name
// @Tags Score Items
// @Param query query string true "Search query (min 2 chars)"
// @Success 200 {array} models.ScoreItem
// @Security BearerAuth
// @Router /api/v1/score-items/search [get]
func (h *ScoreHandler) SearchScoreItems(c *fiber.Ctx) error {
    query := c.Query("query")
    if query == "" {
        return c.Status(400).JSON(fiber.Map{"error": "query parameter required"})
    }

    items, err := h.service.SearchItems(query)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(items)
}
```

```bash
# Passo 4: Registrar rota
vim apps/api/cmd/server/main.go

# Adicionar:
# router.Get("/api/v1/score-items/search", middleware.Auth(cfg), scoreHandler.SearchScoreItems)

# Passo 5: Gerar OpenAPI
pnpm generate:swagger

# Passo 6: Frontend - criar query
vim apps/web/lib/api/score-api.ts
```

```typescript
export function useSearchScoreItems(query: string) {
  return useQuery({
    queryKey: [...scoreKeys.items(), 'search', query] as const,
    queryFn: () => apiClient.get<ScoreItem[]>(`/api/v1/score-items/search?query=${encodeURIComponent(query)}`),
    enabled: query.length >= 2,
    staleTime: 30 * 1000, // 30 seconds
  })
}
```

---

### 25. Variáveis de Ambiente

```bash
# Backend (apps/api/.env)
PORT=3001
ENVIRONMENT=development
DB_HOST=localhost
DB_PORT=5432
DB_USER=plenya_user
DB_PASSWORD=senha_segura
DB_NAME=plenya_db
JWT_SECRET=jwt_secret_256bits
JWT_ACCESS_EXPIRY=15m
JWT_REFRESH_EXPIRY=168h
ENCRYPTION_KEY=encryption_key_256bits
CORS_ORIGIN=http://localhost:3000

# Frontend (apps/web/.env.local)
NEXT_PUBLIC_API_URL=http://localhost:3001
```

---

### 26. Endpoints API (Principais)

**Auth:**
- `POST /api/v1/auth/register` - Criar conta
- `POST /api/v1/auth/login` - Login
- `POST /api/v1/auth/refresh` - Renovar token
- `POST /api/v1/auth/2fa/enable` - Habilitar 2FA

**Patients:**
- `GET /api/v1/patients` - Listar (paginado)
- `GET /api/v1/patients/:id` - Detalhes
- `POST /api/v1/patients` - Criar
- `PUT /api/v1/patients/:id` - Atualizar
- `DELETE /api/v1/patients/:id` - Soft delete

**Score Groups:**
- `GET /api/v1/score-groups` - Listar todos
- `GET /api/v1/score-groups/:id` - Detalhes
- `GET /api/v1/score-groups/:id/tree` - Com nested data
- `GET /api/v1/score-groups/tree` - Todos com nested data (mindmap)
- `POST /api/v1/score-groups` - Criar (admin)
- `PUT /api/v1/score-groups/:id` - Atualizar (admin)
- `DELETE /api/v1/score-groups/:id` - Soft delete (admin)

**Score Items:**
- `GET /api/v1/score-items` - Listar todos
- `GET /api/v1/score-subgroups/:subgroupId/items` - Por subgrupo
- `GET /api/v1/score-items/:id` - Detalhes com levels
- `POST /api/v1/score-items` - Criar (admin)
- `PUT /api/v1/score-items/:id` - Atualizar (admin)
- `DELETE /api/v1/score-items/:id` - Soft delete (admin)

**Score Levels:**
- `GET /api/v1/score-items/:itemId/levels` - Por item
- `GET /api/v1/score-levels/:id` - Detalhes
- `POST /api/v1/score-levels` - Criar (admin)
- `PUT /api/v1/score-levels/:id` - Atualizar (admin)
- `DELETE /api/v1/score-levels/:id` - Soft delete (admin)

**Anamnesis:**
- `GET /api/v1/patients/:id/anamnesis` - Histórico
- `POST /api/v1/anamnesis` - Criar anamnese

**Labs:**
- `GET /api/v1/patients/:id/labs` - Exames
- `POST /api/v1/labs` - Registrar exame

**Audit:**
- `GET /api/v1/audit-logs` - Listar (admin only)

---

### 27. Performance Targets

- Latência p95: < 500ms (rotas simples), < 2s (PDFs)
- Error rate: < 0.5%
- Uptime: > 99.5%
- Database queries: < 100ms (p95)

---

### 28. Custos Mensais

| Item | Custo |
|------|-------|
| VPS Hetzner CPX31 | €9 (~R$50) |
| Backup Snapshots | €2 (~R$11) |
| Backup S3 Glacier | ~R$5 |
| Domínio .com.br | ~R$3 |
| **Total** | **~R$69/mês** |

Opcionais:
- Expo EAS Updates: $99 USD
- App Stores: ~R$50/mês

---

### 29. Roadmap

- [x] **Fase 1:** Fundação (monorepo + Docker) ✅
- [x] **Fase 2:** Backend Core (Auth + RBAC + Migrations) ✅
- [x] **Fase 3:** Features Médicas (Anamnesis + Appointments + Prescriptions + Lab Results) ✅
- [x] **Fase 4 Parte 1:** Frontend Web - Fundação (Next.js 16.1 + shadcn/ui + TanStack Query + Login) ✅
- [x] **Fase 4 Parte 2:** Frontend Web - Dashboard e CRUD Completo + Sistema de Escores ✅
- [ ] **Fase 5:** Mobile Apps (Expo + React Native)
- [ ] **Fase 6:** Hardening LGPD (Relatórios + Exportação)
- [ ] **Fase 7:** Deploy Produção (Hetzner + Coolify)

---

### 30. Observações Importantes

**Desenvolvimento:**
1. **Nunca commitar .env** - Usar .env.example
2. **Go models são única fonte** - Nunca editar arquivos gerados
3. **Testar geração** - Rodar `pnpm generate` após mudar models
4. **UUID v7 obrigatório** - TODOS os models DEVEM ter BeforeCreate hook
5. **useFormNavigation obrigatório** - TODOS os formulários DEVEM usar
6. **useRequireSelectedPatient obrigatório** - TODAS as páginas de dados de paciente DEVEM usar

**Segurança:**
7. **Audit logs são imutáveis** - Nunca deletar
8. **Criptografia obrigatória** - CPF, RG, documentos (BeforeSave/AfterFind hooks)
9. **2FA obrigatório** - Profissionais de saúde
10. **LGPD é lei** - Compliance não opcional

**Sistema de Escores:**
11. **ScoreItem.AppliesToPatient** - Filtros demográficos (gender, age, postMenopause)
12. **ScoreLevel.EvaluatesTrue** - Operadores de comparação (=, >, >=, <, <=, between)
13. **LastReview auto-updated** - BeforeUpdate hook quando campos clínicos mudam
14. **M2N Articles ↔ ScoreItems** - Relação bidirecional para artigos científicos

---

**Última atualização:** Fevereiro 2026 (v2.0 - Sistema de Escores completo)
