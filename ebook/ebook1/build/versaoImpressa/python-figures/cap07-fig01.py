"""
Cap07 Fig01 (PT-BR, B&W vetorial) — Ser sedentário extremo aumenta em ~5x o risco de morte.
Barras horizontais de hazard ratio de mortalidade.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams

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
BAR_TOP  = "#3A3A3A"     # mais escuro — o risco mais alto
BAR_OTHER = "#9A9A9A"    # cinza médio
BAND     = "#F4F4F4"

fig = plt.figure(figsize=(11.0, 7.0))
fig.patch.set_facecolor(BG)

# Título
fig.text(0.025, 0.945, "Figura 1 — Ser sedentário extremo aumenta em ~5x o risco de morte",
         fontsize=15, color=INK, weight="bold")
fig.text(0.025, 0.905, "Comparado a outros fatores de risco conhecidos",
         fontsize=10, color=INK_SOFT, style="italic")

# Eixo
ax = fig.add_axes([0.30, 0.20, 0.62, 0.62])

bars = [
    ("Aptidão cardiorrespiratória\nbaixa (sedentarismo extremo)", 5.0, True),
    ("Doença renal\nem estágio terminal",                          2.0, False),
    ("Tabagismo",                                                  1.4, False),
    ("Diabetes",                                                   1.4, False),
    ("Doença coronariana\n(do coração)",                           1.3, False),
]

labels = [b[0] for b in bars]
values = [b[1] for b in bars]
colors = [BAR_TOP if b[2] else BAR_OTHER for b in bars]

y_pos = list(range(len(bars)))
ax.barh(y_pos, values, color=colors, height=0.65, edgecolor="none", zorder=3)

# Valores no fim de cada barra
for i, v in enumerate(values):
    label_v = f"{v:.1f}x".replace(".", ",")
    weight = "bold" if i == 0 else "bold"
    color  = INK
    fontsize = 14 if i == 0 else 11
    ax.text(v + 0.10, i, label_v,
            color=color, weight=weight, fontsize=fontsize, va="center")

# Linha de referência em 1.0x
ax.axvline(1.0, color=INK, linewidth=1.0, linestyle=(0, (3, 3)), zorder=2)
# "1,0x" label imediatamente à direita da linha tracejada, posicionada
# no espaço entre Tabagismo e Diabetes mas afastada das barras
ax.text(1.07, -0.50, "1,0x = Risco de referência (pessoas em excelente forma)",
        fontsize=8, color=INK_SOFT, weight="bold", va="center", style="italic")

# Eixo Y: rótulos
ax.set_yticks(y_pos)
ax.set_yticklabels(labels, fontsize=10, color=INK, weight="bold",
                    linespacing=1.15)
ax.invert_yaxis()  # primeira barra no topo

# Eixo X
ax.set_xlim(0, 6)
ax.set_xticks([0, 1, 2, 3, 4, 5])
ax.set_xticklabels([str(t) for t in [0, 1, 2, 3, 4, 5]],
                   fontsize=9, color=TICK)
ax.set_xlabel("Risco de morte (vezes maior)",
              fontsize=10, color=INK, weight="bold")

# Cosmética
for spine in ("top", "right"):
    ax.spines[spine].set_visible(False)
ax.spines["left"].set_color("#888888")
ax.spines["bottom"].set_color("#888888")
ax.tick_params(axis="y", length=0)
ax.tick_params(axis="x", colors=TICK)
ax.set_axisbelow(True)
ax.grid(axis="x", color="#EEEEEE", linewidth=0.6, zorder=0)

# Frase abaixo do gráfico
fig.text(0.5, 0.105,
         "Mais letal do que fumar, diabetes ou doença do coração.",
         fontsize=11, color=INK, weight="bold", ha="center", style="italic")

# Footer
fig.text(0.025, 0.055,
         "Fonte: Mandsager et al., JAMA Network Open, 2018. Estudo com 122.007 adultos acompanhados por mediana de 8,4 anos.",
         fontsize=7.5, color=FOOT)
fig.text(0.025, 0.030,
         "Aptidão baixa definida como abaixo do percentil 25 para idade e sexo. Referência: grupo elite (acima do percentil 97,5).",
         fontsize=7.5, color=FOOT)

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap07_Fig01.pdf"
png_path = out_dir / "_preview_Cap07_Fig01.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
