# 02 - Stack Técnica Completa

## Frontend

### Web (Next.js)
- **Framework:** Next.js 16.1 (App Router)
- **React:** 19.2 (React Compiler habilitado)
- **TypeScript:** 5.9.3
- **UI Library:** shadcn/ui (componentes Radix UI)
- **Styling:** Tailwind CSS 4.0
- **Forms:** React Hook Form + Zod validation
- **State Management:**
  - TanStack Query v5 (server state)
  - Zustand 5.0 (client state)
- **Charts:** Recharts + Tremor Raw
- **Icons:** Lucide React
- **Date:** date-fns
- **Build:** Turbopack (dev), Webpack (prod)

### Mobile (React Native)
- **Framework:** Expo SDK 56
- **React Native:** 0.77
- **Navigation:** Expo Router (file-based)
- **UI:** React Native Paper + custom components
- **Forms:** React Hook Form + Zod
- **State:** TanStack Query + Zustand (mesma estrutura web)
- **Storage:** Expo SecureStore (tokens)
- **Camera:** expo-camera
- **File System:** expo-file-system
- **Updates:** EAS Updates (over-the-air)

## Backend

### Language & Framework
- **Language:** Go 1.25
- **Framework:** Fiber v2.52.10 (Express-like para Go)
- **Router:** Fiber native
- **Middleware:** Fiber middleware ecosystem

### Database & ORM
- **Database:** PostgreSQL 17
- **ORM:** GORM v1.25
- **Migrations:** Atlas (schema-as-code)
- **Extensions PostgreSQL:**
  - pgcrypto (criptografia)
  - uuid-ossp (UUID generation)
  - pg_stat_statements (query performance)

### Authentication & Security
- **JWT:** golang-jwt/jwt v5
- **Password Hashing:** bcrypt (cost 12) ou Argon2id
- **Encryption:** AES-256-GCM (via pgcrypto)
- **CORS:** Fiber CORS middleware
- **Rate Limiting:** Fiber Rate Limiter
- **2FA:** TOTP (Time-based One-Time Password)

### Validation & Documentation
- **Validation:** go-playground/validator v10
- **API Docs:** Swag (Swagger/OpenAPI generator)
- **OpenAPI:** 3.0 spec

### Utilities
- **UUID:** google/uuid (v7 support)
- **Config:** viper
- **Logging:** zerolog ou logrus
- **Testing:** testify

## Infrastructure

### Monorepo
- **Tool:** Turborepo 2.7.5
- **Package Manager:** pnpm 10.28.1
- **Workspaces:** pnpm workspaces
- **Node.js:** 24 LTS (Krypton)

### Containerization
- **Engine:** Docker 27
- **Compose:** Docker Compose v2
- **Base Images:**
  - Go: golang:1.25-alpine
  - Node: node:24-alpine
  - PostgreSQL: postgres:17-alpine

### Development Tools
- **Go:**
  - Air (hot reload)
  - golangci-lint (linting)
  - gofmt (formatting)
- **TypeScript/JavaScript:**
  - ESLint 9
  - Prettier 3
  - TypeScript 5.9

### Build Tools
- **Go:** Native go build
- **Next.js:** Turbopack (dev), Webpack (prod)
- **Mobile:** Expo EAS Build

## Production Infrastructure

### Hosting
- **VPS:** Hetzner Cloud CPX31
  - 4 vCPU AMD
  - 8 GB RAM
  - 160 GB NVMe SSD
  - 20 TB traffic
  - €9/mês (~R$50)

### PaaS
- **Platform:** Coolify 4.0 (self-hosted)
  - Docker-based deployment
  - Git integration
  - SSL/TLS automático (Let's Encrypt)
  - Backups automatizados
  - Monitoring integrado

### CDN & DNS
- **Provider:** Cloudflare Free
  - DNS management
  - DDoS protection
  - SSL/TLS
  - Caching rules
  - Page Rules

### Database Production
- **Primary:** PostgreSQL 17
- **Backups:**
  - Daily automated (Coolify)
  - Weekly snapshots (Hetzner)
  - Monthly S3 Glacier (off-site)
- **High Availability:** Futuro (Patroni + HAProxy)

### Monitoring & Observability
- **Logs:** Loki + Promtail
- **Metrics:** Prometheus + Grafana
- **Uptime:** UptimeRobot (free tier)
- **Error Tracking:** Sentry (futuro)

### CI/CD
- **Platform:** GitHub Actions
- **Workflows:**
  - Lint & Test (on PR)
  - Build & Deploy (on push to main)
  - Database migrations (on tag)
- **Environments:**
  - Development (local Docker)
  - Staging (Coolify staging)
  - Production (Coolify prod)

## Package Structure

### Monorepo Workspaces

```json
{
  "name": "plenya",
  "private": true,
  "workspaces": [
    "apps/*",
    "packages/*"
  ]
}
```

### Apps

```
apps/
├── api/              # Go backend
├── web/              # Next.js web app
└── mobile/           # Expo mobile app
```

### Packages (Shared)

```
packages/
├── types/            # TypeScript types (gerados do OpenAPI)
├── ui/               # Componentes UI compartilhados
├── api-client/       # HTTP client wrapper
├── config/           # Shared configs (ESLint, TypeScript, etc.)
└── utils/            # Utility functions
```

## Versões & Compatibilidade

### Runtime
- **Node.js:** >= 24.0.0 LTS
- **Go:** >= 1.25
- **PostgreSQL:** >= 17
- **Docker:** >= 27.0

### Browsers (Web)
- Chrome/Edge: últimas 2 versões
- Firefox: últimas 2 versões
- Safari: >= 16

### Mobile
- **iOS:** >= 15.0
- **Android:** >= 13 (API Level 33)

## Package.json Scripts

### Root

```json
{
  "scripts": {
    "dev": "turbo run dev",
    "build": "turbo run build",
    "lint": "turbo run lint",
    "format": "prettier --write \"**/*.{ts,tsx,js,jsx,json,md}\"",
    "generate": "pnpm generate:migrations && pnpm generate:swagger && pnpm generate:types",
    "generate:migrations": "cd apps/api && atlas migrate diff --env dev",
    "generate:swagger": "cd apps/api && swag init -g cmd/server/main.go -o docs",
    "generate:types": "cd packages/types && pnpm run generate"
  }
}
```

### apps/web

```json
{
  "scripts": {
    "dev": "next dev",
    "build": "next build",
    "start": "next start",
    "lint": "next lint",
    "type-check": "tsc --noEmit"
  }
}
```

### apps/mobile

```json
{
  "scripts": {
    "start": "expo start",
    "android": "expo run:android",
    "ios": "expo run:ios",
    "lint": "eslint .",
    "type-check": "tsc --noEmit"
  }
}
```

### apps/api

```json
{
  "scripts": {
    "dev": "air",
    "build": "go build -o bin/server cmd/server/main.go",
    "test": "go test ./...",
    "lint": "golangci-lint run"
  }
}
```

## Environment Variables

### Development

**apps/api/.env:**
```bash
PORT=3001
ENVIRONMENT=development
DB_HOST=localhost
DB_PORT=5432
DB_USER=plenya_user
DB_PASSWORD=senha_segura
DB_NAME=plenya_db
JWT_SECRET=jwt_secret_256bits_dev
JWT_ACCESS_EXPIRY=15m
JWT_REFRESH_EXPIRY=168h
ENCRYPTION_KEY=encryption_key_256bits_dev
CORS_ORIGIN=http://localhost:3000
```

**apps/web/.env.local:**
```bash
NEXT_PUBLIC_API_URL=http://localhost:3001
```

**apps/mobile/.env:**
```bash
EXPO_PUBLIC_API_URL=http://localhost:3001
```

### Production

**apps/api/.env.production:**
```bash
PORT=3001
ENVIRONMENT=production
DB_HOST=db.internal
DB_PORT=5432
DB_USER=plenya_prod
DB_PASSWORD=<secret-from-env>
DB_NAME=plenya_prod
JWT_SECRET=<secret-from-env>
JWT_ACCESS_EXPIRY=15m
JWT_REFRESH_EXPIRY=168h
ENCRYPTION_KEY=<secret-from-env>
CORS_ORIGIN=https://app.plenya.com.br
SENTRY_DSN=<sentry-dsn>
```

## Dependencies Principais

### Backend (go.mod)

```go
module github.com/plenya/api

go 1.25

require (
    github.com/gofiber/fiber/v2 v2.52.10
    github.com/google/uuid v1.6.0
    gorm.io/gorm v1.25.12
    gorm.io/driver/postgres v1.5.9
    github.com/golang-jwt/jwt/v5 v5.2.1
    github.com/go-playground/validator/v10 v10.22.1
    github.com/swaggo/swag v1.16.3
    github.com/spf13/viper v1.19.0
    golang.org/x/crypto v0.28.0
)
```

### Frontend (package.json)

```json
{
  "dependencies": {
    "next": "16.1.0",
    "react": "19.2.0",
    "react-dom": "19.2.0",
    "@tanstack/react-query": "^5.62.11",
    "zustand": "^5.0.2",
    "react-hook-form": "^7.54.2",
    "zod": "^3.24.1",
    "tailwindcss": "^4.0.0",
    "lucide-react": "^0.469.0",
    "date-fns": "^4.1.0",
    "recharts": "^2.15.0"
  },
  "devDependencies": {
    "typescript": "5.9.3",
    "eslint": "^9.17.0",
    "prettier": "^3.4.2",
    "turbo": "2.7.5"
  }
}
```

### Mobile (package.json)

```json
{
  "dependencies": {
    "expo": "~56.0.0",
    "react-native": "0.77.0",
    "expo-router": "~4.0.0",
    "@tanstack/react-query": "^5.62.11",
    "zustand": "^5.0.2",
    "react-hook-form": "^7.54.2",
    "zod": "^3.24.1",
    "expo-secure-store": "~14.0.0"
  }
}
```

## Performance Targets

### Web
- **First Contentful Paint:** < 1.5s
- **Largest Contentful Paint:** < 2.5s
- **Time to Interactive:** < 3.5s
- **Cumulative Layout Shift:** < 0.1
- **Lighthouse Score:** > 90

### API
- **Response Time (p50):** < 100ms
- **Response Time (p95):** < 500ms
- **Response Time (p99):** < 1s
- **Throughput:** > 1000 req/s
- **Error Rate:** < 0.5%

### Database
- **Query Time (p95):** < 100ms
- **Connection Pool:** 10-25 connections
- **Max Connections:** 100

### Mobile
- **App Launch:** < 3s
- **Screen Transitions:** < 300ms
- **API Calls:** < 2s

## Custo Mensal Estimado

| Item | Custo |
|------|-------|
| Hetzner CPX31 VPS | €9 (~R$50) |
| Hetzner Snapshots (2x) | €2 (~R$11) |
| S3 Glacier Backups | ~R$5 |
| Domínio .com.br | ~R$3 |
| Cloudflare | R$0 (free) |
| **Subtotal Base** | **~R$69/mês** |

Opcionais:
- Expo EAS Build/Updates | $99 USD (~R$520)
- Google Play Store | $25 USD one-time
- Apple Developer | $99 USD/ano (~R$520)
- Sentry | $26 USD/mês (~R$136)

**Total com opcionais:** ~R$750/mês

## Justificativas Técnicas

### Por que Next.js 16.1?
- App Router estável
- React Server Components
- Turbopack (build 700% mais rápido)
- React 19 suporte

### Por que Go?
- Performance (compile to native)
- Concurrency (goroutines)
- Small binary size (~20MB)
- Low memory footprint
- Static typing

### Por que PostgreSQL 17?
- ACID compliant
- JSON support (JSONB)
- Full-text search
- Row-level security
- Extensões (pgcrypto)
- Sem custos de licença

### Por que GORM?
- Idiomatic Go
- Auto migrations (via Atlas)
- Preload/Joins
- Hooks (BeforeCreate, AfterFind)
- Type-safe

### Por que Fiber?
- Express-like API (familiar)
- Fastest Go framework
- Middleware ecosystem
- Low memory usage

### Por que Turborepo?
- Incremental builds
- Remote caching
- Task pipelines
- Monorepo otimizado

### Por que Hetzner?
- Melhor custo/benefício Europa
- Data center na Alemanha (GDPR)
- SSD NVMe rápido
- IPv4 incluído
- Suporte excelente

### Por que Coolify?
- Self-hosted (controle total)
- Git-based deployment
- Docker nativo
- SSL automático
- Backups integrados
- Alternativa free ao Vercel/Heroku
