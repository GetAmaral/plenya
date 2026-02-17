package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
	"gorm.io/gorm"
)

// EnrichmentContext contexto para enrichment
type EnrichmentContext struct {
	ScoreItem ScoreItemData `json:"score_item"`
	Articles  []ArticleData `json:"articles"`
	Tier      string        `json:"tier"`
	Status    string        `json:"status"`
}

type ScoreItemData struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	Unit               *string `json:"unit"`
	Gender             *string `json:"gender"`
	AgeRangeMin        *int    `json:"age_range_min"`
	AgeRangeMax        *int    `json:"age_range_max"`
	Points             *float64 `json:"points"`
	ClinicalRelevance  *string `json:"clinical_relevance"`
	PatientExplanation *string `json:"patient_explanation"`
	Conduct            *string `json:"conduct"`
	LastReview         *time.Time `json:"last_review"`
}

type ArticleData struct {
	Title      string  `json:"title"`
	Authors    string  `json:"authors"`
	Journal    string  `json:"journal"`
	Year       int     `json:"year"`
	Abstract   *string `json:"abstract"`
	DOI        *string `json:"doi"`
	PMID       *string `json:"pmid"`
	Similarity float64 `json:"similarity"`
}

// EnrichmentPayload payload to apply enrichment
type EnrichmentPayload struct {
	ScoreItemID        string  `json:"score_item_id"`
	ClinicalRelevance  string  `json:"clinical_relevance"`
	PatientExplanation string  `json:"patient_explanation"`
	Conduct            string  `json:"conduct"`
	MaxPoints          int     `json:"max_points"`
	Justification      string  `json:"justification"`
	Confidence         float64 `json:"confidence"`
	Tier               string  `json:"tier"`
}

func main() {
	// Parse command line args
	var offset int = 0
	var applyMode bool = false
	var payloadFile string

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "apply":
			if len(os.Args) < 3 {
				log.Fatal("Usage: go run scripts/enrich_next_item.go apply <enrichment.json>")
			}
			applyMode = true
			payloadFile = os.Args[2]
		default:
			// Assume it's an offset number
			_, err := fmt.Sscanf(os.Args[1], "%d", &offset)
			if err != nil {
				log.Fatalf("Invalid offset: %v. Usage: go run scripts/enrich_next_item.go [offset]", err)
			}
		}
	}

	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	if err = database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db := database.DB

	if applyMode {
		// Apply enrichment mode
		applyEnrichment(db, payloadFile)
	} else {
		// Print context mode
		printEnrichmentContext(db, offset)
	}
}

func printEnrichmentContext(db *gorm.DB, offset int) {
	// Query: last_review IS NULL LIMIT 1 OFFSET {offset}
	var item models.ScoreItem
	err := db.Where("last_review IS NULL").
		Order("created_at ASC").
		Offset(offset).
		Limit(1).
		First(&item).Error

	if err != nil {
		log.Fatalf("Failed to find ScoreItem: %v (offset=%d)", err, offset)
	}

	// Fetch linked articles (top 10 by similarity)
	vectorRepo := repository.NewArticleVectorRepository(db)
	articleResults, err := vectorRepo.FindSimilarArticlesForScoreItem(item.ID, 10, 0.7)
	if err != nil {
		log.Printf("Warning: RAG search failed: %v", err)
		articleResults = []repository.ArticleSearchResult{}
	}

	// Determine tier
	tier := determineTier(&item)
	status := determineStatus(&item)

	// Build context
	context := EnrichmentContext{
		ScoreItem: ScoreItemData{
			ID:                 item.ID.String(),
			Name:               item.Name,
			Unit:               item.Unit,
			Gender:             item.Gender,
			AgeRangeMin:        item.AgeRangeMin,
			AgeRangeMax:        item.AgeRangeMax,
			Points:             item.Points,
			ClinicalRelevance:  item.ClinicalRelevance,
			PatientExplanation: item.PatientExplanation,
			Conduct:            item.Conduct,
			LastReview:         item.LastReview,
		},
		Articles: make([]ArticleData, 0, len(articleResults)),
		Tier:     tier,
		Status:   status,
	}

	for _, result := range articleResults {
		context.Articles = append(context.Articles, ArticleData{
			Title:      result.Article.Title,
			Authors:    result.Article.Authors,
			Journal:    result.Article.Journal,
			Year:       result.Article.PublishDate.Year(),
			Abstract:   result.Article.Abstract,
			DOI:        result.Article.DOI,
			PMID:       result.Article.PMID,
			Similarity: result.Similarity,
		})
	}

	// Print report header
	fmt.Printf("====================================\n")
	fmt.Printf("ENRICHMENT CONTEXT (OFFSET=%d)\n", offset)
	fmt.Printf("====================================\n\n")
	fmt.Printf("📋 Name:   %s\n", item.Name)
	fmt.Printf("🎯 Tier:   %s\n", tier)
	fmt.Printf("📊 Status: %s\n", status)
	fmt.Printf("🔗 ID:     %s\n", item.ID.String())
	fmt.Printf("📚 Articles: %d linked\n\n", len(articleResults))

	// Output full JSON
	jsonData, err := json.MarshalIndent(context, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	fmt.Println(string(jsonData))
	fmt.Println()
	fmt.Printf("====================================\n")
	fmt.Printf("NEXT STEPS:\n")
	fmt.Printf("====================================\n")
	fmt.Printf("1. Copy the JSON above\n")
	fmt.Printf("2. Enrich using Claude or other tools\n")
	fmt.Printf("3. Save enriched data to enrichment.json\n")
	fmt.Printf("4. Apply: go run scripts/enrich_next_item.go apply enrichment.json\n\n")
}

func applyEnrichment(db *gorm.DB, payloadFile string) {
	// Read JSON file
	data, err := os.ReadFile(payloadFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	// Parse JSON
	var enrichment EnrichmentPayload
	if err := json.Unmarshal(data, &enrichment); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// Validate
	if err := validateEnrichment(&enrichment); err != nil {
		log.Fatalf("Validation failed: %v", err)
	}

	// Parse UUID
	itemID, err := uuid.Parse(enrichment.ScoreItemID)
	if err != nil {
		log.Fatalf("Invalid score_item_id: %v", err)
	}

	// Fetch current item for reporting
	var beforeItem models.ScoreItem
	if err := db.First(&beforeItem, itemID).Error; err != nil {
		log.Fatalf("Failed to find ScoreItem: %v", err)
	}

	tierBefore := determineTier(&beforeItem)
	statusBefore := determineStatus(&beforeItem)

	// Create before snapshot for audit
	beforeData := map[string]interface{}{
		"clinical_relevance":  beforeItem.ClinicalRelevance,
		"patient_explanation": beforeItem.PatientExplanation,
		"conduct":             beforeItem.Conduct,
		"points":              beforeItem.Points,
	}
	beforeJSON, _ := json.Marshal(beforeData)

	// Apply enrichment using GORM model (triggers BeforeUpdate hook)
	// The hook will auto-set LastReview when clinical fields change
	updateItem := models.ScoreItem{
		ID:                 itemID,
		ClinicalRelevance:  &enrichment.ClinicalRelevance,
		PatientExplanation: &enrichment.PatientExplanation,
		Conduct:            &enrichment.Conduct,
		Points:             floatPtr(float64(enrichment.MaxPoints)),
	}

	// Use Select to only update specific fields (triggers hook)
	err = db.Model(&models.ScoreItem{}).
		Where("id = ?", itemID).
		Select("ClinicalRelevance", "PatientExplanation", "Conduct", "Points").
		Updates(&updateItem).Error

	if err != nil {
		log.Fatalf("Failed to update score_item: %v", err)
	}

	// Fetch updated item to verify hook triggered
	var afterItem models.ScoreItem
	if err := db.First(&afterItem, itemID).Error; err != nil {
		log.Fatalf("Failed to fetch updated ScoreItem: %v", err)
	}

	// Create after snapshot
	afterData := map[string]interface{}{
		"clinical_relevance":  enrichment.ClinicalRelevance,
		"patient_explanation": enrichment.PatientExplanation,
		"conduct":             enrichment.Conduct,
		"max_points":          enrichment.MaxPoints,
		"justification":       enrichment.Justification,
	}
	afterJSON, _ := json.Marshal(afterData)

	// Save audit trail
	now := time.Now()
	err = db.Exec(`
		INSERT INTO score_item_review_history
		(id, score_item_id, review_type, before_snapshot, after_snapshot,
		 tier, confidence_score, model_used, reviewed_at)
		VALUES (?, ?, 'zero_api_enrichment', ?::jsonb, ?::jsonb, ?, ?, 'claude-sonnet-4-5-20250929', ?)
	`,
		uuid.Must(uuid.NewV7()),
		itemID,
		string(beforeJSON),
		string(afterJSON),
		enrichment.Tier,
		enrichment.Confidence,
		now,
	).Error

	if err != nil {
		log.Printf("⚠️  Warning: Failed to save audit trail: %v", err)
	}

	tierAfter := determineTier(&afterItem)
	statusAfter := determineStatus(&afterItem)

	// Print success report
	fmt.Printf("\n")
	fmt.Printf("====================================\n")
	fmt.Printf("✅ ENRICHMENT APPLIED SUCCESSFULLY\n")
	fmt.Printf("====================================\n\n")
	fmt.Printf("📋 Name:        %s\n", afterItem.Name)
	fmt.Printf("🔗 ID:          %s\n\n", enrichment.ScoreItemID)
	fmt.Printf("🎯 Tier:        %s → %s\n", tierBefore, tierAfter)
	fmt.Printf("📊 Status:      %s → %s\n", statusBefore, statusAfter)
	fmt.Printf("💯 Points:      %.1f → %d\n", getFloat(beforeItem.Points), enrichment.MaxPoints)
	fmt.Printf("🔍 Confidence:  %.2f\n", enrichment.Confidence)
	fmt.Printf("📅 Last Review: %v\n", afterItem.LastReview.Format("2006-01-02 15:04:05"))
	fmt.Printf("\n")
	fmt.Printf("📝 Field Lengths:\n")
	fmt.Printf("   - Clinical Relevance:  %d chars\n", len(enrichment.ClinicalRelevance))
	fmt.Printf("   - Patient Explanation: %d chars\n", len(enrichment.PatientExplanation))
	fmt.Printf("   - Conduct:             %d chars\n", len(enrichment.Conduct))
	fmt.Printf("\n")

	if afterItem.LastReview == nil {
		fmt.Printf("⚠️  WARNING: BeforeUpdate hook did NOT trigger!\n")
		fmt.Printf("   LastReview is still NULL\n\n")
	} else {
		fmt.Printf("✅ BeforeUpdate hook triggered successfully\n")
		fmt.Printf("   LastReview auto-updated to: %v\n\n", afterItem.LastReview.Format("2006-01-02 15:04:05"))
	}
}

func determineTier(item *models.ScoreItem) string {
	crLen := len(stringValue(item.ClinicalRelevance))
	peLen := len(stringValue(item.PatientExplanation))
	condLen := len(stringValue(item.Conduct))

	isRecent := item.LastReview != nil &&
		item.LastReview.After(time.Now().AddDate(0, -6, 0)) // 6 months

	// Tier 1: PRESERVE
	if crLen >= 1500 && peLen >= 600 && condLen >= 800 && isRecent {
		return "preserve"
	}

	// Tier 3: REWRITE
	if crLen < 500 {
		return "rewrite"
	}

	// Tier 2: ENRICH
	return "enrich"
}

func determineStatus(item *models.ScoreItem) string {
	crLen := len(stringValue(item.ClinicalRelevance))
	peLen := len(stringValue(item.PatientExplanation))
	condLen := len(stringValue(item.Conduct))

	if item.LastReview == nil {
		if crLen == 0 && peLen == 0 && condLen == 0 {
			return "empty"
		}
		return "not_reviewed"
	}

	daysSince := int(time.Since(*item.LastReview).Hours() / 24)

	if daysSince > 180 {
		return "stale"
	}

	if crLen >= 1500 && peLen >= 600 && condLen >= 800 {
		return "complete"
	}

	return "partial"
}

func validateEnrichment(e *EnrichmentPayload) error {
	if e.ScoreItemID == "" {
		return fmt.Errorf("score_item_id is required")
	}
	if len(e.ClinicalRelevance) < 800 {
		return fmt.Errorf("clinical_relevance too short (min 800 chars)")
	}
	if len(e.PatientExplanation) < 400 {
		return fmt.Errorf("patient_explanation too short (min 400 chars)")
	}
	if len(e.Conduct) < 600 {
		return fmt.Errorf("conduct too short (min 600 chars)")
	}
	if e.MaxPoints < 0 || e.MaxPoints > 50 {
		return fmt.Errorf("max_points must be 0-50")
	}
	if e.Confidence < 0 || e.Confidence > 1 {
		return fmt.Errorf("confidence must be 0-1")
	}
	if e.Tier != "preserve" && e.Tier != "enrich" && e.Tier != "rewrite" {
		return fmt.Errorf("tier must be preserve/enrich/rewrite")
	}
	return nil
}

func stringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func floatPtr(f float64) *float64 {
	return &f
}

func getFloat(p *float64) float64 {
	if p == nil {
		return 0
	}
	return *p
}
