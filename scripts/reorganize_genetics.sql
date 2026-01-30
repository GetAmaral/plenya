-- ============================================================================
-- SCRIPT DE REORGANIZAÇÃO DO GRUPO GENÉTICA
-- ============================================================================
-- PROBLEMA: 61 exames genéticos estão como GRUPOS em vez de ITEMS
-- SOLUÇÃO: Criar subgrupos lógicos e mover genes para items
-- ============================================================================

BEGIN;

-- 1. Verificar grupo Genética
SELECT id, name FROM score_groups WHERE name = 'Genética';

-- 2. Criar subgrupos lógicos dentro de Genética
WITH genetics_group AS (
    SELECT id FROM score_groups WHERE name = 'Genética'
)
INSERT INTO score_subgroups (group_id, name, description)
SELECT
    g.id,
    sg.name,
    sg.description
FROM genetics_group g
CROSS JOIN (VALUES
    ('Metabolismo', 'Genes relacionados ao metabolismo de nutrientes, energia e hormônios (MTHFR, FTO, MC4R, PPARG, TCF7L2, VDR, etc)'),
    ('Cardiovascular', 'Genes relacionados à saúde cardiovascular e pressão arterial (APOE, ACE, AGT, AGTR1, NOS3, LDLR, etc)'),
    ('Neurodegeneração', 'Genes relacionados a doenças neurodegenerativas (APOE, PSEN1, PSEN2, APP, LRRK2, MAPT, etc)'),
    ('Detoxificação', 'Genes de fase I e II de detoxificação hepática (CYP1A1, CYP1A2, CYP2A6, GSTM1, GSTT1, GSTP1, NAT2, etc)'),
    ('Imunidade', 'Genes relacionados ao sistema imune e autoimunidade (HLA-DQ2/DQ8, IL1B, IL6, TNF, CRP, etc)'),
    ('Performance', 'Genes relacionados a performance física e adaptação ao exercício (ACTN3, ACE I/D, COL5A1, PPARGC1A, etc)'),
    ('Outros', 'Genes que não se enquadram nas categorias anteriores')
) sg(name, description)
ON CONFLICT DO NOTHING;

-- 3. Listar grupos genéticos incorretos que precisam ser movidos
SELECT
    g.id,
    g.name,
    (SELECT COUNT(*) FROM score_subgroups WHERE group_id = g.id) as subgroups_count,
    (SELECT COUNT(*) FROM score_items si
     JOIN score_subgroups sg ON si.subgroup_id = sg.id
     WHERE sg.group_id = g.id) as items_count
FROM score_groups g
WHERE g.name ~ '[A-Z]+[0-9]+'  -- Padrão de gene (ex: CYP1A2, MTHFR)
   OR g.name IN (
       'MTHFR C677T rs1801133 (Homocisteína)',
       'APOE (Alzheimer e Lipídios)',
       'FTO rs9939609 (Obesidade)',
       'MC4R rs17782313 (Obesidade)',
       'TCF7L2 rs7903146 (Diabetes)',
       'PPARG Pro12Ala rs1801282 (Diabetes)',
       'VDR FokI rs2228570 (Vitamina D)',
       'ACE I/D rs4646994 (Hipertensão)'
   )
ORDER BY g.name;

-- NOTA: A conversão de GRUPOS → ITEMS requer mapeamento manual por gene
-- Próximo passo: Criar script Python para migrar cada grupo genético
-- para o subgrupo apropriado e converter seus subgrupos em levels

ROLLBACK;  -- Manter como ROLLBACK até ter certeza da estratégia
