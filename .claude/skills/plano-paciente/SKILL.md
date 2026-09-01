---
name: plano-paciente
description: Monta a devolutiva de resultados de um paciente (o "deck") dentro do EMR. Puxa o dossiê derivado do prontuário (réguas por exame, achados de anamnese, sinais vitais, achados ordenados por peso), propõe o arco narrativo, escreve os títulos e as frases de fechamento em voz de paciente, e grava o plano como rascunho para o médico revisar e publicar. Invocar quando o usuário pedir "plano do paciente", "devolutiva", "deck do <nome>", "montar o plano de resultados". NÃO usar para deck comercial Plenya (isso é /plenya-deck) nem para aula (isso é /aula).
---

# Skill `/plano` — devolutiva de resultados do paciente

> Referência interna usada por `.claude/commands/plano.md`.
> Documento vivo do projeto: [docs/emr/plano-planos-paciente.md](../../../docs/emr/plano-planos-paciente.md).
> Gramática e regras de desenho: [referencias/gramatica-v2.md](referencias/gramatica-v2.md).

O EMR já deriva os FATOS. Esta skill escreve o que é JULGAMENTO: qual é o arco, o que vira
slide, e como cada coisa é dita para o paciente. A divisão não é negociável — inventar número é
o pior erro possível aqui.

---

## Modo de operação

Analise `$ARGUMENTS`:

- nome ou id de paciente, sem mais nada → **MODO CRIAR**
- "revisar" / "ajustar slide N" → **MODO REVISAR**
- "publicar" → **MODO PUBLICAR** (só depois de o médico aprovar)

---

## Fase 0 — GATE: o dossiê antes de qualquer frase

**Regra de lei: é proibido escrever um slide antes de ler o dossiê.** Todo número, nome de exame,
data e valor sai de lá. Se um dado que você quer citar não está no dossiê, ele não entra no plano —
ou você vai ao banco confirmar, ou o slide muda.

```
GET /api/v1/patients/{patientId}/plan-dossier
```

O dossiê traz, já derivado das TRÊS fontes do prontuário (exames, anamnese e consulta), inclusive
o que ainda está em rascunho:

| Campo | O que é |
|---|---|
| `rulers` | uma régua por exame: escala do escore aplicável a ESTE paciente + histórico dele |
| `strong` | o que está bem, **ordenado pelo peso do item** |
| `moving` | o que está se movendo, **ordenado por pontos perdidos** |
| `vitals` | pressão, frequência, peso, cintura, IMC das duas últimas consultas |
| `carePlan` | condutas já registradas por pilar AGIR |
| `labRequest` | o último pedido de exames (para o slide "os exames que faltam") |
| `prescriptions` | receitas vigentes (para o slide "para levar") |
| `snapshot` | o escore vigente, quando existe |

Confirme em uma linha, antes de seguir: quantas réguas, quantos achados de cada lado, se há vitais
e se há escore. Dossiê vazio (paciente sem exame e sem anamnese) → **pare** e diga isso; não há
devolutiva a montar.

### O que o dossiê NÃO resolve, e você tem que buscar

- **Por que** um achado está como está. Isso é leitura clínica, não dado.
- Conduta que ainda não está em `carePlan`.
- Contexto de vida do paciente (trabalho, rotina, o que ele já tentou).

Quando faltar, **pergunte ao médico**. Não preencha com plausibilidade.

---

## Fase 1 — o arco, antes das palavras

Proponha o arco em uma lista curta e **espere aprovação** antes de escrever slide. Os dois últimos
decks (Ana, 21 slides; José Ricardo, 20) convergiram para a mesma ordem, e ela funciona:

```
1   capa                      nome, data, uma frase de abertura
2   resumo em uma página      o que está forte | o que está se movendo | o que vamos fazer
3-4 o que está bem            1 a 4 réguas por slide
5-7 o que está se movendo     um achado por slide
8-9 a decisão em aberto       quando há dois caminhos possíveis
10+ o plano                   uma conduta por slide
n-2 a sequência               os próximos três meses, em ordem
n-1 para levar                o que começa a tomar agora, com dose
n   em uma página             o fecho
```

Como escolher o que entra:

- **"O que está bem" vem do topo de `strong`** — que já está ordenado pelo peso do item, não por
  pontos perdidos. Em nível 4-5 ninguém perde ponto, então peso é o único sinal. Boa parte da
  anamnese é checklist de ausência ("Adrenalectomia: não") e **não vira slide**: use os marcadores
  pesados que estão no ótimo.
- **"O que está se movendo" vem do topo de `moving`.** Três é o número que funcionou nos dois
  decks. Um achado com `trend: worsening` merece prioridade mesmo em nível bom: a direção é o sinal.
- **A decisão em aberto** só existe quando há de fato dois caminhos (a ferritina do Ricardo). Não
  invente dilema.

Um assunto por slide. Se dois achados precisam ser explicados juntos, são dois slides.

---

## Fase 2 — escrever

Só agora. Leia [referencias/gramatica-v2.md](referencias/gramatica-v2.md) para os 8 blocos e o
JSON de cada um.

### Voz

Prosa clínica conectiva em PT-BR, dirigida ao paciente, sem infantilizar. O paciente é adulto e
está lendo sobre o próprio corpo.

- **Título de slide é uma AFIRMAÇÃO, não um rótulo.** "A ferritina dobrou em dois anos", não
  "Ferritina". "O rim e o fígado estão impecáveis", não "Função renal e hepática".
- **`display` da régua é o nome que o paciente reconhece**, não o do catálogo: "Ferritina", não
  "Ferritina - Homens". O `sub` explica o que o exame mede, em cinco palavras: "estoque de ferro".
- **O `punch` fecha o slide com a consequência**, e é o único lugar onde `<em>` entra (ele sai
  dourado). "Subir a ferritina e <em>não</em> subir a saturação é a chave."
- Número sempre com a unidade e a data de quando foi medido.

### Regras editoriais invariantes (valem em tudo que o paciente lê)

- **Sem travessão.** Vírgula, ponto ou dois-pontos.
- **Sem "Não é X. É Y."** e sem fecho-slogan.
- **Sem ícone decorativo** em lista.
- **Sem preço, sem marca comercial** (suplemento, wearable, varejista): use a categoria.
- **Sem "medicina preditiva".**
- Nada que identifique terceiros.

### A regra que a evidência impõe

**Nenhuma régua entra num slide sem um rótulo avaliativo visível no mesmo slide.** Barra colorida
sozinha tem desempenho pior do que barra colorida com rótulo. O rótulo pode estar no título do
slide, no `punch` ou no `note` da régua, mas tem que estar em algum deles.

---

## Fase 3 — gravar como rascunho

```
POST /api/v1/patients/{patientId}/plans      { "title": "...", "content": [ ...slides... ] }
```

Depois, **sempre**:

```
GET /api/v1/patients/{patientId}/plans/{planId}/overflow
```

Lista vazia = cabe. Lista com slides = **corte texto e rode de novo**. O slide tem altura fixa e
`overflow:hidden`: o que não cabe some do PDF sem erro nenhum, e a publicação recusa.

Não publique. Apresente ao médico: o arco, os slides e o resultado da conferência. Publicar é
decisão dele.

---

## MODO REVISAR

Carregue o plano (`GET .../plans/{planId}`), aplique o ajuste pedido, salve com
`PUT .../plans/{planId}` e **rode a conferência de novo** — texto novo é a causa mais comum de
estouro. Salvar devolve o plano para rascunho de propósito: o que está no portal continua sendo a
versão publicada até alguém publicar outra.

## MODO PUBLICAR

Só com ordem explícita do médico.

```
POST .../plans/{planId}/publish     → PDF 16:9 + PDF A4 no portal (sem assinatura)
POST .../plans/{planId}/report      → relatório A4 assinado com ICP-Brasil (ato médico)
```

Os três saem do MESMO conteúdo. Se a publicação devolver 422, ela lista quais slides não cabem.

---

## Erros que já aconteceram, não repetir

1. **Escrever antes de ler o dossiê.** Todo número vem de lá.
2. **Usar o nome do catálogo na régua.** "Ferritina - Mulheres Pós-Menopausa" no slide do paciente.
3. **Encher o slide.** Oito réguas não cabem; quatro cabem. Conferir sempre.
4. **Transformar checklist de ausência em "o que está bem".** "Adrenalectomia: não" não é conquista.
5. **Publicar sem o médico mandar.**
