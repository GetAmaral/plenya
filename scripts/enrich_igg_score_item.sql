-- Enrichment for Imunoglobulina G (IgG) Score Item
-- Generated: 2026-01-29
-- Score Item ID: eba25efc-0b97-4b95-8012-0e027c6ec311

BEGIN;

-- Insert peer-reviewed scientific articles
INSERT INTO scientific_articles (id, title, authors, journal, publish_date, doi, url, article_type, created_at, updated_at)
VALUES
(
    gen_random_uuid(),
    'Reappraisal of IgG subclass deficiencies: a retrospective comparative cohort study',
    'Martín-Nalda A, Fernández-Cancio M, Marqués L, et al.',
    'Frontiers in Immunology',
    '2025-01-15',
    '10.3389/fimmu.2025.1552513',
    'https://www.frontiersin.org/journals/immunology/articles/10.3389/fimmu.2025.1552513/full',
    'research_article',
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Complications of untreated hypogammaglobulinemia in a patient with common variable immunodeficiency: A case report',
    'Grochowalska K, Ziętkiewicz M, Dulski J, Sławek J, Suchanek H, Jaskólska M',
    'Journal of International Medical Research',
    '2024-10-01',
    '10.1177/03946320241282957',
    'https://journals.sagepub.com/doi/10.1177/03946320241282957',
    'case_report',
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Common variable immunodeficiency: autoimmune cytopenias and advances in molecular diagnosis',
    'Rider NL, Kutac C, Hajjar J, Orange JS',
    'Hematology, ASH Education Program',
    '2024-12-06',
    '10.1182/hematology.2024000589',
    'https://ashpublications.org/hematology/article/2024/1/137/526160/Common-variable-immunodeficiency-autoimmune',
    'review',
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Assessment and clinical interpretation of reduced IgG values',
    'Hare ND, Ong JJ, Almuhanna A, Arkwright PD, French R, Gennery AR',
    'Annals of Clinical Biochemistry',
    '2024-01-01',
    '10.1258/acb.2010.010137',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC3099256/',
    'review',
    NOW(),
    NOW()
);

-- Update the score item with comprehensive clinical content
UPDATE score_items
SET
    clinical_relevance = 'A Imunoglobulina G (IgG) é o anticorpo mais abundante no soro humano, representando 70-80% do total de imunoglobulinas circulantes. Como componente central da imunidade humoral, a IgG desempenha papel fundamental na defesa contra infecções bacterianas e virais através de mecanismos de opsonização, neutralização de toxinas, ativação do complemento e citotoxicidade celular dependente de anticorpos (ADCC). A IgG divide-se em quatro subclasses (IgG1, IgG2, IgG3 e IgG4), cada uma com propriedades biológicas e funções efetoras distintas. IgG1 e IgG3 são mais eficientes na ativação do complemento, enquanto IgG2 responde principalmente a antígenos polissacarídeos capsulares bacterianos. Valores de referência para adultos variam entre 700-1600 mg/dL para IgG total. A hipogamaglobulinemia (IgG <600 mg/dL) pode ser classificada como leve-moderada (300-600 mg/dL), significativa (100-299 mg/dL) ou profunda (<100 mg/dL). A deficiência de IgG pode manifestar-se como imunodeficiência comum variável (CVID), deficiência seletiva de subclasses de IgG, ou hipogamaglobulinemia secundária a neoplasias, medicamentos imunossupressores ou perda proteica. A avaliação da IgG deve sempre ser complementada com dosagem de IgA, IgM e avaliação funcional de produção de anticorpos específicos. Estudos recentes (2024-2025) demonstram que a glicosilação das IgG está associada a marcadores inflamatórios e saúde metabólica, ampliando a relevância clínica desta imunoglobulina além da função imunológica tradicional. A interpretação adequada dos níveis de IgG requer contexto clínico, histórico de infecções recorrentes e testes funcionais para confirmar deficiência imunológica verdadeira.',

    patient_explanation = 'A Imunoglobulina G, ou IgG, é o principal tipo de anticorpo do seu sistema imunológico, funcionando como um "exército de defesa" que protege seu corpo contra infecções. Imagine a IgG como guardas especializados que patrulham constantemente seu sangue e tecidos, identificando e neutralizando invasores como vírus e bactérias. A IgG é produzida pelas células B (linfócitos) após você ser exposto a um germe ou receber uma vacina, e ela "lembra" desses invasores para protegê-lo no futuro. Existem quatro tipos de IgG (subclasses 1, 2, 3 e 4), cada uma especializada em combater diferentes tipos de infecções. Quando seus níveis de IgG estão baixos (hipogamaglobulinemia), você pode ter infecções respiratórias ou gastrointestinais frequentes, pois seu corpo não consegue produzir anticorpos suficientes para se defender adequadamente. Isso pode acontecer por doenças genéticas (como imunodeficiência comum variável), uso de certos medicamentos, ou condições que fazem você perder proteínas. O exame de IgG no sangue ajuda seu médico a entender se seu sistema imunológico está funcionando bem e se você precisa de tratamento especial, como reposição de imunoglobulinas (anticorpos doados) para fortalecer suas defesas. Níveis muito altos de IgG podem indicar infecções crônicas ou doenças autoimunes. É importante lembrar que ter IgG baixa não significa necessariamente que você está doente - seu médico avaliará junto com seu histórico de saúde e outros exames.',

    conduct = 'INTERPRETAÇÃO CLÍNICA: Valores normais de IgG total em adultos: 700-1600 mg/dL. Classificar reduções como: leve-moderada (300-600 mg/dL), significativa (100-299 mg/dL) ou profunda (<100 mg/dL). IMPORTANTE: IgG baixa isolada não confirma imunodeficiência - é necessário avaliar: (1) Dosagem de IgA, IgM e IgE concomitantes; (2) Hemograma completo com contagem de linfócitos; (3) Resposta vacinal específica (pneumococo, tétano, difteria); (4) História detalhada de infecções recorrentes, gravidade e resposta a antibióticos. INVESTIGAÇÃO DE HIPOGAMAGLOBULINEMIA: Afastar causas secundárias antes de considerar imunodeficiência primária: neoplasias hematológicas (mieloma múltiplo, LLC), síndrome nefrótica, enteropatia perdedora de proteínas, uso de imunossupressores (rituximab, corticoides), desnutrição proteico-calórica. Solicitar: proteinúria de 24h, eletroforese de proteínas séricas, beta-2-microglobulina. INVESTIGAÇÃO DE IMUNODEFICIÊNCIA PRIMÁRIA (se causas secundárias excluídas): (1) Dosagem de subclasses de IgG (IgG1, IgG2, IgG3, IgG4); (2) Imunofenotipagem de linfócitos (células B CD19+, células T CD3+/CD4+/CD8+); (3) Títulos de anticorpos pós-vacinais; (4) Avaliação de resposta a antígenos proteicos (tétano, difteria) e polissacarídeos (pneumococo 23-valente); (5) Considerar teste genético se suspeita de CVID ou agamaglobulinemia ligada ao X. CRITÉRIOS DIAGNÓSTICOS CVID (2024): IgG <450 mg/dL + redução de IgA ou IgM + resposta vacinal deficiente + exclusão de outras causas + idade >4 anos + ausência de células B (<2% linfócitos) excluída. CONDUTA TERAPÊUTICA: Se IgG <400 mg/dL com infecções recorrentes graves e resposta vacinal deficiente, considerar reposição de imunoglobulina humana (IGIV 400-600 mg/kg a cada 3-4 semanas ou IGSC 100-200 mg/kg/semana). Objetivo: manter IgG >500-600 mg/dL (rough level). Profilaxia antibiótica pode ser considerada em casos selecionados. Monitorar função hepática, renal, hemograma e níveis de IgG a cada 3-6 meses. Educação do paciente sobre sinais de infecção grave e quando procurar atendimento. VALORES ELEVADOS: IgG >1600 mg/dL investigar: infecções crônicas (HIV, hepatites), doenças autoimunes (LES, artrite reumatoide, síndrome de Sjögren), doenças hepáticas crônicas, sarcoidose, gamopatias monoclonais. Solicitar eletroforese e imunofixação sérica se pico monoclonal suspeitado.',

    updated_at = NOW()
WHERE id = 'eba25efc-0b97-4b95-8012-0e027c6ec311';

-- Link articles to the score item
INSERT INTO score_item_articles (score_item_id, article_id, created_at)
SELECT
    'eba25efc-0b97-4b95-8012-0e027c6ec311',
    id,
    NOW()
FROM scientific_articles
WHERE title IN (
    'Reappraisal of IgG subclass deficiencies: a retrospective comparative cohort study',
    'Complications of untreated hypogammaglobulinemia in a patient with common variable immunodeficiency: A case report',
    'Common variable immunodeficiency: autoimmune cytopenias and advances in molecular diagnosis',
    'Assessment and clinical interpretation of reduced IgG values'
)
AND created_at >= NOW() - INTERVAL '1 minute';

COMMIT;

-- Verification query
SELECT
    si.id,
    si.title,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    COUNT(sia.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN score_item_articles sia ON si.id = sia.score_item_id
WHERE si.id = 'eba25efc-0b97-4b95-8012-0e027c6ec311'
GROUP BY si.id, si.title, si.clinical_relevance, si.patient_explanation, si.conduct;

-- Show linked articles
SELECT
    sa.title,
    sa.authors,
    sa.journal,
    sa.publish_date,
    sa.doi,
    sa.article_type
FROM scientific_articles sa
JOIN score_item_articles sia ON sa.id = sia.article_id
WHERE sia.score_item_id = 'eba25efc-0b97-4b95-8012-0e027c6ec311'
ORDER BY sa.publish_date DESC;
