# Plano — Templates de exames com layout estruturado (opção C)

> ✅ **DEPLOYADO EM PROD (2026-06-18).** Commits `a2e54e15` (feature) + `28459442` (fix trocar
> paciente). migration 00040 aplicada no prod (goose=40). Dados (justificativas reescritas +
> categoria sono→imagem + layout 437 vínculos) aplicados por seed em transação; **prod ≡ dev**
> (fingerprints md5 idênticos de layout e justificativas). Backups pré-deploy: dev+prod
> `~/db-backups/plenya_db_*_20260618-004910.dump` (prod também na VPS). Deploy do web travou num
> blip de rede VPS→GitHub (clone timeout); fila limpa via coolify-db + redeploy → web novo no ar.

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

## 5. Editor de template (UI) — ✅ FEITO (2026-06-19)
Antes: o dialog usava `DualListSelector`, que reordenava os selecionados **por nome** (perdia a
ordem do template), não mostrava quebra de página nem justificativa, e ao salvar chamava o
replace-all que **zerava `display_order`/`page_break_before`**.

Agora:
- Novo componente `apps/web/components/lab-tests/template-tests-editor.tsx`: coluna de selecionados
  **na ordem do template**, com **drag-to-reorder** (`@dnd-kit/sortable`), **toggle "Nova página"**
  (`page_break_before`) por exame e **preview da justificativa** (read-only — é da definição, global).
- Backend: `PUT /lab-request-templates/:id/tests` aceita `tests: [{testId, displayOrder,
  pageBreakBefore}]` (campo `testIds` legado mantido). Novo caminho `UpdateLabRequestTemplateTestsOrdered`
  (handler→service→repo) faz DELETE+INSERT **gravando display_order + page_break_before** — não
  zera mais o layout.
- Verificado em dev: reordenar + marcar quebra → salvar → join table persiste ordem e `page_break_before`.

## 6. População inicial (seed por template — A DEFINIR o default do layout com o Dr.)
`display_order`: ordem atual (lab alfabético, depois imagem). `page_break_before`: definir o padrão
por painel (ex.: lab sem quebra; cada exame de imagem com quebra = comportamento original "1 por
página"; depois o Dr. ajusta quais imagens compartilham página removendo quebras).

## Verificação (dev)
`migrate up` + `go build` verde; `pnpm generate`; carregar template → texto na ordem/quebras
corretas; gerar PDF. Dev-only; prod sob ordem.
