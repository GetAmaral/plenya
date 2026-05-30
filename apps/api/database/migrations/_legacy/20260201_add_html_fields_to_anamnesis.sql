-- Add HTML fields to anamnesis table for dual storage (plain text + HTML)
-- Created: 2026-02-01
-- Purpose: Store both plain text (for search/indexing) and HTML (for formatted display)

-- Add content_html column (HTML formatted content)
ALTER TABLE anamnesis
ADD COLUMN content_html TEXT;

-- Add summary_html column (HTML formatted summary)
ALTER TABLE anamnesis
ADD COLUMN summary_html TEXT;

-- Comment columns for documentation
COMMENT ON COLUMN anamnesis.content IS 'Plain text content for search and indexing';
COMMENT ON COLUMN anamnesis.content_html IS 'HTML formatted content for display';
COMMENT ON COLUMN anamnesis.summary IS 'Plain text summary for search and indexing';
COMMENT ON COLUMN anamnesis.summary_html IS 'HTML formatted summary for display';
