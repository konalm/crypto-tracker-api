package reports

import (
  "stelita-api/db"
)

/**
 *
 */
func InsertCryptoReport(
  name string, success bool, dbProcessList int,
) {
  db := db.Conn()
  defer db.Close()

  query :=
    `INSERT INTO insert_crypto_reports
    (crypto_currency, success, db_process_list)
    VALUES (?,?,?)`

  stmt, err := db.Prepare(query)
  if err != nil {
    panic(err.Error())
  }

  _, err = stmt.Exec(name, success, dbProcessList)
  if err != nil {
    panic(err.Error())
  }

  defer stmt.Close()
}
