CREATE TABLE wallets (
  id CHAR(36) PRIMARY KEY,
  user_id VARCHAR(255) NOT NULL,
  date_created TIMESTAMP
)
