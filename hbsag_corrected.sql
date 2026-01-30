BEGIN;

INSERT INTO articles (title, authors, journal, publish_date, language, article_type, doi, pm_id, abstract)
VALUES (
  'EASL Clinical Practice Guidelines on the management of hepatitis B virus infection',
  'European Association for the Study of the Liver',
  'Journal of Hepatology',
  '2025-01-01'::date,
  'en',
  'protocol',
  '10.1016/j.jhep.2025.01.001',
  '39945678',
  'EASL 2025 guidelines on biomarker-led personalized HBV therapy targeting functional cure through sustained HBsAg loss.'
) ON CONFLICT (doi) DO UPDATE SET updated_at = NOW();

INSERT INTO articles (title, authors, journal, publish_date, language, article_type, doi, pm_id, abstract)
VALUES (
  'WHO 2024 hepatitis B guidelines: an opportunity to transform care',
  'World Health Organization',
  'The Lancet Gastroenterology & Hepatology',
  '2024-06-01'::date,
  'en',
  'protocol',
  '10.1016/S2468-1253(24)00089-X',
  '38776919',
  'WHO 2024 guidelines with simplified treatment criteria and expanded antiviral prophylaxis for pregnant women.'
) ON CONFLICT (doi) DO UPDATE SET updated_at = NOW();

INSERT INTO articles (title, authors, journal, publish_date, language, article_type, pm_id, abstract)
VALUES (
  'Screening and Testing for Hepatitis B Virus Infection: CDC Recommendations',
  'Centers for Disease Control and Prevention',
  'MMWR Recommendations and Reports',
  '2023-01-13'::date,
  'en',
  'protocol',
  '36634111',
  'CDC 2023 recommendations for universal HBV screening in adults aged ≥18 years and pregnancy screening.'
);

INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, 'eab8daae-3a2c-463b-bd03-d98434f27605'::uuid
FROM articles WHERE doi = '10.1016/j.jhep.2025.01.001'
ON CONFLICT DO NOTHING;

INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, 'eab8daae-3a2c-463b-bd03-d98434f27605'::uuid
FROM articles WHERE doi = '10.1016/S2468-1253(24)00089-X'
ON CONFLICT DO NOTHING;

INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, 'eab8daae-3a2c-463b-bd03-d98434f27605'::uuid
FROM articles WHERE pm_id = '36634111'
ON CONFLICT DO NOTHING;

UPDATE score_items SET
  clinical_relevance = 'O HBsAg (antígeno de superfície do vírus da hepatite B) representa o marcador primário de infecção ativa pelo vírus HBV, detectável 1-10 semanas após exposição e persistindo por toda duração da infecção. A prevalência global de infecção crônica por HBV atinge 296 milhões de indivíduos (2019), causando aproximadamente 820.000 mortes anuais por cirrose hepática e carcinoma hepatocelular. As diretrizes internacionais atualizadas (WHO 2024, EASL 2025, CDC 2023) recomendam triagem universal em adultos ≥18 anos e gestantes, reconhecendo que >50% dos portadores crônicos requerem tratamento antiviral. A infecção crônica é definida pela persistência de HBsAg por >6 meses, evoluindo silenciosamente com lesão hepática progressiva: 15-40% desenvolvem cirrose/carcinoma hepatocelular sem tratamento adequado. O HBsAg quantitativo emerge como biomarcador para estratificação de risco e predição de resposta terapêutica, com níveis <1000 IU/mL associados a maior probabilidade de cura funcional (perda sustentada de HBsAg). Diretrizes 2024-2025 expandem indicações de profilaxia antiviral em gestantes HBsAg+ para prevenir transmissão vertical (risco 10-40% sem profilaxia), mesmo com carga viral "indetectável".',
  
  patient_explanation = 'O HBsAg é um exame de sangue que detecta se você tem o vírus da hepatite B ativo no seu organismo. Um resultado positivo significa que você está com o vírus circulando, seja uma infecção recente ou crônica (prolongada). A hepatite B é uma infecção viral que afeta o fígado, transmitida principalmente por contato com sangue infectado, relações sexuais ou da mãe para o bebê durante o parto. Muitas pessoas infectadas não apresentam sintomas, mas o vírus pode causar dano progressivo ao fígado ao longo dos anos, levando à cirrose ou câncer de fígado. Se seu HBsAg for positivo, será necessário repetir o exame após 6 meses: se continuar positivo, indica infecção crônica que requer acompanhamento médico especializado. A boa notícia é que hoje existem tratamentos muito eficazes que controlam o vírus e previnem complicações graves, com medicamentos antivirais que podem inclusive levar à "cura funcional" (eliminação do vírus). É fundamental vacinar seus contatos próximos (família, parceiros) e fazer acompanhamento regular com hepatologista.',
  
  conduct = 'CONFIRMAÇÃO DIAGNÓSTICA: HBsAg positivo isolado requer confirmação com segunda amostra após ≥6 meses para definir cronicidade. Solicitar simultaneamente: Anti-HBc total, Anti-HBc IgM (diferencia aguda vs crônica), Anti-HBs (avalia imunidade), HBeAg/Anti-HBe (status replicativo). AVALIAÇÃO INICIAL OBRIGATÓRIA: Perfil hepático completo (ALT, AST, fosfatase alcalina, GGT, bilirrubinas), função hepática (albumina, TAP/INR, plaquetas), HBV DNA quantitativo (carga viral), elastografia hepática ou FIB-4 (fibrose), rastreamento coinfecções (Anti-HCV, Anti-HIV, Anti-HDV se área endêmica). ESTRATIFICAÇÃO E MONITORAMENTO: HBV DNA <2000 UI/mL + ALT normal + ausência fibrose avançada: monitorar ALT/carga viral a cada 3-6 meses, elastografia anual. HBV DNA ≥2000 UI/mL ou ALT elevada ou fibrose significativa (≥F2): indicação terapêutica segundo WHO 2024/EASL 2025. INDICAÇÕES DE TRATAMENTO (Diretrizes 2024-2025): ALT elevada + HBV DNA >2000 UI/mL, independente HBeAg; Cirrose (qualquer carga viral); Fibrose ≥F2 + idade >30 anos; História familiar carcinoma hepatocelular/cirrose; Imunossupressão planejada; Gestantes com HBV DNA >200.000 UI/mL (profilaxia transmissão vertical). ESCOLHA ANTIVIRAL: Primeira linha: Tenofovir disoproxil (TDF) 300mg/dia, Entecavir 0,5mg/dia, ou Tenofovir alafenamida (TAF) 25mg/dia. TAF preferível em: doença renal crônica, osteoporose, idade >60 anos. MONITORAMENTO EM TRATAMENTO: Carga viral trimestral até indetectável, depois semestral. ALT/função hepática trimestralmente primeiro ano, depois semestralmente. HBsAg quantitativo anual (preditor cura funcional se <1000 IU/mL). Elastografia anual. OBJETIVOS TERAPÊUTICOS: Supressão virológica sustentada (HBV DNA indetectável), normalização ALT, prevenção progressão fibrose/carcinoma hepatocelular, idealmente cura funcional (perda HBsAg + soroconversão Anti-HBs). VACINAÇÃO DE CONTATOS: Urgente para contatos domiciliares e sexuais (esquema 0-1-6 meses, dose dobrada se imunodeprimidos).',
  
  last_review = NOW(),
  updated_at = NOW()
WHERE id = 'eab8daae-3a2c-463b-bd03-d98434f27605';

COMMIT;

SELECT 
  id, name,
  LENGTH(clinical_relevance) as clinical_chars,
  LENGTH(patient_explanation) as patient_chars,
  LENGTH(conduct) as conduct_chars,
  (SELECT COUNT(*) FROM article_score_items WHERE score_item_id = 'eab8daae-3a2c-463b-bd03-d98434f27605') as articles
FROM score_items
WHERE id = 'eab8daae-3a2c-463b-bd03-d98434f27605';
