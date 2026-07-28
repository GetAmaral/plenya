-- 00064 — Reequilíbrio do escore: café, AINEs separados, peso da alimentação, meio peso no
--         histórico de doenças.
--
-- Origem: estudo docs/emr/estudo-escore-equilibrio-e-devolutiva.md (caso-âncora, 93% com o
-- checklist de ausência valendo 72% do denominador). Decisões do Dr. Getúlio em 2026-07-28.
--
-- (1) CAFÉ. `CAFE_E_CHAS` estava com points=0 (perguntado, classificado, não pontuava) e fora
--     do Continuum | Médico | Inicial. Passa a valer 10 pontos e migra do Continuum | Nutri |
--     Inicial para o Continuum | Médico | Inicial. Os templates não-Continuum ficam como estão;
--     os Continuum de acompanhamento/reavaliação também (eixo de tempo, não de partição).
--
-- (2) AINEs. O item "Analgésicos / Opioides / AINES / Relaxantes musculares" era binário
--     (Sim/Não) e misturava dipirona ocasional com anti-inflamatório crônico — que para o
--     rim, o TGI e a pressão é outra conversa. Vira dois itens:
--       · "Anti-inflamatórios (AINEs)" — NOVO, 12 pts, 4 níveis por frequência de uso
--       · "Analgésicos / Opioides / Relaxantes musculares" — 6 pts, segue binário
--     O item novo entra em todos os templates onde o antigo está.
--
-- (3) ALIMENTAÇÃO. 9 itens `level_choice` com níveis definidos estavam com points 0/NULL —
--     eram avaliados e contribuíam zero (no snapshot do caso-âncora, Frutas e Açúcar casaram N0 e
--     não custaram nada). Passam a valer 10 pontos cada.
--     Ficam de fora de propósito: os 13 campos de texto livre (sem níveis, não pontuáveis) e
--     os 11 booleanos de padrão alimentar (Livre/Vegana/Carnívora/…, cujos níveis são
--     "Não"=N0 / "Sim"=N5 — pontuar significaria zerar o paciente por não seguir 10 dietas
--     ao mesmo tempo).
--
-- (4) HISTÓRICO DE DOENÇAS. Todo item do grupo tem o peso dividido por 2 (12 → 6, 22 → 11).
--     O checklist de ausência deixa de dominar o denominador.

-- +goose Up

-- ███ (1) CAFÉ ██████████████████████████████████████████████████████████████████████████████
UPDATE score_items SET points = 10, updated_at = now()
WHERE anamnese_item_code = 'CAFE_E_CHAS' AND deleted_at IS NULL;

-- entra no Continuum | Médico | Inicial
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order", created_at, updated_at)
SELECT uuid_generate_v7(), '11111111-1111-7111-8111-111111111104', si.id, 0, now(), now()
FROM score_items si
WHERE si.anamnese_item_code = 'CAFE_E_CHAS' AND si.deleted_at IS NULL
  AND NOT EXISTS (
    SELECT 1 FROM anamnesis_template_items x
    WHERE x.anamnesis_template_id = '11111111-1111-7111-8111-111111111104'
      AND x.score_item_id = si.id AND x.deleted_at IS NULL);

-- sai do Continuum | Nutri | Inicial (partição: quem pergunta na linha de base é o médico)
UPDATE anamnesis_template_items t SET deleted_at = now()
FROM score_items si
WHERE t.score_item_id = si.id
  AND si.anamnese_item_code = 'CAFE_E_CHAS'
  AND t.anamnesis_template_id = '11111111-1111-7111-8111-111111111108'
  AND t.deleted_at IS NULL;

-- ███ (4) HISTÓRICO DE DOENÇAS — metade do peso ████████████████████████████████████████████
-- Antes do (2) de propósito: o item de analgésicos recebe o valor final explícito depois.
UPDATE score_items si SET points = si.points / 2, updated_at = now()
FROM score_subgroups sg, score_groups g
WHERE sg.id = si.subgroup_id AND g.id = sg.group_id
  AND g.name = 'Histórico de doenças'
  AND si.deleted_at IS NULL
  AND si.points IS NOT NULL AND si.points > 0;

-- ███ (2) AINEs — item próprio ██████████████████████████████████████████████████████████████
-- abre espaço na ordem alfabética do subgrupo (entre Anti-hipertensivos e Antiosteoporóticos)
UPDATE score_items si SET "order" = si."order" + 1, updated_at = now()
FROM score_subgroups sg, score_groups g
WHERE sg.id = si.subgroup_id AND g.id = sg.group_id
  AND g.name = 'Histórico de doenças' AND sg.name = 'Medicamentos'
  AND si.parent_item_id = '019bf31d-2ef0-78da-9d77-4e8258d3cf8e'
  AND si.deleted_at IS NULL
  AND si."order" >= 14;

INSERT INTO score_items (
  id, name, points, "order", subgroup_id, parent_item_id, anamnese_item_code,
  site_render_type, gender, default_level5, created_at, updated_at
)
SELECT
  '019fa9c0-0000-7000-8000-000000000001'::uuid,
  'Anti-inflamatórios (AINEs)',
  12,
  14,
  si.subgroup_id,
  '019bf31d-2ef0-78da-9d77-4e8258d3cf8e'::uuid,
  'ANTI_INFLAMATORIOS_AINES',
  'level_choice',
  'not_applicable',
  false,
  now(), now()
FROM score_items si
WHERE si.anamnese_item_code = 'ANALGESICOS_OPIOIDES_AINES_RELAXANTES_MUSCULARES'
  AND si.deleted_at IS NULL;

INSERT INTO score_levels (id, level, name, operator, item_id, created_at, updated_at) VALUES
  ('019fa9c0-0000-7000-8000-000000000011'::uuid, 5, 'Não usa',                                   '=', '019fa9c0-0000-7000-8000-000000000001'::uuid, now(), now()),
  ('019fa9c0-0000-7000-8000-000000000012'::uuid, 3, 'Uso esporádico (menos de 2x por mês)',      '=', '019fa9c0-0000-7000-8000-000000000001'::uuid, now(), now()),
  ('019fa9c0-0000-7000-8000-000000000013'::uuid, 1, 'Uso recorrente (2 a 4x por mês)',           '=', '019fa9c0-0000-7000-8000-000000000001'::uuid, now(), now()),
  ('019fa9c0-0000-7000-8000-000000000014'::uuid, 0, 'Uso frequente (mais de 4x por mês)',        '=', '019fa9c0-0000-7000-8000-000000000001'::uuid, now(), now());

-- pilares: herda os do item antigo (Neurológico, Osteomuscular, Medicamentos) e ganha
-- Renal e Gastrointestinal, que são o risco próprio do AINE.
INSERT INTO score_item_method_pillars (method_pillar_id, score_item_id)
SELECT mp.id, '019fa9c0-0000-7000-8000-000000000001'::uuid
FROM method_pillars mp
JOIN method_letters ml ON ml.id = mp.letter_id
WHERE ml.code = 'G'
  AND mp.name IN ('Medicamentos', 'Renal', 'Gastrointestinal', 'Osteomuscular')
  AND mp.deleted_at IS NULL
ON CONFLICT DO NOTHING;

-- item antigo perde os AINEs do rótulo e fica com o peso dos demais analgésicos
UPDATE score_items
SET name = 'Analgésicos / Opioides / Relaxantes musculares',
    points = 6,
    updated_at = now()
WHERE anamnese_item_code = 'ANALGESICOS_OPIOIDES_AINES_RELAXANTES_MUSCULARES'
  AND deleted_at IS NULL;

-- o item novo entra em todo template onde o antigo está
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order", created_at, updated_at)
SELECT uuid_generate_v7(), t.anamnesis_template_id, '019fa9c0-0000-7000-8000-000000000001'::uuid, 0, now(), now()
FROM anamnesis_template_items t
JOIN score_items si ON si.id = t.score_item_id
WHERE si.anamnese_item_code = 'ANALGESICOS_OPIOIDES_AINES_RELAXANTES_MUSCULARES'
  AND t.deleted_at IS NULL
  AND NOT EXISTS (
    SELECT 1 FROM anamnesis_template_items x
    WHERE x.anamnesis_template_id = t.anamnesis_template_id
      AND x.score_item_id = '019fa9c0-0000-7000-8000-000000000001'::uuid
      AND x.deleted_at IS NULL);

-- ███ (3) ALIMENTAÇÃO — peso nos itens classificáveis que estavam zerados ███████████████████
UPDATE score_items SET points = 10, updated_at = now()
WHERE anamnese_item_code IN (
        'AGUA', 'ALCOOL', 'SUCOS',
        'CONSUMO_DE_FRUTAS', 'CONSUMO_DE_ACUCAR', 'ADOCANTES',
        'CONSUMO_DE_PROTEINAS', 'CONSUMO_DE_CALORIAS')
  AND deleted_at IS NULL
  AND (points IS NULL OR points = 0);

-- ███ (5) ORDEM CANÔNICA EM TODOS OS TEMPLATES ██████████████████████████████████████████████
WITH canonica AS (
  SELECT ti.id,
         ROW_NUMBER() OVER (
           PARTITION BY ti.anamnesis_template_id
           ORDER BY g."order", g.name,
                    sg."order", sg.name,
                    COALESCE(pai."order", si."order"),
                    COALESCE(pai.name, si.name),
                    (si.parent_item_id IS NOT NULL),
                    si."order", si.name
         ) AS rn
  FROM anamnesis_template_items ti
  JOIN score_items     si  ON si.id  = ti.score_item_id
  LEFT JOIN score_items pai ON pai.id = si.parent_item_id AND pai.deleted_at IS NULL
  JOIN score_subgroups sg  ON sg.id  = si.subgroup_id
  JOIN score_groups    g   ON g.id   = sg.group_id
  WHERE ti.deleted_at IS NULL
)
UPDATE anamnesis_template_items t SET "order" = c.rn, updated_at = now()
FROM canonica c
WHERE t.id = c.id AND t."order" IS DISTINCT FROM c.rn;

-- +goose Down

DELETE FROM anamnesis_template_items WHERE score_item_id = '019fa9c0-0000-7000-8000-000000000001'::uuid;
DELETE FROM score_item_method_pillars WHERE score_item_id = '019fa9c0-0000-7000-8000-000000000001'::uuid;
DELETE FROM score_levels WHERE item_id = '019fa9c0-0000-7000-8000-000000000001'::uuid;
DELETE FROM score_items WHERE id = '019fa9c0-0000-7000-8000-000000000001'::uuid;

UPDATE score_items si SET "order" = si."order" - 1, updated_at = now()
FROM score_subgroups sg, score_groups g
WHERE sg.id = si.subgroup_id AND g.id = sg.group_id
  AND g.name = 'Histórico de doenças' AND sg.name = 'Medicamentos'
  AND si.parent_item_id = '019bf31d-2ef0-78da-9d77-4e8258d3cf8e'
  AND si.deleted_at IS NULL
  AND si."order" >= 15;

UPDATE score_items
SET name = 'Analgésicos / Opioides / AINES / Relaxantes musculares', points = 12, updated_at = now()
WHERE anamnese_item_code = 'ANALGESICOS_OPIOIDES_AINES_RELAXANTES_MUSCULARES' AND deleted_at IS NULL;

UPDATE score_items si SET points = si.points * 2, updated_at = now()
FROM score_subgroups sg, score_groups g
WHERE sg.id = si.subgroup_id AND g.id = sg.group_id
  AND g.name = 'Histórico de doenças'
  AND si.deleted_at IS NULL
  AND si.anamnese_item_code IS DISTINCT FROM 'ANALGESICOS_OPIOIDES_AINES_RELAXANTES_MUSCULARES'
  AND si.points IS NOT NULL AND si.points > 0;

UPDATE score_items SET points = 0, updated_at = now()
WHERE anamnese_item_code IN (
        'AGUA', 'ALCOOL', 'SUCOS', 'CAFE_E_CHAS',
        'CONSUMO_DE_FRUTAS', 'CONSUMO_DE_ACUCAR', 'ADOCANTES',
        'CONSUMO_DE_PROTEINAS', 'CONSUMO_DE_CALORIAS')
  AND deleted_at IS NULL;

UPDATE anamnesis_template_items t SET deleted_at = NULL, updated_at = now()
FROM score_items si
WHERE t.score_item_id = si.id AND si.anamnese_item_code = 'CAFE_E_CHAS'
  AND t.anamnesis_template_id = '11111111-1111-7111-8111-111111111108';

UPDATE anamnesis_template_items t SET deleted_at = now()
FROM score_items si
WHERE t.score_item_id = si.id AND si.anamnese_item_code = 'CAFE_E_CHAS'
  AND t.anamnesis_template_id = '11111111-1111-7111-8111-111111111104'
  AND t.deleted_at IS NULL;
