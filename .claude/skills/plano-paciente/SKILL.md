---
name: plano-paciente
description: Monta a devolutiva de resultados de um paciente (o "deck") a partir do prontuário do EMR. O EMR MONTA o rascunho mecânico só com código, de graça (arco, réguas, histórico, plano, receita); esta skill conduz a DISCUSSÃO clínica que nenhum código faz sozinho, aplica o resultado ao rascunho e salva de volta no EMR. Invocar quando o usuário pedir "plano do paciente", "devolutiva", "deck do <nome>", "montar o plano de resultados", "elaborar o plano". NÃO usar para deck comercial Plenya (isso é /plenya-deck) nem para aula (isso é /aula).
---

# Skill `/plano` — devolutiva de resultados do paciente

> Gramática e regras de desenho: [referencias/gramatica-v2.md](referencias/gramatica-v2.md).
> Documento vivo: [docs/emr/plano-ui-geracao-plano.md](../../../docs/emr/plano-ui-geracao-plano.md).

## O que mudou, e por que esta skill existe agora

**O EMR MONTA o rascunho sozinho, só com código.** `POST /patients/:id/plans/assemble` devolve 18
slides em milissegundos, **sem chamar modelo nenhum e sem custo**: o arco, as réguas hidratadas do
escore com a glosa do catálogo, a mini-série do histórico ("8,60 em 2024, 6,97 em 2025, 9,84
agora"), os exames que faltam, uma conduta registrada por slide, a sequência e a receita como
"para levar". Tudo isso é montagem, e montagem não se paga por token.

**O que o código deixa em branco de propósito:**

| campo | por quê |
|---|---|
| `punch` | é a consequência clínica ("a ferritina dobrou, e isso muda o que se investiga") |
| título como afirmação | o código escreve o descritivo ("Ferritina, saturação e ferro sérico") |
| tema da seção | "O que o cigarro está cobrando" é leitura, não agrupamento por ranking |
| a narrativa clínica | o que uma investigação já respondeu; mora na `ClinicalNote` |

Então o fluxo é: **o médico preenche → o EMR monta de graça → nós escrevemos a leitura → volta.**

Existe também `POST .../plans/generate`, que faz o mesmo deck com o texto redigido pelo modelo, em
~70 segundos por ~US$ 0,21. Use quando o médico quiser um rascunho já escrito e não houver conversa
para ter. O caminho normal desta skill é **montar e discutir**: sai melhor e custa zero.

---

## Fase 0 — buscar o que já existe

```
GET /api/v1/patients/{id}/plan-dossier          # o prontuário compilado
GET /api/v1/patients/{id}/plans                 # planos que já existem
```

Diga em uma linha o que veio: quantas réguas, quantos achados de cada lado, quantas condutas,
quantas receitas, se há vitais. **Esse resumo já diz o tamanho do deck que vai sair**, e é honesto
dizer isso antes de gastar a geração:

| falta no dossiê | o deck perde |
|---|---|
| `carePlan` vazio | a seção do plano inteira, 5 a 6 slides nos decks aprovados |
| `vitals` vazio | pressão e peso; nenhum deck aprovado cita o que não foi aferido |
| `prescriptions` vazio | o "para levar" sai sem dose |
| `labRequest` ausente | o slide dos exames que faltam |

Dossiê sem exame e sem anamnese: **pare**. Não há devolutiva a montar, e a geração recusa com 422.

---

## Fase 1 — montar o rascunho pelo EMR

```
POST /api/v1/patients/{id}/plans/assemble       # código puro, custo zero, corpo: {"title"?}
POST /api/v1/patients/{id}/plans/generate       # com modelo, ~US$ 0,21, {"instruction"?}
```

Monte primeiro. Só passe ao `generate` se o médico pedir o texto pronto sem conversa; `instruction`
lá dirige o recorte ("foque no ferro e no sono"), nunca dita texto.

A resposta traz `plan` (o rascunho já salvo), `reply`, `warnings` e `usage` (zerado no montado).
**Leia os `warnings` antes de olhar os slides** — eles são de quatro naturezas e a ação é diferente:

| `kind` | o que é | o que fazer |
|---|---|---|
| `numeral` | número escrito que não existe no dossiê | conferir a frase; é o aviso mais sério |
| `regua` | exame que não está no prontuário | a régua foi removida; ver se faz falta |
| `estilo` | fora do padrão medido (punch, contagem de régua, travessão) | ajustar na fase 3 |
| `lacuna` | o prontuário não tinha o dado | **não é defeito do deck**: é registro que falta |

No deck montado dois avisos de lacuna são esperados e não são defeito: o número de slides sem punch
(é o trabalho que sobrou para esta conversa) e o de slides do plano que trazem o `rationale` como
está no prontuário. **Esse segundo merece leitura**: o texto foi escrito para o prontuário e às
vezes traz diretriz e sigla que o paciente não lê ("a ESE/ENSAT 2023 pede confirmação da
independência do ACTH"). Reescrever essas frases é parte da fase 2.

---

## Fase 2 — a discussão, que é o motivo de tudo

Aqui está o valor. Leia o rascunho e traga ao médico **o que o gerador não pôde saber**, em uma
lista curta. As perguntas que produziram os slides bons dos decks aprovados:

1. **Por que este achado está assim?** O dossiê diz que a ferritina dobrou; só o médico diz se é
   inflamação, reposição ou sobrecarga. Isso vira o `punch`, que o montador deixa vazio em todos os
   slides de conteúdo: é a maior parte do que esta conversa acrescenta.
2. **O que uma investigação já respondeu?** "O que a tomografia de maio já respondeu" foi um slide
   inteiro no deck da Ana, e não existe em campo nenhum.
3. **Há dois caminhos de verdade?** Só então existe o slide de decisão, com o cartão de veredicto.
   Não invente dilema: nos dois decks aprovados ele aparece uma vez.
4. **Estes exames contam a mesma história?** É o que transforma "O que está se movendo · 1 de 3" em
   "O que o cigarro está cobrando". O gerador agrupa por ranking; o tema é leitura.
5. **O que o paciente já tentou?** Muda o tom da conduta, e não está no prontuário.

**Nunca preencha com plausibilidade.** Se o médico não respondeu, o slide não existe. Inventar
leitura clínica é pior que um deck curto, porque o paciente lê como se fosse dele.

Quando a resposta do médico for uma conduta nova, **peça para ele registrar no plano de cuidado**
em vez de escrever direto no slide: o deck é derivado, e conduta que só existe no deck some do
prontuário.

---

## Fase 3 — aplicar ao rascunho

**O caminho curto, e o normal desta skill:** `scripts/plano/plano.py`, que lê o deck numerado e
escreve só os campos de texto de volta, com `expectedRevision` e medindo o estouro depois.

```bash
scripts/plano/plano.py ler      --paciente <uuid> --plano <uuid>   # deck numerado, marca "sem punch"
scripts/plano/plano.py escrever --paciente <uuid> --plano <uuid> --edicoes leitura.json
```

`leitura.json` é `{"<slide>": {"campo": "valor"}}` — a chave é a posição ("7") ou o id do slide, e o
campo é um de `title`, `punch`, `lede`, `kicker`, `eyebrow`. Régua, tabela, cartão e dose ele recusa
de propósito: **dado se corrige no prontuário**, e o rascunho se remonta.

Em dev use `--api http://localhost:3001` (bypass ligado, qualquer token). Contra o espelho de prod,
rode de dentro do container com `--api http://127.0.0.1:3005`.

As outras duas vias, quando fazem sentido:

**A conversa do assistente**, quando quem edita é o médico na tela (e não esta conversa):
```
POST /api/v1/patients/{id}/plans/{planId}/assistant/messages
     {"body": "...", "clientMessageId": "<uuid fixo por mensagem>"}
```
O servidor decide o que entra: texto aplica direto e reversível; número, unidade, dose ou régua
vira **sugestão** com a origem do número ao lado, para o médico aceitar slide a slide. Tudo fica no
histórico de revisões.

**O PUT**, quando você está reescrevendo estrutura (acrescentar um slide de narrativa, refazer uma
seção):
```
PUT /api/v1/patients/{id}/plans/{planId}
    {"title": "...", "content": [...], "expectedRevision": <revisionSeq atual>}
```
`expectedRevision` não é opcional: sem ele você sobrescreve em silêncio quem escreveu antes.

Depois de qualquer alteração, **sempre**:
```
GET /api/v1/patients/{id}/plans/{planId}/overflow
```
Lista vazia = cabe. Lista com slides = corte texto e rode de novo. O slide tem altura fixa e
`overflow:hidden`: o que não cabe **some do PDF sem erro nenhum**.

---

## Fase 4 — entregar ao médico, não ao paciente

Apresente: o arco, o que a discussão acrescentou, os avisos que sobraram e o resultado da
conferência. **Publicar é decisão dele**, e é o que entrega ao paciente:

```
POST .../plans/{planId}/publish    # PDF 16:9 + A4 no portal
POST .../plans/{planId}/report     # relatório A4 assinado com ICP-Brasil (ato médico)
```

Os três saem do MESMO conteúdo. Publicação com slide estourando devolve 422 e lista quais.

---

## O alvo, medido nos dois decks aprovados

Não é opinião: foi contado slide a slide em Ana (21) e José Ricardo (20).

- **A tabela é o bloco mais usado**, não a régua: 9 de 21 e 8 de 20 slides. A régua é o átomo
  visual; a tabela é como o plano se explica.
- **Régua: 2 a 4 por slide**, média 3,1. Nunca uma sozinha, nunca cinco.
- **`punch` em 85%** dos slides, ausente só em capa, para-levar e fecho. De 55 a 110 caracteres,
  **exatamente um `<em>`**, e em 9 de 10 a frase termina dentro dele. Duas frases: constatação
  plana, depois a virada que carrega a decisão.
- **Título de 16 a 53 caracteres**, uma linha, e é uma AFIRMAÇÃO: "A ferritina dobrou em dois anos",
  não "Ferritina". Só o fecho é longo, três frases.
- **A legenda da rampa aparece uma vez por deck**, no primeiro slide de régua.
- **`variant: deep` em exatamente dois slides**: capa e fecho. Um tom só nas páginas de conteúdo.
- **A sequência é uma tabela de 2 colunas**, com a primeira em `dose` e valores relativos ("Agora",
  "Em 4 semanas"). O kind `sequence` nunca foi usado em nenhum dos dois.
- **Boas notícias primeiro**, nunca em menos de 2 slides, e sempre antes de qualquer notícia ruim.

O servidor já garante os quatro últimos automaticamente. Os outros são escrita.

---

## Regras editoriais invariantes

Valem em tudo que o paciente lê: **sem travessão** (vírgula, ponto ou dois-pontos); sem
"Não é X. É Y."; sem fecho-slogan; sem ícone decorativo em lista; **sem preço**; **sem marca
comercial** (suplemento, aparelho, laboratório, varejista) — use a categoria; sem "medicina
preditiva"; nada que identifique terceiros.

**Regra de lei da régua:** nenhuma régua entra num slide sem um rótulo avaliativo visível no MESMO
slide — no título, no `punch` ou no `note`. Barra colorida sem rótulo comunica pior que barra com
rótulo.

---

## Erros que já aconteceram, não repetir

1. **Testar com paciente fictício.** Um deck só se julga contra o prontuário real de quem ele
   descreve. Paciente de teste com zero conduta produz um deck sem metade do plano e leva a
   conclusões erradas sobre o gerador.
2. **Escrever antes de ler o dossiê.** Todo número vem de lá.
3. **Usar o nome do catálogo na régua.** "Ferritina - Mulheres Pós-Menopausa" é do escore;
   `display` é "Ferritina", que é o que a paciente reconhece. O montador já resolve isso sozinho
   (`lab_test_definitions.name` + `patient_gloss`); não reescreva à mão.
4. **Pagar o modelo para fazer montagem.** Agrupar exame, hidratar régua, formatar histórico,
   copiar conduta e dose: tudo isso o `assemble` faz de graça. Gerar com IA para depois reescrever
   metade é gastar duas vezes.
5. **Encher o slide.** Oito réguas não cabem, quatro cabem. Conferir sempre.
6. **Transformar checklist de ausência em "o que está bem".** "Adrenalectomia: não" não é conquista.
7. **Afinar o eixo da régua à mão.** O eixo é a escala do escore e o servidor o calcula; o valor
   fora da faixa encosta na borda de propósito, com o número impresso ao lado. Isso é a informação
   de estar fora da escala, não um defeito a corrigir.
8. **Publicar sem o médico mandar.**

---

## Dados clínicos

Ficam em `pacs/<NOME>/`, que é gitignored. Nada de conteúdo de paciente em commit, em log ou em
mensagem de erro.
