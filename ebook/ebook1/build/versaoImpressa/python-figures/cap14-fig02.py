"""
Cap14 Fig02 (PT-BR, B&W vetorial) — A regularidade vence a duração: o novo alvo do sono.
Matriz 2x2 (Duração × Regularidade).
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, FancyBboxPatch

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

CELL_OK   = "#F0F0F0"   # baseline
CELL_MID1 = "#E0E0E0"   # leve
CELL_MID2 = "#C8C8C8"   # médio
CELL_BAD  = "#A8A8A8"   # alto

fig = plt.figure(figsize=(11.0, 7.5))
fig.patch.set_facecolor(BG)

LEFT = 0.025

# Título
fig.text(LEFT, 0.945,
         "Figura 2 — A regularidade vence a duração: o novo alvo do sono.",
         fontsize=15, color=INK, weight="bold")
fig.text(LEFT, 0.910,
         "Risco relativo de mortalidade por todas as causas (referência = 1,0).",
         fontsize=9.5, color=INK_SOFT, style="italic")

# Matriz 2x2
# Position layout: linhas/colunas labels + 4 células
# Headers das colunas
COL_HEAD_Y = 0.835
fig.text(0.5, COL_HEAD_Y + 0.025, "REGULARIDADE DO SONO",
         fontsize=10, color=INK_SOFT, weight="bold", ha="center")

COL_CENTERS = [0.45, 0.75]
col_labels = [
    ("REGULAR",   "(horários consistentes)"),
    ("IRREGULAR", "(variação > 90 min entre dias)"),
]
for cx, (lbl, sub) in zip(COL_CENTERS, col_labels):
    fig.text(cx, COL_HEAD_Y, lbl,
             fontsize=10, color=INK, weight="bold", ha="center")
    fig.text(cx, COL_HEAD_Y - 0.020, sub,
             fontsize=8, color=INK_SOFT, ha="center", style="italic")

# Headers das linhas
ROW_HEAD_X = 0.12
fig.text(ROW_HEAD_X, 0.5, "DURAÇÃO DO SONO",
         fontsize=10, color=INK_SOFT, weight="bold",
         ha="center", va="center", rotation=90)

ROW_CENTERS = [0.65, 0.35]
row_labels = [
    ("ADEQUADA",   "(7–8 horas)"),
    ("INADEQUADA", "(< 7 horas\nou > 9 horas)"),
]
for ry, (lbl, sub) in zip(ROW_CENTERS, row_labels):
    fig.text(ROW_HEAD_X + 0.08, ry, lbl,
             fontsize=10, color=INK, weight="bold",
             ha="center", va="center")
    fig.text(ROW_HEAD_X + 0.08, ry - 0.030, sub,
             fontsize=8, color=INK_SOFT, ha="center", va="center",
             style="italic", linespacing=1.2)

# 4 células
CELL_W = 0.26
CELL_H = 0.24

cells = [
    # (col_idx, row_idx, color, big_text, sub_text)
    (0, 0, CELL_OK,   "1,0",        "REFERÊNCIA"),
    (1, 0, CELL_MID2, "+25% a +30%", "MAIOR RISCO QUE REGULAR\nCOM DURAÇÃO INADEQUADA"),
    (0, 1, CELL_MID1, "+15% a +20%", "RISCO MODERADAMENTE\nAUMENTADO"),
    (1, 1, CELL_BAD,  "+40% a +48%", "MAIOR RISCO"),
]

for col_i, row_i, color, big, sub in cells:
    cx = COL_CENTERS[col_i]
    cy = ROW_CENTERS[row_i]

    fig.patches.append(FancyBboxPatch(
        (cx - CELL_W/2, cy - CELL_H/2), CELL_W, CELL_H,
        boxstyle="round,pad=0.003,rounding_size=0.008",
        facecolor=color, edgecolor=INK, linewidth=0.8,
        transform=fig.transFigure, zorder=1
    ))
    # Cabeçalho da célula (regular + adequado, irregular + adequado etc)
    titles = {
        (0, 0): "Regular + Adequado",
        (1, 0): "Irregular + Adequado",
        (0, 1): "Regular + Inadequado",
        (1, 1): "Irregular + Inadequado",
    }
    fig.text(cx, cy + CELL_H/2 - 0.018, titles[(col_i, row_i)],
             fontsize=8.5, color=INK, weight="bold", ha="center", va="top",
             style="italic")
    # Número grande
    fig.text(cx, cy + 0.010, big,
             fontsize=18 if "REF" in sub else 17,
             color=INK, weight="bold", ha="center", va="center")
    # Sub-text
    fig.text(cx, cy - 0.045, sub,
             fontsize=8, color=INK_SOFT, ha="center", va="center",
             linespacing=1.25, weight="bold")

# Annotation crítica no centro entre as 4 células
fig.text(0.60, 0.50,
         "Sono IRREGULAR\ncom duração normal\né PIOR que sono\nREGULAR com\nduração curta.",
         fontsize=8, color=INK, weight="bold", ha="center", va="center",
         linespacing=1.3, style="italic",
         bbox=dict(boxstyle="round,pad=0.5", facecolor="white",
                   edgecolor=INK, linewidth=0.8))

# Footer left
fig.text(LEFT, 0.105,
         "Derivado de análises de UK Biobank com seguimento objetivo em > 60.000 participantes (Windred et al.,",
         fontsize=7.5, color=FOOT)
fig.text(LEFT, 0.090,
         "Sleep, 2024) e MESA Sleep Study (Huang et al., JACC, 2020). Valores representativos — não precisões pontuais —,",
         fontsize=7.5, color=FOOT)
fig.text(LEFT, 0.075,
         "os estudos usam métricas diferentes (Sleep Regularity Index, variabilidade), mas convergem na direção.",
         fontsize=7.5, color=FOOT)

# Bottom right text
fig.text(0.97, 0.105,
         "Dormir 7 horas todos os dias do mesmo horário",
         fontsize=8.5, color=INK, ha="right")
fig.text(0.97, 0.087,
         "é melhor que dormir 8 horas com horários flutuantes.",
         fontsize=8.5, color=INK, weight="bold", ha="right")

fig.text(0.97, 0.060,
         "Alvo: horários regulares, inclusive aos fins de semana.",
         fontsize=8.5, color=INK_SOFT, ha="right", style="italic")

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap14_Fig02.pdf"
png_path = out_dir / "_preview_Cap14_Fig02.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
