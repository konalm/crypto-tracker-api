package cryptoRatesController

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  // "fmt"
  "stelita-api/structs"
)


func GetCryptoCurrencyRatesForRsi(currency string) []structs.CryptoRate {
  /* open database connection */
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stelita_dev")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  query :=
    `SELECT date, closing_price, min
    FROM crypto_rates
    WHERE currency = ?
      AND date > (NOW() - INTERVAL 16 DAY)`

  rows, err := db.Query(query, currency)
  if err != nil {
    panic(err.Error())
  }
  defer rows.Close()

  var cryptoRates []structs.CryptoRate

  for rows.Next() {
    var cryptoRate structs.CryptoRate

    err := rows.Scan(&cryptoRate.Date, &cryptoRate.ClosingPrice, &cryptoRate.Min)
    if err != nil {
      panic(err.Error())
    }

    cryptoRates = append(cryptoRates, cryptoRate)
  }

  return cryptoRates
}
