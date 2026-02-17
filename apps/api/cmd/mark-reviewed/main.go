package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Mark Reviewed - Updates last_review for Tier 1 (preserve) items
// Use this for items that already have excellent content and don't need LLM enrichment
//
// Usage: go run cmd/mark-reviewed/main.go [offset]
//
// IMPORTANT: Only use this for Tier 1 (preserve) items!
// For Tier 2 (enrich) or Tier 3 (rewrite), use quick-enrich-one instead.

func main() {
	fmt.Println("✅ Mark ScoreItem as Reviewed")
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
	fmt.Printf("   Name: %s\n\n", item.Name)

	// Check if item qualifies for simple review (Tier 1)
	crOk := item.ClinicalRelevance != nil && len(*item.ClinicalRelevance) > 200
	peOk := item.PatientExplanation != nil && len(*item.PatientExplanation) > 200
	condOk := item.Conduct != nil && len(*item.Conduct) > 200

	if !crOk || !peOk || !condOk {
		fmt.Println("⚠️  WARNING: This item does NOT qualify for simple review!")
		fmt.Println("   One or more fields are missing or too short.")
		fmt.Println()
		fmt.Printf("   ClinicalRelevance:   %v (length: %d)\n", crOk, getLength(item.ClinicalRelevance))
		fmt.Printf("   PatientExplanation:  %v (length: %d)\n", peOk, getLength(item.PatientExplanation))
		fmt.Printf("   Conduct:             %v (length: %d)\n", condOk, getLength(item.Conduct))
		fmt.Println()
		fmt.Println("❌ This item needs LLM enrichment, not simple review.")
		fmt.Printf("   Use: go run cmd/quick-enrich-one/main.go %d\n", offset)
		os.Exit(1)
	}

	fmt.Println("✅ Quality check passed - Item qualifies as Tier 1 (Preserve)")
	fmt.Printf("   ClinicalRelevance:   %d chars\n", len(*item.ClinicalRelevance))
	fmt.Printf("   PatientExplanation:  %d chars\n", len(*item.PatientExplanation))
	fmt.Printf("   Conduct:             %d chars\n\n", len(*item.Conduct))

	// Update last_review
	now := time.Now()
	item.LastReview = &now

	if err := db.Save(&item).Error; err != nil {
		log.Fatalf("❌ Failed to update item: %v", err)
	}

	fmt.Println("=" + string(make([]byte, 60)))
	fmt.Println("✅ REVIEW COMPLETED")
	fmt.Println("=" + string(make([]byte, 60)))
	fmt.Printf("Name:        %s\n", item.Name)
	fmt.Printf("Tier:        preserve (Tier 1)\n")
	fmt.Printf("Status:      ✅ SUCCESS\n")
	fmt.Printf("Last Review: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Println()

	fmt.Println("💡 Item marked as reviewed without LLM enrichment")
	fmt.Println("   (Content was already excellent)")
	fmt.Println()

	// Count remaining
	var remaining int64
	db.Model(&models.ScoreItem{}).Where("last_review IS NULL").Count(&remaining)
	fmt.Printf("📊 Remaining unreviewed items: %d\n", remaining)
	fmt.Printf("   To review next: go run cmd/mark-reviewed/main.go %d\n", offset+1)
}

func getLength(s *string) int {
	if s == nil {
		return 0
	}
	return len(*s)
}
