#!/usr/bin/env python3
"""Build slide 72 (closing) with the REAL book cover + contact info.

Layout (2048x1152):
- Petrol background #063b4f
- Real book cover (capa.jpg) on the left, occupying ~38% width, vertically centered
  with elegant drop-shadow for depth
- Right column: vertical text block with contact info
- Plenya P monogram small at bottom-right
"""
from PIL import Image, ImageDraw, ImageFont, ImageFilter
from pathlib import Path

CAPA = Path("/home/user/plenya/ebook/Ebook 1 - Performance e Longevidade/capas/pt-BR/capa.jpg")
LOGO = Path("/home/user/plenya/docs/palestras/antes-janela-silenciosa/plenya_logo_transparent.png")
OUT = Path("/home/user/plenya/docs/palestras/antes-janela-silenciosa/slides/slide-72-capa-livro-contato-final.png")

W, H = 2048, 1152
PETROL = (6, 59, 79)
CREAM = (234, 231, 218)
GOLD = (179, 134, 69)

# Canvas
canvas = Image.new("RGB", (W, H), PETROL)

# Load cover and resize to ~80% of slide height
cover = Image.open(CAPA).convert("RGB")
target_h = int(H * 0.80)          # 921 px
scale = target_h / cover.height
target_w = int(cover.width * scale)  # proportional
cover_resized = cover.resize((target_w, target_h), Image.LANCZOS)

# Drop shadow for cover
shadow_pad = 28
shadow = Image.new("RGBA", (target_w + shadow_pad*2, target_h + shadow_pad*2), (0, 0, 0, 0))
sd = ImageDraw.Draw(shadow)
sd.rectangle([shadow_pad, shadow_pad, shadow_pad+target_w, shadow_pad+target_h], fill=(0, 0, 0, 140))
shadow = shadow.filter(ImageFilter.GaussianBlur(16))

# Place cover on left
cover_x = int(W * 0.12)            # left padding 12%
cover_y = (H - target_h) // 2

# Paste shadow first, then cover
canvas.paste(shadow, (cover_x - shadow_pad + 10, cover_y - shadow_pad + 14), shadow)
canvas.paste(cover_resized, (cover_x, cover_y))

# Right column text block
draw = ImageDraw.Draw(canvas)

serif = "/usr/share/fonts/truetype/liberation/LiberationSerif-Regular.ttf"
serif_italic = "/usr/share/fonts/truetype/liberation/LiberationSerif-Italic.ttf"
sans = "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf"

# Right side begins where cover ends + gutter
right_x = cover_x + target_w + 120

# Block 1: Title
f_eyebrow = ImageFont.truetype(sans, 24)
f_h1 = ImageFont.truetype(serif, 74)
f_h2 = ImageFont.truetype(serif_italic, 38)
f_body = ImageFont.truetype(sans, 28)
f_small = ImageFont.truetype(sans, 22)
f_tiny  = ImageFont.truetype(sans, 19)

y = cover_y + 40

# Eyebrow
draw.text((right_x, y), "LIVRO", font=f_eyebrow, fill=GOLD,
          stroke_width=0, spacing=6)
y += 54

# Title
draw.text((right_x, y), "ANTES", font=f_h1, fill=CREAM)
y += 92

# Subtitle (italic)
draw.text((right_x, y), "A Janela Silenciosa", font=f_h2, fill=CREAM)
y += 50
draw.text((right_x, y), "entre o Normal e o Ótimo", font=f_h2, fill=CREAM)
y += 88

# Divider
draw.line([(right_x, y), (right_x + 120, y)], fill=GOLD, width=2)
y += 44

# Contact block
draw.text((right_x, y), "Dr. Getulio Amaral Filho", font=f_body, fill=CREAM)
y += 42
draw.text((right_x, y), "CRM-PR 21.876 · RQE 16.038", font=f_small, fill=(234, 231, 218, 180))
y += 62

draw.text((right_x, y), "Disponível na Amazon", font=f_body, fill=CREAM)
y += 42
draw.text((right_x, y), "ISBN 978-65-02-06742-0", font=f_small, fill=GOLD)
y += 62

draw.text((right_x, y), "@drGetulioAmaralFilho", font=f_small, fill=CREAM)
y += 32
draw.text((right_x, y), "drgetulioamaralfilho.com.br", font=f_small, fill=CREAM)
y += 32
draw.text((right_x, y), "plenyasaude.com.br", font=f_small, fill=CREAM)

# Plenya P monogram bottom-right
logo = Image.open(LOGO).convert("RGBA")
logo_h = int(H * 0.055)  # 63 px
logo_w = int(logo.width * logo_h / logo.height)
logo_resized = logo.resize((logo_w, logo_h), Image.LANCZOS)
margin = int(H * 0.04)
canvas_rgba = canvas.convert("RGBA")
canvas_rgba.paste(logo_resized, (W - logo_w - margin, H - logo_h - margin), logo_resized)

canvas_rgba.convert("RGB").save(OUT, "PNG", optimize=True)
print(f"Saved: {OUT} ({OUT.stat().st_size:,} bytes)")
