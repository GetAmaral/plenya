-- Respostas individuais de escalas somadas (PHQ-9, GAD-7, Epworth, IIEF-5, …) preenchidas
-- pelo widget pergunta-a-pergunta no EMR. O nível classificado continua em selected_level
-- (fonte do motor de score); esta coluna guarda o detalhe por pergunta para exibição/histórico.
-- Plano: docs/emr/plano-escalas-pergunta-a-pergunta.md (F3).

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.anamnesis_items
  ADD COLUMN IF NOT EXISTS scale_responses jsonb;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.anamnesis_items DROP COLUMN IF EXISTS scale_responses;
-- +goose StatementEnd
