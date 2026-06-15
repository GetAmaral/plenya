-- ============================================================================
-- Seed: Painel Risco Vascular — COMPLEMENTO. Só exames que NÃO estão no
-- Inicial nem no Completo, pra empilhar (Inicial+Vascular ou Completo+Vascular)
-- sem pedir nada em duplicata. Hoje: apenas D-dímero (rule-out de trombose,
-- inespecífico demais p/ rastreio amplo — ver curadoria). Cresce conforme a
-- curadoria tira exames "dirigidos" do Completo.
-- Idempotente. APLICAR EM DEV; prod só sob ordem.
-- ============================================================================
BEGIN;

-- Template (id fixo); ordem 3 (focado), Completo vai p/ 4 --------------------
INSERT INTO lab_request_templates (id, name, description, display_order, is_active, created_at, updated_at)
SELECT '0c9e7d10-0000-7000-8000-0000000000e4'::uuid, 'Painel Risco Vascular',
       'Complemento de risco vascular — empilha sobre o Inicial ou o Completo quando há foco cardiovascular.',
       3, true, now(), now()
WHERE NOT EXISTS (SELECT 1 FROM lab_request_templates t WHERE t.name = 'Painel Risco Vascular');

UPDATE lab_request_templates SET display_order = 5, updated_at = now()
WHERE name = 'Painel Plenya Completo' AND display_order <> 5;

-- Itens (só D-dímero por ora) -------------------------------------------------
INSERT INTO lab_request_template_tests (lab_request_template_id, lab_test_definition_id)
SELECT '0c9e7d10-0000-7000-8000-0000000000e4', d.id FROM lab_test_definitions d
WHERE d.code IN ('PLNB384883B')  -- D-dímero
ON CONFLICT DO NOTHING;

COMMIT;
