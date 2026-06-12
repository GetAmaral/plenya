package pdfdoc

import (
	"strconv"
	"strings"
)

// ExamItem — um exame solicitado já resolvido (Name livre + TUSS opcional, anexado no render).
type ExamItem struct {
	Name          string
	Tuss          string // "" quando a linha não casa com o catálogo
	Justification string // "" quando não há; multi-linha separada por "\n" (linhas ">" no texto livre)
}

// ExamRequest — dados para a Solicitação de Exames (PDF). O título do PDF pode diferir do
// menu do EMR ("Pedido de Exames" no EMR, "Solicitação de Exames" no documento).
type ExamRequest struct {
	Title      string // default "Solicitação de Exames"
	Patient    Patient
	Indication string       // opcional: vazio não renderiza
	ExamPages  [][]ExamItem // páginas (quebra por linha em branco / >40); TUSS resolvido por quem monta
	Doctor     Doctor
	Signature  Signature
	Clinic     Clinic // zero value => DefaultClinic()
}

const maxPerPage = 40 // ≤20 = 1 col; 21–40 = 2 col (20/col); >40 = nova página

const justPrefix = "#" // linha de justificativa, anexa ao exame imediatamente acima
// (escolhido em vez de ">": médico usa ">"/"<" como limiar no início de frases — "# " quase nunca abre linha)

// parseExamBlocks quebra o texto livre em blocos de ExamItem (linha em branco = novo bloco/página).
// Linha iniciada por "#" é justificativa e adere ao exame de cima; linhas "#" consecutivas
// concatenam (justificativa multi-linha). "#" órfão (sem exame acima) ou vazio é ignorado.
func parseExamBlocks(text string) [][]ExamItem {
	var blocks [][]ExamItem
	var cur []ExamItem
	flush := func() {
		if len(cur) > 0 {
			blocks = append(blocks, cur)
			cur = nil
		}
	}
	for _, ln := range strings.Split(text, "\n") {
		t := strings.TrimSpace(ln)
		if t == "" {
			flush()
			continue
		}
		if strings.HasPrefix(t, justPrefix) {
			j := strings.TrimSpace(strings.TrimPrefix(t, justPrefix))
			if j == "" || len(cur) == 0 {
				continue // ">" vazio ou justificativa órfã → ignora
			}
			last := &cur[len(cur)-1]
			if last.Justification == "" {
				last.Justification = j
			} else {
				last.Justification += "\n" + j
			}
			continue
		}
		cur = append(cur, ExamItem{Name: t})
	}
	flush()
	return blocks
}

// ExamPagesFromText — conveniência: texto livre (1 exame/linha, "#" = justificativa do exame de
// cima, linha em branco = nova página) para páginas de ExamItem SEM TUSS. O adaptador resolve o
// TUSS depois (LabTestMatcher), preservando a caixa de texto 100% livre. A paginação conta só
// exames (justificativas não ocupam slot).
func ExamPagesFromText(text string) [][]ExamItem {
	var pages [][]ExamItem
	for _, b := range parseExamBlocks(text) {
		for i := 0; i < len(b); i += maxPerPage {
			end := i + maxPerPage
			if end > len(b) {
				end = len(b)
			}
			pages = append(pages, b[i:end])
		}
	}
	return pages
}

// examPadding — espaçamento dinâmico: generoso com poucos exames, mínimo (1.5px) com 20+ na coluna.
func examPadding(items []ExamItem) string {
	const minPad, maxPad = 1.5, 7.0
	ctrl := len(items)
	if ctrl > 20 {
		ctrl = 20
	}
	pad := minPad + (maxPad-minPad)*float64(20-ctrl)/19.0
	if pad < minPad {
		pad = minPad
	}
	if pad > maxPad {
		pad = maxPad
	}
	return strconv.FormatFloat(pad, 'f', 2, 64)
}

func examColHTML(items []ExamItem) string {
	var b strings.Builder
	b.WriteString(`<ul class="excol">`)
	for _, it := range items {
		b.WriteString(`<li class="exitem"><span class="bullet"></span><span class="exbody"><span class="exname">`)
		b.WriteString(esc(it.Name))
		if it.Tuss != "" {
			b.WriteString(`<span class="extuss"> (` + esc(it.Tuss) + `)</span>`)
		}
		b.WriteString(`</span>`)
		if it.Justification != "" {
			b.WriteString(`<span class="exjust">` + strings.ReplaceAll(esc(it.Justification), "\n", "<br>") + `</span>`)
		}
		b.WriteString(`</span></li>`)
	}
	b.WriteString(`</ul>`)
	return b.String()
}

// examListHTML — ≤20 = 1 coluna; 21–40 = 2 colunas com 20 por coluna (preenche a 1ª até 20).
func examListHTML(items []ExamItem) string {
	pad := examPadding(items)
	var cols string
	if len(items) <= 20 {
		cols = examColHTML(items)
	} else {
		cols = examColHTML(items[:20]) + examColHTML(items[20:])
	}
	return `<div class="exwrap" style="--expad:` + pad + `px">` + cols + `</div>`
}

// RenderExamRequest gera o PDF vetorial da Solicitação de Exames.
func RenderExamRequest(in ExamRequest) ([]byte, error) {
	if in.Title == "" {
		in.Title = "Solicitação de Exames"
	}
	if (in.Clinic == Clinic{}) {
		in.Clinic = DefaultClinic()
	}
	if len(in.ExamPages) == 0 {
		in.ExamPages = [][]ExamItem{{}}
	}
	indic := indicationHTML(in.Indication)
	foot := signatureHTML(in.Doctor, in.Signature) + footerNAPHTML(in.Clinic)

	var body strings.Builder
	for _, items := range in.ExamPages {
		top := headerHTML() + titleHTML(in.Title) + patientHTML(in.Patient) + indic +
			`<div class="sec"><span class="eyebrow">Exames solicitados</span>` + examListHTML(items) + `</div>`
		body.WriteString(pageHTML(top, foot))
	}
	return renderHTMLToPDF(documentHTML(body.String()), a4Options())
}
