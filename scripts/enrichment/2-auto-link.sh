#!/bin/bash
# ==============================================================================
# Script 2: Auto-Link Articles ↔ ScoreItems
# ==============================================================================
#
# O QUE FAZ:
#   - Cria links automáticos entre Articles e ScoreItems via similaridade vetorial
#   - Usa threshold adaptativo (0.2 para cold-start, 0.65-0.80 para warm)
#   - Batch processing com checkpoints (50 items por batch)
#   - Inclui link de orphans (ScoreItems sem nenhum link)
#
# QUANDO USAR:
#   - Após adicionar novos Articles ao sistema
#   - Após regenerar embeddings
#   - Para criar links iniciais
#   - Para adicionar links a ScoreItems órfãos
#
# DURAÇÃO: ~5-10 minutos para 878 ScoreItems
# RESULTADO ESPERADO: ~10,000-12,000 links criados
#
# ==============================================================================

set -e

echo "🔗 SCRIPT 2: Auto-Link Articles ↔ ScoreItems"
echo "=============================================="
echo ""
echo "Usage: $0 [uuid1 uuid2 ...]"
echo "  Sem args: processa TODOS os score items"
echo "  Com args: processa apenas os IDs especificados"
echo ""

# Verificar Docker
if ! docker compose ps | grep -q "api.*Up"; then
    echo "❌ Erro: Docker não está rodando!"
    exit 1
fi

# Estatísticas ANTES
echo "📊 Estado ANTES do auto-link:"
docker compose exec -T db psql -U plenya_user -d plenya_db << 'EOSQL'
SELECT
    COUNT(*) as links_atuais,
    COUNT(DISTINCT score_item_id) as scoreitems_com_links,
    ROUND(100.0 * COUNT(DISTINCT score_item_id) /
        (SELECT COUNT(*) FROM score_items WHERE deleted_at IS NULL), 1) as percent_cobertura
FROM article_score_items
WHERE auto_linked = true;
EOSQL

echo ""
read -p "⚠️  Deseja continuar com auto-link? Isso pode adicionar novos links (s/N): " -n 1 -r
echo ""
if [[ ! $REPLY =~ ^[Ss]$ ]]; then
    echo "❌ Cancelado pelo usuário"
    exit 0
fi

echo ""
echo "🚀 Executando auto-link-all (com batch processing)..."
echo ""

# Executar auto-link (passa IDs se fornecidos)
docker compose exec api go run /app/cmd/auto-link-all/main.go $@ 2>&1 | \
    grep -E "(Auto-linking|específicos|Batch|completed|Total|✅)"

echo ""
echo "🔗 Executando link-orphans (ScoreItems sem links)..."
echo ""

# Link orphans
docker compose exec api go run /app/cmd/link-orphans/main.go 2>&1 | \
    grep -E "(Linking|Completed|linkados|✅|⚠️)" | tail -10

echo ""
echo "✅ Auto-link concluído!"
echo ""

# Estatísticas DEPOIS
echo "📊 Estado DEPOIS do auto-link:"
docker compose exec -T db psql -U plenya_user -d plenya_db << 'EOSQL'
SELECT
    '=== LINKS ===' as categoria,
    COUNT(*)::text as total,
    COUNT(*) FILTER (WHERE confidence_score >= 0.5)::text || ' bons (≥0.5)' as qualidade
FROM article_score_items
WHERE auto_linked = true

UNION ALL

SELECT
    '=== COBERTURA ===' as categoria,
    COUNT(DISTINCT score_item_id)::text || '/' ||
    (SELECT COUNT(*)::text FROM score_items WHERE deleted_at IS NULL) as total,
    ROUND(100.0 * COUNT(DISTINCT score_item_id) /
        (SELECT COUNT(*) FROM score_items WHERE deleted_at IS NULL), 1)::text || '%' as qualidade
FROM article_score_items
WHERE auto_linked = true;
EOSQL

echo ""
echo "💡 Próximo passo: Execute ./3-prepare-with-prompts.sh"
echo ""
