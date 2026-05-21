# Sistema de Escores - Estrutura de Dados

**Data:** 21 de Janeiro de 2026
**Status:** ✅ Estrutura criada - Pronta para popular com dados

---

## Visão Geral

Sistema completo para armazenar e gerenciar escores clínicos de medicina funcional no Plenya EMR. Permite estratificação de risco em 6 níveis (0-5) para qualquer parâmetro laboratorial ou de imagem.

---

## Estrutura Hierárquica

```
ScoreGroup (Hemograma Completo)
  │
  ├─ ScoreSubgroup (Série Vermelha)
  │    │
  │    ├─ ScoreItem (Hemoglobina - Homens)
  │    │    │
  │    │    ├─ ScoreLevel (Nível 0: <12 g/dL - Anemia severa)
  │    │    ├─ ScoreLevel (Nível 1: 12-13 g/dL - Anemia leve)
  │    │    ├─ ScoreLevel (Nível 2: >17 g/dL - Policitemia)
  │    │    ├─ ScoreLevel (Nível 3: 13-13.9 g/dL - Subótimo)
  │    │    ├─ ScoreLevel (Nível 4: 15.1-17 g/dL - Alto-normal)
  │    │    └─ ScoreLevel (Nível 5: 14.0-15.0 g/dL - Ótimo)
  │    │
  │    └─ ScoreItem (Hematócrito - Homens)
  │         └─ [6 níveis]
  │
  ├─ ScoreSubgroup (Série Branca)
  │    │
  │    ├─ ScoreItem (Leucócitos Totais)
  │    └─ ScoreItem (Neutrófilos)
  │
  └─ ScoreSubgroup (Plaquetas)
       └─ ScoreItem (Plaquetas)

ScoreGroup (Ecocardiograma)
  │
  ├─ ScoreSubgroup (Função Sistólica)
  │    │
  │    ├─ ScoreItem (FEVE - Homens)
  │    └─ ScoreItem (FEVE - Mulheres)
  │
  └─ ScoreSubgroup (Função Diastólica)
       └─ ScoreItem (E/e' Média)
```

---

## Tabelas do Banco de Dados

### 1. `score_groups` - Grupos de Escores

**Propósito:** Categoria principal (ex: "Hemograma Completo", "Ecocardiograma")

| Campo | Tipo | Obrigatório | Descrição |
|-------|------|-------------|-----------|
| `id` | UUID | Sim | Identificador único (UUID v7) |
| `name` | VARCHAR(200) | Sim | Nome do grupo (único) |
| `order` | INTEGER | Sim | Ordem de exibição (default: 0) |
| `created_at` | TIMESTAMP | Sim | Data de criação |
| `updated_at` | TIMESTAMP | Sim | Data de atualização |
| `deleted_at` | TIMESTAMP | Não | Soft delete |

**Relacionamentos:**
- `1:N` com `score_subgroups` (um grupo tem vários subgrupos)

**Exemplo de dados:**
```sql
INSERT INTO score_groups (name, "order") VALUES
  ('Hemograma Completo', 1),
  ('Ecocardiograma', 2),
  ('Colonoscopia', 3);
```

---

### 2. `score_subgroups` - Subgrupos

**Propósito:** Subcategoria dentro de um grupo (ex: "Série Vermelha", "Função Sistólica")

| Campo | Tipo | Obrigatório | Descrição |
|-------|------|-------------|-----------|
| `id` | UUID | Sim | Identificador único (UUID v7) |
| `name` | VARCHAR(200) | Sim | Nome do subgrupo |
| `order` | INTEGER | Sim | Ordem dentro do grupo |
| `group_id` | UUID | Sim | FK para `score_groups` |
| `created_at` | TIMESTAMP | Sim | Data de criação |
| `updated_at` | TIMESTAMP | Sim | Data de atualização |
| `deleted_at` | TIMESTAMP | Não | Soft delete |

**Relacionamentos:**
- `N:1` com `score_groups` (subgrupo pertence a um grupo)
- `1:N` com `score_items` (subgrupo tem vários itens)

**Exemplo de dados:**
```sql
-- Assumindo group_id do Hemograma
INSERT INTO score_subgroups (name, "order", group_id) VALUES
  ('Série Vermelha', 1, '<uuid-do-hemograma>'),
  ('Série Branca', 2, '<uuid-do-hemograma>'),
  ('Plaquetas', 3, '<uuid-do-hemograma>');
```

---

### 3. `score_items` - Itens de Escore

**Propósito:** Parâmetro clínico específico (ex: "Hemoglobina - Homens", "FEVE")

| Campo | Tipo | Obrigatório | Descrição |
|-------|------|-------------|-----------|
| `id` | UUID | Sim | Identificador único (UUID v7) |
| `name` | VARCHAR(300) | Sim | Nome do parâmetro |
| `unit` | VARCHAR(50) | Não | Unidade de medida (g/dL, %, mm) |
| `unit_conversion` | TEXT | Não | Conversão (ex: "1 g/dL = 10 g/L") |
| `points` | DOUBLE | Sim | Pontos máximos (0-100, default: 0) |
| `order` | INTEGER | Sim | Ordem dentro do subgrupo |
| `subgroup_id` | UUID | Sim | FK para `score_subgroups` |
| `parent_item_id` | UUID | Não | FK para `score_items` (hierarquia) |
| `created_at` | TIMESTAMP | Sim | Data de criação |
| `updated_at` | TIMESTAMP | Sim | Data de atualização |
| `deleted_at` | TIMESTAMP | Não | Soft delete |

**Relacionamentos:**
- `N:1` com `score_subgroups` (item pertence a um subgrupo)
- `N:1` com `score_items` (self-reference para hierarquia opcional)
- `1:N` com `score_items` (item pode ter sub-itens)
- `1:N` com `score_levels` (item tem 6 níveis de estratificação)

**Exemplo de dados:**
```sql
-- Assumindo subgroup_id da Série Vermelha
INSERT INTO score_items (name, unit, unit_conversion, points, "order", subgroup_id) VALUES
  ('Hemoglobina - Homens', 'g/dL', '1 g/dL = 10 g/L', 20, 1, '<uuid-serie-vermelha>'),
  ('Hemoglobina - Mulheres', 'g/dL', '1 g/dL = 10 g/L', 20, 2, '<uuid-serie-vermelha>');
```

---

### 4. `score_levels` - Níveis de Estratificação

**Propósito:** Define os 6 níveis de risco para cada item (0=pior, 5=melhor, 6=reservado)

| Campo | Tipo | Obrigatório | Descrição |
|-------|------|-------------|-----------|
| `id` | UUID | Sim | Identificador único (UUID v7) |
| `level` | INTEGER | Sim | Nível (0-6, CHECK constraint) |
| `name` | VARCHAR(500) | Sim | Descrição (ex: "55 a 70 (Ótimo)") - aumentado de 200 para 500 |
| `definition` | TEXT | Não | Definição clínica detalhada |
| `lower_limit` | VARCHAR(50) | Não | Limite inferior (ex: "<40", "55") |
| `upper_limit` | VARCHAR(50) | Não | Limite superior (ex: "70", ">75") |
| `item_id` | UUID | Sim | FK para `score_items` |
| `created_at` | TIMESTAMP | Sim | Data de criação |
| `updated_at` | TIMESTAMP | Sim | Data de atualização |
| `deleted_at` | TIMESTAMP | Não | Soft delete |

**Relacionamentos:**
- `N:1` com `score_items` (nível pertence a um item)

**Exemplo de dados (FEVE Homens):**
```sql
-- Assumindo item_id da FEVE Homens
INSERT INTO score_levels (level, name, definition, lower_limit, upper_limit, item_id) VALUES
  (0, '<40 (IC)', 'Insuficiência cardíaca sistólica', '<40', NULL, '<uuid-feve>'),
  (1, '40 a 49 (Leve↓)', 'Disfunção sistólica leve', '40', '49', '<uuid-feve>'),
  (2, '>75 (Hiperdinâmico)', 'Estado hipercinético patológico', '>75', NULL, '<uuid-feve>'),
  (3, '50 a 54 (Limítrofe)', 'Borderline - monitorar', '50', '54', '<uuid-feve>'),
  (4, '71 a 75 (Alto-normal)', 'Alto-normal', '71', '75', '<uuid-feve>'),
  (5, '55 a 70 (Ótimo)', 'Faixa ótima - menor mortalidade CV', '55', '70', '<uuid-feve>');
```

---

## Convenções de Níveis

### Sistema de 6 Níveis (0-5)

**Nível 0 (Crítico):** Valor mais perigoso - extremo baixo ou alto patológico
- Exemplos: FEVE <40% (IC), Hemoglobina <12 g/dL (anemia severa)

**Nível 1 (Muito Baixo/Alto):** Valor anormal requerendo intervenção
- Exemplos: FEVE 40-49% (disfunção leve), Hemoglobina 12-13 g/dL

**Nível 2 (Subótimo ou Extremo Oposto):** Valor abaixo do ideal OU extremo alto patológico
- Exemplos: FEVE >75% (hiperdinâmico), Hemoglobina >17 g/dL (policitemia)
- **IMPORTANTE:** Curva em U - extremos altos também vão aqui!

**Nível 3 (Limítrofe):** Borderline - monitorar
- Exemplos: FEVE 50-54%, Hemoglobina 13-13.9 g/dL

**Nível 4 (Bom):** Dentro do normal convencional
- Exemplos: FEVE 71-75%, Hemoglobina 15.1-17 g/dL

**Nível 5 (Ótimo):** Faixa ideal de medicina funcional
- Exemplos: FEVE 55-70%, Hemoglobina 14.0-15.0 g/dL
- **Gold standard:** Menor mortalidade, melhor performance

**Nível 6:** Reservado para expansão futura (vazio por enquanto)

---

## Casos Especiais

### 1. Curva em U (U-Shaped Curves)

Alguns parâmetros têm valores ótimos no meio da faixa, com extremos baixos E altos sendo patológicos:

**FEVE (Fração de Ejeção):**
- Nível 0: <40% (insuficiência cardíaca)
- Nível 2: >75% (hiperdinâmico - sepsis, tireotoxicose)
- Nível 5: 55-70% (ótimo)

**TAPSE:**
- Nível 0: <16mm (disfunção VD)
- Nível 2: >30mm (hipercinesia anormal)
- Nível 5: 20-26mm (ótimo)

**LAVI (Volume Atrial Esquerdo):**
- Nível 0: >48 mL/m² (dilatação severa)
- Nível 2: <22 mL/m² (muito pequeno - também anormal)
- Nível 5: 22-27.9 mL/m² (ótimo)

**Regra:** Extremos SEMPRE vão em Nível 0 ou 2, NUNCA em Nível 5!

### 2. Hierarquia com `parent_item_id`

Para itens com sub-itens hierárquicos (raro, mas suportado):

```sql
-- Item pai
INSERT INTO score_items (name, subgroup_id) VALUES
  ('Lipidograma', '<subgroup-id>') RETURNING id;

-- Sub-itens
INSERT INTO score_items (name, unit, subgroup_id, parent_item_id) VALUES
  ('Colesterol Total', 'mg/dL', '<subgroup-id>', '<lipidograma-id>'),
  ('LDL', 'mg/dL', '<subgroup-id>', '<lipidograma-id>'),
  ('HDL', 'mg/dL', '<subgroup-id>', '<lipidograma-id>');
```

---

## Constraints e Validações

### Constraints SQL

1. **Unique Index:** `score_groups.name` deve ser único (ignorando soft deletes)
2. **Foreign Keys:** Todas com `ON DELETE CASCADE` ou `SET NULL`
3. **Check Constraints:**
   - `score_items.points` entre 0 e 100
   - `score_levels.level` entre 0 e 6
4. **Soft Delete:** Todas as tabelas têm `deleted_at` para soft delete

### Validações Go (Go Playground Validator)

```go
// score_item.go
validate:"required,min=2,max=300"  // name
validate:"gte=0,lte=100"           // points
validate:"omitempty,max=50"        // unit

// score_level.go
validate:"required,gte=0,lte=6"    // level
validate:"required,min=1,max=200"  // name
```

---

## Índices para Performance

**Índices criados:**
- `idx_score_group_name` (UNIQUE, WHERE deleted_at IS NULL)
- `idx_score_group_order`
- `idx_score_subgroup_group` (FK)
- `idx_score_subgroup_order`
- `idx_score_item_subgroup` (FK)
- `idx_score_item_parent` (FK)
- `idx_score_item_order`
- `idx_score_level_item` (FK)
- `idx_score_level_level`
- Todos os `deleted_at` têm índice para soft delete

**Benefícios:**
- Queries rápidas por grupo/subgrupo
- Ordenação eficiente por `order`
- Joins otimizados nas FKs

---

## Próximos Passos

### 1. Popular com Dados do CSV

Você mencionou ter uma tabela de escores. Formato esperado:

```csv
Nome;Unidade | Conversão;Pontos;Nível 0;Nível 1;Nível 2;Nível 3;Nível 4;Nível 5
Hemoglobina - Homens;g/dL | 1 g/dL = 10 g/L;20;<12 (Anemia severa);12 a 13 (Anemia leve);>17 (Policitemia);13 a 13.9 (Subótimo);15.1 a 17 (Alto-normal);14.0 a 15.0 (Ótimo)
```

**Script de importação necessário para:**
1. Criar `ScoreGroup` para cada categoria (Hemograma, Ecodopplercardiograma, etc.)
2. Criar `ScoreSubgroup` para cada subcategoria
3. Criar `ScoreItem` de cada linha do CSV
4. Criar 6 `ScoreLevel` por item (parseando os níveis 0-5)

### 2. Criar Handlers e Endpoints

**Endpoints sugeridos:**
```
GET    /api/v1/score-groups              # Listar grupos
GET    /api/v1/score-groups/:id          # Detalhes do grupo
GET    /api/v1/score-groups/:id/tree     # Grupo com subgrupos, itens e níveis (nested)
GET    /api/v1/score-items/:id/levels    # Níveis de um item
POST   /api/v1/score-items/:id/evaluate  # Avaliar valor do paciente
```

### 3. Implementar Lógica de Avaliação

```go
// EvaluateScore retorna o nível (0-5) para um valor medido
func EvaluateScore(itemID uuid.UUID, value float64) (*ScoreLevel, error) {
    // 1. Buscar item e seus níveis
    // 2. Comparar valor com lower_limit e upper_limit de cada nível
    // 3. Retornar o nível correspondente
}
```

### 4. Integrar com Lab Results

```go
// LabResult modelo existente precisa referenciar ScoreItem
type LabResult struct {
    // ... campos existentes
    ScoreItemID *uuid.UUID  // FK para score_items
    ScoreLevel  *int        // Nível calculado (0-5)
}
```

---

## Arquivos Criados

**Go Models (Fonte Única de Verdade):**
- `apps/api/internal/models/score_group.go` ✅
- `apps/api/internal/models/score_subgroup.go` ✅
- `apps/api/internal/models/score_item.go` ✅
- `apps/api/internal/models/score_level.go` ✅

**Migration SQL:**
- `apps/api/database/migrations/20260121_create_score_tables.sql` ✅

**Documentação:**
- `SCORE-SYSTEM-STRUCTURE.md` (este arquivo) ✅

---

## Exemplo Completo: Hemoglobina - Homens

```sql
-- 1. Grupo
INSERT INTO score_groups (id, name, "order")
VALUES ('01943c8a-0000-7000-8000-000000000001', 'Hemograma Completo', 1);

-- 2. Subgrupo
INSERT INTO score_subgroups (id, name, "order", group_id)
VALUES ('01943c8a-0000-7000-8000-000000000002', 'Série Vermelha', 1, '01943c8a-0000-7000-8000-000000000001');

-- 3. Item
INSERT INTO score_items (id, name, unit, unit_conversion, points, "order", subgroup_id)
VALUES ('01943c8a-0000-7000-8000-000000000003', 'Hemoglobina - Homens', 'g/dL', '1 g/dL = 10 g/L', 20, 1, '01943c8a-0000-7000-8000-000000000002');

-- 4. Níveis
INSERT INTO score_levels (level, name, lower_limit, upper_limit, item_id) VALUES
  (0, '<12 (Anemia severa)', '<12', NULL, '01943c8a-0000-7000-8000-000000000003'),
  (1, '12 a 13 (Anemia leve)', '12', '13', '01943c8a-0000-7000-8000-000000000003'),
  (2, '>17 (Policitemia)', '>17', NULL, '01943c8a-0000-7000-8000-000000000003'),
  (3, '13 a 13.9 (Subótimo)', '13', '13.9', '01943c8a-0000-7000-8000-000000000003'),
  (4, '15.1 a 17 (Alto-normal)', '15.1', '17', '01943c8a-0000-7000-8000-000000000003'),
  (5, '14.0 a 15.0 (Ótimo)', '14.0', '15.0', '01943c8a-0000-7000-8000-000000000003');
```

**Query de teste:**
```sql
SELECT
    g.name as grupo,
    s.name as subgrupo,
    i.name as item,
    i.unit,
    l.level,
    l.name as nivel_descricao
FROM score_groups g
JOIN score_subgroups s ON s.group_id = g.id
JOIN score_items i ON i.subgroup_id = s.id
JOIN score_levels l ON l.item_id = i.id
WHERE i.name = 'Hemoglobina - Homens'
ORDER BY l.level;
```

---

**Status:** ✅ Estrutura completa e pronta para receber dados
**Próximo passo:** Envie a tabela CSV de escores para popular o banco
**Última atualização:** 21/01/2026
