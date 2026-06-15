-- ============================================================================
-- Seed: Painel Inflamação — COMPLEMENTO dirigido. Só exames que NÃO estão no
-- Inicial nem no Completo, pra empilhar sem duplicar. Hoje: apenas VHS
-- (inespecífico; PCR-us já cobre inflamação nos amplos). Cresce na curadoria.
-- Idempotente. APLICAR EM DEV; prod só sob ordem.
-- ============================================================================
BEGIN;

INSERT INTO lab_request_templates (id, name, description, display_order, is_active, created_at, updated_at)
SELECT '0c9e7d10-0000-7000-8000-0000000000e6'::uuid, 'Painel Inflamação',
       'Complemento inflamatório dirigido — empilha sobre o Inicial ou o Completo.',
       5, true, now(), now()
WHERE NOT EXISTS (SELECT 1 FROM lab_request_templates t WHERE t.name = 'Painel Inflamação');

INSERT INTO lab_request_template_tests (lab_request_template_id, lab_test_definition_id)
SELECT '0c9e7d10-0000-7000-8000-0000000000e6', d.id FROM lab_test_definitions d
WHERE d.code IN ('PLN7472EF93')  -- VHS (ESR)
ON CONFLICT DO NOTHING;

COMMIT;
