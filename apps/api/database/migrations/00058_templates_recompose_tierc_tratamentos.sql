-- Converte para goose os deltas de curadoria de anamnese que viviam no bundle manual
-- (docs/emr/prod-migration-anamnese.sql) e que ficaram pendentes em prod:
--   (1) Renome dos 13 templates (idempotente; já em prod, mantido p/ reprodutibilidade).
--   (2) Tier C (FSS/FSFI/PSQI) soft-delete de score_items + refs em templates.
--   (3) Recompose: moves de itens entre templates. Escalas movidas por CÓDIGO (a 00057
--       removeu o "____/NN" do nome, então busca por nome não casa mais). Inclui o swap
--       ASEX<->IIEF-5 (Continuum Médico Inicial<->Complemento) que faltava em prod.
--   (4) "Tratamentos em uso p/ modificar composição corporal": ajuste do nível 5.
-- Tudo idempotente e portável (itens por nome/código; templates por UUID fixo).
-- Parte equivalente "ASEX /25->/30" do bundle é obsoleta (00057 tirou o número do nome).

-- +goose Up

-- ███ (1) RENOME DOS TEMPLATES ███████████████████████████████████████████████
UPDATE anamnesis_templates SET name='Médico | Inicial',                            updated_at=now() WHERE id='11111111-1111-7111-8111-111111111101';
UPDATE anamnesis_templates SET name='Médico | Acompanhamento',                     updated_at=now() WHERE id='11111111-1111-7111-8111-111111111102';
UPDATE anamnesis_templates SET name='Médico | Revisão de Exames',                  updated_at=now() WHERE id='11111111-1111-7111-8111-111111111103';
UPDATE anamnesis_templates SET name='Continuum | Médico | Inicial',               updated_at=now() WHERE id='11111111-1111-7111-8111-111111111104';
UPDATE anamnesis_templates SET name='Continuum | Médico | Complemento',           updated_at=now() WHERE id='11111111-1111-7111-8111-111111111105';
UPDATE anamnesis_templates SET name='Continuum | Médico | Acompanhamento',        updated_at=now() WHERE id='11111111-1111-7111-8111-111111111106';
UPDATE anamnesis_templates SET name='Continuum | Médico | Reavaliação Trimestral',updated_at=now() WHERE id='11111111-1111-7111-8111-111111111107';
UPDATE anamnesis_templates SET name='Continuum | Nutri | Inicial',                updated_at=now() WHERE id='11111111-1111-7111-8111-111111111108';
UPDATE anamnesis_templates SET name='Continuum | Nutri | Acompanhamento',         updated_at=now() WHERE id='11111111-1111-7111-8111-111111111109';
UPDATE anamnesis_templates SET name='Continuum | Psico | Inicial',                updated_at=now() WHERE id='11111111-1111-7111-8111-111111111110';
UPDATE anamnesis_templates SET name='Continuum | Psico | Acompanhamento',         updated_at=now() WHERE id='11111111-1111-7111-8111-111111111111';
UPDATE anamnesis_templates SET name='Continuum | Ed. Física | Inicial',           updated_at=now() WHERE id='11111111-1111-7111-8111-111111111112';
UPDATE anamnesis_templates SET name='Continuum | Ed. Física | Acompanhamento',    updated_at=now() WHERE id='11111111-1111-7111-8111-111111111113';

-- ███ (2) TIER C — soft-delete ███████████████████████████████████████████████
UPDATE anamnesis_template_items ati SET deleted_at = now()
FROM score_items si
WHERE ati.score_item_id = si.id
  AND si.anamnese_item_code IN ('FSS_9','FSFI_36','PSQI_21')
  AND ati.deleted_at IS NULL;
UPDATE score_items SET deleted_at = now()
WHERE anamnese_item_code IN ('FSS_9','FSFI_36','PSQI_21')
  AND deleted_at IS NULL;

-- ███ (3) RECOMPOSE — moves de itens entre templates █████████████████████████
CREATE TEMP TABLE moves(src uuid, dst uuid, grupo text, subgrupo text, item text, code text) ON COMMIT DROP;
-- Moves por NOME (itens regulares: grupo+subgrupo+item; item=NULL => subgrupo inteiro)
INSERT INTO moves (src, dst, grupo, subgrupo, item) VALUES
 ('11111111-1111-7111-8111-111111111104','11111111-1111-7111-8111-111111111105',
  'Objetivos','Percepção de futuro (6m-5a-10a-30a)', NULL),
 ('11111111-1111-7111-8111-111111111104','11111111-1111-7111-8111-111111111105',
  'Vida Sexual','Atual','Uso recente outros medicamentos/suplementos para libido/desempenho sexual'),
 ('11111111-1111-7111-8111-111111111108','11111111-1111-7111-8111-111111111104',
  'Alimentação','Atual (últmos 6 meses)','Consumo de Frutas'),
 ('11111111-1111-7111-8111-111111111108','11111111-1111-7111-8111-111111111104',
  'Alimentação','Atual (últmos 6 meses)','Consumo de Verduras e Legumes'),
 ('11111111-1111-7111-8111-111111111108','11111111-1111-7111-8111-111111111104',
  'Alimentação','Atual (últmos 6 meses)','Consumo de Proteínas (cálculos com base no recordatório)'),
 ('11111111-1111-7111-8111-111111111108','11111111-1111-7111-8111-111111111104',
  'Composição corporal','Atual','Tratamentos em uso para modificar composição corporal'),
 ('11111111-1111-7111-8111-111111111105','11111111-1111-7111-8111-111111111104',
  'Histórico de doenças','Cirurgias já realizadas', NULL),
 ('11111111-1111-7111-8111-111111111105','11111111-1111-7111-8111-111111111104',
  'Histórico de doenças','Hábitos e vícios nocivos (Questionar ativamente sobre uso passado ou atual):', NULL),
 ('11111111-1111-7111-8111-111111111105','11111111-1111-7111-8111-111111111110',
  'Sono','Histórico', NULL),
 ('11111111-1111-7111-8111-111111111110','11111111-1111-7111-8111-111111111104',
  'Cognição','Atual','Uso atual de psicotrópicos para cognição'),
 ('11111111-1111-7111-8111-111111111112','11111111-1111-7111-8111-111111111104',
  'Movimento e atividade física','Histórico','Lesões relacionadas ao exercício'),
 ('11111111-1111-7111-8111-111111111112','11111111-1111-7111-8111-111111111104',
  'Movimento e atividade física','Histórico','Cirurgias realizadas relacionadas ao exercício');
-- Moves por CÓDIGO (escalas — robusto ao nome sem "____/NN")
INSERT INTO moves (src, dst, code) VALUES
 ('11111111-1111-7111-8111-111111111105','11111111-1111-7111-8111-111111111104','ASEX_25'),
 ('11111111-1111-7111-8111-111111111104','11111111-1111-7111-8111-111111111105','IIEF_5_25'),
 ('11111111-1111-7111-8111-111111111104','11111111-1111-7111-8111-111111111110','ESCALA_PHQ_9_27'),
 ('11111111-1111-7111-8111-111111111104','11111111-1111-7111-8111-111111111110','GAD_7_21'),
 ('11111111-1111-7111-8111-111111111110','11111111-1111-7111-8111-111111111104','5_PALAVRAS_DE_DUBOIS_IMEDIATO_5'),
 ('11111111-1111-7111-8111-111111111110','11111111-1111-7111-8111-111111111104','5_PALAVRAS_DE_DUBOIS_TARDIO_5'),
 ('11111111-1111-7111-8111-111111111110','11111111-1111-7111-8111-111111111104','SPAN_DE_DIGITOS_DIRETO_8'),
 ('11111111-1111-7111-8111-111111111110','11111111-1111-7111-8111-111111111104','SPAN_DE_DIGITOS_INVERSO_7');

CREATE TEMP TABLE resolved ON COMMIT DROP AS
  SELECT m.src, m.dst, si.id AS score_item_id
  FROM moves m
  JOIN score_groups    g  ON g.name  = m.grupo    AND g.deleted_at  IS NULL
  JOIN score_subgroups sg ON sg.group_id = g.id   AND sg.name = m.subgrupo AND sg.deleted_at IS NULL
  JOIN score_items     si ON si.subgroup_id = sg.id AND si.deleted_at IS NULL
                          AND (m.item IS NULL OR si.name = m.item)
  WHERE m.code IS NULL
UNION
  SELECT m.src, m.dst, si.id
  FROM moves m
  JOIN score_items si ON si.anamnese_item_code = m.code AND si.deleted_at IS NULL
  WHERE m.code IS NOT NULL;

-- remove do template de origem
UPDATE anamnesis_template_items t SET deleted_at = now()
FROM resolved r
WHERE t.anamnesis_template_id = r.src AND t.score_item_id = r.score_item_id AND t.deleted_at IS NULL;

-- adiciona no destino (se ainda não tiver ativo)
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order", created_at, updated_at)
SELECT uuid_generate_v7(), r.dst, r.score_item_id, 0, now(), now()
FROM resolved r
WHERE NOT EXISTS (
  SELECT 1 FROM anamnesis_template_items x
  WHERE x.anamnesis_template_id = r.dst AND x.score_item_id = r.score_item_id AND x.deleted_at IS NULL);

-- renumera "order" (ordem natural g/sg/i) nos templates tocados
WITH touched AS (
  SELECT src AS tid FROM resolved UNION SELECT dst FROM resolved
), o2 AS (
  SELECT ti.id,
         ROW_NUMBER() OVER (PARTITION BY ti.anamnesis_template_id
                            ORDER BY g."order", sg."order", si."order") AS rn
  FROM anamnesis_template_items ti
  JOIN score_items     si ON si.id = ti.score_item_id
  JOIN score_subgroups sg ON sg.id = si.subgroup_id
  JOIN score_groups    g  ON g.id  = sg.group_id
  WHERE ti.anamnesis_template_id IN (SELECT tid FROM touched) AND ti.deleted_at IS NULL
)
UPDATE anamnesis_template_items t SET "order" = o2.rn, updated_at = now()
FROM o2 WHERE t.id = o2.id;

-- ███ (4) TRATAMENTOS composição — nível 5 ██████████████████████████████████
UPDATE score_levels sl
SET name        = 'Nenhum ou Acompanhamento nutri/educador físico adequado',
    site_legend = 'Nenhum tratamento em uso para mudar a composição do corpo, ou acompanhamento adequado com nutricionista e educador físico.',
    updated_at  = now()
FROM score_items si
WHERE sl.item_id = si.id
  AND si.anamnese_item_code = 'TRATAMENTOS_EM_USO_PARA_MODIFICAR_COMPOSICAO_CORPORAL'
  AND sl.level = 5 AND sl.deleted_at IS NULL;

-- +goose Down
-- Curadoria de composição de templates: reversão fiel exigiria snapshot das posições e
-- nomes anteriores (não capturados aqui). No-op proposital; restaurar do backup se preciso.
SELECT 1;
