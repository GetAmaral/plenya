-- Add lab_request_template_id column to lab_requests table
-- This tracks which template was used to create the lab request

ALTER TABLE lab_requests
ADD COLUMN lab_request_template_id UUID REFERENCES lab_request_templates(id) ON DELETE SET NULL;

-- Add index for faster lookups
CREATE INDEX idx_lab_requests_template_id ON lab_requests(lab_request_template_id);

-- Add comment
COMMENT ON COLUMN lab_requests.lab_request_template_id IS 'ID do template utilizado para criar este pedido (opcional)';
