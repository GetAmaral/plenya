# Backend - Service Layer

## Responsabilidades

1. Receber DTOs (Data Transfer Objects) do handler
2. Validar regras de negócio complexas
3. Converter DTOs → Models
4. Chamar repository para persistência
5. Orquestrar operações em múltiplos models
6. Retornar models ou DTOs ao handler

## DTOs (Data Transfer Objects)

### Por quê DTOs separados?

- **Separação de concerns:** API contract ≠ model interno
- **Validação específica:** Criar vs Atualizar têm regras diferentes
- **Flexibilidade:** API pode mudar sem mudar DB schema
- **Segurança:** Não expor campos internos

### Create DTO (Todos campos obrigatórios relevantes)

```go
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
```

### Update DTO (Todos campos opcionais - partial update)

```go
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

## Service Pattern

### Estrutura

```go
package services

import (
    "errors"
    "github.com/google/uuid"
    "github.com/plenya/api/internal/models"
    "github.com/plenya/api/internal/repository"
)

type ScoreService struct {
    repo *repository.ScoreRepository
}

func NewScoreService(repo *repository.ScoreRepository) *ScoreService {
    return &ScoreService{repo: repo}
}
```

### CreateItem Pattern

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

### UpdateItem Pattern (Partial Update)

```go
func (s *ScoreService) UpdateItem(id uuid.UUID, dto UpdateScoreItemDTO) (*models.ScoreItem, error) {
    // 1. Buscar item existente
    item, err := s.repo.GetScoreItemByID(id)
    if err != nil {
        return nil, err
    }

    // 2. Atualizar apenas campos fornecidos
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
    if dto.PatientExplanation != nil {
        item.PatientExplanation = dto.PatientExplanation
    }
    if dto.Conduct != nil {
        item.Conduct = dto.Conduct
    }
    if dto.Gender != nil {
        item.Gender = dto.Gender
    }
    if dto.AgeRangeMin != nil {
        item.AgeRangeMin = dto.AgeRangeMin
    }
    if dto.AgeRangeMax != nil {
        item.AgeRangeMax = dto.AgeRangeMax
    }

    // 3. ParentItemID especial (pode ser nil para "unindent")
    if dto.Order != nil {
        item.ParentItemID = dto.ParentItemID
    }

    // 4. Salvar (BeforeUpdate hook executa aqui - atualiza LastReview)
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

### GetByID Pattern

```go
func (s *ScoreService) GetItemByID(id uuid.UUID) (*models.ScoreItem, error) {
    item, err := s.repo.GetScoreItemByID(id)
    if err != nil {
        return nil, errors.New("score item not found")
    }
    return item, nil
}
```

### Delete Pattern

```go
func (s *ScoreService) DeleteItem(id uuid.UUID) error {
    // 1. Verificar se existe
    _, err := s.repo.GetScoreItemByID(id)
    if err != nil {
        return errors.New("score item not found")
    }

    // 2. Deletar (soft delete - seta DeletedAt)
    if err := s.repo.DeleteScoreItem(id); err != nil {
        return err
    }

    return nil
}
```

### List Pattern

```go
func (s *ScoreService) ListItems(filters ListFilters) ([]models.ScoreItem, error) {
    items, err := s.repo.ListScoreItems(filters)
    if err != nil {
        return nil, err
    }
    return items, nil
}
```

## Validação de Regras de Negócio

### Exemplo: Verificar Duplicatas

```go
func (s *PatientService) Create(dto CreatePatientDTO) (*models.Patient, error) {
    // Validar CPF único
    if dto.CPF != nil {
        exists, err := s.repo.ExistsByCPF(*dto.CPF)
        if err != nil {
            return nil, err
        }
        if exists {
            return nil, errors.New("CPF already registered")
        }
    }

    patient := &models.Patient{
        Name:      dto.Name,
        BirthDate: dto.BirthDate,
        CPF:       dto.CPF,
    }

    if err := s.repo.Create(patient); err != nil {
        return nil, err
    }

    return patient, nil
}
```

### Exemplo: Validar Hierarquia

```go
func (s *ScoreService) UpdateItem(id uuid.UUID, dto UpdateScoreItemDTO) (*models.ScoreItem, error) {
    item, err := s.repo.GetScoreItemByID(id)
    if err != nil {
        return nil, err
    }

    // Validar ParentItemID (não pode ser o próprio item)
    if dto.ParentItemID != nil && *dto.ParentItemID == id {
        return nil, errors.New("item cannot be its own parent")
    }

    // Validar SubgroupID (se mudou, verificar se existe)
    if dto.SubgroupID != nil && *dto.SubgroupID != item.SubgroupID {
        _, err := s.repo.GetScoreSubgroupByID(*dto.SubgroupID)
        if err != nil {
            return nil, errors.New("subgroup not found")
        }
    }

    // ... resto da atualização
}
```

## Operações Complexas (Transações)

```go
func (s *ScoreService) DuplicateItem(itemID uuid.UUID, newName string) (*models.ScoreItem, error) {
    // 1. Buscar item original com níveis
    original, err := s.repo.GetScoreItemByID(itemID)
    if err != nil {
        return nil, err
    }

    // 2. Transaction
    var newItem *models.ScoreItem
    err = s.repo.Transaction(func(tx *gorm.DB) error {
        // 2.1. Criar novo item (cópia)
        newItem = &models.ScoreItem{
            Name:               newName,
            LabTestCode:        original.LabTestCode,
            Unit:               original.Unit,
            UnitConversion:     original.UnitConversion,
            Gender:             original.Gender,
            AgeRangeMin:        original.AgeRangeMin,
            AgeRangeMax:        original.AgeRangeMax,
            PostMenopause:      original.PostMenopause,
            ClinicalRelevance:  original.ClinicalRelevance,
            PatientExplanation: original.PatientExplanation,
            Conduct:            original.Conduct,
            Points:             original.Points,
            SubgroupID:         original.SubgroupID,
            Order:              original.Order + 1,
        }

        if err := tx.Create(newItem).Error; err != nil {
            return err
        }

        // 2.2. Duplicar níveis
        for _, level := range original.Levels {
            newLevel := &models.ScoreLevel{
                Level:              level.Level,
                Name:               level.Name,
                LowerLimit:         level.LowerLimit,
                UpperLimit:         level.UpperLimit,
                Operator:           level.Operator,
                ClinicalRelevance:  level.ClinicalRelevance,
                PatientExplanation: level.PatientExplanation,
                Conduct:            level.Conduct,
                ItemID:             newItem.ID,
            }
            if err := tx.Create(newLevel).Error; err != nil {
                return err
            }
        }

        return nil
    })

    if err != nil {
        return nil, err
    }

    // 3. Recarregar com relações
    return s.repo.GetScoreItemByID(newItem.ID)
}
```

## Helper Functions

```go
// Ponteiros para tipos primitivos
func stringPtr(s string) *string {
    return &s
}

func intPtr(i int) *int {
    return &i
}

func boolPtr(b bool) *bool {
    return &b
}

func float64Ptr(f float64) *float64 {
    return &f
}

// Uso
item := &models.ScoreItem{
    Name:  "Hemoglobina",
    Unit:  stringPtr("g/dL"),
    Order: intPtr(1),
}
```

## Regras

- ✅ Validar foreign keys antes de criar
- ✅ Auto-incrementar order se não fornecido
- ✅ Sempre recarregar após salvar (Preload)
- ✅ Update: apenas campos fornecidos (partial)
- ✅ Retornar erros descritivos
- ❌ NÃO fazer queries diretas (usar repository)
- ❌ NÃO retornar erros de DB diretamente
- ❌ NÃO colocar lógica de apresentação

Ver também: `apps/api/internal/services/score_service.go`
