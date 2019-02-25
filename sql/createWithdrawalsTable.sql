CREATE TABLE withdrawals (
  id CHAR(36) PRIMARY KEY,
  crypto_currency VARCHAR(255),
  amount FLOAT NOT NULL,
  date_created TIMESTAMP
)
