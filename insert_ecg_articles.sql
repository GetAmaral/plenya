-- Inserir artigos científicos sobre Sokolow-Lyon e HVE

-- Artigo 1: Meta-análise sobre valor preditivo
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    abstract,
    article_type,
    specialty,
    keywords,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Predictive value of electrocardiographic left ventricular hypertrophy in the general population: A meta-analysis',
    'Zhigang You, Ting He, Ying Ding, Lu Yang, Xinghua Jiang, Lin Huang',
    'Journal of Electrocardiology',
    '2020-09-01',
    'en',
    '10.1016/j.jelectrocard.2020.07.001',
    '32745730',
    'This meta-analysis examined electrocardiographic left ventricular hypertrophy (LVH) as a predictor of adverse cardiac outcomes in 58,400 participants across 10 studies. The Sokolow-Lyon voltage, Cornell voltage, and Cornell product criteria showed comparable ability in predicting major adverse cardiovascular events (MACE), with risk ratios ranging from 1.56 to 1.70. The pooled risk ratio of MACE was 1.62 (95% CI 1.40-1.89) for Sokolow-Lyon voltage criteria. The pooled risk ratio of all-cause mortality was 1.47 (95% CI 1.10-1.97), and cardiovascular mortality was 1.38 (95% CI 1.19-1.60) for Sokolow-Lyon criteria. Cornell voltage demonstrated stronger predictive value for cardiovascular and all-cause mortality outcomes.',
    'meta_analysis',
    'Cardiologia',
    '["left ventricular hypertrophy", "electrocardiography", "Sokolow-Lyon", "cardiovascular risk", "meta-analysis", "mortality prediction"]'::jsonb,
    NOW(),
    NOW()
) RETURNING id;

-- Artigo 2: Diferenças por sexo
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    abstract,
    article_type,
    specialty,
    keywords,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Left ventricular hypertrophy determined by Sokolow-Lyon criteria: a different predictor in women than in men?',
    'R L Antikainen, T Grodzicki, A J Palmer, D G Beevers, J Webster, C J Bulpitt',
    'Journal of Human Hypertension',
    '2006-06-01',
    'en',
    '10.1038/sj.jhh.1002006',
    '16708082',
    'This prospective study examined 3,338 women and 3,330 men with hypertension over 11.2 years to assess whether ECG left ventricular hypertrophy (LVH) by Sokolow-Lyon voltage criteria predicted cardiovascular outcomes differently by gender. Increasing voltage independently predicted CVD mortality in both men and women. The risk of stroke, coronary heart disease (CHD) and cardiovascular disease (CVD) mortality increased significantly for each quantitative 0.1 mV increase in baseline ECG voltage, in women within the range of 1.6-3.9% and in men 1.4-3.0%. Women tended to have a high risk of stroke mortality owing to LVH, while men demonstrated stronger associations between voltage and coronary heart disease mortality.',
    'research_article',
    'Cardiologia',
    '["left ventricular hypertrophy", "Sokolow-Lyon", "gender differences", "cardiovascular mortality", "hypertension", "stroke risk"]'::jsonb,
    NOW(),
    NOW()
) RETURNING id;

-- Buscar IDs dos artigos recém-criados
SELECT id, title, pm_id FROM articles WHERE pm_id IN ('32745730', '16708082');
