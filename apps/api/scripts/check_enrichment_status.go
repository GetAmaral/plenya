package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/models"
	"gorm.io/gorm"
)

func main() {
	tier := flag.String("tier", "all", "Tier to check: letter, pillar, or all")
	flag.Parse()

	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	if *tier == "all" || *tier == "letter" {
		checkLetters(database.DB)
	}

	if *tier == "all" || *tier == "pillar" {
		checkPillars(database.DB)
	}
}

func checkLetters(db *gorm.DB) {
	var totalCount, enrichedCount, unenrichedCount int64

	// Total count
	db.Model(&models.MethodLetter{}).Count(&totalCount)

	// Enriched count
	db.Model(&models.MethodLetter{}).
		Where("last_review IS NOT NULL").
		Count(&enrichedCount)

	// Unenriched count
	db.Model(&models.MethodLetter{}).
		Where("last_review IS NULL").
		Count(&unenrichedCount)

	fmt.Println("\n=== METHOD LETTERS ===")
	fmt.Printf("Total:      %d\n", totalCount)
	fmt.Printf("Enriched:   %d (%.1f%%)\n", enrichedCount, float64(enrichedCount)/float64(totalCount)*100)
	fmt.Printf("Unenriched: %d (%.1f%%)\n", unenrichedCount, float64(unenrichedCount)/float64(totalCount)*100)

	// Show first 5 unenriched
	if unenrichedCount > 0 {
		var letters []models.MethodLetter
		db.Where("last_review IS NULL").
			Order("\"order\" ASC").
			Limit(5).
			Preload("Method").
			Find(&letters)

		fmt.Println("\nNext 5 unenriched letters:")
		for i, letter := range letters {
			fmt.Printf("  [%d] %s - %s (%s)\n", i, letter.Method.ShortName, letter.Name, letter.Code)
		}
	}
}

func checkPillars(db *gorm.DB) {
	var totalCount, enrichedCount, unenrichedCount int64

	// Total count
	db.Model(&models.MethodPillar{}).Count(&totalCount)

	// Enriched count
	db.Model(&models.MethodPillar{}).
		Where("last_review IS NOT NULL").
		Count(&enrichedCount)

	// Unenriched count
	db.Model(&models.MethodPillar{}).
		Where("last_review IS NULL").
		Count(&unenrichedCount)

	fmt.Println("\n=== METHOD PILLARS ===")
	fmt.Printf("Total:      %d\n", totalCount)
	fmt.Printf("Enriched:   %d (%.1f%%)\n", enrichedCount, float64(enrichedCount)/float64(totalCount)*100)
	fmt.Printf("Unenriched: %d (%.1f%%)\n", unenrichedCount, float64(unenrichedCount)/float64(totalCount)*100)

	// Show first 5 unenriched
	if unenrichedCount > 0 {
		var pillars []models.MethodPillar
		db.Where("last_review IS NULL").
			Order("\"order\" ASC").
			Limit(5).
			Preload("Letter").
			Preload("Letter.Method").
			Find(&pillars)

		fmt.Println("\nNext 5 unenriched pillars:")
		for i, pillar := range pillars {
			fmt.Printf("  [%d] %s > %s > %s\n", i, pillar.Letter.Method.ShortName, pillar.Letter.Name, pillar.Name)
		}
	}
}
