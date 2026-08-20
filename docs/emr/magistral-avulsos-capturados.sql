-- Substâncias e fórmulas que existiam só no banco de desenvolvimento.
--
-- POR QUE ESTE ARQUIVO EXISTE: carregar todos os seeds num banco vazio devolveu 281 substâncias e
-- 122 fórmulas, contra 290 e 132 no dev. A diferença veio de comandos avulsos rodados direto no
-- psql durante o trabalho, que nunca foram capturados em arquivo — deploy de produção nasceria
-- diferente do que foi conferido aqui.
--
-- O teste que achou isso vale mais que o conserto: carga limpa num banco vazio, e comparação de
-- contagem com o dev. Idempotente.

BEGIN;

-- ---------------------------------------------------------------------------------------------
-- Substâncias
-- ---------------------------------------------------------------------------------------------

INSERT INTO magistral_components (id, name, synonyms, cas, dcb_code, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density, density_source, eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok, notes, source, evidence_status, indications, dose_reference, indication_bullets, dose_bullets, elemental_percent, correction_note, brand, in28_nutrient, in28_factor, default_category, regulatory_note, assay_interference, assay_interference_dose, is_active, created_at, updated_at)
SELECT uuid_generate_v7(), 'Rhodiola rosea', 'rodiola, raiz-do-ártico', NULL, NULL, 'mg', 150.0, NULL, NULL, 'por_dia', 0.45, 'classe', false, false, false, false, false, NULL, NULL, NULL, 'parceiro', 'suggested', 'Adaptógeno usado em fadiga associada a estresse e em suporte de foco.', NULL, 'fadiga por estresse
foco e desempenho mental', NULL, NULL, NULL, NULL, NULL, 1.0, 'simple', NULL, NULL, NULL, true, now(), now()
 WHERE NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Rhodiola rosea');

INSERT INTO magistral_components (id, name, synonyms, cas, dcb_code, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density, density_source, eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok, notes, source, evidence_status, indications, dose_reference, indication_bullets, dose_bullets, elemental_percent, correction_note, brand, in28_nutrient, in28_factor, default_category, regulatory_note, assay_interference, assay_interference_dose, is_active, created_at, updated_at)
SELECT uuid_generate_v7(), 'Piperina', 'extrato de pimenta preta, BioPerine', NULL, NULL, 'mg', 10.0, NULL, NULL, 'por_dia', 0.45, 'classe', false, false, false, false, false, NULL, NULL, NULL, 'parceiro', 'suggested', 'Usada para aumentar a biodisponibilidade da curcumina e de outros ativos.', NULL, 'potencializa absorção da curcumina', NULL, NULL, NULL, NULL, NULL, 1.0, 'simple', NULL, NULL, NULL, true, now(), now()
 WHERE NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Piperina');

INSERT INTO magistral_components (id, name, synonyms, cas, dcb_code, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density, density_source, eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok, notes, source, evidence_status, indications, dose_reference, indication_bullets, dose_bullets, elemental_percent, correction_note, brand, in28_nutrient, in28_factor, default_category, regulatory_note, assay_interference, assay_interference_dose, is_active, created_at, updated_at)
SELECT uuid_generate_v7(), 'Policosanol', '', NULL, NULL, 'mg', 30.0, NULL, NULL, 'por_dia', 0.45, 'classe', false, false, false, false, false, NULL, NULL, NULL, 'parceiro', 'suggested', 'Álcool graxo de cana usado em fórmulas de perfil lipídico; evidência heterogênea, com resultados positivos concentrados em estudos cubanos.', NULL, 'coadjuvante lipídico', NULL, NULL, NULL, NULL, NULL, 1.0, 'simple', NULL, NULL, NULL, true, now(), now()
 WHERE NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Policosanol');

INSERT INTO magistral_components (id, name, synonyms, cas, dcb_code, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density, density_source, eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok, notes, source, evidence_status, indications, dose_reference, indication_bullets, dose_bullets, elemental_percent, correction_note, brand, in28_nutrient, in28_factor, default_category, regulatory_note, assay_interference, assay_interference_dose, is_active, created_at, updated_at)
SELECT uuid_generate_v7(), 'L-tirosina', 'tirosina, L TIROSINA', NULL, NULL, 'mg', 500.0, NULL, NULL, 'por_dia', 0.6, 'classe', false, false, false, false, false, NULL, NULL, NULL, 'parceiro', 'suggested', 'Aminoácido precursor de dopamina, noradrenalina e hormônio tireoidiano.', NULL, 'precursor de catecolaminas
cofator da síntese tireoidiana', NULL, NULL, NULL, NULL, 'Tirosina', 1.0, 'simple', NULL, NULL, NULL, true, now(), now()
 WHERE NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'L-tirosina');

INSERT INTO magistral_components (id, name, synonyms, cas, dcb_code, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density, density_source, eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok, notes, source, evidence_status, indications, dose_reference, indication_bullets, dose_bullets, elemental_percent, correction_note, brand, in28_nutrient, in28_factor, default_category, regulatory_note, assay_interference, assay_interference_dose, is_active, created_at, updated_at)
SELECT uuid_generate_v7(), 'D-ribose', 'ribose, D RIBOSE', NULL, NULL, 'g', 5.0, NULL, NULL, 'por_dia', 0.55, 'classe', false, false, false, false, false, NULL, NULL, NULL, 'parceiro', 'suggested', 'Açúcar de cinco carbonos usado em suporte energético e recuperação.', NULL, 'fadiga e recuperação
suporte de ATP', NULL, NULL, NULL, NULL, 'D-ribose', 1.0, 'simple', NULL, NULL, NULL, true, now(), now()
 WHERE NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'D-ribose');

INSERT INTO magistral_components (id, name, synonyms, cas, dcb_code, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density, density_source, eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok, notes, source, evidence_status, indications, dose_reference, indication_bullets, dose_bullets, elemental_percent, correction_note, brand, in28_nutrient, in28_factor, default_category, regulatory_note, assay_interference, assay_interference_dose, is_active, created_at, updated_at)
SELECT uuid_generate_v7(), 'Beta-hidroxibutirato', 'BHB, cetona exógena, sais de BHB', NULL, NULL, 'g', 10.0, 3.0, 12.0, 'por_dia', 0.8, 'classe', false, true, false, false, false, 3, true, NULL, 'pesquisa', 'suggested', 'Corpo cetônico usado como fonte energética alternativa em protocolos cetogênicos e de desempenho cognitivo. Vem como sal de cálcio, magnésio, sódio ou potássio, e a carga desses minerais conta na fórmula.', 'De 3 a 12 g/dia do sal, geralmente fracionados. Higroscópico e salgado: sachê é a forma prática. A carga de sódio e cálcio do sal precisa entrar na conta do dia.', NULL, NULL, NULL, NULL, NULL, NULL, 1.0, 'simple', NULL, NULL, NULL, true, now(), now()
 WHERE NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Beta-hidroxibutirato');

INSERT INTO magistral_components (id, name, synonyms, cas, dcb_code, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density, density_source, eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok, notes, source, evidence_status, indications, dose_reference, indication_bullets, dose_bullets, elemental_percent, correction_note, brand, in28_nutrient, in28_factor, default_category, regulatory_note, assay_interference, assay_interference_dose, is_active, created_at, updated_at)
SELECT uuid_generate_v7(), 'Gynostemma pentaphyllum', 'jiaogulan, ginostema, ginseng do sul', NULL, NULL, 'mg', 300.0, 200.0, 450.0, 'por_dia', 0.5, 'classe', false, false, false, false, false, 3, true, NULL, 'pesquisa', 'suggested', 'Adaptógeno com ação sobre a AMPK, usado em resistência insulínica e esteatose. Amargo característico.', 'De 200 a 450 mg/dia do extrato padronizado em gipenosídeos. Os ensaios em esteatose e resistência insulínica usam essa faixa por 8 a 12 semanas.', NULL, NULL, NULL, NULL, NULL, NULL, 1.0, 'simple', NULL, NULL, NULL, true, now(), now()
 WHERE NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Gynostemma pentaphyllum');

INSERT INTO magistral_components (id, name, synonyms, cas, dcb_code, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density, density_source, eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok, notes, source, evidence_status, indications, dose_reference, indication_bullets, dose_bullets, elemental_percent, correction_note, brand, in28_nutrient, in28_factor, default_category, regulatory_note, assay_interference, assay_interference_dose, is_active, created_at, updated_at)
SELECT uuid_generate_v7(), 'Goma cássia', 'cassia gum, goma de cássia', NULL, NULL, 'mg', NULL, NULL, NULL, 'por_dia', 0.7, 'classe', false, true, false, false, false, 0, true, NULL, 'pesquisa', 'suggested', 'Fibra solúvel usada como espessante e agente de corpo em sachês. Não tem dose terapêutica própria: a quantidade vem da textura pretendida.', 'Sem faixa cadastrada de propósito: é excipiente de textura, não ativo com posologia estabelecida.', NULL, NULL, NULL, NULL, NULL, NULL, 1.0, 'simple', NULL, NULL, NULL, true, now(), now()
 WHERE NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Goma cássia');

INSERT INTO magistral_components (id, name, synonyms, cas, dcb_code, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density, density_source, eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok, notes, source, evidence_status, indications, dose_reference, indication_bullets, dose_bullets, elemental_percent, correction_note, brand, in28_nutrient, in28_factor, default_category, regulatory_note, assay_interference, assay_interference_dose, is_active, created_at, updated_at)
SELECT uuid_generate_v7(), 'Ácido hidroxicítrico', 'HCA, Citrimax, Garcinia cambogia, GARCINEA CAMBOJA, GARCINIA CAMBOJA', NULL, NULL, 'mg', 1500.0, 1500.0, 3000.0, 'por_dia', 0.6, 'classe', false, false, false, false, false, 2, true, NULL, 'pesquisa', 'suggested', 'Extrato de Garcinia cambogia padronizado em ácido hidroxicítrico, usado para saciedade e controle de peso. As metanálises mostram efeito pequeno e inconsistente sobre o peso.', 'De 500 a 1.000 mg três vezes ao dia, antes das refeições, padronizado a 50 a 60% de HCA. O efeito sobre o peso nas metanálises é pequeno e de qualidade baixa.', NULL, NULL, NULL, NULL, NULL, NULL, 1.0, 'simple', NULL, NULL, NULL, true, now(), now()
 WHERE NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Ácido hidroxicítrico');


-- ---------------------------------------------------------------------------------------------
-- Fórmulas-base
-- ---------------------------------------------------------------------------------------------

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Fórmula sublingual do sono', 'Indução do sono por via mais fisiológica, antes de recorrer à melatonina. Atua como precursor de serotonina e melatonina com efeito gabaérgico.', 'Insônia inicial
Precursores de melatonina
Efeito gabaérgico noturno', 'sublingual', 'internal', '', '', 60.0, 'cápsulas', 'Uso sublingual antes de dormir', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Fórmula sublingual do sono' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, '5-HTP', 25.0::numeric, 'mg', 'simple', '', false),
    (1, 'L-teanina', 20.0::numeric, 'mg', 'simple', '', false),
    (2, 'Piridoxal-5-fosfato (P5P)', 10.0::numeric, 'mg', 'simple', '', false)
) AS c(ord, s, q, u, cat, n, e);

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Precursores de melatonina sublingual', 'Estímulo à via triptofano, 5-HTP, serotonina e melatonina. Cautela em usuários de ISRS.', 'Insônia
Suporte à via serotoninérgica', 'sublingual', 'internal', '', '', 60.0, 'cápsulas', 'À noite, por via sublingual', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Precursores de melatonina sublingual' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, '5-HTP', 25.0::numeric, 'mg', 'simple', 'sublingual à noite', false),
    (1, 'Piridoxal-5-fosfato (P5P)', 10.0::numeric, 'mg', 'simple', 'sublingual', false)
) AS c(ord, s, q, u, cat, n, e);

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Sachê matinal mitocondrial', 'Suporte à função mitocondrial e ao metabolismo energético.', 'Otimização da função mitocondrial
Baixa energia em condições crônicas
Pacientes acima de 50 anos', 'sachê', 'internal', '', '', 60.0, 'cápsulas', '1 sachê pela manhã', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Sachê matinal mitocondrial' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, 'L-carnitina', 500.0::numeric, 'mg', 'simple', '', false),
    (1, 'D-ribose', 5.0::numeric, 'g', 'simple', 'cautela em diabéticos', false),
    (2, 'Magnésio glicina', 150.0::numeric, 'mg', 'simple', '', true)
) AS c(ord, s, q, u, cat, n, e);

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Fórmula pré-refeição para controle glicêmico', 'Suporte ao controle glicêmico e à resistência insulínica em contexto de emagrecimento.', 'Resistência insulínica
Apoio ao emagrecimento
Uso pré-refeição', 'cápsula', 'internal', '', '', 60.0, 'cápsulas', 'Antes das refeições', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Fórmula pré-refeição para controle glicêmico' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, 'Citrimax (ácido hidroxicítrico - HCA)', 500.0::numeric, 'mg', 'simple', '', false),
    (1, 'Gimnema silvestre', 300.0::numeric, 'mg', 'simple', 'faixa de 200 a 300 mg', false),
    (2, 'Ginostema pentaphyllum (Gynostemma pentaphyllum)', 200.0::numeric, 'mg', 'simple', '150 a 200 mg se padronizado a 80% de gipenosídeos; cerca de 500 mg sem padronização, para 100 a 180 mg de gipenosídeos', false),
    (3, 'Cromo (GTF ou picolinato)', 300.0::numeric, 'mcg', 'simple', 'faixa de 200 a 300 mcg', true)
) AS c(ord, s, q, u, cat, n, e);

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Combinação colinérgica fosfatidilcolina + alfa-GPC', 'Suporte à síntese de acetilcolina por múltiplas vias de absorção, voltado a cognição e foco.', 'Precursores de acetilcolina
Suporte cognitivo
Diversificação de vias de absorção', 'cápsula', 'internal', '', '', 60.0, 'cápsulas', 'Dividir em até 4 doses conforme tolerabilidade', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Combinação colinérgica fosfatidilcolina + alfa-GPC' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, 'Fosfatidilcolina', 250.0::numeric, 'mg', 'simple', '', false),
    (1, 'Alfa-GPC (L-alfa-glicerofosfocolina)', 500.0::numeric, 'mg', 'simple', 'faixa de 250 a 500 mg', false)
) AS c(ord, s, q, u, cat, n, e);

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Curcumina com piperina', 'Suporte anti-inflamatório e antioxidante com melhora da biodisponibilidade da curcumina.', 'Modulação da inflamação crônica
Aumento da absorção da curcumina
Cautela em uso de anticoagulantes', 'cápsula', 'internal', '', '', 60.0, 'cápsulas', '500 mg a 2 g/dia de curcuminoides, respeitando 5 mg de piperina a cada 500 mg', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Curcumina com piperina' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, 'Curcumina padronizada (95% curcuminoides)', 500.0::numeric, 'mg', 'simple', 'faixa de 500 mg a 2 g/dia conforme tolerância; anticoagulados limitar a ≤500 mg/dia ou 250 mg lipossomada', false),
    (1, 'Piperina', 5.0::numeric, 'mg', 'simple', '5 mg para cada 500 mg de curcumina; avaliar alergia', false)
) AS c(ord, s, q, u, cat, n, e);

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Sachê matinal de energia', 'Exemplo prático de prescrição em sachê para uso matinal.', 'Uso matinal
Suporte energético', 'sachê', 'internal', '', '', 60.0, 'cápsulas', 'Diluir e tomar pela manhã', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Sachê matinal de energia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, 'Glutamina', 4.0::numeric, 'g', 'simple', '', false),
    (1, 'Magnésio glicina', 500.0::numeric, 'mg', 'simple', '', false),
    (2, 'L-tirosina', 500.0::numeric, 'mg', 'simple', '', false),
    (3, 'Goma cássia', 2500.0::numeric, 'mg', 'simple', '', false),
    (4, 'BHB', 3.0::numeric, 'g', 'simple', '', false),
    (5, 'D-ribose', 5.0::numeric, 'g', 'simple', '', false)
) AS c(ord, s, q, u, cat, n, e);

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Sachê noturno de relaxamento', 'Combinação em pó para promover relaxamento e foco no período noturno. Formato em sachê evita ingestão de muitas cápsulas.', 'Relaxamento noturno
Suporte ao foco', 'sachê', 'internal', '', '', 60.0, 'cápsulas', '1 sachê ao deitar', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Sachê noturno de relaxamento' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, 'Magnésio glicina', 500.0::numeric, 'mg', 'simple', '', false),
    (1, 'Magnésio treonato', 500.0::numeric, 'mg', 'simple', '', false),
    (2, 'Inositol', 1.0::numeric, 'g', 'simple', '', false),
    (3, 'Taurina', 1.0::numeric, 'g', 'simple', '', false)
) AS c(ord, s, q, u, cat, n, e);

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Vitamina D conforme exame', 'Reposição de vitamina D com a dose ajustada pela 25-hidroxivitamina D mais recente do paciente.', 'Reposição de vitamina D guiada pelo exame
Faixa-alvo de 40 a 60 ng/mL
Reavaliar 25-OH-D em 90 dias', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 60.0, 'cápsulas', '1 cápsula pela manhã', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Vitamina D conforme exame' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, 'Vitamina D3', 2000.0::numeric, 'UI', 'simple', '', false)
) AS c(ord, s, q, u, cat, n, e);

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Cápsulas de suporte mitocondrial', 'Suplementação oral para biogênese e desempenho mitocondrial em condições crônicas degenerativas, neurológicas ou metabólicas.', 'Disfunção mitocondrial
Condições degenerativas crônicas
Fadiga e baixa energia
Pacientes acima de 50 anos', 'cápsula', 'internal', '', '', 60.0, 'cápsulas', '1 dose ao dia', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Cápsulas de suporte mitocondrial' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, 'Acetil-L-carnitina', 500.0::numeric, 'mg', 'simple', 'em jejum, manhã ou tarde', false),
    (1, 'Coenzima Q10 (ubiquinona) ou ubiquinol', 100.0::numeric, 'mg', 'simple', 'preferir com refeição gordurosa', false),
    (2, 'Vitamina B2 (riboflavina)', 25.0::numeric, 'mg', 'simple', '', false),
    (3, 'Vitamina B3 (nicotinamida)', 100.0::numeric, 'mg', 'simple', '', false),
    (4, 'Vitamina B6 (piridoxal-5-fosfato)', 10.0::numeric, 'mg', 'simple', '', false),
    (5, 'Magnésio dimalato', 500.0::numeric, 'mg', 'simple', '', false),
    (6, 'Ácido alfa-lipoico', 600.0::numeric, 'mg', 'simple', '300 a 600 mg, final da tarde em jejum, pode exigir cápsula gastrorresistente', false),
    (7, 'PQQ (pirroloquinolina quinona)', 20.0::numeric, 'mg', 'simple', '', false)
) AS c(ord, s, q, u, cat, n, e);


COMMIT;
