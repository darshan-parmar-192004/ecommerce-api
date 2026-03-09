-- Products filtering by category
CREATE INDEX idx_products_category_id
ON products(category_id);

-- Products price range queries
CREATE INDEX idx_products_price
ON products(price);

-- Customer login lookup
CREATE INDEX idx_customers_email
ON customers(email);

-- Orders per customer
CREATE INDEX idx_orders_customer_id
ON orders(customer_id);

-- Composite index for order items
CREATE INDEX idx_order_items_order_product
ON order_items(order_id, product_id);