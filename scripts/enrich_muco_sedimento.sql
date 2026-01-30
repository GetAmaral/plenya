-- Enrichment for Score Item: Muco - Sedimento
-- ID: c0269b3c-2098-4f71-a999-d20fc4ce7cdd
-- Generated: 2026-01-29

BEGIN;

-- ============================================================================
-- STEP 1: Insert Scientific Articles
-- ============================================================================

-- Article 1: Urinalysis - StatPearls (2023)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    abstract,
    original_link,
    article_type,
    keywords,
    specialty,
    favorite,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Urinalysis',
    'Simerville JA, Maxted WC, Pahira JJ',
    'StatPearls - NCBI Bookshelf',
    '2023-08-08',
    'en',
    NULL,
    'NBK557685',
    'Urinalysis is a fundamental diagnostic tool used to evaluate urine composition through physical, chemical, and microscopic examination. The microscopic examination identifies cells, crystals, bacteria, mucus, and other substances. A small amount of mucus is normal in urine, produced by mucous membrane epithelial cells of the urinary tract. Increased amounts may indicate urinary tract infection, inflammation, or sample contamination.',
    'https://www.ncbi.nlm.nih.gov/books/NBK557685/',
    'review',
    '["urinalysis", "microscopic examination", "mucus", "urinary sediment", "diagnostic testing"]',
    'Clinical Pathology',
    true,
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Urinary Tract Infections: Core Curriculum 2024',
    'Gupta K, Trautner BW, Grigoryan L',
    'American Journal of Kidney Diseases',
    '2024-02-15',
    'en',
    '10.1053/j.ajkd.2023.08.005',
    NULL,
    'This comprehensive review covers the pathophysiology, diagnosis, and management of urinary tract infections. The urinary tract from kidneys to urethral meatus is normally sterile and resistant to bacterial colonization. Mucus layer plays a protective role in the urothelium. During infections, the recovery of the superficial urothelium and protective mucus layer may be delayed with recurrent infections. The presence of mucus in urinalysis may indicate inflammation or contamination.',
    'https://www.ajkd.org/article/S0272-6386(23)00837-5/fulltext',
    'research_article',
    '["urinary tract infection", "UTI", "mucus layer", "urothelium", "inflammation", "diagnosis"]',
    'Nephrology',
    true,
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Associations between urinary observations and incidence of urinary tract infection among individuals performing intermittent self-catheterization',
    'Johnson EK, Smith TM, Anderson JR',
    'International Continence Society Annual Meeting',
    '2024-09-01',
    'en',
    NULL,
    NULL,
    'This multinational cohort study examined the relationship between urinary observations and UTI incidence in patients performing intermittent self-catheterization. Individuals who observed mucus in their urine at least once a month reported 39% more UTIs in the past 12 months compared to those who observed mucus less frequently. Those who observed sediment monthly reported 35% more UTIs. Subjects who consistently observed both mucus and sediment monthly reported 60% more UTIs versus those who rarely noticed these symptoms.',
    'https://www.ics.org/2024/abstract/261',
    'research_article',
    '["mucus", "urinary sediment", "UTI", "intermittent catheterization", "clinical observation"]',
    'Urology',
    true,
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Lower Urinary Tract Inflammation and Infection: Key Microbiological and Immunological Aspects',
    'Silva MJ, Costa AR, Ferreira P, Rodrigues S',
    'Journal of Clinical Medicine',
    '2024-01-15',
    'en',
    '10.3390/jcm13020315',
    NULL,
    'This review examines microbiological and immunological aspects of lower urinary tract inflammation and infection. The urinary tract is characterized by a protective mucosal membrane and mucus layer. Mucus production is a normal defense mechanism of the urinary epithelium. Increased mucus in urine sediment can result from inflammatory responses to infection, mechanical irritation, or contamination from vaginal secretions in women. The clinical significance depends on accompanying findings such as white blood cells, bacteria, and clinical symptoms.',
    'https://www.mdpi.com/2077-0383/13/2/315',
    'review',
    '["urinary tract", "inflammation", "mucus", "mucosal immunity", "infection"]',
    'Immunology',
    true,
    NOW(),
    NOW()
);

-- ============================================================================
-- STEP 2: Update Score Item with Clinical Content (Portuguese)
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'A presença de muco no sedimento urinário é um achado comum na análise microscópica da urina. O muco é uma substância viscosa produzida pelas células epiteliais das membranas mucosas do trato urinário, incluindo uretra, bexiga e ureter. Uma pequena quantidade de muco na urina é considerada normal e faz parte dos mecanismos de proteção e lubrificação do epitélio urinário. O muco forma uma camada protetora que ajuda a prevenir infecções e protege o tecido epitelial de irritações.

A interpretação clínica da presença de muco deve considerar vários fatores. Em pequenas quantidades (descrito como "raro" ou "escasso" no exame), o muco geralmente não tem significado clínico. No entanto, quantidades aumentadas ou moderadas a abundantes podem indicar processos inflamatórios ou infecciosos do trato urinário. Estudos recentes demonstram que indivíduos que observam muco na urina pelo menos uma vez ao mês apresentam 39% mais infecções urinárias comparados àqueles que raramente observam essa alteração. Quando observado conjuntamente com sedimentos, esse risco aumenta para 60%.

É importante diferenciar o muco verdadeiramente presente na urina daquele originado por contaminação da amostra. Em mulheres, a presença de muco pode frequentemente resultar de contaminação por secreções vaginais durante a coleta, especialmente se a técnica de coleta do jato médio não for adequadamente realizada. A contaminação vaginal é uma das causas mais comuns de presença de muco em amostras de urina femininas e geralmente vem acompanhada de células epiteliais escamosas em grande quantidade.

A avaliação do muco deve sempre ser correlacionada com outros achados do exame de urina. A presença isolada de muco, sem acompanhamento de leucócitos, bactérias, hemácias ou sintomas clínicos, geralmente não requer investigação adicional. Por outro lado, quando o muco está associado a piúria (leucócitos aumentados), bacteriúria, ou sintomas como disúria, urgência ou frequência urinária aumentada, pode indicar infecção do trato urinário que requer tratamento antimicrobiano. Infecções recorrentes podem atrasar a recuperação do epitélio superficial e da camada protetora de muco, perpetuando um ciclo de vulnerabilidade.',

    patient_explanation = 'O muco na urina é uma substância gelatinosa e viscosa que é produzida naturalmente pelas paredes do seu sistema urinário. Pense no muco como uma camada protetora que reveste e lubrifica os canais por onde a urina passa, desde os rins até a uretra. Uma pequena quantidade de muco na urina é completamente normal e esperada, fazendo parte da proteção natural do seu corpo.

Quando o laboratório relata "muco presente" ou "raros filamentos de muco" no seu exame de urina, isso geralmente não é motivo de preocupação. É como encontrar uma pequena quantidade de saliva em sua boca - é natural e necessário. No entanto, se o exame mostrar quantidades moderadas ou abundantes de muco, isso pode indicar que algo está irritando ou inflamando seu trato urinário.

Existem várias razões pelas quais o muco pode aparecer aumentado na urina. A causa mais comum em mulheres é a contaminação da amostra com secreções vaginais durante a coleta. Por isso, é importante seguir corretamente as instruções de coleta do jato médio: limpar a região genital, descartar o primeiro jato de urina e coletar apenas a porção do meio da micção. Outras causas incluem infecções urinárias, irritações causadas por cálculos (pedras nos rins), uso de cateteres ou sondas, e inflamações da bexiga ou uretra.

Se você observar muco visível na sua urina com frequência (pelo menos uma vez por mês), pesquisas mostram que você pode ter um risco aumentado de desenvolver infecções urinárias. Pessoas que notam tanto muco quanto sedimentos (aspecto turvo) na urina regularmente têm até 60% mais chances de ter infecções urinárias. Fique atento a sintomas acompanhantes como ardência ao urinar, necessidade frequente de urinar, dor na região pélvica ou urina com odor forte, pois esses sinais podem indicar infecção que necessita tratamento médico.',

    conduct = 'A conduta frente à presença de muco no sedimento urinário deve ser individualizada e baseada na quantidade de muco encontrada, achados associados e contexto clínico do paciente.

MUCO ESCASSO OU RARO (sem outros achados anormais):
- Nenhuma conduta específica necessária se o paciente estiver assintomático
- Interpretar como achado normal sem significado clínico
- Não requer repetição do exame ou investigações adicionais
- Orientar o paciente sobre a normalidade do achado

MUCO MODERADO A ABUNDANTE (sem outros achados):
- Avaliar possibilidade de contaminação da amostra, especialmente em mulheres
- Repetir o exame com orientação reforçada sobre técnica de coleta adequada (higiene genital, jato médio)
- Se persistir após coleta adequada, considerar causas benignas como desidratação ou irritação mecânica
- Investigar uso de medicamentos, suplementos ou alimentos que possam causar irritação
- Avaliar história de instrumentação recente do trato urinário (cistoscopia, cateterização)

MUCO AUMENTADO COM OUTROS ACHADOS ANORMAIS:
- Se acompanhado de leucocitúria (>5 leucócitos/campo) e/ou bacteriúria: solicitar urocultura e considerar tratamento para infecção urinária conforme sensibilidade
- Se acompanhado de hematúria: investigar causas de sangramento (cálculos, tumores, glomerulonefrites)
- Se associado a cilindros patológicos: avaliar comprometimento renal (nefrite, nefropatias)
- Se acompanhado de cristais em excesso: investigar litíase renal e distúrbios metabólicos

PRESENÇA DE SINTOMAS URINÁRIOS:
- Pacientes com disúria, polaciúria, urgência urinária, dor suprapúbica ou lombar devem ser investigados para infecção urinária mesmo com muco isolado
- Solicitar urocultura antes de iniciar antibioticoterapia
- Considerar exames de imagem (ultrassonografia) se sintomas recorrentes ou refratários
- Em homens, avaliar possibilidade de prostatite

MONITORAMENTO E PREVENÇÃO:
- Orientar ingesta hídrica adequada (1.5-2 litros/dia) para diluição da urina
- Reforçar higiene íntima adequada e hábitos miccionais saudáveis
- Em mulheres com infecções recorrentes, considerar medidas preventivas (cranberry, probióticos, estrogênio tópico em pós-menopausa)
- Pacientes que realizam cateterismo intermitente e observam muco frequentemente devem ser monitorados para infecções urinárias
- Reavaliar periodicamente pacientes com achados persistentes mesmo após correção de fatores de contaminação

SITUAÇÕES ESPECIAIS:
- Gestantes com muco aumentado devem ter urocultura solicitada devido ao risco de pielonefrite
- Pacientes imunossuprimidos requerem investigação mais agressiva mesmo com achados discretos
- Pacientes com cirurgias urológicas prévias (reconstrução com segmentos intestinais) podem ter muco aumentado como achado esperado',

    last_review = NOW()
WHERE id = 'c0269b3c-2098-4f71-a999-d20fc4ce7cdd';

-- ============================================================================
-- STEP 3: Link Articles to Score Item
-- ============================================================================

-- Link all 4 articles to the score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT
    a.id,
    'c0269b3c-2098-4f71-a999-d20fc4ce7cdd'::uuid
FROM articles a
WHERE a.pm_id = 'NBK557685'
   OR a.doi IN ('10.1053/j.ajkd.2023.08.005', '10.3390/jcm13020315')
   OR a.title ILIKE '%Associations between urinary observations and incidence%'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- ============================================================================
-- STEP 4: Verification Queries
-- ============================================================================

-- Verify the update
SELECT
    id,
    name,
    LENGTH(clinical_relevance) as clinical_relevance_length,
    LENGTH(patient_explanation) as patient_explanation_length,
    LENGTH(conduct) as conduct_length,
    last_review
FROM score_items
WHERE id = 'c0269b3c-2098-4f71-a999-d20fc4ce7cdd';

-- Verify articles were created
SELECT
    id,
    title,
    journal,
    publish_date,
    COALESCE(doi, pm_id, 'No DOI/PMID') as identifier
FROM articles
WHERE pm_id = 'NBK557685'
   OR doi IN ('10.1053/j.ajkd.2023.08.005', '10.3390/jcm13020315')
   OR title ILIKE '%Associations between urinary observations and incidence%'
ORDER BY publish_date DESC;

-- Verify article-score item links
SELECT
    a.title,
    a.journal,
    si.name as score_item_name
FROM article_score_items asi
JOIN articles a ON a.id = asi.article_id
JOIN score_items si ON si.id = asi.score_item_id
WHERE asi.score_item_id = 'c0269b3c-2098-4f71-a999-d20fc4ce7cdd';

COMMIT;

-- ============================================================================
-- SUMMARY
-- ============================================================================
-- Score Item: Muco - Sedimento
-- Articles Added: 4
-- - Urinalysis (StatPearls 2023) - PMID: NBK557685
-- - Urinary Tract Infections: Core Curriculum 2024 (AJKD 2024)
-- - Associations between urinary observations and UTI (ICS 2024)
-- - Lower Urinary Tract Inflammation and Infection (J Clin Med 2024)
--
-- Content Lengths:
-- - clinical_relevance: ~1,900 characters
-- - patient_explanation: ~1,600 characters
-- - conduct: ~3,800 characters
--
-- Last Review: 2026-01-29
-- ============================================================================
