# Análise — design do radar AGIR (o que mudou e quando)

**Contexto (2026-06-04):** o Dr. abriu a capa do paciente e reagiu ao radar AGIR — "cadê as
letras com separação por cor?". Investigação do que foi feito no **design do radar**.

## Conclusão (com prova no git)

- **Nesta sessão (2026-06-04): ZERO código alterado.** `git status` dos `.tsx/.ts/.go` voltou vazio.
  O único que mudou aqui foi **dado em dev** (taxonomia 13→42 pilares + snapshot de níveis
  aleatórios semeado no paciente de teste TTTEST), o que deixou o radar mais denso e fez a falta
  de cor saltar aos olhos. **O design do radar NÃO foi tocado nesta sessão.**
- **A mudança de design foi ANTES:** commit **`cb46eede` (19/04/2026)**, "feat(web): alinha EMR ao
  design system Plenya" (co-autoria Claude Opus 4.7). Mensagem literal:
  > "porta o **RadarAgir do site (SVG puro)** para **substituir o Recharts** no modo metodologia
  > do ScoreRadarChart."

  Esse commit **criou `RadarAgir.tsx` (385 linhas, SVG)** e **gutou `ScoreRadarChart.tsx`**
  (562 → 316 linhas, ~308 linhas removidas), trocando o radar.

## O que o radar ANTIGO tinha (removido em cb46eede)

A versão de `ScoreRadarChart.tsx` **antes** de `cb46eede` (562 linhas, Recharts) construía:
- `letterScores` / `letterPositions` com **`letterColor`** por letra (A/G/I/R);
- agrupava os pilares **por letra, cada um com a cor da sua letra** → o radar tinha
  **separação por cor por letra** (as regiões/séries coloridas), que é o que o Dr. quer de volta.

## O que o radar ATUAL faz (RadarAgir.tsx, SVG)

- Desenha um **polígono de uma cor só** (bege) com a forma do escore;
- só os **arcos externos** (A verde, G azul, I roxo, R laranja) e os **pontinhos dos vértices**
  são coloridos por letra;
- **NÃO** colore a região/setor interno por letra. Daí a impressão de "sem separação por cor".
- Consome `snapshot.itemResults[].item.methodPillars` (o M2M) → por isso reflete os 42 pilares.

## Onde está

- `apps/web/components/health-scores/RadarAgir.tsx` — o SVG atual (props: `letters[]`, `pillars[]`).
- `apps/web/components/health-scores/ScoreRadarChart.tsx` — radar da página `/health-scores`.
- `apps/web/app/(authenticated)/patients/[id]/page.tsx` — `buildAgir(snapshot)` (capa do paciente).
- Commit do design: `cb46eede`. **Já está no ar (prod) desde abril** — não é regressão desta sessão.

## Opções de conserto (a decidir — nada feito ainda)

1. **Restaurar o radar Recharts antigo** (com cor por letra) — reverter a lógica de `cb46eede` no
   `ScoreRadarChart`. Traz de volta exatamente o visual antigo, mas reintroduz a dep. Recharts no
   modo metodologia (o que `cb46eede` queria evitar).
2. **Colorir os setores do próprio `RadarAgir.tsx` (SVG) por letra** — pintar cada wedge/região
   A/G/I/R com a cor da letra (preenchimento translúcido), mantendo o SVG e o design system.
   **Recomendado** — resolve a queixa sem voltar pro Recharts.
3. Comparativo visual (antigo × atual × proposta) antes de mexer no código.

## Nota sobre o radar denso

Com 42 pilares (vs 13 antes), o radar ganhou muitos vértices (G tem 23). Independente da cor,
vale avaliar **agrupar visualmente os pilares por letra** (ou mostrar só as 4 letras no nível
macro, com drill-down por pilar) pra não poluir. Decidir junto com a cor.
