package pdfdoc

import (
	"fmt"
	"strings"
)

// Os tipos de slide da devolutiva.
//
// Não são um catálogo aberto: saíram da leitura dos decks que já existem. Ana e José Ricardo
// convergiram para a MESMA gramática — capa, resumo em uma página, o que está bem, o que está se
// movendo, decisão em aberto, o plano, a sequência, para levar, fecho — com contagens de bloco
// quase idênticas (régua 230 vs 231, resumo 17 vs 17). Cada slide carrega UMA ideia.
type DeckSlideKind string

const (
	DeckCover    DeckSlideKind = "cover"     // capa: nome, data, uma frase de abertura
	DeckSummary  DeckSlideKind = "summary"   // "onde você está, em uma página"
	DeckRulers   DeckSlideKind = "rulers"    // 1 a 4 réguas + legenda da rampa
	DeckTwoCards DeckSlideKind = "two-cards" // dois caminhos possíveis, um deles em destaque
	DeckPlanStep DeckSlideKind = "plan-step" // uma conduta explicada
	DeckSequence DeckSlideKind = "sequence"  // os próximos meses, em ordem
	DeckTakeaway DeckSlideKind = "takeaway"  // o que começa a tomar agora
	DeckClosing  DeckSlideKind = "closing"   // o fecho, em uma página
	// DeckTableKind — tabela densa: um item por linha, com o porquê e um selo. Nasceu do deck do
	// José Ricardo, onde 8 dos 20 slides são isso (os exames que faltam, o prato, o movimento, a
	// sequência). A gramática v2 tinha 8 blocos e esse escapou.
	DeckTableKind DeckSlideKind = "table"
)

// DeckVariant — o fundo do slide.
const (
	DeckVariantLight = ""
	DeckVariantDark  = "dark"
	DeckVariantDeep  = "deep"
)

// DeckSlide — um slide. Os campos comuns valem para todos os tipos; os blocos específicos são
// preenchidos conforme o Kind. É esta struct que vai e volta do `patient_plans.content` (JSONB),
// então os nomes JSON são o contrato com o front e com a skill que escreve o plano.
type DeckSlide struct {
	// ID — identidade estável do slide, opaca, atribuída no servidor.
	//
	// Sem ela o slide só é endereçável pelo índice no array, e índice muda quando alguém
	// reordena. Uma sugestão criada sobre "o slide 6" e aceita depois de um reorder escreveria
	// no slide errado, sem erro e sem log, num documento que o paciente lê. É a mesma classe de
	// falha do `overflow:hidden`, que engole conteúdo em silêncio.
	//
	// Vazio na entrada é tolerado (deck escrito à mão, plano antigo): o serviço preenche ao
	// carregar. Quem edita nunca inventa um ID: copia o que veio.
	ID string `json:"id,omitempty"`

	Kind    DeckSlideKind `json:"kind"`
	Variant string        `json:"variant,omitempty"`

	Eyebrow string `json:"eyebrow,omitempty"`
	Title   string `json:"title,omitempty"`
	Lede    string `json:"lede,omitempty"`
	Kicker  string `json:"kicker,omitempty"`
	Source  string `json:"source,omitempty"`
	Punch   string `json:"punch,omitempty"`

	Rulers  []DeckRulerBlock  `json:"rulers,omitempty"`
	Legend  bool              `json:"legend,omitempty"`
	Summary *DeckSummaryBlock `json:"summary,omitempty"`
	Cards   []DeckCard        `json:"cards,omitempty"`
	Steps   []DeckSeqStep     `json:"steps,omitempty"`
	Take    *DeckTakeawayBox  `json:"takeaway,omitempty"`
	Table   *DeckTable        `json:"table,omitempty"`
}

// DeckTable — a tabela densa. Cada linha é um item com o porquê ao lado; é como o deck explica
// "o que cada exame decide" ou "o que muda no prato" sem virar parede de texto.
//
// O modelo é por COLUNA, e não um punhado de campos fixos (item/dose/porquê), porque o deck real
// usa combinações diferentes: "exame | decide o quê | prioridade" num slide, "o quê | quanto | por
// quê" no outro. Com campos fixos, a prosa da terceira coluna caía no campo de dose, que tem
// `white-space:nowrap`, e o slide vazava 2472px para a direita.
type DeckTable struct {
	Columns []DeckTableCol `json:"columns,omitempty"`
	// Rows — as células na ordem das colunas. A primeira sai em destaque.
	Rows []DeckTableRow `json:"rows,omitempty"`
	// Dense aperta o espaçamento quando a tabela é longa.
	Dense bool `json:"dense,omitempty"`
}

// DeckTableColStyle — como a coluna se comporta.
type DeckTableColStyle string

const (
	DeckColText DeckTableColStyle = ""     // texto normal, quebra linha
	DeckColWhy  DeckTableColStyle = "why"  // explicação, em cinza
	DeckColDose DeckTableColStyle = "dose" // número; NÃO quebra, então nunca use para prosa
	DeckColTag  DeckTableColStyle = "tag"  // selo curto ("prioridade")
)

// DeckTableCol — cabeçalho e comportamento de uma coluna.
type DeckTableCol struct {
	Label string            `json:"label,omitempty"`
	Style DeckTableColStyle `json:"style,omitempty"`
	Width string            `json:"width,omitempty"` // ex.: "390px"; vazio = automático
}

// DeckTableRow — uma linha da tabela.
type DeckTableRow struct {
	Cells []string `json:"cells"`
	// Muted risca a linha: o item que foi considerado e descartado.
	Muted bool `json:"muted,omitempty"`
}

// DeckRulerBlock — uma régua no slide, com o nome em voz de paciente.
//
// Display e Note são do autor, não do catálogo: "Ferritina" e "estoque de ferro" comunicam; o nome
// do catálogo ("Ferritina - Homens") não. A Note é onde entra o RÓTULO AVALIATIVO quando o título
// do slide não o traz.
type DeckRulerBlock struct {
	Ruler
}

// DeckSummaryLine — uma linha do resumo: nome, mini-régua e valor.
type DeckSummaryLine struct {
	Name string `json:"name"`
	Sub  string `json:"sub,omitempty"`
	// Code — o exame de onde `Value` saiu, no vocabulário de `lab_test_definitions.code`.
	//
	// É o que torna o número da linha auditável. Sem ele `Value` é uma string solta sem âncora
	// no dossiê, e conferir a proveniência do número passa a ser impossível — não difícil,
	// impossível. A linha de resumo é justamente onde o deck concentra mais número por
	// centímetro, então é o pior lugar para não ter âncora.
	//
	// Opcional por compatibilidade com os planos que já existem; obrigatório em linha nova
	// criada pelo assistente.
	Code  string `json:"code,omitempty"`
	Value string `json:"value"`
	Unit  string `json:"unit,omitempty"`
	// Ruler é a escala para a mini-régua; sem ela a linha sai sem barra.
	Ruler *Ruler   `json:"ruler,omitempty"`
	Plot  *float64 `json:"plot,omitempty"`
}

// DeckSummaryCard — um dos cartões do resumo ("o que está forte" / "o que está se movendo").
type DeckSummaryCard struct {
	Title string            `json:"title"`
	Tone  string            `json:"tone,omitempty"` // "bom" | "ruim"
	Lines []DeckSummaryLine `json:"lines,omitempty"`
}

// DeckSummaryBlock — o slide de resumo: os cartões e, embaixo, os passos do que vai ser feito.
type DeckSummaryBlock struct {
	Cards      []DeckSummaryCard `json:"cards,omitempty"`
	Legend     string            `json:"legend,omitempty"`
	StepsTitle string            `json:"stepsTitle,omitempty"`
	Steps      []string          `json:"steps,omitempty"`
}

// DeckCard — um cartão de texto. Dim apaga o caminho descartado; Focus destaca o que vale.
type DeckCard struct {
	Kicker string `json:"kicker,omitempty"`
	Body   string `json:"body,omitempty"`
	Dim    bool   `json:"dim,omitempty"`
	Focus  bool   `json:"focus,omitempty"`
}

// DeckSeqStep — uma linha da sequência ("nas próximas 4 semanas → …").
type DeckSeqStep struct {
	When   string `json:"when"`
	What   string `json:"what"`
	Detail string `json:"detail,omitempty"`
}

// DeckDose — um item para levar, com a dose.
type DeckDose struct {
	Name string `json:"name"`
	Sub  string `json:"sub,omitempty"`
	Dose string `json:"dose,omitempty"`
}

// DeckDoseGroup — um bloco de doses por momento ("de manhã", "todo dia", "na semana").
type DeckDoseGroup struct {
	Title string     `json:"title"`
	Items []DeckDose `json:"items,omitempty"`
}

// DeckHighlight — o item em destaque do "para levar" (o que muda o tratamento).
type DeckHighlight struct {
	When string `json:"when,omitempty"`
	Name string `json:"name"`
	Obs  string `json:"obs,omitempty"`
	Dose string `json:"dose,omitempty"`
	Unit string `json:"unit,omitempty"`
}

// DeckTakeawayBox — o slide "o que você começa a tomar agora".
type DeckTakeawayBox struct {
	Highlight *DeckHighlight  `json:"highlight,omitempty"`
	Groups    []DeckDoseGroup `json:"groups,omitempty"`
	Note      string          `json:"note,omitempty"`
}

// renderBlocks desenha a parte específica de cada tipo de slide.
func renderBlocks(s DeckSlide) string {
	var b strings.Builder
	switch s.Kind {
	case DeckRulers:
		for i := range s.Rulers {
			r := s.Rulers[i].Ruler
			r.Dark = isDarkVariant(s.Variant)
			b.WriteString(rulerSVG(r))
		}
		if s.Legend {
			b.WriteString(rampLegend())
		}
	case DeckSummary:
		if s.Summary != nil {
			b.WriteString(renderSummary(*s.Summary, isDarkVariant(s.Variant)))
		}
	case DeckTableKind:
		b.WriteString(renderTable(s.Table))
	case DeckTwoCards:
		b.WriteString(renderCards(s.Cards))
		b.WriteString(renderTable(s.Table))
	case DeckSequence:
		b.WriteString(renderSequence(s.Steps))
	case DeckTakeaway:
		if s.Take != nil {
			b.WriteString(renderTakeaway(*s.Take))
		}
	case DeckPlanStep:
		// Uma conduta é prosa + cartões + tabela: o texto vem de Lede/Kicker, o resto daqui.
		b.WriteString(renderCards(s.Cards))
		b.WriteString(renderTable(s.Table))
		if s.Take != nil {
			b.WriteString(renderTakeaway(*s.Take))
		}
	case DeckCover, DeckClosing:
		// Só tipografia; nada além do envelope.
	}
	return b.String()
}

func isDarkVariant(v string) bool { return v == DeckVariantDark || v == DeckVariantDeep }

// rampLegend — a única legenda da régua: a rampa inteira é a escala, do pior ao ótimo. Sem rótulo
// por faixa, que é o que mantém o slide com uma ideia só.
func rampLegend() string {
	var b strings.Builder
	b.WriteString(`<div class="rampa"><span>pior</span>`)
	for _, c := range rulerRamp {
		b.WriteString(`<i style="background:` + c + `"></i>`)
	}
	b.WriteString(`<span>ótimo</span><em>a régua inteira é a escala Plenya para este exame</em></div>`)
	return b.String()
}

func renderSummary(s DeckSummaryBlock, dark bool) string {
	var b strings.Builder
	b.WriteString(`<div class="rez">`)
	for _, c := range s.Cards {
		tone := ""
		switch c.Tone {
		case "bom", "ruim":
			tone = " " + c.Tone
		}
		b.WriteString(`<div class="rez-card` + tone + `"><h4>` + esc(c.Title) + `</h4>`)
		for _, ln := range c.Lines {
			b.WriteString(`<div class="rez-linha"><div class="n">` + esc(ln.Name))
			if ln.Sub != "" {
				b.WriteString(`<small>` + esc(ln.Sub) + `</small>`)
			}
			b.WriteString(`</div>`)
			if ln.Ruler != nil {
				r := *ln.Ruler
				r.Dark = dark
				v, has := 0.0, false
				if ln.Plot != nil {
					v, has = *ln.Plot, true
				} else if n := len(r.History); n > 0 {
					v, has = r.History[n-1].Value, true
				}
				b.WriteString(miniRulerSVG(r, v, has))
			} else {
				b.WriteString(`<span></span>`)
			}
			b.WriteString(`<div class="v">` + esc(ln.Value))
			if ln.Unit != "" {
				b.WriteString(` <span style="font-weight:400;color:var(--muted)">` + esc(ln.Unit) + `</span>`)
			}
			b.WriteString(`</div></div>`)
		}
		b.WriteString(`</div>`)
	}
	if s.Legend != "" {
		b.WriteString(`<div class="rez-legenda">` + esc(s.Legend) + `</div>`)
	}
	if len(s.Steps) > 0 {
		title := s.StepsTitle
		if title == "" {
			title = "O que vamos fazer"
		}
		b.WriteString(`<div class="rez-cond"><h4>` + esc(title) + `</h4><div class="rez-passos">`)
		for i, st := range s.Steps {
			fmt.Fprintf(&b, `<div class="rez-passo"><div class="num">%d</div><div class="txt">%s</div></div>`,
				i+1, inlineHTML(st))
		}
		b.WriteString(`</div></div>`)
	}
	b.WriteString(`</div>`)
	return b.String()
}

func renderCards(cards []DeckCard) string {
	if len(cards) == 0 {
		return ""
	}
	var b strings.Builder
	if len(cards) == 2 {
		b.WriteString(`<div class="two">`)
	} else {
		b.WriteString(`<div class="cards" style="grid-template-columns:repeat(` +
			fmt.Sprint(len(cards)) + `,1fr);display:grid;gap:40px">`)
	}
	for _, c := range cards {
		cls := "card"
		if c.Dim {
			cls += " dim"
		}
		if c.Focus {
			cls += " focus"
		}
		b.WriteString(`<div class="` + cls + `">`)
		if c.Kicker != "" {
			b.WriteString(`<div class="k">` + esc(c.Kicker) + `</div>`)
		}
		if c.Body != "" {
			b.WriteString(`<div class="s">` + inlineHTML(c.Body) + `</div>`)
		}
		b.WriteString(`</div>`)
	}
	b.WriteString(`</div>`)
	return b.String()
}

func renderSequence(steps []DeckSeqStep) string {
	if len(steps) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString(`<div class="seq">`)
	for _, st := range steps {
		b.WriteString(`<div class="seq-row"><div class="seq-when">` + esc(st.When) + `</div>`)
		b.WriteString(`<div class="seq-what">` + inlineHTML(st.What))
		if st.Detail != "" {
			b.WriteString(`<small>` + inlineHTML(st.Detail) + `</small>`)
		}
		b.WriteString(`</div></div>`)
	}
	b.WriteString(`</div>`)
	return b.String()
}

func renderTakeaway(t DeckTakeawayBox) string {
	var b strings.Builder
	b.WriteString(`<div class="rx2">`)
	if h := t.Highlight; h != nil {
		b.WriteString(`<div class="rx2-destaque"><div>`)
		if h.When != "" {
			b.WriteString(`<div class="quando">` + esc(h.When) + `</div>`)
		}
		b.WriteString(`<div class="nome">` + esc(h.Name) + `</div>`)
		if h.Obs != "" {
			b.WriteString(`<div class="obs">` + inlineHTML(h.Obs) + `</div>`)
		}
		b.WriteString(`</div><div class="dose">` + esc(h.Dose))
		if h.Unit != "" {
			b.WriteString(`<small>` + esc(h.Unit) + `</small>`)
		}
		b.WriteString(`</div></div>`)
	}
	for _, g := range t.Groups {
		b.WriteString(`<div class="rx2-card"><h4>` + esc(g.Title) + `</h4>`)
		for _, it := range g.Items {
			b.WriteString(`<div class="item"><b>` + esc(it.Name))
			if it.Sub != "" {
				b.WriteString(`<span class="sub">` + esc(it.Sub) + `</span>`)
			}
			b.WriteString(`</b><span class="d">` + esc(it.Dose) + `</span></div>`)
		}
		b.WriteString(`</div>`)
	}
	b.WriteString(`</div>`)
	if t.Note != "" {
		b.WriteString(`<p class="rx2-nota">` + inlineHTML(t.Note) + `</p>`)
	}
	return b.String()
}

// inlineHTML libera um punhado de tags de ênfase e escapa TODO o resto.
//
// O texto do deck vem do `patient_plans.content`, escrito por quem monta o plano, e precisa de
// ênfase de verdade: o `<em>` do punch é o que muda de cor na frase de fechamento, e o `<br>`
// separa linhas dentro de um cartão. Mas é conteúdo de origem externa entrando num HTML que o
// Chromium executa, então a régra é a inversa da intuitiva: escapa tudo primeiro e devolve à mão
// só a lista curta abaixo. Tag desconhecida aparece como texto, que é o comportamento seguro.
var inlineAllowed = []string{"em", "b", "strong", "i", "small"}

func inlineHTML(s string) string {
	out := esc(s)
	for _, tag := range inlineAllowed {
		out = strings.ReplaceAll(out, "&lt;"+tag+"&gt;", "<"+tag+">")
		out = strings.ReplaceAll(out, "&lt;/"+tag+"&gt;", "</"+tag+">")
	}
	out = strings.ReplaceAll(out, "&lt;br&gt;", "<br>")
	out = strings.ReplaceAll(out, "&lt;br/&gt;", "<br>")
	out = strings.ReplaceAll(out, "&lt;br /&gt;", "<br>")
	return out
}

// renderTable desenha a tabela densa.
func renderTable(t *DeckTable) string {
	if t == nil || len(t.Rows) == 0 {
		return ""
	}
	cls := "cond"
	if t.Dense {
		cls += " dense"
	}
	var b strings.Builder
	b.WriteString(`<table class="` + cls + `">`)
	if len(t.Columns) > 0 {
		b.WriteString(`<tr>`)
		for _, c := range t.Columns {
			w := ""
			if c.Width != "" {
				w = ` style="width:` + esc(c.Width) + `"`
			}
			b.WriteString(`<th` + w + `>` + esc(c.Label) + `</th>`)
		}
		b.WriteString(`</tr>`)
	}
	for _, r := range t.Rows {
		if r.Muted {
			b.WriteString(`<tr class="out">`)
		} else {
			b.WriteString(`<tr>`)
		}
		for i, cell := range r.Cells {
			estilo := DeckColText
			if i < len(t.Columns) {
				estilo = t.Columns[i].Style
			}
			switch {
			case i == 0:
				b.WriteString(`<td><b>` + inlineHTML(cell) + `</b></td>`)
			case estilo == DeckColTag:
				if cell == "" {
					b.WriteString(`<td></td>`)
				} else {
					b.WriteString(`<td><span class="gr">` + esc(cell) + `</span></td>`)
				}
			case estilo == DeckColDose:
				b.WriteString(`<td class="dose">` + inlineHTML(cell) + `</td>`)
			case estilo == DeckColWhy:
				b.WriteString(`<td class="why">` + inlineHTML(cell) + `</td>`)
			default:
				b.WriteString(`<td>` + inlineHTML(cell) + `</td>`)
			}
		}
		b.WriteString(`</tr>`)
	}
	b.WriteString(`</table>`)
	return b.String()
}
