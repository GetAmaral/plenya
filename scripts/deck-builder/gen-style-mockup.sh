#!/usr/bin/env bash
# Generate a style-reference mockup slide image via OpenAI gpt-image-2.
# Usage: gen-style-mockup.sh <name> "<prompt>"
# Output: /home/user/plenya/docs/decks/_assets/style-refs/<name>.png
set -euo pipefail

NAME="${1:?name required}"
PROMPT="${2:?prompt required}"

if [ -z "${OPENAI_API_KEY:-}" ]; then
  OPENAI_API_KEY="$(grep -E '^OPENAI_API_KEY=' /home/user/plenya/apps/api/.env | cut -d= -f2-)"
fi

OUT_DIR="/home/user/plenya/docs/decks/_assets/style-refs"
mkdir -p "$OUT_DIR"
OUT="$OUT_DIR/${NAME}.png"

TMP_JSON="$(mktemp)"
trap 'rm -f "$TMP_JSON"' EXIT

PAYLOAD=$(jq -n \
  --arg model "gpt-image-2" \
  --arg prompt "$PROMPT" \
  --arg size "1536x1024" \
  --arg quality "high" \
  '{model:$model, prompt:$prompt, size:$size, quality:$quality, n:1, output_format:"png"}')

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

jq -r '.data[0].b64_json' "$TMP_JSON" | base64 -d > "$OUT"

if [ ! -s "$OUT" ]; then
  echo "empty output file" >&2
  exit 1
fi

echo "$OUT"
