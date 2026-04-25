#!/usr/bin/env python3
"""Extract the gold P shape from plenya_66.PNG with transparent background.

The source logo has a petrol background ~#0B3B50. This script creates a
version with only the gold glyph visible, safe to composite over any
slide background.
"""
from PIL import Image
from pathlib import Path

SRC = Path("/home/user/plenya/docs/palestras/antes-janela-silenciosa/plenya_66.PNG")
DST = Path("/home/user/plenya/docs/palestras/antes-janela-silenciosa/plenya_logo_transparent.png")

BG = (11, 59, 80)       # petrol of the source file
TOLERANCE = 30          # how far from BG still counts as background

img = Image.open(SRC).convert("RGBA")
px = img.load()
w, h = img.size
for y in range(h):
    for x in range(w):
        r, g, b, a = px[x, y]
        dist = abs(r - BG[0]) + abs(g - BG[1]) + abs(b - BG[2])
        if dist <= TOLERANCE:
            px[x, y] = (0, 0, 0, 0)
img.save(DST)
print(f"Saved: {DST} ({img.size[0]}x{img.size[1]} RGBA)")
