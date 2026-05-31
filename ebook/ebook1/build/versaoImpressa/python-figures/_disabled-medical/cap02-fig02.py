"""
Cap02 Fig02 (PT-BR, B&W vetorial) — Da Estria Gordurosa ao Infarto.

4 cortes transversais de artéria mostrando progressão da aterosclerose:
  1. Artéria saudável
  2. Estria gordurosa
  3. Placa estável (70% obstrução, capa fibrosa grossa)
  4. Placa vulnerável + ruptura (40% obstrução, capa fibrosa fina)

Saída: PDF vetorial em build/versaoImpressa/figuras-bw/Cap02_Fig02.pdf
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Circle, Wedge, FancyArrowPatch
import numpy as np

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False
rcParams["pdf.fonttype"] = 42
rcParams["ps.fonttype"] = 42

# ---------- paleta B&W ----------
BG       = "#FFFFFF"
INK      = "#000000"
INK_SOFT = "#3A3A3A"
TICK     = "#555555"
FOOT     = "#666666"
WALL     = "#E0E0E0"   # parede arterial — cinza claro
PLAQUE   = "#B5B5B5"   # placa — cinza médio
THROMBUS = "#1F1F1F"   # trombo — quase preto

# ---------- figura ----------
fig = plt.figure(figsize=(11.0, 6.0))
fig.patch.set_facecolor(BG)

LEFT_MARGIN = 0.025

# Título
fig.text(LEFT_MARGIN, 0.945, "Figura 2 — Da Estria Gordurosa ao Infarto",
         fontsize=17, color=INK, weight="bold", va="center")
fig.text(LEFT_MARGIN, 0.895,
         "A aterosclerose é um processo inflamatório na parede arterial — não é apenas \"colesterol entupindo um cano\".",
         fontsize=10, color=INK_SOFT)

# ---------- 4 colunas com axes (1 por estágio) ----------
columns = [
    ("1. Artéria saudável",      "Infância /\nadolescência",           None, None),
    ("2. Estria gordurosa",      "Acúmulo de lipídios e\ncélulas espumosas na íntima.\nLúmen ainda preservado.", None, None),
    ("3. Placa estável",         "10–20 anos",
     "~70% obstrução,\ncapa fibrosa grossa",  "Pode nunca romper"),
    ("4. Placa vulnerável + ruptura", None,
     "~40% obstrução,\ncapa fibrosa fina, RUPTURA", "Pode matar"),
]

# Posições dos axes
AXES_Y = 0.36
AXES_W = 0.20
AXES_H = 0.36
START_X = 0.04
GAP_X = 0.04

for i, (title, subtitle, descr, outcome) in enumerate(columns):
    ax_x = START_X + i * (AXES_W + GAP_X)
    cx_center = ax_x + AXES_W / 2

    # Título do estágio (acima)
    fig.text(cx_center, 0.835, title,
             fontsize=11, color=INK, weight="bold", ha="center")

    # Eixo de desenho — bounds expandidos pra anotações caberem
    ax = fig.add_axes([ax_x, AXES_Y, AXES_W, AXES_H])
    ax.set_xlim(-1.6, 1.6)
    ax.set_ylim(-1.3, 1.5)
    ax.set_aspect("equal")
    ax.axis("off")

    # --- desenho específico de cada estágio ---
    # Todos têm: anel externo (parede) e área interna (lúmen)
    # Estágios 2-4 adicionam placa; estágio 4 adiciona trombo

    if i == 0:
        # Artéria saudável: 2 anéis concêntricos limpos
        ax.add_patch(Circle((0, 0), 1.0, facecolor=WALL, edgecolor=INK, linewidth=1.2))
        ax.add_patch(Circle((0, 0), 0.72, facecolor="white", edgecolor=INK, linewidth=0.8))
        # Etiqueta interna
        ax.text(0, 0, "Lúmen", fontsize=8.5, color=INK_SOFT,
                ha="center", va="center", style="italic")

    elif i == 1:
        # Estria gordurosa: anel + pequena protrusão (estria) na parede inferior
        ax.add_patch(Circle((0, 0), 1.0, facecolor=WALL, edgecolor=INK, linewidth=1.2))
        ax.add_patch(Circle((0, 0), 0.72, facecolor="white", edgecolor=INK, linewidth=0.8))
        # Pequena placa (wedge) na parte inferior do lúmen
        wedge = Wedge((0, -0.3), 0.42, 35, 145, facecolor=PLAQUE,
                      edgecolor=INK, linewidth=0.6, alpha=0.85)
        ax.add_patch(wedge)
        # Etiqueta interna
        ax.text(0, 0.32, "Lúmen", fontsize=8.5, color=INK_SOFT,
                ha="center", va="center", style="italic")

    elif i == 2:
        # Placa estável: placa grande com capa fibrosa grossa
        ax.add_patch(Circle((0, 0), 1.0, facecolor=WALL, edgecolor=INK, linewidth=1.2))
        ax.add_patch(Circle((0, 0), 0.72, facecolor="white", edgecolor=INK, linewidth=0.8))
        # Placa: grande wedge inferior (núcleo lipídico)
        wedge_outer = Wedge((0, -0.5), 0.85, 25, 155, facecolor=PLAQUE,
                            edgecolor=INK, linewidth=1.0, alpha=0.85)
        ax.add_patch(wedge_outer)
        # Capa fibrosa grossa (linha mais espessa na borda superior da placa)
        theta = np.linspace(25, 155, 60)
        rad = 0.85
        cx_arc = rad * np.cos(np.radians(theta))
        cy_arc = -0.5 + rad * np.sin(np.radians(theta))
        ax.plot(cx_arc, cy_arc, color=INK, linewidth=2.5)

        # Etiqueta interna
        ax.text(0, 0.42, "Lúmen\n(estreito)", fontsize=7.5, color=INK_SOFT,
                ha="center", va="center", style="italic")

        # Anotação "Capa fibrosa grossa" com seta apontando pra arco
        ax.annotate("Capa fibrosa\ngrossa",
                    xy=(0.0, 0.34), xytext=(1.10, 1.30),
                    fontsize=7.5, color=INK, ha="center",
                    arrowprops=dict(arrowstyle="->", color=INK, lw=0.7))
        # "Núcleo lipídico" dentro da placa (texto preto sobre cinza claro)
        ax.text(0, -0.55, "Núcleo\nlipídico", fontsize=7.5, color=INK,
                ha="center", va="center", weight="bold")
        ax.annotate("Inflamação\n(moderada)",
                    xy=(-0.5, -0.50), xytext=(-1.50, -0.90),
                    fontsize=7.0, color=INK_SOFT, ha="center", style="italic",
                    arrowprops=dict(arrowstyle="->", color=INK_SOFT, lw=0.6))

    else:
        # Placa vulnerável + ruptura: placa menor com capa fina, trombo
        ax.add_patch(Circle((0, 0), 1.0, facecolor=WALL, edgecolor=INK, linewidth=1.2))
        ax.add_patch(Circle((0, 0), 0.72, facecolor="white", edgecolor=INK, linewidth=0.8))
        # Placa menor (mas com núcleo lipídico grande proporcionalmente)
        wedge_outer = Wedge((0, -0.45), 0.72, 30, 150, facecolor=PLAQUE,
                            edgecolor=INK, linewidth=1.0, alpha=0.85)
        ax.add_patch(wedge_outer)
        # Capa fibrosa fina (linha tracejada/quebrada para indicar ruptura)
        theta = np.linspace(30, 80, 30)
        rad = 0.72
        cx_arc = rad * np.cos(np.radians(theta))
        cy_arc = -0.45 + rad * np.sin(np.radians(theta))
        ax.plot(cx_arc, cy_arc, color=INK, linewidth=1.2)
        # Ruptura (gap) entre 80 e 100 graus
        theta2 = np.linspace(100, 150, 30)
        cx_arc2 = rad * np.cos(np.radians(theta2))
        cy_arc2 = -0.45 + rad * np.sin(np.radians(theta2))
        ax.plot(cx_arc2, cy_arc2, color=INK, linewidth=1.2)
        # Marker da ruptura: linha em "X" pequena no gap
        ax.plot([0.0, 0.18], [0.25, 0.40], color=INK, linewidth=1.0)
        ax.plot([0.0, -0.18], [0.25, 0.40], color=INK, linewidth=1.0)

        # Trombo (coágulo) — círculo escuro acima da ruptura
        ax.add_patch(Circle((0.05, 0.20), 0.18, facecolor=THROMBUS,
                            edgecolor=INK, linewidth=0.6, alpha=0.9))
        # Etiquetas
        ax.text(0, -0.55, "Núcleo\nlipídico\ngrande", fontsize=7.0, color=INK,
                ha="center", va="center", weight="bold")
        ax.annotate("Ruptura\nda capa fina",
                    xy=(0.0, 0.42), xytext=(-1.20, 1.20),
                    fontsize=7.5, color=INK, ha="center",
                    arrowprops=dict(arrowstyle="->", color=INK, lw=0.7))
        ax.annotate("Trombo\n(coágulo)",
                    xy=(0.20, 0.20), xytext=(1.25, 0.50),
                    fontsize=7.5, color=INK, ha="center",
                    arrowprops=dict(arrowstyle="->", color=INK, lw=0.7))

    # --- Subtitle (debaixo do círculo, na figura) ---
    if subtitle:
        fig.text(cx_center, AXES_Y - 0.020, subtitle,
                 fontsize=8.5, color=INK_SOFT, ha="center", va="top",
                 linespacing=1.15, style="italic")

    if descr:
        fig.text(cx_center, AXES_Y - 0.075, descr,
                 fontsize=8.5, color=INK, weight="bold", ha="center", va="top",
                 linespacing=1.15)

    if outcome:
        # Caixinha embaixo com outcome (Pode nunca romper / Pode matar)
        fig.text(cx_center, AXES_Y - 0.125, outcome,
                 fontsize=9, color=INK, weight="bold", style="italic",
                 ha="center", va="top")

# ---------- footer (frase central + fonte) ----------
fig.text(0.5, 0.090,
         "O check-up convencional procura obstrução.",
         fontsize=10, color=INK, ha="center")
fig.text(0.5, 0.062,
         "A doença real é inflamação.",
         fontsize=11, color=INK, weight="bold", style="italic", ha="center")

# ---------- save ----------
out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap02_Fig02.pdf"
png_path = out_dir / "_preview_Cap02_Fig02.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
