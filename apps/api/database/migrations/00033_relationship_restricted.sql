-- Lívia — campo "restrito" em todo o 360 (Fase E do plano docs/emr/plano-livia-memoria-dossie-360.md §12).
-- Restrito = aparece para a EQUIPE no EMR, mas NÃO é exposto à IA quando ela avalia o que sabe do
-- paciente. Em relationship_facts o conceito já existia como `sensitive` (sem uso) → renomeado para
-- `restricted` por clareza/consistência; em important_people e relationship_events é coluna nova.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.relationship_facts RENAME COLUMN sensitive TO restricted;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.important_people ADD COLUMN IF NOT EXISTS restricted boolean NOT NULL DEFAULT false;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.relationship_events ADD COLUMN IF NOT EXISTS restricted boolean NOT NULL DEFAULT false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.relationship_events DROP COLUMN IF EXISTS restricted;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.important_people DROP COLUMN IF EXISTS restricted;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.relationship_facts RENAME COLUMN restricted TO sensitive;
-- +goose StatementEnd
