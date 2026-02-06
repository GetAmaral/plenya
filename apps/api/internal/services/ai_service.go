package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
		"max_tokens":  8192, // Máximo permitido pelo Haiku (8192)
		"temperature": 0.2,   // Baixa temperatura para extração factual
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
		Usage struct {
			InputTokens  int `json:"input_tokens"`
			OutputTokens int `json:"output_tokens"`
		} `json:"usage"`
		Model string `json:"model"`
		ID    string `json:"id"`
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

	// Log token usage for cost tracking
	fmt.Printf("💰 Token Usage - Model: %s, Input: %d tokens, Output: %d tokens, Total: %d tokens\n",
		apiResp.Model, apiResp.Usage.InputTokens, apiResp.Usage.OutputTokens,
		apiResp.Usage.InputTokens+apiResp.Usage.OutputTokens)

	// Extrair resultado do tool_use
	for _, content := range apiResp.Content {
		if content.Type == "tool_use" {
			fmt.Printf("🔍 Tool use input (first 500 chars): %s\n", string(content.Input[:min(500, len(content.Input))]))

			// Tentar parse direto primeiro
			var result dto.AILabResultExtractionResponse
			if err := json.Unmarshal(content.Input, &result); err != nil {
				// Se falhar, pode ser que labResults veio como string JSON
				// Tentar parse com struct intermediária
				var rawResult struct {
					LaboratoryName string          `json:"laboratoryName"`
					CollectionDate string          `json:"collectionDate"`
					ResultDate     string          `json:"resultDate"`
					LabResults     json.RawMessage `json:"labResults"`
				}
				if err2 := json.Unmarshal(content.Input, &rawResult); err2 != nil {
					return nil, fmt.Errorf("failed to parse tool input: %v (also tried raw: %v)", err, err2)
				}

				result.LaboratoryName = rawResult.LaboratoryName
				result.CollectionDate = rawResult.CollectionDate
				result.ResultDate = rawResult.ResultDate

				// labResults pode ser array ou string com JSON
				labResultsBytes := []byte(rawResult.LabResults)
				// Se começa com aspas, é uma string JSON que precisa ser unescaped
				if len(labResultsBytes) > 0 && labResultsBytes[0] == '"' {
					var labResultsStr string
					if err := json.Unmarshal(labResultsBytes, &labResultsStr); err != nil {
						return nil, fmt.Errorf("failed to unmarshal labResults string: %v", err)
					}
					labResultsBytes = []byte(labResultsStr)
				}
				if err := json.Unmarshal(labResultsBytes, &result.LabResults); err != nil {
					return nil, fmt.Errorf("failed to parse labResults array: %v", err)
				}
			}

			fmt.Printf("✅ Parsed result: laboratoryName=%s, labResults count=%d\n", result.LaboratoryName, len(result.LabResults))
			return &result, nil
		}
	}

	return nil, fmt.Errorf("no tool_use in response")
}

// buildPrompt - prompt otimizado para extração médica estruturada
// NÃO envia definições de exames - apenas extrai dados brutos do OCR
func (s *AIService) buildPrompt(ocrText string, tests []dto.LabTestSummary) string {
	// Ignoramos tests - matching será feito localmente após extração
	_ = tests

	return fmt.Sprintf(`# TAREFA: Extrair TODOS os Resultados de Exames Laboratoriais

Analise o texto OCR abaixo e extraia TODOS os resultados de exames encontrados.

## TEXTO OCR DO LAUDO
%s

## INSTRUÇÕES CRÍTICAS

1. **EXTRAIA ABSOLUTAMENTE TODOS OS EXAMES** - não pule nenhum resultado
2. **Informações gerais do laudo:**
   - laboratoryName: nome do laboratório (procure no cabeçalho)
   - collectionDate: data de coleta (formato YYYY-MM-DD)
   - resultDate: data do resultado (formato YYYY-MM-DD)

3. **Para CADA exame encontrado, extraia:**
   - testName: nome exato do exame como aparece no laudo
   - testType: biochemistry, hematology, hormones, immunology, urinalysis, microbiology, ou other
   - resultNumeric: valor numérico (ex: 95.5) - use ponto decimal, não vírgula
   - resultText: resultado textual para exames qualitativos (ex: "Negativo", "Normal")
   - unit: unidade de medida (ex: "mg/dL", "g/dL", "mEq/L", "%%")
   - referenceRange: faixa de referência como texto (ex: "70-100 mg/dL")
   - interpretation: interpretação se houver (ex: "Normal", "Alterado")
   - matched: sempre false (matching será feito depois)
   - labTestDefinitionId: sempre null (será preenchido depois)

4. **IMPORTANTE:**
   - Extraia TODOS os exames, mesmo que pareçam repetidos
   - Hemograma inclui múltiplos parâmetros: Hemácias, Hemoglobina, Hematócrito, VCM, HCM, CHCM, RDW, Leucócitos, Plaquetas, etc - extraia CADA UM separadamente
   - Perfil Lipídico inclui: Colesterol Total, HDL, LDL, VLDL, Triglicerídeos - extraia CADA UM
   - Exames hormonais: TSH, T4, T3, Testosterona, Estradiol, Cortisol, etc
   - Números brasileiros usam vírgula como decimal (1,5 = 1.5)
   - NUNCA invente dados - extraia apenas o que está no texto`, ocrText)
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
				"description": "Lista de TODOS os resultados de exames extraídos do laudo",
				"items": map[string]interface{}{
					"type":     "object",
					"required": []string{"testName", "testType"},
					"properties": map[string]interface{}{
						"testName": map[string]string{
							"type":        "string",
							"description": "Nome do exame conforme aparece no laudo",
						},
						"testType": map[string]string{
							"type":        "string",
							"description": "Tipo: biochemistry, hematology, hormones, immunology, microbiology, urinalysis, other",
						},
						"resultNumeric": map[string]interface{}{
							"type":        []string{"number", "null"},
							"description": "Valor numérico (use ponto decimal: 95.5, não 95,5)",
						},
						"resultText": map[string]interface{}{
							"type":        []string{"string", "null"},
							"description": "Resultado textual (Negativo, Normal, Reagente, etc)",
						},
						"unit": map[string]interface{}{
							"type":        []string{"string", "null"},
							"description": "Unidade (mg/dL, g/dL, mEq/L, UI/mL, etc)",
						},
						"referenceRange": map[string]interface{}{
							"type":        []string{"string", "null"},
							"description": "Faixa de referência completa",
						},
					},
				},
			},
		},
	}
}

// MatchResultsWithDefinitions - faz matching local dos resultados extraídos com definições de exames
// Usa normalização de texto e altNames para matching flexível
func (s *AIService) MatchResultsWithDefinitions(
	results *dto.AILabResultExtractionResponse,
	testDefs []dto.LabTestSummary,
) {
	// Criar mapa de nomes normalizados para definições
	nameToDefID := make(map[string]string)
	for _, def := range testDefs {
		// Nome principal
		normalizedName := normalizeForMatching(def.Name)
		nameToDefID[normalizedName] = def.ID.String()

		// ShortName
		if def.ShortName != nil && *def.ShortName != "" {
			nameToDefID[normalizeForMatching(*def.ShortName)] = def.ID.String()
		}

		// Code
		if def.Code != "" {
			nameToDefID[normalizeForMatching(def.Code)] = def.ID.String()
		}
	}

	// Tentar match para cada resultado
	for i := range results.LabResults {
		result := &results.LabResults[i]
		normalizedTestName := normalizeForMatching(result.TestName)

		// Busca direta
		if defID, found := nameToDefID[normalizedTestName]; found {
			result.LabTestDefinitionID = &defID
			result.Matched = true
			continue
		}

		// Busca parcial (contém)
		for defName, defID := range nameToDefID {
			if strings.Contains(normalizedTestName, defName) || strings.Contains(defName, normalizedTestName) {
				// Só faz match se tiver pelo menos 5 caracteres em comum
				if len(defName) >= 5 || len(normalizedTestName) >= 5 {
					result.LabTestDefinitionID = &defID
					result.Matched = true
					break
				}
			}
		}
	}
}

// normalizeForMatching - normaliza texto para matching
func normalizeForMatching(text string) string {
	// Lowercase
	text = strings.ToLower(text)

	// Remove acentos (simplificado)
	replacer := strings.NewReplacer(
		"á", "a", "à", "a", "ã", "a", "â", "a",
		"é", "e", "è", "e", "ê", "e",
		"í", "i", "ì", "i", "î", "i",
		"ó", "o", "ò", "o", "õ", "o", "ô", "o",
		"ú", "u", "ù", "u", "û", "u",
		"ç", "c",
	)
	text = replacer.Replace(text)

	// Remove caracteres especiais
	text = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == ' ' {
			return r
		}
		return -1
	}, text)

	// Remove espaços extras
	text = strings.Join(strings.Fields(text), " ")

	return text
}
