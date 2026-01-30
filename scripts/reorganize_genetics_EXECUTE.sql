-- ============================================================================
-- REORGANIZAÇÃO DO GRUPO GENÉTICA - VERSÃO EXECUTÁVEL
-- ============================================================================

BEGIN;

-- 1. Verificar/criar grupo Genética (já existe)
SELECT id, name FROM score_groups WHERE name = 'Genética';

-- 2. Criar subgrupos lógicos dentro de Genética
WITH genetics_group AS (
    SELECT id FROM score_groups WHERE name = 'Genética'
)
INSERT INTO score_subgroups (group_id, name, "order")
SELECT
    g.id,
    sg.name,
    sg.ord
FROM genetics_group g
CROSS JOIN (VALUES
    ('Metabolismo', 1),
    ('Cardiovascular', 2),
    ('Neurodegeneração', 3),
    ('Detoxificação', 4),
    ('Imunidade', 5),
    ('Performance', 6),
    ('Outros', 7)
) sg(name, ord)
ON CONFLICT DO NOTHING
RETURNING id, name;

-- 3. Verificar subgrupos criados
SELECT sg.id, sg.name, sg."order"
FROM score_subgroups sg
JOIN score_groups g ON sg.group_id = g.id
WHERE g.name = 'Genética'
ORDER BY sg."order";

COMMIT;
