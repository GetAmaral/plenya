"""
Cap11 Fig01 (PT-BR, B&W vetorial) — Paulo: 6 meses sem reposição — e a trajetória mudou.
Dot plot com setas de movimento (antes → depois).
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, Ellipse, FancyArrowPatch

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
BAND_OK  = "#F4F4F4"
BAND_MID = "#E0E0E0"
BAND_BAD = "#BFBFBF"
BAND     = "#EDEDED"

_FIG_W, _FIG_H = 11.0, 7.5
_ASPECT = _FIG_W / _FIG_H

# Cada biomarker: (nome, unidade, valor antes, valor depois, alvo,
#                  pos antes, pos depois, pos alvo)
biomarkers = [
    ("TESTOSTERONA\nTOTAL", "(ng/dL)",   "310",  "485",  "> 500", 0.18, 0.55, 0.62),
    ("TESTOSTERONA\nLIVRE", "(pg/mL)",   "4,8",  "11,2", "> 10",  0.20, 0.62, 0.55),
    ("VITAMINA D\n(25-OH)", "(ng/mL)",   "24",   "58",   "40-60", 0.22, 0.66, 0.55),
    ("hs-CRP",              "(mg/L)",    "1,7",  "0,9",  "< 1,0", 0.62, 0.35, 0.30),
    ("IDADE EPIGENÉTICA",   "(anos vs. cronológica)", "+4", "+2", "0", 0.78, 0.55, 0.25),
]

fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

LEFT_MARGIN  = 0.025
BAR_LEFT     = 0.22
BAR_RIGHT    = 0.83
ALVO_X       = 0.91

# Título
fig.text(LEFT_MARGIN, 0.945,
         "Figura 1 — Paulo: 6 meses sem reposição — e a trajetória mudou.",
         fontsize=15, color=INK, weight="bold")
fig.text(LEFT_MARGIN, 0.910,
         "Otimização de sono, vitamina D e treino de força ajustado — sem uso de reposição de testosterona.",
         fontsize=9.5, color=INK_SOFT, style="italic")

# Legenda topo
LEG_Y = 0.860
def small_circle(x, y, fill, edge=INK):
    r = 0.0085
    fig.patches.append(Ellipse(
        (x, y), width=r*2, height=r*2*_ASPECT,
        facecolor=fill, edgecolor=edge, linewidth=1.0,
        transform=fig.transFigure, zorder=4
    ))

small_circle(LEFT_MARGIN + 0.005, LEG_Y, "white")
fig.text(LEFT_MARGIN + 0.020, LEG_Y, "Antes (baseline)",
         fontsize=9, color=INK, weight="bold", va="center")

small_circle(LEFT_MARGIN + 0.160, LEG_Y, INK)
fig.text(LEFT_MARGIN + 0.175, LEG_Y, "Depois (6 meses)",
         fontsize=9, color=INK, weight="bold", va="center")

fig.text(LEFT_MARGIN + 0.305, LEG_Y, "▲  Alvo ótimo",
         fontsize=9, color=INK_SOFT, va="center")

# Header da coluna direita
fig.text(ALVO_X, 0.860, "Alvo ótimo",
         fontsize=8.5, color=TICK, style="italic", ha="center")

# Linhas de biomarcadores
ROW_TOP = 0.785
ROW_BOTTOM = 0.230
ROW_SPACE = (ROW_TOP - ROW_BOTTOM) / (len(biomarkers) - 1)
BAR_HEIGHT = 0.030

for i, (name, unit, antes, depois, alvo,
        pos_antes, pos_depois, pos_alvo) in enumerate(biomarkers):
    y = ROW_TOP - i * ROW_SPACE
    bar_w = BAR_RIGHT - BAR_LEFT

    # Nome
    fig.text(LEFT_MARGIN, y + 0.008, name,
             fontsize=10, color=INK, weight="bold",
             va="center", linespacing=1.15)
    fig.text(LEFT_MARGIN, y - 0.030, unit,
             fontsize=7.5, color=TICK, va="center")

    # Barra de fundo (3 zonas)
    # Pra simplificar: gradient simples — claro à esquerda, escuro à direita
    # representando "fora do ótimo" → "no ótimo" ou vice-versa
    fig.patches.extend([
        Rectangle((BAR_LEFT, y - BAR_HEIGHT/2), bar_w * 0.5, BAR_HEIGHT,
                  facecolor=BAND_BAD, edgecolor="none",
                  transform=fig.transFigure, zorder=1),
        Rectangle((BAR_LEFT + bar_w * 0.5, y - BAR_HEIGHT/2),
                  bar_w * 0.25, BAR_HEIGHT,
                  facecolor=BAND_MID, edgecolor="none",
                  transform=fig.transFigure, zorder=1),
        Rectangle((BAR_LEFT + bar_w * 0.75, y - BAR_HEIGHT/2),
                  bar_w * 0.25, BAR_HEIGHT,
                  facecolor=BAND_OK, edgecolor="none",
                  transform=fig.transFigure, zorder=1),
    ])

    x_antes  = BAR_LEFT + bar_w * pos_antes
    x_depois = BAR_LEFT + bar_w * pos_depois
    x_alvo   = BAR_LEFT + bar_w * pos_alvo

    # Seta antes → depois
    fig.patches.append(FancyArrowPatch(
        (x_antes, y - 0.022), (x_depois, y - 0.022),
        arrowstyle="->", color=INK, lw=1.5, mutation_scale=12,
        transform=fig.transFigure, zorder=4
    ))

    # Marker ANTES (círculo branco)
    small_circle(x_antes, y, "white")
    fig.text(x_antes, y + 0.030, antes,
             fontsize=9, color=INK_SOFT, weight="bold", ha="center", va="bottom")

    # Marker DEPOIS (círculo preto)
    small_circle(x_depois, y, INK)
    fig.text(x_depois, y + 0.030, depois,
             fontsize=10, color=INK, weight="bold", ha="center", va="bottom")

    # Marker ALVO (triângulo)
    fig.text(x_alvo, y, "▲", fontsize=11, color=INK,
             ha="center", va="center", zorder=5)

    # Valor alvo na coluna direita
    fig.text(ALVO_X, y, alvo,
             fontsize=10, color=INK, weight="bold",
             ha="center", va="center")

# ---------- caixa final ----------
BOX_X1, BOX_X2 = 0.04, 0.96
BOX_Y1, BOX_Y2 = 0.090, 0.155
fig.patches.append(Rectangle(
    (BOX_X1, BOX_Y1), BOX_X2 - BOX_X1, BOX_Y2 - BOX_Y1,
    facecolor=BAND, edgecolor=INK, linewidth=0.5,
    transform=fig.transFigure, zorder=1
))
fig.text(0.5, (BOX_Y1 + BOX_Y2) / 2 + 0.010,
         "Diagnóstico correto.",
         fontsize=11, color=INK, weight="bold", ha="center", va="center")
fig.text(0.5, (BOX_Y1 + BOX_Y2) / 2 - 0.012,
         "Plano integrado e seguimento.",
         fontsize=10, color=INK, ha="center", va="center", style="italic")

# Footer
fig.text(LEFT_MARGIN, 0.055,
         "DHEA-S aumentou de 145 para 210 µg/dL no período (não exibido no gráfico para manter a clareza visual).",
         fontsize=7.5, color=FOOT, style="italic")

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap11_Fig01.pdf"
png_path = out_dir / "_preview_Cap11_Fig01.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
