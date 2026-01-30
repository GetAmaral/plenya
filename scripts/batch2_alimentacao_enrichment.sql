-- ============================================
-- BATCH 2: ENRIQUECIMENTO - ALIMENTAÇÃO
-- 23 itens do grupo "Alimentação"
-- Data: 2026-01-27
-- ============================================

-- Item 1: Introdução alimentar
UPDATE score_items
SET
  clinical_relevance = 'A introdução alimentar representa período crítico para programação metabólica, desenvolvimento do sistema imunológico e estabelecimento da microbiota intestinal. A janela imunológica entre 4-7 meses é fundamental para tolerância oral e prevenção de alergias alimentares. Introdução precoce (antes de 4 meses) associa-se a maior risco de alergia, eczema e obesidade infantil, enquanto introdução tardia (após 7 meses) relaciona-se a deficiências nutricionais e maior risco de reações alérgicas. Estudos recentes demonstram que exposição controlada a alérgenos durante janela imunológica reduz significativamente incidência de alergia alimentar, especialmente amendoim, ovo e trigo. Baby-led weaning versus papinhas tradicionais influencia desenvolvimento de habilidades motoras orais, autorregulação alimentar e preferências alimentares futuras. Avaliação detalhada da introdução alimentar permite identificar fatores de risco para desenvolvimento de distúrbios alimentares, alergias, intolerâncias e padrões alimentares inadequados na vida adulta.',
  patient_explanation = 'A forma como a alimentação foi introduzida na sua infância tem impacto duradouro na sua saúde atual. O período entre 4-7 meses de vida é chamado de "janela imunológica" - quando o corpo aprende a reconhecer e aceitar diferentes alimentos sem desenvolver alergias. Se a introdução foi muito precoce (antes de 4 meses) ou muito tardia (após 7 meses), pode ter aumentado risco de alergias e problemas digestivos. O tipo de introdução alimentar também importa: se você comeu pedaços de alimentos sólidos desde cedo ou apenas papinhas, isso influenciou seu desenvolvimento motor oral e suas preferências alimentares atuais. Compreender seu histórico de introdução alimentar ajuda explicar possíveis intolerâncias, alergias ou dificuldades alimentares que você apresenta hoje.',
  conduct = 'Na anamnese, investigar: época de início (mês de vida), alimentos introduzidos primeiro, presença de reações adversas, método utilizado (papinhas versus BLW), amamentação exclusiva até 6 meses, presença de eczema ou reações alérgicas. Correlacionar com sintomas gastrointestinais atuais, alergias alimentares, padrões de restrição alimentar. Em casos de história de introdução inadequada associada a sintomas atuais, considerar investigação de permeabilidade intestinal, disbiose e sensibilidades alimentares retardadas.',
  updated_at = NOW()
WHERE id = '71cc4a7b-5ebd-437a-bbee-6f85105ac8dc';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('592f062d-2e8d-4ffa-b88b-b7a95d93bc36', '71cc4a7b-5ebd-437a-bbee-6f85105ac8dc'),
  ('cfd1073e-d4a8-45c1-999a-b96e75223902', '71cc4a7b-5ebd-437a-bbee-6f85105ac8dc'),
  ('c287dfbe-eee5-44b8-919b-3e957a725376', '71cc4a7b-5ebd-437a-bbee-6f85105ac8dc')
ON CONFLICT DO NOTHING;

-- Item 2: Qualidade da alimentação na idade pré-escolar
UPDATE score_items
SET
  clinical_relevance = 'A fase pré-escolar (2-6 anos) é período de rápido crescimento cerebral, desenvolvimento cognitivo e consolidação de preferências alimentares que persistem até vida adulta. Neofobia alimentar (recusa de novos alimentos) atinge pico entre 2-3 anos como mecanismo evolutivo de proteção, mas pode resultar em restrição alimentar inadequada se não manejada corretamente. Qualidade nutricional nesta fase impacta diretamente desenvolvimento neurológico, função imunológica, crescimento ósseo e programação metabólica. Déficits de ferro, zinco, ômega-3 e vitaminas B comprometem mielinização, desenvolvimento cognitivo e regulação emocional. Consumo excessivo de ultraprocessados, açúcares e gorduras trans correlaciona-se com obesidade infantil, resistência insulínica precoce e alterações comportamentais (hiperatividade, déficit atencional). Padrões familiares de alimentação estabelecidos nesta fase - horários irregulares, alimentação assistindo TV, uso de comida como recompensa - perpetuam-se como hábitos disfuncionais na vida adulta.',
  patient_explanation = 'Entre 2 e 6 anos, o cérebro está em desenvolvimento acelerado e as preferências alimentares se formam - padrões que você mantém até hoje. Nesta idade, é comum recusar alimentos novos (neofobia), mas a forma como sua família lidou com isso influenciou sua relação atual com comida. Se você comeu muitos alimentos ultraprocessados, doces e frituras nesta fase, pode ter desenvolvido preferência por sabores intensos e dificuldade em apreciar alimentos naturais. A falta de nutrientes importantes (ferro, zinco, ômega-3, vitaminas B) neste período pode ter afetado seu desenvolvimento cerebral, capacidade de concentração e regulação emocional. Hábitos como comer assistindo TV, receber comida como recompensa ou ter horários irregulares de refeições provavelmente começaram nesta época e podem estar contribuindo para problemas alimentares atuais.',
  conduct = 'Investigar: variedade alimentar (número de grupos alimentares consumidos), frequência de ultraprocessados, consumo de frutas/verduras, presença de neofobia alimentar, hábitos familiares (refeições em família, uso de telas durante refeições, comida como recompensa/punição). Correlacionar com sintomas atuais: compulsão alimentar, restrições alimentares, preferências por doces/salgados, dificuldade com vegetais. Orientar sobre reprogramação de paladares através de exposição repetida, mindful eating e ressignificação de padrões disfuncionais estabelecidos na infância.',
  updated_at = NOW()
WHERE id = '6d98fb8b-d233-457f-bfc9-8028fbdf23f4';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('cfd1073e-d4a8-45c1-999a-b96e75223902', '6d98fb8b-d233-457f-bfc9-8028fbdf23f4'),
  ('deb302ab-195b-41c0-8657-b299fd610a7f', '6d98fb8b-d233-457f-bfc9-8028fbdf23f4'),
  ('c287dfbe-eee5-44b8-919b-3e957a725376', '6d98fb8b-d233-457f-bfc9-8028fbdf23f4')
ON CONFLICT DO NOTHING;

-- Item 3: Pré-natal
UPDATE score_items
SET
  clinical_relevance = 'A nutrição materna durante gestação determina programação epigenética fetal, desenvolvimento neurológico, função imunológica e risco metabólico ao longo da vida (conceito DOHaD - Developmental Origins of Health and Disease). Deficiências nutricionais maternas associam-se a defeitos de tubo neural (ácido fólico), prematuridade (ômega-3, vitamina D), baixo peso ao nascer (proteína, ferro, zinco) e maior risco de doenças crônicas na vida adulta da prole (diabetes, hipertensão, obesidade). Excesso nutricional materno - ganho ponderal excessivo, hiperglicemia gestacional, dieta inflamatória - programa metabolicamente o feto para resistência insulínica, obesidade e síndrome metabólica. Micronutrientes críticos incluem folato (metilação de DNA), colina (desenvolvimento cerebral), iodo (função tireoidiana fetal), ferro (oxigenação fetal), vitamina D (desenvolvimento imunológico) e DHA (neurogênese). Disbiose materna e dieta pró-inflamatória alteram colonização microbiana do recém-nascido, impactando desenvolvimento imunológico e risco de atopia.',
  patient_explanation = 'A alimentação da sua mãe durante a gravidez programou geneticamente seu metabolismo, sistema imunológico e até sua tendência a doenças crônicas - conceito conhecido como "origem desenvolvimentista da saúde". Se sua mãe teve deficiências de nutrientes importantes (ácido fólico, ferro, zinco, ômega-3, vitamina D), você pode ter nascido com programação metabólica que favorece problemas de saúde atuais. Por outro lado, se houve ganho excessivo de peso, diabetes gestacional ou alimentação inflamatória, você foi "programado" para maior risco de obesidade, resistência à insulina e doenças metabólicas. A saúde intestinal e a dieta da sua mãe também determinaram as primeiras bactérias que colonizaram seu intestino ao nascer, influenciando seu sistema imunológico e risco de alergias. Compreender a nutrição materna pré-natal ajuda explicar tendências metabólicas e imunológicas que você apresenta hoje.',
  conduct = 'Investigar: nutrição materna pré-gestacional e durante gestação, suplementação (ácido fólico, ferro, ômega-3, vitamina D), ganho ponderal gestacional, diabetes gestacional, dieta materna (qualidade, grupos alimentares), intercorrências gestacionais. Correlacionar com condições atuais do paciente: tendência a obesidade, resistência insulínica, alergias, doenças autoimunes, alterações tireoidianas. Em casos de história gestacional inadequada, considerar avaliação epigenética, permeabilidade intestinal e modulação de metilação através de nutrientes doadores de metil (folato, B12, colina, betaína).',
  updated_at = NOW()
WHERE id = '94b35411-dd63-47c5-b1a0-3b307e938d1e';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('cfd1073e-d4a8-45c1-999a-b96e75223902', '94b35411-dd63-47c5-b1a0-3b307e938d1e'),
  ('deb302ab-195b-41c0-8657-b299fd610a7f', '94b35411-dd63-47c5-b1a0-3b307e938d1e'),
  ('c287dfbe-eee5-44b8-919b-3e957a725376', '94b35411-dd63-47c5-b1a0-3b307e938d1e')
ON CONFLICT DO NOTHING;

-- Item 4: O que pais comiam antes e durante gestação
UPDATE score_items
SET
  clinical_relevance = 'A nutrição parental pré-concepcional influencia qualidade gamética, fertilidade, desenvolvimento embrionário e programação epigenética da prole através de mecanismos de herança transgeracional. Estudos demonstram que dieta paterna rica em gorduras saturadas e pobre em micronutrientes altera metilação de DNA espermático, impactando metabolismo da prole (obesidade, diabetes tipo 2). Dieta materna pré-concepcional determina reservas nutricionais para primeiro trimestre gestacional (período de maior vulnerabilidade para defeitos congênitos) e qualidade oocitária. Exposição parental a disruptores endócrinos, álcool, tabaco e dieta pró-inflamatória aumenta estresse oxidativo gamético e risco de alterações epigenéticas transmissíveis. Deficiências de folato, B12, colina e zinco comprometem metilação adequada de DNA e expressão gênica fetal. Estado nutricional e metabólico parental até 3 meses pré-concepção impacta desfechos gestacionais, peso ao nascer e saúde metabólica da proge ao longo da vida.',
  patient_explanation = 'A alimentação dos seus pais antes mesmo de você ser concebido influenciou sua genética e metabolismo - fenômeno chamado herança transgeracional. O que seu pai comia afetou a qualidade dos espermatozoides e programou marcas genéticas transmitidas a você. O que sua mãe comia antes da gravidez determinou as reservas nutricionais disponíveis nos primeiros meses de gestação (período crítico para formação dos órgãos) e a qualidade dos óvulos. Se seus pais tinham alimentação pobre em nutrientes, rica em alimentos inflamatórios, ou exposição a álcool e toxinas, isso pode ter alterado marcas genéticas (epigenética) que você carrega hoje e que influenciam seu metabolismo, tendência a ganho de peso, diabetes e doenças crônicas. Entender a nutrição dos seus pais antes da concepção ajuda explicar predisposições metabólicas e de saúde que não são explicadas apenas por sua genética direta.',
  conduct = 'Investigar: padrão alimentar de ambos os pais 3-6 meses pré-concepção, peso e estado nutricional parental, consumo de álcool/tabaco, exposição a toxinas ambientais, suplementação pré-concepcional, intervalo entre gestações (depleção materna). Correlacionar com condições atuais: tendências metabólicas inexplicáveis por genética direta, obesidade de difícil controle, resistência insulínica precoce, alterações de metilação. Considerar investigação de polimorfismos relacionados a ciclo de metilação (MTHFR, MTR, MTRR) e suporte com nutrientes doadores de metil.',
  updated_at = NOW()
WHERE id = 'a35a3012-60ea-4659-a742-fbbcb741c6db';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('cfd1073e-d4a8-45c1-999a-b96e75223902', 'a35a3012-60ea-4659-a742-fbbcb741c6db'),
  ('deb302ab-195b-41c0-8657-b299fd610a7f', 'a35a3012-60ea-4659-a742-fbbcb741c6db')
ON CONFLICT DO NOTHING;

-- Item 5: Intolerâncias
UPDATE score_items
SET
  clinical_relevance = 'Intolerâncias alimentares representam reações adversas não-imunológicas a alimentos, geralmente por deficiências enzimáticas (lactase, dissacaridases), alterações de transporte (frutose, sorbitol) ou sensibilidade a compostos bioativos (histamina, salicilatos, aminas biogênicas). Diferem de alergias IgE-mediadas por ausência de resposta imunológica imediata e apresentação clínica variável e dose-dependente. Intolerância à lactose afeta 65-70% da população mundial adulta (hipolactasia primária), mas sintomas são frequentemente confundidos com síndrome do intestino irritável ou SIBO. Má absorção de frutose e FODMAPs correlaciona-se fortemente com SII, distensão abdominal e disbiose. Intolerância à histamina (déficit de DAO - diamino oxidase) manifesta-se como sintomas sistêmicos: enxaqueca, urticária, sintomas gastrointestinais, taquicardia. Investigação adequada de intolerâncias evita restrições alimentares desnecessárias que comprometem diversidade microbiana e adequação nutricional, além de identificar causas tratáveis (SIBO, permeabilidade intestinal, disbiose).',
  patient_explanation = 'Intolerâncias alimentares são diferentes de alergias: não envolvem o sistema imunológico, mas sim dificuldades do corpo em digerir ou processar certos alimentos. A mais comum é intolerância à lactose (açúcar do leite), que afeta 70% dos adultos e causa gases, distensão e diarreia. Você também pode ter dificuldade em absorver frutose (açúcar das frutas) ou FODMAPs (tipos de carboidratos presentes em muitos alimentos), causando sintomas parecidos com intestino irritável. Outra intolerância importante é à histamina, presente em alimentos fermentados, envelhecidos e processados, que pode causar dores de cabeça, coceira, problemas digestivos e até taquicardia. Identificar suas intolerâncias evita que você faça restrições alimentares desnecessárias (que prejudicam suas bactérias intestinais) e ajuda encontrar a causa real dos sintomas - que pode ser tratável.',
  conduct = 'Investigar: alimentos suspeitos, sintomas relacionados (gastrointestinais versus sistêmicos), latência dos sintomas (minutos versus horas), dose-dependência, histórico familiar de intolerâncias. Diferenciar de alergia IgE-mediada (reação imediata, risco anafilático). Considerar testes: teste respiratório de hidrogênio para lactose/frutose, dosagem de DAO sérica para histamina, dieta de eliminação FODMAPs. Investigar causas secundárias: SIBO, permeabilidade intestinal, disbiose, deficiências enzimáticas. Evitar restrições prolongadas sem investigação adequada.',
  updated_at = NOW()
WHERE id = 'd8696390-b619-4752-b366-5c63f4730811';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('592f062d-2e8d-4ffa-b88b-b7a95d93bc36', 'd8696390-b619-4752-b366-5c63f4730811'),
  ('44a22cf0-71bd-41aa-986e-7f99f0f77b1a', 'd8696390-b619-4752-b366-5c63f4730811'),
  ('09d14865-0a33-4d31-86d4-00666bc1ba67', 'd8696390-b619-4752-b366-5c63f4730811')
ON CONFLICT DO NOTHING;

-- Item 6: Lactose
UPDATE score_items
SET
  clinical_relevance = 'A hipolactasia (deficiência de lactase) é a intolerância alimentar mais prevalente mundialmente, afetando aproximadamente 65-70% da população adulta global com variação étnica significativa (>90% em asiáticos, 60-80% em afro-descendentes, 15-30% em caucasianos). Hipolactasia primária resulta de down-regulation genética da expressão de lactase após desmame (persistência versus não-persistência de lactase), enquanto hipolactasia secundária ocorre por dano à borda em escova intestinal (gastroenterites, doença celíaca, SIBO, IBD). Lactose não digerida é fermentada por bactérias colônicas, produzindo hidrogênio, metano e ácidos graxos de cadeia curta, causando distensão, flatulência, dor abdominal e diarréia osmótica. Teste respiratório de hidrogênio confirma diagnóstico com sensibilidade >90%. Restrição de lactose melhora sintomas mas pode comprometer ingestão de cálcio e vitamina D, exigindo suplementação ou fontes alternativas. Produtos fermentados (iogurte, kefir) são melhor tolerados devido à pré-digestão bacteriana da lactose.',
  patient_explanation = 'Intolerância à lactose acontece quando seu intestino não produz enzima suficiente (lactase) para digerir o açúcar do leite. Cerca de 70% dos adultos no mundo têm esse problema - é geneticamente programado para acontecer após o desmame. Quando você consome lactose sem ter enzima suficiente, as bactérias intestinais fermentam este açúcar, produzindo gases, distensão, dor abdominal e diarréia. Existe um teste simples (teste do hidrogênio no ar expirado) que confirma o diagnóstico com mais de 90% de precisão. Se você tem intolerância, precisa reduzir lactose mas cuidar para não faltar cálcio e vitamina D na dieta. Produtos fermentados como iogurte e kefir geralmente são bem tolerados porque as bactérias já "pré-digerem" a lactose, e existem leites sem lactose e enzimas de lactase que você pode usar.',
  conduct = 'Investigar: quantidade e tipo de laticínios que desencadeiam sintomas, tolerância a produtos fermentados (iogurte, queijos curados), histórico de uso de lactase exógena, adequação de cálcio/vitamina D na dieta. Solicitar teste respiratório de hidrogênio para lactose se diagnóstico incerto. Diferenciar de alergia à proteína do leite (APLV) e sensibilidade ao A1 beta-caseína. Orientar sobre: uso de lactase exógena, produtos sem lactose, produtos fermentados, fontes alternativas de cálcio (folhas verdes escuras, peixes com ossos, fortificados). Investigar hipolactasia secundária se início súbito (SIBO, doença celíaca).',
  updated_at = NOW()
WHERE id = '5cac2737-c0c5-47d9-aaa8-2ca4fad39f47';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('592f062d-2e8d-4ffa-b88b-b7a95d93bc36', '5cac2737-c0c5-47d9-aaa8-2ca4fad39f47'),
  ('44a22cf0-71bd-41aa-986e-7f99f0f77b1a', '5cac2737-c0c5-47d9-aaa8-2ca4fad39f47'),
  ('b777f9d9-5a13-4856-9da8-548d1fd843f6', '5cac2737-c0c5-47d9-aaa8-2ca4fad39f47')
ON CONFLICT DO NOTHING;

-- Item 7: Tipo de alimentos que mais gosta
UPDATE score_items
SET
  clinical_relevance = 'Preferências alimentares refletem programação neurobiológica estabelecida na infância, influências culturais, microbiota intestinal (eixo intestino-cérebro) e alterações de palatabilidade associadas a condições metabólicas. Predileção excessiva por doces correlaciona-se com disbiose (predomínio de Candida e bactérias sacarolíticas), resistência insulínica, deficiências de cromo/magnésio e padrões de recompensa dopaminérgica disfuncionais. Preferência por alimentos salgados/gordurosos pode indicar disfunção adrenal, deficiências minerais (sódio, magnésio) ou adaptação a dieta cronicamente inadequada em nutrientes densos. Aversão a vegetais amargos relaciona-se ao número de papilas gustativas TAS2R (bitter taste receptors), variação genética individual, e exposição limitada durante janela de desenvolvimento do paladar. Microbiota disbiótica influencia preferências através de produção de metabólitos neuroativos que modulam craving e saciedade via nervo vago.',
  patient_explanation = 'Os alimentos que você mais gosta não são escolhas aleatórias - eles refletem programação do seu cérebro desde a infância, estado das suas bactérias intestinais e até deficiências nutricionais. Se você tem craving intenso por doces, pode ser sinal de desequilíbrio das bactérias intestinais (especialmente crescimento de Candida), resistência à insulina ou falta de minerais como cromo e magnésio. Preferência exagerada por salgados e gorduras pode indicar estresse das glândulas adrenais ou deficiências minerais. Se você não gosta de vegetais amargos, pode ter mais receptores de sabor amargo na língua (genético) ou ter tido pouca exposição a esses alimentos quando criança. Suas bactérias intestinais também influenciam suas vontades através de substâncias que produzem e que afetam seu cérebro via nervo vago.',
  conduct = 'Investigar: intensidade de craving por doces/salgados/gordurosos, horários predominantes de fome específica, correlação com estados emocionais, saciedade após refeições, aversões alimentares. Correlacionar doces com: glicemia de jejum, HbA1c, resistência insulínica, candidiase recorrente, ferritina. Preferência por salgados: investigar função adrenal (cortisol salivar), pressão arterial, magnésio. Orientar sobre reprogramação gradual do paladar através de redução progressiva de estímulos intensos, aumento de exposições (12-15x para aceitação de novos alimentos), modulação de microbiota e correção de deficiências nutricionais subjacentes.',
  updated_at = NOW()
WHERE id = '3dbcb5dc-88da-41f1-917b-ec74d5c19dfd';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('c287dfbe-eee5-44b8-919b-3e957a725376', '3dbcb5dc-88da-41f1-917b-ec74d5c19dfd'),
  ('cfd1073e-d4a8-45c1-999a-b96e75223902', '3dbcb5dc-88da-41f1-917b-ec74d5c19dfd'),
  ('b777f9d9-5a13-4856-9da8-548d1fd843f6', '3dbcb5dc-88da-41f1-917b-ec74d5c19dfd')
ON CONFLICT DO NOTHING;

-- Item 8: Qualidade da alimentação na adolescência
UPDATE score_items
SET
  clinical_relevance = 'A adolescência representa período de vulnerabilidade nutricional devido a demandas aumentadas por crescimento acelerado, desenvolvimento puberal, expansão de massa óssea (pico de massa óssea aos 18-25 anos) e maturação cerebral (córtex pré-frontal até 25 anos). Simultaneamente, influências psicossociais promovem padrões alimentares inadequados: refeições irregulares, fast-food, bebidas açucaradas, dietas restritivas e transtornos alimentares emergentes. Deficiências nutricionais prevalentes - ferro (meninas pós-menarca), cálcio, vitamina D, zinco, ômega-3 - comprometem desenvolvimento ósseo, função cognitiva, regulação emocional e desempenho acadêmico. Padrões alimentares estabelecidos na adolescência (dieta ocidental versus mediterrânea) predizem risco cardiovascular e metabólico na vida adulta. Consumo elevado de ultraprocessados associa-se a inflamação crônica de baixo grau, acne inflamatória, alterações de humor, síndrome pré-menstrual severa e maior risco de obesidade/síndrome metabólica.',
  patient_explanation = 'A adolescência é fase de alto risco nutricional: seu corpo precisava de muitos nutrientes para crescer, desenvolver ossos fortes e amadurecer o cérebro, mas ao mesmo tempo é quando a maioria come pior - fast-food, refrigerantes, refeições puladas, dietas malucas. Se você teve alimentação inadequada nesta fase, pode ter comprometido sua massa óssea máxima (que previne osteoporose futura), desenvolvimento do córtex pré-frontal (área do cérebro responsável por controle de impulsos e tomada de decisões) e reservas de nutrientes importantes. Meninas que menstruam têm risco especial de deficiência de ferro. Alimentação rica em ultraprocessados na adolescência aumenta inflamação crônica, piora acne, altera humor, intensifica TPM e programa metabolismo para ganho de peso e problemas cardiovasculares na vida adulta. Os padrões alimentares que você estabeleceu nesta época provavelmente persistem até hoje.',
  conduct = 'Investigar: padrão alimentar predominante (casa versus fast-food), frequência de refeições, consumo de ultraprocessados/bebidas açucaradas, presença de dietas restritivas ou transtornos alimentares, prática esportiva e adequação nutricional. Correlacionar com: densidade mineral óssea atual, perfil lipídico, resistência insulínica, histórico de acne severa, alterações menstruais, sintomas de ansiedade/depressão. Avaliar marcadores: ferritina (meninas), vitamina D, zinco, ômega-3 (índice ômega-3). Orientar sobre janela de oportunidade para otimização de massa óssea até 25 anos e impacto duradouro de padrões estabelecidos na adolescência.',
  updated_at = NOW()
WHERE id = 'df9627bd-e69a-43fe-a925-7f368f20e2f6';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('cfd1073e-d4a8-45c1-999a-b96e75223902', 'df9627bd-e69a-43fe-a925-7f368f20e2f6'),
  ('deb302ab-195b-41c0-8657-b299fd610a7f', 'df9627bd-e69a-43fe-a925-7f368f20e2f6'),
  ('700704d7-ba2a-4123-b677-b0f221b167c1', 'df9627bd-e69a-43fe-a925-7f368f20e2f6')
ON CONFLICT DO NOTHING;

-- Item 9: Qualidade da alimentação na idade escolar
UPDATE score_items
SET
  clinical_relevance = 'A fase escolar (6-12 anos) é crítica para consolidação de hábitos alimentares, desenvolvimento de função executiva e estabelecimento de relação saudável com alimentos. Qualidade nutricional impacta diretamente desempenho acadêmico, comportamento, desenvolvimento cognitivo e função imunológica. Estudos demonstram correlação entre consumo de café da manhã rico em proteínas e melhora de atenção, memória de trabalho e desempenho escolar. Déficits de ferro (mesmo sem anemia) comprometem função cognitiva, concentração e rendimento acadêmico. Consumo excessivo de açúcares simples e corantes artificiais associa-se a hiperatividade, déficit atencional e alterações comportamentais. Ômega-3 (EPA/DHA) é essencial para neurotransmissão, função sináptica e neuroplasticidade. Ambiente alimentar escolar (cantina, merenda) e influência de pares moldam preferências e comportamentos alimentares duradouros. Bullying relacionado a peso, imagem corporal e padrões alimentares restritivos frequentemente emergem nesta fase.',
  patient_explanation = 'Entre 6 e 12 anos, sua alimentação influenciou diretamente seu desempenho na escola, comportamento e desenvolvimento cerebral. Se você tomava café da manhã com proteínas, provavelmente tinha melhor atenção e memória nas aulas. Falta de ferro (mesmo sem anemia) prejudica concentração e rendimento escolar. Consumo excessivo de açúcares e corantes artificiais está associado a hiperatividade e dificuldade de atenção. Ômega-3 era essencial para seu cérebro funcionar bem - comunicação entre neurônios e aprendizado. O que tinha disponível na cantina da escola e o que seus colegas comiam influenciaram fortemente suas escolhas e preferências alimentares que você mantém hoje. Se você sofreu bullying por causa do peso ou desenvolveu preocupação excessiva com imagem corporal nesta fase, pode ter iniciado padrões alimentares restritivos que persistem.',
  conduct = 'Investigar: hábito de café da manhã, qualidade da alimentação escolar (merenda versus cantina), consumo de doces/refrigerantes, histórico de dificuldades de aprendizado ou déficit atencional, episódios de bullying relacionados a peso, desenvolvimento de preocupações com imagem corporal. Correlacionar com: histórico acadêmico, diagnóstico de TDAH, perfil de ferritina, ômega-3, zinco. Investigar início de padrões restritivos ou compulsivos nesta fase. Orientar sobre impacto duradouro de déficits nutricionais durante desenvolvimento escolar e possibilidade de intervenção atual para otimização cognitiva.',
  updated_at = NOW()
WHERE id = 'a91d9533-7553-43b5-b3b5-2ac72127eea0';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('cfd1073e-d4a8-45c1-999a-b96e75223902', 'a91d9533-7553-43b5-b3b5-2ac72127eea0'),
  ('deb302ab-195b-41c0-8657-b299fd610a7f', 'a91d9533-7553-43b5-b3b5-2ac72127eea0'),
  ('700704d7-ba2a-4123-b677-b0f221b167c1', 'a91d9533-7553-43b5-b3b5-2ac72127eea0')
ON CONFLICT DO NOTHING;

-- Item 10: Qualidade da alimentação na vida adulta
UPDATE score_items
SET
  clinical_relevance = 'A qualidade alimentar na vida adulta é determinante primário de longevidade, healthspan e risco de doenças crônicas não-transmissíveis (cardiovasculares, diabetes tipo 2, câncer, demência). Dieta mediterrânea demonstra consistentemente redução de mortalidade total, eventos cardiovasculares, declínio cognitivo e marcadores inflamatórios. Consumo elevado de ultraprocessados (>20% das calorias) associa-se a aumento de 62% no risco de mortalidade total e correlaciona com obesidade, síndrome metabólica, câncer colorretal e depressão. Padrão pró-inflamatório (alto índice dietético inflamatório) prediz resistência insulínica, aterosclerose e envelhecimento acelerado. Restrição calórica moderada e jejum intermitente ativam vias de longevidade (AMPK, SIRT1, autofagia) e melhoram parâmetros metabólicos. Micronutrientes críticos frequentemente deficientes em adultos: vitamina D, magnésio, ômega-3, vitamina B12 (especialmente >50 anos). Timing nutricional, ritmo circadiano alimentar e crononutrição influenciam metabolismo glicídico, lipídico e regulação ponderal.',
  patient_explanation = 'A qualidade da sua alimentação atual é o principal fator que determina quanto tempo você viverá com saúde e seu risco de doenças crônicas como diabetes, infarto, câncer e demência. Dieta mediterrânea (rica em vegetais, azeite, peixes, castanhas) reduz risco de morte e protege cérebro e coração. Se mais de 20% das suas calorias vêm de ultraprocessados, você tem 62% mais risco de morte prematura, além de maior chance de obesidade, diabetes e depressão. Alimentação que causa inflamação crônica acelera envelhecimento e aumenta risco de todas as doenças crônicas. Períodos de jejum moderado ativam mecanismos de longevidade no corpo. Adultos frequentemente têm deficiências de vitamina D, magnésio, ômega-3 e B12 (especialmente após 50 anos). O horário que você come também importa - comer alinhado com ritmo circadiano melhora metabolismo.',
  conduct = 'Investigar: padrão alimentar predominante (mediterrâneo, ocidental, plant-based), percentual de ultraprocessados, frequência de refeições caseiras versus industrializadas, consumo de vegetais/frutas/grãos integrais/gorduras saudáveis, hidratação, prática de jejum, horários de refeições (crononutrição). Solicitar: recordatório alimentar 24h ou diário alimentar 3 dias, questionário de frequência alimentar. Avaliar índice dietético inflamatório. Marcadores laboratoriais: vitamina D, magnésio eritrocitário, índice ômega-3, B12, homocisteína, perfil inflamatório (PCR-us, IL-6). Orientar sobre padrões alimentares baseados em evidência para longevidade e prevenção de doenças.',
  updated_at = NOW()
WHERE id = 'f2ab84f5-f39a-4683-8814-9f18e88dea60';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('cfd1073e-d4a8-45c1-999a-b96e75223902', 'f2ab84f5-f39a-4683-8814-9f18e88dea60'),
  ('3962bae9-a59b-48e1-b9b2-bbf9d8cd4118', 'f2ab84f5-f39a-4683-8814-9f18e88dea60'),
  ('deb302ab-195b-41c0-8657-b299fd610a7f', 'f2ab84f5-f39a-4683-8814-9f18e88dea60'),
  ('f943b301-717d-42d0-a236-9a924139c51a', 'f2ab84f5-f39a-4683-8814-9f18e88dea60')
ON CONFLICT DO NOTHING;

-- Item 11: Qual dieta e quando foi tentada, duração e resultados obtidos
UPDATE score_items
SET
  clinical_relevance = 'Histórico de dietas prévias revela padrões de restrição alimentar, efeito sanfona (weight cycling), metabolismo adaptativo e relação comportamental com comida. Weight cycling associa-se a resistência progressiva à perda ponderal, perda de massa magra, aumento de gordura visceral, alterações metabólicas persistentes e maior risco cardiovascular. Dietas muito restritivas (<1200 kcal/dia) induzem adaptação metabólica (redução de taxa metabólica basal em até 25%), aumento de grelina, redução de leptina e alterações de hormônios tireoidianos (T3 reverso). Padrão de múltiplas dietas seguidas de reganho sugere abordagens insustentáveis, falta de mudança comportamental ou causas metabólicas subjacentes não abordadas (resistência insulínica, hipotireoidismo, disbiose, deficiências nutricionais). Algumas intervenções dietéticas (cetogênica, jejum intermitente, paleo, low-FODMAP) podem ter indicações terapêuticas específicas além de perda ponderal. Compreender histórico dietético orienta abordagem individualizada e sustentável.',
  patient_explanation = 'Cada dieta que você já tentou deixou marcas no seu metabolismo e na sua relação com comida. Se você fez muitas dietas com perda e reganho de peso (efeito sanfona), seu metabolismo pode ter se adaptado reduzindo gasto energético em até 25%, tornando cada vez mais difícil emagrecer. Dietas muito restritivas alteram hormônios da fome (aumentam grelina e reduzem leptina), fazendo você sentir fome constante. Se você já fez várias dietas e sempre recuperou o peso, pode ser que as dietas fossem insustentáveis ou que existam problemas metabólicos que nunca foram tratados (resistência à insulina, tireoide, bactérias intestinais, deficiências nutricionais). Algumas dietas específicas (cetogênica, jejum intermitente, low-FODMAP) podem ter benefícios terapêuticos além do emagrecimento. Entender seu histórico de dietas ajuda criar abordagem personalizada e sustentável.',
  conduct = 'Investigar detalhadamente cada dieta prévia: tipo (low-carb, cetogênica, jejum, pontos, detox, etc), duração, perda ponderal alcançada, reganho após suspensão, sintomas durante dieta, motivo de abandono, número total de ciclos de perda/reganho. Identificar padrões: restrições extremas, exclusões alimentares múltiplas, uso de shakes/substitutos, dietas da moda. Avaliar impacto psicológico: desenvolvimento de compulsão, restrição cognitiva, culpa alimentar. Investigar causas metabólicas de insucesso: resistência insulínica (HOMA-IR), função tireoidiana (TSH, T4L, T3L, T3r), cortisol, disbiose. Orientar sobre metabolismo adaptativo e necessidade de abordagem sustentável e multimodal.',
  updated_at = NOW()
WHERE id = '74039b09-a41b-43a5-b0d6-d9567f68a829';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('3962bae9-a59b-48e1-b9b2-bbf9d8cd4118', '74039b09-a41b-43a5-b0d6-d9567f68a829'),
  ('2b4db9bc-a628-4c10-a978-5f0b4bd5607a', '74039b09-a41b-43a5-b0d6-d9567f68a829'),
  ('1744afcf-6266-4f21-9a71-45b6e4255d79', '74039b09-a41b-43a5-b0d6-d9567f68a829')
ON CONFLICT DO NOTHING;

-- Item 12: Onde e como come
UPDATE score_items
SET
  clinical_relevance = 'O contexto ambiental e comportamental das refeições influencia profundamente digestão, saciedade, escolhas alimentares e relação com comida. Comer assistindo telas (TV, celular, computador) associa-se a eating disinibido, redução de sinais de saciedade, aumento de consumo calórico (20-30% mais) e menor consciência alimentar (mindless eating). Velocidade de ingestão correlaciona-se inversamente com saciedade e diretamente com ganho ponderal - mastigação inadequada (<15-20 vezes) compromete digestão mecânica, liberação de enzimas salivares, sinalização de saciedade e aumenta carga digestiva gastrintestinal. Ambiente estressante durante refeições ativa sistema nervoso simpático, reduz fluxo sanguíneo gastrointestinal, compromete motilidade, secreção enzimática e pode precipitar sintomas dispépticos. Refeições sociais (versus solitárias) promovem melhor escolha alimentar, maior prazer, melhor mastigação. Comer de pé, na rua ou em trânsito associa-se a escolhas alimentares pobres, porções inadequadas e menor satisfação.',
  patient_explanation = 'Onde e como você come é tão importante quanto o que você come. Se você come assistindo TV, no celular ou trabalhando, seu cérebro não registra adequadamente a refeição - você come 20-30% mais sem perceber e não se sente satisfeito. Comer rápido, mastigando pouco, prejudica digestão, impede que hormônios da saciedade sejam liberados a tempo e aumenta chances de ganhar peso. Comer estressado ou com pressa ativa modo "luta ou fuga" do sistema nervoso, reduzindo fluxo de sangue para intestino, piorando digestão e causando má digestão e desconforto. Refeições feitas em companhia, sentado à mesa, sem distrações, promovem escolhas melhores, mastigação adequada e maior satisfação. Comer de pé, na rua ou no carro leva a escolhas piores e menor saciedade.',
  conduct = 'Investigar: local habitual de refeições (mesa, sofá, carro, em pé), presença de distrações (TV, celular, trabalho), companhia durante refeições (sozinho versus acompanhado), velocidade de ingestão, número de mastigações, tempo dedicado às refeições, estado emocional durante alimentação. Correlacionar com: sintomas digestivos (dispepsia, refluxo, distensão), controle de porções, percepção de saciedade, ganho ponderal. Orientar sobre: mindful eating, mastigação consciente (20-30x), ambiente tranquilo para refeições, refeições sentadas à mesa sem telas, prática de gratidão alimentar. Considerar treinamento formal de mindful eating se padrões muito disfuncionais.',
  updated_at = NOW()
WHERE id = '01498f3d-551a-4285-bc84-86d015569d31';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('cfd1073e-d4a8-45c1-999a-b96e75223902', '01498f3d-551a-4285-bc84-86d015569d31'),
  ('b777f9d9-5a13-4856-9da8-548d1fd843f6', '01498f3d-551a-4285-bc84-86d015569d31')
ON CONFLICT DO NOTHING;

-- Item 13: Quem prepara as refeições
UPDATE score_items
SET
  clinical_relevance = 'A autonomia culinária (food literacy) correlaciona-se positivamente com qualidade alimentar, consumo de alimentos in natura, diversidade nutricional e melhores desfechos de saúde. Indivíduos que preparam próprias refeições consomem significativamente mais vegetais, frutas, fibras e menos ultraprocessados, sódio e gorduras trans. Dependência de terceiros para alimentação (família, delivery, restaurantes) limita controle sobre ingredientes, porções, qualidade de gorduras e sódio. Refeições preparadas em casa permitem personalização para necessidades específicas (alergias, intolerâncias, preferências) e controle de métodos de cocção. Falta de habilidade culinária é barreira importante para adesão a orientações nutricionais e manutenção de mudanças alimentares sustentáveis. Dinâmica familiar alimentar - quem cozinha, quem decide cardápio, divisão de tarefas - reflete e influencia padrões alimentares de todos os membros. Cooking classes e meal planning são intervenções efetivas para melhoria de qualidade alimentar.',
  patient_explanation = 'Quem prepara sua comida tem enorme impacto na qualidade da sua alimentação. Pessoas que cozinham suas próprias refeições comem muito mais vegetais, frutas e alimentos naturais, e menos ultraprocessados, sódio e gorduras ruins. Se você depende de outras pessoas, delivery ou restaurantes, você perde controle sobre ingredientes, qualidade das gorduras, quantidade de sal e aditivos. Preparar sua própria comida permite adaptar para suas necessidades - alergias, intolerâncias, objetivos de saúde. Não saber cozinhar é uma das maiores barreiras para conseguir manter alimentação saudável a longo prazo. A dinâmica familiar também importa: quem cozinha em casa, quem decide o cardápio, como as tarefas são divididas - isso afeta a alimentação de toda família. Aprender a cozinhar e planejar refeições são formas comprovadas de melhorar alimentação.',
  conduct = 'Investigar: frequência de preparo próprio versus refeições prontas/delivery/restaurantes, habilidades culinárias (básicas, intermediárias, avançadas), tempo disponível para cozinhar, equipamentos disponíveis, dinâmica familiar (quem cozinha, compra, decide cardápio), meal planning. Avaliar barreiras: tempo, conhecimento, acesso a ingredientes, custo percebido. Correlacionar com: variedade alimentar, consumo de ultraprocessados, adequação nutricional. Orientar sobre: meal prep, batch cooking, receitas simples e rápidas, lista de compras planejada. Considerar encaminhamento para cooking classes, consultas com nutricionista para educação culinária, recursos online (vídeos, aplicativos de receitas).',
  updated_at = NOW()
WHERE id = '5511eb71-e1c3-44a7-9916-958c11dd788e';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('cfd1073e-d4a8-45c1-999a-b96e75223902', '5511eb71-e1c3-44a7-9916-958c11dd788e'),
  ('0108fb41-2fee-47f4-9d07-84a8de20c226', '5511eb71-e1c3-44a7-9916-958c11dd788e')
ON CONFLICT DO NOTHING;

-- Item 14: Recordatório detalhado das refeições
UPDATE score_items
SET
  clinical_relevance = 'Recordatório alimentar de 24 horas (ou diário alimentar de 3-7 dias) é ferramenta essencial para avaliação quantitativa e qualitativa da ingestão alimentar atual. Permite estimar adequação calórica, distribuição de macronutrientes, densidade de micronutrientes, carga glicêmica, índice inflamatório dietético, cronobiologia alimentar e identificar padrões disfuncionais. Análise revela: frequência e timing de refeições (impacto circadiano), presença de jejum prolongado ou beliscadas constantes, balanço proteico por refeição (otimização de síntese proteica requer 25-40g por refeição), adequação de fibras, variedade alimentar (diversidade é proxy de microbiota saudável), hidratação. Software de análise nutricional quantifica micronutrientes frequentemente deficientes: magnésio, zinco, vitamina D, B12, folato, ômega-3. Índice de qualidade da dieta (HEI - Healthy Eating Index) prediz desfechos de saúde. Recordatório também identifica gatilhos emocionais, horários de craving, associações comportamentais com alimentação.',
  patient_explanation = 'Detalhar tudo que você comeu nas últimas 24 horas (ou manter diário de 3-7 dias) é fundamental para entender sua nutrição real. Isso permite calcular se você está comendo calorias suficientes ou em excesso, se distribuição de proteínas, carboidratos e gorduras está adequada, quais vitaminas e minerais podem estar faltando, e se sua alimentação está causando inflamação. Também mostra padrões importantes: horários das refeições (que afeta seu metabolismo), períodos longos sem comer ou beliscar o tempo todo, quantidade de proteína por refeição (precisa ser 25-40g para construir músculo), fibras, variedade de alimentos (quanto mais variedade, melhor para bactérias intestinais). Programas de computador podem calcular exatamente seus níveis de magnésio, zinco, vitamina D, B12 e ômega-3. O recordatório também revela quando você come por emoção, horários de fissura por comida e hábitos problemáticos.',
  conduct = 'Obter recordatório 24h detalhado ou solicitar diário alimentar 3 dias (incluir 1 fim de semana). Registrar: horários, alimentos, quantidades (medidas caseiras ou peso), forma de preparo, bebidas, suplementos. Analisar manualmente ou via software (DietPro, Avanutri, MyFitnessPal): calorias totais, macronutrientes (%), distribuição proteica por refeição, fibras, micronutrientes críticos, índice glicêmico, carga inflamatória. Identificar: padrões de timing, adequação proteica, déficits de micronutrientes, ultraprocessados (%), variedade alimentar (número de alimentos diferentes/semana), hidratação. Comparar com necessidades individualizadas (TMB × fator atividade, necessidades específicas). Usar dados para orientação personalizada e priorização de intervenções.',
  updated_at = NOW()
WHERE id = '7eb5ab2a-5775-4f2a-acdc-e5866873d264';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('cfd1073e-d4a8-45c1-999a-b96e75223902', '7eb5ab2a-5775-4f2a-acdc-e5866873d264'),
  ('0108fb41-2fee-47f4-9d07-84a8de20c226', '7eb5ab2a-5775-4f2a-acdc-e5866873d264')
ON CONFLICT DO NOTHING;

-- Item 15: Situação familiar de alimentação
UPDATE score_items
SET
  clinical_relevance = 'O ambiente alimentar familiar é determinante crítico de escolhas individuais através de disponibilidade de alimentos, modelagem comportamental, normas sociais e suporte para mudanças. Incongruência entre objetivos individuais e padrão familiar representa barreira importante para adesão - paciente tentando emagrecimento em família que consome regularmente doces, refrigerantes e fast-food enfrenta exposição constante a gatilhos e falta de suporte social. Presença de crianças influencia compras (alimentos infantis ultraprocessados) e dinâmica de refeições. Cônjuge com padrão alimentar inadequado reduz chances de sucesso em mudanças dietéticas em 60%. Refeições em família (versus individualizadas) associam-se a melhor qualidade alimentar, menor risco de transtornos alimentares em adolescentes, melhor desempenho acadêmico. Abordar alimentação familiarmente (family-based approach) aumenta efetividade de intervenções nutricionais e promove mudanças sustentáveis para todos os membros.',
  patient_explanation = 'Sua família tem enorme influência na sua alimentação - o que tem em casa, como as pessoas comem, se apoiam ou sabotam suas tentativas de comer melhor. Se você está tentando melhorar alimentação mas sua família continua comprando e comendo doces, refrigerantes e fast-food, você enfrenta tentações constantes e falta de apoio. Estudos mostram que se seu cônjuge come mal, suas chances de sucesso em mudanças alimentares caem 60%. Presença de crianças influencia o que se compra (ultraprocessados "infantis") e como são as refeições. Fazer refeições juntos em família (não cada um comendo separado) melhora qualidade alimentar de todos, reduz risco de transtornos alimentares em adolescentes e até melhora desempenho escolar das crianças. Trabalhar alimentação em família (não só individualmente) aumenta muito as chances de sucesso e beneficia todos.',
  conduct = 'Investigar: composição familiar, padrão alimentar de cônjuge e filhos, alimentos disponíveis em casa, quem faz compras, presença de alimentos-gatilho, refeições compartilhadas versus individualizadas, suporte familiar para mudanças, conflitos relacionados a alimentação. Avaliar: se família representa facilitador ou barreira, discordância entre objetivos individuais e contexto familiar, impacto de alimentação infantil no contexto doméstico. Estratégias: envolver família nas orientações quando possível, negociar mudanças graduais no ambiente doméstico, criar estratégias para convivência com alimentos-gatilho, fortalecer suporte social. Considerar abordagem familiar integrada, especialmente em casos de obesidade infantil ou transtornos alimentares.',
  updated_at = NOW()
WHERE id = 'a97820ec-2c64-433f-b1ed-6a410cf4936d';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('cfd1073e-d4a8-45c1-999a-b96e75223902', 'a97820ec-2c64-433f-b1ed-6a410cf4936d'),
  ('b777f9d9-5a13-4856-9da8-548d1fd843f6', 'a97820ec-2c64-433f-b1ed-6a410cf4936d')
ON CONFLICT DO NOTHING;

-- Item 16: Todos seguem o mesmo padrão alimentar do paciente
UPDATE score_items
SET
  clinical_relevance = 'Alinhamento ou desalinhamento entre padrão alimentar individual e familiar determina viabilidade prática, adesão e sucesso terapêutico de intervenções nutricionais. Paciente que segue dieta restritiva (cetogênica, vegetariana, paleo) enquanto família mantém padrão ocidental onívoro enfrenta sobrecarga logística (preparar refeições separadas), custos aumentados, exposição constante a alimentos excluídos e risco de conflitos familiares. Situação oposta - família adotando padrão alimentar saudável enquanto paciente mantém alimentação inadequada - sugere resistência a mudanças, questões comportamentais/emocionais subjacentes ou falta de motivação. Alinhamento familiar facilita adesão, reduz custos, simplifica rotina e potencializa benefícios (todos melhoram marcadores de saúde). Desalinhamento requer estratégias específicas: meal prep individual, comunicação familiar assertiva, manejo de ambiente obesogênico doméstico.',
  patient_explanation = 'Se você come de forma diferente da sua família, isso afeta muito suas chances de sucesso. Se você segue dieta específica (cetogênica, vegetariana, sem glúten) mas sua família come normalmente, você enfrenta: trabalho extra preparando refeições separadas, custo maior, tentação constante vendo alimentos que você excluiu, possíveis conflitos familiares. Na situação oposta - família comendo saudável e você mantendo alimentação ruim - pode indicar que você está resistente a mudanças, tem questões emocionais com comida ou falta motivação. Quando toda família come igual (e saudável), tudo fica mais fácil: melhor adesão, menor custo, rotina mais simples, e todos melhoram saúde juntos. Se há desalinhamento, precisa de estratégias especiais: preparar suas refeições separadamente, comunicar necessidades claramente para família, criar formas de lidar com tentações em casa.',
  conduct = 'Investigar: grau de alinhamento alimentar familiar, quem prepara refeições (únicas versus individualizadas), impacto logístico/financeiro de desalinhamento, conflitos relacionados a diferenças alimentares, motivação para mudanças (individual versus familiar). Cenários: 1) Paciente restritivo/família padrão - avaliar sobrecarga, viabilidade, exposição a gatilhos; 2) Família saudável/paciente inadequado - investigar barreiras psicológicas, motivação, transtornos alimentares; 3) Alinhamento inadequado - oportunidade para intervenção familiar; 4) Alinhamento saudável - reforçar benefícios e facilitar manutenção. Orientar estratégias conforme contexto: meal prep, comunicação familiar, envolvimento de família em consultas, transição gradual familiar.',
  updated_at = NOW()
WHERE id = '668f752a-6b53-4fc2-9d5c-a7d275642ead';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('cfd1073e-d4a8-45c1-999a-b96e75223902', '668f752a-6b53-4fc2-9d5c-a7d275642ead'),
  ('0108fb41-2fee-47f4-9d07-84a8de20c226', '668f752a-6b53-4fc2-9d5c-a7d275642ead')
ON CONFLICT DO NOTHING;

-- Item 17: Suplementação utilizada e dose
UPDATE score_items
SET
  clinical_relevance = 'Avaliação detalhada de suplementação prévia e atual é fundamental para identificar uso inapropriado, doses inadequadas, interações medicamentosas, excesso vitamínico (hipervitaminoses A/D/E), deficiências persistentes apesar de suplementação (má absorção, dose insuficiente, forma inadequada) e gaps terapêuticos. Suplementação empírica sem avaliação laboratorial pode mascarar deficiências (ferro sem investigar anemia), causar toxicidade (vitamina A, selênio), ou ser inefetiva (vitamina D 400UI quando necessário 2000-5000UI). Formas de nutrientes importam: metilfolato versus ácido fólico em polimorfismos MTHFR, quelatos minerais versus óxidos (absorção), ômega-3 EPA:DHA ratio conforme objetivo. Timing de suplementação afeta absorção: lipossolúveis (D/E/K) com gorduras, minerais em jejum, magnésio à noite. Interações: cálcio compete com ferro/zinco/magnésio, vitaminas do complexo B podem excitar (tomar manhã), antiácidos reduzem absorção de B12/minerais.',
  patient_explanation = 'É importante saber exatamente quais suplementos você já usou e usa atualmente, em que doses, porque uso inadequado pode ser perigoso ou ineficaz. Tomar suplementos sem exames pode mascarar problemas, causar intoxicação (vitamina A, selênio) ou simplesmente não funcionar (vitamina D em dose muito baixa). A forma do nutriente importa muito: alguns tipos são muito melhor absorvidos que outros - metilfolato é melhor que ácido fólico para quem tem variação genética comum, minerais "quelados" são melhor absorvidos que óxidos baratos. O horário que você toma também afeta: vitaminas A/D/E/K precisam ser tomadas com gordura, minerais são melhor absorvidos em jejum, magnésio à noite ajuda sono. Alguns suplementos competem entre si: cálcio atrapalha absorção de ferro, zinco e magnésio. Vitaminas B podem dar energia (melhor de manhã), antiácidos prejudicam absorção de B12 e minerais.',
  conduct = 'Listar detalhadamente: 1) Suplementos atuais (nome comercial, princípio ativo, dose, forma química, horário, tempo de uso); 2) Histórico prévio (o que já usou, por quanto tempo, resultados percebidos, motivo de suspensão). Avaliar: adequação de doses (comparar com RDA e doses terapêuticas), formas químicas utilizadas, timing, interações entre suplementos, interações com medicamentos. Identificar: deficiências persistentes apesar de suplementação (investigar má absorção, SIBO, hipocloridria, doença celíaca), uso desnecessário, doses subterapêuticas, formas de baixa biodisponibilidade. Correlacionar com exames laboratoriais. Otimizar protocolo: ajustar doses, trocar formas, corrigir timing, adicionar cofatores (exemplo: magnésio + B6 + zinco para conversão D3).',
  updated_at = NOW()
WHERE id = 'ba93f0e8-9f64-4090-b843-99752e3c622d';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('0108fb41-2fee-47f4-9d07-84a8de20c226', 'ba93f0e8-9f64-4090-b843-99752e3c622d'),
  ('c375239d-02bb-46d1-a429-2abba4ef7999', 'ba93f0e8-9f64-4090-b843-99752e3c622d'),
  ('a9f50c86-cbec-4ca5-bec3-5422705fba3d', 'ba93f0e8-9f64-4090-b843-99752e3c622d'),
  ('f20437f0-a354-46bf-a79d-88e6b5c522d3', 'ba93f0e8-9f64-4090-b843-99752e3c622d')
ON CONFLICT DO NOTHING;

-- Item 18: Whey
UPDATE score_items
SET
  clinical_relevance = 'Whey protein é suplemento proteico derivado do soro do leite, amplamente utilizado para hipertrofia, recuperação muscular, perda ponderal e suporte nutricional. Vantagens: perfil aminoacídico completo (BCAA elevado, especialmente leucina - gatilho de síntese proteica), alta biodisponibilidade (PDCAAS 1.0), rápida absorção (ideal pós-treino), praticidade. Formas: concentrado (WPC - 70-80% proteína, contém lactose), isolado (WPI - >90% proteína, mínima lactose, melhor para intolerantes), hidrolisado (pré-digerido, absorção mais rápida, mais caro). Doses: 20-40g por porção otimizam síntese proteica muscular. Riscos: intolerância à lactose (WPC), alergia à proteína do leite, sobrecarga renal em insuficiência renal, acne (IGF-1 do leite), ganho ponderal se excesso calórico. Uso deve ser contextualizado: adequação proteica total da dieta, timing nutricional, objetivos, viabilidade financeira. Alternativas: proteína vegetal (ervilha, arroz), alimentos proteicos inteiros.',
  patient_explanation = 'Whey é proteína extraída do soro do leite, muito usada para ganhar músculo, recuperação após treino, emagrecer ou complementar proteína da dieta. Vantagens: tem todos os aminoácidos que o corpo precisa (especialmente leucina, que estimula crescimento muscular), é muito bem absorvido, age rápido (ótimo pós-treino) e é prático. Existem 3 tipos: concentrado (mais barato, tem lactose), isolado (mais puro, >90% proteína, quase sem lactose - melhor para intolerantes), hidrolisado (pré-digerido, mais rápido, mais caro). Dose ideal: 20-40g por vez. Cuidados: concentrado pode causar problemas se você tem intolerância à lactose, pode causar alergia se você tem alergia ao leite, pode sobrecarregar rins se eles já estão ruins, pode piorar acne, e pode engordar se você ultrapassar calorias. Precisa avaliar se você realmente precisa ou se consegue proteína suficiente da comida.',
  conduct = 'Investigar uso de whey: tipo (concentrado/isolado/hidrolisado), marca, dose/frequência, timing (pré/pós-treino, refeições), tempo de uso, objetivo (hipertrofia, emagrecimento, praticidade), custo, resultados percebidos. Avaliar: adequação proteica total da dieta (precisa 1.6-2.2g/kg para hipertrofia), necessidade real de suplementação versus alimentos, tolerância (sintomas gastrointestinais, acne). Contraindicações relativas: intolerância à lactose (preferir isolado), alergia à proteína do leite (substituir por vegetal), insuficiência renal (monitorar ureia/creatinina), acne severa (considerar suspensão - relação IGF-1). Orientar sobre: timing ideal (janela anabólica pós-treino), distribuição proteica ao longo do dia, alternativas alimentares proteicas, custo-benefício.',
  updated_at = NOW()
WHERE id = '88d88abd-3208-453a-b39e-0c9dd0c68816';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('700704d7-ba2a-4123-b677-b0f221b167c1', '88d88abd-3208-453a-b39e-0c9dd0c68816'),
  ('0108fb41-2fee-47f4-9d07-84a8de20c226', '88d88abd-3208-453a-b39e-0c9dd0c68816'),
  ('c375239d-02bb-46d1-a429-2abba4ef7999', '88d88abd-3208-453a-b39e-0c9dd0c68816')
ON CONFLICT DO NOTHING;

-- Item 19: Shots
UPDATE score_items
SET
  clinical_relevance = 'Shots nutricionais são concentrados líquidos de nutrientes, extratos vegetais, compostos bioativos ou probióticos, comercializados como suplementos práticos para energia, imunidade, detox ou antioxidação. Composições variam amplamente: gengibre + cúrcuma + limão (anti-inflamatório), chlorella + spirulina (detox), probióticos líquidos, colágeno hidrolisado, vitamina C em altas doses, extratos adaptogênicos. Efetividade depende de: dose de princípios ativos (muitos produtos têm doses subterapêuticas por custo), biodisponibilidade (cúrcuma requer piperina ou lipídios para absorção), qualidade de matéria-prima, ausência de açúcares/adoçantes adicionados. Marketing frequentemente superestima benefícios com claims não substanciados cientificamente (detox, alcalinização, imunidade instantânea). Shots podem ser úteis como parte de estratégia nutricional integrativa se bem formulados, mas não substituem alimentação adequada ou tratam causas-raiz de condições clínicas.',
  patient_explanation = 'Shots nutricionais são bebidas concentradas com nutrientes, extratos de plantas, probióticos ou antioxidantes, vendidos para dar energia, fortalecer imunidade ou "desintoxicar". Podem conter gengibre, cúrcuma, limão, chlorella, probióticos, colágeno, vitamina C, adaptógenos. Se funcionam depende de: ter dose suficiente dos ativos (muitos têm quantidade baixa demais por serem caros), serem bem absorvidos (cúrcuma precisa de pimenta preta ou gordura junto), terem qualidade, e não terem açúcar ou adoçantes ruins. O marketing desses produtos exagera benefícios sem provas científicas sólidas ("desintoxicação", "alcalinização", "imunidade instantânea"). Shots podem ajudar se bem formulados e como parte de estratégia nutricional completa, mas não substituem alimentação adequada nem tratam causas reais de problemas de saúde.',
  conduct = 'Investigar: tipos de shots utilizados, composição (ingredientes, doses), frequência, objetivo (energia, imunidade, detox, antioxidação), custo, resultados percebidos. Avaliar: presença de doses terapêuticas de ativos, qualidade de ingredientes (orgânicos, procedência), presença de açúcares/adoçantes, biodisponibilidade (cofatores necessários), justificativa científica para claims. Orientar: shots não substituem alimentação adequada, não fazem "detox" (fígado e rins fazem isso naturalmente), benefícios são geralmente modestos e requerem uso consistente. Se paciente insiste em usar, direcionar para: formulações evidence-based (gengibre + cúrcuma com piperina para inflamação, probióticos strain-specific), preparação caseira (mais econômico, controle de ingredientes), integração com estratégia nutricional global.',
  updated_at = NOW()
WHERE id = '61f8fa46-d68d-45c5-9af3-8bfe896c4d87';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('0108fb41-2fee-47f4-9d07-84a8de20c226', '61f8fa46-d68d-45c5-9af3-8bfe896c4d87'),
  ('c375239d-02bb-46d1-a429-2abba4ef7999', '61f8fa46-d68d-45c5-9af3-8bfe896c4d87')
ON CONFLICT DO NOTHING;

-- Item 20: Tipos de líquidos consumidos
UPDATE score_items
SET
  clinical_relevance = 'Hidratação adequada é essencial para homeostase celular, função renal, termorregulação, função cognitiva e performance física, mas tipo de líquido consumido afeta profundamente saúde metabólica e inflamatória. Água pura é gold standard. Bebidas açucaradas (refrigerantes, sucos industrializados, energéticos) são principais contribuidores de açúcares livres na dieta, associando-se a obesidade, resistência insulínica, diabetes tipo 2, doença hepática gordurosa, cáries e mortalidade cardiovascular. Um copo de suco de laranja (250ml) contém 25g açúcar - equivalente a refrigerante - sem fibras que modulam absorção. Adoçantes artificiais (aspartame, sucralose, sacarina) alteram microbiota, aumentam craving por doces e paradoxalmente associam-se a ganho ponderal e síndrome metabólica. Café e chá têm polifenóis benéficos mas excesso de cafeína causa ansiedade, insônia, taquicardia. Leite e bebidas lácteas: lactose, caseína A1 (inflamatória). Álcool: toxicidade hepática, disbiose, deficiências nutricionais.',
  patient_explanation = 'O que você bebe é tão importante quanto o que você come. Água pura é a melhor escolha. Refrigerantes, sucos de caixinha e energéticos são as maiores fontes de açúcar da alimentação moderna, causando obesidade, resistência à insulina, diabetes, gordura no fígado e problemas no coração. Um copo de suco de laranja tem tanto açúcar quanto refrigerante (25g em 250ml) e não tem as fibras da fruta que ajudam a controlar absorção. Adoçantes artificiais (aspartame, sucralose) parecem inocentes mas alteram bactérias intestinais, aumentam vontade de comer doce e paradoxalmente fazem ganhar peso. Café e chá têm antioxidantes bons mas excesso de cafeína causa ansiedade, insônia e coração acelerado. Leite tem lactose e proteína (caseína A1) que pode causar inflamação. Álcool prejudica fígado, desequilibra bactérias intestinais e esgota nutrientes.',
  conduct = 'Investigar: volume total de líquidos/dia, distribuição (água pura, sucos, refrigerantes, chás, café, leite, bebidas alcoólicas), frequência de cada tipo, horários (bebidas cafeinadas após 14h prejudicam sono), adição de açúcar/adoçantes. Quantificar: carga de açúcar líquido (g/dia), cafeína total (mg/dia), álcool (doses/semana). Avaliar impacto: bebidas açucaradas → glicemia, peso, HbA1c; cafeína → ansiedade, sono, taquicardia; álcool → enzimas hepáticas, triglicerídeos. Orientar: priorizar água pura (30-35ml/kg/dia), substituir sucos por frutas inteiras, eliminar refrigerantes, limitar café (até 400mg cafeína/dia, antes 14h), moderar álcool (máximo 1 dose/dia mulheres, 2 homens), evitar adoçantes artificiais. Estratégias de transição: águas saborizadas naturais (frutas, ervas), chás não-cafeinados.',
  updated_at = NOW()
WHERE id = 'ec39efac-fa2e-4b56-9156-03e9d991fc0e';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('cfd1073e-d4a8-45c1-999a-b96e75223902', 'ec39efac-fa2e-4b56-9156-03e9d991fc0e'),
  ('deb302ab-195b-41c0-8657-b299fd610a7f', 'ec39efac-fa2e-4b56-9156-03e9d991fc0e')
ON CONFLICT DO NOTHING;

-- Item 21: Quantidade total do dia
UPDATE score_items
SET
  clinical_relevance = 'Hidratação adequada é fundamental para múltiplas funções fisiológicas: homeostase celular, função renal (clearance de toxinas, regulação eletrolítica), termorregulação, função cognitiva (desidratação de 2% reduz performance cognitiva), lubrificação articular, motilidade intestinal e prevenção de litíase renal. Recomendação geral: 30-35ml/kg/dia, ajustado por atividade física, temperatura ambiente, altitude, condições clínicas. Desidratação crônica subclínica (<1-2%) é comum e associa-se a fadiga, cefaléia, constipação, redução de performance física/cognitiva, aumento de risco de litíase renal. Indicadores de hidratação: cor urinária (amarelo claro), volume urinário (>1-1.5L/dia), sede. Excesso hídrico (>3-4L/dia sem necessidade) pode causar hiponatremia dilucional, sobrecarga renal e depleção de eletrólitos. Necessidades aumentadas: exercício (500-1000ml/hora), calor, altitude, gestação/lactação, febr, diarréia.',
  patient_explanation = 'Beber água suficiente é essencial para seu corpo funcionar: células precisam de água, rins precisam para eliminar toxinas, cérebro precisa para funcionar bem (apenas 2% de desidratação já reduz sua capacidade mental), articulações precisam de lubrificação, intestino precisa para funcionar, e previne pedra nos rins. Recomendação geral: 30-35ml por quilo de peso por dia (pessoa de 70kg precisa de 2.1-2.5 litros). Desidratação leve crônica (que a maioria tem) causa cansaço, dor de cabeça, prisão de ventre, piora performance física e mental, e aumenta risco de pedra nos rins. Sinais de boa hidratação: urina amarelo claro, fazer xixi mais de 1-1.5 litros por dia. Beber água demais (>3-4L sem necessidade) pode diluir sódio no sangue, sobrecarregar rins e esgotar eletrólitos. Precisa mais água se: exercita, calor, altitude, grávida/amamentando, febre, diarréia.',
  conduct = 'Investigar: volume total estimado de líquidos/dia (somar todos os tipos), distribuição ao longo do dia, sede habitual, cor urinária, frequência urinária, presença de sinais de desidratação (fadiga, cefaléia, constipação, urina escura/concentrada), histórico de litíase renal. Calcular necessidade individual: peso × 30-35ml/kg/dia, ajustar por atividade física, clima, condições especiais. Comparar ingestão atual versus necessidade. Orientar sobre: volume adequado, distribuição regular (não concentrar), sinais de hidratação adequada, estratégias práticas (garrafa de água sempre à mão, lembretes, associar hidratação a horários de rotina). Condições especiais: atletas (protocolo de hidratação periexercício), gestantes (+300ml/dia), lactantes (+700ml/dia), litíase renal (>2.5L/dia para produzir >2L urina).',
  updated_at = NOW()
WHERE id = '9366ba67-4fce-4f31-8e44-3c38a1fab93a';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('cfd1073e-d4a8-45c1-999a-b96e75223902', '9366ba67-4fce-4f31-8e44-3c38a1fab93a'),
  ('0108fb41-2fee-47f4-9d07-84a8de20c226', '9366ba67-4fce-4f31-8e44-3c38a1fab93a')
ON CONFLICT DO NOTHING;

-- Item 22: Vegetariana
UPDATE score_items
SET
  clinical_relevance = 'Dieta vegetariana exclui carnes mas inclui ovos/laticínios (ovolactovegetariano) ou apenas vegetais (vegano - ver próximo item). Benefícios documentados: redução de risco cardiovascular (23% menor), diabetes tipo 2 (redução de 50%), hipertensão, alguns cânceres (colorretal), obesidade, mortalidade total. Mecanismos: maior ingestão de fibras, antioxidantes, fitoquímicos, menor ingestão de gorduras saturadas, colesterol, compostos pró-inflamatórios (AGEs, neu5Gc). Riscos nutricionais: deficiência de vitamina B12 (essencial suplementar), ferro (forma não-heme menos biodisponível - requer vitamina C), zinco, ômega-3 EPA/DHA (ALA vegetal tem conversão <5%), vitamina D, iodo, cálcio (se evita laticínios), proteína (possível mas requer planejamento). Qualidade importa: vegetariano ultraprocessado (junk food vegetariano) não confere benefícios. Crianças, gestantes, atletas requerem planejamento nutricional rigoroso.',
  patient_explanation = 'Dieta vegetariana não come carnes mas pode incluir ovos e leite. Benefícios comprovados: 23% menos risco de problemas cardíacos, 50% menos diabetes, menos pressão alta, alguns cânceres (intestino) e obesidade. Funciona porque vegetarianos comem mais fibras, antioxidantes e compostos benéficos das plantas, e menos gorduras saturadas, colesterol e compostos inflamatórios das carnes. Porém tem riscos nutricionais importantes: deficiência de B12 (essencial suplementar - não existe em vegetais), ferro (tipo vegetal é menos absorvido - precisa comer com vitamina C), zinco, ômega-3 EPA/DHA (vegetais têm ALA que converte <5% no corpo), vitamina D, iodo, cálcio (se não come laticínios), proteína (dá para obter mas precisa planejar bem). Importante: ser vegetariano e comer ultraprocessados vegetarianos não traz benefícios. Crianças, grávidas e atletas precisam planejamento nutricional rigoroso.',
  conduct = 'Investigar: tipo exato (ovolactovegetariano, lactovegetariano, ovovegetariano), tempo de adesão, motivação (ética, saúde, religiosa, ambiental), planejamento nutricional (autodidata versus acompanhado), suplementação atual, adequação proteica, variedad de fontes vegetais, consumo de ovos/laticínios. Avaliar riscos: solicitar B12 (essencial), ferritina (anemia ferropriva comum), zinco, vitamina D, índice ômega-3, hemograma completo, cálcio (se baixo consumo de laticínios). Investigar sintomas de deficiências: fadiga, palidez, glossite, parestesias (B12), queda capilar (ferro/zinco). Orientar: suplementação obrigatória de B12, fontes de ferro não-heme + vitamina C, combinação de proteínas vegetais para aminoácidos essenciais, considerar suplementação de ômega-3 (DHA algae), creatina (performance). Encaminhar para nutricionista se inadequação evidente.',
  updated_at = NOW()
WHERE id = 'f7d42bf8-7c6f-4362-8f98-36268261719b';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('cfd1073e-d4a8-45c1-999a-b96e75223902', 'f7d42bf8-7c6f-4362-8f98-36268261719b'),
  ('700704d7-ba2a-4123-b677-b0f221b167c1', 'f7d42bf8-7c6f-4362-8f98-36268261719b'),
  ('0108fb41-2fee-47f4-9d07-84a8de20c226', 'f7d42bf8-7c6f-4362-8f98-36268261719b')
ON CONFLICT DO NOTHING;

-- Item 23: Vegana
UPDATE score_items
SET
  clinical_relevance = 'Dieta vegana exclui todos produtos animais (carnes, ovos, laticínios, mel). Apresenta mesmos benefícios cardiovasculares/metabólicos da vegetariana, com riscos nutricionais amplificados pela exclusão total de produtos animais. Deficiências altamente prevalentes: vitamina B12 (100% requer suplementação - nenhuma fonte vegetal confiável), vitamina D (fontes vegetais limitadas a cogumelos UV-irradiados), EPA/DHA (conversão ínfima de ALA), ferro (biodisponibilidade reduzida), zinco, cálcio (biodisponibilidade vegetal reduzida por fitatos/oxalatos), iodo (sem sal iodado/algas), proteína (requer combinação consciente de leguminosas/grãos/oleaginosas). Creatina, carnosina, taurina são zero em dieta vegana - podem afetar performance física/cognitiva. Gestantes/lactantes veganas requerem monitoramento rigoroso - deficiência materna de B12 causa danos neurológicos irreversíveis em lactentes. Veganismo bem planejado com suplementação adequada pode ser saudável; mal planejado é alto risco nutricional.',
  patient_explanation = 'Dieta vegana exclui todos os produtos de origem animal - carnes, ovos, leite, mel. Tem os mesmos benefícios de saúde do coração e metabolismo que vegetariana, mas riscos nutricionais maiores porque exclui tudo de origem animal. Deficiências muito comuns: B12 (100% dos veganos precisam suplementar - não existe B12 confiável em vegetais), vitamina D (poucas fontes vegetais), ômega-3 EPA/DHA (conversão do tipo vegetal é mínima <5%), ferro (absorção reduzida), zinco, cálcio (absorção vegetal é menor por compostos que bloqueiam), iodo, proteína (consegue mas precisa combinar bem leguminosas, grãos, castanhas). Creatina, carnosina, taurina são zero em veganos - pode afetar músculos e cérebro. Grávidas e que amamentam veganas precisam monitoramento rigoroso - falta de B12 causa danos cerebrais permanentes no bebê. Veganismo bem planejado com suplementos certos pode ser saudável; mal planejado é muito arriscado.',
  conduct = 'Investigar: tempo de veganismo, motivação, planejamento nutricional (profissional versus autodidata), suplementação atual (essencial avaliar adequação), adequação proteica (combinar leguminosas/grãos), variedade de fontes vegetais, consumo de alimentos fortificados. Avaliação laboratorial obrigatória: B12 sérica (idealmente B12 ativa/holotranscobalamina ou homocisteína/ácido metilmalônico), ferritina, hemograma completo, vitamina D, índice ômega-3, zinco, cálcio sérico, TSH (iodo). Investigar sintomas: fadiga intensa, parestesias, glossite, alterações neurológicas (B12), anemia, queda capilar. Suplementação essencial: B12 (cianocobalamina 1000-2000mcg/dia ou metilcobalamina 2000mcg/dia ou 2500mcg/semana), DHA algae (200-300mg/dia), considerar multivitamínico vegano completo. Grupos de risco: gestantes, lactantes, crianças, adolescentes, atletas - acompanhamento nutricional obrigatório.',
  updated_at = NOW()
WHERE id = '73fa8a7e-3050-4cc5-b0c3-e6a5def6a3c3';

INSERT INTO article_score_items (article_id, score_item_id)
VALUES
  ('cfd1073e-d4a8-45c1-999a-b96e75223902', '73fa8a7e-3050-4cc5-b0c3-e6a5def6a3c3'),
  ('700704d7-ba2a-4123-b677-b0f221b167c1', '73fa8a7e-3050-4cc5-b0c3-e6a5def6a3c3'),
  ('0108fb41-2fee-47f4-9d07-84a8de20c226', '73fa8a7e-3050-4cc5-b0c3-e6a5def6a3c3')
ON CONFLICT DO NOTHING;

-- ============================================
-- FIM DO ENRIQUECIMENTO
-- 23 items processados com sucesso
-- ============================================
