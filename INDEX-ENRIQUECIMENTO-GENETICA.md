# Índice de Arquivos - Enriquecimento Genético Batch 1

**Projeto:** Plenya EMR - Enriquecimento de Score Items Genética
**Data:** 2026-01-26
**Status:** 9 de 20 genes processados (45%)

---

## Arquivos de Dados (JSON)

### Para Importação

**1. genetic_enrichment_batch1_CONSOLIDATED.json** ⭐ **USAR ESTE**
- Arquivo consolidado com 9 genes completos
- Formato: Array JSON
- Campos: gene_id, gene_name, clinical_relevance, patient_explanation, conduct, sources
- Encoding: UTF-8
- Pronto para importação via API

### Arquivos Intermediários (Referência)

**2. genetic_enrichment_batch1.json**
- Genes 1-6 (FTO, MC4R, TCF7L2, PPARG, MTHFR, APOE)
- Consolidado em CONSOLIDATED.json

**3. genetic_enrichment_batch1_part2.json**
- Genes 7-9 (FADS1, FADS2, ACE)
- Consolidado em CONSOLIDATED.json

---

## Documentação Técnica (Markdown)

### Documentação Principal

**4. SUMARIO-EXECUTIVO-GENETICA-BATCH1.md** ⭐ **LEITURA RECOMENDADA**
- Visão executiva do projeto
- Destaques nutrigenômicos
- Próximos passos e ROI
- Ideal para apresentação a stakeholders

**5. GENETIC-ENRICHMENT-FINAL-REPORT.md** ⭐ **RELATÓRIO COMPLETO**
- Relatório técnico detalhado
- Metodologia de pesquisa
- Métricas de qualidade
- Princípios de nutrigenômica aplicados
- 5.800 palavras

**6. IMPORTACAO-GENES-INSTRUCOES.md** ⭐ **GUIA DE IMPORTAÇÃO**
- Instruções passo-a-passo
- 3 métodos de importação:
  - Script Python (recomendado)
  - curl manual
  - Interface web
- Troubleshooting
- Validação

### Documentação Intermediária

**7. GENETIC-ENRICHMENT-BATCH1-SUMMARY.md**
- Sumário após processamento de 6 genes
- Estado intermediário do projeto
- Referência histórica

---

## Outros Arquivos Relacionados

**8. PAINEIS_GENETICOS_BRASIL.md**
- Informações sobre painéis genéticos disponíveis no Brasil
- Não diretamente relacionado ao enriquecimento

**9. PLANO_IMPORTACAO_CSV_ESCORE.md**
- Plano de importação do CSV original de scores
- Contexto: como genes foram importados inicialmente

---

## Estrutura de Diretórios

```
/home/user/plenya/
├── genetic_enrichment_batch1_CONSOLIDATED.json  ← IMPORTAR
├── genetic_enrichment_batch1.json               (referência)
├── genetic_enrichment_batch1_part2.json         (referência)
├── SUMARIO-EXECUTIVO-GENETICA-BATCH1.md         ← LER PRIMEIRO
├── GENETIC-ENRICHMENT-FINAL-REPORT.md           ← TÉCNICO COMPLETO
├── IMPORTACAO-GENES-INSTRUCOES.md               ← GUIA IMPORTAÇÃO
├── GENETIC-ENRICHMENT-BATCH1-SUMMARY.md         (intermediário)
└── INDEX-ENRIQUECIMENTO-GENETICA.md             (este arquivo)
```

---

## Workflow Recomendado

### 1. Leitura Inicial

```bash
# Visão executiva (5 min)
less /home/user/plenya/SUMARIO-EXECUTIVO-GENETICA-BATCH1.md

# Relatório técnico completo (15 min)
less /home/user/plenya/GENETIC-ENRICHMENT-FINAL-REPORT.md
```

### 2. Preparação para Importação

```bash
# Guia de importação (10 min)
less /home/user/plenya/IMPORTACAO-GENES-INSTRUCOES.md

# Verificar arquivo de dados
jq '.[0]' /home/user/plenya/genetic_enrichment_batch1_CONSOLIDATED.json
```

### 3. Importação

```bash
# Método Python (recomendado)
cd /home/user/plenya
pip3 install requests
# Criar script conforme IMPORTACAO-GENES-INSTRUCOES.md
python3 import_genetic_enrichment.py
```

### 4. Validação

```bash
# Verificar no banco
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
  SELECT name, LENGTH(clinical_relevance) as chars
  FROM score_groups
  WHERE clinical_relevance IS NOT NULL
  LIMIT 5;
"
```

---

## Resumo dos 9 Genes Processados

| # | Gene | Categoria | Intervenção Chave |
|---|------|-----------|-------------------|
| 1 | FTO rs9939609 | Obesidade | Exercício anula risco AA |
| 2 | MC4R rs17782313 | Obesidade | Dieta volumétrica + mindful eating |
| 3 | TCF7L2 rs7903146 | Diabetes | Agonistas GLP-1 funcionam melhor em TT |
| 4 | PPARG Pro12Ala | Diabetes | Vitamina E aumenta adiponectina |
| 5 | MTHFR C677T | Homocisteína | Metilfolato normaliza (teste NÃO recomendado) |
| 6 | APOE | Alzheimer | Colina + DHA revertem danos em ε4 |
| 7 | FADS1 rs174546 | Ômega-3 | TT vegano PRECISA suplementar |
| 8 | FADS2 rs174575 | Ômega-3 | Sinérgico com FADS1 |
| 9 | ACE I/D | Hipertensão | DD responde melhor a IECA, II é sal-sensível |

---

## Métricas de Qualidade

- **Total de palavras geradas:** ~10.500 palavras (textos genéticos)
- **Média por gene:** 1.166 palavras
- **Fontes científicas:** 15+ artigos peer-reviewed
- **Artigos recentes (2024-2025):** 53%
- **Genes com nutrigenômica:** 100%
- **Genes com farmacogenômica:** 44%

---

## Próximos Arquivos a Gerar

### Batch 1 (Genes 10-20)

Quando completados, serão gerados:
- `genetic_enrichment_batch1_COMPLETE.json` (20 genes)
- `GENETIC-ENRICHMENT-BATCH1-COMPLETE-REPORT.md`

### Batches Futuros

- `genetic_enrichment_batch2.json` (genes 21-40)
- `genetic_enrichment_batch3.json` (genes 41-60)
- `genetic_enrichment_batch4.json` (genes 61-80)

---

## Comandos Úteis

### Visualizar Dados

```bash
# Ver todos genes no JSON consolidado
jq '.[] | .gene_name' genetic_enrichment_batch1_CONSOLIDATED.json

# Ver detalhes de um gene específico (ex: APOE)
jq '.[] | select(.gene_name | contains("APOE"))' genetic_enrichment_batch1_CONSOLIDATED.json

# Contar palavras por campo
jq '.[] | {
  name: .gene_name,
  clinical: (.clinical_relevance | split(" ") | length),
  patient: (.patient_explanation | split(" ") | length),
  conduct: (.conduct | split(" ") | length)
}' genetic_enrichment_batch1_CONSOLIDATED.json
```

### Buscar Informações

```bash
# Buscar menção a um nutriente (ex: "colina")
grep -i "colina" genetic_enrichment_batch1_CONSOLIDATED.json

# Buscar genes relacionados a diabetes
jq '.[] | select(.gene_name | contains("Diabetes"))' genetic_enrichment_batch1_CONSOLIDATED.json
```

---

## Fontes Científicas Utilizadas

### Journals
- Genes & Nutrition (2024)
- MDPI Biomolecules & Nutrients (2023-2024)
- Nature Scientific Reports (2020-2024)
- Springer Nature journals (2022-2024)
- Frontiers in Nutrition (2023-2024)
- American Heart Association journals

### Guidelines
- CDC - MTHFR guidelines
- American College of Medical Genetics (ACMG)
- RACGP (Royal Australian College of General Practitioners)

### Institutions
- MIT Picower Institute (2024)
- PubMed/PubMed Central

---

## Contato e Suporte

Para dúvidas sobre:
- **Importação:** Ver `IMPORTACAO-GENES-INSTRUCOES.md`
- **Conteúdo técnico:** Ver `GENETIC-ENRICHMENT-FINAL-REPORT.md`
- **Visão geral:** Ver `SUMARIO-EXECUTIVO-GENETICA-BATCH1.md`
- **Sistema Plenya:** Ver `/home/user/plenya/CLAUDE.md`

---

**Última atualização:** 2026-01-26
**Processador:** Claude Sonnet 4.5
**Status:** Batch 1 - 45% completo (9/20 genes)
**Próximo milestone:** Completar genes 10-20
