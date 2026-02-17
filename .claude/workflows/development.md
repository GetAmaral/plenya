# Development Workflow

## Setup Inicial

### Requisitos

- Docker 27+
- Node.js 24 LTS
- pnpm 10+
- Go 1.25+ (opcional, apenas se rodar fora do Docker)

### Primeira Execução

```bash
# 1. Clonar repositório
git clone <repo-url>
cd plenya

# 2. Copiar .env de exemplo
cp apps/api/.env.example apps/api/.env
cp apps/web/.env.local.example apps/web/.env.local

# 3. Iniciar containers
docker compose up -d

# 4. Verificar logs
docker compose logs -f

# 5. Acessar aplicação
# Web: http://localhost:3000
# API: http://localhost:3001
# Swagger: http://localhost:3001/swagger/index.html
```

### Dev Bypass Auth (Opcional)

Para desenvolvimento rápido sem login manual, ative o bypass auth que injeta automaticamente `admin@plenya.com`:

```bash
# apps/api/.env
DEV_BYPASS_AUTH=true

# apps/web/.env.local
NEXT_PUBLIC_DEV_BYPASS_AUTH=true

# Restart
docker compose restart api web
```

**Ver documentação completa:** [dev-bypass-auth.md](./dev-bypass-auth.md)

⚠️ **NUNCA ativar em produção!**

### Verificar Status

```bash
# Status dos containers
docker compose ps

# Logs específicos
docker compose logs -f api
docker compose logs -f web
docker compose logs -f db

# Acessar container
docker compose exec api sh
docker compose exec web sh
docker compose exec db psql -U plenya_user -d plenya_db
```

## Workflow Diário

### Iniciar Desenvolvimento

```bash
# Iniciar todos os serviços
docker compose up -d

# Ver logs em tempo real
docker compose logs -f
```

### Parar Desenvolvimento

```bash
# Parar containers (preserva dados)
docker compose down

# Parar e remover volumes (limpa tudo)
docker compose down -v
```

### Hot Reload

**Backend (Go):**
- Air detecta mudanças em `*.go`
- Recompila e reinicia automaticamente
- Ver logs: `docker compose logs -f api`

**Frontend (Next.js):**
- Fast Refresh detecta mudanças em `*.ts`, `*.tsx`
- Recarrega automaticamente no browser
- Ver logs: `docker compose logs -f web`

**Database:**
- Mudanças em migrations requerem `atlas migrate apply`

## Adicionar Dependências

### Backend (Go)

```bash
# Dentro do container
docker compose exec api go get github.com/pacote/exemplo
docker compose exec api go mod tidy

# OU fora do container (se tiver Go instalado)
cd apps/api
go get github.com/pacote/exemplo
go mod tidy
```

**Importante:** Após adicionar dependências, rebuild:

```bash
docker compose up -d --build api
```

### Frontend (Next.js/Web)

```bash
# Dentro do container
docker compose exec web pnpm add <pacote> --filter web

# OU fora do container
pnpm add <pacote> --filter web
```

Rebuild:

```bash
docker compose up -d --build web
```

### Frontend (Mobile)

```bash
cd apps/mobile
pnpm add <pacote>

# Se dependência nativa
npx expo install <pacote>
```

## Workflow de Mudanças

### Caso 1: Adicionar/Modificar Go Model

```bash
# 1. Editar model
vim apps/api/internal/models/patient.go

# 2. Gerar tudo automaticamente
pnpm generate

# Isso executa:
# - atlas migrate diff (cria migration SQL)
# - swag init (gera OpenAPI)
# - pnpm generate:types (TypeScript types)

# 3. Revisar gerados
git diff apps/api/database/migrations/
git diff apps/api/docs/swagger.json
git diff packages/types/src/generated/

# 4. Aplicar migration
atlas migrate apply --env dev

# 5. Testar
docker compose logs -f api
```

### Caso 2: Adicionar Endpoint na API

```bash
# 1. Adicionar método no repository
vim apps/api/internal/repository/patient_repository.go

# 2. Adicionar método no service
vim apps/api/internal/services/patient_service.go

# 3. Adicionar handler com Swagger annotations
vim apps/api/internal/handlers/patient_handler.go

# 4. Registrar rota
vim apps/api/cmd/server/main.go

# 5. Gerar OpenAPI
pnpm generate:swagger

# 6. Ver Swagger UI
open http://localhost:3001/swagger/index.html
```

### Caso 3: Adicionar Página no Frontend

```bash
# 1. Criar arquivo de página
vim apps/web/app/(authenticated)/pacientes/novo/page.tsx

# 2. Criar queries (se necessário)
vim apps/web/lib/api/patients.ts

# 3. Testar
open http://localhost:3000/pacientes/novo
```

## Debugging

### Backend (Go)

```bash
# Logs detalhados
docker compose logs -f api

# Acessar container
docker compose exec api sh

# Ver variáveis de ambiente
docker compose exec api env | grep DB

# Rodar comando Go
docker compose exec api go version
```

### Frontend (Next.js)

```bash
# Logs do servidor
docker compose logs -f web

# Logs do browser (Chrome DevTools)
# F12 → Console

# Verificar build
docker compose exec web pnpm build --filter web
```

### Database

```bash
# Acesso interativo
docker compose exec db psql -U plenya_user -d plenya_db

# Ver conexões ativas
docker compose exec db psql -U plenya_user -d plenya_db -c "SELECT * FROM pg_stat_activity;"

# Ver tamanho do banco
docker compose exec db psql -U plenya_user -d plenya_db -c "SELECT pg_size_pretty(pg_database_size('plenya_db'));"

# Ver tabelas
docker compose exec db psql -U plenya_user -d plenya_db -c "\dt"

# Descrever tabela
docker compose exec db psql -U plenya_user -d plenya_db -c "\d score_items"
```

## Testes

### Backend (Go)

```bash
# Rodar todos os testes
docker compose exec api go test ./...

# Testes de um pacote específico
docker compose exec api go test ./internal/services

# Com coverage
docker compose exec api go test -cover ./...

# Verbose
docker compose exec api go test -v ./...
```

### Frontend (Web)

```bash
# Rodar testes
docker compose exec web pnpm test --filter web

# Watch mode
docker compose exec web pnpm test:watch --filter web
```

## Migrations

### Criar Migration

```bash
# Após editar Go model
pnpm generate:migrations

# Isso cria migration em:
# apps/api/database/migrations/<timestamp>_<nome>.sql
```

### Aplicar Migrations

```bash
# Desenvolvimento
atlas migrate apply --env dev

# Ver status
atlas migrate status --env dev

# Rollback (manualmente via SQL)
docker compose exec db psql -U plenya_user -d plenya_db -f apps/api/database/migrations/<timestamp>_down.sql
```

### Resetar Database

```bash
# ⚠️ CUIDADO: Remove TODOS os dados
docker compose down -v
docker compose up -d
atlas migrate apply --env dev
```

## Linting e Formatação

### Backend (Go)

```bash
# Formatar código
docker compose exec api gofmt -w .

# Linting
docker compose exec api golangci-lint run
```

### Frontend (TypeScript)

```bash
# Lint
pnpm lint --filter web

# Format com Prettier
pnpm format --filter web
```

## Git Workflow

```bash
# 1. Criar branch
git checkout -b feature/nova-feature

# 2. Fazer mudanças
vim apps/api/internal/models/patient.go
pnpm generate

# 3. Commit (seguir conventional commits)
git add .
git commit -m "feat: adiciona campo email ao Patient"

# 4. Push
git push origin feature/nova-feature

# 5. Abrir PR
# Via GitHub UI ou gh CLI
gh pr create --title "feat: adiciona campo email ao Patient"
```

## Troubleshooting

### Port já em uso

```bash
# Ver o que está usando a porta
lsof -i :3000
lsof -i :3001
lsof -i :5432

# Matar processo
kill -9 <PID>
```

### Container não inicia

```bash
# Ver logs de erro
docker compose logs api

# Rebuild completo
docker compose down
docker compose up -d --build

# Verificar Dockerfile
vim apps/api/Dockerfile
```

### Migrations falhando

```bash
# Ver status
atlas migrate status --env dev

# Ver histórico
docker compose exec db psql -U plenya_user -d plenya_db -c "SELECT * FROM atlas_schema_revisions;"

# Resetar (cuidado!)
docker compose down -v
docker compose up -d
atlas migrate apply --env dev
```

### Hot reload não funciona

**Go (Air):**
```bash
# Verificar .air.toml
vim apps/api/.air.toml

# Restart container
docker compose restart api
```

**Next.js:**
```bash
# Verificar next.config.ts
vim apps/web/next.config.ts

# Limpar cache
docker compose exec web pnpm clean --filter web
docker compose restart web
```

## Performance

### Analisar Build

```bash
# Backend
docker compose exec api go build -o /dev/null -v ./cmd/server

# Frontend
docker compose exec web pnpm build --filter web
```

### Monitorar Recursos

```bash
# Docker stats
docker stats

# Memória do banco
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT
  pg_size_pretty(pg_database_size('plenya_db')) AS db_size,
  pg_size_pretty(sum(pg_total_relation_size(relid))) AS tables_size
FROM pg_stat_user_tables;
"
```

## Dicas

1. **Use `pnpm generate` sempre** após mudar Go models
2. **Nunca edite arquivos gerados** (migrations, swagger, types)
3. **Commit antes de grandes mudanças** (fácil rollback)
4. **Teste migrations em dev** antes de aplicar em prod
5. **Use Docker** para consistência (evita "works on my machine")
6. **Logs são seus amigos** (`docker compose logs -f`)
7. **Backup antes de resetar** (`docker compose exec db pg_dump`)
