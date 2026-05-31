"""
Cap06 Fig04 (PT-BR, B&W vetorial) — O TOTG de André: quando o jejum mente.
Curva dupla glicose + insulina ao longo do TOTG (0-120 min).
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle
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
GLI_COLOR = "#000000"   # glicose — preto sólido
INS_COLOR = "#666666"   # insulina — cinza médio (linha tracejada)
BAND     = "#F4F4F4"

fig = plt.figure(figsize=(11.0, 7.0))
fig.patch.set_facecolor(BG)

# Título
fig.text(0.025, 0.945, "Figura 4 — O TOTG de André: quando o jejum mente.",
         fontsize=17, color=INK, weight="bold")
fig.text(0.025, 0.895,
         "Glicose aparentemente normal com resposta insulínica desproporcional.",
         fontsize=10, color=INK_SOFT)

# ---------- 2 axes lado a lado pra dupla escala ----------
ax1 = fig.add_axes([0.10, 0.18, 0.80, 0.62])
ax2 = ax1.twinx()

x = [0, 30, 60, 90, 120]
glicose = [92, 148, 162, 154, 131]
insulina = [8.5, 78, 124, 118, 89]

# Banda "faixa normal pós-prandial" (até 140 mg/dL)
ax1.axhspan(70, 140, facecolor=BAND, zorder=0)
ax1.text(2, 73, "Faixa normal pós-prandial (≤ 140 mg/dL)",
         fontsize=8, color=INK_SOFT, ha="left", va="bottom", style="italic",
         linespacing=1.2)

# Glicose (linha sólida + círculos)
ax1.plot(x, glicose, color=GLI_COLOR, linewidth=2.4, marker="o",
         markersize=8, markerfacecolor=GLI_COLOR, markeredgecolor=GLI_COLOR,
         label="Glicose (mg/dL)", zorder=5)

# Insulina (linha tracejada + quadrados)
ax2.plot(x, insulina, color=INS_COLOR, linewidth=2.0,
         linestyle=(0, (5, 3)), marker="s",
         markersize=7, markerfacecolor="white", markeredgecolor=INS_COLOR,
         markeredgewidth=2, label="Insulina (µIU/mL)", zorder=5)

# Anotação de valores
for xv, gv, iv in zip(x, glicose, insulina):
    ax1.annotate(f"{gv}", (xv, gv), textcoords="offset points",
                 xytext=(0, 8), ha="center", fontsize=9, color=INK,
                 weight="bold")
    # t=0 caso especial — insulin label sai do caminho da glicose label e da banda
    if xv == 0:
        ax2.annotate(f"{iv}", (xv, iv), textcoords="offset points",
                     xytext=(12, -4), ha="left", fontsize=9, color=INK_SOFT,
                     weight="bold")
    else:
        ax2.annotate(f"{iv}", (xv, iv), textcoords="offset points",
                     xytext=(0, -16), ha="center", fontsize=9, color=INK_SOFT,
                     weight="bold")

# Marcadores "(pico)" — acima dos valores no pico (tempo 60)
ax1.text(58, 178, "(pico)", fontsize=8, color=INK, ha="right", style="italic")
# Anotação dirigida pra glicose: "Glicose moderada" — abaixo da curva, lado direito
ax1.text(110, 175, "Glicose moderada",
         fontsize=9, color=INK, weight="bold", ha="center")
# "Insulina desproporcional" — acima do pico de insulina (no espaço sobre)
ax2.text(60, 152, "Insulina desproporcional",
         fontsize=9, color=INK, weight="bold", ha="center")

# Configurar eixos
ax1.set_xlabel("Tempo (min)", fontsize=10, color=INK, weight="bold")
ax1.set_ylabel("Glicose (mg/dL)", fontsize=10, color=INK, weight="bold")
ax2.set_ylabel("Insulina (µIU/mL)", fontsize=10, color=INK, weight="bold")

ax1.set_xticks(x)
ax1.set_xlim(-5, 130)
ax1.set_ylim(60, 200)
ax2.set_ylim(0, 160)

ax1.tick_params(axis="both", colors=TICK, labelsize=9)
ax2.tick_params(axis="y", colors=TICK, labelsize=9)
for spine in ("top",):
    ax1.spines[spine].set_visible(False)
    ax2.spines[spine].set_visible(False)
for spine_name in ("left", "bottom", "right"):
    if spine_name in ax1.spines:
        ax1.spines[spine_name].set_color("#888888")
    if spine_name in ax2.spines:
        ax2.spines[spine_name].set_color("#888888")

# Legendas das curvas (canto superior)
lgd_lines = [
    plt.Line2D([0], [0], color=GLI_COLOR, linewidth=2.4, marker="o",
               markersize=6, markerfacecolor=GLI_COLOR, label="Glicose (mg/dL)"),
    plt.Line2D([0], [0], color=INS_COLOR, linewidth=2.0,
               linestyle=(0, (5, 3)), marker="s", markersize=6,
               markerfacecolor="white", markeredgecolor=INS_COLOR,
               markeredgewidth=2, label="Insulina (µIU/mL)"),
]
ax1.legend(handles=lgd_lines, loc="upper right", frameon=False, fontsize=9)

# Footer
fig.text(0.025, 0.060,
         "Fonte: dados reconstruídos do caso-tipo apresentado no Capítulo 6.",
         fontsize=8, color=FOOT)
fig.text(0.025, 0.035,
         "Padrão de referência: Kraft JR (Detection of Diabetes Mellitus In Situ, 2008).",
         fontsize=8, color=FOOT)

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap06_Fig04.pdf"
png_path = out_dir / "_preview_Cap06_Fig04.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
