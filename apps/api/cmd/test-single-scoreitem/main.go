package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	godotenv.Load()
	cfg, _ := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	vectorRepo := repository.NewArticleVectorRepository(db)

	// Test Albumina
	albuminaID := uuid.MustParse("c77cedd3-2800-7633-b793-4f5c9ab56b59")

	var item models.ScoreItem
	db.First(&item, albuminaID)

	fmt.Printf("Testing ScoreItem: %s (ID: %s)\n", item.Name, item.ID)

	// Verificar existing links
	var existingCount int64
	db.Model(&models.Article{}).
		Joins("JOIN article_score_items ON articles.id = article_score_items.article_id").
		Where("article_score_items.score_item_id = ?", item.ID).
		Count(&existingCount)

	fmt.Printf("Existing links: %d\n", existingCount)

	// Use cold-start threshold
	threshold := 0.2
	chunkLimit := 50

	fmt.Printf("Searching with threshold: %.2f, limit: %d\n", threshold, chunkLimit)

	chunks, err := vectorRepo.FindTopChunksForScoreItem(item.ID, chunkLimit, threshold)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("\n✅ Found %d chunks\n\n", len(chunks))

	if len(chunks) > 0 {
		fmt.Println("Top 5 chunks:")
		for i, chunk := range chunks {
			if i >= 5 {
				break
			}
			fmt.Printf("  %d. Article: %s\n", i+1, chunk.ArticleTitle)
			fmt.Printf("     Similarity: %.3f\n", chunk.Similarity)
			fmt.Printf("     Section: %s\n", chunk.Section)
			fmt.Printf("     Text preview: %.100s...\n\n", chunk.ChunkText)
		}
	} else {
		fmt.Println("❌ NO CHUNKS FOUND!")
	}
}
