# Estudo — sessão/login persistente em dispositivos conhecidos (EMR + PWA iOS)

**Pedido (2026-06-09):** o PWA na tela de início do iPhone funciona bem (recebe push), mas faz
**logout automático rápido demais**. Queremos manter o login por **dias** em dispositivos
conhecidos. Estudo: causa, melhor forma, e o que a lei diz. **Ainda não implementado.**

## Como o auth funciona hoje (fatos do código)

- **Access token JWT = 15min** · **Refresh token JWT = 168h (7 dias)** (`JWT_ACCESS_EXPIRY` /
  `JWT_REFRESH_EXPIRY` em `config.go`).
- Tokens guardados **no client, em `localStorage`** (`plenya-auth`, Zustand `persist`).
- Refresh tokens **persistidos em `refresh_tokens`** (hash, `user_agent`, `ip_address`,
  `expires_at`, `revoked_at`, `used_at`) — **rotação de uso único com revogação**: a cada
  `/auth/refresh`, o antigo é revogado e um novo par é emitido (`RefreshToken()` em
  `auth_service.go`).
- `IsActive()` mata o token **na hora** se `revoked_at` ou `used_at` != nil — **sem janela de graça**.
- Client (`api-client.ts`): em 401 chama `/auth/refresh`; sucesso → `setAuth` (novo par);
  **falha → `clearAuth` (logout)**. Há guard `isRefreshing` só dentro de uma instância.
- 2FA TOTP existe. **Não existe "lembrar deste aparelho"** no web. Mobile tem `/me/sessions`
  (device_tokens); web não tem tela de sessões.

## Por que o logout é rápido no PWA do iPhone (diagnóstico)

Causa principal = **rotação de refresh de uso único + iOS suspendendo/matando o PWA**:

1. O iOS **congela/mata o PWA** ao ir pro background. Ao reabrir, o `refetchOnWindowFocus` do
   TanStack dispara uma **rajada** de requests. Se o access (15min) expirou, várias batem 401 ao
   mesmo tempo.
2. O refresh **rotaciona e revoga o token antigo imediatamente**. Se o app é morto na janela
   entre receber o novo refresh e o Zustand **persistir** no `localStorage`, ou se uma request
   atrasada usa o token já rotacionado, o próximo uso bate em token revogado → `ErrTokenRevoked`
   → **logout**. O access de 15min faz essa dança acontecer o tempo todo, multiplicando a chance.
3. **Teto do iOS/WebKit:** o WebKit **expira storage script-writable (localStorage) após ~7 dias**
   (ITP). PWA instalado relaxa isso, mas não garante. Como o refresh também é 7 dias, o teto
   prático é ~7 dias de qualquer forma.

Ou seja: o "rápido" (minutos/horas) é a corrida de rotação; o teto (7 dias) é storage + expiry.

## Melhor forma de resolver (defesa em camadas)

### A. Parar o logout prematuro (maior impacto, fazer primeiro)
- **Janela de graça na rotação:** ao rotacionar, aceitar o token antigo por ~30–60s (ou aceitar o
  novo enquanto o antigo "rotacionado" ainda vale brevemente). Detecção de **reuse de família**
  (token chain): reuso real depois da graça = suspeita de roubo → revoga a família inteira. É o
  padrão de mercado pra refresh rotation sem derrubar sessão legítima.
- **Client single-flight + persistência síncrona:** todas as 401 aguardam UM refresh; só seguir
  depois que o novo par foi gravado no storage.

### B. Sessão longa + "lembrar deste aparelho"
- Opt-in **"manter conectado neste aparelho"** no login. Em aparelho confiável:
  - refresh **longo (ex.: 30 dias)**, vinculado ao device (já temos `user_agent`/`ip_address` em
    `refresh_tokens`; somar um `device_label`/`device_id`).
  - **expiração deslizante:** cada refresh estende a validade (janela rolante) até um **teto
    absoluto** (ex.: 30–90 dias). Aparelho em uso fica logado; aparelho ocioso expira.
- Sem o opt-in (computador compartilhado da clínica): sessão curta (ex.: 1 dia ou só a sessão).
- Pode subir o access pra **30–60min** pra reduzir a frequência de refresh (menos corrida).

### C. Durabilidade de storage no iOS
- Chamar **`navigator.storage.persist()`** (PWA instalado reduz evicção).
- **Melhor opção (segurança+durabilidade):** mover o **refresh token pra cookie
  `HttpOnly; Secure; SameSite=None` no domínio `.plenyasaude.com.br`** (api e app são subdomínios).
  Não é legível por JS (imune a XSS) e é mais durável no PWA do que localStorage. Access fica em
  memória. É mudança arquitetural maior, porém é a prática recomendada pra refresh tokens.

## O que a lei diz (LGPD + CFM/SBIS)

- **LGPD (Lei 13.709/2018):** não fixa timeout. Exige **segurança** (art. 46) e trata **dado de
  saúde como sensível** (art. 11), proteção reforçada. "Manter conectado" é permitido desde que
  haja medidas de segurança proporcionais.
- **Prontuário eletrônico — CFM Res. 1.821/2007 + Manual SBIS/CFM de Certificação (S-RES):** os
  níveis de garantia de segurança (NGS) exigem, entre outros, **encerramento/bloqueio automático
  de sessão por inatividade** e **reautenticação**, além de **autenticação forte** e **auditoria**.
- **Leitura prática (como EMRs conformes fazem):** separa-se **"confiança do dispositivo"** de
  **"sessão ativa de dados clínicos"**:
  - O **credencial de longa duração (refresh)** pode durar dias/semanas no aparelho pessoal,
    revogável (opt-in, transparente, só em aparelho pessoal — não compartilhado).
  - MAS a **visualização de dados de paciente** deve ter **bloqueio por inatividade** (após N min,
    trava a tela e pede reautenticação leve: biometria/PIN/senha, e 2FA periódico).
  - Mantém **2FA** (já temos), **auditoria** (já temos) e **revogação remota de sessões**
    ("Minhas sessões" no web, espelhando o que o mobile já tem em `/me/sessions`).

**Conclusão:** o objetivo não é "logar uma vez pra sempre", e sim **manter o aparelho confiável
por dias (refresh longo, deslizante, revogável, opt-in)** + **bloqueio por inatividade com
reautenticação** sobre os dados clínicos. Isso resolve a queixa, é seguro e é defensável perante
LGPD/CFM/SBIS.

## Proposta de faseamento (a discutir, NÃO implementado)
1. **Fix da corrida (A):** janela de graça na rotação + single-flight/persist no client +
   `storage.persist()`. Sozinho já deve acabar com o "logout rápido". Baixo risco.
2. **Remember-me + deslizante (B):** opt-in no login, refresh 30d deslizante p/ device confiável,
   tela "Minhas sessões" no web (revogação).
3. **Bloqueio por inatividade (compliance):** lock de tela + reautenticação (biometria/PIN) sobre
   dados de paciente; define o timeout configurável.
4. **(Opcional, maior) refresh em cookie HttpOnly** no domínio raiz — segurança e durabilidade.

Pendências/decisões pro Getúlio: duração do refresh confiável (30/60/90d?), timeout de inatividade
(ex.: 15–30 min?), e se entra o cookie HttpOnly agora ou depois.

---

## 🟢 Status: IMPLEMENTADO no dev (2026-06-09), aguardando ordem de deploy
- Backend: migration 00036 (família/rotated_at/remember/last_used_at), rotação com janela de
  graça (60s) + detecção de roubo por família, remember 30d deslizante, `/auth/sessions`
  (GET/DELETE), `jti` único nos tokens (corrige colisão de hash), access 15→30min.
- Frontend: checkbox "Manter conectado" no login, `navigator.storage.persist()` no boot,
  card "Aparelhos conectados" no /profile.
- **Validado no dev (E2E):** login remember → refresh encadeia; reuso na graça → 200 (corrida
  do PWA não desloga); reuso fora da graça → 401 + família revogada; remember=30d / sessão=7d.
- **Prod (no deploy):** migration roda no entrypoint; envs novas têm default no código
  (`JWT_REMEMBER_EXPIRY=720h`, `JWT_REFRESH_GRACE_SECONDS=60`). Se quiser access 30min em prod,
  ajustar `JWT_ACCESS_EXPIRY` no Coolify (hoje pode estar 15m).

## ✅ Decisões (2026-06-09)
- **Escopo:** Fix da corrida + "manter conectado". (Bloqueio por inatividade fica pra fase futura.)
- **Duração do aparelho confiável:** **30 dias deslizante**.
- **Bloqueio por inatividade:** **não agora** (registrado como débito de conformidade CFM/SBIS).
- Cookie HttpOnly: **depois** (mantém localStorage por ora; só adiciona `storage.persist()`).

## Plano de implementação (escopo aprovado) — dev primeiro, sem deploy

### Backend (`apps/api`)
1. **Janela de graça na rotação** (`auth_service.RefreshToken`): ao rotacionar, em vez de matar o
   antigo na hora, marcar `rotated_at`/`replaced_by` e **aceitar o antigo por ~60s** após a
   rotação. Reuso do antigo DEPOIS da graça (e com o novo já usado) = suspeita → revoga a
   **família** inteira (`family_id`). Migration: somar a `refresh_tokens` os campos
   `family_id uuid`, `rotated_at timestamptz`, `replaced_by uuid` (e índice em `family_id`).
2. **Remember-me + deslizante:** `POST /auth/login` aceita `rememberDevice bool`. Refresh nasce
   com expiry **30d** (confiável) ou **curto** (ex.: 1d / `JWT_REFRESH_EXPIRY` atual) caso
   contrário; o tipo fica gravado no token/family. Cada `/auth/refresh` **renova a expiry**
   (deslizante) até um **teto absoluto** (ex.: 30d desde o login; configurável
   `JWT_REMEMBER_EXPIRY=720h`, `JWT_REMEMBER_ABSOLUTE`).
3. **Minhas sessões:** `GET /me/sessions` (web) lista os `refresh_tokens` ativos do usuário
   (device/user_agent/ip/última atividade) e `DELETE /me/sessions/:id` revoga. (Mobile já tem
   algo análogo via device_tokens; aqui é pros refresh tokens do web.)
4. Subir o access token pra **30min** (reduz frequência de refresh/corrida). Env, sem migration.

### Frontend (`apps/web`)
5. **Single-flight + persistência síncrona** no `api-client`: garantir que todas as 401 aguardam
   UM refresh e que o novo par é gravado no storage antes de seguir (já há `isRefreshing`; reforçar
   e tratar o caso de cold-start do PWA).
6. **Checkbox "Manter conectado neste aparelho"** no login → manda `rememberDevice` (aviso: só em
   aparelho pessoal, não compartilhado).
7. `navigator.storage.persist()` no boot do app (reduz evicção no PWA iOS).
8. Tela **"Minhas sessões"** (perfil/segurança): listar e revogar aparelhos.

### Bloqueio por inatividade — IMPLEMENTADO (4h, 2026-06-09)
- Decisão atualizada: **trava após 4h de inatividade** (CFM/SBIS).
- Frontend `InactivityLock` (envolve o layout autenticado): rastreia atividade (mouse/teclado/
  toque/scroll, throttle 15s) em `localStorage`; após 4h sem atividade (ou ao reabrir o app
  depois desse tempo) mostra overlay com `backdrop-blur` (esconde dados do paciente) exigindo
  **senha** pra desbloquear. Reautenticação leve: `POST /auth/verify-password` (confere senha,
  NÃO reemite tokens — a sessão continua). Fallback "Sair e entrar com outra conta" (cobre
  contas OAuth sem senha). Login reseta o relógio (`touchActivity`).
- Limite hardcoded `INACTIVITY_MS = 4h` em `components/auth/inactivity-lock.tsx`.

### Fora deste escopo (débito registrado)
- Refresh token em cookie HttpOnly no domínio raiz — melhoria de segurança/durabilidade futura.
- Biometria/PIN no desbloqueio (hoje é senha) — possível evolução.
