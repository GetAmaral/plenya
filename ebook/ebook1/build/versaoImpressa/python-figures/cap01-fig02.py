"""
Cap01 Fig02 (PT-BR, B&W vetorial) — Todos os números de Ricardo estavam 'normais'.
Nenhum estava ótimo.

Dot plot multi-linha com 5 biomarcadores. Cada linha mostra:
  - Alvo ótimo para longevidade   (▲)
  - Valor do Ricardo              (●)
  - Limite de normalidade (lab)   (linha tracejada)

Saída: PDF vetorial em build/versaoImpressa/figuras-bw/Cap01_Fig02.pdf
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False
rcParams["pdf.fonttype"] = 42
rcParams["ps.fonttype"] = 42

# ---------- paleta B&W ----------
BG          = "#FFFFFF"
INK         = "#000000"
INK_SOFT    = "#3A3A3A"
TICK        = "#555555"
FOOT        = "#666666"
BAND_OK     = "#F4F4F4"   # zona ótima
BAND_MID    = "#E0E0E0"   # zona subótima
BAND_BAD    = "#BFBFBF"   # zona "normal" laboratório (acima do limite ainda usado pelo lab)

# ---------- dados ----------
# Posições visuais estimadas a partir do original — o designer NÃO usou escala
# linear value/limit. As barras têm escalas próprias por métrica, com o limite
# do laboratório sempre ancorado em ~85% e os marcadores onde visualmente fazem
# sentido. Preservar essa disposição é parte da identidade visual da figura.
# (nome, unidade, label ótimo, label Ricardo, label limite, pos ótimo, pos Ricardo, pos limite, valor exibido à direita, comment)
biomarkers = [
    ("Insulina de jejum", "(µIU/mL)", "< 8",   "14",    "25",  0.15, 0.32, 0.85, "14",    None),
    ("ApoB",              "(mg/dL)",  "< 90",  "118",   "130", 0.38, 0.72, 0.85, "118",   None),
    ("hs-CRP",            "(mg/L)",   "< 1,0", "2,4",   "3,0", 0.13, 0.58, 0.85, "2,4",   None),
    ("Homocisteína",      "(µmol/L)", "< 10",  "13,8",  "15",  0.38, 0.70, 0.85, "13,8",  None),
    # Idade arterial: sem limite (a "norma" é a idade real do paciente)
    ("Idade arterial",    "(anos)",   "52 (idade real)", "68", None, 0.32, 0.78, 0.85, "68", "+16 anos acima da idade real"),
]

# ---------- figura ----------
fig = plt.figure(figsize=(11.0, 7.4))
fig.patch.set_facecolor(BG)

# Layout: linhas empilhadas, eixos invisíveis, tudo controlado manualmente.
# Coordenadas em fração da figura.
LEFT_MARGIN  = 0.03
RIGHT_MARGIN = 0.97
BAR_LEFT     = 0.18   # onde começa a barra (apertado contra o label, como no original)
BAR_RIGHT    = 0.81   # onde termina a barra (mesmo ponto para todas as linhas)
VALUE_X      = 0.90   # posição da coluna "Valor de Ricardo"

# ---------- título ----------
fig.text(LEFT_MARGIN, 0.945,
         "Todos os números de Ricardo estavam 'normais'. Nenhum estava ótimo.",
         fontsize=17, color=INK, weight="bold")

# ---------- legenda no topo (markers explanation) ----------
LEG_Y = 0.880
# marcadores: triângulo = Valor de Ricardo (como no original), círculo = Alvo ótimo
fig.text(LEFT_MARGIN,        LEG_Y, "▲  Valor de Ricardo",
         fontsize=9.5, color=INK, weight="bold", va="center")
fig.text(LEFT_MARGIN + 0.22, LEG_Y, "●  Alvo ótimo para longevidade",
         fontsize=9.5, color=INK_SOFT, va="center")
fig.text(LEFT_MARGIN + 0.50, LEG_Y, "|  Limite de normalidade (laboratório)",
         fontsize=9.5, color=INK_SOFT, va="center")

# ---------- coluna headers ----------
HEAD_Y = 0.830
fig.text((BAR_LEFT + BAR_RIGHT) / 2, HEAD_Y, "Escala de referência",
         fontsize=8.5, color=TICK, ha="center", style="italic")
fig.text(VALUE_X,                    HEAD_Y, "Valor de Ricardo",
         fontsize=8.5, color=TICK, ha="center", style="italic")

# ---------- linhas de dados ----------
ROW_TOP    = 0.785
ROW_BOTTOM = 0.255
ROW_SPACE  = (ROW_TOP - ROW_BOTTOM) / (len(biomarkers) - 1) if len(biomarkers) > 1 else 0
BAR_HEIGHT = 0.030

for i, (name, unit, otimo_lab, ricardo_lab, limit_lab,
        otimo_pos, ricardo_pos, limit_pos, ricardo_value, comment) in enumerate(biomarkers):
    y = ROW_TOP - i * ROW_SPACE
    bar_w = BAR_RIGHT - BAR_LEFT

    # --- nome do biomarcador (esquerda) ---
    fig.text(LEFT_MARGIN, y + 0.005, name,
             fontsize=11, color=INK, weight="bold", va="center")
    fig.text(LEFT_MARGIN, y - 0.020, unit,
             fontsize=8.5, color=TICK, va="center")

    # --- barra com 3 zonas (sólidas cinza, gradação ótimo→pior) ---
    # otimo zone:    BAR_LEFT          → BAR_LEFT + bar_w * otimo_pos
    # subotimo zone: ... otimo_pos     → BAR_LEFT + bar_w * limit_pos
    # acima limite:  ... limit_pos     → BAR_RIGHT
    x_otimo = BAR_LEFT + bar_w * otimo_pos
    x_limit = BAR_LEFT + bar_w * limit_pos

    # ax = fig in normalized coords; usar fig.add_axes seria mais limpo, mas
    # patches via fig.patches dá controle direto sem axes overhead.
    from matplotlib.patches import Rectangle
    fig.patches.extend([
        Rectangle((BAR_LEFT, y - BAR_HEIGHT/2), x_otimo - BAR_LEFT, BAR_HEIGHT,
                  facecolor=BAND_OK,  edgecolor="none", transform=fig.transFigure, zorder=1),
        Rectangle((x_otimo,  y - BAR_HEIGHT/2), x_limit - x_otimo, BAR_HEIGHT,
                  facecolor=BAND_MID, edgecolor="none", transform=fig.transFigure, zorder=1),
        Rectangle((x_limit,  y - BAR_HEIGHT/2), BAR_RIGHT - x_limit, BAR_HEIGHT,
                  facecolor=BAND_BAD, edgecolor="none", transform=fig.transFigure, zorder=1),
    ])

    # --- linha tracejada do limite (laboratório) ---
    if limit_lab is not None:
        fig.lines.extend([
            plt.Line2D([x_limit, x_limit], [y - BAR_HEIGHT*0.85, y + BAR_HEIGHT*0.85],
                       color=INK, linewidth=0.9, linestyle=(0, (3, 2)),
                       transform=fig.transFigure, zorder=4)
        ])
        # label do limite (abaixo da barra)
        fig.text(x_limit, y - BAR_HEIGHT*0.85 - 0.014, limit_lab,
                 fontsize=8, color=INK_SOFT, ha="center", va="top")

    # --- marker: alvo ótimo (CÍRCULO — como no original) ---
    x_otimo_marker = BAR_LEFT + bar_w * otimo_pos
    fig.text(x_otimo_marker, y, "●", fontsize=12, color=INK_SOFT,
             ha="center", va="center", zorder=5)
    fig.text(x_otimo_marker, y + BAR_HEIGHT*0.85 + 0.005, otimo_lab,
             fontsize=8, color=INK_SOFT, ha="center", va="bottom")

    # --- marker: valor Ricardo (TRIÂNGULO — como no original) ---
    x_ricardo_marker = BAR_LEFT + bar_w * ricardo_pos
    fig.text(x_ricardo_marker, y, "▲", fontsize=12, color=INK,
             ha="center", va="center", zorder=6)
    fig.text(x_ricardo_marker, y + BAR_HEIGHT*0.85 + 0.005, ricardo_lab,
             fontsize=8, color=INK, weight="bold", ha="center", va="bottom")

    # --- valor à direita ---
    fig.text(VALUE_X, y, ricardo_value,
             fontsize=13, color=INK, weight="bold", ha="center", va="center")
    if comment:
        fig.text(VALUE_X, y - 0.025, comment,
                 fontsize=7.5, color=INK_SOFT, style="italic", ha="center", va="center")

# ---------- separador antes da legenda inferior ----------
SEP_Y = ROW_BOTTOM - 0.045
fig.lines.extend([
    plt.Line2D([LEFT_MARGIN, RIGHT_MARGIN], [SEP_Y, SEP_Y],
               color="#CFCFCF", linewidth=0.5, transform=fig.transFigure, zorder=2)
])

# ---------- legenda inferior das 3 zonas (3 caixinhas com texto) ----------
LEG_Y2 = SEP_Y - 0.045
ZONE_W = (RIGHT_MARGIN - LEFT_MARGIN) / 3 - 0.01

zones = [
    ("Zona ótima",
     "Abaixo do alvo para longevidade",
     BAND_OK),
    ("Zona subótima",
     "Entre o ótimo e o limite do laboratório",
     BAND_MID),
    ("Zona \"normal\" pelo laboratório",
     "Dentro da referência — mas longe do ideal para longevidade",
     BAND_BAD),
]

from matplotlib.patches import Rectangle
for i, (label, desc, color) in enumerate(zones):
    x0 = LEFT_MARGIN + i * (ZONE_W + 0.015)
    # swatch (caixinha colorida representando o cinza da zona)
    fig.patches.append(
        Rectangle((x0, LEG_Y2 - 0.005), 0.025, 0.025,
                  facecolor=color, edgecolor="#888888", linewidth=0.5,
                  transform=fig.transFigure, zorder=3)
    )
    fig.text(x0 + 0.032, LEG_Y2 + 0.013, label,
             fontsize=8.5, color=INK, weight="bold", va="center")
    fig.text(x0 + 0.032, LEG_Y2 - 0.005, desc,
             fontsize=7.5, color=INK_SOFT, va="center")

# ---------- footer ----------
foot_lines = [
    "Triângulo (▲): valores reais de Ricardo. Círculo (●): alvo ótimo em estudos de longevidade e centenários.",
    "A linha tracejada marca o limite de \"normalidade\" usado pelo laboratório.",
]
for i, line in enumerate(foot_lines):
    fig.text(LEFT_MARGIN, 0.060 - i * 0.018, line,
             fontsize=7.5, color=FOOT)

# ---------- save ----------
out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap01_Fig02.pdf"
png_path = out_dir / "_preview_Cap01_Fig02.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
