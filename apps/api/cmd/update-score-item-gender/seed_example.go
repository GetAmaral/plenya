package main

// Este arquivo demonstra como integrar a lógica de atualização de gender
// ao sistema de seed do Plenya EMR.
//
// OPÇÃO 1: Adicionar ao arquivo apps/api/internal/database/database.go
// OPÇÃO 2: Criar novo arquivo apps/api/internal/database/seed.go

/*
import (
	"log"
	"strings"

	"github.com/plenya/api/internal/models"
	"gorm.io/gorm"
)

// SeedScoreItemGender atualiza o campo gender de todos score_items baseado no nome
// Pode ser chamado após AutoMigrate() no main.go
func SeedScoreItemGender(db *gorm.DB) error {
	log.Println("🔄 Seeding score_item gender fields...")

	var items []models.ScoreItem
	if err := db.Find(&items).Error; err != nil {
		return err
	}

	if len(items) == 0 {
		log.Println("⚠️  No score items found, skipping gender seed")
		return nil
	}

	updated := 0
	errors := 0

	for _, item := range items {
		gender := detectGenderFromName(item.Name)

		// Só atualizar se for diferente do atual
		currentGender := "not_applicable"
		if item.Gender != nil {
			currentGender = *item.Gender
		}

		if currentGender != gender {
			if err := db.Model(&item).Update("gender", gender).Error; err != nil {
				log.Printf("⚠️  Failed to update item %s: %v", item.ID, err)
				errors++
				continue
			}
			updated++
		}
	}

	log.Printf("✅ Gender seed completed: %d updated, %d errors, %d total",
		updated, errors, len(items))

	return nil
}

// detectGenderFromName analisa o nome e retorna o gênero detectado
func detectGenderFromName(name string) string {
	nameLower := strings.ToLower(name)

	// Palavras-chave masculinas
	maleKeywords := []string{
		"homem", "homens", "masculino", "masculina",
		"masculinos", "masculinas", "homem adulto",
		"sexo masculino", "dos homens", "para homens",
		"em homens", "no homem", "nos homens",
	}
	for _, keyword := range maleKeywords {
		if strings.Contains(nameLower, keyword) {
			return "male"
		}
	}

	// Palavras-chave femininas
	femaleKeywords := []string{
		"mulher", "mulheres", "feminino", "feminina",
		"femininos", "femininas", "mulher adulta",
		"sexo feminino", "das mulheres", "para mulheres",
		"em mulheres", "na mulher", "nas mulheres",
	}
	for _, keyword := range femaleKeywords {
		if strings.Contains(nameLower, keyword) {
			return "female"
		}
	}

	return "not_applicable"
}
*/

// ============================================================================
// EXEMPLO DE USO NO MAIN.GO
// ============================================================================

/*
// apps/api/cmd/server/main.go

func main() {
	// ... código existente ...

	// Conectar ao banco de dados
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Auto migrate (desenvolvimento)
	if cfg.Server.Environment == "development" {
		log.Println("🔄 Running auto migrations...")
		if err := database.AutoMigrate(); err != nil {
			log.Fatalf("❌ Failed to auto migrate: %v", err)
		}
		log.Println("✅ Auto migrations completed")

		// ⭐ ADICIONAR ESTA LINHA
		// Seed gender após migrations
		if err := database.SeedScoreItemGender(database.DB); err != nil {
			log.Printf("⚠️  Failed to seed gender: %v", err)
			// Não é fatal - continua a execução
		}
	}

	// ... resto do código ...
}
*/

// ============================================================================
// EXEMPLO DE USO COMO MIGRATION SQL
// ============================================================================

/*
-- apps/api/database/migrations/20260208000001_update_score_item_gender.sql

-- Function para detectar gênero (similar ao Go)
CREATE OR REPLACE FUNCTION detect_gender(item_name TEXT)
RETURNS VARCHAR(20) AS $$
BEGIN
    -- Verificar masculino
    IF LOWER(item_name) ~ '(homem|homens|masculin)' THEN
        RETURN 'male';
    END IF;

    -- Verificar feminino
    IF LOWER(item_name) ~ '(mulher|mulheres|feminin)' THEN
        RETURN 'female';
    END IF;

    -- Default
    RETURN 'not_applicable';
END;
$$ LANGUAGE plpgsql;

-- Atualizar todos os items
UPDATE score_items
SET gender = detect_gender(name)
WHERE deleted_at IS NULL;

-- Verificar resultados
SELECT
    gender,
    COUNT(*) as count
FROM score_items
WHERE deleted_at IS NULL
GROUP BY gender;

-- Cleanup da função temporária (opcional)
DROP FUNCTION IF EXISTS detect_gender(TEXT);
*/

// ============================================================================
// EXEMPLO DE USO COMO HOOK NO MODEL
// ============================================================================

/*
// apps/api/internal/models/score_item.go

// BeforeSave hook para auto-detectar gender
func (si *ScoreItem) BeforeSave(tx *gorm.DB) error {
	// Auto-detect gender se não foi explicitamente definido
	// OU se o nome foi alterado
	if tx.Statement.Changed("Name") || si.Gender == nil || *si.Gender == "" {
		gender := detectGenderFromName(si.Name)
		si.Gender = &gender
	}
	return nil
}

func detectGenderFromName(name string) string {
	nameLower := strings.ToLower(name)

	maleKeywords := []string{"homem", "homens", "masculino"}
	for _, kw := range maleKeywords {
		if strings.Contains(nameLower, kw) {
			return "male"
		}
	}

	femaleKeywords := []string{"mulher", "mulheres", "feminino"}
	for _, kw := range femaleKeywords {
		if strings.Contains(nameLower, kw) {
			return "female"
		}
	}

	return "not_applicable"
}
*/

// ============================================================================
// EXEMPLO DE USO COMO SCHEDULED JOB
// ============================================================================

/*
// apps/api/internal/scheduler/gender_update_job.go

package scheduler

import (
	"log"
	"strings"
	"time"

	"github.com/plenya/api/internal/models"
	"gorm.io/gorm"
)

type GenderUpdateJob struct {
	db     *gorm.DB
	ticker *time.Ticker
}

func NewGenderUpdateJob(db *gorm.DB) *GenderUpdateJob {
	return &GenderUpdateJob{
		db:     db,
		ticker: time.NewTicker(24 * time.Hour), // Rodar a cada 24h
	}
}

func (j *GenderUpdateJob) Start() {
	log.Println("🤖 Gender update job started (every 24h)")

	go func() {
		// Executar imediatamente
		j.Run()

		// Depois executar periodicamente
		for range j.ticker.C {
			j.Run()
		}
	}()
}

func (j *GenderUpdateJob) Stop() {
	j.ticker.Stop()
	log.Println("🛑 Gender update job stopped")
}

func (j *GenderUpdateJob) Run() {
	log.Println("🔄 Running scheduled gender update...")

	var items []models.ScoreItem
	if err := j.db.Find(&items).Error; err != nil {
		log.Printf("⚠️  Failed to fetch items: %v", err)
		return
	}

	updated := 0
	for _, item := range items {
		gender := detectGenderFromName(item.Name)

		currentGender := "not_applicable"
		if item.Gender != nil {
			currentGender = *item.Gender
		}

		if currentGender != gender {
			if err := j.db.Model(&item).Update("gender", gender).Error; err != nil {
				continue
			}
			updated++
		}
	}

	log.Printf("✅ Gender update completed: %d items updated", updated)
}

func detectGenderFromName(name string) string {
	nameLower := strings.ToLower(name)

	maleKeywords := []string{"homem", "homens", "masculino"}
	for _, kw := range maleKeywords {
		if strings.Contains(nameLower, kw) {
			return "male"
		}
	}

	femaleKeywords := []string{"mulher", "mulheres", "feminino"}
	for _, kw := range femaleKeywords {
		if strings.Contains(nameLower, kw) {
			return "female"
		}
	}

	return "not_applicable"
}

// No main.go:
//
// genderJob := scheduler.NewGenderUpdateJob(database.DB)
// genderJob.Start()
// defer genderJob.Stop()
*/

// ============================================================================
// EXEMPLO DE USO COMO API ENDPOINT
// ============================================================================

/*
// apps/api/internal/handlers/admin_handler.go

func (h *AdminHandler) UpdateScoreItemGender(c *fiber.Ctx) error {
	log.Println("🔄 Admin triggered gender update...")

	var items []models.ScoreItem
	if err := h.db.Find(&items).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch score items",
		})
	}

	updated := 0
	errors := 0
	unchanged := 0

	for _, item := range items {
		gender := detectGenderFromName(item.Name)

		currentGender := "not_applicable"
		if item.Gender != nil {
			currentGender = *item.Gender
		}

		if currentGender == gender {
			unchanged++
			continue
		}

		if err := h.db.Model(&item).Update("gender", gender).Error; err != nil {
			errors++
			continue
		}
		updated++
	}

	return c.JSON(fiber.Map{
		"message":   "Gender update completed",
		"updated":   updated,
		"unchanged": unchanged,
		"errors":    errors,
		"total":     len(items),
	})
}

// No setup de rotas (main.go):
//
// admin := v1.Group("/admin")
// admin.Use(middleware.Auth(cfg))
// admin.Use(middleware.RequireAdmin())
// admin.Post("/score-items/update-gender", adminHandler.UpdateScoreItemGender)
*/
