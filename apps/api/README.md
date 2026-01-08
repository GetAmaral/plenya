# Plenya API - Backend Go

Backend do sistema EMR Plenya construído com Go + Fiber + GORM.

## Requisitos

- Go 1.23+
- PostgreSQL 17 (via Docker Compose na raiz do projeto)

## Instalação

```bash
# 1. Instalar dependências
go mod tidy

# 2. Configurar variáveis de ambiente
cp .env.example .env
# Edite .env se necessário (valores padrão já funcionam com docker-compose)

# 3. Garantir que PostgreSQL está rodando
cd ../.. && docker compose up -d
```

## Desenvolvimento

```bash
# Rodar servidor (com hot reload via air - opcional)
go run cmd/server/main.go

# Ou build e executar
go build -o bin/server cmd/server/main.go
./bin/server
```

## Estrutura

```
apps/api/
├── cmd/
│   └── server/          # Entry point (main.go)
├── internal/
│   ├── config/          # Configuração e env vars
│   ├── database/        # Conexão GORM
│   ├── models/          # GORM models (ÚNICA FONTE DE VERDADE)
│   ├── handlers/        # HTTP handlers (controllers)
│   ├── services/        # Lógica de negócio
│   ├── repository/      # Data access layer
│   └── middleware/      # Auth, RBAC, Audit, etc
├── database/
│   └── migrations/      # Atlas migrations (GERADAS automaticamente)
├── docs/                # Swagger docs (GERADO automaticamente)
└── pkg/
    ├── crypto/          # Criptografia (CPF, RG)
    └── utils/           # Utilidades
```

## Endpoints

### Health Check
```
GET /health
```

### API v1
```
GET /api/v1
```

## Auto Migration (Desenvolvimento)

Em desenvolvimento, o GORM roda auto-migration na inicialização:

```go
// Cria tabelas: users, patients, audit_logs
database.AutoMigrate()
```

**PRODUÇÃO:** Usaremos Atlas para migrations versionadas.

## Próximos Passos (Fase 2)

- [ ] Implementar autenticação JWT
- [ ] Implementar RBAC middleware
- [ ] Implementar audit logging middleware
- [ ] Criar handlers CRUD para patients
- [ ] Configurar Atlas para migrations
- [ ] Configurar Swagger/Swag
- [ ] Gerar TypeScript types e Zod schemas

## Variáveis de Ambiente

Ver `.env.example` para lista completa.

Principais:
- `PORT`: Porta do servidor (padrão: 3001)
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`: PostgreSQL
- `JWT_SECRET`: Secret para assinar JWTs
- `ENCRYPTION_KEY`: Chave para criptografar dados sensíveis
