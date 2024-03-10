CREATE TABLE IF NOT EXISTS orders(
   id VARCHAR PRIMARY KEY,
   customer_id VARCHAR NOT NULL,
   ordered_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS order_details(
    id VARCHAR PRIMARY KEY,
    order_id VARCHAR NOT NULL,
    product_id VARCHAR NOT NULL,
    unit_price DECIMAL NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id)
);

