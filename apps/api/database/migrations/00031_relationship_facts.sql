-- Lívia — memória SEMÂNTICA (fatos atômicos sociais) por pessoa: relationship_facts.
-- N por (owner_type, owner_id). SÓ social/relacional/administrativo — NUNCA clínico (LGPD/CFM:
-- o clínico vive no prontuário). Fase B do plano docs/emr/plano-livia-memoria-dossie-360.md
-- (§2.2, §3). Validade temporal: valid_until NULL = fato ATIVO; UPDATE/DELETE fecham o anterior
-- (valid_until = now) e, no UPDATE, inserem novo ativo. Índice único parcial garante 1 ativo por key.
--   category : taxonomia controlada (identidade_social|familia_rede|preferencias_atendimento|
--              contexto_chegada|relacionamento).
--   key/value: ex. key="profissao" value="engenheiro"; key="indicado_por" value="filha Marina".
--   source   : ai|staff|form|consulta. added_by: user (quando staff). confidence: quando IA.

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.relationship_facts (
    id          uuid PRIMARY KEY,
    owner_type  varchar(20) NOT NULL,
    owner_id    uuid        NOT NULL,
    category    varchar(40) NOT NULL,
    key         varchar(60) NOT NULL,
    value       text        NOT NULL,
    source      varchar(20) NOT NULL DEFAULT 'ai',
    added_by    uuid,
    confidence  real,
    valid_from  timestamptz NOT NULL DEFAULT now(),
    valid_until timestamptz,
    sensitive   boolean     NOT NULL DEFAULT false,
    created_at  timestamptz NOT NULL DEFAULT now(),
    updated_at  timestamptz NOT NULL DEFAULT now()
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE UNIQUE INDEX IF NOT EXISTS uq_relationship_fact_active
    ON public.relationship_facts (owner_type, owner_id, key)
    WHERE valid_until IS NULL;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_relationship_fact_owner
    ON public.relationship_facts (owner_type, owner_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.relationship_facts;
-- +goose StatementEnd
