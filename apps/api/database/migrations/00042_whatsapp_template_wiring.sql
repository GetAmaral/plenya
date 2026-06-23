-- Cabeamento dos templates WhatsApp aprovados (ciclo de consulta + follow-up + reengajamento).
-- Colunas de idempotência por estágio. Ver docs/emr/whatsapp-templates-wiring.md.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.appointments
  ADD COLUMN IF NOT EXISTS dayof_reminder_sent_at timestamptz,
  ADD COLUMN IF NOT EXISTS followup_sent_at       timestamptz;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.leads
  ADD COLUMN IF NOT EXISTS reengaged_at timestamptz;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.appointments
  DROP COLUMN IF EXISTS dayof_reminder_sent_at,
  DROP COLUMN IF EXISTS followup_sent_at;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.leads
  DROP COLUMN IF EXISTS reengaged_at;
-- +goose StatementEnd
