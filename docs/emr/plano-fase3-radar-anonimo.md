# Fase 3 — Radar anônimo (light/triagem) por pilar

Objetivo: o resultado público do escore-light/triagem (`/escore-plenya/resultado/[code]`) passa a
renderizar o radar **por pilar AGIR** via `buildAgir` + `@plenya/ui` `RadarAgir` (mesmo componente do
EMR), em vez do radar por GRUPO atual. Fallback por grupo para sessões antigas (sem per-item result).

## Por que precisa de uma tabela nova
O scoring anônimo (`SubmitAnonymousScore`/`evaluateLightItem`) hoje só persiste **agregados**: snapshot
(total) + group_results (por grupo). `buildAgir` precisa de `itemResults[].item.methodPillars` (por-item,
com pilar). Logo, persistimos o resultado por-item (espelhando `patient_score_item_results` do EMR).

## Backend
1. **Model `AnonymousScoreItemResult`** (`anonymous_score_item_results`) — mínimo: id(UUIDv7), snapshot_id
   (FK CASCADE), item_id, group_id, status, level_number, level_matched_id, value_used, max_points,
   actual_points, timestamps+soft-delete. Relation `Item *ScoreItem` (p/ Preload MethodPillars). Sem
   FK em item_id/group_id (acoplamento leve; Preload usa só a tag foreignKey).
2. **`AnonymousScoreSnapshot`** += relation `ItemResults []AnonymousScoreItemResult`.
3. **Migration 00024** cria a tabela + índices (snapshot/item/group/deleted_at).
4. **`evaluateLightItem`** passa a receber `groupID` e RETORNAR `*AnonymousScoreItemResult` (não-nil só
   p/ avaliados). O loop de `SubmitAnonymousScore` coleta os results; após criar o snapshot, seta
   SnapshotID e faz bulk insert. (Só persiste **avaliados** — é o que o radar usa.)
5. **Payload público**: DTOs `PublicItemResult{status,actualPoints,maxPoints,item:{methodPillars:[{id,
   name,order,letter:{code,name,color,order}}]}}` + `ItemResults` em `PublicSnapshot`. `toPublicSession`
   mapeia. `loadSessionByPublicCode` += Preload `Snapshot.ItemResults.Item.MethodPillars(.Letter)`
   (deleted_at IS NULL, order ASC). Shape idêntico ao `AgirSnapshotInput` do `@plenya/ui`.

## Frontend (site)
6. `lib/score-light/types.ts`: += tipos PublicMethodLetter/Pillar/ItemResult + `itemResults?` em
   PublicSnapshot.
7. `components/escore/escore-light-resultado.tsx`: `const agir = buildAgir(snapshot)`. Se `agir` != null
   (há pilares) → `<RadarAgir letters pillars globalScore={round(totalPct)} labels/>` (PT/EN glyph A·C·T·S).
   Senão → fallback atual (`ScoreRadarChart` por grupo). Sessões antigas (sem itemResults) caem no fallback.

## Compat / deploy
- Sessões antigas: snapshot sem itemResults → `buildAgir` retorna null → fallback por grupo. Sem backfill.
- PROD: o radar por pilar só "acende" quando a distribuição AGIR (score_item_method_pillars) estiver em
  prod (workstream separado). Sem pilares → fallback por grupo. O código cobre os dois casos.
- dev verificado: 83/83 itens light têm methodPillars.

## Verificação
go build 0 · migrate up (00024) · submissão anônima POST → payload tem `snapshot.itemResults[].item.
methodPillars` · site resultado renderiza RadarAgir por pilar (Playwright) · fallback por grupo quando
sem pilares. Tipos gerados re-sincronizam no `pnpm generate` de CI (swag não roda no env).
