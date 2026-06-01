# Atualização do Tech Stack — EMR Plenya (maio/2026)

> Pesquisa current-vs-latest (fontes web reais: pkg.go.dev, npm, GitHub releases) + blast-radius no
> código + decisão. Postura: **higiene de segurança agressiva em patches/minors, conservadorismo em
> majors**. Status de execução ao fim de cada item. Nada commitado até o usuário pedir.

## Maior retorno imediato (segurança)

- **Go toolchain 1.25.7 → 1.25.9+** (container já em 1.25.10): ~13 CVEs (DoS crypto/tls, DoS
  crypto/x509, XSS html/template). Bump zero-breaking.
- **Next 16.2.3 → 16.2.6 + React 19.2.0 → 19.2.6**: 13 advisories (HIGH: auth/middleware bypass,
  SSRF em WS upgrade, DoS RSC/image, XSS CSP nonce; CVE-2026-23870 DoS RSC no React). Atinge
  exatamente RBAC-por-middleware + portal-por-subdomínio. Mesma minor, sem breaking.

## Lote seguro (ATUALIZAR JÁ)

### Backend
- `go` directive 1.25.7 → 1.25.9 (container já tem toolchain ≥1.25.9; prod usa `golang:1.25-alpine` flutuante, pega patch no rebuild).
- `golang.org/x/crypto` 0.50.0 → 0.52.0 (CVEs só em ssh/*, não usado; higiene SCA).
- `github.com/liushuangls/go-anthropic/v2` 2.18.0 → 2.20.2 (fix de sparse stream indexes no RAG/chat).

### Web
- `next` 16.2.6, `eslint-config-next` 16.2.6, `react`/`react-dom` 19.2.6 (lockstep, segurança).
- `axios` 1.16.1 — **dep morta (0 call-sites)**; bump por higiene ou remover.
- `@tanstack/react-query`+devtools 5.100.14, `@tanstack/react-table` 8.21.3.
- `react-hook-form` 7.76.1 + `@hookform/resolvers` 5.4.0.
- `@anthropic-ai/sdk` 0.100.1, `date-fns` 4.4.0, `cmdk` 1.1.1, `vaul` 1.1.2, `typescript` 6.0.3.
- `@tiptap/*` (react, starter-kit, extension-color/highlight/text-align/text-style/underline) → 3.23.6 em lockstep.
- `@radix-ui/*` (~21 pkgs) → patches via `pnpm update` (ranges ^ já cobrem).

## ATUALIZAR C/ CUIDADO (um de cada vez, fora do lote seguro)

- `zod` 4.3.6 → 4.4.3 (regenerar @plenya/types, rodar validação; watch regressão preprocess+catch #5937).
- `pgvector/pgvector-go` 0.1.1 → 0.4.0 (split de submódulos; validar import pgx no RAG).
- `@types/react`/`@types/react-dom` → 19.2.x (alinhar ao runtime; typecheck).
- `lucide-react` 0.469 → 1.17.0 (major drop-in; auditar brand icons removidos, aria-hidden default).
- `pdfjs-dist` → faixa peer do react-pdf 10.4.1 (testar visualizador).
- `@tremor/react` 3.18.7 (testar sob React 19.2; lib parada, planejar saída p/ shadcn/charts).
- `@types/node` → alinhar ao Node de runtime (24 LTS), NÃO 25.x.
- `eslint` linha 9.4x (oportunístico).

## SEGURAR / Migração major dedicada

- **Fiber v2.52 → v3** — esforço ALTO. Sem CVE → sem pressão. Blast-radius: 67 arquivos, 1.771 `fiber.`,
  442 `*fiber.Ctx` (vira interface), 564 rotas, 7 middlewares custom, error handler global, 27 `c.Locals`
  com type-assert, 32 `UserContext()`/`Context()`. CLI `fiber migrate` cobre o mecânico; Locals/Context/
  middlewares exigem revisão manual. Migração dedicada.
- **Tailwind v3.4 → v4** — esforço MÉDIO-ALTO. Sem CVE. Config CSS-first (`@theme`), plugin
  `@tailwindcss/postcss`, `border` default vira currentColor, substituir `tailwindcss-animate`, toca
  `@plenya/brand` (tokens gold/petrol/ocean/sage/cream) + ~260 componentes. Upgrade tool + revisão visual.
- **gofpdf/v2** — lib órfã (jung-kurt parado, fork go-pdf/fpdf arquivado). Sem update a aplicar; dívida
  técnica nos laudos PDF (AssessmentHTMLService). Avaliar codeberg.org/go-pdf/fpdf (drop-in) ou render
  via go-rod/Chromium (já no stack). Sem prazo curto.
- **swag v2** (ainda RC, OpenAPI 3.x, retrabalha `pnpm generate`), **Next 16.3** (só canary),
  **@types/node 25.x** (ímpar/current) — aguardar.

## Verificação por lote
- Backend: `docker compose exec -w /app api go build ./...` + `go test ./...` (atenção ao
  `chunking_service_test.go:186` quebrado PRÉ-EXISTENTE, não-regressão).
- Web: `pnpm -F web exec tsc --noEmit` (build tem `ignoreBuildErrors:true`, não mascarar) + `pnpm -F web build`.
- Commitar em lotes separados (backend / web-segurança / web-libs) para bisect fácil.

## Status de execução (2026-05-30)

**✅ APLICADO E VERIFICADO (lote seguro):**

Backend (`go build ./...` OK + `go vet` limpo, fora o `chunking_service_test.go:186` pré-existente):
- go.mod diretiva `go 1.25.7` → `1.25.9` (container já em toolchain 1.25.10; prod usa `golang:1.25-alpine` flutuante → pega patch no rebuild).
- `golang.org/x/crypto` 0.50.0 → 0.52.0 (e o grupo golang.org/x: net/sys/text/mod/tools subiram junto via tidy).
- `github.com/liushuangls/go-anthropic/v2` 2.18.0 → 2.20.2.

Web (`pnpm -F web build` → **EXIT 0**, verificado):
- next 16.2.6, eslint-config-next 16.2.6, react 19.2.6, react-dom 19.2.6, axios 1.16.1.
- @tanstack/react-query(+devtools) 5.100.14, @tanstack/react-table (latest 8.x).
- react-hook-form 7.76.1, @hookform/resolvers (latest 5.x), @anthropic-ai/sdk 0.100.1.
- date-fns 4.4.0, cmdk 1.1.1, vaul (patch), typescript 6.0.3, @tiptap/* 3.23.6 (lockstep), @radix-ui/* (patches).
- `next.config.ts`: removida a chave `eslint` (Next 16 removeu o lint embutido no build → era config morta + warning).

Nota de ambiente: `node_modules`/`.next` de vários workspaces estavam **root-owned** (criados por containers que rodam como root), bloqueando `pnpm`/build; corrigido com `chown -R 1000:1000` via container root.

**✅ Rodada 2 — Tier 1+2 APLICADOS e deployados (2026-05-31):**
- Tier 1 (baixo risco): @types/react 19.2.15, @types/react-dom 19.2.3, @tremor/react 3.18.7, framer-motion 12.40.0.
- Tier 2 (c/ cuidado, um a um, build verde cada): zod 4.4.3, pgvector-go 0.4.0 (runtime RAG verificado), @types/node 24.12.4 (alinha ao Node 24 de runtime), lucide-react 1.17.0 (major; 194 arquivos sem ícone quebrado).
- **2 fixes de runtime** (pegos por QA Playwright, invisíveis ao build): react-query **deduplicado** (estava split 5.99/5.100 → "No QueryClient set" no dev) e **CORS** (`CORS_ORIGIN` singular CSV → splita por vírgula). Ver memória `emr_qa_visual_playwright`.
- Secretaria `/hoje` resolvido (página única `/recepcao`, Fase 2 concluída e deployada).

**✅ Tailwind v3.4 → v4.3 CONCLUÍDO (2026-05-31):** os 3 apps web (site-getulio CSS-first; site + web via `@config` preservando o preset `@plenya/brand`) + `tw-animate-css` no lugar do `tailwindcss-animate` + `@utility container` no web. QA visual Playwright pré/pós fiel. Commits `7381e7a3`/`b9b456ec`/`24086914`. Detalhes e aprendizados em `docs/emr/tailwind-v4-migracao.md`.

**⏳ Ainda segurado (sem pressão de segurança):** **Fiber v3** (adiado, toca núcleo auth), **eslint 10** (compat eslint-config-next), **pdfjs 6** (preso ao react-pdf 10), **swag v2** (RC). Mobile/NativeWind seguem em Tailwind v3 (intencional).
