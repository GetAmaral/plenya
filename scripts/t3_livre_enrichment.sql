-- Enriquecimento do item T3 Livre
-- ID: d164eacf-a0d7-48f2-899d-3f0d57ec7cc3
-- Data: 2026-01-28

-- Primeiro, vamos inserir os artigos científicos de referência
INSERT INTO articles (
    title, authors, journal, publish_date, doi, pm_id,
    abstract, specialty, language, created_at, updated_at
) VALUES
(
    'The clinical evaluation of patients with subclinical hyperthyroidism and free triiodothyronine (free T3) toxicosis',
    'Figge J, Leinung M, Goodman AD et al.',
    'American Journal of Medicine',
    '1994-01-01',
    NULL,
    '8154510',
    'Estudo clássico sobre a avaliação de pacientes com hipertireoidismo subclínico e T3 toxicosis, estabelecendo critérios diagnósticos e padrões de interpretação para T3 livre em contextos de função tireoidiana limítrofe.',
    'Endocrinologia',
    'en',
    NOW(),
    NOW()
),
(
    'Clinical usage recommendations and analytic performance goals for total and free triiodothyronine measurements',
    'Demers LM, Spencer CA',
    'Clinical Chemistry',
    '1995-01-01',
    NULL,
    '8565219',
    'Recomendações para uso clínico e objetivos de desempenho analítico para medições de T3 total e livre, incluindo padronização de métodos laboratoriais e valores de referência.',
    'Endocrinologia',
    'en',
    NOW(),
    NOW()
),
(
    'To test or not to test? Clinical utility and considerations for triiodothyronine (T3) testing',
    'Association for Diagnostics & Laboratory Medicine',
    'ADLMConnect',
    '2024-01-01',
    NULL,
    NULL,
    'Revisão contemporânea sobre utilidade clínica do teste de T3, incluindo indicações precisas, interpretação em diferentes contextos clínicos e frequência de T3 toxicosis (1.6% em pacientes com TSH suprimido).',
    'Endocrinologia',
    'en',
    NOW(),
    NOW()
)
ON CONFLICT (pm_id) WHERE pm_id IS NOT NULL DO NOTHING;

-- Capturar IDs dos artigos inseridos
DO $$
DECLARE
    article1_id UUID;
    article2_id UUID;
    article3_id UUID;
    item_id UUID := 'd164eacf-a0d7-48f2-899d-3f0d57ec7cc3';
BEGIN
    -- Buscar IDs dos artigos
    SELECT id INTO article1_id FROM articles WHERE pm_id = '8154510';
    SELECT id INTO article2_id FROM articles WHERE pm_id = '8565219';
    SELECT id INTO article3_id FROM articles WHERE title LIKE '%To test or not to test%';

    -- Criar relações many-to-many (se não existirem)
    IF article1_id IS NOT NULL THEN
        INSERT INTO article_score_items (article_id, score_item_id, created_at)
        VALUES (article1_id, item_id, NOW())
        ON CONFLICT DO NOTHING;
    END IF;

    IF article2_id IS NOT NULL THEN
        INSERT INTO article_score_items (article_id, score_item_id, created_at)
        VALUES (article2_id, item_id, NOW())
        ON CONFLICT DO NOTHING;
    END IF;

    IF article3_id IS NOT NULL THEN
        INSERT INTO article_score_items (article_id, score_item_id, created_at)
        VALUES (article3_id, item_id, NOW())
        ON CONFLICT DO NOTHING;
    END IF;
END $$;

-- Atualizar o score_item com conteúdo clínico em PT-BR
UPDATE score_items
SET
    clinical_relevance = 'O T3 livre (triiodotironina livre) é a forma biologicamente ativa do hormônio tireoidiano triiodotironina, representando a fração não ligada a proteínas transportadoras e disponível para ação celular. Enquanto o T4 (tiroxina) é o principal hormônio secretado pela tireoide, o T3 é cerca de 3-4 vezes mais potente e responsável pela maioria dos efeitos biológicos dos hormônios tireoidianos. Aproximadamente 80% do T3 circulante é produzido por conversão periférica do T4 através da enzima 5''-deiodinase, principalmente no fígado e rins.

A dosagem de T3 livre está indicada principalmente na investigação de hipertireoidismo quando o TSH está suprimido mas o T4 livre está normal ou limítrofe, situação conhecida como T3 toxicosis. Esta condição representa uma forma de hipertireoidismo onde apenas o T3 está elevado, ocorrendo em aproximadamente 1,6% dos pacientes com TSH suprimido, sendo mais comum em doença de Graves inicial, bócio multinodular tóxico e adenomas tóxicos. Valores elevados de T3 livre (>4.4 pg/mL) na presença de TSH suprimido (<0.4 mIU/L) confirmam o diagnóstico de T3 toxicosis, indicando necessidade de tratamento mesmo com T4 livre normal.

A interpretação do T3 livre deve sempre considerar o contexto clínico e os demais hormônios tireoidianos. Valores reduzidos (<2.0 pg/mL) podem ocorrer no hipotireoidismo, síndrome do doente eutireoideo (euthyroid sick syndrome) e uso de medicamentos que inibem a conversão periférica de T4 para T3 (propranolol, amiodarona, corticoides, propiltiouracil). O T3 livre é menos útil no diagnóstico de hipotireoidismo, onde TSH e T4 livre são suficientes. A dosagem de T3 livre não deve ser solicitada de rotina em todos os pacientes com alterações tireoidianas, mas reservada para situações específicas onde há suspeita clínica de T3 toxicosis ou necessidade de avaliar a conversão periférica de T4.',

    patient_explanation = 'O T3 livre é um dos hormônios produzidos pela sua glândula tireoide, localizada no pescoço. Ele é a forma mais ativa desse hormônio e ajuda a controlar o metabolismo do seu corpo - ou seja, como seu organismo usa energia. O "livre" no nome significa que é a parte do hormônio que está disponível para ser usada pelas células, não aquela que está presa a outras proteínas no sangue.

Seu médico pode pedir este exame quando suspeita que sua tireoide está trabalhando demais (hipertireoidismo), especialmente se outros exames de tireoide mostraram resultados no limite. Valores aumentados podem causar sintomas como coração acelerado, perda de peso mesmo comendo normalmente, nervosismo, tremores nas mãos, suor excessivo e intolerância ao calor. Valores diminuídos podem ocorrer quando a tireoide está funcionando pouco, causando cansaço, ganho de peso, pele seca e sensação de frio constante.

Os valores normais geralmente ficam entre 2.0 e 4.4 pg/mL, mas podem variar um pouco dependendo do laboratório. É importante que você não se preocupe apenas com um número isolado - seu médico vai interpretar o resultado junto com outros exames de tireoide (TSH e T4 livre) e seus sintomas. Se o resultado estiver alterado, não significa necessariamente que você tem um problema grave, mas que precisa de uma avaliação mais detalhada com seu médico ou um endocrinologista.',

    conduct = 'PROTOCOLO DE INVESTIGAÇÃO QUANDO T3 LIVRE ESTÁ ALTERADO:

T3 livre elevado (>4.4 pg/mL):
• Confirmar com TSH e T4 livre na mesma amostra
• Se TSH suprimido (<0.4 mIU/L) + T4 livre normal + T3 livre elevado = T3 toxicosis confirmada
• Solicitar anticorpos antitireoidianos (anti-TPO, anti-TRAb) para diferenciar etiologia
• Cintilografia de tireoide com captação se indicado (diferenciar Graves de bócio multinodular/adenoma)
• USG de tireoide para avaliar morfologia e presença de nódulos
• Avaliar sinais e sintomas de tireotoxicose: taquicardia, tremor, perda ponderal, irritabilidade
• Considerar beta-bloqueador se sintomas cardiovasculares importantes
• ENCAMINHAR AO ENDOCRINOLOGISTA para manejo definitivo (antitireoidianos, radioiodo ou cirurgia)

T3 livre reduzido (<2.0 pg/mL):
• Verificar contexto: doença aguda grave? (síndrome do doente eutireoideo)
• Revisar medicações: uso de beta-bloqueadores, amiodarona, corticoides, propiltiouracil?
• Se TSH elevado + T4 livre reduzido = hipotireoidismo primário (tratar com levotiroxina)
• Se TSH normal/baixo + T4 livre normal = possível deficiência isolada de conversão T4→T3 ou artefato laboratorial
• Considerar causas de má conversão: deficiência de selênio, doença hepática, insuficiência renal

REAVALIAÇÃO:
• T3 toxicosis em tratamento: repetir TSH, T4 livre e T3 livre a cada 4-6 semanas até normalização
• Pacientes com T3 inicialmente limítrofe: reavaliar em 8-12 semanas com TSH, T4 livre e T3 livre
• Sempre encaminhar ao endocrinologista se dúvida diagnóstica ou necessidade de tratamento específico',

    last_review = NOW(),
    updated_at = NOW()
WHERE id = 'd164eacf-a0d7-48f2-899d-3f0d57ec7cc3';

-- Verificar atualização
SELECT
    id,
    name,
    CASE
        WHEN clinical_relevance IS NOT NULL THEN '✓ Preenchido'
        ELSE '✗ Vazio'
    END as clinical_relevance_status,
    CASE
        WHEN patient_explanation IS NOT NULL THEN '✓ Preenchido'
        ELSE '✗ Vazio'
    END as patient_explanation_status,
    CASE
        WHEN conduct IS NOT NULL THEN '✓ Preenchido'
        ELSE '✗ Vazio'
    END as conduct_status,
    last_review
FROM score_items
WHERE id = 'd164eacf-a0d7-48f2-899d-3f0d57ec7cc3';

-- Verificar artigos vinculados
SELECT
    a.id,
    a.title,
    a.authors,
    a.journal,
    EXTRACT(YEAR FROM a.publish_date) as publication_year,
    a.pm_id
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = 'd164eacf-a0d7-48f2-899d-3f0d57ec7cc3'
ORDER BY a.publish_date DESC;
