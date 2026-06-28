# Escalas — mostrar resultado real + persistir valor final + exibir em leitura

**Aprovado 2026-06-27.** Vale para TODAS as escalas do SCALE_REGISTRY (PHQ-9, GAD-7,
Epworth, IIEF-5, ASEX, Dubois imediato/tardio, Span direto/inverso).

## Problema
- Widget de preenchimento já calcula total + classificação ao vivo (OK).
- Mas (1) o **total não é gravado em `numericValue`** (só fica no JSONB `scaleResponses.total`);
  (2) as telas de **leitura** mostram o nome cru do score item (contém "____/25") e no máximo
  o chip "N{nível}", nunca o total computado.
- Semântica do projeto: `numeric_value` = medida crua, `selected_level` = classificação.
  Escala = total → numericValue; classificação → selectedLevel. O widget só esquecia o numericValue.

## Plano (A+B+C+D — escopo total)

### A. Persistir valor final
`AnamnesisTemplateItemsRenderer.tsx` → `handleScaleResult`: ao completar, setar
`numericValue = total` (limpar quando incompleta/limpa). Não afeta o motor de score (lê selectedLevel).

### B. Helper compartilhado em `@plenya/domain`
`formatScaleResult(total, maxScore, levelName?)` → `"22/25"` ou `"22/25 · Disfunção leve"`.
Pure, sem depender de tipos do web. Exportado no index do domain.

### C. Exibir em modo leitura (frontend)
Mostrar `total/maxScore · N{lvl} · nome` (via scaleResponses.total + getScaleDef.maxScore +
detectLevel/levels) em:
- `AnamnesisTemplateItemsRenderer.tsx` `currentChip` (linha colapsada da consulta).
- `app/(authenticated)/anamnesis/page.tsx` (cards de anamnese salva).
- `components/consultations/consultation-anamnesis-panel.tsx`.
- `app/(authenticated)/patients/[id]/prontuario/page.tsx` (resumo).

### D. Histórico / timeline (backend + frontend)
- `apps/api/internal/services/clinical_timeline_service.go`: incluir `ai.scale_responses` no SELECT.
- DTO `ClinicalTimelineEntry`: novo campo `scaleResponses`.
- `apps/web/lib/api/clinical-timeline.ts`: tipo + `AnamnesisItemHistory.tsx` formatValue mostra total.

## Notas
- Dev only, sem deploy (regra). Compilar Go no container; `pnpm generate` se o DTO entrar no swagger.
- Req "auto-preencher classificação ao preencher" já funciona no ScaleWidget — sem mudança.
