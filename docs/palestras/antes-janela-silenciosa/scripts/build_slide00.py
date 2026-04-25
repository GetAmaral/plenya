#!/usr/bin/env python3
"""Build slide 00 — pre-talk holding slide for the MC / audience wait.

Shows on screen while the master of ceremonies introduces the speaker.
Stays up until Dr. Getulio steps on stage and the palestra begins (which
triggers slide 01 — black absolute — for the opening silence).
"""
from PIL import Image, ImageDraw, ImageFont
from pathlib import Path

LOGO = Path("/home/user/plenya/docs/palestras/antes-janela-silenciosa/plenya_logo_transparent.png")
OUT = Path("/home/user/plenya/docs/palestras/antes-janela-silenciosa/slides/slide-00-pre-talk-final.png")

W, H = 2048, 1152
PETROL = (6, 59, 79)
CREAM = (234, 231, 218)
CREAM_DIM = (234, 231, 218, 160)
GOLD = (179, 134, 69)

canvas = Image.new("RGB", (W, H), PETROL)
draw = ImageDraw.Draw(canvas)

serif = "/usr/share/fonts/truetype/liberation/LiberationSerif-Regular.ttf"
serif_italic = "/usr/share/fonts/truetype/liberation/LiberationSerif-Italic.ttf"
sans = "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf"

f_eyebrow = ImageFont.truetype(sans, 28)
f_title = ImageFont.truetype(serif, 300)
f_subtitle = ImageFont.truetype(serif_italic, 52)
f_author = ImageFont.truetype(serif, 48)
f_credentials = ImageFont.truetype(sans, 24)
f_tagline = ImageFont.truetype(serif_italic, 30)


def centered(text, y, font, fill, tracking=0):
    bbox = draw.textbbox((0, 0), text, font=font)
    w = bbox[2] - bbox[0]
    x = (W - w) // 2
    if tracking:
        # Manual letter-spacing: draw each char with extra space
        cx = x - tracking * (len(text) - 1) // 2 * 0
        running = 0
        # Measure full with tracking
        widths = []
        for ch in text:
            cb = draw.textbbox((0, 0), ch, font=font)
            widths.append(cb[2] - cb[0])
        total = sum(widths) + tracking * (len(text) - 1)
        cx = (W - total) // 2
        for i, ch in enumerate(text):
            draw.text((cx, y), ch, font=font, fill=fill)
            cx += widths[i] + tracking
    else:
        draw.text((x, y), text, font=font, fill=fill)


# Eyebrow "EM INSTANTES" — tight letter-spaced uppercase
y = 160
centered("EM INSTANTES", y, f_eyebrow, GOLD, tracking=10)

# Title "ANTES" — huge serif
y = 250
centered("ANTES", y, f_title, CREAM, tracking=18)

# Subtitle
y = 630
centered("A Janela Silenciosa entre o Normal e o Ótimo", y, f_subtitle, CREAM)

# Hairline gold
y = 740
line_w = 180
draw.line([((W - line_w) // 2, y), ((W + line_w) // 2, y)], fill=GOLD, width=2)

# Author
y = 780
centered("com Dr. Getulio Amaral Filho", y, f_author, CREAM)

# Credentials line
y = 854
centered("Nefrologista · CRM-PR 21.876 · RQE 16.038 · Autor de ANTES (2026)", y, f_credentials, (200, 197, 185))

# Tagline bottom
y = 940
centered("Viva bem. Viva mais.", y, f_tagline, (179, 134, 69))

# Plenya P monogram bottom center
logo = Image.open(LOGO).convert("RGBA")
logo_h = int(H * 0.07)
logo_w = int(logo.width * logo_h / logo.height)
logo_resized = logo.resize((logo_w, logo_h), Image.LANCZOS)
canvas_rgba = canvas.convert("RGBA")
canvas_rgba.paste(logo_resized, ((W - logo_w) // 2, H - logo_h - 50), logo_resized)

canvas_rgba.convert("RGB").save(OUT, "PNG", optimize=True)
print(f"Saved: {OUT} ({OUT.stat().st_size:,} bytes)")
