"""
Cap03 Fig01 (PT-BR, B&W vetorial) — Os 5 Marcadores do Envelhecimento.

5 pontos numa elipse, conectados por linhas tracejadas (interligados).
Cada um com nome + descrição curta. Caixa inferior: "Todos são modificáveis".
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

fig = plt.figure(figsize=(11.0, 7.0))
fig.patch.set_facecolor(BG)

LEFT = 0.025

# Título
fig.text(LEFT, 0.945, "Figura 1 — Os 5 Marcadores do Envelhecimento",
         fontsize=17, color=INK, weight="bold", va="center")

# ---------- elipse com 5 pontos ----------
# Adicionamos eixo coordenado pra plotar mais facilmente
ax = fig.add_axes([0.05, 0.20, 0.90, 0.65])
ax.set_xlim(0, 1)
ax.set_ylim(0, 1)
ax.axis("off")

# 5 pontos numa curva côncava (elipse, parte superior)
# Posições x igualmente espaçadas, y seguindo um arco
xs_pts = np.linspace(0.06, 0.94, 5)
ys_pts = 0.50 + 0.18 * np.sin(np.linspace(0.1, np.pi - 0.1, 5))  # arco curvado

# Elipse de fundo (cinza suave) — desenhada como linha contínua
theta = np.linspace(0, 2*np.pi, 200)
ellipse_cx, ellipse_cy = 0.50, 0.40
ellipse_a, ellipse_b   = 0.46, 0.30
ex = ellipse_cx + ellipse_a * np.cos(theta)
ey = ellipse_cy + ellipse_b * np.sin(theta)
ax.plot(ex, ey, color="#888888", linewidth=1.0, zorder=1)

# Reposicionar pontos na elipse superior
# (mais elegante: pontos exatamente sobre a elipse, distribuídos em arco superior)
arc_angles = np.linspace(np.pi - 0.3, 0.3, 5)
xs_pts = ellipse_cx + ellipse_a * np.cos(arc_angles)
ys_pts = ellipse_cy + ellipse_b * np.sin(arc_angles)

# Linhas tracejadas conectando todos os pontos entre si
for i in range(len(xs_pts)):
    for j in range(i + 1, len(xs_pts)):
        ax.plot([xs_pts[i], xs_pts[j]], [ys_pts[i], ys_pts[j]],
                color=DASH, linewidth=0.7, linestyle=(0, (3, 3)), zorder=2)

# Pontos
for x, y in zip(xs_pts, ys_pts):
    ax.plot([x], [y], "o", color=DOT, markersize=10, zorder=4)

# Rótulos dos marcadores (acima do ponto)
labels = [
    ("Inflammaging",          "O fogo silencioso"),
    ("Disfunção\nmitocondrial",  "Usinas falhando"),
    ("Senescência\ncelular",     "Células zumbis"),
    ("Encurtamento\ntelomérico", "Relógios nas pontas"),
    ("Instabilidade\nepigenética", "Genes ligando\ne desligando"),
]

# Posições verticais dos textos: alguns acima, outros abaixo dependendo da curva
for i, (x, y, (title, sub)) in enumerate(zip(xs_pts, ys_pts, labels)):
    # Título acima do ponto
    ty = y + 0.10
    ax.text(x, ty, title,
            fontsize=10.5, color=INK, weight="bold",
            ha="center", va="bottom", linespacing=1.1)
    # Subtítulo abaixo do título
    ax.text(x, ty - 0.06, sub,
            fontsize=8.5, color=INK_SOFT,
            ha="center", va="bottom", linespacing=1.1, style="italic")

# Texto central dentro da elipse
ax.text(ellipse_cx, ellipse_cy - 0.05,
        "Processos interligados.\nCada um alimenta os outros.",
        fontsize=10, color=INK_SOFT,
        ha="center", va="center", linespacing=1.3)

# ---------- caixa inferior ----------
BOX_X1, BOX_X2 = 0.05, 0.95
BOX_Y1, BOX_Y2 = 0.06, 0.14
fig.patches.append(Rectangle(
    (BOX_X1, BOX_Y1), BOX_X2 - BOX_X1, BOX_Y2 - BOX_Y1,
    facecolor=BAND, edgecolor="none", transform=fig.transFigure, zorder=1
))
fig.text(0.5, (BOX_Y1 + BOX_Y2) / 2,
         "Todos são modificáveis por estilo de vida.",
         fontsize=13, color=INK, weight="bold",
         ha="center", va="center")

# ---------- save ----------
out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap03_Fig01.pdf"
png_path = out_dir / "_preview_Cap03_Fig01.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
