-- Enrichment for DHEA-S - Homens (20-29 anos)
-- Score Item ID: a6742c97-a610-4bd4-b9de-87e91ecc8d34
-- Generated: 2026-01-28

BEGIN;

-- Insert scientific articles (ON CONFLICT DO NOTHING to handle duplicates)
INSERT INTO articles (id, title, authors, journal, publish_date, doi, pm_id, article_type, language, specialty, created_at, updated_at)
VALUES
(
    gen_random_uuid(),
    'Reference Ranges for Serum Dehydroepiandrosterone Sulfate and Testosterone in Adult Men',
    'Friedrich N, Alte D, Völzke H, Spilcke-Liss E, Lüdemann J, Lerch MM, Kohlmann T, Nauck M, Wallaschofski H',
    'Journal of Andrology',
    '2008-11-01'::date,
    '10.2164/jandrol.108.005561',
    '18599883',
    'research_article',
    'en',
    'Endocrinology',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    gen_random_uuid(),
    'Performance of Dehydroepiandrosterone Sulfate and Baseline Cortisol in Assessing Adrenal Insufficiency',
    'Han AJ, Suresh M, Gruber LM, Algeciras-Schimnich A, Achenbach SJ, Atkinson EJ, Bancos I',
    'Journal of Clinical Endocrinology & Metabolism',
    '2024-08-07'::date,
    '10.1210/clinem/dgae855',
    '39657727',
    'research_article',
    'en',
    'Endocrinology',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    gen_random_uuid(),
    'Adrenal Androgens and Aging',
    'Papadopoulou-Marketou N, Kassi E, Chrousos GP',
    'Endotext',
    '2023-01-18'::date,
    NULL,
    '25905237',
    'review',
    'en',
    'Endocrinology',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    gen_random_uuid(),
    'A Review of Age-Related Dehydroepiandrosterone Decline and Its Association with Well-Known Geriatric Syndromes: Is Treatment Beneficial?',
    'Samaras N, Samaras D, Frangos E, Forster A, Philippe J',
    'Rejuvenation Research',
    '2013-08-01'::date,
    '10.1089/rej.2013.1425',
    '23647054',
    'review',
    'en',
    'Geriatrics',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
ON CONFLICT (doi) DO NOTHING;

-- Get the article IDs for linking
WITH article_ids AS (
    SELECT id FROM articles
    WHERE doi IN (
        '10.2164/jandrol.108.005561',
        '10.1210/clinem/dgae855',
        '10.1089/rej.2013.1425'
    )
    OR pm_id IN ('25905237')
)
-- Link articles to score item (ON CONFLICT DO NOTHING to handle duplicates)
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    'a6742c97-a610-4bd4-b9de-87e91ecc8d34'::uuid,
    id
FROM article_ids
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- Update score item with enriched clinical content
UPDATE score_items
SET
    clinical_relevance = 'O DHEA-S (sulfato de dehidroepiandrosterona) é o esteróide adrenal mais abundante no organismo humano, produzido exclusivamente pela zona reticular do córtex adrenal em resposta ao ACTH. Em homens de 20-29 anos, os níveis de DHEA-S estão no pico fisiológico, com valores médios de 286 μg/dL (faixa: 226-361 μg/dL) aos 20-24 anos, representando a fase de máxima produção adrenal. O DHEA-S é considerado o marcador ideal da função da zona reticular adrenal devido à sua longa meia-vida (7-10 horas) e ausência de variação circadiana, ao contrário do DHEA não-sulfatado. Funciona como pró-hormônio essencial, sendo convertido perifericamente em andrógenos ativos (testosterona, diidrotestosterona) e estrógenos, contribuindo com cerca de 30% dos andrógenos totais em homens. Níveis elevados podem indicar tumores adrenais, hiperplasia adrenal congênita ou síndrome de Cushing. Níveis baixos nesta faixa etária (abaixo de 110 μg/dL) são altamente sugestivos de insuficiência adrenal, especialmente quando combinados com cortisol basal reduzido. A combinação de DHEA-S e cortisol basal possui AUROC de 0.81 para diagnóstico de insuficiência adrenal, podendo eliminar necessidade de testes dinâmicos em muitos casos. Após o pico na terceira década, ocorre declínio fisiológico de 2-5% ao ano, com variabilidade individual significativa (fatores genéticos explicam 37-46% da variância). A avaliação do DHEA-S nesta faixa etária estabelece baseline individual para monitoramento futuro e identifica precocemente disfunções adrenais.',

    patient_explanation = 'O DHEA-S (sulfato de dehidroepiandrosterona) é um hormônio produzido pelas glândulas adrenais (localizadas acima dos rins) que funciona como "matéria-prima" para fabricar outros hormônios importantes, incluindo testosterona e estrogênio. Entre 20 e 29 anos, você está na fase da vida em que seu corpo produz a maior quantidade deste hormônio – é o seu "pico de produção". Os valores normais nesta idade ficam entre 110 a 510 μg/dL, com média em torno de 286 μg/dL. Ter níveis adequados de DHEA-S é importante porque ele reflete se suas glândulas adrenais estão funcionando bem. Níveis muito baixos podem indicar problemas nas glândulas adrenais (como fadiga adrenal ou insuficiência adrenal), enquanto níveis muito altos podem sugerir tumores ou outras condições que precisam investigação. Diferente de outros hormônios que variam ao longo do dia, o DHEA-S permanece estável, o que torna o exame confiável independente do horário da coleta. Este hormônio também está relacionado à energia, disposição, massa muscular, saúde óssea e função sexual. Conhecer seu nível de DHEA-S agora, no auge da produção, cria uma "fotografia" importante para comparações futuras, já que após os 30 anos os níveis naturalmente começam a diminuir cerca de 2-5% a cada ano. Manter o acompanhamento ajuda a identificar precocemente qualquer problema e permite intervenções quando necessário.',

    conduct = 'PROTOCOLO DE INTERPRETAÇÃO E CONDUTA PARA DHEA-S EM HOMENS 20-29 ANOS:

FAIXAS DE REFERÊNCIA ESPECÍFICAS POR IDADE:
• 20-24 anos: 226-361 μg/dL (média 286 μg/dL) - baseado em estudo populacional Friedrich et al. 2008
• 25-29 anos: 200-350 μg/dL (faixa laboratorial ampla: 110-510 μg/dL)
• Conversão: μg/dL × 0.027 = μmol/L

INTERPRETAÇÃO CLÍNICA:

NÍVEIS BAIXOS (<110 μg/dL):
1. Investigar insuficiência adrenal primária ou secundária
2. Solicitar: cortisol basal 8h, ACTH, eletrólitos (Na+, K+), glicemia
3. Se cortisol <10 μg/dL + DHEA-S baixo: ~72% probabilidade de insuficiência adrenal
4. Considerar teste de estímulo com ACTH se diagnóstico incerto
5. Avaliar uso recente de glicocorticóides (reduz acurácia diagnóstica do DHEA-S)
6. Descartar: doença de Addison, hipopituitarismo, supressão iatrogênica
7. Sintomas associados: fadiga profunda, hipotensão, hiperpigmentação cutânea

NÍVEIS NORMAIS (110-510 μg/dL):
1. Função adrenal adequada
2. Estabelecer valor baseline individual para monitoramento futuro
3. Reavaliar anualmente se histórico familiar de doenças adrenais
4. Correlacionar com testosterona total e livre para avaliação androgênica completa

NÍVEIS ELEVADOS (>510 μg/dL):
1. Investigar hiperplasia adrenal congênita (deficiência de 21-hidroxilase)
2. Descartar tumor produtor de andrógenos (adrenal ou testicular)
3. Avaliar síndrome de Cushing (solicitar cortisol livre urinário 24h)
4. Solicitar: 17-OH-progesterona, androstenediona, testosterona total
5. Considerar imagem adrenal (TC ou RNM) se elevação marcada (>2x limite superior)
6. Avaliar sinais de virilização ou síndrome metabólica

CONDUTA ESPECÍFICA PARA FAIXA ETÁRIA 20-29 ANOS:
• DHEA-S baixo nesta idade é SEMPRE patológico (não atribuir a envelhecimento)
• Combinação DHEA-S + cortisol basal tem alto valor preditivo negativo
• Se ambos normais: probabilidade de insuficiência adrenal <1.2%
• Considerar fatores genéticos: 37-46% da variabilidade é hereditária
• Avaliar contexto clínico: estresse crônico, exercício extremo, desnutrição podem reduzir níveis
• Monitorar tendência: declínio >5% ao ano pode indicar patologia subjacente

SEGUIMENTO:
• Valores normais + assintomático: reavaliar em 1-2 anos
• Valores borderline: repetir em 3-6 meses com cortisol matinal simultâneo
• Valores anormais: encaminhar endocrinologista para investigação aprofundada
• Documentar medicações em uso (glicocorticóides, esteróides anabolizantes, opioides)

CORRELAÇÕES IMPORTANTES:
• DHEA-S contribui com ~30% dos andrógenos totais em homens
• Avaliar junto com testosterona para visão completa do status androgênico
• Baixo DHEA-S + testosterona normal: sugere problema adrenal isolado
• Ambos baixos: considerar hipogonadismo hipogonadotrófico ou insuficiência adrenal secundária',

    last_review = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = 'a6742c97-a610-4bd4-b9de-87e91ecc8d34';

COMMIT;

-- Verification query
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
WHERE si.id = 'a6742c97-a610-4bd4-b9de-87e91ecc8d34'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;

-- Show linked articles
SELECT
    a.title,
    a.authors,
    a.journal,
    a.publish_date,
    a.doi,
    a.pm_id,
    a.article_type
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = 'a6742c97-a610-4bd4-b9de-87e91ecc8d34'
ORDER BY a.publish_date DESC;
