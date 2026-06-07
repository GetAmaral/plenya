"""
Cap12 Fig02 (PT-BR, B&W vetorial) — Quando o pilar psicológico entra, a biologia responde.

Layout do original 1536×1024:
  Header (FIGURA 2 + título + 2 subtítulos)
  Seção BIOLÓGICO: PCR-us + Cortisol matinal
  Seção PSICOLÓGICO: PHQ-9 + GAD-7
  Quote callout no lado direito
  Legenda + footer

Markers em B&W:
  ANTES   = ○ círculo vazado
  DEPOIS  = ● círculo preenchido
  ZONA-META = banda cinza claro
  CORTE-ALERTA = linha vertical tracejada
"""
from pathlib import Path as _Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, FancyBboxPatch, Circle, Polygon

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
GRAY2    = "#D9D9D9"
ZONE_META = "#E8E8E8"

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
ax.add_patch(Rectangle((40, 28), 124, 38, facecolor=INK, edgecolor="none"))
ax.text(40 + 62, 28 + 19, "FIGURA 2",
        fontsize=11, color="white", weight="bold",
        ha="center", va="center")

ax.text(40, 100,
        "Quando o pilar psicológico entra, a biologia responde.",
        fontsize=22, color=INK, weight="bold", va="center", ha="left")

ax.text(40, 148,
        "Ana, 44 anos: 18 meses de otimização biológica sem mover dois marcadores.",
        fontsize=11, color=SOFT, va="center", ha="left")
ax.text(40, 170,
        "Seis meses de trabalho na mente e os quatro saíram do lugar.",
        fontsize=11, color=INK, weight="bold", va="center", ha="left")

# ============================================================
# QUOTE CALLOUT (lado direito)
# ============================================================
QUOTE_X = 1230
QUOTE_W = 290
QUOTE_TOP = 215
QUOTE_BOT = 830

ax.add_patch(FancyBboxPatch(
    (QUOTE_X, QUOTE_TOP), QUOTE_W, QUOTE_BOT - QUOTE_TOP,
    boxstyle="round,pad=4,rounding_size=12",
    facecolor=BG, edgecolor=GRAY1, linewidth=0.8, zorder=1))

QX = QUOTE_X + 20
QW = QUOTE_W - 40

# Aspas grandes no topo
ax.text(QX, QUOTE_TOP + 46, "“",
        fontsize=36, color=INK, weight="bold",
        ha="left", va="center")

# Parte 1 — declaração principal (bold)
p1_lines = ["Os marcadores", "biológicos não", "mudaram antes."]
y = QUOTE_TOP + 100
for ln in p1_lines:
    ax.text(QX, y, ln,
            fontsize=11, color=INK, weight="bold",
            ha="left", va="center")
    y += 24

# Divider
y += 8
ax.plot([QX, QX + QW], [y, y], color=GRAY1, linewidth=0.5)

# Parte 2 — método (mix normal + bold)
y += 22
p2_normal = ["Mudaram após", "6 meses de"]
for ln in p2_normal:
    ax.text(QX, y, ln,
            fontsize=10, color=INK, ha="left", va="center")
    y += 22
p2_bold = ["terapia cognitivo-", "comportamental (TCC)",
           "+ redução do estresse", "(MBSR) + antidepressivo", "em baixa dose."]
for ln in p2_bold:
    ax.text(QX, y, ln,
            fontsize=10, color=INK, weight="bold",
            ha="left", va="center")
    y += 22

# Divider
y += 6
ax.plot([QX, QX + QW], [y, y], color=GRAY1, linewidth=0.5)

# Parte 3 — resultado final
y += 22
ax.text(QX, y, "Idade epigenética", fontsize=10, color=INK, ha="left", va="center")
y += 22
ax.text(QX, y, "desacelerou", fontsize=10, color=INK, ha="left", va="center")
y += 22
ax.text(QX, y, "2 anos.", fontsize=10, color=INK, weight="bold", ha="left", va="center")

# ============================================================
# CHARTS (BIOLÓGICO + PSICOLÓGICO)
# ============================================================
NAME_X     = 40
NAME_W     = 350
BAR_LEFT   = 390
BAR_RIGHT  = 1050
META_X     = 1140
ROW_HEIGHT = 60

SECTIONS = [
    {
        "title": "BIOLÓGICO",
        "y_top": 230,
        "rows": [
            {"name": ["Proteína C reativa (PCR-us)"], "unit": "marcador de inflamação · mg/L",
             "antes": "1,8", "depois": "0,7", "meta": "< 1,0",
             "ticks": [0, 1, 2, 3, 4],
             "zone_side": "right", "alert": None,
             "row_y": 320},
            {"name": ["Cortisol matinal"], "unit": "hormônio do estresse · µg/dL",
             "antes": "22", "depois": "14", "meta": "10–18",
             "ticks": [0, 10, 20, 30],
             "zone_side": "right", "alert": None,
             "row_y": 470},
        ],
    },
    {
        "title": "PSICOLÓGICO",
        "y_top": 570,
        "rows": [
            {"name": ["Escala de sintomas depressivos", "(PHQ-9)"], "unit": "quanto maior, pior",
             "antes": "6", "depois": "14", "meta": "≥ 10",
             "ticks": [0, 5, 10, 15, 20],
             "zone_side": "left", "alert": True,
             "row_y": 670},
            {"name": ["Escala de ansiedade", "(GAD-7)"], "unit": "quanto maior, pior",
             "antes": "5", "depois": "16", "meta": "≥ 10",
             "ticks": [0, 5, 10, 15, 20],
             "zone_side": "left", "alert": True,
             "row_y": 790},
        ],
    },
]

def value_to_x(val, vmin, vmax):
    return BAR_LEFT + (val - vmin) / (vmax - vmin) * (BAR_RIGHT - BAR_LEFT)

# Posições simbólicas para markers (mesma fração em todas as linhas)
ANTES_FRAC  = 0.30
DEPOIS_FRAC = 0.55
# Zona meta (BIOLÓGICO = direita; PSICOLÓGICO = esquerda)
ZONE_RIGHT_X0 = 0.45
ZONE_RIGHT_X1 = 1.00
ZONE_LEFT_X0  = 0.00
ZONE_LEFT_X1  = 0.45
# Corte-alerta no meio
ALERT_FRAC    = 0.45

for sec in SECTIONS:
    y_header = sec["y_top"]
    ax.text(NAME_X, y_header, sec["title"],
            fontsize=12, color=INK, weight="bold",
            ha="left", va="center")
    # Column headers
    header_label = "CORTE-ALERTA" if any(r["alert"] for r in sec["rows"]) else "META"
    # ANTES at left-of-bar-mid, DEPOIS at right-of-bar-mid
    ax.text((BAR_LEFT + BAR_RIGHT)/2 - 200, y_header, "ANTES",
            fontsize=10, color=SOFT, style="italic",
            ha="center", va="center")
    ax.text((BAR_LEFT + BAR_RIGHT)/2 + 100, y_header, "DEPOIS",
            fontsize=10, color=SOFT, style="italic",
            ha="center", va="center")
    ax.text(META_X, y_header, header_label,
            fontsize=10, color=SOFT, style="italic",
            ha="center", va="center")
    # Divider line under header
    ax.plot([NAME_X, META_X + 80], [y_header + 18, y_header + 18],
            color=GRAY1, linewidth=0.6)

    for row in sec["rows"]:
        y = row["row_y"]
        # Nome do marcador (esquerda) — line-height generoso
        NAME_LH = 22
        n_lines = len(row["name"])
        # Centralizar verticalmente o bloco de nome em torno do marker
        block_h = (n_lines - 1) * NAME_LH
        ny = y - 6 - block_h
        for ln in row["name"]:
            ax.text(NAME_X, ny, ln,
                    fontsize=11, color=INK, weight="bold",
                    ha="left", va="center")
            ny += NAME_LH
        ax.text(NAME_X, ny + 4, row["unit"],
                fontsize=9, color=SOFT, ha="left", va="center")

        # Eixo X (linha cinza com ticks pré-definidos)
        axis_y = y + 22
        ax.plot([BAR_LEFT, BAR_RIGHT], [axis_y, axis_y],
                color=GRAY1, linewidth=0.6)
        ticks = row["ticks"]
        tmin, tmax = ticks[0], ticks[-1]
        for tv in ticks:
            tx = BAR_LEFT + (tv - tmin) / (tmax - tmin) * (BAR_RIGHT - BAR_LEFT)
            ax.plot([tx, tx], [axis_y, axis_y + 4],
                    color=GRAY1, linewidth=0.5)
            ax.text(tx, axis_y + 14, f"{tv:g}",
                    fontsize=8.5, color=SOFT, ha="center", va="center")

        # Zona meta (banda cinza claro acima do eixo) — POSIÇÃO SIMBÓLICA
        bar_w = BAR_RIGHT - BAR_LEFT
        if row["zone_side"] == "right":
            zx0 = BAR_LEFT + bar_w * ZONE_RIGHT_X0
            zx1 = BAR_LEFT + bar_w * ZONE_RIGHT_X1
        else:
            zx0 = BAR_LEFT + bar_w * ZONE_LEFT_X0
            zx1 = BAR_LEFT + bar_w * ZONE_LEFT_X1
        ax.add_patch(Rectangle((zx0, axis_y - 24), zx1 - zx0, 24,
                                facecolor=ZONE_META, edgecolor="none", zorder=1))

        # Corte-alerta (vertical dashed line) — posição simbólica entre markers
        if row["alert"]:
            ax_x = BAR_LEFT + bar_w * ALERT_FRAC
            ax.plot([ax_x, ax_x], [axis_y - 28, axis_y - 2],
                    color=INK, linewidth=0.8, linestyle=(0, (4, 3)), zorder=3)

        # Posições simbólicas dos marcadores
        x_antes  = BAR_LEFT + bar_w * ANTES_FRAC
        x_depois = BAR_LEFT + bar_w * DEPOIS_FRAC

        marker_y = axis_y - 12

        # Seta de ANTES → DEPOIS (no nível do marker)
        sign = 1 if x_depois > x_antes else -1
        line_start = x_antes + sign * 9
        head_tip = x_depois - sign * 9
        head_base = head_tip - sign * 7
        ax.plot([line_start, head_base], [marker_y, marker_y],
                color=INK, linewidth=0.9, solid_capstyle="butt", zorder=4)
        head_tri = [(head_tip, marker_y),
                    (head_base, marker_y - 3),
                    (head_base, marker_y + 3)]
        ax.add_patch(Polygon(head_tri, closed=True, facecolor=INK,
                              edgecolor="none", zorder=4))

        # ANTES (hollow)
        ax.add_patch(Circle((x_antes, marker_y), 8,
                             facecolor=BG, edgecolor=INK,
                             linewidth=1.8, zorder=6))
        ax.text(x_antes, marker_y - 24, row["antes"],
                fontsize=11, color=INK, weight="bold",
                ha="center", va="center")

        # DEPOIS (filled)
        ax.add_patch(Circle((x_depois, marker_y), 8,
                             facecolor=INK, edgecolor="none", zorder=6))
        ax.text(x_depois, marker_y - 24, row["depois"],
                fontsize=11, color=INK, weight="bold",
                ha="center", va="center")

        # META text
        ax.text(META_X, marker_y - 6, row["meta"],
                fontsize=12, color=INK, weight="bold",
                ha="center", va="center")

# ============================================================
# LEGENDA INFERIOR
# ============================================================
LEG_Y = 870
LEG_ITEMS_X = [200, 480, 760, 1040]

# ANTES (hollow)
ax.add_patch(Circle((LEG_ITEMS_X[0], LEG_Y), 8,
                     facecolor=BG, edgecolor=INK, linewidth=1.8, zorder=4))
ax.text(LEG_ITEMS_X[0] + 18, LEG_Y - 4, "ANTES",
        fontsize=10, color=INK, weight="bold",
        ha="left", va="center")
ax.text(LEG_ITEMS_X[0] + 18, LEG_Y + 12, "(situação inicial)",
        fontsize=9, color=SOFT, ha="left", va="center")

# DEPOIS (filled)
ax.add_patch(Circle((LEG_ITEMS_X[1], LEG_Y), 8,
                     facecolor=INK, edgecolor="none", zorder=4))
ax.text(LEG_ITEMS_X[1] + 18, LEG_Y - 4, "DEPOIS",
        fontsize=10, color=INK, weight="bold",
        ha="left", va="center")
ax.text(LEG_ITEMS_X[1] + 18, LEG_Y + 12, "(após 6 meses)",
        fontsize=9, color=SOFT, ha="left", va="center")

# ZONA-META (swatch)
ax.add_patch(Rectangle((LEG_ITEMS_X[2] - 12, LEG_Y - 7), 24, 14,
                        facecolor=ZONE_META, edgecolor=GRAY1, linewidth=0.5, zorder=3))
ax.text(LEG_ITEMS_X[2] + 18, LEG_Y - 4, "ZONA-META",
        fontsize=10, color=INK, weight="bold",
        ha="left", va="center")
ax.text(LEG_ITEMS_X[2] + 18, LEG_Y + 12, "(faixa desejada)",
        fontsize=9, color=SOFT, ha="left", va="center")

# CORTE-ALERTA (dashed line)
ax.plot([LEG_ITEMS_X[3], LEG_ITEMS_X[3]],
        [LEG_Y - 9, LEG_Y + 9],
        color=INK, linewidth=1.0, linestyle=(0, (4, 3)), zorder=3)
ax.text(LEG_ITEMS_X[3] + 14, LEG_Y - 4, "CORTE-ALERTA",
        fontsize=10, color=INK, weight="bold",
        ha="left", va="center")
ax.text(LEG_ITEMS_X[3] + 14, LEG_Y + 12, "(risco aumentado)",
        fontsize=9, color=SOFT, ha="left", va="center")

# ============================================================
# RODAPÉ — FONTE
# ============================================================
ax.plot([40, W_IMG - 40], [925, 925], color=GRAY1, linewidth=0.5)
ax.text(40, 950,
        "Fonte: caso-tipo do Capítulo 12.  Escala de sintomas depressivos (PHQ-9) (Kroenke et al., JGIM 2001);",
        fontsize=9, color=FOOT, va="center", ha="left", style="italic")
ax.text(40, 972,
        "Escala de ansiedade (GAD-7) (Spitzer et al., Arch Intern Med 2006).",
        fontsize=9, color=FOOT, va="center", ha="left", style="italic")

# ============================================================
# EXPORT
# ============================================================
out_dir = _Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap12_Fig02.pdf"
png_path = out_dir / "_preview_Cap12_Fig02.png"
plt.savefig(pdf_path, facecolor=BG, bbox_inches="tight", pad_inches=0.0)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches="tight", pad_inches=0.0)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
