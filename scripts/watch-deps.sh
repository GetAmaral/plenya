#!/bin/bash
# Script para reinstalar dependências automaticamente quando package.json muda

set -e

echo "🔍 Monitorando mudanças em package.json..."

# Última modificação conhecida
LAST_MODIFIED=""

while true; do
  # Verificar se algum package.json mudou
  CURRENT=$(find /app -name "package.json" -type f -exec stat -c %Y {} \; 2>/dev/null | md5sum)

  if [ "$LAST_MODIFIED" != "$CURRENT" ] && [ -n "$LAST_MODIFIED" ]; then
    echo ""
    echo "📦 Detectada mudança em package.json!"
    echo "🔄 Reinstalando dependências..."

    cd /app
    pnpm install --frozen-lockfile || pnpm install

    echo "✅ Dependências atualizadas!"
    echo ""
  fi

  LAST_MODIFIED="$CURRENT"
  sleep 5
done
