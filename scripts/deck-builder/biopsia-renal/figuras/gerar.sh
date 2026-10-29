#!/usr/bin/env bash
# Ilustrações do card do paciente (biópsia renal) com o modelo de imagem da fonte única.
#   ./gerar.sh            -> todas
#   ./gerar.sh conversa   -> só as indicadas
# Só ilustração: texto e prazos de remédio ficam no HTML (imagem de IA erra número).
set -euo pipefail
HERE="$(cd "$(dirname "$0")" && pwd)"
source /home/user/plenya/scripts/lib/openai-image.sh
RAW="$HERE/raw"; mkdir -p "$RAW"
STYLE="$(cat "$HERE/style.txt")"
K="$(grep -E '^OPENAI_API_KEY=' /home/user/plenya/apps/api/.env | cut -d= -f2- | tr -d '"'"'"'\r')"
WANT=("$@")
gen() {
  local id="$1" prompt="$2" tmp; tmp="$(mktemp)"
  local code; code=$(curl -s -o "$tmp" -w "%{http_code}" https://api.openai.com/v1/images/generations \
    -H "Authorization: Bearer $K" -H "Content-Type: application/json" \
    -d "$(jq -n --arg m "$OPENAI_IMAGE_MODEL" --arg p "${prompt}${STYLE}" \
          '{model:$m,prompt:$p,size:"1024x1024",quality:"high",n:1}')")
  [ "$code" = 200 ] || { echo "ERRO $id ($code)"; head -c 300 "$tmp"; rm -f "$tmp"; return 1; }
  jq -r '.data[0].b64_json' "$tmp" | base64 -d > "$RAW/$id.png"; rm -f "$tmp"; echo "ok $id"
}
while IFS=$'\t' read -r id prompt; do
  if [ ${#WANT[@]} -gt 0 ]; then printf '%s\n' "${WANT[@]}" | grep -qx "$id" || continue; fi
  gen "$id" "$prompt" &
done < "$HERE/specs.tsv"
wait
python3 "$HERE/preparar.py"
