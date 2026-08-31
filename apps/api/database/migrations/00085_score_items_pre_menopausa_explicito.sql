-- +goose Up
-- Itens de escore de fase do ciclo continuavam elegíveis para mulher pós-menopausa.
--
-- `AppliesToPatient` (internal/models/score_item.go) só filtra por menopausa quando o ITEM tem
-- `post_menopause` preenchido: NULL no item significa "não me filtre por isso". Os itens de fase
-- do ciclo (folicular, lútea, ovulatória, gestação) e os "Pré-Menopausa" estavam com NULL, então
-- continuavam elegíveis para QUALQUER mulher, inclusive pós-menopausa e inclusive gestação.
--
-- Com isso, um mesmo lab_test_code passava a ter 4 a 6 itens elegíveis para a mesma paciente e o
-- `pickScoringItem` escolhia um arbitrariamente. Efeito medido numa paciente de 63 anos,
-- pós-menopausa, no painel de 02/05/2026:
--
--   Estradiol <5 pg/mL   → nível 5 no item Pós-Menopausa (Sem TRH); nível 0 num item de fase
--   FSH 43 mUI/mL        → nível 5 em Pós-menopausa (faixa 40-100); nível 0 em Pré-menopausa Dia 3
--   Progesterona 0,05    → nível 5 em Pós-Menopausa; nível 0 em Fase Lútea
--
-- Ou seja: valores normais de pós-menopausa pontuando como o pior nível possível. Atinge toda
-- paciente mulher com painel hormonal.
--
-- A correção é tornar o pré-menopausa explícito (false) nos itens que disputam o mesmo
-- lab_test_code com um item pós-menopausa. Os demais itens "female" seguem NULL de propósito:
-- hemoglobina, hematócrito, antropometria e testes físicos valem para toda mulher.

UPDATE public.score_items si
   SET post_menopause = false,
       updated_at     = now()
 WHERE si.deleted_at IS NULL
   AND si.gender = 'female'
   AND si.post_menopause IS NULL
   AND si.lab_test_code IS NOT NULL
   AND EXISTS (
         SELECT 1 FROM public.score_items pm
          WHERE pm.deleted_at IS NULL
            AND pm.gender = 'female'
            AND pm.post_menopause IS TRUE
            AND pm.lab_test_code = si.lab_test_code
       );

-- +goose StatementBegin
DO $$
DECLARE restantes int;
BEGIN
  SELECT count(*) INTO restantes
    FROM public.score_items si
   WHERE si.deleted_at IS NULL AND si.gender = 'female' AND si.post_menopause IS NULL
     AND si.lab_test_code IS NOT NULL
     AND EXISTS (SELECT 1 FROM public.score_items pm
                  WHERE pm.deleted_at IS NULL AND pm.gender = 'female'
                    AND pm.post_menopause IS TRUE AND pm.lab_test_code = si.lab_test_code);
  IF restantes > 0 THEN
    RAISE EXCEPTION 'ainda restam % itens de fase de ciclo com post_menopause NULL', restantes;
  END IF;
END $$;
-- +goose StatementEnd

-- +goose Down
UPDATE public.score_items si
   SET post_menopause = NULL,
       updated_at     = now()
 WHERE si.deleted_at IS NULL
   AND si.gender = 'female'
   AND si.post_menopause IS FALSE
   AND si.lab_test_code IS NOT NULL
   AND EXISTS (
         SELECT 1 FROM public.score_items pm
          WHERE pm.deleted_at IS NULL
            AND pm.gender = 'female'
            AND pm.post_menopause IS TRUE
            AND pm.lab_test_code = si.lab_test_code
       );
