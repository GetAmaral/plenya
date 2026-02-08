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
// Retorna JSON string diretamente com exames extraídos
func (s *AIService) InterpretLabResult(
	ocrText string,
) (string, error) {
	prompt := s.buildPrompt(ocrText, nil)

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
		return "", fmt.Errorf("failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", s.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to call Claude API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("claude api error %d: %s", resp.StatusCode, string(body))
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
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	// Debug: log resposta completa do Claude
	fmt.Printf("🤖 Claude API Response (first 2000 chars): %s\n", string(bodyBytes[:min(2000, len(bodyBytes))]))

	if err := json.Unmarshal(bodyBytes, &apiResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	// Log token usage for cost tracking
	fmt.Printf("💰 Token Usage - Model: %s, Input: %d tokens, Output: %d tokens, Total: %d tokens\n",
		apiResp.Model, apiResp.Usage.InputTokens, apiResp.Usage.OutputTokens,
		apiResp.Usage.InputTokens+apiResp.Usage.OutputTokens)

	// Extrair resultado do tool_use
	for _, content := range apiResp.Content {
		if content.Type == "tool_use" {
			fmt.Printf("🔍 Tool use input (first 500 chars): %s\n", string(content.Input[:min(500, len(content.Input))]))

			// Retornar JSON diretamente como string
			return string(content.Input), nil
		}
	}

	return "", fmt.Errorf("no tool_use in response")
}

// buildPrompt - prompt otimizado para extração médica estruturada
// NÃO envia definições de exames - apenas extrai dados brutos do OCR
func (s *AIService) buildPrompt(ocrText string, tests []dto.LabTestSummary) string {
	// Ignorar tests - não fazemos matching
	_ = tests

	return fmt.Sprintf(`# TAREFA: Extrair Dados de Exames Laboratoriais

Analise o texto OCR abaixo e extraia TODOS os exames em formato JSON.

## TEXTO OCR DO LAUDO
%s

## INSTRUÇÕES

Para CADA exame encontrado, extraia estes campos:

**Obrigatórios:**
1. **nomeExame**: nome do exame como aparece no laudo
2. **resultado**: valor do resultado (número ou texto)

**Opcionais (OMITIR se não encontrar):**
3. **unidade**: unidade de medida (mg/dL, g/dL, etc)
4. **material**: material biológico (Soro, Sangue, Urina, etc)
5. **metodo**: método usado (Enzimático, ELISA, etc)

## REGRAS CRÍTICAS

- Extraia TODOS os exames (Hemograma completo = múltiplos exames separados)
- **OMITA campos opcionais se não encontrar** (não envie campo vazio para economizar tokens)
- Descarte valores de referência, interpretações, notas
- Números brasileiros: use ponto decimal (1.5 não 1,5)
- NUNCA invente dados

## ELETROFORESE DE PROTEÍNAS - REGRA ESPECIAL

Se encontrar "ELETROFORESE DE PROTEÍNAS" ou "PROTEIN ELECTROPHORESIS", extraia CADA fração como exame separado:

**IMPORTANTE:** Cada fração possui dois valores: percentual (%%) e concentração absoluta (g/dL).
**EXTRAIA APENAS o valor em g/dL** (ignore o valor em %%).

Exemplo no laudo:

    Albumina...............: 62,1 %%
    Albumina g/dL..........: 4,04 g/dL
    Alfa-1-globulina.......: 3,6 %%
    Alfa-1-globulina g/dL..: 0,23 g/dL

Deve extrair:

    [
      {"nomeExame": "Albumina", "resultado": "4.04", "unidade": "g/dL"},
      {"nomeExame": "Alfa-1-globulina", "resultado": "0.23", "unidade": "g/dL"}
    ]

**Frações a extrair (APENAS valores g/dL):**
- Albumina (g/dL)
- Alfa-1-globulina (g/dL) ou Alfa 1 (g/dL)
- Alfa-2-globulina (g/dL) ou Alfa 2 (g/dL)
- Beta-1-globulina (g/dL) ou Beta 1 (g/dL)
- Beta-2-globulina (g/dL) ou Beta 2 (g/dL)
- Gama-globulina (g/dL) ou Gama (g/dL)
- Relação A/G (valor numérico sem unidade)
- Proteínas totais (g/dL)

## MICROALBUMINÚRIA - REGRA ESPECIAL

Se encontrar "MICROALBUMINÚRIA" ou "MICROALBUMINURIA", extraia APENAS a relação calculada:

**IMPORTANTE:** O laudo geralmente apresenta 3 valores:
1. Microalbuminúria (valor absoluto em mg/L)
2. Creatinina urinária (valor absoluto em g/L ou mg/dL)
3. Relação Microalbuminuria/Creatinina (mg/g) ← **ESTE É O VALOR CLÍNICO IMPORTANTE**

**EXTRAIA APENAS a Relação Microalbuminuria/Creatinina** (ignore os valores individuais).

Exemplo no laudo:

    Microalbuminúria...........: 12,5 mg/L
    Creatinina urinária........: 200 mg/dL
    Relação Microalbuminuria/Creatinina: 0,625 mg/g

Deve extrair:

    [
      {"nomeExame": "Relação Microalbuminuria/Creatinina", "resultado": "0.625", "unidade": "mg/g"}
    ]

**NÃO extrair** os valores individuais de Microalbuminúria e Creatinina quando estiver no contexto do exame de Microalbuminúria (esses valores sozinhos não têm valor clínico para este exame específico).

## EXEMPLO DE SAÍDA

[
  {
    "nomeExame": "Glicose",
    "resultado": "95",
    "unidade": "mg/dL",
    "material": "Soro",
    "metodo": "Enzimático"
  },
  {
    "nomeExame": "Hemoglobina",
    "resultado": "14.5",
    "unidade": "g/dL"
  },
  {
    "nomeExame": "Albumina",
    "resultado": "4.04",
    "unidade": "g/dL",
    "material": "Soro"
  },
  {
    "nomeExame": "Alfa-1-globulina",
    "resultado": "0.23",
    "unidade": "g/dL"
  }
]`, ocrText)
}

// buildJSONSchema - schema JSON para structured output (tool calling)
func (s *AIService) buildJSONSchema() map[string]interface{} {
	return map[string]interface{}{
		"type":     "object",
		"required": []string{"exames"},
		"properties": map[string]interface{}{
			"exames": map[string]interface{}{
				"type":        "array",
				"description": "Lista de TODOS os exames extraídos do laudo",
				"items": map[string]interface{}{
					"type":     "object",
					"required": []string{"nomeExame", "resultado"},
					"properties": map[string]interface{}{
						"nomeExame": map[string]string{
							"type":        "string",
							"description": "Nome do exame conforme aparece no laudo (OBRIGATÓRIO)",
						},
						"resultado": map[string]string{
							"type":        "string",
							"description": "Valor do resultado (número ou texto) (OBRIGATÓRIO)",
						},
						"unidade": map[string]string{
							"type":        "string",
							"description": "Unidade de medida (mg/dL, g/dL, etc) - OMITIR se não encontrar",
						},
						"material": map[string]string{
							"type":        "string",
							"description": "Material biológico (Soro, Sangue, Urina, etc) - OMITIR se não encontrar",
						},
						"metodo": map[string]string{
							"type":        "string",
							"description": "Método usado (Enzimático, ELISA, etc) - OMITIR se não encontrar",
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
