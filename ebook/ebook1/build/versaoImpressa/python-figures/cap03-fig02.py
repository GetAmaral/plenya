"""
Cap03 Fig02 (PT-BR, B&W vetorial) — O Que Acelera e O Que Freia o Envelhecimento.

Duas colunas simétricas com coluna central indicando velocidade biológica.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, FancyArrowPatch

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
BAND     = "#EDEDED"
COL_BG   = "#F4F4F4"

fig = plt.figure(figsize=(11.0, 7.5))
fig.patch.set_facecolor(BG)

LEFT = 0.025

# Título
fig.text(LEFT, 0.945, "Figura 2 — O Que Acelera e O Que Freia o Envelhecimento",
         fontsize=17, color=INK, weight="bold", va="center")

# ---------- 3 colunas ----------
LEFT_X1, LEFT_X2 = 0.04, 0.38       # coluna FREIA
MID_X1, MID_X2   = 0.40, 0.60       # coluna velocidade
RIGHT_X1, RIGHT_X2 = 0.62, 0.96     # coluna ACELERA

# Headers — com sublinhado decorativo abaixo (cor da coluna no original; preto em B&W)
HEADER_Y = 0.85
# Esquerda
hx_l = (LEFT_X1 + LEFT_X2) / 2
fig.text(hx_l, HEADER_Y, "O QUE FREIA",
         fontsize=14, color=INK, weight="bold", ha="center")
fig.add_artist(plt.Line2D(
    [hx_l - 0.060, hx_l + 0.060], [HEADER_Y - 0.025, HEADER_Y - 0.025],
    color=INK, linewidth=1.4, transform=fig.transFigure, zorder=3
))
# Direita
hx_r = (RIGHT_X1 + RIGHT_X2) / 2
fig.text(hx_r, HEADER_Y, "O QUE ACELERA",
         fontsize=14, color=INK, weight="bold", ha="center")
fig.add_artist(plt.Line2D(
    [hx_r - 0.065, hx_r + 0.065], [HEADER_Y - 0.025, HEADER_Y - 0.025],
    color=INK, linewidth=1.4, transform=fig.transFigure, zorder=3
))

# Centro: velocidade do envelhecimento biológico
fig.text((MID_X1 + MID_X2) / 2, HEADER_Y, "Velocidade do",
         fontsize=9, color=INK_SOFT, ha="center", weight="bold")
fig.text((MID_X1 + MID_X2) / 2, HEADER_Y - 0.025,
         "envelhecimento biológico",
         fontsize=9, color=INK_SOFT, ha="center", weight="bold")

# Sub-rótulos do centro
fig.text(MID_X1 + 0.04, HEADER_Y - 0.055, "MAIS LENTA",
         fontsize=8, color=INK_SOFT, ha="center", weight="bold")
fig.text(MID_X2 - 0.04, HEADER_Y - 0.055, "MAIS RÁPIDA",
         fontsize=8, color=INK_SOFT, ha="center", weight="bold")

# Setas direcionais
fig.patches.append(FancyArrowPatch(
    (MID_X1 + 0.05, HEADER_Y - 0.075), (MID_X1 + 0.02, HEADER_Y - 0.075),
    arrowstyle="->", color=INK_SOFT, lw=1.0, mutation_scale=10,
    transform=fig.transFigure
))
fig.patches.append(FancyArrowPatch(
    (MID_X2 - 0.05, HEADER_Y - 0.075), (MID_X2 - 0.02, HEADER_Y - 0.075),
    arrowstyle="->", color=INK_SOFT, lw=1.0, mutation_scale=10,
    transform=fig.transFigure
))

# Caixa central com GRADIENTE horizontal (claro=LENTA esquerda, escuro=RÁPIDA direita)
COL_Y1, COL_Y2 = 0.18, HEADER_Y - 0.090
# Gradiente via N bandas verticais finas
import numpy as np
N_BANDS = 80
for k in range(N_BANDS):
    frac = k / (N_BANDS - 1)
    # Interpola entre claro (#F5F5F5) e escuro (#999999)
    g = int(245 - frac * (245 - 153))  # 245 → 153
    col = f"#{g:02X}{g:02X}{g:02X}"
    bx0 = MID_X1 + (MID_X2 - MID_X1) * (k / N_BANDS)
    bw  = (MID_X2 - MID_X1) / N_BANDS + 0.001
    fig.patches.append(Rectangle(
        (bx0, COL_Y1), bw, COL_Y2 - COL_Y1,
        facecolor=col, edgecolor="none", transform=fig.transFigure, zorder=1
    ))
# Texto central da coluna
fig.text((MID_X1 + MID_X2) / 2, (COL_Y1 + COL_Y2) / 2,
         "Processos biológicos\nmoduláveis",
         fontsize=9.5, color=INK_SOFT, ha="center", va="center",
         style="italic", linespacing=1.3)

# ---------- 5 linhas de itens ----------
items_left = [
    ("Exercício aeróbio",      "Mitocôndrias ↑, telômeros ↑"),
    ("Treino de força",        "Inflamação ↓, insulina ↓"),
    ("Sono de qualidade",      "Limpeza celular ↑, epigenoma estável"),
    ("Alimentação real",       "Estresse oxidativo ↓, epigenoma modula"),
    ("Gestão do estresse",     "Telômeros ↑, inflamação ↓"),
]
items_right = [
    ("Sedentarismo",           "Mitocôndrias ↓, senescência ↑"),
    ("Ultraprocessados",       "Inflamação ↑, estresse oxidativo ↑"),
    ("Sono ruim / irregular",  "Limpeza celular ↓, epigenoma desregulado"),
    ("Estresse crônico",       "Telômeros ↓, cortisol ↑"),
    ("Gordura visceral",       "Inflamação ↑, insulina ↑"),
]

ROW_TOP = HEADER_Y - 0.120
ROW_BOTTOM = 0.20
ROW_SPACE = (ROW_TOP - ROW_BOTTOM) / (len(items_left) - 1)

for i, ((name_l, desc_l), (name_r, desc_r)) in enumerate(zip(items_left, items_right)):
    y = ROW_TOP - i * ROW_SPACE

    # === Coluna LEFT ===
    fig.text(LEFT_X1 + 0.010, y, name_l,
             fontsize=11.5, color=INK, weight="bold",
             ha="left", va="center")
    # Sublinhado curto decorativo abaixo do título (cor da coluna no original)
    # Largura ~ proporcional ao texto
    underline_w = 0.005 * len(name_l) * 0.7
    fig.add_artist(plt.Line2D(
        [LEFT_X1 + 0.010, LEFT_X1 + 0.010 + underline_w],
        [y - 0.015, y - 0.015],
        color=INK, linewidth=1.0, transform=fig.transFigure, zorder=3
    ))
    fig.text(LEFT_X1 + 0.010, y - 0.030, desc_l,
             fontsize=9, color=INK_SOFT,
             ha="left", va="center", style="italic")

    # === Coluna RIGHT ===
    fig.text(RIGHT_X2 - 0.010, y, name_r,
             fontsize=11.5, color=INK, weight="bold",
             ha="right", va="center")
    # Sublinhado curto decorativo abaixo do título (alinhado à direita)
    underline_w = 0.005 * len(name_r) * 0.7
    fig.add_artist(plt.Line2D(
        [RIGHT_X2 - 0.010 - underline_w, RIGHT_X2 - 0.010],
        [y - 0.015, y - 0.015],
        color=INK, linewidth=1.0, transform=fig.transFigure, zorder=3
    ))
    fig.text(RIGHT_X2 - 0.010, y - 0.030, desc_r,
             fontsize=9, color=INK_SOFT,
             ha="right", va="center", style="italic")

# ---------- caixa inferior ----------
BOX_X1, BOX_X2 = 0.04, 0.96
BOX_Y1, BOX_Y2 = 0.07, 0.14
fig.patches.append(Rectangle(
    (BOX_X1, BOX_Y1), BOX_X2 - BOX_X1, BOX_Y2 - BOX_Y1,
    facecolor=BAND, edgecolor="none", transform=fig.transFigure, zorder=1
))
fig.text(0.5, (BOX_Y1 + BOX_Y2) / 2,
         "Envelhecimento não é destino. É um processo modificável.",
         fontsize=12, color=INK, weight="bold",
         ha="center", va="center")

# ---------- save ----------
out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap03_Fig02.pdf"
png_path = out_dir / "_preview_Cap03_Fig02.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
