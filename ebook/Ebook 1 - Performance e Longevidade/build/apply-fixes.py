"""
Apply targeted PIL-based fixes to EN figures where gpt-image-2 left
paraphrases or untranslated PT decimals.

Each fix: identify a bbox region in the image, fill with the local
background color, and overlay the correct EN text in a matching font.

Usage:
  ./apply-fixes.py <figure-name>     # apply fixes for one figure
  ./apply-fixes.py --all             # apply all configured fixes
  ./apply-fixes.py --inspect <fig> <x> <y> <w> <h>  # crop + show region
"""

import sys
from pathlib import Path
from PIL import Image, ImageDraw, ImageFont

EBOOK = Path("/home/user/plenya/ebook/Ebook 1 - Performance e Longevidade")
EN_DIR = EBOOK / "figuras" / "en"
BACKUP_DIR = EBOOK / "figuras" / "en-pre-pil-fix"
INSPECT_DIR = Path("/tmp/figure-inspect")
INSPECT_DIR.mkdir(exist_ok=True)
BACKUP_DIR.mkdir(exist_ok=True)

# Fonts available on the system that work well for figure text
FONT_CANDIDATES = [
    "/usr/share/fonts/truetype/inter/Inter-Bold.ttf",
    "/usr/share/fonts/truetype/inter/Inter-Regular.ttf",
    "/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf",
    "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf",
    "/usr/share/fonts/truetype/liberation/LiberationSans-Bold.ttf",
    "/usr/share/fonts/truetype/liberation/LiberationSans-Regular.ttf",
]


def get_font(weight: str = "regular", size: int = 24):
    """weight: 'bold' or 'regular'"""
    for path in FONT_CANDIDATES:
        if (weight == "bold" and "Bold" in path) or (weight == "regular" and "Bold" not in path):
            if Path(path).exists():
                return ImageFont.truetype(path, size)
    # fallback
    return ImageFont.load_default()


def sample_bg_color(img: Image.Image, x: int, y: int) -> tuple:
    """Sample background color near a bbox by averaging pixels just outside it."""
    return img.getpixel((x, y))


def apply_fix(img: Image.Image, fix: dict):
    """Apply a single fix to img (modifies in place)."""
    draw = ImageDraw.Draw(img)
    x, y, w, h = fix["bbox"]
    bg = fix.get("bg_color")
    if bg is None:
        # sample from a corner near the bbox
        bg = sample_bg_color(img, max(0, x - 5), max(0, y - 5))
    # erase
    draw.rectangle([x, y, x + w, y + h], fill=bg)

    if "text" in fix:
        text = fix["text"]
        font = get_font(fix.get("weight", "regular"), fix.get("font_size", 24))
        color = fix.get("color", (10, 17, 34))
        anchor = fix.get("anchor", "lt")  # left-top default
        align = fix.get("align", "left")
        # PIL anchor format: first char = horizontal (l/m/r), second = vertical (t/m/b/a/d/s)
        # We compute the (x, y) point that PIL needs based on which corner/edge of bbox.
        ax, ay = anchor[0], anchor[1]
        tx = x if ax == "l" else (x + w // 2 if ax == "m" else x + w)
        ty = y if ay == "t" else (y + h // 2 if ay == "m" else y + h)
        draw.text((tx, ty), text, font=font, fill=color, anchor=anchor, align=align)


# =============================================================================
# FIXES — keyed by figure filename
# =============================================================================
# These were identified by manual inspection of EN figures vs PT.
# Each entry: list of fix dicts with bbox (x, y, w, h) and text.
# =============================================================================

FIXES: dict[str, list[dict]] = {
    "Cap09 Fig01.PNG": [
        # 1. Fix FIGURA badge → FIGURE 1 (the red box is at x=24..191, y=14..57)
        {
            "bbox": (20, 10, 200, 55),
            "bg_color": (173, 4, 21),  # same red
            "text": "FIGURE 1",
            "font_size": 24,
            "weight": "bold",
            "color": (255, 255, 255),
            "anchor": "mm",
        },
        # 2. Fix title glitch "Three syštemas" → "Three systems, one medicine."
        {
            "bbox": (15, 70, 820, 60),
            "bg_color": (254, 254, 254),
            "text": "Three systems, one medicine.",
            "font_size": 38,
            "weight": "bold",
            "color": (11, 25, 66),
            "anchor": "lm",
        },
        # 3. Fix subtitle (sentido invertido) — original at y=95..137
        {
            "bbox": (15, 130, 1100, 50),
            "bg_color": (254, 254, 254),
            "text": "Why three drug classes stopped being 'diabetes drugs' in the last five years.",
            "font_size": 18,
            "weight": "regular",
            "color": (91, 96, 105),
            "anchor": "lm",
        },
    ],
    "Cap13 Fig02.PNG": [
        # Title: '"My exams are good. I'm not okay."' → '"The labs are good. I am not."'
        {
            "bbox": (40, 55, 1300, 80),
            "bg_color": (254, 254, 254),
            "text": '"The labs are good. I am not."',
            "font_size": 36,
            "weight": "bold",
            "color": (11, 25, 66),
            "anchor": "lm",
        },
        # Lay ritual → Secular ritual (body table — coords found via scan)
        {
            "bbox": (200, 620, 200, 25),
            "bg_color": (254, 254, 254),
            "text": "Secular ritual with Marina",
            "font_size": 14,
            "weight": "regular",
            "color": (11, 25, 66),
            "anchor": "lm",
        },
    ],
    "Cap08 Fig01.PNG": [
        # 1. Title: Mark → Marcos (image=1672x941, title bbox x=71..1417 y=97..147, left-aligned)
        {
            "bbox": (60, 88, 1400, 65),
            "bg_color": (252, 251, 249),
            "text": "Marcos, 8 months later: a fitness gain that's worth a statin.",
            "font_size": 36,
            "weight": "bold",
            "color": (11, 25, 66),
            "anchor": "lm",
        },
        # 2. Big red +1,8 MET (between BASELINE and 8 MONTHS bars)
        {
            "bbox": (275, 432, 285, 75),
            "bg_color": (253, 250, 249),
            "text": "+1.8 MET",
            "font_size": 48,
            "weight": "bold",
            "color": (200, 29, 27),
            "anchor": "mm",
        },
        # 3. Replace "34,3" big number near 8 MONTHS bar (preserves "ml/kg/min" subscript)
        {
            "bbox": (625, 543, 100, 60),
            "bg_color": (253, 251, 249),
            "text": "34.3",
            "font_size": 40,
            "weight": "bold",
            "color": (11, 25, 66),
            "anchor": "lm",
        },
        # 4. Footer +1,8 MET sentence (dark navy, centered)
        {
            "bbox": (400, 780, 940, 50),
            "bg_color": (253, 251, 248),
            "text": "+1.8 MET = 25–30% lower all-cause death risk.",
            "font_size": 24,
            "weight": "bold",
            "color": (11, 25, 66),
            "anchor": "mm",
        },
    ],
    "Cap04 Fig01.PNG": [
        # Title regression: "No Risk Factors." → "None Optimal."
        # Repaint entire title centered (image width=1536; full mask span).
        {
            "bbox": (130, 12, 1280, 70),
            "bg_color": (254, 254, 254),
            "text": "The Fernanda Case: All 'Normal'. None Optimal.",
            "font_size": 42,
            "weight": "bold",
            "color": (11, 25, 66),
            "anchor": "mm",
        },
    ],
    "Cap01 Fig02.PNG": [
        # Title: "None were great." → "None were optimal." (whole title repaint)
        {
            "bbox": (40, 18, 1080, 60),
            "bg_color": (253, 253, 253),
            "text": "All of Ricardo's numbers were 'normal'. None were optimal.",
            "font_size": 36,
            "weight": "bold",
            "color": (11, 25, 66),
            "anchor": "lm",
        },
    ],
    "Cap12 Fig02.PNG": [
        # Subtitle: 2 lines starting below large title
        # Line 1 of subtitle (around y=120-150)
        {
            "bbox": (35, 120, 1280, 35),
            "bg_color": (254, 254, 254),
            "text": "Ana, 44 years old: 18 months of biological optimization without moving these markers.",
            "font_size": 18,
            "weight": "regular",
            "color": (91, 96, 105),
            "anchor": "lm",
        },
        # Line 2 of subtitle
        {
            "bbox": (35, 155, 1280, 60),
            "bg_color": (254, 254, 254),
            "text": "Six months of work on the mind and the four moved.",
            "font_size": 18,
            "weight": "regular",
            "color": (91, 96, 105),
            "anchor": "lm",
        },
    ],
}


def inspect_region(figname: str, x: int, y: int, w: int, h: int):
    """Crop a region of the EN figure and save it to /tmp/figure-inspect."""
    src = EN_DIR / figname
    img = Image.open(src)
    region = img.crop((x, y, x + w, y + h))
    out = INSPECT_DIR / f"{figname.replace(' ', '_')}_x{x}_y{y}.png"
    region.save(out)
    print(f"saved → {out}  (image size: {img.size})")
    # also sample bg colors at corners for overlay match
    print(f"corner bg samples:")
    for label, (cx, cy) in [("TL", (x, y)), ("TR", (x + w, y)), ("BL", (x, y + h)), ("BR", (x + w, y + h))]:
        if 0 <= cx < img.size[0] and 0 <= cy < img.size[1]:
            print(f"  {label} ({cx},{cy}): {img.getpixel((cx, cy))}")


def fix_figure(figname: str):
    src = EN_DIR / figname
    if not src.exists():
        print(f"missing: {src}")
        return False
    fixes = FIXES.get(figname)
    if not fixes:
        print(f"no fixes configured for {figname}")
        return False

    # backup once
    bkp = BACKUP_DIR / figname
    if not bkp.exists():
        Image.open(src).save(bkp)

    img = Image.open(src).convert("RGB")
    for fix in fixes:
        apply_fix(img, fix)
    img.save(src)
    print(f"applied {len(fixes)} fix(es) → {src}")
    return True


def main():
    if len(sys.argv) < 2:
        print(__doc__)
        return

    if sys.argv[1] == "--inspect":
        figname = sys.argv[2]
        x, y, w, h = (int(v) for v in sys.argv[3:7])
        inspect_region(figname, x, y, w, h)
        return

    if sys.argv[1] == "--all":
        for figname in FIXES:
            fix_figure(figname)
        return

    fix_figure(sys.argv[1])


if __name__ == "__main__":
    main()
