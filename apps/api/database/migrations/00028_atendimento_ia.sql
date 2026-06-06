-- Atendimento IA — controle GLOBAL do recepcionista virtual (modo baseline + janela de horário
-- semanal + tempos X/Y/Z) e rascunhos gerados no servidor (debounced) por conversa.
-- Plano: docs/emr/plano-atendimento-ia-global.md.
--   reception_settings: singleton (1 linha) com a config global runtime.
--   conversation_suggested_replies: rascunho atual por conversa, mostrado no compositor e usado
--   pelo envio automático (Auto) / espera de revisão (Copiloto).
-- O override por conversa continua em conversation_automations (a presença da linha = flag explícita).

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.reception_settings (
    id                       uuid PRIMARY KEY,
    baseline_mode            varchar(10) NOT NULL DEFAULT 'off' CHECK (baseline_mode IN ('off','copilot','auto')),
    schedule_enabled         boolean     NOT NULL DEFAULT false,
    schedule                 jsonb       NOT NULL DEFAULT '{}'::jsonb,
    debounce_seconds         integer     NOT NULL DEFAULT 30,
    copilot_fallback_minutes integer     NOT NULL DEFAULT 10,
    off_alert_minutes        integer     NOT NULL DEFAULT 20,
    updated_by               uuid,
    created_at               timestamptz NOT NULL DEFAULT now(),
    updated_at               timestamptz NOT NULL DEFAULT now()
);
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO public.reception_settings (id, baseline_mode, schedule_enabled, schedule)
VALUES ('019e9301-0000-7000-a000-0000000a1a01', 'off', false, '{}'::jsonb)
ON CONFLICT (id) DO NOTHING;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.conversation_suggested_replies (
    id                   uuid PRIMARY KEY,
    owner_type           varchar(10) NOT NULL,
    owner_id             uuid        NOT NULL,
    reply                text        NOT NULL,
    action               varchar(20) NOT NULL DEFAULT 'answer',
    model                varchar(60),
    based_on_activity_id uuid,
    created_at           timestamptz NOT NULL DEFAULT now(),
    updated_at           timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT uq_suggested_reply_owner UNIQUE (owner_type, owner_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.conversation_suggested_replies;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS public.reception_settings;
-- +goose StatementEnd
