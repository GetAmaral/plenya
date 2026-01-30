-- ========================================
-- ENRICHMENT: Cristais Patológicos (Score Item)
-- ID: ebcc36fd-d285-4754-adf7-50c7b130b286
-- Generated: 2026-01-28
-- ========================================

BEGIN;

-- ========================================
-- STEP 1: Insert Scientific Articles
-- ========================================

-- Article 1: Outpatient Crystalluria Study (2025)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract
) VALUES (
    'Outpatient crystalluria: prevalence, crystal types, and associations with comorbidities and urinary tract infections at a provincial hospital',
    'Samira Natoubi, Rim Jamal, Nezha Baghdad',
    'Iranian Journal of Microbiology',
    '2025-06-01'::date,
    'en',
    'research_article',
    '10.18502/ijm.v17i3.18820',
    '40612723',
    'Background and Objectives: Crystalluria refers to the occurrence of crystals in urine resulting from urinary supersaturation, which disrupts the balance between factors that promote and those that inhibit crystal formation in urine. This study examined crystalluria occurrence in 1,025 urine samples from outpatients at a provincial hospital. Results: 22.04% of samples showed crystalluria with mean patient age of 51.3 years. The most common crystal types were calcium oxalate (46.4%), uric acid (23.5%), urates (15.1%) and struvite (9.3%). Diabetes, kidney failure, prostatitis, and nephrotic syndrome were associated with crystal formation. The prevalence of urinary tract infections in patients with urinary crystals was 10.6%. Notably, struvite crystals correlated with bacterial infections. Conclusion: Monitoring crystals is vital for preventing kidney stone development and infection-related complications in susceptible populations.'
) ON CONFLICT (doi) DO UPDATE
SET updated_at = NOW();

-- Article 2: Amoxicillin Crystalluria Review (2025)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract
) VALUES (
    'Amoxicillin crystalluria and amoxicillin-induced crystal nephropathy: a narrative review',
    'Dominique Vodovar, Cyril Mousseaux, Michel Daudon, Matthieu Jamme, Emmanuel Letavernier',
    'Kidney International',
    '2025-01-01'::date,
    'en',
    'review',
    '10.1016/j.kint.2024.09.019',
    '39490983',
    'Amoxicillin crystalluria is defined as the precipitation of amoxicillin trihydrate crystals in urine, while amoxicillin-induced crystal nephropathy involves kidney tubule obstruction causing acute kidney injury. This narrative review reveals higher prevalence (24-41%) than previously recognized in patients receiving high-dose intravenous amoxicillin (≥150 mg/kg daily). The condition typically presents with cloudy urine and macroscopic hematuria. Management involves discontinuing the antibiotic and aggressive fluid resuscitation to enhance urine flow. Patients generally recover normal kidney function rapidly after stopping amoxicillin, though 10-40% require renal replacement therapy. No deaths have been directly attributed to this condition.'
) ON CONFLICT (doi) DO UPDATE
SET updated_at = NOW();

-- Article 3: Molecular Mechanisms of Stone Formation (2024)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract
) VALUES (
    'Pathophysiology and Main Molecular Mechanisms of Urinary Stone Formation and Recurrence',
    'Flavia Tamborino, Rossella Cicchetti, Marco Mascitti, Giulio Litterio, Angelo Orsini, Simone Ferretti, Martina Basconi, Antonio De Palma, Matteo Ferro, Michele Marchioni, Luigi Schips',
    'International Journal of Molecular Sciences',
    '2024-03-06'::date,
    'en',
    'review',
    '10.3390/ijms25053075',
    '38474319',
    'Urinary stone disease (nephrolithiasis) is a multifactorial condition affecting 12% of the global population with recurrence rates as high as 50% within 5 years after first stone onset. This comprehensive review examines the pathophysiology and molecular mechanisms underlying urinary stone formation and recurrence. The review covers crystal nucleation, aggregation, and retention processes, highlighting the role of urinary supersaturation, pH imbalances, and inhibitor deficiencies. Special attention is given to calcium oxalate and calcium phosphate stones (80% of cases), uric acid stones, cystine stones, and infection-related struvite crystals. Understanding these mechanisms is essential for developing targeted prevention and treatment strategies for this highly recurrent condition.'
) ON CONFLICT (doi) DO UPDATE
SET updated_at = NOW();

-- Article 4: Sulfamethoxazole-Induced Crystal Nephropathy (2024)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract
) VALUES (
    'Sulfamethoxazole-induced crystal nephropathy: characterization and prognosis in a case series',
    'Ruben Azencot, Camille Saint-Jacques, Jean-Philippe Haymann, Vincent Frochot, Michel Daudon, Emmanuel Letavernier',
    'Scientific Reports',
    '2024-03-13'::date,
    'en',
    'case_study',
    '10.1038/s41598-024-56322-9',
    '38480876',
    'This case series examined patients who developed acute kidney injury under sulfamethoxazole treatment to assess the causal relationship between the drug and kidney injury in the presence of N-acetyl-sulfamethoxazole (NASM) crystals in urine. Drug-induced crystalluria represents an important but underrecognized cause of acute kidney injury. The study characterized the clinical presentation, crystal morphology, and prognosis of sulfamethoxazole-induced crystal nephropathy. Findings emphasize the importance of early crystal identification in urine microscopy and prompt discontinuation of the causative medication to prevent irreversible renal dysfunction. Most patients showed renal function recovery after drug cessation, highlighting the generally favorable prognosis with timely intervention.'
) ON CONFLICT (doi) DO UPDATE
SET updated_at = NOW();

-- ========================================
-- STEP 2: Link Articles to Score Item
-- ========================================

-- Link Article 1 (Outpatient Crystalluria)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'ebcc36fd-d285-4754-adf7-50c7b130b286'::uuid
FROM articles a
WHERE a.doi = '10.18502/ijm.v17i3.18820'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 2 (Amoxicillin Crystalluria)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'ebcc36fd-d285-4754-adf7-50c7b130b286'::uuid
FROM articles a
WHERE a.doi = '10.1016/j.kint.2024.09.019'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 3 (Molecular Mechanisms)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'ebcc36fd-d285-4754-adf7-50c7b130b286'::uuid
FROM articles a
WHERE a.doi = '10.3390/ijms25053075'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 4 (Sulfamethoxazole Nephropathy)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'ebcc36fd-d285-4754-adf7-50c7b130b286'::uuid
FROM articles a
WHERE a.doi = '10.1038/s41598-024-56322-9'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- ========================================
-- STEP 3: Update Score Item with Clinical Content
-- ========================================

UPDATE score_items
SET
    clinical_relevance = 'A presença de cristais patológicos na urina representa um marcador importante de supersaturação urinária e risco elevado para formação de cálculos renais (nefrolitíase). Diferentemente dos cristais fisiológicos (oxalato de cálcio em pequenas quantidades, ácido úrico ocasional), cristais patológicos indicam condições metabólicas específicas ou exposição a medicamentos nefrotóxicos. Os principais tipos incluem: cistina (cistinúria hereditária), estruvita ou fosfato triplo (infecções urinárias por bactérias produtoras de urease como Proteus), 2,8-dihidroxiadenina (deficiência de adenina fosforribosiltransferase), e cristais induzidos por drogas (sulfametoxazol, amoxicilina, atazanavir, aciclovir, metotrexato). A prevalência de cristalúria em pacientes ambulatoriais atinge 22%, sendo oxalato de cálcio (46%), ácido úrico (23%), uratos (15%) e estruvita (9%) os mais comuns. A cristalúria por amoxicilina afeta 24-41% dos pacientes em doses altas (≥150 mg/kg/dia), podendo causar nefropatia cristalina aguda com necessidade de diálise em 10-40% dos casos. A detecção precoce é fundamental pois 50% dos pacientes com primeiro episódio de nefrolitíase apresentam recorrência em 5 anos. Cristais patológicos associam-se a diabetes, insuficiência renal, prostatite, síndrome nefrótica e ITU (10,6% dos casos com cristalúria). A identificação microscópica adequada permite intervenção terapêutica direcionada antes da formação de cálculos ou dano renal irreversível.',

    patient_explanation = 'Cristais patológicos na urina são formações sólidas anormais que indicam risco aumentado para "pedras nos rins" (cálculos renais). Imagine sua urina como água com vários minerais dissolvidos - quando há excesso desses minerais ou a urina fica muito concentrada, eles podem se juntar formando cristais microscópicos visíveis no exame de urina. Existem cristais normais que aparecem ocasionalmente, mas cristais patológicos são sempre preocupantes. Os tipos mais importantes incluem: cristais de cistina (doença genética rara), cristais de estruvita (indicam infecção urinária com bactérias específicas), e cristais causados por medicamentos como alguns antibióticos (sulfametoxazol, amoxicilina em doses altas), antivirais ou quimioterápicos. Pessoas com diabetes, problemas renais ou infecções urinárias de repetição têm maior risco. A boa notícia é que detectar cristais patológicos cedo permite tratamento antes que se formem pedras maiores ou ocorra lesão renal. O tratamento pode incluir: beber mais água (2-3 litros/dia) para diluir a urina, ajustar medicações que estão causando cristais, tratar infecções urinárias, modificar a alimentação (reduzir sal, proteína animal, oxalato) e em casos específicos usar remédios para alcalinizar ou acidificar a urina. Com acompanhamento adequado, é possível prevenir complicações graves como cálculos renais e insuficiência renal.',

    conduct = 'INVESTIGAÇÃO INICIAL: (1) Confirmar cristalúria em amostra fresca (<2h) com microscopia de luz polarizada e contraste de fase para caracterização morfológica precisa; (2) Coletar história detalhada: medicamentos em uso (antibióticos, antivirais, antirretrovirais, quimioterápicos), história familiar de cálculos renais ou doenças metabólicas, episódios prévios de nefrolitíase, sintomas de ITU, hábitos dietéticos (ingestão hídrica, consumo de proteína animal, sal, oxalato); (3) Exames laboratoriais: função renal (creatinina, ureia, TFG), EAS completo com sedimentoscopia, urinocultura se sinais de ITU, pH urinário, urina de 24h (cálcio, oxalato, citrato, ácido úrico, sódio, creatinina, volume total), eletrólitos séricos, gasometria venosa se suspeita de acidose tubular renal. CRISTAIS ESPECÍFICOS: Cistina - investigar cistinúria com teste de nitroprussiato, dosagem de cistina urinária, considerar teste genético; Estruvita - tratar ITU agressivamente, identificar foco estrutural (ecografia renal), antibioticoterapia guiada por cultura; Cristais medicamentosos - avaliar suspensão ou redução de dose do fármaco causador, hidratação EV vigorosa (3-4 L/dia) para aumentar fluxo urinário, alcalinização urinária com bicarbonato de sódio se apropriado. PREVENÇÃO DE CÁLCULOS: (1) Hidratação: objetivo >2,5 L de débito urinário/dia, orientar ingestão hídrica fracionada; (2) Modificação dietética: reduzir sódio (<2g/dia), proteína animal (<1g/kg/dia), manter cálcio dietético adequado (1000-1200mg/dia), reduzir oxalato (espinafre, beterraba, chocolate, nozes) se cristais de oxalato; (3) Farmacoterapia: citrato de potássio 30-60 mEq/dia para hipocitratúria, tiazídicos para hipercalciúria, alopurinol para hiperuricosúria; (4) Acompanhamento: EAS trimestral no primeiro ano, ecografia renal anual, reavaliação de urina 24h após 3-6 meses de intervenção. MONITORAMENTO: Pacientes com cristalúria patológica requerem seguimento em 3 meses para verificar resolução. Se persistência apesar de tratamento ou desenvolvimento de IRA, encaminhar para nefrologista.',

    updated_at = NOW()
WHERE id = 'ebcc36fd-d285-4754-adf7-50c7b130b286';

COMMIT;

-- ========================================
-- VERIFICATION QUERY
-- ========================================

SELECT
    si.id,
    si.name,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    (SELECT COUNT(*)
     FROM article_score_items asi
     WHERE asi.score_item_id = si.id) as linked_articles_count,
    si.updated_at
FROM score_items si
WHERE si.id = 'ebcc36fd-d285-4754-adf7-50c7b130b286';

-- Show linked articles
SELECT
    a.title,
    a.authors,
    a.journal,
    a.publish_date,
    a.doi,
    a.pm_id
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = 'ebcc36fd-d285-4754-adf7-50c7b130b286'
ORDER BY a.publish_date DESC;
