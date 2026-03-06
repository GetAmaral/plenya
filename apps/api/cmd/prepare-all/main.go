package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Prepare ALL - Prepara chunks para todos score items (sem interação)
// Usage: go run cmd/prepare-all/main.go [uuid1 uuid2 ...]
// Sem args: processa todos sem preparation. Com args: força reprocessamento dos IDs especificados.
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

	// Parse optional specific IDs from args
	var specificIDs []uuid.UUID
	for _, arg := range os.Args[1:] {
		id, err := uuid.Parse(arg)
		if err != nil {
			log.Fatalf("❌ ID inválido: %s (%v)", arg, err)
		}
		specificIDs = append(specificIDs, id)
	}

	var scoreItems []models.ScoreItem
	q := db.Order("(unit IS NOT NULL) DESC, created_at ASC")
	if len(specificIDs) > 0 {
		q = q.Where("id IN ?", specificIDs)
		fmt.Printf("🔬 Preparando %d score items específicos (forçando reprocessamento)...\n", len(specificIDs))
		// Deletar preparations existentes para forçar reprocessamento
		db.Exec("DELETE FROM score_item_enrichment_preparation WHERE score_item_id IN ?", specificIDs)
	}
	q.Find(&scoreItems)

	fmt.Printf("🔬 Preparing chunks for %d score items...\n", len(scoreItems))

	prepared := 0
	skipped := 0
	failed := 0
	var failedItems []string

	for i, item := range scoreItems {
		// Verificar se já tem preparation
		existing, _ := prepService.GetPreparation(item.ID)
		if existing != nil {
			skipped++
			continue
		}

		// Preparar chunks com threshold adaptativo
		_, err := prepService.PrepareChunksAdaptive(item.ID)
		if err != nil {
			failed++
			failedItems = append(failedItems, fmt.Sprintf("  - %s (%s)", item.Name, err.Error()))
			continue
		}

		prepared++

		if (i+1)%50 == 0 {
			fmt.Printf("  Processed %d/%d...\n", i+1, len(scoreItems))
		}
	}

	fmt.Printf("✅ Prepared: %d | Skipped (já existia): %d | Failed: %d\n", prepared, skipped, failed)

	if len(failedItems) > 0 {
		fmt.Printf("\n⚠️  Items que falharam (%d):\n", len(failedItems))
		for _, f := range failedItems {
			fmt.Println(f)
		}
	}
}
