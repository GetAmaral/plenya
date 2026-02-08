# Quick Start - Score Item Gender Update

## 1 Minuto: Executar o Script

```bash
# Via Docker Compose
docker compose exec api make update-gender
```

Pronto! O script irá:
- Conectar ao banco de dados
- Analisar todos os score_items
- Atualizar o campo gender baseado no nome
- Exibir um resumo das mudanças

---

## 5 Minutos: Com Validação

```bash
# 1. Fazer backup (recomendado)
docker compose exec db pg_dump -U plenya_user plenya_db > backup_$(date +%Y%m%d).sql

# 2. Executar script
docker compose exec api make update-gender

# 3. Verificar resultados
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT gender, COUNT(*) FROM score_items GROUP BY gender;"
```

---

## 10 Minutos: Teste Completo

```bash
# 1. Testar lógica de detecção
docker compose exec api make test-gender

# 2. Backup do banco
docker compose exec db pg_dump -U plenya_user plenya_db > backup_before_gender.sql

# 3. Executar script (com log)
docker compose exec api make update-gender | tee gender_update_$(date +%Y%m%d_%H%M%S).log

# 4. Verificar distribuição
docker compose exec db psql -U plenya_user -d plenya_db << EOF
SELECT
  gender,
  COUNT(*) as count,
  ROUND(COUNT(*) * 100.0 / SUM(COUNT(*)) OVER(), 2) as percentage
FROM score_items
WHERE deleted_at IS NULL
GROUP BY gender;
EOF

# 5. Verificar se há falsos negativos
docker compose exec db psql -U plenya_user -d plenya_db << EOF
SELECT id, name, gender
FROM score_items
WHERE gender = 'not_applicable'
  AND (
    LOWER(name) LIKE '%homem%' OR
    LOWER(name) LIKE '%mulher%'
  )
LIMIT 10;
EOF
```

---

## Comandos Úteis

### Executar

```bash
# Via Makefile (recomendado)
docker compose exec api make update-gender

# Via go run
docker compose exec api go run cmd/update-score-item-gender/main.go

# Via shell helper
docker compose exec api ./cmd/update-score-item-gender/run.sh
```

### Testar

```bash
# Rodar testes
docker compose exec api make test-gender

# Com coverage
docker compose exec api go test -cover cmd/update-score-item-gender/

# Com verbose
docker compose exec api go test -v cmd/update-score-item-gender/
```

### Verificar no Banco

```bash
# Total por gênero
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT gender, COUNT(*) FROM score_items GROUP BY gender;"

# Listar masculinos
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT name FROM score_items WHERE gender = 'male' LIMIT 10;"

# Listar femininos
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT name FROM score_items WHERE gender = 'female' LIMIT 10;"
```

---

## Troubleshooting Rápido

### Erro: "No score items found"

```bash
# Verificar se há dados no banco
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT COUNT(*) FROM score_items;"

# Se retornar 0, popular com seed data (se disponível)
```

### Erro: "Failed to connect to database"

```bash
# Verificar se PostgreSQL está rodando
docker compose ps db

# Ver logs
docker compose logs db

# Reiniciar containers
docker compose restart
```

### Erro: "Permission denied"

```bash
# Dar permissão ao script shell
chmod +x apps/api/cmd/update-score-item-gender/run.sh
```

---

## Opções de Integração Rápida

### Como Seed (Automático)

Adicione ao `apps/api/internal/database/database.go`:

```go
func SeedScoreItemGender(db *gorm.DB) error {
    // Ver seed_example.go para implementação completa
}
```

Depois chame no `main.go` após AutoMigrate():

```go
if cfg.Server.Environment == "development" {
    database.AutoMigrate()
    database.SeedScoreItemGender(database.DB) // ⭐ Adicionar
}
```

### Como Migration SQL

```bash
# Criar migration
atlas migrate new update_score_item_gender --env dev

# Editar SQL gerado (ver INTEGRATION.md para exemplo)
# Aplicar
atlas migrate apply --env dev
```

---

## Estrutura de Arquivos

```
apps/api/cmd/update-score-item-gender/
├── main.go              # ⭐ Script principal
├── main_test.go         # Testes
├── run.sh               # Helper
├── QUICKSTART.md        # Este arquivo
├── README.md            # Documentação completa
├── EXAMPLES.md          # Casos de uso
├── INTEGRATION.md       # Guia de integração
├── SUMMARY.md           # Sumário executivo
└── seed_example.go      # Exemplos de código
```

---

## Próximos Passos

Depois de executar com sucesso:

1. **Ler** `README.md` para entender detalhes
2. **Ver** `EXAMPLES.md` para casos de uso avançados
3. **Considerar** `INTEGRATION.md` para automatizar

---

## Checklist de Produção

Antes de usar em produção:

- [ ] Testar em ambiente de staging
- [ ] Fazer backup completo do banco
- [ ] Executar durante janela de manutenção
- [ ] Validar resultados com queries SQL
- [ ] Monitorar logs para erros
- [ ] Documentar execução

---

## Suporte

- **Documentação técnica**: `README.md`
- **Casos de uso**: `EXAMPLES.md`
- **Integrações**: `INTEGRATION.md`
- **Visão geral**: `SUMMARY.md`

---

**Tempo estimado de execução**: < 5 segundos para 500 items

**Segurança**: Idempotente, não destrutivo, transaction-safe

**Status**: ✅ Pronto para uso
