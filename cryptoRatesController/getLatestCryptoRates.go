package cryptoRatesController

import (
  // "database/sql"
  _ "github.com/go-sql-driver/mysql"
  // "fmt"
  "stelita-api/structs"
  "stelita-api/db"
  "stelita-api/errorReporter"
)


/**
 *
 */
func GetCryptoCurrencyRatesForRsi(currency string) []structs.CryptoRate {
  dbConn := db.Conn()
  defer dbConn.Close()

  query :=
    `SELECT date, closing_price, min
    FROM crypto_rates
    WHERE currency = ?
      AND date > (NOW() - INTERVAL 16 DAY)`

  rows, err := dbConn.Query(query, currency)
  if err != nil {
    panic(err.Error())
  }
  defer rows.Close()

  var cryptoRates []structs.CryptoRate

  for rows.Next() {
    var cryptoRate structs.CryptoRate

    err := rows.Scan(&cryptoRate.Date, &cryptoRate.ClosingPrice, &cryptoRate.Min)
    if err != nil {
      errorReporter.ReportError("Getting latest crypto currency rates from the DB")
      panic(err.Error())
    }

    cryptoRates = append(cryptoRates, cryptoRate)
  }

  return cryptoRates
}
