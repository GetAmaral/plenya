"""
Cap03 Fig01 (PT-BR, B&W vetorial) — Os 5 Marcadores do Envelhecimento.

Posições EXATAS dos 5 markers detectadas do PNG original via análise pixel-a-pixel:
  Inflammaging   (red):    px(216,517)  → fig(0.141, 0.495)
  Disfunção mit. (orange): px(464,418)  → fig(0.302, 0.592)
  Senescência    (green):  px(760,392)  → fig(0.495, 0.617)
  Encurtamento   (blue):   px(1058,418) → fig(0.689, 0.592)
  Instabilidade  (purple): px(1316,528) → fig(0.857, 0.485)
Elipse: cx=0.499  cy=0.490  a=0.358  b=0.127  (derivado dos markers)

Figsize 11×7.33 preserva aspect ratio do original (1536×1024 = 1.5:1).
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, Ellipse
import numpy as np

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False
rcParams["pdf.fonttype"] = 42
rcParams["ps.fonttype"] = 42

BG       = "#FFFFFF"
INK      = "#000000"
INK_SOFT = "#3A3A3A"
DOT      = "#000000"
DASH     = "#A0A0A0"
BAND     = "#EDEDED"

# Aspect ratio idêntico ao original (1.5:1)
# Aspecto 1.500, igual ao da arte original (1536×1024).
# Era 1.501.
fig = plt.figure(figsize=(11.0, 7.333))
fig.patch.set_facecolor(BG)

# Coordenadas EXATAS dos markers em figure-fraction (do original)
markers = [
    (0.141, 0.495, "Inflammaging",           "O fogo silencioso"),
    (0.302, 0.592, "Disfunção\nmitocondrial", "Usinas falhando"),
    (0.495, 0.617, "Senescência\ncelular",    "Células zumbis"),
    (0.689, 0.592, "Encurtamento\ntelomérico","Relógios nas pontas"),
    (0.857, 0.485, "Instabilidade\nepigenética","Genes ligando\ne desligando"),
]

# Title — posição detectada Y≈85px → fig y=0.917; X=47px → fig x=0.031
fig.text(0.031, 0.917, "Figura 1 — Os 5 Marcadores do Envelhecimento",
         fontsize=20, color=INK, weight="bold", va="center")

# ---------- Elipse ----------
# Parâmetros EXATOS detectados do original via pixel sweep:
#   X span: 175-1349 px → a=587 px = 0.382 fig
#   Y span: 392-830 px → b=219 px = 0.214 fig
#   Center: (762, 611) → fig (0.496, 0.403)
# Os markers ficam acima do equador da elipse (não nas vértices).
ellipse_cx, ellipse_cy = 0.496, 0.403
ellipse_a, ellipse_b   = 0.382, 0.214

theta = np.linspace(0, 2 * np.pi, 240)
ex = ellipse_cx + ellipse_a * np.cos(theta)
ey = ellipse_cy + ellipse_b * np.sin(theta)
fig.add_artist(plt.Line2D(ex, ey, color="#888888", linewidth=1.0, transform=fig.transFigure, zorder=1))

# ---------- Dashed lines (curvas — bezier com sag descendente, como no original) ----------
# Cada par de markers conectado por curva tracejada que afunda em direção ao
# interior da elipse (não é linha reta como flowchart).
for i in range(len(markers)):
    for j in range(i + 1, len(markers)):
        x1, y1 = markers[i][0], markers[i][1]
        x2, y2 = markers[j][0], markers[j][1]
        # Gera bezier quadrático em N pontos com control point afundado
        n_pts = 80
        t = np.linspace(0, 1, n_pts)
        mid_x = (x1 + x2) / 2
        mid_y = (y1 + y2) / 2
        # Sag: control point abaixo do midpoint (em direção ao centro da elipse)
        # Sag proporcional à distância horizontal entre os markers (linhas longas afundam mais)
        dx_h = abs(x2 - x1)
        sag = 0.08 + 0.20 * (dx_h / 0.72)   # min 0.08, max 0.28
        ctrl_x = mid_x
        ctrl_y = mid_y - sag
        # Bezier quadrático: P(t) = (1-t)²·P0 + 2(1-t)t·C + t²·P1
        cx_arr = (1 - t) ** 2 * x1 + 2 * (1 - t) * t * ctrl_x + t ** 2 * x2
        cy_arr = (1 - t) ** 2 * y1 + 2 * (1 - t) * t * ctrl_y + t ** 2 * y2
        fig.add_artist(plt.Line2D(
            cx_arr, cy_arr,
            color=DASH, linewidth=0.7, linestyle=(0, (3, 3)),
            transform=fig.transFigure, zorder=2
        ))

# ---------- Markers (círculos pretos preenchidos) ----------
# Diâmetro original: 21px / 1536 = 1.37% da largura da figura
# Em pontos: 1.37% × 11" × 72pt/" = 10.8pt
MARKER_PT = 11
for x, y, title, sub in markers:
    fig.add_artist(plt.Line2D(
        [x], [y],
        marker='o', markersize=MARKER_PT,
        markerfacecolor=DOT, markeredgecolor=DOT,
        transform=fig.transFigure, zorder=4
    ))

# ---------- Labels (centrados acima dos markers) ----------
# Estrutura igual original: TÍTULO bold + SUBLINHADO curto + SUBTÍTULO regular + MARKER
for x, y, title, sub in markers:
    sub_y     = y + 0.060            # subtítulo regular logo acima do marker
    underline_y = sub_y + 0.045      # sublinhado decorativo entre title e subtitle
    title_y   = underline_y + 0.020  # título acima do sublinhado

    # Título (bold, preto)
    fig.text(x, title_y, title,
             fontsize=12, color=INK, weight="bold",
             ha="center", va="bottom", linespacing=1.05)
    # Sublinhado decorativo curto (no original tem cor do marker; em B&W usamos preto)
    fig.add_artist(plt.Line2D(
        [x - 0.045, x + 0.045], [underline_y, underline_y],
        color=INK, linewidth=1.4, transform=fig.transFigure, zorder=3
    ))
    # Subtítulo (regular, não italic)
    fig.text(x, sub_y, sub,
             fontsize=10, color=INK_SOFT,
             ha="center", va="bottom", linespacing=1.1)

# ---------- Texto central dentro da elipse ----------
# Detectado no original em y=750 px → fig y = 0.268
fig.text(ellipse_cx, 0.268,
         "Processos interligados.\nCada um alimenta os outros.",
         fontsize=11, color=INK_SOFT,
         ha="center", va="center", linespacing=1.3,
         transform=fig.transFigure)

# ---------- Banda inferior ----------
BOX_X1, BOX_X2 = 0.04, 0.96
BOX_Y1, BOX_Y2 = 0.06, 0.14
fig.patches.append(Rectangle(
    (BOX_X1, BOX_Y1), BOX_X2 - BOX_X1, BOX_Y2 - BOX_Y1,
    facecolor=BAND, edgecolor="none", transform=fig.transFigure, zorder=1
))
fig.text(0.5, (BOX_Y1 + BOX_Y2) / 2,
         "Todos são modificáveis por estilo de vida.",
         fontsize=14, color=INK, weight="bold",
         ha="center", va="center")

# ---------- Save ----------
out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap03_Fig01.pdf"
png_path = out_dir / "_preview_Cap03_Fig01.png"

from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG)
plt.savefig(png_path, dpi=170, facecolor=BG)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
