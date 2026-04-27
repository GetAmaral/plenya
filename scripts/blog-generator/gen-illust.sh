#!/usr/bin/env bash
# Generate an ILLUSTRATIVE in-content blog image — meant to clarify a
# specific concept the reader is meeting in that paragraph, not just
# set mood. Lighter on brand palette than gen-image.sh; more variety in
# composition, subject, and color so the article doesn't feel uniform.
#
# Usage: gen-illust.sh <slug> <name> "<prompt>"
#   <name> becomes the filename: /images/blog/<slug>/<name>.webp
#         e.g. concept-1, diagram, comparison, anatomy
# Output path is printed on success (1024x1024 webp).
set -euo pipefail

SLUG="${1:?slug required}"
NAME="${2:?name required (e.g. concept-1)}"
PROMPT="${3:?prompt required}"

# Sanitize name (no path separators)
NAME="${NAME//\//-}"

if [ -z "${OPENAI_API_KEY:-}" ]; then
  OPENAI_API_KEY="$(grep -E '^OPENAI_API_KEY=' /home/user/plenya/apps/api/.env | cut -d= -f2-)"
fi

OUT_DIR="/home/user/plenya/apps/site/public/images/blog/$SLUG"
mkdir -p "$OUT_DIR"
OUT="$OUT_DIR/$NAME.webp"

# Lighter style suffix — informative editorial illustration that varies
# more across the blog. Still warm and cohesive, but not locked into the
# petrol/gold mood of the hero/inline images.
STYLE_SUFFIX=" Illustrative editorial photograph or still-life that clearly communicates the concept above to a reader. Use whatever color palette best fits the subject — natural greens, muted earth tones, warm whites, and clinical neutrals are all welcome; do not force teal-and-gold. Soft natural lighting, real textures, shallow depth of field where appropriate. Composition leaves space for breathing but keeps the subject readable and recognizable. No text, no numbers, no logos, no watermarks, no recognizable celebrity faces."

FULL_PROMPT="${PROMPT}.${STYLE_SUFFIX}"

TMP_JSON="$(mktemp)"
trap 'rm -f "$TMP_JSON"' EXIT

PAYLOAD=$(jq -n \
  --arg model "gpt-image-1" \
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
  head -c 500 "$TMP_JSON" >&2
  exit 1
fi

jq -r '.data[0].b64_json' "$TMP_JSON" | base64 -d > "$OUT"

if [ ! -s "$OUT" ]; then
  echo "empty output file" >&2
  exit 1
fi

if command -v convert >/dev/null 2>&1; then
  TMP_OPT="${OUT}.opt.webp"
  convert "$OUT" -strip -quality 82 -define webp:method=6 "$TMP_OPT" 2>/dev/null && mv "$TMP_OPT" "$OUT" || rm -f "$TMP_OPT"
fi

echo "$OUT"
