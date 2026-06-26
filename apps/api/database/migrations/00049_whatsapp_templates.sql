-- Catálogo de templates WhatsApp no EMR: cache + status sincronizado da Meta + metadados nossos
-- (propósito, variáveis, cabeamento, toggle enabled). Fonte da verdade do conteúdo/status = Meta;
-- aqui é cache sincronizável + o que é nosso. Ver docs/emr/whatsapp-templates-catalogo-plano.md.

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.whatsapp_templates (
    id              uuid PRIMARY KEY,
    meta_id         varchar(64),
    name            varchar(120) NOT NULL,
    language        varchar(12)  NOT NULL DEFAULT 'pt_BR',
    category        varchar(20),
    status          varchar(20)  NOT NULL DEFAULT 'UNKNOWN',
    body_text       text,
    variable_count  int          NOT NULL DEFAULT 0,
    variables       jsonb        NOT NULL DEFAULT '[]',
    purpose         text,
    wiring_notes    text,
    enabled         boolean      NOT NULL DEFAULT true,
    last_synced_at  timestamptz,
    created_at      timestamptz  NOT NULL DEFAULT now(),
    updated_at      timestamptz  NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX IF NOT EXISTS uq_whatsapp_templates_name_lang ON public.whatsapp_templates (name, language);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.whatsapp_templates;
-- +goose StatementEnd
