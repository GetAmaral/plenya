# Atendimento IA — plano FECHADO (2026-06-06)

> Nome: **Atendimento IA**. Estados: **Off · Copiloto · Auto**. **Quem manda é o GLOBAL**;
> conversas individuais sobrepõem temporariamente (flag de 24h).
> Mapeamento p/ enum existente (não mudar valores armazenados): Off=`off`, Copiloto=`copilot`, Auto=`auto`.

## Modos (comportamento)

- **Off:** IA silenciosa. Se chegar inbound e ficar **> Z** sem resposta **e** a conversa estiver Off
  *por herança do global* (sem flag Off explícita) → **alerta o admin** (que pode entrar e ligar o Auto).
- **Copiloto:** após o debounce **X**, a IA **gera o rascunho** (idêntico ao que o Auto enviaria) e
  **preenche o campo de mensagem** (quando a conversa está aberta na tela). **Espera o humano enviar.**
  Se ficar **> Y** sem resposta ao lead → **escala p/ Auto** (regera e envia).
- **Auto:** após o debounce **X**, gera a resposta. Se a conversa estiver **aberta** na tela, preenche o
  campo e mostra um **countdown de X/3** pro atendente ver/intervir; se ninguém mexer, **envia** ao fim do
  X/3. Se a conversa **não** estiver aberta, **envia direto** após X.

> A IA só age quando há um **inbound do lead aguardando resposta** — nunca fala sozinha sem o lead ter
> mandado algo. Handoff (dúvida clínica / pedido de humano): **não envia**, alerta (canal do Z) e tira a
> conversa do Auto até alguém agir.

## Três tempos configuráveis

| | O quê | Default | Observação |
|---|---|---|---|
| **X** | **debounce** + atraso da resposta do Auto. Conta do **último inbound do lead**; **reseta a cada nova msg** da rajada. | **30 s** | usado também no Copiloto (não draftar no meio da rajada). Em Auto + tela aberta, há +X/3 de preview antes do envio. |
| **Y** | sem resposta no Copiloto por Y → escala p/ Auto (regera e envia) | **10 min** | vale p/ Copiloto via global ou via flag |
| **Z** | sem resposta no Off por Z → alerta o admin | **20 min** | só p/ Off **por herança** (flag Off explícita NÃO alerta) |

## Janela de horário do Auto (tela de config, estilo Google)

- **Parametrizável por dia da semana** com faixas de horário (ex.: 18:00→08:00 seg–sex; dia todo no fds).
- **Dentro da janela → Auto.** **Fora da janela → vale o global** (o baseline que o usuário setou).
- **Opção de desligar** a janela (aí vale o global o tempo todo).
- Timezone: **America/Sao_Paulo**.

## Máquina de estados por conversa

- **Conversa nova, sem flag:** segue o **global** (que pode estar elevado a Auto pela janela de horário).
- **Usuário escolhe um modo na conversa** (ex.: global=Auto, marca Copiloto) → vale o override.
- **Auto + humano digita e envia** (ou mexe no campo durante o countdown X/3) → vira **Copiloto** ("assumiu").
- **Copiloto + > Y sem resposta ao lead** → escala p/ **Auto**.
- **Flag Off explícita** → **não** dispara o alerta Z (Z é só p/ conversas novas / sem flag / Off herdado).
- **Validade da flag: 24h após a última mensagem (qualquer direção)** → conversa perde a flag e volta ao
  global. (Alinha com a janela de 24h do WhatsApp.)
- **Hierarquia:** flag de conversa válida **vence** o global.

## Configuração / permissões
- Quem configura o global + tempos X/Y/Z + janela de horário: **admin, secretary, manager, doctor**.

## Arquitetura (build)
- **Rascunho no servidor (Copiloto E Auto):** o rascunho nasce no backend (debounced por X) e fica
  **salvo na conversa** (`suggested_reply` + timestamp + modelo). Mostrado no compositor quando a conversa
  está aberta; badge "rascunho pronto" na lista. No Auto, vira envio (com preview X/3 se aberto).
- **Job** (evolui o `conversation_auto_reply_job`): orquestra X (debounce→draft/preview/send), Y
  (copiloto→auto), Z (off→alerta), janela de horário, expiração de flag 24h, handoff→alerta.
- **Storage:** setting global runtime singleton (modo baseline + janela semanal + X/Y/Z) + auditoria;
  flag por conversa com timestamp (expiração 24h) e origem (herdada vs explícita, p/ a regra do Z).
  Migrations goose.
- **UI:** controle global "Atendimento IA: Off/Copiloto/Auto" + **tela de config** (janela semanal +
  tempos) + **banner de estado** sempre visível ("Atendimento IA: Auto até 08:00"). Preview X/3 com
  opção de assumir/cancelar no compositor.
- **Guardrails mantidos:** kill switch `RECEPTION_BOT_ENABLED`, opt-out LGPD, anti-spam/hora, janela 24h
  (free-form só dentro; fora = template).

## Rótulos (UI) vs enum (storage)
Display: **Manual · Copiloto · Auto**. Enum armazenado (não mudar): `off`=Manual, `copilot`, `auto`.

## Progresso

### ✅ Parte 1/3 — backend data layer + controle global (commit `bae103fd`)
- migration `00028_atendimento_ia.sql`: `reception_settings` (singleton) + `conversation_suggested_replies`.
- models `ReceptionSettings` (+ `WeeklySchedule`/`ParsedSchedule`) e `ConversationSuggestedReply`.
- `ReceptionSettingsService` (Get/Update + janela cruza meia-noite TZ America/Sao_Paulo + EffectiveGlobalMode/AutoActiveNow).
- endpoints **`GET/PUT /api/v1/reception/settings`** (admin/secretary/manager/doctor).

### ✅ Parte 2/3 — job da máquina de estados (commit `158005dc`)
- `conversation_auto_reply_job` reescrito (tick 15s): resolveMode (flag ≤24h → global c/ janela),
  expireStaleFlags (sweep 24h), Off+Z→alerta, Copiloto (draft após X; escala no Y), Auto (draft após X;
  envia após X+X/3), handoff→pausa+alerta, anti-spam, stillSendable. Gated por `RECEPTION_BOT_ENABLED`.
- ⚠️ ainda não exercitado end-to-end em runtime (precisa ligar kill switch + WhatsApp; fazer junto com o front).

### ✅ Parte 3/3 — frontend (CÓDIGO FEITO — falta só `pnpm generate` opcional + deploy)
Config numa **tela dedicada** (`/configuracoes/atendimento-ia`), não dialog. Todos os 7 passos:
1. ✅ **Backend:** `GET /conversations/:type/:id/suggested-reply` — lê o rascunho salvo + `effectiveMode`
   (resolve override→global c/ janela, via `autoSvc.ResolveEffectiveMode`). 204 quando não há rascunho.
2. ✅ **Hooks TanStack:** `useReceptionSettings` (GET/PUT) + `useUpdateReceptionSettings` + `useSuggestedReply`
   (poll 8s, normaliza 204→null) em `conversations-api.ts`.
3. ✅ **Controle global** `AtendimentoIAGlobalBar` (Manual·Copiloto·Auto) no topo do `/conversas/whatsapp`.
4. ✅ **Banner de estado** embutido na mesma barra ("Auto até HH:MM" quando a janela eleva; replica
   `isAutoActiveAt` no cliente p/ achar o fim da janela) + atalho "Configurar".
5. ✅ **Tela dedicada de config** `/configuracoes/atendimento-ia`: modo baseline + grade semanal (faixas/dia,
   cruza meia-noite, "dia todo", "aplicar exemplo") + X/Y/Z. PUT único do schedule inteiro.
6. ✅ **Preview X/3 no compositor:** `AIDraftBar` no viewer — injeta o rascunho no campo (reusa draftBody/token),
   countdown ao vivo até `updatedAt + X/3` no Auto + ações "Assumir" (→copilot) e "Enviar agora". Copiloto/handoff
   mostram aviso sem countdown.
7. ✅ **Relabel** do `AutomationToggle` por conversa → Manual/Copiloto/Auto. ⬜ `pnpm generate` (opcional — front
   usa tipos manuais) + ⬜ **deploy** (api + web) junto com a otimização `a8f65208`.

⚠️ Falta: **deploy de prod** (api+web via Coolify; 00028 sobe sozinha com `RUN_MIGRATIONS=true`) + ligar
`RECEPTION_BOT_ENABLED` + **teste end-to-end em runtime** (job nunca exercitado). Confirmar janela com o Getúlio.

## 🧭 Retomada após /compact
Ler este doc + `git log --oneline` (procurar `bae103fd`/`158005dc`). Backend pronto e commitado; começar
pela **Parte 3a**. Arquivos-chave: API `apps/web/lib/api/conversations-api.ts`; toggle por conversa
`apps/web/components/conversations/automation-toggle.tsx`; tela `apps/web/app/(authenticated)/conversas/`;
compositor `conversation-composer.tsx` (já recebe rascunho via `draftBody`/`draftToken`). Endpoints prontos:
`GET/PUT /api/v1/reception/settings`. Falta o endpoint do passo 1. Mapeamento Manual=`off`.

## Status
Partes 1, 2 e 3 **concluídas, commitadas e DEPLOYADAS em prod** (2026-06-06).
- Front+back no commit `9e35144a`; api+web rebuildados no Coolify (api `kgcuxgvmnbx6yya35e3ca2v0`,
  web `nwbhak0fscs2th13gz5g9zjm`); migration **00028 aplicada** (RUN_MIGRATIONS=true + prod-entrypoint
  `set -e` → server saudável prova que subiu). Rotas `/reception/settings` e `/suggested-reply` no ar (401 sem auth).
- **`RECEPTION_BOT_ENABLED=false`** setado explícito no app da API → **bot OFF**: UI visível, nada
  rascunhado/enviado sozinho. `pnpm generate` pulado (front usa tipos manuais).
- ⬜ **Pendente:** teste end-to-end em runtime (ligar o kill switch num teste controlado) + confirmar
  janela com o Getúlio antes de ativar Copiloto/Auto no número real.
