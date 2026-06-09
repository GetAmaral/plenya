-- Sessão persistente / "manter conectado" + rotação resiliente (janela de graça + famílias).
-- Plano: docs/emr/estudo-sessao-login-persistente.md.
--   family_id     — agrupa a cadeia de tokens de um login (revogação por roubo = família toda)
--   rotated_at    — quando o token foi rotacionado (base da janela de graça)
--   remember      — sessão de aparelho confiável (refresh longo deslizante)
--   last_used_at  — última atividade (tela "Minhas sessões" + diagnóstico)

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.refresh_tokens ADD COLUMN IF NOT EXISTS family_id uuid;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.refresh_tokens ADD COLUMN IF NOT EXISTS rotated_at timestamptz;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.refresh_tokens ADD COLUMN IF NOT EXISTS remember boolean NOT NULL DEFAULT false;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.refresh_tokens ADD COLUMN IF NOT EXISTS last_used_at timestamptz;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_refresh_tokens_family ON public.refresh_tokens (family_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_refresh_tokens_family;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.refresh_tokens DROP COLUMN IF EXISTS last_used_at;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.refresh_tokens DROP COLUMN IF EXISTS remember;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.refresh_tokens DROP COLUMN IF EXISTS rotated_at;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.refresh_tokens DROP COLUMN IF EXISTS family_id;
-- +goose StatementEnd
