package main

import (
	"fmt"
	"log"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
)

func main() {
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

	// Count by review status
	var total, notReviewed, reviewed, empty int64

	db.Table("score_items").Where("deleted_at IS NULL").Count(&total)
	db.Table("score_items").Where("deleted_at IS NULL AND last_review IS NULL").Count(&notReviewed)
	db.Table("score_items").Where("deleted_at IS NULL AND last_review IS NOT NULL").Count(&reviewed)
	db.Table("score_items").
		Where("deleted_at IS NULL AND last_review IS NULL").
		Where("clinical_relevance IS NULL OR clinical_relevance = ''").
		Where("patient_explanation IS NULL OR patient_explanation = ''").
		Where("conduct IS NULL OR conduct = ''").
		Count(&empty)

	// Print report
	fmt.Println("====================================")
	fmt.Println("SCORE ITEMS REVIEW STATUS")
	fmt.Println("====================================")
	fmt.Printf("\n")
	fmt.Printf("📊 Total Items:       %d\n", total)
	fmt.Printf("✅ Reviewed:          %d (%.1f%%)\n", reviewed, float64(reviewed)/float64(total)*100)
	fmt.Printf("⏳ Not Reviewed:      %d (%.1f%%)\n", notReviewed, float64(notReviewed)/float64(total)*100)
	fmt.Printf("📝 Empty (no content): %d\n", empty)
	fmt.Printf("\n")

	// Count by tier (for not reviewed items)
	var tierRewrite, tierEnrich, tierPreserve int64

	// Tier 3: REWRITE (clinical_relevance < 500 chars)
	db.Table("score_items").
		Where("deleted_at IS NULL AND last_review IS NULL").
		Where("LENGTH(COALESCE(clinical_relevance, '')) < 500").
		Count(&tierRewrite)

	// Tier 1: PRESERVE (cr >= 1500, pe >= 600, cond >= 800)
	db.Table("score_items").
		Where("deleted_at IS NULL AND last_review IS NULL").
		Where("LENGTH(COALESCE(clinical_relevance, '')) >= 1500").
		Where("LENGTH(COALESCE(patient_explanation, '')) >= 600").
		Where("LENGTH(COALESCE(conduct, '')) >= 800").
		Count(&tierPreserve)

	// Tier 2: ENRICH (everything else)
	tierEnrich = notReviewed - tierRewrite - tierPreserve

	fmt.Println("====================================")
	fmt.Println("NOT REVIEWED BY TIER")
	fmt.Println("====================================")
	fmt.Printf("\n")
	fmt.Printf("🔴 Tier 3 (REWRITE):  %d (%.1f%%)\n", tierRewrite, float64(tierRewrite)/float64(notReviewed)*100)
	fmt.Printf("🟡 Tier 2 (ENRICH):   %d (%.1f%%)\n", tierEnrich, float64(tierEnrich)/float64(notReviewed)*100)
	fmt.Printf("🟢 Tier 1 (PRESERVE): %d (%.1f%%)\n", tierPreserve, float64(tierPreserve)/float64(notReviewed)*100)
	fmt.Printf("\n")
	fmt.Println("====================================")
	fmt.Println("NEXT STEPS")
	fmt.Println("====================================")
	fmt.Printf("To start enriching:\n")
	fmt.Printf("  go run scripts/enrich_next_item.go 0   # First item\n")
	fmt.Printf("  go run scripts/enrich_next_item.go 1   # Second item\n")
	fmt.Printf("  go run scripts/enrich_next_item.go 2   # Third item (etc)\n")
	fmt.Printf("\n")
}
