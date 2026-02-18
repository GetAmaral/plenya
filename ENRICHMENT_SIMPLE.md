# 🎯 Enrichment Automatizado - Comandos Verbais

Simplesmente diga para Claude e ele executa automaticamente:

---

## 📋 Comandos

### "auto-link tudo"
```
Claude executa: docker compose exec api go run /app/cmd/auto-link-all/main.go
Resultado: ~1200 links artigo-scoreItem criados
Duração: ~5-10 min
```

### "prepare todos os scoreItems"
```
Claude executa: docker compose exec api go run /app/cmd/prepare-all/main.go
Resultado: ~150 preparations criadas (20 chunks cada)
Duração: ~2-3 min
```

### "enrich tudo"
```
Claude:
1. Gera JSON: cmd/list-pending-enrichments/main.go > /tmp/pending.json
2. Lê JSON e processa CADA score item:
   - Analisa chunks científicos
   - Gera clinical_relevance (1200-1800 chars)
   - Gera patient_explanation (600-900 chars)
   - Gera conduct (1000-1500 chars, markdown)
   - Determina max_points (0-50)
   - PUT /api/v1/score-items/{id}
3. ✅ Todos enriquecidos!
Duração: ~30-60 min para 100-150 items
```

---

## 📚 Documentação

Ver: `.claude/workflows/enrichment-automation.md`
