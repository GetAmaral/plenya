-- Backfill da origem dos resultados existentes: marca como 'pdf' os que pertencem a
-- lotes que têm um ProcessingJob com PDF (a 00044 deu default 'manual' a todos).
-- Ver docs/emr/plano-importacao-laudos-melhorias.md (Fase 3, revisão M4).

-- +goose Up
-- +goose StatementBegin
UPDATE public.lab_results r
SET source = 'pdf'
WHERE source = 'manual'
  AND EXISTS (
    SELECT 1 FROM public.processing_jobs j
    WHERE j.lab_result_batch_id = r.lab_result_batch_id
      AND j.pdf_path <> ''
  );
-- +goose StatementEnd

-- +goose StatementBegin
UPDATE public.lab_results
SET match_reason = 'Não encontrado no catálogo de exames'
WHERE source = 'pdf' AND matched = false AND match_reason IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Backfill de dados: sem reversão automática (no-op).
SELECT 1;
-- +goose StatementEnd
