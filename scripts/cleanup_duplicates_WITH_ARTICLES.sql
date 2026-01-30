-- ============================================================================
-- SCRIPT DE LIMPEZA DE DUPLICATAS - COM MIGRAÇÃO DE ARTIGOS
-- ============================================================================
-- BACKUP: backup_before_cleanup_20260127_011846.sql
-- ============================================================================

BEGIN;

-- Criar tabela temporária com mapeamento keep_id → delete_ids
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
    'ANTES' as momento,
    (SELECT COUNT(*) FROM score_items) as total_items,
    (SELECT COUNT(*) FROM items_cleanup) as items_a_deletar,
    (SELECT COUNT(*) FROM score_levels WHERE item_id IN (SELECT delete_id FROM items_cleanup)) as levels_a_deletar,
    (SELECT COUNT(*) FROM article_score_items WHERE score_item_id IN (SELECT delete_id FROM items_cleanup)) as article_links_a_migrar;

-- PASSO 1: Migrar links de artigos dos items duplicados para os items que vamos manter
-- Evitar duplicação de links (ON CONFLICT DO NOTHING)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT DISTINCT
    asi.article_id,
    ic.keep_id
FROM article_score_items asi
JOIN items_cleanup ic ON asi.score_item_id = ic.delete_id
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- PASSO 2: Deletar links antigos (dos items que serão deletados)
DELETE FROM article_score_items
WHERE score_item_id IN (SELECT delete_id FROM items_cleanup);

-- PASSO 3: Deletar levels dos items duplicados
DELETE FROM score_levels
WHERE item_id IN (SELECT delete_id FROM items_cleanup);

-- PASSO 4: Deletar items duplicados
DELETE FROM score_items
WHERE id IN (SELECT delete_id FROM items_cleanup);

-- Estatísticas DEPOIS
SELECT
    'DEPOIS' as momento,
    (SELECT COUNT(*) FROM score_items) as total_items,
    (SELECT COUNT(DISTINCT name) FROM score_items) as nomes_unicos,
    (SELECT COUNT(*) FROM score_items WHERE LENGTH(COALESCE(clinical_relevance, '')) > 100) as items_com_conteudo,
    (SELECT COUNT(*) FROM article_score_items) as total_article_links;

-- COMMIT
COMMIT;
