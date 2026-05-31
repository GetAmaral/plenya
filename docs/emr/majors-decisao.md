# Upgrades MAJOR do EMR — Decisão (maio/2026)

> Pesquisa web (fontes oficiais) + blast-radius no código + decisão. Inclui **PostgreSQL** (que
> faltava na análise de versões `atualizacao-versoes.md`). Ordem e racional abaixo. Nenhum tem
> pressão de CVE no app, exceto o pgvector (ver PG vertente a).

## Ordem recomendada
**PG vertente (a)** → **gofpdf** → **[PG 17→18]** → **[Tailwind v4]** → **(Fiber v3 adiado)**.

| Major | Atual | Latest/Status | Esforço | Risco | Decisão |
|---|---|---|---|---|---|
| **PostgreSQL (a)** patch 17.x + pgvector + imagem prod | PG 17.7, pgvector 0.8.1, prod sem pgvector | PG 17.10 (14/05/2026), pgvector 0.8.2 (25/02/2026) | baixo | baixo | **FAZER-AGORA** [toca prod] |
| **gofpdf** → codeberg.org/go-pdf/fpdf | jung-kurt/gofpdf v2.17.3 (órfão) | go-pdf/fpdf v0.12.0 (18/05/2026) | baixo (~1-2h) | baixo | **FAZER-AGORA** (não toca prod) |
| **PostgreSQL (b)** major 17→18 | PG 17.x | PG 18.4 (GA 25/09/2025) | alto | alto | **JANELA DEDICADA** [toca prod] |
| **Tailwind v4** | 3.4.18 | 4.3.0 (GA jan/2025) | alto | médio | **JANELA DEDICADA** |
| **Fiber v3** | v2.52.12 | v3.0.0 (GA 02/02/2026) | alto | médio | **ADIAR** |

## PostgreSQL — destaque

### (a) FAZER-AGORA (baixo risco, alto valor; casa com o cutover de migrations goose)
1. **Corrigir divergência de imagem em prod:** `postgres:17-alpine` → `pgvector/pgvector:pg17` em
   `docker-compose.prod.yml`. Sem isso prod não faz `CREATE EXTENSION vector` → RAG não funciona em prod.
2. **pgvector 0.8.1 → 0.8.2** (dev+prod): fecha **CVE-2026-3172** (buffer overflow no build paralelo
   de HNSW que vaza dados de outras relações — inaceitável sob LGPD). Após atualizar imagem:
   `ALTER EXTENSION vector UPDATE TO '0.8.2';` + `REINDEX` dos índices HNSW/ivfflat.
3. **Patch do servidor 17.7 → 17.10** (security 14/05/2026): trocar a tag para o patch atual.
- Mesma major 17 → volume preservado. **[TOCA PROD — confirmar + backup + janela curta.]**

### (b) Major 17 → 18 — JANELA DEDICADA (após estabilizar (a), testar em staging)
- Ganhos: Async I/O (io_uring, 2-3x em cargas I/O-bound), `pg_upgrade` preserva planner stats,
  `uuidv7()` nativo (alinha com UUID v7 dos models), skip scan B-tree, virtual generated columns.
- Breaking: formato on-disk muda (exige `pg_upgrade`/dump-restore), **quirk Docker:** `postgres:18`
  mudou o layout do PGDATA (`/var/lib/postgresql/18/docker`) → revisar mapeamento de volume no Coolify;
  pgvector compatível precisa existir na imagem de destino ANTES do upgrade (`pg_upgrade` aborta
  se extensão faltar); `md5` auth deprecado (preferir `scram-sha-256`).
- Caminho: backup (pg_dumpall + snapshot Hetzner) → imagem `pgvector/pgvector:pg18` → parar app →
  `pg_upgrade` (ou `pgautoupgrade`) → `ALTER EXTENSION vector UPDATE;` + `REINDEX` → validar GORM/suíte/RAG.
  Ensaiar em staging com cópia dos dados. Downtime poucos~30min. **NÃO esperar PG19** (só beta 04/06/2026).

## gofpdf — substituição (FEITO/em execução nesta data)
- 3 arquivos: `pdf_service.go`, `payment_service.go`, `prescription_pdf_service.go` (~1.368 linhas).
- `codeberg.org/go-pdf/fpdf` é o MESMO lib que migrou do GitHub p/ Codeberg (não morreu); API idêntica.
- Caminho: `go get codeberg.org/go-pdf/fpdf@v0.12.0` + alias `import gofpdf "codeberg.org/go-pdf/fpdf"`
  (zera diffs no corpo) + `go build ./...` + **smoke: gerar 1 recibo + 1 laudo + 1 prescrição e comparar**
  (fontes OpenSans UTF-8, ImageOptions PNG). NÃO migrar para go-rod (cold-start/memória Chromium).

## Tailwind v4 — JANELA DEDICADA
- v4.3.0 madura, shadcn já v4-first. Ganho só DX/perf (build ~3.5x). Risco: regressão visual em massa
  (~251 componentes), renomes (`shadow`→`shadow-sm`, `ring`→`ring-3`, `rounded`→`rounded-sm`...),
  `border`/`ring` default → currentColor, migração do preset `@plenya/brand` (gold/petrol/ocean/sage/cream)
  para `@theme`, troca `tailwindcss-animate`→`tw-animate-css`. Piso de browser: Safari 16.4+/Chrome 111+.
- Caminho: worktree efêmero + `npx @tailwindcss/upgrade` → portar `@plenya/brand` → shims anti-regressão →
  QA visual telas-a-tela → **coordenar com site e site-getulio** (consomem `@plenya/brand`).

## Fiber v3 — ADIAR
- v3.0.0 GA (02/02/2026), ~3,5 meses. Blast-radius: 442 `*fiber.Ctx` (vira interface), 7 middlewares
  custom + 27 `c.Locals` (passam a `FromContext` tipado) = núcleo auth/RBAC/audit. CVEs da v2 já cobertas
  no v2.52.12 sobre Go 1.25; linha v2 ainda mantida. Sem gatilho → puro custo num núcleo de risco.
- Reavaliar só com feature que dependa de `net/http` std compat / `Bind` unificado, ou janela dedicada.
  Manter v2.52.x no último patch.

## Plano de execução
- **Já (sem prod):** gofpdf swap + build + smoke de PDFs.
- **Já [TOCA PROD — confirmar + backup]:** PG vertente (a) — imagem pgvector/pg17 17.10, `ALTER EXTENSION
  vector UPDATE TO '0.8.2'`, `REINDEX`. **Casar com o cutover de migrations goose pendente** (`docs/emr/migrations-decisao.md`).
- **Janela dedicada + staging:** PG 17→18; Tailwind v4.
- **Adiado:** Fiber v3.

## Status de execução (2026-05-30)

**✅ gofpdf → codeberg.org/go-pdf/fpdf — FEITO (build verificado):**
- `go get codeberg.org/go-pdf/fpdf@v0.12.0`; imports trocados com alias `gofpdf` em
  `prescription_pdf_service.go`, `payment_service.go`, `pdf_service.go` (zero mudança no corpo);
  `go mod tidy` removeu `jung-kurt/gofpdf`; `go build ./...` → **OK**.
- ⚠️ **Pendente:** smoke-test de saída (gerar 1 recibo + 1 laudo + 1 prescrição reais pelo app e
  comparar rendering — fontes OpenSans UTF-8, ImageOptions PNG). Não há teste unitário de PDF; build
  não garante pixel-igualdade. Fazer antes de confiar em prod.

**✅ PostgreSQL 18 — DEV migrado e verificado (2026-05-30):** decisão do usuário foi PG18 em dev E prod
(paridade total). Dev: imagem `pgvector/pgvector:pg18` (PG 18.4 + pgvector 0.8.2), via dump→restore
(296MB-ish, evitou pg_upgrade). Corrigido o quirk PG18: volume montado no PARENT `/var/lib/postgresql`
(dados em `/18/docker`); `init.sql` redundante removido (uuid_v7 vem da baseline goose). Verificado:
94 tabelas, dados intactos (18 pacientes, 38.277 embeddings), pgvector 0.8.2, similaridade + REINDEX OK,
api sobe via goose com AutoMigrate desligado. Dump PG17 de segurança em `/tmp/dev_pg17.dump`.

Estado real verificado das duas pontas ANTES (divergência que a migração alinha):
| | Dev (antes) | Prod (antes) | Alvo |
|---|---|---|---|
| PostgreSQL | 17.7 | 17.9 | **18.4** |
| pgvector | 0.8.1 | 0.8.2 | **0.8.2** |
| Extensões extra | pg_trgm, unaccent | (sem) | reconciliar no cutover |

**✅ PROD PG18 — MIGRADO E VERIFICADO (2026-05-30):** recurso Coolify `mb511beqjtgd7nsjlnngh3m6`
agora `pgvector/pgvector:pg18` (PG 18.4 + pgvector 0.8.2). Caminho usado (beta, dados reais de
RAG/escores preservados): backup verificado off-server (`~/.plenya-vps-secrets/backups/prod_pg17_*.dump`)
→ PATCH image pg18 via API Coolify → stop+rm container → wipe volume → start (recria pg18) → restore
do dump do prod → delta de paridade (pg_trgm/unaccent/immutable_unaccent + 4 índices trigram) →
ivfflat em article_embeddings → REINDEX. **Zero perda** (contagens idênticas: 77 tabelas, 38.277
embeddings, 961 score_items, 272 lab defs, 1.183 artigos, 2 users, 1 paciente). api conecta, RAG ok,
app./minha./api respondem.

**Quirk PG18 + Coolify (aprendizado, registrar):** a imagem pg18 muda o PGDATA default p/
`/var/lib/postgresql/18/docker`, mas o Coolify (beta antigo) monta o volume em `/var/lib/postgresql/data`
→ mismatch → container aborta ("in 18+ ... unused mount"). **Fix:** criar env var `PGDATA=/var/lib/postgresql/data`
no recurso Coolify (via `POST /api/v1/databases/<uuid>/envs`) — `custom_docker_run_options` é bloqueado
pela API. No dev (compose próprio) o fix foi montar o volume no parent `/var/lib/postgresql`.

**Paridade restante (drift, separado):** prod tem 77 tabelas vs 93 do dev (features novas
ainda não deployadas — secretaria/payments/waitlist). Isso é o cutover de schema/goose + deploy de
código, não o PG18.

**✅ DEPLOY CONSOLIDADO + PARIDADE DE SCHEMA (2026-05-30):** commit `a2c0b8f6` (126 arquivos do EMR,
excluindo lixo não-relacionado) → push → Coolify rebuildou api + web. Prod no código novo (goose infra
gated, secretaria backend, deps, gofpdf). AutoMigrate criou as 4 tabelas da secretaria (77→81). Depois
reconciliei o drift pré-existente: prod estava sem ~10 tabelas reais (patient_score_*, enrichment,
api_usage, auto_link_*, etc.) que vinham de .sql legados fora do AutoMigrate — extraí o schema do dev e
criei no prod (81→91, vazias). Corrigida a `.gitignore` que ignorava as migrations goose (quebraria o
build). **Paridade real atingida:** prod=91 tabelas; diff vs dev só nos 2 backups-fantasma + goose_db_version.
Dados de prod 100% preservados (38.277 embeddings etc.). api/web no ar, PG 18.4.

**⏳ Restante:** (a) **cutover goose no prod** (opcional; prod segue em AutoMigrate, funcional — adotar
goose = stamp baseline + RUN_MIGRATIONS=true + MIGRATIONS_AUTO=false); (b) **Tailwind v4** (janela
dedicada); (c) Fiber v3 (adiado); (d) smoke visual do /recepcao pelo usuário.
