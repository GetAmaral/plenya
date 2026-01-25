# Implementação do Sistema de Artigos Científicos

## Resumo

Sistema completo para gerenciar artigos científicos no Plenya EMR, com upload de PDF e extração automática de metadados.

---

## Arquitetura

```
Article Model (Go)
    ↓
Migration (PostgreSQL)
    ↓
Repository Layer
    ↓
Service Layer
    ↓
Handler Layer (REST API)
```

---

## Modelo de Dados

### Tabela: `articles`

**Metadados Bibliográficos:**
- `title` (varchar 1000, indexed) - Título do artigo
- `authors` (varchar 2000) - Autores (texto livre)
- `journal` (varchar 500, indexed) - Nome da revista
- `publish_date` (date, indexed) - Data de publicação
- `language` (varchar 10, default: 'en') - Idioma

**Identificadores Únicos:**
- `doi` (varchar 255, unique index) - Digital Object Identifier
- `pmid` (varchar 50, indexed) - PubMed ID
- `issn` (varchar 20) - ISSN da revista

**Conteúdo:**
- `abstract` (text, nullable) - Resumo do artigo
- `full_content` (text, nullable) - Texto completo extraído do PDF
- `notes` (text, nullable) - Notas pessoais do médico

**Links:**
- `original_link` (varchar 2048, nullable) - URL da publicação original
- `internal_link` (varchar 2048, nullable) - Path do PDF armazenado localmente

**Classificação:**
- `article_type` (enum, default: 'research_article'):
  - `research_article` - Artigo de pesquisa original
  - `review` - Revisão sistemática/narrativa
  - `meta_analysis` - Meta-análise
  - `case_study` - Estudo de caso
  - `clinical_trial` - Ensaio clínico
  - `editorial` - Editorial
  - `letter` - Carta ao editor
  - `protocol` - Protocolo de estudo
  - `lecture` - Palestra/aula ✨ (novo)

- `keywords` (jsonb) - Palavras-chave
- `mesh_terms` (jsonb) - Medical Subject Headings
- `specialty` (varchar 200, indexed) - Especialidade médica

**Gestão Pessoal:**
- `favorite` (boolean, indexed, default: false) - Marcado como favorito
- `rating` (smallint 0-5, nullable) - Avaliação pessoal (estrelas)

**Arquivo:**
- `file_hash` (varchar 64, indexed) - SHA-256 do PDF (anti-duplicação)
- `file_size` (bigint, nullable) - Tamanho do arquivo em bytes

**Metadados do Sistema:**
- `indexed_at` (timestamp) - Data de indexação no sistema
- `last_accessed_at` (timestamp, nullable) - Último acesso

**Audit:**
- `created_by` (uuid, nullable) - User ID que criou
- `updated_by` (uuid, nullable) - User ID que atualizou

**Timestamps:**
- `created_at` (timestamp)
- `updated_at` (timestamp)
- `deleted_at` (timestamp, nullable) - Soft delete

---

## Funcionalidades Implementadas

### 1. Upload de PDF com Extração Automática ✨

**Endpoint:** `POST /api/v1/articles/upload`

**Features:**
- Upload de arquivo PDF (máximo 50MB)
- Validação de tipo de arquivo (application/pdf)
- Cálculo de hash SHA-256 (anti-duplicação)
- **Extração automática de metadados:**
  - DOI (via regex: `doi: 10.xxxx/...`)
  - PMID (via regex: `pmid: xxxxxxx`)
  - Título (primeira linha significativa)
  - Abstract (texto após palavra "abstract")
  - Texto completo do PDF (todas as páginas)

**Integração com APIs Externas:**
- Se DOI for encontrado → Busca metadados via **CrossRef API** (gratuita):
  - Título
  - Autores
  - Abstract
  - Keywords/Subjects

**Armazenamento:**
- Path: `./uploads/articles/{uuid}.pdf`
- UUID v4 para cada arquivo
- Soft delete remove arquivo físico

**Exemplo de uso:**
```bash
curl -X POST http://localhost:3001/api/v1/articles/upload \
  -H "Authorization: Bearer {token}" \
  -F "file=@article.pdf"
```

### 2. CRUD Completo

**Criar artigo manualmente:**
```
POST /api/v1/articles
```

**Listar artigos (com filtros e paginação):**
```
GET /api/v1/articles?page=1&pageSize=20&journal=Nature&favorite=true&rating=5
```

**Filtros disponíveis:**
- `journal` - Filtrar por revista
- `specialty` - Filtrar por especialidade
- `articleType` - Filtrar por tipo
- `favorite` - Apenas favoritos
- `rating` - Filtrar por rating (0-5)
- `publishedAfter` - Data de publicação (formato: YYYY-MM-DD)
- `publishedBefore` - Data de publicação

**Buscar por ID:**
```
GET /api/v1/articles/{id}
```
- Atualiza automaticamente `last_accessed_at` (em goroutine)

**Atualizar artigo:**
```
PUT /api/v1/articles/{id}
```

**Deletar artigo (soft delete):**
```
DELETE /api/v1/articles/{id}
```
- Remove arquivo PDF do filesystem

### 3. Busca Full-Text

**Endpoint:** `GET /api/v1/articles/search?q={query}&page=1&pageSize=20`

Busca por texto em:
- Título
- Autores
- Abstract
- Journal

**Exemplo:**
```bash
curl "http://localhost:3001/api/v1/articles/search?q=cardiovascular&page=1" \
  -H "Authorization: Bearer {token}"
```

### 4. Favoritos

**Listar favoritos:**
```
GET /api/v1/articles/favorites?page=1&pageSize=20
```

**Toggle favorito:**
```
PATCH /api/v1/articles/{id}/favorite
```
- Alterna entre favorito/não favorito

### 5. Rating

**Definir rating (0-5 estrelas):**
```
PATCH /api/v1/articles/{id}/rating
Body: { "rating": 4 }
```

### 6. Download de PDF

**Endpoint:** `GET /api/v1/articles/{id}/download`

Retorna o arquivo PDF original.

### 7. Estatísticas

**Endpoint:** `GET /api/v1/articles/stats`

(Placeholder - a ser implementado)

---

## Segurança e Permissões

**Permissões compartilhadas:**
- ✅ Todos os médicos veem todos os artigos (biblioteca compartilhada)
- ✅ Apenas `doctors` e `nurses` podem criar/editar/deletar artigos
- ✅ Todos os usuários autenticados podem:
  - Marcar favoritos
  - Dar rating
  - Fazer buscas
  - Baixar PDFs

**Audit log:**
- Todas as operações são registradas automaticamente via middleware

**Anti-duplicação:**
- Hash SHA-256 do arquivo impede upload de PDFs duplicados
- DOI único impede duplicação por identificador

---

## Extração de Metadados

### Biblioteca Usada

**Go PDF Parser:** `github.com/ledongthuc/pdf v0.0.0-20250511090121-5959a4027728`

### Regex para Extração

**DOI:**
```regex
(?i)doi[:\s]*([10]\.\d{4,}/[^\s]+)
```

**PMID:**
```regex
(?i)pmid[:\s]*(\d{7,})
```

**Abstract:**
```regex
(?is)abstract[:\s]+(.*?)(?:introduction|keywords|background|\n\n)
```

### CrossRef API

**URL:** `https://api.crossref.org/works/{doi}`

**Sem necessidade de API key** (gratuita)

**Dados retornados:**
- `title[]` - Título do artigo
- `abstract` - Resumo
- `author[]` - Lista de autores (given, family)
- `subject[]` - Keywords/áreas

---

## Arquivos Criados

### Backend Go

1. **`apps/api/internal/models/article.go`** (206 linhas)
   - Modelo Article com annotations Swagger
   - Validações GORM
   - Hooks BeforeCreate

2. **`apps/api/internal/services/article_service.go`** (500+ linhas)
   - DTOs (CreateArticleDTO, UpdateArticleDTO)
   - CRUD completo
   - Upload de PDF
   - Extração de metadados do PDF
   - Integração com CrossRef API
   - Busca full-text
   - Anti-duplicação por hash

3. **`apps/api/internal/repository/article_repository.go`** (280+ linhas)
   - Operações de banco de dados
   - Queries com filtros
   - Paginação
   - Full-text search (ILIKE)
   - Métodos auxiliares (GetByDOI, GetByPMID, GetFavorites, GetStatistics)

4. **`apps/api/internal/handlers/article_handler.go`** (580+ linhas)
   - 12 endpoints REST
   - Validações de input
   - Upload multipart/form-data
   - Annotations Swagger completas

### Configuração

5. **`apps/api/cmd/server/main.go`** (modificado)
   - Registro do ArticleService
   - Registro do ArticleHandler
   - Rotas `/api/v1/articles/*`
   - Servir arquivos estáticos `/uploads`

6. **`apps/api/internal/database/database.go`** (modificado)
   - Adicionado `&models.Article{}` no AutoMigrate

7. **`apps/api/go.mod`** (modificado)
   - Adicionada dependência `github.com/ledongthuc/pdf`

### Infraestrutura

8. **`/home/user/plenya/uploads/articles/`** (diretório)
   - Armazenamento local de PDFs

---

## Próximos Passos (Opcional)

### Melhorias Futuras

1. **Full-Text Search PostgreSQL:**
   - Implementar `ts_vector` e `ts_query` para busca mais sofisticada
   - Suporte a stemming em português/inglês
   - Ranking de relevância

2. **Integração com mais APIs:**
   - **PubMed API** (artigos médicos)
   - **Europe PMC API**
   - **Unpaywall API** (acesso aberto)

3. **OCR para PDFs escaneados:**
   - Usar Tesseract via Docker para extrair texto de imagens

4. **Anotações em PDF:**
   - Highlights
   - Comentários
   - Marcações

5. **Export bibliográfico:**
   - BibTeX (para LaTeX)
   - RIS (para EndNote)
   - EndNote XML

6. **Notificações:**
   - Alertas de novos artigos em tópicos de interesse
   - Recomendações baseadas em especialidade

7. **Compartilhamento:**
   - Compartilhar artigos específicos com outros médicos
   - Comentários e discussões

8. **Estatísticas:**
   - Dashboard com métricas:
     - Total de artigos
     - Artigos por tipo
     - Journals mais comuns
     - Especialidades mais representadas

---

## Endpoints da API

### Autenticação

Todos os endpoints requerem `Authorization: Bearer {token}` no header.

### Lista Completa

| Método | Endpoint | Descrição | Permissão |
|--------|----------|-----------|-----------|
| POST | `/api/v1/articles` | Criar artigo manualmente | Medical Staff |
| POST | `/api/v1/articles/upload` | Upload de PDF | Medical Staff |
| GET | `/api/v1/articles` | Listar artigos | Autenticado |
| GET | `/api/v1/articles/search` | Buscar artigos | Autenticado |
| GET | `/api/v1/articles/favorites` | Listar favoritos | Autenticado |
| GET | `/api/v1/articles/stats` | Estatísticas | Autenticado |
| GET | `/api/v1/articles/:id` | Buscar por ID | Autenticado |
| PUT | `/api/v1/articles/:id` | Atualizar artigo | Medical Staff |
| DELETE | `/api/v1/articles/:id` | Deletar artigo | Medical Staff |
| PATCH | `/api/v1/articles/:id/favorite` | Toggle favorito | Autenticado |
| PATCH | `/api/v1/articles/:id/rating` | Definir rating | Autenticado |
| GET | `/api/v1/articles/:id/download` | Download PDF | Autenticado |

---

## Estrutura de Resposta

### Article Object

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "Cardiovascular Risk Factors in Hypertensive Patients",
  "authors": "Silva JA, Santos MF, Oliveira PR",
  "journal": "Nature Medicine",
  "publishDate": "2024-03-15T00:00:00Z",
  "language": "en",
  "doi": "10.1038/s41591-024-12345-6",
  "pmid": "38355808",
  "issn": "1546-170X",
  "abstract": "Background: Cardiovascular disease remains...",
  "fullContent": "Full text extracted from PDF...",
  "notes": "Important for clinical practice",
  "originalLink": "https://www.nature.com/articles/s41591-024-12345-6",
  "internalLink": "/uploads/articles/550e8400-e29b-41d4-a716-446655440000.pdf",
  "articleType": "research_article",
  "keywords": ["cardiovascular", "hypertension", "risk factors"],
  "meshTerms": ["Cardiovascular Diseases", "Hypertension"],
  "specialty": "Cardiology",
  "favorite": true,
  "rating": 5,
  "fileHash": "a1b2c3d4e5f6...",
  "fileSize": 2457600,
  "indexedAt": "2026-01-25T05:42:57Z",
  "lastAccessedAt": "2026-01-25T06:15:30Z",
  "createdBy": "660e8400-e29b-41d4-a716-446655440000",
  "updatedBy": "660e8400-e29b-41d4-a716-446655440000",
  "createdAt": "2026-01-25T05:42:57Z",
  "updatedAt": "2026-01-25T06:15:30Z"
}
```

### Paginated Response

```json
{
  "articles": [/* array of Article objects */],
  "pagination": {
    "page": 1,
    "pageSize": 20,
    "total": 142,
    "totalPages": 8
  }
}
```

---

## Testes

### Testar Upload de PDF

```bash
# 1. Fazer login
TOKEN=$(curl -X POST http://localhost:3001/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"doctor@example.com","password":"password123"}' \
  | jq -r '.token')

# 2. Upload de PDF
curl -X POST http://localhost:3001/api/v1/articles/upload \
  -H "Authorization: Bearer $TOKEN" \
  -F "file=@path/to/article.pdf"

# 3. Listar artigos
curl http://localhost:3001/api/v1/articles \
  -H "Authorization: Bearer $TOKEN"

# 4. Buscar por texto
curl "http://localhost:3001/api/v1/articles/search?q=cardiovascular" \
  -H "Authorization: Bearer $TOKEN"

# 5. Marcar como favorito
ARTICLE_ID="550e8400-e29b-41d4-a716-446655440000"
curl -X PATCH "http://localhost:3001/api/v1/articles/$ARTICLE_ID/favorite" \
  -H "Authorization: Bearer $TOKEN"

# 6. Dar rating
curl -X PATCH "http://localhost:3001/api/v1/articles/$ARTICLE_ID/rating" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"rating": 5}'

# 7. Download PDF
curl "http://localhost:3001/api/v1/articles/$ARTICLE_ID/download" \
  -H "Authorization: Bearer $TOKEN" \
  -o article_downloaded.pdf
```

---

## Status

✅ **Implementação Backend Completa**
- Modelo de dados
- Repository layer
- Service layer com extração de PDF
- Handlers REST API
- 12 endpoints funcionais
- Integração com CrossRef API

🔲 **Próximos Passos (Frontend):**
- Interface web para listar artigos
- Upload de PDF via interface
- Busca e filtros
- Favoritos e ratings
- Visualizador de PDF integrado

---

**Última atualização:** 25 de Janeiro de 2026
