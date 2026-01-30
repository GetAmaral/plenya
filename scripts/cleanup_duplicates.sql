-- ============================================================================
-- SCRIPT DE LIMPEZA DE DUPLICATAS EM SCORE_ITEMS
-- ============================================================================
-- RESUMO: 1.728 items duplicados + 6.112 levels a deletar
-- Critério: Manter APENAS o item com conteúdo (ou mais antigo se todos vazios)
-- ============================================================================

BEGIN;

-- Criar tabela temporária com IDs a manter vs deletar
CREATE TEMP TABLE items_cleanup AS
WITH duplicates AS (
    SELECT
        name,
        ARRAY_AGG(id ORDER BY
            CASE
                WHEN LENGTH(COALESCE(clinical_relevance, '')) > 100 THEN 0  -- Com conteúdo primeiro
                ELSE 1  -- Vazios depois
            END,
            created_at NULLS LAST  -- Mais antigos primeiro se empate
        ) as ids
    FROM score_items
    GROUP BY name
    HAVING COUNT(*) > 1
)
SELECT
    name,
    ids[1] as keep_id,  -- Primeiro da lista ordenada (com conteúdo se existir)
    UNNEST(ids[2:]) as delete_id  -- Resto para deletar
FROM duplicates;

-- Estatísticas antes da limpeza
SELECT
    'ANTES DA LIMPEZA' as status,
    (SELECT COUNT(*) FROM score_items) as total_items,
    (SELECT COUNT(DISTINCT name) FROM score_items) as nomes_unicos,
    (SELECT COUNT(*) FROM items_cleanup) as items_a_deletar,
    (SELECT COUNT(*) FROM score_levels WHERE item_id IN (SELECT delete_id FROM items_cleanup)) as levels_a_deletar;

-- Mostrar amostra dos items que serão mantidos vs deletados
SELECT
    ic.name,
    COUNT(*) as total_copies,
    ic.keep_id,
    ARRAY_AGG(ic.delete_id) as delete_ids
FROM items_cleanup ic
GROUP BY ic.name, ic.keep_id
ORDER BY COUNT(*) DESC
LIMIT 20;

-- IMPORTANTE: Revise os IDs acima antes de continuar!
-- Para executar a limpeza, descomente as linhas abaixo:

-- PASSO 1: Deletar levels dos items duplicados (FK constraint)
-- DELETE FROM score_levels
-- WHERE item_id IN (SELECT delete_id FROM items_cleanup);

-- PASSO 2: Deletar items duplicados
-- DELETE FROM score_items
-- WHERE id IN (SELECT delete_id FROM items_cleanup);

-- Estatísticas depois (descomente após executar as deleções acima)
-- SELECT
--     'DEPOIS DA LIMPEZA' as status,
--     (SELECT COUNT(*) FROM score_items) as total_items,
--     (SELECT COUNT(DISTINCT name) FROM score_items) as nomes_unicos,
--     (SELECT COUNT(*) FROM score_items WHERE LENGTH(COALESCE(clinical_relevance, '')) > 100) as items_com_conteudo;

-- Se tudo estiver OK, COMMIT. Caso contrário, ROLLBACK.
ROLLBACK;  -- Trocar por COMMIT quando tiver certeza
