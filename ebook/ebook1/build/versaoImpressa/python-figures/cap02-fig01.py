"""
Cap02 Fig01 (PT-BR, B&W vetorial) — Os 20 Anos Silenciosos.

Quatro doenças em timelines paralelas. Cada uma com início silencioso →
janela de intervenção → diagnóstico convencional. Eixo de idade no topo.

Saída: PDF vetorial em build/versaoImpressa/figuras-bw/Cap02_Fig01.pdf
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, FancyArrowPatch

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
BAND_LIGHT  = "#F0F0F0"   # zona "silencioso" — mais claro
BAND_MID    = "#D6D6D6"   # janela de intervenção — médio
BAND_DARK   = "#A8A8A8"   # diagnóstico convencional — escuro

# ---------- dados ----------
# Cada doença: (nome (2 linhas), idade_inicio_silencioso, idade_inicio_janela,
#               idade_fim_janela, idade_diagnostico_inicio, idade_diagnostico_fim,
#               label_inicio, label_diagnostico, has_arrow)
diseases = [
    ("DOENÇA\nCARDIOVASCULAR",
     30, 30, 55, 55, 65,
     "30–50 anos\nestrias gordurosas\nApoB sobe",
     "55–65 anos\nIAM, AVC", False),
    ("DOENÇA\nMETABÓLICA",
     30, 30, 50, 50, 60,
     "30–40 anos\nInsulina começa a subir",
     "50–60 anos\nDiagnóstico de diabetes", False),
    ("NEURODEGENERAÇÃO",
     40, 40, 65, 65, 75,
     "40–50 anos\nDeposição de amilóide",
     "65–75 anos\nDeclínio cognitivo perceptível", False),
    ("CÂNCER",
     40, 40, 65, 65, 75,
     "40–50 anos\nTerreno metabólico estabelecido",
     "Variável\nDiagnóstico oncológico", True),
]

# ---------- figura ----------
# Aspecto 1.777, igual ao da arte original (1672×941). Era 1.719.
fig = plt.figure(figsize=(11.0, 6.19))
fig.patch.set_facecolor(BG)

LEFT_MARGIN  = 0.025
LABEL_COL    = 0.18   # onde termina a coluna do nome da doença
TIMELINE_L   = 0.20   # início da timeline (eixo idade)
TIMELINE_R   = 0.97   # fim da timeline
AGE_MIN, AGE_MAX = 30, 75

def age_to_x(age):
    return TIMELINE_L + (age - AGE_MIN) / (AGE_MAX - AGE_MIN) * (TIMELINE_R - TIMELINE_L)

# ---------- título ----------
fig.text(LEFT_MARGIN, 0.945, "Figura 1 — Os 20 Anos Silenciosos",
         fontsize=17, color=INK, weight="bold", va="center")

# Subtítulo
fig.text(LEFT_MARGIN, 0.905,
         "Quatro doenças. Um mesmo padrão silencioso. 10 a 20 anos de avanço antes do diagnóstico.",
         fontsize=10, color=INK_SOFT)

# ---------- "Check-up: 'tudo normal'" annotation no topo (alinhado com idade 50) ----------
x_50 = age_to_x(50)
fig.text(x_50, 0.853, "Check-up:\n'tudo normal'",
         fontsize=8.5, color=INK, weight="bold", style="italic",
         ha="center", va="bottom", linespacing=1.1)
# Setinha vertical curta abaixo do texto
fig.patches.append(FancyArrowPatch(
    (x_50, 0.845), (x_50, 0.815),
    arrowstyle="->", color=INK, lw=1.0,
    mutation_scale=12, transform=fig.transFigure, zorder=4
))

# ---------- eixo de idade ----------
AXIS_Y = 0.795
fig.text(LEFT_MARGIN, AXIS_Y, "IDADE (ANOS)",
         fontsize=8.5, color=TICK, weight="bold", va="center")
for age in [30, 40, 50, 60, 70]:
    x = age_to_x(age)
    fig.text(x, AXIS_Y, str(age),
             fontsize=10.5, color=INK, weight="bold", ha="center", va="center")

# Linha sutil do eixo
fig.lines.append(plt.Line2D(
    [TIMELINE_L, TIMELINE_R], [AXIS_Y - 0.018, AXIS_Y - 0.018],
    color="#BBBBBB", linewidth=0.5, transform=fig.transFigure, zorder=1
))

# Linha tracejada vertical no 50 (continuando pra baixo, atravessando as timelines)
fig.lines.append(plt.Line2D(
    [x_50, x_50], [0.085, AXIS_Y - 0.018],
    color=INK, linewidth=0.8, linestyle=(0, (3, 3)),
    transform=fig.transFigure, zorder=2
))

# ---------- linhas de doenças ----------
# Centro de cada linha MEDIDO na arte original (a faixa verde da janela de
# intervenção ocupa y 0.332–0.407, 0.467–0.540, 0.597–0.668 e 0.728–0.800,
# contando do topo). A grade anterior punha as quatro linhas mais acima e com
# passo diferente, o que desalinhava tudo em relação ao original.
LINHA_Y    = [0.370, 0.504, 0.633, 0.764]
ROW_TOP    = 1.0 - LINHA_Y[0]
ROW_BOTTOM = 1.0 - LINHA_Y[-1]
ROW_SPACE  = (ROW_TOP - ROW_BOTTOM) / (len(diseases) - 1)
BAR_HEIGHT = 0.058

for i, (name, silent_start, jan_start, jan_end, diag_start, diag_end,
        label_left, label_right, has_arrow) in enumerate(diseases):
    y = 1.0 - LINHA_Y[i]

    # Bullet à esquerda
    fig.text(LEFT_MARGIN + 0.000, y, "●",
             fontsize=12, color=INK, va="center")
    # Nome da doença
    fig.text(LEFT_MARGIN + 0.012, y, name,
             fontsize=10, color=INK, weight="bold", va="center",
             linespacing=1.1)

    # ---- barra única da janela de intervenção ----
    x_jan_l = age_to_x(jan_start)
    x_jan_r = age_to_x(jan_end)
    x_diag_r = age_to_x(diag_end)

    # Uma barra contínua média
    fig.patches.append(Rectangle(
        (x_jan_l, y - BAR_HEIGHT/2), x_jan_r - x_jan_l, BAR_HEIGHT,
        facecolor=BAND_MID, edgecolor="none", transform=fig.transFigure, zorder=2
    ))

    # Markers (bolinhas) início e fim da janela
    fig.text(x_jan_l, y, "●", fontsize=10, color=INK,
             ha="center", va="center", zorder=4)
    fig.text(x_jan_r, y, "●", fontsize=10, color=INK,
             ha="center", va="center", zorder=4)

    # Label "JANELA DE INTERVENÇÃO" centralizado na barra
    fig.text((x_jan_l + x_jan_r) / 2, y, "JANELA DE INTERVENÇÃO",
             fontsize=9, color=INK, weight="bold",
             ha="center", va="center", zorder=5)

    # Label esquerda (início silencioso) — abaixo do começo
    fig.text(x_jan_l + 0.005, y - BAR_HEIGHT/2 - 0.005, label_left,
             fontsize=7.5, color=INK_SOFT, ha="left", va="top",
             linespacing=1.15)

    # Label direita (diagnóstico) — fora da barra, à direita
    fig.text(x_jan_r + 0.008, y, label_right,
             fontsize=7.5, color=INK, weight="bold",
             ha="left", va="center", linespacing=1.15)

    # Seta após a label do câncer ("Variável → ...")
    if has_arrow:
        fig.patches.append(FancyArrowPatch(
            (x_diag_r - 0.005, y), (x_diag_r + 0.012, y),
            arrowstyle="->", color=INK, lw=1.2,
            mutation_scale=14, transform=fig.transFigure, zorder=5
        ))

# ---------- frase abaixo ----------
fig.text(0.5, 0.110,
         "A doença já estava lá. O diagnóstico é que chegou tarde.",
         fontsize=11, color=INK, weight="bold", style="italic",
         ha="center")

# ---------- footer fontes ----------
fig.text(LEFT_MARGIN, 0.045,
         "Fontes: Libby P et al, JACC, 2019; Alba ZJ et al, Circulation, 2019; "
         "Jack CR Jr et al, Lancet Neurol, 2018; Hanahan D, Cell, 2022.",
         fontsize=7.5, color=FOOT)

# ---------- save ----------
out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap02_Fig01.pdf"
png_path = out_dir / "_preview_Cap02_Fig01.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
# recorte automático removido: as posições agora são as medidas na arte
# original em fração de figura, e recortar o branco muda a proporção final
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG)
plt.savefig(png_path, dpi=170, facecolor=BG)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
