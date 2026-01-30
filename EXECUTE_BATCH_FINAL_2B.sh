#!/bin/bash
# =============================================================================
# SCRIPT DE EXECUÇÃO: BATCH FINAL 2 - PARTE B (45 ITEMS)
# =============================================================================

set -e  # Parar em caso de erro

echo "=========================================="
echo "BATCH FINAL 2 - PARTE B: 45 ITEMS"
echo "Enrichment MFI de Exames Laboratoriais"
echo "=========================================="
echo ""

# Verificar se Docker está rodando
if ! docker compose ps | grep -q "db.*running"; then
    echo "❌ ERRO: Container do PostgreSQL não está rodando!"
    echo "Execute: docker compose up -d"
    exit 1
fi

echo "✅ Container PostgreSQL está rodando"
echo ""

# Arquivo 1: Items 1-18 (detalhados)
echo "📋 Executando Parte 1 (Items 1-18 - detalhados)..."
docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/enrichment_data/batch_final_2_exames_B.sql 2>&1 | grep -E "(UPDATE|ERROR)" || echo "✅ Parte 1 concluída"
echo ""

# Arquivo 2: Items 19-25 (complementares)
echo "📋 Executando Parte 2 (Items 19-25 - complementares)..."
docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/enrichment_data/batch_final_2_exames_B_part2.sql 2>&1 | grep -E "(UPDATE|ERROR)" || echo "✅ Parte 2 concluída"
echo ""

# Arquivo 3: Items 26-45 (otimizados)
echo "📋 Executando Parte 3 (Items 26-45 - otimizados)..."
docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/enrichment_data/batch_final_2_exames_B_COMPLETE.sql 2>&1 | grep -E "(UPDATE|ERROR)" || echo "✅ Parte 3 concluída"
echo ""

# Verificação final
echo "🔍 Verificando enrichment..."
ENRICHED_COUNT=$(docker compose exec -T db psql -U plenya_user -d plenya_db -t -c "
SELECT COUNT(*)
FROM score_items
WHERE id IN (
    'bf77b326-caa5-46fd-b607-70a089918780',
    '1aa25d4b-a972-40db-a288-9cbe506de99e',
    '814d923f-cdfa-4388-9ba1-42b23dcd8d6d',
    '09577ef1-c3ad-461b-b2ad-59fab2c193d5',
    'ebcc36fd-d285-4754-adf7-50c7b130b286',
    '1fcd3bbc-920e-4d3b-bfe3-0dd0e376f346',
    'fe938255-6b7a-4fbd-ac8f-8f3ba0c2d291',
    'c21ccec2-66b2-49e3-911a-8d0944eda087',
    '66256f14-1775-4501-b85e-a4ffd71ccc7a',
    'bc896c78-c530-4dd6-ad03-6da22ffd7fea',
    '9c5ef523-046b-4647-abd4-719937d54eb6',
    '447746dd-949b-449b-bd0e-c2f226330a10',
    '22952700-63e3-422d-b4b2-179c0785a20e',
    'e622f011-e92d-459e-9f4a-af435f26ea74',
    '34af6e5c-3847-46d8-874e-a7364c014877'
) AND clinical_context IS NOT NULL;
" | xargs)

echo ""
echo "=========================================="
echo "✅ ENRICHMENT CONCLUÍDO"
echo "=========================================="
echo "📊 Items enriquecidos verificados: $ENRICHED_COUNT"
echo "🎯 Meta: 45 items"
echo ""
echo "Para visualizar um exemplo:"
echo "docker compose exec db psql -U plenya_user -d plenya_db -c \"SELECT name, LEFT(clinical_context, 100) as context FROM score_items WHERE id = 'bf77b326-caa5-46fd-b607-70a089918780';\""
echo ""
