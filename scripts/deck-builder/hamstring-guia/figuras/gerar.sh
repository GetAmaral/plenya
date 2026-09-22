#!/usr/bin/env bash
# Regera as figuras do guia com gpt-image-2.5-sunburst e prepara os JPEG do documento.
#   ./gerar.sh            -> todas
#   ./gerar.sh e1 d1      -> só as indicadas
# Cada figura DEVE ser conferida contra os cues clínicos do specs.tsv antes de entrar no PDF.
# Ver memória: ai_imagem_nao_serve_para_figura_de_exercicio (o que funciona e o que não).
set -euo pipefail
HERE="$(cd "$(dirname "$0")" && pwd)"
RAW="$HERE/raw"; mkdir -p "$RAW"
# esquema anatômico usa um bloco de estilo próprio (ossos, sem figura humana)
STYLE_CORPO="$(cat "$HERE/style.txt")"
STYLE_ANAT="$(cat "$HERE/style-anat.txt")"
K="$(grep -E '^OPENAI_API_KEY=' /home/user/plenya/apps/api/.env | cut -d= -f2- | tr -d '"'"'"'\r')"
WANT=("$@")
gen() {
  local id="$1" prompt="$2" tmp; tmp="$(mktemp)"
  local style="$STYLE_CORPO"
  if [ "$id" = anat ]; then style="$STYLE_ANAT"; fi
  local code; code=$(curl -s -o "$tmp" -w "%{http_code}" https://api.openai.com/v1/images/generations \
    -H "Authorization: Bearer $K" -H "Content-Type: application/json" \
    -d "$(jq -n --arg m gpt-image-2.5-sunburst --arg p "${prompt}${style}" \
          '{model:$m,prompt:$p,size:"1536x1024",quality:"high",n:1}')")
  [ "$code" = 200 ] || { echo "ERRO $id ($code)"; head -c 300 "$tmp"; rm -f "$tmp"; return 1; }
  jq -r '.data[0].b64_json' "$tmp" | base64 -d > "$RAW/$id.png"; rm -f "$tmp"; echo "ok $id"
}
while IFS=$'\t' read -r id prompt; do
  if [ ${#WANT[@]} -gt 0 ]; then printf '%s\n' "${WANT[@]}" | grep -qx "$id" || continue; fi
  gen "$id" "$prompt" &
done < "$HERE/specs.tsv"
wait
python3 "$HERE/preparar.py"
