#!/usr/bin/env python3
"""Composite the transparent Plenya P monogram onto an AI-generated slide.

Usage:
    composite_logo(base_png, out_png, size_pct=8, position="center-bottom", margin_pct=6)

Positions supported: "center-bottom", "center", "footer-right", "footer-center".
size_pct: logo height as percentage of slide height (default 8 = small brand mark).
margin_pct: distance from edge as percentage of slide height.
"""
from PIL import Image
from pathlib import Path

LOGO = Path("/home/user/plenya/docs/palestras/antes-janela-silenciosa/plenya_logo_transparent.png")


def composite_logo(base_png, out_png, size_pct=8, position="center-bottom", margin_pct=6):
    base = Image.open(base_png).convert("RGBA")
    W, H = base.size

    logo = Image.open(LOGO).convert("RGBA")
    target_h = int(H * size_pct / 100)
    scale = target_h / logo.height
    target_w = int(logo.width * scale)
    logo_resized = logo.resize((target_w, target_h), Image.LANCZOS)

    margin = int(H * margin_pct / 100)
    if position == "center-bottom":
        x = (W - target_w) // 2
        y = H - target_h - margin
    elif position == "center":
        x = (W - target_w) // 2
        y = (H - target_h) // 2
    elif position == "footer-right":
        x = W - target_w - margin
        y = H - target_h - margin
    elif position == "footer-center":
        x = (W - target_w) // 2
        y = H - target_h - margin
    else:
        raise ValueError(f"unknown position: {position}")

    canvas = base.copy()
    canvas.paste(logo_resized, (x, y), logo_resized)
    canvas.convert("RGB").save(out_png, "PNG", optimize=True)
    print(f"Saved: {out_png} ({W}x{H}, logo {target_w}x{target_h} at ({x},{y}))")


if __name__ == "__main__":
    import sys
    base = Path(sys.argv[1])
    out = Path(sys.argv[2]) if len(sys.argv) > 2 else base.with_name(base.stem + "-final.png")
    size = float(sys.argv[3]) if len(sys.argv) > 3 else 8
    pos = sys.argv[4] if len(sys.argv) > 4 else "center-bottom"
    composite_logo(base, out, size_pct=size, position=pos)
