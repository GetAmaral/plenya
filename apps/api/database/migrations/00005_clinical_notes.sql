-- Nota de evolução clínica por consulta (SOAP/APSO). Distinta da anamnese
-- (intake inicial): é a evolução de cada visita, ancorada ao appointment.
-- Imutável após assinada (status='signed') — correção só via adendo
-- (amendment_of_id). Reusa o enum de visibilidade da anamnese.

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.clinical_notes (
    id uuid NOT NULL,
    appointment_id uuid,
    patient_id uuid NOT NULL,
    author_id uuid NOT NULL,
    layout character varying(4) NOT NULL DEFAULT 'soap',
    subjective text,
    subjective_html text,
    objective text,
    objective_html text,
    assessment text,
    assessment_html text,
    plan text,
    plan_html text,
    status character varying(10) NOT NULL DEFAULT 'draft',
    signed_at timestamp with time zone,
    amendment_of_id uuid,
    visibility character varying(20) NOT NULL DEFAULT 'all',
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone,
    CONSTRAINT clinical_notes_pkey PRIMARY KEY (id),
    CONSTRAINT chk_clinical_notes_layout CHECK (layout IN ('soap','apso')),
    CONSTRAINT chk_clinical_notes_status CHECK (status IN ('draft','signed')),
    CONSTRAINT chk_clinical_notes_visibility CHECK (visibility IN ('all','medicalOnly','psychOnly','authorOnly'))
);
CREATE INDEX IF NOT EXISTS idx_clinical_notes_appointment_id ON public.clinical_notes (appointment_id);
CREATE INDEX IF NOT EXISTS idx_clinical_notes_patient_id ON public.clinical_notes (patient_id);
CREATE INDEX IF NOT EXISTS idx_clinical_notes_author_id ON public.clinical_notes (author_id);
CREATE INDEX IF NOT EXISTS idx_clinical_notes_amendment_of_id ON public.clinical_notes (amendment_of_id);
CREATE INDEX IF NOT EXISTS idx_clinical_notes_deleted_at ON public.clinical_notes (deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.clinical_notes;
-- +goose StatementEnd
