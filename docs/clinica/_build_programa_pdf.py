# -*- coding: utf-8 -*-
"""Gera o Programa de Necessidades (arquiteto) em PDF A4 branded Plenya.
Pág.1 = espaços e tamanhos; págs. seguintes = explicações. chromium print-to-pdf."""
import pathlib, subprocess, sys

ROOT = pathlib.Path("/home/user/plenya")
ASSETS = ROOT / "docs/branding/papelaria/assets"
INTER = ROOT / "docs/site/fontes_capa"
OUT_HTML = ROOT / "docs/clinica/_programa_plenya.html"
OUT_PDF = ROOT / "docs/clinica/programa-arquitetonico-plenya.pdf"

def f(p): return f"file://{p}"

FONT_FACE = f"""
@font-face {{ font-family:'Cormorant'; src:url('{f(ASSETS)}/cormorant-500.woff2'); font-weight:500; }}
@font-face {{ font-family:'Cormorant'; src:url('{f(ASSETS)}/cormorant-600.woff2'); font-weight:600; }}
@font-face {{ font-family:'Cormorant'; src:url('{f(ASSETS)}/cormorant-700.woff2'); font-weight:700; }}
@font-face {{ font-family:'Inter'; src:url('{f(INTER)}/Inter-Regular.ttf'); font-weight:400; }}
@font-face {{ font-family:'Inter'; src:url('{f(INTER)}/Inter-Medium.ttf'); font-weight:500; }}
@font-face {{ font-family:'Inter'; src:url('{f(INTER)}/Inter-SemiBold.ttf'); font-weight:600; }}
@font-face {{ font-family:'Inter'; src:url('{f(INTER)}/Inter-Bold.ttf'); font-weight:700; }}
"""

CSS = """
:root{ --petrol:#063b4f; --gold:#b38645; --cream:#eae7da; --ocean:#417e8e; --sage:#92b8b4;
  --ink2:rgba(6,59,79,.82); --ink3:rgba(6,59,79,.60); --line:rgba(6,59,79,.14); }
*{ margin:0; padding:0; box-sizing:border-box; }
@page{ size:A4; margin:0; }
html,body{ background:#fff; }
.page{ position:relative; width:210mm; height:297mm; background:var(--cream);
  font-family:'Inter',sans-serif; color:var(--petrol); overflow:hidden; page-break-after:always; }
.page:last-child{ page-break-after:auto; }
.wm{ position:absolute; inset:0; width:100%; height:100%; object-fit:cover; opacity:.035; z-index:0; }
.frame{ position:relative; z-index:1; height:100%; padding:16mm 17mm 12mm; display:flex; flex-direction:column; }
/* header */
.head{ display:flex; align-items:center; justify-content:space-between; }
.head .sym{ height:26px; }
.head .mark{ height:20px; }
.eyebrow{ font-size:9.5px; letter-spacing:1.4px; text-transform:uppercase; color:var(--gold); font-weight:600; }
.rule-gold{ height:1.2px; background:var(--gold); margin-top:10px; }
.titlerow{ margin-top:16px; }
.title{ font-family:'Cormorant',serif; font-size:31px; font-weight:600; color:var(--petrol); line-height:1; }
.subtitle{ font-size:11.5px; color:var(--ink2); margin-top:7px; line-height:1.5; max-width:172mm; }
/* zone columns */
.cols{ display:flex; gap:9mm; margin-top:13px; }
.col{ flex:1; }
.zone{ margin-bottom:11px; break-inside:avoid; }
.ztitle{ font-family:'Cormorant',serif; font-size:15px; font-weight:700; color:var(--petrol);
  display:flex; align-items:baseline; gap:7px; }
.ztitle .zn{ font-size:9.5px; font-family:'Inter'; font-weight:600; color:var(--gold);
  border:1px solid var(--gold); border-radius:3px; padding:1px 5px; letter-spacing:.5px; }
.ztitle .ztot{ margin-left:auto; font-family:'Inter'; font-size:11px; font-weight:700; color:var(--ocean); }
.zsub{ font-size:9px; color:var(--ink3); margin:2px 0 5px; letter-spacing:.3px; text-transform:uppercase; }
.row{ display:flex; justify-content:space-between; gap:8px; font-size:10.7px; line-height:1.5;
  padding:1.5px 0; border-bottom:1px solid var(--line); }
.row .m{ color:var(--petrol); font-weight:600; white-space:nowrap; }
.row .n{ color:var(--ink2); }
/* totals band */
.totals{ margin-top:auto; }
.tband{ display:flex; gap:7mm; }
.tcard{ flex:1; background:rgba(6,59,79,.05); border:1px solid var(--line); border-left:3px solid var(--gold);
  border-radius:5px; padding:10px 12px; }
.tcard .tl{ font-size:9.5px; letter-spacing:.5px; text-transform:uppercase; color:var(--ink3); font-weight:600; }
.tcard .tv{ font-family:'Cormorant',serif; font-size:27px; font-weight:700; color:var(--petrol); line-height:1.05; margin-top:3px; }
.tcard .td{ font-size:9.5px; color:var(--ink2); margin-top:3px; line-height:1.35; }
/* footer */
.foot{ display:flex; align-items:center; justify-content:space-between; margin-top:11px;
  padding-top:8px; border-top:1px solid var(--line); font-size:8.8px; color:var(--ink3); letter-spacing:.4px; }
.foot .tag{ text-transform:uppercase; font-weight:600; color:var(--gold); }
/* explanation pages */
.lead{ font-size:12px; color:var(--ink2); line-height:1.6; margin-top:13px; max-width:174mm; }
.callout{ background:rgba(6,59,79,.05); border:1px solid var(--line); border-left:3px solid var(--gold);
  border-radius:5px; padding:12px 14px; margin-top:14px; }
.callout .ch{ font-family:'Cormorant',serif; font-size:15px; font-weight:700; color:var(--petrol); margin-bottom:5px; }
.callout p{ font-size:11px; color:var(--ink2); line-height:1.55; }
.callout b{ color:var(--petrol); }
.block{ margin-top:15px; break-inside:avoid; }
.bh{ font-family:'Cormorant',serif; font-size:16px; font-weight:700; color:var(--petrol);
  display:flex; align-items:baseline; gap:7px; }
.bh .zn{ font-size:9px; font-family:'Inter'; font-weight:600; color:var(--gold);
  border:1px solid var(--gold); border-radius:3px; padding:1px 5px; }
.bp{ font-size:11px; color:var(--ink2); line-height:1.6; margin-top:5px; }
.bp b{ color:var(--petrol); font-weight:600; }
ul.steps{ margin-top:6px; padding-left:0; list-style:none; }
ul.steps li{ font-size:11px; color:var(--ink2); line-height:1.5; padding:3px 0 3px 20px; position:relative; }
ul.steps li:before{ content:""; position:absolute; left:3px; top:8px; width:5px; height:5px;
  background:var(--gold); border-radius:1px; transform:rotate(45deg); }
ul.steps li b{ color:var(--petrol); }
.decisions li{ counter-increment:d; padding-left:24px; }
.decisions li:before{ content:counter(d); background:var(--petrol); color:var(--cream);
  width:15px; height:15px; border-radius:50%; font-size:9px; font-weight:700; display:flex;
  align-items:center; justify-content:center; transform:none; top:4px; left:2px; }
.decisions{ counter-reset:d; }
"""

def page(inner):
    return f'<div class="page"><img class="wm" src="{f(ASSETS)}/pattern.svg"><div class="frame">{inner}</div></div>'

def header(eyebrow):
    return (f'<div class="head"><img class="sym" src="{f(ASSETS)}/symbol-petrol.svg">'
            f'<span class="eyebrow">{eyebrow}</span>'
            f'<img class="mark" src="{f(ASSETS)}/wordmark-petrol.svg"></div><div class="rule-gold"></div>')

def foot(n):
    return ('<div class="foot"><span class="tag">Saúde, Performance &amp; Longevidade</span>'
            f'<span>Programa de necessidades · Clínica Plenya Londrina · {n}</span>'
            '<span>ordem de grandeza — projetista (CAU/CREA) confirma</span></div>')

def zone(zn, name, sub, tot, rows):
    rr = "".join(f'<div class="row"><span class="n">{n}</span><span class="m">{m}</span></div>' for n,m in rows)
    sb = f'<div class="zsub">{sub}</div>' if sub else ""
    tt = f'<span class="ztot">{tot}</span>' if tot else ""
    return (f'<div class="zone"><div class="ztitle"><span class="zn">{zn}</span>{name}{tt}</div>'
            f'{sb}{rr}</div>')

# ---------------- PÁGINA 1 — espaços e tamanhos ----------------
col1 = (
  zone("Z1","Acolhimento &amp; apoio","área útil 62 m²","",[
    ("Recepção + espera (lounge)","28 m²"),("Sanitários (2 + PCD Ø1,50)","10 m²"),
    ("Copa/apoio da equipe","6 m²"),("DML (depósito limpeza)","3 m²"),
    ("Abrigo de resíduos (PGRSS)","4 m²"),("Administrativo / gestão","11 m²")]) +
  zone("Z2","Consulta &amp; devolutiva","área útil 24 m²","",[
    ("Consultório 1 (entrada)","12 m²"),("Consultório 2 / devolutiva","12 m²")]) +
  zone("Z3","Avaliação funcional","exames leves · 24 m²","",[
    ("Sala composição (bio · grip · retina · MoCA)","12 m²"),
    ("Sala repouso cardiometabólica (RMR · ITB/VOP)","12 m²")])
)
col2 = (
  zone("Z4","Coleta laboratorial","área útil 15 m²","",[
    ("Sala / box de coleta (2 cadeiras)","8 m²"),
    ("Processamento pré-analítico","4 m²"),("Sanitário vinculado","3 m²")]) +
  zone("Z5","Centro de infusão","dispara enfermagem · 41 m²","",[
    ("Sala de aplicação","5 m²"),("Lounge de infusão (4 poltronas)","24 m²"),
    ("Posto de enfermagem","6 m²"),("Preparo + rede frio","6 m²")]) +
  zone("Z6","Expansão (previsão)","cardio/imagem · 30 m²","",[
    ("Sala ergometria / ergoespirometria","18 m²"),
    ("Sala ultrassom / eco","12 m²"),("(opcional) plataforma de força","+20 m²")])
)

totals = (
  '<div class="totals"><div class="tband">'
  '<div class="tcard"><div class="tl">A · Núcleo</div><div class="tv">~170 m²</div>'
  '<div class="td">Exames + consulta, sem infusão (Z1–Z4)</div></div>'
  '<div class="tcard"><div class="tl">B · Com infusão</div><div class="tv">~225 m²</div>'
  '<div class="td">+ centro de infusão / lounge (Z1–Z5)</div></div>'
  '<div class="tcard"><div class="tl">C · Completo</div><div class="tv">~265 m²</div>'
  '<div class="td">+ esforço e ultrassom previstos (Z1–Z6)</div></div>'
  '</div></div>'
)

p1 = page(
  header("Programa de necessidades · Londrina") +
  '<div class="titlerow"><div class="title">Espaços &amp; dimensões da clínica</div>'
  '<div class="subtitle">Tamanhos confortáveis (premium, acima do mínimo RDC 50). '
  'Total construído soma <b>~35% de circulação</b> (corredores, paredes, acessibilidade) sobre a área útil. '
  'Clínica ambulatorial <b>sem radiação e sem sala limpa</b>.</div></div>'
  f'<div class="cols"><div class="col">{col1}</div><div class="col">{col2}</div></div>'
  + totals + foot("pág. 1/3")
)

# ---------------- PÁGINA 2 — o que é cada espaço ----------------
def block(zn, name, body):
    return f'<div class="block"><div class="bh"><span class="zn">{zn}</span>{name}</div><div class="bp">{body}</div></div>'

p2 = page(
  header("O que é cada espaço") +
  '<div class="titlerow"><div class="title">Para que serve cada zona</div></div>'
  '<div class="callout"><div class="ch">Dois simplificadores que barateiam a obra</div>'
  '<p><b>Zero ambientes com radiação</b> — densitometria, tomografia e ressonância são terceirizadas: '
  'nenhuma sala blindada, área controlada ou licenciamento nuclear. &nbsp;<b>Não manipula medicação estéril</b> '
  '(compra pronta) — nenhuma sala limpa, cabine de fluxo laminar ou licença de farmácia. '
  'É uma <b>clínica ambulatorial premium</b>, não um centro de imagem nem farmácia.</p></div>' +
  block("Z1","Acolhimento &amp; apoio",
    "Recepção e espera com ambiência serena (tipo spa-clínico), sanitários acessíveis (NBR 9050), copa, "
    "depósito de limpeza, abrigo de resíduos (PGRSS) e área administrativa.") +
  block("Z2","Consulta &amp; devolutiva",
    "Consultórios médicos — a <b>porta de entrada</b> da jornada e o momento da <b>devolutiva do relatório "
    "integrado de longevidade</b> (o produto de maior valor: consulta → exames → devolutiva → reavaliação).") +
  block("Z3","Avaliação funcional",
    "Exames leves, rápidos e sem radiação, agrupados em duas salas: <b>composição corporal</b> (bioimpedância, "
    "preensão palmar, retinografia, triagem cognitiva) e <b>repouso cardiometabólica</b> (metabolismo de "
    "repouso, índice tornozelo-braço e idade vascular — exige silêncio e penumbra).") +
  block("Z4","Coleta laboratorial",
    "Posto de coleta (sangue + genética/epigenética) com bancada de processamento (centrífuga e geladeira "
    "exclusiva de amostras). Só coleta — a análise é terceirizada no laboratório de apoio. <b>É o coração da "
    "medicina de longevidade</b> (sangue avançado).") +
  block("Z5","Centro de infusão",
    "Sala de aplicação + <b>lounge de poltronas</b> (a experiência premium) + posto de enfermagem + preparo + "
    "rede frio. É o serviço que <b>exige enfermagem</b> e eleva a classificação de risco — pode ser construído "
    "depois do centro de exames.") +
  block("Z6","Expansão cardio/imagem",
    "Sala de esforço (ergometria/ergoespirometria, VO₂máx) e sala de ultrassom/eco. Dependem de <b>trazer "
    "profissional</b> (cardiologista + ultrassonografista), não de obra — por isso entram como <b>previsão de "
    "área e elétrica</b>, evitando uma segunda reforma.")
  + foot("pág. 2/3")
)

# ---------------- PÁGINA 3 — obra + decisões ----------------
p3 = page(
  header("Sequência da obra &amp; decisões") +
  '<div class="titlerow"><div class="title">Como executar e o que decidir</div></div>'
  '<div class="block"><div class="bh">Sequência regulatória (não pular)</div>'
  '<ul class="steps">'
  '<li><b>Projeto aprovado na DAPES/SESA-PR ANTES de reformar</b> (Res. SESA 1.891/2024 — até 90 dias de análise). Nada de quebrar parede antes.</li>'
  '<li><b>Classificação de risco sanitário</b> (Res. SESA 1.034/2020) — a infusão tende a médio/alto risco.</li>'
  '<li>Licença Sanitária Municipal + Alvará (Londrina) · CNES · <b>AVCB/Bombeiros</b> · PGRSS aprovado.</li>'
  '<li>RT médico (CRM-PR) + — quando houver infusão — <b>ERT de enfermagem (COREN-PR)</b>.</li>'
  '<li>Transversais RDC 50: revestimentos laváveis, cantos arredondados em área crítica, <b>fluxo limpo × sujo</b>, climatização e elétrica de emergência.</li>'
  '</ul></div>'
  '<div class="block"><div class="bh">Cenários de obra (faseável)</div>'
  '<div class="bp"><b>A · Núcleo (~170 m²):</b> exames + consulta — leve, sem enfermagem/radiação/sala limpa, '
  'risco regulatório baixo, obra rápida. Já entrega a jornada one-stop. &nbsp;'
  '<b>B · Com infusão (~225 m²):</b> + lounge de infusão (dispara enfermagem e rede frio). &nbsp;'
  '<b>C · Completo (~265 m²):</b> + sala de esforço e ultrassom previstas desde a planta.</div></div>'
  '<div class="block"><div class="bh">Decisões para levar ao arquiteto</div>'
  '<ul class="steps decisions">'
  '<li><b>Infusão na 1ª obra ou faseada depois?</b> Muda metragem, enfermagem e classificação de risco.</li>'
  '<li><b>Reservar a área da Zona 6</b> (esforço + ultrassom) desde já? Evita 2ª reforma.</li>'
  '<li><b>Vertical esportiva</b> entra? Define se precisa de sala ampla p/ plataforma de força.</li>'
  '<li><b>Nº de poltronas de infusão</b> (~6 m² cada) — dimensiona o lounge.</li>'
  '<li><b>Devolutiva</b> no consultório ou em sala de reunião dedicada?</li>'
  '<li><b>Metragem real do imóvel</b> a reformar — define quanto agrupar as zonas.</li>'
  '</ul></div>'
  '<div class="callout"><div class="ch">Mensagem central</div>'
  '<p>Clínica ambulatorial premium, <b>sem radiação e sem sala limpa</b>. O que define a complexidade '
  'regulatória é o <b>centro de infusão</b> (RDC 50 + enfermagem + rede frio + resíduos). O restante é '
  'avaliação funcional leve + coleta + consultórios — agrupáveis para caber no imóvel. '
  'Alavanca de metragem: o lounge escala pelo nº de poltronas; compartilhar salas nas Zonas 3/4/6 '
  'reduz 15–25 m².</p></div>'
  + foot("pág. 3/3")
)

html = f"<!doctype html><html><head><meta charset='utf-8'><style>{FONT_FACE}{CSS}</style></head><body>{p1}{p2}{p3}</body></html>"
OUT_HTML.write_text(html, encoding="utf-8")

cmd = ["chromium","--headless","--no-sandbox","--disable-gpu","--no-pdf-header-footer",
       "--allow-file-access-from-files","--virtual-time-budget=15000",
       "--run-all-compositor-stages-before-draw",
       f"--print-to-pdf={OUT_PDF}", f"file://{OUT_HTML}"]
r = subprocess.run(cmd, capture_output=True, text=True, timeout=180)
print("PDF:", OUT_PDF, "exists:", OUT_PDF.exists(), "size:", OUT_PDF.stat().st_size if OUT_PDF.exists() else 0)
