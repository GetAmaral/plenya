-- ================================================================
-- ENRICHMENT SCRIPT: IGF-1 (20-30 anos)
-- Score Item ID: cbd75469-ca59-4b12-ab18-9b6b202f54fc
-- Generated: 2026-01-29
-- ================================================================

BEGIN;

-- ================================================================
-- STEP 1: INSERT ARTICLES
-- ================================================================

-- Article 1: Serbian Population Reference Values (2021)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    doi,
    article_type,
    abstract,
    original_link,
    language,
    created_at,
    updated_at
) VALUES (
    'Serum Insulin-Like Growth Factor-1 (IGF-1) Age-Specific Reference Values for Healthy Adult Population of Serbia',
    'Stojanovic M, Popevic M, Pekic S, Doknic M, Miljic D, Medic-Stojanoska M, Topalov D, Stojanovic J, Milovanovic A, Petakov M, Damjanovic S, Popovic V',
    'Acta Endocrinologica',
    '2021-12-01',
    '10.4183/aeb.2021.462',
    'research_article',
    'Este estudo estabeleceu valores de referência populacionais específicos para IGF-1 em adultos sérvios utilizando 1.200 participantes saudáveis (idades 21-80 anos) com representação equilibrada de gênero em doze intervalos de idade de cinco anos. Os pesquisadores utilizaram o ensaio Siemens Immulite 2000 sob condições laboratoriais uniformes e calcularam intervalos de referência usando os percentis 5 e 95 para cada faixa etária. Os resultados demonstraram declínio significativo do IGF-1 relacionado à idade, particularmente pronunciado entre 21-50 anos, seguido por platô até os 70 anos. Diferenças de gênero foram mínimas no geral, embora mulheres tenham exibido padrões de declínio mais uniformes. O método de transformação matemática LMS provou-se superior para calcular escores de desvio padronizado, facilitando melhor aplicação clínica das diretrizes de consenso para deficiência de hormônio do crescimento e manejo de acromegalia.',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC9206165/',
    'pt',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE
SET updated_at = NOW()
RETURNING id;

-- Store the article ID
DO $$
DECLARE
    article_1_id uuid;
    article_2_id uuid;
    article_3_id uuid;
    article_4_id uuid;
BEGIN
    -- Get Article 1 ID
    SELECT id INTO article_1_id FROM articles WHERE doi = '10.4183/aeb.2021.462';

    -- Article 2: IGF-1 Maintenance and Inflammation (2025)
    INSERT INTO articles (
        title,
        authors,
        journal,
        publish_date,
        doi,
        article_type,
        abstract,
        original_link,
        language,
        created_at,
        updated_at
    ) VALUES (
        'Long-Term IGF-1 Maintenance in the Upper-Normal Range Has Beneficial Effect on Low-Grade Inflammation Marker in Adults with Growth Hormone Deficiency',
        'Klinc A, Janež A, Jensterle M',
        'International Journal of Molecular Sciences',
        '2025-01-01',
        '10.3390/ijms26052010',
        'research_article',
        'Esta investigação transversal examinou 31 adultos com deficiência de hormônio do crescimento recebendo terapia de reposição de longo prazo para comparar desfechos entre aqueles mantendo níveis de IGF-1 no intervalo superior-normal versus inferior-normal. Os pesquisadores descobriram que a manutenção de longo prazo de escores de desvio padrão de IGF-1 no intervalo superior-normal foi associada a níveis mais baixos de proteína C-reativa de alta sensibilidade, indicando redução da inflamação sistêmica. Pacientes masculinos alcançaram mais frequentemente níveis superiores-normais, enquanto participantes femininas permaneceram predominantemente em intervalos inferiores-normais, provavelmente devido aos efeitos moduladores do estrogênio sobre a sensibilidade ao hormônio do crescimento. Melhorias na composição corporal correlacionaram-se com manutenção de IGF-1 mais elevado, embora esta associação tenha se tornado não significativa após ajuste por sexo. Os achados sugerem que buscar níveis de IGF-1 no intervalo superior-normal pode fornecer benefícios anti-inflamatórios.',
        'https://pmc.ncbi.nlm.nih.gov/articles/PMC11900236/',
        'pt',
        NOW(),
        NOW()
    ) ON CONFLICT (doi) DO UPDATE
    SET updated_at = NOW()
    RETURNING id INTO article_2_id;

    -- Get Article 2 ID
    SELECT id INTO article_2_id FROM articles WHERE doi = '10.3390/ijms26052010';

    -- Article 3: 2024 Update on GH Deficiency Guidelines
    INSERT INTO articles (
        title,
        authors,
        journal,
        publish_date,
        doi,
        article_type,
        abstract,
        original_link,
        language,
        created_at,
        updated_at
    ) VALUES (
        'A 2024 Update on Growth Hormone Deficiency Syndrome in Adults: From Guidelines to Real Life',
        'Multiple Authors',
        'Journal of Clinical Medicine',
        '2024-10-15',
        '10.3390/jcm13206079',
        'review',
        'Esta revisão abrangente de 2024 atualiza as diretrizes clínicas para diagnóstico e manejo da deficiência de hormônio do crescimento em adultos. O artigo enfatiza que pacientes com três ou mais deficiências hormonais hipofisárias e níveis de IGF-1 abaixo de -2 desvios-padrão têm probabilidade superior a 97% de ter deficiência de GH confirmada. Para pacientes com menos de dois déficits hormonais, níveis baixos de IGF-1 isolados não são suficientes para diagnóstico e testes de estímulo de GH devem ser realizados. O documento revisa valores de referência idade e sexo-específicos para IGF-1, enfatizando a importância de intervalos de referência populacionais adequados. A revisão também discute a interpretação de IGF-1 no contexto de condições clínicas como desnutrição, diabetes mellitus mal controlado, doenças crônicas, insuficiência renal e cirrose hepática, que podem reduzir os níveis de IGF-1 independentemente do status de GH.',
        'https://www.mdpi.com/2077-0383/13/20/6079',
        'pt',
        NOW(),
        NOW()
    ) ON CONFLICT (doi) DO UPDATE
    SET updated_at = NOW()
    RETURNING id INTO article_3_id;

    -- Get Article 3 ID
    SELECT id INTO article_3_id FROM articles WHERE doi = '10.3390/jcm13206079';

    -- Article 4: Normative Data for Indian Adult Males (2025)
    INSERT INTO articles (
        title,
        authors,
        journal,
        publish_date,
        doi,
        article_type,
        abstract,
        original_link,
        language,
        created_at,
        updated_at
    ) VALUES (
        'Normative data for insulin-like growth factor-1 and insulin-like growth factor binding protein-3 and their determinants in healthy adult males from India: the INDIIGo study',
        'Multiple Authors',
        'Scientific Reports',
        '2025-01-09',
        '10.1038/s41598-025-25235-6',
        'research_article',
        'O estudo INDIIGo estabeleceu dados normativos para IGF-1 e IGFBP-3 séricos em 1.271 homens adultos saudáveis indianos com idades entre 20 e 65,9 anos, utilizando o imunoensaio de eletroquimioluminescência Roche Elecsys. Os resultados demonstraram que os níveis séricos de IGF-1 diminuíram 30,1 ng/mL por década de idade, com declínio mais pronunciado nas faixas etárias mais jovens. Para adultos jovens na faixa de 20-30 anos, o estudo forneceu intervalos de referência específicos considerando variações étnicas e metodológicas. O estudo também investigou determinantes dos níveis de IGF-1, incluindo índice de massa corporal, circunferência da cintura e outros parâmetros antropométricos. Esta pesquisa preenche uma lacuna importante ao fornecer dados normativos específicos para população do sul da Ásia, reconhecendo que valores de referência podem variar significativamente entre diferentes populações e métodos de ensaio.',
        'https://www.nature.com/articles/s41598-025-25235-6',
        'pt',
        NOW(),
        NOW()
    ) ON CONFLICT (doi) DO UPDATE
    SET updated_at = NOW()
    RETURNING id INTO article_4_id;

    -- Get Article 4 ID
    SELECT id INTO article_4_id FROM articles WHERE doi = '10.1038/s41598-025-25235-6';

    -- ================================================================
    -- STEP 2: LINK ARTICLES TO SCORE ITEM
    -- ================================================================

    INSERT INTO article_score_items (score_item_id, article_id)
    VALUES
        ('cbd75469-ca59-4b12-ab18-9b6b202f54fc', article_1_id),
        ('cbd75469-ca59-4b12-ab18-9b6b202f54fc', article_2_id),
        ('cbd75469-ca59-4b12-ab18-9b6b202f54fc', article_3_id),
        ('cbd75469-ca59-4b12-ab18-9b6b202f54fc', article_4_id)
    ON CONFLICT (score_item_id, article_id) DO NOTHING;

END $$;

-- ================================================================
-- STEP 3: UPDATE SCORE ITEM WITH CLINICAL CONTENT
-- ================================================================

UPDATE score_items
SET
    clinical_relevance = 'O IGF-1 (Fator de Crescimento Insulina-Símile tipo 1) é o principal mediador dos efeitos anabólicos do hormônio do crescimento (GH) e um biomarcador essencial para avaliação do eixo GH-IGF-1 em adultos jovens de 20-30 anos. Nesta faixa etária, o IGF-1 encontra-se próximo aos níveis de pico pós-puberal, com valores de referência tipicamente entre 150-450 ng/mL dependendo do ensaio utilizado, gênero e população específica. Estudos populacionais demonstram que o declínio do IGF-1 inicia-se logo após os 20 anos, com redução média de 30 ng/mL por década. A dosagem de IGF-1 possui vantagens sobre a dosagem direta de GH devido à sua estabilidade ao longo do dia, sem variações circadianas ou pulsos secretórios. Clinicamente, níveis de IGF-1 abaixo do percentil 2,5 (escore Z < -2) em adultos jovens sugerem deficiência de GH ou resistência severa ao GH, especialmente quando associados a múltiplas deficiências hipofisárias (probabilidade >97% de deficiência de GH se três ou mais hormônios hipofisários deficientes). Por outro lado, níveis elevados de IGF-1 indicam excesso de GH, geralmente secundário a adenomas hipofisários secretores de GH (acromegalia). Na faixa etária de 20-30 anos, a interpretação do IGF-1 deve considerar fatores que podem reduzir seus níveis independentemente do status de GH: desnutrição proteico-calórica, diabetes mellitus mal controlado, doenças crônicas, insuficiência renal e cirrose hepática. O IGF-1 também apresenta dimorfismo sexual sutil, com mulheres jovens tendendo a níveis discretamente mais baixos devido aos efeitos moduladores do estrogênio sobre a sensibilidade hepática ao GH. Para pacientes em reposição de GH, o objetivo terapêutico é manter o IGF-1 no terço médio-superior do intervalo de referência, pois estudos recentes demonstram que manutenção no intervalo superior-normal associa-se a menores marcadores inflamatórios (PCR-ultrassensível) e melhor composição corporal.',

    patient_explanation = 'O IGF-1 (Fator de Crescimento Insulina-Símile 1) é uma proteína produzida principalmente pelo fígado em resposta ao hormônio do crescimento (GH) liberado pela glândula hipófise no cérebro. Embora seja mais conhecido por seu papel no crescimento durante a infância e adolescência, o IGF-1 continua tendo funções importantes na vida adulta, especialmente entre 20-30 anos quando seus níveis ainda estão relativamente elevados. Na prática, o IGF-1 funciona como um "mensageiro" que leva os sinais do hormônio do crescimento para todo o corpo, ajudando na manutenção da massa muscular, densidade óssea, metabolismo de gorduras e açúcares, e até na saúde cardiovascular. Diferente do hormônio do crescimento que varia muito ao longo do dia, o IGF-1 permanece estável, tornando-se um exame de sangue muito útil para avaliar se o hormônio do crescimento está funcionando adequadamente. Para pessoas na faixa dos 20-30 anos, valores normais geralmente ficam entre 150 e 450 ng/mL, mas é importante comparar com a tabela específica do laboratório que realizou o exame. Níveis muito baixos de IGF-1 podem indicar deficiência do hormônio do crescimento, causando sintomas como fadiga excessiva, dificuldade em ganhar massa muscular, aumento de gordura corporal e menor densidade óssea. Já níveis muito elevados podem sinalizar produção excessiva de hormônio do crescimento, geralmente devido a um tumor benigno na hipófise, levando a condições como acromegalia. Vários fatores podem afetar os níveis de IGF-1 além do hormônio do crescimento, incluindo nutrição inadequada, diabetes não controlado, doenças crônicas do fígado ou rins, e até mesmo o uso de certos medicamentos. Por isso, a interpretação do resultado sempre deve ser feita por um médico endocrinologista considerando seu histórico clínico completo.',

    conduct = 'A interpretação adequada do IGF-1 em adultos de 20-30 anos requer abordagem sistemática e contextualizada. PRIMEIRA ETAPA - Verificação da adequação do exame: confirmar que o paciente estava em jejum de 8-12 horas, sem uso de biotina nas últimas 72 horas (pode interferir em alguns ensaios), e que o resultado está sendo comparado com valores de referência idade e sexo-específicos do laboratório executor, preferencialmente utilizando escores de desvio padrão (Z-scores ou SDS). SEGUNDA ETAPA - Interpretação de valores normais (entre percentil 2,5 e 97,5, ou SDS entre -2 e +2): indicam função normal do eixo GH-IGF-1 nesta faixa etária. Não requer investigação adicional exceto se sintomas clínicos sugestivos de disfunção do GH estiverem presentes. TERCEIRA ETAPA - Investigação de IGF-1 baixo (< percentil 2,5 ou SDS < -2): antes de confirmar deficiência de GH, excluir causas secundárias de IGF-1 reduzido mediante avaliação de: estado nutricional (albumina, pré-albumina), controle glicêmico (hemoglobina glicada), função hepática (transaminases, bilirrubinas, tempo de protrombina), função renal (creatinina, taxa de filtração glomerular), presença de doenças inflamatórias crônicas (PCR, VHS), e uso de medicações (glicocorticoides, estrogênios orais). Se IGF-1 persistentemente baixo após exclusão de causas secundárias, avaliar presença de outros déficits hormonais hipofisários (TSH/T4 livre, cortisol, LH/FSH, prolactina). Pacientes com três ou mais déficits hipofisários associados E IGF-1 < -2 SDS têm diagnóstico de deficiência de GH com probabilidade >97%, dispensando testes de estímulo. Pacientes com menos de dois déficits hormonais necessitam teste de estímulo de GH (teste de tolerância à insulina, teste do glucagon, ou teste combinado GHRH+arginina) para confirmar deficiência de GH. QUARTA ETAPA - Investigação de IGF-1 elevado (> percentil 97,5 ou SDS > +2): solicitar dosagem de GH basal e teste de supressão oral de glicose (TSOG). No TSOG, administra-se 75g de glicose oral e mede-se GH em 0, 30, 60, 90 e 120 minutos; falha em suprimir GH para <1 ng/mL (ensaios ultrassensíveis) ou <0,4 ng/mL confirma acromegalia. Realizar ressonância magnética de hipófise com contraste para identificar adenoma secretor de GH. Avaliar complicações da acromegalia: ecocardiograma (hipertrofia ventricular), colonoscopia (pólipos), polissonografia (apneia do sono), campo visual, glicemia e TOTG (diabetes secundário). QUINTA ETAPA - Monitoramento de pacientes em reposição de GH: dosar IGF-1 a cada 3-6 meses inicialmente, ajustando dose para manter IGF-1 no terço médio-superior do intervalo de referência (SDS entre 0 e +2). Estudos recentes indicam que manutenção no intervalo superior-normal associa-se a menores níveis de marcadores inflamatórios (PCR-us) e melhor composição corporal. Avaliar periodicamente composição corporal (DEXA), perfil lipídico, glicemia, e sintomas de excesso de GH (artralgia, edema, cefaleia). CONSIDERAÇÕES ESPECIAIS: em mulheres usando estrogênio oral, a biodisponibilidade de GH é reduzida, podendo resultar em IGF-1 mais baixo apesar de dose adequada de GH; considerar trocar para estrogênio transdérmico ou aumentar dose de GH. A interpretação do IGF-1 isoladamente tem limitações; sempre correlacionar com quadro clínico e outros exames complementares.',

    last_review = '2026-01-29',
    updated_at = NOW()
WHERE id = 'cbd75469-ca59-4b12-ab18-9b6b202f54fc';

COMMIT;

-- ================================================================
-- VERIFICATION QUERY
-- ================================================================

SELECT
    si.id,
    si.name,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    si.last_review,
    COUNT(asi.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = 'cbd75469-ca59-4b12-ab18-9b6b202f54fc'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;
