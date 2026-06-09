# Estudo — prevenção de XSS no modelo atual (apps/web)

**Contexto (2026-06-09):** mantendo o refresh token em `localStorage` (sem cookie HttpOnly por
ora), o melhor retorno de segurança vem de **prevenir XSS** — porque um XSS é o que permitiria ler
o token e/ou agir na sessão. Estudo das melhores práticas aplicadas ao nosso stack
(Next 16 + React 19 + TanStack + Tiptap + DOMPurify).

## Onde já estamos bem (não mexer, só manter)
- **React escapa JSX por padrão** — interpolação de texto (`{valor}`) é segura. É a razão de a
  maior parte do app já estar imune.
- **Todos os `dangerouslySetInnerHTML` (5) passam por `DOMPurify.sanitize`** — workout plan
  (público + interno), anamnese, workspace da consulta. (`isomorphic-dompurify` instalado.)
- **CSP + headers de segurança** já existem (`next.config.ts`): `default-src 'self'`,
  `frame-ancestors 'none'`, `object-src 'none'`, `base-uri 'self'`, `form-action 'self'`,
  `X-Frame-Options: DENY`, `X-Content-Type-Options: nosniff`, HSTS, Referrer-Policy,
  Permissions-Policy.
- **Sem `eval`/`new Function` no nosso código.**

## Gaps mapeados (com evidência)
1. **CSP `script-src` permissiva:** `script-src 'self' 'unsafe-inline' 'unsafe-eval' …`. O
   `unsafe-inline` praticamente **anula** a proteção anti-XSS do CSP (um `<script>` injetado
   executaria). O próprio comentário no arquivo reconhece. **É o maior gap.**
2. **URLs vindas de dado em `href`/`src`** (vetor `javascript:`): `campaign.url`/`c.url`
   (campaigns) e `href={url}` no `conversation-viewer` (links de mensagens de WhatsApp/email =
   **não confiável**). React **não** bloqueia `javascript:` em href — só avisa em dev. EpubViewer
   usa blob (download), ok.
3. **DOMPurify descentralizado:** cada sink chama `DOMPurify.sanitize(...)` direto, sem config
   única nem trava — um sink novo sem sanitize passaria despercebido (não há lint
   `react/no-danger`).
4. **`components/ui/markdown-content.tsx`** — verificar se não usa `rehype-raw` (que reabre HTML
   cru); react-markdown é seguro por padrão, mas precisa confirmar.
5. **`unsafe-eval` sem uso real** — está no CSP mas nada nosso usa; candidato a remoção.

## Melhores práticas priorizadas

### P0 — maior impacto
- **CSP estrita com nonce por request.** Remover `'unsafe-inline'` (e `'unsafe-eval'`) do
  `script-src`, usando **nonce por requisição** (Next 16 suporta via middleware: gera nonce,
  injeta no header CSP e o Next o propaga aos `<script>`) + `'strict-dynamic'`. Mantém
  `style-src 'unsafe-inline'` (estilos inline são mais difíceis de remover e o risco é bem menor).
  Fecha o vetor de script inline injetado — a defesa-em-profundidade que falta.
  Esforço: médio (testar telemed/Google/Daily, recharts). Risco: médio (validar que nada quebra).
- **Sanitizar URLs** antes de `href`/`src` que vêm de dado: helper `safeUrl()` que só aceita
  `http:`/`https:`/`mailto:`/`tel:` e descarta `javascript:`/`data:`. Aplicar em campaigns e no
  conversation-viewer. Esforço: baixo. Risco: baixo.

### P1 — consolidação
- **Centralizar a sanitização:** um único `sanitizeHtml()` (config estrita: sem `script`,
  sem handlers `on*`, sem `style` se possível) e/ou componente `<SafeHtml html=…/>`. Trocar os 5
  sinks pra usar ele. Adicionar regra de lint **`react/no-danger`** (erro) pra impedir sink novo
  sem o wrapper. Garante que não nasce buraco no futuro.
- **Confirmar `markdown-content.tsx`** sem `rehype-raw`; se usar, sanitizar a saída.
- **Supply chain (XSS via dependência):** `pnpm audit` no CI + Dependabot/renovate;
  `frozen-lockfile` (já temos); minimizar scripts de terceiros; SRI nos externos (Google/Daily)
  quando aplicável. Com CSP estrita, uma dep comprometida não consegue injetar inline script.

### P2 — avançado (opcional)
- **Trusted Types** (`require-trusted-types-for 'script'` no CSP): força todo sink de DOM a passar
  por uma política; DOMPurify sabe retornar `TrustedHTML`. Defesa forte contra DOM-XSS, mas exige
  ajustes e é só Chromium (degrada gracioso em Safari).
- **Reduzir blast radius do token:** access já é curto (30 min). Avaliar manter o access só em
  memória (não persistir) — o refresh segue sendo o alvo (aí entra o debate do cookie HttpOnly,
  em `estudo-sessao-login-persistente.md`).

## Recomendação
Ordem de melhor custo-benefício, **independente do cookie HttpOnly**:
1. **P0 sanitizar URLs** (rápido, fecha um vetor concreto que existe hoje).
2. **P0 CSP estrita com nonce** (a grande melhora; remove `unsafe-inline`/`unsafe-eval`).
3. **P1 centralizar sanitização + lint `no-danger`** (trava regressão futura).
4. P1 supply chain no CI · P2 conforme apetite.

Isso eleva bastante a postura de XSS **mantendo o modelo atual** (token em localStorage), e é
pré-requisito natural caso um dia se faça o cookie HttpOnly.

## Status — P0 + P1 IMPLEMENTADOS no dev (2026-06-09), aguardando deploy
- [x] **P0 safeUrl** (`lib/security.ts`): bloqueia `javascript:/data:/vbscript:/file:`; aplicado em
  campaigns (2) e conversation-viewer (img/áudio/link/download).
- [x] **P0 CSP estrita com nonce** (`middleware.ts`): `script-src` com `nonce` por request +
  `strict-dynamic`, **sem `unsafe-eval`** (só em dev, p/ HMR). Verificado: header presente, scripts
  do Next recebem o nonce, páginas carregam 200. CSP saiu do `next.config.ts` (lá ficaram os
  demais headers).
- [x] **P1 centralização**: `sanitizeHtml()` único + componente `<SafeHtml>` (único sink
  autorizado); os 5 `dangerouslySetInnerHTML` migrados (treino interno/público, anamnese×2,
  workspace da consulta). Sobrou `dangerouslySetInnerHTML` só no `SafeHtml`.
- [x] **P1 lint** `react/no-danger: error` no `eslint.config.mjs`.
- [x] **P1 markdown**: `markdown-content.tsx` usa react-markdown sem `rehype-raw` → seguro.
- ⚠️ **Tooling:** `next lint` foi removido no Next 16 (o script `"lint"` quebra sozinho). A regra
  `react/no-danger` é válida mas só atua quando o lint rodar — corrigir o runner de lint é tarefa
  à parte (migrar pra `eslint .` com flat config funcional).

### Pendente (P2, fora deste pacote)
- Trusted Types · access token só em memória · `pnpm audit` no CI + Dependabot · cookie HttpOnly.
