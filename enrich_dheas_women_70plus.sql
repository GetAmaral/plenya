-- Enrichment for DHEA-S - Mulheres (70+ anos)
-- Score Item ID: e622f011-e92d-459e-9f4a-af435f26ea74

BEGIN;

-- Insert scientific articles (skip duplicates by DOI)
INSERT INTO articles (id, title, authors, journal, publish_date, doi, article_type, original_link, created_at, updated_at)
VALUES
  (
    'b8f9c012-1a5d-4e8f-9c3a-1f8d9e2a3b4c',
    'Impact of DHEA supplementation on testosterone and estradiol levels in postmenopausal women: a meta-analysis of randomized controlled trials assessing dose and duration effects',
    'He SY, Lu K, Zhang L, Cao H, Tang X, Zhang X',
    'Diabetology & Metabolic Syndrome',
    '2025-07-04',
    '10.1186/s13098-025-01770-0',
    'meta_analysis',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC12231631/',
    NOW(),
    NOW()
  ),
  (
    'c9e8d123-2b6e-4f9a-8d4b-2e9f0a3c4d5e',
    'DHEAS Levels and Mortality in Disabled Older Women: The Women''s Health and Aging Study I',
    'Cappola AR, Xue QL, Walston JD, Leng SX, Ferrucci L, Guralnik J, Fried LP',
    'Journal of Gerontology A: Biological Sciences and Medical Sciences',
    '2006-09-01',
    '10.1093/gerona/61.9.957',
    'research_article',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC2645634/',
    NOW(),
    NOW()
  ),
  (
    'e1a0f345-4d8e-4b1c-0f6d-4a1b2c5e6f7a',
    'Dehydroepiandrosterone Combined with Exercise Improves Muscle Strength and Physical Function in Frail Older Women',
    'Kenny AM, Boxer RS, Kleppinger A, Brindisi J, Feinn R, Burleson JA',
    'Journal of the American Geriatrics Society',
    '2010-08-01',
    '10.1111/j.1532-5415.2010.03019.x',
    'research_article',
    'https://pubmed.ncbi.nlm.nih.gov/20863330/',
    NOW(),
    NOW()
  )
ON CONFLICT (doi) DO NOTHING;

-- Link articles to the score item (including the one that already exists)
-- First, get the existing article ID for the duplicate
INSERT INTO article_score_items (score_item_id, article_id)
VALUES
  ('e622f011-e92d-459e-9f4a-af435f26ea74', 'b8f9c012-1a5d-4e8f-9c3a-1f8d9e2a3b4c'),
  ('e622f011-e92d-459e-9f4a-af435f26ea74', 'c9e8d123-2b6e-4f9a-8d4b-2e9f0a3c4d5e'),
  ('e622f011-e92d-459e-9f4a-af435f26ea74', '3b4bc8da-ff7d-419f-a867-a21872cbc4a4'), -- existing article with DOI 10.1089/rej.2013.1425
  ('e622f011-e92d-459e-9f4a-af435f26ea74', 'e1a0f345-4d8e-4b1c-0f6d-4a1b2c5e6f7a')
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- Update score item with comprehensive clinical content
UPDATE score_items
SET
  clinical_relevance = 'O DHEA-S (sulfato de deidroepiandrosterona) em mulheres acima de 70 anos apresenta padrão único de declínio e associações clínicas distintas das observadas em homens ou mulheres mais jovens. Meta-análises recentes demonstram que a relação entre DHEA-S e mortalidade em mulheres idosas é complexa e não linear: estudos prospectivos identificaram uma curva em "U", onde tanto níveis muito baixos quanto muito elevados associam-se a aumento de mortalidade em 5 anos (RR>2.0), com nível ótimo em torno de 41 μg/dL. Em mulheres com mais de 70 anos, os níveis de DHEA-S frequentemente representam apenas 10-20% dos valores observados em mulheres jovens, refletindo o envelhecimento adrenal fisiológico. A relevância clínica específica nesta faixa etária relaciona-se à fragilidade, função física e qualidade de vida: estudos controlados em mulheres frágeis (idade média 76 anos) demonstraram que suplementação de DHEA 50mg/dia combinada com exercício de baixa intensidade melhorou significativamente força de membros inferiores e função física após 6 meses, sem efeitos adversos graves. Diferentemente dos homens, a associação entre DHEA-S baixo e mortalidade cardiovascular em mulheres idosas é menos consistente, enquanto níveis elevados podem associar-se a maior risco de mortalidade por câncer. A monitorização de DHEA-S em mulheres com 70+ anos é particularmente relevante em contextos de fragilidade, sarcopenia, declínio funcional e como marcador potencial de envelhecimento adrenal, sempre considerando a individualização terapêutica baseada no perfil clínico completo.',

  patient_explanation = 'O DHEA-S é um hormônio produzido pelas glândulas suprarrenais que diminui naturalmente com o envelhecimento. Em mulheres com mais de 70 anos, os níveis costumam estar entre 10% e 20% dos valores encontrados em mulheres jovens, o que é considerado parte normal do processo de envelhecimento. O interessante é que, nesta idade, a relação entre DHEA-S e saúde não é simples: tanto níveis muito baixos quanto muito altos podem não ser ideais. Estudos científicos mostram que mulheres idosas com níveis intermediários de DHEA-S tendem a viver mais tempo do que aquelas com níveis extremos (muito baixo ou muito alto). O DHEA-S está relacionado com força muscular, capacidade física e qualidade de vida em idosas. Pesquisas com mulheres frágeis acima de 70 anos mostraram que a combinação de suplementação de DHEA com exercícios leves (como ioga ou aeróbica na cadeira) pode melhorar a força das pernas e a capacidade de realizar atividades diárias. No entanto, a reposição de DHEA não é indicada para todas as mulheres nesta faixa etária: deve ser avaliada individualmente, considerando sintomas, outros hormônios e condições de saúde. Níveis de DHEA-S devem ser interpretados junto com o quadro clínico completo, histórico de câncer, saúde cardiovascular e objetivos terapêuticos específicos. A decisão sobre suplementação deve sempre ser feita com orientação médica especializada.',

  conduct = 'Para mulheres com 70 anos ou mais, a abordagem diagnóstica e terapêutica do DHEA-S deve ser altamente individualizada e contextualizada. INTERPRETAÇÃO LABORATORIAL: Valores de referência ajustados para idade são essenciais, considerando que níveis entre 15-70 μg/dL podem ser fisiológicos nesta faixa etária. Estudos prospectivos sugerem que o nível ótimo associado a menor mortalidade em mulheres idosas está em torno de 41 μg/dL, mas valores individuais devem ser interpretados considerando comorbidades e fragilidade. AVALIAÇÃO CLÍNICA INTEGRADA: Antes de considerar intervenção, avaliar: presença de fragilidade (Fried criteria), sarcopenia (força de preensão, massa muscular por DEXA), função física (velocidade de marcha, Short Physical Performance Battery), qualidade de vida, histórico oncológico (especialmente câncer de mama) e risco cardiovascular. EVIDÊNCIAS PARA SUPLEMENTAÇÃO: A meta-análise mais recente (2025) demonstra que doses ≥50mg/dia aumentam testosterona (WMD: 29.65 ng/dL) e estradiol (WMD: 8.65 pg/mL) em mulheres pós-menopáusicas, com efeitos mais pronunciados em tratamentos ≥26 semanas e mulheres ≥60 anos. Ensaio clínico controlado em mulheres frágeis (idade média 76 anos) mostrou que DHEA 50mg/dia + exercício de baixa intensidade melhorou força de membros inferiores sem efeitos adversos significativos em 6 meses. INDICAÇÕES POTENCIAIS: Considerar suplementação em mulheres com fragilidade documentada, sarcopenia, níveis muito baixos de DHEA-S (<15 μg/dL) associados a sintomas, e função adrenal preservada. CONTRAINDICAÇÕES: Histórico de câncer hormônio-dependente (mama, endométrio), níveis já elevados de DHEA-S (>100 μg/dL pela curva em U de mortalidade), insuficiência adrenal primária não tratada. PROTOCOLO DE SUPLEMENTAÇÃO: Quando indicado, iniciar com DHEA 25-50mg/dia (dose padrão dos estudos em idosas), sempre combinado com programa de exercício físico supervisionado. Monitorar DHEA-S, estradiol, testosterona, função hepática aos 3 e 6 meses. Reavaliar benefícios clínicos (força, função física, qualidade de vida) após 6 meses antes de manter tratamento prolongado. A suplementação indiscriminada NÃO é recomendada: a relação em U com mortalidade exige cautela.',

  last_review = NOW(),
  updated_at = NOW()
WHERE id = 'e622f011-e92d-459e-9f4a-af435f26ea74';

COMMIT;

-- Verification query
SELECT
  si.id,
  si.name,
  si.code,
  LENGTH(si.clinical_relevance) as clinical_relevance_chars,
  LENGTH(si.patient_explanation) as patient_explanation_chars,
  LENGTH(si.conduct) as conduct_chars,
  COUNT(asi.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = 'e622f011-e92d-459e-9f4a-af435f26ea74'
GROUP BY si.id, si.name, si.code, si.clinical_relevance, si.patient_explanation, si.conduct;
