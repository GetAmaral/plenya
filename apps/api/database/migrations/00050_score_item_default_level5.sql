-- Pré-seleção de nível 5 por padrão em templates de anamnese.
-- Itens marcados (histórico de doença / uso de medicação) chegam preenchidos como
-- "sem doença" / "sem uso" (nível 5), agilizando a anamnese.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.score_items
  ADD COLUMN IF NOT EXISTS default_level5 boolean NOT NULL DEFAULT false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.score_items
  DROP COLUMN IF EXISTS default_level5;
-- +goose StatementEnd
