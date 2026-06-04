-- Recepcionista virtual Fase 2: estado de automação por conversa + usuário-bot.
-- conversation_automations guarda o modo (off|copilot|auto) por (owner_type, owner_id).
-- O usuário "Assistente Plenya" (UUID fixo) é o ActorUserID das mensagens enviadas pela IA.
-- Ver docs/emr/plano-recepcionista-virtual.md.

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.conversation_automations (
    id               uuid PRIMARY KEY,
    owner_type       varchar(10) NOT NULL,
    owner_id         uuid NOT NULL,
    mode             varchar(10) NOT NULL DEFAULT 'off',
    fallback_minutes integer NOT NULL DEFAULT 0,
    paused_until     timestamptz,
    last_bot_at      timestamptz,
    updated_by       uuid,
    created_at       timestamptz NOT NULL DEFAULT now(),
    updated_at       timestamptz NOT NULL DEFAULT now(),
    deleted_at       timestamptz,
    CONSTRAINT chk_conversation_automations_mode CHECK (mode IN ('off','copilot','auto')),
    CONSTRAINT chk_conversation_automations_owner_type CHECK (owner_type IN ('lead','patient'))
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE UNIQUE INDEX IF NOT EXISTS uq_conversation_automation_owner
    ON public.conversation_automations USING btree (owner_type, owner_id)
    WHERE deleted_at IS NULL;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_conversation_automations_mode
    ON public.conversation_automations USING btree (mode) WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- Usuário-bot (Assistente Plenya). UUID fixo referenciado no código (botUserID).
-- +goose StatementBegin
INSERT INTO public.users (id, email, name, roles, certificate_mode, created_at, updated_at)
VALUES (
    '019e9301-0000-7000-a000-0000000b0b01',
    'assistente@plenyasaude.com.br',
    'Assistente Plenya',
    '["secretary"]'::jsonb,
    'local',
    now(), now()
)
ON CONFLICT DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.conversation_automations;
-- +goose StatementEnd
-- +goose StatementBegin
DELETE FROM public.users WHERE id = '019e9301-0000-7000-a000-0000000b0b01';
-- +goose StatementEnd
