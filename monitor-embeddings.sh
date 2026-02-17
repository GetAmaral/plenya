#!/bin/bash
# Monitor de progresso da re-indexação de embeddings

echo "📊 Monitorando re-indexação de embeddings com Voyage AI"
echo "========================================================"
echo ""

while true; do
  clear
  echo "📊 Re-indexação de Embeddings — Voyage AI"
  echo "========================================================"
  date
  echo ""

  # Status da fila
  docker compose exec db psql -U plenya_user -d plenya_db -t -c "
    SELECT
      status || ': ' || COUNT(*)
    FROM embedding_queue
    GROUP BY status
    ORDER BY
      CASE status
        WHEN 'completed' THEN 1
        WHEN 'processing' THEN 2
        WHEN 'pending' THEN 3
        WHEN 'failed' THEN 4
      END;
  " | sed 's/^[ \t]*/  /'

  echo ""

  # Progresso %
  docker compose exec db psql -U plenya_user -d plenya_db -t -c "
    SELECT
      '  Progresso: ' ||
      ROUND((COUNT(*) FILTER (WHERE status = 'completed')::numeric / COUNT(*)::numeric * 100), 1) ||
      '% (' || COUNT(*) FILTER (WHERE status = 'completed') || '/' || COUNT(*) || ')'
    FROM embedding_queue;
  "

  echo ""

  # Últimos processados
  echo "  Últimos 3 processados:"
  docker compose exec db psql -U plenya_user -d plenya_db -t -c "
    SELECT
      '    ' || TO_CHAR(processed_at, 'HH24:MI:SS') || ' - ' ||
      CASE entity_type
        WHEN 'article' THEN 'Article'
        WHEN 'score_item' THEN 'ScoreItem'
      END
    FROM embedding_queue
    WHERE status = 'completed' AND processed_at IS NOT NULL
    ORDER BY processed_at DESC
    LIMIT 3;
  "

  echo ""
  echo "  Pressione Ctrl+C para sair"
  echo ""

  sleep 10
done
