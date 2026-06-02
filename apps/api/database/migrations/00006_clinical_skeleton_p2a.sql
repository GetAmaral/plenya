-- Esqueleto clínico P2a: alergias do paciente (base do CDS de prescrição) e
-- sinais vitais por consulta (distintos do physical_assessments do módulo treino).
-- Aditivo, sem tocar tabelas existentes.

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.patient_allergies (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    substance character varying(200) NOT NULL DEFAULT '',
    substance_type character varying(15) NOT NULL DEFAULT 'drug',
    reaction text,
    severity character varying(12) NOT NULL DEFAULT 'moderate',
    status character varying(10) NOT NULL DEFAULT 'active',
    no_known_allergies boolean NOT NULL DEFAULT false,
    notes text,
    recorded_by_user_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone,
    CONSTRAINT patient_allergies_pkey PRIMARY KEY (id),
    CONSTRAINT chk_patient_allergies_type CHECK (substance_type IN ('drug','food','environmental','other')),
    CONSTRAINT chk_patient_allergies_severity CHECK (severity IN ('mild','moderate','severe','anaphylaxis')),
    CONSTRAINT chk_patient_allergies_status CHECK (status IN ('active','inactive'))
);
CREATE INDEX IF NOT EXISTS idx_patient_allergies_patient_id ON public.patient_allergies (patient_id);
CREATE INDEX IF NOT EXISTS idx_patient_allergies_active ON public.patient_allergies (patient_id) WHERE status = 'active' AND deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_patient_allergies_deleted_at ON public.patient_allergies (deleted_at);

CREATE TABLE IF NOT EXISTS public.consultation_vitals (
    id uuid NOT NULL,
    appointment_id uuid,
    patient_id uuid NOT NULL,
    systolic_bp integer,
    diastolic_bp integer,
    heart_rate integer,
    resp_rate integer,
    temperature numeric(4,1),
    spo2 integer,
    weight numeric(5,2),
    height numeric(5,2),
    waist_circumference numeric(5,2),
    bmi numeric(4,1),
    measured_by_user_id uuid NOT NULL,
    measured_at timestamp with time zone NOT NULL DEFAULT now(),
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone,
    CONSTRAINT consultation_vitals_pkey PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS idx_consultation_vitals_patient_id ON public.consultation_vitals (patient_id);
CREATE INDEX IF NOT EXISTS idx_consultation_vitals_appointment_id ON public.consultation_vitals (appointment_id);
CREATE INDEX IF NOT EXISTS idx_consultation_vitals_deleted_at ON public.consultation_vitals (deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.consultation_vitals;
DROP TABLE IF EXISTS public.patient_allergies;
-- +goose StatementEnd
