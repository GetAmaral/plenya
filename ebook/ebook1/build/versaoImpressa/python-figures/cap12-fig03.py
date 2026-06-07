"""
Cap12 Fig03 (PT-BR, B&W vetorial) — Cinco instrumentos que mudam a consulta em cinco minutos.

Implementação BASEADA EM MEDIÇÕES PIXEL-PRECISAS do original 2048x3072.
Doc canônico: /tmp/cap12f03_canonical.md

Mapeamento de cores → B&W:
  - Vermelho (FIGURA 3, alert box, ≥10 cut) → PRETO
  - Navy/azul (GAD-7, sidebar border) → PRETO
  - Cores instrumento-específicas (orange/maroon/green) → PRETO
"""
from pathlib import Path as _Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, FancyBboxPatch, Circle, Polygon
import math as _m

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False
rcParams["pdf.fonttype"] = 42
rcParams["ps.fonttype"] = 42

BG    = "#FFFFFF"
INK   = "#000000"
SOFT  = "#5A5A5A"
FOOT  = "#6A6A6A"
GRAY1 = "#9E9E9E"
GRAY2 = "#D9D9D9"
ZONE  = "#EDEDED"

W_IMG, H_IMG = 2048, 3072
_FIG_W = 18.0
_FIG_H = _FIG_W * H_IMG / W_IMG
fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

ax = fig.add_axes([0, 0, 1, 1])
ax.set_xlim(0, W_IMG); ax.set_ylim(H_IMG, 0)
ax.set_aspect("equal"); ax.axis("off")

# ============================================================
# CABEÇALHO
# ============================================================
# FIGURA 3 box (medido: x=55..315, y=44..116)
ax.add_patch(Rectangle((55, 44), 260, 72, facecolor=INK, edgecolor="none"))
ax.text(55 + 130, 44 + 36, "FIGURA 3",
        fontsize=22, color="white", weight="bold",
        ha="center", va="center")

# Título serif bold — 3 linhas (medido y=168..279, 286..374, 408..493)
ax.text(60, 224, "Cinco instrumentos que",
        fontsize=58, color=INK, weight="bold", family="serif",
        va="center", ha="left")
ax.text(60, 330, "mudam a consulta em",
        fontsize=58, color=INK, weight="bold", family="serif",
        va="center", ha="left")
ax.text(60, 450, "cinco minutos.",
        fontsize=58, color=INK, weight="bold", family="serif",
        va="center", ha="left")

# Subtítulo gray — 2 linhas (medido y=538..579, 623..663)
ax.text(60, 560, "O mínimo necessário para identificar quando",
        fontsize=26, color=SOFT, va="center", ha="left")
ax.text(60, 643, "o plano deixa de ser biológico.",
        fontsize=26, color=SOFT, va="center", ha="left")

# ============================================================
# SIDEBAR DIREITA (medido: x=1405..1981, y=728..2829)
# ============================================================
SB_X0, SB_X1 = 1405, 1981
SB_Y0, SB_Y1 = 728, 2829

ax.add_patch(FancyBboxPatch(
    (SB_X0, SB_Y0), SB_X1 - SB_X0, SB_Y1 - SB_Y0,
    boxstyle="round,pad=4,rounding_size=30",
    facecolor=BG, edgecolor=INK, linewidth=2.2, zorder=1))

sb_cx = (SB_X0 + SB_X1) / 2

# Star badge no topo (circle preto + estrela branca dentro)
star_cy = SB_Y0 + 110
ax.add_patch(Circle((sb_cx, star_cy), 58,
                     facecolor=INK, edgecolor="none", zorder=3))
star_pts = []
for i in range(10):
    angle = -_m.pi/2 + i * _m.pi / 5
    rr = 32 if i % 2 == 0 else 14
    star_pts.append((sb_cx + rr * _m.cos(angle), star_cy + rr * _m.sin(angle)))
ax.add_patch(Polygon(star_pts, closed=True, facecolor="white",
                      edgecolor="none", zorder=4))

# Título do sidebar 4 linhas centralizadas
sb_y = star_cy + 130
sb_title_lines = ["Quando o", "psicológico", "vira prioridade", "do plano:"]
for ln in sb_title_lines:
    ax.text(sb_cx, sb_y, ln,
            fontsize=34, color=INK, weight="bold",
            ha="center", va="center")
    sb_y += 64

# Divisor
sb_y += 20
ax.plot([SB_X0 + 60, SB_X1 - 60], [sb_y, sb_y],
        color=INK, linewidth=1.2)
sb_y += 75

# 3 bullets numerados
SB_BULLET_CX = SB_X0 + 95
bullets = [
    ["PCR persistente", "> 1,5 apesar de", "pilares biológicos", "otimizados."],
    ["Pontuação", "≥ ponto de corte", "em qualquer", "escala acima."],
    ["Ideação suicida", "ou trauma", "relatados em", "qualquer", "intensidade."],
]
for i, lines in enumerate(bullets, 1):
    ax.add_patch(Circle((SB_BULLET_CX, sb_y + 8), 30,
                         facecolor=INK, edgecolor="none", zorder=3))
    ax.text(SB_BULLET_CX, sb_y + 8, str(i),
            fontsize=24, color="white", weight="bold",
            ha="center", va="center", zorder=4)
    text_x = SB_BULLET_CX + 55
    text_y = sb_y + 8
    for ln in lines:
        ax.text(text_x, text_y, ln,
                fontsize=22, color=INK,
                ha="left", va="center")
        text_y += 48
    sb_y = text_y + 30

# Bottom callout
bb_top = SB_Y1 - 220
bb_h = 180
ax.add_patch(FancyBboxPatch(
    (SB_X0 + 30, bb_top), SB_X1 - SB_X0 - 60, bb_h,
    boxstyle="round,pad=4,rounding_size=16",
    facecolor=ZONE, edgecolor="none", zorder=2))
bottom_lines = ["Identificar cedo muda",
                "a direção do plano",
                "e o desfecho do paciente."]
by2 = bb_top + 50
for ln in bottom_lines:
    ax.text(sb_cx, by2, ln,
            fontsize=22, color=INK, weight="bold",
            ha="center", va="center")
    by2 += 45

# ============================================================
# INSTRUMENTOS — 5 escalas
# ============================================================
INSTR_BAR_X0, INSTR_BAR_X1 = 80, 86  # barra MUITO fina (6 unidades = ~9 px no original)
BADGE_CX = 170
NAME_X = 240
AXIS_X0, AXIS_X1 = 557, 1315  # linha visível

INSTRUMENTS = [
    {
        "num": "1", "name": "PHQ-9", "title": "Depressão",
        "desc_lines": ["Sintomas depressivos", "nas últimas 2 semanas."],
        "bar_y0": 725, "bar_y1": 1295,
        "axis_y": 835,
        # value → x (medido)
        "ticks": [(0, 566), (5, 738), (10, 915), (15, 1111), (20, 1310)],
        "cuts": [10],
        "categories": [(4.5, ["leve"]),
                       (12.5, ["moderada"]),
                       (17, ["moderada-", "severa"]),
                       (20, ["severa"])],
        "cut_n": "≥ 10",
        "cut_text": "entra no plano agora.",
        "alert": "Qualquer resposta positiva sobre ideação suicida = conversa imediata, antes do paciente sair.",
        "alert_y0": 1008, "alert_y1": 1297,
    },
    {
        "num": "2", "name": "GAD-7", "title": "Ansiedade generalizada",
        "desc_lines": ["Sintomas de ansiedade", "nas últimas 2 semanas."],
        "bar_y0": 1338, "bar_y1": 1674,
        "axis_y": 1447,
        "ticks": [(0, 567), (5, 738), (10, 914), (15, 1091), (21, 1304)],
        "cuts": [10],
        "categories": [(4.5, ["leve"]), (12.5, ["moderada"]), (18, ["severa"])],
        "cut_n": "≥ 10",
        "cut_text": "ansiedade relevante.",
    },
    {
        "num": "3", "name": "AUDIT", "title": "Consumo problemático de álcool",
        "desc_lines": ["Padrões de uso de álcool."],
        "bar_y0": 1710, "bar_y1": 2069,
        "axis_y": 1810,
        "ticks": [(0, 567), (8, 745), (16, 952), (40, 1303)],
        "cuts": [8, 16],
        "categories": [(12, ["risco"]), (28, ["uso nocivo /", "dependência"])],
        "cut_n": "≥ 8",
        "cut_text": "risco mesmo em uso social.",
        "note": "Sobretudo em paciente que diz “só bebo socialmente”.",
    },
    {
        "num": "4", "name": "PCL-5", "title": "Transtorno de estresse pós-traumático",
        "desc_lines": ["Sintomas relacionados", "a evento traumático."],
        "bar_y0": 2102, "bar_y1": 2447,
        "axis_y": 2206,
        "ticks": [(0, 567), (33, 923), (80, 1302)],
        "cuts": [33],
        "categories": [],
        "cut_n": "≥ 33",
        "cut_text": "investigar trauma.",
        "note": "Aplicar se a pergunta de trauma na anamnese foi positiva.",
    },
    {
        "num": "5", "name": "UCLA-3", "title": "Solidão percebida",
        "desc_lines": ["Percepção subjetiva de", "conexão e pertencimento."],
        "bar_y0": 2479, "bar_y1": 2784,
        "axis_y": 2572,
        "ticks": [(0, 567), (3, 781), (6, 1023), (9, 1305)],
        "cuts": [6],
        "categories": [(1.5, ["baixa"]), (7.5, ["elevada"])],
        "cut_n": "≥ 6",
        "cut_text": "solidão relevante.",
        "note": "Mesmo em paciente que diz ter “muitos amigos”.",
    },
]


def value_to_x(v, ticks):
    """Interpola x para valor v usando ticks como pontos âncora."""
    pairs = ticks
    if v <= pairs[0][0]:
        return pairs[0][1]
    if v >= pairs[-1][0]:
        return pairs[-1][1]
    for i in range(len(pairs)-1):
        v0, x0 = pairs[i]
        v1, x1 = pairs[i+1]
        if v0 <= v <= v1:
            frac = (v - v0) / (v1 - v0)
            return x0 + frac * (x1 - x0)
    return pairs[-1][1]


for instr in INSTRUMENTS:
    bar_y0, bar_y1 = instr["bar_y0"], instr["bar_y1"]
    axis_y = instr["axis_y"]

    # Barra vertical preta (instrumento)
    ax.add_patch(Rectangle(
        (INSTR_BAR_X0, bar_y0),
        INSTR_BAR_X1 - INSTR_BAR_X0,
        bar_y1 - bar_y0,
        facecolor=INK, edgecolor="none", zorder=2))

    # Badge numerado (círculo preto com número branco)
    badge_cy = bar_y0 + 45
    ax.add_patch(Circle((BADGE_CX, badge_cy), 40,
                         facecolor=INK, edgecolor="none", zorder=3))
    ax.text(BADGE_CX, badge_cy, instr["num"],
            fontsize=30, color="white", weight="bold",
            ha="center", va="center", zorder=4)

    # Name (bold)
    ax.text(NAME_X, badge_cy, instr["name"],
            fontsize=36, color=INK, weight="bold",
            ha="left", va="center")

    # Title (abaixo do name) — wrap se for muito longo
    title_y = bar_y0 + 95
    title_text = instr["title"]
    # Wrap titles longos
    title_max_chars = 18  # cabe na coluna ~330px
    if len(title_text) > title_max_chars:
        words = title_text.split()
        lines, cur = [], ""
        for w in words:
            if len(cur) + len(w) + 1 > title_max_chars:
                lines.append(cur)
                cur = w
            else:
                cur = (cur + " " + w).strip()
        if cur: lines.append(cur)
    else:
        lines = [title_text]
    for ln in lines:
        ax.text(NAME_X, title_y, ln,
                fontsize=22, color=INK,
                ha="left", va="center")
        title_y += 36

    # Description (gray, 1 ou 2 linhas)
    desc_y = title_y + 14
    for line in instr["desc_lines"]:
        ax.text(NAME_X, desc_y, line,
                fontsize=18, color=SOFT, ha="left", va="center")
        desc_y += 32

    # --- Axis ---
    ax.plot([AXIS_X0, AXIS_X1], [axis_y, axis_y],
            color=GRAY1, linewidth=1.8)
    cut_set = set(instr["cuts"])
    for tv, tx in instr["ticks"]:
        # Tick mark
        ax.plot([tx, tx], [axis_y - 9, axis_y + 9],
                color=GRAY1, linewidth=1.5)
        is_cut = tv in cut_set
        ax.text(tx, axis_y - 40, f"{tv:g}",
                fontsize=26 if is_cut else 22,
                color=INK if is_cut else SOFT,
                weight="bold" if is_cut else "normal",
                ha="center", va="center")

    # Markers (cut values) — círculos pretos preenchidos sobre o eixo
    for cv in instr["cuts"]:
        cx = value_to_x(cv, instr["ticks"])
        ax.add_patch(Circle((cx, axis_y), 18,
                             facecolor=INK, edgecolor="none", zorder=4))

    # Categorias abaixo do eixo
    n_cat_max = max((len(cl) for _, cl in instr["categories"]), default=0)
    for cat_v, cat_lines in instr["categories"]:
        cxm = value_to_x(cat_v, instr["ticks"])
        cy = axis_y + 56
        for cln in cat_lines:
            ax.text(cxm, cy, cln,
                    fontsize=20, color=SOFT, ha="center", va="center")
            cy += 32

    # Cut label "→ ≥ N → texto"
    cl_y = axis_y + 56 + max(n_cat_max, 1) * 32 + 28
    # → (arrow)
    ax.text(AXIS_X0 + 8, cl_y, "→",
            fontsize=30, color=INK, weight="bold",
            ha="left", va="center")
    # ≥ N (bold)
    ax.text(AXIS_X0 + 80, cl_y, instr["cut_n"],
            fontsize=30, color=INK, weight="bold",
            ha="left", va="center")
    # → arrow segundo
    ax.text(AXIS_X0 + 220, cl_y, "→",
            fontsize=30, color=INK, weight="bold",
            ha="left", va="center")
    # texto
    ax.text(AXIS_X0 + 285, cl_y, instr["cut_text"],
            fontsize=28, color=INK,
            ha="left", va="center")

    # Note (italic gray)
    if instr.get("note"):
        ax.text(AXIS_X0 + 8, cl_y + 50, instr["note"],
                fontsize=20, color=SOFT, style="italic",
                ha="left", va="center")

    # Alert box (PHQ-9 only)
    if instr.get("alert"):
        ay0 = instr["alert_y0"]
        ay1 = instr["alert_y1"]
        ax_x0 = 562
        ax_x1 = 1244
        box_w = ax_x1 - ax_x0
        box_h = ay1 - ay0
        ax.add_patch(FancyBboxPatch(
            (ax_x0, ay0), box_w, box_h,
            boxstyle="round,pad=4,rounding_size=14",
            facecolor=INK, edgecolor="none", zorder=2))
        # White circle com "!" preto
        ic_cx = ax_x0 + 70
        ic_cy = (ay0 + ay1) / 2
        ax.add_patch(Circle((ic_cx, ic_cy), 38,
                             facecolor="white", edgecolor="none", zorder=3))
        ax.text(ic_cx, ic_cy, "!",
                fontsize=42, color=INK, weight="bold",
                ha="center", va="center", zorder=4)
        # Texto branco (2 linhas)
        alert_lines = [
            "Qualquer resposta positiva sobre ideação suicida =",
            "conversa imediata, antes do paciente sair.",
        ]
        ty = ay0 + 100
        for ln in alert_lines:
            ax.text(ax_x0 + 140, ty, ln,
                    fontsize=24, color="white",
                    ha="left", va="center", zorder=4)
            ty += 60

# ============================================================
# FOOTER — FONTES (abaixo de y=2820)
# ============================================================
foot_div_y = 2860
ax.plot([60, W_IMG - 60], [foot_div_y, foot_div_y],
        color=GRAY1, linewidth=1.0)
foot_lines = [
    "Fontes: PHQ-9 (Kroenke et al., JGIM 2001); GAD-7 (Spitzer et al., Arch Intern Med 2006);",
    "AUDIT (Saunders et al., Addiction 1993); PCL-5 (Weathers et al., NCPTSD 2013);",
    "UCLA-3 (Hughes et al., Research on Aging 2004). Todas validadas em português brasileiro.",
]
fy = foot_div_y + 50
for ln in foot_lines:
    ax.text(60, fy, ln,
            fontsize=20, color=FOOT,
            va="center", ha="left", style="italic")
    fy += 42

# ============================================================
# EXPORT
# ============================================================
out_dir = _Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap12_Fig03.pdf"
png_path = out_dir / "_preview_Cap12_Fig03.png"
plt.savefig(pdf_path, facecolor=BG, pad_inches=0.0)
plt.savefig(png_path, dpi=170, facecolor=BG, pad_inches=0.0)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
