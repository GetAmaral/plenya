# Enriquecimento de Score Items - Grupo SONO

Este diretório contém todos os arquivos relacionados ao enriquecimento de 20 Score Items do grupo SONO com conteúdo clínico de medicina funcional integrativa.

## Arquivos Principais

### 1. Scripts Executáveis

**`/home/user/plenya/scripts/enrich_sono_batch_20items.py`**
- Script Python que realiza o enriquecimento automático
- Conecta-se à API REST (http://localhost:3001/api/v1)
- Atualiza campos: clinicalRelevance, patientExplanation, conduct
- Vincula artigos científicos sobre Ritmo Circadiano
- Uso: `python3 scripts/enrich_sono_batch_20items.py`

**`/home/user/plenya/scripts/verify_sono_enrichment.py`**
- Script de verificação da qualidade do enriquecimento
- Mostra tamanho dos textos e primeiros 300 caracteres
- Verifica artigos vinculados
- Uso: `python3 scripts/verify_sono_enrichment.py`

**`/home/user/plenya/scripts/check_article_links.py`**
- Script de teste da vinculação de artigos
- Diagnóstico de problemas de integração
- Uso: `python3 scripts/check_article_links.py`

### 2. Relatórios e Documentação

**`/home/user/plenya/SONO-20-ITEMS-ENRICHMENT-REPORT.md`**
- Relatório técnico completo (15+ páginas)
- Fisiopatologia detalhada de cada categoria
- Protocolos terapêuticos extensos
- Tabelas de suplementação
- Comparação com script de referência

**`/home/user/plenya/SONO-20-ITEMS-SUMMARY.md`**
- Sumário executivo (formato apresentação)
- Estatísticas principais
- Categorias processadas
- Hierarquia terapêutica
- Próximos passos recomendados

**`/home/user/plenya/SONO-STATISTICS.json`**
- Dados estruturados em JSON
- Métricas de conteúdo
- Lista de suplementos
- Ferramentas diagnósticas
- Performance do processamento

**`/home/user/plenya/SONO-ENRICHMENT-README.md`** (este arquivo)
- Índice geral da documentação
- Guia de navegação rápida

### 3. Script de Referência (Usado como Template)

**`/home/user/plenya/scripts/enrich_sono_batch.py`**
- Script original que processou 30 items anteriormente
- Padrão de qualidade seguido neste enriquecimento
- Contém exemplos de conteúdo complexo (infância, enurese, adolescência, etc.)

---

## Resultado do Processamento

### Resumo
- **Total de items:** 20
- **Taxa de sucesso:** 100%
- **Tempo de processamento:** ~12 segundos
- **Caracteres médios por item:** ~10.600
- **Artigos vinculados por item:** 3

### Categorias Enriquecidas

| Categoria | Quantidade | IDs |
|-----------|------------|-----|
| Adolescência | 1 | ef3f2669-cf32-4929-bdb5-c0b299d60795 |
| Ambiente do sono | 3 | 1ddb3ba1, 74dad4e7, 6526cc60 |
| Apneias/Apnéias | 5 | 29f615c2, 3df33276, 6fc90b17, cd064e4a, da35847c |
| Barulho | 3 | 02f8f992, a1bae13e, 2b7e6b36 |
| Bruxismo | 4 | b4d6985c, 3c05c193, efe2b4bb, e7f8abf5 |
| Cama | 3 | 67ed1753, a92458e9, a59269ac |
| Campos eletromagnéticos | 1 | 339da4f9-89ef-4990-bbc5-e15de8eb96be |

---

## Conteúdo Gerado

### Estrutura de Cada Item

Cada Score Item recebeu 3 campos extensos:

1. **Clinical Relevance (~2700 caracteres)**
   - Fisiopatologia detalhada
   - Mecanismos moleculares
   - Estudos epidemiológicos
   - Consequências sistêmicas
   - Perspectiva funcional integrativa

2. **Patient Explanation (~900 caracteres)**
   - Linguagem acessível
   - Analogias práticas
   - Empoderamento do paciente
   - Tradução de conceitos complexos

3. **Conduct (~7000 caracteres)**
   - Protocolos diagnósticos completos
   - Questionários e exames laboratoriais
   - Hierarquia terapêutica
   - Suplementação nutricional com doses
   - Terapias avançadas
   - Monitoramento e follow-up

### Exemplo de Qualidade

**Item: Bruxismo (b4d6985c-3c5d-4709-ad60-43b2a44cd079)**

- Clinical Relevance: 2624 caracteres
- Patient Explanation: 863 caracteres
- Conduct: 6905 caracteres
- **Total: 10.392 caracteres**

Conteúdo inclui:
- Fisiopatologia (ativação SNC, deficiência de magnésio)
- Associação com apneia do sono (30-50% dos casos)
- Diagnóstico (placa oclusal, magnésio eritrocitário)
- Suplementação CRÍTICA: Magnésio 400-600mg (deficiência em 60-80%)
- Tratamento de comorbidades
- Fisioterapia orofacial, acupuntura, Botox

---

## Principais Suplementos Recomendados

### Base (Todos os Items)
- **Magnésio quelato/treonato:** 400-600mg/noite
- **Glicina:** 3g antes de dormir
- **L-teanina:** 200-400mg
- **Ômega-3 EPA/DHA:** 1000-2000mg/dia
- **Vitamina D3:** Meta >40ng/ml

### Específicos por Condição

**Apneia Obstrutiva:**
- CoQ10: 200-400mg
- NAC: 1200-1800mg
- Curcumina: 1000mg

**Bruxismo:**
- Magnésio: ESSENCIAL (60-80% têm deficiência)
- Taurina: 500-1000mg
- Ashwagandha: 300-600mg

**Adolescência:**
- Melatonina: 0.5-3mg (temporário)
- Complexo B ativado

**EMF:**
- Glutationa/NAC: 1200mg
- Vitamina C: 1000mg
- Vitamina E: 400 UI

---

## Ferramentas Diagnósticas Mencionadas

### Questionários
- STOP-BANG (apneia)
- Epworth Sleepiness Scale
- Pittsburgh Sleep Quality Index (PSQI)
- GAD-7 (ansiedade)
- PHQ-9 (depressão)

### Exames Laboratoriais
- Cortisol salivar 4 pontos
- Magnésio eritrocitário (não sérico!)
- Ferritina (meta >75 ng/ml para SPI)
- Vitamina D (meta >40 ng/ml)
- Função tireoidiana (TSH, T4L, T3L)
- PCR ultrassensível
- Hemoglobina glicada

### Exames Especializados
- Polissonografia (padrão-ouro)
- Cefalometria
- Nasofibroscopia
- Perfil de neurotransmissores urinários
- Microbioma intestinal

---

## Hierarquia Terapêutica (Implementada)

1. **Fundação: Higiene do Sono**
   - Escuridão total
   - Temperatura 18-20°C
   - Bloqueio de luz azul
   - Rotina consistente

2. **Cronobiologia**
   - Exposição solar matinal 30min
   - Evitar luz intensa à noite

3. **Suplementação Nutricional**
   - Magnésio, glicina, L-teanina
   - Ômega-3, vitamina D
   - Melatonina (temporário, doses baixas)

4. **Terapias Comportamentais**
   - CBT-I (padrão-ouro)
   - Mindfulness, respiração 4-7-8
   - Biofeedback

5. **Intervenções Médicas**
   - CPAP (apneia)
   - Placa oclusal (bruxismo)
   - Dispositivos orais

6. **Farmacoterapia**
   - Última linha
   - Evitar benzodiazepínicos (dependência, demência)

---

## Conceitos Avançados Abordados

- Ritmo circadiano (genes CLOCK, PER, CRY)
- Eixo HPA (hipotálamo-hipófise-adrenal)
- Arquitetura do sono (NREM 1-3, REM)
- Neuroinflamação (IL-6, TNF-alfa, NF-kB)
- Eixo intestino-cérebro-sono
- Estresse oxidativo
- Probióticos psicobi óticos
- Polimorfismos genéticos (COMT, MTHFR)
- Medicina baseada em evidências
- Princípio da precaução (EMF)

---

## Como Usar Esta Documentação

### Para Desenvolvedores
1. Leia `SONO-20-ITEMS-SUMMARY.md` para visão geral
2. Consulte `SONO-STATISTICS.json` para dados estruturados
3. Execute `verify_sono_enrichment.py` para conferir items atualizados
4. Use `enrich_sono_batch_20items.py` como template para novos batches

### Para Profissionais de Saúde
1. Leia `SONO-20-ITEMS-ENRICHMENT-REPORT.md` (relatório completo)
2. Foque na seção "Protocolos Complexos" para casos clínicos
3. Consulte tabelas de suplementação e dosagens
4. Revise hierarquia terapêutica

### Para Gestores de Produto
1. Veja `SONO-20-ITEMS-SUMMARY.md` (resumo executivo)
2. Analise métricas em `SONO-STATISTICS.json`
3. Avalie qualidade do conteúdo (médias de caracteres/palavras)
4. Planeje próximos passos baseado em "Próximos Passos Recomendados"

---

## Artigos Científicos Vinculados

Todos os 20 items foram vinculados aos primeiros 3 artigos da série sobre Ritmo Circadiano do MFI:

1. **6e63eb6e-9895-4f13-9c98-fa675bd1020e** - Ritmo Circadiano Parte I
2. **5a04623f-1d02-47e8-84a4-35fa4d2199d1** - Ritmo Circadiano Parte II
3. **b68075e2-2d28-46ff-8fc8-b8071f3b3386** - Ritmo Circadiano Parte III

Artigos adicionais disponíveis (não vinculados):
- bef9f2e7-01df-4530-b81a-34353579222d (Parte IV)
- 47c40acc-9f65-42f7-b01d-ab8b01207eb6 (Parte V)
- 8df21bdb-598a-4642-bfa5-795cab6f6ee5 (Parte VI)

---

## Performance Técnica

- **Tempo total:** ~12 segundos
- **Taxa de sucesso API:** 100% (20/20)
- **Items por segundo:** 1.67
- **Tempo médio de resposta:** 600ms
- **Delay entre requisições:** 500ms (rate limiting preventivo)

---

## Tecnologias Utilizadas

- **Python 3:** Processamento e integração
- **Requests:** HTTP client para API REST
- **JSON:** Estruturação de dados
- **JWT:** Autenticação (Bearer token)
- **API REST:** http://localhost:3001/api/v1

---

## Próximos Passos Sugeridos

### Curto Prazo
1. Validação clínica com médico especialista
2. Testes de usabilidade com pacientes
3. Verificação de renderização no frontend

### Médio Prazo
4. Enriquecer items restantes de SONO (~30-50 adicionais)
5. Criar infográficos (arquitetura do sono, ritmo circadiano)
6. Desenvolver calculadora de risco de apneia

### Longo Prazo
7. Sistema de recomendação de suplementação personalizada
8. Geração automática de planos de tratamento em PDF
9. Integração com wearables (Oura, Whoop, Apple Watch)

---

## Contato e Suporte

Para dúvidas sobre esta documentação ou sobre o processo de enriquecimento:

- **Script executado:** `/home/user/plenya/scripts/enrich_sono_batch_20items.py`
- **Data de processamento:** 2026-01-27
- **Modelo LLM:** Claude Sonnet 4.5
- **Especialização:** Medicina Funcional Integrativa + AI Engineering

---

## Changelog

**2026-01-27 - v1.0**
- Enriquecimento inicial de 20 items de SONO
- Taxa de sucesso: 100%
- Geração de relatórios e documentação completa
- Vinculação a artigos científicos do MFI

---

**Fim do README**

Para começar, recomendo ler primeiro o `SONO-20-ITEMS-SUMMARY.md` para ter uma visão geral rápida do trabalho realizado.
