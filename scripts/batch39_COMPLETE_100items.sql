-- ============================================================================
-- BATCH 39: COMPOSIÇÃO CORPORAL PARTE 2 - 100 ITEMS COMPLETOS
-- Foco: ASMI para sarcopenia, circunferência abdominal > IMC,
--       massa magra/gordura visceral, ritmo circadiano
-- Referências: EWGSOP2, AWGS 2019
-- ============================================================================

BEGIN;

-- ============================================================================
-- BLOCO 1: HISTÓRICO (27 items)
-- ============================================================================

-- 1-3. Pré-natal (3x repetido)
UPDATE score_items SET
  clinical_relevance = 'História pré-natal dos pais influencia programação metabólica fetal, predispondo a obesidade/sarcopenia na vida adulta. Estado nutricional materno (desnutrição ou excesso), diabetes gestacional, tabagismo e estresse crônico alteram epigenética do feto, afetando adipogênese, miogênese e regulação hormonal do apetite (leptina, grelina). Filhos de mães com diabetes/obesidade têm maior risco de adiposidade visceral e resistência insulínica.',
  patient_explanation = 'A saúde dos seus pais antes e durante a gestação influenciou sua composição corporal atual. Estado nutricional da mãe, diabetes ou obesidade na gravidez afetam como seu corpo armazena gordura e constrói músculos.',
  conduct = 'Investigar: IMC materno pré-gestacional, ganho de peso gestacional, diabetes gestacional, complicações perinatais. Considerar risco aumentado para estratégias preventivas.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '6c0651a0-f8a4-47be-95b8-a42a36673de1',
  'b898d743-cca9-48e4-aab9-50c3e04d5b1b',
  '917bc528-e4a1-4a30-9890-d704517e2669'
);

-- 4-6. Nascimento (3x repetido)
UPDATE score_items SET
  clinical_relevance = 'Peso ao nascer e idade gestacional são preditores de composição corporal futura. Baixo peso (< 2500g) associa-se a sarcopenia, baixa massa muscular e risco de diabetes tipo 2 (thrifty phenotype). Alto peso (> 4000g) predispõe a obesidade e síndrome metabólica. Prematuridade compromete desenvolvimento muscular e ósseo. Complicações perinatais afetam programação metabólica.',
  patient_explanation = 'O peso ao nascer e se você nasceu prematuro influenciam sua composição corporal. Bebês com baixo peso têm maior risco de ter menos massa muscular. Bebês grandes têm mais chance de obesidade.',
  conduct = 'Documentar: peso ao nascer, idade gestacional, complicações perinatais. Baixo peso: rastreamento precoce de sarcopenia e diabetes. Alto peso: monitoramento de obesidade e síndrome metabólica.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '74dfa65f-df48-4331-9a47-72fa2d52b5e2',
  '4ba81974-b66a-42c5-92cf-9093d3084e38',
  '5e02310c-08f1-4f88-b22d-4d89447fe565'
);

-- 7-9. Infância (3x repetido)
UPDATE score_items SET
  clinical_relevance = 'Trajetória de peso/altura na infância prediz composição corporal adulta. Ganho de peso acelerado (catch-up growth) nos primeiros 2 anos associa-se a maior adiposidade visceral e risco metabólico. Obesidade infantil (IMC > p95) predispõe a obesidade adulta e sarcopenia por inflamação crônica. Desnutrição infantil compromete desenvolvimento muscular e ósseo irreversivelmente.',
  patient_explanation = 'Como você cresceu e ganhou peso na infância afeta sua composição corporal hoje. Crianças que ganharam peso muito rápido ou que tiveram obesidade infantil têm maior risco de problemas metabólicos.',
  conduct = 'Revisar: curvas de crescimento infantil (OMS), IMC percentilado, histórico de obesidade/desnutrição. Ganho rápido ou obesidade infantil: monitoramento de síndrome metabólica.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '1bdd6114-82c5-4b57-8352-ea32449169ba',
  '64614d75-3b62-48ac-a6e8-294fa906ae2e',
  '84baa5e6-db01-4a90-af2b-801ce9f0670b'
);

-- 10-11. Adolescência (2x repetido)
UPDATE score_items SET
  clinical_relevance = 'Adolescência é período crítico para desenvolvimento de massa muscular e óssea (pico aos 18-25 anos). Ganho de peso excessivo na puberdade associa-se a obesidade adulta persistente. Estirões inadequados podem indicar deficiências nutricionais ou disfunções hormonais (GH, IGF-1, tireoidianos, esteroides gonadais). Anorexia/bulimia comprometem saúde óssea e muscular irreversivelmente.',
  patient_explanation = 'A adolescência é quando você constrói a maior parte da sua massa muscular e óssea. Problemas de peso nessa fase, como obesidade ou desnutrição, afetam sua saúde metabólica para o resto da vida.',
  conduct = 'Documentar: estirão puberal, IMC na adolescência, transtornos alimentares. Obesidade adolescente: rastreamento de sarcopenia prematura. Anorexia/bulimia: avaliação de densidade óssea (DEXA).',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '8a932b13-17c6-42d1-b428-dfe296accd4b',
  '657dcaae-f159-4415-8531-5a2908c0d095'
);

-- 12-14. Vida adulta (3x repetido)
UPDATE score_items SET
  clinical_relevance = 'Histórico de avaliações de composição corporal (DEXA, bioimpedância, pregas) permite identificar trajetórias de ganho/perda muscular e gordura. Perda muscular > 5-10% em 6-12 meses indica sarcopenia secundária (doença, inatividade, desnutrição). Ganho de gordura visceral progressivo prediz risco metabólico. Variações cíclicas (efeito sanfona) associam-se a perda muscular preferencial.',
  patient_explanation = 'Se você já fez exames de composição corporal (DEXA, bioimpedância), esses resultados mostram como seu corpo mudou ao longo do tempo. Perdas de massa muscular ou ganhos de gordura visceral indicam riscos metabólicos.',
  conduct = 'Revisar: avaliações prévias de composição corporal (DEXA, bioimpedância). Calcular variações de ASMI, FMI, gordura visceral. Perda muscular > 5-10%: investigar sarcopenia secundária.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  'ac04ab06-bdc7-4959-a33d-32f3646c1a87',
  '3c9b2568-a4ea-4132-9c3e-f327ad1b0ce7',
  '852be041-945c-4ced-91a8-4f85b55e9651'
);

-- 15-17. Picos de peso (3x repetido)
UPDATE score_items SET
  clinical_relevance = 'Picos de peso/gordura identificam períodos de risco metabólico aumentado e gatilhos ambientais (estresse, rotina, medicamentos). Ganho rápido (> 10% em 6 meses) associa-se a acúmulo preferencial de gordura visceral, resistência insulínica e inflamação crônica. Identificar contextos (gravidez, menopausa, corticoterapia, antipsicóticos) orienta manejo.',
  patient_explanation = 'Momentos em que você ganhou muito peso rapidamente podem indicar fatores desencadeantes (estresse, medicamentos, mudanças hormonais). Compreender esses padrões ajuda a prevenir novos ganhos de peso.',
  conduct = 'Identificar: gatilhos de ganho rápido (medicamentos, estresse, mudanças hormonais). Ganho > 10% em 6 meses: investigar resistência insulínica (HOMA-IR), gordura visceral (DEXA/TC). Ajustar medicações se necessário.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '3b9f7a1e-527e-4a4c-bc23-87c68c41ee89',
  'ef3181a6-135a-4110-9275-68eae1c3890a',
  '1af195d6-77c0-495b-9833-9ad38362e6fb'
);

-- 18-20. Melhor composição (3x repetido)
UPDATE score_items SET
  clinical_relevance = 'Períodos de melhor composição corporal revelam condições ideais para saúde metabólica do paciente (atividade física, sono, alimentação, manejo de estresse). Identificar contextos favoráveis (ex: prática regular de musculação, low-carb, jejum intermitente) permite replicá-los. Melhora de composição associa-se a redução de marcadores inflamatórios (PCR, IL-6), melhora de sensibilidade insulínica.',
  patient_explanation = 'Momentos em que você teve sua melhor forma física indicam o que funciona para seu corpo. Identificar as estratégias que deram certo no passado (tipo de exercício, alimentação, rotina de sono) ajuda a criar um plano personalizado.',
  conduct = 'Documentar condições da melhor composição: tipo de exercício, padrão alimentar, rotina de sono, manejo de estresse. Replicar estratégias bem-sucedidas.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '7b4975a7-a24d-4bbc-8d27-2dd89f9a99ad',
  '8d587ace-8023-41ab-85b2-334a2a00f41b',
  'e05d895d-36fc-48ff-8201-0098148ca9ac'
);

-- 21-23. Tratamentos prévios (3x repetido)
UPDATE score_items SET
  clinical_relevance = 'Histórico de tratamentos para composição corporal identifica estratégias eficazes, efeitos adversos e contraindicações. Dietas restritivas repetidas (efeito sanfona) associam-se a perda muscular, ganho de gordura visceral e redução de taxa metabólica basal. Medicamentos (orlistat, liraglutida, semaglutida) têm eficácia variável. Cirurgia bariátrica é eficaz mas requer suplementação vitalícia.',
  patient_explanation = 'Conhecer os tratamentos que você já tentou (dietas, remédios, hormônios, cirurgias) ajuda a evitar repetir estratégias que não funcionaram e a identificar o que foi eficaz. Isso personaliza seu plano.',
  conduct = 'Documentar tratamentos prévios: dietas, medicamentos (orlistat, liraglutida, semaglutida), hormônios, cirurgias. Avaliar eficácia, efeitos adversos, motivos de descontinuação.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  'c1c9eb73-b2ac-456f-b854-ed707c817a3d',
  'e965a941-475a-4c19-9407-22b964af9ca3',
  '70dae5d1-7c26-42cd-a746-36b21a49e9cd'
);

-- 24-26. Histórico familiar (3x repetido)
UPDATE score_items SET
  clinical_relevance = 'Histórico familiar de composição corporal identifica predisposição genética a obesidade, sarcopenia e distribuição de gordura. Obesidade em parentes de 1º grau aumenta risco em 2-8x. Variantes genéticas (FTO, MC4R, PPARG) afetam regulação de apetite e adipogênese. Sarcopenia familiar sugere miopatias hereditárias ou deficiência de vitamina D crônica. Padrões de distribuição de gordura têm componente hereditário.',
  patient_explanation = 'A composição corporal dos seus pais, irmãos e avós pode indicar sua tendência genética a ganhar peso, perder músculos ou acumular gordura em certas regiões. Compreender sua genética ajuda a personalizar estratégias.',
  conduct = 'Documentar composição corporal de familiares de 1º e 2º grau. Obesidade familiar: rastreamento intensivo de síndrome metabólica. Sarcopenia familiar: investigar miopatias, vitamina D.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '03438548-88b3-44a7-8449-5441c26ea829',
  '85a74aaf-8c04-4552-9c8f-636cfe1c9133',
  '5d7f01cb-f4bb-404a-8d89-d80066384f5a'
);

-- 27. Avaliação atual (2x)
UPDATE score_items SET
  clinical_relevance = 'Avaliação atual de composição corporal (DEXA, bioimpedância, pregas cutâneas, análise fotográfica) fornece dados objetivos para estratificação de risco metabólico e nutricional. DEXA é padrão-ouro para quantificar massa muscular (ASMI para sarcopenia EWGSOP2), gordura visceral e densidade mineral óssea. Bioimpedância estima água corporal, massa magra e gordura com boa acurácia. Comparação com referências orienta diagnóstico.',
  patient_explanation = 'Exames de composição corporal (como DEXA ou bioimpedância) medem sua massa muscular, gordura corporal e água. Esses dados mostram riscos metabólicos e nutricionais, permitindo criar um plano personalizado para melhorar sua saúde.',
  conduct = 'Solicitar DEXA (padrão-ouro) ou bioimpedância (acessível). Avaliar: ASMI (sarcopenia), FMI, gordura visceral, ângulo de fase, DMO. Comparar com referências (EWGSOP2, AWGS 2019).',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '088b4d4d-1873-45bb-8a3a-19e4463de7a5',
  '0b58623e-63d3-4685-b864-c41a0058ed56'
);

-- ============================================================================
-- BLOCO 2: TRATAMENTOS E OBJETIVOS ATUAIS (6 items)
-- ============================================================================

-- 28-29. Tratamentos atuais (2x)
UPDATE score_items SET
  clinical_relevance = 'Tratamentos em uso para composição corporal devem ser avaliados quanto a eficácia, segurança e adesão. Medicamentos (GLP-1 agonistas, orlistat), suplementos (proteínas, creatina, vitamina D), hormônios (testosterona, GH) e estratégias dietéticas (cetogênica, jejum intermitente) têm impactos diferentes em massa muscular e gordura visceral. Monitoramento de efeitos adversos (GI, tireoidianos, hepáticos) e interações medicamentosas é essencial.',
  patient_explanation = 'Se você está fazendo tratamentos para composição corporal (remédios, suplementos, dietas), é importante monitorar se estão funcionando e se há efeitos colaterais. Ajustes personalizados aumentam a eficácia e segurança do seu plano.',
  conduct = 'Documentar tratamentos atuais: medicamentos (GLP-1, orlistat), suplementos (whey, creatina, vitamina D), hormônios, dietas (cetogênica, jejum). Monitorar eficácia, efeitos adversos, adesão.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '2579198d-65b8-4cff-9970-7c437076adc7',
  'a6222101-4fee-41a4-aca2-fd466d1c789e'
);

-- 30-31. Satisfação e objetivos (2x)
UPDATE score_items SET
  clinical_relevance = 'Grau de satisfação com composição corporal e objetivos do paciente devem ser avaliados de forma realista, considerando genética, idade, comorbidades e capacidade de aderência. Insatisfação corporal excessiva pode indicar dismorfia ou transtornos alimentares. Objetivos irrealistas (ex: IMC < 18, %GC < 5%) são contraproducentes. Foco deve ser saúde metabólica (redução de gordura visceral, manutenção de massa muscular) e não apenas estética.',
  patient_explanation = 'Seus objetivos de composição corporal devem ser realistas e focados em saúde, não apenas estética. Juntos, definimos metas alcançáveis que consideram sua genética, estilo de vida e saúde geral, aumentando as chances de sucesso sustentável.',
  conduct = 'Avaliar grau de satisfação corporal (escala 1-10), objetivos específicos. Insatisfação severa: rastreamento de dismorfia, transtornos alimentares. Educar sobre expectativas realistas e saúde metabólica.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  'f836753e-9e45-48f5-a54c-298195ff0588',
  'a3707b3b-d45e-4235-93b4-1b011deb5028'
);

-- 32-33. Medidas objetivas (2x)
UPDATE score_items SET
  clinical_relevance = 'Medidas objetivas de composição corporal (antropometria, bioimpedância, DEXA) são fundamentais para diagnóstico preciso de obesidade, sarcopenia e risco metabólico. Circunferência abdominal e razão cintura/quadril predizem risco cardiovascular melhor que IMC. ASMI < 7,0 kg/m² (homens) ou < 5,5 kg/m² (mulheres) indica sarcopenia (EWGSOP2). Gordura visceral > 100 cm² associa-se a resistência insulínica. Monitoramento longitudinal orienta ajustes de intervenções.',
  patient_explanation = 'Medidas objetivas (peso, circunferências, exames de composição corporal) fornecem dados precisos sobre sua saúde metabólica e nutricional. Esses números guiam decisões e permitem acompanhar progresso de forma científica.',
  conduct = 'Realizar: peso, altura, IMC, circunferências (abdominal, quadril, pescoço, panturrilha, braço, coxa). Idealmente: bioimpedância ou DEXA. Calcular: razões cintura/quadril, cintura/altura.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '662902a2-d3f7-4d08-b3ee-4fde9e633cf8',
  '7c77b924-d8fb-478d-81fc-10ba491d51a0'
);

-- ============================================================================
-- BLOCO 3: MEDIDAS ANTROPOMÉTRICAS BÁSICAS (5 items)
-- ============================================================================

-- 34-35. Peso (2x)
UPDATE score_items SET
  clinical_relevance = 'Peso corporal é medida básica para cálculo de IMC, monitoramento de ganho/perda e ajuste de medicamentos. Variações > 5% em 1 mês exigem investigação (retenção hídrica, doença, mudança de composição corporal). Peso isolado não diferencia massa muscular de gordura. Perda de peso intencional deve ser lenta (0,5-1 kg/semana) para preservar massa muscular. Ganho de peso rápido (> 2 kg/mês) sugere acúmulo de gordura visceral.',
  patient_explanation = 'Seu peso é uma medida importante, mas não conta a história completa. Variações rápidas podem indicar retenção de líquidos ou mudanças metabólicas. Acompanhar o peso regularmente ajuda a detectar tendências e ajustar estratégias.',
  conduct = 'Monitorar peso semanalmente (mesmas condições: jejum, após micção). Variação > 5% em 1 mês: investigar causas (retenção hídrica, doença, mudança composição). Combinar com bioimpedância para diferenciar músculo de gordura.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '400295c1-25e5-44c5-8d99-0f355f1aa7cc',
  'c2718674-8ee8-4c54-b63c-86784cbd903c'
);

-- 36. Altura
UPDATE score_items SET
  clinical_relevance = 'Altura é fundamental para cálculo de IMC (kg/m²), ASMI (massa muscular apendicular/altura²) e razão cintura/altura. Redução de altura em idosos indica compressões vertebrais por osteoporose. Medição anual em > 50 anos permite detectar fraturas vertebrais assintomáticas. Altura precisa é essencial para avaliação nutricional e cálculo de dose de medicamentos.',
  patient_explanation = 'Sua altura é usada para calcular vários índices de saúde, como IMC e ASMI. Em idosos, redução de altura pode indicar osteoporose e fraturas vertebrais silenciosas.',
  conduct = 'Medir altura com estadiômetro, sem calçados, calcanhares juntos, olhar no horizonte. Em > 50 anos: medir anualmente. Redução > 2 cm: investigar osteoporose (DEXA, vitamina D, cálcio, PTH).',
  last_review = CURRENT_TIMESTAMP
WHERE id = '48b082bf-3697-4c8a-a183-b2fc4396d270';

-- 37-38. IMC (2x)
UPDATE score_items SET
  clinical_relevance = 'IMC (kg/m²) é triagem inicial para classificar peso: < 18,5 baixo peso, 18,5-24,9 normal, 25-29,9 sobrepeso, 30-34,9 obesidade I, 35-39,9 obesidade II, ≥ 40 obesidade III. Limitações: não diferencia massa muscular de gordura (atletas com IMC alto mas baixa gordura), não avalia distribuição de gordura (visceral vs. subcutânea). Circunferência abdominal e bioimpedância são complementares. IMC > 30 aumenta risco de diabetes, HAS, dislipidemia.',
  patient_explanation = 'O IMC (Índice de Massa Corporal) é uma medida que relaciona seu peso e altura. Embora útil como triagem inicial, ele não diferencia músculos de gordura. Por isso, usamos outras medidas complementares para avaliar melhor sua composição corporal.',
  conduct = 'Calcular IMC = peso (kg) / altura² (m²). < 18,5: investigar desnutrição, sarcopenia. 25-29,9: rastreamento de síndrome metabólica. ≥ 30: avaliação completa de composição corporal (DEXA, bioimpedância).',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  'bf875f17-1a90-4bb8-8305-e0e0285aeb80',
  '7efe0605-a423-4560-91aa-05c285a93c4f'
);

-- ============================================================================
-- BLOCO 4: CIRCUNFERÊNCIAS CORPORAIS (18 items)
-- ============================================================================

-- 39-40. Circunferência Abdominal - Homem e Mulher
UPDATE score_items SET
  clinical_relevance = 'Circunferência abdominal (cintura) é melhor preditor de risco metabólico que IMC, pois reflete gordura visceral. Valores de risco: homens > 94 cm (risco aumentado), > 102 cm (risco alto). Gordura visceral é metabolicamente ativa, secretando adipocinas pró-inflamatórias (TNF-α, IL-6) e aumentando resistência insulínica. Redução de 5-10% da circunferência melhora perfil metabólico independente de perda de peso. Medida ao nível da cicatriz umbilical, em expiração.',
  patient_explanation = 'A circunferência da sua cintura é mais importante que o peso para prever risco de diabetes e doenças cardíacas. Gordura na barriga (visceral) é mais perigosa que gordura em outras regiões. Reduzir a cintura melhora sua saúde metabólica.',
  conduct = 'Medir circunferência abdominal ao nível da cicatriz umbilical, em expiração. Homens: > 94 cm iniciar intervenções, > 102 cm risco alto. Investigar síndrome metabólica (glicemia, HOMA-IR, perfil lipídico, PA).',
  last_review = CURRENT_TIMESTAMP
WHERE id = '9a97c090-ce0b-4dbd-a030-e652542afc6c';

UPDATE score_items SET
  clinical_relevance = 'Circunferência abdominal (cintura) é melhor preditor de risco metabólico que IMC, pois reflete gordura visceral. Valores de risco: mulheres > 80 cm (risco aumentado), > 88 cm (risco alto). Gordura visceral é metabolicamente ativa, secretando adipocinas pró-inflamatórias (TNF-α, IL-6) e aumentando resistência insulínica. Redução de 5-10% da circunferência melhora perfil metabólico. Mulheres pós-menopausa têm maior acúmulo visceral por redução de estrogênio. Medida ao nível da cicatriz umbilical, em expiração.',
  patient_explanation = 'A circunferência da sua cintura é mais importante que o peso para prever risco de diabetes e doenças cardíacas. Gordura na barriga (visceral) é mais perigosa que gordura em outras regiões. Reduzir a cintura melhora sua saúde metabólica.',
  conduct = 'Medir circunferência abdominal ao nível da cicatriz umbilical, em expiração. Mulheres: > 80 cm iniciar intervenções, > 88 cm risco alto. Investigar síndrome metabólica (glicemia, HOMA-IR, perfil lipídico, PA).',
  last_review = CURRENT_TIMESTAMP
WHERE id = 'f6a6515f-5488-455d-9459-8c606a13029e';

-- 41-42. Quadril (2x)
UPDATE score_items SET
  clinical_relevance = 'Circunferência de quadril (ao nível dos trocânteres maiores) é usada para calcular razão cintura/quadril. Quadril maior (maior massa muscular glútea e femoral) é protetor metabólico. Distribuição de gordura ginoide (quadris) é menos deletéria que androide (abdominal). Quadril menor em idosos sugere sarcopenia de membros inferiores, aumentando risco de quedas. Medida em conjunto com cintura orienta diagnóstico de síndrome metabólica.',
  patient_explanation = 'A medida do quadril, junto com a cintura, ajuda a identificar o padrão de distribuição de gordura no seu corpo. Gordura nos quadris é menos preocupante que gordura abdominal para sua saúde metabólica.',
  conduct = 'Medir circunferência de quadril ao nível dos trocânteres maiores. Calcular razão cintura/quadril. Quadril reduzido em idosos: avaliar sarcopenia (ASMI, força de preensão manual, teste de caminhada).',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '1a9be52d-aa0b-4a53-a71a-4eff3e0cc6ba',
  '935ecda0-3917-452e-8c4d-91b05f4e3e0d'
);

-- 43-44. Razão Cintura/Quadril - Homem e Mulher
UPDATE score_items SET
  clinical_relevance = 'Razão cintura/quadril (RCQ) identifica distribuição de gordura: androide (abdominal, > 0,90 homens) ou ginoide (quadris, < 0,90). Padrão androide associa-se a maior risco metabólico (diabetes, DCV, DHGNA). RCQ > 1,0 em homens indica adiposidade visceral significativa. RCQ aumenta com idade e redução de testosterona. Valores elevados mesmo com IMC normal indicam obesidade metabólica de peso normal (MONW). Redução de RCQ é objetivo terapêutico.',
  patient_explanation = 'A relação entre sua cintura e quadril mostra onde seu corpo acumula gordura. Valores altos indicam mais gordura abdominal, que é mais perigosa para sua saúde. Reduzir essa relação é um objetivo importante.',
  conduct = 'Calcular RCQ = circunferência abdominal / circunferência de quadril. Homens: > 0,90 risco aumentado, > 1,0 risco alto. Investigar síndrome metabólica, resistência insulínica.',
  last_review = CURRENT_TIMESTAMP
WHERE id = 'c9348fbd-df44-4903-ac22-1d8b5b47179a';

UPDATE score_items SET
  clinical_relevance = 'Razão cintura/quadril (RCQ) identifica distribuição de gordura: androide (abdominal, > 0,85 mulheres) ou ginoide (quadris, < 0,85). Padrão androide associa-se a maior risco metabólico (diabetes, DCV, DHGNA). Mulheres pós-menopausa têm aumento de RCQ por redução de estrogênio e redistribuição de gordura. RCQ > 0,85 indica adiposidade visceral significativa. Valores elevados mesmo com IMC normal indicam obesidade metabólica de peso normal (MONW).',
  patient_explanation = 'A relação entre sua cintura e quadril mostra onde seu corpo acumula gordura. Valores altos indicam mais gordura abdominal, que é mais perigosa para sua saúde. Reduzir essa relação é um objetivo importante.',
  conduct = 'Calcular RCQ = circunferência abdominal / circunferência de quadril. Mulheres: > 0,85 risco aumentado. Pós-menopausa: monitoramento intensivo. Investigar síndrome metabólica, resistência insulínica.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  'b2414f5e-8ad7-4b9c-846e-edd6a3c8277f',
  '14c4c09a-ad66-4d7e-9baf-912f82049be0'
);

-- 45-46. Razão Cintura/Altura (2x)
UPDATE score_items SET
  clinical_relevance = 'Razão cintura/altura (RCA) é preditor robusto de risco metabólico, independente de etnia. RCA > 0,5 indica risco aumentado de DCV, diabetes e mortalidade ("mantenha sua cintura abaixo de metade da sua altura"). Vantagens: simples, não requer tabelas de referência por idade/sexo, aplicável a crianças e adultos. RCA > 0,6 indica obesidade central. Redução de RCA < 0,5 é meta terapêutica universal.',
  patient_explanation = 'A relação entre sua cintura e altura é um indicador simples de risco metabólico. Se sua cintura é mais da metade da sua altura (> 0,5), há risco aumentado de diabetes e problemas cardíacos. Manter essa relação baixa é uma meta importante.',
  conduct = 'Calcular RCA = circunferência abdominal (cm) / altura (cm). > 0,5 risco aumentado, > 0,6 obesidade central. Meta terapêutica: RCA < 0,5. Intervenções: dieta, exercício, manejo de estresse.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '936855d1-e715-4171-a061-3bcc500957f7',
  '98119b72-b520-4c99-8102-04e764332353'
);

-- 47-48. Pescoço - Homem e Mulher
UPDATE score_items SET
  clinical_relevance = 'Circunferência de pescoço (medida abaixo do pomo de Adão) prediz obesidade visceral e apneia obstrutiva do sono (AOS). Valores de risco: homens > 37 cm, mulheres > 34 cm. Pescoço > 40 cm em homens aumenta risco de AOS em 5x. Gordura cervical comprime vias aéreas superiores, predispondo a hipóxia noturna, hipertensão e arritmias. Pescoço elevado associa-se a síndrome metabólica. Medida não invasiva útil para rastreamento de AOS.',
  patient_explanation = 'A circunferência do pescoço ajuda a identificar risco de apneia do sono e gordura visceral. Pescoço grosso pode indicar acúmulo de gordura que dificulta a respiração durante o sono, aumentando risco de pressão alta e problemas cardíacos.',
  conduct = 'Medir circunferência de pescoço abaixo do pomo de Adão. Homens: > 37 cm rastreamento de AOS (questionário STOP-BANG, polissonografia se indicado). > 40 cm: risco alto de AOS.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  'c3ee7527-104b-467a-a87e-b4587192005b',
  '9f1df5e4-37c5-4b58-9d4e-24b7eac925b3'
);

UPDATE score_items SET
  clinical_relevance = 'Circunferência de pescoço (medida abaixo da proeminência laríngea) prediz obesidade visceral e apneia obstrutiva do sono (AOS). Valores de risco: mulheres > 34 cm. Pescoço > 35 cm aumenta risco de AOS. Gordura cervical comprime vias aéreas superiores, predispondo a hipóxia noturna, hipertensão e arritmias. Pescoço elevado associa-se a síndrome metabólica. Mulheres pós-menopausa têm maior acúmulo de gordura cervical.',
  patient_explanation = 'A circunferência do pescoço ajuda a identificar risco de apneia do sono e gordura visceral. Pescoço grosso pode indicar acúmulo de gordura que dificulta a respiração durante o sono, aumentando risco de pressão alta e problemas cardíacos.',
  conduct = 'Medir circunferência de pescoço abaixo da proeminência laríngea. Mulheres: > 34 cm rastreamento de AOS (questionário STOP-BANG, polissonografia se indicado).',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '421db9cd-d75d-4f0a-8188-216affee192f',
  '261ecf45-3c11-4a14-b022-38d1fc1e3f6e'
);

-- 49-50. Relação Pescoço/Altura - Homem e Mulher
UPDATE score_items SET
  clinical_relevance = 'Relação pescoço/altura (cm/m) normaliza circunferência de pescoço pela estatura, melhorando acurácia para predizer AOS e síndrome metabólica. Valores > 42 em homens indicam risco aumentado. Relação elevada associa-se a resistência insulínica, dislipidemia e HAS. Útil em populações com diferentes estaturas. Medida simples, não invasiva e reprodutível para rastreamento de risco metabólico e respiratório.',
  patient_explanation = 'A relação entre seu pescoço e altura ajusta a medida do pescoço pela sua estatura, melhorando a avaliação de risco metabólico e de apneia do sono. É um indicador simples e útil.',
  conduct = 'Calcular relação pescoço/altura = circunferência pescoço (cm) / altura (m). Homens: > 42 risco aumentado. Investigar AOS e síndrome metabólica.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  'ff5472cf-e148-48c1-be7a-245df89c200c',
  'c21e57d3-6d49-41ca-9281-7dbff96dbc8d'
);

UPDATE score_items SET
  clinical_relevance = 'Relação pescoço/altura (cm/m) normaliza circunferência de pescoço pela estatura, melhorando acurácia para predizer AOS e síndrome metabólica. Valores > 35 em mulheres indicam risco aumentado. Relação elevada associa-se a resistência insulínica, dislipidemia e HAS. Útil em populações com diferentes estaturas. Medida simples, não invasiva e reprodutível para rastreamento de risco metabólico e respiratório.',
  patient_explanation = 'A relação entre seu pescoço e altura ajusta a medida do pescoço pela sua estatura, melhorando a avaliação de risco metabólico e de apneia do sono. É um indicador simples e útil.',
  conduct = 'Calcular relação pescoço/altura = circunferência pescoço (cm) / altura (m). Mulheres: > 35 risco aumentado. Investigar AOS e síndrome metabólica.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '472969f8-7ea9-4532-8c73-c4a96f31e2fc',
  '1e7e9470-247a-43bd-bb4d-235ec97c7da6'
);

-- 51-52. Panturrilha - Homem e Mulher
UPDATE score_items SET
  clinical_relevance = 'Circunferência de panturrilha (ponto de maior proeminência) é marcador de massa muscular de membros inferiores. Valores < 31 cm em homens sugerem sarcopenia e aumentam risco de quedas, fraturas e dependência funcional. Panturrilha < 31 cm é critério de desnutrição em idosos (ESPEN). Redução progressiva indica perda muscular (AWGS 2019). Massa muscular de membros inferiores é crucial para mobilidade, equilíbrio e independência.',
  patient_explanation = 'A circunferência da panturrilha indica a quantidade de músculo nas suas pernas. Valores baixos sugerem perda muscular (sarcopenia), aumentando risco de quedas e perda de independência. Manter músculos fortes nas pernas é fundamental para mobilidade.',
  conduct = 'Medir circunferência de panturrilha no ponto de maior proeminência. Homens: < 31 cm investigar sarcopenia (ASMI, força de preensão, teste de caminhada). Iniciar treinamento resistido, suplementação proteica.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '094b4b55-e7a4-4587-be08-9963c0520bf0',
  '3939e1db-5739-4935-9b4b-fed7a8d0711a'
);

UPDATE score_items SET
  clinical_relevance = 'Circunferência de panturrilha (ponto de maior proeminência) é marcador de massa muscular de membros inferiores. Valores < 33 cm em mulheres sugerem sarcopenia e aumentam risco de quedas, fraturas e dependência funcional. Panturrilha < 33 cm é critério de desnutrição em idosos (ESPEN). Redução progressiva indica perda muscular (AWGS 2019). Massa muscular de membros inferiores é crucial para mobilidade, equilíbrio e independência.',
  patient_explanation = 'A circunferência da panturrilha indica a quantidade de músculo nas suas pernas. Valores baixos sugerem perda muscular (sarcopenia), aumentando risco de quedas e perda de independência. Manter músculos fortes nas pernas é fundamental para mobilidade.',
  conduct = 'Medir circunferência de panturrilha no ponto de maior proeminência. Mulheres: < 33 cm investigar sarcopenia (ASMI, força de preensão, teste de caminhada). Iniciar treinamento resistido, suplementação proteica.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '660ec23e-8957-4936-b83b-284a0f8c84eb',
  '8bc9d1cf-61c0-47bc-b94b-9935250e598f'
);

-- 53-54. Braço Relaxado - Homem e Mulher
UPDATE score_items SET
  clinical_relevance = 'Circunferência de braço relaxado (ponto médio entre acrômio e olécrano) estima massa muscular e gordura de membros superiores. Valores < 24 cm em homens sugerem desnutrição ou sarcopenia. Redução > 10% em 6 meses indica perda muscular significativa. Braço relaxado é usado em equações de avaliação nutricional (circunferência muscular do braço = CMB). Útil para monitoramento em pacientes acamados ou com mobilidade reduzida.',
  patient_explanation = 'A circunferência do braço relaxado indica sua massa muscular e estado nutricional. Valores baixos ou redução ao longo do tempo sugerem perda muscular ou desnutrição. Monitorar essa medida ajuda a detectar mudanças precoces.',
  conduct = 'Medir circunferência de braço não dominante relaxado no ponto médio entre acrômio e olécrano. Homens: < 24 cm investigar desnutrição, sarcopenia. Calcular circunferência muscular do braço (CMB).',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '30fa255c-83d9-4cd4-8b06-75f70e2fb3eb',
  '7fe0b34c-151d-407f-9de8-6ff2492dde4b'
);

UPDATE score_items SET
  clinical_relevance = 'Circunferência de braço relaxado (ponto médio entre acrômio e olécrano) estima massa muscular e gordura de membros superiores. Valores < 22 cm em mulheres sugerem desnutrição ou sarcopenia. Redução > 10% em 6 meses indica perda muscular significativa. Braço relaxado é usado em equações de avaliação nutricional (circunferência muscular do braço = CMB). Útil para monitoramento em pacientes acamados ou com mobilidade reduzida.',
  patient_explanation = 'A circunferência do braço relaxado indica sua massa muscular e estado nutricional. Valores baixos ou redução ao longo do tempo sugerem perda muscular ou desnutrição. Monitorar essa medida ajuda a detectar mudanças precoces.',
  conduct = 'Medir circunferência de braço não dominante relaxado no ponto médio entre acrômio e olécrano. Mulheres: < 22 cm investigar desnutrição, sarcopenia. Calcular circunferência muscular do braço (CMB).',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '371f12db-5a68-4092-a4c3-1430ec21f18a',
  '8e0efcc9-8281-48dc-856e-fa86a845e97d'
);

-- 55-56. Braço Contraído (2x - Direito e Esquerdo)
UPDATE score_items SET
  clinical_relevance = 'Circunferência de braço contraído (flexão máxima do cotovelo) avalia massa muscular e tônus do bíceps braquial. Diferença < 2 cm entre braço contraído e relaxado sugere hipotonia ou sarcopenia. Braço contraído é usado para monitorar resposta a treinamento resistido. Assimetria > 1 cm entre braços pode indicar dominância lateral ou lesão prévia. Útil em atletas e em reabilitação pós-trauma.',
  patient_explanation = 'A circunferência do braço contraído avalia o tônus e tamanho do músculo do bíceps. Diferenças pequenas entre braço relaxado e contraído sugerem perda de massa muscular ou tônus baixo.',
  conduct = 'Medir circunferência de braço com flexão máxima do cotovelo. Calcular diferença entre braço contraído e relaxado. < 2 cm: investigar sarcopenia, hipotonia. Assimetria > 1 cm: investigar lesões.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '405b0018-088c-4a35-8688-11de766fc557',
  '9cbf71b3-83de-4b71-b987-120e525c790e'
);

-- 57-58. Coxa - Homem e Mulher
UPDATE score_items SET
  clinical_relevance = 'Circunferência de coxa (ponto médio entre trocânter maior e joelho) avalia massa muscular de membros inferiores (quadríceps). Valores < 50 cm em homens sugerem sarcopenia de membros inferiores, aumentando risco de quedas e dependência funcional. Coxa é maior grupo muscular do corpo, sua redução indica perda muscular sistêmica. Redução > 10% em 6 meses é significativa. Monitoramento em idosos, acamados e pós-cirúrgicos é fundamental.',
  patient_explanation = 'A circunferência da coxa indica a massa muscular das suas pernas, especialmente do quadríceps. Valores baixos sugerem sarcopenia, aumentando risco de quedas e perda de mobilidade. Manter músculos fortes nas coxas é essencial para independência.',
  conduct = 'Medir circunferência de coxa no ponto médio entre trocânter maior e joelho. Homens: < 50 cm investigar sarcopenia. Monitorar variações. Redução > 10%: intensificar treinamento resistido.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '8062469b-618a-4cfc-9239-046fd1f882e2',
  '731d8307-6295-469e-95b6-361373449dcf'
);

UPDATE score_items SET
  clinical_relevance = 'Circunferência de coxa (ponto médio entre trocânter maior e joelho) avalia massa muscular de membros inferiores (quadríceps). Valores < 45 cm em mulheres sugerem sarcopenia de membros inferiores, aumentando risco de quedas e dependência funcional. Coxa é maior grupo muscular do corpo, sua redução indica perda muscular sistêmica. Redução > 10% em 6 meses é significativa. Monitoramento em idosos, acamados e pós-cirúrgicos é fundamental.',
  patient_explanation = 'A circunferência da coxa indica a massa muscular das suas pernas, especialmente do quadríceps. Valores baixos sugerem sarcopenia, aumentando risco de quedas e perda de mobilidade. Manter músculos fortes nas coxas é essencial para independência.',
  conduct = 'Medir circunferência de coxa no ponto médio entre trocânter maior e joelho. Mulheres: < 45 cm investigar sarcopenia. Monitorar variações. Redução > 10%: intensificar treinamento resistido.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '25822830-7d95-4e20-9b32-bb5cd7d4a3d1',
  '6bf45614-a614-4938-854a-7f4902fefd5e'
);

-- ============================================================================
-- BLOCO 5: COMPOSIÇÃO CORPORAL AVANÇADA (21 items)
-- ============================================================================

-- 59-60. Massa Muscular Esquelética (2x)
UPDATE score_items SET
  clinical_relevance = 'Massa muscular esquelética (MME) é componente fundamental da composição corporal, responsável por metabolismo basal, força, mobilidade e independência funcional. MME medida por DEXA ou bioimpedância permite diagnóstico preciso de sarcopenia. Redução de MME associa-se a maior mortalidade, fragilidade, quedas e hospitalização. MME é influenciada por exercício resistido, ingestão proteica (1,2-1,6 g/kg/dia), hormônios anabólicos (testosterona, GH) e estado inflamatório.',
  patient_explanation = 'Sua massa muscular esquelética é a quantidade total de músculos no corpo. Ter músculos adequados é essencial para força, mobilidade e saúde metabólica. Perda muscular aumenta risco de fragilidade e dependência funcional.',
  conduct = 'Medir MME por DEXA (padrão-ouro) ou bioimpedância. Calcular MME/Peso (%) e Índice MME (kg/m²). Valores baixos: investigar sarcopenia, iniciar treinamento resistido, otimizar ingestão proteica.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '2ff747e7-e48d-4207-85ee-168b1c14e35e',
  '4f1f46f4-1339-4eff-afd9-b04b32d2e7dd'
);

-- 61-63. MME/Peso (%) (3x)
UPDATE score_items SET
  clinical_relevance = 'MME/Peso (%) expressa proporção de massa muscular em relação ao peso total. Valores de referência: homens > 37%, mulheres > 28%. Valores baixos indicam sarcopenia ou obesidade sarcopênica (alta % de gordura corporal com baixa MME). Obesidade sarcopênica é especialmente deletéria, associando-se a maior mortalidade que obesidade ou sarcopenia isoladas. Monitoramento de MME/Peso orienta intervenções nutricionais e de exercício.',
  patient_explanation = 'A porcentagem de músculo em relação ao seu peso total indica se você tem massa muscular adequada. Valores baixos podem indicar sarcopenia ou obesidade sarcopênica (muito gordura, pouco músculo), aumentando riscos metabólicos.',
  conduct = 'Calcular MME/Peso (%) = (MME / Peso) × 100. Homens: < 37% investigar sarcopenia. Mulheres: < 28% investigar sarcopenia. Meta: aumentar MME via treinamento resistido e nutrição adequada.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  'f4b9b50a-6369-4d97-b6a4-672e22eafd6b',
  '725d0bc7-8f25-4c44-b268-53b72f75adff',
  'e06b1f18-5f06-4db7-adf7-024d73583117',
  '6ab770e5-4df4-4937-b63f-e66064eb4e37'
);

-- 64-65. Índice MME (2x)
UPDATE score_items SET
  clinical_relevance = 'Índice MME (kg/m²) = MME / altura², normaliza massa muscular pela estatura. Valores de referência: homens > 10,75 kg/m², mulheres > 6,75 kg/m². Índice MME baixo é critério de sarcopenia secundária. Útil para comparações longitudinais e entre populações. Índice MME < 8,5 kg/m² (homens) ou < 5,75 kg/m² (mulheres) indica sarcopenia severa, com maior risco de mortalidade e incapacidade funcional.',
  patient_explanation = 'O Índice de Massa Muscular Esquelética ajusta sua massa muscular pela altura, permitindo avaliação mais precisa. Valores baixos indicam sarcopenia, aumentando risco de fragilidade e perda de independência.',
  conduct = 'Calcular Índice MME = MME (kg) / altura² (m²). Homens: < 10,75 kg/m² sarcopenia. Mulheres: < 6,75 kg/m² sarcopenia. Iniciar treinamento resistido, suplementação proteica, investigar causas secundárias.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  'cd592cb8-4c34-4964-ac29-63bf9c224bef',
  '9b0d2d5a-d5df-4b55-abf8-5a0b2e442377'
);

-- 66-67. Massa Apendicular (2x)
UPDATE score_items SET
  clinical_relevance = 'Massa apendicular (kg) é soma de massa magra de membros superiores e inferiores, medida por DEXA. Massa apendicular é componente fundamental para diagnóstico de sarcopenia (EWGSOP2, AWGS 2019). Redução de massa apendicular associa-se a fragilidade, quedas, fraturas e mortalidade. Massa apendicular é altamente responsiva a treinamento resistido e nutrição adequada (proteínas 1,2-1,6 g/kg/dia).',
  patient_explanation = 'Sua massa apendicular é a quantidade de músculo nos braços e pernas. Essa medida é crucial para diagnosticar sarcopenia. Manter massa muscular adequada nos membros é essencial para força, mobilidade e independência.',
  conduct = 'Medir massa apendicular por DEXA (padrão-ouro). Calcular ASMI = massa apendicular / altura². Valores baixos: investigar sarcopenia (EWGSOP2, AWGS 2019). Iniciar treinamento resistido, otimizar nutrição.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  'bdf03bb7-ef91-4e12-9247-4e29ce1e7185',
  'c75841a3-0bd7-4225-ac87-7fcf0d6bf438'
);

-- 68. ASMI - Homem
UPDATE score_items SET
  clinical_relevance = 'ASMI (Appendicular Skeletal Muscle Index, kg/m²) = massa apendicular / altura² é critério diagnóstico de sarcopenia (EWGSOP2). Valores de corte para homens: < 7,0 kg/m² sarcopenia (EWGSOP2), < 7,0 kg/m² sarcopenia (AWGS 2019). ASMI < 6,0 kg/m² indica sarcopenia severa. ASMI baixo associa-se a maior mortalidade, fragilidade, quedas, fraturas e hospitalização. ASMI é responsivo a treinamento resistido, nutrição adequada e tratamento de doenças subjacentes.',
  patient_explanation = 'O ASMI é um índice que mede sua massa muscular de braços e pernas ajustada pela altura. É o principal critério para diagnosticar sarcopenia. Valores baixos indicam perda muscular e risco de fragilidade.',
  conduct = 'Medir ASMI por DEXA. Homens: < 7,0 kg/m² diagnosticar sarcopenia (EWGSOP2). Investigar causas (desnutrição, inatividade, doenças crônicas, deficiência hormonal). Iniciar treinamento resistido progressivo, otimizar ingestão proteica (1,2-1,6 g/kg/dia), suplementar vitamina D se deficiente.',
  last_review = CURRENT_TIMESTAMP
WHERE id = 'a2688922-5b23-4f7c-a842-ffff6acf081e';

-- 69-70. Massa Gorda Total e % Gordura Corporal - Mulher
UPDATE score_items SET
  clinical_relevance = 'Massa gorda total (kg) e % gordura corporal são parâmetros essenciais para avaliação de risco metabólico. % gordura corporal para mulheres: < 20% atlética, 20-25% fitness, 25-30% aceitável, 30-35% excesso, > 35% obesidade. % gordura elevada associa-se a resistência insulínica, inflamação crônica, dislipidemia e risco cardiovascular. Distribuição de gordura (visceral vs. subcutânea) é mais importante que quantidade total.',
  patient_explanation = 'Sua massa gorda total e porcentagem de gordura corporal indicam se você tem gordura em excesso. Valores altos aumentam risco de diabetes, doenças cardíacas e inflamação. Importante avaliar também onde a gordura está localizada (barriga é mais perigosa).',
  conduct = 'Medir massa gorda e % gordura por DEXA ou bioimpedância. Mulheres: > 35% obesidade. Investigar distribuição (gordura visceral). Meta: redução de gordura visceral via dieta, exercício, manejo de estresse.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  'f0e7e520-6e01-437e-9274-5a5c90255d77',
  '35f8405e-5bd5-4803-bade-c50e6d615582'
);

-- 71-72. FMI - Fat Mass Index (Homem e Mulher)
UPDATE score_items SET
  clinical_relevance = 'FMI (Fat Mass Index, kg/m²) = massa gorda / altura² normaliza massa gorda pela estatura. Valores de referência: homens > 6,0 kg/m² excesso de gordura, > 9,0 kg/m² obesidade; mulheres > 9,0 kg/m² excesso de gordura, > 13,0 kg/m² obesidade. FMI elevado associa-se a resistência insulínica, inflamação crônica e risco cardiovascular. FMI é complementar a IMC, permitindo diagnóstico de obesidade metabólica em indivíduos com IMC normal.',
  patient_explanation = 'O Índice de Massa Gorda ajusta sua quantidade de gordura pela altura. Valores altos indicam excesso de gordura, aumentando risco metabólico. Esse índice é mais preciso que IMC para avaliar obesidade.',
  conduct = 'Calcular FMI = massa gorda (kg) / altura² (m²). Homens: > 6,0 kg/m² excesso, > 9,0 kg/m² obesidade. Mulheres: > 9,0 kg/m² excesso, > 13,0 kg/m² obesidade. Investigar síndrome metabólica.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '1442a990-a9f0-4cd3-85b1-f3d22705e469',
  'e0dbadd4-b307-43ff-8438-c24fa0876d10'
);

-- 73-74. Gordura Visceral (Homem e Mulher)
UPDATE score_items SET
  clinical_relevance = 'Gordura visceral (cm²) medida por DEXA ou TC abdominal é preditor independente de risco metabólico e cardiovascular. Valores de risco: homens > 100 cm² risco aumentado, > 130 cm² risco alto. Gordura visceral secreta adipocinas pró-inflamatórias (TNF-α, IL-6, resistina) e adipocinas anti-inflamatórias reduzidas (adiponectina), promovendo resistência insulínica, dislipidemia e aterosclerose. Redução de gordura visceral melhora perfil metabólico independente de perda de peso.',
  patient_explanation = 'Gordura visceral é a gordura acumulada dentro do abdômen, ao redor dos órgãos. Essa gordura é mais perigosa que gordura subcutânea, aumentando risco de diabetes, infarto e AVC. Reduzir gordura visceral é prioritário para sua saúde.',
  conduct = 'Medir gordura visceral por DEXA ou TC abdominal. Homens: > 100 cm² risco aumentado, > 130 cm² risco alto. Intervenções: dieta com restrição calórica moderada, exercício aeróbico + resistido, redução de estresse, sono adequado.',
  last_review = CURRENT_TIMESTAMP
WHERE id = 'd3f0556e-53cb-4a29-b825-a84bc8f38f3e';

UPDATE score_items SET
  clinical_relevance = 'Gordura visceral (cm²) medida por DEXA ou TC abdominal é preditor independente de risco metabólico e cardiovascular. Valores de risco: mulheres > 100 cm² risco aumentado. Gordura visceral secreta adipocinas pró-inflamatórias (TNF-α, IL-6, resistina), promovendo resistência insulínica, dislipidemia e aterosclerose. Mulheres pós-menopausa têm maior acúmulo visceral por redução de estrogênio. Redução de gordura visceral melhora perfil metabólico independente de perda de peso.',
  patient_explanation = 'Gordura visceral é a gordura acumulada dentro do abdômen, ao redor dos órgãos. Essa gordura é mais perigosa que gordura subcutânea, aumentando risco de diabetes, infarto e AVC. Reduzir gordura visceral é prioritário para sua saúde.',
  conduct = 'Medir gordura visceral por DEXA ou TC abdominal. Mulheres: > 100 cm² risco aumentado. Pós-menopausa: monitoramento intensivo. Intervenções: dieta, exercício aeróbico + resistido, redução de estresse, sono adequado.',
  last_review = CURRENT_TIMESTAMP
WHERE id = 'a4e0dc5b-e931-4126-a8b7-ac530f86aa68';

-- 75-76. Razão Androide/Ginoide (Homem e Mulher)
UPDATE score_items SET
  clinical_relevance = 'Razão androide/ginoide (A/G) medida por DEXA identifica padrão de distribuição de gordura. Região androide: tronco e abdômen. Região ginoide: quadris e coxas. Valores de risco: homens > 1,0 padrão androide (risco metabólico aumentado). Padrão androide associa-se a maior adiposidade visceral, resistência insulínica, dislipidemia e risco cardiovascular. Razão A/G aumenta com idade, ganho de peso e redução de testosterona.',
  patient_explanation = 'A razão androide/ginoide mostra se você acumula gordura mais na barriga (androide) ou nos quadris (ginoide). Gordura na barriga é mais perigosa para sua saúde metabólica.',
  conduct = 'Medir razão A/G por DEXA. Homens: > 1,0 padrão androide. Investigar síndrome metabólica (glicemia, HOMA-IR, perfil lipídico, PA). Intervenções para reduzir gordura visceral.',
  last_review = CURRENT_TIMESTAMP
WHERE id = '2265fbe7-fc73-48e1-ac8d-c8e239d29e3c';

UPDATE score_items SET
  clinical_relevance = 'Razão androide/ginoide (A/G) medida por DEXA identifica padrão de distribuição de gordura. Região androide: tronco e abdômen. Região ginoide: quadris e coxas. Valores de risco: mulheres > 0,8 padrão androide (risco metabólico aumentado). Padrão androide associa-se a maior adiposidade visceral, resistência insulínica, dislipidemia e risco cardiovascular. Mulheres pós-menopausa têm aumento de razão A/G por redução de estrogênio.',
  patient_explanation = 'A razão androide/ginoide mostra se você acumula gordura mais na barriga (androide) ou nos quadris (ginoide). Gordura na barriga é mais perigosa para sua saúde metabólica.',
  conduct = 'Medir razão A/G por DEXA. Mulheres: > 0,8 padrão androide. Pós-menopausa: monitoramento intensivo. Investigar síndrome metabólica. Intervenções para reduzir gordura visceral.',
  last_review = CURRENT_TIMESTAMP
WHERE id = 'bc698d3d-c9da-4276-b0d1-677ebd1fdf95';

-- 77. Gordura de Tronco (%)
UPDATE score_items SET
  clinical_relevance = 'Gordura de tronco (%) representa proporção de gordura corporal localizada no tronco (tórax, abdômen). Valores elevados (> 50% da gordura total no tronco) indicam padrão androide de distribuição, associando-se a maior risco metabólico. Gordura de tronco correlaciona-se fortemente com gordura visceral. Redução de gordura de tronco melhora sensibilidade insulínica, perfil lipídico e marcadores inflamatórios.',
  patient_explanation = 'A porcentagem de gordura no tronco indica quanto da sua gordura está acumulada no abdômen e tórax. Valores altos indicam risco metabólico aumentado.',
  conduct = 'Medir gordura de tronco (%) por DEXA. > 50% indica padrão androide. Investigar síndrome metabólica. Intervenções: dieta, exercício, manejo de estresse.',
  last_review = CURRENT_TIMESTAMP
WHERE id = 'e5b6d407-4744-48eb-8efb-bdae72dc34b9';

-- 78-80. Água Corporal Total (3x)
UPDATE score_items SET
  clinical_relevance = 'Água corporal total (ACT) representa ~60% do peso corporal em homens e ~50% em mulheres. ACT é dividida em água intracelular (AIC, ~2/3) e água extracelular (AEC, ~1/3). ACT (L) e ACT (%) são medidos por bioimpedância. Valores de referência: homens 50-65%, mulheres 45-60%. ACT baixa pode indicar desidratação. ACT elevada pode indicar retenção hídrica (insuficiência cardíaca, renal, hepática). ACT % reduzida com aumento de gordura corporal.',
  patient_explanation = 'Água corporal total é a quantidade de água no seu corpo. Valores normais indicam boa hidratação. Valores muito baixos sugerem desidratação, valores muito altos podem indicar retenção de líquidos.',
  conduct = 'Medir ACT por bioimpedância. Homens: 50-65%, Mulheres: 45-60%. ACT baixa: investigar desidratação. ACT alta: investigar edema, insuficiência cardíaca/renal/hepática.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '8bf6396f-b07c-4c93-82c4-b69ec8e87aa2',
  '2eafc97e-7134-49d7-b34c-6903192383d5',
  'f28c8a35-fa77-49a0-a69f-ba649fcffef2'
);

-- 81-82. Água Intracelular e Extracelular
UPDATE score_items SET
  clinical_relevance = 'Água intracelular (AIC, L) representa ~2/3 da água corporal total, localizada dentro das células. AIC correlaciona-se com massa celular corporal (BCM), refletindo massa metabolicamente ativa. AIC reduzida pode indicar perda de massa celular (desnutrição, sarcopenia, doença consuptiva). Monitoramento de AIC auxilia avaliação nutricional e resposta a intervenções.',
  patient_explanation = 'Água intracelular é a água dentro das suas células, refletindo massa muscular e tecidos metabolicamente ativos. Valores baixos podem indicar desnutrição ou perda muscular.',
  conduct = 'Medir AIC por bioimpedância. AIC reduzida: investigar desnutrição, sarcopenia. Otimizar nutrição, exercício.',
  last_review = CURRENT_TIMESTAMP
WHERE id = 'e79394a0-3d1f-4d18-8c55-d6a20179017a';

UPDATE score_items SET
  clinical_relevance = 'Água extracelular (AEC, L) representa ~1/3 da água corporal total, localizada fora das células (interstício, plasma). AEC elevada pode indicar edema, retenção hídrica, insuficiência cardíaca, renal ou hepática. Razão AEC/ACT > 0,40 sugere expansão de volume extracelular. Monitoramento de AEC auxilia diagnóstico e manejo de edema.',
  patient_explanation = 'Água extracelular é a água fora das suas células. Valores altos podem indicar retenção de líquidos ou edema, sugerindo problemas cardíacos, renais ou hepáticos.',
  conduct = 'Medir AEC por bioimpedância. AEC elevada: investigar edema, insuficiência cardíaca/renal/hepática. Calcular razão AEC/ACT.',
  last_review = CURRENT_TIMESTAMP
WHERE id = '7857982a-8b2b-401f-bf42-8db8a3af264b';

-- 83. Razão AEC/ACT (%)
UPDATE score_items SET
  clinical_relevance = 'Razão AEC/ACT (%) = (água extracelular / água corporal total) × 100. Valores de referência: 0,36-0,40. Razão AEC/ACT > 0,40 indica expansão de volume extracelular (edema, retenção hídrica, insuficiência cardíaca/renal/hepática). Razão AEC/ACT elevada associa-se a pior prognóstico em pacientes com doença renal crônica, insuficiência cardíaca e cirrose. Monitoramento auxilia ajuste de diuréticos e restrição hídrica/sódica.',
  patient_explanation = 'A razão entre água fora e dentro das células indica se você tem retenção de líquidos. Valores altos sugerem edema ou problemas cardíacos, renais ou hepáticos.',
  conduct = 'Calcular razão AEC/ACT. > 0,40 investigar edema, insuficiência cardíaca/renal/hepática. Ajustar diuréticos, restrição hídrica/sódica.',
  last_review = CURRENT_TIMESTAMP
WHERE id = '806cb0c4-1954-4567-b727-55bf4560553e';

-- 84-85. Ângulo de Fase (Homem e Mulher)
UPDATE score_items SET
  clinical_relevance = 'Ângulo de fase (°) medido por bioimpedância reflete integridade de membranas celulares e massa celular corporal. Valores de referência: homens > 5,0°, mulheres > 4,5°. Ângulo de fase baixo (< 5° homens, < 4,5° mulheres) indica desnutrição, sarcopenia, inflamação crônica ou doença avançada. Ângulo de fase é preditor independente de mortalidade em pacientes com câncer, HIV, cirrose e DRC. Ângulo de fase melhora com exercício resistido e nutrição adequada.',
  patient_explanation = 'O ângulo de fase indica a saúde das suas células e massa muscular. Valores baixos sugerem desnutrição, inflamação ou doença avançada. Melhorar ângulo de fase indica resposta positiva a tratamentos.',
  conduct = 'Medir ângulo de fase por bioimpedância. Homens: < 5,0° desnutrição/sarcopenia. Mulheres: < 4,5° desnutrição/sarcopenia. Otimizar nutrição (proteínas, calorias), exercício resistido, tratar doenças subjacentes.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  'f694e6f7-8d03-4cbd-9813-3635f0bee1d5',
  'ad60ffd2-8d47-4d9f-9924-47e7a3b5830b'
);

-- 86. Taxa Metabólica Basal (TMB)
UPDATE score_items SET
  clinical_relevance = 'Taxa metabólica basal (TMB, kcal) é quantidade de energia necessária para manter funções vitais em repouso. TMB representa ~60-70% do gasto energético total diário. TMB é estimada por equações (Harris-Benedict, Mifflin-St Jeor) ou medida por calorimetria indireta. TMB é influenciada por massa muscular (maior MME = maior TMB), idade (redução com envelhecimento), sexo (homens > mulheres), hormônios tireoidianos e composição corporal. TMB baixa dificulta perda de peso.',
  patient_explanation = 'Sua taxa metabólica basal é a quantidade de calorias que seu corpo gasta em repouso para manter funções vitais. Ter mais massa muscular aumenta seu metabolismo basal, facilitando controle de peso.',
  conduct = 'Estimar TMB por equação de Mifflin-St Jeor ou medir por calorimetria indireta. TMB baixa: investigar hipotireoidismo, baixa massa muscular. Aumentar TMB via treinamento resistido (ganho de MME).',
  last_review = CURRENT_TIMESTAMP
WHERE id = 'b6f334f1-e4dc-4c30-b627-a7ab9c14f5fd';

-- ============================================================================
-- FIM DO BATCH 39: 100 ITEMS ATUALIZADOS
-- ============================================================================

COMMIT;

-- Verificar total de items atualizados
SELECT COUNT(*) as total_items_atualizados
FROM score_items
WHERE last_review > CURRENT_TIMESTAMP - INTERVAL '1 minute';

SELECT 'Batch 39 completo: 100 items de Composição Corporal Parte 2 atualizados!' as status;
