# Plano — Pontos + Níveis de risco para os genes (Escore Plenya)

**Status:** ✅ CONCLUÍDO · aplicado em **dev + prod** 2026-06-05.

## Resultado (2026-06-05)
- **361 genes** com `points` + escada de níveis (genótipo→nível); **1041 níveis** (era 366).
- Genética: **26 → 851 pts = 8,9%** do escore total (3ª maior atrás de Exames 45,5% e Histórico 20,5%;
  o Dr. escolheu o tier mais leve T2=2/T1=1 pra não pesar mais que os grupos modificáveis).
- **dev ≡ prod byte-idêntico:** `levels_md5=4ac0622…`, `points_md5=2136397…`, 1041/1041 conferidos.
- Prod aplicado via **dry-run-em-prod com ROLLBACK** (txn aplicada→hashes conferidos→rollback→aplicada de verdade);
  backup durável `/tmp/plenya_prod_PREgenes_20260605_232336.dump` (219M) no host.
- **APOE-ε** refeito (estava invertido) → proteção=5 … risco=0, 6 níveis distintos.
- Artefatos canônicos: `genetica-pontos-niveis.sql` (apply idempotente) · `gen-genetica-pontos-niveis.py`
  (gerador determinístico) · `genetica-pontos-niveis-plano.csv` (auditoria por gene).
## Captura: gene é EXAME (não anamnese) — 2026-06-06
Decisão do Dr.: gene é exame laboratorial **qualitativo** (genótipo, não número), capturado via
`lab_results`, não anamnese. Descobertas:
- O motor, no caminho de exame, avaliava **só numérico** (`result_numeric` + `EvaluatesTrue`); ignorava
  `result_text` e `lab_results.level`. Exame categórico de verdade (retinopatia, angio-TC) **nunca casava**.
- **Fix de backend (2 arquivos):**
  - `score_snapshot_service.go`: no ramo de lab, se `ResultNumeric==nil` mas `Level!=nil`, casa o
    `score_level` pelo **número do nível** (`lab_results.level` pré-casado na entrada). Direção alto=bom.
  - `lab_result_repository.go` (`GetHistoricalResultsByLabTestCode`): filtro relaxado de
    `result_numeric IS NOT NULL` → `(result_numeric IS NOT NULL OR level IS NOT NULL)` p/ qualitativo
    chegar ao scorer. Blast radius=0 (nenhum exame não-gene tinha level sem numeric).
  - Bônus: conserta TODOS os exames categóricos que estavam quebrados.
- **Dados:** cada gene virou `lab_test_definition` (`result_type=categorical`, categoria `genetics`,
  code `GEN_<id-hex>` único/determinístico) + `score_items.lab_test_code` setado (361/361).
  SQL `genetica-exames-labdefs.sql`. Os `score_levels` (genótipo→severidade) são as **opções** do exame
  E a escada de pontuação.
- **Captura do resultado:** `lab_results` guarda `result_text` (genótipo "AT") + `level` (severidade).
  Validado E2E no dev: FTO genótipo AT (level 2) → `data_source=lab_result`, actual=8×0,4=**3,20**. ✓
- **Pendência (frontend):** a tela de entrada de exame ainda precisa, p/ teste categórico/genótipo,
  listar as opções (score_levels) e gravar `result_text`+`level`. Backend já aceita (DTO tem `level`).

**Histórico:** em execução (dev) · 2026-06-05
**Objetivo:** fazer os ~361 itens do grupo **Genética** efetivamente pontuarem no Escore,
atribuindo `points` (peso) **e** uma escada de `score_levels` (genótipo → nível) por gene,
via análise individual calibrada pelos genes que já têm peso. Depois migrar dev → prod.

## Como o motor pontua (descoberto em `score_snapshot_service.go`)
- Contribuição de um item = `points × multiplicador(nível_casado)`.
- `GetLevelMultiplier`: nível **0→0% · 1→20% · 2→40% · 3→60% · 4→80% · 5→100% · 6→100%**.
- Item de genética é preenchido via **anamnese**: `anamnesis_items.numeric_value` **é** o número do
  nível (o motor casa `score_levels.level == numeric_value`). Os `score_levels` de um item são, ao
  mesmo tempo, **as opções de resposta do questionário** e a **escada de pontuação**.
- **Consequência crítica:** 360/361 genes hoje têm **só nível 0** → multiplicador 0 → contribuem
  **0**, qualquer que seja `points` (inclusive ACE=10, ABCC8=8, ABCA1=8). Só o APOE-ε tinha 0–5.

## Convenção de direção (decisão do Dr., 2026-06-05)
**Nível ALTO = bom, nível BAIXO = ruim.** O genótipo favorável casa nível alto (multiplicador alto
→ ganha quase todo o `points` do item → escore sobe); o genótipo de risco casa nível 0 (0% → não
ganha pontos → escore cai). Não é obrigatório usar os 6 níveis: pode ser binário (5 bom / 0 ruim),
ternário (5 / 2 / 0) etc.

> ⚠️ O **APOE-ε** existente está **invertido** (E2/E2-proteção=nível 0, E4/E4-risco=nível 5). Será
> **refeito** nesta convenção (proteção→alto, risco→0).

## Escada de níveis (templates, direção alto=bom)
Aleo de risco `R`, alelo favorável/normal `P`:
- **Ternário aditivo/codominante (default p/ SNP comum):** `PP → 5` · `PR (het) → 2` · `RR → 0`.
  (carregar 1 alelo de risco custa 60% dos pontos do item.)
- **Só "RR=risco" na legenda (homozigoto de risco nomeado):** `não-RR → 5` · `het → 3` · `RR → 0`
  (ou binário `não-RR → 5 / RR → 0` quando het não for distinguível).
- **Variante rara deletéria ("Mutação=deficiência", "Expansão"):** `não-portador → 5` · `portador → 0`.
- **Variante rara protetora ("R46L=proteção", "A673T=proteção rara"):** `portador → 5` · `não-portador → 3`
  (não-portador é a norma, não penaliza forte).
- **Trade-off sem bom/ruim (ACTN3 poder/resistência, CYP1A2 rápido/lento):** tratados caso a caso —
  default informativo (níveis neutros ~3–4, `points` baixo) para não premiar nem penalizar. Marcados ⚑.

Cada `score_levels.name` mantém a legenda do genótipo (ex.: `5: TT (favorável)`, `2: AT`, `0: AA (risco)`).

## Calibração de `points` (tiers, ancorados nos existentes ACE=10 · ABCC8=8 · ABCA1=8)
- **T3 — maior, replicado, acionável, desfecho duro (cardiometabólico/neuro):** 7–10.
- **T2 — moderado, replicado, modificável (lipídico, ômega, vit D/B12, inflamação, óssea, detox c/ estilo de vida):** 4–6.
- **T1 — menor/informativo, efeito fraco ou baixa acionabilidade (performance, capilar/GeneHair, várias enzimas de detox):** 1–3.
- **T0/⚑ — trade-off puro:** 0–2 com níveis neutros.

**Peso agregado da Genética (knob a aprovar):** hoje 26 pts (0,3% do total ~8759). A calibração
proposta mantém a maioria em T1 (1–3) e poucos em T3, mirando a Genética em **~6–8%** do total
(não deve dominar Exames/Histórico). O share resultante é mostrado na amostra antes do run completo.

## Direção já disponível nos dados
337/361 genes já trazem a direção do risco na legenda (`A=risco T=proteção`, `TT=risco`,
`G=↓HDL A=normal`). **24** precisam de lookup clínico (APOE rs429358/rs7412, PCSK9 R46L, ADH1B,
CYP1A2, SOD2, IL6, MCM6, APP, C9orf72, ACTN3 + 13 GeneHair capilar).

## Execução
1. **Framework + amostra (~15 genes)** cobrindo todo tier/subgrupo/tipo-de-legenda → aprovação do Dr.
   (calibração de `points`, nível do het, share agregado, trade-offs). ← gate atual.
2. **Geração completa:** tabela estruturada por gene (id, alelo de risco, tier, points, mapa genótipo→nível)
   → SQL idempotente (`UPDATE points` + `INSERT/DELETE score_levels`). Reaproveita UUIDs fixos dos genes.
3. **Aplicar no dev**, recomputar snapshot do paciente demo p/ ver impacto, conferir share por grupo.
4. **Migrar p/ prod** pelo método dry-run-em-clone (backup → restore em `*_dryrun` → aplicar c/
   `ON_ERROR_STOP=1` → comparar counts+md5 c/ dev → aplicar no real). Ver [[agir_distribuicao_bioma_status]].

## Anti-padrões
- Não dar `points` alto à maioria (Genética dominaria o escore).
- Não copiar a direção invertida do APOE-ε.
- Não inventar direção de risco para os 24 sem legenda — fazer lookup (estabelecido na literatura).
- Não tratar trade-off (poder/resistência, rápido/lento) como bom/ruim.
