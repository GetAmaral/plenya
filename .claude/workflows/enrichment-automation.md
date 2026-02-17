# Enrichment Científico Automatizado

Sistema completo de 3 etapas para preparar ScoreItems com prompts prontos para Claude.

**NOVO:** Scripts bash organizados em `scripts/enrichment/` para execução fácil.

---

## 🚀 Uso Rápido (Scripts Prontos)

### Pipeline Completo
```bash
cd ~/plenya/scripts/enrichment
./RUN-ALL.sh
```

Executa automaticamente:
1. Regenera embeddings (se necessário)
2. Cria auto-links (Articles ↔ ScoreItems)
3. Prepara prompts (4 por ScoreItem)

**Duração:** 30-40 minutos
**Resultado:** 878 ScoreItems prontos com 4 prompts cada

---

## 📋 Scripts Individuais

### 1️⃣ Regenerar Embeddings
```bash
./1-regenerate-embeddings.sh
```
- Processa embedding_queue (items stale)
- Rate limiting automático
- Retry em falhas

### 2️⃣ Auto-Link
```bash
./2-auto-link.sh
```
- Threshold adaptativo (0.2-0.8)
- Batch processing com rollback
- Link de orphans automático

### 3️⃣ Prepare com Prompts
```bash
./3-prepare-with-prompts.sh
```
- Threshold adaptativo (0.35→0.15)
- Gera 4 prompts prontos
- Quality grading automático

---

## 🎯 Comandos Go Diretos (Avançado)

### 1. Auto-link Tudo

**Quando o usuário pedir:** "auto-link tudo" ou "linkar artigos"

**Executar:**
```bash
docker compose exec api go run /app/cmd/auto-link-all/main.go
```

**O que faz:**
- Busca 30 chunks por score item via RAG
- Threshold adaptativo (0.65-0.75 baseado em especificidade)
- Requer ≥2 chunks OU similaridade ≥0.85 para linkar
- Usa `fullName` (Group - Subgroup - Parent - Name) para contexto rico
- Cria links em `article_score_items` table

**Duração:** ~5-10 minutos para 878 score items

**Output esperado:**
```
✅ Created 1247 links
```

---

### 2. Prepare Todos os Score Items (COM PROMPTS)

**Quando o usuário pedir:** "prepare todos os scoreItems" ou "preparar prompts"

**Executar (recomendado):**
```bash
cd ~/plenya/scripts/enrichment
./3-prepare-with-prompts.sh
```

**OU direto:**
```bash
docker compose exec api go run /app/cmd/prepare-all/main.go
```

**O que faz (NOVO):**
- Processa TODOS os 878 ScoreItems
- Threshold adaptativo: 0.35 → 0.30 → 0.25 → 0.20 → 0.15
- Busca 30 chunks por item (vs 20 anterior)
- **GERA 4 PROMPTS PRONTOS** salvos no banco:
  * `prompt_clinical_relevance` (~32KB)
  * `prompt_patient_explanation` (~32KB)
  * `prompt_conduct` (~32KB)
  * `prompt_max_points` (~400 chars)
- Inclui **fullName** (Group - Subgroup - Name) em todos prompts
- Quality grading automático (excellent/good/fair/poor)
- Metadata enriquecido (min/max similarity, recency, sections)

**Duração:** ~10-15 minutos para 878 items

**Output esperado:**
```
✅ Prepared: 878 | Skipped: 0
📊 Quality: 735/878 (83.7%) excellent/good
```

**Resultado:** Cada ScoreItem tem 4 prompts COMPLETOS prontos para enviar ao Claude API

---

### 3. Enrich Tudo (NOVO - Com Prompts Prontos)

**Quando o usuário pedir:** "enrich tudo" ou "enriquecer score items"

**NOVO:** Prompts já estão prontos em `score_item_enrichment_preparation`!

**Buscar preparations:**
```sql
SELECT
    si.id,
    si.name,
    prep.prompt_clinical_relevance,
    prep.prompt_patient_explanation,
    prep.prompt_conduct,
    prep.prompt_max_points,
    prep.metadata
FROM score_item_enrichment_preparation prep
JOIN score_items si ON si.id = prep.score_item_id
WHERE prep.status = 'ready'
ORDER BY (prep.metadata->>'quality_grade') DESC
LIMIT 100;
```

**Para CADA ScoreItem:**

1. **Buscar prompts prontos** (já incluem fullName + 30 chunks + instruções)

2. **Enviar 4 prompts ao Claude API:**

**Exemplo de código Python:**
```python
# Buscar preparation
prep = db.query(ScoreItemEnrichmentPreparation).filter_by(status='ready').first()

# Enviar 4 prompts ao Claude API
import anthropic
client = anthropic.Anthropic(api_key=os.environ["ANTHROPIC_API_KEY"])

# Prompt 1: Clinical Relevance
response_cr = client.messages.create(
    model="claude-sonnet-4",
    max_tokens=2000,
    messages=[{"role": "user", "content": prep.prompt_clinical_relevance}]
)
clinical_relevance = response_cr.content[0].text

# Prompt 2: Patient Explanation
response_pe = client.messages.create(
    model="claude-sonnet-4",
    max_tokens=1200,
    messages=[{"role": "user", "content": prep.prompt_patient_explanation}]
)
patient_explanation = response_pe.content[0].text

# Prompt 3: Conduct
response_cond = client.messages.create(
    model="claude-sonnet-4",
    max_tokens=2000,
    messages=[{"role": "user", "content": prep.prompt_conduct}]
)
conduct = response_cond.content[0].text

# Prompt 4: Max Points
response_pts = client.messages.create(
    model="claude-sonnet-4",
    max_tokens=100,
    messages=[{"role": "user", "content": prep.prompt_max_points}]
)
points = int(response_pts.content[0].text.split()[0])

# Atualizar ScoreItem
score_item = db.query(ScoreItem).get(prep.score_item_id)
score_item.clinical_relevance = clinical_relevance
score_item.patient_explanation = patient_explanation
score_item.conduct = conduct
score_item.points = points
db.commit()

print(f"✅ {score_item.name} enriched (cr: {len(clinical_relevance)}, pe: {len(patient_explanation)}, pts: {points})")
```

**Estrutura dos prompts (JÁ PRONTOS):**

Cada prompt inclui:
- ✅ FullName: "Leptina - Mulheres (Exames - Laboratoriais)"
- ✅ Unit, gender, age range
- ✅ 30 chunks científicos COMPLETOS (~1000 chars cada)
- ✅ Instruções específicas de formato e extensão
- ✅ Tamanho: ~32,000 chars por prompt

**Especificações dos campos:**

   **a) clinical_relevance (1200-1800 chars):**
   - Definição fisiológica precisa
   - Valores de referência com NÚMEROS dos chunks
   - Fisiopatologia resumida
   - Dados epidemiológicos (prevalência, RR, OR, sensibilidade/especificidade)
   - Estratificação de risco por valores

   **b) patient_explanation (600-900 chars):**
   - Linguagem SIMPLES (8º ano escolar)
   - O QUE é, POR QUE importa, O QUE significam valores alterados
   - Tom empático e empoderador

   **c) conduct (1000-1500 chars, Markdown):**
   ```markdown
   ## Interpretação de Valores
   ## Investigação Complementar
   ## Protocolo de Tratamento
   ## Critérios de Referência/Encaminhamento
   ```

   **d) max_points (0-50):**
   - Impacto prognóstico (0-20)
   - Modificabilidade (0-15)
   - Prevalência (0-10)
   - Urgência clínica (0-5)

**Duração:** ~1-2 horas para 878 items (4 chamadas API × 878)

---

## 📊 Workflow Completo

```
Usuário: "auto-link tudo"
  → Claude executa: cmd/auto-link-all/main.go
  → ✅ 1247 links criados

Usuário: "prepare todos os scoreItems"
  → Claude executa: cmd/prepare-all/main.go
  → ✅ 142 preparations criadas

Usuário: "enrich tudo"
  → Claude executa: cmd/list-pending-enrichments/main.go
  → Claude lê JSON e processa cada item:
      - Analisa chunks científicos
      - Gera 4 campos
      - PUT /api/v1/score-items/{id}
  → ✅ 142 score items enriquecidos
```

---

## 🔍 Estrutura do pending_enrichments.json

```json
[
  {
    "score_item_id": "019bf31d-2ef0-78aa-a54e-67723b8c035d",
    "score_item_name": "Zinco",
    "score_item_full_name": "Zinco (Exames - Laboratoriais)",
    "unit": "µg/dL",
    "chunks": {
      "items": [
        {
          "article_title": "Base Metabólica das Doenças Crônicas",
          "article_year": 2024,
          "chunk_text": "Zinco é cofator de >300 enzimas...",
          "section": "results",
          "similarity": 0.78
        }
      ]
    },
    "metadata": {
      "total_chunks": 10,
      "articles_count": 4,
      "avg_similarity": 0.675
    }
  }
]
```

---

## ⚙️ Configurações

### Modificar tier de processamento:

Editar `cmd/prepare-all/main.go`:
```go
// Linha 34 - escolher tier
db.Where("clinical_relevance IS NULL OR LENGTH(clinical_relevance) < 500") // rewrite
// ou
db.Where("LENGTH(clinical_relevance) >= 500 AND LENGTH(clinical_relevance) < 1500") // enrich
// ou
db.Where("1=1") // todos
```

### Modificar thresholds:

Editar `cmd/auto-link-all/main.go` função `determineThreshold()`:
- Específicos: 0.75
- Genéricos: 0.65
- Padrão: 0.70

---

## 📈 Métricas de Qualidade

Após enrichment, validar:

```sql
SELECT
  AVG(LENGTH(clinical_relevance)) as avg_cr,
  AVG(LENGTH(patient_explanation)) as avg_pe,
  AVG(LENGTH(conduct)) as avg_cond,
  AVG(points) as avg_points,
  COUNT(*) FILTER (WHERE last_review > NOW() - INTERVAL '1 day') as enriched_today
FROM score_items
WHERE last_review IS NOT NULL;
```

**Metas:**
- clinical_relevance: 1300-1700 chars
- patient_explanation: 700-850 chars
- conduct: 1100-1400 chars
- points: 15-35 (média)

---

## 🚀 Exemplo Completo

```bash
# Usuário diz: "prepare tudo e enrich tudo"

# Claude executa:
docker compose exec api go run /app/cmd/prepare-all/main.go
# → ✅ Prepared: 142

docker compose exec api go run /app/cmd/list-pending-enrichments/main.go > /tmp/pending.json
# → ✅ JSON gerado

# Claude lê /tmp/pending.json e processa automaticamente:
# [1/142] Zinco → Analisa 10 chunks → Gera 4 campos → PUT API → ✅
# [2/142] Triglicerídeos → Analisa 3 chunks → Gera 4 campos → PUT API → ✅
# ...
# [142/142] ✅ Done!

# Resultado final:
✅ 142 score items enriquecidos com textos científicos de alta qualidade
```
