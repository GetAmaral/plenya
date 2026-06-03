-- Hardening de assinatura + e-CPF em nuvem (PSC/broker) + modo de assinatura da prescrição.
-- Aditivo em users (vínculo do certificado em nuvem) e prescriptions (signature_mode).
-- Certificado em nuvem (NGS2): a chave privada do médico fica num HSM do PSC; o backend
-- dispara a assinatura via API e o titular confirma (push/OTP). Guardamos só referências
-- (o token de credencial é cifrado AES-256-GCM, nunca em claro).

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.users
    ADD COLUMN IF NOT EXISTS certificate_mode        varchar(20)  NOT NULL DEFAULT 'local',
    ADD COLUMN IF NOT EXISTS cloud_cert_provider     varchar(40),
    ADD COLUMN IF NOT EXISTS cloud_cert_pem          text,
    ADD COLUMN IF NOT EXISTS cloud_credential_ref    varchar(200),
    ADD COLUMN IF NOT EXISTS cloud_credential_token  text,
    ADD COLUMN IF NOT EXISTS cloud_credential_expiry timestamp with time zone;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.prescriptions
    ADD COLUMN IF NOT EXISTS signature_mode varchar(20);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.users
    DROP COLUMN IF EXISTS certificate_mode,
    DROP COLUMN IF EXISTS cloud_cert_provider,
    DROP COLUMN IF EXISTS cloud_cert_pem,
    DROP COLUMN IF EXISTS cloud_credential_ref,
    DROP COLUMN IF EXISTS cloud_credential_token,
    DROP COLUMN IF EXISTS cloud_credential_expiry;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.prescriptions
    DROP COLUMN IF EXISTS signature_mode;
-- +goose StatementEnd
