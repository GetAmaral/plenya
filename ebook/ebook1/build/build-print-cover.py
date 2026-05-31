#!/usr/bin/env python3
"""Build print-ready cover for 'ANTES' — single-PDF wrap for KDP Paperback (6x9").

Usage:
    python3 build-print-cover.py pt-BR <pagecount>
    python3 build-print-cover.py pt-BR 406

Layout (left → right, when laid flat):
    [bleed] [BACK COVER (6")] [SPINE (variable)] [FRONT COVER (6")] [bleed]

Spine width = pagecount × 0.002252 inches (KDP white paper 75g formula).
Bleed = 0.125 inches on top/bottom and outside edges.
Resolution: 300 DPI in CMYK preview, RGB final (KDP accepts both).
"""

from __future__ import annotations
import sys
from pathlib import Path
from PIL import Image, ImageDraw, ImageFont

BOOK_ROOT = Path(__file__).resolve().parent.parent
LANG = sys.argv[1] if len(sys.argv) > 1 else "pt-BR"
PAGECOUNT = int(sys.argv[2]) if len(sys.argv) > 2 else 406

# --- Paths ---
CAPA_FRONT  = BOOK_ROOT / "capas" / LANG / "capa.jpg"
BARCODE     = BOOK_ROOT / "capas" / LANG / "isbn-barcode-978-65-02-07691-0.png"
AUTHOR_BW   = BOOK_ROOT / "fotos" / "getulio_bw_halfbody_1000.jpg"
SYMBOL_GOLD = BOOK_ROOT.parent.parent / "apps" / "site" / "public" / "brand" / "symbol" / "gold.png"
WORDMARK_GOLD = BOOK_ROOT.parent.parent / "apps" / "site" / "public" / "brand" / "wordmark" / "gold.png"
BUILD_DIR   = BOOK_ROOT / "build"
OUT_PDF     = BUILD_DIR / f"Antes-{LANG}-print-cover-{PAGECOUNT}pp.pdf"
OUT_PNG     = BUILD_DIR / f"Antes-{LANG}-print-cover-{PAGECOUNT}pp.png"

# --- Geometry ---
DPI = 300
TRIM_W_IN = 6.0
TRIM_H_IN = 9.0
BLEED_IN = 0.125
SPINE_IN = PAGECOUNT * 0.002252

CANVAS_W_IN = BLEED_IN + TRIM_W_IN + SPINE_IN + TRIM_W_IN + BLEED_IN
CANVAS_H_IN = BLEED_IN + TRIM_H_IN + BLEED_IN

CANVAS_W = round(CANVAS_W_IN * DPI)
CANVAS_H = round(CANVAS_H_IN * DPI)

# Boundaries (px)
BLEED_PX = round(BLEED_IN * DPI)                   # 38
TRIM_PX = round(TRIM_W_IN * DPI)                   # 1800
SPINE_PX = round(SPINE_IN * DPI)
HEIGHT_PX = round(TRIM_H_IN * DPI)                 # 2700

# X coordinates of major regions
BACK_X0  = BLEED_PX                                # back cover left edge (after bleed)
BACK_X1  = BACK_X0 + TRIM_PX                       # back cover right edge = spine left
SPINE_X0 = BACK_X1
SPINE_X1 = SPINE_X0 + SPINE_PX                     # spine right edge = front cover left
FRONT_X0 = SPINE_X1
FRONT_X1 = FRONT_X0 + TRIM_PX
TOP_PX   = BLEED_PX
BOT_PX   = TOP_PX + HEIGHT_PX

# --- Plenya palette ---
PETROL = (6, 59, 79)         # #063b4f
GOLD   = (179, 134, 69)      # #b38645
CREAM  = (234, 231, 218)     # #eae7da
INK    = (26, 26, 26)        # #1a1a1a

# --- Fonts ---
def load_fonts():
    candidates_serif = [
        "/usr/share/texmf/fonts/opentype/public/tex-gyre/texgyrepagella-italic.otf",
        "/usr/share/fonts/truetype/dejavu/DejaVuSerif-Italic.ttf",
    ]
    candidates_serif_regular = [
        "/usr/share/texmf/fonts/opentype/public/tex-gyre/texgyrepagella-regular.otf",
        "/usr/share/fonts/truetype/dejavu/DejaVuSerif.ttf",
    ]
    candidates_sans_bold = [
        "/usr/share/fonts/opentype/inter/Inter-Bold.otf",
        "/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf",
    ]
    candidates_sans = [
        "/usr/share/fonts/opentype/inter/Inter-Regular.otf",
        "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf",
    ]
    def first(paths):
        for p in paths:
            if Path(p).exists():
                return p
        raise FileNotFoundError(f"No font found among {paths}")
    return {
        "serif_regular": first(candidates_serif_regular),
        "serif_italic":  first(candidates_serif),
        "sans_bold":     first(candidates_sans_bold),
        "sans":          first(candidates_sans),
    }

FONTS = load_fonts()

def f(role, size_pt):
    """Load font at given point size for our DPI."""
    px = round(size_pt * DPI / 72)
    return ImageFont.truetype(FONTS[role], px)


# --- Helpers ---
def fit_into_box(img: Image.Image, max_w: int, max_h: int) -> Image.Image:
    img = img.copy()
    img.thumbnail((max_w, max_h), Image.LANCZOS)
    return img


def circular_crop(img: Image.Image, diameter: int) -> Image.Image:
    """Crop to square center, resize to diameter, mask to circle (RGBA)."""
    w, h = img.size
    side = min(w, h)
    left = (w - side) // 2
    top = (h - side) // 2
    img = img.crop((left, top, left + side, top + side)).resize((diameter, diameter), Image.LANCZOS)
    if img.mode != "RGBA":
        img = img.convert("RGBA")
    mask = Image.new("L", (diameter, diameter), 0)
    ImageDraw.Draw(mask).ellipse((0, 0, diameter, diameter), fill=255)
    out = Image.new("RGBA", (diameter, diameter), (0, 0, 0, 0))
    out.paste(img, (0, 0), mask)
    return out


def draw_wrapped(draw, text, x, y, w, font, fill, line_h_mult=1.35, align="left"):
    """Word-wrap to width w, return final y."""
    words = text.split()
    line = ""
    cur_y = y
    line_h = round(font.size * line_h_mult)
    for word in words:
        trial = (line + " " + word).strip()
        bbox = draw.textbbox((0, 0), trial, font=font)
        tw = bbox[2] - bbox[0]
        if tw <= w or not line:
            line = trial
        else:
            _draw_aligned(draw, line, x, cur_y, w, font, fill, align)
            cur_y += line_h
            line = word
    if line:
        _draw_aligned(draw, line, x, cur_y, w, font, fill, align)
        cur_y += line_h
    return cur_y


def _draw_aligned(draw, line, x, y, w, font, fill, align):
    if align == "left":
        draw.text((x, y), line, font=font, fill=fill)
    elif align == "right":
        bbox = draw.textbbox((0, 0), line, font=font)
        tw = bbox[2] - bbox[0]
        draw.text((x + w - tw, y), line, font=font, fill=fill)
    else:  # center
        bbox = draw.textbbox((0, 0), line, font=font)
        tw = bbox[2] - bbox[0]
        draw.text((x + (w - tw) // 2, y), line, font=font, fill=fill)


# =====================================================================
def main():
    print(f"📕 Building cover: {LANG}, {PAGECOUNT} pages")
    print(f"   Canvas: {CANVAS_W_IN:.4f}\" × {CANVAS_H_IN:.4f}\" → {CANVAS_W} × {CANVAS_H} px @ {DPI} DPI")
    print(f"   Spine: {SPINE_IN:.4f}\" ({SPINE_IN*25.4:.2f} mm)")

    # 1. Base canvas — petrol all over (4ª capa + spine background)
    canvas = Image.new("RGB", (CANVAS_W, CANVAS_H), PETROL)

    # 2. FRONT COVER (right) — paste capa.jpg over the front trim+bleed area
    front_w = TRIM_PX + BLEED_PX  # extends to right bleed
    front_h = HEIGHT_PX + 2 * BLEED_PX
    front_img = Image.open(CAPA_FRONT).convert("RGB")
    front_img = front_img.resize((front_w, front_h), Image.LANCZOS)
    canvas.paste(front_img, (FRONT_X0, 0))

    # 3. BACK COVER content
    draw = ImageDraw.Draw(canvas)
    pad = round(0.5 * DPI)        # 0.5" inset
    bx0 = BACK_X0 + pad
    bx1 = BACK_X1 - pad
    bw  = bx1 - bx0
    by  = TOP_PX + pad

    # 3a. Blurb body
    f_blurb = f("serif_regular", 10.5)
    f_blurb_b = f("serif_italic", 10.5)  # we'll bold via drawing twice
    f_lead = f("sans_bold", 11)

    blurb = (
        "Ricardo tinha 52 anos, fazia check-up todo ano, "
        "corria três vezes por semana. Oito meses depois "
        "do último exame \"dentro da normalidade\", caiu "
        "no estacionamento do escritório com um infarto.\n\n"
        "Entre o que o seu laboratório chama de normal e "
        "o que a ciência da longevidade chama de ótimo "
        "existe uma janela silenciosa de dez a vinte anos. "
        "É nela que a doença é construída — e é nela que "
        "ainda dá para mudar a trajetória.\n\n"
        "Em ANTES, Dr. Getulio Amaral Filho reúne vinte "
        "anos de prática clínica num mapa que o check-up "
        "convencional não desenha: 16 biomarcadores em 4 "
        "grupos que separam o normal do ótimo, o método "
        "AGIR de quatro pilares e a Regra dos Dois — a "
        "cada trimestre, dois focos com maior potencial "
        "de mover trajetória.\n\n"
        "Não é wellness. É medicina baseada em evidência, "
        "vinte anos de consultório, casos reais — para "
        "quem ainda tem a década inteira pela frente, e "
        "para quem só tem a próxima."
    )

    cy = by
    for paragraph in blurb.split("\n\n"):
        cy = draw_wrapped(draw, paragraph, bx0, cy, bw, f_blurb, CREAM, line_h_mult=1.45)
        cy += round(0.10 * DPI)

    # 3b. Slogan-selo
    cy += round(0.20 * DPI)
    # Gold rule
    rule_w = round(1.4 * DPI)
    rule_x0 = bx0 + (bw - rule_w) // 2
    draw.line((rule_x0, cy, rule_x0 + rule_w, cy), fill=GOLD, width=2)
    cy += round(0.30 * DPI)

    f_slogan = f("serif_italic", 13)
    line1 = "O melhor momento para começar"
    line2 = "era dez anos atrás."
    line3 = "O segundo melhor é hoje."
    for ln in (line1, line2, line3):
        bbox = draw.textbbox((0, 0), ln, font=f_slogan)
        tw = bbox[2] - bbox[0]
        draw.text((bx0 + (bw - tw) // 2, cy), ln, font=f_slogan, fill=GOLD)
        cy += round(f_slogan.size * 1.30)

    cy += round(0.30 * DPI)
    draw.line((rule_x0, cy, rule_x0 + rule_w, cy), fill=GOLD, width=2)
    cy += round(0.35 * DPI)

    # 3c. Author photo (circular) + bio
    photo_dia = round(1.4 * DPI)
    photo_img = circular_crop(Image.open(AUTHOR_BW), photo_dia)
    photo_x = bx0
    photo_y = cy
    canvas.paste(photo_img, (photo_x, photo_y), photo_img)

    # Bio text to the right of the photo
    bio_x = photo_x + photo_dia + round(0.25 * DPI)
    bio_w = bx1 - bio_x
    bio_y = photo_y + round(0.10 * DPI)
    f_bio_name = f("sans_bold", 11)
    f_bio = f("serif_regular", 9.5)
    draw.text((bio_x, bio_y), "Dr. Getulio Amaral Filho", font=f_bio_name, fill=CREAM)
    bio_y += round(f_bio_name.size * 1.4)
    bio = (
        "Nefrologista (CRM-PR 21.876), coordena a Residência de "
        "Nefrologia da Santa Casa de Londrina. Fundador da "
        "Plenya — clínica de medicina integrada em saúde, "
        "performance e longevidade."
    )
    draw_wrapped(draw, bio, bio_x, bio_y, bio_w, f_bio, CREAM, line_h_mult=1.35)

    # 3d. Barcode bottom-right of back cover, P symbol bottom-left
    margin = round(0.4 * DPI)
    # Barcode (white plaque + image)
    bc = Image.open(BARCODE).convert("RGB")
    bc_target_w = round(1.6 * DPI)
    bc_ratio = bc.height / bc.width
    bc = bc.resize((bc_target_w, round(bc_target_w * bc_ratio)), Image.LANCZOS)
    plaque_pad = round(0.08 * DPI)
    plaque_w = bc.width + 2 * plaque_pad
    plaque_h = bc.height + 2 * plaque_pad
    plaque_x = BACK_X1 - margin - plaque_w
    plaque_y = BOT_PX - margin - plaque_h
    draw.rectangle((plaque_x, plaque_y, plaque_x + plaque_w, plaque_y + plaque_h), fill=(255, 255, 255))
    canvas.paste(bc, (plaque_x + plaque_pad, plaque_y + plaque_pad))

    # Plenya P symbol bottom-left of back cover
    if SYMBOL_GOLD.exists():
        sym = Image.open(SYMBOL_GOLD).convert("RGBA")
        sym_target = round(0.65 * DPI)
        sym = fit_into_box(sym, sym_target, sym_target)
        sym_x = BACK_X0 + margin
        sym_y = BOT_PX - margin - sym.height
        canvas.paste(sym, (sym_x, sym_y), sym)

    # 4. SPINE — text rotated 90° clockwise (reads from bottom to top when book stands upright)
    if SPINE_PX >= round(0.25 * DPI):
        # Build a horizontal canvas, then rotate 90° and paste
        spine_h_horiz = SPINE_PX                      # becomes the spine's WIDTH after rotation? wait
        # Spine on cover is a vertical strip. We render text horizontally on a horizontal strip
        # of size HEIGHT_PX × SPINE_PX (height matches spine height before rotation), then rotate -90.
        spine_canvas = Image.new("RGB", (HEIGHT_PX, SPINE_PX), PETROL)
        sd = ImageDraw.Draw(spine_canvas)
        # Title left → right; we'll rotate so title reads bottom-to-top on the wrap.
        f_spine_title = f("sans_bold", max(11, min(int(SPINE_PX / 9.5 / DPI * 72), 22)))
        f_spine_author = f("sans", max(8, min(int(SPINE_PX / 14 / DPI * 72), 14)))
        # ANTES (bigger, pushed to the front-cover end)
        title_text = "ANTES"
        author_text = "Getulio Amaral Filho"
        bbox_t = sd.textbbox((0, 0), title_text, font=f_spine_title)
        tt_w = bbox_t[2] - bbox_t[0]; tt_h = bbox_t[3] - bbox_t[1]
        bbox_a = sd.textbbox((0, 0), author_text, font=f_spine_author)
        ta_w = bbox_a[2] - bbox_a[0]; ta_h = bbox_a[3] - bbox_a[1]
        # Layout horizontally: TITLE on the LEFT, AUTHOR on the RIGHT.
        # After -90° rotation (clockwise), LEFT becomes TOP and RIGHT becomes BOTTOM —
        # which gives the standard convention: title at top of spine, author at bottom.
        title_x = round(0.6 * DPI)
        title_y = (SPINE_PX - tt_h) // 2 - bbox_t[1]
        sd.text((title_x, title_y), title_text, font=f_spine_title, fill=CREAM)
        author_x = HEIGHT_PX - ta_w - round(0.6 * DPI)
        author_y = (SPINE_PX - ta_h) // 2 - bbox_a[1]
        sd.text((author_x, author_y), author_text, font=f_spine_author, fill=GOLD)
        # Rotate 90° so it reads top-to-bottom on the standing book
        spine_rot = spine_canvas.rotate(-90, expand=True)
        canvas.paste(spine_rot, (SPINE_X0, TOP_PX))

    # 5. Save
    canvas.save(OUT_PNG, "PNG")
    # Save PDF with prepress-quality JPEG (Q=95). PIL default is Q=75, which
    # produces visible banding on the cover's flat ocre walls and gradient
    # shadows when printed at 6×9 trim. Q=95 is visually indistinguishable
    # from the source PNG at print resolution.
    canvas.save(OUT_PDF, "PDF", resolution=DPI, quality=95)
    print(f"\n✅ Cover: {OUT_PDF}")
    print(f"✅ Preview: {OUT_PNG}")


if __name__ == "__main__":
    main()
