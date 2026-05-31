"""
Cap13 Fig03 (PT-BR, B&W vetorial) — Ikigai: o que te faz levantar da cama.
Diagrama Venn com 4 círculos + callout.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Circle, FancyBboxPatch

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
SHADE_1  = "#D0D0D0"
SHADE_2  = "#B8B8B8"
SHADE_3  = "#A8A8A8"
SHADE_4  = "#909090"
BAND     = "#EDEDED"

fig = plt.figure(figsize=(11.0, 7.5))
fig.patch.set_facecolor(BG)

LEFT = 0.025

# Título
fig.text(LEFT, 0.945,
         "Figura 3 — Ikigai: o que te faz levantar da cama.",
         fontsize=15, color=INK, weight="bold")
fig.text(LEFT, 0.910,
         "Quatro círculos que se cruzam num ponto. O estudo Ohsaki (2008) associou",
         fontsize=9.5, color=INK_SOFT, style="italic")
fig.text(LEFT, 0.892,
         "a resposta afirmativa à pergunta a 30–50% menos mortalidade em 7 anos.",
         fontsize=9.5, color=INK_SOFT, style="italic")

# Diagrama Venn
ax = fig.add_axes([0.02, 0.18, 0.55, 0.65])
ax.set_xlim(-2.2, 2.2)
ax.set_ylim(-1.8, 1.8)
ax.set_aspect("equal")
ax.axis("off")

# 4 círculos sobrepostos (mesmo raio, espaçados)
r = 1.05
# Posições nos cantos NE, NW, SW, SE
positions = [
    ( 0.0,  0.65, SHADE_1, "O que você ama",            "PAIXÃO",    ( 0.0,  1.45), ( 0.0,  0.85)),
    ( 0.65, 0.0,  SHADE_2, "O que o mundo precisa",     "MISSÃO",    ( 1.65, 0.85), ( 0.85, 0.20)),
    ( 0.0, -0.65, SHADE_3, "O que o mundo precisa",     "VOCAÇÃO",   ( 0.0, -1.45), ( 0.0, -0.85)),
    (-0.65, 0.0,  SHADE_4, "O que faz bem",             "PROFISSÃO", (-1.65, 0.85), (-0.85, 0.20)),
]

# 4 círculos com alpha
for cx, cy, color, _, _, _, _ in positions:
    ax.add_patch(Circle((cx, cy), r, facecolor=color, alpha=0.45,
                        edgecolor=INK, linewidth=1.0))

# Labels externas dos círculos
ext_labels = [
    ("O que você ama",       0.0,  1.60),
    ("O que o mundo precisa", 1.70, 0.85),
    ("Pelo que pode ser pago", 1.70, -0.85),
    ("O que faz bem",        -1.70, 0.85),
]
# Atualizar pra posições corretas
actual_ext = [
    ("O que você ama",                 0.0,   1.60, "center"),
    ("O que o mundo precisa",          1.85,  0.65, "left"),
    ("Pelo que pode ser pago",         1.85, -0.65, "left"),
    ("O que faz bem",                 -1.85,  0.65, "right"),
    ("O que o mundo\nprecisa",         0.0,  -1.55, "center"),  # bottom
]

# Sobreposição correta (4 círculos clássicos do Ikigai):
# top: você ama, right: mundo precisa, bottom: pago, left: faz bem
# Ajustando:
ax.clear()
ax.set_xlim(-2.4, 2.4)
ax.set_ylim(-1.8, 1.8)
ax.set_aspect("equal")
ax.axis("off")

circles = [
    ( 0.0,  0.65, SHADE_1),   # topo
    ( 0.65, 0.0,  SHADE_2),   # direita
    ( 0.0, -0.65, SHADE_3),   # baixo
    (-0.65, 0.0,  SHADE_4),   # esquerda
]
for cx, cy, color in circles:
    ax.add_patch(Circle((cx, cy), r, facecolor=color, alpha=0.45,
                        edgecolor=INK, linewidth=1.0))

# Labels externas — 4 cantos do diagrama Venn
# TOP, RIGHT, BOTTOM, LEFT
external = [
    ("O que você ama",          0.0,   1.65, "center", "bottom"),
    ("Pelo que pode\nser pago", 1.70,  0.0,  "left",   "center"),
    ("O que o mundo\nprecisa",  0.0,  -1.65, "center", "top"),
    ("O que faz bem",          -1.70,  0.0,  "right",  "center"),
]

for label, x, y, ha, va in external:
    ax.text(x, y, label,
            fontsize=10, color=INK, weight="bold",
            ha=ha, va=va, linespacing=1.2)

# Label "O que o mundo precisa" — bottom também (era na figura original em embaixo)
# Hmm no original o LEFT é "O que faz bem", BOTTOM é "O que o mundo precisa"
# Vou re-checar. Olhando original:
# - TOP: "O que você ama" (com PAIXÃO no inner label)
# - RIGHT: "O que o mundo precisa" (com MISSÃO)
# - BOTTOM: "O que o mundo precisa" (com VOCAÇÃO)? Hmm seria estranho repetir
# Actually no original talvez:
# - BOTTOM: "Pelo que pode ser pago" (VOCAÇÃO)
# - LEFT: "O que faz bem" (PROFISSÃO)
# OK que é o que eu tenho.

# Inner labels (PAIXÃO, MISSÃO, VOCAÇÃO, PROFISSÃO) — posicionados nos "lóbulos"
inner_labels = [
    ("PAIXÃO",     0.0,  1.05),
    ("MISSÃO",     1.05, 0.0),
    ("VOCAÇÃO",    0.0, -1.05),
    ("PROFISSÃO", -1.05, 0.0),
]
for lbl, x, y in inner_labels:
    ax.text(x, y, lbl,
            fontsize=10, color=INK, weight="bold",
            ha="center", va="center")

# Central IKIGAI
ax.text(0, 0.10, "IKIGAI",
        fontsize=15, color=INK, weight="bold", ha="center", va="center")
ax.text(0, -0.15, "A razão",
        fontsize=8.5, color=INK_SOFT, ha="center", va="center", style="italic")
ax.text(0, -0.27, "de acordar.",
        fontsize=8.5, color=INK_SOFT, ha="center", va="center", style="italic")

# ---------- callout direito ----------
CO_X1, CO_X2 = 0.72, 0.97
CO_Y1, CO_Y2 = 0.30, 0.80

fig.patches.append(FancyBboxPatch(
    (CO_X1, CO_Y1), CO_X2 - CO_X1, CO_Y2 - CO_Y1,
    boxstyle="round,pad=0.005,rounding_size=0.008",
    facecolor="white", edgecolor=INK, linewidth=1.0,
    transform=fig.transFigure, zorder=1
))

fig.text((CO_X1 + CO_X2) / 2, CO_Y2 - 0.030,
         "Pergunta clínica:",
         fontsize=10, color=INK, weight="bold", ha="center")

fig.text(CO_X1 + 0.015, CO_Y2 - 0.075,
         "\"Qual é o motivo que te",
         fontsize=10, color=INK, style="italic", weight="bold")
fig.text(CO_X1 + 0.015, CO_Y2 - 0.097,
         "faz levantar da cama de",
         fontsize=10, color=INK, style="italic", weight="bold")
fig.text(CO_X1 + 0.015, CO_Y2 - 0.119,
         "manhã?\"",
         fontsize=10, color=INK, style="italic", weight="bold")

fig.text(CO_X1 + 0.015, CO_Y2 - 0.170,
         "Se não houver resposta clara",
         fontsize=8.5, color=INK)
fig.text(CO_X1 + 0.015, CO_Y2 - 0.190,
         "em um adulto pós-40 anos,",
         fontsize=8.5, color=INK)
fig.text(CO_X1 + 0.015, CO_Y2 - 0.210,
         "há uma camada do plano que",
         fontsize=8.5, color=INK)
fig.text(CO_X1 + 0.015, CO_Y2 - 0.230,
         "está faltando.",
         fontsize=8.5, color=INK, weight="bold")

# Footer
fig.text(LEFT, 0.105,
         "Fonte: conceito japonês tradicional.",
         fontsize=8, color=FOOT, style="italic")
fig.text(LEFT, 0.080,
         "Evidência: Sone et al (\"Sense of Life Worth Living and Mortality in Japan: Ohsaki Study\"),",
         fontsize=7.5, color=FOOT)
fig.text(LEFT, 0.060,
         "Psychosomatic Medicine 2008, 70(8):709-715. Coorte de 43.391 adultos, 7 anos de seguimento.",
         fontsize=7.5, color=FOOT)

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap13_Fig03.pdf"
png_path = out_dir / "_preview_Cap13_Fig03.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
