-- seed-site-question.sql — preenche site_question (pergunta leiga, voz de paciente) para os
-- 538 itens de anamnese (lab_test_code IS NULL). Autoria item a item: linguagem de paciente,
-- siglas médicas explicadas. Preserva as perguntas já existentes (não sobrescreve as ~74 que já
-- estavam preenchidas / no ar no site público). Idempotente por id.
--
-- Dev: docker compose exec -T db psql -U plenya_user -d plenya_db -f /tmp/seed-site-question.sql
-- (banco direto; nunca API). Prod depois, via dry-run-em-prod com rollback.

BEGIN;

-- =========================================================================================
-- GRUPO 1 — OBJETIVOS
-- =========================================================================================
-- Objetivos iniciais (text — enquadrados como metas do paciente)
UPDATE score_items SET site_question = 'Controlar o diabetes (açúcar alto no sangue) está entre os seus objetivos? Conte o que gostaria de alcançar.' WHERE id = '019c4fad-ba45-71d0-b86d-23b1a37f1310';
UPDATE score_items SET site_question = 'Controlar uma doença autoimune (quando as defesas do corpo passam a atacar o próprio organismo) é um objetivo seu? Qual e como está hoje?' WHERE id = '019c4fad-ba4a-721b-859a-cf7c646e71ca';
UPDATE score_items SET site_question = 'Controlar a pressão alta está entre os seus objetivos? Conte como ela anda atualmente.' WHERE id = '019c4fad-ba4a-75c5-bb6b-3bc87d163ddb';
UPDATE score_items SET site_question = 'Ter mais disposição e energia no dia a dia é um objetivo seu? Descreva como anda a sua energia.' WHERE id = '019c4fad-ba4a-7c22-bdf3-c38e22d6a21c';
UPDATE score_items SET site_question = 'Emagrecer está entre os seus objetivos? Quanto gostaria de perder e por quê?' WHERE id = '019c4fad-ba4a-7c13-8cfa-7545d0c17407';
UPDATE score_items SET site_question = 'Ganhar força muscular é um objetivo seu? Conte o que espera melhorar.' WHERE id = '019c4fad-ba4a-7855-b13f-d338dcbfedca';
UPDATE score_items SET site_question = 'Lidar melhor com o estresse está entre os seus objetivos? Descreva como o estresse afeta você.' WHERE id = '019c4fad-ba4a-77d0-96d0-852f3b2cc87c';
UPDATE score_items SET site_question = 'Aliviar sintomas da perimenopausa (a transição antes da menopausa, quando os hormônios começam a cair) é um objetivo seu? O que mais incomoda?' WHERE id = '019c4fad-ba4a-7294-80a2-2bc7434ee2db';
UPDATE score_items SET site_question = 'Melhorar a sua vida sexual é um objetivo seu? Conte, no seu tempo, o que gostaria de resolver.' WHERE id = '019c4fad-ba4a-7d23-bc98-da5ed288ae7a';
UPDATE score_items SET site_question = 'Melhorar a sua alimentação está entre os seus objetivos? O que gostaria de mudar?' WHERE id = '019c4fad-ba4a-7d8b-9feb-e07c5d24a4e2';
UPDATE score_items SET site_question = 'Melhorar os seus exames de sangue é um objetivo seu? Há algum resultado que te preocupa?' WHERE id = '019c4fad-ba4a-76d8-94b0-bfd152e2e101';
UPDATE score_items SET site_question = 'Melhorar a memória e a concentração está entre os seus objetivos? Conte o que tem notado.' WHERE id = '019c4fad-ba4b-7c71-b17a-afb2d6a0532d';
UPDATE score_items SET site_question = 'Cuidar melhor da sua saúde emocional é um objetivo seu? Descreva como tem se sentido.' WHERE id = '019c4fad-ba4b-7b7d-b137-403c20eb15bf';
UPDATE score_items SET site_question = 'Melhorar o seu desempenho nos exercícios está entre os seus objetivos? O que gostaria de evoluir?' WHERE id = '019c4fad-ba4b-77e0-805a-cc028635b929';
UPDATE score_items SET site_question = 'Dormir melhor é um objetivo seu? Conte como anda o seu sono.' WHERE id = '019c4fad-ba4b-727e-ac7f-3bf722c77501';
-- Percepção de futuro (text)
UPDATE score_items SET site_question = 'Como você se imagina daqui a 6 meses em relação à sua saúde? Que metas tem para esse período?' WHERE id = 'c77cedd3-2800-7db3-aa91-68e188fa8864';
UPDATE score_items SET site_question = 'Como gostaria de estar de saúde daqui a 5 anos? Quais objetivos tem para esse prazo?' WHERE id = 'c77cedd3-2800-734f-862f-6434c4a8522c';
UPDATE score_items SET site_question = 'Quando pensa na sua saúde daqui a 10 anos, o que imagina e deseja para essa fase da vida?' WHERE id = 'c77cedd3-2800-7a0a-9f2c-fdc5ebbc2220';
-- Adesão (level_choice)
UPDATE score_items SET site_question = 'O quanto você costuma conseguir seguir, no dia a dia, um plano de saúde combinado (alimentação, exercícios, remédios)?' WHERE id = 'c77cedd3-2800-759a-bdea-45cd69d48dad';

-- =========================================================================================
-- GRUPO 2 — ALIMENTAÇÃO
-- =========================================================================================
-- Histórico
UPDATE score_items SET site_question = 'Você foi amamentado(a) no peito quando bebê? Até qual idade, se souber?' WHERE id = 'c77cedd3-2800-7b7a-a4c8-1dd10988a8a5';
UPDATE score_items SET site_question = 'O que você sabe sobre a gravidez da sua mãe e o seu pré-natal (a saúde dela e os cuidados antes de você nascer)?' WHERE id = '019bf31d-2ef0-77c5-b1a0-3b307e938d1e';
UPDATE score_items SET site_question = 'Como era a sua alimentação na adolescência (dos 12 aos 18 anos)?' WHERE id = '019bf31d-2ef0-73fe-a925-7f368f20e2f6';
UPDATE score_items SET site_question = 'Como foi a sua introdução alimentar quando bebê (quando começou a comer além do leite, por volta dos 6 meses)?' WHERE id = '019bf31d-2ef0-737a-bbee-6f85105ac8dc';
UPDATE score_items SET site_question = 'Você sabe como era a alimentação dos seus pais antes e durante a gravidez?' WHERE id = '019bf31d-2ef0-7659-a742-fbbcb741c6db';
UPDATE score_items SET site_question = 'Quais foram as piores fases da sua alimentação na vida adulta (muito ultraprocessado, doces, dietas radicais)? Quando e por quê?' WHERE id = '019bf31d-2ef0-7565-83ad-f6ac3335460d';
UPDATE score_items SET site_question = 'Como era a sua alimentação na infância (primeiros anos de vida)?' WHERE id = '019c500a-c35b-7f35-85a5-d935b36b2970';
UPDATE score_items SET site_question = 'Como era a sua alimentação até os 6 anos de idade, se você souber?' WHERE id = '019bf31d-2ef0-757f-bfc9-8028fbdf23f4';
UPDATE score_items SET site_question = 'Conte como era a sua alimentação durante a adolescência (dos 12 aos 18 anos).' WHERE id = 'c77cedd3-2800-7b74-9c6f-4fb1ea0ced62';
UPDATE score_items SET site_question = 'Durante a gravidez, a sua mãe teve diabetes, pressão alta, colesterol alterado ou excesso de peso, se souber?' WHERE id = '019bf31d-2ef0-7cc0-a851-8741fda0f082';
UPDATE score_items SET site_question = 'Você já fez acompanhamento com nutricionista? Como foi essa experiência?' WHERE id = 'c77cedd3-2800-7508-909f-34cb11c23e54';
UPDATE score_items SET site_question = 'Como era a sua alimentação na fase escolar (dos 6 aos 12 anos)?' WHERE id = '019bf31d-2ef0-73b5-b3b5-2ac72127eea0';
UPDATE score_items SET site_question = 'Como foi a sua alimentação ao longo da vida adulta, dos 18 anos até cerca de 6 meses atrás?' WHERE id = '019c500b-a814-7e44-a740-2a5ce14e60e0';
UPDATE score_items SET site_question = 'Na época em que você foi concebido(a), o seu pai tinha diabetes, pressão alta, colesterol alterado ou excesso de peso, se souber?' WHERE id = '019bf31d-2ef0-7b08-9b9b-a6cc8f6e5f25';
UPDATE score_items SET site_question = 'Você tem alguma intolerância alimentar (como à lactose, do leite, ou ao glúten, do trigo)? Quais?' WHERE id = '019bf31d-2ef0-7752-b366-5c63f4730811';
UPDATE score_items SET site_question = 'Você sente desconforto (gases, inchaço, diarreia) ao consumir leite e derivados?' WHERE id = '019bf31d-2ef0-77d9-aaa8-2ca4fad39f47';
UPDATE score_items SET site_question = 'Você sente desconforto ao comer alimentos com glúten (trigo, pães, massas)?' WHERE id = 'c77cedd3-2800-7b5f-9435-08411b954ffe';
UPDATE score_items SET site_question = 'Você tem reação à proteína do leite de vaca (diferente da intolerância à lactose)?' WHERE id = '019c1a2b-a36f-7664-a20d-37666632c3ba';
UPDATE score_items SET site_question = 'Você sente sintomas (dor de cabeça, vermelhidão, coceira) ao consumir queijos curados, vinho ou fermentados, que são ricos em histamina?' WHERE id = 'c77cedd3-2800-722f-8a95-5d4f2202fa49';
UPDATE score_items SET site_question = 'Você tem alguma alergia alimentar (leite, ovo, amendoim, trigo, frutos do mar)? Qual e que tipo de reação?' WHERE id = 'c77cedd3-2800-7e1a-aaa5-c618035a4918';
UPDATE score_items SET site_question = 'Há algum alimento que você evita ou não pode comer por motivo de saúde? Quais e por quê?' WHERE id = 'c77cedd3-2800-7e17-8ebb-585c69447a7c';
UPDATE score_items SET site_question = 'Quais são os 5 alimentos que você mais gosta de comer?' WHERE id = '019bf31d-2ef0-71f1-917b-ec74d5c19dfd';
-- Atual (últimos 6 meses)
UPDATE score_items SET site_question = 'Como é a sua alimentação nos últimos 6 meses, num dia comum?' WHERE id = '019c534a-afc3-70c4-82e0-bfde4b5b8f93';
UPDATE score_items SET site_question = 'Quanta água você costuma beber por dia?' WHERE id = '019c5355-750a-7080-9ac8-602d7c202626';
UPDATE score_items SET site_question = 'Quantas porções de fruta você come por dia, em média?' WHERE id = '019c5375-2d18-7371-a707-8a320f56a635';
UPDATE score_items SET site_question = 'Você come livremente, sem seguir nenhuma dieta ou plano alimentar específico?' WHERE id = 'c77cedd3-2800-72e4-9d17-2f997d743550';
UPDATE score_items SET site_question = 'Você segue uma dieta de baixa histamina (evitando fermentados, queijos curados e vinho)?' WHERE id = '019c534c-a38b-7b52-b4e8-0888c8a17b85';
UPDATE score_items SET site_question = 'Você bebe bebidas alcoólicas? Com que frequência e de que tipo?' WHERE id = '019c5357-9157-7169-a6a7-e1d336feb015';
UPDATE score_items SET site_question = 'Você segue uma dieta carnívora (basicamente só alimentos de origem animal)?' WHERE id = 'c77cedd3-2800-711c-8e1f-b4526da1e7ca';
UPDATE score_items SET site_question = 'Com que frequência você consome açúcar (no café, em doces, biscoitos e industrializados)?' WHERE id = '019c537b-d2e4-7a83-ae5a-e86b1f53c5ed';
UPDATE score_items SET site_question = 'Com que frequência você toma sucos (naturais ou de caixinha)?' WHERE id = '019c535c-c656-7eeb-89c6-d80b78cf1f8d';
UPDATE score_items SET site_question = 'Você usa adoçantes no lugar do açúcar? Com que frequência?' WHERE id = '019c537d-1145-74aa-a8bb-ab28ac3b56fa';
UPDATE score_items SET site_question = 'Você segue uma dieta cetogênica (muito pouco carboidrato e bastante gordura)?' WHERE id = 'c77cedd3-2800-7f39-89b7-5b249a004d08';
UPDATE score_items SET site_question = 'Quantas xícaras de café ou de chás com cafeína (preto, verde, mate) você toma por dia?' WHERE id = '019c5360-48a5-7ee6-8124-d7580d4f0282';
UPDATE score_items SET site_question = 'Como está o seu consumo de proteínas (carnes, ovos, peixe, laticínios) no dia a dia?' WHERE id = '019c5378-9893-70f2-a25f-fd10d9afdbf0';
UPDATE score_items SET site_question = 'Você sente que come mais do que o seu corpo gasta no dia (alimentação muito calórica)?' WHERE id = '019c534c-0bf6-7d6b-9c60-d650fa5f846f';
UPDATE score_items SET site_question = 'Você segue uma dieta low carb (com pouco carboidrato: pães, massas, arroz, açúcar)?' WHERE id = 'c77cedd3-2800-70de-bec4-0b5534637c78';
UPDATE score_items SET site_question = 'Você sabe, mais ou menos, quantas calorias consome por dia?' WHERE id = '019c537e-9b24-7650-836e-7153987d2658';
UPDATE score_items SET site_question = 'Você segue a dieta Low FODMAP (que reduz certos carboidratos que o intestino tem dificuldade de absorver)?' WHERE id = '019c534b-a12b-727c-a251-693f5d84d467';
UPDATE score_items SET site_question = 'Você segue a dieta mediterrânea (rica em azeite, peixes, vegetais, grãos e frutas)?' WHERE id = '019c534b-cdc0-7f24-993f-84d8a1f93fb2';
UPDATE score_items SET site_question = 'Você segue uma dieta sem glúten (sem trigo, cevada e centeio)?' WHERE id = '019c5349-99e6-79f6-8e5d-14051a0929e0';
UPDATE score_items SET site_question = 'Quem prepara as suas refeições no dia a dia (você, alguém da família, delivery, restaurante)?' WHERE id = '019bf31d-2ef0-74a7-9916-958c11dd788e';
UPDATE score_items SET site_question = 'Você segue uma dieta sem lactose (sem o açúcar do leite e derivados)?' WHERE id = '019c534b-4ad4-7564-b8f3-e13286121e42';
UPDATE score_items SET site_question = 'Que líquidos você bebe ao longo do dia e em que quantidade (água, sucos, refrigerantes, café)?' WHERE id = 'c77cedd3-2800-71a8-8555-e461cfd36c12';
UPDATE score_items SET site_question = 'Você segue uma dieta vegana (sem nenhum alimento de origem animal)?' WHERE id = '019bf31d-2ef0-7cc5-b0c3-e6a5def6a3c3';
UPDATE score_items SET site_question = 'Você segue uma dieta vegetariana (sem carnes)?' WHERE id = '019bf31d-2ef0-7362-8f98-36268261719b';
UPDATE score_items SET site_question = 'De modo geral, quais alimentos fazem parte do seu dia a dia nos últimos 6 meses?' WHERE id = '019c5374-7426-7761-8653-9a90cfda4a2e';
UPDATE score_items SET site_question = 'Descreva tudo o que você costuma comer e beber num dia típico, com horários e quantidades.' WHERE id = '019bf31d-2ef0-7f2a-acdc-e5866873d264';
UPDATE score_items SET site_question = 'O quanto você tem conseguido seguir as orientações alimentares prescritas, nos últimos tempos?' WHERE id = 'c77cedd3-2800-70ef-abfd-b3cdd23b2ced';
UPDATE score_items SET site_question = 'Como é a alimentação das pessoas que moram com você (cônjuge, filhos)?' WHERE id = '019bf31d-2ef0-733f-b1ed-6a410cf4936d';
UPDATE score_items SET site_question = 'Quais suplementos você usa atualmente? Informe o nome ou a marca, a dose e a frequência.' WHERE id = '019bf31d-2ef0-7090-b843-99752e3c622d';
UPDATE score_items SET site_question = 'As pessoas que moram com você seguem o mesmo tipo de alimentação que a sua?' WHERE id = '019bf31d-2ef0-7fc2-9d5c-a7d275642ead';
UPDATE score_items SET site_question = 'Você tem algum suplemento prescrito? Está conseguindo usar conforme a orientação?' WHERE id = '019bf31d-2ef0-753a-b39e-0c9dd0c68816';

-- =========================================================================================
-- GRUPO 3 — MOVIMENTO E ATIVIDADE FÍSICA
-- =========================================================================================
-- Histórico
UPDATE score_items SET site_question = 'Como você se movimentava e brincava na infância (até os 12 anos)? Era uma criança ativa?' WHERE id = 'c77cedd3-2800-77ce-8296-6d488eefc152';
UPDATE score_items SET site_question = 'Você praticava esportes ou exercícios na adolescência? Conte como era.' WHERE id = '019c1f9f-a6f0-7228-bef1-4ed504c694ec';
UPDATE score_items SET site_question = 'Você tem ou já teve o hábito de fazer atividades ao ar livre (esportes, caminhadas, sol)? Conte.' WHERE id = '019c1f9f-a6f3-76d3-b12a-6cfe695f0b0a';
UPDATE score_items SET site_question = 'Na infância, com que frequência e intensidade você praticava atividades físicas e esportes?' WHERE id = 'c77cedd3-2800-7513-8469-dc83e1795dca';
UPDATE score_items SET site_question = 'Ao longo da vida, com que frequência e intensidade você fez atividade física (caminhar, nadar, pedalar, musculação)?' WHERE id = '019c9b96-3c22-747a-81d5-f20d7cc58a1d';
UPDATE score_items SET site_question = 'Na infância, com que frequência e intensidade você fazia exercícios (correr, nadar, andar de bicicleta, esportes)?' WHERE id = '019c9b77-6a92-7ba4-bd0e-0bcccf486e6e';
UPDATE score_items SET site_question = 'Com que frequência e intensidade você pratica exercício físico planejado (aquele que acelera os batimentos)?' WHERE id = '019c9b96-9c21-7ce9-aa7d-6a77f8ab2450';
UPDATE score_items SET site_question = 'Você já praticou esporte de competição? Com que frequência e intensidade?' WHERE id = '019c9b96-b40d-757c-97ad-681886cdcd98';
UPDATE score_items SET site_question = 'Como foi a sua rotina de exercícios ao longo da vida adulta (o histórico, não o que faz hoje)?' WHERE id = '019c1f9f-a6f2-7776-abe6-00e1a899f026';
UPDATE score_items SET site_question = 'Na infância, você praticou esporte de competição (com treinos e campeonatos)? Com que frequência?' WHERE id = '019c9b7b-c8f5-7993-af88-360bf4bd779e';
UPDATE score_items SET site_question = 'Na infância, você costumava brincar e se movimentar ao ar livre (praças, quintais, parques)?' WHERE id = 'c77cedd3-2800-76fa-89d5-e4b5917dc239';
UPDATE score_items SET site_question = 'Quais foram as suas melhores fases em forma física? Que esportes praticava, em que idade e por quanto tempo?' WHERE id = 'c77cedd3-2800-71d9-aa16-6047f9ffb0c8';
UPDATE score_items SET site_question = 'Quais foram os períodos em que você ficou mais parado(a), sem se exercitar? Quando e por quê?' WHERE id = 'c77cedd3-2800-7469-b178-16c97cb0c25b';
UPDATE score_items SET site_question = 'Quais tipos de esporte ou exercício você mais gosta de praticar?' WHERE id = 'c77cedd3-2800-74d3-af10-b4210f1ef00a';
UPDATE score_items SET site_question = 'Quais tipos de exercício você não gosta ou evita fazer?' WHERE id = 'c77cedd3-2800-7a2d-803b-7e39bddcb1c0';
UPDATE score_items SET site_question = 'Você já fez alguma cirurgia por causa de esporte ou exercício (joelho, ombro, tornozelo)?' WHERE id = 'c77cedd3-2800-7ad2-8233-b4346812377b';
UPDATE score_items SET site_question = 'Você tem alguma limitação física ou médica que atrapalha ou impede a prática de exercícios? Qual?' WHERE id = 'c77cedd3-2800-7ccd-a370-e461d8980883';
UPDATE score_items SET site_question = 'Na sua família (pais, irmãos), as pessoas costumam ou costumavam praticar exercícios?' WHERE id = 'c77cedd3-2800-716e-924a-f761b5bd82df';
UPDATE score_items SET site_question = 'Na infância e na adolescência, você passava bastante tempo se movimentando ao ar livre?' WHERE id = '019c1f9f-a6f2-7406-b4d9-ba3a4d29fb34';
UPDATE score_items SET site_question = 'Na adolescência, com que frequência e intensidade você praticava atividade física?' WHERE id = '019c9b87-393e-73e9-b168-4e4a5d2ac725';
UPDATE score_items SET site_question = 'Na adolescência, com que frequência e intensidade você fazia exercícios?' WHERE id = '019c9b8b-c57c-74b8-a59f-0df411447bcf';
UPDATE score_items SET site_question = 'Na adolescência, você praticou esporte de competição (treinos regulares e campeonatos)? Com que frequência?' WHERE id = '019c9b8d-8234-735f-9ccd-8d442f49d2d2';
-- Atual
UPDATE score_items SET site_question = 'Como você distribui os exercícios na semana (só um tipo ou variando entre força, fôlego e mobilidade)?' WHERE id = 'c77cedd3-2800-7d37-8e88-8ff5483e6ced';
UPDATE score_items SET site_question = 'Em que horários do dia você costuma se exercitar?' WHERE id = 'c77cedd3-2800-7f06-86a6-9119be8518af';
UPDATE score_items SET site_question = 'Onde e como você faz as suas atividades físicas (academia, rua, casa; sozinho ou com acompanhamento)?' WHERE id = 'c77cedd3-2800-7b8a-a445-311a0f633963';
UPDATE score_items SET site_question = 'Quem se exercita junto com você atualmente (você sozinho, parceiro, família)?' WHERE id = 'c77cedd3-2800-75f3-9aca-0bddd66c8069';
UPDATE score_items SET site_question = 'Você usa algum suplemento antes, durante ou depois dos treinos? Quais?' WHERE id = 'c77cedd3-2800-7f15-940c-a4d47cabb237';
UPDATE score_items SET site_question = 'Você tem alguma prova ou desafio físico marcado para os próximos 6 meses (corrida, ciclismo, natação)?' WHERE id = 'c77cedd3-2800-75bb-ad6a-2f22013e4aca';
UPDATE score_items SET site_question = 'As pessoas próximas a você (família e amigos) praticam exercícios e apoiam a sua prática?' WHERE id = 'c77cedd3-2800-7f9c-9bf6-5c053962ca97';
-- Testes práticos (resultado que o paciente informa; tabela de referência varia por idade/sexo)
UPDATE score_items SET site_question = 'Por quanto tempo você consegue sustentar a posição de prancha (abdômen contraído)?' WHERE id = '019cbe7a-213c-7a96-a3b5-997ac4fcb32c';
UPDATE score_items SET site_question = 'Quantas repetições de abdominal você consegue fazer seguidas, sem parar?' WHERE id = '019c9bb1-db5f-7486-8033-c53296c26c92';
UPDATE score_items SET site_question = 'Quantas flexões de braço no chão você consegue fazer seguidas, sem parar?' WHERE id = '019c9bbf-1f29-7e87-83f4-0dd27520b8dd';
UPDATE score_items SET site_question = 'Quantos burpees (agachar, deitar no chão e saltar) você consegue fazer no tempo do teste?' WHERE id = '019c9bc5-17a3-70dd-b0b3-2ded87054c9a';
UPDATE score_items SET site_question = 'Quantas repetições de abdominal você consegue fazer seguidas, sem parar?' WHERE id = '019c9bb8-ff9a-766f-97f3-56faecff965e';
UPDATE score_items SET site_question = 'Quantas flexões de braço no chão você consegue fazer seguidas, sem parar?' WHERE id = '019c9bc0-a317-736e-8d62-c12f3987e181';
UPDATE score_items SET site_question = 'Quantas flexões de braço no chão você consegue fazer seguidas, sem parar?' WHERE id = '019c9bc0-a319-7b53-b599-b27f33d67503';
UPDATE score_items SET site_question = 'Por quanto tempo você consegue sustentar a posição de prancha (abdômen contraído)?' WHERE id = '019cbe7a-2145-7931-a655-5afec9787eee';
UPDATE score_items SET site_question = 'Quantos burpees (agachar, deitar no chão e saltar) você consegue fazer no tempo do teste?' WHERE id = '019c9bc5-17a5-723f-b9d7-bb23a42dcc67';
UPDATE score_items SET site_question = 'Quantas repetições de abdominal você consegue fazer seguidas, sem parar?' WHERE id = '019c9bb8-ffa6-7053-8c43-431c02fb2379';
UPDATE score_items SET site_question = 'Quantas repetições de abdominal você consegue fazer seguidas, sem parar?' WHERE id = '019c9bb8-ffa6-7283-b178-79bda697293b';
UPDATE score_items SET site_question = 'Quantos burpees (agachar, deitar no chão e saltar) você consegue fazer no tempo do teste?' WHERE id = '019c9bc2-fe21-7dc3-9ad1-2ebedfe1c7ff';
UPDATE score_items SET site_question = 'Quantas flexões de braço no chão você consegue fazer seguidas, sem parar?' WHERE id = '019c9bc0-a31a-71a3-98da-441001600e08';
UPDATE score_items SET site_question = 'Quantas repetições de abdominal você consegue fazer seguidas, sem parar?' WHERE id = '019c9bb8-ffa6-707f-aeef-a0508733d57b';
UPDATE score_items SET site_question = 'Quantas flexões de braço no chão você consegue fazer seguidas, sem parar?' WHERE id = '019c9bc0-a31a-7859-b556-d657b1e06fb1';
UPDATE score_items SET site_question = 'Quantas repetições de abdominal você consegue fazer seguidas, sem parar?' WHERE id = '019c9bb8-ffa7-7888-84ea-6b084755db27';
UPDATE score_items SET site_question = 'Quantas flexões de braço no chão você consegue fazer seguidas, sem parar?' WHERE id = '019c9bc0-a31b-7099-a6d3-d4299afb8eba';
UPDATE score_items SET site_question = 'Quantas repetições de abdominal você consegue fazer seguidas, sem parar?' WHERE id = '019c9bb8-ffa7-74f9-b03c-4d9580cc4281';
UPDATE score_items SET site_question = 'Quantas flexões de braço no chão você consegue fazer seguidas, sem parar?' WHERE id = '019c9bc0-a31b-7c8d-9117-b9f577acf6b2';
UPDATE score_items SET site_question = 'Quantas flexões de braço no chão você consegue fazer seguidas, sem parar?' WHERE id = '019c9bc0-a31b-74bf-8552-e565e32e2716';
UPDATE score_items SET site_question = 'Quantas repetições de abdominal você consegue fazer seguidas, sem parar?' WHERE id = '019c9bb8-ffa7-7576-9227-7fcd19e28af8';
UPDATE score_items SET site_question = 'Quantas flexões de braço no chão você consegue fazer seguidas, sem parar?' WHERE id = '019c9bc0-a31c-7c79-8504-12761feb6450';
UPDATE score_items SET site_question = 'Quantas repetições de abdominal você consegue fazer seguidas, sem parar?' WHERE id = '019c9bb8-ffa9-70f4-836b-10cff120d509';
UPDATE score_items SET site_question = 'Quantas flexões de braço no chão você consegue fazer seguidas, sem parar?' WHERE id = '019c9bc0-a31c-70f0-a806-4f7d4eb474de';
UPDATE score_items SET site_question = 'Quantas repetições de abdominal você consegue fazer seguidas, sem parar?' WHERE id = '019c9bb8-ffa9-7e41-b424-cbe0421eb8c4';
UPDATE score_items SET site_question = 'Quantas flexões de braço no chão você consegue fazer seguidas, sem parar?' WHERE id = '019c9bc0-a31c-7917-b243-642d3ef3cdaf';
UPDATE score_items SET site_question = 'Quantas repetições de abdominal você consegue fazer seguidas, sem parar?' WHERE id = '019c9bb8-ffa9-7b9f-991c-7a5adad9b744';
UPDATE score_items SET site_question = 'Quantas repetições de abdominal você consegue fazer seguidas, sem parar?' WHERE id = '019c9bb8-ffaa-712f-9147-c5fc38a0374e';
UPDATE score_items SET site_question = 'Quantas flexões de braço no chão você consegue fazer seguidas, sem parar?' WHERE id = '019c9bc0-a31d-775b-944c-a4731c861404';
UPDATE score_items SET site_question = 'Quantas repetições de abdominal você consegue fazer seguidas, sem parar?' WHERE id = '019c9bb8-ffaa-779b-8fd0-bdcc640fb16e';
UPDATE score_items SET site_question = 'Quantas flexões de braço no chão você consegue fazer seguidas, sem parar?' WHERE id = '019c9bc0-a31d-73e4-9c15-f566651d55f8';

-- =========================================================================================
-- GRUPO 4 — SONO (histórico + início do atual)
-- =========================================================================================
UPDATE score_items SET site_question = 'Como era o seu sono na infância (dos 0 aos 12 anos)? Dormia bem?' WHERE id = '019c1646-a023-7ba7-b004-97a01e452de9';
UPDATE score_items SET site_question = 'Com que idade você deixou as fraldas (aprendeu a controlar o xixi e o cocô)?' WHERE id = 'c77cedd3-2800-753d-a826-239d2d2bcdbe';
UPDATE score_items SET site_question = 'Você fez xixi na cama dormindo depois dos 5 anos de idade (o que os médicos chamam de enurese)?' WHERE id = 'c77cedd3-2800-7ae9-9784-67595e1a596b';
UPDATE score_items SET site_question = 'Como era a qualidade do seu sono na adolescência?' WHERE id = 'c77cedd3-2800-7a3b-b643-8e229f0aae71';
UPDATE score_items SET site_question = 'Como era o seu sono na adolescência (dos 12 aos 18 anos)?' WHERE id = '019c1647-1212-7573-8459-dc293bfdd04c';
UPDATE score_items SET site_question = 'Como foi o seu sono ao longo da vida adulta, até cerca de 6 meses atrás?' WHERE id = 'c77cedd3-2800-7010-9fa6-80bb6bf67621';
UPDATE score_items SET site_question = 'Você já usou ou usa remédios ou suplementos para dormir? Quais?' WHERE id = '019bf31d-2ef0-77ed-8332-6f2e5242850f';
UPDATE score_items SET site_question = 'Na sua família há problemas de sono (insônia, ronco, apneia)? Quem?' WHERE id = '019bf31d-2ef0-72d7-a8ea-2d6b8ac3da07';
UPDATE score_items SET site_question = 'Quais foram as melhores fases do seu sono (idade, quantas horas dormia, horários)?' WHERE id = 'c77cedd3-2800-76a2-a48c-6a02ec07734f';
UPDATE score_items SET site_question = 'Quais foram as piores fases do seu sono? O que atrapalhava?' WHERE id = 'c77cedd3-2800-7d81-bb48-3c21e00eacb2';
UPDATE score_items SET site_question = 'A que horas você costuma acordar nos dias normais?' WHERE id = '019c5396-5a85-70ca-a228-101168d0e9e4';
UPDATE score_items SET site_question = 'Você consome café ou chás com cafeína à noite? Com que frequência?' WHERE id = 'c77cedd3-2800-7ab2-977c-a6a39cedc3fa';
-- Sono > Atual (continuação)
UPDATE score_items SET site_question = 'Você costuma ir dormir sempre no mesmo horário, inclusive nos fins de semana?' WHERE id = '019c53a3-1947-749e-9cf9-722da80c9d5a';
UPDATE score_items SET site_question = 'Com que frequência você acorda durante a noite?' WHERE id = '019bf31d-2ef0-7efb-b6ec-8c7965bf4374';
UPDATE score_items SET site_question = 'Quando você acorda à noite, quanto tempo costuma levar para voltar a dormir?' WHERE id = '019bf31d-2ef0-7bc5-8cd7-4673f4f13134';
UPDATE score_items SET site_question = 'Você acorda durante a noite? Quantas vezes e por quê?' WHERE id = 'c77cedd3-2800-7a83-befe-9b059fdfa924';
UPDATE score_items SET site_question = 'O que mais costuma te acordar à noite (ir ao banheiro, ronco, calor, ansiedade, dor)?' WHERE id = '019c53b1-d062-728e-8618-d96687bdf99f';
UPDATE score_items SET site_question = 'Você usa algum aparelho ou aplicativo para monitorar o sono (smartwatch, anel, app de celular)?' WHERE id = '019bf31d-2ef0-7815-a7d3-62844e1d3dae';
UPDATE score_items SET site_question = 'Como é a sua cama para dormir (colchão e conforto)?' WHERE id = 'c77cedd3-2800-7aac-9cf8-ea99accfeb68';
UPDATE score_items SET site_question = 'Como é o ambiente onde você dorme (luz, barulho, temperatura)?' WHERE id = 'c77cedd3-2800-7147-972b-e88d4a9b2acc';
UPDATE score_items SET site_question = 'O seu colchão é confortável e adequado para você?' WHERE id = '019c1649-2558-7553-bfbe-3acf1f3236b0';
UPDATE score_items SET site_question = 'O que e quando você costuma comer à noite, antes de dormir?' WHERE id = 'c77cedd3-2800-715c-a182-eba24bbcfd43';
UPDATE score_items SET site_question = 'Quais hábitos você tem na hora de dormir (rotina, telas, horários)?' WHERE id = '019bf31d-2ef0-7ca9-82cf-f55897d2a862';
UPDATE score_items SET site_question = 'O seu travesseiro é confortável e dá bom apoio ao pescoço?' WHERE id = '019bf31d-2ef0-7b18-bc13-059cf0770c8e';
UPDATE score_items SET site_question = 'A roupa que você usa para dormir é confortável e adequada à temperatura?' WHERE id = '019bf31d-2ef0-7eb3-a838-0d959754a6c6';
UPDATE score_items SET site_question = 'Você usa atualmente algum remédio ou suplemento para dormir? Qual?' WHERE id = '019bf31d-2ef0-741d-b04c-a535615a997e';
UPDATE score_items SET site_question = 'As suas roupas de cama (lençóis e cobertas) são confortáveis e adequadas à temperatura?' WHERE id = '019bf31d-2ef0-7d0b-a934-934044a17207';
UPDATE score_items SET site_question = 'Você tem algum sintoma à noite (suor, câimbras, ronco, agitação, despertares)?' WHERE id = '019bf31d-2ef0-7d60-8cd4-21d8a3d87027';
UPDATE score_items SET site_question = 'De modo geral, o seu quarto é um bom ambiente para dormir?' WHERE id = '019bf31d-2ef0-7240-b3ba-5926f65c1bdd';
UPDATE score_items SET site_question = 'A pessoa que divide o quarto com você atrapalha o seu sono (ronco, movimento, celular, horários diferentes)?' WHERE id = '019bf31d-2ef0-7799-abdb-ce0be9930287';
UPDATE score_items SET site_question = 'O seu quarto fica bem escuro na hora de dormir?' WHERE id = '019bf31d-2ef0-794d-aeb9-ac5fbe5b0542';
UPDATE score_items SET site_question = 'Entra claridade de fora no seu quarto à noite (postes, rua, amanhecer)?' WHERE id = '019bf31d-2ef0-764f-9b07-b9a13e5b6a62';
UPDATE score_items SET site_question = 'Você costuma usar telas (TV, celular, tablet) na cama antes de dormir?' WHERE id = '019bf31d-2ef0-7ae9-850f-dac4fb0da126';
UPDATE score_items SET site_question = 'Há barulho no seu quarto durante a noite?' WHERE id = 'c77cedd3-2800-7c46-901c-16e38005e8d9';
UPDATE score_items SET site_question = 'Há odores no seu quarto que ajudam ou atrapalham o seu sono?' WHERE id = '019bf31d-2ef0-7a1c-8338-652e22f41fb3';
UPDATE score_items SET site_question = 'A temperatura do seu quarto é agradável para dormir?' WHERE id = '019c53b4-69f2-7131-a846-1fce769a709d';
UPDATE score_items SET site_question = 'Você mantém aparelhos eletrônicos ligados perto da cama (celular, roteador Wi-Fi)?' WHERE id = 'c77cedd3-2800-711d-996c-e4b6ae6b71e1';
UPDATE score_items SET site_question = 'Você range ou aperta os dentes durante o sono (acorda com dor ou tensão na mandíbula)?' WHERE id = '019c164e-5254-7f7b-8857-fa984531e276';
UPDATE score_items SET site_question = 'Você tem pesadelos que te acordam à noite? Com que frequência?' WHERE id = '019c164e-8f2e-7e6d-b940-931aef234a2b';
UPDATE score_items SET site_question = 'Você sua muito à noite enquanto dorme?' WHERE id = '019c164f-04f8-74a8-97e7-6c25bdbaacd1';
UPDATE score_items SET site_question = 'Você tem dificuldade para pegar no sono quando vai para a cama?' WHERE id = '019c5392-7d84-7de2-9afd-90f03ae9f5df';

-- =========================================================================================
-- GRUPO 5 — COGNIÇÃO
-- =========================================================================================
-- Histórico
UPDATE score_items SET site_question = 'Algum parente próximo (pais, avós) teve depressão?' WHERE id = '019c5507-896e-710c-b304-ec327a7839ab';
UPDATE score_items SET site_question = 'Como foi o seu desenvolvimento na infância (sentar, andar, falar no tempo esperado)?' WHERE id = 'c77cedd3-2800-7cb1-8925-8cdcb6eca5f3';
UPDATE score_items SET site_question = 'Algum parente próximo (pais, avós) teve transtorno bipolar?' WHERE id = '019c5507-896e-7e16-86f9-9536c3f6c0be';
UPDATE score_items SET site_question = 'Como foi o seu desempenho na escola (aprendizado, atenção, notas)?' WHERE id = 'c77cedd3-2800-7d9b-84fa-5b9bf26d4693';
UPDATE score_items SET site_question = 'Algum parente próximo (pais, avós) teve ansiedade?' WHERE id = '019c5507-896e-7f98-ba24-5f7734682050';
UPDATE score_items SET site_question = 'Como foi a sua memória ao longo da vida?' WHERE id = 'c77cedd3-2800-727e-99be-46aebe3f9419';
UPDATE score_items SET site_question = 'Algum parente próximo (pais, avós) foi diagnosticado com esquizofrenia?' WHERE id = '019c5507-896e-79a6-a052-4610f2bc6132';
UPDATE score_items SET site_question = 'Como foi a sua disposição e energia ao longo da vida?' WHERE id = 'c77cedd3-2800-716b-8f22-bb290378358e';
UPDATE score_items SET site_question = 'Como foram o seu foco, a concentração e a facilidade de aprender ao longo da vida?' WHERE id = 'c77cedd3-2800-7a4a-b3a2-cc48ae8931e2';
UPDATE score_items SET site_question = 'Você já passou por períodos de depressão (tristeza profunda e desânimo por semanas) ao longo da vida?' WHERE id = 'c77cedd3-2800-7363-a4ae-d9a9c845eb8f';
UPDATE score_items SET site_question = 'Você tem diagnóstico de Transtorno do Espectro Autista (TEA)?' WHERE id = 'c77cedd3-2800-7285-8c1d-982d66c9b790';
UPDATE score_items SET site_question = 'Você tem diagnóstico de Transtorno de Déficit de Atenção e Hiperatividade (TDAH)?' WHERE id = '019c54ee-0191-70f3-a4a9-6b32b8391e94';
UPDATE score_items SET site_question = 'Como era o convívio com a sua família na infância e na adolescência?' WHERE id = 'c77cedd3-2800-7538-ba65-42299ee272ba';
UPDATE score_items SET site_question = 'Como eram as suas amizades e a vida social na infância e na adolescência?' WHERE id = '019c54f4-4112-750b-8d90-077ad52d23aa';
UPDATE score_items SET site_question = 'Em quais fases da vida a sua mente esteve mais afiada (boa memória, foco e clareza)?' WHERE id = 'c77cedd3-2800-7668-83e9-1a906d19e69d';
UPDATE score_items SET site_question = 'Em quais fases você sentiu a mente pior (memória, foco ou clareza abaixo do normal)? O que estava acontecendo?' WHERE id = '019bf31d-2ef0-7237-9e94-666bc07646ff';
UPDATE score_items SET site_question = 'Você já usou remédios ou suplementos para memória, foco, humor ou disposição mental? Quais?' WHERE id = '019bf31d-2ef0-737d-b74f-520c73096980';
UPDATE score_items SET site_question = 'Há casos de demência ou Alzheimer entre os seus pais e avós?' WHERE id = '019c5507-65cb-793f-a2fa-25d598e7f203';
-- Atual
UPDATE score_items SET site_question = 'Como está a sua energia e vontade de participar de momentos com a família?' WHERE id = 'c77cedd3-2800-7ba6-ab4c-a4b6f7de1229';
UPDATE score_items SET site_question = 'Teste de memória: depois de ouvir cinco palavras, quantas você consegue repetir na hora?' WHERE id = 'c77cedd3-2800-7e6c-a76f-596811b9b4e0';
UPDATE score_items SET site_question = 'Como você avalia a sua memória e o seu raciocínio hoje, no dia a dia?' WHERE id = '019bf31d-2ef0-7929-813b-9ead82399476';
UPDATE score_items SET site_question = 'Como está a sua vontade e energia para estar com outras pessoas e participar de eventos sociais?' WHERE id = '019c550e-dc4c-7edf-8f1e-273359d547e2';
UPDATE score_items SET site_question = 'Nas últimas 2 semanas, com que frequência você se sentiu nervoso, ansioso ou no limite? (questionário GAD-7)' WHERE id = 'c77cedd3-2800-78a5-a272-01c1e573fcc0';
UPDATE score_items SET site_question = 'Teste de atenção: quantos números em sequência você consegue repetir na ordem em que ouviu?' WHERE id = '019bf31d-2ef0-7fa3-a163-84fde09b062f';
UPDATE score_items SET site_question = 'Como está a sua energia e foco para o trabalho no dia a dia?' WHERE id = '019c550e-a320-7e31-8ca9-5c03364d3940';
UPDATE score_items SET site_question = 'Teste de memória de trabalho: quantos números você consegue repetir na ordem inversa à que ouviu?' WHERE id = '019bf31d-2ef0-7934-b23d-98bd66bc03ee';
UPDATE score_items SET site_question = 'Como está a sua vontade e energia para se exercitar?' WHERE id = '019c550f-14a5-7681-891a-9ac8dd96b9eb';
UPDATE score_items SET site_question = 'Teste de memória: depois de alguns minutos, quantas das cinco palavras você ainda lembra?' WHERE id = 'c77cedd3-2800-7085-b998-daab565ddd1c';
UPDATE score_items SET site_question = 'O quanto você sente vontade de cochilar durante o dia em situações comuns? (escala de sonolência de Epworth)' WHERE id = 'c77cedd3-2800-7556-8439-afa139a9bae3';
UPDATE score_items SET site_question = 'Como está a sua disposição e energia para as atividades do dia a dia?' WHERE id = 'c77cedd3-2800-7ad9-bf6e-c4266e3d5df5';
UPDATE score_items SET site_question = 'Como estão o seu foco, a concentração e a capacidade de aprender atualmente?' WHERE id = 'c77cedd3-2800-7ccd-8dfd-cff55da653fd';
UPDATE score_items SET site_question = 'Você tem sentido tristeza, desânimo ou ansiedade com frequência? Descreva.' WHERE id = '019bf31d-2ef0-72fd-8568-8da14cf20af8';
UPDATE score_items SET site_question = 'Para mulheres: você percebe mudanças no humor, energia ou memória conforme o ciclo menstrual ou a proximidade da menopausa?' WHERE id = '019bf31d-2ef0-7fa0-964b-652831c3cd95';
UPDATE score_items SET site_question = 'Você percebeu piora de memória ou disposição depois de alguma infecção forte, doença grave ou trauma?' WHERE id = '019c5514-f0ee-7b08-aa75-64ecf0d9c2cd';
UPDATE score_items SET site_question = 'Você usa atualmente algum medicamento para memória, concentração ou humor? Qual?' WHERE id = '019bf31d-2ef0-7e06-b67d-edfaa498ba25';

-- =========================================================================================
-- GRUPO 6 — STRESS
-- =========================================================================================
UPDATE score_items SET site_question = 'Você passou por situações de estresse intenso ou traumas ao longo da vida (infância, adolescência, fase adulta)?' WHERE id = 'c77cedd3-2800-7276-8c09-8a1263923cde';
UPDATE score_items SET site_question = 'Quando você está muito estressado, o seu corpo reage de alguma forma (intestino, respiração, pele, dores)?' WHERE id = 'c77cedd3-2800-734c-831d-845fdaea4ad7';
UPDATE score_items SET site_question = 'O que você já usou ou faz para aliviar o estresse? O que funciona melhor para você?' WHERE id = 'c77cedd3-2800-7836-af1c-1170baeadee7';
UPDATE score_items SET site_question = 'Você usa alguma forma de relaxar ou aliviar o estresse no dia a dia (exercício, meditação, lazer)? Com que frequência?' WHERE id = 'c77cedd3-2800-7a0b-8f29-a7a79c27b7f3';
UPDATE score_items SET site_question = 'Quais sintomas você sente hoje e relaciona ao estresse (cansaço, insônia, irritação, tensão muscular)?' WHERE id = 'c77cedd3-2800-7e5f-9e0a-6769f842be0f';

-- =========================================================================================
-- GRUPO 7 — VIDA SEXUAL
-- =========================================================================================
-- Histórico
UPDATE score_items SET site_question = 'Com que idade começaram os sinais da puberdade (pelos; primeira menstruação nas mulheres; primeira ejaculação nos homens)?' WHERE id = 'c77cedd3-2800-7dcb-b0c8-53a4a878857f';
UPDATE score_items SET site_question = 'Você viveu, no passado, alguma situação de abuso ou trauma que gostaria de registrar? Responda no seu tempo.' WHERE id = 'c77cedd3-2800-7bc8-a56b-41bec7497c68';
UPDATE score_items SET site_question = 'Como você descreveria a sua vida sexual de forma geral (desejo, frequência, satisfação)?' WHERE id = 'c77cedd3-2800-72b0-a9d4-26bf05843214';
UPDATE score_items SET site_question = 'Em quais fases da vida você teve mais desejo e melhor desempenho sexual?' WHERE id = 'c77cedd3-2800-7181-a820-d836c68cdd2b';
UPDATE score_items SET site_question = 'Em quais fases você percebeu queda no desejo ou no desempenho sexual? O que estava acontecendo?' WHERE id = '019bf31d-2ef0-7f4d-b8a7-fc83c68d28ee';
UPDATE score_items SET site_question = 'Conte o seu histórico reprodutivo (número de filhos, dificuldade para engravidar, gestações, abortos, tipo de parto).' WHERE id = 'c77cedd3-2800-75f3-a287-89c9b36f13b5';
UPDATE score_items SET site_question = 'Você já usou hormônios (como testosterona ou anticoncepcional) ou remédios para o desempenho sexual (para disfunção erétil)?' WHERE id = 'c77cedd3-2800-7e48-a9af-1a71a90c337f';
-- Atual
UPDATE score_items SET site_question = 'Como estão o seu desejo e o seu desempenho sexual atualmente?' WHERE id = 'c77cedd3-2800-795b-8203-e7ba11b5ed8b';
UPDATE score_items SET site_question = 'O que você percebe que melhora a sua vida sexual?' WHERE id = 'c77cedd3-2800-7a18-a61e-67eed6a5cc6f';
UPDATE score_items SET site_question = 'O que você percebe que piora a sua vida sexual?' WHERE id = 'c77cedd3-2800-7859-8bc3-44c9deee8abe';
UPDATE score_items SET site_question = 'Como é o seu ciclo menstrual (duração, intensidade do fluxo, sintomas, uso de anticoncepcional)?' WHERE id = 'c77cedd3-2800-7166-80cd-65e5a7e956fc';
UPDATE score_items SET site_question = 'Você está sem usar anticoncepcional hormonal (pílula, anel, adesivo, injeção)?' WHERE id = 'c77cedd3-2800-784b-b3bc-516e0c8887b4';
UPDATE score_items SET site_question = 'Você usa anticoncepcional hormonal (pílula, DIU hormonal)? Qual?' WHERE id = 'c77cedd3-2800-77eb-885c-476883a4208f';
UPDATE score_items SET site_question = 'Você já está na pós-menopausa (mais de um ano sem menstruar)?' WHERE id = 'c77cedd3-2800-7467-9c04-2dc347979727';
UPDATE score_items SET site_question = 'Você usou recentemente algum hormônio (testosterona, estrogênio, progesterona)?' WHERE id = 'c77cedd3-2800-7831-b6e6-8b53c3be1642';
UPDATE score_items SET site_question = 'Você usou recentemente algum remédio para disfunção erétil (os que ajudam na ereção)?' WHERE id = 'c77cedd3-2800-770b-80f0-418fb3a941d6';
UPDATE score_items SET site_question = 'Você usou recentemente algum outro remédio ou suplemento para o desejo ou o desempenho sexual?' WHERE id = '019bf31d-2ef0-7839-96b6-08260315e87c';
UPDATE score_items SET site_question = 'Você gostaria de relatar como está a sua função sexual hoje (ereção, lubrificação, orgasmo, satisfação)?' WHERE id = 'c77cedd3-2800-7fb6-99ad-dd12f4e9e919';
UPDATE score_items SET site_question = 'Questionário sobre a sua vida sexual: desejo, excitação, lubrificação, orgasmo e satisfação (escala ASEX).' WHERE id = 'c77cedd3-2800-7d12-9f9e-5dd3fc1e6df0';
UPDATE score_items SET site_question = 'Questionário de 5 perguntas sobre a função erétil (IIEF-5).' WHERE id = 'c77cedd3-2800-7e5e-b7c0-bd312f6a7ed4';
UPDATE score_items SET site_question = 'Questionário sobre a função sexual feminina: desejo, excitação, lubrificação, orgasmo e satisfação (escala FSFI).' WHERE id = 'c77cedd3-2800-70e0-a1b5-faeb028289f2';

-- =========================================================================================
-- GRUPO 8 — COMPOSIÇÃO CORPORAL
-- =========================================================================================
-- Histórico
UPDATE score_items SET site_question = 'Você sabe como estavam a saúde e o peso dos seus pais quando você foi concebido e durante a gravidez?' WHERE id = '019bf31d-2ef0-77be-95b8-a42a36673de1';
UPDATE score_items SET site_question = 'Você sabe quanto pesou ao nascer, se nasceu no tempo certo ou prematuro, e se houve complicações no parto?' WHERE id = '019bf31d-2ef0-7331-9a47-72fa2d52b5e2';
UPDATE score_items SET site_question = 'Como foi o seu crescimento na infância (peso, altura, se era magro ou acima do peso)?' WHERE id = '019bf31d-2ef0-7b57-8352-ea32449169ba';
UPDATE score_items SET site_question = 'Como variaram o seu peso e a sua altura na adolescência?' WHERE id = 'c77cedd3-2800-7ce8-bcd4-4e6f6a38f5cf';
UPDATE score_items SET site_question = 'Você já fez avaliações de composição corporal (gordura e músculo) na vida adulta? Como evoluíram?' WHERE id = '019bf31d-2ef0-7959-a33d-32f3646c1a87';
UPDATE score_items SET site_question = 'Em quais fases da vida você teve mais peso ou gordura? O que estava acontecendo?' WHERE id = '019bf31d-2ef0-7a4c-bc23-87c68c41ee89';
UPDATE score_items SET site_question = 'Em quais fases o seu corpo esteve na melhor forma (boa proporção de músculo e gordura)?' WHERE id = '019bf31d-2ef0-7bbc-8d27-2dd89f9a99ad';
UPDATE score_items SET site_question = 'Quais tratamentos você já fez para mudar o corpo (dietas, remédios, hormônios, cirurgias)?' WHERE id = '019bf31d-2ef0-756f-b854-ed707c817a3d';
UPDATE score_items SET site_question = 'Na sua família (pais, irmãos, avós), as pessoas tendem a ter excesso de peso ou boa massa muscular?' WHERE id = '019bf31d-2ef0-74a7-8449-5441c26ea829';
-- Atual
UPDATE score_items SET site_question = 'Você tem alguma avaliação recente de composição corporal (DEXA, bioimpedância, dobras)? Quais os resultados?' WHERE id = '019bf31d-2ef0-75bb-8a3a-19e4463de7a5';
UPDATE score_items SET site_question = 'Você está fazendo algum tratamento para mudar o corpo atualmente (dieta, remédio, hormônio)?' WHERE id = '019bf31d-2ef0-7cff-9970-7c437076adc7';
UPDATE score_items SET site_question = 'O quanto você está satisfeito com o seu corpo hoje, e o que gostaria de mudar?' WHERE id = '019bf31d-2ef0-78f5-a54c-298195ff0588';
-- Medidas Objetivas
UPDATE score_items SET site_question = 'Qual é o seu peso atual, em quilos?' WHERE id = '019bf31d-2ef0-74c5-8d99-0f355f1aa7cc';
UPDATE score_items SET site_question = 'Qual é a sua altura, em centímetros?' WHERE id = '019bf31d-2ef0-71e4-a845-cdfaeedbb599';
UPDATE score_items SET site_question = 'Medida calculada a partir da sua cintura e altura que estima a gordura corporal (BRI).' WHERE id = '019cbe90-f927-74d1-be7b-3eb87777c2e8';
UPDATE score_items SET site_question = 'Você sabe qual é a sua taxa metabólica basal (as calorias que o corpo gasta em repouso)?' WHERE id = '019bf31d-2ef0-7c30-b627-a7ab9c14f5fd';
UPDATE score_items SET site_question = 'Qual é a medida do seu quadril, em centímetros (na parte mais larga)?' WHERE id = '019bf31d-2ef0-7a53-a71a-4eff3e0cc6ba';
UPDATE score_items SET site_question = 'Relação entre as medidas da sua cintura e do seu quadril (mostra onde o corpo acumula gordura).' WHERE id = '019bf31d-2ef0-7903-ac22-1d8b5b47179a';
UPDATE score_items SET site_question = 'Relação entre as medidas da sua cintura e do seu quadril (mostra onde o corpo acumula gordura).' WHERE id = '019bf31d-2ef0-7b9c-846e-edd6a3c8277f';
UPDATE score_items SET site_question = 'Relação entre a medida do seu pescoço e a sua altura (avalia acúmulo de gordura).' WHERE id = '019bf31d-2ef0-78c1-be7a-245df89c200c';
UPDATE score_items SET site_question = 'Relação entre a medida do seu pescoço e a sua altura (avalia acúmulo de gordura).' WHERE id = '019bf31d-2ef0-7532-8c73-c4a96f31e2fc';
UPDATE score_items SET site_question = 'Qual é a medida do seu braço direito contraído, em centímetros?' WHERE id = 'c77cedd3-2800-7f52-b0eb-c1daf5da33fa';
UPDATE score_items SET site_question = 'Qual é a medida do seu braço esquerdo contraído, em centímetros?' WHERE id = 'c77cedd3-2800-761a-b634-7965cfecd6b7';
UPDATE score_items SET site_question = 'Medida do seu braço não dominante, relaxado, em centímetros (estima a massa muscular).' WHERE id = '019bf31d-2ef0-7cd4-8b06-75f70e2fb3eb';
UPDATE score_items SET site_question = 'Medida do seu braço não dominante, relaxado, em centímetros (estima a massa muscular).' WHERE id = '019bf31d-2ef0-7092-a4c3-1430ec21f18a';
UPDATE score_items SET site_question = 'Qual é a medida da sua coxa, em centímetros?' WHERE id = '019bf31d-2ef0-7cfc-9239-046fd1f882e2';
UPDATE score_items SET site_question = 'Qual é a medida da sua coxa, em centímetros?' WHERE id = '019bf31d-2ef0-7e20-9b32-bb5cd7d4a3d1';
UPDATE score_items SET site_question = 'Qual é a medida da sua panturrilha (batata da perna), em centímetros?' WHERE id = '019bf31d-2ef0-7587-be08-9963c0520bf0';
UPDATE score_items SET site_question = 'Qual é a medida da sua panturrilha (batata da perna), em centímetros?' WHERE id = '019bf31d-2ef0-7936-b83b-284a0f8c84eb';
UPDATE score_items SET site_question = 'Você sabe qual é a sua massa de gordura total, em quilos?' WHERE id = '019bf31d-2ef0-737e-9274-5a5c90255d77';
UPDATE score_items SET site_question = 'Índice que mede a quantidade de gordura do corpo em relação à altura (FMI).' WHERE id = '019bf31d-2ef0-7cd3-85b1-f3d22705e469';
UPDATE score_items SET site_question = 'Índice que mede a quantidade de gordura do corpo em relação à altura (FMI).' WHERE id = '019bf31d-2ef0-73ff-8438-c24fa0876d10';
UPDATE score_items SET site_question = 'Você sabe qual é o seu percentual de gordura no tronco (barriga, tórax, costas)?' WHERE id = '019bf31d-2ef0-78eb-8efb-bdae72dc34b9';
UPDATE score_items SET site_question = 'Medida que compara a gordura da barriga com a dos quadris (razão androide/ginoide).' WHERE id = '019bf31d-2ef0-78e1-ac8d-c8e239d29e3c';
UPDATE score_items SET site_question = 'Medida que compara a gordura da barriga com a dos quadris (razão androide/ginoide).' WHERE id = '019bf31d-2ef0-7276-b0d1-677ebd1fdf95';
UPDATE score_items SET site_question = 'Você sabe qual é a sua massa muscular, em quilos?' WHERE id = '019bf31d-2ef0-7207-85ee-168b1c14e35e';
UPDATE score_items SET site_question = 'Você sabe qual proporção do seu peso é músculo?' WHERE id = '019bf31d-2ef0-7c44-b268-53b72f75adff';
UPDATE score_items SET site_question = 'Você sabe o seu índice de massa muscular (músculo dos braços e pernas em relação à altura)?' WHERE id = '019bf31d-2ef0-7964-ac29-63bf9c224bef';
UPDATE score_items SET site_question = 'Você sabe qual é a massa muscular dos seus braços e pernas, em quilos?' WHERE id = '019bf31d-2ef0-7e12-9247-4e29ce1e7185';
UPDATE score_items SET site_question = 'Você sabe qual é a quantidade total de água no seu corpo, em litros?' WHERE id = '019bf31d-2ef0-7c93-82c4-b69ec8e87aa2';
UPDATE score_items SET site_question = 'Percentual do seu peso que é água (água corporal total).' WHERE id = '019bf31d-2ef0-79d7-b34c-6903192383d5';
UPDATE score_items SET site_question = 'Percentual do seu peso que é água (água corporal total).' WHERE id = '019bf31d-2ef0-79a0-a69f-ba649fcffef2';
UPDATE score_items SET site_question = 'Você sabe qual é a quantidade de água dentro das suas células (água intracelular), em litros?' WHERE id = '019bf31d-2ef0-7d18-8c55-d6a20179017a';
UPDATE score_items SET site_question = 'Você sabe qual é a quantidade de água fora das células do seu corpo (água extracelular), em litros?' WHERE id = '019bf31d-2ef0-701f-bf42-8db8a3af264b';
UPDATE score_items SET site_question = 'Medida de como a água do corpo se distribui dentro e fora das células (razão entre água extracelular e água total).' WHERE id = '019bf31d-2ef0-7567-b727-55bf4560553e';
UPDATE score_items SET site_question = 'Medida da bioimpedância que avalia a saúde das suas células por dentro (ângulo de fase).' WHERE id = '019bf31d-2ef0-7cbd-9813-3635f0bee1d5';
UPDATE score_items SET site_question = 'Medida da bioimpedância que avalia a saúde das suas células por dentro (ângulo de fase).' WHERE id = '019bf31d-2ef0-7d9f-9924-47e7a3b5830b';

-- =========================================================================================
-- GRUPO 9 — HISTÓRICO DE DOENÇAS
-- =========================================================================================
-- Doenças crônicas
UPDATE score_items SET site_question = 'Você tem diagnóstico confirmado de diabetes? Há quanto tempo e como está o controle?' WHERE id = '019bf31d-2ef0-71c2-b7af-1e9880f62583';
UPDATE score_items SET site_question = 'Você tem ou já teve doença renal crônica (funcionamento reduzido dos rins)?' WHERE id = '019bf31d-2ef0-7adb-8bd5-a242485e0446';
UPDATE score_items SET site_question = 'Você tem ou já teve alguma outra doença nos rins?' WHERE id = '019bf31d-2ef0-7a7d-815c-a56b3e3735c9';
UPDATE score_items SET site_question = 'Você tem ou já teve nefrite (inflamação nos rins)?' WHERE id = '019bf31d-2ef0-7103-bda8-a522bc357978';
UPDATE score_items SET site_question = 'Você tem ou já teve síndrome nefrótica (perda de proteína pela urina, com inchaço)?' WHERE id = '019bf31d-2ef0-7bdb-b7c8-1c239c373389';
UPDATE score_items SET site_question = 'Você tem ou já teve pedra nos rins (cálculo renal)?' WHERE id = '019bf31d-2ef0-79c2-8fc0-7dcd8dcb75bc';
UPDATE score_items SET site_question = 'Quantas infecções urinárias (ardência ou dor ao urinar) você costuma ter por ano?' WHERE id = '019bf31d-2ef0-734a-aeae-0682b3375758';
UPDATE score_items SET site_question = 'Você tem alguma infecção por vírus de longa duração (como hepatite ou herpes)? Qual?' WHERE id = '019bf31d-2ef0-7d1f-b3bd-8f8e8f731634';
UPDATE score_items SET site_question = 'Você tem diagnóstico de HIV (o vírus que afeta as defesas do corpo)?' WHERE id = '019bf31d-2ef0-74d0-837c-1ad343661dc2';
UPDATE score_items SET site_question = 'Você tem ou já teve hepatite B (uma infecção do fígado causada por vírus)?' WHERE id = '019bf31d-2ef0-7e76-b57a-c79c8e45a162';
UPDATE score_items SET site_question = 'Você tem ou já teve hepatite C (uma infecção do fígado causada por vírus)?' WHERE id = '019bf31d-2ef0-7039-8b37-f70fdec68733';
UPDATE score_items SET site_question = 'Você tem alguma doença autoimune (quando as defesas do corpo atacam o próprio organismo)? Qual?' WHERE id = '019bf31d-2ef0-7f95-b152-ddf14d0ce398';
UPDATE score_items SET site_question = 'Você tem diagnóstico de lúpus (LES, uma doença autoimune)?' WHERE id = '019bf31d-2ef0-7fcc-9a3c-886a6fb709a6';
UPDATE score_items SET site_question = 'Você tem ou já teve artrite reumatoide (doença autoimune que inflama as articulações)?' WHERE id = 'c77cedd3-2800-76bb-bbc2-81329e69c008';
UPDATE score_items SET site_question = 'Você tem ou já teve esclerodermia (doença autoimune que endurece a pele e os tecidos)?' WHERE id = '019bf31d-2ef0-7039-889b-d9b18025392e';
UPDATE score_items SET site_question = 'Você tem ou já teve doença de Crohn (uma inflamação crônica do intestino)?' WHERE id = '019bf31d-2ef0-75f2-8485-f9d26e4e1369';
UPDATE score_items SET site_question = 'Você tem ou já teve retocolite ulcerativa (RCU, inflamação crônica do intestino grosso)?' WHERE id = '019bf31d-2ef0-729d-a9df-4b53569231dc';
UPDATE score_items SET site_question = 'Você tem doença celíaca (intolerância autoimune ao glúten)?' WHERE id = '019bf31d-2ef0-702c-9375-02d892ce7440';
UPDATE score_items SET site_question = 'Você tem alguma alergia (a alimentos, remédios, pólen, poeira, picadas)? Quais?' WHERE id = 'c77cedd3-2800-7f31-8acc-136e5d63aded';
UPDATE score_items SET site_question = 'Você tem alguma outra doença crônica que ainda não mencionamos? Qual?' WHERE id = '019bf31d-2ef0-7186-b669-ed5f667eb1c8';
-- Histórico de saúde (sintomas e doenças, passado ou atual)
UPDATE score_items SET site_question = 'Quais infecções você já teve ao longo da vida (das quais se lembra)?' WHERE id = '019bf31d-2ef0-7ece-b1af-2aecca8ded54';
UPDATE score_items SET site_question = 'Você costuma ter náuseas ou vômitos? Com que frequência?' WHERE id = '019bf31d-2ef0-7314-9051-e9b12c5626eb';
UPDATE score_items SET site_question = 'Você tem ou já teve tuberculose (uma infecção dos pulmões causada por bactéria)?' WHERE id = '019bf31d-2ef0-7829-b9ee-dda47b6b55d4';
UPDATE score_items SET site_question = 'Você tem ou já teve herpes simples (feridas que voltam nos lábios ou na região genital)?' WHERE id = '019bf31d-2ef0-7f2e-b0b3-72f23851469f';
UPDATE score_items SET site_question = 'Você já teve herpes-zóster (o popular cobreiro)?' WHERE id = '019bf31d-2ef0-783f-94a6-86102fae8ca4';
UPDATE score_items SET site_question = 'Você costuma ter diarreia? Com que frequência?' WHERE id = 'c77cedd3-2800-7f9d-a7c2-a1fdb7923a06';
UPDATE score_items SET site_question = 'Você costuma ter dores ou cólicas na barriga?' WHERE id = '019bf31d-2ef0-765b-b291-f80f9dce057b';
UPDATE score_items SET site_question = 'Você já teve catapora (varicela)?' WHERE id = '019bf31d-2ef0-744f-810c-40fa736da05f';
UPDATE score_items SET site_question = 'Você costuma ter excesso de gases (flatulência)?' WHERE id = '019bf31d-2ef0-7044-b745-1313acc819d3';
UPDATE score_items SET site_question = 'Você já teve infecção por citomegalovírus (CMV)?' WHERE id = '019bf31d-2ef0-70b0-be6d-b11e371fd6ee';
UPDATE score_items SET site_question = 'Você costuma arrotar com frequência (eructação)?' WHERE id = '019bf31d-2ef0-7ce2-ae05-c01f8fef9ead';
UPDATE score_items SET site_question = 'Você já teve mononucleose (a doença do beijo, causada pelo vírus EBV)?' WHERE id = '019bf31d-2ef0-716b-8327-11c00851402e';
UPDATE score_items SET site_question = 'Você já teve dengue?' WHERE id = '019bf31d-2ef0-7ebc-8831-6a3925345344';
UPDATE score_items SET site_question = 'Você tem ou já teve hemorroidas?' WHERE id = '019bf31d-2ef0-7a23-9410-9d6034c3e141';
UPDATE score_items SET site_question = 'Você já teve zika?' WHERE id = '019bf31d-2ef0-7422-a5ee-e02540370f7a';
UPDATE score_items SET site_question = 'Você sente dor ou ardência ao urinar (disúria)?' WHERE id = '019bf31d-2ef0-749f-8305-537a8fb85d98';
UPDATE score_items SET site_question = 'Você tem ou já teve dor na parte baixa das costas (dor lombar)?' WHERE id = '019bf31d-2ef0-7d83-a6db-9571aabf9bde';
UPDATE score_items SET site_question = 'Você já teve chikungunya?' WHERE id = '019bf31d-2ef0-7b50-8794-1fb5c541e43a';
UPDATE score_items SET site_question = 'Você já teve COVID-19? Ficou com algum sintoma depois?' WHERE id = '019bf31d-2ef0-734a-aef8-dfd8a07933e8';
UPDATE score_items SET site_question = 'Você já teve sarampo?' WHERE id = '019bf31d-2ef0-7e01-8bd5-06b3ae60e8ec';
UPDATE score_items SET site_question = 'Você já teve alguma outra doença infecciosa que ainda não mencionamos?' WHERE id = '019bf31d-2ef0-70ff-b7c5-b9ae4e7eec42';
UPDATE score_items SET site_question = 'Você tem alguma doença alérgica (rinite, asma, alergia de pele)? Qual?' WHERE id = '019bf31d-2ef0-7591-983a-1ac670c3da1d';
UPDATE score_items SET site_question = 'Você tem ou já teve rinite (nariz entupido, espirros, coriza)?' WHERE id = '019bf31d-2ef0-7b1b-a647-6bf4f6458d30';
UPDATE score_items SET site_question = 'Você tem ou já teve sinusite (inflamação dos seios da face)?' WHERE id = '019bf31d-2ef0-7afd-b8ea-10aa7638fecb';
UPDATE score_items SET site_question = 'Você tem ou já teve bronquite (inflamação dos brônquios, com tosse e catarro)?' WHERE id = 'c77cedd3-2800-7707-8985-9beb91d0768c';
UPDATE score_items SET site_question = 'Você tem ou já teve urticária (placas vermelhas com muita coceira na pele)?' WHERE id = '019bf31d-2ef0-7d1a-b92b-c36dc8335f00';
UPDATE score_items SET site_question = 'Você tem alguma outra alergia que ainda não mencionamos?' WHERE id = '019bf31d-2ef0-759a-8089-fcceffcaa4a5';
UPDATE score_items SET site_question = 'Você tem ou já teve algum problema na cabeça (dores de cabeça, enxaqueca, zumbido, tontura)?' WHERE id = '019bf31d-2ef0-7a61-8855-bcbd3dd8b4fe';
UPDATE score_items SET site_question = 'Você costuma ter dor de cabeça? Com que frequência?' WHERE id = '019bf31d-2ef0-7cf5-a8e2-7be010e0ad5e';
UPDATE score_items SET site_question = 'Você costuma sentir tontura?' WHERE id = '019bf31d-2ef0-7817-8c4f-d399bafe461f';
UPDATE score_items SET site_question = 'Você costuma ouvir zumbido nos ouvidos (chiado ou apito sem som externo)?' WHERE id = '019bf31d-2ef0-7988-ac94-7be0c7725a14';
UPDATE score_items SET site_question = 'Você tem ou já teve algum problema de visão (visão dupla, embaçada, pontos no campo de visão)?' WHERE id = '019bf31d-2ef0-720d-a74a-b2c873931143';
UPDATE score_items SET site_question = 'Você sente algum outro sintoma na região da cabeça que ainda não mencionamos?' WHERE id = '019bf31d-2ef0-7d2c-b5e2-4a992060de4d';
UPDATE score_items SET site_question = 'Você tem ou já teve algum problema no peito (coração, pulmões, respiração)?' WHERE id = '019bf31d-2ef0-7763-a743-0fce76fe2686';
UPDATE score_items SET site_question = 'Você tem ou já teve algum problema nas mamas (dor, nódulo, cisto)?' WHERE id = '019bf31d-2ef0-7ad3-9340-a962ffb32c58';
UPDATE score_items SET site_question = 'Você tem ou já teve algum problema na barriga ou na digestão?' WHERE id = '019bf31d-2ef0-74d4-b3f7-a89665d3bd8f';
UPDATE score_items SET site_question = 'Você tem ou já teve algum problema nos braços ou nas pernas?' WHERE id = '019bf31d-2ef0-7604-b5ae-731c05c90af0';
UPDATE score_items SET site_question = 'Você costuma ter câimbras (contrações dolorosas, geralmente nas pernas)?' WHERE id = '019bf31d-2ef0-7a1a-a7ea-fb9d0cc21878';
UPDATE score_items SET site_question = 'Você tem dores nas articulações (juntas dos braços, pernas, mãos ou pés)?' WHERE id = '019bf31d-2ef0-7ac0-a743-d6cb400f4b27';
UPDATE score_items SET site_question = 'Você tem ou já teve artrite (inflamação nas articulações)?' WHERE id = 'c77cedd3-2800-731b-a1dd-edf619700cca';
UPDATE score_items SET site_question = 'Você tem ou já teve alguma lesão muscular (estiramento, ruptura)?' WHERE id = '019bf31d-2ef0-78ba-95bc-91e454256d84';
UPDATE score_items SET site_question = 'Você tem ou já teve lesão de ligamento ou tendão?' WHERE id = '019bf31d-2ef0-7044-9d59-b71fa222cbaa';
UPDATE score_items SET site_question = 'Você já teve alguma fratura (osso quebrado)?' WHERE id = '019bf31d-2ef0-7a31-b8c9-0f54b201dbcb';
UPDATE score_items SET site_question = 'Você tem ou já teve algum problema de pele, cabelo ou unhas?' WHERE id = '019bf31d-2ef0-7bee-92b4-81a1559562ad';
UPDATE score_items SET site_question = 'Você tem notado os fios de cabelo mais finos, quebradiços ou sem brilho?' WHERE id = '019bf31d-2ef0-7981-b73d-46aa7c71ec6b';
UPDATE score_items SET site_question = 'Você tem notado queda de cabelo acima do normal?' WHERE id = '019bf31d-2ef0-781a-8255-108158a64239';
UPDATE score_items SET site_question = 'As suas unhas andam fracas, quebradiças ou com estrias?' WHERE id = '019bf31d-2ef0-7ce8-845a-a4e176109ad2';
UPDATE score_items SET site_question = 'Você tem notado excesso ou falta de pelos no corpo?' WHERE id = 'c77cedd3-2800-7997-866e-2182f73ed8c1';
UPDATE score_items SET site_question = 'Você tem ou já teve lesões na pele (manchas, vermelhidão, coceira, acne, eczema)?' WHERE id = 'c77cedd3-2800-7bfd-9300-14d50f9475bb';
UPDATE score_items SET site_question = 'Você tem ou já teve algum problema na região genital (pênis, testículos, próstata)?' WHERE id = '019bf31d-2ef0-7ddf-a1ae-2f67c03e54a4';
UPDATE score_items SET site_question = 'Você tem ou já teve algum problema na pele que recobre a ponta do pênis (prepúcio) ou na glande?' WHERE id = '019bf31d-2ef0-7ea1-8117-3d3e30377dab';
UPDATE score_items SET site_question = 'Você tem alguma curvatura fora do normal no pênis, especialmente durante a ereção?' WHERE id = 'c77cedd3-2800-7baf-8756-f0cab6268e1b';
UPDATE score_items SET site_question = 'Você tem ou já teve algum problema na bolsa escrotal ou no epidídimo (atrás do testículo)?' WHERE id = '019bf31d-2ef0-7b42-9d37-487e91411a18';
UPDATE score_items SET site_question = 'Você tem ou já teve algum problema nos testículos (dor, nódulo, inchaço)?' WHERE id = '019bf31d-2ef0-7e35-94d7-0d232cc258ce';
UPDATE score_items SET site_question = 'Você tem ou já teve algum problema na região genital (vulva, vagina, órgãos reprodutivos)?' WHERE id = '019bf31d-2ef0-7616-8f7b-33b1c9d11279';
UPDATE score_items SET site_question = 'Você sente ressecamento, ardência ou desconforto na região genital (comum após a menopausa)?' WHERE id = '019bf31d-2ef0-7e31-b393-f2783549d874';
UPDATE score_items SET site_question = 'Você tem ou já teve perda de urina ou sensação de peso na região íntima (assoalho pélvico enfraquecido)?' WHERE id = '019bf31d-2ef0-70ee-907a-28431ce4858a';
UPDATE score_items SET site_question = 'Você tem ou já teve algum problema na vulva (parte externa da genitália feminina)?' WHERE id = '019bf31d-2ef0-7078-b942-860ee56136a4';
UPDATE score_items SET site_question = 'Você tem ou já teve alguma doença na vagina ou no colo do útero (infecções repetidas, alteração no preventivo)?' WHERE id = '019bf31d-2ef0-7f0c-b76b-806c1d6ff1fd';
-- Cirurgias já realizadas
UPDATE score_items SET site_question = 'Quais cirurgias você já realizou ao longo da vida?' WHERE id = '019bf31d-2ef0-70e0-a47d-2cef065d39e9';
UPDATE score_items SET site_question = 'Você já fez alguma cirurgia que retirou ou alterou um órgão (tireoide, útero, vesícula, entre outros)?' WHERE id = '019bf31d-2ef0-78fe-a64a-61215107d5e3';
UPDATE score_items SET site_question = 'Você já passou por amputação de algum membro (braço ou perna)?' WHERE id = 'c77cedd3-2800-7632-a1de-097fd28e1da7';
UPDATE score_items SET site_question = 'Você já fez mastectomia (retirada total ou parcial da mama)?' WHERE id = '019bf31d-2ef0-7270-8a8b-017b293ca147';
UPDATE score_items SET site_question = 'Você já fez cirurgia de retirada da próstata (prostatectomia)?' WHERE id = '019bf31d-2ef0-73a8-b1ae-3e27305e25d8';
UPDATE score_items SET site_question = 'Você já fez cirurgia de retirada da tireoide (tireoidectomia)?' WHERE id = '019bf31d-2ef0-7db2-9844-d04c37701d92';
UPDATE score_items SET site_question = 'Você já fez cirurgia de retirada do útero (histerectomia)?' WHERE id = '019bf31d-2ef0-77e9-9045-5c994cfcbf94';
UPDATE score_items SET site_question = 'Você já fez cirurgia de retirada dos ovários (ooforectomia)?' WHERE id = '019bf31d-2ef0-72f9-b93c-0f0a852d9d51';
UPDATE score_items SET site_question = 'Você já fez cirurgia de retirada de um ou dos dois testículos (orquiectomia)?' WHERE id = '019bf31d-2ef0-786a-a6b1-3e1b4ca8cc6e';
UPDATE score_items SET site_question = 'Você já fez cirurgia de retirada de um rim (nefrectomia)?' WHERE id = '019bf31d-2ef0-78e0-b6a8-2323b3896449';
UPDATE score_items SET site_question = 'Você já fez cirurgia de retirada de parte do fígado (hepatectomia parcial)?' WHERE id = '019bf31d-2ef0-749e-9e66-045319ceaaa3';
UPDATE score_items SET site_question = 'Você já fez cirurgia de retirada de parte ou de todo um pulmão?' WHERE id = '019bf31d-2ef0-7946-bf90-bc759eddb080';
UPDATE score_items SET site_question = 'Você já fez cirurgia no cérebro com abertura do crânio (craniotomia)?' WHERE id = '019bf31d-2ef0-7b91-9670-b17fab93c6e9';
UPDATE score_items SET site_question = 'Você já fez cirurgia para tratar epilepsia?' WHERE id = '019bf31d-2ef0-7a1f-ab7c-964f8c8114e6';
UPDATE score_items SET site_question = 'Você tem uma derivação (válvula com tubinho) para drenar líquido do cérebro?' WHERE id = '019bf31d-2ef0-73a2-8d2d-a11801e9f19d';
UPDATE score_items SET site_question = 'Você já fez cirurgia de retirada do baço (esplenectomia)?' WHERE id = '019bf31d-2ef0-7461-9f21-b414e3f4c425';
UPDATE score_items SET site_question = 'Você já fez cirurgia para corrigir hérnia (herniorrafia)?' WHERE id = '019bf31d-2ef0-7b62-a444-ddf312915ea0';
UPDATE score_items SET site_question = 'Você já fez cirurgia de retirada da vesícula biliar (colecistectomia)?' WHERE id = '019bf31d-2ef0-7b9c-a1eb-81ef1f2ec646';
UPDATE score_items SET site_question = 'Você já fez cirurgia de retirada de glândula paratireoide (paratireoidectomia)?' WHERE id = '019bf31d-2ef0-7fc3-b4ed-3e9c8cfd5b36';
UPDATE score_items SET site_question = 'Você já fez cirurgia de retirada de glândula suprarrenal (adrenalectomia)?' WHERE id = 'c77cedd3-2800-72ee-8b15-cda98a584c37';
UPDATE score_items SET site_question = 'Você já fez cirurgia de retirada de parte ou de todo o pâncreas (pancreatectomia)?' WHERE id = '019bf31d-2ef0-7522-94c8-c6f9299e4c59';
UPDATE score_items SET site_question = 'Você já fez cirurgia nos olhos para descolamento de retina ou glaucoma?' WHERE id = '019bf31d-2ef0-7998-b857-e28f32fe8a4e';
-- Medicamentos
UPDATE score_items SET site_question = 'Quais remédios você já usou ao longo da vida, e teve alguma reação ou resultado importante?' WHERE id = '019bf31d-2ef0-75aa-838b-0a0928acc4b1';
UPDATE score_items SET site_question = 'Quais remédios você usa atualmente (com receita, de venda livre, anticoncepcionais, suplementos)?' WHERE id = '019bf31d-2ef0-78da-9d77-4e8258d3cf8e';
UPDATE score_items SET site_question = 'Você usa ou já usou remédios para dor (anti-inflamatórios, opioides, relaxantes musculares)?' WHERE id = 'c77cedd3-2800-785b-a255-311ee0fc7bf4';
UPDATE score_items SET site_question = 'Você usa ou já usou remédios para ansiedade ou para acalmar (ansiolíticos, benzodiazepínicos)?' WHERE id = 'c77cedd3-2800-7de4-bb4d-0e886a128fb2';
UPDATE score_items SET site_question = 'Você usa ou já usou remédios para controlar o ritmo do coração (antiarrítmicos)?' WHERE id = 'c77cedd3-2800-7dec-9ee4-e8ed1335b026';
UPDATE score_items SET site_question = 'Você já usou antibióticos por longos períodos ou em muitos ciclos?' WHERE id = 'c77cedd3-2800-78a9-805d-b206544d27b9';
UPDATE score_items SET site_question = 'Você usa ou já usou remédios que afinam o sangue (anticoagulantes ou antiagregantes, como AAS)?' WHERE id = 'c77cedd3-2800-7427-842d-04ab6960c01d';
UPDATE score_items SET site_question = 'Você usa ou já usou anticoncepcional?' WHERE id = 'c77cedd3-2800-7f39-a8fd-1332fa906fa5';
UPDATE score_items SET site_question = 'Você usa ou já usou remédios para convulsão (anticonvulsivantes)?' WHERE id = 'c77cedd3-2800-7dad-baf7-f92b6f987100';
UPDATE score_items SET site_question = 'Você usa ou já usou antidepressivos?' WHERE id = 'c77cedd3-2800-7d33-86b8-0ac8be744547';
UPDATE score_items SET site_question = 'Você usa ou já usou remédios em comprimido para diabetes (como metformina)?' WHERE id = 'c77cedd3-2800-7f7f-964b-88f5c2451e79';
UPDATE score_items SET site_question = 'Você usa ou já usou remédios para osteoporose (ossos fracos)?' WHERE id = 'c77cedd3-2800-7c46-aa62-a9dbb13341dd';
UPDATE score_items SET site_question = 'Você usa ou já usou remédios para Parkinson?' WHERE id = 'c77cedd3-2800-793d-bb4e-43b4db957188';
UPDATE score_items SET site_question = 'Você usa ou já usou antipsicóticos ou estabilizantes de humor?' WHERE id = 'c77cedd3-2800-736d-a224-3d66f3a5d7c1';
UPDATE score_items SET site_question = 'Você usa antivirais de forma contínua (para HIV ou hepatite)?' WHERE id = 'c77cedd3-2800-787b-ab16-51343f66bc18';
UPDATE score_items SET site_question = 'Você usa ou já usou bombinhas para respirar (broncodilatadores ou corticoides inalatórios)?' WHERE id = '019bf31d-2ef0-7f5b-93c1-4590ed06b91f';
UPDATE score_items SET site_question = 'Você usa ou já usou corticoide por via oral ou injetável (como prednisona)?' WHERE id = '019bf31d-2ef0-7f37-99af-bba1e5d57177';
UPDATE score_items SET site_question = 'Você usa ou já usou remédios para fortalecer o coração (digitálicos, como digoxina)?' WHERE id = '019bf31d-2ef0-7982-a41f-ea8bb0eb9da4';
UPDATE score_items SET site_question = 'Você usa ou já usou remédios para disfunção erétil?' WHERE id = '019bf31d-2ef0-737c-a061-b81dddab34cd';
UPDATE score_items SET site_question = 'Você usa ou já usou estimulantes para foco ou atenção (como metilfenidato)?' WHERE id = '019bf31d-2ef0-7b72-94e7-367771f837bd';
UPDATE score_items SET site_question = 'Você usa ou já usou remédios para a tireoide?' WHERE id = '019bf31d-2ef0-7d60-aa28-180a7e672899';
UPDATE score_items SET site_question = 'Você usa ou já usou remédios para próstata aumentada (HPB)?' WHERE id = '019bf31d-2ef0-760f-baec-2680dc817e47';
UPDATE score_items SET site_question = 'Você usa ou já usou medicamentos imunobiológicos (injetáveis para doenças autoimunes ou inflamatórias)?' WHERE id = '019bf31d-2ef0-7cf2-9572-77271dc50145';
UPDATE score_items SET site_question = 'Você usa ou já usou remédios que reduzem o sistema de defesa (imunossupressores)?' WHERE id = '019bf31d-2ef0-7bff-b781-ebde75069941';
UPDATE score_items SET site_question = 'Você usa ou já usou remédios para azia ou refluxo (como omeprazol e outros prazóis)?' WHERE id = 'c77cedd3-2800-7058-9539-d53d1fed6e30';
UPDATE score_items SET site_question = 'Você usa laxantes ou antidiarreicos com frequência?' WHERE id = '019bf31d-2ef0-717c-95d3-12c4bc1ce4f7';
UPDATE score_items SET site_question = 'Você usa ou já usou quimioterapia?' WHERE id = '019bf31d-2ef0-7671-b30a-fbb314fff861';
UPDATE score_items SET site_question = 'Você usa ou já usou reposição hormonal (estrogênio ou testosterona)?' WHERE id = '019bf31d-2ef0-750e-bda7-232a3ab8e49b';
UPDATE score_items SET site_question = 'Você usa ou já usou hormônios para tratar câncer de mama (como tamoxifeno)?' WHERE id = '019bf31d-2ef0-7e4d-9c33-8b8a95f36642';
-- Hábitos e vícios nocivos
UPDATE score_items SET site_question = 'Você usa ou já usou alguma droga ilícita? Responda no seu tempo.' WHERE id = '019bf31d-2ef0-7295-958a-97935e4329ca';
UPDATE score_items SET site_question = 'Você tem o hábito de jogar ou apostar? Sente que às vezes foge do controle?' WHERE id = '019bf31d-2ef0-7fe6-ba81-35b2005375c1';
UPDATE score_items SET site_question = 'Como você descreveria os seus hábitos sexuais (frequência, comportamentos de risco)?' WHERE id = '019bf31d-2ef0-7e2d-8a49-94a50eae4437';
UPDATE score_items SET site_question = 'Você tem algum outro vício ou hábito que sente que prejudica a sua saúde (cafeína em excesso, telas, compras)?' WHERE id = '019bf31d-2ef0-77a4-8a16-8174b61dcb8a';
-- Saúde bucal
UPDATE score_items SET site_question = 'Como é o seu histórico no dentista? Você tem restaurações antigas prateadas (amálgama)?' WHERE id = '019bf31d-2ef0-7a03-b425-bb60dff938b7';
UPDATE score_items SET site_question = 'Você já perdeu dentes? Quantos, aproximadamente?' WHERE id = '019bf31d-2ef0-745c-9291-9ef38957a93d';
UPDATE score_items SET site_question = 'Você já fez tratamentos dentários mais pesados (canal, implante, extrações, obturações de amálgama)?' WHERE id = '019bf31d-2ef0-7607-9893-69d10e7914ae';
UPDATE score_items SET site_question = 'Como está a saúde da sua boca hoje (dentes, gengivas, eventuais dores)?' WHERE id = '019bf31d-2ef0-76fb-aeee-5d3115a2738d';
UPDATE score_items SET site_question = 'Você tem algum sintoma na boca (sangramento na gengiva, mau hálito, feridas, sensibilidade)?' WHERE id = '019bf31d-2ef0-787c-b063-fbe80c9216b1';
UPDATE score_items SET site_question = 'Você consegue mastigar bem os alimentos, sem dor ou dificuldade?' WHERE id = '019bf31d-2ef0-70ea-a03d-e9cb564f837a';
-- Especialistas médicos externos
UPDATE score_items SET site_question = 'Você faz acompanhamento com clínico geral ou médico de família?' WHERE id = 'b1070000-0000-7000-8000-000000000000';
UPDATE score_items SET site_question = 'Você faz acompanhamento com cardiologista (médico do coração)?' WHERE id = 'b1070000-0000-7000-8000-000000000001';
UPDATE score_items SET site_question = 'Você faz acompanhamento com endocrinologista (médico de hormônios e metabolismo)?' WHERE id = 'b1070000-0000-7000-8000-000000000002';
UPDATE score_items SET site_question = 'Você faz acompanhamento com nefrologista (médico dos rins)?' WHERE id = 'b1070000-0000-7000-8000-000000000003';
UPDATE score_items SET site_question = 'Você faz acompanhamento com ginecologista?' WHERE id = 'b1070000-0000-7000-8000-000000000004';
UPDATE score_items SET site_question = 'Você faz acompanhamento com urologista?' WHERE id = 'b1070000-0000-7000-8000-000000000005';
UPDATE score_items SET site_question = 'Você faz acompanhamento com gastroenterologista (médico do aparelho digestivo)?' WHERE id = 'b1070000-0000-7000-8000-000000000006';
UPDATE score_items SET site_question = 'Você faz acompanhamento com dermatologista (médico da pele)?' WHERE id = 'b1070000-0000-7000-8000-000000000007';
UPDATE score_items SET site_question = 'Você faz acompanhamento com oftalmologista (médico dos olhos)?' WHERE id = 'b1070000-0000-7000-8000-000000000008';
UPDATE score_items SET site_question = 'Você faz acompanhamento com psiquiatra?' WHERE id = 'b1070000-0000-7000-8000-000000000009';
UPDATE score_items SET site_question = 'Você faz acompanhamento com algum outro médico especialista? Qual?' WHERE id = 'b1070000-0000-7000-8000-00000000000a';
-- Equipe multiprofissional externa
UPDATE score_items SET site_question = 'Você faz acompanhamento com dentista?' WHERE id = 'b1070000-0000-7000-8000-00000000000b';
UPDATE score_items SET site_question = 'Você faz acompanhamento com nutricionista?' WHERE id = 'b1070000-0000-7000-8000-00000000000c';
UPDATE score_items SET site_question = 'Você faz acompanhamento com educador físico ou personal trainer?' WHERE id = 'b1070000-0000-7000-8000-00000000000d';
UPDATE score_items SET site_question = 'Você faz acompanhamento com psicólogo?' WHERE id = 'b1070000-0000-7000-8000-00000000000e';
UPDATE score_items SET site_question = 'Você faz acompanhamento com fisioterapeuta?' WHERE id = 'b1070000-0000-7000-8000-00000000000f';
UPDATE score_items SET site_question = 'Você faz acompanhamento com algum outro profissional de saúde? Qual?' WHERE id = 'b1070000-0000-7000-8000-000000000010';

-- =========================================================================================
-- GRUPO 10 — HISTÓRICO FAMILIAR DE DOENÇAS
-- =========================================================================================
UPDATE score_items SET site_question = 'Quais doenças crônicas existem na sua família próxima (pais, irmãos, avós, filhos)?' WHERE id = '019bf31d-2ef0-7d56-976a-b6339b4482a6';
UPDATE score_items SET site_question = 'Algum parente próximo tem ou teve doença nos rins (insuficiência renal, nefrite, pedras, infecções de repetição)?' WHERE id = 'c77cedd3-2800-731d-88cb-2ac2422f8e86';
UPDATE score_items SET site_question = 'Algum parente próximo tem ou teve HIV, hepatite B ou hepatite C?' WHERE id = 'c77cedd3-2800-7845-98dc-303f46922403';
UPDATE score_items SET site_question = 'Algum parente próximo tem ou teve doença autoimune (lúpus, artrite reumatoide, Crohn, asma)?' WHERE id = 'c77cedd3-2800-728e-a218-a29391e1e26f';
UPDATE score_items SET site_question = 'Há alguma outra doença que se repete na sua família?' WHERE id = '019bf31d-2ef0-7f5f-9813-561389776254';
UPDATE score_items SET site_question = 'Há alguma doença importante em parentes mais distantes (tios, primos, bisavós)?' WHERE id = '019bf31d-2ef0-7093-ad0d-3e4b7709e99f';
-- Hábitos e vícios dos parentes
UPDATE score_items SET site_question = 'O seu cônjuge ou parceiro(a) tem hábitos como fumar, beber em excesso ou usar drogas?' WHERE id = '019c5004-3ce9-7c9d-8558-71683ef4381e';
UPDATE score_items SET site_question = 'Que hábitos a sua mãe teve ao longo da vida (cigarro, álcool, drogas)?' WHERE id = '019c4ffd-f95a-76aa-bc4b-a397be677057';
UPDATE score_items SET site_question = 'Durante a gravidez de você, a sua mãe fumava, bebia ou usava alguma substância?' WHERE id = 'c77cedd3-2800-74b4-bb86-67d46e33c572';
UPDATE score_items SET site_question = 'Na época em que você foi concebido, o seu pai fumava, bebia ou usava alguma substância?' WHERE id = 'c77cedd3-2800-73fc-8a18-251bc6b8ebdf';
UPDATE score_items SET site_question = 'Os seus filhos têm hábitos prejudiciais à saúde (cigarro, álcool, drogas, sedentarismo)?' WHERE id = '019c5004-b33c-7eb2-a65b-087cef20af54';
UPDATE score_items SET site_question = 'Enquanto você era criança, a sua mãe fumava, bebia ou usava alguma substância?' WHERE id = '019c5000-8523-7979-88a1-9e05034b4dd8';
UPDATE score_items SET site_question = 'Que hábitos o seu pai teve ao longo da vida (cigarro, álcool, drogas)?' WHERE id = '019c4fff-c1c7-7ac5-af9e-a34aa2eb1fe8';
UPDATE score_items SET site_question = 'Enquanto você era criança, o seu pai fumava, bebia ou usava alguma substância?' WHERE id = '019c5001-6674-7350-a7b2-de15db2fbdcc';
UPDATE score_items SET site_question = 'Durante a sua adolescência, a sua mãe fumava, bebia ou usava alguma substância?' WHERE id = '019c5001-2633-7496-ac85-ae56c0863996';
UPDATE score_items SET site_question = 'Que hábitos prejudiciais existem entre os seus parentes próximos (cigarro, álcool, drogas)?' WHERE id = '019c5004-19d2-77a2-9251-2f521152a435';
UPDATE score_items SET site_question = 'Os seus outros parentes (tios, avós, primos) têm ou tiveram hábitos prejudiciais à saúde?' WHERE id = '019c5005-3230-7bbb-bb1f-ad8a4de289f6';
UPDATE score_items SET site_question = 'Durante a sua adolescência, o seu pai fumava, bebia ou usava alguma substância?' WHERE id = '019c5002-4e3a-7d96-b6d0-9b5ea0adce8c';

-- =========================================================================================
-- GRUPO 11 — SOCIAL
-- =========================================================================================
-- Histórico
UPDATE score_items SET site_question = 'Como são as condições da sua moradia (estrutura, umidade, conforto)?' WHERE id = '019c2e63-92df-7a7b-829e-7e0167b7d5cb';
UPDATE score_items SET site_question = 'Qual é a sua situação conjugal (casado, solteiro, separado, viúvo, em união)?' WHERE id = '019bf31d-2ef0-7df0-9a5e-e664e5622f1d';
UPDATE score_items SET site_question = 'Como é a sua situação familiar e a sua rede de apoio?' WHERE id = '019bf31d-2ef0-75a9-bc40-692bbb4c010c';
UPDATE score_items SET site_question = 'Você tem animais de estimação? Quais?' WHERE id = '019c2e63-92e0-7ecb-bd7b-fc0d46bd05d5';
UPDATE score_items SET site_question = 'A fé ou a espiritualidade fazem parte da sua vida? De que forma?' WHERE id = '019bf31d-2ef0-73e5-976d-ee56fc7a2ce3';
UPDATE score_items SET site_question = 'Como você descreveria a sua situação financeira e as suas fontes de renda?' WHERE id = 'c77cedd3-2800-7ef7-880b-1faef5982a55';
UPDATE score_items SET site_question = 'Qual é a sua profissão e como foi a sua trajetória de trabalho?' WHERE id = '019c2e63-92e0-73fa-ae5a-7e33f3f91703';
UPDATE score_items SET site_question = 'Você tem atividades de lazer ou hobbies na sua rotina? Quais?' WHERE id = '019bf31d-2ef0-78e2-9cdc-b892a703c6d5';
-- Atual
UPDATE score_items SET site_question = 'Como é a sua moradia atualmente, e quem mora com você?' WHERE id = '019bf31d-2ef0-7801-bb58-5cf2e4209288';
UPDATE score_items SET site_question = 'Como é a qualidade do ar dentro da sua casa (ventilação, mofo, poeira)?' WHERE id = '019bf31d-2ef0-7901-b3ce-42d6d870fe8c';
UPDATE score_items SET site_question = 'Você está exposto a alguma substância ou condição ambiental em casa (produtos químicos, mofo, fumaça)?' WHERE id = '019bf31d-2ef0-7817-9a38-3e34313af80b';
UPDATE score_items SET site_question = 'O ambiente da sua casa é silencioso ou barulhento?' WHERE id = '019bf31d-2ef0-7a78-8741-441f89630ff7';
UPDATE score_items SET site_question = 'Como é a qualidade da água que você bebe (torneira, filtro, mineral)?' WHERE id = '019bf31d-2ef0-7357-8bc3-70a48b490520';
UPDATE score_items SET site_question = 'A sua casa recebe bastante luz natural do sol?' WHERE id = 'c77cedd3-2800-769d-bd80-42730c51e164';
UPDATE score_items SET site_question = 'Você tem espaço em casa ou por perto para se exercitar?' WHERE id = '019bf31d-2ef0-7ae7-8a07-f92d9edccc4d';
UPDATE score_items SET site_question = 'Como está a sua situação conjugal atualmente?' WHERE id = '019c2e64-1521-7348-aa97-cae24044024d';
UPDATE score_items SET site_question = 'Você convive com animais de estimação atualmente?' WHERE id = '019bf31d-2ef0-7619-a877-a79839c49936';
UPDATE score_items SET site_question = 'Como está a sua situação financeira atualmente?' WHERE id = '019c2e64-1522-78e9-b4d0-832e615a6603';
UPDATE score_items SET site_question = 'Você reserva tempo para o lazer no seu dia a dia?' WHERE id = 'c77cedd3-2800-7580-b577-771aa448b12b';
UPDATE score_items SET site_question = 'Você pratica algum hobby regularmente?' WHERE id = '019bf31d-2ef0-7f37-822d-377fdaaf2fca';
UPDATE score_items SET site_question = 'Qual é a sua profissão atual e como é a sua rotina de trabalho?' WHERE id = 'c77cedd3-2800-73ab-a4ee-d2c0e08de952';
UPDATE score_items SET site_question = 'A fé ou a espiritualidade fazem parte da sua vida hoje? De que forma?' WHERE id = '019c2e64-1522-75cc-9c9c-dea7f149d50b';

COMMIT;
