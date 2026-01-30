-- ============================================================================
-- SCRIPT DE LIMPEZA DE DUPLICATAS EM SCORE_ITEMS - VERSÃO EXECUTÁVEL
-- ============================================================================
-- BACKUP CRIADO: backup_before_cleanup_20260127_011846.sql
-- ============================================================================

BEGIN;

-- Criar tabela temporária com IDs a manter vs deletar
CREATE TEMP TABLE items_cleanup AS
WITH duplicates AS (
    SELECT
        name,
        ARRAY_AGG(id ORDER BY
            CASE
                WHEN LENGTH(COALESCE(clinical_relevance, '')) > 100 THEN 0
                ELSE 1
            END,
            created_at NULLS LAST
        ) as ids
    FROM score_items
    GROUP BY name
    HAVING COUNT(*) > 1
)
SELECT
    name,
    ids[1] as keep_id,
    UNNEST(ids[2:]) as delete_id
FROM duplicates;

-- Estatísticas ANTES
SELECT
    'ANTES DA LIMPEZA' as status,
    (SELECT COUNT(*) FROM score_items) as total_items,
    (SELECT COUNT(DISTINCT name) FROM score_items) as nomes_unicos,
    (SELECT COUNT(*) FROM items_cleanup) as items_a_deletar,
    (SELECT COUNT(*) FROM score_levels WHERE item_id IN (SELECT delete_id FROM items_cleanup)) as levels_a_deletar;

-- PASSO 1: Deletar levels dos items duplicados (FK constraint)
DELETE FROM score_levels
WHERE item_id IN (SELECT delete_id FROM items_cleanup);

-- PASSO 2: Deletar items duplicados
DELETE FROM score_items
WHERE id IN (SELECT delete_id FROM items_cleanup);

-- Estatísticas DEPOIS
SELECT
    'DEPOIS DA LIMPEZA' as status,
    (SELECT COUNT(*) FROM score_items) as total_items,
    (SELECT COUNT(DISTINCT name) FROM score_items) as nomes_unicos,
    (SELECT COUNT(*) FROM score_items WHERE LENGTH(COALESCE(clinical_relevance, '')) > 100) as items_com_conteudo;

-- COMMIT (executar de verdade)
COMMIT;
