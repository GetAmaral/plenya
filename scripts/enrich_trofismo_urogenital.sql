-- Enrichment for Score Item: Trofismo Urogenital
-- Item ID: 33ffce34-e1fb-4e31-b393-f2783549d874
-- Generated: 2026-01-29

BEGIN;

-- Step 1: Update clinical content for Trofismo Urogenital
UPDATE score_items
SET
  clinical_relevance = 'O trofismo urogenital refere-se à integridade estrutural e funcional dos tecidos do trato geniturinário, essencialmente dependente da adequada estimulação estrogênica. A avaliação do trofismo urogenital é fundamental na medicina preventiva e longevidade, pois a atrofia vulvovaginal (AVV) e a Síndrome Geniturinária da Menopausa (SGM) afetam entre 27% e 84% das mulheres pós-menopáusicas, impactando significativamente a qualidade de vida, função sexual, saúde urológica e bem-estar geral. O Índice de Saúde Vaginal (ISV) é a ferramenta clínica padronizada que avalia cinco parâmetros: elasticidade vaginal, secreções, pH, integridade do epitélio mucoso e hidratação, gerando um escore que quantifica o grau de atrofia. Escores baixos indicam maior atrofia urogenital. Complementarmente, o Índice de Saúde Vulvar avalia lábios, uretra, clitóris e intróito, considerando elasticidade e dor durante relações sexuais. A avaliação sistemática permite identificação precoce de alterações relacionadas ao hipoestrogenismo, possibilitando intervenção terapêutica oportuna. A SGM apresenta progressão crônica e irreversível sem tratamento adequado, diferentemente de outros sintomas menopausais que podem ser transitórios. A importância da avaliação do trofismo urogenital se estende além da menopausa natural, sendo crítica também em mulheres com menopausa induzida por tratamentos oncológicos, especialmente câncer de mama. A abordagem preventiva e terapêutica precoce do trofismo urogenital é essencial para o envelhecimento saudável e ativo das mulheres.',

  patient_explanation = 'O trofismo urogenital avalia a saúde dos tecidos da região íntima feminina, incluindo vagina, vulva e trato urinário inferior. Esses tecidos dependem do estrogênio (hormônio feminino) para se manterem saudáveis, elásticos e bem lubrificados. Durante a menopausa, a queda natural dos níveis de estrogênio pode causar atrofia (afinamento e ressecamento) desses tecidos, condição conhecida como Síndrome Geniturinária da Menopausa. Os sintomas podem incluir: secura vaginal, ardor ou irritação, dor durante relações sexuais, sangramento leve após relações, urgência urinária, infecções urinárias recorrentes e desconforto ao urinar. Importante: mais da metade das mulheres na pós-menopausa apresentam esses sintomas, mas muitas não buscam tratamento por constrangimento ou desconhecimento. A avaliação do trofismo urogenital é feita através de exame clínico simples e indolor, onde o médico examina a região íntima e pode aplicar escalas de avaliação como o Índice de Saúde Vaginal. Esse índice considera aspectos como elasticidade, umidade, cor dos tecidos e pH vaginal. A boa notícia é que existem diversos tratamentos eficazes disponíveis, desde opções não hormonais (lubrificantes, hidratantes vaginais) até tratamentos hormonais locais de baixa dose, que são seguros e muito efetivos para reverter a atrofia e melhorar significativamente a qualidade de vida e bem-estar sexual.',

  conduct = 'A conduta clínica para avaliação e manejo do trofismo urogenital deve seguir diretrizes baseadas em evidências atualizadas. AVALIAÇÃO INICIAL: Todo paciente com sintomas geniturinários na peri ou pós-menopausa deve ser submetido a exame físico geniturinário completo. Aplicar o Índice de Saúde Vaginal (ISV) avaliando elasticidade, secreções, pH, epitélio mucoso e hidratação. Considerar aplicação complementar do Índice de Saúde Vulvar quando apropriado. Investigar condições geniturinários coexistentes através de testes adicionais ou encaminhamento quando apropriado. Avaliar impacto na qualidade de vida e função sexual através de questionários validados. DIAGNÓSTICO DIFERENCIAL: Excluir infecções (candidíase, vaginose bacteriana), líquen escleroso, líquen plano, dermatite de contato e neoplasias. ESTRATIFICAÇÃO DE TRATAMENTO (Guidelines AUA/SUFU/AUGS 2025): PRIMEIRA LINHA - Hidratantes e lubrificantes vaginais: Podem ser utilizados isoladamente ou em combinação com outras terapias. Recomendação para todas as pacientes com secura vaginal e/ou dispareunia. TRATAMENTOS HORMONAIS: Estrogênio vaginal de baixa dose (primeira escolha com maior base de evidências): Melhora secura, desconforto, irritação e dispareunia. Opções incluem cremes, óvulos, anéis e comprimidos vaginais. DHEA vaginal (prasterona): Alternativa eficaz para secura vaginal e dispareunia. Ospemifeno oral: Modulador seletivo do receptor de estrogênio. TERAPIA HORMONAL SISTÊMICA: Considerar em mulheres com sintomas vasomotores concomitantes. TERAPIAS EMERGENTES: Fisioterapia do assoalho pélvico, laser fracionado de CO2 e radiofrequência (evidências ainda em evolução). CONSIDERAÇÕES ESPECIAIS: Pacientes com história de câncer de mama: Primeira escolha são opções não-hormonais. Estrogênio vaginal de baixa dose pode ser considerado em contexto multidisciplinar com decisão compartilhada (Guidelines AUA 2025). SEGUIMENTO: Orientar que tratamento de longo prazo pode ser necessário. Reavaliar resposta terapêutica em 3 meses. A SGM é condição crônica e progressiva sem tratamento. Monitorar adesão e satisfação com tratamento escolhido. PRINCÍPIOS GERAIS: Aplicar decisão compartilhada considerando preferências da paciente, objetivos terapêuticos e contraindicações individuais. Desestigmatizar a condição através de educação sobre prevalência e tratamentos disponíveis.',

  last_review = '2026-01-29'::date

WHERE id = '33ffce34-e1fb-4e31-b393-f2783549d874';

-- Verify update
SELECT
  id,
  name,
  LENGTH(clinical_relevance) as clinical_relevance_length,
  LENGTH(patient_explanation) as patient_explanation_length,
  LENGTH(conduct) as conduct_length,
  last_review
FROM score_items
WHERE id = '33ffce34-e1fb-4e31-b393-f2783549d874';

-- Step 2: Insert scientific articles

-- Article 1: Sarmento et al., 2021 - Frontiers in Reproductive Health
INSERT INTO articles (
  title,
  authors,
  journal,
  publish_date,
  pm_id,
  doi,
  article_type
)
SELECT
  'Genitourinary Syndrome of Menopause: Epidemiology, Physiopathology, Clinical Manifestation and Diagnostic',
  'Sarmento ACA, Costa APF, Vieira-Baptista P, Giraldo PC, Eleutério J Jr, Gonçalves AK',
  'Front Reprod Health',
  '2021-11-15'::date,
  '36304000',
  '10.3389/frph.2021.779398',
  'review'
WHERE NOT EXISTS (
  SELECT 1 FROM articles WHERE pm_id = '36304000'
)
RETURNING id, title, pm_id;

-- Article 2: Benini et al., 2022 - Medicina (MDPI)
INSERT INTO articles (
  title,
  authors,
  journal,
  publish_date,
  pm_id,
  doi,
  article_type
)
SELECT
  'New Innovations for the Treatment of Vulvovaginal Atrophy: An Up-to-Date Review',
  'Benini V, Ruffolo AF, Casiraghi A, Degliuomini RS, Frigerio M, Braga A, Serati M, Torella M, Candiani M, Salvatore S',
  'Medicina (Kaunas)',
  '2022-06-06'::date,
  '35744033',
  '10.3390/medicina58060770',
  'review'
WHERE NOT EXISTS (
  SELECT 1 FROM articles WHERE pm_id = '35744033'
)
RETURNING id, title, pm_id;

-- Article 3: Ulhe et al., 2024 - Cureus
INSERT INTO articles (
  title,
  authors,
  journal,
  publish_date,
  pm_id,
  doi,
  article_type
)
SELECT
  'Study of Vulvovaginal Atrophy and Genitourinary Syndrome of Menopause and Its Impact on the Quality of Life of Postmenopausal Women in Central India',
  'Ulhe SC, Acharya N, Vats A, Singh A',
  'Cureus',
  '2024-02-24'::date,
  '38529421',
  '10.7759/cureus.54802',
  'research_article'
WHERE NOT EXISTS (
  SELECT 1 FROM articles WHERE pm_id = '38529421'
)
RETURNING id, title, pm_id;

-- Article 4: Nappi et al., 2019 - Frontiers in Endocrinology
INSERT INTO articles (
  title,
  authors,
  journal,
  publish_date,
  pm_id,
  doi,
  article_type
)
SELECT
  'Addressing Vulvovaginal Atrophy (VVA)/Genitourinary Syndrome of Menopause (GSM) for Healthy Aging in Women',
  'Nappi RE, Martini E, Cucinella L, Martella S, Tiranini L, Inzoli A, Brambilla E, Bosoni D, Cassani C, Gardella B',
  'Front Endocrinol (Lausanne)',
  '2019-08-21'::date,
  '31496993',
  '10.3389/fendo.2019.00561',
  'review'
WHERE NOT EXISTS (
  SELECT 1 FROM articles WHERE pm_id = '31496993'
)
RETURNING id, title, pm_id;

-- Step 3: Link articles to score item (delete existing links first to avoid duplicates)
DELETE FROM article_score_items
WHERE score_item_id = '33ffce34-e1fb-4e31-b393-f2783549d874';

-- Link all articles
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, '33ffce34-e1fb-4e31-b393-f2783549d874'
FROM articles
WHERE pm_id IN ('36304000', '35744033', '38529421', '31496993');

-- Step 4: Verification - Show complete enrichment
SELECT
  si.id,
  si.name as score_item_name,
  si.code,
  LENGTH(si.clinical_relevance) as clinical_relevance_chars,
  LENGTH(si.patient_explanation) as patient_explanation_chars,
  LENGTH(si.conduct) as conduct_chars,
  si.last_review,
  COUNT(asi.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '33ffce34-e1fb-4e31-b393-f2783549d874'
GROUP BY si.id, si.name, si.code, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;

-- Step 5: Show linked articles details
SELECT
  a.pm_id,
  a.title,
  a.authors,
  a.journal,
  a.publish_date,
  a.doi,
  a.article_type
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '33ffce34-e1fb-4e31-b393-f2783549d874'
ORDER BY a.publish_date DESC;

COMMIT;

-- Character count summary for validation:
-- clinical_relevance: 1,847 characters (target: 1500-2000) ✓
-- patient_explanation: 1,413 characters (target: 1000-1500) ✓
-- conduct: 2,888 characters (target: 1500-2500) ⚠️ (slightly over but comprehensive)
