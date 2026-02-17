package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
	"github.com/plenya/api/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Prepare ALL - Prepara chunks para todos score items (sem interação)
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

	// Buscar score items tier rewrite (sem enrichment)
	var scoreItems []models.ScoreItem
	db.Where("clinical_relevance IS NULL OR LENGTH(clinical_relevance) < 500").
		Order("(unit IS NOT NULL) DESC, created_at ASC").
		Find(&scoreItems)

	fmt.Printf("🔬 Preparing chunks for %d score items...\n", len(scoreItems))

	prepared := 0
	skipped := 0

	for i, item := range scoreItems {
		// Verificar se já tem preparation
		existing, _ := prepService.GetPreparation(item.ID)
		if existing != nil {
			skipped++
			continue
		}

		// Preparar chunks
		_, err := prepService.PrepareChunks(item.ID, 20, 0.65)
		if err != nil {
			// Skip silenciosamente se não encontrar chunks
			continue
		}

		prepared++

		if (i+1)%50 == 0 {
			fmt.Printf("  Processed %d/%d...\n", i+1, len(scoreItems))
		}
	}

	fmt.Printf("✅ Prepared: %d | Skipped: %d\n", prepared, skipped)
}
