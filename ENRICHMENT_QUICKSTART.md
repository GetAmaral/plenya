# Score Item Enrichment - Quick Start

**One-page reference for the ZERO-API enrichment workflow**

---

## 🚀 Commands

### Report (Check Status)
```bash
./scripts/enrich_score.sh report
```

### Review (Dry-Run)
```bash
./scripts/enrich_score.sh review 0
```

### Enrich Single Item
```bash
./scripts/enrich_score.sh enrich 0
```

### Batch Process
```bash
./scripts/enrich_score.sh batch 0 9    # Items 0-9
./scripts/enrich_score.sh batch 10 19  # Items 10-19
```

---

## 📊 Current Status

```
Total:     878 items
Reviewed:  439 (50%)
Pending:   439 (50%)

Tiers:
- ENRICH:   298 items (improve existing)
- REWRITE:  141 items (from scratch)
```

---

## 💰 Costs

- **Per item:** ~$0.05
- **Full run (439 items):** ~$22
- **Time:** ~73 minutes (with delays)

---

## 🎯 Workflow

```
1. Generate report    → Understand scope
2. Review item        → Preview without changes
3. Enrich item        → Full LLM enrichment
4. Repeat             → Process next offset
```

---

## 📁 Files

| File | Purpose |
|------|---------|
| `scripts/enrich_score.sh` | Shell wrapper (main interface) |
| `apps/api/scripts/enrich_by_offset.go` | Main enrichment logic |
| `apps/api/scripts/report_pending_enrichment.go` | Report generator |
| `apps/api/scripts/README_ENRICHMENT.md` | Full documentation |
| `ENRICHMENT_WORKFLOW_SUMMARY.md` | Implementation summary |

---

## 🔧 Setup

### Required Environment Variable
```bash
# In docker-compose.yml or .env
ANTHROPIC_API_KEY=sk-ant-...
```

### Make Scripts Executable
```bash
chmod +x scripts/enrich_score.sh
```

---

## 💡 Tips

- **Always start with report** to understand scope
- **Use dry-run** before enriching new offsets
- **Process in batches** of 10-20 items
- **Monitor costs** with report between batches
- **Check audit trail** in `score_item_review_history` table

---

## ⚠️ Important

- This is **development only** (direct DB access)
- Production apps **MUST use HTTP API**
- `last_review` is **automatically set** via GORM hook
- All enrichments are **logged** for audit

---

## 🐛 Troubleshooting

### "API key not found"
```bash
docker compose exec api env | grep ANTHROPIC
# Add to docker-compose.yml if missing
```

### "No more items at offset"
```bash
./scripts/enrich_score.sh report  # Check remaining
```

### "Validation warnings"
- Soft warnings, enrichment still saved
- Review output, may need manual adjustment

---

## 📚 Full Documentation

- `apps/api/scripts/README_ENRICHMENT.md`
- `ENRICHMENT_WORKFLOW_SUMMARY.md`
- `.claude/workflows/database-ops.md`

---

**Ready to start? Run:**
```bash
./scripts/enrich_score.sh report
```
