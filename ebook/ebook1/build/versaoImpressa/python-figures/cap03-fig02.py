"""
Cap03 Fig02 (PT-BR) — O Que Acelera e O Que Freia o Envelhecimento.

Reescrito em 2026-08-24 para reproduzir a COMPOSIÇÃO da arte original, e não
apenas o conteúdo. A versão anterior era do lote pré-v3: mesmos itens, mas
layout aproximado — gradiente como bloco vertical alto no lugar da barra fina,
coluna da direita alinhada à direita, sublinhado no lugar da barra de acento.

Todas as coordenadas abaixo foram MEDIDAS na arte original (1536×1024) por
detecção de cor e de tinta, não estimadas:

  título ................. y 0.040–0.089, x a partir de 0.026
  cabeçalhos ............. y 0.146–0.189; régua verde x 0.073–0.262 e
                           régua vermelha x 0.716–0.902, ambas em y 0.199
  barras de acento ....... x 0.054–0.056 (esq) e 0.699–0.701 (dir), 4 px
  separadores de item .... x 0.054–0.285 (esq) e 0.699–0.957 (dir)
  barra do gradiente ..... x 0.324–0.665, y 0.307–0.334
  MAIS LENTA / RÁPIDA .... y 0.271–0.286
  divisores pontilhados .. x 0.337 e 0.651, y 0.36–0.82
  tarja do rodapé ........ x 0.039–0.960, y 0.863–0.962

FIG_COR=1 gera na paleta da arte original em figuras-cor/; sem a variável sai
em cinza, em figuras-bw/, e a edição P&B não muda.
"""
import os
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
RULE     = "#C8C8C8"

COR         = os.environ.get("FIG_COR") == "1"
FREIA_COR   = "#085828" if COR else INK
ACELERA_COR = "#b00810" if COR else INK
LENTA_COR   = "#085018" if COR else INK_SOFT
RAPIDA_COR  = "#c80000" if COR else INK_SOFT
RODAPE_COR  = "#005018" if COR else INK
TARJA_COR   = "#EAF2EA" if COR else "#EDEDED"

# Perfil do gradiente lido da barra da arte original (linha y=315).
GRAD_STOPS = [(0.00, "#2d7a3c"), (0.10, "#3f8644"), (0.20, "#77aa6b"),
              (0.30, "#a7c796"), (0.40, "#d6e2c4"), (0.50, "#eeeade"),
              (0.60, "#fad3aa"), (0.70, "#f8ac73"), (0.80, "#f1813f"),
              (0.90, "#e34711"), (1.00, "#d92d09")]


def _grad(frac):
    if not COR:
        g = int(238 - frac * (238 - 150))
        return f"#{g:02X}{g:02X}{g:02X}"
    for i in range(len(GRAD_STOPS) - 1):
        f0, c0 = GRAD_STOPS[i]
        f1, c1 = GRAD_STOPS[i + 1]
        if f0 <= frac <= f1:
            k = 0 if f1 == f0 else (frac - f0) / (f1 - f0)
            a0 = [int(c0[j:j + 2], 16) for j in (1, 3, 5)]
            a1 = [int(c1[j:j + 2], 16) for j in (1, 3, 5)]
            return "#%02x%02x%02x" % tuple(
                round(a0[j] + (a1[j] - a0[j]) * k) for j in range(3))
    return GRAD_STOPS[-1][1]


# Aspecto 1.5, igual ao da arte original (1536×1024).
FIG_W, FIG_H = 11.0, 7.333
fig = plt.figure(figsize=(FIG_W, FIG_H))
fig.patch.set_facecolor(BG)

def Y(y_top):
    """Converte y medido a partir do TOPO da arte para fração de figura."""
    return 1.0 - y_top


# ---------- título ----------
fig.text(0.026, Y(0.070), "Figura 2 — O Que Acelera e O Que Freia o Envelhecimento",
         fontsize=18, color=INK, weight="bold", va="center")

# ---------- cabeçalhos das colunas ----------
RULE_L_X = (0.073, 0.262)
RULE_R_X = (0.716, 0.902)
for (x0, x1), texto, cor in ((RULE_L_X, "O QUE FREIA", FREIA_COR),
                             (RULE_R_X, "O QUE ACELERA", ACELERA_COR)):
    fig.text((x0 + x1) / 2, Y(0.165), texto,
             fontsize=17, color=cor, weight="bold", ha="center", va="center")
    fig.add_artist(plt.Line2D([x0, x1], [Y(0.199), Y(0.199)],
                              color=cor, linewidth=2.0,
                              transform=fig.transFigure, zorder=3))

# ---------- coluna central ----------
MID_CX = 0.494
fig.text(MID_CX, Y(0.213), "Velocidade do",
         fontsize=12, color=INK, weight="bold", ha="center", va="center")
fig.text(MID_CX, Y(0.246), "envelhecimento biológico",
         fontsize=12, color=INK, weight="bold", ha="center", va="center")

fig.text(0.365, Y(0.279), "MAIS LENTA",
         fontsize=9, color=LENTA_COR, weight="bold", ha="center", va="center")
fig.text(0.622, Y(0.279), "MAIS RÁPIDA",
         fontsize=9, color=RAPIDA_COR, weight="bold", ha="center", va="center")

# barra fina do gradiente, com o tique central escuro
BAR_X0, BAR_X1 = 0.324, 0.665
BAR_Y0, BAR_Y1 = Y(0.334), Y(0.307)
N = 180
for k in range(N):
    frac = k / (N - 1)
    bx = BAR_X0 + (BAR_X1 - BAR_X0) * (k / N)
    fig.patches.append(Rectangle(
        (bx, BAR_Y0), (BAR_X1 - BAR_X0) / N + 0.0008, BAR_Y1 - BAR_Y0,
        facecolor=_grad(frac), edgecolor="none",
        transform=fig.transFigure, zorder=1))
fig.add_artist(plt.Line2D([MID_CX, MID_CX], [BAR_Y0, BAR_Y1],
                          color="#252d35", linewidth=1.6,
                          transform=fig.transFigure, zorder=4))

# divisores pontilhados que separam a coluna central das laterais
for dx in (0.337, 0.651):
    fig.add_artist(plt.Line2D([dx, dx], [Y(0.820), Y(0.360)],
                              color=RULE, linewidth=0.9, linestyle=(0, (1, 4)),
                              transform=fig.transFigure, zorder=1))

fig.text(MID_CX, Y(0.588), "Processos biológicos",
         fontsize=10.5, color=INK_SOFT, ha="center", va="center", style="italic")
fig.text(MID_CX, Y(0.614), "moduláveis",
         fontsize=10.5, color=INK_SOFT, ha="center", va="center", style="italic")

# ---------- itens ----------
# (título, descrição em linhas) e o topo medido da barra de acento de cada item
items_left = [
    ("Exercício aeróbio",   ["Mitocôndrias ↑, telômeros ↑"]),
    ("Treino de força",     ["Inflamação ↓, insulina ↓"]),
    ("Sono de qualidade",   ["Limpeza celular ↑,", "epigenoma estável"]),
    ("Alimentação real",    ["Estresse oxidativo ↓,", "epigenoma regulado"]),
    ("Gestão do estresse",  ["Telômeros ↑, inflamação ↓"]),
]
items_right = [
    ("Sedentarismo",          ["Mitocôndrias ↓, senescência ↑"]),
    ("Ultraprocessados",      ["Inflamação ↑, estresse oxidativo ↑"]),
    ("Sono ruim / irregular", ["Limpeza celular ↓,", "epigenoma desregulado"]),
    ("Estresse crônico",      ["Telômeros ↓, cortisol ↑"]),
    ("Gordura visceral",      ["Inflamação ↑, insulina ↑"]),
]
BAR_TOP_L = [0.233, 0.349, 0.461, 0.610, 0.754]
BAR_BOT_L = [0.301, 0.415, 0.558, 0.706, 0.815]
BAR_TOP_R = [0.233, 0.349, 0.461, 0.616, 0.742]
BAR_BOT_R = [0.302, 0.416, 0.560, 0.680, 0.807]
SEP_L = [0.326, 0.438, 0.587, 0.731]
SEP_R = [0.326, 0.438, 0.588, 0.716]

def coluna(items, bar_x, texto_x, tops, bots, seps, sep_x, cor):
    for (nome, linhas), t, b in zip(items, tops, bots):
        fig.patches.append(Rectangle(
            (bar_x, Y(b)), 0.003, b - t,
            facecolor=cor, edgecolor="none",
            transform=fig.transFigure, zorder=3))
        fig.text(texto_x, Y(t + 0.006), nome,
                 fontsize=12.5, color=INK, weight="bold", ha="left", va="top")
        for j, ln in enumerate(linhas):
            fig.text(texto_x, Y(t + 0.040 + j * 0.028), ln,
                     fontsize=10.5, color=INK_SOFT, ha="left", va="top")
    for s in seps:
        fig.add_artist(plt.Line2D(list(sep_x), [Y(s), Y(s)],
                                  color=RULE, linewidth=0.8,
                                  transform=fig.transFigure, zorder=1))

coluna(items_left,  0.054, 0.073, BAR_TOP_L, BAR_BOT_L, SEP_L, (0.054, 0.285), FREIA_COR)
coluna(items_right, 0.699, 0.716, BAR_TOP_R, BAR_BOT_R, SEP_R, (0.699, 0.957), ACELERA_COR)

# ---------- tarja do rodapé ----------
fig.patches.append(Rectangle(
    (0.039, Y(0.962)), 0.960 - 0.039, 0.962 - 0.863,
    facecolor=TARJA_COR, edgecolor="none", transform=fig.transFigure, zorder=1))
fig.text(0.5, Y(0.913),
         "Envelhecimento não é destino. É um processo modificável.",
         fontsize=13, color=RODAPE_COR, weight="bold", ha="center", va="center")

# ---------- export ----------
out_dir = Path(__file__).resolve().parents[1] / ("figuras-cor" if COR else "figuras-bw")
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap03_Fig02.pdf"
png_path = out_dir / "_preview_Cap03_Fig02.png"
plt.savefig(pdf_path, facecolor=BG)
plt.savefig(png_path, dpi=170, facecolor=BG)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
