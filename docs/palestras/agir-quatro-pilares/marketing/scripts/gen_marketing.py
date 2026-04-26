#!/usr/bin/env python3
"""AGIR marketing — pieces 02–07 with hybrid pipeline (gpt-image-2 + PIL).

Mirrors gen_marketing_v3.py from the ANTES talk, adapted to AGIR:
- Hero word AGIR (not ANTES) — same Didone treatment.
- Cinematic scenes use the four-beams-of-light metaphor for the four pillars.
- Social comparative replaces "Normal vs Ótimo" with the matrix
  Impacto × Esforço (the slide-rainha of the talk).
- Quote card uses the AGIR-specific frase-pedra: "Foco não é fazer mais.
  Foco é decidir o que não fazer."
- Sales deck synopsis pivots from biomarkers to method.

Run:
  python3 gen_marketing.py            # full pipeline (all pieces)
  python3 gen_marketing.py 02 04      # only those numeric prefixes
  python3 gen_marketing.py --compose  # skip API, recomposite from frames
"""
import json
import base64
import pathlib
import urllib.request
import urllib.error
import sys
from PIL import Image

ROOT = pathlib.Path("/home/user/plenya")
OUT = ROOT / "docs/palestras/agir-quatro-pilares/marketing/output"
OUT.mkdir(parents=True, exist_ok=True)

WORDMARK_GOLD = ROOT / "apps/site/public/brand/wordmark/gold.png"
WORDMARK_CREAM = ROOT / "apps/site/public/brand/wordmark/cream.png"
SYMBOL_GOLD = ROOT / "apps/site/public/brand/symbol/gold.png"
SYMBOL_CREAM = ROOT / "apps/site/public/brand/symbol/cream.png"
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
- DO NOT draw the word "PLENYA" anywhere. Not as a wordmark, not as text, not even small.
- DO NOT draw any "Y" with a diacritic, accent, macron, umlaut, or any mark on it. There is no special Y in this brand.
- DO NOT draw any P symbol, P monogram, P logo, P + infinity (∞) combination, or any infinity symbol of any kind.
- DO NOT draw any logo whatsoever — no medical cross, no caduceus, no abstract mark, nothing.
- The reserved zones (described below per piece) must remain clean and flat colored — wordmark and symbol will be composited afterward from the official PNG files. If you draw anything in those zones it will collide with the real logo and ruin the piece.
- Spell "AGIR" exactly with classic A-G-I-R uppercase — no special letterforms, no diacritics on the I, no ligatures.
"""

# =============================================================================
# PROMPTS — pieces 02 to 07
# =============================================================================

PROMPT_2 = BRAND_BLOCK + """
ASSET: A5 PORTRAIT FLYER FRONT — physical poster-grade piece. Editorial poster feel, not event-marketing.

CANVAS: 1152x2048.

COMPOSITION:

1) TOP 62% of the canvas — CINEMATIC PHOTOGRAPHIC TREATMENT:
   A quiet interior at dawn, deep petrol shadows. A tall window with FOUR slender vertical mullions casting FOUR PARALLEL BEAMS of warm golden light slanting across a petrol stone floor. Each beam slightly different in intensity — the four together suggesting structure without literal labels. Heavy atmospheric grain, painterly. NO PEOPLE. NO furniture. NO objects. NO numbers, letters, or labels visible inside the scene.

   OVERLAID typography:
   - TOP-LEFT (x≈80–460, y≈100–185): RESERVED FOR REAL WORDMARK COMPOSITE — flat petrol, NOTHING DRAWN HERE.
   - TOP-RIGHT: cream small sans-serif "KEYNOTE · 2026", letterspaced.
   - MID-UPPER: "AGIR" in towering editorial serif (Didone/Bodoni) in gold #b38645, centered, absolutely huge (~35% of band height). Spell exactly A-G-I-R.
   - Beneath AGIR: thin cream sans-serif "Os Quatro Pilares da Longevidade Prática".
   - Lower-center on the fading petrol bottom, in cream bold serif: "Saber não é fazer."
   - Below that: cream small sans-serif attribution: "— tese da palestra".

2) BOTTOM 38% of canvas — CREAM #eae7da editorial block:
   - 1pt gold hairline at the top edge separating from the photo band.
   - VERY SUBTLE wallpaper of tiny petrol dots at 4% opacity (no P shapes, just dots).
   - CENTER-LEFT: petrol regular sans-serif paragraph, 4 short lines:
     "Uma keynote do Dr. Getulio Amaral Filho sobre o método operacional que transforma intenção em resultado em três meses. Quatro pilares — A·G·I·R — e a Regra dos Dois."
   - BOTTOM-LEFT (x≈90–180, y≈1925–2010): RESERVED FOR REAL P SYMBOL COMPOSITE — flat cream, NOTHING DRAWN HERE.
   - BOTTOM-RIGHT: three lines in petrol small caps sans:
     "CONTRATAÇÃO  ·  getamaralb002@gmail.com"
     "INSTAGRAM  ·  @drGetulioAmaralFilho"
     "INFO NO VERSO →"
   - Final 1pt petrol hairline along the bottom edge, beneath which in tiny italic sage: "Viva bem. Viva mais."

The piece must feel like a literary event poster, not a marketing flyer.

High resolution A5 portrait. Image size: 1152x2048.
"""

PROMPT_3 = BRAND_BLOCK + """
ASSET: A5 PORTRAIT FLYER BACK — verso of poster flyer.

CANVAS: 1152x2048.

Background: uniform cream hex #eae7da with extremely subtle dot pattern at 3% opacity in petrol (no P shapes).

1) TOP BLOCK:
   - Small gold tag "SOBRE A KEYNOTE" uppercase, letterspaced.
   - Headline beneath in petrol bold editorial serif: "Os Quatro Pilares".
   - Subtitle in petrol regular sans-serif: "Atividade física · Gestão clínica · Integração mente-corpo · Ritmo circadiano."
   - 1pt gold hairline.

2) SYNOPSIS — petrol regular sans-serif, 5 comfortable lines:
   "Toda semana alguém termina uma palestra de bem-estar tomando notas — e na segunda-feira não faz nada. Não por preguiça: aprendeu mais do que conseguia executar. AGIR é o oposto disso. Um framework operacional de quatro pilares e um protocolo — a Regra dos Dois — que produz resultado mensurável em três meses, sem sobrecarga, sem fórmula mágica."

3) FRASE-PEDRA — own breathing space, centered, sage #92b8b4 italic editorial serif: "Foco não é fazer mais. Foco é decidir o que não fazer."

4) THREE TAKEAWAYS — each with gold editorial-serif numeral "01"/"02"/"03" + bold petrol headline + one petrol regular line:
   01 · O método AGIR — quatro letras que organizam tudo
   02 · Quatro casos clínicos reais — Marcos · Paulo · Ana · André
   03 · A Regra dos Dois — dois focos por trimestre, alto impacto e baixo esforço

5) 1pt gold hairline.

6) CONTACT BLOCK — three columns at bottom:
   - LEFT: a clean petrol square with hairline gold border (~200x200) for QR placeholder, caption beneath "SOLICITAR PROPOSTA".
   - CENTER: "getamaralb002@gmail.com" / "plenyasaude.com.br".
   - RIGHT: "Dr. Getulio Amaral Filho" / "CRM-PR 21.876 · RQE 16.038" / "@drGetulioAmaralFilho".

7) BOTTOM EDGE:
   - LEFT (x≈90–180, y≈1955–2030): RESERVED FOR REAL P SYMBOL COMPOSITE — flat cream, NOTHING DRAWN HERE.
   - RIGHT: "VIVA BEM. VIVA MAIS." in petrol small caps letterspaced.

High resolution A5 portrait. Image size: 1152x2048.
"""

PROMPT_4 = BRAND_BLOCK + """
ASSET: 16:9 LANDSCAPE — cover slide of a 10-slide sales deck.

CANVAS: 2048x1152.

COMPOSITION: clean asymmetric split-screen.

LEFT HALF (50% width) — solid petrol #063b4f:
- TOP-LEFT (x≈80–440, y≈80–155): RESERVED FOR REAL WORDMARK COMPOSITE — flat petrol, NOTHING DRAWN HERE.
- Beneath that: tag "KEYNOTE · 45–60 MIN" in cream uppercase sans, small.
- CENTER: hero headline stack:
  Line 1: "AGIR" in towering Didone/Bodoni serif in gold #b38645, ~40% of left-panel height. Spell exactly A-G-I-R.
  Line 2: "Os Quatro Pilares" in cream bold sans.
  Line 3: "da Longevidade Prática" in cream regular sans.
- LOWER: 1pt gold hairline (~200px), then three stacked cream lines:
  "Dr. Getulio Amaral Filho"
  "Nefrologista · 20 anos de prática"
  "CRM-PR 21.876 · RQE 16.038"
- BOTTOM-LEFT (x≈80–170, y≈1050–1130): RESERVED FOR REAL P SYMBOL COMPOSITE — flat petrol, NOTHING DRAWN HERE.

VERTICAL HAIRLINE: 0.5pt gold at exact 50% mark, floor to ceiling.

RIGHT HALF (50% width) — CINEMATIC STILL LIFE:
A low-angle, dimly lit interior at dawn. A tall window with four mullions on the right edge of the panel casts FOUR PARALLEL BEAMS of warm golden light slanting across a deep petrol stone floor. Each beam slightly different in intensity — visual metaphor for the four pillars without literal labels. Background recedes into petrol darkness with subtle warm highlights. Petrol+gold color grading. NO PEOPLE. NO furniture. NO objects. No logos visible.

High resolution landscape. Image size: 2048x1152.
"""

PROMPT_5 = BRAND_BLOCK + """
ASSET: 1:1 SQUARE SOCIAL POST — keynote announcement.

CANVAS: 1024x1024.

Background: deep petrol #063b4f.

TOP QUARTER (0–25%):
- TOP-CENTER (x≈340–680, y≈55–115): RESERVED FOR REAL WORDMARK COMPOSITE — flat petrol, NOTHING DRAWN HERE.
- Beneath: tag "NOVA KEYNOTE · 2026" in cream uppercase sans-serif, small, letterspaced, centered.

MIDDLE BAND (25–72%):
- A horizontal cinematic photographic strip — tall window at dawn casting four parallel beams of warm golden light slanting across a petrol floor, NO PERSON, NO objects. Duotone (petrol + gold highlights), heavy atmospheric grain, ~1/3 of total height.
- Overlaid in gold #b38645 editorial Didone serif, large but partially transparent over the image: "AGIR" — centered horizontally. Spell exactly A-G-I-R.

LOWER THIRD (72–100%):
- 1pt gold hairline.
- Cream bold sans-serif line, centered: "Os Quatro Pilares da Longevidade Prática".
- Below, sage #92b8b4 italic editorial line: "uma keynote do Dr. Getulio Amaral Filho".
- BOTTOM-LEFT corner: "Viva bem. Viva mais." in cream small caps sans, very small.
- BOTTOM-RIGHT corner (x≈920–990, y≈920–990): RESERVED FOR REAL P SYMBOL COMPOSITE — flat petrol, NOTHING DRAWN HERE.

High resolution square. Image size: 1024x1024.
"""

PROMPT_6 = BRAND_BLOCK + """
ASSET: 1:1 SQUARE SOCIAL POST — pure quote card.

CANVAS: 1024x1024.

Background: cream hex #eae7da, uniform, with extremely subtle dot pattern at 3% opacity in petrol (no P shapes).

Top:
- A single tag in gold #b38645 uppercase sans-serif, letterspaced, centered: "PALESTRA · AGIR".

Center (the star):
- Gigantic opening quote mark (a single typographic curly quote ") in gold #b38645, partially transparent (~40% opacity), positioned top-left of the quote, taking ~1/5 of canvas height.
- The QUOTE TEXT in petrol #063b4f editorial serif (Playfair Display Bold), centered, three lines with dramatic line breaks:
  Line 1: "Foco não é"
  Line 2: "fazer mais."
  Line 3: "É decidir o que não fazer."
  Only "decidir" in gold #b38645 for emphasis; rest petrol.
- Generous breathing room.
- Closing quote mark on bottom-right, mirror of the opening, semi-transparent gold.

Below quote:
- Short 1pt gold hairline, centered, ~80px.
- Attribution in petrol sans small: "Dr. Getulio Amaral Filho  ·  Regra dos Dois · palestra AGIR".
- Beneath: "Viva bem. Viva mais." in gold small caps letterspaced, tiny.

BOTTOM-RIGHT corner (x≈920–990, y≈920–990): RESERVED FOR REAL P SYMBOL COMPOSITE — flat cream, NOTHING DRAWN HERE.

High resolution square. Image size: 1024x1024.
"""

PROMPT_7 = BRAND_BLOCK + """
ASSET: 1:1 SQUARE SOCIAL POST — the four pillars at a glance.

CANVAS: 1024x1024.

COMPOSITION: a clean editorial 2×2 grid of four pillars, each anchored to a single number from 20 years of clinical practice. Background: deep petrol #063b4f. A thin gold cross splits the canvas into four equal quadrants (0.5pt gold hairlines, vertical and horizontal, exactly at 50%).

TOP STRIP (0–10%):
- Centered tag in gold #b38645 uppercase sans-serif, letterspaced: "AGIR · OS QUATRO PILARES".

EACH OF THE FOUR QUADRANTS (10–88% area, divided in 2×2):
Each quadrant has the SAME composition:
- A LARGE letter at the top (A / G / I / R) in towering editorial Didone serif in gold #b38645, occupying about 35% of the quadrant height. Spell exactly A then G then I then R, one per quadrant.
- Beneath, in cream uppercase sans-serif small: a short label.
- Beneath that, in cream bold editorial sans, a single number-anchor.
- Beneath, in cream regular sans-serif tiny, a one-line caption.

QUADRANT TOP-LEFT (A):
- Letter: A
- Label: ATIVIDADE FÍSICA
- Number: HR 5,04
- Caption: sedentarismo · maior fator isolado de mortalidade

QUADRANT TOP-RIGHT (G):
- Letter: G
- Label: GESTÃO CLÍNICA
- Number: < 8 µIU/mL
- Caption: insulina ótima vs. < 25 "normal" do laboratório

QUADRANT BOTTOM-LEFT (I):
- Letter: I
- Label: INTEGRAÇÃO MENTE-CORPO
- Number: −33%
- Caption: ikigai · redução de mortalidade all-cause

QUADRANT BOTTOM-RIGHT (R):
- Letter: R
- Label: RITMO CIRCADIANO
- Number: 20–48%
- Caption: regularidade do sono · redução de mortalidade

BOTTOM STRIP (last 12% of canvas):
- 1pt gold hairline above.
- LEFT: small cream sans "Dr. Getulio Amaral Filho  ·  PALESTRA AGIR".
- RIGHT (x≈920–990, y≈920–990): RESERVED FOR REAL P SYMBOL COMPOSITE — flat petrol, NOTHING DRAWN HERE.

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
        "file": "07-social-quatro-pilares.PNG",
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
