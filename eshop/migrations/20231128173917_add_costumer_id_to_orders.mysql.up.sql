ALTER TABLE orders
  ADD COLUMN customer_id INT NOT NULL,
  ADD CONSTRAINT fk_orders_customers
    FOREIGN KEY (customer_id)
    REFERENCES customers(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE;
