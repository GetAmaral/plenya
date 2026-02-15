package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// AnamnesisRepository handles database operations for anamnesis records
type AnamnesisRepository struct {
	db *gorm.DB
}

// NewAnamnesisRepository creates a new anamnesis repository instance
func NewAnamnesisRepository(db *gorm.DB) *AnamnesisRepository {
	return &AnamnesisRepository{db: db}
}

// GetHistoricalItemsByScoreItemID retrieves a map of [ScoreItemID]→AnamnesisItem (most recent)
// Searches entire patient history using DISTINCT ON (PostgreSQL feature)
// Returns only items with non-null NumericValue
func (r *AnamnesisRepository) GetHistoricalItemsByScoreItemID(patientID uuid.UUID) (map[uuid.UUID]models.AnamnesisItem, error) {
	var items []models.AnamnesisItem

	// Optimized query with DISTINCT ON (PostgreSQL)
	// Returns the most recent item for each ScoreItemID
	//
	// Query explanation:
	// 1. JOIN anamnesis_items with anamnesis to get patient_id and consultation_date
	// 2. Filter by patient_id and ensure we have numeric values
	// 3. DISTINCT ON (ai.score_item_id) ensures we get only one item per score_item_id
	// 4. ORDER BY ai.score_item_id, a.consultation_date DESC, ai.created_at DESC ensures we get the most recent
	err := r.db.
		Raw(`
			SELECT DISTINCT ON (ai.score_item_id)
				ai.id,
				ai.anamnesis_id,
				ai.score_item_id,
				ai.text_value,
				ai.numeric_value,
				ai."order",
				ai.created_at,
				ai.updated_at,
				ai.deleted_at
			FROM anamnesis_items ai
			INNER JOIN anamnesis a ON ai.anamnesis_id = a.id
			WHERE a.patient_id = ?
				AND ai.numeric_value IS NOT NULL
				AND ai.deleted_at IS NULL
				AND a.deleted_at IS NULL
			ORDER BY ai.score_item_id, a.consultation_date DESC, ai.created_at DESC
		`, patientID).
		Scan(&items).Error

	if err != nil {
		return nil, err
	}

	// Preload ScoreItem for each item (excluding soft-deleted)
	for i := range items {
		var scoreItem models.ScoreItem
		if err := r.db.Where("deleted_at IS NULL").First(&scoreItem, "id = ?", items[i].ScoreItemID).Error; err == nil {
			items[i].ScoreItem = &scoreItem
		}
	}

	// Convert to map [ScoreItemID]→AnamnesisItem
	itemMap := make(map[uuid.UUID]models.AnamnesisItem)
	for _, item := range items {
		itemMap[item.ScoreItemID] = item
	}

	return itemMap, nil
}
