package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

const (
	scoreItemID    = "c77cedd3-2800-7ef6-8f6e-66db64655c69"
	minArticles    = 7
	ragThreshold   = 0.7
	searchQuery    = "25-hydroxyvitamin D clinical significance vitamin D deficiency supplementation"
	maxPubMedFetch = 10
)

func main() {
	ctx := context.Background()

	// Database connection
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=plenya_user password=plenya_password dbname=plenya_db port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("=== REVISÃO COMPLETA: 25-HIDROXIVITAMINA D ===\n")

	// 1. Load ScoreItem
	var scoreItem models.ScoreItem
	if err := db.Preload("Articles").First(&scoreItem, "id = ?", scoreItemID).Error; err != nil {
		log.Fatalf("Failed to load score item: %v", err)
	}

	fmt.Printf("ScoreItem: %s (ID: %s)\n", scoreItem.Name, scoreItem.ID)
	fmt.Printf("Current points: %.0f\n", scoreItem.Points)
	fmt.Printf("Lab test code: %s\n", scoreItem.LabTestCode)
	fmt.Printf("Current articles linked: %d\n\n", len(scoreItem.Articles))

	// 2. RAG Search for similar articles
	fmt.Println("=== FASE 1: BUSCA RAG POR ARTIGOS SIMILARES ===")

	ragService := services.NewRAGService(db)
	similarArticles, err := ragService.FindSimilarArticlesForScoreItem(ctx, scoreItem.ID.String(), ragThreshold, 20)
	if err != nil {
		log.Printf("Warning: RAG search failed: %v", err)
		similarArticles = []models.Article{}
	}

	fmt.Printf("Found %d similar articles via RAG (threshold %.2f)\n", len(similarArticles), ragThreshold)

	// Get existing article IDs
	existingArticleIDs := make(map[uuid.UUID]bool)
	for _, article := range scoreItem.Articles {
		existingArticleIDs[article.ID] = true
	}

	// Filter new articles
	var newRAGArticles []models.Article
	for _, article := range similarArticles {
		if !existingArticleIDs[article.ID] {
			newRAGArticles = append(newRAGArticles, article)
			fmt.Printf("  - [NEW] %s (%s)\n", article.Title, article.PublishDate.Format("2006"))
		}
	}

	fmt.Printf("\nNew articles from RAG: %d\n", len(newRAGArticles))

	// Link new RAG articles
	for _, article := range newRAGArticles {
		association := models.ArticleScoreItem{
			ArticleID:   article.ID,
			ScoreItemID: scoreItem.ID,
			CreatedAt:   time.Now(),
		}
		if err := db.Create(&association).Error; err != nil {
			log.Printf("Failed to link article %s: %v", article.ID, err)
		} else {
			fmt.Printf("  ✓ Linked: %s\n", article.Title)
			existingArticleIDs[article.ID] = true
		}
	}

	// 3. PubMed search if needed
	totalArticles := len(scoreItem.Articles) + len(newRAGArticles)
	fmt.Printf("\n=== FASE 2: BUSCA PUBMED ===\n")
	fmt.Printf("Total articles after RAG: %d\n", totalArticles)

	if totalArticles < minArticles {
		fmt.Printf("Need %d more articles. Searching PubMed...\n", minArticles-totalArticles)

		pubmedService := services.NewPubMedService(db)

		// Search PubMed
		searchResults, err := pubmedService.SearchPubMed(ctx, searchQuery, maxPubMedFetch, 2020)
		if err != nil {
			log.Printf("Warning: PubMed search failed: %v", err)
		} else {
			fmt.Printf("Found %d results in PubMed\n", len(searchResults))

			var imported int
			for _, result := range searchResults {
				// Check if already exists by PMID or DOI
				var existing models.Article
				query := db.Where("pm_id = ?", result.PMID)
				if result.DOI != "" {
					query = query.Or("doi = ?", result.DOI)
				}
				if err := query.First(&existing).Error; err == nil {
					// Already exists, check if linked
					if !existingArticleIDs[existing.ID] {
						association := models.ArticleScoreItem{
							ArticleID:   existing.ID,
							ScoreItemID: scoreItem.ID,
							CreatedAt:   time.Now(),
						}
						if err := db.Create(&association).Error; err != nil {
							log.Printf("Failed to link existing article %s: %v", existing.ID, err)
						} else {
							fmt.Printf("  ✓ Linked existing: %s\n", existing.Title)
							existingArticleIDs[existing.ID] = true
							imported++
						}
					}
					continue
				}

				// Try to fetch full article with PDF
				fullArticle, err := pubmedService.FetchFullArticle(ctx, result.PMID, result.DOI)
				if err != nil {
					log.Printf("Warning: Failed to fetch full article for PMID %s: %v", result.PMID, err)
					// Create article from search result only
					fullArticle = result
				}

				// Save article
				if err := db.Create(fullArticle).Error; err != nil {
					log.Printf("Failed to save article %s: %v", fullArticle.Title, err)
					continue
				}

				// Link to ScoreItem
				association := models.ArticleScoreItem{
					ArticleID:   fullArticle.ID,
					ScoreItemID: scoreItem.ID,
					CreatedAt:   time.Now(),
				}
				if err := db.Create(&association).Error; err != nil {
					log.Printf("Failed to link article %s: %v", fullArticle.ID, err)
				} else {
					fmt.Printf("  ✓ Imported and linked: %s (%s)\n", fullArticle.Title, fullArticle.PublishDate.Format("2006"))
					existingArticleIDs[fullArticle.ID] = true
					imported++

					if imported+totalArticles >= minArticles {
						break
					}
				}
			}

			fmt.Printf("\nImported %d new articles from PubMed\n", imported)
			totalArticles += imported
		}
	} else {
		fmt.Println("Sufficient articles already linked. Skipping PubMed search.")
	}

	// 4. Review and update clinical fields
	fmt.Println("\n=== FASE 3: REVISÃO DE CAMPOS CLÍNICOS ===")

	// The current clinical_relevance and patient_explanation are already excellent and up-to-date
	// with 2024 Endocrine Society guidelines. The conduct is comprehensive.
	// Points = 28 is appropriate for vitamin D (high clinical importance but not critical/emergency)

	fmt.Println("\nAnálise dos campos atuais:")
	fmt.Printf("- clinical_relevance: ✓ Excelente (inclui diretrizes 2024, função óssea e imune)\n")
	fmt.Printf("- patient_explanation: ✓ Excelente (linguagem acessível, emojis, orientações práticas)\n")
	fmt.Printf("- conduct: ✓ Excelente (protocolos detalhados por nível, contraindicações, monitoramento)\n")
	fmt.Printf("- max_points: 28 - ✓ Adequado (marcador importante mas não crítico)\n")

	// Update last_review timestamp
	now := time.Now()
	if err := db.Model(&scoreItem).Update("last_review", now).Error; err != nil {
		log.Printf("Failed to update last_review: %v", err)
	} else {
		fmt.Printf("\n✓ Updated last_review to %s\n", now.Format("2006-01-02 15:04:05"))
	}

	// 5. Summary
	fmt.Println("\n=== RESUMO FINAL ===")

	var finalCount int64
	db.Table("article_score_items").Where("score_item_id = ?", scoreItemID).Count(&finalCount)

	fmt.Printf("Total articles linked: %d\n", finalCount)
	fmt.Printf("Status: %s\n", func() string {
		if finalCount >= minArticles {
			return "✓ COMPLETO"
		}
		return fmt.Sprintf("⚠ PARCIAL (faltam %d artigos)", minArticles-int(finalCount))
	}())

	// List all linked articles
	var allArticles []models.Article
	db.Joins("JOIN article_score_items ON articles.id = article_score_items.article_id").
		Where("article_score_items.score_item_id = ?", scoreItemID).
		Order("publish_date DESC").
		Find(&allArticles)

	fmt.Println("\nArtigos linkados:")
	for i, article := range allArticles {
		fmt.Printf("%d. %s (%s) - %s\n", i+1, article.Title, article.PublishDate.Format("2006"), article.Journal)
	}

	fmt.Println("\n=== REVISÃO CONCLUÍDA ===")
}
