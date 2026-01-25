# Marcadores Virais - Referência Rápida
## Sistema Plenya EMR

---

## TABELAS DE ESTRATIFICAÇÃO

### 1. Anti-Hbs (QUANTITATIVO)

```
Hepatite B - Anti-Hbs (Anticorpo de Superfície)
| Nível 0 | Nível 1 | Nível 2 | Nível 3 |
| <5 mIU/mL | 5-10 mIU/mL | 10-100 mIU/mL | >100 mIU/mL |
mIU/mL | 1 mIU/mL = 1 IU/L
```

### 2. Anti-Hbc (QUALITATIVO)

```
Hepatite B - Anti-Hbc Total
| Nível 0 | Nível 3 |
| Reagente (sem contexto) | Não-reagente |
Qualitativo | Reagente/Não-reagente
```

**Interpretação obrigatória com HBsAg e Anti-Hbs.**

### 3. HBsAg (QUALITATIVO)

```
Hepatite B - HbsAg (Antígeno de Superfície)
| Nível 0 | Nível 3 |
| Reagente | Não-reagente |
Qualitativo | Reagente/Não-reagente
```

**Reagente = Infecção ativa = CRÍTICO**

### 4. Anti-HCV (QUALITATIVO)

```
Hepatite C - Anti-HCV (Anticorpo)
| Nível 0 | Nível 3 |
| Reagente | Não-reagente |
Qualitativo | Reagente/Não-reagente
```

**Reagente = Necessita HCV RNA PCR**

**Com S/CO Ratio (opcional):**
```
| Nível 0 | Nível 1 | Nível 2 | Nível 3 |
| S/CO ≥10.9 | S/CO 9.0-10.9 | S/CO 1.0-8.9 | S/CO <1.0 |
```

### 5. HIV 1+2 (QUALITATIVO)

```
HIV 1+2 Anticorpo/Antígeno Combo
| Nível 0 | Nível 3 |
| Reagente | Não-reagente |
Qualitativo | Reagente/Não-reagente
```

**Reagente = Algoritmo CDC 3 etapas = CRÍTICO**

---

## INTERPRETAÇÕES RÁPIDAS

### Hepatite B (Triple Panel)

| HBsAg | Anti-Hbs | Anti-Hbc | Interpretação |
|-------|----------|----------|---------------|
| (-) | (-) | (-) | Susceptível → Vacinar |
| (-) | (+) | (-) | Imune por vacinação ✓ |
| (-) | (+) | (+) | Imune por infecção resolvida ✓ |
| (-) | (-) | (+) | Isolado Anti-Hbc → Investigar |
| (+) | (-) | (+) | Infecção ativa 🔴 CRÍTICO |

### Hepatite C

| Anti-HCV | HCV RNA | Interpretação |
|----------|---------|---------------|
| (-) | - | Sem exposição ✓ |
| (+) | Não detectável | Curado (SVR) ⚠️ Monitorar |
| (+) | Detectável | Infecção ativa 🔴 TRATAR |

### HIV (Algoritmo CDC 2024)

1. **HIV 1+2 Ag/Ab Combo:**
   - Não-reagente → Negativo ✓
   - Reagente → Passo 2

2. **HIV-1/2 Differentiation:**
   - HIV-1 (+) → Confirma HIV-1 🔴
   - HIV-2 (+) → Confirma HIV-2 🔴
   - Indeterminado → Passo 3

3. **HIV-1 RNA NAT:**
   - Detectável → HIV-1 agudo 🔴
   - Não detectável → Falso positivo OU repetir

---

## ALERTAS AUTOMÁTICOS

### 🔴 CRÍTICO (Encaminhamento URGENTE)

- **HBsAg reagente** → Hepatologia
- **HIV 1+2 reagente** → Infectologia
- **Anti-HCV reagente + HCV RNA detectável** → Hepatologia

### ⚠️ ALERTA (Teste Confirmatório)

- **Anti-HCV reagente** → Solicitar HCV RNA PCR
- **Anti-Hbc isolado reagente** → Solicitar HBV DNA
- **Anti-Hbs 5-10 mIU/mL** → Considerar reforço vacinal

### ✓ ÓTIMO

- **HBsAg não-reagente**
- **Anti-Hbs >100 mIU/mL**
- **Anti-HCV não-reagente**
- **HIV 1+2 não-reagente**

---

## VALORES DE REFERÊNCIA

### Anti-Hbs (mIU/mL)

- **<5:** Negativo (sem imunidade)
- **5-10:** Indeterminado (repetir)
- **10-100:** Imunidade protetora (padrão CDC)
- **>100:** Imunidade ótima (resposta robusta)

### S/CO Ratio (Anti-HCV)

- **<1.0:** Não-reagente
- **1.0-8.9:** Reagente baixo → RNA qualitativo
- **9.0-10.9:** Reagente médio → RNA qualitativo/quantitativo
- **≥10.9:** Reagente alto → RNA quantitativo

---

## JANELAS IMUNOLÓGICAS

| Teste | Janela |
|-------|--------|
| HIV 1+2 Ag/Ab (4ª geração) | 18-45 dias |
| Anti-HCV | 8-11 semanas |
| HBsAg | 3-5 semanas |
| Anti-Hbs (pós-vacina) | 1-2 meses |

**Se exposição recente, repetir teste após janela.**

---

## ESQUEMA DE BANCO DE DADOS

```sql
-- Adicionar à tabela lab_tests
test_type ENUM('quantitative', 'qualitative', 'semi_quantitative')

-- Adicionar à tabela lab_results
numeric_value DECIMAL(10,4)           -- Para quantitativos
unit VARCHAR(20)                      -- mIU/mL, etc
qualitative_result VARCHAR(20)        -- reactive/non_reactive/indeterminate
s_co_ratio DECIMAL(5,2)              -- Para Anti-HCV
requires_confirmation BOOLEAN
confirmation_test VARCHAR(100)        -- Ex: HCV RNA PCR
```

---

## LÓGICA DE INTERFACE

### Quantitativo (Anti-Hbs)
```
✓ Anti-Hbs: 125 mIU/mL
  Nível 3 - Imunidade Ótima
  Resposta imune robusta contra Hepatite B.
```

### Qualitativo (HBsAg negativo)
```
✓ HBsAg: Não-reagente
  Status: ÓTIMO
  Sem evidência de infecção ativa.
```

### Qualitativo CRÍTICO (HIV reagente)
```
🔴 HIV 1+2: Reagente
   Status: CRÍTICO

   ⚠️ AÇÃO URGENTE NECESSÁRIA
   Seguir algoritmo CDC 3 etapas.
   Encaminhar URGENTEMENTE para infectologia.
```

---

## FONTES PRINCIPAIS

**2024-2026:**
- CDC Guidelines (2023-2024)
- EASL Guidelines (2025)
- Nature npj Vaccines (2025)
- NY Health HIV Algorithm (2024)

**Documentos completos:**
- `/home/user/plenya/VIRAL-MARKERS-RESEARCH.md`
- `/home/user/plenya/VIRAL-MARKERS-STRATIFICATION-TABLES.md`

---

**Versão:** 1.0
**Data:** 18/01/2026
**Sistema:** Plenya EMR
