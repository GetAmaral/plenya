"""
Cap13 Fig01 (PT-BR, B&W vetorial) — O preço da solidão: impacto na mortalidade por todas as causas.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, FancyBboxPatch

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
SOCIAL_C = "#3A3A3A"    # FATORES SOCIAIS — escuro
CLASSIC_C = "#9A9A9A"   # FATORES CLÁSSICOS — médio
BAND     = "#F0F0F0"

fig = plt.figure(figsize=(11.0, 7.5))
fig.patch.set_facecolor(BG)

# Título
fig.text(0.025, 0.945,
         "Figura 2 — O preço da solidão: impacto na mortalidade por todas as causas",
         fontsize=15, color=INK, weight="bold")
fig.text(0.025, 0.910,
         "Aumento do risco de morte em comparação a quem não tem o fator.",
         fontsize=9.5, color=INK_SOFT, style="italic")

ax = fig.add_axes([0.28, 0.28, 0.55, 0.56])

bars = [
    ("Viver sozinho\n(sem companhia)",                32, SOCIAL_C, "social"),
    ("Isolamento social\n(pouco contato social)",     29, SOCIAL_C, "social"),
    ("Solidão\n(sentimento de solidão)",              26, SOCIAL_C, "social"),
    ("Tabagismo\n(fumante atual)",                    70, CLASSIC_C, "classic"),
    ("Sedentarismo\n(inatividade física)",            25, CLASSIC_C, "classic"),
    ("Obesidade\n(IMC > 30)",                         20, CLASSIC_C, "classic"),
    ("Poluição do ar\n(exposição a PM 2.5)",          15, CLASSIC_C, "classic"),
]

labels = [b[0] for b in bars]
values = [b[1] for b in bars]
colors = [b[2] for b in bars]

y_pos = list(range(len(bars)))
ax.barh(y_pos, values, color=colors, height=0.62,
        edgecolor="none", zorder=3)

# Valores
for i, v in enumerate(values):
    ax.text(v + 1.0, i, f"{v}%",
            color=INK, weight="bold", fontsize=10, va="center")

ax.set_yticks(y_pos)
ax.set_yticklabels(labels, fontsize=9, color=INK,
                    linespacing=1.15)
ax.invert_yaxis()

ax.set_xlim(0, 80)
ax.set_xticks([0, 10, 20, 30, 40, 50, 60, 70, 80])
ax.set_xlabel("Aumento do risco de morte (%)",
              fontsize=10, color=INK, weight="bold")

for spine in ("top", "right"):
    ax.spines[spine].set_visible(False)
ax.spines["left"].set_color("#888888")
ax.spines["bottom"].set_color("#888888")
ax.tick_params(axis="y", length=0)
ax.tick_params(axis="x", colors=TICK)
ax.set_axisbelow(True)
ax.grid(axis="x", color="#EEEEEE", linewidth=0.6, zorder=0)

# Divisor "FATORES SOCIAIS" / "FATORES CLÁSSICOS"
# Acima do grupo social
fig.text(0.025, 0.745, "FATORES",
         fontsize=9, color=INK, weight="bold")
fig.text(0.025, 0.728, "SOCIAIS",
         fontsize=11, color=INK, weight="bold")

fig.text(0.025, 0.555, "FATORES",
         fontsize=9, color=INK, weight="bold")
fig.text(0.025, 0.540, "CLÁSSICOS",
         fontsize=11, color=INK, weight="bold")

# Linha separadora entre os 2 grupos (após row 3 do grupo social)
ax.axhline(2.55, color="#BBBBBB", linewidth=0.6, linestyle="-", zorder=1)

# Annotation pra "Fatores sociais"
fig.text(0.84, 0.680, "Fatores sociais —\nraramente avaliados\nno consultório,\nmas com impacto\ncomparável a\nfatores clássicos.",
         fontsize=8, color=INK, weight="bold",
         linespacing=1.3, va="top")

# Annotation tabagismo
fig.text(0.84, 0.480, "Referência clássica:\nfator de risco modificável.",
         fontsize=8, color=INK_SOFT, style="italic",
         linespacing=1.3, va="top")

# ---------- Caixa "Como ler" ----------
BOX_X1, BOX_X2 = 0.025, 0.50
BOX_Y1, BOX_Y2 = 0.140, 0.215

fig.patches.append(FancyBboxPatch(
    (BOX_X1, BOX_Y1), BOX_X2 - BOX_X1, BOX_Y2 - BOX_Y1,
    boxstyle="round,pad=0.005,rounding_size=0.005",
    facecolor=BAND, edgecolor=INK, linewidth=0.6,
    transform=fig.transFigure, zorder=1
))
fig.text(BOX_X1 + 0.010, BOX_Y2 - 0.018,
         "COMO LER ESTE GRÁFICO:",
         fontsize=8.5, color=INK, weight="bold")
fig.text(BOX_X1 + 0.010, BOX_Y2 - 0.038,
         "As barras mostram o quanto o risco de morte aumenta em pessoas com",
         fontsize=7.5, color=INK)
fig.text(BOX_X1 + 0.010, BOX_Y2 - 0.054,
         "aquele fator, em comparação a quem não tem o fator.",
         fontsize=7.5, color=INK)

# Right footer — recuado pra direita pra não colidir com "Como ler"
fig.text(0.97, 0.205,
         "Não é uma comparação direta entre o Fig 1 e o Cap 13.",
         fontsize=7.5, color=FOOT, style="italic", ha="right")
fig.text(0.97, 0.190,
         "Trazido para situar a ordem de magnitude.",
         fontsize=7.5, color=FOOT, style="italic", ha="right")
fig.text(0.97, 0.175,
         "Em 2023, o Surgeon General dos EUA classifica",
         fontsize=7.5, color=FOOT, style="italic", ha="right")
fig.text(0.97, 0.160,
         "a solidão como problema de saúde pública de prioridade nacional.",
         fontsize=7.5, color=FOOT, style="italic", ha="right")

# Footer (sources)
fig.text(0.025, 0.080,
         "Fontes: Holt-Lunstad et al., PLOS Medicine 2010; Holt-Lunstad et al., Perspectives on Psychological Science 2015;",
         fontsize=7, color=FOOT)
fig.text(0.025, 0.060,
         "Berk et al., BMJ 2015; The Lancet 2016; Lee et al., Lancet 2012; WHO Global Health Risks / FBC 2020.",
         fontsize=7, color=FOOT)

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap13_Fig01.pdf"
png_path = out_dir / "_preview_Cap13_Fig01.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
