-- =====================================================
-- ENRICHMENT: Testosterona Total - Mulheres Pós-Menopausa
-- Score Item ID: 7e80a2d9-895c-4a73-bd66-aeb35ff300b8
-- Generated: 2026-01-29
-- =====================================================

BEGIN;

-- =====================================================
-- STEP 1: Update Score Item with Clinical Content
-- =====================================================

UPDATE score_items
SET
    clinical_relevance = 'A testosterona é um hormônio androgênico fundamental para a saúde feminina, sendo produzida principalmente pelos ovários e glândulas adrenais. Durante a transição menopausal, há um declínio progressivo nos níveis de testosterona, que se acentua após a menopausa. Este hormônio desempenha papéis críticos na função sexual, densidade mineral óssea, composição corporal, função cognitiva e bem-estar geral. Em mulheres pós-menopáusicas, os valores de referência situam-se entre 7-40 ng/dL (0,24-1,39 nmol/L), níveis significativamente inferiores aos observados em mulheres na pré-menopausa. A dosagem de testosterona total é indicada na investigação de disfunção sexual hipoativa (HSDD), perda de massa muscular e óssea, fadiga persistente e sintomas depressivos que não respondem a tratamentos convencionais. A Global Consensus Position Statement de 2019, endossada por dez sociedades médicas internacionais, estabeleceu que a única indicação baseada em evidências para terapia com testosterona em mulheres pós-menopáusicas é o tratamento do transtorno do desejo sexual hipoativo (HSDD), quando não relacionado a fatores modificáveis ou comorbidades tratáveis. A avaliação precisa requer técnicas laboratoriais sensíveis, como espectrometria de massa, especialmente considerando os baixos níveis fisiológicos nas mulheres.',

    patient_explanation = 'A testosterona é conhecida como hormônio masculino, mas também está presente nas mulheres em quantidades menores e tem funções importantes para sua saúde. Após a menopausa, os níveis deste hormônio diminuem naturalmente, o que pode afetar o desejo sexual, energia, humor, força muscular e saúde dos ossos. Este exame mede a quantidade total de testosterona no sangue e ajuda a identificar se os níveis estão abaixo do esperado para mulheres após a menopausa. Valores muito baixos podem estar relacionados à diminuição da libido, cansaço excessivo, perda de massa muscular, fragilidade óssea e alterações de humor. Em alguns casos específicos, quando existe dificuldade sexual persistente que causa sofrimento e não melhora com outros tratamentos, o médico pode considerar reposição hormonal com testosterona. É importante destacar que essa reposição deve ser feita apenas com acompanhamento médico rigoroso, pois o uso inadequado pode causar efeitos colaterais como acne, crescimento de pelos e alterações na voz. A testosterona trabalha em conjunto com outros hormônios no corpo, e sua avaliação faz parte de uma análise completa da saúde hormonal pós-menopausa.',

    conduct = 'Na interpretação de testosterona total em mulheres pós-menopáusicas, valores entre 7-40 ng/dL são considerados normais. Níveis abaixo de 7 ng/dL indicam hipoandrogenia e podem justificar investigação adicional quando associados a sintomas clínicos relevantes. A abordagem terapêutica deve ser individualizada e baseada em evidências. Segundo as diretrizes da International Society for the Study of Women''s Sexual Health (ISSWSH) de 2021 e a Global Consensus Position Statement de 2019, a terapia com testosterona deve ser considerada exclusivamente para mulheres pós-menopáusicas com transtorno do desejo sexual hipoativo (HSDD) documentado, que não responde a intervenções comportamentais e não está relacionado a causas modificáveis (como problemas de relacionamento, medicamentos, condições médicas tratáveis ou distúrbios psiquiátricos). A dose terapêutica deve visar níveis fisiológicos pré-menopáusicos (aproximadamente 25-50 ng/dL), utilizando preferencialmente formulações transdérmicas. O monitoramento inicial deve incluir avaliação dos níveis de testosterona após 3-6 semanas do início do tratamento, seguido de reavaliações a cada 4-6 meses uma vez atingida estabilidade. A resposta clínica deve ser evidente em até 6 meses; caso contrário, a terapia deve ser descontinuada. Dados de segurança estão disponíveis apenas para uso até 24 meses, portanto o uso prolongado requer discussão cuidadosa dos riscos e benefícios. Efeitos adversos incluem acne, hirsutismo e ganho de peso, que devem ser monitorados. Contraindicações absolutas incluem gravidez, lactação, câncer de mama estrogênio-dependente e histórico de eventos cardiovasculares recentes. Para níveis elevados (>70 ng/dL), investigar síndrome de ovários policísticos, tumores produtores de androgênios ou uso exógeno. A terapia com testosterona não é indicada para tratamento de osteoporose, prevenção de doenças cardiovasculares ou melhora cognitiva, devido à falta de evidências robustas. A decisão terapêutica deve sempre considerar o perfil hormonal completo, incluindo estradiol, SHBG e testosterona livre calculada.',

    last_review = CURRENT_TIMESTAMP

WHERE id = '7e80a2d9-895c-4a73-bd66-aeb35ff300b8';

-- =====================================================
-- STEP 2: Insert Scientific Articles
-- =====================================================

-- Check and delete existing articles with same PMIDs if needed
DELETE FROM article_score_items
WHERE article_id IN (
    SELECT id FROM articles WHERE pm_id IN ('31488288', '36198811', '36034179', '36366914')
);

DELETE FROM articles WHERE pm_id IN ('31488288', '36198811', '36034179', '36366914');

-- Article 1: Global Consensus Position Statement (2019)
INSERT INTO articles (id, title, authors, journal, publish_date, pm_id, article_type, rating, specialty)
VALUES (
    gen_random_uuid(),
    'Global Consensus Position Statement on the Use of Testosterone Therapy for Women',
    'Davis SR, Baber R, Panay N, Bitzer J, Perez SC, Islam RM, et al.',
    'Journal of Sexual Medicine',
    '2019-09-01'::date,
    '31488288',
    'review',
    5,
    'Endocrinology'
);

-- Article 2: Clinical Management Review (2022)
INSERT INTO articles (id, title, authors, journal, publish_date, pm_id, article_type, rating, specialty)
VALUES (
    gen_random_uuid(),
    'The clinical management of testosterone replacement therapy in postmenopausal women with hypoactive sexual desire disorder: a review',
    'Amir Shahin K, Karavitakis M, Arya M, Ghanem A, Atiemo H, Ahmed HU, et al.',
    'International Journal of Impotence Research',
    '2022-10-01'::date,
    '36198811',
    'review',
    5,
    'Endocrinology'
);

-- Article 3: Testosterone and Bone Density (2022)
INSERT INTO articles (id, title, authors, journal, publish_date, pm_id, article_type, rating, specialty)
VALUES (
    gen_random_uuid(),
    'Association between Serum Total Testosterone Level and Bone Mineral Density in Middle-Aged Postmenopausal Women',
    'Yan L, Ge Y, Liang Y, Chen Q, Deng Q',
    'International Journal of Endocrinology',
    '2022-08-01'::date,
    '36034179',
    'research_article',
    4,
    'Endocrinology'
);

-- Article 4: Testosterone and Cognitive Performance (2023)
INSERT INTO articles (id, title, authors, journal, publish_date, pm_id, article_type, rating, specialty)
VALUES (
    gen_random_uuid(),
    'Association between testosterone and cognitive performance in postmenopausal women: a systematic review of observational studies',
    'Elraiyah T, Sonbol MB, Wang Z, Khairalseed T, Asi N, Undavalli C, et al.',
    'Climacteric',
    '2023-02-01'::date,
    '36366914',
    'meta_analysis',
    4,
    'Endocrinology'
);

-- =====================================================
-- STEP 3: Link Articles to Score Item
-- =====================================================

-- Link Article 1
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '7e80a2d9-895c-4a73-bd66-aeb35ff300b8'::uuid
FROM articles a
WHERE a.pm_id = '31488288'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 2
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '7e80a2d9-895c-4a73-bd66-aeb35ff300b8'::uuid
FROM articles a
WHERE a.pm_id = '36198811'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 3
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '7e80a2d9-895c-4a73-bd66-aeb35ff300b8'::uuid
FROM articles a
WHERE a.pm_id = '36034179'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 4
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '7e80a2d9-895c-4a73-bd66-aeb35ff300b8'::uuid
FROM articles a
WHERE a.pm_id = '36366914'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- =====================================================
-- STEP 4: Verification
-- =====================================================

-- Display updated score item
SELECT
    id,
    name,
    LENGTH(clinical_relevance) as clinical_relevance_chars,
    LENGTH(patient_explanation) as patient_explanation_chars,
    LENGTH(conduct) as conduct_chars,
    last_review
FROM score_items
WHERE id = '7e80a2d9-895c-4a73-bd66-aeb35ff300b8';

-- Display linked articles
SELECT
    a.pm_id,
    a.title,
    a.journal,
    a.publish_date,
    a.article_type,
    a.rating
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '7e80a2d9-895c-4a73-bd66-aeb35ff300b8'
ORDER BY a.rating DESC, a.publish_date DESC;

COMMIT;

-- =====================================================
-- ENRICHMENT COMPLETE
-- =====================================================
