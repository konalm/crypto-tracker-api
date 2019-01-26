CREATE TABLE IF NOT EXISTS fetch_crypto_data_reports (
  id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  third_party_api VARCHAR(255) NOT NULL,
  date_created TIMESTAMP,
  request_status INT NOT NULL,
  notes LONGTEXT NOT NULL
)
