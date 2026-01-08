# Plenya EMR - Arquitetura Técnica

## Stack

**Frontend:** Next.js 15.1 + React 19 (web) | React Native 0.77 + Expo 56 (mobile)
**Backend:** Go 1.23 + Fiber v2.53 + GORM
**Database:** PostgreSQL 17
**Monorepo:** Turborepo 2.3 + pnpm 9.15
**Infra:** Docker 27, Hetzner VPS, Coolify 4.0

---

## Única Fonte de Verdade: Go Models

**Tudo é gerado automaticamente a partir dos Go models:**

```
apps/api/internal/models/*.go  ← VOCÊ EDITA APENAS AQUI
         │
         ├─→ Atlas → SQL migrations
         ├─→ Swag → OpenAPI spec
         │    ├─→ openapi-typescript → TypeScript types
         │    └─→ openapi-zod → Zod schemas
         │
         └─→ pnpm generate (roda tudo)
```

### Ferramentas

- **Atlas:** Gera migrations SQL dos Go models
- **Swag:** Gera OpenAPI spec dos comentários Go
- **openapi-typescript:** Gera TypeScript types do OpenAPI
- **openapi-zod-client:** Gera Zod schemas do OpenAPI

### Exemplo de Model

```go
// apps/api/internal/models/patient.go
package models

type Patient struct {
    // @example 550e8400-e29b-41d4-a716-446655440000
    ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

    // @minLength 3
    // @maxLength 200
    Name string `gorm:"type:varchar(200);not null" json:"name" validate:"required,min=3,max=200"`

    // @pattern ^\d{3}\.\d{3}\.\d{3}-\d{2}$
    CPF string `gorm:"type:text;not null;unique" json:"-" validate:"required,cpf"`

    BirthDate time.Time `gorm:"type:date;not null" json:"birthDate"`

    // @enum male,female,other
    Gender string `gorm:"type:varchar(10);not null;check:gender IN ('male','female','other')" json:"gender" validate:"required,oneof=male female other"`

    Phone *string `gorm:"type:varchar(20)" json:"phone,omitempty"`

    CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
```

### Geração Automática

```bash
# Editar model
vim apps/api/internal/models/patient.go

# Gerar tudo
pnpm generate

# Resultado:
✅ apps/api/database/migrations/*.sql         # SQL gerado
✅ apps/api/docs/swagger.json                 # OpenAPI
✅ packages/types/src/generated/api-types.ts  # TypeScript
✅ packages/types/src/generated/api-schemas.ts # Zod
```

---

## Mapeamento de Tipos

| Go | PostgreSQL | TypeScript | Nota |
|----|------------|------------|------|
| `string` | `VARCHAR` | `string` | ✅ |
| `int`, `int64` | `INTEGER`, `BIGINT` | `number` | ✅ |
| `bool` | `BOOLEAN` | `boolean` | ✅ |
| `uuid.UUID` | `UUID` | `string` | ✅ |
| `time.Time` | `TIMESTAMP` | `string` (ISO) | ✅ |
| `*string` | `NULL` | `string?` | ✅ |
| Enum (string) | `CHECK` | `"a"\|"b"` | ✅ Precisa `@enum` |
| `[]string` | `TEXT[]` | `string[]` | ✅ Precisa `@items.type` |
| Struct | `JSONB` | `Type` | ✅ Define struct separado |
| Money | `BIGINT` | `number` | ✅ Usar centavos |
| Arquivos | `VARCHAR` (URL) | `string` | ✅ Usar S3/storage |

### Tags GORM Importantes

```go
`gorm:"type:varchar(200)"`           // Tipo SQL
`gorm:"not null"`                    // NOT NULL
`gorm:"unique"`                      // UNIQUE
`gorm:"index"`                       // CREATE INDEX
`gorm:"primaryKey"`                  // PRIMARY KEY
`gorm:"default:gen_random_uuid()"`   // DEFAULT
`gorm:"check:..."`                   // CHECK constraint
`gorm:"foreignKey:UserID"`           // FK
`gorm:"constraint:OnDelete:CASCADE"` // CASCADE
`gorm:"autoCreateTime"`              // Auto-fill created_at
`gorm:"autoUpdateTime"`              // Auto-fill updated_at
```

### Annotations Swagger

```go
// @Description Descrição do campo
// @example valor
// @minLength 3
// @maxLength 200
// @pattern ^regex$
// @enum value1,value2,value3
// @items.type string (para arrays)
```

---

## Fluxo de Dados

### Criação (Frontend → Backend → Database)

```
Frontend (Zod valida form)
   ↓ POST JSON
Backend (Go valida + processa)
   ↓ BeforeSave hook (criptografa CPF)
   ↓ GORM INSERT
PostgreSQL (salva dados)
```

### Leitura (Database → Backend → Frontend)

```
PostgreSQL
   ↓ SELECT
Backend (GORM)
   ↓ AfterFind hook (descriptografa)
   ↓ JSON serialization
Frontend (renderiza)
```

---

## Segurança EMR

**Obrigatório:**
- Criptografia de CPF/RG (pgcrypto)
- Audit logs (todos os acessos)
- 2FA para profissionais
- RBAC (admin, doctor, nurse, patient)
- Rate limiting (100 req/min)
- HTTPS/TLS 1.3

**Implementação:**
- BeforeSave hook → criptografa dados sensíveis
- Middleware → registra audit log automático
- JWT: access 15min, refresh 7 dias
- Campos sensíveis: `json:"-"` (nunca expor)

---

## Estrutura de Dados Principal

```
Users (base)
  ├─ Patients (extends Users)
  ├─ Doctors (extends Users)
  └─ Nurses (extends Users)

Patients
  ├─ Anamnesis (histórico médico)
  ├─ LabResults (exames)
  └─ HealthScores (scores calculados)

AuditLogs (rastreabilidade)
```

---

## Comandos

```bash
# Gerar tudo
pnpm generate

# Desenvolvimento
pnpm dev                  # Todos os apps
pnpm dev --filter web     # Apenas web
pnpm dev --filter mobile  # Apenas mobile

# Database
docker compose up -d      # Iniciar PostgreSQL

# Backend Go (após Fase 2)
cd apps/api
go run cmd/server/main.go
```

---

## Custos

- VPS Hetzner CPX31: €9 (~R$50)
- Backup Snapshots: €2 (~R$11)
- Backup S3 Glacier: ~R$5
- Domínio: ~R$3
- **Total: ~R$69/mês**

---

## Roadmap

- [x] **Fase 1:** Monorepo + Docker + Schemas ✅
- [ ] **Fase 2:** Backend Go + Auth + RBAC + Migrations
- [ ] **Fase 3:** Features Médicas (CRUD + Anamnese + Labs)
- [ ] **Fase 4:** Frontend Web
- [ ] **Fase 5:** Mobile Apps
- [ ] **Fase 6:** Hardening LGPD
- [ ] **Fase 7:** Deploy

---

**Última atualização:** Janeiro 2026
