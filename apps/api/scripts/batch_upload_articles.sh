#!/bin/bash

# Script para upload em lote de artigos via API HTTP
# Com DEV_BYPASS_AUTH ativado, não precisa de autenticação
# Uso: ./scripts/batch_upload_articles.sh

API_URL="http://localhost:3001/api/v1"
ORIGINALS_DIR="./uploads/originals"

# Cores para output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "📚 Batch Upload de Artigos via API"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "🔓 DEV_BYPASS_AUTH ativo - sem autenticação necessária"
echo ""

# Listar PDFs
if [ ! -d "$ORIGINALS_DIR" ]; then
  echo -e "${RED}❌ Diretório $ORIGINALS_DIR não encontrado${NC}"
  exit 1
fi

PDF_FILES=("$ORIGINALS_DIR"/*.pdf)
TOTAL_FILES=${#PDF_FILES[@]}

if [ $TOTAL_FILES -eq 0 ]; then
  echo -e "${RED}❌ Nenhum arquivo PDF encontrado em $ORIGINALS_DIR${NC}"
  exit 1
fi

echo -e "${BLUE}📦 Encontrados $TOTAL_FILES arquivos PDF${NC}"
echo ""

# Contadores
SUCCESS=0
FAILED=0
SKIPPED=0

# Upload de cada PDF
for i in "${!PDF_FILES[@]}"; do
  FILE_PATH="${PDF_FILES[$i]}"
  FILE_NAME=$(basename "$FILE_PATH")
  CURRENT=$((i + 1))

  echo -e "${BLUE}[$CURRENT/$TOTAL_FILES]${NC} Processando: $FILE_NAME"

  # Upload via API (sem token - DEV_BYPASS_AUTH)
  # Usar -o para salvar response em arquivo temporário
  TEMP_RESPONSE=$(mktemp)
  HTTP_CODE=$(curl -s -o "$TEMP_RESPONSE" -w "%{http_code}" -X POST "$API_URL/articles/upload" \
    -F "file=@$FILE_PATH")

  if [ "$HTTP_CODE" -eq 201 ]; then
    # Sucesso - extrair apenas ID (primeiros chars)
    ARTICLE_ID=$(grep -o '"id":"[^"]*' "$TEMP_RESPONSE" | head -1 | cut -d'"' -f4 | head -c 8)

    echo -e "   ${GREEN}✅ Upload realizado com sucesso${NC}"
    echo "      ID: $ARTICLE_ID..."
    ((SUCCESS++))
  elif [ "$HTTP_CODE" -eq 400 ]; then
    # Verificar se é duplicação
    if grep -q "já foi importado" "$TEMP_RESPONSE"; then
      echo -e "   ${YELLOW}⏭️  Ignorado (já importado anteriormente)${NC}"
      ((SKIPPED++))
    else
      ERROR_MSG=$(grep -o '"error":"[^"]*' "$TEMP_RESPONSE" | cut -d'"' -f4)
      echo -e "   ${RED}❌ Falha: HTTP $HTTP_CODE - $ERROR_MSG${NC}"
      ((FAILED++))
    fi
  else
    # Outro erro
    ERROR_MSG=$(grep -o '"error":"[^"]*' "$TEMP_RESPONSE" | cut -d'"' -f4)
    echo -e "   ${RED}❌ Falha: HTTP $HTTP_CODE${NC}"
    if [ -n "$ERROR_MSG" ]; then
      echo "      Erro: $ERROR_MSG"
    fi
    ((FAILED++))
  fi

  # Limpar arquivo temporário
  rm -f "$TEMP_RESPONSE"

  echo ""

  # Delay para não sobrecarregar
  sleep 0.5
done

# Resumo final
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "📊 RESUMO DO UPLOAD"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo -e "   ${GREEN}✅ Sucesso:${NC}  $SUCCESS"
echo -e "   ${YELLOW}⏭️  Ignorados:${NC} $SKIPPED (já importados)"
echo -e "   ${RED}❌ Falhas:${NC}   $FAILED"
echo -e "   ${BLUE}📦 Total:${NC}    $TOTAL_FILES"
echo ""

if [ $SUCCESS -gt 0 ]; then
  echo "🤖 Embeddings estão sendo processados pelo worker em background"
  echo "   Acompanhe o progresso com: docker compose logs -f api"
fi

echo ""
echo "✅ Batch upload concluído!"
