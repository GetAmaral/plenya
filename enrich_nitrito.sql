-- Enrichment for Score Item: Nitrito (ID: 1aa25d4b-a972-40db-a288-9cbe506de99e)
-- Generated: 2026-01-29
-- Clinical content based on peer-reviewed research (2022-2025)

BEGIN;

-- Update clinical content for Nitrito
UPDATE score_items
SET
  clinical_relevance = 'O teste de nitrito urinário é um marcador indireto essencial para detecção rápida de infecções do trato urinário (ITU) causadas por bactérias Gram-negativas produtoras de nitrato redutase, principalmente Escherichia coli, Klebsiella e Proteus. Estas bactérias convertem nitrato dietético presente na urina em nitrito, detectável pela tira reagente. Estudos recentes demonstram que o nitrito apresenta especificidade elevada (94-100%) mas sensibilidade variável (23-85%), sendo particularmente útil quando positivo. Um resultado positivo de nitrito tem alta probabilidade (99%) de confirmar ITU, especialmente em combinação com leucócito esterase, alcançando sensibilidade de 87.6% e valor preditivo negativo de 94.6%. A presença de nitrito indica bacteriúria significativa (≥10⁵ UFC/mL), geralmente refletindo infecção estabelecida. Porém, resultados falso-negativos são frequentes (3-5% das culturas positivas) devido à falta de nitrato dietético, diluição da urina, tempo de permanência vesical insuficiente (<4 horas), interferência de ácido ascórbico (vitamina C), pH urinário <6.0, ou infecção por patógenos que não produzem nitrato redutase (Enterococcus, Pseudomonas, Acinetobacter, Staphylococcus saprophyticus). Pesquisas de 2025 demonstram aplicação inovadora do nitrito como biomarcador farmacodinâmico, permitindo monitoramento de resposta antimicrobiana e diferenciação entre cepas resistentes e sensíveis de E. coli. O teste de nitrito permanece ferramenta valiosa de triagem quando interpretado dentro do contexto clínico apropriado, mas não substitui a urocultura definitiva.',
  patient_explanation = 'O nitrito na urina é um sinal que indica presença de bactérias específicas que causam infecção urinária, principalmente a bactéria E. coli. Normalmente não existe nitrito na urina saudável. Quando certas bactérias estão presentes na bexiga por algumas horas, elas transformam uma substância natural da urina (nitrato) em nitrito, que é detectado pelo exame de fita. Um resultado positivo sugere fortemente que você tem infecção urinária, especialmente se acompanhado de sintomas como dor ao urinar, urgência ou frequência aumentada. Porém, um resultado negativo não descarta completamente a possibilidade de infecção, pois cerca de 15-25% das infecções não produzem nitrito detectável. Isso pode acontecer se você urinou há pouco tempo (as bactérias precisam de pelo menos 4 horas na bexiga para produzir nitrito), se está tomando vitamina C em altas doses, se a infecção é causada por bactérias que não produzem nitrito, ou se você bebeu muita água e diluiu sua urina. Por isso, mesmo com nitrito negativo, se você tem sintomas de infecção urinária, seu médico pode solicitar uma urocultura completa. O exame de nitrito é rápido e barato, mas deve sempre ser interpretado junto com seus sintomas e outros resultados de laboratório.',
  conduct = 'NITRITO POSITIVO: Indica alta probabilidade de ITU bacteriana por Gram-negativos produtores de nitrato redutase. Conduta: (1) Avaliar sintomas clínicos (disúria, polaciúria, urgência, dor suprapúbica, febre); (2) Solicitar urocultura com antibiograma antes de iniciar antibioticoterapia, especialmente em casos recorrentes ou complicados; (3) Considerar tratamento empírico imediato em ITU não complicada sintomática conforme diretrizes locais (nitrofurantoína, fosfomicina ou sulfametoxazol-trimetoprima), ajustando após resultado da cultura; (4) Investigar presença concomitante de leucócito esterase, hematúria e proteinúria para estratificação de risco; (5) Em gestantes, pacientes diabéticos, imunossuprimidos ou com sintomas sistêmicos, considerar ITU complicada e intensificar monitoramento. NITRITO NEGATIVO: Não exclui ITU. Conduta: (1) Avaliar leucócito esterase e sedimento urinário (piúria); (2) Se sintomas sugestivos de ITU estão presentes, solicitar urocultura mesmo com nitrito negativo; (3) Considerar causas de falso-negativo: tempo de permanência vesical <4 horas (coletar primeira urina da manhã), diluição excessiva, uso de vitamina C, pH urinário baixo, ou patógenos não produtores de nitrato redutase (Enterococcus, Pseudomonas, Staphylococcus saprophyticus); (4) Em crianças <2 anos, sensibilidade do nitrito é apenas 23%, portanto sempre confirmar com urocultura; (5) Em idosos ou pacientes com cateter vesical, interpretar com cautela devido a alta taxa de bacteriúria assintomática. MONITORAMENTO: Estudos recentes (2025) sugerem que medições seriadas de nitrito podem auxiliar no monitoramento farmacodinâmico da resposta antimicrobiana, diferenciando cepas resistentes de sensíveis. IMPORTANTE: A combinação nitrito + leucócito esterase + coloração de Gram aumenta sensibilidade diagnóstica para 87.6%, reduzindo uroculturas desnecessárias e uso inadequado de antibióticos. Sempre correlacionar achados laboratoriais com apresentação clínica do paciente.',
  last_review = CURRENT_DATE
WHERE id = '1aa25d4b-a972-40db-a288-9cbe506de99e';

-- Insert scientific articles
INSERT INTO articles (id, pm_id, title, authors, journal, publish_date, article_type)
VALUES
  (
    gen_random_uuid(),
    '36762527',
    'Nitrite-negative results in urinary tract infection by Enterobacterales: does the nitrite dipstick test have low sensitivity?',
    'Araújo-Filho CE, Galvão VS, do Espírito Santo RF, et al.',
    'Journal of Medical Microbiology',
    '2023-01-01'::date,
    'research_article'
  ),
  (
    gen_random_uuid(),
    '35865430',
    'The Role of Urinary Nitrite in Predicting Bacterial Resistance in Urine Culture Analysis Among Patients With Uncomplicated Urinary Tract Infection',
    'Papava V, Didbaridze T, Zaalishvili Z, Gogokhia N, Maziashvili G',
    'Cureus',
    '2022-01-01'::date,
    'research_article'
  ),
  (
    gen_random_uuid(),
    '40487910',
    'Performance of Gram Stain, Leukocyte Esterase, and Nitrite in Predicting the Presence of Urinary Tract Infections: A Diagnostic Accuracy Study',
    'Solorzano C, Rubio MC, Licht-Ardila M, Castillo C, Valencia Silva JC, Caro MA, Manrique-Hernández EF, Hurtado-Ortiz A, García LT',
    'Archives of Academic Emergency Medicine',
    '2025-01-01'::date,
    'research_article'
  ),
  (
    gen_random_uuid(),
    '40596670',
    'Feasibility of serial measurement of nitrite for pharmacodynamic monitoring and precision prescribing in urinary tract infections',
    'Stadler EV, Holmes A, O''Hare D, Sutton M, Brown C, Rawson TM',
    'Communications Medicine',
    '2025-01-01'::date,
    'research_article'
  );

-- Link articles to score item via junction table
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '1aa25d4b-a972-40db-a288-9cbe506de99e'
FROM articles a
WHERE a.pm_id IN ('36762527', '35865430', '40487910', '40596670');

COMMIT;

-- Verification queries
SELECT
  LENGTH(clinical_relevance) as clinical_relevance_chars,
  LENGTH(patient_explanation) as patient_explanation_chars,
  LENGTH(conduct) as conduct_chars,
  last_review
FROM score_items
WHERE id = '1aa25d4b-a972-40db-a288-9cbe506de99e';

SELECT
  a.pm_id,
  a.title,
  a.journal,
  EXTRACT(YEAR FROM a.publish_date) as year
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '1aa25d4b-a972-40db-a288-9cbe506de99e'
ORDER BY a.publish_date DESC;
