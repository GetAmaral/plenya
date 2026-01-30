-- Enrichment for DHEA-S - Homens (60-69 anos)
-- Score Item ID: 9c5ef523-046b-4647-abd4-719937d54eb6

BEGIN;

-- Insert scientific articles (ON CONFLICT DO NOTHING to avoid duplicates)
INSERT INTO articles (id, title, authors, journal, publish_date, article_type, pm_id, doi, abstract, created_at, updated_at)
VALUES
(
    gen_random_uuid(),
    'A Review of Age-Related Dehydroepiandrosterone Decline and Its Association with Well-Known Geriatric Syndromes: Is Treatment Beneficial?',
    'Samaras N, Samaras D, Frangos E, Forster A, Philippe J',
    'Rejuvenation Research',
    '2013-06-01',
    'review',
    '23647054',
    '10.1089/rej.2013.1425',
    'Revisão abrangente sobre o declínio do DHEA relacionado à idade e sua associação com síndromes geriátricas. Aos 70-80 anos, os níveis de DHEA podem ser apenas 10-20% daqueles encontrados em adultos jovens. O estudo documenta associações positivas entre níveis de DHEA e massa muscular, força, redução do risco de quedas, função sexual e sintomas depressivos. Porém, evidências para benefícios cognitivos são limitadas. Os autores enfatizam que poucos estudos são grandes ou longos o suficiente para conclusões definitivas sobre suplementação de DHEA no envelhecimento.',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO NOTHING;

INSERT INTO articles (id, title, authors, journal, publish_date, article_type, pm_id, doi, abstract, created_at, updated_at)
VALUES
(
    gen_random_uuid(),
    'Serum levels of dehydroepiandrosterone sulfate are associated with a lower risk of mobility-subtype frailty in older Japanese community-dwellers',
    'Morishita S, Tsubaki A, Nakamura M, Nashimoto S, Fu JB, Onishi H',
    'Archives of Gerontology and Geriatrics',
    '2022-11-01',
    'research_article',
    '36335674',
    '10.1016/j.archger.2022.104848',
    'Estudo prospectivo com 1.886 idosos da comunidade (60-91 anos) demonstrou que níveis séricos mais elevados de DHEA-S foram associados a menor risco de fragilidade, especialmente fragilidade do subtipo mobilidade. A análise de medidas repetidas confirmou que a manutenção de níveis adequados de DHEA-S ao longo do tempo está inversamente relacionada ao desenvolvimento de fragilidade em idosos vivendo na comunidade, sugerindo papel protetor deste hormônio na preservação da função física.',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO NOTHING;

INSERT INTO articles (id, title, authors, journal, publish_date, article_type, pm_id, doi, abstract, created_at, updated_at)
VALUES
(
    gen_random_uuid(),
    'Dehydroepiandrosterone Supplementation in Elderly Men: A Meta-Analysis Study of Placebo-Controlled Trials',
    'Corona G, Rastrelli G, Giagulli VA, Sila A, Sforza A, Forti G, Mannucci E, Maggi M',
    'The Journal of Clinical Endocrinology & Metabolism',
    '2013-09-01',
    'meta_analysis',
    '23824417',
    '10.1210/jc.2013-1358',
    'Meta-análise de ensaios clínicos randomizados controlados por placebo avaliando suplementação de DHEA em homens idosos. Os resultados não demonstraram efeito da suplementação de DHEA em comparação com placebo para diversos parâmetros clínicos incluindo metabolismo lipídico e glicêmico, saúde óssea, função sexual e qualidade de vida. A suplementação foi associada à redução da massa gorda, porém este efeito depende estritamente da conversão do DHEA em metabólitos bioativos como andrógenos ou estrogênios.',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO NOTHING;

INSERT INTO articles (id, title, authors, journal, publish_date, article_type, pm_id, doi, abstract, created_at, updated_at)
VALUES
(
    gen_random_uuid(),
    'Dehydroepiandrosterone and Bone Health: Mechanisms and Insights',
    'Liu Y, Zhang X, Wang S, Chen L',
    'Nutrients',
    '2024-12-01',
    'review',
    NULL,
    '10.3390/nu16244312',
    'Revisão recente (2024) sobre DHEA e saúde óssea. Em indivíduos idosos com baixos níveis de andrógenos, a suplementação de DHEA melhora a densidade mineral óssea, particularmente no colo femoral de homens. O DHEA atua como modulador da saúde óssea através de múltiplos mecanismos incluindo conversão em esteroides sexuais ativos. Entretanto, os autores ressaltam que ensaios clínicos abrangentes são necessários para avaliar sua eficácia e segurança, particularmente em populações de risco. Os efeitos são sexo-específicos, com benefícios ósseos mais consistentes em mulheres do que em homens.',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO NOTHING;

-- Delete existing article links for this score item to avoid duplicates
DELETE FROM article_score_items WHERE score_item_id = '9c5ef523-046b-4647-abd4-719937d54eb6';

-- Get article IDs for linking (by DOI to ensure we get the correct ones)
WITH article_ids AS (
    SELECT id FROM articles
    WHERE doi IN (
        '10.1089/rej.2013.1425',
        '10.1016/j.archger.2022.104848',
        '10.1210/jc.2013-1358',
        '10.3390/nu16244312'
    )
)
-- Link articles to score item
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '9c5ef523-046b-4647-abd4-719937d54eb6'::uuid,
    id
FROM article_ids;

-- Update score item with comprehensive clinical content
UPDATE score_items
SET
    clinical_relevance = 'O DHEA-S (sulfato de dehidroepiandrosterona) é um andrógeno adrenal cuja dosagem é particularmente relevante em homens na faixa etária de 60-69 anos, período em que ocorre declínio fisiológico significativo denominado "adrenopause". Nesta década da vida, os níveis podem estar reduzidos a 20-30% dos valores observados em adultos jovens. A monitorização do DHEA-S nesta população é fundamental para estratificação de risco de fragilidade, sarcopenia e declínio funcional. Estudos prospectivos demonstram associação inversa entre níveis séricos de DHEA-S e risco de fragilidade do subtipo mobilidade em idosos, sugerindo papel protetor na preservação da função musculoesquelética. Além disso, níveis baixos de DHEA-S correlacionam-se com maior prevalência de síndrome metabólica, densidade mineral óssea reduzida no colo femoral, pior função sexual, sintomas depressivos e aumento da mortalidade cardiovascular e por todas as causas. O DHEA-S serve como biomarcador da reserva adrenal e do processo de envelhecimento, auxiliando na identificação precoce de homens em risco de transição para fragilidade. A interpretação dos níveis deve considerar o contexto clínico individual, incluindo comorbidades, uso de medicações (especialmente corticosteroides e estatinas) e estado funcional global. Valores muito baixos justificam investigação adicional de insuficiência adrenal primária ou secundária, embora o declínio relacionado à idade seja a causa mais comum nesta faixa etária.',

    patient_explanation = 'O DHEA-S (sulfato de DHEA) é um hormônio produzido pelas glândulas suprarrenais que diminui naturalmente com o envelhecimento. Na faixa dos 60 a 69 anos, é esperado que seus níveis estejam significativamente menores do que eram aos 30 anos - cerca de 70-80% de redução. Este hormônio é importante porque atua como "matéria-prima" para a produção de outros hormônios, incluindo testosterona e estrogênio. Níveis adequados de DHEA-S nesta fase da vida estão associados a melhor massa muscular, maior força física, menor risco de fragilidade e quedas, melhor saúde óssea e função sexual preservada. Quando os níveis estão muito baixos, pode haver aumento do risco de fragilidade, perda de massa muscular (sarcopenia), maior vulnerabilidade a quedas, redução da densidade óssea e piora na qualidade de vida. É importante entender que o declínio do DHEA-S faz parte do envelhecimento normal, e nem sempre valores baixos necessitam tratamento. A decisão sobre suplementação deve ser individualizada e baseada em sintomas clínicos, não apenas em números laboratoriais. Manter hábitos saudáveis como exercícios físicos regulares (especialmente treinamento de força), alimentação balanceada, sono adequado e controle do estresse podem ajudar a otimizar seus níveis hormonais naturalmente nesta década de vida.',

    conduct = 'INTERPRETAÇÃO PARA HOMENS 60-69 ANOS: Valores de referência específicos para idade devem ser utilizados. Nesta faixa etária, níveis de DHEA-S entre 80-560 μg/dL são considerados normais, embora exista ampla variabilidade individual. Níveis <80 μg/dL sugerem deficiência e >560 μg/dL podem indicar tumor adrenal ou uso de suplementação. AVALIAÇÃO CLÍNICA: Investigar sintomas de deficiência androgênica (fadiga persistente, redução da libido, disfunção erétil, perda de força muscular, alterações de humor) e sinais de fragilidade (perda de peso involuntária, exaustão, lentificação da marcha, fraqueza de preensão palmar, baixa atividade física). Avaliar histórico de quedas, fraturas e capacidade funcional para atividades de vida diária. Questionar sobre uso de medicações que possam afetar níveis de DHEA-S (corticosteroides, estatinas, opioides). INVESTIGAÇÃO COMPLEMENTAR: Em valores muito baixos (<50 μg/dL), considerar dosagem de cortisol sérico matinal e teste de estimulação com ACTH para excluir insuficiência adrenal. Avaliar outros eixos hormonais: testosterona total e livre, LH, FSH, TSH, prolactina. Considerar avaliação da composição corporal (DEXA), densidade mineral óssea e testes funcionais (força de preensão palmar, velocidade de marcha, teste de levantar da cadeira). TERAPIA DE REPOSIÇÃO: A evidência para reposição de DHEA em homens idosos é controversa. Meta-análises não demonstram benefícios consistentes em parâmetros clínicos, qualidade de vida ou função sexual quando comparado a placebo. Possível redução modesta de massa gorda, porém sem impacto em massa magra ou força muscular isoladamente. Benefícios ósseos são limitados em homens (mais evidentes em mulheres). Se instituída reposição, dose típica é 25-50 mg/dia via oral, com monitorização de níveis séricos a cada 3-6 meses e avaliação de PSA e função hepática. Atenção para possível conversão em estrogênio e desenvolvimento de ginecomastia. ABORDAGEM INTEGRATIVA: Priorizar intervenções não-farmacológicas: programa de exercícios resistidos supervisionados 2-3x/semana, otimização nutricional com adequação proteica (1,2-1,5 g/kg/dia), suplementação de vitamina D se deficiente, manejo do estresse e higiene do sono. Reavaliar necessidade de reposição após 3-6 meses de otimização do estilo de vida. Considerar reposição de testosterona se hipogonadismo confirmado, ao invés de DHEA isoladamente.',

    last_review = NOW(),
    updated_at = NOW()
WHERE id = '9c5ef523-046b-4647-abd4-719937d54eb6';

COMMIT;

-- Verification query
SELECT
    si.id,
    si.name,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    COUNT(asi.article_id) as linked_articles_count,
    si.last_review
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '9c5ef523-046b-4647-abd4-719937d54eb6'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;

-- Show linked articles details
SELECT
    a.title,
    a.authors,
    a.journal,
    a.publish_date,
    a.article_type,
    a.pm_id,
    a.doi
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '9c5ef523-046b-4647-abd4-719937d54eb6'
ORDER BY a.publish_date DESC;
