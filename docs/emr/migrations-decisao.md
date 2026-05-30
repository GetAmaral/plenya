# Estratégia de Migrations do EMR Plenya — Análise e Decisão

> **Status:** ✅ DECIDIDO + IMPLEMENTADO EM DEV (2026-05-30) — **goose** (Opção B).
> Pesquisa + verificação direta em 2026-05-30. Fatos do Atlas conferidos na matriz oficial
> `atlasgo.io/features` (Atlas v1.2, abril/2026), não em memória.
>
> **Feito e testado em dev:** `cmd/migrate` (goose, comandos up/status/version/down/stamp, com
> auto-adoção em banco pré-existente); baseline `database/migrations/00001_baseline.sql` gerado de
> `pg_dump` do dev (testado recriando o schema inteiro num banco vazio: 93 tabelas + 402 funções +
> 7 extensões); 58 .sql legados + atlas.sum arquivados em `_legacy/`; AutoMigrate desligado por
> default em `main.go` (vira escape hatch via `MIGRATIONS_AUTO=true`); `entrypoint.sh` (dev) roda
> `migrate up`; dev DB `plenya_db` adotado (goose v1). Server reinicia e sobe com "AutoMigrate
> desligado — schema via migrations goose".
>
> **Preparado, NÃO ativado (toca prod — confirmar):** `Dockerfile` compila/copia o binário `migrate`;
> `prod-entrypoint.sh` roda `migrate up` só com `RUN_MIGRATIONS=true` (gate; default não roda, prod
> segue idêntico). Cutover de prod (conferir drift vs baseline, backup, ativar o gate, garantir
> pgvector na imagem de prod) pendente de OK do usuário.

## 1. O problema (estado atual)

Inconsistência herdada: a documentação (`CLAUDE.md` raiz, `apps/api/CLAUDE.md`, `.claude/03-architecture.md`,
`02-stack.md`) afirma que **Atlas** gera as migrations e que `pnpm generate` roda `atlas diff` — mas:

- Não existe `atlas.hcl`; Atlas nunca foi configurado. `atlas`/`swag` não estão instalados.
- Há **59 `.sql` escritos à mão** em `apps/api/database/migrations/` (cabeçalho manual), de utilidade
  duvidosa, sem runner que os aplique e sem tabela de controle de versão.
- O schema real é criado pelo **GORM AutoMigrate** no startup (`database.go`), em dev sempre e em
  prod só atrás de `MIGRATIONS_AUTO` (`main.go:72-93`). Comentário de "Atlas piloto" abandonado.

Resultado: nenhuma das três peças (docs, .sql, AutoMigrate) concorda com as outras.

## 2. Complexidade do schema (verificado no código)

Pesa na escolha da ferramenta. Counts reais:
- **pgvector**: colunas `vector` + índices ivfflat (embeddings/RAG).
- **EXCLUDE/gist**: anti-overlap de agenda em `appointments` (+ coluna materializada `end_at`).
- **14 funções** plpgsql + **5 triggers**.
- Extensões: `uuid-ossp`, `btree_gist`, `vector`.
- **7 blocos `DO`** em `database.go` (constraints CHECK/EXCLUDE recriadas fora do AutoMigrate).
- ~74 models, status via VARCHAR+CHECK (não enums nativos), UUID v7 por `BeforeCreate`, CPF/RG cripto.

## 3. O que é grátis vs pago no Atlas (oficial, v1.2 — abril/2026)

Atlas é open-core (engine Apache 2.0) + Atlas Pro ($9/assento/mês, trial 30 dias, **sem plano
grátis permanente**; "Hacker License" grátis só para projetos **não-comerciais**, o que exclui o Plenya).

| Recurso | Tier | Aplica ao Plenya |
|---|---|---|
| Schema diffing · Versioned migrations · **External Schema** (provider GORM) | **Grátis** | Núcleo "models como fonte" é grátis |
| Tabelas, colunas, índices, FK, **CHECK/UNIQUE/EXCLUDE**, enums, comments (Postgres) | **Grátis** | Cobre as ~74 tabelas + EXCLUDE da agenda |
| **Extensões, Functions, Triggers, Views, Materialized views, Sequences** (Postgres) | **Pro** | Plenya tem 3 ext + 14 funções + 5 triggers |
| **Composite Schema** + Blob Directory (data sources) | **Pro** | Era a peça central da 1ª proposta |
| Migration lint avançado (concurrent index, table-locking, linear history) | **Pro** | Desejável p/ deploy seguro, não essencial |

**Consequência:** o Atlas grátis gera/versiona as tabelas a partir dos models GORM (mantém a Regra
de Ouro), porém é **cego para extensões/funções/triggers** — justamente o Postgres avançado do
Plenya. Para um fluxo Atlas único e coeso cobrindo o schema inteiro, é **Atlas Pro ($9/mês)**.
A proposta anterior ("Atlas grátis com composite_schema") é **inviável**: composite_schema E os
objetos que ele gerenciaria (extensões/funções/triggers) são todos Pro.

## 4. Opções definitivas (decisão do usuário)

Só há duas respostas "limpas" (mecanismo único). A terceira (Atlas grátis + bootstrap SQL à parte)
reintroduz a fratura de dois mecanismos que causou a bagunça atual e tem risco no tipo `vector`.

### Opção A — Atlas Pro ($9/assento/mês)
- **Prós:** entrega literalmente a Regra de Ouro (models GORM → `atlas migrate diff` gera o `.sql`
  versionado), gerencia o schema **inteiro** (tabelas + extensões + pgvector + funções + triggers +
  EXCLUDE) num fluxo só, lint de segurança (table-locking, concurrent index), `atlas_schema_revisions`,
  apply no deploy. É a opção que faz a doc existente virar verdade.
- **Contras:** custo recorrente ($9/mês para 1 assento — desprezível para um negócio de saúde, mas é
  um SaaS pago a mais; sem login, vira read-only). Dependência de conta Atlas.

### Opção B — goose (grátis, open-source MIT)
- **Prós:** 100% grátis, maduro, simples; lida com **tudo** em SQL (pgvector, funções, triggers,
  EXCLUDE, `CREATE INDEX CONCURRENTLY` via `-- +goose NO TRANSACTION`); versionado, up/down,
  aplicado no deploy por um runner. Formaliza o que o projeto **já faz** (o DDL complexo já é escrito
  à mão hoje em `database.go`).
- **Contras:** abre mão de "models geram o SQL" — ao mudar um model, escreve-se a migration SQL à mão
  (o struct GORM e o DDL ficam sincronizados por disciplina/revisão, não por diff automático).

> **Recomendação:** se $9/mês é aceitável, **Opção A (Atlas Pro)** — é a única que honra a Regra de
> Ouro de ponta a ponta e cobre o schema inteiro. Se preferir custo zero e tolera escrever o SQL das
> migrations à mão, **Opção B (goose)** é a escolha pragmática e robusta (e formaliza a prática atual).

## 5. Plano de implementação (comum às duas opções, ramifica no passo 3)

1. **[SEGURO]** Desligar AutoMigrate como caminho padrão: `MIGRATIONS_AUTO` default `false` (dev e prod);
   AutoMigrate vira só escape hatch manual de prototipagem. Em prod, desligado em definitivo.
2. **[TOCA PROD — confirmar]** Confirmar o que a PROD (Coolify) roda hoje: imagem do Postgres tem
   pgvector? As extensões existem? (Backup verificado antes de qualquer coisa.) `docker-compose.prod.yml`
   versionado diz `postgres:17-alpine` (sem pgvector), mas pode estar obsoleto vs Coolify — checar.
3. **Configurar o mecanismo escolhido:**
   - **A (Atlas Pro):** `atlas login`; `apps/api/tools.go` + `atlas-provider-gorm`; `atlas.hcl` (env gorm,
     dev-db pgvector/pg17 efêmero); baseline do schema de prod (`atlas migrate diff baseline` contra
     prod read-only); marcar baseline como aplicada (`migrate apply --baseline`); arquivar os 59 .sql
     em `_legacy/`.
   - **B (goose):** adicionar `goose` (lib + CLI embutido num pequeno binário Go); pasta
     `apps/api/database/migrations/` no formato goose; gerar a **migration de baseline** a partir de
     `pg_dump --schema-only` do schema atual de prod (registra como já aplicada via `goose`); arquivar
     os 59 .sql antigos.
4. **[SEGURO]** Aplicar migrations no DEPLOY (não no boot): entrypoint/initstep roda `migrate apply`
   (Atlas) ou `goose up` antes de a app ficar Ready; credencial DDL elevada só nesse passo, app em
   privilégio mínimo. Dev roda o mesmo no `entrypoint.sh` (paridade dev/prod).
5. **[SEGURO]** Workflow de mudança: editar model → (A) `atlas migrate diff <nome>` e revisar o .sql /
   (B) escrever a migration SQL à mão → revisar em commit → aplicar em dev → deploy aplica em prod.
   Objetos não-ORM (função/trigger/índice vetorial): (A) Pro gerencia / (B) SQL na migration.
6. **[SEGURO]** Corrigir TODA a documentação (ver §6) para refletir o mecanismo real e único.
7. **[SEGURO]** Forward-only em prod (dados de saúde): erro vira migration corretiva, nunca down;
   backup é a rede. `CREATE INDEX CONCURRENTLY` fora de transação (`txmode none` / `NO TRANSACTION`).

## 6. Documentação a corrigir (independe da opção)

Priorizado (da auditoria): `.claude/03-architecture.md` (todo o fluxo de migrations), `02-stack.md`
(linha 446 "Auto migrations (via Atlas)" está errada — AutoMigrate é do GORM), `.claude/workflows/development.md`
(seção Migrations + troubleshooting), `apps/api/CLAUDE.md` (8,25,49,51-52), `CLAUDE.md` raiz (74,110-111
diagrama fonte única), `.claude/workflows/adding-features.md`, `package.json`/`scripts/generate-all.sh`/`Makefile`
(comando real, sem o `|| echo` que mascara falha), `.claude/01-overview.md:126`, `database.go:70-73`
(TODO M9). Corrigir também o comentário "HNSW" vs índice real `ivfflat` em `20260213_install_pgvector.sql`.

## 7. Riscos e mitigação

- **Drift prod vs models** (AutoMigrate só adiciona): comparar `schema inspect`/`pg_dump` de prod com o
  schema derivado dos models ANTES de marcar baseline; resolver diffs conscientemente.
- **Objetos só em `database.go`** (EXCLUDE, end_at, DO blocks, CHECKs): transcrever para o baseline na
  transição, senão o 1º diff "removeria" o que a ferramenta não conhece.
- **pgvector em prod**: pré-requisito que a imagem de prod tenha a extensão; sem isso o RAG quebra,
  independente da ferramenta.
- **Baseline de banco populado**: backup verificado + credencial read-only na inspeção.

## Fontes
- atlasgo.io/features (matriz open vs Pro, v1.2) · atlasgo.io/pricing ($9/seat) ·
  atlasgo.io/hacker-license (não-comercial) · github.com/pressly/goose · gorm.io/docs/migration.html.
