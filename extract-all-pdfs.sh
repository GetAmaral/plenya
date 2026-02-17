#!/bin/bash
# Extrai conteúdo de todos os PDFs de artigos sem fullContent

echo "📄 Extraindo todos os PDFs com conteúdo faltante"
echo "==============================================="
echo ""

count=0
while true; do
  # Verificar quantos restam
  remaining=$(docker compose exec db psql -U plenya_user -d plenya_db -t -c "
    SELECT COUNT(*)
    FROM articles
    WHERE internal_link IS NOT NULL
      AND internal_link LIKE '/uploads/articles/%'
      AND (full_content IS NULL OR full_content = '')
      AND deleted_at IS NULL;
  " | tr -d ' ')

  if [ "$remaining" = "0" ]; then
    echo ""
    echo "✅ Todos os PDFs extraídos!"
    echo "   Total processados: $count"
    break
  fi

  count=$((count + 1))
  echo "[$count] Restam: $remaining artigos"

  # Processar um
  docker compose exec api go run cmd/extract-one-pdf/main.go 2>&1 | grep -v "^$"

  if [ $? -ne 0 ]; then
    echo "❌ Erro ao processar. Continuando..."
  fi

  echo ""
  sleep 1
done
