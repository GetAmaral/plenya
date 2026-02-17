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

// Auto-Link ALL - Linkador automático completo (sem interação)
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

	fmt.Printf("🔗 Auto-linking %d score items...\n", len(scoreItems))

	total := 0
	for i, item := range scoreItems {
		threshold := determineThreshold(&item)

		chunks, err := vectorRepo.FindTopChunksForScoreItem(item.ID, 30, threshold)
		if err != nil || len(chunks) == 0 {
			continue
		}

		// Agrupar por artigo
		articleChunks := make(map[uuid.UUID][]repository.ChunkSearchResult)
		for _, chunk := range chunks {
			articleChunks[chunk.ArticleID] = append(articleChunks[chunk.ArticleID], chunk)
		}

		// Filtrar qualificados
		for articleID, chunkList := range articleChunks {
			maxSim := 0.0
			var totalSim float64
			for _, c := range chunkList {
				totalSim += c.Similarity
				if c.Similarity > maxSim {
					maxSim = c.Similarity
				}
			}

			if len(chunkList) >= 2 || maxSim >= 0.85 {
				avgSim := totalSim / float64(len(chunkList))
				db.Exec(`
					INSERT INTO article_score_items
					(score_item_id, article_id, confidence_score, auto_linked, linked_at)
					VALUES (?, ?, ?, true, NOW())
					ON CONFLICT (score_item_id, article_id) DO NOTHING
				`, item.ID, articleID, avgSim)
				total++
			}
		}

		if (i+1)%100 == 0 {
			fmt.Printf("  Processed %d/%d...\n", i+1, len(scoreItems))
		}
	}

	fmt.Printf("✅ Created %d links\n", total)
}

func determineThreshold(item *models.ScoreItem) float64 {
	nameLen := len(item.Name)
	hasUnit := item.Unit != nil && *item.Unit != ""
	hasCR := item.ClinicalRelevance != nil && len(*item.ClinicalRelevance) > 500

	if nameLen <= 30 && hasUnit && hasCR {
		return 0.75
	}
	if nameLen > 50 || !hasUnit {
		return 0.65
	}
	return 0.70
}
