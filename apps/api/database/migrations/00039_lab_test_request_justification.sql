-- Justificativa clínica padrão que acompanha o exame no pedido.
-- Carrega automaticamente quando o exame é selecionado em qualquer painel.
-- Uso primário: exames de alto custo (ex.: escore de cálcio coronariano) que
-- exigem fundamentação para autorização/laudo de indicação.
-- Plano: ~/.claude/plans (templates de pedidos de exames).

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.lab_test_definitions
  ADD COLUMN IF NOT EXISTS request_justification text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.lab_test_definitions DROP COLUMN IF EXISTS request_justification;
-- +goose StatementEnd
