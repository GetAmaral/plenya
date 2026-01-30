-- Enriquecimento do Score Item: Urobilinogênio
-- ID: bf77b326-caa5-46fd-b607-70a089918780
-- Data: 2026-01-29

BEGIN;

-- 1. Atualizar campos clínicos do score item
UPDATE score_items
SET
    clinical_relevance = 'O urobilinogênio é um metabólito da bilirrubina formado pela ação de enzimas bacterianas intestinais, especialmente pela BilR de espécies Firmicutes, sobre a bilirrubina conjugada excretada na bile. Sua detecção e quantificação na urina fornecem informações diagnósticas valiosas sobre três sistemas integrados: função hepática, integridade do ciclo êntero-hepático e taxa de hemólise. Valores normais variam de 0,1 a 1,0 mg/dL na urina. A fisiopatologia envolve a circulação êntero-hepática: cerca de 20% do urobilinogênio produzido no intestino é reabsorvido, retorna ao fígado pela veia porta, onde 95-98% é recaptado e reexcretado na bile, enquanto 2-5% escapa para circulação sistêmica e é filtrado pelos rins. Níveis elevados ocorrem em anemia hemolítica (aumento de 6-10 vezes na produção de bilirrubina), doença hepatocelular (redução da recaptação hepática), shunts portossistêmicos e hemólise extravascular aumentada. A ausência ou redução acentuada sugere obstrução biliar completa, colestase grave, uso de antibióticos que alteram microbiota intestinal ou disbiose. A combinação de urobilinogênio urinário com bilirrubina conjugada permite diferenciação diagnóstica: bilirrubina presente + urobilinogênio elevado indica doença hepatocelular; bilirrubina presente + urobilinogênio ausente sugere obstrução biliar; bilirrubina ausente + urobilinogênio elevado aponta hemólise. Descobertas recentes (2024) identificaram a enzima BilR como responsável pela conversão de bilirrubina em urobilinogênio, cuja prevalência está reduzida em neonatos e pacientes com doença inflamatória intestinal, explicando alterações metabólicas nesses grupos. O urobilinogênio possui valor prognóstico emergente como biomarcador precoce de síndrome cardiometabólica (CKM), obesidade e diabetes tipo 2.',

    patient_explanation = 'Urobilinogênio é uma substância amarelada formada no intestino quando bactérias "amigáveis" transformam a bilirrubina que vem da bile. A bilirrubina, por sua vez, é produzida quando glóbulos vermelhos velhos são quebrados naturalmente no corpo. O teste mede quanto urobilinogênio está presente na urina, geralmente através do exame de urina tipo 1 (EAS). Valores normais são de 0,1 a 1,0 mg/dL. Quando está aumentado pode indicar: anemia hemolítica (destruição acelerada de glóbulos vermelhos), hepatite ou cirrose (fígado não consegue reprocessar o urobilinogênio adequadamente), ou infecções graves. Quando está ausente ou muito diminuído pode significar obstrução das vias biliares (cálculos, tumores pancreáticos), uso recente de antibióticos fortes que mataram as bactérias intestinais responsáveis pela conversão, ou problemas graves no fígado. A vantagem deste exame é que ele pode detectar problemas no fígado antes mesmo da pele ficar amarelada (icterícia). É um teste barato, rápido e não invasivo que fornece informações importantes sobre o funcionamento do fígado, vesícula biliar e ritmo de renovação do sangue. Alterações neste parâmetro sempre devem ser interpretadas junto com outros marcadores da urina (como presença de bilirrubina) e exames de sangue para determinar a causa exata.',

    conduct = 'AVALIAÇÃO INICIAL: Confirmar resultado com nova amostra de urina fresca (primeira urina da manhã preferencial), evitando exposição prolongada à luz que oxida urobilinogênio em urobilina. Correlacionar com sintomas clínicos: icterícia, colúria (urina escura), acolia fecal (fezes claras), prurido, dor abdominal, fadiga. Verificar uso de medicamentos fotossensibilizantes, antibióticos de amplo espectro, fenazina ou corantes que geram falsos positivos/negativos. UROBILINOGÊNIO ELEVADO (>1,0 mg/dL): Solicitar hemograma completo com contagem de reticulócitos, bilirrubinas totais e frações, TGO/TGP, FA, GGT, LDH, haptoglobina sérica. Investigar anemia hemolítica: pesquisa de esferócitos, teste de Coombs direto, dosagem de G6PD, eletroforese de hemoglobina. Avaliar doença hepatocelular: sorologias virais (hepatites A, B, C), autoanticorpos (ANA, anti-músculo liso, anti-LKM), ultrassonografia abdominal, considerar elastografia hepática ou biópsia se indicado. Descartar sepse, malária em áreas endêmicas, hematomas extensos em reabsorção. UROBILINOGÊNIO AUSENTE OU REDUZIDO (<0,1 mg/dL): Priorizar investigação de obstrução biliar: ultrassonografia de abdome superior, se positiva prosseguir com colangioressonância magnética ou ecoendoscopia. Dosar bilirrubinas totais e frações (elevação de conjugada), enzimas canaliculares (FA, GGT marcadamente elevadas). Avaliar história de uso recente de antibióticos (ciprofloxacina, metronidazol) que reduzem microbiota produtora de urobilinogênio. Considerar colangite esclerosante primária, coledocolitíase, neoplasias pancreáticas ou de via biliar. MONITORAMENTO: Repetir urinálise semanalmente durante investigação. Acompanhar evolução com perfil hepático completo a cada 2-4 semanas. Em obstrução biliar confirmada, encaminhar para procedimento descompressivo (CPRE com papilotomia, drenagem percutânea ou cirurgia). Em hepatopatias, otimizar tratamento da causa base. Orientar suspensão de hepatotóxicos, álcool, ajustar doses de medicamentos de metabolização hepática. Considerar avaliação da microbiota intestinal em casos de disbiose documentada.',

    updated_at = NOW()
WHERE id = 'bf77b326-caa5-46fd-b607-70a089918780';

-- 2. Inserir artigos científicos (verificando duplicatas antes)

-- Artigo 1: BilR enzyme discovery (Nature Microbiology, 2024)
INSERT INTO articles (id, title, authors, journal, publish_date, pm_id, article_type, notes)
SELECT
    gen_random_uuid(),
    'BilR is a gut microbial enzyme that reduces bilirubin to urobilinogen',
    'Hall B, Levy S, Dufault-Thompson K, Arp G, Zhong A, Ndjite GM, Weiss A, Braccia DJ, Jenkins C, Grant C, Hylemon PB, Sikaroodi M, Gillevet PM, Dailey HA, Narayan ARH, Balskus EP',
    'Nature Microbiology',
    '2024-01-03'::date,
    '38172624',
    'research_article',
    'Identificação da enzima bacteriana BilR (bilirrubina redutase) produzida principalmente por espécies Firmicutes, responsável pela conversão de bilirrubina em urobilinogênio no intestino. Prevalência reduzida em neonatos e pacientes com doença inflamatória intestinal. Descoberta fundamental para compreender o eixo intestino-fígado e potencial alvo terapêutico para doenças cardiometabólicas.'
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '38172624'
);

-- Artigo 2: Bilirubin measurement in liver disease (PMC, 2021)
INSERT INTO articles (id, title, authors, journal, publish_date, pm_id, article_type, notes)
SELECT
    gen_random_uuid(),
    'Measurement and clinical usefulness of bilirubin in liver disease',
    'Guerra Ruiz AR, Crespo J, López Martínez RM, Iruzubieta P, Casals Mercadal G, Lalana Garcés M, Lavin B, Morales Ruiz M',
    'Advances in Laboratory Medicine',
    '2021-11-01'::date,
    '37362415',
    'review',
    'Revisão abrangente dos métodos de medição de bilirrubina (método diazo como padrão-ouro, HPLC, métodos enzimáticos) e interpretação clínica. Classificação da hiperbilirrubinemia em pré-hepática, hepática e pós-hepática. Enfatiza que bilirrubina não é marcador sensível ou específico isolado, requerendo interpretação contextual com história clínica e outros parâmetros laboratoriais.'
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '37362415'
);

-- Artigo 3: Bilirubinuria clinical guide (StatPearls, 2023)
INSERT INTO articles (id, title, authors, journal, publish_date, article_type, notes)
SELECT
    gen_random_uuid(),
    'Bilirubinuria',
    'Hoilat GJ, John S',
    'StatPearls [Internet]',
    '2023-08-08'::date,
    'review',
    'Guia clínico sobre bilirubinúria: apenas bilirrubina conjugada (hidrossolúvel) aparece na urina. Causas intra-hepáticas incluem hepatite viral, cirrose, lesão hepática induzida por drogas. Causas extra-hepáticas incluem coledocolitíase, estenoses biliares, neoplasias pancreáticas. Urina cor de chá/cola é apresentação clássica. Detecção precoce via dipstick antes de icterícia clínica aparente.'
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE title = 'Bilirubinuria' AND authors LIKE '%Hoilat%'
);

-- Artigo 4: Multifaceted role of bilirubin (J Clin Transl Hepatol, 2024)
INSERT INTO articles (id, title, authors, journal, publish_date, pm_id, article_type, notes)
SELECT
    gen_random_uuid(),
    'The Multifaceted Role of Bilirubin in Liver Disease: A Literature Review',
    'Ramírez-Mejía MM, Castillo-Castañeda SM, Pal SC, Qi X, Méndez-Sánchez N',
    'Journal of Clinical and Translational Hepatology',
    '2024-10-28'::date,
    '39544246',
    'review',
    'Revisão do papel multifacetado da bilirrubina como biomarcador diagnóstico e prognóstico (integrado aos escores MELD, Child-Pugh), agente antioxidante potente, modulador anti-inflamatório e regulador metabólico. Produção diária de 250-400 mg de hemoglobina. Propriedades protetoras cardiovasculares e metabólicas da hiperbilirrubinemia leve (síndrome de Gilbert). Neurotoxicidade em níveis severos.'
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '39544246'
);

-- 3. Criar vínculos entre artigos e score item (tabela de junção: article_score_items)

-- Vincular artigo BilR enzyme
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'bf77b326-caa5-46fd-b607-70a089918780'
FROM articles a
WHERE a.pm_id = '38172624'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Vincular artigo bilirubin measurement
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'bf77b326-caa5-46fd-b607-70a089918780'
FROM articles a
WHERE a.pm_id = '37362415'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Vincular artigo bilirubinuria (sem PMID, buscar por título)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'bf77b326-caa5-46fd-b607-70a089918780'
FROM articles a
WHERE a.title = 'Bilirubinuria' AND a.authors LIKE '%Hoilat%'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Vincular artigo multifaceted role
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'bf77b326-caa5-46fd-b607-70a089918780'
FROM articles a
WHERE a.pm_id = '39544246'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

COMMIT;

-- Verificação final
SELECT
    si.name,
    LENGTH(si.clinical_relevance) as len_clinical_relevance,
    LENGTH(si.patient_explanation) as len_patient_explanation,
    LENGTH(si.conduct) as len_conduct,
    COUNT(asi.article_id) as num_articles
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = 'bf77b326-caa5-46fd-b607-70a089918780'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct;
