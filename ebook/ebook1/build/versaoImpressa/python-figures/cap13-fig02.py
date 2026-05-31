"""
Cap13 Fig02 (PT-BR, B&W vetorial) — "Os exames estão bons. Eu não estou."
Ricardo, 3 timeframes × 2 dimensões.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, FancyBboxPatch

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False
rcParams["pdf.fonttype"] = 42
rcParams["ps.fonttype"] = 42

BG       = "#FFFFFF"
INK      = "#000000"
INK_SOFT = "#3A3A3A"
TICK     = "#555555"
FOOT     = "#666666"
COL_1    = "#F4F4F4"   # T+18m
COL_2    = "#E0E0E0"   # T+21m
COL_3    = "#C8C8C8"   # T+30m
BAND     = "#EDEDED"

fig = plt.figure(figsize=(11.0, 7.8))
fig.patch.set_facecolor(BG)

LEFT = 0.025

# Título
fig.text(LEFT, 0.955, "Figura 2",
         fontsize=8.5, color=INK_SOFT, weight="bold")
fig.text(LEFT, 0.928,
         "\"Os exames estão bons. Eu não estou.\"",
         fontsize=16, color=INK, weight="bold")
fig.text(LEFT, 0.895,
         "Ricardo, 18 meses após o infarto: painel biológico em ordem, dimensão relacional em colapso.",
         fontsize=9.5, color=INK_SOFT)
fig.text(LEFT, 0.876,
         "O trabalho em conexão, propósito e sentido restaurou os dois lados.",
         fontsize=9.5, color=INK_SOFT, style="italic")

# 3 colunas de tempo (headers)
COL_X_CENTER = [0.38, 0.58, 0.80]
COL_HEAD_Y = 0.830
COL_SUB_Y  = 0.810
COL_W = 0.18

headers = [
    ("T+18m", "Colapso silencioso", COL_1),
    ("T+21m", "Intervenção multi-frente", COL_2),
    ("T+30m", "Recuperação completa", COL_3),
]

for i, (head, sub, color) in enumerate(headers):
    cx = COL_X_CENTER[i]
    fig.text(cx, COL_HEAD_Y, head,
             fontsize=11, color=INK, weight="bold", ha="center")
    fig.text(cx, COL_SUB_Y, sub,
             fontsize=8.5, color=INK_SOFT, ha="center", style="italic")

# Tabela: 2 categorias × items
# Coluna de labels (esquerda) — recuada pra dar espaço pra label vertical
LBL_COL_X = LEFT + 0.045

# ---------- DIMENSÃO HUMANA — label vertical à esquerda do bloco ----------
fig.text(LEFT - 0.000, 0.560, "DIMENSÃO\nHUMANA",
         fontsize=11, color=INK, weight="bold",
         ha="left", va="center", linespacing=1.15,
         rotation=90)

rows_humana = [
    ("Vida sexual\ncom Marina",
     "8 meses sem",
     "reentrada incipiente\n+ tadalafila 5 mg/dia",
     "plena, qualidade\n> pré-IAM"),
    ("Amigos próximos (n*)",
     "0",
     "2",
     "4"),
    ("Tempo no quarto",
     "0",
     "retraindo",
     "10 (mas semanal firme)"),
    ("Propósito profissional",
     "esvaziado",
     "mentoria voluntária\niniciada",
     "mentoria semanal firme"),
    ("Ritual laico com Marina",
     "0",
     "caminhada silenciosa\ndominical",
     "52 domingos/ano"),
]

ROW_Y_START = 0.730
ROW_H = 0.060

for ri, (label, v1, v2, v3) in enumerate(rows_humana):
    y = ROW_Y_START - ri * ROW_H

    # Label
    fig.text(LBL_COL_X, y, label,
             fontsize=8.5, color=INK, weight="bold",
             va="center", linespacing=1.2)

    # 3 valores nas colunas
    for i, val in enumerate([v1, v2, v3]):
        cx = COL_X_CENTER[i]
        # Bolinha
        fig.text(cx - 0.075, y, "●",
                 fontsize=10, color=INK, va="center")
        fig.text(cx - 0.063, y, val,
                 fontsize=8, color=INK, va="center",
                 linespacing=1.2)

# ---------- CONTROLE BIOLÓGICO — label vertical à esquerda ----------
fig.text(LEFT - 0.000, 0.300, "CONTROLE\nBIOLÓGICO\n(marcadores)",
         fontsize=10, color=INK, weight="bold",
         ha="left", va="center", linespacing=1.15,
         rotation=90)

rows_bio = [
    ("PCR ultrassensível\n(mg/L)",   "1,4",        "1,2",        "0,6"),
    ("Cortisol matinal\n(µg/dL)",    "↑",          "↓",          "adequado"),
]

ROW_Y_START_BIO = 0.350
ROW_H_BIO = 0.060

for ri, (label, v1, v2, v3) in enumerate(rows_bio):
    y = ROW_Y_START_BIO - ri * ROW_H_BIO

    fig.text(LBL_COL_X, y, label,
             fontsize=8.5, color=INK, weight="bold",
             va="center", linespacing=1.2)

    for i, val in enumerate([v1, v2, v3]):
        cx = COL_X_CENTER[i]
        fig.text(cx - 0.075, y, "●",
                 fontsize=10, color=INK, va="center")
        fig.text(cx - 0.063, y, val,
                 fontsize=9, color=INK, weight="bold", va="center")

# ---------- Quote ----------
QUOTE_Y = 0.180
fig.text(LEFT + 0.030, QUOTE_Y + 0.020, "“",
         fontsize=28, color=INK, weight="bold", va="top")
fig.text(LEFT + 0.090, QUOTE_Y,
         "\"Doutor, eu achei que ia sair dessa história mais frágil.",
         fontsize=10, color=INK, style="italic", va="center")
fig.text(LEFT + 0.090, QUOTE_Y - 0.020,
         "Eu acho que sai mais inteiro.\"",
         fontsize=10, color=INK, weight="bold", style="italic", va="center")
fig.text(0.95, QUOTE_Y - 0.030,
         "— Ricardo, T+30m",
         fontsize=8.5, color=INK_SOFT, ha="right", style="italic")

# Footer
fig.text(LEFT, 0.080,
         "Fonte: caso-tipo do Cap. 13. Princeton Consensus III (Nehra et al., Mayo Clinic Proc 2012);",
         fontsize=7.5, color=FOOT)
fig.text(LEFT, 0.060,
         "AHA Sexual Activity and CVD (Levine et al., Circulation 2012).",
         fontsize=7.5, color=FOOT, style="italic")

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap13_Fig02.pdf"
png_path = out_dir / "_preview_Cap13_Fig02.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
