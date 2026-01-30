-- Enriquecimento do Score Item: DHEA-S - Homens (40-49 anos)
-- ID: 66256f14-1775-4501-b85e-a4ffd71ccc7a
-- Data: 2026-01-28

BEGIN;

-- ============================================================================
-- INSERÇÃO DE ARTIGOS CIENTÍFICOS
-- ============================================================================

-- Artigo 1: Mortalidade e DHEA-S (2023)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    original_link,
    article_type,
    abstract,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Dehydroepiandrosterone Sulfate and Mortality in Middle-Aged and Older Men and Women—A J-shaped Relationship',
    'Mukama T, Johnson T, Katzke V, Kaaks R',
    'The Journal of Clinical Endocrinology & Metabolism',
    '2023-06-01',
    '10.1210/clinem/dgac716',
    'https://academic.oup.com/jcem/article/108/6/e313/6883509',
    'research_article',
    'Estudo prospectivo com 7.370 participantes (idade média 55 anos homens, 52,4 anos mulheres) demonstrando relação em J entre DHEA-S e mortalidade. Níveis muito baixos e muito altos de DHEA-S foram associados com maior mortalidade, sugerindo que existe uma faixa ótima de DHEA-S para longevidade em indivíduos de meia-idade e idosos.',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    abstract = EXCLUDED.abstract,
    updated_at = NOW();

-- Artigo 2: DHEA-S e Dislipidemia Diabética (2023)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    original_link,
    article_type,
    abstract,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Low serum dehydroepiandrosterone is associated with diabetic dyslipidemia risk in males with type 2 diabetes',
    'Chen S, Li S, Zhang X, Fan Y, Liu M',
    'Frontiers in Endocrinology',
    '2023-11-24',
    '10.3389/fendo.2023.1272797',
    'https://www.frontiersin.org/journals/endocrinology/articles/10.3389/fendo.2023.1272797/full',
    'research_article',
    'Estudo demonstrou associação entre baixos níveis séricos de DHEA e risco aumentado de dislipidemia diabética em homens com diabetes tipo 2. A deficiência de DHEA pode representar um fator de risco cardiovascular adicional em homens diabéticos de meia-idade, destacando a importância da monitorização deste hormônio nesta população.',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    abstract = EXCLUDED.abstract,
    updated_at = NOW();

-- Artigo 3: DHEA-S e Risco Cardiovascular (2020)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    original_link,
    article_type,
    abstract,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Plasma Dehydroepiandrosterone Sulfate and Cardiovascular Disease Risk in Older Men and Women',
    'Jia X, Sun C, Tang O, Gorlov I, Nambi V, Virani SS, Villareal DT, Taffet GE, Yu B, Bressler J, Boerwinkle E, Windham BG, de Lemos JA, Matsushita K, Selvin E, Michos ED, Hoogeveen RC, Ballantyne CM',
    'The Journal of Clinical Endocrinology & Metabolism',
    '2020-12-01',
    '10.1210/clinem/dgaa518',
    'https://academic.oup.com/jcem/article/105/12/e4304/5891765',
    'research_article',
    'Análise prospectiva de 8.143 indivíduos sem doença cardiovascular prévia (idade média 63 anos) do estudo ARIC. DHEA-S foi medido em plasma armazenado (1996-2013). Em homens, níveis extremamente baixos de DHEA-S (abaixo do percentil 10) mostraram tendência para aumento de risco cardiovascular, embora com diferenças sexo-específicas importantes.',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    abstract = EXCLUDED.abstract,
    updated_at = NOW();

-- Artigo 4: Revisão sobre Declínio de DHEA e Envelhecimento (2013)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    original_link,
    article_type,
    abstract,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'A Review of Age-Related Dehydroepiandrosterone Decline and Its Association with Well-Known Geriatric Syndromes: Is Treatment Beneficial?',
    'Samaras N, Samaras D, Frangos E, Forster A, Philippe J',
    'Rejuvenation Research',
    '2013-08-01',
    '10.1089/rej.2013.1425',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC3746247/',
    'review',
    'Revisão abrangente sobre o declínio de DHEA relacionado à idade e associações com síndromes geriátricas. Níveis de DHEA variam entre 1,33-7,78 ng/mL entre 18-40 anos, declinando para 10-20% desses valores aos 70-80 anos. Associações primárias incluem deterioração musculoesquelética (força muscular, densidade óssea), declínio cognitivo, transtornos de humor e risco cardiovascular.',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    abstract = EXCLUDED.abstract,
    updated_at = NOW();

-- ============================================================================
-- ATUALIZAÇÃO DO SCORE ITEM COM CONTEÚDO CLÍNICO
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'O DHEA-S (sulfato de desidroepiandrosterona) é o andrógeno adrenal mais abundante em homens de 40-49 anos, representando importante marcador de envelhecimento endócrino e reserva adrenal. Nesta faixa etária, os níveis de DHEA-S começam declínio acelerado (1-2% ao ano), podendo atingir apenas 50-60% dos valores observados aos 20-30 anos. Este hormônio possui múltiplas funções metabólicas: serve como precursor para testosterona e estradiol em tecidos periféricos, exerce efeitos neuroprotetores, modula função imune e metabolismo glicídico-lipídico. Em homens de meia-idade, níveis adequados de DHEA-S correlacionam-se com melhor composição corporal (menor adiposidade visceral, preservação de massa magra), função cognitiva preservada, perfil lipídico favorável e menor risco de síndrome metabólica. A relação entre DHEA-S e mortalidade cardiovascular apresenta padrão em forma de J: tanto níveis muito baixos quanto extremamente elevados associam-se a desfechos adversos. Deficiência de DHEA-S em homens 40-49 anos relaciona-se com fadiga crônica, sarcopenia precoce, disfunção erétil, redução de libido, resistência insulínica e maior risco de dislipidemia diabética. A monitorização de DHEA-S nesta faixa etária permite identificar precocemente declínio adrenal excessivo (adrenopausa acelerada), avaliar risco cardiometabólico e orientar intervenções lifestyle ou eventual suplementação em casos selecionados. Importante destacar que DHEA-S não substitui testosterona na avaliação androgênica masculina, mas complementa a investigação hormonal em quadros de hipogonadismo de início tardio (LOH) e síndrome metabólica.',

    patient_explanation = 'O DHEA-S é um hormônio produzido pelas suas glândulas adrenais (localizadas acima dos rins) que funciona como uma "matéria-prima" para produção de outros hormônios importantes, incluindo testosterona. Entre 40-49 anos, os níveis deste hormônio começam a diminuir naturalmente, o que faz parte do processo de envelhecimento. Quando os níveis de DHEA-S estão muito baixos nesta idade, você pode sentir cansaço persistente, dificuldade em manter ou ganhar músculos, diminuição do desejo sexual, acúmulo de gordura abdominal e redução da disposição física e mental. Níveis adequados de DHEA-S ajudam a manter energia, massa muscular, saúde cardiovascular e função cerebral. Este exame ajuda seu médico a entender se suas glândulas adrenais estão funcionando bem e se o declínio hormonal está acontecendo no ritmo esperado para sua idade. Valores muito baixos podem indicar necessidade de investigar estresse crônico excessivo, fadiga adrenal ou outros problemas hormonais. Em alguns casos selecionados, quando os níveis estão muito reduzidos e causam sintomas, seu médico pode considerar suplementação, sempre sob supervisão médica rigorosa. Importante: DHEA-S não é o mesmo que testosterona, mas os dois hormônios trabalham juntos para manter sua saúde hormonal masculina.',

    conduct = 'INTERPRETAÇÃO CLÍNICA EM HOMENS 40-49 ANOS: Valores de referência variam entre laboratórios, mas geralmente: Normal 280-640 μg/dL (7,6-17,4 μmol/L); Limítrofe 180-279 μg/dL; Baixo <180 μg/dL. COLETA: Realizar pela manhã (7-9h) devido ritmo circadiano, jejum não obrigatório mas preferencial. Suspender suplementação de DHEA se em uso (mínimo 48h). INVESTIGAÇÃO DE NÍVEIS BAIXOS: (1) Confirmar com segunda dosagem; (2) Avaliar cortisol sérico matinal e teste ACTH para excluir insuficiência adrenal primária ou secundária; (3) Investigar uso de corticosteroides, opioides ou anticonvulsivantes que suprimem eixo HPA; (4) Avaliar testosterona total e livre, LH, FSH para panorama androgênico completo; (5) Screening metabólico: glicemia, HbA1c, perfil lipídico, insulina, HOMA-IR; (6) Considerar avaliação de composição corporal (DEXA ou bioimpedância) se sarcopenia ou obesidade visceral. CONDUTA TERAPÊUTICA: Níveis normais: manter hábitos saudáveis, reavaliar anualmente. Níveis limítrofes: (1) Otimizar lifestyle: exercício resistido 3-4x/semana, sono 7-9h/noite, manejo de estresse (meditação, yoga), redução de álcool; (2) Corrigir deficiências nutricionais (zinco, magnésio, vitamina D); (3) Reavaliar em 6 meses. Níveis baixos sintomáticos: (1) Após exclusão de causas secundárias, considerar suplementação DHEA 25-50mg/dia pela manhã; (2) Monitorar PSA, hemograma, enzimas hepáticas antes e 3 meses após início; (3) Reavaliar DHEA-S, testosterona, estradiol em 3 meses (risco de conversão excessiva); (4) Manter dose mínima efetiva, não ultrapassar 50mg/dia; (5) Contraindicações: câncer de próstata, hiperplasia prostática severa, hepatopatia. SEGUIMENTO: Níveis muito elevados (>700 μg/dL) exigem investigação de tumor adrenal. Relação DHEA-S/cortisol <10 sugere insuficiência adrenal relativa. Atenção para sintomas de androgênio excessivo se suplementando: acne, oleosidade, agressividade, ginecomastia.',

    last_review = CURRENT_DATE,
    updated_at = NOW()
WHERE id = '66256f14-1775-4501-b85e-a4ffd71ccc7a';

-- ============================================================================
-- VINCULAÇÃO DE ARTIGOS AO SCORE ITEM
-- ============================================================================

-- Vincular artigo 1 (Mukama et al. 2023)
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '66256f14-1775-4501-b85e-a4ffd71ccc7a',
    id
FROM articles
WHERE doi = '10.1210/clinem/dgac716'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- Vincular artigo 2 (Chen et al. 2023)
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '66256f14-1775-4501-b85e-a4ffd71ccc7a',
    id
FROM articles
WHERE doi = '10.3389/fendo.2023.1272797'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- Vincular artigo 3 (Jia et al. 2020)
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '66256f14-1775-4501-b85e-a4ffd71ccc7a',
    id
FROM articles
WHERE doi = '10.1210/clinem/dgaa518'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- Vincular artigo 4 (Samaras et al. 2013)
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '66256f14-1775-4501-b85e-a4ffd71ccc7a',
    id
FROM articles
WHERE doi = '10.1089/rej.2013.1425'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

COMMIT;

-- ============================================================================
-- VERIFICAÇÃO
-- ============================================================================

-- Verificar tamanhos dos campos atualizados
SELECT
    name,
    LENGTH(clinical_relevance) as clinical_relevance_chars,
    LENGTH(patient_explanation) as patient_explanation_chars,
    LENGTH(conduct) as conduct_chars,
    last_review,
    (SELECT COUNT(*)
     FROM article_score_items asi
     WHERE asi.score_item_id = si.id) as articles_count
FROM score_items si
WHERE id = '66256f14-1775-4501-b85e-a4ffd71ccc7a';

-- Listar artigos vinculados
SELECT
    a.title,
    a.authors,
    a.journal,
    a.publish_date,
    a.doi
FROM article_score_items asi
JOIN articles a ON asi.article_id = a.id
WHERE asi.score_item_id = '66256f14-1775-4501-b85e-a4ffd71ccc7a'
ORDER BY a.publish_date DESC;
