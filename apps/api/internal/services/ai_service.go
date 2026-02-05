package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/dto"
)

// Helper min function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// AIService - serviço de integração com Claude API (Anthropic)
type AIService struct {
	apiKey     string
	model      string
	httpClient *http.Client
}

// NewAIService cria uma nova instância do serviço de IA
func NewAIService(cfg *config.Config) *AIService {
	return &AIService{
		apiKey:     cfg.Claude.APIKey,
		model:      cfg.Claude.Model,
		httpClient: &http.Client{Timeout: 180 * time.Second}, // 3 minutos para processar laudos grandes
	}
}

// InterpretLabResult - interpreta laudo médico via Claude API com structured output
func (s *AIService) InterpretLabResult(
	ocrText string,
	availableTests []dto.LabTestSummary,
) (*dto.AILabResultExtractionResponse, error) {
	prompt := s.buildPrompt(ocrText, availableTests)

	payload := map[string]interface{}{
		"model":       s.model,
		"max_tokens":  4000,
		"temperature": 0.2, // Baixa temperatura para extração factual
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
		"tools": []map[string]interface{}{
			{
				"name":         "extract_lab_results",
				"description":  "Extract structured lab results from OCR text",
				"input_schema": s.buildJSONSchema(),
			},
		},
		"tool_choice": map[string]string{
			"type": "tool",
			"name": "extract_lab_results",
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", s.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call Claude API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("claude api error %d: %s", resp.StatusCode, string(body))
	}

	var apiResp struct {
		Content []struct {
			Type  string          `json:"type"`
			Input json.RawMessage `json:"input"`
		} `json:"content"`
	}

	// Ler body inteiro para debug
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Debug: log resposta completa do Claude
	fmt.Printf("🤖 Claude API Response (first 2000 chars): %s\n", string(bodyBytes[:min(2000, len(bodyBytes))]))

	if err := json.Unmarshal(bodyBytes, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	// Extrair resultado do tool_use
	var result dto.AILabResultExtractionResponse
	for _, content := range apiResp.Content {
		if content.Type == "tool_use" {
			fmt.Printf("🔍 Tool use input (first 500 chars): %s\n", string(content.Input[:min(500, len(content.Input))]))
			if err := json.Unmarshal(content.Input, &result); err != nil {
				return nil, fmt.Errorf("failed to parse tool input: %v", err)
			}
			fmt.Printf("✅ Parsed result: laboratoryName=%s, labResults count=%d\n", result.LaboratoryName, len(result.LabResults))
			return &result, nil
		}
	}

	return nil, fmt.Errorf("no tool_use in response")
}

// buildPrompt - prompt otimizado para extração médica estruturada
func (s *AIService) buildPrompt(ocrText string, tests []dto.LabTestSummary) string {
	testsJSON, _ := json.MarshalIndent(tests, "", "  ")

	return fmt.Sprintf(`# TAREFA: Extração de Resultados Laboratoriais

Você receberá o texto OCR de um laudo médico em PDF. Extraia os dados estruturados seguindo as instruções abaixo.

## TEXTO OCR DO LAUDO
%s

## DEFINIÇÕES DE EXAMES DISPONÍVEIS (catálogo do sistema)
%s

## INSTRUÇÕES DETALHADAS

1. **Extrair informações gerais:**
   - Nome do laboratório
   - Data de coleta (formato ISO 8601: YYYY-MM-DD)
   - Data do resultado (formato ISO 8601: YYYY-MM-DD)

2. **Para cada resultado de exame:**
   - Extraia: nome do exame, tipo, resultado (numérico e/ou texto), unidade, faixa de referência, interpretação
   - Tente fazer MATCH com as definições do catálogo acima usando:
     * Nome do exame (case-insensitive, ignore acentos)
     * Código TUSS (se disponível no laudo)
     * Código LOINC (se disponível no laudo)
   - Se encontrar match: use o labTestDefinitionId correspondente e marque matched=true
   - Se NÃO encontrar match: deixe labTestDefinitionId=null e marque matched=false

3. **REGRAS IMPORTANTES:**
   - NUNCA invente dados que não estão no OCR
   - Se um campo não estiver presente no laudo, use null
   - Matching deve ser CASE-INSENSITIVE e tolerante a variações
   - Priorize matching por códigos (TUSS/LOINC) quando disponíveis
   - Se houver múltiplas variações de nome (ex: "Glicemia", "Glicose"), considere match
   - Para resultados qualitativos (ex: "Negativo", "Positivo"), use resultText
   - Para resultados quantitativos (ex: "95.5"), use resultNumeric
   - Mantenha a unidade original (ex: "mg/dL", "g/dL", "mmol/L")

4. **Tipos de exame (testType):**
   - biochemistry: exames bioquímicos (glicemia, ureia, creatinina, etc)
   - hematology: hemograma, coagulograma
   - immunology: sorologias, marcadores tumorais
   - microbiology: culturas, antibiogramas
   - urinalysis: urina tipo 1, urocultura
   - other: outros tipos

Retorne JSON estruturado conforme o schema fornecido.`, ocrText, string(testsJSON))
}

// buildJSONSchema - schema JSON para structured output (tool calling)
func (s *AIService) buildJSONSchema() map[string]interface{} {
	return map[string]interface{}{
		"type":     "object",
		"required": []string{"laboratoryName", "collectionDate", "labResults"},
		"properties": map[string]interface{}{
			"laboratoryName": map[string]string{
				"type":        "string",
				"description": "Nome do laboratório que realizou os exames",
			},
			"collectionDate": map[string]interface{}{
				"type":        "string",
				"pattern":     "^\\d{4}-\\d{2}-\\d{2}$",
				"description": "Data de coleta no formato ISO 8601 (YYYY-MM-DD)",
			},
			"resultDate": map[string]interface{}{
				"type":        "string",
				"pattern":     "^\\d{4}-\\d{2}-\\d{2}$",
				"description": "Data do resultado no formato ISO 8601 (YYYY-MM-DD)",
			},
			"labResults": map[string]interface{}{
				"type":        "array",
				"description": "Lista de resultados de exames extraídos",
				"items": map[string]interface{}{
					"type":     "object",
					"required": []string{"testName", "testType", "matched"},
					"properties": map[string]interface{}{
						"labTestDefinitionId": map[string]interface{}{
							"type":        []string{"string", "null"},
							"format":      "uuid",
							"description": "UUID da definição do exame (null se não matched)",
						},
						"testName": map[string]string{
							"type":        "string",
							"description": "Nome do exame conforme aparece no laudo",
						},
						"testType": map[string]string{
							"type":        "string",
							"description": "Tipo do exame: biochemistry, hematology, immunology, microbiology, urinalysis, other",
						},
						"resultNumeric": map[string]interface{}{
							"type":        []string{"number", "null"},
							"description": "Resultado numérico (para exames quantitativos)",
						},
						"resultText": map[string]interface{}{
							"type":        []string{"string", "null"},
							"description": "Resultado em texto (para exames qualitativos)",
						},
						"unit": map[string]interface{}{
							"type":        []string{"string", "null"},
							"description": "Unidade de medida (ex: mg/dL, g/dL)",
						},
						"referenceRange": map[string]interface{}{
							"type":        []string{"string", "null"},
							"description": "Faixa de referência conforme laudo",
						},
						"interpretation": map[string]interface{}{
							"type":        []string{"string", "null"},
							"description": "Interpretação ou observações do resultado",
						},
						"matched": map[string]string{
							"type":        "boolean",
							"description": "true se foi encontrado match com definição catalogada, false caso contrário",
						},
					},
				},
			},
		},
	}
}
