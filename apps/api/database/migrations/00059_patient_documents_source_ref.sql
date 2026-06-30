-- Liga um patient_document à entidade clínica que o materializou (ex.: "lab_request:<uuid>").
-- Usado pra idempotência: ao "Enviar por WhatsApp" um pedido de exames (que não nasce como
-- patient_document), materializamos uma cópia compartilhável UMA vez — o source_ref evita
-- duplicar a cada envio. NULL para uploads manuais / inbound (Postgres trata NULLs como
-- distintos no índice único, então não conflitam entre si).

-- +goose Up
ALTER TABLE patient_documents ADD COLUMN IF NOT EXISTS source_ref varchar(80);
CREATE UNIQUE INDEX IF NOT EXISTS uq_patient_documents_source_ref
    ON patient_documents (source_ref) WHERE source_ref IS NOT NULL;

-- +goose Down
DROP INDEX IF EXISTS uq_patient_documents_source_ref;
ALTER TABLE patient_documents DROP COLUMN IF EXISTS source_ref;
