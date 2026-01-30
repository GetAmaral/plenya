BEGIN;

INSERT INTO articles (title, authors, journal, publish_date, language, article_type, doi, abstract, created_at, updated_at)
VALUES
(
    'Association between IGF-1 levels ranges and all-cause mortality: A meta-analysis',
    'Rahmani J, Montesanto A, Giovannucci E, et al.',
    'Aging Cell',
    '2022-01-20'::date,
    'en',
    'meta_analysis',
    '10.1111/acel.13540',
    'Meta-analysis of 19 prospective studies with 30,876 participants showing U-shaped relationship between IGF-1 and mortality. Both low and high levels increase mortality risk compared to middle range (120-160 ng/ml).',
    NOW(),
    NOW()
),
(
    'A 2024 Update on Growth Hormone Deficiency Syndrome in Adults: From Guidelines to Real Life',
    'Aversa LS, Cuboni D, Grottoli S, Ghigo E, Gasco V',
    'Journal of Clinical Medicine',
    '2024-10-12'::date,
    'en',
    'review',
    '10.3390/jcm13206079',
    'Systematic review highlighting critical gap: no stimulation test studies include patients >70 years. IGF-I monitoring challenges and heterogeneity in target levels across centers.',
    NOW(),
    NOW()
),
(
    'Association between sarcopenia and levels of growth hormone and insulin-like growth factor-1 in the elderly',
    'Chew J, Tay L, Lim JP, et al.',
    'BMC Musculoskeletal Disorders',
    '2020-04-07'::date,
    'en',
    'research_article',
    '10.1186/s12891-020-03236-y',
    'Cross-sectional study of 120 elderly (≥65 years) showing sarcopenia patients had significantly lower IGF-1 (89.8 vs 118.6 ng/ml). Each 10 ng/ml IGF-1 increase reduced sarcopenia risk by 18%.',
    NOW(),
    NOW()
),
(
    'ROLE of IGF-1 System in the Modulation of Longevity: Controversies and New Insights From a Centenarians Perspective',
    'Vitale G, Pellegrino G, Vollery M, Hofland LJ',
    'Frontiers in Endocrinology',
    '2019-02-01'::date,
    'en',
    'review',
    '10.3389/fendo.2019.00027',
    'Critical review of IGF-1/longevity paradox. Animal models show 38-70% lifespan extension with GH/IGF-1 reduction, but human centenarians show contradictory IGF-1 patterns.',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW();

UPDATE score_items
SET
    clinical_relevance = 'O IGF-1 (Fator de Crescimento Semelhante à Insulina tipo 1) em pacientes acima de 70 anos apresenta importância clínica complexa e controversa na medicina do envelhecimento. O declínio fisiológico do eixo GH/IGF-1 com a idade é universal, mas sua interpretação requer nuances críticas. Níveis baixos de IGF-1 em idosos ≥70 anos associam-se fortemente à sarcopenia, fragilidade, perda de massa muscular, redução de força de preensão e comprometimento da capacidade funcional. Meta-análises demonstram relação em formato de U entre IGF-1 e mortalidade: tanto níveis baixos (HR=1.33) quanto elevados (HR=1.23) aumentam risco comparado ao range médio (120-160 ng/ml). Estudos em centenários revelam o paradoxo da longevidade: enquanto modelos animais com redução de IGF-1 vivem 38-70% mais, humanos centenários apresentam perfis variáveis de IGF-1, desafiando a hipótese simples de \"quanto menor, melhor\". A deficiência de IGF-1 em idosos >70 anos correlaciona-se com perda acelerada de massa muscular apendicular, velocidade de marcha reduzida e maior prevalência de síndrome de fragilidade. Cada redução de 10 ng/ml em IGF-1 sérico aumenta em 22% o risco de sarcopenia (OR=1.22). Contudo, a terapia de reposição permanece controversa: estudos mostram benefícios em composição corporal e densidade óssea, mas sem comprovação de redução de mortalidade ou normalização de risco cardiovascular. Revisões sistemáticas de 2024 destacam lacuna crítica: nenhum estudo de testes estimulatórios de GH incluiu pacientes >70 anos, impossibilitando estabelecer cut-offs diagnósticos específicos para esta faixa etária. O envelhecimento associa-se não apenas à redução quantitativa de IGF-1, mas à resistência periférica aos seus efeitos anabólicos, comprometendo sinalização muscular e endotelial. Disfunção microvascular mediada por deficiência de IGF-1 contribui para hipoperfusão muscular e exacerbação de sarcopenia. A interpretação de níveis de IGF-1 em septuagenários e octogenários deve considerar contexto clínico: desnutrição proteico-calórica, inflamação crônica (IL-6, TNF-α), insuficiência hepática, hipotireoidismo e uso de corticoides suprimem IGF-1 independentemente do eixo GH. Modificações dietéticas (aumento de proteínas animais, laticínios) podem elevar IGF-1, mas o impacto em desfechos clínicos duros permanece incerto. A avaliação de IGF-1 >70 anos exige abordagem individualizada, ponderando risco de fragilidade versus ausência de evidências robustas para reposição hormonal nesta população vulnerável.',

    patient_explanation = 'O IGF-1, ou Fator de Crescimento Similar à Insulina, é um hormônio produzido principalmente pelo fígado em resposta ao hormônio do crescimento (GH). Após os 70 anos, os níveis de IGF-1 naturalmente diminuem como parte do processo de envelhecimento, mas essa queda pode ter impactos importantes na sua saúde. O IGF-1 funciona como um \"construtor\" no seu corpo: ele ajuda a manter seus músculos fortes, seus ossos densos e sua capacidade de se recuperar de lesões. Quando os níveis ficam muito baixos, você pode experimentar perda de massa muscular (sarcopenia), fraqueza para segurar objetos, dificuldade para caminhar e maior risco de quedas e fraturas. Estudos recentes mostram que pessoas com IGF-1 muito baixo têm 33% mais risco de morte prematura, mas curiosamente, níveis muito altos também aumentam o risco em 23% – formando o que os médicos chamam de \"curva em U\". O ideal parece estar no meio-termo, entre 120-160 ng/ml. Existe um paradoxo interessante: em animais de laboratório, níveis baixos de IGF-1 fazem viver mais tempo, mas em humanos centenários (pessoas com mais de 100 anos) os resultados são contraditórios – alguns têm IGF-1 baixo, outros normal. Isso significa que a relação entre IGF-1 e longevidade é mais complexa do que parece. Se seu médico solicitar este exame, ele está investigando possíveis causas de fraqueza muscular, perda de peso involuntária ou fragilidade. Valores baixos podem indicar desnutrição, inflamação crônica, problemas no fígado ou deficiência real de hormônio do crescimento. É importante saber que simplesmente repor IGF-1 com medicamentos não é recomendado para a maioria dos idosos, pois os benefícios não são claros e pode haver riscos. O tratamento, quando necessário, foca em melhorar a nutrição (especialmente proteínas), tratar doenças subjacentes, estimular atividade física supervisionada e, em casos muito selecionados, avaliar reposição de hormônio do crescimento. Lembre-se: níveis de IGF-1 devem ser interpretados junto com sua história clínica completa, estado nutricional, nível de atividade física e presença de outras condições médicas. Converse abertamente com seu geriatra sobre o que os resultados significam especificamente para você.',

    conduct = 'INTERPRETAÇÃO CLÍNICA EM PACIENTES ≥70 ANOS: A avaliação de IGF-1 em septuagenários, octogenários e nonagenários exige abordagem diferenciada dos adultos jovens, devido ao declínio fisiológico idade-dependente e ausência de valores de referência específicos validados. VALORES DE REFERÊNCIA AJUSTADOS: Utilizar sempre faixas específicas por idade e sexo, preferencialmente metodologia por espectrometria de massa (padrão-ouro). Para ≥70 anos, considerar: Homens: 60-150 ng/ml; Mulheres: 50-130 ng/ml. Atenção: cut-offs variam entre laboratórios e métodos (imunoensaio vs espectrometria de massa). INVESTIGAÇÃO DE NÍVEIS BAIXOS (<60 ng/ml): 1) Avaliar causas secundárias antes de atribuir à deficiência de GH: desnutrição proteico-calórica (albumina <3.5 g/dl), inflamação crônica (PCR >5 mg/L, IL-6 elevada), insuficiência hepática (transaminases, INR), hipotireoidismo (TSH, T4L), diabetes descompensado (HbA1c), uso de corticoides sistêmicos, anorexia em caquexia oncológica. 2) Reavaliar após 3-6 meses de correção de causas secundárias e otimização nutricional. 3) Testes estimulatórios de GH (teste de hipoglicemia insulínica, GHRH+arginina) NÃO são recomendados rotineiramente em >70 anos devido a: riscos cardiovasculares, ausência de cut-offs validados nesta faixa etária, baixa especificidade. CORRELAÇÃO CLÍNICA MANDATÓRIA: IGF-1 isolado tem valor limitado; sempre avaliar: 1) Sarcopenia: DEXA, BIA ou antropometria para massa muscular apendicular; força de preensão palmar (dinamometria); velocidade de marcha (4-6 metros); critérios EWGSOP2 ou FNIH. 2) Fragilidade: fenótipo de Fried (5 critérios); escala de fragilidade de Edmonton; timed up-and-go test. 3) Estado nutricional: MAN-SF (Mini Nutritional Assessment), ingestão proteica diária (recomendado >1.2 g/kg/dia para idosos), albumina sérica, pré-albumina. 4) Performance funcional: índice de Katz (AVDs), escala de Lawton-Brody (AIVDs). CONDUTA PARA IGF-1 BAIXO CONFIRMADO: 1) INTERVENÇÃO NUTRICIONAL (primeira linha): Aumentar ingestão proteica para 1.2-1.5 g/kg/dia, priorizando proteínas de alto valor biológico; suplementação de leucina (3g 3x/dia) se ingestão inadequada; vitamina D >30 ng/ml (otimização da função muscular). 2) EXERCÍCIO RESISTIDO SUPERVISIONADO: Treino de força 2-3x/semana demonstrou aumentar IGF-1 em 15-25% e melhorar desfechos funcionais independentemente da reposição hormonal. Considerar fisioterapia especializada em geriatria. 3) TRATAMENTO DE COMORBIDADES: Otimizar controle glicêmico, tratar hipotireoidismo, reduzir corticoides se possível, tratar depressão (associada a níveis baixos de IGF-1). 4) REPOSIÇÃO DE GH/IGF-1 (casos muito selecionados): INDICAÇÕES ESTRITAS em >70 anos: deficiência de GH documentada por teste estimulatório + IGF-1 <2 desvios-padrão + sarcopenia grave refratária + ausência de contraindicações. CONTRAINDICAÇÕES: neoplasia ativa ou história de câncer <5 anos (IGF-1 é mitogênico), retinopatia diabética proliferativa, hipertensão intracraniana benigna, insuficiência cardíaca descompensada. DOSE INICIAL: rhGH 0.1-0.2 mg/dia (metade da dose de adultos jovens), titulação lenta a cada 4-8 semanas conforme IGF-1 e tolerabilidade. ALVO TERAPÊUTICO: IGF-1 no quartil superior da normalidade para idade (NÃO suprafisiológico). MONITORAMENTO: glicemia de jejum e HbA1c (risco de diabetes), edema periférico, artralgias, avaliação oftalmológica semestral. IMPORTANTE: Evidência de benefício em desfechos duros (mortalidade, fraturas) é inexistente em >70 anos. Discussão de risco-benefício detalhada com paciente/família é mandatória. NÍVEIS ELEVADOS (>150-180 ng/ml em >70 anos): Incomum; investigar: acromegalia (GH elevado, alterações faciais, crescimento de extremidades); uso de GH exógeno (doping, \"anti-aging\"); tumores secretores de IGF-1 (raro). Solicitar: GH basal e pós-TOTG 75g (supressão inadequada sugere acromegalia), RM de sela túrcica, avaliação de complicações (hipertrofia cardíaca, apneia do sono, pólipos colônicos). ACOMPANHAMENTO LONGITUDINAL: Reavaliar IGF-1 anualmente em pacientes com sarcopenia/fragilidade, semestral se em reposição de GH. Priorizar desfechos funcionais (força, marcha, AVDs) sobre níveis hormonais isolados. A decisão de tratar deve sempre considerar: expectativa de vida (benefícios levam meses a anos), preferências do paciente, custo-efetividade (rhGH custa milhares de reais/mês), qualidade de vida. EVIDÊNCIAS LIMITADAS: Reconhecer francamente com paciente/família que dados em >70 anos são escassos, com grande heterogeneidade e ausência de ensaios clínicos randomizados robustos nesta faixa etária.',

    updated_at = NOW()
WHERE id = 'd1fd128f-7f92-4498-be3f-3bfe9564ce70';

INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    'd1fd128f-7f92-4498-be3f-3bfe9564ce70',
    id
FROM articles
WHERE doi IN (
    '10.1111/acel.13540',
    '10.3390/jcm13206079',
    '10.1186/s12891-020-03236-y',
    '10.3389/fendo.2019.00027'
)
AND created_at >= NOW() - INTERVAL '1 minute'
ON CONFLICT DO NOTHING;

COMMIT;

SELECT
    si.id,
    si.name,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    COUNT(sia.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN article_score_items sia ON si.id = sia.score_item_id
WHERE si.id = 'd1fd128f-7f92-4498-be3f-3bfe9564ce70'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct;
