# apps/api — Backend EMR (Go)

API REST única que serve web e mobile. **Fonte única de verdade do projeto** — os Go models
aqui geram migrations, OpenAPI e os tipos TS dos packages. Leia também as Regras de Ouro no
[CLAUDE.md raiz](../../CLAUDE.md).

## Stack
Go 1.25 · Fiber v2.52 · GORM v1.31 · PostgreSQL 17 + pgvector · Atlas (migrations) · Swag (OpenAPI).
**Não há Go instalado localmente** — compile no container: `docker compose exec -w /app api go build ./...`.

## Estrutura (`internal/`)
```
models/       ← 74 models — ÚNICA FONTE. Editar aqui dispara a geração.
dto/          ← request/response types (separados dos models)
services/     ← 96 services — lógica de negócio, métodos toDTO
handlers/     ← 55 handlers — finos, delegam para services
repository/   ← DAOs (acesso a dados via GORM)
middleware/   ← auth (JWT+2FA), rbac, rate-limit, audit-log
crypto/       ← AES-256-GCM + blind index (CPF/RG)
workers/      ← jobs assíncronos (embeddings, processing)
scheduler/    ← cron interno
config/ utils/ database/ templates/
```
Rotas: registradas em `cmd/server/main.go` (majoritariamente inline; treino extraído em
`registerTrainingRoutes()`). Migrations geradas em `database/migrations/` (58 .sql) — **nunca editar**.

## Domínios (handlers)
Patients · Scores (Group→Subgroup→Item→Level + snapshots + anonymous/escore-light) · Labs
(definitions, results, batches com OCR, views) · Anamnesis (+ templates) · Appointments +
Calendar (slots, working hours, doctor absences, Google OAuth, telemed lobby) · Prescriptions
(+ medication definitions, certificados ICP-Brasil/SNCR) · CRM (leads, campaigns, conversations
+ attachments, WhatsApp webhook) · Portal do paciente (patient_portal, me_mobile, mobile_config,
patient_workouts, patient_continuum) · Continuum (templates, boxes, dashboard) · Training
(exercises, workout_plans, periodization, physical/postural assessment, training_ai) ·
Subscriptions (+ plans) · Articles + RAG (article_semantic, score_enrichment_preparation) ·
Notifications · Users/auth · uploads · processing_job · certificates.

## Padrões obrigatórios
- **UUID v7** via `BeforeCreate` em todos os models.
- **Patient:** `BeforeSave`/`AfterFind` criptografam CPF/RG; blind index para busca.
- **ScoreItem/ScoreLevel:** `BeforeUpdate` atualiza `LastReview`.
- Fluxo: `model → dto → service (toDTO) → handler (fino) → rota em main.go`.
- `ClinicalDataService` (`services/clinical_data_service.go`) centraliza acesso a anamnese/lab;
  códigos clínicos em `models/clinical_codes.go` — nunca hardcode códigos.
- Auditoria: writes passam por `middleware.AuditLog` (imutável em prod).

## Geração
```bash
pnpm generate        # atlas diff → swag init → openapi-typescript → openapi-zod-client
```
Saídas geradas (não editar): `database/migrations/*.sql`, `docs/swagger.json`,
`packages/types/src/generated/*.ts`.

## Banco direto (dev)
```bash
docker compose exec db psql -U plenya_user -d plenya_db
docker compose exec db psql -U plenya_user -d plenya_db -c "SELECT * FROM score_items LIMIT 5;"
```
Detalhes: [.claude/workflows/database-ops.md](../../.claude/workflows/database-ops.md) ·
[.claude/backend/models.md](../../.claude/backend/models.md) ·
escores em [.claude/domain/score-system.md](../../.claude/domain/score-system.md).
