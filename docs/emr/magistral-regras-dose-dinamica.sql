-- Regras de dose dinâmica das fórmulas-base.
--
-- Cada regra tem procedência declarada na própria nota: o que vem das aulas da pós (RAG) e o que
-- vem da literatura, inclusive quando as duas DIVERGEM. Onde diverge, a faixa segue a conduta da
-- casa e a nota diz qual é a diretriz contrária — quem prescreve precisa ver as duas.
--
-- Piso e teto são obrigatórios em toda regra; a sugestão nunca escreve na receita.
-- Idempotente: apaga a regra do componente antes de recriar.

BEGIN;

-- ---------------------------------------------------------------------------------------------
-- 0. Correções de dado encontradas ao escrever as regras
-- ---------------------------------------------------------------------------------------------

-- Vitamina D3 de 50 UI na fórmula de hipotireoidismo: 50 UI não tem efeito biológico nenhum,
-- é erro de transcrição do formulário. A base vira 2.000 UI e a regra ajusta pelo exame.
UPDATE magistral_formula_template_components c
   SET quantity = 2000
  FROM magistral_formula_templates t
 WHERE t.id = c.template_id AND t.name = 'Hipotireoidismo, suporte de cofatores'
   AND c.substance = 'Vitamina D3' AND c.quantity = 50;

-- Magnésio quelato estava marcado como dose do ELEMENTO numa fórmula e como dose do INSUMO em
-- outra. Mesma substância, duas leituras, três vezes de diferença na cápsula. Padroniza em
-- elemento (260 mg de quelato a 30% ≈ 80 mg de magnésio).
UPDATE magistral_formula_template_components c
   SET quantity = 80, as_elemental = true
  FROM magistral_formula_templates t
 WHERE t.id = c.template_id AND t.name = 'Ansiedade diurna'
   AND c.substance = 'Magnésio quelato' AND c.as_elemental = false;

-- Magnésio glicina do sachê mitocondrial passa a ser escrito em elemento, que é a unidade em que
-- a regra por peso raciocina. 500 mg de bisglicinato a 30% ≈ 150 mg de magnésio.
UPDATE magistral_formula_template_components c
   SET quantity = 150, as_elemental = true
  FROM magistral_formula_templates t
 WHERE t.id = c.template_id AND t.name = 'Sachê matinal mitocondrial'
   AND c.substance = 'Magnésio glicina' AND c.as_elemental = false;

-- Fórmula sem posologia nenhuma: a receita sairia sem dizer como tomar.
UPDATE magistral_formula_templates
   SET posology = '1 sachê ao deitar'
 WHERE name = 'Sachê noturno de relaxamento' AND coalesce(trim(posology), '') = '';

-- O "exemplo de regra" era andaime meu, não fórmula. Vira a fórmula de vitamina D de verdade.
UPDATE magistral_formula_templates
   SET name = 'Vitamina D conforme exame',
       indication = 'Reposição de vitamina D com a dose ajustada pela 25-hidroxivitamina D mais recente do paciente.',
       indication_bullets = 'Reposição de vitamina D guiada pelo exame'||chr(10)||
                            'Faixa-alvo de 40 a 60 ng/mL'||chr(10)||
                            'Reavaliar 25-OH-D em 90 dias'
 WHERE name = 'Vitamina D conforme exame (exemplo de regra)';

-- ---------------------------------------------------------------------------------------------
-- 0.1 Limpeza das regras que este arquivo recria
--
-- Statement PRÓPRIO, e não um CTE junto do INSERT: há UNIQUE em template_component_id, e o
-- INSERT dentro do mesmo statement enxerga o snapshot anterior ao DELETE — a limpeza não teria
-- acontecido ainda quando a chave fosse conferida.
-- ---------------------------------------------------------------------------------------------
DELETE FROM magistral_formula_template_rules r
 USING magistral_formula_template_components c, magistral_formula_templates t
 WHERE c.id = r.template_component_id AND t.id = c.template_id
   AND (t.name, c.substance) IN (
        ('Vitamina D conforme exame',                'Vitamina D3'),
        ('Antioxidante e imunidade',                 'Vitamina D3'),
        ('Hipotireoidismo, suporte de cofatores',    'Vitamina D3'),
        ('Sono completo',                            'Metilcobalamina'),
        ('Fadiga pós atividade física',              'Metilcobalamina'),
        ('Antioxidante amplo',                       'Zinco quelato'),
        ('Fadiga pós atividade física',              'Zinco quelato'),
        ('Hipotireoidismo, suporte de cofatores',    'Zinco quelato'),
        ('Hipotireoidismo, suporte de cofatores',    'Selenometionina'),
        ('Sachê matinal mitocondrial',               'Magnésio glicina')
   );

-- ---------------------------------------------------------------------------------------------
-- 1. Vitamina D3 por faixa de 25-hidroxivitamina D
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id
      FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id
     WHERE c.substance = 'Vitamina D3' AND c.deleted_at IS NULL AND t.deleted_at IS NULL
       AND t.name IN ('Vitamina D conforme exame', 'Antioxidante e imunidade',
                      'Hipotireoidismo, suporte de cofatores')
), ins AS (
    INSERT INTO magistral_formula_template_rules
        (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLN1BF562ED', 'ng/mL', 500, 1000, 7000, 365,
           'Alvo de 40 a 60 ng/mL, como nas aulas da pós. A diretriz da Endocrine Society de 2024 não fixa alvo e desaconselha rastreio: a divergência é deliberada. Acima de 4.000 UI/dia passa do teto de suplemento da IN 28 e vira decisão prescritiva.'
      FROM comp
    RETURNING id
)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot
  FROM ins, (VALUES
      (0, NULL::numeric, 20::numeric,   7000::numeric, 'deficiência'),
      (1, 20::numeric,   30::numeric,   5000::numeric, 'insuficiência'),
      (2, 30::numeric,   40::numeric,   3000::numeric, 'abaixo do alvo'),
      (3, 40::numeric,   60::numeric,   2000::numeric, 'dentro do alvo'),
      (4, 60::numeric,   NULL::numeric, 1000::numeric, 'acima do alvo')
  ) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 2. Metilcobalamina por faixa de vitamina B12
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id
      FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id
     WHERE c.substance = 'Metilcobalamina' AND c.deleted_at IS NULL AND t.deleted_at IS NULL
       AND t.name IN ('Sono completo', 'Fadiga pós atividade física')
), ins AS (
    INSERT INTO magistral_formula_template_rules
        (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLN9B054BBD', 'pg/mL', 50, 100, 1000, 365,
           'Alvo acima de 550 pg/mL, como nas aulas da pós. Na literatura, abaixo de 200 é deficiência e de 200 a 500 há deficiência funcional (a EFNS usa 500 como corte). Oral de 1.000 mcg equivale à via intramuscular. Confirmar com homocisteína.'
      FROM comp
    RETURNING id
)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot
  FROM ins, (VALUES
      (0, NULL::numeric, 300::numeric,  1000::numeric, 'deficiência'),
      (1, 300::numeric,  550::numeric,  500::numeric,  'abaixo do alvo'),
      (2, 550::numeric,  NULL::numeric, 100::numeric,  'dentro do alvo')
  ) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 3. Zinco quelato por faixa de zinco sérico (dose em zinco elementar)
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id
      FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id
     WHERE c.substance = 'Zinco quelato' AND c.deleted_at IS NULL AND t.deleted_at IS NULL
       AND t.name IN ('Antioxidante amplo', 'Fadiga pós atividade física',
                      'Hipotireoidismo, suporte de cofatores')
), ins AS (
    INSERT INTO magistral_formula_template_rules
        (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLN7B3753DD', 'µg/dL', 5, 10, 40, 365,
           'Alvo de pelo menos 100 µg/dL, como nas aulas da pós. Teto de 40 mg/dia de zinco elementar: acima disso, por meses, compete com o cobre. A dose é do elemento; o insumo sai pelo fator de correção.'
      FROM comp
    RETURNING id
)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot
  FROM ins, (VALUES
      (0, NULL::numeric, 70::numeric,   30::numeric, 'deficiência'),
      (1, 70::numeric,   100::numeric,  20::numeric, 'abaixo do alvo'),
      (2, 100::numeric,  NULL::numeric, 10::numeric, 'dentro do alvo')
  ) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 4. Selenometionina por anti-TPO (limiar, não faixa: a conduta aqui é binária)
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id
      FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id
     WHERE c.substance = 'Selenometionina' AND c.deleted_at IS NULL AND t.deleted_at IS NULL
       AND t.name = 'Hipotireoidismo, suporte de cofatores'
)
INSERT INTO magistral_formula_template_rules
    (id, template_component_id, kind, lab_code, lab_unit, lab_operator, lab_threshold,
     dose_if_true, dose_if_false, round_to, min_dose, max_dose, max_data_age_days, note)
SELECT uuid_generate_v7(), comp.id, 'lab_threshold', 'PLNF479B8FF', 'IU/mL', 'gt', 35,
       200, 100, 10, 50, 200, 730,
       'Anti-TPO alterado: 200 mcg de selênio elementar. Nas metanálises, a selenometionina é a única forma com efeito nos anticorpos; a queda de TSH é pequena (0,2 mIU/L) e só um terço dos braços mostrou queda de anti-TPO. Toxicidade a partir de 300 a 400 mcg/dia.'
  FROM comp;

-- ---------------------------------------------------------------------------------------------
-- 5. Magnésio por peso (sachê — a única forma que comporta a dose)
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id
      FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id
     WHERE c.substance = 'Magnésio glicina' AND c.deleted_at IS NULL AND t.deleted_at IS NULL
       AND t.name = 'Sachê matinal mitocondrial'
)
INSERT INTO magistral_formula_template_rules
    (id, template_component_id, kind, per_kg, round_to, min_dose, max_dose, max_data_age_days, note)
SELECT uuid_generate_v7(), comp.id, 'per_kg', 5, 50, 200, 350, 180,
       'Cinco miligramas de magnésio elementar por quilo. Teto de 350 mg/dia de magnésio suplementar (UL do IOM); acima disso o efeito laxativo domina. A dose é do elemento; o bisglicinato sai pelo fator de correção.'
  FROM comp;

COMMIT;
