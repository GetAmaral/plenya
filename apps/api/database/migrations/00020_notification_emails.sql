-- Bucket "Notificações": e-mails automáticos (no-reply, newsletters, notificações de
-- plataformas) que chegam na caixa mas não viram Lead. Antes eram descartados no ingest;
-- agora ficam acessíveis numa aba separada da caixa de e-mail.
-- Ver docs/emr/plano-conversas-redesign.md (Fase 5).

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.notification_emails (
    id          uuid PRIMARY KEY,
    from_email  varchar(320) NOT NULL,
    from_name   varchar(255),
    subject     text NOT NULL DEFAULT '',
    body_text   text NOT NULL DEFAULT '',
    message_id  varchar(998),
    received_at timestamptz NOT NULL DEFAULT now(),
    is_read     boolean NOT NULL DEFAULT false,
    created_at  timestamptz NOT NULL DEFAULT now(),
    updated_at  timestamptz NOT NULL DEFAULT now()
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE UNIQUE INDEX IF NOT EXISTS idx_notification_emails_message_id
    ON public.notification_emails (message_id) WHERE message_id IS NOT NULL;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_notification_emails_received_at
    ON public.notification_emails (received_at DESC);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_notification_emails_is_read
    ON public.notification_emails (is_read);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.notification_emails;
-- +goose StatementEnd
