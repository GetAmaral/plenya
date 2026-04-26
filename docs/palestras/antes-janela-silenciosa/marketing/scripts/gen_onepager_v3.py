#!/usr/bin/env python3
"""One-pager v3 — hybrid: gpt-image-2 frame + PIL composite of real assets.

Fixes from v2 critique:
- Window references the book cover scene (warm clay walls, small bright window
  casting a golden light pool on the floor) — not a generic invented scene.
- Takeaway 3 rewritten: "A virada · da medicina que reage para a que antecipa."
  Anchored to the brand frase "Saúde não é sobre reagir. É sobre antecipar."
- Photo zone is reserved as a clean petrol rectangle — real HD author photo
  (getulio_bw_halfbody_fullres.jpg) is composited via PIL.
- Wordmark zone is reserved — real PLENYA wordmark gold is composited via PIL.
- P+∞ footer is reserved — real symbol gold is composited via PIL.

Run:
  python3 gen_onepager_v3.py            # full pipeline
  python3 gen_onepager_v3.py --compose  # skip API, recomposite from existing frame
"""
import json
import base64
import pathlib
import urllib.request
import urllib.error
import sys
from PIL import Image, ImageDraw

ROOT = pathlib.Path("/home/user/plenya")
OUT = ROOT / "docs/palestras/antes-janela-silenciosa/marketing/output"
OUT.mkdir(parents=True, exist_ok=True)
FRAME_PATH = OUT / "01-onepager-a4_frame.PNG"
FINAL_PATH = OUT / "01-onepager-a4.PNG"

SYMBOL   = ROOT / "apps/site/public/brand/symbol/gold.png"
PHOTO    = ROOT / "ebook/Ebook 1 - Performance e Longevidade/fotos/getulio_bw_halfbody_fullres.jpg"

# Placeholder zones drawn into the prompt — PIL targets the SAME coordinates.
# Canvas: 1152 x 2048
PHOTO_BOX    = (130, 1640, 400, 1840)      # author photo slot — matches model placeholder
SYMBOL_BOX   = (1010, 1955, 1080, 2020)    # bottom-right P symbol

PETROL = "#063b4f"
GOLD   = "#b38645"
CREAM  = "#eae7da"

# =============================================================================
# PROMPT — v3
# =============================================================================

PROMPT = f"""
BRAND GRAMMAR — STRICT, NON-NEGOTIABLE:
- PLENYA brand palette ONLY: petrol #063b4f (dominant dark), gold #b38645 (accents/numerals/hairlines), ocean #417e8e (gradients), sage #92b8b4 (quiet accents), cream #eae7da (light backgrounds, body on dark).
- Typography: 100% editorial. Tall refined SERIF (Didone/Playfair/Bodoni feel) for the word "ANTES" and headline numerals; modern SANS-SERIF (Söhne/Inter/GT America feel) for body and tags.
- Premium editorial magazine treatment — Monocle / Kinfolk / The Gentleman's Journal. Never corporate, never wellness-cliché, never infographic-y.
- Hairlines only for separators (0.5–1pt, gold or petrol at 20%). Zero heavy borders, zero rounded cards, zero drop shadows.
- NO clip-art icons, NO emojis, NO decorative flourishes, NO tech-startup vibe, NO wellness-influencer vibe.
- Brazilian Portuguese with proper accents (ó, ã, é, ç, í).

CRITICAL — DO NOT DRAW:
- DO NOT draw the word "PLENYA" anywhere on the page. There is no PLENYA wordmark on this piece.
- DO NOT draw any human face, person, or portrait anywhere — that area is reserved for a real photograph composite.
- DO NOT draw any P monogram or P-and-infinity symbol — that area is reserved for a real symbol composite at the bottom-right corner.
- DO NOT draw any box, border, frame, or rectangle outline around the reserved bottom-right symbol zone — it must blend seamlessly into the cream background.
- The reserved zones must be CLEAN, FLAT color blocks (petrol or cream), no decoration inside them.

ASSET: A4 PORTRAIT ONE-PAGER for a premium medical keynote, editorial bookstore quality.

CANVAS: 1152 x 2048 px.

COMPOSITION (top to bottom):

1) TOP BAND — first 38% of canvas height (~778 px tall), background petrol #063b4f.
   The CINEMATIC PHOTOGRAPHIC SCENE behind the title (echoing the book cover of "ANTES"):
   - A quiet interior at golden hour: WARM EARTH-TONED CLAY WALLS (terracotta-petrol blend), softly lit.
   - On the RIGHT third, a single small wooden-framed open WINDOW casting a slanting beam of warm golden sunset light.
   - A pool of that golden light spreads across the floor near the window.
   - The rest of the room recedes into deep petrol shadow with subtle warm highlights.
   - Heavy atmospheric grain, painterly, like a still from a Sorrentino or Villeneuve film.
   - NO PEOPLE. NO furniture. NO objects. Just walls, floor, window, and light.
   - This is the "janela silenciosa" — visual heart of the book and the talk.

   OVERLAID typography on this band:
   - Top-right corner only: tag "KEYNOTE · 45–60 MIN" in cream #eae7da uppercase sans-serif, letterspaced, small (~22pt). Top-LEFT must remain empty — pure photographic petrol band with NO text or wordmark whatsoever.
   - Center-anchored, the HERO HEADLINE:
     · "ANTES" — towering editorial serif (Didone/Bodoni), in warm gold #b38645, centered horizontally, occupying about 35-40% of the band height. This is the visual king.
     · Beneath: "A Janela Silenciosa" in cream bold sans-serif, refined.
     · Beneath that: "entre o Normal e o Ótimo" in cream regular sans-serif, lighter.
   - Beneath headline, one short sage (#92b8b4) regular line: "uma keynote do Dr. Getulio Amaral Filho".

2) HAIRLINE — single 1pt gold (#b38645) horizontal line spanning ~65% of canvas width, centered, separating the petrol band from the cream body. Generous breathing room above and below (~80px).

3) CREAM BODY — remaining 62% of canvas, background cream #eae7da:

   a) HOOK — petrol bold sans-serif, large, left-aligned with ~12% outer margins:
      "Por que o seu check-up anual não está procurando o que vai te matar — e o que pedir para que procure."

   b) SYNOPSIS — petrol regular sans-serif, 4 short lines, justified, comfortable leading:
      "A medicina brasileira se tornou excelente em tratar doença instalada — e ruim em prever o que está chegando. Entre a saúde que o laboratório reconhece e a doença que o médico trata existe uma janela silenciosa de 10 a 20 anos, onde a longevidade é construída ou perdida."

   c) THREE TAKEAWAYS — stacked vertically, each with:
      · A large gold numeral "01" / "02" / "03" in editorial serif on the left.
      · A short petrol BOLD sans headline beside or below the numeral.
      · One petrol regular line of body beneath.
      Content (verbatim, with proper accents):
        01 · O mapa
             16 biomarcadores em 4 grupos — do normal ao ótimo.
        02 · O método
             AGIR — médico, nutricionista, psicólogo, educador físico.
        03 · A virada
             Da medicina que reage para a medicina que antecipa.

   d) Another 1pt gold hairline separator across ~65% width, centered.

   e) AUTHOR BAND — two columns:
      LEFT COLUMN (~30%): a CLEAN FLAT PETROL #063b4f RECTANGLE approximately from x=170 to x=430, y=1430 to y=1690 — RESERVED FOR THE AUTHOR PHOTO COMPOSITE. Do not paint anything inside this rectangle. Beneath the rectangle, in petrol small caps sans, centered: "Dr. Getulio Amaral Filho".
      RIGHT COLUMN (~70%): bio in petrol sans-serif, left-aligned:
         "Nefrologista com 20 anos de prática clínica. Coordena a Residência de Nefrologia da Santa Casa de Londrina. Fundador da Plenya — programa de saúde preventiva e longevidade. Autor de ANTES (2026, ISBN 978-65-02-06742-0)."
      Beneath bio, one gold line, smaller: "CRM-PR 21.876 · RQE 16.038".

   f) Bottom claim strip, centered, in petrol bold small caps letterspaced:
      "VIVA BEM. VIVA MAIS."
      followed by a faint hairline beneath.

   g) FOOTER ROW — small petrol-grey sans-serif at 60% opacity:
      "getamaralb002@gmail.com   ·   @drGetulioAmaralFilho   ·   plenyasaude.com.br"
      The far-right corner must be UNTOUCHED CREAM — no text, no shape, no border, no rectangle, no frame. Just clean cream background. The real P symbol will be composited there afterward; if you draw anything in that corner it will collide.

LAYOUT PRINCIPLES:
- Margins ~110–130 px on all sides.
- Generous whitespace, editorial rhythm — the page must breathe.
- Typographic hierarchy unambiguous at a glance.
- The petrol photographic top + cream editorial bottom must read as one coherent piece, like the cover-plus-TOC of a high-end magazine.

High resolution A4 portrait. Image size: 1152x2048.
"""


def call_api(prompt: str, size: str) -> dict:
    env_path = ROOT / ".env"
    key = None
    for line in env_path.read_text().splitlines():
        line = line.strip()
        if line.startswith("OPENAI_API_KEY="):
            key = line.split("=", 1)[1].strip().strip('"').strip("'")
            break
    assert key, "OPENAI_API_KEY not found"

    payload = json.dumps({
        "model": "gpt-image-2",
        "prompt": prompt,
        "size": size,
        "quality": "high",
        "n": 1,
    }).encode()
    req = urllib.request.Request(
        "https://api.openai.com/v1/images/generations",
        data=payload,
        headers={"Authorization": f"Bearer {key}", "Content-Type": "application/json"},
        method="POST",
    )
    with urllib.request.urlopen(req, timeout=900) as r:
        return json.load(r)


def generate_frame() -> None:
    print("=== Generating one-pager v3 frame (1152x2048) ===", flush=True)
    data = call_api(PROMPT, "1152x2048")
    b64 = data["data"][0].get("b64_json")
    FRAME_PATH.write_bytes(base64.b64decode(b64))
    tokens = data.get("usage", {}).get("total_tokens", 0)
    cost = tokens * 0.04 / 1000
    size_kb = FRAME_PATH.stat().st_size // 1024
    print(f"   ✓ frame saved: {FRAME_PATH.name} ({size_kb} KB, {tokens} tokens, ≈ US$ {cost:.2f})",
          flush=True)


def fit_into_box(asset: Image.Image, box: tuple[int, int, int, int]) -> tuple[Image.Image, tuple[int, int]]:
    """Scale asset preserving aspect ratio to fit inside box; return image + top-left paste coords."""
    bx0, by0, bx1, by1 = box
    bw, bh = bx1 - bx0, by1 - by0
    aw, ah = asset.size
    scale = min(bw / aw, bh / ah)
    nw, nh = int(aw * scale), int(ah * scale)
    resized = asset.resize((nw, nh), Image.LANCZOS)
    px = bx0 + (bw - nw) // 2
    py = by0 + (bh - nh) // 2
    return resized, (px, py)


def crop_to_circle(img: Image.Image, size: int) -> Image.Image:
    """Square-crop an image to size x size and apply a circular alpha mask."""
    w, h = img.size
    s = min(w, h)
    left = (w - s) // 2
    top = (h - s) // 2
    sq = img.crop((left, top, left + s, top + s)).resize((size, size), Image.LANCZOS)
    sq = sq.convert("RGBA")
    mask = Image.new("L", (size, size), 0)
    ImageDraw.Draw(mask).ellipse((0, 0, size, size), fill=255)
    sq.putalpha(mask)
    return sq


def composite() -> None:
    print("=== Compositing real assets onto frame ===", flush=True)
    frame = Image.open(FRAME_PATH).convert("RGBA")

    # 1) Real author photo — circular crop
    photo = Image.open(PHOTO).convert("RGB")
    bx0, by0, bx1, by1 = PHOTO_BOX
    diameter = min(bx1 - bx0, by1 - by0)
    photo_circle = crop_to_circle(photo, diameter)
    px = bx0 + ((bx1 - bx0) - diameter) // 2
    py = by0 + ((by1 - by0) - diameter) // 2
    # Thin gold ring around the portrait
    ring = Image.new("RGBA", photo_circle.size, (0, 0, 0, 0))
    ImageDraw.Draw(ring).ellipse(
        (1, 1, diameter - 1, diameter - 1),
        outline=GOLD, width=2,
    )
    frame.alpha_composite(photo_circle, (px, py))
    frame.alpha_composite(ring, (px, py))
    print(f"   ✓ author photo @ ({px}, {py}) (Ø {diameter})", flush=True)

    # 2) P symbol bottom-right footer (no border)
    sym = Image.open(SYMBOL).convert("RGBA")
    sym_fit, sym_xy = fit_into_box(sym, SYMBOL_BOX)
    frame.alpha_composite(sym_fit, sym_xy)
    print(f"   ✓ symbol @ {sym_xy} ({sym_fit.size})", flush=True)

    frame.convert("RGB").save(FINAL_PATH, "PNG", optimize=True)
    size_kb = FINAL_PATH.stat().st_size // 1024
    print(f"   ✓ final saved: {FINAL_PATH.name} ({size_kb} KB)", flush=True)


def main() -> None:
    compose_only = "--compose" in sys.argv
    if not compose_only:
        try:
            generate_frame()
        except urllib.error.HTTPError as e:
            print(f"   ✗ HTTP {e.code}: {e.read().decode('utf-8', 'replace')[:600]}", flush=True)
            sys.exit(1)
        except Exception as e:
            print(f"   ✗ {type(e).__name__}: {e}", flush=True)
            sys.exit(1)
    if not FRAME_PATH.exists():
        print(f"   ✗ frame missing at {FRAME_PATH}", flush=True)
        sys.exit(1)
    composite()


if __name__ == "__main__":
    main()
