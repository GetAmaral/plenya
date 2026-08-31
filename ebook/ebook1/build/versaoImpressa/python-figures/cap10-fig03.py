"""
Cap10 Fig03 (PT-BR, B&W vetorial) — A janela dura cerca de 10 anos.

Timeline horizontal pós-menopausa (0-20 anos) com 3 zonas (PROTEÇÃO/DECISÃO/RISCO),
pill paciente Fernanda no topo, e 4 cards de estudos (ELITE precoce, KEEPS, ELITE
tardio, WHI) ancorados em pontos da linha. Anotação "~10 anos: limite da janela".

Posições medidas do original 1536×1024:
  year 0  → x=80
  year 20 → x=1375 (linear ~64.75 px/year)
  timeline horizontal em y=533
  zonas backgrounds em y=370..530 (acima da linha)
  pill paciente em y=200..305
"""
from pathlib import Path as _Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import (
    Rectangle, FancyBboxPatch, Polygon, Circle, FancyArrowPatch, Ellipse
)
import matplotlib.image as mpimg
import numpy as np

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False
rcParams["pdf.fonttype"] = 42
rcParams["ps.fonttype"] = 42

BG    = "#FFFFFF"
INK   = "#000000"
SOFT  = "#3A3A3A"
FOOT  = "#6A6A6A"
DARK  = "#2A2A2A"
GRAY1 = "#9E9E9E"
GRAY2 = "#D9D9D9"
GRAY3 = "#EEEEEE"
# Zonas B&W (hierarquia: zona PROTEÇÃO mais clara, RISCO mais escura)
ZONE_PROT  = "#F5F5F5"  # PROTEÇÃO (era verde)
ZONE_DEC   = "#EBEBEB"  # DECISÃO INDIVIDUAL (era amarelo)
ZONE_RISCO = "#E0E0E0"  # RISCO (era vermelho)

W_IMG, H_IMG = 1536, 1024
_FIG_W = 10.0
_FIG_H = _FIG_W * H_IMG / W_IMG
fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

ax_bg = fig.add_axes([0, 0, 1, 1])
ax_bg.set_xlim(0, W_IMG); ax_bg.set_ylim(H_IMG, 0)
ax_bg.set_aspect("equal"); ax_bg.axis("off")

def text_with_width(txt, x, y, **kw):
    t = ax_bg.text(x, y, txt, **kw)
    fig.canvas.draw()
    bbox = t.get_window_extent(renderer=fig.canvas.get_renderer())
    inv = ax_bg.transData.inverted()
    return inv.transform(bbox)[1][0]

# Timeline mapping — medido do original (year 0 tick em x=55, year 20 em x=1394)
TIME_LEFT_PX  = 55
TIME_RIGHT_PX = 1394
def year_to_px(y):
    return TIME_LEFT_PX + y / 20 * (TIME_RIGHT_PX - TIME_LEFT_PX)

# Linha do tempo em y 0.494 de 1024 px = 506 na arte original (era 533).
TIMELINE_Y = 506  # main horizontal line
# Topo da banda de zonas MEDIDO na arte original: y 0.302 de 1024 px = 309.
# Estava em 370, o que encurtava as três faixas e as descolava do cartão da
# Fernanda, que na arte encosta nelas.
ZONE_TOP   = 309
ZONE_BOT   = TIMELINE_Y  # zonas vão até a linha

# Zone boundaries — ALINHADOS com year ticks 10 e 15
ZONE_PROT_END = 10
ZONE_DEC_END  = 15

# ============================================================
# CABEÇALHO
# ============================================================
ax_bg.add_patch(Rectangle((25, 15), 175-25, 53-15, facecolor=INK, edgecolor="none"))
ax_bg.text((25+175)/2, (15+53)/2, "FIGURA 3",
           fontsize=11, color="white", weight="bold", va="center", ha="center")

ax_bg.text(41, 97, "A janela dura cerca de 10 anos.",
           fontsize=22, color=INK, weight="bold", va="center", ha="left")
ax_bg.text(43, 150,
           "Começar a reposição hormonal dentro dela protege. Começar fora pode fazer o oposto.",
           fontsize=11, color=SOFT, va="center", ha="left")

# ============================================================
# PILL PACIENTE FERNANDA (no topo, com seta pra baixo)
# ============================================================
pill_x, pill_y, pill_w, pill_h = 75, 195, 605, 115
# Rounded rect pill
ax_bg.add_patch(FancyBboxPatch(
    (pill_x+6, pill_y+6), pill_w-12, pill_h-12,
    boxstyle="round,pad=4,rounding_size=10",
    facecolor=BG, edgecolor=INK, linewidth=1.6, zorder=10))

# Person icon — extraído do original (female silhouette outline)
ICON_DIR = _Path(__file__).resolve().parents[1] / "figuras-bw" / "_icons"
icon_img = mpimg.imread(str(ICON_DIR / "fernanda_t.png"))
ih, iw = icon_img.shape[:2]
# Target width 76 px (similar ao original)
target_w = 62
target_h = target_w * ih / iw
icx, icy = pill_x + 47, pill_y + pill_h/2
extent = (icx - target_w/2, icx + target_w/2,
          icy + target_h/2, icy - target_h/2)
ax_bg.imshow(icon_img, extent=extent, zorder=11)

# Texto pill
TX = pill_x + 115
ax_bg.text(TX, pill_y + 32, "Fernanda, 44 anos",
           fontsize=15, color=INK, weight="bold", va="center", ha="left", zorder=11)
ax_bg.text(TX, pill_y + 60, "FSH 38, estradiol 28 pg/mL.",
           fontsize=12, color=SOFT, va="center", ha="left", zorder=11)
ax_bg.text(TX, pill_y + 88, "Na janela.",
           fontsize=13, color=INK, weight="bold", va="center", ha="left", zorder=11)

# Triangle (arrow down from pill) — posição medida year 1.7 (x=188)
tri_top = pill_y + pill_h - 2
tri_bot = pill_y + pill_h + 22
tri_cx = year_to_px(1.7)
ax_bg.add_patch(Polygon([
    (tri_cx-10, tri_top), (tri_cx+10, tri_top), (tri_cx, tri_bot)
], closed=True, facecolor=INK, edgecolor="none", zorder=10))
# Dashed line vertical from arrow tip DOWN through zone até a timeline
ax_bg.plot([tri_cx, tri_cx], [tri_bot + 4, TIMELINE_Y - 8],
           linestyle=(0, (3, 3)), color=INK, linewidth=1.2, zorder=4)
# Pequeno dot ON timeline marcando posição da Fernanda
ax_bg.add_patch(Circle((tri_cx, TIMELINE_Y), 8,
                       facecolor=INK, edgecolor="none", zorder=6))

# ============================================================
# 3 ZONAS — bandas horizontais
# ============================================================
zone_specs = [
    (0, ZONE_PROT_END, ZONE_PROT,  "PROTEÇÃO",          [
        "Estradiol transdérmico reduz",
        "progressão de aterosclerose,",
        "protege osso e preserva",
        "função cognitiva.",
    ]),
    (ZONE_PROT_END, ZONE_DEC_END, ZONE_DEC, "DECISÃO INDIVIDUAL", [
        "Evidência menos clara.",
        "Decisão caso a caso com",
        "estratificação de risco.",
    ]),
    (ZONE_DEC_END, 21, ZONE_RISCO, "RISCO ↑", [
        "Aterosclerose estabelecida",
        "torna o estradiol potencialmente",
        "deletério. Considerar alternativas",
        "não-hormonais.",
    ]),
]
for y0, y1, color, lbl, desc in zone_specs:
    x0_px = year_to_px(y0)
    x1_px = year_to_px(y1)
    ax_bg.add_patch(Rectangle(
        (x0_px, ZONE_TOP), x1_px - x0_px, ZONE_BOT - ZONE_TOP,
        facecolor=color, edgecolor="none", zorder=1
    ))

# Dividers tracejados entre zonas
for x_year in (ZONE_PROT_END, ZONE_DEC_END):
    x_px = year_to_px(x_year)
    ax_bg.plot([x_px, x_px], [ZONE_TOP, ZONE_BOT],
               linestyle=(0, (4, 4)), color=GRAY1, linewidth=0.8, zorder=2)

# Labels das zonas
ZONE_LBL_Y = 393
ZONE_DESC_START_Y = 425
DESC_LH = 22

zones_centers_years = [(0+ZONE_PROT_END)/2, (ZONE_PROT_END+ZONE_DEC_END)/2, (ZONE_DEC_END+21)/2]
for (yz_center, (y0, y1, color, lbl, desc)) in zip(zones_centers_years, zone_specs):
    cx = year_to_px(yz_center)
    ax_bg.text(cx, ZONE_LBL_Y, lbl,
               fontsize=13, color=INK, weight="bold",
               ha="center", va="center", zorder=6)
    for i, line in enumerate(desc):
        ax_bg.text(cx, ZONE_DESC_START_Y + i * DESC_LH, line,
                   fontsize=10, color=INK,
                   ha="center", va="center", zorder=6)

# ============================================================
# TIMELINE HORIZONTAL + TICKS + LABELS
# ============================================================
# Main timeline line
ax_bg.plot([TIME_LEFT_PX-20, TIME_RIGHT_PX+50], [TIMELINE_Y, TIMELINE_Y],
           color=INK, linewidth=2.0, zorder=5)
# Arrow head at right end
ax_bg.add_patch(Polygon([
    (TIME_RIGHT_PX+50, TIMELINE_Y-8),
    (TIME_RIGHT_PX+50, TIMELINE_Y+8),
    (TIME_RIGHT_PX+70, TIMELINE_Y),
], closed=True, facecolor=INK, edgecolor="none", zorder=5))

# Ticks at 0, 5, 10, 15, 20
TICK_YEARS = [0, 5, 10, 15, 20]
for year in TICK_YEARS:
    x_px = year_to_px(year)
    ax_bg.plot([x_px, x_px], [TIMELINE_Y - 8, TIMELINE_Y + 8],
               color=INK, linewidth=1.8, zorder=5)
    # tick label below
    ax_bg.text(x_px, TIMELINE_Y + 25, str(year),
               fontsize=12, color=INK, weight="bold",
               ha="center", va="center", zorder=10)

# Axis label em 2 LINHAS centralizado com year 10
ax_label_cx = year_to_px(10)
ax_bg.text(ax_label_cx, TIMELINE_Y + 65, "Anos desde a última",
           fontsize=10, color=INK, weight="bold",
           ha="center", va="center")
ax_bg.text(ax_label_cx, TIMELINE_Y + 84, "menstruação",
           fontsize=10, color=INK, weight="bold",
           ha="center", va="center")

# Dashed arrow short — apenas no final do timeline (medido x=1379..1499)
DASH_Y = TIMELINE_Y - 38
dash_x_start = 1379
dash_x_end = 1499
ax_bg.plot([dash_x_start, dash_x_end], [DASH_Y, DASH_Y],
           linestyle=(0, (8, 5)), color=INK, linewidth=1.4, zorder=4)
ax_bg.add_patch(Polygon([
    (dash_x_end, DASH_Y-5), (dash_x_end, DASH_Y+5), (dash_x_end+10, DASH_Y)
], closed=True, facecolor=INK, edgecolor="none", zorder=4))

# Callout "~10 anos: limite da janela" — maior pra ter padding mínimo
callout_cx = year_to_px(10)
callout_y = TIMELINE_Y + 145
CALLOUT_W = 160
CALLOUT_H = 64
# Triangle pointer no topo do callout, apontando pra axis label
ax_bg.add_patch(Polygon([
    (callout_cx - 7, callout_y - CALLOUT_H/2),
    (callout_cx + 7, callout_y - CALLOUT_H/2),
    (callout_cx, callout_y - CALLOUT_H/2 - 8)
], closed=True, facecolor=ZONE_DEC, edgecolor=INK, linewidth=0.8, zorder=8))
ax_bg.add_patch(FancyBboxPatch(
    (callout_cx - CALLOUT_W/2, callout_y - CALLOUT_H/2),
    CALLOUT_W, CALLOUT_H,
    boxstyle="round,pad=2,rounding_size=6",
    facecolor=ZONE_DEC, edgecolor=INK, linewidth=1.0, zorder=8))
ax_bg.text(callout_cx, callout_y - 13, "~10 anos:",
           fontsize=10, color=INK, weight="bold",
           ha="center", va="center", zorder=9)
ax_bg.text(callout_cx, callout_y + 13, "limite da janela",
           fontsize=10, color=INK,
           ha="center", va="center", zorder=9)

# ============================================================
# STUDY CARDS — 4 cards posicionados pra não colidir
# Cards anchors mais espaçados + cards mais estreitos
# Anchor → onde o dot fica na timeline (medido visualmente do original)
# ============================================================
# Posições MEDIDAS do original (centers em ano):
#   ELITE precoce: card_center=3.4 (anchor=5)
#   KEEPS: card_center=6.6 (anchor=7)
#   ELITE tardio: card_center=11.9 (anchor=11)
#   WHI: card_center=14.6 (anchor=13.5)
# ELITE precoce e KEEPS: leaders RETOS (dot e card no mesmo x)
# ELITE tardio e WHI: leaders TORTOS (dot offset)
CARDS = [
    {"year": 3.67, "dot_year": 3.67, "w": 174,
     "name": "ELITE (2016)", "subname": "grupo precoce",
     "text": ["Reduziu progressão", "de aterosclerose."],
     "style": "ok"},
    {"year": 6.69, "dot_year": 6.69, "w": 172,
     "name": "KEEPS", "subname": "(2024 follow-up)",
     "text": ["Confirmou a hipótese", "da janela em", "seguimento longo."],
     "style": "ok"},
    {"year": 12.6, "dot_year": 12.04, "w": 150,
     "name": "ELITE (2016)", "subname": "grupo tardio",
     "text": ["Não reduziu; em", "alguns recortes,", "piorou."],
     "style": "bad"},
    {"year": 15.2, "dot_year": 13.14, "w": 152,
     "name": "WHI (2002)", "subname": "",
     "text": ["Média de 63 anos,", "10+ anos", "pós-menopausa.",
              "Assustou uma", "geração inteira."],
     "style": "neutral"},
]

CARD_LH = 16
CARD_TOP = 662   # medido do original (KEEPS card top border)

for card in CARDS:
    x_card = year_to_px(card["year"])
    x_dot = year_to_px(card["dot_year"])
    card_y_top = CARD_TOP
    CARD_W = card["w"]

    if card["style"] == "ok":
        dot_face = INK
    elif card["style"] == "bad":
        dot_face = INK
    else:
        dot_face = GRAY1
    # Dot BELOW the timeline (gap ~14px) — measured from original
    dot_y = TIMELINE_Y + 27
    ax_bg.add_patch(Circle((x_dot, dot_y), 9,
                           facecolor=dot_face, edgecolor="none", zorder=6))
    # Leader em L: vertical do dot, depois bend horizontal/diagonal pro card top
    if abs(x_dot - x_card) < 5:
        # Same x — straight vertical
        ax_bg.plot([x_dot, x_dot],
                   [dot_y + 9, card_y_top],
                   color=dot_face, linewidth=1.4, zorder=5)
    else:
        # L-shape: vertical 60% down, then bend to card x
        bend_y = dot_y + 9 + (card_y_top - dot_y - 9) * 0.55
        ax_bg.plot([x_dot, x_dot, x_card, x_card],
                   [dot_y + 9, bend_y, bend_y, card_y_top],
                   color=dot_face, linewidth=1.4, zorder=5,
                   solid_joinstyle="round", solid_capstyle="round")
    x_card_center = x_card

    # Card height baseada em conteúdo
    n_text = len(card["text"])
    if card["subname"]:
        card_h = 30 + 22 + n_text * CARD_LH + 25
    else:
        card_h = 30 + n_text * CARD_LH + 25
    card_y_bot = card_y_top + card_h

    edge_color = INK if card["style"] != "neutral" else GRAY1
    line_width = 1.6 if card["style"] == "ok" else (
                  2.0 if card["style"] == "bad" else 1.0)

    ax_bg.add_patch(FancyBboxPatch(
        (x_card_center - CARD_W/2 + 4, card_y_top + 4),
        CARD_W - 8, card_h - 8,
        boxstyle="round,pad=3,rounding_size=8",
        facecolor=BG, edgecolor=edge_color, linewidth=line_width, zorder=8
    ))

    ax_bg.text(x_card_center, card_y_top + 22, card["name"],
               fontsize=10, color=INK, weight="bold",
               ha="center", va="center", zorder=9)
    if card["subname"]:
        ax_bg.text(x_card_center, card_y_top + 42, card["subname"],
                   fontsize=8.5, color=INK, weight="bold",
                   ha="center", va="center", zorder=9)
        divider_y = card_y_top + 60
    else:
        divider_y = card_y_top + 40
    # Soft divider entre title e body (medido no original)
    ax_bg.plot([x_card_center - CARD_W/2 + 14, x_card_center + CARD_W/2 - 14],
               [divider_y, divider_y],
               color=GRAY1, linewidth=0.7, alpha=0.5, zorder=9)
    text_start_y = divider_y + 18
    for i, line in enumerate(card["text"]):
        ax_bg.text(x_card_center, text_start_y + i * CARD_LH, line,
                   fontsize=7, color=SOFT,
                   ha="center", va="center", zorder=9)

# ============================================================
# SOURCE
# ============================================================
ax_bg.plot([45, W_IMG-45], [905, 905], color=GRAY1, linewidth=0.7)

src_x = 50
src_y1 = 928
end_b = text_with_width("Fontes:", src_x, src_y1,
                        fontsize=9, color=FOOT, weight="bold",
                        va="center", ha="left")
ax_bg.text(end_b + 4, src_y1,
           "Hodis et al., ELITE, NEJM 2016; Harman et al., KEEPS primary, Annals 2014;",
           fontsize=9, color=FOOT, va="center", ha="left", style="italic")
ax_bg.text(src_x, src_y1 + 22,
           "Miller et al., KEEPS continuation, PLOS Medicine 2024; Rossouw et al., WHI primary, JAMA 2002.",
           fontsize=9, color=FOOT, va="center", ha="left", style="italic")

# ============================================================
# EXPORT
# ============================================================
out_dir = _Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap10_Fig03.pdf"
png_path = out_dir / "_preview_Cap10_Fig03.png"
plt.savefig(pdf_path, facecolor=BG, bbox_inches="tight", pad_inches=0.0)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches="tight", pad_inches=0.0)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
