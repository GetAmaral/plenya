package services

// A geração em DUAS passadas.
//
// A passada única escrevia o deck inteiro de uma vez, e isso custava três coisas que os decks
// aprovados têm:
//
//   - Os eyebrows de seção dos dois decks são numerados com o total decidido ANTES de escrever
//     ("O que está bem · 1 de 3", "O plano · 1 de 6"). Numa passada só, o modelo tem que adivinhar
//     o total enquanto escreve, e erra: o próprio deck aprovado do Ricardo repete
//     "· 1 de 3" em dois slides seguidos.
//   - O deck inteiro numa resposta já truncou uma vez, e crescer é o caminho natural daqui.
//   - O médico não via o arco antes de gastar a chamada inteira.
//
// Passada 1 decide o arco: quais seções existem, quantos slides cada uma tem, e com que material.
// Passada 2 escreve cada seção, com o arco inteiro no contexto e o detalhe só da sua parte.

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/plenya/api/internal/pdfdoc"
)

// PlanArcSection — uma seção do arco, decidida antes de qualquer slide ser escrito.
type PlanArcSection struct {
	// Key é do conjunto fechado abaixo. É por ela que o SERVIDOR aplica o que não é texto:
	// numeração do eyebrow, variante escura, legenda da rampa.
	Key string `json:"key"`
	// Label é o eyebrow SEM o contador ("O que está bem"). O "· 2 de 3" o servidor compõe.
	Label string `json:"label"`
	// Slides é quantos slides a seção terá. É o denominador do contador.
	Slides int `json:"slides"`
	// Exames e Condutas são o material que o modelo separou para esta seção.
	Exames   []string `json:"exames,omitempty"`
	Condutas []string `json:"condutas,omitempty"`
	// Porque é para o médico ler no arco, nunca vai para o deck.
	Porque string `json:"porque,omitempty"`
}

// As chaves de seção. Fechadas porque o servidor decide comportamento a partir delas.
const (
	SecCapa      = "capa"
	SecResumo    = "resumo"
	SecBem       = "bem"
	SecMovendo   = "movendo"
	SecDecisao   = "decisao"
	SecFaltando  = "faltando"
	SecPlano     = "plano"
	SecSequencia = "sequencia"
	SecLevar     = "levar"
	SecFecho     = "fecho"
)

// planArcSystemPrompt — as regras do ARCO. Curto de propósito: esta passada não escreve texto de
// paciente, decide estrutura.
const planArcSystemPrompt = `Você ajuda um médico a planejar a devolutiva de exames que o PACIENTE vai ler.

Sua tarefa AGORA é só decidir o ARCO: quais seções o deck terá, quantos slides cada uma, e com que
material. Você NÃO escreve slide nesta etapa.

O arco que os dois decks aprovados convergiram, em ordem:
  capa · resumo · o que está bem · o que está se movendo · a decisão em aberto ·
  o que está faltando · o plano · a sequência · para levar · o fecho

REGRAS DO ARCO:
1. "capa", "resumo", "sequencia", "levar" e "fecho" têm SEMPRE 1 slide.
2. "o que está bem" nunca tem menos de 2 slides, e vem SEMPRE antes de qualquer notícia ruim.
   Boa notícia primeiro é regra dos dois decks.
3. "o que está se movendo": um assunto por slide. Três foi o número que funcionou nos dois.
3b. CONTA DOS SLIDES DE RÉGUA: cada slide leva de 2 a 4 réguas, nunca uma só. Então o número de
   slides de uma seção de régua é a quantidade de exames dela dividida por 3, arredondada para
   cima: 3 exames = 1 slide, 5 exames = 2 slides, 8 exames = 3 slides. Pedir 3 slides com 4 exames
   deixa slide com uma régua sozinha, que é o defeito mais comum desta geração. Se a seção tem
   menos de 2 exames, ela não é seção de régua: junte com outra ou corte.
4. "a decisão em aberto" só existe quando há DOIS caminhos possíveis de verdade. Não invente dilema.
5. "o plano" tem um slide por conduta registrada em condutas. Sem conduta no dossiê, a seção NÃO
   EXISTE: não invente conduta, escreva um deck menor.
6. "o que está faltando" sai de pedidoDeExames, quando houver.
7. O label pode ser temático em vez de genérico ("O que o cigarro está cobrando" em vez de "O que
   está se movendo"), e nos decks aprovados os melhores slides são assim. Use quando o material
   permitir dizer do que se trata.
8. O deck inteiro fica entre 12 e 21 slides.
9. Achado com "stale" verdadeiro ou "daysAgo" grande NÃO lidera seção: é exame a refazer, não
   retrato de hoje. Foi assim que um HOMA-IR de dois anos atrás liderou um ranking e o deck teve
   que ser refeito.`

func planArcToolSchema() map[string]any {
	return map[string]any{
		"type":                 "object",
		"additionalProperties": false,
		"required":             []string{"reply", "sections"},
		"properties": map[string]any{
			"reply": map[string]any{
				"type":        "string",
				"description": "O arco que você escolheu e o que deixou de fora, para o MÉDICO ler. Nunca vai para o deck.",
			},
			"sections": map[string]any{
				"type":     "array",
				"minItems": 4,
				"maxItems": 12,
				"items": map[string]any{
					"type":                 "object",
					"additionalProperties": false,
					"required":             []string{"key", "label", "slides"},
					"properties": map[string]any{
						"key": map[string]any{
							"enum": []string{SecCapa, SecResumo, SecBem, SecMovendo, SecDecisao,
								SecFaltando, SecPlano, SecSequencia, SecLevar, SecFecho},
						},
						"label": map[string]any{
							"type": "string",
							"description": `O eyebrow SEM o contador: "O que está bem", "O plano". Pode ser temático ` +
								`("O que o cigarro está cobrando"). O "· 2 de 3" o servidor põe.`,
						},
						"slides": map[string]any{
							"type": "integer", "minimum": 1, "maximum": 8,
							"description": "Quantos slides esta seção terá. Decidir agora é o que faz o contador do eyebrow fechar.",
						},
						"exames":   map[string]any{"type": "array", "items": map[string]any{"type": "string"}, "description": "Os `code` dos exames desta seção, do dossiê."},
						"condutas": map[string]any{"type": "array", "items": map[string]any{"type": "string"}, "description": "As condutas desta seção, pelo texto da recomendação."},
						"porque":   map[string]any{"type": "string", "description": "Em uma frase, para o médico."},
					},
				},
			},
		},
	}
}

func planSectionToolSchema() map[string]any {
	return map[string]any{
		"type":                 "object",
		"additionalProperties": false,
		"required":             []string{"slides"},
		"properties": map[string]any{
			"slides": map[string]any{
				"type":        "array",
				"minItems":    1,
				"maxItems":    8,
				"description": "Os slides DESTA seção, na ordem. Exatamente a quantidade que o arco fixou.",
				"items":       planSlideItemSchema(),
			},
			"numerals": map[string]any{
				"type": "array",
				"items": map[string]any{
					"type": "object", "additionalProperties": false,
					"required":   []string{"numeral", "source"},
					"properties": map[string]any{"numeral": map[string]any{"type": "string"}, "source": map[string]any{"type": "string"}},
				},
			},
		},
	}
}

// PlanGenerateResult — o deck montado, já com o arco aplicado. O nome ficou da geração de uma
// passada, que não existe mais; a estrutura é o que `GenerateDraft` usa para seguir.
type PlanGenerateResult struct {
	Reply  string             `json:"reply"`
	Slides []pdfdoc.DeckSlide `json:"slides"`
}

// PlanArcRequest / PlanArcResult — a primeira passada.
type PlanArcRequest struct {
	DossierJSON string
	Instruction string
	Model       string
}

type PlanArcResult struct {
	Reply    string           `json:"reply"`
	Sections []PlanArcSection `json:"sections"`
}

// GeneratePlanArc decide o arco antes de escrever slide nenhum.
func (s *AIService) GeneratePlanArc(req PlanArcRequest) (*PlanArcResult, AICallMeta, error) {
	var meta AICallMeta
	if !s.IsConfigured() {
		return nil, meta, ErrAINotConfigured
	}
	instrucao := strings.TrimSpace(req.Instruction)
	if instrucao == "" {
		instrucao = "Decida o arco da devolutiva deste paciente."
	}
	payload := map[string]any{
		"model": modeloOu(req.Model, s.model),
		// O arco é curto: uma lista de seções. Não precisa do teto da escrita.
		"max_tokens": 4000,
		"system": []map[string]any{
			{"type": "text", "text": planArcSystemPrompt},
			{"type": "text", "text": "DOSSIÊ (prontuário compilado e congelado deste plano):\n" + req.DossierJSON,
				"cache_control": map[string]any{"type": "ephemeral"}},
		},
		"messages": []map[string]any{{"role": "user", "content": instrucao}},
		"tools": []map[string]any{{
			"name": "planejar_devolutiva", "description": "Decide o arco da devolutiva.",
			"input_schema": planArcToolSchema(),
		}},
		"tool_choice": map[string]any{"type": "tool", "name": "planejar_devolutiva"},
	}
	inicio := time.Now()
	bruto, m, err := s.chamaFerramentaComMeta(payload, false)
	meta = m
	meta.LatencyMs = int(time.Since(inicio).Milliseconds())
	if err != nil {
		return nil, meta, err
	}
	var out PlanArcResult
	if err := json.Unmarshal([]byte(bruto), &out); err != nil {
		return nil, meta, fmt.Errorf("%w: o arco não pôde ser lido", ErrAIUpstream)
	}
	if len(out.Sections) == 0 {
		return nil, meta, errors.New("o modelo não devolveu arco nenhum")
	}
	return &out, meta, nil
}

// PlanSectionRequest / PlanSectionResult — a segunda passada, uma seção por chamada.
type PlanSectionRequest struct {
	DossierJSON string
	ArcoJSON    string
	Secao       PlanArcSection
	Instruction string
	Model       string
}

type PlanSectionResult struct {
	Slides   []pdfdoc.DeckSlide `json:"slides"`
	Numerals []PlanModelNumeral `json:"numerals"`
}

// GeneratePlanSection escreve os slides de UMA seção.
func (s *AIService) GeneratePlanSection(req PlanSectionRequest) (*PlanSectionResult, AICallMeta, error) {
	var meta AICallMeta
	if !s.IsConfigured() {
		return nil, meta, ErrAINotConfigured
	}
	secao, _ := json.Marshal(req.Secao)
	pedido := "ARCO COMPLETO (para você saber o que já foi dito e o que vem depois):\n" + req.ArcoJSON +
		"\n\nESCREVA AGORA, e só, a seção:\n" + string(secao) +
		"\n\nDevolva EXATAMENTE " + fmt.Sprint(req.Secao.Slides) + " slide(s)."
	if t := strings.TrimSpace(req.Instruction); t != "" {
		pedido += "\n\nO médico pediu: " + t
	}
	payload := map[string]any{
		"model":      modeloOu(req.Model, s.model),
		"max_tokens": 16000,
		"system": []map[string]any{
			{"type": "text", "text": planGenerateSystemPrompt},
			{"type": "text", "text": "DOSSIÊ (prontuário compilado e congelado deste plano):\n" + req.DossierJSON,
				// Mesmo ponto de cache em todas as seções: o dossiê é byte-idêntico entre elas, e
				// é isso que faz a segunda passada custar uma fração da primeira.
				"cache_control": map[string]any{"type": "ephemeral"}},
		},
		"messages": []map[string]any{{"role": "user", "content": pedido}},
		"tools": []map[string]any{{
			"name": "escrever_secao", "description": "Escreve os slides de uma seção da devolutiva.",
			"input_schema": planSectionToolSchema(),
		}},
		"tool_choice": map[string]any{"type": "tool", "name": "escrever_secao"},
	}
	inicio := time.Now()
	bruto, m, err := s.chamaFerramentaComMeta(payload, true)
	meta = m
	meta.LatencyMs = int(time.Since(inicio).Milliseconds())
	if err != nil {
		return nil, meta, err
	}
	var out PlanSectionResult
	if err := json.Unmarshal([]byte(bruto), &out); err != nil {
		return nil, meta, fmt.Errorf("%w: a seção %q não pôde ser lida", ErrAIUpstream, req.Secao.Label)
	}
	return &out, meta, nil
}

// comID acrescenta `id` às propriedades do slide. Só o reparo precisa: na geração o id é do
// servidor, e deixá-lo no schema convidaria o modelo a inventar um.
func comID(item map[string]any) map[string]any {
	props, _ := item["properties"].(map[string]any)
	if props == nil {
		return item
	}
	copia := map[string]any{}
	for k, v := range item {
		copia[k] = v
	}
	novasProps := map[string]any{"id": map[string]any{
		"type": "string", "description": "COPIE o id do slide que veio. É por ele que a troca é feita.",
	}}
	for k, v := range props {
		novasProps[k] = v
	}
	copia["properties"] = novasProps
	copia["required"] = []string{"kind", "id"}
	return copia
}

func modeloOu(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

// ---------------------------------------------------------------------------
// A passada de reparo.
//
// O modelo não mede: ele escreve com as faixas de tamanho na cabeça e erra, porque a altura real
// depende de quantas linhas o texto quebra na fonte do deck. O servidor MEDE, no Chromium, e sabe
// exatamente quantos pixels sobraram para fora.
//
// Devolver isso ao modelo uma vez, com o número, é o que transforma "três slides estourando para o
// médico consertar" em "o deck cabe". Uma passada só, e limitada aos slides que estouraram: se a
// segunda tentativa ainda não couber, o aviso vai para a tela e quem corta é o médico.

// PlanRepairRequest — os slides que não couberam, com o excesso medido.
type PlanRepairRequest struct {
	DossierJSON string
	SlidesJSON  string
	Excessos    string
	Model       string
}

func planRepairToolSchema() map[string]any {
	return map[string]any{
		"type":                 "object",
		"additionalProperties": false,
		"required":             []string{"slides"},
		"properties": map[string]any{
			"slides": map[string]any{
				"type":        "array",
				"minItems":    1,
				"description": "Os MESMOS slides, na mesma ordem e com o mesmo id, corrigidos.",
				// O `id` precisa estar no SCHEMA, não só no texto: o casamento do reparo é por id,
				// e sem o campo declarado o modelo omite, `porID` fica vazio e uma chamada de 12k
				// tokens é descartada sem log e sem sinal nenhum para o médico.
				"items": comID(planSlideItemSchema()),
			},
		},
	}
}

const planRepairSystemPrompt = planSystemPrompt + `

OS SLIDES ABAIXO SAÍRAM FORA DO PADRÃO. Para cada um, o servidor diz o que está errado: pixels que
sobraram para fora da altura útil de 912px, punch fora da faixa de 55 a 110 caracteres, régua
sozinha num slide, travessão, ou coluna de dose com prosa.

Devolva OS MESMOS slides, na mesma ordem, com o MESMO id, corrigidos.

Sobre o PUNCH: duas frases, entre 55 e 110 caracteres no total, com EXATAMENTE um <em> fechando a
frase. Uma constatação plana, depois a virada que carrega a decisão. Punch longo demais é sempre a
mesma coisa: explicação que já está no corpo do slide.

Onde cortar, em ordem de preferência:
1. a coluna de prosa da tabela ("por quê"): é onde mora o excesso quase sempre;
2. linhas de tabela inteiras, tirando a menos decisiva, e ligando "dense" se ainda não estiver;
3. o "note" de uma régua que já tem rótulo avaliativo no título ou no punch;
4. o "sub" da régua, que é opcional;
5. o corpo do cartão.

NÃO corte: o punch, o título, nenhuma régua inteira, nenhum número. Cortar régua muda o que o
paciente vê do próprio exame; encurtar prosa não.

Cada 100px de excesso são cerca de duas linhas de texto na fonte do deck. Corte com folga: voltar
com o slide ainda estourando gasta outra rodada.`

// RepairOverflow devolve os slides encurtados.
func (s *AIService) RepairOverflow(req PlanRepairRequest) ([]pdfdoc.DeckSlide, AICallMeta, error) {
	var meta AICallMeta
	if !s.IsConfigured() {
		return nil, meta, ErrAINotConfigured
	}
	payload := map[string]any{
		"model":      modeloOu(req.Model, s.model),
		"max_tokens": 12000,
		"system": []map[string]any{
			{"type": "text", "text": planRepairSystemPrompt},
			{"type": "text", "text": "DOSSIÊ:\n" + req.DossierJSON,
				"cache_control": map[string]any{"type": "ephemeral"}},
		},
		"messages": []map[string]any{{"role": "user",
			"content": "EXCESSO MEDIDO POR SLIDE:\n" + req.Excessos + "\n\nSLIDES:\n" + req.SlidesJSON}},
		"tools": []map[string]any{{
			"name": "encurtar_slides", "description": "Encurta os slides que não couberam.",
			"input_schema": planRepairToolSchema(),
		}},
		"tool_choice": map[string]any{"type": "tool", "name": "encurtar_slides"},
	}
	inicio := time.Now()
	bruto, m, err := s.chamaFerramentaComMeta(payload, true)
	meta = m
	meta.LatencyMs = int(time.Since(inicio).Milliseconds())
	if err != nil {
		return nil, meta, err
	}
	var out struct {
		Slides []pdfdoc.DeckSlide `json:"slides"`
	}
	if err := json.Unmarshal([]byte(bruto), &out); err != nil {
		return nil, meta, fmt.Errorf("%w: o reparo não pôde ser lido", ErrAIUpstream)
	}
	return out.Slides, meta, nil
}
