package main

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

func main() {
	// Parse item ID from command line or use default
	itemIDStr := "019bf31d-2ef0-7f36-aafe-bfcac20f9e46" // ApoB/ApoA1

	itemID, err := uuid.Parse(itemIDStr)
	if err != nil {
		log.Fatalf("Invalid UUID: %v", err)
	}

	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	db := database.DB
	ctx := context.Background()

	log.Printf("===========================================")
	log.Printf("ScoreItem Review - Quick Enrichment")
	log.Printf("===========================================")

	// 1. Fetch ScoreItem
	var item models.ScoreItem
	err = db.Preload("Subgroup.Group").First(&item, itemID).Error
	if err != nil {
		log.Fatalf("Failed to fetch score item: %v", err)
	}

	log.Printf("📋 Reviewing: %s", item.Name)
	log.Printf("   Current state:")
	log.Printf("   - Clinical Relevance: %d chars", len(stringVal(item.ClinicalRelevance)))
	log.Printf("   - Patient Explanation: %d chars", len(stringVal(item.PatientExplanation)))
	log.Printf("   - Conduct: %d chars", len(stringVal(item.Conduct)))
	log.Printf("   - Points: %.1f", floatVal(item.Points))
	log.Printf("   - Last Review: %v", item.LastReview)

	// 2. Count linked articles
	var articleCount int64
	db.Table("article_score_items").
		Where("score_item_id = ?", itemID).
		Count(&articleCount)

	log.Printf("📚 Linked articles: %d", articleCount)

	// 3. Determine enrichment tier
	enrichmentService := services.NewScoreItemEnrichmentService(db)
	tier := enrichmentService.DetermineTier(&item)

	log.Printf("📊 Enrichment tier: %s", tier)

	if tier == services.TierPreserve {
		log.Println("✨ Content is excellent and recent - preserving as-is")
		return
	}

	// 4. Generate enrichment
	log.Println("🧠 Calling LLM for enrichment...")
	result, err := enrichmentService.GenerateEnrichment(ctx, &item, tier)
	if err != nil {
		log.Fatalf("Enrichment failed: %v", err)
	}

	log.Printf("✅ LLM response received (confidence: %.2f)", result.Confidence)
	log.Printf("   New field lengths:")
	log.Printf("   - Clinical Relevance: %d chars", len(result.ClinicalRelevance))
	log.Printf("   - Patient Explanation: %d chars", len(result.PatientExplanation))
	log.Printf("   - Conduct: %d chars", len(result.Conduct))
	log.Printf("   - Max Points: %d", result.MaxPoints)

	// 5. Validate
	validationErrors := enrichmentService.ValidateResult(result)
	if len(validationErrors) > 0 {
		log.Printf("⚠️  Validation warnings: %v", validationErrors)
	} else {
		log.Println("✓ Validation passed")
	}

	// 6. Save review history
	log.Println("💾 Saving review history...")
	err = enrichmentService.SaveReviewHistory(&item, result, tier)
	if err != nil {
		log.Fatalf("Failed to save review history: %v", err)
	}

	// 7. Update ScoreItem
	log.Println("💾 Updating ScoreItem...")
	now := time.Now()
	updates := map[string]interface{}{
		"clinical_relevance":  result.ClinicalRelevance,
		"patient_explanation": result.PatientExplanation,
		"conduct":             result.Conduct,
		"last_review":         now,
		"updated_at":          now,
	}

	// Update points
	maxPointsFloat := float64(result.MaxPoints)
	updates["points"] = &maxPointsFloat

	err = db.Model(&item).Updates(updates).Error
	if err != nil {
		log.Fatalf("Failed to update score item: %v", err)
	}

	// 8. Verify
	var updated models.ScoreItem
	db.First(&updated, itemID)

	log.Printf("\n✅ Review completed successfully!")
	log.Printf("   Name: %s", updated.Name)
	log.Printf("   Clinical Relevance: %d chars", len(stringVal(updated.ClinicalRelevance)))
	log.Printf("   Patient Explanation: %d chars", len(stringVal(updated.PatientExplanation)))
	log.Printf("   Conduct: %d chars", len(stringVal(updated.Conduct)))
	log.Printf("   Points: %.1f", floatVal(updated.Points))
	log.Printf("   Last Review: %v", updated.LastReview)
}

func stringVal(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func floatVal(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}
