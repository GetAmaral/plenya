# -*- coding: utf-8 -*-
"""Gera 3 mockups (A4) de Pedido de Exames, um por direção estética.
Conteúdo idêntico; só o tema visual muda. Assets de marca = SVG vetorial."""
import os, pathlib

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
INDICACAO = ("Avaliação metabólica e de longevidade (rastreio do pilar Gestão, Método AGIR). "
             "Paciente assintomática, em programa de acompanhamento contínuo.")

# painel agrupado de exames
GRUPOS = [
    ("Hematologia", ["Hemograma completo com plaquetas"]),
    ("Metabolismo glicêmico", ["Glicemia de jejum", "Insulina de jejum", "Hemoglobina glicada (HbA1c)"]),
    ("Perfil lipídico", ["Colesterol total e frações (HDL, LDL)", "Triglicérides"]),
    ("Função tireoidiana", ["TSH", "T4 livre"]),
    ("Função renal", ["Creatinina com TFG estimada", "Ureia", "Ácido úrico"]),
    ("Função hepática", ["TGO (AST)", "TGP (ALT)", "Gama-GT"]),
    ("Micronutrientes", ["Vitamina D (25-OH)", "Vitamina B12", "Ferritina"]),
    ("Marcadores inflamatórios", ["PCR ultrassensível"]),
    ("Urinálise", ["Urina tipo I (EAS)"]),
]

FOOTER_NAP = ("Plenya Saúde · Av. Ayrton Senna da Silva, 500, Edifício Torre Pietra, sala 1402, "
              "Gleba Palhano, Londrina/PR · (43) 99974-8899 · plenyasaude.com.br")
ICP = "Documento assinado digitalmente com certificado ICP-Brasil."

FONT_FACE = f"""
@font-face {{ font-family:'Nalieta'; src:url('file://{NALIETA}'); font-weight:400; }}
@font-face {{ font-family:'Inter'; src:url('file://{INTER}/Inter-Regular.ttf'); font-weight:400; }}
@font-face {{ font-family:'Inter'; src:url('file://{INTER}/Inter-Medium.ttf'); font-weight:500; }}
@font-face {{ font-family:'Inter'; src:url('file://{INTER}/Inter-SemiBold.ttf'); font-weight:600; }}
@font-face {{ font-family:'Inter'; src:url('file://{INTER}/Inter-Bold.ttf'); font-weight:700; }}
"""

def exam_panel(label_cls, group_cls, item_cls):
    rows = []
    for g, items in GRUPOS:
        lis = "".join(f'<li class="{item_cls}">{it}</li>' for it in items)
        rows.append(f'<div class="exgroup"><div class="{group_cls}">{g}</div>'
                    f'<ul class="exlist">{lis}</ul></div>')
    return f'<div class="panel">{"".join(rows)}</div>'

# ============================================================ A — EDITORIAL CREME
def direction_A():
    css = f"""
    {FONT_FACE}
    * {{ margin:0; padding:0; box-sizing:border-box; }}
    html,body {{ width:794px; height:1123px; }}
    .page {{ position:relative; width:794px; height:1123px; background:#f4f2ea;
             font-family:'Inter',sans-serif; color:#0c2a36; overflow:hidden;
             padding:58px 64px 0 64px; }}
    .wm {{ position:absolute; top:300px; left:50%; transform:translateX(-50%);
           width:520px; opacity:.05; z-index:0; }}
    .layer {{ position:relative; z-index:1; height:100%; display:flex; flex-direction:column; }}
    /* header */
    .head {{ display:flex; align-items:flex-end; justify-content:space-between; }}
    .head .sym {{ width:34px; }}
    .head .mark {{ height:30px; }}
    .rule-gold {{ height:1px; background:#b38645; margin:16px 0 0 0; opacity:.7; }}
    /* title */
    .title {{ font-family:'Nalieta'; font-size:31px; letter-spacing:.5px; color:#063b4f; margin-top:40px; }}
    .title-rule {{ width:46px; height:2px; background:#b38645; margin-top:12px; }}
    /* patient grid */
    .pgrid {{ display:grid; grid-template-columns:1fr 1fr; gap:6px 40px; margin-top:26px; }}
    .lbl {{ font-size:9px; letter-spacing:2px; text-transform:uppercase; color:#417e8e; font-weight:600; }}
    .val {{ font-size:13.5px; color:#0c2a36; font-weight:500; margin-top:1px; }}
    .indic {{ margin-top:22px; font-size:12px; line-height:1.6; color:#324a54; max-width:640px; }}
    .indic .lbl {{ display:block; margin-bottom:4px; }}
    /* panel */
    .panel {{ margin-top:24px; columns:2; column-gap:46px; }}
    .exgroup {{ break-inside:avoid; margin-bottom:15px; }}
    .grp {{ font-family:'Nalieta'; font-size:14px; color:#063b4f; padding-bottom:5px;
            border-bottom:1px solid rgba(179,134,69,.45); margin-bottom:7px; }}
    .exlist {{ list-style:none; }}
    .exitem {{ font-size:11.5px; line-height:1.75; color:#1b3640; padding-left:14px; position:relative; }}
    .exitem::before {{ content:''; position:absolute; left:0; top:9px; width:5px; height:5px;
                       background:#b38645; border-radius:50%; }}
    /* footer */
    .foot {{ margin-top:auto; padding-bottom:30px; }}
    .sigrow {{ display:flex; justify-content:space-between; align-items:flex-end; margin-bottom:18px; }}
    .sig {{ text-align:center; }}
    .sigline {{ width:280px; border-top:1px solid #063b4f; padding-top:7px;
                font-size:12.5px; font-weight:600; color:#063b4f; }}
    .sigsub {{ font-size:9.5px; letter-spacing:1px; color:#417e8e; margin-top:2px; }}
    .icpbox {{ text-align:right; font-size:8.5px; color:#7d8c91; max-width:200px; line-height:1.5; }}
    .qr {{ width:54px; height:54px; border:1px solid #cdbfa3; margin-left:auto; margin-bottom:6px;
           background:
             repeating-linear-gradient(90deg,#063b4f 0 3px,transparent 3px 6px),
             repeating-linear-gradient(0deg,#063b4f 0 3px,transparent 3px 6px);
           opacity:.55; }}
    .nap {{ border-top:1px solid #b38645; padding-top:10px; font-size:8.5px; letter-spacing:.3px;
            color:#5a7079; text-align:center; line-height:1.55; }}
    """
    body = f"""
    <div class="page">
      <img class="wm" src="{a('symbol-gold.svg')}">
      <div class="layer">
        <div class="head">
          <img class="sym" src="{a('symbol-petrol.svg')}">
          <img class="mark" src="{a('lockup-tagline-petrol.svg')}">
        </div>
        <div class="rule-gold"></div>

        <div class="title">Pedido de Exames</div>
        <div class="title-rule"></div>

        <div class="pgrid">
          <div><div class="lbl">Paciente</div><div class="val">{PACIENTE}</div></div>
          <div><div class="lbl">Data de nascimento</div><div class="val">{NASC} · {IDADE}</div></div>
          <div><div class="lbl">CPF</div><div class="val">{CPF}</div></div>
          <div><div class="lbl">Data</div><div class="val">{DATA}</div></div>
        </div>

        <div class="indic"><span class="lbl">Indicação clínica</span>{INDICACAO}</div>

        {exam_panel('lbl','grp','exitem')}

        <div class="foot">
          <div class="sigrow">
            <div class="sig">
              <div class="sigline">{MEDICO}</div>
              <div class="sigsub">{CRM}</div>
            </div>
            <div>
              <div class="qr"></div>
              <div class="icpbox">{ICP}</div>
            </div>
          </div>
          <div class="nap">{FOOTER_NAP}</div>
        </div>
      </div>
    </div>"""
    return f"<!doctype html><html><head><meta charset='utf-8'><style>{css}</style></head><body>{body}</body></html>"

# ============================================================ B — MASTHEAD PETROL
def direction_B():
    css = f"""
    {FONT_FACE}
    * {{ margin:0; padding:0; box-sizing:border-box; }}
    html,body {{ width:794px; height:1123px; }}
    .page {{ position:relative; width:794px; height:1123px; background:#ffffff;
             font-family:'Inter',sans-serif; color:#16313c; overflow:hidden; }}
    .band {{ background:#063b4f; height:118px; padding:0 56px; display:flex;
             align-items:center; justify-content:space-between; }}
    .band .mark {{ height:34px; }}
    .band .tag {{ font-size:9px; letter-spacing:3px; color:#92b8b4; text-transform:uppercase; margin-top:9px; }}
    .band .sym {{ width:46px; opacity:.92; }}
    .gold-edge {{ height:3px; background:#b38645; }}
    .content {{ padding:40px 56px 0 56px; height:calc(1123px - 118px - 3px); display:flex; flex-direction:column; }}
    .title {{ font-family:'Nalieta'; font-size:29px; color:#063b4f; }}
    .pgrid {{ display:grid; grid-template-columns:1fr 1fr; gap:5px 40px; margin-top:20px;
              background:#f5f7f7; border-left:3px solid #b38645; padding:16px 20px; }}
    .lbl {{ font-size:9px; letter-spacing:2px; text-transform:uppercase; color:#417e8e; font-weight:600; }}
    .val {{ font-size:13.5px; color:#16313c; font-weight:500; margin-top:1px; }}
    .indic {{ margin-top:18px; font-size:12px; line-height:1.6; color:#3a525c; }}
    .indic .lbl {{ display:block; margin-bottom:4px; }}
    .panel {{ margin-top:22px; columns:2; column-gap:44px; }}
    .exgroup {{ break-inside:avoid; margin-bottom:14px; }}
    .grp {{ font-size:10px; letter-spacing:1.5px; text-transform:uppercase; font-weight:700; color:#063b4f;
            padding-bottom:5px; border-bottom:2px solid #92b8b4; margin-bottom:7px; }}
    .exlist {{ list-style:none; }}
    .exitem {{ font-size:11.5px; line-height:1.75; color:#1b3640; padding-left:15px; position:relative; }}
    .exitem::before {{ content:''; position:absolute; left:0; top:7px; width:7px; height:2px; background:#b38645; }}
    .foot {{ margin-top:auto; }}
    .sigrow {{ display:flex; justify-content:space-between; align-items:flex-end; padding:0 0 16px 0; }}
    .sigline {{ width:280px; border-top:1px solid #063b4f; padding-top:7px;
                font-size:12.5px; font-weight:600; color:#063b4f; text-align:center; }}
    .sigsub {{ font-size:9.5px; letter-spacing:1px; color:#417e8e; margin-top:2px; text-align:center; }}
    .qr {{ width:52px; height:52px; opacity:.55; border:1px solid #bcd0d0;
           background:repeating-linear-gradient(90deg,#063b4f 0 3px,transparent 3px 6px),
                      repeating-linear-gradient(0deg,#063b4f 0 3px,transparent 3px 6px); }}
    .footband {{ background:#063b4f; color:#cfe0e0; padding:13px 56px; display:flex;
                 justify-content:space-between; align-items:center; font-size:8.5px; letter-spacing:.4px; }}
    .footband .icp {{ color:#92b8b4; }}
    .footnap {{ text-align:center; line-height:1.5; max-width:560px; }}
    """
    body = f"""
    <div class="page">
      <div class="band">
        <div><img class="mark" src="{a('wordmark-cream.svg')}"><div class="tag">Saúde · Performance · Longevidade</div></div>
        <img class="sym" src="{a('symbol-cream.svg')}">
      </div>
      <div class="gold-edge"></div>
      <div class="content">
        <div class="title">Pedido de Exames</div>
        <div class="pgrid">
          <div><div class="lbl">Paciente</div><div class="val">{PACIENTE}</div></div>
          <div><div class="lbl">Data de nascimento</div><div class="val">{NASC} · {IDADE}</div></div>
          <div><div class="lbl">CPF</div><div class="val">{CPF}</div></div>
          <div><div class="lbl">Data</div><div class="val">{DATA}</div></div>
        </div>
        <div class="indic"><span class="lbl">Indicação clínica</span>{INDICACAO}</div>
        {exam_panel('lbl','grp','exitem')}
        <div class="foot">
          <div class="sigrow">
            <div><div class="sigline">{MEDICO}</div><div class="sigsub">{CRM}</div></div>
            <div class="qr"></div>
          </div>
        </div>
      </div>
      <div class="footband">
        <div class="footnap">{FOOTER_NAP}</div>
        <div class="icp">ICP-Brasil ✓</div>
      </div>
    </div>"""
    return f"<!doctype html><html><head><meta charset='utf-8'><style>{css}</style></head><body>{body}</body></html>"

# ============================================================ C — MINIMAL SUÍÇO
def direction_C():
    css = f"""
    {FONT_FACE}
    * {{ margin:0; padding:0; box-sizing:border-box; }}
    html,body {{ width:794px; height:1123px; }}
    .page {{ position:relative; width:794px; height:1123px; background:#fcfcfa;
             font-family:'Inter',sans-serif; color:#1a2f38; overflow:hidden; padding:60px 70px 0 70px; }}
    .sym {{ width:30px; }}
    .title {{ font-family:'Nalieta'; font-size:34px; color:#063b4f; margin-top:46px; letter-spacing:.5px; }}
    .title-rule {{ width:40px; height:2px; background:#b38645; margin-top:14px; }}
    .pgrid {{ margin-top:38px; }}
    .prow {{ display:flex; align-items:baseline; padding:9px 0; border-bottom:1px solid #ecebe4; }}
    .lbl {{ width:170px; font-size:9px; letter-spacing:2.5px; text-transform:uppercase;
            color:#7d8c91; font-weight:600; }}
    .val {{ font-size:13.5px; color:#1a2f38; font-weight:500; }}
    .indic {{ margin-top:24px; font-size:11.5px; line-height:1.65; color:#56666d; max-width:600px; }}
    .indic .lbl {{ width:auto; display:block; margin-bottom:5px; }}
    .seclabel {{ margin-top:34px; font-size:9px; letter-spacing:2.5px; text-transform:uppercase;
                 color:#7d8c91; font-weight:600; }}
    .panel {{ margin-top:16px; columns:2; column-gap:50px; }}
    .exgroup {{ break-inside:avoid; margin-bottom:14px; }}
    .grp {{ font-size:11.5px; color:#063b4f; font-weight:600; margin-bottom:5px; }}
    .exlist {{ list-style:none; }}
    .exitem {{ font-size:11px; line-height:1.7; color:#3a4a51; }}
    .foot {{ position:absolute; left:70px; right:70px; bottom:48px; }}
    .sigrow {{ display:flex; justify-content:space-between; align-items:flex-end; }}
    .sigline {{ width:260px; border-top:1px solid #1a2f38; padding-top:7px; font-size:12px;
                font-weight:600; color:#063b4f; }}
    .sigsub {{ font-size:9px; letter-spacing:1px; color:#7d8c91; margin-top:2px; }}
    .markwrap {{ text-align:right; }}
    .markwrap img {{ height:16px; opacity:.85; }}
    .markwrap .icp {{ font-size:8px; letter-spacing:.5px; color:#9aa6aa; margin-top:6px; }}
    .nap {{ margin-top:20px; padding-top:10px; border-top:1px solid #ecebe4;
            font-size:8px; letter-spacing:.3px; color:#9aa6aa; text-align:center; }}
    """
    rows = "".join(
        f'<div class="prow"><div class="lbl">{l}</div><div class="val">{v}</div></div>'
        for l, v in [("Paciente", PACIENTE), ("Nascimento", f"{NASC} · {IDADE}"),
                     ("CPF", CPF), ("Data", DATA)])
    body = f"""
    <div class="page">
      <img class="sym" src="{a('symbol-petrol.svg')}">
      <div class="title">Pedido de Exames</div>
      <div class="title-rule"></div>
      <div class="pgrid">{rows}</div>
      <div class="indic"><span class="lbl">Indicação clínica</span>{INDICACAO}</div>
      <div class="seclabel">Exames solicitados</div>
      {exam_panel('lbl','grp','exitem')}
      <div class="foot">
        <div class="sigrow">
          <div><div class="sigline">{MEDICO}</div><div class="sigsub">{CRM}</div></div>
          <div class="markwrap"><img src="{a('wordmark-petrol.svg')}"><div class="icp">{ICP}</div></div>
        </div>
        <div class="nap">{FOOTER_NAP}</div>
      </div>
    </div>"""
    return f"<!doctype html><html><head><meta charset='utf-8'><style>{css}</style></head><body>{body}</body></html>"

for name, fn in [("A-editorial", direction_A), ("B-masthead", direction_B), ("C-minimal", direction_C)]:
    (DRAFTS / f"{name}.html").write_text(fn(), encoding="utf-8")
    print("wrote", name)
