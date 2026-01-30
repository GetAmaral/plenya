# PLANO DE LIMPEZA E REORGANIZAÇÃO DO BANCO DE DADOS

**Data:** 2026-01-27
**Status:** PRONTO PARA EXECUÇÃO (aguardando aprovação)

---

## RESUMO EXECUTIVO

Foram identificados **2 problemas críticos** na estrutura de dados do sistema Plenya:

1. **DUPLICAÇÃO MASSIVA:** 1.728 items duplicados (70% do banco!)
2. **GENÉTICA DESORGANIZADA:** 61 exames genéticos como grupos em vez de items

---

## PROBLEMA 1: DUPLICAÇÃO DE ITEMS

### Estatísticas

- **Total de items:** 2.478
- **Nomes únicos:** 750
- **Duplicatas estimadas:** 1.728 (70%!)
- **Levels a deletar:** 6.112

### Top Duplicatas

| Nome | Total Cópias | Com Conteúdo | Vazias | Ação |
|------|--------------|--------------|--------|------|
| "20" | 162 | 30 | 132 | Manter 1, deletar 161 |
| "Outros sintomas" | 15 | 0 | 15 | Manter 1, deletar 14 |
| "Adolescência" | 9 | 9 | 0 | Manter 1, deletar 8 |
| "Atividades ao ar livre" | 9 | 9 | 0 | Manter 1, deletar 8 |
| "Bruxismo" | 6 | 6 | 0 | Manter 1, deletar 5 |
| "% Gordura corporal - homem" | 3 | 1 | 2 | Manter 1, deletar 2 |

### Critério de Decisão

**Para cada conjunto de duplicatas:**
1. Se houver item(s) COM conteúdo (clinical_relevance > 100 chars) → **manter o primeiro com conteúdo**
2. Se TODOS vazios → **manter o mais antigo** (por created_at)
3. **Deletar todos os outros**

### Execução

**Arquivo:** `/home/user/plenya/scripts/cleanup_duplicates.sql`

**Status atual:** Comentado (seguro)
**Para executar:**
1. Descomentar linhas 48-53 (DELETE FROM score_levels e score_items)
2. Trocar linha 63 `ROLLBACK` por `COMMIT`
3. Executar: `cat scripts/cleanup_duplicates.sql | docker compose exec -T db psql -U plenya_user -d plenya_db`

**Resultado esperado:**
- Items: 2.478 → 750 (redução de 70%)
- Levels: 8.XXX → 2.XXX (aproximado)
- Banco limpo e consistente

---

## PROBLEMA 2: GENÉTICA DESORGANIZADA

### Situação Atual

- **Grupo "Genética" existe** mas está vazio (0 subgrupos, 0 items)
- **61 exames genéticos** estão como GRUPOS em vez de items:
  - MTHFR C677T rs1801133
  - APOE (Alzheimer e Lipídios)
  - FTO rs9939609 (Obesidade)
  - MC4R, TCF7L2, PPARG, VDR, ACE, CYP1A2, etc.

### Estrutura Incorreta

```
❌ ATUAL:
score_groups
  ├─ MTHFR C677T (GRUPO)
  │   └─ rs1801133 (subgrupo?)
  ├─ APOE (GRUPO)
  │   └─ variantes (subgrupos?)
  └─ FTO (GRUPO)
```

### Estrutura Correta

```
✅ DEVE SER:
score_groups
  └─ Genética (GRUPO)
      ├─ Metabolismo (SUBGRUPO)
      │   ├─ MTHFR C677T (ITEM)
      │   │   ├─ C/C (LEVEL)
      │   │   ├─ C/T (LEVEL)
      │   │   └─ T/T (LEVEL)
      │   ├─ FTO rs9939609 (ITEM)
      │   └─ MC4R (ITEM)
      ├─ Cardiovascular (SUBGRUPO)
      │   ├─ APOE (ITEM)
      │   └─ ACE I/D (ITEM)
      └─ Detoxificação (SUBGRUPO)
          ├─ CYP1A2 (ITEM)
          └─ GSTM1 (ITEM)
```

### Subgrupos a Criar

1. **Metabolismo** - Genes de nutrientes/energia/hormônios (MTHFR, FTO, MC4R, PPARG, TCF7L2, VDR)
2. **Cardiovascular** - Genes cardiovasculares (APOE, ACE, AGT, AGTR1, NOS3, LDLR)
3. **Neurodegeneração** - Genes de demências (APOE, PSEN1, PSEN2, APP, LRRK2)
4. **Detoxificação** - Genes fase I/II (CYP1A1, CYP1A2, CYP2A6, GSTM1, GSTT1, NAT2)
5. **Imunidade** - Genes imunes (HLA-DQ2/DQ8, IL1B, IL6, TNF, CRP)
6. **Performance** - Genes de performance física (ACTN3, ACE I/D, COL5A1)
7. **Outros** - Genes não categorizados

### Complexidade

⚠️ **ATENÇÃO:** Esta reorganização é MAIS COMPLEXA que a limpeza de duplicatas porque:

1. Requer **migração de dados** (grupos → items)
2. Requer **categorização manual** de 61 genes
3. Pode haver **items/levels dependentes** desses grupos genéticos incorretos
4. Precisa de **script Python** para fazer a migração de forma segura

**Recomendação:** Fazer DEPOIS da limpeza de duplicatas.

---

## ORDEM DE EXECUÇÃO RECOMENDADA

### PASSO 1: Backup (OBRIGATÓRIO)

```bash
docker compose exec db pg_dump -U plenya_user plenya_db > backup_before_cleanup_$(date +%Y%m%d_%H%M%S).sql
```

### PASSO 2: Limpeza de Duplicatas (MAIS SIMPLES)

```bash
# 1. Revisar análise
cat scripts/cleanup_duplicates.sql | docker compose exec -T db psql -U plenya_user -d plenya_db

# 2. Se OK, editar script:
#    - Descomentar linhas 48-53 (DELETEs)
#    - Trocar linha 63: ROLLBACK → COMMIT

# 3. Executar de verdade
cat scripts/cleanup_duplicates.sql | docker compose exec -T db psql -U plenya_user -d plenya_db

# 4. Verificar resultado
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
  SELECT COUNT(*) as total, COUNT(DISTINCT name) as unicos FROM score_items;
"
```

**Esperado:** total = unicos = 750

### PASSO 3: Reorganização Genética (MAIS COMPLEXO - OPCIONAL POR AGORA)

```bash
# Esta etapa requer script Python customizado para:
# 1. Criar subgrupos
# 2. Migrar grupos → items
# 3. Migrar subgrupos → levels
# 4. Atualizar referências

# Recomendo fazer manualmente OU em sessão dedicada futura
```

---

## ARQUIVOS GERADOS

1. **`scripts/analyze_duplicates.sql`** - Análise detalhada de duplicatas ✅
2. **`scripts/cleanup_duplicates.sql`** - Script de limpeza (pronto, comentado) ✅
3. **`scripts/reorganize_genetics.sql`** - Script de reorganização genética (parcial) ⚠️
4. **`scripts/fix_score_structure.py`** - Script Python completo (precisa Docker) ⚠️
5. **`CLEANUP_PLAN.md`** - Este documento ✅

---

## RISCOS E MITIGAÇÕES

| Risco | Probabilidade | Impacto | Mitigação |
|-------|---------------|---------|-----------|
| Deletar item errado | Baixa | Alto | Backup + script validado + dry-run |
| Perder levels órfãos | Baixa | Médio | DELETE CASCADE automático |
| Quebrar frontend | Baixa | Médio | Frontend usa IDs, não nomes |
| Backup falhar | Muito Baixa | Crítico | Testar restore antes |

---

## APROVAÇÃO NECESSÁRIA

**Preciso da sua aprovação explícita para:**

- [ ] **PASSO 1:** Fazer backup do banco
- [ ] **PASSO 2:** Executar limpeza de duplicatas (1.728 items + 6.112 levels)
- [ ] **PASSO 3 (OPCIONAL):** Reorganizar genética (ou deixar para depois?)

**Responda com:**
- "APROVADO - executar passo 1 e 2" OU
- "APROVADO - executar tudo (1, 2 e 3)" OU
- "REVISAR - preciso ver [detalhes específicos]"

---

## DEPOIS DA LIMPEZA

Com o banco limpo, os agentes de enrichment terão **muito mais eficiência**:

- Menos duplicatas para processar
- Queries mais rápidas
- Dados mais consistentes
- Estrutura mais lógica (se fizer reorganização genética)

**Progresso esperado pós-limpeza:**
- Items totais: 750 (vs 2.478 atual)
- Items enriquecidos: 304 (40,5% vs 12,3% atual!)
- **Muito mais próximo da conclusão! 🎯**

---

**Aguardando sua aprovação para prosseguir.**
