-- Migration: Update all table DEFAULTs to uuid_generate_v7()
-- Ensures all future records use UUID v7
-- Data: 2026-02-02

-- Core tables
ALTER TABLE users ALTER COLUMN id SET DEFAULT uuid_generate_v7();
ALTER TABLE patients ALTER COLUMN id SET DEFAULT uuid_generate_v7();
ALTER TABLE audit_logs ALTER COLUMN id SET DEFAULT uuid_generate_v7();

-- Medical records
ALTER TABLE anamnesis ALTER COLUMN id SET DEFAULT uuid_generate_v7();
ALTER TABLE anamnesis_items ALTER COLUMN id SET DEFAULT uuid_generate_v7();
ALTER TABLE appointments ALTER COLUMN id SET DEFAULT uuid_generate_v7();
ALTER TABLE prescriptions ALTER COLUMN id SET DEFAULT uuid_generate_v7();

-- Lab system
ALTER TABLE lab_results ALTER COLUMN id SET DEFAULT uuid_generate_v7();
ALTER TABLE lab_requests ALTER COLUMN id SET DEFAULT uuid_generate_v7();
ALTER TABLE lab_request_templates ALTER COLUMN id SET DEFAULT uuid_generate_v7();
ALTER TABLE lab_test_definitions ALTER COLUMN id SET DEFAULT uuid_generate_v7();
ALTER TABLE lab_result_values ALTER COLUMN id SET DEFAULT uuid_generate_v7();

-- Score system
ALTER TABLE score_groups ALTER COLUMN id SET DEFAULT uuid_generate_v7();
ALTER TABLE score_subgroups ALTER COLUMN id SET DEFAULT uuid_generate_v7();
ALTER TABLE score_items ALTER COLUMN id SET DEFAULT uuid_generate_v7();
ALTER TABLE score_levels ALTER COLUMN id SET DEFAULT uuid_generate_v7();

-- Articles
ALTER TABLE articles ALTER COLUMN id SET DEFAULT uuid_generate_v7();

-- Anamnesis Templates
ALTER TABLE anamnesis_templates ALTER COLUMN id SET DEFAULT uuid_generate_v7();
ALTER TABLE anamnesis_template_items ALTER COLUMN id SET DEFAULT uuid_generate_v7();

-- Lab request template tests (if it has an ID column)
-- Note: This is a junction table, might not have an ID
DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_name = 'lab_request_template_tests' AND column_name = 'id'
    ) THEN
        ALTER TABLE lab_request_template_tests ALTER COLUMN id SET DEFAULT uuid_generate_v7();
    END IF;
END $$;

-- Verify changes
SELECT
    table_name,
    column_default
FROM information_schema.columns
WHERE
    table_schema = 'public'
    AND column_name = 'id'
    AND column_default IS NOT NULL
ORDER BY table_name;

DO $$
BEGIN
    RAISE NOTICE '✅ All table DEFAULTs updated to uuid_generate_v7()';
END $$;
