# Scripts de Enrichment Científico

Pipeline automatizado de 3 etapas para preparar ScoreItems para enrichment com Claude.

---

## 🚀 Uso Rápido

### Pipeline Completo (Recomendado)
```bash
cd /home/user/plenya/scripts/enrichment
chmod +x *.sh
./RUN-ALL.sh
```

Executa as 3 etapas automaticamente:
1. Regenera embeddings (se necessário)
2. Cria auto-links
3. Prepara prompts

**Duração:** ~30-40 minutos
**Resultado:** Sistema 100% pronto para enrichment

---

## 📋 Scripts Individuais

### 1️⃣ Regenerar Embeddings
```bash
./1-regenerate-embeddings.sh
```

**Quando usar:**
- Após editar ScoreItems (Name, ClinicalRelevance, etc)
- Quando embeddings estão stale
- Para atualizar formato de embeddings

**Duração:** ~15-20 min
**Custo:** ~$0.50

---

### 2️⃣ Auto-Link Articles ↔ ScoreItems
```bash
./2-auto-link.sh
```

**Quando usar:**
- Após adicionar novos Articles
- Após regenerar embeddings
- Para atualizar links

**Duração:** ~5-10 min
**Resultado:** ~11,000 links criados

---

### 3️⃣ Prepare com Prompts
```bash
./3-prepare-with-prompts.sh
```

**Quando usar:**
- Após criar auto-links
- Para gerar prompts novos
- Após mudanças no formato de prompts

**Duração:** ~10-15 min
**Resultado:** 878 preparations com 4 prompts cada

---

## 📊 Comandos de Verificação

### Ver estado geral do sistema
```bash
docker compose exec db psql -U plenya_user -d plenya_db << 'EOF'
SELECT * FROM preparation_health_summary;
SELECT * FROM auto_link_run_summary ORDER BY started_at DESC LIMIT 5;
SELECT * FROM embedding_health_stats;
EOF
```

### Ver preparations problemáticas
```bash
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT * FROM preparation_needs_attention LIMIT 20"
```

### Ver exemplo de preparation completa
```bash
docker compose exec db psql -U plenya_user -d plenya_db << 'EOF'
SELECT
    si.name,
    (prep.metadata->>'quality_grade') as grade,
    (prep.metadata->>'total_chunks') as chunks,
    LENGTH(prep.prompt_clinical_relevance) as prompt_cr_size
FROM score_item_enrichment_preparation prep
JOIN score_items si ON si.id = prep.score_item_id
LIMIT 5;
EOF
```

---

## 🔧 Comandos Avançados

### Re-preparar apenas preparations fracas
```bash
docker compose exec api go run /app/cmd/re-prepare-weak/main.go
```

### Regenerar embeddings de TODOS (forçar)
```bash
docker compose exec db psql -U plenya_user -d plenya_db << 'EOF'
-- Marcar todos como stale
UPDATE score_item_embeddings SET is_stale = true;

-- Enfileirar para regeneração
INSERT INTO embedding_queue (entity_type, entity_id, status)
SELECT 'score_item', id, 'pending'
FROM score_items
ON CONFLICT (entity_type, entity_id)
DO UPDATE SET status = 'pending', retry_count = 0;
EOF

# Processar fila
docker compose exec api go run /app/cmd/regenerate-embeddings/main.go
```

### Verificar progresso de batch processing
```bash
docker compose exec db psql -U plenya_user -d plenya_db << 'EOF'
SELECT * FROM auto_link_run_summary ORDER BY started_at DESC LIMIT 3;
SELECT * FROM auto_link_batch_performance ORDER BY run_id DESC, batch_number LIMIT 10;
EOF
```

---

## 📁 Estrutura de Dados

### Fluxo de Dados
```
ScoreItems (878)
    ↓ (embeddings)
ScoreItemEmbeddings (878) ← is_stale flag
    ↓ (similarity search)
Articles (996) → ArticleEmbeddings (13,484 chunks)
    ↓ (auto-link)
ArticleScoreItems (11,188 links)
    ↓ (prepare)
ScoreItemEnrichmentPreparation (878)
    ├─ selected_chunks (30 chunks)
    ├─ metadata (quality_grade, avg_similarity, etc)
    ├─ prompt_clinical_relevance (~32KB)
    ├─ prompt_patient_explanation (~32KB)
    ├─ prompt_conduct (~32KB)
    └─ prompt_max_points (~400 chars)
```

---

## 🎯 Métricas de Qualidade

### Expectations

| Métrica | Target | Atual |
|---------|--------|-------|
| **Embeddings válidos** | 100% | ✅ 100% |
| **Auto-links** | ~10,000 | ✅ 11,188 |
| **Cobertura de ScoreItems** | > 95% | ✅ 99.8% |
| **Preparations com prompts** | 100% | ✅ 100% |
| **Quality excellent/good** | > 70% | ✅ 83.7% |
| **Avg chunks** | > 20 | ✅ 30 |

---

## 🆘 Troubleshooting

### Erro: "embedding_queue tem muitos pending"
```bash
# Ver fila
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT status, COUNT(*) FROM embedding_queue GROUP BY status"

# Processar manualmente
docker compose exec api go run /app/cmd/regenerate-embeddings/main.go
```

### Erro: "No chunks found"
```bash
# Verificar se Articles têm embeddings
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT embedding_status, COUNT(*) FROM articles GROUP BY embedding_status"

# Se muitos sem embeddings, processar Articles primeiro
```

### Preparations sem prompts
```bash
# Verificar
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT COUNT(*) FROM score_item_enrichment_preparation
   WHERE prompt_clinical_relevance IS NULL"

# Re-preparar
cd /home/user/plenya/scripts/enrichment
./3-prepare-with-prompts.sh
# Escolher opção 2 (re-preparar todos)
```

---

## 📚 Documentação Adicional

- `/home/user/plenya/.claude/workflows/enrichment-automation.md` - Workflow completo
- `/home/user/plenya/IMPLEMENTATION_SUMMARY.md` - Melhorias do sistema RAG
- `/home/user/plenya/PREPARATION_SUCCESS_REPORT.md` - Resultados de preparation
- `/home/user/plenya/PROMPT_REAL_EXEMPLO.md` - Exemplos de prompts reais

---

**Última atualização:** 2026-02-17
**Versão:** 2.0 (Com prompts automáticos)
