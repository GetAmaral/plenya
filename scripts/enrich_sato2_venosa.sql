-- ============================================================================
-- ENRICHMENT: SatO2 Venosa (Venous Oxygen Saturation)
-- Score Item ID: f3792ebc-d50c-448d-97af-8b43a9ac41a6
-- Date: 2026-01-29
-- ============================================================================

BEGIN;

-- ----------------------------------------------------------------------------
-- 1. UPDATE CLINICAL CONTENT
-- ----------------------------------------------------------------------------

UPDATE score_items
SET
  clinical_relevance = 'A saturação venosa de oxigênio (SvO2) representa um parâmetro hemodinâmico essencial que reflete o equilíbrio entre oferta e consumo de oxigênio nos tecidos. Medida no sangue venoso que retorna ao coração direito após perfundir todo o organismo, a SvO2 fornece informações críticas sobre a adequação da oxigenação tecidual sistêmica. Valores normais variam entre 60-80%, sendo que reduções indicam aumento da extração de oxigênio pelos tecidos, geralmente por redução do débito cardíaco, anemia, hipoxemia arterial ou aumento do consumo metabólico. Em medicina intensiva, a monitorização da SvO2 é particularmente valiosa em situações de instabilidade hemodinâmica, incluindo choque de diversas etiologias, insuficiência cardíaca descompensada, sepse e período perioperatório de cirurgias de grande porte. A saturação venosa mista (medida na artéria pulmonar via cateter de Swan-Ganz) representa a extração global de oxigênio, enquanto a saturação venosa central (ScvO2, medida na veia cava superior) oferece alternativa menos invasiva e amplamente utilizada. Estudos demonstram correlação consistente entre valores anormais de SvO2 e maior morbimortalidade, tornando-a ferramenta prognóstica importante. Na insuficiência cardíaca crônica, pacientes podem adaptar-se a valores reduzidos (30-40%), sendo que quedas agudas indicam deterioração cardíaca. O monitoramento contínuo permite detecção precoce de comprometimento hemodinâmico, orientando intervenções terapêuticas como otimização volêmica, uso de inotrópicos, transfusões sanguíneas e ajustes ventilatórios. A interpretação adequada requer análise integrada com outros parâmetros como débito cardíaco, lactato, diferença arteriovenosa de CO2 e saturação arterial de oxigênio.',

  patient_explanation = 'A saturação venosa de oxigênio é um exame que mede quanto oxigênio ainda resta no sangue depois que ele circulou por todo o corpo e entregou oxigênio aos órgãos e tecidos. Imagine que o sangue sai do coração "cheio" de oxigênio (cerca de 95-100%) e, ao passar pelos tecidos, vai "entregando" esse oxigênio. Quando retorna ao coração, normalmente ainda contém 60-80% de oxigênio, o que é saudável e indica que seu coração está bombeando bem e seus tecidos estão recebendo oxigênio adequado. Se esse valor estiver baixo (abaixo de 60%), pode significar que seu coração não está bombeando sangue suficiente, que você está anêmico, ou que seus tecidos estão consumindo mais oxigênio que o normal (como na febre ou infecções graves). Esse exame é especialmente importante em situações críticas, como internação em UTI, insuficiência cardíaca grave, cirurgias de grande porte ou quando você está muito doente. Ele ajuda os médicos a entenderem se seu corpo está recebendo oxigênio suficiente e a ajustarem rapidamente os tratamentos necessários, como soros, medicações para o coração, transfusões de sangue ou ajustes no respirador. É medido através de cateteres especiais colocados em veias grandes, geralmente no pescoço ou na artéria pulmonar. Valores muito baixos indicam sofrimento dos tecidos e necessidade de intervenção urgente, enquanto valores normais tranquilizam que a circulação está adequada.',

  conduct = 'A abordagem da saturação venosa de oxigênio alterada exige análise sistemática dos componentes do transporte de oxigênio: débito cardíaco, concentração de hemoglobina, saturação arterial de oxigênio e consumo metabólico. VALORES BAIXOS (<60% para SvO2 ou <70% para ScvO2): Avaliar débito cardíaco - se reduzido, considerar reposição volêmica guiada por parâmetros de responsividade (variação de pressão de pulso, elevação passiva de pernas), uso de inotrópicos (dobutamina, milrinona) ou vasopressores conforme perfil hemodinâmico. Verificar hemoglobina - transfundir concentrados de hemácias se Hb <7 g/dL (ou <9 g/dL em cardiopatas graves), visando melhorar capacidade de transporte de oxigênio. Otimizar saturação arterial - ajustar fração inspirada de oxigênio (FiO2) e parâmetros ventilatórios para manter SaO2 >92%, tratar causas de hipoxemia (SDRA, pneumonia, edema pulmonar, embolia pulmonar). Controlar consumo metabólico - analgesia e sedação adequadas, controle de febre com antitérmicos, tratar tremores ou agitação, considerar bloqueio neuromuscular em casos refratários. VALORES ELEVADOS (>80%): Podem indicar shunt arteriovenoso (sepse, cirrose), intoxicação por cianeto, hipotermia ou erro de medição. Na sepse, valores elevados refletem incapacidade de extração tecidual de oxigênio e não indicam adequação perfusional. MONITORAMENTO: Tendências são mais importantes que valores isolados. Quedas agudas de 5-10% indicam deterioração hemodinâmica e requerem investigação imediata. Em cirurgias de alto risco, manutenção de SvO2 >65-70% como meta terapêutica associa-se a redução de complicações pós-operatórias. Na insuficiência cardíaca, SvO2 <60% na admissão prediz pior prognóstico em 30 dias. Integrar sempre com outros marcadores: lactato sérico (hipoperfusão), diferença veno-arterial de CO2 (se >6 mmHg sugere inadequação de fluxo), débito urinário e perfusão periférica. Reavaliações frequentes (a cada 1-4 horas conforme gravidade) orientam ajustes terapêuticos e identificam respondedores às intervenções.',

  last_review = CURRENT_DATE

WHERE id = 'f3792ebc-d50c-448d-97af-8b43a9ac41a6';

-- Verify character counts
DO $$
DECLARE
  v_clinical_relevance_length INTEGER;
  v_patient_explanation_length INTEGER;
  v_conduct_length INTEGER;
BEGIN
  SELECT
    LENGTH(clinical_relevance),
    LENGTH(patient_explanation),
    LENGTH(conduct)
  INTO
    v_clinical_relevance_length,
    v_patient_explanation_length,
    v_conduct_length
  FROM score_items
  WHERE id = 'f3792ebc-d50c-448d-97af-8b43a9ac41a6';

  RAISE NOTICE '=== CHARACTER COUNT VERIFICATION ===';
  RAISE NOTICE 'clinical_relevance: % chars (target: 1500-2000)', v_clinical_relevance_length;
  RAISE NOTICE 'patient_explanation: % chars (target: 1000-1500)', v_patient_explanation_length;
  RAISE NOTICE 'conduct: % chars (target: 1500-2500)', v_conduct_length;
END $$;

-- ----------------------------------------------------------------------------
-- 2. INSERT SCIENTIFIC ARTICLES
-- ----------------------------------------------------------------------------

-- First, delete any existing links for this score item to avoid duplicates
DELETE FROM article_score_items
WHERE score_item_id = 'f3792ebc-d50c-448d-97af-8b43a9ac41a6';

-- Article 1: Venous Oxygen Saturation (StatPearls 2024)
WITH article_insert AS (
  INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    pm_id,
    doi,
    article_type,
    publish_date,
    language
  )
  SELECT
    gen_random_uuid(),
    'Venous Oxygen Saturation',
    'Shanmukhappa SC, Lokeshwaran S',
    'StatPearls [Internet]',
    '33232065',
    NULL,
    'review',
    '2024-09-10'::date,
    'en'
  WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '33232065'
  )
  RETURNING id
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, 'f3792ebc-d50c-448d-97af-8b43a9ac41a6'::uuid
FROM article_insert
UNION ALL
SELECT a.id, 'f3792ebc-d50c-448d-97af-8b43a9ac41a6'::uuid
FROM articles a
WHERE a.pm_id = '33232065' AND NOT EXISTS (SELECT 1 FROM article_insert);

-- Article 2: Anesthesia Monitoring of Mixed Venous Saturation (StatPearls 2023)
WITH article_insert AS (
  INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    pm_id,
    doi,
    article_type,
    publish_date,
    language
  )
  SELECT
    gen_random_uuid(),
    'Anesthesia Monitoring of Mixed Venous Saturation',
    'Lee CP, Bora V',
    'StatPearls [Internet]',
    '30969657',
    NULL,
    'review',
    '2023-03-19'::date,
    'en'
  WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '30969657'
  )
  RETURNING id
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, 'f3792ebc-d50c-448d-97af-8b43a9ac41a6'::uuid
FROM article_insert
UNION ALL
SELECT a.id, 'f3792ebc-d50c-448d-97af-8b43a9ac41a6'::uuid
FROM articles a
WHERE a.pm_id = '30969657' AND NOT EXISTS (SELECT 1 FROM article_insert);

-- Article 3: Capnodynamics noninvasive cardiac output and mixed venous oxygen saturation (Frontiers Pediatrics 2023)
WITH article_insert AS (
  INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    pm_id,
    doi,
    article_type,
    publish_date,
    language
  )
  SELECT
    gen_random_uuid(),
    'Capnodynamics - noninvasive cardiac output and mixed venous oxygen saturation monitoring in children',
    'Leino A, Kuusela T, Kukkonen S',
    'Frontiers in Pediatrics',
    '36896976',
    '10.3389/fped.2023.1111270',
    'research_article',
    '2023-02-23'::date,
    'en'
  WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '36896976'
  )
  RETURNING id
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, 'f3792ebc-d50c-448d-97af-8b43a9ac41a6'::uuid
FROM article_insert
UNION ALL
SELECT a.id, 'f3792ebc-d50c-448d-97af-8b43a9ac41a6'::uuid
FROM articles a
WHERE a.pm_id = '36896976' AND NOT EXISTS (SELECT 1 FROM article_insert);

-- Article 4: Monitoring of Tissue Oxygenation (Frontiers Medicine 2018)
WITH article_insert AS (
  INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    pm_id,
    doi,
    article_type,
    publish_date,
    language
  )
  SELECT
    gen_random_uuid(),
    'Monitoring of Tissue Oxygenation: an Everyday Clinical Challenge',
    'Lima A, Bakker J',
    'Frontiers in Medicine',
    '29379781',
    '10.3389/fmed.2017.00247',
    'review',
    '2018-01-08'::date,
    'en'
  WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '29379781'
  )
  RETURNING id
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, 'f3792ebc-d50c-448d-97af-8b43a9ac41a6'::uuid
FROM article_insert
UNION ALL
SELECT a.id, 'f3792ebc-d50c-448d-97af-8b43a9ac41a6'::uuid
FROM articles a
WHERE a.pm_id = '29379781' AND NOT EXISTS (SELECT 1 FROM article_insert);

-- ----------------------------------------------------------------------------
-- 3. VERIFICATION QUERIES
-- ----------------------------------------------------------------------------

\echo '\n=== ENRICHMENT VERIFICATION FOR SatO2 Venosa ==='

SELECT
  id,
  name,
  LENGTH(clinical_relevance) as clinical_relevance_chars,
  LENGTH(patient_explanation) as patient_explanation_chars,
  LENGTH(conduct) as conduct_chars,
  last_review
FROM score_items
WHERE id = 'f3792ebc-d50c-448d-97af-8b43a9ac41a6';

\echo '\n=== LINKED SCIENTIFIC ARTICLES ==='

SELECT
  a.title,
  a.authors,
  a.journal,
  a.pm_id,
  a.doi,
  a.article_type,
  a.publish_date
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = 'f3792ebc-d50c-448d-97af-8b43a9ac41a6'
ORDER BY a.publish_date DESC;

COMMIT;

\echo '\n=== ✅ ENRICHMENT COMPLETED SUCCESSFULLY ==='
