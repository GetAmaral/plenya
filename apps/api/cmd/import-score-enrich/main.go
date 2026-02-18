package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const JSON_DIR = "/app/score-enrich/json"

type EnrichmentJSON struct {
	ScoreItemID        uuid.UUID `json:"score_item_id"`
	ClinicalRelevance  string    `json:"clinical_relevance"`
	Points             float64   `json:"points"`
	PatientExplanation string    `json:"patient_explanation"`
	Conduct            string    `json:"conduct"`
}

func main() {
	godotenv.Load()
	cfg, _ := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	// Listar todos os JSONs processados
	pattern := filepath.Join(JSON_DIR, "*__processed.json")
	files, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatal("❌ Failed to list files:", err)
	}
	sort.Strings(files)

	total := len(files)
	fmt.Printf("📂 Found %d JSON files to import\n\n", total)

	var (
		success  int
		skipped  int
		errCount int
		errors   []string
	)

	for i, file := range files {
		base := filepath.Base(file)
		prefix := strings.TrimSuffix(base, "__processed.json")

		fmt.Printf("[%d/%d] %s ... ", i+1, total, prefix)

		// Ler e parsear JSON
		data, err := os.ReadFile(file)
		if err != nil {
			msg := fmt.Sprintf("%s: failed to read: %v", base, err)
			errors = append(errors, msg)
			errCount++
			fmt.Println("❌ read error")
			continue
		}

		var enrichment EnrichmentJSON
		if err := json.Unmarshal(data, &enrichment); err != nil {
			msg := fmt.Sprintf("%s: invalid JSON: %v", base, err)
			errors = append(errors, msg)
			errCount++
			fmt.Println("❌ parse error")
			continue
		}

		// Validar campos obrigatórios
		if enrichment.ScoreItemID == uuid.Nil {
			msg := fmt.Sprintf("%s: missing score_item_id", base)
			errors = append(errors, msg)
			errCount++
			fmt.Println("❌ missing id")
			continue
		}
		if enrichment.ClinicalRelevance == "" || enrichment.PatientExplanation == "" || enrichment.Conduct == "" {
			msg := fmt.Sprintf("%s: empty required fields", base)
			errors = append(errors, msg)
			errCount++
			fmt.Println("❌ empty fields")
			continue
		}

		// Buscar ScoreItem no banco
		var scoreItem models.ScoreItem
		if err := db.First(&scoreItem, enrichment.ScoreItemID).Error; err != nil {
			msg := fmt.Sprintf("%s: ScoreItem %s not found: %v", base, enrichment.ScoreItemID, err)
			errors = append(errors, msg)
			errCount++
			fmt.Println("❌ not found in db")
			continue
		}

		// Atualizar os 4 campos (BeforeUpdate hook auto-atualiza LastReview)
		updates := map[string]interface{}{
			"clinical_relevance":  enrichment.ClinicalRelevance,
			"patient_explanation": enrichment.PatientExplanation,
			"conduct":             enrichment.Conduct,
			"points":              enrichment.Points,
		}

		if err := db.Model(&scoreItem).Updates(updates).Error; err != nil {
			msg := fmt.Sprintf("%s: DB update failed: %v", base, err)
			errors = append(errors, msg)
			errCount++
			fmt.Println("❌ db error")
			continue
		}

		success++
		fmt.Printf("✅ (cr=%d pat=%d cond=%d pts=%.0f)\n",
			len(enrichment.ClinicalRelevance),
			len(enrichment.PatientExplanation),
			len(enrichment.Conduct),
			enrichment.Points,
		)
	}

	// Resumo final
	fmt.Printf("\n══════════════════════════════════════\n")
	fmt.Printf("✅ Imported:  %d/%d\n", success, total)
	if skipped > 0 {
		fmt.Printf("⏭️  Skipped:   %d\n", skipped)
	}
	if errCount > 0 {
		fmt.Printf("❌ Errors:    %d\n", errCount)
		fmt.Printf("\nErros detalhados:\n")
		for _, e := range errors {
			fmt.Printf("  - %s\n", e)
		}
	}
	fmt.Printf("══════════════════════════════════════\n")
}
