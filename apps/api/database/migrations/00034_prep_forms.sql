-- Formulários pré-consulta (A1/B1/B2) reaproveitando ScoreVersion como entidade do form.
--   score_versions.context      → separa versões públicas (triagem/light) das de prep do paciente.
--   consultation_preps.score_version_id → qual form o paciente respondeu.
--   appointments.prep_form_version_id   → qual form a consulta usa (escolhido ao agendar).
--   appointments.prep_reminder_{48h,24h}_sent_at → idempotência dos lembretes de preparação.
-- Plano: docs/emr (formulários pré-consulta). score_items.prep_order fica dormente (remoção futura).

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.score_versions
    ADD COLUMN IF NOT EXISTS context varchar(20) NOT NULL DEFAULT 'public'
    CHECK (context IN ('public','patient_prep'));
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.consultation_preps
    ADD COLUMN IF NOT EXISTS score_version_id uuid REFERENCES public.score_versions(id) ON DELETE SET NULL;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.appointments
    ADD COLUMN IF NOT EXISTS prep_form_version_id uuid REFERENCES public.score_versions(id) ON DELETE SET NULL,
    ADD COLUMN IF NOT EXISTS prep_reminder_48h_sent_at timestamptz,
    ADD COLUMN IF NOT EXISTS prep_reminder_24h_sent_at timestamptz;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_consultation_prep_version ON public.consultation_preps(score_version_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.appointments
    DROP COLUMN IF EXISTS prep_form_version_id,
    DROP COLUMN IF EXISTS prep_reminder_48h_sent_at,
    DROP COLUMN IF EXISTS prep_reminder_24h_sent_at;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.consultation_preps DROP COLUMN IF EXISTS score_version_id;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.score_versions DROP COLUMN IF EXISTS context;
-- +goose StatementEnd
