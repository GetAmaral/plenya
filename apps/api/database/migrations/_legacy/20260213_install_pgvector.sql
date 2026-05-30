-- Migration: Install pgvector and create RAG semantic search schema
-- Author: Claude Sonnet 4.5
-- Date: 2026-02-13
-- Phase 1: Foundation - pgvector Setup

-- =============================================================================
-- PART 1: Install pgvector extension
-- =============================================================================

-- Install pgvector extension (requires PostgreSQL 11+)
-- Version: 0.8.1 (latest as of Feb 2026)
CREATE EXTENSION IF NOT EXISTS vector;

-- Verify installation
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'vector') THEN
        RAISE EXCEPTION 'pgvector extension installation failed';
    END IF;
    RAISE NOTICE '✅ pgvector extension installed successfully';
END $$;

-- =============================================================================
-- PART 2: Article Embeddings Table
-- =============================================================================

-- Table to store embeddings of article chunks
-- Each article is split into chunks (abstract + sliding window)
-- Each chunk gets its own embedding vector
CREATE TABLE article_embeddings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),

    -- Foreign key to articles table
    article_id UUID NOT NULL REFERENCES articles(id) ON DELETE CASCADE,

    -- Chunk metadata
    chunk_index INT NOT NULL, -- 0 = abstract, 1+ = content chunks
    chunk_text TEXT NOT NULL, -- The actual text of this chunk
    chunk_metadata JSONB DEFAULT '{}', -- {section: "abstract|methods|results", page_range: [1,3], word_count: 250}

    -- Embedding vector (OpenAI text-embedding-3-large: 1024 dimensions)
    embedding vector(1024),

    -- Timestamps
    created_at TIMESTAMP DEFAULT NOW(),

    -- Ensure unique chunks per article
    UNIQUE(article_id, chunk_index)
);

-- Indexes for efficient querying
CREATE INDEX idx_article_embeddings_article_id ON article_embeddings(article_id);

-- HNSW index for fast vector similarity search
-- Using ivfflat for now (pgvector 0.8.1 default)
-- Lists = sqrt(total_rows) heuristic, using 100 for ~10k expected chunks
-- Performance target: <50ms for similarity search on 50k chunks
CREATE INDEX idx_article_embeddings_embedding ON article_embeddings
    USING ivfflat (embedding vector_cosine_ops)
    WITH (lists = 100);

-- Comment
COMMENT ON TABLE article_embeddings IS
'Stores vector embeddings for article chunks. Each article is split into semantic chunks (abstract + sliding window). Used for RAG semantic search.';

COMMENT ON COLUMN article_embeddings.embedding IS
'1024-dimensional vector from OpenAI text-embedding-3-large. Cosine similarity used for semantic search.';

-- =============================================================================
-- PART 3: ScoreItem Embeddings Table
-- =============================================================================

-- Table to store embeddings of ScoreItems (for matching articles to parameters)
-- Each ScoreItem gets one embedding (combination of name + clinical fields)
CREATE TABLE score_item_embeddings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),

    -- Foreign key to score_items table (one-to-one)
    score_item_id UUID NOT NULL UNIQUE REFERENCES score_items(id) ON DELETE CASCADE,

    -- Embedding vector (OpenAI text-embedding-3-large: 1024 dimensions)
    embedding vector(1024),

    -- Source text used to generate embedding
    -- Combination: name + clinical_relevance + patient_explanation + conduct
    text_source TEXT NOT NULL,

    -- Timestamps
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Indexes
CREATE INDEX idx_score_item_embeddings_score_item_id ON score_item_embeddings(score_item_id);

-- HNSW index for vector similarity
-- Lists = 50 (expecting ~500 score items)
CREATE INDEX idx_score_item_embeddings_embedding ON score_item_embeddings
    USING ivfflat (embedding vector_cosine_ops)
    WITH (lists = 50);

-- Comment
COMMENT ON TABLE score_item_embeddings IS
'Stores vector embeddings for ScoreItems. Used to recommend articles for clinical parameters and discover which parameters an article covers.';

-- =============================================================================
-- PART 4: Embedding Processing Queue
-- =============================================================================

-- Queue for asynchronous embedding generation
-- Worker polls this table and processes pending jobs
CREATE TABLE embedding_queue (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),

    -- Entity to process
    entity_type VARCHAR(50) NOT NULL, -- 'article' | 'score_item'
    entity_id UUID NOT NULL,

    -- Status tracking
    status VARCHAR(20) DEFAULT 'pending', -- pending | processing | completed | failed
    retry_count INT DEFAULT 0,
    error_message TEXT,

    -- Timestamps
    created_at TIMESTAMP DEFAULT NOW(),
    processed_at TIMESTAMP,

    -- Constraints
    CHECK (entity_type IN ('article', 'score_item')),
    CHECK (status IN ('pending', 'processing', 'completed', 'failed')),
    CHECK (retry_count >= 0 AND retry_count <= 5) -- Max 5 retries
);

-- Indexes for efficient queue processing
CREATE INDEX idx_embedding_queue_status ON embedding_queue(status, created_at);
CREATE INDEX idx_embedding_queue_entity ON embedding_queue(entity_type, entity_id);

-- Comment
COMMENT ON TABLE embedding_queue IS
'Asynchronous queue for embedding generation. Worker processes pending jobs in background.';

-- =============================================================================
-- PART 5: Extend Articles Table
-- =============================================================================

-- Add metadata fields to track embedding status
ALTER TABLE articles
ADD COLUMN IF NOT EXISTS embedding_status VARCHAR(20) DEFAULT 'pending',
ADD COLUMN IF NOT EXISTS chunk_count INT DEFAULT 0,
ADD COLUMN IF NOT EXISTS last_embedded_at TIMESTAMP;

-- Index for filtering by embedding status
CREATE INDEX IF NOT EXISTS idx_articles_embedding_status ON articles(embedding_status);

-- Constraint
ALTER TABLE articles
ADD CONSTRAINT check_articles_embedding_status
CHECK (embedding_status IN ('pending', 'processing', 'completed', 'failed'));

-- Comment
COMMENT ON COLUMN articles.embedding_status IS
'Status of embedding generation: pending (not started), processing (in progress), completed (done), failed (error).';

COMMENT ON COLUMN articles.chunk_count IS
'Number of chunks generated for this article (typically 1 abstract + N content chunks).';

-- =============================================================================
-- PART 6: Enrich Article-ScoreItem Junction Table
-- =============================================================================

-- Add metadata to track confidence and auto-linking
ALTER TABLE article_score_items
ADD COLUMN IF NOT EXISTS confidence_score FLOAT DEFAULT 1.0,
ADD COLUMN IF NOT EXISTS auto_linked BOOLEAN DEFAULT false,
ADD COLUMN IF NOT EXISTS linked_at TIMESTAMP DEFAULT NOW(),
ADD COLUMN IF NOT EXISTS linked_by UUID REFERENCES users(id);

-- Constraints
ALTER TABLE article_score_items
ADD CONSTRAINT check_article_score_items_confidence
CHECK (confidence_score >= 0.0 AND confidence_score <= 1.0);

-- Index for filtering auto-linked associations
CREATE INDEX IF NOT EXISTS idx_article_score_items_auto_linked ON article_score_items(auto_linked);

-- Comment
COMMENT ON COLUMN article_score_items.confidence_score IS
'Cosine similarity score (0.0-1.0) if auto-linked by RAG, or 1.0 if manually linked. Threshold: 0.7 for auto-suggestions.';

COMMENT ON COLUMN article_score_items.auto_linked IS
'True if link was suggested by RAG semantic search, false if manually linked by user.';

COMMENT ON COLUMN article_score_items.linked_by IS
'User ID who created the link (NULL for system auto-links).';

-- =============================================================================
-- PART 7: API Usage Tracking (Optional but Recommended)
-- =============================================================================

-- Track OpenAI API usage for cost monitoring
CREATE TABLE api_usage_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),

    -- API details
    provider VARCHAR(50) NOT NULL, -- 'openai' | 'anthropic'
    model VARCHAR(100) NOT NULL, -- 'text-embedding-3-large' | 'claude-3-5-haiku-20241022'
    endpoint VARCHAR(100) NOT NULL, -- 'embeddings' | 'messages'

    -- Token usage
    input_tokens INT,
    output_tokens INT,
    total_tokens INT,

    -- Cost calculation (USD)
    cost_usd DECIMAL(10, 6), -- Calculated based on model pricing

    -- User tracking
    user_id UUID REFERENCES users(id),

    -- Metadata
    metadata JSONB DEFAULT '{}', -- {entity_type: 'article', entity_id: '...', operation: 'chunk_embedding'}

    -- Timestamps
    created_at TIMESTAMP DEFAULT NOW(),

    -- Constraints
    CHECK (provider IN ('openai', 'anthropic')),
    CHECK (total_tokens >= 0)
);

-- Indexes for cost analysis
CREATE INDEX idx_api_usage_logs_created_at ON api_usage_logs(created_at);
CREATE INDEX idx_api_usage_logs_provider_model ON api_usage_logs(provider, model);
CREATE INDEX idx_api_usage_logs_user_id ON api_usage_logs(user_id);

-- Comment
COMMENT ON TABLE api_usage_logs IS
'Tracks API usage for cost monitoring and debugging. OpenAI embeddings: $0.13/1M tokens.';

-- =============================================================================
-- PART 8: Verification & Summary
-- =============================================================================

DO $$
DECLARE
    v_embedding_dim INT;
BEGIN
    -- Verify vector dimension support
    SELECT 1024 INTO v_embedding_dim;

    RAISE NOTICE '=============================================================================';
    RAISE NOTICE '✅ RAG Semantic Search Schema Created Successfully';
    RAISE NOTICE '=============================================================================';
    RAISE NOTICE 'Tables created:';
    RAISE NOTICE '  - article_embeddings (vector storage for article chunks)';
    RAISE NOTICE '  - score_item_embeddings (vector storage for clinical parameters)';
    RAISE NOTICE '  - embedding_queue (async processing queue)';
    RAISE NOTICE '  - api_usage_logs (cost tracking)';
    RAISE NOTICE '';
    RAISE NOTICE 'Tables modified:';
    RAISE NOTICE '  - articles (added: embedding_status, chunk_count, last_embedded_at)';
    RAISE NOTICE '  - article_score_items (added: confidence_score, auto_linked, linked_at, linked_by)';
    RAISE NOTICE '';
    RAISE NOTICE 'Vector dimension: % (OpenAI text-embedding-3-large)', v_embedding_dim;
    RAISE NOTICE 'Index type: ivfflat (pgvector 0.8.1)';
    RAISE NOTICE 'Similarity metric: cosine';
    RAISE NOTICE '';
    RAISE NOTICE 'Next steps:';
    RAISE NOTICE '  1. Configure OPENAI_API_KEY in .env';
    RAISE NOTICE '  2. Implement EmbeddingService (Phase 2)';
    RAISE NOTICE '  3. Implement EmbeddingWorker (Phase 3)';
    RAISE NOTICE '  4. Backfill existing articles';
    RAISE NOTICE '=============================================================================';
END $$;
