-- Add password_hash column with default hashed password
-- Default password: password123

ALTER TABLE customers ADD COLUMN IF NOT EXISTS password_hash VARCHAR(255);

-- Default password: password123 (bcrypt hash for "password123")
UPDATE customers 
SET password_hash = '$2a$10$oujE1BE8/qy5.wsbH3WP1u7g7Nil1x1heMF/0J4qZRuyQxyZq9M8a'
WHERE password_hash IS NULL OR password_hash = '';

-- Make password_hash NOT NULL for future inserts
ALTER TABLE customers ALTER COLUMN password_hash SET NOT NULL;
