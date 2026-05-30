-- Add Method System tables
-- Migration created: 2026-02-15

-- Create methods table
CREATE TABLE IF NOT EXISTS methods (
    id UUID PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    short_name VARCHAR(20) NOT NULL,
    description TEXT,
    version VARCHAR(20),
    color VARCHAR(7),
    "order" INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- Create unique indexes for methods
CREATE UNIQUE INDEX IF NOT EXISTS idx_method_name ON methods(name) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX IF NOT EXISTS idx_method_short_name ON methods(short_name) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_method_order ON methods("order");
CREATE INDEX IF NOT EXISTS methods_deleted_at_idx ON methods(deleted_at);

-- Create method_letters table
CREATE TABLE IF NOT EXISTS method_letters (
    id UUID PRIMARY KEY,
    code VARCHAR(10) NOT NULL,
    name VARCHAR(300) NOT NULL,
    description TEXT,
    clinical_relevance TEXT,
    patient_explanation TEXT,
    conduct TEXT,
    last_review TIMESTAMP,
    color VARCHAR(7),
    icon VARCHAR(50),
    "order" INTEGER NOT NULL DEFAULT 0,
    method_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    CONSTRAINT fk_method_letters_method FOREIGN KEY (method_id)
        REFERENCES methods(id) ON DELETE CASCADE
);

-- Create indexes for method_letters
CREATE INDEX IF NOT EXISTS idx_method_letter_code ON method_letters(code);
CREATE INDEX IF NOT EXISTS idx_method_letter_order ON method_letters("order");
CREATE INDEX IF NOT EXISTS idx_method_letter_method ON method_letters(method_id);
CREATE INDEX IF NOT EXISTS method_letters_deleted_at_idx ON method_letters(deleted_at);

-- Create method_pillars table
CREATE TABLE IF NOT EXISTS method_pillars (
    id UUID PRIMARY KEY,
    name VARCHAR(300) NOT NULL,
    description TEXT,
    clinical_relevance TEXT,
    patient_explanation TEXT,
    conduct TEXT,
    last_review TIMESTAMP,
    "order" INTEGER NOT NULL DEFAULT 0,
    letter_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    CONSTRAINT fk_method_pillars_letter FOREIGN KEY (letter_id)
        REFERENCES method_letters(id) ON DELETE CASCADE
);

-- Create indexes for method_pillars
CREATE INDEX IF NOT EXISTS idx_method_pillar_order ON method_pillars("order");
CREATE INDEX IF NOT EXISTS idx_method_pillar_letter ON method_pillars(letter_id);
CREATE INDEX IF NOT EXISTS method_pillars_deleted_at_idx ON method_pillars(deleted_at);

-- Create junction table for M:N relationship (ScoreItem <-> MethodPillar)
CREATE TABLE IF NOT EXISTS score_item_method_pillars (
    score_item_id UUID NOT NULL,
    method_pillar_id UUID NOT NULL,
    PRIMARY KEY (score_item_id, method_pillar_id),
    CONSTRAINT fk_score_item_method_pillars_score_item
        FOREIGN KEY (score_item_id) REFERENCES score_items(id) ON DELETE CASCADE,
    CONSTRAINT fk_score_item_method_pillars_pillar
        FOREIGN KEY (method_pillar_id) REFERENCES method_pillars(id) ON DELETE CASCADE
);

-- Create indexes for junction table
CREATE INDEX IF NOT EXISTS idx_junction_score_item ON score_item_method_pillars(score_item_id);
CREATE INDEX IF NOT EXISTS idx_junction_pillar ON score_item_method_pillars(method_pillar_id);
