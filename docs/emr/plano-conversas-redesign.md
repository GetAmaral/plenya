# Plano — `/conversas`: dock global de WhatsApp (responder de qualquer tela) + separar WhatsApp/E-mail + bucket Notificações

> Status: **APROVADO 2026-06-05**. Substitui o rascunho anterior. Fonte de contexto/bugs:
> [`whatsapp-go-live-runbook.md`](whatsapp-go-live-runbook.md).

## Contexto

Dois pedidos do usuário (2026-06-04/05):
1. A inbox unificada `/conversas` mistura **e-mail (assíncrono, assunto/threads)** e **WhatsApp
   (chat síncrono, janela 24h)** na mesma lista/timeline e **"não ficou bom junto"**. Separar os dois;
   WhatsApp vira chat de verdade.
2. **Principal:** a secretária precisa **responder WhatsApp de qualquer tela do EMR sem navegar pra
   outra página** — o equivalente a dar alt+tab pro WhatsApp do desktop, responder e voltar pro que
   estava fazendo (cadastro, agenda, etc.).

**Pesquisa de padrão (2026-06-05).** O padrão de mercado (Front, Intercom, HubSpot, GetStream) é um
**painel de chat persistente acoplado (docked slide-over)**: botão flutuante + atalho que desliza por
cima da tela atual, sem trocar de rota; responde e fecha, voltando ao contexto. Tecnicamente, no
Next 16/React 19 (doc oficial "Preserving UI state"): o painel **não** é uma página de rota (senão
desmonta e perde o rascunho na navegação); é montado no **layout autenticado compartilhado**
(`app/(authenticated)/layout.tsx`), que o Next preserva entre navegações, com estado em **store
Zustand global** (já usado no projeto). UI = **Sheet** ancorado à direita, aberto por FAB + atalho.

**Decisões já fixadas pelo usuário:**
- WhatsApp e E-mail em **rotas/itens de menu separados** (não abas numa página só).
- Criar **bucket "Notificações"** pras automáticas (hoje descartadas).
- **Dock = só WhatsApp** (síncrono/urgente). E-mail é assíncrono, fica só como página.

## Estado atual confirmado no código
- **Balões já existem** (`MessageBubble` + `WhatsAppMediaView`, `conversation-viewer.tsx:204,330`),
  auto-scroll, mídia WhatsApp nativa. Falta date-dividers + compositor compacto.
- Backend **já aceita `channel=email|whatsapp`** na lista (`conversation_handler.go:52`); e-mail e
  WhatsApp são a mesma tabela `LeadActivity` distinguida por coluna `channel`.
- **Bugs/lacunas:** (1) preview cifrado `v1:` — `decryptIfNeeded` (`conversation_service.go:773`)
  faz base64.Decode na string com prefixo `v1:` e falha; (2) endpoint de **mensagens não filtra por
  canal** (`Messages`/`GetMessagesParams`, `conversation_handler.go:220`/`conversation_service.go:801`);
  (3) `usePatientGuard()` indevido em `conversas/page.tsx:46`; (4) e-mails-bot descartados em
  `email_ingest_service.go:263` (vira o bucket).
- **Reuso:** `crypto.DecryptWithDefaultKey` (trata `vN:`), `conversation-list-item.tsx`,
  `lib/api/conversations-api.ts` (hooks de lista/mensagens/envio/AI/automação já com `channel`),
  Zustand (`lib/auth-store.ts` como molde de store), sidebar `collapsible-sidebar.tsx:100`.

---

## Fase 1 — Backend enabler + ganho visível imediato
1. **Fix `decryptIfNeeded`** (`services/conversation_service.go`): se começa com `v` e tem `:` nos
   primeiros chars, chamar `crypto.DecryptWithDefaultKey` direto (pular pré-check base64); heurística
   antiga como fallback. → preview legível na lista.
2. **Filtro de canal nas mensagens:** add `Channel` em `GetMessagesParams`, `WHERE channel = ?`
   quando preenchido; ler `c.Query("channel")` no handler.
3. **Não-lidas por canal:** dois hooks reusando a lista (`channel=…&unread_only=true`) p/ badges.
4. Build (`docker compose exec -w /app api go build ./...`) + deploy backend.

## Fase 2 — Componente de chat WhatsApp reutilizável + página cheia
A peça central, usada **tanto** pela página quanto pelo dock.
1. Extrair `components/conversations/whatsapp-chat.tsx` (lista + thread em balões + compositor
   compacto), recebendo `channel='whatsapp'`, `selected`, callbacks. **Date-dividers** (agrupar por
   dia), scroll colado no fim, compositor compacto (sem "Assunto"/dropzone de e-mail; 1 mídia),
   badge **janela 24h** + **Template** quando fechada. Controles avançados (Resumir, Agendar, Ver
   ficha, AutomationToggle) atrás de menu **"•••"**.
2. Página `app/(authenticated)/conversas/whatsapp/page.tsx` montando o componente em tela cheia +
   `reception-metrics-bar`. **Remover `usePatientGuard()`**.
3. Item de sidebar "WhatsApp" (→ `/conversas/whatsapp`, badge não-lidas WhatsApp).
4. QA Playwright (enviar/receber, janela 24h, mídia, •••).

## Fase 3 — Dock global de WhatsApp (responder de qualquer tela) ⭐
A entrega-chave do usuário.
1. **Store Zustand** `lib/whatsapp-dock-store.ts`: `{ isOpen, selected, openWith(ownerType,ownerId),
   toggle, close }`. Sobrevive à navegação (estado fora da árvore de rota).
2. **Montar no layout autenticado** `app/(authenticated)/layout.tsx` (fora das pages → persiste entre
   rotas): `<WhatsAppDock />` = **Sheet** ancorado à direita renderizando o **mesmo** `whatsapp-chat`
   (modo compacto). Rascunho/seleção preservados ao navegar e ao fechar/reabrir.
3. **Launcher:** FAB flutuante (canto inf. direito, com badge de não-lidas WhatsApp) **+ atalho**
   global (ex.: Ctrl/Cmd+J) que dá toggle no dock. `Esc` fecha. Visível só pra staff
   (admin/secretary/manager), respeitando RBAC do item atual.
4. **Cuidados (doc Next/React):** fechar dropdowns/popovers transitórios via cleanup; resetar rascunho
   ao trocar de conversa por `key={ownerId}`; não vazar z-index/estilos do Sheet. Playwright com
   `getByRole` (filtra invisíveis).
5. QA: abrir o dock em `/calendario` e `/patients/...`, responder, navegar pra outra rota → dock e
   rascunho intactos; FAB some/aparece certo; atalho funciona; badge atualiza.

## Fase 4 — Página de E-mail separada (webmail)
1. `app/(authenticated)/conversas/email/page.tsx` — lista + thread + compose **com assunto/anexos**,
   "Novo email". Reusa a casca de lista; viewer/compositor estilo webmail (mantém "Sugerir resposta"
   + resumo; **sem** automação/recepção/janela 24h, que são do WhatsApp).
2. `conversas/page.tsx` vira redirect → `/conversas/whatsapp` (preserva deep links).
3. Item de sidebar "E-mails" (badge não-lidas e-mail). Substitui o item único "Conversas".
4. QA Playwright.

## Fase 5 — Bucket "Notificações" (e-mails automáticos)
**Backend:**
1. Model `NotificationEmail` (`models/notification_email.go`): UUID v7 `BeforeCreate`; `From`,
   `Subject`, `Body` (cripto envelope, molde `LeadActivity` BeforeSave/AfterFind), `MessageID`,
   `ReceivedAt`, `IsRead`. Migration goose **`00018_notification_emails.sql`** (confirmar 00017 é a
   última).
2. `email_ingest_service.go:263`: no ramo `IsBot`, **gravar** `NotificationEmail` em vez de descartar.
3. Endpoints `GET /conversations/notifications` (+ unread) e `POST …/:id/read`, sob Auth+RBAC+AuditLog.
**Frontend:**
4. Abas **Caixa | Notificações** na página de e-mail; Notificações = lista read-only do novo endpoint.
5. QA: e-mail-bot novo cai em `notification_emails` (psql), não vira Lead; caixa só com reais; marcar lida.

---

## Guardrails / regras do projeto
Fonte única (Go models) + migration goose (AutoMigrate off; `migrate up`); UUID v7 via BeforeCreate;
cripto de conteúdo em repouso (envelope `vN:`); rotas sob Auth+RBAC+AuditLog; hooks frontend
obrigatórios; **commit direto no master** (nunca branch); sem manipular dado de dev por HTTP (psql
direto).

## Verificação (end-to-end)
- **F1:** build verde; lista mostra preview legível (sem `v1:`); `…/messages?channel=whatsapp` só WA.
- **F2:** página WhatsApp com date-dividers, compositor compacto, ••• com controles, janela 24h/template;
  sem barra de paciente. Playwright (localhost + prod).
- **F3 (chave):** dock abre por FAB e por atalho em qualquer tela; responder sem trocar de rota;
  navegar entre páginas mantém dock + rascunho; badge de não-lidas correto; `Esc`/fechar volta ao
  contexto. Playwright multi-rota.
- **F4:** sidebar com dois itens + badges; `/conversas` redireciona; e-mail estilo webmail.
- **F5:** `migrate up`/`down` da 00018; bucket recebe automáticas; caixa limpa; marcar lida.
