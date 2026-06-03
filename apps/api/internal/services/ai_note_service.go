package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Regras anti-alucinação (telemedicina). É o coração da segurança da feature:
// o modelo só pode afirmar o que está no transcript e NUNCA inventar exame físico.
const clinicalNoteSystemPrompt = `Você é um escriba clínico que estrutura a nota de uma TELECONSULTA a partir da transcrição diarizada (cada fala é de um "Falante"). Você NÃO é o médico e NÃO dá diagnóstico próprio.

REGRAS INVIOLÁVEIS:
1. Use SOMENTE o que está na transcrição. Nunca infira, complete ou adicione conhecimento médico externo. Seção não abordada na conversa = "não informado".
2. Esta é uma TELECONSULTA: NÃO houve exame físico presencial. O campo Objetivo só pode conter o que o PACIENTE relatou ou leu em voz alta (ex.: PA aferida em casa, peso, resultado de exame que trouxe), sempre prefixado por "refere". NUNCA escreva ausculta, palpação, PA aferida pelo médico ou qualquer achado de exame físico. Default do Objetivo = "não avaliado nesta teleconsulta".
3. Nunca invente valores numéricos, doses, datas ou resultados de exame. Se não foi dito, não escreva.
4. Sintomas, história e queixas vêm das falas do PACIENTE; hipóteses, condutas, prescrições e orientações vêm das falas do MÉDICO.
5. Preserve negações como negações ("nega febre"). Use "refere/relata" para o que o paciente disse; não afirme como fato objetivo medido.
6. Em "papeis", mapeie cada "Falante N" para "medico" ou "paciente", inferindo do contexto (quem pergunta/conduz vs quem relata sintomas).
7. Liste em "itens_ambiguos_para_revisao_medica" tudo que ficou incerto, contraditório ou que o médico precisa confirmar.

Português clínico, objetivo, sem travessões, sem fechos-slogan. Responda APENAS chamando a ferramenta com o JSON estruturado.`

// GenerateClinicalNoteFromTranscript gera a nota clínica estruturada (anamnese|soap)
// a partir do transcript diarizado, via Claude com tool_use (schema rígido). Retorna
// o JSON do tool_use (string) — o caller parseia conforme o template do formato.
func (s *AIService) GenerateClinicalNoteFromTranscript(transcript, format string) (string, error) {
	if !s.IsConfigured() {
		return "", ErrAINotConfigured
	}
	tmpl, ok := noteTemplate(format)
	if !ok {
		return "", fmt.Errorf("formato de nota inválido: %q", format)
	}
	model := s.noteModel
	if model == "" {
		model = s.model
	}

	props := map[string]interface{}{}
	required := []string{}
	for _, sec := range tmpl {
		props[sec.Key] = map[string]interface{}{
			"type":        "string",
			"description": sec.Hint + " Se não abordado na conversa: \"não informado\".",
		}
		required = append(required, sec.Key)
	}
	props["itens_ambiguos_para_revisao_medica"] = map[string]interface{}{
		"type":        "array",
		"items":       map[string]interface{}{"type": "string"},
		"description": "Pontos incertos, contraditórios ou que exigem confirmação do médico.",
	}
	props["papeis"] = map[string]interface{}{
		"type":                 "object",
		"description":          "Mapa de cada Falante para \"medico\" ou \"paciente\".",
		"additionalProperties": map[string]interface{}{"type": "string"},
	}
	required = append(required, "itens_ambiguos_para_revisao_medica")

	schema := map[string]interface{}{
		"type":       "object",
		"properties": props,
		"required":   required,
	}

	payload := map[string]interface{}{
		"model":       model,
		"max_tokens":  4096,
		"temperature": 0.2,
		"system":      clinicalNoteSystemPrompt,
		"messages": []map[string]string{
			{"role": "user", "content": "Transcrição diarizada da teleconsulta:\n\n" + transcript},
		},
		"tools": []map[string]interface{}{
			{
				"name":         "registrar_nota_clinica",
				"description":  "Registra a nota clínica estruturada extraída da teleconsulta.",
				"input_schema": schema,
			},
		},
		"tool_choice": map[string]string{"type": "tool", "name": "registrar_nota_clinica"},
	}
	return s.callClaudeToolUse(payload)
}

// callClaudeToolUse faz o POST /v1/messages e devolve o JSON do primeiro bloco
// tool_use. Espelha o tratamento de InterpretLabResult, sem logar PII.
func (s *AIService) callClaudeToolUse(payload map[string]interface{}) (string, error) {
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

	bodyBytes, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		// Não loga o body (pode refletir trechos do payload clínico).
		return "", fmt.Errorf("claude api error %d", resp.StatusCode)
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
	}
	if err := json.Unmarshal(bodyBytes, &apiResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}
	fmt.Printf("💰 Clinical note gen - Model: %s, Input: %d, Output: %d tokens\n",
		apiResp.Model, apiResp.Usage.InputTokens, apiResp.Usage.OutputTokens)

	for _, content := range apiResp.Content {
		if content.Type == "tool_use" {
			return string(content.Input), nil
		}
	}
	return "", fmt.Errorf("no tool_use in response")
}
