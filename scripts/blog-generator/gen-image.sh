#!/usr/bin/env bash
# Generate a blog image via OpenAI gpt-image-2.
# Usage: gen-image.sh <slug> <kind> "<prompt>"
#   kind: hero (1536x1024) | inline (1024x1024)
# Output: /home/user/plenya/apps/site/public/images/blog/<slug>/<kind>.webp
set -euo pipefail

SLUG="${1:?slug required}"
KIND="${2:?kind required: hero|inline}"
PROMPT="${3:?prompt required}"

case "$KIND" in
  hero)   SIZE="1536x1024" ;;
  inline) SIZE="1024x1024" ;;
  *) echo "kind must be hero or inline" >&2; exit 1 ;;
esac

# Load API key from monorepo .env
if [ -z "${OPENAI_API_KEY:-}" ]; then
  OPENAI_API_KEY="$(grep -E '^OPENAI_API_KEY=' /home/user/plenya/apps/api/.env | cut -d= -f2-)"
fi

OUT_DIR="/home/user/plenya/apps/site/public/images/blog/$SLUG"
mkdir -p "$OUT_DIR"
OUT="$OUT_DIR/$KIND.webp"

# Fixed style suffix to keep visual coherence across the blog.
STYLE_SUFFIX=" Editorial photography in muted, warm tones consistent with a sophisticated medical-longevity brand. Color palette: deep teal-petrol (#0a3d3a), antique gold (#b48a4a), cream paper (#f5f1e8), soft warm shadows. Cinematic lighting, shallow depth of field, natural film grain. No text, no logos, no watermarks. Composition with generous negative space."

FULL_PROMPT="${PROMPT}.${STYLE_SUFFIX}"

TMP_JSON="$(mktemp)"
trap 'rm -f "$TMP_JSON"' EXIT

# jq builds the JSON safely (handles quotes, newlines)
PAYLOAD=$(jq -n \
  --arg model "gpt-image-2" \
  --arg prompt "$FULL_PROMPT" \
  --arg size "$SIZE" \
  --arg quality "high" \
  '{model:$model, prompt:$prompt, size:$size, quality:$quality, n:1, output_format:"webp"}')

HTTP_CODE=$(curl -s -o "$TMP_JSON" -w "%{http_code}" \
  https://api.openai.com/v1/images/generations \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -H "Content-Type: application/json" \
  -d "$PAYLOAD")

if [ "$HTTP_CODE" != "200" ]; then
  echo "OpenAI API error ($HTTP_CODE):" >&2
  cat "$TMP_JSON" >&2
  exit 1
fi

# gpt-image-2 returns base64 in data[0].b64_json
jq -r '.data[0].b64_json' "$TMP_JSON" | base64 -d > "$OUT"

if [ ! -s "$OUT" ]; then
  echo "empty output file" >&2
  exit 1
fi

# Optimize: re-encode WebP at quality 82 and downscale hero to max 1600w (next/image will handle responsive variants).
if command -v convert >/dev/null 2>&1; then
  TMP_OPT="${OUT}.opt.webp"
  if [ "$KIND" = "hero" ]; then
    convert "$OUT" -resize '1600x>' -strip -quality 82 -define webp:method=6 "$TMP_OPT" 2>/dev/null && mv "$TMP_OPT" "$OUT" || rm -f "$TMP_OPT"
  else
    convert "$OUT" -strip -quality 82 -define webp:method=6 "$TMP_OPT" 2>/dev/null && mv "$TMP_OPT" "$OUT" || rm -f "$TMP_OPT"
  fi
fi

echo "$OUT"
