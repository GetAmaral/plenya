# Catálogo dinâmico de templates WhatsApp no EMR — plano (aprovado 2026-06-26)

Objetivo: registro estruturado dos templates WhatsApp DENTRO do EMR — status sincronizado da Meta,
variáveis, propósito, cabeamento e toggle — com UI admin. Substitui doc/hardcode como "registro".

Fonte da verdade do conteúdo/status = Meta. Cacheamos + sincronizamos. O mapeamento de params por
fluxo continua no código (regra de negócio).

## Backend
1. Migration `00049_whatsapp_templates.sql` — tabela `whatsapp_templates`:
   - da Meta (sync): `meta_id`, `name`, `language`, `category`, `status`, `body_text`, `variable_count`, `last_synced_at`
   - nossos (preservados no sync): `purpose`, `variables` jsonb `[{index,label,example}]`, `wiring_notes`, `enabled` (bool, default true)
   - chave: unique (`name`,`language`)
2. Model `WhatsAppTemplate` (UUID v7 BeforeCreate).
3. Serviço `WhatsAppTemplateService`:
   - `SyncFromMeta(ctx)` — GET …/message_templates, upsert por (name,language); preserva colunas nossas.
   - `List()`, `GetByName(name,lang)`, `UpdateMeta(id, fields)`.
4. Config: `WHATSAPP_WABA_ID` (default `827642893502160`).
5. Guard de envio: `WhatsAppService.SendTemplate` consulta o cache; se status≠APPROVED ou enabled=false → recusa com erro claro (não envia pendente/rejeitado em silêncio). Injetado por interface (sem ciclo).
6. Cron: sync diário (scheduler) + refresh sob demanda.
7. Handlers admin (RBAC admin/manager):
   - `GET /admin/whatsapp-templates`
   - `POST /admin/whatsapp-templates/sync`
   - `PATCH /admin/whatsapp-templates/:id` (purpose/variables/wiring_notes/enabled)

## Frontend
- Página admin `configuracoes/whatsapp-templates`: tabela (nome, status badge, categoria, nº vars,
  propósito, enabled toggle, last_synced), botão "Sincronizar", edição dos metadados nossos.
- Entrada no nav de configurações.

## Notas
- `lead_activity.TemplateSnapshot` = histórico por mensagem (auditoria), complementa, não substitui.
- 1º sync popula 13 templates atuais (incl. `documento_disponivel` PENDING).
