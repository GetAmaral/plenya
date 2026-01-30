-- Enrichment for Score Item: Vagina e Colo Uterino
-- Item ID: 1ff1e7f8-5b15-4f0c-b76b-806c1d6ff1fd
-- Generated: 2026-01-29

BEGIN;

-- ============================================================================
-- STEP 1: Update clinical content fields
-- ============================================================================

UPDATE score_items
SET
  clinical_relevance = 'O exame vaginal e do colo uterino é um dos pilares fundamentais da saúde preventiva da mulher, permitindo a detecção precoce de alterações pré-malignas e do câncer cervical. A avaliação sistemática inclui a inspeção visual com espéculo, coleta de citologia oncótica (exame de Papanicolaou) e teste de HPV, quando indicado. As diretrizes internacionais atualizadas em 2024-2025 recomendam o rastreamento primário com teste de HPV a cada 5 anos para mulheres de 30-65 anos, ou co-teste (HPV + citologia) no mesmo intervalo, ou ainda citologia isolada a cada 3 anos. Para mulheres de 21-29 anos, recomenda-se citologia a cada 3 anos. A infecção persistente por HPV de alto risco, especialmente os genótipos 16 e 18, é responsável por mais de 70% dos casos de câncer cervical. O exame colposcópico com biópsia dirigida está indicado em casos de citologia alterada (ASCUS HPV-positivo, LGSIL, HGSIL) ou teste de HPV positivo para genótipos de alto risco. A avaliação adequada do colo uterino requer técnica apropriada de inserção do espéculo, visualização completa da zona de transformação, aplicação de ácido acético e, quando necessário, solução de Lugol para delimitar áreas de epitélio escamoso imaturo. Recentes avanços incluem a aprovação pela FDA de testes de HPV em amostras auto-coletadas, que demonstraram acurácia comparável às amostras coletadas por profissionais de saúde, ampliando significativamente o acesso ao rastreamento, especialmente em populações subatendidas ou com barreiras ao exame ginecológico tradicional.',

  patient_explanation = 'O exame do colo do útero, também conhecido como exame preventivo ou Papanicolaou, é essencial para detectar precocemente alterações que podem levar ao câncer de colo uterino. Durante o exame, o médico utiliza um instrumento chamado espéculo para visualizar o interior da vagina e o colo do útero, e coleta células para análise laboratorial. Atualmente, além da citologia tradicional, é possível realizar o teste para HPV (papilomavírus humano), vírus que causa a maioria dos casos de câncer cervical. As recomendações atuais indicam iniciar o rastreamento aos 21 anos com citologia, e após os 30 anos, pode-se optar pelo teste de HPV a cada 5 anos, que oferece maior proteção com intervalos mais longos entre exames. Para mulheres que têm dificuldade de acesso ao exame tradicional ou que se sentem desconfortáveis com o procedimento, uma nova opção é a coleta vaginal auto-realizada para teste de HPV, aprovada em 2024, que você mesma pode fazer em casa ou no consultório. Se o resultado mostrar alterações ou presença de HPV de alto risco, pode ser necessário um exame mais detalhado chamado colposcopia, onde o médico examina o colo do útero com aumento e pode fazer biópsias de áreas suspeitas. A vacinação contra HPV e o rastreamento regular são as estratégias mais eficazes para prevenir o câncer de colo uterino, uma doença evitável quando detectada precocemente.',

  conduct = 'A conduta no rastreamento de câncer cervical segue protocolos estabelecidos pelas principais sociedades médicas (USPSTF, ACS, ASCCP) atualizados em 2024-2025: (1) Iniciar rastreamento aos 21 anos com citologia oncótica a cada 3 anos até 29 anos; (2) Entre 30-65 anos, optar preferencialmente por teste primário de HPV a cada 5 anos, ou co-teste (HPV + citologia) a cada 5 anos, ou citologia isolada a cada 3 anos; (3) Considerar auto-coleta vaginal para teste de HPV em mulheres com barreiras ao exame especular (aprovado pela FDA em 2024), com reteste em 3 anos se negativo; (4) Em caso de citologia ASCUS, realizar genotipagem de HPV - se HPV 16/18 positivo, encaminhar diretamente para colposcopia; (5) Toda citologia LGSIL ou superior em mulheres >24 anos requer colposcopia com biópsia dirigida; (6) Para lesões HGSIL ou ASCUS-H, realizar colposcopia independentemente da idade; (7) Na colposcopia, utilizar espéculo adequado ao tamanho da paciente, aplicar ácido acético a 5% por 1-2 minutos para identificar epitélio acetobranco, examinar sob magnificação de 10-15x, documentar tipo de zona de transformação (TZ1, TZ2, TZ3) e escore de Swede, realizar biópsias dirigidas de áreas suspeitas (acetobrancas densas, mosaico/pontilhado grosseiro, vasos atípicos); (8) Curetagem endocervical está indicada quando a zona de transformação não é completamente visível (TZ3) ou há suspeita de lesão endocervical; (9) Utilizar testes moleculares complementares como p16/Ki-67 dual stain ou metilação de HPV para triagem de casos HPV-positivos com citologia negativa; (10) Manejo pós-tratamento de lesões intraepiteliais (excisão/ablação): co-teste em 6 meses, depois anualmente por 25 anos ou até 65 anos de idade; (11) Critérios para cessação do rastreamento: ≥65 anos com rastreamento adequado prévio negativo (3 citologias ou 2 testes de HPV consecutivos negativos nos últimos 10 anos, sendo o mais recente nos últimos 5 anos) E sem história de CIN2+ nos últimos 25 anos; (12) Mulheres vacinadas contra HPV seguem as mesmas diretrizes de rastreamento; (13) Documentar sempre consentimento informado, aconselhamento sobre HPV e vacinação, e registrar todos os achados em prontuário eletrônico para seguimento longitudinal adequado.',

  last_review = CURRENT_DATE
WHERE id = '1ff1e7f8-5b15-4f0c-b76b-806c1d6ff1fd';

-- ============================================================================
-- STEP 2: Insert scientific articles
-- ============================================================================

-- Article 1: Primary HPV Testing and New Technologies
INSERT INTO articles (
  title,
  authors,
  journal,
  pm_id,
  publish_date,
  article_type,
  abstract
)
SELECT
  'Primary Human Papillomavirus Testing and Other New Technologies for Cervical Cancer Screening',
  'Einstein MH, Zhou N, Gabor L, Sahasrabuddhe VV',
  'Obstet Gynecol',
  '37708516',
  '2023-10-01'::date,
  'review',
  'Revisão abrangente sobre a transição do rastreamento citológico tradicional para abordagens baseadas em HPV. Apresenta dados prospectivos robustos que permitem estratificação precisa de risco para identificar pacientes de maior risco de doença, reduzindo procedimentos desnecessários. Discute tecnologias moleculares emergentes como p16/Ki-67 dual stain e testes de metilação de HPV, que demonstram eficácia comparável ao co-teste atual. Analisa o potencial da auto-coleta domiciliar combinada com novas tecnologias para superar barreiras de acesso ao rastreamento, embora identifique desafios de implementação incluindo fluxo de trabalho, limitações de força de trabalho e custos que precisam ser resolvidos para programas mais equitativos.'
WHERE NOT EXISTS (
  SELECT 1 FROM articles WHERE pm_id = '37708516'
);

-- Article 2: Cervical Cancer Screening Recommendations
INSERT INTO articles (
  title,
  authors,
  journal,
  pm_id,
  publish_date,
  article_type,
  abstract
)
SELECT
  'Cervical Cancer Screening Recommendations: Now and for the Future',
  'Rayner M, Welp A, Stoler MH, Cantrell LA',
  'Healthcare (Basel)',
  '37628471',
  '2023-08-15'::date,
  'review',
  'Análise comparativa das recomendações globais de rastreamento de câncer cervical das principais organizações de saúde. A WHO recomenda teste de DNA-HPV a cada 5-10 anos a partir dos 30 anos, enquanto a Comissão Europeia sugere teste de HPV a cada 5+ anos para 30-65 anos. Nos EUA, as diretrizes de USPSTF, ACOG, ACS e ASCO endossam teste de HPV a cada 5 anos ou alternativas baseadas em citologia. Identifica infraestrutura limitada, restrições de recursos e lacunas na educação de profissionais como barreiras significativas. Enfatiza que estender intervalos de rastreamento com teste primário de HPV melhora eficiência mantendo segurança. Destaca dual-stain testing para pacientes HPV-positivas e métodos de auto-coleta como tecnologias emergentes para melhorar acesso, particularmente em populações vulneráveis.'
WHERE NOT EXISTS (
  SELECT 1 FROM articles WHERE pm_id = '37628471'
);

-- Article 3: Self-Collected Vaginal Specimens
INSERT INTO articles (
  title,
  authors,
  journal,
  pm_id,
  publish_date,
  article_type,
  abstract
)
SELECT
  'Self-Collected Vaginal Specimens for HPV Testing: Recommendations From the Enduring Consensus Cervical Cancer Screening and Management Guidelines Committee',
  'Wentzensen N, Perkins RB, Guido RS, Castle PE, Massad LS, Einstein MH, Smith RA, Waxman AG, Khan MJ, Chelmow D, Fontham ETH, Saraiya M, Conageski C, Feldman S, Tedeschi C, Stier EA, Mayeaux EJ, Saslow D, Schiffman M',
  'J Low Genit Tract Dis',
  '39982254',
  '2025-02-21'::date,
  'clinical_trial',
  'Recomendações do consenso de diretrizes sobre o uso de amostras vaginais auto-coletadas para teste de HPV no rastreamento de câncer cervical. Estabelece que testes de HPV em amostras auto-coletadas demonstram desempenho comparável a amostras coletadas por médicos quando ensaios baseados em PCR são utilizados. Recomenda que, embora amostras cervicais coletadas por profissionais sejam preferidas, amostras vaginais auto-coletadas são aceitáveis para rastreamento em indivíduos assintomáticos. Após resultado de HPV negativo por auto-coleta, reteste em 3 anos é recomendado como margem de segurança. Para HPV 16/18 positivo, referência direta para colposcopia é indicada. Aborda barreiras de implementação, reconhecendo que auto-coleta pode expandir acesso ao rastreamento para indivíduos enfrentando obstáculos relacionados ao sistema de saúde, profissionais ou pessoais.'
WHERE NOT EXISTS (
  SELECT 1 FROM articles WHERE pm_id = '39982254'
);

-- Article 4: Colposcopy Systematic Approach (IARC Technical Report)
INSERT INTO articles (
  title,
  authors,
  journal,
  pm_id,
  publish_date,
  article_type,
  abstract
)
SELECT
  'A systematic approach to colposcopic examination',
  'International Agency for Research on Cancer (IARC)',
  'IARC Technical Publications No. 45 - Colposcopy and Treatment of Cervical Precancer',
  '29897450',
  '2017-12-01'::date,
  'lecture',
  'Guia técnico do IARC sobre abordagem sistemática ao exame colposcópico. Estabelece que colposcopia bem-sucedida requer três elementos fundamentais: treinamento adequado, equipamento apropriado e paciente confortável. Competência envolve conhecimento teórico, reconhecimento de imagens, habilidades de manejo de casos e experiência clínica supervisionada de pelo menos 50 casos sob supervisão seguidos de 100 casos não supervisionados revisados por preceptor. Processo sistemático inclui: inserção do espéculo e visualização cervical, avaliação em baixo poder usando solução salina e filtro verde, aplicação de ácido acético seguida por exame em alto poder, documentação de achados incluindo tipo de zona de transformação e escore de Swede, aplicação de iodo para estabelecer limites da zona de transformação. Enfatiza cuidado centrado na paciente: mulher adequadamente orientada é informada, relaxada e confiante, contrastando com pacientes desinformadas que experimentam maior ansiedade e insatisfação. Enfermeira treinada em colposcopia fornece assistência essencial e suporte à paciente durante todo o exame.'
WHERE NOT EXISTS (
  SELECT 1 FROM articles WHERE pm_id = '29897450'
);

-- ============================================================================
-- STEP 3: Link articles to score item
-- ============================================================================

-- Link Article 1
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '1ff1e7f8-5b15-4f0c-b76b-806c1d6ff1fd'::uuid
FROM articles a
WHERE a.pm_id = '37708516'
AND NOT EXISTS (
  SELECT 1 FROM article_score_items asi
  WHERE asi.article_id = a.id
  AND asi.score_item_id = '1ff1e7f8-5b15-4f0c-b76b-806c1d6ff1fd'::uuid
);

-- Link Article 2
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '1ff1e7f8-5b15-4f0c-b76b-806c1d6ff1fd'::uuid
FROM articles a
WHERE a.pm_id = '37628471'
AND NOT EXISTS (
  SELECT 1 FROM article_score_items asi
  WHERE asi.article_id = a.id
  AND asi.score_item_id = '1ff1e7f8-5b15-4f0c-b76b-806c1d6ff1fd'::uuid
);

-- Link Article 3
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '1ff1e7f8-5b15-4f0c-b76b-806c1d6ff1fd'::uuid
FROM articles a
WHERE a.pm_id = '39982254'
AND NOT EXISTS (
  SELECT 1 FROM article_score_items asi
  WHERE asi.article_id = a.id
  AND asi.score_item_id = '1ff1e7f8-5b15-4f0c-b76b-806c1d6ff1fd'::uuid
);

-- Link Article 4
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '1ff1e7f8-5b15-4f0c-b76b-806c1d6ff1fd'::uuid
FROM articles a
WHERE a.pm_id = '29897450'
AND NOT EXISTS (
  SELECT 1 FROM article_score_items asi
  WHERE asi.article_id = a.id
  AND asi.score_item_id = '1ff1e7f8-5b15-4f0c-b76b-806c1d6ff1fd'::uuid
);

COMMIT;

-- ============================================================================
-- Verification queries
-- ============================================================================

-- Check updated content
SELECT
  name,
  LENGTH(clinical_relevance) as clinical_relevance_chars,
  LENGTH(patient_explanation) as patient_explanation_chars,
  LENGTH(conduct) as conduct_chars,
  last_review
FROM score_items
WHERE id = '1ff1e7f8-5b15-4f0c-b76b-806c1d6ff1fd';

-- Check linked articles
SELECT
  a.title,
  a.authors,
  a.journal,
  a.pm_id,
  a.publish_date,
  a.article_type
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '1ff1e7f8-5b15-4f0c-b76b-806c1d6ff1fd'
ORDER BY a.publish_date DESC;
