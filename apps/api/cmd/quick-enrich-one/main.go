package main

import (
	"context"
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

// Quick Enrich One - Enriquece um único ScoreItem com last_review IS NULL
// Usage: go run cmd/quick-enrich-one/main.go [offset]
//
// Exemplo: go run cmd/quick-enrich-one/main.go 6
// Pega o 7º item sem last_review (OFFSET 6)

func main() {
	fmt.Println("🚀 Quick Enrich One ScoreItem")
	fmt.Println("==============================\n")

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

	// Check API key
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ ANTHROPIC_API_KEY environment variable is required")
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
	fmt.Println()

	// Create enrichment service
	enrichmentService := services.NewScoreItemEnrichmentService(db)
	fmt.Println("⚙️  Using: Claude Sonnet 4.5 (high-quality)")

	// Determine tier
	tier := enrichmentService.DetermineTier(&item)
	fmt.Printf("🎯 Tier: %s\n\n", tier)

	// Show current state
	fmt.Println("📋 Current State:")
	if item.ClinicalRelevance != nil && *item.ClinicalRelevance != "" {
		fmt.Printf("   ClinicalRelevance: %d chars\n", len(*item.ClinicalRelevance))
	} else {
		fmt.Println("   ClinicalRelevance: ❌ empty")
	}
	if item.PatientExplanation != nil && *item.PatientExplanation != "" {
		fmt.Printf("   PatientExplanation: %d chars\n", len(*item.PatientExplanation))
	} else {
		fmt.Println("   PatientExplanation: ❌ empty")
	}
	if item.Conduct != nil && *item.Conduct != "" {
		fmt.Printf("   Conduct: %d chars\n", len(*item.Conduct))
	} else {
		fmt.Println("   Conduct: ❌ empty")
	}
	fmt.Println()

	// Generate enrichment
	fmt.Println("🤖 Generating enrichment with LLM...")
	ctx := context.Background()
	result, err := enrichmentService.GenerateEnrichment(ctx, &item, tier)
	if err != nil {
		log.Fatalf("❌ Enrichment failed: %v", err)
	}

	fmt.Printf("✅ Generated (confidence: %.2f)\n\n", result.Confidence)

	// Validate result
	validationErrors := enrichmentService.ValidateResult(result)
	if len(validationErrors) > 0 {
		fmt.Println("⚠️  Validation errors:")
		for _, verr := range validationErrors {
			fmt.Printf("   - %s\n", verr)
		}
		log.Fatal("\n❌ Validation failed - aborting")
	}

	fmt.Println("✅ Validation passed")

	// Save audit history
	if err := enrichmentService.SaveReviewHistory(&item, result, tier); err != nil {
		log.Printf("⚠️  Failed to save audit: %v", err)
	} else {
		fmt.Println("✅ Audit saved")
	}

	// Update item (GORM BeforeUpdate hook will set LastReview automatically)
	item.ClinicalRelevance = &result.ClinicalRelevance
	item.PatientExplanation = &result.PatientExplanation
	item.Conduct = &result.Conduct
	points := float64(result.MaxPoints)
	item.Points = &points

	if err := db.Save(&item).Error; err != nil {
		log.Fatalf("❌ Failed to save item: %v", err)
	}

	fmt.Println("✅ Database updated")
	fmt.Println()

	// Report
	fmt.Println("=" + string(make([]byte, 60)))
	fmt.Println("📊 ENRICHMENT REPORT")
	fmt.Println("=" + string(make([]byte, 60)))
	fmt.Printf("Name:   %s\n", item.Name)
	fmt.Printf("Tier:   %s\n", tier)
	fmt.Printf("Status: ✅ SUCCESS\n\n")
	fmt.Printf("Fields Updated:\n")
	fmt.Printf("   ClinicalRelevance:   %d chars\n", len(result.ClinicalRelevance))
	fmt.Printf("   PatientExplanation:  %d chars\n", len(result.PatientExplanation))
	fmt.Printf("   Conduct:             %d chars\n", len(result.Conduct))
	fmt.Printf("   Points:              %d\n", result.MaxPoints)
	fmt.Printf("   LastReview:          %s (auto-set by hook)\n", "NOW()")
	fmt.Println()

	// Preview content
	fmt.Println("📝 Generated Content Preview:")
	fmt.Println()
	fmt.Println("ClinicalRelevance:")
	previewText(result.ClinicalRelevance, 200)
	fmt.Println()
	fmt.Println("PatientExplanation:")
	previewText(result.PatientExplanation, 200)
	fmt.Println()
	fmt.Println("Conduct:")
	previewText(result.Conduct, 200)
	fmt.Println()

	fmt.Println("✅ Enrichment completed!")
	fmt.Printf("\n💡 To review next item: go run cmd/quick-enrich-one/main.go %d\n", offset+1)
}

func previewText(text string, maxLen int) {
	if len(text) <= maxLen {
		fmt.Printf("   %s\n", text)
	} else {
		fmt.Printf("   %s...\n", text[:maxLen])
	}
}
