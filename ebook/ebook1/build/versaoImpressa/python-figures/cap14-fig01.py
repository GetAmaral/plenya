"""
Cap14 Fig01 (PT-BR, B&W vetorial) — A arquitetura de uma noite: o que distingue dormir de descansar.
2 hipnogramas: normal vs. fragmentado pela apneia.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
import numpy as np

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
N3_BAND  = "#E8E8E8"
REM_BAND = "#F2F2F2"

fig = plt.figure(figsize=(11.0, 7.5))
fig.patch.set_facecolor(BG)

# Título
fig.text(0.025, 0.945,
         "Figura 1 — A arquitetura de uma noite: o que distingue dormir de descansar",
         fontsize=15, color=INK, weight="bold")
fig.text(0.025, 0.910,
         "Mesmo tempo na cama, resultados completamente diferentes.",
         fontsize=9.5, color=INK_SOFT, style="italic")

# 2 painéis empilhados
ax1 = fig.add_axes([0.10, 0.55, 0.78, 0.30])
ax2 = fig.add_axes([0.10, 0.15, 0.78, 0.30])

stage_y = {"Acordado": 4, "REM": 3, "N1": 2, "N2": 1, "N3": 0}
stage_labels = ["Acordado", "REM", "N1", "N2", "N3"]

def setup_axis(ax, title):
    ax.set_yticks([0, 1, 2, 3, 4])
    ax.set_yticklabels(["N3", "N2", "N1", "REM", "Acordado"][::-1] if False else
                       ["N3", "N2", "N1", "REM", "Acordado"],
                       fontsize=8, color=INK)
    # Já está em ordem certa do bottom (N3) ao top (Acordado)
    # Mas eu prefiro N3 embaixo, vou flipar:
    ax.set_yticks([0, 1, 2, 3, 4])
    ax.set_yticklabels(["Acordado", "REM", "N1", "N2", "N3"])
    ax.set_xlim(0, 8)
    ax.set_ylim(-0.5, 4.5)
    ax.set_xticks([0, 1, 2, 3, 4, 5, 6, 7, 8])
    ax.set_xlabel("Tempo (Horas)", fontsize=9, color=INK, weight="bold")
    ax.tick_params(axis="both", labelsize=8, colors=TICK)
    for spine in ("top", "right"):
        ax.spines[spine].set_visible(False)
    ax.spines["left"].set_color("#888888")
    ax.spines["bottom"].set_color("#888888")
    # Banda N3 (no fundo) e REM (claro)
    ax.axhspan(-0.5, 0.5, facecolor=N3_BAND, zorder=0)
    ax.axhspan(2.5, 3.5, facecolor=REM_BAND, zorder=0)
    ax.set_title(title, fontsize=11, color=INK, weight="bold", loc="left")

# Hipnograma "normal" — ciclos ~90min, N3 profundo nas primeiras horas, REM prolongado depois
# Y: 0=Acordado, 1=REM, 2=N1, 3=N2, 4=N3 (invertido)
# Vou usar simples: 0=N3 (deep), 4=Acordado, then plot
# Atualmente: 0=Acordado, 1=REM, 2=N1, 3=N2, 4=N3 — inversed Y

# REVERTING: Standard hypnogram has Acordado at TOP and N3 at BOTTOM.
# y=4: Acordado (top), y=3: REM, y=2: N1, y=1: N2, y=0: N3 (bottom)
# Set yticklabels accordingly

for ax, title in [(ax1, "Arquitetura normal"), (ax2, "Arquitetura fragmentada pela apneia (Paulo, pré-CPAP)")]:
    ax.set_yticks([0, 1, 2, 3, 4])
    ax.set_yticklabels(["N3", "N2", "N1", "REM", "Acordado"])
    ax.set_xlim(0, 8)
    ax.set_ylim(-0.5, 4.5)
    ax.set_xticks([0, 1, 2, 3, 4, 5, 6, 7, 8])
    ax.set_xlabel("Tempo (Horas)", fontsize=9, color=INK, weight="bold")
    ax.tick_params(axis="both", labelsize=8, colors=TICK)
    for spine in ("top", "right"):
        ax.spines[spine].set_visible(False)
    ax.spines["left"].set_color("#888888")
    ax.spines["bottom"].set_color("#888888")
    # Bands
    ax.axhspan(-0.5, 0.5, facecolor=N3_BAND, zorder=0)
    ax.axhspan(2.5, 3.5, facecolor=REM_BAND, zorder=0)
    ax.set_title(title, fontsize=11, color=INK, weight="bold", loc="left")

# === HIPNOGRAMA NORMAL ===
# Sequence: t=[step times], y=[stage values]
# 4 ciclos: cada ciclo ~90min = 1.5h
# Ciclo 1 (0-1.5h): Acordado→N1→N2→N3 (profundo)→N2→REM curto
# Ciclo 2 (1.5-3.0h): N2→N3→N2→REM
# Ciclo 3 (3.0-4.5h): N2→N3 menor→N2→REM mais longo
# Ciclo 4 (4.5-6.0h): N1→N2→REM longo
# Ciclo 5 (6.0-7.5h): N2→REM final

normal_steps = [
    (0.0, 4),  # Acordado
    (0.1, 2),  # N1
    (0.2, 1),  # N2
    (0.4, 0),  # N3 profundo
    (0.9, 0),  # N3 mantido
    (1.1, 1),  # N2
    (1.3, 3),  # REM curto
    (1.4, 3),
    (1.5, 1),  # N2
    (1.8, 0),  # N3 ciclo 2
    (2.4, 0),
    (2.6, 1),
    (2.8, 3),  # REM ciclo 2
    (3.0, 3),
    (3.1, 1),
    (3.3, 0),  # N3 menos profundo ciclo 3
    (3.8, 0),
    (4.0, 1),
    (4.3, 3),  # REM mais longo
    (4.6, 3),
    (4.7, 2),
    (5.0, 1),
    (5.3, 3),  # REM ciclo 4 longo
    (5.9, 3),
    (6.0, 2),
    (6.2, 1),
    (6.4, 3),  # REM longo ciclo 5
    (7.2, 3),
    (7.5, 4),  # Acordado fim
]

ts, ys = zip(*normal_steps)
ax1.step(ts, ys, where="post", color=INK, linewidth=1.4, zorder=3)

# Annotations no painel 1 — texto inline curto dentro da banda N3 (bottom)
ax1.text(0.7, 0.2, "← N3 profundo: clareance glinfática + limpeza de β-amiloide",
         fontsize=7, color=INK_SOFT, ha="left", va="bottom", style="italic")
ax1.text(7.7, 3.2, "REM prolongado +\nconsolidação emocional",
         fontsize=7, color=INK_SOFT, ha="right", va="bottom",
         style="italic", linespacing=1.2)

ax1.text(8.3, 2.0, "4–5 ciclos\nN3 prolongado na\nprimeira metade.\nREM mais longo\nna segunda.",
         fontsize=7.5, color=INK, weight="bold", va="center",
         linespacing=1.3, clip_on=False)

# === HIPNOGRAMA FRAGMENTADO ===
# Múltiplos despertares + N3 raro
frag_steps = [
    (0.0, 4),
    (0.1, 2),
    (0.3, 1),
    (0.5, 1),
    (0.6, 4),  # despertar
    (0.7, 2),
    (0.9, 1),
    (1.0, 4),  # despertar
    (1.1, 2),
    (1.3, 1),
    (1.5, 0),  # N3 brief
    (1.6, 1),
    (1.7, 4),  # despertar
    (1.9, 2),
    (2.1, 1),
    (2.3, 3),  # REM brief
    (2.4, 4),  # interrompido
    (2.5, 2),
    (2.8, 1),
    (3.0, 4),  # despertar
    (3.1, 2),
    (3.3, 1),
    (3.5, 4),  # despertar
    (3.6, 2),
    (3.9, 1),
    (4.1, 3),  # REM brief
    (4.3, 4),  # interrompido
    (4.4, 2),
    (4.7, 1),
    (4.9, 4),  # despertar
    (5.0, 2),
    (5.3, 1),
    (5.5, 3),  # REM
    (5.7, 4),  # interrompido
    (5.8, 2),
    (6.1, 1),
    (6.3, 4),  # despertar
    (6.4, 2),
    (6.6, 3),  # REM
    (6.8, 4),  # interrompido
    (7.0, 1),
    (7.3, 4),  # despertar
    (7.5, 4),
]
ts, ys = zip(*frag_steps)
ax2.step(ts, ys, where="post", color=INK, linewidth=1.4, zorder=3)

# Annotations painel 2 — inline curto
ax2.text(0.2, 3.5, "Cada acordar = queda de saturação + micro-despertar",
         fontsize=7, color=INK_SOFT, ha="left", va="bottom", style="italic")
ax2.text(0.2, 0.2, "← N3 quase ausente: sono profundo interrompido antes de aprofundar",
         fontsize=7, color=INK_SOFT, ha="left", va="bottom", style="italic")

ax2.text(8.3, 2.0, "Múltiplos micro-\ndespertares\n(~15–20 na noite)\n\nN3 < 5% do total.\nREM frequente,\nmas truncado e\nde rebote.", clip_on=False,
         fontsize=8, color=INK, weight="bold", va="center",
         linespacing=1.3)

# Footer
fig.text(0.025, 0.075,
         "Hipnogramas ilustrativos representando (acima) arquitetura normal de sono em adulto com aproximadamente 7 horas e meia de duração,",
         fontsize=7.5, color=FOOT, style="italic")
fig.text(0.025, 0.060,
         "N3 concentrado na primeira metade da noite e REM prolongado na segunda; (abaixo) arquitetura fragmentada característica de apneia obstrutiva grave,",
         fontsize=7.5, color=FOOT, style="italic")
fig.text(0.025, 0.045,
         "com múltiplos despertares (eventos), N3 menor que 5% do tempo total, e ciclos REM truncados.",
         fontsize=7.5, color=FOOT, style="italic")
fig.text(0.025, 0.025,
         "O tempo de permanência na cama pode ser semelhante — o tempo de sono funcional, não.",
         fontsize=8, color=INK, weight="bold", style="italic")

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap14_Fig01.pdf"
png_path = out_dir / "_preview_Cap14_Fig01.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
