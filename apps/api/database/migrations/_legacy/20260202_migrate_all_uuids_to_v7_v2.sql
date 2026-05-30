-- Migration: Convert All UUIDs from v4 to v7 (Version 2 - Simplified)
-- CRITICAL: Requires full database backup before execution
-- Autor: Claude Sonnet 4.5
-- Data: 2026-02-02

-- Create mapping table
CREATE TABLE IF NOT EXISTS uuid_migration_map (
    old_id UUID NOT NULL,
    new_id UUID NOT NULL,
    table_name VARCHAR(100) NOT NULL,
    migrated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (old_id, table_name)
);

CREATE INDEX IF NOT EXISTS idx_uuid_migration_new_id ON uuid_migration_map(new_id);

DO $$
DECLARE
    v_start_time TIMESTAMP;
    v_count INT;
BEGIN
    v_start_time := clock_timestamp();
    RAISE NOTICE '🚀 Starting UUID v4 → v7 migration';

    -- Level 1: Independent tables (no FKs to update before migration)

    -- score_groups
    ALTER TABLE score_groups DISABLE TRIGGER ALL;
    INSERT INTO uuid_migration_map (old_id, new_id, table_name)
    SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'score_groups' FROM score_groups;
    UPDATE score_groups sg SET id = m.new_id FROM uuid_migration_map m
    WHERE sg.id = m.old_id AND m.table_name = 'score_groups';
    ALTER TABLE score_groups ENABLE TRIGGER ALL;
    GET DIAGNOSTICS v_count = ROW_COUNT;
    RAISE NOTICE '✅ score_groups: % migrated', v_count;

    -- lab_request_templates
    ALTER TABLE lab_request_templates DISABLE TRIGGER ALL;
    INSERT INTO uuid_migration_map (old_id, new_id, table_name)
    SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'lab_request_templates' FROM lab_request_templates;
    UPDATE lab_request_templates t SET id = m.new_id FROM uuid_migration_map m
    WHERE t.id = m.old_id AND m.table_name = 'lab_request_templates';
    ALTER TABLE lab_request_templates ENABLE TRIGGER ALL;
    GET DIAGNOSTICS v_count = ROW_COUNT;
    RAISE NOTICE '✅ lab_request_templates: % migrated', v_count;

    -- articles
    ALTER TABLE articles DISABLE TRIGGER ALL;
    INSERT INTO uuid_migration_map (old_id, new_id, table_name)
    SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'articles' FROM articles;
    UPDATE articles t SET id = m.new_id FROM uuid_migration_map m
    WHERE t.id = m.old_id AND m.table_name = 'articles';
    ALTER TABLE articles ENABLE TRIGGER ALL;
    GET DIAGNOSTICS v_count = ROW_COUNT;
    RAISE NOTICE '✅ articles: % migrated', v_count;

    -- anamnesis_templates
    ALTER TABLE anamnesis_templates DISABLE TRIGGER ALL;
    INSERT INTO uuid_migration_map (old_id, new_id, table_name)
    SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'anamnesis_templates' FROM anamnesis_templates;
    UPDATE anamnesis_templates t SET id = m.new_id FROM uuid_migration_map m
    WHERE t.id = m.old_id AND m.table_name = 'anamnesis_templates';
    ALTER TABLE anamnesis_templates ENABLE TRIGGER ALL;
    GET DIAGNOSTICS v_count = ROW_COUNT;
    RAISE NOTICE '✅ anamnesis_templates: % migrated', v_count;

    -- Level 2: users (CRITICAL - many FK references)

    ALTER TABLE users DISABLE TRIGGER ALL;
    INSERT INTO uuid_migration_map (old_id, new_id, table_name)
    SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'users' FROM users;
    UPDATE users t SET id = m.new_id FROM uuid_migration_map m
    WHERE t.id = m.old_id AND m.table_name = 'users';
    ALTER TABLE users ENABLE TRIGGER ALL;
    GET DIAGNOSTICS v_count = ROW_COUNT;
    RAISE NOTICE '✅ users: % migrated', v_count;

    -- score_subgroups (FK: score_groups)

    ALTER TABLE score_subgroups DISABLE TRIGGER ALL;
    -- Update FK
    UPDATE score_subgroups ss SET group_id = m.new_id FROM uuid_migration_map m
    WHERE ss.group_id = m.old_id AND m.table_name = 'score_groups';
    -- Migrate PKs
    INSERT INTO uuid_migration_map (old_id, new_id, table_name)
    SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'score_subgroups' FROM score_subgroups;
    UPDATE score_subgroups t SET id = m.new_id FROM uuid_migration_map m
    WHERE t.id = m.old_id AND m.table_name = 'score_subgroups';
    ALTER TABLE score_subgroups ENABLE TRIGGER ALL;
    GET DIAGNOSTICS v_count = ROW_COUNT;
    RAISE NOTICE '✅ score_subgroups: % migrated', v_count;

    -- Level 3: Second-level dependencies

    -- patients (FK: users)
    ALTER TABLE patients DISABLE TRIGGER ALL;
    UPDATE patients p SET user_id = m.new_id FROM uuid_migration_map m
    WHERE p.user_id = m.old_id AND m.table_name = 'users';
    INSERT INTO uuid_migration_map (old_id, new_id, table_name)
    SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'patients' FROM patients;
    UPDATE patients t SET id = m.new_id FROM uuid_migration_map m
    WHERE t.id = m.old_id AND m.table_name = 'patients';
    ALTER TABLE patients ENABLE TRIGGER ALL;
    GET DIAGNOSTICS v_count = ROW_COUNT;
    RAISE NOTICE '✅ patients: % migrated', v_count;

    -- score_items (FK: score_subgroups, parent_item_id self-referential)
    ALTER TABLE score_items DISABLE TRIGGER ALL;
    UPDATE score_items si SET subgroup_id = m.new_id FROM uuid_migration_map m
    WHERE si.subgroup_id = m.old_id AND m.table_name = 'score_subgroups';
    INSERT INTO uuid_migration_map (old_id, new_id, table_name)
    SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'score_items' FROM score_items;
    UPDATE score_items t SET id = m.new_id FROM uuid_migration_map m
    WHERE t.id = m.old_id AND m.table_name = 'score_items';
    -- Update self-referential FK
    UPDATE score_items si SET parent_item_id = m.new_id FROM uuid_migration_map m
    WHERE si.parent_item_id = m.old_id AND m.table_name = 'score_items';
    ALTER TABLE score_items ENABLE TRIGGER ALL;
    GET DIAGNOSTICS v_count = ROW_COUNT;
    RAISE NOTICE '✅ score_items: % migrated', v_count;

    -- lab_test_definitions (self-referential)
    ALTER TABLE lab_test_definitions DISABLE TRIGGER ALL;
    INSERT INTO uuid_migration_map (old_id, new_id, table_name)
    SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'lab_test_definitions' FROM lab_test_definitions;
    UPDATE lab_test_definitions t SET id = m.new_id FROM uuid_migration_map m
    WHERE t.id = m.old_id AND m.table_name = 'lab_test_definitions';
    UPDATE lab_test_definitions ltd SET parent_test_id = m.new_id FROM uuid_migration_map m
    WHERE ltd.parent_test_id = m.old_id AND m.table_name = 'lab_test_definitions';
    ALTER TABLE lab_test_definitions ENABLE TRIGGER ALL;
    GET DIAGNOSTICS v_count = ROW_COUNT;
    RAISE NOTICE '✅ lab_test_definitions: % migrated', v_count;

    -- Level 4: Third-level dependencies

    -- anamnesis_template_items
    ALTER TABLE anamnesis_template_items DISABLE TRIGGER ALL;
    UPDATE anamnesis_template_items ati SET anamnesis_template_id = m.new_id FROM uuid_migration_map m
    WHERE ati.anamnesis_template_id = m.old_id AND m.table_name = 'anamnesis_templates';
    UPDATE anamnesis_template_items ati SET score_item_id = m.new_id FROM uuid_migration_map m
    WHERE ati.score_item_id = m.old_id AND m.table_name = 'score_items';
    INSERT INTO uuid_migration_map (old_id, new_id, table_name)
    SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'anamnesis_template_items' FROM anamnesis_template_items;
    UPDATE anamnesis_template_items t SET id = m.new_id FROM uuid_migration_map m
    WHERE t.id = m.old_id AND m.table_name = 'anamnesis_template_items';
    ALTER TABLE anamnesis_template_items ENABLE TRIGGER ALL;
    GET DIAGNOSTICS v_count = ROW_COUNT;
    RAISE NOTICE '✅ anamnesis_template_items: % migrated', v_count;

    -- anamnesis
    ALTER TABLE anamnesis DISABLE TRIGGER ALL;
    UPDATE anamnesis a SET patient_id = m.new_id FROM uuid_migration_map m
    WHERE a.patient_id = m.old_id AND m.table_name = 'patients';
    UPDATE anamnesis a SET author_id = m.new_id FROM uuid_migration_map m
    WHERE a.author_id = m.old_id AND m.table_name = 'users';
    UPDATE anamnesis a SET anamnesis_template_id = m.new_id FROM uuid_migration_map m
    WHERE a.anamnesis_template_id = m.old_id AND m.table_name = 'anamnesis_templates';
    INSERT INTO uuid_migration_map (old_id, new_id, table_name)
    SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'anamnesis' FROM anamnesis;
    UPDATE anamnesis t SET id = m.new_id FROM uuid_migration_map m
    WHERE t.id = m.old_id AND m.table_name = 'anamnesis';
    ALTER TABLE anamnesis ENABLE TRIGGER ALL;
    GET DIAGNOSTICS v_count = ROW_COUNT;
    RAISE NOTICE '✅ anamnesis: % migrated', v_count;

    -- lab_requests
    ALTER TABLE lab_requests DISABLE TRIGGER ALL;
    UPDATE lab_requests lr SET patient_id = m.new_id FROM uuid_migration_map m
    WHERE lr.patient_id = m.old_id AND m.table_name = 'patients';
    UPDATE lab_requests lr SET doctor_id = m.new_id FROM uuid_migration_map m
    WHERE lr.doctor_id = m.old_id AND m.table_name = 'users';
    UPDATE lab_requests lr SET lab_request_template_id = m.new_id FROM uuid_migration_map m
    WHERE lr.lab_request_template_id = m.old_id AND m.table_name = 'lab_request_templates';
    INSERT INTO uuid_migration_map (old_id, new_id, table_name)
    SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'lab_requests' FROM lab_requests;
    UPDATE lab_requests t SET id = m.new_id FROM uuid_migration_map m
    WHERE t.id = m.old_id AND m.table_name = 'lab_requests';
    ALTER TABLE lab_requests ENABLE TRIGGER ALL;
    GET DIAGNOSTICS v_count = ROW_COUNT;
    RAISE NOTICE '✅ lab_requests: % migrated', v_count;

    -- lab_results
    ALTER TABLE lab_results DISABLE TRIGGER ALL;
    UPDATE lab_results lr SET patient_id = m.new_id FROM uuid_migration_map m
    WHERE lr.patient_id = m.old_id AND m.table_name = 'patients';
    UPDATE lab_results lr SET requesting_doctor_id = m.new_id FROM uuid_migration_map m
    WHERE lr.requesting_doctor_id = m.old_id AND m.table_name = 'users';
    INSERT INTO uuid_migration_map (old_id, new_id, table_name)
    SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'lab_results' FROM lab_results;
    UPDATE lab_results t SET id = m.new_id FROM uuid_migration_map m
    WHERE t.id = m.old_id AND m.table_name = 'lab_results';
    ALTER TABLE lab_results ENABLE TRIGGER ALL;
    GET DIAGNOSTICS v_count = ROW_COUNT;
    RAISE NOTICE '✅ lab_results: % migrated', v_count;

    -- appointments
    ALTER TABLE appointments DISABLE TRIGGER ALL;
    UPDATE appointments a SET patient_id = m.new_id FROM uuid_migration_map m
    WHERE a.patient_id = m.old_id AND m.table_name = 'patients';
    UPDATE appointments a SET doctor_id = m.new_id FROM uuid_migration_map m
    WHERE a.doctor_id = m.old_id AND m.table_name = 'users';
    UPDATE appointments a SET anamnesis_id = m.new_id FROM uuid_migration_map m
    WHERE a.anamnesis_id = m.old_id AND m.table_name = 'anamnesis';
    INSERT INTO uuid_migration_map (old_id, new_id, table_name)
    SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'appointments' FROM appointments;
    UPDATE appointments t SET id = m.new_id FROM uuid_migration_map m
    WHERE t.id = m.old_id AND m.table_name = 'appointments';
    ALTER TABLE appointments ENABLE TRIGGER ALL;
    GET DIAGNOSTICS v_count = ROW_COUNT;
    RAISE NOTICE '✅ appointments: % migrated', v_count;

    -- prescriptions
    ALTER TABLE prescriptions DISABLE TRIGGER ALL;
    UPDATE prescriptions p SET patient_id = m.new_id FROM uuid_migration_map m
    WHERE p.patient_id = m.old_id AND m.table_name = 'patients';
    UPDATE prescriptions p SET doctor_id = m.new_id FROM uuid_migration_map m
    WHERE p.doctor_id = m.old_id AND m.table_name = 'users';
    INSERT INTO uuid_migration_map (old_id, new_id, table_name)
    SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'prescriptions' FROM prescriptions;
    UPDATE prescriptions t SET id = m.new_id FROM uuid_migration_map m
    WHERE t.id = m.old_id AND m.table_name = 'prescriptions';
    ALTER TABLE prescriptions ENABLE TRIGGER ALL;
    GET DIAGNOSTICS v_count = ROW_COUNT;
    RAISE NOTICE '✅ prescriptions: % migrated', v_count;

    -- audit_logs
    ALTER TABLE audit_logs DISABLE TRIGGER ALL;
    UPDATE audit_logs al SET user_id = m.new_id FROM uuid_migration_map m
    WHERE al.user_id = m.old_id AND m.table_name = 'users';
    INSERT INTO uuid_migration_map (old_id, new_id, table_name)
    SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'audit_logs' FROM audit_logs;
    UPDATE audit_logs t SET id = m.new_id FROM uuid_migration_map m
    WHERE t.id = m.old_id AND m.table_name = 'audit_logs';
    ALTER TABLE audit_logs ENABLE TRIGGER ALL;
    GET DIAGNOSTICS v_count = ROW_COUNT;
    RAISE NOTICE '✅ audit_logs: % migrated', v_count;

    -- score_levels
    ALTER TABLE score_levels DISABLE TRIGGER ALL;
    UPDATE score_levels sl SET item_id = m.new_id FROM uuid_migration_map m
    WHERE sl.item_id = m.old_id AND m.table_name = 'score_items';
    INSERT INTO uuid_migration_map (old_id, new_id, table_name)
    SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'score_levels' FROM score_levels;
    UPDATE score_levels t SET id = m.new_id FROM uuid_migration_map m
    WHERE t.id = m.old_id AND m.table_name = 'score_levels';
    ALTER TABLE score_levels ENABLE TRIGGER ALL;
    GET DIAGNOSTICS v_count = ROW_COUNT;
    RAISE NOTICE '✅ score_levels: % migrated', v_count;

    -- Level 5: Junction tables

    -- anamnesis_items (if exists - new table)
    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'anamnesis_items') THEN
        ALTER TABLE anamnesis_items DISABLE TRIGGER ALL;
        UPDATE anamnesis_items ai SET anamnesis_id = m.new_id FROM uuid_migration_map m
        WHERE ai.anamnesis_id = m.old_id AND m.table_name = 'anamnesis';
        UPDATE anamnesis_items ai SET score_item_id = m.new_id FROM uuid_migration_map m
        WHERE ai.score_item_id = m.old_id AND m.table_name = 'score_items';
        INSERT INTO uuid_migration_map (old_id, new_id, table_name)
        SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'anamnesis_items' FROM anamnesis_items WHERE id IS NOT NULL;
        UPDATE anamnesis_items t SET id = m.new_id FROM uuid_migration_map m
        WHERE t.id = m.old_id AND m.table_name = 'anamnesis_items';
        ALTER TABLE anamnesis_items ENABLE TRIGGER ALL;
        GET DIAGNOSTICS v_count = ROW_COUNT;
        RAISE NOTICE '✅ anamnesis_items: % migrated', v_count;
    END IF;

    -- lab_result_values
    ALTER TABLE lab_result_values DISABLE TRIGGER ALL;
    UPDATE lab_result_values lrv SET lab_result_id = m.new_id FROM uuid_migration_map m
    WHERE lrv.lab_result_id = m.old_id AND m.table_name = 'lab_results';
    UPDATE lab_result_values lrv SET lab_test_definition_id = m.new_id FROM uuid_migration_map m
    WHERE lrv.lab_test_definition_id = m.old_id AND m.table_name = 'lab_test_definitions';
    INSERT INTO uuid_migration_map (old_id, new_id, table_name)
    SELECT id, uuid_v7_from_timestamp(COALESCE(created_at, NOW())::TIMESTAMP, id), 'lab_result_values' FROM lab_result_values;
    UPDATE lab_result_values t SET id = m.new_id FROM uuid_migration_map m
    WHERE t.id = m.old_id AND m.table_name = 'lab_result_values';
    ALTER TABLE lab_result_values ENABLE TRIGGER ALL;
    GET DIAGNOSTICS v_count = ROW_COUNT;
    RAISE NOTICE '✅ lab_result_values: % migrated', v_count;

    -- Junction tables without IDs (just update FKs)

    UPDATE lab_test_score_mappings ltsm SET lab_test_id = m.new_id FROM uuid_migration_map m
    WHERE ltsm.lab_test_id = m.old_id AND m.table_name = 'lab_test_definitions';
    UPDATE lab_test_score_mappings ltsm SET score_item_id = m.new_id FROM uuid_migration_map m
    WHERE ltsm.score_item_id = m.old_id AND m.table_name = 'score_items';
    RAISE NOTICE '✅ lab_test_score_mappings: FKs updated';

    UPDATE article_score_items asi SET article_id = m.new_id FROM uuid_migration_map m
    WHERE asi.article_id = m.old_id AND m.table_name = 'articles';
    UPDATE article_score_items asi SET score_item_id = m.new_id FROM uuid_migration_map m
    WHERE asi.score_item_id = m.old_id AND m.table_name = 'score_items';
    RAISE NOTICE '✅ article_score_items: FKs updated';

    UPDATE lab_request_template_tests lrtt SET lab_request_template_id = m.new_id FROM uuid_migration_map m
    WHERE lrtt.lab_request_template_id = m.old_id AND m.table_name = 'lab_request_templates';
    UPDATE lab_request_template_tests lrtt SET lab_test_definition_id = m.new_id FROM uuid_migration_map m
    WHERE lrtt.lab_test_definition_id = m.old_id AND m.table_name = 'lab_test_definitions';
    RAISE NOTICE '✅ lab_request_template_tests: FKs updated';

    -- Special: users.selected_patient_id
    UPDATE users u SET selected_patient_id = m.new_id FROM uuid_migration_map m
    WHERE u.selected_patient_id = m.old_id AND m.table_name = 'patients';
    RAISE NOTICE '✅ users.selected_patient_id: FK updated';

    RAISE NOTICE '';
    RAISE NOTICE '═══════════════════════════════════════';
    RAISE NOTICE '✅ Migration Complete!';
    RAISE NOTICE '   Total UUIDs migrated: %', (SELECT COUNT(*) FROM uuid_migration_map);
    RAISE NOTICE '   Duration: % seconds', EXTRACT(EPOCH FROM (clock_timestamp() - v_start_time));
    RAISE NOTICE '═══════════════════════════════════════';

END;
$$;
