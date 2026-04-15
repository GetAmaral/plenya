#!/bin/bash
# export-pdf.sh — Converte slides Marp para PDF
# Uso: ./export-pdf.sh "/caminho/slides.md" "/caminho/slides.pdf"

SLIDES="$1"
OUTPUT="$2"

if [ -z "$SLIDES" ] || [ -z "$OUTPUT" ]; then
  echo "Uso: $0 '/caminho/slides.md' '/caminho/slides.pdf'" >&2
  exit 1
fi

if [ ! -f "$SLIDES" ]; then
  echo "❌ Arquivo não encontrado: $SLIDES" >&2
  exit 1
fi

# Diretório deste script → diretório pai = skill root
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
SKILL_DIR="$(dirname "$SCRIPT_DIR")"
THEME_FILE="${SKILL_DIR}/theme-medical.css"

echo "📑 Exportando PDF: $OUTPUT" >&2

# Verifica se Marp CLI está instalado
if ! command -v marp &>/dev/null; then
  echo "⚠️  Marp CLI não encontrado. Instalando..." >&2
  npm install -g @marp-team/marp-cli --silent 2>/dev/null \
    || npx @marp-team/marp-cli --version >/dev/null 2>&1

  # Tenta com npx se instalação global falhou
  if ! command -v marp &>/dev/null; then
    echo "   Usando npx marp..." >&2
    MARP_CMD="npx --yes @marp-team/marp-cli"
  fi
fi

MARP_CMD="${MARP_CMD:-marp}"

# Verifica se o tema existe
THEME_ARGS=""
if [ -f "$THEME_FILE" ]; then
  THEME_ARGS="--theme \"$THEME_FILE\""
else
  echo "⚠️  Tema não encontrado: $THEME_FILE. Usando tema default." >&2
fi

# Executa exportação
eval "$MARP_CMD \"$SLIDES\" \
  --pdf \
  $THEME_ARGS \
  --allow-local-files \
  --html \
  -o \"$OUTPUT\"" 2>&1

EXIT_CODE=$?

if [ $EXIT_CODE -eq 0 ] && [ -f "$OUTPUT" ]; then
  FILE_SIZE=$(du -sh "$OUTPUT" | cut -f1)
  echo "✅ PDF gerado: $OUTPUT (${FILE_SIZE})"
else
  echo "❌ Falha na exportação do PDF (código $EXIT_CODE)" >&2
  echo "" >&2
  echo "   Dica: Verifique se o Node.js está instalado:" >&2
  echo "   node --version && npm --version" >&2
  exit 1
fi
