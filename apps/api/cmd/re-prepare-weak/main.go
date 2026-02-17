package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Re-Prepare Weak - Re-prepara ScoreItems com preparations fracas (< 15 chunks)
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

	fmt.Println("🔬 Re-preparing weak preparations (< 15 chunks)...")

	// Buscar preparations fracas
	var weakPreps []struct {
		ScoreItemID uuid.UUID
		Name        string
		TotalChunks int
		AvgSim      float64
	}

	db.Raw(`
		SELECT
			prep.score_item_id,
			si.name,
			(prep.metadata->>'total_chunks')::int as total_chunks,
			(prep.metadata->>'avg_similarity')::float as avg_sim
		FROM score_item_enrichment_preparation prep
		JOIN score_items si ON si.id = prep.score_item_id
		WHERE (prep.metadata->>'total_chunks')::int < 15
		ORDER BY total_chunks ASC, avg_sim DESC
	`).Scan(&weakPreps)

	fmt.Printf("📋 Found %d weak preparations to re-process\n\n", len(weakPreps))

	if len(weakPreps) == 0 {
		fmt.Println("✅ No weak preparations found!")
		return
	}

	improved := 0
	same := 0
	failed := 0

	for i, prep := range weakPreps {
		oldChunks := prep.TotalChunks

		// Delete old preparation
		db.Exec("DELETE FROM score_item_enrichment_preparation WHERE score_item_id = ?", prep.ScoreItemID)

		// Re-prepare with adaptive threshold
		newPrep, err := prepService.PrepareChunksAdaptive(prep.ScoreItemID)
		if err != nil {
			failed++
			if i < 10 {
				fmt.Printf("  ❌ %s: failed - %v\n", prep.Name, err)
			}
			continue
		}

		// Safe type conversions from JSONB
		newChunks := 0
		if val, ok := newPrep.Metadata["total_chunks"]; ok {
			switch v := val.(type) {
			case float64:
				newChunks = int(v)
			case int:
				newChunks = v
			}
		}

		newThreshold := 0.0
		if val, ok := newPrep.Metadata["threshold_used"]; ok {
			if f, ok := val.(float64); ok {
				newThreshold = f
			}
		}

		newGrade := "unknown"
		if val, ok := newPrep.Metadata["quality_grade"]; ok {
			if s, ok := val.(string); ok {
				newGrade = s
			}
		}

		if newChunks > oldChunks {
			improved++
			if i < 20 {
				fmt.Printf("  ✅ %s: %d → %d chunks (threshold %.2f, grade: %s)\n",
					prep.Name, oldChunks, newChunks, newThreshold, newGrade)
			}
		} else {
			same++
			if i < 10 {
				fmt.Printf("  ⚠️  %s: still %d chunks (no improvement)\n", prep.Name, newChunks)
			}
		}

		if (i+1)%50 == 0 {
			fmt.Printf("\n  Progress: %d/%d...\n\n", i+1, len(weakPreps))
		}
	}

	fmt.Printf("\n✅ Re-preparation completed!\n")
	fmt.Printf("   Improved: %d\n", improved)
	fmt.Printf("   Same: %d\n", same)
	fmt.Printf("   Failed: %d\n", failed)

	// Estatísticas finais
	var stats struct {
		AvgChunks   float64
		WeakCount   int
		ExcellCount int
	}

	db.Raw(`
		SELECT
			ROUND(AVG((metadata->>'total_chunks')::int), 1) as avg_chunks,
			COUNT(*) FILTER (WHERE (metadata->>'total_chunks')::int < 10) as weak_count,
			COUNT(*) FILTER (WHERE metadata->>'quality_grade' = 'excellent') as excell_count
		FROM score_item_enrichment_preparation
	`).Scan(&stats)

	fmt.Printf("\n📊 Final stats:\n")
	fmt.Printf("   Avg chunks: %.1f\n", stats.AvgChunks)
	fmt.Printf("   Weak (< 10): %d\n", stats.WeakCount)
	fmt.Printf("   Excellent: %d\n", stats.ExcellCount)
}
