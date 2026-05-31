"""
Cap10 Fig01 (PT-BR, B&W vetorial) — A testosterona de um homem de 48 — e a de um homem de 80.
O mesmo número.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import FancyArrowPatch

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
BAND_OK  = "#F4F4F4"   # zona ótima — claro
BAND_BAD = "#C8C8C8"   # hipogonadismo — escuro

fig = plt.figure(figsize=(11.0, 7.2))
fig.patch.set_facecolor(BG)

# Título
fig.text(0.025, 0.945,
         "Figura 1 — A testosterona de um homem de 48 — e a de um homem de 80.",
         fontsize=14, color=INK, weight="bold")
fig.text(0.025, 0.910, "O mesmo número.",
         fontsize=14, color=INK, weight="bold")
fig.text(0.025, 0.870,
         "Médias populacionais de testosterona total (ng/dL) ao longo da vida adulta.",
         fontsize=9.5, color=INK_SOFT, style="italic")

ax = fig.add_axes([0.10, 0.18, 0.82, 0.62])

ages = [20, 30, 40, 50, 60, 70, 80]
testo = [680, 610, 560, 470, 410, 350, 310]

# Bandas de fundo
ax.axhspan(500, 750, facecolor=BAND_OK, zorder=0)
ax.axhspan(100, 300, facecolor=BAND_BAD, zorder=0)

# Labels das bandas
ax.text(85, 620, "ZONA ÓTIMA\n> 500 ng/dL\nAssociada com menor\nmortalidade em estudos\nde longevidade.",
        fontsize=8, color=INK_SOFT, ha="right", va="center",
        linespacing=1.25, style="italic")
ax.text(85, 200, "HIPOGONADISMO\nLABORATORIAL\n< 300 ng/dL",
        fontsize=8, color=INK_SOFT, ha="right", va="center",
        linespacing=1.25, style="italic")

# Curva
ax.plot(ages, testo, color=INK, linewidth=2.4, marker="o",
        markersize=9, markerfacecolor=INK, markeredgecolor=INK, zorder=5)

# Valores acima de cada ponto
for x, y in zip(ages, testo):
    ax.text(x, y + 25, str(y),
            fontsize=9.5, color=INK, weight="bold", ha="center", va="bottom")

# Marker do Paulo (48 anos, 310 ng/dL) — destaque
PAULO_AGE = 48
PAULO_T   = 310

# Linha pontilhada vertical
ax.axvline(PAULO_AGE, color=INK, linewidth=0.9, linestyle=(0, (3, 2)),
           zorder=2, alpha=0.7)

# Marker
ax.plot([PAULO_AGE], [PAULO_T], "o", markersize=14,
        markerfacecolor=INK, markeredgecolor=INK, zorder=6)

# Label do Paulo (à esquerda do marker, posicionado abaixo da curva)
ax.annotate("Paulo, 48 anos:\n310 ng/dL",
            xy=(PAULO_AGE - 1, PAULO_T), xytext=(35, 380),
            fontsize=10, color=INK, weight="bold", ha="center",
            arrowprops=dict(arrowstyle="->", color=INK, lw=0.9),
            linespacing=1.2)

# Label "Valor típico aos 80 anos" — acima da curva à direita
ax.annotate("Valor típico\naos 80 anos",
            xy=(80, 310), xytext=(78, 450),
            fontsize=9, color=INK_SOFT, ha="center", style="italic",
            arrowprops=dict(arrowstyle="->", color=INK_SOFT, lw=0.8),
            linespacing=1.2)

# Seta horizontal "+30 anos de envelhecimento hormonal" — DEBAIXO da curva
ax.annotate("", xy=(80, 165), xytext=(48, 165),
            arrowprops=dict(arrowstyle="<->", color=INK, lw=1.5))
ax.text((48 + 80) / 2, 150, "+30 anos de envelhecimento hormonal",
        fontsize=10, color=INK, weight="bold", ha="center", va="top")

# Eixos
ax.set_xlabel("IDADE (ANOS)", fontsize=10, color=INK, weight="bold")
ax.set_ylabel("TESTOSTERONA TOTAL (ng/dL)", fontsize=10, color=INK, weight="bold")
ax.set_xlim(15, 88)
ax.set_ylim(100, 750)
ax.set_xticks([20, 30, 40, 50, 60, 70, 80, 85])
ax.set_yticks([100, 200, 300, 400, 500, 600, 700, 750])

for spine in ("top", "right"):
    ax.spines[spine].set_visible(False)
ax.spines["left"].set_color("#888888")
ax.spines["bottom"].set_color("#888888")
ax.tick_params(axis="both", colors=TICK, labelsize=9)
ax.set_axisbelow(True)
ax.grid(axis="y", color="#EEEEEE", linewidth=0.5)

# Footer
fig.text(0.025, 0.080,
         "Testosterona total média por faixa etária em homens adultos (dados de coortes populacionais convencionais) — o declínio de 1 a 2% ao ano após os 30 é",
         fontsize=7.5, color=FOOT)
fig.text(0.025, 0.060,
         "fisiológico — mas trajetórias individuais variam bastante. Paulo cai precocemente jovem a valores tipicamente observados décadas depois — caixa verde indica valores",
         fontsize=7.5, color=FOOT)
fig.text(0.025, 0.040,
         "associados a menor mortalidade em estudos de longevidade; a faixa vermelha indica hipogonadismo laboratorial formal.",
         fontsize=7.5, color=FOOT)
fig.text(0.025, 0.018,
         "Fontes: NHANES III (1988-1994); European Male Aging Study; Travison et al., JCEM, 2007; Wu et al., PLoS ONE, 2010.",
         fontsize=7.5, color=FOOT, style="italic")

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap10_Fig01.pdf"
png_path = out_dir / "_preview_Cap10_Fig01.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
