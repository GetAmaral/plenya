-- Create lab_request_templates table
CREATE TABLE IF NOT EXISTS lab_request_templates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(200) NOT NULL UNIQUE,
    description TEXT,
    display_order INTEGER NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- Create many-to-many relationship table
CREATE TABLE IF NOT EXISTS lab_request_template_tests (
    lab_request_template_id UUID NOT NULL REFERENCES lab_request_templates(id) ON DELETE CASCADE,
    lab_test_definition_id UUID NOT NULL REFERENCES lab_test_definitions(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (lab_request_template_id, lab_test_definition_id)
);

-- Indexes
CREATE INDEX idx_lab_request_templates_name ON lab_request_templates(name);
CREATE INDEX idx_lab_request_templates_is_active ON lab_request_templates(is_active);
CREATE INDEX idx_lab_request_templates_deleted_at ON lab_request_templates(deleted_at);

CREATE INDEX idx_lab_request_template_tests_template ON lab_request_template_tests(lab_request_template_id);
CREATE INDEX idx_lab_request_template_tests_test ON lab_request_template_tests(lab_test_definition_id);

-- Comments
COMMENT ON TABLE lab_request_templates IS 'Templates pré-configurados de pedidos de exames';
COMMENT ON TABLE lab_request_template_tests IS 'Relação many-to-many entre templates e exames';
