-- Simplifica a nota de evolução: aposenta SOAP/APSO (4 seções + toggle) por 2 campos livres —
-- "História clínica e evolução" (clinical_history) + "Conduta" (conduct). SOAP não é exigido por
-- lei (CFM Res. 1.638/2002 exige CONTEÚDO: anamnese, exame, hipótese, conduta, evolução — não o
-- formato). Sistema em beta, sem notas reais a preservar → troca limpa, sem migração de dados.
-- Estudo: docs/emr/estudo-simplificacao-consulta-anamnese.md.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.clinical_notes ADD COLUMN IF NOT EXISTS clinical_history text;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.clinical_notes ADD COLUMN IF NOT EXISTS clinical_history_html text;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.clinical_notes ADD COLUMN IF NOT EXISTS conduct text;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.clinical_notes ADD COLUMN IF NOT EXISTS conduct_html text;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.clinical_notes
  DROP COLUMN IF EXISTS layout,
  DROP COLUMN IF EXISTS subjective,
  DROP COLUMN IF EXISTS subjective_html,
  DROP COLUMN IF EXISTS objective,
  DROP COLUMN IF EXISTS objective_html,
  DROP COLUMN IF EXISTS assessment,
  DROP COLUMN IF EXISTS assessment_html,
  DROP COLUMN IF EXISTS plan,
  DROP COLUMN IF EXISTS plan_html;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.clinical_notes
  ADD COLUMN IF NOT EXISTS layout varchar(4) NOT NULL DEFAULT 'soap',
  ADD COLUMN IF NOT EXISTS subjective text,
  ADD COLUMN IF NOT EXISTS subjective_html text,
  ADD COLUMN IF NOT EXISTS objective text,
  ADD COLUMN IF NOT EXISTS objective_html text,
  ADD COLUMN IF NOT EXISTS assessment text,
  ADD COLUMN IF NOT EXISTS assessment_html text,
  ADD COLUMN IF NOT EXISTS plan text,
  ADD COLUMN IF NOT EXISTS plan_html text;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.clinical_notes
  DROP COLUMN IF EXISTS clinical_history,
  DROP COLUMN IF EXISTS clinical_history_html,
  DROP COLUMN IF EXISTS conduct,
  DROP COLUMN IF EXISTS conduct_html;
-- +goose StatementEnd
