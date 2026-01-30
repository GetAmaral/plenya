# COMO EXECUTAR O BATCH FINAL 1

**Missão:** Enriquecer 45 items de exames laboratoriais no banco de dados

---

## ⚡ Execução Rápida (3 comandos)

```bash
# 1. Aplicar enrichments ao banco (45 UPDATEs em 1 transação)
docker compose exec -T db psql -U plenya_user -d plenya_db < batch_final_1_exames_A.sql

# 2. Validar aplicação (9 queries de verificação)
docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/validate_batch_final_1.sql

# 3. Verificar no frontend
open http://localhost:3000/scores
```

---

## 📋 Output Esperado

### Comando 1 (aplicar SQL)
```
BEGIN
UPDATE 1
UPDATE 1
UPDATE 1
... (45 vezes total)
UPDATE 1
COMMIT
```

**Status:** ✅ Se ver "COMMIT" no final, sucesso!

### Comando 2 (validação)
```
===================================================================
VALIDAÇÃO BATCH FINAL 1 - EXAMES LABORATORIAIS PARTE A
===================================================================

1. Total de items atualizados hoje (esperado: 45)
-------------------------------------------------------------------
 total_updated
---------------
            45
(1 row)

2. Items com enrichment específico detalhado (esperado: 3)
-------------------------------------------------------------------
              name              | interp_chars | desc_chars | artigos | last_review
--------------------------------+--------------+------------+---------+-------------
 Doppler Carótidas - Estenose  |          468 |       1187 |       4 | 2026-01-28
 Hidrogênio expirado           |          505 |       1176 |       4 | 2026-01-28
 Mamografia - Densidade Mamária|          553 |        970 |       4 | 2026-01-28
(3 rows)

...
```

**Status:** ✅ Conferir números nas tabelas

---

## 🔍 Verificação Manual (opcional)

Se quiser verificar manualmente:

```bash
# Entrar no container do banco
docker compose exec db psql -U plenya_user -d plenya_db

# Rodar query de verificação
SELECT
  COUNT(*) as total_items,
  COUNT(*) FILTER (WHERE LENGTH(interpretation) > 500) as especificos,
  COUNT(*) FILTER (WHERE LENGTH(interpretation) BETWEEN 200 AND 500) as padrao,
  ROUND(AVG(jsonb_array_length(articles))) as avg_artigos,
  MIN(last_review)::date as primeira_atualizacao,
  MAX(last_review)::date as ultima_atualizacao
FROM score_items
WHERE last_review >= CURRENT_DATE;

# Esperado:
# total_items | especificos | padrao | avg_artigos | primeira_atualizacao | ultima_atualizacao
#          45 |           3 |     42 |           3 | 2026-01-28           | 2026-01-28

# Sair
\q
```

---

## 📊 Checklist de Validação

Após executar os comandos, conferir:

- [ ] **45 items atualizados:** Query 1 retornou 45
- [ ] **3 enrichments específicos:** Query 2 retornou Mamografia, H2 Expirado, Doppler Carótidas
- [ ] **Campos preenchidos:** Query 8 retornou 0 nulls em todos os campos
- [ ] **JSON válido:** Query 5 mostrou artigos corretamente
- [ ] **Frontend funcionando:** Acessar http://localhost:3000/scores e verificar items

---

## ❌ Troubleshooting

### Erro: "relation score_items does not exist"
**Causa:** Banco não está populado
**Solução:**
```bash
# Rodar migrations
docker compose exec api atlas migrate apply --env dev
```

### Erro: "column last_review does not exist"
**Causa:** Schema desatualizado
**Solução:**
```bash
# Verificar schema
docker compose exec db psql -U plenya_user -d plenya_db -c "\d score_items"

# Se faltar coluna, adicionar:
docker compose exec db psql -U plenya_user -d plenya_db -c "ALTER TABLE score_items ADD COLUMN IF NOT EXISTS last_review TIMESTAMP;"
```

### Erro: "duplicate key value violates unique constraint"
**Causa:** SQL já foi executado antes
**Solução:**
- Não é erro crítico, items já estão enriquecidos
- Para reexecutar: deletar e rodar novamente (não recomendado em produção)

### Zero items atualizados (validação retorna 0)
**Causa:** IDs dos items não existem no banco
**Solução:**
```bash
# Verificar se items existem
docker compose exec db psql -U plenya_user -d plenya_db -c "SELECT COUNT(*) FROM score_items WHERE id IN ('341946e7-5833-48bc-b316-71e29954eedd', '348fc460-9959-4648-9d0d-6acafd2f9700');"

# Se retornar 0, items não existem
# Solução: popular banco com seed data primeiro
```

---

## 📁 Arquivos Importantes

| Arquivo | Descrição | Tamanho |
|---------|-----------|---------|
| `batch_final_1_exames_A.sql` | SQL executável (45 UPDATEs) | 640 linhas |
| `scripts/validate_batch_final_1.sql` | Script de validação | 9 queries |
| `batch_final_1_exames_A_results.json` | Resultados detalhados | 548 linhas |
| `BATCH-FINAL-1-EXAMES-A-REPORT.md` | Relatório técnico completo | Documentação |
| `BATCH-FINAL-1-EXECUTIVE-SUMMARY.md` | Sumário executivo | Overview |

---

## 🎯 Resultado Final

Após execução bem-sucedida:

✅ **45 items enriquecidos** com conteúdo MFI completo
✅ **3 items com conteúdo específico** de alta qualidade
✅ **147 artigos científicos** referenciados
✅ **Condutas clínicas específicas** com doses e protocolos
✅ **Sistema pronto** para uso clínico

---

## 📞 Suporte

Se encontrar problemas:

1. Verificar logs do Docker: `docker compose logs db`
2. Verificar se serviços estão rodando: `docker compose ps`
3. Consultar documentação: `BATCH-FINAL-1-EXAMES-A-REPORT.md`

---

**Última atualização:** 2026-01-28
**Versão:** 1.0
**Sistema:** Plenya EMR
