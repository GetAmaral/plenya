# EXEMPLO DE PROMPT REAL COMPLETO

## ScoreItem: Leptina - Mulheres

---

## PROMPT 1: CLINICAL_RELEVANCE

```
Você é um médico especialista em medicina baseada em evidências.

**ScoreItem:** Leptina - Mulheres (Exames - Laboratoriais)
**Unidade:** ng/mL
**Gênero aplicável:** female

**Contexto científico disponível:**
- 30 chunks de 11 artigos científicos
- Avg similarity: 0.644 (EXCELLENT quality)
- Seções: results (12), abstract (8), discussion (7), methods (3)
- Artigos: 2018-2024 (média 2021)

## CHUNKS CIENTÍFICOS

### Chunk 1/30
**Article:** Elevated Serum Leptin Levels as a Predictive Marker for Polycystic Ovary Syndrome (2022)
**Journal:** J Endocrinol Invest
**Section:** abstract | **Similarity:** 0.730

Estudo demonstrando que níveis elevados de leptina sérica estão significativamente associados à síndrome dos ovários policísticos (SOP). Níveis de leptina >11.58 ng/mL apresentam sensibilidade de 77.5% e especificidade de 62.6% para predição de SOP, especialmente em mulheres com hiperandrogenismo e sobrepeso/obesidade. A leptina, hormônio derivado do tecido adiposo, desempenha papel crucial na regulação da homeostase energética e reprodução. Em mulheres com SOP, hiperleptinemia correlaciona-se com resistência insulínica, obesidade central e hiperandrogenemia, formando um ciclo metabólico-reprodutivo adverso.

---

### Chunk 2/30
**Article:** The association between leptin, adiponectin levels and the ovarian reserve in women of reproductive age (2024)
**Journal:** Reprod Health
**Section:** abstract | **Similarity:** 0.700

Estudo investigando a relação entre adipocinas (leptina e adiponectina) e reserva ovariana em mulheres em idade reprodutiva. Demonstra que mulheres com SOP apresentam níveis mais elevados de leptina (média 18.3±6.2 ng/mL vs 8.1±3.4 ng/mL em controles, p<0.001) e menores de adiponectina, sugerindo que alterações hormonais e metabólicas podem influenciar negativamente a fertilidade. A relação leptina/adiponectina >2.0 foi identificada como preditor independente de redução de reserva ovariana (OR 2.8, IC95% 1.6-4.9).

---

### Chunk 3/30
**Article:** Síndrome dos Ovários Policísticos - Parte I (2024)
**Section:** results | **Similarity:** 0.695

Quase 60% das mulheres com SOP reportam insatisfação com os cuidados médicos atuais. O dado sinaliza falhas em educação, rastreio precoce e personalização. Leptina elevada correlaciona com resistência insulínica e obesidade central. Estudos mostram que redução de 5-10% do peso corporal em mulheres obesas com SOP resulta em normalização dos níveis de leptina e melhora significativa dos parâmetros metabólicos e reprodutivos.

---

[... chunks 4-30: artigos sobre leptina em mulheres, SOP, metabolismo, fertilidade ...]

## TAREFA

Analise os 30 chunks científicos acima e gere um texto de **RELEVÂNCIA CLÍNICA** (clinical_relevance) para médicos:

**Requisitos:**
- Extensão: 1200-1800 caracteres
- Linguagem técnica e precisa
- Incluir:
  * Definição fisiológica da leptina
  * Valores de referência (ex: >11.58 ng/mL)
  * Fisiopatologia (tecido adiposo, resistência insulínica, SOP)
  * Dados epidemiológicos com NÚMEROS dos chunks (sensibilidade 77.5%, OR 2.8, etc)
  * Estratificação de risco (normal vs elevado)
  * Considerações para mulheres (idade reprodutiva, SOP, obesidade)
- Base em evidências dos chunks fornecidos

**Retorne APENAS o texto da relevância clínica, sem preâmbulos ou formatação extra.**
```

**Tamanho total do prompt:** ~35,000 caracteres (30 chunks × ~1,000 chars + instruções)

---

## PROMPT 2: PATIENT_EXPLANATION

```
Você é um médico especialista em comunicação com pacientes.

**ScoreItem:** Leptina - Mulheres (Exames - Laboratoriais)
**Unidade:** ng/mL
**Gênero aplicável:** female

[... mesmos 30 chunks ...]

## TAREFA

Gere **EXPLICAÇÃO PARA PACIENTE** (600-900 caracteres):

**Requisitos:**
- Linguagem SIMPLES (8º ano escolar)
- Explicar:
  * O QUE é leptina (hormônio do tecido gorduroso)
  * POR QUE importa (controle de fome, fertilidade, metabolismo)
  * O QUE significam valores elevados (resistência, SOP, obesidade)
- Tom empático e empoderador
- Sem jargão: "adipocina" → "hormônio", "hiperandrogenismo" → "excesso de hormônios masculinos"

**Retorne APENAS o texto da explicação.**
```

---

## PROMPT 3: CONDUCT

```
Você é um médico especialista em protocolos clínicos.

**ScoreItem:** Leptina - Mulheres (Exames - Laboratoriais)
**Unidade:** ng/mL
**Gênero aplicável:** female

[... mesmos 30 chunks ...]

## TAREFA

Gere **CONDUTAS CLÍNICAS** (1000-1500 chars) em Markdown:

**Formato requerido:**
## Interpretação de Valores
- Normal: < 11.58 ng/mL (baseado em chunks)
- Elevado: ≥ 11.58 ng/mL (sensibilidade 77.5% para SOP)

## Investigação Complementar
- Se elevado: solicitar glicemia, insulina, HOMA-IR, perfil androgênico
- USG pélvica para avaliação ovariana
- Perfil lipídico

## Protocolo de Tratamento
- Perda de peso (5-10%) normaliza leptina em obesas
- Metformina se resistência insulínica
- Lifestyle: dieta, exercício

## Critérios de Referência
- Encaminhar para endocrinologia se leptina >20 ng/mL + SOP
- Ginecologia se infertilidade associada

**Retorne APENAS o Markdown, sem preâmbulos.**
```

---

## PROMPT 4: MAX_POINTS

```
Você é um médico especialista em estratificação de risco.

**ScoreItem:** Leptina - Mulheres
[... mesmos 30 chunks ...]

## TAREFA

Determine **MAX_POINTS** (0-50):

**Baseado nos chunks:**
- Preditor de SOP (OR 2.8) - impacto reprodutivo
- Correlaciona com resistência insulínica - impacto metabólico
- Modificável via perda de peso (5-10%)
- Prevalente em mulheres com sobrepeso/SOP

**Critérios:**
- Impacto prognóstico: 12/20 (fertilidade + metabólico)
- Modificabilidade: 12/15 (lifestyle efetivo)
- Prevalência: 6/10 (comum em SOP/obesidade)
- Urgência: 2/5 (não emergencial)

**Retorne número + justificativa (1 linha).**
Exemplo: "32 - Impacto moderado-alto em fertilidade e metabolismo, modificável via lifestyle, prevalente em SOP"
```

---

## 📊 RESUMO DO OBJETO COMPLETO

```json
{
  "id": "...",
  "score_item_id": "019bf31d-2ef0-71b6-a5d1-db49a4fa62fa",

  "selected_chunks": {
    "items": [ /* 30 chunks completos */ ]
  },

  "metadata": {
    "total_chunks": 30,
    "articles_count": 11,
    "avg_similarity": 0.644,
    "quality_grade": "excellent",
    "threshold_used": 0.35
  },

  "prompt_clinical_relevance": "...",      // ~35KB (30 chunks + instruções)
  "prompt_patient_explanation": "...",     // ~35KB (mesmos chunks + instruções)
  "prompt_conduct": "...",                 // ~35KB (mesmos chunks + instruções)
  "prompt_max_points": "...",              // ~35KB (mesmos chunks + instruções)

  "status": "ready"
}
```

**Cada prompt tem ~35,000 caracteres** (30 chunks científicos completos + contexto + instruções)

**Total de 4 prompts prontos para enviar ao Claude API!** ✅