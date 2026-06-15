-- ============================================================================
-- Seed: 3 templates macro de Pedido de Exames (Inicial / Acompanhamento / Completo)
-- Inicial ancorado nos 4 pedidos reais do Dr. + decisões de 2026-06-11.
-- Completo = todos requisitáveis (não-genética, sem os 7 lixo do hemograma,
--            sem Albumina avulsa — coberta pela Eletroforese de proteínas).
-- Idempotente. APLICAR EM DEV; prod só sob ordem.
-- ============================================================================
BEGIN;

-- Templates (ids fixos p/ idempotência) ---------------------------------------
INSERT INTO lab_request_templates (id, name, description, display_order, is_active, created_at, updated_at)
SELECT * FROM (VALUES
  ('0c9e7d10-0000-7000-8000-0000000000e2'::uuid, 'Painel Inicial',        'Baseline de 1ª consulta (longevidade/performance).', 1, true, now(), now()),
  ('0c9e7d10-0000-7000-8000-0000000000e3'::uuid, 'Painel Acompanhamento', 'Rotina enxuta de acompanhamento.',                   2, true, now(), now()),
  ('0c9e7d10-0000-7000-8000-0000000000e1'::uuid, 'Painel Plenya Completo','Avaliação ampla — todo o laboratório + imagem.',      3, true, now(), now())
) AS v(id, name, description, display_order, is_active, created_at, updated_at)
WHERE NOT EXISTS (SELECT 1 FROM lab_request_templates t WHERE t.name = v.name);

-- ACOMPANHAMENTO (~25) --------------------------------------------------------
INSERT INTO lab_request_template_tests (lab_request_template_id, lab_test_definition_id)
SELECT '0c9e7d10-0000-7000-8000-0000000000e3', d.id FROM lab_test_definitions d
WHERE d.code IN (
  'PLNDB36E69C','PLNGLIJEJ01','PLNINSJEJ01','PLNCPEPTD01','PLN3FC5EDA6','PLN6626DC67','PLN543993C6',
  'PLN357DC859','PLNF1744AB9','PLNF0D05E40','PLNBBB9A8C7','PLN74FF1276','PLN27D9E113','PLNFC95CB43',
  'PLNB54EDF5C','PLN2F53AA29','PLN50BB057B','PLNCEFB97FD','PLN1BF562ED','PLN9B054BBD','PLND2C05835',
  'PLN4BF6D992','PLNBB0DDBD4','PLN726D788F','PLN20FCAA74'
) ON CONFLICT DO NOTHING;

-- INICIAL (~66) = Acompanhamento + baseline clínico do Dr. + imagem -----------
INSERT INTO lab_request_template_tests (lab_request_template_id, lab_test_definition_id)
SELECT '0c9e7d10-0000-7000-8000-0000000000e2', d.id FROM lab_test_definitions d
WHERE d.code IN (
  -- herda os de Acompanhamento
  'PLNDB36E69C','PLNGLIJEJ01','PLNINSJEJ01','PLNCPEPTD01','PLN3FC5EDA6','PLN6626DC67','PLN543993C6',
  'PLN357DC859','PLNF1744AB9','PLNF0D05E40','PLNBBB9A8C7','PLN74FF1276','PLN27D9E113','PLNFC95CB43',
  'PLNB54EDF5C','PLN2F53AA29','PLN50BB057B','PLNCEFB97FD','PLN1BF562ED','PLN9B054BBD','PLND2C05835',
  'PLN4BF6D992','PLNBB0DDBD4','PLN726D788F','PLN20FCAA74',
  -- adições do baseline real (laboratório)
  'PLN590DC6E7','PLN6885D35A','PLN33F0400D','PLNA8451657','PLN39ED7B60','PLN0F03A2B1','PLND7C2752F',
  'PLN97EC4BBF','PLN9E2106EC','PLN57E09E4C','PLN00994110','PLN82DBE091','PLNFF2F4B44','PLNF44D1468',
  'PLNBE9B2E32','PLNB453495C','PLN93895A61','PLNA6CA66FE','PLNC01E9624','PLN7FD1B365','PLN1D1920D4',
  'PLNFD087E3B','PLN64222EAD','PLN26FF1748','PLNDE1A5575','PLN20DC8CDB','PLNCFC25332','PLN741297DA',
  'PLN459B19CD','PLN7B3753DD','PLNC61B934A','PLN42D34D96','PLN5B8A10AC',
  -- imagem (Inicial + Completo)
  'PLND1D39793','PLN278D0810','PLN1CF5711F','PLN1B6B4299','PLN56898402','PLNC3A4FA11','PLNEA1BE0BF','PLNUSGTIR01'
) ON CONFLICT DO NOTHING;

-- COMPLETO (~140) = todos requisitáveis, não-genética, sem lixo nem albumina avulsa
INSERT INTO lab_request_template_tests (lab_request_template_id, lab_test_definition_id)
SELECT '0c9e7d10-0000-7000-8000-0000000000e1', d.id FROM lab_test_definitions d
WHERE d.is_active AND d.deleted_at IS NULL AND d.is_requestable AND d.category <> 'genetics'
  AND d.code NOT IN ('BLASTOS','IST','METAMIELO','MIELO','PDW','PROMIELO','VPM','PLN6885D35A',
                     'PLNB384883B',   -- D-dímero: vive só no Painel Risco Vascular (complemento)
                     'PLNERITRO01')   -- Eritropoietina: vive só no Painel Hematológico (complemento)
ON CONFLICT DO NOTHING;

-- Remover templates lixo ------------------------------------------------------
DELETE FROM lab_request_templates WHERE name IN ('teste','teste2');

COMMIT;
