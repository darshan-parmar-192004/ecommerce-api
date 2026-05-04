-- =========================================================
-- 003_add_code_columns.up.sql
-- Add code columns for staging-based seeding
-- =========================================================

-- Add code columns to all tables
ALTER TABLE categories ADD COLUMN IF NOT EXISTS code TEXT UNIQUE;
ALTER TABLE customers ADD COLUMN IF NOT EXISTS code TEXT UNIQUE;
ALTER TABLE products ADD COLUMN IF NOT EXISTS code TEXT UNIQUE;
ALTER TABLE orders ADD COLUMN IF NOT EXISTS code TEXT UNIQUE;

-- Create indexes for FK resolution performance
CREATE INDEX IF NOT EXISTS idx_categories_code ON categories(code);
CREATE INDEX IF NOT EXISTS idx_customers_code ON customers(code);
CREATE INDEX IF NOT EXISTS idx_products_code ON products(code);
CREATE INDEX IF NOT EXISTS idx_orders_code ON orders(code);

-- Make code NOT NULL after ensuring uniqueness (code is required for all records)
ALTER TABLE categories ALTER COLUMN code SET NOT NULL;
ALTER TABLE customers ALTER COLUMN code SET NOT NULL;
ALTER TABLE products ALTER COLUMN code SET NOT NULL;
ALTER TABLE orders ALTER COLUMN code SET NOT NULL;
