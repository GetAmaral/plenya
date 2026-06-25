-- Partição dos 5 templates iniciais do Continuum: cada item pontuável coletável em consulta
-- (níveis>0, sem lab, fora de Genética) é atribuído a EXATAMENTE UM template. Objetivos só na
-- Inicial. Splits clínicos via pilares do método. Labs/Genética/itens não-pontuáveis fora do
-- pool ficam onde estão (não tocados). Idempotente por reconstrução do pool. DEV; prod depois.
--
-- Templates: 104 Inicial · 105 Complemento · 108 Nutrição · 110 Psicologia · 112 Ed. Física

BEGIN;

CREATE TEMP TABLE _assign ON COMMIT DROP AS
WITH pool AS (
  SELECT si.id, g.name AS grp, si.name AS nm, si.anamnese_item_code AS code,
         g."order" AS go, sg."order" AS sgo, si."order" AS io
  FROM score_items si
  JOIN score_subgroups sg ON sg.id = si.subgroup_id
  JOIN score_groups g ON g.id = sg.group_id
  WHERE si.deleted_at IS NULL
    AND si.lab_test_code IS NULL
    AND g.name <> 'Genética'
    AND ((SELECT count(*) FROM score_levels sl WHERE sl.item_id = si.id) > 0
         OR si.id IN (
           SELECT score_item_id FROM anamnesis_template_items
           WHERE anamnesis_template_id IN (
             '11111111-1111-7111-8111-111111111104','11111111-1111-7111-8111-111111111105',
             '11111111-1111-7111-8111-111111111108','11111111-1111-7111-8111-111111111110',
             '11111111-1111-7111-8111-111111111112') AND deleted_at IS NULL))
),
inic AS (
  SELECT score_item_id FROM anamnesis_template_items
  WHERE anamnesis_template_id = '11111111-1111-7111-8111-111111111104' AND deleted_at IS NULL
)
SELECT p.id AS item_id, p.go, p.sgo, p.io,
  CASE
    WHEN p.grp = 'Objetivos' THEN '104'
    WHEN p.code IN ('ESCALA_PHQ_9_27','GAD_7_21','ESCALA_DE_SONOLENCIA_DE_EPWORTH_24','FSS_9','IIEF_5_25','FSFI_36') THEN '104'
    WHEN p.nm ~* 'libido' THEN '104'
    WHEN p.id IN (SELECT score_item_id FROM inic) THEN '104'
    WHEN p.grp IN ('Histórico de doenças','Histórico Familiar de Doenças') THEN '105'
    WHEN p.grp = 'Alimentação' THEN '108'
    WHEN p.grp = 'Movimento e atividade física' THEN '112'
    WHEN p.grp IN ('Cognição','Social','Stress') THEN '110'
    WHEN p.grp = 'Composição corporal' THEN
      CASE WHEN EXISTS (SELECT 1 FROM score_item_method_pillars x WHERE x.score_item_id = p.id
        AND x.method_pillar_id IN ('a91da006-0000-7000-8000-000000000000','a91da004-0000-7000-8000-000000000000'))
      THEN '112' ELSE '108' END
    WHEN p.grp = 'Sono' THEN
      CASE WHEN EXISTS (SELECT 1 FROM score_item_method_pillars x WHERE x.score_item_id = p.id
        AND x.method_pillar_id IN ('a91dd004-0000-7000-8000-000000000000','a91dd003-0000-7000-8000-000000000000','a91dd002-0000-7000-8000-000000000000'))
      THEN '110' ELSE '105' END
    WHEN p.grp = 'Vida Sexual' THEN
      CASE WHEN p.nm ~* 'desenvolvimento|abuso|trauma|reprodutiv|v[ií]nculo' THEN '110' ELSE '105' END
    ELSE '105'
  END AS tgt
FROM pool p;

-- Remove os vínculos atuais desses itens nos 5 templates (labs/Genética/não-pontuáveis intactos).
DELETE FROM anamnesis_template_items ati USING _assign a
WHERE ati.score_item_id = a.item_id
  AND ati.anamnesis_template_id IN (
    '11111111-1111-7111-8111-111111111104','11111111-1111-7111-8111-111111111105',
    '11111111-1111-7111-8111-111111111108','11111111-1111-7111-8111-111111111110',
    '11111111-1111-7111-8111-111111111112');

-- Reinsere cada item no template-alvo, ordenado por grupo/subgrupo/item.
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order", created_at, updated_at)
SELECT
  gen_random_uuid(),
  CASE a.tgt
    WHEN '104' THEN '11111111-1111-7111-8111-111111111104'::uuid
    WHEN '105' THEN '11111111-1111-7111-8111-111111111105'::uuid
    WHEN '108' THEN '11111111-1111-7111-8111-111111111108'::uuid
    WHEN '110' THEN '11111111-1111-7111-8111-111111111110'::uuid
    WHEN '112' THEN '11111111-1111-7111-8111-111111111112'::uuid
  END,
  a.item_id,
  row_number() OVER (PARTITION BY a.tgt ORDER BY a.go, a.sgo, a.io),
  now(), now()
FROM _assign a;

COMMIT;
