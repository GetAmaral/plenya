-- Enrichment for NT-proBNP (50-75 anos) - Score Item ID: 7998e69e-a0e0-488b-bcc3-9da32e59adfb
-- Generated: 2026-01-29
-- Focus: Age-specific cutoffs, heart failure diagnosis, ESC 2023/ACC 2022 guidelines

BEGIN;

-- Step 1: Insert scientific articles
INSERT INTO articles (id, title, authors, journal, publish_date, pm_id, doi, article_type, created_at, updated_at)
VALUES
  -- Article 1: ESC HFA 2023 Consensus - Bayes-Genis et al.
  (
    gen_random_uuid(),
    'Practical algorithms for early diagnosis of heart failure and heart stress using NT-proBNP: A clinical consensus statement from the Heart Failure Association of the ESC',
    'Antoni Bayes-Genis, Kieran F Docherty, Mark C Petrie, James L Januzzi, Christian Mueller, Lisa Anderson, Biykem Bozkurt, Javed Butler, Ovidiu Chioncel, John GF Cleland, Ruxandra Christodorescu, Stefano Del Prato, Finn Gustafsson, Carolyn SP Lam, Brenda Moura, Rodica Pop-Busui, Petar Seferovic, Maurizio Volterrani, Muthiah Vaduganathan, Marco Metra, Giuseppe Rosano',
    'European Journal of Heart Failure',
    '2023-11-01'::date,
    '37712339',
    '10.1002/ejhf.3036',
    'review',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  ),

  -- Article 2: Age-adjusted thresholds community study - Taylor et al. 2025
  (
    gen_random_uuid(),
    'Age-adjusted natriuretic peptide thresholds for a diagnosis of heart failure in the community: Diagnostic accuracy study',
    'Clare J Taylor, Kathryn S Taylor, Nicholas R Jones, Jose M Ordóñez-Mena, Antoni Bayes-Genis, FD Richard Hobbs',
    'ESC Heart Failure',
    '2025-01-01'::date,
    '40717271',
    '10.1002/ehf2.15383',
    'research_article',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  ),

  -- Article 3: Optimal thresholds elderly - Berthelot et al. 2024
  (
    gen_random_uuid(),
    'Setting the optimal threshold of NT-proBNP and BNP for the diagnosis of heart failure in patients over 75 years',
    'Emmanuelle Berthelot, Minh Tam Bailly, Xenia Cerchez Lehova, Manel El Blidi Rahmani, Rahil Bounab, Nathan Mewton, John E Dobbs, Remy Mas, Marie Frank, Nicolas Lellouche, Marion Paclot, Patrick Jourdain',
    'ESC Heart Failure',
    '2024-09-01'::date,
    '38923835',
    '10.1002/ehf2.14894',
    'research_article',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  ),

  -- Article 4: NT-proBNP in atrial fibrillation - Budolfsen et al. 2023
  (
    gen_random_uuid(),
    'NT-proBNP cut-off value for ruling out heart failure in atrial fibrillation patients - A prospective clinical study',
    'Cecilie Budolfsen, Anders Sjørslev Schmidt, Kasper Glerup Lauridsen, Camilla Bang Hoeks, Farhad Waziri, Christian Bo Poulsen, Dung Nguyen Riis, Hans Rickers, Bo Løfgren',
    'American Journal of Emergency Medicine',
    '2023-10-01'::date,
    '37320999',
    '10.1016/j.ajem.2023.05.041',
    'clinical_trial',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  )
ON CONFLICT (doi) DO UPDATE SET
  pm_id = EXCLUDED.pm_id,
  title = EXCLUDED.title,
  authors = EXCLUDED.authors,
  updated_at = CURRENT_TIMESTAMP;

-- Step 2: Link articles to score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '7998e69e-a0e0-488b-bcc3-9da32e59adfb'::uuid
FROM articles a
WHERE a.pm_id IN ('37712339', '40717271', '38923835', '37320999')
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Step 3: Update score item with comprehensive clinical content
UPDATE score_items
SET
  clinical_relevance = 'O NT-proBNP (peptídeo natriurético cerebral N-terminal) é o biomarcador mais validado para diagnóstico e estratificação de insuficiência cardíaca (IC) em adultos de 50-75 anos. Esta faixa etária requer pontos de corte ajustados pela idade devido ao aumento fisiológico dos peptídeos natriuréticos com o envelhecimento. As diretrizes ESC 2023 estabelecem valores de referência específicos: ≥250 pg/mL para rastreio ambulatorial (sensibilidade 88,5%, especificidade 67,8%) e ≥900 pg/mL para IC aguda no departamento de emergência. O NT-proBNP auxilia na diferenciação entre dispneia de origem cardíaca e não-cardíaca, sendo particularmente útil em pacientes com apresentação atípica. Valores <300 pg/mL têm valor preditivo negativo de 98%, excluindo IC com alta confiança. A zona cinzenta (300-900 pg/mL) requer investigação complementar com ecocardiograma. Condições que elevam falsamente o NT-proBNP incluem: fibrilação atrial (presente em 30% dos casos de IC de novo), disfunção renal (TFG <60 mL/min), idade avançada e sepse. Obesidade reduz os níveis (ajuste recomendado para IMC >30). O monitoramento seriado do NT-proBNP auxilia na titulação de terapia e prediz desfechos, com reduções >30% associadas à melhora prognóstica. Em pacientes com IC estabelecida, valores persistentemente elevados (>1000 pg/mL) indicam maior risco de hospitalização e mortalidade cardiovascular.',

  patient_explanation = 'O NT-proBNP é uma substância produzida pelo coração quando ele está sobrecarregado ou funcionando com dificuldade. Entre 50 e 75 anos, os valores normais são naturalmente mais altos do que em pessoas mais jovens, por isso utilizamos referências ajustadas para sua idade. Este exame é fundamental para investigar falta de ar (dispneia) e inchaço nas pernas, ajudando seu médico a diferenciar se esses sintomas são causados por problemas cardíacos ou respiratórios. Valores abaixo de 250 pg/mL em consulta de rotina geralmente indicam que problemas cardíacos graves são improváveis. Valores entre 250-900 pg/mL requerem investigação adicional com ultrassom do coração (ecocardiograma). Valores acima de 900 pg/mL, especialmente se você tiver sintomas agudos, sugerem sobrecarga cardíaca que precisa ser tratada. É importante saber que algumas condições podem alterar o resultado: fibrilação atrial (arritmia comum) tende a elevar os valores; problemas nos rins também aumentam o NT-proBNP; obesidade pode reduzir artificialmente os níveis. Se você já tem diagnóstico de insuficiência cardíaca, seu médico pode solicitar este exame periodicamente para verificar se o tratamento está funcionando - quedas nos valores geralmente indicam melhora. Mantenha seus resultados anteriores para comparação e informe seu médico sobre uso de medicamentos diuréticos, que podem influenciar os valores.',

  conduct = 'Protocolo de interpretação do NT-proBNP para faixa etária 50-75 anos conforme diretrizes ESC 2023/ACC 2022: CONTEXTO AMBULATORIAL - NT-proBNP <250 pg/mL: IC improvável (VPN 94,6%), considerar causas não-cardíacas de dispneia. NT-proBNP 250-500 pg/mL: zona cinzenta, solicitar ecocardiograma transtorácico (ECO) + ECG, avaliar peptídeo natriurético atrial (BNP) se disponível. NT-proBNP >500 pg/mL: encaminhar para cardiologista para investigação de IC, ECO prioritário em <2 semanas. CONTEXTO DE EMERGÊNCIA - NT-proBNP <300 pg/mL: IC aguda improvável (VPN 98%), investigar causas respiratórias/renais. NT-proBNP 300-900 pg/mL: zona cinzenta, realizar ECO à beira-leito, considerar troponina e radiografia de tórax. NT-proBNP >900 pg/mL: IC aguda provável, iniciar terapia dirigida (diuréticos IV, IECA/BRA, betabloqueador após estabilização), ECO em <24h. AJUSTES ESPECÍFICOS - Fibrilação atrial: aumentar ponto de corte em 30-50% ou usar ≥1400 pg/mL para rule-out em emergência. Obesidade (IMC >30): reduzir valores de referência em 50% ou usar BNP. Disfunção renal (TFG 30-60): interpretar com cautela, valores até 50% mais altos podem ser basais. CONDUTA PÓS-DIAGNÓSTICO - IC estabelecida: solicitar NT-proBNP a cada 3-6 meses ou quando piora clínica. Alvo terapêutico: redução >30% em 3 meses indica resposta adequada. Valores persistentemente >1000 pg/mL: otimizar terapia farmacológica (quadrupla terapia GDMT), considerar dispositivos (CDI/ressincronização) se FE <35%. CORRELAÇÃO CLÍNICA OBRIGATÓRIA - Sempre interpretar NT-proBNP junto com sinais/sintomas, ECG, radiografia tórax e função renal. Elevar threshold de decisão na presença de: anemia (Hb <10 g/dL), sepse, embolia pulmonar, cirrose hepática.',

  updated_at = CURRENT_TIMESTAMP
WHERE id = '7998e69e-a0e0-488b-bcc3-9da32e59adfb';

COMMIT;

-- Verification queries
SELECT
  si.code,
  si.name,
  LENGTH(si.clinical_relevance) as clinical_rel_chars,
  LENGTH(si.patient_explanation) as patient_exp_chars,
  LENGTH(si.conduct) as conduct_chars,
  COUNT(DISTINCT asi.article_id) as linked_articles
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '7998e69e-a0e0-488b-bcc3-9da32e59adfb'
GROUP BY si.id, si.code, si.name, si.clinical_relevance, si.patient_explanation, si.conduct;

-- Show linked articles
SELECT
  a.pm_id,
  a.title,
  a.journal,
  a.publish_date,
  a.article_type
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '7998e69e-a0e0-488b-bcc3-9da32e59adfb'
ORDER BY a.publish_date DESC;
