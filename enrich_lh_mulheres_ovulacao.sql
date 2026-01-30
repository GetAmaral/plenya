-- ============================================================================
-- ENRICHMENT: LH - Mulheres Ovulação
-- Score Item ID: 7dc53f7b-5ae3-49cb-ad5c-1e1afd3fcbe8
-- Date: 2026-01-29
-- ============================================================================

BEGIN;

-- ----------------------------------------------------------------------------
-- 1. INSERT SCIENTIFIC ARTICLES
-- ----------------------------------------------------------------------------

-- Article 1: Systematic Review and Meta-Analysis (2022)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type
)
SELECT
    gen_random_uuid(),
    'The LH surge and ovulation re-visited: a systematic review and meta-analysis and implications for true natural cycle frozen thawed embryo transfer',
    'Erden M, Mumusoglu S, Polat M, Ozbek IY, Esteves SC, Humaidan P, Yarali H',
    'Human Reproduction Update',
    '2022-03-01'::date,
    '35258085',
    'meta_analysis'
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '35258085'
);

-- Article 2: Research Article - Novel LH Detection (2022)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type
)
SELECT
    gen_random_uuid(),
    'Identification of the LH surge by measuring intact and total immunoreactivity in urine for prediction of ovulation time',
    'Demir A, Hero M, Alfthan H, Passioni A, Tapanainen JS, Stenman UH',
    'Hormones (Athens)',
    '2022-05-26'::date,
    '35614178',
    'research_article'
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '35614178'
);

-- Article 3: Cochrane Systematic Review (2023)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type
)
SELECT
    gen_random_uuid(),
    'Timed intercourse for couples trying to conceive',
    'Gibbons T, Reavey J, Georgiou EX, Becker CM',
    'Cochrane Database of Systematic Reviews',
    '2023-09-15'::date,
    '37709293',
    'review'
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '37709293'
);

-- Article 4: Retrospective Study - IUI Timing (2022)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type
)
SELECT
    gen_random_uuid(),
    'The Effect of Spontaneous LH Surges on Pregnancy Outcomes in Patients Undergoing Letrozole-HMG IUI: A Retrospective Analysis of 6,285 Cycles',
    'Jiang S, Chen L, Gao Y, Xi Q, Li W, Zhao X, Kuang Y',
    'Frontiers in Endocrinology',
    '2022-05-01'::date,
    '35600574',
    'research_article'
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '35600574'
);

-- ----------------------------------------------------------------------------
-- 2. LINK ARTICLES TO SCORE ITEM
-- ----------------------------------------------------------------------------

INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '7dc53f7b-5ae3-49cb-ad5c-1e1afd3fcbe8'::uuid
FROM articles a
WHERE a.pm_id IN ('35258085', '35614178', '37709293', '35600574')
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- ----------------------------------------------------------------------------
-- 3. UPDATE SCORE ITEM WITH CLINICAL CONTENT
-- ----------------------------------------------------------------------------

UPDATE score_items
SET
    clinical_relevance = 'O hormônio luteinizante (LH) desempenha papel crucial na ovulação feminina, com seu pico sérico (20-80 mIU/mL) ocorrendo 24-36 horas antes da liberação do óvulo. A detecção precisa do surge de LH é fundamental para o planejamento de concepção natural e em reprodução assistida. Meta-análise recente (Erden et al., 2022, Human Reproduction Update) revisou sistematicamente evidências sobre timing do LH surge e suas implicações clínicas, demonstrando que a definição precisa do início do surge deve ser padronizada em estudos futuros. Estudos contemporâneos revelam que os kits comerciais de predição de ovulação (OPKs) detectam apenas LH intacto, perdendo formas não-intactas que podem definir janela de fertilidade mais ampla (Demir et al., 2022, Hormones). Revisão Cochrane com 2.464 participantes (Gibbons et al., 2023) demonstrou que testes urinários de LH provavelmente aumentam taxas de nascidos vivos (RR 1,36; IC95% 1,02-1,81) em mulheres abaixo de 40 anos tentando conceber há menos de 12 meses. Em reprodução assistida, análise de 6.285 ciclos de IUI (Jiang et al., 2022, Frontiers in Endocrinology) mostrou que surge espontâneo de LH em ciclos estimulados não melhora resultados, mas o timing correto permanece crítico. A fecundabilidade é máxima quando relações sexuais ocorrem no intervalo de 3 dias terminando no dia da ovulação, com pico de 42% um dia antes da ovulação versus 20% no dia da ovulação.',

    patient_explanation = 'O LH (hormônio luteinizante) é produzido pela hipófise e provoca a liberação do óvulo maduro pelo ovário - processo chamado ovulação. Cerca de 24-36 horas antes da ovulação, ocorre um aumento súbito do LH no sangue e na urina, chamado de "pico de LH" ou "LH surge". Este pico pode ser detectado por testes de farmácia (testes de ovulação) que identificam LH na urina. Estudos científicos recentes mostram que usar esses testes aumenta as chances de gravidez em mulheres com menos de 40 anos que estão tentando engravidar há menos de 1 ano. O melhor momento para relações sexuais é nos 3 dias que terminam no dia da ovulação, sendo que o dia anterior à ovulação oferece 42% de chance de concepção. Para quem está tentando engravidar, fazer o teste de LH diariamente (geralmente à tarde/noite) ajuda a identificar a janela fértil. Quando o teste fica positivo, significa que a ovulação deve ocorrer nas próximas 24-36 horas. Em tratamentos de fertilidade como inseminação artificial, os médicos usam o timing do pico de LH para programar o procedimento no momento ideal, geralmente 24-40 horas após o surge de LH, aumentando significativamente as taxas de sucesso.',

    conduct = 'INTERPRETAÇÃO CLÍNICA: LH basal em fase folicular: 2-12 mIU/mL; LH surge pré-ovulatório: 20-80 mIU/mL; pós-menopausa: >30 mIU/mL. Valores abaixo de 20 mIU/mL em meio do ciclo sugerem anovulação, enquanto elevação persistente pode indicar síndrome dos ovários policísticos (SOP) ou falência ovariana prematura. MONITORAMENTO DE FERTILIDADE: Recomendar testes urinários de LH em mulheres com ciclos regulares tentando conceber. Orientar testagem diária no período peri-ovulatório (geralmente dia 10-14 em ciclo de 28 dias), preferencialmente à tarde. Revisar Cochrane (Gibbons et al., 2023) demonstrou RR 1,36 para nascidos vivos com uso de OPKs versus relações não programadas. Explicar que surge positivo indica janela de 24-36h até ovulação, sendo ideal ter relações no dia do pico e dia seguinte. REPRODUÇÃO ASSISTIDA: Em ciclos de IUI, agendar procedimento 24-40 horas após surge de LH (evidência de OR 3,0 para gravidez com timing adequado). Análise de 6.285 ciclos (Jiang et al., 2022) mostrou que surge espontâneo em ciclos estimulados não requer alteração de protocolo, mas timing preciso permanece crucial. Considerar que modelos combinando LH, estradiol e progesterona com ultrassom alcançam 93-97% de acurácia preditiva versus LH isolado. INVESTIGAÇÃO DE ANOVULAÇÃO: LH persistentemente elevado com FSH normal sugere SOP; LH elevado com FSH >25 mIU/mL indica insuficiência ovariana. Solicitar LH e FSH basais (dia 2-3 do ciclo) e considerar dosagem seriada mid-cycle se suspeita de anovulação. LIMITAÇÕES: Meta-análise (Erden et al., 2022) enfatiza que surge de LH não garante ovulação - síndrome do folículo luteinizado não roto (LUF) pode ocorrer. Confirmar ovulação com progesterona sérica em fase lútea (>10 ng/mL) ou ultrassom seriado quando indicado.',

    updated_at = NOW()
WHERE id = '7dc53f7b-5ae3-49cb-ad5c-1e1afd3fcbe8';

COMMIT;

-- ============================================================================
-- VERIFICATION QUERIES
-- ============================================================================

-- Check character counts
SELECT
    'clinical_relevance' as field,
    LENGTH(clinical_relevance) as char_count,
    CASE
        WHEN LENGTH(clinical_relevance) BETWEEN 1500 AND 2000 THEN '✓ OK'
        ELSE '✗ OUT OF RANGE'
    END as status
FROM score_items
WHERE id = '7dc53f7b-5ae3-49cb-ad5c-1e1afd3fcbe8'

UNION ALL

SELECT
    'patient_explanation' as field,
    LENGTH(patient_explanation) as char_count,
    CASE
        WHEN LENGTH(patient_explanation) BETWEEN 1000 AND 1500 THEN '✓ OK'
        ELSE '✗ OUT OF RANGE'
    END as status
FROM score_items
WHERE id = '7dc53f7b-5ae3-49cb-ad5c-1e1afd3fcbe8'

UNION ALL

SELECT
    'conduct' as field,
    LENGTH(conduct) as char_count,
    CASE
        WHEN LENGTH(conduct) BETWEEN 1500 AND 2500 THEN '✓ OK'
        ELSE '✗ OUT OF RANGE'
    END as status
FROM score_items
WHERE id = '7dc53f7b-5ae3-49cb-ad5c-1e1afd3fcbe8';

-- Verify article links
SELECT
    si.name as score_item,
    COUNT(asi.article_id) as linked_articles,
    STRING_AGG(a.pm_id, ', ' ORDER BY a.publish_date DESC) as pmids
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
LEFT JOIN articles a ON asi.article_id = a.id
WHERE si.id = '7dc53f7b-5ae3-49cb-ad5c-1e1afd3fcbe8'
GROUP BY si.id, si.name;
