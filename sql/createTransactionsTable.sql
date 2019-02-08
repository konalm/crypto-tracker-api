CREATE TABLE transactions (
  id INT(6) AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  withdrawal_id INT NOT NULL,
  deposit_id INT NOT NULL,
  exchange_rate FLOAT NOT NULL
)
