-- Results inbox (P3 frente 1): revisão clínica de lotes de resultados laboratoriais.
-- Aditivo. reviewed_at/by = ciência do clínico; is_critical/worst_level = derivados do
-- lab_results.level (0=pior .. 5=melhor), backfilled abaixo. Não toca dados existentes além do backfill.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.lab_result_batches
    ADD COLUMN IF NOT EXISTS reviewed_at timestamp with time zone,
    ADD COLUMN IF NOT EXISTS reviewed_by_user_id uuid,
    ADD COLUMN IF NOT EXISTS is_critical boolean NOT NULL DEFAULT false,
    ADD COLUMN IF NOT EXISTS worst_level smallint;

-- Backfill: worst_level = menor level (pior) dos resultados classificados do lote.
UPDATE public.lab_result_batches b
SET worst_level = sub.min_level
FROM (
    SELECT lab_result_batch_id, MIN(level) AS min_level
    FROM public.lab_results
    WHERE level IS NOT NULL
    GROUP BY lab_result_batch_id
) sub
WHERE b.id = sub.lab_result_batch_id;

UPDATE public.lab_result_batches
SET is_critical = true
WHERE worst_level IN (0, 1);

CREATE INDEX IF NOT EXISTS idx_lab_batches_pending_review
    ON public.lab_result_batches (result_date)
    WHERE reviewed_at IS NULL AND deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_lab_batches_reviewed_by ON public.lab_result_batches (reviewed_by_user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS public.idx_lab_batches_pending_review;
DROP INDEX IF EXISTS public.idx_lab_batches_reviewed_by;
ALTER TABLE public.lab_result_batches
    DROP COLUMN IF EXISTS reviewed_at,
    DROP COLUMN IF EXISTS reviewed_by_user_id,
    DROP COLUMN IF EXISTS is_critical,
    DROP COLUMN IF EXISTS worst_level;
-- +goose StatementEnd
