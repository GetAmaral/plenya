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

## Sistema RAG (Retrieval Augmented Generation)

### Visão Geral

Sistema de busca semântica por similaridade vetorial para artigos científicos, permitindo:
- **Busca inteligente** por significado (não apenas palavras-chave)
- **Recomendações automáticas** de artigos ao editar ScoreItems
- **Descoberta bidirecional** - quais artigos cobrem quais parâmetros clínicos

### Stack Técnica RAG

| Componente | Tecnologia | Custo |
|------------|------------|-------|
| Vector Database | pgvector 0.8.1 (PostgreSQL extension) | R$ 0 |
| Embeddings | OpenAI text-embedding-3-large (1024 dims) | R$ 5-10/mês |
| Chunking | Adaptive sliding window (1000 chars, 200 overlap) | R$ 0 |
| Index | IVFFLAT (lists=100) | R$ 0 |
| AI Auxiliar | Claude 3.5 Haiku (PDF extraction) | Já incluído |

### Arquitetura RAG

```
PDF Upload → Claude AI → fullContent + abstract (Article model)
                ↓
     EmbeddingWorker (background)
                ↓
    ChunkingService (1000 chars chunks)
                ↓
    OpenAI API (text-embedding-3-large)
                ↓
    ArticleEmbedding (pgvector, 1024 dims)
                ↓
    IVFFLAT Index (cosine similarity)
```

### Models Principais

**ArticleEmbedding:**
```go
type ArticleEmbedding struct {
    ID            uuid.UUID        `gorm:"type:uuid;primaryKey"`
    ArticleID     uuid.UUID        `gorm:"type:uuid;not null;index"`
    ChunkIndex    int              `gorm:"type:integer;not null"`
    ChunkText     string           `gorm:"type:text;not null"`
    ChunkMetadata datatypes.JSONMap `gorm:"type:jsonb"`
    Embedding     pgvector.Vector  `gorm:"type:vector(1024)"`
    CreatedAt     time.Time        `gorm:"autoCreateTime"`
}
```

**EmbeddingQueue:**
```go
type EmbeddingQueue struct {
    ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
    EntityType   string    `gorm:"type:varchar(50);not null"` // 'article' | 'score_item'
    EntityID     uuid.UUID `gorm:"type:uuid;not null"`
    Status       string    `gorm:"default:'pending'"` // pending|processing|completed|failed
    RetryCount   int       `gorm:"default:0;check:retry_count <= 5"`
    ErrorMessage *string   `gorm:"type:text"`
    CreatedAt    time.Time `gorm:"autoCreateTime"`
    ProcessedAt  *time.Time
}
```

### Endpoints RAG

```bash
# Busca semântica
GET /api/v1/articles/semantic-search?q=diabetes&limit=10&minSimilarity=0.6

# Recomendações para ScoreItem
GET /api/v1/articles/recommend?scoreItemId=UUID&limit=5

# ScoreItems relacionados a um artigo
GET /api/v1/articles/:id/related-score-items?limit=20

# Estatísticas de embeddings
GET /api/v1/articles/embedding-stats

# Health check RAG
GET /health/rag
```

### Performance

| Métrica | Valor | Target | Status |
|---------|-------|--------|--------|
| Latência SQL (busca vetorial) | 0.29 ms | <50ms | ✅ 99.4% melhor |
| Chunks por artigo | ~25 | ~5-10 | ✅ Artigos completos |
| Taxa de sucesso | 100% | >95% | ✅ |
| Custo mensal | R$ 8/mês | <R$20 | ✅ |

### Chunking Strategy

```
Abstract → Chunk 0 (sempre separado)

Full Content → Sliding Window:
  - Tamanho: 1000 chars (~200-250 tokens)
  - Overlap: 200 chars (20%)
  - Fronteiras: quebra em sentenças (. ! ?)
  - Metadata: {section, word_count}
```

### Sanitização UTF-8

**Problema:** PDFs podem conter bytes inválidos UTF-8
**Solução:** Dupla sanitização (ChunkingService + EmbeddingWorker)

```go
func sanitizeUTF8(s string) string {
    if utf8.ValidString(s) {
        return s
    }
    var cleaned strings.Builder
    for i := 0; i < len(s); {
        r, size := utf8.DecodeRuneInString(s[i:])
        if r == utf8.RuneError && size == 1 {
            cleaned.WriteRune(' ') // Substitui por espaço
            i++
            continue
        }
        cleaned.WriteRune(r)
        i += size
    }
    return cleaned.String()
}
```

### Monitoring

```bash
# Health check RAG
curl http://localhost:3001/health/rag

# Resposta:
{
  "status": "healthy",
  "pgvector": true,
  "embeddings_count": 4656,
  "queue": {
    "pending": 631,
    "processing": 0,
    "completed": 186,
    "failed": 1
  },
  "queue_depth": 631
}
```

### Fallback Strategy

Se busca semântica falhar:
1. Log de erro
2. Retorna HTTP 503 Service Unavailable
3. Mensagem amigável sugerindo busca tradicional
4. Frontend pode redirecionar para busca ILIKE

### API Usage Tracking

Todos os embeddings são rastreados em `api_usage_logs`:

```sql
SELECT
  provider,
  model,
  SUM(total_tokens) as tokens,
  COUNT(*) as requests
FROM api_usage_logs
WHERE provider = 'openai'
GROUP BY provider, model;
```

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
