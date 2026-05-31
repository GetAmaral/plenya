"""
Cap12 Fig01 (PT-BR, B&W vetorial) — Como a ansiedade vira doença: a cascata do eixo HPA.
Diagrama de 4 andares descendentes.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, FancyBboxPatch, FancyArrowPatch, Ellipse

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

SHADE_1 = "#5A5A5A"   # estímulo — escuro
SHADE_2 = "#7A7A7A"   # HPA — médio escuro
SHADE_3 = "#9A9A9A"   # consequências — médio
SHADE_4 = "#B5B5B5"   # desfechos — médio claro
SHADE_INSIDE = "#F0F0F0"

_FIG_W, _FIG_H = 8.4, 11.4
_ASPECT = _FIG_W / _FIG_H

fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

LEFT_MARGIN = 0.03

# Título
fig.text(LEFT_MARGIN, 0.965,
         "Figura 1 — Como a ansiedade vira doença: a cascata do eixo HPA",
         fontsize=14, color=INK, weight="bold")

# Andares
def draw_floor(y_top, y_bot, color, number, title, subtitle):
    fig.patches.append(FancyBboxPatch(
        (LEFT_MARGIN, y_bot), 0.94, y_top - y_bot,
        boxstyle="round,pad=0.005,rounding_size=0.008",
        facecolor=color, edgecolor=INK, linewidth=0.8,
        transform=fig.transFigure, zorder=1
    ))
    # Bolinha numerada
    fig.patches.append(Ellipse(
        (LEFT_MARGIN + 0.040, y_top - 0.025),
        width=0.040, height=0.040 * _ASPECT,
        facecolor="white", edgecolor=INK, linewidth=1.0,
        transform=fig.transFigure, zorder=4
    ))
    fig.text(LEFT_MARGIN + 0.040, y_top - 0.025, number,
             fontsize=14, color=INK, weight="bold",
             ha="center", va="center", zorder=5)
    # Title
    fig.text(LEFT_MARGIN + 0.080, y_top - 0.020, title,
             fontsize=10, color="white", weight="bold", va="center")
    fig.text(LEFT_MARGIN + 0.080, y_top - 0.040, subtitle,
             fontsize=11, color="white", weight="bold", va="center")

# Andar 1 — ESTÍMULO PSICOLÓGICO
F1_TOP, F1_BOT = 0.910, 0.795
draw_floor(F1_TOP, F1_BOT, SHADE_1, "1",
           "ESTÍMULO PSICOLÓGICO",
           "Estresse crônico, ruminação, ansiedade")

# Conteúdo interno (3 exemplos)
F1_ICONS_Y = 0.815
for x_offset, label in [(0.10, "E-mails após\no expediente"),
                         (0.40, "Preocupação\ncom filhos"),
                         (0.70, "Pressão\nfinanceira")]:
    x = LEFT_MARGIN + x_offset
    fig.patches.append(FancyBboxPatch(
        (x, F1_ICONS_Y - 0.010), 0.18, 0.040,
        boxstyle="round,pad=0.002,rounding_size=0.005",
        facecolor=SHADE_INSIDE, edgecolor="none",
        transform=fig.transFigure, zorder=2
    ))
    fig.text(x + 0.090, F1_ICONS_Y + 0.010, label,
             fontsize=8, color=INK, ha="center", va="center",
             linespacing=1.15)

# Seta 1→2
fig.patches.append(FancyArrowPatch(
    (0.5, F1_BOT - 0.005), (0.5, 0.770),
    arrowstyle="->", color=INK, lw=1.4, mutation_scale=18,
    transform=fig.transFigure, zorder=3
))
fig.text(0.52, (F1_BOT + 0.770) / 2, "Percepção de ameaça pelo cérebro",
         fontsize=8, color=INK_SOFT, va="center", style="italic")

# Andar 2 — EIXO HPA ATIVADO
F2_TOP, F2_BOT = 0.760, 0.640
draw_floor(F2_TOP, F2_BOT, SHADE_2, "2",
           "EIXO HPA ATIVADO",
           "Eixo Hipotálamo–Hipófise–Adrenal em alerta permanente")

# 3 caixas internas (Hipotálamo, Hipófise, Suprarrenais)
F2_BOXES_Y = 0.665
for x_offset, label in [(0.04, "Hipotálamo\n→ CRH"),
                         (0.36, "Hipófise\n→ ACTH"),
                         (0.68, "Suprarrenais\n→ Cortisol")]:
    x = LEFT_MARGIN + x_offset
    fig.patches.append(FancyBboxPatch(
        (x, F2_BOXES_Y - 0.012), 0.22, 0.045,
        boxstyle="round,pad=0.002,rounding_size=0.005",
        facecolor=SHADE_INSIDE, edgecolor=INK, linewidth=0.5,
        transform=fig.transFigure, zorder=2
    ))
    fig.text(x + 0.110, F2_BOXES_Y + 0.010, label,
             fontsize=9, color=INK, weight="bold",
             ha="center", va="center", linespacing=1.15)

# Setas entre as caixas
for x_arrow_start, x_arrow_end in [(0.30, 0.36), (0.62, 0.68)]:
    fig.patches.append(FancyArrowPatch(
        (LEFT_MARGIN + x_arrow_start, F2_BOXES_Y + 0.010),
        (LEFT_MARGIN + x_arrow_end, F2_BOXES_Y + 0.010),
        arrowstyle="->", color=INK, lw=1.0, mutation_scale=10,
        transform=fig.transFigure, zorder=3
    ))

# Seta 2→3
fig.patches.append(FancyArrowPatch(
    (0.5, F2_BOT - 0.005), (0.5, 0.615),
    arrowstyle="->", color=INK, lw=1.4, mutation_scale=18,
    transform=fig.transFigure, zorder=3
))
fig.text(0.52, (F2_BOT + 0.615) / 2, "Cortisol cronicamente elevado",
         fontsize=8, color=INK_SOFT, va="center", style="italic")

# Andar 3 — CONSEQUÊNCIAS BIOQUÍMICAS
F3_TOP, F3_BOT = 0.605, 0.395
draw_floor(F3_TOP, F3_BOT, SHADE_3, "3",
           "CONSEQUÊNCIAS BIOQUÍMICAS",
           "Cascata bioquímica")

# 6 itens (2 colunas)
F3_ITEMS = [
    "Citocinas inflamatórias\n(IL-6, TNF-α, PCR) ↑",
    "Resistência insulínica ↑",
    "Gordura visceral ↑",
    "Hormônios sexuais ↓",
    "Função imune ↓",
    "Hipocampo / memória ↓",
]

for idx, item in enumerate(F3_ITEMS):
    col = idx // 3
    row = idx % 3
    x = LEFT_MARGIN + 0.04 + col * 0.46
    y = 0.535 - row * 0.040
    fig.text(x, y, "● " + item,
             fontsize=8.5, color=INK, weight="bold",
             linespacing=1.2, va="center")

# Seta 3→4
fig.patches.append(FancyArrowPatch(
    (0.5, F3_BOT - 0.005), (0.5, 0.370),
    arrowstyle="->", color=INK, lw=1.4, mutation_scale=18,
    transform=fig.transFigure, zorder=3
))
fig.text(0.52, (F3_BOT + 0.370) / 2, "Anos a décadas de exposição",
         fontsize=8, color=INK_SOFT, va="center", style="italic")

# Andar 4 — DESFECHOS CLÍNICOS
F4_TOP, F4_BOT = 0.360, 0.190
draw_floor(F4_TOP, F4_BOT, SHADE_4, "4",
           "DESFECHOS CLÍNICOS",
           "As quatro doenças crônicas do Capítulo 2")

# 4 caixas das doenças
F4_BOXES_Y = 0.275
DISEASES = [
    "Doença\ncardiovascular",
    "Doença\nmetabólica",
    "Neuro-\ndegeneração",
    "Câncer",
]
for idx, dis in enumerate(DISEASES):
    x = LEFT_MARGIN + 0.04 + idx * 0.22
    fig.patches.append(FancyBboxPatch(
        (x, F4_BOXES_Y - 0.020), 0.20, 0.055,
        boxstyle="round,pad=0.002,rounding_size=0.005",
        facecolor=SHADE_INSIDE, edgecolor=INK, linewidth=0.5,
        transform=fig.transFigure, zorder=2
    ))
    fig.text(x + 0.100, F4_BOXES_Y + 0.008, dis,
             fontsize=9, color=INK, weight="bold",
             ha="center", va="center", linespacing=1.2)

# Texto bottom "+ envelhecimento biológico acelerado"
fig.text(0.5, F4_BOT + 0.018,
         "+ envelhecimento biológico acelerado",
         fontsize=10, color=INK, weight="bold", ha="center", style="italic")
fig.text(0.5, F4_BOT + 0.003,
         "(telômeros, epigenoma)",
         fontsize=8.5, color=INK, ha="center", style="italic")

# Footer
fig.text(LEFT_MARGIN, 0.140,
         "A ativação crônica do eixo hipotálamo–hipófise–adrenal converte estresse psicológico sustentado",
         fontsize=8, color=INK_SOFT, style="italic")
fig.text(LEFT_MARGIN, 0.125,
         "em consequências biológicas mensuráveis — e essas consequências são exatamente os mecanismos discutidos",
         fontsize=8, color=INK_SOFT, style="italic")
fig.text(LEFT_MARGIN, 0.110,
         "ao longo deste livro como drivers das doenças crônicas do envelhecimento. Por isso a Integração Corpo-Mente",
         fontsize=8, color=INK_SOFT, style="italic")
fig.text(LEFT_MARGIN, 0.095,
         "não é pilar complementar — é pilar constitutivo.",
         fontsize=8.5, color=INK, weight="bold", style="italic")

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap12_Fig01.pdf"
png_path = out_dir / "_preview_Cap12_Fig01.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
