"""
Cap10 Fig03 (PT-BR, B&W vetorial) — A janela dura cerca de 10 anos.
Timeline da janela de reposição hormonal pós-menopausa.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, FancyBboxPatch, FancyArrowPatch

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
ZONE_OK  = "#F4F4F4"   # PROTEÇÃO — claro
ZONE_MID = "#D8D8D8"   # DECISÃO INDIVIDUAL — médio
ZONE_BAD = "#BABABA"   # RISCO — escuro
BAND     = "#EDEDED"

fig = plt.figure(figsize=(11.0, 7.5))
fig.patch.set_facecolor(BG)

# Título
fig.text(0.025, 0.945, "Figura 3 — A janela dura cerca de 10 anos.",
         fontsize=16, color=INK, weight="bold")
fig.text(0.025, 0.905,
         "Começar a reposição hormonal dentro dela protege. Começar fora pode fazer o oposto.",
         fontsize=10, color=INK_SOFT, style="italic")

# ---------- Caixa Fernanda no topo ----------
FER_X1, FER_X2 = 0.025, 0.31
FER_Y1, FER_Y2 = 0.795, 0.860
fig.patches.append(FancyBboxPatch(
    (FER_X1, FER_Y1), FER_X2 - FER_X1, FER_Y2 - FER_Y1,
    boxstyle="round,pad=0.005,rounding_size=0.005",
    facecolor=BAND, edgecolor=INK, linewidth=0.7,
    transform=fig.transFigure, zorder=1
))
fig.text(0.035, 0.842, "● Fernanda, 44 anos",
         fontsize=10, color=INK, weight="bold")
fig.text(0.035, 0.823, "FSH 38, estradiol 28 pg/mL.",
         fontsize=8.5, color=INK)
fig.text(0.035, 0.807, "Na janela.",
         fontsize=9, color=INK, weight="bold", style="italic")

# ---------- 3 zonas (banda horizontal) ----------
ZONE_Y1, ZONE_Y2 = 0.55, 0.72
# Limites em X (escala 0 a 20 anos): 0->0.10, 10->0.55, 15->0.78, 20->0.95
def years_to_x(y):
    return 0.10 + (y / 20.0) * 0.85

# Zona 1: PROTEÇÃO (0-10)
fig.patches.append(Rectangle(
    (years_to_x(0), ZONE_Y1), years_to_x(10) - years_to_x(0), ZONE_Y2 - ZONE_Y1,
    facecolor=ZONE_OK, edgecolor=INK, linewidth=0.5,
    transform=fig.transFigure, zorder=1
))
# Zona 2: DECISÃO INDIVIDUAL (10-15)
fig.patches.append(Rectangle(
    (years_to_x(10), ZONE_Y1), years_to_x(15) - years_to_x(10), ZONE_Y2 - ZONE_Y1,
    facecolor=ZONE_MID, edgecolor=INK, linewidth=0.5,
    transform=fig.transFigure, zorder=1
))
# Zona 3: RISCO (15-20)
fig.patches.append(Rectangle(
    (years_to_x(15), ZONE_Y1), years_to_x(20) - years_to_x(15), ZONE_Y2 - ZONE_Y1,
    facecolor=ZONE_BAD, edgecolor=INK, linewidth=0.5,
    transform=fig.transFigure, zorder=1
))

# Labels das zonas
fig.text((years_to_x(0) + years_to_x(10)) / 2, 0.705, "PROTEÇÃO",
         fontsize=11, color=INK, weight="bold", ha="center", va="top")
fig.text((years_to_x(10) + years_to_x(15)) / 2, 0.705, "DECISÃO INDIVIDUAL",
         fontsize=10, color=INK, weight="bold", ha="center", va="top")
fig.text((years_to_x(15) + years_to_x(20)) / 2, 0.705, "RISCO",
         fontsize=11, color=INK, weight="bold", ha="center", va="top")

# Descrição de cada zona (texto na zona)
fig.text((years_to_x(0) + years_to_x(10)) / 2, 0.625,
         "Estradiol transdérmico reduz\nprogressão de aterosclerose,\nprotege osso e preserva\nfunção cognitiva.",
         fontsize=8, color=INK, ha="center", va="center",
         linespacing=1.3)
fig.text((years_to_x(10) + years_to_x(15)) / 2, 0.625,
         "Evidência menos clara.\nDecisão caso a caso com\nestratificação de risco.",
         fontsize=8, color=INK, ha="center", va="center",
         linespacing=1.3)
fig.text((years_to_x(15) + years_to_x(20)) / 2, 0.625,
         "Aterosclerose estabelecida\ntorna o estradiol\npotencialmente lesivo.\nConsiderar alternativas\nnão-hormonais.",
         fontsize=8, color=INK, ha="center", va="center",
         linespacing=1.3)

# ---------- Linha vertical tracejada em "10 anos" (limite da janela) ----------
x_10 = years_to_x(10)
fig.lines.append(plt.Line2D(
    [x_10, x_10], [0.28, 0.72],
    color=INK, linewidth=1.0, linestyle=(0, (4, 3)),
    transform=fig.transFigure, zorder=4
))
# Label da linha — DENTRO da banda das zonas (na parte superior),
# entre o título e a descrição
fig.text(x_10 + 0.005, 0.685, "← 10 anos\nlimite da janela",
         fontsize=8, color=INK, weight="bold", ha="left", va="top",
         linespacing=1.2)

# ---------- Eixo de tempo (seta) ----------
AXIS_Y = 0.475
fig.patches.append(FancyArrowPatch(
    (years_to_x(0), AXIS_Y), (years_to_x(20) + 0.01, AXIS_Y),
    arrowstyle="->", color=INK, lw=1.5, mutation_scale=20,
    transform=fig.transFigure, zorder=3
))
# Ticks
for yr in [0, 5, 10, 15, 20]:
    x = years_to_x(yr)
    fig.lines.append(plt.Line2D(
        [x, x], [AXIS_Y - 0.010, AXIS_Y + 0.010],
        color=INK, linewidth=1.0, transform=fig.transFigure, zorder=4
    ))
    fig.text(x, AXIS_Y - 0.020, str(yr),
             fontsize=9, color=INK, weight="bold", ha="center", va="top")

fig.text(0.50, AXIS_Y - 0.045, "Anos desde a última menstruação",
         fontsize=9, color=INK_SOFT, ha="center", style="italic")

# ---------- Trials no rodapé (caixinhas com seta) ----------
trials = [
    (3,   "ELITE (2016)\ngrupo precoce", "Reduziu progressão\nde aterosclerose."),
    (7,   "KEEPS (2024)\nfollow-up",     "Confirma a hipótese\nem seguimento longo."),
    (12,  "ELITE (2016)\ngrupo tardio",  "Não reduziu, em\nalguns recortes,\npiorou."),
    (17,  "WHI (2002)",                  "Não reduziu, em\nalgumas\ngeração delas."),
]

for yr, title, desc in trials:
    x = years_to_x(yr)
    # Caixa
    BOX_W, BOX_H = 0.12, 0.16
    BOX_Y_TOP = 0.39
    fig.patches.append(FancyBboxPatch(
        (x - BOX_W/2, BOX_Y_TOP - BOX_H), BOX_W, BOX_H,
        boxstyle="round,pad=0.003,rounding_size=0.005",
        facecolor="white", edgecolor=INK, linewidth=0.6,
        transform=fig.transFigure, zorder=2
    ))
    fig.text(x, BOX_Y_TOP - 0.025, title,
             fontsize=8, color=INK, weight="bold",
             ha="center", va="center", linespacing=1.2)
    fig.text(x, BOX_Y_TOP - 0.075, desc,
             fontsize=7, color=INK_SOFT,
             ha="center", va="center", linespacing=1.2)
    # Seta do trial até o eixo
    fig.patches.append(FancyArrowPatch(
        (x, BOX_Y_TOP + 0.003), (x, AXIS_Y - 0.015),
        arrowstyle="-", color=INK, lw=0.6,
        transform=fig.transFigure, zorder=2
    ))

# Footer
fig.text(0.025, 0.075,
         "Fonte: Hodis et al., ELITE, NEJM 2016; Harman et al., KEEPS primary, Annals 2014;",
         fontsize=7.5, color=FOOT)
fig.text(0.025, 0.055,
         "Miller et al., KEEPS continuation, PLOS Medicine 2024; Rossouw et al., WHI primary, JAMA 2002.",
         fontsize=7.5, color=FOOT, style="italic")

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap10_Fig03.pdf"
png_path = out_dir / "_preview_Cap10_Fig03.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
