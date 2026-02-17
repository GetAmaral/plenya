# Method Content Enrichment Scripts

Enrich method letters and pillars with AI-generated clinical content.

## Overview

The enrichment workflow follows the ZERO-API pattern:
1. Query unenriched items with `last_review IS NULL LIMIT 1 OFFSET N`
2. Enrich using OpenAI GPT-4
3. UPDATE automatically updates `last_review` via BeforeUpdate hook
4. Report name, tier, and status

## Scripts

### 1. Check Enrichment Status

Check how many items need enrichment:

```bash
# Check all tiers
docker compose exec api go run scripts/check_enrichment_status.go --tier all

# Check only letters
docker compose exec api go run scripts/check_enrichment_status.go --tier letter

# Check only pillars
docker compose exec api go run scripts/check_enrichment_status.go --tier pillar
```

**Output:**
- Total count
- Enriched count and percentage
- Unenriched count and percentage
- Next 5 unenriched items

### 2. Enrich Method Content

Enrich a single item using AI:

```bash
# Enrich first unenriched letter (offset 0)
docker compose exec api /bin/sh -c "export OPENAI_API_KEY=your-key && go run scripts/enrich_method_content.go --tier letter --offset 0"

# Enrich 10th unenriched pillar (offset 9)
docker compose exec api /bin/sh -c "export OPENAI_API_KEY=your-key && go run scripts/enrich_method_content.go --tier pillar --offset 9"

# Dry run (preview without saving)
docker compose exec api /bin/sh -c "export OPENAI_API_KEY=your-key && go run scripts/enrich_method_content.go --tier letter --offset 0 --dry-run"
```

**Flags:**
- `--tier letter|pillar` - Tier to enrich (required)
- `--offset N` - Pagination offset (0-based, default: 0)
- `--dry-run` - Preview without updating database

**Output:**
- Item name, code/letter, method
- Tier (Letter or Pillar)
- Status (Unenriched → Enriched)
- Generated content:
  - Clinical Relevance (why it matters)
  - Patient Explanation (simple language)
  - Conduct (clinical recommendations)

## Environment Variables

Set `OPENAI_API_KEY` environment variable:

```bash
export OPENAI_API_KEY=sk-...
```

Or add to `.env`:

```
OPENAI_API_KEY=sk-...
```

## Enrichment Fields

The enrichment generates three fields in Brazilian Portuguese:

1. **Clinical Relevance** (`clinical_relevance`)
   - 2-3 sentences explaining clinical importance
   - Evidence-based rationale
   - Focus on patient outcomes

2. **Patient Explanation** (`patient_explanation`)
   - 2-3 sentences in simple language
   - Avoid medical jargon
   - Patient-friendly explanation

3. **Conduct** (`conduct`)
   - 2-3 bullet points
   - Specific clinical actions
   - Example: "• Solicitar hemograma completo e perfil lipídico"

## Workflow Example

```bash
# 1. Check status
docker compose exec api go run scripts/check_enrichment_status.go --tier all

# Output shows:
# === METHOD LETTERS ===
# Total:      4
# Enriched:   0 (0.0%)
# Unenriched: 4 (100.0%)
# Next 5 unenriched letters:
#   [0] AGIR - Alimentação e Atividade Física (A)
#   [1] AGIR - Gestão Metabólica (G)
#   ...

# 2. Enrich first letter
docker compose exec api /bin/sh -c "export OPENAI_API_KEY=sk-... && go run scripts/enrich_method_content.go --tier letter --offset 0"

# Output shows:
# === ENRICHING METHOD LETTER ===
# Name: Alimentação e Atividade Física
# Code: A
# Method: AGIR
# Tier: Letter
# Status: Unenriched (last_review = NULL)
# Offset: 0
#
# === ENRICHMENT COMPLETE ===
# Status: Enriched (last_review auto-updated via hook)
# Letter ID: fe9915ad-8d57-45a5-a3b5-533a36817cbb

# 3. Enrich next letter
docker compose exec api /bin/sh -c "export OPENAI_API_KEY=sk-... && go run scripts/enrich_method_content.go --tier letter --offset 0"
# (offset 0 again, because the first item is now enriched and skipped)

# 4. Check progress
docker compose exec api go run scripts/check_enrichment_status.go --tier letter

# Output shows:
# === METHOD LETTERS ===
# Total:      4
# Enriched:   2 (50.0%)
# Unenriched: 2 (50.0%)
```

## Database Hooks

The enrichment leverages GORM hooks (see `internal/models/method_letter.go` and `method_pillar.go`):

```go
func (ml *MethodLetter) BeforeUpdate(tx *gorm.DB) error {
    if tx.Statement.Changed("ClinicalRelevance") ||
        tx.Statement.Changed("PatientExplanation") ||
        tx.Statement.Changed("Conduct") {
        now := time.Now()
        ml.LastReview = &now
    }
    return nil
}
```

This automatically sets `last_review` when clinical fields are updated.

## Troubleshooting

### Error: "OPENAI_API_KEY environment variable not set"
- Set the environment variable before running the script
- Add it to `.env` or export it in your shell

### Error: "record not found"
- All items at the given offset are already enriched
- Check status with `check_enrichment_status.go`
- Use offset 0 to always get the next unenriched item

### Dry run shows unexpected content
- Adjust the prompt in `enrichLetterWithClaude` or `enrichPillarWithClaude`
- The prompts are in `scripts/enrich_method_content.go`
- Use dry-run mode to preview before saving

## Future Enhancements

- Batch enrichment (enrich N items at once)
- Support for other AI models (Claude, local models)
- Enrichment revision workflow
- Quality scoring for generated content
