-- Visibilidade da importação de laudos: por que um exame ficou sem casar com o catálogo
-- e de onde veio (PDF vs manual). Ver docs/emr/plano-importacao-laudos-melhorias.md (Fase 3).

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.lab_results
  ADD COLUMN IF NOT EXISTS match_reason text,
  ADD COLUMN IF NOT EXISTS source varchar(20) NOT NULL DEFAULT 'manual';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.lab_results
  DROP COLUMN IF EXISTS match_reason,
  DROP COLUMN IF EXISTS source;
-- +goose StatementEnd
