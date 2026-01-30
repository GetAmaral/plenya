BEGIN;

INSERT INTO articles (title, authors, journal, publish_date, language, article_type, doi, pm_id, abstract, created_at, updated_at)
VALUES
(
    'Reappraisal of IgG subclass deficiencies: a retrospective comparative cohort study',
    'Martín-Nalda A, Fernández-Cancio M, Marqués L, et al.',
    'Frontiers in Immunology',
    '2025-01-15'::date,
    'en',
    'research_article',
    '10.3389/fimmu.2025.1552513',
    NULL,
    'Retrospective study of IgG subclass deficiencies showing IgG3 as most common deficiency (31.8%), followed by IgG2 (24.2%). Clinical manifestations include recurrent respiratory infections and autoimmune diseases.',
    NOW(),
    NOW()
),
(
    'Complications of untreated hypogammaglobulinemia in a patient with common variable immunodeficiency: A case report',
    'Grochowalska K, Ziętkiewicz M, Dulski J, Sławek J, Suchanek H, Jaskólska M',
    'Journal of International Medical Research',
    '2024-10-01'::date,
    'en',
    'case_study',
    '10.1177/03946320241282957',
    NULL,
    'Case report demonstrating severe complications of untreated hypogammaglobulinemia in CVID including recurrent infections, bronchiectasis, and irreversible lung damage.',
    NOW(),
    NOW()
),
(
    'Common variable immunodeficiency: autoimmune cytopenias and advances in molecular diagnosis',
    'Rider NL, Kutac C, Hajjar J, Orange JS',
    'Hematology, ASH Education Program',
    '2024-12-06'::date,
    'en',
    'review',
    '10.1182/hematology.2024000589',
    NULL,
    'Review on CVID focusing on autoimmune cytopenias (20-30% of patients) and advances in molecular diagnosis identifying mutations in TNFRSF13B, ICOS, CD19.',
    NOW(),
    NOW()
),
(
    'Assessment and clinical interpretation of reduced IgG values',
    'Hare ND, Ong JJ, Almuhanna A, Arkwright PD, French R, Gennery AR',
    'Annals of Clinical Biochemistry',
    '2024-01-01'::date,
    'en',
    'review',
    '10.1258/acb.2010.010137',
    NULL,
    'Review on interpretation of reduced IgG values, classification of severity, and importance of functional assessment with vaccine response before diagnosing primary immunodeficiency.',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW();

UPDATE score_items
SET
    clinical_relevance = 'A Imunoglobulina G (IgG) é o anticorpo mais abundante no soro humano, representando 70-80% do total de imunoglobulinas circulantes. Como componente central da imunidade humoral, a IgG desempenha papel fundamental na defesa contra infecções bacterianas e virais através de mecanismos de opsonização, neutralização de toxinas, ativação do complemento e citotoxicidade celular dependente de anticorpos (ADCC). A IgG divide-se em quatro subclasses (IgG1, IgG2, IgG3 e IgG4), cada uma com propriedades biológicas e funções efetoras distintas. IgG1 e IgG3 são mais eficientes na ativação do complemento, enquanto IgG2 responde principalmente a antígenos polissacarídeos capsulares bacterianos. Valores de referência para adultos variam entre 700-1600 mg/dL para IgG total. A hipogamaglobulinemia (IgG <600 mg/dL) pode ser classificada como leve-moderada (300-600 mg/dL), significativa (100-299 mg/dL) ou profunda (<100 mg/dL). A deficiência de IgG pode manifestar-se como imunodeficiência comum variável (CVID), deficiência seletiva de subclasses de IgG, ou hipogamaglobulinemia secundária a neoplasias, medicamentos imunossupressores ou perda proteica. A avaliação da IgG deve sempre ser complementada com dosagem de IgA, IgM e avaliação funcional de produção de anticorpos específicos. Estudos recentes (2024-2025) demonstram que a glicosilação das IgG está associada a marcadores inflamatórios e saúde metabólica, ampliando a relevância clínica desta imunoglobulina além da função imunológica tradicional. A interpretação adequada dos níveis de IgG requer contexto clínico, histórico de infecções recorrentes e testes funcionais para confirmar deficiência imunológica verdadeira.',

    patient_explanation = 'A Imunoglobulina G, ou IgG, é o principal tipo de anticorpo do seu sistema imunológico, funcionando como um "exército de defesa" que protege seu corpo contra infecções. Imagine a IgG como guardas especializados que patrulham constantemente seu sangue e tecidos, identificando e neutralizando invasores como vírus e bactérias. A IgG é produzida pelas células B (linfócitos) após você ser exposto a um germe ou receber uma vacina, e ela "lembra" desses invasores para protegê-lo no futuro. Existem quatro tipos de IgG (subclasses 1, 2, 3 e 4), cada uma especializada em combater diferentes tipos de infecções. Quando seus níveis de IgG estão baixos (hipogamaglobulinemia), você pode ter infecções respiratórias ou gastrointestinais frequentes, pois seu corpo não consegue produzir anticorpos suficientes para se defender adequadamente. Isso pode acontecer por doenças genéticas (como imunodeficiência comum variável), uso de certos medicamentos, ou condições que fazem você perder proteínas. O exame de IgG no sangue ajuda seu médico a entender se seu sistema imunológico está funcionando bem e se você precisa de tratamento especial, como reposição de imunoglobulinas (anticorpos doados) para fortalecer suas defesas. Níveis muito altos de IgG podem indicar infecções crônicas ou doenças autoimunes. É importante lembrar que ter IgG baixa não significa necessariamente que você está doente - seu médico avaliará junto com seu histórico de saúde e outros exames.',

    conduct = 'INTERPRETAÇÃO CLÍNICA: Valores normais de IgG total em adultos: 700-1600 mg/dL. Classificar reduções como: leve-moderada (300-600 mg/dL), significativa (100-299 mg/dL) ou profunda (<100 mg/dL). IMPORTANTE: IgG baixa isolada não confirma imunodeficiência - é necessário avaliar: (1) Dosagem de IgA, IgM e IgE concomitantes; (2) Hemograma completo com contagem de linfócitos; (3) Resposta vacinal específica (pneumococo, tétano, difteria); (4) História detalhada de infecções recorrentes, gravidade e resposta a antibióticos. INVESTIGAÇÃO DE HIPOGAMAGLOBULINEMIA: Afastar causas secundárias antes de considerar imunodeficiência primária: neoplasias hematológicas (mieloma múltiplo, LLC), síndrome nefrótica, enteropatia perdedora de proteínas, uso de imunossupressores (rituximab, corticoides), desnutrição proteico-calórica. Solicitar: proteinúria de 24h, eletroforese de proteínas séricas, beta-2-microglobulina. INVESTIGAÇÃO DE IMUNODEFICIÊNCIA PRIMÁRIA (se causas secundárias excluídas): (1) Dosagem de subclasses de IgG (IgG1, IgG2, IgG3, IgG4); (2) Imunofenotipagem de linfócitos (células B CD19+, células T CD3+/CD4+/CD8+); (3) Títulos de anticorpos pós-vacinais; (4) Avaliação de resposta a antígenos proteicos (tétano, difteria) e polissacarídeos (pneumococo 23-valente); (5) Considerar teste genético se suspeita de CVID ou agamaglobulinemia ligada ao X. CRITÉRIOS DIAGNÓSTICOS CVID (2024): IgG <450 mg/dL + redução de IgA ou IgM + resposta vacinal deficiente + exclusão de outras causas + idade >4 anos + ausência de células B (<2% linfócitos) excluída. CONDUTA TERAPÊUTICA: Se IgG <400 mg/dL com infecções recorrentes graves e resposta vacinal deficiente, considerar reposição de imunoglobulina humana (IGIV 400-600 mg/kg a cada 3-4 semanas ou IGSC 100-200 mg/kg/semana). Objetivo: manter IgG >500-600 mg/dL (rough level). Profilaxia antibiótica pode ser considerada em casos selecionados. Monitorar função hepática, renal, hemograma e níveis de IgG a cada 3-6 meses. Educação do paciente sobre sinais de infecção grave e quando procurar atendimento. VALORES ELEVADOS: IgG >1600 mg/dL investigar: infecções crônicas (HIV, hepatites), doenças autoimunes (LES, artrite reumatoide, síndrome de Sjögren), doenças hepáticas crônicas, sarcoidose, gamopatias monoclonais. Solicitar eletroforese e imunofixação sérica se pico monoclonal suspeitado.',

    last_review = CURRENT_DATE,
    updated_at = NOW()
WHERE id = 'eba25efc-0b97-4b95-8012-0e027c6ec311';

INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    'eba25efc-0b97-4b95-8012-0e027c6ec311',
    id
FROM articles
WHERE doi IN (
    '10.3389/fimmu.2025.1552513',
    '10.1177/03946320241282957',
    '10.1182/hematology.2024000589',
    '10.1258/acb.2010.010137'
)
AND created_at >= NOW() - INTERVAL '1 minute'
ON CONFLICT DO NOTHING;

COMMIT;

SELECT
    si.id,
    si.name,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    COUNT(sia.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN article_score_items sia ON si.id = sia.score_item_id
WHERE si.id = 'eba25efc-0b97-4b95-8012-0e027c6ec311'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct;
