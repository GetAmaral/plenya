// Package pdfdoc gera os PDFs médicos do paciente (papelaria world-class) a partir de
// HTML/CSS renderizado por Chromium headless (go-rod) — saída 100% vetorial.
//
// Substitui a geração legada em gofpdf. O sistema-base (cabeçalho, marca d'água, selo de
// assinatura ICP-Brasil, rodapé, fontes e tokens de marca) é compartilhado por todos os
// documentos; cada documento (solicitação de exames, receituário, atestado, etc.) fornece
// só o miolo. Direção e regras: docs/emr/plano-papelaria-pdf.md.
package pdfdoc

import (
	"fmt"
	"io"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

// a4Options — A4 retrato, margens zero (as margens vêm do CSS .frame; cada .page = 210×297mm
// com page-break-after). Espelha o padrão provado de score_pdf_service.a4PDFOptions.
func a4Options() *proto.PagePrintToPDF {
	w := 8.27  // 210mm
	h := 11.69 // 297mm
	zero := 0.0
	scale := 1.0
	return &proto.PagePrintToPDF{
		PrintBackground:     true,
		PreferCSSPageSize:   false,
		PaperWidth:          &w,
		PaperHeight:         &h,
		MarginTop:           &zero,
		MarginBottom:        &zero,
		MarginLeft:          &zero,
		MarginRight:         &zero,
		Scale:               &scale,
		DisplayHeaderFooter: false,
	}
}

// renderHTMLToPDF converte HTML autossuficiente (fontes/SVGs embutidos via data-URI/inline)
// em bytes de PDF vetorial usando Chromium headless.
func renderHTMLToPDF(html string, opts *proto.PagePrintToPDF) ([]byte, error) {
	chromium := "/usr/bin/chromium-browser"
	if _, err := os.Stat(chromium); os.IsNotExist(err) {
		chromium, _ = launcher.LookPath()
	}

	u := launcher.New().Bin(chromium).Headless(true).NoSandbox(true).MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage()
	defer page.MustClose()

	// HTML grande vai por arquivo temporário (data: URL estoura limite em docs longos).
	tmp, err := os.CreateTemp("", "plenya-doc-*.html")
	if err != nil {
		return nil, fmt.Errorf("temp file: %w", err)
	}
	defer os.Remove(tmp.Name())
	if _, err := tmp.WriteString(html); err != nil {
		return nil, fmt.Errorf("write temp: %w", err)
	}
	tmp.Close()

	page.MustNavigate("file://" + tmp.Name()).MustWaitLoad()

	stream, err := page.PDF(opts)
	if err != nil {
		return nil, fmt.Errorf("print to pdf: %w", err)
	}
	data, err := io.ReadAll(stream)
	if err != nil {
		return nil, fmt.Errorf("read pdf stream: %w", err)
	}
	return data, nil
}
