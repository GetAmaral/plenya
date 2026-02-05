# Diagnóstico: Erro na Importação do Grupo "Movimento e atividade física"

**Data:** 2026-02-02
**Análise:** Claude Code

---

## 🔍 Resumo Executivo

Durante a importação do arquivo `escore.csv` para o banco de dados, ocorreu um erro na estrutura hierárquica do grupo "Movimento e atividade física", resultando em:
- **2 itens pais ausentes** (Adolescência, Vida adulta)
- **4 itens filhos órfãos** (sem parent_item_id correto)
- **6 score_levels ausentes** (3 níveis para cada item com pontos faltante)

**Severidade:** Média
**Impacto:** Dados inconsistentes - estrutura organizacional comprometida
**Perda de dados:** Nenhuma (apenas falta estrutura hierárquica)

---

## 📊 Análise Detalhada

### Estrutura Esperada (CSV)

```
Movimento e atividade física
├── Histórico
│   ├── Infância (0 pts) ✅
│   │   ├── Esportes praticados (5 pts) ⚠️ ÓRFÃO
│   │   └── Atividades ao ar livre (0 pts) ⚠️ ÓRFÃO
│   ├── Adolescência (0 pts) ❌ AUSENTE
│   │   ├── Esportes praticados (5 pts) ❌ AUSENTE
│   │   └── Atividades ao ar livre (0 pts) ❌ AUSENTE
│   ├── Vida adulta (0 pts) ❌ AUSENTE
│   │   ├── Esportes praticados (5 pts) ❌ AUSENTE
│   │   └── Atividades ao ar livre (0 pts) ❌ AUSENTE
│   ├── Melhores fases atléticas ✅
│   ├── Piores fases de sedentarismo ✅
│   ├── Modalidades preferidas ✅
│   ├── Modalidades "odiadas" ✅
│   ├── Lesões relacionadas ✅
│   ├── Cirurgias relacionadas ✅
│   ├── Restrições de atividades ✅
│   └── Histórico familiar ✅
└── Atual
    ├── Estratégia macro atual ✅
    ├── Divisão das atividades ✅
    ├── Horários ✅
    ├── Onde e como faz ✅
    ├── Quem treina ✅
    ├── Suplementação ✅
    ├── Provas/desafios planejados ✅
    └── Situação familiar/amigos ✅
```

### Estado Atual do Banco

**Itens presentes no grupo Movimento:**
- Total de itens: 19
- Itens com pontos: 6
- Itens sem pontos (organizadores): 13
- **Itens com parent_item_id:** 0 ❌

**Query de verificação:**
```sql
SELECT COUNT(*) FROM score_items si
JOIN score_subgroups sg ON si.subgroup_id = sg.id
WHERE sg.group_id = '019bf31d-2ef0-753f-a9df-6afd3b4fafdb'
AND si.parent_item_id IS NOT NULL;
-- Resultado: 0 (esperado: 6)
```

---

## 🐛 Causa Raiz

### Lógica do Script de Importação

O script `import_score_csv.py` implementa hierarquia através de:

```python
# Linha 489-507: Quando col3 preenchida
if col3_item:
    item_name = col3_item
    last_item_col3_id = importer.insert_item(...)  # Guarda como "pai"

# Linha 509-522: Quando col3 vazia, col4 preenchida
elif col4_subitem_ou_unit:
    item_name = col4_subitem_ou_unit
    parent_id = last_item_col3_id  # Usa último col3 como pai
```

**Problema:** A variável `last_item_col3_id` é **GLOBAL e STICKY**, então:

1. Linha 77: `col3=[Infância]` → `last_item_col3_id = ID_INFANCIA` ✅
2. Linha 78: `col4=[Esportes...]` → parent = ID_INFANCIA ✅
3. Linha 79: `col4=[Atividades...]` → parent = ID_INFANCIA ✅
4. Linha 80: `col3=[Adolescência]` → deveria criar item E atualizar last_item_col3_id
5. **MAS** Linha 80 tem col5 VAZIO, então pode ter sido pulada pela lógica:
   ```python
   # Linha 470-472
   if not col3_item and not col4_subitem_ou_unit:
       importer.stats['rows_skipped'] += 1
       continue
   ```

**Hipótese:** Linhas 80 e 83 (Adolescência, Vida adulta) foram **PULADAS** porque:
- col3 preenchida: "Adolescência", "Vida adulta"
- col4 vazia: ""
- col5 vazia: "" (SEM PONTOS)

O script pode ter interpretado como "linha sem conteúdo relevante" e pulou.

---

## 📋 Itens Faltantes

### 1. Adolescência (item pai)
- **Subgroup:** Histórico (ID: 019bf31d-2ef0-7fae-8e83-194607f9a612)
- **Name:** Adolescência
- **Points:** 0 (organizador)
- **Order:** ~2 (entre Infância e Vida adulta)
- **Parent:** NULL

### 2. Esportes praticados - Adolescência (filho)
- **Name:** Esportes praticados (frequência e intensidade)
- **Points:** 5
- **Parent:** Adolescência (a ser criado)
- **Níveis (6):**
  - Nivel 0: Nenhum esporte na adolescência
  - Nivel 1: Atividades irregulares por menos de 5 anos
  - Nivel 2: 1-2h de atividade por semana por 5+ anos
  - Nivel 3: 2-3h de atividade por semana por 5+ anos
  - Nivel 4: 3-5h deatividade por semana por 5+ anos
  - Nivel 5: Mais de 5h de atividades por semana por 5+ anos

### 3. Atividades ao ar livre - Adolescência (filho)
- **Name:** Atividades ao ar livre
- **Points:** 0 (qualitativo)
- **Parent:** Adolescência

### 4. Vida adulta (item pai)
- **Name:** Vida adulta
- **Points:** 0 (organizador)
- **Order:** ~3
- **Parent:** NULL

### 5. Esportes praticados - Vida adulta (filho)
- **Name:** Esportes praticados (frequência e intensidade)
- **Points:** 5
- **Parent:** Vida adulta (a ser criado)
- **Níveis (6):**
  - Nivel 0: Nenhum esporte na vida adulta
  - Nivel 1: Atividades irregulares por menos de 5 anos
  - Nivel 2: 1-2h de atividade por semana por 5+ anos
  - Nivel 3: 2-3h de atividade por semana por 10+ anos
  - Nivel 4: 3-5h deatividade por semana por 10+ anos
  - Nivel 5: Mais de 5h de atividades por semana por 10+ anos

### 6. Atividades ao ar livre - Vida adulta (filho)
- **Name:** Atividades ao ar livre
- **Points:** 0 (qualitativo)
- **Parent:** Vida adulta

---

## ⚠️ Problemas de Hierarquia Existentes

### Itens órfãos (sem parent correto)

**Query de verificação:**
```sql
SELECT si.name, si.parent_item_id
FROM score_items si
JOIN score_subgroups sg ON si.subgroup_id = sg.id
WHERE sg.group_id = '019bf31d-2ef0-753f-a9df-6afd3b4fafdb'
AND sg.name = 'Histórico'
AND si.name IN ('Atividades ao ar livre', 'Esportes praticados (frequência e intensidade)');
```

**Resultado:**
- Atividades ao ar livre: parent_item_id = NULL (deveria ser ID de Infância)
- Esportes praticados: parent_item_id = NULL (deveria ser ID de Infância)

**⚠️ ATENÇÃO:** Esses dois itens JÁ EXISTEM mas estão órfãos. A correção deve:
1. Criar novos itens para Adolescência e Vida adulta
2. Criar novos filhos para Adolescência e Vida adulta
3. **NÃO** alterar os itens órfãos existentes (são de Infância, apenas falta o parent_id)

---

## 🎯 Plano de Ação

### Fase 1: Backup (CRÍTICO)
```bash
# Backup completo antes de qualquer mudança
docker compose exec db pg_dump -U plenya_user plenya_db > \
  backup_before_fix_movimento_$(date +%Y%m%d_%H%M%S).sql

# Backup específico das tabelas de score
docker compose exec db pg_dump -U plenya_user plenya_db \
  -t score_groups -t score_subgroups -t score_items -t score_levels > \
  backup_score_tables_$(date +%Y%m%d_%H%M%S).sql
```

### Fase 2: Script de Correção
Criar arquivo `fix_movimento_hierarchy.sql` que:

1. **Corrige hierarquia dos itens existentes de Infância:**
   ```sql
   -- Buscar ID de Infância
   UPDATE score_items si
   SET parent_item_id = (
     SELECT id FROM score_items
     WHERE name = 'Infância' AND subgroup_id = si.subgroup_id
   )
   WHERE si.name IN (
     'Atividades ao ar livre',
     'Esportes praticados (frequência e intensidade)'
   )
   AND si.subgroup_id = '019bf31d-2ef0-7fae-8e83-194607f9a612'
   AND si.parent_item_id IS NULL;
   ```

2. **Insere item "Adolescência":**
   - Usar uuid_generate_v7() para ID
   - Order: 2 (após Infância, ordem 1)
   - Points: 0
   - Subgroup: Histórico

3. **Insere filhos de Adolescência:**
   - Esportes praticados (5 pontos + 6 níveis)
   - Atividades ao ar livre (0 pontos)

4. **Insere item "Vida adulta":**
   - Order: 3
   - Points: 0

5. **Insere filhos de Vida adulta:**
   - Esportes praticados (5 pontos + 6 níveis)
   - Atividades ao ar livre (0 pontos)

### Fase 3: Validação Pós-Correção
```sql
-- Verificar contagem final
SELECT
  (SELECT COUNT(*) FROM score_items si
   JOIN score_subgroups sg ON si.subgroup_id = sg.id
   WHERE sg.group_id = '019bf31d-2ef0-753f-a9df-6afd3b4fafdb') as total_items,
  (SELECT COUNT(*) FROM score_items si
   JOIN score_subgroups sg ON si.subgroup_id = sg.id
   WHERE sg.group_id = '019bf31d-2ef0-753f-a9df-6afd3b4fafdb'
   AND si.parent_item_id IS NOT NULL) as items_com_parent;

-- Esperado: total_items = 25, items_com_parent = 6
```

---

## ✅ Checklist de Execução

- [ ] **BACKUP COMPLETO** criado e verificado
- [ ] Script SQL de correção criado e revisado
- [ ] Executar correção em transação (BEGIN)
- [ ] Validar contagem de registros (25 itens totais esperados)
- [ ] Validar hierarquia (6 itens com parent)
- [ ] Validar níveis (todos itens com pontos têm 6 níveis)
- [ ] COMMIT apenas se validações OK
- [ ] Testar no frontend (formulários de anamnese)

---

## 🔒 Garantias de Segurança

**O que NÃO será feito:**
- ❌ Deletar itens existentes
- ❌ Alterar pontuações (points)
- ❌ Remover score_levels
- ❌ Modificar nomes de itens existentes

**O que SERÁ feito:**
- ✅ Adicionar 6 novos items
- ✅ Adicionar 6 novos score_levels (para os 2 itens com 5 pontos)
- ✅ Corrigir parent_item_id de 2 itens órfãos existentes

---

## 📊 Impacto Estimado

| Tabela | Registros Adicionados | Registros Atualizados |
|--------|----------------------|----------------------|
| score_items | +6 | 2 (parent_item_id) |
| score_levels | +6 (2 itens × 3 níveis cada) | 0 |
| **TOTAL** | **+12** | **2** |

**Tempo estimado:** ~5 minutos
**Downtime necessário:** Não (operação online)
**Rollback disponível:** Sim (via backup SQL)

---

**Última atualização:** 2026-02-02
**Responsável:** Claude Code
**Aprovação necessária:** Desenvolvedor / DBA
