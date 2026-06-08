-- Web Push (VAPID) — canal de notificação do EMR para desktop/PWA (sem app nativo).
-- Plano: docs/emr/plano-webpush-notificacoes.md. Uma linha = um navegador inscrito.

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.web_push_subscriptions (
    id            uuid PRIMARY KEY,
    user_id       uuid NOT NULL,
    endpoint      text NOT NULL,
    p256dh        text NOT NULL,
    auth          text NOT NULL,
    device_label  varchar(200),
    user_agent    varchar(400),
    last_seen_at  timestamptz NOT NULL DEFAULT now(),
    created_at    timestamptz NOT NULL DEFAULT now(),
    updated_at    timestamptz NOT NULL DEFAULT now()
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE UNIQUE INDEX IF NOT EXISTS idx_web_push_endpoint ON public.web_push_subscriptions (endpoint);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_web_push_user ON public.web_push_subscriptions (user_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_web_push_last_seen ON public.web_push_subscriptions (last_seen_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.web_push_subscriptions;
-- +goose StatementEnd
