# Tabelas de Estratificação - Marcadores Virais
## Sistema Plenya - Medicina Funcional Integrativa

---

## OBSERVAÇÃO CRÍTICA

**Marcadores virais são predominantemente QUALITATIVOS (Reagente/Não-reagente).**

Apenas **Anti-Hbs (Hepatite B)** possui versão quantitativa que permite estratificação tradicional.

Para implementação no sistema:
- **Anti-Hbs:** Usar tabela quantitativa padrão (4 níveis)
- **Demais marcadores:** Usar formato categórico binário ou ternário

---

## TABELA QUANTITATIVA

### 1. Hepatite B - Anti-Hbs (Anticorpo de Superfície)

**Tipo:** Quantitativo - Risco LINEAR (quanto maior, melhor)

```
Hepatite B - Anti-Hbs (Anticorpo de Superfície)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 |
| <5 mIU/mL | 5-10 mIU/mL | 10-100 mIU/mL | >100 mIU/mL |
mIU/mL | 1 mIU/mL = 1 IU/L
```

**Interpretação:**
- **Nível 0:** Negativo - Sem imunidade (necessita vacinação)
- **Nível 1:** Indeterminado - Repetir teste ou considerar reforço
- **Nível 2:** Imunidade protetora - Adequado (padrão CDC ≥10 mIU/mL)
- **Nível 3:** Imunidade ótima - Resposta robusta (ÓTIMO)

---

## TABELAS QUALITATIVAS/CATEGÓRICAS

### 2. Hepatite B - Anti-Hbc (Anticorpo do Core)

**Tipo:** Qualitativo (Total IgG + IgM)

**Formato Simplificado para Sistema:**
```
Hepatite B - Anti-Hbc Total
| Nível 0 | Nível 3 |
| Reagente (sem contexto) | Não-reagente |
Qualitativo | Reagente/Não-reagente
```

**IMPORTANTE:** Resultado isolado tem pouco valor. Deve ser interpretado junto com HBsAg e Anti-Hbs.

**Interpretação Completa (Triple Panel):**

| Anti-Hbc | HBsAg | Anti-Hbs | Significado | Score |
|----------|-------|----------|-------------|-------|
| Não-reagente | Negativo | Negativo | Susceptível (sem exposição) | 1 (necessita vacina) |
| Não-reagente | Negativo | Positivo | Imune por vacinação | 3 (ÓTIMO) |
| Reagente | Negativo | Positivo | Imune por infecção resolvida | 3 (ÓTIMO) |
| Reagente | Negativo | Negativo | Isolado Anti-Hbc (investigar) | 1 (ALERTA) |
| Reagente | Positivo | Negativo | Infecção ativa (aguda/crônica) | 0 (CRÍTICO) |

**Recomendação:** Implementar lógica de interpretação combinada no backend.

---

### 3. Hepatite B - HbsAg (Antígeno de Superfície)

**Tipo:** Qualitativo

```
Hepatite B - HbsAg (Antígeno de Superfície)
| Nível 0 | Nível 3 |
| Reagente | Não-reagente |
Qualitativo | Reagente/Não-reagente
```

**Interpretação:**
- **Nível 0 (Reagente):** Infecção ativa por HBV - **CRÍTICO**
  - Indica presença do vírus no sangue
  - Pessoa é infecciosa
  - Encaminhamento URGENTE para hepatologia

- **Nível 3 (Não-reagente):** Sem infecção ativa - **ÓTIMO**
  - Se Anti-Hbs positivo = imune
  - Se todos negativos = susceptível

**Alerta Automático:**
```
SE HbsAg = Reagente
ENTÃO exibir: "URGENTE: Infecção ativa por Hepatite B. Encaminhar IMEDIATAMENTE para hepatologia."
```

---

### 4. Hepatite C - Anti-HCV (Anticorpo)

**Tipo:** Qualitativo com S/CO Ratio (Semi-quantitativo)

**Formato Básico:**
```
Hepatite C - Anti-HCV (Anticorpo)
| Nível 0 | Nível 3 |
| Reagente | Não-reagente |
Qualitativo | Reagente/Não-reagente
```

**Formato Expandido com S/CO Ratio (opcional):**
```
Hepatite C - Anti-HCV
| Nível 0 | Nível 1 | Nível 2 | Nível 3 |
| S/CO ≥10.9 | S/CO 9.0-10.9 | S/CO 1.0-8.9 | S/CO <1.0 |
S/CO Ratio | Signal-to-Cutoff (≥1.0 = Reagente)
```

**Interpretação:**
- **Nível 0 (Reagente, S/CO alto):** Alta probabilidade de viremia
  - Necessita confirmação com HCV RNA quantitativo
  - ALERTA

- **Nível 1-2 (Reagente, S/CO baixo/médio):** Possível exposição
  - Necessita confirmação com HCV RNA qualitativo
  - ALERTA

- **Nível 3 (Não-reagente):** Sem exposição ao HCV - **ÓTIMO**

**Alerta Automático:**
```
SE Anti-HCV = Reagente
ENTÃO exibir: "ALERTA: Exposição ao vírus da Hepatite C. Confirmar infecção ativa com HCV RNA PCR."
```

**Confirmação Obrigatória:**

| Anti-HCV | HCV RNA PCR | Interpretação Final |
|----------|-------------|---------------------|
| Reagente | Detectável | Infecção ativa - TRATAR (Score 0) |
| Reagente | Não detectável | Infecção curada - Monitorar (Score 2) |
| Não-reagente | - | Sem exposição (Score 3) |

---

### 5. HIV 1+2 (Anticorpo/Antígeno Combo - 4ª Geração)

**Tipo:** Qualitativo

```
HIV 1+2 Anticorpo/Antígeno Combo
| Nível 0 | Nível 3 |
| Reagente | Não-reagente |
Qualitativo | Reagente/Não-reagente
```

**Interpretação:**
- **Nível 0 (Reagente):** Possível infecção por HIV - **CRÍTICO**
  - Iniciar algoritmo CDC de 3 etapas
  - Encaminhamento URGENTE para infectologia

- **Nível 3 (Não-reagente):** Sem infecção detectável - **ÓTIMO**
  - Se exposição recente (<45 dias), repetir teste

**Alerta Automático:**
```
SE HIV 1+2 = Reagente
ENTÃO exibir: "CRÍTICO: Possível infecção por HIV. Seguir algoritmo diagnóstico em 3 etapas (CDC 2024). Encaminhar URGENTEMENTE para infectologia."
```

**Algoritmo Diagnóstico (CDC 2024):**

**Passo 1:** HIV-1/2 Ag/Ab Combo
- Não-reagente → Negativo
- Reagente → Passo 2

**Passo 2:** HIV-1/2 Antibody Differentiation
- HIV-1 (+) → Confirma HIV-1
- HIV-2 (+) → Confirma HIV-2
- Indeterminado → Passo 3

**Passo 3:** HIV-1 RNA NAT
- Detectável → HIV-1 agudo
- Não detectável → Falso positivo OU repetir

---

## RECOMENDAÇÕES DE IMPLEMENTAÇÃO

### Opção 1: Sistema Simples (Recomendado para MVP)

**Para todos os qualitativos:**
- Armazenar resultado: `"reactive"` ou `"non_reactive"`
- Não calcular score numérico
- Exibir flags visuais:
  - ✓ Verde: Não-reagente (ÓTIMO)
  - ⚠️ Amarelo: Reagente (ALERTA - necessita confirmação)
  - 🔴 Vermelho: Confirmado positivo (CRÍTICO)

**Schema banco de dados:**
```sql
CREATE TABLE lab_results (
  id UUID PRIMARY KEY,
  patient_id UUID NOT NULL,
  test_id UUID NOT NULL,

  -- Para quantitativos (Anti-Hbs)
  numeric_value DECIMAL(10,4),
  unit VARCHAR(20),
  risk_level INT, -- 0, 1, 2, 3

  -- Para qualitativos
  qualitative_result VARCHAR(20), -- 'reactive', 'non_reactive', 'indeterminate'
  s_co_ratio DECIMAL(5,2), -- Apenas Anti-HCV

  -- Metadata
  requires_confirmation BOOLEAN DEFAULT FALSE,
  confirmation_test VARCHAR(100),
  clinical_significance TEXT,

  created_at TIMESTAMP,
  updated_at TIMESTAMP
);
```

### Opção 2: Sistema com Score Unificado

Forçar qualitativos em escala 0-3:

| Resultado | Score | Categoria |
|-----------|-------|-----------|
| Reagente (confirmado) | 0 | CRÍTICO |
| Reagente (necessita confirmação) | 1 | ALERTA |
| Não-reagente | 3 | ÓTIMO |

**Limitação:** Perde nuances clínicas importantes.

---

## REGRAS DE NEGÓCIO

### 1. Validação de Entrada
```javascript
// Anti-Hbs: aceitar apenas valores numéricos
if (test_name === "Anti-Hbs") {
  validate_numeric(value, min: 0, max: 1000, unit: "mIU/mL")
}

// Demais: aceitar apenas qualitativo
else if (test_name in ["Anti-Hbc", "HBsAg", "Anti-HCV", "HIV 1+2"]) {
  validate_enum(value, ["reactive", "non_reactive", "indeterminate"])
}
```

### 2. Alertas Automáticos
```javascript
const alerts = {
  "HBsAg": {
    "reactive": {
      level: "CRITICAL",
      message: "URGENTE: Infecção ativa por Hepatite B detectada.",
      action: "Encaminhar IMEDIATAMENTE para hepatologia."
    }
  },
  "Anti-HCV": {
    "reactive": {
      level: "ALERT",
      message: "Exposição ao vírus da Hepatite C detectada.",
      action: "Solicitar HCV RNA PCR para confirmar infecção ativa."
    }
  },
  "HIV 1+2": {
    "reactive": {
      level: "CRITICAL",
      message: "Possível infecção por HIV detectada.",
      action: "Seguir algoritmo CDC 3 etapas. Encaminhar URGENTEMENTE para infectologia."
    }
  },
  "Anti-Hbs": {
    numeric: {
      "<5": {
        level: "ALERT",
        message: "Sem imunidade contra Hepatite B.",
        action: "Considerar vacinação ou reforço."
      },
      "5-10": {
        level: "WARNING",
        message: "Imunidade limítrofe contra Hepatite B.",
        action: "Considerar reforço vacinal, especialmente em profissionais de saúde."
      }
    }
  }
}
```

### 3. Interface do Usuário

**Exibição de Resultado Quantitativo (Anti-Hbs):**
```
┌─────────────────────────────────────────────────┐
│ Hepatite B - Anti-Hbs (Anticorpo de Superfície)│
├─────────────────────────────────────────────────┤
│ Resultado: 125 mIU/mL                           │
│ Nível de Risco: 3 (Imunidade Ótima) ✓          │
│                                                  │
│ Interpretação:                                   │
│ Resposta imune robusta contra Hepatite B.       │
│ Proteção de longo prazo estabelecida.          │
└─────────────────────────────────────────────────┘
```

**Exibição de Resultado Qualitativo (HBsAg):**
```
┌─────────────────────────────────────────────────┐
│ Hepatite B - HbsAg (Antígeno de Superfície)    │
├─────────────────────────────────────────────────┤
│ Resultado: Não-reagente ✓                      │
│ Status: ÓTIMO                                   │
│                                                  │
│ Interpretação:                                   │
│ Sem evidência de infecção ativa por Hepatite B.│
└─────────────────────────────────────────────────┘
```

**Exibição de Resultado CRÍTICO (HIV reagente):**
```
┌─────────────────────────────────────────────────┐
│ HIV 1+2 Anticorpo/Antígeno Combo                │
├─────────────────────────────────────────────────┤
│ Resultado: Reagente 🔴                          │
│ Status: CRÍTICO                                 │
│                                                  │
│ ⚠️ AÇÃO URGENTE NECESSÁRIA                     │
│                                                  │
│ Este resultado indica possível infecção por HIV.│
│                                                  │
│ Próximos passos obrigatórios:                   │
│ 1. Realizar teste de diferenciação HIV-1/HIV-2 │
│ 2. Se indeterminado, realizar HIV-1 RNA NAT    │
│ 3. Encaminhar URGENTEMENTE para infectologia    │
│                                                  │
│ Algoritmo: CDC 2024 (3 etapas)                  │
└─────────────────────────────────────────────────┘
```

---

## CÁLCULO DE HEALTH SCORE GLOBAL

### Problema: Como incluir testes qualitativos no score?

**Opção A: Não incluir no score automático**
- Exames virais são flags de alerta, não scores contínuos
- Exibir separadamente como "Alertas Clínicos"
- Health Score = apenas exames quantitativos

**Opção B: Incluir com peso binário**
- Reagente = 0 pontos (pior)
- Não-reagente = 100 pontos (ótimo)
- Peso igual aos demais exames

**Opção C: Incluir com peso crítico aumentado**
- Se qualquer marcador viral reagente → reduzir Health Score global em X%
- Exemplo: HIV reagente → Score máximo = 50% (independente de outros exames)
- Reflete gravidade de infecção viral ativa

**Recomendação: Opção A (Separar alertas de score)**
- Mais intuitivo clinicamente
- Evita "gamificação" de doenças graves
- Alertas visuais destacados (bandeira vermelha)

---

## CASOS DE USO

### Caso 1: Paciente Vacinado contra Hepatite B
```
Entrada:
- Anti-Hbs: 45 mIU/mL
- Anti-Hbc: Não-reagente
- HBsAg: Não-reagente

Resultado:
Anti-Hbs: Nível 2 (Imunidade protetora) ✓
Interpretação: Imune por vacinação. Proteção adequada contra Hepatite B.
```

### Caso 2: Paciente com Infecção Resolvida
```
Entrada:
- Anti-Hbs: 150 mIU/mL
- Anti-Hbc: Reagente
- HBsAg: Não-reagente

Resultado:
Anti-Hbs: Nível 3 (Imunidade ótima) ✓
Anti-Hbc: Reagente (com contexto de Anti-Hbs+, HBsAg-) ✓
Interpretação: Infecção passada resolvida com imunidade robusta.
```

### Caso 3: Paciente com Hepatite B Crônica
```
Entrada:
- Anti-Hbs: <5 mIU/mL
- Anti-Hbc: Reagente
- HBsAg: Reagente

Resultado:
HBsAg: Reagente 🔴 CRÍTICO
Anti-Hbc: Reagente (contexto de infecção ativa)
Anti-Hbs: Nível 0 (sem imunidade)

Alerta: URGENTE - Infecção crônica por Hepatite B detectada.
Ação: Encaminhar IMEDIATAMENTE para hepatologia.
```

### Caso 4: Hepatite C Curada
```
Entrada:
- Anti-HCV: Reagente
- HCV RNA PCR: Não detectável

Resultado:
Anti-HCV: Reagente ⚠️ ALERTA
HCV RNA: Não detectável ✓

Interpretação: Infecção passada curada (SVR - Resposta Virológica Sustentada).
Nota: Paciente não está imune. Pode reinfectar se exposto novamente.
Recomendação: Monitorar função hepática.
```

---

## CHECKLIST DE IMPLEMENTAÇÃO

- [ ] Criar campo `test_type` na tabela `lab_tests`
  - [ ] Valores: `quantitative`, `qualitative`, `semi_quantitative`

- [ ] Adicionar campos na tabela `lab_results`
  - [ ] `qualitative_result` (reactive/non_reactive/indeterminate)
  - [ ] `s_co_ratio` para Anti-HCV
  - [ ] `requires_confirmation` boolean
  - [ ] `confirmation_test` string

- [ ] Implementar lógica de validação
  - [ ] Validar tipo de entrada baseado em `test_type`
  - [ ] Rejeitar valores inválidos com mensagem clara

- [ ] Criar sistema de alertas
  - [ ] Definir regras de alerta por exame
  - [ ] Categorizar níveis: INFO, WARNING, ALERT, CRITICAL
  - [ ] Gerar notificações para médico responsável

- [ ] Desenvolver interface para qualitativos
  - [ ] Design de card diferenciado (não usar barras de nível)
  - [ ] Badges coloridos: Verde (✓), Amarelo (⚠️), Vermelho (🔴)
  - [ ] Exibir interpretação clínica contextual

- [ ] Implementar interpretação combinada (Triple Panel Hepatite B)
  - [ ] Lógica backend para cruzar HBsAg + Anti-Hbs + Anti-Hbc
  - [ ] Gerar interpretação automática

- [ ] Criar seção separada de "Alertas Infecciosos"
  - [ ] Não incluir no Health Score global
  - [ ] Destacar visualmente na dashboard
  - [ ] Priorizar por gravidade

- [ ] Documentar para equipe médica
  - [ ] Guia de interpretação de resultados
  - [ ] Fluxogramas de conduta (algoritmos CDC/EASL)
  - [ ] Quando solicitar testes confirmatórios

---

**Última atualização:** 18 de janeiro de 2026
**Sistema:** Plenya EMR v1.0
**Módulo:** Estratificação de Risco - Marcadores Virais
