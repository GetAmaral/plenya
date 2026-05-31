"""
Cap07 Fig02 (PT-BR, B&W vetorial) — O maior retorno sobre o investimento está no primeiro passo.
Curva dose-resposta de mortalidade x aptidão cardiorrespiratória.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle

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
HIGHLIGHT = "#E0E0E0"

fig = plt.figure(figsize=(11.0, 6.8))
fig.patch.set_facecolor(BG)

# Título
fig.text(0.025, 0.945, "Figura 2 — O maior retorno sobre o investimento está no primeiro passo",
         fontsize=15, color=INK, weight="bold")
fig.text(0.025, 0.905,
         "Risco ajustado de mortalidade por todas as causas de acordo com a aptidão cardiorrespiratória.",
         fontsize=9.5, color=INK_SOFT, style="italic")

ax = fig.add_axes([0.08, 0.22, 0.86, 0.62])

x = list(range(5))
y = [5.04, 2.10, 1.49, 1.29, 1.00]
labels = ["Baixa\n(≤ P25)",
          "Abaixo da média\n(P25-P49)",
          "Acima da média\n(P50-P74)",
          "Alta\n(P75-P97.6)",
          "Elite\n(≥ P97.7)"]

# Banda destacando o maior ganho (entre x=0 e x=1)
ax.axvspan(-0.2, 1.0, facecolor=HIGHLIGHT, zorder=0)
ax.text(0.4, 5.3, "Maior ganho\nocorre aqui",
        fontsize=10, color=INK, weight="bold", ha="center", va="center",
        linespacing=1.15)

# Curva
ax.plot(x, y, color=INK, linewidth=2.6, marker="o",
        markersize=10, markerfacecolor=INK, markeredgecolor=INK, zorder=5)

# Valores acima de cada ponto
for xv, yv in zip(x, y):
    ax.annotate(f"{yv:.2f}".replace(".", ","),
                (xv, yv), textcoords="offset points",
                xytext=(0, 12), ha="center",
                fontsize=11, color=INK, weight="bold")

# Linha de referência do tabagismo (1.41x)
ax.axhline(1.41, color=INK, linewidth=0.9, linestyle=(0, (3, 3)),
           zorder=2, alpha=0.7)
ax.text(4.1, 1.45, "Risco aproximado\ndo tabagismo (1,41x)",
        fontsize=8.5, color=INK_SOFT, ha="right", va="bottom",
        linespacing=1.2, style="italic")

ax.set_xticks(x)
ax.set_xticklabels(labels, fontsize=9, color=INK,
                   linespacing=1.15)
ax.set_xlim(-0.4, 4.4)
ax.set_ylim(0, 6.0)
ax.set_ylabel("Risco de morte\n(vezes maior)",
              fontsize=10, color=INK, weight="bold")
ax.set_xlabel("Nível de aptidão cardiorrespiratória (por percentil)",
              fontsize=10, color=INK, weight="bold")

for spine in ("top", "right"):
    ax.spines[spine].set_visible(False)
ax.spines["left"].set_color("#888888")
ax.spines["bottom"].set_color("#888888")
ax.tick_params(axis="both", colors=TICK, labelsize=9)
ax.set_axisbelow(True)
ax.grid(axis="y", color="#EEEEEE", linewidth=0.6)

# Footer
fig.text(0.025, 0.065,
         "Risco mais acentuado ocorre entre o grupo de menor aptidão e o seguinte —",
         fontsize=8.5, color=FOOT, style="italic")
fig.text(0.025, 0.045,
         "equivalente a seguir as diretrizes básicas de 150 minutos semanais de atividade moderada.",
         fontsize=8.5, color=FOOT, style="italic")
fig.text(0.025, 0.018,
         "Fonte: Mandsager et al., JAMA Network Open, 2018.",
         fontsize=8, color=FOOT)

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap07_Fig02.pdf"
png_path = out_dir / "_preview_Cap07_Fig02.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
