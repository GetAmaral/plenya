-- Entrega da prescrição por download autenticado (PatientDocument) em vez de /uploads estático.
-- O servir estático de /uploads já foi removido (vazava PDFs sem auth); a prescrição passa a
-- ser publicada como PatientDocument e baixada via endpoint autenticado, como o IssuedDocument.
-- Aditivo: prescriptions.patient_document_id + amplia o CHECK de patient_documents.type p/ 'prescription'.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.prescriptions
    ADD COLUMN IF NOT EXISTS patient_document_id uuid;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.patient_documents DROP CONSTRAINT IF EXISTS chk_patient_documents_type;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.patient_documents
    ADD CONSTRAINT chk_patient_documents_type
    CHECK (type IN ('certificate','report','referral','declaration','other','prescription'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.prescriptions DROP COLUMN IF EXISTS patient_document_id;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.patient_documents DROP CONSTRAINT IF EXISTS chk_patient_documents_type;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.patient_documents
    ADD CONSTRAINT chk_patient_documents_type
    CHECK (type IN ('certificate','report','referral','declaration','other'));
-- +goose StatementEnd
