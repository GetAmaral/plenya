# Plenya - Sistema de Prontuário Médico Eletrônico (EMR)

## Visão Geral

Sistema EMR completo com web app, mobile apps e backend Go. Foco em segurança LGPD e baixo custo.

**Escala:** Centenas de pacientes
**Plataformas:** Web (Next.js) + Mobile (iOS/Android via Expo)
**Infraestrutura:** Self-hosted Hetzner VPS + Coolify

---

## Stack Técnica

**Frontend**
- Web: Next.js 15.1 + React 19
- Mobile: React Native 0.77 + Expo SDK 56
- UI: Tailwind CSS + shadcn/ui
- Charts: Recharts + Tremor Raw
- Forms: React Hook Form + Zod (gerado automaticamente)
- State: TanStack Query + Zustand

**Backend**
- Language: Go 1.23
- Framework: Fiber v2.53
- ORM: GORM v1.25
- Migrations: Atlas (gerado dos Go models)
- Auth: JWT (golang-jwt/jwt v5)
- Validation: go-playground/validator v10

**Database**
- SGBD: PostgreSQL 17
- Extensions: pgcrypto (criptografia), uuid-ossp

**Infraestrutura**
- Monorepo: Turborepo 2.3 + pnpm 9.15
- Containers: Docker 27
- VPS: Hetzner CPX31 (4 vCPU, 8GB, ~R$50/mês)
- PaaS: Coolify 4.0
- CDN/DNS: Cloudflare Free

---

## Arquitetura: Única Fonte de Verdade

**CRÍTICO:** Go models são a única fonte de verdade. Tudo é gerado automaticamente.

```
apps/api/internal/models/*.go  ← ÚNICA FONTE
         │
         ├─→ Atlas → SQL migrations
         ├─→ Swag → OpenAPI spec
         │    ├─→ TypeScript types (gerado)
         │    └─→ Zod schemas (gerado)
         │
         └─→ pnpm generate
```

**Nunca editar:**
- `apps/api/database/migrations/*.sql` (gerado)
- `packages/types/src/generated/*.ts` (gerado)

**Sempre editar:**
- `apps/api/internal/models/*.go` (fonte única)

### Workflow de Mudanças

```bash
# 1. Editar APENAS o Go model
vim apps/api/internal/models/patient.go

# 2. Gerar tudo automaticamente
pnpm generate

# 3. Revisar o que foi gerado
git diff
```

---

## Estrutura do Monorepo

```
plenya/
├── apps/
│   ├── web/              # Next.js 15.1
│   ├── mobile/           # Expo SDK 56
│   └── api/              # Go backend
│       ├── cmd/server/
│       ├── internal/
│       │   ├── handlers/
│       │   ├── models/        ← ÚNICA FONTE DE VERDADE
│       │   ├── repository/
│       │   ├── services/
│       │   └── middleware/
│       ├── database/
│       │   └── migrations/    ← GERADO (não editar)
│       └── docs/
│           └── swagger.json   ← GERADO (não editar)
├── packages/
│   ├── types/
│   │   └── src/
│   │       └── generated/     ← GERADO (não editar)
│   ├── ui/
│   └── api-client/
└── ARQUITETURA.md             # Documentação técnica
```

---

## Convenções de Código

### Go Models (Única Fonte)

```go
// apps/api/internal/models/patient.go

type Patient struct {
    // Annotations para OpenAPI
    // @example 550e8400-e29b-41d4-a716-446655440000
    ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

    // @minLength 3
    // @maxLength 200
    Name string `gorm:"type:varchar(200);not null" json:"name" validate:"required,min=3,max=200"`

    // @pattern ^\d{3}\.\d{3}\.\d{3}-\d{2}$
    CPF string `gorm:"type:text;not null;unique" json:"-" validate:"required,cpf"`

    // @enum male,female,other
    Gender string `gorm:"type:varchar(10);not null;check:gender IN ('male','female','other')" json:"gender"`

    Phone *string `gorm:"type:varchar(20)" json:"phone,omitempty"` // Nullable

    CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete
}
```

**Tags essenciais:**
- `gorm:"type:varchar(200)"` → Define tipo SQL
- `gorm:"not null"` → Constraint NOT NULL
- `gorm:"unique"` → UNIQUE constraint
- `gorm:"index"` → CREATE INDEX
- `json:"-"` → Nunca expor no JSON (segurança)
- `json:"field,omitempty"` → Omitir se vazio
- `validate:"required,min=3"` → Validação server-side

**Annotations Swagger:**
- `// @example valor` → Exemplo na doc
- `// @minLength 3` → Validação mínima
- `// @enum a,b,c` → Enum values
- `// @items.type string` → Tipo de array

---

## Segurança EMR/LGPD (CRÍTICO)

**Obrigatório implementar:**

**Autenticação:**
- JWT access token: 15min
- JWT refresh token: 7 dias
- Senhas: Argon2id ou Bcrypt (cost 12+)
- 2FA/TOTP obrigatório para profissionais
- Biometria em mobile apps

**Autorização:**
- RBAC: admin, doctor, nurse, patient
- Patients só acessam próprios dados
- Doctors acessam pacientes atribuídos
- Admins têm acesso total (com audit log)

**Dados Sensíveis:**
- CPF, RG: criptografados com pgcrypto
- `json:"-"` em campos sensíveis (nunca expor)
- Mascaramento em logs
- Backups criptografados

**Audit Logging:**
- TODOS os acessos a prontuários geram log
- Campos: user_id, action, resource, resource_id, ip_address, created_at
- Retenção: mínimo 5 anos
- Logs imutáveis (append-only)

**Rate Limiting:**
- 100 requests/min por IP
- Alertas de brute force

**Implementação via Hooks:**
```go
func (p *Patient) BeforeSave(tx *gorm.DB) error {
    // Criptografa CPF antes de salvar
    if p.CPF != "" {
        encrypted, _ := crypto.Encrypt(p.CPF)
        p.CPF = encrypted
    }
    return nil
}

func (p *Patient) AfterFind(tx *gorm.DB) error {
    // Descriptografa após buscar
    if p.CPF != "" {
        decrypted, _ := crypto.Decrypt(p.CPF)
        p.CPF = decrypted
    }
    return nil
}
```

---

## Estrutura de Dados

**Entidades Principais:**

```
Users (base para todos)
  ├─ id, email, password_hash, role, 2fa_enabled
  │
  ├─ Patients (extends Users)
  │    └─ cpf_encrypted, name, birth_date, gender
  │
  ├─ Doctors (extends Users)
  │    └─ crm, specialty
  │
  └─ Nurses (extends Users)

Patients (1:N)
  ├─ Anamnesis (histórico médico)
  ├─ LabResults (exames laboratoriais)
  └─ HealthScores (scores calculados)

AuditLogs (rastreabilidade)
  └─ user_id, action, resource, resource_id, ip, created_at
```

**Regras de Negócio:**

- Health scores: low (0-30), medium (31-60), high (61-100)
- Recalculados quando novos exames adicionados
- Relatórios PDF com timestamp + hash SHA-256
- Soft delete em todas entidades principais

---

## Comandos Essenciais

### Desenvolvimento

```bash
# Instalar dependências
pnpm install

# Gerar código (migrations, types, schemas)
pnpm generate

# Desenvolvimento
pnpm dev                    # Todos os apps
pnpm dev --filter web       # Apenas web
pnpm dev --filter mobile    # Apenas mobile

# Database
docker compose up -d        # Iniciar PostgreSQL 17
docker compose logs -f db   # Ver logs
```

### Backend (após Fase 2)

```bash
cd apps/api

# Rodar servidor
go run cmd/server/main.go

# Build
go build -o bin/server cmd/server/main.go

# Testes
go test ./...
```

### Migrations

```bash
# Gerar migration do Go model
pnpm generate:migrations

# Aplicar migrations
atlas migrate apply --env dev

# Status
atlas migrate status --env dev
```

---

## Variáveis de Ambiente

```bash
# Backend (apps/api/.env)
PORT=3001
ENVIRONMENT=development
DB_HOST=localhost
DB_PORT=5432
DB_USER=plenya_user
DB_PASSWORD=senha_segura
DB_NAME=plenya_db
JWT_SECRET=jwt_secret_256bits
JWT_ACCESS_EXPIRY=15m
JWT_REFRESH_EXPIRY=168h
ENCRYPTION_KEY=encryption_key_256bits
CORS_ORIGIN=http://localhost:3000

# Frontend (apps/web/.env.local)
NEXT_PUBLIC_API_URL=http://localhost:3001
```

---

## Endpoints API (Principais)

### Auth
- `POST /api/v1/auth/register` - Criar conta
- `POST /api/v1/auth/login` - Login
- `POST /api/v1/auth/refresh` - Renovar token
- `POST /api/v1/auth/2fa/enable` - Habilitar 2FA

### Patients
- `GET /api/v1/patients` - Listar (paginado)
- `GET /api/v1/patients/:id` - Detalhes
- `POST /api/v1/patients` - Criar
- `PUT /api/v1/patients/:id` - Atualizar
- `DELETE /api/v1/patients/:id` - Soft delete

### Anamnesis
- `GET /api/v1/patients/:id/anamnesis` - Histórico
- `POST /api/v1/anamnesis` - Criar anamnese

### Labs
- `GET /api/v1/patients/:id/labs` - Exames
- `POST /api/v1/labs` - Registrar exame

### Audit
- `GET /api/v1/audit-logs` - Listar (admin only)

---

## Performance Targets

- Latência p95: < 500ms (rotas simples), < 2s (PDFs)
- Error rate: < 0.5%
- Uptime: > 99.5%
- Database queries: < 100ms (p95)

---

## Custos Mensais

| Item | Custo |
|------|-------|
| VPS Hetzner CPX31 | €9 (~R$50) |
| Backup Snapshots | €2 (~R$11) |
| Backup S3 Glacier | ~R$5 |
| Domínio .com.br | ~R$3 |
| **Total** | **~R$69/mês** |

Opcionais:
- Expo EAS Updates: $99 USD
- App Stores: ~R$50/mês

---

## Roadmap

- [x] **Fase 1:** Fundação (monorepo + Docker) ✅
- [x] **Fase 2:** Backend Core (Auth + RBAC + Migrations) ✅
- [x] **Fase 3:** Features Médicas (Anamnesis + Appointments + Prescriptions + Lab Results) ✅
- [x] **Fase 4 Parte 1:** Frontend Web - Fundação (Next.js 15.1 + shadcn/ui + TanStack Query + Login) ✅
- [x] **Fase 4 Parte 2:** Frontend Web - Dashboard e CRUD Completo ✅
- [ ] **Fase 5:** Mobile Apps (Expo + React Native)
- [ ] **Fase 6:** Hardening LGPD (Relatórios + Exportação)
- [ ] **Fase 7:** Deploy Produção (Hetzner + Coolify)

---

## Observações Importantes

1. **Nunca commitar .env** - Usar .env.example
2. **Audit logs são imutáveis** - Nunca deletar
3. **Criptografia obrigatória** - CPF, RG, documentos
4. **2FA obrigatório** - Profissionais de saúde
5. **Go models são única fonte** - Nunca editar arquivos gerados
6. **Testar geração** - Rodar `pnpm generate` após mudar models
7. **LGPD é lei** - Compliance não opcional

---

**Última atualização:** Janeiro 2026
