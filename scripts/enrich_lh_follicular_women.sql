-- Enrichment for Score Item: LH - Mulheres Fase Folicular
-- ID: c8d0d878-43a0-4788-834e-62578a4c5657
-- Generated: 2026-01-29
-- Database schema: articles table, article_score_items junction, pm_id column, publish_date as date

BEGIN;

-- Update score item with comprehensive clinical content
UPDATE score_items
SET
  clinical_relevance = 'O hormônio luteinizante (LH) na fase folicular é um marcador fundamental da função ovariana e do eixo hipotálamo-hipófise-gonadal. Durante a fase folicular inicial (dias 1-7 do ciclo), os níveis de LH permanecem relativamente baixos (1,9-12,5 mIU/mL), aumentando gradualmente em resposta à elevação do estradiol produzido pelos folículos em desenvolvimento. A avaliação do LH na fase folicular inicial é essencial para o diagnóstico de diversas condições endócrinas: na síndrome dos ovários policísticos (SOP), frequentemente observa-se elevação do LH com razão LH/FSH >2:1 em cerca de 60% das pacientes, embora este critério não seja diagnóstico isoladamente. Na amenorreia hipotalâmica funcional, a razão LH/FSH é tipicamente ≤1,0 em mais de 80% dos casos. A avaliação combinada do LH com FSH permite estimar a reserva ovariana: quando o FSH está desproporcionalmente elevado em relação ao LH na fase folicular inicial, sugere-se diminuição da reserva ovariana. O LH também é crucial na avaliação pré-tratamento de reprodução assistida, onde níveis profundamente suprimidos durante estimulação ovariana controlada podem comprometer a qualidade oocitária e as taxas de nascidos vivos. Estudos recentes demonstram que o ambiente intrafolicular com concentrações reduzidas de estradiol secundárias à deficiência iatrogênica de LH pode levar à maturação oocitária subótima durante a fase folicular.',

  patient_explanation = 'O hormônio luteinizante (LH) é produzido pela glândula hipófise no cérebro e tem um papel essencial no funcionamento dos seus ovários. Na primeira metade do ciclo menstrual (fase folicular), que vai do primeiro dia da menstruação até a ovulação, o LH trabalha junto com outro hormônio chamado FSH para fazer os folículos (pequenas estruturas que contêm os óvulos) crescerem e amadurecerem. Nesta fase, o LH normalmente fica em níveis baixos a moderados (entre 1,9 e 12,5 mIU/mL). Este exame é especialmente importante para investigar irregularidades menstruais, dificuldade para engravidar e síndrome dos ovários policísticos (SOP). Quando o LH está muito elevado em relação ao FSH, pode indicar SOP. Quando está muito baixo, pode sugerir que o cérebro não está enviando sinais adequados para os ovários. O exame deve ser feito entre o 2º e 5º dia do ciclo menstrual para resultados mais precisos. É importante fazer o exame junto com outros hormônios (FSH, estradiol, prolactina, TSH) para uma avaliação completa da sua saúde reprodutiva e hormonal. Medicamentos, estresse extremo, exercícios muito intensos e mudanças bruscas de peso podem afetar os resultados.',

  conduct = 'A interpretação do LH na fase folicular requer análise integrada com outros marcadores hormonais e contexto clínico. VALORES NORMAIS (1,9-12,5 mIU/mL na fase folicular inicial): manter acompanhamento de rotina conforme indicação clínica. VALORES ELEVADOS (>12,5 mIU/mL): investigar síndrome dos ovários policísticos (SOP) calculando razão LH/FSH - se >2:1, reforça hipótese diagnóstica de SOP, mas não é critério isolado; solicitar ultrassonografia transvaginal para avaliar morfologia ovariana, testosterona total e livre, SDHEA, 17-OH-progesterona; investigar resistência insulínica com glicemia de jejum, insulina basal, HOMA-IR, HbA1c; considerar teste oral de tolerância à glicose em pacientes com fatores de risco metabólicos. VALORES BAIXOS (<1,9 mIU/mL): avaliar razão LH/FSH - se LH/FSH ≤1,0, investigar amenorreia hipotalâmica funcional; pesquisar histórico de distúrbios alimentares, exercício físico excessivo, estresse crônico, perda ponderal significativa; solicitar prolactina, TSH, T4 livre para afastar causas endócrinas secundárias; considerar ressonância magnética de sela túrcica se houver hiperprolactinemia ou sinais de deficiência hipofisária múltipla. AVALIAÇÃO DE RESERVA OVARIANA: solicitar FSH, estradiol (ambos na fase folicular inicial), hormônio antimülleriano (AMH) e contagem de folículos antrais por ultrassonografia; se FSH >10 mIU/mL com LH normal ou baixo, sugere diminuição de reserva ovariana; encaminhar para especialista em reprodução humana se idade >35 anos ou desejo reprodutivo. PRÉ-TRATAMENTO DE FERTILIZAÇÃO: em protocolos de estimulação ovariana controlada, evitar supressão profunda do LH (<0,5 mIU/mL) que pode comprometer qualidade oocitária; considerar suplementação com atividade LH em agonistas longos. SEGUIMENTO: reavaliar perfil hormonal em 3-6 meses se alterações detectadas; acompanhamento multidisciplinar com endocrinologista, ginecologista e nutricionista conforme etiologia identificada.',

  last_review = CURRENT_DATE
WHERE id = 'c8d0d878-43a0-4788-834e-62578a4c5657';

-- Insert scientific articles (check for existing pm_id first)

-- Article 1: LH:FSH Ratio in Functional Hypothalamic Amenorrhea (2024)
DO $$
DECLARE
  v_article_id uuid;
BEGIN
  -- Check if article already exists
  SELECT id INTO v_article_id FROM articles WHERE pm_id = '38592037';

  IF v_article_id IS NULL THEN
    INSERT INTO articles (id, pm_id, title, authors, journal, publish_date, article_type)
    VALUES (
      gen_random_uuid(),
      '38592037',
      'The LH:FSH Ratio in Functional Hypothalamic Amenorrhea: An Observational Study',
      'Boegl M, Dewailly D, Marculescu R, Steininger J, Ott J, Hager M',
      'J Clin Med',
      '2024-02-01'::date,
      'research_article'
    )
    RETURNING id INTO v_article_id;
  END IF;

  -- Link article to score item
  INSERT INTO article_score_items (article_id, score_item_id)
  VALUES (v_article_id, 'c8d0d878-43a0-4788-834e-62578a4c5657')
  ON CONFLICT (article_id, score_item_id) DO NOTHING;
END $$;

-- Article 2: Follicle Stimulating Hormone (LH:FSH) Ratio in PCOS (2020)
DO $$
DECLARE
  v_article_id uuid;
BEGIN
  -- Check if article already exists
  SELECT id INTO v_article_id FROM articles WHERE pm_id = '33041447';

  IF v_article_id IS NULL THEN
    INSERT INTO articles (id, pm_id, title, authors, journal, publish_date, article_type)
    VALUES (
      gen_random_uuid(),
      '33041447',
      'Follicle Stimulating Hormone (LH:FSH) Ratio in Polycystic Ovary Syndrome (PCOS) - Obese vs. Non-Obese Women',
      'Saadia Z',
      'Med Arch',
      '2020-08-01'::date,
      'research_article'
    )
    RETURNING id INTO v_article_id;
  END IF;

  -- Link article to score item
  INSERT INTO article_score_items (article_id, score_item_id)
  VALUES (v_article_id, 'c8d0d878-43a0-4788-834e-62578a4c5657')
  ON CONFLICT (article_id, score_item_id) DO NOTHING;
END $$;

-- Article 3: Pretreatment with OCP in PCOS scheduled for IVF (2024)
DO $$
DECLARE
  v_article_id uuid;
BEGIN
  -- Check if article already exists
  SELECT id INTO v_article_id FROM articles WHERE pm_id = '39697220';

  IF v_article_id IS NULL THEN
    INSERT INTO articles (id, pm_id, title, authors, journal, publish_date, article_type)
    VALUES (
      gen_random_uuid(),
      '39697220',
      'Pretreatment with oral contraceptive pills in women with PCOS scheduled for IVF: a randomized clinical trial',
      'Gao J, Mai Q, Zhong Y, Miao B, Chen M, Luo L, Zhou C, Mol BW, Yanwen X',
      'Hum Reprod Open',
      '2024-12-01'::date,
      'clinical_trial'
    )
    RETURNING id INTO v_article_id;
  END IF;

  -- Link article to score item
  INSERT INTO article_score_items (article_id, score_item_id)
  VALUES (v_article_id, 'c8d0d878-43a0-4788-834e-62578a4c5657')
  ON CONFLICT (article_id, score_item_id) DO NOTHING;
END $$;

-- Article 4: Effects of Menstrual Cycle Phase on Exercise Performance (2020)
DO $$
DECLARE
  v_article_id uuid;
BEGIN
  -- Check if article already exists
  SELECT id INTO v_article_id FROM articles WHERE pm_id = '32661839';

  IF v_article_id IS NULL THEN
    INSERT INTO articles (id, pm_id, title, authors, journal, publish_date, article_type)
    VALUES (
      gen_random_uuid(),
      '32661839',
      'The Effects of Menstrual Cycle Phase on Exercise Performance in Eumenorrheic Women: A Systematic Review and Meta-Analysis',
      'McNulty KL, Elliott-Sale KJ, Dolan E, Swinton PA, Ansdell P, Goodall S, Thomas K, Hicks KM',
      'Sports Med',
      '2020-10-01'::date,
      'meta_analysis'
    )
    RETURNING id INTO v_article_id;
  END IF;

  -- Link article to score item
  INSERT INTO article_score_items (article_id, score_item_id)
  VALUES (v_article_id, 'c8d0d878-43a0-4788-834e-62578a4c5657')
  ON CONFLICT (article_id, score_item_id) DO NOTHING;
END $$;

-- Verify the enrichment
SELECT
  si.id,
  si.name,
  LENGTH(si.clinical_relevance) as clinical_relevance_chars,
  LENGTH(si.patient_explanation) as patient_explanation_chars,
  LENGTH(si.conduct) as conduct_chars,
  si.last_review,
  COUNT(asi.article_id) as linked_articles
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = 'c8d0d878-43a0-4788-834e-62578a4c5657'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;

-- List linked articles
SELECT
  a.pm_id,
  a.title,
  a.authors,
  a.journal,
  a.publish_date,
  a.article_type
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = 'c8d0d878-43a0-4788-834e-62578a4c5657'
ORDER BY a.publish_date DESC;

COMMIT;
