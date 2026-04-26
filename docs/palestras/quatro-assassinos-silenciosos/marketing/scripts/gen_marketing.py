#!/usr/bin/env python3
"""ASSASSINOS marketing — pieces 02–07 (gpt-image-2 + PIL composite).

Mirrors gen_marketing.py from AGIR adapted for the technical keynote:
- Hero word QUATRO + sub ASSASSINOS SILENCIOSOS — keeps Didone visual
  hierarchy clean despite long subtitle.
- Cinematic scenes use the GNARLED TREE WITH 4 BRANCHES + EXPOSED ROOTS
  metaphor (visual translation of Fig 2.3 of the book).
- Social comparative (07) shows the ROOTS table — 3 raízes biológicas
  → 4 galhos clínicos, instead of biomarker normal vs ótimo.
- Quote card (06) uses the technical frase-pedra: "Quem ataca a raiz
  protege os quatro galhos."

Run:
  python3 gen_marketing.py
  python3 gen_marketing.py 02 04
  python3 gen_marketing.py --compose
"""
import json
import base64
import pathlib
import urllib.request
import urllib.error
import sys
from PIL import Image

ROOT = pathlib.Path("/home/user/plenya")
OUT = ROOT / "docs/palestras/quatro-assassinos-silenciosos/marketing/output"
OUT.mkdir(parents=True, exist_ok=True)

WORDMARK_GOLD = ROOT / "apps/site/public/brand/wordmark/gold.png"
SYMBOL_GOLD = ROOT / "apps/site/public/brand/symbol/gold.png"
SYMBOL_INK = ROOT / "apps/site/public/brand/symbol/ink.png"

BRAND_BLOCK = """
BRAND GRAMMAR — STRICT, NON-NEGOTIABLE:
- PLENYA brand palette ONLY: petrol #063b4f, gold #b38645, ocean #417e8e, sage #92b8b4, cream #eae7da.
- Typography: 100% editorial. Tall refined SERIF (Didone/Playfair/Bodoni feel) for headlines and numerals; modern SANS-SERIF (Söhne/Inter/GT America feel) for body and tags.
- Premium editorial magazine treatment — Monocle / Kinfolk / The Gentleman's Journal. Never corporate, never wellness-cliché, never infographic-y.
- Hairlines only for separators (0.5–1pt, gold or petrol at 20%). Zero heavy borders, zero rounded cards, zero drop shadows.
- NO clip-art icons, NO emojis, NO decorative flourishes, NO tech-startup vibe, NO wellness-influencer vibe.
- Brazilian Portuguese with proper accents (ó, ã, é, ç, í).

CRITICAL — DO NOT DRAW ANY OF THE FOLLOWING:
- DO NOT draw the word "PLENYA" anywhere.
- DO NOT draw any "Y" with a diacritic.
- DO NOT draw any P symbol, P monogram, P + infinity, or any infinity symbol.
- DO NOT draw any logo whatsoever.
- The reserved zones must remain clean and flat colored.
- Spell "QUATRO" exactly with classic Q-U-A-T-R-O uppercase.
- Spell "ASSASSINOS" exactly — no diacritics, no ligatures.
"""

# =============================================================================
# PROMPTS
# =============================================================================

PROMPT_2 = BRAND_BLOCK + """
ASSET: A5 PORTRAIT FLYER FRONT — physical poster for medical congress kit.

CANVAS: 1152x2048.

COMPOSITION:

1) TOP 62% of the canvas — CINEMATIC PHOTOGRAPHIC TREATMENT:
   A single ANCIENT GNARLED TREE at dusk, photographed against a deep petrol sky with hints of cold violet-blue. The tree has FOUR PROMINENT MAIN BRANCHES reaching outward against the sky. At the base, the soil is partially cut away (cross-section style, geological feel) revealing THICK DEEP ROOTS spreading down — the roots illuminated subtly in warm gold tones. Heavy atmospheric grain, painterly. NO PEOPLE. NO objects.

   OVERLAID typography:
   - TOP-LEFT (x≈80–460, y≈100–185): RESERVED FOR REAL WORDMARK COMPOSITE — flat petrol, NOTHING DRAWN HERE.
   - TOP-RIGHT: cream small sans-serif "KEYNOTE TÉCNICA · 2026", letterspaced.
   - MID-UPPER: "QUATRO" in towering editorial serif (Didone/Bodoni) in gold #b38645, centered, large. Spell exactly Q-U-A-T-R-O.
   - Beneath QUATRO: cream bold sans-serif uppercase letterspaced "ASSASSINOS SILENCIOSOS".
   - Beneath that: cream regular sans-serif: "Cardio · Metabólico · Neuro · Câncer".
   - Lower-center on the fading petrol bottom, in cream bold serif italic: "Uma raiz. Quatro galhos."

2) BOTTOM 38% of canvas — CREAM #eae7da editorial block:
   - 1pt gold hairline at the top edge.
   - VERY SUBTLE wallpaper of tiny petrol dots at 4% opacity.
   - CENTER-LEFT: petrol regular sans-serif paragraph, 4 short lines:
     "Versão técnica do Cap. 2 do livro ANTES — para residentes, especialistas e congressos médicos. Os mecanismos compartilhados das quatro classes de doença que respondem por 80% da mortalidade depois dos 40 anos."
   - BOTTOM-LEFT (x≈90–180, y≈1925–2010): RESERVED FOR REAL P SYMBOL COMPOSITE — flat cream, NOTHING DRAWN HERE.
   - BOTTOM-RIGHT: three lines in petrol small caps sans:
     "CONTRATAÇÃO  ·  getamaralb002@gmail.com"
     "INSTAGRAM  ·  @drGetulioAmaralFilho"
     "INFO NO VERSO →"
   - Final 1pt petrol hairline along the bottom edge, beneath which in tiny italic sage: "Viva bem. Viva mais."

High resolution A5 portrait. Image size: 1152x2048.
"""

PROMPT_3 = BRAND_BLOCK + """
ASSET: A5 PORTRAIT FLYER BACK.

CANVAS: 1152x2048.

Background: uniform cream hex #eae7da with extremely subtle dot pattern at 3% opacity in petrol.

1) TOP BLOCK:
   - Small gold tag "SOBRE A KEYNOTE" uppercase, letterspaced.
   - Headline beneath in petrol bold editorial serif: "Quatro Doenças. Uma Raiz."
   - Subtitle in petrol regular sans-serif: "Inflamação crônica · resistência insulínica · disfunção mitocondrial."
   - 1pt gold hairline.

2) SYNOPSIS — petrol regular sans-serif, 5 comfortable lines:
   "Cardiovascular, metabólica, neurodegenerativa, câncer. Quatro classes de doença que dominam a mortalidade depois dos quarenta — e que compartilham três raízes biológicas comuns instaladas dez a vinte anos antes do diagnóstico. Esta keynote apresenta a fisiopatologia compartilhada, os marcadores que se movem antes da clínica e os três marcos farmacológicos da última década."

3) FRASE-PEDRA — own breathing space, centered, sage #92b8b4 italic editorial serif: "Quem ataca a raiz protege os quatro galhos."

4) THREE TAKEAWAYS — each with gold editorial-serif numeral "01"/"02"/"03" + bold petrol headline + one petrol regular line:
   01 · Aterosclerose como inflamação — não encanamento
   02 · Alzheimer como diabetes tipo 3 — APOE4 e p-tau217
   03 · Eixo cardio-reno-metabólico — SGLT2, finerenona, semaglutida

5) 1pt gold hairline.

6) CONTACT BLOCK — three columns at bottom:
   - LEFT: a clean petrol square with hairline gold border (~200x200) for QR placeholder, caption beneath "BIBLIOGRAFIA ANOTADA".
   - CENTER: "getamaralb002@gmail.com" / "plenyasaude.com.br".
   - RIGHT: "Dr. Getulio Amaral Filho" / "CRM-PR 21.876 · RQE 16.038" / "@drGetulioAmaralFilho".

7) BOTTOM EDGE:
   - LEFT (x≈90–180, y≈1955–2030): RESERVED FOR REAL P SYMBOL COMPOSITE — flat cream, NOTHING DRAWN HERE.
   - RIGHT: "VIVA BEM. VIVA MAIS." in petrol small caps letterspaced.

High resolution A5 portrait. Image size: 1152x2048.
"""

PROMPT_4 = BRAND_BLOCK + """
ASSET: 16:9 LANDSCAPE — cover slide of a 10-slide sales deck for medical contracting.

CANVAS: 2048x1152.

COMPOSITION: clean asymmetric split-screen.

LEFT HALF (50% width) — solid petrol #063b4f:
- TOP-LEFT (x≈80–440, y≈80–155): RESERVED FOR REAL WORDMARK COMPOSITE — flat petrol, NOTHING DRAWN HERE.
- Beneath that: tag "KEYNOTE TÉCNICA · 45 MIN" in cream uppercase sans, small.
- CENTER: hero headline stack:
  Line 1: "QUATRO" in towering Didone/Bodoni serif in gold #b38645, ~36% of left-panel height.
  Line 2: "ASSASSINOS SILENCIOSOS" in cream bold sans uppercase letterspaced.
  Line 3: "Cardio · Metabólico · Neuro · Câncer" in cream regular sans.
- LOWER: 1pt gold hairline (~200px), then three stacked cream lines:
  "Dr. Getulio Amaral Filho"
  "Nefrologista · 20 anos de prática · SBN"
  "CRM-PR 21.876 · RQE 16.038"
- BOTTOM-LEFT (x≈80–170, y≈1050–1130): RESERVED FOR REAL P SYMBOL COMPOSITE — flat petrol, NOTHING DRAWN HERE.

VERTICAL HAIRLINE: 0.5pt gold at exact 50% mark.

RIGHT HALF (50% width) — CINEMATIC STILL LIFE:
A single ancient gnarled tree at dusk against deep petrol sky, four prominent branches against the sky, deep roots exposed in cross-section at the base illuminated subtly in warm gold. NO PEOPLE. NO objects. Petrol+gold color grading. Painterly atmosphere.

High resolution landscape. Image size: 2048x1152.
"""

PROMPT_5 = BRAND_BLOCK + """
ASSET: 1:1 SQUARE SOCIAL POST — keynote announcement.

CANVAS: 1024x1024.

Background: deep petrol #063b4f.

TOP QUARTER (0–25%):
- TOP-CENTER (x≈340–680, y≈55–115): RESERVED FOR REAL WORDMARK COMPOSITE — flat petrol, NOTHING DRAWN HERE.
- Beneath: tag "KEYNOTE TÉCNICA · 2026" in cream uppercase sans-serif, small, letterspaced, centered.

MIDDLE BAND (25–72%):
- A horizontal cinematic photographic strip — single ancient gnarled tree at dusk, four prominent branches against deep petrol sky, exposed roots at base illuminated in gold. Duotone (petrol + gold highlights), heavy atmospheric grain, ~1/3 of total height.
- Overlaid in gold #b38645 editorial Didone serif, large but partially transparent over the image: "QUATRO" — centered horizontally. Spell exactly Q-U-A-T-R-O.

LOWER THIRD (72–100%):
- 1pt gold hairline.
- Cream bold sans-serif line, centered: "ASSASSINOS SILENCIOSOS — uma raiz, quatro galhos".
- Below, sage #92b8b4 italic editorial line: "uma keynote técnica do Dr. Getulio Amaral Filho".
- BOTTOM-LEFT corner: "Viva bem. Viva mais." in cream small caps sans, very small.
- BOTTOM-RIGHT corner (x≈920–990, y≈920–990): RESERVED FOR REAL P SYMBOL COMPOSITE.

High resolution square. Image size: 1024x1024.
"""

PROMPT_6 = BRAND_BLOCK + """
ASSET: 1:1 SQUARE SOCIAL POST — pure quote card (technical version).

CANVAS: 1024x1024.

Background: cream hex #eae7da, uniform, with extremely subtle dot pattern at 3% opacity in petrol.

Top:
- A single tag in gold #b38645 uppercase sans-serif, letterspaced, centered: "PALESTRA · QUATRO ASSASSINOS".

Center (the star):
- Gigantic opening quote mark (a single typographic curly quote ") in gold #b38645, partially transparent (~40% opacity), positioned top-left of the quote.
- The QUOTE TEXT in petrol #063b4f editorial serif (Playfair Display Bold), centered, three lines:
  Line 1: "Quem ataca a raiz"
  Line 2: "protege os quatro"
  Line 3: "galhos."
  Only "raiz" in gold #b38645 for emphasis; rest petrol.
- Generous breathing room.
- Closing quote mark on bottom-right, mirror, semi-transparent gold.

Below quote:
- Short 1pt gold hairline, centered, ~80px.
- Attribution in petrol sans small: "Dr. Getulio Amaral Filho  ·  Cap. 2 · ANTES (2026)".
- Beneath: "Viva bem. Viva mais." in gold small caps letterspaced, tiny.

BOTTOM-RIGHT corner (x≈920–990, y≈920–990): RESERVED FOR REAL P SYMBOL COMPOSITE.

High resolution square. Image size: 1024x1024.
"""

PROMPT_7 = BRAND_BLOCK + """
ASSET: 1:1 SQUARE SOCIAL POST — the four killers and three roots.

CANVAS: 1024x1024.

COMPOSITION: clean editorial layout. Background: deep petrol #063b4f.

TOP STRIP (0–10%):
- Centered tag in gold #b38645 uppercase sans-serif, letterspaced: "OS QUATRO ASSASSINOS · TRÊS RAÍZES".

UPPER HALF (10–55%) — THE FOUR BRANCHES:
- A horizontal arrangement of FOUR equal columns, each with:
  · A single editorial Didone gold serif numeral 01 / 02 / 03 / 04 at the top.
  · Beneath, in cream uppercase sans-serif bold letterspaced:
    01 — CARDIO
    02 — METABÓLICO
    03 — NEURO
    04 — CÂNCER
- A 0.5pt gold hairline horizontal across the canvas at y≈55% separating branches from roots.

LOWER HALF (55–88%) — THE THREE ROOTS:
- Centered single label in cream small caps letterspaced gold: "TRÊS RAÍZES BIOLÓGICAS COMUNS".
- Three short lines centered, each with petrol caption tag in gold uppercase + cream regular sans-serif body line, stacked vertically with breathing room:
  · INFLAMAÇÃO CRÔNICA — IL-6, TNF-α, hs-CRP
  · RESISTÊNCIA INSULÍNICA — instalada 5–10 anos antes da glicose subir
  · DISFUNÇÃO MITOCONDRIAL — capacidade oxidativa reduzida

BOTTOM STRIP (last 12% of canvas):
- 1pt gold hairline above.
- LEFT: small cream sans "Dr. Getulio Amaral Filho  ·  PALESTRA QUATRO ASSASSINOS".
- RIGHT (x≈920–990, y≈920–990): RESERVED FOR REAL P SYMBOL COMPOSITE.

High resolution square. Image size: 1024x1024.
"""

JOBS = [
    {
        "key": "02",
        "file": "02-flyer-a5-frente.PNG",
        "prompt": PROMPT_2,
        "size": "1152x2048",
        "overlays": [
            ("wordmark", WORDMARK_GOLD,  ( 80,  100,  460,  185)),
            ("symbol",   SYMBOL_GOLD,    ( 90, 1925,  180, 2010)),
        ],
    },
    {
        "key": "03",
        "file": "03-flyer-a5-verso.PNG",
        "prompt": PROMPT_3,
        "size": "1152x2048",
        "overlays": [
            ("symbol",   SYMBOL_INK,     ( 90, 1955,  180, 2030)),
        ],
    },
    {
        "key": "04",
        "file": "04-sales-deck-cover.PNG",
        "prompt": PROMPT_4,
        "size": "2048x1152",
        "overlays": [
            ("wordmark", WORDMARK_GOLD,  ( 80,   80,  440,  155)),
            ("symbol",   SYMBOL_GOLD,    ( 80, 1050,  170, 1130)),
        ],
    },
    {
        "key": "05",
        "file": "05-social-anuncio.PNG",
        "prompt": PROMPT_5,
        "size": "1024x1024",
        "overlays": [
            ("wordmark", WORDMARK_GOLD,  (340,   55,  680,  115)),
            ("symbol",   SYMBOL_GOLD,    (920,  920,  990,  990)),
        ],
    },
    {
        "key": "06",
        "file": "06-social-frase-pedra.PNG",
        "prompt": PROMPT_6,
        "size": "1024x1024",
        "overlays": [
            ("symbol",   SYMBOL_GOLD,    (920,  920,  990,  990)),
        ],
    },
    {
        "key": "07",
        "file": "07-social-quatro-galhos.PNG",
        "prompt": PROMPT_7,
        "size": "1024x1024",
        "overlays": [
            ("symbol",   SYMBOL_GOLD,    (920,  920,  990,  990)),
        ],
    },
]


def call_api(prompt: str, size: str) -> dict:
    env_path = ROOT / ".env"
    key = None
    for line in env_path.read_text().splitlines():
        line = line.strip()
        if line.startswith("OPENAI_API_KEY="):
            key = line.split("=", 1)[1].strip().strip('"').strip("'")
            break
    assert key
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


def composite_overlays(frame_path: pathlib.Path, overlays: list, out_path: pathlib.Path):
    frame = Image.open(frame_path).convert("RGBA")
    for kind, asset_path, box in overlays:
        asset = Image.open(asset_path).convert("RGBA")
        fit, xy = fit_into_box(asset, box)
        frame.alpha_composite(fit, xy)
        print(f"   ✓ {kind}@{xy} ({fit.size}) from {asset_path.name}", flush=True)
    frame.convert("RGB").save(out_path, "PNG", optimize=True)
    print(f"   ✓ saved: {out_path.name} ({out_path.stat().st_size//1024} KB)", flush=True)


def run_one(job: dict, compose_only: bool) -> int:
    final = OUT / job["file"]
    frame = OUT / job["file"].replace(".PNG", "_frame.PNG")
    tokens = 0
    if not compose_only:
        print(f"=== {job['file']}  ({job['size']}) ===", flush=True)
        try:
            data = call_api(job["prompt"], job["size"])
            b64 = data["data"][0].get("b64_json")
            frame.write_bytes(base64.b64decode(b64))
            tokens = data.get("usage", {}).get("total_tokens", 0)
            cost = tokens * 0.04 / 1000
            print(f"   ✓ frame saved ({frame.stat().st_size//1024} KB, {tokens} tokens, ≈ US$ {cost:.2f})", flush=True)
        except urllib.error.HTTPError as e:
            print(f"   ✗ HTTP {e.code}: {e.read().decode('utf-8','replace')[:600]}", flush=True)
            return 0
        except Exception as e:
            print(f"   ✗ {type(e).__name__}: {e}", flush=True)
            return 0
    if not frame.exists():
        print(f"   ✗ frame missing: {frame}", flush=True)
        return tokens
    composite_overlays(frame, job["overlays"], final)
    return tokens


def main():
    args = sys.argv[1:]
    compose_only = "--compose" in args
    keys = [a for a in args if a not in ("--compose",)]
    selected = [j for j in JOBS if not keys or any(k in j["key"] for k in keys)]
    total = 0
    for job in selected:
        total += run_one(job, compose_only)
    print(f"\nTotal tokens: {total}  ≈  US$ {total * 0.04 / 1000:.2f}")


if __name__ == "__main__":
    main()
