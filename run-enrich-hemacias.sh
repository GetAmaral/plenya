#!/bin/bash
###############################################################################
# Enriquecimento: Hemácias - Mulheres
# Item ID: 501fd84a-a440-4c13-9b11-35e2f69017d1
###############################################################################

set -e

echo ""
echo "################################################################################"
echo "# ENRIQUECIMENTO: Hemácias - Mulheres"
echo "# Item ID: 501fd84a-a440-4c13-9b11-35e2f69017d1"
echo "################################################################################"
echo ""

# Verificar se ANTHROPIC_API_KEY está definida
if [ -z "$ANTHROPIC_API_KEY" ]; then
  echo "⚠️  ANTHROPIC_API_KEY não está definida no ambiente."
  echo ""
  echo "Por favor, execute um dos comandos abaixo:"
  echo ""
  echo "  export ANTHROPIC_API_KEY='sk-ant-...'"
  echo "  ou"
  echo "  ANTHROPIC_API_KEY='sk-ant-...' ./run-enrich-hemacias.sh"
  echo ""
  exit 1
fi

echo "✓ API key configurada (${#ANTHROPIC_API_KEY} caracteres)"
echo ""

# Verificar se o container web está rodando
if ! docker compose ps web | grep -q "Up"; then
  echo "⚠️  Container 'web' não está rodando"
  echo "Iniciando containers..."
  docker compose up -d
  echo "✓ Containers iniciados"
  echo ""
  sleep 2
fi

# Verificar se as dependências estão instaladas
echo "Verificando dependências..."
if ! docker compose exec web test -f /app/apps/web/node_modules/@anthropic-ai/sdk/package.json 2>/dev/null; then
  echo "Instalando @anthropic-ai/sdk e pg..."
  docker compose exec web pnpm add @anthropic-ai/sdk pg --filter web
  echo "✓ Dependências instaladas"
  echo ""
fi

# Executar o script
echo "Iniciando enriquecimento..."
echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

docker compose exec \
  -e ANTHROPIC_API_KEY="$ANTHROPIC_API_KEY" \
  -e NODE_PATH=/app/apps/web/node_modules \
  web node /app/apps/web/enrich.mjs

EXIT_CODE=$?

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

if [ $EXIT_CODE -eq 0 ]; then
  echo "✓ Enriquecimento concluído com sucesso!"
  echo ""
  echo "Próximos passos:"
  echo "  1. Verificar dados no banco:"
  echo "     docker compose exec db psql -U plenya_user -d plenya_db -c \"SELECT name, LENGTH(clinical_relevance) FROM score_items WHERE id='501fd84a-a440-4c13-9b11-35e2f69017d1';\""
  echo ""
  echo "  2. Ver artigos vinculados:"
  echo "     docker compose exec db psql -U plenya_user -d plenya_db -c \"SELECT COUNT(*) FROM score_item_articles WHERE score_item_id='501fd84a-a440-4c13-9b11-35e2f69017d1';\""
  echo ""
else
  echo "❌ Erro durante a execução (código: $EXIT_CODE)"
  echo ""
  echo "Verifique:"
  echo "  - API key é válida"
  echo "  - Container web está rodando (docker compose ps)"
  echo "  - Database está acessível"
  echo ""
fi

exit $EXIT_CODE
