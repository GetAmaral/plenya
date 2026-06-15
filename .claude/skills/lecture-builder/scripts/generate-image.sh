#!/bin/bash
# generate-image.sh — Gera ilustração médica via OpenAI gpt-image-2
# Uso: ./generate-image.sh "prompt" "nome-arquivo.png" "/caminho/output_dir"

PROMPT="$1"
FILENAME="$2"
OUTDIR="$3"

if [ -z "$PROMPT" ] || [ -z "$FILENAME" ] || [ -z "$OUTDIR" ]; then
  echo "Uso: $0 'prompt' 'nome-arquivo.png' '/caminho/output_dir'" >&2
  exit 1
fi

PLENYA_ROOT="/home/user/plenya"
if [ -z "$OPENAI_API_KEY" ] && [ -f "$PLENYA_ROOT/.env" ]; then
  OPENAI_API_KEY=$(grep -m1 '^OPENAI_API_KEY=' "$PLENYA_ROOT/.env" | cut -d= -f2-)
fi
if [ -z "$OPENAI_API_KEY" ] && [ -f "$PLENYA_ROOT/apps/api/.env" ]; then
  OPENAI_API_KEY=$(grep -m1 '^OPENAI_API_KEY=' "$PLENYA_ROOT/apps/api/.env" | cut -d= -f2-)
fi

if [ -z "$OPENAI_API_KEY" ]; then
  echo "⚠️  OPENAI_API_KEY não encontrada." >&2
  exit 1
fi

mkdir -p "${OUTDIR}/images"

# Prompt enriquecido para contexto médico-educacional, on-brand Plenya.
# Sem texto na imagem (legendas vão no HTML), paleta clínica neutra.
FULL_PROMPT="Medical educational illustration: ${PROMPT}. Editorial style suitable for an academic medical lecture slide. Clean light background or restrained color palette (cream, deep teal, muted gold). Anatomical/physiological accuracy. No text overlay, no labels embedded in the image. Composition leaves negative space for caption placement. High detail, crisp lines."

echo "🎨 gpt-image-2: ${FILENAME}" >&2
echo "   Prompt: ${PROMPT}" >&2

RESPONSE=$(curl -s --max-time 180 \
  https://api.openai.com/v1/images/generations \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -H "Content-Type: application/json" \
  -d "{
    \"model\": \"gpt-image-2\",
    \"prompt\": $(echo "$FULL_PROMPT" | jq -Rs .),
    \"size\": \"1792x1024\",
    \"quality\": \"high\",
    \"n\": 1
  }" 2>/dev/null)

# gpt-image-2 retorna b64_json por padrão
B64=$(echo "$RESPONSE" | jq -r '.data[0].b64_json // empty' 2>/dev/null)
URL=$(echo "$RESPONSE" | jq -r '.data[0].url // empty' 2>/dev/null)

if [ -n "$B64" ]; then
  echo "$B64" | base64 -d > "${OUTDIR}/images/${FILENAME}"
elif [ -n "$URL" ]; then
  curl -sL --max-time 60 "$URL" -o "${OUTDIR}/images/${FILENAME}"
else
  ERROR=$(echo "$RESPONSE" | jq -r '.error.message // "Erro desconhecido"' 2>/dev/null)
  echo "❌ Erro: $ERROR" >&2
  echo "   Resposta: $RESPONSE" | head -c 600 >&2
  exit 1
fi

if [ -s "${OUTDIR}/images/${FILENAME}" ]; then
  SIZE=$(du -sh "${OUTDIR}/images/${FILENAME}" | cut -f1)
  echo "✅ ${OUTDIR}/images/${FILENAME} (${SIZE})"
else
  echo "❌ Arquivo vazio" >&2
  exit 1
fi
