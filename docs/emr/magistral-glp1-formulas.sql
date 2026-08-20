-- As 7 fórmulas do material de suporte a análogos de GLP-1.
--
-- Cada uma vem com o objetivo clínico que o material declara. last_review NULO: são sugestões do
-- fornecedor, não conduta conferida. Idempotente pelo nome.

BEGIN;

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'GLP-1 · constipação', 
    'Estimular o peristaltismo, hidratar o bolo fecal e modular a microbiota. Os análogos de GLP-1 retardam o esvaziamento gástrico e a motilidade, e a constipação é o efeito mais persistente.',
    E'constipação em uso de análogo de GLP-1\nretardo de esvaziamento gástrico\ntomar com bastante líquido',
    'sachê', 'internal', 'oral', 'Veículo para sachê qsp', 30::numeric, 'sachês',
    '1 sachê ao dia, diluído em água, preferencialmente pela manhã', 30,
    'Material de suporte a análogos de GLP-1 (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'GLP-1 · constipação' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', c.n, c.e FROM nova, (VALUES
  (0,'Psyllium',1::numeric,'g','',false),
  (1,'Magnésio quelato',200::numeric,'mg','Dose do elemento, como o material escreve (faixa citada de 150 a 600 mg).',true),
  (2,'Bifidobacterium lactis',5::numeric,'bilhões UFC','',false),
  (3,'Motility',150::numeric,'mg','Blend do fornecedor; a única substância do material sem referência externa.',false)
) AS c(ord,s,q,u,n,e);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'GLP-1 · diarreia',
    'Equilibrar a diversidade da microbiota, reduzir inflamação e restaurar a mucosa intestinal na diarreia funcional que acompanha o uso de análogos de GLP-1.',
    E'diarreia em uso de análogo de GLP-1\nalteração de microbiota e permeabilidade',
    'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 30::numeric, 'cápsulas',
    '1 dose ao dia, em jejum', 30,
    'Material de suporte a análogos de GLP-1 (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'GLP-1 · diarreia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', '', false FROM nova, (VALUES
  (0,'Saccharomyces boulardii',150::numeric,'mg'),
  (1,'Bacillus clausii',2::numeric,'bilhões UFC'),
  (2,'Lactobacillus rhamnosus',2::numeric,'bilhões UFC'),
  (3,'Fibregum B',200::numeric,'mg')
) AS c(ord,s,q,u);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'GLP-1 · digestão e gases',
    'Reduzir formação de gases, otimizar a digestão e aliviar o desconforto abdominal. A digestão incompleta e a fermentação aumentada vêm da menor motilidade.',
    E'distensão abdominal e flatulência\ndigestão incompleta por menor motilidade',
    'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 90::numeric, 'cápsulas',
    '1 dose após as refeições principais', 30,
    'O carvão ativado adsorve de forma inespecífica: afastar pelo menos duas horas de qualquer medicamento e das outras fórmulas.',
    'Material de suporte a análogos de GLP-1 (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'GLP-1 · digestão e gases' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', c.n, false FROM nova, (VALUES
  (0,'Simeticone',80::numeric,'mg',''),
  (1,'Carvão ativado',200::numeric,'mg','Adsorve fármaco e nutriente: separar por duas horas do resto.'),
  (2,'Alfa-amilase',30::numeric,'mg',''),
  (3,'Protease',100::numeric,'mg',''),
  (4,'Lipase',300::numeric,'mg','')
) AS c(ord,s,q,u,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'GLP-1 · saciedade e manutenção do peso',
    'Estímulo endógeno de GLP-1 e PYY para manter a regulação do apetite e a perda de peso alcançada.',
    E'manutenção do peso após o análogo\nsaciedade por estímulo de receptores intestinais',
    'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 90::numeric, 'cápsulas',
    '1 dose 30 minutos antes das principais refeições', 30,
    'Material de suporte a análogos de GLP-1 (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'GLP-1 · saciedade e manutenção do peso' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', '', false FROM nova, (VALUES
  (0,'Berberina',300::numeric,'mg'),
  (1,'Akkermansia muciniphila',50::numeric,'mg'),
  (2,'Slendesta',100::numeric,'mg')
) AS c(ord,s,q,u);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'GLP-1 · suporte dérmico pós-emagrecimento',
    'Estimular a biossíntese de colágeno e elastina e restaurar a firmeza cutânea. A perda rápida de gordura subcutânea leva à flacidez, inclusive facial.',
    E'flacidez após perda rápida de peso\nperda de elasticidade cutânea',
    'sachê', 'internal', 'oral', 'Veículo para sachê qsp', 30::numeric, 'sachês',
    '1 sachê ao dia, diluído em 100 mL de água pela manhã', 90,
    'Material de suporte a análogos de GLP-1 (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'GLP-1 · suporte dérmico pós-emagrecimento' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', c.n, c.e FROM nova, (VALUES
  (0,'Verisol',2.5::numeric,'g','',false),
  (1,'Palmitato de ascorbila',200::numeric,'mg','Material pede vitamina C 200 mg; a dose é do ativo e o insumo sai pelo fator de correção.',true),
  (2,'Nutricolin',100::numeric,'mg','',false)
) AS c(ord,s,q,u,n,e);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'GLP-1 · preservação de massa magra',
    'Preservar massa muscular e estimular a síntese proteica durante a perda rápida de peso, sobretudo em quem tem baixa ingestão proteica, sedentarismo ou idade avançada.',
    E'perda de massa magra durante o análogo\nbaixa ingestão proteica ou idade avançada',
    'sachê', 'internal', 'oral', 'Veículo para sachê qsp', 30::numeric, 'sachês',
    '1 sachê ao dia, diluído em 100 mL de líquido, preferencialmente após o treino', 90,
    'Material de suporte a análogos de GLP-1 (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'GLP-1 · preservação de massa magra' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', c.n, false FROM nova, (VALUES
  (0,'PeptiStrong',1.2::numeric,'g','O material dá 2,4 g/dia como dose usual: esta fórmula entrega metade em um sachê.'),
  (1,'HMB',1.5::numeric,'g','')
) AS c(ord,s,q,u,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'GLP-1 · ciclo capilar',
    'Reduzir eflúvio, regenerar a matriz folicular e dar resistência ao fio. A queda parece vir de deficiência nutricional, inflamação e estresse metabólico da perda rápida.',
    E'eflúvio após perda rápida de peso\nfragilidade capilar',
    'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 30::numeric, 'cápsulas',
    '1 dose ao dia', 90,
    'Contém biotina em dose alta: suspender 3 dias antes de qualquer coleta de sangue e avisar o laboratório. Acima de 5 mg/dia a biotina falseia TSH, T4 livre, troponina e hormônios em imunoensaio.',
    'Material de suporte a análogos de GLP-1 (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'GLP-1 · ciclo capilar' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', c.n, c.e FROM nova, (VALUES
  (0,'Cistina',100::numeric,'mg','',false),
  (1,'Metionina',100::numeric,'mg','',false),
  (2,'Nutricolin',100::numeric,'mg','',false),
  (3,'Biotina',10::numeric,'mg','Dose do material. Interfere em imunoensaio acima de 5 mg/dia: suspender antes da coleta.',false),
  (4,'Ferro',10::numeric,'mg','Dose do elemento. Repor ferro sem ferritina medida é decisão a conferir.',true),
  (5,'Cisteína',100::numeric,'mg','',false),
  (6,'Ácido pantotênico',100::numeric,'mg','',false),
  (7,'Saw palmetto',150::numeric,'mg','',false)
) AS c(ord,s,q,u,n,e);

COMMIT;
