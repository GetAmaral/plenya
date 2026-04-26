#!/usr/bin/env python3
"""ASSASSINOS one-pager — A4 portrait, hybrid pipeline.

Mirrors gen_onepager.py from the AGIR talk, adapted to the technical
keynote "Os Quatro Assassinos Silenciosos":
- Hero word "QUATRO" + sub "ASSASSINOS SILENCIOSOS" (long word fits as
  bold sans subtitle, not as Didone hero — keeps visual hierarchy clean).
- Cinematic top band: an ancient tree at dusk with FOUR prominent
  branches and exposed deep roots at the base — visual translation of
  Fig 2.3 (árvore das raízes comuns).
- Three takeaways focused on technical content (raízes biológicas +
  tabela 16 marcadores + condutas).

Run:
  python3 gen_onepager.py
  python3 gen_onepager.py --compose
"""
import json
import base64
import pathlib
import urllib.request
import urllib.error
import sys
from PIL import Image, ImageDraw

ROOT = pathlib.Path("/home/user/plenya")
OUT = ROOT / "docs/palestras/quatro-assassinos-silenciosos/marketing/output"
OUT.mkdir(parents=True, exist_ok=True)
FRAME_PATH = OUT / "01-onepager-a4_frame.PNG"
FINAL_PATH = OUT / "01-onepager-a4.PNG"

SYMBOL = ROOT / "apps/site/public/brand/symbol/gold.png"
PHOTO = ROOT / "ebook/Ebook 1 - Performance e Longevidade/fotos/getulio_bw_halfbody_fullres.jpg"

PHOTO_BOX = (130, 1640, 400, 1840)
SYMBOL_BOX = (1010, 1955, 1080, 2020)

GOLD = "#b38645"

PROMPT = f"""
BRAND GRAMMAR — STRICT, NON-NEGOTIABLE:
- PLENYA brand palette ONLY: petrol #063b4f (dominant dark), gold #b38645 (accents/numerals/hairlines), ocean #417e8e (gradients), sage #92b8b4 (quiet accents), cream #eae7da (light backgrounds, body on dark).
- Typography: 100% editorial. Tall refined SERIF (Didone/Playfair/Bodoni feel) for the word "QUATRO" and headline numerals; modern SANS-SERIF (Söhne/Inter/GT America feel) for body and tags.
- Premium editorial magazine treatment — Monocle / Kinfolk / The Gentleman's Journal. Never corporate, never wellness-cliché, never infographic-y.
- Hairlines only for separators (0.5–1pt, gold or petrol at 20%). Zero heavy borders, zero rounded cards, zero drop shadows.
- NO clip-art icons, NO emojis, NO decorative flourishes, NO tech-startup vibe, NO wellness-influencer vibe.
- Brazilian Portuguese with proper accents (ó, ã, é, ç, í).

CRITICAL — DO NOT DRAW:
- DO NOT draw the word "PLENYA" anywhere on the page.
- DO NOT draw any human face, person, or portrait — that area is reserved for a real photograph composite.
- DO NOT draw any P monogram or P-and-infinity symbol — that area is reserved for a real symbol composite.
- DO NOT draw any box, border, frame, or rectangle outline around the reserved bottom-right symbol zone.
- The reserved zones must be CLEAN, FLAT color blocks (petrol or cream).
- Spell "QUATRO" exactly with classic Q-U-A-T-R-O uppercase — no special letterforms.
- Spell "ASSASSINOS" exactly — no diacritics, no ligatures.

ASSET: A4 PORTRAIT ONE-PAGER for a TECHNICAL medical keynote (the third in a series — companion to ANTES and AGIR), editorial bookstore quality, intended for medical conferences and residency programs.

CANVAS: 1152 x 2048 px.

COMPOSITION (top to bottom):

1) TOP BAND — first 38% of canvas height (~778 px tall), background petrol #063b4f.
   The CINEMATIC PHOTOGRAPHIC SCENE behind the title — visual translation of the book's Fig 2.3 (the tree with common roots):
   - A single ANCIENT, GNARLED TREE at dusk, photographed against a deep petrol sky with hints of cold violet-blue.
   - The tree has FOUR PROMINENT MAIN BRANCHES reaching outward — visible against the sky.
   - At the base, the SOIL IS PARTIALLY CUT AWAY (cross-section style, geological feel) revealing THICK DEEP ROOTS spreading down — the roots illuminated subtly in warm gold tones.
   - Heavy atmospheric grain, painterly, like a still from a Villeneuve or Sorrentino film. Deep color contrast: petrol sky + gold-lit roots.
   - NO PEOPLE. NO furniture. NO objects. NO numbers, letters, or labels visible inside the scene.

   OVERLAID typography on this band:
   - Top-right corner only: tag "KEYNOTE TÉCNICA · 45 MIN" in cream #eae7da uppercase sans-serif, letterspaced, small (~22pt).
   - Top-LEFT must remain empty.
   - Center-anchored, the HERO HEADLINE:
     · "QUATRO" — towering editorial serif (Didone/Bodoni), in warm gold #b38645, centered horizontally, occupying about 30% of the band height. Spell exactly Q-U-A-T-R-O.
     · Beneath: "ASSASSINOS SILENCIOSOS" in cream bold sans-serif, refined, letterspaced, smaller — about 1/3 of the height of QUATRO.
     · Beneath that: "Cardio · Metabólico · Neuro · Câncer" in cream regular sans-serif, lighter.
   - Beneath headline, one short sage (#92b8b4) regular line: "uma keynote do Dr. Getulio Amaral Filho".

2) HAIRLINE — single 1pt gold (#b38645) horizontal line spanning ~65% of canvas width, centered.

3) CREAM BODY — remaining 62% of canvas, background cream #eae7da:

   a) HOOK — petrol bold sans-serif, large, left-aligned with ~12% outer margins:
      "Quatro doenças. Três raízes biológicas. Uma janela de vinte anos antes do diagnóstico."

   b) SYNOPSIS — petrol regular sans-serif, 4 short lines, justified:
      "Versão técnica do Capítulo 2 do livro ANTES. Discute os mecanismos compartilhados das quatro classes de doença que respondem por 80% da mortalidade — inflamação crônica, resistência insulínica, disfunção mitocondrial — e os marcadores que se movem antes da clínica."

   c) THREE TAKEAWAYS — stacked vertically, each with a large gold numeral "01"/"02"/"03" + petrol BOLD sans headline + 1 petrol regular line:
        01 · A árvore das raízes comuns
             Um modelo conceitual para os quatro galhos.
        02 · 16 biomarcadores em 4 grupos
             Tabela completa do livro — normal vs. ótimo.
        03 · Três marcos farmacológicos
             SGLT2 · finerenona · semaglutida — fora do diabetes.

   d) Another 1pt gold hairline separator across ~65% width, centered.

   e) AUTHOR BAND — two columns:
      LEFT COLUMN (~30%): a CLEAN FLAT PETROL #063b4f RECTANGLE approximately from x=170 to x=430, y=1430 to y=1690 — RESERVED FOR THE AUTHOR PHOTO. Beneath the rectangle, in petrol small caps sans, centered: "Dr. Getulio Amaral Filho".
      RIGHT COLUMN (~70%): bio in petrol sans-serif, left-aligned:
         "Nefrologista com 20 anos de prática clínica. Coordena a Residência de Nefrologia da Santa Casa de Londrina. Membro da Sociedade Brasileira de Nefrologia. Fundador da Plenya. Autor de ANTES (2026, ISBN 978-65-02-06742-0)."
      Beneath bio, one gold line, smaller: "CRM-PR 21.876 · RQE 16.038".

   f) Bottom claim strip, centered, in petrol bold small caps letterspaced:
      "VIVA BEM. VIVA MAIS."
      followed by a faint hairline beneath.

   g) FOOTER ROW — small petrol-grey sans-serif at 60% opacity:
      "getamaralb002@gmail.com   ·   @drGetulioAmaralFilho   ·   plenyasaude.com.br"
      The far-right corner must be UNTOUCHED CREAM.

LAYOUT PRINCIPLES:
- Margins ~110–130 px on all sides.
- Generous whitespace, editorial rhythm.
- Typographic hierarchy unambiguous at a glance.

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
    print("=== Generating ASSASSINOS one-pager frame (1152x2048) ===", flush=True)
    data = call_api(PROMPT, "1152x2048")
    b64 = data["data"][0].get("b64_json")
    FRAME_PATH.write_bytes(base64.b64decode(b64))
    tokens = data.get("usage", {}).get("total_tokens", 0)
    cost = tokens * 0.04 / 1000
    size_kb = FRAME_PATH.stat().st_size // 1024
    print(f"   ✓ frame saved: {FRAME_PATH.name} ({size_kb} KB, {tokens} tokens, ≈ US$ {cost:.2f})",
          flush=True)


def fit_into_box(asset, box):
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
    print("=== Compositing real assets onto ASSASSINOS frame ===", flush=True)
    frame = Image.open(FRAME_PATH).convert("RGBA")

    photo = Image.open(PHOTO).convert("RGB")
    bx0, by0, bx1, by1 = PHOTO_BOX
    diameter = min(bx1 - bx0, by1 - by0)
    photo_circle = crop_to_circle(photo, diameter)
    px = bx0 + ((bx1 - bx0) - diameter) // 2
    py = by0 + ((by1 - by0) - diameter) // 2
    ring = Image.new("RGBA", photo_circle.size, (0, 0, 0, 0))
    ImageDraw.Draw(ring).ellipse(
        (1, 1, diameter - 1, diameter - 1),
        outline=GOLD, width=2,
    )
    frame.alpha_composite(photo_circle, (px, py))
    frame.alpha_composite(ring, (px, py))
    print(f"   ✓ author photo @ ({px}, {py}) (Ø {diameter})", flush=True)

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
