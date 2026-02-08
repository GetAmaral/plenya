package main

// Este arquivo demonstra como integrar a lógica de atualização de age range
// ao sistema de seed do Plenya EMR.
//
// OPÇÃO 1: Adicionar ao arquivo apps/api/internal/database/database.go
// OPÇÃO 2: Criar novo arquivo apps/api/internal/database/seed.go

/*
import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/plenya/api/internal/models"
	"gorm.io/gorm"
)

// AgeRangeResult representa o resultado da detecção
type AgeRangeResult struct {
	Min      *int
	Max      *int
	Detected bool
}

// SeedScoreItemAgeRange atualiza os campos age_range_min e age_range_max
// de todos score_items baseado no nome
// Pode ser chamado após AutoMigrate() no main.go
func SeedScoreItemAgeRange(db *gorm.DB) error {
	log.Println("🔄 Seeding score_item age range fields...")

	var items []models.ScoreItem
	if err := db.Find(&items).Error; err != nil {
		return err
	}

	if len(items) == 0 {
		log.Println("⚠️  No score items found, skipping age range seed")
		return nil
	}

	updated := 0
	errors := 0

	for _, item := range items {
		ageRange := detectAgeRangeFromName(item.Name)

		// Só atualizar se detectou algo
		if !ageRange.Detected {
			continue
		}

		// Verificar se já tem os valores corretos
		currentMin := item.AgeRangeMin
		currentMax := item.AgeRangeMax

		if (currentMin == nil && ageRange.Min == nil ||
		    currentMin != nil && ageRange.Min != nil && *currentMin == *ageRange.Min) &&
		   (currentMax == nil && ageRange.Max == nil ||
		    currentMax != nil && ageRange.Max != nil && *currentMax == *ageRange.Max) {
			continue
		}

		// Atualizar
		updates := map[string]interface{}{
			"age_range_min": ageRange.Min,
			"age_range_max": ageRange.Max,
		}

		if err := db.Model(&item).Updates(updates).Error; err != nil {
			log.Printf("⚠️  Failed to update item %s: %v", item.ID, err)
			errors++
			continue
		}
		updated++
	}

	log.Printf("✅ Age range seed completed: %d updated, %d errors, %d total",
		updated, errors, len(items))

	return nil
}

// detectAgeRangeFromName analisa o nome e retorna a faixa etária detectada
func detectAgeRangeFromName(name string) AgeRangeResult {
	nameLower := strings.ToLower(name)

	// Termos a ignorar
	ignoreTerms := []string{
		"pré-menopausa", "pós-menopausa", "gestação", "gestante",
		"lactante", "fase folicular", "fase lútea",
	}
	for _, term := range ignoreTerms {
		if strings.Contains(nameLower, term) {
			return AgeRangeResult{Detected: false}
		}
	}

	// Padrões compilados
	rangePattern := regexp.MustCompile(`(?i)(\d+)\s*[-a]\s*(\d+)\s*anos?`)
	lessThanPattern := regexp.MustCompile(`(?i)(?:<\s*|menor\s+(?:que|de)\s+|abaixo\s+de\s+)(\d+)\s*anos?`)
	greaterThanPattern := regexp.MustCompile(`(?i)(?:>\s*|maior\s+(?:que|de)\s+|acima\s+de\s+)(\d+)\s*anos?`)
	plusPattern := regexp.MustCompile(`(?i)(\d+)\s*\+\s*anos?`)

	// 1. Range "X-Y anos"
	if matches := rangePattern.FindStringSubmatch(name); matches != nil {
		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		if min < max && min >= 0 && max <= 150 {
			return AgeRangeResult{Min: &min, Max: &max, Detected: true}
		}
	}

	// 2. "< X anos"
	if matches := lessThanPattern.FindStringSubmatch(name); matches != nil {
		max, _ := strconv.Atoi(matches[1])
		if max > 0 && max <= 150 {
			min := 0
			return AgeRangeResult{Min: &min, Max: &max, Detected: true}
		}
	}

	// 3. "> X anos"
	if matches := greaterThanPattern.FindStringSubmatch(name); matches != nil {
		min, _ := strconv.Atoi(matches[1])
		if min >= 0 && min < 150 {
			max := 150 // Limite máximo do banco (constraint)
			return AgeRangeResult{Min: &min, Max: &max, Detected: true}
		}
	}

	// 4. "X+ anos"
	if matches := plusPattern.FindStringSubmatch(name); matches != nil {
		min, _ := strconv.Atoi(matches[1])
		if min >= 0 && min < 150 {
			max := 150 // Limite máximo do banco (constraint)
			return AgeRangeResult{Min: &min, Max: &max, Detected: true}
		}
	}

	return AgeRangeResult{Detected: false}
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
		// Seed age range após migrations
		if err := database.SeedScoreItemAgeRange(database.DB); err != nil {
			log.Printf("⚠️  Failed to seed age range: %v", err)
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
-- apps/api/database/migrations/20260208000002_update_score_item_age_range.sql

-- Function para detectar age range (similar ao Go)
CREATE OR REPLACE FUNCTION detect_age_range(item_name TEXT)
RETURNS TABLE(min_age INTEGER, max_age INTEGER) AS $$
DECLARE
    ignore_terms TEXT[] := ARRAY[
        'pré-menopausa', 'pós-menopausa', 'gestação',
        'gestante', 'lactante', 'fase folicular'
    ];
    term TEXT;
BEGIN
    -- Verificar termos a ignorar
    FOREACH term IN ARRAY ignore_terms
    LOOP
        IF LOWER(item_name) LIKE '%' || term || '%' THEN
            RETURN;
        END IF;
    END LOOP;

    -- Padrão: "X-Y anos" ou "X a Y anos"
    IF LOWER(item_name) ~ '(\d+)\s*[-a]\s*(\d+)\s*anos?' THEN
        SELECT
            (regexp_matches(LOWER(item_name), '(\d+)\s*[-a]\s*(\d+)\s*anos?'))[1]::INTEGER,
            (regexp_matches(LOWER(item_name), '(\d+)\s*[-a]\s*(\d+)\s*anos?'))[2]::INTEGER
        INTO min_age, max_age;

        IF min_age < max_age AND min_age >= 0 AND max_age <= 150 THEN
            RETURN NEXT;
            RETURN;
        END IF;
    END IF;

    -- Padrão: "< X anos"
    IF LOWER(item_name) ~ '(<\s*|menor\s+(que|de)\s+|abaixo\s+de\s+)(\d+)\s*anos?' THEN
        SELECT
            0,
            (regexp_matches(LOWER(item_name), '(<\s*|menor\s+(que|de)\s+|abaixo\s+de\s+)(\d+)\s*anos?'))[3]::INTEGER
        INTO min_age, max_age;

        IF max_age > 0 AND max_age <= 150 THEN
            RETURN NEXT;
            RETURN;
        END IF;
    END IF;

    -- Padrão: "> X anos"
    IF LOWER(item_name) ~ '(>\s*|maior\s+(que|de)\s+|acima\s+de\s+)(\d+)\s*anos?' THEN
        SELECT
            (regexp_matches(LOWER(item_name), '(>\s*|maior\s+(que|de)\s+|acima\s+de\s+)(\d+)\s*anos?'))[3]::INTEGER,
            150 -- Limite máximo do banco (constraint)
        INTO min_age, max_age;

        IF min_age >= 0 AND min_age < 150 THEN
            RETURN NEXT;
            RETURN;
        END IF;
    END IF;

    -- Padrão: "X+ anos"
    IF LOWER(item_name) ~ '(\d+)\s*\+\s*anos?' THEN
        SELECT
            (regexp_matches(LOWER(item_name), '(\d+)\s*\+\s*anos?'))[1]::INTEGER,
            150 -- Limite máximo do banco (constraint)
        INTO min_age, max_age;

        IF min_age >= 0 AND min_age < 150 THEN
            RETURN NEXT;
            RETURN;
        END IF;
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Atualizar todos os items
UPDATE score_items si
SET
    age_range_min = ar.min_age,
    age_range_max = ar.max_age
FROM detect_age_range(si.name) ar
WHERE si.deleted_at IS NULL;

-- Verificar resultados
SELECT
    CASE
        WHEN age_range_min IS NULL AND age_range_max IS NULL THEN 'No range'
        WHEN age_range_max = 999 THEN age_range_min || '+'
        ELSE age_range_min || '-' || age_range_max
    END as range,
    COUNT(*) as count
FROM score_items
WHERE deleted_at IS NULL
GROUP BY range
ORDER BY age_range_min NULLS LAST;

-- Cleanup da função temporária (opcional)
DROP FUNCTION IF EXISTS detect_age_range(TEXT);
*/

// ============================================================================
// EXEMPLO DE USO COMO HOOK NO MODEL
// ============================================================================

/*
// apps/api/internal/models/score_item.go

import (
	"regexp"
	"strconv"
	"strings"
)

// BeforeSave hook para auto-detectar age range
func (si *ScoreItem) BeforeSave(tx *gorm.DB) error {
	// Auto-detect age range se o nome foi alterado
	// OU se os campos de age range estão vazios
	if tx.Statement.Changed("Name") ||
	   (si.AgeRangeMin == nil && si.AgeRangeMax == nil) {
		ageRange := detectAgeRangeFromName(si.Name)
		if ageRange.Detected {
			si.AgeRangeMin = ageRange.Min
			si.AgeRangeMax = ageRange.Max
		}
	}
	return nil
}

type ageRangeDetection struct {
	Min      *int
	Max      *int
	Detected bool
}

func detectAgeRangeFromName(name string) ageRangeDetection {
	nameLower := strings.ToLower(name)

	// Termos a ignorar
	ignoreTerms := []string{
		"pré-menopausa", "pós-menopausa", "gestação",
	}
	for _, term := range ignoreTerms {
		if strings.Contains(nameLower, term) {
			return ageRangeDetection{Detected: false}
		}
	}

	// Padrão: "X-Y anos"
	rangePattern := regexp.MustCompile(`(?i)(\d+)\s*[-a]\s*(\d+)\s*anos?`)
	if matches := rangePattern.FindStringSubmatch(name); matches != nil {
		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		if min < max && min >= 0 && max <= 150 {
			return ageRangeDetection{Min: &min, Max: &max, Detected: true}
		}
	}

	// Padrão: "< X anos"
	lessThanPattern := regexp.MustCompile(`(?i)(?:<\s*|menor\s+que\s+)(\d+)\s*anos?`)
	if matches := lessThanPattern.FindStringSubmatch(name); matches != nil {
		max, _ := strconv.Atoi(matches[1])
		if max > 0 && max <= 150 {
			min := 0
			return ageRangeDetection{Min: &min, Max: &max, Detected: true}
		}
	}

	// Padrão: "> X anos"
	greaterThanPattern := regexp.MustCompile(`(?i)(?:>\s*|maior\s+que\s+)(\d+)\s*anos?`)
	if matches := greaterThanPattern.FindStringSubmatch(name); matches != nil {
		min, _ := strconv.Atoi(matches[1])
		if min >= 0 && min < 150 {
			max := 150 // Limite máximo do banco (constraint)
			return ageRangeDetection{Min: &min, Max: &max, Detected: true}
		}
	}

	// Padrão: "X+"
	plusPattern := regexp.MustCompile(`(?i)(\d+)\s*\+\s*anos?`)
	if matches := plusPattern.FindStringSubmatch(name); matches != nil {
		min, _ := strconv.Atoi(matches[1])
		if min >= 0 && min < 150 {
			max := 150 // Limite máximo do banco (constraint)
			return ageRangeDetection{Min: &min, Max: &max, Detected: true}
		}
	}

	return ageRangeDetection{Detected: false}
}
*/

// ============================================================================
// EXEMPLO DE USO COMO SCHEDULED JOB
// ============================================================================

/*
// apps/api/internal/scheduler/age_range_update_job.go

package scheduler

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/plenya/api/internal/models"
	"gorm.io/gorm"
)

type AgeRangeUpdateJob struct {
	db     *gorm.DB
	ticker *time.Ticker
}

func NewAgeRangeUpdateJob(db *gorm.DB) *AgeRangeUpdateJob {
	return &AgeRangeUpdateJob{
		db:     db,
		ticker: time.NewTicker(24 * time.Hour), // Rodar a cada 24h
	}
}

func (j *AgeRangeUpdateJob) Start() {
	log.Println("🤖 Age range update job started (every 24h)")

	go func() {
		// Executar imediatamente
		j.Run()

		// Depois executar periodicamente
		for range j.ticker.C {
			j.Run()
		}
	}()
}

func (j *AgeRangeUpdateJob) Stop() {
	j.ticker.Stop()
	log.Println("🛑 Age range update job stopped")
}

func (j *AgeRangeUpdateJob) Run() {
	log.Println("🔄 Running scheduled age range update...")

	var items []models.ScoreItem
	if err := j.db.Find(&items).Error; err != nil {
		log.Printf("⚠️  Failed to fetch items: %v", err)
		return
	}

	updated := 0
	for _, item := range items {
		ageRange := detectAgeRange(item.Name)

		if !ageRange.Detected {
			continue
		}

		// Verificar se já está correto
		currentMin := item.AgeRangeMin
		currentMax := item.AgeRangeMax

		if (currentMin == nil && ageRange.Min == nil ||
		    currentMin != nil && ageRange.Min != nil && *currentMin == *ageRange.Min) &&
		   (currentMax == nil && ageRange.Max == nil ||
		    currentMax != nil && ageRange.Max != nil && *currentMax == *ageRange.Max) {
			continue
		}

		// Atualizar
		updates := map[string]interface{}{
			"age_range_min": ageRange.Min,
			"age_range_max": ageRange.Max,
		}

		if err := j.db.Model(&item).Updates(updates).Error; err != nil {
			continue
		}
		updated++
	}

	log.Printf("✅ Age range update completed: %d items updated", updated)
}

type ageRangeResult struct {
	Min      *int
	Max      *int
	Detected bool
}

func detectAgeRange(name string) ageRangeResult {
	nameLower := strings.ToLower(name)

	// Ignorar termos específicos
	ignoreTerms := []string{"pré-menopausa", "pós-menopausa", "gestação"}
	for _, term := range ignoreTerms {
		if strings.Contains(nameLower, term) {
			return ageRangeResult{Detected: false}
		}
	}

	// Detectar padrões (implementação simplificada)
	rangePattern := regexp.MustCompile(`(?i)(\d+)\s*[-a]\s*(\d+)\s*anos?`)
	if matches := rangePattern.FindStringSubmatch(name); matches != nil {
		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		if min < max && min >= 0 && max <= 150 {
			return ageRangeResult{Min: &min, Max: &max, Detected: true}
		}
	}

	return ageRangeResult{Detected: false}
}

// No main.go:
//
// ageRangeJob := scheduler.NewAgeRangeUpdateJob(database.DB)
// ageRangeJob.Start()
// defer ageRangeJob.Stop()
*/

// ============================================================================
// EXEMPLO DE USO COMO API ENDPOINT
// ============================================================================

/*
// apps/api/internal/handlers/admin_handler.go

func (h *AdminHandler) UpdateScoreItemAgeRange(c *fiber.Ctx) error {
	log.Println("🔄 Admin triggered age range update...")

	var items []models.ScoreItem
	if err := h.db.Find(&items).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch score items",
		})
	}

	updated := 0
	errors := 0
	unchanged := 0
	ignored := 0

	for _, item := range items {
		ageRange := detectAgeRangeFromName(item.Name)

		if !ageRange.Detected {
			ignored++
			continue
		}

		currentMin := item.AgeRangeMin
		currentMax := item.AgeRangeMax

		if (currentMin == nil && ageRange.Min == nil ||
		    currentMin != nil && ageRange.Min != nil && *currentMin == *ageRange.Min) &&
		   (currentMax == nil && ageRange.Max == nil ||
		    currentMax != nil && ageRange.Max != nil && *currentMax == *ageRange.Max) {
			unchanged++
			continue
		}

		updates := map[string]interface{}{
			"age_range_min": ageRange.Min,
			"age_range_max": ageRange.Max,
		}

		if err := h.db.Model(&item).Updates(updates).Error; err != nil {
			errors++
			continue
		}
		updated++
	}

	return c.JSON(fiber.Map{
		"message":   "Age range update completed",
		"updated":   updated,
		"unchanged": unchanged,
		"ignored":   ignored,
		"errors":    errors,
		"total":     len(items),
	})
}

// No setup de rotas (main.go):
//
// admin := v1.Group("/admin")
// admin.Use(middleware.Auth(cfg))
// admin.Use(middleware.RequireAdmin())
// admin.Post("/score-items/update-age-range", adminHandler.UpdateScoreItemAgeRange)
*/
