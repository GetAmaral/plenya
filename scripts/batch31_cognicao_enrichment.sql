-- ============================================================================
-- BATCH 31: ENRIQUECIMENTO COGNIÇÃO - 6 ITEMS
-- Data: 2026-01-27
-- Baseado em: Literatura científica e artigos MFI - Psiquiatria Metabólica
-- ============================================================================

-- ----------------------------------------------------------------------------
-- ITEM 1: Testes rápidos de memória
-- ID: e8416da2-f19b-4929-813b-9ead82399476
-- ----------------------------------------------------------------------------

UPDATE score_items
SET
  clinical_relevance = 'Os testes rápidos de memória são ferramentas fundamentais na avaliação cognitiva inicial e no screening de declínio cognitivo leve (DCL) ou demências em estágios iniciais. A avaliação da memória episódica (eventos recentes), memória de trabalho (span de dígitos) e memória semântica (conhecimento geral) permite identificar alterações sutis que precedem manifestações clínicas mais evidentes. Na medicina funcional integrativa, a abordagem vai além do diagnóstico nosológico, investigando fatores modificáveis que impactam a cognição: deficiências nutricionais (vitamina B12, folato, ômega-3), alterações metabólicas (resistência insulínica, disfunção tireoidiana), inflamação crônica de baixo grau, disbiose intestinal (eixo intestino-cérebro), distúrbios do sono, estresse crônico com hipercortisolismo, e exposição a toxinas ambientais. A neuroplasticidade permite reversibilidade parcial de déficits cognitivos quando os fatores causais são identificados e corrigidos precocemente. A avaliação seriada permite monitorar resposta a intervenções terapêuticas multimodais.',

  patient_explanation = 'Os testes rápidos de memória são como um "check-up" do funcionamento do seu cérebro. Eles ajudam a identificar se há dificuldades para lembrar informações recentes, manter a concentração ou processar novos conhecimentos. Esses testes são importantes porque muitos fatores do dia a dia podem afetar sua memória: alimentação inadequada, falta de vitaminas, sono ruim, estresse prolongado, problemas intestinais ou alterações hormonais. A boa notícia é que o cérebro tem capacidade de se recuperar quando corrigimos esses fatores. Através de mudanças na alimentação, suplementação adequada, melhora do sono e controle do estresse, é possível melhorar significativamente a memória e a clareza mental. Estes testes também servem para acompanhar sua evolução ao longo do tratamento.',

  conduct = 'Realizar avaliação cognitiva inicial com Mini-Mental, MoCA ou testes específicos de memória. Investigar: hemograma completo, perfil tiroideano (TSH, T3 livre, T4 livre), vitamina B12, ácido fólico, vitamina D, homocisteína, glicemia e hemoglobina glicada, perfil lipídico completo, proteína C reativa ultrassensível. Avaliar qualidade do sono (polissonografia se indicado), rastreio de apneia. Investigar disbiose intestinal e permeabilidade aumentada. Suplementação direcionada conforme deficiências identificadas. Estimular neuroplasticidade através de exercícios cognitivos, atividade física regular (mínimo 150min/semana), dieta anti-inflamatória mediterrânea, controle de estresse (mindfulness, meditação). Reavaliação a cada 3-6 meses.',

  updated_at = NOW()
WHERE id = 'e8416da2-f19b-4929-813b-9ead82399476';

-- Linkar artigos relevantes
INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('bc8d38a2-f684-40ce-8859-04f32570136a', 'e8416da2-f19b-4929-813b-9ead82399476'), -- MFI Psiquiatria Aula 01
  ('edcd8bcc-4b08-4021-b842-caee99b2e632', 'e8416da2-f19b-4929-813b-9ead82399476'), -- MFI Psiquiatria Aula 02
  ('664a4773-d083-4295-a7ec-0c7156f5457e', 'e8416da2-f19b-4929-813b-9ead82399476')  -- Abordagem Funcional Parte I
ON CONFLICT DO NOTHING;

-- ----------------------------------------------------------------------------
-- ITEM 2: Span de dígitos - Direto:___/8
-- ID: 10449bea-87a3-4fa3-a163-84fde09b062f
-- ----------------------------------------------------------------------------

UPDATE score_items
SET
  clinical_relevance = 'O teste de span de dígitos direto avalia especificamente a memória de trabalho auditiva imediata e a capacidade de atenção sustentada. Valores normais variam de 5 a 8 dígitos, sendo que resultados abaixo de 5 sugerem comprometimento da atenção e memória de curto prazo. Este teste é sensível para detectar alterações em lobo frontal e circuitos atencionais, sendo impactado por fadiga mental, privação de sono, transtorno de déficit de atenção (TDA/TDAH), ansiedade, depressão, e declínio cognitivo inicial. Na perspectiva funcional, diversos fatores metabólicos influenciam o desempenho: hipoglicemia ou hiperglicemia (variabilidade glicêmica), deficiência de magnésio (cofator em mais de 300 reações enzimáticas neurais), depleção de neurotransmissores (acetilcolina, dopamina, noradrenalina), estresse oxidativo neuronal, neuroinflamação, e exposição a metais pesados (chumbo, mercúrio, alumínio). A suplementação de precursores de neurotransmissores e otimização metabólica frequentemente melhoram o desempenho neste teste.',

  patient_explanation = 'O teste de span de dígitos direto mede sua capacidade de prestar atenção e guardar informações por alguns segundos - como lembrar um número de telefone que acabou de ouvir. Se você consegue repetir menos de 5 números, pode significar dificuldade de concentração ou sobrecarga mental. Diversos fatores afetam esse resultado: cansaço mental, noites mal dormidas, ansiedade, oscilações do açúcar no sangue, falta de magnésio (mineral importante para o cérebro) ou estresse crônico. Melhorar esses aspectos pode aumentar significativamente sua capacidade de atenção e memória imediata. O tratamento inclui ajustes na alimentação para estabilizar o açúcar no sangue, suplementos específicos para o cérebro, melhora na qualidade do sono e técnicas para reduzir o estresse.',

  conduct = 'Aplicar teste padronizado de span de dígitos direto. Score < 5: investigação aprofundada. Avaliar: glicemia de jejum e pós-prandial (curva glicêmica se indicado), magnésio intracelular (eritrocitário), zinco, cobre, vitaminas do complexo B (B1, B6, B12), ferritina e saturação de transferrina. Rastreio de metais pesados (metalotioneína, mineralograma capilar). Avaliar neurotransmissores através de marcadores indiretos ou testes funcionais. Suplementação: magnésio quelado 300-600mg/dia, complexo B ativo (metilfolato, metilcobalamina), ômega-3 EPA/DHA 2-3g/dia, fosfatidilserina 300mg/dia. Otimizar sono (higiene do sono, melatonina se necessário). Exercícios cognitivos específicos para atenção. Eliminar disruptores (cafeína excessiva, álcool). Retestar após 8-12 semanas.',

  updated_at = NOW()
WHERE id = '10449bea-87a3-4fa3-a163-84fde09b062f';

-- Linkar artigos relevantes
INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('bc8d38a2-f684-40ce-8859-04f32570136a', '10449bea-87a3-4fa3-a163-84fde09b062f'), -- MFI Psiquiatria Aula 01
  ('3df8e817-39cb-4828-8a17-eda0369dba2e', '10449bea-87a3-4fa3-a163-84fde09b062f'), -- MFI Psiquiatria Aula 03
  ('d5369fca-38bc-4cbb-9742-2090b2789892', '10449bea-87a3-4fa3-a163-84fde09b062f')  -- Abordagem Funcional Parte II
ON CONFLICT DO NOTHING;

-- ----------------------------------------------------------------------------
-- ITEM 3: Span de dígitos - Inverso:___/7
-- ID: 4e2aff36-a3c5-4934-b23d-98bd66bc03ee
-- ----------------------------------------------------------------------------

UPDATE score_items
SET
  clinical_relevance = 'O teste de span de dígitos inverso avalia funções executivas mais complexas que o direto, exigindo manipulação mental ativa da informação (memória de trabalho central executiva) e não apenas retenção passiva. Depende primariamente da integridade do córtex pré-frontal dorsolateral e circuitos frontoestriatais. Valores normais situam-se entre 4 e 7 dígitos. Déficits neste teste são marcadores precoces de disfunção executiva, comuns em TDAH, síndrome disexecutiva frontal, demência frontotemporal inicial, e transtornos neuropsiquiátricos. Na abordagem funcional integrativa, identifica-se que este domínio cognitivo é particularmente vulnerável a: hipoperfusão cerebral (anemia, hipotensão, disfunção vascular), hipóxia crônica (apneia do sono), inflamação sistêmica (citocinas pró-inflamatórias IL-1β, IL-6, TNF-α atravessam barreira hematoencefálica), depleção de dopamina e noradrenalina (precursores: tirosina, fenilalanina), déficit de acetilcolina, e toxicidade por organofosforados ou solventes orgânicos.',

  patient_explanation = 'O teste de span de dígitos inverso avalia uma capacidade mental mais complexa: você precisa não só lembrar os números, mas reorganizá-los mentalmente de trás para frente. É como fazer um malabarismo mental. Essa habilidade está ligada ao planejamento, organização e tomada de decisões do dia a dia. Se você tem dificuldade neste teste (consegue menos de 4 números), pode explicar por que às vezes é difícil organizar tarefas, priorizar atividades ou fazer várias coisas ao mesmo tempo. Fatores como má circulação sanguínea no cérebro, ronco com pausas na respiração durante o sono, inflamação no corpo, ou deficiência de substâncias importantes para o funcionamento cerebral podem prejudicar essa função. O tratamento foca em melhorar esses aspectos através de mudanças alimentares, exercícios físicos, suplementos específicos e treino cognitivo.',

  conduct = 'Aplicar teste de span inverso padronizado. Score < 4: investigação de disfunção executiva. Avaliar: hemograma completo (anemia), pressão arterial (hipotensão ou hipertensão), PCR ultrassensível, IL-6 se disponível, homocisteína (marcador de inflamação vascular), perfil lipídico com apolipoproteínas. Investigar apneia do sono (Questionário de Berlim, polissonografia). Avaliar reserva de catecolaminas: dosagem de tirosina plasmática, metabólitos urinários. Suplementação: N-acetilcisteína 600-1200mg/dia (precursor glutationa, antioxidante), acetil-L-carnitina 1500-3000mg/dia (função mitocondrial), L-tirosina 500-1500mg/dia (precursor dopamina/noradrenalina), colina 500mg/dia ou alpha-GPC 300mg/dia (precursor acetilcolina). Exercício aeróbico regular melhora perfusão cerebral. Treino cognitivo de funções executivas. Reavaliar 12 semanas.',

  updated_at = NOW()
WHERE id = '4e2aff36-a3c5-4934-b23d-98bd66bc03ee';

-- Linkar artigos relevantes
INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('bc8d38a2-f684-40ce-8859-04f32570136a', '4e2aff36-a3c5-4934-b23d-98bd66bc03ee'), -- MFI Psiquiatria Aula 01
  ('0dd08507-9078-4c09-9a3a-972f19ccbd3c', '4e2aff36-a3c5-4934-b23d-98bd66bc03ee'), -- MFI Psiquiatria Aula 05
  ('e73ad009-bf8d-4b2b-a37e-d29978f4533f', '4e2aff36-a3c5-4934-b23d-98bd66bc03ee')  -- Abordagem Funcional Parte III
ON CONFLICT DO NOTHING;

-- ----------------------------------------------------------------------------
-- ITEM 4: Presença de sintomas de tristeza/depressão/ansiedade
-- ID: 7fbdd31b-3368-42fd-8568-8da14cf20af8
-- ----------------------------------------------------------------------------

UPDATE score_items
SET
  clinical_relevance = 'Sintomas depressivos e ansiosos frequentemente coexistem com déficits cognitivos, configurando uma tríade que impacta significativamente qualidade de vida e funcionalidade. A neurociência moderna demonstra que depressão não é apenas "deficiência de serotonina", mas um transtorno inflamatório, metabólico e neurodegenerativo complexo. Pacientes deprimidos apresentam níveis elevados de citocinas pró-inflamatórias, ativação da via das quinureninas (desvio do triptofano para ácido quinolínico neurotóxico ao invés de serotonina), disfunção mitocondrial, estresse oxidativo aumentado, redução de BDNF (fator neurotrófico), atrofia hipocampal, e disbiose intestinal com aumento de permeabilidade (endotoxemia metabólica). A medicina funcional integrativa aborda causas raiz: deficiências nutricionais (ômega-3, vitamina D, magnésio, zinco, SAMe, metilfolato), disfunção tireoidiana subclínica, resistência insulínica, hipoglicemia reativa, alergias alimentares, toxinas (glifosato, metais pesados), e trauma psicológico não processado. Estudos demonstram eficácia de intervenções nutricionais e suplementares comparáveis a antidepressivos em casos leves a moderados.',

  patient_explanation = 'Sentimentos de tristeza, ansiedade ou depressão estão frequentemente ligados ao funcionamento do corpo todo, não apenas da mente. Pesquisas recentes mostram que a depressão tem muito a ver com inflamação no corpo, problemas no intestino, falta de vitaminas e minerais importantes, alterações no açúcar do sangue e desequilíbrios hormonais. O intestino, por exemplo, produz a maior parte da serotonina (substância do bem-estar) - quando ele não funciona bem, afeta diretamente seu humor. Deficiências de ômega-3, vitamina D, magnésio e vitaminas do complexo B também podem causar sintomas emocionais. O tratamento integrativo não substitui acompanhamento psicológico ou psiquiátrico quando necessário, mas complementa identificando e corrigindo esses fatores físicos que afetam suas emoções, frequentemente melhorando significativamente os sintomas.',

  conduct = 'Aplicar escalas validadas: PHQ-9 (depressão), GAD-7 (ansiedade), HAM-D ou Beck quando apropriado. Score elevado: avaliação psiquiátrica e psicológica. Investigação laboratorial abrangente: hemograma, TSH, T3 livre, T4 livre, anti-TPO (tireoidite), cortisol salivar em 4 pontos, vitamina D (ideal > 40ng/mL), vitamina B12 (ideal > 500pg/mL), ácido fólico, ferritina (ideal > 50ng/mL), zinco, magnésio eritrocitário, homocisteína, PCR ultrassensível, glicemia e insulina de jejum (HOMA-IR). Avaliar: disbiose (calprotectina fecal, SIBO), permeabilidade intestinal (zonulina), neurotransmissores (teste orgânico urinário). Protocolo suplementar: ômega-3 EPA > 2g/dia, vitamina D3 5000-10000UI/dia, magnésio 400-600mg/dia, complexo B metilado, probióticos específicos (Lactobacillus helveticus, Bifidobacterium longum), SAMe 400-800mg/dia, 5-HTP ou L-triptofano conforme caso. Psicoterapia cognitivo-comportamental. Exercício físico regular. Reavaliação mensal.',

  updated_at = NOW()
WHERE id = '7fbdd31b-3368-42fd-8568-8da14cf20af8';

-- Linkar artigos relevantes
INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('bc8d38a2-f684-40ce-8859-04f32570136a', '7fbdd31b-3368-42fd-8568-8da14cf20af8'), -- MFI Psiquiatria Aula 01
  ('edcd8bcc-4b08-4021-b842-caee99b2e632', '7fbdd31b-3368-42fd-8568-8da14cf20af8'), -- MFI Psiquiatria Aula 02
  ('3df8e817-39cb-4828-8a17-eda0369dba2e', '7fbdd31b-3368-42fd-8568-8da14cf20af8'), -- MFI Psiquiatria Aula 03
  ('05898fe5-54de-4c22-a56a-aca088c6706a', '7fbdd31b-3368-42fd-8568-8da14cf20af8')  -- MFI Psiquiatria Aula 06
ON CONFLICT DO NOTHING;

-- ----------------------------------------------------------------------------
-- ITEM 5: Mulheres: percepção de mudanças conforme época do ciclo menstrual ou proximidade menopausa
-- ID: bc22be14-aa28-4fa0-964b-652831c3cd95
-- ----------------------------------------------------------------------------

UPDATE score_items
SET
  clinical_relevance = 'Flutuações hormonais ao longo do ciclo menstrual e durante a transição menopausal exercem efeitos profundos sobre cognição, humor e bem-estar. Estrogênio e progesterona são neuroesteroides que modulam neurotransmissão, neuroplasticidade, perfusão cerebral e metabolismo energético neuronal. Na fase lútea tardia (pré-menstrual), a queda abrupta de estrogênio e progesterona correlaciona-se com piora de sintomas depressivos/ansiosos (síndrome pré-menstrual ou transtorno disfórico pré-menstrual - TDPM), declínio de memória de trabalho e função executiva. Durante perimenopausa e menopausa, a redução de estradiol está associada a "brain fog" (névoa mental), esquecimentos, dificuldade de concentração, insônia, fogachos, e aumento de risco para depressão. Estudos mostram que estrogênio aumenta produção de serotonina, modula receptores GABA e NMDA, promove neurogênese hipocampal, e tem efeito neuroprotetor antioxidante. A abordagem funcional integrativa investiga dominância estrogênica relativa (relação estrogênio/progesterona), metabolismo hepático de estrogênios (via 2-OH protetora vs 16-OH proliferativa), função tireoidiana (frequentemente alterada na menopausa), e suporte de vias de detoxificação.',

  patient_explanation = 'Muitas mulheres percebem que sua memória, concentração e humor variam ao longo do mês ou pioram na transição para a menopausa. Isso acontece porque os hormônios femininos (estrogênio e progesterona) não afetam apenas o útero - eles têm papel fundamental no funcionamento do cérebro. Esses hormônios ajudam na produção de substâncias que regulam o humor, melhoram a memória e protegem as células cerebrais. Quando seus níveis oscilam muito (antes da menstruação) ou caem (na menopausa), você pode experimentar "névoa mental", esquecimentos, dificuldade de concentração, alterações de humor, insônia e ondas de calor. Entender esse padrão é importante porque permite tratamentos específicos que equilibram os hormônios naturalmente, através de alimentação adequada, suplementos, fitoterapia e, quando necessário, reposição hormonal bioidêntica, melhorando significativamente esses sintomas.',

  conduct = 'Anamnese detalhada: padrão menstrual, sintomas pré-menstruais, sintomas climatéricos (escala de Kupperman/Greene). Avaliar hormonais: dia 3 do ciclo (FSH, LH, estradiol, progesterona, prolactina), ou perfil hormonal salivar ao longo do ciclo. Avaliar: DHEA-S, cortisol, TSH, T3 livre, T4 livre, anti-TPO. Metabolismo de estrogênios (urina de 24h: 2-OH, 4-OH, 16-OH-estrone). Sintomas pré-menstruais: magnésio 400-600mg/dia, vitamina B6 100mg/dia (todo ciclo), vitex agnus-castus 40mg (aumenta progesterona), cálcio 1200mg/dia. Menopausa/perimenopausa: fitoestrógenos (isoflavonas 40-80mg/dia, linhaça), Black cohosh 40-80mg/dia, terapia hormonal bioidêntica quando indicado (estradiol transdérmico + progesterona micronizada). Suporte hepático: DIM 200-300mg/dia ou I3C, cardo mariano, complexo B. Exercício regular. Acompanhamento ginecológico. Reavaliação 3 meses.',

  updated_at = NOW()
WHERE id = 'bc22be14-aa28-4fa0-964b-652831c3cd95';

-- Linkar artigos relevantes
INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('721299e2-7131-417e-9da5-e7e509c088eb', 'bc22be14-aa28-4fa0-964b-652831c3cd95'), -- Abordagem Funcional Parte IV
  ('70298603-a991-437d-80e6-d15af039b80d', 'bc22be14-aa28-4fa0-964b-652831c3cd95'), -- Abordagem Funcional Parte V
  ('a472d98c-ebd8-4b59-bc62-0beddce19126', 'bc22be14-aa28-4fa0-964b-652831c3cd95')  -- Abordagem Funcional Parte VI
ON CONFLICT DO NOTHING;

-- ----------------------------------------------------------------------------
-- ITEM 6: Uso atual de psicotrópicos para cognição
-- ID: 36d39438-5337-4e06-b67d-edfaa498ba25
-- ----------------------------------------------------------------------------

UPDATE score_items
SET
  clinical_relevance = 'O uso de medicamentos psicotrópicos para cognição engloba: nootrópicos (piracetam, modafinil), inibidores de colinesterase (donepezil, rivastigmina, galantamina - para demências), memantina (modulador NMDA), metilfenidato e lisdexanfetamina (TDAH), além de uso off-label de antidepressivos e ansiolíticos. É fundamental documentar indicação, dose, tempo de uso, eficácia percebida e efeitos adversos. Na abordagem funcional integrativa, questiona-se: o medicamento trata causa raiz ou apenas sintoma? Muitos pacientes usam estimulantes para compensar privação crônica de sono, ou antidepressivos para sintomas causados por deficiências nutricionais corrigíveis. Além disso, medicamentos psicotrópicos frequentemente causam: depleção nutricional (antipsicóticos reduzem CoQ10, estatinas reduzem CoQ10, anticonvulsivantes reduzem folato e B12), disfunção mitocondrial, disbiose intestinal, ganho de peso e resistência insulínica. A estratégia integrativa busca otimizar fatores modificáveis para permitir, quando possível e seguro, redução gradual ou descontinuação sob supervisão médica rigorosa.',

  patient_explanation = 'Se você usa medicações para melhorar memória, concentração ou tratar depressão/ansiedade, é importante documentar o que usa, há quanto tempo e se está funcionando. Enquanto esses remédios podem ser necessários e ajudar, a medicina integrativa busca entender por que você precisou deles. Muitas vezes, fatores como falta de vitaminas, alterações intestinais, problemas de sono ou desequilíbrios hormonais causam sintomas que levam à prescrição desses medicamentos. Além disso, alguns medicamentos podem causar efeitos a longo prazo, como depleção de nutrientes importantes para o cérebro. O objetivo não é simplesmente parar as medicações - isso só deve ser feito com acompanhamento médico cuidadoso - mas identificar e corrigir problemas de base para que, eventualmente, você possa usar doses menores ou, em alguns casos, não precisar mais deles.',

  conduct = 'Documentar: medicamento(s), dose, frequência, tempo de uso, indicação original, profissional prescritor, eficácia (escala 0-10), efeitos adversos. Nunca suspender abruptamente - risco de síndrome de descontinuação grave. Investigar: causas tratáveis de sintomas cognitivos/neuropsiquiátricos conforme protocolos anteriores (deficiências nutricionais, disfunção tireoidiana, inflamação, disbiose). Avaliar depletações induzidas por medicamentos: estatinas → avaliar CoQ10; anticonvulsivantes → folato, B12, vitamina D, carnitina; antipsicóticos → CoQ10, vitamina E; metformina → B12. Suplementação para potencializar efeito e reduzir dose necessária: ômega-3, magnésio, complexo B, probióticos. Otimização de estilo de vida: sono, exercício, dieta anti-inflamatória, gestão de estresse. Se objetivo de desmame: protocolo gradual (redução 10-25% a cada 2-4 semanas) sob supervisão psiquiátrica, com suporte intensivo psicoterapêutico e farmacológico complementar. Jamais estimular desmame em pacientes com transtornos graves, histórico de tentativas de suicídio ou psicose.',

  updated_at = NOW()
WHERE id = '36d39438-5337-4e06-b67d-edfaa498ba25';

-- Linkar artigos relevantes
INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('bc8d38a2-f684-40ce-8859-04f32570136a', '36d39438-5337-4e06-b67d-edfaa498ba25'), -- MFI Psiquiatria Aula 01
  ('edcd8bcc-4b08-4021-b842-caee99b2e632', '36d39438-5337-4e06-b67d-edfaa498ba25'), -- MFI Psiquiatria Aula 02
  ('0dd08507-9078-4c09-9a3a-972f19ccbd3c', '36d39438-5337-4e06-b67d-edfaa498ba25'), -- MFI Psiquiatria Aula 05
  ('05898fe5-54de-4c22-a56a-aca088c6706a', '36d39438-5337-4e06-b67d-edfaa498ba25')  -- MFI Psiquiatria Aula 06
ON CONFLICT DO NOTHING;

-- ============================================================================
-- FIM DO SCRIPT
-- Total de items enriquecidos: 6
-- Total de links artigo-item criados: 21
-- ============================================================================
