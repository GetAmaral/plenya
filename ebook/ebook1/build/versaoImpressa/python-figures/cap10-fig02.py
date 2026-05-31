"""
Cap10 Fig02 (PT-BR, B&W vetorial) — IGF-1: a exceção da curva — nem muito baixo, nem muito alto.
Curva em U: risco relativo x IGF-1.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
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
ZONE_BAD = "#C8C8C8"     # zonas de risco — escuro
ZONE_MID = "#E0E0E0"     # subótimas — médio
ZONE_OK  = "#F4F4F4"     # zona ótima — claro

fig = plt.figure(figsize=(11.0, 7.2))
fig.patch.set_facecolor(BG)

# Título
fig.text(0.025, 0.945,
         "Figura 2 — IGF-1: a exceção da curva — nem muito baixo, nem muito alto",
         fontsize=15, color=INK, weight="bold")
fig.text(0.025, 0.908,
         "Risco relativo para longevidade em função de IGF-1 (ng/mL) — adultos de 40-50 anos.",
         fontsize=9.5, color=INK_SOFT, style="italic")

ax = fig.add_axes([0.10, 0.18, 0.85, 0.65])

# Limites
ax.set_xlim(60, 230)
ax.set_ylim(0, 3)

# Bandas verticais por zona
zones = [
    (60, 80,   ZONE_BAD),   # < 80: risco deficiência
    (80, 120,  ZONE_MID),   # 80-120: subótimo inferior
    (120, 160, ZONE_OK),    # 120-160: ZONA ÓTIMA
    (160, 200, ZONE_MID),   # 160-200: subótimo superior
    (200, 230, ZONE_BAD),   # > 200: risco excesso
]
for x1, x2, color in zones:
    ax.axvspan(x1, x2, facecolor=color, zorder=0)

# Labels das zonas no topo
zone_labels_top = [
    (70,  "< 80\nRisco por\ndeficiência"),
    (100, "80 – 120\nSubótimo\ninferior"),
    (140, "120 – 160\nZONA ÓTIMA\npara longevidade"),
    (180, "160 – 200\nSubótimo\nsuperior"),
    (215, "> 200\nRisco por\nexcesso"),
]
for x, lbl in zone_labels_top:
    weight = "bold" if "ZONA ÓTIMA" in lbl else "normal"
    color  = INK if "ZONA ÓTIMA" in lbl else INK_SOFT
    ax.text(x, 2.80, lbl, fontsize=8.5, color=color,
            ha="center", va="top", weight=weight, linespacing=1.3)

# Curva em U (rscaled)
x_curve = np.linspace(60, 230, 100)
# U shape via parabola normalizada com mínimo em 140
y_curve = 0.0014 * (x_curve - 140) ** 2 + 0.4
ax.plot(x_curve, y_curve, color=INK, linewidth=2.6, zorder=5)

# Labels do eixo Y (texto, não números) — invertido conceitualmente
ax.set_yticks([0.5, 1.5, 2.5])
ax.set_yticklabels(["BAIXO", "MODERADO", "ALTO"],
                    fontsize=9, color=INK, weight="bold")

# X axis
ax.set_xticks([60, 80, 100, 120, 140, 160, 180, 200, 220, 230])
ax.tick_params(axis="x", colors=TICK, labelsize=8.5)
ax.tick_params(axis="y", colors=TICK, labelsize=9)

ax.set_xlabel("IGF-1 (ng/mL)", fontsize=10, color=INK, weight="bold")
ax.set_ylabel("RISCO RELATIVO", fontsize=10, color=INK, weight="bold")

for spine in ("top", "right"):
    ax.spines[spine].set_visible(False)
ax.spines["left"].set_color("#888888")
ax.spines["bottom"].set_color("#888888")

# ---------- Anotações abaixo do gráfico ----------
ANNOT_Y = 0.110

annotations = [
    (0.18, "Fragilidade, sarcopenia,\ndeclínio funcional"),
    (0.50, "Zona ótima para longevidade"),
    (0.82, "Aceleração proliferativa,\nrisco oncológico"),
]
for x, txt in annotations:
    weight = "bold" if "Zona ótima" in txt else "normal"
    fig.text(x, ANNOT_Y, txt,
             fontsize=9, color=INK, ha="center", va="center",
             weight=weight, linespacing=1.3)

# Footer
fig.text(0.025, 0.045,
         "Ao contrário da maioria dos biomarcadores discutidos neste livro, o IGF-1 não segue a lógica \"quanto mais baixo, melhor\". Níveis muito baixos se associam",
         fontsize=7.5, color=FOOT)
fig.text(0.025, 0.025,
         "à fragilidade e perda de função; níveis muito altos se associam à proliferação celular e maior risco oncológico. Valores aproximados para adultos de 40-50 anos.",
         fontsize=7.5, color=FOOT)

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap10_Fig02.pdf"
png_path = out_dir / "_preview_Cap10_Fig02.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
