"""
Cap07 Fig01 (PT-BR, B&W vetorial) — Ser sedentário extremo aumenta em ~5x o risco de morte.
Reconstrução pixel-a-pixel do original (1536x1024, aspect 1.5).
Fonte: Mandsager et al., JAMA Network Open, 2018.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, Polygon

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False
rcParams["pdf.fonttype"] = 42
rcParams["ps.fonttype"] = 42

BG, INK, SOFT, FOOT = "#FFFFFF", "#000000", "#3A3A3A", "#666666"
BAR_GRAY = "#9A9A9A"
BAR_RED  = "#000000"   # vermelho original → preto sólido B&W (destaque)

_FIG_W, _FIG_H = 11.0, 7.333
fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

# ============= FIGURA 1 TAG (texto cinza, NÃO box) =============
fig.text(0.024, 0.957, "FIGURA 1",
         fontsize=10, color=SOFT, weight="bold", va="center", ha="left")

# ============= TÍTULO (cy=0.906) =============
fig.text(0.024, 0.906, "Ser sedentário extremo aumenta em ~5x o risco de morte",
         fontsize=20, color=INK, weight="bold", va="center", ha="left")

# ============= SUBTÍTULO (cy=0.839) =============
fig.text(0.024, 0.839, "Comparado a outros fatores de risco conhecidos",
         fontsize=12, color=SOFT, va="center", ha="left")

# ============= BARRAS (posições EXATAS detectadas) =============
bars = [
    # (cy_norm, height_norm, x_right_norm, value_label, is_red)
    (0.745, 0.061, 0.888, "5,0x", True),
    (0.645, 0.055, 0.533, "2,0x", False),
    (0.551, 0.051, 0.451, "1,4x", False),
    (0.460, 0.051, 0.450, "1,4x", False),
    (0.371, 0.051, 0.431, "1,3x", False),
]

BAR_X_LEFT = 0.301
for cy, h, x_right, value_label, is_red in bars:
    y_bot = cy - h/2
    color = BAR_RED if is_red else BAR_GRAY
    fig.patches.append(Rectangle(
        (BAR_X_LEFT, y_bot), x_right - BAR_X_LEFT, h,
        facecolor=color, edgecolor="none",
        transform=fig.transFigure, zorder=2
    ))
    label_color = INK if is_red else BAR_GRAY
    fig.text(x_right + 0.008, cy, value_label,
             fontsize=18 if is_red else 16,
             color=label_color, weight="bold",
             va="center", ha="left", zorder=3)

# ============= CATEGORIA LABELS (esquerda das barras) =============
fig.text(0.024, 0.760, "Aptidão cardiorrespiratória",
         fontsize=12, color=INK, weight="bold", va="center", ha="left")
fig.text(0.024, 0.733, "baixa (sedentarismo extremo)",
         fontsize=12, color=INK, weight="bold", va="center", ha="left")

fig.text(0.024, 0.662, "Doença renal", fontsize=12, color=INK, weight="bold",
         va="center", ha="left")
fig.text(0.024, 0.627, "em estágio terminal", fontsize=12, color=INK,
         weight="bold", va="center", ha="left")

fig.text(0.024, 0.550, "Tabagismo", fontsize=12, color=INK, weight="bold",
         va="center", ha="left")

fig.text(0.024, 0.463, "Diabetes", fontsize=12, color=INK, weight="bold",
         va="center", ha="left")

fig.text(0.024, 0.381, "Doença coronariana", fontsize=12, color=INK,
         weight="bold", va="center", ha="left")
fig.text(0.024, 0.344, "(do coração)", fontsize=12, color=INK,
         weight="bold", va="center", ha="left")

# ============= LINHA DASHED VERTICAL em x=1.0 =============
DASH_X = 0.301
fig.lines.append(plt.Line2D(
    [DASH_X, DASH_X], [0.305, 0.793],
    color=INK, linewidth=1.0, linestyle=(0, (5, 4)),
    transform=fig.transFigure, zorder=1
))

# ============= EIXO X (linha + ticks) =============
AXIS_Y = 0.305
fig.lines.append(plt.Line2D(
    [0.290, 0.930], [AXIS_Y, AXIS_Y],
    color=INK, linewidth=1.2, transform=fig.transFigure
))

# ▲ marker abaixo de x=1.0
fig.patches.append(Polygon(
    [(DASH_X, AXIS_Y - 0.005),
     (DASH_X - 0.006, AXIS_Y - 0.020),
     (DASH_X + 0.006, AXIS_Y - 0.020)],
    closed=True, facecolor=INK, edgecolor=INK,
    transform=fig.transFigure, zorder=3
))

# Labels do ponto referência
fig.text(DASH_X, 0.225, "1,0x",
         fontsize=14, color=INK, weight="bold", va="center", ha="center")
fig.text(DASH_X, 0.199, "Risco de referência",
         fontsize=10, color=INK, weight="bold", va="center", ha="center")
fig.text(DASH_X, 0.176, "(pessoa em excelente forma",
         fontsize=9.5, color=INK, va="center", ha="center")
fig.text(DASH_X, 0.156, "e sem o fator de risco)",
         fontsize=9.5, color=INK, va="center", ha="center")

# Ticks 2, 3, 4, 5
UNIT = 0.147   # (0.888 - 0.301) / 4
for tk in (2, 3, 4, 5):
    tx = DASH_X + (tk - 1) * UNIT
    fig.lines.append(plt.Line2D(
        [tx, tx], [AXIS_Y, AXIS_Y - 0.010],
        color=INK, linewidth=1.0, transform=fig.transFigure
    ))
    fig.text(tx, AXIS_Y - 0.030, str(tk),
             fontsize=13, color=INK, va="center", ha="center")

# "Risco de morte (vezes maior)" axis title
fig.text(0.595, 0.240, "Risco de morte (vezes maior)",
         fontsize=13, color=INK, weight="bold", va="center", ha="center")

# ============= CALLOUT (cy=0.129) =============
fig.text(0.500, 0.129, "Mais que fumar, diabetes ou doença do coração.",
         fontsize=14, color=INK, weight="bold", style="italic",
         va="center", ha="center")

# ============= SEPARADOR + FOOTER =============
fig.lines.append(plt.Line2D(
    [0.024, 0.976], [0.097, 0.097],
    color="#BBBBBB", linewidth=0.5, transform=fig.transFigure
))
fig.text(0.024, 0.062,
         "Fonte: Mandsager et al., JAMA Network Open, 2018. Estudo com 122.007 adultos acompanhados por mediana de 8,4 anos.",
         fontsize=9.5, color=FOOT, va="center", ha="left")
fig.text(0.024, 0.034,
         "Aptidão baixa definida como abaixo do percentil 25 para idade e sexo. Referência: grupo elite (acima do percentil 97,7).",
         fontsize=9.5, color=FOOT, va="center", ha="left")

# ============= EXPORT =============
out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap07_Fig01.pdf"
png_path = out_dir / "_preview_Cap07_Fig01.png"
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
