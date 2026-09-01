package pdfdoc

import (
	"fmt"
	"strings"
)

// O relatório é o TERCEIRO modo do deck: mesmo conteúdo, papel diferente.
//
// Os outros dois (16:9 e A4 paisagem) são o slide de 1920×1080 impresso inteiro. Este achata os
// mesmos slides no documento A4 fluido da papelaria — cabeçalho, rodapé e assinatura ICP em toda
// página, paginação medida in situ pelo `renderDocument`. É o modo que vale como documento clínico
// assinado; os de slide são peça de comunicação e não levam assinatura.
//
// O que muda entre os modos é só a moldura: nenhum texto é reescrito, nenhuma régua é redesenhada.
// A régua é SVG com viewBox, então encolher de 1714px para os 170mm do miolo é só largura.

// PlanReport — o plano rendido como documento assinado.
type PlanReport struct {
	Title     string
	Patient   Patient
	Slides    []DeckSlide
	EmittedAt string
	Doctor    Doctor
	Signature Signature
	Clinic    Clinic
}

// reportDeckCSS — só o que o achatamento precisa. Cabeçalho, rodapé, marca d'água, título e
// paginação vêm do motor único.
const reportDeckCSS = `
.pr-sec{ margin-top:16px; break-inside:avoid; }
.pr-eyebrow{ font-size:10.5px; letter-spacing:1.4px; text-transform:uppercase; color:var(--gold);
  font-weight:700; margin-bottom:3px; }
.pr-title{ font-family:'Cormorant',serif; font-size:19px; font-weight:600; color:var(--petrol);
  line-height:1.15; margin-bottom:5px; }
.pr-lede{ font-size:12.5px; line-height:1.45; color:var(--petrol); }
.pr-kicker{ font-size:12px; line-height:1.45; color:var(--ink2); margin-top:5px; }
.pr-punch{ font-family:'Cormorant',serif; font-size:14px; line-height:1.3; color:var(--petrol);
  margin-top:8px; padding-top:7px; border-top:1px solid rgba(179,134,69,.35); }
.pr-punch em{ font-style:normal; color:var(--gold-deep,#8a6534); }
.pr-src{ font-size:10.5px; color:var(--ink2); margin-top:5px; line-height:1.4; }
/* a régua encolhe pela largura; o viewBox cuida do resto */
.pr-ruler{ margin:6px 0; }
.pr-ruler svg{ width:100%; height:auto; display:block; }
/* Os tamanhos de texto DENTRO da régua precisam ser repetidos aqui.
   Eles vivem no deckCSS, que este documento não carrega, e sem eles o <text> herda o corpo do
   documento (~13px). Como o valor é lido em unidades do viewBox (1714 de largura, exibidas em
   ~170mm), 13 unidades viram menos de 5px impressos: nome do exame, valor e marcações ficam
   ilegíveis num documento clínico assinado. Os números abaixo são os mesmos do deck, de propósito:
   é o viewBox que faz a redução, não uma segunda escala tipográfica. */
.pr-ruler svg text{ font-family:'Inter',sans-serif; }
.pr-ruler .rg-name{ font-size:28px; font-weight:500; }
.pr-ruler .rg-sub{ font-size:28px; font-weight:400; }
.pr-ruler .rg-tick{ font-size:28px; }
.pr-ruler .rg-old{ font-size:30px; font-weight:600; }
.pr-ruler .rg-val{ font-size:32px; font-weight:600; }
.pr-ruler .rg-unit{ font-size:28px; font-weight:400; }
.pr-ruler .rg-note{ font-size:28px; font-style:italic; }
/* listas do resumo e do "para levar" */
.pr-card{ margin-top:8px; padding:8px 12px; background:rgba(6,59,79,.035);
  border-left:3px solid var(--gold); border-radius:2px; break-inside:avoid; }
.pr-card-h{ font-size:11px; font-weight:700; letter-spacing:.8px; text-transform:uppercase;
  color:var(--petrol); margin-bottom:4px; }
.pr-row{ display:flex; justify-content:space-between; gap:12px; font-size:12px; color:var(--petrol);
  padding:2.5px 0; border-bottom:1px solid rgba(6,59,79,.07); }
.pr-row:last-child{ border-bottom:none; }
.pr-row .pr-n b{ font-weight:600; }
.pr-row .pr-n small{ display:block; font-size:10.5px; color:var(--ink2); }
.pr-row .pr-v{ font-weight:600; white-space:nowrap; }
.pr-steps{ margin-top:5px; }
.pr-step{ display:flex; gap:9px; font-size:12px; color:var(--petrol); padding:2px 0; }
.pr-step .pr-num{ flex:0 0 auto; font-weight:700; color:var(--gold-deep,#8a6534); }
/* tabela densa no A4. Os tamanhos são os do documento, não os do slide: aqui não há viewBox
   reduzindo nada, o texto é impresso no tamanho que estiver escrito. */
.cond{ width:100%; border-collapse:collapse; font-size:12px; margin-top:8px; break-inside:avoid; }
.cond th{ text-align:left; font-size:10px; font-weight:700; letter-spacing:.8px; text-transform:uppercase;
  color:var(--ink2); padding:0 8px 5px 0; border-bottom:1px solid rgba(6,59,79,.25); white-space:nowrap; }
.cond td{ padding:4px 10px 4px 0; border-bottom:1px solid rgba(6,59,79,.07); vertical-align:top; line-height:1.35; }
.cond tr:last-child td{ border-bottom:none; }
.cond .why{ color:var(--ink2); }
.cond .dose{ font-weight:600; white-space:nowrap; }
.cond tr.out td{ opacity:.5; }
.gr{ font-size:9.5px; font-weight:700; letter-spacing:.4px; padding:1px 5px; border-radius:3px;
  white-space:nowrap; background:#104862; color:#fff; }

.pr-when{ font-size:10.5px; font-weight:700; letter-spacing:.8px; text-transform:uppercase;
  color:var(--gold-deep,#8a6534); }
`

// RenderPlanReport gera o relatório A4 assinado a partir dos slides do plano.
func RenderPlanReport(in PlanReport) ([]byte, error) {
	if strings.TrimSpace(in.Title) == "" {
		in.Title = "Relatório de Saúde, Performance e Longevidade"
	}
	if len(in.Slides) == 0 {
		return nil, fmt.Errorf("plano sem slides")
	}

	var b strings.Builder
	if in.EmittedAt != "" {
		b.WriteString(`<div class="rpt-dates">Relatório emitido em ` + esc(in.EmittedAt) + `</div>`)
	}
	for _, s := range in.Slides {
		b.WriteString(flattenSlide(s))
	}

	return renderDocument(Doc{
		Title:     in.Title,
		Patient:   &in.Patient,
		Body:      b.String(),
		Signature: signatureBlock(in.Doctor, in.Signature),
		ExtraCSS:  reportDeckCSS + `.rpt-dates{ font-size:12px; color:var(--ink2); margin-top:6px; }`,
		Clinic:    in.Clinic,
	})
}

// flattenSlide transforma um slide numa seção do documento fluido. Cada seção é um bloco que o
// paginador pode mover inteiro para a página seguinte (`break-inside:avoid`).
func flattenSlide(s DeckSlide) string {
	// Slide com tabela sai como VÁRIOS blocos IRMÃOS, e a decisão é a primeira coisa da função:
	// `paginateDoc` só pagina filhos DIRETOS de `#src`, então tudo aqui tem que sair no nível de
	// cima. Fatiar dentro da seção deixa as fatias como netos, o slide continua um bloco único mais
	// alto que a área útil, e desce por cima do rodapé e da assinatura.
	if temTabela(s) {
		var b strings.Builder
		b.WriteString(cabecalhoSecao(s))
		if len(s.Cards) > 0 {
			b.WriteString(`<div class="pr-sec">` + cartoesRelatorio(s.Cards) + `</div>`)
		}
		b.WriteString(fatiaTabela(s.Table))
		if s.Kind == DeckPlanStep && s.Take != nil {
			b.WriteString(`<div class="pr-sec">` + flattenTakeaway(*s.Take) + `</div>`)
		}
		b.WriteString(rodapeSecao(s))
		return b.String()
	}

	var b strings.Builder
	b.WriteString(`<div class="pr-sec">`)
	if s.Eyebrow != "" {
		b.WriteString(`<div class="pr-eyebrow">` + esc(s.Eyebrow) + `</div>`)
	}
	// O título da CAPA é o nome do paciente, que a papelaria já imprime no bloco de identificação
	// logo acima. Repetir a dois centímetros de distância é ruído; a data (eyebrow) e a frase de
	// abertura (lede) continuam, porque essas o cabeçalho não tem.
	if s.Title != "" && s.Kind != DeckCover {
		b.WriteString(`<div class="pr-title">` + inlineHTML(s.Title) + `</div>`)
	}
	if s.Lede != "" {
		b.WriteString(`<div class="pr-lede">` + inlineHTML(s.Lede) + `</div>`)
	}

	switch s.Kind {
	case DeckRulers:
		for i := range s.Rulers {
			r := s.Rulers[i].Ruler
			r.Dark = false // o documento é sempre claro, independente da variante do slide
			b.WriteString(`<div class="pr-ruler">` + rulerSVG(r) + `</div>`)
		}
	case DeckSummary:
		if s.Summary != nil {
			b.WriteString(flattenSummary(*s.Summary))
		}
	case DeckTwoCards, DeckPlanStep:
		b.WriteString(cartoesRelatorio(s.Cards))
		if s.Kind == DeckPlanStep && s.Take != nil {
			b.WriteString(flattenTakeaway(*s.Take))
		}
	case DeckSequence:
		for _, st := range s.Steps {
			b.WriteString(`<div class="pr-card"><div class="pr-when">` + esc(st.When) + `</div>`)
			b.WriteString(`<div class="pr-lede">` + inlineHTML(st.What))
			if st.Detail != "" {
				b.WriteString(`<br><span class="pr-src">` + inlineHTML(st.Detail) + `</span>`)
			}
			b.WriteString(`</div></div>`)
		}
	case DeckTakeaway:
		if s.Take != nil {
			b.WriteString(flattenTakeaway(*s.Take))
		}
	}

	fechaSecao(&b, s)
	return b.String()
}

func temTabela(s DeckSlide) bool {
	return s.Table != nil && len(s.Table.Rows) > 0 &&
		(s.Kind == DeckTableKind || s.Kind == DeckTwoCards || s.Kind == DeckPlanStep)
}

// cabecalhoSecao — tarja, título e abertura, como bloco próprio.
func cabecalhoSecao(s DeckSlide) string {
	var b strings.Builder
	b.WriteString(`<div class="pr-sec">`)
	if s.Eyebrow != "" {
		b.WriteString(`<div class="pr-eyebrow">` + esc(s.Eyebrow) + `</div>`)
	}
	if s.Title != "" && s.Kind != DeckCover {
		b.WriteString(`<div class="pr-title">` + inlineHTML(s.Title) + `</div>`)
	}
	if s.Lede != "" {
		b.WriteString(`<div class="pr-lede">` + inlineHTML(s.Lede) + `</div>`)
	}
	b.WriteString(`</div>`)
	return b.String()
}

// rodapeSecao — nota, fonte e frase de fechamento, como bloco próprio.
func rodapeSecao(s DeckSlide) string {
	if s.Kicker == "" && s.Source == "" && s.Punch == "" {
		return ""
	}
	var b strings.Builder
	b.WriteString(`<div class="pr-sec">`)
	fechaSecao(&b, s)
	return b.String()
}

// cartoesRelatorio usa a papelaria do RELATÓRIO, não a do slide.
//
// `renderCards` emite `.two`/`.card`/`.k`/`.s`, que só existem no `deckCSS`. No A4 assinado, que
// não carrega aquele CSS, os cartões sairiam sem moldura e, com três ou mais, dentro de uma grade
// inline dimensionada para um slide de 1920px, numa folha de 170mm.
func cartoesRelatorio(cards []DeckCard) string {
	var b strings.Builder
	for _, c := range cards {
		b.WriteString(`<div class="pr-card">`)
		if c.Kicker != "" {
			b.WriteString(`<div class="pr-card-h">` + esc(c.Kicker) + `</div>`)
		}
		if c.Body != "" {
			b.WriteString(`<div class="pr-lede">` + inlineHTML(c.Body) + `</div>`)
		}
		b.WriteString(`</div>`)
	}
	return b.String()
}

// fechaSecao escreve o rodapé comum do slide achatado.
func fechaSecao(b *strings.Builder, s DeckSlide) {
	if s.Kicker != "" {
		b.WriteString(`<div class="pr-kicker">` + inlineHTML(s.Kicker) + `</div>`)
	}
	if s.Source != "" {
		b.WriteString(`<div class="pr-src">` + inlineHTML(s.Source) + `</div>`)
	}
	if s.Punch != "" {
		b.WriteString(`<div class="pr-punch">` + inlineHTML(s.Punch) + `</div>`)
	}
	b.WriteString(`</div>`)
}

// linhasPorFatia — quantas linhas por bloco. Medido na folha real: com 16, duas fatias ainda
// couberam na mesma página e a última linha encostava na data da assinatura; com 12 a segunda fatia
// desce para a página seguinte e o rodapé respira.
const linhasPorFatia = 12

// fatiaTabela quebra uma tabela longa em vários blocos, cada um repetindo o cabeçalho.
func fatiaTabela(t *DeckTable) string {
	if t == nil || len(t.Rows) == 0 {
		return ""
	}
	if len(t.Rows) <= linhasPorFatia {
		return renderTable(t)
	}
	var b strings.Builder
	for i := 0; i < len(t.Rows); i += linhasPorFatia {
		fim := i + linhasPorFatia
		if fim > len(t.Rows) {
			fim = len(t.Rows)
		}
		// Cada fatia é um bloco próprio para o paginador, e repete o cabeçalho: uma tabela que
		// continua na página seguinte sem cabeçalho vira uma lista de células sem sentido.
		b.WriteString(`<div class="pr-sec">` + renderTable(&DeckTable{
			Columns: t.Columns, Rows: t.Rows[i:fim], Dense: t.Dense,
		}) + `</div>`)
	}
	return b.String()
}

func flattenSummary(s DeckSummaryBlock) string {
	var b strings.Builder
	for _, c := range s.Cards {
		b.WriteString(`<div class="pr-card"><div class="pr-card-h">` + esc(c.Title) + `</div>`)
		for _, ln := range c.Lines {
			b.WriteString(`<div class="pr-row"><span class="pr-n"><b>` + esc(ln.Name) + `</b>`)
			if ln.Sub != "" {
				b.WriteString(`<small>` + esc(ln.Sub) + `</small>`)
			}
			b.WriteString(`</span><span class="pr-v">` + esc(ln.Value))
			if ln.Unit != "" {
				b.WriteString(` ` + esc(ln.Unit))
			}
			b.WriteString(`</span></div>`)
		}
		b.WriteString(`</div>`)
	}
	if len(s.Steps) > 0 {
		title := s.StepsTitle
		if title == "" {
			title = "O que vamos fazer"
		}
		b.WriteString(`<div class="pr-card"><div class="pr-card-h">` + esc(title) + `</div><div class="pr-steps">`)
		for i, st := range s.Steps {
			fmt.Fprintf(&b, `<div class="pr-step"><span class="pr-num">%d</span><span>%s</span></div>`, i+1, inlineHTML(st))
		}
		b.WriteString(`</div></div>`)
	}
	return b.String()
}

func flattenTakeaway(t DeckTakeawayBox) string {
	var b strings.Builder
	if h := t.Highlight; h != nil {
		b.WriteString(`<div class="pr-card">`)
		if h.When != "" {
			b.WriteString(`<div class="pr-when">` + esc(h.When) + `</div>`)
		}
		b.WriteString(`<div class="pr-row"><span class="pr-n"><b>` + esc(h.Name) + `</b>`)
		if h.Obs != "" {
			b.WriteString(`<small>` + inlineHTML(h.Obs) + `</small>`)
		}
		b.WriteString(`</span><span class="pr-v">` + esc(h.Dose))
		if h.Unit != "" {
			b.WriteString(` ` + esc(h.Unit))
		}
		b.WriteString(`</span></div></div>`)
	}
	for _, g := range t.Groups {
		b.WriteString(`<div class="pr-card"><div class="pr-card-h">` + esc(g.Title) + `</div>`)
		for _, it := range g.Items {
			b.WriteString(`<div class="pr-row"><span class="pr-n"><b>` + esc(it.Name) + `</b>`)
			if it.Sub != "" {
				b.WriteString(`<small>` + esc(it.Sub) + `</small>`)
			}
			b.WriteString(`</span><span class="pr-v">` + esc(it.Dose) + `</span></div>`)
		}
		b.WriteString(`</div>`)
	}
	if t.Note != "" {
		b.WriteString(`<div class="pr-src">` + inlineHTML(t.Note) + `</div>`)
	}
	return b.String()
}
