# Plano — Reabilitar a geração de tipos (Regra de Ouro nº1) e migrar o front pros tipos gerados

**Status:** em execução · **Início:** 2026-06-13 · **Dono:** Claude (sob ordem do Dr.)

## Contexto / causa raiz (descoberta em 2026-06-13)

A Regra de Ouro nº1 (`Go models → swag → OpenAPI → tipos TS/Zod gerados`) estava **inoperante há tempo**:

1. **`swag init` quebrado:** faltavam entradas no `go.sum` p/ as deps de build do próprio swag
   (`sigs.k8s.io/yaml`, `urfave/cli/v2`). `swag` nunca rodava.
2. **`generate-all.sh` engolia o erro:** cada passo terminava em `|| echo "⚠️ ..."`, então a falha
   do swag não abortava — os passos seguintes rodavam sobre um `swagger.json` **velho**. Resultado:
   tipos gerados congelados, sem ninguém perceber.
3. **Consequência cultural:** virou hábito criar tipo TS **à mão**. Hoje `apps/web/lib/api/` tem
   51 clients, **47 com `interface` própria**; só 5 importam `@plenya/types`. O próprio
   `@plenya/types/index.ts` é interface à mão (Method, ScoreItem, SubscriptionPlan…) e **não
   expõe** os tipos de model gerados (`api-types.ts`/`components`), só re-exporta os Zod schemas.
4. **Artefatos gitignored:** `packages/types/src/generated/` está no `.gitignore`. O Dockerfile do
   web faz `COPY . .` e **não** roda `generate` → num clone limpo os gerados não existiriam. Build
   sobrevive por contexto, não por design (fragilidade latente).

## Já feito (foundation) ✅

- `go get -tool github.com/swaggo/swag/cmd/swag@<pin>` → swag vira **tool directive** versionado;
  `go.sum` completo. Roda via `go tool swag`.
- `scripts/generate-all.sh` reescrito **falha-dura** (`set -euo pipefail`, sem `|| echo`), usando
  `go tool swag`.
- `pnpm generate` roda ponta a ponta (EXIT 0). Confirmado: `models.LabTestDefinition` gerado tem os
  **30 campos** do model (inclui `sexApplicability` + `requestJustification`) — **superset** do
  manual de 13. Qualidade do gerado é boa o bastante p/ substituir os manuais.

## Progresso (2026-06-13)

- ✅ **Gerador consertado** (swag tool + falha-dura). `pnpm generate` roda ponta a ponta.
- ✅ **Zod gerado APOSENTADO** (passo 4 removido; `api-schemas.ts` apagado; `index.ts` não
  re-exporta mais). Ganho colateral enorme: **tsc do web caiu de 1475 → 216 erros** (os ~1260
  a mais eram lixo do Zod gerado quebrado poluindo todo typecheck). `@plenya/types` agora
  typecheck **limpo** (era impossível antes).
- ✅ **Camada de aliases** `packages/types/src/models.ts` (base = schema gerado; `WithRequired`
  refina só os campos sempre-presentes na resposta). `index.ts` expõe `components/paths/operations`.
- ✅ **Migrados (consomem o gerado, 0 erro novo):** `lab-request-templates.ts`
  (LabTestDefinition, LabRequestTemplate), `medication-definitions.ts` (MedicationDefinition),
  + bug pré-existente `Patient` ausente no `index.ts` corrigido.

## Progresso (2026-06-14) — Trilha A concluída + cascatas de consumidor

**web tsc: 231 → 115** (≈ −116; package `@plenya/types` compila com **0 erros**). 11 commits limpos
(`3e1a5df0`..`cf9190bb`), todos LOCAIS no master (sem push; aguarda ordem).

- ✅ **Trilha A 100% migrada:** todos os `lib/api/*.ts` com interface duplicada agora importam do
  gerado via `@plenya/types`. Inclui dedup do `index.ts` (Method/MethodLetter/MethodPillar +
  Score Group/Subgroup/Item/Level — eram duplicados e DIVERGENTES: `index.ts` tinha `level:string`
  e `ageMin/ageMax`; o gerado tem `level:number` e `ageRangeMin/Max`). `score-api.ts` idem.
- ✅ **Backend anotado (fechou gap p/ o front):** handlers de prescrição (Create/GetByID/List/Update)
  ganharam godoc `dto.PrescriptionResponse` → swag passou a emitir `dto.PrescriptionResponse`,
  `models.PrescriptionStatus`, `dto.Create/UpdatePrescriptionRequest`. Era o ÚNICO gap real de
  `components['schemas']['models.X']` em todo o front.
- ✅ **Bugs de RUNTIME que o `any` mascarava (achados pela tipagem concreta):**
  - `prescriptions.ts` e `medication-definitions.ts` tinham `.data` residual de axios sobre o body
    já desembrulhado (`apiClient` devolve `response.json()`) → funções retornavam `undefined`.
    **Lista de prescrições aparecia SEMPRE vazia**; **página /admin/medication-definitions quebrava**
    no `.filter`. `signedPDFPath` (front) ≠ `signedPdfPath` (json Go) → lia `undefined`.
  - `apiClient.get` só aceita 1 arg; vários passavam `{params}` (2º arg) — trocado por querystring.
- ✅ **Cascatas RHF/zod resolvidas** (padrão: `.default()`/`.transform()` faz input≠output do
  zodResolver): LabTestDefinitionForm (22→0, via `useForm<input,ctx,output>`), LabResultBatchForm
  (10→1, schema é `z.preprocess` → `zodResolver(schema as any)`), prescrição new/edit (via mesmo cast).
- ✅ **Refinos de alias:** `ScoreLevel.operator` → union; `PatientScoreSnapshot.itemResults/groupResults`
  → obrigatórios; `Prescription`/`MedicationResponse`, Score/Method graph com relações aninhadas.
- ✅ **React 19:** reintroduzido `import type { JSX } from 'react'` onde se anota `JSX.Element`.

### Restante (115) — categorizado
- **~67 consumidores da migração** (Score charts, anamnesis pages, lab-results pages, articles) —
  espalhados 1–3 por arquivo. **Próximo alvo** se for pra continuar zerando.
- **~33 dívida PRÉ-EXISTENTE sem relação com tipos** (training pages, OAuth buttons, calendar.tsx,
  rich-text-editor, tailwind.config, admin/users, appointments) — sempre existiram; `ignoreBuildErrors`
  as tolera. Fora do escopo da Regra de Ouro.
- **~4** resolução de módulo `@plenya/ui/score-radar` (build do package `packages/ui`, não tipos).
- **~11** recharts `Formatter` / `implicitly any` (libs externas / `noImplicitAny` pré-existente).

### Trilha B (gap de anotação swag) — ainda PENDENTE
Mesma técnica do prescription handler resolve: `models.Patient` sem `cpf`/`city`/`zipCode`,
`models.User` sem `cpf`, `LabResult`, `Article`. Anotar o handler/DTO Go → `pnpm generate` → migrar.

### Duas trilhas descobertas (das 28 interfaces com homônimo gerado)

**Trilha A — limpas (gerado é superset):** migram direto (alias + import). Ex.: MedicationDefinition,
LabTestDefinition, LabRequestTemplate. Restantes prováveis: Lead/LeadActivity, Anamnesis(+Item),
AnamnesisTemplate(+Item), Article, LabResultView(+Item), LabResult/LabResultValue, ScoreGroup/
Subgroup/Item/Level (cuidado: ScoreItem etc. estão duplicados à mão também no `index.ts` → dedup),
PatientScore*.

**Trilha B — gerado INCOMPLETO (gap de anotação swag no model Go):** migrar agora perderia campo.
- `models.Patient` **não expõe `cpf`/`city`/`zipCode`** (campos que a API retorna e o front usa).
- `models.User` **não expõe `cpf`**.
- Causa provável: json tag/`swaggerignore`/criptografia no model Go. **Correção é na FONTE**
  (anotar o campo no Go p/ o swag emitir), depois migra. NÃO contornar no TS.

## A fazer

### Fase 1 — Versionar e expor os gerados
- [ ] Remover `packages/types/src/generated/` do `.gitignore`; **commitar** `api-types.ts` +
      `api-schemas.ts` (tornam-se reviewáveis e o build deixa de depender de sorte).
- [ ] Criar `packages/types/src/models.ts` com **aliases ergonômicos** dos schemas gerados:
      `export type LabTestDefinition = components['schemas']['models.LabTestDefinition']` etc.
      (1 alias por model usado no front). `index.ts` passa a `export * from './models'`.
- [ ] Migrar as **interfaces à mão do `index.ts`** (Method, ScoreItem, ScoreLevel, SubscriptionPlan,
      PatientSubscription, Notification…) p/ aliases gerados; manter o nome do tipo p/ não quebrar
      imports. Onde o gerado divergir do que o front espera, **corrigir a anotação swag no handler/DTO
      Go** (fonte), nunca remendar no TS.

### Fase 2 — Migrar os 47 clients à mão (em lotes, com verificação)
- [ ] Inventariar cada `interface` local em `apps/web/lib/api/*.ts` → mapear p/ o schema gerado
      equivalente (model vs request/response DTO).
- [ ] Lote a lote: trocar a `interface` local por `import type { X } from '@plenya/types'`.
      Começar pela feature **lab-requests** (já tocada hoje) como referência.
- [ ] Para DTOs de request/response sem equivalente gerado: garantir que o handler Go tenha o DTO
      anotado p/ swag (gera o schema), depois importar o gerado.
- [ ] `ignoreBuildErrors:true` mascara erro de tipo → **typecheck explícito** por lote
      (`tsc --noEmit` no web) além do build.

### Fase 3 — Garantir que não regrida
- [ ] Rodar `pnpm generate` no fluxo certo (documentar no CLAUDE.md que é obrigatório após mudar model).
- [ ] (Opcional) checagem de CI/pre-commit: `pnpm generate` + `git diff --exit-code` em
      `packages/types/src/generated/` → falha se alguém mudou model e não regenerou.
- [ ] Atualizar a Regra de Ouro nº1 com a realidade: gerado é fonte; manual só p/ tipos que **não**
      vêm do backend (estado de UI, libs externas).

## Verificação
1. `pnpm generate` EXIT 0 e `git diff` mostra os gerados atualizados.
2. `tsc --noEmit` no web sem **novos** erros após cada lote.
3. Build do web (Docker) verde com os gerados versionados.
4. Smoke das telas migradas (lab-requests primeiro).
5. Prod só sob ordem explícita.

## Riscos
- Migração em massa pode revelar divergências model↔front mascaradas pelo `ignoreBuildErrors`.
  Mitigação: lotes pequenos + `tsc --noEmit` por lote.
- Versionar os gerados aumenta o diff dos PRs (aceitável; é o ponto — reviewabilidade).
- Corrigir anotações swag pode exigir tocar muitos handlers Go (fonte real da incompletude).
