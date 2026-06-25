-- Scorear o subgrupo "Histórico" do grupo Social, espelhando os itens do "Atual".
-- Decisão (Getúlio): peso = 1/3 do Atual (arredondado p/ inteiro; soma 15 = 45/3):
--   conjugal 3 · familiar 4 · pets 1 · financeira 5 · lazer 2.
-- Os 5 itens já estão vinculados aos mesmos MethodPillars dos espelhos do Atual; o radar é
-- calculado ao vivo dos points, então não precisa regenerar score_version nem snapshot.
-- Itens narrativos sem espelho (Condições de moradia, Religiosidade, Profissões) ficam como text.
-- Níveis reescritos no passado; gradiente 0=adverso … 5=saudável (pets pula o 2, como no Atual).
-- Idempotente. Aplicado em DEV; rodar em PROD quando autorizado.

BEGIN;

-- ===== 1) itens: text → level_choice, points, pergunta no passado =====
UPDATE score_items SET site_render_type='level_choice', points=3,
  site_question='Como foi o seu histórico conjugal ao longo da vida?', updated_at=now()
WHERE id='019bf31d-2ef0-7df0-9a5e-e664e5622f1d'; -- Situação conjugal

UPDATE score_items SET site_render_type='level_choice', points=4,
  site_question='Como foi a sua situação familiar e rede de apoio ao longo da vida?', updated_at=now()
WHERE id='019bf31d-2ef0-75a9-bc40-692bbb4c010c'; -- Situação familiar

UPDATE score_items SET site_render_type='level_choice', points=1,
  site_question='Como foi a sua convivência com animais de estimação ao longo da vida?', updated_at=now()
WHERE id='019c2e63-92e0-7ecb-bd7b-fc0d46bd05d5'; -- Situação de pets

UPDATE score_items SET site_render_type='level_choice', points=5,
  site_question='Como foi a sua situação financeira ao longo da vida?', updated_at=now()
WHERE id='c77cedd3-2800-7ef7-880b-1faef5982a55'; -- Situação financeira

UPDATE score_items SET site_render_type='level_choice', points=2,
  site_question='Como foi o seu histórico de lazer e hobbies ao longo da vida?', updated_at=now()
WHERE id='019bf31d-2ef0-78e2-9cdc-b892a703c6d5'; -- Laser / hobbies

-- ===== 2) níveis (operator '=', guardado por item+level p/ idempotência) =====
INSERT INTO score_levels (id, item_id, level, operator, name, site_legend, created_at, updated_at)
SELECT uuid_generate_v7(), v.item_id, v.level, '=', v.name, v.legend, now(), now()
FROM (VALUES
  -- Situação conjugal (id 019bf31d-2ef0-7df0-9a5e-e664e5622f1d)
  ('019bf31d-2ef0-7df0-9a5e-e664e5622f1d'::uuid, 0, 'Passei por separação conflituosa ou luto conjugal marcante', 'Separação difícil ou perda marcante de parceiro no passado, que pesou emocionalmente.'),
  ('019bf31d-2ef0-7df0-9a5e-e664e5622f1d'::uuid, 1, 'Vivi relação com conflitos graves frequentes ou abusiva', 'Relacionamento com conflitos frequentes ou abusivo no passado, merece atenção.'),
  ('019bf31d-2ef0-7df0-9a5e-e664e5622f1d'::uuid, 2, 'Tive relação instável ou longo período de isolamento social', 'Relação instável ou longa solidão social no passado, com pouco apoio próximo.'),
  ('019bf31d-2ef0-7df0-9a5e-e664e5622f1d'::uuid, 3, 'Tive relação estável com dificuldades, ou rede social limitada', 'Relação estável com dificuldades, ou rede de apoio limitada no passado.'),
  ('019bf31d-2ef0-7df0-9a5e-e664e5622f1d'::uuid, 4, 'Tive relação estável, construindo parceria', 'Relação estável e em construção ao longo da vida.'),
  ('019bf31d-2ef0-7df0-9a5e-e664e5622f1d'::uuid, 5, 'Tive relação sólida ou boa rede de apoio ao longo da vida', 'Histórico de relação afetiva sólida ou boa rede de apoio.'),
  -- Situação familiar (id 019bf31d-2ef0-75a9-bc40-692bbb4c010c)
  ('019bf31d-2ef0-75a9-bc40-692bbb4c010c'::uuid, 0, 'Cresci sem família ou em rompimento total', 'Sem família ou em rompimento total no passado, situação que pesou.'),
  ('019bf31d-2ef0-75a9-bc40-692bbb4c010c'::uuid, 1, 'Vivi conflitos graves, contato raro ou sobrecarga de cuidado', 'Conflitos graves na família ou sobrecarga de cuidado no passado, merece atenção.'),
  ('019bf31d-2ef0-75a9-bc40-692bbb4c010c'::uuid, 2, 'Tive conflitos recorrentes ou suporte familiar fraco', 'Conflitos frequentes ou pouco apoio da família no passado.'),
  ('019bf31d-2ef0-75a9-bc40-692bbb4c010c'::uuid, 3, 'Tive contato regular, com tensões ocasionais', 'Contato familiar regular, com tensões pontuais no passado.'),
  ('019bf31d-2ef0-75a9-bc40-692bbb4c010c'::uuid, 4, 'Tive família próxima, com suporte consistente', 'Família próxima com apoio consistente ao longo da vida.'),
  ('019bf31d-2ef0-75a9-bc40-692bbb4c010c'::uuid, 5, 'Tive rede familiar sólida e relações harmoniosas', 'Histórico de família presente e relações harmoniosas.'),
  -- Situação de pets (id 019c2e63-92e0-7ecb-bd7b-fc0d46bd05d5) — pula nível 2
  ('019c2e63-92e0-7ecb-bd7b-fc0d46bd05d5'::uuid, 0, 'Passei por luto de pet ou animal com doença grave que pesou', 'Perda de um pet ou bicho doente no passado que gerou sobrecarga emocional.'),
  ('019c2e63-92e0-7ecb-bd7b-fc0d46bd05d5'::uuid, 1, 'Tive pet com problemas comportamentais ou cuidado oneroso', 'Pet com problemas de comportamento ou cuidado pesado no passado.'),
  ('019c2e63-92e0-7ecb-bd7b-fc0d46bd05d5'::uuid, 3, 'Tive pet com pouca interação', 'Teve pet, mas com pouca convivência.'),
  ('019c2e63-92e0-7ecb-bd7b-fc0d46bd05d5'::uuid, 4, 'Quis ter pets mas não pude', 'Gostaria de ter tido um pet, mas não pôde.'),
  ('019c2e63-92e0-7ecb-bd7b-fc0d46bd05d5'::uuid, 5, 'Tive pet saudável com bom vínculo, ou optei por não ter', 'Pet saudável com bom vínculo no passado, ou opção por não ter, situação positiva.'),
  -- Situação financeira (id c77cedd3-2800-7ef7-880b-1faef5982a55)
  ('c77cedd3-2800-7ef7-880b-1faef5982a55'::uuid, 0, 'Passei por endividamento grave ou renda insuficiente para o básico', 'Endividamento grave ou renda insuficiente para o básico no passado, fonte de estresse.'),
  ('c77cedd3-2800-7ef7-880b-1faef5982a55'::uuid, 1, 'Vivi com renda apertada, sem reserva, estresse financeiro frequente', 'Renda apertada sem reserva e estresse financeiro frequente no passado.'),
  ('c77cedd3-2800-7ef7-880b-1faef5982a55'::uuid, 2, 'Tive renda suficiente, mas reserva curta e sem renda passiva', 'Renda cobria os gastos, mas com reserva curta e sem renda passiva.'),
  ('c77cedd3-2800-7ef7-880b-1faef5982a55'::uuid, 3, 'Tive renda confortável, reserva de 6 a 12 meses', 'Renda confortável e reserva para vários meses no passado.'),
  ('c77cedd3-2800-7ef7-880b-1faef5982a55'::uuid, 4, 'Tive renda ativa e passiva, boa reserva, sem estresse', 'Renda ativa e passiva com boa reserva e sem estresse financeiro.'),
  ('c77cedd3-2800-7ef7-880b-1faef5982a55'::uuid, 5, 'Tive independência financeira ou patrimônio sólido', 'Histórico de independência financeira com patrimônio sólido.'),
  -- Laser / hobbies (id 019bf31d-2ef0-78e2-9cdc-b892a703c6d5) — espelha "Lazer"
  ('019bf31d-2ef0-78e2-9cdc-b892a703c6d5'::uuid, 0, 'Passei longos períodos sem tempo ou energia para lazer', 'Longos períodos sem tempo ou energia para lazer no passado.'),
  ('019bf31d-2ef0-78e2-9cdc-b892a703c6d5'::uuid, 1, 'Tive lazer raro, com culpa ao descansar', 'Lazer raro e culpa ao descansar no passado, merece cuidado.'),
  ('019bf31d-2ef0-78e2-9cdc-b892a703c6d5'::uuid, 2, 'Tive lazer ocasional, com dificuldade de desligar', 'Pouco lazer e dificuldade de desligar no passado.'),
  ('019bf31d-2ef0-78e2-9cdc-b892a703c6d5'::uuid, 3, 'Tive lazer regular, mas pouco variado', 'Lazer regular, mas pouco variado no passado.'),
  ('019bf31d-2ef0-78e2-9cdc-b892a703c6d5'::uuid, 4, 'Tive lazer semanal variado, conseguia desconectar', 'Lazer variado na semana e boa capacidade de desconectar.'),
  ('019bf31d-2ef0-78e2-9cdc-b892a703c6d5'::uuid, 5, 'Tive lazer frequente, variado e equilibrado', 'Histórico de lazer frequente, variado e equilibrado com a rotina.')
) AS v(item_id, level, name, legend)
WHERE NOT EXISTS (
  SELECT 1 FROM score_levels sl WHERE sl.item_id = v.item_id AND sl.level = v.level
);

COMMIT;
