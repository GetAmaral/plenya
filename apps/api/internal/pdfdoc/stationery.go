package pdfdoc

import (
	"html"
	"strings"
)

// ---- Tipos de dados compartilhados por todos os documentos ----

// Clinic — identidade da clínica para o rodapé (NAP). Endereço de atendimento por padrão;
// o recibo usa o endereço fiscal.
type Clinic struct {
	LegalName   string // "Plenya Serviços de Saúde Ltda."
	CNPJ        string // "66.991.259/0001-50"
	AddressLine string // "Av. Ayrton Senna da Silva, 500"
	AddressCont string // "Torre Pietra, sala 1402 · Gleba Palhano · Londrina/PR"
	Phone       string // "(43) 99974-8899"
	Email       string // "contato@plenyasaude.com.br"
	Site        string // "plenyasaude.com.br"
}

// DefaultClinic — NAP canônica da Plenya (fonte: docs/seo/01-NAP-master.md).
func DefaultClinic() Clinic {
	return Clinic{
		LegalName:   "Plenya Serviços de Saúde Ltda.",
		CNPJ:        "66.991.259/0001-50",
		AddressLine: "Av. Ayrton Senna da Silva, 500",
		AddressCont: "Torre Pietra, sala 1402 · Gleba Palhano · Londrina/PR",
		Phone:       "(43) 99974-8899",
		Email:       "contato@plenyasaude.com.br",
		Site:        "plenyasaude.com.br",
	}
}

// Patient — identificação do paciente no documento.
type Patient struct {
	Name      string
	BirthInfo string // "12/03/1979 · 47 anos"
	CPFMasked string // "***.456.789-**"
}

// Doctor — identificação do médico (signatário). CPF do médico NÃO entra por padrão
// (LGPD); só em receituários de controle (RDC 1.000/2025) — ver plano-papelaria-pdf.md §5.
type Doctor struct {
	Name        string // "Dr. Getúlio José Mattos do Amaral Filho"
	Credentials string // "CRM-PR 21.876 · RQE 16.038 · Nefrologia"
}

// Signature — estado da assinatura. Digital => selo ICP-Brasil + QR; senão => carimbo manual.
type Signature struct {
	Digital     bool
	SignedAt    string // "10/06/2026, 14:32 (horário de Brasília)"
	ValidateURL string // alvo do QR (URL completa da página de validação por ID)
	ValidateRef string // exibido em texto (ex.: "app.plenyasaude.com.br/lab-requests/validate")
	PlaceDate   string // modo manual: "Londrina, 10 de junho de 2026"
}

func esc(s string) string { return html.EscapeString(s) }

// ---- CSS do sistema-base (direção A, tokens de marca). fontFaces() é prefixado em runtime. ----
const baseCSS = `
:root{
  --petrol:#063b4f; --gold:#b38645; --cream:#eae7da; --ocean:#417e8e; --sage:#92b8b4;
  --ink2:rgba(6,59,79,.80); --ink3:rgba(6,59,79,.62);
}
*{ margin:0; padding:0; box-sizing:border-box; }
@page{ size:A4; margin:0; }
html,body{ background:#fff; }
.page{ position:relative; width:210mm; height:297mm; background:var(--cream);
  font-family:'Inter',sans-serif; color:var(--petrol); overflow:hidden; page-break-after:always; }
.page:last-child{ page-break-after:auto; }
.frame{ position:relative; z-index:1; height:100%; padding:19mm 20mm 13mm; display:flex; flex-direction:column; }
.wm{ position:absolute; inset:0; width:100%; height:100%; object-fit:cover; opacity:.04; z-index:0; }
.eyebrow{ font-size:11px; letter-spacing:1.2px; text-transform:uppercase; color:var(--petrol); font-weight:600; }
.sep{ margin:0 9px; color:var(--gold); }
.head{ display:flex; align-items:center; justify-content:space-between; }
.head .sym{ height:44px; }
/* lockup: wordmark PLENYA (SVG) + tagline (texto vetorial Inter, 8pt, centralizada) */
.lk{ display:flex; flex-direction:column; align-items:center; }
.lk .word{ height:27px; display:block; }
.lk .tag{ font-family:'Inter',sans-serif; font-weight:400; font-size:8pt; text-transform:uppercase;
  letter-spacing:-0.95px; color:var(--petrol); line-height:1; margin-top:7px; white-space:nowrap; }
.rule-gold{ height:1.2px; background:var(--gold); margin-top:13px; }
/* claim 'Viva bem, viva mais.' — marca d'água vertical na lateral esquerda */
.claim-v{ position:absolute; left:9mm; top:50%; transform-origin:left center;
  transform:rotate(-90deg) translateX(-50%); font-family:'Cormorant',serif; font-weight:500;
  font-size:19.2px; letter-spacing:3.6px; color:var(--petrol); opacity:.25; white-space:nowrap; z-index:0; }
.titleblock{ margin-top:25px; }
.title-kind{ font-size:11px; letter-spacing:1.5px; text-transform:uppercase; color:var(--gold); font-weight:700; margin-bottom:4px; }
.titlerow{ display:flex; align-items:baseline; justify-content:space-between; }
.title{ font-family:'Cormorant',serif; font-size:34px; font-weight:600; letter-spacing:.2px; color:var(--petrol); line-height:1.1; }
.title-rule{ width:44px; height:2px; background:var(--gold); margin-top:12px; }
.patient{ margin-top:20px; }
.pname{ font-size:17px; font-weight:600; color:var(--petrol); }
.pmeta{ font-size:13px; color:var(--ink2); margin-top:4px; }
.pmeta .k{ color:var(--petrol); text-transform:uppercase; letter-spacing:.6px; font-size:10.5px; font-weight:600; margin-right:5px; }
.indic{ margin-top:16px; font-size:13.3px; line-height:1.55; color:var(--ink2); max-width:166mm; }
.indic .eyebrow{ display:block; margin-bottom:4px; }
.sec{ margin-top:18px; }
.sec .eyebrow{ display:block; margin-bottom:10px; }
.exwrap{ display:flex; gap:48px; }
.excol{ list-style:none; flex:1; }
.exitem{ display:flex; align-items:flex-start; gap:11px; padding:var(--expad,1.5px) 0; }
.bullet{ flex:0 0 auto; width:4.5px; height:4.5px; background:var(--gold); border-radius:1px; transform:rotate(45deg); margin-top:6px; }
.exbody{ display:flex; flex-direction:column; min-width:0; }
.exname{ font-size:13.3px; line-height:1.35; color:var(--petrol); }
.extuss{ color:var(--ink2); font-size:10.5px; white-space:nowrap; }
.exjust{ font-size:10.8px; line-height:1.3; color:var(--ink2); font-style:italic; margin-top:2px; padding-left:11px; border-left:2px solid rgba(179,134,69,.35); }
.ctrltag{ display:inline-block; margin-top:9px; font-size:10px; letter-spacing:1px; text-transform:uppercase; color:var(--gold); font-weight:700; border:1px solid rgba(179,134,69,.5); border-radius:3px; padding:3px 9px; }
.meds{ margin-top:2px; }
.med{ padding:8px 0; border-top:1px solid rgba(6,59,79,.08); }
.med:first-child{ border-top:none; padding-top:2px; }
.medhead{ font-size:13.5px; font-weight:600; color:var(--petrol); }
.medconc{ color:var(--ink2); font-weight:600; }
.medsub{ font-size:11.5px; color:var(--ink2); margin-top:1px; }
.medpos{ font-size:12.5px; color:var(--petrol); margin-top:4px; }
.medqty{ font-size:11.5px; color:var(--ink2); margin-top:2px; }
.medinstr{ font-size:11.5px; color:var(--ink2); margin-top:2px; font-style:italic; }
/* ---- Receituário magistral (manipulado). Cada .formula é um bloco atômico de paginação, por
   isso o número de componentes é limitado na validação: um bloco maior que a página não é
   quebrado pelo paginador, ele transborda em silêncio. */
.formula{ padding:10px 0 4px; border-top:1px solid rgba(6,59,79,.08); }
.formula:first-child{ border-top:none; padding-top:2px; }
.fhead{ display:flex; align-items:baseline; justify-content:space-between; gap:12px; }
.fname{ font-size:13.5px; font-weight:600; color:var(--petrol); }
.fform{ color:var(--ink2); font-weight:600; }
.fuse{ flex:0 0 auto; font-size:9.5px; letter-spacing:1px; text-transform:uppercase; font-weight:700; color:var(--gold); border:1px solid rgba(179,134,69,.5); border-radius:3px; padding:2px 7px; white-space:nowrap; }
.comp{ display:flex; align-items:baseline; gap:6px; margin-top:5px; font-size:12.8px; color:var(--petrol); }
.compname{ flex:0 1 auto; }
.compnote{ color:var(--ink2); font-size:10.8px; font-style:italic; margin:1px 0 0 14px; }
/* "(do elemento)" vai colado na quantidade e em peso, porque muda o que a farmácia pesa. */
.compelem{ font-size:10.5px; font-weight:600; }
.dots{ flex:1 1 auto; border-bottom:1px dotted rgba(6,59,79,.35); transform:translateY(-3px); min-width:14px; }
.compqty{ flex:0 0 auto; font-weight:600; white-space:nowrap; }
.fveh{ display:flex; align-items:baseline; gap:6px; margin-top:5px; font-size:12.8px; color:var(--ink2); }
.fdisp{ margin-top:8px; font-size:12.8px; color:var(--petrol); font-weight:600; }
.fpos{ margin-top:4px; font-size:12.5px; color:var(--petrol); }
.finstr{ margin-top:2px; font-size:11.5px; color:var(--ink2); font-style:italic; }
.validity{ margin-top:14px; font-size:12px; color:var(--ink2); }
.validity b{ color:var(--petrol); font-weight:700; }
.docbody{ margin-top:20px; font-size:13.5px; line-height:1.75; color:var(--petrol); max-width:170mm; text-align:justify; }
.docbody p{ margin-bottom:11px; }
.docbody h1,.docbody h2,.docbody h3{ font-family:'Cormorant',serif; color:var(--petrol); line-height:1.2; margin:14px 0 6px; font-weight:600; }
.docbody h1{ font-size:22px; }
.docbody h2{ font-size:18px; }
.docbody h3{ font-size:15px; }
.docbody ol{ margin:0 0 11px 22px; }
.docbody ol li{ margin-bottom:4px; }
/* Lista com marcador losango dourado — mesmo bullet do pedido de exames (.bullet). */
.docbody ul{ list-style:none; margin:0 0 11px 0; padding-left:0; }
.docbody ul ul{ margin:4px 0 4px 18px; }
.docbody ul li{ position:relative; padding-left:18px; margin-bottom:5px; }
.docbody ul li::before{ content:""; position:absolute; left:3px; top:0.62em; width:4.5px; height:4.5px; background:var(--gold); border-radius:1px; transform:rotate(45deg); }
.docbody strong,.docbody b{ font-weight:700; }
.docbody u{ text-decoration:underline; }
.docbody s,.docbody del{ text-decoration:line-through; }
.docbody mark{ padding:0 2px; }
.doccid{ margin-top:14px; font-size:12px; color:var(--ink2); }
.rcpt-head{ text-align:right; }
.rcpt-eyebrow{ font-size:10px; letter-spacing:1px; text-transform:uppercase; color:var(--gold); font-weight:600; }
.rcpt-val{ font-size:18px; font-weight:700; color:var(--petrol); margin-top:3px; }
.rcpt-body{ margin-top:26px; font-size:13.5px; line-height:1.8; color:var(--petrol); max-width:170mm; }
.rcpt-body b{ font-weight:700; }
.rcpt-notes{ margin-top:8px; color:var(--ink2); }
.rcpt-refund{ margin-top:14px; font-size:11px; letter-spacing:.5px; text-transform:uppercase; color:var(--petrol); font-weight:700; border:1px solid var(--gold); border-radius:3px; padding:6px 11px; display:inline-block; }
.rcpt-sign{ padding-top:30px; text-align:center; margin-bottom:14px; }
.rcpt-date{ font-size:13px; color:var(--ink2); margin-bottom:30px; text-align:left; }
.rcpt-signline{ display:inline-block; min-width:80mm; border-top:1px solid var(--petrol); padding-top:8px; font-size:13px; font-weight:600; color:var(--petrol); }
.rcpt-signsub{ font-size:11px; font-weight:500; color:var(--ink2); }
.foot{ margin-top:auto; }
.seal{ display:flex; align-items:center; gap:16px; background:rgba(6,59,79,.025);
  border:1px solid rgba(179,134,69,.40); border-radius:3px; padding:10px 15px; margin-bottom:12px; }
.sealbadge{ flex:0 0 auto; height:32px; display:block; padding:4px; border-radius:4px; overflow:hidden; background:var(--petrol); box-sizing:content-box; }
.sealmain{ flex:1; }
.sealeyebrow{ font-size:9px; letter-spacing:1.3px; text-transform:uppercase; color:var(--gold); font-weight:700; margin-bottom:3px; }
.sealname{ font-size:13px; font-weight:600; color:var(--petrol); }
.sealreg{ font-size:11px; color:var(--petrol); margin-top:1px; }
.seallegal{ font-size:10.6px; color:var(--ink2); line-height:1.45; margin-top:6px; }
.sealvalid{ font-size:10.6px; color:var(--ink2); line-height:1.45; margin-top:2px; }
.seallegal b, .sealvalid b{ color:var(--petrol); font-weight:700; }
.sealqr{ text-align:center; flex:0 0 auto; border-left:1px solid rgba(6,59,79,.14); padding-left:16px; }
.sealqr .qrbox{ width:60px; height:60px; line-height:0; }
.sealqr .qrbox svg{ width:60px; height:60px; display:block; }
.sealqrcap{ font-size:9px; color:var(--ink3); line-height:1.3; margin-top:3px; }
.sigman{ margin-bottom:13px; }
.sigman .date{ font-size:13.3px; color:var(--ink2); margin-bottom:30px; }
.sigmancenter{ text-align:center; }
.sigmanline{ display:inline-block; width:86mm; border-top:1px solid var(--petrol); padding-top:8px; font-size:14px; font-weight:600; color:var(--petrol); }
.sigmansub{ font-size:13.3px; color:var(--ink2); margin-top:2px; }
.nap{ border-top:1px solid var(--gold); padding-top:11px; display:flex; justify-content:space-between; align-items:flex-start; gap:26px; }
.napcol{ font-size:9pt; line-height:1.6; color:var(--ink2); flex:1; }
.napcol.c{ text-align:center; }
.napcol.r{ text-align:right; }
.napcol b{ display:block; color:var(--petrol); font-weight:700; font-size:9.5pt; margin-bottom:1px; }

/* ---- Documento FLUÍDO (atestado/declaração/laudo/orientações): PAGINAÇÃO REAL em caixas .page.
   O paginador (paginateIssuedDoc, em Go/JS) mede cada bloco e os distribui em páginas A4 fixas,
   reusando o modelo .page/.frame/.foot dos demais PDFs: cabeçalho no topo, miolo, e rodapé preso
   ao PÉ por margin-top:auto. Como cada página é uma caixa de altura fixa e o conteúdo é empacotado
   pra caber, o rodapé fica SEMPRE colado no pé (inclusive na última página, mesmo curta).
   .src é a fonte medida fora da tela; .pagebody é o miolo de cada página. */
.src{ position:absolute; left:-9999px; top:0; width:170mm; }
.docbody{ max-width:none; }
.indic{ max-width:none; }
.pagebody .docbody{ margin-top:14px; }
.foot .sig-block{ margin-bottom:13px; }
`

// ---- Blocos compartilhados ----

// claimText — claim canônico da marca, usado como marca d'água nos documentos.
const claimText = "Viva bem, viva mais."

func headerHTML() string {
	return `<div class="head">` + imgSVG("symbol-petrol.svg", "sym") +
		`<div class="lk">` + imgSVG("wordmark-petrol.svg", "word") +
		`<div class="tag">Saúde, Performance &amp; Longevidade</div></div>` +
		`</div><div class="rule-gold"></div>`
}

func patientHTML(p Patient) string {
	var meta []string
	if p.BirthInfo != "" {
		meta = append(meta, `<span class="k">Nascimento</span>`+esc(p.BirthInfo))
	}
	if p.CPFMasked != "" {
		meta = append(meta, `<span class="k">CPF</span>`+esc(p.CPFMasked))
	}
	metaHTML := ""
	if len(meta) > 0 {
		metaHTML = `<div class="pmeta">` + strings.Join(meta, `<span class="sep">·</span>`) + `</div>`
	}
	return `<div class="patient"><div class="pname">` + esc(p.Name) + `</div>` + metaHTML + `</div>`
}

func indicationHTML(s string) string {
	if s == "" {
		return ""
	}
	return `<div class="indic"><span class="eyebrow">Indicação clínica</span>` + esc(s) + `</div>`
}

func footerNAPHTML(c Clinic) string {
	// Endereço em até 3 linhas (1ª " · " vira quebra) e contato com e-mail/site em linhas
	// separadas — evita "·" órfão e quebra no meio de "Gleba Palhano" (igual ao papel timbrado).
	addrCont := strings.Replace(esc(c.AddressCont), " · ", "<br>", 1)
	return `<div class="nap">` +
		`<div class="napcol"><b>Plenya Saúde</b>` + esc(c.LegalName) + `<br>CNPJ ` + esc(c.CNPJ) + `</div>` +
		`<div class="napcol c"><b>Atendimento</b>` + esc(c.AddressLine) + `<br>` + addrCont + `</div>` +
		`<div class="napcol r"><b>Contato</b>` + esc(c.Phone) + `<br>` + esc(c.Email) + `<br>` + esc(c.Site) + `</div>` +
		`</div>`
}

// signatureBlock — assinatura embrulhada no .sig-block (espaçamento padrão acima da NAP).
func signatureBlock(d Doctor, s Signature) string {
	return `<div class="sig-block">` + signatureHTML(d, s) + `</div>`
}

// signatureHTML — selo ICP-Brasil (digital) ou carimbo manual.
func signatureHTML(d Doctor, s Signature) string {
	if !s.Digital {
		return `<div class="sigman"><div class="date">` + esc(s.PlaceDate) + `.</div>` +
			`<div class="sigmancenter"><div class="sigmanline">` + esc(d.Name) + `</div>` +
			`<div class="sigmansub">` + esc(d.Credentials) + `</div></div></div>`
	}
	qr := qrSVG(s.ValidateURL, "#063b4f")
	return `<div class="seal">` + imgSVG("icp-onbrand.svg", "sealbadge") +
		`<div class="sealmain">` +
		`<div class="sealeyebrow">Assinatura digital</div>` +
		`<div class="sealname">` + esc(d.Name) + `</div>` +
		`<div class="sealreg">` + esc(d.Credentials) + `</div>` +
		`<div class="seallegal">Assinado em ` + esc(s.SignedAt) + `. Padrão <b>PAdES</b>, certificado ICP-Brasil.</div>` +
		`<div class="sealvalid">Validade jurídica: <b>MP nº 2.200-2/2001</b> e <b>Lei nº 14.063/2020</b>. ` +
		`Valide pelo QR ou em <b>validar.iti.gov.br</b>. Se impresso, autenticar na forma da lei.</div>` +
		`</div>` +
		`<div class="sealqr"><div class="qrbox">` + qr + `</div>` +
		`<div class="sealqrcap">Aponte a câmera<br>para validar</div></div>` +
		`</div>`
}

// documentHTMLCSS embrulha o conteúdo no HTML final com fontes + CSS base + CSS específico do
// documento (extra). documentHTML é o atalho sem CSS extra.
func documentHTMLCSS(content, extra string) string {
	return `<!doctype html><html><head><meta charset="utf-8"><style>` +
		fontFaces() + baseCSS + extra + `</style></head><body>` + content + `</body></html>`
}

func documentHTML(content string) string { return documentHTMLCSS(content, "") }
