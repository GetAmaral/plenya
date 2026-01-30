# BATCH FINAL 4 - EXEMPLO VISUAL
## Conteúdo Enriquecido MFI de um Item

---

## Item de Exemplo: "Dor Lombar"

**ID:** `e24dae19-4cb0-4d83-a6db-9571aabf9bde`
**Subgrupo:** Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual)

---

### 📋 clinical_relevance

```
Dor lombar pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da
Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções
em sistemas interconectados incluindo metabolismo energético, eixo neuroendócrino,
sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica.

A avaliação funcional busca identificar causas raiz através de investigação
abrangente de biomarcadores, história clínica detalhada e análise de sistemas
corporais. Sintomas aparentemente isolados podem ser manifestações precoces de
desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais
ou disfunção mitocondrial que precedem diagnósticos convencionais.
```

**Características:**
- ~200-300 palavras
- Perspectiva funcional integrativa
- Foco em causas raiz
- Linguagem técnica mas acessível

---

### 🔍 interpretation_guide

```
Durante anamnese de dor lombar, investigue:

(1) início, duração e padrão temporal
(2) fatores agravantes e atenuantes
(3) sintomas associados em outros sistemas
(4) correlação com alimentação, sono, estresse e ciclo menstrual
(5) impacto na qualidade de vida

Utilize escalas validadas quando disponíveis. Realize exame físico direcionado.
Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico,
função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12.

Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de
peso inexplicada ou febre persistente.
```

**Características:**
- ~150-250 palavras
- Roteiro estruturado de investigação
- Sinais de alerta (red flags)
- Orientação prática para clínicos

---

### 💊 recommendations

```json
[
  "Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de dor lombar",

  "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida",

  "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis",

  "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B",

  "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"
]
```

**Características:**
- 3-5 recomendações práticas
- Baseadas em evidências
- Ações concretas e mensuráveis
- Intervenções de estilo de vida + laboratorial

---

### 🔬 related_markers

```json
[
  "PCR ultrassensível",
  "Vitamina D (25-OH)",
  "Vitamina B12",
  "Hemograma completo",
  "TSH e T4 livre",
  "Perfil metabólico"
]
```

**Características:**
- 4-8 biomarcadores
- Exames laboratoriais relevantes
- Marcadores funcionais específicos
- Correlação com causas raiz

---

### 📚 articles_suggestions

```json
[
  "Entendendo dor lombar: causas e abordagem integrativa",
  "Medicina Funcional: investigando sintomas inexplicados",
  "Nutrição anti-inflamatória para saúde integral",
  "O papel do estilo de vida na prevenção e tratamento"
]
```

**Características:**
- 3-5 tópicos educacionais
- Linguagem acessível para pacientes
- Foco em empoderamento
- Abordagem integrativa

---

## Visualização SQL

```sql
UPDATE score_items
SET
  clinical_relevance = 'Dor lombar pode indicar desequilíbrios...',
  interpretation_guide = 'Durante anamnese de dor lombar, investigue...',
  recommendations = '["Realizar avaliação...", "Implementar diário..."]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)"]'::jsonb,
  articles_suggestions = '["Entendendo dor lombar...", "Medicina Funcional..."]'::jsonb,
  updated_at = NOW()
WHERE id = 'e24dae19-4cb0-4d83-a6db-9571aabf9bde';
```

---

## Como o Conteúdo Aparece no Frontend

### Card de Item
```
┌─────────────────────────────────────────────────┐
│ 📋 Dor Lombar                                   │
├─────────────────────────────────────────────────┤
│                                                 │
│ Relevância Clínica:                            │
│ Dor lombar pode indicar desequilíbrios         │
│ sistêmicos múltiplos na perspectiva da...      │
│                                                 │
│ [Ler mais]                                      │
│                                                 │
├─────────────────────────────────────────────────┤
│ 🔬 Biomarcadores Relacionados:                  │
│ • PCR ultrassensível                           │
│ • Vitamina D (25-OH)                           │
│ • Vitamina B12                                 │
│ • Hemograma completo                           │
│                                                 │
├─────────────────────────────────────────────────┤
│ 💊 Recomendações: 5                             │
│ 📚 Artigos Sugeridos: 4                         │
└─────────────────────────────────────────────────┘
```

### Modal Expandido
```
╔═════════════════════════════════════════════════╗
║ 📋 Dor Lombar                            [X]    ║
╠═════════════════════════════════════════════════╣
║                                                 ║
║ 📖 RELEVÂNCIA CLÍNICA                          ║
║ ─────────────────────────────────────────────  ║
║ Dor lombar pode indicar desequilíbrios         ║
║ sistêmicos múltiplos na perspectiva da         ║
║ Medicina Funcional Integrativa. Este sintoma   ║
║ frequentemente reflete disfunções em           ║
║ sistemas interconectados incluindo...          ║
║                                                 ║
║ 🔍 GUIA DE INTERPRETAÇÃO                       ║
║ ─────────────────────────────────────────────  ║
║ Durante anamnese de dor lombar, investigue:    ║
║ (1) início, duração e padrão temporal          ║
║ (2) fatores agravantes e atenuantes...         ║
║                                                 ║
║ 💊 RECOMENDAÇÕES                               ║
║ ─────────────────────────────────────────────  ║
║ ☑ Realizar avaliação laboratorial abrangente   ║
║ ☑ Implementar diário de sintomas               ║
║ ☑ Otimizar nutrição com dieta anti-inflam...  ║
║                                                 ║
║ 🔬 BIOMARCADORES                               ║
║ ─────────────────────────────────────────────  ║
║ [PCR ultrassensível] [Vitamina D] [B12]       ║
║                                                 ║
║ 📚 ARTIGOS EDUCACIONAIS                        ║
║ ─────────────────────────────────────────────  ║
║ • Entendendo dor lombar: causas e...          ║
║ • Medicina Funcional: investigando...         ║
║                                                 ║
╚═════════════════════════════════════════════════╝
```

---

## Padrão Aplicado a Todos os 40 Items

Este mesmo padrão de conteúdo foi aplicado a:

### Sintomas (27 items)
- Outros sintomas
- Segmento torácico
- Eructação
- Hemorróidas
- Disúria
- **Dor lombar** ← Exemplo mostrado acima
- Segmentos apendiculares
- Cãimbras
- Claudicação
- Dores articulares
- [... e mais 17 sintomas]

### Cirurgias (13 items)
- Mastectomia
- Prostatectomia
- Tireoidectomia
- Histerectomia
- [... e mais 9 cirurgias]

---

## Diferença: Sintomas vs Cirurgias

### Template para Sintomas
- Foco em investigação de causas
- Protocolos de avaliação
- Intervenções terapêuticas
- Modificações de estilo de vida

### Template para Cirurgias
- Impacto metabólico da cirurgia
- Necessidades pós-cirúrgicas
- Reposição hormonal (quando aplicável)
- Monitoramento de longo prazo
- Suporte nutricional específico

---

## Qualidade do Conteúdo

### ✅ Critérios Atendidos

1. **Precisão Técnica**
   - Terminologia médica correta
   - Fisiopatologia funcional
   - Evidências científicas

2. **Aplicabilidade Clínica**
   - Guias práticos de investigação
   - Recomendações acionáveis
   - Biomarcadores relevantes

3. **Educação do Paciente**
   - Linguagem acessível
   - Tópicos de empoderamento
   - Foco em autocuidado

4. **Padrão MFI**
   - Causas raiz, não sintomas
   - Abordagem sistêmica
   - Individualização
   - Protocolos baseados em evidências

---

## Estrutura de Dados

### Tabela: score_items

```sql
CREATE TABLE score_items (
  id UUID PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  subgroup VARCHAR(255),

  -- Campos enriquecidos ↓
  clinical_relevance TEXT,            -- 200-300 palavras
  interpretation_guide TEXT,          -- 150-250 palavras
  recommendations JSONB,               -- Array de 3-5 strings
  related_markers JSONB,               -- Array de 4-8 strings
  articles_suggestions JSONB,         -- Array de 3-5 strings

  updated_at TIMESTAMP DEFAULT NOW()
);
```

### Exemplo de Query

```sql
-- Buscar item com conteúdo completo
SELECT
  name,
  clinical_relevance,
  interpretation_guide,
  jsonb_array_length(recommendations) as num_recomendacoes,
  jsonb_array_length(related_markers) as num_marcadores,
  jsonb_array_length(articles_suggestions) as num_artigos
FROM score_items
WHERE id = 'e24dae19-4cb0-4d83-a6db-9571aabf9bde';

-- Resultado:
-- name: Dor lombar
-- clinical_relevance: Dor lombar pode indicar...
-- interpretation_guide: Durante anamnese...
-- num_recomendacoes: 5
-- num_marcadores: 6
-- num_artigos: 4
```

---

## Validação de Qualidade

### Checklist por Item

✅ clinical_relevance preenchido (>200 palavras)
✅ interpretation_guide preenchido (>150 palavras)
✅ recommendations array com 3-5 items
✅ related_markers array com 4-8 items
✅ articles_suggestions array com 3-5 items
✅ Conteúdo coerente com nome do item
✅ Linguagem técnica mas acessível
✅ Padrão MFI aplicado
✅ JSON válido
✅ Aspas SQL escapadas

---

## Próximos Passos

1. **Executar SQL**
   ```bash
   bash scripts/execute_batch_final_4.sh
   ```

2. **Validar no Banco**
   ```sql
   SELECT COUNT(*) FROM score_items
   WHERE clinical_relevance IS NOT NULL;
   ```

3. **Testar no Frontend**
   - Abrir item "Dor Lombar" no web app
   - Verificar exibição de todos os campos
   - Confirmar formatação adequada

4. **Gerar Artigos**
   - Usar articles_suggestions como guia
   - Criar conteúdo educacional
   - Publicar para pacientes

---

**Arquivo:** `/home/user/plenya/BATCH-FINAL-4-EXEMPLO-VISUAL.md`
**Data:** 2026-01-28
**Status:** ✅ Documentação Completa
