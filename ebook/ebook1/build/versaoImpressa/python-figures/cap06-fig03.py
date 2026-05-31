"""
Cap06 Fig03 (PT-BR, B&W vetorial) — Da Resistência Insulínica ao Diabetes:
A Timeline que o Check-up Não Vê.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, Circle, Ellipse, FancyArrowPatch

# figsize=(11,7.5) → aspect correction
_FIG_ASPECT = 11.0 / 7.5

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
BAR_LIGHT = "#E0E0E0"
BAR_DARK  = "#9A9A9A"
BAND      = "#EDEDED"

fig = plt.figure(figsize=(11.0, 7.5))
fig.patch.set_facecolor(BG)

LEFT = 0.025

# Título
fig.text(LEFT, 0.945, "Figura 3 — Da Resistência Insulínica ao Diabetes:",
         fontsize=16, color=INK, weight="bold")
fig.text(LEFT, 0.905, "A Timeline que o Check-up Não Vê",
         fontsize=16, color=INK, weight="bold")

# ---------- 4 fases — header com nomes ----------
PHASE_TITLES = [
    ("FASE 1", "Resistência insulínica\ncompensada"),
    ("FASE 2", "Disfunção\nmetabólica manifesta"),
    ("FASE 3", "Pré-diabetes"),
    ("FASE 4", "Diabetes\ntipo 2"),
]

# 4 colunas equidistantes
PHASE_X = [0.16, 0.40, 0.64, 0.88]
HEADER_Y = 0.825

for x, (phase, title) in zip(PHASE_X, PHASE_TITLES):
    fig.text(x, HEADER_Y, phase,
             fontsize=10, color=INK, weight="bold", ha="center")
    fig.text(x, HEADER_Y - 0.030, title,
             fontsize=8.5, color=INK_SOFT, ha="center",
             linespacing=1.15)

# ---------- "O QUE ESTÁ ACONTECENDO" — barra horizontal de progressão ----------
fig.text(LEFT, 0.715, "O QUE ESTÁ\nACONTECENDO",
         fontsize=8.5, color=INK_SOFT, weight="bold", va="center",
         linespacing=1.15)

# Barra principal — vai de PHASE_X[0] a PHASE_X[3], gradiente cinza
BAR_Y = 0.715
BAR_HEIGHT = 0.016

# Segments: 1-2 light, 2-3 médio, 3-4 escuro
for i in range(3):
    fig.patches.append(Rectangle(
        (PHASE_X[i], BAR_Y - BAR_HEIGHT/2),
        PHASE_X[i+1] - PHASE_X[i], BAR_HEIGHT,
        facecolor=BAR_LIGHT if i == 0 else (BAR_DARK if i == 2 else "#C0C0C0"),
        edgecolor="none", transform=fig.transFigure, zorder=1
    ))

# Markers (círculos) em cada fase
for x in PHASE_X:
    r = 0.012
    fig.patches.append(Ellipse(
        (x, BAR_Y), width=r*2, height=r*2*_FIG_ASPECT,
        facecolor="white", edgecolor=INK, linewidth=1.2,
        transform=fig.transFigure, zorder=3
    ))

# Detalhes biológicos por fase (texto abaixo do marker)
bio_details = [
    "Insulina ↑\nGlicose normal",
    "Triglicerídeos ↑\nHDL ↓\nGordura hepática ↑\nInflamação ↑",
    "HbA1c\n> 5,7%",
    "HbA1c\n> 6,5%\nDiagnóstico formal",
]

for x, detail in zip(PHASE_X, bio_details):
    fig.text(x, BAR_Y - 0.030, detail,
             fontsize=8, color=INK, ha="center", va="top",
             linespacing=1.2)

# ---------- linha tracejada vertical entre FASE 2 e FASE 3 ("aqui o check-up começa a ver") ----------
VERT_X = (PHASE_X[1] + PHASE_X[2]) / 2
fig.lines.append(plt.Line2D(
    [VERT_X, VERT_X], [0.27, 0.78],
    color=INK, linewidth=0.9, linestyle=(0, (3, 3)),
    transform=fig.transFigure, zorder=2
))
# Label "Aqui o check-up começa a ver"
fig.text(VERT_X + 0.005, 0.490, "Aqui o check-up\ncomeça a ver",
         fontsize=8, color=INK, weight="bold", ha="left", va="center",
         linespacing=1.2)
fig.patches.append(FancyArrowPatch(
    (VERT_X - 0.005, 0.485), (VERT_X - 0.020, 0.485),
    arrowstyle="<-", color=INK, lw=0.9, mutation_scale=10,
    transform=fig.transFigure, zorder=3
))

# ---------- "O QUE O CHECK-UP DIZ" ----------
fig.text(LEFT, 0.450, "O QUE O\nCHECK-UP DIZ",
         fontsize=8.5, color=INK_SOFT, weight="bold", va="center",
         linespacing=1.15)

checkup_quotes = [
    "\"Tudo normal\"",
    "\"Limítrofe,\nacompanhar\"",
    "\"Pré-diabetes —\ndieta e exercício\"",
    "\"Diabetes —\niniciar tratamento\"",
]
CHECKUP_Y = 0.450

for x, q in zip(PHASE_X, checkup_quotes):
    fig.text(x, CHECKUP_Y, q,
             fontsize=9, color=INK, ha="center", va="center",
             linespacing=1.2, style="italic")

# ---------- Fernanda + Janela de intervenção ----------
# Marker da Fernanda — entre fase 1 e fase 2
FERN_X = PHASE_X[0] + 0.45 * (PHASE_X[1] - PHASE_X[0]) + 0.02
FERN_Y = 0.340

_r = 0.012
fig.patches.append(Ellipse(
    (FERN_X, FERN_Y), width=_r*2, height=_r*2*_FIG_ASPECT,
    facecolor=INK, edgecolor=INK, linewidth=1.0,
    transform=fig.transFigure, zorder=3
))
fig.text(FERN_X, FERN_Y - 0.025, "Fernanda\n(41 anos)",
         fontsize=8.5, color=INK, weight="bold", ha="center", va="top",
         linespacing=1.2)

# Janela de intervenção: barra abaixo
JI_Y = 0.265
JI_X1 = PHASE_X[0] - 0.02
JI_X2 = (PHASE_X[1] + PHASE_X[2]) / 2

fig.lines.append(plt.Line2D(
    [JI_X1, JI_X2], [JI_Y, JI_Y],
    color=INK, linewidth=2.0, transform=fig.transFigure
))
fig.lines.append(plt.Line2D(
    [JI_X1, JI_X1], [JI_Y - 0.010, JI_Y + 0.010],
    color=INK, linewidth=1.2, transform=fig.transFigure
))
fig.lines.append(plt.Line2D(
    [JI_X2, JI_X2], [JI_Y - 0.010, JI_Y + 0.010],
    color=INK, linewidth=1.2, transform=fig.transFigure
))

fig.text((JI_X1 + JI_X2) / 2, JI_Y - 0.025,
         "JANELA DE INTERVENÇÃO",
         fontsize=10, color=INK, weight="bold", ha="center", va="top")
fig.text((JI_X1 + JI_X2) / 2, JI_Y - 0.052,
         "5 a 10 anos antes do diagnóstico",
         fontsize=9, color=INK_SOFT, ha="center", va="top", style="italic")

# ---------- linha separadora ----------
SEP_Y = 0.130
fig.lines.append(plt.Line2D(
    [LEFT, 1 - LEFT], [SEP_Y, SEP_Y],
    color="#CFCFCF", linewidth=0.5, transform=fig.transFigure
))

# Footer
fig.text(0.5, 0.085,
         "A doença já estava lá. O diagnóstico é que chegou tarde.",
         fontsize=11, color=INK, weight="bold", style="italic",
         ha="center")

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap06_Fig03.pdf"
png_path = out_dir / "_preview_Cap06_Fig03.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
