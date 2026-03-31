-- Add role column to customers table
ALTER TABLE customers ADD COLUMN IF NOT EXISTS role VARCHAR(20) NOT NULL DEFAULT 'customer';

-- Update existing migrated customers to have 'customer' role (they already have 'N/A' or bcrypt hash as password)
UPDATE customers SET role = 'customer' WHERE role IS NULL OR role = '';

-- Create index on role for faster queries
CREATE INDEX IF NOT EXISTS idx_customers_role ON customers(role);

-- Grant admin role to specific customer emails (update as needed)
-- Example: UPDATE customers SET role = 'admin' WHERE email = 'admin@example.com';
