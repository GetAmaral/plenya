-- +goose Up
-- Item de escore que só faz sentido dentro de um contexto de outro exame.
--
-- Caso que motivou: `PSA Livre/Total (%Free PSA)` pontua 28 e marca ≤10% como o PIOR nível. A razão
-- livre/total só é interpretável quando o PSA TOTAL está na zona cinzenta (4 a 10 ng/mL); com PSA
-- total baixo ela não diz nada. Num paciente com PSA total de 1,81 e razão de 8,8%, o item virava o
-- achado número 1 da devolutiva — um alarme de próstata onde não há alarme — e ainda tirava 28
-- pontos do escore dele.
--
-- Não é caso isolado nem se resolve desligando o item: com PSA total na zona cinzenta a razão é
-- justamente o exame que decide biópsia. O que faltava era o catálogo saber dizer "só vale quando".

ALTER TABLE public.score_items
  ADD COLUMN IF NOT EXISTS requires_lab_code varchar(100),
  ADD COLUMN IF NOT EXISTS requires_min      double precision,
  ADD COLUMN IF NOT EXISTS requires_max      double precision;

COMMENT ON COLUMN public.score_items.requires_lab_code IS
  'Código de lab_test_definitions de que este item depende. NULL = sem condição.';
COMMENT ON COLUMN public.score_items.requires_min IS
  'O item só se aplica quando o resultado mais recente de requires_lab_code é >= este valor. NULL = sem piso.';
COMMENT ON COLUMN public.score_items.requires_max IS
  'O item só se aplica quando o resultado mais recente de requires_lab_code é <= este valor. NULL = sem teto.';

CREATE INDEX IF NOT EXISTS idx_score_items_requires_lab
  ON public.score_items (requires_lab_code) WHERE requires_lab_code IS NOT NULL;

-- %Free PSA passa a valer só na zona cinzenta do PSA total.
-- Limites: 4 a 10 ng/mL, a faixa em que a razão livre/total discrimina hiperplasia de neoplasia.
UPDATE public.score_items si
   SET requires_lab_code = (SELECT code FROM lab_test_definitions
                             WHERE name ILIKE 'PSA total%' AND deleted_at IS NULL
                             ORDER BY length(name) LIMIT 1),
       requires_min = 4,
       requires_max = 10
 WHERE si.name ILIKE '%Free PSA%' OR si.name ILIKE '%Livre/Total%';

-- +goose Down
DROP INDEX IF EXISTS idx_score_items_requires_lab;
ALTER TABLE public.score_items
  DROP COLUMN IF EXISTS requires_lab_code,
  DROP COLUMN IF EXISTS requires_min,
  DROP COLUMN IF EXISTS requires_max;
