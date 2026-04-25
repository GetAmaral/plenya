package services

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

type ContinuumBoxTemplateService struct {
	db *gorm.DB
}

func NewContinuumBoxTemplateService(db *gorm.DB) *ContinuumBoxTemplateService {
	return &ContinuumBoxTemplateService{db: db}
}

func (s *ContinuumBoxTemplateService) List(includeArchived bool) ([]models.ContinuumBoxTemplate, error) {
	var rows []models.ContinuumBoxTemplate
	q := s.db.Model(&models.ContinuumBoxTemplate{}).Order("created_at DESC")
	if !includeArchived {
		q = q.Where("status = ?", "active")
	}
	if err := q.Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (s *ContinuumBoxTemplateService) Get(id uuid.UUID) (*models.ContinuumBoxTemplate, error) {
	var t models.ContinuumBoxTemplate
	if err := s.db.First(&t, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *ContinuumBoxTemplateService) Create(payload dto.ContinuumBoxTemplatePayload, createdBy uuid.UUID) (*models.ContinuumBoxTemplate, error) {
	t := models.ContinuumBoxTemplate{
		Name:            payload.Name,
		Description:     payload.Description,
		Contents:        payload.Contents,
		Notes:           payload.Notes,
		Status:          firstNonEmpty(payload.Status, "active"),
		CreatedByUserID: createdBy,
	}
	if err := s.db.Create(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *ContinuumBoxTemplateService) Update(id uuid.UUID, payload dto.ContinuumBoxTemplatePayload) (*models.ContinuumBoxTemplate, error) {
	var t models.ContinuumBoxTemplate
	if err := s.db.First(&t, "id = ?", id).Error; err != nil {
		return nil, err
	}
	t.Name = payload.Name
	t.Description = payload.Description
	t.Contents = payload.Contents
	t.Notes = payload.Notes
	if payload.Status != "" {
		t.Status = payload.Status
	}
	if err := s.db.Save(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *ContinuumBoxTemplateService) Delete(id uuid.UUID) error {
	return s.db.Delete(&models.ContinuumBoxTemplate{}, "id = ?", id).Error
}

func (s *ContinuumBoxTemplateService) Clone(sourceID uuid.UUID, newName string, createdBy uuid.UUID) (*models.ContinuumBoxTemplate, error) {
	src, err := s.Get(sourceID)
	if err != nil {
		return nil, err
	}
	clone := models.ContinuumBoxTemplate{
		Name:            newName,
		Description:     src.Description,
		Contents:        src.Contents,
		Notes:           src.Notes,
		Status:          "active",
		CreatedByUserID: createdBy,
	}
	if err := s.db.Create(&clone).Error; err != nil {
		return nil, err
	}
	return &clone, nil
}
