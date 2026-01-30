-- Enrichment for Score Item: Cor da Urina (ID: df525b04-d0b4-4b11-a5e1-198374bf32e1)
-- Generated: 2026-01-28
-- Articles: 4 peer-reviewed publications (2013-2025)

BEGIN;

-- Insert Article 1: Blue-Green Urine Diagnostic Approach (2025)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract,
    created_at,
    updated_at
) VALUES (
    'Diagnostic Approach to Abnormal Urine Colors: Lessons From a Case of Blue-Green Urine',
    'Carson Balen, Zayd Chishti, Jason W Wilson',
    'Cureus',
    '2025-04-11'::date,
    'en',
    'case_study',
    '10.7759/cureus.82122',
    '40357097',
    'Abnormal urine coloration can be a perplexing finding in clinical practice. This case report examines blue-green urine discoloration caused by methylene blue, identifying three main etiological categories: iatrogenic (medications like methylene blue, propofol, cimetidine), infectious (Pseudomonas aeruginosa producing pyocyanin), and metabolic/genetic (alkaptonuria, Hartnup disease). The diagnostic approach emphasizes clinical history, urinalysis, and recognition that normal urine ranges from pale yellow to amber due to urochrome. Awareness of medication effects can avoid unnecessary investigations.',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    publish_date = EXCLUDED.publish_date,
    pm_id = EXCLUDED.pm_id,
    abstract = EXCLUDED.abstract,
    updated_at = NOW();

-- Insert Article 2: Urinometer Technology and Color Analysis (2025)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract,
    created_at,
    updated_at
) VALUES (
    'Recent Advances in Urinometers: Enhancing Monitoring of Urine Output, pH, and Color: A Narrative Review',
    'Arati Raut, Ranjana Sharma, Anil Wanjari, Sheetal Mude, Samruddhi Gujar',
    'Journal of Pharmacy & Bioallied Sciences',
    '2025-03-01'::date,
    'en',
    'review',
    '10.4103/jpbs.jpbs_1705_24',
    '40511253',
    'Advancements in urinometer technology have revolutionized urine monitoring, providing clinicians with accurate, real-time data on urine output, pH, and color. Enhanced digital sensors, AI-driven data analytics, and integration with electronic health systems have transformed traditional urinometers into sophisticated diagnostic tools. Urine color analysis can indicate dehydration, hematuria, or other pathological conditions early. Standardized color grading systems ensure consistent and accurate color analysis across different clinical settings, supporting timely interventions and improved patient outcomes.',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    publish_date = EXCLUDED.publish_date,
    pm_id = EXCLUDED.pm_id,
    abstract = EXCLUDED.abstract,
    updated_at = NOW();

-- Insert Article 3: Hematuria Clinical Guidelines (2025)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract,
    created_at,
    updated_at
) VALUES (
    'Gross and Microscopic Hematuria',
    'Stephen W. Leslie, Karim Hamawy, Muhammad O. Saleem',
    'StatPearls',
    '2025-01-01'::date,
    'en',
    'review',
    NULL,
    '22710',
    'Comprehensive clinical guidelines for hematuria evaluation. Microhematuria is defined as 3 or more urinary RBCs per high-power field requiring microscopic confirmation. Risk stratification includes low-risk (<1% cancer risk), intermediate-risk (0.2-3.1% risk), and high-risk (1.3-6.3% risk) categories based on age, smoking history, and RBC count. Differential diagnosis distinguishes glomerular causes (proteinuria, dysmorphic RBCs, casts) from nonglomerular causes (urological pathology). Cystoscopy remains 98% sensitive for bladder cancer detection and is underutilized.',
    NOW(),
    NOW()
);

-- Insert Article 4: Hemoglobinuria vs Hematuria Differentiation (2013)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract,
    created_at,
    updated_at
) VALUES (
    'Hemoglobinuria Misidentified as Hematuria: Review of Discolored Urine and Paroxysmal Nocturnal Hemoglobinuria',
    'Prashant Veerreddy',
    'Clinical Medicine Insights: Blood Disorders',
    '2013-06-01'::date,
    'en',
    'review',
    '10.4137/CMBD.S11517',
    '25512715',
    'Comprehensive review on differentiating hemoglobinuria from hematuria. Hematuria shows dipstick positive with proportionate RBCs on microscopic urinalysis, while hemoglobinuria shows dipstick positive with disproportionately low or absent RBCs. When centrifuged, the sediment is red in hematuria versus red supernatant in hemoglobinuria, myoglobinuria, or drug-induced discoloration. Hemoglobinuria is easily confused with other causes of discolored urine, potentially resulting in extensive unnecessary urological workup. Case demonstrates importance of proper diagnostic differentiation.',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    publish_date = EXCLUDED.publish_date,
    pm_id = EXCLUDED.pm_id,
    abstract = EXCLUDED.abstract,
    updated_at = NOW();

-- Update Score Item with comprehensive clinical content
UPDATE score_items
SET
    clinical_relevance = 'A cor da urina é um parâmetro clínico fundamental que reflete múltiplos aspectos da fisiologia renal, estado de hidratação e presença de condições patológicas. A urina normal varia de amarelo pálido a âmbar devido ao pigmento urocromo, sendo a intensidade da cor diretamente proporcional ao estado de concentração urinária. Alterações na coloração urinária podem indicar desde situações benignas (desidratação, uso de medicamentos, ingestão de alimentos) até condições graves que requerem investigação imediata. A hematúria macroscópica (urina avermelhada) é particularmente significativa, pois pode sinalizar desde infecção urinária até neoplasias do trato urológico, com risco de câncer de bexiga variando de <1% em baixo risco até 6,3% em pacientes de alto risco (homens ≥60 anos, tabagistas >30 maços-ano, >25 hemácias/campo). A diferenciação entre hematúria verdadeira (presença proporcional de hemácias ao exame microscópico), hemoglobinúria (hemoglobina livre sem hemácias proporcionais, sugerindo hemólise intravascular) e pigmentúria (descoloração por pigmentos, medicamentos ou alimentos sem positividade para sangue) é crucial para direcionar a investigação adequada. Urina "cor de coca-cola" sugere hematúria de origem glomerular, enquanto urina vermelha clara indica sangramento do trato urinário inferior. Urina esverdeada pode indicar infecção por Pseudomonas aeruginosa (produção de piocianina) ou uso de medicamentos como azul de metileno. Avanços tecnológicos recentes incluem sistemas de gradação de cor padronizados integrados a urinômetros digitais, permitindo análise objetiva e consistente, eliminando a subjetividade da avaliação visual tradicional e possibilitando detecção precoce de desidratação, hematúria e outras condições patológicas.',

    patient_explanation = 'A cor da sua urina fornece informações importantes sobre seu estado de saúde e hidratação. Urina saudável normalmente varia de amarelo claro (quando você está bem hidratado) a amarelo mais escuro ou âmbar (quando precisa beber mais água). A cor vem principalmente de um pigmento chamado urocromo que seu corpo produz naturalmente. Mudanças na cor podem ser normais ou indicar problemas que precisam de atenção médica. Urina muito clara e quase transparente geralmente significa que você está bebendo bastante líquido. Urina amarelo escuro ou âmbar indica que você precisa se hidratar mais. Urina alaranjada pode ser causada por desidratação, vitaminas do complexo B ou certos medicamentos. Urina rosa ou avermelhada merece atenção especial: pode ser apenas por comer beterraba ou tomar certos remédios, mas também pode indicar presença de sangue (hematúria), que requer avaliação médica imediata. Urina marrom escura ou "cor de coca-cola" pode sugerir problemas no fígado, nos rins ou desidratação grave. Urina verde ou azulada é rara e geralmente relacionada a medicamentos específicos (como azul de metileno) ou infecções urinárias por bactérias especiais. Urina espumosa persistente pode indicar excesso de proteínas. Urina turva ou com mau cheiro frequentemente sugere infecção urinária. É importante informar seu médico sobre qualquer mudança persistente na cor da urina, especialmente se acompanhada de dor ao urinar, febre, dor nas costas ou na barriga. Mantenha-se bem hidratado bebendo água regularmente ao longo do dia. Se notar sangue na urina (mesmo que apareça apenas uma vez), procure avaliação médica, pois pode indicar desde infecções simples até condições mais sérias que precisam de investigação e tratamento adequados.',

    conduct = 'AVALIAÇÃO INICIAL: Realizar anamnese detalhada investigando tempo de início da alteração de cor, padrão (constante vs intermitente), sintomas associados (disúria, febre, dor lombar, dor abdominal), história medicamentosa completa (incluindo vitaminas, suplementos e fitoterápicos), ingestão alimentar recente (beterraba, corantes alimentares), histórico de tabagismo (quantificar maços-ano), exposição ocupacional a químicos, e histórico pessoal/familiar de neoplasias urológicas ou doenças renais. EXAME FÍSICO: Avaliar sinais vitais, realizar exame abdominal buscando massas ou dor à palpação, avaliar punho-percussão lombar (sinal de Giordano), examinar presença de edema periférico e icterícia. INVESTIGAÇÃO LABORATORIAL PRIMÁRIA: (1) Urina tipo I com sedimentoscopia: diferenciar hematúria verdadeira (≥3 hemácias/campo) de hemoglobinúria ou pigmentúria; avaliar presença de cilindros hemáticos (origem glomerular), dismorfismo eritrocitário, proteinúria, leucocitúria, cristais, bactérias. (2) Urocultura se suspeita de ITU. (3) Teste de centrifugação: sedimento vermelho indica hematúria; sobrenadante vermelho indica hemoglobinúria, mioglobinúria ou pigmentúria. ESTRATIFICAÇÃO DE RISCO para hematúria (diretrizes AUA/SUFU 2025): BAIXO RISCO (risco câncer <1%): mulheres <60 anos, homens <40 anos, não-tabagistas, 3-10 hemácias/campo - conduta: repetir urina I em 6 meses, investigar causas benignas. RISCO INTERMEDIÁRIO (0,2-3,1%): mulheres ≥60 anos, homens 40-59 anos, tabagistas 10-30 maços-ano, 10-25 hemácias/campo - conduta: ultrassom de rins e vias urinárias, considerar TC urografia, avaliar necessidade de cistoscopia. ALTO RISCO (1,3-6,3%): homens ≥60 anos, tabagistas >30 maços-ano, >25 hemácias/campo, hematúria macroscópica (qualquer idade) - conduta OBRIGATÓRIA: TC urografia (sensibilidade 95% para lesões renais) + cistoscopia (sensibilidade 98% para câncer de bexiga). DIFERENCIAÇÃO GLOMERULAR vs NÃO-GLOMERULAR: Origem glomerular (nefropatia): presença de cilindros hemáticos, hemácias dismórficas, proteinúria significativa (>500mg/24h), hipertensão recente, creatinina elevada - encaminhar para Nefrologia. Origem não-glomerular (urológica): hemácias normomórficas, ausência de cilindros/proteinúria - prosseguir investigação urológica. CORES ESPECÍFICAS E CONDUTAS: Vermelho/rosa: excluir menstruação, alimentos (beterraba, amora), medicamentos (rifampicina, fenazopiridina, doxorrubicina); se persistir após exclusão, tratar como hematúria. Marrom/cor de coca-cola: investigar glomerulonefrite aguda, rabdomiólise (dosagem CPK, mioglobinúria), hepatopatias (bilirrubinas, transaminases). Verde/azul: revisar medicações (azul de metileno, propofol, amitriptilina), investigar ITU por Pseudomonas (urocultura), considerar erros inatos do metabolismo se persistente (alkaptonúria, doença de Hartnup). Laranja: avaliar hidratação, medicamentos (rifampicina, nitrofurantoína, sulfassalazina), hepatopatias. Turva/leitosa: investigar ITU (urocultura), quilúria (filariose), piúria, fosfatúria. SEGUIMENTO: Pacientes com hematúria de baixo risco sem causa identificada: repetir urina I anualmente por 2 anos. Risco intermediário/alto: seguimento urológico conforme achados de imagem e cistoscopia. Se diagnóstico de ITU, reavaliar após tratamento para confirmar resolução. Hematúria persistente sem diagnóstico: considerar biópsia renal se suspeita glomerular. Educação ao paciente sobre sinais de alarme: dor intensa, febre, impossibilidade de urinar, coágulos volumosos. NOTAS CRÍTICAS: Cistoscopia permanece subutilizada apesar de sua alta sensibilidade; considerar em todos os casos de hematúria de risco intermediário/alto. Dipstick positivo para sangue com microscopia negativa sugere hemoglobinúria ou mioglobinúria, não hematúria verdadeira.',

    updated_at = NOW()
WHERE id = 'df525b04-d0b4-4b11-a5e1-198374bf32e1';

-- Link articles to the score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'df525b04-d0b4-4b11-a5e1-198374bf32e1'::uuid
FROM articles a
WHERE a.doi IN (
    '10.7759/cureus.82122',
    '10.4103/jpbs.jpbs_1705_24',
    '10.4137/CMBD.S11517'
)
OR a.pm_id = '22710'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

COMMIT;

-- Verification query
SELECT
    'Cor da Urina' as item_name,
    LENGTH(clinical_relevance) as clinical_relevance_chars,
    LENGTH(patient_explanation) as patient_explanation_chars,
    LENGTH(conduct) as conduct_chars,
    (SELECT COUNT(*) FROM article_score_items WHERE score_item_id = 'df525b04-d0b4-4b11-a5e1-198374bf32e1') as linked_articles_count
FROM score_items
WHERE id = 'df525b04-d0b4-4b11-a5e1-198374bf32e1';
