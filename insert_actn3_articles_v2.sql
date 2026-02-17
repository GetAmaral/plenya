-- ================================================================
-- INSERIR ARTIGOS CIENTÍFICOS: ACTN3 rs1815739 R577X (Performance)
-- ScoreItem ID: 019c1a2b-a36f-77ac-984e-0dcda6b517f0
-- ================================================================

BEGIN;

-- ================================================================
-- ARTIGO 1: Meta-análise 2024 - Power vs Endurance (PMID: 38609671)
-- ================================================================
DO $$
DECLARE
    v_article_id uuid;
    v_existing_id uuid;
BEGIN
    -- Verificar se artigo já existe por PMID ou DOI
    SELECT id INTO v_existing_id
    FROM articles
    WHERE pm_id = '38609671' OR doi = '10.1186/s40798-024-00711-x'
    LIMIT 1;

    IF v_existing_id IS NULL THEN
        -- Inserir novo artigo
        INSERT INTO articles (
            id, pm_id, doi, title, authors, journal, publish_date, language,
            abstract, article_type, keywords, embedding_status, created_at, updated_at
        ) VALUES (
            uuid_generate_v7(),
            '38609671',
            '10.1186/s40798-024-00711-x',
            'A Systematic Review and Meta-analysis of the Association Between ACTN3 R577X Genotypes and Performance in Endurance Versus Power Athletes and Non-athletes',
            'El Ouali EM, Barthelemy B, Granacher U, Zouhal H',
            'Sports Med Open',
            '2024-04-12',
            'en',
            'This meta-analysis of 25 studies involving 14,541 participants examined ACTN3 R577X genotype frequencies across athlete types. Key findings indicate the RX genotype predominated in power athletes, with the RR genotype and R allele being overrepresented in power athletes compared to non-athletes and endurance athletes. The research suggests the R allele, which maintains normal α-actinin-3 expression in fast-twitch muscle, may enhance strength and power performance development.',
            'meta_analysis',
            ARRAY['ACTN3', 'R577X', 'athletic performance', 'meta-analysis', 'power athletes', 'endurance athletes', 'genetic polymorphism'],
            'pending',
            NOW(),
            NOW()
        ) RETURNING id INTO v_article_id;
        RAISE NOTICE 'Artigo 1 INSERIDO: PMID 38609671';
    ELSE
        v_article_id := v_existing_id;
        RAISE NOTICE 'Artigo 1 JÁ EXISTE: PMID 38609671 (ID: %)', v_article_id;
    END IF;

    -- Linkar ao ScoreItem
    INSERT INTO article_score_items (score_item_id, article_id, confidence_score, auto_linked, linked_at)
    VALUES ('019c1a2b-a36f-77ac-984e-0dcda6b517f0', v_article_id, 0.95, true, NOW())
    ON CONFLICT (score_item_id, article_id) DO NOTHING;
END $$;

-- ================================================================
-- ARTIGO 2: Meta-análise Football 2020 (PMID: 32856541)
-- ================================================================
DO $$
DECLARE
    v_article_id uuid;
    v_existing_id uuid;
BEGIN
    SELECT id INTO v_existing_id FROM articles WHERE pm_id = '32856541' OR doi = '10.1080/02640414.2020.1812195' LIMIT 1;
    IF v_existing_id IS NULL THEN
        INSERT INTO articles (id, pm_id, doi, title, authors, journal, publish_date, language, abstract, article_type, keywords, embedding_status, created_at, updated_at)
        VALUES (
            uuid_generate_v7(), '32856541', '10.1080/02640414.2020.1812195',
            'The association of the ACTN3 R577X and ACE I/D polymorphisms with athlete status in football: a systematic review and meta-analysis',
            'McAuley ABT, Hughes DC, Tsaprouni LG, Varley I, Suraci B, Roos TR, Herbert AJ, Kelly AL',
            'J Sports Sci', '2021-01-01', 'en',
            'This meta-analysis examined whether two genetic variants—ACTN3 R577X and ACE I/D polymorphisms—associate with football player status. The researchers analyzed 17 ACTN3 and 19 ACE studies. Key findings indicate that the ACTN3 R allele and professional footballer status and the ACE D allele and youth footballers showed significant associations.',
            'meta_analysis',
            ARRAY['ACTN3', 'R577X', 'football', 'soccer', 'athletic performance', 'meta-analysis'],
            'pending', NOW(), NOW()
        ) RETURNING id INTO v_article_id;
        RAISE NOTICE 'Artigo 2 INSERIDO: PMID 32856541';
    ELSE
        v_article_id := v_existing_id;
        RAISE NOTICE 'Artigo 2 JÁ EXISTE: PMID 32856541';
    END IF;
    INSERT INTO article_score_items (score_item_id, article_id, confidence_score, auto_linked, linked_at)
    VALUES ('019c1a2b-a36f-77ac-984e-0dcda6b517f0', v_article_id, 0.92, true, NOW())
    ON CONFLICT (score_item_id, article_id) DO NOTHING;
END $$;

-- ================================================================
-- ARTIGO 3: Non-contact Injury 2021 (PMID: 34284153)
-- ================================================================
DO $$
DECLARE
    v_article_id uuid;
    v_existing_id uuid;
BEGIN
    SELECT id INTO v_existing_id FROM articles WHERE pm_id = '34284153' OR doi = '10.1016/j.jshs.2021.07.003' LIMIT 1;
    IF v_existing_id IS NULL THEN
        INSERT INTO articles (id, pm_id, doi, title, authors, journal, publish_date, language, abstract, article_type, keywords, embedding_status, created_at, updated_at)
        VALUES (
            uuid_generate_v7(), '34284153', '10.1016/j.jshs.2021.07.003',
            'Association between ACTN3 R577X genotype and risk of non-contact injury in trained athletes: A systematic review',
            'Zouhal H, et al', 'J Sport Health Sci', '2023-05-01', 'en',
            'This systematic review examined 13 high-quality studies involving 1,093 participants to assess links between the ACTN3 R577X genetic polymorphism and non-contact injury risk in athletes. The analysis found that possessing the ACTN3 XX genotype may predispose athletes to a higher probability of some non-contact injuries, such as muscle injury and ankle sprains.',
            'review',
            ARRAY['ACTN3', 'R577X', 'injury risk', 'non-contact injury', 'muscle damage'],
            'pending', NOW(), NOW()
        ) RETURNING id INTO v_article_id;
        RAISE NOTICE 'Artigo 3 INSERIDO: PMID 34284153';
    ELSE
        v_article_id := v_existing_id;
        RAISE NOTICE 'Artigo 3 JÁ EXISTE: PMID 34284153';
    END IF;
    INSERT INTO article_score_items (score_item_id, article_id, confidence_score, auto_linked, linked_at)
    VALUES ('019c1a2b-a36f-77ac-984e-0dcda6b517f0', v_article_id, 0.90, true, NOW())
    ON CONFLICT (score_item_id, article_id) DO NOTHING;
END $$;

-- ================================================================
-- ARTIGO 4: Muscle Stiffness 2018 (PMID: 29032593)
-- ================================================================
DO $$
DECLARE
    v_article_id uuid;
    v_existing_id uuid;
BEGIN
    SELECT id INTO v_existing_id FROM articles WHERE pm_id = '29032593' OR doi = '10.1111/sms.12994' LIMIT 1;
    IF v_existing_id IS NULL THEN
        INSERT INTO articles (id, pm_id, doi, title, authors, journal, publish_date, language, abstract, article_type, keywords, embedding_status, created_at, updated_at)
        VALUES (
            uuid_generate_v7(), '29032593', '10.1111/sms.12994',
            'Association analysis of the ACTN3 R577X polymorphism with passive muscle stiffness and muscle strain injury',
            'Miyamoto N, Miyamoto-Mikami E, Hirata K, Kimura N, Fuku N',
            'Scand J Med Sci Sports', '2018-03-01', 'en',
            'The study examined whether the ACTN3 R577X genetic polymorphism influences muscle stiffness and hamstring injuries. Using ultrasound elastography, researchers found that carriers of the R-allele showed significantly higher muscle stiffness compared to XX genotype carriers.',
            'research_article',
            ARRAY['ACTN3', 'R577X', 'muscle stiffness', 'hamstring'],
            'pending', NOW(), NOW()
        ) RETURNING id INTO v_article_id;
        RAISE NOTICE 'Artigo 4 INSERIDO: PMID 29032593';
    ELSE
        v_article_id := v_existing_id;
        RAISE NOTICE 'Artigo 4 JÁ EXISTE: PMID 29032593';
    END IF;
    INSERT INTO article_score_items (score_item_id, article_id, confidence_score, auto_linked, linked_at)
    VALUES ('019c1a2b-a36f-77ac-984e-0dcda6b517f0', v_article_id, 0.88, true, NOW())
    ON CONFLICT (score_item_id, article_id) DO NOTHING;
END $$;

-- ================================================================
-- ARTIGO 5: Type IIa Fiber Stiffness 2019 (PMID: 30360698)
-- ================================================================
DO $$
DECLARE
    v_article_id uuid;
    v_existing_id uuid;
BEGIN
    SELECT id INTO v_existing_id FROM articles WHERE pm_id = '30360698' OR doi = '10.1080/17461391.2018.1529200' LIMIT 1;
    IF v_existing_id IS NULL THEN
        INSERT INTO articles (id, pm_id, doi, title, authors, journal, publish_date, language, abstract, article_type, keywords, embedding_status, created_at, updated_at)
        VALUES (
            uuid_generate_v7(), '30360698', '10.1080/17461391.2018.1529200',
            'The stiffness response of type IIa fibres after eccentric exercise-induced muscle damage is dependent on ACTN3 r577X polymorphism',
            'Broos S, Malisoux L, Theisen D, Van Thienen R, Francaux M, Thomis MA, Deldicque L',
            'Eur J Sport Sci', '2019-05-01', 'en',
            'The research examined how the ACTN3 gene variant affects muscle response to eccentric exercise damage. Key findings included increased stiffness in fast-twitch fibers among RR genotype individuals, suggesting a protection mechanism against muscle damage.',
            'research_article',
            ARRAY['ACTN3', 'R577X', 'muscle fibers', 'eccentric exercise'],
            'pending', NOW(), NOW()
        ) RETURNING id INTO v_article_id;
        RAISE NOTICE 'Artigo 5 INSERIDO: PMID 30360698';
    ELSE
        v_article_id := v_existing_id;
        RAISE NOTICE 'Artigo 5 JÁ EXISTE: PMID 30360698';
    END IF;
    INSERT INTO article_score_items (score_item_id, article_id, confidence_score, auto_linked, linked_at)
    VALUES ('019c1a2b-a36f-77ac-984e-0dcda6b517f0', v_article_id, 0.91, true, NOW())
    ON CONFLICT (score_item_id, article_id) DO NOTHING;
END $$;

-- ================================================================
-- ARTIGO 6: Sports Performance Review 2020 (PMID: 32668587)
-- ================================================================
DO $$
DECLARE
    v_article_id uuid;
    v_existing_id uuid;
BEGIN
    SELECT id INTO v_existing_id FROM articles WHERE pm_id = '32668587' OR doi = '10.3390/sports8070099' LIMIT 1;
    IF v_existing_id IS NULL THEN
        INSERT INTO articles (id, pm_id, doi, title, authors, journal, publish_date, language, abstract, article_type, keywords, embedding_status, created_at, updated_at)
        VALUES (
            uuid_generate_v7(), '32668587', '10.3390/sports8070099',
            'Effect of ACTN3 Genotype on Sports Performance, Exercise-Induced Muscle Damage, and Injury Epidemiology',
            'Baltazar-Martins G, Gutiérrez-Hellín J, Aguilar-Navarro M, Ruiz-Moreno C, Moreno-Pérez V, López-Samanes Á, Domínguez R, Del Coso J',
            'Sports (Basel)', '2020-07-13', 'en',
            'This review examines how the ACTN3 gene polymorphism affects athletic performance. The XX genotype is underrepresented in elite sprint and power athletes but may carry advantages for metabolic efficiency. The review also discusses emerging evidence linking XX genotype status to increased injury susceptibility.',
            'review_article',
            ARRAY['ACTN3', 'R577X', 'sports performance', 'muscle damage'],
            'pending', NOW(), NOW()
        ) RETURNING id INTO v_article_id;
        RAISE NOTICE 'Artigo 6 INSERIDO: PMID 32668587';
    ELSE
        v_article_id := v_existing_id;
        RAISE NOTICE 'Artigo 6 JÁ EXISTE: PMID 32668587';
    END IF;
    INSERT INTO article_score_items (score_item_id, article_id, confidence_score, auto_linked, linked_at)
    VALUES ('019c1a2b-a36f-77ac-984e-0dcda6b517f0', v_article_id, 0.94, true, NOW())
    ON CONFLICT (score_item_id, article_id) DO NOTHING;
END $$;

-- ================================================================
-- ARTIGO 7: Inflammatory Response 2022 (PMID: 36567902)
-- ================================================================
DO $$
DECLARE
    v_article_id uuid;
    v_existing_id uuid;
BEGIN
    SELECT id INTO v_existing_id FROM articles WHERE pm_id = '36567902' OR doi = '10.1155/2022/5447100' LIMIT 1;
    IF v_existing_id IS NULL THEN
        INSERT INTO articles (id, pm_id, doi, title, authors, journal, publish_date, language, abstract, article_type, keywords, embedding_status, created_at, updated_at)
        VALUES (
            uuid_generate_v7(), '36567902', '10.1155/2022/5447100',
            'Influence of Alpha-Actinin-3 R577X Polymorphism on Muscle Damage and the Inflammatory Response after an Acute Strength Training Session',
            'Pereira MA, Rosse IC, Silva AC, Coelho PJFN, Castro BM, Becker LK, Ferreira-Júnior JB, Oliveira EC, Pinto KMC, Talvani A, Queiroz KB, Coelho DB',
            'Biomed Res Int', '2022-12-16', 'en',
            'This investigation examined how the ACTN3 R577X genetic polymorphism affects muscle injury markers and immune responses following intense leg resistance exercise. The XX genotype showed elevated inflammatory markers and greater delayed-onset muscle soreness, suggesting greater vulnerability compared to RR and RX profiles.',
            'research_article',
            ARRAY['ACTN3', 'R577X', 'muscle damage', 'inflammation', 'strength training'],
            'pending', NOW(), NOW()
        ) RETURNING id INTO v_article_id;
        RAISE NOTICE 'Artigo 7 INSERIDO: PMID 36567902';
    ELSE
        v_article_id := v_existing_id;
        RAISE NOTICE 'Artigo 7 JÁ EXISTE: PMID 36567902';
    END IF;
    INSERT INTO article_score_items (score_item_id, article_id, confidence_score, auto_linked, linked_at)
    VALUES ('019c1a2b-a36f-77ac-984e-0dcda6b517f0', v_article_id, 0.93, true, NOW())
    ON CONFLICT (score_item_id, article_id) DO NOTHING;
END $$;

-- ================================================================
-- ARTIGO 8: Japanese Elite Athletes 2022 (PMID: 36247951)
-- ================================================================
DO $$
DECLARE
    v_article_id uuid;
    v_existing_id uuid;
BEGIN
    SELECT id INTO v_existing_id FROM articles WHERE pm_id = '36247951' OR doi = '10.5114/biolsport.2022.108704' LIMIT 1;
    IF v_existing_id IS NULL THEN
        INSERT INTO articles (id, pm_id, doi, title, authors, journal, publish_date, language, abstract, article_type, keywords, embedding_status, created_at, updated_at)
        VALUES (
            uuid_generate_v7(), '36247951', '10.5114/biolsport.2022.108704',
            'The association of ACTN3 R577X polymorphism with sports specificity in Japanese elite athletes',
            'Akazawa N, Ohiwa N, Shimizu K, Suzuki N, Kumagai H, Fuku N, Suzuki Y',
            'Biol Sport', '2022-10-01', 'en',
            'The research examined whether the ACTN3 R577X gene variant relates to athletic performance across different sports categories in 906 Japanese elite athletes. Athletes carrying the R allele showed higher prevalence in sprint/power sports. A notable linear trend connected increasing R allele frequency with enhanced performance in sprint/power sports.',
            'research_article',
            ARRAY['ACTN3', 'R577X', 'elite athletes', 'sports specificity', 'Japanese population'],
            'pending', NOW(), NOW()
        ) RETURNING id INTO v_article_id;
        RAISE NOTICE 'Artigo 8 INSERIDO: PMID 36247951';
    ELSE
        v_article_id := v_existing_id;
        RAISE NOTICE 'Artigo 8 JÁ EXISTE: PMID 36247951';
    END IF;
    INSERT INTO article_score_items (score_item_id, article_id, confidence_score, auto_linked, linked_at)
    VALUES ('019c1a2b-a36f-77ac-984e-0dcda6b517f0', v_article_id, 0.89, true, NOW())
    ON CONFLICT (score_item_id, article_id) DO NOTHING;
END $$;

-- ================================================================
-- VERIFICAÇÃO FINAL
-- ================================================================
SELECT '=== RESUMO DE ARTIGOS LINKADOS ===' as info;

SELECT COUNT(*) as total_artigos_pubmed
FROM article_score_items asi
JOIN articles a ON asi.article_id = a.id
WHERE asi.score_item_id = '019c1a2b-a36f-77ac-984e-0dcda6b517f0'
  AND a.pm_id IS NOT NULL;

SELECT COUNT(*) as total_artigos_todos
FROM article_score_items
WHERE score_item_id = '019c1a2b-a36f-77ac-984e-0dcda6b517f0';

SELECT
    a.pm_id,
    LEFT(a.title, 70) as title,
    SUBSTRING(a.journal FROM 1 FOR 25) as journal,
    EXTRACT(YEAR FROM a.publish_date) as year,
    asi.confidence_score
FROM article_score_items asi
JOIN articles a ON asi.article_id = a.id
WHERE asi.score_item_id = '019c1a2b-a36f-77ac-984e-0dcda6b517f0'
ORDER BY a.publish_date DESC NULLS LAST;

COMMIT;

-- ================================================================
-- SUCESSO! 8 artigos científicos PubMed inseridos e linkados
-- ================================================================
