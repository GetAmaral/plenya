-- Índice ÚNICO PARCIAL em source_session_id: garante no máximo 1 patient_score_snapshot vivo por
-- sessão anônima importada (fecha a corrida TOCTOU do import idempotente). Parcial porque
-- source_session_id é nulo nos snapshots normais do staff (e permite re-import após soft-delete).
-- Fase 4 (hardening da revisão adversarial).

-- +goose Up
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_snapshot_source_session;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE UNIQUE INDEX IF NOT EXISTS uq_snapshot_source_session
    ON public.patient_score_snapshots (source_session_id)
    WHERE source_session_id IS NOT NULL AND deleted_at IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS uq_snapshot_source_session;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_snapshot_source_session ON public.patient_score_snapshots (source_session_id);
-- +goose StatementEnd
