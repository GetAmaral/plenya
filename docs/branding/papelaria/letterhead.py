# -*- coding: utf-8 -*-
"""Papel timbrado Plenya — bloco A4 pronto pra gráfica.

Gera papel timbrado genérico (cabeçalho + rodapé impressos, miolo liso pra
escrever à mão) em DUAS versões de fundo:
  - branco : miolo branco, só topo/rodapé em petrol/gold  (impressão barata)
  - creme  : folha inteira em creme #eae7da (full-bleed)   (mais sofisticado)

Especificação de impressão (gráfica):
  - Página PDF: 222 × 309 mm  (= A4 210×297 + 3 mm de sangria + 3 mm de marcas)
  - Sangria (bleed): 3 mm em cada lado  -> caixa sangrada 216 × 303 mm
  - Corte final (trim): 210 × 297 mm centralizado
  - Marcas de corte: 4 cantos, com folga de 2 mm da linha de corte
  - Fundo creme estende até a sangria; conteúdo dentro de margem de segurança

Saída: timbrado-branco.html / timbrado-creme.html em drafts/ (render via render-letterhead.js).
"""
import pathlib

HERE = pathlib.Path(__file__).parent.resolve()
ASSETS = HERE / "assets"
DRAFTS = HERE / "drafts"
DRAFTS.mkdir(exist_ok=True)
INTER = "/home/user/plenya/docs/site/fontes_capa"

def a(name): return f"file://{ASSETS / name}"

# ---- parâmetros do lockup do cabeçalho (ajustados por medição; ver tune-lockup) ----
WORD_H   = 27.0   # altura do wordmark PLENYA (px CSS) — aumento ~meio do caminho (cap ~20pt)
TAG_LS   = -0.95  # letter-spacing da tagline (px) — kerning p/ largura = PLENYA − 2px
TAG_WEIGHT = 400  # peso da tagline (Inter Regular — mais fina, menos bold)
TAG_GAP  = 7.0    # respiro entre wordmark e tagline (px)
SYM_H    = 44.0   # altura do símbolo P∞ = altura total do conjunto (medida e fixada)

# ---- NAP real (fonte: render-final.py / docs de marca) ----
NAP_JURIDICA = ("<b>Plenya Saúde</b>Plenya Serviços de Saúde Ltda.<br>"
                "CNPJ 66.991.259/0001-50")
NAP_ENDERECO = ("<b>Atendimento</b>Av. Ayrton Senna da Silva, 500<br>"
                "Torre Pietra, sala 1402<br>Gleba Palhano · Londrina/PR")
NAP_CONTATO  = ("<b>Contato</b>(43) 99974-8899<br>"
                "contato@plenyasaude.com.br<br>plenyasaude.com.br")

CLAIM = "Viva bem, viva mais."   # claim canônico da marca

CORM = str(ASSETS)
FONT_FACE = f"""
@font-face {{ font-family:'Inter'; src:url('file://{INTER}/Inter-Regular.ttf'); font-weight:400; }}
@font-face {{ font-family:'Inter'; src:url('file://{INTER}/Inter-Medium.ttf'); font-weight:500; }}
@font-face {{ font-family:'Inter'; src:url('file://{INTER}/Inter-SemiBold.ttf'); font-weight:600; }}
@font-face {{ font-family:'Inter'; src:url('file://{INTER}/Inter-Bold.ttf'); font-weight:700; }}
@font-face {{ font-family:'Cormorant'; src:url('file://{CORM}/cormorant-400.woff2'); font-weight:400; }}
@font-face {{ font-family:'Cormorant'; src:url('file://{CORM}/cormorant-500.woff2'); font-weight:500; }}
"""

def css(paper):
    """paper = '#fff' (versão branca) ou 'var(--cream)' (versão creme)."""
    return f"""
{FONT_FACE}
:root {{
  --petrol:#063b4f; --gold:#b38645; --cream:#eae7da; --ocean:#417e8e; --sage:#92b8b4;
  --ink2:rgba(6,59,79,.80); --ink3:rgba(6,59,79,.62);
}}
* {{ margin:0; padding:0; box-sizing:border-box; }}
/* página = caixa sangrada + anel de 3mm pras marcas de corte */
@page {{ size:222mm 309mm; margin:0; }}
html,body {{ background:#fff; }}
.sheet {{ position:relative; width:222mm; height:309mm; background:#fff;
          font-family:'Inter',sans-serif; color:var(--petrol); }}
/* caixa sangrada (216×303) — o fundo do papel vive aqui e vaza até a sangria */
.bleed {{ position:absolute; top:3mm; left:3mm; width:216mm; height:303mm; background:{paper}; }}
/* corte final (210×297) — toda a arte fica DENTRO desta caixa */
.trim {{ position:absolute; top:6mm; left:6mm; width:210mm; height:297mm; overflow:hidden; }}
.frame {{ position:absolute; inset:0; padding:18mm 20mm 14mm; display:flex; flex-direction:column; }}

/* marca d'água — modo MARCA: símbolo P∞ centralizado, petrol bem suave */
.wm {{ position:absolute; top:50%; left:50%; width:115mm; transform:translate(-50%,-50%);
       opacity:.045; }}
/* marca d'água — modo ESTAMPADO: padrão P∞ cobrindo a folha (igual aos PDFs digitais) */
.estampa {{ position:absolute; inset:0; width:100%; height:100%; object-fit:cover; opacity:.04; }}
/* claim 'Viva bem, viva mais.' — modo MARCA: à direita, abaixo do P∞ (10%) */
.claim-r {{ position:absolute; right:47.8mm; top:calc(50% + 47mm + 15px);
   font-family:'Cormorant',serif; font-weight:500; font-size:25px; letter-spacing:2px;
   color:var(--petrol); opacity:.10; white-space:nowrap; }}
/* claim — modo ESTAMPADO: vertical na lateral esquerda (25%) */
.claim-v {{ position:absolute; left:9mm; top:50%; transform-origin:left center;
   transform:rotate(-90deg) translateX(-50%);
   font-family:'Cormorant',serif; font-weight:500; font-size:19.2px; letter-spacing:3.6px;
   color:var(--petrol); opacity:.25; white-space:nowrap; }}

/* cabeçalho — símbolo P∞ (esq) com a MESMA altura do conjunto do lockup (dir) */
.head {{ display:flex; align-items:center; justify-content:space-between; }}
.head .sym {{ height:{SYM_H}px; }}
/* lockup: wordmark PLENYA (SVG vetorial) + tagline (TEXTO vetorial, 8pt) empilhados */
.lk {{ display:flex; flex-direction:column; align-items:center; }}
.lk .word {{ height:{WORD_H}px; display:block; }}
.lk .tag {{ font-family:'Inter',sans-serif; font-weight:{TAG_WEIGHT}; font-size:8pt;
            text-transform:uppercase; letter-spacing:{TAG_LS}px; color:var(--petrol);
            line-height:1; margin-top:{TAG_GAP}px; white-space:nowrap; }}
.rule-gold {{ height:1.2px; background:var(--gold); margin-top:13px; }}

/* miolo liso (área de escrever) — só ocupa o espaço */
.body {{ flex:1; }}

/* rodapé NAP — 3 colunas, regra dourada no topo */
.foot {{ margin-top:auto; }}
.nap {{ border-top:1px solid var(--gold); padding-top:11px; display:flex;
        justify-content:space-between; align-items:flex-start; gap:26px; }}
.napcol {{ font-size:9pt; line-height:1.6; color:var(--ink2); flex:1; }}
.napcol.c {{ text-align:center; }}
.napcol.r {{ text-align:right; }}
.napcol b {{ display:block; color:var(--petrol); font-weight:700; font-size:9.5pt; margin-bottom:1px; }}

/* marcas de corte — 4 cantos do trim (6mm), com folga de 2mm da linha de corte */
.cm {{ position:absolute; background:#000; }}
.cm.v {{ width:.3mm; height:4mm; }}
.cm.h {{ height:.3mm; width:4mm; }}
/* topo-esquerda */
.cm.tl-v {{ left:6mm; top:0; }}
.cm.tl-h {{ top:6mm; left:0; }}
/* topo-direita */
.cm.tr-v {{ right:6mm; top:0; }}
.cm.tr-h {{ top:6mm; right:0; }}
/* base-esquerda */
.cm.bl-v {{ left:6mm; bottom:0; }}
.cm.bl-h {{ bottom:6mm; left:0; }}
/* base-direita */
.cm.br-v {{ right:6mm; bottom:0; }}
.cm.br-h {{ bottom:6mm; right:0; }}
"""

def crop_marks():
    return "".join(f'<div class="cm {c}"></div>' for c in [
        "v tl-v", "h tl-h", "v tr-v", "h tr-h",
        "v bl-v", "h bl-h", "v br-v", "h br-h"])

def header():
    return (f'<div class="head"><img class="sym" src="{a("symbol-petrol.svg")}">'
            f'<div class="lk"><img class="word" src="{a("wordmark-petrol.svg")}">'
            f'<div class="tag">Saúde, Performance &amp; Longevidade</div></div></div>'
            f'<div class="rule-gold"></div>')

def footer_nap():
    return (f'<div class="nap">'
            f'<div class="napcol">{NAP_JURIDICA}</div>'
            f'<div class="napcol c">{NAP_ENDERECO}</div>'
            f'<div class="napcol r">{NAP_CONTATO}</div>'
            f'</div>')

def build(paper, wm_mode):
    """wm_mode = 'marca' (símbolo P∞ central) ou 'estampado' (padrão cobrindo a folha)."""
    # estampado: padrão na caixa sangrada (vaza até a borda). marca: símbolo central no frame.
    bleed_inner = (f'<img class="estampa" src="{a("pattern.svg")}">'
                   if wm_mode == "estampado" else "")
    wm = (f'<img class="wm" src="{a("symbol-petrol.svg")}">'
          if wm_mode == "marca" else "")
    # claim: marca -> à direita abaixo do P∞ ; estampado -> vertical na lateral esquerda
    claim = (f'<div class="claim-r">{CLAIM}</div>' if wm_mode == "marca"
             else f'<div class="claim-v">{CLAIM}</div>')
    inner = (f'{wm}{claim}<div class="frame">{header()}'
             f'<div class="body"></div>'
             f'<div class="foot">{footer_nap()}</div></div>')
    sheet = (f'<div class="sheet"><div class="bleed">{bleed_inner}</div>'
             f'<div class="trim">{inner}</div>{crop_marks()}</div>')
    return (f"<!doctype html><html><head><meta charset='utf-8'>"
            f"<style>{css(paper)}</style></head><body>{sheet}</body></html>")

VARIANTS = [
    ("timbrado-branco-marca",      "#fff",          "marca"),
    ("timbrado-branco-estampado",  "#fff",          "estampado"),
    ("timbrado-creme-marca",       "var(--cream)",  "marca"),
    ("timbrado-creme-estampado",   "var(--cream)",  "estampado"),
]
for slug, paper, wm_mode in VARIANTS:
    (DRAFTS / f"{slug}.html").write_text(build(paper, wm_mode), encoding="utf-8")
    print("wrote", slug)
