-- Nota clínica/anamnese gerada por IA (AI scribe) a partir do transcript da
-- teleconsulta. Rascunho NÃO assinável — o médico revisa e insere na ClinicalNote.
-- Aditivo em telemed_recordings.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.telemed_recordings
    ADD COLUMN IF NOT EXISTS generated_note_json   text,
    ADD COLUMN IF NOT EXISTS generated_note_format varchar(16),
    ADD COLUMN IF NOT EXISTS generated_note_status varchar(16) NOT NULL DEFAULT 'none',
    ADD COLUMN IF NOT EXISTS generated_note_model  varchar(64),
    ADD COLUMN IF NOT EXISTS generated_note_at     timestamp with time zone,
    ADD COLUMN IF NOT EXISTS generated_note_error  text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.telemed_recordings
    DROP COLUMN IF EXISTS generated_note_json,
    DROP COLUMN IF EXISTS generated_note_format,
    DROP COLUMN IF EXISTS generated_note_status,
    DROP COLUMN IF EXISTS generated_note_model,
    DROP COLUMN IF EXISTS generated_note_at,
    DROP COLUMN IF EXISTS generated_note_error;
-- +goose StatementEnd
