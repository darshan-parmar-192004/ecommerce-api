-- =========================================================
-- 001_initial_schema.sql
-- Database Schema
-- =========================================================

-- =========================================================
-- 1️⃣ DROP TABLES (Safe Re-run)
-- =========================================================

DROP TABLE IF EXISTS order_items CASCADE;
DROP TABLE IF EXISTS orders CASCADE;
DROP TABLE IF EXISTS inventory CASCADE;
DROP TABLE IF EXISTS products CASCADE;
DROP TABLE IF EXISTS customers CASCADE;
DROP TABLE IF EXISTS categories CASCADE;

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
    password_hash TEXT NOT NULL
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
