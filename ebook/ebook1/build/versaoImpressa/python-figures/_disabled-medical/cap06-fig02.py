"""
Cap06 Fig02 (PT-BR, B&W vetorial) — Mesmo IMC. Riscos Diferentes.

Dois cortes transversais abdominais estilizados:
  - Gordura subcutânea (anel externo)
  - Gordura visceral TOFI (entre os órgãos)
Mesmo IMC = 24, riscos radicalmente diferentes.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, Circle, Ellipse, Wedge
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
SUBCUT   = "#D0D0D0"     # gordura subcutânea — médio claro
VISCERAL = "#7A7A7A"     # gordura visceral — médio escuro
ORGAN    = "#FFFFFF"     # órgãos — branco
SKIN     = "#000000"     # contorno
BAND     = "#EDEDED"

fig = plt.figure(figsize=(11.0, 7.5))
fig.patch.set_facecolor(BG)

LEFT = 0.025

# Título
fig.text(LEFT, 0.945, "Figura 2 — Mesmo IMC. Riscos Diferentes.",
         fontsize=16, color=INK, weight="bold")

# ---------- 2 painéis ----------
panel_titles = [
    "IMC 24 — GORDURA SUBCUTÂNEA",
    "IMC 24 — GORDURA VISCERAL (TOFI)",
]

panel_y = 0.42
panel_h = 0.45

for i, title in enumerate(panel_titles):
    px0 = 0.04 + i * 0.48
    px1 = px0 + 0.40

    fig.text((px0 + px1) / 2, 0.870, title,
             fontsize=11, color=INK, weight="bold", ha="center")
    fig.lines.append(plt.Line2D(
        [px0, px1], [0.853, 0.853],
        color=INK, linewidth=1.0, transform=fig.transFigure
    ))

    ax = fig.add_axes([px0, panel_y, px1 - px0, panel_h - 0.05])
    ax.set_xlim(-1.4, 1.4)
    ax.set_ylim(-1.2, 1.2)
    ax.set_aspect("equal")
    ax.axis("off")

    # Contorno externo da seção abdominal (elipse)
    ax.add_patch(Ellipse((0, 0), 2.4, 2.0, facecolor="white",
                         edgecolor=INK, linewidth=1.3))

    if i == 0:
        # PADRÃO SUBCUTÂNEO: gordura forma anel externo
        # Anel externo (subcutâneo)
        ax.add_patch(Ellipse((0, 0), 2.4, 2.0, facecolor=SUBCUT,
                             edgecolor=INK, linewidth=0))
        # Cavidade interna (sem gordura, órgãos livres)
        ax.add_patch(Ellipse((0, 0), 1.9, 1.5, facecolor="white",
                             edgecolor=INK, linewidth=1.0))
        # Órgãos (vísceras) — círculos brancos
        for cx, cy, w, h in [(-0.55, 0.25, 0.55, 0.45),  # fígado
                              (0.30, 0.25, 0.45, 0.40),  # estômago
                              (-0.10, -0.30, 0.40, 0.35),  # intestino
                              (0.45, -0.25, 0.30, 0.30)]:  # rim
            ax.add_patch(Ellipse((cx, cy), w, h, facecolor=ORGAN,
                                 edgecolor=INK, linewidth=0.8))

        # Labels com setas
        ax.annotate("Gordura\nsubcutânea",
                    xy=(-1.1, 0.0), xytext=(-1.7, 0.50),
                    fontsize=8, color=INK, ha="center",
                    arrowprops=dict(arrowstyle="->", color=INK, lw=0.7))
        ax.annotate("Órgãos\nlivres",
                    xy=(-0.55, 0.25), xytext=(-1.7, -0.40),
                    fontsize=8, color=INK, ha="center",
                    arrowprops=dict(arrowstyle="->", color=INK, lw=0.7))

    else:
        # PADRÃO VISCERAL: gordura entre os órgãos
        # Pequeno anel externo (subcutâneo mínimo)
        ax.add_patch(Ellipse((0, 0), 2.4, 2.0, facecolor=SUBCUT,
                             edgecolor="none", linewidth=0))
        # Cavidade interna — preenchida com gordura visceral
        ax.add_patch(Ellipse((0, 0), 2.1, 1.7, facecolor=VISCERAL,
                             edgecolor=INK, linewidth=1.0))
        # Órgãos pequenos imersos na gordura visceral
        for cx, cy, w, h in [(-0.55, 0.25, 0.50, 0.40),
                              (0.30, 0.25, 0.40, 0.35),
                              (-0.10, -0.30, 0.35, 0.30),
                              (0.45, -0.25, 0.28, 0.28)]:
            ax.add_patch(Ellipse((cx, cy), w, h, facecolor=ORGAN,
                                 edgecolor=INK, linewidth=0.8))

        # Label com seta
        ax.annotate("Gordura\nvisceral",
                    xy=(0.1, 0.05), xytext=(1.6, 0.55),
                    fontsize=8, color=INK, ha="center",
                    arrowprops=dict(arrowstyle="->", color=INK, lw=0.7))

# ---------- tabela de comparação ----------
TBL_Y_TOP = 0.36
TBL_Y_BOTTOM = 0.16
n_rows = 3
ROW_H = (TBL_Y_TOP - TBL_Y_BOTTOM) / n_rows

table_data = [
    ("IMC",              "24",         "24"),
    ("RISCO METABÓLICO", "BAIXO",      "ELEVADO"),
    ("O CHECK-UP VÊ?",   "\"NORMAL\"", "\"NORMAL\""),
]

COL_X = [0.04, 0.34, 0.66, 0.96]   # bordas das colunas: label | esq | dir | fim

# linhas horizontais da tabela
for r in range(n_rows + 1):
    y = TBL_Y_TOP - r * ROW_H
    fig.lines.append(plt.Line2D(
        [COL_X[0], COL_X[3]], [y, y],
        color=INK, linewidth=0.6, transform=fig.transFigure
    ))

# células
for ridx, (label, val_l, val_r) in enumerate(table_data):
    y_top = TBL_Y_TOP - ridx * ROW_H
    y_mid = y_top - ROW_H / 2

    weight_label = "bold"
    weight_val   = "bold" if ridx > 0 else "normal"

    fig.text(COL_X[0] + 0.005, y_mid, label,
             fontsize=10, color=INK, weight=weight_label,
             va="center")
    fig.text((COL_X[1] + COL_X[2]) / 2, y_mid, val_l,
             fontsize=11, color=INK, weight=weight_val,
             ha="center", va="center")
    fig.text((COL_X[2] + COL_X[3]) / 2, y_mid, val_r,
             fontsize=11, color=INK, weight=weight_val,
             ha="center", va="center")

# ---------- footer ----------
fig.text(0.5, 0.110,
         "O IMC é idêntico. O risco é radicalmente diferente.",
         fontsize=10, color=INK_SOFT, ha="center", style="italic")
fig.text(0.5, 0.080,
         "A balança não distingue os dois.",
         fontsize=11, color=INK, weight="bold", ha="center")

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap06_Fig02.pdf"
png_path = out_dir / "_preview_Cap06_Fig02.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
