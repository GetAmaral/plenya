"""
Cap07 Fig03 (PT-BR, B&W vetorial) — Os 4 pilares do exercício para longevidade (na prática).
Donut chart com 4 segmentos. Layout match original:
  - Zona 2 ocupa metade esquerda + topo esquerdo
  - Força ocupa topo direito
  - HIIT direita
  - Mobilidade base
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Circle, Wedge
import numpy as np

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

SHADE_ZONA2  = "#4A4A4A"
SHADE_FORCA  = "#7A7A7A"
SHADE_HIIT   = "#A8A8A8"
SHADE_MOB    = "#D0D0D0"

fig = plt.figure(figsize=(11.0, 8.0))
fig.patch.set_facecolor(BG)

fig.text(0.025, 0.945, "Figura 3 — Os 4 pilares do exercício para longevidade (na prática)",
         fontsize=15, color=INK, weight="bold")
fig.text(0.025, 0.905,
         "Um sistema integrado — cada pilar tem uma função específica.",
         fontsize=10, color=INK_SOFT, style="italic")

ax = fig.add_axes([0.22, 0.20, 0.56, 0.66])
ax.set_xlim(-1.7, 1.7)
ax.set_ylim(-1.4, 1.4)
ax.set_aspect("equal")
ax.axis("off")

# Wedges definidos por (theta1, theta2) onde counterclockwise vai de theta1 → theta2
# Ordem visual no original (sentido horário começando do topo):
#   Topo-esquerda ao topo-direita: Zona 2 (parte superior esquerda)
#   Topo-direita ao direita: Força
#   Direita ao baixo-direita: HIIT
#   Baixo: Mobilidade
#   Baixo-esquerda ao topo-esquerda: Zona 2 (continua)
segments = [
    # (theta1, theta2, color, pct, title)
    (-90, 70,     SHADE_ZONA2, "50–60%", "1. ZONA 2"),       # esquerda inteira
    (70, 165,     SHADE_FORCA, "25–30%", "2. FORÇA"),        # topo (cruzando 90°)
    # Hmm needs different approach
]

# Vamos refazer com angulos finais (CCW pra matplotlib Wedge):
# Pizza começa em 90° (topo) — em CCW positivo (esquerda)
# Original tem:
#  - Zona 2 (55%): da posição "9h" até "1h" passando pela "11h" (esquerda + topo-esquerda)
#  - Força (28%): "1h" até "3h" (topo-direita + direita-alta)
#  - HIIT (12%): "3h" até "4h30" (direita-baixa)
#  - Mobilidade (5%): "4h30" até "6h" (baixo)
#  - depois o gap fecha de 6h ate 9h vazio? Não — todo donut é preenchido
#
# OK, em ângulos matplotlib (0°=leste, 90°=norte, 180°=oeste):
#  - Mobilidade: 270° (sul) — vai de 270° a 290° (~5%)
#  - HIIT:       290° a 333° (~12%)
#  - Força:      333° a 70° (~27%, passa pelo 0°)
#  - Zona 2:     70° a 270° (~55%, lado esquerdo + topo-esquerdo)

segments = [
    ("Zona 2",     70,  270, SHADE_ZONA2, "50–60%", "1. ZONA 2",     "ccw"),
    ("Força",      333, 70,  SHADE_FORCA, "25–30%", "2. FORÇA",      "ccw_wrap"),
    ("HIIT",       290, 333, SHADE_HIIT,  "10–15%", "3. HIIT",       "ccw"),
    ("Mobilidade", 270, 290, SHADE_MOB,   "5–10%",  "4. MOBILIDADE", "ccw"),
]

R_OUT   = 1.05
R_WIDTH = 0.45

for name, t1, t2, color, pct, title, mode in segments:
    # Wedge counterclockwise from t1 to t2
    if mode == "ccw_wrap":
        # passa pelo 0° (ex.: 333° → 70° = 333° → 360° → 70°)
        w = Wedge((0, 0), R_OUT, t1, t2 + 360 if t2 < t1 else t2,
                  width=R_WIDTH, facecolor=color, edgecolor="white", linewidth=2.5)
    else:
        w = Wedge((0, 0), R_OUT, t1, t2,
                  width=R_WIDTH, facecolor=color, edgecolor="white", linewidth=2.5)
    ax.add_patch(w)

    # Ângulo médio para posicionar o texto
    if mode == "ccw_wrap":
        mid = (t1 + (t2 + 360 if t2 < t1 else t2)) / 2
        mid = mid % 360
    else:
        mid = (t1 + t2) / 2
    rad_mid = R_OUT - R_WIDTH / 2
    tx = rad_mid * np.cos(np.radians(mid))
    ty = rad_mid * np.sin(np.radians(mid))

    text_color = INK if color == SHADE_MOB else "white"

    ax.text(tx, ty + 0.08, pct,
            fontsize=10, color=text_color, weight="bold",
            ha="center", va="center")
    ax.text(tx, ty - 0.08, title,
            fontsize=10, color=text_color, weight="bold",
            ha="center", va="center")

# Texto central
ax.text(0, 0.07, "LONGEVIDADE",
        fontsize=11.5, color=INK, weight="bold", ha="center", va="center")
ax.text(0, -0.07, "FUNCIONAL",
        fontsize=11.5, color=INK, weight="bold", ha="center", va="center")

# Labels externas
ax.text(-1.65, 0.55, "A BASE DE TUDO",
        fontsize=10, color=INK, weight="bold", ha="left", va="top")
ax.text(-1.65, 0.40,
        "Eficiência\nmetabólica\ne resistência.\n\nDeve ocupar\na maior parte\ndo seu tempo.",
        fontsize=8, color=INK_SOFT, ha="left", va="top", linespacing=1.3)

ax.text(1.65, 0.85, "O ESCUDO",
        fontsize=10, color=INK, weight="bold", ha="right", va="top")
ax.text(1.65, 0.70,
        "Preserva massa\nmuscular e protege\no metabolismo\nao longo da vida.",
        fontsize=8, color=INK_SOFT, ha="right", va="top", linespacing=1.3)

ax.text(1.65, -0.30, "O ESTÍMULO",
        fontsize=10, color=INK, weight="bold", ha="right", va="top")
ax.text(1.65, -0.45,
        "Melhora a\ncapacidade\ncardiorrespiratória\nmáxima.",
        fontsize=8, color=INK_SOFT, ha="right", va="top", linespacing=1.3)

ax.text(0, -1.20, "A BASE DA LIBERDADE",
        fontsize=10, color=INK, weight="bold", ha="center", va="top")
ax.text(0, -1.32, "Mantém amplitude de movimento e previne quedas.",
        fontsize=8.5, color=INK_SOFT, ha="center", va="top", style="italic")

# Footer
fig.text(0.5, 0.105,
         "Exercício para longevidade não é um tipo de treino — é um sistema.",
         fontsize=11, color=INK, weight="bold", style="italic", ha="center")
fig.text(0.025, 0.060,
         "Proporção de tempo sugerida para um adulto saudável. Ajustes devem ser feitos conforme",
         fontsize=8, color=FOOT)
fig.text(0.025, 0.040,
         "idade, histórico clínico e objetivos individuais.",
         fontsize=8, color=FOOT)
fig.text(0.025, 0.015,
         "Fontes: Diretrizes ACSM, 2021; Ekkekakis et al., 2020; Pedersen & Saltin, 2015.",
         fontsize=8, color=FOOT)

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap07_Fig03.pdf"
png_path = out_dir / "_preview_Cap07_Fig03.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
