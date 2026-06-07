-- Lívia — rede de relacionamento + calendário de eventos/avisos (Fase C do plano
-- docs/emr/plano-livia-memoria-dossie-360.md §2.3/§2.4/§5). Só social (LGPD/CFM): pessoas
-- importantes da pessoa e datas/eventos para o time interagir proativamente. Nada clínico.
--   important_people  : rede (cônjuge, filhos, quem indicou…), com aniversário opcional.
--   relationship_events: aniversários, formaturas, nascimentos, follow-ups… com antecedência
--                        do aviso (lead_time_days) e status do ciclo.

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.important_people (
    id         uuid PRIMARY KEY,
    owner_type varchar(20) NOT NULL,
    owner_id   uuid        NOT NULL,
    name       varchar(160) NOT NULL,
    relation   varchar(60)  NOT NULL DEFAULT '',
    birthday   date,
    notes      text         NOT NULL DEFAULT '',
    source     varchar(20)  NOT NULL DEFAULT 'ai',
    added_by   uuid,
    created_at timestamptz  NOT NULL DEFAULT now(),
    updated_at timestamptz  NOT NULL DEFAULT now()
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_important_people_owner
    ON public.important_people (owner_type, owner_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.relationship_events (
    id                uuid PRIMARY KEY,
    owner_type        varchar(20) NOT NULL,
    owner_id          uuid        NOT NULL,
    related_person_id uuid,
    type              varchar(20) NOT NULL,
    title             varchar(200) NOT NULL,
    event_date        date        NOT NULL,
    recurring         boolean     NOT NULL DEFAULT false,
    lead_time_days    integer     NOT NULL DEFAULT 7,
    status            varchar(20) NOT NULL DEFAULT 'pending',
    source            varchar(20) NOT NULL DEFAULT 'ai',
    added_by          uuid,
    notes             text        NOT NULL DEFAULT '',
    created_at        timestamptz NOT NULL DEFAULT now(),
    updated_at        timestamptz NOT NULL DEFAULT now()
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_relationship_events_owner
    ON public.relationship_events (owner_type, owner_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_relationship_events_date
    ON public.relationship_events (event_date);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.relationship_events;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS public.important_people;
-- +goose StatementEnd
