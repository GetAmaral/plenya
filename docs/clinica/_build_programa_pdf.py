# -*- coding: utf-8 -*-
"""Programa de Necessidades (arquiteto) — PDF A4 branded Plenya, estrutura MULTI-CLÍNICA.
Pág.1 = PLENYA-EXCLUSIVO (espaços+tamanhos); pág.2 = COMUM/EMPREENDIMENTO; pág.3 = explicações/obra.
chromium print-to-pdf."""
import pathlib, subprocess

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
  --ink2:rgba(6,59,79,.82); --ink3:rgba(6,59,79,.58); --line:rgba(6,59,79,.14); }
*{ margin:0; padding:0; box-sizing:border-box; }
@page{ size:A4; margin:0; }
html,body{ background:#fff; }
.page{ position:relative; width:210mm; height:297mm; background:var(--cream);
  font-family:'Inter',sans-serif; color:var(--petrol); overflow:hidden; page-break-after:always; }
.page:last-child{ page-break-after:auto; }
.wm{ position:absolute; inset:0; width:100%; height:100%; object-fit:cover; opacity:.035; z-index:0; }
.frame{ position:relative; z-index:1; height:100%; padding:15mm 17mm 18mm; display:flex; flex-direction:column; }
.head{ display:flex; align-items:center; justify-content:space-between; }
.head .sym{ height:26px; } .head .mark{ height:20px; }
.eyebrow{ font-size:9.5px; letter-spacing:1.4px; text-transform:uppercase; color:var(--gold); font-weight:600; }
.rule-gold{ height:1.2px; background:var(--gold); margin-top:9px; }
.titlerow{ margin-top:14px; }
.title{ font-family:'Cormorant',serif; font-size:30px; font-weight:600; color:var(--petrol); line-height:1.02; }
.subtitle{ font-size:11px; color:var(--ink2); margin-top:6px; line-height:1.5; max-width:174mm; }
.subtitle b{ color:var(--petrol); }
/* grupos de ambientes */
.grp{ margin-top:13px; }
.gh{ font-family:'Cormorant',serif; font-size:16px; font-weight:700; color:var(--petrol);
  display:flex; align-items:baseline; gap:8px; border-bottom:1.2px solid var(--gold); padding-bottom:4px; }
.gh .gt{ margin-left:auto; font-family:'Inter'; font-size:10px; font-weight:600; color:var(--ocean); }
.row{ display:flex; align-items:baseline; gap:10px; font-size:11px; line-height:1.4;
  padding:4px 0; border-bottom:1px solid var(--line); }
.row .nm{ flex:1; color:var(--petrol); }
.row .nm b{ font-weight:600; }
.row .nm .d{ color:var(--ink3); font-size:10px; }
.row .ref{ width:74px; text-align:right; color:var(--ink3); font-size:9.5px; white-space:nowrap; }
.row .m{ width:62px; text-align:right; color:var(--petrol); font-weight:700; white-space:nowrap; }
.colhead{ display:flex; gap:10px; font-size:8.5px; letter-spacing:.4px; text-transform:uppercase;
  color:var(--ink3); padding-bottom:2px; }
.colhead .a{ flex:1; } .colhead .b{ width:74px; text-align:right; } .colhead .c{ width:62px; text-align:right; }
/* totais */
.totals{ margin-top:auto; }
.tband{ display:flex; gap:6mm; }
.tcard{ flex:1; background:rgba(6,59,79,.05); border:1px solid var(--line); border-left:3px solid var(--gold);
  border-radius:5px; padding:9px 11px; }
.tcard .tl{ font-size:9px; letter-spacing:.5px; text-transform:uppercase; color:var(--ink3); font-weight:600; }
.tcard .tv{ font-family:'Cormorant',serif; font-size:25px; font-weight:700; color:var(--petrol); line-height:1.05; margin-top:2px; }
.tcard .td{ font-size:9px; color:var(--ink2); margin-top:2px; line-height:1.3; }
.note{ font-size:9.5px; color:var(--ink3); margin-top:8px; line-height:1.4; }
.foot{ position:absolute; left:17mm; right:17mm; bottom:10mm; display:flex; align-items:center;
  justify-content:space-between; padding-top:7px; border-top:1px solid var(--line);
  font-size:8.8px; color:var(--ink3); letter-spacing:.3px; }
.foot .tag{ text-transform:uppercase; font-weight:600; color:var(--gold); }
/* callout / blocos */
.callout{ background:rgba(6,59,79,.05); border:1px solid var(--line); border-left:3px solid var(--gold);
  border-radius:5px; padding:11px 13px; margin-top:13px; }
.callout.crit{ border-left-color:#9a3b2e; }
.callout .ch{ font-family:'Cormorant',serif; font-size:15px; font-weight:700; color:var(--petrol); margin-bottom:4px; }
.callout p{ font-size:10.5px; color:var(--ink2); line-height:1.5; }
.callout b{ color:var(--petrol); }
.block{ margin-top:13px; break-inside:avoid; }
.bh{ font-family:'Cormorant',serif; font-size:15px; font-weight:700; color:var(--petrol); }
.bp{ font-size:10.5px; color:var(--ink2); line-height:1.55; margin-top:4px; }
.bp b{ color:var(--petrol); font-weight:600; }
ul.steps{ margin-top:5px; padding:0; list-style:none; }
ul.steps li{ font-size:10.5px; color:var(--ink2); line-height:1.45; padding:2.5px 0 2.5px 19px; position:relative; }
ul.steps li:before{ content:""; position:absolute; left:3px; top:8px; width:5px; height:5px;
  background:var(--gold); border-radius:1px; transform:rotate(45deg); }
ul.steps li b{ color:var(--petrol); }
.decisions{ counter-reset:d; } .decisions li{ counter-increment:d; padding-left:23px; }
.decisions li:before{ content:counter(d); background:var(--petrol); color:var(--cream); width:15px; height:15px;
  border-radius:50%; font-size:9px; font-weight:700; display:flex; align-items:center; justify-content:center;
  transform:none; top:3px; left:2px; }
"""

def page(inner):
    return f'<div class="page"><img class="wm" src="{f(ASSETS)}/pattern.svg"><div class="frame">{inner}</div></div>'
def header(eb):
    return (f'<div class="head"><img class="sym" src="{f(ASSETS)}/symbol-petrol.svg">'
            f'<span class="eyebrow">{eb}</span>'
            f'<img class="mark" src="{f(ASSETS)}/wordmark-petrol.svg"></div><div class="rule-gold"></div>')
def foot(n):
    return ('<div class="foot"><span class="tag">Saúde, Performance &amp; Longevidade</span>'
            f'<span>Programa de necessidades · {n}</span>'
            '<span>valores: ordem de grandeza</span></div>')
def grp(title, tot, rows):
    ch = '<div class="colhead"><span class="a">Ambiente</span><span class="b">mín. RDC 50</span><span class="c">adotado</span></div>'
    rr = "".join(f'<div class="row"><span class="nm">{nm}</span><span class="ref">{ref}</span><span class="m">{m}</span></div>' for nm,ref,m in rows)
    gt = f'<span class="gt">{tot}</span>' if tot else ""
    return f'<div class="grp"><div class="gh">{title}{gt}</div>{ch}{rr}</div>'
def block(name, body):
    return f'<div class="block"><div class="bh">{name}</div><div class="bp">{body}</div></div>'

# ---------------- PÁGINA 1 — PLENYA-EXCLUSIVO ----------------
p1 = page(
  header("Parte 1 · Plenya-exclusivo") +
  '<div class="titlerow"><div class="title">Espaços &amp; dimensões — área Plenya</div>'
  '<div class="subtitle">O que a Plenya constrói/ocupa dentro do empreendimento multi-clínica. '
  'Sizing <b>premium</b> (acima do mínimo RDC 50, mostrado ao lado). Recepção, espera e sanitários são '
  '<b>comuns do empreendimento</b> (pág. 2). Total construído soma <b>~25% de circulação interna</b>.</div></div>'
  + grp("Consulta &nbsp;<span style='font-size:10px;color:var(--ink3)'>(2 salas)</span>", "37 m² útil", [
      ("<b>Consultório médico premium</b> — Dr. Getúlio <span class='d'>+ devolutiva do relatório</span>","7,5 m²","25 m²"),
      ("<b>Consultório multidisciplinar</b> <span class='d'>menor — nutri/equipe</span>","7,5 m²","12 m²"),
    ])
  + grp("Exames &nbsp;<span style='font-size:10px;color:var(--ink3)'>(consolidados em 1 sala)</span>", "28 m² útil", [
      ("<b>Sala de Avaliação Funcional &amp; Longevidade</b><br><span class='d'>bioimpedância · preensão · retinografia · MoCA/olfato · espirometria · ITB/VOP · RMR (calorimetria) · VO₂ estimado · ponto de coleta</span>","—","28 m²"),
    ])
  + grp("Centro de infusão &nbsp;<span style='font-size:10px;color:var(--ink3)'>(diferencial Plenya)</span>", "46 m² útil", [
      ("<b>Lounge de infusão</b> <span class='d'>4 poltronas × 8 m²</span>","5/polt.","32 m²"),
      ("Posto de enfermagem","—","8 m²"),
      ("Preparo + rede frio <span class='d'>(punção no lounge)</span>","—","6 m²"),
    ])
  + grp("Expansão &nbsp;<span style='font-size:10px;color:var(--ink3)'>(previsão — depende de profissional)</span>", "", [
      ("Sala de esforço <span class='d'>ergometria/ergoespirometria — médico presente</span>","—","22 m²"),
      ("Ultrassom/eco <span class='d'>compartilha a sala de avaliação, ou dedicada</span>","—","0–16 m²"),
    ])
  + '<div class="totals"><div class="tband">'
    '<div class="tcard"><div class="tl">A · Consulta + exames</div><div class="tv">~80 m²</div>'
    '<div class="td">2 consultórios + sala multi-exame</div></div>'
    '<div class="tcard"><div class="tl">B · Com infusão</div><div class="tv">~140 m²</div>'
    '<div class="td">+ lounge / enfermagem / preparo</div></div>'
    '<div class="tcard"><div class="tl">C · Completo</div><div class="tv">~165–185 m²</div>'
    '<div class="td">+ esforço (ultrassom compart.→dedicado)</div></div>'
    '</div></div>'
  + foot("pág. 1/3")
)

# ---------------- PÁGINA 2 — COMUM / EMPREENDIMENTO ----------------
def row_common(nm, ind, req):
    return (f'<div class="row"><span class="nm">{nm}</span>'
            f'<span class="ref" style="width:60px">{ind}</span>'
            f'<span class="nm" style="flex:1.3;color:var(--ink2);font-size:10px">{req}</span></div>')

p2 = page(
  header("Parte 2 · Comum do empreendimento") +
  '<div class="titlerow"><div class="title">Espaços compartilhados</div>'
  '<div class="subtitle">Não entram no CAPEX exclusivo da Plenya, mas a clínica <b>depende</b> deles. '
  'Definir com o empreendimento o que é garantido e o rateio.</div></div>'
  '<div class="grp"><div class="gh">Áreas comuns<span class="gt">empreendimento</span></div>'
  '<div class="colhead"><span class="a">Ambiente comum</span>'
  '<span class="b" style="width:60px">indicativo</span>'
  '<span class="c" style="flex:1.3;width:auto;text-align:left;padding-left:10px">requisito que a Plenya impõe</span></div>'
  + row_common("Recepção + espera / lounge","35–50 m²","ambiência premium (marca) · 1,2 m²/pessoa")
  + row_common("Sanitários (masc / fem / <b>PCD</b>)","~12 m²","NBR 9050 — PCD giro Ø 1,50 m + barras")
  + row_common("Copa / estar de funcionários","—","—")
  + row_common("DML + <b>abrigo de resíduos (PGRSS)</b>","—","aceitar Grupo B (farmacológico) + E (perfurocortante)")
  + row_common("Circulação comum (corredores)","—","≥1,20 m (≥2,00 m se maca)")
  + '</div>'
  '<div class="callout crit"><div class="ch">Alinhar com o empreendimento (crítico)</div>'
  '<p>A <b>infusão da Plenya eleva a classificação de risco sanitário</b> e exige PGRSS / abrigo de '
  'resíduos (e possivelmente rede frio) compatíveis — o prédio precisa suportar. <b>Confirmar com a VISA '
  'Londrina</b> se a classificação de risco recai só sobre a sala Plenya ou sobre o empreendimento inteiro.</p></div>'
  '<div class="callout"><div class="ch">Dois simplificadores que barateiam a obra</div>'
  '<p><b>Zero ambientes com radiação</b> (densitometria/tomografia/ressonância terceirizadas) — sem sala '
  'blindada, área controlada ou licenciamento nuclear. &nbsp;<b>Não manipula medicação estéril</b> (compra '
  'pronta) — sem sala limpa, fluxo laminar ou licença de farmácia. É uma <b>clínica ambulatorial premium</b>.</p></div>'
  '<div class="callout"><div class="ch">A otimização dos exames</div>'
  '<p>Em vez de uma sala por exame, <b>uma Sala de Avaliação Funcional faz quase tudo</b> (composição, força, '
  'pulmão, idade vascular, metabolismo, cognição/retina + coleta). Ressalva: a RMR ocupa ~30 min reclinado — '
  'se o volume crescer, desmembra-se uma 2ª sala de repouso (~12 m²). <b>Começa com uma só.</b></p></div>'
  + foot("pág. 2/3")
)

# ---------------- PÁGINA 3 — espaços, obra, decisões ----------------
p3 = page(
  header("Parte 3 · Detalhes &amp; decisões") +
  '<div class="titlerow"><div class="title">O que é cada espaço · obra · decisões</div></div>'
  + block("Os espaços Plenya em uma linha",
    "<b>Consultório premium (25 m²)</b>: entrada da jornada + devolutiva do relatório. "
    "<b>Consultório multidisciplinar (12 m²)</b>: nutri/equipe. "
    "<b>Sala de Avaliação (28 m²)</b>: todos os exames próprios sem radiação, em estações, + coleta. "
    "<b>Centro de infusão</b>: lounge + enfermagem + preparo (dispara enfermagem e classificação de risco). "
    "<b>Esforço + ultrassom</b>: previsão, entram com o cardiologista/ultrassonografista.")
  + '<div class="block"><div class="bh">Sequência regulatória da obra (não pular)</div>'
    '<ul class="steps">'
    '<li><b>Projeto aprovado na DAPES/SESA-PR ANTES de reformar</b> (Res. SESA 1.891/2024 — até 90 dias).</li>'
    '<li><b>Classificação de risco sanitário</b> (Res. SESA 1.034/2020) — a infusão tende a médio/alto.</li>'
    '<li>Licença Sanitária + Alvará (Londrina) · CNES · AVCB/Bombeiros · PGRSS.</li>'
    '<li>RT médico (CRM-PR) + — com infusão — <b>ERT enfermagem (COREN-PR)</b>.</li>'
    '<li>RDC 50: revestimentos laváveis, cantos arredondados em área crítica, <b>fluxo limpo × sujo</b>.</li>'
    '</ul></div>'
  + '<div class="block"><div class="bh">Decisões para a reunião</div>'
    '<ul class="steps decisions">'
    '<li><b>Infusão na 1ª obra ou faseada?</b> Muda metragem, enfermagem e risco.</li>'
    '<li><b>Reservar área de esforço + ultrassom</b> desde já? Evita 2ª reforma.</li>'
    '<li><b>Ultrassom compartilha a sala de avaliação</b> (econômico) ou dedicada?</li>'
    '<li><b>Nº de poltronas</b> de infusão (~8 m² cada) — dimensiona o lounge.</li>'
    '<li><b>Áreas comuns garantidas</b> pelo empreendimento + rateio.</li>'
    '<li><b>Classificação de risco</b>: recai sobre a Plenya ou o empreendimento?</li>'
    '<li><b>Metragem real</b> disponível para a Plenya no empreendimento.</li>'
    '</ul></div>'
  + '<div class="callout"><div class="ch">Mensagem central</div>'
    '<p>A área <b>Plenya-exclusiva é enxuta (~80 · ~140 · ~165–185 m²</b> nas fases A·B·C) porque recepção, espera e '
    'sanitários são <b>comuns do empreendimento</b> e os exames foram <b>consolidados numa sala multi-exame</b>. '
    'O que define a complexidade regulatória é o <b>centro de infusão</b> — e o empreendimento precisa suportá-lo.</p></div>'
  + foot("pág. 3/3")
)

html = f"<!doctype html><html><head><meta charset='utf-8'><style>{FONT_FACE}{CSS}</style></head><body>{p1}{p2}{p3}</body></html>"
OUT_HTML.write_text(html, encoding="utf-8")
cmd = ["chromium","--headless","--no-sandbox","--disable-gpu","--no-pdf-header-footer",
       "--allow-file-access-from-files","--virtual-time-budget=15000",
       "--run-all-compositor-stages-before-draw", f"--print-to-pdf={OUT_PDF}", f"file://{OUT_HTML}"]
subprocess.run(cmd, capture_output=True, text=True, timeout=180)
print("PDF:", OUT_PDF.exists(), OUT_PDF.stat().st_size if OUT_PDF.exists() else 0)
