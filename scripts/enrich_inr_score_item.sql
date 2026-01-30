-- ============================================================================
-- ENRICHMENT: Tempo de Protrombina (INR)
-- Score Item ID: 459b1285-86d6-408f-9735-029dd00e67b6
-- ============================================================================
-- Data: 2026-01-29
-- Artigos: 4 artigos peer-reviewed (2024-2025)
-- Caracteres: clinical_relevance ~1800, patient_explanation ~1200, conduct ~2200
-- ============================================================================

BEGIN;

-- ============================================================================
-- STEP 1: Insert Scientific Articles
-- ============================================================================

-- Article 1: StatPearls - INR Monitoring (2025)
INSERT INTO articles (id, pm_id, title, authors, journal, publish_date, article_type, abstract)
VALUES (
    gen_random_uuid(),
    '29939529',
    'International Normalized Ratio: Assessment, Monitoring, and Clinical Implications',
    'Shikdar S, Vashisht R, Zubair M, Bhattacharya PT',
    'StatPearls',
    '2025-02-14'::date,
    'review',
    'Comprehensive review on INR monitoring for anticoagulation therapy, covering assessment methods, therapeutic ranges, and clinical implications. Updated February 2025 with current guidelines for warfarin management and monitoring protocols.'
)
ON CONFLICT DO NOTHING
RETURNING id;

-- Article 2: DOAC Guidelines 2025
INSERT INTO articles (id, pm_id, title, authors, journal, publish_date, article_type, abstract)
VALUES (
    gen_random_uuid(),
    '40448969',
    '2025 Guidelines for direct oral anticoagulants: a practical guidance on the prescription, laboratory testing, peri-operative and bleeding management',
    'Tran HA, et al',
    'Internal Medicine Journal',
    '2025-05-31'::date,
    'review',
    'Updated guidelines from the Thrombosis and Haemostasis Society of Australia and New Zealand (THANZ) comparing direct oral anticoagulants (DOACs) with warfarin/VKAs. Addresses advantages of DOACs including predictable effect, no routine monitoring requirement, and comparison with INR-based warfarin management.'
)
ON CONFLICT DO NOTHING
RETURNING id;

-- Article 3: Home INR Monitoring (2024)
INSERT INTO articles (id, pm_id, title, authors, journal, publish_date, article_type, abstract)
VALUES (
    gen_random_uuid(),
    '38100006',
    'Outcomes of Warfarin Home INR Monitoring vs Office-Based Monitoring: a Retrospective Claims-Based Analysis',
    'Van Beek EJR, et al',
    'Journal of General Internal Medicine',
    '2024-01-15'::date,
    'research_article',
    'Retrospective claims-based analysis comparing patient self-testing (PST) versus laboratory-based INR monitoring. Demonstrates real-world outcomes and effectiveness of home INR monitoring technologies for warfarin management, with implications for time in therapeutic range and patient safety.'
)
ON CONFLICT DO NOTHING
RETURNING id;

-- Article 4: TTR-INR Protocol (2024)
INSERT INTO articles (id, pm_id, title, authors, journal, publish_date, article_type, abstract)
VALUES (
    gen_random_uuid(),
    '38773162',
    'Utility of TTR-INR guided warfarin adjustment protocol to improve time in therapeutic range in patients with atrial fibrillation receiving warfarin',
    'Kosum S, et al',
    'Scientific Reports',
    '2024-05-22'::date,
    'research_article',
    'Prospective cohort study demonstrating that TTR-INR guided warfarin adjustment protocol significantly improves time in therapeutic range from 65% to 80% in atrial fibrillation patients. Proportion of patients achieving adequate anticoagulation control increased from 47% to 88% over 12 months.'
)
ON CONFLICT DO NOTHING
RETURNING id;

-- ============================================================================
-- STEP 2: Link Articles to Score Item
-- ============================================================================

INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '459b1285-86d6-408f-9735-029dd00e67b6'::uuid
FROM articles a
WHERE a.pm_id IN ('29939529', '40448969', '38100006', '38773162')
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- ============================================================================
-- STEP 3: Update Score Item with Clinical Content
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'O Tempo de Protrombina expresso como INR (International Normalized Ratio) é o método padronizado internacionalmente para monitoramento da anticoagulação oral com antagonistas da vitamina K (varfarina). O INR normaliza as variações entre diferentes reagentes de tromboplastina, permitindo comparação consistente dos resultados entre laboratórios. Em indivíduos saudáveis, o INR normal varia de 0,8 a 1,2. Durante terapia anticoagulante, o alvo terapêutico é 2,0-3,0 para a maioria das indicações (fibrilação atrial, tromboembolismo venoso, prevenção de AVC), e 2,5-3,5 para válvulas mecânicas cardíacas. A qualidade do controle anticoagulante é medida pelo tempo na faixa terapêutica (TTR - Time in Therapeutic Range): TTR >70% indica controle adequado, TTR 60-70% é subótimo, e TTR <60% representa controle inadequado com maior risco de eventos tromboembólicos ou hemorrágicos. Estudos recentes (2024) demonstram que protocolos guiados por TTR-INR melhoram significativamente o controle anticoagulante, elevando o TTR de 65% para 80% e a proporção de pacientes adequadamente anticoagulados de 47% para 88%. Diretrizes de 2025 destacam que, embora anticoagulantes orais diretos (DOACs) ofereçam vantagens como efeito previsível e dispensa de monitoramento rotineiro, a varfarina permanece essencial em situações específicas (válvulas mecânicas, insuficiência renal grave, extremos de peso corporal), tornando o monitoramento do INR clinicamente relevante para milhões de pacientes globalmente.',

    patient_explanation = 'O INR (sigla em inglês para "razão normalizada internacional") é um exame de sangue que mede quanto tempo seu sangue leva para coagular. Valores normais ficam entre 0,8 e 1,2 em pessoas que não tomam anticoagulantes. Se você toma varfarina (Marevan® ou Coumadin®), seu médico ajusta a dose para manter o INR em uma faixa-alvo específica, geralmente entre 2,0 e 3,0. Isso significa que seu sangue está levando 2 a 3 vezes mais tempo para coagular do que o normal, o que previne formação de coágulos perigosos, mas sem aumentar demais o risco de sangramentos. Pessoas com válvulas cardíacas mecânicas precisam de anticoagulação mais intensa (INR 2,5-3,5). O monitoramento regular do INR é essencial: muito baixo (abaixo de 2,0) significa risco de formar coágulos e ter AVC ou trombose; muito alto (acima de 4,0) aumenta risco de sangramentos graves. A frequência ideal de testes é geralmente a cada 3-4 semanas quando o INR está estável, mas pode ser semanal ou até diária durante ajustes de dose. Muitos fatores afetam o INR: mudanças na alimentação (especialmente vegetais verdes ricos em vitamina K como couve, brócolis, espinafre), novos medicamentos, doenças do fígado, diarreia ou vômitos. Por isso, é fundamental comunicar ao médico qualquer alteração em sua saúde ou medicações. Dispositivos modernos permitem monitoramento domiciliar do INR, oferecendo maior comodidade e resultados imediatos.',

    conduct = 'MONITORAMENTO INICIAL: Ao iniciar varfarina, realizar INR diariamente até atingir faixa terapêutica por 2 dias consecutivos. Após estabilização, espaçar gradualmente para semanal, quinzenal e finalmente mensal (se TTR >70%). ALVOS TERAPÊUTICOS: INR 2,0-3,0 para fibrilação atrial, TVP/TEP, prevenção de AVC em cardiopatias; INR 2,5-3,5 para próteses valvares mecânicas (posição mitral ou aórtica de alto risco); INR 1,5-2,0 raramente usado em situações especiais (baixa intensidade). AJUSTES DE DOSE: Para INR <1,5: aumentar dose semanal em 10-20%, repetir INR em 3-7 dias. Para INR 1,5-1,9: aumentar dose em 5-15%. Para INR 2,0-3,0 (alvo padrão): manter dose, próximo INR em 4 semanas se estável. Para INR 3,1-4,0: reduzir dose em 5-15%, repetir em 1-2 semanas. Para INR 4,1-5,0 sem sangramento: omitir 1 dose, reduzir dose semanal em 10-20%, repetir em 3-5 dias. Para INR >5,0: suspender varfarina, considerar vitamina K oral (1-2,5mg para INR 5-9; 5mg para INR >9), avaliar sinais de sangramento, repetir INR diariamente até <3,0. REVERSÃO DE EMERGÊNCIA (sangramento grave): vitamina K 10mg IV lento + concentrado de complexo protrombínico (CCP) ou plasma fresco congelado. MONITORAMENTO DOMICILIAR: Pacientes selecionados (educados, aderentes, capazes) podem usar dispositivos point-of-care para autoteste, mantendo comunicação com equipe médica. Evidências de 2024 mostram equivalência de desfechos clínicos entre monitoramento domiciliar e laboratorial. INTERAÇÕES MEDICAMENTOSAS: Sempre verificar interações antes de prescrever novos fármacos; antibióticos (azólicos, macrolídeos, sulfametoxazol), amiodarona e anti-inflamatórios são particularmente problemáticos. ORIENTAÇÕES DIETÉTICAS: Manter consumo consistente de vitamina K (evitar grandes variações, não é necessário eliminar vegetais verdes). TRANSIÇÃO PARA DOACs: Considerar em pacientes com TTR persistentemente <60%, dificuldade de acesso a monitoramento, múltiplas interações medicamentosas ou preferência do paciente (desde que não haja contraindicação como válvula mecânica ou clearance creatinina <15ml/min).',

    last_review = CURRENT_TIMESTAMP

WHERE id = '459b1285-86d6-408f-9735-029dd00e67b6';

-- ============================================================================
-- VERIFICATION QUERIES
-- ============================================================================

-- Verify article insertion
SELECT
    a.pm_id,
    a.title,
    a.journal,
    a.publish_date,
    a.article_type
FROM articles a
WHERE a.pm_id IN ('29939529', '40448969', '38100006', '38773162')
ORDER BY a.publish_date DESC;

-- Verify article-score_item links
SELECT
    si.name,
    COUNT(asi.article_id) as num_articles
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '459b1285-86d6-408f-9735-029dd00e67b6'
GROUP BY si.name;

-- Verify character counts
SELECT
    name,
    LENGTH(clinical_relevance) as clinical_relevance_chars,
    LENGTH(patient_explanation) as patient_explanation_chars,
    LENGTH(conduct) as conduct_chars,
    last_review
FROM score_items
WHERE id = '459b1285-86d6-408f-9735-029dd00e67b6';

COMMIT;

-- ============================================================================
-- EXPECTED OUTPUT:
-- - 4 articles inserted/updated (PMIDs: 29939529, 40448969, 38100006, 38773162)
-- - 4 article-score_item links created
-- - clinical_relevance: ~1800 chars
-- - patient_explanation: ~1200 chars
-- - conduct: ~2200 chars
-- - last_review: 2026-01-29
-- ============================================================================
