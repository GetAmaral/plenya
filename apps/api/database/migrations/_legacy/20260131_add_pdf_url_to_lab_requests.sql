-- Add pdf_url column to lab_requests table
-- When PDF is generated, the lab request becomes read-only

ALTER TABLE lab_requests
ADD COLUMN pdf_url TEXT;

-- Add index for faster lookups of requests with/without PDF
CREATE INDEX idx_lab_requests_pdf_url ON lab_requests(pdf_url) WHERE pdf_url IS NOT NULL;

-- Add comment
COMMENT ON COLUMN lab_requests.pdf_url IS 'URL do PDF gerado. Quando preenchido, o pedido fica bloqueado para edição';
