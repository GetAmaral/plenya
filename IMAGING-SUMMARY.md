# Resumo Executivo - Estratificação de Risco por Imagem

## Quick Reference: CSV vs. Qualitative

### ✅ ADEQUADO PARA CSV (Quantitativo/Semi-quantitativo)

| Exame | Campo CSV | Tipo de Dado | Níveis de Risco |
|-------|-----------|--------------|-----------------|
| **USG Abdome** | `hepatic_steatosis_grade` | Enum: 0, I, II, III | 4 níveis (III=pior, 0=melhor) |
| **USG Abdome** | `cap_score_dbm` | Decimal (dB/m) | 4 níveis (>290=pior, <238=melhor) |
| **USG Abdome** | `fib4_score` | Decimal (calculado) | Varia por idade |
| **USG Próstata** | `prostate_volume_cc` | Decimal (cc) | 4 níveis (>80=pior, <30=melhor) |
| **USG Próstata** | `psa_density` | Decimal (ng/mL/cc) | 4 níveis (>0.20=pior, <0.10=melhor) |
| **TC Tórax** | `largest_solid_nodule_mm` | Decimal (mm) | 4 níveis (>8=pior, ausente=melhor) |
| **TC Tórax** | `emphysema_goddard_score` | Integer 0-24 | 4 níveis (13-24=pior, 0=melhor) |
| **TC Tórax** | `emphysema_laa_percent` | Decimal (%) | 4 níveis (>25%=pior, <1%=melhor) |

**Total: 8 campos quantificáveis para CSV**

### ❌ NÃO ADEQUADO PARA CSV (Qualitativo - Laudo Textual)

**USG Abdome:**
- Colelitíase, nefrolitíase, lesões hepáticas, esplenomegalia, cistos renais

**USG Próstata:**
- Ecotextura, calcificações, lesões focais, resíduo pós-miccional

**TC Tórax:**
- Nódulos vidro fosco, fibrose pulmonar, bronquiectasias, derrame pleural, linfadenopatia

**Recomendação:** Usar campos `report_text` e `additional_findings` (TEXT) no banco de dados.

---

## Clinical Decision Support (CDS) - Alertas Automáticos

### 1. MASLD/NAFLD (USG Abdome + Labs)

```
SE esteatose grau I/II/III detectada:
  → Calcular FIB-4 Score (necessita: idade, AST, ALT, plaquetas)

SE FIB-4 ≥1.3 (idade 36-65) OU ≥2.0 (idade >65):
  ⚠️ WARNING: "Solicitar elastografia transitória (VCTE) ou ELF"

SE FIB-4 >2.67:
  🔴 CRITICAL: "Alto risco fibrose avançada - Encaminhar hepatologia"

SE esteatose grau II ou III:
  ℹ️ INFO: "Intervenção estilo de vida: perda peso 7-10%, dieta mediterrânea, exercício"
```

**Evidência 2024-2025:**
- MASLD substituiu NAFLD como terminologia oficial (consenso 2023)
- Semaglutida aprovada FDA agosto 2025 para MASH
- Resmetirom aprovado FDA março 2024 (primeiro medicamento específico)

### 2. Câncer de Próstata (USG Próstata + PSA)

```
SE PSAD ≥0.20:
  🔴 CRITICAL: "Risco muito alto - RM multiparamétrica + urologia URGENTE"

SE PSAD ≥0.15:
  ⚠️ WARNING: "Alto risco - Solicitar mpMRI próstata"

SE PSAD ≥0.10:
  ℹ️ INFO: "Risco intermediário - Seguimento rigoroso PSA"

SE volume >80cc E PSAD ≥0.30:
  ⚠️ WARNING: "BPH severo - Considerar biópsia mesmo RM negativa (evidência 2025)"
```

**Evidência 2024-2025:**
- 4K Density supera PSAD tradicional (2025)
- TAUS equivalente a RM para volume (diferença média 2.5 mL)
- PSAD 0.30 cutoff para BPH severo com RM negativa

### 3. Nódulos Pulmonares (TC Tórax)

```
SE nódulo >8mm:
  🔴 CRITICAL: "Alto risco - CT 3 meses, considerar PET/CT ou biópsia"

SE nódulo 6-8mm:
  ⚠️ WARNING: "Risco intermediário - CT 6-12 meses, depois 18-24 meses (Fleischner)"

SE nódulo <6mm:
  ℹ️ INFO: "Baixo risco - Sem seguimento necessário (consenso clínico)"
```

**Guidelines 2024-2025:**
- Fleischner Society 2017 ainda vigente em 2025
- Diretrizes Japonesas 6ª edição (2024) usam 6mm como limiar
- USPSTF Grade B: LDCT rastreamento 50-80 anos, ≥20 pack-years

### 4. Enfisema/COPD (TC Tórax)

```
SE Goddard Score ≥13:
  🔴 CRITICAL: "Enfisema severo - Espirometria + pneumologia URGENTE"

SE Goddard Score 7-12:
  ⚠️ WARNING: "Enfisema moderado - Espirometria, cessação tabagismo, broncodilatadores"

SE Goddard Score 1-6:
  ℹ️ INFO: "Enfisema leve - Cessação tabagismo, seguimento anual"
```

**Goddard Score 2024-2025:**
- Método visual simples, não requer software especializado
- Confiável em TC com dose raio-X (PCD-CT 2025)
- Preditor de complicações pós-operatórias (vazamento aéreo)

---

## Tabelas de Estratificação para CSV

### Esteatose Hepática (Grau USG)
```
| Nível 0 | Nível 1 | Nível 2 | Nível 3 |
| Grau III (>66%) | Grau II (34-66%) | Grau I (5-33%) | Grau 0 (Normal) |
Grau USG | Correlação histologia r=0.82
```

### FIB-4 Score (Idade 36-65 anos)
```
| Nível 0 | Nível 1 | Nível 2 |
| >2.67 | 1.3-2.67 | <1.3 |
Índice | (Idade × AST) / (Plaquetas × √ALT)
```

### Volume Prostático
```
| Nível 0 | Nível 1 | Nível 2 | Nível 3 |
| >80 cc | 50-80 cc | 30-50 cc | <30 cc |
cc | Normal: <30 cc
```

### Densidade PSA
```
| Nível 0 | Nível 1 | Nível 2 | Nível 3 |
| >0.20 ng/mL/cc | 0.15-0.20 | 0.10-0.15 | <0.10 |
ng/mL/cc | PSAD = PSA / Volume Próstata
```

### Nódulo Pulmonar (Tamanho)
```
| Nível 0 | Nível 1 | Nível 2 | Nível 3 |
| >8 mm | 6-8 mm | <6 mm | Ausente |
mm | Fleischner Society 2017
```

### Enfisema (Goddard Score)
```
| Nível 0 | Nível 1 | Nível 2 | Nível 3 |
| 13-24 | 7-12 | 1-6 | 0 |
Score 0-24 | 6 áreas × 0-4 pontos cada
```

---

## Implementação Backend - Schema SQL

```sql
CREATE TABLE imaging_exams (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    exam_type VARCHAR(50) NOT NULL CHECK (exam_type IN ('usg_abdomen','usg_prostate','ct_chest')),
    exam_date DATE NOT NULL,

    -- USG Abdome
    hepatic_steatosis_grade VARCHAR(10) CHECK (hepatic_steatosis_grade IN ('0','I','II','III')),
    cap_score_dbm DECIMAL(5,1),

    -- USG Próstata
    prostate_volume_cc DECIMAL(6,2),
    psa_density DECIMAL(5,3),

    -- TC Tórax
    largest_nodule_mm DECIMAL(5,2),
    emphysema_goddard_score SMALLINT CHECK (emphysema_goddard_score BETWEEN 0 AND 24),
    emphysema_laa_percent DECIMAL(5,2),

    -- Laudo Qualitativo
    report_text TEXT NOT NULL,
    additional_findings TEXT,

    radiologist_id UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE INDEX idx_imaging_patient ON imaging_exams(patient_id);
CREATE INDEX idx_imaging_type ON imaging_exams(exam_type);
CREATE INDEX idx_imaging_date ON imaging_exams(exam_date);
```

---

## Medicina Funcional - Perspectivas

### MASLD Reversal

**Intervenções Lifestyle (Pedra Angular):**
- Perda peso 7-10% → Reverte esteatose, melhora fibrose
- Dieta mediterrânea, redução frutose, aumento ômega-3
- Exercício 150-200 min/semana moderado
- Jejum intermitente (emergente evidência)

**Farmacoterapia 2024-2025:**
- **GLP-1 Agonistas:** Semaglutida, liraglutida (reduções significativas gordura hepática)
- **Inibidores SGLT2:** Empagliflozina, dapagliflozina (melhoram sensibilidade insulina)
- **Pioglitazona:** Melhora histologia, reduz fibrose
- **Vitamina E:** 800 IU/dia (não-diabéticos com MASH)

**Projeção:** Prevalência MASLD aumentará de 25-30% (2024) para >55% (2040)

### BPH Management Funcional

- Saw Palmetto (Serenoa repens): Evidência mista
- Beta-sitosterol: Melhora sintomas IPSS
- Licopeno: Antioxidante, pode reduzir progressão
- Exercício regular: Reduz sintomas LUTS
- Redução cafeína/álcool: Melhora urgência

### Saúde Pulmonar Funcional

- **Cessação Tabagismo:** CRÍTICO (reduz progressão enfisema)
- **Antioxidantes:** NAC (N-acetilcisteína), Vitamina C, E
- **Ômega-3:** Reduz inflamação pulmonar
- **Exercício Respiratório:** Melhora capacidade funcional
- **Evitar Poluição:** Filtros HEPA, evitar horários pico

---

## Fontes-Chave (2024-2025)

**MASLD:**
- [RSNA RadioGraphics MASLD Update 2024](https://pubs.rsna.org/doi/10.1148/rg.240221)
- [AASLD FIB-4 Guidelines](https://www.aasld.org/liver-fellow-network)
- [Semaglutide FDA Approval MASH 2025](https://pubmed.ncbi.nlm.nih.gov/41201884/)

**Próstata:**
- [4K Density Study 2025 - The Prostate (Wiley)](https://onlinelibrary.wiley.com/doi/10.1002/pros.70036)
- [EAU Guidelines Prostate Cancer 2025](https://uroweb.org/guidelines/prostate-cancer)
- [PSAD 0.30 Cutoff BMC Urology 2025](https://bmcurol.biomedcentral.com/articles/10.1186/s12894-025-01719-5)

**Pulmão:**
- [Fleischner Society Radiopaedia](https://radiopaedia.org/articles/fleischner-society-pulmonary-nodule-recommendations-1)
- [Japanese Guidelines 6th Edition 2024](https://link.springer.com/article/10.1007/s11604-024-01695-0)
- [Goddard Score 2026 Study](https://onlinelibrary.wiley.com/doi/10.1155/ijbi/7436511)

---

**Documento:** Resumo Executivo Imaging Risk Stratification
**Data:** Janeiro 2026
**Ver documento completo:** `/home/user/plenya/IMAGING-RISK-STRATIFICATION.md`
