-- ============================================================================
-- Script: Atualização Padrões Alimentares - 2026
-- Descrição: Reordena items, adiciona níveis e preenche campos clínicos
-- Data: 2026-02-12
-- ============================================================================

BEGIN;

-- ============================================================================
-- PARTE 1: REORDENAÇÃO DOS ITEMS
-- ============================================================================

-- Nova ordem: Livre primeiro, demais alfabética
UPDATE score_items SET "order" = 1 WHERE id = 'c77cedd3-2800-72e4-9d17-2f997d743550'; -- Livre
UPDATE score_items SET "order" = 2 WHERE id = '019c534c-a38b-7b52-b4e8-0888c8a17b85'; -- Baixa histamina
UPDATE score_items SET "order" = 3 WHERE id = 'c77cedd3-2800-711c-8e1f-b4526da1e7ca'; -- Carnívora
UPDATE score_items SET "order" = 4 WHERE id = 'c77cedd3-2800-7f39-89b7-5b249a004d08'; -- Cetogênica
UPDATE score_items SET "order" = 5 WHERE id = '019c534c-0bf6-7d6b-9c60-d650fa5f846f'; -- Hipercalórica
UPDATE score_items SET "order" = 6 WHERE id = 'c77cedd3-2800-70de-bec4-0b5534637c78'; -- Low carb
UPDATE score_items SET "order" = 7 WHERE id = '019c534b-a12b-727c-a251-693f5d84d467'; -- Low fodmap
UPDATE score_items SET "order" = 8 WHERE id = '019c534b-cdc0-7f24-993f-84d8a1f93fb2'; -- Mediterrâneo
UPDATE score_items SET "order" = 9 WHERE id = '019c5349-99e6-79f6-8e5d-14051a0929e0'; -- Sem gluten
UPDATE score_items SET "order" = 10 WHERE id = '019c534b-4ad4-7564-b8f3-e13286121e42'; -- Sem lactose
UPDATE score_items SET "order" = 11 WHERE id = '019bf31d-2ef0-7cc5-b0c3-e6a5def6a3c3'; -- Vegana
UPDATE score_items SET "order" = 12 WHERE id = '019bf31d-2ef0-7362-8f98-36268261719b'; -- Vegetariana

-- ============================================================================
-- PARTE 2: ATUALIZAÇÃO DOS CAMPOS CLÍNICOS DOS ITEMS
-- ============================================================================

-- Livre
UPDATE score_items SET
  clinical_relevance = 'Padrão alimentar sem restrições específicas, indicado para indivíduos metabolicamente saudáveis sem condições que exijam intervenção dietética. Serve como baseline comparativo em estudos clínicos e permite maior flexibilidade alimentar. Requer monitoramento para garantir adequação nutricional e prevenir desenvolvimento de comorbidades.',
  patient_explanation = 'Você pode consumir todos os grupos alimentares sem restrições específicas, desde que mantenha equilíbrio nutricional. Este padrão é adequado quando não há necessidades médicas especiais. É importante manter variedade alimentar e acompanhamento regular com profissional de saúde.',
  conduct = 'Indicar para pacientes sem comorbidades metabólicas, alergias ou intolerâncias documentadas. Realizar avaliação nutricional periódica para detectar deficiências ou excessos. Monitorar marcadores metabólicos (glicemia, lipidograma, inflamatórios) anualmente e orientar sobre alimentação equilibrada.',
  last_review = CURRENT_TIMESTAMP
WHERE id = 'c77cedd3-2800-72e4-9d17-2f997d743550';

-- Baixa histamina
UPDATE score_items SET
  clinical_relevance = 'Indicada para síndrome de ativação mastocitária (MCAS), intolerância à histamina por deficiência de DAO, e enxaqueca relacionada à histamina. Evidências de 2024 mostram melhora sintomática em 60-70% dos pacientes após 2-6 semanas de restrição. Importante distinguir MCAS (liberação de múltiplos mediadores) de intolerância à histamina (deficiência enzimática).',
  patient_explanation = 'Esta dieta reduz alimentos ricos em histamina (queijos maturados, embutidos, fermentados, álcool) para diminuir sintomas como urticária, enxaqueca, sintomas gastrointestinais e fadiga. A restrição geralmente é temporária (2-6 semanas) seguida de reintrodução gradual. Nem todos com sintomas precisam eliminar histamina completamente.',
  conduct = 'Prescrever por 2-6 semanas como diagnóstico terapêutico, monitorando intensidade e frequência dos sintomas. Considerar teste de desafio duplo-cego para confirmar intolerância. Avaliar deficiência de DAO e considerar suplementação enzimática. Encaminhar para alergologista se suspeita de MCAS (sintomas sistêmicos múltiplos).',
  last_review = CURRENT_TIMESTAMP
WHERE id = '019c534c-a38b-7b52-b4e8-0888c8a17b85';

-- Carnívora
UPDATE score_items SET
  clinical_relevance = 'Padrão alimentar baseado exclusivamente em produtos animais. Evidências de 2024-2025 mostram redução de peso e melhora de marcadores inflamatórios em curto prazo, mas riscos nutricionais significativos (deficiência de vitaminas C, D, cálcio, magnésio, fibras). Estudos mostram aumento de LDL-C e colesterol total, com respostas metabólicas divergentes entre indivíduos saudáveis e com disfunções metabólicas. Ausência completa de estudos de longo prazo.',
  patient_explanation = 'Dieta que inclui apenas alimentos de origem animal (carnes, peixes, ovos, laticínios). Pode trazer perda de peso e saciedade aumentada no curto prazo, mas apresenta riscos de deficiências nutricionais e elevação do colesterol. Não existem estudos sobre segurança a longo prazo (mais de 6 meses).',
  conduct = 'Não recomendar como primeira linha de tratamento. Se paciente optar pela dieta, limitar a 3-6 meses com monitoramento rigoroso: perfil lipídico completo, função renal, vitaminas C e D, cálcio, magnésio a cada 8-12 semanas. Contraindicada para pacientes com dislipidemia, doença cardiovascular ou renal. Suplementar vitaminas deficientes.',
  last_review = CURRENT_TIMESTAMP
WHERE id = 'c77cedd3-2800-711c-8e1f-b4526da1e7ca';

-- Cetogênica
UPDATE score_items SET
  clinical_relevance = 'Padrão very low-carb (5-10% carboidratos) com forte evidência para epilepsia refratária (redução >50% crises em ~50-60% dos pacientes) e diabetes tipo 2 (melhora de HbA1c e glicemia em 3-6 meses). Guidelines 2024 recomendam como terapia padrão em epilepsias genéticas (Dravet, Glut1, Angelman). Benefícios metabólicos atenuam após 6 meses, com aderência de 70% no primeiro ano caindo para 38% em 3 anos.',
  patient_explanation = 'Dieta muito restrita em carboidratos (pão, arroz, frutas) que faz o corpo usar gordura como energia, produzindo cetonas. Comprovadamente eficaz para epilepsia resistente a medicamentos e pode ajudar no controle de diabetes tipo 2 nos primeiros 6 meses. Requer acompanhamento médico próximo devido a efeitos colaterais (constipação, fadiga inicial, cálculos renais).',
  conduct = 'Indicar para epilepsia refratária (especialmente síndromes genéticas) sob supervisão de neurologista e nutricionista especializado. Para diabetes tipo 2, considerar por 3-6 meses com monitoramento de glicemia, cetonúria, função renal, lipidograma mensal. Ajustar medicações antidiabéticas para evitar hipoglicemia. Suplementar eletrólitos, vitaminas e fibras. Contraindicada em gestação, IRC, insuficiência hepática.',
  last_review = CURRENT_TIMESTAMP
WHERE id = 'c77cedd3-2800-7f39-89b7-5b249a004d08';

-- Hipercalórica
UPDATE score_items SET
  clinical_relevance = 'Indicada para desnutrição, sarcopenia, caquexia oncológica (afeta 15-40% no diagnóstico, 40-80% durante tratamento) e insuficiência cardíaca com perda de massa magra. Estudos 2024 com suplementos hipercalóricos hiperproteicos (whey, leucina, ômega-3) mostram ganho de massa magra, melhora funcional, qualidade de vida e parâmetros bioquímicos nutricionais em 24 semanas. Screening nutricional precoce é altamente recomendado em pacientes oncológicos e cardiológicos.',
  patient_explanation = 'Dieta com calorias aumentadas para ganho de peso e recuperação de massa muscular. Indicada quando há perda de peso não intencional, fraqueza muscular ou durante tratamento de câncer. Geralmente inclui suplementos nutricionais específicos com proteínas de qualidade (whey) e ácidos graxos essenciais.',
  conduct = 'Realizar triagem nutricional (Mini Nutritional Assessment, ASG-PPP) em pacientes oncológicos, cardíacos, idosos e pós-cirúrgicos. Calcular necessidades calórico-proteicas (1,2-2,0g proteína/kg/dia). Prescrever suplementos hipercalóricos hiperproteicos com leucina ou whey protein. Monitorar composição corporal (bioimpedância, DEXA), funcionalidade (força de preensão, teste de caminhada) e marcadores nutricionais (albumina, pré-albumina) a cada 4-6 semanas.',
  last_review = CURRENT_TIMESTAMP
WHERE id = '019c534c-0bf6-7d6b-9c60-d650fa5f846f';

-- Low carb
UPDATE score_items SET
  clinical_relevance = 'Redução moderada de carboidratos (<130g/dia ou <26% do VET) com evidência robusta para síndrome metabólica e diabetes tipo 2. Meta-análises 2024 mostram maior eficácia em 3 meses para redução de HbA1c, glicemia de jejum, triglicerídeos e peso, com benefícios atenuados após 6 meses. Revisões sistemáticas indicam benefícios modestos de curto prazo, mas inconsistência entre populações. Recomendado clinicamente por até 6 meses.',
  patient_explanation = 'Dieta que reduz carboidratos (pães, massas, arroz, açúcares) moderadamente, mantendo frutas, vegetais e grãos integrais em menor quantidade. Eficaz para melhorar diabetes tipo 2, triglicerídeos e perder peso nos primeiros 3-6 meses. Menos restritiva que a dieta cetogênica, permitindo maior variedade de alimentos.',
  conduct = 'Indicar para diabetes tipo 2, pré-diabetes, síndrome metabólica e hipertrigliceridemia como opção de primeira linha por 3-6 meses. Monitorar HbA1c, glicemia, triglicerídeos e peso mensalmente. Ajustar medicações antidiabéticas conforme necessário. Após 6 meses, reavaliar aderência e resultados; considerar manutenção se benefícios sustentados. Garantir adequação de fibras, vitaminas e minerais.',
  last_review = CURRENT_TIMESTAMP
WHERE id = 'c77cedd3-2800-70de-bec4-0b5534637c78';

-- Low fodmap
UPDATE score_items SET
  clinical_relevance = 'Padrão alimentar com forte evidência para síndrome do intestino irritável (SII), sendo a intervenção dietética mais recomendada clinicamente. Estudos 2024 mostram 60% de resposta vs. 26% em dieta controle, com melhora significativa em dor abdominal, distensão, forma das fezes e qualidade de vida em 4 semanas. Fase de restrição tem impacto negativo em Bifidobactérias e qualidade nutricional, necessitando reintrodução e personalização. Integração com dieta mediterrânea está sendo estudada para otimizar microbiota e saúde cardiometabólica.',
  patient_explanation = 'Dieta que reduz temporariamente carboidratos fermentáveis (FODMAPs) presentes em trigo, cebola, alho, leguminosas, alguns laticínios e frutas que causam gases, dor e diarreia em pessoas sensíveis. Tem três fases: restrição (2-6 semanas), reintrodução gradual e personalização. Não é para ser seguida permanentemente.',
  conduct = 'Prescrever para SII confirmada (critérios de Roma IV) por 2-6 semanas sob orientação de nutricionista especializado. Avaliar resposta sintomática com diário alimentar e escalas validadas (IBS-SSS). Iniciar reintrodução gradual por grupos de FODMAPs para identificar gatilhos individuais. Monitorar adequação nutricional e microbiota intestinal (se disponível). Considerar integração com padrão mediterrâneo para benefícios adicionais.',
  last_review = CURRENT_TIMESTAMP
WHERE id = '019c534b-a12b-727c-a251-693f5d84d467';

-- Mediterrâneo
UPDATE score_items SET
  clinical_relevance = 'Padrão com nível mais alto de evidência para redução de risco cardiovascular, AVC, demência e depressão. Estudos 2024-2025 mostram benefícios anti-inflamatórios, cardiometabólicos e sobre microbiota intestinal. Diretrizes americanas 2025-2030 recomendam fortemente para saúde cardiovascular. Integração com low FODMAP está sendo proposta para otimizar controle de SII com adequação nutricional. Caracterizado por alto consumo de vegetais, frutas, grãos integrais, leguminosas, azeite, peixes, consumo moderado de laticínios e baixo de carnes vermelhas.',
  patient_explanation = 'Dieta baseada nos hábitos alimentares de países mediterrâneos: muitos vegetais, frutas, azeite de oliva, peixes, grãos integrais, oleaginosas e pouco consumo de carne vermelha e alimentos processados. Comprovadamente reduz riscos de doenças cardíacas, derrame, diabetes e declínio cognitivo. É um padrão sustentável e saboroso a longo prazo.',
  conduct = 'Recomendar como primeira linha para prevenção cardiovascular, síndrome metabólica, diabetes tipo 2, hipertensão e pacientes com risco de declínio cognitivo. Prescrever com orientações práticas: azeite como gordura principal, peixes 2-3x/semana, leguminosas diárias, limitação de carnes vermelhas. Monitorar aderência com questionários validados (MEDAS), perfil lipídico, glicemia e pressão arterial a cada 3-6 meses. Combinar com low FODMAP se SII coexistir.',
  last_review = CURRENT_TIMESTAMP
WHERE id = '019c534b-cdc0-7f24-993f-84d8a1f93fb2';

-- Sem gluten
UPDATE score_items SET
  clinical_relevance = 'Tratamento obrigatório e único aceito para doença celíaca (DC) confirmada. Estudo Lancet 2025 revelou que sensibilidade ao glúten não celíaca (SGNC) é frequentemente causada por FODMAPs ou outros componentes do trigo, não pelo glúten. Prevalência autorrelatada é 10-36% globalmente, mas apenas <30% confirmam diagnóstico em testes duplo-cego. Eliminação completa de glúten não é necessária na maioria dos casos de SGNC. Para DC, dieta sem glúten é vitalícia e previne dano intestinal e complicações.',
  patient_explanation = 'Dieta que elimina trigo, centeio e cevada (fontes de glúten). Obrigatória para quem tem doença celíaca confirmada por exames médicos. Pesquisas recentes (2025) mostram que a maioria das pessoas que se acham sensíveis ao glúten realmente reagem a outros componentes do trigo (FODMAPs). Se você não tem doença celíaca, eliminar glúten sem orientação médica pode ser desnecessário.',
  conduct = 'Exigir diagnóstico confirmatório de DC (sorologia + biópsia) ANTES de prescrever dieta sem glúten, pois retirada prévia invalida testes. Para DC confirmada, prescrever eliminação estrita vitalícia, educar sobre contaminação cruzada e monitorar aderência com sorologia e sintomas. Para suspeita de SGNC, realizar desafio duplo-cego ou considerar low FODMAP primeiro. Suplementar ferro, cálcio, vitaminas B se deficiente. Reavaliar anualmente.',
  last_review = CURRENT_TIMESTAMP
WHERE id = '019c5349-99e6-79f6-8e5d-14051a0929e0';

-- Sem lactose
UPDATE score_items SET
  clinical_relevance = 'Indicada para intolerância à lactose (deficiência de lactase), secundária a doenças intestinais ou pós-gastrectomia. Maioria dos pacientes tolera até 12-15g lactose/dia (~250ml leite) especialmente se consumida com outros alimentos. Exclusão completa raramente necessária. Doses únicas toleradas: até 5g (~100ml leite). Queijos duros (suíço, cheddar) têm lactose mínima e não causam sintomas. Exclusão total de laticínios aumenta risco de osteopenia e osteoporose por deficiência de cálcio.',
  patient_explanation = 'Dieta que reduz ou elimina lactose (açúcar do leite) presente em leite, iogurte e alguns queijos. A maioria das pessoas com intolerância tolera pequenas quantidades (até 1 copo de leite/dia), especialmente se consumido junto com refeições. Queijos duros geralmente não causam sintomas. Eliminar laticínios completamente pode enfraquecer os ossos por falta de cálcio.',
  conduct = 'Confirmar intolerância com teste de hidrogênio expirado ou teste de tolerância. Educar que restrição completa raramente é necessária: orientar consumo fracionado (até 250ml leite/dia dividido em porções menores com alimentos). Recomendar queijos duros, iogurtes fermentados e produtos com lactase. Prescrever suplementação de lactase se necessário. Garantir ingestão adequada de cálcio (1000-1200mg/dia) via laticínios com lactose reduzida ou suplementos. Monitorar densidade óssea em exclusão prolongada.',
  last_review = CURRENT_TIMESTAMP
WHERE id = '019c534b-4ad4-7564-b8f3-e13286121e42';

-- Vegana
UPDATE score_items SET
  clinical_relevance = 'Padrão plant-based com evidência robusta para redução de risco cardiovascular, diabetes tipo 2, hipertensão, certos cânceres e obesidade. Position Paper da Academia de Nutrição (janeiro 2025, válido até 2032) confirma adequação nutricional quando bem planejada. Micronutrientes críticos: vitamina B12 (suplementação obrigatória), ferro, iodo, cálcio, vitamina D e colina. Benefícios cardiometabólicos bem estabelecidos, mas requer orientação profissional para evitar deficiências.',
  patient_explanation = 'Dieta que exclui todos os produtos animais (carnes, peixes, ovos, laticínios, mel). Quando bem planejada, reduz risco de doenças cardíacas, diabetes e alguns tipos de câncer. É fundamental suplementar vitamina B12 obrigatoriamente e garantir fontes adequadas de ferro, cálcio, iodo e ômega-3. Acompanhamento com nutricionista é essencial.',
  conduct = 'Avaliar motivações e fornecer orientação baseada em evidências (Position Paper 2025). Prescrever suplementação obrigatória de B12 (cianocobalamina 1000mcg/semana ou metilcobalamina diária). Monitorar B12 sérica, ferritina, vitamina D, cálcio, iodo, zinco e ácidos graxos ômega-3 a cada 6-12 meses. Educar sobre proteínas completas (combinação de leguminosas com cereais), ferro não-heme com vitamina C, e fontes de DHA/EPA (algas). Atenção especial em gestação, lactação, infância.',
  last_review = CURRENT_TIMESTAMP
WHERE id = '019bf31d-2ef0-7cc5-b0c3-e6a5def6a3c3';

-- Vegetariana
UPDATE score_items SET
  clinical_relevance = 'Padrão plant-based que permite ovos e laticínios, com mesmo nível de evidência que dieta vegana para redução de risco cardiovascular, diabetes tipo 2, hipertensão e obesidade. Academia de Nutrição (2025) confirma adequação nutricional em todas as fases da vida quando bem planejada. Diretrizes Americanas 2025-2030 fornecem recomendações específicas. Micronutrientes de atenção: vitamina B12, ferro, iodo, vitamina D e colina. Menor risco de deficiências que dieta vegana devido à inclusão de ovos e laticínios.',
  patient_explanation = 'Dieta que exclui carnes e peixes, mas permite ovos, leite e derivados. Comprovadamente reduz risco de doenças cardíacas, diabetes e pressão alta. Ovos e laticínios fornecem vitamina B12, cálcio e proteínas de qualidade, facilitando adequação nutricional. Mesmo assim, acompanhamento nutricional é recomendado para garantir todos os nutrientes.',
  conduct = 'Avaliar tipo de vegetarianismo (ovolacto, lacto, ovo) e fornecer orientação individualizada conforme Diretrizes 2025-2030. Monitorar B12 (especialmente se baixo consumo de ovos/laticínios), ferro, vitamina D, iodo e zinco a cada 12 meses. Educar sobre fontes de proteína completa, ferro com vitamina C, e adequação de cálcio via laticínios ou fortificados. Considerar suplementação de B12 se ingestão inadequada (<2,4mcg/dia). Atenção especial para gestantes, lactantes e crianças.',
  last_review = CURRENT_TIMESTAMP
WHERE id = '019bf31d-2ef0-7362-8f98-36268261719b';

-- ============================================================================
-- PARTE 3: CRIAÇÃO DOS NÍVEIS (0 - Não, 5 - Sim)
-- ============================================================================

-- Livre
INSERT INTO score_levels (id, item_id, level, name, operator, lower_limit, upper_limit, clinical_relevance, patient_explanation, conduct, last_review)
VALUES
  (uuid_generate_v7(), 'c77cedd3-2800-72e4-9d17-2f997d743550', 0, 'Não', '=', '0', '0',
   'Indica que o paciente segue algum padrão alimentar específico com restrições.',
   'Você está seguindo uma dieta específica com alguma restrição alimentar.',
   'Investigar qual padrão alimentar está sendo seguido e avaliar adequação nutricional e necessidade de acompanhamento profissional.',
   CURRENT_TIMESTAMP),
  (uuid_generate_v7(), 'c77cedd3-2800-72e4-9d17-2f997d743550', 5, 'Sim', '=', '5', '5',
   'Paciente segue alimentação livre sem restrições específicas, permitindo maior flexibilidade alimentar. Importante garantir que mesmo sem restrições, a alimentação seja equilibrada e variada.',
   'Você não segue nenhuma dieta restritiva específica, pode consumir todos os grupos alimentares com equilíbrio.',
   'Realizar orientação sobre alimentação equilibrada mesmo na ausência de restrições. Monitorar marcadores metabólicos anualmente.',
   CURRENT_TIMESTAMP);

-- Baixa histamina
INSERT INTO score_levels (id, item_id, level, name, operator, lower_limit, upper_limit, clinical_relevance, patient_explanation, conduct, last_review)
VALUES
  (uuid_generate_v7(), '019c534c-a38b-7b52-b4e8-0888c8a17b85', 0, 'Não', '=', '0', '0',
   'Paciente não segue dieta de baixa histamina. Se houver sintomas compatíveis com intolerância à histamina, considerar avaliação.',
   'Você não está restringindo alimentos ricos em histamina na sua dieta.',
   'Se paciente apresentar sintomas de intolerância à histamina (urticária, enxaqueca, sintomas GI), considerar teste diagnóstico.',
   CURRENT_TIMESTAMP),
  (uuid_generate_v7(), '019c534c-a38b-7b52-b4e8-0888c8a17b85', 5, 'Sim', '=', '5', '5',
   'Paciente em dieta de baixa histamina. Monitorar aderência, resposta sintomática e adequação nutricional, especialmente se a restrição for prolongada.',
   'Você está evitando alimentos ricos em histamina para controlar sintomas como urticária, dor de cabeça ou problemas digestivos.',
   'Acompanhar por 2-6 semanas, avaliar melhora sintomática. Planejar reintrodução gradual para identificar alimentos gatilhos específicos.',
   CURRENT_TIMESTAMP);

-- Carnívora
INSERT INTO score_levels (id, item_id, level, name, operator, lower_limit, upper_limit, clinical_relevance, patient_explanation, conduct, last_review)
VALUES
  (uuid_generate_v7(), 'c77cedd3-2800-711c-8e1f-b4526da1e7ca', 0, 'Não', '=', '0', '0',
   'Paciente não segue dieta carnívora, consumindo variedade de grupos alimentares.',
   'Você não está seguindo uma dieta baseada apenas em produtos animais.',
   'Se paciente relatar interesse em dieta carnívora, educar sobre riscos nutricionais e ausência de evidências de longo prazo.',
   CURRENT_TIMESTAMP),
  (uuid_generate_v7(), 'c77cedd3-2800-711c-8e1f-b4526da1e7ca', 5, 'Sim', '=', '5', '5',
   'Paciente em dieta carnívora exclusiva. ALERTA: Alto risco de deficiências nutricionais (vitaminas C, D, cálcio, magnésio, fibras) e elevação de LDL-C. Monitoramento rigoroso obrigatório.',
   'Você está consumindo apenas alimentos de origem animal. Esta dieta pode causar deficiências nutricionais e elevar colesterol - requer acompanhamento médico frequente.',
   'Monitorar perfil lipídico, função renal, vitaminas C e D, cálcio, magnésio a cada 8-12 semanas. Limitar a 3-6 meses. Contraindicar se dislipidemia ou doença cardiovascular. Suplementar deficiências.',
   CURRENT_TIMESTAMP);

-- Cetogênica
INSERT INTO score_levels (id, item_id, level, name, operator, lower_limit, upper_limit, clinical_relevance, patient_explanation, conduct, last_review)
VALUES
  (uuid_generate_v7(), 'c77cedd3-2800-7f39-89b7-5b249a004d08', 0, 'Não', '=', '0', '0',
   'Paciente não segue dieta cetogênica. Se houver indicação clínica específica (epilepsia refratária, diabetes tipo 2), considerar como opção terapêutica.',
   'Você não está seguindo uma dieta cetogênica (muito restrita em carboidratos).',
   'Avaliar se há indicações clínicas para dieta cetogênica e educar sobre benefícios e efeitos colaterais caso paciente demonstre interesse.',
   CURRENT_TIMESTAMP),
  (uuid_generate_v7(), 'c77cedd3-2800-7f39-89b7-5b249a004d08', 5, 'Sim', '=', '5', '5',
   'Paciente em dieta cetogênica. Indicação estabelecida para epilepsia refratária e evidência para diabetes tipo 2. Monitorar cetonúria, função renal, lipidograma e efeitos adversos (constipação, cálculos renais).',
   'Você está em dieta cetogênica, muito restrita em carboidratos. Seu corpo produz cetonas como fonte de energia. Requer acompanhamento médico próximo.',
   'Monitorar glicemia, cetonúria, função renal e lipidograma mensalmente. Ajustar medicações antidiabéticas. Suplementar eletrólitos e fibras. Atenção a sintomas de cetoacidose.',
   CURRENT_TIMESTAMP);

-- Hipercalórica
INSERT INTO score_levels (id, item_id, level, name, operator, lower_limit, upper_limit, clinical_relevance, patient_explanation, conduct, last_review)
VALUES
  (uuid_generate_v7(), '019c534c-0bf6-7d6b-9c60-d650fa5f846f', 0, 'Não', '=', '0', '0',
   'Paciente não necessita de dieta hipercalórica no momento. Se houver perda de peso não intencional ou sarcopenia, reavaliar necessidade.',
   'Você não precisa aumentar calorias na dieta atual.',
   'Monitorar peso, composição corporal e funcionalidade. Se houver perda de peso >5% em 6 meses ou sarcopenia, considerar dieta hipercalórica.',
   CURRENT_TIMESTAMP),
  (uuid_generate_v7(), '019c534c-0bf6-7d6b-9c60-d650fa5f846f', 5, 'Sim', '=', '5', '5',
   'Paciente em dieta hipercalórica hiperproteica. Indicada para desnutrição, sarcopenia ou caquexia. Monitorar ganho de massa magra, funcionalidade e marcadores nutricionais (albumina, pré-albumina).',
   'Você está consumindo mais calorias e proteínas para ganhar peso e recuperar massa muscular. Importante para sua recuperação.',
   'Calcular necessidades (1,2-2,0g proteína/kg/dia). Prescrever suplementos se necessário. Monitorar composição corporal e funcionalidade a cada 4-6 semanas.',
   CURRENT_TIMESTAMP);

-- Low carb
INSERT INTO score_levels (id, item_id, level, name, operator, lower_limit, upper_limit, clinical_relevance, patient_explanation, conduct, last_review)
VALUES
  (uuid_generate_v7(), 'c77cedd3-2800-70de-bec4-0b5534637c78', 0, 'Não', '=', '0', '0',
   'Paciente não restringe carboidratos. Se houver diabetes tipo 2, pré-diabetes ou síndrome metabólica, considerar dieta low carb como opção terapêutica.',
   'Você não está restringindo carboidratos na sua dieta atual.',
   'Avaliar marcadores metabólicos. Se indicação para controle glicêmico ou perda de peso, considerar dieta low carb por 3-6 meses.',
   CURRENT_TIMESTAMP),
  (uuid_generate_v7(), 'c77cedd3-2800-70de-bec4-0b5534637c78', 5, 'Sim', '=', '5', '5',
   'Paciente em dieta low carb. Evidência para redução de HbA1c, glicemia, triglicerídeos e peso em 3-6 meses. Benefícios mais pronunciados em curto prazo.',
   'Você está reduzindo carboidratos na dieta para melhorar diabetes ou perder peso. Eficaz nos primeiros 3-6 meses.',
   'Monitorar HbA1c, glicemia, triglicerídeos e peso mensalmente. Ajustar medicações antidiabéticas. Reavaliar após 6 meses e considerar manutenção se benefícios sustentados.',
   CURRENT_TIMESTAMP);

-- Low fodmap
INSERT INTO score_levels (id, item_id, level, name, operator, lower_limit, upper_limit, clinical_relevance, patient_explanation, conduct, last_review)
VALUES
  (uuid_generate_v7(), '019c534b-a12b-727c-a251-693f5d84d467', 0, 'Não', '=', '0', '0',
   'Paciente não segue dieta low FODMAP. Se houver sintomas de síndrome do intestino irritável (SII), considerar como primeira linha de intervenção dietética.',
   'Você não está restringindo FODMAPs (carboidratos fermentáveis) na sua dieta.',
   'Se sintomas de SII (dor abdominal, distensão, alteração do hábito intestinal), considerar teste diagnóstico e dieta low FODMAP.',
   CURRENT_TIMESTAMP),
  (uuid_generate_v7(), '019c534b-a12b-727c-a251-693f5d84d467', 5, 'Sim', '=', '5', '5',
   'Paciente em dieta low FODMAP. Intervenção dietética mais eficaz para SII (60% resposta em 4 semanas). Importante fase de reintrodução para evitar restrição desnecessária e impacto negativo em microbiota.',
   'Você está reduzindo FODMAPs para controlar sintomas intestinais. Esta dieta tem 3 fases e não deve ser permanente.',
   'Acompanhar por 2-6 semanas na fase de restrição. Avaliar resposta com diário e escalas validadas. Iniciar reintrodução gradual por grupos para personalizar dieta.',
   CURRENT_TIMESTAMP);

-- Mediterrâneo
INSERT INTO score_levels (id, item_id, level, name, operator, lower_limit, upper_limit, clinical_relevance, patient_explanation, conduct, last_review)
VALUES
  (uuid_generate_v7(), '019c534b-cdc0-7f24-993f-84d8a1f93fb2', 0, 'Não', '=', '0', '0',
   'Paciente não segue dieta mediterrânea. Fortemente recomendada para prevenção cardiovascular, diabetes e declínio cognitivo. Considerar como primeira linha em múltiplas condições.',
   'Você não está seguindo o padrão alimentar mediterrâneo.',
   'Educar sobre benefícios cardiovasculares e metabólicos da dieta mediterrânea. Recomendar como primeira linha para prevenção e tratamento de múltiplas condições.',
   CURRENT_TIMESTAMP),
  (uuid_generate_v7(), '019c534b-cdc0-7f24-993f-84d8a1f93fb2', 5, 'Sim', '=', '5', '5',
   'Paciente segue dieta mediterrânea. Nível mais alto de evidência para redução de risco cardiovascular, AVC, demência e depressão. Padrão sustentável e com benefícios amplamente comprovados.',
   'Você segue a dieta mediterrânea, comprovadamente benéfica para coração, cérebro e metabolismo. Continue!',
   'Reforçar aderência com orientações práticas. Monitorar com questionário MEDAS, perfil lipídico, glicemia e PA a cada 3-6 meses. Elogiar escolha alimentar baseada em evidências.',
   CURRENT_TIMESTAMP);

-- Sem gluten
INSERT INTO score_levels (id, item_id, level, name, operator, lower_limit, upper_limit, clinical_relevance, patient_explanation, conduct, last_review)
VALUES
  (uuid_generate_v7(), '019c5349-99e6-79f6-8e5d-14051a0929e0', 0, 'Não', '=', '0', '0',
   'Paciente consome glúten. Se houver sintomas sugestivos de doença celíaca, realizar investigação diagnóstica ANTES de eliminar glúten.',
   'Você consome alimentos com glúten (trigo, centeio, cevada) normalmente.',
   'Se sintomas sugestivos de DC (diarreia crônica, distensão, anemia, perda de peso), solicitar sorologia e biópsia ANTES de retirar glúten.',
   CURRENT_TIMESTAMP),
  (uuid_generate_v7(), '019c5349-99e6-79f6-8e5d-14051a0929e0', 5, 'Sim', '=', '5', '5',
   'Paciente em dieta sem glúten. Obrigatória em doença celíaca confirmada. Para SGNC, evidências recentes (Lancet 2025) mostram que maioria reage a FODMAPs, não glúten. Confirmar diagnóstico correto.',
   'Você está eliminando glúten da dieta. Se tem doença celíaca confirmada, mantenha sempre. Se não, discuta com médico - pode ser desnecessário.',
   'Confirmar diagnóstico de DC (sorologia + biópsia). Se DC, manter eliminação vitalícia e educar sobre contaminação cruzada. Se SGNC, considerar teste de desafio ou dieta low FODMAP.',
   CURRENT_TIMESTAMP);

-- Sem lactose
INSERT INTO score_levels (id, item_id, level, name, operator, lower_limit, upper_limit, clinical_relevance, patient_explanation, conduct, last_review)
VALUES
  (uuid_generate_v7(), '019c534b-4ad4-7564-b8f3-e13286121e42', 0, 'Não', '=', '0', '0',
   'Paciente consome lactose sem restrições. Se houver sintomas de intolerância (gases, diarreia, cólicas após lácteos), considerar avaliação.',
   'Você consome leite e derivados normalmente sem restrição de lactose.',
   'Se sintomas após consumo de lácteos, considerar teste de intolerância à lactose. A maioria tolera pequenas quantidades fracionadas.',
   CURRENT_TIMESTAMP),
  (uuid_generate_v7(), '019c534b-4ad4-7564-b8f3-e13286121e42', 5, 'Sim', '=', '5', '5',
   'Paciente em dieta sem lactose. Restrição completa raramente necessária - maioria tolera até 12-15g/dia fracionados. Atenção à adequação de cálcio (1000-1200mg/dia) para prevenir osteopenia/osteoporose.',
   'Você está evitando lactose. A maioria das pessoas tolera pequenas quantidades. Importante garantir cálcio suficiente para os ossos.',
   'Educar sobre tolerância fracionada. Recomendar queijos duros, iogurtes fermentados e produtos com lactase. Garantir 1000-1200mg cálcio/dia via laticínios com lactose reduzida ou suplementos.',
   CURRENT_TIMESTAMP);

-- Vegana
INSERT INTO score_levels (id, item_id, level, name, operator, lower_limit, upper_limit, clinical_relevance, patient_explanation, conduct, last_review)
VALUES
  (uuid_generate_v7(), '019bf31d-2ef0-7cc5-b0c3-e6a5def6a3c3', 0, 'Não', '=', '0', '0',
   'Paciente consome produtos animais. Se houver interesse em transição para dieta plant-based, fornecer orientação baseada em evidências (Position Paper 2025).',
   'Você consome produtos animais na sua dieta.',
   'Se paciente demonstrar interesse em veganismo, educar sobre planejamento nutricional adequado e suplementação obrigatória de B12.',
   CURRENT_TIMESTAMP),
  (uuid_generate_v7(), '019bf31d-2ef0-7cc5-b0c3-e6a5def6a3c3', 5, 'Sim', '=', '5', '5',
   'Paciente em dieta vegana. Evidência robusta para benefícios cardiometabólicos. OBRIGATÓRIO: suplementação de B12. Monitorar ferro, iodo, cálcio, vitamina D, zinco e ômega-3 regularmente.',
   'Você não consome produtos animais. Importante suplementar B12 obrigatoriamente e garantir todos os nutrientes. Acompanhamento nutricional é essencial.',
   'Prescrever B12 obrigatoriamente (1000mcg/semana). Monitorar B12, ferritina, vitamina D, cálcio, iodo, zinco e ômega-3 a cada 6-12 meses. Educar sobre proteínas completas e ferro com vitamina C.',
   CURRENT_TIMESTAMP);

-- Vegetariana
INSERT INTO score_levels (id, item_id, level, name, operator, lower_limit, upper_limit, clinical_relevance, patient_explanation, conduct, last_review)
VALUES
  (uuid_generate_v7(), '019bf31d-2ef0-7362-8f98-36268261719b', 0, 'Não', '=', '0', '0',
   'Paciente consome carnes. Se houver interesse em alimentação plant-based, fornecer orientação sobre opções vegetarianas e veganas.',
   'Você consome carnes e peixes na sua dieta.',
   'Se interesse em vegetarianismo, educar sobre adequação nutricional e diferenças entre padrões (ovolacto, lacto, ovo, vegano).',
   CURRENT_TIMESTAMP),
  (uuid_generate_v7(), '019bf31d-2ef0-7362-8f98-36268261719b', 5, 'Sim', '=', '5', '5',
   'Paciente vegetariano (permite ovos/laticínios). Evidência para benefícios cardiovasculares e metabólicos. Menor risco de deficiências que dieta vegana. Monitorar B12, ferro, vitamina D, iodo e zinco.',
   'Você não consome carnes nem peixes, mas inclui ovos e/ou laticínios. Padrão saudável quando bem planejado.',
   'Identificar tipo (ovolacto, lacto, ovo). Monitorar B12, ferro, vitamina D, iodo e zinco anualmente. Educar sobre proteínas completas e ferro com vitamina C. Suplementar B12 se inadequado.',
   CURRENT_TIMESTAMP);

-- ============================================================================
-- VERIFICAÇÃO FINAL
-- ============================================================================

-- Confirmar reordenação
SELECT id, name, "order"
FROM score_items
WHERE parent_item_id = '019c534a-afc3-70c4-82e0-bfde4b5b8f93'
AND deleted_at IS NULL
ORDER BY "order";

-- Confirmar níveis criados (devem ser 24 níveis = 12 items × 2 níveis)
SELECT si.name as item_name, sl.level, sl.name as level_name
FROM score_items si
JOIN score_levels sl ON si.id = sl.item_id
WHERE si.parent_item_id = '019c534a-afc3-70c4-82e0-bfde4b5b8f93'
AND si.deleted_at IS NULL
AND sl.deleted_at IS NULL
ORDER BY si."order", sl.level;

COMMIT;

-- ============================================================================
-- FIM DO SCRIPT
-- ============================================================================
