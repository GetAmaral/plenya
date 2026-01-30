# Quick Start: Enriquecer Hemácias - Mulheres

## TL;DR - Execução em 3 Comandos

```bash
# 1. Definir API key
export ANTHROPIC_API_KEY="sk-ant-api03-..."

# 2. Executar script
./run-enrich-hemacias.sh

# 3. Verificar resultado
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT name, LENGTH(clinical_relevance), LENGTH(patient_explanation), LENGTH(conduct)
   FROM score_items WHERE id='501fd84a-a440-4c13-9b11-35e2f69017d1';"
```

## O Que Será Feito

```
[1] Buscar 3-4 artigos sobre hemácias em mulheres (PubMed/Scholar)
    ↓
[2] Gerar conteúdo clínico em PT-BR (clinical/patient/conduct)
    ↓
[3] Salvar artigos no banco de dados (table: articles)
    ↓
[4] Vincular artigos ao item (table: score_item_articles)
    ↓
[5] Atualizar score_item com campos enriquecidos
```

## Pré-requisitos Checklist

- [ ] Docker rodando (`docker ps`)
- [ ] Containers ativos (`docker compose ps`)
- [ ] Anthropic API key válida
- [ ] Dependências instaladas (`@anthropic-ai/sdk`, `pg`)

## Verificação Rápida

```bash
# Containers rodando?
docker compose ps

# Database acessível?
docker compose exec db pg_isready -U plenya_user

# Item existe?
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT id, name FROM score_items WHERE id='501fd84a-a440-4c13-9b11-35e2f69017d1';"
```

## Resultado Esperado

```
✓ Artigos salvos: 3-4
✓ clinical_relevance: ~800-1200 chars
✓ patient_explanation: ~600-900 chars
✓ conduct: ~500-700 chars
✓ Tempo: 2-3 minutos
✓ Custo: ~$0.05-0.10
```

## Se Algo Der Errado

| Erro | Solução |
|------|---------|
| API key inválida | `export ANTHROPIC_API_KEY="sk-ant-..."` |
| Container parado | `docker compose up -d` |
| Dep. faltando | `docker compose exec web pnpm add @anthropic-ai/sdk pg --filter web` |
| DB inacessível | `docker compose restart db && sleep 10` |

## Arquivos Criados

- `/home/user/plenya/scripts/enrich_hemacias_mulheres.mjs` ← Script principal
- `/home/user/plenya/run-enrich-hemacias.sh` ← Wrapper executável
- `/home/user/plenya/HEMACIAS-MULHERES-SUMMARY.md` ← Documentação completa
- Este arquivo ← Quick start

## Mais Informações

- **Documentação completa:** `HEMACIAS-MULHERES-SUMMARY.md`
- **Instruções detalhadas:** `EXECUTE-HEMACIAS-MULHERES.md`
- **Item ID:** `501fd84a-a440-4c13-9b11-35e2f69017d1`
- **Valores ref.:** 4.0-5.2 milhões/µL

---

**Pronto para executar!** 🚀
