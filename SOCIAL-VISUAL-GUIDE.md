# Batch SOCIAL - Guia Visual

Representação visual do sistema de enriquecimento SOCIAL.

---

## 🎯 Visão Geral em 1 Minuto

```
┌─────────────────────────────────────────────────────────────┐
│                   BATCH SOCIAL ENRICHMENT                   │
│                                                             │
│  30 Items SOCIAL  →  Claude Sonnet 4  →  Items Enriquecidos│
│                                                             │
│  Tempo: 25 min    |  Taxa: 100%      |  Output: 4 campos   │
└─────────────────────────────────────────────────────────────┘
```

---

## 📊 Arquitetura do Sistema

```
┌───────────────────┐
│   INPUT: 30 UUIDs │
└─────────┬─────────┘
          │
          ▼
┌─────────────────────────┐
│  1. Login (JWT)         │
│     ↓                   │
│  2. For each item:      │
│     • Fetch item        │
│     • Generate (Claude) │
│     • Parse JSON        │
│     • Update item       │
│     ↓                   │
│  3. Generate report     │
└─────────┬───────────────┘
          │
          ▼
┌───────────────────────────┐
│  OUTPUT: 30 Items + 4     │
│  campos clínicos cada     │
└───────────────────────────┘
```

---

## 🗂️ Estrutura de Documentação

```
SOCIAL-INDEX.md ◀── Você está aqui (índice)
    │
    ├─── SOCIAL-README.md ◀── Comece aqui (overview)
    │
    ├─── SOCIAL-QUICK-START.md ◀── Execução rápida (3 passos)
    │
    ├─── SOCIAL-BATCH-EXECUTIVE-SUMMARY.md ◀── Sumário executivo
    │
    ├─── SOCIAL-ENRICHMENT-METHODOLOGY.md ◀── Metodologia detalhada
    │
    ├─── SOCIAL-SCIENTIFIC-REFERENCES.md ◀── 50+ referências
    │
    ├─── SOCIAL-EXPECTED-OUTPUT-EXAMPLES.md ◀── Exemplos de qualidade
    │
    └─── SOCIAL-POST-EXECUTION-VALIDATION.md ◀── Checklist validação
```

---

## 🔄 Fluxo de Trabalho Visual

```
╔════════════════════════════════════════════════════════════╗
║                    WORKFLOW COMPLETO                       ║
╚════════════════════════════════════════════════════════════╝

┌─────────────┐
│ 1. PRÉ-REQ  │  • API rodando
│   (5 min)   │  • ANTHROPIC_API_KEY
│             │  • Dependências Python
└──────┬──────┘
       │
       ▼
┌─────────────┐
│ 2. EXECUÇÃO │  ./execute_social_batch.sh
│   (25 min)  │  [████████████████] 30/30
│             │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│ 3. VALIDAÇÃO│  • Técnica (queries SQL)
│   (30 min)  │  • Qualidade (amostragem)
│             │  • API/Frontend (testes)
└──────┬──────┘
       │
       ▼
┌─────────────┐
│ 4. REVISÃO  │  • Médico especialista
│   (2-4h)    │  • Ajustes se necessário
│             │  • Aprovação final
└──────┬──────┘
       │
       ▼
┌─────────────┐
│ 5. DEPLOY   │  ✅ Batch SOCIAL completo
│   (30 min)  │  🚀 Produção
└─────────────┘
```

---

## 📁 Items SOCIAL por Categoria

```
╔═══════════════════════════════════════════════════════════════╗
║                  30 ITEMS SOCIAL (7 CATEGORIAS)               ║
╚═══════════════════════════════════════════════════════════════╝

┌──────────────────────────────────────────────────────────────┐
│ 🔊 AMBIENTE SONORO (4-5 items)                               │
├──────────────────────────────────────────────────────────────┤
│ • Nível de ruído ambiente                                    │
│ • Exposição a poluição sonora                                │
│ • Ruído noturno                                              │
│ → Foco: Cortisol, HTA, sono                                  │
└──────────────────────────────────────────────────────────────┘

┌──────────────────────────────────────────────────────────────┐
│ 🏠 CONDIÇÕES DE MORADIA (5-6 items)                          │
├──────────────────────────────────────────────────────────────┤
│ • Presença de mofo/umidade                                   │
│ • Qualidade do ar interno                                    │
│ • Uso de produtos químicos                                   │
│ → Foco: Micotoxinas, SIRS, toxinas domésticas               │
└──────────────────────────────────────────────────────────────┘

┌──────────────────────────────────────────────────────────────┐
│ 🚶 ESPAÇO PARA MOVIMENTO (3-4 items)                         │
├──────────────────────────────────────────────────────────────┤
│ • Acesso a áreas de caminhada                                │
│ • Espaço para atividade física                               │
│ → Foco: Sedentarismo ambiental, NEAT                         │
└──────────────────────────────────────────────────────────────┘

┌──────────────────────────────────────────────────────────────┐
│ 🌫️ EXPOSIÇÃO AMBIENTAL EXTERNA (5-6 items)                  │
├──────────────────────────────────────────────────────────────┤
│ • Poluição do ar (PM2.5)                                     │
│ • Exposição a metais pesados                                 │
│ • Proximidade a indústrias                                   │
│ → Foco: Estresse oxidativo, inflamação                       │
└──────────────────────────────────────────────────────────────┘

┌──────────────────────────────────────────────────────────────┐
│ 🎨 HOBBIES E LAZER (4-5 items)                               │
├──────────────────────────────────────────────────────────────┤
│ • Horas/semana de hobbies                                    │
│ • Atividades de lazer                                        │
│ • Conexão social através de hobbies                          │
│ → Foco: Longevidade, isolamento social                       │
└──────────────────────────────────────────────────────────────┘

┌──────────────────────────────────────────────────────────────┐
│ ☀️ LUMINOSIDADE NATURAL (3-4 items)                          │
├──────────────────────────────────────────────────────────────┤
│ • Exposição à luz solar                                      │
│ • Horas de luz natural/dia                                   │
│ → Foco: Ritmo circadiano, vitamina D                         │
└──────────────────────────────────────────────────────────────┘

┌──────────────────────────────────────────────────────────────┐
│ 💼 PROFISSÕES (4-5 items)                                    │
├──────────────────────────────────────────────────────────────┤
│ • Tipo de trabalho                                           │
│ • Trabalho noturno                                           │
│ • Exposições ocupacionais                                    │
│ • Nível de estresse/burnout                                  │
│ → Foco: Dessincronização circadiana, exposições             │
└──────────────────────────────────────────────────────────────┘
```

---

## 📝 Estrutura de Conteúdo por Item

```
╔═══════════════════════════════════════════════════════════════╗
║                 ITEM SOCIAL (ANTES vs DEPOIS)                 ║
╚═══════════════════════════════════════════════════════════════╝

ANTES (Vazio):
┌─────────────────────────────────────────────────────────────┐
│ ID: c84412f7-393f-41d0-8bd7-0a28824dbeb0                    │
│ Nome: Nível de ruído no ambiente                            │
│ Descrição: Avalia exposição a poluição sonora               │
│                                                             │
│ clinical_relevance: null          ❌                         │
│ interpretation_guidelines: null   ❌                         │
│ actionable_insights: null         ❌                         │
│ red_flags: null                   ❌                         │
└─────────────────────────────────────────────────────────────┘

         ↓ ↓ ↓  ENRIQUECIMENTO (Claude Sonnet 4)  ↓ ↓ ↓

DEPOIS (Enriquecido):
┌─────────────────────────────────────────────────────────────┐
│ ID: c84412f7-393f-41d0-8bd7-0a28824dbeb0                    │
│ Nome: Nível de ruído no ambiente                            │
│ Descrição: Avalia exposição a poluição sonora               │
│                                                             │
│ ✅ clinical_relevance (1,200 chars)                         │
│    "A exposição crônica a ruído ambiental representa        │
│     um estressor fisiológico significativo... [mecanismos   │
│     fisiopatológicos + evidências epidemiológicas]"         │
│                                                             │
│ ✅ interpretation_guidelines (1,500 chars)                  │
│    "PADRÃO 1: Ruído >70dB - Significado: Ativação eixo     │
│     HPA... Investigar: Cortisol, MAPA 24h..."              │
│                                                             │
│ ✅ actionable_insights (2,000 chars)                        │
│    "1. SE ruído >85dB: Protetores auriculares...           │
│     2. SE cortisol elevado: Ashwagandha 600mg..."          │
│                                                             │
│ ✅ red_flags (900 chars)                                    │
│    "🚩 RED FLAG 1: Ruído >85dB sem EPI + perda auditiva    │
│     - Risco: PAIR irreversível - Ação: Audiometria..."     │
└─────────────────────────────────────────────────────────────┘
```

---

## 🎯 4 Campos Clínicos (Detalhado)

```
╔═══════════════════════════════════════════════════════════════╗
║              ESTRUTURA DOS 4 CAMPOS CLÍNICOS                  ║
╚═══════════════════════════════════════════════════════════════╝

┌───────────────────────────────────────────────────────────────┐
│ 1️⃣  CLINICAL RELEVANCE (800-1500 chars)                      │
├───────────────────────────────────────────────────────────────┤
│                                                               │
│  POR QUÊ IMPORTA?                                            │
│                                                               │
│  ✓ Mecanismos fisiopatológicos específicos                   │
│    Ex: "Ruído → Eixo HPA → Cortisol ↑ → Resistência         │
│         insulínica"                                           │
│                                                               │
│  ✓ Evidências epidemiológicas                                │
│    Ex: "Cada 10dB aumento = 8% ↑ risco HTA (Münzel 2021)"   │
│                                                               │
│  ✓ Impacto em sistemas                                       │
│    Ex: "Cardiovascular, endócrino, neurológico"              │
│                                                               │
│  ✓ Conexão com doenças crônicas                              │
│    Ex: "HTA, DM2, síndrome metabólica"                        │
│                                                               │
└───────────────────────────────────────────────────────────────┘

┌───────────────────────────────────────────────────────────────┐
│ 2️⃣  INTERPRETATION GUIDELINES (1000-2000 chars)              │
├───────────────────────────────────────────────────────────────┤
│                                                               │
│  COMO INTERPRETAR?                                           │
│                                                               │
│  PADRÃO 1: [Descrição resposta]                              │
│  ├─ Significado clínico: [O que significa]                   │
│  ├─ Sistemas comprometidos: [Quais sistemas]                 │
│  ├─ Investigar: [Exames complementares]                      │
│  └─ DD: [Diagnósticos diferenciais]                          │
│                                                               │
│  PADRÃO 2: [Outra resposta]                                  │
│  ├─ Significado clínico: ...                                 │
│  └─ ...                                                       │
│                                                               │
│  PADRÃO 3: ...                                               │
│                                                               │
└───────────────────────────────────────────────────────────────┘

┌───────────────────────────────────────────────────────────────┐
│ 3️⃣  ACTIONABLE INSIGHTS (1500-2500 chars)                    │
├───────────────────────────────────────────────────────────────┤
│                                                               │
│  O QUE FAZER?                                                │
│                                                               │
│  1. SE [condição]:                                           │
│     • IMEDIATO: [Ação urgente]                               │
│     • MÉDIO PRAZO: [Intervenção programada]                  │
│     • LONGO PRAZO: [Mudança sustentável]                     │
│     • Monitorar: [Métricas de acompanhamento]                │
│                                                               │
│  2. SE [outra condição]:                                     │
│     • AMBIENTAL: [Mudança ambiente]                          │
│     • NUTRACÊUTICOS: [Suplementação]                         │
│     • FARMACOLÓGICO: [Se necessário]                         │
│                                                               │
│  3-8. [Mais intervenções...]                                 │
│                                                               │
│  Reavaliação: [Timeline]                                     │
│     • Métricas: [O que medir]                                │
│                                                               │
└───────────────────────────────────────────────────────────────┘

┌───────────────────────────────────────────────────────────────┐
│ 4️⃣  RED FLAGS (600-1200 chars)                               │
├───────────────────────────────────────────────────────────────┤
│                                                               │
│  QUANDO AGIR IMEDIATAMENTE?                                  │
│                                                               │
│  🚩 RED FLAG 1: [Situação crítica]                           │
│     - Risco: [Consequência grave]                            │
│     - Ação: [Intervenção imediata]                           │
│                                                               │
│  🚩 RED FLAG 2: [Outra situação]                             │
│     - Risco: ...                                             │
│     - Ação: ...                                              │
│                                                               │
│  🚩 RED FLAGS 3-5: [Mais alertas...]                         │
│                                                               │
└───────────────────────────────────────────────────────────────┘
```

---

## 🔬 Base Científica Visual

```
╔═══════════════════════════════════════════════════════════════╗
║            PIRÂMIDE DE EVIDÊNCIAS (50+ Referências)           ║
╚═══════════════════════════════════════════════════════════════╝

                    ┌───────────────┐
                    │  Guidelines   │
                    │  Internacionais│
                    │  (WHO, IARC,  │
                    │   EPA, CDC)   │
                    └───────┬───────┘
                            │
                ┌───────────┴───────────┐
                │  Meta-análises &      │
                │  Revisões Sistemáticas│
                │  (Cochrane, Lancet    │
                │   Commissions)        │
                └───────────┬───────────┘
                            │
            ┌───────────────┴───────────────┐
            │  Estudos de Coorte Longos     │
            │  (Framingham, Nurses' Health, │
            │   Blue Zones)                 │
            └───────────────┬───────────────┘
                            │
        ┌───────────────────┴───────────────────┐
        │  Estudos Mecanísticos                 │
        │  (Fisiopatologia, Biomarcadores)      │
        └───────────────────────────────────────┘

        Total: 50+ referências de alta qualidade
```

---

## 📊 Métricas de Sucesso

```
╔═══════════════════════════════════════════════════════════════╗
║                    DASHBOARD DE MÉTRICAS                      ║
╚═══════════════════════════════════════════════════════════════╝

┌──────────────────────────┬──────────────────────────────────┐
│  MÉTRICA                 │  VALOR ESPERADO                  │
├──────────────────────────┼──────────────────────────────────┤
│  Items processados       │  30/30 (100%)        ████████████│
│  Taxa de sucesso         │  100%                ████████████│
│  Tempo total             │  20-25 min           ████████░░░░│
│  Tempo por item          │  45-60s              ████████░░░░│
│                          │                                  │
│  Clinical Relevance      │  800-1500 chars      ████████░░░│
│  Interpretation          │  1000-2000 chars     ████████████│
│  Actionable Insights     │  1500-2500 chars     ████████████│
│  Red Flags               │  600-1200 chars      ████████░░░│
│                          │                                  │
│  Total por item          │  ~4000-6500 chars    ████████████│
│  Total batch             │  ~120K-195K chars    ████████████│
└──────────────────────────┴──────────────────────────────────┘
```

---

## ✅ Checklist Visual de Validação

```
╔═══════════════════════════════════════════════════════════════╗
║              CHECKLIST PÓS-EXECUÇÃO (Rápido)                  ║
╚═══════════════════════════════════════════════════════════════╝

TÉCNICO
  ☐ Relatório JSON gerado (SOCIAL-BATCH-REPORT.json)
  ☐ 30 items em array "success"
  ☐ 0 items em array "failed"
  ☐ Query banco: 30 items enriquecidos
  ☐ Campos não-null verificados
  ☐ Teste API retorna JSON válido
  ☐ Frontend exibe campos corretamente

CONTEÚDO
  ☐ Clinical Relevance: mecanismos + evidências
  ☐ Interpretation: padrões práticos
  ☐ Actionable: 5-8 intervenções
  ☐ Red Flags: 3-5 alertas críticos
  ☐ Referências científicas citadas
  ☐ Linguagem técnica mas acessível

QUALIDADE CLÍNICA
  ☐ Amostra aleatória revisada (5 items)
  ☐ Acurácia científica verificada
  ☐ Intervenções práticas e seguras
  ☐ Revisão médica agendada/completa
  ☐ Feedback incorporado

APROVAÇÃO FINAL
  ☐ 100% items validados
  ☐ Aprovação médica recebida
  ☐ Pronto para produção
```

---

## 🚀 Status Visual

```
╔═══════════════════════════════════════════════════════════════╗
║                    STATUS ATUAL DO PROJETO                    ║
╚═══════════════════════════════════════════════════════════════╝

PLANEJAMENTO         [████████████] 100% ✅ COMPLETO
DOCUMENTAÇÃO         [████████████] 100% ✅ COMPLETO
IMPLEMENTAÇÃO        [████████████] 100% ✅ COMPLETO
EXECUÇÃO             [░░░░░░░░░░░░]   0% ⏳ PENDENTE
VALIDAÇÃO            [░░░░░░░░░░░░]   0% ⏳ PENDENTE
REVISÃO MÉDICA       [░░░░░░░░░░░░]   0% ⏳ PENDENTE
DEPLOY               [░░░░░░░░░░░░]   0% ⏳ PENDENTE

───────────────────────────────────────────────────────────────
PROGRESSO GERAL:     [████░░░░░░░░]  33%
───────────────────────────────────────────────────────────────

PRÓXIMA AÇÃO: ./execute_social_batch.sh
```

---

## 🎯 Call-to-Action

```
╔═══════════════════════════════════════════════════════════════╗
║                                                               ║
║                 🚀 PRONTO PARA EXECUTAR! 🚀                  ║
║                                                               ║
║  1. Configurar: export ANTHROPIC_API_KEY='...'               ║
║  2. Executar:   ./execute_social_batch.sh                    ║
║  3. Validar:    Ver SOCIAL-POST-EXECUTION-VALIDATION.md      ║
║                                                               ║
║  Tempo estimado: 25 minutos                                  ║
║  Output: 30 items enriquecidos com 4 campos cada             ║
║                                                               ║
╚═══════════════════════════════════════════════════════════════╝
```

---

## 📚 Links Rápidos

```
┌───────────────────────────────────────────────────────────────┐
│                    DOCUMENTAÇÃO RÁPIDA                        │
├───────────────────────────────────────────────────────────────┤
│                                                               │
│  🏠 Overview:      SOCIAL-README.md                           │
│  ⚡ Quick Start:   SOCIAL-QUICK-START.md                      │
│  📊 Sumário:       SOCIAL-BATCH-EXECUTIVE-SUMMARY.md          │
│  📖 Metodologia:   SOCIAL-ENRICHMENT-METHODOLOGY.md           │
│  📚 Referências:   SOCIAL-SCIENTIFIC-REFERENCES.md            │
│  📝 Exemplos:      SOCIAL-EXPECTED-OUTPUT-EXAMPLES.md         │
│  ✅ Validação:     SOCIAL-POST-EXECUTION-VALIDATION.md        │
│  📑 Índice:        SOCIAL-INDEX.md                            │
│  🎨 Visual:        SOCIAL-VISUAL-GUIDE.md (você está aqui)    │
│                                                               │
└───────────────────────────────────────────────────────────────┘
```

---

## 💡 Dica Final

```
┌───────────────────────────────────────────────────────────────┐
│                                                               │
│  💡 NOVO NO PROJETO?                                         │
│                                                               │
│  Comece por:                                                 │
│  1. SOCIAL-README.md (10 min)                                │
│  2. SOCIAL-QUICK-START.md (5 min)                            │
│  3. ./execute_social_batch.sh (25 min)                       │
│                                                               │
│  Total: ~40 minutos para ter 30 items enriquecidos!          │
│                                                               │
└───────────────────────────────────────────────────────────────┘
```

---

**Criado**: 2026-01-27
**Última atualização**: 2026-01-27
**Versão**: 1.0.0
**Status**: ✅ Completo e pronto para uso
