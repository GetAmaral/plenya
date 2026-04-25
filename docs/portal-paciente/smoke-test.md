# Portal do Paciente — Smoke Test

Teste manual end-to-end. Roda em `http://localhost:3000` (host = `minha.localhost` simulado via `?host=` ou `/etc/hosts`).

> Em produção, o subdomínio `minha.plenyasaude.com.br` é detectado pelo middleware Next.js e o tráfego é reescrito para `/patient-portal/*` automaticamente.

## Pré-requisitos

1. API rodando (`docker compose up -d api`)
2. Web rodando (`docker compose up -d web`)
3. Banco com pelo menos 1 paciente real (via UI EMR ou seed)

---

## Passo 1 — Convite via staff

1. Como admin, abrir `app.plenyasaude.com.br/patients/<id>`
2. Ver card "Acesso à área do paciente" → status "Sem convite"
3. Marcar checkbox "Email", clicar "Convidar para o portal"
4. Toast "Convite enviado por email"
5. Card mostra link gerado (botão Copy)
6. Esperado: `minha.plenyasaude.com.br/auth/invite?token=<hex64>`

## Passo 2 — Consume invite + senha opcional

1. Abrir o link em janela anônima
2. Tela "Olá, [nome]! Crie senha (opcional) ou continue só com magic link"
3. Definir senha (mínimo 8 chars) → Salvar e entrar
4. Cai em `/` (home dashboard)

## Passo 3 — Home dashboard

1. Header: "Olá, [primeiro nome]"
2. Pendências (banner âmbar) se houver consulta a confirmar nas próx 48h
3. Cards: Próxima consulta, Continuum, Último escore, Boxes, Mensagens
4. Empty states elegantes em cards sem dados

## Passo 4 — Continuum read-only

1. `/continuum`: barra de progresso, próximo marco destacado
2. Plano integrado markdown renderizado
3. Timeline agrupada por mês com badges status
4. Click numa consulta ancorada → vai pra `/consultas/[id]`

## Passo 5 — Consultas + telemed

1. `/consultas` mostra Próximas/Histórico
2. Detalhe `/consultas/[id]`: confirmar presença → toast + badge "Confirmada por você"
3. Equipe vê `PatientConfirmedAt` no EMR
4. Solicitar reagendamento → dialog motivo → notificação chega pra secretária no EMR
5. Se `type=telemedicine` e dentro de janela (-30min..+2h): card "Sala da consulta" com Daily.co inline

## Passo 6 — Lobby standalone

1. Sair do portal (logout)
2. Acessar `minha.plenyasaude.com.br/sala/<token>` (token gerado quando email/WA enviado)
3. Sem auth, mostra Olá [primeiro nome] + countdown se janela fechada
4. Botão "Entrar na sala" aparece automaticamente quando abre
5. Click → DailyCoEmbed inline

## Passo 7 — Mensagens

1. `/mensagens`: empty state se primeira vez
2. Escrever mensagem → Ctrl+Enter envia
3. Bubble própria à direita (cor primary)
4. Equipe vê alerta in-app no EMR (NotificationGeneral)
5. Equipe responde no EMR via Central de Conversas
6. Resposta aparece no portal em até 30s (polling)

## Passo 8 — Exames + prescrições + avaliações

1. `/exames` tabs: Resultados, Pedidos, Prescrições, Avaliações
2. Click num batch → detalhe com tabela de valores
3. Avaliação física → abre HTML completo em nova aba (autenticado)

## Passo 9 — Escores

1. `/escores`: radar do mais recente + timeline
2. Delta vs anterior (verde/âmbar)
3. Click numa entry Light → abre radar público no site

## Passo 10 — Perfil + LGPD

1. `/perfil` tab Dados: editar telefone/cidade/UF → Salvar
2. Tab Segurança: trocar senha (precisa 8+)
3. Tab LGPD:
   - "Baixar dados" → JSON com tudo
   - "Solicitar exclusão" → motivo opcional → notificação pra admin

## Passo 11 — Boxes + Documentos

1. `/boxes`: cards com tracking e status
2. `/documentos`: lista com botão Baixar (autenticado)

## Passo 12 — Família

1. `/familia`: clicar "Convidar"
2. Email + label + scope checkboxes → Enviar
3. Card aparece com "Convite enviado"
4. Editar permissões inline → Salvar
5. Revogar → desaparece

## Passo 13 — Auth alternativo

1. Logout
2. `/login` → "Receber link mágico" → email
3. Clicar link → entra (`/auth/magic?token=...`)
4. `/esqueci-senha` → mesmo fluxo magic link

---

## Smoke automatizado (futuro)

- Playwright cobrindo passos 2, 3, 7, 10
- API tests com testcontainers Postgres
- Por ora, manual via checklist acima

## Rate limit

- `/api/v1/patient/me/*`: 120 req/min por IP
- `/api/v1/auth/patient/*`: 10 req/min
- `/api/v1/sala/:token`: 60 req/min

## Audit log

`AuditLog` middleware no `patientMe` registra todas operações na tabela `audit_logs`:
- userID + action (view/create/update/delete) + resource + resourceID

Útil pra rastrear acessos LGPD e investigar incidentes.
