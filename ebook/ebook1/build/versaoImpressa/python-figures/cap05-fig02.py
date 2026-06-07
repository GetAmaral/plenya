"""
Cap05 Fig02 (PT-BR, B&W vetorial) — Idade Cronológica vs. Idade Arterial.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, Circle, Ellipse

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
GRAY_DOT = "#888888"

fig = plt.figure(figsize=(11.0, 6.4))
fig.patch.set_facecolor(BG)

LEFT = 0.025

# ---------- "FIGURA 2" tag (caixa preta com texto branco como label) ----------
TAG_X1, TAG_Y1 = LEFT, 0.910
TAG_W, TAG_H   = 0.085, 0.045
fig.patches.append(Rectangle(
    (TAG_X1, TAG_Y1), TAG_W, TAG_H,
    facecolor=INK, edgecolor=INK, transform=fig.transFigure, zorder=2
))
fig.text(TAG_X1 + TAG_W / 2, TAG_Y1 + TAG_H / 2, "FIGURA 2",
         fontsize=10, color="white", weight="bold",
         ha="center", va="center", zorder=3)

# Título — depois da tag
fig.text(TAG_X1 + TAG_W + 0.015, TAG_Y1 + TAG_H / 2,
         "Idade Cronológica vs. Idade Arterial",
         fontsize=18, color=INK, weight="bold", va="center")
fig.text(LEFT, 0.860,
         "O escore de cálcio traduzido em anos de envelhecimento arterial.",
         fontsize=10, color=INK_SOFT)

# ---------- Escala de IDADES compartilhada (mesma X scale pras 2 rows) ----------
# Mapeia idade → posição X em fig coords
AGE_MIN, AGE_MAX = 50, 85
AGE_X_START, AGE_X_END = 0.32, 0.92  # range visual

def age_to_x(age):
    return AGE_X_START + (age - AGE_MIN) / (AGE_MAX - AGE_MIN) * (AGE_X_END - AGE_X_START)

# Sub-headers das colunas — posicionados em "57" (típico cronológica) e "80" (extremo)
HEADER_Y = 0.760
fig.text(age_to_x(57), HEADER_Y, "IDADE CRONOLÓGICA",
         fontsize=10, color=INK, weight="bold", ha="center")
fig.text(age_to_x(57), HEADER_Y - 0.025, "(idade real)",
         fontsize=8.5, color=INK_SOFT, ha="center", style="italic")

fig.text(age_to_x(80), HEADER_Y, "IDADE ARTERIAL",
         fontsize=10, color=INK, weight="bold", ha="center")
fig.text(age_to_x(80), HEADER_Y - 0.025, "(equivalente pelo percentil 50 da MESA)",
         fontsize=8.5, color=INK_SOFT, ha="center", style="italic")

# ---------- 2 linhas de pacientes (cronológica + delta + arterial proporcionais à idade) ----------
patients = [
    # (nome, CAC, chrono_idade, delta_anos, arterial_idade)
    ("MARCOS",  "CAC 412",   57, 23, 80),
    ("RICARDO", "CAC ≈187",  52, 16, 68),
]

ROW_Y = [0.55, 0.28]

for (name, cac, chrono_age, delta_age, arterial_age), y in zip(patients, ROW_Y):
    chrono_x   = age_to_x(chrono_age)
    arterial_x = age_to_x(arterial_age)
    delta_x    = (chrono_x + arterial_x) / 2

    # Nome do paciente
    fig.text(0.10, y + 0.020, name,
             fontsize=18, color=INK, weight="bold",
             ha="center", va="center")
    fig.text(0.10, y - 0.015, cac,
             fontsize=9, color=INK_SOFT, ha="center", va="center")

    # Idade cronológica
    fig.text(chrono_x, y + 0.040, str(chrono_age),
             fontsize=30, color=INK, weight="bold",
             ha="center", va="center")
    fig.text(chrono_x, y + 0.010, "anos",
             fontsize=9, color=INK_SOFT, ha="center", va="center")

    # Delta (+anos) no meio do dumbbell desta row
    fig.text(delta_x, y + 0.040, f"+{delta_age}",
             fontsize=30, color=INK, weight="bold",
             ha="center", va="center")
    fig.text(delta_x, y + 0.010, "anos",
             fontsize=9, color=INK_SOFT, ha="center", va="center")

    # Idade arterial (com ~)
    fig.text(arterial_x, y + 0.040, f"~{arterial_age}",
             fontsize=30, color=INK, weight="bold",
             ha="center", va="center")
    fig.text(arterial_x, y + 0.010, "anos",
             fontsize=9, color=INK_SOFT, ha="center", va="center")

    # Linha horizontal proporcional à idade (de chrono até arterial)
    bar_y = y - 0.040
    fig.lines.append(plt.Line2D(
        [chrono_x, arterial_x], [bar_y, bar_y],
        color=INK, linewidth=2.5, transform=fig.transFigure, zorder=2
    ))

    # Dots: cinza esquerda (cronológica), preto direita (arterial)
    r = 0.012
    aspect = 11.0 / 6.4
    fig.patches.append(Ellipse(
        (chrono_x, bar_y), width=r*2, height=r*2*aspect,
        facecolor=GRAY_DOT, edgecolor=INK, linewidth=1.0,
        transform=fig.transFigure, zorder=4
    ))
    fig.patches.append(Ellipse(
        (arterial_x, bar_y), width=r*2, height=r*2*aspect,
        facecolor=INK, edgecolor=INK, linewidth=1.0,
        transform=fig.transFigure, zorder=4
    ))

# ---------- linha separadora ----------
SEP_Y = 0.13
fig.lines.append(plt.Line2D(
    [LEFT, 1 - LEFT], [SEP_Y, SEP_Y],
    color="#CFCFCF", linewidth=0.5, transform=fig.transFigure
))

# Footer — coração OUTLINE (♡), não filled (♥)
fig.text(0.5, 0.075,
         "♡  Suas artérias parecem ter anos a mais que você.",
         fontsize=13, color=INK, weight="bold", style="italic",
         ha="center")

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap05_Fig02.pdf"
png_path = out_dir / "_preview_Cap05_Fig02.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
