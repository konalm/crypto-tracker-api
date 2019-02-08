CREATE TABLE wallet_states (
  id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  wallet_id INT NOT NULL,
  wallet_state_json LONGTEXT NOT NULL,
  transaction_id INT NOT NULL
)
