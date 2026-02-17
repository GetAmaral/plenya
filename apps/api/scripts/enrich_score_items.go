package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/plenya/api/internal/models"
)

// Report structure for tracking enrichment progress
type EnrichmentReport struct {
	Name          string
	Tier          string // "ScoreItem" or "ScoreLevel"
	Status        string // "enriched", "skipped", "pending"
	ID            uuid.UUID
	EnrichedAt    time.Time
	EnrichedFields []string
}

var reports []EnrichmentReport

func main() {
	// Connect to database
	dsn := "host=localhost user=plenya_user password=senha_segura dbname=plenya_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("🔍 Score Item Enrichment Tool")
	fmt.Println("=" + strings.Repeat("=", 50))

	// Main menu
	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. Enrich next ScoreItem (last_review IS NULL)")
		fmt.Println("2. Enrich next ScoreLevel (last_review IS NULL)")
		fmt.Println("3. Show pending items count")
		fmt.Println("4. Show enrichment report")
		fmt.Println("5. Batch enrich (N items)")
		fmt.Println("0. Exit")
		fmt.Print("\nSelect option: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			enrichNextScoreItem(db, 0)
		case "2":
			enrichNextScoreLevel(db, 0)
		case "3":
			showPendingCounts(db)
		case "4":
			showReport()
		case "5":
			batchEnrich(db)
		case "0":
			fmt.Println("\n✅ Final Report:")
			showReport()
			return
		default:
			fmt.Println("❌ Invalid option")
		}
	}
}

// Enrich next ScoreItem with last_review IS NULL
func enrichNextScoreItem(db *gorm.DB, offset int) {
	var item models.ScoreItem

	result := db.Where("last_review IS NULL").
		Preload("Subgroup.Group").
		Order("\"order\" ASC").
		Limit(1).
		Offset(offset).
		First(&item)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("\n✅ No more ScoreItems to enrich!")
			return
		}
		log.Printf("❌ Error fetching item: %v\n", result.Error)
		return
	}

	// Display item info
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("📊 ScoreItem: %s\n", item.Name)
	fmt.Printf("   ID: %s\n", item.ID)
	if item.Subgroup != nil {
		fmt.Printf("   Subgroup: %s", item.Subgroup.Name)
		if item.Subgroup.Group != nil {
			fmt.Printf(" > Group: %s\n", item.Subgroup.Group.Name)
		} else {
			fmt.Println()
		}
	}
	if item.Unit != nil {
		fmt.Printf("   Unit: %s\n", *item.Unit)
	}
	if item.Gender != nil {
		fmt.Printf("   Gender: %s\n", *item.Gender)
	}
	if item.AgeRangeMin != nil || item.AgeRangeMax != nil {
		fmt.Printf("   Age Range: %v - %v\n", item.AgeRangeMin, item.AgeRangeMax)
	}

	// Show current enrichment state
	fmt.Println("\n📝 Current Enrichment:")
	fmt.Printf("   Clinical Relevance: %s\n", getFieldPreview(item.ClinicalRelevance))
	fmt.Printf("   Patient Explanation: %s\n", getFieldPreview(item.PatientExplanation))
	fmt.Printf("   Conduct: %s\n", getFieldPreview(item.Conduct))

	fmt.Println("\n" + strings.Repeat("=", 60))

	// Prompt for enrichment
	reader := bufio.NewReader(os.Stdin)
	enrichedFields := []string{}

	fmt.Print("\n🩺 Clinical Relevance (for professionals): ")
	clinicalRelevance, _ := reader.ReadString('\n')
	clinicalRelevance = strings.TrimSpace(clinicalRelevance)
	if clinicalRelevance != "" && clinicalRelevance != "skip" {
		item.ClinicalRelevance = &clinicalRelevance
		enrichedFields = append(enrichedFields, "ClinicalRelevance")
	}

	fmt.Print("👤 Patient Explanation (simple language): ")
	patientExplanation, _ := reader.ReadString('\n')
	patientExplanation = strings.TrimSpace(patientExplanation)
	if patientExplanation != "" && patientExplanation != "skip" {
		item.PatientExplanation = &patientExplanation
		enrichedFields = append(enrichedFields, "PatientExplanation")
	}

	fmt.Print("💊 Conduct (clinical recommendations): ")
	conduct, _ := reader.ReadString('\n')
	conduct = strings.TrimSpace(conduct)
	if conduct != "" && conduct != "skip" {
		item.Conduct = &conduct
		enrichedFields = append(enrichedFields, "Conduct")
	}

	// Check if user wants to skip
	if len(enrichedFields) == 0 {
		fmt.Println("\n⏭️  Skipped (no changes)")
		reports = append(reports, EnrichmentReport{
			Name:   item.Name,
			Tier:   "ScoreItem",
			Status: "skipped",
			ID:     item.ID,
		})
		return
	}

	// Confirm save
	fmt.Print("\n💾 Save changes? (y/n): ")
	var confirm string
	fmt.Scanln(&confirm)

	if strings.ToLower(confirm) != "y" {
		fmt.Println("❌ Cancelled")
		return
	}

	// Save (BeforeUpdate hook will set last_review automatically)
	if err := db.Save(&item).Error; err != nil {
		log.Printf("❌ Error saving: %v\n", err)
		return
	}

	fmt.Println("✅ Saved successfully!")
	reports = append(reports, EnrichmentReport{
		Name:           item.Name,
		Tier:           "ScoreItem",
		Status:         "enriched",
		ID:             item.ID,
		EnrichedAt:     time.Now(),
		EnrichedFields: enrichedFields,
	})
}

// Enrich next ScoreLevel with last_review IS NULL
func enrichNextScoreLevel(db *gorm.DB, offset int) {
	var level models.ScoreLevel

	result := db.Where("last_review IS NULL").
		Preload("Item.Subgroup.Group").
		Order("item_id ASC, level DESC").
		Limit(1).
		Offset(offset).
		First(&level)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("\n✅ No more ScoreLevels to enrich!")
			return
		}
		log.Printf("❌ Error fetching level: %v\n", result.Error)
		return
	}

	// Display level info
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("📊 ScoreLevel: %s\n", level.Name)
	fmt.Printf("   ID: %s\n", level.ID)
	fmt.Printf("   Level: %d\n", level.Level)
	fmt.Printf("   Operator: %s\n", level.Operator)
	if level.LowerLimit != nil {
		fmt.Printf("   Lower Limit: %s\n", *level.LowerLimit)
	}
	if level.UpperLimit != nil {
		fmt.Printf("   Upper Limit: %s\n", *level.UpperLimit)
	}
	if level.Item != nil {
		fmt.Printf("   Parent Item: %s\n", level.Item.Name)
		if level.Item.Subgroup != nil {
			fmt.Printf("   Subgroup: %s", level.Item.Subgroup.Name)
			if level.Item.Subgroup.Group != nil {
				fmt.Printf(" > Group: %s\n", level.Item.Subgroup.Group.Name)
			} else {
				fmt.Println()
			}
		}
	}

	// Show current enrichment state
	fmt.Println("\n📝 Current Enrichment:")
	fmt.Printf("   Clinical Relevance: %s\n", getFieldPreview(level.ClinicalRelevance))
	fmt.Printf("   Patient Explanation: %s\n", getFieldPreview(level.PatientExplanation))
	fmt.Printf("   Conduct: %s\n", getFieldPreview(level.Conduct))

	fmt.Println("\n" + strings.Repeat("=", 60))

	// Prompt for enrichment
	reader := bufio.NewReader(os.Stdin)
	enrichedFields := []string{}

	fmt.Print("\n🩺 Clinical Relevance (for professionals): ")
	clinicalRelevance, _ := reader.ReadString('\n')
	clinicalRelevance = strings.TrimSpace(clinicalRelevance)
	if clinicalRelevance != "" && clinicalRelevance != "skip" {
		level.ClinicalRelevance = &clinicalRelevance
		enrichedFields = append(enrichedFields, "ClinicalRelevance")
	}

	fmt.Print("👤 Patient Explanation (simple language): ")
	patientExplanation, _ := reader.ReadString('\n')
	patientExplanation = strings.TrimSpace(patientExplanation)
	if patientExplanation != "" && patientExplanation != "skip" {
		level.PatientExplanation = &patientExplanation
		enrichedFields = append(enrichedFields, "PatientExplanation")
	}

	fmt.Print("💊 Conduct (clinical recommendations): ")
	conduct, _ := reader.ReadString('\n')
	conduct = strings.TrimSpace(conduct)
	if conduct != "" && conduct != "skip" {
		level.Conduct = &conduct
		enrichedFields = append(enrichedFields, "Conduct")
	}

	// Check if user wants to skip
	if len(enrichedFields) == 0 {
		fmt.Println("\n⏭️  Skipped (no changes)")
		reports = append(reports, EnrichmentReport{
			Name:   level.Name,
			Tier:   "ScoreLevel",
			Status: "skipped",
			ID:     level.ID,
		})
		return
	}

	// Confirm save
	fmt.Print("\n💾 Save changes? (y/n): ")
	var confirm string
	fmt.Scanln(&confirm)

	if strings.ToLower(confirm) != "y" {
		fmt.Println("❌ Cancelled")
		return
	}

	// Save (BeforeUpdate hook will set last_review automatically)
	if err := db.Save(&level).Error; err != nil {
		log.Printf("❌ Error saving: %v\n", err)
		return
	}

	fmt.Println("✅ Saved successfully!")
	reports = append(reports, EnrichmentReport{
		Name:           level.Name,
		Tier:           "ScoreLevel",
		Status:         "enriched",
		ID:             level.ID,
		EnrichedAt:     time.Now(),
		EnrichedFields: enrichedFields,
	})
}

// Show counts of pending items
func showPendingCounts(db *gorm.DB) {
	var itemCount, levelCount int64

	db.Model(&models.ScoreItem{}).Where("last_review IS NULL").Count(&itemCount)
	db.Model(&models.ScoreLevel{}).Where("last_review IS NULL").Count(&levelCount)

	fmt.Println("\n📊 Pending Enrichment:")
	fmt.Printf("   ScoreItems: %d\n", itemCount)
	fmt.Printf("   ScoreLevels: %d\n", levelCount)
	fmt.Printf("   Total: %d\n", itemCount+levelCount)
}

// Show enrichment report
func showReport() {
	if len(reports) == 0 {
		fmt.Println("\n📋 No enrichments performed yet")
		return
	}

	fmt.Println("\n📋 Enrichment Report")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("%-40s %-15s %-10s\n", "Name", "Tier", "Status")
	fmt.Println(strings.Repeat("-", 80))

	enrichedCount := 0
	skippedCount := 0

	for _, report := range reports {
		status := report.Status
		if report.Status == "enriched" {
			status = "✅ " + status
			enrichedCount++
		} else {
			status = "⏭️  " + status
			skippedCount++
		}

		// Truncate name if too long
		name := report.Name
		if len(name) > 35 {
			name = name[:32] + "..."
		}

		fmt.Printf("%-40s %-15s %-10s\n", name, report.Tier, status)

		if report.Status == "enriched" && len(report.EnrichedFields) > 0 {
			fmt.Printf("   Fields: %s\n", strings.Join(report.EnrichedFields, ", "))
		}
	}

	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("Total: %d | Enriched: %d | Skipped: %d\n", len(reports), enrichedCount, skippedCount)
}

// Batch enrich multiple items
func batchEnrich(db *gorm.DB) {
	fmt.Print("\nHow many items to process? ")
	var n int
	fmt.Scanln(&n)

	if n <= 0 {
		fmt.Println("❌ Invalid number")
		return
	}

	fmt.Println("1. ScoreItems")
	fmt.Println("2. ScoreLevels")
	fmt.Print("Select type: ")

	var typeChoice string
	fmt.Scanln(&typeChoice)

	for i := 0; i < n; i++ {
		fmt.Printf("\n--- Processing %d/%d ---\n", i+1, n)

		switch typeChoice {
		case "1":
			enrichNextScoreItem(db, 0)
		case "2":
			enrichNextScoreLevel(db, 0)
		default:
			fmt.Println("❌ Invalid type")
			return
		}

		// Check if user wants to continue
		if i < n-1 {
			fmt.Print("\nContinue to next? (y/n): ")
			var cont string
			fmt.Scanln(&cont)
			if strings.ToLower(cont) != "y" {
				fmt.Println("⏹️  Batch enrichment stopped")
				break
			}
		}
	}
}

// Helper to preview field content
func getFieldPreview(field *string) string {
	if field == nil || *field == "" {
		return "[empty]"
	}

	preview := *field
	if len(preview) > 60 {
		preview = preview[:57] + "..."
	}

	return preview
}
