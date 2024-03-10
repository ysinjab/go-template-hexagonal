CREATE TABLE IF NOT EXISTS payment(
   id VARCHAR PRIMARY KEY,
   customer_id VARCHAR NOT NULL,
   status VARCHAR NOT NULL,
   created_at TIMESTAMP NOT NULL,
   order_id VARCHAR NOT NULL,
   total_price DECIMAL NOT NULL
);
