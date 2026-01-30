-- ========================================
-- ENRICHMENT: TRAb (Anticorpos Anti-Receptor de TSH)
-- Score Item ID: 3bb82be7-cda4-49bc-8bd1-bf28a1359a6f
-- Category: Exames Laboratoriais - Hormônios Tireoidianos
-- ========================================

BEGIN;

-- Step 1: Update score_items with comprehensive clinical content
UPDATE score_items
SET
  clinical_relevance = 'Os anticorpos anti-receptor de TSH (TRAb) são autoanticorpos direcionados contra o receptor do hormônio estimulante da tireoide (TSH-R) na superfície das células tireoidianas. Representam o teste de escolha para confirmação diagnóstica da doença de Graves, com sensibilidade de 96-98% e especificidade de 97-99% quando utilizados ensaios de terceira geração. Níveis elevados (>1,75 IU/L) confirmam hipertireoidismo autoimune e distinguem de outras causas de tireotoxicose. TRAb possui três subtipos funcionais: anticorpos estimulantes (TSAb), bloqueadores (TBAb) e neutros, sendo TSAb predominante na doença de Graves ativa. O monitoramento seriado durante tratamento com drogas antitireoidianas é crucial: níveis >12 IU/L ao diagnóstico associam-se a 60% de risco de recidiva em 2 anos e 84% em 4 anos. Valores >7,5 IU/L aos 12 meses ou >3,85 IU/L ao cessar terapia predizem recorrência >90%. TRAb elevado é fator de risco estabelecido para orbitopatia de Graves (TED), com níveis pré-tratamento correlacionando-se com severidade e resposta terapêutica. Cut-off de 5,035 IU/L pós-tratamento prediz má resposta à pulsoterapia com metilprednisolona (AUC 0,752). Útil em diagnóstico diferencial de hipertireoidismo, seleção terapêutica, avaliação gestacional (risco neonatal quando >3x valor de referência no terceiro trimestre), e decisão entre terapia medicamentosa versus ablativa. Diretrizes atuais recomendam TRAb na workup diagnóstica, monitoramento de atividade da doença e prognóstico de remissão.',

  patient_explanation = 'TRAb (anticorpos anti-receptor de TSH) são proteínas produzidas pelo sistema imunológico que atacam equivocadamente a glândula tireoide. Em condições normais, o hormônio TSH se liga a receptores específicos para estimular a tireoide a produzir seus hormônios. Na doença de Graves, os TRAb imitam o TSH e mantêm a tireoide constantemente estimulada, resultando em produção excessiva de hormônios (hipertireoidismo). Valores normais são abaixo de 1,75 IU/L. Níveis elevados confirmam doença de Graves e ajudam a diferenciar de outras causas de tireoide hiperativa. Se seus níveis são muito altos (acima de 12 IU/L), existe maior chance da doença voltar após o tratamento. O acompanhamento deste exame durante o uso de medicações é importante: se os níveis não diminuírem adequadamente, pode indicar que a doença não está bem controlada. TRAb elevado também aumenta risco de desenvolver "olhos saltados" (orbitopatia de Graves), uma complicação que afeta os olhos. Para gestantes com histórico de Graves, medir TRAb no terceiro trimestre é essencial, pois níveis muito altos podem afetar o bebê. Este exame ajuda seu médico a escolher o melhor tratamento: remédios, iodo radioativo ou cirurgia, e a prever se haverá recaída após parar os medicamentos.',

  conduct = 'INDICAÇÕES DE SOLICITAÇÃO: Confirmação diagnóstica de hipertireoidismo quando suspeita de doença de Graves (presença de bócio difuso, oftalmopatia, TSH suprimido com T4L/T3L elevados); diagnóstico diferencial entre doença de Graves e hipertireoidismo por outras causas (adenoma tóxico, bócio multinodular tóxico, tireoidite); monitoramento durante tratamento com drogas antitireoidianas (solicitar aos 12 meses e ao final da terapia); avaliação de risco de recidiva pré-descontinuação de metimazol/propiltiouracil; gestantes com histórico de Graves (solicitar no 1º e 3º trimestres para estratificar risco fetal/neonatal); orbitopatia de Graves ativa (correlação com severidade e resposta terapêutica). INTERPRETAÇÃO: <1,0 IU/L = negativo, improvável doença de Graves; 1,0-1,75 IU/L = limítrofe, considerar repetir em 4-6 semanas ou solicitar TSI (ensaio bioativo); >1,75 IU/L = positivo, confirma doença de Graves; >5 IU/L = alto risco de orbitopatia e recidiva; >12 IU/L ao diagnóstico = risco 60-84% de recorrência em 2-4 anos. CONDUTA PÓS-RESULTADO: TRAb positivo + hipertireoidismo bioquímico: iniciar metimazol 10-30 mg/dia ou propiltiouracil 100-300 mg 3x/dia; monitorar TRAb aos 12 meses de tratamento: se >7,5 IU/L, considerar extensão da terapia por 18-24 meses; ao cessar drogas antitireoidianas: se TRAb >3,85 IU/L, risco elevado de recaída (>90%), discutir terapia ablativa (I-131 ou tireoidectomia); gestantes com TRAb >3x valor de referência no 3º trimestre: alertar obstetra/neonatologista para monitoramento fetal/neonatal (risco de hiper ou hipotireoidismo transitório); TRAb persistentemente elevado apesar de tratamento adequado: considerar terapias alternativas (cirurgia, iodo radioativo) ou investigar má adesão. ASSOCIAÇÕES: Solicitar conjuntamente TSH, T4 livre, T3 livre, TPOAb (diferencia tireoidite autoimune), ultrassom tireoidiano com Doppler (padrão "inferno tireoidiano" sugere Graves), cintilografia com captação (se diagnóstico incerto). Em orbitopatia ativa: solicitar TSAb (thyroid-stimulating immunoglobulin) para melhor correlação prognóstica (AUC 0,90 vs 0,82 do TRAb). SEGUIMENTO: Repetir TRAb a cada 6-12 meses durante tratamento; após remissão, monitorar anualmente por 2-3 anos devido risco de recidiva tardia.',

  last_review = CURRENT_TIMESTAMP
WHERE id = '3bb82be7-cda4-49bc-8bd1-bf28a1359a6f';

-- Verify update
SELECT id, name,
       LENGTH(clinical_relevance) as relevance_chars,
       LENGTH(patient_explanation) as explanation_chars,
       LENGTH(conduct) as conduct_chars,
       last_review
FROM score_items
WHERE id = '3bb82be7-cda4-49bc-8bd1-bf28a1359a6f';

-- Step 2: Insert scientific articles (if not already exist)

-- Article 1: Xu et al. 2023 - TSI vs TRAb diagnostic performance
INSERT INTO articles (pm_id, doi, title, authors, journal, publish_date, article_type)
VALUES (
  '37161617',
  '10.1002/jcla.24890',
  'Evaluation of the diagnostic performance of thyroid-stimulating immunoglobulin and thyrotropin receptor antibodies for Graves'' disease',
  'Xu S, Shao W, Wu Q, Zhu J, Pan B, Wang B, Guo W',
  'J Clin Lab Anal',
  '2023-01-01'::date,
  'research_article'
)
ON CONFLICT (doi) DO NOTHING;

-- Article 2: Hoang et al. 2022 - Clinical management update
INSERT INTO articles (pm_id, doi, title, authors, journal, publish_date, article_type)
VALUES (
  '35662442',
  '10.1016/j.ecl.2021.12.004',
  '2022 Update on Clinical Management of Graves'' Disease and Thyroid Eye Disease',
  'Hoang TD, Stocker DJ, Chou EL, Burch HB',
  'Endocrinol Metab Clin North Am',
  '2022-01-01'::date,
  'review'
)
ON CONFLICT (doi) DO NOTHING;

-- Article 3: Lee et al. 2023 - TRAb prognostic significance in orbitopathy
INSERT INTO articles (pm_id, doi, title, authors, journal, publish_date, article_type)
VALUES (
  '37223049',
  '10.3389/fendo.2023.1153312',
  'Prognostic significance of thyroid-stimulating hormone receptor antibodies in moderate-to-severe Graves'' orbitopathy',
  'Lee JH, Jeon MJ, Park S, Lee YJ, Lim DJ, Kang MI, Cha BY, Lee KW, Son HY',
  'Front Endocrinol (Lausanne)',
  '2023-01-01'::date,
  'research_article'
)
ON CONFLICT (doi) DO NOTHING;

-- Article 4: Kalra et al. 2024 - Best practices guidelines
INSERT INTO articles (pm_id, doi, title, authors, journal, publish_date, article_type)
VALUES (
  '39707289',
  '10.1186/s12902-024-01809-9',
  'Best practices in the laboratory diagnosis, prognostication, prediction, and monitoring of Graves'' disease: role of TRAbs',
  'Kalra S, Selim S, Shrestha D, Unnikrishnan AG, Sahay R, Dharmalingam M, Deka PP, Bantwal G, Bhattacharya S, Dharmapuri S',
  'BMC Endocr Disord',
  '2024-01-01'::date,
  'review'
)
ON CONFLICT (doi) DO NOTHING;

-- Step 3: Link articles to score item (junction table)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '3bb82be7-cda4-49bc-8bd1-bf28a1359a6f'::uuid
FROM articles a
WHERE a.pm_id IN ('37161617', '35662442', '37223049', '39707289')
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Step 4: Verification queries
SELECT
  si.name,
  COUNT(DISTINCT asi.article_id) as linked_articles,
  LENGTH(si.clinical_relevance) as relevance_length,
  LENGTH(si.patient_explanation) as explanation_length,
  LENGTH(si.conduct) as conduct_length
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '3bb82be7-cda4-49bc-8bd1-bf28a1359a6f'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct;

-- List linked articles
SELECT
  a.pm_id,
  a.title,
  a.authors,
  a.journal,
  a.publish_date,
  a.article_type
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '3bb82be7-cda4-49bc-8bd1-bf28a1359a6f'
ORDER BY a.publish_date DESC;

COMMIT;

-- ========================================
-- ENRICHMENT COMPLETE
-- ========================================
-- Score Item: TRAb (Anticorpos Anti-Receptor de TSH)
-- Articles Linked: 4
-- Clinical Relevance: ~2000 chars
-- Patient Explanation: ~1500 chars
-- Conduct: ~2500 chars
-- Date: 2026-01-29
-- ========================================
