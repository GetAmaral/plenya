# Melhorias no form de anamnese e na aplicação de templates

**Data:** 2026-07-27 · **Estado:** em execução (dev)

Lista de 10 pontos levantados pelo Dr. Getúlio ao usar o form de anamnese, com a causa raiz
apurada de cada um e a correção. Três pontos foram decididos por ele em 2026-07-27 (marcados
com ✅ decisão).

---

## 1. Itens colados — falta separador

**Causa.** Na visão compacta (`AnamnesisTemplateItemsRenderer`), o `divide-y` só existe entre
*subgrupos*. As linhas de item dentro de um subgrupo têm `last:border-b-0` mas nunca ganharam
`border-b`, então nada separa um item do outro.

**Correção.** Linha fina (`border-b border-border/60`) em cada linha de item, exceto a última.

## 2. Span de dígitos — números confusos

**Causa.** `DigitSpanAdmin` renderiza as DUAS tentativas do mesmo comprimento lado a lado, sem
rótulo: sob "3 dígitos" aparece `3 8 6  6 1 2`, que lido corrido vira seis números.

✅ **Decisão.** É só isso — os comprimentos (direto 3-8, inverso 2-7) e as sequências ficam.
Rotular `1ª` / `2ª` e separar visualmente.

## 3. "Capacidade da memória percebida" tem que vir antes do Dubois

**Causa.** `anamnesis_template_items.order` foi gravado embaralhado em relação à árvore canônica
do escore. Ex. em `Médico | Inicial`, Cognição/Atual sai como 107 PHQ-9, 108 Dubois imediato,
109 Disposição familiares, 110 Capacidade da memória percebida — enquanto na árvore do escore a
ordem é Capacidade(1), Testes rápidos de memória(2) → filhos. É o mesmo sintoma do ponto 8
("no grupo original me parece certo").

**Correção.** Renormalizar `anamnesis_template_items.order` em TODOS os templates seguindo
`(score_groups.order, score_subgroups.order, score_items.order)`. A anamnese passa a herdar a
ordem do grupo original, e o ponto 3 se resolve sozinho.

## 4. Dubois — lembrar com dica pontua igual a lembrar sem dica

**Causa.** Bug real. `wordRecallScore` em `packages/domain/src/scales.ts` faz
`Object.values(answers).filter(v => v >= 1).length` — conta quantas palavras foram evocadas,
com ou sem dica. Ou seja, usa o *score total* antigo (/5 por fase), não o **Score Total
Ponderado**, que dobra o rappel libre e é o mais discriminante.

**Fonte primária.** Dubois B. et al., *"Les 5 mots"*, Presse Med 2002;31(36):1696-9 ·
Cowppli-Bony P. et al., Rev Neurol 2005;161(12 Pt 1):1205-12 · Croisile B. et al., Rev Neurol
2010;166(8-9):711-20 · ficha de administração INESSS/HAS. Cotação canônica:

- rappel libre (evocação espontânea, sem erro nem ajuda): **2 pontos** por palavra
- rappel indicé (evocação com a dica da categoria): **1 ponto** por palavra
- não evocada nem com dica: **0 pontos**
- score do rappel imediato = livre + indicé, **/10**
- score do rappel diferido = livre + indicé, **/10**
- **Score Total Ponderado = imediato + diferido, /20** — corte publicado: 19-20 declínio pouco
  provável; **≤18** possibilidade de declínio, investigar

✅ **Decisão.** Adotar o Score Total Ponderado. `wordRecallScore` passa a somar os pesos
(0-10 por fase), `maxScore` 5 → 10 nos dois itens, e as faixas do banco são recalibradas de /5
para /10.

Faixas novas — mesma quantidade de níveis das antigas, só reescaladas de /5 para /10:

| Item | antes (/5) | agora (/10) |
|---|---|---|
| Dubois imediato | N0 ≤3 · N3 =4 · N5 =5 | N0 ≤8 · N3 =9 · N5 =10 |
| Dubois tardio | N0 ≤2 · N2 =3 · N4 =4 · N5 =5 | N0 ≤7 · N2 =8 · N4 =9 · N5 =10 |

As respostas já gravadas (`scale_responses.answers`) sempre usaram os pesos 0/1/2, então a
migration recalcula `total`, `numeric_value` e `selected_level` sem perda.

## 5. ASEX e épocas de libido → Continuum Médico Complemento

**Estado.** `IIEF-5` e o pai `Escalas de desempenho:` já estão no
`Continuum | Médico | Complemento`; a ASEX e as duas "Épocas de melhor/pior libido/desempenho"
ficaram no `Continuum | Médico | Inicial`.

**Correção.** Mover os 3 itens do Inicial para o Complemento (o template completo
`Médico | Inicial`, que é o superset, mantém tudo).

## 6. Rever a escala de sono

**Causa.** Bug de classificação, não de conteúdo. O `evaluatesTrue` do renderer diverge do
motor Go (`models/score_level.go`) e do `matchLevel` de `@plenya/domain`:

| operador | renderer (errado) | canônico |
|---|---|---|
| `<` / `<=` | lê `lowerLimit` | lê `upperLimit` |
| `between` | `[lower, upper]` fechado | `(lower, upper]` meio-aberto |

Efeito no Epworth (níveis `>18` / `between 15-18` / `between 10-15` / `<=10`): o N5 "≤10", que
é o resultado normal e mais comum, **nunca é marcado** (lowerLimit é NULL), e um total de 10 cai
em N2 em vez de N5. Vale para todo item com `<`/`<=` — Dubois, Span, e os laboratoriais.

✅ **Decisão.** Corrigir o operador. Conteúdo do Epworth fica.

## 7. Peso/altura sem campo numérico e sem cálculo de IMC

**Causa (a).** Todos os 47 itens de `Composição corporal > Medidas Objetivas` estão com
`score_items.unit` NULL. O renderer só mostra o input numérico quando `scoreItem.unit` existe,
então peso, altura, quadril etc. aparecem sem nenhuma forma de entrada.

**Causa (b).** Não existe nenhum cálculo derivado no EMR — IMC e as demais razões teriam de ser
digitadas à mão.

**Correção.** (a) Popular `unit` nos 47 itens. (b) Motor de métricas derivadas
(`DERIVED_METRICS` em `@plenya/domain`, keyed por `anamnese_item_code`) que o renderer recalcula
a cada mudança e preenche automaticamente, com o nível auto-detectado, deixando a digitação
manual sobrepor:

| Derivado | Fórmula |
|---|---|
| IMC | peso / (altura/100)² |
| BRI | 364,2 − 365,5·√(1 − ((cintura/2π)² / (0,5·altura)²)) |
| Razão cintura/altura | cintura / altura |
| Razão cintura/quadril | cintura / quadril |
| Relação pescoço-altura | pescoço / (altura/100) |
| % gordura corporal | massa gorda total / peso · 100 |
| FMI | massa gorda total / (altura/100)² |
| MME / Peso | massa muscular esquelética / peso · 100 |
| Índice MME | massa muscular esquelética / (altura/100)² |
| ASMI | massa apendicular / (altura/100)² |
| Água corporal total (%) | ACT / peso · 100 |
| Razão AEC/ACT | água extracelular / ACT · 100 |

## 8. Cirurgias e medicamentos fora do pai

**Causa.** Dados, não template. `score_items.parent_item_id` está NULL em itens que deveriam
pendurar no pai; como `nestByParent` promove a raiz todo filho sem pai presente, eles saem soltos
no meio da lista:

- `Cirurgias já realizadas` → "Amputação de membro" e "Adrenalectomia" (os outros 18 já apontam
  para "Cirurgias que interferem diretamente no escore")
- `Medicamentos` → as 15 classes de "Analgésicos…" a "Antivirais de uso contínuo" + "Inibidores
  de bomba de prótons" (as outras 16 já apontam para "Uso atual de medicamentos")
- `Vida Sexual > Atual` → ASEX sem o pai "Escalas de desempenho:" (o IIEF-5 tem)

**Correção.** Setar `parent_item_id` nos 18 órfãos.

## 9. Vícios > Sexo com N5 default

**Causa.** `default_level5` é `true` em Tabaco, Álcool, Drogas ilícitas e Jogos/apostas, e
`false` em "Sexo" e "Outros vícios".

**Correção.** Ligar em "Sexo". ("Outros vícios" fica de fora — é campo aberto; confirmar com o
Getúlio se ele quer também.)

---

## Entregáveis

- `apps/web/components/anamnesis/AnamnesisTemplateItemsRenderer.tsx` — separador, operadores,
  derivados
- `apps/web/components/anamnesis/ScaleAdminWidgets.tsx` — tentativas do Span
- `packages/domain/src/scales.ts` — Score Total Ponderado
- `packages/domain/src/anthropometry.ts` — métricas derivadas
- `packages/domain/src/{anthropometry,scales}.test.ts` — 15 testes (vitest), verdes
- `apps/api/database/migrations/00063_anamnese_melhorias_form.sql` — units, órfãos, ordem dos
  templates, default_level5, faixas do Dubois, ASEX → Complemento

## Verificação (dev, 2026-07-27)

Migration 00063 aplicada (junto das 00060-00062, que estavam pendentes no dev). `tsc --noEmit`
do web limpo, 15 testes do domain verdes, e QA de runtime via Playwright no form real, sem
nenhum erro de console:

- peso 80 + altura 180 → **IMC 24,7 preenchido sozinho, com selo AUTO e N5 (18,5-25) marcado**
- Dubois imediato com 4 espontâneas + 1 com dica → **9/10 e N3** (antes marcava 5/5 e N5)
- Span de dígitos → `3 dígitos  1ª [3 8 6]  2ª [6 1 2]`, tentativas separadas
- Epworth 0/24 → **N5 ≤10 marcado** (antes nenhum nível marcava)
- Cognição/Atual abre em "Capacidade da memória percebida", depois "Testes rápidos de memória"
  com os 4 testes aninhados
- linha fina separando todos os itens

## Fora de escopo (decidido em 2026-07-27)

Levantados na revisão e **descartados pelo Getúlio** — não reabrir:

- **"Outros vícios"** fica sem `default_level5`. Só "Sexo" foi ligado.
- **"Qualidade percebida do sono"** fica com os níveis como estão (N2 "Acordo bem, mas a energia
  acaba no meio do dia" / N3 "Acordo cansado e depois pego no tranco").

## Deploy (2026-07-27) — ✅ EM PRODUÇÃO

Commit `9c18efca`. Api primeiro (roda a migration), depois web, um de cada vez.

- **api** — container `...171757298319`. Goose aplicou `00062` e `00063` (as 00060/00061 já
  estavam). Health 200.
- **web** — container `...174616369320`. Health 200. A **primeira tentativa falhou**: o
  `git clone` do GitHub caiu no meio (`fetch-pack: unexpected disconnect`) e deixou o container
  helper órfão segurando a fila, com o deploy parado em `in_progress` por 25 min sem criar
  container. Recuperado pelo procedimento de sempre (remover o helper órfão → limpar a fila →
  um deploy). O app antigo ficou no ar o tempo todo, sem indisponibilidade.

Conferido no banco de prod: 47/47 unidades, "Sexo" com `default_level5`, faixas do Dubois em
/10, 4 raízes remanescentes em Cirurgias/Medicamentos (são os 4 cabeçalhos, correto),
Cognição/Atual abrindo em "Capacidade da memória percebida" com os 4 testes aninhados, e ASEX
no Continuum | Médico | Complemento.

### O push também redeployou o site-getulio (comportamento correto)

O `site-getulio` (`qkdzqaauicc001qfkghfur0s`) rebuildou sozinho neste commit. Não é bug: os
`watch_paths` dele são `apps/site-getulio/**` + **`packages/**`** + `apps/site/content/blog/**`,
e nós mexemos em `packages/domain`. Subiu bem, no commit `9c18efca`, health 200.

O ponto de atenção é de **agenda, não de configuração**: esse deploy automático rodou ao mesmo
tempo que o deploy manual do api (ambos ~17:17), e três clones do monorepo concorrendo é a
hipótese mais provável para o clone do web ter estourado logo em seguida. Ao deployar depois de
um push que toca `packages/**`, esperar o auto-deploy do getulio terminar antes de disparar o
próximo app.

Nota: o `plenya-site` NÃO rebuildou (segue em `d1b6ad01`), ou seja, os `watch_paths` dele não
incluem `packages/**`. Se algum dia o site passar a consumir `@plenya/domain`, isso vira
divergência dev≡prod silenciosa.
