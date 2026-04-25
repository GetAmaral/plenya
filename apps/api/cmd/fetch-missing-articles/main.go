//go:build legacy_scripts
// +build legacy_scripts

// Script ad-hoc — pra rodar use `go run -tags legacy_scripts ./cmd/fetch-missing-articles`
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

// Fetch Missing Articles - Busca artigos no PubMed para ScoreItems com <7 artigos
// Integra PubMed API + Unpaywall para download de PDFs
//
// Uso: go run cmd/fetch-missing-articles/main.go [--threshold 7] [--max-per-item 5] [--dry-run]

func main() {
	fmt.Println("🔍 Fetch Missing Articles - PubMed Integration")
	fmt.Println("=" + string(make([]byte, 60)))

	// Configuração
	threshold := 7     // Mínimo de artigos desejado
	maxPerItem := 5    // Máximo de artigos novos a buscar por item
	dryRun := false    // Se true, apenas simula

	// Carregar .env
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️  .env file not found (using environment variables)")
	}

	// Verificar variáveis necessárias
	email := godotenv.Get("PUBMED_EMAIL")
	if email == "" {
		log.Fatal("❌ PUBMED_EMAIL environment variable is required")
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

	fmt.Println("✅ Connected to database")
	fmt.Printf("📧 PubMed email: %s\n", email)
	fmt.Printf("⚙️  Threshold: %d articles per item\n", threshold)
	fmt.Printf("⚙️  Max new articles per item: %d\n", maxPerItem)
	if dryRun {
		fmt.Println("🧪 DRY RUN MODE - No changes will be made")
	}
	fmt.Println()

	// Criar serviços
	pubmedService := services.NewPubMedService(db)
	queueService := services.NewEmbeddingQueueService(db)

	// Buscar ScoreItems que precisam de artigos
	type ItemWithCount struct {
		ID           string
		Name         string
		Unit         *string
		ArticleCount int64
	}

	var itemsNeedingArticles []ItemWithCount
	err = db.Raw(`
		SELECT
			si.id,
			si.name,
			si.unit,
			COUNT(asi.article_id) as article_count
		FROM score_items si
		LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
		GROUP BY si.id, si.name, si.unit
		HAVING COUNT(asi.article_id) < ?
		ORDER BY article_count ASC, name
	`, threshold).Scan(&itemsNeedingArticles).Error

	if err != nil {
		log.Fatalf("❌ Failed to query items: %v", err)
	}

	fmt.Printf("📊 Found %d items with <%d articles\n", len(itemsNeedingArticles), threshold)

	// Filtrar apenas items laboratoriais (com unidade)
	labItems := []ItemWithCount{}
	for _, item := range itemsNeedingArticles {
		if item.Unit != nil && *item.Unit != "" {
			labItems = append(labItems, item)
		}
	}

	fmt.Printf("🔬 %d are laboratory items (with unit)\n", len(labItems))

	if len(labItems) == 0 {
		fmt.Println("✅ All laboratory items have sufficient articles!")
		return
	}

	fmt.Println("\n🚀 Starting article search...\n")

	// Estatísticas
	totalArticlesFound := 0
	totalArticlesCreated := 0
	totalPDFsDownloaded := 0
	itemsProcessed := 0
	itemsWithErrors := 0

	ctx := context.Background()

	// Processar cada item
	for i, item := range labItems {
		fmt.Printf("[%d/%d] Processing: %s (%d articles)\n",
			i+1, len(labItems), item.Name, item.ArticleCount)

		// Buscar ScoreItem completo
		var scoreItem models.ScoreItem
		if err := db.First(&scoreItem, "id = ?", item.ID).Error; err != nil {
			fmt.Printf("   ❌ Error loading item: %v\n", err)
			itemsWithErrors++
			continue
		}

		// Gerar query PubMed
		query := pubmedService.GenerateQueryForScoreItem(&scoreItem)
		fmt.Printf("   🔍 Query: %s\n", query)

		// Buscar no PubMed (com rate limiting)
		pmids, err := pubmedService.RateLimitedSearch(ctx, query, maxPerItem)
		if err != nil {
			fmt.Printf("   ❌ PubMed search failed: %v\n", err)
			itemsWithErrors++
			continue
		}

		if len(pmids) == 0 {
			fmt.Printf("   ⚠️  No articles found\n\n")
			continue
		}

		fmt.Printf("   ✓ Found %d articles\n", len(pmids))
		totalArticlesFound += len(pmids)

		if dryRun {
			fmt.Printf("   [DRY RUN] Would fetch metadata for PMIDs: %v\n\n", pmids)
			continue
		}

		// Fetch metadata
		pubmedArticles, err := pubmedService.FetchArticleDetails(ctx, pmids)
		if err != nil {
			fmt.Printf("   ⚠️  Failed to fetch metadata: %v\n", err)
			itemsWithErrors++
			continue
		}

		// Processar cada artigo
		articlesCreated := 0
		pdfsDownloaded := 0

		for _, pmArticle := range pubmedArticles {
			// Tentar download PDF (opcional - pode falhar)
			var pdfPath string
			if pmArticle.DOI != "" {
				// Criar ID temporário para o artigo
				tempID := uuid.Must(uuid.NewV7())

				pdf, err := pubmedService.DownloadPDF(ctx, pmArticle.DOI, tempID)
				if err != nil {
					fmt.Printf("      ⚠️  PDF download failed (PMID %s): %v\n",
						pmArticle.PMID, err)
					// Continuar sem PDF
				} else {
					pdfPath = pdf
					pdfsDownloaded++
					fmt.Printf("      ✓ Downloaded PDF\n")
				}
			}

			// Criar artigo no banco
			article, err := pubmedService.CreateArticleFromPubMed(pmArticle, pdfPath)
			if err != nil {
				fmt.Printf("      ⚠️  Failed to create article: %v\n", err)
				continue
			}

			articlesCreated++

			// Linkar ao ScoreItem
			err = db.Exec(`
				INSERT INTO article_score_items
				(score_item_id, article_id, confidence_score, auto_linked, linked_at)
				VALUES (?, ?, 1.0, true, NOW())
				ON CONFLICT (score_item_id, article_id) DO NOTHING
			`, scoreItem.ID, article.ID).Error

			if err != nil {
				fmt.Printf("      ⚠️  Failed to link article: %v\n", err)
				continue
			}

			// Enfileirar para embedding
			if pdfPath != "" {
				if err := queueService.QueueArticle(article.ID); err != nil {
					fmt.Printf("      ⚠️  Failed to queue for embedding: %v\n", err)
				}
			}
		}

		fmt.Printf("   ✅ Created %d articles, downloaded %d PDFs\n\n",
			articlesCreated, pdfsDownloaded)

		totalArticlesCreated += articlesCreated
		totalPDFsDownloaded += pdfsDownloaded
		itemsProcessed++

		// Rate limiting entre items
		time.Sleep(500 * time.Millisecond)
	}

	// Relatório final
	fmt.Println("=" + string(make([]byte, 60)))
	fmt.Println("📊 Summary:")
	fmt.Printf("   Items processed: %d/%d\n", itemsProcessed, len(labItems))
	fmt.Printf("   Items with errors: %d\n", itemsWithErrors)
	fmt.Printf("   Articles found (PubMed): %d\n", totalArticlesFound)
	fmt.Printf("   Articles created (DB): %d\n", totalArticlesCreated)
	fmt.Printf("   PDFs downloaded: %d\n", totalPDFsDownloaded)

	if !dryRun && totalArticlesCreated > 0 {
		fmt.Printf("\n💡 Next steps:\n")
		fmt.Printf("   1. Wait for embedding worker to process %d articles\n", totalArticlesCreated)
		fmt.Printf("   2. Run auto-link-articles-rag again to connect via similarity\n")
		fmt.Printf("   3. Verify coverage improved\n")
	}

	fmt.Println("\n✅ Fetch completed!")
}
