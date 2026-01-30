-- ANÁLISE DE DUPLICATAS EM SCORE_ITEMS

-- 1. Resumo geral
SELECT
    COUNT(*) as total_items,
    COUNT(DISTINCT name) as nomes_unicos,
    COUNT(*) - COUNT(DISTINCT name) as duplicatas_estimadas
FROM score_items;

-- 2. Top 30 duplicatas com detalhes
WITH duplicates AS (
    SELECT
        name,
        COUNT(*) as total_copies,
        COUNT(CASE WHEN LENGTH(COALESCE(clinical_relevance, '')) > 100 THEN 1 END) as with_content,
        COUNT(CASE WHEN LENGTH(COALESCE(clinical_relevance, '')) <= 100 THEN 1 END) as empty,
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
    total_copies,
    with_content,
    empty,
    ids[1] as keep_id,
    array_length(ids, 1) - 1 as will_delete
FROM duplicates
ORDER BY total_copies DESC, name
LIMIT 30;

-- 3. Estatísticas de items a deletar
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
),
items_to_delete AS (
    SELECT UNNEST(ids[2:]) as item_id
    FROM duplicates
)
SELECT
    (SELECT COUNT(*) FROM items_to_delete) as items_to_delete,
    (SELECT COUNT(*) FROM score_levels WHERE item_id IN (SELECT item_id FROM items_to_delete)) as levels_to_delete;

-- 4. Verificar grupos genéticos incorretos
SELECT COUNT(*) as genetic_groups_as_groups
FROM score_groups
WHERE name ~ '[A-Z]+[0-9]+'  -- Padrão de gene
   OR name IN ('MTHFR', 'APOE', 'FTO', 'MC4R', 'TCF7L2', 'PPARG', 'VDR', 'ACE');
