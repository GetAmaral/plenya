# SHBG - Mulheres: Enriquecimento Concluído

**Status:** ✓ COMPLETO
**Data:** 2026-01-29
**ID:** c21ccec2-66b2-49e3-911a-8d0944eda087

---

## Métricas Finais

| Campo | Caracteres | Status |
|-------|-----------|--------|
| clinical_relevance | 1.561 | ✓ OK (1500-2000) |
| patient_explanation | 1.289 | ✓ OK (1000-1500) |
| conduct | 1.826 | ✓ OK (1500-2500) |

## Artigos Científicos (3)

1. **Sex hormone binding globulin as a potential drug candidate for liver-related metabolic disorders treatment**
   - Biomedicine & Pharmacotherapy, 2022
   - PMID: 35738176 | DOI: 10.1016/j.biopha.2022.113261

2. **Contraception for Women with Polycystic Ovary Syndrome: Dealing with a Complex Condition**
   - Revista Brasileira de Ginecologia e Obstetrícia, 2022
   - PMID: 35623618 | DOI: 10.1055/s-0042-1748036

3. **Sex Hormone-Binding Globulin (SHBG) as an Early Biomarker and Therapeutic Target in Polycystic Ovary Syndrome**
   - International Journal of Molecular Sciences, 2020
   - PMID: 33139661 | DOI: 10.3390/ijms21218191

## Principais Tópicos Abordados

- SHBG como marcador de resistência insulínica e síndrome metabólica
- Valores de referência: <40 nmol/L (baixo), 40-114 nmol/L (normal), >114 nmol/L (elevado)
- Índice de Andrógenos Livres (FAI): (Testosterona Total/SHBG) × 100
- Associação com SOP e hiperandrogenismo
- Efeito de contraceptivos hormonais (aumentam SHBG)
- Protocolo de investigação e manejo clínico completo

## Arquivos Gerados

- `/home/user/plenya/scripts/enrich_shbg_mulheres_v2.sql` (executado)
- `/home/user/plenya/scripts/fix_shbg_character_count.sql` (executado)
- `/home/user/plenya/scripts/verify_shbg_enrichment.sql`
- `/home/user/plenya/SHBG-MULHERES-ENRICHMENT-REPORT.md`

## Comando de Verificação

```bash
cat /home/user/plenya/scripts/verify_shbg_enrichment.sql | docker compose -f /home/user/plenya/docker-compose.yml exec -T db psql -U plenya_user -d plenya_db
```
