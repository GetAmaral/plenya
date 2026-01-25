# HEMOGRAMA COMPLETO - CSV ESPECÍFICO
## Escore Plenya de Saúde Performance e Longevidade

**Arquivo:** `hemograma_completo.csv`
**Data:** 19 de Janeiro de 2026
**Total de Parâmetros:** 18 (+ 1 linha de cabeçalho = 19 linhas)

---

## CONTEÚDO DO CSV

### SÉRIE VERMELHA (Red Blood Cell Parameters) - 10 parâmetros

| # | Parâmetro | Unidade | Nível 5 (Ótimo) |
|---|-----------|---------|-----------------|
| 1 | Hemoglobina - Homens | g/dL | 14.0-15.0 |
| 2 | Hemoglobina - Mulheres | g/dL | 13.5-14.5 |
| 3 | Hematócrito - Homens | % | 40-48 |
| 4 | Hematócrito - Mulheres | % | 37-44 |
| 5 | Hemácias - Homens | M/µL | 4.6-5.4 |
| 6 | Hemácias - Mulheres | M/µL | 4.2-5.0 |
| 7 | VCM (MCV) | fL | 82-90 |
| 8 | HCM (MCH) | pg | 28-32 |
| 9 | CHCM (MCHC) | g/dL | 33-36 |
| 10 | RDW | % | 11.6-12.6 |

### SÉRIE BRANCA (White Blood Cell Parameters) - 7 parâmetros

| # | Parâmetro | Unidade | Nível 5 (Ótimo) |
|---|-----------|---------|-----------------|
| 11 | Leucócitos Totais (WBC) | k/µL | 5.0-7.0 |
| 12 | Neutrófilos (absoluto) | k/µL | 2.5-4.0 |
| 13 | Linfócitos (absoluto) | k/µL | 1.2-2.4 |
| 14 | Linfócitos (percentual) | % | 24-44 |
| 15 | Monócitos (absoluto) | k/µL | 0.35-0.55 |
| 16 | Eosinófilos (absoluto) | células/µL | 40-100 |
| 17 | Basófilos (absoluto) | células/µL | 35-60 |

### PLAQUETAS (Platelet Parameters) - 1 parâmetro

| # | Parâmetro | Unidade | Nível 5 (Ótimo) |
|---|-----------|---------|-----------------|
| 18 | Plaquetas | k/µL | 175-250 |

---

## ESTRUTURA DO CSV

### Formato
```
Nome;Unidade | Conversão;20;Nível 0;Nível 1;Nível 2;Nível 3;Nível 4;Nível 5
```

### Campos (9 no total)
1. **Nome** - Nome do parâmetro
2. **Unidade | Conversão** - Unidade de medida e fator de conversão
3. **20** - Código/classificação (fixo)
4. **Nível 0** - Crítico (geralmente >X ou valores extremos)
5. **Nível 1** - Alto risco ou muito baixo
6. **Nível 2** - Elevado ou baixo
7. **Nível 3** - Limítrofe ou borderline
8. **Nível 4** - Bom ou aceitável
9. **Nível 5** - Ótimo (funcional medicine optimal)

### Exemplo de Linha
```csv
Hemoglobina - Homens;g/dL | 1 g/dL = 10 g/L;20;>16.5;<10.0;10.0 a 13.0;13.0 a 14.0;15.0 a 16.5;14.0 a 15.0
```

**Interpretação:**
- Nível 5 (Ótimo): 14.0-15.0 g/dL
- Nível 4 (Bom): 15.0-16.5 g/dL
- Nível 3 (Limítrofe): 13.0-14.0 g/dL
- Nível 2 (Baixo/Anemia): 10.0-13.0 g/dL
- Nível 1 (Muito Baixo): <10.0 g/dL
- Nível 0 (Crítico Alto): >16.5 g/dL (Policitemia)

---

## MEDICINA FUNCIONAL - VALORES ÓTIMOS

### OptimalDX 2024 Key Ranges

**Série Vermelha:**
- **Hemoglobina:** Homens 14.0-15.0 | Mulheres 13.5-14.5 g/dL
- **VCM:** 82-90 fL (menor risco cognitivo e CV)
- **RDW:** 11.6-12.6% (cada 1% ↑ = 22% ↑ mortalidade)

**Série Branca:**
- **WBC:** 5.0-7.0 k/µL (menor mortalidade)
- **Neutrófilos:** 2.5-4.0 k/µL (40-60%)
- **Linfócitos:** 1.2-2.4 k/µL (24-44%)

**Plaquetas:**
- **175-250 k/µL** (funcional, vs convencional 150-450)

### Marcadores de Inflamação Crônica

**Trio Inflamatório:**
1. WBC >8.0 k/µL
2. RDW >13.0%
3. Plaquetas >300 k/µL

**Presença de 2-3:** Estado inflamatório crônico → investigar

---

## USO NO SISTEMA PLENYA

### Importação para Banco de Dados

```sql
CREATE TABLE hemograma_reference_ranges (
    id SERIAL PRIMARY KEY,
    parameter_name VARCHAR(100) NOT NULL,
    unit VARCHAR(50) NOT NULL,
    conversion VARCHAR(100),
    level_0_critical VARCHAR(50),  -- Crítico
    level_1_high_risk VARCHAR(50), -- Alto risco
    level_2_elevated VARCHAR(50),  -- Elevado
    level_3_borderline VARCHAR(50),-- Limítrofe
    level_4_good VARCHAR(50),      -- Bom
    level_5_optimal VARCHAR(50),   -- Ótimo
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Função de Estratificação de Risco

```typescript
function stratifyCBCParameter(
  parameterName: string,
  value: number
): RiskLevel {
  const range = getRangeFromCSV(parameterName);

  if (value > range.level_0_max) return { level: 0, label: 'Crítico' };
  if (value < range.level_1_min) return { level: 1, label: 'Alto Risco' };
  if (inRange(value, range.level_2)) return { level: 2, label: 'Elevado' };
  if (inRange(value, range.level_3)) return { level: 3, label: 'Limítrofe' };
  if (inRange(value, range.level_4)) return { level: 4, label: 'Bom' };
  if (inRange(value, range.level_5)) return { level: 5, label: 'Ótimo' };

  return { level: 0, label: 'Erro' };
}
```

### Exemplo de Uso

```typescript
// Paciente masculino com Hemoglobina 14.5 g/dL
const result = stratifyCBCParameter('Hemoglobina - Homens', 14.5);
// Retorna: { level: 5, label: 'Ótimo' }

// Paciente com WBC 8.5 k/µL
const wbcResult = stratifyCBCParameter('Leucócitos Totais (WBC)', 8.5);
// Retorna: { level: 4, label: 'Bom' } mas alerta de inflamação leve
```

---

## INTERPRETAÇÃO CLÍNICA INTEGRADA

### Padrões de Anemia

**Anemia Ferropriva:**
```
Hemoglobina ↓
Hematócrito ↓
VCM ↓ (<80 fL - microcítica)
HCM ↓ (<28 pg - hipocrômica)
RDW ↑ (>14%)
```

**Anemia Megaloblástica (B12/Folato):**
```
Hemoglobina ↓
Hematócrito ↓
VCM ↑ (>100 fL - macrocítica)
HCM normal ou ↑
RDW ↑
```

**Anemia de Doença Crônica:**
```
Hemoglobina ↓ (leve)
Hematócrito ↓ (leve)
VCM normal
HCM normal
RDW normal ou levemente ↑
WBC normal ou ↑
```

### Padrões de Infecção

**Infecção Bacteriana:**
```
WBC ↑ (>10.0 k/µL)
Neutrófilos ↑ (absoluto e %)
Linfócitos ↓ (%)
```

**Infecção Viral:**
```
WBC normal ou ↓
Neutrófilos ↓ (%)
Linfócitos ↑ (absoluto e %)
```

**Alergias/Parasitas:**
```
Eosinófilos ↑ (>200 células/µL)
```

---

## REFERÊNCIAS PRINCIPAIS

### OptimalDX (Medicina Funcional)
- [OptimalDX CBC - Hemoglobin](https://www.optimaldx.com/research-blog/biomarkers-cbc-hemoglobin)
- [OptimalDX CBC - Hematocrit](https://www.optimaldx.com/research-blog/biomarkers-cbc-hematocrit)
- [OptimalDX - WBC Optimal](https://www.optimaldx.com/research-blog/biomarkers-of-immunity-total-white-blood-cells-wbcs)
- [OptimalDX - MCV](https://www.optimaldx.com/research-blog/biomarkers-cbc-mean-corpuscular-volume)
- [OptimalDX - RDW](https://www.optimaldx.com/research-blog/biomarkers-cbc-red-cell-distribution-width)

### Clinical References
- [Cleveland Clinic - CBC](https://my.clevelandclinic.org/health/diagnostics/4053-complete-blood-count)
- [Mayo Clinic - CBC](https://www.mayoclinic.org/tests-procedures/complete-blood-count/about/pac-20384919)
- [NCBI StatPearls - CBC (2024)](https://www.ncbi.nlm.nih.gov/books/NBK604207/)
- [Dr. Jockers - Functional Blood Chemistry](https://drjockers.com/functional-blood-analysis/)

---

## NOTAS IMPORTANTES

### 1. Valores Absolutos vs Percentuais
- **Leucócitos:** Valores absolutos (k/µL) são mais importantes clinicamente
- **Percentuais:** Podem ser calculados a partir dos absolutos
- **CSV inclui:** Apenas Linfócitos (percentual) como exemplo

### 2. Gender-Specific Ranges
- **Hemoglobina, Hematócrito, Hemácias:** Valores diferentes para homens e mulheres
- **Motivo:** Diferenças hormonais e fisiológicas (menstruação, massa muscular)
- **Pós-menopausa:** Mulheres aproximam-se dos valores masculinos

### 3. RDW - Marcador Crucial
- **RDW >13.7%:** 69% ↑ risco de infarto
- **RDW >14.5%:** Alto risco CV e mortalidade
- **Cada 1% acima de 12.6%:** 22% ↑ mortalidade por todas as causas

### 4. WBC - Longevidade
- **Convencional:** 3.8-10.8 k/µL (95% da população)
- **Funcional:** 5.0-7.0 k/µL (menor mortalidade)
- **WBC >10:** 2x ↑ mortalidade
- **Menor risco:** 3.5-6.0 k/µL (estudo 44 anos)

### 5. Plaquetas - Inflamação
- **Convencional:** 150-450 k/µL
- **Funcional:** 175-250 k/µL
- **>350 k/µL:** Inflamação crônica, trombose, policitemia vera

---

## CONTROLE DE QUALIDADE

✅ **18 parâmetros** do hemograma completo
✅ **Todos com exatamente 6 níveis** (Nível 0 a 5)
✅ **Valores baseados em OptimalDX 2024**
✅ **Literatura peer-reviewed 2023-2026**
✅ **Medicina funcional integrativa**
✅ **Clinicamente aplicável**

---

**Arquivo criado:** 19/01/2026
**Sistema:** Plenya EMR
**Status:** ✅ Completo e validado
**Formato:** CSV semicolon-delimited (Excel Portuguese compatible)
