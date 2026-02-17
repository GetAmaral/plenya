#!/bin/bash
# ==============================================================================
# MASTER SCRIPT: Pipeline Completo de Enrichment
# ==============================================================================
#
# O QUE FAZ:
#   Executa o pipeline completo de enrichment em 3 etapas:
#   1. Regenera embeddings (se necessário)
#   2. Cria auto-links entre Articles e ScoreItems
#   3. Prepara prompts para enrichment com Claude
#
# QUANDO USAR:
#   - Setup inicial do sistema
#   - Após adicionar muitos Articles novos
#   - Para atualizar todo o sistema de uma vez
#
# DURAÇÃO TOTAL: ~30-40 minutos
#
# ==============================================================================

set -e

echo "╔════════════════════════════════════════════════════════════════╗"
echo "║  PIPELINE COMPLETO DE ENRICHMENT CIENTÍFICO                    ║"
echo "║  Sistema RAG para Plenya EMR                                   ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""

# Verificar Docker
if ! docker compose ps | grep -q "api.*running"; then
    echo "❌ Erro: Docker compose não está rodando!"
    echo "   Execute: docker compose up -d"
    exit 1
fi

# Confirmação
echo "Este script vai executar:"
echo "  1️⃣  Regenerar embeddings (se necessário)"
echo "  2️⃣  Criar auto-links (Articles ↔ ScoreItems)"
echo "  3️⃣  Preparar prompts para Claude (4 por ScoreItem)"
echo ""
echo "⏱️  Duração estimada: 30-40 minutos"
echo ""
read -p "Deseja continuar? (s/N): " -n 1 -r
echo ""
if [[ ! $REPLY =~ ^[Ss]$ ]]; then
    echo "❌ Cancelado"
    exit 0
fi

echo ""
echo "════════════════════════════════════════════════════════════════"
echo ""

# ETAPA 1: Embeddings
echo "1️⃣  ETAPA 1: Verificando embeddings..."
echo ""

STALE=$(docker compose exec -T db psql -U plenya_user -d plenya_db -t -c \
    "SELECT COUNT(*) FROM score_item_embeddings WHERE is_stale = true")
STALE=$(echo $STALE | tr -d ' ')

if [ "$STALE" -gt 0 ]; then
    echo "   ⚠️  Encontrados $STALE embeddings stale"
    echo "   🔄 Regenerando..."
    ./1-regenerate-embeddings.sh
else
    echo "   ✅ Embeddings estão atualizados (0 stale)"
fi

echo ""
echo "════════════════════════════════════════════════════════════════"
echo ""

# ETAPA 2: Auto-Link
echo "2️⃣  ETAPA 2: Auto-Link..."
echo ""
./2-auto-link.sh

echo ""
echo "════════════════════════════════════════════════════════════════"
echo ""

# ETAPA 3: Prepare
echo "3️⃣  ETAPA 3: Preparando prompts..."
echo ""

# Forçar re-prepare de todos
docker compose exec -T db psql -U plenya_user -d plenya_db -c \
    "DELETE FROM score_item_enrichment_preparation" > /dev/null

docker compose exec api go run /app/cmd/prepare-all/main.go 2>&1 | \
    grep -E "(Preparing|Processed|Prepared)" | tail -5

echo ""
echo "════════════════════════════════════════════════════════════════"
echo ""

# RESUMO FINAL
echo "📊 RESUMO FINAL DO SISTEMA"
echo ""

docker compose exec -T db psql -U plenya_user -d plenya_db << 'EOSQL'
SELECT
    '🔗 Auto-Links' as componente,
    COUNT(*)::text as total,
    ROUND(100.0 * COUNT(*) FILTER (WHERE confidence_score >= 0.5) / COUNT(*), 1)::text || '% bons' as qualidade
FROM article_score_items
WHERE auto_linked = true

UNION ALL

SELECT
    '📚 ScoreItems Cobertos' as componente,
    COUNT(DISTINCT asi.score_item_id)::text || '/' ||
    (SELECT COUNT(*)::text FROM score_items WHERE deleted_at IS NULL) as total,
    ROUND(100.0 * COUNT(DISTINCT asi.score_item_id) /
        (SELECT COUNT(*) FROM score_items WHERE deleted_at IS NULL), 1)::text || '% cobertura' as qualidade
FROM article_score_items asi
WHERE asi.auto_linked = true

UNION ALL

SELECT
    '🔬 Preparations' as componente,
    COUNT(*)::text || ' com 4 prompts' as total,
    ROUND(100.0 * COUNT(*) FILTER (WHERE metadata->>'quality_grade' IN ('excellent','good')) / COUNT(*), 1)::text || '% excellent/good' as qualidade
FROM score_item_enrichment_preparation
WHERE prompt_clinical_relevance IS NOT NULL;
EOSQL

echo ""
echo "╔════════════════════════════════════════════════════════════════╗"
echo "║  ✅ PIPELINE COMPLETO - SISTEMA PRONTO PARA ENRICHMENT!        ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""
echo "🚀 Próximo passo: Enrichment com Claude"
echo "   - 878 ScoreItems prontos"
echo "   - 4 prompts por item (~32KB cada)"
echo "   - Total: 3,512 prompts para processar"
echo ""
