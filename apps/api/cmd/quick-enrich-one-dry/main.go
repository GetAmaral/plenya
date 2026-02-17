package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Quick Enrich One - DRY RUN MODE
// Shows what would be enriched without calling LLM API
//
// Usage: go run cmd/quick-enrich-one-dry/main.go [offset]

func main() {
	fmt.Println("🔍 Quick Enrich One ScoreItem - DRY RUN")
	fmt.Println("========================================\n")

	// Parse offset from args (default 0)
	offset := 0
	if len(os.Args) > 1 {
		if val, err := strconv.Atoi(os.Args[1]); err == nil {
			offset = val
		}
	}

	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️  .env file not found (using environment variables)")
	}

	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	// Connect to database
	dsn := cfg.Database.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	fmt.Println("✅ Connected to database")

	// Query: last_review IS NULL LIMIT 1 OFFSET N
	var item models.ScoreItem
	if err := db.Where("last_review IS NULL").
		Order("created_at ASC").
		Limit(1).
		Offset(offset).
		First(&item).Error; err != nil {
		log.Fatalf("❌ Failed to find item (offset=%d): %v", offset, err)
	}

	fmt.Printf("📄 Found item at offset %d:\n", offset)
	fmt.Printf("   ID: %s\n", item.ID)
	fmt.Printf("   Name: %s\n", item.Name)
	if item.Unit != nil {
		fmt.Printf("   Unit: %s\n", *item.Unit)
	}
	if item.Gender != nil {
		fmt.Printf("   Gender: %s\n", *item.Gender)
	}
	fmt.Println()

	// Create enrichment service
	enrichmentService := services.NewScoreItemEnrichmentService(db)

	// Determine tier
	tier := enrichmentService.DetermineTier(&item)

	// Count unreviewed items
	var totalUnreviewed int64
	db.Model(&models.ScoreItem{}).Where("last_review IS NULL").Count(&totalUnreviewed)

	fmt.Println("=" + string(make([]byte, 60)))
	fmt.Println("📊 ENRICHMENT ANALYSIS - DRY RUN")
	fmt.Println("=" + string(make([]byte, 60)))
	fmt.Printf("Name:               %s\n", item.Name)
	fmt.Printf("Tier:               %s\n", tier)
	fmt.Printf("Position:           %d / %d unreviewed items\n", offset+1, totalUnreviewed)
	fmt.Println()

	// Current state analysis
	fmt.Println("📋 Current State:")
	crExists := item.ClinicalRelevance != nil && *item.ClinicalRelevance != ""
	peExists := item.PatientExplanation != nil && *item.PatientExplanation != ""
	condExists := item.Conduct != nil && *item.Conduct != ""

	if crExists {
		fmt.Printf("   ✅ ClinicalRelevance:   %d chars\n", len(*item.ClinicalRelevance))
	} else {
		fmt.Println("   ❌ ClinicalRelevance:   empty")
	}
	if peExists {
		fmt.Printf("   ✅ PatientExplanation:  %d chars\n", len(*item.PatientExplanation))
	} else {
		fmt.Println("   ❌ PatientExplanation:  empty")
	}
	if condExists {
		fmt.Printf("   ✅ Conduct:             %d chars\n", len(*item.Conduct))
	} else {
		fmt.Println("   ❌ Conduct:             empty")
	}

	if item.Points != nil {
		fmt.Printf("   ✅ Points:              %.0f\n", *item.Points)
	} else {
		fmt.Println("   ❌ Points:              not set")
	}
	fmt.Println()

	// Tier explanation
	fmt.Println("🎯 Tier Classification Logic:")
	switch tier {
	case services.TierPreserve:
		fmt.Println("   Tier 1 - PRESERVE (Skip enrichment)")
		fmt.Println("   Reason: All 3 fields exist and are high quality")
		fmt.Println("   Action: Would skip (already excellent)")
	case services.TierEnrich:
		fmt.Println("   Tier 2 - ENRICH (Incremental improvement)")
		fmt.Println("   Reason: Some fields exist but need enhancement")
		fmt.Println("   Action: Would preserve good content, improve weak areas")
	case services.TierRewrite:
		fmt.Println("   Tier 3 - REWRITE (Generate from scratch)")
		fmt.Println("   Reason: Missing or very low quality content")
		fmt.Println("   Action: Would generate all 3 fields from scratch")
	}
	fmt.Println()

	// Content preview
	if tier != services.TierRewrite {
		fmt.Println("📝 Current Content Preview:")
		fmt.Println()
		if item.ClinicalRelevance != nil && *item.ClinicalRelevance != "" {
			fmt.Println("ClinicalRelevance:")
			previewText(*item.ClinicalRelevance, 200)
		}
		fmt.Println()
		if item.PatientExplanation != nil && *item.PatientExplanation != "" {
			fmt.Println("PatientExplanation:")
			previewText(*item.PatientExplanation, 200)
		}
		fmt.Println()
		if item.Conduct != nil && *item.Conduct != "" {
			fmt.Println("Conduct:")
			previewText(*item.Conduct, 200)
		}
		fmt.Println()
	}

	// Next steps
	fmt.Println("💡 Next Steps:")
	fmt.Println()
	fmt.Println("   To enrich this item with LLM (requires ANTHROPIC_API_KEY credits):")
	fmt.Printf("   → docker compose exec -T api sh -c \"cd /app && go run cmd/quick-enrich-one/main.go %d\"\n", offset)
	fmt.Println()
	fmt.Println("   To analyze next item:")
	fmt.Printf("   → docker compose exec -T api sh -c \"cd /app && go run cmd/quick-enrich-one-dry/main.go %d\"\n", offset+1)
	fmt.Println()
	fmt.Println("   To manually update via SQL (development only):")
	fmt.Printf("   → docker compose exec db psql -U plenya_user -d plenya_db -c \"UPDATE score_items SET last_review = NOW() WHERE id = '%s';\"\n", item.ID)
	fmt.Println()
}

func previewText(text string, maxLen int) {
	if len(text) <= maxLen {
		fmt.Printf("   %s\n", text)
	} else {
		fmt.Printf("   %s...\n", text[:maxLen])
	}
}
