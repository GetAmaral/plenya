package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/plenya/api/internal/pdfdoc"
)

// Operações sobre o rascunho, e o que cada uma pode tocar.
//
// A entrada não é JSON Pointer livre: caminho arbitrário deixaria escrever em
// `rulers[0].segments[2].b`, que é faixa do catálogo do escore e não texto de ninguém. O que existe
// é um VOCABULÁRIO fechado, e cada caminho tem uma classe fixa.

// FieldClass — o que se pode fazer com um caminho.
type FieldClass int

const (
	// FieldUnknown — caminho fora do vocabulário. Recusado, nem vira sugestão.
	FieldUnknown FieldClass = iota
	// FieldAuthoredText — prosa que o médico escreve. Aplica direto SE não introduzir número novo;
	// a checagem é do diff, não do campo, porque `punch` e `title` carregam número o tempo todo.
	FieldAuthoredText
	// FieldNumeric — número, unidade ou dose. Sempre sugestão.
	FieldNumeric
	// FieldDossierOwned — vem do dossiê e não é autoral. Recusado.
	FieldDossierOwned
	// FieldNumericAuthored — o eixo da régua: sai do dossiê mas é o único número legitimamente
	// ajustado à mão, quando um valor extremo esmaga a escala. Sugestão, com invariante.
	FieldNumericAuthored
)

func (c FieldClass) String() string {
	switch c {
	case FieldAuthoredText:
		return "texto"
	case FieldNumeric:
		return "numérico"
	case FieldDossierOwned:
		return "do dossiê"
	case FieldNumericAuthored:
		return "eixo"
	}
	return "desconhecido"
}

var (
	ErrCaminhoDesconhecido = errors.New("caminho fora do vocabulário do slide")
	ErrCaminhoDoDossie     = errors.New("este campo vem do dossiê e não pode ser reescrito")
	ErrSlideNaoEncontrado  = errors.New("slide não encontrado")
	ErrEixoCortaDado       = errors.New("o eixo proposto deixa de fora dado do paciente")
)

// indice troca `[3]` por `[i]`, para o vocabulário não precisar listar todas as posições.
var indice = regexp.MustCompile(`\[\d+\]`)

// qualquerIndice tira o colchete inteiro, já normalizado ou não. Usar `indice` para isto foi um
// bug: depois da normalização o caminho já tem `[i]`, que `\[\d+\]` não casa, e a raiz do bloco
// saía como "rulers[i]" — o que recusava TODA edição de régua.
var qualquerIndice = regexp.MustCompile(`\[(?:\d+|i)\]`)

// vocabulario mapeia o caminho normalizado para a classe.
//
// Só o que está aqui pode ser tocado. O que falta é recusa, não permissão implícita: uma feature
// que escreve num documento lido por paciente erra melhor negando.
var vocabulario = map[string]FieldClass{
	// Envelope
	"title":   FieldAuthoredText,
	"eyebrow": FieldAuthoredText,
	"lede":    FieldAuthoredText,
	"kicker":  FieldAuthoredText,
	"source":  FieldAuthoredText,
	"punch":   FieldAuthoredText,
	"variant": FieldAuthoredText,
	"legend":  FieldAuthoredText,

	// Cartões
	"cards[i].kicker": FieldAuthoredText,
	"cards[i].body":   FieldAuthoredText,
	"cards[i].dim":    FieldAuthoredText,
	"cards[i].focus":  FieldAuthoredText,

	// Régua: três campos autorais, o resto é do dossiê.
	"rulers[i].display":  FieldAuthoredText,
	"rulers[i].sub":      FieldAuthoredText,
	"rulers[i].note":     FieldAuthoredText,
	"rulers[i].unit":     FieldNumeric,
	"rulers[i].axis":     FieldNumericAuthored,
	"rulers[i].code":     FieldDossierOwned,
	"rulers[i].segments": FieldDossierOwned,
	"rulers[i].history":  FieldDossierOwned,

	// Resumo
	"summary.stepsTitle":              FieldAuthoredText,
	"summary.steps[i]":                FieldAuthoredText,
	"summary.legend":                  FieldAuthoredText,
	"summary.cards[i].title":          FieldAuthoredText,
	"summary.cards[i].tone":           FieldAuthoredText,
	"summary.cards[i].lines[i].name":  FieldAuthoredText,
	"summary.cards[i].lines[i].sub":   FieldAuthoredText,
	"summary.cards[i].lines[i].value": FieldNumeric,
	"summary.cards[i].lines[i].unit":  FieldNumeric,
	"summary.cards[i].lines[i].plot":  FieldNumeric,
	"summary.cards[i].lines[i].code":  FieldNumeric,
	"summary.cards[i].lines[i].ruler": FieldDossierOwned,

	// Sequência
	"steps[i].when":   FieldAuthoredText,
	"steps[i].what":   FieldAuthoredText,
	"steps[i].detail": FieldAuthoredText,

	// Para levar
	"takeaway.note":                    FieldAuthoredText,
	"takeaway.highlight.name":          FieldAuthoredText,
	"takeaway.highlight.obs":           FieldAuthoredText,
	"takeaway.highlight.when":          FieldAuthoredText,
	"takeaway.highlight.dose":          FieldNumeric,
	"takeaway.highlight.unit":          FieldNumeric,
	"takeaway.groups[i].title":         FieldAuthoredText,
	"takeaway.groups[i].items[i].name": FieldAuthoredText,
	"takeaway.groups[i].items[i].sub":  FieldAuthoredText,
	"takeaway.groups[i].items[i].dose": FieldNumeric,

	// Tabela. A célula é classificada em ClassifyPath, porque depende do estilo da COLUNA:
	// coluna de dose é numérica, coluna de prosa não é.
	"table.dense":            FieldAuthoredText,
	"table.columns[i].label": FieldAuthoredText,
	"table.columns[i].style": FieldAuthoredText,
	"table.columns[i].width": FieldAuthoredText,
	"table.rows[i].muted":    FieldAuthoredText,
}

// blocosPorKind diz qual bloco cada tipo de slide aceita. Escrever `table.*` num slide `cover` é
// caminho válido no vocabulário e inválido naquele slide.
var blocosPorKind = map[pdfdoc.DeckSlideKind]map[string]bool{
	pdfdoc.DeckCover:     {},
	pdfdoc.DeckClosing:   {"cards": true},
	pdfdoc.DeckSummary:   {"summary": true},
	pdfdoc.DeckRulers:    {"rulers": true},
	pdfdoc.DeckTwoCards:  {"cards": true, "table": true},
	pdfdoc.DeckPlanStep:  {"cards": true, "table": true, "takeaway": true},
	pdfdoc.DeckSequence:  {"steps": true},
	pdfdoc.DeckTakeaway:  {"takeaway": true},
	pdfdoc.DeckTableKind: {"table": true},
}

// ClassifyPath diz o que se pode fazer com um caminho naquele tipo de slide.
//
// `slide` entra porque a célula de tabela depende do estilo da coluna: numa coluna `dose` a célula
// é número, e numa coluna de prosa não é. Sem o slide, essa distinção não existe.
func ClassifyPath(slide *pdfdoc.DeckSlide, path string) (FieldClass, error) {
	p := strings.TrimSpace(path)
	if p == "" {
		return FieldUnknown, ErrCaminhoDesconhecido
	}
	normal := indice.ReplaceAllString(p, "[i]")

	// Célula de tabela: a classe sai do estilo da coluna correspondente.
	if m := regexp.MustCompile(`^table\.rows\[(\d+)\]\.cells\[(\d+)\]$`).FindStringSubmatch(p); m != nil {
		if slide == nil || slide.Table == nil {
			return FieldUnknown, ErrCaminhoDesconhecido
		}
		var col int
		fmt.Sscanf(m[2], "%d", &col)
		if col < len(slide.Table.Columns) && slide.Table.Columns[col].Style == pdfdoc.DeckColDose {
			return FieldNumeric, nil
		}
		return FieldAuthoredText, nil
	}

	classe, ok := vocabulario[normal]
	if !ok {
		return FieldUnknown, ErrCaminhoDesconhecido
	}
	if slide != nil {
		raiz := qualquerIndice.ReplaceAllString(strings.SplitN(normal, ".", 2)[0], "")
		if permitidos, temKind := blocosPorKind[slide.Kind]; temKind {
			if _, ehEnvelope := envelope[raiz]; !ehEnvelope && !permitidos[raiz] {
				return FieldUnknown, fmt.Errorf("%w: %q não vale num slide %q", ErrCaminhoDesconhecido, raiz, slide.Kind)
			}
		}
	}
	if classe == FieldDossierOwned {
		return classe, ErrCaminhoDoDossie
	}
	return classe, nil
}

var envelope = map[string]struct{}{
	"title": {}, "eyebrow": {}, "lede": {}, "kicker": {},
	"source": {}, "punch": {}, "variant": {}, "legend": {},
}

// ---- Aplicação ----

// PlanOpKind — as quatro operações possíveis sobre a lista de slides.
type PlanOpKind string

const (
	OpAdd     PlanOpKind = "add"
	OpEdit    PlanOpKind = "edit"
	OpRemove  PlanOpKind = "remove"
	OpReorder PlanOpKind = "reorder"
)

// PlanOp — uma operação. O alvo é sempre o ID do slide, nunca o índice: índice muda quando alguém
// reordena, e uma operação criada sobre "o slide 6" aplicada depois escreveria no lugar errado.
type PlanOp struct {
	Op           PlanOpKind        `json:"op"`
	SlideID      string            `json:"slideId,omitempty"`
	AfterSlideID string            `json:"afterSlideId,omitempty"`
	Path         string            `json:"path,omitempty"`
	Value        any               `json:"value,omitempty"`
	Slide        *pdfdoc.DeckSlide `json:"slide,omitempty"`
	Order        []string          `json:"order,omitempty"`
}

// ApplyOps aplica as operações em ordem sobre uma CÓPIA dos slides.
//
// Não muta a entrada: quem chama precisa do estado anterior para gravar a revisão e montar o diff.
func ApplyOps(slides []pdfdoc.DeckSlide, ops []PlanOp) ([]pdfdoc.DeckSlide, error) {
	out := make([]pdfdoc.DeckSlide, len(slides))
	copy(out, slides)
	for i, op := range ops {
		novo, err := aplicaUma(out, op)
		if err != nil {
			return nil, fmt.Errorf("operação %d (%s): %w", i+1, op.Op, err)
		}
		out = novo
	}
	return out, nil
}

func aplicaUma(slides []pdfdoc.DeckSlide, op PlanOp) ([]pdfdoc.DeckSlide, error) {
	switch op.Op {
	case OpRemove:
		i := indiceDoSlide(slides, op.SlideID)
		if i < 0 {
			return nil, ErrSlideNaoEncontrado
		}
		return append(append([]pdfdoc.DeckSlide{}, slides[:i]...), slides[i+1:]...), nil

	case OpAdd:
		if op.Slide == nil {
			return nil, errors.New("add sem slide")
		}
		novo := *op.Slide
		if novo.ID == "" {
			com, _ := EnsureSlideIDs([]pdfdoc.DeckSlide{novo})
			novo = com[0]
		}
		pos := 0
		if op.AfterSlideID != "" {
			i := indiceDoSlide(slides, op.AfterSlideID)
			if i < 0 {
				return nil, ErrSlideNaoEncontrado
			}
			pos = i + 1
		}
		out := append([]pdfdoc.DeckSlide{}, slides[:pos]...)
		out = append(out, novo)
		return append(out, slides[pos:]...), nil

	case OpReorder:
		if len(op.Order) != len(slides) {
			return nil, fmt.Errorf("reorder com %d ids para %d slides: a ordem tem que ser completa",
				len(op.Order), len(slides))
		}
		porID := map[string]pdfdoc.DeckSlide{}
		for _, s := range slides {
			porID[s.ID] = s
		}
		out := make([]pdfdoc.DeckSlide, 0, len(slides))
		for _, id := range op.Order {
			s, ok := porID[id]
			if !ok {
				return nil, fmt.Errorf("%w: %q", ErrSlideNaoEncontrado, id)
			}
			delete(porID, id)
			out = append(out, s)
		}
		if len(porID) > 0 {
			return nil, errors.New("reorder deixou slide de fora")
		}
		return out, nil

	case OpEdit:
		i := indiceDoSlide(slides, op.SlideID)
		if i < 0 {
			return nil, ErrSlideNaoEncontrado
		}
		classe, err := ClassifyPath(&slides[i], op.Path)
		if err != nil {
			return nil, err
		}
		if classe == FieldDossierOwned {
			return nil, ErrCaminhoDoDossie
		}
		out := make([]pdfdoc.DeckSlide, len(slides))
		copy(out, slides)
		if err := setPath(&out[i], op.Path, op.Value); err != nil {
			return nil, err
		}
		return out, nil
	}
	return nil, fmt.Errorf("operação desconhecida: %q", op.Op)
}

func indiceDoSlide(slides []pdfdoc.DeckSlide, id string) int {
	if id == "" {
		return -1
	}
	for i := range slides {
		if slides[i].ID == id {
			return i
		}
	}
	return -1
}

// ValidateAxis confere o invariante do eixo: ele pode ser apertado à mão, mas não pode deixar de
// fora nenhuma faixa nem nenhum ponto medido do paciente.
//
// Eixo que corta um ponto esconde o dado sem erro nenhum, que é a mesma classe de falha do
// `overflow:hidden` engolindo conteúdo do slide.
func ValidateAxis(novo []float64, r pdfdoc.DeckRulerBlock) error {
	if len(novo) != 2 || novo[0] >= novo[1] {
		return fmt.Errorf("eixo inválido: %v", novo)
	}
	fora := func(v float64) bool { return v < novo[0] || v > novo[1] }
	for _, sg := range r.Segments {
		if fora(sg.A) || fora(sg.B) {
			return fmt.Errorf("%w: a faixa de nível %d vai de %g a %g", ErrEixoCortaDado, sg.Level, sg.A, sg.B)
		}
	}
	for _, p := range r.History {
		if fora(p.Value) {
			return fmt.Errorf("%w: o ponto de %s vale %g", ErrEixoCortaDado, p.Date, p.Value)
		}
	}
	return nil
}

// ---- Escrita por caminho ----

// setPath grava `valor` no caminho dentro do slide.
//
// Vai e volta por JSON de propósito, em vez de reflexão: o contrato do slide É o JSON (é o que está
// no JSONB e o que o front manda), então percorrer a forma serializada é percorrer o contrato. Uma
// travessia por reflexão precisaria reimplementar as regras de tag, `omitempty` e ponteiro, e
// divergiria do que de fato é gravado no dia em que alguém mudar uma tag.
func setPath(slide *pdfdoc.DeckSlide, path string, valor any) error {
	bruto, err := json.Marshal(slide)
	if err != nil {
		return err
	}
	var raiz map[string]any
	if err := json.Unmarshal(bruto, &raiz); err != nil {
		return err
	}
	if err := gravaEm(raiz, quebraCaminho(path), valor); err != nil {
		return err
	}
	novo, err := json.Marshal(raiz)
	if err != nil {
		return err
	}
	var out pdfdoc.DeckSlide
	if err := json.Unmarshal(novo, &out); err != nil {
		return err
	}
	*slide = out
	return nil
}

// passo — um degrau do caminho: um nome de campo ou um índice de lista.
type passo struct {
	campo   string
	indice  int
	ehLista bool
}

var passoRe = regexp.MustCompile(`^([A-Za-z0-9_]+)((?:\[\d+\])*)$`)

func quebraCaminho(path string) []passo {
	var out []passo
	for _, parte := range strings.Split(path, ".") {
		m := passoRe.FindStringSubmatch(parte)
		if m == nil {
			return nil
		}
		out = append(out, passo{campo: m[1]})
		for _, idx := range regexp.MustCompile(`\[(\d+)\]`).FindAllStringSubmatch(m[2], -1) {
			var i int
			fmt.Sscanf(idx[1], "%d", &i)
			out = append(out, passo{indice: i, ehLista: true})
		}
	}
	return out
}

func gravaEm(atual any, passos []passo, valor any) error {
	if len(passos) == 0 {
		return errors.New("caminho vazio")
	}
	p := passos[0]
	ultimo := len(passos) == 1

	if p.ehLista {
		lista, ok := atual.([]any)
		if !ok || p.indice < 0 || p.indice >= len(lista) {
			return fmt.Errorf("índice %d fora da lista", p.indice)
		}
		if ultimo {
			lista[p.indice] = valor
			return nil
		}
		return gravaEm(lista[p.indice], passos[1:], valor)
	}

	obj, ok := atual.(map[string]any)
	if !ok {
		return fmt.Errorf("campo %q não está num objeto", p.campo)
	}
	if ultimo {
		obj[p.campo] = valor
		return nil
	}
	prox, existe := obj[p.campo]
	if !existe || prox == nil {
		// Bloco ausente não é erro de caminho: é slide que ainda não tem aquele bloco. Criar um
		// objeto vazio deixa o próximo passo falhar com mensagem específica, em vez de estourar.
		if passos[1].ehLista {
			return fmt.Errorf("a lista %q ainda não existe neste slide", p.campo)
		}
		novo := map[string]any{}
		obj[p.campo] = novo
		prox = novo
	}
	return gravaEm(prox, passos[1:], valor)
}

// qualquerIndiceNum captura o primeiro índice numérico do caminho, para achar de qual régua se
// está falando.
var qualquerIndiceNum = regexp.MustCompile(`\[(\d+)\]`)

// getPath lê o valor que está no caminho. É a travessia de setPath ao contrário, e reusa a mesma
// serialização de propósito: duas implementações do mesmo caminho divergem no primeiro campo novo.
func getPath(slide *pdfdoc.DeckSlide, path string) (any, error) {
	bruto, err := json.Marshal(slide)
	if err != nil {
		return nil, err
	}
	var raiz any
	if err := json.Unmarshal(bruto, &raiz); err != nil {
		return nil, err
	}
	atual := raiz
	for _, p := range quebraCaminho(path) {
		if p.ehLista {
			lista, ok := atual.([]any)
			if !ok || p.indice < 0 || p.indice >= len(lista) {
				return nil, fmt.Errorf("índice %d fora da lista", p.indice)
			}
			atual = lista[p.indice]
			continue
		}
		obj, ok := atual.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("campo %q não está num objeto", p.campo)
		}
		atual = obj[p.campo]
	}
	return atual, nil
}
