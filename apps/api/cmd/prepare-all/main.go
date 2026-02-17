package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
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

	// Buscar TODOS os score items
	var scoreItems []models.ScoreItem
	db.Order("(unit IS NOT NULL) DESC, created_at ASC").
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

		// Preparar chunks com threshold adaptativo e limite maior
		_, err := prepService.PrepareChunksAdaptive(item.ID)
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
