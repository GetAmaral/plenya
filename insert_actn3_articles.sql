-- ================================================================
-- INSERIR ARTIGOS CIENTÍFICOS: ACTN3 rs1815739 R577X (Performance)
-- ScoreItem ID: 019c1a2b-a36f-77ac-984e-0dcda6b517f0
-- ================================================================

BEGIN;

-- ================================================================
-- ARTIGO 1: Meta-análise 2024 - Power vs Endurance
-- ================================================================
DO $$
DECLARE
    v_article_id uuid;
BEGIN
    -- Inserir artigo se não existir
    INSERT INTO articles (
        id,
        pm_id,
        doi,
        title,
        authors,
        journal,
        publish_date,
        language,
        abstract,
        article_type,
        keywords,
        embedding_status,
        created_at,
        updated_at
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
    )
    ON CONFLICT (pm_id) DO UPDATE SET
        doi = EXCLUDED.doi,
        title = EXCLUDED.title,
        authors = EXCLUDED.authors,
        abstract = EXCLUDED.abstract,
        updated_at = NOW()
    RETURNING id INTO v_article_id;

    -- Linkar ao ScoreItem se ainda não estiver linkado
    INSERT INTO article_score_items (
        score_item_id,
        article_id,
        confidence_score,
        auto_linked,
        linked_at
    ) VALUES (
        '019c1a2b-a36f-77ac-984e-0dcda6b517f0',
        v_article_id,
        0.95,
        true,
        NOW()
    )
    ON CONFLICT (score_item_id, article_id) DO NOTHING;

    RAISE NOTICE 'Artigo 1 inserido/atualizado: PMID 38609671';
END $$;

-- ================================================================
-- ARTIGO 2: Meta-análise Football 2020
-- ================================================================
DO $$
DECLARE
    v_article_id uuid;
BEGIN
    INSERT INTO articles (
        id,
        pm_id,
        doi,
        title,
        authors,
        journal,
        publish_date,
        language,
        abstract,
        article_type,
        keywords,
        embedding_status,
        created_at,
        updated_at
    ) VALUES (
        uuid_generate_v7(),
        '32856541',
        '10.1080/02640414.2020.1812195',
        'The association of the ACTN3 R577X and ACE I/D polymorphisms with athlete status in football: a systematic review and meta-analysis',
        'McAuley ABT, Hughes DC, Tsaprouni LG, Varley I, Suraci B, Roos TR, Herbert AJ, Kelly AL',
        'J Sports Sci',
        '2021-01-01',
        'en',
        'This meta-analysis examined whether two genetic variants—ACTN3 R577X and ACE I/D polymorphisms—associate with football player status. The researchers analyzed 17 ACTN3 and 19 ACE studies. Key findings indicate that the ACTN3 R allele and professional footballer status and the ACE D allele and youth footballers showed significant associations. The ACTN3 RR and ACE DD genotypes demonstrated the strongest correlations, likely due to their connection with power-oriented athletic traits relevant to soccer success.',
        'meta_analysis',
        ARRAY['ACTN3', 'R577X', 'football', 'soccer', 'athletic performance', 'meta-analysis', 'ACE polymorphism'],
        'pending',
        NOW(),
        NOW()
    )
    ON CONFLICT (pm_id) DO UPDATE SET
        doi = EXCLUDED.doi,
        title = EXCLUDED.title,
        authors = EXCLUDED.authors,
        abstract = EXCLUDED.abstract,
        updated_at = NOW()
    RETURNING id INTO v_article_id;

    INSERT INTO article_score_items (
        score_item_id,
        article_id,
        confidence_score,
        auto_linked,
        linked_at
    ) VALUES (
        '019c1a2b-a36f-77ac-984e-0dcda6b517f0',
        v_article_id,
        0.92,
        true,
        NOW()
    )
    ON CONFLICT (score_item_id, article_id) DO NOTHING;

    RAISE NOTICE 'Artigo 2 inserido/atualizado: PMID 32856541';
END $$;

-- ================================================================
-- ARTIGO 3: Revisão Sistemática - Non-contact Injury 2021
-- ================================================================
DO $$
DECLARE
    v_article_id uuid;
BEGIN
    INSERT INTO articles (
        id,
        pm_id,
        doi,
        title,
        authors,
        journal,
        publish_date,
        language,
        abstract,
        article_type,
        keywords,
        embedding_status,
        created_at,
        updated_at
    ) VALUES (
        uuid_generate_v7(),
        '34284153',
        '10.1016/j.jshs.2021.07.003',
        'Association between ACTN3 R577X genotype and risk of non-contact injury in trained athletes: A systematic review',
        'Zouhal H, et al',
        'J Sport Health Sci',
        '2023-05-01',
        'en',
        'This systematic review examined 13 high-quality studies involving 1,093 participants across six countries to assess links between the ACTN3 R577X genetic polymorphism and non-contact injury risk in athletes. The analysis found that possessing the ACTN3 XX genotype may predispose athletes to a higher probability of some non-contact injuries, such as muscle injury, ankle sprains and exercise-induced muscle damage. Six studies demonstrated significant associations between this genetic variant and injury outcomes, though findings varied by injury type.',
        'systematic_review',
        ARRAY['ACTN3', 'R577X', 'injury risk', 'non-contact injury', 'muscle damage', 'ankle sprain'],
        'pending',
        NOW(),
        NOW()
    )
    ON CONFLICT (pm_id) DO UPDATE SET
        doi = EXCLUDED.doi,
        title = EXCLUDED.title,
        authors = EXCLUDED.authors,
        abstract = EXCLUDED.abstract,
        updated_at = NOW()
    RETURNING id INTO v_article_id;

    INSERT INTO article_score_items (
        score_item_id,
        article_id,
        confidence_score,
        auto_linked,
        linked_at
    ) VALUES (
        '019c1a2b-a36f-77ac-984e-0dcda6b517f0',
        v_article_id,
        0.90,
        true,
        NOW()
    )
    ON CONFLICT (score_item_id, article_id) DO NOTHING;

    RAISE NOTICE 'Artigo 3 inserido/atualizado: PMID 34284153';
END $$;

-- ================================================================
-- ARTIGO 4: Muscle Stiffness 2018
-- ================================================================
DO $$
DECLARE
    v_article_id uuid;
BEGIN
    INSERT INTO articles (
        id,
        pm_id,
        doi,
        title,
        authors,
        journal,
        publish_date,
        language,
        abstract,
        article_type,
        keywords,
        embedding_status,
        created_at,
        updated_at
    ) VALUES (
        uuid_generate_v7(),
        '29032593',
        '10.1111/sms.12994',
        'Association analysis of the ACTN3 R577X polymorphism with passive muscle stiffness and muscle strain injury',
        'Miyamoto N, Miyamoto-Mikami E, Hirata K, Kimura N, Fuku N',
        'Scand J Med Sci Sports',
        '2018-03-01',
        'en',
        'The study examined whether the ACTN3 R577X genetic polymorphism influences muscle stiffness and hamstring injuries in 76 healthy young males. Using ultrasound elastography, researchers found that carriers of the R-allele showed significantly higher muscle stiffness in certain hamstring muscles compared to XX genotype carriers. However, injury frequency did not differ between genetic groups. The findings suggest that while the R-allele variant associates with increased passive muscle stiffness, this mechanical difference might not affect the risk of hamstring muscle strain injury.',
        'research_article',
        ARRAY['ACTN3', 'R577X', 'muscle stiffness', 'hamstring', 'injury risk', 'elastography'],
        'pending',
        NOW(),
        NOW()
    )
    ON CONFLICT (pm_id) DO UPDATE SET
        doi = EXCLUDED.doi,
        title = EXCLUDED.title,
        authors = EXCLUDED.authors,
        abstract = EXCLUDED.abstract,
        updated_at = NOW()
    RETURNING id INTO v_article_id;

    INSERT INTO article_score_items (
        score_item_id,
        article_id,
        confidence_score,
        auto_linked,
        linked_at
    ) VALUES (
        '019c1a2b-a36f-77ac-984e-0dcda6b517f0',
        v_article_id,
        0.88,
        true,
        NOW()
    )
    ON CONFLICT (score_item_id, article_id) DO NOTHING;

    RAISE NOTICE 'Artigo 4 inserido/atualizado: PMID 29032593';
END $$;

-- ================================================================
-- ARTIGO 5: Type IIa Fiber Stiffness 2019
-- ================================================================
DO $$
DECLARE
    v_article_id uuid;
BEGIN
    INSERT INTO articles (
        id,
        pm_id,
        doi,
        title,
        authors,
        journal,
        publish_date,
        language,
        abstract,
        article_type,
        keywords,
        embedding_status,
        created_at,
        updated_at
    ) VALUES (
        uuid_generate_v7(),
        '30360698',
        '10.1080/17461391.2018.1529200',
        'The stiffness response of type IIa fibres after eccentric exercise-induced muscle damage is dependent on ACTN3 r577X polymorphism',
        'Broos S, Malisoux L, Theisen D, Van Thienen R, Francaux M, Thomis MA, Deldicque L',
        'Eur J Sport Sci',
        '2019-05-01',
        'en',
        'The research examined how the ACTN3 gene variant affects muscle response to eccentric exercise damage. Participants with different genotypes performed intensive knee flexion exercises, and researchers measured muscle damage markers and fiber properties. Key findings included a substantial drop in quadriceps strength (averaging 37% at 24 hours post-exercise) and increased stiffness in fast-twitch fibers among RR genotype individuals only. The study suggests that increased stiffness response in fast RR fibres might be a protection mechanism against subsequent muscle damage, though no major genotype differences emerged in acute damage susceptibility.',
        'research_article',
        ARRAY['ACTN3', 'R577X', 'muscle fibers', 'eccentric exercise', 'muscle damage', 'fiber stiffness'],
        'pending',
        NOW(),
        NOW()
    )
    ON CONFLICT (pm_id) DO UPDATE SET
        doi = EXCLUDED.doi,
        title = EXCLUDED.title,
        authors = EXCLUDED.authors,
        abstract = EXCLUDED.abstract,
        updated_at = NOW()
    RETURNING id INTO v_article_id;

    INSERT INTO article_score_items (
        score_item_id,
        article_id,
        confidence_score,
        auto_linked,
        linked_at
    ) VALUES (
        '019c1a2b-a36f-77ac-984e-0dcda6b517f0',
        v_article_id,
        0.91,
        true,
        NOW()
    )
    ON CONFLICT (score_item_id, article_id) DO NOTHING;

    RAISE NOTICE 'Artigo 5 inserido/atualizado: PMID 30360698';
END $$;

-- ================================================================
-- ARTIGO 6: Sports Performance Review 2020
-- ================================================================
DO $$
DECLARE
    v_article_id uuid;
BEGIN
    INSERT INTO articles (
        id,
        pm_id,
        doi,
        title,
        authors,
        journal,
        publish_date,
        language,
        abstract,
        article_type,
        keywords,
        embedding_status,
        created_at,
        updated_at
    ) VALUES (
        uuid_generate_v7(),
        '32668587',
        '10.3390/sports8070099',
        'Effect of ACTN3 Genotype on Sports Performance, Exercise-Induced Muscle Damage, and Injury Epidemiology',
        'Baltazar-Martins G, Gutiérrez-Hellín J, Aguilar-Navarro M, Ruiz-Moreno C, Moreno-Pérez V, López-Samanes Á, Domínguez R, Del Coso J',
        'Sports (Basel)',
        '2020-07-13',
        'en',
        'This review examines how the ACTN3 gene polymorphism affects athletic performance. The polymorphism produces individuals with an XX genotype lacking functional α-actinin-3 protein, while R-allele carriers (RR and RX genotypes) express this protein. The authors note that α-actinin-3 is a protein located within the skeletal muscle with a key role in the production of sarcomeric force. Key findings indicate the XX genotype is underrepresented in elite sprint and power athletes but may carry advantages for metabolic efficiency. The review also discusses emerging evidence linking XX genotype status to increased injury susceptibility, higher exercise-induced muscle damage markers, and greater ankle sprain risk compared to R-allele carriers.',
        'review_article',
        ARRAY['ACTN3', 'R577X', 'sports performance', 'muscle damage', 'injury', 'athletic performance'],
        'pending',
        NOW(),
        NOW()
    )
    ON CONFLICT (pm_id) DO UPDATE SET
        doi = EXCLUDED.doi,
        title = EXCLUDED.title,
        authors = EXCLUDED.authors,
        abstract = EXCLUDED.abstract,
        updated_at = NOW()
    RETURNING id INTO v_article_id;

    INSERT INTO article_score_items (
        score_item_id,
        article_id,
        confidence_score,
        auto_linked,
        linked_at
    ) VALUES (
        '019c1a2b-a36f-77ac-984e-0dcda6b517f0',
        v_article_id,
        0.94,
        true,
        NOW()
    )
    ON CONFLICT (score_item_id, article_id) DO NOTHING;

    RAISE NOTICE 'Artigo 6 inserido/atualizado: PMID 32668587';
END $$;

-- ================================================================
-- ARTIGO 7: Inflammatory Response 2022
-- ================================================================
DO $$
DECLARE
    v_article_id uuid;
BEGIN
    INSERT INTO articles (
        id,
        pm_id,
        doi,
        title,
        authors,
        journal,
        publish_date,
        language,
        abstract,
        article_type,
        keywords,
        embedding_status,
        created_at,
        updated_at
    ) VALUES (
        uuid_generate_v7(),
        '36567902',
        '10.1155/2022/5447100',
        'Influence of Alpha-Actinin-3 R577X Polymorphism on Muscle Damage and the Inflammatory Response after an Acute Strength Training Session',
        'Pereira MA, Rosse IC, Silva AC, Coelho PJFN, Castro BM, Becker LK, Ferreira-Júnior JB, Oliveira EC, Pinto KMC, Talvani A, Queiroz KB, Coelho DB',
        'Biomed Res Int',
        '2022-12-16',
        'en',
        'This investigation examined how the ACTN3 R577X genetic polymorphism affects muscle injury markers and immune responses following intense leg resistance exercise. Twenty-seven healthy males were classified by genotype (18 RR/RX carriers, 9 XX carriers). The RR/RX group demonstrated superior strength performance and elevated muscle damage indicators at 24 hours. Conversely, the XX genotype showed elevated inflammatory markers at baseline and greater delayed-onset muscle soreness despite lower training loads, suggesting greater vulnerability associated with genetic profile XX compared to genetic profiles RR and RX.',
        'research_article',
        ARRAY['ACTN3', 'R577X', 'muscle damage', 'inflammation', 'strength training', 'DOMS'],
        'pending',
        NOW(),
        NOW()
    )
    ON CONFLICT (pm_id) DO UPDATE SET
        doi = EXCLUDED.doi,
        title = EXCLUDED.title,
        authors = EXCLUDED.authors,
        abstract = EXCLUDED.abstract,
        updated_at = NOW()
    RETURNING id INTO v_article_id;

    INSERT INTO article_score_items (
        score_item_id,
        article_id,
        confidence_score,
        auto_linked,
        linked_at
    ) VALUES (
        '019c1a2b-a36f-77ac-984e-0dcda6b517f0',
        v_article_id,
        0.93,
        true,
        NOW()
    )
    ON CONFLICT (score_item_id, article_id) DO NOTHING;

    RAISE NOTICE 'Artigo 7 inserido/atualizado: PMID 36567902';
END $$;

-- ================================================================
-- ARTIGO 8: Japanese Elite Athletes 2022
-- ================================================================
DO $$
DECLARE
    v_article_id uuid;
BEGIN
    INSERT INTO articles (
        id,
        pm_id,
        doi,
        title,
        authors,
        journal,
        publish_date,
        language,
        abstract,
        article_type,
        keywords,
        embedding_status,
        created_at,
        updated_at
    ) VALUES (
        uuid_generate_v7(),
        '36247951',
        '10.5114/biolsport.2022.108704',
        'The association of ACTN3 R577X polymorphism with sports specificity in Japanese elite athletes',
        'Akazawa N, Ohiwa N, Shimizu K, Suzuki N, Kumagai H, Fuku N, Suzuki Y',
        'Biol Sport',
        '2022-10-01',
        'en',
        'The research examined whether the ACTN3 R577X gene variant relates to athletic performance across different sports categories. The study involved 906 Japanese elite athletes and 649 controls, categorized into sprint/power, endurance, artistic, martial arts, and ball game sports. Key findings: Athletes carrying the R allele showed higher prevalence in sprint/power, martial arts, and ball game sports compared to controls and endurance athletes. The RR genotype was significantly more frequent among sprint/power, martial arts, and ball game athletes. A notable linear trend connected increasing R allele frequency with enhanced performance in sprint/power sports at elite competitive levels.',
        'research_article',
        ARRAY['ACTN3', 'R577X', 'elite athletes', 'sports specificity', 'Japanese population', 'sprint performance'],
        'pending',
        NOW(),
        NOW()
    )
    ON CONFLICT (pm_id) DO UPDATE SET
        doi = EXCLUDED.doi,
        title = EXCLUDED.title,
        authors = EXCLUDED.authors,
        abstract = EXCLUDED.abstract,
        updated_at = NOW()
    RETURNING id INTO v_article_id;

    INSERT INTO article_score_items (
        score_item_id,
        article_id,
        confidence_score,
        auto_linked,
        linked_at
    ) VALUES (
        '019c1a2b-a36f-77ac-984e-0dcda6b517f0',
        v_article_id,
        0.89,
        true,
        NOW()
    )
    ON CONFLICT (score_item_id, article_id) DO NOTHING;

    RAISE NOTICE 'Artigo 8 inserido/atualizado: PMID 36247951';
END $$;

-- ================================================================
-- VERIFICAÇÃO FINAL
-- ================================================================
SELECT
    '=== RESUMO DE ARTIGOS LINKADOS ===' as info;

SELECT
    COUNT(*) as total_artigos_cientificos_pubmed
FROM article_score_items asi
JOIN articles a ON asi.article_id = a.id
WHERE asi.score_item_id = '019c1a2b-a36f-77ac-984e-0dcda6b517f0'
  AND a.pm_id IS NOT NULL;

SELECT
    COUNT(*) as total_artigos_todos
FROM article_score_items
WHERE score_item_id = '019c1a2b-a36f-77ac-984e-0dcda6b517f0';

SELECT
    a.pm_id,
    LEFT(a.title, 80) as title,
    SUBSTRING(a.journal FROM 1 FOR 30) as journal,
    EXTRACT(YEAR FROM a.publish_date) as year,
    asi.confidence_score
FROM article_score_items asi
JOIN articles a ON asi.article_id = a.id
WHERE asi.score_item_id = '019c1a2b-a36f-77ac-984e-0dcda6b517f0'
ORDER BY a.publish_date DESC;

COMMIT;

-- ================================================================
-- SUCESSO! 8 artigos científicos PubMed inseridos e linkados
-- ================================================================
