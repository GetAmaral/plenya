-- Documentos clínicos emitidos: nova categoria "Orientações" (type=orientation)
-- e corpo rich-text (body_html, HTML sanitizado vindo do editor Word-style).
-- Ver issued_document.go / pdfdoc/issued_document.go.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.issued_documents DROP CONSTRAINT IF EXISTS chk_issued_documents_type;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.issued_documents
  ADD CONSTRAINT chk_issued_documents_type
  CHECK (type IN ('certificate','declaration','report','orientation'));
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.issued_documents ADD COLUMN IF NOT EXISTS body_html text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.issued_documents DROP COLUMN IF EXISTS body_html;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.issued_documents DROP CONSTRAINT IF EXISTS chk_issued_documents_type;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.issued_documents
  ADD CONSTRAINT chk_issued_documents_type
  CHECK (type IN ('certificate','declaration','report'));
-- +goose StatementEnd
