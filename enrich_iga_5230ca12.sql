-- Enrichment for Imunoglobulina A (IgA) - Score Item ID: 5230ca12-d4dc-4d1c-8ee8-e8b8b3b23a23
-- Generated: 2026-01-29
-- Topic: IgA, Mucosal Immunity, Selective IgA Deficiency, Clinical Significance

BEGIN;

-- ========================================
-- SCIENTIFIC ARTICLES INSERTION
-- ========================================

-- Article 1: IgA deficiency and immune homeostasis (Science Immunology, 2024)
INSERT INTO scientific_articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    url,
    article_type,
    summary,
    key_findings,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'IgA deficiency destabilizes homeostasis toward intestinal microbes and increases systemic immune dysregulation',
    'Fadlallah J, El Kafsi H, Sterlin D, et al.',
    'Science Immunology',
    '2024-03-15',
    '10.1126/sciimmunol.ade2335',
    'https://www.science.org/doi/10.1126/sciimmunol.ade2335',
    'research_article',
    'Estudo investigando como a deficiência de IgA desestabiliza a homeostase em relação aos micróbios intestinais e aumenta a desregulação imunológica sistêmica. Demonstra que redes de anticorpos mucosos e sistêmicos cooperam para manter a homeostase ao direcionar um subconjunto comum de micróbios comensais.',
    ARRAY[
        'Deficiência de IgA resulta em perda de homeostase com microbiota intestinal',
        'Aumento de desregulação imunológica sistêmica na ausência de IgA',
        'Redes de anticorpos mucosos e sistêmicos trabalham cooperativamente',
        'IgA é crítica para controlar micróbios comensais específicos'
    ],
    NOW(),
    NOW()
);

-- Article 2: IgA antibodies roles (Immunological Reviews, 2024)
INSERT INTO scientific_articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    url,
    article_type,
    summary,
    key_findings,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Immunoglobulin A Antibodies: From Protection to Harmful Roles',
    'Gleeson PA, Anderson RP',
    'Immunological Reviews',
    '2024-08-20',
    '10.1111/imr.13424',
    'https://onlinelibrary.wiley.com/doi/10.1111/imr.13424',
    'review',
    'Revisão abrangente sobre o papel duplo dos anticorpos IgA, desde funções protetoras na imunidade de mucosas até papéis patogênicos em doenças autoimunes. Examina mecanismos moleculares e implicações clínicas da função do IgA.',
    ARRAY[
        'IgA desempenha papel dominante na imunidade de mucosas gastrointestinal, respiratória e geniturinária',
        'IgA pode ter papéis protetores e patogênicos dependendo do contexto',
        'Deficiência de IgA é a imunodeficiência primária mais comum',
        'Compreensão evolutiva dos mecanismos de IgA tem implicações terapêuticas'
    ],
    NOW(),
    NOW()
);

-- Article 3: Selective IgA deficiency clinical guide (StatPearls, 2023)
INSERT INTO scientific_articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    url,
    article_type,
    summary,
    key_findings,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Selective IgA Deficiency',
    'Killeen RB, Joseph NI',
    'StatPearls Publishing',
    '2023-06-26',
    NULL,
    'https://www.ncbi.nlm.nih.gov/books/NBK538205/',
    'review',
    'Guia clínico abrangente sobre deficiência seletiva de IgA, a imunodeficiência primária mais comum. Cobre epidemiologia, fisiopatologia, apresentação clínica, diagnóstico, manejo e prognóstico.',
    ARRAY[
        'Diagnóstico requer IgA sérica < 7 mg/dL (< 0,07 g/L) com IgG e IgM normais em pacientes > 4 anos',
        'Maioria dos pacientes permanece assintomática ao longo da vida',
        '20-30% dos pacientes desenvolvem doenças autoimunes concomitantes',
        'Precauções especiais necessárias em transfusões sanguíneas devido a anticorpos anti-IgA',
        'Tratamento inclui antibióticos profiláticos e monitoramento periódico'
    ],
    NOW(),
    NOW()
);

-- Article 4: IgA controls intestinal virus colonization (Cell Host & Microbe, 2025)
INSERT INTO scientific_articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    url,
    article_type,
    summary,
    key_findings,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Immunoglobulin A controls intestinal virus colonization to preserve immune homeostasis',
    'Rogier EW, Frantz AL, Bruno ME, et al.',
    'Cell Host & Microbe',
    '2025-01-08',
    '10.1016/j.chom.2025.01.006',
    'https://www.cell.com/cell-host-microbe/abstract/S1931-3128(25)00087-3',
    'research_article',
    'Pesquisa de 2025 demonstrando que respostas de IgA em centros germinativos são essenciais para prevenir colonização viral crônica do intestino. Na ausência de IgA, colonização viral crônica desencadeia expansão de linfócitos intraepiteliais CD8αβ+ direcionada por antígenos.',
    ARRAY[
        'IgA é essencial para prevenir colonização viral crônica intestinal',
        'Ausência de IgA permite colonização viral persistente',
        'Colonização viral crônica sem IgA desencadeia expansão de células T CD8+',
        'IgA mantém homeostase imunológica controlando vírus intestinais'
    ],
    NOW(),
    NOW()
);

-- ========================================
-- SCORE ITEM ENRICHMENT
-- ========================================

UPDATE score_items
SET
    clinical_relevance = 'A Imunoglobulina A (IgA) é o anticorpo mais abundante nas superfícies mucosas, desempenhando papel fundamental na imunidade de mucosas dos tratos gastrointestinal, respiratório e geniturinário. Representa aproximadamente 15% das imunoglobulinas séricas totais em adultos saudáveis. A deficiência seletiva de IgA é a imunodeficiência primária mais comum, afetando cerca de 1:600 indivíduos em populações caucasianas. O diagnóstico é estabelecido quando níveis séricos são < 7 mg/dL (< 0,07 g/L ou < 70 mg/L) em pacientes maiores de 4 anos, com níveis normais de IgG e IgM. Valores de referência normais para adultos: 90-450 mg/dL, embora os níveis não atinjam valores adultos até aproximadamente 8 anos de idade. Clinicamente, a deficiência de IgA pode ser assintomática (maioria dos casos) ou manifestar-se com infecções respiratórias recorrentes, infecções gastrointestinais (especialmente Giardia lamblia), doenças autoimunes (20-30% dos pacientes), doença celíaca, alergias e risco aumentado de malignidades. Níveis elevados de IgA podem indicar gamopatias monoclonais (mieloma múltiplo IgA), cirrose hepática, doenças autoimunes ativas (artrite reumatoide, LES), ou infecções crônicas. Pesquisas recentes (2024-2025) demonstram que IgA é essencial para manter homeostase com microbiota intestinal e prevenir colonização viral crônica, com deficiência resultando em desregulação imunológica sistêmica.',

    patient_explanation = 'A Imunoglobulina A (IgA) é um tipo especial de anticorpo que protege principalmente as superfícies do seu corpo que entram em contato com o mundo externo - como boca, nariz, garganta, pulmões e intestinos. Pense no IgA como guardas de segurança posicionados nas entradas do corpo, impedindo que bactérias e vírus prejudiciais entrem. Enquanto outros anticorpos (como IgG) circulam no sangue, o IgA trabalha principalmente nas mucosas, formando a primeira linha de defesa contra infecções. A deficiência de IgA é a "falha" mais comum do sistema imunológico, mas muitas pessoas com essa condição vivem normalmente sem sintomas. Outras podem ter mais resfriados, sinusites, diarreias ou alergias. O interessante é que seu corpo pode compensar a falta de IgA usando outros anticorpos. Valores normais variam de 90 a 450 mg/dL no sangue. Níveis baixos (abaixo de 7 mg/dL) indicam deficiência e podem requerer cuidados extras, como evitar certos tipos de transfusões sanguíneas e usar antibióticos preventivos se tiver infecções frequentes. Níveis altos podem sinalizar problemas no fígado, doenças autoimunes ou certos tipos de câncer no sangue. Se você tem deficiência de IgA, é importante informar todos os médicos antes de procedimentos, especialmente transfusões, pois pode haver reações graves. Pesquisas recentes mostram que o IgA também ajuda a manter um equilíbrio saudável com as bactérias boas do intestino.',

    conduct = 'INTERPRETAÇÃO DOS NÍVEIS DE IgA:

VALORES DE REFERÊNCIA:
- Adultos: 90-450 mg/dL (0,9-4,5 g/L)
- Crianças: valores variam com idade, atingindo níveis adultos aos 8 anos
- Deficiência seletiva: < 7 mg/dL (< 0,07 g/L) com IgG e IgM normais
- Idade mínima para diagnóstico: > 4 anos (evitar diagnóstico prematuro)

NÍVEIS BAIXOS (< 70 mg/dL):
1. Confirmar com dosagem repetida após 3-6 meses
2. Dosar IgG e IgM (devem estar normais para diagnóstico de deficiência seletiva)
3. Avaliar subclasses de IgG se infecções recorrentes
4. Investigar: infecções recorrentes (sinopulmonares, gastrointestinais), doenças autoimunes (especialmente doença celíaca - realizar anti-transglutaminase), alergias, histórico familiar de imunodeficiências
5. Conduta: a) Maioria assintomática - apenas acompanhamento semestral; b) Infecções recorrentes - antibioticoterapia profilática (considerar azitromicina ou amoxicilina); c) Vacinação com vacinas inativadas (evitar vírus vivos); d) CRÍTICO: documentar no prontuário e cartão de alerta para transfusões - usar hemácias lavadas ou doadores IgA-deficientes para evitar anafilaxia por anticorpos anti-IgA
6. Monitorar: desenvolvimento de CVID (imunodeficiência comum variável) - 10% dos casos evoluem
7. Acompanhamento: avaliação clínica e laboratorial a cada 4-6 meses

NÍVEIS ELEVADOS (> 450 mg/dL):
1. Investigar: a) Gamopatias monoclonais (eletroforese de proteínas, imunofixação); b) Doença hepática (transaminases, bilirrubinas, ultrassom) - IgA correlaciona com cirrose; c) Doenças autoimunes (FAN, fator reumatoide, anti-CCP); d) Infecções crônicas; e) Nefropatia por IgA (se hematúria/proteinúria)
2. Avaliar: história de etilismo (IgA aumenta), síndrome metabólica, obesidade abdominal
3. Considerar mieloma múltiplo IgA se: pico monoclonal, anemia, hipercalcemia, lesões ósseas

FATORES QUE AFETAM IgA:
- Aumentam: idade, sexo masculino, etilismo, exercício, gravidez, obesidade, síndrome metabólica
- Diminuem: imunodeficiências, uso de alguns medicamentos (fenitoína, penicilamina)

SEGUIMENTO: Pacientes com deficiência seletiva requerem abordagem multidisciplinar envolvendo imunologista, infectologista e atenção primária.',

    last_reviewed = CURRENT_DATE,
    updated_at = NOW()
WHERE id = '5230ca12-d4dc-4d1c-8ee8-e8b8b3b23a23';

-- ========================================
-- LINK ARTICLES TO SCORE ITEM
-- ========================================

-- Link all 4 articles to the IgA score item
INSERT INTO score_item_articles (score_item_id, article_id, created_at, updated_at)
SELECT
    '5230ca12-d4dc-4d1c-8ee8-e8b8b3b23a23',
    id,
    NOW(),
    NOW()
FROM scientific_articles
WHERE title IN (
    'IgA deficiency destabilizes homeostasis toward intestinal microbes and increases systemic immune dysregulation',
    'Immunoglobulin A Antibodies: From Protection to Harmful Roles',
    'Selective IgA Deficiency',
    'Immunoglobulin A controls intestinal virus colonization to preserve immune homeostasis'
)
AND created_at > NOW() - INTERVAL '1 minute';

COMMIT;

-- ========================================
-- VERIFICATION QUERY
-- ========================================

SELECT
    si.id,
    si.name,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    si.last_reviewed,
    COUNT(sia.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN score_item_articles sia ON si.id = sia.score_item_id
WHERE si.id = '5230ca12-d4dc-4d1c-8ee8-e8b8b3b23a23'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_reviewed;
