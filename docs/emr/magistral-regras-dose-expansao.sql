-- Expansão das regras de dose dinâmica: de 10 para o conjunto que cobre o repertório.
--
-- Toda regra é escrita em dose DIÁRIA. A tela divide pelas tomadas que lê da posologia e mostra a
-- conta — sem isso, uma regra de 5.000 UI/dia numa fórmula tomada duas vezes ao dia entregaria
-- 10.000. A trava também é diária, e corta antes de dividir.
--
-- Idempotente: apaga a regra do componente antes de recriar.

BEGIN;

-- ---------------------------------------------------------------------------------------------
-- 0. Duas doses da fonte que não dá para inferir — ficam como estão, anotadas
-- ---------------------------------------------------------------------------------------------
UPDATE magistral_formula_template_components c
   SET note = 'Fonte traz 100 mcg. Zinco terapêutico é de 10 a 30 mg: a unidade parece errada, mas o documento não dá o número certo em nenhum outro lugar. Conferir com a farmácia.'
  FROM magistral_formula_templates t
 WHERE t.id = c.template_id AND c.substance = 'Zinco quelato' AND c.unit = 'mcg';

UPDATE magistral_formula_template_components c
   SET note = 'Fonte traz 100 UI. As demais fórmulas do documento usam 1.000 UI ou mais; a regra por exame ajusta.'
  FROM magistral_formula_templates t
 WHERE t.id = c.template_id AND c.substance = 'Vitamina D3' AND c.unit IN ('UI','ui') AND c.quantity <= 100;

-- ---------------------------------------------------------------------------------------------
-- 1. Limpeza das regras que este arquivo cria
-- ---------------------------------------------------------------------------------------------
DELETE FROM magistral_formula_template_rules r
 USING magistral_formula_template_components c, magistral_formula_templates t
 WHERE c.id = r.template_component_id AND t.id = c.template_id AND t.deleted_at IS NULL
   AND (
     (c.substance = 'Vitamina D3'              AND c.unit IN ('UI','ui')) OR
     (c.substance = 'Metilcobalamina'          AND c.unit = 'mcg')        OR
     (c.substance = 'Zinco quelato'            AND c.unit = 'mg')         OR
     (c.substance = 'Metilfolato'              AND c.unit = 'mcg')        OR
     (c.substance = 'Ferro'                    AND c.unit = 'mg')         OR
     (c.substance = 'Berberina'                AND c.unit = 'mg')         OR
     (c.substance = 'Testosterona micronizada' AND c.unit = 'mg')         OR
     (c.substance = 'Selenometionina'          AND c.unit = 'mcg' AND t.name ~* 'tireoid|hipotireo')
   );

-- ---------------------------------------------------------------------------------------------
-- 2. Vitamina D3 por faixa de 25-OH-D — todas as fórmulas que a contêm
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id AND t.deleted_at IS NULL
     WHERE c.substance = 'Vitamina D3' AND c.unit IN ('UI','ui') AND c.deleted_at IS NULL
), ins AS (
    INSERT INTO magistral_formula_template_rules
      (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLN1BF562ED', 'ng/mL', 250, 1000, 7000, 365,
      'Alvo de 40 a 60 ng/mL, como nas aulas da pós. A diretriz da Endocrine Society de 2024 não fixa alvo e desaconselha rastreio: a divergência é deliberada. Acima de 4.000 UI/dia passa do teto de suplemento da IN 28.'
      FROM comp RETURNING id)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot FROM ins, (VALUES
  (0, NULL::numeric, 20::numeric,   7000::numeric, 'deficiência'),
  (1, 20::numeric,   30::numeric,   5000::numeric, 'insuficiência'),
  (2, 30::numeric,   40::numeric,   3000::numeric, 'abaixo do alvo'),
  (3, 40::numeric,   60::numeric,   2000::numeric, 'dentro do alvo'),
  (4, 60::numeric,   NULL::numeric, 1000::numeric, 'acima do alvo')
) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 3. Metilcobalamina por faixa de B12
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id AND t.deleted_at IS NULL
     WHERE c.substance = 'Metilcobalamina' AND c.unit = 'mcg' AND c.deleted_at IS NULL
), ins AS (
    INSERT INTO magistral_formula_template_rules
      (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLN9B054BBD', 'pg/mL', 50, 100, 2000, 365,
      'Alvo acima de 550 pg/mL, como nas aulas da pós. Na literatura, abaixo de 200 é deficiência e de 200 a 500 há deficiência funcional (a EFNS usa 500). Oral de 1.000 mcg equivale à via intramuscular.'
      FROM comp RETURNING id)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot FROM ins, (VALUES
  (0, NULL::numeric, 300::numeric,  1000::numeric, 'deficiência'),
  (1, 300::numeric,  550::numeric,  500::numeric,  'abaixo do alvo'),
  (2, 550::numeric,  NULL::numeric, 100::numeric,  'dentro do alvo')
) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 4. Zinco por faixa de zinco sérico
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id AND t.deleted_at IS NULL
     WHERE c.substance = 'Zinco quelato' AND c.unit = 'mg' AND c.deleted_at IS NULL
), ins AS (
    INSERT INTO magistral_formula_template_rules
      (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLN7B3753DD', 'µg/dL', 5, 10, 29, 365,
      'Alvo de pelo menos 100 µg/dL, como nas aulas da pós. O teto de 29,59 mg/dia é o do Anexo IV da IN 28; acima disso, por meses, o zinco compete com o cobre. A dose é do elemento.'
      FROM comp RETURNING id)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot FROM ins, (VALUES
  (0, NULL::numeric, 70::numeric,   29::numeric, 'deficiência'),
  (1, 70::numeric,   100::numeric,  20::numeric, 'abaixo do alvo'),
  (2, 100::numeric,  NULL::numeric, 10::numeric, 'dentro do alvo')
) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 5. Metilfolato por homocisteína
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id AND t.deleted_at IS NULL
     WHERE c.substance = 'Metilfolato' AND c.unit = 'mcg' AND c.deleted_at IS NULL
), ins AS (
    INSERT INTO magistral_formula_template_rules
      (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLNC01E9624', 'µmol/L', 100, 200, 1000, 365,
      'Homocisteína acima de 15 µmol/L é hiper-homocisteinemia; o alvo é abaixo de 10. Repor B12 antes ou junto: folato isolado corrige o hemograma e deixa a lesão neurológica da falta de B12 avançar.'
      FROM comp RETURNING id)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot FROM ins, (VALUES
  (0, 15::numeric,   NULL::numeric, 1000::numeric, 'hiper-homocisteinemia'),
  (1, 10::numeric,   15::numeric,   800::numeric,  'acima do alvo'),
  (2, NULL::numeric, 10::numeric,   400::numeric,  'dentro do alvo')
) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 6. Ferro por ferritina — com um buraco de propósito acima do alvo
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id AND t.deleted_at IS NULL
     WHERE c.substance = 'Ferro' AND c.unit = 'mg' AND c.deleted_at IS NULL
), ins AS (
    INSERT INTO magistral_formula_template_rules
      (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLNCEFB97FD', 'ng/mL', 5, 20, 60, 180,
      'Ferritina acima de 70 ng/mL não tem faixa cadastrada de propósito: repor ferro sem falta comprovada é risco, não conveniência. Em queda capilar o alvo costuma ser acima de 40 a 70. Dose em ferro elementar, longe de cálcio e zinco; dia sim, dia não absorve proporcionalmente mais.'
      FROM comp RETURNING id)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot FROM ins, (VALUES
  (0, NULL::numeric, 15::numeric, 60::numeric, 'deficiência absoluta'),
  (1, 15::numeric,   30::numeric, 45::numeric, 'deficiência'),
  (2, 30::numeric,   70::numeric, 30::numeric, 'abaixo do alvo capilar')
) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 7. Berberina por hemoglobina glicada
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id AND t.deleted_at IS NULL
     WHERE c.substance = 'Berberina' AND c.unit = 'mg' AND c.deleted_at IS NULL
), ins AS (
    INSERT INTO magistral_formula_template_rules
      (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLN3FC5EDA6', '%', 50, 500, 1500, 365,
      'Os ensaios usam 1 a 1,5 g/dia fracionados. A berberina inibe CYP3A4 e P-glicoproteína: conferir interação com o que o paciente já toma, sobretudo estatina, ciclosporina e anticoagulante.'
      FROM comp RETURNING id)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot FROM ins, (VALUES
  (0, 6.4::numeric,  NULL::numeric, 1500::numeric, 'faixa de diabetes'),
  (1, 5.7::numeric,  6.4::numeric,  1000::numeric, 'pré-diabetes'),
  (2, NULL::numeric, 5.7::numeric,  500::numeric,  'normal')
) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 8. Testosterona por testosterona total — o buraco acima de 350 é a regra do CFM
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id AND t.deleted_at IS NULL
     WHERE c.substance = 'Testosterona micronizada' AND c.unit = 'mg' AND c.deleted_at IS NULL
), ins AS (
    INSERT INTO magistral_formula_template_rules
      (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLNDE1A5575', 'ng/dL', 10, 20, 90, 180,
      'Acima de 350 ng/dL não há faixa cadastrada de propósito: a Resolução CFM 2.333/2023 só admite reposição com deficiência comprovada e nexo causal. O sistema não sugere dose para quem não tem indicação.'
      FROM comp RETURNING id)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot FROM ins, (VALUES
  (0, NULL::numeric, 250::numeric, 60::numeric, 'deficiência'),
  (1, 250::numeric,  350::numeric, 40::numeric, 'limítrofe')
) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 9. Selenometionina por anti-TPO — só nas fórmulas de tireoide
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id AND t.deleted_at IS NULL
     WHERE c.substance = 'Selenometionina' AND c.unit = 'mcg' AND c.deleted_at IS NULL
       AND t.name ~* 'tireoid|hipotireo'
)
INSERT INTO magistral_formula_template_rules
  (id, template_component_id, kind, lab_code, lab_unit, lab_operator, lab_threshold,
   dose_if_true, dose_if_false, round_to, min_dose, max_dose, max_data_age_days, note)
SELECT uuid_generate_v7(), comp.id, 'lab_threshold', 'PLNF479B8FF', 'IU/mL', 'gt', 35,
       200, 100, 10, 50, 200, 730,
       'Anti-TPO alterado: 200 mcg de selênio elementar. Nas metanálises, a selenometionina é a única forma com efeito nos anticorpos; a queda de TSH é pequena (0,2 mIU/L) e só um terço dos braços mostrou queda de anti-TPO. Toxicidade a partir de 300 a 400 mcg/dia.'
  FROM comp;

COMMIT;
