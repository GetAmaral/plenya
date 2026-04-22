#!/bin/bash
# Gera 3 ilustrações instrucionais via OpenAI gpt-image-1
# Salva em apps/site/public/escore-light/instructions/
# Custom size 1024x1024, formato PNG.
set -euo pipefail

cd "$(dirname "$0")/../../.."
set -a; source .env; set +a

OUT=/home/user/plenya/apps/site/public/escore-light/instructions
mkdir -p "$OUT"

generate() {
  local name="$1"
  local prompt="$2"
  echo "→ Gerando $name..."
  curl -sS https://api.openai.com/v1/images/generations \
    -H "Authorization: Bearer $OPENAI_API_KEY" \
    -H "Content-Type: application/json" \
    -d "{
      \"model\": \"gpt-image-1\",
      \"prompt\": $(printf '%s' "$prompt" | jq -Rs .),
      \"size\": \"1024x1024\",
      \"quality\": \"high\",
      \"n\": 1
    }" > "/tmp/${name}.json"

  # gpt-image-1 returns base64 in data[0].b64_json
  jq -r '.data[0].b64_json' "/tmp/${name}.json" | base64 -d > "$OUT/${name}.png"

  if [ ! -s "$OUT/${name}.png" ]; then
    echo "ERROR: empty image for $name. API response:"
    cat "/tmp/${name}.json"
    exit 1
  fi
  echo "  ✓ $OUT/${name}.png ($(stat -c%s "$OUT/${name}.png") bytes)"
}

PROMPT_PLANK="A clean, minimalist medical instructional illustration showing a person performing a forearm plank exercise correctly. Side view, single human figure in plank position with forearms on the ground, body in a straight horizontal line from head to heels. Subtle dotted line above showing the perfect alignment from head to feet. Style: thin elegant line drawing in dark teal color, on a soft cream/paper background. Professional health-book aesthetic, no shading, no labels or text, ample white space. Editorial quality."

PROMPT_WAIST="A clean, minimalist medical instructional illustration showing how to measure waist circumference. Front view of a slim human torso silhouette (gender neutral), with a measuring tape wrapped horizontally around the waist exactly at navel level. A small marker indicates the navel. Style: thin elegant line drawing in dark teal color, with the measuring tape accented in warm gold, on a soft cream/paper background. Professional health-book aesthetic, no shading, no labels or text, ample white space. Editorial quality."

PROMPT_NECK="A clean, minimalist medical instructional illustration showing how to measure neck circumference. Front view of a person's head, neck, and upper shoulders. A measuring tape is wrapped around the neck just below the Adam's apple at the narrowest point. Style: thin elegant line drawing in dark teal color, with the measuring tape accented in warm gold, on a soft cream/paper background. Professional health-book aesthetic, no shading, no labels or text, ample white space. Editorial quality."

generate "plank" "$PROMPT_PLANK"
generate "waist" "$PROMPT_WAIST"
generate "neck" "$PROMPT_NECK"

echo "✓ Done"
