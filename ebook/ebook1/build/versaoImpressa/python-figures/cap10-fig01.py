"""
Cap10 Fig01 (PT-BR, B&W vetorial) — A testosterona de um homem de 48 — e a
de um homem de 80. O mesmo número.

Gráfico de linha mostrando média populacional de testosterona total vs idade.
Bandas: ZONA ÓTIMA (y>500, cinza claro hatched) e HIPOGONADISMO (y<300, cinza
mais escuro). Paulo destacado em (48, 310) — nível típico de 80 anos.
"""
from pathlib import Path as _Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import (
    Rectangle, FancyBboxPatch, Circle, FancyArrowPatch, Polygon
)
import numpy as np

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False
rcParams["pdf.fonttype"] = 42
rcParams["ps.fonttype"] = 42

BG    = "#FFFFFF"
INK   = "#000000"
SOFT  = "#3A3A3A"
FOOT  = "#6A6A6A"
GRAY1 = "#9E9E9E"
GRAY2 = "#D9D9D9"
GRAY3 = "#EEEEEE"
ZONE_OK   = "#F2F2F2"   # zona ótima (era verde)
ZONE_HIPO = "#E0E0E0"   # zona hipogonadismo (era vermelha)

W_IMG, H_IMG = 1536, 1024
_FIG_W = 10.0
_FIG_H = _FIG_W * H_IMG / W_IMG
fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

# Eixo full-figure em image px
ax_bg = fig.add_axes([0, 0, 1, 1])
ax_bg.set_xlim(0, W_IMG); ax_bg.set_ylim(H_IMG, 0)
ax_bg.set_aspect("equal"); ax_bg.axis("off")

def text_with_width(txt, x, y, **kw):
    t = ax_bg.text(x, y, txt, **kw)
    fig.canvas.draw()
    bbox = t.get_window_extent(renderer=fig.canvas.get_renderer())
    inv = ax_bg.transData.inverted()
    return inv.transform(bbox)[1][0]

# ============================================================
# CABEÇALHO
# ============================================================
ax_bg.add_patch(Rectangle((22, 18), 156-22, 45-18, facecolor=INK, edgecolor="none"))
ax_bg.text((22+156)/2, (18+45)/2, "FIGURA 1",
           fontsize=11, color="white", weight="bold", va="center", ha="center")

ax_bg.text(45, 100, "A testosterona de um homem de 48 — e a de um homem de 80.",
           fontsize=22, color=INK, weight="bold", va="center", ha="left")
ax_bg.text(45, 140, "O mesmo número.",
           fontsize=22, color=INK, weight="bold", va="center", ha="left")
ax_bg.text(45, 180,
           "Médias populacionais de testosterona total (ng/dL) ao longo da vida adulta.",
           fontsize=11, color=SOFT, va="center", ha="left")

# ============================================================
# CHART AREA — dedicated ax with data coords
# ============================================================
# Chart area in image px: roughly x=110..1430, y=240..830
# Use fig.add_axes with normalized coords
chart_left_px  = 110
chart_right_px = 1430
# MEDIDO pelos rótulos do eixo y da arte original: as onze marcas vão de 750
# em y 0.267 a 250 em y 0.707 (de 1024 px). A caixa resolvida a partir disso é
# 264..769 (com ylim 200–760); estava em 240..830, o que esticava o gráfico para baixo e fazia o
# colchete "+30 anos" encavalar nos rótulos do eixo.
chart_top_px   = 264
chart_bot_px   = 769

# Convert to fig fraction (note: matplotlib origin bottom-left)
ax_x0 = chart_left_px / W_IMG
ax_x1 = chart_right_px / W_IMG
ax_y0 = 1 - chart_bot_px / H_IMG    # bottom (in fig coords)
ax_y1 = 1 - chart_top_px / H_IMG    # top
ax = fig.add_axes([ax_x0, ax_y0, ax_x1 - ax_x0, ax_y1 - ax_y0])
ax.set_xlim(15, 87)
ax.set_ylim(200, 760)
ax.spines["top"].set_visible(False)
ax.spines["right"].set_visible(False)
ax.spines["left"].set_color(GRAY1)
ax.spines["bottom"].set_color(GRAY1)
ax.tick_params(axis="both", colors=SOFT, labelsize=9)
ax.set_xticks([20, 30, 40, 50, 60, 70, 80, 85])
ax.set_yticks([250, 300, 350, 400, 450, 500, 550, 600, 650, 700, 750])
ax.grid(False)

# ZONA ÓTIMA (y > 500) — banda cinza claro
ax.axhspan(500, 760, facecolor=ZONE_OK, alpha=0.9, zorder=1)
# HIPOGONADISMO (y < 320) — banda cinza médio
ax.axhspan(200, 320, facecolor=ZONE_HIPO, alpha=0.9, zorder=1)

# Linhas tracejadas nas bordas das zonas
ax.axhline(500, linestyle=(0, (4, 4)), color=INK, linewidth=1.0,
           alpha=0.6, zorder=2)
ax.axhline(320, linestyle=(0, (4, 4)), color=INK, linewidth=1.0,
           alpha=0.6, zorder=2)

# Data points (idade, testosterona)
AGES = [20, 30, 40, 50, 60, 70, 80]
TEST = [680, 610, 540, 470, 410, 350, 310]

# Linha de tendência (excluindo Paulo)
ax.plot(AGES, TEST, color=SOFT, linewidth=2.2, zorder=4)
ax.scatter(AGES, TEST, s=70, color=SOFT, edgecolor=BG, linewidth=2,
           zorder=5)

# Labels dos pontos (valor acima do ponto)
for age, val in zip(AGES, TEST):
    ax.text(age, val + 22, str(val),
            fontsize=10, color=SOFT, weight="bold",
            ha="center", va="bottom", zorder=6)

# Paulo highlight em (48, 310) — círculo preto cheio + halo branco
ax.add_patch(Circle((48, 310), 0.9 * 1.3, facecolor="white",
                     edgecolor="none", zorder=7,
                     transform=ax.transData))
ax.scatter([48], [310], s=180, color=INK, edgecolor=BG, linewidth=2.5,
           zorder=8)

# Linha vertical pontilhada de "48" no x-axis pra (48, 310)
ax.plot([48, 48], [200, 310], linestyle=(0, (2, 3)),
        color=INK, linewidth=1.0, alpha=0.7, zorder=3)

# ============================================================
# LABELS DAS ZONAS (dentro do gráfico)
# ============================================================
# ZONA ÓTIMA (canto sup-esquerdo da banda)
# Posicionado BAIXO no zone pra ficar embaixo da linha de tendência;
# x maior pra não ficar sobre o label do eixo Y.
# Linha de tendência: (20,680)→(30,610)→(40,540)→(50,470)...
# A faixa entre y=510 e y=620 fica livre da linha quando x está entre 17 e 30.
zx = 17
ax.text(zx, 615, "ZONA ÓTIMA",
        fontsize=10.5, color=INK, weight="bold", va="center", ha="left", zorder=6)
ax.text(zx, 590, "> 500 ng/dL",
        fontsize=9, color=SOFT, va="center", ha="left", zorder=6)
ax.text(zx, 565, "Associação com menor",
        fontsize=8.5, color=SOFT, va="center", ha="left", zorder=6)
ax.text(zx, 545, "mortalidade em estudos",
        fontsize=8.5, color=SOFT, va="center", ha="left", zorder=6)
ax.text(zx, 525, "de longevidade.",
        fontsize=8.5, color=SOFT, va="center", ha="left", zorder=6)

# HIPOGONADISMO (canto inf-esquerdo da banda)
ax.text(17, 280, "HIPOGONADISMO",
        fontsize=10, color=INK, weight="bold", va="center", ha="left", zorder=6)
ax.text(17, 260, "LABORATORIAL",
        fontsize=10, color=INK, weight="bold", va="center", ha="left", zorder=6)
ax.text(17, 240, "< 300 ng/dL",
        fontsize=9, color=SOFT, va="center", ha="left", zorder=6)

# ============================================================
# PAULO CALLOUT — label + arrow
# ============================================================
# Box texto à esquerda do ponto Paulo
ax.text(40, 380, "Paulo, 48 anos:",
        fontsize=11, color=INK, weight="bold",
        ha="center", va="center", zorder=8)
ax.text(40, 355, "310 ng/dL",
        fontsize=11, color=INK, weight="bold",
        ha="center", va="center", zorder=8)
# Seta curva do label pro ponto
arrow = FancyArrowPatch((43, 345), (47.3, 318),
                       arrowstyle="-|>,head_length=8,head_width=6",
                       connectionstyle="arc3,rad=-0.3",
                       color=INK, linewidth=1.2, zorder=8)
ax.add_patch(arrow)

# Label "48" em bold sob o eixo
ax.text(48, 192, "48",
        fontsize=10, color=INK, weight="bold",
        ha="center", va="top", zorder=6, transform=ax.transData,
        clip_on=False)

# X-axis label
# Será desenhado pelo ax_bg

# ============================================================
# VALOR TÍPICO AOS 80 ANOS (label na zona hipogonadismo)
# ============================================================
ax.text(67, 305, "Valor típico aos 80 anos",
        fontsize=9, color=SOFT, style="italic",
        ha="center", va="center", zorder=6)

# ============================================================
# +30 anos de envelhecimento hormonal — bracket horizontal
# ============================================================
# Bracket de age=48 ao age=80, abaixo do eixo x (em image coords via ax_bg)
# Mapear: age=48 → x_px ; age=80 → x_px
# ax range 15..87 in x, output px range chart_left_px..chart_right_px = 110..1430
def age_to_px(age):
    return chart_left_px + (age - 15) / (87 - 15) * (chart_right_px - chart_left_px)

x48 = age_to_px(48)
x80 = age_to_px(80)
y_brk_top = 875
y_brk_mid = 890
y_brk_label = 905

ax_bg.plot([x48, x48], [y_brk_top, y_brk_mid], color=INK, linewidth=1.0)
ax_bg.plot([x80, x80], [y_brk_top, y_brk_mid], color=INK, linewidth=1.0)
ax_bg.plot([x48, x80], [y_brk_mid, y_brk_mid], color=INK, linewidth=1.0,
           linestyle=(0, (5, 3)))
ax_bg.text((x48 + x80) / 2, y_brk_label,
           "+30 anos de envelhecimento hormonal",
           fontsize=11, color=INK, weight="bold",
           va="center", ha="center")

# X-axis label IDADE (ANOS)
ax_bg.text(45, 870, "IDADE (ANOS)",
           fontsize=9, color=SOFT, weight="bold", va="center", ha="left")

# Y-axis label
ax_bg.text(110, 230, "TESTOSTERONA TOTAL (ng/dL)",
           fontsize=9, color=SOFT, weight="bold", va="center", ha="left")

# ============================================================
# SOURCE (rodapé)
# ============================================================
ax_bg.plot([45, W_IMG-45], [930, 930], color=GRAY1, linewidth=0.7)

ax_bg.text(45, 952,
           "Testosterona total média por faixa etária em homens adultos (dados de coortes populacionais consolidadas). O declínio de 1 a 2% ao ano após os 30 é",
           fontsize=8, color=FOOT, va="center", ha="left")
ax_bg.text(45, 970,
           "fisiológico — mas trajetórias acentuadas levam pacientes relativamente jovens a valores tipicamente observados em idosos. A faixa verde indica valores",
           fontsize=8, color=FOOT, va="center", ha="left")
ax_bg.text(45, 988,
           "associados a menor mortalidade em estudos de longevidade; a faixa vermelha indica hipogonadismo laboratorial formal.",
           fontsize=8, color=FOOT, va="center", ha="left")

# Source line bold "Fontes:"
fonte_x = 45
fonte_y = 1010
end_b = text_with_width("Fontes:", fonte_x, fonte_y,
                        fontsize=8.5, color=FOOT, weight="bold",
                        va="center", ha="left")
ax_bg.text(end_b + 4, fonte_y,
           "NHANES III (1988–1994); European Male Ageing Study; Travison et al., JCEM, 2007; Wu et al., PLoS ONE, 2010.",
           fontsize=8.5, color=FOOT, va="center", ha="left")

# ============================================================
# EXPORT
# ============================================================
out_dir = _Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap10_Fig01.pdf"
png_path = out_dir / "_preview_Cap10_Fig01.png"
plt.savefig(pdf_path, facecolor=BG, bbox_inches="tight", pad_inches=0.0)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches="tight", pad_inches=0.0)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
