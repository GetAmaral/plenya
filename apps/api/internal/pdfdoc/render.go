// Package pdfdoc gera os PDFs médicos do paciente (papelaria world-class) a partir de
// HTML/CSS renderizado por Chromium headless (go-rod) — saída 100% vetorial.
//
// Substitui a geração legada em gofpdf. O sistema-base (cabeçalho, marca d'água, selo de
// assinatura ICP-Brasil, rodapé, fontes e tokens de marca) é compartilhado por todos os
// documentos; cada documento (solicitação de exames, receituário, atestado, etc.) fornece
// só o miolo. Direção e regras: docs/emr/plano-papelaria-pdf.md.
//
// Concorrência: um único Chromium é mantido vivo e **reutilizado** entre renders, e os renders
// são **serializados** por um mutex. Lançar um Chromium novo por request, sob concorrência e
// pouca RAM, fazia as instâncias se atropelarem e travarem para sempre (os métodos Must* do rod
// não têm timeout) — deixando processos zumbis queimando CPU/RAM. Agora cada render reusa o
// browser (≈0,3–1 s), nunca há dois Chromiums disputando, e um timeout duro aborta qualquer
// render preso em vez de pendurar o handler.
package pdfdoc

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

// renderTimeout — teto duro por render. O normal é ~1 s; passou disso, aborta com erro
// (handler responde 500 na hora) em vez de pendurar para sempre.
const renderTimeout = 35 * time.Second

var (
	browserMu     sync.Mutex // serializa renders + protege o singleton abaixo
	sharedBrowser *rod.Browser
)

func chromiumBin() string {
	chromium := "/usr/bin/chromium-browser"
	if _, err := os.Stat(chromium); os.IsNotExist(err) {
		chromium, _ = launcher.LookPath()
	}
	return chromium
}

// getBrowser devolve um Chromium headless vivo, relançando se o anterior morreu.
// Sempre chamado sob browserMu.
func getBrowser() (*rod.Browser, error) {
	if sharedBrowser != nil {
		// Sonda barata de liveness; se o browser morreu, descarta e relança.
		if _, err := sharedBrowser.Timeout(3 * time.Second).Version(); err == nil {
			return sharedBrowser, nil
		}
		sharedBrowser = nil
	}
	// disable-dev-shm-usage: /dev/shm em container é 64MB; sem isto o Chromium trava em renders
	// pesados (pedido com muitas páginas) — escreve em /tmp. disable-gpu/disable-software-rasterizer
	// estabilizam o headless sob limite de memória. Foi o "new page: context deadline exceeded" em prod.
	u, err := launcher.New().Bin(chromiumBin()).Headless(true).NoSandbox(true).
		Set("disable-dev-shm-usage").
		Set("disable-gpu").
		Set("disable-software-rasterizer").
		Launch()
	if err != nil {
		return nil, fmt.Errorf("launch chromium: %w", err)
	}
	b := rod.New().ControlURL(u)
	if err := b.Connect(); err != nil {
		return nil, fmt.Errorf("connect chromium: %w", err)
	}
	sharedBrowser = b
	return b, nil
}

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

// RenderHTML converte HTML autossuficiente em PDF vetorial, com paper/margens custom (pôster,
// A4, etc.), pelo MESMO Chromium serializado que toda a papelaria usa. Ponto de entrada único
// para quem está fora do pacote (ex.: ScorePDFService) — garante uma fila/browser global, sem
// dois Chromiums disputando RAM. opts nil cai no A4 padrão.
func RenderHTML(html string, opts *proto.PagePrintToPDF) ([]byte, error) {
	if opts == nil {
		opts = a4Options()
	}
	return renderHTMLToPDF(html, opts)
}

// renderHTMLToPDF converte HTML autossuficiente (fontes/SVGs embutidos via data-URI/inline)
// em bytes de PDF vetorial usando Chromium headless. Serializado e com timeout duro.
func renderHTMLToPDF(html string, opts *proto.PagePrintToPDF) (result []byte, err error) {
	browserMu.Lock()
	defer browserMu.Unlock()

	// Qualquer panic interno do rod (timeout/conexão) vira erro — nunca derruba a API.
	// E em QUALQUER falha (panic OU erro retornado, ex.: "new page: context deadline exceeded"),
	// descarta o browser compartilhado: ele pode ter ficado wedged e, se mantido no singleton,
	// faria todo render seguinte falhar para sempre. Nulificar força relançar um Chromium limpo.
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("pdf render: %v", r)
		}
		if err != nil && sharedBrowser != nil {
			old := sharedBrowser
			sharedBrowser = nil
			go old.Close() // assíncrono: não trava o request se o browser estiver pendurado
		}
	}()

	b, err := getBrowser()
	if err != nil {
		return nil, err
	}

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

	ctx, cancel := context.WithTimeout(context.Background(), renderTimeout)
	defer cancel()

	page, err := b.Context(ctx).Page(proto.TargetCreateTarget{})
	if err != nil {
		return nil, fmt.Errorf("new page: %w", err)
	}
	defer page.Close()
	page = page.Context(ctx)

	if err := page.Navigate("file://" + tmp.Name()); err != nil {
		return nil, fmt.Errorf("navigate: %w", err)
	}
	if err := page.WaitLoad(); err != nil {
		return nil, fmt.Errorf("wait load: %w", err)
	}

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
