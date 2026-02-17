package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/liushuangls/go-anthropic/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Enrichment result structure
type EnrichmentResult struct {
	ClinicalRelevance  string `json:"clinical_relevance"`
	PatientExplanation string `json:"patient_explanation"`
	Conduct            string `json:"conduct"`
}

// ScoreItemEnrich represents the minimal data needed for enrichment
type ScoreItemEnrich struct {
	ID                 uuid.UUID
	Name               string
	Unit               *string
	Gender             *string
	AgeRangeMin        *int
	AgeRangeMax        *int
	PostMenopause      *bool
	ClinicalRelevance  *string
	PatientExplanation *string
	Conduct            *string
	SubgroupName       string
	GroupName          string
}

// Progress report structure
type ProgressReport struct {
	Index       int
	Total       int
	ItemID      string
	ItemName    string
	Tier        string
	Status      string
	Error       string
	ProcessedAt time.Time
}

func main() {
	// Parse command line flags
	startOffset := flag.Int("offset", 0, "Starting offset (default: 0)")
	batchSize := flag.Int("batch", 10, "Number of items to process (default: 10)")
	dryRun := flag.Bool("dry-run", false, "Dry run mode - no database updates")
	flag.Parse()

	// Load environment variables
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		log.Fatal("ANTHROPIC_API_KEY environment variable not set")
	}

	// Database connection
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "plenya_user"),
		getEnv("DB_PASSWORD", "plenya_password"),
		getEnv("DB_NAME", "plenya_db"),
		getEnv("DB_PORT", "5432"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Get total count
	var totalCount int64
	if err := db.Table("score_items").
		Where("last_review IS NULL AND deleted_at IS NULL").
		Count(&totalCount).Error; err != nil {
		log.Fatalf("Failed to count items: %v", err)
	}

	log.Printf("=== Score Item Enrichment Tool ===")
	log.Printf("Total items with NULL last_review: %d", totalCount)
	log.Printf("Starting offset: %d", *startOffset)
	log.Printf("Batch size: %d", *batchSize)
	log.Printf("Dry run mode: %v\n", *dryRun)

	// Create Anthropic client
	client := anthropic.NewClient(apiKey)

	// Process items
	successCount := 0
	errorCount := 0
	skippedCount := 0

	for i := 0; i < *batchSize; i++ {
		offset := *startOffset + i

		log.Printf("\n[%d/%d] Processing item at offset %d...", i+1, *batchSize, offset)

		// Fetch single item with hierarchy
		var item ScoreItemEnrich
		err := db.Table("score_items si").
			Select(`
				si.id,
				si.name,
				si.unit,
				si.gender,
				si.age_range_min,
				si.age_range_max,
				si.post_menopause,
				si.clinical_relevance,
				si.patient_explanation,
				si.conduct,
				ss.name as subgroup_name,
				sg.name as group_name
			`).
			Joins("JOIN score_subgroups ss ON si.subgroup_id = ss.id").
			Joins("JOIN score_groups sg ON ss.group_id = sg.id").
			Where("si.last_review IS NULL AND si.deleted_at IS NULL").
			Order("si.created_at").
			Limit(1).
			Offset(offset).
			Scan(&item).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				log.Printf("No more items found at offset %d", offset)
				break
			}
			log.Printf("ERROR: Failed to fetch item: %v", err)
			errorCount++
			continue
		}

		// Build tier string
		tier := fmt.Sprintf("%s > %s", item.GroupName, item.SubgroupName)

		// Check if already enriched
		if item.ClinicalRelevance != nil && item.PatientExplanation != nil && item.Conduct != nil {
			log.Printf("SKIPPED: %s [%s] - Already has enrichment", item.Name, tier)
			skippedCount++
			continue
		}

		// Enrich with LLM
		log.Printf("ENRICHING: %s [%s]", item.Name, tier)
		enrichment, err := enrichWithLLM(client, &item)
		if err != nil {
			log.Printf("ERROR: LLM enrichment failed: %v", err)
			errorCount++
			continue
		}

		// Update database
		if !*dryRun {
			now := time.Now()
			updateData := map[string]interface{}{
				"clinical_relevance":  enrichment.ClinicalRelevance,
				"patient_explanation": enrichment.PatientExplanation,
				"conduct":             enrichment.Conduct,
				"last_review":         now,
				"updated_at":          now,
			}

			if err := db.Table("score_items").
				Where("id = ?", item.ID).
				Updates(updateData).Error; err != nil {
				log.Printf("ERROR: Database update failed: %v", err)
				errorCount++
				continue
			}
		}

		log.Printf("SUCCESS: %s [%s]", item.Name, tier)
		successCount++

		// Rate limiting - respect Anthropic's rate limits
		if i < *batchSize-1 {
			time.Sleep(1 * time.Second)
		}
	}

	// Final report
	log.Printf("\n=== Enrichment Summary ===")
	log.Printf("Processed: %d items", successCount+errorCount+skippedCount)
	log.Printf("Success: %d", successCount)
	log.Printf("Skipped: %d", skippedCount)
	log.Printf("Errors: %d", errorCount)
	log.Printf("Remaining: %d", totalCount-int64(*startOffset+*batchSize))
}

func enrichWithLLM(client *anthropic.Client, item *ScoreItemEnrich) (*EnrichmentResult, error) {
	// Build context prompt
	prompt := buildEnrichmentPrompt(item)

	// Call Claude API
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	resp, err := client.CreateMessages(ctx, anthropic.MessagesRequest{
		Model: anthropic.ModelClaude3_5Sonnet20241022,
		Messages: []anthropic.Message{
			{
				Role: anthropic.RoleUser,
				Content: []anthropic.MessageContent{
					anthropic.NewTextMessageContent(prompt),
				},
			},
		},
		MaxTokens: 2000,
		Temperature: float32Ptr(0.3), // Lower temperature for more consistent medical information
	})

	if err != nil {
		return nil, fmt.Errorf("API call failed: %w", err)
	}

	// Extract text from response
	if len(resp.Content) == 0 {
		return nil, fmt.Errorf("empty response from API")
	}

	textContent, ok := resp.Content[0].(*anthropic.MessageContentText)
	if !ok {
		return nil, fmt.Errorf("unexpected content type")
	}

	// Parse JSON response
	var result EnrichmentResult
	if err := json.Unmarshal([]byte(textContent.Text), &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w\nResponse: %s", err, textContent.Text)
	}

	return &result, nil
}

func buildEnrichmentPrompt(item *ScoreItemEnrich) string {
	// Build demographic filters text
	demographics := ""
	if item.Gender != nil && *item.Gender != "not_applicable" {
		demographics += fmt.Sprintf("Gênero: %s\n", *item.Gender)
	}
	if item.AgeRangeMin != nil && item.AgeRangeMax != nil {
		demographics += fmt.Sprintf("Faixa etária: %d-%d anos\n", *item.AgeRangeMin, *item.AgeRangeMax)
	} else if item.AgeRangeMin != nil {
		demographics += fmt.Sprintf("Idade mínima: %d anos\n", *item.AgeRangeMin)
	} else if item.AgeRangeMax != nil {
		demographics += fmt.Sprintf("Idade máxima: %d anos\n", *item.AgeRangeMax)
	}
	if item.PostMenopause != nil && *item.PostMenopause {
		demographics += "Aplicável apenas para mulheres pós-menopausa\n"
	}

	unit := ""
	if item.Unit != nil {
		unit = fmt.Sprintf("Unidade: %s\n", *item.Unit)
	}

	prompt := fmt.Sprintf(`Você é um médico especialista criando conteúdo para um sistema de prontuário eletrônico (EMR).

CONTEXTO:
Grupo: %s
Subgrupo: %s
Parâmetro: %s
%s%s

Gere 3 campos de enriquecimento clínico:

1. **clinical_relevance**: Explicação técnica para profissionais de saúde (2-3 frases). Foco em fisiopatologia, associações clínicas, impacto na saúde.

2. **patient_explanation**: Explicação simples para o paciente (1-2 frases). Linguagem acessível, evite jargões médicos.

3. **conduct**: Recomendações de conduta clínica (2-4 frases). Quando investigar, quando encaminhar, quando tratar, pontos de corte importantes.

IMPORTANTE:
- Se for um parâmetro laboratorial, inclua valores de referência e interpretação clínica
- Se for anamnese/histórico, foque em relevância diagnóstica e fatores de risco
- Se for exame de imagem, mencione achados importantes e quando solicitar
- Considere os filtros demográficos na interpretação
- Seja conciso e objetivo
- Use linguagem em português brasileiro

Responda APENAS com um JSON válido (sem markdown, sem explicações):

{
  "clinical_relevance": "...",
  "patient_explanation": "...",
  "conduct": "..."
}`,
		item.GroupName,
		item.SubgroupName,
		item.Name,
		unit,
		demographics,
	)

	return prompt
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func float32Ptr(f float32) *float32 {
	return &f
}
