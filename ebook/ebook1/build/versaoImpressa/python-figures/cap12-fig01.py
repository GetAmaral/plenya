"""
Cap12 Fig01 (PT-BR) — Como a ansiedade vira doença: a cascata do eixo HPA.

Reescrito em 2026-08-24 reproduzindo a COMPOSIÇÃO da arte original. A versão
anterior era do lote pré-v3: mesmo conteúdo, mas quatro caixas brancas de
borda fina no lugar dos quatro painéis de fundo sólido, e proporção 0.78 contra
0.667 do original.

Coordenadas MEDIDAS na arte original (2048×3072), y a partir do topo:
  painéis ....... x 0.087–0.913
                  1  y 0.097–0.244   navy   #18426c
                  2  y 0.303–0.443   cinza  claro
                  3  y 0.501–0.664   âmbar  #fce4ae
                  4  y 0.718–0.889   salmão #de6c5a
  setas ......... x 0.480–0.518, y 0.248–0.294 / 0.447–0.492 / 0.670–0.713
  caixas p2 ..... x 0.150–0.333 / 0.407–0.587 / 0.663–0.853, y 0.374–0.426
  caixas p4 ..... x 0.110–0.305 / 0.316–0.495 / 0.505–0.705 / 0.716–0.890,
                  y 0.788–0.842

FIG_COR=1 gera na paleta original em figuras-cor/; sem a variável sai em cinza,
em figuras-bw/.
"""
import os
from pathlib import Path

import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, Circle, FancyBboxPatch, Polygon

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False
rcParams["pdf.fonttype"] = 42
rcParams["ps.fonttype"] = 42

COR = os.environ.get("FIG_COR") == "1"

BG   = "#FFFFFF"
INK  = "#000000"
SOFT = "#3A3A3A"
FOOT = "#555555"

# Painéis: cor na arte original, cinza equivalente em P&B (mantendo o contraste
# claro→escuro que faz a cascata ser lida de cima para baixo).
P1_BG = "#18426c" if COR else "#3A3A3A"
P2_BG = "#F2F2F2"
P3_BG = "#fce4ae" if COR else "#E4E4E4"
P4_BG = "#de6c5a" if COR else "#B4B4B4"
P1_FG, P4_FG = "#FFFFFF", "#FFFFFF"
BADGE3 = "#b8801e" if COR else "#6E6E6E"
SETA   = "#8A8A8A"
SETINHA = "#18426c" if COR else "#3A3A3A"
UP_DOWN = "#c0392b" if COR else INK          # ↑ / ↓ do painel 3

FIG_W, FIG_H = 8.0, 12.0                     # aspecto 0.667, igual ao original
fig = plt.figure(figsize=(FIG_W, FIG_H))
fig.patch.set_facecolor(BG)
ASP = FIG_W / FIG_H

def Y(y_top):
    return 1.0 - y_top

PX0, PX1 = 0.087, 0.913


def painel(y0, y1, cor, borda=None):
    fig.patches.append(FancyBboxPatch(
        (PX0, Y(y1)), PX1 - PX0, y1 - y0,
        boxstyle="round,pad=0,rounding_size=0.008",
        facecolor=cor, edgecolor=borda or cor, linewidth=1.0,
        transform=fig.transFigure, zorder=1))


def badge(n, y_top, face, texto_cor):
    cx = 0.132
    r = 0.019
    fig.patches.append(Circle((cx, Y(y_top)), r, facecolor=face,
                              edgecolor=face, transform=fig.transFigure,
                              zorder=3))
    fig.text(cx, Y(y_top), str(n), fontsize=13, color=texto_cor,
             weight="bold", ha="center", va="center", zorder=4)


def seta(y0, y1, linhas):
    """Seta grossa para baixo, com o rótulo em itálico à direita."""
    cx = 0.499
    w = 0.017
    fig.patches.append(Rectangle((cx - w/2, Y(y1 - 0.013)), w, (y1 - y0) - 0.013,
                                 facecolor=SETA, edgecolor="none",
                                 transform=fig.transFigure, zorder=2))
    fig.patches.append(Polygon([[cx - 0.020, Y(y1 - 0.013)],
                                [cx + 0.020, Y(y1 - 0.013)],
                                [cx, Y(y1)]],
                               closed=True, facecolor=SETA, edgecolor="none",
                               transform=fig.transFigure, zorder=2))
    for i, ln in enumerate(linhas):
        fig.text(0.545, Y(y0 + 0.012 + i * 0.014), ln, fontsize=10.5,
                 color=SOFT, style="italic", ha="left", va="center")


# ---------- cabeçalho ----------
fig.text(PX0, Y(0.023), "FIGURA 1", fontsize=10, color=FOOT,
         weight="bold", va="center")
fig.text(PX0, Y(0.053), "Como a ansiedade vira doença: a cascata do eixo HPA",
         fontsize=19, color=INK, weight="bold", va="center")

# ---------- painel 1 ----------
painel(0.097, 0.244, P1_BG)
badge(1, 0.130, "#FFFFFF", P1_BG)
fig.text(0.5, Y(0.130), "ESTÍMULO PSICOLÓGICO", fontsize=13, color=P1_FG,
         weight="bold", ha="center", va="center")
fig.text(0.5, Y(0.163), "Estresse crônico, ruminação, ansiedade",
         fontsize=16, color=P1_FG, weight="bold", ha="center", va="center")
fig.add_artist(plt.Line2D([0.128, 0.872], [Y(0.186), Y(0.186)],
                          color=P1_FG, linewidth=0.8, alpha=0.55,
                          transform=fig.transFigure, zorder=3))
p1_itens = [("E-mails após", "o expediente"), ("Preocupação", "com filhos"),
            ("Pressão", "financeira")]
p1_cx = [0.278, 0.500, 0.722]
for cx, (l1, l2) in zip(p1_cx, p1_itens):
    fig.text(cx, Y(0.208), l1, fontsize=11.5, color=P1_FG,
             ha="center", va="center")
    fig.text(cx, Y(0.227), l2, fontsize=11.5, color=P1_FG,
             ha="center", va="center")
for dx in (0.389, 0.611):
    fig.add_artist(plt.Line2D([dx, dx], [Y(0.232), Y(0.199)],
                              color=P1_FG, linewidth=0.7, alpha=0.45,
                              transform=fig.transFigure, zorder=3))

seta(0.248, 0.294, ["Percepção de ameaça", "pelo cérebro"])

# ---------- painel 2 ----------
painel(0.303, 0.443, P2_BG, borda="#DCDCDC")
badge(2, 0.334, "#6E6E6E", "#FFFFFF")
fig.text(0.5, Y(0.334), "EIXO HPA ATIVADO", fontsize=13, color=INK,
         weight="bold", ha="center", va="center")
fig.text(0.5, Y(0.360), "Eixo Hipotálamo–Hipófise–Adrenal em alerta permanente",
         fontsize=11.5, color=SOFT, ha="center", va="center")
p2_boxes = [(0.150, 0.333, "Hipotálamo", "→ CRH"),
            (0.407, 0.587, "Hipófise", "→ ACTH"),
            (0.663, 0.853, "Suprarrenais", "→ Cortisol")]
for x0, x1, l1, l2 in p2_boxes:
    fig.patches.append(FancyBboxPatch(
        (x0, Y(0.426)), x1 - x0, 0.426 - 0.374,
        boxstyle="round,pad=0,rounding_size=0.006",
        facecolor="#FFFFFF", edgecolor="#C8C8C8", linewidth=0.9,
        transform=fig.transFigure, zorder=2))
    fig.text((x0 + x1) / 2, Y(0.391), l1, fontsize=11.5, color=INK,
             ha="center", va="center", zorder=3)
    fig.text((x0 + x1) / 2, Y(0.411), l2, fontsize=11.5, color=INK,
             ha="center", va="center", zorder=3)
for xa in (0.348, 0.604):
    fig.patches.append(Polygon(
        [[xa, Y(0.393)], [xa, Y(0.407)], [xa + 0.030, Y(0.400)]],
        closed=True, facecolor="#8A8A8A", edgecolor="none",
        transform=fig.transFigure, zorder=3))

seta(0.447, 0.492, ["Cortisol cronicamente", "elevado"])

# ---------- painel 3 ----------
painel(0.501, 0.664, P3_BG)
badge(3, 0.532, BADGE3, "#FFFFFF")
fig.text(0.5, Y(0.532), "CONSEQUÊNCIAS BIOQUÍMICAS", fontsize=13, color=INK,
         weight="bold", ha="center", va="center")
fig.text(0.5, Y(0.560), "Cascata bioquímica", fontsize=12.5, color=INK,
         weight="bold", ha="center", va="center")
esq = [("Citocinas inflamatórias ", "↑", "(IL-6, TNF-α, PCR)"),
       ("Resistência insulínica ", "↑", None),
       ("Gordura visceral ", "↑", None)]
dire = [("Hormônios sexuais ", "↓", None),
        ("Função imune ", "↓", None),
        ("Hipocampo / memória ", "↓", None)]
def _largura(txt, fontsize, weight="normal"):
    """Largura de `txt` em fração da figura, medida com o renderer."""
    probe = fig.text(0, 0, txt, fontsize=fontsize, weight=weight)
    fig.canvas.draw()
    w = probe.get_window_extent(renderer=fig.canvas.get_renderer()).width
    probe.remove()
    return w / fig.bbox.width


def bullets(itens, x, ys):
    """Bullet + rótulo + a seta ↑/↓ logo depois do texto.

    A posição da seta é MEDIDA: preencher com espaços não funciona, porque a
    largura do espaço não corresponde à das letras e a seta acaba caindo em
    cima da palavra.

    `ys` traz o y de cada LINHA (contando a continuação como linha própria),
    nas posições medidas na arte original — a coluna esquerda tem 4 linhas de
    passo uniforme e a direita tem 3 itens mais espaçados, ambas terminando em
    0.650, dentro do painel.
    """
    i = 0
    for txt, seta_c, cont in itens:
        y = ys[i]; i += 1
        fig.text(x, Y(y), "•", fontsize=12, color=INK, ha="left", va="center")
        fig.text(x + 0.016, Y(y), txt, fontsize=12, color=INK,
                 ha="left", va="center")
        fig.text(x + 0.016 + _largura(txt, 12), Y(y), seta_c, fontsize=12,
                 color=UP_DOWN, weight="bold", ha="left", va="center")
        if cont:
            fig.text(x + 0.016, Y(ys[i]), cont, fontsize=12, color=INK,
                     ha="left", va="center")
            i += 1
bullets(esq, 0.132, [0.585, 0.607, 0.629, 0.650])
bullets(dire, 0.545, [0.585, 0.617, 0.650])
fig.add_artist(plt.Line2D([0.508, 0.508], [Y(0.648), Y(0.582)],
                          color=BADGE3, linewidth=0.9, alpha=0.7,
                          transform=fig.transFigure, zorder=3))

seta(0.670, 0.713, ["Anos a décadas", "de exposição"])

def cabe(texto, largura_frac, base, minimo=9.0, folga=0.90):
    """Maior corpo <= base em que ``texto`` cabe em ``largura_frac`` da figura.

    Mede de verdade (TextPath na fonte já configurada) em vez de chutar largura
    média de caractere — rótulo de caixa que estoura a borda é pior que rótulo
    pequeno.
    """
    from matplotlib.textpath import TextPath
    from matplotlib.font_manager import FontProperties
    alvo = largura_frac * fig.get_size_inches()[0] * 72 * folga
    tam = base
    while tam > minimo:
        largura = TextPath((0, 0), texto, size=tam,
                           prop=FontProperties(family=rcParams["font.sans-serif"])
                           ).get_extents().width
        if largura <= alvo:
            break
        tam -= 0.25
    return round(tam, 2)


# ---------- painel 4 ----------
painel(0.718, 0.889, P4_BG)
badge(4, 0.748, "#FFFFFF", P4_BG)
fig.text(0.5, Y(0.748), "DESFECHOS CLÍNICOS", fontsize=13, color=P4_FG,
         weight="bold", ha="center", va="center")
fig.text(0.5, Y(0.775), "As quatro doenças crônicas do Capítulo 2",
         fontsize=15, color=P4_FG, weight="bold", ha="center", va="center")
# "Neurodegeneração" era quebrada à mão em ["Neuro-", "degeneração"]: virava a
# única palavra hifenizada da arte e saltava na página impressa. Agora vai
# inteira e o corpo encolhe até caber na caixa (nunca abaixo de 9pt).
p4_boxes = [(0.110, 0.305, ["Doença", "cardiovascular"]),
            (0.316, 0.495, ["Doença", "metabólica"]),
            (0.505, 0.705, ["Neurodegeneração"]),
            (0.716, 0.890, ["Câncer"])]
for x0, x1, linhas in p4_boxes:
    fig.patches.append(FancyBboxPatch(
        (x0, Y(0.842)), x1 - x0, 0.842 - 0.788,
        boxstyle="round,pad=0,rounding_size=0.006",
        facecolor="#FDF1EE" if COR else "#F4F4F4", edgecolor="none",
        transform=fig.transFigure, zorder=2))
    cy = 0.815 - (0.010 if len(linhas) > 1 else 0)
    for j, ln in enumerate(linhas):
        fig.text((x0 + x1) / 2, Y(cy + j * 0.019), ln,
                 fontsize=cabe(ln, x1 - x0, 11.5), color=INK,
                 ha="center", va="center", zorder=3)
fig.text(0.5, Y(0.858), "+ envelhecimento biológico acelerado",
         fontsize=12.5, color=P4_FG, weight="bold", style="italic",
         ha="center", va="center")
fig.text(0.5, Y(0.876), "(telômeros, epigenoma)", fontsize=11.5,
         color=P4_FG, style="italic", ha="center", va="center")

# ---------- rodapé ----------
fig.add_artist(plt.Line2D([PX0, PX1], [Y(0.903), Y(0.903)],
                          color="#C8C8C8", linewidth=0.8,
                          transform=fig.transFigure))
rodape = [
    "A ativação crônica do eixo hipotálamo–hipófise–adrenal converte estresse psicológico sustentado em",
    "consequências biológicas mensuráveis — e essas consequências são exatamente os mecanismos discutidos",
    "ao longo deste livro como drivers das doenças crônicas do envelhecimento. Por isso a Integração Corpo-Mente",
    "não é pilar complementar — é pilar constitutivo.",
]
for i, ln in enumerate(rodape):
    fig.text(PX0, Y(0.921 + i * 0.017), ln, fontsize=9.2, color=SOFT,
             ha="left", va="center")

out_dir = Path(__file__).resolve().parents[1] / ("figuras-cor" if COR else "figuras-bw")
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap12_Fig01.pdf"
png_path = out_dir / "_preview_Cap12_Fig01.png"
plt.savefig(pdf_path, facecolor=BG)
plt.savefig(png_path, dpi=150, facecolor=BG)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
