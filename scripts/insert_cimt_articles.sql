-- Inserção de artigos científicos sobre CIMT Carótidas
-- Data: 2026-01-28

-- Artigo 1: UK Biobank Study 2024
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    abstract,
    original_link,
    article_type,
    specialty,
    keywords
) VALUES (
    'Carotid intima-media thickness, cardiovascular disease, and risk factors in 29,000 UK Biobank adults',
    'UK Biobank Research Team',
    'PMC - PubMed Central',
    '2024-06-01',
    'en',
    NULL,
    'PMC12162042',
    'In a prospective cohort study of 29,292 participants free from cardiovascular disease at baseline, higher cIMT values were associated with an increased risk of coronary heart disease and myocardial infarction. The study demonstrated that carotid intima-media thickness is a well-established surrogate marker of atherosclerosis and the impact of cardiometabolic risk factor burden on cIMT and future MACE risk.',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC12162042/',
    'research_article',
    'Cardiologia',
    '["carotid intima-media thickness", "cardiovascular disease", "atherosclerosis", "UK Biobank", "risk assessment", "coronary heart disease"]'::jsonb
) ON CONFLICT (doi) DO NOTHING
RETURNING id;

-- Artigo 2: Clinical Applications Review
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    abstract,
    original_link,
    article_type,
    specialty,
    keywords
) VALUES (
    'Clinical and Research Applications of Carotid Intima-Media Thickness',
    'Multiple Authors',
    'PMC - PubMed Central',
    '2023-01-01',
    'en',
    NULL,
    'PMC2691892',
    'Comprehensive review of clinical and research applications of carotid intima-media thickness measurement. B-mode ultrasound is most commonly used to measure CIMT. The CIMT is defined as the distance from the lumen-intima interface to the media-adventitia interface. Strict ultrasound protocols are necessary to ensure reproducibility.',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC2691892/',
    'review',
    'Cardiologia',
    '["carotid intima-media thickness", "ultrasound", "atherosclerosis", "clinical applications", "cardiovascular risk"]'::jsonb
) ON CONFLICT (doi) DO NOTHING
RETURNING id;

-- Artigo 3: Cardiovascular Risk Assessment with CIMT and CAC 2024
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    abstract,
    original_link,
    article_type,
    specialty,
    keywords
) VALUES (
    'Assessing Cardiovascular Risk with Coronary Artery Calcium and Carotid Intima-Media Thickness in Patients with Negative Stress Echocardiography',
    'Multiple Authors',
    'PMC - PubMed Central',
    '2024-09-01',
    'en',
    NULL,
    'PMC11429111',
    'Findings suggest that increased cIMT and CAC scores may serve as useful markers for predicting CV outcomes even in symptomatic patients with negative stress echocardiography results. The study demonstrated that increased cIMT and CAC scores may be a risk factor for cardiovascular events.',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC11429111/',
    'research_article',
    'Cardiologia',
    '["carotid intima-media thickness", "coronary artery calcium", "cardiovascular risk", "stress echocardiography", "risk prediction"]'::jsonb
) ON CONFLICT (doi) DO NOTHING
RETURNING id;

-- Artigo 4: Unravelling the role of carotid atherosclerosis
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    abstract,
    original_link,
    article_type,
    specialty,
    keywords
) VALUES (
    'Unravelling the role of carotid atherosclerosis in predicting cardiovascular disease risk: A review',
    'Multiple Authors',
    'PMC - PubMed Central',
    '2024-12-01',
    'en',
    NULL,
    'PMC11663449',
    'A 2021 meta-analysis provided strong evidence of a direct relationship between carotid intima-media thickness (CIMT) and the severity of coronary artery disease (p < 0.001). Research demonstrated a strong association between a maximum CIMT above 1.54 mm and the presence of severe coronary artery disease.',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC11663449/',
    'review',
    'Cardiologia',
    '["carotid atherosclerosis", "cardiovascular disease", "risk prediction", "coronary artery disease", "CIMT"]'::jsonb
) ON CONFLICT (doi) DO NOTHING
RETURNING id;

-- Verificar os IDs inseridos
SELECT id, title, pm_id
FROM articles
WHERE pm_id IN ('PMC12162042', 'PMC2691892', 'PMC11429111', 'PMC11663449')
ORDER BY publish_date DESC;
