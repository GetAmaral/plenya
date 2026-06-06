# Atendimento IA — plano FECHADO (2026-06-06)

> Nome: **Atendimento IA**. Estados: **Off · Copiloto · Auto**. **Quem manda é o GLOBAL**;
> conversas individuais sobrepõem temporariamente (flag de 24h).
> Mapeamento p/ enum existente (não mudar valores armazenados): Off=`off`, Copiloto=`copilot`, Auto=`auto`.

## Modos (comportamento)

- **Off:** IA silenciosa. Se chegar inbound e ficar **> Z** sem resposta **e** a conversa estiver Off
  *por herança do global* (sem flag Off explícita) → **alerta o admin** (que pode entrar e ligar o Auto).
- **Copiloto:** após o debounce **X**, a IA **gera o rascunho** (idêntico ao que o Auto enviaria) e
  **preenche o campo de mensagem** (quando a conversa está aberta na tela). **Espera o humano enviar.**
  Se ficar **> Y** sem resposta ao lead → **escala p/ Auto** (regera e envia).
- **Auto:** após o debounce **X**, gera a resposta. Se a conversa estiver **aberta** na tela, preenche o
  campo e mostra um **countdown de X/3** pro atendente ver/intervir; se ninguém mexer, **envia** ao fim do
  X/3. Se a conversa **não** estiver aberta, **envia direto** após X.

> A IA só age quando há um **inbound do lead aguardando resposta** — nunca fala sozinha sem o lead ter
> mandado algo. Handoff (dúvida clínica / pedido de humano): **não envia**, alerta (canal do Z) e tira a
> conversa do Auto até alguém agir.

## Três tempos configuráveis

| | O quê | Default | Observação |
|---|---|---|---|
| **X** | **debounce** + atraso da resposta do Auto. Conta do **último inbound do lead**; **reseta a cada nova msg** da rajada. | **30 s** | usado também no Copiloto (não draftar no meio da rajada). Em Auto + tela aberta, há +X/3 de preview antes do envio. |
| **Y** | sem resposta no Copiloto por Y → escala p/ Auto (regera e envia) | **10 min** | vale p/ Copiloto via global ou via flag |
| **Z** | sem resposta no Off por Z → alerta o admin | **20 min** | só p/ Off **por herança** (flag Off explícita NÃO alerta) |

## Janela de horário do Auto (tela de config, estilo Google)

- **Parametrizável por dia da semana** com faixas de horário (ex.: 18:00→08:00 seg–sex; dia todo no fds).
- **Dentro da janela → Auto.** **Fora da janela → vale o global** (o baseline que o usuário setou).
- **Opção de desligar** a janela (aí vale o global o tempo todo).
- Timezone: **America/Sao_Paulo**.

## Máquina de estados por conversa

- **Conversa nova, sem flag:** segue o **global** (que pode estar elevado a Auto pela janela de horário).
- **Usuário escolhe um modo na conversa** (ex.: global=Auto, marca Copiloto) → vale o override.
- **Auto + humano digita e envia** (ou mexe no campo durante o countdown X/3) → vira **Copiloto** ("assumiu").
- **Copiloto + > Y sem resposta ao lead** → escala p/ **Auto**.
- **Flag Off explícita** → **não** dispara o alerta Z (Z é só p/ conversas novas / sem flag / Off herdado).
- **Validade da flag: 24h após a última mensagem (qualquer direção)** → conversa perde a flag e volta ao
  global. (Alinha com a janela de 24h do WhatsApp.)
- **Hierarquia:** flag de conversa válida **vence** o global.

## Configuração / permissões
- Quem configura o global + tempos X/Y/Z + janela de horário: **admin, secretary, manager, doctor**.

## Arquitetura (build)
- **Rascunho no servidor (Copiloto E Auto):** o rascunho nasce no backend (debounced por X) e fica
  **salvo na conversa** (`suggested_reply` + timestamp + modelo). Mostrado no compositor quando a conversa
  está aberta; badge "rascunho pronto" na lista. No Auto, vira envio (com preview X/3 se aberto).
- **Job** (evolui o `conversation_auto_reply_job`): orquestra X (debounce→draft/preview/send), Y
  (copiloto→auto), Z (off→alerta), janela de horário, expiração de flag 24h, handoff→alerta.
- **Storage:** setting global runtime singleton (modo baseline + janela semanal + X/Y/Z) + auditoria;
  flag por conversa com timestamp (expiração 24h) e origem (herdada vs explícita, p/ a regra do Z).
  Migrations goose.
- **UI:** controle global "Atendimento IA: Off/Copiloto/Auto" + **tela de config** (janela semanal +
  tempos) + **banner de estado** sempre visível ("Atendimento IA: Auto até 08:00"). Preview X/3 com
  opção de assumir/cancelar no compositor.
- **Guardrails mantidos:** kill switch `RECEPTION_BOT_ENABLED`, opt-out LGPD, anti-spam/hora, janela 24h
  (free-form só dentro; fora = template).

## Ordem de build
1. Backend: model/migration (setting global + flag por conversa com timestamp/origem) → serviço (resolver
   modo efetivo: flag → janela → global) → rotas (GET/PUT global + config) → evolução do job (X/Y/Z +
   janela + draft persistido + handoff). Compilar no container.
2. Frontend: controle global + tela de config + banner + preview X/3 no compositor + alinhar rótulos do
   toggle por conversa p/ Off/Copiloto/Auto.

## Status
**Plano fechado.** Aguardando "go" pra começar pelo backend.
