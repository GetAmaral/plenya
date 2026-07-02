-- Blindagem de whatsapp_templates.variables: deve ser SEMPRE um array JSON.
--
-- Uma linha gravada como objeto `{...}` (insert manual — o sync/PATCH sempre gravam slice→array)
-- fazia o `serializer:json` do GORM (campo `Variables []WhatsAppTemplateVar`) abortar a leitura da
-- tabela INTEIRA: List/ListSendable retornavam 500 (tela de seleção de templates zerada) e o cron
-- `[wa-template-sync]` travava (o próprio First() na linha ruim estoura, então os status também
-- paravam de sincronizar). Ver services/whatsapp_template_service.go e models/whatsapp_template.go.
--
-- Normaliza legados (objeto→array de 1 elemento) e adiciona CHECK: dá pra falhar no WRITE (loud)
-- em vez de corromper o READ (silent). Idempotente.

-- +goose Up
-- +goose StatementBegin
UPDATE public.whatsapp_templates
   SET variables = jsonb_build_array(variables)
 WHERE jsonb_typeof(variables) = 'object';

DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint WHERE conname = 'whatsapp_templates_variables_is_array'
  ) THEN
    ALTER TABLE public.whatsapp_templates
      ADD CONSTRAINT whatsapp_templates_variables_is_array
      CHECK (jsonb_typeof(variables) = 'array');
  END IF;
END $$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.whatsapp_templates
  DROP CONSTRAINT IF EXISTS whatsapp_templates_variables_is_array;
-- +goose StatementEnd
