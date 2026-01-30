# Relatório de Enriquecimento: Metano Basal (CH₄ Jejum)

**Data de Execução:** 2026-01-28
**Item ID:** `a92d20ce-702f-4b15-8817-098b9539c0f0`
**Grupo:** Exames > Imagem
**Status:** ✅ CONCLUÍDO

---

## 1. Resumo Executivo

Foi realizado o enriquecimento completo do item "Metano Basal (CH₄ Jejum)" no sistema Plenya, incluindo:

- 3 artigos científicos de alta qualidade salvos no banco de dados
- Conteúdo clínico completo em português brasileiro
- Relações many-to-many estabelecidas entre artigos e score_item
- Validação da integridade dos dados

---

## 2. Contexto Clínico

### O que é Metano Basal?

O metano basal é medido através do teste respiratório em jejum e representa um biomarcador fundamental para diagnóstico de **Supercrescimento de Metanógenos Intestinais (IMO - Intestinal Methanogen Overgrowth)**.

### Características Principais

- **Threshold diagnóstico:** ≥10 ppm define IMO positivo
- **Organismos responsáveis:** Arqueias (não bactérias), principalmente *Methanobrevibacter smithii*
- **Fenótipo clínico:** Fortemente associado à constipação crônica
- **Mecanismo:** Inibição direta do trânsito intestinal (até 59% em modelos animais)
- **Especificidade:** Pacientes com IMO têm 5x mais probabilidade de constipação vs. SIBO clássico

### Metodologia Diagnóstica

- **Teste simplificado:** Medição única em jejum (SMM - Single Methane Measurement)
- **Performance:** Sensibilidade 86,4%, Especificidade 100%
- **Vantagem:** Comparável aos testes de 2 horas, mais prático e econômico

---

## 3. Artigos Científicos Salvos

### 3.1 North American Consensus Guidelines (2017)

**Título:** Hydrogen and Methane-Based Breath Testing in Gastrointestinal Disorders: The North American Consensus

**Autores:** Rezaie A, Buresi M, Lembo A, Lin H, McCallum R, Rao S, Schmulson M, Valdovinos M, Zakko S, Pimentel M

**Journal:** American Journal of Gastroenterology

**DOI:** 10.1038/ajg.2017.46

**PMID:** 28323273

**Relevância:** Consenso oficial que estabelece o threshold de ≥10 ppm para metano, define a correlação com constipação (5x mais provável), e orienta seleção de antibióticos específicos para metanógenos.

---

### 3.2 Understanding Hydrogen-Methane Breath Testing (2023)

**Título:** Understanding Our Tests: Hydrogen-Methane Breath Testing to Diagnose Small Intestinal Bacterial Overgrowth

**Autores:** Tansel A, Levinthal DJ

**Journal:** Clinical and Translational Gastroenterology

**DOI:** 10.14309/ctg.0000000000000567

**PMID:** 36744854

**Relevância:** Revisão detalhada introduzindo o conceito de IMO, validando a medição única em jejum (SMM ≥10 ppm) com alta performance diagnóstica, e explicando o mecanismo fisiopatológico da inibição do trânsito intestinal.

---

### 3.3 Meta-Analysis: Methane and Constipation (2011)

**Título:** Methane on breath testing is associated with constipation: a systematic review and meta-analysis

**Autores:** Kunkel D, Basseri RJ, Makhani MD, Chong K, Chang C, Pimentel M

**Journal:** Digestive Diseases and Sciences

**DOI:** 10.1007/s10620-010-1455-4

**PMID:** 21286935

**Relevância:** Meta-análise demonstrando OR de 3,51 (IC 95%: 2,00-6,16) para associação metano-constipação. Evidencia correlação entre níveis de *M. smithii* fecal, metano expirado e constipação.

---

## 4. Conteúdo Clínico Criado (PT-BR)

### 4.1 Clinical Relevance (919 caracteres)

Explica que o metano basal é biomarcador para IMO, define o threshold ≥10 ppm, descreve que metanógenos são arqueias (não bactérias), detalha a inibição do trânsito intestinal (59%), apresenta o risco 5x maior de constipação vs. SIBO clássico, e menciona a performance diagnóstica (sensibilidade 86,4%, especificidade 100%) da medição em jejum.

### 4.2 Patient Explanation (868 caracteres)

Linguagem acessível explicando que o metano é medido após jejum de 8-12 horas, é produzido por metanógenos intestinais, e causa lentificação intestinal resultando em constipação. Descreve sintomas associados (distensão, gases, desconforto) e menciona tratamentos específicos com antibióticos e mudanças alimentares.

### 4.3 Conduct (1452 caracteres)

Protocolo terapêutico detalhado:

- **Primeira linha:** Rifaximina 550mg 3x/dia + Neomicina 500mg 2x/dia por 14 dias
- **Alternativa:** Metronidazol 500mg 3x/dia por 14 dias
- **Suporte nutricional:** Dieta FODMAP baixa por 4-6 semanas
- **Procinéticos:** Prucaloprida 2mg/dia ou linaclotida 290mcg/dia
- **Probióticos:** Evitar Bifidobacterium; preferir S. boulardii ou L. plantarum
- **Reavaliação:** Repetir teste após 4 semanas
- **Investigação de causas:** Dismotilidade, IBP crônico, divertículos, estenoses
- **Taxa de recorrência:** 40-50%, pode necessitar ciclos adicionais

---

## 5. Validação dos Dados

### Verificação Final

```
Score Item ID: a92d20ce-702f-4b15-8817-098b9539c0f0
Nome: Metano Basal (CH₄ Jejum)
Clinical Relevance: 919 caracteres ✓
Patient Explanation: 868 caracteres ✓
Conduct: 1452 caracteres ✓
Last Review: 2026-01-28 16:28:40
Artigos Vinculados: 12 total (3 científicos com DOI + 9 lectures MFI)
```

### Artigos Científicos com DOI

1. Understanding Our Tests... (2023) - DOI: 10.14309/ctg.0000000000000567
2. North American Consensus (2017) - DOI: 10.1038/ajg.2017.46
3. Meta-analysis (2011) - DOI: 10.1007/s10620-010-1455-4

---

## 6. Arquivos Gerados

- **Script SQL:** `/home/user/plenya/scripts/enrich_metano_basal.sql`
- **Relatório:** `/home/user/plenya/METANO-BASAL-ENRICHMENT-REPORT.md`

---

## 7. Referências Utilizadas

### Fontes Primárias

- [Understanding Our Tests: Hydrogen-Methane Breath Testing](https://pmc.ncbi.nlm.nih.gov/articles/PMC10132719/)
- [North American Consensus Guidelines](https://pmc.ncbi.nlm.nih.gov/articles/PMC5418558/)
- [Methane and Constipation Meta-analysis](https://pubmed.ncbi.nlm.nih.gov/21286935/)

### Buscas Realizadas

- "methane breath test IMO intestinal methanogen overgrowth 2024 2025 2026"
- "methane breath test constipation clinical guidelines PubMed"

---

## 8. Próximos Passos Sugeridos

1. Integrar o score_item ao sistema de laudos automatizados
2. Criar algoritmo de interpretação automática (≥10 ppm = IMO positivo)
3. Vincular protocolo terapêutico a prescrições eletrônicas
4. Adicionar alertas para investigação de causas subjacentes (dismotilidade, IBP)
5. Implementar follow-up automático após 4 semanas para reteste

---

**Executado via Docker:** ✓
**Transação completa:** COMMIT successful
**Integridade referencial:** Verificada (foreign keys válidas)

---

**Status Final:** 🎯 MISSÃO CUMPRIDA
