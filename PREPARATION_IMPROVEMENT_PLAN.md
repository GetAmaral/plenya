# Plano de Melhoria: Sistema de Preparation (ScoreItem Enrichment)

## 📊 Análise do Estado Atual

### Estatísticas Gerais
- **Cobertura:** 878/878 ScoreItems (100%) ✅
- **Avg chunks por preparation:** 13.0
- **Range:** 1-30 chunks

### 🔴 Problemas Críticos Identificados

| Problema | Qtd | % | Impacto |
|----------|-----|---|---------|
| **Apenas 1 chunk** | 103 | 11.7% | Contexto científico INSUFICIENTE para enrichment |
| **< 10 chunks** | 376 | 42.8% | Contexto FRACO - enrichment limitado |
| **No limite (20 ou 30)** | 401 | 45.7% | Poderiam ter MAIS chunks disponíveis |

### Distribuição Atual

```
Chunks por Preparation:
  1 chunk:        103 (11.7%) ❌ MUITO FRACO
  2-5 chunks:     197 (22.4%) ❌ FRACO
  6-15 chunks:    177 (20.2%) ⚠️  RAZOÁVEL
  16-25 chunks:   367 (41.8%) ✅ BOM
  26-30 chunks:    34 (3.9%)  ✅ EXCELENTE
```

### Parâmetros Atuais (Problemáticos)

```go
// prepare-all/main.go linha 49
prepService.PrepareChunks(item.ID, 30, 0.3)
//                                 ^^  ^^^
//                          limite  threshold FIXO
```

**Problemas:**
1. **Threshold 0.3 fixo** - muito alto para items genéricos
2. **Limite 30** - muito baixo (poderia ser 40-50)
3. **Sem fallback** - se não encontrar chunks com 0.3, desiste
4. **Sem re-preparation** - ignora items que já têm preparation (mesmo que fraca)

---

## 🎯 Proposta de Melhorias

### Melhoria 1: Threshold Adaptativo com Fallback (CRÍTICO)

**Problema:** 42.8% têm < 10 chunks porque threshold 0.3 é muito alto.

**Solução:** Tentativas progressivas até atingir 15+ chunks

```go
func prepareWithAdaptiveThreshold(scoreItemID uuid.UUID) {
    targetChunks := 20  // Alvo mínimo
    maxLimit := 50      // Limite máximo

    // Tentativa 1: Threshold normal (0.3)
    chunks := FindTopChunks(scoreItemID, maxLimit, 0.3)

    if len(chunks) >= targetChunks {
        return chunks[:min(30, len(chunks))]  // Sucesso
    }

    // Tentativa 2: Threshold permissivo (0.25)
    chunks = FindTopChunks(scoreItemID, maxLimit, 0.25)

    if len(chunks) >= targetChunks {
        return chunks[:min(30, len(chunks))]
    }

    // Tentativa 3: Threshold muito permissivo (0.2)
    chunks = FindTopChunks(scoreItemID, maxLimit, 0.2)

    if len(chunks) >= targetChunks {
        return chunks[:min(35, len(chunks))]  // Mais chunks se threshold baixo
    }

    // Tentativa 4: Fallback extremo (0.15)
    chunks = FindTopChunks(scoreItemID, maxLimit, 0.15)

    return chunks[:min(40, len(chunks))]  // Retorna o que encontrar
}
```

**Resultado esperado:**
- Preparations fracas (< 10): 42.8% → < 5%
- Avg chunks: 13 → 22-25
- Preparations com 1 chunk: 11.7% → 0%

---

### Melhoria 2: Re-Prepare Items Fracos

**Problema:** 376 preparations têm < 10 chunks (insuficiente para enrichment).

**Solução:** Comando para re-preparar apenas os fracos

```bash
# Novo comando: re-prepare-weak/main.go
docker compose exec api go run /app/cmd/re-prepare-weak/main.go

# Estratégia:
# 1. Selecionar preparations com < 15 chunks
# 2. DELETE preparation antiga
# 3. Re-preparar com threshold adaptativo (0.25 → 0.2 → 0.15)
# 4. Validar que tem >= 15 chunks ou abortar
```

**Verificação:**
```sql
SELECT
    COUNT(*) FILTER (WHERE (metadata->>'total_chunks')::int < 10) as antes_weak,
    -- Após re-prepare
    COUNT(*) FILTER (WHERE (metadata->>'total_chunks')::int < 10) as depois_weak
FROM score_item_enrichment_preparation;
```

---

### Melhoria 3: Aumentar Limite de Chunks (FÁCIL)

**Problema:** 45.7% (401 preparations) estão no limite (20 ou 30), significando que poderiam ter MAIS.

**Solução:** Aumentar limite de 30 → 50

```go
// prepare-all/main.go
// Antes:
prepService.PrepareChunks(item.ID, 30, 0.3)

// Depois:
prepService.PrepareChunks(item.ID, 50, 0.3)
```

**Impacto:**
- Preparations com 30 chunks → podem chegar a 40-50
- Mais contexto científico para enrichment
- Custo quase zero (só storage)

---

### Melhoria 4: Priorização Inteligente de Chunks (MÉDIO)

**Problema:** Não sabemos QUAIS chunks são selecionados - pode estar pegando chunks ruins.

**Solução:** Ranking multi-critério

```go
// Fórmula de ranking (já implementada parcialmente)
score = base_similarity × section_weight × recency_factor

// Section weights (já existe):
- results:     1.10 (+10%)
- discussion:  1.05 (+5%)
- methods:     1.02 (+2%)
- introduction: 1.01 (+1%)

// NOVO - Recency factor:
recency_factor = {
    if year >= 2020: 1.05 (+5% para artigos recentes)
    if year >= 2015: 1.0
    else: 0.95
}

// NOVO - Length bonus:
length_bonus = {
    if word_count >= 200: 1.02 (+2% para chunks substanciais)
    else: 1.0
}
```

**Implementação:**
Modificar `FindTopChunksForScoreItem` para incluir:
- Recency boost
- Length quality filter

---

### Melhoria 5: Metadata Enriquecido (FÁCIL)

**Problema:** Metadata atual não tem informações suficientes para QA.

**Solução:** Adicionar campos úteis

```go
// Score_item_enrichment_preparation.metadata
{
    "total_chunks": 25,
    "articles_count": 8,
    "avg_similarity": 0.68,
    "min_similarity": 0.45,          // NOVO
    "max_similarity": 0.92,          // NOVO
    "sections_distribution": {...},
    "total_word_count": 5420,        // NOVO
    "avg_chunk_length": 216,         // NOVO
    "recency_stats": {               // NOVO
        "newest_year": 2024,
        "oldest_year": 2010,
        "avg_year": 2018
    },
    "quality_grade": "excellent"     // NOVO (based on chunks + similarity)
}
```

**Uso:**
- QA: identificar preparations ruins antes de enrichment
- Analytics: tracking de qualidade ao longo do tempo

---

### Melhoria 6: Invalidação Automática (CRÍTICO)

**Problema:** Se ScoreItem é editado, preparation fica DESATUALIZADA (mesma issue dos embeddings).

**Solução:** Hook BeforeUpdate no ScoreItem (similar ao embedding)

```go
// score_item.go - BeforeUpdate
func (si *ScoreItem) BeforeUpdate(tx *gorm.DB) error {
    // ... código existente de embedding invalidation

    // NOVO: Invalidar preparation também
    if needsReembedding {  // Mesma condição
        tx.Exec(`
            UPDATE score_item_enrichment_preparation
            SET status = 'stale'
            WHERE score_item_id = ?
        `, si.ID)

        // Ou deletar diretamente para forçar re-prepare
        tx.Exec(`
            DELETE FROM score_item_enrichment_preparation
            WHERE score_item_id = ?
        `, si.ID)
    }

    return nil
}
```

**Adicionar status 'stale':**
```sql
ALTER TABLE score_item_enrichment_preparation
ALTER COLUMN status TYPE varchar(20);

ALTER TABLE score_item_enrichment_preparation
DROP CONSTRAINT IF EXISTS score_item_enrichment_preparation_status_check;

ALTER TABLE score_item_enrichment_preparation
ADD CONSTRAINT score_item_enrichment_preparation_status_check
CHECK (status IN ('ready','processing','completed','failed','stale'));
```

---

## 🚀 Plano de Implementação

### Fase 1: Quick Wins (1 hora)

#### 1.1. Aumentar limite de chunks (10 min)
```go
// prepare-all/main.go linha 49
prepService.PrepareChunks(item.ID, 50, 0.3)  // 30 → 50
```

#### 1.2. Re-prepare items com < 10 chunks (20 min)
Criar comando `re-prepare-weak/main.go`:
```go
// DELETE preparations com < 10 chunks
// Re-executar prepare-all apenas para esses
```

#### 1.3. Adicionar metadata enriquecido (30 min)
Modificar `PrepareChunks` para calcular min/max similarity, word count total, etc.

---

### Fase 2: Threshold Adaptativo (2 horas)

#### 2.1. Implementar tentativas progressivas
```go
func PrepareChunksAdaptive(scoreItemID uuid.UUID) {
    thresholds := []float64{0.3, 0.25, 0.2, 0.15}
    targetChunks := 20

    for _, threshold := range thresholds {
        chunks := FindTopChunks(scoreItemID, 50, threshold)

        if len(chunks) >= targetChunks {
            return savePreparation(chunks[:min(30, len(chunks))])
        }
    }

    // Salvaguarda: retornar o que encontrou
    return savePreparation(chunks)
}
```

#### 2.2. Logging de fallback
Adicionar log quando usa threshold < 0.3:
```
⚠️  ScoreItem "Albumina": fallback to threshold 0.25 (30 chunks found)
```

---

### Fase 3: Priorização Avançada (3 horas - OPCIONAL)

#### 3.1. Recency boost
```sql
-- Modificar FindTopChunksForScoreItem
recency_factor = CASE
    WHEN EXTRACT(YEAR FROM a.publish_date) >= 2020 THEN 1.05
    WHEN EXTRACT(YEAR FROM a.publish_date) >= 2015 THEN 1.0
    ELSE 0.95
END
```

#### 3.2. Length quality filter
```sql
-- Filtrar chunks muito curtos (< 50 words)
WHERE COALESCE((ae.chunk_metadata->>'word_count')::int, 0) >= 50
```

---

## 📊 Resultados Esperados

### Antes (Atual)
```
Preparations:
  Fracas (< 10 chunks):    376 (42.8%) ❌
  1 chunk apenas:          103 (11.7%) ❌
  Avg chunks:              13.0
  No limite (20/30):       401 (45.7%)
```

### Depois (Projetado)
```
Preparations:
  Fracas (< 10 chunks):    < 50 (< 6%) ✅
  1 chunk apenas:          0 (0%) ✅
  Avg chunks:              24-27
  Ótimas (20-40 chunks):   ~750 (85%+) ✅
```

---

## 🛠️ Arquivos a Modificar

### Quick Wins
1. `apps/api/cmd/prepare-all/main.go` - aumentar limite 30→50
2. `apps/api/cmd/re-prepare-weak/main.go` - NOVO comando
3. `apps/api/internal/services/score_enrichment_preparation_service.go` - metadata enriquecido

### Threshold Adaptativo
4. `apps/api/internal/services/score_enrichment_preparation_service.go` - lógica adaptativa
5. `apps/api/cmd/prepare-all/main.go` - usar método adaptativo

### Invalidação Automática
6. `apps/api/internal/models/score_item.go` - hook BeforeUpdate
7. `apps/api/database/migrations/` - adicionar status 'stale'

---

## 🧪 Verificação

### Após Quick Wins
```sql
-- Deverá mostrar < 50 preparations fracas
SELECT COUNT(*)
FROM score_item_enrichment_preparation
WHERE (metadata->>'total_chunks')::int < 10;
```

### Após Threshold Adaptativo
```sql
-- Avg chunks deverá ser 24-27
SELECT
    ROUND(AVG((metadata->>'total_chunks')::int), 1) as avg_chunks,
    COUNT(*) FILTER (WHERE (metadata->>'total_chunks')::int >= 20) as boas
FROM score_item_enrichment_preparation;
```

---

## 💭 Considerações

### Por que tantas preparations fracas?

1. **ScoreItems muito específicos** (ex: "ACTN3 rs1815739")
   - Poucos artigos sobre tópicos muito nichados
   - Threshold 0.3 não encontra matches suficientes

2. **ScoreItems demográficos** (ex: "Hemoglobina - Mulheres pós-menopausa")
   - Embeddings com demografia podem ter reduzido matches
   - Threshold mais baixo ajudaria

3. **Artigos com poucos chunks**
   - Se artigo só tem 5 chunks total, não dá para pegar 20
   - Solução: aceitar menos chunks para esses casos

### Trade-offs

| Abordagem | Prós | Contras |
|-----------|------|---------|
| **Threshold fixo 0.3** | Alta precision | 42.8% fracas |
| **Threshold adaptativo** | Cobertura melhor | Pode incluir chunks menos relevantes |
| **Limite 50 chunks** | Mais contexto | Mais noise potencial |

**Recomendação:** Threshold adaptativo + limite 50 + filtro de qualidade mínima (similarity ≥ 0.2 absoluto)

---

## 📝 Próximos Passos Recomendados

### Imediato (FAZER AGORA)
1. ✅ Aumentar limite de 30 → 50
2. ✅ Deletar preparations com < 10 chunks
3. ✅ Re-executar prepare-all com threshold adaptativo

### Curto Prazo (1 semana)
4. Implementar threshold adaptativo completo
5. Adicionar invalidação automática via hook
6. Criar analytics dashboard de qualidade de preparation

### Médio Prazo (1 mês)
7. Recency boost para artigos recentes
8. Length quality filter (chunks < 50 words)
9. A/B testing de strategies de seleção

---

## 🎯 Objetivo Final

```
✅ 100% preparations com >= 15 chunks
✅ Avg chunks: 25-30
✅ 0% preparations com 1 chunk
✅ Auto-invalidação quando ScoreItem muda
✅ Metadata rico para QA e analytics
```

**Enrichment científico de alta qualidade para TODOS os ScoreItems!** 🚀
