-- ============================================================
-- BATCH 3 PARTE 2: Items restantes (razão cintura/altura até água extracelular)
-- ============================================================

-- Item: Razão cintura/altura (cm/cm)
UPDATE score_items SET
  clinical_relevance = 'A razão cintura/altura (RCA ou WHtR) é marcador simples e poderoso de adiposidade abdominal e risco cardiometabólico, considerado superior ao IMC e comparável à RCQ. Regra simples: cintura deve ser <50% da altura ("Keep your waist to less than half your height"). Valores: <0.4 muito magro, 0.4-0.49 saudável, 0.5-0.59 sobrepeso/risco aumentado, ≥0.6 obesidade abdominal/risco muito elevado. RCA ≥0.5 associa-se a risco aumentado de diabetes tipo 2, doenças cardiovasculares, mortalidade, independente de sexo, idade e etnia. Vantagens sobre IMC: não requer tabelas por sexo/idade, considera distribuição de gordura, aplicável em crianças e adultos. Meta universal: RCA <0.5.',
  patient_explanation = 'A relação entre sua cintura e sua altura é um dos indicadores mais simples e eficazes de risco à saúde. A regra é fácil: sua cintura deve ser menor que METADE da sua altura. Se você tem 170cm, sua cintura deveria ser menor que 85cm. Valores acima de 0.5 (cintura igual ou maior que metade da altura) indicam acúmulo de gordura abdominal e risco aumentado de diabetes e problemas cardiovasculares. Essa medida funciona para qualquer pessoa, independente de ser homem ou mulher, jovem ou idoso, alto ou baixo. É mais útil que o IMC porque considera onde a gordura está (barriga vs coxas).',
  conduct = 'Calcular RCA = circunferência cintura (cm) / altura (cm). Interpretar: <0.4 muito magro, 0.4-0.49 saudável, 0.5-0.59 risco aumentado, ≥0.6 risco muito elevado. RCA ≥0.5 requer investigação metabólica completa e intervenção. Meta universal: <0.5. Útil para explicar risco ao paciente (conceito simples: cintura <metade da altura).',
  updated_at = NOW()
WHERE id = '936855d1-e715-4171-a061-3bcc500957f7';

-- Items de circunferência de pescoço (homem e mulher) - marcador de apneia do sono e gordura visceral
UPDATE score_items SET
  clinical_relevance = 'Circunferência de pescoço correlaciona-se com gordura visceral, síndrome metabólica e apneia obstrutiva do sono (AOS). Valores de risco: >40cm homens, >36cm mulheres. Pescoço aumentado em conjunto com cintura aumentada e IMC elevado é forte preditor de AOS (Escore de Mallampati, Escala de Epworth). AOS associa-se a hipertensão refratária, arritmias, resistência insulínica, inflamação crônica. Investigação com polissonografia se suspeita.',
  patient_explanation = 'A medida do pescoço pode parecer estranha, mas é um indicador importante de gordura visceral e risco de apneia do sono (paradas respiratórias durante o sono). Pescoço grosso (>40cm homens, >36cm mulheres) junto com ronco, sonolência diurna e cansaço sugerem apneia do sono, que prejudica a qualidade do sono, aumenta pressão arterial e piora o metabolismo.',
  conduct = 'Medir circunferência de pescoço abaixo da proeminência laríngea (pomo de Adão), paciente sentado ou em pé, cabeça ereta olhando para frente, fita ajustada sem comprimir. Risco AOS se >40cm (H) ou >36cm (M). Se suspeita de AOS (pescoço aumentado + ronco + sonolência diurna + obesidade), aplicar Escala de Epworth e encaminhar para polissonografia.',
  updated_at = NOW()
WHERE id IN ('c3ee7527-104b-467a-a87e-b4587192005b', '421db9cd-d75d-4f0a-8188-216affee192f');

-- Items de relação pescoço-altura (homem e mulher)
UPDATE score_items SET
  clinical_relevance = 'Relação pescoço/altura (RPA) ajustada combina informações de adiposidade cervical com estatura. Valores de risco sugeridos: >0.24 homens, >0.215 mulheres. Correlaciona-se com síndrome metabólica e risco cardiovascular.',
  patient_explanation = 'A relação entre o pescoço e a altura refina a avaliação de risco, considerando que pessoas mais altas naturalmente têm pescoços maiores. Ajuda a identificar acúmulo desproporcional de gordura na região cervical.',
  conduct = 'Calcular RPA = circunferência pescoço (cm) / altura (cm). Valores >0.24 (H) ou >0.215 (M) sugerem risco aumentado. Complementar avaliação de síndrome metabólica.',
  updated_at = NOW()
WHERE id IN ('ff5472cf-e148-48c1-be7a-245df89c200c', '472969f8-7ea9-4532-8c73-c4a96f31e2fc');

-- Items de panturrilha (homem e mulher) - marcador de sarcopenia
UPDATE score_items SET
  clinical_relevance = 'Circunferência de panturrilha é marcador de massa muscular de membros inferiores e rastreio de sarcopenia, especialmente em idosos. Ponto de corte: <31cm indica risco de sarcopenia independente de sexo. Panturrilha reduzida associa-se a fragilidade, quedas, fraturas, perda funcional, mortalidade. Avaliar junto com força de preensão palmar e testes funcionais (velocidade de marcha, levantar da cadeira).',
  patient_explanation = 'A medida da panturrilha (batata da perna) avalia sua massa muscular nos membros inferiores. Panturrilhas finas (<31cm) indicam perda muscular (sarcopenia), especialmente preocupante em pessoas mais velhas. Músculo das pernas é essencial para caminhar, equilíbrio, prevenir quedas e manter independência.',
  conduct = 'Medir circunferência máxima da panturrilha com paciente sentado, pé apoiado, músculo relaxado. Valores <31cm sugerem sarcopenia. Complementar com força de preensão palmar, velocidade de marcha, massa muscular por bioimpedância/DEXA. Se sarcopenia, investigar causas (desnutrição proteica, inatividade, doenças consumptivas, déficits hormonais) e intervir (exercício resistido, suplementação proteica, correção hormonal).',
  updated_at = NOW()
WHERE id IN ('094b4b55-e7a4-4587-be08-9963c0520bf0', '660ec23e-8957-4936-b83b-284a0f8c84eb');

-- Items de braço relaxado (homem e mulher)
UPDATE score_items SET
  clinical_relevance = 'Circunferência de braço reflete massa muscular de membros superiores e reserva proteica. Útil para avaliar sarcopenia e desnutrição. Valores de referência variam por sexo, idade e etnia. Redução progressiva sugere catabolismo muscular (desnutrição, doenças consumptivas, envelhecimento). Aumento pode refletir ganho muscular (treinamento) ou adiposo.',
  patient_explanation = 'A medida do braço avalia a massa muscular dos membros superiores. Braços muito finos sugerem perda muscular ou desnutrição. Monitoramos essa medida para garantir que você está ganhando ou mantendo massa muscular, especialmente se está em programa de treinamento de força.',
  conduct = 'Medir braço não-dominante relaxado no ponto médio entre acrômio (ombro) e olécrano (cotovelo), com braço pendente relaxado. Comparar com valores de referência por sexo e idade. Combinar com outras medidas de composição corporal e função muscular.',
  updated_at = NOW()
WHERE id IN ('30fa255c-83d9-4cd4-8b06-75f70e2fb3eb', '371f12db-5a68-4092-a4c3-1430ec21f18a');

-- Items de coxa (homem e mulher)
UPDATE score_items SET
  clinical_relevance = 'Circunferência de coxa reflete massa muscular do maior grupo muscular do corpo (quadríceps, isquiotibiais). Importante para força, funcionalidade, metabolismo. Redução de coxa associa-se a sarcopenia, fragilidade, risco de quedas. Estudos mostram correlação inversa entre circunferência de coxa e mortalidade (coxas maiores são protetoras, refletindo massa muscular).',
  patient_explanation = 'A medida da coxa avalia a musculatura das pernas, que é o maior grupo muscular do corpo. Coxas musculosas indicam boa massa muscular geral e bom metabolismo. Perda de massa nas coxas sugere sarcopenia e perda de força, aumentando risco de quedas e limitações físicas.',
  conduct = 'Medir circunferência de coxa no ponto médio entre prega inguinal e borda superior da patela, paciente em pé, peso distribuído, músculo relaxado. Monitorar mudanças: perda sugere catabolismo, ganho (em contexto de treinamento) indica hipertrofia muscular.',
  updated_at = NOW()
WHERE id IN ('8062469b-618a-4cfc-9239-046fd1f882e2', '25822830-7d95-4e20-9b32-bb5cd7d4a3d1');

-- Items de massa muscular esquelética (MME)
UPDATE score_items SET
  clinical_relevance = 'Massa muscular esquelética (MME) é componente crítico da composição corporal. Representa ~40-50% do peso corporal em adultos saudáveis. Funções: locomoção, postura, termogênese, reserva proteica, órgão endócrino (miocinas), metabolismo glicose. Sarcopenia: MME reduzida (<7kg/m² homens, <5.5kg/m² mulheres por DEXA). Causas: envelhecimento, inatividade, desnutrição proteica, déficits hormonais (testosterona, GH, tireoide), doenças crônicas, inflamação. Consequências: fragilidade, quedas, fraturas, resistência insulínica (músculo é principal sítio de captação de glicose), mortalidade. Objetivo: preservar/aumentar MME mesmo durante perda ponderal (através de exercício resistido e aporte proteico adequado 1.6-2.2g/kg/dia).',
  patient_explanation = 'Sua massa muscular esquelética é a quantidade total de músculo no seu corpo (excluindo órgãos). Músculo é essencial não apenas para força e movimento, mas também para metabolismo: ele queima calorias (quanto mais músculo, maior seu metabolismo basal), armazena glicose (prevenindo diabetes), produz substâncias anti-inflamatórias, e serve como reserva proteica para recuperação de doenças. Perda muscular (sarcopenia) é um dos principais fatores de envelhecimento não saudável. Nosso objetivo é preservar ou aumentar seu músculo através de treino de força e alimentação proteica adequada.',
  conduct = 'Avaliar MME por bioimpedância multifrequencial ou DEXA. Valores absolutos (kg) e relativos (% peso, kg/m²). Sarcopenia se índice MME <7kg/m² (H) ou <5.5kg/m² (M). Complementar com testes funcionais. Intervenção: exercício resistido progressivo 3-4x/semana, proteína 1.6-2.2g/kg/dia (distribuída nas refeições, 25-40g por refeição), correção de déficits hormonais, controle de inflamação. Reavaliação trimestral.',
  updated_at = NOW()
WHERE id IN ('2ff747e7-e48d-4207-85ee-168b1c14e35e', '725d0bc7-8f25-4c44-b268-53b72f75adff', 'cd592cb8-4c34-4964-ac29-63bf9c224bef');

-- Massa apendicular
UPDATE score_items SET
  clinical_relevance = 'Massa apendicular (soma de massa magra de membros superiores e inferiores) é o principal marcador de sarcopenia. Índice de massa apendicular (IMA = massa apendicular/altura²) é critério diagnóstico: sarcopenia se <7.0kg/m² homens, <5.5kg/m² mulheres (EWGSOP2). Perda progressiva de massa apendicular é marcador robusto de fragilidade e mortalidade.',
  patient_explanation = 'Massa apendicular é a soma dos músculos dos seus braços e pernas. É o melhor marcador de sarcopenia (perda muscular), pois esses músculos são essenciais para mobilidade, independência e qualidade de vida. Valores baixos indicam maior risco de quedas, fraturas e perda de autonomia.',
  conduct = 'Avaliar por DEXA (padrão-ouro) ou bioimpedância. Calcular IMA = massa apendicular(kg)/altura²(m). Sarcopenia se <7.0 (H) ou <5.5 (M). Intervir com treino resistido e nutrição proteica.',
  updated_at = NOW()
WHERE id = 'bdf03bb7-ef91-4e12-9247-4e29ce1e7185';

-- Items de massa gorda e FMI
UPDATE score_items SET
  clinical_relevance = 'Massa gorda total (MG) e Fat Mass Index (FMI = MG/altura²) quantificam adiposidade. Percentual de gordura (MG/peso): ótimo <15% H / <23% M; normal 15-20% H / 23-30% M; sobrepeso 20-25% H / 30-35% M; obesidade >25% H / >35% M. FMI >6.0kg/m² (H) ou >9.0kg/m² (M) indica excesso de gordura. Distinguir obesidade subcutânea vs visceral (esta última mais deletéria). Massa gorda não é inerte: tecido endócrino secretando adipocinas (pró-inflamatórias se excessiva). Meta: reduzir MG preservando MME.',
  patient_explanation = 'Sua massa gorda total e o índice de massa gorda (FMI) mostram QUANTO de gordura você tem no corpo. Não basta saber o peso - duas pessoas de mesmo peso podem ter quantidades muito diferentes de gordura. O objetivo é ter massa gorda em faixa saudável (não muito baixa, que prejudica hormônios, nem muito alta, que causa inflamação e doenças). Durante tratamento, queremos REDUZIR gordura enquanto PRESERVAMOS ou GANHAMOS músculo.',
  conduct = 'Avaliar por bioimpedância ou DEXA. Registrar MG absoluta (kg), percentual (%), FMI (kg/m²). Interpretar segundo sexo e idade. Estabelecer meta individualizada. Monitorar mudanças: perda de MG com preservação MME é padrão ideal.',
  updated_at = NOW()
WHERE id IN ('f0e7e520-6e01-437e-9274-5a5c90255d77', '1442a990-a9f0-4cd3-85b1-f3d22705e469', 'e0dbadd4-b307-43ff-8438-c24fa0876d10');

-- Gordura visceral
UPDATE score_items SET
  clinical_relevance = 'Gordura visceral (intra-abdominal, ao redor de vísceras) é o depósito adiposo mais metabolicamente ativo e deletério. Secreta adipocinas pró-inflamatórias, sofre lipólise elevada liberando ácidos graxos livres na circulação portal → resistência insulínica hepática, esteatose, dislipidemia aterogênica, estado pró-inflamatório. Associa-se a síndrome metabólica, diabetes tipo 2, doenças cardiovasculares, esteatose hepática, câncer. Avaliada por DEXA, TC, RM (padrão-ouro >100cm² área visceral = risco elevado), bioimpedância (nível >12-15 = risco). Responde bem a intervenções: perda ponderal modesta (5-10%) reduz preferencialmente gordura visceral. Exercício, especialmente intervalado e resistido, é particularmente eficaz.',
  patient_explanation = 'Gordura visceral é aquela que fica DENTRO do abdômen, ao redor dos órgãos (fígado, intestino, pâncreas), diferente da gordura subcutânea que fica debaixo da pele. É a gordura mais perigosa: libera substâncias inflamatórias, prejudica o fígado, aumenta resistência à insulina, eleva colesterol ruim, aumenta risco de diabetes e infarto. Barriga "dura" sugere gordura visceral; barriga "mole" é mais subcutânea. A boa notícia: gordura visceral RESPONDE MUITO BEM a exercícios e alimentação adequada - é a primeira a diminuir quando você emagrece de forma saudável.',
  conduct = 'Avaliar por DEXA (área cm²: risco se >100cm²) ou bioimpedância (nível: risco se >12-15 depende do equipamento). Gordura visceral elevada demanda investigação metabólica completa. Intervenção prioritária: dieta com controle glicêmico, exercício (resistido + intervalado HIIT), sono, controle estresse, considerar jejum intermitente. Monitorar redução preferencial de gordura visceral (maior redução que gordura subcutânea).',
  updated_at = NOW()
WHERE id IN ('d3f0556e-53cb-4a29-b825-a84bc8f38f3e', 'a4e0dc5b-e931-4126-a8b7-ac530f86aa68');

-- Razão androide/ginoide
UPDATE score_items SET
  clinical_relevance = 'Razão androide/ginoide (A/G) avalia distribuição de gordura: androide (abdominal) vs ginoide (glúteo-femoral). Valores elevados (>1.0 H, >0.8 M) indicam distribuição androide de risco metabólico. Reflete perfil hormonal: testosterona/cortisol favorecem androide; estrogênio favorece ginoide. Mulheres na menopausa migram de ginoide para androide. Razão A/G correlaciona-se com resistência insulínica, síndrome metabólica, risco cardiovascular, independente de IMC.',
  patient_explanation = 'Essa relação mostra se sua gordura está mais concentrada na região abdominal (androide, tipo maçã, mais arriscada) ou nas coxas e glúteos (ginoide, tipo pera, menos arriscada). Valores altos indicam acúmulo abdominal perigoso. Padrão muda com hormônios: homens naturalmente têm padrão androide; mulheres jovens têm ginoide, mas na menopausa migram para androide devido à queda de estrogênio.',
  conduct = 'Avaliar por DEXA. Interpretar: >1.0 (H) ou >0.8 (M) indica risco metabólico. Correlacionar com perfil hormonal. Intervenção visa reduzir gordura androide preferencialmente.',
  updated_at = NOW()
WHERE id IN ('2265fbe7-fc73-48e1-ac8d-c8e239d29e3c', 'bc698d3d-c9da-4276-b0d1-677ebd1fdf95');

-- Gordura de tronco
UPDATE score_items SET
  clinical_relevance = 'Percentual de gordura de tronco (%GT) avalia adiposidade central (tórax, abdômen). Correlaciona-se com gordura visceral e risco cardiometabólico. Valores elevados associam-se a síndrome metabólica, diabetes, doenças cardiovasculares. Meta: <30% em homens, <35% em mulheres (valores variam conforme referência).',
  patient_explanation = 'O percentual de gordura do tronco (peito, abdômen, costas) indica quanto da gordura corporal total está concentrada no centro do corpo. Valores altos indicam risco metabólico aumentado, pois gordura central (especialmente abdominal) é mais perigosa que gordura periférica (braços, pernas).',
  conduct = 'Avaliar por DEXA ou bioimpedância. Valores elevados requerem intervenção para redução preferencial de gordura central através de exercício e dieta.',
  updated_at = NOW()
WHERE id = 'e5b6d407-4744-48eb-8efb-bdae72dc34b9';

-- Items de água corporal
UPDATE score_items SET
  clinical_relevance = 'Água corporal total (ACT) representa ~50-60% peso corporal (maior em homens e indivíduos musculosos, menor em mulheres e obesos). ACT distribui-se em intracelular (AIC, ~60-65% ACT, dentro das células) e extracelular (AEC, ~35-40% ACT, plasma e interstício). Razão AEC/ACT aumentada (>0.40) indica retenção hídrica extracelular (edema, insuficiência cardíaca, renal, hepática, desnutrição, inflamação). ACT% baixa sugere desidratação ou excesso de massa gorda (gordura contém pouca água). Monitorar durante perda ponderal: perda rápida inicial é frequentemente água, não gordura.',
  patient_explanation = 'Seu corpo é composto majoritariamente de água (50-60% do peso). Essa água fica dentro das células (intracelular) e fora delas (extracelular, no sangue e entre os tecidos). A distribuição da água nos diz sobre sua hidratação e saúde: excesso de água extracelular indica inchaço (retenção hídrica); água total baixa pode indicar desidratação ou excesso de gordura (gordura tem pouca água). Durante emagrecimento, a perda inicial rápida é geralmente água, não gordura - por isso o peso pode cair muito no início e depois estabilizar.',
  conduct = 'Avaliar ACT por bioimpedância: litros absolutos e percentual do peso. Valores de referência: ~60% H, ~50-55% M. Avaliar AIC, AEC, razão AEC/ACT. Razão AEC/ACT >0.40 sugere retenção hídrica (investigar causas: cardíaca, renal, hepática, nutricional, inflamatória). Educar paciente sobre variações de água vs mudanças reais em composição.',
  updated_at = NOW()
WHERE id IN ('8bf6396f-b07c-4c93-82c4-b69ec8e87aa2', '2eafc97e-7134-49d7-b34c-6903192383d5',
             'f28c8a35-fa77-49a0-a69f-ba649fcffef2', 'e79394a0-3d1f-4d18-8c55-d6a20179017a',
             '7857982a-8b2b-401f-bf42-8db8a3af264b', '806cb0c4-1954-4567-b727-55bf4560553e');

-- Linkar artigos aos items da Parte 2
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, si.id
FROM articles a
CROSS JOIN (
  SELECT unnest(ARRAY[
    '936855d1-e715-4171-a061-3bcc500957f7', 'c3ee7527-104b-467a-a87e-b4587192005b',
    '421db9cd-d75d-4f0a-8188-216affee192f', '2ff747e7-e48d-4207-85ee-168b1c14e35e',
    'f0e7e520-6e01-437e-9274-5a5c90255d77', 'd3f0556e-53cb-4a29-b825-a84bc8f38f3e'
  ]::uuid[]) AS id
) si
WHERE a.title ILIKE '%emagrecimento%'
   OR a.title ILIKE '%composição%'
   OR a.title ILIKE '%metabol%'
   OR a.title ILIKE '%exercício%'
LIMIT 100
ON CONFLICT DO NOTHING;
