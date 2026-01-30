# Score Items - Enriquecimento com IA
## Progresso da Tarefa Massiva

**Iniciado em:** 26 de Janeiro de 2026, 00:16 UTC
**Status:** 🔄 **EM ANDAMENTO**

---

## 📊 Visão Geral

### Escopo Total
- **Total de Score Items:** 2.316 (atualizado)
- **Items já processados:** 1 (NT-proBNP)
- **Items pendentes:** 2.315
- **Taxa de conclusão:** 0.04%

### Estratégia
- **Processamento:** Batches de 5-10 items
- **Agentes:** ai-engineer (especializados em IA/LLM)
- **Paralelização:** 1-2 agentes simultâneos
- **Tempo estimado total:** 40-80 horas (distribuídas)

---

## 🎯 Objetivos por Item

Para cada um dos 772 Score Items:

### 1. Coleta de Evidências
- ✅ Buscar nos 241 lectures MFI
- ✅ Buscar nos 6 research articles
- ✅ Buscar online (PubMed + Google Scholar)
- ✅ Importar PDFs/textos gratuitos relevantes (≥30% relevância)
- ✅ Criar links many-to-many (article ↔ score item)

### 2. Geração de Conteúdo Clínico (Português-BR)
- ✅ **clinicalRelevance** (200-400 palavras)
  - Explicação técnica para médicos
  - Base científica e mecanismos
  - Importância clínica

- ✅ **patientExplanation** (100-200 palavras)
  - Linguagem simples e acessível
  - Sem jargões médicos
  - Foco em compreensão do paciente

- ✅ **conduct** (150-300 palavras)
  - Orientações práticas baseadas em evidências
  - Quando recomendar
  - Como implementar
  - Monitoramento e contraindicações

### 3. Atualização Automática
- ✅ `lastReview` → NOW() (via hook GORM)

---

## 🚀 Batches Planejados

### Batch 1 - Alimentação (5 items) ✅ CONCLUÍDO
**Agente:** aa22287
**Iniciado:** 26/01/2026 00:16 UTC
**Concluído:** 26/01/2026 00:26 UTC
**Status:** Conteúdo gerado, documentado em DIETARY-ITEMS-CLINICAL-CONTENT-REPORT.md

**Items:**
1. Estratégia macro atual
2. Livre
3. Low carb
4. Vegetariana
5. Vegana

**Resultado:**
- ✅ 17.491 palavras de conteúdo clínico gerado
- ✅ Baseado em 241 lectures MFI + literatura 2026
- ✅ Script Python criado com todo o conteúdo
- ⚠️ Aguardando aplicação no banco via API

---

### Batch 2 - Exames Laboratoriais (5 items) 🔄 EM ANDAMENTO
**Agente:** a42e1aa
**Iniciado:** 26/01/2026 00:17 UTC
**Status:** Processando

**Items:**
- Colesterol Total
- Colesterol não-HDL
- Hemoglobina
- Glicose
- Outros exames lab

**Progresso:**
- ✅ Login autenticado
- 🔄 Buscando evidências
- ⏳ Gerando textos clínicos

---

### Batch 3 - Vitaminas (5 items) 🔄 EM ANDAMENTO
**Agente:** aee6cb8
**Iniciado:** 26/01/2026 00:18 UTC
**Status:** Processando

**Items:**
- 25-hidroxivitamina D
- Vitamina A
- Vitamina B1 (Tiamina)
- Vitamina B2 (Riboflavina)
- Vitamina B12 (Cobalamina)

**Progresso:**
- ✅ Login autenticado
- 🔄 Buscando evidências

---

### Batch 4 - Sono (10 items) 🔄 EM ANDAMENTO
**Agente:** a9c31d0
**Iniciado:** 26/01/2026 00:35 UTC
**Status:** Iniciando

---

### Batch 5 - Cognição (10 items) 🔄 EM ANDAMENTO
**Agente:** a584e16
**Iniciado:** 26/01/2026 00:35 UTC
**Status:** Iniciando

---

### Batch 6 - Stress (9 items) 🔄 EM ANDAMENTO
**Agente:** a9c31d0 (novo)
**Iniciado:** 26/01/2026 00:36 UTC
**Status:** Iniciando
**Nota:** Grupo pequeno - processar todos os 9 items

---

### Batch 7 - Movimento (10 items) 🔄 EM ANDAMENTO
**Agente:** (novo)
**Iniciado:** 26/01/2026 00:36 UTC
**Status:** Iniciando

---

### Batch 8 - Composição Corporal (10 items) 🔄 EM ANDAMENTO
**Agente:** (novo)
**Iniciado:** 26/01/2026 00:36 UTC
**Status:** Iniciando

---

### Batches 9+ - Demais Grupos (~2.250 items) ⏸️ PLANEJADO

**Grupos principais restantes:**
- Exames laboratoriais (~880 items restantes)
- Histórico de doenças (~393 items)
- Composição corporal (~152 items restantes)
- Alimentação (~137 items restantes)
- Sono (~128 items restantes)
- Cognição (~68 items restantes)
- Outros grupos (~492 items)

---

## 📈 Estatísticas em Tempo Real

### Items Processados por Categoria

| Categoria | Total | Processados | Pendentes | % Concluído |
|-----------|-------|-------------|-----------|-------------|
| Alimentação | ~30 | 1 | 29 | 3.3% |
| Exames Lab | ~400 | 0 | 400 | 0% |
| Genética | ~50 | 0 | 50 | 0% |
| Hábitos | ~40 | 0 | 40 | 0% |
| Sintomas | ~80 | 0 | 80 | 0% |
| Histórico | ~100 | 0 | 100 | 0% |
| Medicações | ~50 | 0 | 50 | 0% |
| Outros | ~22 | 0 | 22 | 0% |
| **TOTAL** | **772** | **1** | **771** | **0.1%** |

### Articles por Score Item

| Métrica | Valor |
|---------|-------|
| Items com 0 articles | 771 |
| Items com 1-5 articles | 1 |
| Items com 6-10 articles | 0 |
| Items com 11+ articles | 0 |
| **Média de articles/item** | **0.001** |

### Campos Clínicos Preenchidos

| Campo | Preenchidos | Pendentes | % |
|-------|-------------|-----------|---|
| clinicalRelevance | 1 | 771 | 0.1% |
| patientExplanation | 1 | 771 | 0.1% |
| conduct | 1 | 771 | 0.1% |

---

## ⚙️ Tecnologias e Ferramentas

### Agentes Utilizados
- **ai-engineer**: Análise de conteúdo + geração de textos clínicos
- **general-purpose**: Busca online e coleta de PDFs
- **WebSearch**: Google Scholar + PubMed
- **Grep**: Busca full-text nos articles existentes

### APIs Utilizadas
- **Plenya API:** `http://localhost:3001/api/v1`
  - PUT `/score-items/{id}` - Atualizar campos clínicos
  - POST `/articles/{id}/score-items` - Criar links
  - GET `/articles` - Buscar articles existentes

### Banco de Dados
- **PostgreSQL 17**
- **Tabelas:** score_items, articles, article_score_items
- **Hooks GORM:** BeforeUpdate (lastReview automático)

---

## 🎓 Qualidade do Conteúdo Gerado

### Critérios de Validação

**clinicalRelevance:**
- ✅ Base em evidências científicas
- ✅ Linguagem técnica adequada
- ✅ Referências a lectures MFI quando aplicável
- ✅ 200-400 palavras

**patientExplanation:**
- ✅ Linguagem acessível (Flesch-Kincaid ≥ 60)
- ✅ Sem jargões médicos
- ✅ Foco em compreensão
- ✅ 100-200 palavras

**conduct:**
- ✅ Orientações práticas e aplicáveis
- ✅ Baseadas em medicina funcional integrativa
- ✅ Inclusão de quando/como/monitoramento
- ✅ 150-300 palavras

---

## 📚 Fontes de Evidências

### Articles Existentes no Sistema

| Tipo | Quantidade | Journal |
|------|------------|---------|
| lecture | 241 | Pos Graduacao MFI |
| research_article | 6 | Diversos |
| **TOTAL** | **247** | - |

### Busca Online (Critérios)

**Importar como article:**
- ✅ Gratuito E completo E relevância ≥30%
- ✅ OU relevância ≥80% (mesmo que não gratuito/completo)

**Fontes prioritárias:**
- PubMed / PubMed Central (PMC)
- Google Scholar
- Journals de medicina funcional integrativa
- Guidelines de sociedades médicas

---

## ⏱️ Estimativas de Tempo

### Por Batch (5-10 items)
- **Busca de evidências:** 5-10 min
- **Geração de textos (IA):** 5-10 min
- **Salvamento + links:** 2-5 min
- **TOTAL por batch:** 15-30 min

### Tempo Total Estimado
- **772 items ÷ 10 items/batch** = 77 batches
- **77 batches × 20 min/batch** = 1.540 minutos
- **= 25.7 horas** (com 1 agente)
- **= 12.8 horas** (com 2 agentes paralelos)

### Processamento 24/7
- **Com 1 agente:** ~2 dias
- **Com 2 agentes:** ~1 dia
- **Com pausas/revisões:** ~3-5 dias

---

## 🐛 Desafios e Soluções

### Desafio 1: Context Window Limits
**Problema:** Processar 772 items estoura context
**Solução:** ✅ Batches de 5-10 items + agentes especializados

### Desafio 2: Rate Limits (APIs)
**Problema:** Muitas requisições simultâneas
**Solução:** ✅ Throttling + agentes sequenciais

### Desafio 3: Qualidade Variável
**Problema:** Alguns items têm poucas evidências
**Solução:** ✅ Ser conservador mas informativo + busca online

### Desafio 4: Tempo de Execução
**Problema:** 772 items levam muito tempo
**Solução:** ✅ Processamento em background + monitoramento

---

## 📝 Próximos Passos

### Imediato (Próximas Horas)
1. ✅ Concluir Batch 1 (5 items Alimentação)
2. ⏳ Validar qualidade dos textos gerados
3. ⏳ Iniciar Batch 2 (10 items Alimentação)
4. ⏳ Ajustar prompts se necessário

### Curto Prazo (Próximos Dias)
1. Processar todos items de Alimentação (~30 total)
2. Processar items de Exames Laboratoriais principais (~50)
3. Processar items de Genética (~50)
4. Revisar amostra de 10% dos items processados

### Médio Prazo (Próxima Semana)
1. Completar todos os 772 items
2. Revisão médica de amostra (50 items aleatórios)
3. Ajustes baseados em feedback
4. Documentação final

---

## 📊 Métricas de Sucesso

### KPIs
- **Cobertura:** 100% dos 772 items com os 3 campos preenchidos
- **Qualidade:** ≥ 90% dos textos aprovados em revisão médica
- **Links:** Média ≥ 2 articles por item
- **Evidências:** ≥ 70% dos items com base em lectures MFI
- **Tempo:** Completar em ≤ 5 dias corridos

### Métricas Atuais
- ✅ Cobertura: 0.1% (1/772)
- ⏳ Qualidade: A avaliar após Batch 1
- ⏳ Links: 0 (average)
- ⏳ Evidências: A mapear
- ⏳ Tempo: Dia 0 de 5

---

## 🔗 Links Úteis

- **Banco de dados:** `docker compose exec db psql -U plenya_user -d plenya_db`
- **API:** http://localhost:3001/api/v1
- **Frontend:** http://localhost:3000/scores
- **Agente atual:** `/tmp/claude/-home-user-plenya/tasks/aa22287.output`

---

## 📌 Notas Importantes

1. **Revisão humana OBRIGATÓRIA** após processamento
2. **lastReview automático** garante rastreabilidade
3. **Links many-to-many** preservam relações article ↔ item
4. **Textos em português-BR** para sistema brasileiro
5. **Base em evidências** (lectures MFI + literatura científica)

---

**Última atualização:** 26/01/2026 00:20 UTC
**Atualizado por:** Sistema automatizado
**Próxima atualização:** Após conclusão do Batch 1

---

*Plenya EMR - Sistema de Prontuário Eletrônico Baseado em Evidências*
*Versão: 2026.01*
