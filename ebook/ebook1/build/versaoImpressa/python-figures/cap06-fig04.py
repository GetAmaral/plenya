"""
Cap06 Fig04 (PT-BR, B&W vetorial) — O TOTG de André: quando o jejum mente.
Reconstrução pixel-a-pixel do original PNG (1536x1024, aspect 1.5).
Estrutura: FIGURA tag -> TÍTULO em linha própria -> SUBTÍTULO -> LEGENDA
com amostras de linha -> PLOT com axis labels dentro + curvas com anotações
empilhadas no pico de insulina + bracket inferior + labels conclusivos.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, FancyArrowPatch
import numpy as np

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False
rcParams["pdf.fonttype"] = 42
rcParams["ps.fonttype"] = 42

BG, INK, SOFT, FOOT = "#FFFFFF", "#000000", "#3A3A3A", "#666666"
GLI = "#000000"     # azul no original
INS = "#787878"     # vermelho no original
BAND = "#ECECEC"
TICK = "#555555"

_FIG_W, _FIG_H = 11.0, 7.333
fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

# =================== FIGURA 4 TAG (X=0.015-0.130, Y=0.943-0.987) ===================
fig.patches.append(Rectangle(
    (0.015, 0.943), 0.115, 0.044,
    facecolor=INK, edgecolor=INK, transform=fig.transFigure, zorder=10
))
fig.text(0.0725, 0.965, "FIGURA 4",
         fontsize=11, color="white", weight="bold",
         ha="center", va="center", zorder=11)

# =================== TÍTULO em linha própria (cy=0.898) ===================
fig.text(0.028, 0.898, "O TOTG de André: quando o jejum mente.",
         fontsize=20, color=INK, weight="bold", va="center")

# =================== SUBTÍTULO (cy=0.844) ===================
fig.text(0.028, 0.844,
         "Glicose aparentemente normal com resposta insulínica desproporcional.",
         fontsize=11, color=SOFT, va="center")

# =================== LEGENDA (cy=0.785) com amostras de linha ===================
LGD_Y = 0.785
# amostra glicose: linha sólida preta
fig.lines.append(plt.Line2D([0.035, 0.065], [LGD_Y, LGD_Y],
                            color=GLI, linewidth=2.4,
                            transform=fig.transFigure))
fig.text(0.072, LGD_Y, "Glicose (mg/dL)",
         fontsize=10.5, color=INK, weight="bold", va="center")
# amostra insulina: linha tracejada cinza
fig.lines.append(plt.Line2D([0.205, 0.235], [LGD_Y, LGD_Y],
                            color=INS, linewidth=1.8, linestyle=(0, (5, 3)),
                            transform=fig.transFigure))
fig.text(0.242, LGD_Y, "Insulina (µIU/mL)",
         fontsize=10.5, color=INS, weight="bold", va="center")

# =================== PLOT AREA  X=[0.079, 0.909]  Y=[0.206, 0.745] ===================
PL_LEFT, PL_BOT = 0.079, 0.206
PL_W, PL_H = 0.830, 0.539
ax1 = fig.add_axes([PL_LEFT, PL_BOT, PL_W, PL_H])
ax2 = ax1.twinx()

x = np.array([0, 30, 60, 90, 120])
glicose  = np.array([92, 148, 162, 154, 131])
insulina = np.array([8.5, 78, 124, 118, 89])

ax1.set_xlim(-10, 130)
ax1.set_ylim(70, 195)
ax2.set_ylim(-10, 200)

# =================== AXIS LABELS dentro do plot (top-left/top-right) ===================
# Em fig coords pra precisão das posições detectadas
fig.text(0.031, 0.732, "Glicose (mg/dL)",
         fontsize=10.5, color=INK, weight="bold", va="center", ha="left")
fig.text(0.953, 0.740, "Insulina (µIU/mL)",
         fontsize=10.5, color=INS, weight="bold", va="center", ha="right")

# =================== BANDA FAIXA NORMAL (glicose 70-140) ===================
ax1.axhspan(70, 140, facecolor=BAND, edgecolor="none", zorder=0)

# =================== CURVAS ===================
# Glicose: linha sólida + círculos cheios pretos
ax1.plot(x, glicose, color=GLI, linewidth=2.4, zorder=5)
ax1.scatter(x, glicose, s=70, facecolor=GLI, edgecolor=GLI,
            zorder=6, linewidths=1.0)

# Insulina: dashed cinza + círculos. Pico (t=60) é hollow.
ax2.plot(x, insulina, color=INS, linewidth=1.8, linestyle=(0, (5, 3)), zorder=5)
ins_face = ["white" if xv == 60 else INS for xv in x]
ins_edge_w = [2.0 if xv == 60 else 1.0 for xv in x]
ins_size = [110 if xv == 60 else 55 for xv in x]
ax2.scatter(x, insulina, s=ins_size, facecolor=ins_face, edgecolor=INS,
            zorder=6, linewidths=ins_edge_w)

# =================== EIXOS ===================
ax1.set_xticks([0, 30, 60, 90, 120])
ax1.set_yticks([120, 140, 160])
ax2.set_yticks([40, 60, 80, 100])
ax1.set_xlabel("Tempo (min)", fontsize=11, color=INK, weight="bold", labelpad=8)
ax1.tick_params(axis="both", colors=TICK, labelsize=10)
ax2.tick_params(axis="y", colors=TICK, labelsize=10)
ax1.spines["top"].set_visible(False); ax2.spines["top"].set_visible(False)
for s in ("left", "bottom"): ax1.spines[s].set_color("#888888")
ax2.spines["right"].set_color("#888888")
ax2.spines["left"].set_visible(False)

# =================== "Faixa normal pós-prandial (≤ 140 mg/dL)" 3 linhas italic ===================
fig.text(0.085, 0.483, "Faixa normal",
         fontsize=10, color=SOFT, ha="left", va="center", style="italic")
fig.text(0.085, 0.458, "pós-prandial",
         fontsize=10, color=SOFT, ha="left", va="center", style="italic")
fig.text(0.085, 0.434, "(≤ 140 mg/dL)",
         fontsize=10, color=SOFT, ha="left", va="center", style="italic")

# =================== VALORES GLICOSE — pos detectadas (cx, cy_top) ===================
glicose_labels = [
    (0.145, 0.354, "92"),
    (0.319, 0.604, "148"),
    (0.498, 0.675, "162"),
    (0.676, 0.641, "154"),
    (0.846, 0.497, "131"),
]
for cx, cy_top, txt in glicose_labels:
    fig.text(cx, cy_top, txt, fontsize=14, color=INK, weight="bold",
             ha="center", va="bottom")

# "(pico)" italic ABAIXO de "162" (empilhado, NÃO ao lado)
fig.text(0.498, 0.640, "(pico)", fontsize=10, color=INK, style="italic",
         ha="center", va="top")

# =================== VALORES INSULINA (cinza) abaixo dos markers ===================
# t=30 marker em cy=0.436; t=90 em 0.560; t=120 em 0.429
fig.text(0.321, 0.413, "78", fontsize=12, color=INS, weight="bold",
         ha="center", va="top")
fig.text(0.676, 0.537, "118", fontsize=12, color=INS, weight="bold",
         ha="center", va="top")
fig.text(0.846, 0.405, "89", fontsize=12, color=INS, weight="bold",
         ha="center", va="top")

# =================== ANOTAÇÕES PICO INSULINA — STACK abaixo do hollow marker ===================
# Hollow marker em fig cy ≈ 0.553
# Pequena seta apontando UP pra marker
fig.patches.append(FancyArrowPatch(
    (0.498, 0.535), (0.498, 0.555),
    arrowstyle="-|>", color=INS, lw=1.0, mutation_scale=8,
    transform=fig.transFigure, zorder=8,
))
fig.text(0.498, 0.520, "124 µIU/mL", fontsize=11, color=INS, weight="bold",
         ha="center", va="top")
fig.text(0.498, 0.496, "(pico)", fontsize=10, color=INS, style="italic",
         ha="center", va="top")
fig.text(0.498, 0.470, "Kraft II:", fontsize=10, color=INK,
         ha="center", va="top")
fig.text(0.498, 0.448, "hiperinsulinemia", fontsize=10, color=INK,
         ha="center", va="top")
fig.text(0.498, 0.426, "compensatória.", fontsize=10, color=INK,
         ha="center", va="top")

# =================== BRACKET inferior + labels (t=60 → t=120) ===================
# Bracket horizontal em cy=0.395
BR_Y = 0.395
BR_X1 = 0.498  # x_norm de t=60
BR_X2 = 0.848  # x_norm de t=120
fig.lines.append(plt.Line2D([BR_X1, BR_X2], [BR_Y, BR_Y],
                            color=INS, linewidth=1.8, transform=fig.transFigure))
fig.lines.append(plt.Line2D([BR_X1, BR_X1], [BR_Y, BR_Y+0.012],
                            color=INS, linewidth=1.8, transform=fig.transFigure))
fig.lines.append(plt.Line2D([BR_X2, BR_X2], [BR_Y, BR_Y+0.012],
                            color=INS, linewidth=1.8, transform=fig.transFigure))

# "Glicose moderada." (azul → preto)
fig.text(0.560, 0.362, "Glicose moderada.",
         fontsize=12, color=INK, weight="bold", ha="left", va="center")
# "Insulina desproporcional." (vermelho → cinza)
fig.text(0.560, 0.335, "Insulina desproporcional.",
         fontsize=12, color=INS, weight="bold", ha="left", va="center")

# =================== ANOTAÇÃO "8,5 µIU/mL" com arrow → t=0 insulin marker ===================
# Marker t=0 insulin em fig (0.140, 0.260). Texto à direita.
fig.text(0.205, 0.272, "8,5 µIU/mL",
         fontsize=10, color=INS, weight="bold", ha="left", va="center")
fig.text(0.205, 0.247, "\"normal\" pelo",
         fontsize=9.5, color=INK, ha="left", va="center")
fig.text(0.205, 0.225, "laboratório",
         fontsize=9.5, color=INK, ha="left", va="center")
# arrow da text esq pro marker
fig.patches.append(FancyArrowPatch(
    (0.200, 0.272), (0.155, 0.262),
    arrowstyle="-|>", color=INS, lw=1.0, mutation_scale=8,
    transform=fig.transFigure, zorder=8,
))

# =================== FOOTER (cy=0.087, cy=0.060) ===================
fig.text(0.033, 0.087,
         "Fonte: dados reconstruídos do caso-tipo apresentado no Capítulo 6.",
         fontsize=9, color=FOOT, va="center")
fig.text(0.033, 0.060,
         "Padrão de referência: Kraft JR (Detection of Diabetes Mellitus In Situ, 2008).",
         fontsize=9, color=FOOT, va="center")

# =================== EXPORT ===================
out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap06_Fig04.pdf"
png_path = out_dir / "_preview_Cap06_Fig04.png"
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
