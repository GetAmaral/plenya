# WhatsApp → EMR — runbook de go-live (provisionamento Meta)

> Estado em 2026-06-04. **O código está COMPLETO e deployado.** Toda a integração WhatsApp do CRM
> existe e roda em modo fallback-log enquanto não há credenciais Meta: inbound (texto + mídia +
> áudio com transcrição), outbound (texto + mídia + template), janela de 24h, opt-in/PARAR,
> coexistence (ecos do app do celular), criptografia em repouso, dedup por `wa_message_id`,
> roteamento de anexo de paciente pro prontuário + "interpretar como exame". Revisão de segurança
> feita (commits 21329cec/9a84c88c/ddc12af9). Detalhes: memória `crm_whatsapp_phase2_plan`.

Falta só **ligar a conta Meta** e setar as envs no Coolify. Decisão fixada (usuário, 2026-05-22):
conectar o **número existente +55 43 99974-8899 via Coexistence** (mantém o app no celular
funcionando). Ver memória `whatsapp_coexistence_available`.

---

## Pré-requisitos (você)

- **App WhatsApp Business** instalado no celular com o número 99974-8899 (não o WhatsApp comum),
  versão ≥ 2.24.17. Coexistence exige o número já ativo no app **Business**.
- **CNPJ + contrato social** da Plenya Serviços de Saúde Ltda. (66.991.259/0001-50) + documento do
  sócio Dr. Getúlio — para a Business Verification.
- Acesso ao **Meta Business Manager** (business.facebook.com) com a conta que administra a Plenya.

---

## Caminho crítico: Business Verification primeiro (leva dias)

### Passo 1 — Meta App + WhatsApp Business Account (WABA)
**Quem:** você (logado na Meta). Eu te guio campo a campo.

1. developers.facebook.com → criar **App** (tipo: Business) → adicionar produto **WhatsApp**.
2. Criar/conectar a **WABA** à empresa Plenya no Business Manager.

### Passo 2 — Business Verification (o gargalo)
**Quem:** você. Submeter CNPJ + contrato social + doc do sócio. Aprovação leva **dias**.
➜ **Começar por aqui**, em paralelo com tudo o mais. Sem isso o número não vira "produção".

### Passo 3 — Conectar o número via Coexistence (mantém o celular)
**Quem:** você, via **Embedded Signup**.

- O onboarding de coexistence exige que o App Meta esteja como **Tech Provider/Solution Partner**.
  Decisão a tomar (Passo 7): usar nosso próprio App como Tech Provider **ou** ir por um BSP.
- No Embedded Signup, escolher "usar número já no app Business" → sincroniza histórico (até 6 meses)
  e passa a espelhar mensagens nos dois lados.
- ⚠️ Não funciona em coexistence: voz/vídeo, grupos, transmissão, localização ao vivo, catálogo.
  O código já trata os ecos do app (`WHATSAPP_COEXISTENCE=true`, já é o default).

### Passo 4 — Token, IDs e webhook
**Quem:** você gera na Meta; eu coloco no Coolify.

1. Gerar **Permanent Access Token** (System User, sem expiração).
2. Anotar **Phone Number ID** e **App Secret**.
3. Configurar **webhook**: URL `https://api.plenyasaude.com.br/api/v1/webhooks/whatsapp`,
   gerar um `WHATSAPP_WEBHOOK_VERIFY_TOKEN` aleatório (32+ chars), subscrever o field **messages**.
4. Garantir HTTPS válido e o path liberado no firewall (Meta bate de IPs variáveis).

### Passo 5 — Templates (aprovação ~24h)
**Quem:** você submete na Meta; o texto já está versionado.

- `magic_link` (utility, pt_BR) — Escore Light. Body em `crm_phase1_pending_setup`.
- `lead_alert` (utility, pt_BR) — notificação interna de novo lead. Body em `crm_phase2_pending_setup`.
- `docs/lgpd/whatsapp-templates.md` é o registro canônico dos bodies.

### Passo 6 — Setar envs no Coolify (app api `kgcuxgvmnbx6yya35e3ca2v0`)
**Quem:** eu (quando você me passar os valores do Passo 4).

```
WHATSAPP_APP_SECRET=<app secret>
WHATSAPP_ACCESS_TOKEN=<permanent token>
WHATSAPP_PHONE_NUMBER_ID=<id do número>
WHATSAPP_WEBHOOK_VERIFY_TOKEN=<random 32+ chars>
WHATSAPP_COEXISTENCE=true            # default; manter ligado p/ não duplicar ecos do celular
# templates já têm default magic_link / lead_alert
```
Depois: backfill de phones legados (uma vez): `cmd/backfill-patient-phones` (ver `crm_phase2_pending_setup`).

### Passo 7 — Decisão: App próprio (Tech Provider) vs BSP
**Quem:** decisão sua, eu explico o trade-off.

- **App próprio como Tech Provider:** mais controle, sem custo de intermediário, mais setup Meta.
- **BSP (ex.: provedor parceiro):** onboarding de coexistence mais guiado, custo por conversa/mensal.
- Para um único número clínico, o App próprio costuma bastar. Decidir antes do Passo 3.

### Passo 8 — DPA + LGPD
**Quem:** você assina; eu aponto os débitos.

- Assinar o **DPA da Meta** (Business Manager → Compliance).
- ⚠️ **Transcrição de áudio (Whisper/OpenAI)** = sub-processador. Antes de ligar com paciente real,
  sign-off do DPO + RoPA (débito já registrado na Fase 6 LGPD). Até lá, manter transcrição off
  ou só para leads não-pacientes.

---

## Verificação após go-live (smoke)

1. `GET /webhooks/whatsapp?hub.mode=subscribe&hub.verify_token=<correto>&hub.challenge=hello` → `hello`.
2. Mandar mensagem real de um celular pro 99974-8899 → Lead com `source=whatsapp_inbound` em `/conversas` em segundos.
3. Responder pelo CRM (dentro da janela 24h) → chega no celular.
4. Mandar mensagem **pelo app do celular** (coexistence) → aparece no CRM como `via app` (eco), sem criar lead duplicado.
5. Mandar foto/PDF de paciente cadastrado → vira anexo + botão "salvar no prontuário"/"interpretar como exame".
6. Enviar "PARAR" do celular → lead vira `unsubscribed`, opt-out respeitado.

---

## Progresso real (2026-06-04)

- Portfólio empresarial **drgetulioamaralfilho** (business_id `1965827327477102`): dados da empresa
  preenchidos com a Plenya (razão social, CNPJ `66.991.259/0001-50`, endereço fiscal, tel, site).
- **Verificação de Negócio submetida → status "Em análise"** (Meta achou o registro pelo CNPJ, não
  pediu documento). Aguardando aprovação por email (dias).
- **App Meta criado:** "Plenya CRM", app ID `919138737851741`, caso de uso "Conectar-se com clientes
  pelo WhatsApp", conectado ao portfólio drgetulioamaralfilho.
- Etapa 1 (número de teste) provisionada — sandbox descartável.
- ⚠️ **Bloqueio de decisão:** a Etapa 2 ("Integrar com API") só oferece **"Adicionar novo número"**
  (registro/migração via SMS → tira o número do app do celular). **Não há coexistence nesse caminho.**
  Coexistence (manter o celular) exige o fluxo de **Embedded Signup**, que vem pela trilha
  **"Torne-se um Provedor de Tecnologia"** (Tech Provider) ou por um **BSP**. Decisão pendente do
  usuário antes de prosseguir (não registrar o 99974-8899 pela trilha de migração).

## Pesquisa coexistence (2026-06-04, fontes oficiais Meta)

Doc oficial (developers.facebook.com/.../embedded-signup/onboarding-business-app-users, atualizado
2026-05-21) + confirmação em fóruns. Conclusões:

- Coexistence (número no app do celular **e** no Cloud API, histórico 6 meses sincronizado, espelho
  de mensagens) **existe e é GA**, MAS o onboarding **só** acontece via **Embedded Signup** e
  **exige que o app seja Tech Provider ou Solution Partner** + **App Review** aprovado + assinatura
  dos webhooks extras `history`, `smb_app_state_sync`, `smb_message_echoes` (session logging).
- **Confirmado: não há caminho de coexistence para o número próprio direto no Cloud API sem virar
  Tech Provider** (ou usar um BSP que já seja um). Fonte: doc "Requirements" + Reddit r/WhatsappBusinessAPI.
- Limitações mesmo COM coexistence: throughput fixo 20 mps; **sem grupos, sem chamadas voz/vídeo pela
  API, sem catálogo/etiquetas/respostas rápidas/mensagem de ausência/perfil comercial editável pela
  API**; listas de transmissão desativadas; disappearing/view-once/live-location desligados nos 1:1;
  companion devices são desvinculados no onboarding (re-vincular depois); abrir o app a cada ~13 dias.

### Implicação e opções reais
- **A) Migrar o número pro Cloud API (direto, nosso app):** simples, encaixa 100% no backend que já
  temos (Cloud API direto, zero dev novo). O número **sai do app do celular**; atendimento 100% pelo
  CRM `/conversas`. Precisa só da verificação (em análise) + forma de pagamento. **Recomendado.**
- **B) Coexistence virando Tech Provider (nosso app):** mantém o celular, mas exige Tech Provider +
  Embedded Signup (fluxo web novo) + App Review + webhooks de histórico. Semanas de trabalho +
  burocracia, para um único número. Desproporcional.
- **C) Coexistence via BSP:** mantém o celular, o BSP é o Tech Provider; porém o número passa a viver
  no BSP → nossa integração deixaria de ser Cloud API direto (rework) + custo por conversa/mensal.

Decisão revisada do usuário pendente (a escolha original por coexistence foi feita sem saber do peso
de Tech Provider + Embedded Signup + App Review).

**Empírico (2026-06-04):** no nosso app "Plenya CRM", o fluxo "Adicionar número" é Perfil → Adicionar
número → Verificar (OTP) = registro/migração. **Não há opção self-serve de coexistence** (nenhum
"conectar número já no app"/QR). Confirma que coexistence = só via Tech Provider/BSP.

### DECISÃO 2026-06-04: número real 99974-8899 em STAND-BY
Não migrar nem fazer coexistence por ora. Próxima fase = **testar as funcionalidades do EMR usando o
número de TESTE da Meta** (sandbox da Etapa 1: `+1 555 646-5351`, Phone Number ID `1042915892248397`,
WABA de teste `27890641940528227`), sem tocar no número real. App "Plenya CRM" id `919138737851741`.

## Resultado do teste com número de teste (2026-06-04) — INBOUND e OUTBOUND OK

Prod (app api `kgcuxgvmnbx6yya35e3ca2v0`) foi apontado pro **número de TESTE da Meta** pra validar o EMR.

**Credenciais setadas no Coolify (via API, bulk) + redeploy aplicado:**
- `WHATSAPP_APP_SECRET`, `WHATSAPP_ACCESS_TOKEN` (token de teste, expira ~24h), `WHATSAPP_PHONE_NUMBER_ID=1042915892248397`,
  `WHATSAPP_WEBHOOK_VERIFY_TOKEN=ecf6abd2c6d0edd78e5d6388ab237a606f3c6dc7`, `WHATSAPP_GRAPH_API_VERSION=v25.0`.
- Cofre local: `~/.plenya-vps-secrets/whatsapp-test.env`. App "Plenya CRM" id `919138737851741`.
- ⚠️ **Coolify `restart` NÃO reaplica env alterada; precisa `deploy` (redeploy).**

**Webhook (configurado no painel + via Graph):** callback `https://api.plenyasaude.com.br/api/v1/webhooks/whatsapp`
ativo; app inscrito na WABA de teste (`POST /{waba}/subscribed_apps`); **campo `messages` assinado**
(`POST /{app}/subscriptions fields=messages`). ⚠️ **Bloqueio que travou o 1º teste:** o `fields` estava `[]`
(só o objeto inscrito, sem o campo `messages`) → nenhuma mensagem chegava. Assinar `messages` resolveu.

**Inbound:** ✅ funciona ponta a ponta. Mensagem do celular → `POST /webhooks/whatsapp 200` → Lead criado
(`source=whatsapp_inbound`) + `lead_activities` (`message_received`, conteúdo cifrado `v1:`).
**Outbound:** ✅ funciona. Erro inicial `(#131030) Recipient not in allowed list` era **só do sandbox**
(lista de permitidos + nono dígito BR: wa_id vem `554399123452`). Resolvido adicionando o número na lista.
Em número real não existe lista de permitidos.

**Bugs/pendências achados no teste:**
1. 🐛 **Preview da lista de `/conversas` mostra texto cifrado `v1:...`** — `decryptIfNeeded`
   (`conversation_service.go:773`) não trata o envelope `v1:` (faz `base64.Decode("v1:...")` que falha →
   devolve cifrado). Fix pequeno no backend. (Entra no plano de redesenho.)
2. ✅ **CORRIGIDO (2026-06-04):** `User.ProfessionalPhone` do Dr. Getúlio era `"(43)999748899"` (não-E.164)
   e quebrava a notificação de lead via WhatsApp. Normalizado pra `+5543999748899` direto no banco de prod.
   (Obs: entrega completa do alerta de lead via WhatsApp ainda depende do template `lead_alert` aprovado;
   o sino in-app já funciona.)
3. ⚠️ **`/conversas` mal resolvido** (mistura email+WA, sem cara de chat, barra de paciente indevida).
   Plano dedicado: `plano-conversas-redesign.md`.

> **Para voltar ao normal / ir pro número real depois:** trocar as 5 envs do Coolify pelas credenciais
> do número real, re-inscrever o app na WABA real + campo `messages`. Número real 99974-8899 segue em stand-by.

## Teste do recepcionista virtual (Copiloto + Automático) com número particular — 2026-06-05

Objetivo: validar, ponta a ponta no EMR de produção (ainda apontado pro número de TESTE da Meta), os
dois modos do recepcionista virtual usando o WhatsApp pessoal do usuário como "lead".

### Arquitetura (confirmada no código, nada a desenvolver)
- Toggle por conversa **Off · Copiloto · Automático** — `components/conversations/automation-toggle.tsx`
  (badge "global off" aparece se escolher Automático com `RECEPTION_BOT_ENABLED=false`).
- **Copiloto** = botão **"Recepção IA"** no compositor (`conversation-composer.tsx`) → chama
  `POST /conversations/:type/:id/ai/reception-reply` (on-demand). **NÃO depende do kill switch**;
  funciona assim que as credenciais WA de teste estiverem válidas.
- **Automático** = job `ConversationAutoReplyJob` (1/min) → exige `RECEPTION_BOT_ENABLED=true` +
  conversa em modo `auto` + inbound parado há ≥ `RECEPTION_BOT_FALLBACK_MINUTES` (default 5) dentro
  da janela 24h. Handoff (dúvida clínica / "falar com alguém") pausa 24h + notifica staff.
- Lead de inbound entra com `WhatsAppOptIn=true` (consent implícito, `lead_service.go:477`) → auto
  não trava por opt-in. "PARAR" → `unsubscribed`, IA silencia (enforce no job, não só no prompt).
- Cérebro = `receptionSystemPrompt` já revisado à luz do curso (engenharia da escuta, sem pressão).

### Pré-requisitos — só no painel Meta (app "Plenya CRM" `919138737851741`), só o usuário faz
1. **Allow-list**: API Setup → campo "To" → "Manage phone number list" → adicionar o número pessoal
   (OTP). Sem isso, todo outbound falha `(#131030) Recipient not in allowed list` (quirk do sandbox;
   atenção ao 9º dígito BR — wa_id chega sem ele). Em número real não existe lista.
2. **Token**: o token de teste é temporário (~24h) e o de 2026-06-04 está expirado. Pegar um novo em
   API Setup → "Temporary access token", OU gerar um **System User token permanente** pra não expirar
   durante a fase de teste (recomendado).

### Coolify (app api `kgcuxgvmnbx6yya35e3ca2v0`) — aplicar e **REDEPLOY** (restart não reaplica env)
```
WHATSAPP_ACCESS_TOKEN=<novo token de teste>     # o antigo expirou
RECEPTION_BOT_ENABLED=true                       # necessário só pro modo Automático
RECEPTION_BOT_FALLBACK_MINUTES=1                 # acelera o teste (volta pra 5 depois)
# RECEPTION_BOT_DEFAULT_MODE fica 'off'; o modo é setado por conversa na UI
```

### Roteiro
**Copiloto** (pode rodar sem mexer no kill switch):
1. Do WhatsApp pessoal, mandar mensagem pro número de teste (+1 555 646-5351) → Lead
   `source=whatsapp_inbound` aparece em `/conversas → WhatsApp` em segundos.
2. Abrir a conversa → toggle **Copiloto**.
3. Clicar **"Recepção IA"** no compositor → gera a resposta ancorada no script revisado; revisar,
   editar se quiser, **enviar** → chega no celular.

**Automático**:
1. Toggle **Automático** (não pode aparecer o badge "global off" — confirma `RECEPTION_BOT_ENABLED=true`).
2. Mandar nova mensagem do celular e **não** responder no CRM.
3. Após `FALLBACK_MINUTES`, o job gera e envia sozinho → chega no celular. Testar também um caso de
   handoff (ex.: "tô com dor no peito" / "quero falar com uma pessoa") → IA manda cortesia, pausa,
   e dispara o sino pra staff.

### Caveats
- Bug cosmético conhecido: preview da lista em `/conversas` mostra `v1:...` cifrado (não afeta o teste).
- Áudio do celular aciona transcrição (Whisper = sub-processador) — pra teste tudo bem; com paciente
  real exige sign-off DPO (débito Fase 6 LGPD).
- Token de teste permanente do System User evita ter que refazer o passo 2 a cada dia.

## GO-LIVE do número REAL 99974-8899 (2026-06-05)

Decisão revisada: **migrar o número real direto pro Cloud API** (sem coexistence). Coexistence exigiria
Tech Provider (App Review, gated por thresholds) ou BSP (custo + rework); desproporcional p/ 1 número.
Migração é **reversível** (deletar o número da WABA → re-registrar no app WhatsApp Business com OTP;
libera em ~1h; perde histórico daquele número). O WhatsApp pessoal fica em outro número, intocado.

**Provisionado e confirmado via Graph API (token permanente):**
- App "Plenya CRM" `919138737851741`, System User "Employee" (token permanente, não expira, escopos
  `whatsapp_business_management` + `whatsapp_business_messaging` + `business_management`).
- **WABA real `827642893502160`** ("Plenya Saúde | Dr Getulio Amaral Filho"), `account_review_status: APPROVED`.
  ⚠️ Após registrar o número, a WABA nova **não vem atribuída ao System User** — atribuir em
  Configurações do negócio → Usuários do sistema → Employee → Adicionar ativos → Contas do WhatsApp →
  Controle total. Sem isso o token dá `error_subcode 33` (missing permissions).
- **Phone Number ID `1193269897200545`**, display +55 43 99974-8899, `VERIFIED`, `CLOUD_API`, `LIVE`,
  nome aprovado, `messaging_limit_tier: TIER_250`.
- App **inscrito na WABA real** (`POST /{waba}/subscribed_apps` → success). Campo `messages` já assinado
  no app (mesmo app do teste).
- Número **registrado no Cloud API** (`POST /{phone_number_id}/register` com PIN `141719` → success).
- Credenciais no cofre: `~/.plenya-vps-secrets/whatsapp-real.env` (phone, PIN, token permanente).

**Pendente (último passo) — swap de 2 envs no Coolify (app api `kgcuxgvmnbx6yya35e3ca2v0`) + REDEPLOY:**
```
WHATSAPP_PHONE_NUMBER_ID=1193269897200545
WHATSAPP_ACCESS_TOKEN=<token permanente do System User>   # ver cofre whatsapp-real.env
# manter WHATSAPP_APP_SECRET, WHATSAPP_WEBHOOK_VERIFY_TOKEN, WHATSAPP_GRAPH_API_VERSION (mesmo app)
```
Redeploy (não restart). Depois: smoke test (WhatsApp → lead em /conversas → Copiloto → Automático).

## Resumo do caminho crítico

1. **Business Verification** (Passo 2) — começa já, leva dias.
2. Em paralelo: criar App+WABA (1), decidir Tech Provider vs BSP (7), submeter templates (5).
3. Depois da verificação: Embedded Signup coexistence (3) → gerar token/IDs/webhook (4) → eu seto
   as envs no Coolify (6) → smoke test.

**O que já está pronto e não precisa de nada:** todo o código do CRM/WhatsApp. É só plugar a conta.
</content>

---

## Template de aniversário (Lívia — item 2 / Fase E) — APROVADO p/ criar no Meta

Cópia aprovada pelo Getúlio (2026-06-07). Falta só criar/aprovar no Meta (test token estava
expirado em 07/06; criação de template é ação no WhatsApp Manager + aprovação assíncrona da Meta).

**Definição do template:**
- **name:** `aniversario_plenya`  (minúsculo, sem espaços — exigência Meta)
- **category:** `MARKETING` (saudação sem transação; a Meta costuma classificar greeting como marketing)
- **language:** `pt_BR`
- **body:** `Oi, {{1}}! Hoje é seu aniversário e a equipe da Plenya passa para desejar um dia feliz e um ano de saúde e bem-estar. Conte com a gente.`
- **exemplo do {{1}}:** `Maria`
- sem header, sem footer, sem botões.

**Como criar via Graph API** (rodar com um token que tenha `whatsapp_business_management` + o WABA ID;
o WABA ID NÃO está no cofre hoje — pegar no WhatsApp Manager ou via `GET /{phone_number_id}` autorizado):

```bash
curl -X POST "https://graph.facebook.com/v25.0/<WABA_ID>/message_templates" \
  -H "Authorization: Bearer <MGMT_TOKEN>" -H "Content-Type: application/json" \
  -d '{
    "name": "aniversario_plenya",
    "category": "MARKETING",
    "language": "pt_BR",
    "components": [
      {"type":"BODY",
       "text":"Oi, {{1}}! Hoje é seu aniversário e a equipe da Plenya passa para desejar um dia feliz e um ano de saúde e bem-estar. Conte com a gente.",
       "example": {"body_text": [["Maria"]]}}
    ]
  }'
```

**Depois de aprovado:** setar a env `WHATSAPP_TEMPLATE_BIRTHDAY=aniversario_plenya` no Coolify (api).
Sem isso, o job só notifica o time (não envia). O auto-envio ainda exige: Atendimento IA em `auto` +
kill switch ligado + aniversário próprio não-restrito no dia.

**Pré-requisito atual:** renovar o token do WhatsApp (o de teste no cofre venceu em 04/06/2026).
