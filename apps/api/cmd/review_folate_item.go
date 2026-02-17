package main

import (
	"context"
	"encoding/json"
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
	"gorm.io/gorm"
)

func main() {
	// Load environment
	if err := godotenv.Load("/home/user/plenya/apps/api/.env"); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Database connection
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// ScoreItem ID to review
	scoreItemID := uuid.MustParse("019bf31d-2ef0-7ecf-87d3-66cd04d3a8cb")

	// Initialize services
	vectorRepo := repository.NewArticleVectorRepository(db)
	embeddingService := services.NewEmbeddingService(db)
	articleSemanticService := services.NewArticleSemanticService(vectorRepo, embeddingService)
	enrichmentService := services.NewScoreItemEnrichmentService(db)
	pubmedService := services.NewPubMedService(db)

	ctx := context.Background()

	// Step 1: Fetch current ScoreItem
	var scoreItem models.ScoreItem
	if err := db.First(&scoreItem, scoreItemID).Error; err != nil {
		log.Fatalf("Failed to fetch ScoreItem: %v", err)
	}

	fmt.Printf("=== Reviewing ScoreItem: %s ===\n\n", scoreItem.Name)

	// Step 2: Check current article count
	var currentArticleCount int64
	if err := db.Table("article_score_items").
		Where("score_item_id = ?", scoreItemID).
		Count(&currentArticleCount).Error; err != nil {
		log.Fatalf("Failed to count articles: %v", err)
	}

	fmt.Printf("Current article count: %d\n\n", currentArticleCount)

	// Step 3: Search for similar articles using RAG (threshold 0.7)
	fmt.Println("Searching for similar articles using RAG...")
	similarArticles, err := vectorRepo.FindSimilarArticlesForScoreItem(scoreItemID, 20, 0.7)
	if err != nil {
		log.Fatalf("Failed to search similar articles: %v", err)
	}

	fmt.Printf("Found %d similar articles (similarity ≥ 0.7)\n\n", len(similarArticles))

	// Step 4: Find new articles to link (not already linked)
	var alreadyLinkedIDs []uuid.UUID
	if err := db.Table("article_score_items").
		Where("score_item_id = ?", scoreItemID).
		Pluck("article_id", &alreadyLinkedIDs).Error; err != nil {
		log.Fatalf("Failed to fetch linked article IDs: %v", err)
	}

	alreadyLinkedMap := make(map[uuid.UUID]bool)
	for _, id := range alreadyLinkedIDs {
		alreadyLinkedMap[id] = true
	}

	newArticlesToLink := []repository.ArticleSearchResult{}
	for _, result := range similarArticles {
		if !alreadyLinkedMap[result.Article.ID] {
			newArticlesToLink = append(newArticlesToLink, result)
		}
	}

	fmt.Printf("New articles to link: %d\n", len(newArticlesToLink))
	for i, result := range newArticlesToLink {
		fmt.Printf("  %d. [%.3f] %s (%s)\n",
			i+1,
			result.Similarity,
			result.Article.Title,
			result.Article.Journal)
	}
	fmt.Println()

	// Step 5: Link new articles
	if len(newArticlesToLink) > 0 {
		fmt.Println("Linking new articles...")
		for _, result := range newArticlesToLink {
			linkRecord := map[string]interface{}{
				"article_id":    result.Article.ID,
				"score_item_id": scoreItemID,
				"created_at":    time.Now(),
			}
			if err := db.Table("article_score_items").Create(linkRecord).Error; err != nil {
				log.Printf("Warning: Failed to link article %s: %v", result.Article.ID, err)
			}
		}
		fmt.Printf("Linked %d new articles\n\n", len(newArticlesToLink))
	}

	// Step 6: Check if we need more articles from PubMed
	totalArticles := int(currentArticleCount) + len(newArticlesToLink)
	if totalArticles < 7 {
		fmt.Printf("Need more articles (current: %d, target: 7). Searching PubMed...\n", totalArticles)

		// Search PubMed
		query := fmt.Sprintf("(%s) AND (functional medicine OR integrative medicine OR reference ranges)", scoreItem.Name)
		pmids, err := pubmedService.SearchArticles(ctx, query, 10)
		if err != nil {
			log.Printf("Warning: PubMed search failed: %v", err)
		} else {
			fmt.Printf("Found %d PubMed articles\n", len(pmids))
			// Note: Fetching details and downloading PDFs would require full implementation
			// For now, just log the PMIDs
			for i, pmid := range pmids {
				fmt.Printf("  %d. PMID: %s\n", i+1, pmid)
			}
			fmt.Println("Note: Full PubMed integration (fetch details + download PDFs) requires additional implementation")
		}
		fmt.Println()
	}

	// Step 7: Review clinical fields using enrichment service
	fmt.Println("Reviewing clinical fields...")

	// Determine tier
	tier := enrichmentService.DetermineTier(&scoreItem)
	fmt.Printf("Enrichment tier: %s\n", tier)

	// Check current field quality
	fmt.Printf("\nCurrent field lengths:\n")
	fmt.Printf("  clinical_relevance: %d chars\n", len(ptrString(scoreItem.ClinicalRelevance)))
	fmt.Printf("  patient_explanation: %d chars\n", len(ptrString(scoreItem.PatientExplanation)))
	fmt.Printf("  conduct: %d chars\n", len(ptrString(scoreItem.Conduct)))
	fmt.Printf("  points: %.1f\n", ptrFloat(scoreItem.Points))

	if tier == services.TierPreserve {
		fmt.Println("\nContent is excellent and recent. No enrichment needed.")
	} else {
		fmt.Println("\nGenerating enrichment using Claude API...")

		// Generate enrichment
		result, err := enrichmentService.GenerateEnrichment(ctx, &scoreItem, tier)
		if err != nil {
			log.Fatalf("Failed to generate enrichment: %v", err)
		}

		// Validate result
		validationErrors := enrichmentService.ValidateResult(result)
		if len(validationErrors) > 0 {
			fmt.Println("\nValidation warnings:")
			for _, err := range validationErrors {
				fmt.Printf("  - %s\n", err)
			}
			fmt.Println()
		}

		// Display result
		fmt.Printf("\nEnrichment result (confidence: %.2f):\n", result.Confidence)
		fmt.Printf("\nclinical_relevance (%d chars):\n%s\n",
			len(result.ClinicalRelevance),
			truncate(result.ClinicalRelevance, 200))
		fmt.Printf("\npatient_explanation (%d chars):\n%s\n",
			len(result.PatientExplanation),
			truncate(result.PatientExplanation, 200))
		fmt.Printf("\nconduct (%d chars):\n%s\n",
			len(result.Conduct),
			truncate(result.Conduct, 200))
		fmt.Printf("\nmax_points: %d\n", result.MaxPoints)
		fmt.Printf("\njustification: %s\n", result.Justification)

		// Prompt user for confirmation
		fmt.Println("\n=== REVIEW REQUIRED ===")
		fmt.Println("Please review the enriched content above.")
		fmt.Print("Apply changes? (yes/no): ")

		var response string
		fmt.Scanln(&response)

		if response == "yes" || response == "y" {
			// Save review history
			if err := enrichmentService.SaveReviewHistory(&scoreItem, result, tier); err != nil {
				log.Printf("Warning: Failed to save review history: %v", err)
			}

			// Update ScoreItem
			now := time.Now()
			updates := map[string]interface{}{
				"clinical_relevance":  result.ClinicalRelevance,
				"patient_explanation": result.PatientExplanation,
				"conduct":             result.Conduct,
				"points":              float64(result.MaxPoints),
				"last_review":         now,
				"updated_at":          now,
			}

			if err := db.Model(&scoreItem).Updates(updates).Error; err != nil {
				log.Fatalf("Failed to update ScoreItem: %v", err)
			}

			fmt.Println("\n✓ ScoreItem updated successfully!")
		} else {
			fmt.Println("\n✗ Changes discarded.")
		}
	}

	// Step 8: Display final summary
	fmt.Println("\n=== Final Summary ===")

	var finalArticleCount int64
	db.Table("article_score_items").
		Where("score_item_id = ?", scoreItemID).
		Count(&finalArticleCount)

	fmt.Printf("Final article count: %d\n", finalArticleCount)
	fmt.Printf("Last review: %v\n", scoreItem.LastReview)
	fmt.Println("\nReview complete!")
}

// Helper functions
func ptrString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func ptrFloat(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
