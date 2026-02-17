package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	godotenv.Load()
	cfg, _ := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	prepService := services.NewScoreEnrichmentPreparationService(db)

	// Leptina - Mulheres
	leptinaID := uuid.MustParse("019bf31d-2ef0-71b6-a5d1-db49a4fa62fa")

	// Delete existing
	db.Exec("DELETE FROM score_item_enrichment_preparation WHERE score_item_id = ?", leptinaID)

	fmt.Println("🔬 Preparando Leptina - Mulheres com prompts completos...")

	// Prepare with adaptive threshold
	prep, err := prepService.PrepareChunksAdaptive(leptinaID)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("\n✅ Preparation criada!\n")
	fmt.Printf("   Chunks: %d\n", len(prep.SelectedChunks["items"].([]interface{})))
	fmt.Printf("   Grade: %s\n", prep.Metadata["quality_grade"])

	if prep.PromptClinicalRelevance != nil {
		fmt.Printf("   Prompt CR: %d chars\n", len(*prep.PromptClinicalRelevance))
	}
	if prep.PromptPatientExplanation != nil {
		fmt.Printf("   Prompt PE: %d chars\n", len(*prep.PromptPatientExplanation))
	}
	if prep.PromptConduct != nil {
		fmt.Printf("   Prompt Conduct: %d chars\n", len(*prep.PromptConduct))
	}
	if prep.PromptMaxPoints != nil {
		fmt.Printf("   Prompt Points: %d chars\n", len(*prep.PromptMaxPoints))
	}

	fmt.Println("\n📄 Mostrando prompts completos...\n")

	if prep.PromptClinicalRelevance != nil {
		fmt.Println("=== PROMPT 1: CLINICAL RELEVANCE ===")
		fmt.Println(*prep.PromptClinicalRelevance)
		fmt.Println("\n" + strings.Repeat("-", 80) + "\n")
	}
}
