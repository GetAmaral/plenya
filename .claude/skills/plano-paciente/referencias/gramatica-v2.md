# Gramática v2 da devolutiva — os 8 blocos

Referência do skill `/plano`. O contrato é `pdfdoc.DeckSlide`
(`apps/api/internal/pdfdoc/deck_blocks.go`); o que está aqui é o mesmo, do lado de quem escreve.

Esta gramática não foi inventada: saiu da leitura dos decks que já existem. Ana e José Ricardo
convergiram para a MESMA estrutura, com contagem de bloco quase idêntica (régua 230 vs 231, resumo
17 vs 17, punch 18 vs 17). O deck do João é de uma geração anterior, por sistema em vez de por
narrativa, e não é o modelo.

---

## Envelope comum

Todo slide aceita:

```json
{
  "kind": "rulers",
  "variant": "",
  "eyebrow": "O que está se movendo · 1 de 3",
  "title": "A ferritina dobrou em dois anos",
  "lede": "frase de abertura, opcional",
  "kicker": "parágrafo de apoio, opcional",
  "source": "nota de rodapé, opcional",
  "punch": "a frase que fecha, com <em>uma ênfase</em>"
}
```

`variant`: `""` (creme, padrão), `"dark"` (petróleo) ou `"deep"` (petróleo escuro). Use escuro
só na capa e no fecho.

Inline permitido em qualquer texto: `<em>` `<strong>` `<b>` `<i>` `<small>` `<br>`. Qualquer outra
tag vira texto literal. O `<em>` do `punch` é o que sai dourado.

---

## 1 · `cover` — a capa

```json
{ "kind": "cover", "variant": "deep",
  "eyebrow": "Seus exames · 27 de agosto de 2026",
  "title": "José Ricardo",
  "lede": "Três anos de exames lidos juntos. A boa notícia primeiro, depois as três coisas que estão se movendo." }
```

`title` é o primeiro nome. No relatório A4 assinado ele é omitido de propósito (a papelaria já
imprime o nome no bloco de identificação); o `eyebrow` e o `lede` continuam.

## 2 · `summary` — onde você está, em uma página

Dois cartões e os passos. É o slide que o paciente mais relê.

```json
{ "kind": "summary", "eyebrow": "Resumo", "title": "Onde você está, em uma página",
  "summary": {
    "cards": [
      { "title": "O que está forte", "tone": "bom", "lines": [
        { "name": "Rim", "sub": "sem perda de proteína", "value": "2,7", "unit": "mg/g",
          "ruler": { "...": "cole a régua do dossiê para sair a mini-barra" } }
      ]},
      { "title": "O que está se movendo", "tone": "ruim", "lines": [
        { "name": "Ferritina", "sub": "dobrou em dois anos", "value": "500", "unit": "ng/mL" }
      ]}
    ],
    "stepsTitle": "O que vamos fazer",
    "steps": ["Tirzepatida uma vez por semana", "Treino de força com carga"]
  },
  "punch": "Nenhum dos quatro é doença ainda. <em>É cedo o bastante para mudar a direção.</em>" }
```

`tone`: `"bom"` ou `"ruim"` (a barra colorida na lateral do cartão). Quatro linhas por cartão é o
teto confortável; quatro passos também.

## 3 · `rulers` — as réguas

O átomo da devolutiva. **1 a 4 por slide.**

```json
{ "kind": "rulers", "title": "A ferritina dobrou em dois anos", "legend": true,
  "rulers": [
    { "code": "PLNCEFB97FD", "display": "Ferritina", "sub": "estoque de ferro", "unit": "ng/mL",
      "axis": [0, 520],
      "segments": [ { "level": 0, "a": 0, "b": 15 }, { "level": 5, "a": 50, "b": 200 } ],
      "history": [ { "value": 432, "text": "432" }, { "value": 500, "text": "500" } ],
      "note": "239 em 2024, 432 em 2025, 500 agora" } ] }
```

- **Copie `axis`, `segments` e `history` do dossiê.** Não recalcule.
- `display` e `sub` são SEUS: o nome que o paciente reconhece e o que o exame mede.
- Só os **dois últimos** pontos do histórico aparecem, com a seta do anterior para o atual. Mande a
  série inteira; a régua escolhe.
- **`axis` é o único número que às vezes se ajusta à mão**: quando um valor extremo esmaga a escala
  (PCR de 63 num eixo que termina em 15), aperte o teto e deixe o ponto encostar na borda. Foi o que
  a Ana precisou. Ajuste o eixo, nunca o valor.
- `legend: true` no primeiro slide de régua de cada bloco, não em todos.

## 4 · `two-cards` — a decisão em aberto

Dois caminhos, um deles descartado.

```json
{ "kind": "two-cards", "title": "Os dois caminhos, e qual é o seu",
  "cards": [
    { "kicker": "Caminho 1 · sobrecarga de origem genética", "dim": true,
      "body": "A marca é a saturação de transferrina alta, acima de 45%." },
    { "kicker": "Caminho 2 · causa metabólica", "focus": true,
      "body": "A sua saturação é 30%, e caiu de 36% no ano passado." } ] }
```

`dim` apaga o caminho descartado; `focus` destaca o que vale. Use os dois no mesmo slide — é o
contraste que ensina.

## 5 · `plan-step` — uma conduta

```json
{ "kind": "plan-step", "eyebrow": "O plano · 1 de 5", "title": "A tirzepatida: como começa",
  "lede": "Uma aplicação por semana, sempre no mesmo dia.",
  "cards": [ { "kicker": "Semanas 1 a 4", "body": "1,25 mg por semana." } ] }
```

Uma conduta por slide. Aceita também um bloco `takeaway` quando a conduta tem doses.

## 6 · `sequence` — os próximos meses

```json
{ "kind": "sequence", "title": "Os próximos três meses, em ordem",
  "steps": [
    { "when": "Esta semana", "what": "Começa a tirzepatida", "detail": "Junto com creatina e ômega-3." },
    { "when": "Em 4 semanas", "what": "Consulta de ajuste" } ] }
```

`when` é relativo ("em 4 semanas"), não data absoluta: o paciente lê isso semanas depois.

## 7 · `takeaway` — o que começa a tomar agora

```json
{ "kind": "takeaway", "eyebrow": "Para levar", "title": "O que você começa a tomar agora",
  "takeaway": {
    "highlight": { "when": "Uma vez por semana", "name": "Tirzepatida",
                   "obs": "1,25 mg nas quatro primeiras semanas",
                   "dose": "1,25", "unit": "mg por semana" },
    "groups": [
      { "title": "De manhã", "items": [ { "name": "Creatina", "dose": "5 g" } ] },
      { "title": "Todo dia", "items": [ { "name": "Água", "sub": "mais nos dias de treino", "dose": "3,0 a 3,5 L" } ] } ],
    "note": "A fórmula manipulada ainda não entra: ela é montada na consulta de quatro semanas." } }
```

`highlight` é o que muda o tratamento, um só. Três `groups` cabem lado a lado.

## 8 · `closing` — em uma página

```json
{ "kind": "closing", "variant": "deep", "eyebrow": "Em uma página",
  "title": "Você está bem, e três marcadores estão se movendo na mesma direção.",
  "lede": "Começamos pela tirzepatida, junto com proteína e treino de carga." }
```

O `title` aqui é longo de propósito: é o resumo que o paciente leva na cabeça.

---

## As quatro regras de desenho, e por que existem

1. **A régua mostra a faixa-meta NO LUGAR da faixa de referência do laboratório**, não somada a
   ela. Substituir mede melhor que somar.
2. **A direção da mudança tem que estar visível** (a seta do valor anterior). Direção é o que o
   leitor mais erra e o que mais importa clinicamente.
3. **Toda régua precisa de um rótulo avaliativo no mesmo slide.** Barra colorida com rótulo bate
   barra colorida sozinha. O rótulo pode vir do título, do `punch` ou do `note`.
4. **Um assunto por slide.** Informação em blocos pequenos, verificada a cada bloco, é o que
   sustenta a compreensão.

Fontes: [revisão sistemática JMIR 2024](https://www.jmir.org/2024/1/e53993) ·
[faixa-meta substituindo referência](https://www.ncbi.nlm.nih.gov/pmc/articles/PMC6231727/) ·
[harm anchors](https://www.ncbi.nlm.nih.gov/pmc/articles/PMC5891666/) ·
[direção da mudança](https://www.advancesinpro.org/article/S3050-6964\(26\)00018-2/fulltext) ·
[AHRQ teach-back](https://www.ahrq.gov/health-literacy/improve/precautions/tool5.html).
