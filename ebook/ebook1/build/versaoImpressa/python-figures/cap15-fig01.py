"""
Cap15 Fig01 (PT-BR, B&W vetorial) — Trajetória de Marcos.
Painel de biomarcadores antes/depois 8 meses.
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
BAND_OK  = "#F4F4F4"   # ótimo
BAND_MID = "#DEDEDE"   # alerta
BAND_BAD = "#BFBFBF"   # risco
BAND     = "#EDEDED"

_FIG_W, _FIG_H = 11.0, 7.5
_ASPECT = _FIG_W / _FIG_H

fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

LEFT = 0.025

fig.text(LEFT, 0.945,
         "Figura 1 — Trajetória de Marcos: painel de biomarcadores antes e oito meses depois",
         fontsize=14, color=INK, weight="bold")
fig.text(LEFT, 0.910,
         "6 de 7 marcadores migraram para zona ótima. 1 marcador estrutural permaneceu inalterado.",
         fontsize=9.5, color=INK_SOFT, style="italic")

# Headers
HEAD_Y = 0.860
fig.text(LEFT,        HEAD_Y, "MARCADOR",
         fontsize=9, color=INK_SOFT, weight="bold")
fig.text(0.21,        HEAD_Y, "VALOR",
         fontsize=9, color=INK_SOFT, weight="bold", ha="center")
fig.text(0.52,        HEAD_Y, "FAIXA DE REFERÊNCIA (ZONA)",
         fontsize=9, color=INK_SOFT, weight="bold", ha="center")
fig.text(0.85,        HEAD_Y, "INTERPRETAÇÃO",
         fontsize=9, color=INK_SOFT, weight="bold")

# Linha
fig.lines.append(plt.Line2D(
    [LEFT, 0.98], [0.838, 0.838],
    color="#888888", linewidth=0.5, transform=fig.transFigure
))

# Sub-headers: ANTES / DEPOIS
fig.text(0.165, 0.815, "Mês 0",
         fontsize=7.5, color=INK_SOFT, ha="center", style="italic")
fig.text(0.165, 0.802, "(baseline)",
         fontsize=7, color=INK_SOFT, ha="center", style="italic")
fig.text(0.245, 0.815, "8 meses",
         fontsize=7.5, color=INK_SOFT, ha="center", style="italic")
fig.text(0.245, 0.802, "depois",
         fontsize=7, color=INK_SOFT, ha="center", style="italic")

# Sub-headers das zonas
fig.text(0.36, 0.815, "RISCO",
         fontsize=7.5, color=INK, weight="bold", ha="center")
fig.text(0.52, 0.815, "ALERTA",
         fontsize=7.5, color=INK, weight="bold", ha="center")
fig.text(0.68, 0.815, "ÓTIMO",
         fontsize=7.5, color=INK, weight="bold", ha="center")

# 7 marcadores
# (name, antes, depois, pos_antes, pos_depois, interpretação)
biomarkers = [
    ("ApoB",                       "82",  "58",   0.50, 0.70, "Alvo atingido"),
    ("Insulina de jejum",         "11",  "6",    0.40, 0.78, "Normalização metabólica"),
    ("PCR ultrassensível",        "1,6", "0,6",  0.50, 0.75, "Inflamação controlada"),
    ("Vitamina D",                 "28",  "52",  0.30, 0.78, "Faixa ótima"),
    ("Ergometria\n(MET)",          "—",   "+1,8 MET", 0.40, 0.75, "Ganho de capacidade\nfuncional"),
    ("Composição corporal",        "—",   "−5 kg gordura\n+2 kg massa magra", 0.40, 0.75, "Recomposição"),
    ("CAC",                        "412", "412",  0.35, 0.35, "Marcador estrutural\ninalterado"),
]

# Bar zone backgrounds
BAR_LEFT  = 0.30
BAR_RIGHT = 0.78

def small_circle(x, y, fill, r=0.0085):
    fig.patches.append(Ellipse(
        (x, y), width=r*2, height=r*2*_ASPECT,
        facecolor=fill, edgecolor=INK, linewidth=0.8,
        transform=fig.transFigure, zorder=4
    ))

ROW_TOP = 0.760
ROW_BOTTOM = 0.180
ROW_SPACE = (ROW_TOP - ROW_BOTTOM) / (len(biomarkers) - 1)
BAR_H = 0.022

for i, (name, antes, depois, pos_a, pos_d, interp) in enumerate(biomarkers):
    y = ROW_TOP - i * ROW_SPACE
    bar_w = BAR_RIGHT - BAR_LEFT

    # Nome
    fig.text(LEFT, y + 0.005, name,
             fontsize=9, color=INK, weight="bold", va="center",
             linespacing=1.2)

    # Valor antes / depois
    fig.text(0.165, y + 0.005, antes,
             fontsize=10, color=INK_SOFT, weight="bold", ha="center", va="center")
    fig.text(0.245, y + 0.005, depois,
             fontsize=10, color=INK, weight="bold", ha="center", va="center",
             linespacing=1.15)

    # Zone bar (3 zonas)
    fig.patches.extend([
        Rectangle((BAR_LEFT, y - BAR_H/2), bar_w * 0.33, BAR_H,
                  facecolor=BAND_BAD, edgecolor="none",
                  transform=fig.transFigure, zorder=1),
        Rectangle((BAR_LEFT + bar_w*0.33, y - BAR_H/2), bar_w * 0.34, BAR_H,
                  facecolor=BAND_MID, edgecolor="none",
                  transform=fig.transFigure, zorder=1),
        Rectangle((BAR_LEFT + bar_w*0.67, y - BAR_H/2), bar_w * 0.33, BAR_H,
                  facecolor=BAND_OK, edgecolor="none",
                  transform=fig.transFigure, zorder=1),
    ])

    # Markers + seta
    x_a = BAR_LEFT + bar_w * pos_a
    x_d = BAR_LEFT + bar_w * pos_d

    if pos_a != pos_d:
        fig.patches.append(FancyArrowPatch(
            (x_a, y), (x_d, y),
            arrowstyle="->", color=INK, lw=1.3, mutation_scale=12,
            transform=fig.transFigure, zorder=4
        ))
    small_circle(x_a, y + 0.015, "white")
    small_circle(x_d, y + 0.015, INK)

    # Interpretação
    fig.text(0.80, y + 0.005, interp,
             fontsize=8.5, color=INK, va="center",
             linespacing=1.2)

# Legenda inferior
LEG_Y = 0.110
small_circle(LEFT + 0.005, LEG_Y, "white")
fig.text(LEFT + 0.018, LEG_Y, "Mês 0 (baseline)",
         fontsize=8, color=INK, va="center")

small_circle(LEFT + 0.150, LEG_Y, INK)
fig.text(LEFT + 0.163, LEG_Y, "8 meses depois",
         fontsize=8, color=INK, va="center")

# Swatches das zonas
for x_off, label, color in [(0.35, "RISCO", BAND_BAD),
                              (0.46, "ALERTA", BAND_MID),
                              (0.57, "ÓTIMO", BAND_OK)]:
    fig.patches.append(Rectangle(
        (x_off, LEG_Y - 0.008), 0.016, 0.016,
        facecolor=color, edgecolor="#888888", linewidth=0.5,
        transform=fig.transFigure, zorder=3
    ))
    fig.text(x_off + 0.020, LEG_Y, label,
             fontsize=8, color=INK_SOFT, va="center")

fig.text(0.65, LEG_Y, "● → ● = direção e magnitude da mudança clínica",
         fontsize=8, color=INK_SOFT, va="center", style="italic")

# Footer
fig.text(LEFT, 0.058,
         "Cálcio coronariano não se reduz, é estabilizado e o programa modifica a progressão. A queda do CAC reflete falha do programa, não meta clínica.",
         fontsize=7.5, color=FOOT, style="italic")
fig.text(LEFT, 0.035,
         "Marcos: trajetória de 8 meses (cf. Capítulos 7-8). Tabela representativa baseada nos achados clínicos.",
         fontsize=7.5, color=FOOT)

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap15_Fig01.pdf"
png_path = out_dir / "_preview_Cap15_Fig01.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
