-- OAuth Support and Nullable CPF
-- Created: 2026-02-03
-- Description: Add OAuth provider fields to users table and make CPF nullable in patients table

-- Add OAuth fields to users table
ALTER TABLE users ALTER COLUMN password_hash DROP NOT NULL;
ALTER TABLE users ADD COLUMN oauth_provider VARCHAR(20);
ALTER TABLE users ADD COLUMN oauth_provider_id VARCHAR(255);
ALTER TABLE users ADD COLUMN oauth_picture_url TEXT;

-- Create indexes for OAuth lookups
CREATE INDEX idx_users_oauth_provider ON users(oauth_provider) WHERE oauth_provider IS NOT NULL;
CREATE UNIQUE INDEX idx_users_oauth_unique ON users(oauth_provider, oauth_provider_id) WHERE oauth_provider IS NOT NULL AND oauth_provider_id IS NOT NULL;

-- Make CPF nullable in patients table (already nullable in code, but enforce in DB)
ALTER TABLE patients ALTER COLUMN cpf DROP NOT NULL;

-- Add check constraint for OAuth provider enum
ALTER TABLE users ADD CONSTRAINT chk_oauth_provider CHECK (oauth_provider IS NULL OR oauth_provider IN ('google', 'apple'));
