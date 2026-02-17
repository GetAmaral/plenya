package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// EmbeddingQueueItem represents a pending embedding regeneration
type EmbeddingQueueItem struct {
	ID         uuid.UUID
	EntityType string
	EntityID   uuid.UUID
	Status     string
	RetryCount int
	MaxRetries int
}

// RegenerateEmbeddings - Worker para processar fila de embeddings stale
func main() {
	godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	embeddingService := services.NewEmbeddingService(cfg, db)
	chunkingService := services.NewChunkingService()

	fmt.Println("🔄 Starting embedding regeneration worker...")
	fmt.Println("📋 Checking queue...")

	// Get pending items from queue (ordered by priority DESC, created_at ASC)
	var queueItems []EmbeddingQueueItem
	err = db.Raw(`
		SELECT id, entity_type, entity_id, status, retry_count, max_retries
		FROM embedding_queue
		WHERE status = 'pending' AND retry_count < max_retries
		ORDER BY priority DESC, created_at ASC
		LIMIT 100
	`).Scan(&queueItems).Error

	if err != nil {
		log.Fatalf("Failed to fetch queue: %v", err)
	}

	if len(queueItems) == 0 {
		fmt.Println("✅ Queue is empty - nothing to process")
		return
	}

	fmt.Printf("📦 Found %d items in queue\n\n", len(queueItems))

	successCount := 0
	failCount := 0

	for i, item := range queueItems {
		fmt.Printf("[%d/%d] Processing %s %s...\n", i+1, len(queueItems), item.EntityType, item.EntityID)

		// Mark as processing
		db.Exec(`
			UPDATE embedding_queue
			SET status = 'processing', processed_at = NOW()
			WHERE id = ?
		`, item.ID)

		var err error
		if item.EntityType == "score_item" {
			err = regenerateScoreItemEmbedding(db, embeddingService, chunkingService, item.EntityID)
		} else if item.EntityType == "article" {
			err = regenerateArticleEmbedding(db, embeddingService, chunkingService, item.EntityID)
		} else {
			err = fmt.Errorf("unknown entity_type: %s", item.EntityType)
		}

		if err != nil {
			failCount++
			newRetryCount := item.RetryCount + 1

			// Mark as failed or retry
			if newRetryCount >= item.MaxRetries {
				fmt.Printf("  ❌ FAILED (max retries): %v\n", err)
				db.Exec(`
					UPDATE embedding_queue
					SET status = 'failed', error_message = ?, retry_count = ?
					WHERE id = ?
				`, err.Error(), newRetryCount, item.ID)

				// Log to audit trail
				db.Exec(`
					INSERT INTO embedding_audit_log (entity_type, entity_id, action, reason, triggered_by)
					VALUES (?, ?, 'failed', ?, 'regenerate-embeddings worker')
				`, item.EntityType, item.EntityID, err.Error())
			} else {
				fmt.Printf("  ⚠️  Failed (will retry): %v\n", err)
				db.Exec(`
					UPDATE embedding_queue
					SET status = 'pending', error_message = ?, retry_count = ?, processed_at = NULL
					WHERE id = ?
				`, err.Error(), newRetryCount, item.ID)
			}
		} else {
			successCount++
			fmt.Println("  ✅ Regenerated successfully")

			// Mark as completed
			db.Exec(`
				UPDATE embedding_queue
				SET status = 'completed', processed_at = NOW()
				WHERE id = ?
			`, item.ID)

			// Log to audit trail
			db.Exec(`
				INSERT INTO embedding_audit_log (entity_type, entity_id, action, triggered_by)
				VALUES (?, ?, 'regenerated', 'regenerate-embeddings worker')
			`, item.EntityType, item.EntityID)
		}

		// Rate limiting: wait 200ms between items to avoid API throttling
		if i < len(queueItems)-1 {
			time.Sleep(200 * time.Millisecond)
		}
	}

	fmt.Printf("\n✅ Completed: %d/%d successful, %d failed\n", successCount, len(queueItems), failCount)
}

// regenerateScoreItemEmbedding regenerates embedding for a single ScoreItem
func regenerateScoreItemEmbedding(
	db *gorm.DB,
	embeddingService *services.EmbeddingService,
	chunkingService *services.ChunkingService,
	scoreItemID uuid.UUID,
) error {
	// Fetch ScoreItem
	var item models.ScoreItem
	if err := db.First(&item, scoreItemID).Error; err != nil {
		return fmt.Errorf("failed to fetch ScoreItem: %w", err)
	}

	// Get full name (with relationships)
	db.Preload("Subgroup.Group").Preload("ParentItem").First(&item, scoreItemID)
	fullName := item.GetFullName()

	// Generate text for embedding
	textSource := chunkingService.ChunkScoreItem(
		fullName,
		item.ClinicalRelevance,
		item.PatientExplanation,
		item.Conduct,
		item.Gender,
		item.AgeRangeMin,
		item.AgeRangeMax,
		item.PostMenopause,
	)

	// Generate embedding
	ctx := context.Background()
	embedding, err := embeddingService.GenerateEmbedding(ctx, textSource)
	if err != nil {
		return fmt.Errorf("failed to generate embedding: %w", err)
	}

	// Update embedding in database
	return db.Transaction(func(tx *gorm.DB) error {
		// Check if embedding already exists
		var existing models.ScoreItemEmbedding
		err := tx.Where("score_item_id = ?", scoreItemID).First(&existing).Error

		if err == nil {
			// Update existing
			existing.TextSource = textSource
			existing.IsStale = false
			if err := existing.SetEmbeddingFromSlice(embedding); err != nil {
				return fmt.Errorf("failed to set embedding: %w", err)
			}
			if err := tx.Save(&existing).Error; err != nil {
				return fmt.Errorf("failed to update embedding: %w", err)
			}
		} else if err == gorm.ErrRecordNotFound {
			// Create new
			newEmbed := &models.ScoreItemEmbedding{
				ScoreItemID: scoreItemID,
				TextSource:  textSource,
				IsStale:     false,
			}
			if err := newEmbed.SetEmbeddingFromSlice(embedding); err != nil {
				return fmt.Errorf("failed to set embedding: %w", err)
			}
			if err := tx.Create(newEmbed).Error; err != nil {
				return fmt.Errorf("failed to create embedding: %w", err)
			}
		} else {
			return fmt.Errorf("failed to check existing embedding: %w", err)
		}

		return nil
	})
}

// regenerateArticleEmbedding regenerates embeddings for all chunks of an Article
func regenerateArticleEmbedding(
	db *gorm.DB,
	embeddingService *services.EmbeddingService,
	chunkingService *services.ChunkingService,
	articleID uuid.UUID,
) error {
	// Fetch Article
	var article models.Article
	if err := db.First(&article, articleID).Error; err != nil {
		return fmt.Errorf("failed to fetch Article: %w", err)
	}

	// Get fullContent and abstract
	fullContent := ""
	if article.FullContent != nil {
		fullContent = *article.FullContent
	}

	abstract := ""
	if article.Abstract != nil {
		abstract = *article.Abstract
	}

	// Check if article has content
	if fullContent == "" && abstract == "" {
		return fmt.Errorf("article has no content")
	}

	// Generate chunks
	chunks, err := chunkingService.ChunkArticle(fullContent, abstract)
	if err != nil {
		return fmt.Errorf("failed to chunk article: %w", err)
	}
	if len(chunks) == 0 {
		return fmt.Errorf("no chunks generated")
	}

	// Generate embeddings for all chunks
	ctx := context.Background()
	return db.Transaction(func(tx *gorm.DB) error {
		// Delete old embeddings
		if err := tx.Where("article_id = ?", articleID).Delete(&models.ArticleEmbedding{}).Error; err != nil {
			return fmt.Errorf("failed to delete old embeddings: %w", err)
		}

		// Insert new embeddings
		for _, chunk := range chunks {
			embeddingVals, err := embeddingService.GenerateEmbedding(ctx, chunk.Text)
			if err != nil {
				return fmt.Errorf("failed to generate embedding for chunk %d: %w", chunk.Index, err)
			}

			newEmbed := &models.ArticleEmbedding{
				ArticleID:     articleID,
				ChunkIndex:    chunk.Index,
				ChunkText:     chunk.Text,
				ChunkMetadata: chunk.Metadata,
				IsStale:       false,
			}

			if err := newEmbed.SetEmbeddingFromSlice(embeddingVals); err != nil {
				return fmt.Errorf("failed to set embedding for chunk %d: %w", chunk.Index, err)
			}

			if err := tx.Create(newEmbed).Error; err != nil {
				return fmt.Errorf("failed to save embedding for chunk %d: %w", chunk.Index, err)
			}

			// Rate limiting between chunks
			time.Sleep(100 * time.Millisecond)
		}

		return nil
	})
}
