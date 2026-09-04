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
			"label": map[string]any{
				"type": "string",
				"description": `Um rótulo TEMÁTICO para esta seção, se o material permitir dizer do que se ` +
					`trata ("O que o cigarro está cobrando" em vez de "O que está se movendo"). Vazio = fica o padrão.`,
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

// PlanArcResult — o arco, hoje calculado em código. O tipo sobrevive porque o orquestrador
// carrega as seções nele.
type PlanArcResult struct {
	Sections []PlanArcSection `json:"sections"`
}

// PlanSectionRequest / PlanSectionResult — a segunda passada, uma seção por chamada.
type PlanSectionRequest struct {
	DossierJSON string
	ArcoJSON    string
	Secao       PlanArcSection
	Indice      int
	Instruction string
	Model       string
}

type PlanSectionResult struct {
	Slides []pdfdoc.DeckSlide `json:"slides"`
	// Label — o rótulo temático que o modelo propõe para esta seção. O código escolhe o genérico
	// ("O que está se movendo"); nos decks aprovados os melhores slides têm rótulo temático ("O que
	// o cigarro está cobrando"), e isso é leitura clínica do conjunto, não conta. Custa dez tokens.
	Label string `json:"label,omitempty"`
}

// GeneratePlanSection escreve os slides de UMA seção.
func (s *AIService) GeneratePlanSection(req PlanSectionRequest) (*PlanSectionResult, AICallMeta, error) {
	var meta AICallMeta
	if !s.IsConfigured() {
		return nil, meta, ErrAINotConfigured
	}
	// O pedido é MÍNIMO de propósito. O arco inteiro vivia aqui e era reenviado nas oito chamadas,
	// ~1,5 mil tokens de entrada não cacheada cada; agora ele está no bloco de sistema, dentro do
	// prefixo que o cache cobre.
	pedido := fmt.Sprintf("Escreva a seção %d ([%s] %s), e só ela: exatamente %d slide(s).",
		req.Indice+1, req.Secao.Key, req.Secao.Label, req.Secao.Slides)
	if len(req.Secao.Exames) > 0 {
		pedido += fmt.Sprintf(" Exames desta seção: %v.", req.Secao.Exames)
	}
	if len(req.Secao.Condutas) > 0 {
		c, _ := json.Marshal(req.Secao.Condutas)
		pedido += " Condutas desta seção: " + string(c)
	}
	if t := strings.TrimSpace(req.Instruction); t != "" {
		pedido += "\n\nO médico pediu: " + t
	}
	payload := map[string]any{
		"model":      modeloOu(req.Model, s.model),
		"max_tokens": 16000,
		// `effort` baixo: escrever uma seção com o arco decidido, os exames escolhidos e a
		// gramática dada é execução, não raciocínio aberto. No alto, o pensamento (que é cobrado
		// como saída) dominava a conta.
		"output_config": map[string]any{"effort": "low"},
		"system": []map[string]any{
			{"type": "text", "text": planGenerateSystemPrompt},
			{"type": "text",
				// O ARCO entra aqui, e não no pedido: ele é idêntico nas oito chamadas, então
				// dentro do prefixo ele é escrito uma vez e lido sete, a 0,1x em vez de 1x.
				"text": "DOSSIÊ (prontuário compilado e congelado deste plano):\n" + req.DossierJSON +
					"\n\n" + req.ArcoJSON,
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
	SlidesJSON string
	Excessos   string
	Model      string
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
		// `effort` baixo e SEM dossiê: o reparo encurta prosa e tem ordem explícita de não cortar
		// número nenhum, então o dossiê seria ~5 mil tokens de prefixo escrito (a 1,25x) para uma
		// tarefa que não consulta dado. Foi a segunda maior linha de custo depois da escrita de
		// cache das seções.
		"output_config": map[string]any{"effort": "low"},
		"system": []map[string]any{
			{"type": "text", "text": planRepairSystemPrompt},
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
