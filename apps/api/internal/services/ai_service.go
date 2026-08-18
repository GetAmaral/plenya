package services

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/dto"
)

// ErrAIUpstream indica que a Claude API retornou erro não-200 ou indisponibilidade
// (timeout, rede, parse). Handlers mapeiam pra 502/504 sem vazar detalhes do provider.
var ErrAIUpstream = errors.New("ai: upstream Claude API failure")

// ErrAINotConfigured ocorre quando CLAUDE_API_KEY não está setada.
var ErrAINotConfigured = errors.New("ai: claude api key not configured")

// CompleteTextOptions controla parâmetros de uma chamada plain-text à Claude API.
// Quando Model vazio, usa s.model (default haiku). Timeout < 1s vira default 30s.
type CompleteTextOptions struct {
	Model       string
	MaxTokens   int
	Temperature float64
	Timeout     time.Duration
}

// CompleteText faz uma chamada plain-text (sem tools) à Claude Messages API e retorna
// o texto da resposta. Pra features de produto (resumo de conversa, sugestão de resposta,
// etc.) — não pra extração estruturada (use as funções com tool_use).
//
// LGPD: NÃO loga prompt nem resposta — só metadata (tokens, model, latência).
func (s *AIService) CompleteText(ctx context.Context, prompt string, opts CompleteTextOptions) (string, error) {
	if s.apiKey == "" {
		return "", ErrAINotConfigured
	}

	model := opts.Model
	if model == "" {
		model = s.model
	}
	maxTokens := opts.MaxTokens
	if maxTokens <= 0 {
		maxTokens = 1024
	}
	temperature := opts.Temperature
	if temperature == 0 {
		temperature = 0.4
	}
	timeout := opts.Timeout
	if timeout < time.Second {
		timeout = 30 * time.Second
	}

	payload := map[string]interface{}{
		"model":       model,
		"max_tokens":  maxTokens,
		"temperature": temperature,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("%w: marshal: %v", ErrAIUpstream, err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(reqCtx, "POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("%w: build request: %v", ErrAIUpstream, err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", s.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	start := time.Now()
	// Cliente local com timeout customizado — não reusa s.httpClient (3min) pra
	// respeitar Timeout passado via opts (default 30s).
	client := &http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	if err != nil {
		// Distingue timeout/cancelamento explicitamente pro handler retornar 504.
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(reqCtx.Err(), context.DeadlineExceeded) {
			return "", fmt.Errorf("%w: timeout after %s", ErrAIUpstream, timeout)
		}
		return "", fmt.Errorf("%w: %v", ErrAIUpstream, err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("%w: read body: %v", ErrAIUpstream, err)
	}

	if resp.StatusCode != 200 {
		// LGPD: NÃO logar conteúdo. Body do erro Claude pode ecoar trecho do prompt.
		fmt.Printf("⚠️  Claude CompleteText status=%d size=%d\n", resp.StatusCode, len(respBytes))
		return "", fmt.Errorf("%w: status %d", ErrAIUpstream, resp.StatusCode)
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
		Model      string `json:"model"`
		StopReason string `json:"stop_reason"`
	}
	if err := json.Unmarshal(respBytes, &apiResp); err != nil {
		return "", fmt.Errorf("%w: decode: %v", ErrAIUpstream, err)
	}

	// LGPD: log apenas metadata (model, tokens, latência) — nunca conteúdo.
	fmt.Printf("💬 Claude CompleteText - model=%s in=%d out=%d latency=%dms\n",
		apiResp.Model, apiResp.Usage.InputTokens, apiResp.Usage.OutputTokens,
		time.Since(start).Milliseconds())

	var sb strings.Builder
	for _, c := range apiResp.Content {
		if c.Type == "text" {
			sb.WriteString(c.Text)
		}
	}
	out := strings.TrimSpace(sb.String())
	if out == "" {
		return "", fmt.Errorf("%w: empty response", ErrAIUpstream)
	}
	return out, nil
}

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
	labModel   string
	noteModel  string
	httpClient *http.Client
	// labClient é usado só pela extração de laudo, que pode gerar dezenas de milhares de
	// tokens de saída (ver labExtractionMaxTokens). Roda no worker de processing jobs, fora
	// do ciclo de request do usuário, então o timeout largo não segura ninguém.
	labClient *http.Client
}

// NewAIService cria uma nova instância do serviço de IA
func NewAIService(cfg *config.Config) *AIService {
	return &AIService{
		apiKey:     cfg.Claude.APIKey,
		model:      cfg.Claude.Model,
		labModel:   cfg.Claude.LabModel,
		noteModel:  cfg.Claude.NoteModel,
		httpClient: &http.Client{Timeout: 180 * time.Second}, // 3 minutos para processar laudos grandes
		labClient:  &http.Client{Timeout: 10 * time.Minute},
	}
}

// IsConfigured retorna true se a CLAUDE_API_KEY está setada.
func (s *AIService) IsConfigured() bool { return s.apiKey != "" }

// labExtractionMaxTokens é o teto de saída da extração de laudo. Um check-up completo
// (hemograma + bioquímica + urina + imagem) passa fácil de 200 exames, e cada exame custa
// ~40-60 tokens no JSON do tool_use. O valor antigo (8192, herdado do Haiku 3) truncava a
// resposta no meio do array: a API devolve stop_reason=max_tokens e um input parcial que
// vira "0 exames" — silenciosamente, já que tool_use com JSON incompleto não dá erro HTTP.
// Os modelos atuais aceitam até 128k de saída e a chamada é em streaming (sem pressão de
// timeout), então sobra folga: 32k cobre laudo grande + os tokens de thinking, que também
// contam aqui.
const labExtractionMaxTokens = 32000

// labExtractionEffort controla quanto o modelo raciocina antes de responder. Ler laudo é
// transcrição estruturada, não raciocínio aberto: "medium" mantém o cuidado com os casos
// especiais do prompt (eletroforese, leucograma, urina) sem gastar milhares de tokens de
// thinking.
const labExtractionEffort = "medium"

// ErrAITruncated indica que a resposta da IA bateu no teto de max_tokens e veio incompleta.
// Não adianta repetir a mesma chamada (é determinístico) — o laudo precisa ser dividido.
var ErrAITruncated = errors.New("ai: resposta truncada por max_tokens")

// InterpretLabResult - interpreta laudo médico via Claude API com structured output
// Retorna JSON string diretamente com exames extraídos.
//
// Roda em streaming (SSE): com laudo grande a resposta passa de 10k tokens e uma chamada
// não-streaming ficaria minutos com a conexão parada, sujeita a timeout de infra no meio
// do caminho. O streaming também dá o stop_reason final de forma confiável.
func (s *AIService) InterpretLabResult(
	ocrText string,
) (string, error) {
	prompt := s.buildPrompt(ocrText, nil)

	// Sem "temperature": os modelos atuais (Opus 5 e a família 4.7+) rejeitam parâmetros de
	// sampling com 400. O que garante extração factual aqui é o schema do tool + o prompt.
	payload := map[string]interface{}{
		"model":      s.labModel,
		"max_tokens": labExtractionMaxTokens,
		"stream":     true,
		"output_config": map[string]interface{}{
			"effort": labExtractionEffort,
		},
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

	start := time.Now()
	resp, err := s.labClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to call Claude API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		// LGPD: o body de erro da Claude pode ecoar trecho do prompt (dado clínico).
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("⚠️  Claude InterpretLabResult status=%d size=%d\n", resp.StatusCode, len(body))
		return "", fmt.Errorf("claude api error %d", resp.StatusCode)
	}

	res, err := readToolUseStream(resp.Body, "extract_lab_results")
	if err != nil {
		return "", err
	}

	// LGPD: só metadata — nunca o conteúdo extraído (valores de exames).
	fmt.Printf("💰 Token Usage - Model: %s, Input: %d tokens, Output: %d tokens, Stop: %s, Latency: %ds\n",
		res.model, res.inputTokens, res.outputTokens, res.stopReason, int(time.Since(start).Seconds()))

	// Resposta truncada: o tool_use vem com JSON incompleto, o que aparecia downstream como
	// "0 exames". Falha explícita aqui em vez de gravar um laudo pela metade.
	if res.stopReason == "max_tokens" {
		return "", fmt.Errorf("%w (limite de %d tokens de saída)", ErrAITruncated, labExtractionMaxTokens)
	}

	if res.input == "" {
		return "", fmt.Errorf("no tool_use in response (stop_reason=%s)", res.stopReason)
	}
	fmt.Printf("🔍 Tool use input: %d bytes\n", len(res.input))
	return res.input, nil
}

// streamResult é o que sobra de um stream SSE da Messages API depois de descartar
// thinking/texto: o JSON do tool_use pedido + a metadata de billing e parada.
type streamResult struct {
	input        string
	model        string
	stopReason   string
	inputTokens  int
	outputTokens int
}

// readToolUseStream consome o SSE da Messages API e remonta o input JSON do tool_use de
// nome toolName (que chega em pedaços via input_json_delta). Blocos de thinking e texto são
// ignorados de propósito — com tool_choice forçado, o dado está no tool_use.
func readToolUseStream(body io.Reader, toolName string) (streamResult, error) {
	var res streamResult
	var jsonParts strings.Builder
	capturing := false

	sc := bufio.NewScanner(body)
	// Um input_json_delta é pequeno, mas um bloco de thinking pode vir gordo numa linha só.
	sc.Buffer(make([]byte, 0, 64*1024), 8*1024*1024)

	for sc.Scan() {
		line := sc.Text()
		if !strings.HasPrefix(line, "data: ") {
			continue // linhas "event:" e brancas não carregam payload
		}

		var ev struct {
			Type         string `json:"type"`
			Index        int    `json:"index"`
			ContentBlock struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"content_block"`
			Delta struct {
				Type        string `json:"type"`
				PartialJSON string `json:"partial_json"`
				StopReason  string `json:"stop_reason"`
			} `json:"delta"`
			Message struct {
				Model string `json:"model"`
				Usage struct {
					InputTokens  int `json:"input_tokens"`
					OutputTokens int `json:"output_tokens"`
				} `json:"usage"`
			} `json:"message"`
			Usage struct {
				OutputTokens int `json:"output_tokens"`
			} `json:"usage"`
			Error struct {
				Type    string `json:"type"`
				Message string `json:"message"`
			} `json:"error"`
		}
		if err := json.Unmarshal([]byte(strings.TrimPrefix(line, "data: ")), &ev); err != nil {
			continue // evento desconhecido/parcial: ignora em vez de derrubar a extração
		}

		switch ev.Type {
		case "message_start":
			res.model = ev.Message.Model
			res.inputTokens = ev.Message.Usage.InputTokens
		case "content_block_start":
			capturing = ev.ContentBlock.Type == "tool_use" && ev.ContentBlock.Name == toolName
		case "content_block_delta":
			if capturing && ev.Delta.Type == "input_json_delta" {
				jsonParts.WriteString(ev.Delta.PartialJSON)
			}
		case "content_block_stop":
			capturing = false
		case "message_delta":
			if ev.Delta.StopReason != "" {
				res.stopReason = ev.Delta.StopReason
			}
			if ev.Usage.OutputTokens > 0 {
				res.outputTokens = ev.Usage.OutputTokens
			}
		case "error":
			// LGPD: a mensagem da API pode ecoar o prompt — só o tipo do erro vai pro log.
			return res, fmt.Errorf("%w: stream error (%s)", ErrAIUpstream, ev.Error.Type)
		}
	}
	if err := sc.Err(); err != nil {
		return res, fmt.Errorf("%w: stream interrompido: %v", ErrAIUpstream, err)
	}

	res.input = jsonParts.String()
	return res, nil
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

### Dados do laudo (no topo, OMITIR se não encontrar)
- **laboratorio**: nome do laboratório que emitiu o laudo (ex: "Laboratório Oswaldo Cruz").
- **dataColeta**: data da COLETA do material, no formato **YYYY-MM-DD**. Use a data da COLETA,
  não a data de emissão/liberação/impressão do laudo. Se houver datas diferentes por exame,
  use a coleta mais comum/representativa.

### Para CADA exame encontrado, extraia estes campos:

**Obrigatórios:**
1. **nomeExame**: nome do exame como aparece no laudo
2. **resultado**: valor do resultado (número ou texto)

**Opcionais (OMITIR se não encontrar):**
3. **unidade**: unidade de medida (mg/dL, g/dL, etc)
4. **material**: material biológico/espécime (Soro, Sangue, Plasma, Urina, Fezes, etc)
5. **metodo**: método usado (Enzimático, ELISA, ECLIA, HPLC, etc)
6. **valorReferencia**: faixa de referência impressa do lab para ESTE exame, como texto
   (ex: "70 a 99 mg/dL", "Inferior a 190 mg/dL", "Negativo"). NÃO é o resultado.
7. **dataColetaExame**: data da COLETA deste exame (formato YYYY-MM-DD), SEMPRE que o laudo
   trouxer "Coletado em:" para o exame — mesmo que igual à coleta geral. Usada para renderizar
   a tabela por data do RESULTADO.

## REGRAS CRÍTICAS

- Extraia TODOS os exames (Hemograma completo = múltiplos exames separados)
- **OMITA campos opcionais se não encontrar** (não envie campo vazio para economizar tokens)
- Capture a faixa de referência em **valorReferencia** (não no resultado); descarte
  interpretações/comentários/notas e os "Resultados Anteriores" (histórico)
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

## HEMOGRAMA / LEUCOGRAMA (diferencial) - REGRA ESPECIAL

No LEUCOGRAMA, cada tipo celular costuma ter DOIS valores na MESMA linha: percentual (%%)
e contagem ABSOLUTA (/mm3 ou /µL). O valor clínico para o escore é o **ABSOLUTO**.
Extraia DOIS exames por célula:
- "<Célula> (absoluto)" com a contagem absoluta (/mm3) ← OBRIGATÓRIO
- "<Célula> (percentual)" com o percentual

Trate "Segmentados" como **Neutrófilos**.

Exemplo no laudo:

    Segmentados.....: 41,17 %%   2.548,42 /mm3
    Linfócitos......: 42,73 %%   2.644,99 /mm3
    Monócitos.......: 12,41 %%   768,18 /mm3

Deve extrair:

    [
      {"nomeExame": "Neutrófilos (absoluto)", "resultado": "2548.42", "unidade": "/mm3"},
      {"nomeExame": "Neutrófilos (percentual)", "resultado": "41.17", "unidade": "%%"},
      {"nomeExame": "Linfócitos (absoluto)", "resultado": "2644.99", "unidade": "/mm3"},
      {"nomeExame": "Linfócitos (percentual)", "resultado": "42.73", "unidade": "%%"},
      {"nomeExame": "Monócitos (absoluto)", "resultado": "768.18", "unidade": "/mm3"},
      {"nomeExame": "Monócitos (percentual)", "resultado": "12.41", "unidade": "%%"}
    ]

Aplica-se a Neutrófilos/Segmentados, Linfócitos, Monócitos, Eosinófilos, Basófilos e
Bastonetes. Se a linha tiver SÓ o %% (sem absoluto), extraia apenas o %%.

## ROTINA DE URINA / EAS - REGRA ESPECIAL

Para componentes do exame de urina (EAS / Rotina de urina / sedimento), use o nome URINÁRIO
explícito e SEMPRE informe material="Urina" — evita confundir com o exame de SANGUE de mesmo
nome (glicose, proteínas, hemoglobina têm versão sangue E urina):
- Glicose → "Glicose Urinária"; Proteínas → "Proteínas Urinárias"; Hemoglobina → "Hemoglobina Urinária"
- Corpos cetônicos → "Corpos Cetônicos Urinários"; Bilirrubina → "Bilirrubina Urinária"
- Densidade → "Densidade Urinária"; pH → "pH Urinário"; Nitrito; Urobilinogênio (mantêm o nome)
- Leucócitos/Hemácias/Células epiteliais/Bactérias/Leveduras/Muco → "<Componente> - Sedimento"

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
			"laboratorio": map[string]string{
				"type":        "string",
				"description": "Nome do laboratório que emitiu o laudo (ex: Oswaldo Cruz, Fleury) - OMITIR se não encontrar",
			},
			"dataColeta": map[string]string{
				"type":        "string",
				"description": "Data da COLETA do material no formato YYYY-MM-DD (NÃO a data de emissão/liberação/impressão) - OMITIR se não encontrar",
			},
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
							"description": "Método usado (Enzimático, ELISA, ECLIA, HPLC, etc) - OMITIR se não encontrar",
						},
						"valorReferencia": map[string]string{
							"type":        "string",
							"description": "Faixa de referência impressa do lab para este exame (ex: '70 a 99 mg/dL', 'Negativo') - NÃO é o resultado - OMITIR se não encontrar",
						},
						"dataColetaExame": map[string]string{
							"type":        "string",
							"description": "Data da coleta DESTE exame (YYYY-MM-DD), SEMPRE que o laudo trouxer 'Coletado em:' para o exame (mesmo se igual à geral) - OMITIR só se não houver",
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

// ExtractArticleMetadata - extrai metadados de artigo científico da primeira página
// Usa Claude Haiku para extração estruturada inteligente
func (s *AIService) ExtractArticleMetadata(firstPageText string) (map[string]interface{}, error) {
	if s.apiKey == "" {
		return nil, fmt.Errorf("Claude API key not configured")
	}

	prompt := fmt.Sprintf(`# TAREFA: Extrair Metadados de Artigo Científico

Analise o texto da primeira página do PDF abaixo e extraia os metadados bibliográficos.

## TEXTO DA PRIMEIRA PÁGINA
%s

## INSTRUÇÕES

Extraia os seguintes campos (omita campos que não encontrar):

**Obrigatórios:**
1. **title**: Título completo do artigo
2. **authors**: Nomes dos autores (formato: "Sobrenome A, Sobrenome B, Sobrenome C")
3. **journal**: Nome da revista/journal

**Opcionais:**
4. **publicationDate**: Data de publicação (formato YYYY-MM-DD, ou apenas YYYY se não tiver mês/dia)
5. **doi**: DOI do artigo (apenas o identificador, ex: "10.1038/s41586-024-07146-0")
6. **pmid**: PubMed ID (apenas números)
7. **issn**: ISSN da revista
8. **abstract**: Resumo/abstract do artigo (se presente na primeira página)
9. **keywords**: Array de palavras-chave (se presentes)
10. **articleType**: Tipo do artigo - escolha entre:
    - "research_article" (artigo de pesquisa original)
    - "review" (revisão narrativa)
    - "meta_analysis" (meta-análise, revisão sistemática)
    - "case_study" (estudo de caso, relato de caso)
    - "clinical_trial" (ensaio clínico)
    - "editorial" (editorial, comentário)
    - "letter" (carta, correspondência)

## REGRAS CRÍTICAS

- **NUNCA invente dados** - omita campos se não encontrar
- Para **authors**: se houver muitos autores, liste até 10 principais + "et al."
- Para **doi**: extraia apenas o identificador (sem URL completa)
- Para **publicationDate**: tente inferir o ano mesmo que não esteja explícito
- Para **articleType**: analise o conteúdo e tipo de estudo para classificar
- Se o texto estiver em português, traduza title/abstract para inglês se possível

Retorne um JSON com os campos encontrados.`, firstPageText)

	payload := map[string]interface{}{
		"model":       s.model,
		"max_tokens":  4096,
		"temperature": 0.2,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
		"tools": []map[string]interface{}{
			{
				"name":         "extract_article_metadata",
				"description":  "Extract bibliographic metadata from scientific article first page",
				"input_schema": s.buildArticleMetadataSchema(),
			},
		},
		"tool_choice": map[string]string{
			"type": "tool",
			"name": "extract_article_metadata",
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
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	if err := json.Unmarshal(bodyBytes, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	fmt.Printf("💰 Article Metadata Extraction - Model: %s, Input: %d tokens, Output: %d tokens\n",
		apiResp.Model, apiResp.Usage.InputTokens, apiResp.Usage.OutputTokens)

	// Extrair resultado do tool_use
	for _, content := range apiResp.Content {
		if content.Type == "tool_use" {
			var result map[string]interface{}
			if err := json.Unmarshal(content.Input, &result); err != nil {
				return nil, fmt.Errorf("failed to parse tool input: %v", err)
			}
			return result, nil
		}
	}

	return nil, fmt.Errorf("no tool_use in response")
}

// ExtractTableOfContents extrai títulos de capítulos de um sumário de livro via Claude Haiku
// Retorna lista de títulos de capítulos encontrados no texto
func (s *AIService) ExtractTableOfContents(text string) ([]string, error) {
	if s.apiKey == "" {
		return nil, fmt.Errorf("Claude API key not configured")
	}

	prompt := fmt.Sprintf(`Analise o texto abaixo que pode conter um sumário ou índice de livro.

Extraia os títulos dos capítulos principais (não subseções, não sub-capítulos — apenas capítulos de nível 1).

TEXTO:
%s

Retorne apenas os títulos dos capítulos como array JSON de strings.
Se não houver sumário ou capítulos identificáveis, retorne um array vazio [].
Retorne no mínimo 3 capítulos para ser útil.

Exemplo de formato: ["Capítulo 1: Introdução", "Capítulo 2: Metabolismo", "Capítulo 3: Conclusão"]

Retorne APENAS o JSON, sem explicações.`, text)

	payload := map[string]interface{}{
		"model":       "claude-haiku-4-5-20251001",
		"max_tokens":  1024,
		"temperature": 0.1,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", s.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("claude api error %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var apiResp struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	}

	if err := json.Unmarshal(bodyBytes, &apiResp); err != nil {
		return nil, err
	}

	for _, c := range apiResp.Content {
		if c.Type == "text" {
			// Extrair array JSON do texto
			start := strings.Index(c.Text, "[")
			end := strings.LastIndex(c.Text, "]")
			if start >= 0 && end > start {
				jsonStr := c.Text[start : end+1]
				var titles []string
				if err := json.Unmarshal([]byte(jsonStr), &titles); err == nil {
					return titles, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("resposta Claude sem array JSON válido")
}

// buildArticleMetadataSchema - schema JSON para extração de metadados de artigos
func (s *AIService) buildArticleMetadataSchema() map[string]interface{} {
	return map[string]interface{}{
		"type":     "object",
		"required": []string{"title", "authors", "journal"},
		"properties": map[string]interface{}{
			"title": map[string]string{
				"type":        "string",
				"description": "Full article title",
			},
			"authors": map[string]string{
				"type":        "string",
				"description": "Authors in format: 'LastName A, LastName B, LastName C' or 'LastName A, et al.'",
			},
			"journal": map[string]string{
				"type":        "string",
				"description": "Journal/publication name",
			},
			"publicationDate": map[string]string{
				"type":        "string",
				"description": "Publication date in YYYY-MM-DD format (or YYYY if only year available)",
			},
			"doi": map[string]string{
				"type":        "string",
				"description": "DOI identifier (without URL prefix)",
			},
			"pmid": map[string]string{
				"type":        "string",
				"description": "PubMed ID (numbers only)",
			},
			"issn": map[string]string{
				"type":        "string",
				"description": "ISSN of the journal",
			},
			"abstract": map[string]string{
				"type":        "string",
				"description": "Article abstract/summary if present on first page",
			},
			"keywords": map[string]interface{}{
				"type":        "array",
				"description": "Keywords if present",
				"items": map[string]string{
					"type": "string",
				},
			},
			"articleType": map[string]interface{}{
				"type": "string",
				"enum": []string{
					"research_article",
					"review",
					"meta_analysis",
					"case_study",
					"clinical_trial",
					"editorial",
					"letter",
				},
				"description": "Type of article based on content",
			},
		},
	}
}

// ExtractRequestedExams - extrai a LISTA DE NOMES de exames SOLICITADOS num pedido (foto/PDF de
// outro médico). Sem resultados (é um pedido, não um laudo). Usado p/ dedup ao importar.
func (s *AIService) ExtractRequestedExams(ocrText string) ([]string, error) {
	if !s.IsConfigured() {
		return nil, fmt.Errorf("CLAUDE_API_KEY não configurada")
	}
	prompt := "Abaixo está o texto (OCR) de um PEDIDO/SOLICITAÇÃO de exames médicos emitido por outro " +
		"profissional. Não há resultados — apenas a lista de exames pedidos. Extraia SOMENTE os nomes dos " +
		"exames solicitados (laboratoriais e de imagem), um por item, sem numeração, sem dados do paciente, " +
		"sem cabeçalho de laboratório. Mantenha o nome do exame como aparece.\n\n--- TEXTO ---\n" + ocrText

	schema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"exames": map[string]interface{}{
				"type":        "array",
				"items":       map[string]interface{}{"type": "string"},
				"description": "Nomes dos exames solicitados",
			},
		},
		"required": []string{"exames"},
	}
	payload := map[string]interface{}{
		"model":       s.model,
		"max_tokens":  4096,
		"temperature": 0.1,
		"messages":    []map[string]string{{"role": "user", "content": prompt}},
		"tools": []map[string]interface{}{{
			"name": "extract_requested_exams", "description": "Lista de exames solicitados", "input_schema": schema,
		}},
		"tool_choice": map[string]string{"type": "tool", "name": "extract_requested_exams"},
	}
	jsonPayload, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", s.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("claude api error %d: %s", resp.StatusCode, string(body))
	}
	var apiResp struct {
		Content []struct {
			Type  string          `json:"type"`
			Input json.RawMessage `json:"input"`
		} `json:"content"`
	}
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, err
	}
	for _, c := range apiResp.Content {
		if c.Type == "tool_use" {
			var parsed struct {
				Exames []string `json:"exames"`
			}
			if err := json.Unmarshal(c.Input, &parsed); err != nil {
				return nil, err
			}
			return parsed.Exames, nil
		}
	}
	return nil, fmt.Errorf("no tool_use in response")
}
