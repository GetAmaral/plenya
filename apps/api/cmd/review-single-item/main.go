package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
	"github.com/plenya/api/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Review Single ScoreItem - Revisão completa de um ScoreItem específico
// Inclui busca RAG, PubMed e enriquecimento de campos clínicos
//
// Uso: go run cmd/review-single-item/main.go <score-item-id>

func main() {
	if len(os.Args) < 2 {
		log.Fatal("❌ Usage: go run cmd/review-single-item/main.go <score-item-id>")
	}

	scoreItemID := os.Args[1]
	itemUUID, err := uuid.Parse(scoreItemID)
	if err != nil {
		log.Fatalf("❌ Invalid UUID: %v", err)
	}

	fmt.Printf("🔍 Reviewing ScoreItem: %s\n", scoreItemID)
	fmt.Println("=" + string(make([]byte, 80)))

	// Carregar .env
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️  .env file not found (using environment variables)")
	}

	// Verificar API key
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ ANTHROPIC_API_KEY environment variable is required")
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

	// Buscar ScoreItem
	var item models.ScoreItem
	if err := db.First(&item, "id = ?", itemUUID).Error; err != nil {
		log.Fatalf("❌ ScoreItem not found: %v", err)
	}

	fmt.Printf("📋 ScoreItem: %s\n", item.Name)
	fmt.Printf("   Gender: %s\n", item.Gender)
	fmt.Printf("   Unit: %s\n", item.Unit)
	fmt.Println()

	// ========================================
	// ETAPA 1: Revisar Artigos (RAG + PubMed)
	// ========================================
	fmt.Println("📚 STEP 1: Article Review")
	fmt.Println("-" + string(make([]byte, 79)))

	ctx := context.Background()

	// Contar artigos atuais
	var currentArticleCount int64
	db.Raw(`
		SELECT COUNT(*)
		FROM article_score_items
		WHERE score_item_id = ?
	`, itemUUID).Scan(&currentArticleCount)

	fmt.Printf("Current articles linked: %d\n", currentArticleCount)

	// Buscar artigos via RAG
	vectorRepo := repository.NewArticleVectorRepository(db)
	ragResults, err := vectorRepo.FindSimilarArticlesForScoreItem(
		itemUUID,
		30,
		0.7,
	)

	if err != nil {
		log.Printf("⚠️  RAG search error: %v", err)
	} else {
		fmt.Printf("RAG search found: %d similar articles (threshold 0.7)\n", len(ragResults))

		// Filtrar artigos já linkados
		var existingLinks []struct {
			ArticleID uuid.UUID
		}
		db.Raw(`
			SELECT article_id
			FROM article_score_items
			WHERE score_item_id = ?
		`, itemUUID).Scan(&existingLinks)

		existingIDs := make(map[uuid.UUID]bool)
		for _, link := range existingLinks {
			existingIDs[link.ArticleID] = true
		}

		// Link novos artigos
		newLinksCreated := 0
		for _, result := range ragResults {
			if !existingIDs[result.Article.ID] {
				err := db.Exec(`
					INSERT INTO article_score_items
					(score_item_id, article_id, confidence_score, auto_linked, linked_at)
					VALUES (?, ?, ?, true, NOW())
					ON CONFLICT DO NOTHING
				`, itemUUID, result.Article.ID, result.Similarity).Error

				if err == nil {
					newLinksCreated++
					fmt.Printf("   ✓ Linked: %s (similarity: %.3f)\n",
						truncate(result.Article.Title, 60), result.Similarity)
				}
			}
		}

		if newLinksCreated > 0 {
			fmt.Printf("✅ Added %d new articles via RAG\n", newLinksCreated)
			currentArticleCount += int64(newLinksCreated)
		}
	}

	// Buscar no PubMed se < 7 artigos
	minArticles := 7
	if currentArticleCount < int64(minArticles) {
		fmt.Printf("\n⚠️  Only %d articles (need %d), searching PubMed...\n",
			currentArticleCount, minArticles)

		pubmedService := services.NewPubMedService(db)
		query := "waist circumference men cardiovascular risk metabolic syndrome AND (Review[ptyp] OR Meta-Analysis[ptyp])"

		searchResults, err := pubmedService.SearchArticles(ctx, query, 20)
		if err != nil {
			log.Printf("⚠️  PubMed search error: %v", err)
		} else {
			fmt.Printf("PubMed search: %d results found\n", len(searchResults))

			if len(searchResults) > 0 {
				// Fetch article details
				pubmedArticles, err := pubmedService.FetchArticleDetails(ctx, searchResults)
				if err != nil {
					log.Printf("⚠️  Failed to fetch article details: %v", err)
				} else {
					pubmedAdded := 0
					for _, pmArticle := range pubmedArticles {
						if currentArticleCount >= int64(minArticles) {
							break
						}

						// Verificar se já existe
						var count int64
						db.Model(&models.Article{}).Where("pm_id = ?", pmArticle.PMID).Count(&count)
						if count > 0 {
							continue
						}

						// Criar artigo no banco (sem PDF por enquanto)
						article, err := pubmedService.CreateArticleFromPubMed(pmArticle, "")
						if err != nil {
							log.Printf("   ⚠️  Failed to save PMID %s: %v", pmArticle.PMID, err)
							continue
						}

						// Link ao ScoreItem
						db.Exec(`
							INSERT INTO article_score_items
							(score_item_id, article_id, auto_linked, linked_at)
							VALUES (?, ?, true, NOW())
						`, itemUUID, article.ID)

						pubmedAdded++
						currentArticleCount++

						if pmArticle.Title != "" {
							fmt.Printf("   ✓ Added PMID %s: %s\n", pmArticle.PMID, truncate(pmArticle.Title, 60))
						} else {
							fmt.Printf("   ✓ Added PMID %s\n", pmArticle.PMID)
						}

						// Download PDF se tiver DOI
						if pmArticle.DOI != "" {
							pdfPath, err := pubmedService.DownloadPDF(ctx, pmArticle.DOI, article.ID)
							if err == nil {
								// Atualizar internal_link
								article.InternalLink = &pdfPath
								db.Save(article)
								fmt.Printf("      ✓ PDF downloaded: %s\n", pdfPath)
							}
						}

						time.Sleep(500 * time.Millisecond) // Rate limiting
					}

					if pubmedAdded > 0 {
						fmt.Printf("✅ Added %d articles from PubMed\n", pubmedAdded)
					}
				}
			}
		}
	}

	fmt.Printf("\n📊 Final article count: %d\n", currentArticleCount)
	fmt.Println()

	// ========================================
	// ETAPA 2: Revisar Campos Clínicos
	// ========================================
	fmt.Println("🧠 STEP 2: Clinical Fields Review (LLM)")
	fmt.Println("-" + string(make([]byte, 79)))

	enrichmentService := services.NewScoreItemEnrichmentService(db)
	tier := enrichmentService.DetermineTier(&item)

	fmt.Printf("Enrichment tier: %s\n", tier)

	// Se já está excelente (Tier Preserve), perguntar se quer revisar mesmo assim
	if tier == services.TierPreserve {
		fmt.Println("⚠️  This item is already in Tier 1 (excellent quality)")
		fmt.Println("   Current fields are comprehensive and well-structured.")
		fmt.Println("   Enrichment will be skipped to preserve quality.")
		fmt.Println("\n✅ Review completed! No changes needed.")
		return
	}

	// Gerar enriquecimento
	startTime := time.Now()

	result, err := enrichmentService.GenerateEnrichment(ctx, &item, tier)
	if err != nil {
		log.Fatalf("❌ Enrichment failed: %v", err)
	}

	duration := time.Since(startTime)
	fmt.Printf("✓ Generated in %.1fs (confidence: %.2f)\n", duration.Seconds(), result.Confidence)

	// Validar resultado
	validationErrors := enrichmentService.ValidateResult(result)
	if len(validationErrors) > 0 {
		fmt.Println("\n⚠️  Validation errors:")
		for _, verr := range validationErrors {
			fmt.Printf("   - %s\n", verr)
		}
		log.Fatal("❌ Enrichment validation failed")
	}

	// Mostrar preview
	fmt.Println("\n📝 New content preview:")
	fmt.Printf("   ClinicalRelevance: %d chars\n", len(result.ClinicalRelevance))
	fmt.Printf("   PatientExplanation: %d chars\n", len(result.PatientExplanation))
	fmt.Printf("   Conduct: %d chars\n", len(result.Conduct))
	fmt.Printf("   MaxPoints: %d\n", result.MaxPoints)
	fmt.Println()

	// Salvar auditoria
	if err := enrichmentService.SaveReviewHistory(&item, result, tier); err != nil {
		log.Printf("⚠️  Failed to save audit: %v", err)
	} else {
		fmt.Println("✓ Audit trail saved")
	}

	// Atualizar item
	item.ClinicalRelevance = &result.ClinicalRelevance
	item.PatientExplanation = &result.PatientExplanation
	item.Conduct = &result.Conduct
	points := float64(result.MaxPoints)
	item.Points = &points

	if err := db.Save(&item).Error; err != nil {
		log.Fatalf("❌ Failed to save item: %v", err)
	}

	fmt.Println("✅ ScoreItem updated successfully")
	fmt.Println()

	// Relatório final
	fmt.Println("=" + string(make([]byte, 80)))
	fmt.Println("✅ REVIEW COMPLETED!")
	fmt.Println()
	fmt.Printf("📋 ScoreItem: %s\n", item.Name)
	fmt.Printf("   Articles: %d\n", currentArticleCount)
	fmt.Printf("   LastReview: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("   Tier: %s → Tier 1 (Preserve)\n", tier)
	fmt.Println()
	fmt.Println("💡 Next steps:")
	fmt.Println("   1. Review changes in database")
	fmt.Println("   2. Check audit trail: SELECT * FROM score_item_review_history WHERE score_item_id = '" + scoreItemID + "';")
	fmt.Println("   3. Validate quality manually")
	fmt.Println("   4. Rollback if needed using before_snapshot from audit")
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
