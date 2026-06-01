-- Tabela dedicada para os jti single-use dos magic links do Score Light.
-- Antes o jti era gravado em refresh_tokens com user_id=uuid.Nil, violando a FK
-- refresh_tokens→users (SQLSTATE 23503), porque o User só é criado no ConfirmClaim.
-- Esta tabela não tem FK de usuário de propósito.

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.magic_link_tokens (
    id uuid NOT NULL,
    token_hash character varying(64) NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    used_at timestamp with time zone,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    CONSTRAINT magic_link_tokens_pkey PRIMARY KEY (id)
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_magic_link_tokens_token_hash ON public.magic_link_tokens (token_hash);
CREATE INDEX IF NOT EXISTS idx_magic_link_tokens_expires_at ON public.magic_link_tokens (expires_at);
CREATE INDEX IF NOT EXISTS idx_magic_link_tokens_used_at ON public.magic_link_tokens (used_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.magic_link_tokens;
-- +goose StatementEnd
