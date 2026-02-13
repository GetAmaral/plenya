# Plenya - Sistema de Prontuário Médico Eletrônico (EMR)

Sistema EMR completo com Go backend + Next.js frontend + mobile apps React Native.

## 🚨 Regras de Ouro (CRÍTICAS)

### 1. Fonte Única de Verdade
**Go models são a ÚNICA fonte.** Nunca editar arquivos gerados.

```
apps/api/internal/models/*.go  ← EDITAR AQUI
         │
         ├─→ Atlas → migrations/*.sql (gerado)
         ├─→ Swag → swagger.json (gerado)
         └─→ packages/types/generated/*.ts (gerado)
```

### 2. Desenvolvimento vs Produção

| Contexto | Método | Quando |
|----------|--------|--------|
| **Desenvolvimento (você, Claude)** | Docker exec → psql | Manipulação manual de dados |
| **Desenvolvimento (você, Claude)** | Go scripts | Operações complexas |
| **Produção (apps)** | API HTTP | Web/mobile usam API |

**❌ NUNCA usar API HTTP (curl/POST/PUT) para desenvolvimento manual**
**✅ SEMPRE usar banco direto via Docker ou Go scripts**

### 3. Hooks Obrigatórios

#### Backend (Go)
- **TODOS os models:** `BeforeCreate` hook para UUID v7
- **Patient:** `BeforeSave`/`AfterFind` para criptografia (CPF, RG)
- **ScoreItem, ScoreLevel:** `BeforeUpdate` para auto-atualizar `LastReview`

#### Frontend (React)
- **TODOS os formulários:** `useFormNavigation({ formRef })`
- **TODAS as páginas de dados de paciente:** `useRequireSelectedPatient()`

## 📚 Documentação Completa

**Leia `.claude/` para detalhes completos.** Abaixo apenas índice.

### Fundação
- [**01-overview.md**](.claude/01-overview.md) - Visão geral, escala, objetivos
- [**02-stack.md**](.claude/02-stack.md) - Stack técnica e versões
- [**03-architecture.md**](.claude/03-architecture.md) - Arquitetura e geração automática

### Backend
- [**models.md**](.claude/backend/models.md) - Go models: tags GORM, validação, Swagger
- [**🔥 database.md**](.claude/backend/database.md) - **COMO TRABALHAR COM BANCO DIRETO**
- [**hooks.md**](.claude/backend/hooks.md) - GORM lifecycle hooks
- [**service-layer.md**](.claude/backend/service-layer.md) - DTOs, business logic
- [**api-endpoints.md**](.claude/backend/api-endpoints.md) - Quando usar API HTTP

### Frontend
- [**form-navigation.md**](.claude/frontend/form-navigation.md) - useFormNavigation (obrigatório)
- [**patient-context.md**](.claude/frontend/patient-context.md) - useRequireSelectedPatient
- [**tanstack-query.md**](.claude/frontend/tanstack-query.md) - Query patterns, invalidação

### Domínio
- [**🎯 score-system.md**](.claude/domain/score-system.md) - **SISTEMA DE ESCORES (hierarquia, manipulação)**
- [**patients.md**](.claude/domain/patients.md) - Workflows de pacientes
- [**security.md**](.claude/domain/security.md) - LGPD, criptografia, audit

### Workflows
- [**development.md**](.claude/workflows/development.md) - Como desenvolver
- [**🔥 database-ops.md**](.claude/workflows/database-ops.md) - **OPERAÇÕES DIRETAS NO BANCO**
- [**adding-features.md**](.claude/workflows/adding-features.md) - Adicionar features

## 🛠 Comandos Essenciais

### Acessar Banco Direto (DESENVOLVIMENTO)

```bash
# Acesso interativo
docker compose exec db psql -U plenya_user -d plenya_db

# Executar SQL direto
docker compose exec db psql -U plenya_user -d plenya_db -c "SELECT * FROM score_items;"

# Ver estrutura de tabela
docker compose exec db psql -U plenya_user -d plenya_db -c "\d score_items"
```

### Geração Automática

```bash
# Após editar Go models
pnpm generate  # Gera: migrations, OpenAPI, TypeScript types, Zod schemas
```

### Docker

```bash
# Iniciar tudo
docker compose up -d

# Ver logs
docker compose logs -f api
docker compose logs -f web

# Rebuild após mudar dependências
docker compose up -d --build
```

## 📖 Leitura Obrigatória Por Contexto

### Vou manipular dados manualmente (adicionar/editar score items, etc.)
1. 🔥 [Database Operations](.claude/workflows/database-ops.md)
2. 🎯 [Sistema de Escores](.claude/domain/score-system.md)

### Vou adicionar um novo model/feature
1. [Architecture](.claude/03-architecture.md)
2. [Models](.claude/backend/models.md)
3. [Adding Features](.claude/workflows/adding-features.md)

### Vou trabalhar no frontend
1. [Form Navigation](.claude/frontend/form-navigation.md)
2. [Patient Context](.claude/frontend/patient-context.md)
3. [TanStack Query](.claude/frontend/tanstack-query.md)

### Vou trabalhar com segurança/LGPD
1. [Security](.claude/domain/security.md)
2. [Hooks](.claude/backend/hooks.md) (criptografia)

## 🏗 Estrutura do Monorepo

```
plenya/
├── CLAUDE.md                    # Este arquivo (índice)
├── .claude/                     # Documentação detalhada
│   ├── 01-overview.md
│   ├── 02-stack.md
│   ├── 03-architecture.md
│   ├── backend/
│   │   ├── models.md
│   │   ├── database.md          # ⭐ IMPORTANTE
│   │   ├── hooks.md
│   │   ├── service-layer.md
│   │   └── api-endpoints.md
│   ├── frontend/
│   │   ├── form-navigation.md
│   │   ├── patient-context.md
│   │   └── tanstack-query.md
│   ├── domain/
│   │   ├── score-system.md      # ⭐ IMPORTANTE
│   │   ├── patients.md
│   │   └── security.md
│   └── workflows/
│       ├── development.md
│       ├── database-ops.md      # ⭐ IMPORTANTE
│       └── adding-features.md
├── apps/
│   ├── api/                     # Go backend
│   │   ├── internal/
│   │   │   └── models/          ← ÚNICA FONTE DE VERDADE
│   │   └── database/
│   │       └── migrations/      ← GERADO (não editar)
│   ├── web/                     # Next.js 16.1
│   └── mobile/                  # Expo SDK 56
└── packages/
    ├── types/
    │   └── src/generated/       ← GERADO (não editar)
    └── ui/
```

## 🎯 Stack Resumida

- **Backend:** Go 1.25 + Fiber v2 + GORM v1.25 + PostgreSQL 17
- **Frontend:** Next.js 16.1 + React 19.2 + shadcn/ui + TanStack Query
- **Mobile:** React Native 0.77 + Expo SDK 56
- **Infra:** Docker 27 + Turborepo 2.7 + Hetzner VPS + Coolify 4.0

## 🔐 Segurança LGPD

- CPF/RG: Criptografados via hooks (BeforeSave/AfterFind)
- Audit logs: Imutáveis, retenção 5 anos
- JWT: Access 15min, Refresh 7 dias
- 2FA obrigatório para profissionais

## 📊 Sistema de Escores (Core Feature)

Hierarquia de 4 níveis para estratificação de risco:

```
ScoreGroup → ScoreSubgroup → ScoreItem → ScoreLevel
```

- **Filtros demográficos:** gender, age range, post-menopause
- **Operadores:** =, >, >=, <, <=, between
- **Enriquecimento clínico:** clinical_relevance, patient_explanation, conduct

Ver [score-system.md](.claude/domain/score-system.md) para detalhes completos.

## 🚀 Roadmap

- [x] Fase 1-3: Backend core + Auth + RBAC + Migrations
- [x] Fase 4: Frontend web + Dashboard + Sistema de Escores
- [ ] Fase 5: Mobile apps (Expo)
- [ ] Fase 6: Hardening LGPD
- [ ] Fase 7: Deploy produção (Hetzner + Coolify)

---

**Última atualização:** Fevereiro 2026
**Versão:** 3.0 (Documentação Modular)

Para detalhes técnicos, sempre consulte arquivos em `.claude/`.
