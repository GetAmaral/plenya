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

// PrepareChunksAdaptive busca chunks com threshold adaptativo (fallback progressivo)
// Garante pelo menos 15 chunks ou usa o threshold mais baixo possível
func (s *ScoreEnrichmentPreparationService) PrepareChunksAdaptive(
	scoreItemID uuid.UUID,
) (*models.ScoreItemEnrichmentPreparation, error) {
	targetChunks := 20  // Alvo mínimo ideal
	maxLimit := 50      // Buscar até 50 chunks

	// Tentativas progressivas de threshold (do mais rigoroso ao mais permissivo)
	thresholds := []struct {
		value float64
		label string
	}{
		{0.35, "normal"},
		{0.30, "permissivo"},
		{0.25, "muito permissivo"},
		{0.20, "extremo"},
		{0.15, "fallback final"},
	}

	var bestChunks []repository.ChunkSearchResult
	var usedThreshold float64
	var thresholdLabel string

	// Validar que score_item existe
	var scoreItem models.ScoreItem
	err := s.db.Preload("Subgroup.Group").First(&scoreItem, "id = ?", scoreItemID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("score_item not found: %s", scoreItemID)
		}
		return nil, fmt.Errorf("failed to fetch score_item: %w", err)
	}

	// Tentar cada threshold até encontrar >= targetChunks
	for _, t := range thresholds {
		chunks, err := s.vectorRepo.FindTopChunksForScoreItem(scoreItemID, maxLimit, t.value)
		if err != nil {
			continue
		}

		if len(chunks) >= targetChunks {
			bestChunks = chunks
			usedThreshold = t.value
			thresholdLabel = t.label
			break
		}

		// Guardar o melhor resultado mesmo se < targetChunks
		if len(chunks) > len(bestChunks) {
			bestChunks = chunks
			usedThreshold = t.value
			thresholdLabel = t.label
		}
	}

	if len(bestChunks) == 0 {
		return nil, fmt.Errorf("no chunks found even with threshold 0.15")
	}

	// Limitar chunks ao máximo desejado (30 para qualidade)
	finalLimit := 30
	if usedThreshold >= 0.30 {
		finalLimit = 30 // Threshold bom, 30 chunks suficientes
	} else if usedThreshold >= 0.20 {
		finalLimit = 35 // Threshold mais baixo, incluir mais chunks
	} else {
		finalLimit = 40 // Fallback extremo, incluir ainda mais
	}

	if len(bestChunks) > finalLimit {
		bestChunks = bestChunks[:finalLimit]
	}

	return s.savePreparation(scoreItemID, bestChunks, usedThreshold, thresholdLabel)
}

// PrepareChunks busca chunks relevantes via RAG e salva no banco (SEM IA)
// Retorna preparation criado ou erro
// DEPRECATED: Use PrepareChunksAdaptive para melhor qualidade
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

	// Delegar para savePreparation com threshold usado
	return s.savePreparation(scoreItemID, chunks, minSimilarity, "manual")
}

// savePreparation extrai lógica comum de salvar preparation (com metadata enriquecido)
func (s *ScoreEnrichmentPreparationService) savePreparation(
	scoreItemID uuid.UUID,
	chunks []repository.ChunkSearchResult,
	usedThreshold float64,
	thresholdLabel string,
) (*models.ScoreItemEnrichmentPreparation, error) {
	// 1. Converter chunks para JSON structure
	chunksData := make([]map[string]interface{}, len(chunks))
	articleIDs := make(map[uuid.UUID]bool)
	sectionsCount := make(map[string]int)
	yearStats := make([]int, 0, len(chunks))

	var totalSimilarity float64
	var totalWordCount int
	minSimilarity := 1.0
	maxSimilarity := 0.0

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
		totalWordCount += chunk.WordCount

		// Track min/max similarity
		if chunk.Similarity < minSimilarity {
			minSimilarity = chunk.Similarity
		}
		if chunk.Similarity > maxSimilarity {
			maxSimilarity = chunk.Similarity
		}

		// Track years for recency stats
		if chunk.ArticleYear > 0 {
			yearStats = append(yearStats, chunk.ArticleYear)
		}
	}

	// 2. Calcular metadata enriquecido
	avgSimilarity := totalSimilarity / float64(len(chunks))
	avgChunkLength := 0
	if len(chunks) > 0 {
		avgChunkLength = totalWordCount / len(chunks)
	}

	// Recency stats
	newestYear := 0
	oldestYear := 9999
	var totalYear int
	for _, year := range yearStats {
		if year > newestYear {
			newestYear = year
		}
		if year < oldestYear {
			oldestYear = year
		}
		totalYear += year
	}
	avgYear := 0
	if len(yearStats) > 0 {
		avgYear = totalYear / len(yearStats)
	}

	// Quality grade
	qualityGrade := "poor"
	if len(chunks) >= 25 && avgSimilarity >= 0.6 {
		qualityGrade = "excellent"
	} else if len(chunks) >= 20 && avgSimilarity >= 0.5 {
		qualityGrade = "good"
	} else if len(chunks) >= 15 && avgSimilarity >= 0.4 {
		qualityGrade = "fair"
	}

	// 3. Criar preparation com metadata enriquecido
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
			"min_similarity":        minSimilarity,
			"max_similarity":        maxSimilarity,
			"sections_distribution": sectionsCount,
			"total_word_count":      totalWordCount,
			"avg_chunk_length":      avgChunkLength,
			"recency_stats": map[string]interface{}{
				"newest_year": newestYear,
				"oldest_year": oldestYear,
				"avg_year":    avgYear,
			},
			"quality_grade":  qualityGrade,
			"threshold_used": usedThreshold,
			"threshold_type": thresholdLabel,
		},
		Status: "ready",
	}

	// 4. Salvar no banco (upsert)
	err := s.prepRepo.Create(prep)
	if err != nil {
		return nil, fmt.Errorf("failed to save preparation: %w", err)
	}

	// 5. Recarregar com ScoreItem preloaded
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
