CREATE TABLE IF NOT EXISTS insert_crypto_reports (
  id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  date_created TIMESTAMP,
  success TINYINT(1) NOT NULL,
  db_process_list INT NOT NULL
)
