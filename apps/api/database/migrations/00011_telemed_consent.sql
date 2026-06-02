-- Consentimento de telemedicina (CFM 2.314/2022) no appointment (P3 — peça telemed).
-- Carimbo no prontuário: quando/quem registrou + texto exato apresentado + modalidade.
-- Aditivo, sem tocar dados existentes.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.appointments
    ADD COLUMN IF NOT EXISTS telemed_consent_at timestamp with time zone,
    ADD COLUMN IF NOT EXISTS telemed_consent_text text,
    ADD COLUMN IF NOT EXISTS telemed_consent_by_user_id uuid,
    ADD COLUMN IF NOT EXISTS telemed_consent_mode character varying(10);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.appointments
    DROP COLUMN IF EXISTS telemed_consent_at,
    DROP COLUMN IF EXISTS telemed_consent_text,
    DROP COLUMN IF EXISTS telemed_consent_by_user_id,
    DROP COLUMN IF EXISTS telemed_consent_mode;
-- +goose StatementEnd
