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

// ReviewScoreItemTask coordena revisão completa de um ScoreItem
type ReviewScoreItemTask struct {
	db                  *gorm.DB
	pubmedService       *services.PubMedService
	enrichmentService   *services.ScoreItemEnrichmentService
	vectorRepo          *repository.ArticleVectorRepository
	scoreItemID         uuid.UUID
	minArticles         int
	ragThreshold        float64
}

func main() {
	// Load .env
	if err := godotenv.Load("/home/user/plenya/apps/api/.env"); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Parse ScoreItem ID from args
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run review_score_item.go <score_item_id>")
	}

	scoreItemID, err := uuid.Parse(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid UUID: %v", err)
	}

	// Initialize DB
	cfg := config.LoadConfig()
	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create services
	task := &ReviewScoreItemTask{
		db:                db,
		pubmedService:     services.NewPubMedService(db),
		enrichmentService: services.NewScoreItemEnrichmentService(db),
		vectorRepo:        repository.NewArticleVectorRepository(db),
		scoreItemID:       scoreItemID,
		minArticles:       7,
		ragThreshold:      0.7,
	}

	// Run review
	ctx := context.Background()
	if err := task.Run(ctx); err != nil {
		log.Fatalf("Review failed: %v", err)
	}

	log.Println("✅ Review completed successfully!")
}

func (t *ReviewScoreItemTask) Run(ctx context.Context) error {
	log.Printf("🔍 Starting review for ScoreItem: %s\n", t.scoreItemID)

	// 1. Fetch ScoreItem
	var item models.ScoreItem
	if err := t.db.Preload("Subgroup.Group").First(&item, t.scoreItemID).Error; err != nil {
		return fmt.Errorf("failed to fetch ScoreItem: %w", err)
	}

	log.Printf("📌 ScoreItem: %s (Group: %s > Subgroup: %s)\n",
		item.Name,
		item.Subgroup.Group.Name,
		item.Subgroup.Name)

	// 2. Review articles
	if err := t.reviewArticles(ctx, &item); err != nil {
		return fmt.Errorf("article review failed: %w", err)
	}

	// 3. Review clinical fields
	if err := t.reviewClinicalFields(ctx, &item); err != nil {
		return fmt.Errorf("clinical fields review failed: %w", err)
	}

	// 4. Update last_review
	if err := t.db.Model(&item).Update("last_review", time.Now()).Error; err != nil {
		return fmt.Errorf("failed to update last_review: %w", err)
	}

	return nil
}

func (t *ReviewScoreItemTask) reviewArticles(ctx context.Context, item *models.ScoreItem) error {
	log.Println("\n📚 STEP 1: Article Review")

	// Count current linked articles
	var linkedCount int64
	if err := t.db.Model(&models.ArticleScoreItem{}).
		Where("score_item_id = ?", item.ID).
		Count(&linkedCount).Error; err != nil {
		return fmt.Errorf("failed to count linked articles: %w", err)
	}

	log.Printf("Current linked articles: %d\n", linkedCount)

	// Fetch linked articles to evaluate quality
	var linkedArticles []models.Article
	if err := t.db.
		Joins("INNER JOIN article_score_items asi ON asi.article_id = articles.id").
		Where("asi.score_item_id = ?", item.ID).
		Find(&linkedArticles).Error; err != nil {
		return fmt.Errorf("failed to fetch linked articles: %w", err)
	}

	// Check quality of linked articles (have DOI or PMID?)
	qualityArticles := 0
	for _, article := range linkedArticles {
		if (article.DOI != nil && *article.DOI != "") ||
			(article.PMID != nil && *article.PMID != "") {
			qualityArticles++
		}
	}

	log.Printf("Quality articles (with DOI/PMID): %d\n", qualityArticles)

	// Need more articles?
	articlesNeeded := t.minArticles - qualityArticles
	if articlesNeeded <= 0 {
		log.Println("✅ Sufficient quality articles already linked")
		return nil
	}

	log.Printf("🔍 Need %d more quality articles\n", articlesNeeded)

	// Try RAG search first
	ragArticles, err := t.searchArticlesViaRAG(item, articlesNeeded)
	if err != nil {
		log.Printf("⚠️  RAG search failed: %v\n", err)
	} else {
		log.Printf("Found %d articles via RAG\n", len(ragArticles))
		if err := t.linkArticles(item.ID, ragArticles, true); err != nil {
			return fmt.Errorf("failed to link RAG articles: %w", err)
		}
		articlesNeeded -= len(ragArticles)
	}

	// Still need more? Search PubMed
	if articlesNeeded > 0 {
		log.Printf("🔍 Searching PubMed for %d more articles\n", articlesNeeded)
		pubmedArticles, err := t.searchAndDownloadFromPubMed(ctx, item, articlesNeeded)
		if err != nil {
			log.Printf("⚠️  PubMed search failed: %v\n", err)
		} else {
			log.Printf("Found %d articles from PubMed\n", len(pubmedArticles))
			if err := t.linkArticles(item.ID, pubmedArticles, true); err != nil {
				return fmt.Errorf("failed to link PubMed articles: %w", err)
			}
		}
	}

	return nil
}

func (t *ReviewScoreItemTask) searchArticlesViaRAG(item *models.ScoreItem, limit int) ([]*models.Article, error) {
	// Use FindSimilarArticlesForScoreItem
	results, err := t.vectorRepo.FindSimilarArticlesForScoreItem(
		item.ID,
		limit,
		t.ragThreshold,
	)
	if err != nil {
		return nil, err
	}

	articles := make([]*models.Article, 0, len(results))
	for _, result := range results {
		// Only include if has DOI or PMID
		if (result.Article.DOI != nil && *result.Article.DOI != "") ||
			(result.Article.PMID != nil && *result.Article.PMID != "") {
			articles = append(articles, result.Article)
		}
	}

	return articles, nil
}

func (t *ReviewScoreItemTask) searchAndDownloadFromPubMed(
	ctx context.Context,
	item *models.ScoreItem,
	limit int,
) ([]*models.Article, error) {
	// Generate optimized PubMed query
	query := t.generatePubMedQuery(item)
	log.Printf("PubMed query: %s\n", query)

	// Search PubMed
	pmids, err := t.pubmedService.RateLimitedSearch(ctx, query, limit*3) // Search more to filter later
	if err != nil {
		return nil, fmt.Errorf("PubMed search failed: %w", err)
	}

	if len(pmids) == 0 {
		log.Println("No articles found on PubMed")
		return []*models.Article{}, nil
	}

	log.Printf("Found %d PMIDs\n", len(pmids))

	// Fetch article details
	pubmedArticles, err := t.pubmedService.FetchArticleDetails(ctx, pmids)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch article details: %w", err)
	}

	// Try to download PDFs and create articles
	articles := make([]*models.Article, 0, limit)
	for i, pmArticle := range pubmedArticles {
		if len(articles) >= limit {
			break
		}

		// Try to download PDF if DOI exists
		var pdfPath string
		if pmArticle.DOI != "" {
			log.Printf("Trying to download PDF for DOI: %s\n", pmArticle.DOI)
			articleID := uuid.Must(uuid.NewV7())
			pdfPath, err = t.pubmedService.DownloadPDF(ctx, pmArticle.DOI, articleID)
			if err != nil {
				log.Printf("⚠️  PDF download failed for %s: %v\n", pmArticle.DOI, err)
				pdfPath = "" // Continue without PDF
			} else {
				log.Printf("✅ Downloaded PDF: %s\n", pdfPath)
			}
		}

		// Create article in DB
		article, err := t.pubmedService.CreateArticleFromPubMed(pmArticle, pdfPath)
		if err != nil {
			log.Printf("⚠️  Failed to create article %d: %v\n", i+1, err)
			continue
		}

		articles = append(articles, article)

		// Rate limit
		time.Sleep(200 * time.Millisecond)
	}

	return articles, nil
}

func (t *ReviewScoreItemTask) generatePubMedQuery(item *models.ScoreItem) string {
	// Build query based on item context
	var query string

	// Base query from item name and context
	groupName := item.Subgroup.Group.Name
	subgroupName := item.Subgroup.Name

	// For "6 meses" in "Objetivos > Percepção de futuro"
	// Focus on: goal setting, patient engagement, short-term health goals
	if groupName == "Objetivos" {
		query = "(goal setting[MeSH] OR patient engagement[MeSH] OR health goals OR treatment adherence[MeSH]) AND (short-term OR 6 months)"
	} else {
		// Generic query
		query = fmt.Sprintf("\"%s\" AND %s", item.Name, subgroupName)
	}

	// Add filters
	query += " AND (Review[ptyp] OR Meta-Analysis[ptyp] OR Systematic Review[ptyp])"

	// Last 10 years
	currentYear := time.Now().Year()
	query += fmt.Sprintf(" AND %d:%d[dp]", currentYear-10, currentYear)

	// English only
	query += " AND english[LA]"

	return query
}

func (t *ReviewScoreItemTask) linkArticles(scoreItemID uuid.UUID, articles []*models.Article, autoLinked bool) error {
	for _, article := range articles {
		// Check if already linked
		var exists int64
		if err := t.db.Model(&models.ArticleScoreItem{}).
			Where("score_item_id = ? AND article_id = ?", scoreItemID, article.ID).
			Count(&exists).Error; err != nil {
			return err
		}

		if exists > 0 {
			log.Printf("Article %s already linked, skipping\n", article.Title)
			continue
		}

		// Link
		link := &models.ArticleScoreItem{
			ScoreItemID:     scoreItemID,
			ArticleID:       article.ID,
			ConfidenceScore: 0.8, // Default confidence
			AutoLinked:      autoLinked,
		}

		if err := t.db.Create(link).Error; err != nil {
			return fmt.Errorf("failed to link article %s: %w", article.Title, err)
		}

		log.Printf("✅ Linked article: %s\n", article.Title)
	}

	return nil
}

func (t *ReviewScoreItemTask) reviewClinicalFields(ctx context.Context, item *models.ScoreItem) error {
	log.Println("\n🏥 STEP 2: Clinical Fields Review")

	// Determine enrichment tier
	tier := t.enrichmentService.DetermineTier(item)
	log.Printf("Enrichment tier: %s\n", tier)

	if tier == services.TierPreserve {
		log.Println("✅ Content is excellent and recent, preserving")
		return nil
	}

	// Generate enrichment
	log.Println("🤖 Generating enrichment with LLM...")
	result, err := t.enrichmentService.GenerateEnrichment(ctx, item, tier)
	if err != nil {
		return fmt.Errorf("enrichment generation failed: %w", err)
	}

	// Validate result
	validationErrors := t.enrichmentService.ValidateResult(result)
	if len(validationErrors) > 0 {
		log.Println("⚠️  Validation warnings:")
		for _, e := range validationErrors {
			log.Printf("  - %s\n", e)
		}
	}

	// Display result
	log.Println("\n📝 Generated content:")
	log.Printf("Clinical Relevance (%d chars):\n%s\n\n", len(result.ClinicalRelevance), result.ClinicalRelevance)
	log.Printf("Patient Explanation (%d chars):\n%s\n\n", len(result.PatientExplanation), result.PatientExplanation)
	log.Printf("Conduct (%d chars):\n%s\n\n", len(result.Conduct), result.Conduct)
	log.Printf("Max Points: %d\n", result.MaxPoints)
	log.Printf("Confidence: %.2f\n", result.Confidence)
	log.Printf("Justification: %s\n", result.Justification)

	// Save review history
	if err := t.enrichmentService.SaveReviewHistory(item, result, tier); err != nil {
		log.Printf("⚠️  Failed to save review history: %v\n", err)
	}

	// Update ScoreItem
	updates := map[string]interface{}{
		"clinical_relevance":  result.ClinicalRelevance,
		"patient_explanation": result.PatientExplanation,
		"conduct":             result.Conduct,
	}

	// Only update points if currently 0 or result suggests change
	currentPoints := float64(0)
	if item.Points != nil {
		currentPoints = *item.Points
	}

	if currentPoints == 0 && result.MaxPoints > 0 {
		points := float64(result.MaxPoints)
		updates["points"] = points
		log.Printf("Updating points: 0 → %d\n", result.MaxPoints)
	}

	if err := t.db.Model(item).Updates(updates).Error; err != nil {
		return fmt.Errorf("failed to update ScoreItem: %w", err)
	}

	log.Println("✅ Clinical fields updated successfully")

	return nil
}
