-- As fórmulas do formulário das parceiras. Parser conferido contra o sumário do documento;
-- nomes canonizados contra o catálogo; formas preferidas do prescritor aplicadas depois, em
-- magistral-formulario-correcoes.sql. last_review NULO: nenhuma foi conferida.
BEGIN;

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Formula da beleza 30+ anos', 'Fórmula do formulário das parceiras, seção Beleza.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Formula da beleza 30+ anos' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Silício orgânico', 100.0::numeric, 'mg'),
    (1, 'Bio Arct', 100.0::numeric, 'mg'),
    (2, 'Extrato de semente de uva', 200.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Formula da beleza 40+ anos', 'Fórmula do formulário das parceiras, seção Beleza.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Formula da beleza 40+ anos' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Silício orgânico', 100.0::numeric, 'mg'),
    (1, 'Bio Arct', 100.0::numeric, 'mg'),
    (2, 'Glycoxil', 50.0::numeric, 'mg'),
    (3, 'Green coffee', 30.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Formula da beleza 50+ anos', 'Fórmula do formulário das parceiras, seção Beleza.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Tomar 1 sache ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Formula da beleza 50+ anos' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Verisol', 2.5::numeric, 'g'),
    (1, 'Bio Arct', 100.0::numeric, 'mg'),
    (2, 'Silício orgânico', 150.0::numeric, 'mg'),
    (3, 'Green coffee', 30.0::numeric, 'mg'),
    (4, 'Glycoxil', 50.0::numeric, 'mg'),
    (5, 'Licopeno', 10.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Peeling oral', 'Fórmula do formulário das parceiras, seção Beleza.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Peeling oral' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Oli-Ola', 300.0::numeric, 'mg'),
    (1, 'Nutricolin', 100.0::numeric, 'mg'),
    (2, 'Vitamina C', 100.0::numeric, 'mg'),
    (3, 'Fosfolipídeos de caviar', 100.0::numeric, 'mg'),
    (4, 'Licopeno', 10.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Celulite I', 'Fórmula do formulário das parceiras, seção Celulite.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose ao dia, após café da manhã', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Celulite I' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Silício orgânico', 100.0::numeric, 'mg'),
    (1, 'Bio Arct', 100.0::numeric, 'mg'),
    (2, 'DMAE bitartarato', 100.0::numeric, 'mg'),
    (3, 'Castanha-da-índia', 150.0::numeric, 'mg'),
    (4, 'Centella asiatica', 150.0::numeric, 'mg'),
    (5, 'Rutina', 50.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Celulite II', 'Fórmula do formulário das parceiras, seção Celulite.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose pela manhã após o café', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Celulite II' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Dimpless', 40.0::numeric, 'mg'),
    (1, 'Ásiaticosídeo', 30.0::numeric, 'mg'),
    (2, 'Ácido alfa-lipoico', 200.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Aumento do metabolismo', 'Fórmula do formulário das parceiras, seção Emagrecimento.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Tomar 1 cápsula ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Aumento do metabolismo' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Glisodim', 100.0::numeric, 'mg'),
    (1, 'Extrato de chá verde', 75.0::numeric, 'mg'),
    (2, 'Lactobacillus plantarum', 1.0::numeric, 'bilhões UFC')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Compulsao por doces e carboidratos', 'Fórmula do formulário das parceiras, seção Emagrecimento.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula pela manhã e à noite', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Compulsao por doces e carboidratos' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Açafrão padronizado', 90.0::numeric, 'mg'),
    (1, 'Griffonia', 50.0::numeric, 'mg'),
    (2, 'Magnésio quelato', 50.0::numeric, 'mg'),
    (3, 'Rhodiola rosea', 100.0::numeric, 'mg'),
    (4, 'L-teanina', 100.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Redução de medidas', 'Fórmula do formulário das parceiras, seção Emagrecimento.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula pela manhã', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Redução de medidas' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Morosil', 400.0::numeric, 'mg'),
    (1, 'Ácido alfa-lipoico', 75.0::numeric, 'mg'),
    (2, 'Cactinea', 1.0::numeric, 'g'),
    (3, 'Vitamina C', 100.0::numeric, 'mg'),
    (4, 'Magnésio quelato', 75.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Redução gordura, apetite e ação diurética', 'Fórmula do formulário das parceiras, seção Emagrecimento.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 1h antes do almoço e jantar', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Redução gordura, apetite e ação diurética' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Koubo', 200.0::numeric, 'mg'),
    (1, 'Morosil', 250.0::numeric, 'mg'),
    (2, 'Açafrão padronizado', 90.0::numeric, 'mg'),
    (3, 'Cactinea', 500.0::numeric, 'mg'),
    (4, 'Hibisco', 150.0::numeric, 'mg'),
    (5, 'Abacateiro', 250.0::numeric, 'mg'),
    (6, 'Cavalinha', 250.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Termogênica acelera metabolismo', 'Fórmula do formulário das parceiras, seção Emagrecimento.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 1h antes do almoço e jantar', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Termogênica acelera metabolismo' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Picolinato de cromo', 250.0::numeric, 'mcg'),
    (1, 'Teacrina', 50.0::numeric, 'mg'),
    (2, 'Ácido hidroxicítrico', 150.0::numeric, 'mg'),
    (3, 'Gymnema silvestre', 100.0::numeric, 'mg'),
    (4, 'Extrato de chá verde', 250.0::numeric, 'mg'),
    (5, 'Cactinea', 500.0::numeric, 'mg'),
    (6, 'Gengibre', 100.0::numeric, 'mg'),
    (7, 'Capsiate', 3.0::numeric, 'mg'),
    (8, 'Cássia angustifólia', 150.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Termogênica queima gordura', 'Fórmula do formulário das parceiras, seção Emagrecimento.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 1h antes do almoço e jantar', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Termogênica queima gordura' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'ID-alG', 100.0::numeric, 'mg'),
    (1, 'Meratrim', 200.0::numeric, 'mg'),
    (2, 'Lowat', 100.0::numeric, 'mg'),
    (3, 'Ácido hidroxicítrico', 200.0::numeric, 'mg'),
    (4, 'Gymnema silvestre', 100.0::numeric, 'mg'),
    (5, 'Green coffee', 100.0::numeric, 'mg'),
    (6, 'Cavalinha', 500.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Queima gordura e ganho massa', 'Fórmula do formulário das parceiras, seção Emagrecimento.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 dose 2 vezes ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Queima gordura e ganho massa' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Coleus forskohlii', 300.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Gripes, resfriados e herpes', 'Fórmula do formulário das parceiras, seção Imunomodulação.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 2 vezes ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Gripes, resfriados e herpes' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Equinácea', 150.0::numeric, 'mg'),
    (1, 'L-lisina', 300.0::numeric, 'mg'),
    (2, 'Vitamina C', 100.0::numeric, 'mg'),
    (3, 'Chlorella', 100.0::numeric, 'mg'),
    (4, 'Zinco quelato', 10.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Imunomoduladora', 'Fórmula do formulário das parceiras, seção Imunomodulação.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Imunomoduladora' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Epicor', 50.0::numeric, 'mg'),
    (1, 'L-lisina', 300.0::numeric, 'mg'),
    (2, 'Vitamina C', 100.0::numeric, 'mg'),
    (3, 'Romã', 100.0::numeric, 'mg'),
    (4, 'Zinco quelato', 15.0::numeric, 'mg'),
    (5, 'Extrato de semente de uva', 50.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Fortalecimento sistema imunológico', 'Fórmula do formulário das parceiras, seção Imunomodulação.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 2 vezes ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Fortalecimento sistema imunológico' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Astragalus', 100.0::numeric, 'mg'),
    (1, 'Curcumina', 200.0::numeric, 'mg'),
    (2, 'Cyanotis vagas', 200.0::numeric, 'mg'),
    (3, 'Extrato de chá verde', 100.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Detox', 'Fórmula do formulário das parceiras, seção Nutrição esportiva.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose pela manhã após o café', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Detox' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Chlorella', 250.0::numeric, 'mg'),
    (1, 'Folha de oliveira', 200.0::numeric, 'mg'),
    (2, 'Gengibre', 100.0::numeric, 'mg'),
    (3, 'Green coffee', 50.0::numeric, 'mg'),
    (4, 'Altilix', 50.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Controle do apetite', 'Fórmula do formulário das parceiras, seção Nutrição esportiva.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 dose 1h antes do almoço e jantar', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Controle do apetite' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Koubo', 200.0::numeric, 'mg'),
    (1, 'Griffonia', 50.0::numeric, 'mg'),
    (2, 'Citrus aurantium', 400.0::numeric, 'mg'),
    (3, 'Ácido hidroxicítrico', 300.0::numeric, 'mg'),
    (4, 'Gymnema silvestre', 75.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Termogenico – queima gordura', 'Fórmula do formulário das parceiras, seção Nutrição esportiva.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 dose 1h antes do almoço e jantar', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Termogenico – queima gordura' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Picolinato de cromo', 250.0::numeric, 'mcg'),
    (1, 'Sinetrol', 500.0::numeric, 'mg'),
    (2, 'Citrus aurantium', 200.0::numeric, 'mg'),
    (3, 'Morosil', 250.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pré treino i– fornecedor de energia', 'Fórmula do formulário das parceiras, seção Nutrição esportiva.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 a 2 cáps antes do treino', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pré treino i– fornecedor de energia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Teacrina', 100.0::numeric, 'mg'),
    (1, 'L-taurina', 200.0::numeric, 'mg'),
    (2, 'AAKG', 500.0::numeric, 'mg'),
    (3, 'L-citrulina malato', 300.0::numeric, 'mg'),
    (4, 'Beta-alanina', 500.0::numeric, 'mg'),
    (5, 'Piridoxal-5-fosfato', 5.0::numeric, 'mg'),
    (6, 'Selenometionina', 40.0::numeric, 'mcg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pré treino II – força e explosão', 'Fórmula do formulário das parceiras, seção Nutrição esportiva.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Dissolver 1 sache em líquido e tomar 30min antes do treino', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pré treino II – força e explosão' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'BCAA', 5.0::numeric, 'g'),
    (1, 'L-tirosina', 400.0::numeric, 'mg'),
    (2, 'Piperina', 5.0::numeric, 'mg'),
    (3, 'L-arginina', 2000.0::numeric, 'mg'),
    (4, 'Beta-alanina', 1000.0::numeric, 'mg'),
    (5, 'HMB', 1000.0::numeric, 'mg'),
    (6, 'D-ribose', 2.0::numeric, 'g'),
    (7, 'Palatinose', 10.0::numeric, 'g')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pós treino - recuperação muscular', 'Fórmula do formulário das parceiras, seção Nutrição esportiva.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Dissolver 1 sache em líquido e tomar após o treino', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pós treino - recuperação muscular' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'BCAA', 5.0::numeric, 'g'),
    (1, 'Phosfator', 500.0::numeric, 'mg'),
    (2, 'Gama-orizanol', 300.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pos treino – aumento de testosterona', 'Fórmula do formulário das parceiras, seção Nutrição esportiva.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose ào dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pos treino – aumento de testosterona' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Ácido D-aspártico', 500.0::numeric, 'mg'),
    (1, 'Tribulus terrestris', 500.0::numeric, 'mg'),
    (2, 'Mucuna pruriens', 200.0::numeric, 'mg'),
    (3, 'Maca peruana', 200.0::numeric, 'mg'),
    (4, 'Pygeum africanum', 100.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Equilibrio da flora - biodisponibilidade de nutrientes', 'Fórmula do formulário das parceiras, seção Prebiótiocos e próbioticos.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Tomar 1 dose ao deitar, não ingerir com alimentos quentes', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Equilibrio da flora - biodisponibilidade de nutrientes' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Lactobacillus acidophilus', 1.0::numeric, 'bilhões UFC'),
    (1, 'Bifidobacterium bifidum', 1.0::numeric, 'bilhões UFC'),
    (2, 'Lactobacillus bulgaricus', 1.0::numeric, 'bilhões UFC'),
    (3, 'Lactobacillus rhamnosus', 1.0::numeric, 'bilhões UFC'),
    (4, 'Lactobacillus plantarum', 1.0::numeric, 'bilhões UFC'),
    (5, 'Lactobacillus salivarius', 1.0::numeric, 'bilhões UFC'),
    (6, 'Inulina', 200.0::numeric, 'MG')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Intolerância a lactose-melhora digestão e reduz sintomas', 'Fórmula do formulário das parceiras, seção Prebiótiocos e próbioticos.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Tomar 1 dose ao deitar, não ingerir com alimentos quentes', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Intolerância a lactose-melhora digestão e reduz sintomas' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Lactobacillus delbrueckii', 4.0::numeric, 'bilhões UFC'),
    (1, 'Lactobacillus bulgaricus', 2.0::numeric, 'bilhões UFC'),
    (2, 'Streptococcus thermophilus', 2.0::numeric, 'bilhões UFC'),
    (3, 'Inulina', 200.0::numeric, 'MG')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Gerenciamento de peso', 'Fórmula do formulário das parceiras, seção Prebiótiocos e próbioticos.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Tomar 1 dose ao deitar, não ingerir com alimentos quentes', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Gerenciamento de peso' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Lactobacillus gasseri', 1.0::numeric, 'bilhões UFC'),
    (1, 'Lactobacillus reuteri', 1.0::numeric, 'bilhões UFC'),
    (2, 'Lactobacillus casei', 1.0::numeric, 'bilhões UFC'),
    (3, 'Bifidobacterium breve', 1.0::numeric, 'bilhões UFC'),
    (4, 'Bifidobacterium longum', 1.0::numeric, 'bilhões UFC'),
    (5, 'Inulina', 200.0::numeric, 'MG')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Restaura mucosa intestinal', 'Fórmula do formulário das parceiras, seção Prebiótiocos e próbioticos.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         7::numeric, 'sachês', 'Dissolver 1 sache em liquido e por 7 dias', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Restaura mucosa intestinal' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'L-glutamina', 5.0::numeric, 'G')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Melhora constipação intestinal infantil', 'Fórmula do formulário das parceiras, seção Prebiótiocos e próbioticos.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Dissolver 1 sache em liquido e tomar diariamente', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Melhora constipação intestinal infantil' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Peg 4000 5g a', 10.0::numeric, 'G')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Melhora constipação intestinal adulto gestantes e idosos', 'Fórmula do formulário das parceiras, seção Prebiótiocos e próbioticos.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Dissolver 1 a 2 colheres de sopa em liquido e tomar diariamente', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Melhora constipação intestinal adulto gestantes e idosos' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Peg 4000 10g a', 20.0::numeric, 'g')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Anemia gestacional', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula após o almoço', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Anemia gestacional' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Metilfolato', 100.0::numeric, 'mcg'),
    (1, 'Cobre', 1.0::numeric, 'mg'),
    (2, 'Ferro', 30.0::numeric, 'mg'),
    (3, 'Piridoxal-5-fosfato', 50.0::numeric, 'mg'),
    (4, 'Metilcobalamina', 50.0::numeric, 'mcg'),
    (5, 'Vitamina C', 100.0::numeric, 'mg'),
    (6, 'Vitamina E', 60.0::numeric, 'ui'),
    (7, 'Zinco quelato', 15.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Antidiabetes funcional', 'Fórmula do formulário das parceiras, seção Patologias.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Tomar 1 cápsula ao dia, não ingerir com líquidos', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Antidiabetes funcional' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Weg lem 70(ganoderma lucidum)', 200.0::numeric, 'mg'),
    (1, 'Glisodim', 100.0::numeric, 'mg'),
    (2, 'Bifidobacterium lactis', 1.0::numeric, 'bilhões UFC'),
    (3, 'Streptococcus thermophilus', 1.0::numeric, 'bilhões UFC'),
    (4, 'Lactobacillus delbrueckii', 1.0::numeric, 'bilhões UFC')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Antiobesidade', 'Fórmula do formulário das parceiras, seção Patologias.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Tomar 1 cápsula ao dia, não ingerir com líquidos', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Antiobesidade' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Weg lem 70(ganoderma lucidum)', 200.0::numeric, 'mg'),
    (1, 'Glisodim', 200.0::numeric, 'mg'),
    (2, 'Lactobacillus gasseri', 1.0::numeric, 'bilhões UFC'),
    (3, 'Lactobacillus paracasei', 1.0::numeric, 'bilhões UFC')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Anorexia', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Anorexia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Metilfolato', 100.0::numeric, 'mcg'),
    (1, 'Betaína anidra', 50.0::numeric, 'mg'),
    (2, 'Cianocobalamina', 100.0::numeric, 'mcg'),
    (3, 'Cobre', 0.5::numeric, 'mg'),
    (4, 'L-lisina', 100.0::numeric, 'mg'),
    (5, 'Magnésio quelato', 100.0::numeric, 'mg'),
    (6, 'Nicotinamida', 5.0::numeric, 'mg'),
    (7, 'Vitamina B1', 5.0::numeric, 'mg'),
    (8, 'Riboflavina', 5.0::numeric, 'mg'),
    (9, 'Vitamina C', 100.0::numeric, 'mg'),
    (10, 'Zinco quelato', 8.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Anorexia e inapetência II', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula antes de dormir', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Anorexia e inapetência II' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'L-carnitina', 200.0::numeric, 'mg'),
    (1, 'Buclizina', 50.0::numeric, 'mg'),
    (2, 'Ciproeptadina', 6.0::numeric, 'mg'),
    (3, 'Metilcobalamina', 0.5::numeric, 'mg'),
    (4, 'L-lisina', 100.0::numeric, 'mg'),
    (5, 'Magnésio quelato', 100.0::numeric, 'mg'),
    (6, 'Metilfolato', 10.0::numeric, 'mcg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Ansiedade generalizada', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula de 12/12 horas', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Ansiedade generalizada' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, '5-HTP', 25.0::numeric, 'mg'),
    (1, 'Ashwagandha', 200.0::numeric, 'mg'),
    (2, 'Kava-kava', 50.0::numeric, 'mg'),
    (3, 'Cálcio', 25.0::numeric, 'mg'),
    (4, 'Valeriana', 25.0::numeric, 'mg'),
    (5, 'L-glutamina', 100.0::numeric, 'mg'),
    (6, 'L-taurina', 75.0::numeric, 'mg'),
    (7, 'Magnésio quelato', 100.0::numeric, 'mg'),
    (8, 'Melissa', 100.0::numeric, 'mg'),
    (9, 'L-teanina', 100.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Antioxidante básico', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula de 12/12 horas', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Antioxidante básico' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Betacaroteno', 50.0::numeric, 'mg'),
    (1, 'Cobre', 1.0::numeric, 'mg'),
    (2, 'Manganês', 1.0::numeric, 'mg'),
    (3, 'Selenometionina', 30.0::numeric, 'mcg'),
    (4, 'Vitamina C', 100.0::numeric, 'mg'),
    (5, 'Vitamina E', 30.0::numeric, 'ui'),
    (6, 'Zinco quelato', 15.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Antioxidante maxi', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Antioxidante maxi' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Ácido alfa-lipoico', 50.0::numeric, 'mg'),
    (1, 'Coenzima Q10', 40.0::numeric, 'mg'),
    (2, 'Cobre', 1.0::numeric, 'mg'),
    (3, 'Zinco quelato', 15.0::numeric, 'mg'),
    (4, 'Cisteína', 100.0::numeric, 'mg'),
    (5, 'Manganês', 1.0::numeric, 'mg'),
    (6, 'NADH', 1.0::numeric, 'mg'),
    (7, 'Picnogenol', 35.0::numeric, 'mg'),
    (8, 'Trans-resveratrol', 5.0::numeric, 'mg'),
    (9, 'Selênio', 30.0::numeric, 'mcg'),
    (10, 'Dimpless', 40.0::numeric, 'mg'),
    (11, 'Riboflavina', 10.0::numeric, 'mg'),
    (12, 'Vitamina C', 100.0::numeric, 'mg'),
    (13, 'Vitamina E', 45.0::numeric, 'ui')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Anti-ox maxi ultra', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Anti-ox maxi ultra' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Ácido alfa-lipoico', 50.0::numeric, 'mg'),
    (1, 'Cobre', 1.0::numeric, 'mg'),
    (2, 'Coenzima Q10', 40.0::numeric, 'mg'),
    (3, 'Cisteína', 100.0::numeric, 'mg'),
    (4, 'Licopeno', 5.0::numeric, 'mg'),
    (5, 'Luteína', 5.0::numeric, 'mg'),
    (6, 'Manganês', 1.0::numeric, 'mg'),
    (7, 'NADH', 1.0::numeric, 'mg'),
    (8, 'Picnogenol', 20.0::numeric, 'mg'),
    (9, 'Trans-resveratrol', 5.0::numeric, 'mg'),
    (10, 'PQQ', 10.0::numeric, 'mg'),
    (11, 'Selênio', 30.0::numeric, 'mg'),
    (12, 'Riboflavina', 10.0::numeric, 'mg'),
    (13, 'Vitamina C', 100.0::numeric, 'mg'),
    (14, 'Vitamina E', 45.0::numeric, 'ui'),
    (15, 'Zinco quelato', 15.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Antioxidantes para fumantes', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 2x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Antioxidantes para fumantes' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Cobre', 0.5::numeric, 'mg'),
    (1, 'Cisteína', 100.0::numeric, 'mg'),
    (2, 'Luteína', 1.0::numeric, 'mg'),
    (3, 'Nicotinamida', 50.0::numeric, 'mg'),
    (4, 'Selênio', 30.0::numeric, 'mcg'),
    (5, 'L-taurina', 75.0::numeric, 'mg'),
    (6, 'Vitamina C', 100.0::numeric, 'mg'),
    (7, 'Vitamina E', 90.0::numeric, 'ui'),
    (8, 'Zinco quelato', 8.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Artrite artrose I', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose via oral ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Artrite artrose I' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Colágeno tipo II', 40.0::numeric, 'mg'),
    (1, 'Boswellia', 500.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Artrite artrose I I', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 dose via oral 2 vezes ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Artrite artrose I I' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Sucupira', 500.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Cãimbras', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Cãimbras' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Ácido alfa-lipoico', 50.0::numeric, 'mg'),
    (1, 'Potássio', 5.0::numeric, 'mg'),
    (2, 'Coenzima Q10', 30.0::numeric, 'mg'),
    (3, 'Creatina', 250.0::numeric, 'mg'),
    (4, 'Magnésio quelato', 50.0::numeric, 'mg'),
    (5, 'MSM', 100.0::numeric, 'mg'),
    (6, 'Vitamina A', 50.0::numeric, 'mg'),
    (7, 'Piridoxal-5-fosfato', 10.0::numeric, 'mg'),
    (8, 'Vitamina C', 100.0::numeric, 'mg'),
    (9, 'Vitamina E', 90.0::numeric, 'ui'),
    (10, 'Vitamina K2 MK-7', 25.0::numeric, 'mcg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Circulacao periférica - varizes , varicoses e hemorroidas', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 dose 2 vezes ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Circulacao periférica - varizes , varicoses e hemorroidas' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Diosmina', 250.0::numeric, 'mg'),
    (1, 'Hesperidina', 100.0::numeric, 'mg'),
    (2, 'Castanha-da-índia', 150.0::numeric, 'mg'),
    (3, 'Centella asiatica', 150.0::numeric, 'mg'),
    (4, 'Rutina', 100.0::numeric, 'mg'),
    (5, 'Ásiaticosídeo', 20.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Cistite infecções urinárias', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 dose via oral 2 vezes ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Cistite infecções urinárias' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Cranberry', 500.0::numeric, 'mg'),
    (1, 'Vitamina C', 250.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Colesterol I', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose à noite', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Colesterol I' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Policosanol', 30.0::numeric, 'mg'),
    (1, 'Cissus quadrangularis', 250.0::numeric, 'mg'),
    (2, 'Trans-resveratrol', 30.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Colesterol II', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose à noite', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Colesterol II' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Policosanol', 30.0::numeric, 'mg'),
    (1, 'Trans-resveratrol', 30.0::numeric, 'mg'),
    (2, 'Ácido alfa-lipoico', 150.0::numeric, 'mg'),
    (3, 'Coenzima Q10', 50.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Cortisol elevado – modulação estresse', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cáps 2x/dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Cortisol elevado – modulação estresse' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Relora', 150.0::numeric, 'mg'),
    (1, 'L-teanina', 200.0::numeric, 'mg'),
    (2, 'Ashwagandha', 200.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Diabetes mellitus', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Diabetes mellitus' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Ácido alfa-lipoico', 10.0::numeric, 'mg'),
    (1, 'Beta-alanina', 500.0::numeric, 'mg'),
    (2, 'Potássio', 100.0::numeric, 'mg'),
    (3, 'Cobre', 1.0::numeric, 'mg'),
    (4, 'Coenzima Q10', 25.0::numeric, 'mg'),
    (5, 'Picolinato de cromo', 25.0::numeric, 'mg'),
    (6, 'Cisteína', 50.0::numeric, 'mg'),
    (7, 'Licopeno', 3.0::numeric, 'mg'),
    (8, 'Magnésio quelato', 100.0::numeric, 'mg'),
    (9, 'Selenometionina', 30.0::numeric, 'mcg'),
    (10, 'Vanádio', 50.0::numeric, 'mcg'),
    (11, 'Piridoxal-5-fosfato', 50.0::numeric, 'mg'),
    (12, 'Metilcobalamina', 100.0::numeric, 'mcg'),
    (13, 'Vitamina C', 100.0::numeric, 'mg'),
    (14, 'Vitamina E', 90.0::numeric, 'ui'),
    (15, 'Zinco quelato', 15.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Diabetes tipo I', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 2x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Diabetes tipo I' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Ácido alfa-lipoico', 10.0::numeric, 'mg'),
    (1, 'Cobre', 1.0::numeric, 'mg'),
    (2, 'Picolinato de cromo', 50.0::numeric, 'mcg'),
    (3, 'Magnésio quelato', 50.0::numeric, 'mg'),
    (4, 'Manganês', 1.0::numeric, 'mg'),
    (5, 'Potássio', 50.0::numeric, 'mg'),
    (6, 'Selênio', 30.0::numeric, 'mcg'),
    (7, 'Vanádio', 25.0::numeric, 'mcg'),
    (8, 'Vitamina C', 100.0::numeric, 'mg'),
    (9, 'Vitamina E', 50.0::numeric, 'ui'),
    (10, 'Zinco quelato', 15.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Diabetes tipo II', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula antes do almoço e jantar', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Diabetes tipo II' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Ácido alfa-lipoico', 10.0::numeric, 'mg'),
    (1, 'Metilfolato', 10.0::numeric, 'mg'),
    (2, 'Magnésio aspartato', 100.0::numeric, 'mg'),
    (3, 'Faseolamina', 100.0::numeric, 'mg'),
    (4, 'Glutationa reduzida', 50.0::numeric, 'mg'),
    (5, 'L-arginina', 100.0::numeric, 'mg'),
    (6, 'Psyllium', 1.0::numeric, 'g'),
    (7, 'Selênio', 30.0::numeric, 'mcg'),
    (8, 'Vanádio', 100.0::numeric, 'mcg'),
    (9, 'Vitamina C', 100.0::numeric, 'mg'),
    (10, 'Vitamina E', 90.0::numeric, 'ui')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Doença cardiovascular', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Doença cardiovascular' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Nattoquinase', 50.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Doença celíaca (sachês)', 'Fórmula do formulário das parceiras, seção Patologias.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Tomar 1 sachê 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Doença celíaca (sachês)' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Metilfolato', 100.0::numeric, 'mcg'),
    (1, 'Cobre', 1.0::numeric, 'mg'),
    (2, 'Enzimas pancreáticas', 250.0::numeric, 'mg'),
    (3, 'FOS', 250.0::numeric, 'mg'),
    (4, 'L-glutamina', 100.0::numeric, 'mg'),
    (5, 'Lecitina', 100.0::numeric, 'mg'),
    (6, 'Magnésio quelato', 100.0::numeric, 'mg'),
    (7, 'Psyllium', 1.0::numeric, 'g'),
    (8, 'Polidextrose', 1.0::numeric, 'g'),
    (9, 'Vitamina A', 1000.0::numeric, 'ui'),
    (10, 'Piridoxal-5-fosfato', 10.0::numeric, 'mg'),
    (11, 'Vitamina C', 100.0::numeric, 'mg'),
    (12, 'Vitamina D3', 100.0::numeric, 'ui'),
    (13, 'Vitamina E', 100.0::numeric, 'ui'),
    (14, 'Zinco quelato', 100.0::numeric, 'mcg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Energético mineral (catalisadores)', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 2x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Energético mineral (catalisadores)' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Boro', 1.0::numeric, 'mg'),
    (1, 'Cálcio', 100.0::numeric, 'mg'),
    (2, 'Cobre', 1.0::numeric, 'mg'),
    (3, 'Magnésio quelato', 50.0::numeric, 'mg'),
    (4, 'Manganês', 10.0::numeric, 'mg'),
    (5, 'Picolinato de cromo', 25.0::numeric, 'mcg'),
    (6, 'Molibdênio', 25.0::numeric, 'mcg'),
    (7, 'Potássio', 50.0::numeric, 'mg'),
    (8, 'Selenometionina', 30.0::numeric, 'mcg'),
    (9, 'Zinco quelato', 10.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Energético vitamínico', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Energético vitamínico' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Metilfolato', 100.0::numeric, 'mcg'),
    (1, 'Vitamina B1', 1.0::numeric, 'mg'),
    (2, 'Nicotinamida', 10.0::numeric, 'mg'),
    (3, 'Ácido pantotênico', 50.0::numeric, 'mg'),
    (4, 'Piridoxal-5-fosfato', 10.0::numeric, 'mg'),
    (5, 'Metilcobalamina', 100.0::numeric, 'mcg'),
    (6, 'Vitamina C', 100.0::numeric, 'mg'),
    (7, 'Vitamina E', 100.0::numeric, 'ui'),
    (8, 'Vitamina D3', 400.0::numeric, 'ui')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Energizante – combate a fadiga', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose ao dia, após café da manhã', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Energizante – combate a fadiga' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'L-taurina', 200.0::numeric, 'mg'),
    (1, 'Piridoxal-5-fosfato', 7.0::numeric, 'mg'),
    (2, 'Magnésio quelato', 100.0::numeric, 'mg'),
    (3, 'Cálcio', 200.0::numeric, 'mg'),
    (4, 'Zinco quelato', 20.0::numeric, 'mg'),
    (5, 'Piridoxal-5-fosfato', 100.0::numeric, 'mg'),
    (6, 'Griffonia', 100.0::numeric, 'mg'),
    (7, 'Coenzima Q10', 50.0::numeric, 'mg'),
    (8, 'Selenometionina', 50.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Esteatose hepática, resistencia insulinica e dislipidemia', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 2x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Esteatose hepática, resistencia insulinica e dislipidemia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Morosil', 200.0::numeric, 'mg'),
    (1, 'Ácido alfa-lipoico', 75.0::numeric, 'mg'),
    (2, 'Piridoxal-5-fosfato', 3.0::numeric, 'mg'),
    (3, 'Ácido pantotênico', 30.0::numeric, 'mg'),
    (4, 'Vitamina C', 60.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Enzimas digestivas MIX', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 2x ao dia, 30 min antes do almoço e jantar', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Enzimas digestivas MIX' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Alfa-amilase', 100.0::numeric, 'mg'),
    (1, 'Bromelina', 150.0::numeric, 'mg'),
    (2, 'Lactase', 85.0::numeric, 'mg'),
    (3, 'Lipase', 50.0::numeric, 'mg'),
    (4, 'Papaína', 50.0::numeric, 'mg'),
    (5, 'Protease', 120.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Flacidez', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Flacidez' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Cobre', 100.0::numeric, 'mcg'),
    (1, 'L-carnitina', 100.0::numeric, 'mg'),
    (2, 'Glicina', 100.0::numeric, 'mg'),
    (3, 'L-lisina', 100.0::numeric, 'mg'),
    (4, 'L-prolina', 100.0::numeric, 'mg'),
    (5, 'Manganês', 10.0::numeric, 'mg'),
    (6, 'Silício orgânico', 100.0::numeric, 'mg'),
    (7, 'Vitamina C', 150.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Fotoproteção oral', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 dose via oral 2 vezes ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Fotoproteção oral' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Polypodium leucotomos', 250.0::numeric, 'mg'),
    (1, 'Betacaroteno', 15.0::numeric, 'mg'),
    (2, 'Picnogenol', 75.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Hipoglicemia', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Hipoglicemia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Cobre', 1.0::numeric, 'mg'),
    (1, 'Picolinato de cromo', 50.0::numeric, 'mg'),
    (2, 'D-ribose', 250.0::numeric, 'mg'),
    (3, 'Manganês', 1.0::numeric, 'mg'),
    (4, 'Vanádio', 25.0::numeric, 'mg'),
    (5, 'Zinco quelato', 15.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Hiperglicemia', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Hiperglicemia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Ácido alfa-lipoico', 20.0::numeric, 'mg'),
    (1, 'Betaína anidra', 50.0::numeric, 'mg'),
    (2, 'Cobre', 0.5::numeric, 'mg'),
    (3, 'Picolinato de cromo', 50.0::numeric, 'mcg'),
    (4, 'Faseolamina', 100.0::numeric, 'mg'),
    (5, 'FOS', 250.0::numeric, 'mg'),
    (6, 'Gymnema silvestre', 200.0::numeric, 'mg'),
    (7, 'Cisteína', 100.0::numeric, 'mg'),
    (8, 'Manganês', 1.0::numeric, 'mg'),
    (9, 'Nicotinamida', 10.0::numeric, 'mg'),
    (10, 'Potássio', 50.0::numeric, 'mg'),
    (11, 'Vanádio', 50.0::numeric, 'mcg'),
    (12, 'Metilcobalamina', 100.0::numeric, 'mcg'),
    (13, 'Riboflavina', 25.0::numeric, 'mg'),
    (14, 'Piridoxal-5-fosfato', 25.0::numeric, 'mg'),
    (15, 'Vitamina C', 100.0::numeric, 'mg'),
    (16, 'Zinco quelato', 15.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Hipotireoidismo', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Hipotireoidismo' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Betacaroteno', 2.5::numeric, 'mg'),
    (1, 'Cobre', 1.0::numeric, 'mg'),
    (2, 'Iodo', 100.0::numeric, 'mcg'),
    (3, 'L-tirosina', 100.0::numeric, 'mg'),
    (4, 'Selenometionina', 30.0::numeric, 'mcg'),
    (5, 'Vitamina A', 1000.0::numeric, 'ui'),
    (6, 'Vitamina C', 100.0::numeric, 'mg'),
    (7, 'Vitamina D3', 50.0::numeric, 'ui'),
    (8, 'Zinco quelato', 15.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Insonia I', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose via oral antes de dormir', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Insonia I' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Passiflora', 120.0::numeric, 'mg'),
    (1, 'Melissa', 120.0::numeric, 'mg'),
    (2, 'Mulungu', 80.0::numeric, 'mg'),
    (3, 'Valeriana', 50.0::numeric, 'mg'),
    (4, 'Griffonia', 75.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Insonia II', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose via oral antes de dormir', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Insonia II' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Melatonina', 3.0::numeric, 'mg'),
    (1, 'Magnésio quelato', 50.0::numeric, 'mg'),
    (2, 'Piridoxal-5-fosfato', 100.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Menopausa', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 dose via oral 2 vezes ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Menopausa' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Amora', 500.0::numeric, 'mg'),
    (1, 'Cimicifuga racemosa', 100.0::numeric, 'mg'),
    (2, 'Isoflavona', 50.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Menopausa II', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 dose via oral 2 vezes ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Menopausa II' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Isoflavona', 80.0::numeric, 'mg'),
    (1, 'Trevo-vermelho', 40.0::numeric, 'mg'),
    (2, 'Dong quai', 80.0::numeric, 'mg'),
    (3, 'Yam mexicano', 100.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Modulação nutricional em infecções', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Modulação nutricional em infecções' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Equinácea', 25.0::numeric, 'mg'),
    (1, 'Epicor', 250.0::numeric, 'mg'),
    (2, 'Cobre', 1.0::numeric, 'mg'),
    (3, 'Quercetina', 75.0::numeric, 'mg'),
    (4, 'L-taurina', 500.0::numeric, 'mg'),
    (5, 'Vitamina C', 500.0::numeric, 'mg'),
    (6, 'Zinco quelato', 10.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Osteoporose I', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula ao deitar', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Osteoporose I' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Boro', 1.0::numeric, 'mg'),
    (1, 'Cálcio', 500.0::numeric, 'mg'),
    (2, 'Magnésio dimalato', 200.0::numeric, 'mg'),
    (3, 'Ácido pantotênico', 100.0::numeric, 'mg'),
    (4, 'Vitamina K2 MK-7', 50.0::numeric, 'mcg'),
    (5, 'Vitamina D3', 400.0::numeric, 'ui'),
    (6, 'Colágeno tipo II', 40.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Osteoporose II hormonal', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Osteoporose II hormonal' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Boro', 1.0::numeric, 'mg'),
    (1, 'Cálcio', 500.0::numeric, 'mg'),
    (2, 'Isoflavona', 120.0::numeric, 'mg'),
    (3, 'Vitamina D3', 400.0::numeric, 'ui'),
    (4, 'Magnésio dimalato', 200.0::numeric, 'mg'),
    (5, 'Vitamina K2 MK-7', 50.0::numeric, 'mcg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Osteoporose – MIX vit lipolíticas', 'Fórmula do formulário das parceiras, seção Patologias.', 'solução', 'internal', 'oral', 'Veículo hidroalcoólico qsp',
         30::numeric, 'ml', 'Tomar conforme orientação', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Osteoporose – MIX vit lipolíticas' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Vitamina D3', 1000.0::numeric, 'UI'),
    (1, 'Vitamina A', 1000.0::numeric, 'UI'),
    (2, 'Vitamina K2 MK-7', 100.0::numeric, 'mcg'),
    (3, 'Veiculo oleoso qsp', 5.0::numeric, 'Gotas')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Ovário policistico', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 2x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Ovário policistico' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'D-quiro-inositol', 500.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Suplemento geriátrico', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Suplemento geriátrico' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Ácido alfa-lipoico', 20.0::numeric, 'mg'),
    (1, 'Ashwagandha', 200.0::numeric, 'mg'),
    (2, 'Coenzima Q10', 10.0::numeric, 'mg'),
    (3, 'Ginseng', 50.0::numeric, 'mg'),
    (4, 'Fosfatidilserina', 100.0::numeric, 'mg'),
    (5, 'L-carnitina', 250.0::numeric, 'mg'),
    (6, 'L-isoleucina', 100.0::numeric, 'mg'),
    (7, 'L-leucina', 100.0::numeric, 'mg'),
    (8, 'L-valina', 100.0::numeric, 'mg'),
    (9, 'NADH', 10.0::numeric, 'mg'),
    (10, 'Picnogenol', 15.0::numeric, 'mg'),
    (11, 'Pregnenolona', 10.0::numeric, 'mg'),
    (12, 'Trans-resveratrol', 5.0::numeric, 'mg'),
    (13, 'Vinpocetina', 5.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'TPM control', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'TPM control' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Açafrão padronizado', 100.0::numeric, 'mg'),
    (1, 'Griffonia', 50.0::numeric, 'mg'),
    (2, 'Metilfolato', 400.0::numeric, 'mcg'),
    (3, 'Rhodiola rosea', 200.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Tintura anti refluxo', 'Fórmula do formulário das parceiras, seção Patologias.', 'solução', 'internal', 'oral', 'Veículo hidroalcoólico qsp',
         30::numeric, 'ml', 'Tomar 20 gotas em água antes das principais refeições', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Tintura anti refluxo' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Tintura de alecrim', 60.0::numeric, 'mL')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Tintura anti flatulência', 'Fórmula do formulário das parceiras, seção Patologias.', 'solução', 'internal', 'oral', 'Veículo hidroalcoólico qsp',
         30::numeric, 'ml', 'Tomar 20 gotas em água antes das principais refeições', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Tintura anti flatulência' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Tintura de hortelã', 60.0::numeric, '%'),
    (1, 'Tintura de alcachofra', 40.0::numeric, '%')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Tintura proteção gastrica MIX', 'Fórmula do formulário das parceiras, seção Patologias.', 'solução', 'internal', 'oral', 'Veículo hidroalcoólico qsp',
         30::numeric, 'ml', 'Tomar 20 gotas em água antes das principais refeições', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Tintura proteção gastrica MIX' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Tintura de funcho', 30.0::numeric, '%'),
    (1, 'Tintura de espinheira-santa', 30.0::numeric, '%'),
    (2, 'Tintura de alecrim', 40.0::numeric, '%')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Ansiedade', 'Fórmula do formulário das parceiras, seção Formas farmacêuticas.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', '1 dose ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Ansiedade' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Griffonia', 75.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Anti –aging', 'Fórmula do formulário das parceiras, seção Formas farmacêuticas.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', '1 dose ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Anti –aging' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Verisol', 2.5::numeric, 'G')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Compulsão alimentar', 'Fórmula do formulário das parceiras, seção Formas farmacêuticas.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', '1 dose ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Compulsão alimentar' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Picolinato de cromo', 350.0::numeric, 'mcg'),
    (1, 'Açafrão padronizado', 160.0::numeric, 'mg'),
    (2, 'Koubo', 400.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Diurético', 'Fórmula do formulário das parceiras, seção Formas farmacêuticas.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', '1 dose ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Diurético' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Cactinea', 1000.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Queima barriga', 'Fórmula do formulário das parceiras, seção Formas farmacêuticas.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', '1 dose ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Queima barriga' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Morosil', 500.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Saciedade', 'Fórmula do formulário das parceiras, seção Formas farmacêuticas.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', '1 dose ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Saciedade' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Chia', 1.5::numeric, 'g')
  ) AS c(ord, subst, qtd, un);

COMMIT;
