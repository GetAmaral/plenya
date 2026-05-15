#!/usr/bin/env python3
"""Build front cover for AGORA · Vol. 1 — DO NORMAL AO ÓTIMO.

v5 (master) refinements over v4:
    - tighter leading on two-line title (×0.92 of nominal line height)
    - VOL. 1 right-aligned to LAST line ("AO ÓTIMO"), not widest line
    - VOL. 1 sits closer to title baseline (editorial stamp position)
    - gancho italic size 0.048 (was 0.052) for clearer subordination
    - more breathing room between gancho and territory tag
    - territory tag with single spaces + tracking 0.34 (cleaner editorial)
    - thin gold rule centered below territory tag (closes text block, brand anchor)
"""

from pathlib import Path
from PIL import Image, ImageDraw, ImageFont

ROOT = Path(__file__).resolve().parent.parent
INPUT = ROOT / "capas" / "pt-BR" / "capa-base.jpg"
OUTPUT = ROOT / "capas" / "pt-BR" / "capa-v7-final.jpg"

IVORY = (240, 233, 217)        # #F0E9D9 — title, author (full presence)
SOFT_IVORY = (229, 220, 198)   # #E5DCC6 — gancho + tag (voice/label, legible)
DIM_IVORY = (210, 200, 180)    # #D2C8B4 — vol prefix only (most subordinated)
GOLD = (201, 149, 84)          # #C99554 — Plenya brand accent, saturated for visibility

SERIF_BOLD = "/usr/share/fonts/truetype/adf/BerenisADFPro-Bold.otf"
SERIF_ITALIC = "/usr/share/fonts/truetype/ebgaramond/EBGaramond12-Italic.ttf"
SERIF_REGULAR = "/usr/share/fonts/truetype/ebgaramond/EBGaramond12-Regular.ttf"
SANS_MEDIUM = "/usr/share/fonts/opentype/inter/Inter-Medium.otf"
SANS_BOLD = "/usr/share/fonts/opentype/inter/Inter-Bold.otf"


def text_w(draw, text, font):
    bbox = draw.textbbox((0, 0), text, font=font)
    return bbox[2] - bbox[0]


def text_h(draw, text, font):
    bbox = draw.textbbox((0, 0), text, font=font)
    return bbox[3] - bbox[1]


def draw_tracked(draw, text, x, y, font, fill, tracking_em=0.20):
    tracking_px = round(font.size * tracking_em)
    cur_x = x
    for ch in text:
        draw.text((cur_x, y), ch, font=font, fill=fill)
        cur_x += text_w(draw, ch, font) + tracking_px


def tracked_w(draw, text, font, tracking_em=0.20):
    tracking_px = round(font.size * tracking_em)
    total = 0
    for ch in text:
        total += text_w(draw, ch, font) + tracking_px
    return total - tracking_px


def main():
    if not INPUT.exists():
        raise FileNotFoundError(f"Base image not found: {INPUT}")

    img = Image.open(INPUT).convert("RGB")
    W, H = img.size
    draw = ImageDraw.Draw(img)
    print(f"Base: {W} × {H} px")

    # =================================================================
    # TITLE — "DO NORMAL / AO ÓTIMO" — two-line serif, tight leading
    # =================================================================
    title_pt_ratio = 0.090
    title_size = round(H * title_pt_ratio)
    f_title = ImageFont.truetype(SERIF_BOLD, title_size)
    title_lines = ["DO NORMAL", "AO ÓTIMO"]
    title_y = round(H * 0.075)
    ascent, descent = f_title.getmetrics()
    line_h = ascent + descent

    # Tight leading: 92% of nominal line height. Two lines feel as one unit.
    leading_factor = 0.92
    line_advance = round(line_h * leading_factor)

    line_widths = [text_w(draw, t, font=f_title) for t in title_lines]
    last_line_right_edge = (W - line_widths[-1]) // 2 + line_widths[-1]

    cur_y = title_y
    for i, (tl, tw) in enumerate(zip(title_lines, line_widths)):
        draw.text(((W - tw) // 2, cur_y), tl, font=f_title, fill=IVORY)
        if i < len(title_lines) - 1:
            cur_y += line_advance  # tight advance only between lines
    title_bottom = cur_y + line_h  # last line extends full line_h below its top

    # =================================================================
    # VOL. 1 — "VOL. " in dim ivory + "1" in gold, larger and bolder
    # Right-aligned to LAST title line. Baselines aligned across sizes.
    # =================================================================
    vol_size = round(H * 0.018)
    vol_size_big = round(vol_size * 1.30)  # digit 30% bigger to declare itself
    f_vol = ImageFont.truetype(SANS_MEDIUM, vol_size)
    f_vol_big = ImageFont.truetype(SANS_BOLD, vol_size_big)

    prefix = "VOL. "
    vol_tracking_px = round(f_vol.size * 0.30)
    prefix_w = tracked_w(draw, prefix, f_vol, tracking_em=0.30)
    digit_w = text_w(draw, "1", font=f_vol_big)
    # End-to-end width: prefix tracked + trailing tracking gap + digit
    total_vol_w = prefix_w + vol_tracking_px + digit_w

    vol_y = title_bottom - round(H * 0.005)
    vol_x = last_line_right_edge - total_vol_w

    # Draw "VOL. " in DIM_IVORY at vol_y
    cur_x = vol_x
    for ch in prefix:
        draw.text((cur_x, vol_y), ch, font=f_vol, fill=DIM_IVORY)
        cur_x += text_w(draw, ch, font=f_vol) + vol_tracking_px

    # Draw "1" in GOLD, larger, baselines aligned with the small caps
    small_ascent, _ = f_vol.getmetrics()
    big_ascent, _ = f_vol_big.getmetrics()
    digit_y = vol_y + small_ascent - big_ascent
    draw.text((cur_x, digit_y), "1", font=f_vol_big, fill=GOLD)

    # =================================================================
    # GANCHO — "O cansaço que não passa" — italic serif, smaller for hierarchy
    # =================================================================
    sub_y = vol_y + vol_size + round(H * 0.045)
    sub_size = round(H * 0.048)
    f_sub = ImageFont.truetype(SERIF_ITALIC, sub_size)
    sub_text = "O cansaço que não passa"
    stw = text_w(draw, sub_text, font=f_sub)
    draw.text(((W - stw) // 2, sub_y), sub_text, font=f_sub, fill=SOFT_IVORY)
    s_ascent, s_descent = f_sub.getmetrics()
    sub_bottom = sub_y + s_ascent + s_descent

    # =================================================================
    # TERRITORY TAG — single spaces, tracking 0.34, more breathing above
    # =================================================================
    terr_y = sub_bottom + round(H * 0.030)
    terr_size = round(H * 0.018)
    f_terr = ImageFont.truetype(SANS_MEDIUM, terr_size)
    terr_text = "ENERGIA E DISPOSIÇÃO"
    terr_tw = tracked_w(draw, terr_text, f_terr, tracking_em=0.34)
    draw_tracked(
        draw, terr_text,
        (W - terr_tw) // 2, terr_y,
        f_terr, SOFT_IVORY, tracking_em=0.34,
    )

    # =================================================================
    # AUTHOR
    # =================================================================
    author_size = round(H * 0.022)
    f_author = ImageFont.truetype(SANS_BOLD, author_size)
    author_text = "DR. GETULIO AMARAL FILHO"
    author_tw = tracked_w(draw, author_text, f_author, tracking_em=0.25)
    author_y = round(H * 0.940)
    draw_tracked(
        draw, author_text,
        (W - author_tw) // 2, author_y,
        f_author, IVORY, tracking_em=0.25,
    )

    OUTPUT.parent.mkdir(parents=True, exist_ok=True)
    img.save(OUTPUT, "JPEG", quality=95)
    print(f"Saved: {OUTPUT}")


if __name__ == "__main__":
    main()
