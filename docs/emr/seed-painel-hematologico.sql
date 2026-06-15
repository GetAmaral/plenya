-- ============================================================================
-- Seed: Painel Hematológico — COMPLEMENTO dirigido. Só exames que NÃO estão no
-- Inicial nem no Completo, pra empilhar sem duplicar. Hoje: apenas Eritropoietina
-- (reflexo de hemograma alterado — poliglobulia/anemia; sem valor de triagem com
-- hemograma normal — ver curadoria). Cresce conforme a curadoria.
-- Idempotente. APLICAR EM DEV; prod só sob ordem.
-- ============================================================================
BEGIN;

-- Template (id fixo); ordem 4 (focado), Completo vai p/ 5 --------------------
INSERT INTO lab_request_templates (id, name, description, display_order, is_active, created_at, updated_at)
SELECT '0c9e7d10-0000-7000-8000-0000000000e5'::uuid, 'Painel Hematológico',
       'Complemento hematológico dirigido — empilha sobre o Inicial ou o Completo.',
       4, true, now(), now()
WHERE NOT EXISTS (SELECT 1 FROM lab_request_templates t WHERE t.name = 'Painel Hematológico');

UPDATE lab_request_templates SET display_order = 5, updated_at = now()
WHERE name = 'Painel Plenya Completo' AND display_order <> 5;

-- Itens (só Eritropoietina por ora) ------------------------------------------
INSERT INTO lab_request_template_tests (lab_request_template_id, lab_test_definition_id)
SELECT '0c9e7d10-0000-7000-8000-0000000000e5', d.id FROM lab_test_definitions d
WHERE d.code IN ('PLNERITRO01')  -- Eritropoietina
ON CONFLICT DO NOTHING;

COMMIT;
