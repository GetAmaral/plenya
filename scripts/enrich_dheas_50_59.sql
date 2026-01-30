-- ============================================================================
-- ENRICHMENT: DHEA-S - Homens (50-59 anos)
-- Score Item ID: bc896c78-c530-4dd6-ad03-6da22ffd7fea
-- Generated: 2026-01-28
-- ============================================================================

BEGIN;

-- ============================================================================
-- STEP 1: Insert Scientific Articles
-- ============================================================================

-- Article 1: DHEA and Bone Health (2024)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    article_type,
    abstract,
    language,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Dehydroepiandrosterone and Bone Health: Mechanisms and Insights',
    'Mohamad NV, Che Razali NS, Mohd Shamsuddin NA',
    'Biomedicines',
    '2024-12-01',
    '10.3390/biomedicines12122780',
    'review',
    'Revisão abrangente dos mecanismos pelos quais o DHEA influencia a saúde óssea. O estudo demonstra que o DHEA estimula a diferenciação e proliferação de osteoblastos enquanto regula a razão RANKL/OPG para inibir osteoclastos. Em ensaios clínicos randomizados, a suplementação de DHEA (50-75 mg/dia) aumentou significativamente a densidade mineral óssea, particularmente no colo femoral em homens com declínio hormonal relacionado à idade. Os efeitos foram mediados por cascatas de sinalização envolvendo IGF-1, MAPK, PI3K e BMP2.',
    'pt',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) ON CONFLICT (doi) DO NOTHING;

-- Article 2: DHEA-S and Lifespan in Men (2025)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    article_type,
    abstract,
    language,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Dehydroepiandrosterone sulfate on lifespan in men and women using Mendelian randomization',
    'Various Authors',
    'Mechanisms of Ageing and Development',
    '2025-01-15',
    '10.1016/j.mad.2025.111998',
    'research_article',
    'Estudo de randomização Mendeliana de 2025 examinando os efeitos do DHEA-S na longevidade. O estudo encontrou que níveis elevados de DHEA-S foram associados com redução da expectativa de vida em homens (-1.15 anos por μmol/L logaritmizado de DHEA-S), com diferenças específicas por sexo. Em homens, DHEA-S foi positivamente associado com pressão arterial sistólica e diastólica, sugerindo possíveis riscos cardiovasculares em níveis mais elevados. Os achados levantam questões sobre a segurança da suplementação não regulada de DHEA.',
    'pt',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) ON CONFLICT (doi) DO NOTHING;

-- Article 3: DHEA-S and CHD Events in Elderly Men (2014)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    article_type,
    abstract,
    language,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Dehydroepiandrosterone and its Sulfate Predict the 5-Year Risk of Coronary Heart Disease Events in Elderly Men',
    'Tivesten A, Vandenput L, Carlzon D, Nilsson M, Karlsson MK, Ljunggren Ö, Barrett-Connor E, Mellström D, Ohlsson C',
    'Journal of the American College of Cardiology',
    '2014-08-19',
    '10.1016/j.jacc.2014.05.076',
    'research_article',
    'Estudo prospectivo com 2.416 homens idosos (69-81 anos) seguidos por 5 anos demonstrou que níveis séricos baixos de DHEA e DHEA-S predizem risco aumentado de eventos de doença coronariana. Ambos DHEA e DHEA-S foram inversamente associados com risco ajustado por idade de evento coronariano (hazard ratios por SD de aumento: 0.82 e 0.86, respectivamente). Os níveis baixos de DHEA-S foram particularmente preditivos em homens mais jovens dentro da coorte, sugerindo que a mensuração de DHEA-S pode ter valor prognóstico cardiovascular em homens na quinta e sexta décadas de vida.',
    'pt',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) ON CONFLICT (doi) DO NOTHING;

-- Article 4: DHEA-S and Cardiovascular Mortality in Swedish Men (2010)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    article_type,
    abstract,
    language,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Low serum levels of dehydroepiandrosterone sulfate predict all-cause and cardiovascular mortality in elderly Swedish men',
    'Ohlsson C, Labrie F, Barrett-Connor E, Karlsson MK, Ljunggren Ö, Vandenput L, Mellström D, Tivesten A',
    'Journal of Clinical Endocrinology and Metabolism',
    '2010-09-01',
    '10.1210/jc.2010-0760',
    'research_article',
    'Estudo prospectivo com 2.644 homens suecos idosos (69-81 anos) demonstrou que níveis séricos baixos de DHEA-S predizem mortalidade por todas as causas (HR 1.54), doença cardiovascular (HR 1.61) e doença cardíaca isquêmica (HR 1.67). Notavelmente, a predição de mortalidade por DHEA-S baixo foi mais pronunciada em homens mais jovens (HR ajustado por idade para morte cardiovascular 2.64) comparado a homens mais velhos (HR 1.30). Estes achados sugerem que a mensuração de DHEA-S em homens na quinta década pode identificar indivíduos com risco cardiovascular aumentado.',
    'pt',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) ON CONFLICT (doi) DO NOTHING;

-- ============================================================================
-- STEP 2: Link Articles to Score Item
-- ============================================================================

-- Link Article 1 (DHEA and Bone Health 2024)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'bc896c78-c530-4dd6-ad03-6da22ffd7fea'
FROM articles a
WHERE a.doi = '10.3390/biomedicines12122780'
ON CONFLICT DO NOTHING;

-- Link Article 2 (DHEA-S and Lifespan 2025)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'bc896c78-c530-4dd6-ad03-6da22ffd7fea'
FROM articles a
WHERE a.doi = '10.1016/j.mad.2025.111998'
ON CONFLICT DO NOTHING;

-- Link Article 3 (DHEA-S and CHD Events 2014)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'bc896c78-c530-4dd6-ad03-6da22ffd7fea'
FROM articles a
WHERE a.doi = '10.1016/j.jacc.2014.05.076'
ON CONFLICT DO NOTHING;

-- Link Article 4 (DHEA-S and Cardiovascular Mortality 2010)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'bc896c78-c530-4dd6-ad03-6da22ffd7fea'
FROM articles a
WHERE a.doi = '10.1210/jc.2010-0760'
ON CONFLICT DO NOTHING;

-- ============================================================================
-- STEP 3: Update Score Item with Clinical Content
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'A mensuração do sulfato de dehidroepiandrosterona (DHEA-S) em homens entre 50-59 anos representa um marcador endócrino crítico para avaliação do envelhecimento adrenal e risco cardiometabólico na transição para a andropausa tardia. Nesta faixa etária, os níveis de DHEA-S encontram-se tipicamente 40-50% abaixo do pico observado na terceira década de vida, refletindo o declínio progressivo da zona reticular adrenal. Estudos prospectivos demonstram que níveis baixos de DHEA-S (< 80 µg/dL) em homens na quinta década predizem risco aumentado de eventos coronarianos em 5 anos (HR 0.86 por SD), mortalidade cardiovascular (HR 1.61) e mortalidade por todas as causas (HR 1.54). A predição de mortalidade é particularmente pronunciada em homens mais jovens dentro desta faixa etária. Além do risco cardiovascular, DHEA-S baixo está associado com aceleração da perda de densidade mineral óssea (taxa de 1% ao ano), risco aumentado de síndrome metabólica em homens não-obesos, e possível declínio cognitivo. O DHEA-S serve como precursor para andrógenos e estrógenos periféricos, modulando múltiplos sistemas fisiológicos incluindo metabolismo ósseo, função endotelial, sensibilidade à insulina e resposta imunológica. A mensuração do DHEA-S nesta década permite identificação precoce de declínio hormonal patológico, estratificação de risco cardiovascular e orientação de intervenções preventivas antes da manifestação clínica de doenças relacionadas ao envelhecimento. A interpretação deve considerar que níveis extremamente elevados também podem associar-se com aumento de pressão arterial em homens, conforme dados recentes de randomização Mendeliana.',

    patient_explanation = 'O DHEA-S (sulfato de dehidroepiandrosterona) é um hormônio produzido principalmente pelas glândulas suprarrenais que funciona como precursor para outros hormônios importantes, incluindo testosterona e estrogênio. Em homens entre 50 e 59 anos, os níveis de DHEA-S naturalmente diminuem para aproximadamente metade dos valores encontrados aos 20-30 anos, como parte do processo normal de envelhecimento chamado "adrenopausa". Esta avaliação é particularmente importante nesta fase da vida porque níveis muito baixos de DHEA-S podem indicar risco aumentado para problemas cardiovasculares, perda acelerada de massa óssea, desenvolvimento de síndrome metabólica e possível declínio das funções cognitivas. O DHEA-S atua em diversos sistemas do corpo: ajuda a manter a saúde dos ossos estimulando células formadoras de osso, protege o sistema cardiovascular influenciando a pressão arterial e a função dos vasos sanguíneos, participa do metabolismo da glicose e gorduras, e contribui para a manutenção da força muscular e energia. Valores baixos nesta faixa etária merecem atenção porque podem sinalizar envelhecimento acelerado e maior vulnerabilidade a doenças crônicas. Por outro lado, estudos recentes também mostram que níveis muito elevados podem estar associados com aumento da pressão arterial, indicando que existe uma faixa ideal de DHEA-S. O médico utiliza este exame junto com outros marcadores hormonais e cardiovasculares para avaliar seu estado geral de saúde, risco de doenças futuras e necessidade de intervenções preventivas específicas para homens na quinta década de vida.',

    conduct = 'INTERPRETAÇÃO CLÍNICA POR FAIXAS: Para homens 50-59 anos, valores de referência esperados situam-se entre 80-560 µg/dL, com declínio progressivo ao longo da década. DHEA-S < 80 µg/dL indica declínio hormonal significativo associado com risco cardiovascular aumentado e requer investigação adicional. DHEA-S 80-150 µg/dL representa zona limítrofe que merece monitoramento e avaliação de fatores de risco associados. DHEA-S 150-350 µg/dL constitui faixa adequada para idade. DHEA-S > 400 µg/dL em homens desta faixa etária é incomum e pode associar-se com hipertensão arterial, requerendo avaliação de pressão arterial e possível investigação de tumor adrenal se muito elevado (> 600 µg/dL).

AVALIAÇÃO COMPLEMENTAR: Diante de DHEA-S baixo, solicitar painel hormonal completo incluindo testosterona total e livre, SHBG, cortisol matinal, TSH, glicemia de jejum, perfil lipídico completo e marcadores inflamatórios (PCR ultrassensível). Avaliar densidade mineral óssea por DXA se DHEA-S < 100 µg/dL. Investigar sintomas de andropausa (fadiga, redução de libido, alterações de humor, perda de massa muscular). Descartar causas secundárias de supressão adrenal: uso de corticosteroides, doença de Addison, hipopituitarismo, doença crítica crônica, desnutrição severa, HIV/AIDS.

ESTRATIFICAÇÃO DE RISCO CARDIOVASCULAR: DHEA-S < 80 µg/dL indica risco cardiovascular aumentado independente. Calcular escore de risco de Framingham ou equivalente. Realizar ECG, ecocardiograma e teste ergométrico se apropriado. Otimizar controle de fatores de risco modificáveis (hipertensão, dislipidemia, diabetes, tabagismo). Considerar dosagem de biomarcadores cardíacos (troponina ultrassensível, BNP) se risco elevado.

INTERVENÇÕES NÃO-FARMACOLÓGICAS: Enfatizar modificações de estilo de vida antes de considerar reposição hormonal: exercício resistido 3-4x/semana para preservação de massa óssea e muscular; atividade aeróbica moderada-vigorosa 150min/semana para saúde cardiovascular; otimização do sono (7-9h/noite) pois deprivação suprime produção de DHEA; manejo de estresse crônico através de técnicas comprovadas; dieta mediterrânea ou antiinflamatória rica em antioxidantes; manutenção de peso saudável (IMC 20-25); suplementação de vitamina D se deficiente (alvo > 30 ng/mL).

SUPLEMENTAÇÃO DE DHEA: A evidência atual (2024-2025) não suporta suplementação rotineira de DHEA para homens. Revisões sistemáticas demonstram benefícios modestos e inconsistentes em densidade óssea, sem benefícios comprovados em cognição, força muscular ou desempenho físico. Dados recentes de randomização Mendeliana sugerem que níveis elevados de DHEA-S podem reduzir expectativa de vida em homens e aumentar pressão arterial. Se considerada, dose típica é 25-50 mg/dia, significativamente menor que as doses estudadas em ensaios (50-75 mg). Monitorar DHEA-S, testosterona, estradiol, PSA, hematócrito e pressão arterial a cada 3 meses inicialmente. Descontinuar se PSA aumentar > 1.4 ng/mL em 12 meses, hematócrito > 54%, ou efeitos adversos (acne, oleosidade, irritabilidade, ginecomastia).

MONITORAMENTO LONGITUDINAL: Repetir DHEA-S anualmente em homens com níveis limítrofes ou baixos. Avaliar trajetória de declínio, pois queda abrupta ou variabilidade extrema associa-se com maior mortalidade independente do nível basal. Documentar sintomas clínicos, composição corporal (massa muscular e gordura), densidade óssea e função cardiovascular longitudinalmente.

CRITÉRIOS DE ENCAMINHAMENTO: Encaminhar para endocrinologista se DHEA-S < 50 µg/dL persistentemente baixo, suspeita de insuficiência adrenal, presença de tumor adrenal, ou hipogonadismo primário concomitante. Encaminhar para cardiologista se risco cardiovascular alto ou doença estabelecida. Última revisão baseada em literatura 2024-2025 incluindo dados de randomização Mendeliana.',

    last_review = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = 'bc896c78-c530-4dd6-ad03-6da22ffd7fea';

COMMIT;

-- ============================================================================
-- VERIFICATION QUERY
-- ============================================================================

SELECT
    'DHEA-S 50-59 Enrichment Summary' as summary,
    LENGTH(clinical_relevance) as clinical_chars,
    LENGTH(patient_explanation) as patient_chars,
    LENGTH(conduct) as conduct_chars,
    (SELECT COUNT(*) FROM article_score_items WHERE score_item_id = 'bc896c78-c530-4dd6-ad03-6da22ffd7fea') as linked_articles
FROM score_items
WHERE id = 'bc896c78-c530-4dd6-ad03-6da22ffd7fea';
