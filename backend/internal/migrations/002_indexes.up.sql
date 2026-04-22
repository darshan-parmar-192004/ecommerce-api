-- =========================================================
-- 002_indexes.up.sql
-- Database Indexes (UP Migration)
-- =========================================================

CREATE INDEX idx_products_category_id ON products(category_id);
CREATE INDEX idx_products_price ON products(price);
CREATE INDEX idx_customers_email ON customers(email);
CREATE INDEX idx_orders_customer_id ON orders(customer_id);
CREATE INDEX idx_order_items_order_product ON order_items(order_id, product_id);
