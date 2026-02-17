package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
	"gorm.io/gorm"
)

// ScoreItemEnrichmentService enriquece campos clínicos usando LLM baseado em artigos linkados
type ScoreItemEnrichmentService struct {
	db         *gorm.DB
	vectorRepo *repository.ArticleVectorRepository
	apiKey     string
	modelID    string // claude-sonnet-4-5-20250929 ou claude-haiku-4-5-20251001
}

// EnrichmentTier define estratégia de processamento
type EnrichmentTier string

const (
	TierPreserve  EnrichmentTier = "preserve"  // Não alterar (conteúdo excelente)
	TierEnrich    EnrichmentTier = "enrich"    // Melhorar incrementalmente
	TierRewrite   EnrichmentTier = "rewrite"   // Reescrever do zero
)

// EnrichmentResult resultado do enriquecimento LLM
type EnrichmentResult struct {
	ClinicalRelevance  string  `json:"clinical_relevance"`
	PatientExplanation string  `json:"patient_explanation"`
	Conduct            string  `json:"conduct"`
	MaxPoints          int     `json:"max_points"`
	Justification      string  `json:"justification"`
	Confidence         float64 `json:"confidence"`
}

// NewScoreItemEnrichmentService cria nova instância
func NewScoreItemEnrichmentService(db *gorm.DB) *ScoreItemEnrichmentService {
	return &ScoreItemEnrichmentService{
		db:         db,
		vectorRepo: repository.NewArticleVectorRepository(db),
		apiKey:     os.Getenv("ANTHROPIC_API_KEY"),
		modelID:    "claude-sonnet-4-5-20250929", // Default: Sonnet
	}
}

// SetModel altera o modelo (sonnet ou haiku para economia)
func (s *ScoreItemEnrichmentService) SetModel(modelID string) {
	s.modelID = modelID
}

// DetermineTier classifica ScoreItem em tier de processamento
func (s *ScoreItemEnrichmentService) DetermineTier(item *models.ScoreItem) EnrichmentTier {
	// Tier 1: PRESERVAR - conteúdo excelente e recente
	crLen := len(stringValue(item.ClinicalRelevance))
	peLen := len(stringValue(item.PatientExplanation))
	condLen := len(stringValue(item.Conduct))

	isRecent := item.LastReview != nil &&
		item.LastReview.After(time.Now().AddDate(0, -6, 0)) // 6 meses

	if crLen >= 1500 && peLen >= 600 && condLen >= 800 && isRecent {
		return TierPreserve
	}

	// Tier 3: REESCREVER - campos ausentes ou muito curtos
	if crLen < 500 {
		return TierRewrite
	}

	// Tier 2: ENRIQUECER - tudo no meio
	return TierEnrich
}

// GenerateEnrichment usa LLM para gerar/melhorar campos clínicos
func (s *ScoreItemEnrichmentService) GenerateEnrichment(
	ctx context.Context,
	item *models.ScoreItem,
	tier EnrichmentTier,
) (*EnrichmentResult, error) {
	if s.apiKey == "" {
		return nil, fmt.Errorf("ANTHROPIC_API_KEY environment variable is required")
	}

	// 1. Buscar artigos linkados (top 5 por similaridade)
	articleResults, err := s.vectorRepo.FindSimilarArticlesForScoreItem(
		item.ID,
		5,     // limit
		0.7,   // minSimilarity
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch linked articles: %w", err)
	}

	// 2. Construir contexto de artigos
	articlesContext := s.buildArticlesContext(articleResults)

	// 3. Construir prompt baseado no tier
	prompt := s.buildPrompt(item, tier, articlesContext)

	// 4. Chamar Claude API
	result, err := s.callClaudeAPI(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("Claude API call failed: %w", err)
	}

	return result, nil
}

// buildArticlesContext formata artigos para o prompt
func (s *ScoreItemEnrichmentService) buildArticlesContext(results []repository.ArticleSearchResult) string {
	if len(results) == 0 {
		return "Nenhum artigo disponível."
	}

	var builder strings.Builder
	for i, result := range results {
		builder.WriteString(fmt.Sprintf("%d. [%s] (%s, %d)\n",
			i+1,
			result.Article.Title,
			result.Article.Journal,
			result.Article.PublishDate.Year()))
		builder.WriteString(fmt.Sprintf("   Similaridade: %.3f\n", result.Similarity))

		if result.Article.Abstract != nil {
			abstract := *result.Article.Abstract
			if len(abstract) > 500 {
				abstract = abstract[:500] + "..."
			}
			builder.WriteString(fmt.Sprintf("   Resumo: %s\n", abstract))
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

// buildPrompt constrói prompt para Claude baseado no tier
func (s *ScoreItemEnrichmentService) buildPrompt(
	item *models.ScoreItem,
	tier EnrichmentTier,
	articlesContext string,
) string {
	var instruction string
	switch tier {
	case TierPreserve:
		return "" // Não deve chamar API para Tier Preserve
	case TierEnrich:
		instruction = `TAREFA: ENRIQUECER campos existentes, mantendo o conteúdo bom e expandindo onde necessário.
Preserve informações corretas. Adicione dados quantitativos, anos, e detalhes faltantes.`
	case TierRewrite:
		instruction = `TAREFA: GERAR campos clínicos completos do zero, baseado nos artigos fornecidos.`
	}

	unitStr := "N/A"
	if item.Unit != nil {
		unitStr = *item.Unit
	}

	genderStr := "not_applicable"
	if item.Gender != nil {
		genderStr = *item.Gender
	}

	existingCR := stringValue(item.ClinicalRelevance)
	existingPE := stringValue(item.PatientExplanation)
	existingCond := stringValue(item.Conduct)

	prompt := fmt.Sprintf(`Você é especialista clínico escrevendo conteúdo educacional para sistema EMR.

ScoreItem: %s
Unidade: %s
Filtros demográficos: gender=%s, age_range=%s
Pontuação atual: %.1f

Artigos linkados (ordenados por similaridade):
%s

%s

**Diretrizes de formato:**

1. clinical_relevance (1000-1500 caracteres):
   - Definição clínica e fisiopatologia
   - Dados epidemiológicos recentes (com anos específicos)
   - Significância clínica quantificada (risco relativo, odds ratio, etc)
   - Estratificação de risco por valores
   - Considerações populacionais (idade, gênero, comorbidades)

2. patient_explanation (500-800 caracteres):
   - Linguagem acessível (8º ano escolar, Flesch-Kincaid ≥60)
   - O QUE é este parâmetro
   - POR QUE importa para a saúde
   - O QUE significam valores anormais (evitar alarmismo)
   - Tom: empático e empoderador

3. conduct (800-1200 caracteres):
   - Estruturado com headers e bullets markdown
   - Thresholds numéricos específicos
   - Workup diagnóstico (exames complementares)
   - Protocolos de tratamento baseados em evidências
   - Critérios de referral/especialista
   - Periodicidade de monitoramento/follow-up

4. max_points (0-50):
   - Baseado em importância clínica e impacto prognóstico
   - Justificativa breve

**Conteúdo existente** (usar como base se Tier=enrich):
clinical_relevance: %s
patient_explanation: %s
conduct: %s

Retorne APENAS JSON válido (sem markdown, sem explicações):
{
  "clinical_relevance": "...",
  "patient_explanation": "...",
  "conduct": "...",
  "max_points": 30,
  "justification": "...",
  "confidence": 0.92
}`,
		item.Name,
		unitStr,
		genderStr,
		formatAgeRange(item.AgeRangeMin, item.AgeRangeMax),
		getPoints(item.Points),
		articlesContext,
		instruction,
		existingCR,
		existingPE,
		existingCond,
	)

	return prompt
}

// callClaudeAPI chama Anthropic API (simplificado - em produção usar SDK oficial)
func (s *ScoreItemEnrichmentService) callClaudeAPI(ctx context.Context, prompt string) (*EnrichmentResult, error) {
	// TODO: Integrar com Anthropic SDK oficial (github.com/anthropics/anthropic-sdk-go)
	// Por ora, retornar mock para demonstração

	// Em produção:
	// client := anthropic.NewClient(s.apiKey)
	// response, err := client.Messages.Create(ctx, &anthropic.MessageCreateParams{
	// 	Model: s.modelID,
	// 	Messages: []anthropic.Message{{
	// 		Role: "user",
	// 		Content: prompt,
	// 	}},
	// 	MaxTokens: 2000,
	// })

	// Preparar request body
	requestBody := map[string]interface{}{
		"model":      s.modelID,
		"max_tokens": 8000,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Criar HTTP request
	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", s.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	// Executar request
	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	// Ler resposta
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(respBody))
	}

	// Parse resposta da API
	var apiResponse struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	}

	if err := json.Unmarshal(respBody, &apiResponse); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %w", err)
	}

	if len(apiResponse.Content) == 0 {
		return nil, fmt.Errorf("empty response from API")
	}

	// Extrair texto
	text := apiResponse.Content[0].Text

	// Parse JSON do texto (Claude retorna JSON puro)
	// Limpar possível markdown
	text = strings.TrimPrefix(text, "```json")
	text = strings.TrimPrefix(text, "```")
	text = strings.TrimSuffix(text, "```")
	text = strings.TrimSpace(text)

	var result EnrichmentResult
	if err := json.Unmarshal([]byte(text), &result); err != nil {
		return nil, fmt.Errorf("failed to parse LLM JSON response: %w (text: %s)", err, text[:minInt(200, len(text))])
	}

	return &result, nil
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ValidateResult valida qualidade do resultado LLM
func (s *ScoreItemEnrichmentService) ValidateResult(result *EnrichmentResult) []string {
	var errors []string

	// Validar length requirements (limites mais flexíveis para aceitar conteúdo de qualidade)
	if len(result.ClinicalRelevance) < 800 {
		errors = append(errors, "clinical_relevance too short (<800 chars)")
	}
	if len(result.ClinicalRelevance) > 3000 {
		errors = append(errors, "clinical_relevance too long (>3000 chars)")
	}

	if len(result.PatientExplanation) < 400 {
		errors = append(errors, "patient_explanation too short (<400 chars)")
	}
	if len(result.PatientExplanation) > 2000 {
		errors = append(errors, "patient_explanation too long (>2000 chars)")
	}

	if len(result.Conduct) < 600 {
		errors = append(errors, "conduct too short (<600 chars)")
	}
	if len(result.Conduct) > 6000 {
		errors = append(errors, "conduct too long (>6000 chars)")
	}

	// Validar max_points range
	if result.MaxPoints < 0 || result.MaxPoints > 50 {
		errors = append(errors, "max_points out of range (must be 0-50)")
	}

	// Validar confidence
	if result.Confidence < 0 || result.Confidence > 1 {
		errors = append(errors, "confidence out of range (must be 0-1)")
	}

	// Detectar placeholders
	placeholders := []string{"[", "TODO", "TBD", "..."}
	for _, ph := range placeholders {
		if strings.Contains(result.ClinicalRelevance, ph) ||
			strings.Contains(result.PatientExplanation, ph) ||
			strings.Contains(result.Conduct, ph) {
			errors = append(errors, fmt.Sprintf("placeholder detected: %s", ph))
			break
		}
	}

	return errors
}

// SaveReviewHistory cria snapshot antes de atualizar (auditoria)
func (s *ScoreItemEnrichmentService) SaveReviewHistory(
	item *models.ScoreItem,
	newResult *EnrichmentResult,
	tier EnrichmentTier,
) error {
	// Serializar estado antes
	beforeJSON, err := json.Marshal(map[string]interface{}{
		"clinical_relevance":  item.ClinicalRelevance,
		"patient_explanation": item.PatientExplanation,
		"conduct":             item.Conduct,
		"points":              item.Points,
	})
	if err != nil {
		return fmt.Errorf("failed to serialize before state: %w", err)
	}

	// Serializar estado depois
	afterJSON, err := json.Marshal(map[string]interface{}{
		"clinical_relevance":  newResult.ClinicalRelevance,
		"patient_explanation": newResult.PatientExplanation,
		"conduct":             newResult.Conduct,
		"max_points":          newResult.MaxPoints,
	})
	if err != nil {
		return fmt.Errorf("failed to serialize after state: %w", err)
	}

	// Inserir em score_item_review_history
	err = s.db.Exec(`
		INSERT INTO score_item_review_history
		(id, score_item_id, review_type, before_snapshot, after_snapshot,
		 tier, confidence_score, model_used, reviewed_at)
		VALUES (?, ?, 'llm_enrichment', ?::jsonb, ?::jsonb, ?, ?, ?, NOW())
	`,
		uuid.Must(uuid.NewV7()),
		item.ID,
		string(beforeJSON),
		string(afterJSON),
		string(tier),
		newResult.Confidence,
		s.modelID,
	).Error

	return err
}

// Helper functions
func stringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func formatAgeRange(min, max *int) string {
	if min == nil && max == nil {
		return "all ages"
	}
	if min != nil && max == nil {
		return fmt.Sprintf("≥%d years", *min)
	}
	if min == nil && max != nil {
		return fmt.Sprintf("≤%d years", *max)
	}
	return fmt.Sprintf("%d-%d years", *min, *max)
}

func getPoints(p *float64) float64 {
	if p == nil {
		return 0
	}
	return *p
}
