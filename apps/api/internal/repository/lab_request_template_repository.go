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
// ordenados pelo display_order DO TEMPLATE, com page_break_before anexado em cada exame.
func (r *LabRequestTemplateRepository) GetLabRequestTemplateByID(id uuid.UUID) (*models.LabRequestTemplate, error) {
	var template models.LabRequestTemplate
	err := r.db.
		Preload("LabTests", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_active = ? AND is_requestable = ?", true, true).
				Order("lab_request_template_tests.display_order ASC, lab_test_definitions.name ASC")
		}).
		First(&template, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("lab request template not found")
		}
		return nil, err
	}
	if err := r.attachLayout(&template); err != nil {
		return nil, err
	}
	return &template, nil
}

// attachLayout preenche o campo transiente PageBreakBefore de cada LabTests a partir do join.
func (r *LabRequestTemplateRepository) attachLayout(templates ...*models.LabRequestTemplate) error {
	for _, tpl := range templates {
		if len(tpl.LabTests) == 0 {
			continue
		}
		var links []models.LabRequestTemplateTest
		if err := r.db.Where("lab_request_template_id = ?", tpl.ID).Find(&links).Error; err != nil {
			return err
		}
		brk := make(map[uuid.UUID]bool, len(links))
		for _, l := range links {
			brk[l.LabTestDefinitionID] = l.PageBreakBefore
		}
		for i := range tpl.LabTests {
			tpl.LabTests[i].PageBreakBefore = brk[tpl.LabTests[i].ID]
		}
	}
	return nil
}

// GetAllLabRequestTemplates retrieves all active lab request templates
func (r *LabRequestTemplateRepository) GetAllLabRequestTemplates() ([]models.LabRequestTemplate, error) {
	var templates []models.LabRequestTemplate
	err := r.db.
		Where("is_active = ?", true).
		Order("display_order ASC, name ASC").
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
				Order("lab_request_template_tests.display_order ASC, lab_test_definitions.name ASC")
		}).
		Order("display_order ASC, name ASC").
		Find(&templates).Error
	if err != nil {
		return nil, err
	}
	ptrs := make([]*models.LabRequestTemplate, len(templates))
	for i := range templates {
		ptrs[i] = &templates[i]
	}
	if err := r.attachLayout(ptrs...); err != nil {
		return nil, err
	}
	return templates, nil
}

// UpdateLabRequestTemplate updates an existing lab request template
func (r *LabRequestTemplateRepository) UpdateLabRequestTemplate(template *models.LabRequestTemplate) error {
	return r.db.Save(template).Error
}

// UpdateLabRequestTemplateTests updates the lab tests associated with a template
func (r *LabRequestTemplateRepository) UpdateLabRequestTemplateTests(templateID uuid.UUID, testIDs []uuid.UUID) error {
	// Use a simpler approach to avoid deadlocks
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Delete all existing associations for this template
		if err := tx.Exec("DELETE FROM lab_request_template_tests WHERE lab_request_template_id = ?", templateID).Error; err != nil {
			return err
		}

		// Insert new associations if any
		if len(testIDs) > 0 {
			// Build values for bulk insert
			for _, testID := range testIDs {
				if err := tx.Exec(
					"INSERT INTO lab_request_template_tests (lab_request_template_id, lab_test_definition_id) VALUES (?, ?) ON CONFLICT DO NOTHING",
					templateID, testID,
				).Error; err != nil {
					return err
				}
			}
		}

		// Update template's updated_at timestamp
		if err := tx.Exec("UPDATE lab_request_templates SET updated_at = NOW() WHERE id = ?", templateID).Error; err != nil {
			return err
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
