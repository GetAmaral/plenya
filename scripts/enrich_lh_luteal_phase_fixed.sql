-- Enrichment Script for LH - Mulheres Fase Lútea
-- Score Item ID: 31436e15-994d-4a39-b1c5-ae6cf494dacb
-- Generated: 2026-01-29
-- Articles: 4 peer-reviewed publications (2015-2023)

BEGIN;

-- ============================================================================
-- STEP 1: Insert Scientific Articles
-- ============================================================================

-- Article 1: The inadequate corpus luteum (Duncan, 2021)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type,
    abstract
) VALUES (
    gen_random_uuid(),
    'The inadequate corpus luteum',
    'Duncan WC',
    'Reproduction and Fertility',
    '2021-02-01'::date,
    '35128435',
    'review',
    'Estudo revisional sobre a função do corpo lúteo demonstrando que a produção de progesterona lútea é absolutamente dependente da estimulação do receptor de LH/hCG. O LH é essencial em três momentos críticos: formação do corpo lúteo durante a luteinização folicular, manutenção da produção de progesterona durante a fase lútea, e suporte da gravidez inicial até que o hCG placentário assuma esta função. Identificou que corpos lúteos inadequados em ciclos naturais refletem desenvolvimento oocitário subótimo ao invés de defeito intrínseco do corpo lúteo.'
);

-- Article 2: Anatomy of Corpus Luteum (Oliver & Pillarisetty, 2023)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type,
    abstract
) VALUES (
    gen_random_uuid(),
    'Anatomy, Abdomen and Pelvis, Ovary Corpus Luteum',
    'Oliver R, Pillarisetty LS',
    'StatPearls',
    '2023-01-01'::date,
    '30969526',
    'review',
    'Revisão anatômica e fisiológica do corpo lúteo demonstrando que esta estrutura endócrina temporária desenvolve-se após ovulação a partir de células da teca e granulosa foliculares. O corpo lúteo regula o eixo hipotálamo-hipofisário através da inibição de GnRH, diminuindo consequentemente a liberação de LH e FSH pela hipófise anterior. Persiste por aproximadamente 14 dias pós-ovulação se não houver fertilização, degenerando em corpus albicans. Secreta progesterona, inibina A e estradiol durante sua fase funcional.'
);

-- Article 3: Progesterone and the Luteal Phase (Mesen & Young, 2015)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type,
    abstract
) VALUES (
    gen_random_uuid(),
    'Progesterone and the Luteal Phase: A Requisite to Reproduction',
    'Mesen TB, Young SL',
    'Obstetrics and Gynecology Clinics of North America',
    '2015-03-01'::date,
    '25681845',
    'review',
    'Estudo abrangente sobre progesterona e fisiologia da fase lútea. Demonstra que o corpo lúteo desenvolve-se com neovascularização imediata resultando em fluxo sanguíneo excepcionalmente elevado. Células lúteas diferenciam-se em dois tipos morfológicos complementares: células pequenas contendo receptores de LH regulando captação de colesterol, e células grandes com maior capacidade esteroidogênica mas sem receptores de LH, conectadas por gap junctions para síntese coordenada de progesterona. A fase lútea normal dura 11-17 dias (média 14,2 dias).'
);

-- Article 4: Diagnosis and treatment of luteal phase deficiency (ASRM, 2021)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type,
    abstract
) VALUES (
    gen_random_uuid(),
    'Diagnosis and treatment of luteal phase deficiency: a committee opinion',
    'Practice Committee of the American Society for Reproductive Medicine',
    'Fertility and Sterility',
    '2021-05-01'::date,
    '33827766',
    'review',
    'Opinião de comitê da ASRM estabelecendo que progesterona é secretada em pulsos sob controle do LH, com produção pulsátil pelo corpo lúteo em resposta aos pulsos de LH. Os pulsos de progesterona são mais pronunciados nas fases média e tardia da fase lútea, podendo flutuar até 8 vezes dentro de 90 minutos. Valores de progesterona oscilam entre 5 e 40 ng/mL em curtos períodos em mulheres ovulatórias normais. Nenhum teste diagnóstico para deficiência de fase lútea provou ser confiável em diferenciar mulheres férteis de inférteis.'
);

-- ============================================================================
-- STEP 2: Link Articles to Score Item
-- ============================================================================

-- Link Article 1 (Duncan 2021)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '31436e15-994d-4a39-b1c5-ae6cf494dacb'::uuid
FROM articles a
WHERE a.pm_id = '35128435'
AND NOT EXISTS (
    SELECT 1 FROM article_score_items asi
    WHERE asi.article_id = a.id AND asi.score_item_id = '31436e15-994d-4a39-b1c5-ae6cf494dacb'::uuid
);

-- Link Article 2 (Oliver & Pillarisetty 2023)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '31436e15-994d-4a39-b1c5-ae6cf494dacb'::uuid
FROM articles a
WHERE a.pm_id = '30969526'
AND NOT EXISTS (
    SELECT 1 FROM article_score_items asi
    WHERE asi.article_id = a.id AND asi.score_item_id = '31436e15-994d-4a39-b1c5-ae6cf494dacb'::uuid
);

-- Link Article 3 (Mesen & Young 2015)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '31436e15-994d-4a39-b1c5-ae6cf494dacb'::uuid
FROM articles a
WHERE a.pm_id = '25681845'
AND NOT EXISTS (
    SELECT 1 FROM article_score_items asi
    WHERE asi.article_id = a.id AND asi.score_item_id = '31436e15-994d-4a39-b1c5-ae6cf494dacb'::uuid
);

-- Link Article 4 (ASRM 2021)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '31436e15-994d-4a39-b1c5-ae6cf494dacb'::uuid
FROM articles a
WHERE a.pm_id = '33827766'
AND NOT EXISTS (
    SELECT 1 FROM article_score_items asi
    WHERE asi.article_id = a.id AND asi.score_item_id = '31436e15-994d-4a39-b1c5-ae6cf494dacb'::uuid
);

-- ============================================================================
-- STEP 3: Update Score Item with Clinical Content
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'O hormônio luteinizante (LH) na fase lútea desempenha papel essencial na manutenção do corpo lúteo e produção de progesterona necessária para receptividade endometrial e estabelecimento da gravidez. Durante a fase lútea do ciclo menstrual, que dura aproximadamente 11-17 dias (média 14 dias), o LH é secretado de forma pulsátil pela hipófise anterior e mantém a função esteroidogênica do corpo lúteo através da estimulação de receptores LH/hCG nas células luteínicas. A produção de progesterona pelo corpo lúteo é absolutamente dependente desta estimulação por LH, sendo que a remoção ou insuficiência de LH causa declínio acentuado nas concentrações de progesterona. Valores de referência para LH na fase lútea variam de 1,0 a 15,0 mIU/mL, com secreção pulsátil característica. A avaliação de LH lúteo é particularmente importante na investigação de deficiência de fase lútea (DFL), condição caracterizada por fase lútea menor que 11 dias e/ou progesterona sérica meio-lútea inferior a 10 ng/mL. Estudos demonstram que mulheres com DFL apresentam frequência significativamente maior de pulsos de LH comparadas a controles normais (10,5 versus 5,2 pulsos/8 horas), além de pico de LH inadequado ou prematuro na ovulação. O LH é crítico em três momentos: formação do corpo lúteo durante luteinização folicular pós-ovulatória, manutenção da função lútea durante toda a fase lútea, e suporte da gravidez inicial até que hCG embrionário assuma esta função por volta de 8-9 semanas gestacionais. Níveis inadequados de LH lúteo podem refletir disfunção hipotalâmica ou comprometimento do desenvolvimento folicular pré-ovulatório, sendo marcador tanto de função ovariana quanto de qualidade oocitária.',

    patient_explanation = 'O LH (hormônio luteinizante) na fase lútea é o hormônio responsável por manter funcionando o corpo lúteo, uma glândula temporária que se forma no ovário após a ovulação. A fase lútea é o período entre a ovulação e a próxima menstruação, durando normalmente 12-14 dias. Durante esta fase, o LH estimula o corpo lúteo a produzir progesterona, hormônio fundamental para preparar o útero para receber um possível embrião. Valores normais de LH nesta fase ficam entre 1 e 15 mIU/mL. Diferente do pico muito alto de LH que acontece na ovulação (podendo atingir 95 mIU/mL), na fase lútea os níveis são mais baixos mas constantemente presentes, sendo liberados em pequenos pulsos ao longo do dia. Quando os níveis de LH estão baixos ou inadequados nesta fase, o corpo lúteo não consegue produzir progesterona suficiente, levando à condição chamada "deficiência de fase lútea". Esta condição pode causar dificuldade para engravidar ou manter a gravidez inicial, além de ciclos menstruais irregulares e fase lútea encurtada (menor que 11 dias). Se você engravidar, o LH mantém o corpo lúteo funcionando nas primeiras 8-9 semanas até que a placenta assuma a produção de progesterona. A dosagem de LH na fase lútea é especialmente útil em casos de infertilidade, abortos recorrentes, ciclos irregulares ou quando há suspeita de problemas na ovulação. Valores alterados podem indicar necessidade de investigação hormonal mais ampla incluindo função hipofisária e ovariana.',

    conduct = 'AVALIAÇÃO LABORATORIAL: A dosagem de LH na fase lútea deve ser realizada idealmente entre 5-9 dias após ovulação confirmada (fase meio-lútea), correspondendo a 19-23 dias em ciclo regular de 28 dias. Devido à secreção pulsátil característica do LH, com flutuações que podem atingir variações significativas em períodos de 90 minutos, a interpretação de valores isolados requer cautela. Em investigações de deficiência de fase lútea, considerar coletas seriadas (2-3 amostras com intervalo de 30-60 minutos) ou dosagem concomitante de progesterona sérica para melhor caracterização funcional do corpo lúteo. Valores de referência: 1,0-15,0 mIU/mL (fase lútea). INTERPRETAÇÃO DE RESULTADOS: LH lúteo persistentemente elevado (>15 mIU/mL) pode indicar: insuficiência ovariana primária, síndrome dos ovários policísticos com anovulação, ou hiperprolactinemia com ciclos inadequados. LH lúteo baixo (<1,0 mIU/mL) sugere: hipopituitarismo, disfunção hipotalâmica (estresse, exercício excessivo, distúrbios alimentares), uso de agonistas de GnRH, ou defeito intrínseco na pulsatilidade de GnRH. Padrões anormais de pulsatilidade (frequência aumentada de pulsos) são característicos de deficiência de fase lútea mesmo com valores médios dentro da normalidade. CORRELAÇÃO CLÍNICA: Avaliar sempre em conjunto com progesterona meio-lútea (valor adequado >10 ng/mL), estradiol, FSH, prolactina e TSH. História menstrual detalhada incluindo duração da fase lútea (calcular do pico de LH urinário ou temperatura basal até menstruação) é fundamental - fase lútea <11 dias é anormal. Investigar sintomas de tensão pré-menstrual, fertilidade, regularidade dos ciclos e confirmação de ovulação através de ultrassonografia seriada ou pico de LH urinário. CONDUTAS CLÍNICAS: Em casos de deficiência de fase lútea confirmada, a evidência científica atual não suporta suplementação de progesterona em ciclos naturais para melhora de fertilidade. O foco terapêutico deve dirigir-se à otimização do desenvolvimento folicular e qualidade do pico ovulatório, já que corpo lúteo inadequado frequentemente reflete oocitogênese subótima. Considerar: citrato de clomifeno ou letrozol para melhora da qualidade folicular, correção de fatores hipotalâmicos (redução de estresse, normalização de peso, adequação de exercício físico), tratamento de hiperprolactinemia ou disfunção tireoidiana se presentes. Em ciclos de reprodução assistida com múltiplos folículos, a suplementação de progesterona é necessária devido à inibição do LH endógeno por feedback negativo dos níveis suprafisiológicos de esteroides. SEGUIMENTO: Pacientes com DFL confirmada em investigação de infertilidade devem ser referenciadas para especialista em reprodução humana. Monitorização da resposta ao tratamento através de mapeamento folicular ultrassonográfico, dosagens hormonais seriadas e documentação da duração da fase lútea em ciclos subsequentes.',

    last_review = CURRENT_DATE,
    updated_at = CURRENT_TIMESTAMP
WHERE id = '31436e15-994d-4a39-b1c5-ae6cf494dacb';

COMMIT;

-- ============================================================================
-- VERIFICATION QUERIES (Run separately after commit)
-- ============================================================================

-- Verify articles were inserted
SELECT
    pm_id,
    LEFT(title, 60) as title,
    authors,
    journal,
    publish_date
FROM articles
WHERE pm_id IN ('35128435', '30969526', '25681845', '33827766')
ORDER BY publish_date DESC;

-- Verify article linkages
SELECT
    a.pm_id,
    LEFT(a.title, 50) as article_title,
    si.name as score_item
FROM article_score_items asi
JOIN articles a ON asi.article_id = a.id
JOIN score_items si ON asi.score_item_id = si.id
WHERE si.id = '31436e15-994d-4a39-b1c5-ae6cf494dacb'
ORDER BY a.publish_date DESC;

-- Verify character counts of clinical content
SELECT
    name,
    LENGTH(clinical_relevance) as clinical_relevance_chars,
    LENGTH(patient_explanation) as patient_explanation_chars,
    LENGTH(conduct) as conduct_chars,
    last_review
FROM score_items
WHERE id = '31436e15-994d-4a39-b1c5-ae6cf494dacb';

-- Verify complete enrichment
SELECT
    si.name,
    si.code,
    COUNT(DISTINCT a.id) as num_articles,
    CASE
        WHEN si.clinical_relevance IS NOT NULL AND LENGTH(si.clinical_relevance) > 1000 THEN '✓'
        ELSE '✗'
    END as has_clinical_relevance,
    CASE
        WHEN si.patient_explanation IS NOT NULL AND LENGTH(si.patient_explanation) > 800 THEN '✓'
        ELSE '✗'
    END as has_patient_explanation,
    CASE
        WHEN si.conduct IS NOT NULL AND LENGTH(si.conduct) > 1000 THEN '✓'
        ELSE '✗'
    END as has_conduct
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
LEFT JOIN articles a ON asi.article_id = a.id
WHERE si.id = '31436e15-994d-4a39-b1c5-ae6cf494dacb'
GROUP BY si.id, si.name, si.code, si.clinical_relevance, si.patient_explanation, si.conduct;
