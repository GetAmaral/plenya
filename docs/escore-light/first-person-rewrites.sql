-- Escore Light — reescrita das respostas para 1ª pessoa do singular ("eu/tenho/tive")
-- Auditoria: muitas opções estavam em terceira pessoa ("tem", "teve") quando a pergunta usa "você",
-- criando incoerência de sujeito. Esta SQL alinha tudo para "Eu tenho/tive/sinto/sou".

BEGIN;

-- =========================================================
-- 1. Lesões relacionadas ao exercício
-- =========================================================
UPDATE score_levels SET name = 'Tive uma lesão grave que deixou limitação permanente' WHERE item_id = 'c77cedd3-2800-7de9-ad64-b75ed1037b5e' AND level = 0;
UPDATE score_levels SET name = 'Tive uma lesão leve que deixou limitação permanente' WHERE item_id = 'c77cedd3-2800-7de9-ad64-b75ed1037b5e' AND level = 1;
UPDATE score_levels SET name = 'Tive lesão que limitou minha rotina por 1 a 5 anos' WHERE item_id = 'c77cedd3-2800-7de9-ad64-b75ed1037b5e' AND level = 2;
UPDATE score_levels SET name = 'Tive lesão grave nos últimos 12 meses' WHERE item_id = 'c77cedd3-2800-7de9-ad64-b75ed1037b5e' AND level = 3;
UPDATE score_levels SET name = 'Tive lesão leve nos últimos 12 meses' WHERE item_id = 'c77cedd3-2800-7de9-ad64-b75ed1037b5e' AND level = 4;
UPDATE score_levels SET name = 'Nunca tive lesão grave' WHERE item_id = 'c77cedd3-2800-7de9-ad64-b75ed1037b5e' AND level = 5;

-- =========================================================
-- 2. Doença cardiovascular (pessoal)
-- =========================================================
UPDATE score_levels SET name = 'Tive 2 ou mais eventos graves' WHERE item_id = '019bf31d-2ef0-7de7-a98d-603c41d12ae6' AND level = 0;
UPDATE score_levels SET name = 'Tive pelo menos um evento grave' WHERE item_id = '019bf31d-2ef0-7de7-a98d-603c41d12ae6' AND level = 1;
UPDATE score_levels SET name = 'Tive mais de um evento leve sem sequelas' WHERE item_id = '019bf31d-2ef0-7de7-a98d-603c41d12ae6' AND level = 2;
UPDATE score_levels SET name = 'Tive um único evento leve, sem sequelas' WHERE item_id = '019bf31d-2ef0-7de7-a98d-603c41d12ae6' AND level = 3;
UPDATE score_levels SET name = 'Nunca tive, mas tenho histórico familiar forte' WHERE item_id = '019bf31d-2ef0-7de7-a98d-603c41d12ae6' AND level = 4;
UPDATE score_levels SET name = 'Nunca tive' WHERE item_id = '019bf31d-2ef0-7de7-a98d-603c41d12ae6' AND level = 5;

-- =========================================================
-- 3. Insuficiência cardíaca — só N4 e N5
-- =========================================================
UPDATE score_levels SET name = 'Não tenho, mas tenho histórico familiar forte' WHERE item_id = '019bf31d-2ef0-7e03-bb24-dbdc747e5fd4' AND level = 4;
UPDATE score_levels SET name = 'Não tenho' WHERE item_id = '019bf31d-2ef0-7e03-bb24-dbdc747e5fd4' AND level = 5;

-- =========================================================
-- 4. Diabetes mellitus
-- =========================================================
UPDATE score_levels SET name = 'Tenho há mais de 5 anos, com controle inadequado' WHERE item_id = '019bf31d-2ef0-7f17-a053-7f45b7162dd2' AND level = 0;
UPDATE score_levels SET name = 'Tenho há menos de 5 anos, com controle inadequado' WHERE item_id = '019bf31d-2ef0-7f17-a053-7f45b7162dd2' AND level = 1;
UPDATE score_levels SET name = 'Tenho há mais de 5 anos, com controle adequado' WHERE item_id = '019bf31d-2ef0-7f17-a053-7f45b7162dd2' AND level = 2;
UPDATE score_levels SET name = 'Tenho há menos de 5 anos, com controle adequado' WHERE item_id = '019bf31d-2ef0-7f17-a053-7f45b7162dd2' AND level = 3;
UPDATE score_levels SET name = 'Não tenho, mas tenho histórico familiar forte' WHERE item_id = '019bf31d-2ef0-7f17-a053-7f45b7162dd2' AND level = 4;
UPDATE score_levels SET name = 'Não tenho' WHERE item_id = '019bf31d-2ef0-7f17-a053-7f45b7162dd2' AND level = 5;

-- =========================================================
-- 5. Pré-diabetes / Resistência a insulina
-- =========================================================
UPDATE score_levels SET name = 'Tenho diabetes já confirmado' WHERE item_id = '019bf31d-2ef0-7c5c-bc5b-00fbf09daf98' AND level = 0;
UPDATE score_levels SET name = 'Tenho pré-diabetes' WHERE item_id = '019bf31d-2ef0-7c5c-bc5b-00fbf09daf98' AND level = 1;
UPDATE score_levels SET name = 'Tenho resistência insulínica alta' WHERE item_id = '019bf31d-2ef0-7c5c-bc5b-00fbf09daf98' AND level = 2;
UPDATE score_levels SET name = 'Tenho resistência insulínica moderada' WHERE item_id = '019bf31d-2ef0-7c5c-bc5b-00fbf09daf98' AND level = 3;
UPDATE score_levels SET name = 'Tenho resistência insulínica leve' WHERE item_id = '019bf31d-2ef0-7c5c-bc5b-00fbf09daf98' AND level = 4;
UPDATE score_levels SET name = 'Não tenho' WHERE item_id = '019bf31d-2ef0-7c5c-bc5b-00fbf09daf98' AND level = 5;

-- =========================================================
-- 6. Arritmia — N5
-- =========================================================
UPDATE score_levels SET name = 'Nunca tive' WHERE item_id = 'c77cedd3-2800-7b94-ac6e-c404259efa2e' AND level = 5;

-- =========================================================
-- 7. Asma
-- =========================================================
UPDATE score_levels SET name = 'Tenho asma moderada/grave sem tratamento' WHERE item_id = 'c77cedd3-2800-7ef5-b90a-7b4b6cc19d5b' AND level = 0;
UPDATE score_levels SET name = 'Tenho asma moderada/grave de difícil controle' WHERE item_id = 'c77cedd3-2800-7ef5-b90a-7b4b6cc19d5b' AND level = 1;
UPDATE score_levels SET name = 'Tenho asma moderada/grave bem controlada' WHERE item_id = 'c77cedd3-2800-7ef5-b90a-7b4b6cc19d5b' AND level = 2;
UPDATE score_levels SET name = 'Tenho asma leve, sem precisar de remédios fortes' WHERE item_id = 'c77cedd3-2800-7ef5-b90a-7b4b6cc19d5b' AND level = 3;
UPDATE score_levels SET name = 'Não tenho, mas tenho histórico familiar forte' WHERE item_id = 'c77cedd3-2800-7ef5-b90a-7b4b6cc19d5b' AND level = 4;
UPDATE score_levels SET name = 'Não tenho' WHERE item_id = 'c77cedd3-2800-7ef5-b90a-7b4b6cc19d5b' AND level = 5;

-- =========================================================
-- 8-13. Histórico Familiar (6 items idênticos)
-- =========================================================
UPDATE score_levels SET name = 'Sim, 3 ou mais parentes próximos'
  WHERE level = 0 AND item_id IN (
    'c77cedd3-2800-70c7-942e-a1134e3aa05e', -- Doença CV familiar
    'c77cedd3-2800-7e0c-b44f-ccf9c3b926ef', -- Câncer
    'c77cedd3-2800-7524-9049-f4559170db14', -- DM/RI familiar
    'c77cedd3-2800-74b8-976d-8e90ecd896be', -- Dislipidemia
    'c77cedd3-2800-75fe-b483-ea0183478225', -- HAS familiar
    'c77cedd3-2800-70b8-b444-b05d15b96a57'  -- Obesidade familiar
  );
UPDATE score_levels SET name = 'Sim, 2 parentes próximos'
  WHERE level = 1 AND item_id IN (
    'c77cedd3-2800-70c7-942e-a1134e3aa05e', 'c77cedd3-2800-7e0c-b44f-ccf9c3b926ef',
    'c77cedd3-2800-7524-9049-f4559170db14', 'c77cedd3-2800-74b8-976d-8e90ecd896be',
    'c77cedd3-2800-75fe-b483-ea0183478225', 'c77cedd3-2800-70b8-b444-b05d15b96a57'
  );
UPDATE score_levels SET name = 'Sim, 1 parente próximo'
  WHERE level = 2 AND item_id IN (
    'c77cedd3-2800-70c7-942e-a1134e3aa05e', 'c77cedd3-2800-7e0c-b44f-ccf9c3b926ef',
    'c77cedd3-2800-7524-9049-f4559170db14', 'c77cedd3-2800-74b8-976d-8e90ecd896be',
    'c77cedd3-2800-75fe-b483-ea0183478225', 'c77cedd3-2800-70b8-b444-b05d15b96a57'
  );
UPDATE score_levels SET name = 'Não tenho conhecimento de casos na família'
  WHERE level = 5 AND item_id IN (
    'c77cedd3-2800-70c7-942e-a1134e3aa05e', 'c77cedd3-2800-7e0c-b44f-ccf9c3b926ef',
    'c77cedd3-2800-7524-9049-f4559170db14', 'c77cedd3-2800-74b8-976d-8e90ecd896be',
    'c77cedd3-2800-75fe-b483-ea0183478225', 'c77cedd3-2800-70b8-b444-b05d15b96a57'
  );

-- =========================================================
-- 14-17. Sintomas torácicos / pernas (Dor torácica, Dispnéia, Palpitação, Claudicação)
-- Padrão: "Quadro ativo / Teve quadro grave/leve / Já teve / Não tem"
-- =========================================================
UPDATE score_levels SET name = 'Tenho atualmente'
  WHERE level = 0 AND item_id IN (
    '019bf31d-2ef0-73fc-b353-73eadb4940d5', -- Dor torácica
    '019bf31d-2ef0-71ae-8f72-8330e3072647', -- Dispnéia
    '019bf31d-2ef0-7f0a-b7d9-5d22d7a40484'  -- Palpitação
  );
UPDATE score_levels SET name = 'Tive episódio grave no último ano'
  WHERE level = 1 AND item_id IN (
    '019bf31d-2ef0-73fc-b353-73eadb4940d5',
    '019bf31d-2ef0-71ae-8f72-8330e3072647',
    '019bf31d-2ef0-7f0a-b7d9-5d22d7a40484'
  );
UPDATE score_levels SET name = 'Tive episódio leve no último ano'
  WHERE level = 2 AND item_id IN (
    '019bf31d-2ef0-73fc-b353-73eadb4940d5',
    '019bf31d-2ef0-71ae-8f72-8330e3072647',
    '019bf31d-2ef0-7f0a-b7d9-5d22d7a40484'
  );
UPDATE score_levels SET name = 'Já tive um episódio grave (no passado)'
  WHERE level = 3 AND item_id IN (
    '019bf31d-2ef0-73fc-b353-73eadb4940d5',
    '019bf31d-2ef0-71ae-8f72-8330e3072647',
    '019bf31d-2ef0-7f0a-b7d9-5d22d7a40484'
  );
UPDATE score_levels SET name = 'Já tive no passado, sem deixar sequela'
  WHERE level = 4 AND item_id IN (
    '019bf31d-2ef0-73fc-b353-73eadb4940d5',
    '019bf31d-2ef0-71ae-8f72-8330e3072647',
    '019bf31d-2ef0-7f0a-b7d9-5d22d7a40484'
  );
UPDATE score_levels SET name = 'Nunca senti'
  WHERE level = 5 AND item_id IN (
    '019bf31d-2ef0-73fc-b353-73eadb4940d5',
    '019bf31d-2ef0-71ae-8f72-8330e3072647',
    '019bf31d-2ef0-7f0a-b7d9-5d22d7a40484'
  );

-- Claudicação (padrão "Quadro grave/leve ativo / Teve quadro grave/leve no último ano /...")
UPDATE score_levels SET name = 'Tenho atualmente, em forma grave' WHERE item_id = '019bf31d-2ef0-7193-b496-afe7ce51ed91' AND level = 0;
UPDATE score_levels SET name = 'Tenho atualmente, em forma leve' WHERE item_id = '019bf31d-2ef0-7193-b496-afe7ce51ed91' AND level = 1;
UPDATE score_levels SET name = 'Tive episódio grave no último ano' WHERE item_id = '019bf31d-2ef0-7193-b496-afe7ce51ed91' AND level = 2;
UPDATE score_levels SET name = 'Tive episódio leve no último ano' WHERE item_id = '019bf31d-2ef0-7193-b496-afe7ce51ed91' AND level = 3;
UPDATE score_levels SET name = 'Já tive no passado, há anos não tenho' WHERE item_id = '019bf31d-2ef0-7193-b496-afe7ce51ed91' AND level = 4;
UPDATE score_levels SET name = 'Nunca senti' WHERE item_id = '019bf31d-2ef0-7193-b496-afe7ce51ed91' AND level = 5;

-- =========================================================
-- 18-19. Sintomas GI (Azia, Obstipação) — mesmo padrão
-- =========================================================
UPDATE score_levels SET name = 'Tenho atualmente, em forma grave' WHERE item_id IN ('c77cedd3-2800-73d3-8d0f-6bdeb10d1680','019bf31d-2ef0-751d-b2c8-1ef3913bf002') AND level = 0;
UPDATE score_levels SET name = 'Tenho atualmente, em forma leve' WHERE item_id IN ('c77cedd3-2800-73d3-8d0f-6bdeb10d1680','019bf31d-2ef0-751d-b2c8-1ef3913bf002') AND level = 1;
UPDATE score_levels SET name = 'Tive episódio grave no último ano' WHERE item_id IN ('c77cedd3-2800-73d3-8d0f-6bdeb10d1680','019bf31d-2ef0-751d-b2c8-1ef3913bf002') AND level = 2;
UPDATE score_levels SET name = 'Tive episódio leve no último ano' WHERE item_id IN ('c77cedd3-2800-73d3-8d0f-6bdeb10d1680','019bf31d-2ef0-751d-b2c8-1ef3913bf002') AND level = 3;
UPDATE score_levels SET name = 'Já tive no passado, há anos não tenho' WHERE item_id IN ('c77cedd3-2800-73d3-8d0f-6bdeb10d1680','019bf31d-2ef0-751d-b2c8-1ef3913bf002') AND level = 4;
UPDATE score_levels SET name = 'Nunca senti' WHERE item_id IN ('c77cedd3-2800-73d3-8d0f-6bdeb10d1680','019bf31d-2ef0-751d-b2c8-1ef3913bf002') AND level = 5;

-- =========================================================
-- 20. Edema — N5
-- =========================================================
UPDATE score_levels SET name = 'Não tenho' WHERE item_id = '019bf31d-2ef0-7297-bd6d-4c83d8081d2d' AND level = 5;

-- =========================================================
-- 21. Socialização atual
-- =========================================================
UPDATE score_levels SET name = 'Estou isolado(a)' WHERE item_id = '019c54fc-edfb-73d7-8525-3e326e91a976' AND level = 0;
UPDATE score_levels SET name = 'Convivo só com a família' WHERE item_id = '019c54fc-edfb-73d7-8525-3e326e91a976' AND level = 2;
UPDATE score_levels SET name = 'Convivo só com amigos' WHERE item_id = '019c54fc-edfb-73d7-8525-3e326e91a976' AND level = 3;
UPDATE score_levels SET name = 'Convivo com família e amigos (rede social ativa)' WHERE item_id = '019c54fc-edfb-73d7-8525-3e326e91a976' AND level = 5;

-- =========================================================
-- 22. Fontes de stress percebidas
-- =========================================================
UPDATE score_levels SET name = 'Estou em burnout completo' WHERE item_id = 'c77cedd3-2800-7360-bf3c-5b4f28c660ef' AND level = 0;
UPDATE score_levels SET name = 'Estou em estresse intenso, próximo do colapso' WHERE item_id = 'c77cedd3-2800-7360-bf3c-5b4f28c660ef' AND level = 1;
UPDATE score_levels SET name = 'Tenho estresse difícil de manejar, com sintomas físicos' WHERE item_id = 'c77cedd3-2800-7360-bf3c-5b4f28c660ef' AND level = 2;
UPDATE score_levels SET name = 'Tenho muitas fontes de estresse, mas bem manejadas' WHERE item_id = 'c77cedd3-2800-7360-bf3c-5b4f28c660ef' AND level = 3;
UPDATE score_levels SET name = 'Tenho pouco estresse, bem manejado' WHERE item_id = 'c77cedd3-2800-7360-bf3c-5b4f28c660ef' AND level = 4;
UPDATE score_levels SET name = 'Não percebo estresse danoso na minha vida' WHERE item_id = 'c77cedd3-2800-7360-bf3c-5b4f28c660ef' AND level = 5;

-- =========================================================
-- 23. Tabaco
-- =========================================================
UPDATE score_levels SET name = 'Sou fumante atual, com longa história (≈ 1 maço/dia por 20+ anos)' WHERE item_id = '019bf31d-2ef0-7a08-ab7e-d9c06d8d2103' AND level = 0;
UPDATE score_levels SET name = 'Sou fumante atual, com história mais curta ou menos cigarros' WHERE item_id = '019bf31d-2ef0-7a08-ab7e-d9c06d8d2103' AND level = 1;
UPDATE score_levels SET name = 'Sou ex-fumante (parei há até 10 anos)' WHERE item_id = '019bf31d-2ef0-7a08-ab7e-d9c06d8d2103' AND level = 2;
UPDATE score_levels SET name = 'Sou ex-fumante (parei há mais de 10 anos), com longa história de fumo' WHERE item_id = '019bf31d-2ef0-7a08-ab7e-d9c06d8d2103' AND level = 3;
UPDATE score_levels SET name = 'Sou ex-fumante (parei há mais de 10 anos), com história curta de fumo' WHERE item_id = '019bf31d-2ef0-7a08-ab7e-d9c06d8d2103' AND level = 4;
UPDATE score_levels SET name = 'Nunca fumei' WHERE item_id = '019bf31d-2ef0-7a08-ab7e-d9c06d8d2103' AND level = 5;

-- =========================================================
-- 24. Álcool
-- =========================================================
UPDATE score_levels SET name = 'Tenho dependência ativa, com complicações' WHERE item_id = '019bf31d-2ef0-7eff-b67d-00e5eec8c2d0' AND level = 0;
UPDATE score_levels SET name = 'Bebo pesado todo dia (mais de 4 doses/dia)' WHERE item_id = '019bf31d-2ef0-7eff-b67d-00e5eec8c2d0' AND level = 1;
UPDATE score_levels SET name = 'Bebo de 2 a 4 doses por dia' WHERE item_id = '019bf31d-2ef0-7eff-b67d-00e5eec8c2d0' AND level = 2;
UPDATE score_levels SET name = 'Bebo até 1 dose por dia' WHERE item_id = '019bf31d-2ef0-7eff-b67d-00e5eec8c2d0' AND level = 3;
UPDATE score_levels SET name = 'Sou ex-etilista, ou bebo só ocasionalmente em eventos sociais' WHERE item_id = '019bf31d-2ef0-7eff-b67d-00e5eec8c2d0' AND level = 4;
UPDATE score_levels SET name = 'Não bebo (ou bebo raramente)' WHERE item_id = '019bf31d-2ef0-7eff-b67d-00e5eec8c2d0' AND level = 5;

-- =========================================================
-- 25. Situação familiar
-- =========================================================
UPDATE score_levels SET name = 'Não tenho família ou estou em rompimento total' WHERE item_id = '019c2e64-1522-74db-a6b5-97595412d14a' AND level = 0;
UPDATE score_levels SET name = 'Tenho conflitos graves ativos, contato raro ou sobrecarga de cuidado' WHERE item_id = '019c2e64-1522-74db-a6b5-97595412d14a' AND level = 1;
UPDATE score_levels SET name = 'Tenho conflitos recorrentes ou suporte familiar fraco' WHERE item_id = '019c2e64-1522-74db-a6b5-97595412d14a' AND level = 2;
UPDATE score_levels SET name = 'Tenho contato regular, com tensões ocasionais' WHERE item_id = '019c2e64-1522-74db-a6b5-97595412d14a' AND level = 3;
UPDATE score_levels SET name = 'Tenho família próxima, com suporte consistente' WHERE item_id = '019c2e64-1522-74db-a6b5-97595412d14a' AND level = 4;
UPDATE score_levels SET name = 'Tenho rede familiar sólida e relações harmoniosas' WHERE item_id = '019c2e64-1522-74db-a6b5-97595412d14a' AND level = 5;

-- =========================================================
-- 26. Qualidade percebida do sono
-- =========================================================
UPDATE score_levels SET name = 'Passo o dia todo cansado, sonolento e sem energia' WHERE item_id = 'c77cedd3-2800-735f-bf28-c5d07d7d7092' AND level = 0;
UPDATE score_levels SET name = 'Acordo bem, mas a energia acaba no meio do dia' WHERE item_id = 'c77cedd3-2800-735f-bf28-c5d07d7d7092' AND level = 2;
UPDATE score_levels SET name = 'Acordo cansado e depois "pego no tranco"' WHERE item_id = 'c77cedd3-2800-735f-bf28-c5d07d7d7092' AND level = 3;
UPDATE score_levels SET name = 'É reparador, com energia o dia todo' WHERE item_id = 'c77cedd3-2800-735f-bf28-c5d07d7d7092' AND level = 5;

-- =========================================================
-- 27. Roncos
-- =========================================================
UPDATE score_levels SET name = 'Ronco com frequência e intensidade' WHERE item_id = '019c164e-cb2c-747a-bfc7-ffb3e787229f' AND level = 0;
UPDATE score_levels SET name = 'Ronco esporadicamente' WHERE item_id = '019c164e-cb2c-747a-bfc7-ffb3e787229f' AND level = 1;
UPDATE score_levels SET name = 'Não ronco' WHERE item_id = '019c164e-cb2c-747a-bfc7-ffb3e787229f' AND level = 5;

-- =========================================================
-- 28. Tela noturna — "Usa..." → "Uso..."
-- =========================================================
UPDATE score_levels SET name = 'Uso todas as noites' WHERE item_id = '019bf31d-2ef0-7f44-a617-84562c569b71' AND level = 0;
UPDATE score_levels SET name = 'Uso 4-6x por semana' WHERE item_id = '019bf31d-2ef0-7f44-a617-84562c569b71' AND level = 2;
UPDATE score_levels SET name = 'Uso 3-4x por semana' WHERE item_id = '019bf31d-2ef0-7f44-a617-84562c569b71' AND level = 3;
UPDATE score_levels SET name = 'Uso 1-2x por semana' WHERE item_id = '019bf31d-2ef0-7f44-a617-84562c569b71' AND level = 4;
UPDATE score_levels SET name = 'Não uso' WHERE item_id = '019bf31d-2ef0-7f44-a617-84562c569b71' AND level = 5;

-- =========================================================
-- 29. Refrigerantes — "Consome..." → "Consumo..."
-- =========================================================
UPDATE score_levels SET name = 'Consumo refrigerante/energético com açúcar' WHERE item_id = '019c535e-d7b0-7f20-8c8c-b3c2dee441fc' AND level = 0;
UPDATE score_levels SET name = 'Consumo apenas versões sem açúcar (zero/diet)' WHERE item_id = '019c535e-d7b0-7f20-8c8c-b3c2dee441fc' AND level = 2;
UPDATE score_levels SET name = 'Não consumo' WHERE item_id = '019c535e-d7b0-7f20-8c8c-b3c2dee441fc' AND level = 5;

-- =========================================================
-- 30. Apneias — manter o padrão da reescrita anterior; só ajustar N5
-- =========================================================
UPDATE score_levels SET name = 'Não apresento (que eu saiba)' WHERE item_id = 'c77cedd3-2800-7ca1-9dd2-ac6e478cd0e5' AND level = 5;

COMMIT;
