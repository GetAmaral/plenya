package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
)

// ScoreService handles business logic for score operations
type ScoreService struct {
	repo       *repository.ScoreRepository
	pdfService *ScorePDFService
}

// NewScoreService creates a new score service instance
func NewScoreService(repo *repository.ScoreRepository) *ScoreService {
	return &ScoreService{
		repo:       repo,
		pdfService: NewScorePDFService(),
	}
}

// ============================================================
// DTOs (Data Transfer Objects)
// ============================================================

// CreateScoreGroupDTO represents the request to create a score group
type CreateScoreGroupDTO struct {
	Name  string `json:"name" validate:"required,min=2,max=200"`
	Order *int   `json:"order,omitempty"`
}

// UpdateScoreGroupDTO represents the request to update a score group
type UpdateScoreGroupDTO struct {
	Name  *string `json:"name,omitempty" validate:"omitempty,min=2,max=200"`
	Order *int    `json:"order,omitempty" validate:"omitempty,gte=0,lte=9999"`
}

// CreateScoreSubgroupDTO represents the request to create a score subgroup
type CreateScoreSubgroupDTO struct {
	Name    string    `json:"name" validate:"required,min=2,max=200"`
	GroupID uuid.UUID `json:"groupId" validate:"required"`
	Order   *int      `json:"order,omitempty"`
}

// UpdateScoreSubgroupDTO represents the request to update a score subgroup
type UpdateScoreSubgroupDTO struct {
	Name  *string `json:"name,omitempty" validate:"omitempty,min=2,max=200"`
	Order *int    `json:"order,omitempty" validate:"omitempty,gte=0,lte=9999"`
}

// CreateScoreItemDTO represents the request to create a score item
type CreateScoreItemDTO struct {
	Name           string     `json:"name" validate:"required,min=2,max=300"`
	Unit           *string    `json:"unit,omitempty" validate:"omitempty,max=50"`
	UnitConversion *string    `json:"unitConversion,omitempty"`
	Points         *float64   `json:"points,omitempty" validate:"omitempty,gte=0,lte=100"`
	SubgroupID     uuid.UUID  `json:"subgroupId" validate:"required"`
	ParentItemID   *uuid.UUID `json:"parentItemId,omitempty"`
	Order          *int       `json:"order,omitempty"`
}

// UpdateScoreItemDTO represents the request to update a score item
type UpdateScoreItemDTO struct {
	Name               *string     `json:"name,omitempty" validate:"omitempty,min=2,max=300"`
	Unit               *string     `json:"unit,omitempty" validate:"omitempty,max=50"`
	UnitConversion     *string     `json:"unitConversion,omitempty"`
	Points             *float64    `json:"points,omitempty" validate:"omitempty,gte=0,lte=100"`
	Order              *int        `json:"order,omitempty" validate:"omitempty,gte=0,lte=9999"`
	SubgroupID         *uuid.UUID  `json:"subgroupId,omitempty"`
	ParentItemID       *uuid.UUID  `json:"parentItemId,omitempty"`
	ClinicalRelevance  *string     `json:"clinicalRelevance,omitempty"`
	PatientExplanation *string     `json:"patientExplanation,omitempty"`
	Conduct            *string     `json:"conduct,omitempty"`
	LastReview         *time.Time  `json:"lastReview,omitempty"`
}

// CreateScoreLevelDTO represents the request to create a score level
type CreateScoreLevelDTO struct {
	Level      int        `json:"level" validate:"required,gte=0,lte=6"`
	Name       string     `json:"name" validate:"required,min=1,max=500"`
	LowerLimit *string    `json:"lowerLimit,omitempty" validate:"omitempty,max=50"`
	UpperLimit *string    `json:"upperLimit,omitempty" validate:"omitempty,max=50"`
	Operator   string     `json:"operator" validate:"required,oneof== > >= < <= between"`
	ItemID     uuid.UUID  `json:"itemId" validate:"required"`
}

// UpdateScoreLevelDTO represents the request to update a score level
type UpdateScoreLevelDTO struct {
	Level      *int    `json:"level,omitempty" validate:"omitempty,gte=0,lte=6"`
	Name       *string `json:"name,omitempty" validate:"omitempty,min=1,max=500"`
	LowerLimit *string `json:"lowerLimit,omitempty" validate:"omitempty,max=50"`
	UpperLimit *string `json:"upperLimit,omitempty" validate:"omitempty,max=50"`
	Operator   *string `json:"operator,omitempty" validate:"omitempty,oneof== > >= < <= between"`
}

// ============================================================
// ScoreGroup Operations
// ============================================================

// CreateGroup creates a new score group with auto-incremented order
func (s *ScoreService) CreateGroup(dto CreateScoreGroupDTO) (*models.ScoreGroup, error) {
	// Auto-increment order if not provided
	order := 0
	if dto.Order != nil {
		order = *dto.Order
	} else {
		maxOrder, err := s.repo.GetMaxOrderForGroup()
		if err != nil {
			return nil, err
		}
		order = maxOrder + 1
	}

	group := &models.ScoreGroup{
		Name:  dto.Name,
		Order: order,
	}

	if err := s.repo.CreateScoreGroup(group); err != nil {
		return nil, err
	}

	return group, nil
}

// GetGroupByID retrieves a score group by ID
func (s *ScoreService) GetGroupByID(id uuid.UUID) (*models.ScoreGroup, error) {
	return s.repo.GetScoreGroupByID(id)
}

// GetAllGroups retrieves all score groups
func (s *ScoreService) GetAllGroups() ([]models.ScoreGroup, error) {
	return s.repo.GetAllScoreGroups()
}

// GetGroupTree retrieves a score group with all nested data
func (s *ScoreService) GetGroupTree(id uuid.UUID) (*models.ScoreGroup, error) {
	return s.repo.GetScoreGroupTree(id)
}

// GetAllGroupTrees retrieves all score groups with nested data
func (s *ScoreService) GetAllGroupTrees() ([]models.ScoreGroup, error) {
	return s.repo.GetAllScoreGroupTrees()
}

// UpdateGroup updates an existing score group
func (s *ScoreService) UpdateGroup(id uuid.UUID, dto UpdateScoreGroupDTO) (*models.ScoreGroup, error) {
	group, err := s.repo.GetScoreGroupByID(id)
	if err != nil {
		return nil, err
	}

	if dto.Name != nil {
		group.Name = *dto.Name
	}
	if dto.Order != nil {
		group.Order = *dto.Order
	}

	if err := s.repo.UpdateScoreGroup(group); err != nil {
		return nil, err
	}

	return group, nil
}

// DeleteGroup soft deletes a score group
func (s *ScoreService) DeleteGroup(id uuid.UUID) error {
	return s.repo.DeleteScoreGroup(id)
}

// ============================================================
// ScoreSubgroup Operations
// ============================================================

// CreateSubgroup creates a new score subgroup with auto-incremented order
func (s *ScoreService) CreateSubgroup(dto CreateScoreSubgroupDTO) (*models.ScoreSubgroup, error) {
	// Verify parent group exists
	_, err := s.repo.GetScoreGroupByID(dto.GroupID)
	if err != nil {
		return nil, errors.New("parent group not found")
	}

	// Auto-increment order if not provided
	order := 0
	if dto.Order != nil {
		order = *dto.Order
	} else {
		maxOrder, err := s.repo.GetMaxOrderForSubgroup(dto.GroupID)
		if err != nil {
			return nil, err
		}
		order = maxOrder + 1
	}

	subgroup := &models.ScoreSubgroup{
		Name:    dto.Name,
		GroupID: dto.GroupID,
		Order:   order,
	}

	if err := s.repo.CreateScoreSubgroup(subgroup); err != nil {
		return nil, err
	}

	// Load the group relation
	subgroup, err = s.repo.GetScoreSubgroupByID(subgroup.ID)
	if err != nil {
		return nil, err
	}

	return subgroup, nil
}

// GetSubgroupByID retrieves a score subgroup by ID
func (s *ScoreService) GetSubgroupByID(id uuid.UUID) (*models.ScoreSubgroup, error) {
	return s.repo.GetScoreSubgroupByID(id)
}

// GetSubgroupsByGroupID retrieves all subgroups for a group
func (s *ScoreService) GetSubgroupsByGroupID(groupID uuid.UUID) ([]models.ScoreSubgroup, error) {
	return s.repo.GetSubgroupsByGroupID(groupID)
}

// UpdateSubgroup updates an existing score subgroup
func (s *ScoreService) UpdateSubgroup(id uuid.UUID, dto UpdateScoreSubgroupDTO) (*models.ScoreSubgroup, error) {
	subgroup, err := s.repo.GetScoreSubgroupByID(id)
	if err != nil {
		return nil, err
	}

	if dto.Name != nil {
		subgroup.Name = *dto.Name
	}
	if dto.Order != nil {
		subgroup.Order = *dto.Order
	}

	if err := s.repo.UpdateScoreSubgroup(subgroup); err != nil {
		return nil, err
	}

	return subgroup, nil
}

// DeleteSubgroup soft deletes a score subgroup
func (s *ScoreService) DeleteSubgroup(id uuid.UUID) error {
	return s.repo.DeleteScoreSubgroup(id)
}

// ============================================================
// ScoreItem Operations
// ============================================================

// CreateItem creates a new score item with auto-incremented order
func (s *ScoreService) CreateItem(dto CreateScoreItemDTO) (*models.ScoreItem, error) {
	// Verify parent subgroup exists
	_, err := s.repo.GetScoreSubgroupByID(dto.SubgroupID)
	if err != nil {
		return nil, errors.New("parent subgroup not found")
	}

	// Verify parent item exists if provided
	if dto.ParentItemID != nil {
		_, err := s.repo.GetScoreItemByID(*dto.ParentItemID)
		if err != nil {
			return nil, errors.New("parent item not found")
		}
	}

	// Auto-increment order if not provided
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

	item := &models.ScoreItem{
		Name:           dto.Name,
		Unit:           dto.Unit,
		UnitConversion: dto.UnitConversion,
		Points:         dto.Points,
		SubgroupID:     dto.SubgroupID,
		ParentItemID:   dto.ParentItemID,
		Order:          order,
	}

	if err := s.repo.CreateScoreItem(item); err != nil {
		return nil, err
	}

	// Load relations
	item, err = s.repo.GetScoreItemByID(item.ID)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// GetItemByID retrieves a score item by ID
func (s *ScoreService) GetItemByID(id uuid.UUID) (*models.ScoreItem, error) {
	return s.repo.GetScoreItemByID(id)
}

// GetItemsBySubgroupID retrieves all items for a subgroup
func (s *ScoreService) GetItemsBySubgroupID(subgroupID uuid.UUID) ([]models.ScoreItem, error) {
	return s.repo.GetItemsBySubgroupID(subgroupID)
}

// GetAllScoreItems retrieves all score items
func (s *ScoreService) GetAllScoreItems() ([]models.ScoreItem, error) {
	return s.repo.GetAllScoreItems()
}

// UpdateItem updates an existing score item
func (s *ScoreService) UpdateItem(id uuid.UUID, dto UpdateScoreItemDTO) (*models.ScoreItem, error) {
	item, err := s.repo.GetScoreItemByID(id)
	if err != nil {
		return nil, err
	}

	// Update required fields
	if dto.Name != nil {
		item.Name = *dto.Name
	}
	if dto.Points != nil {
		item.Points = dto.Points
	}
	if dto.Order != nil {
		item.Order = *dto.Order
	}
	if dto.SubgroupID != nil {
		item.SubgroupID = *dto.SubgroupID
	}

	// Always update optional fields (allows clearing them with null)
	// Frontend sends null to clear, backend updates even if nil
	item.Unit = dto.Unit
	item.UnitConversion = dto.UnitConversion
	item.ParentItemID = dto.ParentItemID
	item.ClinicalRelevance = dto.ClinicalRelevance
	item.PatientExplanation = dto.PatientExplanation
	item.Conduct = dto.Conduct

	// LastReview is optional but not typically cleared
	if dto.LastReview != nil {
		item.LastReview = dto.LastReview
	}

	if err := s.repo.UpdateScoreItem(item); err != nil {
		return nil, err
	}

	return item, nil
}

// DeleteItem soft deletes a score item
func (s *ScoreService) DeleteItem(id uuid.UUID) error {
	return s.repo.DeleteScoreItem(id)
}

// ============================================================
// ScoreLevel Operations
// ============================================================

// CreateLevel creates a new score level
func (s *ScoreService) CreateLevel(dto CreateScoreLevelDTO) (*models.ScoreLevel, error) {
	// Verify parent item exists
	_, err := s.repo.GetScoreItemByID(dto.ItemID)
	if err != nil {
		return nil, errors.New("parent item not found")
	}

	// Validate operator and limits
	if err := s.validateLevelLimits(dto.Operator, dto.LowerLimit, dto.UpperLimit); err != nil {
		return nil, err
	}

	level := &models.ScoreLevel{
		Level:      dto.Level,
		Name:       dto.Name,
		LowerLimit: dto.LowerLimit,
		UpperLimit: dto.UpperLimit,
		Operator:   dto.Operator,
		ItemID:     dto.ItemID,
	}

	if err := s.repo.CreateScoreLevel(level); err != nil {
		return nil, err
	}

	// Load relations
	level, err = s.repo.GetScoreLevelByID(level.ID)
	if err != nil {
		return nil, err
	}

	return level, nil
}

// GetLevelByID retrieves a score level by ID
func (s *ScoreService) GetLevelByID(id uuid.UUID) (*models.ScoreLevel, error) {
	return s.repo.GetScoreLevelByID(id)
}

// GetLevelsByItemID retrieves all levels for an item
func (s *ScoreService) GetLevelsByItemID(itemID uuid.UUID) ([]models.ScoreLevel, error) {
	return s.repo.GetLevelsByItemID(itemID)
}

// UpdateLevel updates an existing score level
func (s *ScoreService) UpdateLevel(id uuid.UUID, dto UpdateScoreLevelDTO) (*models.ScoreLevel, error) {
	level, err := s.repo.GetScoreLevelByID(id)
	if err != nil {
		return nil, err
	}

	if dto.Level != nil {
		level.Level = *dto.Level
	}
	if dto.Name != nil {
		level.Name = *dto.Name
	}
	if dto.LowerLimit != nil {
		level.LowerLimit = dto.LowerLimit
	}
	if dto.UpperLimit != nil {
		level.UpperLimit = dto.UpperLimit
	}
	if dto.Operator != nil {
		level.Operator = *dto.Operator
	}

	// Validate operator and limits
	if err := s.validateLevelLimits(level.Operator, level.LowerLimit, level.UpperLimit); err != nil {
		return nil, err
	}

	if err := s.repo.UpdateScoreLevel(level); err != nil {
		return nil, err
	}

	return level, nil
}

// DeleteLevel soft deletes a score level
func (s *ScoreService) DeleteLevel(id uuid.UUID) error {
	return s.repo.DeleteScoreLevel(id)
}

// ============================================================
// Validation Helpers
// ============================================================

// validateLevelLimits validates operator and limit combinations
func (s *ScoreService) validateLevelLimits(operator string, lowerLimit, upperLimit *string) error {
	switch operator {
	case "between":
		if lowerLimit == nil || upperLimit == nil {
			return errors.New("both lowerLimit and upperLimit are required for 'between' operator")
		}
	case "=", ">", ">=", "<", "<=":
		// Single limit operators can have either lower or upper, or neither for simple comparisons
		// No strict validation needed
	default:
		return fmt.Errorf("invalid operator: %s", operator)
	}
	return nil
}

// ============================================================
// PDF Generation
// ============================================================

// GeneratePosterPDF generates a PDF poster (60cm x 300cm) for all scores
func (s *ScoreService) GeneratePosterPDF() ([]byte, error) {
	// Get all score groups with complete tree
	groups, err := s.repo.GetAllScoreGroupTrees()
	if err != nil {
		return nil, fmt.Errorf("failed to get score groups: %v", err)
	}

	// Generate PDF using dedicated PDF service
	return s.pdfService.GeneratePosterPDF(groups)
}
