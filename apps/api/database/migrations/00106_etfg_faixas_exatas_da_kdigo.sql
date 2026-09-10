-- +goose Up
-- eTFG: as faixas da 00103 estagiavam a fronteira um estágio PIOR do que a KDIGO.
--
-- A 00103 pôs os limites em 15/30/45/60/90 e justificou o desencontro com a diretriz dizendo que a
-- eTFG calculada é contínua e cai em cima da fronteira raramente. Isso estava errado por um motivo
-- que estava no próprio código: `SincronizaTFGDerivada` arredonda o resultado para INTEIRO antes de
-- gravar. Então todo valor real entre 59,5 e 60,5 vira exatamente 60 — e com a convenção meio-aberta
-- `(inferior, superior]` do projeto, 60 cai em `(45,60]`, que é G3a. A KDIGO diz G2. O mesmo em 90
-- (saía G2 onde a KDIGO diz G1, que é ≥90), em 45, em 30 e em 15. Não é raro: é o caso comum.
--
-- Estágio errado não é detalhe de escore. G3a contra G2 muda meta de pressão, muda ajuste de dose
-- de droga com eliminação renal e muda o momento de encaminhar ao nefrologista.
--
-- A correção mantém a convenção do projeto e desloca os limites em uma unidade, o que reproduz a
-- KDIGO EXATAMENTE para valor inteiro, que é como filtração é reportada:
--
--   >89        -> G1   (KDIGO ≥90)          90 e acima
--   (59, 89]   -> G2   (KDIGO 60-89)        60 a 89
--   (44, 59]   -> G3a  (KDIGO 45-59)        45 a 59
--   (29, 44]   -> G3b  (KDIGO 30-44)        30 a 44
--   (14, 29]   -> G4   (KDIGO 15-29)        15 a 29
--   <=14       -> G5   (KDIGO <15)          14 e abaixo
--
-- Sem sobreposição, e a ordem de avaliação (do nível 0 para cima, primeira faixa que bate vence)
-- devolve o estágio certo em cada fronteira.
-- +goose StatementBegin
DO $faixas$
DECLARE
  v_item uuid;
  v_n    int;
BEGIN
  SELECT id INTO v_item FROM public.score_items
   WHERE lab_test_code = 'ETFG' AND deleted_at IS NULL
   LIMIT 1;

  IF v_item IS NULL THEN
    RAISE NOTICE 'sem item de escore para ETFG: 00106 não fez nada';
    RETURN;
  END IF;

  UPDATE public.score_levels sl
     SET name = v.rotulo, operator = v.op, lower_limit = v.inf, upper_limit = v.sup, updated_at = now()
    FROM (VALUES
      (5, 'G1 · ≥90',      '>',       '89',  NULL),
      (4, 'G2 · 60-89',    'between', '59',  '89'),
      (3, 'G3a · 45-59',   'between', '44',  '59'),
      (2, 'G3b · 30-44',   'between', '29',  '44'),
      (1, 'G4 · 15-29',    'between', '14',  '29'),
      (0, 'G5 · ≤14',      '<=',      NULL,  '14')
    ) AS v(nivel, rotulo, op, inf, sup)
   WHERE sl.item_id = v_item AND sl.level = v.nivel AND sl.deleted_at IS NULL;
  GET DIAGNOSTICS v_n = ROW_COUNT;

  IF v_n <> 6 THEN
    RAISE EXCEPTION 'esperava corrigir 6 faixas da eTFG, corrigiu %', v_n;
  END IF;
  RAISE NOTICE 'faixas da eTFG alinhadas à KDIGO (% linhas)', v_n;
END
$faixas$;
-- +goose StatementEnd

-- +goose Down
DO $faixas_down$
DECLARE
  v_item uuid;
BEGIN
  SELECT id INTO v_item FROM public.score_items
   WHERE lab_test_code = 'ETFG' AND deleted_at IS NULL LIMIT 1;
  IF v_item IS NULL THEN RETURN; END IF;

  UPDATE public.score_levels sl
     SET name = v.rotulo, operator = v.op, lower_limit = v.inf, upper_limit = v.sup, updated_at = now()
    FROM (VALUES
      (5, 'G1 · >90',      '>',       '90',  NULL),
      (4, 'G2 · 60-90',    'between', '60',  '90'),
      (3, 'G3a · 45-60',   'between', '45',  '60'),
      (2, 'G3b · 30-45',   'between', '30',  '45'),
      (1, 'G4 · 15-30',    'between', '15',  '30'),
      (0, 'G5 · ≤15',      '<=',      NULL,  '15')
    ) AS v(nivel, rotulo, op, inf, sup)
   WHERE sl.item_id = v_item AND sl.level = v.nivel AND sl.deleted_at IS NULL;
END
$faixas_down$;
