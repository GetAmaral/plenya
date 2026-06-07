"""
Cap09 Fig01 (PT-BR, B&W vetorial) — Três sistemas, uma medicina só.

Diagrama de Venn com 3 círculos INTERSECTANTES (Coração, Rim, Metabolismo).
Caixa central no overlap 3-vias lista as 3 classes farmacológicas com estudos.
Caixa lateral direita ("Antes/Hoje") com leader pontilhado.

Ícones de órgãos extraídos do original e convertidos pra silhueta B&W
(via _icons/heart_t.png, kidney_t.png, metab_t.png).

Implementação: usa ax.text() sequencial para bold+regular inline (NÃO usa
mathtext que renderiza espaços extras e converte hyphens em minus signs).
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
GRAY1 = "#9E9E9E"
GRAY2 = "#D9D9D9"
GRAY3 = "#EEEEEE"

W_IMG, H_IMG = 1536, 1024
_FIG_W = 10.0
_FIG_H = _FIG_W * H_IMG / W_IMG
fig = plt.figure(figsize=(_FIG_W, _FIG_H))
fig.patch.set_facecolor(BG)

ax = fig.add_axes([0, 0, 1, 1])
ax.set_xlim(0, W_IMG)
ax.set_ylim(H_IMG, 0)
ax.set_aspect("equal")
ax.axis("off")

# --- HELPER: text bbox em data coords (pra encadear bold + regular) ---
def text_with_width(txt, x, y, **kw):
    """Render text and return data-coord right edge (for next inline element)."""
    t = ax.text(x, y, txt, **kw)
    fig.canvas.draw()
    bbox = t.get_window_extent(renderer=fig.canvas.get_renderer())
    inv = ax.transData.inverted()
    right = inv.transform(bbox)[1][0]
    return right

# ============================================================
# CABEÇALHO
# ============================================================
ax.add_patch(Rectangle((24, 13), 191-24, 57-13, facecolor=INK, edgecolor="none"))
ax.text((24+191)/2, (13+57)/2, "FIGURA 1",
        fontsize=11, color="white", weight="bold", va="center", ha="center")

ax.text(45, 105,
        "Três sistemas, uma medicina só.",
        fontsize=22, color=INK, weight="bold", va="center", ha="left")

ax.text(45, 160,
        "Por que três classes de medicamentos deixaram de ser ‘remédio de diabético’ nos últimos cinco anos.",
        fontsize=11, color=SOFT, va="center", ha="left")

# ============================================================
# DIAGRAMA DE VENN — 3 CÍRCULOS INTERSECTANTES
# ============================================================
# Posições calibradas pra match visual do original (medidas em bbox 1536x1024)
# Originais (Cor=(479,442), Rim=(934,442), Met=(706,638) com R~270)
# Ajusto pra R menor + posições proporcionais pra caber no rodapé do source
R_VENN = 250
C_COR  = (470, 440)
C_RIM  = (895, 440)
C_MET  = (685, 660)

FILL_ALPHA = 0.16
for center in (C_COR, C_RIM, C_MET):
    ax.add_patch(Circle(center, R_VENN, facecolor=INK, edgecolor="none",
                        alpha=FILL_ALPHA, zorder=1))
for center in (C_COR, C_RIM, C_MET):
    ax.add_patch(Circle(center, R_VENN, facecolor="none", edgecolor=INK,
                        linewidth=1.6, zorder=2))

# ============================================================
# ÍCONES + LABELS DOS ÓRGÃOS
# ============================================================
ICON_DIR = _Path(__file__).resolve().parents[1] / "figuras-bw" / "_icons"

def place_icon(name, cx, cy, target_w):
    img = mpimg.imread(str(ICON_DIR / f"{name}_t.png"))
    h, w = img.shape[:2]
    scale = target_w / w
    ext_w, ext_h = target_w, h * scale
    extent = (cx - ext_w/2, cx + ext_w/2,
              cy + ext_h/2, cy - ext_h/2)
    ax.imshow(img, extent=extent, zorder=5)

# Coração — icon GRANDE pra ficar igual em tamanho ao original (~140 px wide)
ICON_OFFSET_X_COR = -60
# Heart icon menor (110 wide, h=110) e posicionado mais alto pra não tocar label
place_icon("heart",  C_COR[0]+ICON_OFFSET_X_COR, C_COR[1]-160, target_w=110)
ax.text(C_COR[0]+ICON_OFFSET_X_COR, C_COR[1]-65, "Coração",
        fontsize=16, color=INK, weight="bold",
        va="center", ha="center", zorder=6)
ax.plot([C_COR[0]+ICON_OFFSET_X_COR-70, C_COR[0]+ICON_OFFSET_X_COR+70],
        [C_COR[1]-47, C_COR[1]-47],
        color=INK, linewidth=2.0, zorder=6)

# Rim — icon GRANDE
ICON_OFFSET_X_RIM = 60
place_icon("kidney", C_RIM[0]+ICON_OFFSET_X_RIM, C_RIM[1]-140, target_w=130)
ax.text(C_RIM[0]+ICON_OFFSET_X_RIM, C_RIM[1]-50, "Rim",
        fontsize=16, color=INK, weight="bold",
        va="center", ha="center", zorder=6)
ax.plot([C_RIM[0]+ICON_OFFSET_X_RIM-35, C_RIM[0]+ICON_OFFSET_X_RIM+35],
        [C_RIM[1]-32, C_RIM[1]-32],
        color=INK, linewidth=2.0, zorder=6)

# Metabolismo — icon GRANDE embaixo
place_icon("metab", C_MET[0], C_MET[1]+90, target_w=200)
ax.text(C_MET[0], C_MET[1]+170, "Metabolismo",
        fontsize=16, color=INK, weight="bold",
        va="center", ha="center", zorder=6)
ax.plot([C_MET[0]-100, C_MET[0]+100], [C_MET[1]+185, C_MET[1]+185],
        color=INK, linewidth=2.0, zorder=6)

# ============================================================
# CAIXA CENTRAL — 3 classes farmacológicas (no centro do Venn)
# Largura aumentada pra não cortar texto
# ============================================================
# Caixa central: largura ajustada pra caber "(dapagliflozina · empagliflozina)"
cb_x, cb_y, cb_w, cb_h = 490, 370, 400, 350
ax.add_patch(FancyBboxPatch(
    (cb_x+6, cb_y+6), cb_w-12, cb_h-12,
    boxstyle="round,pad=4,rounding_size=10",
    facecolor=BG, edgecolor=INK, linewidth=1.8, zorder=10))

ENTRIES = [
    ("1", "SGLT2",
     "(dapagliflozina · empagliflozina)",
     "DAPA-HF · EMPA-KIDNEY · DELIVER"),
    ("2", "Finerenona",
     "(Firialta)",
     "FIDELIO-DKD · FIGARO-DKD · FIND-CKD"),
    ("3", "GLP-1 / GIP-GLP-1",
     "(semaglutida · tirzepatida)",
     "SELECT (2023) · retatrutida fase 2"),
]

entry_h = 105
for i, (num, title, paren, studies) in enumerate(ENTRIES):
    ey = cb_y + 28 + i*entry_h
    bx = cb_x + 32
    ax.add_patch(Circle((bx, ey + 12), 14, facecolor=INK,
                        edgecolor="none", zorder=11))
    # Number centered via baseline alignment + slight y nudge
    ax.text(bx, ey + 12, num,
            fontsize=12, color="white", weight="bold",
            va="center_baseline", ha="center", zorder=12)
    ax.text(bx + 30, ey + 6, title,
            fontsize=12.5, color=INK, weight="bold",
            va="center", ha="left", zorder=11)
    ax.text(bx + 30, ey + 32, paren,
            fontsize=9.5, color=SOFT, va="center", ha="left", zorder=11)
    ax.text(bx + 30, ey + 53, studies,
            fontsize=8, color=SOFT, style="italic",
            va="center", ha="left", zorder=11)
    if i < len(ENTRIES) - 1:
        ax.plot([cb_x + 18, cb_x + cb_w - 18],
                [ey + entry_h - 18, ey + entry_h - 18],
                color=GRAY1, linewidth=0.6, zorder=10)

# ============================================================
# CAIXA LATERAL DIREITA — Antes / Hoje
# Posicionada mais à direita pra não tocar caixa central
# ============================================================
# Side note: REAL posição do box no original (não confundir com leader)
# Box: x=1130..1420 (w=290), y=510..760 (h=250). Leader vai do canto sup-dir
# da caixa central diagonalmente até o canto sup-esq do side note.
sb_x, sb_y, sb_w, sb_h = 1130, 510, 290, 250
ax.add_patch(FancyBboxPatch(
    (sb_x+6, sb_y+6), sb_w-12, sb_h-12,
    boxstyle="round,pad=4,rounding_size=10",
    facecolor=BG, edgecolor=INK, linewidth=1.6, zorder=10))

# Leader sai do MEIO VERTICAL da borda direita da caixa central
LEADER_START = (cb_x + cb_w, cb_y + cb_h/2)
ax.add_patch(Circle(LEADER_START, 5, facecolor=INK,
                    edgecolor="none", zorder=12))
# Linha até o meio vertical da borda esquerda do side note
ax.plot([LEADER_START[0], sb_x],
        [LEADER_START[1], sb_y + sb_h/2],
        color=INK, linewidth=1.4, zorder=11)

# Side note: 2 seções "Antes" e "Hoje" separadas por divider, como no original
TXT_X = sb_x + 22
# Antes — título bold no topo, texto regular abaixo (2 linhas, fontsize 11)
y1 = sb_y + 30
ax.text(TXT_X, y1, "Antes:",
        fontsize=12, color=INK, weight="bold",
        va="center", ha="left", zorder=11)
ax.text(TXT_X, y1 + 24, "remédio de diabético.",
        fontsize=11, color=INK, va="center", ha="left", zorder=11)

# Divider entre Antes e Hoje
ax.plot([sb_x+18, sb_x+sb_w-18], [sb_y+82, sb_y+82],
        color=GRAY1, linewidth=0.7, zorder=11)

# Hoje — título bold, depois 4 linhas regulares com "sem diabetes." bold
y2 = sb_y + 105
end_b2 = text_with_width("Hoje:", TXT_X, y2,
                         fontsize=12, color=INK, weight="bold",
                         va="center", ha="left", zorder=11)
ax.text(end_b2 + 6, y2, "proteção",
        fontsize=11, color=INK, va="center", ha="left", zorder=11)
ax.text(TXT_X, y2 + 24, "cardiorrenal +",
        fontsize=11, color=INK, va="center", ha="left", zorder=11)
ax.text(TXT_X, y2 + 48, "metabólica em",
        fontsize=11, color=INK, va="center", ha="left", zorder=11)
# Última linha: "paciente " regular + "sem diabetes." bold
end_pac = text_with_width("paciente ", TXT_X, y2 + 72,
                          fontsize=11, color=INK,
                          va="center", ha="left", zorder=11)
ax.text(end_pac, y2 + 72, "sem diabetes.",
        fontsize=11, color=INK, weight="bold",
        va="center", ha="left", zorder=11)

# ============================================================
# SOURCE (rodapé) — bold "Fonte:" + texto regular com nomes em italic
# ============================================================
ax.plot([45, W_IMG-45], [905, 905], color=GRAY1, linewidth=0.7)

# Helper pra adicionar trecho regular ou italic sequencial
def src_chain(x, y, segments, fontsize=9.5):
    """segments = list of (text, style, weight). Renderiza em cadeia."""
    cur_x = x
    for text, style, weight in segments:
        right = text_with_width(text, cur_x, y,
                                fontsize=fontsize, color=FOOT,
                                weight=weight, style=style,
                                va="center", ha="left")
        cur_x = right

# Linha 1: Fonte: ensaios DAPA-HF (...), EMPA-KIDNEY (...), DELIVER (...),
src_chain(45, 935, [
    ("Fonte:", "normal", "bold"),
    (" ensaios ", "normal", "normal"),
    ("DAPA-HF", "italic", "normal"),
    (" (McMurray, NEJM 2019), ", "normal", "normal"),
    ("EMPA-KIDNEY", "italic", "normal"),
    (" (Herrington, NEJM 2023), ", "normal", "normal"),
    ("DELIVER", "italic", "normal"),
    (" (Solomon, NEJM 2022),", "normal", "normal"),
])
# Linha 2: FIDELIO-DKD (...), FIGARO-DKD (...), FIND-CKD (...), SELECT (...)
src_chain(45, 960, [
    ("FIDELIO-DKD", "italic", "normal"),
    (" (Bakris, NEJM 2020), ", "normal", "normal"),
    ("FIGARO-DKD", "italic", "normal"),
    (" (Pitt, NEJM 2021), ", "normal", "normal"),
    ("FIND-CKD", "italic", "normal"),
    (" (2025), ", "normal", "normal"),
    ("SELECT", "italic", "normal"),
    (" (Lincoff, NEJM 2023).", "normal", "normal"),
])
ax.text(45, 990, "Síntese clínica do Capítulo 9.",
        fontsize=9.5, color=FOOT, va="center", ha="left")

# ============================================================
# EXPORT
# ============================================================
out_dir = _Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap09_Fig01.pdf"
png_path = out_dir / "_preview_Cap09_Fig01.png"
plt.savefig(pdf_path, facecolor=BG, bbox_inches="tight", pad_inches=0.0)
plt.savefig(png_path, dpi=170, facecolor=BG, bbox_inches="tight", pad_inches=0.0)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
