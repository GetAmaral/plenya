#!/usr/bin/env bash
# Translate an existing PT figure to EN by passing it as input to
# gpt-image-2 /v1/images/edits — the model preserves the original
# layout/colors/icons and only swaps the inline text.
#
# Usage: gen-figure-translate.sh <slug> <name> "<text-replacement-prompt>"
#   slug: blog folder, e.g. ferritina-30-a-100-normalidade-que-esgota-mulheres
#   name: source name (figura-1) — output will be figura-1-en
#   prompt: explicit text-by-text replacement instructions
#
# Example:
#   gen-figure-translate.sh slug figura-1 "Replace 'O método AGIR' with 'The ACTS Method'. ..."
set -euo pipefail

SLUG="${1:?slug required}"
NAME="${2:?name required}"
PROMPT="${3:?prompt required}"
NAME="${NAME//\//-}"

if [ -z "${OPENAI_API_KEY:-}" ]; then
  OPENAI_API_KEY="$(grep -E '^OPENAI_API_KEY=' /home/user/plenya/apps/api/.env | cut -d= -f2-)"
fi

DIR="/home/user/plenya/apps/site/public/images/blog/$SLUG"
SRC="$DIR/$NAME.webp"
OUT="$DIR/${NAME}-en.webp"

[ -f "$SRC" ] || { echo "source not found: $SRC" >&2; exit 1; }

# gpt-image-2 /v1/images/edits requires PNG input. Convert WebP -> PNG.
TMP_PNG="$(mktemp --suffix=.png)"
trap 'rm -f "$TMP_PNG" "$TMP_JSON"' EXIT
convert "$SRC" "$TMP_PNG" 2>/dev/null || { echo "ImageMagick convert failed" >&2; exit 1; }

INSTRUCTION="Translate the inline text of this infographic from Brazilian Portuguese to US English. PRESERVE EXACTLY: the overall layout, every panel position, every color, every icon, every chart shape, every numeric value, the typography style, the cream background, and the proportions. CHANGE ONLY the textual labels — every Portuguese word, phrase, title, axis label, legend, annotation, source citation. Keep the visual identity 1:1; this is a localization, not a redesign. ${PROMPT}"

TMP_JSON="$(mktemp)"

HTTP_CODE=$(curl -s -o "$TMP_JSON" -w "%{http_code}" \
  https://api.openai.com/v1/images/edits \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -F model="gpt-image-2" \
  -F image="@${TMP_PNG}" \
  -F "prompt=${INSTRUCTION}" \
  -F size="1024x1024" \
  -F quality="high" \
  -F output_format="webp")

if [ "$HTTP_CODE" != "200" ]; then
  echo "OpenAI API error ($HTTP_CODE):" >&2
  head -c 1500 "$TMP_JSON" >&2
  exit 1
fi

jq -r '.data[0].b64_json' "$TMP_JSON" | base64 -d > "$OUT"
[ -s "$OUT" ] || { echo "empty output file" >&2; exit 1; }

if command -v convert >/dev/null 2>&1; then
  TMP_OPT="${OUT}.opt.webp"
  convert "$OUT" -strip -quality 85 -define webp:method=6 "$TMP_OPT" 2>/dev/null && mv "$TMP_OPT" "$OUT" || rm -f "$TMP_OPT"
fi

echo "$OUT"
