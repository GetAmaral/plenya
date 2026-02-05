# Plano de Ação: Correção do Grupo "Movimento e atividade física"

**Data:** 2026-02-02
**Responsável:** Claude Code
**Status:** ✅ PRONTO PARA EXECUÇÃO

---

## 📋 Resumo

**Problema identificado:**
- Itens "Adolescência" e "Vida adulta" ausentes no grupo Movimento
- Hierarquia parent_item_id incorreta em 2 itens de Infância
- Total de 6 itens e 12 níveis faltando no banco

**Solução:**
- Script SQL automatizado com validações
- Backup completo criado
- Rollback disponível

---

## ✅ Pré-requisitos (CONCLUÍDOS)

- [x] **Diagnóstico completo** → Ver `DIAGNOSTICO_MOVIMENTO.md`
- [x] **Backup full criado** → `backups/backup_full_before_fix_movimento_20260202_123021.sql` (11M)
- [x] **Backup score tables** → `backups/backup_score_tables_20260202_134345.sql` (4.1M)
- [x] **Script de correção** → `fix_movimento_hierarchy.sql`

---

## 🚀 Instruções de Execução

### Opção 1: Execução Via Docker (RECOMENDADO)

```bash
# 1. Verificar que backups existem
ls -lh /home/user/plenya/backups/*.sql

# 2. Executar script de correção (em transação)
docker compose exec -T db psql -U plenya_user -d plenya_db < \
  /home/user/plenya/fix_movimento_hierarchy.sql

# 3. Revisar output das validações
# Se tudo estiver OK, executar COMMIT:
docker compose exec db psql -U plenya_user -d plenya_db -c "COMMIT;"

# Se algo der errado, executar ROLLBACK:
docker compose exec db psql -U plenya_user -d plenya_db -c "ROLLBACK;"
```

### Opção 2: Execução Interativa (Mais Segura)

```bash
# 1. Entrar no container do banco
docker compose exec db psql -U plenya_user -d plenya_db

# 2. No psql, executar:
\i /fix_movimento_hierarchy.sql

# 3. Revisar output das validações

# 4. Se tudo OK:
COMMIT;

# Ou reverter:
ROLLBACK;

# 5. Sair do psql
\q
```

---

## 📊 Output Esperado

O script deve exibir mensagens como:

```
NOTICE:  ✅ Validações pré-execução OK
NOTICE:  ✅ Parte 1: 2 items órfãos de Infância corrigidos
NOTICE:  ✅ Item "Adolescência" criado: [UUID]
NOTICE:  ✅ Filho "Esportes praticados" (Adolescência) criado: [UUID]
NOTICE:  ✅ 6 níveis criados para Esportes praticados (Adolescência)
NOTICE:  ✅ Filho "Atividades ao ar livre" (Adolescência) criado: [UUID]
NOTICE:  ✅ Parte 2: Item "Adolescência" e filhos criados com sucesso
NOTICE:  ✅ Item "Vida adulta" criado: [UUID]
NOTICE:  ✅ Filho "Esportes praticados" (Vida adulta) criado: [UUID]
NOTICE:  ✅ 6 níveis criados para Esportes praticados (Vida adulta)
NOTICE:  ✅ Filho "Atividades ao ar livre" (Vida adulta) criado: [UUID]
NOTICE:  ✅ Parte 3: Item "Vida adulta" e filhos criados com sucesso
NOTICE:
NOTICE:  === VALIDAÇÕES PÓS-EXECUÇÃO ===
NOTICE:  Total de items no grupo Movimento: 25 (esperado: 25)
NOTICE:  Items com parent_item_id: 6 (esperado: 6)
NOTICE:  Item "Adolescência" existe: 1 (esperado: 1)
NOTICE:  Item "Vida adulta" existe: 1 (esperado: 1)
NOTICE:  Níveis de Esportes praticados: 18 (esperado: 18 = 3×6)
NOTICE:
NOTICE:  ✅ TODAS AS VALIDAÇÕES PASSARAM!
NOTICE:

              status              | total_items_movimento | items_com_parent
----------------------------------+-----------------------+------------------
 Correção concluída com sucesso! |                    25 |                6
(1 row)
```

---

## ⚠️ O Que Fazer Se Algo Der Errado

### Erro Durante Execução

```bash
# 1. ROLLBACK imediatamente
docker compose exec db psql -U plenya_user -d plenya_db -c "ROLLBACK;"

# 2. Verificar mensagem de erro no output

# 3. Restaurar backup se necessário (apenas se COMMIT foi executado)
docker compose exec -T db psql -U plenya_user -d plenya_db < \
  backups/backup_score_tables_20260202_134345.sql
```

### Restaurar Backup Completo (Último Recurso)

```bash
# ⚠️ ATENÇÃO: Isso irá sobrescrever TODOS os dados do banco!

# 1. Parar aplicação
docker compose stop web api

# 2. Dropar banco e recriar
docker compose exec db psql -U plenya_user -c "DROP DATABASE plenya_db;"
docker compose exec db psql -U plenya_user -c "CREATE DATABASE plenya_db;"

# 3. Restaurar backup
docker compose exec -T db psql -U plenya_user -d plenya_db < \
  backups/backup_full_before_fix_movimento_20260202_123021.sql

# 4. Reiniciar aplicação
docker compose start web api
```

---

## 🔍 Validação Pós-Execução

### Queries de Verificação Manual

```sql
-- 1. Verificar total de items no grupo Movimento
SELECT COUNT(*) as total
FROM score_items si
JOIN score_subgroups sg ON si.subgroup_id = sg.id
JOIN score_groups g ON sg.group_id = g.id
WHERE g.name = 'Movimento e atividade física'
AND si.deleted_at IS NULL;
-- Esperado: 25

-- 2. Verificar hierarquia
SELECT si.name, si.parent_item_id IS NOT NULL as tem_parent
FROM score_items si
JOIN score_subgroups sg ON si.subgroup_id = sg.id
JOIN score_groups g ON sg.group_id = g.id
WHERE g.name = 'Movimento e atividade física'
AND sg.name = 'Histórico'
AND si.deleted_at IS NULL
ORDER BY si."order";
-- Deve mostrar 6 items com tem_parent = true

-- 3. Verificar estrutura hierárquica completa
SELECT
  COALESCE(p.name, '(raiz)') as pai,
  si.name as filho,
  si.points
FROM score_items si
LEFT JOIN score_items p ON si.parent_item_id = p.id
JOIN score_subgroups sg ON si.subgroup_id = sg.id
JOIN score_groups g ON sg.group_id = g.id
WHERE g.name = 'Movimento e atividade física'
AND sg.name = 'Histórico'
AND si.deleted_at IS NULL
ORDER BY COALESCE(p."order", si."order"), si."order";

-- 4. Verificar níveis dos items com pontos
SELECT
  si.name,
  si.points,
  COUNT(sl.id) as qtd_niveis
FROM score_items si
LEFT JOIN score_levels sl ON sl.item_id = si.id
JOIN score_subgroups sg ON si.subgroup_id = sg.id
JOIN score_groups g ON sg.group_id = g.id
WHERE g.name = 'Movimento e atividade física'
AND si.points > 0
AND si.deleted_at IS NULL
GROUP BY si.id, si.name, si.points
ORDER BY si.name;
-- Todos devem ter 6 níveis
```

---

## 📊 Impacto da Correção

| Métrica | Antes | Depois | Delta |
|---------|-------|--------|-------|
| Items totais (Movimento) | 19 | 25 | +6 |
| Items com parent | 0 | 6 | +6 |
| Score levels | 42 | 54 | +12 |
| Items organizadores | 13 | 15 | +2 |
| Items com pontos | 6 | 8 | +2 |

---

## 🧪 Testes de Validação Frontend

Após executar a correção, testar no frontend:

1. **Acessar formulário de Anamnese**
   ```
   http://localhost:3000/anamnesis
   ```

2. **Verificar grupo "Movimento e atividade física"**
   - Subgrupo "Histórico" deve ter:
     - Infância (com 2 filhos)
     - **Adolescência (com 2 filhos)** ← NOVO
     - **Vida adulta (com 2 filhos)** ← NOVO
     - Outros items...

3. **Testar preenchimento**
   - Clicar em "Adolescência"
   - Verificar que "Esportes praticados" tem 6 opções
   - Selecionar um nível e salvar
   - Verificar que dados são persistidos

4. **Verificar hierarquia no UI**
   - Items filhos devem aparecer identados
   - Drag-and-drop deve respeitar hierarquia

---

## 📝 Checklist Final

### Pré-Execução
- [x] Diagnóstico completo realizado
- [x] Backup full criado e verificado (11M)
- [x] Backup score tables criado (4.1M)
- [x] Script de correção criado e revisado
- [ ] Aplicação em modo manutenção (opcional)

### Execução
- [ ] Script SQL executado via psql
- [ ] Todas as validações passaram (✅ verde no output)
- [ ] Query manual de verificação executada
- [ ] COMMIT realizado

### Pós-Execução
- [ ] Total de 25 items no grupo Movimento
- [ ] 6 items com parent_item_id
- [ ] Adolescência e Vida adulta existem
- [ ] 18 níveis para Esportes praticados (3×6)
- [ ] Frontend testado - estrutura visível
- [ ] Frontend testado - preenchimento funciona
- [ ] Frontend testado - hierarquia correta
- [ ] Aplicação voltou ao normal (se estava em manutenção)

---

## 📞 Suporte

Se houver problemas durante a execução:

1. **ROLLBACK imediatamente**
2. Verificar mensagem de erro
3. Consultar seção "O Que Fazer Se Algo Der Errado"
4. Restaurar backup se necessário

---

## 🔐 Garantias de Segurança

**O script NÃO fará:**
- ❌ Deletar items existentes
- ❌ Alterar pontuações (points)
- ❌ Remover score_levels
- ❌ Modificar items de outros grupos

**O script FARÁ:**
- ✅ Adicionar 6 novos items
- ✅ Adicionar 12 novos score_levels
- ✅ Corrigir parent_item_id de 2 items órfãos

**Transação:**
- ✅ Tudo em uma transação única
- ✅ Rollback automático se validações falharem
- ✅ COMMIT manual necessário

---

**Última atualização:** 2026-02-02 13:45 UTC
**Status:** ✅ PRONTO PARA EXECUÇÃO
**Aprovado por:** Aguardando aprovação do desenvolvedor
