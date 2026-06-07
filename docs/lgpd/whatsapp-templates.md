# Templates WhatsApp Business — registro versionado

Registro de todos os templates submetidos à Meta WhatsApp Cloud API e seu status de aprovação.
Este arquivo é a **fonte de verdade humana** — quando atualizar texto, bumpe a versão e mantenha o histórico.

---

## Template: `magic_link`

| Campo | Valor |
|-------|-------|
| **Nome técnico** | `magic_link` |
| **Categoria** | `utility` |
| **Idioma** | `pt_BR` |
| **Variáveis** | `{{1}}` = URL do magic link |
| **Sub-tipo botões** | nenhum (link no body) |
| **Status atual** | ⏳ PENDING (submetido 2026-06-05 na WABA real `827642893502160`, id `1413933667249898`) |
| **Justificativa categoria** | Mensagem transacional vinculada a ação iniciada pelo usuário (claim do Escore Light). Sem cunho promocional. |

### Versão 1.0 (2026-04-23)

```
Olá! Aqui é da Plenya. Seu resultado do Escore Plenya Light está pronto. Acesse pelo link seguro: {{1}}. Esse link é único e expira em 7 dias. Se não foi você, ignore esta mensagem.
```

**ATENÇÃO:** o "7 dias" precisa bater com o JWT TTL em
`apps/api/internal/services/anonymous_score_service.go` (busca por `7 * 24 * time.Hour`).
Se mudar um, mudar o outro **e** ressubmeter o template à Meta (texto novo = nova aprovação).

### Histórico de versões

| Versão | Data | Mudança | Submetido | Aprovado |
|--------|------|---------|-----------|----------|
| 1.0    | 2026-04-23 | Criação inicial (7 dias) | 2026-06-05 | ⏳ PENDING |

---

## Processo pra adicionar/alterar template

1. Editar o arquivo aqui primeiro (versão + texto + categoria + justificativa).
2. Acessar Meta Business Manager → WhatsApp Manager → Message templates → Create template.
3. Colar o texto exato. Selecionar categoria + idioma `pt_BR`.
4. Submeter. Aprovação Meta leva até 24h.
5. Quando aprovado, atualizar status nesta tabela.
6. Se rejeitado, anotar motivo aqui e iterar (categoria errada é o motivo mais comum — utility ↔ marketing).

## Princípios

- **Categoria correta importa MUITO.** Marketing custa 5×–10× mais que utility e exige opt-in explícito por mensagem (não basta opt-in geral).
- **Conteúdo proibido:** ofertas, descontos, "promoção", linguagem persuasiva forte → vira marketing automaticamente.
- **Variáveis:** usar `{{N}}` numéricas, nunca nomeadas.
- **Botões:** evitar URL_BUTTON com variável dinâmica; preferir link no body (mais simples, aprovado mais rápido).
- **Mudança de texto** após aprovação requer **nova submissão**. Meta NÃO permite editar in-place.

## Aprovações pendentes

- [x] `magic_link` v1.0 (utility, pt_BR) — **submetido 2026-06-05** na WABA real, ⏳ PENDING
- [x] `lead_alert` v1.1 (utility, pt_BR) — **submetido 2026-06-05** na WABA real, ⏳ PENDING (v1.0 recusada, ver abaixo)

---

## Template: `lead_alert` (Phase 2)

| Campo | Valor |
|-------|-------|
| **Nome técnico** | `lead_alert` |
| **Categoria** | `utility` |
| **Idioma** | `pt_BR` |
| **Variáveis** | `{{1}}` = nome do lead · `{{2}}` = contato (email/phone) · `{{3}}` = origem · `{{4}}` = URL admin |
| **Sub-tipo botões** | nenhum |
| **Status atual** | ⏳ PENDING v1.1 (submetido 2026-06-05 na WABA real `827642893502160`, id `1026052723178640`) |
| **Recipients** | Vendedores Plenya (telefones em `User.ProfessionalPhone`, opt-in via flag) — **NÃO é mensagem ao paciente**. É notificação interna ao time. |
| **Justificativa categoria** | Notificação operacional vinculada a evento (lead capturado). Time recebeu por ser parte do processo de vendas — base legal: legítimo interesse + consentimento na contratação. |

### Versão 1.1 (2026-06-05) — submetida

```
Um novo lead foi registrado no sistema da Plenya e aguarda retorno do time comercial. Nome: {{1}}. Contato: {{2}}. Origem: {{3}}. Acesse o painel administrativo pelo link {{4}} para ver os detalhes e dar seguimento ao atendimento.
```

Ordem dos parâmetros preservada ({{1}}=nome, {{2}}=contato, {{3}}=origem, {{4}}=URL admin), então
o backend (`SendLeadAlert`, params `[leadName, contact, source, adminURL]`) não muda.

### Versão 1.0 (2026-04-23) — RECUSADA pela Meta (2026-06-05)

```
Novo lead Plenya: {{1}} ({{2}}). Origem: {{3}}. Abra no admin: {{4}}
```

Dois motivos de recusa, corrigidos na v1.1:
1. `error_subcode 2388293` — proporção variáveis/texto: 4 variáveis num texto curto demais. Fix: alongar o texto fixo.
2. `error_subcode 2388299` — variável no fim do modelo (`...{{4}}`). Meta não permite variável no início nem no fim. Fix: texto fixo depois do `{{4}}`.

**Atenção Meta:** templates de "internal alerts" ocasionalmente são reclassificados como marketing
se tom soa promocional. Manter texto neutro/factual reduz risco.

### Histórico de versões

| Versão | Data | Mudança | Submetido | Aprovado |
|--------|------|---------|-----------|----------|
| 1.0    | 2026-04-23 | Criação inicial | 2026-06-05 | ❌ recusada (ratio + var no fim) |
| 1.1    | 2026-06-05 | Texto alongado, var fora das pontas | 2026-06-05 | ⏳ PENDING |

---

## Família "ciclo da consulta" + reengajamento (criados 2026-06-05, WABA real `827642893502160`)

Todos UTILITY · pt_BR · ⏳ PENDING. Voz Plenya (anti-pressão, sem travessão). Variáveis nunca no
início/fim do corpo. Enviados via `ConversationService.SendWhatsAppTemplate` (genérico) — para
disparo automático (véspera/dia/semana/pós) falta um **job agendador** no backend (a criar; hoje o
envio é manual pela tela `/conversas` ou por integração futura). Os nomes técnicos são a referência.

### `confirmacao_consulta_semana`
Variáveis: `{{1}}`=nome · `{{2}}`=data · `{{3}}`=hora · `{{4}}`=modalidade.
```
Olá, {{1}}. Sua Consulta Plenya está marcada para {{2}}, às {{3}}, na modalidade {{4}}. Na semana da consulta enviaremos as orientações de preparo. Se precisar ajustar a data ou o horário, responda esta mensagem que a equipe ajuda.
```

### `confirmacao_consulta_vespera`
Variáveis: `{{1}}`=nome · `{{2}}`=hora · `{{3}}`=modalidade.
```
Olá, {{1}}. Passando para lembrar que sua Consulta Plenya é amanhã, às {{2}}, na modalidade {{3}}. Leve seus exames anteriores e a lista de medicações em uso. Para confirmar a presença ou remarcar, é só responder esta mensagem.
```

### `confirmacao_consulta_dia`
Variáveis: `{{1}}`=nome · `{{2}}`=hora · `{{3}}`=detalhe (ex.: "O atendimento é presencial, na nossa unidade em Londrina" ou "O atendimento é online; o link de acesso chega em seguida").
```
Olá, {{1}}. Sua Consulta Plenya é hoje, às {{2}}. {{3}}. A equipe já está preparada para te receber. Se tiver qualquer imprevisto, responda por aqui que a gente reorganiza.
```

### `followup_pos_consulta`
Variáveis: `{{1}}`=nome.
```
Olá, {{1}}. Foi muito bom te receber na Plenya. Seu plano e suas orientações ficaram registrados, e a equipe segue acompanhando seu cuidado ao longo do tempo. Se surgir alguma dúvida sobre as condutas combinadas na consulta, responda por aqui.
```

### `reengajamento_lead`
Variáveis: `{{1}}`=nome. **Risco de reclassificação para MARKETING** (reengajar lead frio é, aos
olhos da Meta, promocional). Submetido como UTILITY; se a Meta reclassificar/recusar, vira MARKETING
(exige opt-in explícito por mensagem + custa mais). Inclui opt-out "responda PARAR".
```
Olá, {{1}}. Aqui é da Plenya. Conversamos sobre a Consulta Plenya e você ficou de avaliar com calma. Seguimos à disposição para esclarecer dúvidas e ver uma data quando fizer sentido para você. Se preferir não receber mais mensagens, responda PARAR.
```

---

## Template: `aniversario_plenya` (Lívia — auto-cumprimento de aniversário)

| Campo | Valor |
|-------|-------|
| **Nome técnico** | `aniversario_plenya` |
| **Categoria** | `MARKETING` (saudação, não-transacional; submetido já como MARKETING) |
| **Idioma** | `pt_BR` |
| **Variáveis** | `{{1}}` = primeiro nome |
| **Status atual** | ⏳ PENDING (submetido 2026-06-07 na WABA real `827642893502160`, id `1304293918582848`) |
| **Justificativa categoria** | Cumprimento de relacionamento, sem transação nem ação iniciada pelo usuário → MARKETING (honesto; evita reclassificação). |
| **Uso** | `RelationshipEventReminderJob` envia automático no dia do aniversário PRÓPRIO (não-restrito), só em modo `auto` + `WHATSAPP_TEMPLATE_BIRTHDAY=aniversario_plenya` + kill switch ligado. Senão, vira aviso ao time. |

### Versão 1.0 (2026-06-07)
```
Oi, {{1}}! Hoje é seu aniversário e a equipe da Plenya passa para desejar um dia feliz e um ano de saúde e bem-estar. Conte com a gente.
```

### Histórico de versões
| Versão | Data | Mudança | Submetido | Aprovado |
|--------|------|---------|-----------|----------|
| 1.0    | 2026-06-07 | Criação inicial | 2026-06-07 | ⏳ PENDING |
