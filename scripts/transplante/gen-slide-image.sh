#!/usr/bin/env bash
# Gera uma imagem de slide da palestra "Cuidados com o Paciente Transplantado Renal".
#
# Uso: gen-slide-image.sh <nome> "<prompt>"
#   <nome> vira o arquivo: docs/transplante/images/<nome>.webp
#   ex.: gen-slide-image.sh 04-caso "mesa de cabeceira na madrugada..."
#
# Modelo: gpt-image-2.5-sunburst (qualidade). Para iterar rápido, exporte
# TX_IMG_MODEL=gpt-image-2.5-flare antes de chamar.
#
# O estilo é fixo de propósito: a aula é da Santa Casa / NefroClínica, não da
# Plenya, então NÃO usa a paleta petrol/gold da marca. Documentário clínico,
# luz natural, sem texto, sem rosto reconhecível.
set -euo pipefail

NAME="${1:?nome required (ex.: 04-caso)}"
PROMPT="${2:?prompt required}"
NAME="${NAME//\//-}"

MODEL="${TX_IMG_MODEL:-gpt-image-2.5-sunburst}"
SIZE="${TX_IMG_SIZE:-1536x1024}"

if [ -z "${OPENAI_API_KEY:-}" ]; then
  OPENAI_API_KEY="$(grep -E '^OPENAI_API_KEY=' /home/user/plenya/apps/api/.env | cut -d= -f2-)"
fi

OUT_DIR="/home/user/plenya/docs/transplante/images"
mkdir -p "$OUT_DIR"
OUT="$OUT_DIR/$NAME.webp"

STYLE_SUFFIX=" Documentary editorial photograph for a hospital nursing lecture. Cool natural daylight, muted palette of soft greens, warm greys and clinical whites, real textures, shallow depth of field, calm and unstaged. Generous negative space on one side so a title can sit over the image. ABSOLUTELY NO TEXT anywhere in the image: no words, no letters, no numbers, no labels, no signage, no readable screens, no logos, no watermarks. Any monitor or display in frame must be switched off or fully out of focus. No recognizable human faces; if people appear, show hands, backs or silhouettes only. Photorealistic, not illustrated, not 3D render."

FULL_PROMPT="${PROMPT}.${STYLE_SUFFIX}"

TMP_JSON="$(mktemp)"
trap 'rm -f "$TMP_JSON"' EXIT

PAYLOAD=$(jq -n \
  --arg model "$MODEL" \
  --arg prompt "$FULL_PROMPT" \
  --arg size "$SIZE" \
  '{model:$model, prompt:$prompt, size:$size, quality:"high", n:1, output_format:"webp"}')

HTTP_CODE=$(curl -s -o "$TMP_JSON" -w "%{http_code}" \
  https://api.openai.com/v1/images/generations \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -H "Content-Type: application/json" \
  -d "$PAYLOAD")

if [ "$HTTP_CODE" != "200" ]; then
  echo "OpenAI API error ($HTTP_CODE) em $NAME:" >&2
  head -c 600 "$TMP_JSON" >&2
  exit 1
fi

jq -r '.data[0].b64_json' "$TMP_JSON" | base64 -d > "$OUT"
[ -s "$OUT" ] || { echo "saida vazia em $NAME" >&2; exit 1; }
echo "$OUT ($(du -h "$OUT" | cut -f1), $MODEL)"
