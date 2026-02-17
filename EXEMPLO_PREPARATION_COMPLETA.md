# Exemplo de ScoreItemEnrichmentPreparation Completa

## Objeto Final Esperado

```json
{
  "id": "019c6abc-1234-7890-abcd-ef1234567890",
  "score_item_id": "019bf31d-2ef0-71b6-a5d1-db49a4fa62fa",

  "selected_chunks": {
    "items": [
      {
        "article_id": "019c63d8-7154-745e-91b2-4a3987e83f94",
        "article_title": "Elevated Serum Leptin Levels as a Predictive Marker for Polycystic Ovary Syndrome",
        "article_year": 2022,
        "journal": "J Endocrinol Invest",
        "chunk_index": 0,
        "chunk_text": "Estudo demonstrando que níveis elevados de leptina sérica estão significativamente associados à síndrome dos ovários policísticos (SOP). Níveis de leptina >11.58 ng/mL apresentam sensibilidade de 77.5% e especificidade de 62.6% para predição de SOP, especialmente em mulheres com hiperandrogenismo e sobrepeso/obesidade. A leptina, hormônio derivado do tecido adiposo, desempenha papel crucial na regulação da homeostase energética e reprodução. Em mulheres com SOP, hiperleptinemia correlaciona-se com resistência insulínica, obesidade central e hiperandrogenemia, formando um ciclo metabólico-reprodutivo adverso.",
        "section": "abstract",
        "similarity": 0.730,
        "word_count": 94
      },
      {
        "article_id": "019c63d8-7157-7655-a885-a9c4b417ba61",
        "article_title": "The association between leptin, adiponectin levels and the ovarian reserve in women of reproductive age",
        "article_year": 2024,
        "journal": "Reprod Health",
        "chunk_index": 0,
        "chunk_text": "Estudo investigando a relação entre adipocinas (leptina e adiponectina) e reserva ovariana em mulheres em idade reprodutiva. Demonstra que mulheres com SOP apresentam níveis mais elevados de leptina (média 18.3±6.2 ng/mL vs 8.1±3.4 ng/mL em controles, p<0.001) e menores de adiponectina, sugerindo que alterações hormonais e metabólicas podem influenciar negativamente a fertilidade. A relação leptina/adiponectina >2.0 foi identificada como preditor independente de redução de reserva ovariana (OR 2.8, IC95% 1.6-4.9).",
        "section": "abstract",
        "similarity": 0.700,
        "word_count": 87
      }
      // ... mais 28 chunks
    ]
  },

  "metadata": {
    "total_chunks": 30,
    "articles_count": 11,
    "avg_similarity": 0.644,
    "min_similarity": 0.501,
    "max_similarity": 0.730,
    "sections_distribution": {
      "abstract": 8,
      "results": 12,
      "discussion": 7,
      "methods": 3
    },
    "total_word_count": 6840,
    "avg_chunk_length": 228,
    "recency_stats": {
      "newest_year": 2024,
      "oldest_year": 2018,
      "avg_year": 2021
    },
    "quality_grade": "excellent",
    "threshold_used": 0.35,
    "threshold_type": "normal"
  },

  "prompt_clinical_relevance": "Você é um médico especialista em medicina baseada em evidências.\n\n**ScoreItem:** Leptina - Mulheres (Exames - Laboratoriais)\n**Nome curto:** Leptina - Mulheres\n**Unidade:** ng/mL\n**Gênero aplicável:** female\n\n**Contexto científico disponível:**\n- 30 chunks de 11 artigos científicos\n- Avg similarity: 0.644\n- Seções: results, discussion, methods, abstract\n\n## CHUNKS CIENTÍFICOS\n\n### Chunk 1/30\n**Article:** Elevated Serum Leptin Levels as a Predictive Marker for Polycystic Ovary Syndrome (2022)\n**Journal:** J Endocrinol Invest\n**Section:** abstract | **Similarity:** 0.730\n\nEstudo demonstrando que níveis elevados de leptina sérica estão significativamente associados à síndrome dos ovários policísticos (SOP). Níveis de leptina >11.58 ng/mL apresentam sensibilidade de 77.5% e especificidade de 62.6% para predição de SOP, especialmente em mulheres com hiperandrogenismo e sobrepeso/obesidade. A leptina, hormônio derivado do tecido adiposo, desempenha papel crucial na regulação da homeostase energética e reprodução. Em mulheres com SOP, hiperleptinemia correlaciona-se com resistência insulínica, obesidade central e hiperandrogenemia, formando um ciclo metabólico-reprodutivo adverso.\n\n---\n\n### Chunk 2/30\n**Article:** The association between leptin, adiponectin levels and the ovarian reserve in women of reproductive age (2024)\n**Journal:** Reprod Health\n**Section:** abstract | **Similarity:** 0.700\n\nEstudo investigando a relação entre adipocinas (leptina e adiponectina) e reserva ovariana em mulheres em idade reprodutiva. Demonstra que mulheres com SOP apresentam níveis mais elevados de leptina (média 18.3±6.2 ng/mL vs 8.1±3.4 ng/mL em controles, p<0.001) e menores de adiponectina, sugerindo que alterações hormonais e metabólicas podem influenciar negativamente a fertilidade. A relação leptina/adiponectina >2.0 foi identificada como preditor independente de redução de reserva ovariana (OR 2.8, IC95% 1.6-4.9).\n\n---\n\n[... chunks 3-30 ...]\n\n## TAREFA\n\nAnalise os 30 chunks científicos acima e gere um texto de **RELEVÂNCIA CLÍNICA** (clinical_relevance) para médicos:\n\n**Requisitos:**\n- Extensão: 1200-1800 caracteres\n- Linguagem técnica e precisa\n- Incluir: definição fisiológica, valores de referência, fisiopatologia resumida\n- Citar dados epidemiológicos com NÚMEROS (prevalência, RR, OR, sensibilidade/especificidade)\n- Estratificação de risco por valores\n- Considerações populacionais (idade, gênero, comorbidades)\n- Base em evidências dos chunks fornecidos\n\n**Retorne APENAS o texto da relevância clínica, sem preâmbulos.**",

  "prompt_patient_explanation": "Você é um médico especialista em comunicação com pacientes.\n\n**ScoreItem:** Leptina - Mulheres (Exames - Laboratoriais)\n**Nome curto:** Leptina - Mulheres\n**Unidade:** ng/mL\n**Gênero aplicável:** female\n\n**Contexto científico disponível:**\n- 30 chunks de 11 artigos científicos\n- Avg similarity: 0.644\n- Seções: results, discussion, methods, abstract\n\n## CHUNKS CIENTÍFICOS\n\n[... mesmos 30 chunks ...]\n\n## TAREFA\n\nAnalise os 30 chunks científicos acima e gere uma **EXPLICAÇÃO PARA PACIENTE** (patient_explanation):\n\n**Requisitos:**\n- Extensão: 600-900 caracteres\n- Linguagem SIMPLES (8º ano escolar)\n- Explicar: O QUE é o parâmetro, POR QUE importa, O QUE significam valores alterados\n- Tom empático e empoderador\n- Sem jargão médico complexo\n- Base em evidências dos chunks\n\n**Retorne APENAS o texto da explicação, sem preâmbulos.**",

  "prompt_conduct": "Você é um médico especialista em protocolos clínicos.\n\n**ScoreItem:** Leptina - Mulheres (Exames - Laboratoriais)\n**Nome curto:** Leptina - Mulheres\n**Unidade:** ng/mL\n**Gênero aplicável:** female\n\n**Contexto científico disponível:**\n- 30 chunks de 11 artigos científicos\n- Avg similarity: 0.644\n- Seções: results, discussion, methods, abstract\n\n## CHUNKS CIENTÍFICOS\n\n[... mesmos 30 chunks ...]\n\n## TAREFA\n\nAnalise os 30 chunks científicos acima e gere **CONDUTAS CLÍNICAS** (conduct) em formato Markdown:\n\n**Requisitos:**\n- Extensão: 1000-1500 caracteres\n- Formato Markdown com seções:\n  ## Interpretação de Valores\n  ## Investigação Complementar\n  ## Protocolo de Tratamento\n  ## Critérios de Referência/Encaminhamento\n- Condutas baseadas em evidências dos chunks\n- Incluir valores de corte quando aplicável (ex: >11.58 ng/mL)\n- Especificar exames complementares\n- Critérios objetivos de encaminhamento\n\n**Retorne APENAS o texto da conduta em Markdown, sem preâmbulos.**",

  "prompt_max_points": "Você é um médico especialista em estratificação de risco.\n\n**ScoreItem:** Leptina - Mulheres (Exames - Laboratoriais)\n**Nome curto:** Leptina - Mulheres\n**Unidade:** ng/mL\n**Gênero aplicável:** female\n\n**Contexto científico disponível:**\n- 30 chunks de 11 artigos científicos\n- Avg similarity: 0.644\n\n## CHUNKS CIENTÍFICOS\n\n[... mesmos 30 chunks ...]\n\n## TAREFA\n\nAnalise os 30 chunks científicos acima e determine **MAX_POINTS** (0-50) para este parâmetro:\n\n**Critérios de pontuação:**\n- Impacto prognóstico (0-20 pontos): RR de mortalidade, eventos CV, fertilidade\n- Modificabilidade (0-15 pontos): Lifestyle, farmacológico, cirúrgico\n- Prevalência (0-10 pontos): Comum vs raro\n- Urgência clínica (0-5 pontos): Emergência vs rotina\n\n**Baseado nos chunks:**\n- Leptina >11.58 prediz SOP (OR 2.8)\n- Correlaciona com resistência insulínica e obesidade\n- Modificável via lifestyle e perda de peso\n- Prevalente em mulheres com sobrepeso\n\n**Retorne APENAS o número (0-50) seguido de 1 linha de justificativa.**\nExemplo: '25 - Impacto moderado em fertilidade e metabolismo, modificável via lifestyle, prevalente em SOP'",

  "status": "ready",
  "created_at": "2026-02-17T17:45:00Z",
  "updated_at": "2026-02-17T17:45:00Z"
}
```

---

## 📋 Estrutura dos 4 Prompts

### Prompt 1: Clinical Relevance (1200-1800 chars)

```
✅ Inclui: fullName com contexto (Group - Subgroup - Name)
✅ Inclui: unit, gender, age range
✅ Inclui: 30 chunks científicos COMPLETOS
✅ Requisitos específicos de extensão e conteúdo
✅ Instrução clara de formato de resposta
```

### Prompt 2: Patient Explanation (600-900 chars)

```
✅ Mesmos chunks científicos
✅ Requisitos de linguagem simples
✅ Foco em comunicação paciente
```

### Prompt 3: Conduct (1000-1500 chars)

```
✅ Mesmos chunks científicos
✅ Formato Markdown com seções específicas
✅ Valores de corte dos chunks
```

### Prompt 4: Max Points (0-50)

```
✅ Mesmos chunks científicos
✅ Critérios objetivos de pontuação
✅ Exemplos baseados em evidências
```

---

## 🎯 Vantagens deste Formato

1. **Prompts prontos** - só enviar para Claude API
2. **Contexto completo** - fullName + chunks + metadados
3. **4 chamadas separadas** - uma por campo
4. **Validação fácil** - checar extensão de cada resposta
5. **Rastreável** - metadata mostra quality_grade e threshold usado

---

## ⚙️ Como Usar (Enrichment)

```bash
# 1. Buscar preparation
prep = GET /api/preparations/:score_item_id

# 2. Enviar 4 prompts para Claude API

# Prompt 1
response_cr = claude_api.messages.create(
    model="claude-sonnet-4",
    messages=[{"role": "user", "content": prep.prompt_clinical_relevance}]
)

# Prompt 2
response_pe = claude_api.messages.create(
    model="claude-sonnet-4",
    messages=[{"role": "user", "content": prep.prompt_patient_explanation}]
)

# Prompt 3
response_cond = claude_api.messages.create(
    model="claude-sonnet-4",
    messages=[{"role": "user", "content": prep.prompt_conduct}]
)

# Prompt 4
response_pts = claude_api.messages.create(
    model="claude-sonnet-4",
    messages=[{"role": "user", "content": prep.prompt_max_points}]
)

# 3. Atualizar ScoreItem
PUT /api/score-items/:id {
    "clinical_relevance": response_cr.content,
    "patient_explanation": response_pe.content,
    "conduct": response_cond.content,
    "points": parse_int(response_pts.content)
}
```

---

## 🔧 Correções Necessárias no Código

O código atual NÃO está salvando os prompts. Preciso corrigir:

1. `savePreparationWithPrompts()` não está atribuindo os prompts corretamente ao objeto
2. `generateEnrichmentPrompts()` pode não estar sendo chamada
3. Verificar se `prepRepo.Create()` está salvando os campos TEXT

**TODO:** Debugar e corrigir a cadeia de preparação para popular os 4 campos de prompt.
