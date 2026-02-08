# Score Item Age Range Update Script

Script utilitário para detectar e atualizar automaticamente os campos `ageRangeMin` e `ageRangeMax` nos `score_items` baseado em padrões encontrados no campo `name`.

## Visão Geral

O sistema EMR Plenya armazena parâmetros clínicos (`score_items`) que podem ter faixas etárias específicas. Este script analisa o campo `name` de cada item e preenche automaticamente os campos `age_range_min` e `age_range_max` quando detecta padrões de idade.

## Resumo Visual

```
Score Item Name                          Padrão           age_range_min  age_range_max
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Hemoglobina - Adultos 18-65 anos         X-Y anos         18             65
Colesterol < 20 anos                     < X anos         0              20
PSA > 50 anos                            > X anos         50             150
Mamografia 40+                           X+ anos          40             150
Estradiol - Pré-menopausa                (ignorado)       NULL           NULL
FEVE                                     (sem padrão)     NULL           NULL
```

## Padrões Detectados

### 1. Range "X-Y anos"
```
"Hemoglobina - 18-65 anos"        → min: 18, max: 65
"Colesterol 20 a 60 anos"         → min: 20, max: 60
"TSH adultos (30-70 anos)"        → min: 30, max: 70
```

### 2. Menor que "< X anos"
```
"Hemoglobina < 18 anos"           → min: 0, max: 18
"Glicose menor que 20 anos"       → min: 0, max: 20
"TSH abaixo de 15 anos"           → min: 0, max: 15
```

### 3. Maior que "> X anos"
```
"PSA > 50 anos"                   → min: 50, max: 150
"Glicose maior que 65 anos"       → min: 65, max: 150
"Colesterol acima de 60 anos"     → min: 60, max: 150
```

### 4. Mais de "X+ anos"
```
"Mamografia 40+ anos"             → min: 40, max: 150
"PSA 50 +"                        → min: 50, max: 150
"Colonoscopia 45+anos"            → min: 45, max: 150
```

**Nota:** O valor máximo é 150 devido à constraint do banco de dados (`age_range_max <= 150`), que reflete o limite biológico humano.

## Termos Ignorados

O script **NÃO** detecta idade em termos que representam fases fisiológicas (não são idades específicas):

```
Pré-menopausa
Pós-menopausa
Gestação / Gestante / Grávida
Lactante / Lactação
Puerpério
Fase folicular / Fase lútea
Ovulatório
Ciclo menstrual
```

Esses termos são ignorados porque representam contextos clínicos específicos que não podem ser reduzidos a uma faixa etária simples.

## Características

- **Case-insensitive:** Funciona com MAIÚSCULAS, minúsculas ou MiStUrAdO
- **Regex otimizado:** Padrões compilados uma vez (performance)
- **Idempotente:** Pode rodar múltiplas vezes sem efeitos colaterais
- **Validação:** Apenas ranges válidos (0-150 anos, min < max)
- **Logs detalhados:** Mostra exatamente o que foi detectado e atualizado
- **Seguro:** Não atualiza se o valor já estiver correto

## Como Usar

### Via Makefile (Recomendado)

```bash
cd apps/api
make update-age-range
```

### Diretamente com Go

```bash
cd apps/api
go run cmd/update-score-item-age-range/main.go
```

### Com Docker

```bash
docker compose exec api make update-age-range
```

## Output Esperado

```
=== Score Item Age Range Update Script ===

✅ Database connected

📊 Found 342 score items

Processing score items:
----------------------------------------
[15/342] ✓ Hemoglobina - Adultos 18-65 anos
        Pattern: 18-65 anos
        Range: null → 18-65

[42/342] ✓ PSA - Homens > 50 anos
        Pattern: > 50 anos
        Range: null → 50-∞

[89/342] ✓ Glicose < 18 anos
        Pattern: < 18 anos
        Range: null → 0-18

[120/342] ✓ Mamografia 40+ anos
        Pattern: 40+ anos
        Range: null → 40-∞

----------------------------------------
=== Summary ===
Total processed:     342
Range detected:      87
Range updated:       85
Unchanged:           2
Ignored (no range):  255
Errors:              0

✅ Script completed successfully!
```

## Testes

```bash
cd apps/api

# Rodar testes
make test-age-range

# Ou diretamente
go test -v cmd/update-score-item-age-range/
```

### Coverage

```bash
go test -cover cmd/update-score-item-age-range/
```

## Casos de Uso

### 1. Após Importação em Lote
Quando importar score_items de planilhas/CSVs sem age range:
```bash
make update-age-range
```

### 2. Correção de Dados Existentes
Quando descobrir que os ranges estão incorretos:
```bash
# Rodar script
make update-age-range

# Verificar mudanças no banco
psql -d plenya_db -c "
  SELECT name, age_range_min, age_range_max
  FROM score_items
  WHERE age_range_min IS NOT NULL
  ORDER BY age_range_min;
"
```

### 3. Validação Pós-Deploy
Após deploy de novos score_items via migrations:
```bash
make migrate-up
make update-age-range
```

## Integração com Sistema

### Opção 1: Hook no Model (Automático)
Adicionar ao `apps/api/internal/models/score_item.go`:

```go
func (si *ScoreItem) BeforeSave(tx *gorm.DB) error {
    // Auto-detect age range se o nome foi alterado
    if tx.Statement.Changed("Name") {
        ageRange := detectAgeRangeFromName(si.Name)
        if ageRange.Detected {
            si.AgeRangeMin = ageRange.Min
            si.AgeRangeMax = ageRange.Max
        }
    }
    return nil
}
```

### Opção 2: API Endpoint (Manual)
```go
// POST /api/v1/admin/score-items/update-age-range
func (h *AdminHandler) UpdateAgeRange(c *fiber.Ctx) error {
    // Chama lógica do script
    result := runAgeRangeUpdate(h.db)
    return c.JSON(result)
}
```

### Opção 3: Scheduled Job (Periódico)
```go
// Rodar diariamente às 3h da manhã
cron := cron.New()
cron.AddFunc("0 3 * * *", func() {
    runAgeRangeUpdate(db)
})
```

## Limitações

1. **Apenas português:** Padrões em inglês não são detectados
2. **Formato específico:** Requer "anos" ou "ano" no texto
3. **Ranges implícitos:** "Adultos" sem número não é detectado
4. **Contexto clínico:** Pré/Pós-menopausa não são convertidos

## Exemplos Reais

### Casos Detectados ✅
```
"Hemoglobina - Homens 18-65 anos"          → 18-65
"PSA - Screening > 50 anos"                → 50-150
"Glicose pediátrica < 12 anos"             → 0-12
"Colesterol LDL (adultos 20 a 60 anos)"    → 20-60
"Mamografia de rastreio 40+"               → 40-150
```

### Casos Ignorados ❌
```
"Estradiol - Pré-menopausa"                → ignorado (fase, não idade)
"Progesterona - Fase lútea"                → ignorado (contexto clínico)
"hCG - Gestação"                           → ignorado (estado fisiológico)
"Hemoglobina - Homens"                     → ignorado (sem idade)
"FEVE (Fração de Ejeção)"                  → ignorado (sem idade)
```

## SQL Queries Úteis

### Visualizar todos os ranges detectados
```sql
SELECT
    name,
    age_range_min,
    age_range_max,
    CASE
        WHEN age_range_min IS NULL AND age_range_max IS NULL THEN 'No range'
        WHEN age_range_max = 999 THEN age_range_min || '+ anos'
        ELSE age_range_min || '-' || age_range_max || ' anos'
    END as range_display
FROM score_items
WHERE deleted_at IS NULL
ORDER BY age_range_min NULLS LAST, name;
```

### Estatísticas de ranges
```sql
SELECT
    CASE
        WHEN age_range_min IS NULL AND age_range_max IS NULL THEN 'Sem range'
        WHEN age_range_max = 999 THEN age_range_min || '+ anos'
        ELSE age_range_min || '-' || age_range_max || ' anos'
    END as range_group,
    COUNT(*) as count
FROM score_items
WHERE deleted_at IS NULL
GROUP BY range_group
ORDER BY MIN(age_range_min) NULLS LAST;
```

### Itens com range detectado
```sql
SELECT
    name,
    age_range_min || '-' || age_range_max as range
FROM score_items
WHERE age_range_min IS NOT NULL
  AND deleted_at IS NULL
ORDER BY age_range_min, age_range_max;
```

### Verificar padrões não detectados
```sql
-- Buscar nomes que contenham números mas não tenham range
SELECT name
FROM score_items
WHERE name ~ '\d+'
  AND age_range_min IS NULL
  AND deleted_at IS NULL
ORDER BY name;
```

## Troubleshooting

### Erro: "Failed to connect to database"
```bash
# Verificar se PostgreSQL está rodando
docker compose ps

# Verificar .env
cat apps/api/.env | grep DB_
```

### Erro: "No score items found"
```bash
# Verificar se há dados no banco
psql -d plenya_db -c "SELECT COUNT(*) FROM score_items;"

# Rodar migrations
make migrate-up
```

### Age range não foi detectado
```bash
# Verificar o nome do item
psql -d plenya_db -c "SELECT id, name FROM score_items WHERE id = 'UUID_AQUI';"

# Testar manualmente
go run cmd/update-score-item-age-range/main.go | grep "Nome do Item"
```

## Manutenção

### Adicionar Novo Padrão
1. Editar `main.go` e adicionar regex:
```go
newPattern := regexp.MustCompile(`(?i)seu-padrao-aqui`)
```

2. Adicionar lógica em `detectAgeRange()`:
```go
if matches := newPattern.FindStringSubmatch(name); matches != nil {
    // processar matches
}
```

3. Adicionar testes em `main_test.go`:
```go
{
    name:        "Novo padrão",
    input:       "Exemplo do novo padrão",
    expectMin:   intPtr(10),
    expectMax:   intPtr(20),
    expectFound: true,
},
```

### Adicionar Termo a Ignorar
Editar `ignoreTerms` em `main.go`:
```go
var ignoreTerms = []string{
    // ... termos existentes ...
    "novo termo",
    "outro termo",
}
```

## Referências

- Model: `apps/api/internal/models/score_item.go`
- Migrations: `apps/api/database/migrations/`
- Script de Gender: `apps/api/cmd/update-score-item-gender/`

## Changelog

### v1.0.0 (2026-02-08)
- ✨ Detecção inicial de age ranges
- ✨ Suporte a 5 padrões diferentes
- ✨ Lista de termos a ignorar (fases clínicas)
- ✅ Testes unitários completos
- 📝 Documentação completa
