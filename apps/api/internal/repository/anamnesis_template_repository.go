package repository

import (
	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"gorm.io/gorm"
)

type AnamnesisTemplateRepository struct {
	db *gorm.DB
}

func NewAnamnesisTemplateRepository(db *gorm.DB) *AnamnesisTemplateRepository {
	return &AnamnesisTemplateRepository{db: db}
}

// Create creates a new anamnesis template
func (r *AnamnesisTemplateRepository) Create(template *models.AnamnesisTemplate) error {
	return r.db.Create(template).Error
}

// GetByID retrieves an anamnesis template by ID with optional items preload
func (r *AnamnesisTemplateRepository) GetByID(id uuid.UUID, withItems bool) (*models.AnamnesisTemplate, error) {
	var template models.AnamnesisTemplate
	query := r.db

	if withItems {
		query = query.Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC").
				Preload("ScoreItem.Subgroup.Group").
				Preload("ScoreItem.Levels", func(db *gorm.DB) *gorm.DB {
					return db.Order("level ASC")
				})
		})
	}

	if err := query.First(&template, id).Error; err != nil {
		return nil, err
	}

	return &template, nil
}

// GetAll retrieves all anamnesis templates with optional items preload
func (r *AnamnesisTemplateRepository) GetAll(withItems bool) ([]models.AnamnesisTemplate, error) {
	var templates []models.AnamnesisTemplate
	query := r.db.Order("name ASC")

	if withItems {
		query = query.Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC").
				Preload("ScoreItem.Subgroup.Group").
				Preload("ScoreItem.Levels", func(db *gorm.DB) *gorm.DB {
					return db.Order("level ASC")
				})
		})
	}

	if err := query.Find(&templates).Error; err != nil {
		return nil, err
	}

	return templates, nil
}

// Update updates an anamnesis template
func (r *AnamnesisTemplateRepository) Update(template *models.AnamnesisTemplate) error {
	return r.db.Save(template).Error
}

// Delete soft deletes an anamnesis template
func (r *AnamnesisTemplateRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.AnamnesisTemplate{}, id).Error
}

// UpdateItems updates all items of a template (replace all)
func (r *AnamnesisTemplateRepository) UpdateItems(templateID uuid.UUID, items []models.AnamnesisTemplateItem) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Delete all existing items
		if err := tx.Where("anamnesis_template_id = ?", templateID).Delete(&models.AnamnesisTemplateItem{}).Error; err != nil {
			return err
		}

		// Create new items
		if len(items) > 0 {
			for i := range items {
				items[i].AnamnesisTemplateID = templateID
			}
			if err := tx.Create(&items).Error; err != nil {
				return err
			}
		}

		// Update template's updated_at timestamp
		return tx.Model(&models.AnamnesisTemplate{}).Where("id = ?", templateID).Update("updated_at", gorm.Expr("CURRENT_TIMESTAMP")).Error
	})
}

// Search searches anamnesis templates by name or area
func (r *AnamnesisTemplateRepository) Search(query string) ([]models.AnamnesisTemplate, error) {
	var templates []models.AnamnesisTemplate
	searchPattern := "%" + query + "%"

	if err := r.db.Where("name ILIKE ? OR area ILIKE ?", searchPattern, searchPattern).
		Order("name ASC").
		Find(&templates).Error; err != nil {
		return nil, err
	}

	return templates, nil
}
