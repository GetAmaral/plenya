-- Enriquecimento do item "Urocultura com contagem de colônias e antibiograma automatizado"
-- Score Item ID: 7ec43763-493a-481b-9865-4f793840ab20
-- Data: 2026-01-29

BEGIN;

-- 1. Atualizar campos clínicos do score_items
UPDATE score_items
SET
    clinical_relevance = 'A urocultura com contagem de colônias e antibiograma automatizado representa o padrão-ouro diagnóstico para infecções do trato urinário (ITU), permitindo identificação precisa do agente etiológico e seu perfil de susceptibilidade antimicrobiana. O limiar clássico de ≥100.000 UFC/mL (10⁵ CFU/mL) tem sido revisado por diretrizes contemporâneas: em pacientes sintomáticos, contagens entre 10.000-100.000 UFC/mL (10⁴-10⁵) podem representar ITU verdadeira, especialmente em amostras obtidas por cateterismo ou punção suprapúbica. A automação por sistemas como VITEK 2 e Phoenix BD reduz o tempo de resposta para 4-6 horas, crucial para iniciar terapia direcionada precocemente. Em 2024, a resistência de E. coli produtora de ESBL (beta-lactamase de espectro estendido) atingiu níveis críticos globalmente, com 65.9% de resistência a cefepima e 56.1% a cefotaxima, tornando o antibiograma indispensável para evitar falha terapêutica. As diretrizes IDSA 2025 recomendam obtenção de cultura antes de iniciar antibioticoterapia em ITU complicada, cistite recorrente e pielonefrite, enfatizando de-escalation baseado em resultados. A interpretação adequada diferencia contaminação (múltiplos organismos, <10⁴ UFC/mL) de infecção verdadeira, reduzindo uso inapropriado de antimicrobianos e custos hospitalares. O antibiograma automatizado detecta resistências emergentes (carbapenemases, fluoroquinolonas) essenciais para vigilância epidemiológica institucional.',
    patient_explanation = 'Este exame é considerado o mais preciso para diagnosticar infecções urinárias, pois identifica exatamente qual bactéria está causando a infecção e quais antibióticos funcionam contra ela. A coleta adequada da urina (geralmente jato médio após higiene) é fundamental para evitar contaminação. O laboratório cultiva sua urina em placas especiais e conta quantas colônias de bactérias crescem: quando há ≥100.000 colônias por mililitro, indica infecção estabelecida, mas valores entre 10.000-100.000 também podem ser significativos se você tem sintomas como ardência ao urinar, dor baixo-ventre ou febre. O "antibiograma" é o teste que mostra quais antibióticos matam a bactéria encontrada - hoje feito por máquinas automatizadas (VITEK, Phoenix) que dão resultado em 4-6 horas, muito mais rápido que métodos antigos. Isso é especialmente importante porque algumas bactérias desenvolveram resistência a antibióticos comuns: por exemplo, a E. coli (causa de 80% das infecções urinárias) pode produzir enzimas ESBL que a tornam resistente a muitos medicamentos, exigindo tratamentos alternativos. Seu médico usa este exame para escolher o antibiótico ideal, evitar medicamentos que não funcionariam e reduzir o risco de complicações. Em infecções de repetição, o exame ajuda a identificar padrões e guiar estratégias preventivas personalizadas.',
    conduct = 'A interpretação clínica da urocultura deve integrar contagem de colônias, características do isolado e contexto clínico. Em cistite não-complicada com sintomas típicos (disúria, urgência, polaciúria) e contagem ≥10⁵ UFC/mL de uropatógeno único (E. coli, Klebsiella, Proteus, Enterococcus), iniciar terapia empírica baseada em antibiograma local enquanto aguarda antibiograma individual. Para contagens 10⁴-10⁵ UFC/mL, considerar ITU se sintomas presentes, descartando contaminação (ausência de células epiteliais escamosas >10/campo em microscopia). Em pielonefrite e ITU complicada, solicitar sempre cultura pré-tratamento e reavaliar terapia após antibiograma, implementando de-escalation quando possível (diretrizes IDSA 2025). Resistência prévia documentada contraindica uso do mesmo antimicrobiano; exposição a fluoroquinolonas nos últimos 12 meses indica evitar esta classe. Para E. coli produtora de ESBL, priorizar nitrofurantoína ou sulfametoxazol-trimetoprima em cistite não-complicada se sensível; reservar carbapenêmicos (ertapenem, meropenem) para pielonefrite ou ITU complicada por ESBL. Culturas polimicrobianas (≥3 organismos) sugerem contaminação, exigindo nova coleta com técnica asséptica rigorosa. Bacteriúria assintomática (≥10⁵ UFC/mL sem sintomas) não deve ser tratada, exceto em gestantes e pré-procedimentos urológicos invasivos. Documentar resultados em prontuário eletrônico, incluindo MIC (concentração inibitória mínima) de antimicrobianos, e notificar perfis MDR (multirresistentes) à CCIH. Repetir cultura 48-72h após início de terapia em pacientes que não melhoram clinicamente. Considerar urologia se ITU recorrente (≥2 episódios/6 meses ou ≥3/ano), sugerindo anormalidades anatômicas ou funcionais. Educar pacientes sobre técnica de coleta, sinais de alerta (febre alta, dor lombar, náuseas) e importância de completar ciclo antibiótico mesmo com melhora precoce dos sintomas.',
    last_review = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = '7ec43763-493a-481b-9865-4f793840ab20';

-- Verificar se a atualização foi bem-sucedida
DO $$
DECLARE
    v_count int;
BEGIN
    SELECT COUNT(*) INTO v_count
    FROM score_items
    WHERE id = '7ec43763-493a-481b-9865-4f793840ab20'
      AND clinical_relevance IS NOT NULL
      AND patient_explanation IS NOT NULL
      AND conduct IS NOT NULL;

    IF v_count = 0 THEN
        RAISE EXCEPTION 'Score item com ID 7ec43763-493a-481b-9865-4f793840ab20 não encontrado ou não foi atualizado';
    END IF;

    RAISE NOTICE 'Score item atualizado com sucesso!';
END $$;

-- 2. Inserir artigos científicos na tabela "articles"

-- Artigo 1: Core Curriculum UTI 2024 (AJKD)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type,
    created_at,
    updated_at
)
SELECT
    gen_random_uuid(),
    'Urinary Tract Infections: Core Curriculum 2024',
    'Hawra Al Lawati, Barbra M Blair, Jeffrey Larnard',
    'American Journal of Kidney Diseases',
    '2024-01-01'::date,
    '37906240',
    'review',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '37906240'
);

-- Artigo 2: VITEK Reveal para AST rápido (2024)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type,
    created_at,
    updated_at
)
SELECT
    gen_random_uuid(),
    'Evaluation of the Vitek® Reveal™ system for rapid antimicrobial susceptibility testing of Gram-negative pathogens, including ESBL, CRE and CRAB, from positive blood cultures',
    'Alberto Antonelli, Sara Cuffari, Benedetta Casciato, Tommaso Giani, Gian Maria Rossolini',
    'Diagnostic Microbiology and Infectious Disease',
    '2024-12-01'::date,
    '39197326',
    'research_article',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '39197326'
);

-- Artigo 3: IDSA Guidance ESBL-E (2022)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type,
    created_at,
    updated_at
)
SELECT
    gen_random_uuid(),
    'Infectious Diseases Society of America 2022 Guidance on the Treatment of Extended-Spectrum β-lactamase Producing Enterobacterales (ESBL-E), Carbapenem-Resistant Enterobacterales (CRE), and Pseudomonas aeruginosa with Difficult-to-Treat Resistance (DTR-P. aeruginosa)',
    'Pranita D Tamma, Paul G Ambrose, David L Paterson, et al.',
    'Clinical Infectious Diseases',
    '2022-04-19'::date,
    '35439291',
    'clinical_trial',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '35439291'
);

-- Artigo 4: Continuum of UTI (2024)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type,
    created_at,
    updated_at
)
SELECT
    gen_random_uuid(),
    'Proposing the Continuum of UTI for a Nuanced Approach to Diagnosis and Management of Urinary Tract Infections',
    'Sonali D Advani, Nicholas A Turner, Rebecca North, Rebekah W Moehring, Valerie M Vaughn, Charles D Scales Jr, Nazema Y Siddiqui, Kenneth E Schmader, Deverick J Anderson',
    'The Journal of Urology',
    '2024-02-08'::date,
    '38330392',
    'review',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '38330392'
);

-- 3. Criar vínculos na tabela de junção "article_score_items"
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '7ec43763-493a-481b-9865-4f793840ab20'::uuid,
    a.id
FROM articles a
WHERE a.pm_id IN ('37906240', '39197326', '35439291', '38330392')
  AND NOT EXISTS (
    SELECT 1 FROM article_score_items asi
    WHERE asi.score_item_id = '7ec43763-493a-481b-9865-4f793840ab20'
      AND asi.article_id = a.id
  );

-- 4. Verificações finais
DO $$
DECLARE
    v_clinical_length int;
    v_patient_length int;
    v_conduct_length int;
    v_article_count int;
BEGIN
    -- Verificar comprimentos dos textos
    SELECT
        length(clinical_relevance),
        length(patient_explanation),
        length(conduct)
    INTO v_clinical_length, v_patient_length, v_conduct_length
    FROM score_items
    WHERE id = '7ec43763-493a-481b-9865-4f793840ab20';

    RAISE NOTICE 'Comprimento clinical_relevance: % caracteres', v_clinical_length;
    RAISE NOTICE 'Comprimento patient_explanation: % caracteres', v_patient_length;
    RAISE NOTICE 'Comprimento conduct: % caracteres', v_conduct_length;

    -- Verificar se está dentro dos limites esperados
    IF v_clinical_length < 1500 OR v_clinical_length > 2000 THEN
        RAISE WARNING 'clinical_relevance fora do range esperado (1500-2000): %', v_clinical_length;
    END IF;

    IF v_patient_length < 1000 OR v_patient_length > 1500 THEN
        RAISE WARNING 'patient_explanation fora do range esperado (1000-1500): %', v_patient_length;
    END IF;

    IF v_conduct_length < 1500 OR v_conduct_length > 2500 THEN
        RAISE WARNING 'conduct fora do range esperado (1500-2500): %', v_conduct_length;
    END IF;

    -- Verificar quantidade de artigos vinculados
    SELECT COUNT(*)
    INTO v_article_count
    FROM article_score_items
    WHERE score_item_id = '7ec43763-493a-481b-9865-4f793840ab20';

    RAISE NOTICE 'Total de artigos vinculados: %', v_article_count;

    IF v_article_count < 4 THEN
        RAISE WARNING 'Menos de 4 artigos vinculados (encontrados: %)', v_article_count;
    END IF;
END $$;

COMMIT;

-- Query para visualizar o resultado final
SELECT
    si.id,
    si.name,
    length(si.clinical_relevance) as clinical_length,
    length(si.patient_explanation) as patient_length,
    length(si.conduct) as conduct_length,
    si.last_review,
    COUNT(asi.article_id) as article_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '7ec43763-493a-481b-9865-4f793840ab20'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;
