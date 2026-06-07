"""
Cap10 Fig02 (PT-BR, B&W vetorial) — IGF-1: a exceção da curva.

Posições medidas pixel-precisas do original 1536×1024 via crop-grid 400×400 +
máscaras de cor. Não há adivinhação: cada elemento usa as coords medidas.

Mapeamento principal:
  Chart x: IGF=40 → px=200, IGF=230 → px=1450 (linear, ~6.58 px/IGF)
  Chart y: y_top (ALTO) ≈ 320, y_bot (BAIXO) ≈ 610, x-axis em y=707
  Zonas:   value labels em y=234; descrições em y=275..305
  Y-axis arrow: x=175, y=62..649
"""
from pathlib import Path as _Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, FancyArrowPatch, Polygon, Circle
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
# Zonas B&W (preservando hierarquia: extremos mais escuros, ótima mais clara)
ZONE_BAD  = "#E2E2E2"
ZONE_SUB  = "#EEEEEE"
ZONE_OK   = "#F7F7F7"

W_IMG, H_IMG = 1536, 1024
_FIG_W = 10.0
_FIG_H = _FIG_W * H_IMG / W_IMG
fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

ax_bg = fig.add_axes([0, 0, 1, 1])
ax_bg.set_xlim(0, W_IMG); ax_bg.set_ylim(H_IMG, 0)
ax_bg.set_aspect("equal"); ax_bg.axis("off")

# Mapeamento IGF → x_px (medido do original)
CHART_LEFT_PX = 200
CHART_RIGHT_PX = 1450
IGF_MIN = 40
IGF_MAX = 230
def igf_to_px(igf):
    return CHART_LEFT_PX + (igf - IGF_MIN) / (IGF_MAX - IGF_MIN) * (CHART_RIGHT_PX - CHART_LEFT_PX)

# Y-bounds chart (curve area)
Y_ALTO   = 320   # topo do gráfico (linha pontilhada ALTO)
Y_MODERADO = 475 # nível MODERADO
Y_BAIXO  = 605   # nível BAIXO
Y_XAXIS  = 707   # linha do eixo x
Y_CHART_TOP = 210  # zonas cobrem desde acima dos labels até o eixo X

# ============================================================
# CABEÇALHO
# ============================================================
ax_bg.add_patch(Rectangle((24, 15), 164-24, 45-15, facecolor=INK, edgecolor="none"))
ax_bg.text((24+164)/2, (15+45)/2, "FIGURA 2",
           fontsize=11, color="white", weight="bold", va="center", ha="center")

# Título: y=60..119 (medido) → fontsize 22pt
ax_bg.text(61, 90,
           "IGF-1: a exceção da curva — nem muito baixo, nem muito alto",
           fontsize=22, color=INK, weight="bold", va="center", ha="left")
# Subtítulo centralizado verticalmente entre fim do título (y=119) e início do gráfico (y=210)
ax_bg.text(59, 164,
           "Risco relativo para longevidade em função do IGF-1 (ng/mL) — adultos de 40–50 anos.",
           fontsize=11, color=SOFT, va="center", ha="left")

# ============================================================
# RISCO RELATIVO label (y=205, x=20..170)
# ============================================================
ax_bg.text(20, 210, "RISCO RELATIVO",
           fontsize=9, color=INK, weight="bold", va="center", ha="left")

# Y-axis arrow (vertical line x=175, y top=215, bot=649)
ax_bg.annotate("", xy=(175, 215), xytext=(175, 649),
               arrowprops=dict(arrowstyle="-|>", color=INK, linewidth=1.4))

# ============================================================
# ZONAS — bandas verticais de fundo + dividers
# Zonas spans em IGF: [40-80, 80-120, 120-160, 160-200, 200-230]
# ============================================================
ZONE_SPANS = [
    (40, 80,   ZONE_BAD),
    (80, 120,  ZONE_SUB),
    (120, 160, ZONE_OK),
    (160, 200, ZONE_SUB),
    (200, 230, ZONE_BAD),
]
for x0, x1, color in ZONE_SPANS:
    rect_x0 = igf_to_px(x0)
    rect_x1 = igf_to_px(x1)
    ax_bg.add_patch(Rectangle(
        (rect_x0, Y_CHART_TOP), rect_x1 - rect_x0, Y_XAXIS - Y_CHART_TOP,
        facecolor=color, edgecolor="none", zorder=1
    ))

# Dividers verticais entre zonas (tracejado)
for x_div in (80, 120, 160, 200):
    x_px = igf_to_px(x_div)
    ax_bg.plot([x_px, x_px], [Y_CHART_TOP, Y_XAXIS],
               linestyle=(0, (4, 4)), color=GRAY1, linewidth=0.8, zorder=2)

# ============================================================
# TOP ZONE LABELS — valores (y=234) + descrições (y=270, 295)
# Centros medidos do original (em px)
# ============================================================
ZONE_LABELS = [
    (60,  "< 80",      ["Risco por", "deficiência"]),
    (100, "80 – 120",  ["Subótimo", "inferior"]),
    (140, "120 – 160", ["ZONA ÓTIMA", "para longevidade"]),
    (180, "160 – 200", ["Subótimo", "superior"]),
    (215, "> 200",     ["Risco por", "excesso"]),
]
VAL_Y  = 234   # y do valor numérico
DESC_Y = 273   # y descrição linha 1
DESC2_Y= 298   # y descrição linha 2

for x_center_igf, val_lbl, desc_lines in ZONE_LABELS:
    cx = igf_to_px(x_center_igf)
    ax_bg.text(cx, VAL_Y, val_lbl,
               fontsize=13, color=INK, weight="bold",
               ha="center", va="center", zorder=10)
    if "ZONA ÓTIMA" in desc_lines[0]:
        ax_bg.text(cx, DESC_Y, desc_lines[0],
                   fontsize=10.5, color=INK, weight="bold",
                   ha="center", va="center", zorder=10)
        ax_bg.text(cx, DESC2_Y, desc_lines[1],
                   fontsize=10, color=INK,
                   ha="center", va="center", zorder=10)
    else:
        ax_bg.text(cx, DESC_Y, desc_lines[0],
                   fontsize=10, color=INK,
                   ha="center", va="center", zorder=10)
        ax_bg.text(cx, DESC2_Y, desc_lines[1],
                   fontsize=10, color=INK,
                   ha="center", va="center", zorder=10)

# ============================================================
# Y-AXIS LABELS — ALTO (y=316), MODERADO (y=475), BAIXO (y=608)
# ============================================================
Y_LABEL_X_END = 175    # label termina pouco antes do eixo Y
for label, y_px in [("ALTO", 316), ("MODERADO", 475), ("BAIXO", 608)]:
    ax_bg.text(Y_LABEL_X_END - 5, y_px, label,
               fontsize=10, color=INK, weight="bold",
               ha="right", va="center", zorder=10)
    # Linha horizontal tracejada no nível
    ax_bg.plot([CHART_LEFT_PX, CHART_RIGHT_PX], [y_px, y_px],
               linestyle=(0, (2, 5)), color=GRAY1, linewidth=0.6, zorder=2)

# ============================================================
# CURVA EM U — pontos medidos pixel-precisos do original
# 84 amostras detectadas via máscara navy. Interpolação por cubic spline
# pra suavidade.
# ============================================================
# (IGF, y_px) — amostras do original
CURVE_PTS = [
    (46.08, 331.0), (47.60, 334.0), (49.12, 337.0), (50.64, 340.5),
    (52.16, 344.0), (55.20, 352.5), (58.24, 362.5), (70.40, 413.5),
    (71.92, 420.5), (73.44, 428.5), (74.96, 436.5), (76.48, 445.0),
    (78.00, 453.0), (79.52, 462.0), (81.04, 470.0), (82.56, 479.0),
    (84.08, 487.0), (85.60, 496.0), (87.12, 504.5), (88.64, 513.0),
    (90.16, 521.0), (91.68, 529.5), (93.20, 537.5), (99.28, 568.5),
    (102.32, 583.5), (103.84, 590.0), (105.36, 596.5), (106.88, 603.0),
    (108.40, 609.5), (109.92, 615.0), (111.44, 621.0), (112.96, 626.0),
    (114.48, 631.5), (116.00, 636.0), (117.52, 641.0), (119.04, 645.0),
    (120.56, 649.0), (122.08, 652.0), (123.60, 655.5), (128.16, 663.5),
    (131.20, 666.5), (132.72, 667.5), (140.32, 668.5), (144.88, 664.0),
    (146.40, 662.0), (150.96, 653.5), (152.48, 650.0), (154.00, 646.0),
    (155.52, 642.0), (157.04, 638.0), (158.56, 633.0), (160.08, 628.0),
    (161.60, 622.5), (163.12, 617.0), (164.64, 611.0), (166.16, 604.5),
    (167.68, 598.5), (169.20, 591.5), (173.76, 570.5), (181.36, 531.0),
    (182.88, 522.5), (184.40, 514.0), (185.92, 505.5), (187.44, 497.0),
    (188.96, 488.0), (190.48, 479.5), (192.00, 471.0), (193.52, 462.5),
    (195.04, 454.0), (196.56, 445.5), (198.08, 437.5), (199.60, 429.5),
    (201.12, 421.5), (202.64, 414.5), (204.16, 406.5), (205.68, 399.5),
    (207.20, 392.5), (213.28, 369.0), (214.80, 363.5), (220.88, 346.0),
    (222.40, 343.0), (223.92, 340.0), (225.44, 337.5), (226.96, 335.5),
]
curve_igf = np.array([p[0] for p in CURVE_PTS])
curve_y   = np.array([p[1] for p in CURVE_PTS])

# Densify via cubic spline (scipy se disponível, fallback np.interp)
try:
    from scipy.interpolate import CubicSpline
    cs = CubicSpline(curve_igf, curve_y)
    igf_dense = np.linspace(curve_igf[0], curve_igf[-1], 500)
    y_dense = cs(igf_dense)
except Exception:
    igf_dense = np.linspace(curve_igf[0], curve_igf[-1], 500)
    y_dense = np.interp(igf_dense, curve_igf, curve_y)

xs_curve_px = np.array([igf_to_px(igf) for igf in igf_dense])
ax_bg.plot(xs_curve_px, y_dense, color=INK, linewidth=2.8, zorder=5)

# ============================================================
# X-AXIS — linha + ticks + labels
# Ticks em IGF: 40, 60, 80, 100, 120, 140, 160, 180, 200, 220, 230
# ============================================================
ax_bg.plot([CHART_LEFT_PX-5, CHART_RIGHT_PX+5], [Y_XAXIS, Y_XAXIS],
           color=GRAY1, linewidth=1.0, zorder=3)

X_TICKS_IGF = [40, 60, 80, 100, 120, 140, 160, 180, 200, 220, 230]
for igf in X_TICKS_IGF:
    x_px = igf_to_px(igf)
    # tick mark
    ax_bg.plot([x_px, x_px], [Y_XAXIS, Y_XAXIS+6],
               color=GRAY1, linewidth=0.8, zorder=3)
    # label
    ax_bg.text(x_px, Y_XAXIS+22, str(igf),
               fontsize=9, color=SOFT,
               ha="center", va="center", zorder=10)

# ============================================================
# BRACKETS INFERIORES — 3 (esq red, centro green, dir red)
# Bracket horizontal y=771 (medido)
# Texto bold abaixo
# ============================================================
BRK_Y_TOP = 765
BRK_Y_BAR = 778
BRK_TEXT_Y1 = 808
BRK_TEXT_Y2 = 832
BRK_TEXT_Y3 = 856

def bracket(x_left_igf, x_right_igf, text_lines):
    x0 = igf_to_px(x_left_igf)
    x1 = igf_to_px(x_right_igf)
    cx = (x0 + x1) / 2
    # bracket (linha horizontal + 2 verticais nos cantos)
    ax_bg.plot([x0, x0], [BRK_Y_TOP, BRK_Y_BAR], color=INK, linewidth=1.0)
    ax_bg.plot([x1, x1], [BRK_Y_TOP, BRK_Y_BAR], color=INK, linewidth=1.0)
    ax_bg.plot([x0, x1], [BRK_Y_BAR, BRK_Y_BAR], color=INK, linewidth=1.0)
    # dot no centro
    ax_bg.add_patch(Circle((cx, BRK_Y_BAR), 4, facecolor=INK, edgecolor="none", zorder=5))
    # dotted line down
    ax_bg.plot([cx, cx], [BRK_Y_BAR + 6, BRK_TEXT_Y1 - 12],
               linestyle=(0, (2, 3)), color=INK, linewidth=0.8)
    # texto bold
    y_positions = [BRK_TEXT_Y1, BRK_TEXT_Y2, BRK_TEXT_Y3]
    for i, line in enumerate(text_lines):
        if i < len(y_positions):
            ax_bg.text(cx, y_positions[i], line,
                       fontsize=11, color=INK, weight="bold",
                       ha="center", va="center")

# IGF-1 (ng/mL) label acima do bracket central
center_cx = igf_to_px(140)
ax_bg.text(center_cx, 755, "IGF-1 (ng/mL)",
           fontsize=10, color=INK, weight="bold",
           ha="center", va="center", zorder=10)

# Brackets (em IGF units)
bracket(45, 78, ["Fragilidade,", "sarcopenia,", "declínio funcional"])
bracket(122, 158, ["Zona ótima", "para longevidade"])
bracket(202, 228, ["Aceleração proliferativa,", "risco oncológico"])

# ============================================================
# SOURCE (rodapé)
# ============================================================
ax_bg.plot([60, W_IMG-60], [905, 905], color=GRAY1, linewidth=0.6)

ax_bg.text(60, 925,
           "Ao contrário da maioria dos biomarcadores discutidos neste livro, o IGF-1 não segue a lógica “quanto mais baixo, melhor”. Níveis muito baixos se associam a fragilidade",
           fontsize=8, color=FOOT, va="center", ha="left", style="italic")
ax_bg.text(60, 945,
           "e perda de função; níveis muito altos se associam a proliferação celular e maior risco de alguns cânceres.",
           fontsize=8, color=FOOT, va="center", ha="left", style="italic")
ax_bg.text(60, 975,
           "A faixa ótima (zona verde) varia com idade, mas o princípio é constante. Valores aproximados para adultos de 40–50 anos.",
           fontsize=8, color=FOOT, va="center", ha="left", style="italic")

# ============================================================
# EXPORT
# ============================================================
out_dir = _Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap10_Fig02.pdf"
png_path = out_dir / "_preview_Cap10_Fig02.png"
plt.savefig(pdf_path, facecolor=BG, bbox_inches="tight", pad_inches=0.0)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches="tight", pad_inches=0.0)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
