"""
Cap07 Fig02 (PT-BR, B&W vetorial) — O maior retorno sobre o investimento está no primeiro passo.
Reconstrução pixel-a-pixel do original (1536x1024, aspect 1.5).

Mapeamento (medições pixel-a-pixel):
- Plot ax: [0.103, 0.265, 0.797, 0.435]  (Y-axis em X=159, X-axis em Y=753)
- ylim ≈ (0.7, 5.5)  (marker 5,04 em cy=0.675 detectado)
- Markers cx: 0.183, 0.354, 0.525, 0.696, 0.867
- "Maior ganho ocorre aqui" + bracket FORA do plot (acima), entre Y=223 e Y=307
- "5,04" label INSIDE highlight, just above marker

Fonte: Mandsager et al., JAMA Network Open, 2018.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle
from matplotlib.path import Path as MPath
from matplotlib.patches import PathPatch
import numpy as np

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False
rcParams["pdf.fonttype"] = 42
rcParams["ps.fonttype"] = 42

BG, INK, SOFT, FOOT = "#FFFFFF", "#000000", "#3A3A3A", "#666666"
TICK = "#555555"
CURVE = "#000000"
HILITE = "#EFEFEF"
HL_INK = "#3A3A3A"

_FIG_W, _FIG_H = 11.0, 7.333
fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

# =============== FIGURA 2 TAG ===============
fig.text(0.026, 0.957, "FIGURA 2",
         fontsize=10, color=SOFT, weight="bold", va="center", ha="left")

# =============== TÍTULO + SUBTÍTULO ===============
fig.text(0.026, 0.901, "O maior retorno sobre o investimento está no primeiro passo",
         fontsize=20, color=INK, weight="bold", va="center", ha="left")
fig.text(0.026, 0.845,
         "Risco ajustado de mortalidade por todas as causas de acordo com a aptidão cardiorrespiratória.",
         fontsize=11.5, color=SOFT, va="center", ha="left")

# =============== PLOT AREA (medidas EXATAS pixel-a-pixel) ===============
# Y-axis line X=159 → cx_n=0.103. X-axis line X=159-1358 → cx_n=0.103-0.884 → PL_W=0.781
# X-axis at Y=753 → cy_n=0.265 → PL_BOT=0.265
# y=5.5 tick at cy_n=0.707 → PL_H=0.442
# Categories center cx_n: Baixa=0.170, Abaixo=0.322, Acima=0.487, Alta=0.656, Elite=0.819
# Back-calc xlim: (-0.413, 4.402)
PL_LEFT, PL_BOT = 0.103, 0.265
PL_W, PL_H = 0.781, 0.442
ax = fig.add_axes([PL_LEFT, PL_BOT, PL_W, PL_H])

categories = ["Baixa", "Abaixo da média", "Acima da média", "Alta", "Elite"]
subcats = ["(< P25)", "(P25–P49)", "(P50–P74)", "(P75–P97,6)", "(≥ P97,7)"]
values = [5.04, 2.10, 1.49, 1.29, 1.00]
x_pos = np.arange(5)

ax.set_xlim(-0.413, 4.402)
ax.set_ylim(0.7, 5.5)

# =============== HIGHLIGHT BOX (X=256-472 → axes -0.016 a 0.847) ===============
ax.axvspan(-0.016, 0.847, ymin=0.0, ymax=1.0, facecolor=HILITE,
           edgecolor="none", zorder=0)

# =============== CURVA + MARKERS ===============
ax.plot(x_pos, values, color=CURVE, linewidth=2.2, zorder=5)
ax.scatter(x_pos, values, s=100, facecolor=CURVE, edgecolor=CURVE,
           zorder=6, linewidths=1.0)

# =============== VALUE LABELS — posições EXATAS detectadas no original ===============
# Detectadas (cx_n, cy_n) das char-clusters de cada label:
# "5,04" centro cx=0.170 cy=0.676 (mesma X do marker; label 11px acima)
# "2,10" centro cx=0.331 cy=0.441
# "1,49" centro cx=0.491 cy=0.394
# "1,29" centro cx=0.669 cy=0.363
# "1,00" centro cx=0.862 cy=0.293 (mesma Y do Elite marker em 0.819, label à direita)
value_labels_pos = [
    (0.170, 0.690, "5,04", "center"),
    (0.331, 0.445, "2,10", "center"),
    (0.491, 0.395, "1,49", "center"),
    (0.669, 0.365, "1,29", "center"),
    (0.862, 0.293, "1,00", "center"),
]
for cx, cy, txt, ha in value_labels_pos:
    fig.text(cx, cy, txt, fontsize=12, color=INK, weight="bold",
             va="center", ha=ha, zorder=7)

# =============== DASHED LINE em y=1.41 (tabagismo) ===============
# Spans middle of plot to right edge (axes x≈1.5 a 4.4)
ax.axhline(y=1.41, color=INK, linewidth=1.0, linestyle=(0, (5, 4)),
           xmin=0.40, xmax=0.95, zorder=3)

# =============== Y-AXIS TICKS + SPINES ===============
ax.set_yticks([1.0, 2.0, 3.0, 4.0, 5.0, 5.5])
ax.set_yticklabels(["1,0", "2,0", "3,0", "4,0", "5,0", "5,5"])

ax.set_xticks(x_pos)
ax.set_xticklabels([""] * 5)

ax.spines["top"].set_visible(False)
ax.spines["right"].set_visible(False)
ax.spines["left"].set_color("#555555")
ax.spines["left"].set_linewidth(0.8)
ax.spines["bottom"].set_color("#555555")
ax.spines["bottom"].set_linewidth(0.8)
ax.tick_params(axis="y", colors=TICK, labelsize=11, length=5, width=0.8, direction="out")
ax.tick_params(axis="x", length=0, labelsize=11)

# =============== Y-AXIS LABEL "Risco de morte / (vezes maior)" ===============
fig.text(0.026, 0.776, "Risco de morte",
         fontsize=11, color=INK, weight="bold", va="center", ha="left")
fig.text(0.026, 0.747, "(vezes maior)",
         fontsize=11, color=INK, va="center", ha="left")

# =============== BRACKET (CURLY) + "Maior ganho ocorre aqui" — FORA do plot ===============
# Bracket spans the highlight box width, ABOVE plot top
# Detected green range Y=223-307 → fig_y 0.700-0.782
# Plot top em fig_y=0.700; bracket logo acima.

# X-fig do bracket: corresponde a x=-0.5 (left edge highlight) e x=1.5 (right edge)
# Conversão axes→fig: x_fig = PL_LEFT + (x_axes - (-0.5))/(4.5-(-0.5)) * PL_W
def ax_x_to_fig(x):
    # xlim=(-0.413, 4.402), range=4.815
    return PL_LEFT + (x + 0.413) / 4.815 * PL_W

# Bracket span = highlight span (axes -0.016 a 0.847)
BR_X_LEFT_FIG = ax_x_to_fig(-0.016)
BR_X_RIGHT_FIG = ax_x_to_fig(0.847)
# Bracket subido pra dar respiro ao 5,04
BR_Y = 0.722
BR_CAP_Y = 0.714

# Bracket horizontal
fig.lines.append(plt.Line2D(
    [BR_X_LEFT_FIG, BR_X_RIGHT_FIG], [BR_Y, BR_Y],
    color=HL_INK, linewidth=1.2, transform=fig.transFigure, zorder=8
))
# Caps verticais nas pontas, descem até o topo do plot
fig.lines.append(plt.Line2D(
    [BR_X_LEFT_FIG, BR_X_LEFT_FIG], [BR_Y, BR_CAP_Y],
    color=HL_INK, linewidth=1.2, transform=fig.transFigure, zorder=8
))
fig.lines.append(plt.Line2D(
    [BR_X_RIGHT_FIG, BR_X_RIGHT_FIG], [BR_Y, BR_CAP_Y],
    color=HL_INK, linewidth=1.2, transform=fig.transFigure, zorder=8
))
# Center pin going DOWN (matches original curly bracket style)
BR_CENTER_X = (BR_X_LEFT_FIG + BR_X_RIGHT_FIG) / 2
fig.lines.append(plt.Line2D(
    [BR_CENTER_X, BR_CENTER_X], [BR_Y, BR_Y - 0.006],
    color=HL_INK, linewidth=1.2, transform=fig.transFigure, zorder=8
))

# Texto subido pra dar respiro
TXT_CX = (BR_X_LEFT_FIG + BR_X_RIGHT_FIG) / 2
fig.text(TXT_CX, 0.783, "Maior ganho",
         fontsize=13, color=HL_INK, weight="bold", va="center", ha="center")
fig.text(TXT_CX, 0.748, "ocorre aqui",
         fontsize=13, color=HL_INK, weight="bold", va="center", ha="center")

# =============== RISCO TABAGISMO label ===============
# Ambas linhas ACIMA do dashed (em fig_y=0.330). Right-aligned em fig_x=0.954.
fig.text(0.954, 0.378,
         "Risco aproximado",
         fontsize=10, color=SOFT, va="center", ha="right", zorder=5)
fig.text(0.954, 0.354,
         "do tabagismo (1,41x)",
         fontsize=10, color=SOFT, va="center", ha="right", zorder=5)

# =============== X-AXIS LABELS (2 linhas) ===============
for i, (cat, sub) in enumerate(zip(categories, subcats)):
    cx_fig = PL_LEFT + (i + 0.5) / 5 * PL_W
    fig.text(cx_fig, 0.234, cat,
             fontsize=12, color=INK, weight="bold", va="center", ha="center")
    fig.text(cx_fig, 0.201, sub,
             fontsize=10, color=SOFT, va="center", ha="center")

# =============== X-AXIS TITLE ===============
fig.text(0.500, 0.150,
         "Nível de aptidão cardiorrespiratória (por percentil)",
         fontsize=12, color=INK, weight="bold", va="center", ha="center")

# =============== SEPARATOR + FOOTER ===============
fig.lines.append(plt.Line2D(
    [0.026, 0.974], [0.117, 0.117],
    color="#BBBBBB", linewidth=0.5, transform=fig.transFigure
))
fig.text(0.026, 0.092,
         "A queda mais acentuada ocorre entre o grupo de menor aptidão e o seguinte —",
         fontsize=10, color=FOOT, va="center", ha="left")
fig.text(0.026, 0.068,
         "equivalente a seguir as diretrizes básicas de 150 minutos semanais de atividade moderada.",
         fontsize=10, color=FOOT, va="center", ha="left")
fig.text(0.026, 0.031,
         "Fonte: Mandsager et al., JAMA Network Open, 2018.",
         fontsize=10, color=FOOT, va="center", ha="left")

# =============== EXPORT ===============
out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap07_Fig02.pdf"
png_path = out_dir / "_preview_Cap07_Fig02.png"
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG)
plt.savefig(png_path, dpi=170, facecolor=BG)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
