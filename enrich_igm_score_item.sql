-- Enrichment for Score Item: Imunoglobulina M (IgM)
-- ID: 50b7bdd2-8b09-4c53-aef1-5d36ccd30ebc
-- Generated: 2026-01-29

BEGIN;

-- Insert Article 1: IgM Structure and Pentamer Formation (Nature Communications, 2022)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract,
    created_at,
    updated_at
)
VALUES (
    'Cryomicroscopy reveals the structural basis for a flexible hinge motion in the immunoglobulin M pentamer',
    'Qu Chen, Rajesh Menon, Lesley J Calder, Pavel Tolar, Peter B Rosenthal',
    'Nature Communications',
    '2022-10-23'::date,
    'en',
    'research_article',
    '10.1038/s41467-022-34090-2',
    '36274064',
    'IgM represents the most ancient of the five isotypes of immunoglobulin (Ig) molecules and serves as the first line of defence. Using cryo-EM imaging, researchers visualized the complete human IgM pentamer structure, revealing that antigen-binding domains are flexibly connected to an asymmetric core formed by constant regions and the J-chain. A hinge located at the Cμ3/Cμ2 domain interface permits coordinated pivoting of Fabs and Cμ2 both parallel and perpendicular to the plane. This differs from IgG and IgA, where Fab arms move independently. Asymmetry in the Cμ3 domain creates a biased orientation in one Fab pair, potentially affecting multivalent antigen binding and complement activation. The monomer Fc structure resembles the pentamer version but exhibits greater dynamic behavior in the Cμ4 domain.',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW();

-- Insert Article 2: IgM Receptor Recognition (Nature, 2023)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract,
    created_at,
    updated_at
)
VALUES (
    'Immunoglobulin M perception by FcμR',
    'Yaxin Li, Hao Shen, Ruixue Zhang, Chenggong Ji, Yuxin Wang, Chen Su, Junyu Xiao',
    'Nature',
    '2023-03-22'::date,
    'en',
    'research_article',
    '10.1038/s41586-023-05835-w',
    '36949194',
    'This research examines how FcμR, the sole IgM-specific receptor in mammals, recognizes and interacts with different forms of immunoglobulin M. Using crystallography and cryo-electron microscopy, researchers determined that two FcμR molecules interact with a Fcμ-Cμ4 dimer in membrane-bound contexts. The study reveals FcμR binds with different stoichiometries depending on IgM form: a 2:1 ratio for membrane-bound IgM within the B cell receptor complex, while pentameric IgM recruits four FcμR molecules on one side, and secretory IgM associates with four FcμR molecules on its alternate side. These findings illuminate the structural mechanisms underlying IgM recognition across its multiple cellular forms and provide insights into B cell regulation and immune responses mediated by IgM.',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW();

-- Insert Article 3: Selective IgM Deficiency (Diagnostics, 2023)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract,
    created_at,
    updated_at
)
VALUES (
    'Selective IgM Deficiency: Evidence, Controversies, and Gaps',
    'Ivan Taietti, Martina Votto, Maria De Filippo, Matteo Naso, Lorenza Montagna, Daniela Montagna, Amelia Licari, Gian Luigi Marseglia, Riccardo Castagnoli',
    'Diagnostics (Basel)',
    '2023-09-04'::date,
    'en',
    'review',
    '10.3390/diagnostics13172861',
    '37685399',
    'This review examines selective immunoglobulin M deficiency (SIgMD), recently classified as an inborn error of immunity. The authors note that the understanding of SIgMD is still extremely limited, especially in the pediatric population. The pathogenesis remains elusive with no established genetic or molecular basis identified. Recurrent respiratory infections represent the main clinical manifestations in children, followed by allergic and autoimmune diseases. According to ESID criteria, SIgMD is defined as repeatedly absent or reduced serum IgM levels (less than 2 SD or <10% of values from healthy controls or an absolute value <20 mg/dL in pediatric age) with normal IgA and IgG levels, normal vaccine responses, and absence of T cell defects. Immunoglobulin replacement therapy is not universally required, though it may be recommended for patients with significantly associated antibody deficiency. Prophylactic antibiotics and prompt treatment of fever remain important management strategies.',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW();

-- Insert Article 4: IgM-enriched immunoglobulin therapeutic use (Minerva Anestesiol, 2023)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract,
    created_at,
    updated_at
)
VALUES (
    'Potential role of IgM-enriched immunoglobulin as adjuvant treatment in severe SARS-CoV-2 infection',
    'Carlo Tascini, Marco Cotrufo, Emanuela Sozio, Matteo Fanin, Fabiana Dellai, Agnese Zanus Forte, Daniela Cesselli, Paola DE Stefanis, Andrea Ripoli, Francesco Sbrana, Simone Giuliano, Martina Fabris, Massimo Girardis, Francesco Curcio, Flavio Bassi',
    'Minerva Anestesiol',
    '2023-10-01'::date,
    'en',
    'clinical_trial',
    '10.23736/S0375-9393.23.17244-0',
    '37822148',
    'This study examined IgM-enriched intravenous immunoglobulins (Pentaglobin) in severe COVID-19 patients presenting during late disease stages. In this single-center retrospective case-control analysis, 56 treated patients were matched against 169 untreated controls. Results indicated that the Pentaglobin treatment was identified as a significant protective factor for death outcome with improved D-dimer and P/F ratios in the treatment group. The research offers insight into immunoglobulin preparations for severely infected patients and identifies candidate profiles for this therapeutic approach, demonstrating the potential clinical utility of IgM-enriched preparations in critical care settings with severe viral infections.',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW();

-- Update score_items with comprehensive clinical content in Portuguese
UPDATE score_items
SET
    clinical_relevance = 'A Imunoglobulina M (IgM) é o anticorpo mais antigo filogeneticamente e representa a primeira linha de defesa da resposta imune adaptativa. É a primeira imunoglobulina secretada pelas células B em resposta a um antígeno estranho, caracterizando a resposta imune primária. A IgM possui estrutura pentamérica única com peso molecular superior a 900 kDa, contendo 10 sítios de ligação ao antígeno e uma cadeia J (joining chain), o que confere alta avidez apesar de menor afinidade individual. Esta conformação pentamérica permite neutralização eficaz de patógenos e ativação potente do sistema complemento. Os valores de referência para adultos variam entre 40-230 mg/dL, com meia-vida sérica de 5-10 dias. A IgM não atravessa a placenta, portanto sua presença em recém-nascidos indica infecção intrauterina. Níveis elevados sugerem infecção aguda ou recente, enquanto deficiência seletiva (IgM <20 mg/dL com IgG e IgA normais) associa-se a infecções respiratórias recorrentes, doenças autoimunes e alérgicas. Alterações nos níveis de IgM possuem significado clínico importante no diagnóstico diferencial de imunodeficiências primárias (como síndrome de Hiper-IgM), infecções agudas, doenças autoimunes (lúpus eritematoso sistêmico) e gamopatias monoclonais (macroglobulinemia de Waldenström). Pesquisas recentes revelaram detalhes estruturais através de criomicroscopia eletrônica, demonstrando a base molecular da flexibilidade conformacional da IgM e seus mecanismos de reconhecimento pelo receptor FcμR, único receptor específico para IgM em mamíferos.',
    patient_explanation = 'A Imunoglobulina M, ou IgM, é um tipo especial de proteína de defesa que seu corpo produz como primeira resposta quando entra em contato com uma infecção ou invasor estranho. Imagine a IgM como os "soldados de resposta rápida" do seu sistema imunológico - eles são os primeiros a chegar ao local de uma infecção, antes mesmo que outros anticorpos mais especializados sejam produzidos. A IgM tem uma forma única parecida com uma estrela de cinco pontas, com dez "braços" que podem agarrar e neutralizar vírus, bactérias e outros germes. Por ser a primeira a aparecer, a presença de IgM no sangue geralmente indica que você teve uma infecção recente ou está com uma infecção ativa no momento. Valores normais para adultos ficam entre 40 e 230 mg/dL. Quando seus níveis de IgM estão aumentados, pode significar que seu corpo está combatendo ativamente uma infecção. Por outro lado, níveis muito baixos podem indicar que seu sistema imunológico não está respondendo adequadamente, deixando você mais vulnerável a infecções repetidas, especialmente do sistema respiratório. É importante lembrar que a IgM não passa da mãe para o bebê durante a gravidez, então se um recém-nascido tem IgM, isso indica que ele próprio teve uma infecção ainda no útero. Seu médico pode solicitar este exame junto com outros anticorpos (IgG e IgA) para ter uma visão completa de como seu sistema imunológico está funcionando.',
    conduct = 'A interpretação dos níveis de IgM deve ser realizada em conjunto com outros parâmetros imunológicos (IgG, IgA, subclasses de IgG) e o contexto clínico do paciente. Para IgM elevada (>230 mg/dL): investigar infecções agudas virais (hepatite A, citomegalovírus, Epstein-Barr, rubéola) ou bacterianas (Mycoplasma pneumoniae), realizar sorologia específica com pesquisa de IgM anti-patógeno. Considerar doenças autoimunes como lúpus eritematoso sistêmico (solicitar FAN, anti-DNA, complemento) ou síndrome de Sjögren. Descartar gamopatia monoclonal IgM através de eletroforese de proteínas e imunofixação sérica; se presente, avaliar macroglobulinemia de Waldenström com mielograma. Em níveis persistentemente elevados sem etiologia definida, considerar síndrome de Hiper-IgM - solicitar citometria de fluxo para avaliar expressão de CD40L em linfócitos T ativados. Para IgM baixa (<40 mg/dL): confirmar com segunda dosagem após 3-4 semanas. Definir se é deficiência seletiva de IgM (IgG e IgA normais, <20 mg/dL) ou parte de imunodeficiência combinada. Avaliar histórico de infecções recorrentes, especialmente respiratórias por bactérias encapsuladas (Streptococcus pneumoniae, Haemophilus influenzae). Solicitar avaliação funcional com resposta vacinal a antígenos polissacarídicos (pneumococo) e proteicos (tétano). Hemograma completo com contagem de linfócitos e subpopulações (CD3, CD4, CD8, CD19, CD16/56). Excluir causas secundárias: medicamentos imunossupressores, neoplasias hematológicas, perdas proteicas (síndrome nefrótica, enteropatia perdedora de proteínas). Para deficiência seletiva de IgM sintomática: considerar profilaxia antibiótica com amoxicilina ou azitromicina, vacinação pneumocócica, tratamento precoce e agressivo de infecções. Reposição de imunoglobulina endovenosa geralmente não é indicada na deficiência isolada de IgM (preparações comerciais contêm principalmente IgG), mas pode ser considerada se houver deficiência funcional associada de IgG. Encaminhamento ao imunologista é recomendado para casos de deficiência grave, infecções recorrentes ou suspeita de imunodeficiência primária. Monitoramento periódico a cada 6-12 meses com dosagem de imunoglobulinas, hemograma e avaliação clínica.',
    last_review = CURRENT_DATE,
    updated_at = NOW()
WHERE id = '50b7bdd2-8b09-4c53-aef1-5d36ccd30ebc';

-- Link articles to the score item
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '50b7bdd2-8b09-4c53-aef1-5d36ccd30ebc',
    id
FROM articles
WHERE doi IN (
    '10.1038/s41467-022-34090-2',
    '10.1038/s41586-023-05835-w',
    '10.3390/diagnostics13172861',
    '10.23736/S0375-9393.23.17244-0'
)
AND created_at >= NOW() - INTERVAL '1 minute'
ON CONFLICT DO NOTHING;

COMMIT;

-- Verification query
SELECT
    si.id,
    si.name,
    si.last_review,
    LENGTH(si.clinical_relevance) as clinical_relevance_length,
    LENGTH(si.patient_explanation) as patient_explanation_length,
    LENGTH(si.conduct) as conduct_length,
    COUNT(asi.article_id) as linked_articles
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '50b7bdd2-8b09-4c53-aef1-5d36ccd30ebc'
GROUP BY si.id, si.name, si.last_review, si.clinical_relevance, si.patient_explanation, si.conduct;
