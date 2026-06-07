-- ============================================================================
-- Seed dos 13 templates de anamnese (Escore Plenya)
-- Plano: docs/emr/plano-templates-anamnese.md · Curadoria: curadoria-templates-anamnese.md
-- Idempotente: templates com UUID fixo (upsert), itens deletados+reinseridos.
-- Universo = itens de anamnese (lab_test_code IS NULL). Exames/Genética ficam fora.
-- Escolhas item-a-item incluem o pai (parent_item_id, profundidade 1).
-- Rodar dentro de transação; trocar ROLLBACK->COMMIT após conferir os counts.
-- ============================================================================
BEGIN;

-- 0) Base: todos os itens de anamnese com chaves de ordenação ----------------
CREATE TEMP TABLE ai AS
SELECT si.id, si.parent_item_id,
       g.name  AS grupo,    g."order"  AS g_ord,
       sg.name AS subgrupo, sg."order" AS sg_ord,
       si.name AS item,     si."order" AS i_ord
FROM score_items si
JOIN score_subgroups sg ON sg.id = si.subgroup_id
JOIN score_groups   g  ON g.id  = sg.group_id
WHERE si.deleted_at IS NULL AND si.lab_test_code IS NULL;

-- 1) Aposenta os placeholders ------------------------------------------------
UPDATE anamnesis_templates SET deleted_at = now()
WHERE name IN ('Teste','Medica Inicial') AND deleted_at IS NULL;

-- 2) Upsert dos 13 templates (UUID fixo p/ idempotência) ---------------------
INSERT INTO anamnesis_templates (id, name, area, created_at, updated_at) VALUES
 ('11111111-1111-7111-8111-111111111101','Avaliação Médica Inicial','Medicina',now(),now()),
 ('11111111-1111-7111-8111-111111111102','Acompanhamento Médico','Medicina',now(),now()),
 ('11111111-1111-7111-8111-111111111103','Revisão de Exames','Medicina',now(),now()),
 ('11111111-1111-7111-8111-111111111104','Avaliação Médica de Entrada (Continuum)','Medicina',now(),now()),
 ('11111111-1111-7111-8111-111111111105','Complemento da Avaliação Médica (Continuum)','Medicina',now(),now()),
 ('11111111-1111-7111-8111-111111111106','Acompanhamento Médico (Continuum)','Medicina',now(),now()),
 ('11111111-1111-7111-8111-111111111107','Reavaliação Médica Trimestral (Continuum)','Medicina',now(),now()),
 ('11111111-1111-7111-8111-111111111108','Avaliação Nutricional Inicial (Continuum)','Nutricao',now(),now()),
 ('11111111-1111-7111-8111-111111111109','Acompanhamento Nutricional (Continuum)','Nutricao',now(),now()),
 ('11111111-1111-7111-8111-111111111110','Avaliação Psicológica Inicial (Continuum)','Psicologia',now(),now()),
 ('11111111-1111-7111-8111-111111111111','Acompanhamento Psicológico (Continuum)','Psicologia',now(),now()),
 ('11111111-1111-7111-8111-111111111112','Avaliação Física Inicial (Continuum)','Educacao Fisica',now(),now()),
 ('11111111-1111-7111-8111-111111111113','Acompanhamento (Educação Física) (Continuum)','Educacao Fisica',now(),now())
ON CONFLICT (id) DO UPDATE SET name=EXCLUDED.name, area=EXCLUDED.area, updated_at=now(), deleted_at=NULL;

-- 3) Limpa itens existentes dos 13 -------------------------------------------
DELETE FROM anamnesis_template_items
WHERE anamnesis_template_id IN (
  '11111111-1111-7111-8111-111111111101','11111111-1111-7111-8111-111111111102',
  '11111111-1111-7111-8111-111111111103','11111111-1111-7111-8111-111111111104',
  '11111111-1111-7111-8111-111111111105','11111111-1111-7111-8111-111111111106',
  '11111111-1111-7111-8111-111111111107','11111111-1111-7111-8111-111111111108',
  '11111111-1111-7111-8111-111111111109','11111111-1111-7111-8111-111111111110',
  '11111111-1111-7111-8111-111111111111','11111111-1111-7111-8111-111111111112',
  '11111111-1111-7111-8111-111111111113');

-- 4) Componentes reutilizáveis (id-only) -------------------------------------
-- Rastreio Alimentação (13) + pai
CREATE TEMP TABLE t_rastreio_alim AS
WITH leaves AS (
  SELECT id FROM ai WHERE grupo='Alimentação' AND item IN (
    'Padrão alimentar atual','Água (30 ml/kg/dia)','Líquidos no dia','Álcool (discriminar o tipo)',
    'Consumo de Açúcar','Refrigerantes e energéticos','Consumo de Frutas','Consumo de Verduras e Legumes',
    'Café e chás (com cafeína - chá preto, chá verde, chá mate)','Intolerâncias','Restrições alimentares',
    'Suplementações utilizadas (marcas e doses)','Tem suplementação prescrita? Está usando adequadamente?'))
SELECT id FROM leaves
UNION SELECT parent_item_id FROM ai WHERE id IN (SELECT id FROM leaves)
      AND parent_item_id IS NOT NULL AND parent_item_id IN (SELECT id FROM ai);

-- Rastreio Movimento (6) + pai
CREATE TEMP TABLE t_rastreio_mov AS
WITH leaves AS (
  SELECT id FROM ai WHERE grupo='Movimento e atividade física' AND item IN (
    'Estratégia macro atual','Divisão das atividades','Restrições de atividades',
    'Lesões relacionadas ao exercício','Cirurgias realizadas relacionadas ao exercício',
    'Piores fases de sedentarismo (idade, duração, fatores associados)'))
SELECT id FROM leaves
UNION SELECT parent_item_id FROM ai WHERE id IN (SELECT id FROM leaves)
      AND parent_item_id IS NOT NULL AND parent_item_id IN (SELECT id FROM ai);

-- Grupos médicos integrais (núcleo A1)
CREATE TEMP TABLE t_med_groups AS
SELECT id FROM ai WHERE grupo IN (
  'Objetivos','Histórico de doenças','Histórico Familiar de Doenças','Sono',
  'Vida Sexual','Composição corporal','Cognição','Stress','Social');

-- A1 = núcleo médico + rastreios
CREATE TEMP TABLE t_a1 AS
SELECT id FROM t_med_groups
UNION SELECT id FROM t_rastreio_alim
UNION SELECT id FROM t_rastreio_mov;

-- B1 cross-AGIR (escolhas item-a-item) + pais
CREATE TEMP TABLE t_b1 AS
WITH leaves AS (
  SELECT id FROM ai WHERE grupo='Objetivos'
  UNION ALL SELECT id FROM ai WHERE grupo='Alimentação' AND item IN (
    'Padrão alimentar atual','Álcool (discriminar o tipo)','Consumo de Açúcar',
    'Refrigerantes e energéticos','Água (30 ml/kg/dia)','Suplementações utilizadas (marcas e doses)')
  UNION ALL SELECT id FROM ai WHERE grupo='Movimento e atividade física' AND item IN (
    'Estratégia macro atual','Divisão das atividades')
  UNION ALL SELECT id FROM ai WHERE grupo='Histórico de doenças' AND subgrupo LIKE 'Doenças crônicas%'
  UNION ALL SELECT id FROM ai WHERE grupo='Histórico de doenças' AND subgrupo='Medicamentos'
  UNION ALL SELECT id FROM ai WHERE grupo='Histórico Familiar de Doenças' AND subgrupo LIKE 'Parentes próximos%'
  UNION ALL SELECT id FROM ai WHERE grupo='Composição corporal' AND subgrupo='Medidas Objetivas' AND item IN (
    'Peso (kg)','Altura (cm)','IMC (kg/m²)','Abdominal (cintura em cm) - homem','Abdominal (cintura em cm) - mulher',
    'Razão cintura/altura (cm/cm)','% Gordura corporal - homem','% Gordura corporal - mulher',
    'Gordura visceral (cm²) - homem','Gordura visceral (cm²) - mulher','Massa muscular esquelética (kg)',
    'Ângulo de fase (°) - homem','Ângulo de fase (°) - mulher')
  UNION ALL SELECT id FROM ai WHERE grupo='Cognição' AND subgrupo='Atual' AND item IN (
    'Escala PHQ-9 (humor): ___/27','GAD-7 (ansiedade): ___/21','Escala de Sonolência de Epworth: ___/24',
    'Disposição/energia','Capacidade da memória percebida','Capacidade de foco/concentração/aprendizado')
  UNION ALL SELECT id FROM ai WHERE grupo='Stress' AND subgrupo='Atual'
  UNION ALL SELECT id FROM ai WHERE grupo='Social' AND subgrupo='Atual' AND item IN (
    'Situação conjugal','Situação familiar','Profissões','Lazer','Hobbies')
  UNION ALL SELECT id FROM ai WHERE grupo='Sono' AND subgrupo='Atual' AND item IN (
    'Qualidade percebida do sono','Tempo de sono','Hora de dormir','Hora de acordar','Regularidade no dormir',
    'Roncos','Apneias','Insônia / dificuldade em iniciar o sono','Uso atual de medicamentos/suplementos para dormir')
)
SELECT id FROM leaves
UNION SELECT parent_item_id FROM ai WHERE id IN (SELECT id FROM leaves)
      AND parent_item_id IS NOT NULL AND parent_item_id IN (SELECT id FROM ai);

-- A2/B3 (acompanhamento médico mensal — dinâmico leve)
CREATE TEMP TABLE t_a2 AS
SELECT id FROM ai WHERE grupo='Objetivos'
UNION SELECT id FROM ai WHERE grupo='Histórico de doenças' AND (subgrupo LIKE 'Doenças crônicas%' OR subgrupo='Medicamentos')
UNION SELECT id FROM ai WHERE grupo='Composição corporal' AND subgrupo='Medidas Objetivas'
UNION SELECT id FROM ai WHERE grupo='Sono' AND subgrupo='Atual'
UNION SELECT id FROM ai WHERE grupo='Cognição' AND subgrupo='Atual'
UNION SELECT id FROM t_rastreio_alim
UNION SELECT id FROM t_rastreio_mov;

-- B4 (reavaliação trimestral — recalcular escore: A2 + sexual/social/stress atual + composição inteira)
CREATE TEMP TABLE t_b4 AS
SELECT id FROM t_a2
UNION SELECT id FROM ai WHERE grupo='Vida Sexual' AND subgrupo='Atual'
UNION SELECT id FROM ai WHERE grupo='Social' AND subgrupo='Atual'
UNION SELECT id FROM ai WHERE grupo='Stress'
UNION SELECT id FROM ai WHERE grupo='Composição corporal';

-- 5) Inserts por template (ordem natural g_ord, sg_ord, i_ord) ---------------
-- helper macro: INSERT ... SELECT tpl, x.id, row_number() ... FROM <set> x JOIN ai a USING(id)

-- A1
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order")
SELECT gen_random_uuid(), '11111111-1111-7111-8111-111111111101', a.id, ROW_NUMBER() OVER (ORDER BY a.g_ord,a.sg_ord,a.i_ord)
FROM t_a1 x JOIN ai a ON a.id=x.id;

-- A2  (= B3)
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order")
SELECT gen_random_uuid(), '11111111-1111-7111-8111-111111111102', a.id, ROW_NUMBER() OVER (ORDER BY a.g_ord,a.sg_ord,a.i_ord)
FROM t_a2 x JOIN ai a ON a.id=x.id;

-- A3 (mínimo: Medicamentos)
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order")
SELECT gen_random_uuid(), '11111111-1111-7111-8111-111111111103', a.id, ROW_NUMBER() OVER (ORDER BY a.g_ord,a.sg_ord,a.i_ord)
FROM ai a WHERE a.grupo='Histórico de doenças' AND a.subgrupo='Medicamentos';

-- B1
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order")
SELECT gen_random_uuid(), '11111111-1111-7111-8111-111111111104', a.id, ROW_NUMBER() OVER (ORDER BY a.g_ord,a.sg_ord,a.i_ord)
FROM t_b1 x JOIN ai a ON a.id=x.id;

-- B2 = A1 - B1
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order")
SELECT gen_random_uuid(), '11111111-1111-7111-8111-111111111105', a.id, ROW_NUMBER() OVER (ORDER BY a.g_ord,a.sg_ord,a.i_ord)
FROM (SELECT id FROM t_a1 EXCEPT SELECT id FROM t_b1) x JOIN ai a ON a.id=x.id;

-- B3 = A2
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order")
SELECT gen_random_uuid(), '11111111-1111-7111-8111-111111111106', a.id, ROW_NUMBER() OVER (ORDER BY a.g_ord,a.sg_ord,a.i_ord)
FROM t_a2 x JOIN ai a ON a.id=x.id;

-- B4
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order")
SELECT gen_random_uuid(), '11111111-1111-7111-8111-111111111107', a.id, ROW_NUMBER() OVER (ORDER BY a.g_ord,a.sg_ord,a.i_ord)
FROM t_b4 x JOIN ai a ON a.id=x.id;

-- B5 Nutri Inicial = Alimentação + Composição corporal + Objetivos
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order")
SELECT gen_random_uuid(), '11111111-1111-7111-8111-111111111108', a.id, ROW_NUMBER() OVER (ORDER BY a.g_ord,a.sg_ord,a.i_ord)
FROM ai a WHERE a.grupo IN ('Alimentação','Composição corporal','Objetivos');

-- B6 Nutri Acomp = Alimentação Atual + Composição (Medidas + Atual)
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order")
SELECT gen_random_uuid(), '11111111-1111-7111-8111-111111111109', a.id, ROW_NUMBER() OVER (ORDER BY a.g_ord,a.sg_ord,a.i_ord)
FROM ai a WHERE (a.grupo='Alimentação' AND a.subgrupo LIKE 'Atual%')
            OR (a.grupo='Composição corporal' AND a.subgrupo IN ('Medidas Objetivas','Atual'));

-- B7 Psico Inicial = Cognição + Stress + Social + Sono + Objetivos
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order")
SELECT gen_random_uuid(), '11111111-1111-7111-8111-111111111110', a.id, ROW_NUMBER() OVER (ORDER BY a.g_ord,a.sg_ord,a.i_ord)
FROM ai a WHERE a.grupo IN ('Cognição','Stress','Social','Sono','Objetivos');

-- B8 Psico Acomp = Atual de Cognição/Stress/Social/Sono
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order")
SELECT gen_random_uuid(), '11111111-1111-7111-8111-111111111111', a.id, ROW_NUMBER() OVER (ORDER BY a.g_ord,a.sg_ord,a.i_ord)
FROM ai a WHERE a.grupo IN ('Cognição','Stress','Social','Sono') AND a.subgrupo='Atual';

-- B9 Ed.Física Inicial = Movimento (inteiro, inc Testes) + Composição + Objetivos
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order")
SELECT gen_random_uuid(), '11111111-1111-7111-8111-111111111112', a.id, ROW_NUMBER() OVER (ORDER BY a.g_ord,a.sg_ord,a.i_ord)
FROM ai a WHERE a.grupo IN ('Movimento e atividade física','Composição corporal','Objetivos');

-- B10 Ed.Física Acomp = Movimento Atual + Testes práticos + Composição Medidas
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order")
SELECT gen_random_uuid(), '11111111-1111-7111-8111-111111111113', a.id, ROW_NUMBER() OVER (ORDER BY a.g_ord,a.sg_ord,a.i_ord)
FROM ai a WHERE (a.grupo='Movimento e atividade física' AND a.subgrupo IN ('Atual','Testes práticos'))
            OR (a.grupo='Composição corporal' AND a.subgrupo='Medidas Objetivas');

-- 6) Verificação -------------------------------------------------------------
SELECT t.name, t.area, count(i.id) AS itens
FROM anamnesis_templates t
LEFT JOIN anamnesis_template_items i ON i.anamnesis_template_id=t.id
WHERE t.id::text LIKE '11111111-1111-7111-8111-1111111111%'
GROUP BY t.id, t.name, t.area ORDER BY t.name;

-- placeholders aposentados?
SELECT name, deleted_at IS NOT NULL AS aposentado FROM anamnesis_templates WHERE name IN ('Teste','Medica Inicial');

ROLLBACK;
