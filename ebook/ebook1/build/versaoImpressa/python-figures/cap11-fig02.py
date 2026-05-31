"""
Cap11 Fig02 (PT-BR, B&W vetorial) — Oito genes. Oito decisões clínicas diferentes.
Grid 3 categorias × cards.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, FancyBboxPatch, Circle, Ellipse

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
CAT_1    = "#F0F0F0"   # nutrição/suplementação — mais claro
CAT_2    = "#E0E0E0"   # comportamento — médio
CAT_3    = "#C8C8C8"   # decisão estratégica — mais escuro
BAND_HEADER = "#3A3A3A"

fig = plt.figure(figsize=(11.0, 8.4))
fig.patch.set_facecolor(BG)

# Título
fig.text(0.025, 0.955,
         "Figura 2 — Oito genes. Oito decisões clínicas diferentes.",
         fontsize=15, color=INK, weight="bold")
fig.text(0.025, 0.920,
         "Quando o teste genético muda o plano — e quando apenas adiciona ansiedade.",
         fontsize=9.5, color=INK_SOFT, style="italic")

# 3 categorias (headers)
categories = [
    ("1.  SUPLEMENTAÇÃO\nE METABOLISMO",
     [
        ("1", "MTHFR",      "(C677T,\nA1298C)",
         "Suplementar com folato\nmetilado.",
         "Monitorar\nhomocisteína."),
        ("2", "FADS1/2",    "(rs174576)",
         "Conversão ALA →\nEPA/DHA reduzida.",
         "Linhaça não resolve."),
        ("3", "VDR FokI",   "(rs2228570)",
         "Menor responsividade\nà vitamina D.",
         "Dose maior\n(3.000–10.000 UI/dia em\nadultos comuns)\n40 ng/mL."),
     ]),
    ("2.  COMPORTAMENTO\nE ESTILO DE VIDA",
     [
        ("4", "CYP1A2",     "(rs762551)",
         "Metabolizador lento\nde cafeína.",
         "Lento (CC): risco\nhipertensão e IAM.\nRápido (AA):\nrisco menor."),
        ("5", "FTO",        "(rs9939609)",
         "Maior tendência\nao ganho.",
         "Exercício de\nmoderada–alta intensidade\nresponde acima\nda média.\n(Frayling 2007)"),
        ("6", "ALDH2",      "(rs671)",
         "Risco câncer com\nqualquer dose de álcool.",
         "Abstinência ou\nminimização, não\n\"moderação\".\n(Aging Cancer 2024)"),
     ]),
    ("3.  DECISÃO CLÍNICA\nESTRATÉGICA",
     [
        ("7", "APOE4",                 "",
         "Risco aumentado\nAlzheimer tardio.",
         "Prevenção multidomínio\n(FINGER).\nBenefício maior em\nrisco intermediário."),
        ("8", "ESR1 / COL1A1",         "",
         "Maior perda óssea\npós-menopausa.",
         "Pesa no cálculo\nde TRH em mulheres\nem conjunto com\nhistória familiar\nde osteoporose."),
     ]),
]

# Layout: 3 colunas, cards empilhados
COL_X = [0.04, 0.36, 0.68]
COL_W = 0.28
HEAD_Y = 0.875

for col, (head, cards) in enumerate(categories):
    x0 = COL_X[col]
    x1 = x0 + COL_W
    # Header da categoria
    fig.patches.append(Rectangle(
        (x0, HEAD_Y - 0.035), COL_W, 0.045,
        facecolor=BAND_HEADER, edgecolor="none",
        transform=fig.transFigure, zorder=2
    ))
    fig.text(x0 + 0.010, HEAD_Y - 0.012, head,
             fontsize=9.5, color="white", weight="bold", va="center",
             linespacing=1.15)

    # Cards
    n_cards = len(cards)
    CARDS_TOP = HEAD_Y - 0.060
    CARDS_BOTTOM = 0.150
    if n_cards == 2:
        CARD_H = 0.32
    else:
        CARD_H = 0.21
    CARD_SPACE = 0.010

    cat_color = [CAT_1, CAT_2, CAT_3][col]

    for ci, (num, name, variant, meaning, action) in enumerate(cards):
        y_top = CARDS_TOP - ci * (CARD_H + CARD_SPACE)
        y_bot = y_top - CARD_H

        # Card box
        fig.patches.append(FancyBboxPatch(
            (x0, y_bot), COL_W, CARD_H,
            boxstyle="round,pad=0.003,rounding_size=0.005",
            facecolor=cat_color, edgecolor=INK, linewidth=0.5,
            transform=fig.transFigure, zorder=1
        ))
        # Número (círculo escuro)
        ASPECT = 11.0 / 8.4
        fig.patches.append(Ellipse(
            (x0 + 0.020, y_top - 0.022),
            width=0.020, height=0.020 * ASPECT,
            facecolor=INK, edgecolor="none",
            transform=fig.transFigure, zorder=3
        ))
        fig.text(x0 + 0.020, y_top - 0.022, num,
                 fontsize=10, color="white", weight="bold",
                 ha="center", va="center", zorder=4)
        # Nome do gene
        fig.text(x0 + 0.040, y_top - 0.022, name,
                 fontsize=12, color=INK, weight="bold",
                 va="center", zorder=4)
        # Variante
        if variant:
            fig.text(x0 + 0.040, y_top - 0.045, variant,
                     fontsize=7.5, color=INK_SOFT, style="italic",
                     va="top", linespacing=1.15)
        # Significado
        fig.text(x0 + 0.010, y_top - 0.085, meaning,
                 fontsize=8, color=INK,
                 va="top", linespacing=1.25)
        # Linha separadora
        fig.lines.append(plt.Line2D(
            [x0 + 0.010, x1 - 0.010],
            [y_top - 0.130, y_top - 0.130],
            color="#888888", linewidth=0.4,
            transform=fig.transFigure
        ))
        # Ação
        fig.text(x0 + 0.010, y_top - 0.140, action,
                 fontsize=8, color=INK, weight="bold",
                 va="top", linespacing=1.25)

# Footer
fig.text(0.025, 0.105,
         "Categorias indicam quando o resultado é acionável:",
         fontsize=8, color=INK_SOFT, style="italic")
fig.text(0.025, 0.085,
         "1. Suplementação e metabolismo  ·  2. Comportamento e estilo de vida  ·  3. Decisão clínica estratégica.",
         fontsize=8, color=INK_SOFT)
fig.text(0.025, 0.055,
         "Fontes: variantes de impacto clínico estabelecido referenciadas em: MTHFR (Frosst 1995), APOE (Farrer, Nature Medicine 2024),",
         fontsize=7.5, color=FOOT)
fig.text(0.025, 0.038,
         "FTO (Frayling 2007), ALDH2 (Brooks 2009), CYP1A2 (Cornelis 2006). Síntese do Capítulo 11.",
         fontsize=7.5, color=FOOT)

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap11_Fig02.pdf"
png_path = out_dir / "_preview_Cap11_Fig02.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
