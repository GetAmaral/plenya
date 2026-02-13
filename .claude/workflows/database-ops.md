# Database Operations - Operações Diretas no Banco

## 🚨 REGRA DE OURO

**Durante DESENVOLVIMENTO (quando você, Claude, está manipulando dados manualmente):**
- ✅ Usar Docker exec → psql
- ✅ Usar Go scripts temporários com GORM
- ❌ **NUNCA** usar API HTTP (POST/PUT/DELETE via curl/fetch)

**Em PRODUÇÃO (apps web/mobile):**
- ✅ Usar API HTTP via TanStack Query

## Por Quê?

A API exige autenticação, middleware, e é destinada a aplicações finais. Para desenvolvimento/testes, acesso direto ao banco é mais eficiente e correto.

## Método 1: PostgreSQL via Docker (SQL Direto)

### Acessar banco interativo

```bash
docker compose exec db psql -U plenya_user -d plenya_db
```

### Executar SQL direto

```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
INSERT INTO score_items (
  id, name, unit, gender, subgroup_id, \"order\", created_at, updated_at
) VALUES (
  gen_random_uuid(),
  'Hemoglobina - Homens',
  'g/dL',
  'male',
  '019c1a1f-5678-7f3b-a1bd-4767008e844c',
  1,
  NOW(),
  NOW()
);
"
```

### Queries úteis

```bash
# Listar score items
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT id, name, unit, gender, subgroup_id
FROM score_items
ORDER BY \"order\" ASC;
"

# Listar levels de um item
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT level, name, operator, lower_limit, upper_limit
FROM score_levels
WHERE item_id = '019c1a20-1234-7f3b-a1bd-4767008e844c'
ORDER BY level ASC;
"

# Buscar por nome
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT id, name FROM score_items WHERE name ILIKE '%hemoglobina%';
"

# Deletar (soft delete - seta deleted_at)
docker compose exec db psql -U plenya_user -d plenya_db -c "
UPDATE score_items SET deleted_at = NOW() WHERE id = 'uuid-aqui';
"

# Deletar permanente (cuidado!)
docker compose exec db psql -U plenya_user -d plenya_db -c "
DELETE FROM score_items WHERE id = 'uuid-aqui';
"
```

## Método 2: Go Script Temporário (Recomendado para Operações Complexas)

### Quando usar?

- Operações que envolvem múltiplas tabelas
- Necessita validação de regras de negócio
- Quer usar métodos do model (AppliesToPatient, EvaluatesTrue)
- Quer usar Preload/relações GORM

### Template de script

Criar `apps/api/scripts/add_score_item.go`:

```go
package main

import (
    "fmt"
    "log"
    "github.com/google/uuid"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/plenya/api/internal/models"
)

func main() {
    // Conectar ao banco
    dsn := "host=localhost user=plenya_user password=senha_segura dbname=plenya_db port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect:", err)
    }

    // Criar ScoreItem
    subgroupID := uuid.MustParse("019c1a1f-5678-7f3b-a1bd-4767008e844c")

    item := &models.ScoreItem{
        Name:        "Hemoglobina - Homens",
        Unit:        stringPtr("g/dL"),
        Gender:      stringPtr("male"),
        AgeRangeMin: intPtr(18),
        AgeRangeMax: intPtr(65),
        SubgroupID:  subgroupID,
        Order:       1,
    }

    // BeforeCreate hook gera UUID v7 automaticamente
    if err := db.Create(item).Error; err != nil {
        log.Fatal("Failed to create item:", err)
    }

    fmt.Printf("✅ Item criado: %s (ID: %s)\n", item.Name, item.ID)

    // Criar níveis
    levels := []models.ScoreLevel{
        {
            Level:      6,
            Name:       ">= 14 g/dL (Ótimo)",
            LowerLimit: stringPtr("14"),
            Operator:   ">=",
            ItemID:     item.ID,
        },
        {
            Level:      5,
            Name:       "13-13.9 g/dL (Normal)",
            LowerLimit: stringPtr("13"),
            UpperLimit: stringPtr("13.9"),
            Operator:   "between",
            ItemID:     item.ID,
        },
        {
            Level:      3,
            Name:       "11-12.9 g/dL (Anemia leve)",
            LowerLimit: stringPtr("11"),
            UpperLimit: stringPtr("12.9"),
            Operator:   "between",
            ItemID:     item.ID,
        },
        {
            Level:      1,
            Name:       "< 11 g/dL (Anemia moderada/grave)",
            LowerLimit: stringPtr("11"),
            Operator:   "<",
            ItemID:     item.ID,
        },
    }

    for _, level := range levels {
        if err := db.Create(&level).Error; err != nil {
            log.Fatal("Failed to create level:", err)
        }
        fmt.Printf("✅ Level criado: %s\n", level.Name)
    }

    // Recarregar com Preload para verificar
    var reloaded models.ScoreItem
    db.Preload("Levels", func(db *gorm.DB) *gorm.DB {
        return db.Order("level ASC")
    }).First(&reloaded, "id = ?", item.ID)

    fmt.Printf("\n📊 Item final: %s com %d níveis\n", reloaded.Name, len(reloaded.Levels))
}

func stringPtr(s string) *string { return &s }
func intPtr(i int) *int          { return &i }
```

### Executar script

```bash
# Dentro do container API
docker compose exec api go run scripts/add_score_item.go

# OU fora do container (se tiver Go instalado localmente)
cd apps/api
go run scripts/add_score_item.go
```

## Método 3: Repository Layer via Go Test

### Criar arquivo de teste temporário

`apps/api/internal/repository/score_repository_manual_test.go`:

```go
package repository

import (
    "testing"
    "github.com/google/uuid"
    "github.com/plenya/api/internal/models"
)

// Execute com: go test -run TestManualCreateScoreItem
func TestManualCreateScoreItem(t *testing.T) {
    // Setup database connection
    db := setupTestDB(t) // Você precisa ter um helper
    repo := NewScoreRepository(db)

    subgroupID := uuid.MustParse("019c1a1f-5678-7f3b-a1bd-4767008e844c")

    item := &models.ScoreItem{
        Name:       "Hemoglobina - Homens",
        Unit:       stringPtr("g/dL"),
        Gender:     stringPtr("male"),
        SubgroupID: subgroupID,
        Order:      1,
    }

    err := repo.CreateScoreItem(item)
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("✅ Item criado: %s", item.ID)
}
```

```bash
cd apps/api
go test -run TestManualCreateScoreItem ./internal/repository
```

## Comparação de Métodos

| Método | Quando Usar | Prós | Contras |
|--------|-------------|------|---------|
| **psql direto** | Queries simples, SELECT, UPDATE rápido | Rápido, direto | Não executa hooks Go |
| **Go script** | CRUD com validação, múltiplas tabelas | Hooks executam, type-safe | Precisa compilar/executar |
| **Go test** | Operações repetíveis, CI/CD | Testável, versionado | Overhead de teste |
| **API HTTP** | ❌ NÃO USAR PARA DEV | - | Precisa auth, middleware |

## Exemplos Práticos

### Adicionar ScoreItem completo com níveis

```bash
docker compose exec db psql -U plenya_user -d plenya_db << 'EOF'
BEGIN;

-- Inserir item
INSERT INTO score_items (id, name, unit, gender, age_range_min, age_range_max, subgroup_id, "order", created_at, updated_at)
VALUES (
  '019c1a20-1234-7f3b-a1bd-4767008e844c',
  'Hemoglobina - Homens',
  'g/dL',
  'male',
  18,
  65,
  '019c1a1f-5678-7f3b-a1bd-4767008e844c',
  1,
  NOW(),
  NOW()
);

-- Inserir níveis
INSERT INTO score_levels (id, level, name, lower_limit, upper_limit, operator, item_id, created_at, updated_at) VALUES
  (gen_random_uuid(), 6, '>= 14 g/dL (Ótimo)', '14', NULL, '>=', '019c1a20-1234-7f3b-a1bd-4767008e844c', NOW(), NOW()),
  (gen_random_uuid(), 5, '13-13.9 g/dL (Normal)', '13', '13.9', 'between', '019c1a20-1234-7f3b-a1bd-4767008e844c', NOW(), NOW()),
  (gen_random_uuid(), 3, '11-12.9 g/dL (Anemia leve)', '11', '12.9', 'between', '019c1a20-1234-7f3b-a1bd-4767008e844c', NOW(), NOW()),
  (gen_random_uuid(), 1, '< 11 g/dL (Anemia moderada/grave)', NULL, '11', '<', '019c1a20-1234-7f3b-a1bd-4767008e844c', NOW(), NOW());

COMMIT;
EOF
```

### Atualizar campos de um item existente

```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
UPDATE score_items
SET clinical_relevance = 'Hemoglobina baixa indica anemia',
    patient_explanation = 'Transporta oxigênio no sangue',
    conduct = 'Investigar causa, suplementar se necessário',
    last_review = NOW()
WHERE id = '019c1a20-1234-7f3b-a1bd-4767008e844c';
"
```

### Duplicar item com novos filtros

```bash
docker compose exec db psql -U plenya_user -d plenya_db << 'EOF'
-- Duplicar item mudando gender e post_menopause
INSERT INTO score_items (id, name, unit, gender, age_range_min, age_range_max, post_menopause, subgroup_id, "order", created_at, updated_at)
SELECT
  gen_random_uuid(),
  'Hemoglobina - Mulheres pós-menopausa',
  unit,
  'female',
  age_range_min,
  age_range_max,
  true,
  subgroup_id,
  "order" + 1,
  NOW(),
  NOW()
FROM score_items
WHERE id = '019c1a20-original-uuid';

-- Copiar níveis para o novo item
INSERT INTO score_levels (id, level, name, lower_limit, upper_limit, operator, item_id, created_at, updated_at)
SELECT
  gen_random_uuid(),
  level,
  name,
  lower_limit,
  upper_limit,
  operator,
  'novo-item-uuid-aqui',
  NOW(),
  NOW()
FROM score_levels
WHERE item_id = '019c1a20-original-uuid';
EOF
```

## Debugging

### Ver estrutura de tabela

```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "\d score_items"
```

### Ver constraints

```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT conname, contype, pg_get_constraintdef(oid)
FROM pg_constraint
WHERE conrelid = 'score_items'::regclass;
"
```

### Ver foreign keys

```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT
  tc.table_name,
  kcu.column_name,
  ccu.table_name AS foreign_table_name,
  ccu.column_name AS foreign_column_name
FROM information_schema.table_constraints AS tc
JOIN information_schema.key_column_usage AS kcu
  ON tc.constraint_name = kcu.constraint_name
JOIN information_schema.constraint_column_usage AS ccu
  ON ccu.constraint_name = tc.constraint_name
WHERE tc.constraint_type = 'FOREIGN KEY' AND tc.table_name='score_items';
"
```

## ⚠️ Avisos

1. **UUIDs:** Use `gen_random_uuid()` no SQL ou deixe o BeforeCreate hook gerar (Go)
2. **Soft Delete:** `deleted_at` marca como deletado (GORM ignora automaticamente)
3. **Order:** Campo reservado, sempre escapar com `"order"`
4. **Timestamps:** `created_at` e `updated_at` devem sempre ser preenchidos
5. **Relações:** Use Preload no Go, JOINs no SQL
