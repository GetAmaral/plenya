package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ProcessingState tracks the state of batch processing
type ProcessingState struct {
	RunID          uuid.UUID
	TotalProcessed int
	TotalLinked    int
	TotalSkipped   int
	FailedItems    []uuid.UUID
}

// BatchResult tracks results for a single batch
type BatchResult struct {
	ItemsProcessed int
	LinksCreated   int
	ItemsSkipped   int
	ItemsFailed    int
	TotalConfidence float64
	ProcessingTime time.Duration
}

const BATCH_SIZE = 50

// Auto-Link ALL - Linkador automático com batch processing e rollback
func main() {
	godotenv.Load()
	cfg, _ := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	vectorRepo := repository.NewArticleVectorRepository(db)

	var scoreItems []models.ScoreItem
	db.Preload("Subgroup.Group").Preload("ParentItem").Find(&scoreItems)

	fmt.Printf("🔗 Auto-linking %d score items (batch size: %d)...\n", len(scoreItems), BATCH_SIZE)

	// Create processing run
	var runID uuid.UUID
	db.Raw("SELECT create_auto_link_run(?, ?)", len(scoreItems), BATCH_SIZE).Scan(&runID)
	fmt.Printf("📋 Run ID: %s\n\n", runID)

	state := &ProcessingState{RunID: runID}

	// Process in batches
	batchNumber := 0
	for i := 0; i < len(scoreItems); i += BATCH_SIZE {
		batchNumber++
		end := min(i+BATCH_SIZE, len(scoreItems))
		batch := scoreItems[i:end]

		fmt.Printf("📦 Batch %d: Processing items %d-%d...\n", batchNumber, i+1, end)

		result := processBatch(db, vectorRepo, batch, state)

		// Save checkpoint
		avgConfidence := 0.0
		if result.LinksCreated > 0 {
			avgConfidence = result.TotalConfidence / float64(result.LinksCreated)
		}

		db.Exec(`
			SELECT save_batch_checkpoint(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
			runID,
			batchNumber,
			i,
			end-1,
			result.ItemsProcessed,
			result.LinksCreated,
			result.ItemsSkipped,
			result.ItemsFailed,
			avgConfidence,
			result.ProcessingTime.Milliseconds(),
			batch[len(batch)-1].ID,
		)

		fmt.Printf("  ✓ Completed: %d processed, %d linked, %d skipped, %d failed (%.2fs)\n\n",
			result.ItemsProcessed,
			result.LinksCreated,
			result.ItemsSkipped,
			result.ItemsFailed,
			result.ProcessingTime.Seconds(),
		)
	}

	// Retry failed items (max 3 attempts)
	if len(state.FailedItems) > 0 {
		fmt.Printf("🔄 Retrying %d failed items...\n", len(state.FailedItems))
		retryFailed(db, vectorRepo, state, scoreItems)
	}

	// Complete run
	db.Exec("SELECT complete_auto_link_run(?, 'completed', NULL)", runID)

	fmt.Printf("\n✅ Auto-link completed!\n")
	fmt.Printf("   Total processed: %d\n", state.TotalProcessed)
	fmt.Printf("   Total linked: %d\n", state.TotalLinked)
	fmt.Printf("   Total skipped: %d\n", state.TotalSkipped)
	fmt.Printf("   Total failed: %d\n", len(state.FailedItems))
}

// processBatch processes a single batch of ScoreItems transactionally
func processBatch(
	db *gorm.DB,
	vectorRepo *repository.ArticleVectorRepository,
	batch []models.ScoreItem,
	state *ProcessingState,
) BatchResult {
	startTime := time.Now()
	result := BatchResult{}

	// Process each item in the batch (non-transactional to allow partial success)
	for _, item := range batch {
		itemStartTime := time.Now()
		linksCreated, avgConf, err := processItem(db, vectorRepo, &item)

		processingTime := time.Since(itemStartTime)

		if err != nil {
			result.ItemsFailed++
			state.FailedItems = append(state.FailedItems, item.ID)

			// Log failure
			db.Exec(`
				INSERT INTO auto_link_item_log
				(run_id, score_item_id, score_item_name, status, error_message, processing_time_ms)
				VALUES (?, ?, ?, 'failed', ?, ?)
			`, state.RunID, item.ID, item.Name, err.Error(), processingTime.Milliseconds())

			log.Printf("  ⚠️  Item %s failed: %v", item.ID, err)
		} else {
			result.ItemsProcessed++
			state.TotalProcessed++

			if linksCreated > 0 {
				result.LinksCreated += linksCreated
				result.TotalConfidence += avgConf * float64(linksCreated)
				state.TotalLinked += linksCreated

				// Log success
				db.Exec(`
					INSERT INTO auto_link_item_log
					(run_id, score_item_id, score_item_name, status, links_created, avg_confidence, processing_time_ms)
					VALUES (?, ?, ?, 'success', ?, ?, ?)
				`, state.RunID, item.ID, item.Name, linksCreated, avgConf, processingTime.Milliseconds())
			} else {
				result.ItemsSkipped++
				state.TotalSkipped++

				// Log skipped
				db.Exec(`
					INSERT INTO auto_link_item_log
					(run_id, score_item_id, score_item_name, status, processing_time_ms)
					VALUES (?, ?, ?, 'skipped', ?)
				`, state.RunID, item.ID, item.Name, processingTime.Milliseconds())
			}
		}
	}

	result.ProcessingTime = time.Since(startTime)
	return result
}

// processItem processes a single ScoreItem and returns (linksCreated, avgConfidence, error)
func processItem(
	db *gorm.DB,
	vectorRepo *repository.ArticleVectorRepository,
	item *models.ScoreItem,
) (int, float64, error) {
	// Verificar se já tem artigos linkados
	var existingCount int64
	db.Model(&models.Article{}).
		Joins("JOIN article_score_items ON articles.id = article_score_items.article_id").
		Where("article_score_items.score_item_id = ?", item.ID).
		Count(&existingCount)

	// Determinar threshold baseado em características do item
	baseThreshold := determineThreshold(item)
	threshold := baseThreshold
	chunkLimit := 30
	minChunksRequired := 2
	minSimilarityForSingle := 0.85

	if existingCount == 0 {
		// Cold-start: redução moderada do threshold (balanço entre qualidade e cobertura)
		threshold = baseThreshold - 0.10       // Ex: 0.70 → 0.60 (mais permissivo que 0.55)
		chunkLimit = 40                        // Buscar mais chunks para compensar
		minChunksRequired = 2                  // Exigir pelo menos 2 chunks
		minSimilarityForSingle = baseThreshold - 0.05 // Ex: 0.70 → 0.65 (mais permissivo)
	}

	chunks, err := vectorRepo.FindTopChunksForScoreItem(item.ID, chunkLimit, threshold)
	if err != nil {
		return 0, 0, err
	}

	if len(chunks) == 0 {
		return 0, 0, nil // No error, just no links created
	}

	// Agrupar por artigo
	articleChunks := make(map[uuid.UUID][]repository.ChunkSearchResult)
	for _, chunk := range chunks {
		articleChunks[chunk.ArticleID] = append(articleChunks[chunk.ArticleID], chunk)
	}

	// Filtrar qualificados
	linksCreated := 0
	var totalConfidence float64

	for articleID, chunkList := range articleChunks {
		maxSim := 0.0
		var totalSim float64
		for _, c := range chunkList {
			totalSim += c.Similarity
			if c.Similarity > maxSim {
				maxSim = c.Similarity
			}
		}

		if len(chunkList) >= minChunksRequired || maxSim >= minSimilarityForSingle {
			avgSim := totalSim / float64(len(chunkList))

			result := db.Exec(`
				INSERT INTO article_score_items
				(score_item_id, article_id, confidence_score, auto_linked, linked_at)
				VALUES (?, ?, ?, true, NOW())
				ON CONFLICT (score_item_id, article_id) DO NOTHING
			`, item.ID, articleID, avgSim)

			if result.RowsAffected > 0 {
				linksCreated++
				totalConfidence += avgSim
			}
		}
	}

	avgConfidence := 0.0
	if linksCreated > 0 {
		avgConfidence = totalConfidence / float64(linksCreated)
	}

	return linksCreated, avgConfidence, nil
}

// retryFailed retries failed items up to 3 times
func retryFailed(
	db *gorm.DB,
	vectorRepo *repository.ArticleVectorRepository,
	state *ProcessingState,
	allItems []models.ScoreItem,
) {
	maxRetries := 3
	itemMap := make(map[uuid.UUID]*models.ScoreItem)
	for i := range allItems {
		itemMap[allItems[i].ID] = &allItems[i]
	}

	for retry := 1; retry <= maxRetries; retry++ {
		if len(state.FailedItems) == 0 {
			break
		}

		fmt.Printf("  Retry %d/%d: %d items\n", retry, maxRetries, len(state.FailedItems))

		stillFailing := []uuid.UUID{}

		for _, itemID := range state.FailedItems {
			item := itemMap[itemID]
			if item == nil {
				continue
			}

			linksCreated, avgConf, err := processItem(db, vectorRepo, item)

			if err != nil {
				stillFailing = append(stillFailing, itemID)
				log.Printf("    ⚠️  Item %s still failing: %v", itemID, err)
			} else {
				state.TotalProcessed++
				if linksCreated > 0 {
					state.TotalLinked += linksCreated
					fmt.Printf("    ✓ Item %s recovered: %d links (avg: %.3f)\n", itemID, linksCreated, avgConf)
				}
			}
		}

		state.FailedItems = stillFailing

		if len(stillFailing) > 0 && retry < maxRetries {
			time.Sleep(2 * time.Second) // Wait before retry
		}
	}

	if len(state.FailedItems) > 0 {
		fmt.Printf("  ❌ %d items failed after %d retries\n", len(state.FailedItems), maxRetries)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// determineThreshold calculates optimal threshold based on ScoreItem characteristics
// Higher threshold = more conservative (fewer but higher quality links)
// Lower threshold = more permissive (more links but may include false positives)
func determineThreshold(item *models.ScoreItem) float64 {
	nameLen := len(item.Name)
	hasUnit := item.Unit != nil && *item.Unit != ""
	hasCR := item.ClinicalRelevance != nil && len(*item.ClinicalRelevance) > 200
	hasCRLong := item.ClinicalRelevance != nil && len(*item.ClinicalRelevance) > 500
	hasConduct := item.Conduct != nil && len(*item.Conduct) > 100
	hasDemographics := (item.Gender != nil && *item.Gender != "not_applicable") ||
		item.AgeRangeMin != nil ||
		item.AgeRangeMax != nil ||
		(item.PostMenopause != nil && *item.PostMenopause)

	// Score based on specificity indicators
	specificityScore := 0

	// Name characteristics (shorter + with unit = more specific)
	if nameLen <= 30 && hasUnit {
		specificityScore += 3
	} else if nameLen <= 50 && hasUnit {
		specificityScore += 2
	} else if nameLen > 80 {
		specificityScore -= 1 // Long descriptive names are less specific
	}

	// Clinical content richness (rich content = easier to match accurately)
	if hasCRLong {
		specificityScore += 2
	} else if hasCR {
		specificityScore += 1
	}

	if hasConduct {
		specificityScore += 1
	}

	// Demographic filters (more specific targeting)
	if hasDemographics {
		specificityScore += 2
	}

	// Convert specificity score to threshold
	// High specificity (6-8) → high threshold (0.75-0.80)
	// Medium specificity (3-5) → medium threshold (0.65-0.75)
	// Low specificity (0-2) → low threshold (0.60-0.65)
	switch {
	case specificityScore >= 7:
		return 0.80 // Very specific item, demand high similarity
	case specificityScore >= 5:
		return 0.75 // Specific item
	case specificityScore >= 3:
		return 0.70 // Moderately specific
	case specificityScore >= 1:
		return 0.65 // Generic item
	default:
		return 0.60 // Very generic, be more permissive
	}
}
