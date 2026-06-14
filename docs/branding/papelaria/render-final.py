# -*- coding: utf-8 -*-
"""Versão final — direção A (papel editorial creme).
Lista única na ordem registrada; linha em branco no texto = quebra de página.
Saída: A4 vetorial real (print-to-pdf) + PNG por página."""
import pathlib

HERE = pathlib.Path(__file__).parent.resolve()
ASSETS = HERE / "assets"
DRAFTS = HERE / "drafts"
DRAFTS.mkdir(exist_ok=True)

NALIETA = "/home/user/plenya/packages/brand/src/fonts/nalieta/Nalieta-Regular.otf"
INTER = "/home/user/plenya/docs/site/fontes_capa"
def a(name): return f"file://{ASSETS / name}"

# ---- dados de exemplo (NAP real) ----
PACIENTE = "Maria Helena Soares"
IDADE = "47 anos"
NASC = "12/03/1979"
CPF = "***.456.789-**"
DATA = "Londrina, 10 de junho de 2026"
MEDICO = "Dr. Getúlio José Mattos do Amaral Filho"
CRM = "CRM-PR 21.876 · RQE 16.038 · Nefrologia"
MEDICO_CPF = "***.123.456-**"   # CPF do signatário (vem do certificado ICP-Brasil)
INDICACAO = ("Avaliação metabólica e de longevidade, rastreio do pilar Gestão do Método AGIR. "
             "Paciente assintomática, em programa de acompanhamento contínuo.")

# Texto do pedido tal como armazenado: 1 exame por linha, ordem da definição.
# LINHA EM BRANCO = nova página (aqui: laboratório | imagem).
EXAMS = """Hemograma completo com plaquetas
Glicemia de jejum
Insulina de jejum
Hemoglobina glicada (HbA1c)
Colesterol total e frações (HDL, LDL)
Triglicérides
Apolipoproteína B
Lipoproteína (a)
TSH
T4 livre
T3 livre
Creatinina com taxa de filtração glomerular estimada
Ureia
Ácido úrico
Cistatina C
TGO (AST)
TGP (ALT)
Gama-GT
Fosfatase alcalina
Bilirrubinas total e frações
Proteínas totais e frações
Albumina
Vitamina D (25-hidroxivitamina D)
Vitamina B12
Ácido fólico
Ferritina
Ferro sérico e saturação de transferrina
Zinco sérico
Magnésio
Cálcio
Fósforo
Sódio
Potássio
Proteína C reativa ultrassensível
Homocisteína
VHS (velocidade de hemossedimentação)
Cortisol matinal
Testosterona total e livre
Microalbuminúria em amostra isolada
Urina tipo I (EAS)

Ultrassonografia de abdome total
Ecocardiograma transtorácico com Doppler
Densitometria óssea (coluna lombar e fêmur)
Tomografia computadorizada de tórax de baixa dose"""

FOOTER_NAP = ("Plenya Saúde · Av. Ayrton Senna da Silva, 500, Edifício Torre Pietra, sala 1402, "
              "Gleba Palhano, Londrina/PR · (43) 99974-8899 · plenyasaude.com.br")

# --- Dados da assinatura digital (carimbo ICP-Brasil, padrão PAdES) ---
SIGNED_AT  = "10/06/2026, 14:32 (horário de Brasília)"
DOC_ID     = "019eb4a2-7c10-7f3a-9c21-8d4e5b60a1b2"   # ID do documento = chave de validação por URL
DOC_REF    = "PLN-0610-7F3A"                            # referência curta legível (traceabilidade)
SHA256     = "7f3a 9c21 8d4e 5b60 a1b2 … d47"          # impressão SHA-256 do PDF assinado
VALID_URL  = "app.plenyasaude.com.br/lab-requests/validate"  # página real (resolve por ID)
ITI_URL    = "validar.iti.gov.br"             # validador oficial do governo (ITI/gov.br)
DIGITAL_SIGNATURE = True                       # False => modo manual (linha + carimbo, sem QR/ICP)

MAX_PER_PAGE = 40   # ≤20 → 1 coluna; 21–40 → 2 colunas (20/coluna); >40 → nova página

def pages_from_exams(text):
    """Quebra em páginas: por linha em branco E a cada 40 exames (2 col × 20)."""
    blocks, cur = [], []
    for line in text.split("\n"):
        if line.strip() == "":
            if cur: blocks.append(cur); cur = []
        else:
            cur.append(line.strip())
    if cur: blocks.append(cur)
    pages = []
    for b in blocks:
        for i in range(0, len(b), MAX_PER_PAGE):
            pages.append(b[i:i + MAX_PER_PAGE])
    return pages

PAGES = pages_from_exams(EXAMS)

# Título: Cormorant Garamond (livre, OFL) — substituto da Nalieta (paga), igual ao fallback do site.
CORM = str(ASSETS)
FONT_FACE = f"""
@font-face {{ font-family:'Cormorant'; src:url('file://{CORM}/cormorant-500.woff2'); font-weight:500; }}
@font-face {{ font-family:'Cormorant'; src:url('file://{CORM}/cormorant-600.woff2'); font-weight:600; }}
@font-face {{ font-family:'Inter'; src:url('file://{INTER}/Inter-Regular.ttf'); font-weight:400; }}
@font-face {{ font-family:'Inter'; src:url('file://{INTER}/Inter-Medium.ttf'); font-weight:500; }}
@font-face {{ font-family:'Inter'; src:url('file://{INTER}/Inter-SemiBold.ttf'); font-weight:600; }}
@font-face {{ font-family:'Inter'; src:url('file://{INTER}/Inter-Bold.ttf'); font-weight:700; }}
"""

CSS = f"""
{FONT_FACE}
/* Cores: exclusivamente tokens de marca. petrol=texto · gold=acentos · cream=papel/selo.
   Secundário = petrol com opacidade (mesmo matiz da marca, contraste AA garantido). */
:root {{
  --petrol:#063b4f; --gold:#b38645; --cream:#eae7da; --ocean:#417e8e; --sage:#92b8b4;
  --ink2:rgba(6,59,79,.80);   /* secundário */
  --ink3:rgba(6,59,79,.62);   /* legenda mínima */
}}
* {{ margin:0; padding:0; box-sizing:border-box; }}
@page {{ size:A4; margin:0; }}
html,body {{ background:#fff; }}
.page {{ position:relative; width:210mm; height:297mm; background:var(--cream);
         font-family:'Inter',sans-serif; color:var(--petrol); overflow:hidden;
         page-break-after:always; }}
.page:last-child {{ page-break-after:auto; }}
.frame {{ position:relative; z-index:1; height:100%; padding:19mm 20mm 13mm;
          display:flex; flex-direction:column; }}
/* marca d'água: padrão P∞ dourado da papelaria (papel timbrado), 4% */
.wm {{ position:absolute; inset:0; width:100%; height:100%; object-fit:cover;
       opacity:.04; z-index:0; }}
/* eyebrow/labels — petrol (contraste AA) */
.eyebrow {{ font-size:11px; letter-spacing:1.2px; text-transform:uppercase;
            color:var(--petrol); font-weight:600; }}
.sep {{ margin:0 9px; color:var(--gold); }}
/* cabeçalho */
.head {{ display:flex; align-items:center; justify-content:space-between; }}
.head .sym {{ height:29px; }}
.head .mark {{ height:23px; }}
.rule-gold {{ height:1.2px; background:var(--gold); margin-top:13px; }}
/* título — Cormorant Garamond */
.titlerow {{ display:flex; align-items:baseline; justify-content:space-between; margin-top:25px; }}
.title {{ font-family:'Cormorant',serif; font-size:34px; font-weight:600; letter-spacing:.2px;
          color:var(--petrol); line-height:1; }}
.title-rule {{ width:44px; height:2px; background:var(--gold); margin-top:12px; }}
/* paciente */
.patient {{ margin-top:20px; }}
.pname {{ font-size:17px; font-weight:600; color:var(--petrol); }}
.pmeta {{ font-size:13px; color:var(--ink2); margin-top:4px; }}
.pmeta .k {{ color:var(--petrol); text-transform:uppercase; letter-spacing:.6px;
             font-size:10.5px; font-weight:600; margin-right:5px; }}
/* indicação */
.indic {{ margin-top:16px; font-size:13.3px; line-height:1.55; color:var(--ink2); max-width:166mm; }}
.indic .eyebrow {{ display:block; margin-bottom:4px; }}
/* exames — itens 10pt, espaçamento dinâmico via --expad */
.sec {{ margin-top:18px; }}
.sec .eyebrow {{ display:block; margin-bottom:10px; }}
.exwrap {{ display:flex; gap:48px; }}
.excol {{ list-style:none; flex:1; }}
.exitem {{ display:flex; align-items:flex-start; gap:11px; padding:var(--expad,1.5px) 0; }}
.bullet {{ flex:0 0 auto; width:4.5px; height:4.5px; background:var(--gold); border-radius:1px;
           transform:rotate(45deg); margin-top:6px; }}
.exname {{ font-size:13.3px; line-height:1.35; color:var(--petrol); }}
/* rodapé (área) */
.foot {{ margin-top:auto; }}
/* selo de assinatura digital — compacto (linhas legais em 8pt) */
.seal {{ display:flex; align-items:center; gap:16px;
         background:rgba(6,59,79,.025); border:1px solid rgba(179,134,69,.40); border-radius:3px;
         padding:10px 15px; margin-bottom:12px; }}
.sealbadge {{ flex:0 0 auto; height:32px; display:block; padding:4px; border-radius:4px; overflow:hidden;
              background:var(--petrol); box-sizing:content-box; }}
.sealmain {{ flex:1; }}
.sealeyebrow {{ font-size:9px; letter-spacing:1.3px; text-transform:uppercase;
                color:var(--gold); font-weight:700; margin-bottom:3px; }}
.sealname {{ font-size:13px; font-weight:600; color:var(--petrol); }}
.sealreg {{ font-size:11px; color:var(--petrol); margin-top:1px; }}
.seallegal {{ font-size:10.6px; color:var(--ink2); line-height:1.45; margin-top:6px; }}
.sealvalid {{ font-size:10.6px; color:var(--ink2); line-height:1.45; margin-top:2px; }}
.seallegal b, .sealvalid b {{ color:var(--petrol); font-weight:700; }}
.sealqr {{ text-align:center; flex:0 0 auto; border-left:1px solid rgba(6,59,79,.14);
           padding-left:16px; }}
.sealqr img {{ width:60px; height:60px; display:block; }}
.sealqrcap {{ font-size:9px; color:var(--ink3); line-height:1.3; margin-top:3px; }}
/* assinatura manual — 10pt: data à esquerda, nome/dados centralizados */
.sigman {{ margin-bottom:13px; }}
.sigman .date {{ font-size:13.3px; color:var(--ink2); margin-bottom:30px; }}
.sigmancenter {{ text-align:center; }}
.sigmanline {{ display:inline-block; width:86mm; border-top:1px solid var(--petrol); padding-top:8px;
               font-size:14px; font-weight:600; color:var(--petrol); }}
.sigmansub {{ font-size:13.3px; color:var(--ink2); margin-top:2px; }}
/* rodapé NAP — 3 colunas estruturadas */
.nap {{ border-top:1px solid var(--gold); padding-top:11px; display:flex;
        justify-content:space-between; align-items:flex-start; gap:26px; }}
.napcol {{ font-size:11px; line-height:1.6; color:var(--ink2); flex:1; }}
.napcol.c {{ text-align:center; }}
.napcol.r {{ text-align:right; }}
.napcol b {{ display:block; color:var(--petrol); font-weight:700; font-size:11.5px; margin-bottom:1px; }}
"""

def header():
    return (f'<div class="head"><img class="sym" src="{a("symbol-petrol.svg")}">'
            f'<img class="mark" src="{a("lockup-tagline-petrol.svg")}"></div>'
            f'<div class="rule-gold"></div>')

def title_block():
    # Título do PDF é "Solicitação de Exames" (no EMR os menus/páginas seguem "Pedido de Exames").
    return ('<div class="titlerow"><div class="title">Solicitação de Exames</div></div>'
            '<div class="title-rule"></div>')

def patient_block():
    return (f'<div class="patient"><div class="pname">{PACIENTE}</div>'
            f'<div class="pmeta"><span class="k">Nascimento</span>{NASC} · {IDADE}'
            f'<span class="sep">·</span><span class="k">CPF</span>{CPF}</div></div>')

def signature():
    if not DIGITAL_SIGNATURE:
        # Modo manual: data + local à esquerda; nome e dados do médico centralizados.
        return (f'<div class="sigman">'
                f'<div class="date">{DATA}.</div>'
                f'<div class="sigmancenter">'
                f'<div class="sigmanline">{MEDICO}</div>'
                f'<div class="sigmansub">{CRM}</div>'
                f'</div></div>')
    # Modo digital: SELO compacto de autenticidade ICP-Brasil (PAdES).
    # Sem códigos impressos: o QR carrega o link de validação e o validador do ITI
    # confere por upload do próprio PDF — código/UUID/hash seriam redundantes.
    return (f'<div class="seal">'
            f'<img class="sealbadge" src="{a("icp-onbrand.svg")}">'
            f'<div class="sealmain">'
            f'<div class="sealeyebrow">Assinatura digital</div>'
            f'<div class="sealname">{MEDICO}</div>'
            f'<div class="sealreg">{CRM}</div>'
            f'<div class="seallegal">Assinado em {SIGNED_AT}. Padrão <b>PAdES</b>, certificado ICP-Brasil.</div>'
            f'<div class="sealvalid">Validade jurídica: <b>MP nº 2.200-2/2001</b> e <b>Lei nº 14.063/2020</b>. '
            f'Valide pelo QR ou em <b>{ITI_URL}</b>. Se impresso, autenticar na forma da lei.</div>'
            f'</div>'
            f'<div class="sealqr"><img src="{a("qr-sample.svg")}">'
            f'<div class="sealqrcap">Aponte a câmera<br>para validar</div></div>'
            f'</div>')

def footer_nap():
    # 3 colunas: identidade jurídica · endereço de atendimento · contato.
    return ('<div class="nap">'
            '<div class="napcol"><b>Plenya Saúde</b>Plenya Serviços de Saúde Ltda.<br>'
            'CNPJ 66.991.259/0001-50</div>'
            '<div class="napcol c"><b>Atendimento</b>Av. Ayrton Senna da Silva, 500<br>'
            'Torre Pietra, sala 1402 · Gleba Palhano · Londrina/PR</div>'
            '<div class="napcol r"><b>Contato</b>(43) 99974-8899<br>'
            'contato@plenyasaude.com.br · plenyasaude.com.br</div>'
            '</div>')

def exam_padding(items):
    # Espaçamento dinâmico: generoso com poucos exames, comprime até o mínimo com 20+
    # na coluna mais cheia (controla pela altura: min(n,20) linhas por coluna).
    MIN_PAD, MAX_PAD = 1.5, 7.0
    ctrl = min(len(items), 20)
    pad = MIN_PAD + (MAX_PAD - MIN_PAD) * (20 - ctrl) / 19.0
    return round(max(MIN_PAD, min(MAX_PAD, pad)), 2)

def exam_list(items):
    # ≤20: 1 coluna (largura quase total). 21–40: 2 colunas com 20 por coluna
    # (preenche a 1ª até 20, depois a 2ª) — preserva a ordem de leitura.
    def col(lst):
        lis = "".join(f'<li class="exitem"><span class="bullet"></span>'
                      f'<span class="exname">{name}</span></li>' for name in lst)
        return f'<ul class="excol">{lis}</ul>'
    pad = exam_padding(items)
    cols = col(items) if len(items) <= 20 else col(items[:20]) + col(items[20:])
    return f'<div class="exwrap" style="--expad:{pad}px">{cols}</div>'

def build():
    # Cada bloco (separado por linha em branco) = um pedido COMPLETO e independente.
    out = []
    # Indicação clínica é opcional: se vazia, nem renderiza (não ocupa espaço).
    indic_html = (f'<div class="indic"><span class="eyebrow">Indicação clínica</span>{INDICACAO}</div>'
                  if INDICACAO.strip() else '')
    for items in PAGES:
        top = (f'{header()}{title_block()}{patient_block()}'
               f'{indic_html}'
               f'<div class="sec"><span class="eyebrow">Exames solicitados</span>'
               f'{exam_list(items)}</div>')
        foot = f'<div class="foot">{signature()}{footer_nap()}</div>'
        out.append(f'<div class="page"><img class="wm" src="{a("pattern.svg")}">'
                   f'<div class="frame">{top}{foot}</div></div>')
    return (f"<!doctype html><html><head><meta charset='utf-8'><style>{CSS}</style></head>"
            f"<body>{''.join(out)}</body></html>")

(DRAFTS / "A-final.html").write_text(build(), encoding="utf-8")
print(f"wrote A-final.html · {len(PAGES)} páginas")
