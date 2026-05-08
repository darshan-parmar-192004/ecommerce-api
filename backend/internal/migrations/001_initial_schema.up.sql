-- =========================================================
-- 001_initial_schema.up.sql
-- Database Schema (UP Migration)
-- =========================================================

CREATE TABLE IF NOT EXISTS categories (
    category_id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    parent_category_id TEXT REFERENCES categories(category_id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS customers (
    customer_id TEXT PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    country TEXT,
    phone VARCHAR(50),
    created_at TIMESTAMP NOT NULL,
    status VARCHAR(50),
    password_hash TEXT,
    role VARCHAR(20) NOT NULL DEFAULT 'customer' CHECK (role IN ('customer', 'admin'))
);

CREATE TABLE IF NOT EXISTS products (
    product_id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    category_id TEXT NOT NULL REFERENCES categories(category_id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE,
    price NUMERIC NOT NULL CHECK (price > 0),
    description TEXT,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
    order_id TEXT PRIMARY KEY,
    customer_id TEXT NOT NULL REFERENCES customers(customer_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    order_date TIMESTAMP NOT NULL,
    status TEXT,
    total_amount NUMERIC NOT NULL CHECK (total_amount >= 0),
    shipping_address TEXT
);

CREATE TABLE IF NOT EXISTS order_items (
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

CREATE TABLE IF NOT EXISTS inventory (
    product_id TEXT NOT NULL REFERENCES products(product_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    warehouse_id VARCHAR(50) NOT NULL,
    quantity INT NOT NULL CHECK (quantity >= 0),
    last_updated TIMESTAMP,
    PRIMARY KEY (product_id, warehouse_id)
);
