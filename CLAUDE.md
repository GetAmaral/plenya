# Plenya — Monorepo

Plenya é uma plataforma de saúde (EMR + sites + conteúdo + social) construída como monorepo
pnpm + Turborepo. Este arquivo é o **índice raiz** e contém as **regras invariantes
compartilhadas** por todos os subprojetos. Cada app tem seu próprio `CLAUDE.md` com detalhes
locais — o Claude carrega automaticamente o mais próximo do diretório em que você está
trabalhando.

> **Como usar este monorepo com o Claude:** trabalhe a partir do diretório do subprojeto
> (`apps/api`, `apps/web`, etc.) para que o `CLAUDE.md` local entre em contexto. Para tarefas
> paralelas e isoladas, use git worktrees (`git worktree add ../plenya-<x> <branch>`). As
> regras deste arquivo valem em qualquer subprojeto; os `CLAUDE.md` locais nunca as duplicam,
> só referenciam.

> 🧭 **Roteamento de subprojeto (instrução ao Claude):** quando o usuário disser que vai
> trabalhar num subprojeto — em linguagem natural ("vamos trabalhar no site plenya", "mexer no
> backend", "no app do paciente") ou via slash command (`/site`, `/api`, ...) — **leia
> imediatamente o `CLAUDE.md` do app correspondente**, confirme em uma linha que carregou o
> contexto, e só então prossiga. Mapa de frases → arquivo:
>
> | O usuário menciona… | Carregue |
> |---|---|
> | site plenya, institucional, plenyasaude | `apps/site/CLAUDE.md` (cmd `/site`) |
> | site do Getúlio, pessoal, drgetulioamaralfilho | `apps/site-getulio/CLAUDE.md` (`/site-getulio`) |
> | backend, api, Go, models, escore/banco | `apps/api/CLAUDE.md` (`/api`) |
> | web, EMR, frontend, portal do paciente | `apps/web/CLAUDE.md` (`/web`) |
> | app profissional, Plenya Pro, mobile-pro | `apps/mobile-pro/CLAUDE.md` (`/mobile-pro`) |
> | app do paciente, Plenya app, mobile-app | `apps/mobile-app/CLAUDE.md` (`/mobile-app`) |
> | social, instagram, linkedin, MCP, DMs | `apps/social-mcp/CLAUDE.md` + `.claude/social/` (`/social`) |
> | deck, eBook, blog, imagem | `.claude/content/*.md` |

## 🗺 Mapa dos subprojetos

### `apps/` — aplicações

| App | O que é | Stack | CLAUDE.md |
|-----|---------|-------|-----------|
| **api** | Backend EMR — fonte única dos dados (74 models, RAG/pgvector, CRM, portal, telemed, calendar, training, subscriptions) | Go 1.25 + Fiber v2 + GORM v1.31 + PostgreSQL 18 (pgvector) | [apps/api/CLAUDE.md](apps/api/CLAUDE.md) |
| **web** | Frontend EMR profissional **+ portal do paciente** (mesmo app, roteado por subdomínio via middleware) | Next 16.2 + React 19.2 + TanStack Query v5 + Tailwind v4 + shadcn/ui | [apps/web/CLAUDE.md](apps/web/CLAUDE.md) |
| **site** | plenyasaude.com.br — site institucional (blog MDX, escore-light, leads→CRM) | Next 16.2 + next-intl + MDX | [apps/site/CLAUDE.md](apps/site/CLAUDE.md) |
| **site-getulio** | drgetulioamaralfilho.com.br — site pessoal do Dr. Getúlio (livros, palestras, artigos) | Next 16.2 + next-intl + MDX | [apps/site-getulio/CLAUDE.md](apps/site-getulio/CLAUDE.md) |
| **social-mcp** | Servidor MCP (Python) que conecta o Claude ao Instagram/Facebook via Meta Graph API | Python 3.11 + MCP + httpx | [apps/social-mcp/CLAUDE.md](apps/social-mcp/CLAUDE.md) |
| **mobile-pro** | App profissional (Plenya Pro) — `com.plenya.pro` | Expo SDK 52 + RN 0.76.5 | [apps/mobile-pro/CLAUDE.md](apps/mobile-pro/CLAUDE.md) |
| **mobile-app** | App do paciente (Plenya) — `com.plenya.app` | Expo SDK 52 + RN 0.76.5 | [apps/mobile-app/CLAUDE.md](apps/mobile-app/CLAUDE.md) |

### `packages/` — bibliotecas compartilhadas

| Package | O que é | Gerado? |
|---------|---------|---------|
| **@plenya/types** | Tipos TS + schemas Zod a partir do OpenAPI do backend | ⚙️ 90% gerado |
| **@plenya/api-client** | Cliente HTTP tipado + query options do TanStack Query | ✍️ manual |
| **@plenya/brand** | Tokens de marca (paleta gold/petrol/ocean/sage/cream), tipografia, logos | ✍️ manual |
| **@plenya/domain** | Lógica de domínio pura (cálculo de escore, CPF, formatters) | ✍️ manual |
| **@plenya/emails** | Templates React Email → HTML estático | ✍️ manual |
| **@plenya/ui-mobile** | Primitivos React Native + NativeWind | ✍️ manual |
| `packages/ui` | Reservado/vazio (sem package.json) | — |

### Conteúdo, scripts e skills

- **`scripts/deck-builder/`** — decks comerciais (HTML/CSS → PDF via Playwright). Doc viva: `continuum/EDITORIAL.md`. Ver [.claude/content/decks.md](.claude/content/decks.md).
- **`scripts/blog-generator/`** — geração de imagens com **gpt-image-2** (`gen-figure.sh` / `gen-image.sh` / `gen-illust.sh`). Ver [.claude/content/images.md](.claude/content/images.md).
- **`scripts/linkedin/`** — fila (`queue.yaml`) + cron publisher de posts do Dr. Getúlio. Skill `/linkedin-week`.
- **`scripts/enrichment/`** — enrichment científico dos ScoreItems (RAG + Claude). Ver [.claude/workflows/enrichment-automation.md](.claude/workflows/enrichment-automation.md).
- **`ebook/`** — eBooks do Dr. Getúlio (Série AGORA, Série Bases). Skill `/ebook`. Ver [.claude/content/ebooks.md](.claude/content/ebooks.md).
- **`.claude/skills/`** — `plenya-deck`, `linkedin-week`, `lecture-builder` (`/aula`), `ebook-builder` (`/ebook`), `plano-paciente` (`/plano` — devolutiva de resultados), `responder-insta`, `pptx`. Ver [.claude/social/](.claude/social/).

## 🚨 Regras de Ouro (valem em TODO o monorepo)

### 1. Fonte única de verdade
**Go models (`apps/api/internal/models/*.go`) são a única fonte.** Nunca editar arquivos gerados:

```
apps/api/internal/models/*.go  ← EDITAR AQUI
         ├─→ Swag  → apps/api/docs/swagger.json                (gerado)
         ├─→ openapi-typescript → packages/types/src/generated/api-types.ts   (gerado)
         └─→ openapi-zod-client → packages/types/src/generated/api-schemas.ts (gerado)
```

**Schema/migrations NÃO são derivados automaticamente dos models.** São migrations **goose**
escritas à mão (`apps/api/database/migrations/`), aplicadas no deploy via `cmd/migrate`; AutoMigrate
fica desligado por default. Mudou um model → escreva a migration SQL correspondente.
Ver [docs/emr/migrations-decisao.md](docs/emr/migrations-decisao.md).

Para conteúdo de marca/clínico, a fonte é o **site** + os arquivos canônicos
(`apps/site/lib/agir-structure.ts`, brandbook). Antes de gerar deck/post/copy, leia a fonte —
nunca chute dados verificáveis (versões, contagens, nomes, métricas).

### 2. Desenvolvimento vs Produção
| Contexto | Método |
|----------|--------|
| Dev (você, Claude) manipulando dados | Docker exec → psql, ou Go scripts |
| Produção (apps web/mobile) | API HTTP |

**❌ NUNCA usar API HTTP (curl/POST) para manipulação manual em dev. ✅ SEMPRE banco direto.**

### 3. Hooks obrigatórios
- **Backend:** todos os models com `BeforeCreate` (UUID v7); `Patient` com `BeforeSave`/`AfterFind` (cripto CPF/RG); `ScoreItem`/`ScoreLevel` com `BeforeUpdate` (`LastReview`).
- **Frontend:** todos os formulários com `useFormNavigation({ formRef })`; páginas de dados de paciente com `useRequireSelectedPatient()`.

### 4. Voz e regras editoriais (conteúdo público)
Em IG/FB/LinkedIn/blog/decks/DMs: **sem maneirismos de IA** (sem travessões/em-dash, sem
"Não é X. É Y.", sem fechos-slogan, sem ícones decorativos em listas), **sem preços**, **sem
marcas comerciais** (wearables, suplementos, varejistas), **sem "medicina preditiva"**, **sem
casos clínicos identificáveis**. Tom: prosa clínica conectiva PT-BR. Tagline: "Saúde,
Performance & Longevidade". Detalhes na memória (`MEMORY.md`) e em `.claude/social/`.

## 🛠 Comandos essenciais

```bash
# Banco direto (DESENVOLVIMENTO)
docker compose exec db psql -U plenya_user -d plenya_db
docker compose exec db psql -U plenya_user -d plenya_db -c "\d score_items"

# Migrations (goose) — schema NÃO é auto-derivado dos models
docker compose exec -w /app api go run ./cmd/migrate up        # aplica pendentes
docker compose exec -w /app api go run ./cmd/migrate status    # estado

# Geração de OpenAPI + tipos TS/Zod após editar Go models (NÃO inclui migrations)
pnpm generate

# Compilar Go (não há Go local — usar container)
docker compose exec -w /app api go build ./...

# Subir tudo / logs
docker compose up -d
docker compose logs -f api
```

## 🚀 Deploy (produção / Coolify)

**Auto-deploy está DESLIGADO de propósito.** `git push origin master` NÃO deploya nada — a VPS
tem 8GB e rebuildar os 3 apps juntos (via webhook) causa OOM. Deploy é **manual, deliberado e
por-app**: só deploya o app cujo código mudou, um de cada vez.

```bash
# Deploy de UM app (faz higiene da fila + dispara + espera container novo subir):
scripts/deploy/deploy-app.sh api    # apps/api  → plenya-api
scripts/deploy/deploy-app.sh web    # apps/web  → plenya-web (EMR + portal)
scripts/deploy/deploy-app.sh site   # apps/site → plenya-site
```

Regras: **nunca deployar sem ordem explícita do usuário** (memória `plenya_no_deploy_sem_ordem`);
mudou só `apps/web`? deploya só `web`. Migrations goose rodam no deploy do `api`
(`RUN_MIGRATIONS=true`). Detalhe do procedimento e recuperação de deploy travado nas memórias
`plenya_deploy_manual` e `coolify_deploy_orphan_lock_procedure`.

## 📚 Documentação detalhada (`.claude/`)

- Fundação: [01-overview](.claude/01-overview.md) · [02-stack](.claude/02-stack.md) · [03-architecture](.claude/03-architecture.md)
- Backend: [models](.claude/backend/models.md) · [hooks](.claude/backend/hooks.md) · [service-layer](.claude/backend/service-layer.md) · [api-endpoints](.claude/backend/api-endpoints.md) · 🔥 banco direto em [workflows/database-ops](.claude/workflows/database-ops.md)
- Frontend: [form-navigation](.claude/frontend/form-navigation.md) · [patient-context](.claude/frontend/patient-context.md) · [tanstack-query](.claude/frontend/tanstack-query.md)
- Domínio: [🎯 score-system](.claude/domain/score-system.md) · [patients](.claude/domain/patients.md) · [security](.claude/domain/security.md)
- Workflows: [development](.claude/workflows/development.md) · [🔥 database-ops](.claude/workflows/database-ops.md) · [🤖 enrichment-automation](.claude/workflows/enrichment-automation.md) · [adding-features](.claude/workflows/adding-features.md) · [dev-bypass-auth](.claude/workflows/dev-bypass-auth.md)
- Conteúdo: [decks](.claude/content/decks.md) · [ebooks](.claude/content/ebooks.md) · [images](.claude/content/images.md)
- Social: [.claude/social/](.claude/social/)
- Mobile: [setup](.claude/mobile/setup.md) · [security](.claude/mobile/security.md) · [deploy](.claude/mobile/deploy.md) · [ota-policy](.claude/mobile/ota-policy.md) · [release-checklist](.claude/mobile/release-checklist.md)
- Infra/VPS: ver memórias `plenya_vps*` (Coolify, Stalwart mailserver, deploy git-push)
- **Backups (banco + uploads): [scripts/backups/README.md](scripts/backups/README.md).** VPS gera por
  cron (db 03:00 / uploads 03:30, rotação) e a máquina de dev espelha em `/home/user/backups/prod/vps/`.
  **Sempre salvar backup do Plenya em `/home/user/backups/{dev,prod}` — nunca criar pasta solta.**
  Persistência dos uploads (bind mount Coolify) em [docs/emr/plano-persistencia-uploads-vps.md](docs/emr/plano-persistencia-uploads-vps.md).

## 🚀 Roadmap

- [x] Fase 1-4: Backend core + Auth/RBAC + Frontend web + Sistema de Escores
- [x] Portal do paciente, CRM/Central de Conversas, Calendar V1 + telemedicina, RAG (embeddings + semantic search), Subscriptions, módulo Training
- [ ] **Fase 5: Mobile** — `apps/mobile-pro` (~MVP em curso), `apps/mobile-app` (paciente, inicial)
- [ ] Fase 6: Hardening LGPD · Fase 7: Deploy produção (já no ar em Hetzner/KingHost + Coolify)

---
**Versão:** 4.0 (multi-subprojeto) · **Atualizado:** Maio 2026
Para detalhes, sempre consulte o `CLAUDE.md` do subprojeto e os arquivos em `.claude/`.
