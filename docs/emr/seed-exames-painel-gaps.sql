-- ============================================================================
-- Seed: exames faltantes p/ os painéis + ativação de Albumina + flags de sexo.
-- Decisões do Dr. (2026-06-11). TUSS confirmados online. Idempotente.
-- APLICAR EM DEV; prod só sob ordem. (Peptídeo C já criado em seed próprio.)
-- ============================================================================
BEGIN;

-- 1) Exames avulsos a criar ---------------------------------------------------
INSERT INTO lab_test_definitions
  (id, code, name, short_name, tuss_code, category, is_requestable, result_type,
   unit, fasting_hours, sex_applicability, description, display_order, is_active, created_at, updated_at)
SELECT * FROM (VALUES
  (uuid_generate_v7(), 'PLNGLIJEJ01', 'Glicemia de jejum', 'Gli jejum', '40302040', 'biochemistry', true, 'numeric', 'mg/dL', 8,  'all', 'Glicose plasmática de jejum.',                                  0, true, now(), now()),
  (uuid_generate_v7(), 'PLNINSJEJ01', 'Insulina (jejum)',  'Insulina',  '40316360', 'biochemistry', true, 'numeric', 'µUI/mL', 8, 'all', 'Insulina de jejum — sensibilidade insulínica / HOMA-IR.',       0, true, now(), now()),
  (uuid_generate_v7(), 'PLNERITRO01', 'Eritropoietina',    'EPO',       '40305295', 'hematology',   true, 'numeric', 'mUI/mL', NULL, 'all', 'Eritropoietina sérica.',                                       0, true, now(), now()),
  (uuid_generate_v7(), 'PLNUSGTIR01', 'Ultrassonografia de tireoide', 'USG Tireoide', '40901203', 'imaging', true, 'text', NULL, NULL, 'all', 'US de órgãos superficiais (tireoide).',                  0, true, now(), now()),
  (uuid_generate_v7(), 'PLNRXTORX01', 'Radiografia de tórax PA + perfil', 'RX Tórax PA+P', '40805026', 'imaging', true, 'text', NULL, NULL, 'all', 'RX de tórax, 2 incidências (PA e perfil).',          0, true, now(), now()),
  -- Curadoria 2026-06-12: TTPA pareado ao TP/INR (TUSS 40304639, verificado online)
  (uuid_generate_v7(), 'PLNTTPA0001', 'Tempo de tromboplastina parcial ativada (TTPa)', 'TTPa', '40304639', 'hematology', true, 'numeric', 'segundos', NULL, 'all', 'Via intrínseca da coagulação; pareado ao TP/INR.', 0, true, now(), now())
) AS v(id, code, name, short_name, tuss_code, category, is_requestable, result_type, unit, fasting_hours, sex_applicability, description, display_order, is_active, created_at, updated_at)
WHERE NOT EXISTS (SELECT 1 FROM lab_test_definitions d WHERE d.code = v.code);

-- 2) Ativar Albumina (já existe, estava desligada) ----------------------------
UPDATE lab_test_definitions SET is_requestable = true, updated_at = now()
WHERE code = 'PLN6885D35A' AND is_requestable = false;

-- 3) Flags de sexo (conservador) ----------------------------------------------
UPDATE lab_test_definitions SET sex_applicability = 'male', updated_at = now()
WHERE code IN ('PLN1BFE6CA3','PLNE9B271B9','PLN1B6B4299')          -- PSA Total, %fPSA, USG próstata
  AND sex_applicability <> 'male';

UPDATE lab_test_definitions SET sex_applicability = 'female', updated_at = now()
WHERE code IN ('PLNFDA4FECB','PLN56898402','PLNC3A4FA11','PLNEA1BE0BF')  -- CA-125, Mamografia, USG mamas, USG transvaginal
  AND sex_applicability <> 'female';

COMMIT;
