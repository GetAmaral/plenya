-- ========================================
-- BATCH 29: ENRIQUECIMENTO GRUPO SOCIAL
-- Total: 9 items
-- Data: 2026-01-27
-- ========================================

BEGIN;

-- ========================================
-- ITEM 1: Situação familiar
-- ID: b2edeb6b-4855-45a9-bc40-692bbb4c010c
-- ========================================

UPDATE score_items
SET
  clinical_relevance = 'A situação familiar representa um dos determinantes sociais de saúde mais significativos, influenciando diretamente desfechos clínicos através de múltiplos mecanismos biopsicossociais. O suporte social familiar está associado a menor risco cardiovascular, melhor controle glicêmico em diabéticos, maior adesão terapêutica e redução de 50% na mortalidade por todas as causas. Pacientes com suporte familiar adequado apresentam níveis reduzidos de cortisol, melhor função imunológica e menor incidência de transtornos mentais. A composição familiar, dinâmica relacional, presença de cuidadores e qualidade dos vínculos afetivos modulam a resposta ao estresse crônico e a expressão de genes relacionados à inflamação. Conflitos familiares crônicos elevam biomarcadores inflamatórios (PCR, IL-6) e aumentam risco de doenças autoimunes. A avaliação da rede de suporte deve incluir: número de pessoas no domicílio, qualidade das relações, presença de dependentes, sobrecarga de cuidador e isolamento social. Intervenções baseadas em terapia familiar sistêmica demonstram benefícios mensuráveis em parâmetros metabólicos e cardiovasculares.',
  patient_explanation = 'A sua situação familiar tem impacto direto na sua saúde física e mental. Pesquisas mostram que pessoas com bom suporte familiar apresentam sistema imunológico mais forte, menor risco de doenças do coração e recuperam-se mais rápido de enfermidades. O contrário também é verdadeiro: conflitos familiares constantes, isolamento ou sobrecarga de cuidados podem aumentar o estresse crônico e facilitar o desenvolvimento de doenças. Avaliar sua rede de apoio ajuda a identificar recursos disponíveis e necessidades de suporte, seja através de familiares, amigos ou grupos comunitários. Fortalecer vínculos positivos e desenvolver habilidades de comunicação familiar são estratégias terapêuticas que melhoram não apenas o bem-estar emocional, mas também indicadores clínicos objetivos como pressão arterial, glicemia e marcadores inflamatórios.',
  conduct = 'Realizar genograma familiar para mapear rede de suporte e conflitos. Avaliar qualidade das relações através de escalas validadas (Escala de Suporte Social). Identificar fatores de risco: isolamento social, conflitos crônicos, sobrecarga de cuidador. Encaminhar para terapia familiar quando indicado. Orientar sobre estratégias de comunicação não-violenta e resolução de conflitos. Conectar paciente a grupos de apoio e recursos comunitários. Reavaliar suporte social a cada consulta, especialmente em doenças crônicas. Considerar suporte familiar como fator prognóstico em planejamento terapêutico.',
  updated_at = NOW()
WHERE id = 'b2edeb6b-4855-45a9-bc40-692bbb4c010c';

-- Linkar artigos relevantes
INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('664a4773-d083-4295-a7ec-0c7156f5457e', 'b2edeb6b-4855-45a9-bc40-692bbb4c010c'),
  ('d5369fca-38bc-4cbb-9742-2090b2789892', 'b2edeb6b-4855-45a9-bc40-692bbb4c010c')
ON CONFLICT DO NOTHING;


-- ========================================
-- ITEM 2: Laser / hobbies
-- ID: c8ac45bc-f7fa-48e2-9cdc-b892a703c6d5
-- ========================================

UPDATE score_items
SET
  clinical_relevance = 'Atividades de lazer e hobbies representam componentes essenciais da saúde integral, com impacto mensurável em biomarcadores de estresse e longevidade. Engajamento regular em atividades prazerosas está associado a redução de 20-30% no risco de mortalidade por todas as causas, independente de exercício físico estruturado. Hobbies criativos, sociais ou cognitivamente estimulantes modulam o eixo HPA, reduzem cortisol sérico e aumentam BDNF (fator neurotrófico derivado do cérebro). Pacientes com hobbies regulares apresentam menor incidência de demência (35% redução), depressão e síndrome metabólica. A neurobiologia do prazer envolve ativação do sistema dopaminérgico mesolímbico, promovendo neuroplasticidade e resiliência ao estresse. Ausência de atividades de lazer está associada a burnout, insônia e maior risco cardiovascular. A avaliação deve incluir: frequência, variedade, presença de flow state, componente social e barreiras percebidas. Prescrição de hobbies como intervenção terapêutica demonstra eficácia comparável a intervenções farmacológicas leves em quadros ansiosos e depressivos leves.',
  patient_explanation = 'Ter hobbies e atividades de lazer não é luxo, mas necessidade biológica para saúde plena. Quando você se envolve em atividades prazerosas, seu cérebro libera substâncias que protegem contra estresse, fortalecem memória e melhoram o humor. Estudos mostram que pessoas com hobbies regulares vivem mais, têm menor risco de demência e apresentam melhor qualidade de sono. Não precisa ser algo complexo ou caro: jardinagem, leitura, artesanato, música, jogos de tabuleiro ou qualquer atividade que traga satisfação genuína já produz benefícios mensuráveis. O ideal é dedicar pelo menos 2-3 horas semanais a atividades que não sejam trabalho ou obrigações domésticas. Se você sente que não tem tempo para lazer, isso pode ser um sinal de desequilíbrio que precisa ser abordado terapeuticamente.',
  conduct = 'Investigar histórico de hobbies e identificar barreiras atuais (tempo, recursos, fadiga). Prescrever "receita de lazer": 2-3 sessões semanais de atividades prazerosas. Orientar sobre flow state e engajamento significativo. Sugerir variedade: atividades criativas, sociais, cognitivas e ao ar livre. Encaminhar para grupos comunitários ou oficinas terapêuticas. Monitorar adesão e ajustar conforme necessidades. Educar sobre neurobiologia do prazer e importância da desconexão produtiva. Reavaliar impacto em sintomas de estresse e qualidade de vida.',
  updated_at = NOW()
WHERE id = 'c8ac45bc-f7fa-48e2-9cdc-b892a703c6d5';

-- Linkar artigos relevantes
INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('e73ad009-bf8d-4b2b-a37e-d29978f4533f', 'c8ac45bc-f7fa-48e2-9cdc-b892a703c6d5'),
  ('721299e2-7131-417e-9da5-e7e509c088eb', 'c8ac45bc-f7fa-48e2-9cdc-b892a703c6d5')
ON CONFLICT DO NOTHING;


-- ========================================
-- ITEM 3: Condições de moradia
-- ID: da38d7a5-f201-4801-bb58-5cf2e4209288
-- ========================================

UPDATE score_items
SET
  clinical_relevance = 'Condições de moradia constituem determinante social de saúde com impacto direto em múltiplos desfechos clínicos. Habitações inadequadas (umidade excessiva, ventilação deficiente, superlotação) elevam risco de doenças respiratórias, alergias, asma e infecções em 40-60%. Presença de mofo está associada a asma grave, rinosinusite crônica e exacerbação de doenças autoimunes. Densidade habitacional excessiva (>2 pessoas/cômodo) aumenta transmissão de patógenos respiratórios e gastrointestinais, além de elevar cortisol e prejudicar qualidade do sono. Falta de saneamento básico está diretamente relacionada a parasitoses intestinais, que prejudicam absorção de nutrientes e modulam negativamente o microbioma. Temperatura inadequada da moradia afeta termorregulação, aumentando risco cardiovascular em idosos. Segurança estrutural e risco de acidentes domésticos impactam morbimortalidade. A avaliação deve incluir: tipo de construção, número de cômodos, ventilação, luminosidade natural, presença de umidade/mofo, saneamento básico, segurança e percepção de conforto. Intervenções ambientais demonstram redução significativa em hospitalizações por asma e doenças infecciosas.',
  patient_explanation = 'Sua casa não é apenas um endereço, mas um fator que influencia diretamente sua saúde. Moradias com umidade ou mofo podem piorar alergias, asma e infecções respiratórias. Casas muito cheias dificultam o descanso adequado e facilitam a transmissão de doenças. Falta de ventilação e luz natural prejudica o sono e o humor. Ausência de saneamento básico expõe a parasitas intestinais que roubam nutrientes do corpo. Temperatura inadequada pode agravar problemas cardíacos, especialmente em idosos. Avaliar suas condições de moradia ajuda a identificar riscos ambientais que podem estar contribuindo para sintomas crônicos. Mesmo pequenas melhorias - como abrir janelas regularmente, eliminar mofo ou reorganizar espaços - podem trazer benefícios significativos para sua saúde.',
  conduct = 'Realizar questionário detalhado sobre condições de moradia. Investigar presença de mofo, umidade, ventilação inadequada e superlotação. Avaliar saneamento básico e acesso a água potável. Orientar sobre ventilação cruzada diária (15-30min). Encaminhar para assistência social quando necessário. Solicitar inspeção sanitária em casos graves. Educar sobre relação entre ambiente doméstico e sintomas respiratórios/alérgicos. Considerar mudanças ambientais como parte do plano terapêutico. Documentar condições de moradia como fator prognóstico.',
  updated_at = NOW()
WHERE id = 'da38d7a5-f201-4801-bb58-5cf2e4209288';

-- Linkar artigos relevantes
INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('70298603-a991-437d-80e6-d15af039b80d', 'da38d7a5-f201-4801-bb58-5cf2e4209288'),
  ('a472d98c-ebd8-4b59-bc62-0beddce19126', 'da38d7a5-f201-4801-bb58-5cf2e4209288')
ON CONFLICT DO NOTHING;


-- ========================================
-- ITEM 4: Qualidade do Ar Interior
-- ID: 7e4167cb-509e-4901-b3ce-42d6d870fe8c
-- ========================================

UPDATE score_items
SET
  clinical_relevance = 'A qualidade do ar interior representa fator crítico em medicina ambiental, com concentrações de poluentes frequentemente 2-5 vezes superiores ao ar externo. Exposição crônica a material particulado (PM2.5, PM10), compostos orgânicos voláteis (VOCs) e CO2 elevado está associada a disfunção mitocondrial, estresse oxidativo sistêmico e inflamação de baixo grau. Poluentes internos (produtos de limpeza, tintas, móveis novos, fumaça de tabaco, fogões a gás) liberam formaldeído, benzeno e tolueno, que são desreguladores endócrinos e carcinogênicos. CO2 acima de 1000ppm prejudica função cognitiva em até 50%. Exposição a fungos e ácaros desencadeia resposta Th2, perpetuando quadros alérgicos e inflamatórios crônicos. Má qualidade do ar interior está associada a asma, DPOC, alergias, fadiga crônica, disfunção cognitiva e risco cardiovascular aumentado. A avaliação deve incluir: ventilação, uso de produtos químicos, presença de fumantes, tipo de fogão, plantas purificadoras, filtros de ar. Intervenções simples (ventilação adequada, redução de VOCs) melhoram marcadores inflamatórios e função pulmonar.',
  patient_explanation = 'O ar que você respira dentro de casa pode estar mais poluído que o ar da rua, afetando pulmões, cérebro e sistema imunológico. Produtos de limpeza, tintas, móveis novos, fumaça de cigarro e até fogões a gás liberam substâncias tóxicas no ar. Ambientes fechados acumulam gás carbônico, reduzindo sua capacidade de concentração e causando fadiga. Mofo e poeira acumulada pioram alergias e asma. Esses poluentes invisíveis causam inflamação crônica no corpo, aumentam o risco de doenças respiratórias e cardiovasculares e prejudicam o sono. Melhorar a qualidade do ar interno é intervenção terapêutica eficaz: abrir janelas diariamente, usar produtos de limpeza naturais, ter plantas purificadoras e eliminar fontes de mofo fazem diferença real na sua saúde.',
  conduct = 'Investigar fontes de poluição interna: produtos de limpeza, fumaça, fogão a gás, tintas recentes. Orientar ventilação cruzada diária (mínimo 15min). Recomendar substituição de produtos químicos por alternativas naturais (vinagre, bicarbonato). Sugerir plantas purificadoras (espada-de-são-jorge, jiboia, lírio-da-paz). Avaliar necessidade de filtro HEPA em casos de alergias graves. Educar sobre VOCs e seus efeitos à saúde. Orientar manutenção de sistemas de ventilação e ar-condicionado. Considerar qualidade do ar como fator terapêutico em doenças respiratórias.',
  updated_at = NOW()
WHERE id = '7e4167cb-509e-4901-b3ce-42d6d870fe8c';

-- Linkar artigos relevantes
INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('d664415f-fccf-4c20-b183-ab0af64490be', '7e4167cb-509e-4901-b3ce-42d6d870fe8c'),
  ('ee30962a-8025-4c25-8811-2d5b8f862d9d', '7e4167cb-509e-4901-b3ce-42d6d870fe8c')
ON CONFLICT DO NOTHING;


-- ========================================
-- ITEM 5: Exposição Ambiental Externa
-- ID: cd6b1f77-00a3-4817-9a38-3e34313af80b
-- ========================================

UPDATE score_items
SET
  clinical_relevance = 'Exposição ambiental externa engloba poluentes atmosféricos, pesticidas, metais pesados, radiação eletromagnética e outros xenobióticos que modulam epigeneticamente a expressão gênica e aceleram envelhecimento celular. Poluição do ar por PM2.5 está associada a aumento de 10-15% no risco cardiovascular e 8% na mortalidade por todas as causas para cada 10μg/m³ de incremento. Exposição ocupacional ou residencial a pesticidas organofosforados prejudica função mitocondrial, depleta glutationa e está relacionada a Parkinson, declínio cognitivo e disruptores endócrinos. Metais pesados (chumbo, mercúrio, arsênio) bioacumulam em tecido adiposo e ósseo, causando neurotoxicidade, nefrotoxicidade e risco carcinogênico. Proximidade de rodovias, indústrias ou áreas agrícolas aumenta carga alostática e marcadores inflamatórios. Radiação não-ionizante (torres, dispositivos) permanece controversa, mas princípio da precaução é recomendado. Avaliação deve incluir: localização residencial/ocupacional, proximidade de fontes poluidoras, uso de pesticidas, exposição ocupacional. Testes de metais pesados e screening de desreguladores endócrinos podem ser indicados.',
  patient_explanation = 'Onde você mora e trabalha expõe seu corpo a substâncias ambientais que podem prejudicar sua saúde ao longo do tempo. Poluição do ar aumenta risco de doenças do coração e pulmões. Morar perto de rodovias movimentadas, indústrias ou áreas agrícolas com uso intenso de agrotóxicos pode expô-lo a toxinas que se acumulam no corpo, causando inflamação crônica. Metais pesados como chumbo e mercúrio prejudicam o cérebro, rins e sistema nervoso. Exposição ocupacional a produtos químicos também é fator de risco. Identificar suas exposições ambientais permite tomar medidas protetoras: uso de filtros de ar, consumo de alimentos orgânicos, suplementação com nutrientes que ajudam na desintoxicação (glutationa, vitamina C, selênio) e, quando possível, mudanças de ambiente.',
  conduct = 'Mapear exposições ambientais: localização residencial/ocupacional, proximidade de fontes poluidoras. Investigar uso doméstico de pesticidas e produtos químicos. Avaliar exposição ocupacional a toxinas. Solicitar dosagem de metais pesados se indicado (chumbo, mercúrio, arsênio). Orientar sobre medidas protetoras: filtros de ar, alimentos orgânicos, EPIs ocupacionais. Prescrever suporte nutricional para desintoxicação: NAC, glutationa, vitamina C, selênio, magnésio. Considerar quelação em casos de intoxicação confirmada. Documentar exposições como fator de risco em prontuário.',
  updated_at = NOW()
WHERE id = 'cd6b1f77-00a3-4817-9a38-3e34313af80b';

-- Linkar artigos relevantes
INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('3bc699f0-37df-442a-9166-60af73e1ac60', 'cd6b1f77-00a3-4817-9a38-3e34313af80b'),
  ('3cd27396-2dc3-4bea-9aa9-192db82cb88d', 'cd6b1f77-00a3-4817-9a38-3e34313af80b')
ON CONFLICT DO NOTHING;


-- ========================================
-- ITEM 6: Ambiente Sonoro
-- ID: 91e450db-29df-4a78-8741-441f89630ff7
-- ========================================

UPDATE score_items
SET
  clinical_relevance = 'Poluição sonora representa estressor ambiental crônico com efeitos mensuráveis em múltiplos sistemas fisiológicos. Exposição a ruído acima de 55dB durante o sono ou 65dB durante o dia está associada a hipertensão arterial, infarto agudo do miocárdio, AVC e diabetes tipo 2. Cada incremento de 10dB no ruído ambiental aumenta risco cardiovascular em 8-10%. Mecanismos incluem ativação crônica do sistema nervoso simpático, elevação persistente de cortisol, disfunção endotelial e inflamação sistêmica (elevação de IL-6 e PCR). Ruído noturno fragmenta arquitetura do sono, reduzindo sono REM e profundo, prejudicando consolidação de memória e regulação metabólica. Exposição ocupacional a ruído intenso (>85dB) causa perda auditiva irreversível e está associada a maior prevalência de ansiedade e depressão. Crianças expostas a ruído crônico apresentam atraso no desenvolvimento cognitivo e linguístico. Avaliação deve incluir: fontes de ruído (tráfego, vizinhança, ocupacional), decibéis estimados, uso de proteção auditiva e qualidade do sono. Intervenções de redução de ruído melhoram pressão arterial e qualidade de vida.',
  patient_explanation = 'Ruído constante não é apenas incômodo, mas fator de risco real para doenças do coração, pressão alta, diabetes e problemas de sono. Quando você está exposto a barulho contínuo - tráfego, vizinhos, construções, ambiente de trabalho - seu corpo permanece em estado de alerta, liberando hormônios do estresse que elevam pressão arterial e inflamação. Ruído durante a noite prejudica a qualidade do sono, mesmo que você não acorde completamente, impedindo que seu corpo se recupere adequadamente. Com o tempo, isso aumenta o risco de infarto, AVC e transtornos mentais. Avaliar seu ambiente sonoro permite tomar medidas protetoras: uso de protetores auriculares, isolamento acústico, máquinas de ruído branco ou mudanças de ambiente quando possível. Ambientes silenciosos são intervenção terapêutica real.',
  conduct = 'Investigar exposições a ruído: residencial (tráfego, vizinhança), ocupacional, recreacional. Estimar níveis de decibéis e duração de exposição. Avaliar impacto no sono através de polissonografia ou diário de sono. Orientar sobre limites seguros: <55dB noturno, <65dB diurno. Recomendar protetores auriculares para exposições intensas. Sugerir isolamento acústico residencial quando viável. Prescrever técnicas de manejo de estresse se exposição inevitável. Monitorar pressão arterial e marcadores cardiovasculares em expostos crônicos. Encaminhar para audiometria se exposição ocupacional >85dB. Considerar ruído como fator de risco cardiovascular.',
  updated_at = NOW()
WHERE id = '91e450db-29df-4a78-8741-441f89630ff7';

-- Linkar artigos relevantes
INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('c7f2d451-b59a-421b-8e52-abfe7440a98a', '91e450db-29df-4a78-8741-441f89630ff7'),
  ('94565e8c-a859-42e2-9f4e-c1977b3201ae', '91e450db-29df-4a78-8741-441f89630ff7')
ON CONFLICT DO NOTHING;


-- ========================================
-- ITEM 7: Qualidade da Água
-- ID: bdd9ed17-65fc-4357-8bc3-70a48b490520
-- ========================================

UPDATE score_items
SET
  clinical_relevance = 'A qualidade da água de consumo impacta diretamente múltiplos sistemas fisiológicos através de contaminantes químicos, metais pesados, microbiológicos e desreguladores endócrinos. Água contaminada por metais pesados (chumbo, arsênio, mercúrio) causa neurotoxicidade, nefrotoxicidade e risco carcinogênico cumulativo. Fluoreto em excesso (>1.5mg/L) está associado a disfunção tireoidiana e prejuízo cognitivo em crianças. Resíduos de pesticidas e fármacos (anticoncepcionais, antibióticos) funcionam como disruptores endócrinos, alterando eixo hipotálamo-hipófise-gonadal. Cloro e subprodutos da desinfecção (trihalometanos) estão relacionados a risco aumentado de câncer de bexiga. Microplásticos presentes em água engarrafada carregam toxinas lipossolúveis e xenoestrógenos. Dureza excessiva da água (alto cálcio/magnésio) pode causar litíase renal, enquanto água muito mole aumenta absorção de metais das tubulações. Contaminação microbiológica causa diarreias, parasitoses e disbiose intestinal. Avaliação deve incluir: fonte de água, tratamento, análise físico-química periódica, tipo de armazenamento. Filtração adequada e remineralização quando necessário são intervenções fundamentais.',
  patient_explanation = 'A água que você bebe todos os dias pode conter substâncias invisíveis que prejudicam sua saúde a longo prazo. Mesmo água tratada pode ter resíduos de cloro, metais pesados das tubulações, hormônios de anticoncepcionais, restos de agrotóxicos e microplásticos. Essas substâncias se acumulam no corpo, causando problemas na tireoide, hormônios sexuais, rins e aumentando risco de câncer. Água contaminada também pode alterar seu microbioma intestinal. Por outro lado, água muito "pura" (destilada ou de osmose reversa sem remineralização) pode depletar minerais essenciais. O ideal é consumir água filtrada adequadamente, preferencialmente em filtros de carvão ativado ou osmose reversa com remineralização. Evite água de plástico sempre que possível. Avaliar a qualidade da sua água é parte importante do cuidado preventivo.',
  conduct = 'Investigar fonte de água: rede pública, poço, mineral engarrafada. Avaliar sistema de filtração doméstica. Recomendar análise físico-química e microbiológica da água. Orientar sobre tipos de filtros: carvão ativado (remove cloro, VOCs), osmose reversa (remove metais, fluoreto). Alertar sobre necessidade de remineralização pós-osmose reversa. Orientar consumo mínimo: 30-35ml/kg/dia. Desencorajar água em garrafas plásticas (BPA, ftalatos). Avaliar hidratação através de osmolalidade urinária. Suplementar minerais se água muito mole ou pós-osmose. Considerar qualidade da água em doenças renais, tireoideopatias e disruptores endócrinos.',
  updated_at = NOW()
WHERE id = 'bdd9ed17-65fc-4357-8bc3-70a48b490520';

-- Linkar artigos relevantes
INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('ec953cdf-f2cd-42b1-a5da-0bd1675a40c6', 'bdd9ed17-65fc-4357-8bc3-70a48b490520'),
  ('b8a9fb8d-5149-467b-bfc5-371e422b84eb', 'bdd9ed17-65fc-4357-8bc3-70a48b490520')
ON CONFLICT DO NOTHING;


-- ========================================
-- ITEM 8: Espaço para Movimento
-- ID: b7736f9d-5498-4ae7-8a07-f92d9edccc4d
-- ========================================

UPDATE score_items
SET
  clinical_relevance = 'Disponibilidade de espaço para movimento no ambiente doméstico e comunitário representa determinante estrutural crítico para atividade física habitual e saúde metabólica. Indivíduos residindo em ambientes com espaços adequados para movimento apresentam 30-40% mais atividade física não estruturada (NEAT - Non-Exercise Activity Thermogenesis), contribuindo significativamente para gasto energético diário. Moradias com espaço limitado estão associadas a comportamento sedentário aumentado, elevando risco de obesidade, resistência insulínica e síndrome metabólica. Crianças sem espaço para brincar apresentam maior prevalência de obesidade infantil e atraso no desenvolvimento motor. Acesso a áreas verdes, parques e ciclovias está inversamente relacionado a mortalidade cardiovascular e depressão. Ambientes construídos walkable aumentam mobilidade ativa, reduzindo risco cardiometabólico em 20%. Espaço interno para prática de exercícios domésticos melhora adesão a programas de atividade física. Avaliação deve incluir: metragem da moradia, presença de áreas externas, proximidade de parques/academias, segurança para circulação. Urbanismo em saúde considera espaço para movimento como infraestrutura de saúde pública.',
  patient_explanation = 'Ter espaço adequado para se movimentar em casa e no bairro influencia diretamente sua saúde metabólica e risco de doenças. Quando você tem espaço para caminhar, alongar, brincar ou praticar exercícios, é mais provável que se mova ao longo do dia, mesmo sem treino formal. Esse movimento natural queima calorias, melhora circulação, regula açúcar no sangue e protege o coração. Morar em casas muito pequenas ou bairros sem áreas verdes aumenta comportamento sedentário e risco de obesidade, diabetes e depressão. Crianças precisam especialmente de espaço para desenvolver coordenação motora adequada. Avaliar seu ambiente permite identificar oportunidades ou barreiras para movimento. Mesmo em espaços pequenos, é possível criar rotinas de exercício doméstico. Proximidade de parques e ruas seguras para caminhar são fatores protetores importantes.',
  conduct = 'Avaliar metragem e configuração da moradia quanto a espaço para movimento. Investigar presença de áreas externas (quintal, varanda). Mapear proximidade de parques, praças e ciclovias. Avaliar segurança do bairro para caminhadas. Identificar barreiras arquitetônicas ao movimento. Orientar sobre exercícios domésticos em espaços reduzidos. Prescrever caminhadas em parques próximos quando disponível. Encaminhar para programas comunitários de atividade física. Educar sobre NEAT e movimento não estruturado. Considerar mudança de ambiente como intervenção terapêutica em casos graves de sedentarismo.',
  updated_at = NOW()
WHERE id = 'b7736f9d-5498-4ae7-8a07-f92d9edccc4d';

-- Linkar artigos relevantes
INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('1e824dec-145e-49f1-962a-5ada38614bce', 'b7736f9d-5498-4ae7-8a07-f92d9edccc4d'),
  ('1081414c-5d15-41bc-8f08-b2c618927913', 'b7736f9d-5498-4ae7-8a07-f92d9edccc4d')
ON CONFLICT DO NOTHING;


-- ========================================
-- ITEM 9: Hobbies (Atual)
-- ID: f335cb3d-df6b-4f37-822d-377fdaaf2fca
-- ========================================

UPDATE score_items
SET
  clinical_relevance = 'Manutenção de hobbies e atividades de interesse pessoal constitui fator protetor documentado contra múltiplas condições crônicas, atuando através de modulação neuroendócrina e redução de carga alostática. Meta-análises demonstram que engajamento regular em hobbies reduz risco de demência em 35%, depressão em 30% e mortalidade por todas as causas em 20-30%, independentemente de fatores confundidores. Mecanismos incluem: ativação do sistema dopaminérgico mesolímbico, aumento de BDNF (neuroplasticidade), redução de cortisol e citocinas pró-inflamatórias (IL-6, TNF-α). Hobbies cognitivamente estimulantes (leitura, xadrez, instrumentos musicais) aumentam reserva cognitiva e conectividade neural. Atividades criativas modulam resposta parassimpática, reduzindo frequência cardíaca e pressão arterial. Hobbies sociais ampliam rede de suporte e combatem isolamento, fator de risco independente para mortalidade. Flow state (absorção completa na atividade) está associado a bem-estar sustentado e resiliência ao estresse. Ausência de hobbies correlaciona-se com burnout, insônia e síndrome metabólica. Prescrição de hobbies como intervenção terapêutica demonstra custo-efetividade superior a muitas intervenções farmacológicas.',
  patient_explanation = 'Manter hobbies regulares é uma das intervenções mais poderosas para proteger sua saúde física e mental. Quando você se dedica a atividades prazerosas, seu cérebro produz substâncias que combatem o estresse, melhoram a memória e protegem contra doenças. Pesquisas mostram que pessoas com hobbies vivem mais tempo, têm menor risco de demência, depressão e doenças do coração. Não importa qual hobby - jardinagem, leitura, pintura, música, jogos, artesanato - o importante é dedicar tempo regular a algo que traga satisfação genuína. O ideal é pelo menos 2-3 sessões semanais de atividades não relacionadas a trabalho ou obrigações. Se você sente que não tem tempo para hobbies, isso é um sinal de alerta de desequilíbrio que precisa ser tratado. Hobbies não são luxo, são necessidade biológica para saúde plena.',
  conduct = 'Investigar hobbies atuais e histórico de interesses. Identificar barreiras: tempo, recursos, fadiga, anhedonia. Prescrever formalmente hobbies: 2-3 sessões semanais, 30-60min cada. Orientar sobre diversidade: atividades cognitivas, criativas, sociais e físicas. Sugerir experimentação de novas atividades. Encaminhar para grupos comunitários, oficinas ou aulas. Educar sobre flow state e engajamento significativo. Monitorar adesão como parte do tratamento. Avaliar impacto em sintomas depressivos, ansiosos e qualidade de sono. Reavaliar periodicamente e ajustar conforme evolução.',
  updated_at = NOW()
WHERE id = 'f335cb3d-df6b-4f37-822d-377fdaaf2fca';

-- Linkar artigos relevantes
INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('3b83c045-0caf-4e09-bc55-22d0cb22c4bd', 'f335cb3d-df6b-4f37-822d-377fdaaf2fca'),
  ('92ad5318-8cc0-49b5-9afd-0e47743a28d0', 'f335cb3d-df6b-4f37-822d-377fdaaf2fca')
ON CONFLICT DO NOTHING;

COMMIT;

-- ========================================
-- FIM DO SCRIPT
-- Total de items enriquecidos: 9
-- Total de artigos linkados: 18 (2 por item)
-- ========================================
