-- Conserta drift do baseline: a tabela notifications tem DUAS check constraints de `type`.
-- chk_notifications_type (boa, inclui lead_new/lead_assigned/etc.) e notifications_type_check
-- (legada, só tipos antigos). A legada bloqueia silenciosamente TODA notificação de lead
-- (lead_new no inbound, lead_assigned, handoff da recepção IA, intent cancel/reschedule).
-- Remove a legada; a boa permanece como única fonte. Forward-only.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.notifications DROP CONSTRAINT IF EXISTS notifications_type_check;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 1;
-- +goose StatementEnd
