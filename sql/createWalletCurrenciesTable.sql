CREATE TABLE wallet_currencies (
  id CHAR(36) PRIMARY KEY,
  wallet_id CHAR(36) NOT NULL,
  currency VARCHAR(255) NOT NULL,
  amount FLOAT NOT NULL
)
