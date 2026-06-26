package pdfdoc

import (
	"strings"

	"github.com/go-rod/rod"
)

// Doc — especificação genérica de QUALQUER documento da papelaria Plenya (receituário, pedido de
// exames, atestado/declaração/laudo, recibo, relatório). O motor renderDocument cuida de TUDO que
// é comum a todos: paginação automática, marca d'água/claim, cabeçalho (logo + régua) e rodapé
// repetidos no topo/pé de toda página, e a assinatura ancorada no pé da ÚLTIMA página. Cada
// documento só monta o seu MIOLO específico (Body) e alguns campos de título/assinatura.
//
// É a fonte ÚNICA de layout: nenhum gerador deve montar página/cabeçalho/rodapé por conta própria.
type Doc struct {
	Kind       string   // tarja de categoria acima do título (opcional)
	Title      string   // título do documento
	TitleRight string   // HTML opcional à direita do título (ex.: recibo nº + valor)
	Patient    *Patient // bloco de identificação do paciente (opcional — recibo não usa)
	Body       string   // HTML do conteúdo específico (miolo); paginado automaticamente
	Signature  string   // HTML do bloco de assinatura na última página (vazio = sem assinatura)
	Footer     string   // HTML do rodapé; vazio => NAP padrão da Clinic
	ExtraCSS   string   // CSS específico do documento (ex.: relatório); vazio na maioria
	Clinic     Clinic
}

// renderDocument é o ponto de entrada ÚNICO da papelaria: monta o scaffold (fonte + templates de
// cabeçalho/rodapé/assinatura/marca-d'água) e delega a paginação real ao Chromium via paginateDoc.
func renderDocument(d Doc) ([]byte, error) {
	if (d.Clinic == Clinic{}) {
		d.Clinic = DefaultClinic()
	}
	footer := d.Footer
	if footer == "" {
		footer = footerNAPHTML(d.Clinic)
	}

	// Cabeçalho COMPLETO (logo+régua + título + identificação do paciente) — repetido no topo de
	// TODA página, igual à primeira. O miolo (Body) é só o conteúdo específico.
	var head strings.Builder
	head.WriteString(headerHTML())
	if strings.TrimSpace(d.Title) != "" || d.Kind != "" || d.TitleRight != "" {
		head.WriteString(titleBlockHTML(d.Kind, d.Title, d.TitleRight))
	}
	if d.Patient != nil {
		head.WriteString(patientHTML(*d.Patient))
	}

	src := `<div id="src" class="src">` + d.Body + `</div>`

	tpl := `<template id="tpl-head">` + head.String() + `</template>` +
		`<template id="tpl-nap">` + footer + `</template>` +
		`<template id="tpl-sig">` + d.Signature + `</template>` +
		`<template id="tpl-wm">` + imgSVG("pattern.svg", "wm") + `<div class="claim-v">` + esc(claimText) + `</div></template>`

	html := documentHTMLCSS(src+tpl+`<div id="out"></div>`, d.ExtraCSS)
	return renderHTMLToPDFHook(html, a4Options(), paginateDoc)
}

// titleBlockHTML — bloco de título unificado: tarja de categoria (opcional) + título + conteúdo à
// direita (opcional, HTML) + a régua dourada. Substitui titleHTML/titleKindHTML antigos.
func titleBlockHTML(kind, title, right string) string {
	k := ""
	if strings.TrimSpace(kind) != "" {
		k = `<div class="title-kind">` + esc(kind) + `</div>`
	}
	return `<div class="titleblock">` + k +
		`<div class="titlerow"><div class="title">` + esc(title) + `</div>` + right + `</div>` +
		`<div class="title-rule"></div></div>`
}

// paginateDoc faz a PAGINAÇÃO REAL no navegador, em situ: adiciona os blocos do miolo a caixas A4
// fixas (.page) e mede o overflow de verdade (mesma largura/fonte/margens do PDF final), abrindo
// página nova quando estoura. Cada página reusa o modelo .page/.frame/.foot: cabeçalho no topo,
// miolo, e rodapé preso ao PÉ por margin-top:auto. A assinatura entra só na última página, acima
// da NAP. Medir em caixa fora da página ou antes das webfonts carregarem dava medida errada e
// cortava o rodapé — por isso forçamos o load das fontes e medimos dentro da própria página.
func paginateDoc(page *rod.Page) error {
	_, err := page.Eval(`async () => {
		await Promise.all(Array.from(document.fonts).map(f => f.load().catch(() => {})));
		await document.fonts.ready;
		const mm = v => v * 96 / 25.4;
		const FRAME = mm(297 - 19 - 13);   // área útil do .frame (padding 19mm topo / 13mm pé)
		// Sem margem de segurança: footH já reserva o rodapé e a medição é IN SITU (= impressão).
		// Página cheia empacota até a borda (como o layout antigo do pedido de exames); página não
		// cheia ganha folga natural (rodapé com margin-top:auto).
		const SAFE = 0;
		const src = document.getElementById('src');
		const headHTML = document.getElementById('tpl-head').innerHTML;
		const napHTML  = document.getElementById('tpl-nap').innerHTML;
		const sigHTML  = document.getElementById('tpl-sig').innerHTML;
		const wmHTML   = document.getElementById('tpl-wm').innerHTML;
		const out = document.getElementById('out');
		const hasSig = sigHTML.trim().length > 0;

		// mede header / NAP / assinatura na largura real do miolo (170mm)
		const meas = document.createElement('div');
		meas.style.cssText = 'position:absolute;left:-9999px;top:0;width:170mm;';
		document.body.appendChild(meas);
		const measure = html => { meas.innerHTML = html; return meas.getBoundingClientRect().height; };
		const headH = measure(headHTML);
		const napH  = measure(napHTML);
		const sigH  = hasSig ? measure(sigHTML) + mm(8) : 0;
		const footH = napH + sigH;   // rodapé COMPLETO (assinatura + NAP) repetido em TODA página
		meas.remove();

		// blocos (NÓS) na ordem; cada filho de um .docbody é um bloco do miolo (paginável).
		const blocks = [];
		for (const el of Array.from(src.children)) {
			if (el.classList.contains('docbody')) {
				for (const c of Array.from(el.children)) blocks.push({body:true, el:c});
			} else {
				blocks.push({body:false, el});
			}
		}
		src.remove();

		// EMPACOTAMENTO IN SITU: adiciona blocos à página REAL e mede o overflow de verdade. O rodapé
		// (assinatura + NAP) é montado em TODA página, então a capacidade do miolo já desconta footH.
		// TOL: tolerância de transbordo. O .frame tem 13mm de padding inferior abaixo do rodapé, então
		// uns mm a mais só empurram o rodapé um tico pra dentro desse padding (continua visível) em vez
		// de quebrar um bloco grande pra outra página deixando a anterior vazia (caso do lab de 40 em
		// 2 colunas, que enche a folha "até a borda" — igual ao layout antigo).
		const avail = FRAME - headH - footH - SAFE;
		const TOL = mm(6);
		const mkPage = () => {
			const pg = document.createElement('div'); pg.className = 'page'; pg.innerHTML = wmHTML;
			const frame = document.createElement('div'); frame.className = 'frame';
			const head = document.createElement('div'); head.innerHTML = headHTML; frame.appendChild(head);
			const pb = document.createElement('div'); pb.className = 'pagebody'; frame.appendChild(pb);
			const foot = document.createElement('div'); foot.className = 'foot';
			foot.innerHTML = (hasSig ? sigHTML : '') + napHTML;   // assinatura em TODA página, acima da NAP
			frame.appendChild(foot);
			pg.appendChild(frame); out.appendChild(pg);
			return {pb, foot, doc: null};
		};
		const place = (b, pg) => {
			if (b.body) {
				if (!pg.doc) { pg.doc = document.createElement('div'); pg.doc.className = 'docbody'; pg.pb.appendChild(pg.doc); }
				pg.doc.appendChild(b.el);
			} else { pg.doc = null; pg.pb.appendChild(b.el); }
		};
		let cur = mkPage(); let n = 0; let brk = false;
		for (const b of blocks) {
			// .page-break: QUEBRA FORÇADA (ex.: separar grupos de exames lab × imagem). Abre página
			// nova antes do próximo bloco, sem criar página vazia se a atual ainda estiver em branco.
			if (b.el.classList && b.el.classList.contains('page-break')) { brk = true; continue; }
			if (brk && n > 0) { cur = mkPage(); n = 0; }
			brk = false;
			place(b, cur); n++;
			if (n > 1 && cur.pb.getBoundingClientRect().height > avail + TOL) {
				b.el.remove();                 // não coube: tira da página atual
				cur = mkPage(); n = 1;
				place(b, cur);                 // recoloca na página nova
			}
		}
		return out.children.length;
	}`)
	return err
}
