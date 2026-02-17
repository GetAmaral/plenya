package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Auto Review - Smart router: Tier 1 → mark reviewed, Tier 2/3 → LLM enrich
// Automatically decides the best processing method based on content quality
//
// Usage: go run cmd/auto-review/main.go [offset] [--skip-llm]
//
// Examples:
//   go run cmd/auto-review/main.go 7          # Process offset 7 (auto-detect tier)
//   go run cmd/auto-review/main.go 7 --skip-llm  # Skip LLM even for Tier 2/3

func main() {
	fmt.Println("🤖 Auto Review ScoreItem")
	fmt.Println("=========================\n")

	// Parse args
	offset := 0
	skipLLM := false

	if len(os.Args) > 1 {
		if val, err := strconv.Atoi(os.Args[1]); err == nil {
			offset = val
		}
	}
	if len(os.Args) > 2 && os.Args[2] == "--skip-llm" {
		skipLLM = true
		fmt.Println("⚠️  LLM enrichment disabled (--skip-llm flag)")
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

	fmt.Printf("📄 Item at offset %d:\n", offset)
	fmt.Printf("   ID:   %s\n", item.ID)
	fmt.Printf("   Name: %s\n\n", item.Name)

	// Create enrichment service for tier detection
	enrichmentService := services.NewScoreItemEnrichmentService(db)
	tier := enrichmentService.DetermineTier(&item)

	fmt.Printf("🎯 Tier: %s\n", tier)

	// Decision tree
	switch tier {
	case services.TierPreserve:
		fmt.Println("   → Strategy: Mark as reviewed (content already excellent)\n")

		// Update last_review
		now := time.Now()
		item.LastReview = &now

		if err := db.Save(&item).Error; err != nil {
			log.Fatalf("❌ Failed to update item: %v", err)
		}

		fmt.Println("=" + string(make([]byte, 60)))
		fmt.Println("✅ REVIEW COMPLETED - TIER 1 PRESERVE")
		fmt.Println("=" + string(make([]byte, 60)))
		fmt.Printf("Name:        %s\n", item.Name)
		fmt.Printf("Method:      Timestamp update only\n")
		fmt.Printf("Last Review: %s\n", now.Format("2006-01-02 15:04:05"))
		fmt.Printf("Cost:        $0\n")

	case services.TierEnrich, services.TierRewrite:
		if skipLLM {
			fmt.Println("   → Strategy: LLM enrichment required but skipped (--skip-llm)")
			fmt.Println()
			fmt.Println("⚠️  This item needs LLM enrichment but was skipped.")
			fmt.Printf("   To process: go run cmd/auto-review/main.go %d\n", offset)
			fmt.Println("   (Remove --skip-llm flag and ensure ANTHROPIC_API_KEY has credits)")
			return
		}

		fmt.Printf("   → Strategy: LLM enrichment (%s)\n\n", tier)

		// Check API key
		apiKey := os.Getenv("ANTHROPIC_API_KEY")
		if apiKey == "" {
			log.Fatal("❌ ANTHROPIC_API_KEY required for Tier 2/3 items")
		}

		// Generate enrichment
		fmt.Println("🤖 Calling Claude API...")
		ctx := context.Background()
		result, err := enrichmentService.GenerateEnrichment(ctx, &item, tier)
		if err != nil {
			log.Fatalf("❌ Enrichment failed: %v", err)
		}

		fmt.Printf("✅ Generated (confidence: %.2f)\n", result.Confidence)

		// Validate
		validationErrors := enrichmentService.ValidateResult(result)
		if len(validationErrors) > 0 {
			fmt.Println("⚠️  Validation errors:")
			for _, verr := range validationErrors {
				fmt.Printf("   - %s\n", verr)
			}
			log.Fatal("\n❌ Validation failed")
		}

		// Save audit
		if err := enrichmentService.SaveReviewHistory(&item, result, tier); err != nil {
			log.Printf("⚠️  Failed to save audit: %v", err)
		}

		// Update item (GORM hook sets LastReview)
		item.ClinicalRelevance = &result.ClinicalRelevance
		item.PatientExplanation = &result.PatientExplanation
		item.Conduct = &result.Conduct
		points := float64(result.MaxPoints)
		item.Points = &points

		if err := db.Save(&item).Error; err != nil {
			log.Fatalf("❌ Failed to save: %v", err)
		}

		fmt.Println()
		fmt.Println("=" + string(make([]byte, 60)))
		fmt.Printf("✅ ENRICHMENT COMPLETED - %s\n", tier)
		fmt.Println("=" + string(make([]byte, 60)))
		fmt.Printf("Name:   %s\n", item.Name)
		fmt.Printf("Method: LLM enrichment\n")
		fmt.Printf("Fields: CR=%d, PE=%d, Cond=%d chars\n",
			len(result.ClinicalRelevance),
			len(result.PatientExplanation),
			len(result.Conduct))
		if tier == services.TierEnrich {
			fmt.Printf("Cost:   ~$0.015\n")
		} else {
			fmt.Printf("Cost:   ~$0.020\n")
		}
	}

	// Count remaining
	var remaining int64
	db.Model(&models.ScoreItem{}).Where("last_review IS NULL").Count(&remaining)

	fmt.Println()
	fmt.Printf("📊 Remaining unreviewed: %d\n", remaining)
	fmt.Printf("   Next: go run cmd/auto-review/main.go %d\n", offset+1)
}
