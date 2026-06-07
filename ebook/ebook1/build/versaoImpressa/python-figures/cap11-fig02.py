"""
Cap11 Fig02 (PT-BR, B&W vetorial) — Oito genes. Oito decisões clínicas diferentes.

Layout horizontal espelhando o original 1536×1024:
  - Header: FIGURA 2 em caixa preta + título + subtítulo
  - 3 grupos com arcos curvos:
      G1 SUPLEMENTAÇÃO E METABOLISMO: cols 1,2,3 (centros x=98, 268, 431)
      G2 COMPORTAMENTO E ESTILO DE VIDA: cols 4,5,6 (centros x=643, 829, 1007)
      G3 DECISÃO CLÍNICA ESTRATÉGICA: cols 7,8 (centros x=1216, 1407)
  - Cada coluna: badge numerado → oval c/ nome+polimorfismo → linha dotted →
    descrição → seta → recomendação (bold) → detalhe
  - Legenda 4 categorias na base + footer fontes

Categoria por gene (não por grupo): verde=ação simples, azul=comportamento,
âmbar=decisão complexa, vermelho=evitar. Em B&W diferenciamos por TOM de cinza:
  verde → cinza claro
  azul  → cinza médio
  âmbar → cinza escuro
  vermelho → preto
"""
from pathlib import Path as _Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import (
    Rectangle, FancyBboxPatch, Circle, Ellipse, Polygon, FancyArrowPatch, Wedge
)


def draw_cat_badge(ax, cx, cy, r, shape, num=None, fs=10):
    """Desenha badge categórico com forma distintiva em B&W.
    shape ∈ {filled, half, hollow, square}"""
    if shape == "filled":
        ax.add_patch(Circle((cx, cy), r, facecolor="#000000", edgecolor="white",
                             linewidth=1.5, zorder=5))
        if num:
            ax.text(cx, cy, num, fontsize=fs, color="white",
                    weight="bold", ha="center", va="center", zorder=6)
    elif shape == "half":
        # Círculo branco com borda + wedge esquerdo preto
        ax.add_patch(Circle((cx, cy), r, facecolor="white", edgecolor="#000000",
                             linewidth=1.4, zorder=5))
        ax.add_patch(Wedge((cx, cy), r, 90, 270, facecolor="#000000",
                            edgecolor="none", zorder=5))
        if num:
            # Numero centralizado — fica em cima da divisória, fica difícil ler.
            # Solução: número branco do lado preto (esquerda), preto do lado branco (direita)
            # Mais simples: número preto no canto-direito (lado branco)
            ax.text(cx + r*0.40, cy, num, fontsize=fs, color="#000000",
                    weight="bold", ha="center", va="center", zorder=6)
    elif shape == "hollow":
        ax.add_patch(Circle((cx, cy), r, facecolor="white", edgecolor="#000000",
                             linewidth=2.2, zorder=5))
        if num:
            ax.text(cx, cy, num, fontsize=fs, color="#000000",
                    weight="bold", ha="center", va="center", zorder=6)
    elif shape == "square":
        side = r * 1.8
        ax.add_patch(Rectangle((cx - side/2, cy - side/2), side, side,
                                facecolor="#000000", edgecolor="white",
                                linewidth=1.5, zorder=5))
        if num:
            ax.text(cx, cy, num, fontsize=fs, color="white",
                    weight="bold", ha="center", va="center", zorder=6)
from matplotlib.path import Path as MPath
from matplotlib.patches import PathPatch

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False
rcParams["pdf.fonttype"] = 42
rcParams["ps.fonttype"] = 42

BG       = "#FFFFFF"
INK      = "#000000"
SOFT     = "#3A3A3A"
FOOT     = "#6A6A6A"
GRAY1    = "#9E9E9E"

# Em B&W usamos FORMAS distintas (não tom de cinza) pra categoria:
#   verde    = "filled"    → ● círculo preenchido
#   azul     = "half"      → ◐ semicírculo (esquerda preta)
#   âmbar    = "hollow"    → ○ círculo vazado borda grossa
#   vermelho = "square"    → ■ quadrado preenchido
CAT_VERDE = "filled"
CAT_AZUL  = "half"
CAT_AMBAR = "hollow"
CAT_VERM  = "square"

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
# FIGURA 2 em caixa preta (como no original)
ax.add_patch(Rectangle((40, 28), 124, 38, facecolor=INK, edgecolor="none"))
ax.text(40 + 62, 28 + 19, "FIGURA 2",
        fontsize=11, color="white", weight="bold",
        ha="center", va="center")

# Título
ax.text(40, 100,
        "Oito genes. Oito decisões clínicas diferentes.",
        fontsize=23, color=INK, weight="bold", va="center", ha="left")

# Subtítulo
ax.text(40, 152,
        "Quando o teste genético muda o plano — e quando apenas adiciona ansiedade.",
        fontsize=11, color=SOFT, va="center", ha="left")

# ============================================================
# GRUPOS (headers + arcos)
# ============================================================
# Header text positions (centros dos grupos)
GROUPS = [
    {"label": ["1. SUPLEMENTAÇÃO", "E METABOLISMO"], "x_left": 30,   "x_right": 510,
     "col_centers": [98, 268, 431]},
    {"label": ["2. COMPORTAMENTO",  "E ESTILO DE VIDA"], "x_left": 555, "x_right": 1090,
     "col_centers": [643, 829, 1007]},
    {"label": ["3. DECISÃO CLÍNICA", "ESTRATÉGICA"], "x_left": 1140, "x_right": 1500,
     "col_centers": [1216, 1407]},
]
HEAD_Y = 202
ARC_Y_TOP = 268
ARC_Y_BOT = 288

for grp in GROUPS:
    cx = (grp["x_left"] + grp["x_right"]) / 2
    # Header text (2 linhas) — line height generoso
    for j, line in enumerate(grp["label"]):
        ax.text(cx, HEAD_Y - 12 + j * 30, line,
                fontsize=12, color=INK, weight="bold",
                ha="center", va="center")
    # Arc — bezier cúbico: parêntese-inverso suave
    x_l = grp["x_left"]; x_r = grp["x_right"]
    r = 14  # raio dos cantos
    y_top = ARC_Y_TOP
    y_bot = ARC_Y_BOT
    verts = [
        (x_l, y_top),                    # início vertical esq
        (x_l, y_bot - r),                # antes da curva
        (x_l, y_bot),                    # CP1 da curva esq
        (x_l + r, y_bot),                # fim da curva esq (entra na horizontal)
        (x_r - r, y_bot),                # horizontal até fim
        (x_r, y_bot),                    # CP1 da curva dir
        (x_r, y_bot),                    # CP2 (mesmo)
        (x_r, y_bot - r),                # fim da curva dir
        (x_r, y_top),                    # vertical até topo
    ]
    codes = [
        MPath.MOVETO, MPath.LINETO,
        MPath.CURVE4, MPath.CURVE4, MPath.CURVE4,
        MPath.CURVE4, MPath.CURVE4, MPath.CURVE4,
        MPath.LINETO,
    ]
    # Simpler: 3 segments + 2 quarter-arcs via curve3
    verts = [
        (x_l, y_top),
        (x_l, y_bot - r),
        (x_l, y_bot),
        (x_l + r, y_bot),
        (x_r - r, y_bot),
        (x_r, y_bot),
        (x_r, y_bot - r),
        (x_r, y_top),
    ]
    codes = [
        MPath.MOVETO, MPath.LINETO,
        MPath.CURVE3, MPath.CURVE3,
        MPath.LINETO,
        MPath.CURVE3, MPath.CURVE3,
        MPath.LINETO,
    ]
    path = MPath(verts, codes)
    ax.add_patch(PathPatch(path, facecolor="none", edgecolor=INK, linewidth=0.9))
    # Dot no centro do arco
    ax.add_patch(Circle((cx, ARC_Y_BOT), 6, facecolor=INK, edgecolor="none", zorder=4))

# ============================================================
# GENES (8 colunas)
# ============================================================
GENES = [
    {
        "x": 98, "num": "1", "cat": CAT_VERDE,
        "name": "MTHFR",
        "polym": ["C677T,", "A1298C"],
        "desc": ["Metaboliza", "folato mal."],
        "reco": ["Suplementar", "L-metilfolato +", "metilcobalamina."],
        "detail": ["Monitorar", "homocisteína."],
    },
    {
        "x": 268, "num": "2", "cat": CAT_VERDE,
        "name": "FADS1/2",
        "polym": [],
        "desc": ["Conversão", "ruim de", "ALA →", "EPA/DHA."],
        "reco": ["Peixe ou", "suplemento", "direto."],
        "detail": ["Linhaça não", "resolve."],
    },
    {
        "x": 431, "num": "3", "cat": CAT_AMBAR,
        "name": "VDR",
        "name2": "FokI",
        "polym": ["rs2228570"],
        "desc": ["Menor", "responsividade", "à vitamina D."],
        "reco": ["Dose maior", "(7.000–10.000 UI/dia)", "sob monitoramento", "até atingir", "40 ng/mL."],
        "detail": [],
    },
    {
        "x": 643, "num": "4", "cat": CAT_AZUL,
        "name": "CYP1A2",
        "polym": ["rs762551"],
        "desc": ["Metabolizador", "lento ou rápido", "da cafeína."],
        "reco": ["Lento (CC):", "parar cafeína", "antes das 12h."],
        "detail": ["Rápido (AA):", "tolera até 16h."],
    },
    {
        "x": 829, "num": "5", "cat": CAT_AZUL,
        "name": "FTO",
        "polym": ["rs9939609"],
        "desc": ["Maior apetite", "e preferência", "calórica."],
        "reco": ["Exercício de", "intensidade", "moderada-alta", "responde acima", "da média."],
        "detail": ["Torna treino", "inegociável."],
    },
    {
        "x": 1007, "num": "6", "cat": CAT_VERM,
        "name": "ALDH2",
        "polym": ["rs671"],
        "desc": ["Flush + risco", "de câncer", "esofágico com", "álcool."],
        "reco": ["Abstinência ou", "consumo muito", "ocasional, não", "“com moderação”."],
        "detail": [],
    },
    {
        "x": 1216, "num": "7", "cat": CAT_AMBAR,
        "name": "APOE4",
        "polym": ["ε4/ε4"],
        "desc": ["Risco aumentado", "de Alzheimer", "tardio."],
        "reco": ["Prevenção", "multidomínio", "intensiva (FINGER)."],
        "detail": ["Benefício maior", "em portadores."],
    },
    {
        "x": 1407, "num": "8", "cat": CAT_AMBAR,
        "name": "ESR1 /",
        "name2": "COL1A1",
        "polym": ["polimorfismos"],
        "desc": ["Maior perda", "óssea na", "pós-menopausa."],
        "reco": ["Pesar no cálculo", "da decisão sobre", "TRH em mulher", "com história", "familiar de", "osteoporose."],
        "detail": [],
    },
]

# Y positions (medidas do original)
BADGE_Y = 305
BADGE_R = 14
OVAL_CY = 370
OVAL_RX = 76
OVAL_RY = 62
DOTTED_TOP = 440
DOTTED_BOT = 482
DESC_Y = 510
ARROW_TOP = 610
ARROW_BOT = 650
RECO_Y = 680
DETAIL_GAP = 14   # extra gap antes do detalhe

RECO_INK = "#000000"   # cor única pro texto da recomendação (sem mais gradiente)

for g in GENES:
    cx = g["x"]
    cat_shape = g["cat"]

    # Oval (gene name container) — borda neutra cinza médio
    ax.add_patch(Ellipse((cx, OVAL_CY), 2*OVAL_RX, 2*OVAL_RY,
                          facecolor=BG, edgecolor=GRAY1,
                          linewidth=1.0, zorder=2))

    # Gene name (1 or 2 lines) + polimorfismo dentro do oval
    has_polym = bool(g["polym"])
    n_polym = len(g["polym"])
    if "name2" in g:
        # 2-line gene name
        ax.text(cx, OVAL_CY - 26, g["name"],
                fontsize=12, color=INK, weight="bold",
                ha="center", va="center", zorder=3)
        ax.text(cx, OVAL_CY - 8, g["name2"],
                fontsize=12, color=INK, weight="bold",
                ha="center", va="center", zorder=3)
        polym_y = OVAL_CY + 18
    else:
        if has_polym and n_polym >= 2:
            # 1-line name + 2-line polym (ex.: MTHFR)
            ax.text(cx, OVAL_CY - 22, g["name"],
                    fontsize=13, color=INK, weight="bold",
                    ha="center", va="center", zorder=3)
            polym_y = OVAL_CY + 2
        else:
            ax.text(cx, OVAL_CY - 14, g["name"],
                    fontsize=13, color=INK, weight="bold",
                    ha="center", va="center", zorder=3)
            polym_y = OVAL_CY + 12

    # Polymorphism inside oval
    for pi, pline in enumerate(g["polym"]):
        ax.text(cx, polym_y + pi * 14, pline,
                fontsize=8.5, color=SOFT, ha="center", va="center", zorder=3)

    # Badge — círculo preto preenchido com número branco (uniforme)
    ax.add_patch(Circle((cx, BADGE_Y), BADGE_R,
                         facecolor=INK, edgecolor=BG,
                         linewidth=1.5, zorder=5))
    ax.text(cx, BADGE_Y, g["num"],
            fontsize=10, color="white", weight="bold",
            ha="center", va="center", zorder=6)

    # Dotted line (vertical, between oval and description)
    n_dots = 5
    for di in range(n_dots):
        dot_y = DOTTED_TOP + di * (DOTTED_BOT - DOTTED_TOP) / (n_dots - 1)
        ax.add_patch(Circle((cx, dot_y), 1.8, facecolor=SOFT, edgecolor="none", zorder=2))

    # Description (centered)
    for di, line in enumerate(g["desc"]):
        ax.text(cx, DESC_Y + di * 17, line,
                fontsize=9, color=SOFT, ha="center", va="center")

    # Arrow down (between desc and reco) — preto neutro
    ax.plot([cx, cx], [ARROW_TOP, ARROW_BOT - 8],
            color=INK, linewidth=1.4, solid_capstyle="butt", zorder=3)
    head = [
        (cx, ARROW_BOT),
        (cx - 6, ARROW_BOT - 8),
        (cx + 6, ARROW_BOT - 8),
    ]
    ax.add_patch(Polygon(head, closed=True, facecolor=INK, edgecolor="none", zorder=3))

    # Recommendation (bold preto)
    for ri, line in enumerate(g["reco"]):
        ax.text(cx, RECO_Y + ri * 17, line,
                fontsize=9, color=INK, weight="bold",
                ha="center", va="center")

    # Detail (gray, regular)
    detail_y = RECO_Y + len(g["reco"]) * 17 + DETAIL_GAP
    for ti, line in enumerate(g["detail"]):
        ax.text(cx, detail_y + ti * 17, line,
                fontsize=9, color=SOFT, ha="center", va="center")

# ============================================================
# RODAPÉ — FONTE
# ============================================================
ax.plot([40, W_IMG - 40], [930, 930], color=GRAY1, linewidth=0.5)
ax.text(40, 955,
        "Fonte: variantes de impacto clínico estabelecido referenciadas em: MTHFR (Frosst 1995); APOE (Fortea, Nature Medicine 2024);",
        fontsize=8.5, color=FOOT, va="center", ha="left")
ax.text(40, 975,
        "FTO (Frayling 2007); ALDH2 (Brooks 2009); CYP1A2 (Cornelis 2006).   Síntese do Capítulo 11.",
        fontsize=8.5, color=FOOT, va="center", ha="left")

# ============================================================
# EXPORT
# ============================================================
out_dir = _Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap11_Fig02.pdf"
png_path = out_dir / "_preview_Cap11_Fig02.png"
plt.savefig(pdf_path, facecolor=BG, bbox_inches="tight", pad_inches=0.0)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches="tight", pad_inches=0.0)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
