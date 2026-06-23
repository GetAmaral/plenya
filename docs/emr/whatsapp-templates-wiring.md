# Cabeamento dos templates WhatsApp aprovados (EMR)

Contexto: 9 templates pt_BR aprovados na Meta (WABA `827642893502160`, número real
+55 43 99974-8899), mas só 4 em uso. O fluxo de consulta disparava nomes em inglês
(`appointment_confirmation/reminder_24h/cancelled/rescheduled`) que **não existem aprovados**,
então confirmações/lembretes por WhatsApp falhavam em silêncio (e-mail continuava saindo).

Este plano cabeia os templates aprovados e cria os 2 jobs que faltavam.

## Mapeamento final

| Momento | Template aprovado | Params | Como dispara |
|---|---|---|---|
| Ao agendar (confirmação) | `confirmacao_consulta_semana` | nome · data ("12 de junho") · hora ("14h") · modalidade ("presencial em Londrina") | `SendConfirmation` |
| Véspera (~24h antes) | `confirmacao_consulta_vespera` | nome · hora · modalidade | cron `AppointmentReminderJob` janela [now+23h, now+25h] |
| No dia (~2–4h antes) | `confirmacao_consulta_dia` | nome · hora · frase-dia ("O atendimento é presencial, na nossa unidade em Londrina") | cron, nova janela [now+2h, now+4h] |
| Pós-consulta | `followup_pos_consulta` | nome | novo cron `AppointmentFollowupJob` (EndAt < now, recém-concluída) |
| Lead frio | `reengajamento_lead` | nome | novo cron `LeadReengageJob` (**OFF por default**) |
| Cancelamento / remarcação | — (sem template aprovado) | — | só e-mail; WA desligado (era quebrado) |

## Mudanças

- **config**: 5 nomes de template configuráveis (`WHATSAPP_TEMPLATE_APPT_CONFIRM/VESPERA/DIA`,
  `WHATSAPP_TEMPLATE_FOLLOWUP`, `WHATSAPP_TEMPLATE_REENGAGE`). Defaults = nomes aprovados, exceto
  `REENGAGE` (vazio = OFF, exige opt-in explícito por ser outreach a lead).
- **models**: `appointments.dayof_reminder_sent_at`, `appointments.followup_sent_at`,
  `leads.reengaged_at` (idempotência). Migration goose **00042**.
- **appointment_notification_service.go**: reescreve os params do WA pra cada template PT; remove
  os envios WA quebrados de cancel/reschedule (mantém e-mail); novos helpers de data/hora/modalidade;
  novo método `SendReminderDayOf`.
- **appointment_reminder_job.go**: segunda janela (dia-da-consulta).
- **followup_job.go** (novo): consulta concluída → `followup_pos_consulta`. Limitado a EndAt nas
  últimas 48h (não faz backfill histórico).
- **lead_reengage_job.go** (novo, gated): leads `contacted/qualified`, opt-in WA, sem inbound há 7d+,
  via `ConversationService.SendWhatsAppTemplate` (já respeita opt-out/unsubscribe). Default OFF.

## Garantias

- **Idempotência**: cada estágio tem sua coluna `*_sent_at`.
- **Opt-out**: já existe no webhook (`IsUnsubscribeKeyword` PARAR/SAIR/STOP/CANCELAR/DESCADASTRAR +
  `UnsubscribeByPhone`); o envio a leads checa `WhatsAppOptIn` e `status != unsubscribed`.
- **Sem blast histórico**: janelas temporais limitam o que cada job pega no 1º deploy.

Deploy: migration roda no deploy do `api` (`RUN_MIGRATIONS=true`).
