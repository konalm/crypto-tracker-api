package reports

import (
  "stelita-api/db"
)

/**
 *
 */
func InsertUpdateCryptoTrendStatReport(
  cryptoCurrency string, success bool, dbProcessList int,
) {
  db := db.Conn()
  defer db.Close()

  query :=
    `INSERT INTO update_crypto_trend_stat_reports
    (crypto_currency, success, db_process_list)
    VALUES (?,?,?)`

  stmt, err := db.Prepare(query)
  if err != nil {
    panic(err.Error())
  }

  _, err = stmt.Exec(cryptoCurrency, success, dbProcessList)
  if err != nil {
    panic(err.Error())
  }

  defer stmt.Close()
}
