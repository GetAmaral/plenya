package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"gorm.io/gorm"
)

var (
	ErrLabResultViewNotFound = errors.New("lab result view not found")
)

type LabResultViewRepository struct {
	db *gorm.DB
}

func NewLabResultViewRepository(db *gorm.DB) *LabResultViewRepository {
	return &LabResultViewRepository{db: db}
}

// Create cria uma nova view
func (r *LabResultViewRepository) Create(view *models.LabResultView) error {
	return r.db.Create(view).Error
}

// GetByID busca uma view por ID
func (r *LabResultViewRepository) GetByID(id uuid.UUID, withItems bool) (*models.LabResultView, error) {
	var view models.LabResultView
	query := r.db

	if withItems {
		query = query.Preload("Items.LabTestDefinition")
	}

	err := query.Where("id = ?", id).First(&view).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLabResultViewNotFound
		}
		return nil, err
	}

	return &view, nil
}

// GetAll busca todas as views
func (r *LabResultViewRepository) GetAll(withItems bool, includeInactive bool) ([]models.LabResultView, error) {
	var views []models.LabResultView
	query := r.db

	if withItems {
		query = query.Preload("Items.LabTestDefinition")
	}

	if !includeInactive {
		query = query.Where("is_active = ?", true)
	}

	err := query.Order("display_order ASC, name ASC").Find(&views).Error
	if err != nil {
		return nil, err
	}

	return views, nil
}

// Update atualiza uma view
func (r *LabResultViewRepository) Update(view *models.LabResultView) error {
	return r.db.Save(view).Error
}

// Delete faz soft delete de uma view
func (r *LabResultViewRepository) Delete(id uuid.UUID) error {
	result := r.db.Delete(&models.LabResultView{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrLabResultViewNotFound
	}
	return nil
}

// UpdateItems substitui TODOS os items de uma view (transação atômica)
func (r *LabResultViewRepository) UpdateItems(viewID uuid.UUID, items []models.LabResultViewItem) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 1. Verificar se view existe
		var view models.LabResultView
		if err := tx.Where("id = ?", viewID).First(&view).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrLabResultViewNotFound
			}
			return err
		}

		// 2. DELETE todos items existentes (soft delete)
		if err := tx.Where("lab_result_view_id = ?", viewID).Delete(&models.LabResultViewItem{}).Error; err != nil {
			return err
		}

		// 3. CREATE novos items (ordem preservada)
		if len(items) > 0 {
			for i := range items {
				items[i].LabResultViewID = viewID
			}
			if err := tx.Create(&items).Error; err != nil {
				return err
			}
		}

		// 4. UPDATE view.updated_at
		if err := tx.Model(&models.LabResultView{}).Where("id = ?", viewID).Update("updated_at", gorm.Expr("CURRENT_TIMESTAMP")).Error; err != nil {
			return err
		}

		return nil
	})
}

// Search busca views por nome ou descrição
func (r *LabResultViewRepository) Search(query string, withItems bool) ([]models.LabResultView, error) {
	var views []models.LabResultView
	db := r.db

	if withItems {
		db = db.Preload("Items.LabTestDefinition")
	}

	searchPattern := "%" + query + "%"
	err := db.Where("name ILIKE ? OR description ILIKE ?", searchPattern, searchPattern).
		Order("display_order ASC, name ASC").
		Find(&views).Error

	if err != nil {
		return nil, err
	}

	return views, nil
}
