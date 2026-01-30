-- Enrichment for IGF-1 (Somatomedina C) Score Item
-- ID: 039f1542-7596-4671-8ed0-049b3b41cfc4
-- Generated: 2026-01-29

BEGIN;

-- Insert scientific articles supporting IGF-1 clinical content
INSERT INTO articles (id, title, authors, journal, publish_date, doi, original_link, article_type, created_at, updated_at)
VALUES
  (
    gen_random_uuid(),
    'Insulin-Like Growth Factor-1 (IGF-1) and Its Monitoring in Medical Diagnostic and in Sports',
    'Bailes J, Soloviev M',
    'Biomolecules',
    '2021-02-04',
    '10.3390/biom11020217',
    'https://www.mdpi.com/2218-273X/11/2/217',
    'review',
    NOW(),
    NOW()
  ),
  (
    gen_random_uuid(),
    'Reference Intervals for Insulin-like Growth Factor-1 (IGF-I) From Birth to Senescence: Results From a Multicenter Study Using a New Automated Chemiluminescence IGF-I Immunoassay Conforming to Recent International Recommendations',
    'Bidlingmaier M, Friedrich N, Emeny RT, Spranger J, Wolthers OD, Roswall J, Körner A, Obermayer-Pietsch B, Hübener C, Dahlgren J, Frystyk J, Pfeiffer AF, Doering A, Bielohuby M, Wallaschofski H, Arafat AM',
    'J Clin Endocrinol Metab',
    '2014-05-01',
    '10.1210/jc.2013-3059',
    'https://academic.oup.com/jcem/article/99/5/1712/2537423',
    'research_article',
    NOW(),
    NOW()
  ),
  (
    gen_random_uuid(),
    'IGF1 for the diagnosis of growth hormone deficiency in children and adolescents: a reappraisal',
    'Ibba A, Corrias F, Guzzetti C, Casula L, Salerno M, di Iorgi N, Tornese G, Patti G, Radetti G, Maghnie M, Cappa M, Loche S',
    'Endocr Connect',
    '2020-11-01',
    '10.1530/EC-20-0347',
    'https://ec.bioscientifica.com/view/journals/ec/9/11/EC-20-0347.xml',
    'research_article',
    NOW(),
    NOW()
  ),
  (
    gen_random_uuid(),
    'Acromegaly: An Endocrine Society Clinical Practice Guideline',
    'Katznelson L, Laws ER Jr, Melmed S, Molitch ME, Murad MH, Utz A, Wass JA; Endocrine Society',
    'J Clin Endocrinol Metab',
    '2014-11-01',
    '10.1210/jc.2014-2700',
    'https://academic.oup.com/jcem/article/99/11/3933/2836347',
    'review',
    NOW(),
    NOW()
  )
ON CONFLICT (doi) DO NOTHING;

-- Update score_items with comprehensive clinical content
UPDATE score_items
SET
  clinical_relevance = 'O IGF-1 (Fator de Crescimento Semelhante à Insulina tipo 1), também conhecido como Somatomedina C, é o principal marcador bioquímico da ação do hormônio do crescimento (GH) e constitui ferramenta fundamental na avaliação do eixo somatotrófico. Produzido predominantemente no fígado sob estímulo do GH, o IGF-1 apresenta níveis séricos estáveis ao longo do dia, ao contrário da secreção pulsátil do GH, tornando-o o marcador de escolha para avaliação da produção e ação do hormônio do crescimento. Aproximadamente 98% do IGF-1 circulante está ligado a proteínas de ligação (IGFBPs), especialmente IGFBP-3, formando complexo ternário que modula sua biodisponibilidade. Clinicamente, é essencial no diagnóstico e monitoramento de deficiência de GH em crianças com baixa estatura e adultos, na avaliação de excesso de GH (acromegalia e gigantismo), no acompanhamento de terapia de reposição com GH recombinante, e na investigação de distúrbios do crescimento. Os valores de referência são fortemente dependentes da idade, com declínio após o nascimento, elevação progressiva até pico puberal (aos 15 anos), e declínio contínuo na vida adulta e senescência. Valores abaixo do percentil 2,5 (Z-score <-2) ajustados para idade sugerem deficiência ou resistência ao GH, enquanto valores elevados acima de 1,3 vezes o limite superior da normalidade em pacientes com características clínicas típicas confirmam acromegalia. A interpretação adequada requer valores de referência específicos para idade, método analítico utilizado, e consideração de fatores que afetam níveis de IGF-1 como estado nutricional, doenças crônicas, hipotireoidismo, doença hepática e renal.',

  patient_explanation = 'O IGF-1, também chamado de Somatomedina C, é uma proteína produzida principalmente pelo fígado quando estimulado pelo hormônio do crescimento (GH). Este exame mede a quantidade de IGF-1 no sangue e é muito importante para avaliar se o hormônio do crescimento está funcionando adequadamente no seu organismo. Diferente do GH que varia muito ao longo do dia, o IGF-1 se mantém mais estável, tornando-o um marcador mais confiável. O IGF-1 é fundamental para o crescimento normal na infância e adolescência, e continua tendo funções importantes na vida adulta, como manutenção da massa muscular, saúde óssea e metabolismo. Níveis baixos de IGF-1 podem indicar deficiência de hormônio do crescimento, que em crianças pode causar baixa estatura e em adultos pode causar cansaço, perda de massa muscular e óssea, aumento de gordura corporal e piora da qualidade de vida. Níveis muito altos podem indicar excesso de hormônio do crescimento, condição chamada acromegalia em adultos, que pode causar aumento das mãos, pés e características faciais. É importante saber que os valores normais variam muito com a idade: são baixos em bebês, aumentam durante a infância, atingem valores máximos na adolescência e diminuem gradualmente após os 20 anos. Por isso, o resultado sempre deve ser comparado com valores de referência específicos para sua faixa etária. O médico também considera outros fatores como nutrição, doenças crônicas e função da tireoide ao interpretar o resultado.',

  conduct = 'A avaliação clínica do IGF-1 exige interpretação cuidadosa no contexto de valores de referência específicos para idade, método laboratorial utilizado e condição clínica investigada. VALORES BAIXOS (IGF-1 < percentil 2,5 ou Z-score <-2): Em crianças com baixa estatura, solicitar também IGFBP-3 e considerar testes de estímulo de GH se suspeita de deficiência de GH, lembrando que IGF-1 isolado tem acurácia limitada (sensibilidade 67%, especificidade 63%, melhor em puberdade). Valores baixos também podem ocorrer em desnutrição, hipotireoidismo, doença hepática, doença renal crônica, diabetes mal controlado, e uso de glicocorticoides, exigindo investigação e correção destes fatores antes de diagnosticar deficiência de GH. Em adultos com sintomas sugestivos de deficiência de GH (fadiga, perda de massa muscular, osteoporose), considerar teste de estímulo de GH se IGF-1 baixo, especialmente em pacientes com histórico de lesão hipotálamo-hipofisária, cirurgia ou radioterapia craniana. VALORES ELEVADOS (IGF-1 > 1,3x LSN): Em pacientes com características clínicas de acromegalia (prognatismo, aumento de extremidades, cefaleia, hipertensão, diabetes), confirmar com teste de supressão de GH após sobrecarga oral de glicose (TOTG), onde GH que não suprime abaixo de 1 µg/L confirma diagnóstico. Solicitar ressonância magnética de sela túrcica para identificar adenoma hipofisário. No monitoramento de acromegalia tratada, o objetivo terapêutico é normalização do IGF-1 ajustado para idade e sexo, idealmente no terço inferior da faixa normal. MONITORAMENTO DE TERAPIA COM GH: Durante reposição de GH, o objetivo é manter IGF-1 na faixa normal, idealmente no terço médio-superior, com ajustes de dose baseados em níveis seriados de IGF-1. Evitar níveis suprafisiológicos que aumentam risco de efeitos adversos. CONSIDERAÇÕES ESPECIAIS: Na transição da adolescência para idade adulta, reteste é recomendado em pacientes com deficiência isolada idiopática de GH após completar crescimento longitudinal. Pacientes com deficiências hipofisárias múltiplas (≥3 deficiências) e IGF-1 <-2 DP não necessariamente precisam de reteste. Solicitar IGF-1 preferencialmente pela manhã, jejum não é obrigatório mas recomendado. Sempre utilizar valores de referência do método específico do laboratório.',

  updated_at = NOW()
WHERE id = '039f1542-7596-4671-8ed0-049b3b41cfc4';

-- Link articles to the score item
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
  '039f1542-7596-4671-8ed0-049b3b41cfc4',
  id
FROM articles
WHERE doi IN (
  '10.3390/biom11020217',
  '10.1210/jc.2013-3059',
  '10.1530/EC-20-0347',
  '10.1210/jc.2014-2700'
)
ON CONFLICT (score_item_id, article_id) DO NOTHING;

COMMIT;

-- Verification query
SELECT
  si.name,
  si.code,
  LENGTH(si.clinical_relevance) as clinical_relevance_chars,
  LENGTH(si.patient_explanation) as patient_explanation_chars,
  LENGTH(si.conduct) as conduct_chars,
  COUNT(asi.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '039f1542-7596-4671-8ed0-049b3b41cfc4'
GROUP BY si.id, si.name, si.code, si.clinical_relevance, si.patient_explanation, si.conduct;
