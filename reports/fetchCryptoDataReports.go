package reports

import (
  "fmt"
  "stelita-api/db"
)

/**
 *
 */
func InsertFetchCryptoDataReport(
  thirdPartyApi string, requestStatus int, requestMessage string,
) {
  fmt.Println("insert fetch crypto data report !!")

  db := db.Conn()
  defer db.Close()

  query :=
    `INSERT INTO fetch_crypto_data_reports
    (third_party_api, request_status, notes)
    VALUES(?,?,?)`

  stmt, err := db.Prepare(query)
  if err != nil {
    panic(err.Error())
  }

  _, err = stmt.Exec(thirdPartyApi, requestStatus, requestMessage)
  if err != nil {
    panic(err.Error())
  }

  defer stmt.Close()
}
