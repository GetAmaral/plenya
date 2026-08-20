-- Fórmulas do material do Pentravan (Fagron), por via transdérmica e vaginal.
--
-- Cada componente já entra com a CATEGORIA de receita que carrega: as fórmulas com
-- testosterona e oxandrolona saem como Controle Especial (lista C5), que é o que a Portaria
-- 344/98 pede e o que o sistema não fazia sozinho. last_review NULO.
BEGIN;

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · miomatose e endometriose vaginal', 'Miodesin e Pentravan por via vaginal na miomatose uterina e na endometriose.', E'endometriose\nmiomatose uterina',
         'vaginal', 'external', 'vaginal', 'Pentravan qsp 1 g', 60::numeric, 'g',
         'Aplicar 1 g por via vaginal, à noite, por até 2 meses', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · miomatose e endometriose vaginal' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Miodesin', 170::numeric, 'mg', 'simple', '')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · proteção endometrial', 'Progesterona por via vaginal para proteção do endométrio.', E'proteção endometrial\núltimos 13 a 15 dias do mês',
         'vaginal', 'external', 'vaginal', 'Pentravan qsp 1 g', 60::numeric, 'g',
         'Aplicar 1 g por via vaginal à noite, nos últimos 13 a 15 dias do mês', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · proteção endometrial' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Progesterona micronizada', 50::numeric, 'mg', 'simple', 'Material dá faixa de 20 a 80 mg por grama.')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · endometriose com gestrinona', 'Gestrinona por via vaginal na endometriose, com os estudos de Maia Jr. em Pentravan.', E'dor de endometriose\naplicação três vezes por semana',
         'vaginal', 'external', 'vaginal', 'Pentravan qsp 1 g', 60::numeric, 'g',
         'Aplicar 1 g por via vaginal, 3 vezes por semana', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · endometriose com gestrinona' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Gestrinona', 5::numeric, 'mg', 'simple', 'Material dá 2,5 mg ou 5 mg por grama.')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · deficiência androgênica feminina', 'Reposição androgênica feminina por via transdérmica em região de pouco pelo e pouco tecido adiposo.', E'deficiência androgênica feminina comprovada\naplicar em pulsos ou antebraços',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL com pump dosador, 1 vez ao dia ou em dias alternados', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · deficiência androgênica feminina' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Testosterona micronizada', 3::numeric, 'mg', 'c5', 'Material dá faixa de 0,5 a 5 mg por mL.')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · climatério transdérmico', 'Estradiol e estriol por via transdérmica no alívio de sintomas climatéricos.', E'sintomas climatéricos\n25 dias de uso com 5 de intervalo',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL com pump dosador ao dia por 25 dias, com intervalo de 5 dias', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · climatério transdérmico' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, '17-beta-estradiol', 1::numeric, 'mg', 'simple', 'Material dá 0,25 a 2 mg por mL.'),
  (1, 'Estriol', 4::numeric, 'mg', 'simple', 'Material dá 2 a 8 mg por mL.')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · estimulante sexual feminino', 'Sildenafila por uso vulvar como estimulante sexual feminino.', E'uso sob demanda\naplicar 30 minutos antes',
         'vaginal', 'external', 'vaginal', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL na região dos lábios vaginais 30 minutos antes da relação sexual', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · estimulante sexual feminino' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Citrato de sildenafila', 0.25::numeric, '%', 'simple', '')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · vulvodínia', 'PEA e baclofeno por via tópica no alívio da vulvodínia.', E'vulvodínia\naplicar nas áreas afetadas',
         'vaginal', 'external', 'vaginal', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL nas áreas afetadas, 1 a 2 vezes ao dia', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · vulvodínia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Palmitoiletanolamida', 10::numeric, 'mg', 'simple', ''),
  (1, 'Baclofeno', 50::numeric, 'mg', 'simple', '')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · mastalgia cíclica', 'Danazol por via transdérmica na mama, na mastalgia cíclica.', E'mastalgia cíclica\naplicação na mama',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 g', 60::numeric, 'g',
         'Aplicar 1 mL (1 pump) na mama, 1 vez ao dia', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · mastalgia cíclica' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Danazol', 50::numeric, 'mg', 'simple', '')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · AMPK e longevidade', 'Metformina por via transdérmica para modulação de AMPK.', E'modulação de AMPK\nduas aplicações ao dia',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL (1 pump), 2 vezes ao dia, em região com poucos pelos', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · AMPK e longevidade' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Metformina', 75::numeric, 'mg', 'simple', 'Material dá 50 a 100 mg por mL.')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · flacidez e envelhecimento cutâneo', 'Silício, estriol e resveratrol por via tópica facial, na prevenção de flacidez.', E'flacidez facial\nuso em rosto e pescoço',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL com pump dosador no rosto e pescoço, 1 vez ao dia', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · flacidez e envelhecimento cutâneo' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'SiliciuMax', 5::numeric, '%', 'simple', ''),
  (1, 'Estriol', 0.3::numeric, '%', 'simple', ''),
  (2, 'Trans-resveratrol', 3::numeric, '%', 'simple', '')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · modulação de testosterona com resveratrol', 'Testosterona com trans-resveratrol como inibidor de aromatase, por via transdérmica.', E'declínio androgênico masculino\nassociação com inibidor de aromatase',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL (1 pump), 1 vez ao dia, em pulsos, antebraços ou ombros', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · modulação de testosterona com resveratrol' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Testosterona micronizada', 50::numeric, 'mg', 'c5', ''),
  (1, 'Trans-resveratrol', 50::numeric, 'mg', 'simple', '')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · declínio androgênico masculino', 'Testosterona transdérmica no declínio androgênico masculino com deficiência comprovada.', E'declínio androgênico com deficiência documentada\naplicar em região de pouco pelo',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL (1 pump), 1 vez ao dia, em pulsos, antebraços ou ombros', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · declínio androgênico masculino' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Testosterona micronizada', 60::numeric, 'mg', 'c5', 'Material dá faixa de 40 a 90 mg por mL.')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · testosterona com tadalafila', 'Testosterona associada a inibidor de fosfodiesterase-5, uso diário.', E'declínio androgênico com disfunção erétil\nuso diário',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL (1 pump), 1 vez ao dia, em pulsos, antebraços ou ombros', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · testosterona com tadalafila' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Testosterona micronizada', 50::numeric, 'mg', 'c5', ''),
  (1, 'Tadalafila', 5::numeric, 'mg', 'simple', '')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · disfunção erétil sob demanda', 'Alprostadil tópico, uso sob demanda.', E'disfunção erétil\n5 a 10 pumps por aplicação',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar de 5 a 10 pumps, no mínimo 3 vezes por semana, de 5 a 30 minutos antes da atividade sexual', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · disfunção erétil sob demanda' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Alprostadil', 100::numeric, 'mcg', 'simple', 'Cada pump contém 100 mcg.')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · disfunção erétil com fentolamina', 'Alprostadil com mesilato de fentolamina, uso tópico sob demanda.', E'disfunção erétil\nintervalo mínimo de 24 horas entre aplicações',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar de 5 a 10 pumps, no mínimo 3 vezes por semana, mantendo 24 horas entre aplicações', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · disfunção erétil com fentolamina' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Alprostadil', 100::numeric, 'mcg', 'simple', ''),
  (1, 'Mesilato de fentolamina', 4::numeric, 'mg', 'simple', '')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · oxandrolona em sarcopenia', 'Oxandrolona transdérmica em sarcopenia, no material do fornecedor.', E'sarcopenia com deficiência documentada\nregistrar indicação no prontuário',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL (1 pump), 1 vez ao dia, em região com poucos pelos', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · oxandrolona em sarcopenia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Oxandrolona', 10::numeric, 'mg', 'c5', 'A fórmula original do material se intitula ''sarcopenia e ganho de peso''; ganho de massa é finalidade vedada pela Resolução CFM 2.333/2023.')
) AS c(ord,s,q,u,cat,n);

COMMIT;
