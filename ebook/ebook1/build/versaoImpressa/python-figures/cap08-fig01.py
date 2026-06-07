"""
Cap08 Fig01 (PT-BR, B&W vetorial) — Marcos, 8 meses depois.

Reconstrução fiel do original 2048x1152 via metodologia CROP-GRID 400x400 +
medição pixel-precisa por amostragem de máscara colorida (sample_positions.py).

Layout: bloco left = bars BASELINE vs 8 MESES; bloco right = chart hazard vs aptidão.
Cores originais (red + slate) convertidas pra escala cinza B&W mantendo separação visual.
"""
from pathlib import Path as _Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, FancyArrowPatch, Circle
from matplotlib.path import Path as MplPath
from matplotlib.patches import PathPatch
import numpy as np

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False
rcParams["pdf.fonttype"] = 42
rcParams["ps.fonttype"] = 42

# ---------- paleta B&W ----------
BG    = "#FFFFFF"
INK   = "#000000"
SOFT  = "#3A3A3A"
FOOT  = "#666666"
GRAY  = "#BDBDBD"   # BASELINE bar e elementos "antes" (era slate gray)
GRAY2 = "#8E8E8E"   # ring outline, divisor

# ---------- figura px-coord ----------
W_IMG, H_IMG = 2048, 1152
_FIG_W, _FIG_H = 12.0, 12.0 * H_IMG / W_IMG   # 12 x 6.75 in
fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

# Eixo full-figure em pixels da imagem original (origem topo-esq)
ax = fig.add_axes([0, 0, 1, 1])
ax.set_xlim(0, W_IMG)
ax.set_ylim(H_IMG, 0)
ax.set_aspect("equal")
ax.axis("off")

# px -> pt (figura 12in = 2048px logical → 1px = 12/2048 in = 0.422 pt)
def px_pt(px): return px * (_FIG_W * 72.0) / W_IMG / 1.6  # /1.6 = cap-height vs em scaling

# ============================================================
# CABEÇALHO
# ============================================================
ax.text(87, 78, "FIGURA 1",
        fontsize=11, color=SOFT, weight="bold", va="center", ha="left")

ax.text(87, 154,
        "Marcos, 8 meses depois: a aptidão que vale por uma estatina.",
        fontsize=22, color=INK, weight="bold", va="center", ha="left")

ax.text(87, 290, "VO₂ MAX  ·  ERGOMETRIA",
        fontsize=12, color=SOFT, weight="bold", va="center", ha="left", alpha=0.85)

ax.text(1130, 290, "RISCO DE MORTE  ×  APTIDÃO",
        fontsize=12, color=SOFT, weight="bold", va="center", ha="left", alpha=0.85)

# ============================================================
# BLOCO ESQUERDA — BARRAS
# ============================================================
# Bar GRAY (BASELINE) — px (264,413)..(683,482)
ax.add_patch(Rectangle((264, 413), 683-264, 482-413,
                       facecolor=GRAY, edgecolor="none"))
# Bar BLACK (8 MESES, era vermelha) — px (261,659)..(748,730)
ax.add_patch(Rectangle((261, 659), 748-261, 730-659,
                       facecolor=INK, edgecolor="none"))

# Labels esquerda
ax.text(87, 425, "BASELINE",
        fontsize=13, color=INK, weight="bold", va="center", ha="left")
ax.text(87, 678, "8 MESES",
        fontsize=13, color=INK, weight="bold", va="center", ha="left")

# Valores: "28 ml/kg/min" e "34,3 ml/kg/min"
# centro de cada bar height
# "28" grande + "ml/kg/min" pequeno embaixo/lado
ax.text(720, 442, "28",
        fontsize=22, color=INK, weight="bold", va="center", ha="left")
ax.text(790, 452, "ml/kg/min",
        fontsize=10, color=SOFT, va="center", ha="left")

ax.text(770, 690, "34,3",
        fontsize=22, color=INK, weight="bold", va="center", ha="left")
ax.text(885, 700, "ml/kg/min",
        fontsize=10, color=SOFT, va="center", ha="left")

# +1,8 MET — destaque central entre as barras
ax.text(370, 580, "+1,8 MET",
        fontsize=42, color=INK, weight="bold", va="center", ha="center",
        family="sans-serif")

# ============================================================
# BLOCO DIREITA — DOT PLOT (HR × APTIDÃO)
# ============================================================
# Separador vertical entre bloco esquerda e bloco direita (px x=1023, y=248..857)
ax.plot([1023, 1023], [248, 857], color=GRAY2, linewidth=0.7, alpha=0.65, zorder=1)

# 5 dots filled (curva tracejada): BAIXA, ABAIXO, ACIMA, ALTA, ELITE
DOTS = [
    (1189.7, 366.5),   # BAIXA  (HR 5.04)
    (1356.8, 473.2),   # ABAIXO
    (1570.6, 573.3),   # ACIMA
    (1743.6, 638.7),   # ALTA
    (1924.9, 691.4),   # ELITE  (HR 1.00)
]

# Linha tracejada conectando os 5 dots (curva monotônica decrescente)
xs = np.array([d[0] for d in DOTS])
ys = np.array([d[1] for d in DOTS])
# Interpolação suave: usar curva quadrática que passe pelos pontos
# Como são 5 pontos quase em curva, plot direto com linestyle dashed em segmentos
# Mas pra ficar smooth como o original (Bezier suave), usar interp
from numpy.polynomial import polynomial as P
# Spline simples: usar np.interp por segmentos densificados via curva cúbica
xs_dense = np.linspace(xs[0], xs[-1], 400)
# Ajustar polinômio grau 3 aos 5 pontos
coef = np.polyfit(xs, ys, 3)
ys_dense = np.polyval(coef, xs_dense)
ax.plot(xs_dense, ys_dense, linestyle=(0, (8, 5)),
        color=INK, linewidth=1.6, alpha=0.85, zorder=3)

# Dots cheios (preto)
for (x, y) in DOTS:
    ax.add_patch(Circle((x, y), 12.5, facecolor=INK, edgecolor="none", zorder=5))

# Hollow ring (BAIXA, "baseline" Marcos) — center (1192, 623)
ax.add_patch(Circle((1192, 623), 15.5, facecolor=BG,
                    edgecolor=INK, linewidth=2.2, zorder=5))

# Red filled dot ACIMA "8 meses" Marcos — preto sólido maior + halo branco
# center (1565.8, 618.2)
ax.add_patch(Circle((1566, 618), 22.5, facecolor=BG, edgecolor="none", zorder=5))
ax.add_patch(Circle((1566, 618), 19.5, facecolor=INK, edgecolor="none", zorder=6))

# Arrow: hollow → red dot. Curva arc UP, apex em (1370, 565), inicio (1206, 613), fim (1546, 605)
arrow = FancyArrowPatch(
    (1212, 615), (1543, 612),
    connectionstyle="arc3,rad=-0.32",
    arrowstyle="-|>,head_length=14,head_width=10",
    color=INK, linewidth=2.0, zorder=7, capstyle="round",
)
ax.add_patch(arrow)

# Y-axis labels "5,04" e "1,00"
ax.text(1155, 368, "5,04",
        fontsize=12, color=INK, weight="bold", va="center", ha="right")
ax.text(1155, 700, "1,00",
        fontsize=12, color=INK, weight="bold", va="center", ha="right")

# X-axis labels
X_TICKS = [
    (1199, "BAIXA"),
    (1374, "ABAIXO"),
    (1571, "ACIMA"),
    (1750, "ALTA"),
    (1893, "ELITE"),
]
for x, label in X_TICKS:
    ax.text(x, 765, label,
            fontsize=11, color=INK, weight="bold", va="center", ha="center")

# Sublabels: "baseline" sob BAIXA, "8 meses" sob ACIMA
ax.text(1199, 800, "baseline",
        fontsize=9, color=SOFT, va="center", ha="center", style="italic")
ax.text(1571, 800, "8 meses",
        fontsize=9, color=SOFT, va="center", ha="center", style="italic")

# Leaders verticais pontilhados — saem do hollow ring (x=1193) e do red dot (x=1571)
# Dois segmentos por leader (interrompidos pelo label principal):
#   seg1: y=642..725 (logo abaixo do dot/ring → topo do label)
#   seg2: y=771..791 (entre label e sublabel)
LEADER_STYLE = dict(color=GRAY2, linewidth=1.0, linestyle=(0, (1.5, 2.5)),
                    alpha=0.9, zorder=2)
for lead_x in (1193, 1571):
    ax.plot([lead_x, lead_x], [642, 725], **LEADER_STYLE)
    ax.plot([lead_x, lead_x], [771, 791], **LEADER_STYLE)

# ============================================================
# RODAPÉ
# ============================================================
# Divider line
ax.plot([87, 1960], [912, 912], color=GRAY2, linewidth=0.8, zorder=2)

# Resultado: "+1,8 MET = 25–30% menos risco de morte por todas as causas."
ax.text(1024, 980,
        "+1,8 MET  =  25–30% menos risco de morte por todas as causas.",
        fontsize=15, color=INK, weight="bold", va="center", ha="center")

# Fonte
ax.text(1024, 1035,
        "Mandsager et al., JAMA Network Open 2018.",
        fontsize=10, color=FOOT, va="center", ha="center", style="italic")

# ============================================================
# EXPORT
# ============================================================
out_dir = _Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap08_Fig01.pdf"
png_path = out_dir / "_preview_Cap08_Fig01.png"
plt.savefig(pdf_path, facecolor=BG, bbox_inches="tight", pad_inches=0.0)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches="tight", pad_inches=0.0)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
