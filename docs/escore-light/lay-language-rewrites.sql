-- Escore Light — reescrita de jargão médico para linguagem leiga
-- Aplicado em 2026-04-22 após auditoria das perguntas e respostas do Light.

BEGIN;

-- =========================================================
-- 1. Uso de BiPAP/CPAP — explicar o que são os aparelhos
-- =========================================================
UPDATE score_items
SET light_question = 'Você usa CPAP ou BiPAP para dormir? (aparelhos que mantêm a via aérea aberta durante o sono — tratamento padrão da apneia obstrutiva do sono)'
WHERE id = '019c28f0-05c1-7c0f-92b5-9cc9adac6eb3';

-- =========================================================
-- 2. Edema — substituir "anasarca" e "edema" por "inchaço"
-- =========================================================
UPDATE score_items
SET light_question = 'Você apresenta inchaço (edema) nos braços, pernas ou no corpo?'
WHERE id = '019bf31d-2ef0-7297-bd6d-4c83d8081d2d';

UPDATE score_levels SET name = 'Inchaço severo no corpo todo' WHERE item_id = '019bf31d-2ef0-7297-bd6d-4c83d8081d2d' AND level = 0;
UPDATE score_levels SET name = 'Inchaço intenso ao longo do dia' WHERE item_id = '019bf31d-2ef0-7297-bd6d-4c83d8081d2d' AND level = 1;
UPDATE score_levels SET name = 'Inchaço leve persistente o dia todo' WHERE item_id = '019bf31d-2ef0-7297-bd6d-4c83d8081d2d' AND level = 2;
UPDATE score_levels SET name = 'Inchaço nas pernas no fim do dia' WHERE item_id = '019bf31d-2ef0-7297-bd6d-4c83d8081d2d' AND level = 3;
UPDATE score_levels SET name = 'Inchaço esporádico em mãos e pés' WHERE item_id = '019bf31d-2ef0-7297-bd6d-4c83d8081d2d' AND level = 4;

-- =========================================================
-- 3. Insuficiência cardíaca — explicar "IC"
-- =========================================================
UPDATE score_items
SET light_question = 'Você tem diagnóstico de insuficiência cardíaca? (condição em que o coração tem dificuldade de bombear sangue)'
WHERE id = '019bf31d-2ef0-7e03-bb24-dbdc747e5fd4';

UPDATE score_levels SET name = 'Tenho, em descompensação atual (sintomas piorando)' WHERE item_id = '019bf31d-2ef0-7e03-bb24-dbdc747e5fd4' AND level = 0;
UPDATE score_levels SET name = 'Tenho em estado grave' WHERE item_id = '019bf31d-2ef0-7e03-bb24-dbdc747e5fd4' AND level = 1;
UPDATE score_levels SET name = 'Tenho em estado moderado' WHERE item_id = '019bf31d-2ef0-7e03-bb24-dbdc747e5fd4' AND level = 2;
UPDATE score_levels SET name = 'Tenho em estado leve' WHERE item_id = '019bf31d-2ef0-7e03-bb24-dbdc747e5fd4' AND level = 3;

-- =========================================================
-- 4. Arritmia — explicar CDI/MP
-- =========================================================
UPDATE score_levels SET name = 'Tenho desfibrilador implantado (CDI) ou ressincronizador' WHERE item_id = 'c77cedd3-2800-7b94-ac6e-c404259efa2e' AND level = 0;
UPDATE score_levels SET name = 'Tenho marcapasso implantado' WHERE item_id = 'c77cedd3-2800-7b94-ac6e-c404259efa2e' AND level = 1;
UPDATE score_levels SET name = 'Tenho arritmia moderada/grave em acompanhamento' WHERE item_id = 'c77cedd3-2800-7b94-ac6e-c404259efa2e' AND level = 2;
UPDATE score_levels SET name = 'Já tive arritmia moderada/grave ou precisei de procedimento no coração' WHERE item_id = 'c77cedd3-2800-7b94-ac6e-c404259efa2e' AND level = 3;
UPDATE score_levels SET name = 'Já tive ou tenho arritmias leves (sem gravidade)' WHERE item_id = 'c77cedd3-2800-7b94-ac6e-c404259efa2e' AND level = 4;

-- =========================================================
-- 5. Asma — substituir "imunossupressor"
-- =========================================================
UPDATE score_levels SET name = 'Asma moderada/grave sem tratamento' WHERE item_id = 'c77cedd3-2800-7ef5-b90a-7b4b6cc19d5b' AND level = 0;
UPDATE score_levels SET name = 'Asma moderada/grave de difícil controle' WHERE item_id = 'c77cedd3-2800-7ef5-b90a-7b4b6cc19d5b' AND level = 1;
UPDATE score_levels SET name = 'Asma moderada/grave bem controlada' WHERE item_id = 'c77cedd3-2800-7ef5-b90a-7b4b6cc19d5b' AND level = 2;
UPDATE score_levels SET name = 'Asma leve, sem precisar de remédios fortes' WHERE item_id = 'c77cedd3-2800-7ef5-b90a-7b4b6cc19d5b' AND level = 3;

-- =========================================================
-- 6. Histórico Periodontal — substituir periodontite/gengivite por linguagem clara
-- =========================================================
UPDATE score_items
SET light_question = 'Como é seu histórico de problemas nas gengivas (sangramento, inflamação, infecções)?'
WHERE id = '019bf31d-2ef0-7f92-900b-138dfe299a5f';

UPDATE score_levels SET name = 'Doença grave nas gengivas já tratada, ou várias infecções (abscessos)' WHERE item_id = '019bf31d-2ef0-7f92-900b-138dfe299a5f' AND level = 0;
UPDATE score_levels SET name = 'Doença moderada nas gengivas ou 1-2 infecções' WHERE item_id = '019bf31d-2ef0-7f92-900b-138dfe299a5f' AND level = 1;
UPDATE score_levels SET name = 'Inflamação leve nas gengivas ou sangramento frequente' WHERE item_id = '019bf31d-2ef0-7f92-900b-138dfe299a5f' AND level = 2;
UPDATE score_levels SET name = 'Já tive gengivas inflamadas, mas tratei e não voltou' WHERE item_id = '019bf31d-2ef0-7f92-900b-138dfe299a5f' AND level = 3;
UPDATE score_levels SET name = 'Episódio isolado de inflamação na gengiva' WHERE item_id = '019bf31d-2ef0-7f92-900b-138dfe299a5f' AND level = 4;
UPDATE score_levels SET name = 'Nunca tive problemas nas gengivas' WHERE item_id = '019bf31d-2ef0-7f92-900b-138dfe299a5f' AND level = 5;

-- =========================================================
-- 7. Tabaco — explicar "maços-ano"
-- =========================================================
UPDATE score_levels SET name = 'Fumante atual, longa história (equivale a 1 maço/dia por 20+ anos)' WHERE item_id = '019bf31d-2ef0-7a08-ab7e-d9c06d8d2103' AND level = 0;
UPDATE score_levels SET name = 'Fumante atual, história mais curta ou menos cigarros' WHERE item_id = '019bf31d-2ef0-7a08-ab7e-d9c06d8d2103' AND level = 1;
UPDATE score_levels SET name = 'Ex-fumante (parou há até 10 anos)' WHERE item_id = '019bf31d-2ef0-7a08-ab7e-d9c06d8d2103' AND level = 2;
UPDATE score_levels SET name = 'Ex-fumante (parou há mais de 10 anos), com longa história de fumo' WHERE item_id = '019bf31d-2ef0-7a08-ab7e-d9c06d8d2103' AND level = 3;
UPDATE score_levels SET name = 'Ex-fumante (parou há mais de 10 anos), com história curta de fumo' WHERE item_id = '019bf31d-2ef0-7a08-ab7e-d9c06d8d2103' AND level = 4;

-- =========================================================
-- 8. Apneias — quantificar de forma compreensível
-- =========================================================
UPDATE score_levels SET name = 'Frequentes e intensas (várias por noite)' WHERE item_id = 'c77cedd3-2800-7ca1-9dd2-ac6e478cd0e5' AND level = 0;
UPDATE score_levels SET name = 'Algumas por semana (5 a 10 por noite)' WHERE item_id = 'c77cedd3-2800-7ca1-9dd2-ac6e478cd0e5' AND level = 1;
UPDATE score_levels SET name = 'Esporádicas (1 episódio por noite)' WHERE item_id = 'c77cedd3-2800-7ca1-9dd2-ac6e478cd0e5' AND level = 3;
UPDATE score_levels SET name = 'Não apresento' WHERE item_id = 'c77cedd3-2800-7ca1-9dd2-ac6e478cd0e5' AND level = 5;

COMMIT;
