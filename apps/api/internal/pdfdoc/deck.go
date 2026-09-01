package pdfdoc

import (
	"fmt"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// Deck — a devolutiva de resultados do paciente, em slides.
//
// Não passa pelo motor de papelaria (renderDocument): aquele é um paginador de FLUXO A4, que mede
// blocos e abre página quando estoura. Aqui é o contrário — o slide tem tamanho fixo de 1920×1080,
// cada um é composto à mão, e transbordo é erro a corrigir, não a paginar. O que os dois
// compartilham é a camada cara: RenderHTML, com Chromium único, serializado e com timeout.

// Dimensões do slide, em px CSS.
const (
	SlideW = 1920.0
	SlideH = 1080.0
)

// DeckPaper — em que papel o deck sai.
type DeckPaper string

const (
	// DeckPaper169 — uma página por slide, 1920×1080. Para apresentar e para mandar.
	DeckPaper169 DeckPaper = "16:9"
	// DeckPaperA4 — A4 paisagem, o slide encaixado na largura e centralizado. Para imprimir.
	DeckPaperA4 DeckPaper = "a4"
)

// A4 paisagem a 96dpi, em px CSS. O slide é escalado para caber na LARGURA (a folha é mais
// "quadrada" que 16:9) e centralizado, com a sobra preenchida pelo próprio fundo do slide.
const (
	a4LandW = 1122.52
	a4LandH = 793.70
)

// a4SlideScale — o fator de encaixe do slide na folha A4 paisagem. Sai daqui e de nenhum outro
// lugar: número derivado tem que ter uma fonte só.
func a4SlideScale() float64 { return a4LandW / SlideW }

// Deck — o documento inteiro.
type Deck struct {
	Title  string
	Slides []DeckSlide
}

// RenderDeck gera o PDF do deck no papel pedido.
func RenderDeck(d Deck, paper DeckPaper) ([]byte, error) {
	html, err := deckHTML(d)
	if err != nil {
		return nil, err
	}
	switch paper {
	case DeckPaperA4:
		return renderHTMLToPDFHook(html, a4LandscapeOptions(), fitSlidesToA4)
	case DeckPaper169, "":
		return renderHTMLToPDFHook(html, slideOptions(), awaitFonts)
	default:
		return nil, fmt.Errorf("papel de deck desconhecido: %q", paper)
	}
}

// awaitFonts segura a impressão até as webfontes carregarem.
//
// Sem isso o layout é medido com a fonte de fallback, cujas métricas não são as da Fraunces nem as
// da Inter. Era o único render do pacote que não esperava, e o efeito era pior do que "ficou um
// pouco diferente": a conferência de estouro mede COM as fontes e diz que cabe, o PDF imprime SEM
// elas e o `overflow:hidden` corta o excesso em silêncio — exatamente o que a conferência existe
// para impedir. Dava também para o A4 e o 16:9 do mesmo conteúdo saírem diferentes.
func awaitFonts(page *rod.Page) error {
	_, err := page.Eval(`async () => {
		await Promise.all(Array.from(document.fonts).map(f => f.load().catch(() => {})));
		await document.fonts.ready;
	}`)
	return err
}

// DeckHTML devolve o HTML do deck, para a tela do portal e para conferência visual.
func DeckHTML(d Deck) (string, error) { return deckHTML(d) }

func deckHTML(d Deck) (string, error) {
	if len(d.Slides) == 0 {
		return "", fmt.Errorf("deck sem slides")
	}
	var b strings.Builder
	for i := range d.Slides {
		b.WriteString(renderSlide(d.Slides[i], i+1, len(d.Slides)))
	}
	title := d.Title
	if strings.TrimSpace(title) == "" {
		title = "Plenya"
	}
	return `<!doctype html><html lang="pt-BR"><head><meta charset="utf-8"><title>` + esc(title) +
		`</title><style>` + deckFontFaces() + deckCSS + `</style></head><body>` + b.String() + `</body></html>`, nil
}

// slideOptions — uma página por slide, no tamanho exato do slide.
func slideOptions() *proto.PagePrintToPDF {
	w := SlideW / 96.0
	h := SlideH / 96.0
	zero := 0.0
	scale := 1.0
	return &proto.PagePrintToPDF{
		PrintBackground: true, PreferCSSPageSize: false,
		PaperWidth: &w, PaperHeight: &h,
		MarginTop: &zero, MarginBottom: &zero, MarginLeft: &zero, MarginRight: &zero,
		Scale: &scale, DisplayHeaderFooter: false,
	}
}

// a4LandscapeOptions — 297×210mm.
func a4LandscapeOptions() *proto.PagePrintToPDF {
	w := 297.0 / 25.4
	h := 210.0 / 25.4
	zero := 0.0
	scale := 1.0
	return &proto.PagePrintToPDF{
		PrintBackground: true, PreferCSSPageSize: false,
		PaperWidth: &w, PaperHeight: &h,
		MarginTop: &zero, MarginBottom: &zero, MarginLeft: &zero, MarginRight: &zero,
		Scale: &scale, DisplayHeaderFooter: false,
	}
}

// fitSlidesToA4 é o MESMO HTML impresso em folha diferente: envolve cada slide numa caixa do
// tamanho da folha, escala o slide para a largura e centraliza. Não existe segundo HTML para o A4.
func fitSlidesToA4(page *rod.Page) error {
	_, err := page.Eval(`async (pageW, pageH, scale) => {
		await Promise.all(Array.from(document.fonts).map(f => f.load().catch(() => {})));
		await document.fonts.ready;
		for (const slide of Array.from(document.querySelectorAll('section.slide'))) {
			const bg = getComputedStyle(slide).backgroundColor;
			const wrap = document.createElement('div');
			wrap.className = 'a4-page';
			wrap.style.cssText = 'width:' + pageW + 'px;height:' + pageH + 'px;background:' + bg +
				';display:flex;align-items:center;justify-content:center;overflow:hidden;' +
				'page-break-after:always;break-after:page;';
			slide.style.transform = 'scale(' + scale + ')';
			slide.style.transformOrigin = 'center center';
			slide.style.flexShrink = '0';
			slide.style.margin = '0';
			slide.style.pageBreakAfter = 'auto';
			slide.parentNode.insertBefore(wrap, slide);
			wrap.appendChild(slide);
		}
		document.body.style.margin = '0';
		const st = document.createElement('style');
		st.textContent = '@page { size: 297mm 210mm; margin: 0; }';
		document.head.appendChild(st);
		return document.querySelectorAll('.a4-page').length;
	}`, a4LandW, a4LandH, a4SlideScale())
	return err
}

// renderSlide monta o envelope comum: tarja, título, corpo, frase de fechamento e numeração.
func renderSlide(s DeckSlide, page, total int) string {
	cls := "slide"
	switch s.Variant {
	case DeckVariantDark:
		cls += " dark"
	case DeckVariantDeep:
		cls += " deep"
	}

	var b strings.Builder
	fmt.Fprintf(&b, `<section class="%s" id="s%02d">`, cls, page)
	b.WriteString(`<div class="body">`)
	if s.Eyebrow != "" {
		b.WriteString(`<div class="eyebrow">` + esc(s.Eyebrow) + `</div>`)
	}
	if s.Title != "" {
		tag := "h2"
		if s.Kind == DeckCover {
			tag = "h1"
		}
		b.WriteString(`<` + tag + `>` + inlineHTML(s.Title) + `</` + tag + `>`)
	}
	if s.Lede != "" {
		b.WriteString(`<p class="lede">` + inlineHTML(s.Lede) + `</p>`)
	}
	b.WriteString(renderBlocks(s))
	if s.Kicker != "" {
		b.WriteString(`<p class="kicker">` + inlineHTML(s.Kicker) + `</p>`)
	}
	if s.Source != "" {
		b.WriteString(`<p class="src">` + inlineHTML(s.Source) + `</p>`)
	}
	b.WriteString(`</div>`)
	if s.Punch != "" {
		b.WriteString(`<div class="punch">` + inlineHTML(s.Punch) + `</div>`)
	}
	fmt.Fprintf(&b, `<div class="slide-footer"><span class="brand"></span><span class="pagenum">%02d</span></div>`, page)
	b.WriteString(`</section>`)
	return b.String()
}

// deckCSS — o sistema visual do deck, extraído do `deck.css` que era copiado entre pacientes.
// Os dois últimos decks (Ana e José Ricardo) já o usavam praticamente idêntico: 5 diferenças
// triviais em 419 linhas. A regra de composição é 1 ideia por slide.
const deckCSS = `
:root{
  /* marca — packages/brand/src/tokens/colors.ts */
  --gold:#B38645; --gold-soft:#D4A86B; --gold-deep:#8A6534;
  --petrol:#063B4F; --petrol-deep:#041F2A;
  --ocean:#417E8E; --sage:#92B8B4;
  --cream:#EAE7DA; --cream-soft:#F5F1E8;
  --ink:#0A1F26; --muted:#5A6B70;
  /* escala ordinal dos níveis: o nível É o sinal do escore */
  --n0:#84AEBD; --n1:#5F93A8; --n2:#3D7A93; --n3:#22607C; --n4:#104862; --n5:#041F2A;
  --alert:#A33A1F;
  --serif:'Cormorant Garamond',Georgia,serif;
  --display:'Fraunces',Georgia,serif;
  --sans:'Inter',system-ui,sans-serif;
  --slide-w:1920px; --slide-h:1080px; --pad:84px;
}
*{ margin:0; padding:0; box-sizing:border-box; }
html,body{ background:#111; }
body{ font-family:var(--sans); -webkit-font-smoothing:antialiased; text-rendering:geometricPrecision; }

.slide{ position:relative; width:var(--slide-w); height:var(--slide-h); padding:var(--pad);
  overflow:hidden; display:flex; flex-direction:column; page-break-after:always;
  background:var(--cream-soft); color:var(--ink); }
.slide + .slide{ margin-top:40px; }

/* Na TELA os slides são uma pilha rolável, com respiro entre eles e fundo escuro em volta.
   Na IMPRESSÃO isso não pode existir: 40px de margem somados a um slide de 1080px estouram a
   página e cada slide vira duas folhas, a segunda em branco. */
@media print{
  html,body{ background:#fff; }
  .slide + .slide{ margin-top:0; }
  .slide{ margin:0; }
}
.slide.dark{ background:var(--petrol); color:var(--cream); }
.slide.deep{ background:var(--petrol-deep); color:var(--cream); }

/* tipografia */
.eyebrow{ font-family:var(--sans); font-size:28px; font-weight:600; letter-spacing:.14em;
  text-transform:uppercase; color:var(--gold); margin-bottom:12px; }
.dark .eyebrow,.deep .eyebrow{ color:var(--gold-soft); }
h1{ font-family:var(--display); font-size:92px; font-weight:400; line-height:1.04; letter-spacing:-.02em; }
h2{ font-family:var(--display); font-size:56px; font-weight:400; line-height:1.08; letter-spacing:-.015em; margin-bottom:10px; }
h3{ font-family:var(--sans); font-size:28px; font-weight:600; letter-spacing:.02em; }
.lede{ font-family:var(--serif); font-size:34px; line-height:1.4; color:var(--muted); max-width:1180px; }
.dark .lede,.deep .lede{ color:var(--sage); }
.kicker{ font-size:28px; color:var(--muted); line-height:1.5; max-width:1300px; margin-top:24px; }
.dark .kicker,.deep .kicker{ color:var(--sage); }
.body{ flex:1; display:flex; flex-direction:column; justify-content:center; }
.body.top{ justify-content:flex-start; }

/* a frase de fechamento — a "1 frase" do slide */
.punch{ margin-top:auto; margin-bottom:26px; font-family:var(--serif); font-size:36px;
  line-height:1.24; font-weight:500; color:var(--petrol); max-width:1400px;
  padding-top:22px; border-top:2px solid rgba(179,134,69,.35); }
.dark .punch,.deep .punch{ color:var(--cream); border-top-color:rgba(212,168,107,.4); }
.punch em{ font-style:normal; color:var(--gold-deep); }
.dark .punch em,.deep .punch em{ color:var(--gold-soft); }

.src{ font-size:30px; color:var(--muted); line-height:1.5; margin-top:16px; }
.dark .src,.deep .src{ color:rgba(146,184,180,.75); }
.two{ display:grid; grid-template-columns:1fr 1fr; gap:70px; align-items:start; }
.card{ background:rgba(255,255,255,.6); border:1px solid rgba(10,31,38,.1); border-radius:10px; padding:20px 24px; }
.dark .card,.deep .card{ background:rgba(234,231,218,.06); border-color:rgba(234,231,218,.14); }
.card .k{ font-size:28px; font-weight:600; letter-spacing:.08em; text-transform:uppercase; color:var(--muted); margin-bottom:8px; }
.dark .card .k,.deep .card .k{ color:var(--sage); }
.card .s{ font-size:30px; color:var(--muted); margin-top:8px; line-height:1.45; }
.dark .card .s,.deep .card .s{ color:var(--sage); }
.card.dim{ opacity:.62; }
.card.focus{ border-color:#B3503C; border-width:2px; }

/* footer canônico — numeração */
.slide-footer{ position:absolute; bottom:30px; left:84px; right:84px; display:flex;
  justify-content:flex-end; align-items:center; font-family:var(--sans); font-size:28px;
  font-weight:600; letter-spacing:.12em; text-transform:uppercase; }
.slide-footer .brand{ display:none; }
.slide-footer .pagenum{ color:var(--gold); }
.dark .slide-footer .pagenum,.deep .slide-footer .pagenum{ color:var(--gold-soft); }

/* régua (SVG) — o átomo visual do deck */
svg.rg{ display:block; margin:0 0 12px; overflow:visible; }
svg.rg text{ font-family:'Inter',system-ui,sans-serif; }
.rg-name{ font-size:28px; font-weight:500; }
.rg-sub{ font-size:28px; font-weight:400; }
.rg-tick{ font-size:28px; font-variant-numeric:tabular-nums; }
.rg-old{ font-size:30px; font-weight:600; font-variant-numeric:tabular-nums; }
.rg-val{ font-size:32px; font-weight:600; font-variant-numeric:tabular-nums; }
.rg-unit{ font-size:28px; font-weight:400; }
.rg-note{ font-size:28px; font-style:italic; }

/* legenda da rampa — uma linha, sem rótulo por faixa */
.rampa{ display:flex; align-items:center; gap:0; margin:20px 0 2px 490px; font-size:28px; color:var(--muted); }
.rampa span{ letter-spacing:.1em; text-transform:uppercase; font-size:28px; font-weight:600; }
.rampa span:first-child{ margin-right:12px; }
.rampa span:last-of-type{ margin-left:12px; }
.rampa i{ display:block; width:60px; height:16px; }
.rampa i:first-of-type{ border-radius:6px 0 0 6px; }
.rampa i:last-of-type{ border-radius:0 6px 6px 0; }
.rampa em{ font-style:normal; font-size:28px; margin-left:26px; color:var(--muted); opacity:.8; }

/* tabela densa — "o que cada exame decide", "o que muda no prato" */
.cond{ width:100%; border-collapse:collapse; font-size:31px; margin-top:24px; }
.cond th{ text-align:left; font-size:28px; font-weight:700; letter-spacing:.12em; text-transform:uppercase;
  color:var(--muted); padding:0 14px 12px 0; border-bottom:1.5px solid rgba(10,31,38,.18); white-space:nowrap; }
.dark .cond th,.deep .cond th{ color:var(--sage); border-bottom-color:rgba(234,231,218,.25); }
.cond td{ padding:10px 22px 10px 0; border-bottom:1px solid rgba(10,31,38,.07); vertical-align:top; line-height:1.3; }
.cond.dense{ font-size:30px; }
.cond.dense td{ padding:5px 14px 5px 0; }
.cond.dense .why{ font-size:28px; }
.dark .cond td,.deep .cond td{ border-bottom-color:rgba(234,231,218,.1); }
.cond tr:last-child td{ border-bottom:none; }
.cond .dose{ font-weight:600; font-variant-numeric:tabular-nums; white-space:nowrap; }
.cond .why{ color:var(--muted); font-size:28px; }
.dark .cond .why,.deep .cond .why{ color:rgba(146,184,180,.85); }
.cond tr.out td{ opacity:.5; }
.cond tr.out .dose{ text-decoration:line-through; }
.gr{ font-size:28px; font-weight:700; letter-spacing:.06em; padding:2px 6px; border-radius:3px;
  white-space:nowrap; background:#104862; color:#fff; }

/* resumo executivo */
.rez{ display:grid; grid-template-columns:1fr 1fr; gap:30px; }
.rez-card{ border-radius:12px; padding:26px 30px 22px; background:rgba(255,255,255,.62); border:1px solid rgba(10,31,38,.11); }
.rez-card.bom{ border-left:6px solid #0E4C6B; }
.rez-card.ruim{ border-left:6px solid #B3503C; }
.rez-card h4{ font-family:var(--sans); font-size:26px; font-weight:700; letter-spacing:.14em;
  text-transform:uppercase; margin:0 0 16px; padding-bottom:10px; border-bottom:1.5px solid rgba(10,31,38,.1); }
.rez-card.bom h4{ color:#0E4C6B; }
.rez-card.ruim h4{ color:#B3503C; }
.rez-linha{ display:grid; grid-template-columns:1fr 264px 205px; align-items:center; gap:16px; padding:9px 0; }
.rez-linha + .rez-linha{ border-top:1px solid rgba(10,31,38,.07); }
.rez-linha .n{ font-size:29px; font-weight:600; line-height:1.15; }
.rez-linha .n small{ display:block; font-size:24px; font-weight:400; color:var(--muted); }
.rez-linha .v{ font-size:29px; font-weight:700; text-align:right; font-variant-numeric:tabular-nums; white-space:nowrap; }
svg.mini{ display:block; }
.rez-legenda{ grid-column:1 / -1; text-align:center; font-size:25px; color:var(--muted);
  letter-spacing:.05em; margin:-4px 0 2px; }
.rez-cond{ grid-column:1 / -1; background:var(--petrol); color:var(--cream); border-radius:12px; padding:24px 32px; margin-top:2px; }
.rez-cond h4{ font-family:var(--sans); font-size:26px; font-weight:700; letter-spacing:.14em;
  text-transform:uppercase; color:var(--gold-soft); margin:0 0 14px; }
.rez-passos{ display:grid; grid-template-columns:repeat(4,1fr); gap:26px; }
.rez-passo{ display:grid; grid-template-columns:auto 1fr; gap:14px; align-items:start; }
.rez-passo .num{ font-family:var(--display); font-size:44px; line-height:.9; color:var(--gold-soft); }
.rez-passo .txt{ font-size:27px; line-height:1.28; color:var(--cream); }

/* a sequência — os próximos meses, em ordem */
.seq{ display:grid; gap:20px; }
.seq-row{ display:grid; grid-template-columns:300px 1fr; gap:32px; align-items:baseline;
  padding:18px 0; border-top:1px solid rgba(10,31,38,.1); }
.dark .seq-row,.deep .seq-row{ border-top-color:rgba(234,231,218,.16); }
.seq-when{ font-family:var(--sans); font-size:28px; font-weight:700; letter-spacing:.1em;
  text-transform:uppercase; color:var(--gold-deep); }
.dark .seq-when,.deep .seq-when{ color:var(--gold-soft); }
.seq-what{ font-size:32px; line-height:1.3; }
.seq-what small{ display:block; font-size:27px; color:var(--muted); margin-top:6px; line-height:1.3; }
.dark .seq-what small,.deep .seq-what small{ color:var(--sage); }

/* para levar — o que tomar e quando */
.rx2{ display:grid; grid-template-columns:repeat(3,1fr); gap:28px; }
.rx2-destaque{ grid-column:1 / -1; background:var(--petrol); color:var(--cream); border-radius:12px;
  padding:32px 38px; display:grid; grid-template-columns:1fr auto; align-items:center; gap:30px; }
.rx2-destaque .quando{ font-size:26px; font-weight:700; letter-spacing:.14em; text-transform:uppercase;
  color:var(--gold-soft); margin-bottom:8px; }
.rx2-destaque .nome{ font-family:var(--display); font-size:52px; line-height:1; }
.rx2-destaque .obs{ font-size:28px; color:var(--sage); margin-top:10px; }
.rx2-destaque .dose{ font-family:var(--display); font-size:76px; line-height:.95; text-align:right; white-space:nowrap; }
.rx2-destaque .dose small{ display:block; font-family:var(--sans); font-size:26px; font-weight:500;
  color:var(--sage); letter-spacing:.02em; margin-top:8px; }
.rx2-card{ background:rgba(255,255,255,.62); border:1px solid rgba(10,31,38,.11); border-radius:12px; padding:28px 30px 24px; }
.rx2-card h4{ font-family:var(--sans); font-size:26px; font-weight:700; letter-spacing:.14em;
  text-transform:uppercase; color:var(--gold-deep); margin:0 0 14px; padding-bottom:10px;
  border-bottom:1.5px solid rgba(179,134,69,.35); }
.rx2-card .item{ display:grid; grid-template-columns:1fr auto; align-items:baseline; gap:14px; padding:7px 0; }
.rx2-card .item + .item{ border-top:1px solid rgba(10,31,38,.07); }
.rx2-card .item b{ font-size:30px; font-weight:600; line-height:1.15; }
.rx2-card .item .d{ font-size:30px; font-weight:700; font-variant-numeric:tabular-nums; white-space:nowrap; color:var(--petrol); }
.rx2-card .sub{ display:block; font-size:26px; font-weight:400; color:var(--muted); line-height:1.3; margin-top:2px; }
.rx2-nota{ font-size:28px; color:var(--muted); margin-top:24px; line-height:1.35; }
`

// RenderDeckMeasured renderiza o 16:9 E mede o transbordo na MESMA passada.
//
// A publicação media com `CheckDeckOverflow` (que renderiza um PDF inteiro só para rodar o script
// de medição) e depois renderizava de novo: três passagens pelo Chromium, que é serializado por um
// mutex global e também atende receita e pedido de exames. O hook já roda antes de imprimir, então
// medir ali sai de graça.
func RenderDeckMeasured(d Deck) ([]byte, []DeckOverflow, error) {
	html, err := deckHTML(d)
	if err != nil {
		return nil, nil, err
	}
	var over []DeckOverflow
	pdf, err := renderHTMLToPDFHook(html, slideOptions(), func(page *rod.Page) error {
		found, mErr := measureOverflow(page)
		if mErr != nil {
			return mErr
		}
		over = found
		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	return pdf, over, nil
}

// DeckOverflow — um slide cujo conteúdo passou da moldura de 1920×1080.
type DeckOverflow struct {
	Slide  int     `json:"slide"`  // 1-based, igual à numeração impressa
	Title  string  `json:"title"`  // para o autor saber qual é sem contar slides
	Right  float64 `json:"right"`  // px que vazaram à direita
	Bottom float64 `json:"bottom"` // px que vazaram embaixo
}

// CheckDeckOverflow mede, no próprio Chromium, se algum slide transborda.
//
// O slide tem altura fixa e `overflow:hidden`: conteúdo demais não empurra nada, simplesmente
// SOME na impressão, sem erro nenhum. Era o risco número um do processo antigo, contornado por um
// script solto em /tmp que dependia de alguém lembrar de rodar. Agora é função do pacote, medida
// depois das webfontes carregarem (medir antes dá altura errada), e o serviço se recusa a publicar
// um deck que transborda.
func CheckDeckOverflow(d Deck) ([]DeckOverflow, error) {
	html, err := deckHTML(d)
	if err != nil {
		return nil, err
	}
	var found []DeckOverflow
	_, err = renderHTMLToPDFHook(html, slideOptions(), func(page *rod.Page) error {
		f, mErr := measureOverflow(page)
		found = f
		return mErr
	})
	if err != nil {
		return nil, err
	}
	return found, nil
}

// measureOverflow roda a medição no Chromium, depois das webfontes carregarem: medir antes das
// fontes dá altura errada.
func measureOverflow(page *rod.Page) ([]DeckOverflow, error) {
	res, err := page.Eval(`async (W, H) => {
		await Promise.all(Array.from(document.fonts).map(f => f.load().catch(() => {})));
		await document.fonts.ready;
		const out = [];
		document.querySelectorAll('section.slide').forEach((slide, i) => {
			const box = slide.getBoundingClientRect();
			let right = 0, bottom = 0;
			slide.querySelectorAll('*').forEach(el => {
				const r = el.getBoundingClientRect();
				if (r.width === 0 && r.height === 0) return;
				right  = Math.max(right,  r.right  - box.left - W);
				bottom = Math.max(bottom, r.bottom - box.top  - H);
			});
			// 1px de tolerância: arredondamento de subpixel não é estouro.
			if (right > 1 || bottom > 1) {
				const h = slide.querySelector('h1, h2');
				out.push({ slide: i + 1, title: h ? h.textContent.trim() : '',
				           right: Math.round(right), bottom: Math.round(bottom) });
			}
		});
		return out;
	}`, SlideW, SlideH)
	if err != nil {
		return nil, err
	}
	var found []DeckOverflow
	if uErr := res.Value.Unmarshal(&found); uErr != nil {
		return nil, uErr
	}
	return found, nil
}
