CREATE TABLE update_crypto_trend_stat_reports (
  id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  crypto_currency VARCHAR(255) NOT NULL,
  date_created TIMESTAMP,
  success TINYINT(1) NOT NULL,
  db_process_list INT NOT NULL
)
