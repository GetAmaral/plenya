# Score Item Enrichment Tool

Zero-API approach for enriching score items with clinical content using LLM.

## Overview

This tool processes score items with `last_review IS NULL`, enriching them with:
- **Clinical Relevance**: Technical explanation for healthcare professionals
- **Patient Explanation**: Simple language for patients
- **Conduct**: Clinical recommendations and guidelines

## Architecture

```
┌─────────────────────────────────────────────────────────┐
│  Enrichment Flow (Zero-API)                             │
├─────────────────────────────────────────────────────────┤
│                                                         │
│  1. PostgreSQL Query                                    │
│     SELECT ... LIMIT 1 OFFSET N                         │
│     WHERE last_review IS NULL                           │
│     ↓                                                   │
│  2. Build Context                                       │
│     - Group > Subgroup > Item (tier)                    │
│     - Demographic filters                               │
│     - Unit of measurement                               │
│     ↓                                                   │
│  3. LLM Enrichment                                      │
│     Anthropic Claude 3.5 Sonnet                         │
│     (Medical domain specialist)                         │
│     ↓                                                   │
│  4. Direct DB UPDATE                                    │
│     UPDATE score_items SET                              │
│       clinical_relevance = ...,                         │
│       patient_explanation = ...,                        │
│       conduct = ...,                                    │
│       last_review = NOW()                               │
│     ↓                                                   │
│  5. Progress Report                                     │
│     [offset/total] name [tier] → STATUS                 │
│                                                         │
└─────────────────────────────────────────────────────────┘
```

## Prerequisites

1. **Database running**:
   ```bash
   docker compose up -d db
   ```

2. **Anthropic API Key**:
   ```bash
   export ANTHROPIC_API_KEY="your-api-key-here"
   ```

3. **Go dependencies** (auto-installed by script):
   - `github.com/liushuangls/go-anthropic/v2`

## Usage

### Basic Usage

Process 10 items starting from offset 0:
```bash
./scripts/enrich_score_items.sh
```

### Custom Offset and Batch

Process 5 items starting from offset 20:
```bash
./scripts/enrich_score_items.sh 20 5
```

### Dry Run Mode

Test without database updates:
```bash
./scripts/enrich_score_items.sh 0 3 true
```

### Process All Items (Incremental)

```bash
# Check total count first
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT COUNT(*) FROM score_items WHERE last_review IS NULL;"

# Process in batches of 50
./scripts/enrich_score_items.sh 0 50
./scripts/enrich_score_items.sh 50 50
./scripts/enrich_score_items.sh 100 50
# ... continue until all processed
```

## Output Format

```
=== Score Item Enrichment Tool ===
Total items with NULL last_review: 441
Starting offset: 0
Batch size: 10
Dry run mode: false

[1/10] Processing item at offset 0...
ENRICHING: Luminosidade Natural [Social > Atual]
SUCCESS: Luminosidade Natural [Social > Atual]

[2/10] Processing item at offset 1...
ENRICHING: Vitamina B2 por HPLC em sangue total [Exames > Laboratoriais]
SUCCESS: Vitamina B2 por HPLC em sangue total [Exames > Laboratoriais]

...

=== Enrichment Summary ===
Processed: 10 items
Success: 8
Skipped: 1
Errors: 1
Remaining: 431
```

## Status Codes

- **SUCCESS**: Item enriched and updated successfully
- **SKIPPED**: Item already has all enrichment fields filled
- **ERROR**: LLM or database error (see logs for details)

## Database Schema

The tool updates these fields in `score_items`:

```sql
-- Enrichment fields (all TEXT)
clinical_relevance   -- Technical explanation
patient_explanation  -- Patient-friendly text
conduct             -- Clinical recommendations

-- Auto-updated by BeforeUpdate hook
last_review         -- Timestamp of enrichment
updated_at          -- Standard GORM timestamp
```

## Rate Limiting

- **1 second delay** between LLM calls
- Respects Anthropic API rate limits
- For large batches, consider multiple smaller runs

## Error Handling

### Common Errors

1. **"ANTHROPIC_API_KEY environment variable not set"**
   ```bash
   export ANTHROPIC_API_KEY="sk-ant-..."
   ```

2. **"Failed to connect to database"**
   ```bash
   docker compose up -d db
   # Wait 5 seconds for DB to be ready
   ```

3. **"API call failed: rate limit exceeded"**
   - Wait 60 seconds
   - Reduce batch size
   - Use smaller batches with delays

4. **"failed to parse JSON response"**
   - LLM returned invalid JSON
   - Check logs for raw response
   - Item will be retried on next run

## Verification

Check enriched items:
```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT
  name,
  last_review,
  LENGTH(clinical_relevance) as clin_len,
  LENGTH(patient_explanation) as pat_len,
  LENGTH(conduct) as cond_len
FROM score_items
WHERE last_review IS NOT NULL
ORDER BY last_review DESC
LIMIT 5;
"
```

View specific enrichment:
```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT
  name,
  clinical_relevance,
  patient_explanation,
  conduct
FROM score_items
WHERE name = 'Hemoglobina - Homens';
"
```

## Development

Run directly with Go:
```bash
cd apps/api

# Dry run with verbose logging
go run cmd/enrich_score_items.go -offset=0 -batch=1 -dry-run

# Process single item
go run cmd/enrich_score_items.go -offset=4 -batch=1

# Process batch
go run cmd/enrich_score_items.go -offset=10 -batch=20
```

## Best Practices

1. **Start with dry run**: Test the prompt and LLM responses
2. **Small batches first**: Process 5-10 items to verify quality
3. **Review samples**: Check database for content quality
4. **Incremental processing**: Use offset to continue from where you left off
5. **Monitor costs**: Track Anthropic API usage

## Cost Estimation

- Model: Claude 3.5 Sonnet
- Avg tokens per item: ~500 input + 300 output
- Cost per 1M tokens: $3 (input) + $15 (output)
- **Estimated cost per 100 items**: ~$0.50

For 441 items: ~$2.20 total

## Troubleshooting

### Items not being enriched

Check if items exist:
```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT id, name, last_review
FROM score_items
WHERE last_review IS NULL
LIMIT 5;
"
```

### LLM generating poor content

1. Adjust temperature in `enrich_score_items.go` (currently 0.3)
2. Modify prompt in `buildEnrichmentPrompt()` function
3. Test with dry-run first

### Database connection issues

Check Docker:
```bash
docker compose ps
docker compose logs db
```

Check connection:
```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "SELECT 1;"
```

## Related Documentation

- [Database Operations](./.claude/workflows/database-ops.md)
- [Score System](./.claude/domain/score-system.md)
- [Backend Models](./.claude/backend/models.md)

---

**Last Updated**: February 2026
**Version**: 1.0
