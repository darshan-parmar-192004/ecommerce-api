-- =========================================================
-- 003_add_code_columns.down.sql
-- Remove code columns from all tables
-- =========================================================

-- Drop indexes first
DROP INDEX IF EXISTS idx_categories_code;
DROP INDEX IF EXISTS idx_customers_code;
DROP INDEX IF EXISTS idx_products_code;
DROP INDEX IF EXISTS idx_orders_code;

-- Drop code columns
ALTER TABLE categories DROP COLUMN IF EXISTS code;
ALTER TABLE customers DROP COLUMN IF EXISTS code;
ALTER TABLE products DROP COLUMN IF EXISTS code;
ALTER TABLE orders DROP COLUMN IF EXISTS code;
