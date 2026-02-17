#!/bin/bash
# ==============================================================================
# Script 1: Regenerar Embeddings
# ==============================================================================
#
# O QUE FAZ:
#   - Regenera embeddings para ScoreItems que estão marcados como stale
#   - Processa fila embedding_queue (criada por hooks ou manualmente)
#   - Rate limiting automático (200ms entre items)
#   - Retry automático (max 3x) para falhas
#
# QUANDO USAR:
#   - Após editar ScoreItems (Name, ClinicalRelevance, Gender, etc)
#   - Quando embedding_queue tem items pendentes
#   - Para regenerar embeddings com novo formato/modelo
#
# DURAÇÃO: ~15-20 minutos para 878 ScoreItems (~1 segundo por item)
# CUSTO: ~$0.50 (878 items × 800 chars × $0.13/1M tokens)
#
# ==============================================================================

set -e

echo "🔄 SCRIPT 1: Regenerar Embeddings"
echo "=================================="
echo ""

# Verificar se Docker está rodando
if ! docker compose ps | grep -q "api.*running"; then
    echo "❌ Erro: Docker compose não está rodando!"
    echo "   Execute: docker compose up -d"
    exit 1
fi

# Verificar quantos embeddings estão na fila
PENDING=$(docker compose exec -T db psql -U plenya_user -d plenya_db -t -c \
    "SELECT COUNT(*) FROM embedding_queue WHERE status = 'pending' AND entity_type = 'score_item'")
PENDING=$(echo $PENDING | tr -d ' ')

echo "📋 Fila de embeddings:"
echo "   Pending: $PENDING ScoreItems"
echo ""

if [ "$PENDING" -eq 0 ]; then
    echo "✅ Nenhum embedding pendente. Sistema está atualizado!"
    echo ""
    echo "💡 Para marcar TODOS como stale e regenerar:"
    echo "   docker compose exec db psql -U plenya_user -d plenya_db -c \"UPDATE score_item_embeddings SET is_stale = true\""
    exit 0
fi

echo "🚀 Iniciando regeneração de $PENDING embeddings..."
echo "   Tempo estimado: ~$((PENDING / 5)) segundos"
echo ""

# Executar regeneração
docker compose exec api go run /app/cmd/regenerate-embeddings/main.go 2>&1 | \
    grep -E "(Starting|Found|Completed|✅|❌|Progress)"

echo ""
echo "✅ Regeneração concluída!"
echo ""

# Estatísticas finais
docker compose exec -T db psql -U plenya_user -d plenya_db << 'EOSQL'
SELECT
    '📊 RESULTADO FINAL' as status,
    COUNT(*) FILTER (WHERE is_stale = false) as embeddings_validos,
    COUNT(*) FILTER (WHERE is_stale = true) as embeddings_stale,
    ROUND(100.0 * COUNT(*) FILTER (WHERE is_stale = false) / COUNT(*), 1) as percent_validos
FROM score_item_embeddings;
EOSQL

echo ""
echo "💾 Embeddings atualizados e prontos para auto-link!"
echo ""
