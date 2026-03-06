package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

// ExerciseJSON matches the actual exercises.json structure
type ExerciseJSON struct {
	ExerciseID       string   `json:"exerciseId"`
	Name             string   `json:"name"`
	GifURL           string   `json:"gifUrl"`
	TargetMuscles    []string `json:"targetMuscles"`
	BodyParts        []string `json:"bodyParts"`
	Equipments       []string `json:"equipments"`
	SecondaryMuscles []string `json:"secondaryMuscles"`
	Instructions     []string `json:"instructions"`
}

// ExercisePtJSON matches exercicios_mapeados_pt.json
type ExercisePtJSON struct {
	ID            string `json:"id"`
	NomeOriginal  string `json:"nome_original"`
	CategoriaPt   string `json:"categoria_pt"`
	CaminhoLocal  string `json:"caminho_local"`
	URLRemota     string `json:"url_remota"`
	TemMidiaLocal bool   `json:"tem_midia_local"`
}

func main() {
	exercisesFile := flag.String("exercises", "", "Path to exercises.json")
	ptFile := flag.String("pt", "", "Path to exercicios_mapeados_pt.json (optional)")
	gifDir := flag.String("gif-dir", "/uploads/exercises", "URL prefix for GIF files")
	gifFallbackPrefix := flag.String("gif-fallback", "https://static.exercisedb.dev/media", "CDN fallback URL prefix")
	dryRun := flag.Bool("dry-run", false, "Print what would be imported without inserting")
	flag.Parse()

	if *exercisesFile == "" {
		log.Fatal("--exercises flag is required (path to exercises.json)")
	}

	godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Auto-migrate Exercise table
	if err := db.AutoMigrate(&models.Exercise{}); err != nil {
		log.Fatal("Error migrating:", err)
	}

	// Load exercises.json
	exercisesData, err := os.ReadFile(*exercisesFile)
	if err != nil {
		log.Fatalf("Error reading %s: %v", *exercisesFile, err)
	}

	var exercises []ExerciseJSON
	if err := json.Unmarshal(exercisesData, &exercises); err != nil {
		log.Fatalf("Error parsing exercises.json: %v", err)
	}
	fmt.Printf("Loaded %d exercises from %s\n", len(exercises), *exercisesFile)

	// Load PT mappings (optional - provides categoria_pt for body part)
	ptMap := make(map[string]*ExercisePtJSON)
	if *ptFile != "" {
		ptData, err := os.ReadFile(*ptFile)
		if err != nil {
			log.Printf("Could not read PT file %s: %v (continuing without PT)", *ptFile, err)
		} else {
			var ptExercises []ExercisePtJSON
			if err := json.Unmarshal(ptData, &ptExercises); err != nil {
				log.Printf("Error parsing PT file: %v (continuing without PT)", err)
			} else {
				for i := range ptExercises {
					ptMap[ptExercises[i].ID] = &ptExercises[i]
				}
				fmt.Printf("Loaded %d PT mappings from %s\n", len(ptMap), *ptFile)
			}
		}
	}

	// Process and upsert
	created := 0
	errors := 0

	for _, ex := range exercises {
		if ex.ExerciseID == "" {
			log.Printf("Skipping exercise with empty ID: %s", ex.Name)
			errors++
			continue
		}

		model := models.Exercise{
			ExternalId:       ex.ExerciseID,
			Name:             ex.Name,
			BodyParts:        ex.BodyParts,
			TargetMuscles:    ex.TargetMuscles,
			Equipments:       ex.Equipments,
			SecondaryMuscles: ex.SecondaryMuscles,
			Instructions:     ex.Instructions,
			GifUrl:           fmt.Sprintf("%s/%s.gif", *gifDir, ex.ExerciseID),
			GifUrlFallback:   fmt.Sprintf("%s/%s.gif", *gifFallbackPrefix, ex.ExerciseID),
			IsActive:         true,
		}

		// Apply PT body part from mapping file
		if pt, ok := ptMap[ex.ExerciseID]; ok {
			if pt.CategoriaPt != "" {
				model.BodyPartsPt = []string{pt.CategoriaPt}
			}
		}

		if *dryRun {
			fmt.Printf("  [DRY] %s -> %s (bodyParts: %v)\n", ex.ExerciseID, ex.Name, model.BodyParts)
			continue
		}

		// Upsert by external_id using ON CONFLICT
		result := db.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "external_id"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"name",
				"body_parts", "body_parts_pt",
				"target_muscles",
				"equipments",
				"secondary_muscles",
				"instructions",
				"gif_url", "gif_url_fallback",
				"updated_at",
			}),
		}).Create(&model)

		if result.Error != nil {
			log.Printf("Error upserting %s: %v", ex.ExerciseID, result.Error)
			errors++
			continue
		}
		created++
	}

	fmt.Printf("\nDone! Upserted: %d, Errors: %d, Total: %d\n",
		created, errors, len(exercises))
}
