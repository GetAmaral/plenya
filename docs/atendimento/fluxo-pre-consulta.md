# Fluxo de pré-consulta Plenya (preparação pós-agendamento)

> Como o paciente envia exames antes da consulta, o Dr. analisa e pede exames complementares, e
> preenche um formulário curto de preparação. Objetivo: chegar no dia com painel completo e história
> levantada, para os ~60 minutos irem para a conduta, não para coleta de dados.

Decisões tomadas (2026-06-03, com o Getúlio):
- **Pré-análise dos exames: caso a caso, pelo Dr.** (não painel padrão genérico).
- **Dois canais de envio:** WhatsApp + CRM (com OCR/interpretação e salvamento no prontuário) **e**
  portal dedicado do paciente.
- **Formulário de preparação: curto, pós-agendamento**, no mesmo estilo da Triagem pública, porém com
  **itens do Escore curados** especificamente para esse momento (curadoria a fazer com o Dr.).

---

## 1. Princípio que organiza tudo: agendar primeiro, otimizar depois

Auditoria do EMR (2026-06-03): no estado **"lead"** não existe upload de exames nem formulário
estruturado próprio (todo upload exige autenticação; a anamnese é preenchida pela equipe depois da
consulta). O que destrava a infra (portal, upload, pedido de exame, results inbox) é o paciente
**existir no EMR**.

Logo, o **agendamento pago é o gatilho** que converte o lead em paciente. Duas consequências:

1. **Nunca colocar exame ou formulário como pré-requisito para agendar.** Isso vira barreira e derruba
   conversão. Primeiro fecha e paga, depois roda a preparação.
2. **A consulta paga é o filtro que protege a agenda do Dr.** A pré-análise caso a caso (tempo médico)
   só acontece para quem já pagou e marcou. Resolve a preocupação de "tempo não pago": ninguém é
   pré-analisado sem compromisso firmado, e a relação médico-paciente já está estabelecida quando o
   pedido de exame é emitido.

---

## 2. Visão geral do fluxo

```
LEAD
  │  (recepção qualifica e agenda; ver script-recepcao-conversao-leads.md)
  ▼
AGENDAMENTO + PAGAMENTO ANTECIPADO (R$ 800, à vista)
  │  recepção converte Lead → Paciente no EMR
  ▼
DISPARO DA PREPARAÇÃO  (mensagem automática/manual com 2 coisas)
  ├─ link do portal do paciente (magic link)
  └─ instrução para enviar exames anteriores (portal ou WhatsApp)
  ▼
ENVIO DE EXAMES ANTERIORES
  ├─ Canal A: WhatsApp → CRM → OCR/interpret-exam → salva no prontuário
  └─ Canal B: portal do paciente → upload seguro → prontuário
  ▼
PRÉ-ANÁLISE DO DR. (caso a caso)
  │  Dr. lê os exames enviados + o formulário de preparação
  ▼
PEDIDO DE EXAMES COMPLEMENTARES (IssuedDocument, com CRM)
  │  recepção envia o pedido ao paciente
  ▼
PACIENTE REALIZA OS EXAMES (laboratório de escolha)
  │  resultados sobem (portal/WhatsApp) → results inbox (/lab-results/revisar)
  ▼
FORMULÁRIO DE PREPARAÇÃO (curto, curado, estilo Triagem)
  ▼
DIA DA CONSULTA: Dr. já tem exames antigos + novos + história. 60 min para conduta.
```

> **Janela de tempo:** o loop de exames complementares só fecha se a consulta estiver marcada com
> folga (exame de laboratório leva dias). Regra prática: se a consulta for em **7 dias ou mais**, roda
> o loop completo; se for antes, a recepção só pede para o paciente enviar/levar os exames que já tem,
> e o Dr. pede os novos na própria consulta (como hoje).

---

## 3. Os dois canais de envio (quando cada um)

Os dois coexistem; o paciente usa o que for mais confortável.

| | Canal A: WhatsApp + CRM | Canal B: Portal do paciente |
|---|---|---|
| **Estado** | Funciona hoje (MVP) | A construir (fase 2) |
| **Como** | Paciente manda PDF/foto no WhatsApp; equipe vê no CRM (`/conversas`); `interpret-exam` faz OCR e salva no prontuário | Paciente entra com magic link; faz upload seguro; anexa ao prontuário |
| **Vantagem** | Zero atrito para o paciente, já é o canal natural | Estruturado, organizado, melhor para volume e para o formulário |
| **LGPD** | Consentimento explícito antes do envio; conteúdo já é cifrado em repouso no CRM | Canal próprio, consentimento no fluxo, cifrado |

A recepção **roteia** os arquivos; não interpreta nem comenta resultado. Interpretação é ato médico.

---

## 4. Formulário de preparação (curado, estilo Triagem)

### Princípio de design
- **Curto: 5 a 8 minutos.** Formulário longo antes da consulta derruba conclusão e aumenta no-show.
  O aprofundamento do Escore completo (800+ itens) fica com a equipe na consulta, nunca como pré-portão.
- **Pós-agendamento, não pré.** Pedir depois de marcar (compromisso alto) tem conclusão muito melhor.
- **Lembretes** (ex.: 24h e 2h antes) elevam a taxa de conclusão de forma relevante.
- **Mobile, sem login pesado, "curado".** O público 35-55 premium tolera formulário bem feito e
  rejeita ficha burocrática.
- **Não repetir o que o Dr. vai perguntar de novo:** o que entra no formulário, o Dr. usa.

### Como construir (reaproveitando o que existe)
A Triagem pública já é **config-driven**: itens vêm do EMR via `apps/site/scripts/sync-score-light.ts`,
gravados em `apps/site/content/data/score-triagem-config.json` (36 itens) e renderizados por
`apps/site/components/escore/escore-light-form.tsx`. O score-light completo tem 83 itens, incluindo um
grupo "Exames" (22).

O formulário de preparação é **o mesmo motor, com outro recorte de itens**:
1. Marcar no EMR quais ScoreItems compõem a "preparação" (flag/seleção curada, análoga ao `lightOrder`
   da Triagem).
2. Gerar `apps/site/content/data/score-preparacao-config.json` pelo mesmo sync.
3. Renderizar com o componente do escore-light (reuso), porém **atrás do agendamento** (vinculado ao
   paciente, não anônimo): entregue pelo portal do paciente e/ou por link pós-agendamento.
4. Respostas alimentam a anamnese/Escore do paciente, prontas para o Dr. revisar antes da consulta.

### Curadoria dos itens (a fazer com o Dr.)
Base de partida (grupos da Triagem atual, contagem entre parênteses): Alimentação (3), Movimento e
atividade física (3), Sono (5), Cognição (2), Stress (1), Composição corporal (6), Histórico de doenças
(12), Histórico Familiar de Doenças (4).

Candidatos sugeridos para o formulário de preparação (rascunho, o Dr. aprova/ajusta):
- **Queixa e objetivo principal** (campo aberto curto: "o que mais te incomoda hoje" / "o que você quer
  resolver"). Não existe na Triagem; vale acrescentar.
- **Medicações e suplementos em uso** (lista). Essencial e some no dia.
- **Histórico de doenças** (recorte do grupo de 12: as condições de maior impacto clínico).
- **Histórico familiar** (os 4 itens, ou os de maior peso: cardiovascular, metabólico, renal, oncológico).
- **Sono** (2 a 3 dos 5: duração, qualidade, ronco/apneia).
- **Composição corporal / antropometria** (peso, altura, e o que o paciente souber).
- **Consentimento LGPD** + **upload de exames anteriores**.

> Critério de corte: cada item entra só se (a) o Dr. usa de fato na consulta e (b) ajuda a decidir o
> painel de exames complementares. Tudo que for aprofundamento fica para a consulta. Meta de tamanho:
> caber em 5 a 8 minutos no celular.

---

## 5. LGPD e CFM

- **Consentimento explícito** antes de qualquer envio de exame ou preenchimento (dado de saúde
  sensível). O fluxo registra a versão do consentimento.
- **Canal seguro + cifrado.** WhatsApp/CRM já cifra conteúdo em repouso; o portal terá upload próprio.
- **Recepção não interpreta exame nem dá orientação clínica.** Só roteia e organiza. Interpretação e
  pedido de exame são atos médicos.
- **Pedido de exame com CRM, paciente já registrado.** Como a conversão lead→paciente acontece no
  agendamento, o Dr. emite o pedido (IssuedDocument) sobre um prontuário real, com relação
  médico-paciente estabelecida. Não se emite pedido para "lead solto".
- **Sem promessa de resultado** em nenhuma comunicação (publicidade médica).

---

## 6. O que existe vs. o que falta construir

Com base na auditoria do EMR (2026-06-03):

**Já existe e dá para usar agora (fase 1):**
- Captura de lead (`apps/site` `/api/leads` → EMR `lead_handler.go`).
- Conversão Lead → Paciente (`/api/v1/leads/:id/convert`).
- Magic link do portal do paciente (`patient_portal_handler.go`, `/patient-portal/auth/magic`).
- CRM de conversas com anexo + OCR (`conversation_handler.go`, `interpret-exam`).
- Upload de documento do paciente (`POST /api/v1/patients/:id/documents`).
- Pedido/documento emitido com assinatura (`IssuedDocument`).
- Results inbox para o Dr. revisar exames (`/lab-results/revisar`, `lab_result*`).

**Falta construir (fase 2):**
- **Formulário de preparação curado**, vinculado ao paciente: seleção de itens no EMR +
  `score-preparacao-config.json` + sync + reuso do `escore-light-form` atrás do agendamento.
- **Entrega da preparação** no portal do paciente (tela "antes da sua consulta": formulário + upload de
  exames num lugar só).
- **Disparo + lembretes** automáticos da preparação ao confirmar agendamento (24h/2h).
- (Opcional) trazer o resultado da Triagem pública, quando o lead já fez, para dentro da preparação,
  evitando repergunta.

---

## 7. Plano de implementação faseado

**Fase 1 (agora, sem build novo), operar pelo WhatsApp + CRM:**
1. Recepção fecha e cobra a consulta, converte lead→paciente.
2. Pede exames anteriores pelo WhatsApp (com consentimento); equipe salva no prontuário via CRM.
3. Dr. analisa e emite pedido de exames complementares; recepção envia.
4. Paciente faz; resultados entram na results inbox.
5. "Formulário" provisório como roteiro de perguntas-chave da recepção (sem app), só o essencial.

**Fase 2 (build), portal + formulário curado:**
1. Curar com o Dr. os itens do Escore que entram na preparação.
2. Implementar a config `score-preparacao` + sync + render no portal.
3. Tela de preparação no portal (formulário + upload).
4. Disparo e lembretes automáticos no agendamento.

---

## 8. Reflexo no script da recepção
As falas e mensagens da recepção para esse fluxo (pedir exames, explicar o pedido complementar, enviar
o formulário, consentimento) estão em
[`script-recepcao-conversao-leads.md`](script-recepcao-conversao-leads.md), seção "Preparação da consulta".

## Fontes
- Auditoria de infra do EMR (apps/api, apps/web, apps/site), 2026-06-03.
- `apps/site/content/data/score-triagem-config.json` (36 itens) e `score-light-config.json` (83 itens).
- `apps/site/scripts/sync-score-light.ts`, `apps/site/components/escore/escore-light-form.tsx`.
- Evidência de intake digital: preferência forte por preenchimento online e ganho de conclusão com
  lembretes (benchmarks do setor de patient intake).
