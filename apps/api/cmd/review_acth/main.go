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
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
	"github.com/plenya/api/internal/services"
	"gorm.io/gorm"
)

const (
	acthID = "c77cedd3-2800-77e4-b510-d4cccf642c0d"
	minArticles = 7
	ragThreshold = 0.7
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Connect to database
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	ctx := context.Background()

	// Initialize services
	vectorRepo := repository.NewArticleVectorRepository(db)
	pubmedService := services.NewPubMedService(db)
	enrichmentService := services.NewScoreItemEnrichmentService(db)

	// 1. Buscar ScoreItem
	var scoreItem models.ScoreItem
	if err := db.First(&scoreItem, "id = ?", acthID).Error; err != nil {
		log.Fatalf("Failed to fetch ScoreItem: %v", err)
	}

	fmt.Printf("\n=== REVISÃO DE SCORE ITEM: %s ===\n", scoreItem.Name)
	fmt.Printf("ID: %s\n", scoreItem.ID)
	fmt.Printf("Unidade: %s\n", stringValue(scoreItem.Unit))
	fmt.Printf("Pontos atuais: %.1f\n\n", getPoints(scoreItem.Points))

	// 2. Buscar artigos já linkados
	existingArticles, err := getLinkedArticles(db, scoreItem.ID)
	if err != nil {
		log.Fatalf("Failed to fetch linked articles: %v", err)
	}

	fmt.Printf("Artigos já linkados: %d\n", len(existingArticles))
	for i, article := range existingArticles {
		pmid := "N/A"
		if article.PMID != nil {
			pmid = *article.PMID
		}
		fmt.Printf("  %d. [PMID: %s] %s (%d)\n", i+1, pmid, article.Title, article.PublishDate.Year())
	}

	// Avaliar qualidade: artigos sem PMID = não científicos
	scientificCount := 0
	for _, article := range existingArticles {
		if article.PMID != nil && *article.PMID != "" {
			scientificCount++
		}
	}
	fmt.Printf("\nArtigos científicos (com PMID): %d\n", scientificCount)

	// 3. Buscar artigos similares via RAG
	fmt.Printf("\n=== BUSCA RAG (threshold: %.2f) ===\n", ragThreshold)
	ragResults, err := vectorRepo.FindSimilarArticlesForScoreItem(
		scoreItem.ID,
		20,           // buscar mais para avaliar
		ragThreshold,
	)
	if err != nil {
		log.Printf("RAG search failed: %v (continuing...)", err)
	} else {
		fmt.Printf("Artigos encontrados via RAG: %d\n", len(ragResults))
		for i, result := range ragResults {
			if i >= 5 { // mostrar apenas top 5
				break
			}
			pmid := "N/A"
			if result.Article.PMID != nil {
				pmid = *result.Article.PMID
			}
			fmt.Printf("  %d. [PMID: %s, Sim: %.3f] %s\n",
				i+1, pmid, result.Similarity, result.Article.Title)
		}
	}

	// 4. Se < 7 artigos científicos, buscar no PubMed
	totalQualityArticles := scientificCount + len(ragResults)
	if totalQualityArticles < minArticles {
		fmt.Printf("\n=== BUSCA PUBMED (necessário: %d artigos) ===\n", minArticles - totalQualityArticles)

		// Gerar query PubMed
		query := generatePubMedQuery(scoreItem.Name)
		fmt.Printf("Query: %s\n", query)

		pmids, err := pubmedService.SearchArticles(ctx, query, minArticles * 2)
		if err != nil {
			log.Printf("PubMed search failed: %v", err)
		} else {
			fmt.Printf("PMIDs encontrados: %d\n", len(pmids))

			// Fetch details
			if len(pmids) > 0 {
				articles, err := pubmedService.FetchArticleDetails(ctx, pmids)
				if err != nil {
					log.Printf("Failed to fetch article details: %v", err)
				} else {
					fmt.Printf("Artigos obtidos: %d\n", len(articles))

					// Salvar artigos e linkar ao ScoreItem
					for _, pubmedArticle := range articles {
						article, err := pubmedService.CreateArticleFromPubMed(pubmedArticle, "")
						if err != nil {
							log.Printf("Failed to create article %s: %v", pubmedArticle.PMID, err)
							continue
						}

						// Linkar ao ScoreItem
						if err := linkArticleToScoreItem(db, article.ID, scoreItem.ID); err != nil {
							log.Printf("Failed to link article %s: %v", article.ID, err)
						} else {
							fmt.Printf("  ✓ Linkado: %s\n", article.Title)
						}
					}
				}
			}
		}
	} else {
		fmt.Printf("\n✓ Artigos suficientes (%d ≥ %d)\n", totalQualityArticles, minArticles)
	}

	// 5. Gerar enriquecimento LLM
	fmt.Printf("\n=== ENRIQUECIMENTO LLM ===\n")

	// Problema detectado: conteúdo atual está ERRADO (Anti-TPO em vez de ACTH)
	fmt.Println("⚠️  PROBLEMA DETECTADO: Conteúdo clínico atual é sobre Anti-TPO, não ACTH!")
	fmt.Println("    Tier selecionado: REWRITE (reescrever do zero)\n")

	result, err := enrichmentService.GenerateEnrichment(ctx, &scoreItem, services.TierRewrite)
	if err != nil {
		log.Fatalf("Enrichment failed: %v", err)
	}

	// Validar resultado
	validationErrors := enrichmentService.ValidateResult(result)
	if len(validationErrors) > 0 {
		fmt.Println("⚠️  Erros de validação:")
		for _, e := range validationErrors {
			fmt.Printf("  - %s\n", e)
		}

		if shouldContinue() {
			fmt.Println("Continuando apesar dos erros...")
		} else {
			log.Fatal("Abortado pelo usuário")
		}
	}

	// Mostrar preview
	fmt.Printf("\n=== PREVIEW DO RESULTADO ===\n")
	fmt.Printf("Confidence: %.2f\n", result.Confidence)
	fmt.Printf("Max Points: %d\n\n", result.MaxPoints)

	fmt.Println("Clinical Relevance:")
	fmt.Printf("%s\n\n", truncate(result.ClinicalRelevance, 300))

	fmt.Println("Patient Explanation:")
	fmt.Printf("%s\n\n", truncate(result.PatientExplanation, 200))

	fmt.Println("Conduct:")
	fmt.Printf("%s\n\n", truncate(result.Conduct, 300))

	fmt.Printf("Justification: %s\n\n", result.Justification)

	// 6. Confirmar e salvar
	if !confirmSave() {
		fmt.Println("Abortado pelo usuário")
		return
	}

	// Salvar histórico de revisão
	if err := enrichmentService.SaveReviewHistory(&scoreItem, result, services.TierRewrite); err != nil {
		log.Printf("Warning: Failed to save review history: %v", err)
	}

	// Atualizar ScoreItem
	now := time.Now()
	updates := map[string]interface{}{
		"clinical_relevance":  result.ClinicalRelevance,
		"patient_explanation": result.PatientExplanation,
		"conduct":             result.Conduct,
		"points":              float64(result.MaxPoints),
		"last_review":         now,
	}

	if err := db.Model(&scoreItem).Updates(updates).Error; err != nil {
		log.Fatalf("Failed to update ScoreItem: %v", err)
	}

	fmt.Printf("\n✅ ScoreItem atualizado com sucesso!\n")
	fmt.Printf("Last Review: %s\n", now.Format("2006-01-02 15:04:05"))
}

func getLinkedArticles(db *gorm.DB, scoreItemID uuid.UUID) ([]models.Article, error) {
	var articles []models.Article
	err := db.
		Joins("INNER JOIN article_score_items ON articles.id = article_score_items.article_id").
		Where("article_score_items.score_item_id = ?", scoreItemID).
		Order("articles.publish_date DESC").
		Find(&articles).Error
	return articles, err
}

func linkArticleToScoreItem(db *gorm.DB, articleID, scoreItemID uuid.UUID) error {
	// Verificar se já existe
	var count int64
	db.Table("article_score_items").
		Where("article_id = ? AND score_item_id = ?", articleID, scoreItemID).
		Count(&count)

	if count > 0 {
		return nil // Já existe
	}

	return db.Exec(`
		INSERT INTO article_score_items (id, article_id, score_item_id, created_at)
		VALUES (?, ?, ?, NOW())
	`, uuid.Must(uuid.NewV7()), articleID, scoreItemID).Error
}

func generatePubMedQuery(itemName string) string {
	// Query específica para ACTH
	if strings.Contains(strings.ToUpper(itemName), "ACTH") {
		return "ACTH[tiab] AND (reference values[tiab] OR normal range[tiab] OR diagnosis[tiab] OR clinical significance[tiab]) AND (Review[ptyp] OR Meta-Analysis[ptyp]) AND 2019:2026[dp]"
	}

	// Query genérica
	return fmt.Sprintf("%s AND (Review[ptyp] OR Meta-Analysis[ptyp]) AND 2019:2026[dp]", itemName)
}

func stringValue(s *string) string {
	if s == nil {
		return "N/A"
	}
	return *s
}

func getPoints(p *float64) float64 {
	if p == nil {
		return 0
	}
	return *p
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

func shouldContinue() bool {
	fmt.Print("Continuar mesmo assim? (y/N): ")
	var response string
	fmt.Scanln(&response)
	return strings.ToLower(response) == "y"
}

func confirmSave() bool {
	fmt.Print("\nSalvar mudanças no banco de dados? (y/N): ")
	var response string
	fmt.Scanln(&response)
	return strings.ToLower(response) == "y"
}
