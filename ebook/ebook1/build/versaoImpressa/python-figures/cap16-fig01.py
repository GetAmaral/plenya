"""
Cap16 Fig01 (PT-BR, B&W vetorial) — Dois modelos de acompanhamento médico ao longo de um ano.
Layout portrait com 2 painéis.
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
PANEL_1  = "#F2F2F2"
PANEL_2  = "#E8E8E8"
BAND     = "#EDEDED"

_FIG_W, _FIG_H = 8.4, 11.4
_ASPECT = _FIG_W / _FIG_H

fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

LEFT = 0.03

# Título
fig.text(LEFT, 0.965, "FIGURA 1.",
         fontsize=8, color=INK_SOFT, weight="bold")
fig.text(LEFT, 0.945,
         "Dois modelos de acompanhamento médico ao longo de um ano.",
         fontsize=15, color=INK, weight="bold")
fig.text(LEFT, 0.918,
         "Check-up convencional vs. avaliação preventiva ampliada.",
         fontsize=10, color=INK_SOFT, style="italic")

# ============== PAINEL 1: MODELO CONVENCIONAL ==============
P1_TOP, P1_BOT = 0.890, 0.560

fig.patches.append(FancyBboxPatch(
    (LEFT, P1_BOT), 0.94, P1_TOP - P1_BOT,
    boxstyle="round,pad=0.005,rounding_size=0.008",
    facecolor=PANEL_1, edgecolor=INK, linewidth=0.8,
    transform=fig.transFigure, zorder=1
))

# Header lateral esquerdo
fig.text(LEFT + 0.020, P1_TOP - 0.020, "MODELO",
         fontsize=10, color=INK, weight="bold")
fig.text(LEFT + 0.020, P1_TOP - 0.040, "CONVENCIONAL",
         fontsize=11, color=INK, weight="bold")

fig.text(LEFT + 0.020, P1_TOP - 0.075,
         "Acompanhamento\npontual e com painel\nbásico e faixas\nde referência.",
         fontsize=8, color=INK_SOFT, va="top", linespacing=1.3)

# Timeline com 3 markers
TL1_LEFT  = LEFT + 0.27
TL1_RIGHT = LEFT + 0.92
TL1_Y     = P1_BOT + 0.085

def draw_timeline_arrow(x1, x2, y):
    fig.patches.append(FancyArrowPatch(
        (x1, y), (x2, y),
        arrowstyle="->", color=INK, lw=1.0, mutation_scale=14,
        transform=fig.transFigure, zorder=3
    ))

draw_timeline_arrow(TL1_LEFT, TL1_RIGHT, TL1_Y)

# Markers ao longo da timeline
def small_dot(x, y, r=0.008, fill=INK):
    fig.patches.append(Ellipse(
        (x, y), width=r*2, height=r*2*_ASPECT,
        facecolor=fill, edgecolor=INK, linewidth=0.8,
        transform=fig.transFigure, zorder=4
    ))

tl1_markers = [
    (TL1_LEFT + 0.00, "MÊS 0",  "Consulta anual\n(20–30 min)\n\nPainel básico"),
    (TL1_LEFT + 0.08, "MÊS 1",  "Resultados\n\"Dentro da faixa\nde referência\""),
    (TL1_RIGHT - 0.005, "MÊS 12", "Nova consulta\nanual\n\nCiclo reinicia"),
]
for x, label, desc in tl1_markers:
    small_dot(x, TL1_Y)
    fig.text(x, TL1_Y + 0.025, label,
             fontsize=8.5, color=INK, weight="bold", ha="center", va="bottom")
    fig.text(x, TL1_Y - 0.022, desc,
             fontsize=7, color=INK, ha="center", va="top",
             linespacing=1.25)

# Texto entre os markers
mid_x = (tl1_markers[1][0] + tl1_markers[2][0]) / 2
fig.text(mid_x, TL1_Y - 0.005, "Sem acompanhamento",
         fontsize=7.5, color=INK_SOFT, ha="center", va="center", style="italic")

# Escala de tempo (0 → 12 MESES)
fig.text(TL1_LEFT, TL1_Y - 0.080, "0",
         fontsize=8, color=INK_SOFT, ha="center", weight="bold")
fig.text(TL1_RIGHT, TL1_Y - 0.080, "12   MESES",
         fontsize=8, color=INK_SOFT, ha="center", weight="bold")
# Linha base
fig.lines.append(plt.Line2D(
    [TL1_LEFT, TL1_RIGHT], [TL1_Y - 0.065, TL1_Y - 0.065],
    color=INK, linewidth=0.8, transform=fig.transFigure
))

# ============== PAINEL 2: AVALIAÇÃO PREVENTIVA AMPLIADA ==============
P2_TOP, P2_BOT = 0.540, 0.180

fig.patches.append(FancyBboxPatch(
    (LEFT, P2_BOT), 0.94, P2_TOP - P2_BOT,
    boxstyle="round,pad=0.005,rounding_size=0.008",
    facecolor=PANEL_2, edgecolor=INK, linewidth=1.0,
    transform=fig.transFigure, zorder=1
))

fig.text(LEFT + 0.020, P2_TOP - 0.020, "AVALIAÇÃO",
         fontsize=10, color=INK, weight="bold")
fig.text(LEFT + 0.020, P2_TOP - 0.040, "PREVENTIVA",
         fontsize=11, color=INK, weight="bold")
fig.text(LEFT + 0.020, P2_TOP - 0.057, "AMPLIADA",
         fontsize=11, color=INK, weight="bold")

fig.text(LEFT + 0.020, P2_TOP - 0.090,
         "Acompanhamento\ncontínuo e iterativo\ncom painel estendido,\nfaixas ótimas e\nplano personalizado.",
         fontsize=8, color=INK_SOFT, va="top", linespacing=1.3)

# Timeline com 6 markers
TL2_LEFT  = LEFT + 0.27
TL2_RIGHT = LEFT + 0.92
TL2_Y     = P2_BOT + 0.180

draw_timeline_arrow(TL2_LEFT, TL2_RIGHT, TL2_Y)

# 6 markers — offsets aumentados pra evitar overlap
TL2_MARKERS = [
    (0.00, "MÊS 0",       "Primeira consulta\n(60–120 min)\n\nAnamnese profunda\n+ solicitação\nde painel"),
    (0.16, "SEMANAS\n2-3", "Exames\nrealizados\n\nPainel ampliado\n(biomarcadores\n+ imagem)"),
    (0.32, "MÊS 1",       "Retorno estruturado\n(60–90 min)\n\nLeitura conjunta\nFaixas ótimas\nPlano com paciente"),
    (0.50, "MÊS 4",       "Primeiro retorno\n\nReavaliação dos\nmarcadores-alvo.\nAjuste de plano."),
    (0.70, "MÊS 8",       "Segundo retorno\n\nReavaliação.\nNovos marcadores\nse necessário."),
    (0.95, "MÊS 12",      "Avaliação anual\ncompleta\n\nPainel completo.\nComposição corporal.\nPlacar revisado."),
]
for offset, label, desc in TL2_MARKERS:
    x = TL2_LEFT + offset * (TL2_RIGHT - TL2_LEFT)
    small_dot(x, TL2_Y)
    fig.text(x, TL2_Y + 0.025, label,
             fontsize=7, color=INK, weight="bold", ha="center", va="bottom",
             linespacing=1.1)
    fig.text(x, TL2_Y - 0.022, desc,
             fontsize=6, color=INK, ha="center", va="top",
             linespacing=1.3)

# Ciclo iterativo (texto abaixo)
fig.text((TL2_LEFT + TL2_RIGHT) / 2, P2_BOT + 0.060,
         "Ciclo iterativo: testar → intervir → reavaliar → ajustar",
         fontsize=9, color=INK, weight="bold", ha="center", style="italic")

# Escala (atualizada pros novos offsets)
fig.text(TL2_LEFT, P2_BOT + 0.035, "0",
         fontsize=8, color=INK_SOFT, ha="center", weight="bold")
fig.text(TL2_LEFT + 0.32 * (TL2_RIGHT - TL2_LEFT), P2_BOT + 0.035, "1",
         fontsize=8, color=INK_SOFT, ha="center", weight="bold")
fig.text(TL2_LEFT + 0.50 * (TL2_RIGHT - TL2_LEFT), P2_BOT + 0.035, "4",
         fontsize=8, color=INK_SOFT, ha="center", weight="bold")
fig.text(TL2_LEFT + 0.70 * (TL2_RIGHT - TL2_LEFT), P2_BOT + 0.035, "8",
         fontsize=8, color=INK_SOFT, ha="center", weight="bold")
fig.text(TL2_RIGHT, P2_BOT + 0.035, "12   MESES",
         fontsize=8, color=INK_SOFT, ha="center", weight="bold")

fig.lines.append(plt.Line2D(
    [TL2_LEFT, TL2_RIGHT], [P2_BOT + 0.050, P2_BOT + 0.050],
    color=INK, linewidth=0.8, transform=fig.transFigure
))

# ============== Caixa final ==============
BOX_X1, BOX_X2 = LEFT, LEFT + 0.94
BOX_Y1, BOX_Y2 = 0.080, 0.170

fig.patches.append(FancyBboxPatch(
    (BOX_X1, BOX_Y1), BOX_X2 - BOX_X1, BOX_Y2 - BOX_Y1,
    boxstyle="round,pad=0.005,rounding_size=0.008",
    facecolor=BAND, edgecolor=INK, linewidth=0.6,
    transform=fig.transFigure, zorder=1
))

# Ícone (i)
fig.patches.append(Ellipse(
    (BOX_X1 + 0.030, (BOX_Y1 + BOX_Y2)/2),
    width=0.024, height=0.024 * _ASPECT,
    facecolor=INK, edgecolor=INK, linewidth=0.5,
    transform=fig.transFigure, zorder=3
))
fig.text(BOX_X1 + 0.030, (BOX_Y1 + BOX_Y2)/2, "i",
         fontsize=11, color="white", weight="bold",
         ha="center", va="center", style="italic", zorder=4)

# Texto da caixa
TX = BOX_X1 + 0.060
fig.text(TX, BOX_Y2 - 0.020,
         "O modelo convencional opera em ciclos anuais com avaliação pontual e ausência de acompanhamento",
         fontsize=8, color=INK)
fig.text(TX, BOX_Y2 - 0.035,
         "intermediário. A avaliação preventiva ampliada utiliza ciclos iterativos ao longo do ano, com painel",
         fontsize=8, color=INK)
fig.text(TX, BOX_Y2 - 0.050,
         "estendido, definição de alvos, intervenções priorizadas e reavaliações periódicas.",
         fontsize=8, color=INK)
fig.text(TX, BOX_Y2 - 0.072,
         "A diferença central é estrutural: continuidade e ajuste ao longo do tempo.",
         fontsize=8.5, color=INK, weight="bold", style="italic")

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap16_Fig01.pdf"
png_path = out_dir / "_preview_Cap16_Fig01.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
