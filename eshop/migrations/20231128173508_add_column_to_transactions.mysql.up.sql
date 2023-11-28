ALTER TABLE transactions
  ADD COLUMN expiry_month INT DEFAULT 0,
  ADD COLUMN expiry_year INT DEFAULT 0;
