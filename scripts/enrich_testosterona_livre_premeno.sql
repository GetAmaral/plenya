-- Enriquecimento: Testosterona Livre - Mulheres Pré-Menopausa
-- Score Item ID: bb8e93f4-97f3-45de-8446-e138235953f0
-- Data: 2026-01-29

BEGIN;

-- =============================================================================
-- PARTE 1: INSERIR ARTIGOS CIENTÍFICOS
-- =============================================================================

-- Artigo 1: Hyperandrogenism and microRNAs in PCOS (Frontiers Med 2019)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type,
    abstract,
    created_at,
    updated_at
)
SELECT
    gen_random_uuid(),
    'Hyperandrogenism and Metabolic Syndrome Are Associated With Changes in Serum-Derived microRNAs in Women With Polycystic Ovary Syndrome',
    'Sørensen AE, Udesen PB, Maciag G, Geiger J, Saliani N, et al.',
    'Front Med (Lausanne)',
    '2019-11-01'::date,
    '31737638',
    'research_article',
    'Estudo investigando 42 mulheres com SOP estratificadas por testosterona livre sérica (≥0.034 nmol/L) em grupos normoandrêmicas e hiperandrêmicas. Demonstrou que mulheres com SOP hiperandrêmica apresentam maior resistência insulínica e prevalência de síndrome metabólica. Identificou perfil específico de microRNAs circulantes associados à testosterona livre elevada, sugerindo biomarcadores para complicações metabólicas.',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '31737638'
);

-- Artigo 2: Practical Approach to Hyperandrogenism (Med Clin North Am 2021)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type,
    abstract,
    created_at,
    updated_at
)
SELECT
    gen_random_uuid(),
    'Practical Approach to Hyperandrogenism in Women',
    'Sharma A, Welt CK',
    'Med Clin North Am',
    '2021-11-01'::date,
    '34688417',
    'review',
    'Revisão prática sobre avaliação do hiperandrogenismo em mulheres, enfatizando que testosterona total medida por LC-MS/MS é o método mais confiável, embora imunoensaios tradicionais frequentemente falhem em detectar elevações leves típicas de SOP. Discute diferenciação clínica por fase da vida e estratégias de manejo, priorizando anticoncepcionais orais como terapia de primeira linha para hirsutismo.',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '34688417'
);

-- Artigo 3: Reference ranges testosterone premenopausal (Clin Biochem 2012)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type,
    abstract,
    created_at,
    updated_at
)
SELECT
    gen_random_uuid(),
    'Reference ranges for total and calculated free and bioavailable testosterone in a young healthy women population with normal menstrual cycles or using oral contraception',
    'Pesant MH, Desmarais G, Fink GD, Baillargeon JP',
    'Clin Biochem',
    '2012-01-01'::date,
    '22019954',
    'research_article',
    'Estabeleceu valores de referência para testosterona livre calculada (3-39 pmol/L) e biodisponível (0.06-0.81 nmol/L) em 155 mulheres jovens saudáveis de 18-40 anos com ciclos menstruais normais ou usando contraceptivos orais. Enfatiza a necessidade de valores de referência estabelecidos a partir de população bem caracterizada para interpretação adequada dos níveis androgênicos.',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '22019954'
);

-- Artigo 4: Society for Endocrinology Guideline 2025
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type,
    abstract,
    created_at,
    updated_at
)
SELECT
    gen_random_uuid(),
    'Society for Endocrinology Clinical Practice Guideline for the Evaluation of Androgen Excess in Women',
    'Elhassan YS, Cortés I, Jayasena CN, et al.',
    'Clin Endocrinol (Oxf)',
    '2025-01-01'::date,
    '39757890',
    'clinical_trial',
    'Diretriz clínica da Sociedade de Endocrinologia recomendando avaliação de índice de testosterona livre em mulheres investigadas para hiperandrogenismo, particularmente se testosterona total é normal com excesso androgênico clínico. Resultados devem ser interpretados contra valores de referência baseados em evidência, categorizados por faixa etária específica e/ou status pré e pós-menopausa.',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '39757890'
);

-- =============================================================================
-- PARTE 2: VINCULAR ARTIGOS AO SCORE ITEM
-- =============================================================================

-- Vincular Artigo 1
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'bb8e93f4-97f3-45de-8446-e138235953f0'::uuid
FROM articles a
WHERE a.pm_id = '31737638'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Vincular Artigo 2
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'bb8e93f4-97f3-45de-8446-e138235953f0'::uuid
FROM articles a
WHERE a.pm_id = '34688417'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Vincular Artigo 3
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'bb8e93f4-97f3-45de-8446-e138235953f0'::uuid
FROM articles a
WHERE a.pm_id = '22019954'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Vincular Artigo 4
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'bb8e93f4-97f3-45de-8446-e138235953f0'::uuid
FROM articles a
WHERE a.pm_id = '39757890'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- =============================================================================
-- PARTE 3: ATUALIZAR CONTEÚDO CLÍNICO DO SCORE ITEM
-- =============================================================================

UPDATE score_items
SET
    clinical_relevance = 'A testosterona livre representa a fração biologicamente ativa do hormônio não ligada à SHBG (globulina ligadora de hormônios sexuais) ou albumina, sendo superior à testosterona total para diagnóstico de hiperandrogenismo em mulheres pré-menopáusicas. Valores de referência variam de 0.3-1.9 pg/mL (1.0-6.6 pmol/L), com faixa até 39 pmol/L considerada normal em mulheres sem hiperandrogenismo. A medição direta por diálise de equilíbrio com LC-MS/MS é padrão-ouro, porém a testosterona livre calculada pela fórmula de Vermeulen (utilizando testosterona total e SHBG) é alternativa validada e mais acessível, especialmente quando SHBG > 30 nmol/L. A Sociedade de Endocrinologia recomenda avaliação de testosterona livre quando testosterona total está normal mas há excesso androgênico clínico, particularmente na Síndrome dos Ovários Policísticos (SOP), que afeta 5-20% das mulheres em idade reprodutiva. Mulheres com SOP hiperandrêmica (testosterona livre ≥0.034 nmol/L ou ≥1.18 pg/mL) apresentam 4.42 vezes maior risco de resistência insulínica comparado a mulheres com níveis normais, além de maior prevalência de síndrome metabólica, dislipidemia e risco cardiovascular. Testosterona livre também varia durante ciclo menstrual, com pico no meio do ciclo (fase ovulatória) cerca de 15.6 pg/mL, e declínio gradual na fase lútea, necessitando interpretação com referências específicas para fase menstrual. Condições que alteram SHBG (obesidade reduz, uso de anticoncepcionais orais aumenta) modificam proporção de testosterona livre, tornando sua medição mais específica que testosterona total nestas situações.',

    patient_explanation = 'A testosterona livre é a parcela do hormônio testosterona que circula "solta" no sangue, disponível para agir nas células do corpo. Embora seja considerado um "hormônio masculino", a testosterona é produzida normalmente pelos ovários e glândulas suprarrenais das mulheres em pequenas quantidades (cerca de 10-20 vezes menos que homens). Valores normais variam entre 0.3-1.9 pg/mL, podendo oscilar durante o ciclo menstrual com pico no meio do ciclo (período ovulatório). Níveis elevados de testosterona livre podem causar sintomas como crescimento excessivo de pelos em padrão masculino (face, tórax, abdômen), acne persistente, queda de cabelo no topo da cabeça, irregularidade menstrual, dificuldade para engravidar, ganho de peso e resistência à insulina. A causa mais comum de testosterona livre elevada é a Síndrome dos Ovários Policísticos (SOP), condição que afeta 1 em cada 10 mulheres. Este exame é especialmente importante quando há sintomas de excesso hormonal mas a testosterona total aparece normal, pois pode revelar alterações que outros exames não detectam. A testosterona livre também pode estar alterada por uso de anticoncepcionais, obesidade, síndrome metabólica ou em raros casos por tumores produtores de hormônios. O médico solicitará este exame junto com outros hormônios (SHBG, LH, FSH) para avaliar melhor sua saúde hormonal e função ovariana. Valores normais indicam boa função hormonal, enquanto alterações requerem investigação e acompanhamento especializado.',

    conduct = 'INVESTIGAÇÃO DIAGNÓSTICA: Solicitar testosterona livre em mulheres com sinais clínicos de hiperandrogenismo (hirsutismo, acne refratária, alopecia androgenética, irregularidade menstrual) especialmente quando testosterona total está normal. Coletar preferencialmente na fase folicular precoce (dias 3-5 do ciclo) pela manhã (8-10h) por variação circadiana, embora valores de referência específicos para fase menstrual sejam ideais. Preferir testosterona livre calculada pela fórmula de Vermeulen (requer testosterona total por LC-MS/MS + SHBG) se diálise de equilíbrio não disponível. Solicitar painel completo: testosterona total e livre, SHBG, DHEA-S, androstenediona, 17-hidroxiprogesterona (excluir hiperplasia adrenal congênita), LH/FSH (razão LH:FSH elevada sugere SOP), prolactina, TSH, glicemia jejum e HbA1c (avaliar resistência insulínica). Em casos de testosterona livre elevada: realizar ultrassom transvaginal para morfologia ovariana (critérios Rotterdam para SOP: ≥12 folículos 2-9mm ou volume ovariano >10mL), avaliar síndrome metabólica (circunferência abdominal, pressão arterial, lipidograma, glicemia), calcular HOMA-IR. Se testosterona livre >2-3x limite superior normal ou progressão rápida (<6 meses), investigar tumor secretor (ultrassom/RNM pélvica e adrenal, TC abdômen). INTERPRETAÇÃO: Valores 0.3-1.9 pg/mL (1.0-6.6 pmol/L) normais; 2.0-3.5 pg/mL leve elevação (SOP provável); >4.0 pg/mL significativo (investigar tumor se >8.0 pg/mL). Considerar influência de SHBG: obesidade reduz SHBG (aumenta testosterona livre desproporcionalmente), anticoncepcionais orais aumentam SHBG (reduzem testosterona livre). CONDUTA TERAPÊUTICA: SOP confirmada sem desejo gestacional imediato: anticoncepcionais orais combinados primeira linha (etinilestradiol 20-35mcg + progestágeno antiandrogênico como drospirenona, ciproterona, dienogeste) reduzem testosterona livre por suprimir LH e aumentar SHBG. Se hirsutismo persiste após 6 meses: adicionar espironolactona 50-200mg/dia (antiandrogênico, monitorar potássio) ou finasterida 2.5-5mg/dia. Se resistência insulínica ou síndrome metabólica: metformina 1500-2000mg/dia melhora sensibilidade insulínica e pode reduzir testosterona livre. SOP com infertilidade: encaminhar reprodução assistida (citrato clomifeno, letrozol, gonadotrofinas). Modificações estilo vida fundamentais: perda ponderal 5-10% melhora hiperandrogenismo, exercício resistido e aeróbico, dieta baixo índice glicêmico. SEGUIMENTO: Reavaliar testosterona livre, glicemia e lipidograma após 3-6 meses de tratamento. Monitorar anualmente função metabólica (síndrome metabólica em 40% das SOP), risco cardiovascular e saúde reprodutiva. Valores persistentemente elevados sem resposta terapêutica requerem reavaliação diagnóstica.',

    last_review = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP

WHERE id = 'bb8e93f4-97f3-45de-8446-e138235953f0';

-- =============================================================================
-- VERIFICAÇÃO FINAL
-- =============================================================================

-- Verificar artigos inseridos
SELECT
    a.pm_id,
    a.title,
    a.journal,
    a.publish_date,
    a.article_type,
    CASE
        WHEN asi.score_item_id IS NOT NULL THEN 'Vinculado'
        ELSE 'NÃO vinculado'
    END as status_vinculo
FROM articles a
LEFT JOIN article_score_items asi
    ON a.id = asi.article_id
    AND asi.score_item_id = 'bb8e93f4-97f3-45de-8446-e138235953f0'::uuid
WHERE a.pm_id IN ('31737638', '34688417', '22019954', '39757890')
ORDER BY a.publish_date DESC;

-- Verificar tamanhos dos campos
SELECT
    'clinical_relevance' as campo,
    LENGTH(clinical_relevance) as caracteres,
    CASE
        WHEN LENGTH(clinical_relevance) BETWEEN 1500 AND 2000 THEN '✓ OK'
        ELSE '✗ Fora do range 1500-2000'
    END as validacao
FROM score_items
WHERE id = 'bb8e93f4-97f3-45de-8446-e138235953f0'

UNION ALL

SELECT
    'patient_explanation' as campo,
    LENGTH(patient_explanation) as caracteres,
    CASE
        WHEN LENGTH(patient_explanation) BETWEEN 1000 AND 1500 THEN '✓ OK'
        ELSE '✗ Fora do range 1000-1500'
    END as validacao
FROM score_items
WHERE id = 'bb8e93f4-97f3-45de-8446-e138235953f0'

UNION ALL

SELECT
    'conduct' as campo,
    LENGTH(conduct) as caracteres,
    CASE
        WHEN LENGTH(conduct) BETWEEN 1500 AND 2500 THEN '✓ OK'
        ELSE '✗ Fora do range 1500-2500'
    END as validacao
FROM score_items
WHERE id = 'bb8e93f4-97f3-45de-8446-e138235953f0';

-- Verificar total de artigos vinculados
SELECT
    COUNT(*) as total_artigos_vinculados
FROM article_score_items
WHERE score_item_id = 'bb8e93f4-97f3-45de-8446-e138235953f0';

COMMIT;
