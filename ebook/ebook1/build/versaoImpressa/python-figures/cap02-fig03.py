"""
Cap02 Fig03 (PT-BR, B&W vetorial) — As Raízes Comuns.

Diagrama em árvore invertida:
  - 4 doenças no topo convergem
  - Banda central "JANELA DE INTERVENÇÃO"
  - 3 raízes abaixo (resistência insulínica, inflamação crônica, disfunção metabólica)
    com setas mostrando reforço mútuo
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, FancyArrowPatch, PathPatch
from matplotlib.path import Path as MPLPath
import numpy as np


def rounded_polyline(waypoints, radius=0.012):
    """Path matplotlib pra polyline com cantos arredondados.

    Os trechos entre waypoints são linhas RETAS; em cada waypoint interno é
    feita uma curva quadrática de raio `radius` que suaviza o canto.
    Estética igual aos diagramas de flowchart bem desenhados.
    """
    verts = [waypoints[0]]
    codes = [MPLPath.MOVETO]

    for i in range(1, len(waypoints) - 1):
        prev = waypoints[i - 1]
        curr = waypoints[i]
        nxt  = waypoints[i + 1]

        # Direção de entrada (prev → curr), normalizada
        din = (curr[0] - prev[0], curr[1] - prev[1])
        lin = (din[0] ** 2 + din[1] ** 2) ** 0.5
        uin = (din[0] / lin, din[1] / lin) if lin else (0, 0)

        # Direção de saída (curr → nxt), normalizada
        dout = (nxt[0] - curr[0], nxt[1] - curr[1])
        lout = (dout[0] ** 2 + dout[1] ** 2) ** 0.5
        uout = (dout[0] / lout, dout[1] / lout) if lout else (0, 0)

        # Ponto antes do canto (recua pelo vetor de entrada)
        p_pre  = (curr[0] - uin[0] * radius,  curr[1] - uin[1] * radius)
        # Ponto depois do canto (avança pelo vetor de saída)
        p_post = (curr[0] + uout[0] * radius, curr[1] + uout[1] * radius)

        # Reta até antes do canto
        verts.append(p_pre)
        codes.append(MPLPath.LINETO)
        # Quadrática: control no canto, end depois do canto
        verts.append(curr)
        verts.append(p_post)
        codes.append(MPLPath.CURVE3)
        codes.append(MPLPath.CURVE3)

    # Reta final até último waypoint
    verts.append(waypoints[-1])
    codes.append(MPLPath.LINETO)

    return MPLPath(verts, codes)

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
BAND     = "#EDEDED"   # banda janela de intervenção

# Figsize reduzido: original tem proporção mais compacta vertical.
fig = plt.figure(figsize=(11.0, 6.2))
fig.patch.set_facecolor(BG)

LEFT = 0.025

# ---------- título ----------
fig.text(LEFT, 0.935, "Figura 3 — As Raízes Comuns",
         fontsize=17, color=INK, weight="bold", va="center")

# ---------- 4 doenças no topo ----------
diseases = ["DOENÇA\nCARDIOVASCULAR", "DOENÇA\nMETABÓLICA",
            "NEURODEGENERAÇÃO", "CÂNCER"]
disease_xs = [0.16, 0.39, 0.62, 0.85]
disease_y  = 0.84

for x, name in zip(disease_xs, diseases):
    fig.text(x, disease_y, name,
             fontsize=10, color=INK, weight="bold",
             ha="center", va="center", linespacing=1.1)
    # Sublinhado curto + fino (proporção ~ texto, elegante como o original)
    fig.lines.append(plt.Line2D(
        [x - 0.038, x + 0.038], [disease_y - 0.040, disease_y - 0.040],
        color=INK, linewidth=0.9, transform=fig.transFigure, zorder=2
    ))

# Curvas de cada doença até o ponto central (topo da banda)
funnel_top_y    = 0.60
funnel_center_x = 0.50

# Polyline: vertical curto sob o label → diagonal pra dentro → vertical curto entrando na banda
for x in disease_xs:
    y_start = disease_y - 0.055
    y_corner1 = y_start - 0.030          # canto onde vira de vertical pra diagonal
    y_corner2 = funnel_top_y + 0.030     # canto onde vira de diagonal pra vertical
    waypoints = [
        (x, y_start),               # sob o label
        (x, y_corner1),             # primeiro canto
        (funnel_center_x, y_corner2),  # segundo canto
        (funnel_center_x, funnel_top_y),  # entrada na banda
    ]
    path = rounded_polyline(waypoints, radius=0.018)
    fig.patches.append(PathPatch(
        path, facecolor="none", edgecolor=INK, linewidth=1.0,
        capstyle="round", joinstyle="round",
        transform=fig.transFigure, zorder=2
    ))

# ---------- banda central "JANELA DE INTERVENÇÃO" (fina e leve) ----------
BAND_X1, BAND_X2 = 0.04, 0.96
BAND_Y1, BAND_Y2 = 0.46, 0.535

fig.patches.append(Rectangle(
    (BAND_X1, BAND_Y1), BAND_X2 - BAND_X1, BAND_Y2 - BAND_Y1,
    facecolor=BAND, edgecolor="none",
    transform=fig.transFigure, zorder=1
))

fig.text(0.5, 0.515, "JANELA DE INTERVENÇÃO:",
         fontsize=12, color=INK, weight="bold", ha="center", va="center")
fig.text(0.5, 0.490, "agir aqui previne todas as quatro.",
         fontsize=10, color=INK_SOFT, ha="center", va="center", style="italic")

# ---------- curvas da banda central até as 3 raízes ----------
root_xs = [0.22, 0.50, 0.78]
root_y_top = 0.34    # topo dos labels das raízes
funnel_bottom_y = 0.46  # base da banda

# 3 polylines saindo da base da banda em direção a cada raiz
for x in root_xs:
    y_start = funnel_bottom_y
    y_corner1 = y_start - 0.030          # primeiro canto (vira pra diagonal)
    y_corner2 = root_y_top + 0.030       # segundo canto (vira pra vertical)
    if abs(x - funnel_center_x) < 0.001:
        # Branch central: linha reta vertical, sem cantos
        waypoints = [
            (funnel_center_x, y_start),
            (funnel_center_x, root_y_top),
        ]
    else:
        waypoints = [
            (funnel_center_x, y_start),
            (funnel_center_x, y_corner1),
            (x, y_corner2),
            (x, root_y_top),
        ]
    path = rounded_polyline(waypoints, radius=0.018)
    fig.patches.append(PathPatch(
        path, facecolor="none", edgecolor=INK, linewidth=1.0,
        capstyle="round", joinstyle="round",
        transform=fig.transFigure, zorder=2
    ))

# ---------- 3 raízes (labels com sublinhado, SEM caixas — como no original) ----------
roots = [
    ("RESISTÊNCIA\nINSULÍNICA", root_xs[0]),
    ("INFLAMAÇÃO\nCRÔNICA",     root_xs[1]),
    ("DISFUNÇÃO\nMETABÓLICA",   root_xs[2]),
]

# "Virtual width" — onde começam/terminam as setas adjacentes (após o texto do label)
ROOT_W = 0.10
root_label_y = root_y_top - 0.040   # centro do label

for name, x in roots:
    fig.text(x, root_label_y, name,
             fontsize=10, color=INK, weight="bold",
             ha="center", va="center", linespacing=1.15, zorder=3)
    # Sublinhado curto + fino (mesma estética dos topos)
    fig.lines.append(plt.Line2D(
        [x - 0.038, x + 0.038], [root_label_y - 0.035, root_label_y - 0.035],
        color=INK, linewidth=0.9, transform=fig.transFigure, zorder=2
    ))

# ---------- setas de reforço mútuo entre as raízes ----------
arrow_y = root_label_y
# RI <-> IC (esquerda <-> meio) — setas finas e pequenas
fig.patches.append(FancyArrowPatch(
    (root_xs[0] + ROOT_W/2 + 0.005, arrow_y),
    (root_xs[1] - ROOT_W/2 - 0.005, arrow_y),
    arrowstyle="<->", color=INK, lw=0.8, mutation_scale=8,
    transform=fig.transFigure, zorder=3
))
# IC <-> DM (meio <-> direita)
fig.patches.append(FancyArrowPatch(
    (root_xs[1] + ROOT_W/2 + 0.005, arrow_y),
    (root_xs[2] - ROOT_W/2 - 0.005, arrow_y),
    arrowstyle="<->", color=INK, lw=0.8, mutation_scale=8,
    transform=fig.transFigure, zorder=3
))
# RI <-> DM (loop externo) — curva discreta, bem baixa, fina
fig.patches.append(FancyArrowPatch(
    (root_xs[0], root_label_y - 0.045),
    (root_xs[2], root_label_y - 0.045),
    arrowstyle="<->", color=INK, lw=0.8, mutation_scale=8,
    connectionstyle="arc3,rad=0.12",
    transform=fig.transFigure, zorder=3
))

# ---------- footer ----------
SEP_Y = 0.135
fig.lines.append(plt.Line2D(
    [LEFT, 1 - LEFT], [SEP_Y, SEP_Y],
    color="#CFCFCF", linewidth=0.5, transform=fig.transFigure
))
fig.text(0.5, 0.090,
         "A prevenção das quatro doenças converge para as mesmas ações —",
         fontsize=10, color=INK_SOFT, ha="center")
fig.text(0.5, 0.050,
         "porque compartilham as mesmas raízes.",
         fontsize=11, color=INK, weight="bold", ha="center")

# ---------- save ----------
out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap02_Fig03.pdf"
png_path = out_dir / "_preview_Cap02_Fig03.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
