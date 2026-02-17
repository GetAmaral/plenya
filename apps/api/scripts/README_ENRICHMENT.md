# Score Item Enrichment Scripts

ZERO-API workflow for enriching clinical fields in score items.

## Overview

These scripts implement a **direct database manipulation** workflow (development only) for:
1. Finding score items with `last_review IS NULL`
2. Enriching them manually or via LLM
3. Auto-updating via GORM hooks (sets `last_review` automatically)
4. Generating detailed reports

## Two Approaches

### 🔹 Approach 1: Manual Enrichment (Recommended)
Use `enrich_next_item.go` for full control with external enrichment tools.

### 🔹 Approach 2: Automated LLM Enrichment
Use legacy `enrich_by_offset.go` for fully automated Claude API calls.

---

## Scripts

### 0. `count_review_status.go` ⭐ NEW
**Purpose:** Quick overview of review status and tier distribution.

**Usage:**
```bash
docker compose exec -T api go run scripts/count_review_status.go
```

**Output:**
```
====================================
SCORE ITEMS REVIEW STATUS
====================================

📊 Total Items:       878
✅ Reviewed:          439 (50.0%)
⏳ Not Reviewed:      439 (50.0%)
📝 Empty (no content): 69

====================================
NOT REVIEWED BY TIER
====================================

🔴 Tier 3 (REWRITE):  146 (33.3%)
🟡 Tier 2 (ENRICH):   172 (39.2%)
🟢 Tier 1 (PRESERVE): 121 (27.6%)
```

### 1. `enrich_next_item.go` ⭐ NEW - Manual Enrichment
**Purpose:** Fetch context → Manually enrich → Apply via GORM with auto-update.

**Usage:**
```bash
# STEP 1: Fetch enrichment context
docker compose exec -T api go run scripts/enrich_next_item.go <OFFSET>

# Examples:
docker compose exec -T api go run scripts/enrich_next_item.go 0   # First not-reviewed
docker compose exec -T api go run scripts/enrich_next_item.go 2   # Third not-reviewed

# STEP 2: Apply enrichment
docker compose exec -T api go run scripts/enrich_next_item.go apply enrichment.json
```

**Query Pattern:**
```sql
SELECT * FROM score_items
WHERE last_review IS NULL
ORDER BY created_at ASC
LIMIT 1 OFFSET <OFFSET>
```

**Context Output:**
```json
{
  "score_item": {
    "id": "c77cedd3-2800-70c7-942e-a1134e3aa05e",
    "name": "Doença cardiovascular (IAM, revascularização...)",
    "clinical_relevance": "...",
    "patient_explanation": null,
    "conduct": "...",
    "last_review": null
  },
  "articles": [
    {
      "title": "Cardiologia II",
      "similarity": 0.80,
      "abstract": "..."
    }
  ],
  "tier": "enrich",
  "status": "not_reviewed"
}
```

**Enrichment Payload Format:**
```json
{
  "score_item_id": "c77cedd3-2800-70c7-942e-a1134e3aa05e",
  "clinical_relevance": "Full text 800-3000 chars...",
  "patient_explanation": "Patient-friendly 400-2000 chars...",
  "conduct": "Clinical protocols 600-6000 chars...",
  "max_points": 35,
  "justification": "Why this point value...",
  "confidence": 0.92,
  "tier": "enrich"
}
```

**Apply Output:**
```
====================================
✅ ENRICHMENT APPLIED SUCCESSFULLY
====================================

📋 Name:        Doença cardiovascular (IAM...)
🔗 ID:          c77cedd3-2800-70c7-942e-a1134e3aa05e

🎯 Tier:        enrich → enrich
📊 Status:      not_reviewed → partial
💯 Points:      3.0 → 35
🔍 Confidence:  0.92
📅 Last Review: 2026-02-16 20:30:45

📝 Field Lengths:
   - Clinical Relevance:  1847 chars
   - Patient Explanation: 612 chars
   - Conduct:             1523 chars

✅ BeforeUpdate hook triggered successfully
   LastReview auto-updated to: 2026-02-16 20:30:45
```

**Advantages:**
- Full control over enrichment process
- Use any LLM (Claude, ChatGPT, etc.)
- Review before applying
- No API costs from scripts
- Transparent workflow

### 2. `report_pending_enrichment.go`
**Purpose:** Generate comprehensive report of all pending enrichments.

**Usage:**
```bash
docker compose exec -T api go run /app/scripts/report_pending_enrichment.go
```

**Output:**
- Overall statistics (total, reviewed, pending)
- Tier breakdown (preserve, enrich, rewrite)
- Items by score group
- First 20 pending items
- Next steps instructions

### 3. `enrich_by_offset.go` (Legacy - Automated)
**Purpose:** Enrich ONE score item using pagination (LIMIT 1 OFFSET n).

**Usage:**
```bash
# Dry-run (review without changes)
docker compose exec -T api go run /app/scripts/enrich_by_offset.go -offset 0 -dry-run

# Actual enrichment
docker compose exec -T api go run /app/scripts/enrich_by_offset.go -offset 0
```

**Parameters:**
- `-offset N`: Pagination offset (0-based index)
- `-dry-run`: Preview item without enriching

**Process Flow:**
1. **FETCH:** `SELECT * WHERE last_review IS NULL ORDER BY name LIMIT 1 OFFSET n`
2. **ANALYZE:** Determine enrichment tier (preserve/enrich/rewrite)
3. **ENRICH:** Call Claude API with linked articles context
4. **VALIDATE:** Check field lengths and quality
5. **SAVE HISTORY:** Create audit snapshot
6. **UPDATE:** Automatic update (BeforeUpdate hook sets `last_review`)
7. **REPORT:** Display results and next command

**Output Example:**
```
═══════════════════════════════════════════
✅ ENRICHMENT COMPLETE
═══════════════════════════════════════════
Name:            Hemoglobina - Homens
Tier:            Exames → Hematologia
Status:          ENRICHED (tier=enrich)
Confidence:      0.94
Processing Time: 8.52 seconds
Progress:        1 of 439 items

💡 Next command:
   go run enrich_by_offset.go -offset 1
═══════════════════════════════════════════
```

## Enrichment Tiers

The system automatically determines the appropriate enrichment strategy:

| Tier | Criteria | Strategy | Items |
|------|----------|----------|-------|
| **PRESERVE** ✨ | CR≥1500, PE≥600, Conduct≥800, Recent (<6mo) | Skip (excellent) | 0 |
| **ENRICH** 📝 | CR≥500 but not excellent | Improve incrementally | 298 |
| **REWRITE** 🔨 | CR<500 or missing | Generate from scratch | 141 |

**Legend:**
- CR = Clinical Relevance
- PE = Patient Explanation

## Batch Processing

### Manual Loop (Recommended)
Process multiple items with human supervision:

```bash
# Process 10 items with 5-second delay
for i in {0..9}; do
  docker compose exec -T api go run /app/scripts/enrich_by_offset.go -offset $i
  sleep 5
done
```

### Process Specific Ranges
```bash
# Process items 50-59
for i in {50..59}; do
  docker compose exec -T api go run /app/scripts/enrich_by_offset.go -offset $i
  sleep 5
done
```

### Continuous Processing
```bash
# Process all pending items (use with caution - API costs!)
for i in $(seq 0 438); do
  docker compose exec -T api go run /app/scripts/enrich_by_offset.go -offset $i
  sleep 10  # Rate limiting
done
```

## Important Notes

### ⚠️ Development Only
These scripts use **direct database access**, which is appropriate for:
- ✅ Development/testing
- ✅ Batch data operations
- ✅ Manual enrichment workflows

**NEVER use in production.** Production apps must use HTTP API.

### 🔐 API Key Required
Set `ANTHROPIC_API_KEY` environment variable:

```bash
# In docker-compose.yml or .env
ANTHROPIC_API_KEY=sk-ant-...
```

### 💰 Cost Management
Each enrichment costs approximately:
- **Input tokens:** ~2,000 tokens ($6/MTok)
- **Output tokens:** ~1,500 tokens ($24/MTok)
- **Per item:** ~$0.05-0.10

**439 items × $0.075 ≈ $33 total**

Use dry-run mode liberally to preview before enriching.

### 🪝 Automatic last_review
The `BeforeUpdate` hook in `score_item.go` automatically sets `last_review` when:
- `clinical_relevance` changes
- `patient_explanation` changes
- `conduct` changes

No manual timestamp management needed!

### 📊 Review History
Every enrichment is logged in `score_item_review_history`:
- Before/after snapshots (JSON)
- Tier used
- Confidence score
- Model ID
- Timestamp

Access via:
```sql
SELECT * FROM score_item_review_history
WHERE score_item_id = '...'
ORDER BY reviewed_at DESC;
```

## Workflow Example

### Complete Workflow
```bash
# 1. Generate initial report
docker compose exec -T api go run /app/scripts/report_pending_enrichment.go

# 2. Review first item (dry-run)
docker compose exec -T api go run /app/scripts/enrich_by_offset.go -offset 0 -dry-run

# 3. Enrich first item
docker compose exec -T api go run /app/scripts/enrich_by_offset.go -offset 0

# 4. Continue with next items
docker compose exec -T api go run /app/scripts/enrich_by_offset.go -offset 1
docker compose exec -T api go run /app/scripts/enrich_by_offset.go -offset 2

# 5. Check progress
docker compose exec -T api go run /app/scripts/report_pending_enrichment.go
```

### Targeting Specific Groups
To enrich items from a specific group, modify SQL in `enrich_by_offset.go`:

```go
// Instead of:
Where("last_review IS NULL")

// Use:
Where("last_review IS NULL").
Joins("JOIN score_subgroups ss ON score_items.subgroup_id = ss.id").
Joins("JOIN score_groups sg ON ss.group_id = sg.id").
Where("sg.name = ?", "Exames")
```

## Troubleshooting

### "API key not found"
```bash
# Check environment variable
docker compose exec api env | grep ANTHROPIC

# Set if missing (docker-compose.yml)
services:
  api:
    environment:
      - ANTHROPIC_API_KEY=${ANTHROPIC_API_KEY}
```

### "Empty response from API"
- Check API key validity
- Verify network connectivity
- Check Anthropic service status

### "Validation warnings"
The LLM sometimes generates content outside ideal ranges. Warnings include:
- Field too short/long
- Placeholders detected (`[`, `TODO`, `...`)
- Max points out of range

These are **soft warnings** - review will still save unless critical errors occur.

### "Failed to fetch score item"
No more items with `last_review IS NULL` at that offset:
```bash
# Check remaining count
docker compose exec db psql -U plenya_user -d plenya_db \
  -c "SELECT COUNT(*) FROM score_items WHERE last_review IS NULL;"
```

## Database Queries

### Check enrichment status
```sql
-- Overall stats
SELECT
  COUNT(*) as total,
  COUNT(last_review) as reviewed,
  COUNT(*) - COUNT(last_review) as pending
FROM score_items;

-- By group
SELECT
  sg.name,
  COUNT(*) as total,
  COUNT(si.last_review) as reviewed,
  COUNT(*) - COUNT(si.last_review) as pending
FROM score_items si
JOIN score_subgroups ss ON si.subgroup_id = ss.id
JOIN score_groups sg ON ss.group_id = sg.id
GROUP BY sg.name
ORDER BY pending DESC;
```

### View recent enrichments
```sql
SELECT
  si.name,
  si.last_review,
  h.tier,
  h.confidence_score,
  h.reviewed_at
FROM score_items si
LEFT JOIN score_item_review_history h ON si.id = h.score_item_id
WHERE si.last_review IS NOT NULL
ORDER BY si.last_review DESC
LIMIT 20;
```

### Reset enrichment (for testing)
```sql
-- CAUTION: This removes enrichment data
UPDATE score_items
SET
  clinical_relevance = 'N/A',
  patient_explanation = 'N/A',
  conduct = 'N/A',
  last_review = NULL
WHERE id = '...';
```

## Architecture Notes

This workflow follows Plenya's **golden rules**:

1. ✅ **Go models are source of truth** - No direct SQL edits
2. ✅ **Development uses direct DB** - Scripts use GORM
3. ✅ **Hooks are respected** - BeforeUpdate auto-sets last_review
4. ✅ **Audit trail** - Review history table preserves snapshots

See `.claude/workflows/database-ops.md` for more on development workflows.

---

**Last Updated:** 2026-02-16
**Author:** Claude Sonnet 4.5
