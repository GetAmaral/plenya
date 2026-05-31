"""
Cap14 Fig03 (PT-BR, B&W vetorial) — Paulo: quatro tempos, quatro pilares, uma tese.
Trajetória 24 meses, 6 biomarcadores × 4 tempos.
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
COL_BG   = "#F4F4F4"
ROW_ALT  = "#FAFAFA"
BAND     = "#EDEDED"

_FIG_W, _FIG_H = 11.0, 7.8
_ASPECT = _FIG_W / _FIG_H

fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

LEFT = 0.025

# Título
fig.text(LEFT, 0.945,
         "Figura 3 — Paulo: quatro tempos, quatro pilares, uma tese.",
         fontsize=15, color=INK, weight="bold")
fig.text(LEFT, 0.910,
         "Trajetória de 24 meses mostrando a regressão silenciosa e o resgate pós-CPAP.",
         fontsize=9.5, color=INK_SOFT, style="italic")

# Header da tabela: 5 colunas
LBL_COL_X = 0.025
LBL_COL_W = 0.18
ALVO_X    = 0.215
T_COL_X = [0.34, 0.50, 0.66, 0.82]   # T0, T+6m, T+18m, T+24m

# Header texts
COL_HEADERS = [
    ("T0\nInício (Cap. 6)",          "baseline"),
    ("T+6m\nAlvo aparente",          "antes do CPAP"),
    ("T+18m\nRegressão silenciosa",  "antes do CPAP"),
    ("T+24m\nCPAP + protocolo\ncircadiano",  "6 m pós-falha"),
]
HEAD_Y = 0.840
for x, (head, sub) in zip(T_COL_X, COL_HEADERS):
    fig.text(x, HEAD_Y, head,
             fontsize=9, color=INK, weight="bold", ha="center",
             linespacing=1.15)
fig.text(LBL_COL_X, HEAD_Y, "BIOMARCADOR",
         fontsize=9, color=INK_SOFT, weight="bold")
fig.text(ALVO_X, HEAD_Y, "ALVO ÓTIMO",
         fontsize=9, color=INK_SOFT, weight="bold", ha="center")

# Linha horizontal abaixo dos headers
fig.lines.append(plt.Line2D(
    [LBL_COL_X - 0.005, 0.97], [0.770, 0.770],
    color="#888888", linewidth=0.6, transform=fig.transFigure
))

# 6 biomarcadores
biomarkers = [
    # (name, alvo, t0, t6, t18, t24)
    ("Testosterona total",        "> 500",   "210",       "485",      "410",       "540"),
    ("Testosterona livre",        "—",       "4,8",       "11,2",     "6,4",       "12,5"),
    ("PCR ultrassensível",        "< 1,0",   "1,7",       "0,9",      "1,4",       "0,5"),
    ("HbA1c",                     "5,4–5,6", "5,8",       "5,6",      "5,7",       "5,4"),
    ("N3 (% do tempo total)",     "15–20%",  "não medido aqui", "<5%", "<5%",      "17%"),
    ("IAH (apneia–hipopneia)",    "< 5",     "—",         "—",        "22",        "2"),
]

ROW_TOP = 0.745
ROW_H = 0.090

ASPECT = 11.0 / 7.8

for ri, (name, alvo, t0, t6, t18, t24) in enumerate(biomarkers):
    y = ROW_TOP - ri * ROW_H
    # Background alterno
    if ri % 2 == 1:
        fig.patches.append(Rectangle(
            (LBL_COL_X - 0.005, y - ROW_H/2 + 0.005), 0.97 - LBL_COL_X + 0.005, ROW_H - 0.005,
            facecolor=ROW_ALT, edgecolor="none",
            transform=fig.transFigure, zorder=0
        ))

    # Nome
    fig.text(LBL_COL_X, y + 0.005, name,
             fontsize=9, color=INK, weight="bold", va="center")
    # Alvo
    fig.text(ALVO_X, y + 0.005, alvo,
             fontsize=9, color=INK_SOFT, ha="center", va="center")

    # 4 valores
    for x, val in zip(T_COL_X, [t0, t6, t18, t24]):
        # Bolinha + valor
        fig.patches.append(Ellipse(
            (x - 0.040, y + 0.005),
            width=0.010, height=0.010 * ASPECT,
            facecolor=INK, edgecolor=INK, linewidth=0.5,
            transform=fig.transFigure, zorder=3
        ))
        fig.text(x - 0.025, y + 0.005, val,
                 fontsize=9, color=INK, weight="bold", va="center")

# ---------- Annotations / setas explicativas no topo ----------
ANNOT_Y = 0.795
fig.text(T_COL_X[2], ANNOT_Y,
         "Regressão silenciosa:\no que nenhum check-up\nconvencional capta",
         fontsize=7.5, color=INK_SOFT, ha="center", va="bottom",
         style="italic", linespacing=1.2)

fig.text(T_COL_X[3], ANNOT_Y,
         "CPAP + protocolo\ncircadiano = a peça\nque faltava",
         fontsize=7.5, color=INK_SOFT, ha="center", va="bottom",
         style="italic", linespacing=1.2)

# ---------- Callout footer ----------
BOX_X1, BOX_X2 = 0.025, 0.97
BOX_Y1, BOX_Y2 = 0.115, 0.180

fig.patches.append(FancyBboxPatch(
    (BOX_X1, BOX_Y1), BOX_X2 - BOX_X1, BOX_Y2 - BOX_Y1,
    boxstyle="round,pad=0.005,rounding_size=0.005",
    facecolor=BAND, edgecolor=INK, linewidth=0.5,
    transform=fig.transFigure, zorder=1
))
fig.text(0.5, (BOX_Y1 + BOX_Y2)/2 + 0.012,
         "Este é o fechamento do sistema.",
         fontsize=11, color=INK, weight="bold", ha="center", va="center")
fig.text(0.5, (BOX_Y1 + BOX_Y2)/2 - 0.012,
         "Cada peça depende da outra — o N3 é o que faltava.",
         fontsize=10, color=INK, ha="center", va="center", style="italic")

# Footer
fig.text(LEFT, 0.085,
         "Trajetória completa do Paulo ao longo de 24 meses. Os Capítulos 6 e 10 mostram como o painel pode estar em \"alvo\" sem mover o gargalo real.",
         fontsize=7.5, color=FOOT, style="italic")
fig.text(LEFT, 0.068,
         "A polissonografia revelou apneia obstrutiva grave (IAH = 22) — sem tratamento, otimização nutrológica e nem reposição hormonal teriam fechado a equação.",
         fontsize=7.5, color=FOOT, style="italic")
fig.text(LEFT, 0.048,
         "O CPAP + protocolo circadiano, em conjunto, fecharam apenas os marcadores que ainda voltavam ao primeiro alvo — superaram-no.",
         fontsize=7.5, color=FOOT, style="italic")
fig.text(LEFT, 0.025,
         "Sistemas integrados precisam de fechamento integral. E é por isso que o sono é, de fato, o primeiro pilar.",
         fontsize=8, color=INK, weight="bold", style="italic")

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap14_Fig03.pdf"
png_path = out_dir / "_preview_Cap14_Fig03.png"
# Tight crop sem padding lateral/topo; 0.08" no inferior pra não colar na legenda.
from matplotlib.transforms import Bbox as _Bbox
fig.canvas.draw()
_tb = fig.get_tightbbox(fig.canvas.get_renderer())  # já em inches
_bbox_in = _Bbox.from_extents(_tb.x0, _tb.y0 - 0.08, _tb.x1, _tb.y1)
plt.savefig(pdf_path, facecolor=BG, bbox_inches=_bbox_in)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches=_bbox_in)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
