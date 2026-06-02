-- Documentos clínicos emitidos/assináveis (P3 frente 2): atestado, declaração, laudo.
-- Gerados no EMR (gofpdf), assinados com ICP-Brasil quando há certificado (degrada para
-- assinatura manual), publicados como PatientDocument (Type=...) no portal. Metadados de
-- assinatura ficam aqui para validação pública por QR. Aditivo.

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.issued_documents (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    appointment_id uuid,
    doctor_id uuid NOT NULL,
    type character varying(20) NOT NULL DEFAULT 'certificate',
    title character varying(200) NOT NULL DEFAULT '',
    body text NOT NULL DEFAULT '',
    purpose character varying(300),
    days_off integer,
    includes_cid boolean NOT NULL DEFAULT false,
    cid_code character varying(10),
    cid_consent boolean NOT NULL DEFAULT false,
    status character varying(10) NOT NULL DEFAULT 'draft',
    has_digital_signature boolean NOT NULL DEFAULT false,
    signed_at timestamp with time zone,
    signed_pdf_path character varying(500),
    signed_pdf_hash character varying(64),
    certificate_serial character varying(100),
    qr_code_data text,
    patient_document_id uuid,
    issued_by_user_id uuid NOT NULL,
    issued_at timestamp with time zone NOT NULL DEFAULT now(),
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone,
    CONSTRAINT issued_documents_pkey PRIMARY KEY (id),
    CONSTRAINT chk_issued_documents_type CHECK (type IN ('certificate','declaration','report')),
    CONSTRAINT chk_issued_documents_status CHECK (status IN ('draft','signed'))
);
CREATE INDEX IF NOT EXISTS idx_issued_documents_patient_id ON public.issued_documents (patient_id);
CREATE INDEX IF NOT EXISTS idx_issued_documents_appointment_id ON public.issued_documents (appointment_id);
CREATE INDEX IF NOT EXISTS idx_issued_documents_deleted_at ON public.issued_documents (deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.issued_documents;
-- +goose StatementEnd
