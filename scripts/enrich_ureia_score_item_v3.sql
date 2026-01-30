-- ============================================================================
-- ENRICHMENT V3: Score Item "Ureia" - Optimized Character Counts
-- Generated: 2026-01-29
-- Target: clinical_relevance (1500-2000), patient_explanation (1000-1500), conduct (1500-2500)
-- ============================================================================

BEGIN;

UPDATE score_items
SET
    clinical_relevance = 'A ureia, mensurada como nitrogênio ureico no sangue (BUN), é um marcador fundamental da função renal e representa o produto final do metabolismo proteico hepático. Cerca de 85% da ureia é eliminada pelos rins através da filtração glomerular, tornando-a um indicador sensível da TFG. Valores de referência situam-se entre 15-40 mg/dL (ou 7-20 mg/dL quando expressos como BUN). O Guideline KDIGO 2024 estabelece que a ureia deve ser avaliada em conjunto com creatinina e TFG estimada para estratificação de risco em DRC. Estudos prospectivos publicados entre 2019-2024 demonstram que níveis elevados de ureia associam-se independentemente a progressão da DRC estágios 3-5 e aumento de eventos cardiovasculares, mesmo após ajuste para TFGe. A razão ureia/creatinina (BUN/Cr) é crucial na diferenciação de azotemia: razão >20:1 sugere azotemia pré-renal (desidratação, hipoperfusão), enquanto razão 10-15:1 indica lesão renal intrínseca. Na DRC, ureia elevada representa acúmulo de toxinas urêmicas com consequências sistêmicas: anemia por supressão medular, disfunção endotelial, sarcopenia e síndrome urêmica. A razão urina/plasma de ureia correlaciona-se com capacidade de concentração tubular, funcionando como biomarcador de função tubular e preditor de progressão da DRC. Dados do estudo CKD-REIN demonstraram que pacientes no tertil mais elevado de ureia apresentam risco significativamente aumentado de eventos cardiovasculares ateromatosos e não-ateromatosos. O estudo BMC Nephrology com seguimento de 1893 pacientes japoneses mostrou que BUN elevado associa-se independentemente a desfechos renais adversos (início de terapia renal substitutiva ou declínio ≥50% da TFG). Valores persistentemente elevados acima de 40 mg/dL exigem investigação etiológica e medidas terapêuticas para prevenir complicações renais e cardiovasculares.',

    patient_explanation = 'A ureia é uma substância produzida pelo fígado quando o corpo quebra as proteínas dos alimentos. Seus rins funcionam como filtros que retiram a ureia do sangue e a eliminam pela urina. Quando a ureia está alta, pode significar que seus rins não estão filtrando bem, você está desidratado, ou está comendo muita proteína. Valores normais ficam entre 15-40 mg/dL. Se sua ureia estiver elevada, o médico pedirá outros exames como creatinina e taxa de filtração renal para investigar melhor. Causas comuns de ureia alta: desidratação, insuficiência renal, dieta muito rica em proteínas, sangramento no estômago ou intestino, uso de anti-inflamatórios ou corticoides, e insuficiência cardíaca. Se a ureia estiver baixa, pode indicar desnutrição, doença hepática grave, ou gravidez. Sintomas de ureia muito elevada incluem: cansaço, perda de apetite, náuseas, confusão mental, sonolência e inchaço nas pernas. O tratamento depende da causa: reidratação se desidratação, ajuste de medicamentos, controle da pressão e diabetes na doença renal crônica, restrição moderada de proteínas, e em casos graves pode ser necessário diálise.',

    conduct = 'INTERPRETAÇÃO: Valores de referência: 15-40 mg/dL (BUN: 7-20 mg/dL). Calcular razão BUN/Creatinina: 10-20 considerada normal. Razão >20 sugere azotemia pré-renal; <10-15 sugere lesão renal intrínseca. ESTRATIFICAÇÃO KDIGO 2024: Integrar ureia com creatinina, TFGe (CKD-EPI 2021) e albuminúria para classificação CGA. INVESTIGAÇÃO SE UREIA ELEVADA: 1) Avaliar hidratação e fatores pré-renais (ICC, hipotensão, AINES, IECA/BRA, diuréticos). 2) Revisar dieta proteica (>1.5g/kg/dia pode elevar ureia sem disfunção renal). 3) Investigar sangramento GI oculto se desproporção ureia/creatinina. 4) Solicitar: TFGe, creatinina, EAS com albuminúria/proteinúria, eletrólitos (Na, K, Ca, P, bicarbonato), gasometria venosa. 5) USG renal com Doppler se TFGe <60 mL/min/1.73m² para avaliar tamanho renal, hidronefrose, cistos, fluxo arterial. MANEJO DRC (KDIGO 2024): Estágios G3a-G5 (TFGe <60): controle PA <130/80 mmHg, iSGLT2 se TFGe ≥20, IECA/BRA se albuminúria >30 mg/g, restrição proteica 0.8 g/kg/dia (0.6 g/kg/dia em G4-G5). Monitorar ureia, creatinina, TFGe, eletrólitos, PTH trimestralmente em G3-G4; mensalmente em G5. Encaminhar nefrologia se: TFGe <30, progressão >5 mL/min/ano, albuminúria >300 mg/g, hipercalemia >5.5 mEq/L refratária, síndrome urêmica. AZOTEMIA PRÉ-RENAL: Reposição volêmica IV com cristaloides, suspender nefrotóxicos e diuréticos, otimizar hemodinâmica. Reavaliar ureia/creatinina em 24-48h. SEGUIMENTO: Ureia >60 mg/dL requer investigação completa mesmo com creatinina normal (pode indicar estados hipercatabólicos, sangramento GI oculto, início de disfunção renal).',

    updated_at = CURRENT_TIMESTAMP
WHERE id = '6111a95b-6b98-4534-b623-80601070d80d';

-- Verify final character counts
SELECT
    name,
    LENGTH(clinical_relevance) as clinical_chars,
    LENGTH(patient_explanation) as patient_chars,
    LENGTH(conduct) as conduct_chars,
    CASE WHEN LENGTH(clinical_relevance) BETWEEN 1500 AND 2000 THEN '✓' ELSE '✗' END as clinical_ok,
    CASE WHEN LENGTH(patient_explanation) BETWEEN 1000 AND 1500 THEN '✓' ELSE '✗' END as patient_ok,
    CASE WHEN LENGTH(conduct) BETWEEN 1500 AND 2500 THEN '✓' ELSE '✗' END as conduct_ok
FROM score_items
WHERE id = '6111a95b-6b98-4534-b623-80601070d80d';

-- Verify article linkages with PubMed IDs
SELECT
    'Ureia' as score_item,
    a.pm_id,
    LEFT(a.title, 80) || '...' as article_title,
    a.journal,
    a.publish_date,
    a.article_type
FROM score_items si
JOIN article_score_items asi ON si.id = asi.score_item_id
JOIN articles a ON asi.article_id = a.id
WHERE si.id = '6111a95b-6b98-4534-b623-80601070d80d'
  AND a.pm_id IS NOT NULL
ORDER BY a.publish_date DESC;

COMMIT;

-- ============================================================================
-- EXPECTED RESULTS
-- ============================================================================
-- Clinical relevance: ~1850 chars ✓
-- Patient explanation: ~1150 chars ✓
-- Conduct: ~1900 chars ✓
-- PubMed articles: 4 (PMIDs: 38490803, 36356680, 35544273, 30940101)
-- ============================================================================
