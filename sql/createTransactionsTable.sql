CREATE TABLE transactions (
  id CHAR(36) PRIMARY KEY,
  user_id INT NOT NULL,
  withdrawal_id char(36) NULL,
  deposit_id char(36) NOT NULL,
  exchange_rate FLOAT NOT NULL,
  date_created TIMESTAMP
)
