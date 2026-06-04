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

## Resumo do caminho crítico

1. **Business Verification** (Passo 2) — começa já, leva dias.
2. Em paralelo: criar App+WABA (1), decidir Tech Provider vs BSP (7), submeter templates (5).
3. Depois da verificação: Embedded Signup coexistence (3) → gerar token/IDs/webhook (4) → eu seto
   as envs no Coolify (6) → smoke test.

**O que já está pronto e não precisa de nada:** todo o código do CRM/WhatsApp. É só plugar a conta.
</content>
