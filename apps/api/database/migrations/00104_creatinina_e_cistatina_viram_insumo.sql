-- +goose Up
-- CREATININA E CISTATINA C VIRAM INSUMO: 20 e 18 pontos passam a 15 cada.
--
-- A migration 00103 pôs a eTFG no escore, com 22 pontos, e criou uma sobreposição que ficou
-- registrada lá como pendência: a eTFG é conta sobre a creatinina e a cistatina C, então a mesma
-- fisiologia renal passou a pontuar duas vezes — 60 pontos entre os três itens, num escore onde
-- inflar peso é justamente o defeito conhecido.
--
-- Decisão do Getúlio, em 2026-09-09: **15 pontos cada** para creatinina e cistatina C. A eTFG fica
-- com os 22 porque é ela que estagia a doença e organiza a conduta; as outras duas continuam
-- pontuando, e devem continuar, porque cada uma diz algo que a eTFG sozinha não diz — creatinina
-- alta com eTFG preservada aponta massa muscular ou creatina, e cistatina C alterada com creatinina
-- normal aponta o contrário. São insumos com leitura própria, e não redundância pura.
--
-- Total renal: 60 → 52 pontos.
--
-- Por que por MIGRATION e não por psql: peso de item de escore é dado de catálogo, roda no deploy
-- do api e precisa chegar igual em dev e prod. Ver a Regra de Ouro 1 e
-- docs/emr/migrations-decisao.md.
--
-- Não recalcula snapshot nenhum: o escore de cada paciente é refeito no próximo `recalc-scores` ou
-- na próxima anamnese. Mexer em snapshot antigo reescreveria o número que já foi mostrado ao
-- paciente, com uma régua que não era a daquele dia.
-- +goose StatementBegin
DO $pesos$
DECLARE
  v_n int;
BEGIN
  UPDATE public.score_items
     SET points = 15, updated_at = now()
   WHERE lab_test_code IN ('PLN357DC859', 'PLN42D34D96')
     AND deleted_at IS NULL
     AND points <> 15;
  GET DIAGNOSTICS v_n = ROW_COUNT;

  -- Zero linhas é estado válido (migration reaplicada, ou base onde o catálogo não foi semeado);
  -- mais de dois é sinal de que o catálogo ganhou variantes por sexo desses itens desde 00104, e
  -- aí o peso precisa de decisão nova em vez de um UPDATE cego.
  IF v_n > 2 THEN
    RAISE EXCEPTION 'esperava no máximo 2 itens (creatinina e cistatina C), atualizou %', v_n;
  END IF;
  RAISE NOTICE 'peso ajustado para 15 em % item(ns)', v_n;
END
$pesos$;
-- +goose StatementEnd

-- +goose Down
UPDATE public.score_items SET points = 20, updated_at = now()
 WHERE lab_test_code = 'PLN357DC859' AND deleted_at IS NULL;
UPDATE public.score_items SET points = 18, updated_at = now()
 WHERE lab_test_code = 'PLN42D34D96' AND deleted_at IS NULL;
