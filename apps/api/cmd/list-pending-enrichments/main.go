package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// List Pending Enrichments - Lista score items prontos para Claude processar
// Output: JSON com todas preparations prontas
func main() {
	godotenv.Load()
	cfg, _ := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Buscar preparations prontas
	var preparations []models.ScoreItemEnrichmentPreparation
	err = db.Preload("ScoreItem.Subgroup.Group").
		Preload("ScoreItem.ParentItem").
		Where("status = ?", "ready").
		Order("created_at ASC").
		Find(&preparations).Error

	if err != nil {
		log.Fatal(err)
	}

	type PrepSummary struct {
		ID              string                 `json:"id"`
		ScoreItemID     string                 `json:"score_item_id"`
		ScoreItemName   string                 `json:"score_item_name"`
		ScoreItemFullName string               `json:"score_item_full_name"`
		Unit            string                 `json:"unit"`
		Chunks          interface{}            `json:"chunks"`
		Metadata        map[string]interface{} `json:"metadata"`
	}

	summaries := make([]PrepSummary, len(preparations))
	for i, prep := range preparations {
		item := prep.ScoreItem
		fullName := item.GetFullName()

		unitStr := "N/A"
		if item.Unit != nil {
			unitStr = *item.Unit
		}

		summaries[i] = PrepSummary{
			ID:                prep.ID.String(),
			ScoreItemID:       prep.ScoreItemID.String(),
			ScoreItemName:     item.Name,
			ScoreItemFullName: fullName,
			Unit:              unitStr,
			Chunks:            prep.SelectedChunks,
			Metadata:          prep.Metadata,
		}
	}

	output, _ := json.MarshalIndent(summaries, "", "  ")
	fmt.Println(string(output))
}
