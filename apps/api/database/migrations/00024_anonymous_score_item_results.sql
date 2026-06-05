-- Per-item results das sessões anônimas (Escore Light/Triagem), espelhando o mínimo de
-- patient_score_item_results. Habilita o radar por PILAR (buildAgir) no resultado público,
-- via Item.MethodPillars. Fase 3 do plano docs/emr/plano-score-versions-unificacao.md.

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.anonymous_score_item_results (
    id               uuid PRIMARY KEY,
    snapshot_id      uuid NOT NULL,
    item_id          uuid NOT NULL,
    group_id         uuid NOT NULL,
    status           varchar(30) NOT NULL CHECK (status IN ('evaluated','not_applicable','no_data_available')),
    level_number     integer CHECK (level_number >= 0 AND level_number <= 6),
    level_matched_id uuid,
    value_used       double precision,
    max_points       double precision NOT NULL DEFAULT 0,
    actual_points    double precision NOT NULL DEFAULT 0,
    created_at       timestamptz NOT NULL DEFAULT now(),
    updated_at       timestamptz NOT NULL DEFAULT now(),
    deleted_at       timestamptz,
    CONSTRAINT fk_anon_item_result_snapshot FOREIGN KEY (snapshot_id)
        REFERENCES public.anonymous_score_snapshots(id) ON DELETE CASCADE
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_anon_item_result_snapshot ON public.anonymous_score_item_results (snapshot_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_anon_item_result_item ON public.anonymous_score_item_results (item_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_anon_item_result_group ON public.anonymous_score_item_results (group_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_anon_item_result_deleted_at ON public.anonymous_score_item_results (deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.anonymous_score_item_results;
-- +goose StatementEnd
