"""
Cap12 Fig03 (PT-BR, B&W vetorial) — Cinco instrumentos que mudam a consulta em cinco minutos.
Layout portrait (fullpage). 5 escalas psicológicas + callout direito.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, FancyBboxPatch, Ellipse

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
BAND_OK  = "#F4F4F4"
BAND_MID = "#E0E0E0"
BAND_BAD = "#BFBFBF"
BAND_WARN = "#A8A8A8"
BAND     = "#EDEDED"

_FIG_W, _FIG_H = 8.4, 11.4
_ASPECT = _FIG_W / _FIG_H

fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

LEFT = 0.025

# Título
fig.text(LEFT, 0.965, "Figura 3",
         fontsize=8.5, color="white", weight="bold")
# Badge
fig.patches.append(Rectangle(
    (LEFT, 0.957), 0.060, 0.020,
    facecolor="#3A3A3A", edgecolor="none",
    transform=fig.transFigure, zorder=1
))
fig.text(LEFT + 0.030, 0.967, "FIGURA 3",
         fontsize=8, color="white", weight="bold",
         ha="center", va="center", zorder=2)

fig.text(LEFT, 0.935, "Cinco instrumentos que",
         fontsize=15, color=INK, weight="bold")
fig.text(LEFT, 0.915, "mudam a consulta em",
         fontsize=15, color=INK, weight="bold")
fig.text(LEFT, 0.895, "cinco minutos.",
         fontsize=15, color=INK, weight="bold")
fig.text(LEFT, 0.870,
         "O mínimo necessário para identificar quando",
         fontsize=8.5, color=INK_SOFT)
fig.text(LEFT, 0.855,
         "o plano deixa de ser biológico.",
         fontsize=8.5, color=INK_SOFT)

# ---------- 5 escalas (lado esquerdo) ----------
# Layout: cada escala tem nome+desc à esquerda, barra com ticks e bandas de gravidade
# (Right column for callout)

INSTRUMENTS = [
    ("1", "PHQ-9", "Depressão",
     "Sintomas depressivos\nnas últimas 2 semanas.",
     [0, 5, 10, 15, 20],
     [(0, 5, "leve"), (5, 10, "moderada"), (10, 15, "moderada–severa"), (15, 20, "severa")],
     "≥ 10 entra no plano agora.",
     "Qualquer resposta positiva sobre ideação suicida = conversa imediata, antes do paciente sair."),
    ("2", "GAD-7", "Ansiedade generalizada",
     "Sintomas de ansiedade\nnas últimas 2 semanas.",
     [0, 5, 10, 15, 21],
     [(0, 5, "leve"), (5, 10, "moderada"), (10, 21, "severa")],
     "≥ 10 ansiedade relevante.",
     None),
    ("3", "AUDIT", "Consumo problemático de álcool",
     "Padrões de uso\nde álcool.",
     [0, 8, 16, 40],
     [(0, 8, "risco"), (8, 16, "nocivo / dependência"), (16, 40, "")],
     "≥ 8 risco em uso social.",
     "Sobretudo em paciente que diz \"só bebo socialmente\"."),
    ("4", "PCL-5", "Transtorno de estresse pós-traumático",
     "Sintomas relacionados\na evento traumático.",
     [0, 33, 80],
     [(0, 33, ""), (33, 80, "")],
     "≥ 33 investigar trauma.",
     "Aplicar se a pergunta de trauma na anamnese for positiva."),
    ("5", "UCLA-3", "Solidão percebida",
     "Percepção subjetiva\nde conexão e ligação.",
     [0, 3, 6, 9],
     [(0, 3, "baixa"), (3, 6, ""), (6, 9, "elevada")],
     "≥ 6 solidão relevante.",
     "Sobretudo em paciente que diz \"muitos amigos\"."),
]

# Coluna esquerda (instrumentos)
LEFT_COL_W = 0.62   # use only left 62% — right 38% pra callout

# Y posições: 5 escalas distribuídas
ROW_TOP    = 0.820
ROW_BOTTOM = 0.110
n = len(INSTRUMENTS)
ROW_SPACE  = (ROW_TOP - ROW_BOTTOM) / n
BAR_INNER_H = 0.018

def small_circle(x, y, fill, edge=INK, r=0.011):
    fig.patches.append(Ellipse(
        (x, y), width=r*2, height=r*2*_ASPECT,
        facecolor=fill, edgecolor=edge, linewidth=1.0,
        transform=fig.transFigure, zorder=4
    ))

for i, (num, name, descr, sub_desc, ticks, ranges, corte, warn) in enumerate(INSTRUMENTS):
    y_top = ROW_TOP - i * ROW_SPACE
    y_center = y_top - 0.025
    bar_y = y_center - 0.015

    # Número (círculo)
    small_circle(LEFT + 0.013, y_center + 0.015, INK, INK)
    fig.text(LEFT + 0.013, y_center + 0.015, num,
             fontsize=9, color="white", weight="bold",
             ha="center", va="center", zorder=5)

    # Nome + descrição (com mais espaço entre o nome e a descrição)
    fig.text(LEFT + 0.035, y_center + 0.025, name,
             fontsize=11, color=INK, weight="bold", va="center")
    fig.text(LEFT + 0.110, y_center + 0.025, descr,
             fontsize=9, color=INK, va="center")
    fig.text(LEFT + 0.035, y_center + 0.005, sub_desc,
             fontsize=7.5, color=INK_SOFT, va="center",
             linespacing=1.2, style="italic")

    # ---- Escala ----
    SCALE_LEFT = LEFT + 0.030
    SCALE_RIGHT = LEFT_COL_W - 0.005
    scale_w = SCALE_RIGHT - SCALE_LEFT
    val_min, val_max = ticks[0], ticks[-1]

    def val_to_x(v):
        return SCALE_LEFT + (v - val_min) / (val_max - val_min) * scale_w

    # Bandas de gravidade (tons crescentes)
    n_ranges = len(ranges)
    band_tones = [BAND_OK, BAND_MID, BAND_BAD, BAND_WARN]
    for ri, (vstart, vend, label) in enumerate(ranges):
        color = band_tones[min(ri, len(band_tones) - 1)]
        x1 = val_to_x(vstart)
        x2 = val_to_x(vend)
        fig.patches.append(Rectangle(
            (x1, bar_y - BAR_INNER_H/2), x2 - x1, BAR_INNER_H,
            facecolor=color, edgecolor="none",
            transform=fig.transFigure, zorder=1
        ))
        # Label da faixa
        if label:
            fig.text((x1 + x2) / 2, bar_y - 0.014, label,
                     fontsize=7, color=INK_SOFT, ha="center", va="top")

    # Ticks
    for v in ticks:
        x = val_to_x(v)
        fig.lines.append(plt.Line2D(
            [x, x], [bar_y + BAR_INNER_H/2, bar_y + BAR_INNER_H/2 + 0.005],
            color=INK, linewidth=0.6, transform=fig.transFigure
        ))
        fig.text(x, bar_y + BAR_INNER_H/2 + 0.009, str(v),
                 fontsize=7.5, color=INK_SOFT, ha="center", va="bottom")

    # Corte (texto abaixo da escala)
    fig.text(SCALE_LEFT, bar_y - 0.035, "▶ " + corte,
             fontsize=9, color=INK, weight="bold", va="top")
    if warn:
        fig.text(SCALE_LEFT, bar_y - 0.053, warn,
                 fontsize=7, color=INK_SOFT, va="top",
                 linespacing=1.2, style="italic")

# ---------- callout direito ----------
CO_X1, CO_X2 = LEFT_COL_W + 0.015, 0.97
CO_Y1, CO_Y2 = 0.150, 0.820

fig.patches.append(FancyBboxPatch(
    (CO_X1, CO_Y1), CO_X2 - CO_X1, CO_Y2 - CO_Y1,
    boxstyle="round,pad=0.005,rounding_size=0.008",
    facecolor=BAND, edgecolor=INK, linewidth=0.8,
    transform=fig.transFigure, zorder=1
))

# Header com estrela
fig.text((CO_X1 + CO_X2) / 2, CO_Y2 - 0.020, "★",
         fontsize=18, color=INK, weight="bold", ha="center")

fig.text((CO_X1 + CO_X2) / 2, CO_Y2 - 0.060,
         "Quando o",
         fontsize=11, color=INK, weight="bold", ha="center")
fig.text((CO_X1 + CO_X2) / 2, CO_Y2 - 0.082,
         "psicológico",
         fontsize=11, color=INK, weight="bold", ha="center")
fig.text((CO_X1 + CO_X2) / 2, CO_Y2 - 0.104,
         "vira prioridade",
         fontsize=11, color=INK, weight="bold", ha="center")
fig.text((CO_X1 + CO_X2) / 2, CO_Y2 - 0.126,
         "do plano:",
         fontsize=11, color=INK, weight="bold", ha="center")

# 3 pontos
points = [
    ("1", "PCR persistente\n> 1,5 apesar de\npilares biológicos\notimizados."),
    ("2", "Pontuação\n≥ ponto de corte\nem qualquer\nescala acima."),
    ("3", "Ideação suicida\nou trauma\nrelatados em\nqualquer\nintensidade."),
]
co_y = CO_Y2 - 0.170
for num, text in points:
    # Bolinha numerada
    small_circle(CO_X1 + 0.018, co_y, INK, INK, r=0.011)
    fig.text(CO_X1 + 0.018, co_y, num,
             fontsize=9, color="white", weight="bold",
             ha="center", va="center", zorder=5)
    fig.text(CO_X1 + 0.040, co_y + 0.005, text,
             fontsize=8, color=INK, va="top",
             linespacing=1.25, weight="bold")
    co_y -= 0.085

# Caixa final no callout
fig.patches.append(Rectangle(
    (CO_X1 + 0.010, CO_Y1 + 0.010), CO_X2 - CO_X1 - 0.020, 0.060,
    facecolor="none", edgecolor=INK, linewidth=0.8,
    transform=fig.transFigure, zorder=2
))
fig.text((CO_X1 + CO_X2) / 2, CO_Y1 + 0.050,
         "Identificar cedo muda",
         fontsize=8, color=INK, weight="bold", ha="center")
fig.text((CO_X1 + CO_X2) / 2, CO_Y1 + 0.035,
         "a direção do plano",
         fontsize=8, color=INK, weight="bold", ha="center")
fig.text((CO_X1 + CO_X2) / 2, CO_Y1 + 0.020,
         "e o desfecho do paciente.",
         fontsize=8, color=INK, weight="bold", ha="center")

# Footer
fig.text(LEFT, 0.080,
         "Fontes: PHQ-9 (Kroenke et al., JGM 2001); GAD-7 (Spitzer et al., Arch Intern Med 2006);",
         fontsize=7, color=FOOT)
fig.text(LEFT, 0.065,
         "AUDIT (Saunders et al., Addiction 1993); PCL-5 (Weathers et al., NCPTSD 2013);",
         fontsize=7, color=FOOT)
fig.text(LEFT, 0.050,
         "UCLA-3 (Hughes et al., Research on Aging 2004). Todos validados em português brasileiro.",
         fontsize=7, color=FOOT)

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap12_Fig03.pdf"
png_path = out_dir / "_preview_Cap12_Fig03.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
