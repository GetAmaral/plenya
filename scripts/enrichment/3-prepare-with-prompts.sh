#!/bin/bash
# ==============================================================================
# Script 3: Prepare ScoreItems com Prompts para Claude
# ==============================================================================
#
# O QUE FAZ:
#   - Busca top 30 chunks científicos para cada ScoreItem via RAG
#   - Usa threshold adaptativo (0.35→0.30→0.25→0.20→0.15)
#   - Gera 4 prompts prontos para Claude:
#     1. Clinical Relevance (~32K chars)
#     2. Patient Explanation (~32K chars)
#     3. Conduct (~32K chars)
#     4. Max Points (~400 chars)
#   - Inclui fullName (Group - Subgroup - Name) em todos prompts
#   - Salva em score_item_enrichment_preparation
#
# QUANDO USAR:
#   - Após criar/atualizar auto-links
#   - Quando preparations ficam stale (após editar ScoreItems)
#   - Para re-gerar prompts com novo formato
#
# DURAÇÃO: ~10-15 minutos para 878 ScoreItems
# RESULTADO: 878 preparations com 4 prompts cada (100% cobertura)
#
# ==============================================================================

set -e

echo "🔬 SCRIPT 3: Prepare ScoreItems com Prompts"
echo "============================================"
echo ""
echo "Usage: $0 [uuid1 uuid2 ...]"
echo "  Sem args: processa todos sem preparation (incremental)"
echo "  Com args: força reprocessamento dos IDs especificados"
echo ""

# Verificar Docker
if ! docker compose ps | grep -q "api.*Up"; then
    echo "❌ Erro: Docker não está rodando!"
    exit 1
fi

# Estatísticas ANTES
echo "📊 Estado ANTES do prepare:"
docker compose exec -T db psql -U plenya_user -d plenya_db << 'EOSQL'
SELECT
    COUNT(*) as preparations_existentes,
    COUNT(*) FILTER (WHERE prompt_clinical_relevance IS NOT NULL) as com_prompts,
    ROUND(100.0 * COUNT(*) FILTER (WHERE prompt_clinical_relevance IS NOT NULL) /
        NULLIF(COUNT(*), 0), 1) as percent_com_prompts
FROM score_item_enrichment_preparation;
EOSQL

echo ""

if [ "$#" -gt 0 ]; then
    # IDs específicos: reprocessa direto, sem interação
    echo "🎯 Modo específico: reprocessando $# item(s) fornecido(s)..."
    echo ""
else
    # Modo interativo para todos
    echo "Opções:"
    echo "  1) Preparar apenas items SEM preparation (incremental)"
    echo "  2) RE-PREPARAR TODOS (deleta e recria - recomendado após mudanças)"
    echo ""
    read -p "Escolha (1/2): " -n 1 -r OPTION
    echo ""

    if [[ $OPTION == "2" ]]; then
        echo ""
        read -p "⚠️  Isto vai DELETAR todas preparations existentes. Confirma? (s/N): " -n 1 -r
        echo ""
        if [[ ! $REPLY =~ ^[Ss]$ ]]; then
            echo "❌ Cancelado"
            exit 0
        fi

        echo "🗑️  Deletando preparations antigas..."
        docker compose exec -T db psql -U plenya_user -d plenya_db -c \
            "DELETE FROM score_item_enrichment_preparation"
        echo "✅ Deletado"
        echo ""
    fi
fi

echo "🚀 Executando prepare-all com threshold adaptativo..."
echo "   Gerando 4 prompts por ScoreItem (CR, PE, Conduct, Points)"
echo "   Incluindo fullName e 30 chunks científicos completos"
echo ""

# Executar prepare (passa IDs se fornecidos)
docker compose exec api go run /app/cmd/prepare-all/main.go $@ 2>&1 | \
    grep -E "(Preparing|Processed|Prepared|Skipped|Progress)" | \
    while read line; do
        echo "   $line"
    done

echo ""
echo "✅ Preparation concluída!"
echo ""

# Estatísticas DEPOIS
echo "📊 Estado FINAL:"
docker compose exec -T db psql -U plenya_user -d plenya_db << 'EOSQL'
SELECT
    '=== PREPARATIONS ===' as categoria,
    COUNT(*)::text as total,
    COUNT(*) FILTER (WHERE prompt_clinical_relevance IS NOT NULL)::text || ' com 4 prompts' as status
FROM score_item_enrichment_preparation

UNION ALL

SELECT
    '=== QUALITY GRADES ===' as categoria,
    COUNT(*) FILTER (WHERE metadata->>'quality_grade' IN ('excellent', 'good'))::text || '/' ||
    COUNT(*)::text as total,
    ROUND(100.0 * COUNT(*) FILTER (WHERE metadata->>'quality_grade' IN ('excellent', 'good')) / COUNT(*), 1)::text || '% excellent/good' as status
FROM score_item_enrichment_preparation

UNION ALL

SELECT
    '=== AVG CHUNKS ===' as categoria,
    ROUND(AVG((metadata->>'total_chunks')::int), 1)::text as total,
    'chunks por preparation' as status
FROM score_item_enrichment_preparation;

-- Exemplo de 1 preparation
SELECT
    E'\n📋 EXEMPLO: ' || si.name as exemplo,
    E'\n   FullName: ' || si.name || ' (' || sg.name || ' - ' || ssg.name || ')' ||
    E'\n   Chunks: ' || (prep.metadata->>'total_chunks') || ' | Grade: ' || (prep.metadata->>'quality_grade') ||
    E'\n   Prompts:' ||
    E'\n     ✅ Clinical Relevance: ' || LENGTH(prep.prompt_clinical_relevance) || ' chars' ||
    E'\n     ✅ Patient Explanation: ' || LENGTH(prep.prompt_patient_explanation) || ' chars' ||
    E'\n     ✅ Conduct: ' || LENGTH(prep.prompt_conduct) || ' chars' ||
    E'\n     ✅ Max Points: ' || LENGTH(prep.prompt_max_points) || ' chars'
FROM score_item_enrichment_preparation prep
JOIN score_items si ON si.id = prep.score_item_id
JOIN score_subgroups ssg ON ssg.id = si.subgroup_id
JOIN score_groups sg ON sg.id = ssg.group_id
LIMIT 1;
EOSQL

echo ""
echo "✅ SISTEMA PRONTO PARA ENRICHMENT!"
echo ""
echo "💡 Os prompts estão salvos em score_item_enrichment_preparation"
echo "   Cada ScoreItem tem 4 prompts prontos (~32KB cada) para enviar ao Claude"
echo ""
echo "🚀 Próximo passo: Enrichment com Claude (manual ou via API)"
echo ""
