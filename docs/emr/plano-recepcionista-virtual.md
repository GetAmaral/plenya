# Plano — Recepcionista virtual com IA acoplado ao CRM

> Plano aprovado 2026-06-04. Cópia versionada do plano de implementação (regra: persistir plano em arquivo).

## Contexto

Os leads da Plenya chegam por WhatsApp (funil IG → site/form → WhatsApp no 43 99974-8899). Hoje o
CRM já tem IA pontual (sugestão de resposta manual, resumo de conversa, auto-confirmação de consulta
por intenção), mas **não existe um recepcionista virtual** que atenda. O guia/script da recepção +
banco de 12 objeções (`docs/atendimento/`) é a **base de conhecimento** que a IA usa para atender.

Objetivo: um recepcionista virtual **híbrido e intercambiável** acoplado ao CRM. Por conversa, o
atendente escolhe **Copiloto** (IA redige, humano envia) ou **Automático** (IA atende sozinha). O
modo automático **entra sozinho se nenhum humano responder após X minutos** (configurável), para não
deixar lead esperando (cobre noite/fim de semana naturalmente). O bot **se identifica** como
assistente virtual, segue os guardrails CFM/LGPD/marca, e **faz handoff** pro humano em sinal de
fechamento, dúvida clínica ou incerteza. Canal: **WhatsApp + copiloto no `/conversas`**. IG/site
widget ficam para depois (envio de IG não existe no backend; não há widget de site).

## Decisões fixadas (usuário, 2026-06-04)
- Híbrido **copiloto ⇄ automático**, alternável por conversa (e default global), ligado/desligado pelo atendente.
- Automático dispara **fallback por tempo**: se humano não responde em **X min** (configurável), a IA assume.
- Bot **se identifica** como assistente virtual na primeira mensagem automática.
- Canal WhatsApp + copiloto no CRM. IG/site = expansão futura.

## Arquitetura e reuso (o que já existe)
- **Geração de texto:** `AIService.CompleteText` (`services/ai_service.go`) e o padrão de
  `ConversationService.SuggestReply` / `SummarizeConversation` (`services/conversation_ai_service.go`).
- **Cérebro (knowledge):** system prompt versionado derivado do script + objeções
  (`docs/atendimento/script-recepcao-conversao-leads.md`), inline no prompt. Inclui R$ 800 fixo e os guardrails.
- **Envio:** `ConversationService.SendMessage`/`sendWhatsApp` (`services/conversation_service.go`) —
  já valida opt-in e a **janela de 24h** via `Lead.LastInboundAt`. Fora da janela, vira sugestão/handoff.
- **Gatilho inbound:** padrão `PatientInboundWAHook` (`lead_service.go SetPatientInboundWAHook` +
  `cmd/server/appointment_intent_hook.go`). Estender para cobrir **lead** também.
- **Scheduler:** molde `scheduler/appointment_reminder_job.go` para o job de fallback por tempo.
- **Handoff/notificação:** `NotificationService` (reuso).

## Fase 1 (executável agora) — Copiloto ancorado no script + cérebro
1. **Cérebro:** `services/reception_brain.go` com o system prompt (script condensado + objeções +
   guardrails + R$ 800 + voz da marca + regra de handoff).
2. **Geração ancorada:** `ConversationService.GenerateReceptionReply(ctx, ownerType, ownerID)` em
   `conversation_ai_service.go`, reusando `CompleteText`, saída estruturada
   `{reply, action: ask|answer|handle_objection|propose_schedule|handoff, handoffReason, discloseAI}`.
3. **Handler + rota:** `POST /conversations/:type/:id/ai/reception-reply` (junto das rotas AI em `main.go`).
4. **Frontend:** botão "Resposta da recepção" em `conversation-composer.tsx` → insere no compositor
   para revisar e enviar com 1 clique (humano sempre envia).

## Fase 2 — Modo automático + fallback timer + handoff + disclosure
1. Model `ConversationAutomation` (owner_type/owner_id XOR, `mode` off|copilot|auto, `fallback_minutes`,
   `paused_until`, `last_bot_at`, `updated_by`) + migration goose **`00017_conversation_automation.sql`**.
   Default global via config (`RECEPTION_BOT_DEFAULT_MODE`, `RECEPTION_BOT_FALLBACK_MINUTES`).
2. Usuário "Assistente Plenya" (seed) para `ActorUserID` das mensagens do bot.
3. Job `scheduler/conversation_auto_reply_job.go` (~1 min): conversas `mode=auto` com inbound sem
   resposta humana há ≥ fallback_minutes → `GenerateReceptionReply` → guardrails → `SendMessage`
   (janela 24h) → `LeadActivity{message_sent, metadata:{ai_generated:true}}` + `last_bot_at`. 1ª msg do bot identifica-se.
4. Handoff: `action=handoff`/sinais → `paused_until` + notifica equipe; "Assumir" no UI pausa o bot.
5. Frontend: toggle Off/Copiloto/Automático + `fallback_minutes` por conversa, default global em
   configurações, badge "respondido pela IA", botão "Assumir".
6. Anti-spam: limite de N mensagens do bot por conversa/hora.

## Fase 3 — Refinos (outline)
Propor horários reais (slots do calendar), métricas do bot, ajuste fino do cérebro, expansão IG/site.

## Guardrails do bot (lei)
Identifica-se como assistente; **nunca** diagnostica, interpreta exame ou promete resultado (CFM);
não pede dado clínico sensível (LGPD); **não fala preço do Continuum**; consulta é **R$ 800 fixo, não
negociável**; sem marcas/varejo, sem "medicina preditiva", sem maneirismos de IA nem travessão; voz
clínica conectiva PT-BR; **handoff imediato** em dúvida clínica ou sinal de fechamento; só envia
free-form dentro da janela de 24h.

## Verificação
- **Fase 1:** unit do prompt builder; Claude real em dev gera resposta ancorada (cita R$ 800, trata
  objeção, sem violar guardrails); Playwright no `/conversas`; `go build ./...` verde.
- **Fase 2:** migrate up/down 00017; teste do job; handoff pausa + notifica; toggle; rate-limit. QA + DB.
- **Guardrails:** bateria adversarial (diagnóstico, preço Continuum, "é robô?", negociar preço).

## Status de execução
- [x] **Fase 1 CONCLUÍDA (2026-06-04):** `reception_brain.go` (cérebro) + `GenerateReceptionReply`
  (saída estruturada reply/action/handoffReason/discloseAI) + handler `AIReceptionReply` + rota
  `POST /conversations/:type/:id/ai/reception-reply` + frontend (hook `useConversationReceptionReply`
  + botão "Recepção IA" no compositor, ambos canais). Verificado com Claude real em dev: resposta
  ancorada (segura R$ 800, sem desconto, sem preço do Continuum), **handoff** em dúvida clínica (não
  diagnostica/medica), **disclosure** na 1ª mensagem. `go build` verde, tsc web sem erro novo.
  Inclui **fix:** `aiModelSuggestion` era snapshot inválido `claude-sonnet-4-6-20251001` (404 na
  Anthropic) → alias `claude-sonnet-4-6` (conserta também o "Sugerir resposta" pré-existente) +
  `sanitizeReceptionVoice` remove travessões que o modelo insere.
- [x] **Fase 2 CONCLUÍDA (2026-06-04):** modo automático com fallback por tempo.
  - Model `ConversationAutomation` (owner_type/owner_id, mode off|copilot|auto, fallback_minutes,
    paused_until, last_bot_at) + migration **00017** (tabela + seed do usuário-bot "Assistente Plenya",
    UUID fixo `019e9301-0000-7000-a000-0000000b0b01` = `services.BotUserID`).
  - Config `ReceptionBotConfig` (env): `RECEPTION_BOT_ENABLED` (kill switch global, **default false**),
    `RECEPTION_BOT_DEFAULT_MODE`, `RECEPTION_BOT_FALLBACK_MINUTES` (5), `RECEPTION_BOT_MAX_MSGS_HOUR` (6).
  - `ConversationAutomationService` (Get efetivo/Set/Pause/TouchLastBot/ListAutoCandidates).
  - Job `scheduler/conversation_auto_reply_job.go` (ticker 1 min): elegibilidade (inbound sem resposta
    humana ≥ fallback, dentro da janela 24h, anti-spam por hora) → `GenerateReceptionReply` →
    `SendMessage` (bot user, `AIGenerated`) → idempotente. `action=handoff` → envia cortesia + pausa
    (paused_until) + notifica staff (bot excluído). Só roda com kill switch ligado.
  - `SendMessageInput.AIGenerated/AIModel` → metadata `ai_generated` na activity.
  - Endpoints `GET/PUT /conversations/:type/:id/automation`.
  - Frontend: toggle Off/Copiloto/Automático no header do viewer (`automation-toggle.tsx`), badge
    "respondido pela IA" nas mensagens, hooks `useConversationAutomation`/`useSetConversationAutomation`.
  - **FIX pré-existente (migration 00018):** `notifications` tinha 2 check constraints de `type` no
    baseline; a legada `notifications_type_check` (só tipos antigos) bloqueava silenciosamente TODA
    notificação de lead (lead_new/lead_assigned/handoff) — **inclusive em prod**. Removida; fica só
    `chk_notifications_type`.
  - **Verificado end-to-end em dev** (kill switch ligado temporariamente): auto-reply enviado pelo bot
    (ai_generated, idempotente), handoff em dúvida clínica (cortesia + pausa + 3 notificações ao staff),
    toggle GET/PUT. `go build` verde; tsc web sem erro novo. **Próxima migration = 00019.**
- [x] **Fase 3 (parcial) CONCLUÍDA (2026-06-04):** horários reais + métricas + guardrails.
  - **Horários reais:** `BuildUpcomingSlotsText` (`reception_slots.go`) busca a disponibilidade real
    via `CalendarSlotService.ListAvailable` (médico padrão = 1º user role doctor; consulta 60 min) e
    injeta no cérebro via `ConversationService.SetReceptionSlotsProvider`. O bot oferece horários
    concretos (action `propose_schedule`); quando o paciente escolhe, vira `handoff` com o horário no
    `handoffReason` (humano finaliza agendamento + pagamento; NÃO auto-booka). Sem working_hours → cai
    gracioso em "quer que eu veja uma data?".
  - **Métricas:** `ComputeMetrics` + `GET /conversations/ai/metrics?days=30` (auto-respostas, handoffs,
    conversas atendidas, convertidos após bot, por modo) + faixa `ReceptionMetricsBar` no topo do `/conversas`.
  - **Guardrails:** job pula lead `unsubscribed`/sem opt-in; handoff registrado como `LeadActivity`
    (`metadata.ai_handoff`, auditoria + métrica); cérebro não re-engaja após "PARAR".
  - **Verificado em dev:** ofereceu slots reais (seed de working_hours) → escolha → handoff com horário;
    métricas retornam JSON. `go build` verde; tsc web sem erro novo. **Nenhuma migration nova.**
  - **NÃO feito (fora de escopo desta rodada):** expansão IG/site (precisa envio IG no backend / widget).

## Revisão pré-deploy (2026-06-04) — achados corrigidos
Revisão adversarial de código + a11y antes do rollout. Corrigido (migration nova **00019**):
- **C1 (corrida de envio duplo):** o job re-checa `stillSendable` imediatamente antes de enviar
  (inbound ainda sem resposta + não pausada + modo auto). Evita a IA falar por cima de um humano
  que respondeu durante a chamada do Claude.
- **C2 (custo/latência):** `BuildUpcomingSlotsText` cacheia o texto de horários por médico
  (`receptionSlotsCacheTTL` 10 min). Antes fazia até 10 chamadas ao free/busy do Google por
  reception-reply (copiloto + cada candidato do job/min).
- **H1 (LGPD):** o job pula qualquer conversa cuja última mensagem seja stop-keyword
  (`IsUnsubscribeKeyword`), cobrindo **paciente** (que não tem flag de opt-out como o lead). Enforce,
  não confia só no prompt.
- **H2 (médico errado):** horários vêm de `RECEPTION_CONSULT_DOCTOR_ID` (quando setado e ainda doctor);
  fallback = 1º doctor por nome. Evita ofertar agenda do médico errado quando houver >1 doctor.
- **H3 (Claude em loop):** falha persistente de envio (não-handoff) agora pausa a conversa por
  `sendFailCooldown` (15 min), em vez de re-chamar o Claude a cada minuto.
- **M1 (privilégio do bot):** usuário-bot rebaixado de `secretary` p/ role própria `bot` (migration
  00019); não aparece mais em listas/notificações de staff. (Sem login: sem password_hash.)
- **M2/M3/M4:** limpeza de var morta no `Set`; métricas com `deleted_at` só onde a tabela tem
  (leads sim; lead_activities é log imutável sem a coluna); `sanitizeReceptionVoice` não mexe em
  en-dash colado a números (ex. faixas "18,5–24,9").
- **a11y:** `SheetContent` do viewer mobile (`/conversas`) ganhou `SheetTitle` (sr-only). Era a fonte
  do aviso "DialogContent requires DialogTitle". QA console: **zero erros** agora.
- Verificado em dev: copiloto OK (slots reais + cache, 0 travessões), job OK (happy path envia 1/2;
  lead "PARAR" pulado), métricas sem erro de coluna. `go build` verde; tsc web sem erro novo.

## Checklist de deploy em produção
1. `migrate up` (chega na **00019**: cria automation + bot user, fix notifications, rebaixa role do bot).
2. Envs no Coolify (app api):
   - `RECEPTION_BOT_ENABLED=true` (kill switch; deixe `false` p/ rollout só-copiloto).
   - `RECEPTION_CONSULT_DOCTOR_ID=<user id do Dr. Getúlio>` (recomendado; senão usa o 1º doctor).
   - opcional: `RECEPTION_BOT_FALLBACK_MINUTES` (default 5), `RECEPTION_BOT_MAX_MSGS_HOUR` (default 6).
3. **Mantenha `RECEPTION_BOT_DEFAULT_MODE=off`**: o job só processa conversas com **linha** `mode=auto`
   (criada pelo toggle no `/conversas`). Default global `auto` NÃO é auto-aplicado sem linha por conversa.
4. Rollout sugerido: começar só-copiloto (kill switch off), depois ligar `auto` por conversa via toggle.

## Como ligar em produção
1. Setar no Coolify (app api): `RECEPTION_BOT_ENABLED=true` (+ ajustar `RECEPTION_BOT_FALLBACK_MINUTES`
   se quiser ≠ 5). Rodar `migrate up` (chega na 00018).
2. Por conversa, o atendente escolhe Off/Copiloto/Automático no `/conversas`. Default global =
   `RECEPTION_BOT_DEFAULT_MODE` (off).
