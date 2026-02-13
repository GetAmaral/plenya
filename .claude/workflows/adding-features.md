# Workflows - Adding Features

## Workflow Completo: Criar Novo Model do Zero

Exemplo: Criar `ClinicalProtocol` (protocolo clínico).

### Passo 1: Criar Go Model

```bash
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

### Passo 2: Gerar Migrations + OpenAPI + TypeScript

```bash
pnpm generate
```

### Passo 3: Criar Repository

```bash
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

### Passo 4: Criar Service com DTOs

```bash
vim apps/api/internal/services/clinical_protocol_service.go
```

```go
package services

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
```

### Passo 5: Criar Handler

```bash
vim apps/api/internal/handlers/clinical_protocol_handler.go
```

```go
package handlers

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
        return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch"})
    }
    return c.JSON(protocols)
}

// Create godoc
// @Summary Create a new clinical protocol
// @Tags Clinical Protocols
// @Accept json
// @Produce json
// @Param body body services.CreateClinicalProtocolDTO true "Data"
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
```

### Passo 6: Registrar Rotas

```bash
vim apps/api/cmd/server/main.go
```

```go
// Adicionar no setupRoutes:
protocolRepo := repository.NewClinicalProtocolRepository(db)
protocolService := services.NewClinicalProtocolService(protocolRepo)
protocolHandler := handlers.NewClinicalProtocolHandler(protocolService, validate)

router.Get("/api/v1/clinical-protocols", middleware.Auth(cfg), protocolHandler.GetAll)
router.Post("/api/v1/clinical-protocols", middleware.Auth(cfg), middleware.RequireAdmin(), protocolHandler.Create)
```

### Passo 7: Gerar OpenAPI

```bash
pnpm generate:swagger
```

### Passo 8: Frontend - Criar Queries

```bash
vim apps/web/lib/api/clinical-protocol-api.ts
```

```typescript
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

### Passo 9: Criar Página

```bash
vim apps/web/app/(authenticated)/clinical-protocols/page.tsx
```

## Checklist Completo

- [ ] Criar Go model com BeforeCreate hook
- [ ] Rodar `pnpm generate`
- [ ] Criar repository
- [ ] Criar service com DTOs
- [ ] Criar handler com Swagger annotations
- [ ] Registrar rotas
- [ ] Gerar OpenAPI
- [ ] Criar queries frontend
- [ ] Criar página

## Workflow: Adicionar Endpoint em Model Existente

Exemplo: `GET /api/v1/score-items/search?query=hemoglobina`

### 1. Repository

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

### 2. Service

```go
func (s *ScoreService) SearchItems(query string) ([]models.ScoreItem, error) {
    if len(query) < 2 {
        return nil, errors.New("query must be at least 2 characters")
    }
    return s.repo.SearchItems(query)
}
```

### 3. Handler

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
        return c.Status(400).JSON(fiber.Map{"error": "query required"})
    }

    items, err := h.service.SearchItems(query)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(items)
}
```

### 4. Rota

```go
router.Get("/api/v1/score-items/search", middleware.Auth(cfg), scoreHandler.SearchScoreItems)
```

### 5. Gerar OpenAPI

```bash
pnpm generate:swagger
```

### 6. Frontend Query

```typescript
export function useSearchScoreItems(query: string) {
  return useQuery({
    queryKey: [...scoreKeys.items(), 'search', query] as const,
    queryFn: () => apiClient.get<ScoreItem[]>(`/api/v1/score-items/search?query=${encodeURIComponent(query)}`),
    enabled: query.length >= 2,
    staleTime: 30 * 1000,
  })
}
```

Ver também:
- `database-ops.md` - Para manipular dados manualmente
- `score-system.md` - Para entender hierarquia de escores
- `models.md` - Para padrões de Go models
