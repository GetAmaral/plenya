-- +goose Up
-- Terapia de reposição hormonal como dado da paciente e como filtro de elegibilidade.
--
-- O catálogo tem dois itens de estradiol pós-menopausa, "Com TRH" e "Sem TRH", com faixas
-- opostas: 5 pg/mL é o MELHOR nível (5) em quem não repõe e o PIOR (0) em quem repõe. Como
-- nada no modelo dizia se a paciente repõe, os dois ficavam elegíveis ao mesmo tempo e o
-- `pickScoringItem` escolhia um arbitrariamente. Medido numa paciente de 63 anos sem TRH:
-- estradiol <5 pg/mL classificado como nível 0 quando o correto era 5.
--
-- Mesma mecânica do `post_menopause` já existente: NULL no item significa "não filtro por
-- isso"; item que declara exige o dado na paciente e exige que bata (models.ScoreItem.
-- AppliesToPatient).

ALTER TABLE public.patients
  ADD COLUMN IF NOT EXISTS hormone_therapy boolean;

COMMENT ON COLUMN public.patients.hormone_therapy IS
  'Paciente faz terapia de reposição hormonal. Muda a faixa de referência dos hormônios. NULL = não informado.';

ALTER TABLE public.score_items
  ADD COLUMN IF NOT EXISTS hormone_therapy boolean;

COMMENT ON COLUMN public.score_items.hormone_therapy IS
  'Restringe o item por uso de TRH: true = só quem repõe, false = só quem não repõe, NULL = não filtra.';

-- Os dois itens que motivaram a coluna. Casados pelo nome porque é o que distingue os dois
-- irmãos do mesmo lab_test_code.
UPDATE public.score_items
   SET hormone_therapy = true, updated_at = now()
 WHERE deleted_at IS NULL AND name = 'Estradiol - Mulheres Pós-Menopausa (Com TRH)';

UPDATE public.score_items
   SET hormone_therapy = false, updated_at = now()
 WHERE deleted_at IS NULL AND name = 'Estradiol - Mulheres Pós-Menopausa (Sem TRH)';

-- +goose StatementBegin
DO $$
DECLARE marcados int;
BEGIN
  SELECT count(*) INTO marcados FROM public.score_items
   WHERE deleted_at IS NULL AND hormone_therapy IS NOT NULL;
  IF marcados <> 2 THEN
    RAISE EXCEPTION 'esperava exatamente 2 itens de estradiol marcados por TRH, encontrei %', marcados;
  END IF;
END $$;
-- +goose StatementEnd

-- +goose Down
ALTER TABLE public.score_items DROP COLUMN IF EXISTS hormone_therapy;
ALTER TABLE public.patients   DROP COLUMN IF EXISTS hormone_therapy;
