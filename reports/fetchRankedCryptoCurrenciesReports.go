package reports

import (
  "stelita-api/db"
)

/**
 *
 */
func InsertFetchRankedCryptoCurrenciesReports(requestStatus int, notes string) {
  db := db.Conn()
  defer db.Close()

  query :=
    `INSERT INTO fetch_ranked_crypto_currencies_reports
    (request_status, notes)
    VALUES (?,?)`

  stmt, err := db.Prepare(query)
  if err != nil {
    panic(err.Error())
  }

  _, err = stmt.Exec(requestStatus, notes)
  if err != nil {
    panic(err.Error())
  }

  defer stmt.Close()
}
