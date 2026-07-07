# Continuum · deck-v5 — Plano de reestruturação

> Origem: briefing de reunião (2026-07-06). v5 = cópia da v4, reestruturada.
> v4 fica **congelada** como backup. Editar só `deck-v5.html`.

## Decisão base
Deck encolhe de **21 → 13 slides**. Corta o Ato I (hook/janela/origem), funde
crenças/paradigma/diferença na fala do médico-gestor, tira o Escore do deck
(migrou pro fim da consulta), reordena o miolo e reforça a ancoragem de valor
antes do preço (técnica Pedro Quintanilha).

## Nova sequência (13 slides)

| # v5 | Slide | Vem do v4 | O que muda |
|---|---|---|---|
| 1 | **Capa** | s04 (reveal «Continuum.») | repropor como capa; cerimonial, sem page mark |
| 2 | **Para quem não é** | s16 | ordem invertida (não-é antes de é); **2 slides separados** |
| 3 | **Para quem é** | s15 | — |
| 4 | **Médico Gestor** | s08 | fala absorve 5/6/7; citar de leve (não poluir) |
| 5 | **Corpo como um só sistema** | s10 | tirar moldura AGIR (microcounter A·G·I·R) — foco gestor |
| 6 | **Método AGIR** | s09 | — |
| 7 | **Jornada do Programa** | s12 | — |
| 8 | **Equipe Plenya** | s13 | — |
| 9 | **Box Plenya** | s17 | — |
| 10 | **Resumo das entregas** | s18 | repetição proposital da entrega |
| 11 | **Ancoragem de valor** | s19 (reformulado) | «quanto vale seu sonho / futuro / longevidade / ver o neto crescer? 100k? 200k? 300k?» — desenho aberto, aspiracional-com-gravidade (não fear) |
| 12 | **Investimento** | s20 | destacar a **parcela** (grande); à vista **menor com benefício** |
| 13 | **Fechamento cerimonial** | s21 | mantém «Viva bem, viva mais.»; sem page mark |

## Cortados do deck
s01 (hook 10 anos), s02 (janela silenciosa), s03 (vinte anos/origem),
s05 (paradigma 3 verbos), s06 (três crenças), s07 (a diferença/estrada),
s11 (Escore Plenya → fim da consulta), s14 (o que 20 anos ensinaram — redundante).

## Sequência do golpe de ancoragem (Pedro Quintanilha)
Slide 11 mostra valores altos na tela (100k / 200k / 300k — «quanto você pagaria
pra alcançar seu futuro?»). Slide 12 revela a **parcela** (~ordem de R$ 3-5k/mês)
em destaque = quebra de paradigma. Executar em voz Plenya (quiet luxury, editorial),
não em tom Hormozi/game-show.

## Regras invariantes que continuam valendo
Sem preços fora do slide 12, sem marcas comerciais, sem casos clínicos
identificáveis, sem em-dash, sem «não é X. é Y.» empilhado. Aspiracional, não
fear-based (a ancoragem é sobre o que se preserva/conquista, não sobre susto).
Fonte de copy = canonical do site. Ver EDITORIAL.md + memórias editoriais.

## Fases de execução — ✅ CONCLUÍDA (2026-07-07)
1. **Mecânica (script):** reordenar/renumerar ids + page marks (denominador → 13). ✅
2. **Conteúdo (slide a slide, PNG aprovado):**
   - s04 Médico Gestor: linha-cue «Normal não é ótimo · Enxergar antes · Integrar o todo». ✅
   - s05 Corpo: microcounter A·G·I·R removido. ✅
   - s11 Ancoragem: refeito 3×. Final = imagem estrada-ao-amanhecer (Slide7-novo-b, universal, sem neto) + frase «Quanto vale viver bem por mais tempo?» + escada 100/200/300 mil em crescendo (tamanho+brilho) + eyebrow «O valor do tempo». ✅
   - s12 Investimento: parcela em destaque (12× R$ 5.417 / 6× R$ 6.167), à vista discreto sem desconto, duas modalidades. ✅
3. **PDF gerado:** `docs/decks/continuum-plenya-20260707-v5.pdf` (13 slides). ✅

## Pós-rodada (pendente de ordem)
- Deploy do PDF/previews pro VPS (`decks.plenyasaude.com.br`) — só sob ordem.
- Commit da v5 (deck-v5.html + V5-PLANO.md) — só sob ordem.

## Pendências abertas
- Preço/parcela exata do slide 12 (v4 tem 12× R$3.750; briefing citou «~5k»
  ilustrativo — confirmar com o user ao chegar no slide).
- Imagem nova possível só na Ancoragem (prompt gpt-image-2 a definir).
