# Reavaliação — UX da Recepção/Secretária do EMR Plenya (2026-06-01)

> **Proveniência:** reavaliação conduzida por workflow multi-agente (12 agentes, ~1M tokens, 221
> fontes web reais consultadas + auditoria do codebase com evidências arquivo:linha). Pesquisa
> online de casos de uso de recepção de clínica (5 ângulos) + auditoria da UX real (5 dimensões) +
> síntese + crítico de completude adversarial. Complementa, não substitui, o plano canônico em
> `plano-recepcao-secretaria.md` (feature `/recepcao` concluída em 2026-05-31).

---

## 1. Veredito

**A UX da secretária ainda NÃO é "superb". Nota honesta ~5.3/10.** A leitura precisa é de duas
camadas com qualidades opostas:

- **Backend ~7.5** — sólido e correto: busca por blind index HMAC (CPF criptografado), EXCLUDE
  constraint anti-overlap no Postgres, recibo PDF com numeração sequencial atômica e CNPJ real,
  RBAC financeiro granular (secretária cobra mas não mexe em preço), tratamento maduro da janela
  de 24h do WhatsApp.
- **UI da secretária ~5.0** — furada justo nos pontos que mais importam: **o cadastro de paciente
  está quebrado de ponta a ponta em produção**, o handoff lead→agenda não existe, e o financeiro
  tem todos os hooks/endpoints prontos mas quase nenhuma tela de consumo.

O produto tem a fundação certa. O que falta é fechar os buracos da camada que a secretária toca o
dia inteiro. A boa notícia: a maioria das lacunas de maior impacto é "UI sobre backend que já
existe" (effort S/M), não construção do zero.

---

## 2. O que já está forte (manter)

- **Cockpit `/recepcao` com fluxo de balcão 1-clique** (Chegou → Iniciar → Concluir + Pagamento) e
  agenda viva por polling de 15s que reflete entre colunas automaticamente. É comportamento de
  "command center" onde mais importa.
- **Check-in com carimbo e tempo de espera** (`checked_in`/`in_progress`, `CheckedInAt`/`StartedAt`).
- **Busca de paciente server-side por blind index** (nome/telefone ILIKE + CPF por HMAC) — a
  GlobalSearch Cmd+K e a PatientQuickSearch usam corretamente.
- **Recibo PDF sequencial atômico** (UPSERT ON CONFLICT RETURNING) com CNPJ real, e RBAC financeiro
  granular.
- **Janela de 24h do WhatsApp + bloco LGPD/consentimento** tratados com maturidade; IA no tom Plenya
  (resumir/sugerir) reduz tempo de redação.
- **Encadeamento de contexto** (cadastrar/waitlist → paciente em contexto → `/appointments/new`);
  login roteia secretária-pura direto para `/recepcao`.

---

## 3. P0 — Bugs que quebram o fluxo central (corrigir JÁ)

Estes não são melhorias; são coisas que **a secretária não consegue fazer hoje** ou que a enganam.

| # | Bug | Evidência | Conserto |
|---|-----|-----------|----------|
| P0-1 | **Cadastro de paciente quebrado**. `lib/api/patients.ts` chama `/patients` sem prefixo `/api/v1` e lê `response.data` (incompatível com o apiClient que devolve JSON já parseado); e a rota `/patients/new` **não existe**. Não há nenhum caminho funcional de cadastro pela UI. | `lib/api/patients.ts`; rota ausente | Aposentar `lib/api/patients.ts`, apontar QuickAddPatientDialog/GlobalSearch para `lib/api/patient-api.ts` (já usa `/api/v1`); criar `app/(authenticated)/patients/new/page.tsx` reusando o form do `edit`. |
| P0-2 | **Falha de rede vira empty-state caloroso**. API caída renderiza "Nenhuma consulta para hoje" / "Tudo confirmado". Risco operacional: a secretária acredita que a agenda está vazia. | cards de `/recepcao` sem ramo `isError` | Ramo `isError` em cada card → "Não foi possível carregar" + botão "Tentar de novo", distinto do empty-state. |
| P0-3 | **Reagendar com conflito → 500 cru**. O handler `Update` não trata `ErrAppointmentConflict`. | `appointment_handler.go:205` | Mapear `errors.Is(err, services.ErrAppointmentConflict)` → 409 com mensagem PT-BR ("Este horário já foi reservado, escolha outro"), espelhando o Create; tratar no dialog. |
| P0-4 | **Leads e waitlist não fazem polling**, contrariando a promessa do cockpit e um comentário enganoso no código. Lead novo do Instagram/site só aparece após refocar a janela. | `leads-api.ts` staleTime 60s, sem refetchInterval | `refetchInterval` ~30s em `useLeads`/`useWaitlist` e corrigir o comentário. |

---

## 4. Casos de uso ESQUECIDOS (a pergunta central)

A pesquisa online cruzada com a auditoria revelou casos de uso que **não estão nas 6 features nem
no plano**. Ordenados por prioridade. `[BR]` = especificidade brasileira.

### Prioridade ALTA

| Caso de uso | Canal | Por quê |
|-------------|-------|---------|
| **Handoff lead → agendar consulta** `[BR]` | lead | O passo mais quente do funil não tem caminho. `appointments/new` aceita `patientId` mas não `leadId`; não há botão "Agendar" no lead nem no viewer de conversa. A secretária sai do CRM e remonta tudo à mão. Buraco de maior impacto para quem capta por IG/FB/site. |
| **Recall de retorno programado** (agendar retorno no checkout) | gestão | O modelo de longevidade é acompanhamento recorrente (reavaliação anual de biomarcadores, reanálise do Escore). "Concluir" hoje é desacoplado de qualquer retorno. Sem fila de recall por data-alvo, o retorno depende da memória do paciente. Receita previsível sem custo de aquisição na mesa. |
| **Ingestão de Instagram + Facebook DMs na inbox** `[BR]` | online | A inbox só conhece email+WhatsApp (`conversation_service.go:553`). IG/FB DMs são o principal canal de leads sociais do Plenya e a infra Meta já existe (`social-mcp`). Hoje a secretária responde DM no app por fora, sem registro no CRM. |
| **Régua de confirmação/lembrete do lado do PACIENTE** (48h/24h/2h) `[BR]` | gestão | O card "a confirmar" confirma do lado da secretária; não existe régua automática com botão confirmar/remarcar para o paciente. Tablestakes do mercado BR que o Plenya ainda não tem. (Sobre a magnitude da redução de no-show, ver §8.) |
| **Reativação de pacientes inativos (win-back)** | gestão | A base de quem fez avaliação inicial e não voltou é o maior reservatório de receita ociosa de uma clínica de acompanhamento contínuo. Inexistente. Exige base legal LGPD distinta da assistencial. (Validar tom, ver §8.) |
| **Alerta de lead novo + SLA visível (speed-to-lead acionável)** `[BR]` | online | Hoje leads nem fazem polling, o sidebar não tem badge de lead novo, e não há SLA por conversa. Responder rápido depende de vigilância manual. |
| **Telas financeiras de consumo** | gestão | Histórico de pagamento do paciente, reimpressão de recibo, estorno na UI, admin de preços por tipo. **Todos os hooks/endpoints já existem** (`payments.ts`), só falta UI. Hoje um erro de cobrança só se corrige via banco direto. |

### Prioridade MÉDIA

| Caso de uso | Canal | Por quê |
|-------------|-------|---------|
| **Biblioteca de respostas rápidas (canned text)** `[BR]` | online | A secretária redige à mão orientações repetitivas (horário, endereço, preparo, link). Snippets com variável de nome aceleram o dia e padronizam a voz da marca. |
| **Registrar "Não compareceu" (no_show) + ciclo de recuperação** | presencial | Status `no_show` existe no backend (`calendar-api.ts:551`) mas não há gatilho na UI. Sem registrar falta, não há recontato em 48h nem flag de risco (paciente com falta prévia tem mais chance de faltar de novo → alimenta a régua de confirmação). |
| **Pré-cadastro / intake digital antes da chegada (link WhatsApp)** `[BR]` | ambos | Anamnese de longevidade é extensa; preencher no balcão atrasa todos. Pré-intake transforma a chegada em "conferir e seguir". O portal do paciente já é a base. |
| **NPS / satisfação pós-consulta** `[BR]` | online | Termômetro de retenção e motor de boca-a-boca numa marca premium. (Fricção LGPD + tom, ver §8.) |
| **Dashboard do gestor** (ocupação, no-show, ticket médio, inadimplência) `[BR]` | gestão | O dashboard de leads cobre só o topo do funil. Sem ocupação/no-show/ticket o gestor decide no escuro. |
| **Sinal antecipado por link de pagamento (PIX/cartão)** `[BR]` | gestão | Encaixa no inbox que a secretária já usa; reduz no-show em agenda particular de poucos slots. Eticamente permitido (CEM veda cobrar por atendimento NÃO realizado, não o sinal). (Magnitude vendor, ver §8.) |
| **Encaixe ativo da lista de espera** | gestão | A waitlist hoje é passiva; não notifica quando uma vaga é liberada por cancelamento/no-show. |
| **Conferência de identidade/documentos na chegada com registro** `[BR]` | presencial | Identificação errada é risco clínico e legal. Conferir e atualizar contato na chegada (discretamente) mantém a base limpa para cobrança e comunicação. *(Achado adicionado pelo crítico: a síntese havia omitido.)* |
| **Gestão proativa de fila / transparência de espera** | presencial | O Plenya já mede o tempo de espera mas não o usa para comunicar. "Espera comunicada parece mais curta." Hospitalidade premium coerente com o anti-"drip bar". *(Crítico.)* |
| **Política de tolerância a atraso (encaixe vs. remarcar)** | presencial | Sem regra, cada atraso vira negociação de balcão. Política clara protege os pacientes pontuais. *(Crítico.)* |
| **Reativação de LEADS frios** (distinto de win-back de pacientes) | online | Interessados que nunca agendaram são um reservatório separado dos pacientes inativos. A síntese fundira os dois. *(Crítico.)* |
| **Origem/atribuição do lead** (incl. "como nos conheceu?" + indicação) | online | O lead detail já exibe UTM; falta capturar origem de leads crus e rastrear indicação de paciente (origem de maior qualidade). *(Crítico.)* |

### Prioridade BAIXA

| Caso de uso | Por quê |
|-------------|---------|
| **Agenda com recurso de SALA** (centro de infusão/exames) `[BR]` | A clínica física prevê infusão + exames além dos consultórios (memória `clinica_fisica`). Sala é recurso escasso que precisa entrar no agendamento, não só o médico. |
| **Datas de relacionamento** (aniversário / 1 ano de acompanhamento) `[BR]` | Toque relacional de menor custo. Exige base legal LGPD adequada. Pós-fundação de recall/reativação. |
| **Painel de chamada com privacidade / "sala pronta"** `[BR]` | Para a clínica física LGPD-sensível: chamada discreta (sem expor nome/motivo) + sinalizar ao médico que o paciente chegou. Diferenciador, não tablestakes. |

---

## 5. Auditoria de UX por dimensão

| Dimensão | Nota | Principais achados |
|----------|:----:|--------------------|
| **Cockpit `/recepcao`** | 6.5 | Forte no balcão 1-clique; fraco em estados de erro (P0-2), polling de leads/waitlist (P0-4), e em usar o tempo de espera para comunicar a fila. |
| **Agendamento / Calendário** | 6.5 | Multi-médico bom; reagendar custa 5+ passos, sem drag-and-drop nem clique em slot vazio; janela de horas 08-20 hardcoded; sem sombreamento de expediente/ausência; `statusConfig` só tem 4 status (ignora `checked_in`/`in_progress`/`no_show`); reagendar não avisa se o paciente será notificado. |
| **Leads + Conversas** | 6.0 | Inbox unifica por contato mas só email+WhatsApp; sem SLA na lista de conversas (não distingue "esperando há 3h" de "já respondemos"); filtro multi-status filtra só a página atual (client-side sobre 25 itens, contagem enganosa); `alert()` nativo no envio e metadados como JSON cru (`leads/[id]/page.tsx:413,169`). |
| **Cadastro + Busca de paciente** | 4.5 | **Mais baixa.** Cadastro quebrado (P0-1); lista `/patients` bufferiza 500 e filtra no cliente (teto silencioso, busca por CPF não funciona na lista); sem dedupe por CPF (violação de unique estoura 500); busca de nome não normaliza acento (Joao ≠ João, `patient_service.go:174`). |
| **Financeiro + Arquitetura de Informação** | 5.5 | Backend financeiro completo, UI de consumo quase ausente (histórico/estorno/admin de preços); sem badge "pago" na agenda nem dedupe de pagamento; superfície clínica da página de paciente não é gateada para roles não-clínicas; secretária pula entre 5 páginas no dia (`/recepcao`↔`/calendario`↔`/conversas`↔`/leads`↔`/patients`). |

---

## 6. Roadmap priorizado

### P0 — Bugs (corrigir já, antes de qualquer feature nova)
1. Consertar cadastro: aposentar `lib/api/patients.ts` + criar `/patients/new` (M)
2. Tratar `isError` em todos os cards do cockpit (M)
3. Mapear 409 de conflito no reagendamento (S)
4. Polling em leads/waitlist + corrigir comentário (S)

### P1 — Alto impacto, backend já existe (UI sobre o que está pronto)
5. Botão "Agendar consulta" direto do lead e da conversa (handoff) (M) — fazer `/appointments/new` aceitar `leadId`
6. Liberar "Configurar Agenda" para a secretária (front + RBAC: `RoleSecretary` em `canManageDoctor`, `working_hours_handler.go:34-42`) (S)
7. Telas financeiras de consumo: histórico do paciente + reimpressão + estorno + admin de preços (M)
8. Dedupe de paciente por CPF na criação + mensagem amigável (M)
9. Status de pagamento na agenda + impedir cobrança duplicada (M)
10. Busca server-side na lista de pacientes (eliminar buffer de 500) (M)
11. Alerta de lead novo no sidebar + SLA na inbox de conversas (não só leads) (M)
12. Biblioteca de respostas rápidas no composer + EmailReplyBlock (M)

### P2 — Novos casos de uso de maior retorno (construção)
13. Régua de confirmação/lembrete automática do lado do paciente (48h/24h/2h) (L)
14. Recall de retorno programado + agendar retorno no checkout (L)
15. Ingestão de Instagram + Facebook DMs na inbox (L) — infra Meta já existe
16. Reativação de inativos (win-back) — validar tom (L)
17. Registrar `no_show` na UI + ciclo de recuperação 48h + flag de risco (M)
18. Criar consulta clicando em slot vazio + drag-and-drop para reagendar (L)
19. Encaixe ativo da waitlist quando vaga é liberada (L)

### P3 — Refinamentos e diferenciadores premium
20. Pré-cadastro/intake digital por link de WhatsApp (L)
21. Dashboard do gestor (ocupação/no-show/ticket/inadimplência) (L)
22. Janela de horas dinâmica + sombreamento de expediente/ausência na grade (L)
23. Sinal antecipado por link de pagamento (L) — validar (§8)
24. NPS pós-consulta (L) — resolver LGPD + tom (§8)
25. Recurso de SALA na agenda (L), datas de relacionamento, painel de chamada com privacidade
26. Polimentos: busca unaccent, marcador de "agora" na agenda, overflow de ações, acentuação de strings, metadados de lead legíveis, toast em vez de `alert()`

---

## 7. Achados de auditoria que o crítico resgatou (não perder)

O crítico de completude flagou achados reais que a síntese diluiu. Incorporados acima, registrados aqui:

- **SLA é da inbox de conversas INTEIRA** (pacientes incluídos), não só de "lead novo". A lista não
  ordena por mais-antiga-sem-resposta nem exibe `lastInboundAt` como "aguardando".
- **Filtro de leads server-side** (multi-status/range/ordenação) — hoje filtra só a página atual.
- **Lista global de consultas para staff** (sem forçar contexto de paciente) + ampliar `statusConfig`.
- **Reagendar/cancelar não comunica** se o paciente será notificado; telefone não é clicável (tel:/wa.me).
- **Busca insensível a acento** (`unaccent`) — barato e br_relevant.
- **Metadados de lead como JSON cru** e **`alert()` nativo** → padronizar feedback (toast).

---

## 8. Ressalvas: o que NÃO transplantar dos EUA / validar com o Getúlio

A pesquisa trouxe muitos números de blogs de fornecedores (vendor-claims) e padrões americanos.
Antes de virar roadmap, filtrar (regra de marca: [[plenya_no_padroes_americanos_transplantados]]):

- **Fiscal: NFS-e (PJ), NÃO Receita Saúde (PF).** A Plenya fatura como **PJ** (Plenya Serviços de
  Saúde Ltda., CNPJ 66.991.259/0001-50), então o documento fiscal é **NFS-e municipal**, não o app
  Receita Saúde (que é obrigatório só para médico PF, pessoa física, desde jan/2025). O recibo PDF
  numerado atual é comprovante de pagamento legítimo, mas **não substitui a NFS-e**. Integração
  NFS-e é o caminho fiscal correto; nunca recomendar Receita Saúde para o faturamento da clínica PJ.
- **Magnitudes de no-show são vendor-claims.** "No-show de 50%", "confirmação reduz 38-60%",
  "sinal eleva comparecimento de 65% para 95%" vêm de blogs de software médico, não de literatura
  revisada. A direção está certa (confirmar reduz falta); a magnitude para uma clínica de
  longevidade de público engajado e poucos slots é provavelmente bem menor. Priorizar pela lógica
  (no-show particular = receita perdida direta), não pelo número.
- **Self-booking 24/7 não é ganho inequívoco.** Em clínica concierge/relacional onde o toque humano
  É o produto, muitos pacientes de alto valor (e mais velhos) preferem falar com a secretária.
  Avaliar como opção, não como tablestakes. O handoff lead→agenda (humano) tem retorno maior e menos
  risco de posicionamento.
- **Win-back / "reativar custa 5-25x menos"** é clichê de e-commerce. A mecânica de "campanha
  multi-toque" precisa caber no tom anti-coach/anti-propaganda da marca
  ([[linkedin_tom_anti_propaganda]]). Para uma base pequena com relação pessoal, reativação é mais
  "o médico lembrou de você" que "campanha".
- **NPS por WhatsApp tem fricção LGPD própria.** Disparar pesquisa é uso relacional/marketing, exige
  base legal distinta da assistencial; mandar logo após consulta clínica beira uso secundário de
  dado sensível. Resolver a base legal antes de implementar.
- **Régua de lembrete e reativação exigem base legal LGPD por TIPO de contato** (assistencial vs.
  marketing) com opt-out automático. O bloco LGPD do lead é um começo; falta a base por finalidade.

---

## 9. Apêndice — fontes consultadas (221 no total)

Pesquisa com buscas reais (WebSearch/firecrawl), URLs citadas pelos agentes. Amostra representativa
por ângulo:

- **Balcão presencial (54):** curogram.com (patient check-in / front-desk automation), conclinica.com.br
  (fluxo de atendimento), portaltelemedicina.com.br, pixeon.com, blog.iclinic.com.br, blog.apolo.app,
  wavetec.com (queue management), wingassistant.com.
- **Captação/leads (35):** leadcareteam.com (speed-to-lead), influxmd.com (lead conversion 2025),
  doctible.com, blog.meets.com.br, weramp.com.br (funil clínicas), pro.doctoralia.com.br,
  healthcare-marketing.agency (patient funnel/nurturing).
- **Omnichannel (42):** clint.digital (automações WhatsApp 2026), socialhub.pro, clinicorp.com
  (confirmação WhatsApp), amplimed.com.br (chatbot), respond.io (multichannel), quickconnect.biz
  (unified inbox), dealism.ai, zealousweb.com (no-show reminders).
- **Financeiro BR (36):** conclinica.com.br (fechamento de caixa), validadortiss.com.br (PIX),
  contmatic / clinicorp / contabilizei (NFS-e, Receita Saúde), **portal.cfm.org.br** (recibos fiscais
  Receita Saúde), gov.br/receitafederal (prazo Receita Saúde 2026).
- **Retenção/gestão BR (54):** gestaods.com.br (LTV), elogrowth.com.br (indicação/reativação),
  news.hidoctor.com.br, conclinica.com.br (campanha de reativação / recall), imaginasoft.pt (recall),
  comtele.com.br (confirmação), **socialhub.pro/lgpd-clinica-whatsapp-2026**, oabcampinas.org.br (LGPD clínicas).

---

**Status:** reavaliação concluída em 2026-06-01. Próximo passo aguarda decisão do usuário sobre o
escopo a executar (sugestão: começar pelos 4 bugs P0, que são baratos e destravam o fluxo central).
Quando um escopo for aprovado, atualizar `plano-recepcao-secretaria.md` e a memória
[[recepcao_secretaria_status]].
