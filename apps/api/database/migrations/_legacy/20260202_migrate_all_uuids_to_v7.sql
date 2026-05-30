-- Migration: Convert All UUIDs from v4 to v7
-- CRITICAL: Requires full database backup before execution
-- Estimated time: ~5 minutes for current data volume
-- Autor: Claude Sonnet 4.5
-- Data: 2026-02-02

-- ============================================================================
-- PHASE 1: Create UUID Mapping Table
-- ============================================================================

CREATE TABLE IF NOT EXISTS uuid_migration_map (
    old_id UUID NOT NULL,
    new_id UUID NOT NULL,
    table_name VARCHAR(100) NOT NULL,
    migrated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (old_id, table_name)
);

CREATE INDEX IF NOT EXISTS idx_uuid_migration_new_id ON uuid_migration_map(new_id);
CREATE INDEX IF NOT EXISTS idx_uuid_migration_table ON uuid_migration_map(table_name);

-- ============================================================================
-- PHASE 2: Helper Function - Update FK References
-- ============================================================================

CREATE OR REPLACE FUNCTION migrate_fk_column(
    p_table_name TEXT,
    p_column_name TEXT,
    p_referenced_table TEXT
)
RETURNS VOID
LANGUAGE plpgsql
AS $$
DECLARE
    v_sql TEXT;
    v_updated_count BIGINT;
BEGIN
    v_sql := format(
        'UPDATE %I SET %I = m.new_id FROM uuid_migration_map m ' ||
        'WHERE %I.%I = m.old_id AND m.table_name = %L',
        p_table_name, p_column_name,
        p_table_name, p_column_name, p_referenced_table
    );

    EXECUTE v_sql;
    GET DIAGNOSTICS v_updated_count = ROW_COUNT;

    RAISE NOTICE 'Updated %.%: % rows', p_table_name, p_column_name, v_updated_count;
END;
$$;
-- ============================================================================
-- PHASE 3: Migrate Tables (Ordered by Dependencies)
-- ============================================================================

DO $$
DECLARE
    v_start_time TIMESTAMP;
    v_total_migrated INT := 0;
    v_table RECORD;
BEGIN
    v_start_time := clock_timestamp();
    RAISE NOTICE '🚀 Starting UUID v4 → v7 migration at %', v_start_time;

    -- Define migration order (respecting FK dependencies)
    FOR v_table IN
        SELECT table_name, fk_columns::TEXT[] AS fk_cols, fk_tables::TEXT[] AS fk_refs
        FROM (VALUES
            -- Level 1: No dependencies
            ('score_groups', ARRAY[]::TEXT[], ARRAY[]::TEXT[]),
            ('lab_request_templates', ARRAY[]::TEXT[], ARRAY[]::TEXT[]),
            ('articles', ARRAY[]::TEXT[], ARRAY[]::TEXT[]),
            ('anamnesis_templates', ARRAY[]::TEXT[], ARRAY[]::TEXT[]),
            
            -- Level 2: First-level dependencies
            ('users', ARRAY[]::TEXT[], ARRAY[]::TEXT[]),
            ('score_subgroups', ARRAY['group_id'], ARRAY['score_groups']),
            
            -- Level 3: Second-level dependencies
            ('patients', ARRAY['user_id'], ARRAY['users']),
            ('score_items', ARRAY['subgroup_id', 'parent_item_id'], ARRAY['score_subgroups', 'score_items']),
            ('lab_test_definitions', ARRAY['parent_test_id'], ARRAY['lab_test_definitions']),
            
            -- Level 4: Third-level dependencies
            ('anamnesis_template_items', ARRAY['anamnesis_template_id', 'score_item_id'], ARRAY['anamnesis_templates', 'score_items']),
            ('anamnesis', ARRAY['patient_id', 'author_id', 'anamnesis_template_id'], ARRAY['patients', 'users', 'anamnesis_templates']),
            ('lab_requests', ARRAY['patient_id', 'doctor_id', 'lab_request_template_id'], ARRAY['patients', 'users', 'lab_request_templates']),
            ('lab_results', ARRAY['patient_id', 'requesting_doctor_id'], ARRAY['patients', 'users']),
            ('appointments', ARRAY['patient_id', 'doctor_id', 'anamnesis_id'], ARRAY['patients', 'users', 'anamnesis']),
            ('prescriptions', ARRAY['patient_id', 'doctor_id'], ARRAY['patients', 'users']),
            ('audit_logs', ARRAY['user_id'], ARRAY['users']),
            ('score_levels', ARRAY['item_id'], ARRAY['score_items']),
            
            -- Level 5: Junction tables
            ('anamnesis_items', ARRAY['anamnesis_id', 'score_item_id'], ARRAY['anamnesis', 'score_items']),
            ('lab_result_values', ARRAY['lab_result_id', 'lab_test_definition_id'], ARRAY['lab_results', 'lab_test_definitions'])
        ) AS t(table_name, fk_columns, fk_tables)
    LOOP
        BEGIN
            -- Disable triggers for this table
            EXECUTE format('ALTER TABLE %I DISABLE TRIGGER ALL', v_table.table_name);
            
            -- Update foreign keys first
            IF v_table.fk_cols IS NOT NULL THEN
                FOR i IN 1..array_length(v_table.fk_cols, 1) LOOP
                    PERFORM migrate_fk_column(v_table.table_name, v_table.fk_cols[i], v_table.fk_refs[i]);
                END LOOP;
            END IF;
            
            -- Migrate primary keys
            EXECUTE format(
                'INSERT INTO uuid_migration_map (old_id, new_id, table_name) ' ||
                'SELECT id, uuid_v7_from_timestamp(created_at, id), %L FROM %I WHERE id IS NOT NULL',
                v_table.table_name, v_table.table_name
            );
            
            EXECUTE format(
                'UPDATE %I t SET id = m.new_id FROM uuid_migration_map m WHERE t.id = m.old_id AND m.table_name = %L',
                v_table.table_name, v_table.table_name
            );
            
            -- Re-enable triggers
            EXECUTE format('ALTER TABLE %I ENABLE TRIGGER ALL', v_table.table_name);
            
            RAISE NOTICE '✅ Migrated %', v_table.table_name;
            
        EXCEPTION WHEN OTHERS THEN
            EXECUTE format('ALTER TABLE %I ENABLE TRIGGER ALL', v_table.table_name);
            RAISE EXCEPTION 'Failed migrating %: %', v_table.table_name, SQLERRM;
        END;
    END LOOP;
    
    -- Handle junction tables without PKs (just update FKs)
    PERFORM migrate_fk_column('lab_test_score_mappings', 'lab_test_id', 'lab_test_definitions');
    PERFORM migrate_fk_column('lab_test_score_mappings', 'score_item_id', 'score_items');
    RAISE NOTICE '✅ Updated lab_test_score_mappings FKs';
    
    PERFORM migrate_fk_column('article_score_items', 'article_id', 'articles');
    PERFORM migrate_fk_column('article_score_items', 'score_item_id', 'score_items');
    RAISE NOTICE '✅ Updated article_score_items FKs';
    
    PERFORM migrate_fk_column('lab_request_template_tests', 'lab_request_template_id', 'lab_request_templates');
    PERFORM migrate_fk_column('lab_request_template_tests', 'lab_test_definition_id', 'lab_test_definitions');
    RAISE NOTICE '✅ Updated lab_request_template_tests FKs';
    
    -- Special: users.selected_patient_id FK
    PERFORM migrate_fk_column('users', 'selected_patient_id', 'patients');
    RAISE NOTICE '✅ Updated users.selected_patient_id FK';
    
    SELECT COUNT(*) INTO v_total_migrated FROM uuid_migration_map;
    
    RAISE NOTICE '';
    RAISE NOTICE '═══════════════════════════════════════════════════════';
    RAISE NOTICE '✅ Migration Complete!';
    RAISE NOTICE '   Total UUIDs migrated: %', v_total_migrated;
    RAISE NOTICE '   Duration: % seconds', EXTRACT(EPOCH FROM (clock_timestamp() - v_start_time));
    RAISE NOTICE '═══════════════════════════════════════════════════════';
    RAISE NOTICE '';
    RAISE NOTICE '⚠️  IMPORTANT: Keep uuid_migration_map table for rollback!';
    
END;
$$;

-- ============================================================================
-- PHASE 4: Validation Queries
-- ============================================================================

DO $$
DECLARE
    v_table RECORD;
    v_v4_count INT;
    v_v7_count INT;
BEGIN
    RAISE NOTICE '';
    RAISE NOTICE 'UUID Version Distribution After Migration:';
    RAISE NOTICE '────────────────────────────────────────────────────';

    FOR v_table IN
        SELECT table_name FROM information_schema.columns
        WHERE column_name = 'id' AND table_schema = 'public'
        AND table_name != 'uuid_migration_map'
        ORDER BY table_name
    LOOP
        EXECUTE format(
            'SELECT
                COUNT(*) FILTER (WHERE (get_byte(uuid_send(id), 6) >> 4) = 4),
                COUNT(*) FILTER (WHERE (get_byte(uuid_send(id), 6) >> 4) = 7)
            FROM %I',
            v_table.table_name
        ) INTO v_v4_count, v_v7_count;

        IF v_v4_count + v_v7_count > 0 THEN
            RAISE NOTICE '%-30s  v4: %, v7: % %',
                v_table.table_name,
                v_v4_count,
                v_v7_count,
                CASE WHEN v_v4_count > 0 THEN '⚠️  STILL HAS v4!' ELSE '✅' END;
        END IF;
    END LOOP;

    RAISE NOTICE '────────────────────────────────────────────────────';
END;
$$;
