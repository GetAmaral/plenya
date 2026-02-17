# Zero-API Enrichment Workflow

**Quick reference for enriching ScoreItems without HTTP APIs.**

## Core Command Pattern

```bash
# 1. CHECK STATUS
docker compose exec -T api go run scripts/count_review_status.go

# 2. FETCH CONTEXT (LIMIT 1 OFFSET N)
docker compose exec -T api go run scripts/enrich_next_item.go <OFFSET>

# 3. APPLY ENRICHMENT (AUTO-UPDATE last_review)
docker compose exec -T api go run scripts/enrich_next_item.go apply enrichment.json
```

## Query Pattern

The script uses this exact pattern you requested:

```sql
SELECT * FROM score_items
WHERE last_review IS NULL
ORDER BY created_at ASC
LIMIT 1 OFFSET <OFFSET>
```

## Step-by-Step Example

### Step 1: Get the 3rd un-reviewed item (OFFSET=2)

```bash
docker compose exec -T api go run scripts/enrich_next_item.go 2
```

**Output:**
```
====================================
ENRICHMENT CONTEXT (OFFSET=2)
====================================

📋 Name:   Doença cardiovascular (IAM, revascularização, arritmias, AVC)
🎯 Tier:   enrich
📊 Status: not_reviewed
🔗 ID:     c77cedd3-2800-70c7-942e-a1134e3aa05e
📚 Articles: 19 linked

{
  "score_item": { ... },
  "articles": [ ... ],
  "tier": "enrich",
  "status": "not_reviewed"
}
```

### Step 2: Enrich manually

Use the JSON context with Claude, ChatGPT, or manual editing.

Create `enrichment.json`:

```json
{
  "score_item_id": "c77cedd3-2800-70c7-942e-a1134e3aa05e",
  "clinical_relevance": "Detailed clinical text 800-3000 chars...",
  "patient_explanation": "Patient-friendly text 400-2000 chars...",
  "conduct": "Clinical protocols 600-6000 chars...",
  "max_points": 35,
  "justification": "High cardiovascular risk factor",
  "confidence": 0.92,
  "tier": "enrich"
}
```

### Step 3: Apply enrichment

```bash
docker compose exec -T api go run scripts/enrich_next_item.go apply enrichment.json
```

**Output:**
```
====================================
✅ ENRICHMENT APPLIED SUCCESSFULLY
====================================

📋 Name:        Doença cardiovascular (IAM...)
🎯 Tier:        enrich → enrich
📊 Status:      not_reviewed → partial
📅 Last Review: 2026-02-16 20:30:45

✅ BeforeUpdate hook triggered successfully
   LastReview auto-updated to: 2026-02-16 20:30:45
```

## Report: Name, Tier, Status

The script automatically reports:

### During Fetch (OFFSET command):
- **Name:** ScoreItem name
- **Tier:** preserve | enrich | rewrite
- **Status:** empty | not_reviewed | partial | complete | stale
- **ID:** UUID
- **Articles:** Count of linked articles

### During Apply:
- **Name:** ScoreItem name
- **Tier Before → After:** Classification change
- **Status Before → After:** Review status change
- **Last Review:** Auto-updated timestamp (proof hook triggered)

## Tier Definitions

| Tier | Criteria | Action |
|------|----------|--------|
| **PRESERVE** | CR≥1500, PE≥600, Conduct≥800, Recent | Skip |
| **ENRICH** | Has content but incomplete | Improve |
| **REWRITE** | CR<500 or empty | Generate from scratch |

## Status Definitions

| Status | Meaning |
|--------|---------|
| **empty** | All clinical fields NULL/empty |
| **not_reviewed** | Has content but `last_review IS NULL` |
| **partial** | Reviewed but doesn't meet thresholds |
| **complete** | Reviewed + meets all thresholds |
| **stale** | Reviewed but >180 days ago |

## Auto-UPDATE Mechanism

The `BeforeUpdate` hook in `/home/user/plenya/apps/api/internal/models/score_item.go` automatically sets `last_review = NOW()` when:

- `clinical_relevance` changes
- `patient_explanation` changes
- `conduct` changes

**No manual UPDATE needed!**

## Validation Rules

Applied during `apply` command:

- `clinical_relevance`: 800-3000 chars
- `patient_explanation`: 400-2000 chars
- `conduct`: 600-6000 chars
- `max_points`: 0-50
- `confidence`: 0.0-1.0
- `tier`: preserve | enrich | rewrite

## Current Status (Feb 16, 2026)

```
📊 Total Items:       878
✅ Reviewed:          439 (50.0%)
⏳ Not Reviewed:      439 (50.0%)

🔴 Tier 3 (REWRITE):  146 (33.3%)
🟡 Tier 2 (ENRICH):   172 (39.2%)
🟢 Tier 1 (PRESERVE): 121 (27.6%)
```

## Batch Processing

```bash
# Process first 10 items
for i in {0..9}; do
  echo "=== Fetching item $i ==="
  docker compose exec -T api go run scripts/enrich_next_item.go $i > context_$i.json

  # Manual enrichment here (use Claude, etc.)
  # Then apply:
  # docker compose exec -T api go run scripts/enrich_next_item.go apply enrichment_$i.json
done
```

## Files

- **Script:** `/home/user/plenya/apps/api/scripts/enrich_next_item.go`
- **Count Script:** `/home/user/plenya/apps/api/scripts/count_review_status.go`
- **Model Hook:** `/home/user/plenya/apps/api/internal/models/score_item.go:120-130`
- **Full Docs:** `/home/user/plenya/apps/api/scripts/README_ENRICHMENT.md`

---

**Created:** 2026-02-16
**Pattern:** `LIMIT 1 OFFSET N` → Enrich → Auto-UPDATE
