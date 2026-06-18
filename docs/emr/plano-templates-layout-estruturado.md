# Plano — Templates de exames com layout estruturado (opção C)

**Decisão (2026-06-18):** o **template** passa a ser estruturado (ordem + quebra de página) e é a
fonte da verdade do layout; o **pedido** (`lab_requests.exams`) continua **texto livre** (mantém a
flexibilidade atual). Padrão de *order set* de EHR: catálogo → template ordenado/seccionado →
instância editável. Regra única de paginação: **linha em branco = nova página**. NENHUMA regra de
agrupamento no código.

## 1. Schema (migration goose, próximo número incremental)
```sql
-- up
ALTER TABLE lab_request_template_tests
  ADD COLUMN display_order      integer NOT NULL DEFAULT 0,
  ADD COLUMN page_break_before  boolean NOT NULL DEFAULT false;
-- down
ALTER TABLE lab_request_template_tests
  DROP COLUMN display_order, DROP COLUMN page_break_before;
```
Defaults tornam as 437 linhas existentes válidas sem backfill obrigatório.

## 2. Model (`internal/models/lab_request_template.go`)
No struct do join `LabRequestTemplateTest`: `DisplayOrder int` (json `displayOrder`) e
`PageBreakBefore bool` (json `pageBreakBefore`). Vira join-table explícita; preload ordenado por
`display_order`.

## 3. Serializer / API
`GET /lab-request-templates/:id?withTests=true` devolve os exames **na ordem do template**, cada um
anotado com `pageBreakBefore`. (Consulta o join explicitamente ordenado por `display_order`; DTO =
campos do exame + `pageBreakBefore`.) `pnpm generate` → tipos TS.

## 4. applyTemplate (front, `lib/lab-request-apply.ts`)
Remove TODO o sort/agrupamento. Emite os exames **na ordem do template**; insere uma **linha em
branco antes** de cada exame com `pageBreakBefore=true`. Mantém só o filtro por sexo + dedup do
pedido importado. (O `examBlock`/justificativa continua igual.)

## 5. Criar template (UI — follow-up, não-bloqueante)
Admin reordena (drag) + toggle "quebra de página antes" por exame. Por ora a ordem/quebras são
**dados**, configurados por SQL.

## 6. População inicial (seed por template — A DEFINIR o default do layout com o Dr.)
`display_order`: ordem atual (lab alfabético, depois imagem). `page_break_before`: definir o padrão
por painel (ex.: lab sem quebra; cada exame de imagem com quebra = comportamento original "1 por
página"; depois o Dr. ajusta quais imagens compartilham página removendo quebras).

## Verificação (dev)
`migrate up` + `go build` verde; `pnpm generate`; carregar template → texto na ordem/quebras
corretas; gerar PDF. Dev-only; prod sob ordem.
