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
	bruto, m, err := s.chamaFerramentaComMeta(payload, false)
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
// chamaFerramentaComMeta faz a chamada. `longa` escolhe o cliente: a GERAÇÃO pede 32000 tokens de
// saída, o mesmo teto da extração de laudo, e precisa dos 10 minutos do `labClient` — com os 3
// minutos do `httpClient` um deck longo tem a conexão cortada DEPOIS de a chamada já ter sido
// cobrada. O turno de conversa continua nos 3 minutos: ele devolve um punhado de operações, e
// esperar dez por um turno travado seria pior que falhar.
func (s *AIService) chamaFerramentaComMeta(payload map[string]any, longa bool) (string, AICallMeta, error) {
	var meta AICallMeta
	cliente := s.httpClient
	if longa && s.labClient != nil {
		cliente = s.labClient
	}
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

	resp, err := cliente.Do(req)
	if err != nil {
		return "", meta, fmt.Errorf("%w: %v", ErrAIUpstream, err)
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		// O corpo de erro da Anthropic NÃO ecoa o prompt: traz `{"error":{"type","message"}}`
		// descrevendo o que está errado na REQUISIÇÃO (schema inválido, parâmetro recusado). Sem
		// ele, um 400 é indepurável — foi exatamente o que aconteceu aqui, e a mensagem dizia só
		// "status 400". Extrai só type e message, nunca o corpo cru.
		var apiErr struct {
			Error struct {
				Type    string `json:"type"`
				Message string `json:"message"`
			} `json:"error"`
		}
		if json.Unmarshal(b, &apiErr) == nil && apiErr.Error.Message != "" {
			return "", meta, fmt.Errorf("%w: status %d (%s: %s)", ErrAIUpstream, resp.StatusCode,
				apiErr.Error.Type, apiErr.Error.Message)
		}
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

// ---------------------------------------------------------------------------
// Geração do rascunho inteiro, a partir do dossiê.
//
// É o passo que faltava: a edição turno a turno (EditPatientPlan) mexe num deck que alguém já
// escreveu, e até aqui ninguém escrevia. O "novo plano" nascia com dois slides vazios.
//
// Vale a mesma regra de ouro da edição, com força maior: aqui o modelo escreve o documento inteiro
// de uma vez, então TODO número precisa vir declarado com origem no dossiê, e o servidor confere
// slide a slide antes de gravar.

// planGenerateSystemPrompt — as regras da edição mais o arco e a gramática.
//
// O arco não é invenção: é o que os dois decks reais convergiram (Ana, 21 slides; José Ricardo,
// 20), e está escrito na skill `/plano`. Os tetos (4 réguas, 8 linhas de tabela, 3 colunas) são
// medidos, não estimados: com 8 réguas o slide estoura e existe teste provando.
const planGenerateSystemPrompt = planSystemPrompt + `

VOCÊ ESTÁ ESCREVENDO O RASCUNHO INTEIRO, do zero, a partir do dossiê.

O ARCO (use como espinha, corte o que não se aplica a este paciente):
1   capa                  nome do paciente não; só uma frase de abertura
2   resumo em uma página  o que está forte | o que está se movendo | o que vamos fazer
3-4 o que está bem        1 a 4 réguas por slide, tiradas do topo de "strong"
5-7 o que está se movendo UM achado por slide, do topo de "moving"
8+  o plano               uma conduta por slide, vindas de carePlan
n-2 a sequência           os próximos meses, em ordem
n-1 para levar            o que começa agora, com dose, vindo de prescriptions
n   em uma página         o fecho

COMO ESCOLHER O QUE ENTRA:
- "o que está bem" sai de strong, que já vem ordenado por PESO do item. Em nível 4-5 ninguém perde
  ponto, então peso é o único sinal.
- Boa parte da anamnese é checklist de ausência ("Adrenalectomia: não"). Isso NÃO é conquista e
  NÃO vira slide.
- "o que está se movendo" sai do topo de moving. Três é o número que funcionou nos dois decks.
  Um achado com trend "worsening" merece prioridade mesmo em nível bom: a direção é o sinal.
- Um assunto por slide. Se dois achados precisam ser explicados juntos, são dois slides.
- Se o dossiê não tem conduta registrada, NÃO invente conduta. Escreva menos slides.

VOZ (o paciente é adulto e está lendo sobre o próprio corpo):
- Título de slide é uma AFIRMAÇÃO, não um rótulo: "A ferritina dobrou em dois anos", não "Ferritina".
- "display" da régua é o nome que o paciente reconhece ("Ferritina"), nunca o do catálogo
  ("Ferritina - Homens"). "sub" explica o que o exame mede em até cinco palavras.
- "punch" fecha o slide com a consequência, e é o único lugar onde <em> entra.
- Todo número vem com unidade e com a data em que foi medido.

PROIBIDO no texto que o paciente lê: travessão; a construção "Não é X. É Y."; fecho em slogan;
ícone decorativo em lista; preço; marca comercial (suplemento, aparelho, laboratório, varejista);
a expressão "medicina preditiva"; qualquer coisa que identifique outra pessoa.

REGRA DE LEI DA RÉGUA: nenhuma régua entra num slide sem um rótulo avaliativo visível no MESMO
slide. Pode estar no título, no punch ou no note da régua, mas tem que estar em algum deles.
Barra colorida sozinha comunica pior que barra com rótulo.

TETOS MEDIDOS, não sugestões: no máximo 4 réguas por slide (com 8 o slide estoura), 8 linhas de
tabela, 3 colunas, 4 linhas por cartão de resumo, 3 grupos no "para levar".

O QUE OS DOIS DECKS APROVADOS FAZEM, medido slide a slide:

- A TABELA é o bloco mais usado, não a régua: 9 dos 21 slides de um, 8 dos 20 do outro. A régua é o
  átomo visual; a tabela é COMO O PLANO SE EXPLICA. Toda seção de conduta é tabela. Cabeçalhos que
  se repetem: "O quê | Quanto | Por quê", "O quê | Dose | Por quê", "Quando | O que acontece".
- A SEQUÊNCIA é uma TABELA de 2 colunas, não o bloco "sequence": primeira coluna com style "dose"
  e valores relativos ("Agora", "Em 4 semanas", "Em 12 semanas"), nunca data absoluta. O bloco
  "sequence" não foi usado em nenhum dos dois decks.
- RÉGUA por slide: 2, 3 ou 4. NUNCA uma só, nunca cinco. Média 3,1 nos dois.
- PUNCH em 85% dos slides, e ausente só em capa, para-levar e fecho. Entre 55 e 110 caracteres,
  com EXATAMENTE UM <em>, e em 9 de cada 10 a frase termina dentro dele. Duas frases na maioria:
  uma constatação plana, depois a virada que carrega a decisão.
- TÍTULO entre 16 e 53 caracteres, uma linha. Só o fecho é longo: lá são três frases, ~280
  caracteres, e é o resumo que o paciente leva.
- "sub" da régua: 6 a 27 caracteres, minúsculo, sem ponto final, dizendo o que o exame MEDE
  ("estoque de ferro", "o rim filtrando"), nunca o resultado. Deixe vazio quando o nome já explica.
- "note" da régua aparece em ~30% delas, e é onde mora a mini-série temporal
  ("239 em 2024, 432 em 2025, 500 agora") ou o rótulo avaliativo.

STRINGS QUE OS DOIS DECKS COMPARTILHAM, use as mesmas: título do resumo "Onde você está, em uma
página"; cartões do resumo "O que está forte" e "O que está se movendo"; título dos passos "O que
vamos fazer"; eyebrow e título da sequência "A sequência" / "Os próximos três meses, em ordem";
"Para levar" / "O que você começa a tomar agora"; eyebrow do fecho "Em uma página".`

// planSlideItemSchema — a forma de UM slide. Compartilhada pelas duas passadas da geração: o
// contrato do slide é o mesmo escrevendo o deck inteiro ou escrevendo uma seção.
func planSlideItemSchema() map[string]any {
	return map[string]any{
		"type":     "object",
		"required": []string{"kind"},
		"properties": map[string]any{
			// `kind` estava livre e o modelo devolveu os onze slides com ele VAZIO, o
			// que deixa o render sem saber o que desenhar. Enum resolve na origem.
			"kind": map[string]any{
				"enum": []string{"cover", "summary", "rulers", "rulers-cards", "two-cards",
					"plan-step", "sequence", "takeaway", "closing", "table"},
			},
			// O render conhece "", "dark" e "deep". O enum antigo oferecia "light",
			// que não é valor de render nenhum e caía em creme por acidente. Nos dois
			// decks aprovados só capa e fecho são "deep", e "dark" nunca foi usado.
			//
			// O vazio precisa estar no enum: com um valor legal só, o modelo tende a
			// preencher o campo em todo slide e o deck sai inteiro escuro, que é o
			// oposto da regra ("um tom só nas páginas de conteúdo").
			"variant": map[string]any{
				"enum":        []string{"", "deep"},
				"description": `Vazio em TODO slide de conteúdo. "deep" só na capa e no fecho.`,
			},
			"eyebrow": map[string]any{"type": "string"},
			"title":   map[string]any{"type": "string"},
			"lede":    map[string]any{"type": "string"},
			"punch":   map[string]any{"type": "string"},
			// Os blocos abaixo precisam estar no schema com a forma exata. Sem eles o
			// modelo devolvia o slide com cabeçalho e NADA no miolo: `summary`,
			// `two-cards`, `sequence` e `closing` saíam ocos, bonitos por fora e sem
			// conteúdo nenhum. Só `rulers` vinha preenchido, que era o único descrito.
			"summary": map[string]any{
				"type":        "object",
				"description": "Só em kind=summary. Dois cartões (o que está forte, o que está se movendo) e os passos embaixo.",
				"properties": map[string]any{
					"cards": map[string]any{
						"type": "array", "maxItems": 2,
						"items": map[string]any{
							"type":     "object",
							"required": []string{"title"},
							"properties": map[string]any{
								"title": map[string]any{"type": "string"},
								"tone":  map[string]any{"enum": []string{"bom", "ruim"}},
								"lines": map[string]any{
									"type": "array", "maxItems": 4,
									"items": map[string]any{
										"type":     "object",
										"required": []string{"name", "value", "code"},
										"properties": map[string]any{
											"name":  map[string]any{"type": "string"},
											"sub":   map[string]any{"type": "string"},
											"code":  map[string]any{"type": "string", "description": "O code do exame de onde o valor saiu. Obrigatório: é o que torna o número auditável."},
											"value": map[string]any{"type": "string"},
											"unit":  map[string]any{"type": "string"},
										},
									},
								},
							},
						},
					},
					"stepsTitle": map[string]any{"type": "string"},
					"steps":      map[string]any{"type": "array", "maxItems": 4, "items": map[string]any{"type": "string"}},
				},
			},
			"cards": map[string]any{
				"type":     "array",
				"maxItems": 4,
				"description": "Em two-cards e rulers-cards, 2 ou 4 (a grade 2x2 da decisão). " +
					"Em plan-step, 2 ou 3. NÃO use em cover nem em closing: o render descarta.",
				"items": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"kicker":  map[string]any{"type": "string", "description": "Rótulo em caixa alta, geralmente a PERGUNTA que o cartão responde."},
						"verdict": map[string]any{"type": "string", "description": "A resposta, em uma ou duas palavras: \"Não precisa\", \"Sim\". É o que faz o slide decidir."},
						"tone":    map[string]any{"enum": []string{"ok", "flag"}, "description": "ok = tranquiliza; flag = regra de segurança. Vazio = neutro."},
						"body":    map[string]any{"type": "string"},
						"dim":     map[string]any{"type": "boolean", "description": "Apaga o caminho descartado."},
						"focus":   map[string]any{"type": "boolean", "description": "Destaca o caminho recomendado."},
					},
				},
			},
			"steps": map[string]any{
				"type":        "array",
				"maxItems":    6,
				"description": "Só em kind=sequence: os próximos meses, em ordem.",
				"items": map[string]any{
					"type":     "object",
					"required": []string{"when", "what"},
					"properties": map[string]any{
						"when":   map[string]any{"type": "string", "description": `Ex.: "nas próximas 4 semanas".`},
						"what":   map[string]any{"type": "string"},
						"detail": map[string]any{"type": "string"},
					},
				},
			},
			"takeaway": map[string]any{
				"type":        "object",
				"description": "Em kind=takeaway, e também em plan-step quando a conduta tem doses. Dose só sai de prescricoes do dossiê; sem prescrição, não escreva dose.",
				"properties": map[string]any{
					"highlight": map[string]any{
						"type": "object", "required": []string{"name"},
						"properties": map[string]any{
							"name": map[string]any{"type": "string"},
							"dose": map[string]any{"type": "string"},
							"unit": map[string]any{"type": "string"},
							"when": map[string]any{"type": "string"},
							"obs":  map[string]any{"type": "string"},
						},
					},
					"groups": map[string]any{
						"type": "array", "maxItems": 3,
						"items": map[string]any{
							"type": "object", "required": []string{"title"},
							"properties": map[string]any{
								"title": map[string]any{"type": "string"},
								"items": map[string]any{
									"type": "array",
									"items": map[string]any{
										"type": "object", "required": []string{"name"},
										"properties": map[string]any{
											"name": map[string]any{"type": "string"},
											"sub":  map[string]any{"type": "string"},
											"dose": map[string]any{"type": "string"},
										},
									},
								},
							},
						},
					},
					"note": map[string]any{"type": "string"},
				},
			},
			"table": map[string]any{
				"type":        "object",
				"description": "Em kind=table, e também em two-cards e plan-step, onde o render a desenha abaixo dos cartões. Máximo 3 colunas e 8 linhas.",
				"properties": map[string]any{
					// `dense` NÃO estava no schema e o render precisa dele: uma tabela de 6 linhas
					// com prosa estourou o slide em 580px na primeira geração de duas passadas. É
					// o botão que existe exatamente para isso.
					"dense": map[string]any{
						"type":        "boolean",
						"description": "Aperta o espaçamento. Use SEMPRE que a tabela passar de 4 linhas ou tiver coluna de prosa.",
					},
					"columns": map[string]any{
						"type": "array", "maxItems": 3,
						"items": map[string]any{
							"type": "object", "required": []string{"label"},
							"properties": map[string]any{
								"label": map[string]any{"type": "string"},
								"style": map[string]any{
									"enum": []string{"", "why", "dose", "tag"},
									"description": `"" texto normal; "why" a coluna que explica, menor e cinza; ` +
										`"dose" NÃO quebra linha, então nunca ponha prosa nela; "tag" vira selo.`,
								},
								"width": map[string]any{
									"type":        "string",
									"description": `Largura fixa, ex.: "390px". Use na primeira coluna quando ela for curta e a de prosa precisar do resto.`,
								},
							},
						},
					},
					"rows": map[string]any{
						"type": "array", "maxItems": 8,
						"items": map[string]any{
							"type": "object", "required": []string{"cells"},
							"properties": map[string]any{
								"cells": map[string]any{"type": "array", "items": map[string]any{"type": "string"}},
								"muted": map[string]any{"type": "boolean"},
							},
						},
					},
				},
			},
			"rulers": map[string]any{
				"type":     "array",
				"maxItems": 4,
				"description": "Em kind=rulers, 2 a 4. Em rulers-cards, exatamente 2. Só o exame e o texto autoral; a escala e o histórico o SERVIDOR copia " +
					"do dossiê; não escreva axis, segments nem history.",
				"items": map[string]any{
					"type":     "object",
					"required": []string{"code", "display"},
					"properties": map[string]any{
						"code":    map[string]any{"type": "string", "description": "O `code` da régua, copiado do dossiê."},
						"display": map[string]any{"type": "string", "description": "O nome que o paciente reconhece."},
						"sub":     map[string]any{"type": "string", "description": "O que o exame mede, em até cinco palavras."},
						"note":    map[string]any{"type": "string", "description": "Leitura clínica em uma linha. É onde o rótulo avaliativo mora."},
					},
				},
			},
		},
	}
}

// PlanModelNumeral — um número declarado pelo modelo, com a origem que ele alega.
type PlanModelNumeral struct {
	Numeral string `json:"numeral"`
	Source  string `json:"source"`
}
