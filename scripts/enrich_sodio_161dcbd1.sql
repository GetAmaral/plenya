-- ================================================================
-- ENRICHMENT: Sódio (ID: 161dcbd1-6694-4175-958b-2b260ae48a40)
-- Generated: 2026-01-29
-- Category: Exames Laboratoriais
-- ================================================================

BEGIN;

-- ================================================================
-- STEP 1: Insert Scientific Articles
-- ================================================================

-- Article 1: Hyponatraemia-treatment standard 2024
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    doi,
    article_type
) VALUES (
    gen_random_uuid(),
    'Hyponatraemia-treatment standard 2024',
    'Spasovski G',
    'Nephrology Dialysis Transplantation',
    '2024-09-27'::date,
    '39009016',
    '10.1093/ndt/gfae162',
    'review'
) ON CONFLICT (doi) DO UPDATE SET
    pm_id = EXCLUDED.pm_id,
    title = EXCLUDED.title,
    authors = EXCLUDED.authors;

-- Article 2: Treatment of hyponatremia: comprehension and best clinical practice
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    doi,
    article_type
) VALUES (
    gen_random_uuid(),
    'Treatment of hyponatremia: comprehension and best clinical practice',
    'Sumi H, Tominaga N, Fujita Y, Verbalis JG',
    'Clinical and Experimental Nephrology',
    '2025-01-23'::date,
    '39847310',
    '10.1007/s10157-024-02606-3',
    'review'
) ON CONFLICT (doi) DO UPDATE SET
    pm_id = EXCLUDED.pm_id,
    title = EXCLUDED.title,
    authors = EXCLUDED.authors;

-- Article 3: Correction Rates and Clinical Outcomes in Severe Hyponatremia (Meta-Analysis)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    doi,
    article_type
) VALUES (
    gen_random_uuid(),
    'Correction Rates and Clinical Outcomes in Hospitalized Adults With Severe Hyponatremia: A Systematic Review and Meta-Analysis',
    'Ayus JC, Moritz ML, Fuentes NA, Mejia JR, Alfonso JM, Shin S, Fralick M, Ciapponi A',
    'JAMA Internal Medicine',
    '2025-01-01'::date,
    '39556338',
    '10.1001/jamainternmed.2024.5981',
    'meta_analysis'
) ON CONFLICT (doi) DO UPDATE SET
    pm_id = EXCLUDED.pm_id,
    title = EXCLUDED.title,
    authors = EXCLUDED.authors;

-- Article 4: Syndrome of Inappropriate Antidiuretic Hormone Secretion (StatPearls - no DOI)
-- Check if article with this pm_id already exists
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM articles WHERE pm_id = '29939554') THEN
        INSERT INTO articles (
            id,
            title,
            authors,
            journal,
            publish_date,
            pm_id,
            doi,
            article_type
        ) VALUES (
            gen_random_uuid(),
            'Syndrome of Inappropriate Antidiuretic Hormone Secretion',
            'Yasir M, Mechanic OJ',
            'StatPearls',
            '2023-03-06'::date,
            '29939554',
            NULL,
            'review'
        );
    END IF;
END $$;

-- ================================================================
-- STEP 2: Update Score Item with Clinical Content
-- ================================================================

UPDATE score_items
SET
    clinical_relevance = 'O sódio sérico é o principal cátion extracelular e determinante fundamental da osmolalidade plasmática, sendo essencial para a manutenção da homeostase hidroeletrolítica. A concentração normal varia de 135 a 145 mEq/L, e alterações nesta faixa refletem distúrbios no balanço de água mais do que de sódio propriamente dito. A hiponatremia (Na+ <135 mEq/L) é o distúrbio eletrolítico mais comum em pacientes hospitalizados, associada a aumento significativo de morbimortalidade, tempo de internação e custos hospitalares. Pode manifestar-se desde casos assintomáticos até quadros graves com alterações neurológicas (confusão, convulsões, coma) devido ao edema cerebral. A hipernatremia (Na+ >145 mEq/L), embora menos frequente, também apresenta alta mortalidade, especialmente em idosos e pacientes críticos, causando desidratação celular cerebral. As causas de hiponatremia incluem SIADH (síndrome da secreção inapropriada de hormônio antidiurético), insuficiência cardíaca, cirrose, nefropatias, uso de diuréticos e polidipsia psicogênica. A hipernatremia geralmente resulta de perda de água livre (diabetes insipidus, perdas gastrointestinais) ou ganho excessivo de sódio. O diagnóstico diferencial requer avaliação da volemia clínica, osmolalidade sérica e urinária, sódio urinário e função renal. Estudos recentes de 2024-2025 demonstram que a velocidade de correção da hiponatremia grave permanece controversa: meta-análise publicada no JAMA Internal Medicine (2025) sugere que correção mais rápida associa-se a menor mortalidade hospitalar (32 mortes a menos por 1000 pacientes tratados), desafiando diretrizes tradicionais que limitam a velocidade de correção para prevenir síndrome de desmielinização osmótica. A avaliação do sódio sérico é fundamental no manejo de pacientes críticos, idosos, portadores de doenças crônicas e usuários de medicações que afetam o balanço hidroeletrolítico.',

    patient_explanation = 'O sódio é um mineral essencial presente no sangue que ajuda a controlar a quantidade de água no seu corpo e é fundamental para o funcionamento dos nervos e músculos. O valor normal no sangue fica entre 135 e 145 mEq/L. Quando o sódio está baixo (hiponatremia, abaixo de 135), geralmente significa que há excesso de água no corpo em relação ao sódio, e não necessariamente falta de sal. Isso pode causar sintomas como náuseas, dor de cabeça, confusão mental, fraqueza muscular e, em casos graves, convulsões ou coma. As causas mais comuns incluem uso de diuréticos ("remédios para urinar"), insuficiência cardíaca, problemas renais, cirrose hepática, beber água em excesso e uma condição chamada SIADH (quando o corpo retém água demais). Quando o sódio está alto (hipernatremia, acima de 145), significa desidratação - seu corpo tem pouca água em relação ao sal. Os sintomas incluem sede intensa, boca seca, confusão e sonolência. Isso pode acontecer por perda excessiva de líquidos (diarreia, vômitos, suor excessivo), diabetes não controlada ou não beber água suficiente, sendo especialmente perigoso em idosos. Alterações no sódio são sérias e requerem investigação médica imediata. O tratamento depende da causa: pode incluir ajuste de medicações, restrição ou reposição de líquidos, e em casos graves, internação hospitalar para correção cuidadosa dos níveis, pois corrigir muito rápido ou muito devagar pode ser perigoso. Seu médico pode solicitar outros exames (como osmolalidade, sódio urinário) para identificar a causa exata e definir o melhor tratamento.',

    conduct = 'INVESTIGAÇÃO DIAGNÓSTICA: Diante de alteração nos níveis de sódio sérico, realizar avaliação criteriosa incluindo: (1) História clínica detalhada investigando sintomas neurológicos (confusão, letargia, convulsões), volemia (edema, ascite, desidratação), uso de medicações (diuréticos, ISRS, carbamazepina, AINEs, inibidores de bomba de prótons), comorbidades (insuficiência cardíaca, cirrose, nefropatia, hipotireoidismo, insuficiência adrenal), ingestão hídrica e perdas (vômitos, diarreia, poliúria); (2) Exame físico avaliando estado de hidratação, volemia (PVC, turgor cutâneo, mucosas), sinais neurológicos e edemas; (3) Exames complementares essenciais: osmolalidade sérica e urinária, sódio urinário, potássio sérico, função renal (ureia, creatinina), glicemia, TSH, cortisol matinal, gasometria se sintomático. Para hiponatremia, classificar quanto à: severidade (leve 130-134, moderada 125-129, grave <125 mEq/L), duração (aguda <48h, crônica >48h ou desconhecida), sintomatologia (assintomática, moderada, grave) e volemia (hipovolêmica, euvolêmica, hipervolêmica). CONDUTA NA HIPONATREMIA: (1) HIPONATREMIA GRAVE SINTOMÁTICA: Emergência médica - administrar solução salina hipertônica 3% em bolus de 150 mL em 20 minutos (pode repetir até 2x se sintomas persistirem), monitorização em UTI com dosagens de sódio a cada 2-4h, meta de elevação 4-6 mEq/L nas primeiras 6h ou até resolução dos sintomas, limite máximo de correção 8 mEq/L em 24h e 18 mEq/L em 48h para evitar síndrome de desmielinização osmótica (especialmente em pacientes com hiponatremia crônica, alcoolismo, desnutrição ou hipocalemia); nota: meta-análise recente (JAMA 2025) sugere que correções mais rápidas podem associar-se a melhor desfecho, mas diretrizes europeias 2024 mantêm recomendação de cautela; (2) HIPONATREMIA ASSINTOMÁTICA LEVE-MODERADA: Investigar e tratar causa base - SIADH (restrição hídrica 500 mL/dia ajustada conforme resposta, considerar ureia oral 15-60g/dia ou tolvaptan 15mg/dia se refratária), hipovolêmica (reposição volêmica com SF 0,9%), hipervolêmica (restrição hídrica + sal, diuréticos de alça), hipotireoidismo (levotiroxina), insuficiência adrenal (hidrocortisona); (3) Suspender medicações causadoras quando possível; (4) Aumentar aporte de solutos (sal, proteínas). CONDUTA NA HIPERNATREMIA: (1) Calcular déficit de água livre: Déficit = 0,6 x peso(kg) x [(Na atual/140) - 1]; (2) Reposição gradual: corrigir no máximo 10-12 mEq/L em 24h (0,5 mEq/L/h) para evitar edema cerebral, preferir água via oral se possível, ou soluções hipotônicas IV (SF 0,45% ou SG 5%); (3) Investigar e tratar causa: diabetes insipidus (desmopressina), perdas gastrointestinais (correção volêmica), reposição hídrica inadequada (aumentar oferta). SEGUIMENTO: Monitorização laboratorial conforme severidade (grave: 2-4h, moderada: 6-12h, leve: 24-48h), ajuste terapêutico baseado em resposta clínico-laboratorial, investigação etiológica aprofundada se causa não evidente, encaminhamento a endocrinologista/nefrologista em casos complexos (SIADH refratária, diabetes insipidus, insuficiência adrenal). Educar paciente sobre sinais de alerta e importância da adesão ao plano terapêutico.',

    last_review = CURRENT_DATE
WHERE id = '161dcbd1-6694-4175-958b-2b260ae48a40';

-- ================================================================
-- STEP 3: Link Articles to Score Item
-- ================================================================

-- Link Article 1
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '161dcbd1-6694-4175-958b-2b260ae48a40'
FROM articles a
WHERE a.pm_id = '39009016'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 2
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '161dcbd1-6694-4175-958b-2b260ae48a40'
FROM articles a
WHERE a.pm_id = '39847310'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 3
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '161dcbd1-6694-4175-958b-2b260ae48a40'
FROM articles a
WHERE a.pm_id = '39556338'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 4
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '161dcbd1-6694-4175-958b-2b260ae48a40'
FROM articles a
WHERE a.pm_id = '29939554'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- ================================================================
-- VERIFICATION QUERY
-- ================================================================

-- Verify enrichment
SELECT
    si.id,
    si.name,
    si.code,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    si.last_review,
    COUNT(DISTINCT asis.article_id) as linked_articles
FROM score_items si
LEFT JOIN article_score_items asis ON si.id = asis.score_item_id
WHERE si.id = '161dcbd1-6694-4175-958b-2b260ae48a40'
GROUP BY si.id, si.name, si.code, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;

-- Show linked articles
SELECT
    a.pm_id,
    a.title,
    a.authors,
    a.journal,
    a.publish_date,
    a.article_type
FROM articles a
INNER JOIN article_score_items asis ON a.id = asis.article_id
WHERE asis.score_item_id = '161dcbd1-6694-4175-958b-2b260ae48a40'
ORDER BY a.publish_date DESC;

COMMIT;

-- ================================================================
-- EXECUTION SUMMARY
-- ================================================================
-- Score Item: Sódio (161dcbd1-6694-4175-958b-2b260ae48a40)
-- Articles Added: 4 peer-reviewed articles (2023-2025)
-- Clinical Content: Comprehensive PT-BR content with 2024-2025 guidelines
-- Focus: Sodium homeostasis, hyponatremia, hypernatremia, SIADH
-- ================================================================
