package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

// ZERO-API Enrichment Workflow
// Usage: go run enrich_by_offset.go -offset 0
// Fetches ONE score_item WHERE last_review IS NULL (paginated)
// Enriches using LLM + Automatic UPDATE (BeforeUpdate hook sets last_review)
// Generates detailed report

func main() {
	// Parse command-line flags
	offset := flag.Int("offset", 0, "Pagination offset for score_items")
	dryRun := flag.Bool("dry-run", false, "Print item details without enriching")
	flag.Parse()

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

	// ============================================================
	// STEP 1: FETCH ONE ITEM (last_review IS NULL) with OFFSET
	// ============================================================
	log.Printf("===========================================")
	log.Printf("ZERO-API Score Item Enrichment")
	log.Printf("Offset: %d | Dry-run: %v", *offset, *dryRun)
	log.Printf("===========================================\n")

	var item models.ScoreItem
	err = db.
		Preload("Subgroup.Group").
		Where("last_review IS NULL").
		Order("name ASC").
		Limit(1).
		Offset(*offset).
		First(&item).Error

	if err != nil {
		log.Fatalf("❌ Failed to fetch score item: %v", err)
	}

	// Extract hierarchy
	groupName := ""
	subgroupName := ""
	if item.Subgroup != nil {
		subgroupName = item.Subgroup.Name
		if item.Subgroup.Group != nil {
			groupName = item.Subgroup.Group.Name
		}
	}

	// ============================================================
	// STEP 2: REPORT CURRENT STATE
	// ============================================================
	fmt.Printf("\n📋 ITEM DETAILS\n")
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("ID:              %s\n", item.ID)
	fmt.Printf("Name:            %s\n", item.Name)
	fmt.Printf("Tier:            %s → %s\n", groupName, subgroupName)
	fmt.Printf("Unit:            %s\n", stringPtrValue(item.Unit))
	fmt.Printf("Gender:          %s\n", stringPtrValue(item.Gender))
	fmt.Printf("Age Range:       %s\n", formatAgeRange(item.AgeRangeMin, item.AgeRangeMax))
	fmt.Printf("Post-Menopause:  %s\n", boolPtrValue(item.PostMenopause))
	fmt.Printf("Lab Test Code:   %s\n", stringPtrValue(item.LabTestCode))
	fmt.Printf("\n📊 CURRENT ENRICHMENT STATUS\n")
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("Clinical Relevance:   %4d chars | %s\n",
		len(stringPtrValue(item.ClinicalRelevance)),
		enrichmentStatus(len(stringPtrValue(item.ClinicalRelevance)), 1000, 1500))
	fmt.Printf("Patient Explanation:  %4d chars | %s\n",
		len(stringPtrValue(item.PatientExplanation)),
		enrichmentStatus(len(stringPtrValue(item.PatientExplanation)), 500, 800))
	fmt.Printf("Conduct:              %4d chars | %s\n",
		len(stringPtrValue(item.Conduct)),
		enrichmentStatus(len(stringPtrValue(item.Conduct)), 800, 1200))
	fmt.Printf("Points:               %.1f\n", floatPtrValue(item.Points))
	fmt.Printf("Last Review:          %s\n", timeStatus(item.LastReview))

	// Count linked articles
	var articleCount int64
	db.Table("article_score_items").
		Where("score_item_id = ?", item.ID).
		Count(&articleCount)
	fmt.Printf("Linked Articles:      %d\n", articleCount)

	// Determine enrichment tier
	enrichmentService := services.NewScoreItemEnrichmentService(db)
	tier := enrichmentService.DetermineTier(&item)
	fmt.Printf("\n🎯 ENRICHMENT TIER: %s\n", string(tier))

	// Total items remaining
	var totalRemaining int64
	db.Model(&models.ScoreItem{}).
		Where("last_review IS NULL").
		Count(&totalRemaining)
	fmt.Printf("📈 Progress: %d of %d items pending review\n", *offset+1, totalRemaining)

	if *dryRun {
		fmt.Printf("\n✋ DRY-RUN MODE - No changes made\n")
		fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
		return
	}

	if tier == services.TierPreserve {
		fmt.Printf("\n✨ Content is excellent and recent - PRESERVING as-is\n")
		fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
		return
	}

	// ============================================================
	// STEP 3: ENRICH via LLM
	// ============================================================
	fmt.Printf("\n🧠 CALLING LLM (Model: claude-sonnet-4-5)\n")
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")

	startTime := time.Now()
	result, err := enrichmentService.GenerateEnrichment(ctx, &item, tier)
	if err != nil {
		log.Fatalf("❌ Enrichment failed: %v", err)
	}
	elapsed := time.Since(startTime)

	fmt.Printf("✅ LLM Response received in %.2f seconds\n", elapsed.Seconds())
	fmt.Printf("   Confidence: %.2f\n", result.Confidence)

	// ============================================================
	// STEP 4: VALIDATE RESULT
	// ============================================================
	fmt.Printf("\n🔍 VALIDATING RESULT\n")
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("Clinical Relevance:   %4d chars | %s\n",
		len(result.ClinicalRelevance),
		enrichmentStatus(len(result.ClinicalRelevance), 1000, 1500))
	fmt.Printf("Patient Explanation:  %4d chars | %s\n",
		len(result.PatientExplanation),
		enrichmentStatus(len(result.PatientExplanation), 500, 800))
	fmt.Printf("Conduct:              %4d chars | %s\n",
		len(result.Conduct),
		enrichmentStatus(len(result.Conduct), 800, 1200))
	fmt.Printf("Max Points:           %d/50\n", result.MaxPoints)
	fmt.Printf("Justification:        %s\n", truncate(result.Justification, 80))

	validationErrors := enrichmentService.ValidateResult(result)
	if len(validationErrors) > 0 {
		fmt.Printf("\n⚠️  VALIDATION WARNINGS:\n")
		for _, err := range validationErrors {
			fmt.Printf("   - %s\n", err)
		}
	} else {
		fmt.Printf("\n✅ Validation passed\n")
	}

	// ============================================================
	// STEP 5: SAVE REVIEW HISTORY (AUDIT)
	// ============================================================
	fmt.Printf("\n💾 SAVING REVIEW HISTORY\n")
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")

	err = enrichmentService.SaveReviewHistory(&item, result, tier)
	if err != nil {
		log.Fatalf("❌ Failed to save review history: %v", err)
	}
	fmt.Printf("✅ Review history saved\n")

	// ============================================================
	// STEP 6: AUTOMATIC UPDATE (BeforeUpdate hook sets last_review)
	// ============================================================
	fmt.Printf("\n💾 UPDATING SCORE ITEM (auto last_review via hook)\n")
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")

	// Note: We manually set last_review to ensure it's updated
	// even though BeforeUpdate hook will handle it if clinical fields change
	now := time.Now()
	maxPointsFloat := float64(result.MaxPoints)

	updates := map[string]interface{}{
		"clinical_relevance":  result.ClinicalRelevance,
		"patient_explanation": result.PatientExplanation,
		"conduct":             result.Conduct,
		"points":              &maxPointsFloat,
		"last_review":         now,
		"updated_at":          now,
	}

	err = db.Model(&item).Updates(updates).Error
	if err != nil {
		log.Fatalf("❌ Failed to update score item: %v", err)
	}

	// Verify update
	var updated models.ScoreItem
	db.Preload("Subgroup.Group").First(&updated, item.ID)

	fmt.Printf("✅ Update successful\n")
	fmt.Printf("   Last Review: %s\n", updated.LastReview.Format("2006-01-02 15:04:05"))
	fmt.Printf("   Points: %.1f\n", floatPtrValue(updated.Points))

	// ============================================================
	// FINAL REPORT
	// ============================================================
	fmt.Printf("\n")
	fmt.Printf("═══════════════════════════════════════════\n")
	fmt.Printf("✅ ENRICHMENT COMPLETE\n")
	fmt.Printf("═══════════════════════════════════════════\n")
	fmt.Printf("Name:            %s\n", updated.Name)
	fmt.Printf("Tier:            %s → %s\n", groupName, subgroupName)
	fmt.Printf("Status:          ENRICHED (tier=%s)\n", tier)
	fmt.Printf("Confidence:      %.2f\n", result.Confidence)
	fmt.Printf("Processing Time: %.2f seconds\n", elapsed.Seconds())
	fmt.Printf("Progress:        %d of %d items\n", *offset+1, totalRemaining)
	fmt.Printf("\n💡 Next command:\n")
	fmt.Printf("   go run enrich_by_offset.go -offset %d\n", *offset+1)
	fmt.Printf("═══════════════════════════════════════════\n")
}

// Helper functions

func stringPtrValue(s *string) string {
	if s == nil {
		return "N/A"
	}
	if *s == "" {
		return "N/A"
	}
	return *s
}

func boolPtrValue(b *bool) string {
	if b == nil {
		return "N/A"
	}
	if *b {
		return "Yes"
	}
	return "No"
}

func floatPtrValue(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

func formatAgeRange(min, max *int) string {
	if min == nil && max == nil {
		return "All ages"
	}
	if min != nil && max == nil {
		return fmt.Sprintf("≥%d years", *min)
	}
	if min == nil && max != nil {
		return fmt.Sprintf("≤%d years", *max)
	}
	return fmt.Sprintf("%d-%d years", *min, *max)
}

func timeStatus(t *time.Time) string {
	if t == nil {
		return "❌ NEVER REVIEWED"
	}
	return t.Format("2006-01-02 15:04:05")
}

func enrichmentStatus(length, minTarget, maxTarget int) string {
	if length == 0 {
		return "❌ EMPTY"
	}
	if length < minTarget/2 {
		return "⚠️  VERY SHORT"
	}
	if length < minTarget {
		return "⚠️  SHORT"
	}
	if length >= minTarget && length <= maxTarget {
		return "✅ OPTIMAL"
	}
	if length <= maxTarget*2 {
		return "✅ GOOD"
	}
	return "⚠️  VERY LONG"
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
