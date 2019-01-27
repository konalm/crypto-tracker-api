package reports

import (
  "fmt"
  "stelita-api/db"
)

/**
 *
 */
func InsertFetchRankedCryptoCurrenciesReports(requestStatus int, notes string) {
  fmt.Println("update ranked crypto currencies report !!")

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
