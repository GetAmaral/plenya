# UI de geração do plano no EMR — estudo

> Estudo pedido em 01/09/2026, depois de gerar o deck da Ana pelo EMR e comparar com o aprovado.
> Documento vivo. O plano de implantação anterior está em
> [plano-planos-paciente.md](plano-planos-paciente.md); este trata só da **tela de autoria**.

## O que a tela precisa fazer

Compilar o prontuário, montar o rascunho visual, e sustentar a discussão até a versão final.
As três coisas existem hoje em algum grau, e nenhuma está na tela.

## Diagnóstico do que já existe

### Funciona, e a tela não mostra

`GET /patients/:id/plan-dossier` compila o prontuário inteiro. Na Ana devolveu 83 réguas, 50
achados fortes, 53 se movendo, 26 condutas do plano de cuidado, 3 prescrições e o último pedido de
exames. **O médico nunca vê isso.** É um JSON que só a API devolve e que hoje só eu leio.

O resto do motor está pronto: render 16:9 e A4 do mesmo HTML, conferência de estouro por slide,
publicação com `published_content` congelado (editar um plano publicado não tira a devolutiva do
paciente do ar).

### Existe infraestrutura de IA, e é boa

`AIService.CompleteText` já fala com a API da Anthropic, com saída em JSON estruturado
(`buildJSONSchema`) e modelo configurável por finalidade (`cfg.Claude.LabModel`, `NoteModel`).

Mais relevante: `ai_note_service.go` é o precedente de disciplina. O prompt da nota de teleconsulta
abre com "REGRAS INVIOLÁVEIS: use SOMENTE o que está na transcrição, nunca infira, nunca invente
valores numéricos". É exatamente a postura que o gerador de plano precisa, e já é a da casa.

### Os três buracos

**1. O dossiê não tem tela.** O trabalho de curadoria (o que entra no deck) acontece fora do EMR.

**2. O editor é uma caixa de JSON.** `plano/page.tsx` tem título, um `<Textarea>` com os slides
crus e cinco botões. É ferramenta de desenvolvedor.

**3. Não há histórico de rascunho, e este é o buraco estrutural.** `PatientPlanService.Update`
faz `plan.Content = req.Content` e pronto. `patient_plans.version` conta **publicações**, não
edições. Na sessão de hoje o deck da Ana passou por quatro versões (10 slides, 18, 19, 20) e as
três primeiras **não existem mais**.

> Não se discute um documento cujos estados anteriores foram apagados. Enquanto o rascunho não
> tiver revisões, "discussão" é só sobrescrita com outro nome.

## Desenho proposto

### Modelo de dados

**`patient_plan_revisions`** (novo). Toda gravação cria uma revisão; `patient_plans.content` passa
a ser o "head".

```
plan_id · seq · content jsonb · author_kind ('clinician'|'assistant') · author_user_id
       · note (o que mudou e por quê) · created_at
```

É pequeno e é o que destrava tudo o mais: diff, voltar atrás, rastro da discussão, auditoria de
quem escreveu cada frase que o paciente leu.

**`patient_plans.dossier jsonb`** (novo). O dossiê congelado no momento em que o plano nasce, com
um botão explícito de "atualizar com o prontuário de hoje".

Duas razões. A primeira é que hoje ele é recalculado a cada consulta e o chão se move no meio da
discussão: enquanto eu escrevia o deck da Ana, o escore dela mudou três vezes (60,2 → 59,3 depois
de completar a anamnese). A segunda é auditoria: dá para saber o que a máquina sabia quando o deck
foi escrito.

**`patient_plan_comments`** (depois). Thread por slide, para revisão a quatro mãos.

### A tela, em três colunas

**Esquerda — o prontuário compilado.** O dossiê virando interface.

Abas: se movendo · está bem · réguas · condutas · prescrições · exames pedidos · **lacunas**.
Cada linha traz nome, valor, unidade, nível, pontos perdidos, seta de direção e data. Um checkbox
por linha: *entra no deck*. Essa seleção é a curadoria do médico e vira a entrada do gerador.

A aba **lacunas** é a que eu mais quero: lista o que está travando o plano deste paciente. Sem
vitais aferidos. Sem data de nascimento. Resultados marcados `unit_conversion_status = 'revisar'`.
Itens de anamnese que dependem de dado do cadastro. Hoje isso só aparece se alguém rodar SQL.

**Centro — o rascunho visual.** Lista vertical de cartões, um por slide, cada um com miniatura
real. Não precisa de pipeline novo: o `preview` já devolve o HTML com um `<section class="slide">`
por slide.

No cartão: edição dos campos daquele tipo de slide (título, eyebrow, lede, punch, réguas, linhas de
tabela), selo vermelho com quantos pixels estouram, arrastar para reordenar, duplicar, apagar.
Sem JSON. O JSON continua atrás de um "avançado", para depuração.

A conferência de estouro roda **ao salvar**, não em botão separado. O botão existe hoje porque a
medição é cara; medir só o slide editado resolve.

**Direita — a conversa.** É o pedido central, e é onde mora o risco.

O médico escreve em texto livre: "tira o slide do sono", "junta os dois de colesterol", "o punch do
6 está fraco". Cada rodada do assistente **produz uma revisão com diff** (slides adicionados,
removidos, alterados) que o médico aceita ou descarta. A conversa fica gravada junto do plano, então
um segundo profissional entende por que o deck ficou como ficou.

### O contrato com o modelo

Esta é a parte que não pode ser frouxa.

**Entrada:** o dossiê congelado (ou o subconjunto selecionado na coluna da esquerda) + os slides
atuais + a instrução do médico.

**Saída:** JSON estruturado com uma **lista de operações** sobre os slides (`add`, `edit`,
`remove`, `reorder`), não um deck inteiro reescrito. Operação torna o diff trivial e reversível, e
impede que uma instrução pequena reescreva o deck todo em silêncio.

**Validação no servidor, antes de aplicar:**

- todo `ruler.code` citado tem que existir no dossiê congelado;
- todo valor numérico de régua tem que ser **igual** ao do dossiê;
- `axis`, `segments` e `history` são copiados do dossiê, nunca aceitos do modelo;
- slide que cita número fora desse conjunto é **recusado**, não exibido.

O prompt do `ai_note_service` pede "nunca invente valores numéricos". Pedir é o que dá para fazer
com uma transcrição. Aqui dá para **verificar**, porque os números do deck saem de um conjunto
fechado. Verificar é melhor que pedir.

### Onde fica a skill

Hoje a autoria acontece fora, pela skill `/plano`. Duas opções: manter fora, ou mover para um
endpoint.

**Recomendo os dois, com papéis diferentes.** Um endpoint `POST /plans/:id/draft` dá ao médico um
botão "gerar rascunho" que funciona sem mim, e é o caminho de rotina. A skill continua útil para o
primeiro passe difícil, como foi o da Ana. As duas escrevem na mesma tabela de revisões, então
interoperam sem esforço.

## Fases

| Fase | O quê | Entrega sozinha? |
|---|---|---|
| 1 | Revisões + dossiê congelado + coluna esquerda + cartões de slide no lugar do JSON | Sim: o médico vê o prontuário compilado e edita sem JSON |
| 2 | `POST /plans/:id/draft` com o contrato de operações e a validação numérica no servidor | Sim: botão "gerar rascunho" |
| 3 | Coluna da conversa, com diff aceitar/descartar por rodada | Sim: é o pedido central |
| 4 | Comentários por slide, revisão a quatro mãos | Incremental |

A fase 1 é a que precisa vir primeiro por dependência real, não por conveniência: sem revisão não
há diff, e sem diff a conversa da fase 3 não tem como ser auditada.

## O que eu não construiria

- **Editor visual de slide livre** (arrastar caixa, escolher fonte). O deck tem gramática fechada
  de nove blocos, e é isso que mantém a qualidade. Editar campos de um bloco é o certo; desenhar
  slide do zero devolve o problema para o médico.
- **Geração automática sem médico no meio.** O ranking acerta o que pesa; o arco, o título e a
  decisão de o que virar slide são leitura clínica. A comparação com o deck aprovado da Ana mostrou
  os dois lados: a máquina reproduziu sozinha os quatro achados centrais e não produziu nenhum dos
  três slides de narrativa.
- **Entidade nova para a leitura clínica.** `clinical_notes` já é isso, com "História clínica e
  evolução" e "Conduta", e tem **2 registros em toda a produção**. Expor ela no dossiê resolve; criar
  outra coisa ao lado só dá mais um lugar para a informação se perder.

## Referência do que foi medido

Tudo acima saiu de números de produção em 01/09/2026, não de estimativa:

- dossiê da Ana: 83 réguas, 50 fortes, 53 se movendo, 26 condutas, 3 prescrições
- deck do EMR sem condutas registradas: 10 slides. Com as 26 condutas: 20 slides. Aprovado: 21
- dos 20 marcadores citados nos dois decks, 18 batem com o mesmo valor
- `clinical_notes` em produção: 2 registros, 2 pacientes
- `consultation_vitals` da Ana: 0 registros
- revisões de rascunho preservadas hoje: nenhuma
