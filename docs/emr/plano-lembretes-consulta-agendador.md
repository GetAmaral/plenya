# Agendador de lembretes/templates de consulta (WhatsApp) — A FAZER

> Registrado 2026-06-06. Pré-requisito (templates) **já feito**: os templates UTILITY pt_BR estão
> criados e PENDING na WABA real `827642893502160` (ver `docs/lgpd/whatsapp-templates.md`).
> Falta a **automação de envio**. Hoje os templates só dão pra enviar manualmente pela `/conversas`.

## Objetivo
Job(s) que leem a agenda real do EMR e disparam o template certo na hora certa, dentro/fora da
janela 24h (template reabre a janela). Mesmo padrão do `scheduler/conversation_auto_reply_job.go`
(tick periódico + idempotência por activity já enviada).

## Disparos
| Template (já criado) | Gatilho |
|---|---|
| `confirmacao_consulta_semana` | ~7 dias antes do `Appointment` |
| `confirmacao_consulta_vespera` | 1 dia antes (manhã da véspera) |
| `confirmacao_consulta_dia` | no dia, X horas antes |
| `followup_pos_consulta` | algumas horas/1 dia após a consulta concluída |
| `reengajamento_lead` | lead sem resposta há N dias (revisar categoria: pode virar MARKETING) |

## Pontos de projeto (decidir na execução)
- **Fonte:** `Appointment` (data/hora, paciente, modalidade presencial/online). Mapear params de
  cada template a partir do appointment + patient.
- **Idempotência:** registrar `LeadActivity`/marcador por (appointment_id, template) pra não duplicar;
  re-tick natural não reenvia.
- **Janela/consentimento:** template reabre 24h, mas respeitar opt-out (`IsUnsubscribeKeyword`,
  `WhatsAppOptIn`) e LGPD. Confirmação de consulta é UTILITY (transacional) → ok; reengajamento é o
  caso sensível.
- **Envio:** reusar `ConversationService.SendWhatsAppTemplate` (já existe) + `WhatsAppService.SendTemplate`.
- **Config:** nomes de template via env (já há `WHATSAPP_TEMPLATE_*`; adicionar os novos) + offsets
  configuráveis (dias/horas) + kill switch por tipo.
- **Resposta do paciente:** "Confirmar"/"Remarcar" chega como inbound texto → cai no fluxo de
  `/conversas` (copiloto/secretária). Futuro: quick-reply buttons (exige tratar payload de botão no webhook).

## Escopo fora (por ora)
- Quick-reply buttons nos templates (UX melhor, mas exige o backend tratar button payload).
- Painel de configuração visual dos offsets.
