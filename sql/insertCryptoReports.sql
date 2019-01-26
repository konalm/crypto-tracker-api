CREATE TABLE IF NOT EXISTS insert_crypto_reports (
  id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  crypto_currency VARCHAR(255) NOT NULL
  date_created TIMESTAMP,
  success TINYINT(1) NOT NULL
)
