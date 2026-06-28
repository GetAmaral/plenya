-- ============================================================================
-- Escore — adiciona "Hipertensão arterial" em Histórico de doenças › Doenças crônicas
-- ============================================================================
-- Lacuna: HAS não existia como doença crônica (só "Anti-hipertensivos" em Medicamentos).
-- Posicionada ANTES de "Insuficiência cardíaca" (order 7; IC é 8).
-- Níveis espelham o Diabetes (matriz duração × controle) — análogo crônico/controlável.
-- points=25 (igual Diabetes e Doença renal crônica). default_level5=true (L5 = "Não tenho").
-- Adicionada aos 5 templates que contêm os irmãos do subgrupo.
-- Idempotente: identifica por anamnese_item_code; só insere o que faltar; renumera ordem.
-- ============================================================================
BEGIN;

-- 1) score_item HAS (insere se não existir) ----------------------------------
INSERT INTO score_items (id, name, anamnese_item_code, subgroup_id, "order", points, default_level5, gender, created_at, updated_at)
SELECT uuid_generate_v7(), 'Hipertensão arterial', 'HIPERTENSAO_ARTERIAL', sg.id, 7, 25, true, 'not_applicable', now(), now()
FROM score_subgroups sg JOIN score_groups g ON g.id = sg.group_id
WHERE g.name = 'Histórico de doenças' AND sg.name LIKE 'Doenças crônicas%'
  AND NOT EXISTS (SELECT 1 FROM score_items x WHERE x.anamnese_item_code = 'HIPERTENSAO_ARTERIAL' AND x.deleted_at IS NULL);

-- 2) 6 níveis (matriz duração × controle, espelhando Diabetes) ----------------
INSERT INTO score_levels (id, item_id, level, operator, name, site_legend, created_at, updated_at)
SELECT uuid_generate_v7(), si.id, v.level, '=', v.lname, v.legend, now(), now()
FROM score_items si
CROSS JOIN (VALUES
  (0, 'Tenho há mais de 5 anos, com controle inadequado', 'Hipertensão há mais de cinco anos e mal controlada, exige cuidado intensivo.'),
  (1, 'Tenho há menos de 5 anos, com controle inadequado', 'Hipertensão recente com a pressão ainda mal controlada, requer ajuste.'),
  (2, 'Tenho há mais de 5 anos, com controle adequado', 'Hipertensão de longa data, porém bem controlada.'),
  (3, 'Tenho há menos de 5 anos, com controle adequado', 'Hipertensão há menos de cinco anos e bem controlada.'),
  (4, 'Não tenho, mas tenho histórico familiar forte', 'Sem hipertensão, mas com histórico forte na família.'),
  (5, 'Não tenho', 'Não tem hipertensão, situação favorável.')
) AS v(level, lname, legend)
WHERE si.anamnese_item_code = 'HIPERTENSAO_ARTERIAL' AND si.deleted_at IS NULL
  AND NOT EXISTS (SELECT 1 FROM score_levels sl WHERE sl.item_id = si.id AND sl.level = v.level AND sl.deleted_at IS NULL);

-- 3) adiciona aos 5 templates que têm os irmãos (Doenças crônicas) ------------
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order", created_at, updated_at)
SELECT uuid_generate_v7(), t.id, si.id, 0, now(), now()
FROM score_items si
JOIN anamnesis_templates t ON t.id IN (
  '11111111-1111-7111-8111-111111111101', -- Médico | Inicial
  '11111111-1111-7111-8111-111111111102', -- Médico | Acompanhamento
  '11111111-1111-7111-8111-111111111104', -- Continuum | Médico | Inicial
  '11111111-1111-7111-8111-111111111106', -- Continuum | Médico | Acompanhamento
  '11111111-1111-7111-8111-111111111107') -- Continuum | Médico | Reavaliação Trimestral
WHERE si.anamnese_item_code = 'HIPERTENSAO_ARTERIAL' AND si.deleted_at IS NULL
  AND NOT EXISTS (
    SELECT 1 FROM anamnesis_template_items x
    WHERE x.anamnesis_template_id = t.id AND x.score_item_id = si.id AND x.deleted_at IS NULL);

-- 4) renumera "order" (natural g/sg/i) nos 5 templates tocados ----------------
WITH o AS (
  SELECT ti.id, ROW_NUMBER() OVER (PARTITION BY ti.anamnesis_template_id
                                   ORDER BY g."order", sg."order", si."order") AS rn
  FROM anamnesis_template_items ti
  JOIN score_items si ON si.id = ti.score_item_id
  JOIN score_subgroups sg ON sg.id = si.subgroup_id
  JOIN score_groups g ON g.id = sg.group_id
  WHERE ti.anamnesis_template_id IN (
    '11111111-1111-7111-8111-111111111101','11111111-1111-7111-8111-111111111102',
    '11111111-1111-7111-8111-111111111104','11111111-1111-7111-8111-111111111106',
    '11111111-1111-7111-8111-111111111107') AND ti.deleted_at IS NULL
)
UPDATE anamnesis_template_items t SET "order" = o.rn, updated_at = now() FROM o WHERE t.id = o.id;

-- conferência ----------------------------------------------------------------
SELECT si.name, si.points, si."order",
  (SELECT COUNT(*) FROM score_levels sl WHERE sl.item_id=si.id AND sl.deleted_at IS NULL) AS niveis,
  (SELECT COUNT(*) FROM anamnesis_template_items ti WHERE ti.score_item_id=si.id AND ti.deleted_at IS NULL) AS em_templates
FROM score_items si WHERE si.anamnese_item_code='HIPERTENSAO_ARTERIAL' AND si.deleted_at IS NULL;

COMMIT;
