ALTER TABLE customers DROP COLUMN IF EXISTS role;
DROP INDEX IF EXISTS idx_customers_role;
