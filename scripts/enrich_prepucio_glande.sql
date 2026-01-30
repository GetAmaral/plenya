-- Enrichment for Score Item: Prepúcio / glande (Foreskin / Glans)
-- Item ID: 5ddc4af9-dd19-4ea1-8117-3d3e30377dab
-- Generated: 2026-01-29

BEGIN;

-- Insert scientific articles
INSERT INTO articles (id, title, authors, journal, pm_id, doi, article_type, publish_date)
VALUES
  (
    gen_random_uuid(),
    'Foreskin care: Hygiene, importance of counselling, and management of common complications',
    'Leeson C, Vigil H, Witherspoon L',
    'Canadian Family Physician',
    '39965976',
    '10.46747/cfp.710297',
    'review',
    '2025-02-01'::date
  ),
  (
    gen_random_uuid(),
    'Lichen sclerosus: The 2023 update',
    'De Luca DA, Papara C, Vorobyev A, Staiger H, Bieber K, Thaçi D, Ludwig RJ',
    'Frontiers in Medicine',
    '36873861',
    '10.3389/fmed.2023.1106318',
    'review',
    '2023-02-16'::date
  ),
  (
    gen_random_uuid(),
    'Penile lichen sclerosus, circumcision and sequelae, what are the questions?',
    'Morrel B, t Hoen LA, Pasmans SGMA',
    'Translational Andrology and Urology',
    '35958904',
    '10.21037/tau-22-343',
    'review',
    '2022-07-01'::date
  ),
  (
    gen_random_uuid(),
    'Balanitis',
    'Wray AA, Velasquez J, Leslie SW, Khetarpal S',
    'StatPearls Publishing',
    '30725828',
    '10.StatPearls.30725828',
    'review',
    '2024-08-01'::date
  )
ON CONFLICT (doi) DO UPDATE SET
  pm_id = EXCLUDED.pm_id,
  title = EXCLUDED.title,
  authors = EXCLUDED.authors,
  journal = EXCLUDED.journal,
  publish_date = EXCLUDED.publish_date;

-- Link articles to score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '5ddc4af9-dd19-4ea1-8117-3d3e30377dab'::uuid
FROM articles a
WHERE a.pm_id IN ('39965976', '36873861', '35958904', '30725828')
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Update score item with comprehensive clinical content
UPDATE score_items
SET
  clinical_relevance = 'A avaliação do prepúcio e glande é fundamental no exame urogenital masculino, permitindo identificar condições que variam desde fisiológicas até patológicas graves. A fimose, definida como incapacidade de retrair o prepúcio sobre a glande, pode ser fisiológica (comum em crianças até 5-6 anos) ou patológica (resultante de processos inflamatórios ou cicatriciais). A distinção é crítica: fimose fisiológica geralmente resolve espontaneamente, enquanto fimose patológica requer intervenção. A balanite (inflamação da glande) e balanopostite (incluindo prepúcio) são condições infecciosas frequentes, com Candida como agente mais comum, especialmente em diabéticos. O líquen escleroso (balanite xerótica obliterante) afeta 32% das crianças com fimose e apresenta risco de estenose meatal (7-19% dos casos) e carcinoma peniano se não tratado adequadamente. A higiene inadequada favorece acúmulo de esmegma, inflamações recorrentes e aderências prepuciais. Em adultos, balanite recorrente pode indicar diabetes oculto, exigindo investigação metabólica. A inspeção cuidadosa identifica lesões hipopigmentadas, placas atróficas, cicatrizes, parafimose (prepúcio retraído que não retorna) e lesões suspeitas de malignidade. A circuncisão é curativa em >75% dos casos de líquen escleroso em estágios iniciais, mas circuncisões parciais apresentam 50% de recorrência. O envolvimento de glande e meato são fatores prognósticos negativos para estenoses uretrais. A avaliação criteriosa do prepúcio e glande permite diagnóstico precoce, prevenção de complicações, orientação sobre higiene adequada e indicação cirúrgica quando necessária, impactando significativamente qualidade de vida e saúde urogenital masculina.',

  patient_explanation = 'O prepúcio é a pele móvel que cobre a glande (cabeça do pênis). Avaliar essa região é importante para identificar problemas comuns e prevenir complicações. A fimose (dificuldade ou impossibilidade de retrair o prepúcio) é normal em crianças pequenas, mas pode ser problemática em adultos, causando dor, dificuldade de higiene e infecções. Inflamações como balanite (da glande) são frequentes, especialmente se houver diabetes ou higiene inadequada, causando vermelhidão, inchaço, dor e secreção. O líquen escleroso é uma condição inflamatória crônica que causa manchas brancas, endurecimento da pele e pode estreitar a abertura uretral, sendo mais comum do que se pensava (1 em 3 crianças com fimose). A higiene correta é fundamental: retrair suavemente o prepúcio durante o banho e enxaguar apenas com água corrente, sem sabonetes fortes que irritam a região. Infecções recorrentes podem indicar diabetes não diagnosticado e merecem investigação. Em alguns casos, a cirurgia de circuncisão (remoção do prepúcio) é necessária e resolve a maioria dos problemas, especialmente quando tratamentos com pomadas não funcionam. A avaliação médica regular permite identificar condições precocemente, orientar cuidados adequados e prevenir complicações sérias como estreitamento da uretra ou lesões pré-malignas. Qualquer alteração persistente, dor, dificuldade de urinar ou lesões na região deve ser avaliada por profissional de saúde.',

  conduct = 'AVALIAÇÃO CLÍNICA INICIAL: História detalhada investigando sintomas (dor, prurido, dificuldade de retração, disúria, secreção), idade de início, episódios prévios de inflamação, higiene genital, comorbidades (diabetes, imunossupressão), história sexual e possível exposição a DSTs. Exame físico com inspeção cuidadosa do prepúcio (mobilidade, aderências, cicatrizes, anel fimótico) e glande (lesões, hiperemia, secreção, placas brancas atróficas sugestivas de líquen escleroso). Avaliar meato uretral (calibre, estenose) e linfonodos inguinais. Em crianças, diferenciar fimose fisiológica de patológica: na fisiológica o prepúcio tem aparência saudável; na patológica há cicatrizes, anel fibrótico ou sinais de líquen escleroso. MANEJO CONSERVADOR: Fimose fisiológica em crianças: orientação familiar sobre retração suave diária durante banho com água morna, sem forçar. Corticoides tópicos (betametasona 0,05-0,1% 2x/dia por 4-8 semanas) são seguros e eficazes como primeira linha para fimose fisiológica, com taxa de sucesso 67-95%. Balanite aguda: antifúngicos tópicos (clotrimazol, miconazol 2x/dia por 7-14 dias) associados a corticoides de baixa potência para candidíase; antibióticos tópicos (mupirocina) para infecções bacterianas. Higiene adequada: orientar retração suave do prepúcio e enxágue apenas com água corrente, evitando sabonetes irritantes. Líquen escleroso inicial: corticoides potentes (clobetasol 0,05% 1-2x/dia por 2-3 meses) podem ser tentados em casos leves pediátricos, mas resposta é variável. INDICAÇÕES CIRÚRGICAS: Circuncisão indicada em fimose patológica (após falha de tratamento conservador), parafimose recorrente, balanite/balanopostite recorrente (>2-3 episódios), líquen escleroso (preferencial quando há envolvimento prepucial), lesões suspeitas de malignidade. Circuncisão completa é essencial: circuncisões parciais apresentam 50% de recorrência em líquen escleroso. Prepucioplastia com corticoides intralesionais é alternativa que preserva prepúcio, mas dados de longo prazo são limitados. SEGUIMENTO: Pacientes com balanite recorrente devem ser investigados para diabetes (glicemia de jejum, HbA1c). Líquen escleroso requer seguimento prolongado pelo risco de estenose meatal (7-19%) e estenose uretral proximal, especialmente se houver envolvimento de glande/meato. Reavaliação periódica para detecção de lesões pré-malignas ou carcinoma peniano em pacientes com líquen escleroso de longa data. Educação contínua sobre higiene adequada e sinais de alerta. ENCAMINHAMENTO ESPECIALIZADO: Urologia/cirurgia pediátrica para casos cirúrgicos, líquen escleroso com envolvimento meatal/uretral, lesões suspeitas de malignidade, fimose patológica complexa ou parafimose. Dermatologia para casos de líquen escleroso extenso ou refratário.',

  last_review = CURRENT_DATE
WHERE id = '5ddc4af9-dd19-4ea1-8117-3d3e30377dab';

-- Verify character counts
SELECT
  'Prepúcio / glande' as item,
  LENGTH(clinical_relevance) as clinical_chars,
  LENGTH(patient_explanation) as patient_chars,
  LENGTH(conduct) as conduct_chars,
  (SELECT COUNT(*) FROM article_score_items asi
   JOIN articles a ON asi.article_id = a.id
   WHERE asi.score_item_id = '5ddc4af9-dd19-4ea1-8117-3d3e30377dab') as linked_articles
FROM score_items
WHERE id = '5ddc4af9-dd19-4ea1-8117-3d3e30377dab';

COMMIT;

-- Final verification query
SELECT
  si.name,
  si.last_review,
  LENGTH(si.clinical_relevance) as clinical_relevance_length,
  LENGTH(si.patient_explanation) as patient_explanation_length,
  LENGTH(si.conduct) as conduct_length,
  COUNT(DISTINCT a.id) as article_count,
  STRING_AGG(a.title, ' | ' ORDER BY a.publish_date DESC) as article_titles
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
LEFT JOIN articles a ON asi.article_id = a.id
WHERE si.id = '5ddc4af9-dd19-4ea1-8117-3d3e30377dab'
GROUP BY si.id, si.name, si.last_review, si.clinical_relevance, si.patient_explanation, si.conduct;
