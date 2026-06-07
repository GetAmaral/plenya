# Plano — Unificar o radar AGIR em TODO o monorepo

**Objetivo (Dr.):** TODOS os locais com radar de escore usam o MESMO mecanismo de geração
(`RadarAgir` proporcional + `buildAgir`), em web E site, incluindo escore-light/triagem.

## Censo (2026-06-05)

| # | Local | Componente hoje | Dados | Unificado |
|---|---|---|---|---|
| 1 | EMR `/health-scores` | RadarAgir (web) | buildAgir(snapshot) | ✅ |
| 2 | EMR `/health-scores/[id]` | RadarAgir (web) | buildAgir | ✅ |
| 3 | EMR `/patients/[id]` (capa) | RadarAgir (web) | buildAgir | ✅ |
| 4 | Portal `/escores` | RadarAgir (web) | buildAgir(latest) | ✅ |
| 5 | Site `/escore-plenya` (marketing) | RadarAgir **do site** | mock (agir-structure) | ❌ |
| 6 | Site home `score-section` | RadarAgir **do site** | mock | ❌ |
| 7 | Site `/escore-plenya/resultado/[code]` (light) | ScoreRadarChart **do site** | groupResults (por GRUPO) | ❌ |
| 8 | EMR portal `/escore-light` | nenhum (link p/ #7) | — | ❌ |

Mobile (pro/app) e packages: sem radar.

## Frente 1 — Pacote compartilhado (sem backend)

Extrair para um pacote (`packages/...`) consumido por web e site:
- `RadarAgir` (renderização — proporcional, piso 10°/letra, A no topo, paleta cravada).
  Trocar classes semânticas (`text-foreground`, `bg-card`, `heading-section`, `label-upper`)
  por hex/estilos próprios → portável entre os dois apps (themes diferentes).
- `buildAgir(snapshot)` (dados a partir de um snapshot com itemResults+methodPillars).
- `buildAgirMock(agirStructure)` (adapter p/ os radares de marketing do site: gera letras+pilares
  com scores fictícios determinísticos a partir de `agir-structure.ts`).

Migrações de consumo:
- EMR (#1-4): trocar import local pelo pacote.
- Site #5, #6: usar RadarAgir compartilhado + `buildAgirMock` (vira layout proporcional, igual EMR).
- Site #7: usar RadarAgir compartilhado — MAS depende da Frente 2 (precisa de dados por pilar).

Config: adicionar o path do pacote ao `content` do Tailwind dos dois apps.

## Frente 2 — Escore-light por pilar (backend)

Hoje o light é por grupo. Para o radar do light (#7 e #8) ser por PILAR (mesmo mecanismo):
1. Novo model `AnonymousScoreItemResult` (snapshot_id, score_item_id, actual/max_points, status).
2. `AnonymousScoreService.CreateSession`: persistir os item-results ao calcular.
3. `GetSessionByPublicCode`: `Preload("Snapshot.ItemResults.Item.MethodPillars.Letter.Method")`.
4. Migration goose (nova tabela).
5. Frontend: `escore-light-resultado.tsx` usa `buildAgir(snapshot)` → RadarAgir; fallback por grupo
   se vier sem pilares (sessões antigas).
6. EMR portal `/escore-light`: opcional renderizar o radar inline (em vez de só link).

Custo: ~1-2 migrations + model + service + preload + frontend.

## Decisões pendentes
- (A) Local do pacote compartilhado: `packages/ui` (reservado/vazio hoje) vs novo `@plenya/score-radar`.
- (B) Frente 2 (backend do light) agora ou depois.
