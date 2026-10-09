-- +goose Up
-- GORDURA VISCERAL POR DXA, EM GRAMAS: o número que o densitômetro entrega e o escore não lia.
--
-- O escore só tinha gordura visceral como ÁREA (cm²), nos itens de anamnese GORDURA_VISCERAL_*.
-- Área vem de tomografia num corte L4-L5 ou de DXA de outro fabricante. O DXA GE Lunar com CoreScan
-- (Prodigy e iDXA) entrega MASSA em gramas e VOLUME em cm³, e esses números ficavam fora do escore
-- em silêncio. Caso que motivou: mulher de 52 anos com 1.005 g / 1.065 cm³ e nenhum item que os lesse.
--
-- MASSA E VOLUME SÃO A MESMA MEDIDA. O CoreScan converte por densidade constante de 0,94 g/cm³
-- (da Rosa et al., Healthcare 2021: "VAT mass (g) was automatically transposed into volume (cm3)
-- using a constant correction factor (0.94 g/mL)"). Conferido num laudo real: 1.005 g ÷ 0,94 = 1.069,
-- o laudo traz 1.065. Por isso UM exame em gramas, e o laudo em cm³ entra pela conversão curada:
-- principal = secundária ÷ fator, então fator = 1 ÷ 0,94 = 1,0638 (cm³ = g × 1,0638).
--
-- AS FAIXAS, E POR QUE SÃO ESTAS (estudo completo em docs/emr/escore-vat-dxa-gramas.md)
--
-- Todos os estudos abaixo usaram GE Lunar + CoreScan. DXA de outro fabricante NÃO é comparável.
--   Miazgowski, PLoS One 2017, jovens saudáveis 20-30 a: P90 mulher 515 g; P70 homem 677 g.
--   Meredith-Jones, IJO 2021, 1.482 adultos: risco cardiometabólico a partir de 800 g (mulher ≥40 a)
--     e 1.200 g (homem ≥40 a); 700 g e 1.000 g abaixo dos 40.
--   SPAH (São Paulo), ACR 2024, idosos ≥65 a: síndrome metabólica a partir de 600 g (M) e 642 g (H).
--   da Rosa, Healthcare 2021, militares brasileiros: 1.086 g (AUC 0,80).
--   Bestaoui, PLoS One 2025, Argélia: 1.082 g (M, AUC 0,855) e 1.369 g (H, AUC 0,868).
--   Lundblad (Tromsø), J Obes 2021, ≥40 a: 1.134 g (M) e 1.859 g (H).
--
-- Mulher: ≤500 ótimo (≈P90 de jovens saudáveis, abaixo de todo corte de risco, inclusive o do idoso
--   brasileiro) · 500-800 bom (abaixo do corte de risco ≥40 a) · 800-1.100 médio (entre o corte de
--   risco e os de síndrome metabólica) · >1.100 ruim (nos cortes de síndrome metabólica).
-- Homem: ≤700 ótimo (≈P70 de jovens saudáveis) · 700-1.100 bom (abaixo do militar brasileiro e de
--   Otago ≥40) · 1.100-1.400 médio · >1.400 ruim (acima da Argélia; Tromsø é mais tolerante).
-- Sem nível 1 e 3: não há número publicado que sustente essa granulação. Faixas aprovadas pelo
-- Dr. Getúlio em 2026-10-09.
--
-- VAT% (visceral ÷ gordura androide) fica de fora: em Tromsø não acrescentou nada à massa para
-- prever síndrome metabólica.
--
-- TUDO OU NADA, como a 00102: se o exame-pai da densitometria não existir (base recriada, CI), avisa
-- e não deixa definição órfã.
-- +goose StatementBegin
DO $vat$
DECLARE
  v_pai      uuid;
  v_subgrupo uuid;
  v_ordem    int;
  v_item     uuid;
  v_sexo     text;
  v_def      uuid;
BEGIN
  SELECT id INTO v_pai FROM public.lab_test_definitions
   WHERE code = 'PLND1D39793' AND deleted_at IS NULL LIMIT 1;

  SELECT si.subgroup_id, si."order" INTO v_subgrupo, v_ordem
    FROM public.score_items si
   WHERE si.lab_test_code = 'PLND1D39793' AND si.deleted_at IS NULL
   ORDER BY si."order"
   LIMIT 1;

  IF v_pai IS NULL OR v_subgrupo IS NULL THEN
    RAISE NOTICE 'densitometria de corpo inteiro (PLND1D39793) ausente: 00108 não tem onde pendurar a gordura visceral e não fez nada';
    RETURN;
  END IF;

  -- `id` fixo para dev e prod terem o MESMO id.
  INSERT INTO public.lab_test_definitions
    (id, code, name, short_name, alt_names, category, is_requestable, parent_test_id, unit,
     result_type, sex_applicability, is_active, display_order, patient_gloss, clinical_significance,
     created_at, updated_at)
  SELECT '01a120d4-8ebd-7fa5-bd1f-db5b5e0d0995'::uuid, 'PLNDXAVAT01',
         'Densitometria - Gordura Visceral (VAT)', 'VAT-DXA',
         '["vat", "tecido adiposo visceral", "gordura visceral", "visceral adipose tissue", "corescan", "tav"]'::jsonb,
         'imaging', false, v_pai, 'g', 'numeric', 'all', true,
         coalesce((SELECT display_order FROM public.lab_test_definitions WHERE id = v_pai), 0) + 3,
         'gordura entre os órgãos',
         'Tecido adiposo visceral estimado pelo DXA GE Lunar com CoreScan, na região androide. O laudo '
         'traz massa (g) e volume (cm³), que são a mesma medida pela densidade de 0,94 g/cm³. Não é '
         'comparável a DXA de outro fabricante nem à área (cm²) de tomografia.',
         now(), now()
  WHERE NOT EXISTS (SELECT 1 FROM public.lab_test_definitions WHERE code = 'PLNDXAVAT01');

  -- O id vem pelo código, não pelo literal: se o exame já existia com outro id (criado à mão), a
  -- conversão tem de pendurar NELE, senão o laudo em cm³ fica sem converter.
  SELECT id INTO v_def FROM public.lab_test_definitions
   WHERE code = 'PLNDXAVAT01' AND deleted_at IS NULL LIMIT 1;

  -- Laudo em cm³ vira gramas. As duas grafias, porque a busca da conversão é literal.
  INSERT INTO public.lab_test_unit_conversions
    (id, lab_test_definition_id, main_unit, secondary_unit, conversion_factor, created_at, updated_at)
  SELECT uuid_generate_v7(), v_def, 'g', u.sec, 1 / 0.94, now(), now()
  FROM (VALUES ('cm³'), ('cm3')) AS u(sec)
  WHERE NOT EXISTS (SELECT 1 FROM public.lab_test_unit_conversions c
                     WHERE c.lab_test_definition_id = v_def
                       AND c.secondary_unit = u.sec AND c.deleted_at IS NULL);

  -- Um item por sexo, com o mesmo peso dos itens de gordura visceral em cm² (18 mulher, 22 homem).
  FOREACH v_sexo IN ARRAY ARRAY['female', 'male'] LOOP
    v_item := NULL;
    SELECT id INTO v_item FROM public.score_items
     WHERE lab_test_code = 'PLNDXAVAT01' AND gender = v_sexo AND deleted_at IS NULL
     LIMIT 1;

    IF v_item IS NULL THEN
      INSERT INTO public.score_items
        (id, name, unit, points, "order", subgroup_id, lab_test_code, gender, clinical_relevance,
         created_at, updated_at)
      VALUES (uuid_generate_v7(),
        CASE v_sexo WHEN 'female' THEN 'Densitometria - Gordura Visceral (g) - Mulheres'
                    ELSE 'Densitometria - Gordura Visceral (g) - Homens' END,
        'g', CASE v_sexo WHEN 'female' THEN 18 ELSE 22 END, v_ordem + 3, v_subgrupo, 'PLNDXAVAT01', v_sexo,
        'Gordura visceral é a que envolve os órgãos do abdome e é metabolicamente ativa: prediz '
        'síndrome metabólica, resistência à insulina e risco cardiovascular melhor que o peso ou o '
        'IMC. Faixas ancoradas em sete estudos com o mesmo equipamento (GE Lunar com CoreScan), dos '
        'percentis de jovens saudáveis aos cortes de síndrome metabólica; ver '
        'docs/emr/escore-vat-dxa-gramas.md.',
        now(), now())
      RETURNING id INTO v_item;
    END IF;

    INSERT INTO public.score_levels (id, item_id, level, name, operator, lower_limit, upper_limit, created_at, updated_at)
    SELECT uuid_generate_v7(), v_item, v.nivel, v.rotulo, v.op, v.inf, v.sup, now(), now()
    FROM (VALUES
      ('female', 5, '≤500',        '<=',      NULL,   '500'),
      ('female', 4, '500-800',     'between', '500',  '800'),
      ('female', 2, '800-1.100',   'between', '800',  '1100'),
      ('female', 0, '>1.100',      '>',       '1100', NULL),
      ('male',   5, '≤700',        '<=',      NULL,   '700'),
      ('male',   4, '700-1.100',   'between', '700',  '1100'),
      ('male',   2, '1.100-1.400', 'between', '1100', '1400'),
      ('male',   0, '>1.400',      '>',       '1400', NULL)
    ) AS v(sexo, nivel, rotulo, op, inf, sup)
    WHERE v.sexo = v_sexo
      AND NOT EXISTS (SELECT 1 FROM public.score_levels sl
                       WHERE sl.item_id = v_item AND sl.deleted_at IS NULL);
  END LOOP;
END
$vat$;
-- +goose StatementEnd

-- +goose Down
DELETE FROM public.score_levels sl USING public.score_items si
 WHERE sl.item_id = si.id AND si.lab_test_code = 'PLNDXAVAT01';
DELETE FROM public.score_items WHERE lab_test_code = 'PLNDXAVAT01';
DELETE FROM public.lab_test_unit_conversions c USING public.lab_test_definitions d
 WHERE c.lab_test_definition_id = d.id AND d.code = 'PLNDXAVAT01';
DELETE FROM public.lab_test_definitions WHERE code = 'PLNDXAVAT01';
