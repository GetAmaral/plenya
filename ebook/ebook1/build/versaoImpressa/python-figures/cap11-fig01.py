"""
Cap11 Fig01 (PT-BR, B&W vetorial) — Paulo: 6 meses sem reposição.

Chart comparativo de 5 biomarcadores (TESTOSTERONA TOTAL, LIVRE, VITAMINA D,
hs-CRP, IDADE EPIGENÉTICA) com 3 marcadores cada: Antes (○), Depois (●),
Alvo (▲). Backgrounds em zonas: salmão (ruim), verde (alvo), peach (intermediário).

Posições e ranges MEDIDOS do original 1536×1024 (ticks detectados via OCR de strip
abaixo de cada axis_y):
  Row 1 (T.TOTAL):   y=307, vmin=200, vmax=800, CL=308,  CR=1391
  Row 2 (T.LIVRE):   y=431, vmin=0,   vmax=18,  CL=265,  CR=1398
  Row 3 (VIT D):     y=553, vmin=0,   vmax=80,  CL=265,  CR=1379
  Row 4 (hs-CRP):    y=674, vmin=0,   vmax=3,   CL=265,  CR=1363
  Row 5 (IDADE):     y=786, vmin=-2,  vmax=6,   CL=281,  CR=1395
"""
from pathlib import Path as _Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import (
    Rectangle, FancyBboxPatch, Polygon, Circle, FancyArrowPatch
)

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
ZONE_BAD  = "#D8D8D8"   # darker (pink → escuro)
ZONE_INT  = "#EDEDED"   # médio (peach)
ZONE_OK   = "#F6F6F6"   # mais claro (verde)

W_IMG, H_IMG = 1536, 1024
_FIG_W = 10.0
_FIG_H = _FIG_W * H_IMG / W_IMG
fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

ax = fig.add_axes([0, 0, 1, 1])
ax.set_xlim(0, W_IMG); ax.set_ylim(H_IMG, 0)
ax.set_aspect("equal"); ax.axis("off")

# ============================================================
# CABEÇALHO
# ============================================================
ax.text(40, 38, "FIGURA 3",
        fontsize=10, color=INK, weight="semibold", va="center", ha="left",
        family="sans-serif")

ax.text(40, 88,
        "Paulo: 6 meses sem reposição — e a trajetória mudou",
        fontsize=22, color=INK, weight="bold", va="center", ha="left")

ax.text(40, 138,
        "Otimização de sono, vitamina D e treino de força ajustado — sem uso de reposição de testosterona.",
        fontsize=11, color=SOFT, va="center", ha="left")

# ============================================================
# LEGENDA
# ============================================================
LEGEND_Y = 200
ax.add_patch(Circle((310, LEGEND_Y), 9, facecolor=BG, edgecolor=INK,
                     linewidth=2.0, zorder=5))
ax.text(330, LEGEND_Y, "Antes (baseline)",
        fontsize=11, color=INK, va="center", ha="left")

ax.add_patch(Circle((570, LEGEND_Y), 9, facecolor=INK, edgecolor="none", zorder=5))
ax.text(590, LEGEND_Y, "Depois (6 meses)",
        fontsize=11, color=INK, va="center", ha="left")

tri_pts = [(840, LEGEND_Y-10), (852, LEGEND_Y+8), (828, LEGEND_Y+8)]
ax.add_patch(Polygon(tri_pts, closed=True, facecolor=INK, edgecolor="none", zorder=5))
ax.text(862, LEGEND_Y, "Alvo ótimo",
        fontsize=11, color=INK, va="center", ha="left")

# ============================================================
# CHART AREA — 5 rows
# ============================================================
ROW_HEIGHT = 50
# Note: posições dos marcadores no original NÃO correspondem ao valor literal da label.
# O original posiciona em valores "redondos" visualmente legíveis e mostra a label numérica
# do dado clínico real ao lado. Mantemos a label do dado, mas posicionamos onde o original posiciona.
ROWS = [
    {
        "name": ["TESTOSTERONA", "TOTAL"], "unit": "(ng/dL)",
        "y_axis": 307,
        "vmin": 200, "vmax": 800,
        "chart_left": 308, "chart_right": 1391,
        "zone_left": 260, "zone_right": 1410,
        "ticks": [200, 300, 400, 500, 600, 700, 800],
        "antes_pos": 320, "depois_pos": 500, "alvo_pos": 560,
        "antes_label": "310", "depois_label": "485", "alvo_label": "> 500",
        "zones": [(200, 300, "bad"), (300, 500, "int"), (500, 720, "ok"), (720, 800, "int")],
    },
    {
        "name": ["TESTOSTERONA", "LIVRE"], "unit": "(pg/mL)",
        "y_axis": 431,
        "vmin": 0, "vmax": 18,
        "chart_left": 265, "chart_right": 1398,
        "zone_left": 260, "zone_right": 1410,
        "ticks": [0, 2, 4, 6, 8, 10, 12, 14, 16, 18],
        "antes_pos": 4.5, "depois_pos": 10.5, "alvo_pos": 12,
        "antes_label": "4,8", "depois_label": "11,2", "alvo_label": "> 10",
        "zones": [(0, 3, "bad"), (3, 9, "int"), (9, 16, "ok"), (16, 18, "int")],
    },
    {
        "name": ["VITAMINA D", "(25-OH)"], "unit": "(ng/mL)",
        "y_axis": 553,
        "vmin": 0, "vmax": 80,
        "chart_left": 265, "chart_right": 1379,
        "zone_left": 260, "zone_right": 1410,
        "ticks": [0, 10, 20, 30, 40, 50, 60, 70, 80],
        "antes_pos": 21, "depois_pos": 60, "alvo_pos": 50,
        "antes_label": "24", "depois_label": "58", "alvo_label": "40–60",
        "zones": [(0, 20, "bad"), (20, 40, "int"), (40, 67, "ok"), (67, 80, "int")],
    },
    {
        "name": ["hs-CRP"], "unit": "(mg/L)",
        "y_axis": 674,
        "vmin": 0, "vmax": 3.0,
        "chart_left": 265, "chart_right": 1363,
        "zone_left": 260, "zone_right": 1410,
        "ticks": [0, 0.5, 1.0, 1.5, 2.0, 2.5, 3.0],
        "antes_pos": 1.65, "depois_pos": 1.0, "alvo_pos": 0.75,
        "antes_label": "1,7", "depois_label": "0,9", "alvo_label": "< 1,0",
        "zones": [(0, 1.3, "ok"), (1.3, 2.0, "int"), (2.0, 3.0, "bad")],
        "tick_labels": ["0", "0,5", "1,0", "1,5", "2,0", "2,5", "3,0"],
    },
    {
        "name": ["IDADE EPIGENÉTICA"], "subunit": "(relativa à cronológica)", "unit": "(anos)",
        "y_axis": 786,
        "vmin": -2, "vmax": 6,
        "chart_left": 281, "chart_right": 1395,
        "zone_left": 260, "zone_right": 1410,
        "ticks": [-2, -1, 0, 1, 2, 3, 4, 5, 6],
        "antes_pos": 4, "depois_pos": 2, "alvo_pos": 0,
        "antes_label": "+4", "depois_label": "+2", "alvo_label": "≤ 0",
        "zones": [(-2, 0, "ok"), (0, 2.5, "int"), (2.5, 6, "bad")],
        "tick_labels": ["−2", "−1", "0", "+1", "+2", "+3", "+4", "+5", "+6"],
    },
]

LABEL_X = 40

def value_to_x(val, row):
    return row["chart_left"] + (val - row["vmin"]) / (row["vmax"] - row["vmin"]) * (row["chart_right"] - row["chart_left"])

for row in ROWS:
    y_axis = row["y_axis"]
    y_top = y_axis - ROW_HEIGHT

    # Label (biomarker name + unit)
    label_y_start = y_axis - 38
    for j, line in enumerate(row["name"]):
        ax.text(LABEL_X, label_y_start + j * 19, line,
                fontsize=11, color=INK, weight="bold",
                ha="left", va="center")
    extra_y = label_y_start + len(row["name"]) * 19
    if "subunit" in row:
        ax.text(LABEL_X, extra_y, row["subunit"],
                fontsize=9.5, color=SOFT, ha="left", va="center")
        extra_y += 17
    ax.text(LABEL_X, extra_y, row["unit"],
            fontsize=9.5, color=SOFT, ha="left", va="center")

    # Zone backgrounds
    for v0, v1, ztype in row["zones"]:
        x0 = value_to_x(v0, row)
        x1 = value_to_x(v1, row)
        if v0 == row["vmin"]: x0 = row["zone_left"]
        if v1 == row["vmax"]: x1 = row["zone_right"]
        color = {"bad": ZONE_BAD, "ok": ZONE_OK, "int": ZONE_INT}[ztype]
        ax.add_patch(Rectangle((x0, y_top), x1-x0, ROW_HEIGHT,
                                facecolor=color, edgecolor="none", zorder=1))

    # Axis line
    ax.plot([row["zone_left"], row["zone_right"]], [y_axis, y_axis],
            color=GRAY1, linewidth=0.8, zorder=2)

    # Tick marks and labels
    tick_labels = row.get("tick_labels", [str(t) for t in row["ticks"]])
    for tick, lbl in zip(row["ticks"], tick_labels):
        tx = value_to_x(tick, row)
        ax.plot([tx, tx], [y_axis, y_axis+4], color=GRAY1, linewidth=0.7, zorder=2)
        ax.text(tx, y_axis + 18, lbl,
                fontsize=9, color=SOFT, ha="center", va="center")

    # ALVO marker (triangle) — bottom touches axis line
    alvo_x = value_to_x(row["alvo_pos"], row)
    tri_half_w = 12
    tri_h = 22
    tri = [(alvo_x, y_axis + 2 - tri_h),     # top
           (alvo_x + tri_half_w, y_axis + 2),  # right
           (alvo_x - tri_half_w, y_axis + 2)]  # left
    ax.add_patch(Polygon(tri, closed=True, facecolor=INK, edgecolor="none", zorder=5))
    ax.text(alvo_x, y_axis - 36, row["alvo_label"],
            fontsize=9.5, color=INK, weight="bold",
            ha="center", va="center")

    # ANTES / DEPOIS
    antes_x = value_to_x(row["antes_pos"], row)
    depois_x = value_to_x(row["depois_pos"], row)
    marker_y = y_axis - 16

    # Arrow from ANTES to DEPOIS — discrete: thin line + tiny triangular head
    sign = 1 if depois_x > antes_x else -1
    # Stop line a bit before the dot, head sits between line-end and dot
    line_start_x = antes_x + sign * 8     # leave gap from antes ring
    head_tip_x = depois_x - sign * 8       # tip just before dot
    head_base_x = head_tip_x - sign * 7    # head length 7
    ax.plot([line_start_x, head_base_x], [marker_y, marker_y],
            color=INK, linewidth=0.9, solid_capstyle="butt", zorder=4)
    head_tri = [(head_tip_x, marker_y),
                (head_base_x, marker_y - 3),
                (head_base_x, marker_y + 3)]
    ax.add_patch(Polygon(head_tri, closed=True, facecolor=INK,
                          edgecolor="none", zorder=4))

    # ANTES hollow circle
    ax.add_patch(Circle((antes_x, marker_y), 8,
                         facecolor=BG, edgecolor=INK,
                         linewidth=1.8, zorder=6))
    ax.text(antes_x, y_axis - 36, row["antes_label"],
            fontsize=9.5, color=INK, weight="bold",
            ha="center", va="center")

    # DEPOIS filled circle
    ax.add_patch(Circle((depois_x, marker_y), 8,
                         facecolor=INK, edgecolor="none", zorder=6))
    ax.text(depois_x, y_axis - 36, row["depois_label"],
            fontsize=9.5, color=INK, weight="bold",
            ha="center", va="center")

# ============================================================
# DIAGNOSTIC BOX (bottom-left)
# ============================================================
diag_top = 850
diag_h = 80
diag_w = 360
diag_x = 40

ax.add_patch(FancyBboxPatch(
    (diag_x + 4, diag_top + 4), diag_w - 8, diag_h - 8,
    boxstyle="round,pad=2,rounding_size=10",
    facecolor=BG, edgecolor=INK, linewidth=0.8, zorder=3))

chk_cx, chk_cy = diag_x + 44, diag_top + diag_h/2
ax.add_patch(Circle((chk_cx, chk_cy), 16, facecolor=INK, edgecolor="none", zorder=4))
ax.plot([chk_cx-7, chk_cx-1.5, chk_cx+8], [chk_cy, chk_cy+6, chk_cy-7],
        color="white", linewidth=2.4, solid_capstyle="round",
        solid_joinstyle="round", zorder=5)

text_left = chk_cx + 26
ax.text(text_left, diag_top + 22, "Diagnóstico correto.",
        fontsize=11, color=INK, weight="bold", va="center", ha="left")
ax.text(text_left, diag_top + 42, "Plano integrado.",
        fontsize=11, color=INK, weight="bold", va="center", ha="left")
ax.text(text_left, diag_top + 62, "Seguimento.",
        fontsize=11, color=INK, weight="bold", va="center", ha="left")

# ============================================================
# SOURCE/FOOTER text (right of diagnostic box)
# ============================================================
foot_left = diag_x + diag_w + 30
ax.text(foot_left, diag_top + 8,
        "Em escuro, valores de Paulo na primeira consulta. Em claro, valores 6 meses depois.",
        fontsize=9.5, color=FOOT, va="center", ha="left")
ax.text(foot_left, diag_top + 28,
        "As setas mostram a direção da mudança; o triângulo marca o alvo ótimo para longevidade.",
        fontsize=9.5, color=FOOT, va="center", ha="left")
ax.text(foot_left, diag_top + 48,
        "A idade epigenética, medida por relógios de metilação do DNA, reduziu 2 anos —",
        fontsize=9.5, color=FOOT, va="center", ha="left")
ax.text(foot_left, diag_top + 68,
        "sinal biológico de que o envelhecimento desacelerou.",
        fontsize=9.5, color=FOOT, va="center", ha="left")

# Divider line above DHEA-S note
ax.plot([40, W_IMG-40], [diag_top + diag_h + 18, diag_top + diag_h + 18],
        color=GRAY1, linewidth=0.5)
ax.text(40, diag_top + diag_h + 36,
        "DHEA-S aumentou de 145 para 210 µg/dL no período (não exibido no gráfico para manter a clareza visual).",
        fontsize=9, color=FOOT, va="center", ha="left")

# ============================================================
# EXPORT
# ============================================================
out_dir = _Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap11_Fig01.pdf"
png_path = out_dir / "_preview_Cap11_Fig01.png"
plt.savefig(pdf_path, facecolor=BG, bbox_inches="tight", pad_inches=0.0)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches="tight", pad_inches=0.0)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
