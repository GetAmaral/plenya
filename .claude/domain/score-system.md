# Sistema de Escores - Hierarquia e Manipulação

## Visão Geral

Sistema de estratificação de risco clínico baseado em exames laboratoriais e parâmetros demográficos.

## Hierarquia (4 níveis)

```
ScoreGroup (ex: "Cardiovascular", "Metabólico")
  │
  └─ ScoreSubgroup (ex: "Função Cardíaca", "Lipídios")
      │
      └─ ScoreItem (ex: "FEVE", "Hemoglobina - Homens")
          │
          ├─ ScoreLevel (ex: "Nível 5: 55-70% (Normal)")
          │
          └─ Articles (M:N) - Artigos científicos de referência
```

## Models

### ScoreGroup

```sql
CREATE TABLE score_groups (
  id UUID PRIMARY KEY,
  name VARCHAR(200) NOT NULL,
  "order" INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
```

**Campos:**
- `name` - Nome do grupo (ex: "Cardiovascular")
- `order` - Ordem de exibição

### ScoreSubgroup

```sql
CREATE TABLE score_subgroups (
  id UUID PRIMARY KEY,
  name VARCHAR(200) NOT NULL,
  "order" INTEGER NOT NULL DEFAULT 0,
  max_select INTEGER NOT NULL DEFAULT 0,
  group_id UUID NOT NULL REFERENCES score_groups(id),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
```

**Campos:**
- `name` - Nome do subgrupo (ex: "Hemograma")
- `order` - Ordem dentro do grupo
- `max_select` - Quantos items podem ser selecionados em forms (0=ilimitado, 1=radio, N=checkbox)
- `group_id` - FK para ScoreGroup

### ScoreItem (❗ Model mais importante)

```sql
CREATE TABLE score_items (
  id UUID PRIMARY KEY,
  lab_test_code VARCHAR(100),           -- Liga ao código do exame laboratorial
  name VARCHAR(300) NOT NULL,
  unit VARCHAR(50),
  unit_conversion TEXT,                 -- Ex: "1 g/dL = 10 g/L"

  -- Filtros demográficos (aplicabilidade)
  gender VARCHAR(20) DEFAULT 'not_applicable',  -- not_applicable, male, female
  age_range_min INTEGER CHECK (age_range_min >= 0 AND age_range_min <= 150),
  age_range_max INTEGER CHECK (age_range_max >= 0 AND age_range_max <= 150),
  post_menopause BOOLEAN,               -- Apenas para gender=female

  -- Enriquecimento clínico
  clinical_relevance TEXT,              -- Para profissionais
  patient_explanation TEXT,             -- Para pacientes
  conduct TEXT,                         -- Conduta clínica
  last_review TIMESTAMP,                -- Auto-updated quando campos clínicos mudam

  points DOUBLE PRECISION,
  "order" INTEGER NOT NULL DEFAULT 0,

  -- Foreign Keys
  subgroup_id UUID NOT NULL REFERENCES score_subgroups(id),
  parent_item_id UUID REFERENCES score_items(id),  -- Self-referencing (hierarquia)

  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
```

**Campos especiais:**

- **lab_test_code**: Código único que liga ao `lab_test_definitions.code`
  - Permite associar resultados de exames automaticamente
  - Ex: "PLN585CE3E3" → "Hemoglobina"

- **Filtros demográficos**: Definem quando o item se aplica a um paciente
  - `gender`: "not_applicable" (padrão), "male", "female"
  - `age_range_min`/`max`: Faixa etária (ex: 18-65 anos)
  - `post_menopause`: true/false (apenas para mulheres)

- **Enriquecimento clínico**: 3 campos de texto rico
  - `clinical_relevance`: Explicação técnica (médicos)
  - `patient_explanation`: Linguagem simples (pacientes)
  - `conduct`: Conduta recomendada
  - `last_review`: Auto-updated via BeforeUpdate hook

- **parent_item_id**: Self-referencing para criar subitems
  - Ex: "Hemograma" → "Hemoglobina" → "Hemoglobina - Homens"

**Método de negócio AppliesToPatient:**

```go
func (si *ScoreItem) AppliesToPatient(patient *Patient) bool {
    // Filtro de gênero
    if si.Gender != nil && *si.Gender != "not_applicable" {
        if *si.Gender != string(patient.Gender) {
            return false
        }
    }

    // Filtro de idade
    if si.AgeRangeMin != nil && patient.Age < *si.AgeRangeMin {
        return false
    }
    if si.AgeRangeMax != nil && patient.Age > *si.AgeRangeMax {
        return false
    }

    // Filtro de menopausa (apenas mulheres)
    if patient.Gender == "female" && si.PostMenopause != nil {
        if patient.Menopause == nil || *si.PostMenopause != *patient.Menopause {
            return false
        }
    }

    return true
}
```

**Hook BeforeUpdate:**

```go
func (si *ScoreItem) BeforeUpdate(tx *gorm.DB) error {
    // Auto-update LastReview quando campos clínicos mudam
    if tx.Statement.Changed("ClinicalRelevance") ||
       tx.Statement.Changed("PatientExplanation") ||
       tx.Statement.Changed("Conduct") {
        now := time.Now()
        si.LastReview = &now
    }
    return nil
}
```

### ScoreLevel

```sql
CREATE TABLE score_levels (
  id UUID PRIMARY KEY,
  level INTEGER NOT NULL CHECK (level >= 0 AND level <= 6),
  name VARCHAR(500) NOT NULL,

  -- Limites numéricos
  lower_limit VARCHAR(50),
  upper_limit VARCHAR(50),
  operator VARCHAR(10) NOT NULL CHECK (operator IN ('=','>','>=','<','<=','between')),

  -- Enriquecimento clínico (mesmos campos do ScoreItem)
  clinical_relevance TEXT,
  patient_explanation TEXT,
  conduct TEXT,
  last_review TIMESTAMP,

  -- Foreign Key
  item_id UUID NOT NULL REFERENCES score_items(id),

  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
```

**Operadores:**

| Operador | LowerLimit | UpperLimit | Exemplo | Significado |
|----------|------------|------------|---------|-------------|
| `=` | Obrigatório | - | `lowerLimit="50"` | Valor = 50 |
| `>` | Obrigatório | - | `lowerLimit="70"` | Valor > 70 |
| `>=` | Obrigatório | - | `lowerLimit="70"` | Valor ≥ 70 |
| `<` | Obrigatório | - | `lowerLimit="40"` | Valor < 40 |
| `<=` | Obrigatório | - | `lowerLimit="40"` | Valor ≤ 40 |
| `between` | Obrigatório | Obrigatório | `lower="55", upper="70"` | 55 ≤ Valor ≤ 70 |

**Método de negócio EvaluatesTrue:**

```go
func (sl *ScoreLevel) EvaluatesTrue(value float64) bool {
    parseFloat := func(s *string) float64 {
        if s == nil {
            return 0
        }
        var result float64
        _, _ = fmt.Sscanf(*s, "%f", &result)
        return result
    }

    switch sl.Operator {
    case "=":
        return value == parseFloat(sl.LowerLimit)
    case ">":
        return value > parseFloat(sl.LowerLimit)
    case ">=":
        return value >= parseFloat(sl.LowerLimit)
    case "<":
        return value < parseFloat(sl.LowerLimit)
    case "<=":
        return value <= parseFloat(sl.LowerLimit)
    case "between":
        lower := parseFloat(sl.LowerLimit)
        upper := parseFloat(sl.UpperLimit)
        return value >= lower && value <= upper
    default:
        return false
    }
}
```

## Casos de Uso Práticos

### Caso 1: Hemoglobina com filtros demográficos

**Problema:** Valores de referência diferentes para homens, mulheres pré e pós-menopausa.

**Solução:** 3 ScoreItems separados

```sql
-- Item 1: Homens
INSERT INTO score_items (id, name, unit, gender, age_range_min, age_range_max, subgroup_id, "order")
VALUES (
  gen_random_uuid(),
  'Hemoglobina - Homens',
  'g/dL',
  'male',
  18,
  65,
  'uuid-subgroup-hemograma',
  1
);

-- Item 2: Mulheres pré-menopausa
INSERT INTO score_items (id, name, unit, gender, post_menopause, subgroup_id, "order")
VALUES (
  gen_random_uuid(),
  'Hemoglobina - Mulheres pré-menopausa',
  'g/dL',
  'female',
  false,
  'uuid-subgroup-hemograma',
  2
);

-- Item 3: Mulheres pós-menopausa
INSERT INTO score_items (id, name, unit, gender, post_menopause, subgroup_id, "order")
VALUES (
  gen_random_uuid(),
  'Hemoglobina - Mulheres pós-menopausa',
  'g/dL',
  'female',
  true,
  'uuid-subgroup-hemograma',
  3
);
```

Cada um com seus próprios níveis (valores de referência diferentes).

### Caso 2: FEVE (Fração de Ejeção) - Avaliação Cardiovascular

**Problema:** Indicador crítico de função cardíaca com múltiplos níveis.

```sql
-- ScoreItem
INSERT INTO score_items (id, name, unit, gender, subgroup_id, "order")
VALUES (
  '019c-feve-uuid',
  'FEVE (Fração de Ejeção do VE)',
  '%',
  'not_applicable',
  'uuid-subgroup-funcao-cardiaca',
  1
);

-- Níveis
INSERT INTO score_levels (id, level, name, lower_limit, upper_limit, operator, item_id) VALUES
  (gen_random_uuid(), 6, '≥ 70% (Hipercontrátil)', '70', NULL, '>=', '019c-feve-uuid'),
  (gen_random_uuid(), 5, '55-69% (Normal)', '55', '69', 'between', '019c-feve-uuid'),
  (gen_random_uuid(), 4, '50-54% (Limítrofe)', '50', '54', 'between', '019c-feve-uuid'),
  (gen_random_uuid(), 3, '40-49% (Disfunção leve)', '40', '49', 'between', '019c-feve-uuid'),
  (gen_random_uuid(), 2, '30-39% (Disfunção moderada)', '30', '39', 'between', '019c-feve-uuid'),
  (gen_random_uuid(), 1, '< 30% (Disfunção grave)', NULL, '30', '<', '019c-feve-uuid');
```

### Caso 3: Hierarquia de Items (Parent-Child)

**Problema:** "Pressão Arterial" → "PAS" e "PAD" como subitems

```sql
-- Parent
INSERT INTO score_items (id, name, unit, subgroup_id, "order")
VALUES (
  '019c-pa-parent',
  'Pressão Arterial',
  'mmHg',
  'uuid-subgroup-sinais-vitais',
  1
);

-- Child 1: PAS
INSERT INTO score_items (id, name, unit, parent_item_id, subgroup_id, "order")
VALUES (
  '019c-pa-pas',
  'PAS (Sistólica)',
  'mmHg',
  '019c-pa-parent',
  'uuid-subgroup-sinais-vitais',
  1
);

-- Child 2: PAD
INSERT INTO score_items (id, name, unit, parent_item_id, subgroup_id, "order")
VALUES (
  '019c-pa-pad',
  'PAD (Diastólica)',
  'mmHg',
  '019c-pa-parent',
  'uuid-subgroup-sinais-vitais',
  2
);
```

## Workflows Comuns

### Adicionar novo ScoreItem com níveis

**Via psql (recomendado para desenvolvimento):**

Ver arquivo `.claude/workflows/database-ops.md` para exemplos completos.

**Resumo:**

```bash
docker compose exec db psql -U plenya_user -d plenya_db << 'EOF'
BEGIN;

-- Inserir item
INSERT INTO score_items (id, name, unit, gender, subgroup_id, "order", created_at, updated_at)
VALUES (gen_random_uuid(), 'Nome', 'unidade', 'male', 'uuid-subgroup', 1, NOW(), NOW())
RETURNING id;

-- Copiar o UUID retornado e inserir níveis
INSERT INTO score_levels (id, level, name, lower_limit, operator, item_id, created_at, updated_at)
VALUES
  (gen_random_uuid(), 6, 'Ótimo', '70', '>=', 'uuid-do-item-acima', NOW(), NOW()),
  (gen_random_uuid(), 5, 'Normal', '55', 'between', 'uuid-do-item-acima', NOW(), NOW());

COMMIT;
EOF
```

### Duplicar ScoreItem mudando filtros

**Cenário:** Já existe "Hemoglobina - Homens", criar "Hemoglobina - Mulheres pós-menopausa"

```bash
docker compose exec db psql -U plenya_user -d plenya_db << 'EOF'
-- Passo 1: Duplicar item
INSERT INTO score_items (
  id, name, unit, gender, age_range_min, age_range_max, post_menopause,
  lab_test_code, unit_conversion, clinical_relevance, patient_explanation, conduct,
  points, subgroup_id, parent_item_id, "order", created_at, updated_at
)
SELECT
  gen_random_uuid(),
  'Hemoglobina - Mulheres pós-menopausa',
  unit,
  'female',
  age_range_min,
  age_range_max,
  true,
  lab_test_code,
  unit_conversion,
  clinical_relevance,
  patient_explanation,
  conduct,
  points,
  subgroup_id,
  parent_item_id,
  "order" + 1,
  NOW(),
  NOW()
FROM score_items
WHERE name = 'Hemoglobina - Homens'
RETURNING id;

-- Passo 2: Copiar o UUID retornado (ex: 019c-novo-uuid)
-- Passo 3: Duplicar níveis para o novo item
INSERT INTO score_levels (
  id, level, name, lower_limit, upper_limit, operator,
  clinical_relevance, patient_explanation, conduct,
  item_id, created_at, updated_at
)
SELECT
  gen_random_uuid(),
  level,
  name,
  lower_limit,
  upper_limit,
  operator,
  clinical_relevance,
  patient_explanation,
  conduct,
  '019c-novo-uuid',  -- UUID do novo item
  NOW(),
  NOW()
FROM score_levels
WHERE item_id = (SELECT id FROM score_items WHERE name = 'Hemoglobina - Homens');
EOF
```

### Atualizar enriquecimento clínico

```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
UPDATE score_items
SET
  clinical_relevance = 'Hemoglobina baixa indica anemia, associada a fadiga...',
  patient_explanation = 'Hemoglobina transporta oxigênio. Valores baixos causam cansaço',
  conduct = 'Investigar causa (ferro, B12, folato). Suplementar conforme indicado',
  last_review = NOW()
WHERE name = 'Hemoglobina - Homens';
"
```

**Nota:** `last_review` é atualizado automaticamente pelo BeforeUpdate hook quando usar GORM, mas no SQL direto precisa setar manualmente.

## Relação M:N com Articles

ScoreItems podem ter múltiplos artigos científicos de referência.

```sql
-- Tabela de junção
CREATE TABLE article_score_items (
  article_id UUID REFERENCES articles(id) ON DELETE CASCADE,
  score_item_id UUID REFERENCES score_items(id) ON DELETE CASCADE,
  PRIMARY KEY (article_id, score_item_id)
);

-- Adicionar artigo a um item
INSERT INTO article_score_items (article_id, score_item_id)
VALUES (
  'uuid-do-artigo',
  'uuid-do-score-item'
);
```

**Query para buscar items com artigos:**

```sql
SELECT
  si.id,
  si.name,
  json_agg(json_build_object('id', a.id, 'title', a.title, 'doi', a.doi)) AS articles
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
LEFT JOIN articles a ON asi.article_id = a.id
WHERE si.id = 'uuid-do-item'
GROUP BY si.id, si.name;
```

## Queries Úteis

### Buscar items aplicáveis a um paciente

```sql
-- Paciente: homem, 45 anos
SELECT *
FROM score_items
WHERE
  (gender = 'not_applicable' OR gender = 'male')
  AND (age_range_min IS NULL OR age_range_min <= 45)
  AND (age_range_max IS NULL OR age_range_max >= 45)
  AND deleted_at IS NULL
ORDER BY "order" ASC;
```

### Árvore completa de grupos

```sql
SELECT
  g.id AS group_id,
  g.name AS group_name,
  json_agg(
    json_build_object(
      'id', sg.id,
      'name', sg.name,
      'items', (
        SELECT json_agg(
          json_build_object(
            'id', si.id,
            'name', si.name,
            'unit', si.unit,
            'levels', (
              SELECT json_agg(
                json_build_object('level', sl.level, 'name', sl.name)
                ORDER BY sl.level ASC
              )
              FROM score_levels sl
              WHERE sl.item_id = si.id AND sl.deleted_at IS NULL
            )
          )
          ORDER BY si."order" ASC
        )
        FROM score_items si
        WHERE si.subgroup_id = sg.id AND si.deleted_at IS NULL
      )
    )
    ORDER BY sg."order" ASC
  ) AS subgroups
FROM score_groups g
LEFT JOIN score_subgroups sg ON g.id = sg.group_id AND sg.deleted_at IS NULL
WHERE g.deleted_at IS NULL
GROUP BY g.id, g.name
ORDER BY g."order" ASC;
```

### Items sem níveis (data quality check)

```sql
SELECT si.id, si.name
FROM score_items si
LEFT JOIN score_levels sl ON si.id = sl.item_id AND sl.deleted_at IS NULL
WHERE si.deleted_at IS NULL
GROUP BY si.id, si.name
HAVING COUNT(sl.id) = 0;
```

## Validações Importantes

1. **Sempre ter ao menos 1 nível:** ScoreItem sem níveis não é útil
2. **Operadores consistentes:** `between` DEVE ter lower_limit E upper_limit
3. **Níveis únicos:** Não pode ter dois níveis com mesmo `level` no mesmo item
4. **Gender válido:** "not_applicable", "male", "female"
5. **PostMenopause só para female:** Se gender != "female", post_menopause deve ser NULL
6. **Age ranges lógicos:** age_range_min <= age_range_max
