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
	existing.DisplayOrder = def.DisplayOrder
	existing.IsActive = def.IsActive

	return s.repo.UpdateLabTestDefinition(existing)
}

// DeleteLabTestDefinition soft deletes a lab test definition
func (s *LabTestDefinitionService) DeleteLabTestDefinition(id uuid.UUID) error {
	return s.repo.DeleteLabTestDefinition(id)
}

// ============================================================
// LabTestScoreMapping Operations
// ============================================================

// CreateLabTestScoreMapping creates a new mapping
func (s *LabTestDefinitionService) CreateLabTestScoreMapping(mapping *models.LabTestScoreMapping) error {
	return s.repo.CreateLabTestScoreMapping(mapping)
}

// GetLabTestScoreMappingByID retrieves a mapping by ID
func (s *LabTestDefinitionService) GetLabTestScoreMappingByID(id uuid.UUID) (*models.LabTestScoreMapping, error) {
	return s.repo.GetLabTestScoreMappingByID(id)
}

// GetMappingsForLabTest retrieves all score mappings for a lab test
func (s *LabTestDefinitionService) GetMappingsForLabTest(labTestID uuid.UUID) ([]models.LabTestScoreMapping, error) {
	return s.repo.GetMappingsForLabTest(labTestID)
}

// GetMappingForPatient retrieves the appropriate mapping for a specific patient
func (s *LabTestDefinitionService) GetMappingForPatient(labTestID uuid.UUID, gender models.Gender, age int) (*models.LabTestScoreMapping, error) {
	return s.repo.GetMappingForPatient(labTestID, gender, age)
}

// UpdateLabTestScoreMapping updates a mapping
func (s *LabTestDefinitionService) UpdateLabTestScoreMapping(id uuid.UUID, mapping *models.LabTestScoreMapping) error {
	existing, err := s.repo.GetLabTestScoreMappingByID(id)
	if err != nil {
		return err
	}

	existing.LabTestID = mapping.LabTestID
	existing.ScoreItemID = mapping.ScoreItemID
	existing.Gender = mapping.Gender
	existing.MinAge = mapping.MinAge
	existing.MaxAge = mapping.MaxAge
	existing.Notes = mapping.Notes
	existing.IsActive = mapping.IsActive

	return s.repo.UpdateLabTestScoreMapping(existing)
}

// DeleteLabTestScoreMapping soft deletes a mapping
func (s *LabTestDefinitionService) DeleteLabTestScoreMapping(id uuid.UUID) error {
	return s.repo.DeleteLabTestScoreMapping(id)
}

// ============================================================
// LabTestReferenceRange Operations
// ============================================================

// CreateLabTestReferenceRange creates a new reference range
func (s *LabTestDefinitionService) CreateLabTestReferenceRange(refRange *models.LabTestReferenceRange) error {
	return s.repo.CreateLabTestReferenceRange(refRange)
}

// GetLabTestReferenceRangeByID retrieves a reference range by ID
func (s *LabTestDefinitionService) GetLabTestReferenceRangeByID(id uuid.UUID) (*models.LabTestReferenceRange, error) {
	return s.repo.GetLabTestReferenceRangeByID(id)
}

// GetReferenceRangesForLabTest retrieves all reference ranges for a lab test
func (s *LabTestDefinitionService) GetReferenceRangesForLabTest(labTestID uuid.UUID) ([]models.LabTestReferenceRange, error) {
	return s.repo.GetReferenceRangesForLabTest(labTestID)
}

// GetReferenceRangeForPatient retrieves the appropriate reference range for a specific patient
func (s *LabTestDefinitionService) GetReferenceRangeForPatient(labTestID uuid.UUID, gender models.Gender, age int) (*models.LabTestReferenceRange, error) {
	return s.repo.GetReferenceRangeForPatient(labTestID, gender, age)
}

// UpdateLabTestReferenceRange updates a reference range
func (s *LabTestDefinitionService) UpdateLabTestReferenceRange(id uuid.UUID, refRange *models.LabTestReferenceRange) error {
	existing, err := s.repo.GetLabTestReferenceRangeByID(id)
	if err != nil {
		return err
	}

	existing.LabTestID = refRange.LabTestID
	existing.Gender = refRange.Gender
	existing.MinAge = refRange.MinAge
	existing.MaxAge = refRange.MaxAge
	existing.LowerLimit = refRange.LowerLimit
	existing.UpperLimit = refRange.UpperLimit
	existing.TextReference = refRange.TextReference
	existing.Unit = refRange.Unit
	existing.Description = refRange.Description
	existing.Source = refRange.Source
	existing.IsActive = refRange.IsActive

	return s.repo.UpdateLabTestReferenceRange(existing)
}

// DeleteLabTestReferenceRange soft deletes a reference range
func (s *LabTestDefinitionService) DeleteLabTestReferenceRange(id uuid.UUID) error {
	return s.repo.DeleteLabTestReferenceRange(id)
}
