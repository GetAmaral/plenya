"""
Cap07 Fig03 (PT-BR, B&W vetorial) — Os 4 pilares do exercício para longevidade (na prática).
Donut chart com 4 segmentos + labels com leader dots + ícone de pessoa no centro.
Reconstrução pixel-a-pixel do original (1254x1254, square).
"""
from pathlib import Path as _Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Circle, Wedge, PathPatch
from matplotlib.path import Path
import numpy as np

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False
rcParams["pdf.fonttype"] = 42
rcParams["ps.fonttype"] = 42

BG       = "#FFFFFF"
INK      = "#000000"
SOFT     = "#3A3A3A"
FOOT     = "#666666"

# Shades B&W para os 4 segmentos
S_ZONA2 = "#3A3A3A"    # verde escuro → preto sólido com texto branco
S_FORCA = "#000000"    # azul navy → preto
S_HIIT  = "#8A8A8A"    # laranja → cinza médio
S_MOB   = "#BFBFBF"    # cinza → cinza claro

# Square figure (aspect 1.0)
_FIG = 10.0
fig = plt.figure(figsize=(_FIG, _FIG))
fig.patch.set_facecolor(BG)

# =================== TÍTULO ===================
fig.text(0.030, 0.970, "FIGURA 3",
         fontsize=11, color=SOFT, weight="bold", va="center")
fig.text(0.030, 0.935, "Os 4 pilares do exercício para longevidade (na prática)",
         fontsize=22, color=INK, weight="bold", va="center")
fig.text(0.030, 0.895,
         "Um sistema integrado — cada pilar tem uma função específica.",
         fontsize=12, color=SOFT, va="center")

# =================== DONUT (centered, large) ===================
# Plot ax centrado, square
ax = fig.add_axes([0.20, 0.18, 0.60, 0.60])
ax.set_xlim(-1, 1)
ax.set_ylim(-1, 1)
ax.set_aspect("equal")
ax.axis("off")

R_OUT = 0.85   # raio externo
R_IN  = 0.27   # raio interno (~32% R_OUT, conforme original)

# Segments (CCW angles). Larguras VISUAIS pra acomodar texto, não estritamente proporcionais.
# Total 360°: Zona2(175°) + Forca(90°) + HIIT(50°) + Mob(45°)
segments = [
    # Ângulos medidos pixel-a-pixel do original
    # Zona2 95°-239° (144°), Mob 239°-312° (73°), HIIT 312°-12° wrap (60°), Forca 12°-95° (83°)
    ("Zona 2", 95,  239, S_ZONA2, "50–60%", "1. ZONA 2",
     ["Base metabólica", "Mitocôndrias"]),
    ("Mob",   239,  312, S_MOB,   "5–10%",  "4. MOBILIDADE",
     ["Mobilidade", "Prevenção", "de quedas"]),
    ("HIIT",  312,  372, S_HIIT,  "10–15%", "3. HIIT",
     ["VO₂ máx", "Capacidade", "máxima"]),
    ("Forca",  12,   95, S_FORCA, "25–30%", "2. FORÇA",
     ["Massa muscular", "Sensibilidade", "insulínica"]),
]

for name, t1, t2, color, pct, num, desc_lines in segments:
    w = Wedge((0, 0), R_OUT, t1, t2, width=R_OUT - R_IN,
              facecolor=color, edgecolor="white", linewidth=2)
    ax.add_patch(w)

    mid_angle = (t1 + t2) / 2
    mid_rad = (R_OUT + R_IN) / 2
    tx = mid_rad * np.cos(np.radians(mid_angle))
    ty = mid_rad * np.sin(np.radians(mid_angle))

    text_color = "white" if color in (S_ZONA2, S_FORCA) else INK

    # Stack vertical: pct on top, pillar name middle, description below
    ax.text(tx, ty + 0.13, pct,
            fontsize=14, color=text_color, weight="bold",
            ha="center", va="center")
    ax.text(tx, ty + 0.05, num,
            fontsize=11.5, color=text_color, weight="bold",
            ha="center", va="center")
    for i, line in enumerate(desc_lines):
        ax.text(tx, ty - 0.04 - i * 0.043, line,
                fontsize=8.5, color=text_color,
                ha="center", va="center")

# =================== CENTER WHITE CIRCLE ===================
ax.add_patch(Circle((0, 0), R_IN, facecolor=BG, edgecolor="white",
                    linewidth=2, zorder=3))

# =================== PESSOA ICON (Path-based proper silhouette) ===================
# Standing person silhouette: head (circle) + body (trapezoid) + arms (rect) + legs (rect)
# Centered at (0, 0.08), height ~0.18
# Head circle
HEAD_CY = 0.13
HEAD_R = 0.030
ax.add_patch(Circle((0, HEAD_CY), HEAD_R, facecolor=SOFT, edgecolor="none", zorder=5))

# Body silhouette (trapezoid: wide shoulders to narrow waist, then legs)
# Using Path for SVG-like shape
body_pts = [
    (-0.060, 0.085),   # left shoulder
    ( 0.060, 0.085),   # right shoulder
    ( 0.050, 0.055),   # right under-arm
    ( 0.038, 0.030),   # right waist
    ( 0.022, -0.010),  # right outer leg top
    ( 0.022, -0.075),  # right foot
    ( 0.005, -0.075),  # inner right foot
    ( 0.005, -0.010),  # crotch (right)
    (-0.005, -0.010),  # crotch (left)
    (-0.005, -0.075),  # inner left foot
    (-0.022, -0.075),  # left foot
    (-0.022, -0.010),  # left outer leg top
    (-0.038, 0.030),   # left waist
    (-0.050, 0.055),   # left under-arm
    (-0.060, 0.085),   # close
]
body_codes = [Path.MOVETO] + [Path.LINETO] * (len(body_pts) - 2) + [Path.CLOSEPOLY]
body_path = Path(body_pts, body_codes)
ax.add_patch(PathPatch(body_path, facecolor=SOFT, edgecolor="none", zorder=5))

# =================== CENTER TEXT ===================
ax.text(0, -0.14, "LONGEVIDADE", fontsize=10, color=INK, weight="bold",
        ha="center", va="center", zorder=6)
ax.text(0, -0.205, "FUNCIONAL", fontsize=10, color=INK, weight="bold",
        ha="center", va="center", zorder=6)

# =================== CORNER LABELS ===================
# Function: figure-norm position of a point on donut outer at given angle (degrees)
# ax position: [0.20, 0.18, 0.60, 0.60], center fig (0.50, 0.48)
DONUT_CX_FIG = 0.20 + 0.60 / 2  # 0.50
DONUT_CY_FIG = 0.18 + 0.60 / 2  # 0.48
# R_OUT=0.85 in ax data, ax half-width=0.30 fig (since width=0.60), so R_OUT_FIG = 0.85*0.30 = 0.255
R_OUT_FIG = R_OUT * 0.30

def donut_edge_fig(angle_deg):
    rad = np.radians(angle_deg)
    return (DONUT_CX_FIG + R_OUT_FIG * np.cos(rad),
            DONUT_CY_FIG + R_OUT_FIG * np.sin(rad))

def draw_label(title_pos, title_text, lines, dot_color, ha="left"):
    """Draw label with bold title, divider, description lines, and HORIZONTAL leader.
    Leader is at the divider Y level, going to where it intersects the donut edge.
    Description lines are BELOW divider, so leader (above them) doesn't cross."""
    tx, ty = title_pos
    div_y = ty - 0.022
    # Title
    fig.text(tx, ty, title_text, fontsize=12, color=INK,
             weight="bold", va="center", ha=ha)
    # Description lines BELOW divider
    for i, line in enumerate(lines):
        fig.text(tx, ty - 0.052 - i * 0.022, line,
                 fontsize=9.5, color=SOFT, va="center", ha=ha)
    # HORIZONTAL leader from divider end to donut edge at the SAME Y level
    R = 0.255  # R_OUT_FIG
    cx, cy = 0.50, 0.48
    dy = div_y - cy
    if abs(dy) >= R:
        return  # leader Y outside donut
    dx = (R**2 - dy**2) ** 0.5
    # Determine which side of donut the leader hits
    if ha == "left":
        # Leader from label going RIGHT to donut LEFT edge
        dot_x = cx - dx
        div_x0 = tx
        div_x1 = tx + 0.060
        leader_start = div_x1 + 0.005
    elif ha == "right":
        # Leader from label going LEFT to donut RIGHT edge
        dot_x = cx + dx
        div_x0 = tx
        div_x1 = tx - 0.060
        leader_start = div_x1 - 0.005
    else:
        return  # center labels use vertical leader (separate)
    dot_y = div_y
    # Divider line
    fig.lines.append(plt.Line2D([div_x0, div_x1], [div_y, div_y],
                                color=SOFT, linewidth=1.0, transform=fig.transFigure))
    # Horizontal leader
    fig.lines.append(plt.Line2D([leader_start, dot_x], [div_y, div_y],
                                color=dot_color, linewidth=1.0,
                                transform=fig.transFigure, zorder=9))
    # Dot at donut edge
    fig.patches.append(plt.Circle((dot_x, dot_y), 0.011, facecolor=dot_color,
                                  edgecolor="white", linewidth=1.5,
                                  transform=fig.transFigure, zorder=10))
    # Dot at label end (small)
    fig.patches.append(plt.Circle((leader_start, div_y), 0.006, facecolor=dot_color,
                                  edgecolor="none",
                                  transform=fig.transFigure, zorder=10))

# TOP-LEFT: "A BASE DE TUDO" → connects to Zona 2 segment via HORIZONTAL leader
draw_label((0.045, 0.720), "A BASE DE TUDO",
           ["Eficiência", "metabólica", "e resistência.", "",
            "Deve ocupar", "a maior parte", "do seu tempo."],
           dot_color=S_ZONA2, ha="left")

# TOP-RIGHT: "O ESCUDO" → connects to Força segment
draw_label((0.955, 0.720), "O ESCUDO",
           ["Preserva massa", "muscular e protege",
            "o metabolismo", "ao longo da vida."],
           dot_color=S_FORCA, ha="right")

# MID-RIGHT: "O ESTÍMULO" → connects to HIIT segment
draw_label((0.955, 0.380), "O ESTÍMULO",
           ["Melhora a", "capacidade",
            "cardiorrespiratória", "máxima."],
           dot_color=S_HIIT, ha="right")

# BOTTOM: "A BASE DA LIBERDADE" → connects to Mobilidade (~268°)
# Centered below donut
bottom_dot_x, bottom_dot_y = donut_edge_fig(268)
fig.patches.append(plt.Circle((bottom_dot_x, bottom_dot_y), 0.011, facecolor=S_MOB,
                              edgecolor="white", linewidth=1.5,
                              transform=fig.transFigure, zorder=10))
fig.lines.append(plt.Line2D([bottom_dot_x, 0.500],
                            [bottom_dot_y, 0.180],
                            color=S_MOB, linewidth=1.0,
                            transform=fig.transFigure, zorder=9))
fig.text(0.500, 0.165, "A BASE DA LIBERDADE",
         fontsize=12, color=INK, weight="bold", va="center", ha="center")
fig.text(0.500, 0.140, "Mantém amplitude de movimento e previne quedas.",
         fontsize=9.5, color=SOFT, va="center", ha="center")

# =================== CALLOUT ===================
fig.text(0.500, 0.105,
         "Exercício para longevidade não é um tipo de treino — é um sistema.",
         fontsize=13, color=INK, weight="bold", style="italic", ha="center")

# =================== FOOTER ===================
fig.lines.append(plt.Line2D([0.030, 0.970], [0.080, 0.080],
                            color="#BBBBBB", linewidth=0.5, transform=fig.transFigure))
fig.text(0.030, 0.058,
         "Proporção de tempo sugerida para um adulto saudável. Ajustes devem ser feitos conforme",
         fontsize=10, color=FOOT, va="center")
fig.text(0.030, 0.038,
         "idade, histórico clínico e objetivos individuais.",
         fontsize=10, color=FOOT, va="center")
fig.text(0.030, 0.015,
         "Fontes: Diretrizes ACSM, 2021; Ekkekakis et al., 2020; Pedersen & Saltin, 2015.",
         fontsize=10, color=FOOT, va="center")

# =================== EXPORT ===================
out_dir = _Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap07_Fig03.pdf"
png_path = out_dir / "_preview_Cap07_Fig03.png"
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG)
plt.savefig(png_path, dpi=170, facecolor=BG)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
