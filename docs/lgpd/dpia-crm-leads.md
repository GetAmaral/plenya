# DPIA — CRM (captura e processamento de leads)

**Versão:** 1.0 — 2026-04-23
**Responsável:** Encarregado de Proteção de Dados (DPO) — Dr. Getúlio Amaral
**Próxima revisão:** semestral, ou ao introduzir novo canal/sub-processador
**Vinculado a:** [DPIA Escore Light](dpia-escore-light.md), [Plano de Resposta a Incidentes](plano-resposta-incidentes.md)

---

## 1. Objeto e finalidade

O CRM da Plenya, embarcado dentro do EMR (modelo `Lead`, tabelas `leads` + `lead_activities`), centraliza contatos capturados em quatro pontos:

1. **Claim do Escore Plenya Light** — usuário recebe magic link por email e/ou WhatsApp pra salvar resultado.
2. **Formulário público `/contato`** — contato comercial direto.
3. **WhatsApp inbound** — qualquer mensagem inbound no número Plenya WhatsApp Business cria/atualiza Lead.
4. **Newsletter opt-in** — flag opcional dentro do claim do Light (Phase 1 não tem signup standalone).

Finalidade legítima: identificar prospects qualificados de pacientes e operacionalizar contato comercial, com base legal **consentimento** (art. 7º I LGPD) para email e WhatsApp; **legítimo interesse** (art. 7º IX) para criação do registro de Lead em si.

## 2. Categorias de dados pessoais tratadas

| Categoria | Campo | Origem | Sensível? |
|-----------|-------|--------|-----------|
| Identificação | `name` | Form / Meta profile | Não |
| Contato | `email` | Form / claim Light | Não |
| Contato | `phone` (E.164) | Form / claim Light / Meta WaID | Não |
| Comportamento | `message`, `metadata` (motivo, janela) | Form / WhatsApp inbound | Não |
| Consent | `consent_version`, `consent_timestamp`, `consent_ip_hash` | Submit do form | Não |
| Atividade | `lead_activities` (mensagens, status, notas) | Sistema | **Pode conter** texto livre do cliente |

**Não são tratados:** CPF, RG, dados de saúde, exames, dados financeiros. O Lead é pré-paciente — converter em Patient transfere o tratamento pro escopo do EMR (DPIA separado).

## 3. Sub-processadores

| Sub-processador | Localização | Dados expostos | Base contratual |
|-----------------|-------------|----------------|-----------------|
| **Resend Inc.** | EUA | Email do destinatário, nome, conteúdo do email transacional | Cláusulas-padrão LGPD/GDPR + Resend DPA |
| **Meta Platforms (WhatsApp Business Cloud API)** | EUA | Telefone E.164, nome do perfil WhatsApp, conteúdo das mensagens enviadas/recebidas | Meta DPA (assinado em Business Manager) + cláusulas-padrão |
| **Anthropic** | EUA | Não recebe dados de Lead (apenas usado pra extração de PDF de exames no Light, escopo separado) | DPIA Light, item 3 |
| **KingHost / Coolify** | Brasil | Todos os dados em repouso | Contrato de hospedagem |

Transferência internacional EUA: justificada art. 33 LGPD (cláusulas-padrão), comunicada na Política de Privacidade pública.

## 4. Fluxos e ciclo de vida do dado

| Etapa | O que acontece | Retenção |
|-------|----------------|----------|
| Captura | Lead criado com `consent_version`, `consent_ip_hash` (SHA-256 do IP) | Até exclusão pelo titular ou pela equipe |
| Envio outbound | Email via Resend OU WhatsApp via Meta (template utility "magic_link"). LeadActivity registra envio | Permanente em `lead_activities` |
| Inbound WhatsApp | Webhook Meta valida HMAC, cria/atualiza Lead, registra mensagem | Permanente em `lead_activities` |
| Conversão | Lead vira Patient (manual ou via claim do Light) — `status=converted`, `converted_patient_id` populado | Lead permanece em "histórico"; Patient é regido pela DPIA do EMR |
| Exclusão | Soft delete via `deleted_at` (atividades em cascade) | Permanente após exclusão |

## 5. Direitos do titular (LGPD art. 18)

| Direito | Como exercer | Onde implementado |
|---------|--------------|-------------------|
| Confirmação e acesso | Solicitar via DPO (`dpo@plenyasaude.com.br`) — admin exporta JSON do Lead | Manual (UI de export Phase 2) |
| Correção | Pelo DPO ou via UI admin | `PATCH /api/v1/leads/:id` |
| Portabilidade | Sob demanda via DPO (export JSON estruturado) | Manual Phase 1 |
| Eliminação | `DELETE /api/v1/leads/:id` (admin) ou pedido ao DPO | `lead_handler.go:Delete` |
| Revogação de consentimento | Mudar status pra `unsubscribed` | UI admin Phase 1 |
| Oposição ao WhatsApp | Cliente envia "PARAR" → admin marca `whatsapp_opt_in=false` | Manual Phase 1; automatizar Phase 2 |

Prazo de resposta: 15 dias corridos (art. 19 §3º LGPD) — prazo legal cumprido por SLA interno.

## 6. Riscos identificados e controles

| Risco | Severidade | Probabilidade | Controle |
|-------|------------|---------------|----------|
| Vazamento de Leads via API pública (`POST /leads`) | Alto | Baixa | Rate limit 5 req/h por IP + CSRF Origin no proxy do site |
| Webhook WhatsApp aceita payloads forjados | Crítico | Baixa | HMAC SHA-256 validado contra `WHATSAPP_APP_SECRET`; bloqueia request em prod se secret ausente |
| Magic link interceptado em phishing | Médio | Muito baixa | Token JWT single-use, expira em 7 dias; só dá acesso a um Patient específico (não login universal) |
| Lead duplicado por race em claim simultâneo | Baixo | Média | `uniqueIndex` em `anonymous_score_session_id` + recovery por `unique_violation` |
| Sub-processador EUA sob lei estrangeira (CLOUD Act, FISA) | Médio | Médio | Cláusulas-padrão; mínimo necessário enviado; cliente avisado na Política e no consentimento granular WhatsApp |
| Conteúdo sensível em texto livre de WhatsApp inbound | Alto | Média | Não publicizado fora do EMR; acesso restrito a admin/secretary/manager; auditoria via `audit_logs` |
| Atacante envia spam de inbound forjado pelo número | Baixo | Muito baixa | Rate limit Meta no número Plenya; `lead_activity` versionada permite identificação rápida |

## 7. Avaliação residual

Risco residual **ACEITÁVEL** após controles. Próxima revisão deve verificar:
- Tempo médio de resposta a pedidos LGPD via DPO
- Volume de inbound WhatsApp e proporção converted/lost (sinaliza qualidade da captura)
- Contagem de exclusões mensais (sinaliza UX de unsubscribe)
- Auditoria de acesso a `leads` (quem visualiza dados sem necessidade operacional)

## 8. Aprovação

DPO Dr. Getúlio Amaral — autorização para deploy em produção condicionada a:
1. Resend com DKIM/SPF/DMARC configurados antes do primeiro envio real
2. Meta Business Verification concluída e DPA assinado antes do primeiro envio WhatsApp
3. Variáveis `WHATSAPP_APP_SECRET` e `RESEND_API_KEY` populadas no Coolify (Phase 1 falha fechado se ausentes em prod)
