"""
Cap08 Fig01 (PT-BR, B&W vetorial) — Marcos, 8 meses depois: a aptidão que vale por uma estatina.
2 painéis: bar chart baseline vs 8 meses, e curva dose-resposta.
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
GRAY     = "#9A9A9A"
DARK     = "#000000"

fig = plt.figure(figsize=(11.0, 6.0))
fig.patch.set_facecolor(BG)

# Título
fig.text(0.025, 0.945,
         "Figura 1 — Marcos, 8 meses depois: a aptidão que vale por uma estatina.",
         fontsize=15, color=INK, weight="bold")

# ---------- PAINEL ESQUERDO: VO2 max ergometria ----------
fig.text(0.025, 0.870, "VO₂ MAX – ERGOMETRIA",
         fontsize=9, color=INK_SOFT, weight="bold")

ax1 = fig.add_axes([0.04, 0.30, 0.40, 0.50])
ax1.set_xlim(0, 50)
ax1.set_ylim(-1, 2)
ax1.axis("off")

# Baseline (cinza)
ax1.barh([1.3], [28], height=0.45, color=GRAY, edgecolor="none")
ax1.text(-1, 1.3, "BASELINE", fontsize=10, color=INK_SOFT,
         weight="bold", ha="right", va="center")
ax1.text(29, 1.3, "28 ml/kg/min", fontsize=10, color=INK,
         weight="bold", va="center")

# 8 meses (preto)
ax1.barh([-0.3], [34.3], height=0.45, color=DARK, edgecolor="none")
ax1.text(-1, -0.3, "8 MESES", fontsize=10, color=INK,
         weight="bold", ha="right", va="center")
ax1.text(35.3, -0.3, "34,3 ml/kg/min", fontsize=10, color=INK,
         weight="bold", va="center")

# Delta no meio
ax1.text(22, 0.5, "+1,8 MET",
         fontsize=22, color=INK, weight="bold", ha="center", va="center")

# ---------- PAINEL DIREITO: Risco de morte × aptidão ----------
fig.text(0.50, 0.870, "RISCO DE MORTE × APTIDÃO",
         fontsize=9, color=INK_SOFT, weight="bold")

ax2 = fig.add_axes([0.55, 0.30, 0.42, 0.50])

x = list(range(5))
y = [5.04, 2.10, 1.49, 1.29, 1.00]
labels = ["BAIXA\nbaseline", "ABAIXO", "ACIMA\n8 meses", "ALTA", "ELITE"]

# Curva dashed
ax2.plot(x, y, color=INK, linewidth=1.6, linestyle=(0, (4, 3)),
         marker="o", markersize=8, markerfacecolor=INK,
         markeredgecolor=INK, zorder=4)

# Marker no baseline (branco/oco)
ax2.plot([0], [5.04], "o", markersize=12,
         markerfacecolor="white", markeredgecolor=INK,
         markeredgewidth=2, zorder=5)
ax2.text(0, 5.50, "5,04", fontsize=10, color=INK,
         weight="bold", ha="center")

# Marker em "ACIMA" - filled black (8 meses do Marcos)
ax2.plot([2], [1.49], "o", markersize=12,
         markerfacecolor=INK, markeredgecolor=INK, zorder=5)

# Seta curvada do baseline ao ACIMA
ax2.annotate("", xy=(2, 1.55), xytext=(0.2, 4.9),
             arrowprops=dict(arrowstyle="->", color=INK, lw=2.0,
                             connectionstyle="arc3,rad=-0.35"))

# Linha 1.0x referência
ax2.axhline(1.0, color="#888888", linewidth=0.6, linestyle=(0, (2, 2)), zorder=1)
ax2.text(-0.4, 1.0, "1,00", fontsize=9, color=INK_SOFT, va="center")

ax2.set_xticks(x)
ax2.set_xticklabels(labels, fontsize=8.5, color=INK,
                    linespacing=1.15, weight="bold")
ax2.set_xlim(-0.5, 4.4)
ax2.set_ylim(0, 6.2)

for spine in ("top", "right"):
    ax2.spines[spine].set_visible(False)
ax2.spines["left"].set_color("#888888")
ax2.spines["bottom"].set_color("#888888")
ax2.tick_params(axis="x", colors=TICK)
ax2.tick_params(axis="y", colors=TICK, labelsize=8)
ax2.set_yticks([0, 1, 2, 3, 4, 5])

# Footer
fig.text(0.025, 0.18,
         "+1,8 MET  =  25–30% menos risco de morte por todas as causas.",
         fontsize=12, color=INK, weight="bold")
fig.text(0.025, 0.10,
         "Mandsager et al., JAMA Network Open 2018.",
         fontsize=8.5, color=FOOT, style="italic")

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap08_Fig01.pdf"
png_path = out_dir / "_preview_Cap08_Fig01.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
