#!/usr/bin/env bash
# Generate an INFOGRAPHIC-style figure (chart/diagram) like the ones in
# the book "Antes — A Janela Silenciosa". Pass an extremely detailed
# prompt that specifies title text, axis labels, every numeric value,
# color zones, and source citation — gpt-image-2 will render it.
#
# Usage: gen-figure.sh <slug> <name> "<prompt>"
# Output: /images/blog/<slug>/<name>.webp (1024x1024)
set -euo pipefail

SLUG="${1:?slug required}"
NAME="${2:?name required}"
PROMPT="${3:?prompt required}"
NAME="${NAME//\//-}"

if [ -z "${OPENAI_API_KEY:-}" ]; then
  OPENAI_API_KEY="$(grep -E '^OPENAI_API_KEY=' /home/user/plenya/apps/api/.env | cut -d= -f2-)"
fi

OUT_DIR="/home/user/plenya/apps/site/public/images/blog/$SLUG"
mkdir -p "$OUT_DIR"
OUT="$OUT_DIR/$NAME.webp"

# Style suffix tuned for the book "Antes" infographic look:
# flat 2D editorial chart on cream paper, Plenya brand palette,
# clean sans-serif typography, color-coded zones, bottom source line.
STYLE_SUFFIX=" Render as a clean flat 2D editorial infographic in the visual style of a serious medical/longevity book. Background: warm cream paper color #f5f1e8 with very subtle paper grain. Brand palette: deep teal-petrol #0a3d3a (primary lines, headlines, axis ink), antique gold-mustard #b48a4a (highlight bars, gold zone band), bordo dark red #6b1f2a (warning zone, alert highlights), soft sage green #c5d4ba (optimal zone band), dusty pink #e8c8c4 (risk zone band), warm gray #6b6157 (secondary text and gridlines). Typography: title at top in a tasteful serif (Source Serif 4 or similar) deep petrol color. Axis labels and data labels in clean sans-serif (Inter-style) warm gray. Numeric values placed exactly where described, large and clearly readable. Faint horizontal gridlines if a chart. A small italic source citation at bottom in warm gray. Generous internal padding, balanced negative space, centered composition. NO photographic elements, NO 3D, NO shadows, NO glossy effects. Flat, editorial, like a textbook diagram. Text MUST be readable and accurately spelled in Brazilian Portuguese as specified."

FULL_PROMPT="${PROMPT}.${STYLE_SUFFIX}"

TMP_JSON="$(mktemp)"
trap 'rm -f "$TMP_JSON"' EXIT

PAYLOAD=$(jq -n \
  --arg model "gpt-image-2" \
  --arg prompt "$FULL_PROMPT" \
  --arg size "1024x1024" \
  --arg quality "high" \
  '{model:$model, prompt:$prompt, size:$size, quality:$quality, n:1, output_format:"webp"}')

HTTP_CODE=$(curl -s -o "$TMP_JSON" -w "%{http_code}" \
  https://api.openai.com/v1/images/generations \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -H "Content-Type: application/json" \
  -d "$PAYLOAD")

if [ "$HTTP_CODE" != "200" ]; then
  echo "OpenAI API error ($HTTP_CODE):" >&2
  head -c 800 "$TMP_JSON" >&2
  exit 1
fi

jq -r '.data[0].b64_json' "$TMP_JSON" | base64 -d > "$OUT"
[ -s "$OUT" ] || { echo "empty output file" >&2; exit 1; }

if command -v convert >/dev/null 2>&1; then
  TMP_OPT="${OUT}.opt.webp"
  convert "$OUT" -strip -quality 85 -define webp:method=6 "$TMP_OPT" 2>/dev/null && mv "$TMP_OPT" "$OUT" || rm -f "$TMP_OPT"
fi

echo "$OUT"
