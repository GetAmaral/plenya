"""
Cap08 Fig02 (PT-BR, B&W vetorial) — Finasterida: quando sim, quando não.

Fluxograma clínico convertido pra B&W. Mantém semântica via ícones, peso de borda
e shape (não cor). Layout vertical compacto pra caber 1024 px de altura sem overlap.
"""
from pathlib import Path as _Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import (
    Rectangle, FancyBboxPatch, Polygon, Circle, FancyArrowPatch, Ellipse
)
import numpy as np

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False
rcParams["pdf.fonttype"] = 42
rcParams["ps.fonttype"] = 42

BG    = "#FFFFFF"
INK   = "#000000"
SOFT  = "#3A3A3A"
FOOT  = "#6A6A6A"
DARK  = "#2A2A2A"
GRAY1 = "#9E9E9E"
GRAY2 = "#D9D9D9"
GRAY3 = "#F2F2F2"

W_IMG, H_IMG = 1536, 1024
_FIG_W = 10.0
_FIG_H = _FIG_W * H_IMG / W_IMG
fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

ax = fig.add_axes([0, 0, 1, 1])
ax.set_xlim(0, W_IMG)
ax.set_ylim(H_IMG, 0)
ax.set_aspect("equal")
ax.axis("off")

# ============================================================
# CABEÇALHO
# ============================================================
ax.add_patch(Rectangle((22, 18), 156-22, 45-18, facecolor=INK, edgecolor="none"))
ax.text((22+156)/2, (18+45)/2, "FIGURA 2",
        fontsize=9.5, color="white", weight="bold", va="center", ha="center")

ax.text(41, 100,
        "Finasterida: quando sim, quando não.",
        fontsize=22, color=INK, weight="bold", va="center", ha="left")

ax.text(43, 160,
        "O posicionamento clínico do autor — indicação regulatória vs. uso cosmético, à luz da síndrome pós-finasterida (PFS).",
        fontsize=10, color=SOFT, va="center", ha="left")

# ============================================================
# PILL PACIENTE
# ============================================================
pill_x, pill_y, pill_w, pill_h = 412, 212, 589, 78
pad = pill_h / 2 - 2
ax.add_patch(FancyBboxPatch(
    (pill_x + pad, pill_y + pad),
    pill_w - 2*pad, pill_h - 2*pad,
    boxstyle=f"round,pad={pad-1},rounding_size={pad-1}",
    facecolor=DARK, edgecolor="none"))

icx, icy = pill_x + 42, pill_y + pill_h/2
ax.add_patch(Circle((icx, icy - 10), 9, facecolor="white", edgecolor="none"))
ax.add_patch(Polygon([
    (icx-15, icy+18), (icx+15, icy+18), (icx+12, icy+3), (icx-12, icy+3)
], closed=True, facecolor="white", edgecolor="none"))

ax.text(pill_x + 80, pill_y + 28, "Paciente adulto com",
        fontsize=11, color="white", weight="bold", va="center", ha="left")
ax.text(pill_x + 80, pill_y + 55, "queda de cabelo e/ou sintomas prostáticos",
        fontsize=11, color="white", weight="bold", va="center", ha="left")

# Seta pill → diamond
ax.add_patch(FancyArrowPatch(
    (716, pill_y + pill_h), (716, 308),
    arrowstyle="-|>,head_length=8,head_width=7",
    color=INK, linewidth=1.5, zorder=3))

# ============================================================
# DIAMANTE DE DECISÃO — center (716, 370)
# ============================================================
DCX, DCY = 716, 370
DW, DH = 244, 124
diamond_pts = [(DCX, DCY-DH/2), (DCX+DW/2, DCY), (DCX, DCY+DH/2), (DCX-DW/2, DCY)]
ax.add_patch(Polygon(diamond_pts, closed=True,
                     facecolor=GRAY2, edgecolor=INK, linewidth=1.6, zorder=3))
ax.text(DCX, DCY-22, "Tem HPB",
        fontsize=11, color=INK, weight="bold", va="center", ha="center", zorder=4)
ax.text(DCX, DCY, "sintomática",
        fontsize=11, color=INK, weight="bold", va="center", ha="center", zorder=4)
ax.text(DCX, DCY+22, "(IPSS ≥ 8)?",
        fontsize=11, color=INK, weight="bold", va="center", ha="center", zorder=4)

# ============================================================
# CONECTORES SIM / NÃO
# ============================================================
# SIM: diamante esquerdo (594, 370) → horizontal pra (359, 370) → desce pra box1 top (359, 450)
# Label SIM ACIMA do segmento horizontal
ax.text(490, 350, "SIM",
        fontsize=16, color=INK, weight="bold", va="center", ha="center", zorder=4)
ax.plot([DCX-DW/2, 359], [DCY, DCY], color=INK, linewidth=1.4, zorder=2)
ax.plot([359, 359], [DCY, 445], color=INK, linewidth=1.4, zorder=2)
ax.add_patch(Polygon([(355, 445), (363, 445), (359, 453)],
                     closed=True, facecolor=INK, edgecolor="none", zorder=4))

# NÃO: diamante direito (838, 370) → horizontal (1108, 370) → desce pra red box top (1108, 450)
ax.text(940, 350, "NÃO",
        fontsize=16, color=INK, weight="bold", va="center", ha="center", zorder=4)
ax.plot([DCX+DW/2, 1108], [DCY, DCY], color=INK, linewidth=1.4, zorder=2)
ax.plot([1108, 1108], [DCY, 445], color=INK, linewidth=1.4, zorder=2)
ax.add_patch(Polygon([(1104, 445), (1112, 445), (1108, 453)],
                     closed=True, facecolor=INK, edgecolor="none", zorder=4))

# ============================================================
# CAIXA SIM 1 — Indicação clínica consistente
# Layout original: ícone top-left + título à direita do ícone + body abaixo,
# alinhado em coluna com o título (não cruza o ícone)
# ============================================================
gb1_x, gb1_y, gb1_w, gb1_h = 169, 453, 381, 175
ax.add_patch(FancyBboxPatch(
    (gb1_x+6, gb1_y+6), gb1_w-12, gb1_h-12,
    boxstyle="round,pad=4,rounding_size=10",
    facecolor=BG, edgecolor=INK, linewidth=1.4))

# Icon top-left
icx, icy = gb1_x + 40, gb1_y + 40
ax.add_patch(Circle((icx, icy), 20, facecolor=INK, edgecolor="none"))
ax.plot([icx-10, icx-2, icx+11], [icy, icy+8, icy-9],
        color="white", linewidth=3, solid_capstyle="round", zorder=5)

# Título à direita do ícone (mesmo y do ícone)
TEXT_X = gb1_x + 75
ax.text(TEXT_X, gb1_y + 40, "Indicação clínica consistente",
        fontsize=9.5, color=INK, weight="bold", va="center", ha="left")
# Body alinhado no mesmo x do título
ax.text(TEXT_X, gb1_y + 80, "Finasterida 5 mg/dia OU",
        fontsize=10, color=SOFT, va="center", ha="left")
ax.text(TEXT_X, gb1_y + 105, "dutasterida 0,5 mg/dia —",
        fontsize=10, color=SOFT, va="center", ha="left")
ax.text(TEXT_X, gb1_y + 130, "décadas de evidência.",
        fontsize=10, color=SOFT, va="center", ha="left")

# Conector box1 → box2
ax.plot([359, 359], [gb1_y + gb1_h, 648], color=INK, linewidth=1.2, zorder=2)
ax.add_patch(Circle((359, 648), 4, facecolor=INK, edgecolor="none", zorder=3))

# ============================================================
# CAIXA SIM 2 — Consentimento informado
# Layout: ícone vertical-centro + TODOS os textos à direita do ícone
# ============================================================
gb2_x, gb2_y, gb2_w, gb2_h = 170, 648, 379, 135
ax.add_patch(FancyBboxPatch(
    (gb2_x+6, gb2_y+6), gb2_w-12, gb2_h-12,
    boxstyle="round,pad=4,rounding_size=10",
    facecolor=BG, edgecolor=INK, linewidth=1.4))

# Icon vertical center
icx, icy = gb2_x + 40, gb2_y + gb2_h/2
ax.add_patch(Circle((icx, icy), 20, facecolor=INK, edgecolor="none"))
ax.add_patch(Circle((icx, icy-6), 5, facecolor="white", edgecolor="none"))
ax.add_patch(Polygon([
    (icx-9, icy+10), (icx+9, icy+10), (icx+6, icy+1), (icx-6, icy+1)
], closed=True, facecolor="white", edgecolor="none"))

# Todos os 3 textos à direita do ícone, alinhados em coluna
TEXT_X = gb2_x + 75
ax.text(TEXT_X, gb2_y + 38, "Consentimento informado",
        fontsize=9.5, color=INK, weight="bold", va="center", ha="left")
ax.text(TEXT_X, gb2_y + 70, "inclui conversa sobre",
        fontsize=9.5, color=SOFT, va="center", ha="left")
ax.text(TEXT_X, gb2_y + 92, "efeitos sexuais possíveis.",
        fontsize=9.5, color=SOFT, va="center", ha="left")

# ============================================================
# CAIXA NÃO — Indicação cosmética (borda preta grossa)
# ============================================================
rb_x, rb_y, rb_w, rb_h = 882, 453, 453, 200
ax.add_patch(FancyBboxPatch(
    (rb_x+6, rb_y+6), rb_w-12, rb_h-12,
    boxstyle="round,pad=4,rounding_size=10",
    facecolor=BG, edgecolor=INK, linewidth=2.2))

# Icon top-left
picx, picy = rb_x + 40, rb_y + 40
ax.add_patch(Circle((picx, picy), 20, facecolor=BG, edgecolor=INK, linewidth=3.0))
slash = 15
ax.plot([picx - slash*0.707, picx + slash*0.707],
        [picy + slash*0.707, picy - slash*0.707],
        color=INK, linewidth=3.0, solid_capstyle="round")

# TODOS os textos alinhados à direita do ícone, em coluna
TEXT_X = rb_x + 75
ax.text(TEXT_X, rb_y + 30, "Indicação apenas cosmética",
        fontsize=10, color=INK, weight="bold", va="center", ha="left")
ax.text(TEXT_X, rb_y + 53, "(alopecia androgenética)",
        fontsize=10, color=INK, weight="bold", va="center", ha="left")
ax.text(TEXT_X, rb_y + 90, "NÃO prescrever finasterida",
        fontsize=10, color=INK, weight="bold", va="center", ha="left")
ax.text(TEXT_X, rb_y + 112, "nem dutasterida.",
        fontsize=10, color=INK, weight="bold", va="center", ha="left")
ax.text(TEXT_X, rb_y + 145, "Risco de síndrome pós-finasterida",
        fontsize=9, color=SOFT, va="center", ha="left")
ax.text(TEXT_X, rb_y + 165, "(PFS) desproporcional ao desfecho",
        fontsize=9, color=SOFT, va="center", ha="left")
ax.text(TEXT_X, rb_y + 185, "estético.",
        fontsize=9, color=SOFT, va="center", ha="left")

# Conector red → leaf
ax.plot([1108, 1108], [rb_y + rb_h, 670], color=INK, linewidth=1.2, zorder=2)
ax.add_patch(Circle((1108, 670), 4, facecolor=INK, edgecolor="none", zorder=3))

# ============================================================
# CAIXA NÃO 2 — Alternativas com evidência (leaf)
# ============================================================
glb_x, glb_y, glb_w, glb_h = 870, 670, 465, 132
ax.add_patch(FancyBboxPatch(
    (glb_x+6, glb_y+6), glb_w-12, glb_h-12,
    boxstyle="round,pad=4,rounding_size=10",
    facecolor=BG, edgecolor=INK, linewidth=1.4))

# Icon vertical center
lcx, lcy = glb_x + 40, glb_y + glb_h/2
ax.add_patch(Circle((lcx, lcy), 20, facecolor=INK, edgecolor="none"))
ax.add_patch(Ellipse((lcx, lcy), 24, 11, angle=-35,
                     facecolor="white", edgecolor="none"))
ax.plot([lcx-8, lcx+10], [lcy+7, lcy-7],
        color=INK, linewidth=1.4, solid_capstyle="round", zorder=6)

# TODOS os textos à direita do ícone, em coluna
TEXT_X = glb_x + 75
ax.text(TEXT_X, glb_y + 22, "Alternativas com evidência:",
        fontsize=10, color=INK, weight="bold", va="center", ha="left")
ax.text(TEXT_X, glb_y + 50, "minoxidil tópico 5% ou oral baixa dose;",
        fontsize=8.5, color=SOFT, va="center", ha="left")
ax.text(TEXT_X, glb_y + 70, "PRP capilar; microneedling; correção",
        fontsize=8.5, color=SOFT, va="center", ha="left")
ax.text(TEXT_X, glb_y + 90, "de ferritina ≥ 40; otimização de vit. D,",
        fontsize=8.5, color=SOFT, va="center", ha="left")
ax.text(TEXT_X, glb_y + 110, "B12, zinco, função tireoidiana.",
        fontsize=8.5, color=SOFT, va="center", ha="left")

# ============================================================
# BASE REGULATÓRIA + TIMELINE
# ============================================================
ax.plot([168, 1110], [820, 820], color=GRAY1, linewidth=1.0, zorder=1)
ax.text(1329, 820, "BASE REGULATÓRIA",
        fontsize=11, color=INK, weight="bold", va="center", ha="right",
        bbox=dict(facecolor=BG, edgecolor="none", pad=4))

NODES = [
    (288, ["2011 — FDA:", "advertência sobre", "depressão em bula."]),
    (715, ["2022 — FDA:", "atualização da bula", "alertando sobre",
           "ideação suicida."]),
    (1149, ["2025 — EMA:", "reconhecimento formal da",
            "síndrome pós-finasterida",
            "(disfunção sexual persistente +",
            "sintomas neuropsiquiátricos)."]),
]
# Note: 2022 has 3 lines body, 2025 has 4 lines body. Original confirmed via crop.

ax.plot([NODES[0][0], NODES[-1][0]], [850, 850],
        color=GRAY1, linewidth=1.0, zorder=1)
for cx, _ in NODES:
    ax.add_patch(Circle((cx, 850), 6.5, facecolor=INK, edgecolor="none", zorder=2))
    ax.plot([cx, cx], [850, 862], color=INK, linewidth=1.0, zorder=2)

BOXES = [
    (163, 862, 250, 88),
    (587, 862, 257, 88),
    (984, 862, 330, 88),
]
for (bx, by, bw, bh), (cx, lines) in zip(BOXES, NODES):
    ax.add_patch(FancyBboxPatch(
        (bx+4, by+4), bw-8, bh-8,
        boxstyle="round,pad=3,rounding_size=8",
        facecolor=BG, edgecolor=GRAY1, linewidth=1.0))
    icx, icy = bx + 26, by + bh/2
    ax.add_patch(FancyBboxPatch(
        (icx-16, icy-16), 32, 32,
        boxstyle="round,pad=0,rounding_size=4",
        facecolor=GRAY3, edgecolor=GRAY1, linewidth=1.0))
    ax.add_patch(Rectangle((icx-16, icy-16), 32, 7,
                           facecolor=GRAY1, edgecolor="none"))
    ax.plot([icx-9, icx-9], [icy-19, icy-12], color=INK, linewidth=1.6)
    ax.plot([icx+9, icx+9], [icy-19, icy-12], color=INK, linewidth=1.6)
    for dx in (-7, 0, 7):
        for dy in (-2, 5):
            ax.add_patch(Circle((icx+dx, icy+dy), 1.1,
                                facecolor=GRAY1, edgecolor="none"))

    ax.text(bx + 56, by + 16, lines[0],
            fontsize=9.5, color=INK, weight="bold", va="center", ha="left")
    for i, ln in enumerate(lines[1:]):
        ax.text(bx + 56, by + 32 + i*12, ln,
                fontsize=7.5, color=SOFT, va="center", ha="left")

# ============================================================
# SOURCE
# ============================================================
# Source — uso matplotlib mathtext pra misturar bold + italic na mesma linha
ax.text(43, 975,
        r"$\bf{Fonte:}$ $\it{FDA\ Drug\ Safety\ Communications}$ (2011, 2022); "
        r"$\it{EMA\ PRAC\ Recommendation}$ (2025).",
        fontsize=9.5, color=FOOT, va="center", ha="left")
ax.text(43, 1000,
        "Posicionamento clínico do autor expresso no Capítulo 8.",
        fontsize=9.5, color=FOOT, va="center", ha="left")

# ============================================================
# EXPORT
# ============================================================
out_dir = _Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap08_Fig02.pdf"
png_path = out_dir / "_preview_Cap08_Fig02.png"
plt.savefig(pdf_path, facecolor=BG, bbox_inches="tight", pad_inches=0.0)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches="tight", pad_inches=0.0)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
