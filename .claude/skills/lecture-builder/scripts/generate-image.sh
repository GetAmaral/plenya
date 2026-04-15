#!/bin/bash
# generate-image.sh — Gera imagem médica via DALL-E 3
# Uso: ./generate-image.sh "prompt descritivo" "nome-arquivo.png" "/caminho/output_dir"

PROMPT="$1"
FILENAME="$2"
OUTDIR="$3"

if [ -z "$PROMPT" ] || [ -z "$FILENAME" ] || [ -z "$OUTDIR" ]; then
  echo "Uso: $0 'prompt' 'nome-arquivo.png' '/caminho/output'" >&2
  echo "Exemplo: $0 'Fisiopatologia da resistência à insulina' 'section-3.png' '~/lecture-output/sm'" >&2
  exit 1
fi

# Carrega OPENAI_API_KEY do .env do projeto se não estiver no ambiente
PLENYA_ROOT="/home/user/plenya"
if [ -z "$OPENAI_API_KEY" ] && [ -f "$PLENYA_ROOT/.env" ]; then
  OPENAI_API_KEY=$(grep -m1 '^OPENAI_API_KEY=' "$PLENYA_ROOT/.env" | cut -d= -f2-)
fi
if [ -z "$OPENAI_API_KEY" ] && [ -f "$PLENYA_ROOT/apps/api/.env" ]; then
  OPENAI_API_KEY=$(grep -m1 '^OPENAI_API_KEY=' "$PLENYA_ROOT/apps/api/.env" | cut -d= -f2-)
fi

if [ -z "$OPENAI_API_KEY" ]; then
  echo "⚠️  OPENAI_API_KEY não encontrada (ambiente nem .env)." >&2
  echo "   Placeholder mantido: <!-- IMAGE:${FILENAME} -->" >&2
  exit 1
fi

# Cria diretório de imagens
mkdir -p "${OUTDIR}/images"

# Prompt enriquecido para contexto médico-educacional
FULL_PROMPT="Medical educational diagram: ${PROMPT}. Professional clinical illustration style, clean white background, suitable for medical lecture slides. High detail anatomical or physiological accuracy. No text overlaid on image. Color-coded components for clarity."

echo "🎨 Gerando imagem: ${FILENAME}..." >&2
echo "   Prompt: ${PROMPT}" >&2

# Chamada DALL-E 3
RESPONSE=$(curl -s --max-time 90 \
  https://api.openai.com/v1/images/generations \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -H "Content-Type: application/json" \
  -d "{
    \"model\": \"dall-e-3\",
    \"prompt\": $(echo "$FULL_PROMPT" | jq -Rs .),
    \"size\": \"1792x1024\",
    \"quality\": \"standard\",
    \"n\": 1
  }" 2>/dev/null)

# Extrai URL da imagem
IMAGE_URL=$(echo "$RESPONSE" | jq -r '.data[0].url' 2>/dev/null)

if [ -z "$IMAGE_URL" ] || [ "$IMAGE_URL" = "null" ]; then
  ERROR=$(echo "$RESPONSE" | jq -r '.error.message' 2>/dev/null || echo "Erro desconhecido")
  echo "❌ Erro ao gerar imagem: $ERROR" >&2
  echo "   Resposta completa: $RESPONSE" >&2
  exit 1
fi

# Download da imagem
echo "📥 Baixando imagem..." >&2
curl -sL --max-time 60 "$IMAGE_URL" -o "${OUTDIR}/images/${FILENAME}"

if [ $? -eq 0 ] && [ -s "${OUTDIR}/images/${FILENAME}" ]; then
  FILE_SIZE=$(du -sh "${OUTDIR}/images/${FILENAME}" | cut -f1)
  echo "✅ ${OUTDIR}/images/${FILENAME} (${FILE_SIZE})"
else
  echo "❌ Falha no download da imagem" >&2
  exit 1
fi
