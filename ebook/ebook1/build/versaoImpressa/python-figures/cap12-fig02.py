"""
Cap12 Fig02 (PT-BR, B&W vetorial) — Quando o pilar psicológico entra, a biologia responde.
Ana — biológico vs psicológico, antes/depois.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, Ellipse, FancyArrowPatch, FancyBboxPatch

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

# Estrutura: 2 seções (BIOLÓGICO, PSICOLÓGICO), cada uma com 2 markers
sections = [
    ("BIOLÓGICO", [
        # (nome, unidade, antes, depois, meta, pos_antes, pos_depois)
        ("Proteína C reativa (PCR-us)", "mg/L",   "1,8", "0,7", "< 1,0", 0.70, 0.25),
        ("Cortisol matinal",            "µg/dL",  "22",  "14",  "10–18", 0.85, 0.45),
    ]),
    ("PSICOLÓGICO", [
        ("Escala de sintomas depressivos\n(PHQ-9)", "", "14", "6", "≥ 10", 0.70, 0.30),
        ("Escala de ansiedade\n(GAD-7)",            "", "16", "5", "≥ 10", 0.80, 0.25),
    ]),
]

fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

LEFT_MARGIN = 0.025

# Título
fig.text(LEFT_MARGIN, 0.945,
         "Figura 2 — Quando o pilar psicológico entra, a biologia responde.",
         fontsize=15, color=INK, weight="bold")
fig.text(LEFT_MARGIN, 0.905,
         "Ana, 44 anos: 18 meses de otimização biológica sem mover dois marcadores.",
         fontsize=9.5, color=INK_SOFT)
fig.text(LEFT_MARGIN, 0.880,
         "Seis meses de trabalho na mente e os quatro saíram do lugar.",
         fontsize=9.5, color=INK_SOFT, weight="bold")

# Legenda topo
LEG_Y = 0.825
def small_circle(x, y, fill, edge=INK, r=0.0085):
    fig.patches.append(Ellipse(
        (x, y), width=r*2, height=r*2*_ASPECT,
        facecolor=fill, edgecolor=edge, linewidth=1.0,
        transform=fig.transFigure, zorder=4
    ))

small_circle(LEFT_MARGIN + 0.005, LEG_Y, "white")
fig.text(LEFT_MARGIN + 0.018, LEG_Y, "ANTES",
         fontsize=9, color=INK, weight="bold", va="center")

small_circle(LEFT_MARGIN + 0.090, LEG_Y, INK)
fig.text(LEFT_MARGIN + 0.103, LEG_Y, "DEPOIS (após 6 meses)",
         fontsize=9, color=INK, weight="bold", va="center")

# Swatch zona meta
fig.patches.append(Rectangle(
    (LEFT_MARGIN + 0.260, LEG_Y - 0.010), 0.020, 0.018,
    facecolor=BAND_OK, edgecolor="#888888", linewidth=0.5,
    transform=fig.transFigure, zorder=3
))
fig.text(LEFT_MARGIN + 0.285, LEG_Y, "ZONA META",
         fontsize=9, color=INK_SOFT, va="center")

# Linha vertical CORTE-ALERTA
fig.lines.append(plt.Line2D(
    [LEFT_MARGIN + 0.395, LEFT_MARGIN + 0.395],
    [LEG_Y - 0.013, LEG_Y + 0.013],
    color=INK, linewidth=1.0, linestyle=(0, (3, 2)),
    transform=fig.transFigure, zorder=3
))
fig.text(LEFT_MARGIN + 0.405, LEG_Y, "CORTE-ALERTA (risco aumentado)",
         fontsize=9, color=INK_SOFT, va="center")

# 2 seções
ROWS = []  # accumulate row y positions for drawing
SEC_TOP_LIST = [0.755, 0.430]
SEC_HEIGHT  = 0.260

for s_idx, (section_name, markers) in enumerate(sections):
    sec_top = SEC_TOP_LIST[s_idx]
    sec_bot = sec_top - SEC_HEIGHT

    # Header da seção
    fig.text(LEFT_MARGIN, sec_top + 0.005, section_name,
             fontsize=11, color=INK, weight="bold")

    # Headers ANTES / DEPOIS / META
    headers_y = sec_top - 0.030
    fig.text(0.30, headers_y, "ANTES",
             fontsize=8.5, color=TICK, ha="center", style="italic")
    fig.text(0.55, headers_y, "DEPOIS",
             fontsize=8.5, color=TICK, ha="center", style="italic")
    fig.text(0.91, headers_y, "META",
             fontsize=8.5, color=TICK, ha="center", style="italic")

    # 2 markers por seção
    BAR_LEFT  = 0.27
    BAR_RIGHT = 0.83
    META_X    = 0.91

    n = len(markers)
    row_top = sec_top - 0.050
    row_bot = sec_bot
    if n > 1:
        row_space = (row_top - row_bot) / (n - 1)
    else:
        row_space = 0

    BAR_HEIGHT = 0.028

    for r, (name, unit, antes, depois, meta, pos_a, pos_d) in enumerate(markers):
        y = row_top - r * row_space if n > 1 else (row_top + row_bot)/2
        bar_w = BAR_RIGHT - BAR_LEFT

        # Nome
        fig.text(LEFT_MARGIN, y + 0.010, name,
                 fontsize=9.5, color=INK, weight="bold",
                 va="center", linespacing=1.2)
        if unit:
            fig.text(LEFT_MARGIN, y - 0.015, unit,
                     fontsize=8, color=TICK, va="center")

        # Barra background
        fig.patches.extend([
            Rectangle((BAR_LEFT, y - BAR_HEIGHT/2), bar_w * 0.40, BAR_HEIGHT,
                      facecolor=BAND_OK, edgecolor="none",
                      transform=fig.transFigure, zorder=1),
            Rectangle((BAR_LEFT + bar_w * 0.40, y - BAR_HEIGHT/2),
                      bar_w * 0.30, BAR_HEIGHT,
                      facecolor=BAND_MID, edgecolor="none",
                      transform=fig.transFigure, zorder=1),
            Rectangle((BAR_LEFT + bar_w * 0.70, y - BAR_HEIGHT/2),
                      bar_w * 0.30, BAR_HEIGHT,
                      facecolor=BAND_BAD, edgecolor="none",
                      transform=fig.transFigure, zorder=1),
        ])

        # Linha CORTE-ALERTA (dashed)
        x_cut = BAR_LEFT + bar_w * 0.70
        fig.lines.append(plt.Line2D(
            [x_cut, x_cut], [y - BAR_HEIGHT*0.85, y + BAR_HEIGHT*0.85],
            color=INK, linewidth=0.8, linestyle=(0, (3, 2)),
            transform=fig.transFigure, zorder=4
        ))

        x_antes  = BAR_LEFT + bar_w * pos_a
        x_depois = BAR_LEFT + bar_w * pos_d

        # Seta antes → depois
        fig.patches.append(FancyArrowPatch(
            (x_antes, y - 0.022), (x_depois, y - 0.022),
            arrowstyle="->", color=INK, lw=1.5, mutation_scale=12,
            transform=fig.transFigure, zorder=4
        ))

        # Marker ANTES (círculo branco)
        small_circle(x_antes, y, "white")
        fig.text(x_antes, y + 0.028, antes,
                 fontsize=9.5, color=INK_SOFT, weight="bold",
                 ha="center", va="bottom")

        # Marker DEPOIS (círculo preto)
        small_circle(x_depois, y, INK)
        fig.text(x_depois, y + 0.028, depois,
                 fontsize=11, color=INK, weight="bold",
                 ha="center", va="bottom")

        # Meta (texto à direita)
        fig.text(META_X, y, meta,
                 fontsize=11, color=INK, weight="bold",
                 ha="center", va="center")

# Footer
fig.text(LEFT_MARGIN, 0.060,
         "Fonte: caso-tipo do Capítulo 12. Escala de sintomas depressivos (PHQ-9) (Kroenke et al., JGM 2001).",
         fontsize=7.5, color=FOOT)
fig.text(LEFT_MARGIN, 0.040,
         "Escala de ansiedade (GAD-7) (Spitzer et al., Arch Intern Med 2006).",
         fontsize=7.5, color=FOOT, style="italic")

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap12_Fig02.pdf"
png_path = out_dir / "_preview_Cap12_Fig02.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
