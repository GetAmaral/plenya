"""
Cap09 Fig01 (PT-BR, B&W vetorial) — Três sistemas, uma medicina só.
Diagrama de Venn (3 círculos: Coração, Rim, Metabolismo) com 3 classes de drogas
no centro que afetam os três sistemas.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Circle, FancyArrowPatch, FancyBboxPatch
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
CIRC_1   = "#A0A0A0"   # Coração - cinza médio
CIRC_2   = "#B8B8B8"   # Rim
CIRC_3   = "#CCCCCC"   # Metabolismo
BAND     = "#F4F4F4"

fig = plt.figure(figsize=(11.0, 7.5))
fig.patch.set_facecolor(BG)

# Título
fig.text(0.025, 0.945, "Figura 1 — Três sistemas, uma medicina só.",
         fontsize=16, color=INK, weight="bold")
fig.text(0.025, 0.905,
         "Por que três classes de medicamentos deixaram de ser \"remédio de diabético\" nos últimos cinco anos.",
         fontsize=9.5, color=INK_SOFT, style="italic")

# Venn diagram
ax = fig.add_axes([0.15, 0.18, 0.55, 0.66])
ax.set_xlim(-2.2, 2.2)
ax.set_ylim(-2.0, 2.0)
ax.set_aspect("equal")
ax.axis("off")

# 3 círculos sobrepostos
r = 1.30
# Coração - topo esquerdo
ax.add_patch(Circle((-0.55, 0.55), r, facecolor=CIRC_1, alpha=0.55,
                    edgecolor=INK, linewidth=1.0))
# Rim - topo direito
ax.add_patch(Circle((0.55, 0.55), r, facecolor=CIRC_2, alpha=0.55,
                    edgecolor=INK, linewidth=1.0))
# Metabolismo - centro inferior
ax.add_patch(Circle((0, -0.45), r, facecolor=CIRC_3, alpha=0.55,
                    edgecolor=INK, linewidth=1.0))

# Labels dos círculos (nas bordas externas, em cima)
ax.text(-1.55, 1.45, "Coração", fontsize=12, color=INK, weight="bold")
ax.text(1.10, 1.45, "Rim", fontsize=12, color=INK, weight="bold")
ax.text(0.85, -1.65, "Metabolismo", fontsize=12, color=INK, weight="bold")

# ---------- Caixa central com 3 classes de drogas ----------
# Caixa branca de fundo
ax.add_patch(FancyBboxPatch(
    (-1.30, -0.65), 2.60, 1.40,
    boxstyle="round,pad=0.02,rounding_size=0.05",
    facecolor="white", edgecolor=INK, linewidth=1.0, zorder=5
))

ax.text(-1.18, 0.60, "❶ SGLT2",
        fontsize=10, color=INK, weight="bold", va="center", zorder=6)
ax.text(-1.18, 0.45, "(dapagliflozina, empagliflozina)",
        fontsize=8, color=INK_SOFT, va="center", style="italic", zorder=6)
ax.text(-1.18, 0.30, "DAPA-HF | EMPA-KIDNEY | DELIVER",
        fontsize=7, color=INK_SOFT, va="center", zorder=6)

ax.text(-1.18, 0.05, "❷ Finerenona",
        fontsize=10, color=INK, weight="bold", va="center", zorder=6)
ax.text(-1.18, -0.10, "FIDELIO-DKD | FIGARO-DKD",
        fontsize=7, color=INK_SOFT, va="center", zorder=6)

ax.text(-1.18, -0.30, "❸ GLP-1 / GIP-GLP-1",
        fontsize=10, color=INK, weight="bold", va="center", zorder=6)
ax.text(-1.18, -0.45, "(semaglutida, tirzepatida)",
        fontsize=8, color=INK_SOFT, va="center", style="italic", zorder=6)
ax.text(-1.18, -0.60, "SELECT (2023) | retatrutide fase 2",
        fontsize=7, color=INK_SOFT, va="center", zorder=6)

# ---------- Caixa "Antes / Hoje" à direita ----------
RIGHT_BOX = FancyBboxPatch(
    (0.72, 0.30), 0.25, 0.30,
    boxstyle="round,pad=0.005,rounding_size=0.008",
    facecolor=BAND, edgecolor=INK, linewidth=0.6,
    transform=fig.transFigure, zorder=2
)
fig.patches.append(RIGHT_BOX)
fig.text(0.845, 0.575, "Antes:",
         fontsize=9.5, color=INK_SOFT, weight="bold",
         ha="center", va="center")
fig.text(0.845, 0.555, "remédio de diabético.",
         fontsize=9, color=INK_SOFT, ha="center", va="center", style="italic")

fig.text(0.845, 0.490, "Hoje:",
         fontsize=10, color=INK, weight="bold", ha="center", va="center")
fig.text(0.845, 0.460, "proteção cardiorrenal",
         fontsize=9, color=INK, ha="center", va="center")
fig.text(0.845, 0.440, "e metabólica em",
         fontsize=9, color=INK, ha="center", va="center")
fig.text(0.845, 0.420, "paciente sem diabetes.",
         fontsize=9, color=INK, weight="bold", ha="center", va="center")

# Seta da caixa central da Venn → caixa Antes/Hoje
fig.patches.append(FancyArrowPatch(
    (0.65, 0.470), (0.72, 0.470),
    arrowstyle="->", color=INK, lw=1.0, mutation_scale=12,
    transform=fig.transFigure, zorder=3
))

# Footer
fig.text(0.025, 0.090,
         "Fonte: ensaios DAPA-HF (McMurray, NEJM 2019), EMPA-KIDNEY (Herrington, NEJM 2023), DELIVER (Solomon, NEJM 2022),",
         fontsize=7.5, color=FOOT)
fig.text(0.025, 0.072,
         "FIDELIO-DKD (Bakris, NEJM 2020), FIGARO-DKD (Pitt, NEJM 2021), FIND-CKD (2025), SELECT (Lincoff, NEJM 2023).",
         fontsize=7.5, color=FOOT)
fig.text(0.025, 0.045,
         "Discussão completa do Capítulo 9.",
         fontsize=8, color=FOOT, style="italic")

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap09_Fig01.pdf"
png_path = out_dir / "_preview_Cap09_Fig01.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
