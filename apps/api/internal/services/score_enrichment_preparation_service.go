package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// ScoreEnrichmentPreparationService prepara chunks para enrichment (ETAPA 1 - SEM IA)
type ScoreEnrichmentPreparationService struct {
	db         *gorm.DB
	vectorRepo *repository.ArticleVectorRepository
	prepRepo   *repository.ScoreEnrichmentPreparationRepository
}

// NewScoreEnrichmentPreparationService cria nova instância
func NewScoreEnrichmentPreparationService(db *gorm.DB) *ScoreEnrichmentPreparationService {
	return &ScoreEnrichmentPreparationService{
		db:         db,
		vectorRepo: repository.NewArticleVectorRepository(db),
		prepRepo:   repository.NewScoreEnrichmentPreparationRepository(db),
	}
}

// PrepareChunks busca chunks relevantes via RAG e salva no banco (SEM IA)
// Retorna preparation criado ou erro
func (s *ScoreEnrichmentPreparationService) PrepareChunks(
	scoreItemID uuid.UUID,
	limit int,              // Total de chunks (default: 20)
	minSimilarity float64,  // Threshold (default: 0.65)
) (*models.ScoreItemEnrichmentPreparation, error) {
	// Validar que score_item existe
	var scoreItem models.ScoreItem
	err := s.db.Preload("Subgroup.Group").First(&scoreItem, "id = ?", scoreItemID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("score_item not found: %s", scoreItemID)
		}
		return nil, fmt.Errorf("failed to fetch score_item: %w", err)
	}

	// Defaults
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if minSimilarity < 0 || minSimilarity > 1 {
		minSimilarity = 0.65
	}

	// 1. Buscar chunks via RAG
	chunks, err := s.vectorRepo.FindTopChunksForScoreItem(scoreItemID, limit, minSimilarity)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch chunks: %w", err)
	}

	if len(chunks) == 0 {
		return nil, fmt.Errorf("no chunks found with similarity >= %.2f", minSimilarity)
	}

	// 2. Converter chunks para JSON structure
	chunksData := make([]map[string]interface{}, len(chunks))
	articleIDs := make(map[uuid.UUID]bool)
	sectionsCount := make(map[string]int)
	var totalSimilarity float64

	for i, chunk := range chunks {
		chunksData[i] = map[string]interface{}{
			"article_id":    chunk.ArticleID.String(),
			"article_title": chunk.ArticleTitle,
			"article_year":  chunk.ArticleYear,
			"journal":       chunk.Journal,
			"chunk_index":   chunk.ChunkIndex,
			"chunk_text":    chunk.ChunkText,
			"section":       chunk.Section,
			"similarity":    chunk.Similarity,
			"word_count":    chunk.WordCount,
		}

		if chunk.PageRange != nil {
			chunksData[i]["page_range"] = *chunk.PageRange
		}

		articleIDs[chunk.ArticleID] = true
		sectionsCount[chunk.Section]++
		totalSimilarity += chunk.Similarity
	}

	// 3. Calcular metadata
	avgSimilarity := totalSimilarity / float64(len(chunks))

	// 4. Criar preparation usando datatypes.JSONMap (map[string]interface{})
	// SelectedChunks precisa ser wrapped em um objeto porque datatypes.JSONMap é map
	chunksInterface := make([]interface{}, len(chunksData))
	for i, chunk := range chunksData {
		chunksInterface[i] = chunk
	}

	prep := &models.ScoreItemEnrichmentPreparation{
		ID:          uuid.Must(uuid.NewV7()),
		ScoreItemID: scoreItemID,
		SelectedChunks: datatypes.JSONMap{
			"items": chunksInterface,
		},
		Metadata: datatypes.JSONMap{
			"total_chunks":          len(chunks),
			"articles_count":        len(articleIDs),
			"avg_similarity":        avgSimilarity,
			"sections_distribution": sectionsCount,
		},
		Status: "ready",
	}

	// 6. Salvar no banco (upsert)
	err = s.prepRepo.Create(prep)
	if err != nil {
		return nil, fmt.Errorf("failed to save preparation: %w", err)
	}

	// 7. Recarregar com ScoreItem preloaded
	prep, err = s.prepRepo.FindByScoreItemID(scoreItemID)
	if err != nil {
		return nil, fmt.Errorf("failed to reload preparation: %w", err)
	}

	return prep, nil
}

// GetPreparation busca preparation por score_item_id
func (s *ScoreEnrichmentPreparationService) GetPreparation(scoreItemID uuid.UUID) (*models.ScoreItemEnrichmentPreparation, error) {
	prep, err := s.prepRepo.FindByScoreItemID(scoreItemID)
	if err != nil {
		return nil, fmt.Errorf("failed to get preparation: %w", err)
	}

	if prep == nil {
		return nil, fmt.Errorf("preparation not found for score_item: %s", scoreItemID)
	}

	return prep, nil
}

// DeletePreparation remove preparation por score_item_id
func (s *ScoreEnrichmentPreparationService) DeletePreparation(scoreItemID uuid.UUID) error {
	return s.prepRepo.DeleteByScoreItemID(scoreItemID)
}

// GetStats retorna estatísticas de preparações
func (s *ScoreEnrichmentPreparationService) GetStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// Total
	total, err := s.prepRepo.Count()
	if err != nil {
		return nil, err
	}
	stats["total"] = total

	// Por status
	ready, _ := s.prepRepo.CountByStatus("ready")
	processing, _ := s.prepRepo.CountByStatus("processing")
	completed, _ := s.prepRepo.CountByStatus("completed")
	failed, _ := s.prepRepo.CountByStatus("failed")

	stats["by_status"] = map[string]int64{
		"ready":      ready,
		"processing": processing,
		"completed":  completed,
		"failed":     failed,
	}

	return stats, nil
}
