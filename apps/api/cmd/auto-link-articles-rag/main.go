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

// Auto-Link Articles Script - Conecta artigos a ScoreItems via similaridade semântica (RAG)
// Usa FindSimilarArticlesForScoreItem() com threshold 0.7 para auto-linking
//
// Uso: go run cmd/auto-link-articles-rag/main.go [--threshold 0.7] [--limit 30] [--dry-run]

func main() {
	fmt.Println("🔗 Auto-Link Articles via RAG - Plenya System")
	fmt.Println("=" + string(make([]byte, 60)))

	// Configurar threshold e limite
	threshold := 0.7  // Mínimo de similaridade (cosine similarity)
	limit := 30       // Máximo de artigos por ScoreItem
	dryRun := false   // Se true, apenas simula sem criar links

	// Parse flags simples (opcional)
	// TODO: Adicionar flag parsing se necessário

	// Carregar .env
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️  .env file not found (using environment variables)")
	}

	// Carregar configuração
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	// Conectar ao banco
	dsn := cfg.Database.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	fmt.Println("✅ Connected to database\n")

	// Criar repository vetorial
	vectorRepo := repository.NewArticleVectorRepository(db)

	// Buscar todos ScoreItems ativos
	var scoreItems []models.ScoreItem
	if err := db.Find(&scoreItems).Error; err != nil {
		log.Fatalf("❌ Failed to fetch score items: %v", err)
	}

	fmt.Printf("📊 Found %d score items to process\n", len(scoreItems))
	fmt.Printf("⚙️  Similarity threshold: %.2f\n", threshold)
	fmt.Printf("⚙️  Max articles per item: %d\n", limit)
	if dryRun {
		fmt.Println("🧪 DRY RUN MODE - No links will be created")
	}
	fmt.Println()

	// Estatísticas
	totalLinksCreated := 0
	itemsWithNewLinks := 0
	itemsSkipped := 0
	itemsWithErrors := 0

	// Processar cada ScoreItem
	for i, item := range scoreItems {
		fmt.Printf("[%d/%d] Processing: %s\n", i+1, len(scoreItems), item.Name)

		// 1. Buscar artigos similares via RAG
		results, err := vectorRepo.FindSimilarArticlesForScoreItem(
			item.ID,
			limit,
			threshold,
		)

		if err != nil {
			fmt.Printf("   ❌ Error finding similar articles: %v\n", err)
			itemsWithErrors++
			continue
		}

		if len(results) == 0 {
			fmt.Printf("   ⚠️  No similar articles found (threshold=%.2f)\n", threshold)
			itemsSkipped++
			continue
		}

		fmt.Printf("   🔍 Found %d similar articles (similarity ≥ %.2f)\n", len(results), threshold)

		// 2. Buscar artigos JÁ linkados a este ScoreItem
		var existingLinks []struct {
			ArticleID uuid.UUID
		}
		err = db.Raw(`
			SELECT article_id
			FROM article_score_items
			WHERE score_item_id = ?
		`, item.ID).Scan(&existingLinks).Error

		if err != nil {
			fmt.Printf("   ❌ Error checking existing links: %v\n", err)
			itemsWithErrors++
			continue
		}

		existingIDs := make(map[uuid.UUID]bool)
		for _, link := range existingLinks {
			existingIDs[link.ArticleID] = true
		}

		// 3. Filtrar artigos novos (não linkados)
		newArticles := []repository.ArticleSearchResult{}
		for _, result := range results {
			if !existingIDs[result.Article.ID] {
				newArticles = append(newArticles, result)
			}
		}

		if len(newArticles) == 0 {
			fmt.Printf("   ✓ All similar articles already linked (%d existing)\n", len(existingLinks))
			itemsSkipped++
			continue
		}

		fmt.Printf("   📎 Will create %d new links (%d already exist)\n", len(newArticles), len(existingLinks))

		// 4. Criar links (ou apenas simular se dry-run)
		if !dryRun {
			linksCreated := 0
			for _, result := range newArticles {
				// Insert direto SQL (tabela join GORM)
				err := db.Exec(`
					INSERT INTO article_score_items
					(score_item_id, article_id, confidence_score, auto_linked, linked_at, linked_by)
					VALUES (?, ?, ?, true, NOW(), NULL)
					ON CONFLICT (score_item_id, article_id) DO NOTHING
				`, item.ID, result.Article.ID, result.Similarity).Error

				if err != nil {
					fmt.Printf("   ⚠️  Failed to link article %s: %v\n",
						result.Article.ID.String()[:8], err)
					continue
				}

				linksCreated++
			}

			if linksCreated > 0 {
				fmt.Printf("   ✅ Created %d links\n", linksCreated)
				totalLinksCreated += linksCreated
				itemsWithNewLinks++
			}
		} else {
			// Dry run - apenas mostrar
			for _, result := range newArticles {
				fmt.Printf("      [DRY RUN] Would link: %s (similarity: %.3f)\n",
					result.Article.Title[:60], result.Similarity)
			}
			totalLinksCreated += len(newArticles)
			itemsWithNewLinks++
		}

		fmt.Println()
	}

	// Relatório final
	fmt.Println("=" + string(make([]byte, 60)))
	fmt.Println("📊 Auto-Linking Summary:")
	fmt.Printf("   Total ScoreItems processed: %d\n", len(scoreItems))
	fmt.Printf("   Items with new links: %d\n", itemsWithNewLinks)
	fmt.Printf("   Items skipped (no new articles): %d\n", itemsSkipped)
	fmt.Printf("   Items with errors: %d\n", itemsWithErrors)
	fmt.Printf("   Total links created: %d\n", totalLinksCreated)
	if dryRun {
		fmt.Println("   ⚠️  DRY RUN - No changes were saved")
	}
	fmt.Println()

	// Estatísticas finais do banco
	if !dryRun {
		var stats struct {
			TotalLinks     int64
			AutoLinked     int64
			ManualLinked   int64
			AvgConfidence  float64
			ItemsWithLinks int64
		}

		db.Raw(`
			SELECT
				COUNT(*) as total_links,
				COUNT(*) FILTER (WHERE auto_linked = true) as auto_linked,
				COUNT(*) FILTER (WHERE auto_linked = false) as manual_linked,
				AVG(confidence_score) as avg_confidence,
				COUNT(DISTINCT score_item_id) as items_with_links
			FROM article_score_items
		`).Scan(&stats)

		fmt.Println("📈 Database Statistics:")
		fmt.Printf("   Total article-item links: %d\n", stats.TotalLinks)
		fmt.Printf("   Auto-linked: %d (%.1f%%)\n",
			stats.AutoLinked,
			float64(stats.AutoLinked)/float64(stats.TotalLinks)*100)
		fmt.Printf("   Manual-linked: %d (%.1f%%)\n",
			stats.ManualLinked,
			float64(stats.ManualLinked)/float64(stats.TotalLinks)*100)
		fmt.Printf("   Average confidence: %.3f\n", stats.AvgConfidence)
		fmt.Printf("   ScoreItems with articles: %d/%d (%.1f%%)\n",
			stats.ItemsWithLinks,
			len(scoreItems),
			float64(stats.ItemsWithLinks)/float64(len(scoreItems))*100)
		fmt.Println()
	}

	fmt.Println("✅ Auto-linking completed!")
	if !dryRun {
		fmt.Println("\n💡 Next steps:")
		fmt.Println("   1. Review coverage: SELECT COUNT(*) FROM article_score_items WHERE auto_linked = true;")
		fmt.Println("   2. Check items with <7 articles (may need PubMed search)")
		fmt.Println("   3. Validate similarity scores are reasonable (avg ≥ 0.75)")
	}
}
