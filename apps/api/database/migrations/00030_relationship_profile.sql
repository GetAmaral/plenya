-- Lívia — memória de longo prazo (SOCIAL) por pessoa: relationship_profiles.
-- 1 por (owner_type, owner_id). SÓ dado social/relacional/administrativo — NUNCA clínico
-- (LGPD/CFM: o clínico vive no prontuário; o que o lead conta de clínico fica só na janela
-- curta da conversa, nunca no resumo). Fase A do plano
-- docs/emr/plano-livia-memoria-dossie-360.md (§0, §3.1, §9).
--   rolling_summary    : resumo rolante social da relação/conversa, injetado no prompt da Lívia.
--   summary_msg_count  : nº de mensagens já resumidas (gatilho de regeneração a cada +5).
--   summary_updated_at : quando o resumo foi regerado pela última vez.
--   relationship_stage : estágio do funil (reservado p/ fases seguintes).

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.relationship_profiles (
    id                 uuid PRIMARY KEY,
    owner_type         varchar(20) NOT NULL,
    owner_id           uuid        NOT NULL,
    rolling_summary    text        NOT NULL DEFAULT '',
    summary_updated_at timestamptz,
    summary_msg_count  integer     NOT NULL DEFAULT 0,
    relationship_stage varchar(40) NOT NULL DEFAULT '',
    created_at         timestamptz NOT NULL DEFAULT now(),
    updated_at         timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT uq_relationship_profile_owner UNIQUE (owner_type, owner_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.relationship_profiles;
-- +goose StatementEnd
