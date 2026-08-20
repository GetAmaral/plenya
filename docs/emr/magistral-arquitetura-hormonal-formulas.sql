-- As três fórmulas do material de arquitetura hormonal.
BEGIN;

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Eixo androgênico · manutenção',
    'Equilíbrio de testosterona total e livre, regulação de SHBG e redução de cortisol, com suporte de zinco, magnésio e boro.',
    E'suporte androgênico fisiológico\ndisposição e desempenho',
    'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 30::numeric, 'cápsulas',
    '1 dose pela manhã, diariamente', 90,
    'Fitoterápico não é reposição: a expectativa a combinar com o paciente é de suporte, não de elevação de testosterona comparável a hormônio.',
    'Material de arquitetura hormonal (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Eixo androgênico · manutenção' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', c.n, c.e FROM nova, (VALUES
  (0,'Testofen',50::numeric,'mg','Ensaios do ingrediente usam 300 a 600 mg/dia: a dose da fórmula fica bem abaixo.',false),
  (1,'Tribulus terrestris',300::numeric,'mg','Revisão sistemática de 2025 não sustenta elevação de testosterona.',false),
  (2,'Eurycoma longifolia',200::numeric,'mg','',false),
  (3,'Boro',3::numeric,'mg','Dose do elemento.',true),
  (4,'Zinco quelato',20::numeric,'mg','Dose do elemento.',true),
  (5,'Magnésio quelato',50::numeric,'mg','Dose do elemento.',true)
) AS c(ord,s,q,u,n,e);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Eixo adrenal · cortisol e energia',
    'Equilíbrio do eixo hipotálamo-hipófise-adrenal com adaptógenos, para regulação de cortisol, disposição e recuperação.',
    E'estresse crônico\ncortisol alto com fadiga',
    'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 30::numeric, 'cápsulas',
    '1 dose ao dia, preferencialmente à tarde', 90,
    'Material de arquitetura hormonal (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Eixo adrenal · cortisol e energia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', c.n, false FROM nova, (VALUES
  (0,'Robuvit',120::numeric,'mg','Ensaios do ingrediente usam 200 a 300 mg/dia.'),
  (1,'Rhodiola rosea',150::numeric,'mg',''),
  (2,'Maca peruana',300::numeric,'mg',''),
  (3,'Teacrina',75::numeric,'mg',''),
  (4,'Ashwagandha',300::numeric,'mg','KSM-66 no material.')
) AS c(ord,s,q,u,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Eixo DHEA · função adrenal',
    'Estímulo fisiológico da produção de DHEA e suporte adrenal, com ginseng, yam mexicano e cofatores.',
    E'queda de DHEA\nfadiga e baixa capacidade adaptativa',
    'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 30::numeric, 'cápsulas',
    '1 dose pela manhã, diariamente', 90,
    'Material de arquitetura hormonal (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Eixo DHEA · função adrenal' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', c.n, c.e FROM nova, (VALUES
  (0,'Panax ginseng',150::numeric,'mg','',false),
  (1,'UbiQsome',100::numeric,'mg','',false),
  (2,'Magnésio quelato',150::numeric,'mg','Dose do elemento.',true),
  (3,'Turkesterone',300::numeric,'mg','Dados humanos escassos; ver a observação da substância.',false),
  (4,'Yam mexicano',200::numeric,'mg','',false),
  (5,'Boro',3::numeric,'mg','Dose do elemento.',true)
) AS c(ord,s,q,u,n,e);

COMMIT;
