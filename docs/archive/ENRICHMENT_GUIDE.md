# Score Item Enrichment Guide

Complete guide for the ZERO-API enrichment workflow that processes score items with `last_review IS NULL`.

## Overview

The Plenya project has a sophisticated enrichment system that:
- Fetches score items with NULL `last_review` (paginated with LIMIT 1 OFFSET N)
- Uses LLM (Claude Sonnet 4.5) to enrich clinical fields
- Automatically updates database (BeforeUpdate hook sets `last_review`)
- Generates detailed progress reports with name, tier, and status
- Integrates with PubMed and RAG for article linking

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│  ZERO-API Enrichment Flow                                   │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  1. PostgreSQL Query                                        │
│     SELECT * FROM score_items                               │
│     WHERE last_review IS NULL                               │
│     LIMIT 1 OFFSET N                                        │
│     ↓                                                       │
│  2. Build Context (Automatic)                               │
│     - Hierarchy: Group > Subgroup > Item                    │
│     - Demographics: Gender, Age, Menopause                  │
│     - Lab Test Code, Unit                                   │
│     - Linked Articles (count)                               │
│     ↓                                                       │
│  3. Determine Tier                                          │
│     - PRESERVE: Excellent content (>1000 chars)             │
│     - ENRICH: Has content, needs improvement                │
│     - REWRITE: Empty or minimal content                     │
│     ↓                                                       │
│  4. RAG Article Search (Optional)                           │
│     - Semantic search via pgvector                          │
│     - Auto-link relevant articles                           │
│     ↓                                                       │
│  5. PubMed Search (If needed)                               │
│     - Target: 7+ articles per item                          │
│     - Filters: Reviews, Meta-analyses, 2018-2026            │
│     ↓                                                       │
│  6. LLM Enrichment                                          │
│     - Model: Claude Sonnet 4.5                              │
│     - Generates: clinical_relevance, patient_explanation,   │
│       conduct, max_points, confidence score                 │
│     ↓                                                       │
│  7. Validation                                              │
│     - Length checks (min/max thresholds)                    │
│     - Content quality validation                            │
│     ↓                                                       │
│  8. Audit Trail                                             │
│     - Save to score_item_review_history                     │
│     - Before/after snapshots (JSON)                         │
│     ↓                                                       │
│  9. Database UPDATE                                         │
│     - Direct SQL UPDATE on score_items                      │
│     - BeforeUpdate hook auto-sets last_review               │
│     ↓                                                       │
│  10. Progress Report                                        │
│      [offset/total] Name [Tier: Group > Subgroup] → STATUS  │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

## Available Scripts

### 1. Report Pending Items (Read-Only)
Shows statistics without making changes.

```bash
# Inside container
docker compose exec api go run scripts/report_pending_enrichment.go

# Output
📊 SCORE ITEM ENRICHMENT REPORT
═══════════════════════════════════════════════════════
Total Items:       441
Reviewed:          0 (0.0%)
Pending:           441 (100.0%)
═══════════════════════════════════════════════════════

🎯 ENRICHMENT TIER BREAKDOWN
PRESERVE (excellent):     0 items
ENRICH (improve):        23 items
REWRITE (from scratch): 418 items

📁 PENDING ITEMS BY GROUP
...
```

### 2. Enrich By Offset (Single Item)
Process ONE specific item by offset.

```bash
# Dry-run (no changes)
docker compose exec api go run scripts/enrich_by_offset.go -offset=0 -dry-run

# Enrich item at offset 0
docker compose exec api go run scripts/enrich_by_offset.go -offset=0

# Enrich item at offset 5
docker compose exec api go run scripts/enrich_by_offset.go -offset=5

# Output
📋 ITEM DETAILS
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
ID:              c77cedd3-2800-769d-bd80-42730c51e164
Name:            Luminosidade Natural
Tier:            Social → Atual
Unit:            N/A
Gender:          not_applicable
Age Range:       All ages
Post-Menopause:  N/A

📊 CURRENT ENRICHMENT STATUS
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Clinical Relevance:      0 chars | ❌ EMPTY
Patient Explanation:     0 chars | ❌ EMPTY
Conduct:                 0 chars | ❌ EMPTY
Last Review:          ❌ NEVER REVIEWED

🎯 ENRICHMENT TIER: rewrite
📈 Progress: 1 of 441 items pending review

🧠 CALLING LLM (Model: claude-sonnet-4-5)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
✅ LLM Response received in 2.34 seconds
   Confidence: 0.85

💾 UPDATING SCORE ITEM
✅ Update successful
   Last Review: 2026-02-16 10:30:00
   Points: 5.0

═══════════════════════════════════════════
✅ ENRICHMENT COMPLETE
═══════════════════════════════════════════
Name:            Luminosidade Natural
Tier:            Social → Atual
Status:          ENRICHED (tier=rewrite)
Confidence:      0.85
Processing Time: 2.34 seconds
Progress:        1 of 441 items

💡 Next command:
   go run scripts/enrich_by_offset.go -offset=1
═══════════════════════════════════════════
```

### 3. Batch Enrichment (Multiple Items)
Process multiple items with full RAG + PubMed integration.

```bash
# Process items at offset 5, batch size 1
docker compose exec api go run scripts/batch_enrich_score_items.go 5 1

# Process 10 items starting from offset 0
docker compose exec api go run scripts/batch_enrich_score_items.go 0 10

# Output
🚀 Iniciando batch enrichment (offset=5, limit=1)
================================================================================
📋 Encontrados 1 ScoreItems para processar

────────────────────────────────────────────────────────────────────────────────
🔄 Processando item 1/1: Vitamina B2 por HPLC em sangue total
────────────────────────────────────────────────────────────────────────────────
🔎 Buscando artigos via RAG...
📊 RAG encontrou 3 artigos
🔗 Linkados 3 artigos via RAG
🌐 Buscando 4 artigos no PubMed...
🔍 Query: ("Vitamina B2"[Title/Abstract]) AND ("reference values"[Title/Abstract]) AND (Review[ptyp]) AND 2018:2026[dp]
📋 PMIDs: [12345678, 23456789, 34567890, 45678901]
🤖 Enriquecendo com LLM...
✅ Item enriquecido com sucesso

📊 Resumo [1/1]:
   Nome: Vitamina B2 por HPLC em sangue total
   Tier: rewrite
   Status: success
   Confidence: 0.92
   Artigos: 0 → 7
   Duração: 15s

================================================================================
📊 RELATÓRIO FINAL DO BATCH
================================================================================

📈 Estatísticas:
   Total processado: 1
   ✅ Sucesso: 1
   ⏭️  Skipped: 0
   ❌ Falhas: 0
   📊 Confidence média: 0.92
   ⏱️  Tempo total: 15s
================================================================================
```

### 4. Count Review Status
Quick count of pending vs reviewed items.

```bash
docker compose exec api go run scripts/count_review_status.go
```

## Enrichment Tiers

The system automatically determines the enrichment tier based on existing content:

| Tier | Condition | Action | Description |
|------|-----------|--------|-------------|
| **PRESERVE** | All 3 fields >1000 chars<br/>Recent last_review | Skip | Content is excellent and recent |
| **ENRICH** | Some fields present<br/>Incomplete content | Improve | Enhance existing content with LLM |
| **REWRITE** | Empty or minimal<br/>No last_review | Generate | Create content from scratch |

## Database Schema

### score_items Table
```sql
CREATE TABLE score_items (
    id UUID PRIMARY KEY,
    name VARCHAR(300) NOT NULL,

    -- Enrichment fields (populated by LLM)
    clinical_relevance TEXT,
    patient_explanation TEXT,
    conduct TEXT,

    -- Auto-updated by BeforeUpdate hook when clinical fields change
    last_review TIMESTAMP,

    -- Demographics filters
    gender VARCHAR(20) DEFAULT 'not_applicable',
    age_range_min INTEGER,
    age_range_max INTEGER,
    post_menopause BOOLEAN,

    -- Lab test integration
    lab_test_code VARCHAR(100),
    unit VARCHAR(50),
    unit_conversion TEXT,

    -- Scoring
    points DOUBLE PRECISION,

    -- Hierarchy
    subgroup_id UUID NOT NULL REFERENCES score_subgroups(id),
    parent_item_id UUID REFERENCES score_items(id),

    -- Timestamps
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);
```

### score_item_review_history Table
```sql
CREATE TABLE score_item_review_history (
    id UUID PRIMARY KEY,
    score_item_id UUID NOT NULL REFERENCES score_items(id),
    review_type VARCHAR(50) NOT NULL, -- 'llm_enrichment', 'manual_review'
    before_snapshot JSONB,
    after_snapshot JSONB,
    tier VARCHAR(20), -- 'preserve', 'enrich', 'rewrite'
    confidence_score DOUBLE PRECISION,
    model_used VARCHAR(100), -- 'claude-sonnet-4-5-20250929'
    reviewed_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ
);
```

## Workflow Patterns

### Pattern 1: Single Item Testing
Test enrichment on one specific item before running batch.

```bash
# 1. Find interesting item
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT id, name FROM score_items
WHERE last_review IS NULL
ORDER BY created_at
LIMIT 10 OFFSET 0;
"

# 2. Dry-run to see context
docker compose exec api go run scripts/enrich_by_offset.go -offset=0 -dry-run

# 3. Enrich single item
docker compose exec api go run scripts/enrich_by_offset.go -offset=0

# 4. Verify in database
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT name, last_review,
       LENGTH(clinical_relevance) as clin_len,
       LENGTH(patient_explanation) as pat_len
FROM score_items
WHERE last_review IS NOT NULL
ORDER BY last_review DESC
LIMIT 1;
"
```

### Pattern 2: Batch Processing
Process multiple items with rate limiting.

```bash
# Process 5 items at a time (recommended)
docker compose exec api go run scripts/batch_enrich_score_items.go 0 5

# Continue from offset 5
docker compose exec api go run scripts/batch_enrich_score_items.go 5 5

# Continue from offset 10
docker compose exec api go run scripts/batch_enrich_score_items.go 10 5

# Rate limiting: 5 seconds between items is built-in
```

### Pattern 3: Progressive Enrichment
Enrich items one by one with manual verification.

```bash
# Loop script (bash)
for i in {0..50}; do
    echo "Processing item $i..."
    docker compose exec api go run scripts/enrich_by_offset.go -offset=$i

    # Review result before continuing
    read -p "Continue to next item? (y/n) " -n 1 -r
    echo
    if [[ ! $REGREPLY =~ ^[Yy]$ ]]; then
        break
    fi
done
```

### Pattern 4: Group-Specific Enrichment
Enrich items from specific groups.

```bash
# Find offset for specific group
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT
    ROW_NUMBER() OVER (ORDER BY si.created_at) - 1 as offset,
    si.name,
    sg.name as group_name
FROM score_items si
JOIN score_subgroups ss ON si.subgroup_id = ss.id
JOIN score_groups sg ON ss.group_id = sg.id
WHERE si.last_review IS NULL
  AND sg.name = 'Exames'
ORDER BY si.created_at;
"

# Enrich specific group items by offset
```

## Best Practices

### 1. Start Small
- Begin with dry-run to understand tier distribution
- Process 1-5 items first to verify quality
- Review generated content before scaling up

### 2. Monitor Progress
```bash
# Check progress periodically
docker compose exec api go run scripts/count_review_status.go

# View recent enrichments
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT name, last_review,
       LENGTH(clinical_relevance) + LENGTH(patient_explanation) + LENGTH(conduct) as total_chars
FROM score_items
WHERE last_review IS NOT NULL
ORDER BY last_review DESC
LIMIT 10;
"
```

### 3. Rate Limiting
- **Batch script**: 5 seconds between items (built-in)
- **Single item script**: No delay (manual control)
- **API limits**: Claude API has rate limits (~50 requests/min)
- For 441 items: ~2.5 hours with batch script

### 4. Error Handling
- Review history is saved even if update fails
- Failed items remain with `last_review IS NULL`
- Re-run same offset to retry failed items
- Check logs for specific error messages

### 5. Content Validation
After enrichment, verify quality:

```bash
# Check empty fields
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT COUNT(*) as empty_fields
FROM score_items
WHERE last_review IS NOT NULL
  AND (clinical_relevance IS NULL
    OR patient_explanation IS NULL
    OR conduct IS NULL);
"

# Check field lengths
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT
    AVG(LENGTH(clinical_relevance)) as avg_clin,
    AVG(LENGTH(patient_explanation)) as avg_pat,
    AVG(LENGTH(conduct)) as avg_cond
FROM score_items
WHERE last_review IS NOT NULL;
"
```

## Cost Estimation

### LLM Costs (Claude Sonnet 4.5)
- Input: $3 per 1M tokens
- Output: $15 per 1M tokens
- Avg per item: ~800 input + ~400 output tokens
- **Cost per item**: ~$0.01
- **Total for 441 items**: ~$4.40

### PubMed Integration
- Free API (NCBI E-utilities)
- PDF downloads may fail (paywalls)
- Rate limit: 3 requests/second (with API key)

### Time Estimation
- Single item: ~2-5 seconds (LLM only)
- Batch with PubMed: ~15-30 seconds per item
- **Total time for 441 items**:
  - Fast mode (no PubMed): ~30 minutes
  - Full mode (with PubMed): ~2-3 hours

## Troubleshooting

### "Failed to connect to database"
```bash
# Check Docker
docker compose ps

# Restart database
docker compose restart db

# Wait for healthy status
docker compose ps db
```

### "ANTHROPIC_API_KEY not set"
```bash
# Check .env file
cat /home/user/plenya/.env | grep ANTHROPIC

# Or export manually
export ANTHROPIC_API_KEY="sk-ant-..."
docker compose restart api
```

### "LLM enrichment failed"
- Check API key validity
- Verify internet connection
- Check rate limits (wait 60 seconds)
- Review error message in logs

### "Validation warnings"
- Content too short: LLM didn't generate enough text
- Content too long: LLM was too verbose
- Warnings don't block update, but indicate quality issues

### Items already enriched are being skipped
This is correct behavior when tier=PRESERVE:
```
✨ Content is excellent and recent - PRESERVING as-is
```

To force re-enrichment:
```sql
-- Reset last_review for specific item
UPDATE score_items
SET last_review = NULL
WHERE id = 'uuid-here';
```

## Manual Database Operations

### Query pending items
```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT
    COUNT(*) as total,
    COUNT(*) FILTER (WHERE last_review IS NULL) as pending,
    COUNT(*) FILTER (WHERE last_review IS NOT NULL) as reviewed
FROM score_items;
"
```

### View enrichment history
```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT
    si.name,
    rh.tier,
    rh.confidence_score,
    rh.reviewed_at
FROM score_item_review_history rh
JOIN score_items si ON rh.score_item_id = si.id
ORDER BY rh.reviewed_at DESC
LIMIT 10;
"
```

### Reset enrichment for testing
```bash
# DANGER: Only for testing!
docker compose exec db psql -U plenya_user -d plenya_db -c "
UPDATE score_items
SET last_review = NULL,
    clinical_relevance = NULL,
    patient_explanation = NULL,
    conduct = NULL
WHERE id IN (
    SELECT id FROM score_items
    WHERE last_review IS NOT NULL
    LIMIT 5
);
"
```

## Integration with Frontend

The enriched content is automatically available via the API:

```typescript
// Web: apps/web/lib/api/score-api.ts
const scoreItem = await fetchScoreItem(itemId);

// Contains enriched fields
scoreItem.clinicalRelevance   // For healthcare professionals
scoreItem.patientExplanation  // For patients
scoreItem.conduct             // Clinical recommendations
scoreItem.lastReview          // Last enrichment date
```

## Related Documentation

- [Database Operations](./.claude/workflows/database-ops.md)
- [Score System](./.claude/domain/score-system.md)
- [Backend Models](./.claude/backend/models.md)
- [Service Layer](./.claude/backend/service-layer.md)

---

**Last Updated**: February 2026
**Version**: 1.0
**Status**: 441 items pending enrichment
