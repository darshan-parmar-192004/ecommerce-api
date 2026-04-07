-- =========================================================
-- 003_add_roles.up.sql
-- Add Role-Based Access Control (UP Migration)
-- =========================================================

ALTER TABLE customers ADD COLUMN IF NOT EXISTS role VARCHAR(20) NOT NULL DEFAULT 'customer' CHECK (role IN ('customer', 'admin'));

-- Update existing customers to have 'customer' role
UPDATE customers SET role = 'customer' WHERE role IS NULL OR role = '';

-- Create index on role for faster queries
CREATE INDEX IF NOT EXISTS idx_customers_role ON customers(role);
