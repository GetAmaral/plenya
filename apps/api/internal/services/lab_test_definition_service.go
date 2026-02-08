package services

import (
	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
)

// LabTestDefinitionService handles business logic for lab test definitions
type LabTestDefinitionService struct {
	repo *repository.LabTestDefinitionRepository
}

// NewLabTestDefinitionService creates a new lab test definition service
func NewLabTestDefinitionService(repo *repository.LabTestDefinitionRepository) *LabTestDefinitionService {
	return &LabTestDefinitionService{repo: repo}
}

// ============================================================
// LabTestDefinition Operations
// ============================================================

// CreateLabTestDefinition creates a new lab test definition
func (s *LabTestDefinitionService) CreateLabTestDefinition(def *models.LabTestDefinition) error {
	return s.repo.CreateLabTestDefinition(def)
}

// GetLabTestDefinitionByID retrieves a lab test definition by ID
func (s *LabTestDefinitionService) GetLabTestDefinitionByID(id uuid.UUID) (*models.LabTestDefinition, error) {
	return s.repo.GetLabTestDefinitionByID(id)
}

// GetLabTestDefinitionByCode retrieves a lab test definition by code
func (s *LabTestDefinitionService) GetLabTestDefinitionByCode(code string) (*models.LabTestDefinition, error) {
	return s.repo.GetLabTestDefinitionByCode(code)
}

// GetAllLabTestDefinitions retrieves all lab test definitions
func (s *LabTestDefinitionService) GetAllLabTestDefinitions() ([]models.LabTestDefinition, error) {
	return s.repo.GetAllLabTestDefinitions()
}

// GetRequestableLabTests retrieves only requestable lab tests
func (s *LabTestDefinitionService) GetRequestableLabTests(category *models.LabTestCategory) ([]models.LabTestDefinition, error) {
	return s.repo.GetRequestableLabTests(category)
}

// GetSubTests retrieves all sub-tests for a parent test
func (s *LabTestDefinitionService) GetSubTests(parentID uuid.UUID) ([]models.LabTestDefinition, error) {
	return s.repo.GetSubTests(parentID)
}

// SearchLabTestDefinitions searches lab tests by name or code
func (s *LabTestDefinitionService) SearchLabTestDefinitions(searchTerm string) ([]models.LabTestDefinition, error) {
	return s.repo.SearchLabTestDefinitions(searchTerm)
}

// UpdateLabTestDefinition updates an existing lab test definition
func (s *LabTestDefinitionService) UpdateLabTestDefinition(id uuid.UUID, def *models.LabTestDefinition) error {
	existing, err := s.repo.GetLabTestDefinitionByID(id)
	if err != nil {
		return err
	}

	// Update fields
	existing.Code = def.Code
	existing.Name = def.Name
	existing.ShortName = def.ShortName
	existing.TussCode = def.TussCode
	existing.LoincCode = def.LoincCode
	existing.Category = def.Category
	existing.IsRequestable = def.IsRequestable
	existing.ParentTestID = def.ParentTestID
	existing.Unit = def.Unit
	existing.UnitConversion = def.UnitConversion
	existing.ResultType = def.ResultType
	existing.CollectionMethod = def.CollectionMethod
	existing.FastingHours = def.FastingHours
	existing.SpecimenType = def.SpecimenType
	existing.Description = def.Description
	existing.ClinicalIndications = def.ClinicalIndications
	existing.ClinicalSignificance = def.ClinicalSignificance
	existing.LongevityContext = def.LongevityContext
	existing.ClinicalRecommendations = def.ClinicalRecommendations
	existing.DisplayOrder = def.DisplayOrder
	existing.IsActive = def.IsActive

	return s.repo.UpdateLabTestDefinition(existing)
}

// DeleteLabTestDefinition soft deletes a lab test definition
func (s *LabTestDefinitionService) DeleteLabTestDefinition(id uuid.UUID) error {
	return s.repo.DeleteLabTestDefinition(id)
}

