-- ============================================================================
-- PSAD (Densidade PSA) - Clinical Enrichment Script
-- Item ID: 317acc85-3ce9-4f97-8e14-799354166f5e
-- Grupo: Exames > Imagem
--
-- Evidência científica baseada em:
-- - Peng et al. (2025) - BMC Urology - DOI: 10.1186/s12894-025-01719-5
-- - Chou et al. (2025) - Diagnostics - DOI: 10.3390/diagnostics15162027
-- - Yusim et al. (2020) - Scientific Reports - DOI: 10.1038/s41598-020-76786-9
-- ============================================================================

BEGIN;

-- 1. Insert articles (if not exist)
WITH article_data AS (
    SELECT
        'Optimal PSA density threshold for prostate biopsy in benign prostatic obstruction patients with elevated PSA levels but negative MRI findings' AS title,
        'Peng Y, Wei C, Li Y, Zhao F, Liu Y, Jiang T, Chen Z, Zheng J, Fu J, Wang P, Shen W' AS authors,
        'BMC Urology' AS journal,
        '2025-03-03'::DATE AS publish_date,
        '10.1186/s12894-025-01719-5' AS doi,
        'https://pmc.ncbi.nlm.nih.gov/articles/PMC11874838/' AS original_link,
        NULL::TEXT AS pm_id,
        'Study identifying optimal PSAD cutoff of 0.30 ng/ml/cm³ for biopsy decision in BPH patients with elevated PSA but negative MRI, demonstrating 93% specificity and 65% sensitivity. ROC analysis showed PSAD achieved AUC 0.848, outperforming PSA alone (0.722) and free/total PSA ratio (0.635).' AS abstract
),
inserted_article_1 AS (
    INSERT INTO articles (title, authors, journal, publish_date, doi, original_link, pm_id, abstract, created_at, updated_at)
    SELECT title, authors, journal, publish_date, doi, original_link, pm_id, abstract, NOW(), NOW()
    FROM article_data
    ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
    RETURNING id
)
SELECT id AS article_1_id INTO TEMP TABLE article_1 FROM inserted_article_1;

-- Article 2
WITH article_data AS (
    SELECT
        'Integrating PSA Change with PSA Density Enhances Diagnostic Accuracy and Helps Avoid Unnecessary Prostate Biopsies' AS title,
        'Chou YJ, Jong BE, Tsai YC' AS authors,
        'Diagnostics (Basel)' AS journal,
        '2025-08-13'::DATE AS publish_date,
        '10.3390/diagnostics15162027' AS doi,
        'https://pmc.ncbi.nlm.nih.gov/articles/PMC12385582/' AS original_link,
        '40870878' AS pm_id,
        'Demonstrates that PSA density shows superior diagnostic performance (AUC 0.77-0.81) compared to PSA change alone. Combining both metrics provides optimal results, with >20% PSA decline criterion improving PSAD performance, especially valuable in prostates >80 mL where PSAD accuracy decreases.' AS abstract
),
inserted_article_2 AS (
    INSERT INTO articles (title, authors, journal, publish_date, doi, original_link, pm_id, abstract, created_at, updated_at)
    SELECT title, authors, journal, publish_date, doi, original_link, pm_id, abstract, NOW(), NOW()
    FROM article_data
    ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
    RETURNING id
)
SELECT id AS article_2_id INTO TEMP TABLE article_2 FROM inserted_article_2;

-- Article 3
WITH article_data AS (
    SELECT
        'The use of prostate specific antigen density to predict clinically significant prostate cancer' AS title,
        'Yusim I, Krenawi M, Mazor E, Novack V, Mabjeesh NJ' AS authors,
        'Scientific Reports' AS journal,
        '2020-11-17'::DATE AS publish_date,
        '10.1038/s41598-020-76786-9' AS doi,
        'https://pmc.ncbi.nlm.nih.gov/articles/PMC7672084/' AS original_link,
        '33203873' AS pm_id,
        'Evaluated 992 men undergoing biopsy, finding PSAD AUC of 0.78 vs PSA AUC of 0.64 for predicting clinically significant cancer. Key thresholds: PSAD <0.09 ng/ml² only 4% risk; PSAD 0.09-0.19 ng/ml² risk increases with smaller prostates; PSAD ≥0.20 ng/ml² optimal cutoff with 70% sensitivity and 79% specificity.' AS abstract
),
inserted_article_3 AS (
    INSERT INTO articles (title, authors, journal, publish_date, doi, original_link, pm_id, abstract, created_at, updated_at)
    SELECT title, authors, journal, publish_date, doi, original_link, pm_id, abstract, NOW(), NOW()
    FROM article_data
    ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
    RETURNING id
)
SELECT id AS article_3_id INTO TEMP TABLE article_3 FROM inserted_article_3;

-- 2. Create many-to-many relationships (if not exist)
INSERT INTO article_score_items (score_item_id, article_id)
SELECT '317acc85-3ce9-4f97-8e14-799354166f5e', article_1_id FROM article_1
ON CONFLICT (score_item_id, article_id) DO NOTHING;

INSERT INTO article_score_items (score_item_id, article_id)
SELECT '317acc85-3ce9-4f97-8e14-799354166f5e', article_2_id FROM article_2
ON CONFLICT (score_item_id, article_id) DO NOTHING;

INSERT INTO article_score_items (score_item_id, article_id)
SELECT '317acc85-3ce9-4f97-8e14-799354166f5e', article_3_id FROM article_3
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- 3. Update score_item with clinical content
UPDATE score_items
SET
    clinical_relevance = 'A Densidade do PSA (PSAD) é um parâmetro calculado pela divisão do valor do PSA sérico (ng/mL) pelo volume prostático obtido por ultrassonografia (cm³). Este índice melhora significativamente a capacidade de predição de câncer de próstata clinicamente significativo em comparação ao PSA isolado.

**Fundamento Clínico:**
A PSAD corrige o efeito do volume prostático no valor do PSA, permitindo diferenciar melhor entre elevações do PSA causadas por hiperplasia prostática benigna (HPB) e aquelas relacionadas a neoplasia maligna. Estudos demonstram que a PSAD apresenta área sob a curva ROC de 0.78-0.85, superior ao PSA isolado (AUC 0.64-0.72).

**Pontos de Corte Baseados em Evidências:**

1. **PSAD < 0.09-0.10 ng/mL/cm³:**
   - Baixíssimo risco de câncer clinicamente significativo (4-6%)
   - Pode-se considerar evitar biópsia, especialmente se RM negativa

2. **PSAD 0.10-0.15 ng/mL/cm³:**
   - Risco intermediário
   - Decisão compartilhada considerando RM e fatores individuais

3. **PSAD 0.15-0.20 ng/mL/cm³:**
   - Alto risco
   - Forte indicação de biópsia se lesão suspeita na RM

4. **PSAD ≥ 0.20-0.30 ng/mL/cm³:**
   - Risco muito alto de câncer significativo
   - Indicação clara de biópsia mesmo com RM negativa ou equívoca

**Integração com Ressonância Magnética:**
Para pacientes com PI-RADS 3 (achados equívocos), um limiar de PSAD ≥ 0.20 ng/mL/cm³ demonstra melhor balanço entre sensibilidade e especificidade. Em pacientes com RM negativa (PI-RADS ≤2) e PSAD < 0.15 ng/mL/cm³, pode-se considerar evitar biópsia através de decisão compartilhada.

**Considerações Especiais:**
- A acurácia da PSAD diminui em próstatas muito volumosas (>80 mL)
- A combinação de PSAD com cinética do PSA (variação percentual ao longo do tempo) pode reduzir ainda mais biópsias desnecessárias
- Em pacientes com HPB obstrutiva e RM negativa, o limiar de 0.30 ng/mL/cm³ demonstrou especificidade de 93% com sensibilidade de 65%

**Referências Científicas:**
- Peng Y, et al. BMC Urol. 2025;25:38. DOI: 10.1186/s12894-025-01719-5
- Chou YJ, et al. Diagnostics. 2025;15(16):2027. DOI: 10.3390/diagnostics15162027
- Yusim I, et al. Sci Rep. 2020;10:20015. DOI: 10.1038/s41598-020-76786-9',

    patient_explanation = 'A **Densidade do PSA (PSAD)** é um cálculo que ajuda os médicos a entenderem melhor se uma alteração no seu exame de PSA pode ser preocupante ou não.

**Como funciona:**
O médico pega o valor do seu PSA (aquele exame de sangue) e divide pelo tamanho da sua próstata (medido no ultrassom). Por exemplo, se seu PSA é 6 ng/mL e sua próstata tem 40 cm³ de volume, sua PSAD seria 6 ÷ 40 = 0.15 ng/mL/cm³.

**Por que isso é importante:**
Próstatas maiores naturalmente produzem mais PSA. Então, um PSA de 6 ng/mL pode ser normal se você tem uma próstata grande por hiperplasia benigna (crescimento não-canceroso), mas pode ser preocupante se sua próstata for pequena.

**O que os números significam:**

• **PSAD abaixo de 0.10:** Muito tranquilo. O risco de ter um câncer significativo é baixíssimo (cerca de 4%). Geralmente não precisa fazer biópsia.

• **PSAD entre 0.10 e 0.15:** Risco intermediário. Seu médico vai avaliar junto com outros exames (como a ressonância) se precisa investigar mais.

• **PSAD entre 0.15 e 0.20:** Risco aumentado. Provavelmente será necessária uma biópsia, especialmente se houver algo suspeito na ressonância.

• **PSAD acima de 0.20:** Risco alto. Forte indicação de fazer biópsia para investigar melhor, mesmo que a ressonância não mostre algo muito evidente.

**Vantagens deste cálculo:**
- Ajuda a evitar biópsias desnecessárias em homens com próstatas grandes e PSA elevado apenas por isso
- Identifica melhor os casos que realmente precisam de investigação aprofundada
- Quando combinado com a ressonância magnética, melhora ainda mais a precisão do diagnóstico

**Importante:**
A PSAD é uma ferramenta auxiliar. Seu médico vai considerar este valor junto com sua idade, histórico familiar, toque retal, resultado da ressonância e outros fatores para decidir os próximos passos do seu acompanhamento.',

    conduct = '**Interpretação e Conduta Baseada na PSAD:**

**1. PSAD < 0.10 ng/mL/cm³:**
- ✓ Considerar vigilância ativa sem biópsia imediata
- ✓ Repetir PSA e USG em 12 meses
- ✓ Se RM foi realizada e negativa (PI-RADS ≤2), risco de câncer significativo é ~4%
- ⚠️ Discussão compartilhada com paciente sobre risco residual

**2. PSAD 0.10-0.15 ng/mL/cm³:**
- ⚡ Zona cinzenta - individualizar decisão
- ✓ RM de próstata mandatória se ainda não realizada
- ✓ Se PI-RADS ≤2: considerar seguimento sem biópsia (risco ~6%)
- ✓ Se PI-RADS 3-5: indicar biópsia dirigida por fusão
- ✓ Avaliar cinética do PSA (variação temporal)
- ✓ Considerar testes complementares (PHI, 4K Score, PCA3)

**3. PSAD 0.15-0.20 ng/mL/cm³:**
- 🔴 Alto risco - biópsia geralmente indicada
- ✓ RM obrigatória antes da biópsia
- ✓ Biópsia dirigida por fusão RM-US se disponível
- ✓ Se PI-RADS ≥3, sensibilidade para câncer significativo é ~70%

**4. PSAD ≥ 0.20 ng/mL/cm³:**
- 🔴🔴 Risco muito alto - biópsia fortemente indicada
- ✓ Indicação de biópsia independente do resultado da RM
- ✓ Mesmo com PI-RADS ≤2, considerar biópsia sistemática
- ✓ Atenção especial ao Gleason score e volume tumoral

**5. PSAD ≥ 0.30 ng/mL/cm³ (HPB com RM negativa):**
- 🔴🔴🔴 Em contexto de HPB obstrutiva com RM negativa
- ✓ Especificidade de 93%, sensibilidade 65%
- ✓ Forte indicação de biópsia antes de considerar cirurgia para HPB

**Protocolo de Biópsia Recomendado:**
- **1ª linha:** Biópsia por fusão RM-ultrassom (se RM prévia)
- **2ª linha:** Biópsia sistemática 12-14 fragmentos se fusão indisponível
- **Seguimento pós-biópsia negativa:** Repetir PSA + PSAD em 6-12 meses

**Situações Especiais:**

**Próstatas Volumosas (>80 mL):**
- A PSAD perde acurácia
- Considerar peso maior para cinética do PSA e RM
- Limiares mais conservadores (favorecer biópsia se PSAD ≥ 0.15)

**Idade Avançada (>75 anos):**
- Avaliar expectativa de vida e comorbidades
- Considerar limiares mais altos (PSAD ≥ 0.20) para indicar biópsia
- Privilegiar qualidade de vida vs. diagnóstico precoce

**Vigilância Ativa de Câncer Confirmado:**
- PSAD > 0.15-0.20 pode sugerir reclassificação
- Aumento progressivo da PSAD indica progressão
- Trigger para biópsia de confirmação

**Integração com Outros Biomarcadores:**
- **PHI (Prostate Health Index):** Complementa PSAD na zona cinzenta
- **Cinética PSA:** Variação >20% ao ano aumenta risco
- **PSA livre/total < 15%:** Somado a PSAD >0.15 reforça indicação de biópsia

**Documentação Mandatória:**
- Volume prostático (método: elipsoide, planimetria)
- PSA sérico contemporâneo ao USG (idealmente <3 meses)
- Cálculo explícito: PSAD = PSA (ng/mL) / Volume (cm³)
- Contexto clínico: toque retal, sintomas urinários, histórico familiar',

    last_review = NOW(),
    updated_at = NOW()
WHERE id = '317acc85-3ce9-4f97-8e14-799354166f5e';

-- Verify update
DO $$
DECLARE
    v_item_name TEXT;
    v_article_count INTEGER;
BEGIN
    -- Check if item was updated
    SELECT name INTO v_item_name
    FROM score_items
    WHERE id = '317acc85-3ce9-4f97-8e14-799354166f5e'
    AND clinical_relevance IS NOT NULL;

    IF v_item_name IS NULL THEN
        RAISE EXCEPTION 'Item not found or not updated';
    END IF;

    -- Count linked articles
    SELECT COUNT(*) INTO v_article_count
    FROM article_score_items
    WHERE score_item_id = '317acc85-3ce9-4f97-8e14-799354166f5e';

    RAISE NOTICE '✓ SUCCESS! PSAD enrichment completed';
    RAISE NOTICE '  - Item: %', v_item_name;
    RAISE NOTICE '  - Articles linked: %', v_article_count;
    RAISE NOTICE '  - Last review: %', NOW()::DATE;
END $$;

COMMIT;
