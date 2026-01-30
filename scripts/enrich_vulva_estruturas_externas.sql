-- Enriquecimento do item "Vulva e Estruturas Externas"
-- Score Item ID: 30a7764c-4c11-4078-b942-860ee56136a4
-- Categoria: Vida Sexual
-- Data: 2026-01-29

BEGIN;

-- =======================
-- ARTIGOS CIENTÍFICOS
-- =======================

-- Artigo 1: Gynecologic Pelvic Examination (StatPearls, 2024)
INSERT INTO articles (
    title,
    authors,
    journal,
    pm_id,
    doi,
    publish_date,
    article_type
) VALUES (
    'Gynecologic Pelvic Examination',
    'Mikes BA, Wray AA',
    'StatPearls',
    '30480956',
    NULL,
    '2024-02-01'::date,
    'review'
)
ON CONFLICT (id) DO NOTHING;

-- Artigo 2: Anatomy, histology, and nerve density of clitoris (AJOG, 2021)
INSERT INTO articles (
    title,
    authors,
    journal,
    pm_id,
    doi,
    publish_date,
    article_type
) VALUES (
    'Anatomy, histology, and nerve density of clitoris and associated structures: clinical applications to vulvar surgery',
    'Pin PG, Pin J',
    'American Journal of Obstetrics and Gynecology',
    '32888920',
    '10.1016/j.ajog.2020.08.112',
    '2021-01-01'::date,
    'research_article'
)
ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    pm_id = EXCLUDED.pm_id,
    publish_date = EXCLUDED.publish_date,
    article_type = EXCLUDED.article_type;

-- Artigo 3: 2021 European guideline for vulval conditions (J Eur Acad Dermatol Venereol, 2022)
INSERT INTO articles (
    title,
    authors,
    journal,
    pm_id,
    doi,
    publish_date,
    article_type
) VALUES (
    '2021 European guideline for the management of vulval conditions',
    'van der Meijden WI, Boffa MJ, Ter Harmsel B, Kirtschig G, Lewis F, Moyal-Barracco M, Tiplica GS, Sherrard J',
    'Journal of the European Academy of Dermatology and Venereology',
    '35411963',
    '10.1111/jdv.18102',
    '2022-07-01'::date,
    'review'
)
ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    pm_id = EXCLUDED.pm_id,
    publish_date = EXCLUDED.publish_date,
    article_type = EXCLUDED.article_type;

-- Artigo 4: Lichen sclerosus: The 2023 update (Front Med, 2023)
INSERT INTO articles (
    title,
    authors,
    journal,
    pm_id,
    doi,
    publish_date,
    article_type
) VALUES (
    'Lichen sclerosus: The 2023 update',
    'De Luca DA, Papara C, Vorobyev A, Staiger H, Bieber K, Thaçi D, Ludwig RJ',
    'Frontiers in Medicine',
    '36873861',
    '10.3389/fmed.2023.1106318',
    '2023-02-01'::date,
    'review'
)
ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    pm_id = EXCLUDED.pm_id,
    publish_date = EXCLUDED.publish_date,
    article_type = EXCLUDED.article_type;

-- =======================
-- VINCULAR ARTIGOS AO SCORE ITEM
-- =======================

-- Vincular artigo 1 (StatPearls - sem DOI)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '30a7764c-4c11-4078-b942-860ee56136a4'::uuid
FROM articles a
WHERE a.pm_id = '30480956' AND a.journal = 'StatPearls'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Vincular artigo 2
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '30a7764c-4c11-4078-b942-860ee56136a4'::uuid
FROM articles a
WHERE a.doi = '10.1016/j.ajog.2020.08.112'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Vincular artigo 3
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '30a7764c-4c11-4078-b942-860ee56136a4'::uuid
FROM articles a
WHERE a.doi = '10.1111/jdv.18102'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Vincular artigo 4
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '30a7764c-4c11-4078-b942-860ee56136a4'::uuid
FROM articles a
WHERE a.doi = '10.3389/fmed.2023.1106318'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- =======================
-- ATUALIZAR CONTEÚDO CLÍNICO
-- =======================

UPDATE score_items
SET
    clinical_relevance = 'A avaliação das estruturas vulvares externas é um componente fundamental do exame ginecológico completo, fornecendo informações cruciais sobre saúde genital, dermatológica e sexual. A vulva compreende estruturas anatômicas complexas incluindo monte púbico, grandes lábios, pequenos lábios, clitóris, vestíbulo vulvar, glândulas de Bartholin e Skene, e abertura vaginal. O exame sistemático permite identificar variações anatômicas normais, alterações fisiológicas relacionadas à idade e ciclo hormonal, e detectar precocemente condições patológicas. As dermatoses vulvares representam diagnósticos desafiadores que afetam significativamente a qualidade de vida, incluindo líquen escleroso (condição crônica autoimune que causa atrofia epitelial com placas esbranquiçadas), líquen plano (doença inflamatória mucocutânea), dermatite de contato, e líquen simples crônico. A vulvodínia, caracterizada por dor vulvar crônica sem lesão visível, frequentemente representa dor neuropática secundária a irritação ou lesão do sistema nervoso. O líquen escleroso merece atenção especial pois acomete aproximadamente 1:300 a 1:1000 mulheres e apresenta risco de 5% de evolução para carcinoma espinocelular, necessitando vigilância longitudinal. O exame tátil é essencial para avaliar firmeza, nodularidade ou alterações não visíveis à inspeção, sendo indicada biópsia quando há suspeita de neoplasia ou resposta inadequada ao tratamento. A correlação clínico-patológica é fundamental para diagnóstico preciso, e biópsias devem ser realizadas antes do início do tratamento pois mudanças patológicas características podem ser alteradas por terapias tópicas. A compreensão da anatomia clitoriana, incluindo densidade nervosa e histologia, é crucial para procedimentos cirúrgicos vulvares e preservação da função sexual. O conhecimento das medidas anatômicas normais da genitália externa feminina é necessário para desenho de equipamentos médico-cirúrgicos apropriados e avaliação de variações individuais. A avaliação vulvar deve considerar distribuição pilosa, simetria, coloração, presença de eritema, edema, equimoses, lesões, massas, secreções anormais e integridade da mucosa.',

    patient_explanation = 'A vulva é o conjunto de estruturas genitais externas femininas que incluem os lábios maiores e menores, o clitóris, a entrada da vagina e pequenas glândulas. O exame da vulva é uma parte importante da consulta ginecológica para avaliar sua saúde íntima. Durante o exame, o médico observa cuidadosamente a pele, mucosas, formato e características de cada estrutura. Este exame permite identificar variações normais entre mulheres, mudanças naturais que ocorrem com a idade ou ciclo menstrual, e detectar precocemente problemas como infecções, inflamações ou lesões cutâneas. Algumas condições comuns que afetam a vulva incluem dermatites (inflamações da pele por alergia ou irritação), líquen escleroso (condição que deixa a pele mais fina e esbranquiçada), e vulvodínia (dor crônica na região sem causa aparente). É normal que existam diferenças no tamanho, formato e cor das estruturas vulvares entre mulheres, e essas variações geralmente não indicam problemas. O médico também avalia se há coceira persistente, dor, feridas, manchas, caroços ou mudanças na textura da pele. Algumas alterações necessitam de biópsia (retirada de pequeno fragmento para análise) para confirmar o diagnóstico antes de iniciar tratamento. O exame não deve causar dor significativa, mas pode gerar desconforto emocional, sendo importante que seja realizado com respeito, privacidade e explicações claras. Manter boa higiene íntima com água e sabonete neutro, evitar produtos irritantes, usar roupas íntimas de algodão e consultar regularmente o ginecologista são cuidados importantes para manter a saúde vulvar. Qualquer sintoma persistente como coceira, ardor, dor ou mudanças visíveis deve ser comunicado ao médico.',

    conduct = 'PROTOCOLO DE AVALIAÇÃO VULVAR: O exame deve ser realizado em ambiente privativo, bem iluminado, com a paciente em posição ginecológica confortável. Iniciar com anamnese detalhada investigando sintomas (prurido, dor, ardor, dispareunia), tempo de evolução, tratamentos prévios, história menstrual, uso de medicamentos tópicos, produtos de higiene íntima, e histórico de doenças dermatológicas ou sexualmente transmissíveis. INSPEÇÃO VISUAL: Avaliar sistematicamente monte púbico (distribuição pilosa, presença de lesões), grandes lábios (simetria, coloração, edema, lesões), pequenos lábios (tamanho, pigmentação, simetria), clitóris e capuz clitoriano, intróito vaginal, meato uretral, glândulas de Bartholin bilateralmente, períneo e região perianal. Documentar presença de eritema, placas esbranquiçadas, equimoses, fissuras, escoriações, atrofia, hipertrofia, massas, secreções, cicatrizes ou assimetrias. PALPAÇÃO: Essencial para detectar nodularidades, firmeza aumentada ou alterações não visíveis, especialmente em casos de líquen escleroso com risco de transformação maligna. Palpar glândulas de Bartholin para detectar cistos ou abscessos. CRITÉRIOS PARA BIÓPSIA: Realizar em lesões com aspecto atípico, ulcerações persistentes, áreas de firmeza ou nodularidade à palpação, falha terapêutica após 8-12 semanas de tratamento adequado, e sempre que houver suspeita de neoplasia. A biópsia deve preceder o início do tratamento para não alterar características histopatológicas. DERMATOSES VULVARES: Para líquen escleroso, prescrever corticosteroides potentes (propionato de clobetasol 0,05%) aplicados diariamente por 4 semanas, seguido de redução gradual. Estabelecer seguimento de longo prazo com reavaliações periódicas devido ao risco de malignização. Para líquen plano, considerar corticosteroides tópicos de alta potência ou tacrolimus. Dermatite de contato requer identificação e eliminação do agente irritante/alérgeno, com corticosteroides de potência moderada. Líquen simples crônico necessita corticosteroides tópicos e abordagem do ciclo coceira-trauma. VULVODÍNIA: Abordagem multidisciplinar incluindo fisioterapia do assoalho pélvico, terapia cognitivo-comportamental, antidepressivos tricíclicos ou anticonvulsivantes para dor neuropática. Em casos refratários, considerar inibidores de JAK (upadacitinibe, abrocitinibe). EDUCAÇÃO: Orientar higiene íntima adequada (água morna e sabonete neutro sem fragrância), evitar duchas vaginais, produtos perfumados, roupas sintéticas apertadas. DOCUMENTAÇÃO: Registrar achados detalhadamente, idealmente com fotografias médicas quando apropriado e consentido. Utilizar escalas de qualidade de vida validadas (Vulvar Quality of Life Index) para monitoramento objetivo. SEGUIMENTO: Reavaliações programadas conforme gravidade da condição, com vigilância rigorosa em casos de líquen escleroso pela possibilidade de evolução neoplásica.',

    last_review = CURRENT_DATE
WHERE id = '30a7764c-4c11-4078-b942-860ee56136a4';

-- =======================
-- VERIFICAÇÃO
-- =======================

-- Verificar artigos inseridos
SELECT
    a.pm_id,
    a.title,
    a.journal,
    a.publish_date,
    a.article_type
FROM articles a
WHERE a.pm_id IN ('30480956', '32888920', '35411963', '36873861')
ORDER BY a.publish_date DESC;

-- Verificar vínculos criados
SELECT
    si.name as score_item,
    a.title as article_title,
    a.pm_id,
    a.publish_date
FROM article_score_items asi
JOIN score_items si ON asi.score_item_id = si.id
JOIN articles a ON asi.article_id = a.id
WHERE si.id = '30a7764c-4c11-4078-b942-860ee56136a4'
ORDER BY a.publish_date DESC;

-- Verificar tamanhos dos campos clínicos
SELECT
    name,
    LENGTH(clinical_relevance) as len_clinical_relevance,
    LENGTH(patient_explanation) as len_patient_explanation,
    LENGTH(conduct) as len_conduct,
    last_review
FROM score_items
WHERE id = '30a7764c-4c11-4078-b942-860ee56136a4';

COMMIT;

-- Exibir resumo final
SELECT
    '=== ENRIQUECIMENTO CONCLUÍDO ===' as status,
    COUNT(DISTINCT a.id) as total_artigos,
    COUNT(DISTINCT asi.score_item_id) as items_vinculados
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '30a7764c-4c11-4078-b942-860ee56136a4';
