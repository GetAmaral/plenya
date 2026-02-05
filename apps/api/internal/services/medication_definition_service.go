package services

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

var (
	ErrMedicationDefinitionNotFound = errors.New("medication definition not found")
)

type MedicationDefinitionService struct {
	db *gorm.DB
}

func NewMedicationDefinitionService(db *gorm.DB) *MedicationDefinitionService {
	return &MedicationDefinitionService{db: db}
}

// List lista todas as definições de medicamentos com filtros
func (s *MedicationDefinitionService) List(category *models.MedicationCategory, limit, offset int) ([]models.MedicationDefinition, int64, error) {
	var medications []models.MedicationDefinition
	var total int64

	query := s.db.Model(&models.MedicationDefinition{})

	// Filter by category if provided
	if category != nil && *category != "" {
		query = query.Where("category = ?", *category)
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	if err := query.Order("common_name ASC").
		Limit(limit).
		Offset(offset).
		Find(&medications).Error; err != nil {
		return nil, 0, err
	}

	return medications, total, nil
}

// Search busca medicamentos por nome ou princípio ativo
func (s *MedicationDefinitionService) Search(query string, limit int) ([]models.MedicationDefinition, error) {
	var medications []models.MedicationDefinition

	searchPattern := "%" + query + "%"

	if err := s.db.Where("LOWER(common_name) LIKE LOWER(?) OR LOWER(active_ingredient) LIKE LOWER(?)", searchPattern, searchPattern).
		Limit(limit).
		Order("common_name ASC").
		Find(&medications).Error; err != nil {
		return nil, err
	}

	return medications, nil
}

// GetByID busca uma definição por ID
func (s *MedicationDefinitionService) GetByID(id uuid.UUID) (*models.MedicationDefinition, error) {
	var medication models.MedicationDefinition
	if err := s.db.First(&medication, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMedicationDefinitionNotFound
		}
		return nil, err
	}

	return &medication, nil
}

// Create cria uma nova definição de medicamento
func (s *MedicationDefinitionService) Create(medication *models.MedicationDefinition) error {
	return s.db.Create(medication).Error
}

// Update atualiza uma definição de medicamento
func (s *MedicationDefinitionService) Update(id uuid.UUID, medication *models.MedicationDefinition) error {
	result := s.db.Model(&models.MedicationDefinition{}).
		Where("id = ?", id).
		Updates(medication)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrMedicationDefinitionNotFound
	}

	return nil
}

// Delete deleta uma definição de medicamento (soft delete)
func (s *MedicationDefinitionService) Delete(id uuid.UUID) error {
	result := s.db.Delete(&models.MedicationDefinition{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrMedicationDefinitionNotFound
	}

	return nil
}
