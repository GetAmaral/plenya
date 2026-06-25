# Plano — Importação de laudos por PDF: 6 correções/melhorias

> Plano APROVADO (plan mode). Cópia versionada de `~/.claude/plans/jolly-hatching-wall.md`.
> Implementação a iniciar quando o usuário mandar. Ordem: Fase 0 → 1 → 2 → 3 → 4 → 5.

## Context
Ao importar exames do paciente João pela tela "Importar Laudo via PDF" (lote de resultados), o
usuário encontrou 6 problemas. A investigação (backend `apps/api` + frontend `apps/web`) confirmou
as causas-raiz, incluindo **um bug de perda silenciosa de dados** (editar a data do lote apaga os
exames classificados). Decisão do usuário no #6: vários PDFs viram **1 lote por PDF**, em massa.

## Fluxo atual
Cria lote (`POST /lab-result-batches`, lab+data manuais NOT NULL) → upload de 1 PDF
(`POST /:batchId/upload-pdf`) → `ProcessingJob` async: OCR → Claude (`InterpretLabResult`) →
`createLabResultsFromJSON` (cria `LabResult` + *matching* com `LabTestDefinition`). A atribuição de
NÍVEL (`ClassifyBatchResults`) só roda no botão "Classificar" e no `Update`.

## Os 6 problemas e causa-raiz (verificada)
1. **Não classifica ao importar.** `processJob` termina em `processing_job_service.go:212` sem chamar
   `ClassifyBatchResults`. O matching roda; a atribuição de `Level`/criticidade não.
2. **Não traz laboratório + data de coleta do PDF.** Prompt (`ai_service.go` buildPrompt) e
   `dto/pdf_extraction.go` só extraem exames. `pre_matching_service.go` tem regex p/ lab+data mas está
   instanciado e **nunca chamado**. `LaboratoryName`/`CollectionDate` do lote são manuais (NOT NULL).
3. **🔴 Editar data do lote apaga os exames (perda de dados).** Dois filtros inconsistentes no front:
   o schema (`lib/validations/lab-result-batch.ts:22-29`) mantém linhas com `definição OU nome OU valor`,
   mas `formToApiValues:78` mantém **só linhas com `testName`**. O backend grava `TestName=""` nos
   exames **casados** (`processing_job_service.go:287`). Logo os casados são descartados do payload, e o
   `Update` (`lab_result_batch_service.go:401-408`) **deleta** todo result ausente do payload.
4. **Sem visibilidade.** Só existe `LabResult.Matched bool` (`lab_result.go:51`); o motivo do não-match
   não é persistido (só `fmt.Printf` em `processing_job_service.go:309`). Sem resumo do job
   (extraídos/casados/falhos). `UnclassifiedResultsAlert` mostra contagem, não os itens/razão.
5. **Sem acesso ao PDF original.** PDF salvo em `/app/uploads/lab-result-batches/{id}_{ts}.pdf`, path em
   `ProcessingJob.PDFPath`. Não há endpoint GET nem link na UI. Padrão a reusar:
   `issuedDocs.Get("/:docId/pdf", DownloadPDF)` em `main.go:947`.
6. **Vários PDFs = lotes separados, 1 por vez.** Handler usa `c.FormFile("file")` (1 arquivo);
   `PDFUploadZone.tsx` `maxFiles:1`. `createLabResultsFromJSON:263` deleta results existentes do lote
   (irrelevante no modelo 1-lote-por-PDF).

## Abordagem (faseada)

### Fase 0 — Bug de perda de dados (#3) — URGENTE
- Frontend `lib/validations/lab-result-batch.ts`: corrigir `formToApiValues:78` para o mesmo critério do
  schema — manter linha com `id` OU `labTestDefinitionId` OU `testName` OU valor. Preserva os casados.
- Backend `lab_result_batch_service.go` `Update`: **remover o delete em massa (401-408)** → upsert-only.
  Remoção individual já tem rota dedicada (`DELETE /:batchId/results/:resultId`, `main.go:1018`).
- Recuperação: checar lote(s) do João que perderam exames e recriá-los de `batch.PDFContentJSON`
  (`createLabResultsFromJSON` + `ClassifyBatchResults`), via script Go one-off contra prod.

### Fase 1 — Classificação automática (#1)
- `processJob`: após `createLabResultsFromJSON` (antes do `return nil`, ~linha 211) chamar
  `s.labResultBatchService.ClassifyBatchResults(batchID)` (já existe). Step de progresso. Botão
  "Classificar" vira re-classificar; UI refaz fetch ao concluir.

### Fase 2 — Extrair laboratório + data de coleta (#2)
- `ai_service.go` buildPrompt + `dto/pdf_extraction.go`: extrair `laboratorio` + `dataColeta`. No
  `processJob`, sobrescrever lab/data placeholder do lote. `pre_matching_service.go` (regex) como
  fallback. Tornar lab/data opcionais na importação (placeholder + preenchimento pós-extração; sem
  migration de NOT NULL).

### Fase 3 — Visibilidade: status + razão por item + resumo (#4)
- Migration goose: `lab_results.match_reason text` (+ opcional `source` pdf|manual). Model + DTO +
  `pnpm generate`. `createLabResultsFromJSON` grava razão ("não catalogado", "valor não numérico"…) e
  origem. Resumo do job (extraídos/casados/não-casados/falhos). UI lista itens sem casar/classificar e
  o porquê; marca origem PDF vs manual.

### Fase 4 — Acesso ao PDF original (#5)
- `GET /lab-result-batches/:id/pdf` (espelhar `DownloadPDF` de `main.go:947`): stream do
  `ProcessingJob.PDFPath` com RBAC staff + checagem de paciente. Flag `hasPdf` no DTO; botão "Ver PDF"
  em `[id]/page.tsx`, `[id]/edit`, `revisar`. Apagar arquivo no `Delete` do lote (hoje vaza).

### Fase 5 — Vários PDFs, 1 lote por PDF (#6)
- Frontend `PDFUploadZone.tsx` + `lab-results/new`: aceitar múltiplos PDFs; por arquivo, criar lote +
  upload + processar; progresso consolidado ("3/5"). Cada lote pega lab/data do próprio PDF (Fase 2).
- Backend: reutilizar o pipeline por arquivo (1 `ProcessingJob` por PDF).

## Arquivos centrais
- Backend: `services/processing_job_service.go`, `lab_result_batch_service.go`, `ai_service.go`,
  `pre_matching_service.go`; `handlers/lab_result_batch_handler.go`; `models/lab_result.go`,
  `lab_result_batch.go`; `dto/pdf_extraction.go`, `dto/lab_result_batch.go`; `database/migrations/`
  (Fase 3); rotas `cmd/server/main.go:1004-1024`.
- Frontend: `lib/validations/lab-result-batch.ts`, `lib/api/lab-result-batch-api.ts`;
  `components/lab-results/{LabResultBatchForm,PDFUploadZone,ProcessingStatus,UnclassifiedResultsAlert,UnmatchedBadge}.tsx`;
  `app/(authenticated)/lab-results/{new,[id],[id]/edit,revisar}/page.tsx`.

## Verificação
- Build backend no container; goose up (Fase 3); `pnpm generate` após mudar models.
- Dev: criar lote + subir PDF → confirmar auto-classificação (níveis+criticidade) e lab/data do PDF.
- Regressão #3: editar a data do lote e salvar → exames (inclusive casados) permanecem.
- Razões dos não-casados na UI; "Ver PDF" abre o original; 3 PDFs → 3 lotes. Playwright local (memória
  `emr_qa_visual_playwright`). Deploy por-app com ordem explícita.
