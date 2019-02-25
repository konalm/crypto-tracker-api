CREATE TABLE wallet_states (
  id CHAR(36) PRIMARY KEY,
  wallet_id CHAR(36) NOT NULL,
  wallet_state_json LONGTEXT NOT NULL,
  transaction_id CHAR(36) NOT NULL,
  date_created TIMESTAMP
)
