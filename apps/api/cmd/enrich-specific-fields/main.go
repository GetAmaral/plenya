package main

// Enrich Specific Fields - Enriquece apenas patient_explanation e conduct via LLM
// Usa os prompts já preparados em score_item_enrichment_preparation
// NÃO toca em clinical_relevance
//
// Usage: go run cmd/enrich-specific-fields/main.go <uuid1> [uuid2 ...]

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const modelID = "claude-sonnet-4-6"
const apiURL = "https://api.anthropic.com/v1/messages"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run cmd/enrich-specific-fields/main.go <uuid1> [uuid2 ...]")
	}

	godotenv.Load()
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ ANTHROPIC_API_KEY não definida")
	}

	cfg, _ := config.Load()
	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("❌ DB: %v", err)
	}

	for _, arg := range os.Args[1:] {
		id, err := uuid.Parse(arg)
		if err != nil {
			log.Printf("❌ ID inválido: %s", arg)
			continue
		}
		if err := enrichItem(db, apiKey, id); err != nil {
			log.Printf("❌ Erro em %s: %v", id, err)
		}
	}
}

func enrichItem(db *gorm.DB, apiKey string, id uuid.UUID) error {
	// Carregar item
	var item models.ScoreItem
	if err := db.First(&item, id).Error; err != nil {
		return fmt.Errorf("item não encontrado: %w", err)
	}

	fmt.Printf("\n📄 %s\n", item.Name)
	fmt.Printf("   ID: %s\n\n", item.ID)

	// Carregar preparation
	var prep models.ScoreItemEnrichmentPreparation
	if err := db.First(&prep, "score_item_id = ?", id).Error; err != nil {
		return fmt.Errorf("preparation não encontrada (rode prepare-all primeiro): %w", err)
	}

	if prep.PromptPatientExplanation == nil || prep.PromptConduct == nil {
		return fmt.Errorf("prompts não gerados na preparation")
	}

	ctx := context.Background()

	// --- Patient Explanation ---
	fmt.Println("   🤖 Gerando patient_explanation...")
	pe, err := callClaude(ctx, apiKey, *prep.PromptPatientExplanation)
	if err != nil {
		return fmt.Errorf("patient_explanation: %w", err)
	}
	fmt.Printf("   ✅ patient_explanation: %d chars\n", len(pe))

	// --- Conduct ---
	fmt.Println("   🤖 Gerando conduct...")
	conduct, err := callClaude(ctx, apiKey, *prep.PromptConduct)
	if err != nil {
		return fmt.Errorf("conduct: %w", err)
	}
	fmt.Printf("   ✅ conduct: %d chars\n", len(conduct))

	// Salvar apenas PE e Conduct (NÃO toca clinical_relevance)
	if err := db.Model(&item).Updates(map[string]interface{}{
		"patient_explanation": pe,
		"conduct":             conduct,
	}).Error; err != nil {
		return fmt.Errorf("falha ao salvar: %w", err)
	}

	fmt.Println("   💾 Salvo no banco (clinical_relevance preservada)")
	return nil
}

func callClaude(ctx context.Context, apiKey, prompt string) (string, error) {
	body, _ := json.Marshal(map[string]interface{}{
		"model":      modelID,
		"max_tokens": 4096,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	})

	req, err := http.NewRequestWithContext(ctx, "POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API status %d: %s", resp.StatusCode, string(respBody))
	}

	var apiResp struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	}
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return "", fmt.Errorf("parse resposta: %w", err)
	}
	if len(apiResp.Content) == 0 {
		return "", fmt.Errorf("resposta vazia")
	}

	text := strings.TrimSpace(apiResp.Content[0].Text)
	return text, nil
}
