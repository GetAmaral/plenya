"""
Cap06 Fig01 (PT-BR, B&W vetorial) — Fígado Normal vs. Esteatose: O Que o Ultrassom Revela.

Versão estilizada (esquemática) — o original é foto de ultrassom, impossível de
vetorizar exatamente. Aqui mostramos a IDEIA visual: contraste entre fígado e
rim em ecotextura normal vs hiperecogenicidade difusa.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, Ellipse, FancyArrowPatch, Polygon
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
US_BG    = "#1F1F1F"     # fundo do "ultrassom"
LIVER_OK = "#6E6E6E"     # fígado saudável (médio)
LIVER_BAD= "#C8C8C8"     # fígado esteatótico (muito claro)
KIDNEY   = "#5F5F5F"     # rim (referência)
BAND     = "#EDEDED"

fig = plt.figure(figsize=(11.0, 7.0))
fig.patch.set_facecolor(BG)

LEFT = 0.025

# Título
fig.text(LEFT, 0.945, "Figura 1 — Fígado Normal vs. Esteatose: O Que o Ultrassom Revela",
         fontsize=16, color=INK, weight="bold")

# ---------- 2 painéis lado a lado ----------
panels = [
    ("FÍGADO SAUDÁVEL", "Ecotextura normal", LIVER_OK, KIDNEY, False),
    ("ESTEATOSE GRAU I", "Hiperecogenicidade difusa", LIVER_BAD, KIDNEY, True),
]

PANEL_Y = 0.30
PANEL_H = 0.55

for i, (title, subtitle, liver_color, kidney_color, has_arrow) in enumerate(panels):
    px0 = 0.04 + i * 0.48
    px1 = px0 + 0.40

    # Header
    fig.text((px0 + px1) / 2, 0.870, title,
             fontsize=13, color=INK, weight="bold", ha="center")
    # Linha sob header
    fig.lines.append(plt.Line2D(
        [px0, px1], [0.852, 0.852],
        color=INK, linewidth=1.0, transform=fig.transFigure
    ))

    # Caixa "ultrassom" — fundo escuro
    ax = fig.add_axes([px0, PANEL_Y, px1 - px0, PANEL_H - 0.05])
    ax.set_xlim(0, 1)
    ax.set_ylim(0, 1)
    ax.set_aspect("auto")
    ax.set_facecolor(US_BG)
    ax.set_xticks([])
    ax.set_yticks([])
    # esconde axes lines
    for spine in ax.spines.values():
        spine.set_color(INK)
        spine.set_linewidth(1.0)

    # Fígado: forma orgânica grande na parte superior
    if liver_color == LIVER_OK:
        # ecotextura "normal" — textura sutil de pontos
        liver_pts_x = np.random.default_rng(seed=1).uniform(0.12, 0.85, 200)
        liver_pts_y = np.random.default_rng(seed=2).uniform(0.35, 0.85, 200)
        ax.plot(liver_pts_x, liver_pts_y, ",", color=LIVER_OK, alpha=0.6)
        # Forma do fígado (elipse)
        ax.add_patch(Ellipse((0.50, 0.62), 0.78, 0.42,
                             facecolor=LIVER_OK, alpha=0.55, edgecolor="none"))
    else:
        # esteatose — mais brilhante
        liver_pts_x = np.random.default_rng(seed=3).uniform(0.12, 0.85, 200)
        liver_pts_y = np.random.default_rng(seed=4).uniform(0.35, 0.85, 200)
        ax.plot(liver_pts_x, liver_pts_y, ",", color=LIVER_BAD, alpha=0.6)
        ax.add_patch(Ellipse((0.50, 0.62), 0.78, 0.42,
                             facecolor=LIVER_BAD, alpha=0.70, edgecolor="none"))

    # Rim: forma feijão na parte inferior direita
    ax.add_patch(Ellipse((0.65, 0.22), 0.32, 0.18,
                         facecolor=kidney_color, alpha=0.85, edgecolor="none"))
    # Centro mais claro do rim
    ax.add_patch(Ellipse((0.65, 0.22), 0.20, 0.10,
                         facecolor="#9A9A9A", alpha=0.6, edgecolor="none"))

    # Labels brancas dentro
    ax.text(0.25, 0.62, "FÍGADO",
            fontsize=10, color="white", weight="bold", ha="center", va="center")
    ax.text(0.65, 0.10, "RIM DIREITO",
            fontsize=8, color="white", weight="bold", ha="center", va="center")

    # Subtítulo abaixo
    fig.text((px0 + px1) / 2, PANEL_Y - 0.030, subtitle,
             fontsize=10, color=INK, weight="bold",
             ha="center", va="top", style="italic")

    # Anotação (apenas no painel direito)
    if has_arrow:
        # Linha + texto explicativo
        annot_x = px1 + 0.005
        annot_y = PANEL_Y + 0.32
        fig.text(annot_x + 0.005, annot_y, "Fígado mais\nbrilhante\nque o rim\n= acúmulo\nde gordura",
                 fontsize=8, color=INK, weight="bold", ha="left", va="center",
                 linespacing=1.3)
        fig.patches.append(FancyArrowPatch(
            (annot_x, annot_y), (px1 - 0.10, PANEL_Y + 0.30),
            arrowstyle="<-", color=INK, lw=1.0, mutation_scale=12,
            transform=fig.transFigure
        ))

# ---------- caixa final ----------
BOX_X1, BOX_X2 = 0.04, 0.96
BOX_Y1, BOX_Y2 = 0.08, 0.18
fig.patches.append(Rectangle(
    (BOX_X1, BOX_Y1), BOX_X2 - BOX_X1, BOX_Y2 - BOX_Y1,
    facecolor=BAND, edgecolor=INK, linewidth=0.5,
    transform=fig.transFigure, zorder=1
))
fig.text(0.5, (BOX_Y1 + BOX_Y2) / 2 + 0.015,
         "A esteatose é visível antes de qualquer sintoma —",
         fontsize=10, color=INK, ha="center")
fig.text(0.5, (BOX_Y1 + BOX_Y2) / 2 - 0.012,
         "e antes de qualquer alteração nos exames de sangue convencionais.",
         fontsize=10, color=INK, weight="bold", ha="center")

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap06_Fig01.pdf"
png_path = out_dir / "_preview_Cap06_Fig01.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
