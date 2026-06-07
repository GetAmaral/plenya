-- seed-prep-forms.sql — cria os 3 formulários de preparação pré-consulta como ScoreVersion
-- (context='patient_prep') + seus itens curados (subset dos 538 de anamnese, todos já com
-- site_question). Idempotente: upsert das versions por id + delete/insert dos itens por version.
-- Curadoria INICIAL (ajustável pelo Dr. no builder; é DADO, não código).
--   prep-a1  Consulta Plenya avulsa (geral)
--   prep-b1  Entrada Continuum (prospect/encantar — objetivos + escalas + medidas + rastreios)
--   prep-b2  Continuum aprofundamento (história mais funda + reconciliação + composição)
-- Dev: docker compose exec -T db psql -U plenya_user -d plenya_db -f /tmp/seed-prep-forms.sql

BEGIN;

INSERT INTO score_versions (id, name, slug, description, site_intro, "order", active, context, created_at, updated_at)
VALUES
 ('11111111-2222-7333-8444-000000000001', 'Preparação — Consulta Plenya', 'prep-a1',
  'Formulário de preparação para a consulta médica Plenya (avulsa).',
  'Separe alguns minutos antes da sua consulta. Quanto mais o seu médico souber de antemão, melhor a gente aproveita o tempo de vocês.',
  10, true, 'patient_prep', now(), now()),
 ('11111111-2222-7333-8444-000000000002', 'Preparação — Entrada Continuum', 'prep-b1',
  'Formulário de preparação para a primeira conversa do Continuum.',
  'Para a nossa primeira conversa render, conte um pouco dos seus objetivos e responda algumas perguntas rápidas. Leva poucos minutos.',
  11, true, 'patient_prep', now(), now()),
 ('11111111-2222-7333-8444-000000000003', 'Preparação — Continuum (aprofundamento)', 'prep-b2',
  'Formulário de preparação aprofundado para o início do acompanhamento Continuum.',
  'Agora vamos montar a sua base completa. Reserve um tempo tranquilo e, se puder, tenha seus exames recentes em mãos para enviar.',
  12, true, 'patient_prep', now(), now())
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name, slug = EXCLUDED.slug, description = EXCLUDED.description,
  site_intro = EXCLUDED.site_intro, "order" = EXCLUDED."order",
  active = EXCLUDED.active, context = EXCLUDED.context, updated_at = now();

DELETE FROM score_version_items WHERE version_id IN (
  '11111111-2222-7333-8444-000000000001',
  '11111111-2222-7333-8444-000000000002',
  '11111111-2222-7333-8444-000000000003'
);

-- prep-a1 (Consulta Plenya avulsa) — núcleo + histórico familiar + crônicas + hábitos
INSERT INTO score_version_items (id, version_id, score_item_id, display_order, created_at, updated_at)
SELECT gen_random_uuid(), '11111111-2222-7333-8444-000000000001', x.id, x.ord, now(), now()
FROM (VALUES
  ('019bf31d-2ef0-78da-9d77-4e8258d3cf8e'::uuid, 1),   -- Uso atual de medicamentos
  ('019bf31d-2ef0-7090-b843-99752e3c622d'::uuid, 2),   -- Suplementações utilizadas
  ('019bf31d-2ef0-74c5-8d99-0f355f1aa7cc'::uuid, 3),   -- Peso
  ('019bf31d-2ef0-71e4-a845-cdfaeedbb599'::uuid, 4),   -- Altura
  ('019c534a-afc3-70c4-82e0-bfde4b5b8f93'::uuid, 5),   -- Padrão alimentar atual
  ('c77cedd3-2800-735f-bf28-c5d07d7d7092'::uuid, 6),   -- Qualidade percebida do sono
  ('019c53a6-f1a3-7704-9868-354859c750cd'::uuid, 7),   -- Tempo de sono
  ('019c164e-cb2c-747a-bfc7-ffb3e787229f'::uuid, 8),   -- Roncos
  ('019bf31d-2ef0-7a08-ab7e-d9c06d8d2103'::uuid, 9),   -- Tabaco (hábito)
  ('019bf31d-2ef0-7eff-b67d-00e5eec8c2d0'::uuid, 10),  -- Álcool (hábito)
  ('019bf31d-2ef0-7f17-a053-7f45b7162dd2'::uuid, 11),  -- Diabetes mellitus (crônica)
  ('019bf31d-2ef0-7de7-a98d-603c41d12ae6'::uuid, 12),  -- Doença cardiovascular (crônica)
  ('c77cedd3-2800-75fe-b483-ea0183478225'::uuid, 13),  -- Familiar: hipertensão
  ('c77cedd3-2800-7524-9049-f4559170db14'::uuid, 14),  -- Familiar: diabetes
  ('c77cedd3-2800-7e0c-b44f-ccf9c3b926ef'::uuid, 15),  -- Familiar: câncer
  ('c77cedd3-2800-70c7-942e-a1134e3aa05e'::uuid, 16)   -- Familiar: doença cardiovascular
) AS x(id, ord)
JOIN score_items si ON si.id = x.id AND si.deleted_at IS NULL;

-- prep-b1 (Entrada Continuum) — objetivos/futuro + escalas + medidas + rastreios cross-AGIR
INSERT INTO score_version_items (id, version_id, score_item_id, display_order, created_at, updated_at)
SELECT gen_random_uuid(), '11111111-2222-7333-8444-000000000002', x.id, x.ord, now(), now()
FROM (VALUES
  ('c77cedd3-2800-7db3-aa91-68e188fa8864'::uuid, 1),   -- Percepção de futuro: 6 meses
  ('c77cedd3-2800-734f-862f-6434c4a8522c'::uuid, 2),   -- Percepção de futuro: 5 anos
  ('c77cedd3-2800-7a0a-9f2c-fdc5ebbc2220'::uuid, 3),   -- Percepção de futuro: 10 anos
  ('c77cedd3-2800-759a-bdea-45cd69d48dad'::uuid, 4),   -- Adesão
  ('c77cedd3-2800-721b-94d8-b5515b895753'::uuid, 5),   -- PHQ-9 (humor)
  ('c77cedd3-2800-78a5-a272-01c1e573fcc0'::uuid, 6),   -- GAD-7 (ansiedade)
  ('c77cedd3-2800-7556-8439-afa139a9bae3'::uuid, 7),   -- Epworth (sonolência)
  ('019bf31d-2ef0-74c5-8d99-0f355f1aa7cc'::uuid, 8),   -- Peso
  ('019bf31d-2ef0-71e4-a845-cdfaeedbb599'::uuid, 9),   -- Altura
  ('019c534a-afc3-70c4-82e0-bfde4b5b8f93'::uuid, 10),  -- Padrão alimentar atual
  ('c77cedd3-2800-7bb3-9e74-b76061206583'::uuid, 11),  -- Estratégia macro de movimento atual
  ('c77cedd3-2800-735f-bf28-c5d07d7d7092'::uuid, 12),  -- Qualidade percebida do sono
  ('c77cedd3-2800-7360-bf3c-5b4f28c660ef'::uuid, 13)   -- Fontes de stress percebidas
) AS x(id, ord)
JOIN score_items si ON si.id = x.id AND si.deleted_at IS NULL;

-- prep-b2 (Continuum aprofundamento) — núcleo + história + reconciliação + composição + social
INSERT INTO score_version_items (id, version_id, score_item_id, display_order, created_at, updated_at)
SELECT gen_random_uuid(), '11111111-2222-7333-8444-000000000003', x.id, x.ord, now(), now()
FROM (VALUES
  ('019bf31d-2ef0-78da-9d77-4e8258d3cf8e'::uuid, 1),   -- Uso atual de medicamentos
  ('019bf31d-2ef0-75aa-838b-0a0928acc4b1'::uuid, 2),   -- Histórico de medicamentos
  ('019bf31d-2ef0-7090-b843-99752e3c622d'::uuid, 3),   -- Suplementações utilizadas
  ('019bf31d-2ef0-74c5-8d99-0f355f1aa7cc'::uuid, 4),   -- Peso
  ('019bf31d-2ef0-71e4-a845-cdfaeedbb599'::uuid, 5),   -- Altura
  ('c77cedd3-2800-7aac-985c-c075f020e9e0'::uuid, 6),   -- Cintura (homem)
  ('c77cedd3-2800-7a74-99ad-56ca4b6dddc1'::uuid, 7),   -- Cintura (mulher)
  ('019bf31d-2ef0-7f17-a053-7f45b7162dd2'::uuid, 8),   -- Diabetes mellitus (crônica)
  ('019bf31d-2ef0-7de7-a98d-603c41d12ae6'::uuid, 9),   -- Doença cardiovascular (crônica)
  ('c77cedd3-2800-735f-bf28-c5d07d7d7092'::uuid, 10),  -- Qualidade percebida do sono
  ('019c53a6-f1a3-7704-9868-354859c750cd'::uuid, 11),  -- Tempo de sono
  ('019bf31d-2ef0-7a08-ab7e-d9c06d8d2103'::uuid, 12),  -- Tabaco
  ('019bf31d-2ef0-7eff-b67d-00e5eec8c2d0'::uuid, 13),  -- Álcool
  ('c77cedd3-2800-75fe-b483-ea0183478225'::uuid, 14),  -- Familiar: hipertensão
  ('c77cedd3-2800-7524-9049-f4559170db14'::uuid, 15),  -- Familiar: diabetes
  ('c77cedd3-2800-7e0c-b44f-ccf9c3b926ef'::uuid, 16),  -- Familiar: câncer
  ('c77cedd3-2800-70c7-942e-a1134e3aa05e'::uuid, 17),  -- Familiar: doença cardiovascular
  ('c77cedd3-2800-731d-88cb-2ac2422f8e86'::uuid, 18),  -- Familiar: doença renal
  ('019c2e64-1522-74db-a6b5-97595412d14a'::uuid, 19)   -- Situação familiar (atual)
) AS x(id, ord)
JOIN score_items si ON si.id = x.id AND si.deleted_at IS NULL;

COMMIT;
