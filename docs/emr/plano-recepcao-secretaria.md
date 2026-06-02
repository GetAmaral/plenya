# Plano — UX da Secretaria (Recepção) no EMR Plenya

> **Proveniência:** plano aprovado via plan mode (ExitPlanMode) na sessão de 2026-05-30
> (transcript `c36529d5`). Recuperado e persistido em arquivo em 2026-05-30 após a conversa
> de implementação ter sido limpa por um erro de API. Este é o documento canônico do recurso —
> consultar e manter atualizado o **status** ao final.

## RESOLUÇÃO `/hoje` vs `/recepcao` (2026-05-30)

**Decisão: `/recepcao` é a página única da recepção. `/hoje` foi removida.**

### Por que houve a confusão (assumido)
A tarefa nasceu certa (estudo de casos de uso da secretaria → plano → implementar), mas a
implementação atravessou **duas execuções que não compartilharam estado**: a sessão original
(`c36529d5`) construiu `/hoje` (monolito) e adicionou em `calendar-api.ts` os hooks de check-in/
start + `waitMinutes` + labels; um workflow posterior refatorou para `/recepcao` (cards limpos),
repontou o sidebar, **mas reverteu aqueles hooks** e **não portou os fluxos de balcão**, deixando
`/hoje` órfã e quebrada. Resultado: duas páginas com o mesmo propósito, uma delas derrubando o
build. Raiz: plano vivia na memória da conversa e o trabalho se espalhou por sessões/workflows
sem reconciliação (lição em `persistir_plano_em_arquivo`).

### Estado de cada página (verificado)
- **`/recepcao`** (171 linhas, no sidebar, **builda verde**): cockpit limpo em 8 cards
  (`components/recepcao/*`). Cobre **Feature 1 (Dashboard)**: agenda do dia (todos os médicos),
  a confirmar (1-clique), mensagens não lidas, leads novos, busca inline de paciente, ações
  rápidas. NÃO cobre ainda: check-in/iniciar/concluir, pagamento, lista de espera, Cmd+K.
- **`/hoje`** (379 linhas, órfã, **quebrava o build**): monolito que tinha os fluxos de balcão
  wired, mas importava de `calendar-api.ts` símbolos que o refactor reverteu (`useCheckInAppointment`,
  `useStartAppointment`, `waitMinutes`, `APPOINTMENT_STATUS_COLORS`). Preservada como referência
  em `docs/emr/ref-hoje-balcao-page.tsx.txt` para reaproveitar a fiação na Fase 2.

### Fase 1 — finalizar `/hoje` (FEITO nesta data)
1. Remover a rota `/hoje` (movida para referência). 2. Corrigir o título `Recepcao` → `Recepção`
em `/recepcao`. 3. Build do web verde. Confusão eliminada, uma página só.

### Fase 2 — ✅ CONCLUÍDA E VERIFICADA (2026-05-30) — balcão completo no `/recepcao`
Implementado sobre o `/recepcao` (sem páginas novas), build web + go verdes. Itens 1-8 feitos:
fundação recuperada no `calendar-api.ts` (hooks/type/LABELS/COLORS/`waitMinutes`); check-in/iniciar/
concluir + tempo de espera + pagamento (RegisterPaymentDialog) em `appointment-row.tsx`; card de lista
de espera (`lista-espera-card.tsx` novo) + cadastro rápido (QuickAddPatientDialog) em `page.tsx`;
`<GlobalSearch>` Cmd+K montado em `app/(authenticated)/layout.tsx`; status maps de `recepcao-helpers.ts`
completados; toDTO de Patient devolve BirthDate vazio quando zerada; `/hoje` e a referência removidos.
**Pendente:** smoke visual pelo usuário (não há browser no ambiente) + deploy (parte de "o resto").
Detalhe original abaixo (mantido como registro):
1. **Recuperar de `c36529d5` (L1943-1947)** e reaplicar em `calendar-api.ts`: type `AppointmentStatus`
   com `checked_in`/`in_progress`, `APPOINTMENT_STATUS_LABELS/COLORS`, hooks `useCheckInAppointment`/
   `useStartAppointment`, helper `waitMinutes`. (Usar `ref-hoje-balcao-page.tsx.txt` como guia.)
2. **Feature 2 (check-in):** botões Chegou/Iniciar/Concluir + tempo de espera no `appointment-row.tsx`/
   `agenda-hoje-card.tsx`.
3. **Feature 6 (financeiro):** `RegisterPaymentDialog` na linha da consulta e nas ações rápidas.
4. **Feature 5 (lista de espera):** card de waitlist na coluna lateral + UI de adicionar (`useCreateWaitlistEntry`).
5. **Feature 3 (cadastro rápido):** trocar o "Novo paciente" (rota) pelo `QuickAddPatientDialog` (nome+telefone, encadeia pra Nova Consulta).
6. **Feature 4 (Cmd+K):** montar `<GlobalSearch>` no layout autenticado.
7. Resolver o `Patient.BirthDate not null` (cadastro rápido grava `0001-01-01`).
8. Ao fim, apagar `ref-hoje-balcao-page.tsx.txt`.

---

## Context

O papel `secretary` já existe no RBAC e tem acesso operacional (agenda, conversas, leads,
pacientes, campanhas). O que falta é a **camada de balcão**: hoje a secretária improvisa usando o
Calendário como agenda do dia, não há fluxo de chegada do paciente, cadastro rápido, busca global,
nem qualquer registro financeiro. Este plano transforma os 6 casos de uso levantados em uma
entrega única que dá à recepção um fluxo de trabalho coeso.

Decisões já tomadas com o usuário:
- **Check-in**: 2 status novos (`checked_in`, `in_progress`) + carimbos de horário (chegada e
  início) para medir tempo de espera.
- **Financeiro**: construir do zero — registro de pagamento de consulta avulsa, valor por tipo de
  consulta e emissão de recibo.
- **Entrega**: plano único, implementar tudo de uma vez e revisar o conjunto no fim.

Regras invariantes que valem aqui: Go models são fonte única (editar model → `pnpm generate`);
sem Go local (compilar no container); voz de marca sem travessões nos textos visíveis (recibo,
labels); secretária nunca deleta (admin-only). UUID v7 via `BeforeCreate` em todo model novo.

---

## Feature 1 — Dashboard "Hoje" (home da recepção)

Nova página que consolida o dia em uma tela só, eliminando o pula-pula entre Calendário/Conversas/Leads.

**Frontend (somente):**
- Nova rota `apps/web/app/(authenticated)/hoje/page.tsx`.
- Reaproveita hooks existentes: `useAppointments({ dateFrom: início-do-dia, dateTo: fim-do-dia })`
  de `lib/api/calendar-api.ts`; `useConversationsUnreadCount()` e `useConversations({ unreadOnly:true, limit:10 })`
  de `lib/api/conversations-api.ts`; `useLeads({ status:'new', limit:5 })` de `lib/api/leads-api.ts`.
- Layout: coluna esquerda = **Fila de hoje** (agenda do dia, todos os médicos, ordenada por hora,
  com badge de status e ações de check-in/iniciar — ver Feature 2); coluna direita = cards
  "Não confirmados" (filtra status `scheduled` de hoje/amanhã com botão Confirmar 1-clique via
  `useConfirmAppointment`), "Mensagens não lidas" e "Leads novos".
- Item no sidebar `apps/web/components/layout/collapsible-sidebar.tsx`:
  `{ name:"Hoje", href:"/hoje", icon:Home, staffOnly:true, requiredRoles:['secretary','admin','manager'] }`,
  posicionado no topo. Padrão de role-gate já existe (linhas ~57-69, 257-260).
- Reusar `APPOINTMENT_STATUS_LABELS`/`_COLORS` de `calendar-api.ts` para badges.

Sem mudança de backend (a query "hoje, todos os médicos, com nomes" já é suportada pelo
`List` com `DateFrom`/`DateTo` + `Preload("Patient")`).

---

## Feature 2 — Fluxo de check-in (chegou → em atendimento → concluído)

### Backend
- **Model** `apps/api/internal/models/appointment.go`:
  - Adicionar constantes: `AppointmentCheckedIn = "checked_in"` (aguardando) e
    `AppointmentInProgress = "in_progress"` (em atendimento).
  - Atualizar a tag `check:` do campo `Status` para incluir os 2 valores novos.
  - Adicionar campos `CheckedInAt *time.Time` e `StartedAt *time.Time` (`type:timestamptz`).
- **Migração de constraint** `apps/api/internal/database/database.go`: a tag `check:` do GORM **não
  altera** um CHECK já existente. Adicionar bloco pré-AutoMigrate (padrão dos blocos DO já no
  arquivo) que dropa `chk_appointments_status` e recria com os 5+2 valores. A EXCLUDE
  anti-overlap (linha ~415, exclui só `cancelled`/`no_show`) **não muda** — `checked_in`/
  `in_progress` devem continuar bloqueando conflito de horário.
- **Service** `apps/api/internal/services/appointment_service.go`, seguindo o padrão de `Confirm`:
  - `CheckIn(ctx, apptID, actorUserID) (*dto.AppointmentResponse, error)`: transição
    `scheduled|confirmed → checked_in`, seta `CheckedInAt=now`. Idempotente.
  - `StartConsultation(ctx, apptID, actorUserID) (*dto.AppointmentResponse, error)`: transição
    `checked_in|confirmed|scheduled → in_progress`, seta `StartedAt=now`.
  - `Complete` continua via `Update` com `status=completed` (já existe); opcionalmente expor
    `/complete` por simetria.
- **Handler** `apps/api/internal/handlers/appointment_handler.go`: `CheckIn` e `StartConsultation`
  (padrão fino do handler atual).
- **Rotas** `apps/api/cmd/server/main.go` (bloco `/appointments`, ~691-706):
  `POST /:id/check-in` e `POST /:id/start` — herdam `RequireAnyStaff` (recepção e médico podem).
- **DTO** `apps/api/internal/dto/`: adicionar `checkedInAt`/`startedAt` ao `AppointmentResponse`.

### Frontend
- `lib/api/calendar-api.ts`: mutations `useCheckInAppointment(id)` e `useStartAppointment(id)`
  (padrão de `useConfirmAppointment`, invalidando `calendarKeys.all` + `.appointment(id)`).
- Badges para `checked_in` ("Aguardando") e `in_progress` ("Em atendimento") em
  `appointments/page.tsx` (objeto `statusConfig`) e nos labels/cores de `calendar-api.ts`.
- Botões de ação: na fila do Dashboard Hoje, no drawer do Calendário
  (`components/calendar/calendar-grid.tsx` → Sheet) e em `appointments/[id]/page.tsx`.
- **Tempo de espera**: exibir `StartedAt − CheckedInAt` na fila e na ficha; base para um relatório
  de pontualidade simples (lista do dia com média de espera) na própria página Hoje.

---

## Feature 3 — Cadastro rápido de paciente

Dialog leve para cadastrar no meio da ligação com o mínimo (Nome + Telefone), resto depois.

- **Backend**: nenhum. O `CreatePatientRequest` exige só `name`, `birthDate`, `gender`. Para
  permitir cadastro só com nome+telefone, tornar `birthDate`/`gender` opcionais no DTO **ou**
  (preferido, menos risco) manter obrigatórios mas com defaults na UI rápida. **Decisão de
  implementação**: relaxar `birthDate` e `gender` para opcionais em `dto/patient.go` +
  `patient_service.go` (parse condicional da data; `gender` default `other` se ausente). Isso é a
  mudança mínima e não quebra o cadastro completo.
- **Frontend**: componente `components/patients/quick-add-patient-dialog.tsx` usando `Dialog` +
  `react-hook-form` + `useFormNavigation({ formRef })`. Campos: Nome (req), Telefone, e
  opcionalmente "completar depois". Reutilizável a partir de: botão no Dashboard Hoje, no topo da
  tela de Nova Consulta (`appointments/new/page.tsx`) e na lista de pacientes
  (`patients/page.tsx`). Ao salvar, opção de já selecionar o paciente (`setSelectedPatient`) e
  encadear para Nova Consulta.

---

## Feature 4 — Busca global (Cmd+K) por nome / CPF / telefone

- **Backend** `apps/api/internal/handlers/patient_handler.go` + `patient_service.go`: adicionar
  param `?search=` no `List` de pacientes (hoje não existe; só há lookup exato de CPF por
  `cpf_blind_index`). Lógica: `ILIKE` em `name` e `phone`; se o termo for 11 dígitos, calcular o
  blind index e fazer match exato de CPF (CPF é criptografado, não dá `ILIKE`). Phone é texto puro
  (normalizado E.164) → `ILIKE` direto.
- **Frontend**: componente `components/global-search.tsx` com `Command` (cmdk já é dependência,
  `components/ui/command.tsx`) dentro de `Dialog`, atalho `Cmd/Ctrl+K` global (montar no layout
  autenticado). Debounce 300ms → `GET /api/v1/patients?search=`. Resultado navega para
  `/patients/[id]` ou oferece "selecionar paciente". Atalho disponível de qualquer tela.

---

## Feature 5 — Reagendar (drag-and-drop), filtro "não confirmados", lista de espera

- **Drag-to-reschedule** no calendário: `@dnd-kit/core` já instalado (usado em `methods` e
  `lab-results`). Tornar os blocos de `components/calendar/calendar-grid.tsx` arrastáveis; no drop,
  calcular novo `scheduledAt` a partir da posição e chamar `useUpdateAppointment(id, { scheduledAt })`.
  A EXCLUDE constraint do banco já protege contra sobreposição (erro tratado com toast).
- **Filtro "não confirmados"**: na tela `appointments/page.tsx` e no Dashboard Hoje, filtro/toggle
  passando `status:'scheduled'` para `useAppointments`. UI no padrão dos filtros já existentes.
- **Lista de espera** (encaixe se abrir vaga): feature menor e nova.
  - Model novo `apps/api/internal/models/waitlist_entry.go` (UUID v7): `PatientID`, `DoctorID *`,
    `PreferredType`, `Notes`, `Status` (`waiting`/`scheduled`/`cancelled`), `CreatedByUserID`.
  - Service/handler/rotas CRUD sob `RequireAnyStaff`; DTO próprio.
  - Frontend: aba/painel simples na página Hoje ou em Consultas para adicionar/converter em
    agendamento. (Sem notificação automática nesta fase; conversão é manual pela secretária.)

---

## Feature 6 — Financeiro de recepção (pagamento + recibo) — maior bloco, tudo novo

### Backend
- **Model preço** `apps/api/internal/models/consultation_price.go` (UUID v7): `Type`
  (AppointmentType), `Amount` (em centavos, `int64`), `Active bool`. Admin gerencia; demais leem.
- **Model pagamento** `apps/api/internal/models/appointment_payment.go` (UUID v7):
  `PatientID` (req), `AppointmentID *uuid` (opcional — avulsa pode não ter consulta), `Amount`
  (centavos), `Method` (`cash`/`pix`/`card`/`transfer`/`other`), `Status` (`paid`/`refunded`),
  `PaidAt`, `ReceiptNumber` (sequencial por ano, ex. `2026-000123`), `Notes`, `CreatedByUserID`.
  Hooks: `BeforeCreate` (UUID v7 + geração de `ReceiptNumber` sequencial transacional).
- **Service** `services/payment_service.go`: `Create` (registra pagamento, resolve valor sugerido
  do `ConsultationPrice` pelo tipo), `List` (por paciente / período), `Refund`, `GenerateReceipt`.
- **Recibo**: gerar HTML→PDF reaproveitando o padrão `AssessmentHTMLService`
  (`services/assessment_html_service.go`) — recibo com nº sequencial, dados da clínica (NAP),
  paciente, valor, método, data. Texto sem travessões (voz de marca).
- **Handlers/rotas** `cmd/server/main.go`:
  - `/payments` (CRUD + `GET /:id/receipt` retornando PDF) sob `RequireAdminOps`
    (admin/manager/secretary).
  - `/consultation-prices` (CRUD) sob `RequireAdmin`; `GET` liberado a `RequireAdminOps` para a
    recepção ver o valor sugerido.
- **DTO** correspondentes em `dto/`.

### Frontend
- `lib/api/payments.ts`: hooks de pagamento e preços (padrão TanStack Query do projeto).
- Na ficha do paciente e no Dashboard Hoje: "Registrar pagamento" (Dialog) com valor pré-preenchido
  pelo tipo de consulta, método, e botão "Emitir recibo" (abre/baixa o PDF).
- (Opcional, decidir depois) visão read-only de status de assinatura na ficha — fora do mínimo.

---

## Ordem de implementação

1. **Backend primeiro** (models → migração de constraint → services → handlers → rotas → DTOs),
   numa leva, para rodar `pnpm generate` uma vez só ao fim do backend.
2. `pnpm generate` (atlas migration + swagger + tipos TS + zod) — regra de ouro.
3. **Frontend** por feature, reusando os hooks/clients citados.
4. Build + validação local antes de commit.

Arquivos backend principais: `models/appointment.go`, `models/waitlist_entry.go` (novo),
`models/consultation_price.go` (novo), `models/appointment_payment.go` (novo),
`database/database.go` (bloco de constraint), `services/appointment_service.go`,
`services/payment_service.go` (novo), `handlers/appointment_handler.go`,
`handlers/patient_handler.go`, `services/patient_service.go`, `dto/appointment.go`,
`dto/patient.go`, `cmd/server/main.go`.

Arquivos frontend principais: `app/(authenticated)/hoje/page.tsx` (novo),
`components/layout/collapsible-sidebar.tsx`, `lib/api/calendar-api.ts`,
`components/calendar/calendar-grid.tsx`, `app/(authenticated)/appointments/page.tsx`,
`app/(authenticated)/appointments/[id]/page.tsx`, `components/patients/quick-add-patient-dialog.tsx`
(novo), `components/global-search.tsx` (novo), `lib/api/payments.ts` (novo),
`app/(authenticated)/patients/page.tsx`.

---

## Verificação (end-to-end)

- **Compilar backend**: `docker compose exec -w /app api go build ./...`.
- **Migração**: confirmar no banco que `chk_appointments_status` aceita os novos valores —
  `docker compose exec db psql -U plenya_user -d plenya_db -c "\d+ appointments"`; e que a EXCLUDE
  anti-overlap segue bloqueando `checked_in`/`in_progress`.
- **Fluxo check-in**: criar consulta → check-in (vira `checked_in`, `CheckedInAt` setado) →
  iniciar (`in_progress`, `StartedAt`) → concluir; conferir tempo de espera calculado.
- **Cadastro rápido**: criar paciente só com nome+telefone via Dialog; confirmar persistência.
- **Busca global**: `Cmd+K`, buscar por nome, por telefone parcial, e por CPF de 11 dígitos
  (match exato via blind index).
- **Reagendar**: arrastar bloco no calendário; tentar sobrepor outro agendamento do mesmo médico e
  ver o toast de conflito (constraint).
- **Financeiro**: cadastrar preço por tipo (admin), registrar pagamento (secretária), emitir recibo
  PDF com nº sequencial; conferir que recibos têm numeração contínua.
- **RBAC**: logar como `secretary` e confirmar acesso a Hoje, check-in, busca, pagamentos; e
  bloqueio em consultation-prices CRUD e deletes.
- **Geração**: `pnpm generate` sem erros; tipos novos aparecem em `packages/types`.

Implementação direto no `master` (regra do projeto), commit/deploy só quando solicitado.

---

## ✅ STATUS FINAL — CONCLUÍDO (2026-06-01, prod @14287d43)

**A feature Recepção/Secretária está completa, deployada e verificada.** As 6 features originais +
todos os itens não-excluídos do P0/P1/P2 da reavaliação estão presentes, ligados ponta a ponta e
corretos. Verificação final por workflow adversarial de 9 agentes (7 verificadores por grupo +
crítico de completude + síntese): veredito **GO**, zero blockers.

- **6 features originais:** Dashboard `/recepcao` (cards reais + RBAC), check-in Chegou/Iniciar/
  Concluir + tempo de espera (BE idempotente com guarda de transição), cadastro rápido + `/patients/new`,
  GlobalSearch Cmd+K (montado 1× no layout autenticado), waitlist (card + dialog), financeiro
  (RegisterPaymentDialog + recibo PDF com ReceiptNumber sequencial atômico + CNPJ real).
- **P0 (4 bugs):** cadastro consertado (aposentou `lib/api/patients.ts`), `isError` distinto do
  empty-state nos cards, 409 de conflito mapeado no reagendamento, polling 30s em leads/waitlist.
- **P1 (8):** handoff lead→agenda (via conversão de lead → `?patientId`), RBAC secretária em
  working_hours, telas financeiras (PatientFinanceCard: histórico+reimprimir+estorno gateado),
  dedupe CPF + dedupe pagamento, busca server-side, badge de lead + SLA na inbox, respostas rápidas.
- **P2:** no_show ("Faltou"), encaixe ativo (nudge interno), handoff no viewer, admin de preços
  (`/configuracoes/precos`), badge "Pago R$X", SLA-sort, canned no EmailReplyBlock, busca unaccent,
  recall de retorno (stack completo, **interno**, migration 00003), slot-click, **drag-and-drop**
  (@dnd-kit: mouse/touch/teclado + DragOverlay + update otimista).
- **Excluídos (confirmado AUSENTES e sem código morto):** ingestão IG/FB DMs, win-back, e
  **qualquer mensageria automática ao paciente** (régua 48h/24h/2h). Recall, encaixe e no_show são
  **internos** — verificado que NÃO disparam mensagem automática ao paciente.
- **Hardening de sign-off (commit desta data):** estorno restrito a admin/gerente também no backend
  (`POST /payments/:id/refund` ganhou `RequireRole(admin, manager)` — antes só a UI escondia o botão,
  a secretária ainda podia estornar via API) + guarda anti-duplo-estorno (`ErrPaymentAlreadyRefunded`).
- **Integridade:** `go build ./...` exit 0; tsc web = 812 (baseline, 0 erro novo); migrations
  00001-00004 aplicadas (goose); `HEAD == origin/master == prod`.

**Follow-ups aceitos (não-bloqueantes, fora do escopo desta entrega):** QuickAdd sem dedupe de
telefone (cadastro mínimo por design); recall fica `pending` após "Agendar" (falta auto-linkar o
`scheduled_appointment_id` — backend já suporta); picker de `appointments/new` ainda carrega
pacientes client-side (busca server-side só foi pra lista `/patients`); janela de horas 08-20
hardcoded; pré-intake/dashboard do gestor/NFS-e (P3).

---

## STATUS DA IMPLEMENTAÇÃO (histórico — pré-conclusão, mantido como registro)

> Reconstruído em 2026-05-30 lendo os transcripts das duas sessões anteriores (`c36529d5` =
> sessão 1, e a sessão de refactor seguinte) + verificação direta no codebase e no git. Corrige
> uma primeira auditoria que tinha errado dois pontos ("pnpm generate nunca rodou" e "GlobalSearch
> nunca montado"). **Conteúdo abaixo é o snapshot da época (`/hoje` vs `/recepcao`, drag-drop
> ausente, etc.) — SUPERSEDIDO pelo STATUS FINAL acima.** Mantido pela cronologia das 2 sessões.

### O que realmente aconteceu (cronologia das 2 sessões)

A implementação passou por **duas sessões**, e a segunda reverteu trabalho da primeira. Entender
isso é o que muda a decisão de retomada.

- **Sessão 1 (`c36529d5`)** implementou o plano fielmente: backend completo (compila), e no frontend
  o `calendar-api.ts` ganhou o type `AppointmentStatus` estendido com `checked_in`/`in_progress`,
  os Records LABELS/COLORS, os hooks `useCheckInAppointment`/`useStartAppointment` e o helper
  `waitMinutes`; criou `/hoje/page.tsx` (cockpit monolítico com check-in/pagamento/cadastro
  rápido/filtro não-confirmados wired), pôs **"Hoje"→/hoje** no sidebar e **montou `<GlobalSearch>`
  no `layout.tsx`**. (Tentou `pnpm generate`, que erra em atlas/swag, mas isso é inócuo: o schema
  vem do AutoMigrate.) Travou no erro de API `thinking blocks cannot be modified`.
- **Sessão 2 (refactor)** começou a trocar o monolito `/hoje` por uma arquitetura de cards limpos
  (`components/recepcao/*`): criou os 8 cards + `/recepcao/page.tsx`, repontou o sidebar para
  **"Recepção"→/recepcao**. No processo, **reverteu os arquivos versionados editados pela sessão 1**:
  `calendar-api.ts` voltou a +2 linhas vs HEAD (só `checkedInAt`/`startedAt` na interface; o type,
  labels, hooks e `waitMinutes` sumiram), `layout.tsx` voltou ao HEAD (sem GlobalSearch). **Não
  terminou** (cards não cobrem check-in/pagamento/waitlist) e travou no mesmo erro de API. Foi essa
  sessão que o usuário limpou.

**Implicação central:** `/hoje` não está quebrada por estar "pela metade" — ela está quebrada
porque o refactor da sessão 2 **reverteu de `calendar-api.ts` exatamente os hooks/type/`waitMinutes`
de que ela depende**. Esse código não foi perdido: está no transcript da sessão 1 (`c36529d5`,
edições em torno de L1943-1947) e pode ser recuperado em vez de reescrito.

### Bloqueador único para retomar

**Uma página de recepção pela metade (refactor inacabado).** `/recepcao` (cards, é a do sidebar,
funciona, mas sem balcão) vs `/hoje` (monolito com balcão, órfã, quebrada porque suas deps em
`calendar-api.ts` foram revertidas). Precisa de UMA decisão (ver "Decisão de retomada").

> **NÃO é bloqueador: `pnpm generate` / Atlas / swag.** O fluxo real do projeto NÃO usa Atlas:
> o schema é criado pelo **GORM AutoMigrate** no startup (`database.go:63`), as migrations `.sql`
> são escritas à mão quando precisa, e o frontend desta feature usa **tipos próprios inline** em
> `lib/api/payments.ts`/`waitlist.ts`/`calendar-api.ts` (não importa `@plenya/types`). Portanto
> `atlas`/`swag` não estarem instalados é irrelevante aqui. (Os docs `CLAUDE.md`/`03-architecture.md`/
> `02-stack.md` documentam Atlas como pipeline oficial, mas isso diverge da prática — ver
> [[plenya_atlas_nao_e_usado]].)

### Decisão de retomada (pendente do usuário)

Independente do caminho, **a primeira coisa é recuperar de `c36529d5` o trecho de `calendar-api.ts`**
(type `AppointmentStatus` + LABELS/COLORS + hooks `useCheckInAppointment`/`useStartAppointment` +
`waitMinutes`), porque os dois caminhos precisam dele e ele já estava pronto.

- **Caminho A — terminar `/recepcao` (arquitetura melhor, é a do menu):** portar os fluxos de
  balcão de `/hoje` para os cards, remontar `<GlobalSearch>`, depois deletar `/hoje`. Mais
  alinhado ao destino do refactor, porém re-trabalha o que a sessão 1 já tinha em `/hoje`.
- **Caminho B — restaurar `/hoje` (menos re-trabalho):** recuperar o `calendar-api.ts` e o
  `layout.tsx` (GlobalSearch) da sessão 1, repontar o sidebar para `/hoje`, descartar `/recepcao`
  e os cards. Recupera o máximo de código pronto, mas mantém o monolito de 379 linhas.

### Status por feature

| Feature | Backend | Frontend | Geral | O que falta (1 linha) |
|---|---|---|---|---|
| 1. Dashboard recepção | ✅ completo (sem mudança de BE) | 🟡 parcial | 🟡 quase | `/recepcao` (linkada) não tem balcão; `/hoje` (com balcão) órfã e quebrada — unificar |
| 2. Check-in (5+2 status) | ✅ completo | 🟡 parcial | 🟡 quase | **recuperar** de `c36529d5` os hooks `useCheckIn`/`useStart`/`waitMinutes` + type/LABELS (sessão 2 reverteu) |
| 3. Cadastro rápido paciente | ✅ completo | 🟡 parcial | 🟡 quase | wire dos call-sites (appointments/new, patients/page); resolver `BirthDate not null` no model |
| 4. Busca global Cmd+K | ✅ completo | 🟡 parcial | 🟡 quase | **remontar** `<GlobalSearch>` no layout (sessão 1 montou, sessão 2 reverteu — recuperável L2019) |
| 5. Reagendar + não-confirmados + waitlist | ✅ completo (waitlist BE) | 🟡 parcial | 🟡 parcial | drag-and-drop ausente do calendar-grid; sem UI de "adicionar à lista de espera" |
| 6. Financeiro (pagamento + recibo) | ✅ completo | ✅ completo | ✅ completo | nada (BE compila, FE com tipos próprios, AutoMigrate cria as tabelas) |

### O que está PRONTO

**Backend (compila, exit 0)**
- **Check-in:** constantes `AppointmentCheckedIn`/`AppointmentInProgress` (`models/appointment.go:17-18`),
  campos `CheckedInAt`/`StartedAt` timestamptz (`appointment.go:116,120`), CHECK de 7 valores
  (`appointment.go:87`), bloco DO pré-AutoMigrate idempotente que recria `chk_appointments_status`
  (`database/database.go:184-211`), EXCLUDE anti-overlap mantida bloqueando os novos status
  (`database.go:439-448`), `CheckIn` (`appointment_service.go:605`) e `StartConsultation`
  (`:637`) com guarda de transição + idempotência, toDTO mapeando os carimbos (`:1066-1072`),
  handlers (`appointment_handler.go:306,331`), rotas `POST /:id/check-in` e `POST /:id/start` sob
  RequireAnyStaff (`main.go:708-709`).
- **Waitlist:** model com UUID v7 (`waitlist_entry.go:65-70`), service CRUD+filtro, handler+dto
  completos, rotas GET/POST/PUT/DELETE sob Auth+RequireAnyStaff+AuditLog (`main.go:716-723`),
  AutoMigrate (`database.go:308`).
- **Pagamentos:** `ConsultationPrice`+`PaymentReceiptCounter` (`consultation_price.go:33-47`),
  `AppointmentPayment` UUID v7 (`appointment_payment.go:72-77`), ReceiptNumber sequencial
  transacional via UPSERT ON CONFLICT RETURNING (`payment_service.go:81-93`), `GenerateReceipt`
  gera PDF real via gofpdf com CNPJ 66.991.259/0001-50 (`:144-158,258-364`),
  Create/List/Refund/UpsertPrice/ListPrices, handlers completos, rotas `/payments` sob
  RequireAdminOps e `/consultation-prices` (GET RequireAdminOps, PUT RequireAdmin)
  (`main.go:727-742`), AutoMigrate dos 3 models (`database.go:309-311`).
- **Busca de pacientes:** `List` com param `search` (ILIKE nome/telefone + blind index CPF p/ 11
  dígitos via `crypto.BlindIndexCPF`) (`patient_service.go:148,157-177`), helper `digitsOnly()`
  (`:17-26`), handler lê `?search=` (`patient_handler.go:142,147`), rota GET /patients sob
  RequireAnyStaff (`main.go:596`).
- **Cadastro rápido:** `CreatePatientRequest` com birthDate/gender opcionais (`dto/patient.go:14-15`),
  parse condicional + gender default `other` (`patient_service.go:50-66`), `CalculateAge` trata
  BirthDate zerada (`models/patient.go:208-218`).

**Frontend**
- `/recepcao` completa e ligada no sidebar (`collapsible-sidebar.tsx:73`, staffOnly admin/secretary/manager)
  com 8 components em `components/recepcao/*` auto-consistentes (appointment-row, agenda-hoje-card,
  a-confirmar-card, mensagens-card, leads-novos-card, patient-quick-search, empty-state, recepcao-helpers).
- **Financeiro:** `lib/api/payments.ts` (hooks + openReceipt via getBlob + formatBRL),
  `register-payment-dialog.tsx` com valor pré-preenchido e abertura automática do recibo.
- **Busca global:** `components/global-search.tsx` completo (cmdk+Dialog, Cmd/Ctrl+K, debounce 300ms,
  GET /api/v1/patients?search=, push /patients/:id). cmdk no package.json e `ui/command.tsx` presentes.
- **Cadastro rápido:** `quick-add-patient-dialog.tsx` com `useFormNavigation`.
- **Waitlist client:** `lib/api/waitlist.ts` completo (useWaitlist/useCreate/useUpdate/useDelete).
- **Filtro "não confirmados":** card com `useAppointments({status:'scheduled'})` + `useConfirmAppointment`.
- Interface `Appointment` ganhou `checkedInAt`/`startedAt` tipados (`calendar-api.ts:53-54`).

### O que FALTA / está QUEBRADO

**Bloqueadores (prioridade máxima)**
- `/hoje/page.tsx` importa `useCheckInAppointment`, `useStartAppointment` e `waitMinutes` de
  `@/lib/api/calendar-api`, que **hoje não existem** (verificado, grep = 0) — **mas existiram**: a
  sessão 1 os criou e a sessão 2 os reverteu. Recuperáveis de `c36529d5` (L1943-1947).
- `AppointmentStatus` e os Records LABELS/COLORS em `calendar-api.ts` voltaram aos 5 status antigos
  (revertidos); `/hoje` indexa por `checked_in`/`in_progress` → undefined. Idem: recuperável.
- `<GlobalSearch>` não está montado em nenhum layout AGORA (a sessão 1 montou em `layout.tsx`, a
  sessão 2 reverteu). Cmd+K não registrado. Recuperável de `c36529d5` (L2019).
- Tudo untracked/uncommitted.

**Gaps de feature**
- Duas páginas paralelas (`/recepcao` linkada sem balcão; `/hoje` com balcão, órfã e quebrada) — unificar.
- Drag-and-drop de reagendar **ausente**: `calendar-grid.tsx:145-148` usa `<button onClick>` estático;
  @dnd-kit está no package.json mas não é importado; sem `useUpdateAppointment({scheduledAt})`.
- Sem UI para adicionar à lista de espera (`useCreateWaitlistEntry` nunca usado); só dá pra cancelar.
- Sem card de Lista de espera na coluna lateral de `/recepcao`.
- Botões Chegou/Iniciar/Concluir só em `/hoje`; `appointment-row` de `/recepcao` só faz Confirmar/Abrir/Teleconsulta.
- Bugs de campo herdados em `/hoje`: usa `c.displayName`/`c.lastMessagePreview` (corretos: `c.name`/`c.lastSnippet`)
  e `useLeads(...,1,8)` lendo `.leads` (correto: `.items`, page 0-based) → renderiza vazio. Corrigir ao migrar pra /recepcao.
- Cadastro rápido sem call-sites em appointments/new e patients/page (o encadeamento "cadastrar e já selecionar" não existe).
- Inconsistência DTO vs model: `Patient.BirthDate` segue `gorm:"not null" validate:"required"`
  (`models/patient.go:91`), então cadastro rápido grava `0001-01-01` e toDTO devolve "0001-01-01" (`patient_service.go:309`).
- (Pré-existente, fora da feature) `lib/api/patients.ts` chama `/patients` sem prefixo `/api/v1` e lê `response.data` — 404 contra o APIClient atual. Não usado pela GlobalSearch.

### Próximos passos para retomar (ordem)

0. **Decidir o Caminho A ou B** (ver "Decisão de retomada" acima). Os passos abaixo assumem que
   primeiro recupera-se de `c36529d5` o `calendar-api.ts` (type + LABELS/COLORS + hooks + `waitMinutes`),
   que os dois caminhos precisam e já estava pronto.
1. (Schema já resolvido pelo AutoMigrate no startup; sem passo de `pnpm generate`/atlas.)
2. Recuperar de `c36529d5` (L1943-1947) e reaplicar em `lib/api/calendar-api.ts`: type
   `AppointmentStatus` com `checked_in`/`in_progress`, Records LABELS/COLORS, hooks
   `useCheckInAppointment`/`useStartAppointment` e helper `waitMinutes`.
3. **Caminho A:** portar os fluxos de balcão de `/hoje` para os cards de `/recepcao`
   (Chegou/Iniciar/Concluir + pagamento + cadastro rápido + lista de espera) e deletar `/hoje`,
   corrigindo os bugs de campo herdados (`c.name`/`c.lastSnippet`; `useLeads(filter,0,5)` lendo `.items`).
   **Caminho B:** recuperar `layout.tsx` (GlobalSearch) e o `calendar-api.ts` da sessão 1, repontar o
   sidebar para `/hoje`, descartar `/recepcao` + os 8 cards.
4. Montar `<GlobalSearch>` no layout autenticado (recuperável de `c36529d5` L2019) para registrar o Cmd+K
   (se for Caminho A; no B já vem na recuperação do layout.tsx).
5. Adicionar card de Lista de espera na página sobrevivente + UI de "adicionar à lista de espera" via `useCreateWaitlistEntry`.
6. Drag-and-drop em `components/calendar/calendar-grid.tsx` com @dnd-kit (DndContext + draggable/droppable →
   `useUpdateAppointment({scheduledAt})` no onDragEnd; a EXCLUDE do banco já trata conflito via toast).
7. Wire do cadastro rápido em `appointments/new/page.tsx` (cadastrar e já selecionar) e em `patients/page.tsx`.
8. Resolver `Patient.BirthDate`: torná-lo nullable de verdade (remover `not null`+`required` em
   `models/patient.go:91`) e ajustar toDTO (`patient_service.go:309`) p/ devolver string vazia quando zerada.
   (Mudança de model → AutoMigrate aplica no próximo startup; sem passo de geração.)
9. Validar end-to-end (ver seção Verificação acima), commitar tudo no `master` e remover a página descartada.
10. (Opcional, fora do plano original) telas de histórico/refund de pagamentos e admin de `ConsultationPrice`.
