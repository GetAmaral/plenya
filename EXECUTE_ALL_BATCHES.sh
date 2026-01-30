#!/bin/bash
# ============================================================================
# EXECUÇÃO CONSOLIDADA - TODOS OS 5 BATCHES (168 ITEMS)
# ============================================================================

set -e  # Exit on error

echo "================================================================================"
echo "EXECUTANDO TODOS OS BATCHES DA RODADA 3 (168 ITEMS)"
echo "================================================================================"
echo ""

# Verificar se Docker está rodando
echo "🔍 Verificando Docker..."
if ! docker compose ps db | grep -q "healthy\|running"; then
    echo "❌ ERRO: Container PostgreSQL não está rodando"
    echo "Execute: docker compose up -d"
    exit 1
fi
echo "✅ Docker está rodando"
echo ""

# Contador de items
TOTAL_ITEMS=0

# ============================================================================
# BATCH 1 - EXAMES PARTE A (45 ITEMS)
# ============================================================================
echo "📦 BATCH 1: Exames Parte A (45 items)"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

if [ -f "batch_final_1_exames_A.sql" ]; then
    echo "Executando: batch_final_1_exames_A.sql"
    docker compose exec -T db psql -U plenya_user -d plenya_db < batch_final_1_exames_A.sql > /tmp/batch1_output.txt 2>&1

    if grep -q "COMMIT" /tmp/batch1_output.txt; then
        BATCH1_COUNT=$(grep -c "UPDATE 1" /tmp/batch1_output.txt || echo "0")
        echo "✅ Batch 1 concluído: $BATCH1_COUNT items atualizados"
        TOTAL_ITEMS=$((TOTAL_ITEMS + BATCH1_COUNT))
    else
        echo "❌ Erro no Batch 1 - verificar /tmp/batch1_output.txt"
        cat /tmp/batch1_output.txt
        exit 1
    fi
else
    echo "⚠️  Arquivo batch_final_1_exames_A.sql não encontrado - pulando"
fi
echo ""

# ============================================================================
# BATCH 2B - EXAMES PARTE B (45 ITEMS - 3 PARTES)
# ============================================================================
echo "📦 BATCH 2B: Exames Parte B (45 items em 3 partes)"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

BATCH2_TOTAL=0

# Parte 1
if [ -f "scripts/enrichment_data/batch_final_2_exames_B.sql" ]; then
    echo "  Parte 1: Items detalhados (18 items)"
    docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/enrichment_data/batch_final_2_exames_B.sql > /tmp/batch2_part1.txt 2>&1
    PART1_COUNT=$(grep -c "UPDATE 1" /tmp/batch2_part1.txt || echo "0")
    echo "  ✅ Parte 1: $PART1_COUNT items"
    BATCH2_TOTAL=$((BATCH2_TOTAL + PART1_COUNT))
fi

# Parte 2
if [ -f "scripts/enrichment_data/batch_final_2_exames_B_part2.sql" ]; then
    echo "  Parte 2: Items complementares (7 items)"
    docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/enrichment_data/batch_final_2_exames_B_part2.sql > /tmp/batch2_part2.txt 2>&1
    PART2_COUNT=$(grep -c "UPDATE 1" /tmp/batch2_part2.txt || echo "0")
    echo "  ✅ Parte 2: $PART2_COUNT items"
    BATCH2_TOTAL=$((BATCH2_TOTAL + PART2_COUNT))
fi

# Parte 3
if [ -f "scripts/enrichment_data/batch_final_2_exames_B_COMPLETE.sql" ]; then
    echo "  Parte 3: Items otimizados (20 items)"
    docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/enrichment_data/batch_final_2_exames_B_COMPLETE.sql > /tmp/batch2_part3.txt 2>&1
    PART3_COUNT=$(grep -c "UPDATE 1" /tmp/batch2_part3.txt || echo "0")
    echo "  ✅ Parte 3: $PART3_COUNT items"
    BATCH2_TOTAL=$((BATCH2_TOTAL + PART3_COUNT))
fi

echo "✅ Batch 2B concluído: $BATCH2_TOTAL items atualizados"
TOTAL_ITEMS=$((TOTAL_ITEMS + BATCH2_TOTAL))
echo ""

# ============================================================================
# BATCH 3 - EXAMES PARTE C (35 ITEMS)
# ============================================================================
echo "📦 BATCH 3: Exames Parte C (35 items)"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

if [ -f "scripts/batch_final_3_ALL_35_items.sql" ]; then
    echo "Executando: batch_final_3_ALL_35_items.sql"
    docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/batch_final_3_ALL_35_items.sql > /tmp/batch3_output.txt 2>&1

    if grep -q "COMMIT" /tmp/batch3_output.txt; then
        BATCH3_COUNT=$(grep -c "UPDATE 1" /tmp/batch3_output.txt || echo "0")
        echo "✅ Batch 3 concluído: $BATCH3_COUNT items atualizados"
        TOTAL_ITEMS=$((TOTAL_ITEMS + BATCH3_COUNT))
    else
        echo "❌ Erro no Batch 3 - verificar /tmp/batch3_output.txt"
        cat /tmp/batch3_output.txt
        exit 1
    fi
else
    echo "⚠️  Arquivo batch_final_3_ALL_35_items.sql não encontrado - pulando"
fi
echo ""

# ============================================================================
# BATCH 4 - HISTÓRICO DE DOENÇAS (40 ITEMS)
# ============================================================================
echo "📦 BATCH 4: Histórico de Doenças (40 items)"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

if [ -f "scripts/batch_final_4_doencas_EXECUTAVEL.sql" ]; then
    echo "Executando: batch_final_4_doencas_EXECUTAVEL.sql"
    docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/batch_final_4_doencas_EXECUTAVEL.sql > /tmp/batch4_output.txt 2>&1

    if grep -q "COMMIT" /tmp/batch4_output.txt; then
        BATCH4_COUNT=$(grep -c "UPDATE 1" /tmp/batch4_output.txt || echo "0")
        echo "✅ Batch 4 concluído: $BATCH4_COUNT items atualizados"
        TOTAL_ITEMS=$((TOTAL_ITEMS + BATCH4_COUNT))
    else
        echo "❌ Erro no Batch 4 - verificar /tmp/batch4_output.txt"
        cat /tmp/batch4_output.txt
        exit 1
    fi
else
    echo "⚠️  Arquivo batch_final_4_doencas_EXECUTAVEL.sql não encontrado - pulando"
fi
echo ""

# ============================================================================
# BATCH 5 - COMPOSIÇÃO CORPORAL (3 ITEMS)
# ============================================================================
echo "📦 BATCH 5: Composição Corporal (3 items)"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

if [ -f "scripts/enrichment_data/batch_final_5_composicao.sql" ]; then
    echo "Executando: batch_final_5_composicao.sql"
    docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/enrichment_data/batch_final_5_composicao.sql > /tmp/batch5_output.txt 2>&1

    if grep -q "COMMIT" /tmp/batch5_output.txt; then
        BATCH5_COUNT=$(grep -c "UPDATE 1" /tmp/batch5_output.txt || echo "0")
        echo "✅ Batch 5 concluído: $BATCH5_COUNT items atualizados"
        TOTAL_ITEMS=$((TOTAL_ITEMS + BATCH5_COUNT))
    else
        echo "❌ Erro no Batch 5 - verificar /tmp/batch5_output.txt"
        cat /tmp/batch5_output.txt
        exit 1
    fi
else
    echo "⚠️  Arquivo batch_final_5_composicao.sql não encontrado - pulando"
fi
echo ""

# ============================================================================
# VALIDAÇÃO FINAL
# ============================================================================
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "🔍 VALIDAÇÃO FINAL"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

# Contar items enriquecidos no banco
echo "Verificando banco de dados..."
docker compose exec -T db psql -U plenya_user -d plenya_db << 'EOF'
SELECT
  COUNT(*) as total_items,
  COUNT(*) FILTER (WHERE clinical_relevance IS NOT NULL AND LENGTH(clinical_relevance) > 100) as enriched_items,
  COUNT(*) FILTER (WHERE clinical_relevance IS NULL OR LENGTH(clinical_relevance) <= 100) as pending_items,
  ROUND(100.0 * COUNT(*) FILTER (WHERE clinical_relevance IS NOT NULL AND LENGTH(clinical_relevance) > 100) / NULLIF(COUNT(*), 0), 1) as percentage_complete
FROM score_items;
EOF

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "✅ EXECUÇÃO COMPLETA"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "📊 Resumo:"
echo "   • Items processados nesta execução: $TOTAL_ITEMS"
echo "   • Meta: 168 items (5 batches)"
echo ""
echo "📋 Logs salvos em /tmp/batch*_output.txt"
echo ""
echo "🎯 Próximo passo:"
echo "   Verificar quantos items ainda faltam enriquecer"
echo ""
