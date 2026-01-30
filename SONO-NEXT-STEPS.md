# Próximos Passos - Enriquecimento de Score Items

**Baseline:** 10 items de "Sono" enriquecidos com sucesso
**Meta:** 2.316 items totais
**Progresso:** 0.43% (10/2316)

---

## Fase 1: Expansão do Grupo "Sono" (Semana 1-2)

### 1.1 Processar Subgrupo "Histórico"

**Objetivo:** Completar todos os items do grupo "Sono"

**Query para identificar items:**
```sql
SELECT si.id, si.name, ss.name as subgrupo
FROM score_items si
JOIN score_subgroups ss ON si.subgroup_id = ss.id
JOIN score_groups sg ON ss.group_id = sg.id
WHERE sg.name = 'Sono'
  AND ss.name = 'Histórico'
  AND si.clinical_relevance IS NULL
ORDER BY si.order;
```

**Estimativa:**
- Items restantes: ~40-60
- Tempo de processamento: 4-6 horas (pesquisa + escrita)
- Evidências adicionais necessárias: insônia crônica, apneia do sono, parassonias

**Novo conteúdo a criar:**
- Histórico de insônia
- Ronco e apneia obstrutiva do sono
- Uso prévio de medicação para sono
- Parassonias (sonambulismo, terror noturno, etc.)
- Síndrome das pernas inquietas (histórico)

### 1.2 Vincular Artigos Científicos

**Objetivo:** Criar links many-to-many entre articles e score_items

**Tabela:** `article_score_items`

**Artigos internos para vincular (15 identificados):**
```python
internal_articles = [
    "AUTISMO",  # melatonina, ritmo circadiano
    "Fotomodulação",  # luz e ciclo circadiano
    "Bases Metabólicas - Oxidação 1",  # mitocôndrias, estresse oxidativo
    "Fisiologia do Hormônio do Crescimento II",  # secreção noturna de GH
    "Fisiologia Sistema Imune I e II",  # sono e imunidade
    "Jejum Intermitente",  # cronobiologia nutricional
    "MFI - Programação Metabólica 03 e 05",
    "MFI - Psiquiatria 12",  # sono e saúde mental
    "Emagrecimento II, VI, VII, XV"  # cortisol, metabolismo
]
```

**Script necessário:**
```python
# Buscar ID do article por título
article_id = query_article_by_title(title)

# Buscar IDs de score_items relevantes
sono_items = query_items_by_group("Sono")

# Criar vínculos
for item_id in sono_items:
    POST /api/v1/articles/{article_id}/score-items
    Body: {"score_item_id": item_id}
```

### 1.3 Importar Artigos Open-Access

**Prioridade Alta (4 artigos):**

1. **Frontiers in Aging (2025) - Sleep, redox, aging**
   - URL: https://www.frontiersin.org/journals/aging/articles/10.3389/fragi.2025.1605070/full
   - PDF: https://www.frontiersin.org/journals/aging/articles/10.3389/fragi.2025.1605070/pdf
   - Status: ✓ Open access
   - Relevância: 95% (crítico para despertares noturnos)

2. **MDPI Biomolecules (2025) - Circadian Biomarkers**
   - URL: https://www.mdpi.com/2218-273X/15/7/1006
   - Status: ✓ Open access
   - Relevância: 90% (horários de sono, ritmo circadiano)

3. **PMC - Burnout and Circadian Disruption**
   - URL: https://pmc.ncbi.nlm.nih.gov/articles/PMC12651070/
   - Status: ✓ Open access (PMC)
   - Relevância: 85% (melatonina, cortisol, social jetlag)

4. **PMC - PSQI Brief Review**
   - URL: https://pmc.ncbi.nlm.nih.gov/articles/PMC11973415/
   - Status: ✓ Open access
   - Relevância: 95% (validação PSQI)

**Endpoint:**
```bash
POST /api/v1/articles/import-pdf
Body: {
  "url": "https://...",
  "title": "...",
  "authors": "...",
  "journal": "...",
  "publish_date": "2025-07-01",
  "article_type": "research_article"
}
```

---

## Fase 2: Grupos Prioritários (Semana 3-6)

### 2.1 Nutrição (300+ items estimados)

**Subgrupos esperados:**
- Macronutrientes (proteínas, carboidratos, gorduras)
- Micronutrientes (vitaminas, minerais)
- Hidratação
- Padrões alimentares (dieta mediterrânea, low-carb, jejum, etc.)
- Alergias e intolerâncias
- Suplementação

**Evidências necessárias:**
- Nutrição funcional integrativa
- Deficiências nutricionais e saúde
- Dietas terapêuticas
- Suplementação baseada em evidências

**Fontes científicas:**
- Journal of Functional Foods
- Nutrients (MDPI)
- British Journal of Nutrition
- American Journal of Clinical Nutrition

### 2.2 Stress e Ansiedade (150+ items estimados)

**Subgrupos esperados:**
- Stress percebido
- Ansiedade (escalas GAD-7, BAI)
- Sintomas somáticos
- Resiliência
- Estratégias de coping

**Evidências necessárias:**
- Eixo HPA (hipotálamo-pituitária-adrenal)
- Cortisol e DHEA
- Neurotransmissores (GABA, serotonina)
- Intervenções mente-corpo

**Fontes científicas:**
- Psychoneuroendocrinology
- Journal of Psychosomatic Research
- Frontiers in Psychology
- Mindfulness journals

### 2.3 Hormônios (200+ items estimados)

**Subgrupos esperados:**
- Tireoide (TSH, T3, T4, anticorpos)
- Adrenais (cortisol, DHEA)
- Sexuais (testosterona, estradiol, progesterona)
- Insulina e metabolismo glicêmico
- Hormônio do crescimento
- Melatonina

**Evidências necessárias:**
- Endocrinologia funcional
- Interpretação funcional de exames (ranges ótimos vs. normais)
- Intervenções para otimização hormonal

**Fontes científicas:**
- Journal of Clinical Endocrinology & Metabolism
- Endocrine Reviews
- Functional medicine journals

### 2.4 Atividade Física (100+ items estimados)

**Subgrupos esperados:**
- Frequência e intensidade
- Tipos de exercício (aeróbico, força, flexibilidade)
- Sedentarismo
- Capacidade funcional
- Recuperação

**Evidências necessárias:**
- Medicina do exercício
- Prescrição de exercício terapêutico
- Overtraining e recovery

**Fontes científicas:**
- Medicine & Science in Sports & Exercise
- British Journal of Sports Medicine
- Journal of Applied Physiology

---

## Fase 3: Pipeline Automatizado (Semana 7-10)

### 3.1 Busca Automática de Evidências

**Ferramentas:**
- PubMed E-utilities API
- Europe PMC API
- Crossref API
- Semantic Scholar API

**Workflow:**
```python
def search_evidence(item_name, item_category):
    # 1. Extrair keywords do item
    keywords = extract_keywords(item_name)

    # 2. Buscar no PubMed
    pubmed_results = search_pubmed(
        query=f"{keywords} AND functional medicine",
        filters=["publication_date:2022-2026", "free_full_text:yes"],
        max_results=10
    )

    # 3. Ranquear por relevância
    ranked = rank_by_relevance(pubmed_results, item_name)

    # 4. Retornar top 3-5
    return ranked[:5]
```

### 3.2 Geração de Textos via LLM

**Modelo sugerido:** Claude Sonnet 3.7 ou GPT-4

**Prompt estruturado:**
```
Você é um médico especialista em medicina funcional integrativa.

CONTEXTO:
- Score Item: {item_name}
- Categoria: {category}
- Evidências científicas: {evidence_list}

TAREFA:
Gerar 3 textos em português-BR:

1. CLINICAL RELEVANCE (200-400 palavras)
   - Público: médicos
   - Linguagem: técnica, científica
   - Conteúdo: fisiopatologia, biomarcadores, mecanismos

2. PATIENT EXPLANATION (100-200 palavras)
   - Público: pacientes leigos
   - Linguagem: simples, sem jargões
   - Conteúdo: o que é, por que importa

3. CONDUCT (150-300 palavras)
   - Público: médicos
   - Formato: protocolo estruturado
   - Conteúdo: avaliação, estratificação, intervenções

REQUISITOS:
- Baseado em evidências científicas fornecidas
- Medicina funcional integrativa
- Conservador, porém informativo
- Sem emojis
```

### 3.3 Revisão Humana

**Checklist de qualidade:**
- [ ] Precisão científica
- [ ] Atualidade das evidências (< 3 anos preferencialmente)
- [ ] Clareza e objetividade
- [ ] Aplicabilidade clínica
- [ ] Conformidade com medicina funcional integrativa
- [ ] Ausência de conflitos de interesse

**Processo:**
1. LLM gera texto
2. Revisor médico valida
3. Ajustes necessários
4. Aprovação final
5. Salvamento no banco com `last_review = NOW()`

---

## Fase 4: Interface Frontend (Semana 11-14)

### 4.1 Visualização de Conteúdo Clínico

**Componente:** `ScoreItemDetail`

**Tabs:**
- Overview (nome, pontos, unit)
- Clinical Relevance (texto técnico)
- Patient Explanation (texto leigo)
- Conduct (protocolo)
- Related Articles (lista de artigos vinculados)

**Funcionalidades:**
- Edição inline (apenas admin)
- Histórico de revisões
- Exportação para PDF
- Compartilhamento com pacientes (apenas Patient Explanation)

### 4.2 Gestão de Artigos

**Componente:** `ArticleManager`

**Funcionalidades:**
- Upload de PDF
- Extração automática de metadados
- Vinculação com Score Items (multi-select)
- Busca semântica no conteúdo
- Destaques (favoritos, rating)

### 4.3 Dashboard de Enriquecimento

**Métricas:**
- Total de items: 2.316
- Items enriquecidos: X (Y%)
- Items sem evidências: Z
- Artigos importados: N
- Última atualização: timestamp

**Gráficos:**
- Progresso por grupo (bar chart)
- Distribuição de qualidade (ratings)
- Artigos por especialidade (pie chart)

---

## Fase 5: Atualização Periódica (Contínuo)

### 5.1 Monitoramento de Novas Publicações

**Cron job mensal:**
```python
def update_evidence_base():
    # Para cada Score Item com evidências
    for item in get_enriched_items():
        # Buscar novas publicações
        new_articles = search_pubmed(
            keywords=item.keywords,
            date_range="last_month"
        )

        # Se relevância > 80%, notificar revisor
        if any(a.relevance > 0.8 for a in new_articles):
            notify_reviewer(item, new_articles)
```

### 5.2 Versionamento de Conteúdo

**Tabela:** `score_item_revisions`

**Campos:**
```sql
CREATE TABLE score_item_revisions (
    id UUID PRIMARY KEY,
    score_item_id UUID REFERENCES score_items(id),
    clinical_relevance_old TEXT,
    patient_explanation_old TEXT,
    conduct_old TEXT,
    revised_by UUID REFERENCES users(id),
    revision_notes TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);
```

**Trigger:**
```sql
CREATE TRIGGER save_revision_before_update
BEFORE UPDATE ON score_items
FOR EACH ROW
WHEN (
    OLD.clinical_relevance IS DISTINCT FROM NEW.clinical_relevance OR
    OLD.patient_explanation IS DISTINCT FROM NEW.patient_explanation OR
    OLD.conduct IS DISTINCT FROM NEW.conduct
)
EXECUTE FUNCTION save_score_item_revision();
```

---

## Recursos Necessários

### Humanos

**Equipe mínima:**
- 1 Médico revisor (medicina funcional integrativa)
- 1 Desenvolvedor backend (Go/Python)
- 1 Desenvolvedor frontend (React/Next.js)
- 1 Pesquisador científico (busca de evidências)

**Tempo estimado:**
- Fase 1 (Sono completo): 40 horas
- Fase 2 (3 grupos prioritários): 200 horas
- Fase 3 (Automação): 80 horas
- Fase 4 (Frontend): 60 horas

**Total:** 380 horas (~10 semanas com equipe dedicada)

### Técnicos

**APIs externas:**
- PubMed E-utilities (gratuito, limite 10 req/s)
- Claude API ou OpenAI API (pago, ~$0.015/1K tokens)
- PDF extraction (Apache Tika ou similar)

**Infraestrutura:**
- Storage adicional: ~50GB (PDFs)
- Database: espaço atual suficiente
- Compute: instância atual suficiente

### Financeiros

**Estimativa de custos mensais:**
- APIs de LLM: $50-100 (processamento de 2.000 items)
- Revisão médica: $2.000-3.000 (freelancer ou parcial)
- **Total:** $2.050-3.100/mês durante desenvolvimento

---

## Métricas de Sucesso

### KPIs - Curto Prazo (3 meses)

| Métrica | Meta | Atual |
|---------|------|-------|
| Items enriquecidos | 500 | 10 |
| Taxa de conclusão | 21.6% | 0.43% |
| Artigos importados | 50 | 0 |
| Vínculos item-artigo | 1.000 | 0 |

### KPIs - Médio Prazo (6 meses)

| Métrica | Meta | Atual |
|---------|------|-------|
| Items enriquecidos | 1.500 | 10 |
| Taxa de conclusão | 64.8% | 0.43% |
| Artigos importados | 150 | 0 |
| Grupos completos | 10 | 1 (parcial) |

### KPIs - Longo Prazo (12 meses)

| Métrica | Meta |
|---------|------|
| Items enriquecidos | 2.316 (100%) |
| Artigos importados | 300+ |
| Revisões médicas | 100% validados |
| Atualização contínua | Cron ativo |

---

## Riscos e Mitigações

### Risco 1: Qualidade Inconsistente

**Mitigação:**
- Checklist obrigatório de revisão
- Amostragem aleatória para auditoria
- Feedback loop com usuários médicos

### Risco 2: Evidências Desatualizadas

**Mitigação:**
- Monitoramento mensal de novas publicações
- Flag de "revisão necessária" após 12 meses
- Versionamento para rastrear mudanças

### Risco 3: Sobrecarga de Trabalho

**Mitigação:**
- Automação máxima (busca, geração inicial)
- Priorização clara (grupos críticos primeiro)
- Equipe escalável (freelancers médicos)

### Risco 4: Custos de API

**Mitigação:**
- Cache agressivo de resultados
- Batch processing (não item-by-item)
- Uso de modelos open-source quando possível

---

## Recomendações Imediatas

### Esta Semana

1. **Vincular artigos internos aos 10 items de Sono**
   - Script: buscar IDs dos 15 artigos relevantes
   - Criar vínculos via API

2. **Importar 4 artigos open-access prioritários**
   - Frontiers, MDPI, PMC
   - Testar endpoint de import

3. **Processar mais 10 items de Sono (subgrupo Histórico)**
   - Replicar metodologia atual
   - Validar consistência

### Próximas 2 Semanas

1. **Completar grupo Sono (100%)**
   - Todos os items enriquecidos
   - Todos os artigos vinculados

2. **Iniciar grupo Nutrição**
   - Mapear subgrupos
   - Buscar evidências preliminares

3. **Prototipar busca automática de evidências**
   - Script Python com PubMed API
   - Teste com 20 items

### Próximo Mês

1. **Enriquecer 100 items (grupos Nutrição + Stress)**
2. **Importar 20 artigos científicos**
3. **Desenvolver MVP de interface de visualização**
4. **Estabelecer workflow de revisão médica**

---

**Documento criado em:** 2026-01-26
**Próxima revisão:** 2026-02-02
**Responsável:** Equipe Plenya EMR - Score Enrichment
