"""
Cap06 Fig03 (PT-BR, B&W vetorial) — Da Resistência Insulínica ao Diabetes:
A Timeline que o Check-up Não Vê.

Posições baseadas em detecção pixel-a-pixel do original (1536×1024, aspect 1.5).
"""
import os
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, Ellipse, FancyArrowPatch

_FIG_W, _FIG_H = 11.0, 7.333
_FIG_ASPECT = _FIG_W / _FIG_H  # 1.5

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False
rcParams["pdf.fonttype"] = 42
rcParams["ps.fonttype"] = 42

BG       = "#FFFFFF"
INK      = "#000000"
INK_SOFT = "#3A3A3A"
BAR_LIGHT = "#E0E0E0"
BAR_MID   = "#BDBDBD"
BAR_DARK  = "#7E7E7E"

# --- Modo colorido (opt-in) -------------------------------------------------
# Com FIG_COR=1 a figura sai na paleta da arte original e vai para figuras-cor/.
# Sem a variável, a saída é exatamente a de antes, em figuras-bw/ — a edição
# P&B não muda.
#
# Por que aqui e não recolorindo o PDF pronto, como nas outras figuras: nesta o
# matplotlib agrupa elementos de cores diferentes num mesmo operador de cor —
# "FASE 1" sai junto com o título, e as réguas das duas colunas saem juntas.
# Recolorir o PDF não consegue separá-los; o gerador consegue.
#
# Cores medidas na arte original em 2026-08-24.
COR = os.environ.get("FIG_COR") == "1"
FASE_COR   = ["#105010", "#c88808", "#e05810", "#b80808"] if COR else [INK] * 4
SEG_COR    = (["#94a02c", "#e8a828", "#e05808"] if COR
              else [BAR_LIGHT, BAR_MID, BAR_DARK])
TAG_COR    = "#b81010" if COR else INK       # selo FIGURA 3
ALERTA_COR = "#b80808" if COR else INK       # linha tracejada + "começa a ver"
JANELA_COR = "#205820" if COR else INK       # Fernanda + bracket da janela

fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

# ============= LINHAS HORIZONTAIS AUXILIARES (Y detectados no original) =============
# Y=0.408 (separador bio details / check-up quotes)
# Y=0.314 (separador check-up quotes / Fernanda+bracket)
# Y=0.089 (separador acima do rodapé)
GUIDE_X_FROM = 0.020
GUIDE_X_TO   = 0.980
for hy in (0.408, 0.314, 0.089):
    fig.lines.append(plt.Line2D(
        [GUIDE_X_FROM, GUIDE_X_TO], [hy, hy],
        color="#D0D0D0", linewidth=0.6,
        transform=fig.transFigure, zorder=0
    ))

# ============= LINHAS VERTICAIS AUXILIARES (X detectados no original) =============
# Guias leves entre F1-F2 (X=0.356) e F3-F4 (X=0.755).
# A vertical entre F2-F3 é a dashed principal "aqui o check-up começa a ver" (X=0.568, desenhada abaixo).
GUIDE_Y_TOP    = 0.700
GUIDE_Y_BOTTOM = 0.196
for vx in (0.356, 0.755):
    fig.lines.append(plt.Line2D(
        [vx, vx], [GUIDE_Y_BOTTOM, GUIDE_Y_TOP],
        color="#D0D0D0", linewidth=0.6,
        transform=fig.transFigure, zorder=0
    ))

# ============= FIGURA 3 TAG (substitui vermelho do original) =============
TAG_X, TAG_Y = 0.014, 0.918
TAG_W, TAG_H = 0.097, 0.057
fig.patches.append(Rectangle(
    (TAG_X, TAG_Y), TAG_W, TAG_H,
    facecolor=TAG_COR, edgecolor=TAG_COR, transform=fig.transFigure, zorder=2
))
fig.text(TAG_X + TAG_W/2, TAG_Y + TAG_H/2, "FIGURA 3",
         fontsize=11, color="white", weight="bold",
         ha="center", va="center", zorder=3)

# ============= TÍTULO (Y detectado: 0.945 / 0.877) =============
TITLE_X = TAG_X + TAG_W + 0.018
fig.text(TITLE_X, 0.945, "Da Resistência Insulínica ao Diabetes:",
         fontsize=17, color=INK, weight="bold", va="center")
fig.text(TITLE_X, 0.877, "A Timeline que o Check-up Não Vê",
         fontsize=17, color=INK, weight="bold", va="center")

# ============= 4 COLUNAS FASE (cx detectado) =============
PHASE_X = [0.237, 0.451, 0.650, 0.847]
PHASE_TITLES = [
    ("FASE 1", "Resistência insulínica\ncompensada"),
    ("FASE 2", "Disfunção\nmetabólica manifesta"),
    ("FASE 3", "Pré-diabetes"),
    ("FASE 4", "Diabetes\ntipo 2"),
]

HEADER_Y   = 0.785
SUBTIT_Y1  = 0.729

for i, (x, (phase, title)) in enumerate(zip(PHASE_X, PHASE_TITLES)):
    fig.text(x, HEADER_Y, phase,
             fontsize=12, color=FASE_COR[i], weight="bold",
             ha="center", va="center")
    fig.text(x, SUBTIT_Y1, title,
             fontsize=9.5, color=FASE_COR[i] if COR else INK_SOFT,
             ha="center", va="center", linespacing=1.20)

# ============= TIMELINE =============
LEFT_LABEL_X = 0.018
fig.text(LEFT_LABEL_X, 0.644, "O QUE ESTÁ\nACONTECENDO",
         fontsize=9, color=INK_SOFT, weight="bold", va="center",
         linespacing=1.20)

BAR_Y = 0.644
BAR_HEIGHT = 0.014

# Segmentos (gradiente 1→4: claro → escuro)
seg_colors = SEG_COR
for i in range(3):
    fig.patches.append(Rectangle(
        (PHASE_X[i], BAR_Y - BAR_HEIGHT/2),
        PHASE_X[i+1] - PHASE_X[i], BAR_HEIGHT,
        facecolor=seg_colors[i], edgecolor="none",
        transform=fig.transFigure, zorder=1
    ))

# 4 markers (círculos abertos)
for i, x in enumerate(PHASE_X):
    r = 0.013
    fig.patches.append(Ellipse(
        (x, BAR_Y), width=r*2, height=r*2*_FIG_ASPECT,
        facecolor="white", edgecolor=FASE_COR[i], linewidth=1.4,
        transform=fig.transFigure, zorder=3
    ))

# ============= BIO DETAILS (block Y=0.595 top, ~4 lines) =============
BIO_Y_TOP = 0.595
bio_details = [
    "Insulina ↑\nGlicose normal",
    "Triglicerídeos ↑\nHDL ↓\nGordura hepática ↑\nInflamação ↑",
    "HbA1c > 5,7%",
    "HbA1c > 6,5%\nDiagnóstico formal",
]
for x, detail in zip(PHASE_X, bio_details):
    fig.text(x, BIO_Y_TOP, detail,
             fontsize=9, color=INK, ha="center", va="top",
             linespacing=1.30)

# ============= "O QUE O CHECK-UP DIZ" (Y=0.369 detectado) =============
CHECKUP_Y = 0.369
fig.text(LEFT_LABEL_X, CHECKUP_Y, "O QUE O\nCHECK-UP DIZ",
         fontsize=9, color=INK_SOFT, weight="bold", va="center",
         linespacing=1.20)

checkup_quotes = [
    '"Tudo normal"',
    '"Limítrofe,\nacompanhar"',
    '"Pré-diabetes,\ndieta e exercício"',
    '"Diabetes,\niniciar tratamento"',
]
for x, q in zip(PHASE_X, checkup_quotes):
    fig.text(x, CHECKUP_Y, q,
             fontsize=10.5, color=INK, ha="center", va="center",
             linespacing=1.25, style="italic")

# ============= LINHA DASHED VERTICAL (X=0.568, detectado no original) =============
VERT_X = 0.568
fig.lines.append(plt.Line2D(
    [VERT_X, VERT_X], [0.196, 0.700],
    color=ALERTA_COR, linewidth=1.0, linestyle=(0, (3, 3)),
    transform=fig.transFigure, zorder=2
))

# ============= ANNOTATION "Aqui o check-up começa a ver" (Y=0.317) =============
ANNOT_Y = 0.317
fig.text(VERT_X + 0.012, ANNOT_Y, "Aqui o check-up\ncomeça a ver",
         fontsize=9, color=ALERTA_COR, weight="bold", ha="left", va="center",
         linespacing=1.25)
fig.patches.append(FancyArrowPatch(
    (VERT_X + 0.010, ANNOT_Y), (VERT_X + 0.001, ANNOT_Y),
    arrowstyle="-|>", color=ALERTA_COR, lw=1.0, mutation_scale=10,
    transform=fig.transFigure, zorder=3
))

# ============= FERNANDA (marker @ 0.352, 0.315) =============
FERN_X = 0.352
FERN_Y = 0.315
_r = 0.014
fig.patches.append(Ellipse(
    (FERN_X, FERN_Y), width=_r*2, height=_r*2*_FIG_ASPECT,
    facecolor=JANELA_COR, edgecolor=JANELA_COR, linewidth=1.0,
    transform=fig.transFigure, zorder=3
))
# Labels: Y=0.265 e Y=0.236 detectados
fig.text(FERN_X, 0.265, "Fernanda",
         fontsize=10, color=JANELA_COR, weight="bold", ha="center", va="center")
fig.text(FERN_X, 0.236, "(41 anos)",
         fontsize=9, color=INK, ha="center", va="center")

# ============= BRACKET JANELA (X=0.193 → 0.530, Y=0.196) =============
JI_Y = 0.196
JI_X1 = 0.193
JI_X2 = 0.530
fig.lines.append(plt.Line2D(
    [JI_X1, JI_X2], [JI_Y, JI_Y],
    color=JANELA_COR, linewidth=2.2, transform=fig.transFigure
))
fig.lines.append(plt.Line2D(
    [JI_X1, JI_X1], [JI_Y - 0.012, JI_Y + 0.012],
    color=JANELA_COR, linewidth=1.4, transform=fig.transFigure
))
fig.lines.append(plt.Line2D(
    [JI_X2, JI_X2], [JI_Y - 0.012, JI_Y + 0.012],
    color=JANELA_COR, linewidth=1.4, transform=fig.transFigure
))

# JANELA Y=0.166, "5 a 10 anos" Y=0.136
fig.text((JI_X1 + JI_X2)/2, 0.166,
         "JANELA DE INTERVENÇÃO",
         fontsize=11, color=JANELA_COR, weight="bold", ha="center", va="center")
fig.text((JI_X1 + JI_X2)/2, 0.136,
         "5 a 10 anos antes do diagnóstico",
         fontsize=10, color=INK_SOFT, ha="center", va="center", style="italic")

# ============= FOOTER (Y=0.054) =============
# Na arte original a segunda oração é vermelha — é ela que carrega a tese.
# Em P&B as duas metades saem em tinta, exatamente como antes.
if COR:
    fig.text(0.5, 0.054, "A doença já estava lá.  ",
             fontsize=12, color=INK, weight="bold", style="italic",
             ha="right", va="center")
    fig.text(0.5, 0.054, "  O diagnóstico é que chegou tarde.",
             fontsize=12, color=ALERTA_COR, weight="bold", style="italic",
             ha="left", va="center")
else:
    fig.text(0.5, 0.054,
             "A doença já estava lá. O diagnóstico é que chegou tarde.",
             fontsize=12, color=INK, weight="bold", style="italic",
             ha="center", va="center")

# ============= EXPORT =============
out_dir = Path(__file__).resolve().parents[1] / ("figuras-cor" if COR else "figuras-bw")
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap06_Fig03.pdf"
png_path = out_dir / "_preview_Cap06_Fig03.png"
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
