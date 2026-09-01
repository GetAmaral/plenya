package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// A chamada ao modelo para editar o rascunho da devolutiva.
//
// Segue o padrão da casa para saída estruturada — `tool_use` forçado, como em `InterpretLabResult`
// e no escriba de teleconsulta — com três diferenças que valem por documentação:
//
//  1. NÃO manda `temperature`. Opus 5 e Sonnet 5 respondem 400 com parâmetro de sampling junto de
//     tool forçado. O código atual do resto do serviço só funciona porque o modelo default ainda é
//     um 4.6; quem trocar o default e não souber disto quebra as outras chamadas.
//  2. Usa cache de prompt no bloco de sistema. O dossiê é congelado, portanto byte-idêntico entre
//     turnos, e é justamente por isso que ele é congelado: com dossiê vivo o cache seria invalidado
//     a cada turno. A partir do segundo turno, `cache_read_input_tokens` tem que vir maior que zero.
//  3. O que muda entre turnos (o médico editou um slide à mão, uma sugestão foi recusada) vai como
//     mensagem de sistema NO FIM de `messages`, não no bloco de topo, para não invalidar o prefixo.

// planSystemPrompt — as regras invioláveis.
//
// O prompt PEDE que o modelo não invente número. A diferença desta feature para o escriba de
// teleconsulta é que aqui dá para VERIFICAR: os números saem de um conjunto fechado, e o servidor
// confere cada um contra o dossiê congelado antes de aplicar. O prompt é a primeira barreira, não
// a única, e é escrito sabendo disso.
const planSystemPrompt = `Você ajuda um médico a escrever a devolutiva de exames que o PACIENTE vai ler. Você não é o médico e não dá diagnóstico próprio.

REGRAS INVIOLÁVEIS:
1. Não escreva número que não esteja no DOSSIÊ. Todo número que aparecer em qualquer texto que você escrever tem que ser declarado em "numerals", com a origem no dossiê. Número sem origem: não escreva o número.
2. Os campos "code", "segments" e "history" da régua vêm do dossiê. Copie exatamente ou não mexa. Nunca invente faixa nem medida.
3. Os ids de slide são opacos. Copie os que estão no rascunho; nunca invente um id.
4. Operação de texto NÃO muda número. Se a reescrita mexe em algum número, unidade ou dose, declare a classe como "numeric".
5. Voz de paciente, prosa clínica em português do Brasil, sem jargão. Sem travessão, sem "Não é X. É Y.", sem fecho-slogan, sem ícone em lista, sem preço, sem marca comercial.
6. Título de slide é uma AFIRMAÇÃO, não um rótulo: "A ferritina dobrou em dois anos", não "Ferritina".
7. Se o médico pedir algo que o dossiê não sustenta, diga isso em "reply" e devolva "operations" vazio. Silêncio é melhor que invenção.

O campo "reply" é para o MÉDICO ler: o que você fez, ou por que não fez. Nunca é conteúdo de slide.`

// planEditToolSchema — o contrato de saída.
//
// A saída é uma lista de OPERAÇÕES, não um deck reescrito. Operação torna o diff trivial e
// reversível, e impede que um pedido pequeno reescreva o deck inteiro em silêncio.
func planEditToolSchema() map[string]any {
	return map[string]any{
		"type":                 "object",
		"additionalProperties": false,
		"required":             []string{"reply", "operations"},
		"properties": map[string]any{
			"reply": map[string]any{
				"type":        "string",
				"description": "O que você fez ou por que não fez, em português, para o médico ler. Nunca é conteúdo do slide.",
			},
			"operations": map[string]any{
				"type":     "array",
				"maxItems": 12,
				"items": map[string]any{
					"type":                 "object",
					"additionalProperties": false,
					"required":             []string{"op", "class", "rationale"},
					"properties": map[string]any{
						"op": map[string]any{"enum": []string{"add", "edit", "remove", "reorder"}},
						"class": map[string]any{
							"enum":        []string{"text", "numeric", "structural"},
							"description": "text = nenhum número novo entra. numeric = toca número, unidade, dose ou régua. structural = adiciona, remove ou reordena slide. Na dúvida, numeric.",
						},
						"slideId":      map[string]any{"type": "string", "description": "Id do slide alvo, copiado do rascunho."},
						"afterSlideId": map[string]any{"type": "string", "description": "Só em add: insere depois deste slide. Vazio = no início."},
						"path": map[string]any{
							"type":        "string",
							"description": "Só em edit. Do vocabulário permitido: title, eyebrow, lede, kicker, source, punch, cards[i].kicker, cards[i].body, rulers[i].display, rulers[i].sub, rulers[i].note, rulers[i].axis, summary.cards[i].lines[j].value, table.rows[i].cells[j], takeaway.note, e afins.",
						},
						"value":     map[string]any{"description": "Só em edit. O novo valor do campo."},
						"slide":     map[string]any{"type": "object", "description": "Só em add. O slide completo, sem id."},
						"order":     map[string]any{"type": "array", "items": map[string]any{"type": "string"}, "description": "Só em reorder. TODOS os ids, na ordem nova."},
						"rationale": map[string]any{"type": "string", "description": "Por que, em uma frase, para o médico decidir."},
						"numerals": map[string]any{
							"type":        "array",
							"description": "TODO número presente no que você escreveu, com a origem no dossiê. Número sem origem: não escreva o número.",
							"items": map[string]any{
								"type":                 "object",
								"additionalProperties": false,
								"required":             []string{"numeral", "source"},
								"properties": map[string]any{
									"numeral": map[string]any{"type": "string", "description": `O número como você escreveu: "96", "1,023", "112 g".`},
									"source":  map[string]any{"type": "string", "description": "De onde veio no dossiê, ex.: ruler:PLNFERR:history:2026-02-06."},
								},
							},
						},
					},
				},
			},
		},
	}
}

// PlanEditTurn — um turno anterior da conversa.
type PlanEditTurn struct {
	Role string
	Body string
}

// PlanEditRequest — tudo o que a chamada precisa.
type PlanEditRequest struct {
	DossierJSON   string // já podado; ver PodaDossieParaPrompt
	SlidesJSON    string
	Turns         []PlanEditTurn
	StateNote     string // o que mudou desde o último turno
	Instruction   string
	PromptVersion string
	Model         string
}

// AICallMeta — metadados da chamada, para gravar na mensagem e na revisão.
type AICallMeta struct {
	Model           string
	StopReason      string
	InputTokens     int
	CacheReadTokens int
	OutputTokens    int
	LatencyMs       int
}

// PlanEditResult — o que o modelo devolveu, ainda sem triagem.
type PlanEditResult struct {
	Reply      string        `json:"reply"`
	Operations []PlanModelOp `json:"operations"`
}

// PlanModelOp — a operação como o modelo a declara. A classe declarada NÃO é autoritativa: o
// servidor reclassifica, e discordância vira registro.
type PlanModelOp struct {
	Op            string          `json:"op"`
	SlideID       string          `json:"slideId,omitempty"`
	AfterSlideID  string          `json:"afterSlideId,omitempty"`
	Path          string          `json:"path,omitempty"`
	Value         json.RawMessage `json:"value,omitempty"`
	Slide         json.RawMessage `json:"slide,omitempty"`
	Order         []string        `json:"order,omitempty"`
	DeclaredClass string          `json:"class"`
	Rationale     string          `json:"rationale,omitempty"`
	Numerals      []struct {
		Numeral string `json:"numeral"`
		Source  string `json:"source"`
	} `json:"numerals,omitempty"`
}

// EditPatientPlan chama o modelo e devolve a resposta crua mais os metadados.
func (s *AIService) EditPatientPlan(req PlanEditRequest) (*PlanEditResult, AICallMeta, error) {
	var meta AICallMeta
	if !s.IsConfigured() {
		return nil, meta, ErrAINotConfigured
	}

	// O bloco de sistema é o prefixo cacheado: regras + dossiê. Nada volátil entra aqui.
	sistema := []map[string]any{
		{"type": "text", "text": planSystemPrompt},
		{
			"type": "text",
			"text": "DOSSIÊ (prontuário compilado e congelado deste plano):\n" + req.DossierJSON,
			// O ponto de cache fica no fim do bloco: turnos seguintes leem o prefixo a 0,1x.
			"cache_control": map[string]any{"type": "ephemeral"},
		},
	}

	mensagens := make([]map[string]any, 0, len(req.Turns)+3)
	for _, t := range req.Turns {
		mensagens = append(mensagens, map[string]any{"role": t.Role, "content": t.Body})
	}
	// O rascunho e o estado vão DEPOIS do histórico e fora do bloco de sistema, porque mudam a
	// cada turno: no topo, invalidariam o prefixo cacheado toda vez.
	mensagens = append(mensagens, map[string]any{
		"role":    "user",
		"content": "RASCUNHO ATUAL (slides, na ordem):\n" + req.SlidesJSON,
	})
	if req.StateNote != "" {
		mensagens = append(mensagens, map[string]any{"role": "user", "content": req.StateNote})
	}
	mensagens = append(mensagens, map[string]any{"role": "user", "content": req.Instruction})

	modelo := req.Model
	if modelo == "" {
		modelo = s.model
	}
	payload := map[string]any{
		"model":      modelo,
		"max_tokens": 8000,
		"system":     sistema,
		"messages":   mensagens,
		"tools": []map[string]any{{
			"name":         "editar_devolutiva",
			"description":  "Responde ao médico e propõe operações sobre os slides do rascunho.",
			"input_schema": planEditToolSchema(),
		}},
		"tool_choice": map[string]any{"type": "tool", "name": "editar_devolutiva"},
		// Sem `temperature`: com tool forçado, os modelos atuais respondem 400.
	}

	inicio := time.Now()
	bruto, m, err := s.chamaFerramentaComMeta(payload)
	meta = m
	meta.LatencyMs = int(time.Since(inicio).Milliseconds())
	if err != nil {
		return nil, meta, err
	}

	var out PlanEditResult
	if err := json.Unmarshal([]byte(bruto), &out); err != nil {
		// Diferente do escriba de teleconsulta, que devolve nil em silêncio quando o parse falha:
		// aqui um turno mal formado não pode aplicar operação nenhuma, e o médico precisa saber
		// que a rodada falhou em vez de achar que a IA não teve o que sugerir.
		return nil, meta, fmt.Errorf("%w: resposta do modelo não pôde ser lida", ErrAIUpstream)
	}
	return &out, meta, nil
}

// chamaFerramentaComMeta é o `callClaudeToolUse` com os metadados que a revisão precisa guardar.
func (s *AIService) chamaFerramentaComMeta(payload map[string]any) (string, AICallMeta, error) {
	var meta AICallMeta
	corpo, err := json.Marshal(payload)
	if err != nil {
		return "", meta, err
	}
	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(corpo))
	if err != nil {
		return "", meta, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", s.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", meta, fmt.Errorf("%w: %v", ErrAIUpstream, err)
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		// Não loga o corpo: pode ecoar trecho do prompt clínico.
		return "", meta, fmt.Errorf("%w: status %d", ErrAIUpstream, resp.StatusCode)
	}

	var api struct {
		Content []struct {
			Type  string          `json:"type"`
			Input json.RawMessage `json:"input"`
		} `json:"content"`
		Model      string `json:"model"`
		StopReason string `json:"stop_reason"`
		Usage      struct {
			InputTokens      int `json:"input_tokens"`
			OutputTokens     int `json:"output_tokens"`
			CacheReadTokens  int `json:"cache_read_input_tokens"`
			CacheWriteTokens int `json:"cache_creation_input_tokens"`
		} `json:"usage"`
	}
	if err := json.Unmarshal(b, &api); err != nil {
		return "", meta, fmt.Errorf("%w: resposta ilegível", ErrAIUpstream)
	}
	meta = AICallMeta{
		Model: api.Model, StopReason: api.StopReason,
		InputTokens: api.Usage.InputTokens, OutputTokens: api.Usage.OutputTokens,
		CacheReadTokens: api.Usage.CacheReadTokens,
	}
	// Só metadados no log; nunca prompt nem resposta.
	fmt.Printf("💰 Devolutiva - modelo %s, entrada %d (cache %d), saída %d, stop %s\n",
		meta.Model, meta.InputTokens, meta.CacheReadTokens, meta.OutputTokens, meta.StopReason)

	if api.StopReason == "max_tokens" {
		return "", meta, fmt.Errorf("%w: resposta truncada", ErrAITruncated)
	}
	for _, c := range api.Content {
		if c.Type == "tool_use" {
			return string(c.Input), meta, nil
		}
	}
	return "", meta, fmt.Errorf("%w: sem tool_use na resposta", ErrAIUpstream)
}
