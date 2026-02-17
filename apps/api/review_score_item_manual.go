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
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Load .env
	if err := godotenv.Load("/home/user/plenya/apps/api/.env"); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Parse ScoreItem ID from args
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run review_score_item_manual.go <score_item_id>")
	}

	scoreItemID, err := uuid.Parse(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid UUID: %v", err)
	}

	// Initialize DB
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Import database package
	db, err := initDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run review
	ctx := context.Background()
	if err := reviewScoreItem(ctx, db, scoreItemID); err != nil {
		log.Fatalf("Review failed: %v", err)
	}

	log.Println("\n✅ Review analysis completed!")
}

func initDatabase(cfg *config.Config) (*gorm.DB, error) {
	gormLogger := logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: gormLogger,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func reviewScoreItem(ctx context.Context, db *gorm.DB, scoreItemID uuid.UUID) error {
	log.Printf("🔍 Starting review for ScoreItem: %s\n", scoreItemID)

	// 1. Fetch ScoreItem with relationships
	var item models.ScoreItem
	if err := db.Preload("Subgroup.Group").First(&item, scoreItemID).Error; err != nil {
		return fmt.Errorf("failed to fetch ScoreItem: %w", err)
	}

	log.Printf("\n📌 SCORE ITEM DETAILS\n")
	log.Printf("Name: %s\n", item.Name)
	log.Printf("Group: %s\n", item.Subgroup.Group.Name)
	log.Printf("Subgroup: %s\n", item.Subgroup.Name)

	if item.Points != nil {
		log.Printf("Current Points: %.1f\n", *item.Points)
	} else {
		log.Printf("Current Points: 0 (nil)\n")
	}

	if item.LastReview != nil {
		log.Printf("Last Review: %s (%d days ago)\n",
			item.LastReview.Format("2006-01-02"),
			int(time.Since(*item.LastReview).Hours()/24))
	} else {
		log.Printf("Last Review: Never\n")
	}

	// 2. Fetch and analyze linked articles
	log.Printf("\n📚 LINKED ARTICLES ANALYSIS\n")

	// First, get article IDs and link metadata
	type linkInfo struct {
		ArticleID       uuid.UUID
		ConfidenceScore float64
		AutoLinked      bool
		LinkedAt        time.Time
	}

	var links []linkInfo
	err := db.Table("article_score_items").
		Select("article_id, confidence_score, auto_linked, linked_at").
		Where("score_item_id = ?", scoreItemID).
		Order("confidence_score DESC").
		Find(&links).Error

	if err != nil {
		return fmt.Errorf("failed to fetch article links: %w", err)
	}

	// Then, fetch articles
	var articles []models.Article
	if len(links) > 0 {
		articleIDs := make([]uuid.UUID, len(links))
		for i, link := range links {
			articleIDs[i] = link.ArticleID
		}

		err = db.Where("id IN ?", articleIDs).Find(&articles).Error
		if err != nil {
			return fmt.Errorf("failed to fetch articles: %w", err)
		}
	}

	// Create a map for quick lookup
	articleMap := make(map[uuid.UUID]*models.Article)
	for i := range articles {
		articleMap[articles[i].ID] = &articles[i]
	}

	log.Printf("Total linked articles: %d\n", len(links))

	qualityArticles := 0
	didacticArticles := 0

	for i, link := range links {
		article := articleMap[link.ArticleID]
		if article == nil {
			continue
		}

		hasDOI := article.DOI != nil && *article.DOI != ""
		hasPMID := article.PMID != nil && *article.PMID != ""

		articleType := "📘 Didactic"
		if hasDOI || hasPMID {
			articleType = "📄 Peer-reviewed"
			qualityArticles++
		} else {
			didacticArticles++
		}

		log.Printf("\n%d. %s %s\n", i+1, articleType, article.Title)
		log.Printf("   Journal: %s (%d)\n", article.Journal, article.PublishDate.Year())
		log.Printf("   Authors: %s\n", article.Authors)

		if hasDOI {
			log.Printf("   DOI: %s\n", *article.DOI)
		}
		if hasPMID {
			log.Printf("   PMID: %s\n", *article.PMID)
		}

		log.Printf("   Confidence: %.2f | Auto-linked: %v | Linked: %s\n",
			link.ConfidenceScore,
			link.AutoLinked,
			link.LinkedAt.Format("2006-01-02"))
	}

	log.Printf("\n📊 ARTICLE QUALITY SUMMARY\n")
	log.Printf("Peer-reviewed articles (with DOI/PMID): %d\n", qualityArticles)
	log.Printf("Didactic materials (course content): %d\n", didacticArticles)
	log.Printf("✅ Minimum requirement (7 quality articles): ")
	if qualityArticles >= 7 {
		log.Printf("MET (%d/7)\n", qualityArticles)
	} else {
		log.Printf("NOT MET (%d/7) - Need %d more\n", qualityArticles, 7-qualityArticles)
	}

	// 3. Analyze clinical fields
	log.Printf("\n🏥 CLINICAL FIELDS ANALYSIS\n")

	crLen := 0
	if item.ClinicalRelevance != nil {
		crLen = len(*item.ClinicalRelevance)
	}

	peLen := 0
	if item.PatientExplanation != nil {
		peLen = len(*item.PatientExplanation)
	}

	condLen := 0
	if item.Conduct != nil {
		condLen = len(*item.Conduct)
	}

	log.Printf("\nField lengths:\n")
	log.Printf("  Clinical Relevance: %d chars (recommended: 1000-1500)\n", crLen)
	log.Printf("  Patient Explanation: %d chars (recommended: 500-800)\n", peLen)
	log.Printf("  Conduct: %d chars (recommended: 800-1200)\n", condLen)

	// Determine enrichment tier
	tier := determineTier(&item, crLen, peLen, condLen)
	log.Printf("\n🎯 ENRICHMENT TIER: %s\n", tier)

	switch tier {
	case "PRESERVE":
		log.Println("Content is excellent and recent. No changes needed.")
	case "ENRICH":
		log.Println("Content exists but could be improved. Consider:")
		log.Println("  - Adding more quantitative data")
		log.Println("  - Including recent research findings")
		log.Println("  - Expanding clinical recommendations")
	case "REWRITE":
		log.Println("Content needs significant improvement. Consider:")
		log.Println("  - Complete rewrite based on current articles")
		log.Println("  - Adding missing sections")
		log.Println("  - Improving structure and depth")
	}

	// 4. Display current content
	log.Printf("\n📝 CURRENT CONTENT PREVIEW\n")

	if item.ClinicalRelevance != nil {
		preview := *item.ClinicalRelevance
		if len(preview) > 500 {
			preview = preview[:500] + "..."
		}
		log.Printf("\nClinical Relevance:\n%s\n", preview)
	} else {
		log.Printf("\nClinical Relevance: (empty)\n")
	}

	if item.PatientExplanation != nil {
		preview := *item.PatientExplanation
		if len(preview) > 300 {
			preview = preview[:300] + "..."
		}
		log.Printf("\nPatient Explanation:\n%s\n", preview)
	} else {
		log.Printf("\nPatient Explanation: (empty)\n")
	}

	if item.Conduct != nil {
		preview := *item.Conduct
		if len(preview) > 400 {
			preview = preview[:400] + "..."
		}
		log.Printf("\nConduct:\n%s\n", preview)
	} else {
		log.Printf("\nConduct: (empty)\n")
	}

	// 5. Recommendations
	log.Printf("\n💡 RECOMMENDATIONS\n")

	if qualityArticles < 7 {
		log.Printf("\n1. ARTICLES: Need to find %d more peer-reviewed articles\n", 7-qualityArticles)
		log.Println("   Suggested PubMed search terms:")
		log.Println("   - goal setting AND patient engagement AND (6 months OR short-term)")
		log.Println("   - treatment adherence AND health goals AND motivation")
		log.Println("   - patient-centered care AND goal-setting intervention")
		log.Println("   Filters: Review OR Meta-Analysis, Last 10 years, English")
	}

	if tier == "REWRITE" || tier == "ENRICH" {
		log.Printf("\n2. CONTENT: Clinical fields need updating\n")
		log.Println("   After gathering sufficient articles, use LLM enrichment to:")
		log.Println("   - Generate evidence-based clinical relevance")
		log.Println("   - Create patient-friendly explanations")
		log.Println("   - Develop structured clinical conduct guidelines")
	}

	currentPoints := float64(0)
	if item.Points != nil {
		currentPoints = *item.Points
	}

	if currentPoints == 0 {
		log.Printf("\n3. POINTS: Currently 0 (not scored)\n")
		log.Println("   Suggested points for 'Percepção de futuro - 6 meses':")
		log.Println("   - 5-10 points: Important for treatment planning and patient engagement")
		log.Println("   Justification: Short-term goal setting is crucial for adherence")
	}

	return nil
}

func determineTier(item *models.ScoreItem, crLen, peLen, condLen int) string {
	// Check if recently reviewed
	isRecent := item.LastReview != nil &&
		item.LastReview.After(time.Now().AddDate(0, -6, 0)) // 6 months

	// PRESERVE: excellent and recent content
	if crLen >= 1500 && peLen >= 600 && condLen >= 800 && isRecent {
		return "PRESERVE"
	}

	// REWRITE: missing or very short content
	if crLen < 500 {
		return "REWRITE"
	}

	// ENRICH: everything in between
	return "ENRICH"
}
