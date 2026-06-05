-- Marca a origem do patient_score_snapshot: nulo = cálculo normal pelo staff; "anonymous_import"
-- = materializado a partir de uma sessão do Escore Plenya (Light/Triagem). source_session_id liga
-- de volta à sessão anônima (idempotência + rastreio). Fase 4 do plano score-versions-unificacao.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.patient_score_snapshots
    ADD COLUMN IF NOT EXISTS source            varchar(50),
    ADD COLUMN IF NOT EXISTS source_session_id uuid;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_snapshot_source ON public.patient_score_snapshots (source);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_snapshot_source_session ON public.patient_score_snapshots (source_session_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.patient_score_snapshots
    DROP COLUMN IF EXISTS source,
    DROP COLUMN IF EXISTS source_session_id;
-- +goose StatementEnd
