#!/usr/bin/env python3
"""Generate all marketing assets for palestra ANTES via gpt-image-2.

Mirrors the prompt density of the successful book figure scripts
(gen_book_fig41_v5.py, gen_fig81_landscape.py) but applies the Plenya
brand grammar (petrol/gold/cream + Nalieta/Playfair/sans-serif editorial).
"""
import json
import base64
import pathlib
import urllib.request
import urllib.error
import sys

env_path = pathlib.Path("/home/user/plenya/.env")
key = None
for line in env_path.read_text().splitlines():
    line = line.strip()
    if line.startswith("OPENAI_API_KEY="):
        key = line.split("=", 1)[1].strip().strip('"').strip("'")
        break
assert key, "OPENAI_API_KEY not found"

OUT = pathlib.Path("/home/user/plenya/docs/palestras/antes-janela-silenciosa/marketing/output")
OUT.mkdir(parents=True, exist_ok=True)

# =============================================================================
# SHARED BRAND GRAMMAR BLOCK (prefixed to every prompt for consistency)
# =============================================================================
BRAND_BLOCK = """
BRAND GRAMMAR — STRICT, NON-NEGOTIABLE:
- PLENYA brand palette (use ONLY these colors):
  · Petrol dark: hex #063b4f (dominant dark backgrounds)
  · Gold: hex #b38645 (accents, wordmark, numerals, hairlines)
  · Ocean: hex #417e8e (secondary dark, gradients)
  · Sage: hex #92b8b4 (organic mids, quiet accents)
  · Cream paper: hex #eae7da (light backgrounds, body on dark)
- Typography: 100% editorial. Mix of (a) a tall refined SERIF for headline numerals and the word "ANTES" (Didone/Playfair/Bodoni feel) and (b) a modern SANS-SERIF for body and tags (Söhne/Inter/GT America feel).
  The PLENYA wordmark should feel like a custom elegant serif with the letter "Y" subtly elongated (Nalieta-inspired).
- Monograms present: a tiny P + infinity loop symbol in gold used as discreet signature in corners/footers.
- Premium editorial magazine treatment — think Monocle, Kinfolk, The Gentleman's Journal. Never corporate, never wellness-cliché, never infographic-y.
- Hairlines ONLY for separators (0.5–1pt, gold or petrol at 20% opacity). ZERO heavy borders, ZERO rounded cards, ZERO drop shadows, ZERO gradients except subtle atmospheric ones on photography.
- Photography (when present) should be cinematic, editorial, muted color palette, shallow depth of field, shot on medium format feel. Never stock-happy.
- NO clip-art icons. NO emojis. NO decorative flourishes. NO tech-startup vibe. NO wellness-influencer vibe.
- Language is Brazilian Portuguese. All typography spelled correctly with proper accents (ó, ã, é, ç, í).
"""

# =============================================================================
# PROMPTS — one per asset
# =============================================================================

PROMPT_1_ONEPAGER = BRAND_BLOCK + """
ASSET: A4 PORTRAIT ONE-PAGER for a premium medical keynote, editorial quality, bookstore-ready.

COMPOSITION (top to bottom, vertical stack):

1) TOP BAND — 38% of page height, background hex #063b4f (petrol dark):
   - Ambient cinematic photography treatment behind the title: slight texture/grain suggestion, subtle vertical gradient from deep petrol at top fading to petrol + 10% ocean toward the bottom edge of the band.
   - Top-left: wordmark "PLENYA" in gold hex #b38645, letterspaced, small uppercase, elegant editorial serif (Nalieta-style), about 22pt.
   - Top-right, same baseline: tag "KEYNOTE · 45–60 MIN" in cream hex #eae7da uppercase sans-serif, letterspaced, small.
   - Center-left-anchored, the HERO HEADLINE:
     Line 1: "ANTES" — absolutely enormous, in a tall editorial SERIF, in warm gold #b38645, occupying dramatic presence (about 1/3 of the band height).
     Line 2 beneath: "A Janela Silenciosa" in cream bold sans-serif, refined.
     Line 3 beneath: "entre o Normal e o Ótimo" in cream regular sans-serif, same line but thinner.
   - Beneath headline, one short italic-feeling sans line in sage (#92b8b4) regular: "uma keynote do Dr. Getulio Amaral Filho".

2) HAIRLINE — a single horizontal 1pt gold (#b38645) line spanning 65% of page width, centered, separating the petrol band from the cream body below. ~80px of breathing room above and below this line.

3) CREAM BODY — 62% of page height, background hex #eae7da:

   a) Opening quote/hook in petrol bold sans-serif, large, left-aligned with generous outer margins (~12% each side):
      "Por que o seu check-up anual não está procurando o que vai te matar — e o que pedir para que procure."

   b) Synopsis paragraph in petrol regular sans-serif, 4 short lines, justified or left-aligned, comfortable leading:
      "A medicina brasileira se tornou excelente em tratar doença instalada — e ruim em prever o que está chegando. Entre a saúde que o laboratório reconhece e a doença que o médico trata existe uma janela silenciosa de 10 a 20 anos, onde a longevidade é construída ou perdida."

   c) Three-takeaway block, horizontally arranged OR stacked if more elegant. Each takeaway has:
      - A large gold numeral "01" / "02" / "03" in editorial serif
      - A short bold petrol headline in sans-serif beneath
      - One-line body in petrol regular beneath
      Takeaway 1:   "01 · Os 10 biomarcadores · que transformam triagem em mapa de prevenção."
      Takeaway 2:   "02 · O método AGIR · quatro pilares que definem as próximas duas décadas."
      Takeaway 3:   "03 · A Regra dos Dois · o único protocolo que produz resultado em 3 meses."

   d) Another 1pt gold hairline separator across 65% width, centered.

   e) AUTHOR BAND — two columns:
      - LEFT column (~30%): a clean petrol #063b4f circular placeholder, about 260px diameter, with a tiny sage label beneath "Dr. Getulio Amaral Filho" in small caps sans.
      - RIGHT column (~70%): a compact bio in petrol sans-serif, left-aligned:
        "Nefrologista com 20 anos de prática clínica. Coordena a Residência de Nefrologia da Santa Casa de Londrina. Fundador da Plenya — herdeira da Nefroclínica (40 anos). Autor de ANTES (2026, ISBN 978-65-02-06742-0)."
      Beneath bio, one gold line: "CRM-PR 21.876 · RQE 16.038".

   f) Bottom claim strip, centered, in petrol bold small caps: "VIVA BEM. VIVA MAIS." letterspaced, followed by a faint hairline.

   g) Footer row, small petrol-grey sans-serif at 60% opacity, left-to-right:
      "getamaralb002@gmail.com   ·   @drGetulioAmaralFilho   ·   plenyasaude.com.br"
      And at the far right: a tiny gold P-and-infinity monogram.

LAYOUT PRINCIPLES:
- Margins ~110–130px on all sides.
- Generous whitespace, editorial rhythm — the page must breathe.
- Typographic hierarchy must be unambiguous at a glance.
- The mix of the petrol band (image-like) at top and the cream editorial body at bottom must feel like a single coherent editorial piece, like the cover-plus-table-of-contents of a high-end magazine.

High resolution A4 portrait. Image size: 1152x2048.
"""

PROMPT_2_FLYER_FRONT = BRAND_BLOCK + """
ASSET: A5 PORTRAIT FLYER FRONT — a physical poster-grade piece for conference distribution and speaker kit. Editorial poster feel, not event-marketing feel.

COMPOSITION:

1) TOP 62% of the page — a CINEMATIC PHOTOGRAPHIC TREATMENT (not a literal photo, rather a painterly editorial scene):
   A quiet, early-morning physician's study at dawn. Dark warm-wood desk, a single brass task lamp casting a pool of light. On the desk, partly in shadow: a paper laboratory report with "DENTRO DA NORMALIDADE" faintly visible in green printed ink, a fountain pen resting beside it, a simple ceramic coffee cup, and a pair of reading glasses. A tall window on the right side lets in a slender beam of cold blue-grey dawn light, cutting across the room. The rest of the scene is submerged in deep petrol hex #063b4f darkness with warm gold highlights from the lamp. Shallow depth of field, medium-format cinematic mood, like a still from a Denis Villeneuve or Paolo Sorrentino film. Absolutely no people visible.
   Overlaid on this image, typography:
   - Top-left: wordmark "PLENYA" in gold #b38645 (Nalieta-style editorial serif), small, letterspaced.
   - Top-right: tiny cream sans-serif "KEYNOTE · 2026".
   - Mid-upper: the one-word headline "ANTES" in a towering editorial serif (Didone/Bodoni style) in gold #b38645, centered, absolutely enormous (about 35% of the image height for just that word).
   - Beneath "ANTES": thin cream sans-serif line "A Janela Silenciosa entre o Normal e o Ótimo".
   - Lower-center of this photo band, on the fading petrol bottom, in cream bold serif (smaller than hero, about 1/3 the size of ANTES): "O normal já é tarde demais."
   - Below that, cream small sans-serif attribution: "— tese da palestra".

2) BOTTOM 38% of the page — CREAM #eae7da editorial block:
   - 1pt gold hairline at the top edge separating from the photo band.
   - Subtle P-monogram pattern at 4% opacity in the cream background — very faint wallpaper of tiny Ps in petrol, barely visible, adds texture.
   - Center-left: a paragraph in petrol regular sans-serif, four short lines:
     "Uma keynote do Dr. Getulio Amaral Filho sobre a década silenciosa em que a longevidade é construída ou perdida. Baseada no livro ANTES (2026) e em 20 anos de prática clínica."
   - Bottom-left corner: small gold P+infinity monogram + wordmark "PLENYA" in small caps serif beneath.
   - Bottom-right corner: three lines in petrol small caps sans:
     "CONTRATAÇÃO  ·  getamaralb002@gmail.com"
     "INSTAGRAM  ·  @drGetulioAmaralFilho"
     "INFO NO VERSO →"
   - A final 1pt petrol hairline spans the bottom edge, beneath which in tiny italic sage: "Viva bem. Viva mais."

The piece must feel like a literary event poster in a bookstore window, not a marketing flyer.

High resolution A5 portrait. Image size: 1152x2048.
"""

PROMPT_3_FLYER_BACK = BRAND_BLOCK + """
ASSET: A5 PORTRAIT FLYER BACK — the verso of the poster flyer. Must share visual DNA with the front (same typography, same palette) but serve a different information function: content + contact.

COMPOSITION (plain cream paper aesthetic, editorial):

Background: uniform cream hex #eae7da with an extremely subtle repeating pattern of tiny P monograms at 3% opacity in petrol, creating a faint wallpaper texture.

1) TOP BLOCK:
   - Small gold tag "SOBRE A KEYNOTE" uppercase, letterspaced, petrol sans.
   - Headline beneath in petrol bold editorial serif (medium — about 1/6 of page width): "A Janela Silenciosa".
   - Subtitle in petrol regular sans-serif: "Por que o seu check-up anual não está procurando o que vai te matar."
   - 1pt gold hairline separator.

2) SYNOPSIS BLOCK — petrol regular sans-serif, 5 comfortable lines, left-aligned, body-copy leading:
   "A medicina brasileira se tornou excelente em tratar doença instalada — e ruim em prever o que está chegando. Entre a saúde que o laboratório reconhece e a doença que o médico trata existe uma janela silenciosa de 10 a 20 anos. Esta keynote apresenta o painel ampliado de biomarcadores que o check-up convencional ignora, o método AGIR de quatro pilares e a Regra dos Dois — o protocolo que funciona na prática."

3) FRASE-PEDRA — set in its own breathing space, centered, sage (#92b8b4) italic editorial serif, medium size: "Saúde não é sobre reagir. É sobre antecipar."

4) THREE TAKEAWAYS BLOCK — each with:
   - Gold editorial-serif numeral "01" / "02" / "03" large
   - Bold petrol headline in sans beside or beneath
   - One short petrol regular line
   Content exactly:
   01 · Os 10 biomarcadores que seu check-up esquece
   02 · Método AGIR — os quatro pilares da longevidade
   03 · A Regra dos Dois — o protocolo dos 3 meses

5) 1pt gold hairline separator.

6) CONTACT BLOCK, three columns at the bottom:
   - LEFT column: petrol block with a reserved CLEAN GOLD SQUARE OUTLINE (hairline only, no fill, about 200x200px) — this is a placeholder for a QR code — with the caption beneath "SOLICITAR PROPOSTA" in petrol small caps sans.
   - CENTER column: two stacked lines in petrol regular sans:
     "getamaralb002@gmail.com"
     "plenyasaude.com.br"
   - RIGHT column:
     "Dr. Getulio Amaral Filho"
     "CRM-PR 21.876 · RQE 16.038"
     "@drGetulioAmaralFilho"

7) Bottom edge:
   - Tiny gold P+infinity monogram on the left
   - "VIVA BEM. VIVA MAIS." in petrol small caps letterspaced on the right

The piece must read quickly, breathe, and feel unmistakably matched to the front of the flyer.

High resolution A5 portrait. Image size: 1152x2048.
"""

PROMPT_4_DECK_COVER = BRAND_BLOCK + """
ASSET: 16:9 LANDSCAPE SLIDE — the cover slide of a 10-slide sales deck. Think editorial book-cover-meets-cinematic-opening-title, not corporate PowerPoint.

COMPOSITION: a clean asymmetric split-screen.

LEFT HALF (50% width) — deep petrol hex #063b4f solid background:
- Top: wordmark "PLENYA" in gold #b38645 (Nalieta-style editorial serif), letterspaced, small.
- Beneath: tag "KEYNOTE · 45–60 MIN" in cream sans-serif uppercase, small, letterspaced.
- Center: hero headline stack:
  Line 1: "ANTES" in towering editorial serif (Didone/Bodoni) in gold #b38645 — the absolute visual king of the slide, occupying about 40% of the left-panel height.
  Line 2 beneath: "A Janela Silenciosa" in cream bold sans, medium large.
  Line 3: "entre o Normal e o Ótimo" in cream regular sans, smaller.
- Lower area: a 1pt gold hairline (about 200px long) then three stacked lines in cream:
  "Dr. Getulio Amaral Filho"
  "Nefrologista · 20 anos de prática"
  "CRM-PR 21.876 · RQE 16.038"
- Bottom-left: tiny gold P+infinity monogram.

VERTICAL HAIRLINE — a 0.5pt gold vertical line at the exact 50% mark, running floor-to-ceiling, as a subtle seam.

RIGHT HALF (50% width) — CINEMATIC STILL LIFE PHOTOGRAPH:
A low-angle, dimly lit scene of a physician's desk at dawn: dark warm walnut, a pool of gold lamplight falling on a paper laboratory report whose top band is visible with "DENTRO DA NORMALIDADE" printed in pale green legible ink. Beside the report, a fountain pen, a ceramic espresso cup still steaming, a slim hardcover book face-down (evocative of the ANTES book cover in petrol with gold title), and a brass-framed reading glasses. Background recedes into petrol darkness with a sliver of dawn-blue from a window at the far right edge. Shallow depth of field, editorial color grading (petrol + gold only), cinematic mood. Absolutely NO PEOPLE. Absolutely no logos visible on the book cover itself.

The whole slide should read as a film poster or the opening shot of a prestige documentary — NOT a marketing keynote cover.

High resolution landscape. Image size: 2048x1152.
"""

PROMPT_5_SOCIAL_ANUNCIO = BRAND_BLOCK + """
ASSET: 1:1 SQUARE SOCIAL POST — announcement of a new keynote. Destination: LinkedIn and Instagram feeds.

COMPOSITION: editorial split horizontal.

Background: deep petrol hex #063b4f.

TOP QUARTER (0–25%):
- Wordmark "PLENYA" in gold #b38645 (Nalieta-style serif), small, letterspaced, centered OR top-left.
- Beneath: tag "NOVA KEYNOTE · 2026" in cream sans-serif uppercase, small, letterspaced.

MIDDLE BAND (25–72%):
- A SINGLE CINEMATIC PHOTOGRAPHIC STRIP forms the visual anchor — a quiet, horizontal landscape image of a tall office window at dawn, cold blue-grey light spilling in, empty leather chair visible in faint silhouette, no person. Treated in duotone (petrol #063b4f + gold #b38645 highlights), heavy atmospheric grain, about 1/3 of total image height.
- Overlaid on top of the photographic strip, in gold #b38645 editorial serif (Didone), absolutely enormous but allowed to be partially transparent over the image: "ANTES" — the word alone, centered horizontally.

LOWER THIRD (72–100%):
- 1pt gold hairline separator.
- Beneath: cream bold sans-serif line, centered or left-aligned: "A Janela Silenciosa entre o Normal e o Ótimo".
- Below that, sage #92b8b4 italic editorial line: "uma keynote do Dr. Getulio Amaral Filho".
- Bottom-right corner: tiny gold P+infinity monogram.
- Bottom-left corner: "Viva bem. Viva mais." in cream small caps sans, letterspaced, very small.

The composition must feel like a Monocle magazine cover teaser, not a social-media post template.

High resolution square. Image size: 1024x1024.
"""

PROMPT_6_SOCIAL_QUOTE = BRAND_BLOCK + """
ASSET: 1:1 SQUARE SOCIAL POST — pure quote card. The tese central of the entire keynote, condensed to one sentence, as visually arresting as possible.

COMPOSITION:

Background: cream hex #eae7da, uniform, with an EXTREMELY subtle pattern of tiny P monograms at 3% opacity in petrol — just visible as a whisper of texture.

Top:
- A single tag in gold #b38645 uppercase sans-serif, letterspaced, centered: "PALESTRA · ANTES".

Center (the star of the composition):
- A gigantic opening quote mark "“" in gold #b38645, partially transparent (about 40% opacity), positioned top-left of the quote as if pulled from a book, taking up about 1/5 of the entire canvas height.
- The QUOTE TEXT itself, in petrol hex #063b4f, editorial serif (Playfair Display Bold or similar), centered, set across two or three lines with dramatic line breaks:
  Line 1: "Normal não é"
  Line 2: "o mesmo que"
  Line 3: "ótimo."
  (Only the word "ótimo." is set in gold #b38645 for emphasis; the rest in petrol.)
- Generous breathing room around the text.
- A closing quote mark "”" in gold semi-transparent on the bottom-right, mirror of the opening one.

Below quote:
- A short 1pt gold hairline, centered, about 80px long.
- Then attribution line in petrol sans-serif small: "Dr. Getulio Amaral Filho  ·  tese da palestra ANTES".
- Beneath: "Viva bem. Viva mais." in gold small caps letterspaced, tiny.

Bottom-right corner: a tiny gold P+infinity monogram.

The result must feel like a framed literary quote in an editorial magazine spread — never an Instagram-template quote.

High resolution square. Image size: 1024x1024.
"""

PROMPT_7_SOCIAL_NORMAL_OTIMO = BRAND_BLOCK + """
ASSET: 1:1 SQUARE SOCIAL POST — a TECHNICAL TEASER contrasting "normal" vs "ótimo" biomarker ranges. Designed to stop the scroll with numerical substance.

COMPOSITION: clean horizontal split into two equal halves, top and bottom.

TOP HALF (0–50%) — CREAM #eae7da background:
- Top: small gold tag "NORMAL  ·  O QUE SEU LAB IMPRIME" uppercase sans-serif, letterspaced.
- Beneath, in petrol #063b4f regular editorial sans-serif, aligned as a clean vertical column of 4 biomarker rows. Each row is: BIOMARKER NAME (petrol bold) on the left, VALUE (petrol regular) on the right:
  · INSULINA DE JEJUM  —  < 25 µIU/mL
  · APOB  —  < 130 mg/dL
  · HBA1C  —  < 6,5%
  · HS-CRP  —  < 3,0 mg/L
- Plenty of whitespace; each row breathes.

MIDDLE DIVIDING BAND — a narrow horizontal strip (about 8% of total height) centered at y=50%:
- Background: gold #b38645.
- Centered inline text in petrol #063b4f bold sans-serif, small caps, letterspaced: "← UMA DÉCADA DE DIFERENÇA EM RISCO →".

BOTTOM HALF (58–100%) — DEEP PETROL #063b4f background:
- Top: small gold tag "ÓTIMO  ·  O QUE A EVIDÊNCIA INDICA" uppercase sans-serif, letterspaced.
- Beneath, in cream #eae7da bold editorial sans-serif, same 4 biomarker rows but with optimal ranges, emphasized:
  · INSULINA DE JEJUM  —  < 8 µIU/mL
  · APOB  —  < 90 mg/dL
  · HBA1C  —  ≤ 5,4%
  · HS-CRP  —  < 1,0 mg/L
- The values specifically ("< 8", "< 90", "≤ 5,4%", "< 1,0") set in gold #b38645 for emphasis.

BOTTOM STRIP (last 8% of the canvas, inside the petrol block):
- Left: small cream sans "Dr. Getulio Amaral Filho  ·  PALESTRA ANTES".
- Right: tiny gold P+infinity monogram.
- Ultra-fine gold hairline beneath the biomarker rows separating them from this footer.

Typography hierarchy must make it obvious, in 2 seconds, that the PETROL half values are the goal and the CREAM half values are the trap.

High resolution square. Image size: 1024x1024.
"""

# =============================================================================
# RUNNER
# =============================================================================

JOBS = [
    ("01-onepager-a4.PNG",            PROMPT_1_ONEPAGER,            "1152x2048"),
    ("02-flyer-a5-frente.PNG",        PROMPT_2_FLYER_FRONT,         "1152x2048"),
    ("03-flyer-a5-verso.PNG",         PROMPT_3_FLYER_BACK,          "1152x2048"),
    ("04-sales-deck-cover.PNG",       PROMPT_4_DECK_COVER,          "2048x1152"),
    ("05-social-anuncio.PNG",         PROMPT_5_SOCIAL_ANUNCIO,      "1024x1024"),
    ("06-social-frase-pedra.PNG",     PROMPT_6_SOCIAL_QUOTE,        "1024x1024"),
    ("07-social-normal-vs-otimo.PNG", PROMPT_7_SOCIAL_NORMAL_OTIMO, "1024x1024"),
]


def call(prompt: str, size: str) -> dict:
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


def main():
    targets = sys.argv[1:] if len(sys.argv) > 1 else None
    total_tokens = 0
    for filename, prompt, size in JOBS:
        if targets and not any(t in filename for t in targets):
            continue
        print(f"=== {filename}  ({size}) ===", flush=True)
        try:
            data = call(prompt, size)
            b64 = data["data"][0].get("b64_json")
            out = OUT / filename
            out.write_bytes(base64.b64decode(b64))
            tokens = data.get("usage", {}).get("total_tokens", 0)
            total_tokens += tokens
            size_kb = out.stat().st_size // 1024
            cost = tokens * 0.04 / 1000
            print(f"   ✓ {out.name}  ({size_kb} KB, {tokens} tokens, ≈ US$ {cost:.2f})",
                  flush=True)
        except urllib.error.HTTPError as e:
            print(f"   ✗ HTTP {e.code}: {e.read().decode('utf-8', 'replace')[:600]}",
                  flush=True)
        except Exception as e:
            print(f"   ✗ {type(e).__name__}: {e}", flush=True)
    print(f"\nTotal tokens: {total_tokens}  ≈  US$ {total_tokens * 0.04 / 1000:.2f}")


if __name__ == "__main__":
    main()
