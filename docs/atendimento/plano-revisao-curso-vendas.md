# Plano de revisão — Atendimento Plenya à luz do curso "Meta Batida em 15 Dias"

> **Status:** EXECUTADO (2026-06-05). Guia (C1) e cérebro da IA (C2) editados; build Go verde (C3).
> Falta só o passo D.3: revisão final do Getúlio antes de qualquer deploy do bot (kill switch
> `RECEPTION_BOT_ENABLED` segue ativo). Opção 1 — adaptado e suave.
> **Decisão:** trazer a *engenharia da escuta* do curso, **não** o tom de condução comercial forte.
> **Âncora de voz:** brandbook oficial — a Plenya "orienta sem impor, educa sem simplificar demais,
> inspira sem exageros; não promete resultados rápidos ou milagrosos". Ver
> [`../branding/README.md`](../branding/README.md). Qualquer edição abaixo é subordinada a esse filtro.

## Contexto

Dois artefatos a revisar, derivados um do outro:
1. **Guia humano:** [`script-recepcao-conversao-leads.md`](script-recepcao-conversao-leads.md) (570 linhas).
2. **Cérebro da IA:** [`../../apps/api/internal/services/reception_brain.go`](../../apps/api/internal/services/reception_brain.go)
   (`receptionSystemPrompt`) — condensado do guia. **Ao mudar o guia, refletir aqui.**

Fonte do método: curso transcrito em `/home/user/curso-hubla/` (12 aulas + "O Livro Secreto das
Objeções"). PDF diagramado em `Meta Batida em 15 Dias — Curso Completo.pdf`. O curso é de infoproduto;
filtramos o que não serve a uma clínica regulada (CFM/LGPD).

## Princípio-mestre da revisão

O guia Plenya **já nasceu alinhado** (escutar antes de oferecer, espelhar, conduzir sem pressão, valor
depois da descoberta, firmeza no preço). Esta revisão **enxerta técnicas concretas que faltam** e
**registra o filtro** do que o curso ensina e nós rejeitamos. Não é reescrita.

---

## A. O que ADOTAR (adaptado a CFM/LGPD/voz) — com texto pronto

### A1. Descoberta em 4 tempos *(curso: Aula 3 — ordem das perguntas)*
**Problema atual:** §3.2/§4.3 fazem só 1–3 perguntas soltas; faltam *consequência* e *cenário ideal*,
que são o que constrói valor antes do preço.
**Adaptação:** sequência Contexto → Incômodo → Peso na vida (não-clínico) → Como seria melhor. Os dois
últimos tempos reformulados para **não virar anamnese** (recepção não entra em sintoma/diagnóstico).

Inserir no guia (substituir o miolo da §3.2) e espelhar na §4.3 (telefone):
```
A descoberta segue quatro tempos, em ritmo de conversa (nunca um questionário):
1. Contexto   — "o que te fez procurar um acompanhamento agora?"
2. Incômodo   — "e o que mais tem te incomodado nisso?"
3. Peso na vida (não clínico) — "isso tem pesado no seu dia a dia, na sua rotina?"
4. Cenário ideal — "como você imagina se sentir se isso entrasse nos eixos?"
```
> Guardrail a repetir: tempos 3 e 4 falam de **vida/rotina**, nunca de sintoma, resultado de exame ou
> diagnóstico. Isso é da consulta, com o médico (LGPD/CFM).

### A2. Espelhar-e-aprofundar *(curso: Aula 4)*
**Problema:** o Princípio 2 manda "espelhar o momento" mas não dá a mecânica.
**Adição:** nova subseção curta após §3.2 (e nota na §4.3): repetir a palavra/emoção do lead + devolver
uma pergunta.
```
Técnica de espelhamento: pegue a palavra ou a emoção que a pessoa usou e devolva com uma pergunta.
Ex.: lead diz "ando exausta o tempo todo" → "quando você diz exausta o tempo todo, como isso aparece
no seu dia?". Mostra escuta, reduz a defesa, e não promete nada.
```

### A3. Antídoto + recapitulação antes do preço *(curso: Aulas 8 e 9 — o ouro)*
**Problema:** hoje, ao perguntarem preço (cenário D / objeção 1), ancoramos genérico e soltamos o valor.
**Adaptação (sem promessa de resultado):** verbalizar a situação da pessoa com as palavras dela →
recapitular o que a consulta olha → confirmar ("faz sentido até aqui?") → **só então** o preço.
Inserir como novo bloco no topo da §6 (Fechamento):
```
Antes de dizer o valor a um lead aquecido, recapitule e confirme:
"Então, pelo que você me conta, o que pesa é [situação dela, nas palavras dela]. A consulta foi
pensada justamente para olhar isso com calma: cerca de 60 minutos, com um painel ampliado e uma
conduta para o seu caso. Faz sentido até aqui?"
Depois do "sim", o valor (ver A4). O 'faz sentido até aqui?' não é pedido de validação: é mais um
ponto de acordo antes do preço.
```

### A4. Disciplina do preço *(curso: Aula 8)*
**Problema:** a objeção 1 explica demais *depois* do preço (soa inseguro).
**Adaptação:** valor em **uma frase, tom neutro, sem emendar** formas de pagamento/álibis; deixe a
pessoa reagir. Reescrever a objeção 1 (WhatsApp + telefone) para encurtar o pós-preço e adicionar nota:
```
Regra do preço: diga o valor em uma frase calma e pare. Não emende parcelamento, justificativa ou
"mas inclui X, Y, Z" logo depois — isso lê como insegurança. O contexto vem ANTES do preço (A3),
não depois.
```

### A5. Dúvida real vs. desculpa *(curso: Aula 9 — "quem tem dúvida pergunta, quem tem medo foge")*
**Adição:** regra-mãe no topo da §5 (banco de objeções):
```
Antes de responder, classifique:
- Dúvida real (pergunta concreta: prazo, formato, como funciona, o que está incluso) → responda direto
  e com clareza.
- Evasiva ("vou pesquisar", "depois eu vejo", "preciso pensar" sem pergunta) → acolha e segure um
  próximo passo leve, sem insistir.
O tom muda: à dúvida, clareza; à evasiva, espaço com uma porta aberta.
```

### A6. Pedir permissão para avançar *(curso: Aula 7)*
**Adição:** micro-move na transição valor→agendamento (nota na §3.3/§3.5 e no cérebro):
```
Para avançar sem empurrar, peça passagem: "quer que eu te mostre como funciona na prática / que eu
veja uma data para você?". A pessoa concede o próximo passo; você não força.
```

### A7. Nova seção no guia — "Do método: o que tomamos e o que deixamos de fora"
Registrar explicitamente o filtro (conteúdo da seção B abaixo) como nova §12 do guia, para manutenção
futura e para a recepção entender o porquê.

---

## B. O que REJEITAR (o filtro — metade do valor da revisão)

O curso é de infoproduto; numa clínica regulada, isto **não entra**:
- **Agressividade comercial** ("sou mais agressiva", "quando vai ser o momento? quando fechar as
  portas?"). A marca recua o público quando pressiona.
- **"Perda" amedrontadora** do trio dor-ganho-perda → usar só a *janela silenciosa* com calma, nunca medo.
- **Promessa de transformação/resultado** ("vou resolver os problemas da sua vida") → CFM. O "ganho"
  só descreve a **qualidade do olhar/cuidado**, jamais um desfecho de saúde.
- **"Negocie o escopo para dar desconto"** → não transfere: preço único, sem escopo a reduzir. A
  firmeza gentil atual é a versão correta e superior.
- **Escada de micro-sims mecânica, bajulação, fechos-slogan, "Não é X, é Y"** → já vetados pela voz.

---

## C. Edições concretas por arquivo

### C1. `docs/atendimento/script-recepcao-conversao-leads.md`
1. §3.2 + §4.3 → descoberta em 4 tempos (A1).
2. Nova subseção §3.2.1 "Espelhar e aprofundar" + nota no telefone (A2).
3. §3.3/§3.5 + §6 → "pedir permissão para avançar" (A6).
4. §6 → novo bloco "Antídoto + recap antes do preço" (A3) no topo; e a "Regra do preço" (A4).
5. §5 (abertura) → regra "dúvida real vs. desculpa" (A5).
6. Objeção 1 → encurtar pós-preço (A4).
7. Nova §12 "Do método: o que tomamos e o que deixamos de fora" (A7/B).
8. Atualizar §10 (cartão de referência) se algum CTA/atalho mudar.
> Não tocar nos guardrails da §11 a não ser para reforçar (eles já cobrem o filtro B).

### C2. `apps/api/internal/services/reception_brain.go` (`receptionSystemPrompt`)
Adicionar/ajustar blocos no system prompt, refletindo o guia:
1. Bloco **DESCOBERTA EM 4 TEMPOS** (A1) com o guardrail não-clínico.
2. Bloco **ESPELHAR E APROFUNDAR** (A2).
3. Bloco **ANTES DE FALAR O PREÇO** (A3 + A4): recap → "faz sentido até aqui?" → preço em uma frase,
   sem emendar.
4. Bloco **DÚVIDA REAL vs. DESCULPA** (A5) orientando as actions (`answer` vs. acolher+`propose_schedule`).
5. Em **NUNCA FAÇA**, reforçar veto explícito a: agressividade, urgência forçada, "perda"
   amedrontadora, promessa de resultado (já há base; tornar explícito).
6. Não mudar o contrato de saída JSON (`reply/action/handoffReason/discloseAI`) nem a lógica de handoff.

### C3. Validação
- Recompilar Go no container: `docker compose exec -w /app api go build ./...`.
- Conferir que o guia e o prompt não divergem (mesmas frases-âncora).
- Sanidade de voz: reler trechos novos contra o filtro do brandbook (sem travessão, sem promessa, sem
  fecho-slogan).
- (Opcional) rodar 2–3 conversas de teste mentais: lead que pergunta preço a frio; lead "meus exames
  deram normais"; lead "vou pensar".

---

## D. Ordem de execução
1. Editar o guia (C1) — é a fonte; fechar a redação primeiro.
2. Refletir no cérebro da IA (C2) e recompilar (C3).
3. Reportar diffs para aprovação final do Getúlio antes de qualquer deploy do bot
   (o `RECEPTION_BOT_ENABLED` segue como kill switch; nada vai ao ar sem decisão).

## E. Não-objetivos
- Não alterar preço (R$ 800), política de Continuum, ou guardrails CFM/LGPD.
- Não mudar arquitetura do bot (modos copiloto/auto, slots, handoff).
- Não importar nada da seção B.

---

### Anexos de referência
- Curso: `/home/user/curso-hubla/curso-completo.txt` (+ PDF diagramado).
- Voz da marca: `docs/branding/README.md` · `docs/branding/brandbook-transcricao.md`.
- Memória: [[plenya_brand_essence]] · [[plenya_anti_ai_maneirismos]] · [[plenya_regras_editoriais]].
