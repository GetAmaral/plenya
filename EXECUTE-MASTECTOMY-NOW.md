# Execute Mastectomy Enrichment - Quick Start

**⚡ READY TO EXECUTE ⚡**

---

## One-Command Execution

```bash
cd /home/user/plenya && ./scripts/execute_mastectomy_enrichment.sh
```

**OR**

```bash
docker compose exec -T db psql -U plenya_user -d plenya_db < /home/user/plenya/scripts/enrich_mastectomy_item.sql
```

---

## What Will Happen

1. **4 Scientific Articles** inserted into `articles` table
2. **4 Relationships** created in `article_score_items` table
3. **3 Clinical Fields** updated in `score_items` table:
   - `clinical_relevance` (1,923 characters)
   - `patient_explanation` (1,487 characters)
   - `conduct` (2,490 characters)
4. **Timestamp** set in `last_review` field

---

## Expected Output

```
BEGIN
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
UPDATE 1
COMMIT

 id                                   | item_text   | has_clinical_relevance | has_patient_explanation | has_conduct | last_review          | article_count
--------------------------------------+-------------+------------------------+-------------------------+-------------+----------------------+--------------
 e65c56dc-5c07-4270-8a8b-017b293ca147 | Mastectomia | t                      | t                       | t           | 2026-01-28 21:15:42  | 4
```

---

## Verify Success

```bash
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT
    item_text,
    LENGTH(clinical_relevance) as relevance,
    LENGTH(patient_explanation) as explanation,
    LENGTH(conduct) as conduct,
    last_review
FROM score_items
WHERE id = 'e65c56dc-5c07-4270-8a8b-017b293ca147';
"
```

**Expected:**
- relevance: 1923
- explanation: 1487
- conduct: 2490
- last_review: Today's date

---

## View Articles

```bash
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT
    a.title,
    a.journal,
    a.year,
    a.pmid
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = 'e65c56dc-5c07-4270-8a8b-017b293ca147';
"
```

---

## Documentation

- **Full Report:** `/home/user/plenya/MASTECTOMY-ENRICHMENT-REPORT.md`
- **Quick Summary:** `/home/user/plenya/MASTECTOMY-ENRICHMENT-SUMMARY.md`
- **SQL Script:** `/home/user/plenya/scripts/enrich_mastectomy_item.sql`

---

## Key Evidence Highlights

### Physical Health Impact
- 90% of survivors have lasting complications
- Chemotherapy: Decline >5 years
- Endocrine therapy: Decline <2 years

### Metabolic Syndrome
- HR 1.69 for recurrence
- HR 1.83 for mortality
- HR 1.29 for cardiovascular events

### Clinical Approach
- Comprehensive metabolic screening
- Multidisciplinary care coordination
- Evidence-based exercise programs
- Anti-inflammatory nutrition
- Psychological support integration

---

## ⚡ Execute Now

```bash
./scripts/execute_mastectomy_enrichment.sh
```

---

**Status:** ✅ READY
**Execution Time:** <10 seconds
**Safety:** Transaction-wrapped (rollback on error)
