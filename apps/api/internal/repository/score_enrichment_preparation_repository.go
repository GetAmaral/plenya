package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"gorm.io/gorm"
)

// ScoreEnrichmentPreparationRepository gerencia operações de preparação de enrichment
type ScoreEnrichmentPreparationRepository struct {
	db *gorm.DB
}

// NewScoreEnrichmentPreparationRepository cria nova instância
func NewScoreEnrichmentPreparationRepository(db *gorm.DB) *ScoreEnrichmentPreparationRepository {
	return &ScoreEnrichmentPreparationRepository{db: db}
}

// Create cria nova preparação (ou substitui se já existir devido a UNIQUE constraint)
func (r *ScoreEnrichmentPreparationRepository) Create(prep *models.ScoreItemEnrichmentPreparation) error {
	// Upsert: INSERT ... ON CONFLICT DO UPDATE
	// Se já existe preparação para este score_item_id, atualiza
	err := r.db.Exec(`
		INSERT INTO score_item_enrichment_preparation
		(id, score_item_id, selected_chunks, metadata, status, created_at, updated_at)
		VALUES (?, ?, ?::jsonb, ?::jsonb, ?, NOW(), NOW())
		ON CONFLICT (score_item_id)
		DO UPDATE SET
			selected_chunks = EXCLUDED.selected_chunks,
			metadata = EXCLUDED.metadata,
			status = EXCLUDED.status,
			updated_at = NOW()
	`, prep.ID, prep.ScoreItemID, prep.SelectedChunks, prep.Metadata, prep.Status).Error

	return err
}

// FindByScoreItemID busca preparação por score_item_id
func (r *ScoreEnrichmentPreparationRepository) FindByScoreItemID(scoreItemID uuid.UUID) (*models.ScoreItemEnrichmentPreparation, error) {
	var prep models.ScoreItemEnrichmentPreparation
	err := r.db.
		Preload("ScoreItem.Subgroup.Group").
		Where("score_item_id = ?", scoreItemID).
		First(&prep).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Não é erro, apenas não existe
		}
		return nil, fmt.Errorf("failed to fetch preparation: %w", err)
	}

	return &prep, nil
}

// FindByID busca preparação por ID
func (r *ScoreEnrichmentPreparationRepository) FindByID(id uuid.UUID) (*models.ScoreItemEnrichmentPreparation, error) {
	var prep models.ScoreItemEnrichmentPreparation
	err := r.db.
		Preload("ScoreItem.Subgroup.Group").
		Where("id = ?", id).
		First(&prep).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to fetch preparation: %w", err)
	}

	return &prep, nil
}

// FindByStatus busca preparações por status (ready, processing, completed, failed)
func (r *ScoreEnrichmentPreparationRepository) FindByStatus(status string, limit int) ([]models.ScoreItemEnrichmentPreparation, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}

	var preps []models.ScoreItemEnrichmentPreparation
	err := r.db.
		Preload("ScoreItem.Subgroup.Group").
		Where("status = ?", status).
		Order("created_at DESC").
		Limit(limit).
		Find(&preps).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch preparations by status: %w", err)
	}

	return preps, nil
}

// UpdateStatus atualiza status da preparação
func (r *ScoreEnrichmentPreparationRepository) UpdateStatus(id uuid.UUID, status string, errorMsg *string) error {
	updates := map[string]interface{}{
		"status":     status,
		"updated_at": gorm.Expr("NOW()"),
	}

	if errorMsg != nil {
		updates["error_message"] = *errorMsg
	}

	err := r.db.Model(&models.ScoreItemEnrichmentPreparation{}).
		Where("id = ?", id).
		Updates(updates).Error

	return err
}

// Delete remove preparação
func (r *ScoreEnrichmentPreparationRepository) Delete(id uuid.UUID) error {
	err := r.db.Delete(&models.ScoreItemEnrichmentPreparation{}, "id = ?", id).Error
	return err
}

// DeleteByScoreItemID remove preparação por score_item_id
func (r *ScoreEnrichmentPreparationRepository) DeleteByScoreItemID(scoreItemID uuid.UUID) error {
	err := r.db.Delete(&models.ScoreItemEnrichmentPreparation{}, "score_item_id = ?", scoreItemID).Error
	return err
}

// Count retorna total de preparações
func (r *ScoreEnrichmentPreparationRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.ScoreItemEnrichmentPreparation{}).Count(&count).Error
	return count, err
}

// CountByStatus retorna total de preparações por status
func (r *ScoreEnrichmentPreparationRepository) CountByStatus(status string) (int64, error) {
	var count int64
	err := r.db.Model(&models.ScoreItemEnrichmentPreparation{}).
		Where("status = ?", status).
		Count(&count).Error
	return count, err
}
