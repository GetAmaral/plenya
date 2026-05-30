-- Add maxSelect field to score_subgroups
-- +goose Up
ALTER TABLE score_subgroups
ADD COLUMN max_select INTEGER NOT NULL DEFAULT 0;

-- +goose Down
ALTER TABLE score_subgroups
DROP COLUMN max_select;
