-- Migration: Create score_item_review_history table for audit trail
-- Armazena snapshots antes/depois de alterações em ScoreItems
-- Habilita rollback e auditoria completa de enriquecimento LLM

-- +goose Up
CREATE TABLE IF NOT EXISTS score_item_review_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    score_item_id UUID NOT NULL REFERENCES score_items(id) ON DELETE CASCADE,

    -- Tipo de revisão
    review_type VARCHAR(50) NOT NULL CHECK (review_type IN ('manual', 'llm_enrichment', 'pubmed_fetch', 'article_link')),

    -- Snapshots JSON
    before_snapshot JSONB NOT NULL,
    after_snapshot JSONB NOT NULL,

    -- Metadata de enriquecimento LLM
    tier VARCHAR(20) CHECK (tier IN ('preserve', 'enrich', 'rewrite')),
    confidence_score DOUBLE PRECISION CHECK (confidence_score >= 0 AND confidence_score <= 1),
    model_used VARCHAR(100), -- ex: claude-sonnet-4-5-20250929

    -- Revisão humana
    reviewed_by UUID REFERENCES users(id) ON DELETE SET NULL,
    approved BOOLEAN DEFAULT NULL, -- NULL=pending, true=approved, false=rejected
    review_notes TEXT,

    -- Timestamps
    reviewed_at TIMESTAMP NOT NULL DEFAULT NOW(),
    approved_at TIMESTAMP,

    -- Índices
    CONSTRAINT fk_score_item FOREIGN KEY (score_item_id) REFERENCES score_items(id) ON DELETE CASCADE
);

-- Índices para performance
CREATE INDEX idx_review_history_score_item ON score_item_review_history(score_item_id);
CREATE INDEX idx_review_history_type ON score_item_review_history(review_type);
CREATE INDEX idx_review_history_tier ON score_item_review_history(tier);
CREATE INDEX idx_review_history_approved ON score_item_review_history(approved);
CREATE INDEX idx_review_history_reviewed_at ON score_item_review_history(reviewed_at DESC);

-- Comentários
COMMENT ON TABLE score_item_review_history IS 'Audit trail completo de alterações em ScoreItems, incluindo enriquecimento LLM';
COMMENT ON COLUMN score_item_review_history.before_snapshot IS 'Estado dos campos clínicos ANTES da alteração (JSON)';
COMMENT ON COLUMN score_item_review_history.after_snapshot IS 'Estado dos campos clínicos DEPOIS da alteração (JSON)';
COMMENT ON COLUMN score_item_review_history.tier IS 'Estratégia de processamento LLM: preserve/enrich/rewrite';
COMMENT ON COLUMN score_item_review_history.confidence_score IS 'Confiança do modelo LLM (0-1)';

-- +goose Down
DROP TABLE IF EXISTS score_item_review_history CASCADE;
