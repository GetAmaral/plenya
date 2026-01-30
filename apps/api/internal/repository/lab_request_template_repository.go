package repository

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// LabRequestTemplateRepository handles database operations for lab request templates
type LabRequestTemplateRepository struct {
	db *gorm.DB
}

// NewLabRequestTemplateRepository creates a new lab request template repository instance
func NewLabRequestTemplateRepository(db *gorm.DB) *LabRequestTemplateRepository {
	return &LabRequestTemplateRepository{db: db}
}

// CreateLabRequestTemplate creates a new lab request template
func (r *LabRequestTemplateRepository) CreateLabRequestTemplate(template *models.LabRequestTemplate) error {
	return r.db.Create(template).Error
}

// GetLabRequestTemplateByID retrieves a lab request template by ID with all associated lab tests
func (r *LabRequestTemplateRepository) GetLabRequestTemplateByID(id uuid.UUID) (*models.LabRequestTemplate, error) {
	var template models.LabRequestTemplate
	err := r.db.
		Preload("LabTests", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_active = ? AND is_requestable = ?", true, true).
				Order("name ASC")
		}).
		First(&template, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("lab request template not found")
		}
		return nil, err
	}
	return &template, nil
}

// GetAllLabRequestTemplates retrieves all active lab request templates
func (r *LabRequestTemplateRepository) GetAllLabRequestTemplates() ([]models.LabRequestTemplate, error) {
	var templates []models.LabRequestTemplate
	err := r.db.
		Where("is_active = ?", true).
		Order("name ASC").
		Find(&templates).Error
	return templates, err
}

// GetAllLabRequestTemplatesWithTests retrieves all active templates with their lab tests
func (r *LabRequestTemplateRepository) GetAllLabRequestTemplatesWithTests() ([]models.LabRequestTemplate, error) {
	var templates []models.LabRequestTemplate
	err := r.db.
		Where("is_active = ?", true).
		Preload("LabTests", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_active = ? AND is_requestable = ?", true, true).
				Order("name ASC")
		}).
		Order("name ASC").
		Find(&templates).Error
	return templates, err
}

// UpdateLabRequestTemplate updates an existing lab request template
func (r *LabRequestTemplateRepository) UpdateLabRequestTemplate(template *models.LabRequestTemplate) error {
	return r.db.Save(template).Error
}

// UpdateLabRequestTemplateTests updates the lab tests associated with a template
func (r *LabRequestTemplateRepository) UpdateLabRequestTemplateTests(templateID uuid.UUID, testIDs []uuid.UUID) error {
	// Start transaction
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Get template
		var template models.LabRequestTemplate
		if err := tx.First(&template, "id = ?", templateID).Error; err != nil {
			return err
		}

		// Clear existing associations
		if err := tx.Model(&template).Association("LabTests").Clear(); err != nil {
			return err
		}

		// Add new associations
		if len(testIDs) > 0 {
			var tests []models.LabTestDefinition
			if err := tx.Where("id IN ? AND is_active = ? AND is_requestable = ?", testIDs, true, true).
				Find(&tests).Error; err != nil {
				return err
			}

			if err := tx.Model(&template).Association("LabTests").Append(tests); err != nil {
				return err
			}
		}

		return nil
	})
}

// AddLabTestToTemplate adds a single lab test to a template
func (r *LabRequestTemplateRepository) AddLabTestToTemplate(templateID, testID uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Verify template exists
		var template models.LabRequestTemplate
		if err := tx.First(&template, "id = ?", templateID).Error; err != nil {
			return err
		}

		// Verify test exists and is requestable
		var test models.LabTestDefinition
		if err := tx.First(&test, "id = ? AND is_active = ? AND is_requestable = ?", testID, true, true).Error; err != nil {
			return err
		}

		// Add association
		return tx.Model(&template).Association("LabTests").Append(&test)
	})
}

// RemoveLabTestFromTemplate removes a single lab test from a template
func (r *LabRequestTemplateRepository) RemoveLabTestFromTemplate(templateID, testID uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Verify template exists
		var template models.LabRequestTemplate
		if err := tx.First(&template, "id = ?", templateID).Error; err != nil {
			return err
		}

		// Verify test exists
		var test models.LabTestDefinition
		if err := tx.First(&test, "id = ?", testID).Error; err != nil {
			return err
		}

		// Remove association
		return tx.Model(&template).Association("LabTests").Delete(&test)
	})
}

// DeleteLabRequestTemplate soft deletes a lab request template
func (r *LabRequestTemplateRepository) DeleteLabRequestTemplate(id uuid.UUID) error {
	result := r.db.Delete(&models.LabRequestTemplate{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("lab request template not found")
	}
	return nil
}

// SearchLabRequestTemplates searches templates by name
func (r *LabRequestTemplateRepository) SearchLabRequestTemplates(searchTerm string) ([]models.LabRequestTemplate, error) {
	var templates []models.LabRequestTemplate
	err := r.db.
		Where("name ILIKE ? AND is_active = ?", "%"+searchTerm+"%", true).
		Order("name ASC").
		Limit(50).
		Find(&templates).Error
	return templates, err
}
