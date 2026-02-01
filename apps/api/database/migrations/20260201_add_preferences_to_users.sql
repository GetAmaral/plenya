-- Migration: Add preferences column to users table
-- Date: 2026-02-01
-- Description: Adiciona campo JSONB para armazenar preferências do usuário (viewport do mindmap, etc.)

-- Add preferences column
ALTER TABLE users
ADD COLUMN IF NOT EXISTS preferences JSONB NOT NULL DEFAULT '{}';

-- Create index for better performance on JSON queries
CREATE INDEX IF NOT EXISTS idx_users_preferences ON users USING gin(preferences);

-- Add comment
COMMENT ON COLUMN users.preferences IS 'Preferências do usuário em formato JSON (viewport do mindmap, configurações de interface, etc.)';
