# Score Item Gender Update Script

Script standalone para atualizar o campo `gender` de todos os `score_items` no banco de dados, baseado em análise de palavras-chave no campo `name`.

## Funcionalidade

O script:

1. Conecta ao banco de dados PostgreSQL
2. Busca todos os `score_items`
3. Analisa o campo `name` de cada item (case-insensitive)
4. Identifica palavras-chave relacionadas a gênero:
   - **Masculino**: "homem", "homens", "masculino", "masculina", etc.
   - **Feminino**: "mulher", "mulheres", "feminino", "feminina", etc.
5. Atualiza o campo `gender` com:
   - `"male"` - se contiver palavras relacionadas a masculino
   - `"female"` - se contiver palavras relacionadas a feminino
   - `"not_applicable"` - se não houver indicação de gênero
6. Exibe um resumo das alterações

## Pré-requisitos

- Go 1.25+
- PostgreSQL em execução
- Variáveis de ambiente configuradas (`.env` ou export)

## Variáveis de Ambiente

O script usa as mesmas variáveis de ambiente do backend principal:

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=plenya_user
DB_PASSWORD=senha_segura
DB_NAME=plenya_db
```

## Como Executar

### Opção 1: Via Docker Compose (Recomendado)

```bash
# Executar dentro do container da API
docker compose exec api go run cmd/update-score-item-gender/main.go
```

### Opção 2: Diretamente no Host

```bash
# Dentro do diretório apps/api
cd apps/api

# Carregar variáveis de ambiente (se usar .env)
export $(cat .env | xargs)

# Executar script
go run cmd/update-score-item-gender/main.go
```

### Opção 3: Build e Executar

```bash
cd apps/api

# Build
go build -o bin/update-gender cmd/update-score-item-gender/main.go

# Executar
./bin/update-gender
```

## Exemplo de Saída

```
=== Score Item Gender Update Script ===

✅ Database connected

📊 Found 150 score items

Processing score items:
----------------------------------------
[1/150] ✓ Hemoglobina - Homens
        not_applicable → male
[2/150] ✓ Hemoglobina - Mulheres
        not_applicable → female
[3/150] ✓ Glicose em jejum
        not_applicable → not_applicable
...

----------------------------------------
=== Summary ===
Total processed:     150
Updated to 'male':   25
Updated to 'female': 30
Updated to 'not_applicable': 10
Unchanged:           85
Errors:              0

✅ Script completed successfully!
```

## Palavras-chave Detectadas

### Masculino (male)
- homem, homens
- masculino, masculina, masculinos, masculinas
- homem adulto, sexo masculino
- dos homens, para homens, em homens, no homem, nos homens

### Feminino (female)
- mulher, mulheres
- feminino, feminina, femininos, femininas
- mulher adulta, sexo feminino
- das mulheres, para mulheres, em mulheres, na mulher, nas mulheres

## Segurança

- **Idempotente**: Pode ser executado múltiplas vezes sem problemas
- **Não destrutivo**: Apenas atualiza o campo `gender`, não deleta dados
- **Read-only em outros campos**: Não modifica nenhum outro campo
- **Transaction-safe**: Usa transações GORM individuais por update

## Troubleshooting

### Erro de conexão com banco

```
❌ Failed to connect to database: dial tcp: connect: connection refused
```

**Solução**: Verificar se PostgreSQL está em execução e variáveis de ambiente corretas.

```bash
# Verificar containers Docker
docker compose ps

# Ver logs do PostgreSQL
docker compose logs db
```

### Nenhum item encontrado

```
⚠️  No score items found in database
```

**Solução**: Verificar se o banco foi populado com dados de seed.

```bash
# Verificar no PostgreSQL
docker compose exec db psql -U plenya_user -d plenya_db -c "SELECT COUNT(*) FROM score_items;"
```

## Integração como Migration/Seed

Para executar automaticamente durante o deploy, você pode:

1. **Como migration manual** (recomendado para dados existentes):
   ```bash
   atlas migrate new update_score_item_gender --sql
   # Adicionar chamada ao script no SQL
   ```

2. **Como seed no código**:
   ```go
   // Copiar a lógica de detectGender() para internal/database/seeds.go
   // e executar durante bootstrap inicial
   ```

3. **Como job agendado**:
   ```go
   // Adicionar ao scheduler para executar periodicamente
   // (útil se novos items forem adicionados manualmente)
   ```

## Notas

- O script é **case-insensitive** nas buscas
- A ordem de verificação é: masculino primeiro, depois feminino
- Se um nome contém ambas palavras-chave (ex: "Comparativo Homens e Mulheres"), retorna a primeira detectada (masculino)
- Soft-deleted items **são incluídos** na análise

## Autor

Script criado para automação de atualização de metadados clínicos do Plenya EMR.
