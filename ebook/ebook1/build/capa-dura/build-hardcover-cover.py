#!/usr/bin/env python3
"""Build print-ready cover for 'ANTES' — Brazilian HARDCOVER (capa dura).

Usage:
    python3 build-hardcover-cover.py pt-BR                   # uses default spine
    python3 build-hardcover-cover.py pt-BR 25.168            # custom spine in mm

Geometry (from gabarito.pdf supplied by the printer):
    Trim per side    : 16 × 23 cm
    Spine            : printer-calculated, depends on paper + page count.
                       Current value (25,168 mm) returned by Fábrica do Livro
                       for 378 pages on Couché Fosco 115 g/m². Pass a different
                       number on the CLI to override without editing the file.
    Empastamento     : 2 cm top + 2 cm bottom (folds glued onto the horlle)
    No lateral fold  : the gabarito specifies fold areas only on top/bottom

Total canvas:
    Width  = 16 cm + spine + 16 cm
    Height = 2 cm  + 23 cm + 2 cm           = 270 mm

Layout (left to right, viewed flat on table):
    [contracapa 16 cm] [lombada] [capa 16 cm]

    Vertically each panel sits inside a 27 cm canvas with 2 cm fold strips
    above and below — those strips wrap around the cardboard horlle, so we
    extend the petrol background into them but place no critical content
    (slogans / barcode / author photo) within the empastamento area.
"""
from __future__ import annotations
import sys
import io
from pathlib import Path
from PIL import Image, ImageDraw, ImageFont
import img2pdf

THIS_DIR  = Path(__file__).resolve().parent
BUILD_DIR = THIS_DIR.parent
BOOK_ROOT = BUILD_DIR.parent
LANG = sys.argv[1] if len(sys.argv) > 1 else "pt-BR"
DEFAULT_SPINE_MM = 25.168
SPINE_MM_ARG = float(sys.argv[2]) if len(sys.argv) > 2 else DEFAULT_SPINE_MM

# --- Paths ---
CAPA_FRONT    = BOOK_ROOT / "capas" / LANG / "capa.jpg"
BARCODE       = BOOK_ROOT / "capas" / LANG / "isbn-barcode-978-65-02-07691-0.png"
AUTHOR_BW     = BOOK_ROOT / "fotos" / "getulio_bw_halfbody_fullres.jpg"
SYMBOL_GOLD   = BOOK_ROOT.parent.parent / "apps" / "site" / "public" / "brand" / "symbol" / "gold.png"
WORDMARK_GOLD = BOOK_ROOT.parent.parent / "apps" / "site" / "public" / "brand" / "wordmark" / "gold.png"
OUT_PDF       = THIS_DIR / f"Antes-{LANG}-hardcover-capa.pdf"
OUT_PNG       = THIS_DIR / f"Antes-{LANG}-hardcover-capa.png"

# --- Geometry (mm → px @ DPI) ---
# 600 DPI for premium hardcover — offset printers can resolve up to ~600 lpi
# on coated stock, and lossless PNG → img2pdf embedding means there is no
# JPEG generation in the pipeline (no banding on the petrol gradients, no
# softness on the gold filetes).
DPI = 600
def mm_to_px(mm): return round(mm * DPI / 25.4)

TRIM_W_MM   = 160.0          # 16 cm
TRIM_H_MM   = 230.0          # 23 cm
SPINE_MM    = SPINE_MM_ARG   # CLI override or DEFAULT_SPINE_MM
WRAP_TOP_MM = 20.0           # 2 cm empastamento
WRAP_BOT_MM = 20.0           # 2 cm empastamento

CANVAS_W_MM = TRIM_W_MM + SPINE_MM + TRIM_W_MM           # 347,632 mm
CANVAS_H_MM = WRAP_TOP_MM + TRIM_H_MM + WRAP_BOT_MM      # 270 mm

CANVAS_W = mm_to_px(CANVAS_W_MM)
CANVAS_H = mm_to_px(CANVAS_H_MM)

# Boundaries (x, in px)
BACK_X0  = 0
BACK_X1  = mm_to_px(TRIM_W_MM)
SPINE_X0 = BACK_X1
SPINE_X1 = SPINE_X0 + mm_to_px(SPINE_MM)
FRONT_X0 = SPINE_X1
FRONT_X1 = FRONT_X0 + mm_to_px(TRIM_W_MM)

# Boundaries (y, in px)
WRAP_TOP_Y0 = 0
WRAP_TOP_Y1 = mm_to_px(WRAP_TOP_MM)              # top of the trim
TRIM_Y0     = WRAP_TOP_Y1
TRIM_Y1     = TRIM_Y0 + mm_to_px(TRIM_H_MM)      # bottom of the trim
WRAP_BOT_Y0 = TRIM_Y1
WRAP_BOT_Y1 = WRAP_BOT_Y0 + mm_to_px(WRAP_BOT_MM)

# --- Plenya palette ---
PETROL = (6, 59, 79)         # #063b4f
GOLD   = (179, 134, 69)      # #b38645
CREAM  = (234, 231, 218)     # #eae7da
INK    = (26, 26, 26)        # #1a1a1a

# --- Fonts ---
def load_fonts():
    candidates_serif_italic = [
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
        "serif_italic":  first(candidates_serif_italic),
        "sans_bold":     first(candidates_sans_bold),
        "sans":          first(candidates_sans),
    }

FONTS = load_fonts()

def f(role, size_pt):
    px = round(size_pt * DPI / 72)
    return ImageFont.truetype(FONTS[role], px)


# --- Helpers ---
def fit_into_box(img, max_w, max_h):
    img = img.copy()
    img.thumbnail((max_w, max_h), Image.LANCZOS)
    return img


def circular_crop(img, diameter):
    w, h = img.size
    side = min(w, h)
    left = (w - side) // 2
    top  = (h - side) // 2
    img = img.crop((left, top, left + side, top + side)).resize((diameter, diameter), Image.LANCZOS)
    if img.mode != "RGBA":
        img = img.convert("RGBA")
    mask = Image.new("L", (diameter, diameter), 0)
    ImageDraw.Draw(mask).ellipse((0, 0, diameter, diameter), fill=255)
    out = Image.new("RGBA", (diameter, diameter), (0, 0, 0, 0))
    out.paste(img, (0, 0), mask)
    return out


def draw_wrapped(draw, text, x, y, w, font, fill, line_h_mult=1.35, align="left"):
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
    print(f"📕 Building HARDCOVER cover: {LANG}")
    print(f"   Trim per side : {TRIM_W_MM:.2f} × {TRIM_H_MM:.2f} mm")
    print(f"   Spine         : {SPINE_MM:.3f} mm")
    print(f"   Empastamento  : {WRAP_TOP_MM:.0f} mm topo + {WRAP_BOT_MM:.0f} mm base")
    print(f"   Canvas total  : {CANVAS_W_MM:.3f} × {CANVAS_H_MM:.0f} mm → {CANVAS_W} × {CANVAS_H} px @ {DPI} DPI")

    # 1. Base canvas — petrol everywhere (covers fold strips + back + spine)
    canvas = Image.new("RGB", (CANVAS_W, CANVAS_H), PETROL)

    # 2. FRONT COVER (right) — paste capa.jpg over the front trim area, then
    #    EXTEND THE TOP/BOTTOM EDGES into the empastamento strips so the
    #    fold area visually continues the cover photo instead of jumping to
    #    petrol blue.
    #
    #    Why: the empastamento (2 cm top + 2 cm bottom) wraps around the
    #    cardboard horlle. After binding, those bands sit on the top and
    #    bottom edges of the front cover. capa.jpg has near-black ceiling
    #    at the top and dark-terracotta floor at the bottom, so the natural
    #    extension is those dark colors — petrol there reads as a foreign
    #    band when you look at the closed book from the side.
    #
    #    Technique: edge-clamp / edge-replication. Take the top 8 px (very
    #    thin strip — averages out any noise but stays in the same lighting
    #    band) and stretch it vertically across the empastamento. Same for
    #    the bottom. This matches Adobe InDesign's "extend" bleed mode,
    #    which is how Penguin/Knopf etc. handle hardcover wraps.
    front_w = FRONT_X1 - FRONT_X0
    front_h = TRIM_Y1 - TRIM_Y0
    front_img = Image.open(CAPA_FRONT).convert("RGB").resize((front_w, front_h), Image.LANCZOS)

    # 2a. Top empastamento — replicate top edge of cover photo upward
    top_strip = front_img.crop((0, 0, front_w, 8))
    top_extended = top_strip.resize((front_w, WRAP_TOP_Y1 - WRAP_TOP_Y0), Image.LANCZOS)
    canvas.paste(top_extended, (FRONT_X0, WRAP_TOP_Y0))

    # 2b. Cover photo in the trim area
    canvas.paste(front_img, (FRONT_X0, TRIM_Y0))

    # 2c. Bottom empastamento — replicate bottom edge downward
    bot_strip = front_img.crop((0, front_h - 8, front_w, front_h))
    bot_extended = bot_strip.resize((front_w, WRAP_BOT_Y1 - WRAP_BOT_Y0), Image.LANCZOS)
    canvas.paste(bot_extended, (FRONT_X0, WRAP_BOT_Y0))

    # 3. BACK COVER content — laid out within the safety margin (≥0,5 cm
    #    from trim) AND above/below the empastamento (no important elements
    #    in the fold area).
    draw = ImageDraw.Draw(canvas)
    safe_mm = 8.0  # bit more than the 5 mm minimum to breathe
    bx0 = BACK_X0 + mm_to_px(safe_mm)
    bx1 = BACK_X1 - mm_to_px(safe_mm)
    bw  = bx1 - bx0
    by  = TRIM_Y0 + mm_to_px(safe_mm)
    by_max = TRIM_Y1 - mm_to_px(safe_mm)

    # 3a. Blurb body
    f_blurb = f("serif_regular", 10.5)

    blurb = (
        "Ricardo tinha 52 anos, fazia check-up todo ano, "
        "corria três vezes por semana. Oito meses depois "
        "do último exame \"dentro da normalidade\", caiu "
        "no estacionamento do escritório com um infarto.\n\n"
        "Entre o que o seu laboratório chama de normal e "
        "o que a ciência da longevidade chama de ótimo "
        "existe uma janela silenciosa de dez a vinte anos. "
        "É nela que a doença é construída, e é nela que "
        "ainda dá para mudar a trajetória.\n\n"
        "Em ANTES, Dr. Getulio Amaral Filho reúne vinte "
        "anos de prática clínica num mapa que o check-up "
        "convencional não desenha: 16 biomarcadores em 4 "
        "grupos que separam o normal do ótimo, o método "
        "AGIR de quatro pilares e a Regra dos Dois (a "
        "cada trimestre, dois focos com maior potencial "
        "de mover trajetória).\n\n"
        "Não é wellness. É medicina baseada em evidência, "
        "vinte anos de consultório, casos reais. Para "
        "quem ainda tem a década inteira pela frente, e "
        "para quem só tem a próxima."
    )

    cy = by
    for paragraph in blurb.split("\n\n"):
        cy = draw_wrapped(draw, paragraph, bx0, cy, bw, f_blurb, CREAM, line_h_mult=1.45)
        cy += round(0.10 * DPI)

    # 3b. Slogan-selo
    cy += round(0.20 * DPI)
    rule_w = round(1.4 * DPI)
    rule_x0 = bx0 + (bw - rule_w) // 2
    draw.line((rule_x0, cy, rule_x0 + rule_w, cy), fill=GOLD, width=2)
    cy += round(0.30 * DPI)

    f_slogan = f("serif_italic", 13)
    for ln in ("O melhor momento para começar",
               "era dez anos atrás.",
               "O segundo melhor é hoje."):
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
        "Plenya, clínica de medicina integrada em saúde, "
        "performance e longevidade."
    )
    draw_wrapped(draw, bio, bio_x, bio_y, bio_w, f_bio, CREAM, line_h_mult=1.35)

    # 3d. Barcode bottom-right of back trim, P symbol bottom-left of back trim
    margin_mm = 10.0
    margin = mm_to_px(margin_mm)
    bc = Image.open(BARCODE).convert("RGB")
    bc_target_w = round(1.6 * DPI)
    bc_ratio = bc.height / bc.width
    bc = bc.resize((bc_target_w, round(bc_target_w * bc_ratio)), Image.LANCZOS)
    plaque_pad = round(0.08 * DPI)
    plaque_w = bc.width + 2 * plaque_pad
    plaque_h = bc.height + 2 * plaque_pad
    plaque_x = BACK_X1 - margin - plaque_w
    plaque_y = TRIM_Y1 - margin - plaque_h
    draw.rectangle((plaque_x, plaque_y, plaque_x + plaque_w, plaque_y + plaque_h), fill=(255, 255, 255))
    canvas.paste(bc, (plaque_x + plaque_pad, plaque_y + plaque_pad))

    if SYMBOL_GOLD.exists():
        sym = Image.open(SYMBOL_GOLD).convert("RGBA")
        sym_target = round(0.65 * DPI)
        sym = fit_into_box(sym, sym_target, sym_target)
        sym_x = BACK_X0 + margin
        sym_y = TRIM_Y1 - margin - sym.height
        canvas.paste(sym, (sym_x, sym_y), sym)

    # 4. SPINE — title and author rotated 90° clockwise (standing-book convention)
    spine_w_px = SPINE_X1 - SPINE_X0
    spine_h_px = TRIM_Y1 - TRIM_Y0
    if spine_w_px >= round(0.25 * DPI):
        spine_canvas = Image.new("RGB", (spine_h_px, spine_w_px), PETROL)
        sd = ImageDraw.Draw(spine_canvas)
        # Sizes scale with the spine width (27,6 mm gives plenty of room).
        spine_w_mm = SPINE_MM
        title_size = max(11, min(int(spine_w_mm * 0.95), 26))   # ~26 pt with this width
        author_size = max(9, min(int(spine_w_mm * 0.55), 16))
        f_spine_title  = f("sans_bold", title_size)
        f_spine_author = f("sans", author_size)
        title_text  = "ANTES"
        author_text = "Getulio Amaral Filho"
        bbox_t = sd.textbbox((0, 0), title_text, font=f_spine_title)
        tt_w = bbox_t[2] - bbox_t[0]; tt_h = bbox_t[3] - bbox_t[1]
        bbox_a = sd.textbbox((0, 0), author_text, font=f_spine_author)
        ta_w = bbox_a[2] - bbox_a[0]; ta_h = bbox_a[3] - bbox_a[1]
        # Title at top of spine (rotates to LEFT before -90° rotation)
        title_x = round(0.6 * DPI)
        title_y = (spine_w_px - tt_h) // 2 - bbox_t[1]
        sd.text((title_x, title_y), title_text, font=f_spine_title, fill=CREAM)
        author_x = spine_h_px - ta_w - round(0.6 * DPI)
        author_y = (spine_w_px - ta_h) // 2 - bbox_a[1]
        sd.text((author_x, author_y), author_text, font=f_spine_author, fill=GOLD)
        spine_rot = spine_canvas.rotate(-90, expand=True)
        canvas.paste(spine_rot, (SPINE_X0, TRIM_Y0))

    # 5. Save — lossless pipeline.
    # PIL's save("PDF") wraps the canvas in JPEG (quality=95 still leaves
    # visible banding on the petrol gradients at 600 DPI). img2pdf, by
    # contrast, embeds PNG with zlib/flate compression — pixel-perfect,
    # zero generation loss. PDF page size is set explicitly so the printer
    # gets exactly the trim + empastamento dimensions we computed.
    canvas.save(OUT_PNG, "PNG", optimize=True)

    png_bytes = io.BytesIO()
    canvas.save(png_bytes, "PNG", optimize=True)
    png_bytes.seek(0)
    layout_fn = img2pdf.get_layout_fun(
        pagesize=(img2pdf.mm_to_pt(CANVAS_W_MM), img2pdf.mm_to_pt(CANVAS_H_MM))
    )
    with open(OUT_PDF, "wb") as pdf_out:
        pdf_out.write(img2pdf.convert(png_bytes.read(), layout_fun=layout_fn))

    pdf_size_mb = OUT_PDF.stat().st_size / (1024 * 1024)
    png_size_mb = OUT_PNG.stat().st_size / (1024 * 1024)
    print(f"\n✅ Cover (lossless, {DPI} DPI): {OUT_PDF}  —  {pdf_size_mb:.1f} MB")
    print(f"✅ Preview PNG: {OUT_PNG}  —  {png_size_mb:.1f} MB")


if __name__ == "__main__":
    main()
