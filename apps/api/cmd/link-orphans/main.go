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

// Link Orphans - Criar links para ScoreItems que não têm nenhum link
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

	// Buscar ScoreItems SEM NENHUM link auto
	var orphanItems []models.ScoreItem
	db.Raw(`
		SELECT si.*
		FROM score_items si
		WHERE NOT EXISTS (
			SELECT 1 FROM article_score_items asi
			WHERE asi.score_item_id = si.id AND asi.auto_linked = true
		)
		AND si.deleted_at IS NULL
		ORDER BY si.name
	`).Scan(&orphanItems)

	fmt.Printf("🔗 Linking %d orphan ScoreItems (threshold: 0.15 - MUITO permissivo)...\n\n", len(orphanItems))

	totalLinked := 0
	totalWithNoMatches := 0

	for i, item := range orphanItems {
		// Threshold MUITO baixo para forçar matches
		threshold := 0.15
		chunkLimit := 50

		chunks, err := vectorRepo.FindTopChunksForScoreItem(item.ID, chunkLimit, threshold)
		if err != nil || len(chunks) == 0 {
			totalWithNoMatches++
			if i < 10 {
				fmt.Printf("  ⚠️  %s: NO MATCHES (err: %v)\n", item.Name, err)
			}
			continue
		}

		// Agrupar por artigo
		articleChunks := make(map[uuid.UUID][]repository.ChunkSearchResult)
		for _, chunk := range chunks {
			articleChunks[chunk.ArticleID] = append(articleChunks[chunk.ArticleID], chunk)
		}

		// Criar links (mínimo 1 chunk com sim >= 0.3)
		linksCreated := 0
		for articleID, chunkList := range articleChunks {
			maxSim := 0.0
			var totalSim float64
			for _, c := range chunkList {
				totalSim += c.Similarity
				if c.Similarity > maxSim {
					maxSim = c.Similarity
				}
			}

			// Aceitar se: >= 2 chunks OU 1 chunk com sim >= 0.3
			if len(chunkList) >= 2 || maxSim >= 0.3 {
				avgSim := totalSim / float64(len(chunkList))
				db.Exec(`
					INSERT INTO article_score_items
					(score_item_id, article_id, confidence_score, auto_linked, linked_at)
					VALUES (?, ?, ?, true, NOW())
					ON CONFLICT (score_item_id, article_id) DO NOTHING
				`, item.ID, articleID, avgSim)
				linksCreated++
			}
		}

		if linksCreated > 0 {
			totalLinked++
			if i < 20 {
				fmt.Printf("  ✅ %s: %d links created\n", item.Name, linksCreated)
			}
		}

		if (i+1)%50 == 0 {
			fmt.Printf("\n  Progress: %d/%d...\n\n", i+1, len(orphanItems))
		}
	}

	fmt.Printf("\n✅ Completed:\n")
	fmt.Printf("   ScoreItems linkados: %d/%d\n", totalLinked, len(orphanItems))
	fmt.Printf("   ScoreItems sem matches: %d\n", totalWithNoMatches)
}
