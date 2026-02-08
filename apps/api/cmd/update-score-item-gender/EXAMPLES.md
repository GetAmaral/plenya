# Exemplos de Uso - Score Item Gender Update

## Casos de Uso

### 1. Primeira execução (configuração inicial)

Executar após popular o banco com dados iniciais de `score_items` que não têm o campo `gender` configurado.

```bash
# Via Docker Compose (Recomendado)
docker compose exec api make update-gender

# Ou diretamente
docker compose exec api go run cmd/update-score-item-gender/main.go
```

### 2. Após adicionar novos items manualmente

Se você adicionou novos `score_items` via SQL ou API e precisa atualizar o campo `gender`:

```bash
# Dentro do container
docker compose exec api make update-gender
```

### 3. Corrigir dados inconsistentes

Se alguns items foram criados com `gender` incorreto:

```bash
# O script é idempotente - pode rodar múltiplas vezes
docker compose exec api make update-gender
```

### 4. Validação antes de rodar em produção

Testar em ambiente de staging antes de produção:

```bash
# 1. Testar a lógica de detecção
docker compose exec api make test-gender

# 2. Fazer backup do banco
docker compose exec db pg_dump -U plenya_user plenya_db > backup_before_gender_update.sql

# 3. Executar script
docker compose exec api make update-gender

# 4. Verificar resultados
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT gender, COUNT(*) FROM score_items GROUP BY gender;"
```

## Verificações no Banco de Dados

### Verificar distribuição de gêneros

```sql
SELECT
  gender,
  COUNT(*) as count,
  ROUND(COUNT(*) * 100.0 / SUM(COUNT(*)) OVER(), 2) as percentage
FROM score_items
WHERE deleted_at IS NULL
GROUP BY gender
ORDER BY count DESC;
```

### Listar items masculinos

```sql
SELECT id, name, gender
FROM score_items
WHERE gender = 'male'
  AND deleted_at IS NULL
ORDER BY name;
```

### Listar items femininos

```sql
SELECT id, name, gender
FROM score_items
WHERE gender = 'female'
  AND deleted_at IS NULL
ORDER BY name;
```

### Buscar possíveis falsos negativos

Items que deveriam ter gênero mas estão como `not_applicable`:

```sql
SELECT id, name, gender
FROM score_items
WHERE gender = 'not_applicable'
  AND (
    LOWER(name) LIKE '%homem%'
    OR LOWER(name) LIKE '%mulher%'
    OR LOWER(name) LIKE '%masculin%'
    OR LOWER(name) LIKE '%feminin%'
  )
  AND deleted_at IS NULL;
```

### Verificar items atualizados recentemente

```sql
SELECT id, name, gender, updated_at
FROM score_items
WHERE updated_at > NOW() - INTERVAL '1 hour'
  AND deleted_at IS NULL
ORDER BY updated_at DESC;
```

## Integração com CI/CD

### GitHub Actions

```yaml
name: Update Score Item Gender

on:
  workflow_dispatch: # Trigger manual
  schedule:
    - cron: '0 2 * * 0' # Todo domingo às 2h

jobs:
  update-gender:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.25'

      - name: Load environment
        run: |
          echo "${{ secrets.ENV_FILE }}" > .env

      - name: Run gender update
        run: |
          cd apps/api
          make update-gender

      - name: Verify results
        run: |
          # Adicionar verificação de resultados aqui
```

### Como migration de dados

Se preferir executar como parte do processo de migration:

```go
// apps/api/internal/database/migrations/20260208000000_update_score_item_gender.go

package migrations

import (
    "github.com/plenya/api/internal/models"
    "gorm.io/gorm"
    "strings"
)

func UpdateScoreItemGender(db *gorm.DB) error {
    var items []models.ScoreItem
    if err := db.Find(&items).Error; err != nil {
        return err
    }

    for _, item := range items {
        gender := detectGender(item.Name)
        if err := db.Model(&item).Update("gender", gender).Error; err != nil {
            return err
        }
    }

    return nil
}

func detectGender(name string) string {
    nameLower := strings.ToLower(name)

    maleKeywords := []string{"homem", "homens", "masculino"}
    for _, kw := range maleKeywords {
        if strings.Contains(nameLower, kw) {
            return "male"
        }
    }

    femaleKeywords := []string{"mulher", "mulheres", "feminino"}
    for _, kw := range femaleKeywords {
        if strings.Contains(nameLower, kw) {
            return "female"
        }
    }

    return "not_applicable"
}
```

## Troubleshooting

### Problema: Script não encontra items

```bash
# Verificar se há items no banco
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT COUNT(*) FROM score_items;"

# Se retornar 0, popular o banco com seed data
docker compose exec api go run cmd/seed/main.go
```

### Problema: Erros de conexão

```bash
# Verificar se PostgreSQL está rodando
docker compose ps db

# Verificar logs do banco
docker compose logs db

# Testar conexão manualmente
docker compose exec api psql \
  -h db \
  -U plenya_user \
  -d plenya_db \
  -c "SELECT version();"
```

### Problema: Permissões no script shell

```bash
# Dar permissão de execução
chmod +x apps/api/cmd/update-score-item-gender/run.sh

# Executar
./apps/api/cmd/update-score-item-gender/run.sh
```

## Performance

O script processa aproximadamente:

- **150 items**: ~0.5 segundos
- **500 items**: ~1.5 segundos
- **1000 items**: ~3 segundos

Benchmark:

```bash
cd apps/api
go test -bench=. -benchmem cmd/update-score-item-gender/
```

## Logs e Auditoria

### Salvar output do script

```bash
docker compose exec api make update-gender | tee gender_update_$(date +%Y%m%d_%H%M%S).log
```

### Criar audit log das mudanças

```sql
-- Criar tabela de auditoria (opcional)
CREATE TABLE score_item_gender_audit (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    score_item_id UUID NOT NULL,
    old_gender VARCHAR(20),
    new_gender VARCHAR(20),
    changed_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (score_item_id) REFERENCES score_items(id)
);

-- Trigger para log automático
CREATE OR REPLACE FUNCTION log_gender_change()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.gender IS DISTINCT FROM NEW.gender THEN
        INSERT INTO score_item_gender_audit
            (score_item_id, old_gender, new_gender)
        VALUES
            (NEW.id, OLD.gender, NEW.gender);
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER score_item_gender_change
    AFTER UPDATE ON score_items
    FOR EACH ROW
    EXECUTE FUNCTION log_gender_change();
```

## Notas Importantes

1. **Backup antes de executar em produção**: Sempre fazer backup do banco antes de rodar scripts de atualização em massa.

2. **Idempotência**: O script pode ser executado múltiplas vezes sem causar problemas.

3. **Soft-deleted items**: Por padrão, o script atualiza TODOS os items, incluindo soft-deleted. Isso garante consistência caso um item seja restaurado.

4. **Case-insensitive**: A detecção é case-insensitive, então "HOMEM", "Homem" e "homem" são tratados da mesma forma.

5. **Prioridade de detecção**: Se um nome contém tanto palavras masculinas quanto femininas, a primeira detectada (masculino) tem prioridade.
