"""
Cap04 Fig01 (PT-BR, B&W vetorial) — O Caso Fernanda: Todos 'Normais'. Nenhum Ótimo.

Dot plot multi-linha com 7 biomarcadores. Mesma estrutura visual da Cap01 Fig02.
Inclui Vitamina D que é INVERTIDA (maior = melhor).
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle

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

# (nome, unidade, label ótimo, label Fernanda, label limite,
#  pos ótimo, pos Fernanda, pos limite, valor exibido)
# Vitamina D é INVERTIDA — ótimo é à direita, Fernanda à esquerda
biomarkers = [
    ("Insulina de jejum", "(µIU/mL)", "< 8",     "13",   "25",  0.15, 0.32, 0.85, "13 µIU/mL",   False),
    ("HOMA-IR",           "(—)",      "< 1,5",   "3,0",  "sem limite\nde lab definido", 0.20, 0.50, None, "3,0",  False),
    ("ApoB",              "(mg/dL)",  "< 90",    "108",  "< 130", 0.38, 0.65, 0.85, "108 mg/dL", False),
    ("hs-CRP",            "(mg/L)",   "< 1,0",   "1,9",  "< 3,0", 0.13, 0.48, 0.85, "1,9 mg/L",  False),
    ("Homocisteína",      "(µmol/L)", "< 10",    "12,4", "< 15",  0.38, 0.68, 0.85, "12,4 µmol/L", False),
    # Vitamina D — INVERTIDA
    ("Vitamina D (25-OH)", "(ng/mL)", "40-60",   "22",   "> 20",  0.78, 0.30, 0.18, "22 ng/mL",   True),
    ("TG/HDL*",           "(—)",      "< 2,0",   "3,9",  "< 3,5\n(limite de risco)", 0.18, 0.88, 0.78, "3,9\n(calculado do\nlipidograma existente)", False),
]

# Aspecto 1.500, igual ao da arte original (1536×1024). Era 1.279, o que
# esticava a figura verticalmente e sobrava faixa branca no rodapé.
fig = plt.figure(figsize=(11.0, 7.333))
fig.patch.set_facecolor(BG)

LEFT_MARGIN  = 0.03
RIGHT_MARGIN = 0.97
# Geometria MEDIDA na arte original: a barra vai de x 0.182 a 0.835 e as três
# zonas têm fronteiras FIXAS em todas as linhas (0.383 e 0.684) — o gerador
# calculava a largura de cada zona por linha, o que fazia cada barra ter uma
# escala diferente. Os triângulos de alvo ficam todos alinhados em x 0.302 e
# só o círculo do valor varia. Não há linha invertida: a Vitamina D segue o
# mesmo arranjo das demais.
BAR_LEFT     = 0.182
BAR_RIGHT    = 0.835
ZONA_OTIMA_X = 0.383
ZONA_LAB_X   = 0.684
TRI_X        = 0.302
CIRC_X       = [0.549, 0.526, 0.606, 0.545, 0.587, 0.684, 0.775]
LINHA_Y      = [0.246, 0.341, 0.433, 0.520, 0.605, 0.687, 0.767]
VALUE_X      = 0.915

# Título
fig.text(LEFT_MARGIN, 0.96,
         "O Caso Fernanda: Todos 'Normais'. Nenhum Ótimo.",
         fontsize=17, color=INK, weight="bold")

# Legenda topo
LEG_Y = 0.910
fig.text(LEFT_MARGIN,        LEG_Y, "●  Valor de Fernanda",
         fontsize=9.5, color=INK, weight="bold", va="center")
fig.text(LEFT_MARGIN + 0.22, LEG_Y, "▲  Alvo ótimo para longevidade",
         fontsize=9.5, color=INK_SOFT, va="center")
fig.text(LEFT_MARGIN + 0.50, LEG_Y, "|  Limite de normalidade (laboratório)",
         fontsize=9.5, color=INK_SOFT, va="center")

# Headers das colunas
HEAD_Y = 0.870
fig.text((BAR_LEFT + BAR_RIGHT) / 2, HEAD_Y, "Escala de referência",
         fontsize=8.5, color=TICK, ha="center", style="italic")
fig.text(VALUE_X, HEAD_Y, "Valor de Fernanda",
         fontsize=8.5, color=TICK, ha="center", style="italic")

# Réguas horizontais: uma sob o cabeçalho e uma entre cada par de linhas.
# Posições MEDIDAS na arte original — a versão anterior não tinha nenhuma, o
# que deixava as sete linhas soltas no branco.
REGUAS_Y = [0.195, 0.287, 0.380, 0.468, 0.554, 0.639, 0.718]
for _ry in REGUAS_Y:
    fig.lines.append(plt.Line2D(
        [0.028, 0.965], [1.0 - _ry, 1.0 - _ry],
        color="#D6D6D6", linewidth=0.7, transform=fig.transFigure, zorder=0))

# Cabeçalho da tabela (a arte original nomeia as três colunas)
fig.text(0.045, 1.0 - 0.176, "Biomarcador",
         fontsize=9, color=TICK, va="center")

# Linhas de dados
ROW_TOP = 0.830
ROW_BOTTOM = 0.260
ROW_SPACE = (ROW_TOP - ROW_BOTTOM) / (len(biomarkers) - 1)
BAR_HEIGHT = 0.026

for i, (name, unit, otimo_lab, fer_lab, lim_lab,
        otimo_pos, fer_pos, lim_pos, fer_value, inverted) in enumerate(biomarkers):
    y = 1.0 - LINHA_Y[i]          # posições de linha medidas no original
    bar_w = BAR_RIGHT - BAR_LEFT

    # Nome
    fig.text(LEFT_MARGIN, y + 0.005, name,
             fontsize=10.5, color=INK, weight="bold", va="center")
    fig.text(LEFT_MARGIN, y - 0.018, unit,
             fontsize=8, color=TICK, va="center")

    # Barra — 3 zonas de fronteira fixa, como na arte original
    fig.patches.extend([
        Rectangle((BAR_LEFT, y - BAR_HEIGHT/2), ZONA_OTIMA_X - BAR_LEFT, BAR_HEIGHT,
                  facecolor=BAND_OK, edgecolor="none", transform=fig.transFigure, zorder=1),
        Rectangle((ZONA_OTIMA_X, y - BAR_HEIGHT/2), ZONA_LAB_X - ZONA_OTIMA_X, BAR_HEIGHT,
                  facecolor=BAND_MID, edgecolor="none", transform=fig.transFigure, zorder=1),
        Rectangle((ZONA_LAB_X, y - BAR_HEIGHT/2), BAR_RIGHT - ZONA_LAB_X, BAR_HEIGHT,
                  facecolor=BAND_BAD, edgecolor="none", transform=fig.transFigure, zorder=1),
    ])

    # Linha tracejada do limite
    if lim_lab and lim_pos is not None:
        # A tracejada do limite de laboratório fica em x 0.716 na arte
        # original, à direita da fronteira das zonas (0.684) — deixá-la sobre
        # a fronteira fazia o círculo da Vitamina D cair em cima dela.
        x_lim_line = 0.716
        fig.lines.append(plt.Line2D(
            [x_lim_line, x_lim_line], [y - BAR_HEIGHT*1.0, y + BAR_HEIGHT*1.0],
            color=INK, linewidth=0.8, linestyle=(0, (3, 2)),
            transform=fig.transFigure, zorder=4
        ))
        fig.text(x_lim_line, y - BAR_HEIGHT*1.0 - 0.010, lim_lab,
                 fontsize=7.5, color=INK_SOFT, ha="center", va="top",
                 linespacing=1.1)

    # Marker ótimo (triangulo)
    x_otimo_m = TRI_X
    fig.text(x_otimo_m, y, "▲", fontsize=12, color=INK_SOFT,
             ha="center", va="center", zorder=5)
    fig.text(x_otimo_m, y + BAR_HEIGHT*1.0 + 0.004, otimo_lab,
             fontsize=8, color=INK_SOFT, ha="center", va="bottom")

    # Marker Fernanda (círculo)
    x_fer_m = CIRC_X[i]
    fig.text(x_fer_m, y, "●", fontsize=12, color=INK,
             ha="center", va="center", zorder=6)
    fig.text(x_fer_m, y + BAR_HEIGHT*1.0 + 0.004, fer_lab,
             fontsize=8, color=INK, weight="bold", ha="center", va="bottom")

    # Valor à direita
    fig.text(VALUE_X, y, fer_value,
             fontsize=11, color=INK, weight="bold",
             ha="center", va="center", linespacing=1.15)

# Separador
SEP_Y = 1.0 - LINHA_Y[-1] - 0.055
fig.lines.append(plt.Line2D(
    [LEFT_MARGIN, RIGHT_MARGIN], [SEP_Y, SEP_Y],
    color="#CFCFCF", linewidth=0.5, transform=fig.transFigure
))

# Legenda inferior das zonas
LEG_Y2 = SEP_Y - 0.040
ZONE_W = (RIGHT_MARGIN - LEFT_MARGIN) / 3 - 0.01

zones = [
    ("ZONA ÓTIMA",   "Abaixo do alvo para longevidade",     BAND_OK),
    ("ZONA SUBÓTIMA", "Entre o ótimo e o limite do laboratório", BAND_MID),
    ("ZONA \"NORMAL\" DO LABORATÓRIO",
     "Dentro da referência — mas longe do ideal para longevidade", BAND_BAD),
]

for i, (label, desc, color) in enumerate(zones):
    x0 = LEFT_MARGIN + i * (ZONE_W + 0.015)
    fig.patches.append(
        Rectangle((x0, LEG_Y2 - 0.005), 0.025, 0.022,
                  facecolor=color, edgecolor="#888888", linewidth=0.5,
                  transform=fig.transFigure, zorder=3)
    )
    fig.text(x0 + 0.032, LEG_Y2 + 0.013, label,
             fontsize=8.5, color=INK, weight="bold", va="center")
    fig.text(x0 + 0.032, LEG_Y2 - 0.005, desc,
             fontsize=7.5, color=INK_SOFT, va="center")

# Footer
foot_lines = [
    "*TG/HDL calculado a partir do lipidograma convencional de Fernanda (TG 180, HDL 46), consistentes com TC 201 e LDL 119 pela equação de Friedewald.",
    "Todos os valores de Fernanda estavam dentro da referência do laboratório — mas todos na zona subótima, onde o risco se acumula silenciosamente.",
]
for i, line in enumerate(foot_lines):
    fig.text(LEFT_MARGIN, 0.060 - i * 0.020, line,
             fontsize=7.5, color=FOOT)

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap04_Fig01.pdf"
png_path = out_dir / "_preview_Cap04_Fig01.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
