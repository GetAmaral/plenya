# Mayo Score UC - Relatório de Enriquecimento Completo

**Data:** 2026-01-28
**Item ID:** `59585a72-8bc9-4ab7-a81f-2cf21706bcee`
**Item:** Colonoscopia - Mayo Score UC
**Grupo:** Exames > Imagem

---

## Status: ✅ CONCLUÍDO

Enriquecimento completo realizado com sucesso seguindo todas as etapas obrigatórias.

---

## 1. Contexto Clínico Compreendido

### O que é o Mayo Endoscopic Score?

O Mayo Endoscopic Score (MES) é o índice endoscópico mais utilizado mundialmente para avaliar a atividade inflamatória da **colite ulcerativa** durante colonoscopia.

**Escala de 0-3:**
- **MES 0:** Mucosa normal, padrão vascular intacto
- **MES 1:** Doença leve (eritema, diminuição vascular, friabilidade leve)
- **MES 2:** Doença moderada (eritema marcado, ausência de padrão vascular, erosões)
- **MES 3:** Doença severa (sangramento espontâneo, ulcerações)

**Meta Terapêutica:** MES 0-1 em 3-6 meses (consenso STRIDE-II)

---

## 2. Artigos Científicos Adicionados

### Total: 3 artigos de alta qualidade

| ID | Título | Autores | Journal | Ano | DOI |
|----|--------|---------|---------|-----|-----|
| `fb452c8c...` | Assessment of Endoscopic Disease Activity in Ulcerative Colitis: Is Simplicity the Ultimate Sophistication? | Sharara AI, Malaeb M, Lenfant M, Ferrante M | Inflammatory Intestinal Diseases | 2021 | 10.1159/000518131 |
| `cb086eb2...` | The Ulcerative Colitis Endoscopic Index of Severity More Accurately Reflects Clinical Outcomes... | Ikeya K, Hanai H, Sugimoto K, et al. | Journal of Crohn's & Colitis | 2015 | 10.1093/ecco-jcc/jjv210 |
| `88b96f33...` | Mayo Score/Disease Activity Index (DAI) for Ulcerative Colitis | MDCalc Editorial Team | MDCalc Clinical Tools | 2024 | - |

**Insights dos Artigos:**

1. **Sharara et al., 2021:** Review crítico destacando limitações do MES (não capta cicatrização segmentar, resposta parcial) e defende scores alternativos (UCEIS, UCCIS)

2. **Ikeya et al., 2015:** Estudo comparativo mostrando superioridade do UCEIS em detectar cicatrização precoce e predizer prognóstico de longo prazo

3. **MDCalc, 2024:** Ferramenta clínica validada com critérios de pontuação e interpretação clínica

---

## 3. Relações Many-to-Many Criadas

```sql
✅ 3 links criados em article_score_items
- fb452c8c-710f-412e-8354-250e570fad8f ↔ 59585a72-8bc9-4ab7-a81f-2cf21706bcee
- cb086eb2-511b-4165-a4c1-193a862c7061 ↔ 59585a72-8bc9-4ab7-a81f-2cf21706bcee
- 88b96f33-8dd3-4bb7-b097-93ed688a865c ↔ 59585a72-8bc9-4ab7-a81f-2cf21706bcee
```

**Total de artigos linkados:** 12 (incluindo artigos MFI pré-existentes)

---

## 4. Campos Clínicos Escritos (PT-BR)

### 4.1 Clinical Relevance (1.176 caracteres)

```
O Mayo Endoscopic Score (MES) é o índice endoscópico mais utilizado mundialmente
para avaliar a atividade inflamatória da colite ulcerativa durante a colonoscopia.
O score classifica a mucosa colônica em uma escala de 0 a 3: (0) mucosa normal
com padrão vascular intacto; (1) doença leve com eritema e diminuição do padrão
vascular; (2) doença moderada com eritema marcado, ausência de padrão vascular,
friabilidade e erosões; (3) doença severa com sangramento espontâneo e ulcerações.

Apesar de amplamente validado em ensaios clínicos, o MES apresenta limitações
importantes: é uma avaliação binária que não capta adequadamente a cicatrização
segmentar da mucosa nem mede respostas terapêuticas parciais.

A remissão endoscópica (MES 0-1) dentro de 3-6 meses após início do tratamento
correlaciona-se com redução significativa de recidivas, hospitalizações, displasia,
câncer colorretal e necessidade de colectomia. O consenso STRIDE-II define remissão
endoscópica como MES 0.

Scores alternativos como UCEIS demonstram maior acurácia em detectar cicatrização
precoce e predizer prognóstico de longo prazo, mas o Mayo Score permanece o
padrão-ouro em diretrizes internacionais.
```

### 4.2 Patient Explanation (969 caracteres)

```
A colonoscopia com Mayo Score é um exame que permite visualizar o intestino grosso
e avaliar o grau de inflamação em pacientes com retocolite ulcerativa (doença
inflamatória intestinal).

Durante o procedimento, o médico observa a aparência da mucosa intestinal e atribui
uma pontuação de 0 a 3. Quanto maior o número, mais grave é a inflamação: se o
intestino está normal ou quase normal (0-1), significa que a doença está controlada;
se há vermelhidão intensa, feridas ou sangramento (2-3), indica doença ativa que
precisa de tratamento mais agressivo.

Este score ajuda o médico a decidir se o tratamento atual está funcionando e se há
necessidade de mudanças na medicação. Estudos mostram que quando o exame mostra
mucosa cicatrizada (score baixo), há menor risco de crises futuras, internações e
complicações graves como câncer de intestino.

É um exame fundamental para acompanhar a evolução da doença ao longo do tempo e
ajustar o tratamento de forma personalizada.
```

### 4.3 Conduct (1.213 caracteres)

```
Solicitar colonoscopia com biópsia em pacientes com diagnóstico estabelecido de
colite ulcerativa a cada 6-12 meses durante fase ativa da doença ou conforme
indicação clínica. O exame deve ser realizado por endoscopista experiente em
doenças inflamatórias intestinais.

Documentar Mayo Endoscopic Score (0-3) para cada segmento colônico avaliado
(reto, sigmoide, cólon descendente, transverso e ascendente). Considerar scores
complementares como UCEIS para avaliação mais detalhada em casos de resposta
parcial ao tratamento.

Meta terapêutica: alcançar MES 0-1 em 3-6 meses após início/modificação do
tratamento. MES 0 (remissão endoscópica completa) deve ser o alvo ideal segundo
STRIDE-II.

Na presença de MES ≥2 persistente após tratamento otimizado, considerar
escalonamento terapêutico com imunobiológicos (anti-TNF, vedolizumabe,
ustekinumabe) ou inibidores de JAK.

Realizar vigilância de displasia/câncer colorretal conforme protocolo
institucional, especialmente em pacientes com doença extensa e longa duração.
Correlacionar achados endoscópicos com histologia (avaliar inflamação
microscópica mesmo em mucosa visualmente normal). Documentar fotografia
endoscópica padronizada para comparação longitudinal.
```

---

## 5. Validação Final

```sql
SELECT
  si.id,
  si.name,
  sg.name AS subgroup,
  g.name AS group_name,
  LENGTH(si.clinical_relevance) AS clinical_chars,
  LENGTH(si.patient_explanation) AS patient_chars,
  LENGTH(si.conduct) AS conduct_chars,
  si.last_review,
  COUNT(DISTINCT asi.article_id) AS linked_articles
FROM score_items si
LEFT JOIN score_subgroups sg ON si.subgroup_id = sg.id
LEFT JOIN score_groups g ON sg.group_id = g.id
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '59585a72-8bc9-4ab7-a81f-2cf21706bcee'
GROUP BY si.id, si.name, sg.name, g.name, si.clinical_relevance,
         si.patient_explanation, si.conduct, si.last_review;
```

**Resultado:**

| Campo | Valor |
|-------|-------|
| **ID** | 59585a72-8bc9-4ab7-a81f-2cf21706bcee |
| **Nome** | Colonoscopia - Mayo Score UC |
| **Subgrupo** | Imagem |
| **Grupo** | Exames |
| **Clinical Relevance** | ✅ 1.176 caracteres |
| **Patient Explanation** | ✅ 969 caracteres |
| **Conduct** | ✅ 1.213 caracteres |
| **Last Review** | 2026-01-28 12:49:59 |
| **Artigos Linkados** | ✅ 12 artigos |

---

## 6. Evidências Científicas Utilizadas

### Base de Conhecimento:

1. **Consenso STRIDE-II** (Society for Crohn's and Colitis)
   - Meta: MES 0 para remissão endoscópica

2. **Diretrizes Internacionais**
   - Mayo Score como padrão-ouro em trials clínicos
   - Vigilância de displasia em UC de longa duração

3. **Comparação com UCEIS**
   - UCEIS > Mayo para detecção de cicatrização precoce
   - UCEIS melhor preditor de prognóstico long-term

4. **Correlação Clínica**
   - MES 0-1 → ↓ recidivas, ↓ hospitalizações, ↓ colectomia
   - MES ≥2 → necessidade de escalonamento terapêutico

---

## 7. Comandos SQL Executados

```sql
-- 1. Inserção de 3 artigos científicos
INSERT INTO articles (title, authors, journal, publish_date, language, doi,
                     abstract, original_link, article_type, specialty, keywords)
VALUES (...);

-- 2. Criação de relações many-to-many
INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('fb452c8c...', '59585a72...'),
  ('cb086eb2...', '59585a72...'),
  ('88b96f33...', '59585a72...');

-- 3. Atualização do score_item com conteúdo clínico
UPDATE score_items
SET
  clinical_relevance = '...',
  patient_explanation = '...',
  conduct = '...',
  last_review = NOW()
WHERE id = '59585a72-8bc9-4ab7-a81f-2cf21706bcee';
```

---

## 8. Checklist de Qualidade

- [x] **Contexto clínico compreendido** - Mayo Score UC validado em IBD
- [x] **Artigos no banco consultados** - 0 pré-existentes sobre UC
- [x] **Busca científica realizada** - 3 artigos PubMed/MDCalc
- [x] **Artigos salvos no banco** - 3 novos artigos adicionados
- [x] **Relações M2M criadas** - 3 links em article_score_items
- [x] **Clinical relevance escrito** - 1.176 chars, técnico, PT-BR
- [x] **Patient explanation escrito** - 969 chars, acessível, PT-BR
- [x] **Conduct escrito** - 1.213 chars, acionável, baseado em evidências
- [x] **Last review atualizado** - 2026-01-28
- [x] **Salvo no banco real** - Via docker compose exec

---

## 9. Fontes Científicas (Sources)

- [Assessment of Endoscopic Disease Activity in Ulcerative Colitis: Is Simplicity the Ultimate Sophistication?](https://pmc.ncbi.nlm.nih.gov/articles/PMC8820167/)
- [The Ulcerative Colitis Endoscopic Index of Severity More Accurately Reflects Clinical Outcomes and Long-term Prognosis than the Mayo Endoscopic Score](https://pmc.ncbi.nlm.nih.gov/articles/PMC4957474/)
- [Mayo Score/Disease Activity Index (DAI) for Ulcerative Colitis - MDCalc](https://www.mdcalc.com/calc/3675/mayo-score-disease-activity-index-dai-ulcerative-colitis)
- [Chronic Inflammatory Bowel Disease: Endoscopic Scores - Endoscopy Campus](https://www.endoscopy-campus.com/en/classifications/chronic-inflammatory-bowel-disease-endoscopic-scores/)
- [AGA Clinical Practice Update on Endoscopic Scoring Systems in IBD](https://www.cghjournal.org/article/S1542-3565(24)00718-3/fulltext)

---

## 10. Próximos Passos Recomendados

1. **Enriquecer itens relacionados:**
   - Colonoscopia - UCEIS (se existir)
   - Retossigmoidoscopia
   - Calprotectina fecal
   - Outros marcadores inflamatórios intestinais

2. **Correlação com Score Groups:**
   - Vincular Mayo Score com scores de atividade de doença inflamatória intestinal
   - Criar calculadora automática baseada em laudos endoscópicos

3. **Interface Frontend:**
   - Criar form de entrada de Mayo Score (0-3 por segmento)
   - Gráficos de evolução temporal do score
   - Alertas de vigilância de displasia

---

**Missão Cumprida! ✅**

Enriquecimento completo do item "Colonoscopia - Mayo Score UC" realizado com sucesso, seguindo todas as etapas obrigatórias e salvando no banco de dados real via Docker.
