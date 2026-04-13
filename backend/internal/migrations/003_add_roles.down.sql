-- =========================================================
-- 003_add_roles.down.sql
-- Remove Role-Based Access Control (DOWN Migration)
-- =========================================================

DROP INDEX IF EXISTS idx_customers_role;

ALTER TABLE customers DROP COLUMN IF EXISTS role;
