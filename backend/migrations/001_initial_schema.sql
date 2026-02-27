-- =========================================================
-- 001_initial_schema_and_import.sql
-- Schema + Staging + Data Import (Idempotent)
-- =========================================================

BEGIN;

-- =========================================================
-- 1️⃣ DROP TABLES (Safe Re-run)
-- =========================================================

DROP TABLE IF EXISTS order_items CASCADE;
DROP TABLE IF EXISTS orders CASCADE;
DROP TABLE IF EXISTS inventory CASCADE;
DROP TABLE IF EXISTS products CASCADE;
DROP TABLE IF EXISTS customers CASCADE;
DROP TABLE IF EXISTS categories CASCADE;

DROP TABLE IF EXISTS staging_categories;
DROP TABLE IF EXISTS staging_customers;
DROP TABLE IF EXISTS staging_products;
DROP TABLE IF EXISTS staging_orders;
DROP TABLE IF EXISTS staging_order_items;
DROP TABLE IF EXISTS staging_inventory;
DROP TABLE IF EXISTS import_warnings;

-- =========================================================
-- 2️⃣ PRODUCTION TABLES
-- =========================================================

CREATE TABLE categories (
    category_id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    parent_category_id TEXT REFERENCES categories(category_id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);

CREATE TABLE customers (
    customer_id TEXT PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    country TEXT,
    phone VARCHAR(50),
    created_at TIMESTAMP NOT NULL,
    status VARCHAR(50),
    password_hash TEXT
);

CREATE TABLE products (
    product_id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    category_id TEXT NOT NULL REFERENCES categories(category_id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE,
    price NUMERIC NOT NULL CHECK (price > 0),
    description TEXT,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE orders (
    order_id TEXT PRIMARY KEY,
    customer_id TEXT NOT NULL REFERENCES customers(customer_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    order_date TIMESTAMP NOT NULL,
    status TEXT,
    total_amount NUMERIC NOT NULL CHECK (total_amount >= 0),
    shipping_address TEXT
);

CREATE TABLE order_items (
    order_item_id TEXT PRIMARY KEY,
    order_id TEXT NOT NULL REFERENCES orders(order_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    product_id TEXT NOT NULL REFERENCES products(product_id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE,
    quantity INT NOT NULL CHECK (quantity >= 0),
    unit_price NUMERIC NOT NULL CHECK (unit_price >= 0)
);

CREATE TABLE inventory (
    product_id TEXT NOT NULL REFERENCES products(product_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    warehouse_id VARCHAR(50) NOT NULL,
    quantity INT NOT NULL CHECK (quantity >= 0),
    last_updated TIMESTAMP,
    PRIMARY KEY (product_id, warehouse_id)
);

-- =========================================================
-- 3️⃣ WARNING TABLE
-- =========================================================

CREATE TABLE import_warnings (
    id SERIAL PRIMARY KEY,
    table_name TEXT NOT NULL,
    record_id TEXT,
    issue TEXT NOT NULL,
    logged_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =========================================================
-- 4️⃣ STAGING TABLES (ALL TEXT)
-- =========================================================

CREATE TABLE staging_categories (
    category_id TEXT,
    name TEXT,
    parent_category_id TEXT
);

CREATE TABLE staging_customers (
    customer_id TEXT,
    email TEXT,
    name TEXT,
    country TEXT,
    phone TEXT,
    created_at TEXT,
    status TEXT
);

CREATE TABLE staging_products (
    product_id TEXT,
    name TEXT,
    category_id TEXT,
    price TEXT,
    description TEXT,
    created_at TEXT
);

CREATE TABLE staging_orders (
    order_id TEXT,
    customer_id TEXT,
    order_date TEXT,
    status TEXT,
    total_amount TEXT,
    shipping_address TEXT
);

CREATE TABLE staging_order_items (
    order_item_id TEXT,
    order_id TEXT,
    product_id TEXT,
    quantity TEXT,
    unit_price TEXT
);

CREATE TABLE staging_inventory (
    product_id TEXT,
    warehouse_id TEXT,
    quantity TEXT,
    last_updated TEXT
);

COMMIT;

-- =========================================================
-- 5️⃣ COPY CSV INTO STAGING
-- =========================================================

\copy staging_categories FROM '/home/darshan.parmar/Desktop/Documentation/ecommerce-api/backend/datasets/ecommerce/categories.csv' CSV HEADER
\copy staging_customers FROM '/home/darshan.parmar/Desktop/Documentation/ecommerce-api/backend/datasets/ecommerce/customers.csv' CSV HEADER
\copy staging_products FROM '/home/darshan.parmar/Desktop/Documentation/ecommerce-api/backend/datasets/ecommerce/products.csv' CSV HEADER
\copy staging_orders FROM '/home/darshan.parmar/Desktop/Documentation/ecommerce-api/backend/datasets/ecommerce/orders.csv' CSV HEADER
\copy staging_order_items FROM '/home/darshan.parmar/Desktop/Documentation/ecommerce-api/backend/datasets/ecommerce/order_items.csv' CSV HEADER
\copy staging_inventory FROM '/home/darshan.parmar/Desktop/Documentation/ecommerce-api/backend/datasets/ecommerce/inventory.csv' CSV HEADER

-- =========================================================
-- 6️⃣ DATA CLEANING + MIGRATION
-- =========================================================

-- Categories
INSERT INTO categories
SELECT category_id, name, parent_category_id
FROM staging_categories
ON CONFLICT DO NOTHING;

-- Customers (basic email validation)
INSERT INTO import_warnings (table_name, record_id, issue)
SELECT 'customers', customer_id, 'Invalid email'
FROM staging_customers
WHERE email NOT LIKE '%@%';

INSERT INTO customers (
    customer_id, email, name, country,
    phone, created_at, status, password_hash
)
SELECT
    customer_id,
    email,
    name,
    country,
    phone,
    created_at::timestamp,
    status,
    NULL
FROM staging_customers
WHERE email LIKE '%@%'
ON CONFLICT (email) DO NOTHING;

-- Products (skip negative price)
INSERT INTO import_warnings (table_name, record_id, issue)
SELECT 'products', product_id, 'Invalid price'
FROM staging_products
WHERE price::numeric <= 0;

INSERT INTO products
SELECT
    product_id,
    name,
    category_id,
    price::numeric,
    description,
    created_at::timestamp
FROM staging_products
WHERE price::numeric > 0
ON CONFLICT DO NOTHING;

-- Orders (skip if customer missing)
INSERT INTO orders
SELECT
    o.order_id,
    o.customer_id,
    o.order_date::timestamp,
    o.status,
    o.total_amount::numeric,
    o.shipping_address
FROM staging_orders o
JOIN customers c ON o.customer_id = c.customer_id
ON CONFLICT DO NOTHING;

-- Order Items (skip invalid quantity)
INSERT INTO order_items (
    order_item_id,
    order_id,
    product_id,
    quantity,
    unit_price
)
SELECT
    oi.order_item_id,
    oi.order_id,
    oi.product_id,
    oi.quantity::int,
    oi.unit_price::numeric
FROM staging_order_items oi
JOIN orders o ON oi.order_id = o.order_id
JOIN products p ON oi.product_id = p.product_id
WHERE oi.quantity::int >= 0
ON CONFLICT DO NOTHING;

-- Inventory
INSERT INTO inventory
SELECT
    product_id,
    warehouse_id,
    quantity::int,
    last_updated::timestamp
FROM staging_inventory
WHERE quantity::int >= 0
ON CONFLICT DO NOTHING;

-- =========================================================
-- 7️⃣ CLEAN STAGING
-- =========================================================

TRUNCATE staging_categories;
TRUNCATE staging_customers;
TRUNCATE staging_products;
TRUNCATE staging_orders;
TRUNCATE staging_order_items;
TRUNCATE staging_inventory;
