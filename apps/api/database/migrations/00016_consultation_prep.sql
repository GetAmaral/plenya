-- Preparação pré-consulta (Fase 2): formulário curado pós-agendamento + respostas.
-- Espelha o Escore Light (anonymous_score_session/item), porém vinculado ao Patient.
-- Aditivo e idempotente: score_items.prep_order (flag de curadoria) + 2 tabelas novas.
-- Ver docs/emr/plano-preparacao-pre-consulta.md e docs/atendimento/fluxo-pre-consulta.md.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.score_items
    ADD COLUMN IF NOT EXISTS prep_order integer;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_score_item_prep_order ON public.score_items USING btree (prep_order);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.consultation_preps (
    id                uuid PRIMARY KEY,
    patient_id        uuid NOT NULL REFERENCES public.patients(id) ON DELETE CASCADE,
    appointment_id    uuid,
    status            varchar(20) NOT NULL DEFAULT 'draft',
    chief_complaint   text,
    submitted_at      timestamptz,
    consent_version   varchar(20),
    consent_timestamp timestamptz,
    created_at        timestamptz NOT NULL DEFAULT now(),
    updated_at        timestamptz NOT NULL DEFAULT now(),
    deleted_at        timestamptz,
    CONSTRAINT chk_consultation_preps_status CHECK (status IN ('draft','submitted'))
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_consultation_prep_patient ON public.consultation_preps USING btree (patient_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_consultation_prep_appointment ON public.consultation_preps USING btree (appointment_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_consultation_preps_deleted_at ON public.consultation_preps USING btree (deleted_at);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.consultation_prep_responses (
    id            uuid PRIMARY KEY,
    prep_id       uuid NOT NULL REFERENCES public.consultation_preps(id) ON DELETE CASCADE,
    score_item_id uuid NOT NULL REFERENCES public.score_items(id) ON DELETE RESTRICT,
    numeric_value double precision,
    text_value    text,
    selected_level integer,
    created_at    timestamptz NOT NULL DEFAULT now(),
    updated_at    timestamptz NOT NULL DEFAULT now(),
    deleted_at    timestamptz,
    CONSTRAINT chk_consultation_prep_responses_level CHECK (selected_level IS NULL OR (selected_level >= 0 AND selected_level <= 6))
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_consultation_prep_response_prep ON public.consultation_prep_responses USING btree (prep_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_consultation_prep_response_item ON public.consultation_prep_responses USING btree (score_item_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_consultation_prep_responses_deleted_at ON public.consultation_prep_responses USING btree (deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.consultation_prep_responses;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS public.consultation_preps;
-- +goose StatementEnd
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_score_item_prep_order;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.score_items DROP COLUMN IF EXISTS prep_order;
-- +goose StatementEnd
