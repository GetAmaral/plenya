# Migração Tailwind v3.4 → v4.3 — EMR + sites (maio/2026)

> Último major segurado da modernização do EMR (ver `docs/emr/atualizacao-versoes.md` e memória
> `emr_modernizacao_status`). Sem CVE — risco é **regressão visual**, não segurança. Por isso o
> gate é **QA visual Playwright pré/pós** (memória `emr_qa_visual_playwright`), não só build verde.
> Trabalho direto no `master` (regra de ouro: nunca branch). Commit + push por app.

## Escopo

3 apps web em Tailwind `^3.4.18` + o preset compartilhado `@plenya/brand`:

| App | Porta dev | Preset @plenya/brand | tailwindcss-animate | darkMode | Particularidades |
|-----|-----------|----------------------|---------------------|----------|------------------|
| `apps/web` | 3000 (compose, bypass auth) | **sim** (`presets:[plenyaPreset]`) | **sim** | `class` | shadcn HSL `hsl(var(--x))`, `theme.container{center,padding,screens}`, keyframes accordion/fade/slide/scale no config, `backdropBlur.xs`, `@apply bg-gradient-to-*` no globals |
| `apps/site` | 3002 | **sim** | **sim** | — | `@apply` pesado em `@layer components/utilities`, `bg-gold-600` |
| `apps/site-getulio` | 3003 | **não** (paleta própria paper/ink/gold/navy) | não | — | `font-family: theme('fontFamily.sans')` em CSS, `autoprefixer` |

Fora de escopo: `apps/mobile-*` usam **NativeWind** (sintaxe Tailwind v3, versionamento próprio) — NÃO migrar.

Alvos npm confirmados: `tailwindcss@4.3.0`, `@tailwindcss/postcss@4.3.0`, `tw-animate-css@1.4.0`.

## Estratégia (decisão)

**Caminho conservador: manter os JS configs + o preset `@plenya/brand` via `@config`**, mudando
só o necessário no entrypoint de cada app. Minimiza blast-radius e mantém os 3 apps independentes
(não preciso reescrever o preset em `@theme`). Por app:

1. PostCSS: `tailwindcss: {}` → `@tailwindcss/postcss: {}`; remover `autoprefixer` (v4 já prefixa via Lightning CSS).
2. `globals.css`: `@tailwind base/components/utilities` → `@import "tailwindcss";` + `@config "../../tailwind.config.ts";` (mantém o JS config/preset).
3. Deixar o codemod `npx @tailwindcss/upgrade` renomear utilitários (shadow-sm→shadow-xs, rounded→rounded-sm, bg-gradient→bg-linear, outline-none→outline-hidden, blur→blur-sm, etc.) em templates **e** no globals.css, e converter `@layer utilities {.x}` → `@utility x`.
4. Shims anti-regressão de Preflight v4 (border currentColor→cor antiga, ring 1px/3px, cursor de button) conforme o codemod injetar / o QA exigir.
5. `tailwindcss-animate` (web, site) → `tw-animate-css` (import CSS); classes shadcn (`animate-in`, `fade-in-0`, `zoom-in-95`, `slide-in-from-*`, `data-[state]`) são compatíveis.
6. Gotchas específicos: `theme.container` (web) e `theme('...')` (getulio) podem não sobreviver ao compat `@config` — tratar caso a caso (validado no QA).

> **Decisão CSS-first `@theme` vs `@config`:** ficar no `@config` por ora. Migrar o preset para
> `@theme` é um segundo passo opcional, depois que os 3 apps estiverem verdes e validados.

## Ordem de execução

Por app, isolado, com commit+push cada (bisect fácil). Cada app: **baseline screenshot → migrar →
build → screenshot pós → comparar → corrigir → commit+push**.

1. `apps/site-getulio` primeiro (mais simples: sem preset, sem animate) — valida o fluxo.
2. `apps/site` (usa preset + animate, mas menos superfície que o web).
3. `apps/web` (maior superfície: shadcn + container + dark mode + EMR inteiro).

Racional: subir a complexidade gradualmente; o getulio serve de ensaio do codemod.

## QA visual (gate)

Harness: `/tmp/tw4-shot.js` (Playwright 1.60 local, Chromium cached). Baselines `before` já
capturados em `/tmp/tw4-qa/{web,site,getulio}/before__*.png` (12/10/8 rotas, status 200).
Dev: web :3000 (bypass), site :3002, getulio :3003 (servers de dev no host).

Fluxo pós-migração de cada app: rodar o mesmo conjunto de rotas com label `after`, ler os PNGs,
comparar com `before`. Build (`pnpm -F <app> build`) tem que ficar verde **e** o visual idêntico.

## Rollback

Cada app é 1 commit isolado no master. Regressão detectada no QA pós-commit → `git revert <sha>`
do app afetado (não afeta os outros). Antes do commit, regressão → desfazer no working tree.

## Status

- [x] Recon + baselines (workflow `tw4-recon` + Playwright). Configs/CSS mapeados.
- [ ] site-getulio
- [ ] site
- [ ] web
- [ ] docs/memória atualizados
