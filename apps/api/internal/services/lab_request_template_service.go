package services

import (
	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
)

// LabRequestTemplateService handles business logic for lab request templates
type LabRequestTemplateService struct {
	repo *repository.LabRequestTemplateRepository
}

// NewLabRequestTemplateService creates a new lab request template service
func NewLabRequestTemplateService(repo *repository.LabRequestTemplateRepository) *LabRequestTemplateService {
	return &LabRequestTemplateService{repo: repo}
}

// CreateLabRequestTemplate creates a new lab request template
func (s *LabRequestTemplateService) CreateLabRequestTemplate(template *models.LabRequestTemplate) error {
	return s.repo.CreateLabRequestTemplate(template)
}

// GetLabRequestTemplateByID retrieves a lab request template by ID with all associated lab tests
func (s *LabRequestTemplateService) GetLabRequestTemplateByID(id uuid.UUID) (*models.LabRequestTemplate, error) {
	return s.repo.GetLabRequestTemplateByID(id)
}

// GetAllLabRequestTemplates retrieves all active lab request templates
func (s *LabRequestTemplateService) GetAllLabRequestTemplates() ([]models.LabRequestTemplate, error) {
	return s.repo.GetAllLabRequestTemplates()
}

// GetAllLabRequestTemplatesWithTests retrieves all active templates with their lab tests
func (s *LabRequestTemplateService) GetAllLabRequestTemplatesWithTests() ([]models.LabRequestTemplate, error) {
	return s.repo.GetAllLabRequestTemplatesWithTests()
}

// UpdateLabRequestTemplate updates an existing lab request template
func (s *LabRequestTemplateService) UpdateLabRequestTemplate(id uuid.UUID, template *models.LabRequestTemplate) error {
	existing, err := s.repo.GetLabRequestTemplateByID(id)
	if err != nil {
		return err
	}

	// Update fields
	existing.Name = template.Name
	existing.Description = template.Description
	existing.DisplayOrder = template.DisplayOrder
	existing.IsActive = template.IsActive

	return s.repo.UpdateLabRequestTemplate(existing)
}

// UpdateLabRequestTemplateTests updates the lab tests associated with a template
func (s *LabRequestTemplateService) UpdateLabRequestTemplateTests(templateID uuid.UUID, testIDs []uuid.UUID) error {
	return s.repo.UpdateLabRequestTemplateTests(templateID, testIDs)
}

// AddLabTestToTemplate adds a single lab test to a template
func (s *LabRequestTemplateService) AddLabTestToTemplate(templateID, testID uuid.UUID) error {
	return s.repo.AddLabTestToTemplate(templateID, testID)
}

// RemoveLabTestFromTemplate removes a single lab test from a template
func (s *LabRequestTemplateService) RemoveLabTestFromTemplate(templateID, testID uuid.UUID) error {
	return s.repo.RemoveLabTestFromTemplate(templateID, testID)
}

// DeleteLabRequestTemplate soft deletes a lab request template
func (s *LabRequestTemplateService) DeleteLabRequestTemplate(id uuid.UUID) error {
	return s.repo.DeleteLabRequestTemplate(id)
}

// SearchLabRequestTemplates searches templates by name
func (s *LabRequestTemplateService) SearchLabRequestTemplates(searchTerm string) ([]models.LabRequestTemplate, error) {
	return s.repo.SearchLabRequestTemplates(searchTerm)
}
