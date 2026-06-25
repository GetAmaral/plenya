-- Motivo da NÃO-classificação (level nulo) de cada exame, para o usuário ver POR QUE
-- um exame casado não recebeu nível: sem item de score, não se aplica ao paciente,
-- valor fora das faixas, resultado não numérico, etc. Diferente de match_reason (catálogo).
-- Ver docs/emr/plano-importacao-laudos-melhorias.md (Fase 3, revisão de visibilidade).

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.lab_results
  ADD COLUMN IF NOT EXISTS classify_reason text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.lab_results
  DROP COLUMN IF EXISTS classify_reason;
-- +goose StatementEnd
