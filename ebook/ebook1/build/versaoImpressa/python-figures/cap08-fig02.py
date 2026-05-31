"""
Cap08 Fig02 (PT-BR, B&W vetorial) — Finasterida: quando sim, quando não.
Fluxograma de decisão clínica.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import FancyBboxPatch, Polygon, FancyArrowPatch, Rectangle

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
BAND_OK  = "#EDEDED"     # caixa "SIM"
BAND_BAD = "#D0D0D0"     # caixa "NÃO"
BAND_NEU = "#F4F4F4"     # caixas neutras
BAND_REG = "#E8E8E8"     # base regulatória

fig = plt.figure(figsize=(11.0, 8.6))
fig.patch.set_facecolor(BG)

# Título
fig.text(0.025, 0.955, "Figura 2 — Finasterida: quando sim, quando não.",
         fontsize=15, color=INK, weight="bold")
fig.text(0.025, 0.918,
         "O posicionamento clínico do autor — indicação regulatória vs. uso cosmético, à luz da síndrome pós-finasterida (PFS).",
         fontsize=9, color=INK_SOFT, style="italic")

# ---------- Caixa do paciente (topo) ----------
PAT_BOX = FancyBboxPatch(
    (0.30, 0.83), 0.40, 0.045,
    boxstyle="round,pad=0.005,rounding_size=0.005",
    facecolor=BAND_NEU, edgecolor=INK, linewidth=1.0,
    transform=fig.transFigure, zorder=2
)
fig.patches.append(PAT_BOX)
fig.text(0.50, 0.852, "Paciente adulto com",
         fontsize=10, color=INK, weight="bold", ha="center", va="center")
fig.text(0.50, 0.838, "queda de cabelo e/ou sintomas prostáticos",
         fontsize=9, color=INK_SOFT, ha="center", va="center")

# Seta pra baixo
fig.patches.append(FancyArrowPatch(
    (0.50, 0.830), (0.50, 0.795),
    arrowstyle="->", color=INK, lw=1.2, mutation_scale=14,
    transform=fig.transFigure, zorder=3
))

# ---------- Diamante de decisão ----------
diamond_pts = [
    (0.50, 0.810),  # topo
    (0.66, 0.745),  # direita
    (0.50, 0.680),  # fundo
    (0.34, 0.745),  # esquerda
]
fig.patches.append(Polygon(
    diamond_pts, closed=True,
    facecolor=BAND_NEU, edgecolor=INK, linewidth=1.0,
    transform=fig.transFigure, zorder=2
))
fig.text(0.50, 0.768, "Tem HBP",
         fontsize=10, color=INK, weight="bold", ha="center", va="center")
fig.text(0.50, 0.748, "sintomática",
         fontsize=10, color=INK, weight="bold", ha="center", va="center")
fig.text(0.50, 0.722, "(IPSS ≥ 8)?",
         fontsize=10, color=INK, weight="bold", ha="center", va="center")

# ---------- Setas SIM (esquerda) e NÃO (direita) ----------
# SIM para a esquerda
fig.patches.append(FancyArrowPatch(
    (0.34, 0.745), (0.23, 0.745),
    arrowstyle="->", color=INK, lw=1.2, mutation_scale=14,
    transform=fig.transFigure, zorder=3
))
fig.text(0.305, 0.757, "SIM", fontsize=11, color=INK, weight="bold",
         ha="center", va="center")

# NÃO para a direita
fig.patches.append(FancyArrowPatch(
    (0.66, 0.745), (0.77, 0.745),
    arrowstyle="->", color=INK, lw=1.2, mutation_scale=14,
    transform=fig.transFigure, zorder=3
))
fig.text(0.695, 0.757, "NÃO", fontsize=11, color=INK, weight="bold",
         ha="center", va="center")

# ---------- Caixa SIM (à esquerda) ----------
SIM_BOX = FancyBboxPatch(
    (0.04, 0.45), 0.19, 0.27,
    boxstyle="round,pad=0.005,rounding_size=0.008",
    facecolor=BAND_OK, edgecolor=INK, linewidth=1.0,
    transform=fig.transFigure, zorder=1
)
fig.patches.append(SIM_BOX)
# título com check
fig.text(0.135, 0.700, "✓  Indicação clínica consistente",
         fontsize=10, color=INK, weight="bold", ha="center", va="center")
fig.text(0.135, 0.680, "Finasterida 5 mg/dia OU",
         fontsize=9, color=INK, ha="center", va="center")
fig.text(0.135, 0.665, "dutasterida 0,5 mg/dia",
         fontsize=9, color=INK, ha="center", va="center")
fig.text(0.135, 0.640, "— décadas de evidência",
         fontsize=8.5, color=INK_SOFT, style="italic",
         ha="center", va="center")

# Separador horizontal dentro da caixa
fig.lines.append(plt.Line2D(
    [0.06, 0.21], [0.610, 0.610],
    color="#AAAAAA", linewidth=0.5, transform=fig.transFigure
))

fig.text(0.135, 0.580, "Consentimento informado",
         fontsize=9, color=INK, weight="bold",
         ha="center", va="center")
fig.text(0.135, 0.560, "sobre efeitos sexuais",
         fontsize=9, color=INK, ha="center", va="center")
fig.text(0.135, 0.545, "possíveis.",
         fontsize=9, color=INK, ha="center", va="center")

# ---------- Caixa NÃO (à direita) ----------
NAO_BOX = FancyBboxPatch(
    (0.77, 0.30), 0.19, 0.42,
    boxstyle="round,pad=0.005,rounding_size=0.008",
    facecolor=BAND_BAD, edgecolor=INK, linewidth=1.0,
    transform=fig.transFigure, zorder=1
)
fig.patches.append(NAO_BOX)
# título com X
fig.text(0.865, 0.700, "✗  Indicação apenas cosmética",
         fontsize=10, color=INK, weight="bold", ha="center", va="center")
fig.text(0.865, 0.683, "(alopecia androgenética)",
         fontsize=8.5, color=INK_SOFT, style="italic",
         ha="center", va="center")

fig.text(0.865, 0.654, "NÃO prescrever finasterida",
         fontsize=9, color=INK, weight="bold",
         ha="center", va="center")
fig.text(0.865, 0.638, "sem dutasterida.",
         fontsize=9, color=INK, weight="bold",
         ha="center", va="center")

fig.text(0.865, 0.614, "Risco de pós-finasterida (PFS)",
         fontsize=8, color=INK, ha="center", va="center", style="italic")
fig.text(0.865, 0.600, "— uso cosmético é off-label.",
         fontsize=8, color=INK, ha="center", va="center", style="italic")

# Separador
fig.lines.append(plt.Line2D(
    [0.79, 0.95], [0.578, 0.578],
    color="#AAAAAA", linewidth=0.5, transform=fig.transFigure
))

fig.text(0.865, 0.555, "Alternativas com evidência:",
         fontsize=9, color=INK, weight="bold",
         ha="center", va="center")
alt_lines = [
    "minoxidil tópico 5% ou oral",
    "em baixa dose; PRP capilar;",
    "microneedling; correção",
    "de vitamina D, B12, zinco,",
    "função tireoidiana.",
]
for i, ln in enumerate(alt_lines):
    fig.text(0.865, 0.530 - i * 0.020, ln,
             fontsize=8.5, color=INK_SOFT, ha="center", va="center")

# ---------- Base regulatória (timeline embaixo) ----------
fig.text(0.50, 0.22, "BASE REGULATÓRIA",
         fontsize=10, color=INK_SOFT, weight="bold", ha="center")

# Linha horizontal
fig.lines.append(plt.Line2D(
    [0.10, 0.90], [0.180, 0.180],
    color=INK, linewidth=1.0, transform=fig.transFigure
))

# 3 marcos
milestones = [
    (0.20, "2011 — FDA",  "Sexual adverse\nevents"),
    (0.50, "2022 — FDA",  "Long-term\npersistence"),
    (0.80, "2025 — EMA",  "Recomendações\nformais sobre\nsíndrome\npós-finasterida"),
]

for x, lbl, sub in milestones:
    # marker
    fig.patches.append(Rectangle(
        (x - 0.003, 0.176), 0.006, 0.008,
        facecolor=INK, edgecolor="none", transform=fig.transFigure, zorder=3
    ))
    fig.text(x, 0.160, lbl,
             fontsize=9, color=INK, weight="bold", ha="center", va="top")
    fig.text(x, 0.140, sub,
             fontsize=7.5, color=INK_SOFT, ha="center", va="top",
             linespacing=1.2, style="italic")

# Footer
fig.text(0.025, 0.060,
         "Fonte: FDA Drug Safety Communications (2011, 2022); EMA PRAC Recommendations (2025).",
         fontsize=8, color=FOOT)
fig.text(0.025, 0.038,
         "PFS: post-finasteride syndrome. IPSS: International Prostate Symptom Score. HBP: hiperplasia benigna da próstata.",
         fontsize=8, color=FOOT)
fig.text(0.025, 0.015,
         "Discussão clínica completa do autor no Capítulo 8.",
         fontsize=8, color=FOOT, style="italic")

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap08_Fig02.pdf"
png_path = out_dir / "_preview_Cap08_Fig02.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
