-- Add selected_patient_id to users table
-- This allows users to have a "current working patient" context

ALTER TABLE users
ADD COLUMN selected_patient_id UUID;

-- Add index for performance
CREATE INDEX idx_users_selected_patient_id ON users(selected_patient_id);

-- Add foreign key constraint with SET NULL on delete
ALTER TABLE users
ADD CONSTRAINT fk_users_selected_patient
FOREIGN KEY (selected_patient_id)
REFERENCES patients(id)
ON DELETE SET NULL;

-- Add comment
COMMENT ON COLUMN users.selected_patient_id IS 'ID of the currently selected patient for this user context';
