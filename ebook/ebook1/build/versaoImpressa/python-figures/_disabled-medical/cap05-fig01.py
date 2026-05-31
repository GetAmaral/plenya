"""
Cap05 Fig01 (PT-BR, B&W vetorial) — Escore de Cálcio Coronariano: O Que Cada Faixa Significa.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, Circle, FancyArrowPatch
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
BAND_1   = "#F4F4F4"
BAND_2   = "#E2E2E2"
BAND_3   = "#C8C8C8"
BAND_4   = "#A8A8A8"
CALCIUM  = "#000000"
WALL     = "#D6D6D6"

fig = plt.figure(figsize=(11.0, 8.0))
fig.patch.set_facecolor(BG)

LEFT = 0.025

# Título
fig.text(LEFT, 0.95, "Figura 1 — Escore de Cálcio Coronariano: O Que Cada Faixa Significa",
         fontsize=15, color=INK, weight="bold")
fig.text(LEFT, 0.910,
         "Uma escala contínua de risco. Quanto maior o escore, maior a carga de placa aterosclerótica e o risco futuro.",
         fontsize=9.5, color=INK_SOFT)

# ---------- 4 faixas (barras horizontais) ----------
zones = [
    ("CAC = 0",       "O poder do zero",          BAND_1, 0.04, 0.27),
    ("CAC 1–99",      "A aterosclerose começou",  BAND_2, 0.28, 0.51),
    ("CAC 100–399",   "Carga significativa",      BAND_3, 0.52, 0.75),
    ("CAC ≥ 400",     "Prevenção secundária",     BAND_4, 0.76, 0.97),
]

BAR_Y1, BAR_Y2 = 0.785, 0.840
for title, sub, color, x1, x2 in zones:
    fig.patches.append(Rectangle(
        (x1, BAR_Y1), x2 - x1, BAR_Y2 - BAR_Y1,
        facecolor=color, edgecolor="none", transform=fig.transFigure, zorder=1
    ))
    # Título da faixa
    text_color = INK if color != BAND_4 else "white"
    fig.text((x1 + x2) / 2, 0.825, title,
             fontsize=11, color=text_color, weight="bold",
             ha="center", va="center")
    fig.text((x1 + x2) / 2, 0.798, sub,
             fontsize=8.5, color=text_color, style="italic",
             ha="center", va="center")

# ---------- escala numérica abaixo da barra ----------
SCALE_Y = 0.755
# Ticks: 0, 100, 400 (logaritmico-ish visualmente)
ticks = [(0.040, "0"), (0.280, "100"), (0.520, "400")]
for x, lbl in ticks:
    fig.lines.append(plt.Line2D(
        [x, x], [BAR_Y1 - 0.005, BAR_Y1 - 0.020],
        color=INK, linewidth=0.8, transform=fig.transFigure
    ))
    fig.text(x, SCALE_Y, lbl, fontsize=9, color=INK,
             ha="center", va="top", weight="bold")

# Marcadores Ricardo e Marcos
RICARDO_X = 0.380  # CAC ~187 visualmente na faixa 100-399
MARCOS_X  = 0.870  # CAC = 412 visualmente na faixa ≥400

# Ricardo
fig.lines.append(plt.Line2D(
    [RICARDO_X, RICARDO_X], [BAR_Y1 + 0.005, BAR_Y2 - 0.005],
    color=INK, linewidth=2.0, transform=fig.transFigure, zorder=3
))
fig.text(RICARDO_X, SCALE_Y, "Ricardo (52 anos*)\nCAC ≈ 187",
         fontsize=8, color=INK, weight="bold", ha="center", va="top",
         linespacing=1.2)

# Marcos
fig.lines.append(plt.Line2D(
    [MARCOS_X, MARCOS_X], [BAR_Y1 + 0.005, BAR_Y2 - 0.005],
    color=INK, linewidth=2.0, transform=fig.transFigure, zorder=3
))
fig.text(MARCOS_X, SCALE_Y, "Marcos (57 anos*)\nCAC = 412",
         fontsize=8, color=INK, weight="bold", ha="center", va="top",
         linespacing=1.2)

# ---------- 4 blocos de descrição clínica ----------
desc_blocks = [
    ("Risco muito baixo",
     "nos próximos 5–10 anos.\nEstatina pode ser postergada\nse risco for incerto."),
    ("Risco baixo a moderado.",
     "Conduta depende do contexto\nclínico, percentil para idade\ne sexo."),
    ("Estatina recomendada",
     "(ACC/AHA).\nO risco cai ~14% do risco\ncardiovascular."),
    ("Terapia intensiva.",
     "Equivalente a já ter evento\ncardiovascular."),
]

DESC_Y = 0.640
for i, (title, body) in enumerate(desc_blocks):
    x_center = (zones[i][3] + zones[i][4]) / 2
    fig.text(x_center, DESC_Y, title,
             fontsize=10, color=INK, weight="bold", ha="center", va="top")
    fig.text(x_center, DESC_Y - 0.030, body,
             fontsize=8.5, color=INK_SOFT, ha="center", va="top",
             linespacing=1.3)

# ---------- 4 ilustrações de artéria (cross-section com cálcio progressivo) ----------
fig.text(LEFT, 0.380, "EXEMPLOS\nDE EXAMES",
         fontsize=8.5, color=INK_SOFT, weight="bold", va="center", linespacing=1.15)

# Posicionar 4 axes (artérias)
AXES_Y = 0.30
AXES_H = 0.18
for i, n_calcium in enumerate([0, 2, 8, 25]):
    x1 = zones[i][3]
    x2 = zones[i][4]
    ax = fig.add_axes([x1 + 0.04, AXES_Y, (x2 - x1 - 0.08), AXES_H])
    ax.set_xlim(-1.2, 1.2)
    ax.set_ylim(-1.2, 1.2)
    ax.set_aspect("equal")
    ax.axis("off")
    # Anel arterial
    ax.add_patch(Circle((0, 0), 1.0, facecolor=WALL, edgecolor=INK, linewidth=1.0))
    ax.add_patch(Circle((0, 0), 0.72, facecolor="white", edgecolor=INK, linewidth=0.8))
    # Cálcio: deposições pequenas pontos pretos crescendo em número
    rng = np.random.default_rng(seed=42 + i)
    for _ in range(n_calcium):
        # posição aleatória no anel (entre 0.72 e 1.0)
        angle = rng.uniform(0, 2 * np.pi)
        r = rng.uniform(0.72, 0.95)
        cx = r * np.cos(angle)
        cy = r * np.sin(angle)
        size = rng.uniform(0.04, 0.10)
        ax.add_patch(Circle((cx, cy), size, facecolor=CALCIUM,
                            edgecolor="none", zorder=3))

    # Label da quantidade
    fig.text((x1 + x2) / 2, AXES_Y - 0.025, zones[i][0],
             fontsize=9, color=INK, weight="bold", ha="center", va="top")

# ---------- footer ----------
FOOT_Y = 0.075
fig.text(LEFT, FOOT_Y,
         "Cada duplicação do escore = ~14% ↑ risco cardiovascular.",
         fontsize=9, color=INK_SOFT, style="italic")
fig.text(LEFT, FOOT_Y - 0.025,
         "*idade no momento de Ricardo no infarto; idade do CAC base do Marcos.",
         fontsize=7.5, color=FOOT)
fig.text(LEFT, FOOT_Y - 0.045,
         "* CAC: escore Agatston. Fontes: Greenland P et al, JACC, 2018; Mortensen MB et al, JAMA Cardiol, 2019.",
         fontsize=7.5, color=FOOT)

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap05_Fig01.pdf"
png_path = out_dir / "_preview_Cap05_Fig01.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
