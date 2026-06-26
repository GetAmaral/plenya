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

// examItemsHTML — cada exame é um bloco .exitem independente (paginado pelo motor único).
// examPadding — espaçamento dinâmico por item: generoso com poucos exames, mínimo com 20+ na coluna.
func examPadding(n int) string {
	const minPad, maxPad = 1.5, 7.0
	if n > 20 {
		n = 20
	}
	pad := minPad + (maxPad-minPad)*float64(20-n)/19.0
	if pad < minPad {
		pad = minPad
	}
	if pad > maxPad {
		pad = maxPad
	}
	return strconv.FormatFloat(pad, 'f', 2, 64)
}

func examItemHTML(it ExamItem) string {
	var b strings.Builder
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
	return b.String()
}

func examColHTML(items []ExamItem) string {
	var b strings.Builder
	b.WriteString(`<ul class="excol">`)
	for _, it := range items {
		b.WriteString(examItemHTML(it))
	}
	b.WriteString(`</ul>`)
	return b.String()
}

// examListHTML — ≤20 exames numa página: 1 coluna; >20: 2 colunas balanceadas (mecanismo do EMR).
func examListHTML(items []ExamItem) string {
	pad := examPadding(len(items))
	var cols string
	if len(items) <= 20 {
		cols = examColHTML(items)
	} else {
		mid := (len(items) + 1) / 2
		cols = examColHTML(items[:mid]) + examColHTML(items[mid:])
	}
	return `<div class="exwrap" style="--expad:` + pad + `px">` + cols + `</div>`
}

// RenderExamRequest gera o PDF da Solicitação de Exames pelo motor único.
func RenderExamRequest(in ExamRequest) ([]byte, error) {
	if in.Title == "" {
		in.Title = "Solicitação de Exames"
	}
	// Grupos de ExamPages são quebras de página INTENCIONAIS (linha em branco no texto): cada grupo
	// começa em página nova (ex.: laboratoriais × imagem). Preserva a separação original.
	var body strings.Builder
	body.WriteString(indicationHTML(in.Indication))
	for i, group := range in.ExamPages {
		if len(group) == 0 {
			continue
		}
		if i > 0 {
			body.WriteString(`<div class="page-break"></div>`)
		}
		body.WriteString(`<div class="sec"><span class="eyebrow">Exames solicitados</span></div>`)
		body.WriteString(examListHTML(group))
	}
	return renderDocument(Doc{
		Title:     in.Title,
		Patient:   &in.Patient,
		Body:      body.String(),
		Signature: signatureBlock(in.Doctor, in.Signature),
		Clinic:    in.Clinic,
	})
}
