package workers

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
	"gorm.io/gorm"
)

// EmbeddingWorker processa fila de embeddings em background
// Poll embedding_queue e gera embeddings para articles e score_items
type EmbeddingWorker struct {
	db               *gorm.DB
	embeddingService *services.EmbeddingService
	chunkingService  *services.ChunkingService
	interval         time.Duration // Intervalo entre polls (default: 10s)
	maxConcurrent    int           // Max jobs simultâneos (default: 5)
}

// NewEmbeddingWorker cria nova instância do worker
func NewEmbeddingWorker(
	db *gorm.DB,
	embeddingService *services.EmbeddingService,
	chunkingService *services.ChunkingService,
) *EmbeddingWorker {
	return &EmbeddingWorker{
		db:               db,
		embeddingService: embeddingService,
		chunkingService:  chunkingService,
		interval:         10 * time.Second, // Poll a cada 10 segundos
		maxConcurrent:    5,                // Processar até 5 jobs simultaneamente
	}
}

// Start inicia o worker (blocking)
// Executa em loop até context ser cancelado
func (w *EmbeddingWorker) Start(ctx context.Context) error {
	fmt.Println("🤖 EmbeddingWorker started")
	fmt.Printf("   Poll interval: %v\n", w.interval)
	fmt.Printf("   Max concurrent: %d\n\n", w.maxConcurrent)

	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	// Processar fila imediatamente na inicialização
	w.processQueue(ctx)

	// Loop principal
	for {
		select {
		case <-ctx.Done():
			fmt.Println("🛑 EmbeddingWorker stopped (context cancelled)")
			return ctx.Err()

		case <-ticker.C:
			w.processQueue(ctx)
		}
	}
}

// processQueue busca e processa jobs pendentes da fila
func (w *EmbeddingWorker) processQueue(ctx context.Context) {
	// Buscar jobs pendentes (status = 'pending' ou 'failed' com retry_count < 5)
	var jobs []models.EmbeddingQueue
	err := w.db.WithContext(ctx).
		Where("status = ? OR (status = ? AND retry_count < 5)", "pending", "failed").
		Order("created_at ASC").
		Limit(w.maxConcurrent).
		Find(&jobs).Error

	if err != nil {
		fmt.Printf("❌ Error fetching jobs from queue: %v\n", err)
		return
	}

	if len(jobs) == 0 {
		// Nenhum job pendente (normal)
		return
	}

	fmt.Printf("📋 Found %d pending jobs\n", len(jobs))

	// Processar jobs em paralelo (até maxConcurrent)
	for _, job := range jobs {
		// Clonar job para evitar closure issues
		j := job

		// Processar em goroutine
		go w.processJob(ctx, &j)
	}
}

// processJob processa um único job da fila
func (w *EmbeddingWorker) processJob(ctx context.Context, job *models.EmbeddingQueue) {
	jobID := job.ID.String()[:8] // Primeiros 8 chars do UUID para logging

	fmt.Printf("⚙️  Processing job %s (type: %s, entity: %s)\n",
		jobID, job.EntityType, job.EntityID.String()[:8])

	// Marcar como "processing"
	if err := job.MarkAsProcessing(w.db); err != nil {
		fmt.Printf("❌ Job %s: Failed to mark as processing: %v\n", jobID, err)
		return
	}

	// Processar baseado no tipo de entidade
	var err error
	switch job.EntityType {
	case "article":
		err = w.processArticle(ctx, job.EntityID)
	case "score_item":
		err = w.processScoreItem(ctx, job.EntityID)
	default:
		err = fmt.Errorf("unknown entity type: %s", job.EntityType)
	}

	// Atualizar status do job
	if err != nil {
		fmt.Printf("❌ Job %s: Failed: %v\n", jobID, err)
		if markErr := job.MarkAsFailed(w.db, err.Error()); markErr != nil {
			fmt.Printf("⚠️  Job %s: Failed to mark as failed: %v\n", jobID, markErr)
		}
	} else {
		fmt.Printf("✅ Job %s: Completed successfully\n", jobID)
		if markErr := job.MarkAsCompleted(w.db); markErr != nil {
			fmt.Printf("⚠️  Job %s: Failed to mark as completed: %v\n", jobID, markErr)
		}
	}
}

// processArticle gera embeddings para um artigo
func (w *EmbeddingWorker) processArticle(ctx context.Context, articleID uuid.UUID) error {
	// 1. Buscar article
	var article models.Article
	if err := w.db.First(&article, "id = ?", articleID).Error; err != nil {
		return fmt.Errorf("article not found: %w", err)
	}

	// 2. Atualizar status para "processing"
	if err := w.db.Model(&article).Update("embedding_status", "processing").Error; err != nil {
		return fmt.Errorf("failed to update article status: %w", err)
	}

	// 3. Chunkar artigo
	abstract := ""
	if article.Abstract != nil {
		abstract = *article.Abstract
	}

	fullContent := ""
	if article.FullContent != nil {
		fullContent = *article.FullContent
	}

	chunks, err := w.chunkingService.ChunkArticle(fullContent, abstract)
	if err != nil {
		w.db.Model(&article).Update("embedding_status", "failed")
		return fmt.Errorf("failed to chunk article: %w", err)
	}

	fmt.Printf("   📄 Article chunked into %d chunks\n", len(chunks))

	// 4. Gerar embeddings para cada chunk (em batch se possível)
	chunkTexts := make([]string, len(chunks))
	for i, chunk := range chunks {
		chunkTexts[i] = chunk.Text
	}

	// Processar em batches de 20 (ótimo para OpenAI API)
	batchSize := 20
	allEmbeddings := [][]float32{}

	for i := 0; i < len(chunkTexts); i += batchSize {
		end := i + batchSize
		if end > len(chunkTexts) {
			end = len(chunkTexts)
		}

		batch := chunkTexts[i:end]
		embeddings, err := w.embeddingService.BatchGenerateEmbeddings(ctx, batch)
		if err != nil {
			w.db.Model(&article).Update("embedding_status", "failed")
			return fmt.Errorf("failed to generate embeddings (batch %d): %w", i/batchSize, err)
		}

		allEmbeddings = append(allEmbeddings, embeddings...)
	}

	fmt.Printf("   🧠 Generated %d embeddings\n", len(allEmbeddings))

	// 5. Salvar embeddings no banco (transaction)
	err = w.db.Transaction(func(tx *gorm.DB) error {
		// Deletar embeddings antigos se existirem
		if err := tx.Where("article_id = ?", articleID).Delete(&models.ArticleEmbedding{}).Error; err != nil {
			return fmt.Errorf("failed to delete old embeddings: %w", err)
		}

		// Inserir novos embeddings
		for i, chunk := range chunks {
			embedding := &models.ArticleEmbedding{
				ArticleID:     articleID,
				ChunkIndex:    chunk.Index,
				ChunkText:     chunk.Text,
				ChunkMetadata: chunk.Metadata,
			}

			// Converter []float32 para pgvector.Vector
			if err := embedding.SetEmbeddingFromSlice(allEmbeddings[i]); err != nil {
				return fmt.Errorf("failed to set embedding for chunk %d: %w", i, err)
			}

			if err := tx.Create(embedding).Error; err != nil {
				return fmt.Errorf("failed to save embedding for chunk %d: %w", i, err)
			}
		}

		// Atualizar article metadata
		now := time.Now()
		updates := map[string]interface{}{
			"embedding_status":  "completed",
			"chunk_count":       len(chunks),
			"last_embedded_at":  now,
		}

		if err := tx.Model(&article).Updates(updates).Error; err != nil {
			return fmt.Errorf("failed to update article metadata: %w", err)
		}

		return nil
	})

	if err != nil {
		w.db.Model(&article).Update("embedding_status", "failed")
		return err
	}

	fmt.Printf("   💾 Saved %d embeddings to database\n", len(chunks))
	return nil
}

// processScoreItem gera embedding para um score item
func (w *EmbeddingWorker) processScoreItem(ctx context.Context, scoreItemID uuid.UUID) error {
	// 1. Buscar score item
	var scoreItem models.ScoreItem
	if err := w.db.First(&scoreItem, "id = ?", scoreItemID).Error; err != nil {
		return fmt.Errorf("score item not found: %w", err)
	}

	// 2. Gerar texto combinado
	textSource := w.chunkingService.ChunkScoreItem(
		scoreItem.Name,
		scoreItem.ClinicalRelevance,
		scoreItem.PatientExplanation,
		scoreItem.Conduct,
	)

	if textSource == "" {
		return fmt.Errorf("failed to generate text for score item (empty)")
	}

	fmt.Printf("   🎯 ScoreItem text generated (%d chars)\n", len(textSource))

	// 3. Gerar embedding
	embedding, err := w.embeddingService.GenerateEmbedding(ctx, textSource)
	if err != nil {
		return fmt.Errorf("failed to generate embedding: %w", err)
	}

	fmt.Printf("   🧠 Embedding generated (%d dims)\n", len(embedding))

	// 4. Salvar embedding no banco (upsert)
	err = w.db.Transaction(func(tx *gorm.DB) error {
		// Verificar se já existe
		var existing models.ScoreItemEmbedding
		err := tx.Where("score_item_id = ?", scoreItemID).First(&existing).Error

		if err == nil {
			// Já existe - atualizar
			existing.TextSource = textSource
			if err := existing.SetEmbeddingFromSlice(embedding); err != nil {
				return fmt.Errorf("failed to set embedding: %w", err)
			}
			if err := tx.Save(&existing).Error; err != nil {
				return fmt.Errorf("failed to update embedding: %w", err)
			}
		} else if err == gorm.ErrRecordNotFound {
			// Não existe - criar
			newEmbedding := &models.ScoreItemEmbedding{
				ScoreItemID: scoreItemID,
				TextSource:  textSource,
			}
			if err := newEmbedding.SetEmbeddingFromSlice(embedding); err != nil {
				return fmt.Errorf("failed to set embedding: %w", err)
			}
			if err := tx.Create(newEmbedding).Error; err != nil {
				return fmt.Errorf("failed to create embedding: %w", err)
			}
		} else {
			return fmt.Errorf("failed to check existing embedding: %w", err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Printf("   💾 Saved embedding to database\n")
	return nil
}

// GetQueueStats retorna estatísticas da fila
func (w *EmbeddingWorker) GetQueueStats() (map[string]int64, error) {
	stats := make(map[string]int64)

	// Count por status
	statuses := []string{"pending", "processing", "completed", "failed"}
	for _, status := range statuses {
		var count int64
		if err := w.db.Model(&models.EmbeddingQueue{}).
			Where("status = ?", status).
			Count(&count).Error; err != nil {
			return nil, err
		}
		stats[status] = count
	}

	// Count total
	var total int64
	if err := w.db.Model(&models.EmbeddingQueue{}).Count(&total).Error; err != nil {
		return nil, err
	}
	stats["total"] = total

	return stats, nil
}
