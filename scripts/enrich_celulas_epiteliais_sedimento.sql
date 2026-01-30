-- Enrichment for Score Item: Células Epiteliais - Sedimento
-- ID: 09577ef1-c3ad-461b-b2ad-59fab2c193d5
-- Generated: 2026-01-28

BEGIN;

-- Insert Article 1: Brazilian survey on epithelial cells reporting
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract
) VALUES (
    'Survey on reporting of epithelial cells in urine sediment as part of external quality assessment programs in Brazilian laboratories',
    'José A T Poloni, Adriana de Oliveira Vieira, Caroline R M Dos Santos, Ana-Maria Simundic, Liane N Rotta',
    'Biochem Med (Zagreb)',
    '2021-06-15'::date,
    'en',
    'research_article',
    '10.11613/BM.2021.020711',
    '34140834',
    'This research investigates how Brazilian laboratories report epithelial cells during urine microscopy examination. The study surveyed 1,336 laboratories participating in an external quality assessment program, organizing them by their classification methods. Key findings revealed that while most laboratories recognize the clinical significance of distinguishing renal tubular epithelial cells from other epithelial subtypes, the majority do not currently differentiate between the three main types. The authors conclude that education of laboratory staff about the clinical significance of urinary particles should be considered a key priority.'
) ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id AS article_1_id;

-- Insert Article 2: Diagnostic utility of urine microscopy
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract
) VALUES (
    'Diagnostic Utility of Urine Microscopy in Kidney Diseases',
    'Payal Gaggar, Sree B Raju',
    'Indian Journal of Nephrology',
    '2024-05-28'::date,
    'en',
    'review',
    '10.25259/ijn_362_23',
    '39114391',
    'Urine sediment analysis is a highly valuable yet underutilized test in contemporary medicine. Microscopic examination represents a simple, cost-effective, and powerful diagnostic tool for nephrologists, particularly beneficial in acute kidney injury contexts. The review discusses how timely sediment analysis enables compartmental AKI diagnosis and prognostic guidance. A key concern raised is that automated systems fail identifying pathological casts, crystals, and dysmorphic red blood cells, necessitating renewed focus on manual microscopic analysis for accurate identification of renal tubular epithelial cells and other critical urinary particles.'
) ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id AS article_2_id;

-- Insert Article 3: Urinary sediment microscopy correlations with kidney biopsy
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract
) VALUES (
    'Urinary Sediment Microscopy and Correlations with Kidney Biopsy: Red Flags Not To Be Missed',
    'David Navarro, Nuno Moreira Fonseca, Ana Carina Ferreira, Rui Barata, Mário Góis, Helena Sousa, Fernando Nolasco',
    'Kidney360',
    '2022-09-12'::date,
    'en',
    'research_article',
    '10.34067/KID.0003082022',
    '36700902',
    'This study investigated associations between manual urinary sediment particles and histologic lesions identified through kidney biopsy in 131 patients. Key findings revealed that urinary lipids correlated with mesangial expansion and renal tubular epithelial cells (RTECs) associated with endocapillary hypercellularity and cellular crescents. Manual microscopic evaluation identified particles including dysmorphic erythrocytes, RTECs, and lipids that automated analyzers frequently miss, suggesting these findings warrant kidney biopsy consideration and highlighting the clinical importance of proper epithelial cell identification.'
) ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id AS article_3_id;

-- Insert Article 4: Acute Renal Tubular Necrosis - StatPearls
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    pm_id,
    abstract
)
SELECT
    'Acute Renal Tubular Necrosis',
    'Muhammad O Hanif, Preeti Rout, Kamleshun Ramphul',
    'StatPearls [Internet]',
    '2025-08-08'::date,
    'en',
    'review',
    '29939592',
    'This educational resource describes ATN as the most common intrinsic cause of acute kidney injury, particularly among hospitalized patients. The condition results from ischemic or nephrotoxic mechanisms affecting renal tubules while typically sparing glomeruli. Key findings include ferroptosis identified as the predominant cell death mechanism. Diagnostic approach is primarily clinical, with urinalysis showing muddy brown casts or renal tubular epithelial cells secondary to the sloughing of tubular cells into the lumen. Recognition that ATN may progress to chronic kidney disease or end-stage renal disease, especially in patients with diabetes, heart failure, or pre-existing kidney disease.'
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '29939592'
)
RETURNING id AS article_4_id;

-- Update Score Item with clinical content
UPDATE score_items
SET
    clinical_relevance = 'As células epiteliais no sedimento urinário são classificadas em três tipos principais, cada um com significado clínico distinto: células epiteliais escamosas (originárias da uretra externa e região genital), células epiteliais transicionais ou uroteliais (provenientes da bexiga e sistema coletor), e células epiteliais tubulares renais (originadas dos túbulos renais). As células epiteliais escamosas são as mais comuns e geralmente indicam contaminação da amostra, sendo normal encontrar 1-5 células/campo; valores superiores a 15-20 células/campo sugerem coleta inadequada. As células epiteliais transicionais podem aparecer em processos inflamatórios da bexiga ou sistema coletor. As células tubulares renais (RTECs) possuem a maior relevância clínica, pois sua presença sempre indica lesão tubular ativa. Estudos recentes demonstram que RTECs se correlacionam com hipercelularidade endocapilar, crescentes celulares e necrose tubular aguda (NTA) em biópsias renais. Na NTA, a descamação do epitélio tubular produz RTECs, cilindros de células epiteliais tubulares e os característicos cilindros granulares castanho-escuros ("muddy brown casts"), que apresentam especificidade de 100% para NTA. A diferenciação adequada entre os subtipos celulares é fundamental, porém estudos brasileiros mostram que a maioria dos laboratórios não realiza essa distinção, apesar de seu impacto clínico significativo. Sistemas automatizados falham na identificação de RTECs, reforçando a importância da microscopia manual por profissionais treinados.',

    patient_explanation = 'As células epiteliais são células de revestimento que normalmente cobrem diversas partes do sistema urinário. Quando aparecem na urina, podem ter diferentes significados dependendo de qual região se originam. Existem três tipos principais: células da pele ao redor da uretra (as mais comuns e geralmente sem importância clínica), células da bexiga (podem indicar inflamação), e células dos rins (as mais importantes clinicamente). As células da pele aparecem em quase todas as amostras de urina e apenas indicam que a amostra pode estar contaminada se estiverem em grande quantidade - isso acontece frequentemente quando a coleta não é feita adequadamente. Já as células que vêm dos rins (células tubulares renais) são mais raras e sempre merecem atenção, pois indicam que os túbulos renais estão sofrendo alguma lesão. Quando essas células aparecem junto com cilindros escuros na urina, podem sinalizar uma condição chamada necrose tubular aguda, que é uma lesão mais grave dos rins que precisa de acompanhamento médico. Por isso, é importante que a coleta de urina seja feita corretamente, seguindo as orientações do laboratório, para que o resultado seja confiável. Se seu exame mostrar células epiteliais, seu médico avaliará qual tipo está presente e se há necessidade de investigação adicional.',

    conduct = 'A conduta clínica diante da presença de células epiteliais no sedimento urinário deve ser estratificada conforme o subtipo identificado. Para células epiteliais escamosas: valores de 1-5 células/campo são normais e não requerem ação; contagens superiores a 15-20 células/campo indicam contaminação e a amostra deve ser rejeitada, solicitando-se nova coleta com orientações adequadas sobre higiene perineal e técnica de jato médio. Para células epiteliais transicionais: presença em pequena quantidade pode ser normal; aumento significativo justifica investigação de processos inflamatórios ou infecciosos do trato urinário inferior, correlacionando com sintomas clínicos, urocultura e exames de imagem quando apropriado. Para células tubulares renais (RTECs): sua presença sempre é patológica e indica lesão tubular ativa, exigindo investigação imediata. Protocolo recomendado: (1) Confirmar achado com nova amostra adequadamente coletada para excluir contaminação; (2) Avaliar contexto clínico de lesão renal aguda (LRA) - verificar creatinina sérica, ureia, eletrólitos e balanço hídrico; (3) Pesquisar ativamente cilindros associados (especialmente cilindros granulares castanho-escuros patognomônicos de NTA, cilindros de células epiteliais tubulares); (4) Investigar etiologia: isquemia renal (hipotensão, sepse, cirurgia), nefrotoxicidade (aminoglicosídeos, contraste iodado, AINEs, quimioterápicos), rabdomiólise, ou processos glomerulares; (5) Considerar biópsia renal se achados sugerem doença glomerular ou quando o diagnóstico permanece incerto, especialmente se houver RTECs associadas a cilindros lipídicos, hematúria dismórfica ou proteinúria significativa; (6) Monitoramento evolutivo com exames seriados de função renal e sedimento urinário para avaliar resolução ou progressão da lesão; (7) Implementar medidas de nefroproteção: otimizar volemia, suspender nefrotóxicos, ajustar doses de medicamentos para função renal. A identificação correta de RTECs tem valor prognóstico, pois estudos demonstram associação com desfechos renais adversos em pacientes com nefropatia diabética e outras glomerulopatias. A microscopia manual por profissional treinado é essencial, pois sistemas automatizados apresentam sensibilidade inferior a 50% para identificação de RTECs, apesar de especificidade superior a 98%.',

    last_review = CURRENT_TIMESTAMP,
    updated_at = NOW()
WHERE id = '09577ef1-c3ad-461b-b2ad-59fab2c193d5';

-- Link articles to score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '09577ef1-c3ad-461b-b2ad-59fab2c193d5'::uuid
FROM articles a
WHERE (a.doi IN (
    '10.11613/BM.2021.020711',
    '10.25259/ijn_362_23',
    '10.34067/KID.0003082022'
) OR a.pm_id = '29939592')
AND NOT EXISTS (
    SELECT 1 FROM article_score_items asi
    WHERE asi.article_id = a.id
    AND asi.score_item_id = '09577ef1-c3ad-461b-b2ad-59fab2c193d5'::uuid
);

COMMIT;

-- Verification Query
SELECT
    si.id,
    si.name,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    si.last_review,
    COUNT(DISTINCT asi.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '09577ef1-c3ad-461b-b2ad-59fab2c193d5'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;
