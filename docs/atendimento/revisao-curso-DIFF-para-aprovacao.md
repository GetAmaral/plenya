# Revisão do atendimento à luz do curso — diff para aprovação

> Para o Getúlio aprovar antes de qualquer go-live do recepcionista de IA.
> Abordagem: **Opção 1 (suave)** — trouxemos a engenharia da escuta do curso, **não** o tom
> de condução comercial forte. Tudo subordinado à voz da marca (anti-pressão, anti-promessa).
> Preço, política de Continuum e guardrails CFM/LGPD **não mudaram**.
> Arquivos alterados: `script-recepcao-conversao-leads.md` (guia humano) e `reception_brain.go`
> (cérebro da IA). As frases são idênticas nos dois, para não divergirem.

---

## 1. Descoberta virou conversa em quatro tempos

**Antes:** "1 a 3 perguntas soltas" — na prática, contexto + "fez exames / tem médico?".

**Depois:** uma sequência com ordem pensada, uma pergunta de cada vez (nunca questionário):

1. **Contexto** — "o que te fez procurar um acompanhamento agora?"
2. **Incômodo** — "o que mais tem te incomodado nisso?"
3. **Peso na rotina** — "isso tem pesado no seu dia a dia?"
4. **Cenário ideal** — "como você imagina se sentir se isso entrasse nos eixos?"

**Por quê:** a ordem leva a pessoa do fato ao que de fato a move, e é isso que dá sentido ao valor
da consulta quando ele aparece. Os dois primeiros tempos costumam bastar.

**Trava de compliance:** os tempos 3 e 4 falam de **vida e rotina**, nunca de sintoma, exame ou
diagnóstico. "Tem pesado na sua rotina?" é contexto; "que exame deu alterado?" seria anamnese, e
isso continua sendo da consulta, com o médico.

---

## 2. Espelhar a palavra da pessoa

**Antes:** o guia mandava "espelhar o momento", mas sem dizer como.

**Depois:** quando a pessoa usa uma palavra carregada (exausta, perdida, travada), devolvemos a
palavra dela com uma pergunta:
> Lead: "ando exausta o tempo todo." → "quando você diz exausta o tempo todo, como isso aparece no
> seu dia?"

**Por quê:** mostra escuta, reduz a defesa, aprofunda — e não promete nada.

---

## 3. Recapitular antes de dizer o preço (lead aquecido)

**Antes:** ao perguntarem preço, a gente ancorava genérico e soltava o valor.

**Depois:** quando a pessoa já se abriu, recapitulamos o que ela trouxe e confirmamos **antes** do
valor:
> "Pelo que você me conta, o que mais pesa é [a situação dela]. A consulta foi pensada para olhar
> isso com calma: cerca de 60 minutos, painel ampliado e uma conduta para o seu caso. Faz sentido
> até aqui?"

Depois do "sim", o valor.

**Por quê:** o preço chega em terreno já preparado, não no escuro. O "faz sentido até aqui?" é um
ponto de acordo, não pedido de validação.

**Trava CFM:** a recapitulação descreve a **qualidade do cuidado**, nunca um desfecho de saúde. Não
prometemos resultado.

---

## 4. Disciplina ao dizer o valor

**Antes:** a tendência era explicar demais logo depois do preço.

**Depois:** valor em **uma frase calma, e parar**. Não emendar parcelamento, justificativa ou "mas
inclui X, Y e Z". O contexto vem antes do preço, não depois.

**Por quê:** firmeza tranquila lê como segurança; emendar álibis lê como insegurança. (O valor segue
sendo R$ 800, único e não negociável — isso não mudou.)

---

## 5. Distinguir dúvida real de evasiva

**Depois (regra nova no topo do banco de objeções):**
- **Dúvida real** (pergunta concreta: prazo, formato, o que está incluso, convênio) → responder
  direto e com clareza, e devolver o próximo passo.
- **Evasiva** ("vou pesquisar / depois eu vejo / preciso pensar" sem nenhuma pergunta) → acolher,
  segurar um próximo passo leve, **não insistir**.

**Por quê:** quem pergunta quer entender; quem foge quer espaço. Tratar evasiva com argumento empurra,
e o público da Plenya recua quando sente que está sendo conduzido à força.

---

## 6. Pedir passagem para avançar

**Depois:** em vez de empurrar para o agendamento, pedimos licença para o próximo passo: "quer que eu
te mostre como funciona?" / "quer que eu veja uma data para você?".

**Por quê:** a pessoa concede o passo; conduzimos sem forçar.

---

## 7. Seção nova no guia: o que tomamos e o que deixamos de fora

Registramos por escrito o filtro, para manutenção futura. **Rejeitamos** do curso: agressividade
comercial, urgência forçada, medo como alavanca ("imagina o que você perde"), promessa de
transformação/resultado, e "negociar escopo para dar desconto". Mantivemos só a engenharia da escuta.

---

## O que NÃO mudou (garantias)

- Preço R$ 800, único e não negociável; pagamento antecipado.
- Continuum apresentado pelo médico na consulta; recepção não fala valor dele.
- Guardrails CFM/LGPD: sem diagnóstico, sem interpretar exame, sem "medicina preditiva", sem marcas.
- Arquitetura do bot (modos copiloto/automático, slots, handoff, contrato JSON de saída) intacta.
- O recepcionista de IA só vai ao ar por decisão sua — `RECEPTION_BOT_ENABLED` segue como kill switch.

---

### Anexos
- Guia completo: [`script-recepcao-conversao-leads.md`](script-recepcao-conversao-leads.md)
- Plano e racional: [`plano-revisao-curso-vendas.md`](plano-revisao-curso-vendas.md)
- Trecho técnico do cérebro da IA: `apps/api/internal/services/reception_brain.go` (bloco
  "COMO CONDUZIR A CONVERSA" + reforços no "NUNCA FAÇA").
