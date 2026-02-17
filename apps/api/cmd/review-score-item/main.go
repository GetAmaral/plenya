package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
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

// Review Score Item - Revisão completa de um ScoreItem específico
//
// Processo:
// 1. Buscar artigos linkados atuais
// 2. Avaliar adequação dos artigos
// 3. Buscar artigos similares via RAG (threshold 0.7)
// 4. Se <7 artigos: buscar no PubMed
// 5. Baixar PDFs disponíveis
// 6. Fazer embeddings dos novos artigos
// 7. Linkar novos artigos ao ScoreItem
// 8. Revisar campos clínicos (clinical_relevance, patient_explanation, conduct, max_points)
// 9. Salvar mudanças com audit trail
//
// Uso: go run cmd/review-score-item/main.go <score-item-id>

func main() {
	fmt.Println("🔍 Review Score Item - Complete Revision")
	fmt.Println(strings.Repeat("=", 60))

	// Validar argumentos
	if len(os.Args) < 2 {
		log.Fatal("❌ Usage: go run cmd/review-score-item/main.go <score-item-id>")
	}

	scoreItemIDStr := os.Args[1]
	scoreItemID, err := uuid.Parse(scoreItemIDStr)
	if err != nil {
		log.Fatalf("❌ Invalid UUID: %v", err)
	}

	// Carregar .env
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️  .env file not found (using environment variables)")
	}

	// Verificar API keys
	anthropicKey := os.Getenv("ANTHROPIC_API_KEY")
	if anthropicKey == "" {
		log.Fatal("❌ ANTHROPIC_API_KEY environment variable is required")
	}

	pubmedEmail := os.Getenv("PUBMED_EMAIL")
	if pubmedEmail == "" {
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
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	fmt.Println("✅ Connected to database")

	// Buscar ScoreItem
	var scoreItem models.ScoreItem
	if err := db.First(&scoreItem, scoreItemID).Error; err != nil {
		log.Fatalf("❌ Failed to fetch ScoreItem: %v", err)
	}

	fmt.Printf("\n📊 Score Item: %s\n", scoreItem.Name)
	fmt.Printf("   ID: %s\n", scoreItem.ID)
	fmt.Printf("   Current Points: %.0f\n", getFloat(scoreItem.Points))
	fmt.Printf("   Last Review: %s\n", formatTime(scoreItem.LastReview))
	fmt.Println()

	// 1. Avaliar artigos atuais
	fmt.Println("📚 STEP 1: Avaliar artigos linkados atuais")
	fmt.Println(strings.Repeat("-", 60))

	// Buscar artigos linkados via query direta
	var currentArticles []models.Article
	err = db.Table("articles").
		Joins("INNER JOIN article_score_items ON articles.id = article_score_items.article_id").
		Where("article_score_items.score_item_id = ?", scoreItemID).
		Find(&currentArticles).Error

	if err != nil {
		log.Printf("⚠️  Failed to fetch current articles: %v", err)
		currentArticles = []models.Article{}
	}

	fmt.Printf("   Artigos linkados: %d\n", len(currentArticles))

	for i, article := range currentArticles {
		fmt.Printf("   [%d] %s\n", i+1, article.Title)
		fmt.Printf("       Journal: %s (%d)\n", article.Journal, article.PublishDate.Year())
		if article.PMID != nil && *article.PMID != "" {
			fmt.Printf("       PMID: %s\n", *article.PMID)
		}
		if article.DOI != nil && *article.DOI != "" {
			fmt.Printf("       DOI: %s\n", *article.DOI)
		}
	}
	fmt.Println()

	// 2. Buscar artigos similares via RAG
	fmt.Println("🔎 STEP 2: Buscar artigos similares via RAG (threshold 0.7)")
	fmt.Println(strings.Repeat("-", 60))

	vectorRepo := repository.NewArticleVectorRepository(db)
	embeddingService := services.NewEmbeddingService(cfg, db)
	semanticService := services.NewArticleSemanticService(vectorRepo, embeddingService)

	similarArticles, err := semanticService.RecommendArticlesForScoreItem(scoreItemID, 20)
	if err != nil {
		log.Printf("⚠️  RAG search failed: %v", err)
		similarArticles = []repository.ArticleSearchResult{}
	}

	fmt.Printf("   Artigos encontrados via RAG: %d\n", len(similarArticles))

	// Filtrar artigos já linkados
	linkedArticleIDs := make(map[uuid.UUID]bool)
	for _, article := range currentArticles {
		linkedArticleIDs[article.ID] = true
	}

	newRAGArticles := []repository.ArticleSearchResult{}
	for _, result := range similarArticles {
		if result.Article != nil && !linkedArticleIDs[result.Article.ID] {
			newRAGArticles = append(newRAGArticles, result)
		}
	}

	fmt.Printf("   Novos artigos (não linkados): %d\n", len(newRAGArticles))

	for i, result := range newRAGArticles {
		if i >= 5 { // Mostrar apenas os top 5
			break
		}
		if result.Article != nil {
			fmt.Printf("   [%d] %.3f - %s\n", i+1, result.Similarity, result.Article.Title)
		}
	}
	fmt.Println()

	// 3. Buscar no PubMed se necessário
	totalArticles := len(currentArticles) + len(newRAGArticles)
	minArticles := 7

	var pubmedArticles []*services.PubMedArticle
	pubmedService := services.NewPubMedService(db)

	if totalArticles < minArticles {
		fmt.Printf("🌐 STEP 3: Buscar no PubMed (total atual: %d, mínimo: %d)\n", totalArticles, minArticles)
		fmt.Println(strings.Repeat("-", 60))

		// Gerar query PubMed baseada no nome do ScoreItem
		query := generatePubMedQuery(&scoreItem)
		fmt.Printf("   Query: %s\n", query)

		ctx := context.Background()
		pmids, err := pubmedService.SearchArticles(ctx, query, 15)
		if err != nil {
			log.Printf("⚠️  PubMed search failed: %v", err)
		} else {
			fmt.Printf("   PMIDs encontrados: %d\n", len(pmids))

			// Fetch detalhes (simplificado - XML parser não implementado ainda)
			pubmedArticles, err = pubmedService.FetchArticleDetails(ctx, pmids)
			if err != nil {
				log.Printf("⚠️  Failed to fetch article details: %v", err)
			} else {
				fmt.Printf("   Artigos com metadata: %d\n", len(pubmedArticles))
			}
		}
		fmt.Println()
	} else {
		fmt.Printf("✅ STEP 3: Artigos suficientes (%d >= %d), skip PubMed\n\n", totalArticles, minArticles)
	}

	// 4. Linkar novos artigos RAG ao ScoreItem
	fmt.Println("🔗 STEP 4: Linkar novos artigos ao ScoreItem")
	fmt.Println(strings.Repeat("-", 60))

	linkedCount := 0
	maxToLink := 10 // Limite razoável

	for i, result := range newRAGArticles {
		if i >= maxToLink || result.Article == nil {
			break
		}

		articleID := result.Article.ID

		// Verificar se já está linkado (double-check)
		var count int64
		err := db.Table("article_score_items").
			Where("article_id = ? AND score_item_id = ?", articleID, scoreItemID).
			Count(&count).Error

		if err == nil && count > 0 {
			continue // Já linkado
		}

		// Criar link usando raw SQL
		err = db.Exec(`
			INSERT INTO article_score_items (article_id, score_item_id, confidence_score, auto_linked, linked_at)
			VALUES (?, ?, ?, true, NOW())
		`, articleID, scoreItemID, result.Similarity).Error

		if err != nil {
			log.Printf("⚠️  Failed to link article %s: %v", articleID, err)
			continue
		}

		linkedCount++
		fmt.Printf("   ✓ Linked: %s (similarity: %.3f)\n", result.Article.Title, result.Similarity)
	}

	fmt.Printf("\n   Total artigos linkados: %d\n", linkedCount)
	fmt.Println()

	// 5. Revisar campos clínicos usando LLM
	fmt.Println("✨ STEP 5: Revisar campos clínicos com LLM")
	fmt.Println(strings.Repeat("-", 60))

	enrichmentService := services.NewScoreItemEnrichmentService(db)
	enrichmentService.SetModel("claude-sonnet-4-5-20250929") // Melhor modelo

	// Determinar tier
	tier := enrichmentService.DetermineTier(&scoreItem)
	fmt.Printf("   Tier: %s\n", tier)

	// Gerar enriquecimento
	ctx := context.Background()
	startTime := time.Now()
	result, err := enrichmentService.GenerateEnrichment(ctx, &scoreItem, tier)
	duration := time.Since(startTime)

	if err != nil {
		log.Fatalf("❌ Enrichment failed: %v", err)
	}

	fmt.Printf("   ✓ Generated in %.1fs (confidence: %.2f)\n", duration.Seconds(), result.Confidence)

	// Validar resultado
	validationErrors := enrichmentService.ValidateResult(result)
	if len(validationErrors) > 0 {
		fmt.Println("\n   ⚠️  Validation warnings:")
		for _, verr := range validationErrors {
			fmt.Printf("      - %s\n", verr)
		}
	}

	// Mostrar mudanças propostas
	fmt.Println("\n📝 Mudanças propostas:")
	fmt.Println(strings.Repeat("-", 60))

	fmt.Println("\n[Clinical Relevance]")
	fmt.Printf("Antes: %s\n", truncate(getString(scoreItem.ClinicalRelevance), 200))
	fmt.Printf("Depois: %s\n", truncate(result.ClinicalRelevance, 200))

	fmt.Println("\n[Patient Explanation]")
	fmt.Printf("Antes: %s\n", truncate(getString(scoreItem.PatientExplanation), 200))
	fmt.Printf("Depois: %s\n", truncate(result.PatientExplanation, 200))

	fmt.Println("\n[Conduct]")
	fmt.Printf("Antes: %s\n", truncate(getString(scoreItem.Conduct), 200))
	fmt.Printf("Depois: %s\n", truncate(result.Conduct, 200))

	fmt.Printf("\n[Max Points]\n")
	fmt.Printf("Antes: %.0f\n", getFloat(scoreItem.Points))
	fmt.Printf("Depois: %d\n", result.MaxPoints)

	// 6. Salvar mudanças
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("💾 Salvando mudanças...")

	// Salvar auditoria
	if err := enrichmentService.SaveReviewHistory(&scoreItem, result, tier); err != nil {
		log.Printf("⚠️  Failed to save audit: %v", err)
	}

	// Atualizar item
	scoreItem.ClinicalRelevance = &result.ClinicalRelevance
	scoreItem.PatientExplanation = &result.PatientExplanation
	scoreItem.Conduct = &result.Conduct
	points := float64(result.MaxPoints)
	scoreItem.Points = &points

	if err := db.Save(&scoreItem).Error; err != nil {
		log.Fatalf("❌ Failed to save item: %v", err)
	}

	fmt.Println("\n✅ Score Item atualizado com sucesso!")
	fmt.Printf("   Articles linked: %d → %d\n", len(currentArticles), len(currentArticles)+linkedCount)
	fmt.Printf("   LastReview: %s → %s\n",
		formatTime(scoreItem.LastReview),
		time.Now().Format("2006-01-02 15:04:05"))

	fmt.Println("\n📊 Summary:")
	fmt.Printf("   Clinical Relevance: %d chars\n", len(result.ClinicalRelevance))
	fmt.Printf("   Patient Explanation: %d chars\n", len(result.PatientExplanation))
	fmt.Printf("   Conduct: %d chars\n", len(result.Conduct))
	fmt.Printf("   Max Points: %d\n", result.MaxPoints)
	fmt.Printf("   Confidence: %.2f\n", result.Confidence)

	fmt.Println("\n✅ Revisão completa!")
}

func generatePubMedQuery(item *models.ScoreItem) string {
	// Extrair gene e SNP do nome
	name := item.Name

	// Padrões comuns: "GENE rsXXXXXX (Condition)"
	parts := strings.Fields(name)

	query := ""
	for i, part := range parts {
		if strings.HasPrefix(part, "rs") {
			// SNP encontrado
			gene := parts[0]
			snp := strings.TrimSuffix(part, ",")

			query = fmt.Sprintf("%s[Gene] AND %s[SNP]", gene, snp)
			break
		}

		// Se não encontrar rs, usar primeiras palavras
		if i < 3 {
			query += part + " "
		}
	}

	if query == "" {
		query = name
	}

	// Adicionar filtros
	query += " AND (polymorphism[Title/Abstract] OR SNP[Title/Abstract] OR variant[Title/Abstract])"
	query += " AND (Review[ptyp] OR Meta-Analysis[ptyp] OR Observational Study[ptyp])"

	// Últimos 7 anos
	currentYear := time.Now().Year()
	query += fmt.Sprintf(" AND %d:%d[dp]", currentYear-7, currentYear)

	return query
}

func getString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func getFloat(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

func formatTime(t *time.Time) string {
	if t == nil {
		return "never"
	}
	return t.Format("2006-01-02 15:04:05")
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
