#!/usr/bin/env bash
# Imagem para a AULA de transplante renal (deck Plenya), com direção de arte
# casada com a fotografia real da equipe em docs/site/images/FotosHd:
# fundo navy/petrol, luz quente de recorte, sombra densa, pouca profundidade.
#
# Uso: gen-aula-image.sh <nome> "<prompt>"
# Saída: ~/lecture-output/cuidados-transplante-renal/images/<nome>.png
set -euo pipefail

# modelo de imagem: fonte única em scripts/lib/openai-image.sh
. "$(dirname "${BASH_SOURCE[0]}")/../lib/openai-image.sh"

NAME="${1:?nome required}"; PROMPT="${2:?prompt required}"; NAME="${NAME//\//-}"
QUALITY="${TX_IMG_QUALITY:-xhigh}"
SIZE="${TX_IMG_SIZE:-1536x1024}"

[ -n "${OPENAI_API_KEY:-}" ] || OPENAI_API_KEY="$(grep -E '^OPENAI_API_KEY=' /home/user/plenya/apps/api/.env | cut -d= -f2-)"

OUT_DIR="/home/user/lecture-output/cuidados-transplante-renal/images"
mkdir -p "$OUT_DIR"; OUT="$OUT_DIR/$NAME.png"

STYLE=" Cinematic editorial photograph for a nephrology lecture by a Brazilian medical practice. Art direction matched to a brand photo library shot against deep navy and petrol backgrounds with warm amber practical light: dense shadows, one strong warm key light, soft falloff, shallow depth of field like a 50mm lens at f/1.8, fine film grain, gentle highlight roll-off, no flat overhead fluorescent look. Restrained palette: deep petrol blue, teal shadow, warm gold and amber accents, cream and bone whites. Composed with generous clean negative space on one side so a title can sit over it. Absolutely NO text anywhere: no words, letters, numbers, labels, signage, readable screens, logos or watermarks; any monitor or display must be switched off or far out of focus. No recognizable faces: hands, backs or silhouettes only. Photorealistic, not illustrated, not 3D render."

PAYLOAD=$(jq -n --arg m "$OPENAI_IMAGE_MODEL" --arg p "${PROMPT}.${STYLE}" --arg s "$SIZE" --arg q "$QUALITY" \
  '{model:$m, prompt:$p, size:$s, quality:$q, n:1}')

TMP="$(mktemp)"; trap 'rm -f "$TMP"' EXIT
CODE=$(curl -s --max-time 300 -o "$TMP" -w "%{http_code}" https://api.openai.com/v1/images/generations \
  -H "Authorization: Bearer $OPENAI_API_KEY" -H "Content-Type: application/json" -d "$PAYLOAD")
[ "$CODE" = "200" ] || { echo "erro $CODE em $NAME:" >&2; head -c 500 "$TMP" >&2; exit 1; }
jq -r '.data[0].b64_json' "$TMP" | base64 -d > "$OUT"
[ -s "$OUT" ] || { echo "vazio: $NAME" >&2; exit 1; }
echo "$OUT ($(du -h "$OUT"|cut -f1), $OPENAI_IMAGE_MODEL/$QUALITY)"
