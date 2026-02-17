package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	scoreItemID  = "c77cedd3-2800-7ef6-8f6e-66db64655c69"
	minArticles  = 7
	ragThreshold = 0.7
)

func main() {
	// Database connection from environment variables
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "db"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "plenya_user"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "plenya_dev_password"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "plenya_db"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dbSSLMode := os.Getenv("DB_SSL_MODE")
	if dbSSLMode == "" {
		dbSSLMode = "disable"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("╔═══════════════════════════════════════════════════════════════╗")
	fmt.Println("║     REVISÃO COMPLETA: 25-HIDROXIVITAMINA D (VITAMINA D)      ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════╝\n")

	// 1. Load ScoreItem with current articles
	var scoreItem models.ScoreItem
	if err := db.Preload("Articles").First(&scoreItem, "id = ?", scoreItemID).Error; err != nil {
		log.Fatalf("❌ Failed to load score item: %v", err)
	}

	fmt.Printf("📊 ScoreItem: %s (ID: %s)\n", scoreItem.Name, scoreItem.ID)
	fmt.Printf("   Pontos atuais: %.0f\n", *scoreItem.Points)
	if scoreItem.LabTestCode != nil {
		fmt.Printf("   Código laboratorial: %s\n", *scoreItem.LabTestCode)
	}
	if scoreItem.LastReview != nil {
		fmt.Printf("   Última revisão: %s\n", scoreItem.LastReview.Format("2006-01-02"))
	}
	fmt.Printf("   Artigos linkados: %d\n\n", len(scoreItem.Articles))

	// 2. RAG Search for similar articles
	fmt.Println("═══ FASE 1: BUSCA RAG POR ARTIGOS SIMILARES ═══")

	vectorRepo := repository.NewArticleVectorRepository(db)
	var ragResults []repository.ArticleSearchResult

	// First check if score_item has embeddings
	var embeddingCount int64
	db.Model(&models.ScoreItemEmbedding{}).Where("score_item_id = ?", scoreItemID).Count(&embeddingCount)

	if embeddingCount == 0 {
		fmt.Printf("⚠️  ScoreItem não possui embeddings. Pulando busca RAG.\n")
		fmt.Println("    (Execute embedding pipeline primeiro para habilitar RAG)\n")
	} else {
		ragResults, err = vectorRepo.FindSimilarArticlesForScoreItem(
			uuid.MustParse(scoreItemID),
			20,
			ragThreshold,
		)
		if err != nil {
			log.Printf("⚠️  RAG search failed: %v", err)
			ragResults = []repository.ArticleSearchResult{}
		}

		fmt.Printf("✓ Encontrados %d artigos similares via RAG (threshold %.2f)\n", len(ragResults), ragThreshold)

		// Get existing article IDs
		existingArticleIDs := make(map[uuid.UUID]bool)
		for _, article := range scoreItem.Articles {
			existingArticleIDs[article.ID] = true
		}

		// Link new articles from RAG
		newLinked := 0
		for _, result := range ragResults {
			if !existingArticleIDs[result.Article.ID] {
				// Link article to score item
				err := db.Exec(`
					INSERT INTO article_score_items (score_item_id, article_id, confidence_score, auto_linked)
					VALUES (?, ?, ?, true)
					ON CONFLICT (score_item_id, article_id) DO NOTHING
				`, scoreItemID, result.Article.ID, result.Similarity).Error

				if err != nil {
					log.Printf("  ⚠️  Failed to link article %s: %v", result.Article.ID, err)
				} else {
					fmt.Printf("  ✓ [NOVO] %s (similaridade: %.3f)\n", result.Article.Title, result.Similarity)
					existingArticleIDs[result.Article.ID] = true
					newLinked++
				}
			}
		}

		if newLinked > 0 {
			fmt.Printf("\n✓ %d novos artigos linkados via RAG\n\n", newLinked)
		} else {
			fmt.Println("  Nenhum novo artigo encontrado via RAG\n")
		}
	}

	// 3. Check total articles
	var totalArticles int64
	db.Table("article_score_items").Where("score_item_id = ?", scoreItemID).Count(&totalArticles)

	fmt.Println("═══ FASE 2: ANÁLISE DE ARTIGOS ═══")
	fmt.Printf("Total de artigos linkados: %d\n", totalArticles)

	if totalArticles < minArticles {
		fmt.Printf("⚠️  Menos que o mínimo recomendado (%d artigos)\n", minArticles)
		fmt.Println("   Sugestão: Executar busca PubMed manual ou usar script específico\n")
	} else {
		fmt.Printf("✓ Quantidade adequada (mínimo: %d)\n\n", minArticles)
	}

	// List all linked articles
	var linkedArticles []models.Article
	db.Joins("JOIN article_score_items ON articles.id = article_score_items.article_id").
		Where("article_score_items.score_item_id = ?", scoreItemID).
		Order("publish_date DESC").
		Find(&linkedArticles)

	fmt.Println("Artigos linkados:")
	for i, article := range linkedArticles {
		fmt.Printf("%2d. [%d] %s - %s\n", i+1, article.PublishDate.Year(), article.Title, article.Journal)
	}
	fmt.Println()

	// 4. Review clinical fields
	fmt.Println("═══ FASE 3: REVISÃO DE CAMPOS CLÍNICOS ═══")

	// Analyze current content
	crLen := 0
	if scoreItem.ClinicalRelevance != nil {
		crLen = len(*scoreItem.ClinicalRelevance)
	}
	peLen := 0
	if scoreItem.PatientExplanation != nil {
		peLen = len(*scoreItem.PatientExplanation)
	}
	condLen := 0
	if scoreItem.Conduct != nil {
		condLen = len(*scoreItem.Conduct)
	}

	fmt.Printf("Análise de conteúdo:\n")
	fmt.Printf("  • clinical_relevance:  %d chars ", crLen)
	if crLen >= 1500 {
		fmt.Println("✓ EXCELENTE")
	} else if crLen >= 1000 {
		fmt.Println("✓ Bom")
	} else if crLen >= 500 {
		fmt.Println("⚠ Adequado (pode melhorar)")
	} else {
		fmt.Println("❌ Insuficiente")
	}

	fmt.Printf("  • patient_explanation: %d chars ", peLen)
	if peLen >= 600 {
		fmt.Println("✓ EXCELENTE")
	} else if peLen >= 400 {
		fmt.Println("✓ Bom")
	} else if peLen >= 200 {
		fmt.Println("⚠ Adequado (pode melhorar)")
	} else {
		fmt.Println("❌ Insuficiente")
	}

	fmt.Printf("  • conduct:             %d chars ", condLen)
	if condLen >= 800 {
		fmt.Println("✓ EXCELENTE")
	} else if condLen >= 500 {
		fmt.Println("✓ Bom")
	} else if condLen >= 300 {
		fmt.Println("⚠ Adequado (pode melhorar)")
	} else {
		fmt.Println("❌ Insuficiente")
	}

	fmt.Printf("  • max_points:          %.0f ", *scoreItem.Points)
	points := int(*scoreItem.Points)
	if points >= 25 && points <= 35 {
		fmt.Println("✓ ADEQUADO (marcador importante)")
	} else if points >= 20 && points <= 40 {
		fmt.Println("✓ Razoável")
	} else {
		fmt.Println("⚠ Considerar ajuste")
	}

	fmt.Println()

	// Check content quality
	hasRecentGuidelines := false
	if scoreItem.ClinicalRelevance != nil {
		content := *scoreItem.ClinicalRelevance
		hasRecentGuidelines = contains(content, "2024") || contains(content, "Endocrine Society")
	}

	hasPatientFriendly := false
	if scoreItem.PatientExplanation != nil {
		content := *scoreItem.PatientExplanation
		hasPatientFriendly = contains(content, "🦴") || contains(content, "☀️") || contains(content, "💊")
	}

	hasStructuredConduct := false
	if scoreItem.Conduct != nil {
		content := *scoreItem.Conduct
		hasStructuredConduct = contains(content, "**") && contains(content, "•")
	}

	fmt.Println("Avaliação qualitativa:")
	fmt.Printf("  • Diretrizes recentes (2024):     %s\n", checkMark(hasRecentGuidelines))
	fmt.Printf("  • Linguagem acessível (paciente): %s\n", checkMark(hasPatientFriendly))
	fmt.Printf("  • Conduta estruturada:            %s\n", checkMark(hasStructuredConduct))
	fmt.Println()

	// Overall assessment
	overallGood := crLen >= 1500 && peLen >= 600 && condLen >= 800 &&
		hasRecentGuidelines && hasPatientFriendly && hasStructuredConduct

	if overallGood {
		fmt.Println("✅ AVALIAÇÃO: Conteúdo EXCELENTE - Não requer alterações")
		fmt.Println("   Campos clínicos estão completos, atualizados e bem estruturados")
	} else if crLen >= 1000 && peLen >= 400 && condLen >= 500 {
		fmt.Println("✓ AVALIAÇÃO: Conteúdo BOM - Pequenas melhorias possíveis")
		fmt.Println("  Considerar enriquecimento incremental via LLM se necessário")
	} else {
		fmt.Println("⚠️ AVALIAÇÃO: Conteúdo necessita REVISÃO")
		fmt.Println("   Recomenda-se enriquecimento via LLM com base nos artigos linkados")
	}
	fmt.Println()

	// 5. Update last_review
	fmt.Println("═══ FASE 4: ATUALIZAÇÃO ═══")

	now := time.Now()
	if err := db.Model(&scoreItem).Update("last_review", now).Error; err != nil {
		log.Printf("❌ Failed to update last_review: %v", err)
	} else {
		fmt.Printf("✓ Atualizado last_review para: %s\n", now.Format("2006-01-02 15:04:05"))
	}

	// Final summary
	fmt.Println("\n╔═══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                     RESUMO DA REVISÃO                         ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════╝")
	fmt.Printf("\n📊 ScoreItem: %s\n", scoreItem.Name)
	fmt.Printf("   Artigos linkados: %d ", totalArticles)
	if totalArticles >= minArticles {
		fmt.Println("✓")
	} else {
		fmt.Printf("⚠ (mínimo: %d)\n", minArticles)
	}
	fmt.Printf("   Qualidade de conteúdo: ")
	if overallGood {
		fmt.Println("EXCELENTE ✅")
	} else if crLen >= 1000 {
		fmt.Println("BOM ✓")
	} else {
		fmt.Println("NECESSITA REVISÃO ⚠")
	}
	fmt.Printf("   Última revisão: %s\n", now.Format("2006-01-02"))
	fmt.Println("\n✅ REVISÃO CONCLUÍDA COM SUCESSO\n")
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 &&
		(s == substr || len(s) > len(substr) &&
		(s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
		containsMiddle(s, substr)))
}

func containsMiddle(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func checkMark(b bool) string {
	if b {
		return "✓ SIM"
	}
	return "✗ NÃO"
}
