-- Migration: Create score_item_enrichment_preparation table
-- Armazena chunks selecionados do RAG para posterior geração de enrichment via Claude
-- ETAPA 1 do plano de enrichment científico avançado (SEM IA - apenas preparação)

-- +goose Up
CREATE TABLE IF NOT EXISTS score_item_enrichment_preparation (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    score_item_id UUID UNIQUE NOT NULL REFERENCES score_items(id) ON DELETE CASCADE,

    -- Chunks selecionados via RAG (array de objetos JSON)
    -- Estrutura: [{article_id, article_title, article_year, journal, chunk_index, chunk_text, section, similarity, word_count}, ...]
    selected_chunks JSONB NOT NULL DEFAULT '[]'::jsonb,

    -- Metadata agregado da seleção
    -- Estrutura: {total_chunks, articles_count, avg_similarity, sections_distribution: {results: 8, discussion: 6, methods: 4, introduction: 2}}
    metadata JSONB DEFAULT '{}'::jsonb,

    -- Status do processamento
    -- ready: pronto para geração (chunks já selecionados)
    -- processing: geração em andamento
    -- completed: geração completada
    -- failed: erro na geração
    status VARCHAR(20) NOT NULL DEFAULT 'ready' CHECK (status IN ('ready', 'processing', 'completed', 'failed')),

    -- Mensagem de erro (se status = failed)
    error_message TEXT,

    -- Timestamps
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Índices
CREATE INDEX idx_enrichment_prep_score_item ON score_item_enrichment_preparation(score_item_id);
CREATE INDEX idx_enrichment_prep_status ON score_item_enrichment_preparation(status);
CREATE INDEX idx_enrichment_prep_created ON score_item_enrichment_preparation(created_at DESC);

-- Índice GIN para busca em selected_chunks JSONB
CREATE INDEX idx_enrichment_prep_chunks ON score_item_enrichment_preparation USING gin(selected_chunks);

-- Trigger para atualizar updated_at automaticamente
CREATE OR REPLACE FUNCTION update_enrichment_prep_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_enrichment_prep_updated_at
    BEFORE UPDATE ON score_item_enrichment_preparation
    FOR EACH ROW
    EXECUTE FUNCTION update_enrichment_prep_timestamp();

-- Comentários
COMMENT ON TABLE score_item_enrichment_preparation IS 'Preparação de chunks científicos para enrichment via Claude (Etapa 1 - sem IA)';
COMMENT ON COLUMN score_item_enrichment_preparation.selected_chunks IS 'Array JSONB com 15-20 chunks dos artigos mais relevantes via RAG';
COMMENT ON COLUMN score_item_enrichment_preparation.metadata IS 'Estatísticas da seleção: total_chunks, articles_count, avg_similarity, sections_distribution';
COMMENT ON COLUMN score_item_enrichment_preparation.status IS 'ready (chunks prontos) | processing (gerando) | completed (pronto) | failed (erro)';

-- +goose Down
DROP TRIGGER IF EXISTS trigger_enrichment_prep_updated_at ON score_item_enrichment_preparation;
DROP FUNCTION IF EXISTS update_enrichment_prep_timestamp();
DROP TABLE IF EXISTS score_item_enrichment_preparation CASCADE;
