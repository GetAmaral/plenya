# Enriquecimento: Hemácias - Mulheres - Sumário Executivo

## Informações do Item

| Campo | Valor |
|-------|-------|
| **Nome** | Hemácias - Mulheres |
| **ID** | `501fd84a-a440-4c13-9b11-35e2f69017d1` |
| **Grupo** | Exames > Laboratoriais |
| **Tipo** | Contagem de eritrócitos |
| **Valores Referência** | 4.0 - 5.2 milhões/µL |

## Contexto Clínico

### Diferenças de Gênero
- **Mulheres:** 4.0-5.2 milhões/µL
- **Homens:** 4.5-5.5 milhões/µL

### Por Que Valores São Menores em Mulheres?

1. **Menstruação**
   - Perda mensal de sangue (30-80 mL/ciclo)
   - Depleção de ferro ao longo do tempo
   - Maior risco de anemia ferropriva

2. **Fatores Hormonais**
   - Estrogênio tem efeito supressor leve na eritropoiese
   - Testosterona (mais alta em homens) estimula produção de eritrócitos
   - EPO (eritropoietina) sensibilidade diferente

3. **Fisiologia**
   - Menor massa muscular (menos demanda de O₂)
   - Volume plasmático pode ser proporcionalmente maior

### Alterações Patológicas

#### Hemácias Baixas (<4.0 milhões/µL)

**Causas Comuns:**
- Anemia ferropriva (menstruação abundante)
- Deficiência de vitamina B12/folato
- Doença crônica (inflamatória, renal)
- Gravidez (hemodiluição fisiológica)
- Hemólise (talassemia, anemia falciforme)

**Sintomas:**
- Fadiga, cansaço crônico
- Palidez cutâneo-mucosa
- Dispneia aos esforços
- Taquicardia
- Cefaleia, tontura

**Investigação:**
- Hemoglobina, hematócrito
- VCM, HCM, CHCM (índices)
- Ferritina, ferro sérico, transferrina
- Vitamina B12, folato
- Reticulócitos (avaliar resposta medular)

#### Hemácias Altas (>5.2 milhões/µL)

**Causas Comuns:**
- Desidratação (pseudopolicitemia)
- Hipóxia crônica (altitude, DPOC)
- Policitemia vera (neoplasia mieloproliferativa)
- Uso de EPO exógena (doping, IRC em hemodiálise)

**Sintomas:**
- Rubor facial
- Cefaleia, vertigem
- Prurido (especialmente após banho quente)
- Risco trombótico aumentado

**Investigação:**
- Hematócrito (>48% em mulheres suspeita policitemia)
- Gasometria arterial (saturação O₂)
- Eritropoietina sérica
- JAK2 mutation (policitemia vera)
- Avaliação cardiovascular

## Busca de Artigos Científicos

### Palavras-Chave Recomendadas

**PubMed:**
```
"red blood cell count"[MeSH] AND "women"[MeSH] AND "reference values"
"erythrocyte count" AND "sex differences" AND "menstruation"
"iron deficiency anemia" AND "women" AND "menstrual blood loss"
"hemoglobin levels" AND "gender differences"
```

**Google Scholar:**
```
RBC reference ranges women men differences
erythrocyte count menstruation iron deficiency
sex hormones erythropoiesis women
```

### Tópicos a Cobrir nos Artigos

1. **Valores de referência** para hemácias em mulheres vs homens
2. **Impacto da menstruação** na contagem de eritrócitos
3. **Anemia ferropriva** em mulheres em idade reprodutiva
4. **Efeitos hormonais** (estrogênio, testosterona) na eritropoiese
5. **Gravidez e alterações hematológicas**

## Campos a Serem Enriquecidos

### 1. `clinical_relevance` (150-200 palavras)

**Conteúdo esperado:**
- Explicação das diferenças fisiológicas entre homens e mulheres
- Por que valores são menores em mulheres (menstruação, hormônios)
- Importância da avaliação conjunta com Hb, Hct, ferritina
- Quando suspeitar de patologia (anemia, policitemia)
- Relação com sintomas clínicos

**Tom:** Médico, técnico, objetivo

### 2. `patient_explanation` (100-150 palavras)

**Conteúdo esperado:**
- "Suas hemácias são células vermelhas que transportam oxigênio..."
- Valores normais para mulheres (4.0-5.2)
- Por que podem estar baixos (menstruação, alimentação)
- Sintomas de anemia (cansaço, palidez)
- Quando procurar médico

**Tom:** Leigo, acessível, empático

### 3. `conduct` (80-120 palavras)

**Conteúdo esperado:**

**Se hemácias < 4.0:**
- Solicitar: ferritina, ferro sérico, TIBC, transferrina
- Vitamina B12, ácido fólico
- Reticulócitos (avaliar resposta medular)
- História menstrual detalhada
- Considerar endoscopia digestiva alta/colonoscopia se idade >45 anos

**Se hemácias > 5.2:**
- Repetir exame em hidratação adequada
- Gasometria arterial
- EPO sérica
- Se persistente, considerar JAK2 mutation
- Referenciar para hematologia

**Tom:** Prático, protocolar, orientador

## Execução

### Pré-requisitos
- Docker e Docker Compose rodando
- Anthropic API key válida
- Containers `web` e `db` ativos

### Comando Rápido

```bash
# Definir API key
export ANTHROPIC_API_KEY="sk-ant-api03-..."

# Executar
./run-enrich-hemacias.sh
```

### Arquivos Criados

| Arquivo | Descrição |
|---------|-----------|
| `/home/user/plenya/scripts/enrich_hemacias_mulheres.mjs` | Script principal Node.js |
| `/home/user/plenya/apps/web/enrich.mjs` | Cópia no container web |
| `/home/user/plenya/run-enrich-hemacias.sh` | Script shell wrapper |
| `/home/user/plenya/EXECUTE-HEMACIAS-MULHERES.md` | Instruções detalhadas |
| `/home/user/plenya/HEMACIAS-MULHERES-SUMMARY.md` | Este arquivo |

## Fluxo de Execução

```
┌─────────────────────────────────────┐
│ 1. BUSCAR ARTIGOS CIENTÍFICOS      │
│    - Claude API com web search      │
│    - PubMed, Google Scholar         │
│    - 3-4 artigos relevantes         │
└──────────────┬──────────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│ 2. GERAR CONTEÚDO CLÍNICO (PT-BR)  │
│    - clinical_relevance             │
│    - patient_explanation            │
│    - conduct                        │
└──────────────┬──────────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│ 3. SALVAR ARTIGOS NO BANCO         │
│    - INSERT INTO articles           │
│    - article_type: 'scientific'     │
└──────────────┬──────────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│ 4. CRIAR RELAÇÕES                  │
│    - score_item_articles            │
│    - score_item_id + article_id     │
└──────────────┬──────────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│ 5. ATUALIZAR SCORE_ITEM            │
│    - UPDATE score_items             │
│    - WHERE id = '501fd8...'         │
└─────────────────────────────────────┘
```

## Validação Pós-Execução

### SQL Queries

```sql
-- 1. Verificar item atualizado
SELECT
  id,
  name,
  LENGTH(clinical_relevance) as clinical_chars,
  LENGTH(patient_explanation) as patient_chars,
  LENGTH(conduct) as conduct_chars,
  updated_at
FROM score_items
WHERE id = '501fd84a-a440-4c13-9b11-35e2f69017d1';

-- 2. Contar artigos vinculados
SELECT COUNT(*) as total_articles
FROM score_item_articles
WHERE score_item_id = '501fd84a-a440-4c13-9b11-35e2f69017d1';

-- 3. Ver artigos completos
SELECT
  a.id,
  a.title,
  a.url,
  a.article_type,
  a.created_at
FROM score_item_articles sia
JOIN articles a ON a.id = sia.article_id
WHERE sia.score_item_id = '501fd84a-a440-4c13-9b11-35e2f69017d1'
ORDER BY a.created_at DESC;

-- 4. Preview do conteúdo
SELECT
  name,
  LEFT(clinical_relevance, 200) as clinical_preview,
  LEFT(patient_explanation, 150) as patient_preview,
  LEFT(conduct, 100) as conduct_preview
FROM score_items
WHERE id = '501fd84a-a440-4c13-9b11-35e2f69017d1';
```

### Comandos Docker

```bash
# Verificar item atualizado
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT name, LENGTH(clinical_relevance) FROM score_items WHERE id='501fd84a-a440-4c13-9b11-35e2f69017d1';"

# Contar artigos
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT COUNT(*) FROM score_item_articles WHERE score_item_id='501fd84a-a440-4c13-9b11-35e2f69017d1';"
```

## Troubleshooting

### Erro: "API key não configurada"

**Causa:** Variável `ANTHROPIC_API_KEY` não está definida.

**Solução:**
```bash
export ANTHROPIC_API_KEY="sk-ant-api03-..."
```

### Erro: "Cannot find package '@anthropic-ai/sdk'"

**Causa:** Dependências não instaladas no container web.

**Solução:**
```bash
docker compose exec web pnpm add @anthropic-ai/sdk pg --filter web
docker compose restart web
```

### Erro: "Connection refused" ao banco

**Causa:** Container `db` não está rodando.

**Solução:**
```bash
docker compose up -d db
# Aguardar 5-10 segundos
docker compose exec db pg_isready -U plenya_user
```

### Erro: "Score item não existe"

**Causa:** ID `501fd84a-a440-4c13-9b11-35e2f69017d1` não encontrado.

**Solução:**
```bash
# Verificar se item existe
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT id, name FROM score_items WHERE name ILIKE '%hemácias%mulher%';"
```

## Métricas Esperadas

| Métrica | Valor Esperado |
|---------|----------------|
| **Tempo de execução** | 2-3 minutos |
| **Custo API (Claude)** | $0.05 - $0.10 USD |
| **Artigos salvos** | 3-4 artigos |
| **clinical_relevance** | 800-1200 caracteres |
| **patient_explanation** | 600-900 caracteres |
| **conduct** | 500-700 caracteres |
| **Tokens consumidos** | ~8,000-12,000 tokens |

## Próximos Itens Sugeridos

Após sucesso com "Hemácias - Mulheres", considerar enriquecer:

1. **Hemácias - Homens** (valores 4.5-5.5)
2. **Hemoglobina - Mulheres** (12-16 g/dL)
3. **Hemoglobina - Homens** (13-18 g/dL)
4. **Hematócrito - Mulheres** (36-48%)
5. **Ferritina - Mulheres** (diferenças pré/pós-menopausa)
6. **VCM** (Volume Corpuscular Médio)
7. **HCM** (Hemoglobina Corpuscular Média)
8. **CHCM** (Concentração de Hemoglobina)

## Referências Clínicas

### Guidelines
- WHO: Haemoglobin concentrations for the diagnosis of anaemia (2011)
- CDC: Iron Deficiency in Women's Health
- NICE: Anaemia - Iron Deficiency (CKS)

### Textbooks
- Harrison's Principles of Internal Medicine, 21st ed.
- Williams Hematology, 10th ed.
- Wintrobe's Clinical Hematology, 15th ed.

---

**Criado em:** 2026-01-28
**Autor:** Claude Code (Sonnet 4.5)
**Status:** Pronto para execução
**Revisão:** Pendente validação médica
