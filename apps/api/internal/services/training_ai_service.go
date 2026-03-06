package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/plenya/api/internal/config"
	"gorm.io/gorm"
)

type TrainingAIService struct {
	db              *gorm.DB
	apiKey          string
	model           string
	httpClient      *http.Client
	semanticService *ArticleSemanticService
}

func NewTrainingAIService(
	db *gorm.DB,
	cfg *config.Config,
	semanticService *ArticleSemanticService,
) *TrainingAIService {
	return &TrainingAIService{
		db:              db,
		apiKey:          cfg.Claude.APIKey,
		model:           cfg.Claude.Model,
		httpClient:      &http.Client{Timeout: 120 * time.Second},
		semanticService: semanticService,
	}
}

type ChatMessage struct {
	Role    string `json:"role"` // "user" or "assistant"
	Content string `json:"content"`
}

type PatientContext struct {
	Name        string   `json:"name"`
	Age         int      `json:"age"`
	Gender      string   `json:"gender"`
	Weight      float64  `json:"weight,omitempty"`
	Height      float64  `json:"height,omitempty"`
	BMI         float64  `json:"bmi,omitempty"`
	Objective   string   `json:"objective,omitempty"`
	RiskLevel   string   `json:"riskLevel,omitempty"`
	RiskFactors []string `json:"riskFactors,omitempty"`
}

type ChatRequest struct {
	Messages       []ChatMessage  `json:"messages" validate:"required,min=1"`
	PatientContext *PatientContext `json:"patientContext,omitempty"`
}

type ChatResponse struct {
	Message string   `json:"message"`
	Sources []string `json:"sources,omitempty"`
}

// Chat processes a training-related chat with RAG context
func (s *TrainingAIService) Chat(ctx context.Context, req *ChatRequest) (*ChatResponse, error) {
	if s.apiKey == "" {
		return nil, fmt.Errorf("Claude API key not configured")
	}

	// Get the last user message for RAG search
	lastMessage := ""
	for i := len(req.Messages) - 1; i >= 0; i-- {
		if req.Messages[i].Role == "user" {
			lastMessage = req.Messages[i].Content
			break
		}
	}

	// RAG: search for relevant scientific chunks
	var ragContext string
	var sources []string
	if s.semanticService != nil && lastMessage != "" {
		results, err := s.semanticService.SemanticSearch(ctx, SemanticSearchDTO{
			Query:         lastMessage,
			Limit:         5,
			MinSimilarity: 0.4,
		})
		if err == nil && len(results) > 0 {
			ragContext = "\n\n## Referências Científicas Relevantes\n"
			for i, r := range results {
				ragContext += fmt.Sprintf("\n### Fonte %d: %s\n%s\n", i+1, r.Article.Title, r.ChunkText)
				sources = append(sources, r.Article.Title)
			}
		}
	}

	// Build system prompt
	systemPrompt := `Você é um assistente especializado em fisiologia do exercício, prescrição de treinamento e avaliação física.

Suas áreas de expertise incluem:
- Estratificação de risco ACSM (American College of Sports Medicine)
- Prescrição de exercícios baseada em evidências
- Periodização de treinamento (Bompa, Linear, Ondulatória, Bloco)
- Composição corporal e antropometria
- Bioenergética e fisiologia do exercício
- Biomecânica do movimento humano

Regras:
1. Sempre baseie suas respostas em evidências científicas
2. Cite as fontes quando disponíveis no contexto
3. Use terminologia técnica adequada mas explique quando necessário
4. Considere o contexto do paciente quando disponível
5. Responda em português brasileiro
6. Seja conciso e prático`

	if req.PatientContext != nil {
		pc := req.PatientContext
		systemPrompt += fmt.Sprintf("\n\n## Contexto do Paciente\n- Nome: %s\n- Idade: %d anos\n- Sexo: %s",
			pc.Name, pc.Age, pc.Gender)
		if pc.Weight > 0 {
			systemPrompt += fmt.Sprintf("\n- Peso: %.1f kg", pc.Weight)
		}
		if pc.Height > 0 {
			systemPrompt += fmt.Sprintf("\n- Altura: %.0f cm", pc.Height)
		}
		if pc.BMI > 0 {
			systemPrompt += fmt.Sprintf("\n- IMC: %.1f", pc.BMI)
		}
		if pc.Objective != "" {
			systemPrompt += fmt.Sprintf("\n- Objetivo: %s", pc.Objective)
		}
		if pc.RiskLevel != "" {
			systemPrompt += fmt.Sprintf("\n- Nível de Risco ACSM: %s", pc.RiskLevel)
		}
		if len(pc.RiskFactors) > 0 {
			systemPrompt += fmt.Sprintf("\n- Fatores de Risco: %v", pc.RiskFactors)
		}
	}

	if ragContext != "" {
		systemPrompt += ragContext
	}

	// Build messages for Claude API
	claudeMessages := make([]map[string]string, len(req.Messages))
	for i, m := range req.Messages {
		claudeMessages[i] = map[string]string{
			"role":    m.Role,
			"content": m.Content,
		}
	}

	payload := map[string]interface{}{
		"model":       s.model,
		"max_tokens":  4096,
		"temperature": 0.7,
		"system":      systemPrompt,
		"messages":    claudeMessages,
	}

	resp, err := s.callClaude(payload)
	if err != nil {
		return nil, err
	}

	return &ChatResponse{
		Message: resp,
		Sources: sources,
	}, nil
}

// GenerateMesocycles generates periodization mesocycles using RAG + Claude
func (s *TrainingAIService) GenerateMesocycles(
	ctx context.Context,
	patientContext *PatientContext,
	framework string,
	totalWeeks int,
	objective string,
) ([]MesocycleGenerated, string, error) {
	if s.apiKey == "" {
		return nil, "", fmt.Errorf("Claude API key not configured")
	}

	// RAG: search for periodization literature
	ragQuery := fmt.Sprintf("periodization %s training %s mesocycle phases volume intensity", framework, objective)
	var ragContext string
	if s.semanticService != nil {
		results, err := s.semanticService.SemanticSearch(ctx, SemanticSearchDTO{
			Query:         ragQuery,
			Limit:         8,
			MinSimilarity: 0.35,
		})
		if err == nil && len(results) > 0 {
			ragContext = "\n\n## Literatura Científica de Referência\n"
			for i, r := range results {
				ragContext += fmt.Sprintf("\n### %s\n%s\n", r.Article.Title, r.ChunkText)
				if i >= 7 {
					break
				}
			}
		}
	}

	patientInfo := ""
	if patientContext != nil {
		patientInfo = fmt.Sprintf("Paciente: %s, %d anos, %s", patientContext.Name, patientContext.Age, patientContext.Gender)
		if patientContext.Weight > 0 {
			patientInfo += fmt.Sprintf(", %.0fkg", patientContext.Weight)
		}
		if patientContext.RiskLevel != "" {
			patientInfo += fmt.Sprintf(", Risco ACSM: %s", patientContext.RiskLevel)
		}
	}

	prompt := fmt.Sprintf(`# Gerar Periodização de Treinamento

## Parâmetros
- Framework: %s
- Duração total: %d semanas
- Objetivo: %s
- %s
%s

## Instrução
Gere os mesociclos para esta periodização. Para cada mesociclo, defina:
1. phase: fase do mesociclo (accumulation, transformation, realization, hypertrophy, strength, endurance, power)
2. durationWeeks: duração em semanas
3. volumePercent: volume relativo (0-100)
4. intensityPercent: intensidade relativa (0-100)
5. physiologicalFocus: foco fisiológico principal

A soma das durações deve ser igual a %d semanas.

Também forneça uma justificativa científica para a distribuição dos mesociclos.

Responda em JSON com o formato:
{
  "mesocycles": [
    {"phase": "...", "durationWeeks": N, "volumePercent": N, "intensityPercent": N, "physiologicalFocus": "..."},
    ...
  ],
  "scientificJustification": "texto explicando a razão científica..."
}

IMPORTANTE: Retorne APENAS o JSON, sem explicações adicionais.`,
		framework, totalWeeks, objective, patientInfo, ragContext, totalWeeks)

	payload := map[string]interface{}{
		"model":       s.model,
		"max_tokens":  4096,
		"temperature": 0.5,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}

	resp, err := s.callClaude(payload)
	if err != nil {
		return nil, "", fmt.Errorf("Claude API error: %w", err)
	}

	// Parse JSON response
	var result struct {
		Mesocycles              []MesocycleGenerated `json:"mesocycles"`
		ScientificJustification string               `json:"scientificJustification"`
	}

	// Find JSON in response (may have markdown code blocks)
	jsonStr := resp
	if start := findJSONStart(jsonStr); start >= 0 {
		if end := findJSONEnd(jsonStr, start); end > start {
			jsonStr = jsonStr[start : end+1]
		}
	}

	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, "", fmt.Errorf("failed to parse AI response: %w (response: %s)", err, resp[:min(500, len(resp))])
	}

	return result.Mesocycles, result.ScientificJustification, nil
}

type MesocycleGenerated struct {
	Phase              string `json:"phase"`
	DurationWeeks      int    `json:"durationWeeks"`
	VolumePercent      int    `json:"volumePercent"`
	IntensityPercent   int    `json:"intensityPercent"`
	PhysiologicalFocus string `json:"physiologicalFocus"`
}

// GenerateWorkoutRecommendation generates AI-powered workout recommendations
func (s *TrainingAIService) GenerateWorkoutRecommendation(
	ctx context.Context,
	patientContext *PatientContext,
	objective string,
) (string, error) {
	if s.apiKey == "" {
		return "", fmt.Errorf("Claude API key not configured")
	}

	ragQuery := fmt.Sprintf("resistance training %s exercise prescription NSCA guidelines", objective)
	var ragContext string
	if s.semanticService != nil {
		results, err := s.semanticService.SemanticSearch(ctx, SemanticSearchDTO{
			Query:         ragQuery,
			Limit:         5,
			MinSimilarity: 0.35,
		})
		if err == nil && len(results) > 0 {
			ragContext = "\n\n## Referências:\n"
			for i, r := range results {
				ragContext += fmt.Sprintf("%d. %s: %s\n", i+1, r.Article.Title, r.ChunkText[:min(300, len(r.ChunkText))])
			}
		}
	}

	patientInfo := "Não disponível"
	if patientContext != nil {
		patientInfo = fmt.Sprintf("%s, %d anos, %s", patientContext.Name, patientContext.Age, patientContext.Gender)
		if patientContext.RiskLevel != "" {
			patientInfo += fmt.Sprintf(", Risco: %s", patientContext.RiskLevel)
		}
	}

	prompt := fmt.Sprintf(`Gere uma recomendação de treino baseada em evidências.

Paciente: %s
Objetivo: %s
%s

Forneça:
1. Recomendação de volume (séries/repetições) baseada nas diretrizes NSCA/ACSM
2. Intensidade recomendada
3. Frequência semanal ideal
4. Exercícios prioritários por grupo muscular
5. Progressão sugerida

Responda em português, de forma concisa e prática.`, patientInfo, objective, ragContext)

	payload := map[string]interface{}{
		"model":       s.model,
		"max_tokens":  2048,
		"temperature": 0.6,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}

	return s.callClaude(payload)
}

// callClaude makes a raw HTTP call to Claude API and returns the text response
func (s *TrainingAIService) callClaude(payload map[string]interface{}) (string, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", s.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to call Claude API: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Claude API error %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var apiResp struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
		Usage struct {
			InputTokens  int `json:"input_tokens"`
			OutputTokens int `json:"output_tokens"`
		} `json:"usage"`
	}

	if err := json.Unmarshal(bodyBytes, &apiResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	for _, content := range apiResp.Content {
		if content.Type == "text" {
			return content.Text, nil
		}
	}

	return "", fmt.Errorf("no text content in Claude response")
}

// findJSONStart finds the start of a JSON object in text
func findJSONStart(s string) int {
	for i, c := range s {
		if c == '{' {
			return i
		}
	}
	return -1
}

// findJSONEnd finds the matching closing brace
func findJSONEnd(s string, start int) int {
	depth := 0
	for i := start; i < len(s); i++ {
		switch s[i] {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return i
			}
		}
	}
	return -1
}
