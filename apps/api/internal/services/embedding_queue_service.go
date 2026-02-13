package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"gorm.io/gorm"
)

// EmbeddingQueueService gerencia a fila de processamento de embeddings
type EmbeddingQueueService struct {
	db *gorm.DB
}

// NewEmbeddingQueueService cria nova instância do serviço
func NewEmbeddingQueueService(db *gorm.DB) *EmbeddingQueueService {
	return &EmbeddingQueueService{db: db}
}

// QueueArticle adiciona artigo à fila de processamento
// Retorna nil se já existe na fila (idempotente)
func (s *EmbeddingQueueService) QueueArticle(articleID uuid.UUID) error {
	// Verificar se já existe job pendente/processing para este artigo
	var existing models.EmbeddingQueue
	err := s.db.Where(
		"entity_type = ? AND entity_id = ? AND status IN (?)",
		"article",
		articleID,
		[]string{"pending", "processing"},
	).First(&existing).Error

	if err == nil {
		// Já existe - retornar sucesso (idempotente)
		return nil
	}

	if err != gorm.ErrRecordNotFound {
		// Erro ao buscar
		return fmt.Errorf("failed to check existing queue job: %w", err)
	}

	// Não existe - criar novo job
	job := &models.EmbeddingQueue{
		EntityType: "article",
		EntityID:   articleID,
		Status:     "pending",
		RetryCount: 0,
	}

	if err := s.db.Create(job).Error; err != nil {
		return fmt.Errorf("failed to create queue job: %w", err)
	}

	fmt.Printf("📥 Article %s queued for embedding\n", articleID.String()[:8])
	return nil
}

// QueueScoreItem adiciona score item à fila de processamento
func (s *EmbeddingQueueService) QueueScoreItem(scoreItemID uuid.UUID) error {
	// Verificar se já existe job pendente/processing
	var existing models.EmbeddingQueue
	err := s.db.Where(
		"entity_type = ? AND entity_id = ? AND status IN (?)",
		"score_item",
		scoreItemID,
		[]string{"pending", "processing"},
	).First(&existing).Error

	if err == nil {
		// Já existe
		return nil
	}

	if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("failed to check existing queue job: %w", err)
	}

	// Criar novo job
	job := &models.EmbeddingQueue{
		EntityType: "score_item",
		EntityID:   scoreItemID,
		Status:     "pending",
		RetryCount: 0,
	}

	if err := s.db.Create(job).Error; err != nil {
		return fmt.Errorf("failed to create queue job: %w", err)
	}

	fmt.Printf("📥 ScoreItem %s queued for embedding\n", scoreItemID.String()[:8])
	return nil
}

// QueueAllArticles adiciona todos os artigos sem embeddings à fila (backfill)
func (s *EmbeddingQueueService) QueueAllArticles() (int, error) {
	// Buscar artigos sem embeddings (embedding_status = 'pending' ou NULL)
	var articles []models.Article
	err := s.db.Where(
		"embedding_status = ? OR embedding_status IS NULL",
		"pending",
	).Find(&articles).Error

	if err != nil {
		return 0, fmt.Errorf("failed to fetch articles: %w", err)
	}

	queued := 0
	for _, article := range articles {
		if err := s.QueueArticle(article.ID); err != nil {
			fmt.Printf("⚠️  Failed to queue article %s: %v\n", article.ID.String()[:8], err)
			continue
		}
		queued++
	}

	fmt.Printf("📦 Queued %d articles for embedding (backfill)\n", queued)
	return queued, nil
}

// QueueAllScoreItems adiciona todos os score items sem embeddings à fila
func (s *EmbeddingQueueService) QueueAllScoreItems() (int, error) {
	// Buscar score items que não têm embedding
	var scoreItems []models.ScoreItem
	err := s.db.
		Where("id NOT IN (SELECT score_item_id FROM score_item_embeddings)").
		Find(&scoreItems).Error

	if err != nil {
		return 0, fmt.Errorf("failed to fetch score items: %w", err)
	}

	queued := 0
	for _, item := range scoreItems {
		if err := s.QueueScoreItem(item.ID); err != nil {
			fmt.Printf("⚠️  Failed to queue score item %s: %v\n", item.ID.String()[:8], err)
			continue
		}
		queued++
	}

	fmt.Printf("📦 Queued %d score items for embedding (backfill)\n", queued)
	return queued, nil
}

// RetryFailed reprocessa jobs falhados (retry_count < 5)
func (s *EmbeddingQueueService) RetryFailed() (int, error) {
	result := s.db.Model(&models.EmbeddingQueue{}).
		Where("status = ? AND retry_count < 5", "failed").
		Update("status", "pending")

	if result.Error != nil {
		return 0, fmt.Errorf("failed to retry failed jobs: %w", result.Error)
	}

	fmt.Printf("🔄 Retrying %d failed jobs\n", result.RowsAffected)
	return int(result.RowsAffected), nil
}

// ClearCompleted remove jobs completados (limpeza)
func (s *EmbeddingQueueService) ClearCompleted() (int, error) {
	result := s.db.Where("status = ?", "completed").Delete(&models.EmbeddingQueue{})

	if result.Error != nil {
		return 0, fmt.Errorf("failed to clear completed jobs: %w", result.Error)
	}

	fmt.Printf("🧹 Cleared %d completed jobs\n", result.RowsAffected)
	return int(result.RowsAffected), nil
}

// GetStats retorna estatísticas da fila
func (s *EmbeddingQueueService) GetStats() (map[string]int64, error) {
	stats := make(map[string]int64)

	// Count por status
	statuses := []string{"pending", "processing", "completed", "failed"}
	for _, status := range statuses {
		var count int64
		if err := s.db.Model(&models.EmbeddingQueue{}).
			Where("status = ?", status).
			Count(&count).Error; err != nil {
			return nil, err
		}
		stats[status] = count
	}

	// Count total
	var total int64
	if err := s.db.Model(&models.EmbeddingQueue{}).Count(&total).Error; err != nil {
		return nil, err
	}
	stats["total"] = total

	// Count por entity type
	var articleCount, scoreItemCount int64
	s.db.Model(&models.EmbeddingQueue{}).Where("entity_type = ?", "article").Count(&articleCount)
	s.db.Model(&models.EmbeddingQueue{}).Where("entity_type = ?", "score_item").Count(&scoreItemCount)
	stats["articles"] = articleCount
	stats["score_items"] = scoreItemCount

	return stats, nil
}
