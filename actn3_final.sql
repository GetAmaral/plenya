BEGIN;

-- Artigo 1 (Meta-análise)
INSERT INTO articles (id, pm_id, doi, title, authors, journal, publish_date, language, abstract, article_type, keywords, embedding_status)
SELECT uuid_generate_v7(), '38609671', '10.1186/s40798-024-00711-x',
'A Systematic Review and Meta-analysis of the Association Between ACTN3 R577X Genotypes and Performance in Endurance Versus Power Athletes and Non-athletes',
'El Ouali EM, et al', 'Sports Med Open', '2024-04-12', 'en',
'Meta-analysis of 25 studies with 14,541 participants on ACTN3 R577X and athletic performance. RR genotype overrepresented in power athletes.',
'meta_analysis', ARRAY['ACTN3','R577X','performance'], 'pending'
WHERE NOT EXISTS (SELECT 1 FROM articles WHERE pm_id='38609671')
RETURNING id \gset art1_

INSERT INTO article_score_items (score_item_id, article_id, confidence_score, auto_linked, linked_at)
SELECT '019c1a2b-a36f-77ac-984e-0dcda6b517f0', COALESCE(:'art1_id', (SELECT id FROM articles WHERE pm_id='38609671')), 0.95, true, NOW()
ON CONFLICT DO NOTHING;

-- Artigo 2 (Meta-análise Football)
INSERT INTO articles (id, pm_id, doi, title, authors, journal, publish_date, language, abstract, article_type, keywords, embedding_status)
SELECT uuid_generate_v7(), '32856541', '10.1080/02640414.2020.1812195',
'The association of the ACTN3 R577X and ACE I/D polymorphisms with athlete status in football',
'McAuley ABT, et al', 'J Sports Sci', '2021-01-01', 'en',
'Meta-analysis of ACTN3 R577X in football. R allele associated with professional footballer status.',
'meta_analysis', ARRAY['ACTN3','football','performance'], 'pending'
WHERE NOT EXISTS (SELECT 1 FROM articles WHERE pm_id='32856541')
RETURNING id \gset art2_

INSERT INTO article_score_items (score_item_id, article_id, confidence_score, auto_linked, linked_at)
SELECT '019c1a2b-a36f-77ac-984e-0dcda6b517f0', COALESCE(:'art2_id', (SELECT id FROM articles WHERE pm_id='32856541')), 0.92, true, NOW()
ON CONFLICT DO NOTHING;

-- Artigo 3 (Revisão - Injury)
INSERT INTO articles (id, pm_id, doi, title, authors, journal, publish_date, language, abstract, article_type, keywords, embedding_status)
SELECT uuid_generate_v7(), '34284153', '10.1016/j.jshs.2021.07.003',
'Association between ACTN3 R577X genotype and risk of non-contact injury in trained athletes',
'Zouhal H, et al', 'J Sport Health Sci', '2023-05-01', 'en',
'Systematic review of 13 studies. XX genotype associated with higher injury risk.',
'review', ARRAY['ACTN3','injury','muscle damage'], 'pending'
WHERE NOT EXISTS (SELECT 1 FROM articles WHERE pm_id='34284153')
RETURNING id \gset art3_

INSERT INTO article_score_items (score_item_id, article_id, confidence_score, auto_linked, linked_at)
SELECT '019c1a2b-a36f-77ac-984e-0dcda6b517f0', COALESCE(:'art3_id', (SELECT id FROM articles WHERE pm_id='34284153')), 0.90, true, NOW()
ON CONFLICT DO NOTHING;

-- Artigo 4 (Research - Stiffness)
INSERT INTO articles (id, pm_id, doi, title, authors, journal, publish_date, language, abstract, article_type, keywords, embedding_status)
SELECT uuid_generate_v7(), '29032593', '10.1111/sms.12994',
'Association analysis of the ACTN3 R577X polymorphism with passive muscle stiffness',
'Miyamoto N, et al', 'Scand J Med Sci Sports', '2018-03-01', 'en',
'R-allele carriers showed higher muscle stiffness in hamstrings.',
'research_article', ARRAY['ACTN3','muscle stiffness'], 'pending'
WHERE NOT EXISTS (SELECT 1 FROM articles WHERE pm_id='29032593')
RETURNING id \gset art4_

INSERT INTO article_score_items (score_item_id, article_id, confidence_score, auto_linked, linked_at)
SELECT '019c1a2b-a36f-77ac-984e-0dcda6b517f0', COALESCE(:'art4_id', (SELECT id FROM articles WHERE pm_id='29032593')), 0.88, true, NOW()
ON CONFLICT DO NOTHING;

-- Artigo 5 (Research - Fiber Stiffness)
INSERT INTO articles (id, pm_id, doi, title, authors, journal, publish_date, language, abstract, article_type, keywords, embedding_status)
SELECT uuid_generate_v7(), '30360698', '10.1080/17461391.2018.1529200',
'The stiffness response of type IIa fibres after eccentric exercise-induced muscle damage',
'Broos S, et al', 'Eur J Sport Sci', '2019-05-01', 'en',
'RR genotype showed increased fiber stiffness after eccentric exercise.',
'research_article', ARRAY['ACTN3','muscle fibers','exercise'], 'pending'
WHERE NOT EXISTS (SELECT 1 FROM articles WHERE pm_id='30360698')
RETURNING id \gset art5_

INSERT INTO article_score_items (score_item_id, article_id, confidence_score, auto_linked, linked_at)
SELECT '019c1a2b-a36f-77ac-984e-0dcda6b517f0', COALESCE(:'art5_id', (SELECT id FROM articles WHERE pm_id='30360698')), 0.91, true, NOW()
ON CONFLICT DO NOTHING;

-- Artigo 6 (Review - Sports Performance)
INSERT INTO articles (id, pm_id, doi, title, authors, journal, publish_date, language, abstract, article_type, keywords, embedding_status)
SELECT uuid_generate_v7(), '32668587', '10.3390/sports8070099',
'Effect of ACTN3 Genotype on Sports Performance, Exercise-Induced Muscle Damage, and Injury Epidemiology',
'Baltazar-Martins G, et al', 'Sports (Basel)', '2020-07-13', 'en',
'Comprehensive review on ACTN3 effects. XX genotype linked to injury susceptibility.',
'review', ARRAY['ACTN3','performance','injury'], 'pending'
WHERE NOT EXISTS (SELECT 1 FROM articles WHERE pm_id='32668587')
RETURNING id \gset art6_

INSERT INTO article_score_items (score_item_id, article_id, confidence_score, auto_linked, linked_at)
SELECT '019c1a2b-a36f-77ac-984e-0dcda6b517f0', COALESCE(:'art6_id', (SELECT id FROM articles WHERE pm_id='32668587')), 0.94, true, NOW()
ON CONFLICT DO NOTHING;

-- Artigo 7 (Research - Inflammation)
INSERT INTO articles (id, pm_id, doi, title, authors, journal, publish_date, language, abstract, article_type, keywords, embedding_status)
SELECT uuid_generate_v7(), '36567902', '10.1155/2022/5447100',
'Influence of Alpha-Actinin-3 R577X Polymorphism on Muscle Damage and Inflammatory Response',
'Pereira MA, et al', 'Biomed Res Int', '2022-12-16', 'en',
'XX genotype showed elevated inflammatory markers and greater muscle soreness.',
'research_article', ARRAY['ACTN3','inflammation','muscle damage'], 'pending'
WHERE NOT EXISTS (SELECT 1 FROM articles WHERE pm_id='36567902')
RETURNING id \gset art7_

INSERT INTO article_score_items (score_item_id, article_id, confidence_score, auto_linked, linked_at)
SELECT '019c1a2b-a36f-77ac-984e-0dcda6b517f0', COALESCE(:'art7_id', (SELECT id FROM articles WHERE pm_id='36567902')), 0.93, true, NOW()
ON CONFLICT DO NOTHING;

-- Artigo 8 (Research - Japanese Athletes)
INSERT INTO articles (id, pm_id, doi, title, authors, journal, publish_date, language, abstract, article_type, keywords, embedding_status)
SELECT uuid_generate_v7(), '36247951', '10.5114/biolsport.2022.108704',
'The association of ACTN3 R577X polymorphism with sports specificity in Japanese elite athletes',
'Akazawa N, et al', 'Biol Sport', '2022-10-01', 'en',
'Study of 906 Japanese elite athletes. R allele more prevalent in sprint/power sports.',
'research_article', ARRAY['ACTN3','elite athletes','Japanese'], 'pending'
WHERE NOT EXISTS (SELECT 1 FROM articles WHERE pm_id='36247951')
RETURNING id \gset art8_

INSERT INTO article_score_items (score_item_id, article_id, confidence_score, auto_linked, linked_at)
SELECT '019c1a2b-a36f-77ac-984e-0dcda6b517f0', COALESCE(:'art8_id', (SELECT id FROM articles WHERE pm_id='36247951')), 0.89, true, NOW()
ON CONFLICT DO NOTHING;

-- Verificação final
SELECT '=== RESULTADO FINAL ===' as status;
SELECT COUNT(*) as total_pubmed FROM article_score_items asi JOIN articles a ON asi.article_id = a.id
WHERE asi.score_item_id = '019c1a2b-a36f-77ac-984e-0dcda6b517f0' AND a.pm_id IS NOT NULL;

SELECT COUNT(*) as total_todos FROM article_score_items
WHERE score_item_id = '019c1a2b-a36f-77ac-984e-0dcda6b517f0';

COMMIT;
