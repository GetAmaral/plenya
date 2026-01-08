# Plenya - Sistema de Prontuário Médico Eletrônico

Sistema completo de prontuário eletrônico com web app, mobile apps (iOS/Android) e backend Go.

## Stack

- **Web:** Next.js 15.1 + React 19
- **Mobile:** React Native 0.77 + Expo SDK 56
- **Backend:** Go 1.23 + Fiber v2 + GORM
- **Database:** PostgreSQL 17
- **Monorepo:** Turborepo + pnpm

## Requisitos

- Docker 27+
- Docker Compose
- Node.js 22 LTS (apenas para frontend/mobile)
- pnpm 9.15+ (apenas para frontend/mobile)

**NÃO precisa instalar Go** - roda em Docker!

## Setup e Execução

```bash
# 1. Subir TUDO (banco + API)
docker compose up -d

# Ver logs
docker compose logs -f

# Parar tudo
docker compose down
```

**Pronto!** A API está rodando em: http://localhost:3001

## Testar

```bash
# Health check
curl http://localhost:3001/health

# Deve retornar:
# {"status":"ok","database":"connected"}

# API info
curl http://localhost:3001/api/v1

# Registrar usuário
curl -X POST http://localhost:3001/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@plenya.com","password":"Test123","role":"doctor"}'

# Login e obter token JWT
curl -X POST http://localhost:3001/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@plenya.com","password":"Test123"}'

# Usar o accessToken retornado nas próximas requisições
curl http://localhost:3001/api/v1/patients \
  -H "Authorization: Bearer SEU_ACCESS_TOKEN"
```

## API Features Implementadas

### Autenticação e Autorização
- ✅ JWT Authentication (access + refresh tokens)
- ✅ RBAC com 4 roles: admin, doctor, nurse, patient
- ✅ Middleware de autenticação e autorização
- ✅ Refresh token rotation

### Recursos Médicos
- ✅ **Patients** - CRUD completo com criptografia AES-256-GCM (CPF/RG)
- ✅ **Anamnesis** - Histórico médico do paciente
- ✅ **Appointments** - Agendamento de consultas com confirmação/cancelamento
- ✅ **Prescriptions** - Prescrições médicas com status tracking
- ✅ **Lab Results** - Resultados de exames laboratoriais

### Segurança e Compliance
- ✅ Audit Logs automáticos (LGPD compliance)
- ✅ Criptografia de dados sensíveis (AES-256-GCM)
- ✅ Soft delete em todos os recursos
- ✅ Validação de inputs com validators
- ✅ Proteção contra SQL injection (prepared statements)

### Total de Endpoints
**64 handlers** implementados e funcionando

## Desenvolvimento Frontend

### Web App (Next.js)

```bash
# Instalar dependências (na raiz do projeto)
pnpm install

# Iniciar web app em modo dev
pnpm dev --filter @plenya/web

# Build production
pnpm build --filter @plenya/web

# A web app roda em: http://localhost:3000
```

**Features Implementadas - Design Moderno Janeiro 2026:**

**Fase 4 Parte 1 - Fundação:**
- ✅ Next.js 15.1 + React 19
- ✅ **Design System Completo** com Design Tokens
- ✅ Tailwind CSS + shadcn/ui
- ✅ **Framer Motion** - Animações fluidas e micro-interactions
- ✅ **Sonner** - Toasts elegantes
- ✅ **Tremor + Recharts** - Data visualization
- ✅ TanStack Query + Zustand
- ✅ React Hook Form + Zod
- ✅ **Login Page Moderna** - Glassmorphism + animated gradients
- ✅ **Bento Grid Dashboard** - Layout moderno com cards responsivos
- ✅ **Componentes UI Médicos** - Badge com status de paciente, Avatar, Skeleton
- ✅ **CSS Utilities** - Glass effects, gradients, scrollbar customizado
- ✅ **Cores Semânticas Médicas** - Estável, Observação, Crítico
- ✅ API client com autenticação

**Fase 4 Parte 2 - Dashboard e CRUD:**
- ✅ **Middleware de Proteção** - Auth hooks e route protection
- ✅ **Sidebar Moderna** - Navegação animada com Framer Motion
- ✅ **Layout Dashboard** - Sidebar fixa com conteúdo responsivo
- ✅ **CRUD de Pacientes** - Listagem completa com TanStack Table
- ✅ **TanStack React Table** - Sorting, filtering, pagination
- ✅ **Skeleton Loading States** - Loading UX moderna
- ✅ **Busca Global** - Filtro em tempo real
- ✅ **Paginação** - Controles next/previous
- ✅ **Formatação de Datas** - date-fns com locale pt-BR
- ✅ **Badge de Gênero** - Indicadores visuais
- ✅ **Integração com API** - React Query para cache e fetching

**Fase 4 Parte 3 - CRUD Completo e Features Avançadas:**
- ✅ **Appointments Page** - Agendamento de consultas com status tracking
- ✅ **Prescriptions Page** - Prescrições médicas com detalhes de medicação
- ✅ **Anamnesis Page** - Histórico médico com alergias e medicamentos
- ✅ **Lab Results Page** - Exames laboratoriais com download de resultados
- ✅ **Command Palette (⌘K)** - Navegação rápida global com cmdk
- ✅ **Charts com Tremor** - AreaChart, BarChart, DonutChart no dashboard
- ✅ **Dialog Component** - Radix UI Dialog com animações
- ✅ **Todas Layouts** - 4 novos layouts com sidebar para cada seção médica

### Mobile App (futuramente)

```bash
# Mobile app
pnpm dev --filter mobile
```

## Estrutura

```
plenya/
├── apps/
│   ├── web/          # Next.js (Fase 4)
│   ├── mobile/       # Expo (Fase 5)
│   └── api/          # Go backend ✅ RODANDO
│       └── internal/models/  ← ÚNICA FONTE DE VERDADE
├── packages/
│   └── types/src/generated/  ← Gerado automaticamente
└── docker-compose.yml        ← Sobe tudo
```

## Geração de Código (quando necessário)

```bash
# Gerar migrations, TypeScript types e Zod schemas
pnpm generate
```

## Logs e Debug

```bash
# Ver logs da API
docker compose logs -f api

# Ver logs do banco
docker compose logs -f db

# Rebuild da API (após mudar código Go)
docker compose up -d --build api
```

## Acesso ao Banco

```bash
# Acessar PostgreSQL
docker exec -it plenya-db psql -U plenya_user -d plenya_db

# Ver tabelas criadas
\dt

# Ver dados de uma tabela
SELECT * FROM users;
```

## Documentação

- **CLAUDE.md** - Contexto completo do projeto (para AI)
- **ARQUITETURA.md** - Documentação técnica
- **apps/api/README.md** - Detalhes do backend

## Roadmap

- [x] Fase 1: Fundação (monorepo + Docker) ✅
- [x] Fase 2 Parte 1: Backend básico (models + servidor) ✅
- [x] Fase 2 Parte 2: Auth JWT + RBAC + CRUD ✅
- [x] Fase 3: Features Médicas (Anamnesis, Appointments, Prescriptions, Lab Results) ✅
- [x] Fase 4 Parte 1: Frontend Web - Fundação (Next.js 15.1 + shadcn/ui + TanStack Query) ✅
- [x] Fase 4 Parte 2: Frontend Web - Dashboard e CRUD ✅
- [x] Fase 4 Parte 3: Frontend Web - CRUD Completo + Command Palette + Charts ✅
- [ ] Fase 5: Mobile Apps (React Native + Expo)
- [ ] Fase 6: Hardening LGPD
- [ ] Fase 7: Deploy

## Troubleshooting

**API não sobe:**
```bash
# Ver erro
docker compose logs api

# Rebuild
docker compose up -d --build api
```

**Banco não conecta:**
```bash
# Verificar se está rodando
docker compose ps

# Reiniciar
docker compose restart db
```

## Licença

Proprietary - Todos os direitos reservados
