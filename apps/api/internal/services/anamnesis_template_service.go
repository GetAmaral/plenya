package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
	"gorm.io/gorm"
)

var (
	ErrAnamnesisTemplateNotFound = errors.New("anamnesis template not found")
)

type AnamnesisTemplateService struct {
	repo *repository.AnamnesisTemplateRepository
}

func NewAnamnesisTemplateService(repo *repository.AnamnesisTemplateRepository) *AnamnesisTemplateService {
	return &AnamnesisTemplateService{repo: repo}
}

// Create creates a new anamnesis template
func (s *AnamnesisTemplateService) Create(name string, area string) (*models.AnamnesisTemplate, error) {
	template := &models.AnamnesisTemplate{
		Name: name,
		Area: area,
	}

	if err := s.repo.Create(template); err != nil {
		return nil, err
	}

	return template, nil
}

// GetByID retrieves an anamnesis template by ID
func (s *AnamnesisTemplateService) GetByID(id uuid.UUID, withItems bool) (*models.AnamnesisTemplate, error) {
	template, err := s.repo.GetByID(id, withItems)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAnamnesisTemplateNotFound
		}
		return nil, err
	}

	return template, nil
}

// GetAll retrieves all anamnesis templates
func (s *AnamnesisTemplateService) GetAll(withItems bool) ([]models.AnamnesisTemplate, error) {
	return s.repo.GetAll(withItems)
}

// Update updates an anamnesis template
func (s *AnamnesisTemplateService) Update(id uuid.UUID, name *string, area *string) (*models.AnamnesisTemplate, error) {
	template, err := s.repo.GetByID(id, false)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAnamnesisTemplateNotFound
		}
		return nil, err
	}

	if name != nil {
		template.Name = *name
	}
	if area != nil {
		template.Area = *area
	}

	if err := s.repo.Update(template); err != nil {
		return nil, err
	}

	return template, nil
}

// Delete deletes an anamnesis template
func (s *AnamnesisTemplateService) Delete(id uuid.UUID) error {
	// Check if template exists
	if _, err := s.repo.GetByID(id, false); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAnamnesisTemplateNotFound
		}
		return err
	}

	return s.repo.Delete(id)
}

// UpdateItems updates all items of a template
func (s *AnamnesisTemplateService) UpdateItems(templateID uuid.UUID, itemsData []struct {
	ScoreItemID string
	Order       int
}) error {
	// Check if template exists
	if _, err := s.repo.GetByID(templateID, false); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAnamnesisTemplateNotFound
		}
		return err
	}

	// Parse and create items
	items := make([]models.AnamnesisTemplateItem, len(itemsData))
	for i, data := range itemsData {
		scoreItemID, err := uuid.Parse(data.ScoreItemID)
		if err != nil {
			return errors.New("invalid score item id")
		}

		items[i] = models.AnamnesisTemplateItem{
			AnamnesisTemplateID: templateID,
			ScoreItemID:         scoreItemID,
			Order:               data.Order,
		}
	}

	return s.repo.UpdateItems(templateID, items)
}

// Search searches anamnesis templates
func (s *AnamnesisTemplateService) Search(query string) ([]models.AnamnesisTemplate, error) {
	return s.repo.Search(query)
}
