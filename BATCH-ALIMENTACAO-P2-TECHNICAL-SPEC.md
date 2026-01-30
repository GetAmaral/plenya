# Batch ALIMENTAÇÃO Parte 2 - Technical Specification

## System Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                   ENRICHMENT PIPELINE                        │
└─────────────────────────────────────────────────────────────┘

Input Layer:
  ├─ 20 score_items (IDs + current questions)
  ├─ Clinical context per item
  └─ Group classification: ALIMENTAÇÃO

Processing Layer:
  ├─ Python orchestrator (batch_alimentacao_parte2.py)
  ├─ Anthropic API client (SDK)
  ├─ Claude Opus 4.5 model
  │   ├─ Extended thinking: 10k tokens
  │   ├─ Temperature: 1.0
  │   └─ Max output: 16k tokens
  └─ Rate limiting: 2s delay between requests

Output Layer:
  ├─ JSON file (structured data)
  ├─ SQL file (PostgreSQL migration)
  └─ Execution logs (console + file)

Database Layer:
  └─ PostgreSQL 17
      └─ UPDATE score_items SET ... WHERE id = ...
```

## Data Model

### Input Schema (score_items)
```sql
CREATE TABLE score_items (
  id UUID PRIMARY KEY,
  title VARCHAR(200) NOT NULL,
  question TEXT,
  group_name VARCHAR(100),
  -- Enrichment fields (to be populated)
  clinical_relevance TEXT,
  interpretation_guide TEXT,
  health_implications TEXT[],
  followup_questions TEXT[],
  red_flags TEXT[],
  recommendations TEXT[],
  scientific_background TEXT,
  -- Metadata
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);
```

### Output Schema (enrichment object)
```typescript
interface ScoreItemEnrichment {
  question: string;                    // 50-200 chars
  clinical_relevance: string;          // 200-400 chars
  interpretation_guide: string;        // 300-500 chars
  health_implications: string[];       // 4-6 items, 80-150 chars each
  followup_questions: string[];        // 3-5 items, 50-100 chars each
  red_flags: string[];                 // 3-4 items, 60-120 chars each
  recommendations: string[];           // 4-6 items, 100-200 chars each
  scientific_background: string;       // 200-400 chars
}
```

## API Integration

### Anthropic API Configuration
```python
client = anthropic.Anthropic(api_key=API_KEY)

message = client.messages.create(
    model="claude-opus-4-5-20250929",
    max_tokens=16000,
    temperature=1,
    thinking={
        "type": "enabled",
        "budget_tokens": 10000
    },
    messages=[{
        "role": "user",
        "content": prompt
    }]
)
```

### Prompt Engineering
```
System: Clinical nutritionist creating evidence-based EMR content
Context: Item title, current question, clinical context
Task: Generate 8 enrichment fields in Portuguese
Output: JSON only (no markdown, no extra text)
Constraints:
  - Medical accuracy
  - Evidence-based recommendations
  - Patient-friendly language
  - Functional medicine approach
```

## Processing Flow

### Step-by-Step Execution

```python
1. Load items (20 score_items)
   └─ Validate IDs exist in database

2. For each item:
   a. Build clinical prompt
   b. Call Claude API
   c. Parse JSON response
   d. Validate enrichment structure
   e. Store in results array
   f. Sleep 2s (rate limiting)

3. Generate outputs:
   a. Write JSON file (all results)
   b. Generate SQL migration (UPDATE statements)
   c. Print execution summary

4. Exit with status code
```

### Error Handling

```python
try:
    # API call
    response = client.messages.create(...)
    enrichment = json.loads(response_text)
    return {"success": True, "data": enrichment}
except anthropic.APIError as e:
    return {"success": False, "error": f"API error: {e}"}
except json.JSONDecodeError as e:
    return {"success": False, "error": f"JSON parse error: {e}"}
except Exception as e:
    return {"success": False, "error": f"Unknown error: {e}"}
```

## SQL Generation

### Template
```sql
UPDATE score_items SET
  question = 'escaped_question',
  clinical_relevance = 'escaped_relevance',
  interpretation_guide = 'escaped_guide',
  health_implications = ARRAY['item1', 'item2', ...]::text[],
  followup_questions = ARRAY['q1', 'q2', ...]::text[],
  red_flags = ARRAY['flag1', 'flag2', ...]::text[],
  recommendations = ARRAY['rec1', 'rec2', ...]::text[],
  scientific_background = 'escaped_background'
WHERE id = 'uuid';
```

### String Escaping
```python
# PostgreSQL single quote escaping
escaped = text.replace("'", "''")

# Array formatting
array_sql = f"ARRAY[{', '.join([f\"'{escape(item)}'\" for item in items])}]::text[]"
```

## Performance Metrics

### Target SLAs
- **API response time**: < 30s per item (p95)
- **Success rate**: > 95%
- **JSON parse success**: 100%
- **SQL syntax validity**: 100%
- **Total execution time**: < 15 minutes

### Monitoring
```python
metrics = {
    "total_items": 20,
    "successful": count_success,
    "failed": count_failed,
    "thinking_tokens": sum_thinking_tokens,
    "output_tokens": sum_output_tokens,
    "total_tokens": sum_all_tokens,
    "execution_time_seconds": elapsed_time,
    "average_time_per_item": avg_time
}
```

## Token Economics

### Per-Item Breakdown
```
Input:
  - System prompt: ~500 tokens
  - Item context: ~200 tokens
  - Instructions: ~300 tokens
  ─────────────────────────
  Subtotal: ~1,000 tokens

Thinking:
  - Analysis: ~3,000 tokens
  - Planning: ~2,000 tokens
  - Refinement: ~1,000 tokens
  ─────────────────────────
  Subtotal: ~6,000 tokens

Output:
  - JSON fields: ~2,000 tokens
  - Arrays: ~500 tokens
  ─────────────────────────
  Subtotal: ~2,500 tokens

TOTAL PER ITEM: ~9,500 tokens
```

### Batch Totals (20 items)
```
Input:    20 × 1,000  =  20,000 tokens × $15/1M = $0.30
Thinking: 20 × 6,000  = 120,000 tokens × $15/1M = $1.80
Output:   20 × 2,500  =  50,000 tokens × $75/1M = $3.75
──────────────────────────────────────────────────────
TOTAL:                  190,000 tokens           $5.85
```

## Quality Assurance

### Automated Checks
```python
def validate_enrichment(enrichment: dict) -> bool:
    required_fields = [
        'question', 'clinical_relevance', 'interpretation_guide',
        'health_implications', 'followup_questions', 'red_flags',
        'recommendations', 'scientific_background'
    ]

    # Check all fields exist
    for field in required_fields:
        if field not in enrichment:
            return False

    # Check array fields
    for array_field in ['health_implications', 'followup_questions', 'red_flags', 'recommendations']:
        if not isinstance(enrichment[array_field], list):
            return False
        if len(enrichment[array_field]) < 3:
            return False

    # Check string lengths
    if len(enrichment['clinical_relevance']) < 100:
        return False
    if len(enrichment['interpretation_guide']) < 200:
        return False

    return True
```

### Manual Review Checklist
- [ ] Clinical accuracy
- [ ] Portuguese grammar and spelling
- [ ] Medical terminology precision
- [ ] Evidence-based recommendations
- [ ] Actionable guidance
- [ ] Patient comprehension level
- [ ] No bias or unsubstantiated claims

## Security Considerations

### API Key Management
```bash
# Environment variable (preferred)
export ANTHROPIC_API_KEY='sk-ant-...'

# Config file (alternative)
echo 'sk-ant-...' > ~/.anthropic_key
chmod 600 ~/.anthropic_key

# Never commit to git
echo '.anthropic_key' >> .gitignore
echo '*.env' >> .gitignore
```

### Data Privacy
- No PHI (Personal Health Information) in prompts
- Only generic clinical contexts
- No patient-specific data
- Anonymized examples only

### SQL Injection Prevention
```python
# Proper escaping for PostgreSQL
def escape_sql_string(text: str) -> str:
    return text.replace("'", "''")

# Using parameterized queries (if available)
cursor.execute(
    "UPDATE score_items SET question = %s WHERE id = %s",
    (question, item_id)
)
```

## Deployment

### Prerequisites
```bash
# Python 3.8+
python3 --version

# Anthropic SDK
pip install anthropic

# PostgreSQL client (for validation)
docker compose exec db psql --version
```

### Execution Environment
```
Platform: Linux/macOS/WSL2
Python: 3.8+
Memory: Minimal (~100MB)
Network: Stable internet (API calls)
Disk: ~10MB for outputs
```

### Rollback Strategy
```sql
-- Backup before applying
CREATE TABLE score_items_backup_20260127 AS
SELECT * FROM score_items
WHERE id IN ('uuid1', 'uuid2', ...);

-- Apply migration
\i batch_alimentacao_parte2.sql

-- Rollback if needed
UPDATE score_items si SET
  question = sib.question,
  clinical_relevance = sib.clinical_relevance,
  -- ... other fields
FROM score_items_backup_20260127 sib
WHERE si.id = sib.id;
```

## Logging and Debugging

### Log Levels
```python
import logging

logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(levelname)s - %(message)s',
    handlers=[
        logging.FileHandler('batch_alimentacao_parte2.log'),
        logging.StreamHandler()
    ]
)
```

### Debug Mode
```python
DEBUG = os.getenv('DEBUG', 'false').lower() == 'true'

if DEBUG:
    print(f"Prompt: {prompt}")
    print(f"Response: {response_text}")
    print(f"Parsed: {json.dumps(enrichment, indent=2)}")
```

## Testing

### Unit Tests
```python
def test_generate_enrichment():
    item = ITEMS[0]
    result = generate_enrichment(item)
    assert result['success'] == True
    assert 'data' in result
    assert validate_enrichment(result['data']) == True

def test_sql_escaping():
    text = "It's a test with 'quotes'"
    escaped = escape_sql_string(text)
    assert escaped == "It''s a test with ''quotes''"
```

### Integration Tests
```bash
# Test with single item
python3 -c "
from batch_alimentacao_parte2 import generate_enrichment, ITEMS
result = generate_enrichment(ITEMS[0])
print(f'Success: {result[\"success\"]}')
"

# Test SQL syntax
docker compose exec db psql -U plenya_user -d plenya_db --dry-run -f batch_alimentacao_parte2.sql
```

## Monitoring and Observability

### Real-time Metrics
```bash
# Watch execution progress
watch -n 1 "tail -n 20 batch_alimentacao_parte2.log"

# Monitor API calls
grep "Processing:" batch_alimentacao_parte2.log | wc -l

# Check success rate
grep "Success" batch_alimentacao_parte2.log | wc -l
```

### Post-execution Analysis
```bash
# Token usage summary
cat batch_alimentacao_parte2_results.json | jq '[.[].tokens | select(.) | .thinking, .output] | add'

# Error analysis
cat batch_alimentacao_parte2_results.json | jq '.[] | select(.error) | {id, title, error}'

# Success rate
cat batch_alimentacao_parte2_results.json | jq '[.[] | select(.enrichment)] | length'
```

## Maintenance

### Regular Updates
- Review prompt effectiveness monthly
- Update scientific references quarterly
- Refine output quality based on clinical feedback
- Optimize token usage if costs increase

### Scalability Considerations
- Current: 20 items per batch
- Maximum: 50 items per batch (limited by time)
- Parallelization: Possible with async/await
- Caching: Not needed (one-time enrichment)

---

**Version**: 1.0
**Date**: 2026-01-27
**Author**: AI Clinical Content System
**Status**: Production-ready
