# Plano — Memória da Lívia + Dossiê 360 da Pessoa + Relacionamento

> **Status:** PROPOSTA para aprovação (2026-06-06). NÃO implementar até o ok do Getúlio.
> Decisões já tomadas pelo usuário: (1) alimentam o dossiê **a IA e a equipe**; (2) registrar o
> **máximo** possível **sem ferir CFM/dados clínicos**; (3) o dossiê **migra lead→paciente** e
> continua coletando como paciente; (4) **desenhar tudo**. Mais: **tela 360** para a equipe ver/editar,
> e **sistema de avisos** (aniversários, eventos de vida) para o time interagir.
> Relacionado: [[atendimento_ia_status]] · [[livia_cerebro_processo]] · [[plenya_brand_essence]].

## 0. Princípio jurídico que organiza tudo (LGPD/CFM)

A pesquisa confirma a regra de ouro: numa clínica, os dados se separam em **administrativos/sociais**
(nome, contato, profissão, família, preferências, relacionamento) e **dados de saúde** (sintomas,
exames, diagnósticos, condutas). O **prontuário** (dados de saúde) é categoria especial, fica no EMR,
guarda de 20 anos, fiscalizado pelo CFM, e **não** é o lugar do CRM.

➡️ **Decisão de arquitetura:** o Dossiê da Lívia é um **CRM de relacionamento — só dado social/
administrativo**, fisicamente separado do prontuário. Nada clínico entra nele. Isso é o que permite
"registrar o máximo" com segurança: tudo que é social é livre; tudo que é clínico fica no prontuário
(onde já existe controle). Guardrail técnico + de prompt impede vazamento clínico para o dossiê.

## 1. Arquitetura de memória (melhores práticas pesquisadas)

Padrão consolidado em 2025/26 para agentes: **memória em duas camadas + resumo**, com **90% menos
tokens** que mandar tudo no prompt:

- **Episódica** = o transcrito completo (já temos em `lead_activities`). Não vai inteiro pro prompt.
- **Resumo rolante** = um "resumo do que já se sabe da conversa", atualizado incrementalmente.
- **Semântica** = **fatos atômicos** extraídos (ex.: "veio por indicação da filha"), com operações
  **ADD/UPDATE/DELETE/NOOP** (não só append — atualiza e remove o que mudou), validade temporal e fonte.

A Lívia, a cada resposta, recebe: **resumo rolante + últimas N mensagens + fatos do dossiê + flags
derivadas**. Assim ela "lembra de tudo" sem estourar custo, e **nunca repete pergunta já respondida**.

## 2. Modelo de dados (proposto)

Tudo keyed por **dono** (`owner_type` lead|patient, `owner_id`) — o mesmo padrão já usado em
`conversation_automations`/`conversation_suggested_replies`. Migração lead→paciente: ver §6.

### 2.1 `relationship_profile` (1 por pessoa)
- `owner_type`, `owner_id` (unique)
- `rolling_summary text` — resumo rolante da relação/conversa (Feature 1)
- `summary_updated_at`, `summary_msg_count` — controle de quando regerar
- `relationship_stage varchar` — estágio do relacionamento (novo, em conversa, agendou, paciente ativo, Continuum, inativo)
- `created_at/updated_at`

### 2.2 `relationship_fact` (N por pessoa — memória semântica)
- `owner_type`, `owner_id`
- `category varchar` — taxonomia controlada (ver §3)
- `key varchar` (ex.: "profissao", "indicado_por") + `value text` (ex.: "engenheiro", "filha Marina")
- `source varchar` — `ai` | `staff` | `form` | `consulta`
- `added_by uuid` (user, quando staff) · `confidence` (quando IA)
- `valid_from`, `valid_until` (validade temporal; UPDATE/DELETE marcam fim)
- `sensitive bool` — falso por padrão; bloqueio se classificado como clínico (não deveria entrar)
- unique parcial por (owner, key) ativo → permite UPDATE limpo
- índice por owner

### 2.3 `important_person` (N por pessoa — rede de relacionamento)
- `owner_type`, `owner_id`
- `name`, `relation` (filha, esposo, sócio, mãe…), `birthday date` (nullable)
- `notes text` (ex.: "quem incentivou o cuidado")
- `source`, `added_by`

### 2.4 `relationship_event` (N — calendário de relacionamento / avisos)
- `owner_type`, `owner_id`, `related_person_id` (nullable → evento da própria pessoa ou de alguém da rede)
- `type varchar` — `birthday` | `graduation` | `birth` | `wedding` | `loss` | `milestone` | `followup` | `custom`
- `title`, `event_date date`, `recurring bool` (aniversário=true), `lead_time_days int` (antecedência do aviso)
- `status varchar` — `pending` | `acknowledged` | `done` | `dismissed`
- `source`, `added_by`, `notes`
- Alimenta o **sistema de avisos** (§5).

> Migrations goose (00030+). Nada disso toca o prontuário.

## 3. Taxonomia de fatos (o "máximo" registrável, sem clínico)

Categorias permitidas (social/administrativo/relacional):
- **Identidade social:** como gosta de ser chamado, profissão/ocupação, cidade, idioma.
- **Família/rede:** cônjuge, filhos, quem o acompanha/influencia (liga ao `important_person`).
- **Preferências de atendimento:** canal preferido (WhatsApp/ligação), presencial vs telemedicina, melhor horário, como prefere ser abordado.
- **Contexto de chegada:** como conheceu a Plenya, quem indicou, o que motivou o contato (não-clínico).
- **Relacionamento:** estágio no funil, histórico de interações, sensibilidades de abordagem ("não gosta de insistência"), datas importantes.
- **Status derivado (calculado, não digitado):** é lead ou paciente; **Continuum ativo** (`patient_continuum`/`patient_subscription`); **cliente frequente** (nº de consultas); última consulta; nº de conversas.

🚫 **Proibido no dossiê (fica no prontuário):** sintomas, queixas clínicas, exames, diagnósticos,
medicações, condutas, qualquer dado de saúde. Guardrail: (a) instrução no prompt de extração; (b)
classificador que rejeita fatos marcados clínicos antes de gravar; (c) auditoria.

### 3.1 Clínico mencionado na conversa — curto prazo SIM, longo prazo NÃO (decisão do usuário)
Se o paciente compartilhar algo clínico na **conversa atual** (ex.: "minha memória falha", "não durmo"),
isso fica na **memória curta** — ou seja, na janela de mensagens recentes que vai pro prompt — só para a
Lívia **não repetir a pergunta durante aquela conversa**. **NÃO** é extraído para `relationship_fact`
(longo prazo) nem entra no `rolling_summary` (que é social). Se for clinicamente relevante para guardar,
o caminho é o **prontuário**, registrado pelo médico/profissional — nunca o CRM. Resumindo:
- **Curto prazo (episódico, efêmero da conversa):** pode conter o que foi dito, inclusive menção clínica → evita repetição no atendimento.
- **Longo prazo (dossiê social):** só social; clínico é descartado na extração.
- **Prontuário:** o lugar do clínico que precisa persistir (fluxo do médico, já existente).

## 4. Tela 360 para a equipe (ver + adicionar)

**Onde:** painel lateral "Dossiê" na conversa (`/conversas`) + aba "360/Relacionamento" na ficha do
paciente (`/patients/:id`) e do lead (`/leads/:id`). Mesma fonte, mesma UI.

**O que mostra (estruturado):**
- Cabeçalho: nome, foto/iniciais, **badges** (Lead/Paciente, Continuum, Frequente, canal preferido), última interação.
- **Resumo da relação** (rolling_summary, legível).
- **Fatos** agrupados por categoria, com **proveniência** (ícone IA vs pessoa, quem/quando) e botão editar/remover.
- **Pessoas importantes** (rede) com datas.
- **Próximos eventos/avisos** (aniversários, marcos) + ação rápida.
- Timeline de quem adicionou o quê (auditoria leve).

**Quem edita:** recepção/secretária, médico, profissionais (RBAC). Cada um adiciona fatos/eventos
durante conversa **e durante consulta** (um campo rápido "anotar algo sobre a pessoa" na tela clínica,
gravando no dossiê social, separado da nota clínica SOAP). A IA também adiciona automaticamente.

## 5. Sistema de avisos / lembretes (relacionamento proativo)

Best practice (CRM de relacionamento high-touch): calendário de engajamento + lembretes automáticos,
timing inteligente. Aplicado à Plenya:

- **Job diário** varre `relationship_event` e gera avisos com a antecedência (`lead_time_days`).
- **Tipos:** aniversário do paciente; aniversário de pessoa importante (se souber); marcos (formatura,
  nascimento, casamento); follow-ups pós-consulta; reativação de inativo.
- **Onde aparece:** painel "Hoje/Próximos" na Recepção + notificação ao time (reusar `NotificationService`).
- **Ação:** **auto nas coisas óbvias** (ex.: feliz aniversário simples, no tom da marca); em casos
  sensíveis ou não-óbvios, a Lívia **rascunha** e o humano revisa/envia. Allowlist do que é "óbvio"
  definida na Fase C. A equipe também pode agir manualmente.
- **Aniversário:** vem do `birth_date` do Patient quando existe; senão, é capturado quando surge na
  conversa ou em documento enviado, e vira `relationship_event` recorrente.
- **Guardrails:** opt-out por pessoa; nada de evento sensível/clínico; tom da marca (sem promoção).

## 6. Migração lead→paciente e continuidade

- O dossiê é por (owner_type, owner_id). Quando um **lead vira paciente** (conversão já existente),
  **repontar** `relationship_profile` + `relationship_fact` + `important_person` + `relationship_event`
  do `lead/<id>` para `patient/<novo_id>` (transação no fluxo de conversão). Nada se perde.
- Como paciente, a coleta continua (IA + equipe + consultas).
- Alternativa avaliada (descartada por complexidade agora): um `person_id` unificando lead+patient. Fica
  como evolução futura; o repontar resolve o caso real.

## 7. Pipeline da IA (como a Lívia usa e alimenta)

1. **Leitura (a cada resposta):** monta o prompt com `rolling_summary` + últimas N (subir 14→~40) +
   fatos do dossiê + flags derivadas. Já cobre "nunca esquecer / nunca repetir".
2. **Extração (no job, a cada 5 mensagens OU ao fim do atendimento):** lê a conversa recente e propõe
   operações de fato **social** (ADD/UPDATE/DELETE/NOOP) + atualiza o `rolling_summary` (social) quando
   cresce o suficiente. Dedup por `key`. **Classificador anti-clínico descarta menção clínica** antes de
   gravar (o clínico fica só no curto prazo da conversa — §3.1). Se surgir `birth_date`/data importante na
   conversa ou em documento, registra como evento.
3. **Custo:** resumo + fatos em vez de history bruto → ordem de 90% menos tokens em conversas longas.

## 8. Compliance (LGPD/CFM) — checklist

- Separação física do prontuário (dossiê = social/administrativo).
- Base legal: legítimo interesse/execução de relacionamento p/ dado administrativo; consentimento p/
  comunicações (já temos opt-in WhatsApp/email).
- Guardrail anti-clínico (prompt + classificador + auditoria de origem).
- RBAC de acesso ao dossiê; trilha de auditoria (quem viu/editou).
- Direito de eliminação: deletar dossiê por titular sob demanda.
- Avisos: opt-out; sem dado sensível; tom de marca.

## 9. Faseamento (cada fase atrás de revisão/aprovação)

- **Fase A — Memória (Feature 1):** subir janela 14→~40 + `relationship_profile.rolling_summary`
  mantido pelo job + injeção no prompt. Resolve "esquecer/repetir" já. (Menor risco.)
- **Fase B — Dossiê (Feature 2):** `relationship_fact` + extração IA + tela 360 (leitura) + edição manual.
- **Fase C — Rede + Eventos + Avisos:** `important_person` + `relationship_event` + job de avisos + painel.
- **Fase D — Refinos:** flags derivadas (Continuum/frequente), migração lead→paciente robusta, métricas.

## 10. Decisões — FECHADAS (2026-06-06)

1. ✅ **Extração de fatos:** rodar **a cada 5 mensagens OU ao fim do atendimento** (idle/encerramento).
2. ✅ **Aniversário:** usar `birth_date` do Patient se houver; senão, **obter** se surgir na conversa ou
   em documento enviado (extrair e gravar como evento/recorrente).
3. ✅ **Cumprimentos:** **auto nas coisas óbvias** (ex.: feliz aniversário simples); casos sensíveis ou
   não-óbvios continuam como rascunho para humano enviar. (Definir a allowlist de "óbvio" na Fase C.)
4. ✅ **Clínico na conversa:** curto prazo sim (não repetir no atendimento), longo prazo não — ver §3.1.
5. **Defaults adotados (ajustáveis):** nome interno "**Dossiê**"; taxonomia de categorias conforme §3
   (revisar quando montarmos a tela).

## 11. 🧭 Retomada após /compact — kickoff da implementação

**Estado:** plano APROVADO, decisões §10 fechadas. Pronto pra implementar **fase a fase**, cada fase com
diff pra aprovação do Getúlio antes de deploy (`RECEPTION_BOT_ENABLED` é kill switch). Ver
[[livia_cerebro_processo]] — não sair deployando sem ok.

**Pendência separada (não perder):** a revisão de TOM já aprovada pelo usuário (ritmo / acolhimento sem
bajulação / sem agressividade) ainda **NÃO foi aplicada** no `reception_brain.go`. O texto proposto está
no histórico da conversa; prod hoje (commit `80b2c963`) ainda tem a regra "máximo 2 perguntas". Aplicar
junto ou antes da Fase A, com diff.

**Fase A (começar por aqui) — Memória, menor risco:**
- migration goose (00030) `relationship_profile` (owner_type, owner_id, rolling_summary, summary_updated_at, summary_msg_count, relationship_stage).
- model + service `RelationshipProfile` (get/upsert por owner).
- subir `aiReceptionMaxMessages` 14→~40 e reduzir truncamento (`aiMaxContentChars`) em `conversation_ai_service.go`.
- job: manter `rolling_summary` (social) — extração/resumo a cada 5 msgs ou no fim do atendimento.
- `buildReceptionPrompt`: injetar `rolling_summary` + flags derivadas (lead/paciente, Continuum, frequente).
- guardrail §3.1: clínico só no curto prazo (janela), nunca no rolling_summary.

**Fases seguintes:** B dossiê (`relationship_fact` + extração + tela 360 leitura/edição) · C rede+eventos+avisos
(`important_person`, `relationship_event`, job de avisos, painel) · D flags derivadas + migração lead→paciente + métricas.

**Arquivos-chave:** `apps/api/internal/services/conversation_ai_service.go` (janela/prompt) ·
`reception_brain.go` (prompt) · `internal/scheduler/conversation_auto_reply_job.go` (extração/resumo) ·
`internal/services/conversation_service.go` · migrations `database/migrations/` (próxima 00030) ·
front: ficha paciente/lead + viewer de conversa (`apps/web/.../conversas`, `/patients`, `/leads`).
Detecção Continuum: `patient_continuum` + `patient_subscription`. Conversão lead→paciente: achar o fluxo atual.

### Fontes da pesquisa (best practices)
- Memória de agentes: mem0 (state of AI agent memory 2026; chat summarization guide), Vectorize (frameworks), Towards Data Science (guia prático), Atlan (episódica/semântica).
- Patient 360 / CRM saúde: Salesforce Health Cloud (patient relationship management), Cleveland Clinic Consult QD, Keona/Cured (healthcare CRM).
- CRM de relacionamento/eventos: crm.org, monday.com, Streak (personal CRM), Visionary Square (life-event touchpoints).
- LGPD/CFM: cartilha CFM sobre LGPD, CRM-DF manual LGPD, Migalhas (bases legais em saúde), ConJur (sigilo médico).
- Segurança de memória/PII: MemTrust (zero-trust memory), FutureAGI (PII redaction), Oracle (observability).
