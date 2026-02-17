package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/models"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

type EnrichmentResponse struct {
	ClinicalRelevance  string `json:"clinical_relevance"`
	PatientExplanation string `json:"patient_explanation"`
	Conduct            string `json:"conduct"`
}

func main() {
	// Flags
	tier := flag.String("tier", "letter", "Tier to enrich: letter or pillar")
	offset := flag.Int("offset", 0, "Offset for pagination (0-based)")
	dryRun := flag.Bool("dry-run", false, "Preview without updating database")
	flag.Parse()

	// Validate tier
	if *tier != "letter" && *tier != "pillar" {
		log.Fatal("Error: tier must be 'letter' or 'pillar'")
	}

	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Check OpenAI API key
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("Error: OPENAI_API_KEY environment variable not set")
	}

	client := openai.NewClient(apiKey)

	// Process based on tier
	if *tier == "letter" {
		processLetter(database.DB, client, *offset, *dryRun)
	} else {
		processPillar(database.DB, client, *offset, *dryRun)
	}
}

func processLetter(db *gorm.DB, client *openai.Client, offset int, dryRun bool) {
	var letter models.MethodLetter

	// Query for unenriched letter
	result := db.Where("last_review IS NULL").
		Order("\"order\" ASC").
		Limit(1).
		Offset(offset).
		Preload("Method").
		First(&letter)

	if result.Error != nil {
		log.Fatalf("Error querying letter: %v", result.Error)
	}

	// Report current item
	fmt.Printf("\n=== ENRICHING METHOD LETTER ===\n")
	fmt.Printf("Name: %s\n", letter.Name)
	fmt.Printf("Code: %s\n", letter.Code)
	fmt.Printf("Method: %s\n", letter.Method.ShortName)
	fmt.Printf("Tier: Letter\n")
	fmt.Printf("Status: Unenriched (last_review = NULL)\n")
	fmt.Printf("Offset: %d\n\n", offset)

	// Generate enrichment using Claude
	enrichment := enrichLetterWithClaude(client, letter)

	if dryRun {
		fmt.Println("=== DRY RUN - Preview Only ===")
		fmt.Printf("\nClinical Relevance:\n%s\n", enrichment.ClinicalRelevance)
		fmt.Printf("\nPatient Explanation:\n%s\n", enrichment.PatientExplanation)
		fmt.Printf("\nConduct:\n%s\n", enrichment.Conduct)
		fmt.Println("\nNo database changes made (dry-run mode)")
		return
	}

	// Update database
	letter.ClinicalRelevance = &enrichment.ClinicalRelevance
	letter.PatientExplanation = &enrichment.PatientExplanation
	letter.Conduct = &enrichment.Conduct

	if err := db.Save(&letter).Error; err != nil {
		log.Fatalf("Error updating letter: %v", err)
	}

	fmt.Println("=== ENRICHMENT COMPLETE ===")
	fmt.Printf("Status: Enriched (last_review auto-updated via hook)\n")
	fmt.Printf("Letter ID: %s\n", letter.ID)
}

func processPillar(db *gorm.DB, client *openai.Client, offset int, dryRun bool) {
	var pillar models.MethodPillar

	// Query for unenriched pillar
	result := db.Where("last_review IS NULL").
		Order("\"order\" ASC").
		Limit(1).
		Offset(offset).
		Preload("Letter").
		Preload("Letter.Method").
		First(&pillar)

	if result.Error != nil {
		log.Fatalf("Error querying pillar: %v", result.Error)
	}

	// Report current item
	fmt.Printf("\n=== ENRICHING METHOD PILLAR ===\n")
	fmt.Printf("Name: %s\n", pillar.Name)
	fmt.Printf("Letter: %s (%s)\n", pillar.Letter.Name, pillar.Letter.Code)
	fmt.Printf("Method: %s\n", pillar.Letter.Method.ShortName)
	fmt.Printf("Tier: Pillar\n")
	fmt.Printf("Status: Unenriched (last_review = NULL)\n")
	fmt.Printf("Offset: %d\n\n", offset)

	// Generate enrichment using Claude
	enrichment := enrichPillarWithClaude(client, pillar)

	if dryRun {
		fmt.Println("=== DRY RUN - Preview Only ===")
		fmt.Printf("\nClinical Relevance:\n%s\n", enrichment.ClinicalRelevance)
		fmt.Printf("\nPatient Explanation:\n%s\n", enrichment.PatientExplanation)
		fmt.Printf("\nConduct:\n%s\n", enrichment.Conduct)
		fmt.Println("\nNo database changes made (dry-run mode)")
		return
	}

	// Update database
	pillar.ClinicalRelevance = &enrichment.ClinicalRelevance
	pillar.PatientExplanation = &enrichment.PatientExplanation
	pillar.Conduct = &enrichment.Conduct

	if err := db.Save(&pillar).Error; err != nil {
		log.Fatalf("Error updating pillar: %v", err)
	}

	fmt.Println("=== ENRICHMENT COMPLETE ===")
	fmt.Printf("Status: Enriched (last_review auto-updated via hook)\n")
	fmt.Printf("Pillar ID: %s\n", pillar.ID)
}

func enrichLetterWithClaude(client *openai.Client, letter models.MethodLetter) EnrichmentResponse {
	prompt := fmt.Sprintf(`You are a medical expert specializing in integrative health protocols.

Context:
- Method: %s
- Letter: %s (%s)
- Description: %s

Task: Generate clinical enrichment content for this methodology letter. Return ONLY a valid JSON object with exactly three fields:

{
  "clinical_relevance": "2-3 sentences explaining why this letter/phase is clinically important for patient outcomes",
  "patient_explanation": "2-3 sentences explaining this letter/phase in simple, patient-friendly language (avoid jargon)",
  "conduct": "2-3 bullet points with specific clinical recommendations for practitioners"
}

Guidelines:
- Be concise and evidence-based
- Use Brazilian Portuguese medical terminology
- Focus on practical, actionable information
- Conduct should be specific clinical actions (e.g., "• Solicitar hemograma completo e perfil lipídico")
- Do not include markdown formatting or code blocks
- Return ONLY the JSON object, nothing else`,
		letter.Method.ShortName,
		letter.Name,
		letter.Code,
		getStringOrDefault(letter.Description, "N/A"),
	)

	ctx := context.Background()
	response, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONObject,
		},
		MaxTokens: 2000,
	})

	if err != nil {
		log.Fatalf("Error calling OpenAI API: %v", err)
	}

	// Extract text from response
	text := response.Choices[0].Message.Content
	var enrichment EnrichmentResponse
	if err := json.Unmarshal([]byte(text), &enrichment); err != nil {
		log.Fatalf("Error parsing OpenAI response as JSON: %v\nResponse: %s", err, text)
	}

	if enrichment.ClinicalRelevance == "" {
		log.Fatal("Error: OpenAI response missing clinical_relevance")
	}

	return enrichment
}

func enrichPillarWithClaude(client *openai.Client, pillar models.MethodPillar) EnrichmentResponse {
	prompt := fmt.Sprintf(`You are a medical expert specializing in integrative health protocols.

Context:
- Method: %s
- Letter: %s (%s)
- Pillar: %s
- Description: %s

Task: Generate clinical enrichment content for this methodology pillar. Return ONLY a valid JSON object with exactly three fields:

{
  "clinical_relevance": "2-3 sentences explaining why this pillar is clinically important within this letter/phase",
  "patient_explanation": "2-3 sentences explaining this pillar in simple, patient-friendly language (avoid jargon)",
  "conduct": "2-3 bullet points with specific clinical recommendations for practitioners"
}

Guidelines:
- Be concise and evidence-based
- Use Brazilian Portuguese medical terminology
- Focus on practical, actionable information
- Conduct should be specific clinical actions (e.g., "• Solicitar hemograma completo e perfil lipídico")
- Do not include markdown formatting or code blocks
- Return ONLY the JSON object, nothing else`,
		pillar.Letter.Method.ShortName,
		pillar.Letter.Name,
		pillar.Letter.Code,
		pillar.Name,
		getStringOrDefault(pillar.Description, "N/A"),
	)

	ctx := context.Background()
	response, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONObject,
		},
		MaxTokens: 2000,
	})

	if err != nil {
		log.Fatalf("Error calling OpenAI API: %v", err)
	}

	// Extract text from response
	text := response.Choices[0].Message.Content
	var enrichment EnrichmentResponse
	if err := json.Unmarshal([]byte(text), &enrichment); err != nil {
		log.Fatalf("Error parsing OpenAI response as JSON: %v\nResponse: %s", err, text)
	}

	if enrichment.ClinicalRelevance == "" {
		log.Fatal("Error: OpenAI response missing clinical_relevance")
	}

	return enrichment
}

func getStringOrDefault(s *string, defaultVal string) string {
	if s == nil {
		return defaultVal
	}
	return *s
}
