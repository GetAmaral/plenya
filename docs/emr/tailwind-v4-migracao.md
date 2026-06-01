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

## Status — CONCLUÍDO (2026-05-31)

- [x] Recon + baselines (workflow `tw4-recon` + Playwright). Configs/CSS mapeados.
- [x] **site-getulio** — commit `7381e7a3`. Codemod escolheu **CSS-first** (deletou `tailwind.config.ts`, palette → `@theme`). Fix: tokens de fonte em `@theme inline` (evita `--font-*` circular vs next/font).
- [x] **site** — commit `b9b456ec`. Codemod manteve `@config` (preserva preset). Removido `tailwindcss-animate` (0 utils usados).
- [x] **web** — commit `24086914`. `@config` + shadcn `hsl(var())` ok. globals.css migrado À MÃO (codemod não resolve `@apply` de utils do @config). `tailwindcss-animate` → `tw-animate-css` (18 usos shadcn). `@utility container` (v4 dropou `theme.container`).
- [x] `@plenya/brand` tailwindcss dep → ^4.3 (só tipo `Config`; preset segue JS via @config).
- [x] docs/memória atualizados.

## Resultado do QA visual (Playwright pré/pós, dev)

Fiel em todos os apps. `web` o mais limpo (<0.21% por rota, `/recepcao` pixel-idêntico) graças
ao `* { @apply border-border }` global. Sites com drift vertical sub-perceptual (~1%, direções
mistas por página) = reflow de métrica sub-pixel do pipeline Lightning CSS do v4; sem elemento
quebrado (heatmaps mostram ghosting cumulativo, não bordas/cores sumindo). Marca, gradientes
(`bg-linear`), bordas (compat shim / border-border), shadcn e charts intactos.

## Aprendizados (ler antes de repetir o codemod em monorepo pnpm hoisted)

1. **Permissões:** `node_modules` root-owned (containers rodando como root) fazem o `pnpm` do
   codemod abortar com `ERR_PNPM_EACCES`. Corrigir antes: `sudo chown -R 1000:1000 node_modules packages apps`.
2. **Corrupção de prosa:** o codemod faz find-replace de classes em TEMPLATES *e* em conteúdo;
   corrompeu a palavra inglesa "ring" → "ring-3" num excerpt MDX do getulio. **Auditar SEMPRE**
   todo diff de `.tsx`/`.mdx`/strings por tokens (`ring-3`, `shadow-xs`, etc.) fora de `className`.
3. **`@apply` de utils do @config:** o renderer isolado do codemod não transpila o `.ts` config,
   então engasga em `@apply border-border` (web). Solução: stub mínimo no globals durante o
   codemod (só p/ ele migrar os TEMPLATES), depois migrar o globals à mão.
4. **`npm` vs `pnpm`:** o codemod às vezes chama `npm add` e quebra no `workspace:*`. Setar as
   deps à mão (tailwindcss ^4.3 + @tailwindcss/postcss, remover autoprefixer) é mais confiável.
5. **animate:** `tailwindcss-animate` (plugin v3) quebra no v4 → `tw-animate-css` (import CSS, classes shadcn compatíveis). O codemod NÃO faz essa troca.
6. **Dev container:** após o bump, recriar o `web` com `docker compose up -d --force-recreate --renew-anon-volumes web` (o entrypoint roda `pnpm install`); o node_modules do container é volume anônimo e fica stale.
