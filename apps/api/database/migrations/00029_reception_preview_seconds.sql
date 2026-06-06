-- Atendimento IA — X2: janela de preview do Auto (segundos que o rascunho fica visível no
-- compositor antes do envio automático). Separado do X (debounce). Default 10s.
-- Plano: docs/emr/plano-atendimento-ia-global.md.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.reception_settings
    ADD COLUMN IF NOT EXISTS preview_seconds integer NOT NULL DEFAULT 10;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.reception_settings DROP COLUMN IF EXISTS preview_seconds;
-- +goose StatementEnd
