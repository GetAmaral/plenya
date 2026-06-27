-- Metadados por resultado de exame (material/espécime, método, faixa de referência do lab,
-- e data de coleta POR EXAME). Tudo opcional. O material/método já vinham da extração do PDF
-- mas eram descartados; collection_date permite renderizar a tabela por data do resultado
-- (com fallback p/ a data do lote). Ver docs/emr/plano-metadados-labtest.md.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.lab_results
  ADD COLUMN IF NOT EXISTS specimen        varchar(50),
  ADD COLUMN IF NOT EXISTS method          varchar(150),
  ADD COLUMN IF NOT EXISTS reference_range text,
  ADD COLUMN IF NOT EXISTS collection_date timestamptz;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.lab_results
  DROP COLUMN IF EXISTS specimen,
  DROP COLUMN IF EXISTS method,
  DROP COLUMN IF EXISTS reference_range,
  DROP COLUMN IF EXISTS collection_date;
-- +goose StatementEnd
