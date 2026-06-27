-- Flag de curadoria: o template pode ser escolhido manualmente pelo atendente no compositor de
-- conversas? (false p/ internos/automáticos: lead_alert, magic_link, prep_invite, documento, teste).
-- Independe de enabled (que é "o sistema pode enviar"). Ver services/whatsapp_template_service.go.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.whatsapp_templates ADD COLUMN IF NOT EXISTS chat_selectable boolean NOT NULL DEFAULT true;
-- Curadoria inicial: internos/automáticos não são escolhíveis à mão no chat.
UPDATE public.whatsapp_templates SET chat_selectable = false
 WHERE name IN ('lead_alert','magic_link','consultation_prep_invite','documento_disponivel','hello_world');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.whatsapp_templates DROP COLUMN IF EXISTS chat_selectable;
-- +goose StatementEnd
