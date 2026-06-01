# Cutover AutoMigrate → goose em produção + reconciliação de drift (2026-06-01)

## Contexto
Prod rodava GORM AutoMigrate (`MIGRATIONS_AUTO=true`, `RUN_MIGRATIONS` off). `goose_db_version`
nem existia em prod. Decisão: virar prod pro goose (alinhar à arquitetura documentada —
ver `migrations-decisao.md`). Ao diffar dev×prod (read-only) descobriu-se **drift real**: o
AutoMigrate de prod loga "✅ concluído" mas silenciosamente NÃO aplica ALTERs nem cria algumas
tabelas novas. Prod congelou perto do estado do restore PG18 (30/05).

## Drift verificado (prod ≠ models), 2026-06-01

**Tabela faltando em prod:**
- `appointment_recalls` (feature recall do /recepcao; model + AutoMigrate presentes, tabela não nasceu)

**Colunas faltando em prod (declaradas nos models):**
- `users`: oauth_provider, oauth_provider_id, oauth_picture_url (nullable)
- `prescriptions`: active_ingredient, category, concentration, dosage, duration, frequency,
  instructions, medication_definition_id, medication_name, quantity, quantity_in_words, route
  (prod tem 0 linhas nesta tabela)
- `article_score_items`: auto_linked, confidence_score, feedback_at, feedback_by, linked_at,
  linked_by, user_feedback (todas nullable)
- `embedding_queue`: max_retries, metadata, priority (nullable)
- `lab_request_template_tests`: created_at (NOT NULL — usar DEFAULT now())

**Nullability:**
- `patients.birth_date`: prod nullable, dev/model NOT NULL (decidir após checar dados de prod)

**Ruído conhecido (ignorar):** `goose_db_version` (controle), `article_score_items_backup_20260217`
e `lab_test_definitions_backup_consolidation` (backups só no dev).

## Plano (em ordem)

1. **Backup** do Postgres de prod (`pg_dump` full → `~/.plenya-vps-secrets/backups/`).
2. Levantar row counts das tabelas afetadas + checar nulls em `patients.birth_date`.
3. Extrair do dev as definições exatas: `CREATE TABLE appointment_recalls` (+índices/constraints)
   e os defs de coluna faltantes.
4. Escrever migration goose **`00003_reconcile_prod_schema.sql`** — idempotente
   (CREATE TABLE IF NOT EXISTS, ADD COLUMN IF NOT EXISTS) pra ser no-op no dev e aplicar o delta
   em prod. Cobrir TODA a lista acima.
5. **Dev:** `migrate up` → 00003 aplica como no-op; confirmar `go build` + app de pé.
6. **Prod (cutover):**
   a. Criar `goose_db_version` e **stamp** 00001 + 00002 (prod já tem o equivalente; o que falta da
      baseline é coberto pelo 00003 com IF NOT EXISTS).
   b. `RUN_MIGRATIONS=true` + `MIGRATIONS_AUTO=false` no Coolify (app `kgcuxgvmnbx6yya35e3ca2v0`).
   c. Redeploy → `prod-entrypoint.sh` roda `migrate up` → aplica só 00003.
7. **Verificar:** re-dump schema de prod e diffar contra dev (zero diff fora do ruído);
   `goose_db_version` com 00001/00002/00003; health 200; testar features que estavam quebradas
   (criar receita, recall).

## Status de execução
- [x] 1 backup (`~/.plenya-vps-secrets/backups/prod_20260601_pre_goose_cutover.dump`, 219M)
- [x] 2 row counts — prescriptions/article_score_items/embedding_queue/lab_request_template_tests = 0 linhas; patients = 1 (birth_date sem null); users = 2 → NOT NULL seguro
- [x] 3 defs do dev extraídos
- [x] 4 migration — **00003 já existia** (recepcao, commit ed824925, cria appointment_recalls); minha reconciliação virou **`00004_reconcile_prod_schema.sql`** (só colunas + birth_date)
- [x] 5 dev `migrate up` → v4 (00003+00004 no-op no dev)
- [x] **Validação de ouro:** restaurei o schema de prod em `prod_sim`, rodei `migrate up` (auto-stamp baseline + 02/03/04) e diffei contra dev → **1236=1236 colunas, zero drift**; tabelas idênticas (fora o backup-fantasma do dev). Cutover provado.
- [ ] 6 cutover prod (RUN_MIGRATIONS=true + MIGRATIONS_AUTO=false + redeploy)
- [ ] 7 verificação

> Nota: o `migrate up` faz o stamp da baseline sozinho (`adoptIfLegacy`) quando acha schema sem `goose_db_version`. Não precisa stamp manual em prod.

## Rollback
Se algo falhar no passo 6/7: restaurar o backup do passo 1; reverter envs (`MIGRATIONS_AUTO=true`,
`RUN_MIGRATIONS` off); redeploy. 00003 é idempotente e aditivo (não dropa nada), risco baixo.
