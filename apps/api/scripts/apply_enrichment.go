package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"gorm.io/gorm"
)

// EnrichmentData dados de enrichment gerados externamente
type EnrichmentData struct {
	ScoreItemID        string  `json:"score_item_id"`
	ClinicalRelevance  string  `json:"clinical_relevance"`
	PatientExplanation string  `json:"patient_explanation"`
	Conduct            string  `json:"conduct"`
	MaxPoints          int     `json:"max_points"`
	Justification      string  `json:"justification"`
	Confidence         float64 `json:"confidence"`
	Tier               string  `json:"tier"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run scripts/apply_enrichment.go <enrichment.json>")
	}

	// Ler arquivo JSON
	jsonFile := os.Args[1]
	data, err := os.ReadFile(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	// Parse JSON
	var enrichment EnrichmentData
	if err := json.Unmarshal(data, &enrichment); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// Validar
	if err := validate(&enrichment); err != nil {
		log.Fatalf("Validation failed: %v", err)
	}

	// Conectar banco
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err = database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db := database.DB

	// Parse UUID
	itemID, err := uuid.Parse(enrichment.ScoreItemID)
	if err != nil {
		log.Fatalf("Invalid score_item_id: %v", err)
	}

	// Criar audit trail (before snapshot)
	var beforeJSON []byte
	err = db.Raw(`
		SELECT jsonb_build_object(
			'clinical_relevance', clinical_relevance,
			'patient_explanation', patient_explanation,
			'conduct', conduct,
			'points', points
		) FROM score_items WHERE id = ?
	`, itemID).Scan(&beforeJSON).Error
	if err != nil {
		log.Fatalf("Failed to create before snapshot: %v", err)
	}

	// Aplicar enrichment
	now := time.Now()
	err = db.Exec(`
		UPDATE score_items SET
			clinical_relevance = ?,
			patient_explanation = ?,
			conduct = ?,
			points = ?,
			last_review = ?,
			updated_at = ?
		WHERE id = ?
	`,
		enrichment.ClinicalRelevance,
		enrichment.PatientExplanation,
		enrichment.Conduct,
		enrichment.MaxPoints,
		now,
		now,
		itemID,
	).Error

	if err != nil {
		log.Fatalf("Failed to update score_item: %v", err)
	}

	log.Println("✅ ScoreItem updated successfully")

	// Criar after snapshot
	afterData := map[string]interface{}{
		"clinical_relevance":  enrichment.ClinicalRelevance,
		"patient_explanation": enrichment.PatientExplanation,
		"conduct":             enrichment.Conduct,
		"max_points":          enrichment.MaxPoints,
		"justification":       enrichment.Justification,
	}
	afterJSON, _ := json.Marshal(afterData)

	// Salvar audit trail
	err = db.Exec(`
		INSERT INTO score_item_review_history
		(id, score_item_id, review_type, before_snapshot, after_snapshot,
		 tier, confidence_score, model_used, reviewed_at)
		VALUES (?, ?, 'claude_code_enrichment', ?::jsonb, ?::jsonb, ?, ?, 'claude-sonnet-4-5-20250929', ?)
	`,
		uuid.Must(uuid.NewV7()),
		itemID,
		string(beforeJSON),
		string(afterJSON),
		enrichment.Tier,
		enrichment.Confidence,
		now,
	).Error

	if err != nil {
		log.Printf("⚠️  Warning: Failed to save audit trail: %v", err)
	} else {
		log.Println("✅ Audit trail saved")
	}

	fmt.Printf("\n✅ Enrichment applied successfully!\n")
	fmt.Printf("   Item ID: %s\n", enrichment.ScoreItemID)
	fmt.Printf("   Tier: %s\n", enrichment.Tier)
	fmt.Printf("   Confidence: %.2f\n", enrichment.Confidence)
	fmt.Printf("   Max Points: %d\n", enrichment.MaxPoints)
}

func validate(e *EnrichmentData) error {
	if e.ScoreItemID == "" {
		return fmt.Errorf("score_item_id is required")
	}
	if len(e.ClinicalRelevance) < 800 {
		return fmt.Errorf("clinical_relevance too short (min 800 chars)")
	}
	if len(e.PatientExplanation) < 400 {
		return fmt.Errorf("patient_explanation too short (min 400 chars)")
	}
	if len(e.Conduct) < 600 {
		return fmt.Errorf("conduct too short (min 600 chars)")
	}
	if e.MaxPoints < 0 || e.MaxPoints > 50 {
		return fmt.Errorf("max_points must be 0-50")
	}
	if e.Confidence < 0 || e.Confidence > 1 {
		return fmt.Errorf("confidence must be 0-1")
	}
	if e.Tier != "preserve" && e.Tier != "enrich" && e.Tier != "rewrite" {
		return fmt.Errorf("tier must be preserve/enrich/rewrite")
	}
	return nil
}
