"""
Cap01 Fig01 (PT-BR, B&W vetorial) — Três check-ups 'normais' — uma tendência perigosa.
Trajetória de HbA1c do Ricardo ao longo de quatro check-ups.

Versão preto-e-branco para a impressão paperback KDP.
Saída: PDF vetorial em build/versaoImpressa/figuras-bw/Cap01_Fig01.pdf
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams, patches
from matplotlib.patches import Patch
from matplotlib.ticker import FixedLocator, FixedFormatter

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False
rcParams["pdf.fonttype"] = 42       # TrueType nos PDFs (texto selecionável + vetor)
rcParams["ps.fonttype"] = 42

# ---------- paleta B&W ----------
BG          = "#FFFFFF"
INK         = "#000000"      # texto principal
INK_SOFT    = "#3A3A3A"      # subtítulos / band labels
TICK        = "#555555"
FOOT        = "#666666"
BAND_OK     = "#FAFAFA"      # zona ótima — quase branco
BAND_MID    = "#EEEEEE"      # subótima — cinza claro
BAND_BAD    = "#D8D8D8"      # pré-diabetes — cinza médio

# ---------- dados ----------
xs = [1, 2, 3, 4]
hba1c = [4.9, 5.2, 5.4, 5.7]

# ---------- figura ----------
# Aspecto 1.802, igual ao da arte original (1683×934). Era 1.774.
fig = plt.figure(figsize=(11.0, 6.104))
fig.patch.set_facecolor(BG)

# Caixa de plotagem MEDIDA na arte original: x 0.100–0.837, y 0.126–0.740
# (a partir do topo). A anterior era mais curta e começava mais abaixo, o que
# comprimia as faixas e desalinhava os rótulos de zona à direita.
ax = fig.add_axes([0.100, 0.260, 0.737, 0.614])
ax.set_facecolor(BG)

# ---------- bandas de fundo (3 tons de cinza sólidos, gradação do ótimo ao pior) ----------
ax.axhspan(5.7, 6.1, facecolor=BAND_BAD, edgecolor="none", linewidth=0, zorder=1)
ax.axhspan(5.2, 5.7, facecolor=BAND_MID, edgecolor="none", linewidth=0, zorder=1)
# Fronteira ótima/subótima em 5,2%, como na arte original — o gerador usava
# 5,0%, o que fazia a zona ótima invadir a faixa subótima.
ax.axhspan(4.45, 5.2, facecolor=BAND_OK,  edgecolor="none", linewidth=0, zorder=1)

# ---------- linha tracejada de referência em 5,7% ----------
ax.axhline(5.7, color=INK, linewidth=0.9,
           linestyle=(0, (4, 3)), zorder=2, alpha=0.85)
ax.text(0.6, 5.74, "Limite de normalidade (lab)",
        fontsize=8.5, color=INK_SOFT, style="italic", zorder=5)

# ---------- curva: sólida pontos 1-3, tracejada projeção 3→4 ----------
ax.plot(xs[:3], hba1c[:3],
        color=INK, linewidth=2.6, zorder=4,
        marker="o", markersize=10,
        markerfacecolor=INK,
        markeredgecolor=INK, markeredgewidth=2)

ax.plot(xs[2:], hba1c[2:],
        color=INK, linewidth=2.2, zorder=4,
        linestyle=(0, (5, 3)),
        marker="None")

# Ponto 4 — outline preto, preenchimento branco (projeção)
ax.plot([xs[3]], [hba1c[3]],
        marker="o", markersize=11,
        markerfacecolor="white",
        markeredgecolor=INK, markeredgewidth=2.4,
        linestyle="None", zorder=5)

# Valores acima de cada ponto
for x, y in zip(xs[:3], hba1c[:3]):
    ax.text(x, y + 0.08, f"{y:.1f}%".replace(".", ","),
            ha="center", va="bottom",
            fontsize=11, color=INK, weight="bold", zorder=6)

# Ponto 4 — valor em itálico bold (sem o vermelho original)
ax.text(xs[3], hba1c[3] + 0.08, f"{hba1c[3]:.1f}%".replace(".", ","),
        ha="center", va="bottom",
        fontsize=11, color=INK, weight="bold", style="italic", zorder=6)

# "✓ Normal" abaixo dos 3 primeiros pontos
for x, y in zip(xs[:3], hba1c[:3]):
    ax.text(x, y - 0.10, "✓ Normal",
            ha="center", va="top",
            fontsize=9, color=INK_SOFT, zorder=6)

# "Pré-diabetes" abaixo do ponto 4 — bold italic (substitui o vermelho)
ax.text(xs[3], hba1c[3] - 0.08, "Pré-diabetes",
        ha="center", va="top",
        fontsize=8.5, color=INK, weight="bold", style="italic", zorder=6)

# ---------- anotação com seta para o ponto 3 ----------
ax.annotate(
    "Cada resultado estava 'normal'.\nA tendência não estava.",
    xy=(2.85, 5.35),
    xytext=(2.5, 4.78),
    fontsize=10, color=INK, ha="center", va="center",
    style="italic",
    arrowprops=dict(
        arrowstyle="->",
        color=TICK,
        linewidth=1.0,
        connectionstyle="arc3,rad=-0.35",
    ),
    zorder=7,
)

# ---------- labels das bandas à direita ----------
ax.text(4.55, 5.85, "Pré-diabetes",
        fontsize=10, color=INK, weight="bold",
        va="center", ha="left", zorder=5)
ax.text(4.55, 5.35, "Normal —\nmas subótimo",
        fontsize=10, color=INK, weight="bold",
        va="center", ha="left", zorder=5)
ax.text(4.55, 4.75, "Zona ótima\npara longevidade",
        fontsize=10, color=INK_SOFT, weight="bold",
        va="center", ha="left", zorder=5)

# ---------- eixos ----------
ax.set_xlim(0.5, 4.5)
# ylim expandido 0.05 abaixo do menor tick (4,5%) para garantir que o label não
# seja escondido pelo matplotlib quando coincide com a borda do axes.
ax.set_ylim(4.45, 6.05)
ax.set_xticks(xs)

xtick_main = ["Check-up 1", "Check-up 2", "Check-up 3", "Check-up 4"]
xtick_sub  = ["(5 anos atrás)", "(2 anos atrás)", "(hoje)", "(projeção, daqui 2 anos)"]
ax.set_xticklabels([f"{m}\n{s}" for m, s in zip(xtick_main, xtick_sub)])

# FixedLocator + FixedFormatter forçam exibir TODOS os 6 ticks,
# evitando o auto-hide do matplotlib quando 5,0/5,2 e 5,5/5,7 ficam próximos.
_yticks = [4.5, 5.0, 5.2, 5.5, 5.7, 6.0]
_ylabels = ["4,5%", "5,0%", "5,2%", "5,5%", "5,7%", "6,0%"]
ax.yaxis.set_major_locator(FixedLocator(_yticks))
ax.yaxis.set_major_formatter(FixedFormatter(_ylabels))
ax.tick_params(axis="x", colors=TICK, labelsize=9, length=0, pad=8)
ax.tick_params(axis="y", colors=TICK, labelsize=8, length=0, pad=4)
# Garantir todos os yticks visíveis (matplotlib pode auto-esconder por colisão)
for _lbl in ax.yaxis.get_ticklabels():
    _lbl.set_visible(True)
for spine in ("top", "right"):
    ax.spines[spine].set_visible(False)
ax.spines["left"].set_color("#888888")
ax.spines["bottom"].set_color("#888888")

# Rótulo do eixo Y
fig.text(0.04, 0.49, "HbA1c (%)",
         fontsize=10, color=INK, weight="bold",
         rotation=90, va="center", ha="center")

# ---------- título ----------
fig.text(0.04, 0.910,
         "Três check-ups 'normais' — uma tendência perigosa",
         fontsize=20, color=INK, weight="bold")

# ---------- fonte/rodapé (3 linhas) ----------
source_lines = [
    "HbA1c (hemoglobina glicada) reflete a média do açúcar no sangue nos últimos 2-3 meses.",
    "A 'faixa normal' do laboratório (abaixo de 5,7%) inclui valores que, embora não configurem pré-diabetes,",
    "já estão longe do ótimo para longevidade. A tendência importa mais que o número isolado.",
]
for i, line in enumerate(source_lines):
    fig.text(0.04, 0.080 - i * 0.020, line,
             fontsize=8, color=FOOT)

# ---------- save: PDF vetorial (para o build) + PNG (para comparação visual) ----------
out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap01_Fig01.pdf"
png_path = out_dir / "_preview_Cap01_Fig01.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
