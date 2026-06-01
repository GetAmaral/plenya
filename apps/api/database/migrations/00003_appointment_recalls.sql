-- Fila interna de retornos previstos (recall) da recepção. A secretária trabalha
-- a fila manualmente (contata o paciente e agenda); sem aviso automático nesta
-- fase. Artefato operacional do balcão, não clínico.

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.appointment_recalls (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    doctor_id uuid,
    source_appointment_id uuid,
    due_date date NOT NULL,
    reason text,
    notes text,
    status character varying(20) NOT NULL DEFAULT 'pending',
    scheduled_appointment_id uuid,
    created_by_user_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone,
    CONSTRAINT appointment_recalls_pkey PRIMARY KEY (id),
    CONSTRAINT chk_appointment_recalls_status CHECK (status IN ('pending','scheduled','dismissed'))
);
CREATE INDEX IF NOT EXISTS idx_appointment_recalls_patient_id ON public.appointment_recalls (patient_id);
CREATE INDEX IF NOT EXISTS idx_appointment_recalls_doctor_id ON public.appointment_recalls (doctor_id);
CREATE INDEX IF NOT EXISTS idx_appointment_recalls_due_date ON public.appointment_recalls (due_date);
CREATE INDEX IF NOT EXISTS idx_appointment_recalls_status ON public.appointment_recalls (status);
CREATE INDEX IF NOT EXISTS idx_appointment_recalls_deleted_at ON public.appointment_recalls (deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.appointment_recalls;
-- +goose StatementEnd
