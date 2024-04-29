-- Add foreign key constraint for Customer
ALTER TABLE orders
ADD CONSTRAINT fk_orders_customer
    FOREIGN KEY (customer_id)
    REFERENCES users (user_id);

CREATE UNIQUE INDEX unique_email
		ON users ( email);