-- Enrichment for Hemácias (RBC) - Sedimento Urinário
-- Score Item ID: 814d923f-cdfa-4388-9ba1-42b23dcd8d6d
-- Generated: 2026-01-29

BEGIN;

-- Insert scientific articles
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    original_link,
    article_type,
    language,
    specialty,
    created_at,
    updated_at
) VALUES
(
    gen_random_uuid(),
    'Glomerular Hematuria and the Utility of Urine Microscopy: A Review',
    'Saha MK, Massicotte-Azarniouch D, Reynolds ML, Mottl AK, Falk RJ, Jennette JC, Derebail VK',
    'American Journal of Kidney Diseases',
    '2022-09-01',
    '10.1053/j.ajkd.2022.02.022',
    'https://www.ajkd.org/article/S0272-6386(22)00584-4/fulltext',
    'review',
    'en',
    'Nephrology',
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Comparison of Urinary Red Blood Cell Distribution (URD) and Dysmorphic Red Blood Cells for Detecting Glomerular Hematuria: A Multicenter Study',
    'Lee HJ, Bang SH, Kim KH, Park JY, Shin JH, Kim YJ, Song W, Jeon CH',
    'Journal of Clinical Laboratory Analysis',
    '2025-02-03',
    '10.1002/jcla.25159',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC11904819/',
    'research_article',
    'en',
    'Laboratory Medicine',
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Updates to Microhematuria: AUA/SUFU Guideline (2025)',
    'Barocas DA, Lotan Y, Matulewicz RS, Raman JD, Westerman ME, Kirkby E, Pak LJ, Souter L',
    'Journal of Urology',
    '2025-05-01',
    '10.1097/JU.0000000000004490',
    'https://www.auajournals.org/doi/10.1097/JU.0000000000004490',
    'review',
    'en',
    'Urology',
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases',
    'Yang WS, Kim YJ, Park JY, Kim SB, Lee SK, Park JS',
    'Scientific Reports',
    '2021-10-25',
    '10.1038/s41598-021-00457-6',
    'https://www.nature.com/articles/s41598-021-00457-6',
    'research_article',
    'en',
    'Laboratory Medicine',
    NOW(),
    NOW()
);

-- Link articles to score item
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '814d923f-cdfa-4388-9ba1-42b23dcd8d6d',
    id
FROM articles
WHERE doi IN (
    '10.1053/j.ajkd.2022.02.022',
    '10.1002/jcla.25159',
    '10.1097/JU.0000000000004490',
    '10.1038/s41598-021-00457-6'
);

-- Update score item with clinical content
UPDATE score_items
SET
    clinical_relevance = 'A presença de hemácias no sedimento urinário é um achado clínico fundamental que requer investigação cuidadosa para determinar sua origem e significado. Valores de referência consideram normal até 3 hemácias/campo de alta potência (HPF), sendo a hematuria microscópica definida como ≥3 RBC/HPF. A distinção crítica entre hematúria glomerular e não-glomerular baseia-se na morfologia eritrocitária: hemácias dismórficas (especialmente acantócitos com projeções em forma de anel) indicam lesão glomerular com especificidade de 90-100%, enquanto hemácias isomórficas sugerem origem urológica (cálculos, tumores, infecções). A presença de ≥40% hemácias dismórficas ou ≥5% acantócitos em 2 de 3 amostras indica fortemente origem glomerular. Cilindros hemáticos são patognomônicos de glomerulonefrite. As diretrizes AUA/SUFU 2025 estratificam o risco de malignidade: pacientes de alto risco com hematúria microscópica apresentam 3,8% de incidência de neoplasia urotelial, versus 0,2% em baixo risco. Em doenças glomerulares proliferativas (nefropatia por IgA, vasculite ANCA, lúpus classe III-IV), a hematúria dismórfica é marcador diagnóstico essencial. Importante: analisadores automatizados subestimam a severidade da hematúria glomerular devido à fragilidade das hemácias dismórficas, tornando a microscopia manual com contraste de fase o padrão-ouro. Estudo multicêntrico 2025 demonstrou que URD (distribuição de tamanho eritrocitário) oferece performance diagnóstica comparável (AUC 0,83 vs 0,79) com maior objetividade.',

    patient_explanation = 'As hemácias (glóbulos vermelhos) no sedimento urinário indicam que há sangue na urina, mesmo que não seja visível a olho nu. O normal é ter menos de 3 hemácias por campo microscópico. Quando encontramos mais que isso, precisamos investigar de onde vem esse sangramento. O formato das hemácias na urina é muito importante: se estiverem "amassadas" ou deformadas (dismórficas), geralmente indica problema nos rins, especificamente nos filtros renais chamados glomérulos. Se estiverem com formato normal (isomórficas), pode indicar problemas nas vias urinárias como cálculos renais, infecções ou até tumores de bexiga. Nem sempre a presença de hemácias significa algo grave: exercícios intensos, menstruação, infecções urinárias simples ou traumas podem causar esse achado temporariamente. Porém, a persistência de hemácias na urina requer investigação mais aprofundada, especialmente em pessoas acima de 35 anos ou fumantes, pois pode indicar desde inflamações renais (glomerulonefrites) até câncer de bexiga. Seu médico pode solicitar exames complementares como ultrassom renal, cistoscopia (visualização da bexiga) ou testes específicos para avaliar a função dos rins. O importante é não ignorar esse achado e seguir as orientações médicas para investigação adequada.',

    conduct = 'INVESTIGAÇÃO INICIAL: Confirmar hematúria microscópica em 2-3 amostras de urina em diferentes ocasiões (evitar coleta durante menstruação, após exercício intenso ou relação sexual). Solicitar EAS completo com análise de sedimento por microscopia manual com contraste de fase (método padrão-ouro). Avaliar morfologia eritrocitária: calcular percentual de hemácias dismórficas e acantócitos. Investigar proteinúria associada, cilindros hemáticos e outras alterações sedimentares. ESTRATIFICAÇÃO DE RISCO (AUA/SUFU 2025): Alto risco: idade ≥60 anos, tabagismo ativo/prévio (≥30 maços-ano), ≥25 RBC/HPF, história de exposição ocupacional a químicos (benzenos, aminas aromáticas), doença urológica prévia, sintomas irritativos sem ITU, exposição à ciclofosfamida. Risco intermediário: idade 40-59 anos, 10-24 RBC/HPF, tabagismo <30 maços-ano. Baixo risco: idade <40 anos, 3-9 RBC/HPF, ausência de fatores de risco. CONDUTA POR ORIGEM: Hematúria glomerular (≥40% dismórficas ou ≥5% acantócitos ou cilindros hemáticos): encaminhar à nefrologia, solicitar creatinina sérica, TFG, relação proteína/creatinina urinária, complemento (C3/C4), FAN, ANCA, anti-DNA, sorologias (hepatites B/C, HIV), considerar biópsia renal se proteinúria >500mg/dia ou TFG reduzida. Hematúria não-glomerular: investigação urológica com ultrassom de vias urinárias, citologia urinária oncótica (se risco moderado/alto), cistoscopia para pacientes de risco intermediário/alto, TC de abdome sem contraste (rastreio de litíase) ou urotomografia (se suspeita de neoplasia). SEGUIMENTO: Baixo risco com hematúria resolvida: EAS em 6 e 12 meses. Persistência da hematúria: investigação urológica completa. Alto risco negativo na investigação inicial: vigilância anual com citologia e cistoscopia por 3 anos. Biomarcadores urinários (UroVysion, Cxbladder) podem ser considerados em casos selecionados.',

    updated_at = NOW()
WHERE id = '814d923f-cdfa-4388-9ba1-42b23dcd8d6d';

COMMIT;

-- Verification query
SELECT
    si.code,
    si.name,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    COUNT(asi.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '814d923f-cdfa-4388-9ba1-42b23dcd8d6d'
GROUP BY si.id, si.code, si.name, si.clinical_relevance, si.patient_explanation, si.conduct;
