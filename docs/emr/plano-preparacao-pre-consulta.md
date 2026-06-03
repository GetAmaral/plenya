# Plano de build — Preparação pré-consulta (Fase 2)

Feature do EMR que materializa o fluxo desenhado em
[`../atendimento/fluxo-pre-consulta.md`](../atendimento/fluxo-pre-consulta.md): depois de agendar e
pagar, o paciente (já convertido em paciente no EMR) preenche um **formulário curto de preparação**
(subconjunto curado do Escore, no motor da Triagem) e **envia exames anteriores** pelo portal. O Dr.
revisa antes e pede exames complementares.

Decisões (2026-06-03): pré-análise caso a caso pelo Dr.; dois canais (WhatsApp+CRM já existe; **portal
dedicado = esta fase**); formulário curto pós-agendamento. **Itens do formulário serão curados depois**
— esta fase entrega a infra item-agnóstica (flag `PrepOrder` vazia até a curadoria).

## Arquitetura (espelha o Escore Light / Triagem)

- **Flag no ScoreItem:** novo campo `PrepOrder *int` (espelha `LightOrder`). Item entra no formulário
  de preparação se `PrepOrder != nil`; ordenação por `PrepOrder`.
- **Config endpoint:** `BuildPrepConfig()` no `AnonymousScoreService` (irmão de `BuildLightConfig`),
  filtra por `PrepOrder != nil` e reusa o mesmo DTO `LightConfig` (mapeando `LightOrder := *PrepOrder`
  para o front ordenar igual). Exposto em `GET /api/v1/patient/me/prep/config` (auth paciente).
- **Armazenamento:** novos models espelhando `AnonymousScoreSession`/`AnonymousScoreItem`, porém
  ligados ao paciente:
  - `ConsultationPrep` (PatientID, AppointmentID?, Status draft|submitted, ChiefComplaint?, SubmittedAt,
    consent) + `Responses`.
  - `ConsultationPrepResponse` (PrepID, ScoreItemID, NumericValue?/SelectedLevel?/TextValue?).
  Ambos com `BeforeCreate` UUID v7. Pode semear uma `Anamnesis` depois (campos espelham `AnamnesisItem`).
- **Upload de exame pelo paciente:** `POST /api/v1/patient/me/documents` reusando
  `PatientDocumentsService.Create` com `Source = portal` (hoje só existe upload por staff). Adiciona
  `Source` opcional ao `CreateDocumentInput`.

## Increment 1 — Backend (esta entrega)
1. `models/score_item.go`: add `PrepOrder *int`.
2. `models/consultation_prep.go`: novos `ConsultationPrep` + `ConsultationPrepResponse` (+ TableName + hooks).
3. `database/migrations/00016_consultation_prep.sql` (goose, idempotente): cria as 2 tabelas + add coluna
   `prep_order` em `score_items` + índices.
4. `dto/consultation_prep.go`: SubmitPrepRequest, PrepResponseDTO, ConsultationPrepView.
5. `services/consultation_prep_service.go`: GetForAppointment (cria draft), Submit (transação upsert).
6. `services/anonymous_score_service.go`: add `BuildPrepConfig()`.
7. `services/patient_documents_service.go`: `CreateDocumentInput.Source` opcional.
8. `handlers/consultation_prep_handler.go`: GetConfig, GetPrep, SubmitPrep, UploadExam.
9. `cmd/server/main.go`: instanciar service+handler, registrar rotas em `patientMe`.
10. `docker compose exec -w /app api go build ./...` verde + `migrate up` no dev.

## Increment 2 — Portal (esta entrega, se couber)
1. `apps/web/lib/api/patient-portal-api.ts`: tipos + hooks (usePrepConfig, useMyPrep, useSubmitPrep, useUploadExam).
2. `apps/web/components/patient-portal/prep-form.tsx`: renderer portado do escore-light (grupos → itens;
   inputs: botões de nível, numérico com classificação, texto). Sem next-intl, strings PT-BR. PHQ-9 e
   imagens de instrução ficam para quando esses itens forem curados.
3. `apps/web/app/patient-portal/preparacao/page.tsx`: pega a próxima consulta, carrega/cria draft, mostra
   queixa/objetivo + formulário + upload de exames + enviar. `useRequirePatientAuth()`.
4. Entrada no `PatientSidebar`/`PatientBottomNav`.

## Increment 3 — Lembretes (follow-up, NÃO nesta entrega)
- `Appointment.PrepReminderSentAt` + `ConsultationPrepReminderJob` (espelha `AppointmentReminderJob`):
  consultas em janela T-48h..T-24h sem prep submetida → nudge com link do portal. Depende de template
  Meta/WhatsApp; por isso fica para depois. Hoje a recepção envia o link manualmente (script seção 7).

## Regras respeitadas
- Fonte única (Go models) + migration goose `00016` (AutoMigrate off em prod). Rodar `migrate up`.
- UUID v7 via BeforeCreate; rotas sob `middleware.Auth + RequirePatient + AuditLog`.
- Upload: validação de magic bytes/limite já no `PatientDocumentsService`.
- Sem itens hardcoded: formulário dirigido por config (curadoria posterior).

## Verificação
- `go build ./...` verde; `migrate up` aplica 00016; `migrate status` mostra 16.
- Smoke manual (dev, bypass): abrir `/patient-portal/preparacao`, ver formulário vazio (sem itens ainda),
  enviar queixa + upload de um PDF, confirmar persistência (psql: consultation_preps, patient_documents
  com source=portal).
- Front: typecheck/build do apps/web.
