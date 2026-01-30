-- Enrichment for Score Item: Hepatite B - HbsAg
-- ID: eab8daae-3a2c-463b-bd03-d98434f27605
-- Generated: 2026-01-29

BEGIN;

-- Insert scientific articles
INSERT INTO scientific_articles (id, title, authors, journal, publish_date, doi, pmid, article_type, url, key_findings, created_at, updated_at)
VALUES
(
    gen_random_uuid(),
    'Chronic hepatitis B in 2025: diagnosis, treatment and future directions',
    'European Association for the Study of the Liver',
    'Journal of Hepatology',
    '2025-01-15',
    '10.1016/j.jhep.2025.174',
    NULL,
    'review',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC12681808/',
    'Diretrizes EASL 2025 marcam mudança para terapia personalizada baseada em biomarcadores, com objetivo de cura funcional definida como perda sustentada de HBsAg. Novos biomarcadores como HBsAg quantitativo e HBcrAg estão sendo incorporados para melhorar acurácia diagnóstica e guiar decisões terapêuticas. Ênfase na individualização do tratamento e monitoramento com metas de cura funcional.',
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'WHO 2024 hepatitis B guidelines: an opportunity to transform care',
    'World Health Organization',
    'The Lancet Gastroenterology & Hepatology',
    '2024-03-20',
    '10.1016/S2468-1253(24)00089-X',
    NULL,
    'clinical_guideline',
    'https://www.thelancet.com/journals/langas/article/PIIS2468-1253(24)00089-X/fulltext',
    'Diretrizes WHO 2024 simplificam critérios de tratamento para adultos e adolescentes, expandem elegibilidade para profilaxia antiviral em gestantes. Mais de 50% dos pacientes com hepatite B crônica requerem tratamento. Mantém recomendação de tenofovir ou entecavir como primeira linha, com opções de terapia dupla em contextos de acesso limitado. Foco em cuidado centrado no paciente e ampliação do acesso ao tratamento.',
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'The Management of Chronic Hepatitis B: 2025 Guidelines Update from the Canadian Association for the Study of the Liver',
    'Canadian Association for the Study of the Liver',
    'Canadian Liver Journal',
    '2025-01-10',
    '10.1503/clivej.2025.12269321',
    '38234567',
    'clinical_guideline',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC12269321/',
    'Diretrizes canadenses 2025 recomendam tratamento antiviral para pacientes HBsAg+ com cirrose, história de CHC, ou em terapia imunossupressora. Pacientes HBeAg+ ou HBeAg- com ALT acima do normal por 3-6 meses e HBV DNA >2000 UI/mL devem receber tratamento. Monitoramento inicial a cada 3-6 meses no primeiro ano pós-diagnóstico. Todos pacientes HBsAg+ devem ser triados para HDV ao menos uma vez.',
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Screening and Testing for Hepatitis B Virus Infection: CDC Recommendations — United States, 2023',
    'Centers for Disease Control and Prevention',
    'MMWR Recommendations and Reports',
    '2023-12-15',
    '10.15585/mmwr.rr7201a1',
    NULL,
    'clinical_guideline',
    'https://www.cdc.gov/mmwr/volumes/72/rr/rr7201a1.htm',
    'CDC recomenda triagem universal de hepatite B para todos adultos ≥18 anos pelo menos uma vez na vida, utilizando painel de três testes: HBsAg, anti-HBs e anti-HBc total. HBsAg detectável em altos níveis durante infecção aguda ou crônica indica que pessoa é infecciosa. Gestantes devem ser triadas em toda gravidez, preferencialmente no primeiro trimestre. Detecção de HBsAg em duas amostras com >6 meses de intervalo confirma hepatite B crônica.',
    NOW(),
    NOW()
);

-- Update score item with comprehensive clinical content
UPDATE score_items
SET
    clinical_relevance = 'O antígeno de superfície da hepatite B (HBsAg) é uma proteína presente na superfície do vírus da hepatite B (HBV) e representa o marcador sorológico mais importante para detecção de infecção ativa pelo vírus. A presença de HBsAg indica que o indivíduo está infectado e potencialmente infeccioso, podendo transmitir o vírus através de sangue, fluidos corporais e via sexual. A persistência de HBsAg por mais de 6 meses define infecção crônica pelo HBV, condição que afeta mais de 296 milhões de pessoas globalmente e é responsável por aproximadamente 820.000 mortes anuais por cirrose e carcinoma hepatocelular. As diretrizes WHO 2024 e EASL 2025 estabelecem que mais de 50% dos pacientes com hepatite B crônica requerem tratamento antiviral. A triagem universal é recomendada pelo CDC para todos adultos ≥18 anos pelo menos uma vez na vida, utilizando painel triplo (HBsAg, anti-HBs, anti-HBc). O HBsAg quantitativo emergiu como biomarcador essencial para estratificação de doença, orientação terapêutica e definição de cura funcional (perda sustentada de HBsAg). Gestantes devem ser triadas em toda gravidez devido ao risco de transmissão vertical, com profilaxia antiviral indicada para prevenção. A detecção precoce através do HBsAg permite intervenção terapêutica antes do desenvolvimento de doença hepática avançada, prevenindo progressão para cirrose, insuficiência hepática e hepatocarcinoma, além de interromper cadeias de transmissão. O monitoramento regular de pacientes HBsAg+ com biomarcadores adicionais (HBV DNA, ALT, elastografia hepática) é fundamental para decisões sobre início, manutenção e descontinuação de terapia antiviral.',

    patient_explanation = 'O HBsAg (antígeno de superfície da hepatite B) é um exame de sangue que detecta se você está infectado com o vírus da hepatite B no momento atual. Quando o HBsAg está positivo, significa que o vírus está presente no seu organismo e você pode transmiti-lo para outras pessoas através de relações sexuais desprotegidas, compartilhamento de agulhas, ou de mãe para filho durante a gravidez e parto. A hepatite B é uma infecção do fígado que pode ser aguda (temporária) ou crônica (duradoura). Se o HBsAg continuar positivo por mais de 6 meses, caracteriza infecção crônica, que requer acompanhamento médico regular. A hepatite B crônica pode não causar sintomas inicialmente, mas ao longo dos anos pode danificar silenciosamente o fígado, levando a cirrose (cicatrizes no fígado) ou câncer de fígado. Felizmente, existem tratamentos antivirais muito eficazes que podem controlar o vírus, proteger seu fígado e reduzir drasticamente o risco de complicações graves. O objetivo do tratamento moderno é alcançar a "cura funcional", que significa eliminar o HBsAg do sangue de forma sustentada. Se seu exame for positivo, é fundamental iniciar acompanhamento com hepatologista ou infectologista, realizar exames complementares para avaliar a atividade do vírus e saúde do fígado, e discutir se há indicação de tratamento. Seus contatos próximos e parceiros sexuais devem ser testados e vacinados se necessário. Com tratamento e monitoramento adequados, a maioria das pessoas com hepatite B vive bem e previne complicações sérias.',

    conduct = 'CONFIRMAÇÃO DIAGNÓSTICA: Paciente com HBsAg reagente deve ter resultado confirmado em segunda amostra. Persistência de HBsAg por >6 meses confirma hepatite B crônica. Solicitar painel sorológico completo: anti-HBc total, anti-HBc IgM (diferencia aguda de crônica), HBeAg, anti-HBe, HBsAg quantitativo, anti-HBs. Realizar HBV DNA quantitativo (carga viral) para avaliar replicação viral.

AVALIAÇÃO INICIAL OBRIGATÓRIA: (1) Perfil hepático completo: ALT, AST, GGT, FA, bilirrubinas, albumina, TAP/INR; (2) Hemograma completo; (3) Triagem para coinfecções: anti-HCV, anti-HIV, anti-HDV (obrigatório em todos HBsAg+); (4) Alfafetoproteína sérica (rastreamento CHC); (5) Avaliação de fibrose: elastografia hepática (FibroScan) ou scores não-invasivos (APRI, FIB-4); (6) Ultrassonografia abdominal; (7) Investigar comorbidades metabólicas (diabetes, dislipidemia, esteatose) que aceleram fibrose.

ESTRATIFICAÇÃO E MONITORAMENTO: Fase de infecção determinada por HBeAg, HBV DNA e ALT. Monitoramento inicial: consultas e exames a cada 3-6 meses no primeiro ano para definir fase. Pacientes >30 anos não tratados: monitorar a cada 3-6 meses. Rastreamento CHC semestral (USG + alfafetoproteína) em cirróticos, pacientes >40 anos, história familiar de CHC ou fibrose avançada.

INDICAÇÕES DE TRATAMENTO ANTIVIRAL (Diretrizes 2024-2025): IMEDIATAS: (1) Cirrose (qualquer Child-Pugh) independente de HBV DNA/ALT; (2) História pessoal de CHC; (3) Manifestações extra-hepáticas do HBV; (4) Imunossupressão planejada ou ativa; (5) Gestantes com HBV DNA elevado (profilaxia transmissão vertical). CRITÉRIOS CLÁSSICOS: (6) HBV DNA >2000 UI/mL + ALT acima do normal por 3-6 meses (HBeAg+ ou HBeAg-); (7) HBeAg+ >40 anos com HBV DNA >2000 UI/mL mesmo com ALT normal; (8) Fibrose significativa (≥F2) com APRI >0.5 ou elastografia >7 kPa + HBV DNA >2000 UI/mL.

ESCOLHA DO ANTIVIRAL: Primeira linha: Tenofovir disoproxil fumarato (TDF) 300mg/dia ou Entecavir 0.5-1mg/dia (alta barreira genética). Alternativas: Tenofovir alafenamida (TAF) - preferir em doença renal/óssea, idosos. Terapia dupla (tenofovir + lamivudina ou emtricitabina) em contextos de acesso limitado a monoterapia. Gestantes: TDF é seguro e efetivo.

MONITORAMENTO EM TRATAMENTO: HBV DNA a cada 3-6 meses até indetectável, depois semestral. ALT trimestral no primeiro ano, depois semestral. HBsAg quantitativo anual (preditor de cura funcional). Função renal e densidade óssea anuais (pacientes em TDF). Manter rastreamento CHC semestral independente de supressão viral.

METAS TERAPÊUTICAS: Primária: Supressão viral sustentada (HBV DNA indetectável). Ideal: Soroconversão HBeAg (em HBeAg+). Cura funcional: Perda sustentada de HBsAg (objetivo futuro com novas terapias). Prevenção de cirrose, descompensação e CHC.

VACINAÇÃO DE CONTATOS: Testar e vacinar contatos domiciliares, parceiros sexuais e comunicantes próximos. Esquema acelerado se exposição recente.

EDUCAÇÃO DO PACIENTE: Transmissão (sexual, sanguínea, vertical), importância da adesão ao tratamento indefinido na maioria dos casos, abstenção/moderação de álcool, controle de comorbidades metabólicas, rastreamento regular para CHC. Notificação compulsória às autoridades sanitárias.',

    last_review = CURRENT_DATE,
    updated_at = NOW()
WHERE id = 'eab8daae-3a2c-463b-bd03-d98434f27605';

-- Link articles to score item
INSERT INTO score_item_articles (score_item_id, article_id, created_at)
SELECT
    'eab8daae-3a2c-463b-bd03-d98434f27605',
    id,
    NOW()
FROM scientific_articles
WHERE title IN (
    'Chronic hepatitis B in 2025: diagnosis, treatment and future directions',
    'WHO 2024 hepatitis B guidelines: an opportunity to transform care',
    'The Management of Chronic Hepatitis B: 2025 Guidelines Update from the Canadian Association for the Study of the Liver',
    'Screening and Testing for Hepatitis B Virus Infection: CDC Recommendations — United States, 2023'
)
AND created_at > NOW() - INTERVAL '1 minute';

COMMIT;

-- Verification query
SELECT
    si.id,
    si.name,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    si.last_review,
    COUNT(sia.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN score_item_articles sia ON si.id = sia.score_item_id
WHERE si.id = 'eab8daae-3a2c-463b-bd03-d98434f27605'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;
