-- LeadActivity: corpo HTML original do e-mail (content_html). Content guarda o
-- texto plano (preview/busca/IA); content_html guarda o HTML pra render fiel no
-- viewer (iframe sandbox + sanitização no front). Cifrado em repouso igual a
-- content (canais email/whatsapp). Ver models/lead_activity.go.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.lead_activities ADD COLUMN IF NOT EXISTS content_html text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.lead_activities DROP COLUMN IF EXISTS content_html;
-- +goose StatementEnd
