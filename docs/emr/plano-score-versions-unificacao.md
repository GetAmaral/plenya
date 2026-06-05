# Plano — Score Versions + Campos de site + Gerador + Radar único + Conversão anônimo→paciente

**Status:** PLANO PARA DISCUSSÃO (2026-06-05). Nada executado. Fonte: estudo de 5 frentes.

## Objetivo (Dr.)
1. Radar do site `/escore` deixar de ser 90° fixo → proporcional (mesmo do EMR).
2. Backend `score_version` (1→N) → `score_version_item` (M2M com `score_item`, com ordem). Tela de
   config. Criar Triagem + Light dos itens atuais. Depois **remover `is_light_version`**.
3. Campos novos: **legenda de site** em `score_item` e `score_level` (mensagem leiga ≠ clínica);
   **tipo de render no site** em `score_item` (≠ anamnese/exame do EMR).
4. **Gerador** que produz os configs estáticos do site (triagem/light) a partir da fonte de verdade (EMR).
5. **Radar sempre pela mesma base, por pilar** (do `score_item`), em todo lugar.
6. Resultado continua **anônimo**, mas com **conversão anônimo → escore real no prontuário** quando virar paciente.

## Estado atual (do estudo)

**Models** — `score_item`: `IsLightVersion`(bool), `LightOrder`(*int), `LightQuestion`(*string),
`PrepOrder`, `PatientExplanation`, `ClinicalRelevance`, `Conduct`, `Points`, `Order`, `MethodPillars`(M2M).
`score_level`: `Level`, `Name`, `LowerLimit/UpperLimit`, `Operator`, `PatientExplanation`, `ClinicalRelevance`, `Conduct`.
**Não existe** campo de legenda de site nem de tipo-de-render (o tipo é INFERIDO).

**Render hoje (site `escore-light-form.tsx`)** — 4 tipos inferidos: PHQ9 (widget), NumericClassifier
(numérico com auto-classificação por ranges), LevelChoice (grid de botões), numérico fallback.
Resposta salva em `SessionResponse {scoreItemId, numericValue?, selectedLevel?, textValue?}`.

**Anônimo** — `anonymous_score_item` guarda só a resposta crua (numeric/text/selectedLevel).
`anonymous_score_snapshot` guarda só totais + `anonymous_score_group_result` (por GRUPO). **Não há**
item-result por pilar. Claim (`RequestClaim`→magic link→`ConfirmClaim`) só liga `claimed_by_patient_id`;
**não converte** em `patient_score_snapshot`.

**Engine** — `CalculateSnapshot` (paciente) recalcula de anamnese+lab → `PatientScoreItemResult` +
agregação. `actual = Points * GetLevelMultiplier(level)` (0→0, 1→.2, …, 5/6→1.0). Helpers compartilhados:
`GetLevelMultiplier`, `SortLevelsAsc`, `AppliesToPatient`, `EvaluatesTrue`, `matchLevelForResponse`.

**Radar** — Site `RadarAgir` = setor FIXO 90° (`SECTOR_MID`). EMR `RadarAgir` = proporcional (piso 10°/letra).
EMR é superior. `buildAgir` (web) é a fonte única por pilar.

**Gerador** — `sync-score-light.ts` (apontado no package.json) chamava `GET /api/v1/score-light/config`
(`BuildLightConfig`, filtra `isLightVersion=true`) e gravava o JSON. Configs hoje são estáticos
commitados (light=83 itens/10 grupos; triagem=36/8, curadoria manual, subconjunto do light).

---

## Modelos novos

```
ScoreVersion (score_versions)
  id, name ("Triagem"/"Light"), slug ("triagem"/"light"), description?,
  site_intro?  (texto da página), order, active(bool), timestamps, soft-delete
  Hooks: BeforeCreate UUIDv7

ScoreVersionItem (score_version_items)   -- M2M version↔item, com ordem
  id, version_id (FK→score_versions), score_item_id (FK→score_items),
  display_order (int), timestamps, soft-delete
  UNIQUE (version_id, score_item_id)
  Hooks: BeforeCreate UUIDv7
```
A version **só seleciona + ordena itens**; grupo/subgrupo/pilares vêm do `score_item` (não duplica).

## Campos novos em score_item / score_level

`score_item`:
- `site_render_type` varchar — enum explícito do input do site: `level_choice | numeric_classifier |
  boolean | text | slider | scale_0_3 (PHQ9-like)`. (Hoje é inferido; passa a ser configurado.)
- `site_explanation` text? — explicação leiga genérica (≠ `PatientExplanation` clínica, ≠ `LightQuestion`).
- `site_question` text? — o **prompt leigo** da pergunta no site (substitui o uso de `LightQuestion`).

`score_level`:
- `site_legend` text? — o que aquele nível significa em linguagem leiga (≠ `PatientExplanation`).

`light_question`/`light_order`/`is_light_version`: **a remover** depois da migração (substituídos por
`site_question`/`score_version_item.display_order`/`score_version`). `prep_order` fica (é outra coisa).

---

## Fases (para discutir ordem/escopo)

### Fase 0 — Radar único (pacote compartilhado) + fix do site
- Criar pacote `packages/ui` → `@plenya/ui` (slot reservado) com `RadarAgir` + `buildAgir` +
  `buildAgirMock`. Cores de `@plenya/brand/tokens` (gold/petrol/ocean/sage/cream) e **estilos inline**
  (sem `text-foreground`/`bg-card` — portável). Tailwind `content` dos 2 apps inclui o pacote.
- EMR (#1-4) trocam import local pelo pacote.
- Site marketing (`/escore-plenya`, `score-section`) usa `RadarAgir` + `buildAgirMock(agir-structure)`
  → **vira proporcional** (resolve o 90° fixo). `agir-structure.ts` já está sincronizado (42 pilares).
- Site light result ainda por grupo (até a Fase 3).

### Fase 1 — Campos de site nos models
- Migration goose: `score_items += site_render_type, site_explanation, site_question`;
  `score_levels += site_legend`. Backfill `site_question = light_question`.
- Models + DTO + admin (`ScoreItemDialog`/`ScoreLevelDialog`) ganham os campos. `pnpm generate`.

### Fase 2 — score_version + config + gerador (e remover is_light_version)
- Models `ScoreVersion` + `ScoreVersionItem` + migration.
- **Seed:** criar version "Light" (itens `is_light_version=true`, ordem `light_order`) e "Triagem"
  (os 36 da curadoria atual) como `score_version_item`.
- **Tela de config** `/(authenticated)/scores/versions`: lista de versions; por version, multi-select
  de itens da árvore + reordenar (`display_order`, dnd-kit). CRUD `ScoreVersion`/`ScoreVersionItem`.
- **Gerador:** endpoint `GET /api/v1/score-versions/:slug/config` (monta `LightConfig` a partir de
  `score_version_item` + `score_item` com campos de site + levels) + script `sync-score-versions`
  que grava `apps/site/content/data/<slug>-config.json`. Substitui o `sync-score-light.ts` sumido.
- Depois que os configs saem de version: **drop `is_light_version`, `light_order`, `light_question`**.

### Fase 3 — Light por pilar + radar unificado no site
- Nova tabela `anonymous_score_item_result` (espelha `patient_score_item_result` mínimo:
  snapshot_id, score_item_id, status, level_number, actual/max_points). Persistir no `CreateSession`.
- `GetSessionByPublicCode`: `Preload("Snapshot.ItemResults.Item.MethodPillars.Letter")`.
  DTO público expõe `itemResults`.
- Site light result usa `@plenya/ui` `RadarAgir` + `buildAgir(snapshot)` → **por pilar**. Fallback por
  grupo para sessões antigas (sem itemResults).
- EMR portal `/escore-light`: renderizar o radar inline (em vez de só link).

### Fase 4 — Conversão anônimo → paciente
- `ScoreSnapshotService.ConvertAnonymousSessionToPatientSnapshot(sessionID, patientID, userID)`:
  carrega respostas anônimas → reavalia via engine (`matchLevelForResponse`/`GetLevelMultiplier`) →
  cria `PatientScoreSnapshot` + `PatientScoreItemResult[]` (status evaluated p/ os respondidos;
  resto `no_data_available`) + `PatientScoreGroupResult[]`, em transação. `CalculatedByUserID` = o user.
- Disparo: **manual** ("Importar pro prontuário") no claim/portal, ou no `ConfirmClaim`. (A decidir.)
- Resultado vira snapshot REAL no prontuário (parcial — é triagem), aparece em `/health-scores`.

---

## Decisões (FECHADAS pelo Dr., 2026-06-05)
1. **Render type explícito** (`site_render_type`) — SIM, campo explícito (não inferir mais).
2. **Campos de site no `score_item`** — o item é source of truth; `score_version` só **seleciona+ordena**
   (sem override por version).
3. **`site_question` substitui `light_question`** — MIGRAR (não manter os dois).
4. **Conversão anônimo→paciente: MANUAL** — botão "Importar pro prontuário" (claim não cria registro
   clínico sozinho).
5. **Snapshot convertido é PARCIAL** — OK, marcar **origem: triagem/light**.
6. **Pacote:** `packages/ui` → **`@plenya/ui`**.
7. **Ordem:** começar pela **Fase 0** (radar) e seguir pro backend.
