-- =========================================================
-- 004_add_timestamps.down.sql
-- Remove created_at and updated_at from all tables
-- =========================================================

ALTER TABLE inventory DROP COLUMN IF EXISTS updated_at;
ALTER TABLE inventory DROP COLUMN IF EXISTS created_at;

ALTER TABLE order_items DROP COLUMN IF EXISTS updated_at;
ALTER TABLE order_items DROP COLUMN IF EXISTS created_at;

ALTER TABLE orders DROP COLUMN IF EXISTS updated_at;
ALTER TABLE orders DROP COLUMN IF EXISTS created_at;

ALTER TABLE products DROP COLUMN IF EXISTS updated_at;

ALTER TABLE customers DROP COLUMN IF EXISTS updated_at;

ALTER TABLE categories DROP COLUMN IF EXISTS updated_at;
ALTER TABLE categories DROP COLUMN IF EXISTS created_at;
